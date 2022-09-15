package messaging

import (
	"sync"

	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/common/logger/tag"
	"github.com/nats-io/nats.go"
)

var natsClientSetupOnce sync.Once
var natsClient *nats.Conn

// TODO: Wrap NATS over an interface for abstraction
func NatsClient() *nats.Conn {
	natsClientSetupOnce.Do(func() {
		natsAddress := config.Nats().Address
		var err error

		natsClient, err = nats.Connect(natsAddress)
		if err != nil {
			logger.Error("failed to get NATS client", err, tag.Str("natsAddress", natsAddress))
		}
	})

	return natsClient
}
