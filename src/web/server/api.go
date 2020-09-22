package server

import "time"

// Server implements server interface
type Server interface {
	Start()
	StartTimestampUTC() time.Time
}
