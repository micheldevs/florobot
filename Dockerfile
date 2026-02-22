FROM golang:1.24.13-alpine AS builder

WORKDIR /app
COPY /. ./

RUN go mod download \
 && go build -o /florobot .

FROM alpine:3.16.2
WORKDIR /app

RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/community" > /etc/apk/repositories \
 && echo "http://dl-cdn.alpinelinux.org/alpine/edge/main" >> /etc/apk/repositories \
 && apk update \
 && apk add chromium chromium-chromedriver

COPY --from=builder florobot .

CMD ["./florobot"]
