package redact

import (
	"log"
	"os"
	"fmt"
	"net/http"
	"github.com/gufranmirza/redact-api-golang/src/models"
)

type redact struct {
	logger *log.Logger
}

// NewRedact returns new object of RedactService
func NewRedact () Redact {
	return &redact{
		logger: log.New(os.Stdout, "redact :=> ", log.LstdFlags),
	}
}

func (rs *redact)RedactJSON()  http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		txID, _ := r.Context().Value(models.HdrRequestID).(string)

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "to be implemented txID: ", txID)
		return
	})
}
