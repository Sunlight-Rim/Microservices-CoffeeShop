# Microservices application

Simple application with microservice architecture pattern on Go.

## Description
### **Coffee shop web-application.**
You can register, login, order some coffee, view the history of your coffee orders, list other users and do some related actions. For view all available API methods you can go to [Documentation](/docs/documentation.md).

## Structure
Application architecture:
![schema](https://i.imgur.com/qcioNiX.jpeg)

There are Orders and Users services that realizes CRUD operations. Both services are connected to each other via gRPC and are calling externally via REST through API Gateway. Endpoints and API methods you can see in [Documentation](/docs/documentation.md) \
The folders structure was inspired by [Standard Go Project Layout](https://github.com/golang-standards/project-layout) and persist in each service. Here are:
- `api`: some REST requests examples
- `cmd`: contains entry point of application
- `config`: stores all app configurations
- `docs`: complete documentation of app
- `internal/<service>/proto`: gRPC definitions for microservices
- `internal/<service>/pb`: generated code from protobuffers
- `internal/<service>`: implementation of microservices methods

`go.mod` and `go.sum` is a Go Modules files.

## Usage and Run
Application startup url is **http://localhost:8080**, socket can be changed in `/config/config.yaml`.

### By hand
```
go run cmd/main.go
```
### With Makefile
```
make run
```
You can also **build**, **compile** and **clean** program.
### With Docker
```
docker build -t coffeeshop .
docker run -p 8080:8080 coffeeshop
```
