package healthinterface

import "time"

// ServiceStatus represents the service status
type ServiceStatus string

// ConnectionStatus represents the connection status
type ConnectionStatus string

var (
	// ServiceRunning represents service is Running
	ServiceRunning ServiceStatus = "Running"
	// ServiceDegraded represents service health is Degraded
	ServiceDegraded ServiceStatus = "Degraded"
	// ServiceStopped represents service is Stopped
	ServiceStopped ServiceStatus = "Stopped"
	// ConnectionActive represents connection is active
	ConnectionActive ConnectionStatus = "Active"
	// ConnectionDisconnected represents connection is disconnected
	ConnectionDisconnected ConnectionStatus = "Disconnected"
)

// Health represents health response
type Health struct {
	TimeStampUTC        time.Time           `json:"TimeStampUTC,omitempty"`
	ServiceName         string              `json:"ServiceName,omitempty"`
	ServiceProvider     string              `json:"ServiceProvider,omitempty"`
	ServiceVersion      string              `json:"ServiceVersion,omitempty"`
	ServiceStatus       ServiceStatus       `json:"ServiceStatus,omitempty"`
	ServiceStartTimeUTC time.Time           `json:"ServiceStartTimeUTC,omitempty"`
	Uptime              float64             `json:"Uptime,omitempty"`
	InboundInterfaces   []InboundInterface  `json:"InboundInterfaces,omitempty"`
	OutboundInterfaces  []OutboundInterface `json:"OutboundInterfaces,omitempty"`
}

// InboundInterface is inbound network inferfaces
type InboundInterface struct {
	ApplicationName  string           `json:"ApplicationName,omitempty"`
	ConnectionStatus ConnectionStatus `json:"ConnectionStatus,omitempty"`
	TimeStampUTC     time.Time        `json:"TimeStampUTC,omitempty"`
	Hostname         string           `json:"Hostname,omitempty"`
	Address          string           `json:"Address,omitempty"`
	OS               string           `json:"OS,omitempty"`
}

// OutboundInterface is outbound network interfaces
type OutboundInterface struct {
	ApplicationName  string           `json:"ApplicationName,omitempty"`
	TimeStampUTC     time.Time        `json:"TimeStampUTC,omitempty"`
	URLs             []string         `json:"URLs,omitempty"`
	ConnectionStatus ConnectionStatus `json:"ConnectionStatus,omitempty"`
}
