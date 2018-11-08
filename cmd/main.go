package main

import (
	"context"

	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"inwinstack/cgmh/apiserver/pkg/models"
	"inwinstack/cgmh/apiserver/pkg/router"
	"inwinstack/cgmh/apiserver/pkg/util"

	flag "github.com/spf13/pflag"
)

const (
	readTimeout    = 10 * time.Second
	writeTimeout   = 10 * time.Second
	maxHeaderBytes = 1 << 20
)

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func initAdminUser(init bool) {
	if init {
		hex, err := util.RandomHex(8)
		if err != nil {
			log.Fatal("Server initing error:", err)
		}

		pwdStr := getenv("INIT_ADMIN_PASSWORD", hex)
		pwd := &models.Password{Secret: util.Base64Encode(pwdStr)}
		user := &models.User{
			Email:   getenv("INIT_ADMIN_EMAIL", "admin@example.com"),
			IsAdmin: true,
			Active:  true,
		}

		dao := &models.User{}
		if !dao.IsExistByEmail(user.Email) {
			log.Println("Server initing admin...")
			if err := dao.Insert(user, pwd); err != nil {
				log.Fatal("Server initing error:", err)
			}
			log.Printf("Admin init email: %s", user.Email)
			log.Printf("Admin init password: %s", pwdStr)
		}
	}
}

func main() {
	addr := flag.StringP("listen-addr", "", ":8080", "API server listen address.")
	init := flag.BoolP("init", "", true, "Init admin user.")
	flag.Parse()

	r := router.NewRouter()
	s := &http.Server{
		Addr:           *addr,
		Handler:        r,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.SetFlags(log.LstdFlags)
	log.SetPrefix("[SERVER-debug] ")

	initAdminUser(*init)
	go func() {
		log.Println("Server starting...")
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
