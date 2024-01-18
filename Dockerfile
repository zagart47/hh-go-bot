FROM golang:latest-bullseye as builder
RUN apt-get update && apt-get install -y ca-certificates openssl
ARG cert_location=/usr/local/share/ca-certificates
RUN openssl s_client -showcerts -connect github.com:443 </dev/null 2>/dev/null|openssl x509 -outform PEM > ${cert_location}/github.crt
RUN openssl s_client -showcerts -connect proxy.golang.org:443 </dev/null 2>/dev/null|openssl x509 -outform PEM >  ${cert_location}/proxy.golang.crt
RUN update-ca-certificates
WORKDIR /build
COPY go.mod .
RUN go mod download
COPY .. .
RUN env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main cmd/hh-go-bot/main.go

FROM alpine:latest
COPY internal/config/config.yaml internal/config/
COPY --from=builder main /main
ENTRYPOINT ["/main"]
