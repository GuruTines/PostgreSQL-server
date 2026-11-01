package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/internal/models"
	"server/internal/service"
)

type ServerAPI struct {
	service *service.ServerService
}

func NewServerAPI(service *service.ServerService) *ServerAPI {
	return &ServerAPI{
		service: service,
	}
}

func (s *ServerAPI) ExecRequest(w http.ResponseWriter, r *http.Request) {
	var req models.ServerRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&req); err != nil {
		fmt.Printf("Ошибка декодирования запроса!\n%v\n", err)
		return
	}

	if err := s.service.ExecRequest(&req); err != nil {
		return
	}
}
