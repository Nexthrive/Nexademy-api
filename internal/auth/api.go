package auth

import (
	"net/http"
	"nexademy/internal/auth/request"
	"nexademy/internal/errors"
	"nexademy/pkg/log"

	routing "github.com/go-ozzo/ozzo-routing/v2"
)

// RegisterHandlers registers handlers for different HTTP requests.
func RegisterHandlers(rg *routing.RouteGroup, service Service, logger log.Logger) {
	res := resource{service, logger}
	rg.Post("/login", res.login)
	rg.Post("/createUser", res.Create)
	rg.Post("/user/<id>", res.GetUserByID)
	rg.Put("/user/<id>", res.UpdateUser)
	rg.Get("/users", res.Query)
}

type resource struct {
	service Service
	logger  log.Logger
}

func NewResource(service Service, logger log.Logger) *resource {
	return &resource{
		service: service,
		logger:  logger,
	}
}

func (r resource) login(c *routing.Context) error {
	var req request.Login

	if err := c.Read(&req); err != nil {
		r.logger.Errorf("invalid request: %v", err)
		return errors.BadRequest("invalid request")
	}

	token, err := r.service.Login(c.Request.Context(), req)
	if err != nil {
		return err
	}

	return c.WriteWithStatus(token, http.StatusOK)
}

func (r resource) Create(c *routing.Context) error {
	var input request.CreateUser

	if err := c.Read(&input); err != nil {
		return errors.BadRequest("")
	}

	response, err := r.service.signup(c.Request.Context(), input)

	if err != nil {
		return err
	}

	return c.WriteWithStatus(response, http.StatusCreated)
}

func (r resource) GetUserByID(c *routing.Context) error {
	// Mendapatkan ID pengguna dari parameter URL
	userID := c.Param("id")

	// Memanggil service untuk mendapatkan pengguna berdasarkan ID
	user, err := r.service.GetUser(c.Request.Context(), userID)
	if err != nil {
		// Mengembalikan kesalahan jika terjadi kesalahan saat mengambil pengguna
		return err
	}

	// Mengembalikan pengguna sebagai respons
	return c.WriteWithStatus(user, http.StatusOK)
}

func (r resource) UpdateUser(c *routing.Context) error {
	// Mendapatkan ID pengguna dari parameter URL
	userID := c.Param("id")

	// Membaca data pembaruan dari body permintaan
	var updateUserReq request.UpdateUser
	if err := c.Read(&updateUserReq); err != nil {
		return errors.BadRequest("invalid request")
	}

	// Memanggil service untuk melakukan pembaruan pengguna
	updatedUser, err := r.service.Update(c.Request.Context(), userID, updateUserReq)
	if err != nil {
		return err
	}

	// Mengembalikan respons sukses jika tidak ada kesalahan
	return c.WriteWithStatus(updatedUser, http.StatusOK)
}

func (r resource) Query(c *routing.Context) error {
	users, err := r.service.Query(c.Request.Context())
	if err != nil {
		return err
	}

	return c.WriteWithStatus(users, http.StatusOK)
}
