package main

import (
	"context"
	"envoy-sample/usersvc/api"
)

func main() {
	ctx := context.Background()
	server := api.NewServer()
	server.Init(ctx)
	server.Run("9001")
}
