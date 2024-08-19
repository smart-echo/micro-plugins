// Package mucp provides an mucp client
package mucp

import (
	"github.com/smart-echo/micro/client"
	"github.com/smart-echo/micro/cmd"
)

func init() {
	cmd.DefaultClients["mucp"] = NewClient
}

// NewClient returns a new micro client interface.
func NewClient(opts ...client.Option) client.Client {
	return client.NewClient(opts...)
}
