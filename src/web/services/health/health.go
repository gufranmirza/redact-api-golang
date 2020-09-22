package health

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/gufranmirza/redact-api-golang/src/config"
	"github.com/gufranmirza/redact-api-golang/src/models"
	"github.com/gufranmirza/redact-api-golang/src/web/interfaces/healthinterface"
)

type health struct {
	logger *log.Logger
	config *models.AppConfig
}

// NewHealth returns new health object
func NewHealth() Health {
	return &health{
		logger: log.New(os.Stdout, "health :=> ", log.LstdFlags),
		config: config.Config,
	}
}

// GetHealth returns heath of service
func (h *health) GetHealth() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		txID, _ := r.Context().Value(models.HdrRequestID).(string)

		healthStatus := healthinterface.Health{}
		healthStatus.ServiceName = h.config.ServiceName
		healthStatus.ServiceProvider = h.config.ServiceProvider
		healthStatus.ServiceVersion = h.config.ServiceVersion
		healthStatus.TimeStampUTC = time.Now().UTC()
		healthStatus.ServiceStatus = healthinterface.ServiceRunning

		inbound := []healthinterface.InboundInterface{}
		outbound := []healthinterface.OutboundInterface{}

		// add internal server details
		name, _ := os.Hostname()

		server := healthinterface.InboundInterface{}
		server.Hostname = name
		server.OS = runtime.GOOS
		server.TimeStampUTC = time.Now().UTC()
		server.ApplicationName = h.config.ServiceName
		server.ConnectionStatus = healthinterface.ConnectionActive

		exIP, err := externalIP()
		if err != nil {
			h.logger.Printf("%s:%s Failed to obtain inbound ip address with error %v\n", models.HdrRequestID, txID, err)
			server.ConnectionStatus = healthinterface.ConnectionDisconnected
		}
		server.Address = exIP
		inbound = append(inbound, server)

		healthStatus.InboundInterfaces = inbound
		healthStatus.OutboundInterfaces = outbound
		buff, err := json.Marshal(healthStatus)
		if err != nil {
			h.logger.Printf("%s:%s Failed unmarshal health object error %v\n", models.HdrRequestID, txID, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, string(buff))
		return
	})
}
