package healthcheck

import (
	"context"
	"gitlab.artin.ai/backend/courier-management/offering/db"
	"net/http"
)

func Readyz(w http.ResponseWriter, _ *http.Request) {
	redisClient := db.OfferRedisClient()
	// TODO replace the following test code with "ping" whenever pining is available in OKD
	err := redisClient.Set(context.Background(), "test", "ok", -1).Err()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	tile38Client := db.Tile38Client()
	// TODO replace the following test code with "ping" whenever pining is available in OKD
	err = tile38Client.Keys.Set("test", "ok").String("ok").Do()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
