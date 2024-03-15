package mapel

import (
	"context"
	"database/sql"
	"errors"
	"math/rand"
	"nexademy/internal/entity"
	"nexademy/internal/mapel/request"
	"nexademy/internal/mapel/response"
	"nexademy/pkg/log"
	"strconv"
	"time"
)

type Service interface {
	Get(ctx context.Context, id_mapel string) (response.MapelResponse, error)
	Query(ctx context.Context) (response.MapelResponse, error)
	Create(ctx context.Context, input request.CreateMapelRequest) (response.MapelResponse, error)
	Update(ctx context.Context, id string, req request.UpdateMapelRequest) (response.MapelResponse, error)
	Delete(ctx context.Context, id string) (response.ResponseDel, error)
}

type service struct {
	repo   Repository
	logger log.Logger
}

type Mapel struct {
	entity.Mapel
}

func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

func (s service) Get(ctx context.Context, id string) (response.MapelResponse, error) {
	mapel, err := s.repo.Get(ctx, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response.MapelResponse{}, err
		}

		return response.MapelResponse{}, err
	}

	return response.MapelResponse{
		Status:  200,
		Message: "Query successful",
		Data:    mapel,
	}, nil
}

func (s service) Query(ctx context.Context) (response.MapelResponse, error) {
	mapeles, err := s.repo.Query(ctx)

	if err != nil {
		return response.MapelResponse{}, err
	}

	return response.MapelResponse{
		Status:  200,
		Message: "Query successful",
		Data:    mapeles,
	}, nil
}

func (s service) Create(ctx context.Context, input request.CreateMapelRequest) (response.MapelResponse, error) {
	now := time.Now()
	rand.Seed(now.UnixNano())

	id := rand.Intn(900000) + 100000
	idstr := strconv.Itoa(id)
	newMapel := entity.Mapel{
		ID:        idstr,
		Judul:     input.Judul,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := s.repo.Create(ctx, newMapel)
	if err != nil {
		return response.MapelResponse{}, err
	}

	return response.MapelResponse{
		Status:  201,
		Message: "Create successful",
		Data:    newMapel,
	}, nil
}

func (s service) Update(ctx context.Context, id string, req request.UpdateMapelRequest) (response.MapelResponse, error) {
	mapel, err := s.repo.Get(ctx, id)
	if err != nil {
		return response.MapelResponse{}, err
	}
	now := time.Now()
	updateMapel := entity.Mapel{
		ID:        mapel.ID,
		Judul:     req.Judul,
		UpdatedAt: now,
	}

	err = s.repo.Update(ctx, updateMapel)
	if err != nil {
		return response.MapelResponse{}, err
	}

	res := response.MapelResponse{
		Status:  200,
		Message: "User updated successfully",
		Data:    updateMapel,
	}
	return res, nil

}

func (s service) Delete(ctx context.Context, id string) (response.ResponseDel, error) {
	if err := s.repo.Delete(ctx, id); err != nil {
		return response.ResponseDel{}, err
	}

	res := response.ResponseDel{
		Status:  200,
		Message: "Data deleted successfully",
	}
	return res, nil
}
