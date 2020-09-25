package redact

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/render"
	"github.com/gufranmirza/redact-api-golang/src/models"
	"github.com/gufranmirza/redact-api-golang/src/web/interfaces/errorinterface"
	"github.com/gufranmirza/redact-api-golang/src/web/interfaces/redactinterface"
)

type redact struct {
	logger *log.Logger
}

// NewRedact returns new object of RedactService
func NewRedact() Redact {
	return &redact{
		logger: log.New(os.Stdout, "redact :=> ", log.LstdFlags),
	}
}

func (rs *redact) RedactJSON() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		txID, _ := r.Context().Value(models.HdrRequestID).(string)
		data := &redactRequest{}
		if err := render.Bind(r, data); err != nil {
			render.Render(w, r, &errorinterface.ErrorResponse{
				Err:            err,
				HTTPStatusCode: http.StatusBadRequest,
				Code:           http.StatusBadRequest,
				Status:         http.StatusText(http.StatusBadRequest),
				Error:          err.Error(),
			})
			return
		}

		// Apply full path replace
		for _, item := range data.RedactCompletely {
			rs.logger.Printf("%s applying full path replace at path: %s", txID, item)
			keys := strings.Split(item, ".")
			err := rs.redactJSON(data.JSONToRedact, keys, []string{}, true)
			if err != nil {
				rs.logger.Printf("%s failed to apply full path replace at path: %s with error %s", txID, item, err)
			}
		}

		// Apply regex
		for _, item := range data.RedactRegexes {
			rs.logger.Printf("%s applying regex replace at path: %s", txID, item.Path)
			keys := strings.Split(item.Path, ".")
			err := rs.redactJSON(data.JSONToRedact, keys, item.Regexes, false)
			if err != nil {
				rs.logger.Printf("%s failed to apply regex at path: %s with error %s", txID, item, err)
			}
		}

		buff, err := json.Marshal(data.JSONToRedact)
		if err != nil {
			render.Render(w, r, &errorinterface.ErrorResponse{
				Err:            err,
				HTTPStatusCode: http.StatusInternalServerError,
				Code:           http.StatusInternalServerError,
				Status:         http.StatusText(http.StatusInternalServerError),
				Error:          err.Error(),
			})
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, string(buff))
		return
	})
}

// ===================== Bindings ========================= //

type redactRequest struct {
	*redactinterface.RedactInterface
}

func (d *redactRequest) Bind(r *http.Request) error {
	return nil
}
