package main

import (
	"bytes"
	"context"
	"io/ioutil"

	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"inwinstack/cgmh/apiserver/pkg/db"
	"inwinstack/cgmh/apiserver/pkg/ldap"
	model "inwinstack/cgmh/apiserver/pkg/models"
	"inwinstack/cgmh/apiserver/pkg/router"
	"inwinstack/cgmh/apiserver/pkg/services"

	"github.com/gin-contrib/cors"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	readTimeout    = 10 * time.Second
	writeTimeout   = 10 * time.Second
	initRetryDelay = 5 * time.Second
	maxHeaderBytes = 1 << 20
)

var (
	config string
)

func parseFlags() {
	flag.StringVarP(&config, "config", "", "", "Absolute path to the config file.")
	flag.Parse()
}

func loadConfig() {
	if config == "" {
		log.Fatal("Must be set config path from the flag.")
	}

	content, err := ioutil.ReadFile(config)
	if err != nil {
		log.Fatalf("Can't not load config: %+v.", err)
	}
	viper.SetConfigType("yaml")
	viper.ReadConfig(bytes.NewBuffer(content))
}

func initDatabase() *db.Mongo {
	log.Printf("Connecting database...")
	f := &db.Flag{
		Host:     viper.GetString("db.host"),
		Source:   viper.GetString("db.source"),
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
		DB:       viper.GetString("db.name"),
	}

	// Wait for connecting database
	for {
		database, err := db.New(f)
		if err == nil {
			return database
		}
		log.Printf("Failed to connect database. %+v. retrying...", err)
		<-time.After(initRetryDelay)
	}
}

func main() {
	parseFlags()
	// Load config to viper
	loadConfig()
	log.SetFlags(log.LstdFlags)

	ldapFlag := &ldap.Flag{
		Protocol: viper.GetString("ldap.protocol"),
		Host:     viper.GetString("ldap.host"),
		Username: viper.GetString("ldap.bind.username"),
		Password: viper.GetString("ldap.bind.password"),
		DN:       viper.GetString("ldap.dn"),
		OU:       viper.GetString("ldap.userOU"),
	}
	ldapServer := ldap.NewServer(ldapFlag)

	db := initDatabase()
	svc := service.New(db)
	if err := svc.CreateConfig(); err != nil {
		log.Fatal("Server creating db config error:", err)
	}

	// Init the default levels
	if !svc.AlreadyInitLevel() {
		if err := svc.InitLevels(viper.Get("levels")); err != nil {
			log.Fatal("Server initing levels error:", err)
		}
	}

	// Init admin user
	if !svc.AlreadyInitAdmin() {
		pwd := viper.GetString("admin.password")
		reg := &model.User{
			Email:    viper.GetString("admin.email"),
			Name:     viper.GetString("admin.name"),
			Agency:   viper.GetString("admin.agency"),
			Unit:     viper.GetString("admin.unit"),
			JobTitle: viper.GetString("admin.jobTitle"),
			Phone:    viper.GetString("admin.phone"),
		}
		if err := svc.InitAdmin(reg, pwd); err != nil {
			log.Fatal("Server initing admin error:", err)
		}

		if err := ldapServer.AddOU(ldapFlag.OU, "User account OU."); err != nil {
			log.Fatal("LDAP creating user account OU error:", err)
		}

		if err := ldapServer.AddUser(reg, pwd); err != nil {
			log.Fatal("LDAP creating admin account error:", err)
		}
	}

	r := router.New(svc)
	server := &http.Server{
		Addr:           viper.GetString("global.listen"),
		Handler:        r.GetEngine(),
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	handler := r.GetHandler()
	handler.SetLDAP(ldapServer)

	r.LinkSwaggerAPI(viper.GetBool("global.swagger"))
	r.LinkHandlers()

	if viper.GetStringSlice("global.allowOrigins") != nil {
		config := cors.Config{
			AllowOrigins:     viper.GetStringSlice("global.allowOrigins"),
			AllowMethods:     []string{"GET", "POST", "PUT", "HEAD"},
			AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: false,
			MaxAge:           12 * time.Hour,
		}
		r.SetCORS(config)
	}

	go func() {
		log.Println("API server starting...")
		if err := server.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown:", err)
	}
	log.Println("Server exiting...")
}
