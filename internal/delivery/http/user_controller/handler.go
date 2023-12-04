package user_controller

import (
	"github.com/gorilla/mux"
	"rest-api/internal/config"
)

type userController struct {
	cfg *config.ConfigServer
}

func UserController(r *mux.Router, config *config.ConfigServer) {
	handler := &userController{
		cfg: config,
	}
	r.HandleFunc("/user/info", handler.GetInfoUser).Methods("GET")
}
