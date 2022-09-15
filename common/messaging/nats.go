package messaging

import (
	"github.com/nats-io/nats.go"
	"gitlab.artin.ai/backend/courier-management/common/config"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/common/logger/tag"
	"sync"
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
