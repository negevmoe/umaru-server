##########
# 全局变量 #
##########
NAME	= PROJECT_NAME
VERSION = PROJECT_VERSION

# 指定proto_path
PROTO_PATH := 	--proto_path=./proto \
				--proto_path=$(subst \,/,$(GOPATH)) \

# 生成PB文件
PROTO_PB 		:= $(PROTO_PATH) --go_out=./proto/pb ./proto/message.proto
# PB文件参数校验
PROTO_VALIDATE 	:= $(PROTO_PATH) --validate_out=lang=go:./proto/pb ./proto/message.proto
# GRPC
PROTO_GRPC 		:= $(PROTO_PATH) --go-grpc_out=./proto/pb --go-grpc_opt=require_unimplemented_servers=false ./proto/service.proto
# RESTFUL
PROTO_GATEWAY 	:= $(PROTO_PATH) --grpc-gateway_out=logtostderr=true:./proto/pb ./proto/service.proto

.PHONY:proto
proto:
	@protoc $(PROTO_PB)
	@protoc $(PROTO_VALIDATE)
	@protoc $(PROTO_GRPC)
	@protoc $(PROTO_GATEWAY)

.PHONY:run
	@go run main.go

.PHONY:build
build:
	@GOOS=linux go build -ldflags="-s -w" -trimpath -o ./dist/$(NAME)-$(VERSION) ./main.go


.PHONY:run
run:
	@\
	SERVER_DEBUG=true \
    SERVER_PORT=8081 \
    SERVER_TOKEN_EXPIRATION_TIME=36000 \
    SERVER_USERNAME="admin" \
    SERVER_PASSWORD="admin" \
    SERVER_LOG_DIR="D:/home/negevmoe/umaru-server/.log" \
	DB_PATH="D:/home/negevmoe/umaru-server/umaru.db" \
	DB_MAX_CONNS=30 \
    MEDIA_PATH="D:/home/negevmoe/umaru-server/docker/jellyfin/media" \
    SOURCE_PATH="D:/home/negevmoe/umaru-server/docker/qbittorrent/downloads" \
    QB_DOWNLOAD_PATH="/downloads" \
    QB_URL="http://localhost:9999" \
    QB_USERNAME="admin" \
    QB_PASSWORD="adminadmin" \
    QB_CATEGORY="umaru" \
    QB_RSS_FOLDER="umaru" \
    go run main.go