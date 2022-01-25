# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.17 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /goapi

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /goapi /goapi

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/goapi"]
