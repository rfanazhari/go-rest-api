package main

import (
	"fmt"
	"rest-api/internal/config"
	"rest-api/internal/delivery/http"
)

var (
	cfg = config.LoadConfigServer()
)

func main() {
	fmt.Println("hello world", cfg.ServerPort)

	// Call Rest-API
	http.StartServerHttp(cfg.ServerPort)
}
