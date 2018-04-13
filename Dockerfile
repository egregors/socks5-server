FROM golang:1.10.1 as builder
RUN mkdir -p /go/src/github.com/Egregors/socks5-server
WORKDIR /go/src/github.com/Egregors/socks5-server
COPY . .
RUN go build -ldflags "-linkmode external -extldflags -static" -a s5.go

FROM scratch
COPY --from=builder /go/src/github.com/Egregors/socks5-server/s5 ./
COPY --from=builder /go/src/github.com/Egregors/socks5-server/users ./

EXPOSE 1111
CMD ["./s5"]