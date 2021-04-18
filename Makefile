.PHONY: build

build: build-proto

build-proto:
	protoc -I=./pkg/domain --go_out=./pkg/domain --go_opt=paths=source_relative relay.proto
