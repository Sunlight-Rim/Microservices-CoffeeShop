## Build
FROM golang:latest

WORKDIR /gorestapi

COPY . ./

RUN go build -o bin cmd/main.go

## Deploy
ENTRYPOINT ["./bin"]

EXPOSE 8080