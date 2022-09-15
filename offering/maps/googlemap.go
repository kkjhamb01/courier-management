package maps

import (
	"sync"

	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/common/logger/tag"
	"googlemaps.github.io/maps"
)

var googleClient *maps.Client
var googleClientSetupOnce sync.Once

func SetupGoogleClient() {
	googleClientSetupOnce.Do(func() {
		var err error
		googleClient, err = maps.NewClient(maps.WithAPIKey(config.GoogleMap().ApiKey))
		if err != nil {
			logger.Fatal("failed to create the google maps client", tag.Err("err", err))
		}
	})
}

func GoogleClient() *maps.Client {
	return googleClient
}
