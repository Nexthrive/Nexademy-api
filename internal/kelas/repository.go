package kelas

import (
	"context"
	"nexademy/internal/entity"
	"nexademy/pkg/dbcontext"
	"nexademy/pkg/log"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

type Repository interface {
	Get(ctx context.Context, id_kelas string) (entity.Kelas, error)
	Query(ctx context.Context) ([]entity.Kelas, error)
	Create(ctx context.Context, kelas entity.Kelas) error
}

type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

func NewRepo(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Create implements Repository.
func (r repository) Create(ctx context.Context, kelas entity.Kelas) error {
	return r.db.With(ctx).Model(&kelas).Insert()
}

// Get implements Repository.
func (r repository) Get(ctx context.Context, id_kelas string) (entity.Kelas, error) {
	var kelas entity.Kelas
	err := r.db.With(ctx).Select().From("kelas").Where(dbx.HashExp{"id_kelas": id_kelas}).One(&kelas)
	if err != nil {
			return kelas, err
	}
	return kelas, nil
}

// Query implements Repository.
func (r repository) Query(ctx context.Context) ([]entity.Kelas, error) {
	var kelases []entity.Kelas

	err := r.db.With(ctx).Select().All(&kelases)

	return kelases, err
}
