// Package api configures an http server for administration and application resources.

package router

import (
	"net/http"

	"github.com/gufranmirza/redact-api-golang/src/config"
	"github.com/gufranmirza/redact-api-golang/src/models"
	"github.com/gufranmirza/redact-api-golang/src/web/services/health"
)

type router struct {
	config *models.AppConfig
	health health.Health
}

// NewRouter returns the router implementation
func NewRouter() Router {
	return &router{
		config: config.Config,
		health: health.NewHealth(),
	}
}

// Router configures application resources and routes.
func (router *router) Router() *http.ServeMux {
	r := http.NewServeMux()

	// v1 URL router prefix
	// v1Prefix := router.config.URLPrefix + router.config.APIVersionV1

	// // =================  health routes ======================
	r.Handle(router.config.URLPrefix+"/health", router.health.GetHealth())

	// // =================  recruiters routes ======================
	// recruiterPrefix := v1Prefix + "/recruiters"
	// r.Post(recruiterPrefix+"/signup", router.recruiter.CreateRecruiter)

	return r
}
