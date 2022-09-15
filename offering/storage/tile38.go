package storage

import (
	"context"
)

type tile38Impl struct {
}

func (t tile38Impl) rollback() error {
	// TODO tx is not supported on tile38 for now, rollback procedure needs to be handled manually
	return nil
}

func (t tile38Impl) commit(ctx context.Context) error {
	// TODO tx is not supported on tile38 for now, rollback procedure needs to be handled manually
	return nil
}

func createTile38Tx() tile38Impl {
	return tile38Impl{}
}
