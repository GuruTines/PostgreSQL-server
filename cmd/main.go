package main

import (
	"net/http"
	"server/internal/api"
	"server/internal/service"
)

func main() {
	strConn := "postgresql://postgres:postgres@localhost:5432/server_db"
	pool, err := service.NewServerService(strConn)
	if err != nil {
		return
	}

	serverAPI := api.NewServerAPI(pool)

	http.HandleFunc("GET /query", serverAPI.ExecRequest)
	if err := http.ListenAndServe(":9091", nil); err != nil {
		return
	}
}
