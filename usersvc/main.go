package usersvc

import "context"

func main() {
	ctx := context.Background()
	server := NewServer()
	server.Init(ctx)
	server.Run("9001")
}
