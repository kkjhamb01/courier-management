package db

import (
	t38c "github.com/axvq/tile38-client"
	"gitlab.artin.ai/backend/courier-management/common/config"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/common/logger/tag"
	"sync"
)

// to make sure tile38 would be set up only once
var tileSetupOnce sync.Once
var t38client *t38c.Client

func SetupTile38Client() {
	tileSetupOnce.Do(func() {
		var err error
		// TODO get pool size from config
		t38client, err = t38c.New(config.Tile38().Address, t38c.SetPoolSize(10))
		if err != nil {
			logger.Fatal("failed to create tile38 client", tag.Err("err", err))
		}

		logger.Info("created connection to tile (Not guaranteed as pinging is disabled)")
	})
}

func Tile38Client() *t38c.Client {
	return t38client
}
