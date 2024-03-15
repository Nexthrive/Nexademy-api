package kelas

import (
	"context"
	"database/sql"
	"errors"
	"nexademy/internal/entity"
	"nexademy/internal/kelas/request"
	"nexademy/internal/kelas/response"
	"nexademy/pkg/log"
	"time"
)

type Service interface {
	Get(ctx context.Context, id string) (response.EmptyResponse, error)
	Query(ctx context.Context) (response.EmptyResponse, error)
	Create(ctx context.Context, input request.CreateKelasRequest) (response.EmptyResponse, error)
}

type service struct {
	repo   Repository
	logger log.Logger
}

func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Create implements Service.
func (s service) Create(ctx context.Context, input request.CreateKelasRequest) (response.EmptyResponse, error) {
	now := time.Now()

	err := s.repo.Create(ctx, entity.Kelas{
		ID_Kelas:  input.ID_Kelas,
		Walas:     input.Walas,
		CreatedAt: now,
		UpdatedAt: now,
	})

	if err != nil {
		return response.EmptyResponse{}, err
	}

	return response.EmptyResponse{
		Status:  "201",
		Message: "Create successful",
		Data:    nil,
	}, nil
}

// Get implements Service.
func (s service) Get(ctx context.Context, id_kelas string) (response.EmptyResponse, error) {
	kelas, err := s.repo.Get(ctx, id_kelas)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response.EmptyResponse{}, err
		}

		return response.EmptyResponse{}, err
	}

	return response.EmptyResponse{
		Status:  "200",
		Message: "Query successful",
		Data:    kelas,
	}, nil
}

// Query implements Service.
func (s service) Query(ctx context.Context) (response.EmptyResponse, error) {
	kelases, err := s.repo.Query(ctx)

	if err != nil {
		return response.EmptyResponse{}, err
	}

	return response.EmptyResponse{
		Status:  "200",
		Message: "Query successful",
		Data:    kelases,
	}, nil

}
