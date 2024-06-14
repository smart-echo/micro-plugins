// Package mucp provides an mucp server
package mucp

import (
	"github.com/smart-echo/micro/server"
	"github.com/smart-echo/micro/util/cmd"
)

func init() {
	cmd.DefaultServers["mucp"] = NewServer
}

// NewServer returns a micro server interface.
func NewServer(opts ...server.Option) server.Server {
	return server.NewServer(opts...)
}