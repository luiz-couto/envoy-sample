FROM golang:1.16.3-alpine3.13 AS build

WORKDIR /envoy-sample
COPY go.mod go.mod
RUN go mod download
COPY . .

WORKDIR usersvc
RUN go build

FROM lyft/envoy:latest
WORKDIR /deploy

COPY --from=build envoy-sample/usersvc .
COPY usersvc/envoy.json .
COPY usersvc/entrypoint.sh .

RUN chmod +x entrypoint.sh

CMD ./usersvc
ENTRYPOINT ./entrypoint.sh
