package main

import (
	"fmt"
	"log"

	"github.com/c-pro/go-socks5"
	"github.com/kelseyhightower/envconfig"

	"os"
)

type config struct {
	Host  string `default:"0.0.0.0"`
	Port  int    `default:"1080"`
	Users socks5.StaticCredentials
}

func main() {
	log.Println("SOCKS5 server: https://github.com/Egregors/socks5-server")

	var cfg config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("Failed to read ENV: %s\n", err.Error())
	}

	if len(cfg.Users) < 1 {
		log.Fatal("No user credentials specified in USERS environment variable")
	}

	auth := socks5.UserPassAuthenticator{Credentials: cfg.Users}

	log.Println("Configuration..")
	srvConfig := &socks5.Config{
		AuthMethods: []socks5.Authenticator{auth},
		Logger:      log.New(os.Stdout, "", log.LstdFlags),
	}

	srv, err := socks5.New(srvConfig)
	if err != nil {
		log.Panic(err.Error())
	}

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	log.Printf("Starting server on %s\n", addr)
	if err := srv.ListenAndServe("tcp", addr); err != nil {
		log.Fatalf("Failed to start socks server on %s: %s\n", addr, err)
	}
}
