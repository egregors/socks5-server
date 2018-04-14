package main

import (
	"log"
	"github.com/armon/go-socks5"
	"os"
	"io/ioutil"
	"strings"
)

func readUsersFromFile(path string) map[string]string {
	us, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panic("Can't get users", err.Error())
	}

	users := make(map[string]string)

	for _, l := range strings.Split(string(us), "\n") {
		if l != "" {
			userPass := strings.Split(l, ":")
			users[userPass[0]] = userPass[1]
		}
	}

	return users
}

func main() {
	log.Println("SOCKS5 server: https://github.com/Egregors/socks5-server")
	log.Println("Init users..")
	users := make(socks5.StaticCredentials)
	for u, p := range readUsersFromFile("users") {
		users[u] = p
		log.Printf("User: %s", u)
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
	log.Println("Litsen on 0.0.0.0:1111")
	srv.ListenAndServe("tcp", "0.0.0.0:1111")
}
