package main

import (
	"log"
	"github.com/armon/go-socks5"
	"os"
)

func main() {
	log.Println("SOCKS5 server: https://github.com/Egregors/socks5-server")

	// todo: make normal auth. from ENV?
	user := os.Args[1]
	password := os.Args[2]

	var users socks5.StaticCredentials

	users = socks5.StaticCredentials{
		user: password,
	}

	auth := socks5.UserPassAuthenticator{Credentials: users}

	log.Println("Configuration..")
	srvConfig := &socks5.Config{
		AuthMethods: []socks5.Authenticator{auth},
		Logger:      log.New(os.Stdout, "", log.LstdFlags),
	}

	srv, err := socks5.New(srvConfig)
	if err != nil {
		log.Panic(err.Error())
	}

	log.Println("Start server")
	srv.ListenAndServe("tcp", "0.0.0.0:1111")
}
