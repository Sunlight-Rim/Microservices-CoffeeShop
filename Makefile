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