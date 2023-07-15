autogen:
# Auth
	protoc	--go_out=./internal/auth/grpc/pb/ \
			--go-grpc_out=./internal/auth/grpc/pb/ \
			--grpc-gateway_out=./internal/auth/grpc/pb/ \
			./internal/auth/grpc/proto/*
# Users
	protoc	--go_out=./internal/users/grpc/pb/ \
			--go-grpc_out=./internal/users/grpc/pb/ \
			--grpc-gateway_out=./internal/users/grpc/pb/ \
			./internal/users/grpc/proto/*
# Orders
	protoc	--go_out=./internal/orders/grpc/pb/ \
			--go-grpc_out=./internal/orders/grpc/pb/ \
			--grpc-gateway_out=./internal/orders/grpc/pb/ \
			./internal/orders/grpc/proto/*

build:
	go build -o bin/server cmd/main.go

run:
	go run cmd/main.go

compile:
# Linux
	GOOS=linux GOARCH=amd64 go build -o bin/main-linux-amd64 cmd/main.go
# MacOS
	GOOS=darwin GOARCH=amd64 go build -o bin/main-darwin-amd64 cmd/main.go
# Windows
	GOOS=windows GOARCH=amd64 go build -o bin/main-windows-amd64 cmd/main.go

clean:
	go clean
	rm -rf ./bin