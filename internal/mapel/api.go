package mapel

import (
	"net/http"
	"nexademy/internal/errors"
	"nexademy/internal/mapel/request"
	"nexademy/pkg/log"

	routing "github.com/go-ozzo/ozzo-routing/v2"
)

func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}

	r.Get("/mapel", res.query)
	r.Get("/mapel/<id>", res.get)
	r.Post("/mapel", res.create)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) get(c *routing.Context) error {
	id := c.Param("id")

	mapel, err := r.service.Get(c.Request.Context(), id)

	if err != nil {
		return err
	}

	return c.WriteWithStatus(mapel, http.StatusOK)
}

func (r resource) query(c *routing.Context) error {
	mapeles, err := r.service.Query(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(mapeles)
}

func (r resource) create(c *routing.Context) error {
	var input request.CreateMapelRequest

	if err := c.Read(&input); err != nil {
		return errors.BadRequest("")
	}

	response, err := r.service.Create(c.Request.Context(), input)

	if err != nil {
		return err
	}

	return c.WriteWithStatus(response, http.StatusCreated)
}




