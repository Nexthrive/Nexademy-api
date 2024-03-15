package auth

import (
	"context"
	"nexademy/internal/entity"
	"nexademy/pkg/dbcontext"
	"nexademy/pkg/log"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	GetByUsernameAndPassword(ctx context.Context, nis int, Passphrase string) (entity.User, error)
	Create(ctx context.Context, user entity.User) error
	Get(ctx context.Context, id string) (entity.User, error)
	Query(ctx context.Context) ([]entity.User, error)
	Update(ctx context.Context, req entity.User) error
}

type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

func NewRepo(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

func (r repository) GetByUsernameAndPassword(ctx context.Context, nis int, Passphrase string) (entity.User, error) {
	var user entity.User
	err := r.db.With(ctx).Select().From("user").Where(dbx.HashExp{"nis": nis}).One(&user)
	if err != nil {
		return user, err
	}

	// Compare stored hash with the hash of the provided passphrase
	if err := bcrypt.CompareHashAndPassword([]byte(user.Passphrase), []byte(Passphrase)); err != nil {
		return user, err
	}

	return user, nil
}

func (r repository) Create(ctx context.Context, req entity.User) error {
	return r.db.With(ctx).Model(&req).Insert()
}

func (r repository) Get(ctx context.Context, id string) (entity.User, error) {
	var user entity.User
	err := r.db.With(ctx).Select().Model(id, &user)
	return user, err
}

func (r repository) Update(ctx context.Context, req entity.User) error {
	return r.db.With(ctx).Model(&req).Update()
}

func (r repository) Query(ctx context.Context) ([]entity.User, error) {
	var user []entity.User

	err := r.db.With(ctx).Select().All(&user)

	return user, err
}
