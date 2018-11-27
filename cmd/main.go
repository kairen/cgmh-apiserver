package main

import (
	"context"

	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"inwinstack/cgmh/apiserver/pkg/dao"
	"inwinstack/cgmh/apiserver/pkg/db"
	"inwinstack/cgmh/apiserver/pkg/models"
	"inwinstack/cgmh/apiserver/pkg/router"
	"inwinstack/cgmh/apiserver/pkg/util"

	"github.com/gin-contrib/cors"
	flag "github.com/spf13/pflag"
)

const (
	readTimeout    = 10 * time.Second
	writeTimeout   = 10 * time.Second
	initRetryDelay = 5 * time.Second
	maxHeaderBytes = 1 << 20
)

var (
	address    string
	origins    []string
	enableCORS bool
	initAdmin  bool
	swagger    bool

	// Database infos
	dbHost   string
	dbSource string
	dbUser   string
	dbPasswd string
	dbName   string
)

func parseFlags() {
	flag.StringVarP(&address, "listen-addr", "", ":8080", "API server listen address.")
	flag.StringSliceVarP(&origins, "allow-origins", "", nil, "List of allowed origins for CORS, comma separated.")
	flag.BoolVarP(&initAdmin, "init", "", true, "Init admin user.")
	flag.BoolVarP(&swagger, "enable-swagger", "", true, "Set to enable/disable swagger API.")
	flag.StringVarP(&dbHost, "db-host", "", "127.0.0.1:27017", "Database host address.")
	flag.StringVarP(&dbSource, "db-source", "", "admin", "Database source name.")
	flag.StringVarP(&dbUser, "db-user", "", "root", "Database user name.")
	flag.StringVarP(&dbPasswd, "db-password", "", "", "Database user password.")
	flag.StringVarP(&dbName, "db-name", "", "CGMH", "Database name.")
	flag.Parse()
}

func initDatabase() *db.Database {
	log.Printf("Connecting database...")
	f := &db.Flag{
		Host:     dbHost,
		Source:   dbSource,
		User:     dbUser,
		Password: dbPasswd,
		DB:       dbName,
	}

	// Wait for Connecting database
	for {
		database, err := db.New(f)
		if err == nil {
			return database
		}
		log.Printf("Failed to connect database. %+v. retrying...", err)
		<-time.After(initRetryDelay)
	}
}

func initAdminUser(dao *dao.DataAccess) {
	if initAdmin {
		hex, err := util.RandomHex(8)
		if err != nil {
			log.Fatal("Server initing error:", err)
		}

		pwd := util.GetEnv("INIT_ADMIN_PASSWORD", hex)
		secret := util.MD5Encode(pwd)
		reg := &models.User{
			Email: util.GetEnv("INIT_ADMIN_EMAIL", "admin@inwinstack.com"),
			Name:  "administrator",
		}

		if !dao.User.IsExistByEmail(reg.Email) {
			log.Println("Server initing admin...")
			if err := dao.Auth.Register(reg, secret); err != nil {
				log.Fatal("Server initing error:", err)
			}

			user, err := dao.User.FindByEmail(reg.Email)
			if err != nil {
				log.Fatal("Server initing error:", err)
			}

			stat := &models.UserStatus{UserUUID: user.UUID, Block: false, Active: true}
			if err := dao.User.UpdateStatus(stat); err != nil {
				log.Fatal("Server initing error:", err)
			}

			role := &models.UserRole{UserUUID: user.UUID, Name: models.RoleAdmin}
			if err := dao.User.UpdateRole(role); err != nil {
				log.Fatal("Server initing error:", err)
			}

			log.Printf("Admin init email: %s", reg.Email)
			log.Printf("Admin init password: %s", pwd)
		}
	}
}

func main() {
	parseFlags()
	log.SetFlags(log.LstdFlags)

	db := initDatabase()
	dao := dao.New(db)
	r := router.New(dao)
	s := &http.Server{
		Addr:           address,
		Handler:        r.GetEngine(),
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	if origins != nil {
		config := cors.Config{
			AllowOrigins:     origins,
			AllowMethods:     []string{"GET", "POST", "PUT", "HEAD"},
			AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: false,
			MaxAge:           12 * time.Hour,
		}
		r.SetCORS(config)
	}

	// Init admin user and handlers
	initAdminUser(dao)
	r.LinkSwaggerAPI(swagger)
	r.LinkHandlers()

	go func() {
		log.Println("API server starting...")
		if err := s.ListenAndServe(); err != nil {
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
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown:", err)
	}
	log.Println("Server exiting...")
}
