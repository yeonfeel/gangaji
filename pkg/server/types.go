package server

import (
	"github.com/yeonfeel/gopkg/logger"

	"github.com/yeonfeel/gangaji/config"
	"github.com/yeonfeel/gangaji/pkg/pb/v1beta1/common"
)

var log = logger.Get("server")

// Server is an interface for gRPC Server
type Server interface {
	common.PingTestServer
}

type server struct {
	cfg *config.Config
}

// NewServer returns the common interface
func NewServer(cfg *config.Config) Server {
	s := &server{
		cfg: cfg,
	}

	return s
}
