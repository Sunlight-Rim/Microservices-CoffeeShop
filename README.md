# Microservices application

Simple application with microservice architecture pattern on Go.

## Description
### **Coffee shop web-application.**
You can register, login, order some coffee, view the history of your coffee orders, list other users and do some related actions. To view all available API methods you can go to [Documentation](/docs/documentation.md).

## Structure
Application architecture:
![schema](https://i.imgur.com/lIdaGaP.jpg)

There are Auth authentication service by JWT tokens, Orders and Users services that realizes CRUD operations. All these services are connected to API Gateway via gRPC and are calling externally via REST. Auth also connected to Users for registration. Endpoints and API methods you can see in [Documentation](/docs/documentation.md). \
The folders structure was inspired by [Standard Go Project Layout](https://github.com/golang-standards/project-layout) and persist in each service. Here are:
- `api`: some REST requests examples
- `cmd`: contains entry point of application
- `config`: stores all app configurations
- `docs`: complete documentation of app
- `scripts`: standalone scripts for this app
- `internal/<service>/database`: microservice database files
- `internal/<service>/proto`: gRPC definitions for microservices
- `internal/<service>/pb`: generated code from protobuffers
- `internal/<service>`: implementation of microservices methods

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
