# Microservices application

Simple application with microservice architecture pattern on Go.

## Description
### **Coffee shop web-application.**
There are Orders and Users services that realizes CRUD operations. Both services connected to each other via gRPC and calling externally via REST through API Gateway.

## Structure
Application architecture:
![map](https://i.ibb.co/mT5MvWY/Frame-6-2.jpg)

You can see all available API methods in `/docs/documentation.md`. \
The folders structure was inspired by [Standard Go Project Layout](https://github.com/golang-standards/project-layout) and persist in each service. Here are:
- `cmd`: contains entry point of application
- `config`: stores all app configurations
- `docs`: complete documentation of app
- `proto`: gRPC definitions for microservices
- `pb`: generated code from protobuffers
- `internal`: implementation of microservices methods
- `api`: some REST requests examples

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