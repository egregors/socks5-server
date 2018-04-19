FROM golang as builder
RUN mkdir -p /go/src/github.com/Egregors/socks5-server
WORKDIR /go/src/github.com/Egregors/socks5-server
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o s5

FROM scratch
COPY --from=builder /go/src/github.com/Egregors/socks5-server/s5 ./

ENV HOST=0.0.0.0
ENV PORT=1111
EXPOSE 1111

CMD ["./s5"]
