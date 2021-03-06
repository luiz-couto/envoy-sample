package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"envoy-sample/rest"
)

// initRoutes defines which routes the UID Provider API will have
func (s *Server) initRoutes() {
	s.mux.HandleFunc("/newuser", s.CreateUserHandler())
	s.mux.HandleFunc("/getuser", s.GetUserHandler())
}

func (s *Server) CreateUserHandler() rest.Handler {
	return rest.Handler(func(w http.ResponseWriter, r *http.Request) {
		data := User{}
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			rest.ERROR(w, err)
			return
		}

		if err := set(context.Background(), s.rdb, data); err != nil {
			rest.ERROR(w, errors.New("fail when creating user"))
			return
		}
		rest.JSON(w, 200, struct {
			Message string
		}{
			Message: fmt.Sprintf("Successfully created %s user", data.Username),
		})
	})
}

func (s *Server) GetUserHandler() rest.Handler {
	return rest.Handler(func(w http.ResponseWriter, r *http.Request) {
		data := GetModel{}
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			rest.ERROR(w, err)
			return
		}

		user, err := get(context.Background(), s.rdb, data.Username)

		if err != nil {
			rest.ERROR(w, errors.New("fail when creating user"))
			return
		}
		rest.JSON(w, 200, user)
	})
}
