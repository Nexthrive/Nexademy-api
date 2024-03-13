package kelas

import (
	"context"
	"nexademy/internal/kelas/request"
	"nexademy/internal/kelas/response"
	"nexademy/pkg/log"
)

type Service interface {
	Get(ctx context.Context, id string) (response.EmptyResponse, error)
	Query(ctx context.Context) (response.EmptyResponse, error)
	Create(ctx context.Context, input request.CreateKelasRequest) error
}

type service struct {
	repo   Repository
	logger log.Logger
}

func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Create implements Service.
func (s service) Create(ctx context.Context, input request.CreateKelasRequest) error {
	panic("unimplemented")
}

// Get implements Service.
func (s service) Get(ctx context.Context, id string) (response.EmptyResponse, error) {
	panic("unimplemented")
}

// Query implements Service.
func (s service) Query(ctx context.Context) (response.EmptyResponse, error) {
	kelases, err := s.repo.Query(ctx)

	if err != nil {
		return response.EmptyResponse{
				Status:  "500",
				Message: "An error occurred",
				Data:    nil,
		}, err
	}

	return response.EmptyResponse{
		Status:     "200",
		Message:    "Query successful",
		Data:       kelases,
}, nil

}
