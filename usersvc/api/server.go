package api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
)

const (
	redisPassword = ""
	redisHost     = "redis-cluster.default.svc"
	redisPort     = "6379"
)

type Server struct {
	mux *http.ServeMux
	ctx context.Context
	rdb *redis.ClusterClient
}

func NewServer() *Server {
	return &Server{}
}

// Init - configures the server
func (s *Server) Init(ctx context.Context) {
	s.mux = http.NewServeMux()
	s.ctx = ctx
	s.rdb = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{fmt.Sprintf("%s:%s", redisHost, redisPort)},
		Password: redisPassword,
	})
	s.initRoutes()
}

// Run starts the server on the port given in addr
func (s *Server) Run(addr string) {
	fmt.Printf("usersvc is up! listening on port: %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, s.mux))
}
