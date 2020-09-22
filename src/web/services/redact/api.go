package redact

import "net/http"

// Redact interface
type Redact interface {
	RedactJSON()  http.Handler
}
