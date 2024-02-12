FROM golang:alpine AS build

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o app .

FROM alpine:latest

WORKDIR /root/

COPY --from=build /app/app .

EXPOSE 8080

ENTRYPOINT ["/bin/sh", "-c", "./app"]