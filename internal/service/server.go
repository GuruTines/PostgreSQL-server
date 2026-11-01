package service

import (
	"context"
	"fmt"
	"server/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ServerService struct {
	pool *pgxpool.Pool
}

func NewServerService(strConn string) (*ServerService, error) {
	pool, err := pgxpool.New(context.Background(), strConn)
	if err != nil {
		fmt.Printf("Ошибка создания пула\nОписание: %v\n", err)
		return nil, err
	}

	return &ServerService{pool: pool}, nil
}

func (s *ServerService) ExecRequest(request *models.ServerRequest) error {

	sqlQuery := request.Request

	if _, err := s.pool.Exec(context.Background(), sqlQuery); err != nil {
		fmt.Printf("Ошибка драйвера\n%v\n", err)
		return err
	}

	return nil
}
