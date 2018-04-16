FROM golang:1.10.1 as builder
RUN mkdir -p /go/src/github.com/Egregors/socks5-server
WORKDIR /go/src/github.com/Egregors/socks5-server
COPY . .

RUN go get -u github.com/golang/dep/...
RUN dep ensure

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o s5

FROM scratch
COPY --from=builder /go/src/github.com/Egregors/socks5-server/s5 ./
COPY --from=builder /go/src/github.com/Egregors/socks5-server/users ./

EXPOSE 1111
CMD ["./s5"]