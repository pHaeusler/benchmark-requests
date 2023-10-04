FROM golang:1.17 AS build

WORKDIR /app

COPY src/go.mod ./
RUN go mod download

COPY src/* .

RUN go build server.go

EXPOSE 80

CMD ["/app/server"]
