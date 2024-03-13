package kelas

import (
	"nexademy/pkg/log"

	routing "github.com/go-ozzo/ozzo-routing/v2"
)

func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}

	r.Get("/kelas", res.query)

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