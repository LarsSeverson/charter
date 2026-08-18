PROTO_ROOT := api
GEN_ROOT := gen
GEN_MODULE := github.com/LarsSeverson/charter/gen

USER_ROOT := $(PROTO_ROOT)/internal/charter/user/v1
USER_PROTOS := \
               $(USER_ROOT)/user.proto \
               $(USER_ROOT)/user_service.proto

PROTOC_GEN_GO_VERSION := v1.36.12
PROTOC_GEN_GO_GRPC_VERSION := v1.6.2

ifeq ($(OS),Windows_NT)
  RM_RF = cmd /C "if exist \"$(1)\" rmdir /S /Q \"$(1)\""
else
  RM_RF = rm -rf "$(1)"
endif

generate: proto
	go -C $(GEN_ROOT) mod tidy

proto: proto-user

proto-user:
	protoc \
		--proto_path=$(PROTO_ROOT) \
		--go_out=$(GEN_ROOT) \
		--go_opt=module=$(GEN_MODULE) \
		--go-grpc_out=$(GEN_ROOT) \
		--go-grpc_opt=module=$(GEN_MODULE) \
		$(USER_PROTOS)

.PHONY: generate proto proto-user
