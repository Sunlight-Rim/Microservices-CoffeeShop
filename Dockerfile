## Build
FROM golang:latest

WORKDIR /gorestapi

COPY . .

RUN go build -o shop cmd/main.go

## Deploy
EXPOSE 8080

RUN chmod a+x shop

CMD ["./shop"]