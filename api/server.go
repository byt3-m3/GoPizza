package api

import (
	"GoPizza/api/handlers"
	"GoPizza/repo"
	"github.com/gorilla/mux"
	"github.com/pmoule/go2hal/hal"
	"log/slog"
	"net/http"
)

type ServerConfig struct {
	ListenAddress string
}

type Server struct {
	Config       *ServerConfig
	Mux          *mux.Router
	CustomerRepo repo.CustomerRepo
}

var (
	v1CurrieLink, _ = hal.NewCurieLink("v1", "https://server.com/v1/{rel}")
	v2CurrieLink, _ = hal.NewCurieLink("v2", "https://server.com/v2/{rel}")

	LinkCuries = []*hal.LinkObject{
		v1CurrieLink,
		v2CurrieLink,
	}
)

func (s *Server) Run() error {

	// Define Routes
	s.Mux.HandleFunc("/v1/customer", func(writer http.ResponseWriter, request *http.Request) {
		handlers.CustomerEndpointHandler(writer, request, &handlers.CustomerEndpointHandlerInput{
			Curries:      LinkCuries,
			CustomerRepo: s.CustomerRepo,
		})
	})

	slog.Info("starting server",
		slog.String("ListenAddress", s.Config.ListenAddress),
	)

	return http.ListenAndServe(s.Config.ListenAddress, s.Mux)

}
