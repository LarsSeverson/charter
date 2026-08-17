PROTO_ROOT := api/internal
GEN_ROOT := gen

USER_PROTO := grpc/user/v1/user.proto

PROTOC_GEN_GO_VERSION := v1.36.12
PROTOC_GEN_GO_GRPC_VERSION := v1.6.2

generate: proto
	go -C $(GEN_ROOT) mod tidy

proto: proto-user

proto-user:
	protoc --proto_path=$(PROTO_ROOT) --go_out=$(GEN_ROOT) --go_opt=paths=source_relative --go-grpc_out=$(GEN_ROOT) --go-grpc_opt=paths=source_relative $(USER_PROTO)

.PHONY: generate proto proto-user
