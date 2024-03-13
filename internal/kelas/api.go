package kelas

import (
	"net/http"
	"nexademy/internal/errors"
	"nexademy/internal/kelas/request"
	"nexademy/pkg/log"

	routing "github.com/go-ozzo/ozzo-routing/v2"
)

func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}

	r.Get("/kelas", res.query)
	r.Get("/kelas/<id>", res.get)
	r.Post("/kelas", res.create)

}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) query(c *routing.Context) error {
	kelases, err := r.service.Query(c.Request.Context())

	if err != nil {
		return err
	}

	return c.Write(kelases)
}

func (r resource) get(c *routing.Context) error {
	// Get the ID from the request path parameters
	id := c.Param("id")

	// Call the service to get the kelas data
	response, err := r.service.Get(c.Request.Context(), id)

	if err != nil {
		return err
	}

	// If the call to the service is successful, return the response with HTTP status 200
	return c.Write(response)
}

func (r resource) create(c *routing.Context) error {
	var input request.CreateKelasRequest

	if err := c.Read(&input); err != nil {
		return errors.BadRequest("")
	}

	response, err := r.service.Create(c.Request.Context(), input)

	if err != nil {
		return err
	}

	return c.WriteWithStatus(response, http.StatusCreated)
}
