# socks5-server
Simple socks5 server. You may use it as a proxy for Telegram

![build status](https://img.shields.io/docker/build/cpro29a/socks5-server.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/C-Pro/socks5-server)](https://goreportcard.com/report/github.com/C-Pro/socks5-server)


# Usage

```
docker pull cpro29a/socks5-server
docker run -name socks -d --restart=always -p 1111:1111 -e "USERS=user1:pass1,user2:pass2" cpro29a/socks5-server
```

Do not forget to replace usernames and passwords with your (secure) values! :)

Now you may try to connect to 1111 port of your host

## Build
```
dep ensure
go build -o s5 .
USERS="user1:pass1" ./s5
```

# Contributing
Bug reports, bug fixes and new features are always welcome.
Please open issues and submit pull requests for any new code.
