##
# checkov:skip=CKV_DOCKER_2:Build stage image
# checkov:skip=CKV_DOCKER_3:Build stage image

FROM golang:alpine AS build

WORKDIR /app

COPY . .

RUN  GOOS=linux GOARCH=arm64 go build -o server

FROM alpine:3.14
LABEL org.opencontainers.image.source https://github.com/codingric/http-request-logger

ENTRYPOINT [ "/server" ]

COPY --from=build /app/server /server

