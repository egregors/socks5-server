# socks5-server
Simple socks5 server. You may user it as a proxy for Telegram

![size](https://img.shields.io/badge/images%20size-4.47MB-brightgreen.svg)

# Usage

First, you need to add user-pass into `users.example` and rename it to `users`

## bin
```
go build s5
./s5
```

## Docker
```
docker build . -t s5
docker run -it -p 1111:1111 s5
```

Now you may try to connect to 1111 port of your host

# Contributing
Bug reports, bug fixes and new features are always welcome.
Please open issues and submit pull requests for any new code.