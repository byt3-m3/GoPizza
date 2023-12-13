package main

import (
	"GoPizza/api"
	"GoPizza/repo"
	"github.com/gorilla/mux"
)

func main() {
	s := api.Server{
		Config: &api.ServerConfig{
			ListenAddress: "0.0.0.0:8080",
		},
		Mux:          &mux.Router{},
		CustomerRepo: repo.NewLocalCustomerRepo(),
	}

	if err := s.Run(); err != nil {
		panic(err)
	}
}
