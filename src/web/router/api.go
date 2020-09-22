package router

import (
	"net/http"
)

// Router interface
type Router interface {
	Router() *http.ServeMux
}
