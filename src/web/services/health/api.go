package health

import (
	"net/http"
)

const (
	// FailedToObtainInboundIP is error code used if it failed to obtain inbound connection status
	FailedToObtainInboundIP = "Failed-To-Obtain-Intbound-IP"
)

// Health interface
type Health interface {
	GetHealth() http.Handler
}
