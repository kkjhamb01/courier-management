package maps

import (
	"gitlab.artin.ai/backend/courier-management/common/config"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/common/logger/tag"
	"googlemaps.github.io/maps"
	"sync"
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
