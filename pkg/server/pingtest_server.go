package server

import (
	"context"

	"github.com/yeonfeel/gangaji/pkg/pb/v1beta1/common"
)

func (s *server) PingPong(_ context.Context, ping *common.Ping) (*common.Pong, error) {
	return &common.Pong{
		Pong: ping.Ping,
	}, nil
}
