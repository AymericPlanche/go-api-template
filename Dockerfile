# Compiles the binary
FROM golang:1 as builder
WORKDIR /go/src/myapp
COPY . .
RUN CGO_ENABLED=0 go build -o /myapp main.go

# up to date certificates
FROM alpine:latest as certs
RUN apk --update add ca-certificates
RUN update-ca-certificates

# Main image
FROM scratch
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /myapp /
COPY ./internal/app/templates /usr/local/myapp/internal/app/templates
ENTRYPOINT ["/myapp"]
