package auth

import (
	"context"
	"math/rand"
	"nexademy/internal/auth/request"
	"nexademy/internal/auth/response"
	"nexademy/internal/entity"
	"nexademy/internal/errors"
	"nexademy/pkg/log"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// Service encapsulates the authentication logic.
type Service interface {
	// authenticate authenticates a user using username and password.
	// It returns a JWT token if authentication succeeds. Otherwise, an error is returned.
	authenticate(ctx context.Context, loginReq request.Login) (entity.User, error)
	Login(ctx context.Context, loginreq request.Login) (response.EmptyResponse, error)
	signup(ctx context.Context, Req request.CreateUser) (response.EmptyResponse, error)
	GetUser(ctx context.Context, id string) (response.EmptyResponse, error)
	Update(ctx context.Context, id string, req request.UpdateUser) (response.EmptyResponse, error)
	Query(ctx context.Context) (response.EmptyResponse, error)
}

// Identity represents an authenticated user identity.
type Identity interface {
	// GetID returns the user ID.
	GetID() string
	// GetName returns the user name.
	GetName() string
}

type service struct {
	signingKey      string
	tokenExpiration int
	logger          log.Logger
	repo            Repository
}

// NewService creates a new authentication service.
func NewService(signingKey string, tokenExpiration int, logger log.Logger, repo Repository) Service {
	return &service{signingKey, tokenExpiration, logger, repo}
}

// Login authenticates a user and generates a JWT token if authentication succeeds.
// Otherwise, an error is returned.
func (s service) Login(ctx context.Context, loginreq request.Login) (response.EmptyResponse, error) {
	user, err := s.authenticate(ctx, loginreq)
	if err != nil {
		s.logger.Error("Authentication failed", "error", err)
		return response.EmptyResponse{}, errors.Unauthorized("authentication failed")
	}
	token, err := s.generateJWT(user)
	if err != nil {
		s.logger.Error("JWT generation failed", "error", err)
		return response.EmptyResponse{}, errors.InternalServerError("internal server error")
	}
	res := response.EmptyResponse{
		Status:  200,
		Message: "Login successful",
		Data:    token,
	}
	return res, nil
}

// authenticate authenticates a user using username and password.
// If username and password are correct, an identity is returned. Otherwise, nil is returned.
func (s *service) authenticate(ctx context.Context, loginReq request.Login) (entity.User, error) {
	user, err := s.repo.GetByUsernameAndPassword(ctx, loginReq.Nis, loginReq.Passphrase)
	if err != nil {
		// If authentication fails, return unauthorized error
		s.logger.Error("Authentication failed", "error", err)
		return entity.User{}, errors.Unauthorized("authentication failed")
	}

	return user, nil
}

func (s service) signup(ctx context.Context, req request.CreateUser) (response.EmptyResponse, error) {
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(90000) + 10000
	idstr := strconv.Itoa(id)

	// Hashing passphrase
	hashedPassphrase, err := bcrypt.GenerateFromPassword([]byte(req.Passphrase), bcrypt.DefaultCost)
	if err != nil {
		return response.EmptyResponse{}, err
	}

	err = s.repo.Create(ctx, entity.User{
		ID:         idstr,
		Nis:        req.Nis,
		Name:       req.Name,
		Passphrase: string(hashedPassphrase), // Simpan passphrase yang telah di-hash
		Email:      req.Email,
		No_telp:    req.No_telp,
		Gender:     req.Gender,
		Religion:   req.Religion,
	})
	if err != nil {
		return response.EmptyResponse{}, err
	}
	user, err := s.GetUser(ctx, idstr)
	if err != nil {
		return response.EmptyResponse{}, err
	}
	res := response.EmptyResponse{
		Status:  201,
		Message: "User created successfully",
		Data:    user,
	}
	return res, nil
}
func (s service) GetUser(ctx context.Context, id string) (response.EmptyResponse, error) {
	user, err := s.repo.Get(ctx, id)
	if err != nil {
		return response.EmptyResponse{}, err
	}

	res := response.EmptyResponse{
		Status:  200,
		Message: "User retrieved successfully",
		Data:    user,
	}
	return res, nil
}

func (s service) Query(ctx context.Context) (response.EmptyResponse, error) {
	user, err := s.repo.Query(ctx)

	if err != nil {
		return response.EmptyResponse{}, err
	}

	return response.EmptyResponse{
		Status:  200,
		Message: "Query successful",
		Data:    user,
	}, nil

}

func (s service) Update(ctx context.Context, id string, req request.UpdateUser) (response.EmptyResponse, error) {
	user, err := s.repo.Get(ctx, id)
	if err != nil {
		return response.EmptyResponse{}, err
	}
	updateUser := entity.User{
		ID:         user.ID,
		Name:       req.Name,
		Email:      user.Email,
		No_telp:    req.No_telp,
		Gender:     user.Gender,
		Nis:        user.Nis,
		Passphrase: user.Passphrase,
		Religion:   req.Religion,
	}

	if req.Passphrase != "" {
		hashedPassphrase, err := bcrypt.GenerateFromPassword([]byte(req.Passphrase), bcrypt.DefaultCost)
		if err != nil {
			return response.EmptyResponse{}, err
		}
		updateUser.Passphrase = string(hashedPassphrase)
	}

	err = s.repo.Update(ctx, updateUser)
	if err != nil {
		return response.EmptyResponse{}, err
	}

	res := response.EmptyResponse{
		Status:  200,
		Message: "User updated successfully",
		Data:    updateUser,
	}
	return res, nil

}

// generateJWT generates a JWT that encodes an identity.
func (s service) generateJWT(identity Identity) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   identity.GetID(),
		"name": identity.GetName(),
		"exp":  time.Now().Add(time.Duration(s.tokenExpiration) * time.Hour).Unix(),
	}).SignedString([]byte(s.signingKey))
}
