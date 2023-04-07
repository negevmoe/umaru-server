FROM golang:1.18-alpine AS builder

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update && apk add --no-cache musl-dev gcc build-base

WORKDIR /build

ENV CGO_ENABLED=1

COPY go.mod ./go.mod
COPY go.sum ./go.sum
COPY main.go ./main.go
COPY application ./application

RUN GOOS=linux GOPROXY=https://goproxy.cn go build -ldflags="-s -w" -trimpath -o app .


FROM alpine

ARG MODE=dev

# 时区
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

COPY --from=builder /build/app /
COPY dist/ /dist/

ENV SERVER_PORT=8081
ENV SERVER_TOKEN_EXPIRATION_TIME=36000
ENV SERVER_USERNAME="admin"
ENV SERVER_PASSWORD="admin"
ENV SERVER_LOG_DIR="/logs"
ENV DB_PATH="/umaru.db"
ENV DB_MAX_CONNS=30
ENV MEDIA_PATH="/media"
ENV SOURCE_PATH="/downloads"
ENV QB_DOWNLOAD_PATH="/downloads"
ENV QB_URL="http://localhost:7999"
ENV QB_USERNAME="admin"
ENV QB_PASSWORD="adminadmin"
ENV QB_CATEGORY="umaru"
ENV QB_RSS_FOLDER="umaru"

ENTRYPOINT ["/app"]