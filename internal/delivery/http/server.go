package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"rest-api/internal/config"
	"rest-api/internal/delivery/http/user_controller"
	"rest-api/pkg/utils"
)

func fatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// StartServerHttp StartServer initializes and starts the HTTP server.
func StartServerHttp(port string) {
	// Set up the router
	router := mux.NewRouter()

	router.Use(utils.LoggingMiddleware)

	http.Handle("/", router)
	cfg := config.LoadConfigServer()
	// controllers
	user_controller.UserController(router, cfg)
	// Start the server
	serverAddr := fmt.Sprintf(":%s", port)
	fmt.Printf("Server is starting on http://localhost%s\n", serverAddr)
	err := http.ListenAndServe(serverAddr, router)
	if err != nil {
		fatal(errors.Wrap(err, "Error starting server:"))
	}
}
