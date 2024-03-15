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
	r.Put("/mapel/<id>", res.UpdateMapel)
	r.Delete("/mapel/<id>", res.delete)
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

func (r resource) UpdateMapel(c *routing.Context) error {
	ID := c.Param("id")

	var updateMapelReq request.UpdateMapelRequest
	if err := c.Read(&updateMapelReq); err != nil {
		return errors.BadRequest("invalid request")
	}

	updatedMapel, err := r.service.Update(c.Request.Context(), ID, updateMapelReq)
	if err != nil {
		return err
	}

	return c.WriteWithStatus(updatedMapel, http.StatusOK)
}

func (r resource) delete(c *routing.Context) error {
	mapel, err := r.service.Delete(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(mapel)
}
