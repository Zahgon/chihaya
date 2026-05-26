// Package metrics implements a standalone HTTP server for serving pprof
// profiles and Prometheus metrics.
package metrics

import (
	"net/http"

	"github.com/chihaya/chihaya/pkg/stop"
)

// Server represents a standalone HTTP server for serving a Prometheus metrics
// endpoint.
type Server struct {
	srv *http.Server
}

// Stop shuts down the server.
func (s *Server) Stop() stop.Result { _ = "STUB: not implemented"; return *new(stop.Result) }

// NewServer creates a new instance of a Prometheus server that asynchronously
// serves requests.
func NewServer(addr string) *Server { _ = "STUB: not implemented"; return nil }
