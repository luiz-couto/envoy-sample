FROM golang:1.16.3-alpine3.13 AS build

WORKDIR /envoy-sample
COPY go.mod go.mod
RUN go mod download
COPY . .

WORKDIR /envoy-sample/usersvc
RUN go build

FROM alpine
WORKDIR /deploy
COPY --from=build envoy-sample/usersvc usersvc
ENTRYPOINT ./usersvc