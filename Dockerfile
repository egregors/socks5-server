FROM golang as builder
RUN mkdir -p /go/src/github.com/C-Pro/socks5-server
WORKDIR /go/src/github.com/C-Pro/socks5-server
COPY . .

RUN go get -u github.com/golang/dep/...
RUN dep ensure

RUN go build -ldflags "-linkmode external -extldflags -static" -a s5.go

FROM scratch
COPY --from=builder /go/src/github.com/C-Pro/socks5-server/s5 ./

ENV HOST=0.0.0.0
ENV PORT=1111
EXPOSE 1111
CMD ["./s5"]
