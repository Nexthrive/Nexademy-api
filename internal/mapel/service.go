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

func (s service) Get(ctx context.Context, id_mapel string) (response.MapelResponse, error) {
	mapel, err := s.repo.Get(ctx, id_mapel)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response.MapelResponse{}, err
		}

		return response.MapelResponse{}, err
	}

	return response.MapelResponse{
		Status:  "200",
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
		Status:  "200",
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
		Id_mapel:  idstr,
		Judul:     input.Judul,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := s.repo.Create(ctx, newMapel)
	if err != nil {
		return response.MapelResponse{}, err
	}

	return response.MapelResponse{
		Status:  "201",
		Message: "Create successful",
		Data:    newMapel,
	}, nil
}

