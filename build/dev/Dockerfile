FROM golang:1.20

ARG DEV_UID
ARG DEV_GID

RUN addgroup --gid $DEV_GID go-user && \
    adduser --disabled-password --uid $DEV_UID --ingroup go-user go-user

RUN apt-get update && \
    apt-get install -y postgresql-client

USER go-user:go-user

RUN go install github.com/cweill/gotests/gotests@v1.6.0 && \
    go install github.com/fatih/gomodifytags@v1.16.0 && \
    go install github.com/josharian/impl@v1.1.0 && \
    go install github.com/haya14busa/goplay/cmd/goplay@v1.0.0 && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install honnef.co/go/tools/cmd/staticcheck@latest && \
    go install golang.org/x/tools/gopls@latest && \
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3 && \
    go install github.com/go-task/task/v3/cmd/task@latest && \
    go install github.com/cosmtrek/air@latest

EXPOSE 80
