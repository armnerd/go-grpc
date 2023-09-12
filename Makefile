pb:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    internal/proto/blog.proto

wire:
	cd cmd && wire

all:
	-make pb;
	-make wire;

run:
	cd cmd && go run . -env dev

build:
	cd cmd && go build .