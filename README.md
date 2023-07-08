# Microservices application

Simple application with microservice pattern and clean architecture on Go.

## Description
### **Coffee shop**
You can register, login, order some coffee, view the history of your coffee orders, list other users and do some related actions. To view all available API methods you can go to [Documentation](/docs/documentation.md).

## Structure
Application architecture:
![schema](https://i.imgur.com/Z6z7mAs.jpg)

There are Auth authentication service works by JWT tokens; Orders and Users services that realizes CRUD operations. All of these are connected to API Gateway via gRPC and are calling externally via REST. Auth also connected to Users through gRPC for registration functionality. Endpoints and API methods you can see in [Documentation](/docs/documentation.md). \
The folders structure was inspired by [Standard Go Project Layout](https://github.com/golang-standards/project-layout) and persist in each service. Here are:
- `api`: some REST requests examples
- `docs`: complete documentation
- `cmd`: minimalist entry point of application
- `config`: stores all app configurations
- `scripts`: standalone scripts for app
- `internal/<service>/domain`: domain layer with entities of microservice
- `internal/<service>/business`: implementation of microservices business logic
- `internal/<service>/database`: microservice repository layer, contains database files
- `internal/<service>/grpc`: transport layer, contains gRPC server, adapters, protofiles and generated code

`go.mod` and `go.sum` is a Go Modules files.

## Usage and Run
Application startup url is **http://localhost:8080**, socket can be changed in `/config/config.yaml`.

### By hand
```shell
go run cmd/main.go
```
### With Makefile
```shell
make run
```
You can also **build**, **compile** and **clean** program.
### With Docker
```shell
docker build -t coffeeshop .
docker run -p 8080:8080 coffeeshop
```
