FROM golang:1.24.13-alpine AS builder

WORKDIR /app
COPY /. ./

# line 7 to 9 needed to use `github.com/mattn/go-sqlite3`, an indirect dependency of gorm
RUN apk update \
 && apk add build-base \
 && go env -w CGO_ENABLED=1 \
 && go mod download \
 && go build -o /florobot .

FROM alpine:3.16.2
WORKDIR /app

RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/community" > /etc/apk/repositories \
 && echo "http://dl-cdn.alpinelinux.org/alpine/edge/main" >> /etc/apk/repositories \
 && apk update \
 && apk add chromium chromium-chromedriver

COPY --from=builder florobot .

CMD ["./florobot"]
