package db

import (
	"sync"

	t38c "github.com/axvq/tile38-client"
	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/common/logger/tag"
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
