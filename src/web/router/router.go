// Package api configures an http server for administration and application resources.

package router

import (
	"net/http"

	"github.com/gufranmirza/redact-api-golang/src/config"
	"github.com/gufranmirza/redact-api-golang/src/models"
	"github.com/gufranmirza/redact-api-golang/src/web/services/health"
	"github.com/gufranmirza/redact-api-golang/src/web/services/redact"
)

type router struct {
	config *models.AppConfig
	health health.Health
	redact redact.Redact
}

// NewRouter returns the router implementation
func NewRouter() Router {
	return &router{
		config: config.Config,
		health: health.NewHealth(),
		redact: redact.NewRedact(),
	}
}

// Router configures application resources and routes.
func (router *router) Router() *http.ServeMux {
	r := http.NewServeMux()

	// URL router prefix
	urlPrefix := router.config.URLPrefix

	// =================  health routes ======================
	r.Handle(urlPrefix+"/health/", router.health.GetHealth())

	// =================  redact routes ======================
	r.Handle(urlPrefix+"/redact/", router.redact.RedactJSON())

	return r
}
