package rest

import (
	"fmt"
	"net/http"
	"strconv"
	"sysagent/config"
	"sysagent/rest/handlers/product"
	"sysagent/rest/handlers/user"
	middleware "sysagent/rest/middlewares"
)

type Server struct {
	cnf            *config.Config
	productHandler *product.Handler
	userHandler    *user.Handler
}

func NewServer(
	cnf *config.Config,
	productHandler *product.Handler,
	userHandler *user.Handler,
) *Server {
	return &Server{
		cnf:            cnf,
		productHandler: productHandler,
		userHandler:    userHandler,
	}
}

func (server *Server) Start() {
	manager := middleware.NewManager()

	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux()

	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)

	wrapppedMux := manager.WrapMux(mux)

	addr := ":" + strconv.Itoa(server.cnf.HttpPort)
	fmt.Println("Server starting on port", addr)
	err := http.ListenAndServe(addr, wrapppedMux)
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}
