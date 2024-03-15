package mapel

import (
	"context"
	"nexademy/internal/entity"
	"nexademy/pkg/dbcontext"
	"nexademy/pkg/log"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

type Repository interface {
	Get(ctx context.Context, id_mapel string) (entity.Mapel, error)
	Query(ctx context.Context) ([]entity.Mapel, error)
	Create(ctx context.Context, mapel entity.Mapel) error
	Update(ctx context.Context, req entity.Mapel) error
	Delete(ctx context.Context, id string) error
}

type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

func NewRepo(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

func (r repository) Get(ctx context.Context, id string) (entity.Mapel, error) {
	var mapel entity.Mapel
	err := r.db.With(ctx).Select().From("mapel").Where(dbx.HashExp{"id": id}).One(&mapel)
	if err != nil {
		return mapel, err
	}
	return mapel, nil
}

func (r repository) Query(ctx context.Context) ([]entity.Mapel, error) {
	var mapeles []entity.Mapel
	err := r.db.With(ctx).Select().All(&mapeles)
	return mapeles, err
}

func (r repository) Create(ctx context.Context, mapel entity.Mapel) error {
	return r.db.With(ctx).Model(&mapel).Insert()
}

func (r repository) Update(ctx context.Context, req entity.Mapel) error {
	return r.db.With(ctx).Model(&req).Update()
}

func (r repository) Delete(ctx context.Context, id string) error {
	mapel, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.With(ctx).Model(&mapel).Delete()
}