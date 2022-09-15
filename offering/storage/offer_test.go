package storage

import (
	"context"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestStorageImplIncreaseOfferRetriesTransaction(t *testing.T) {
	numberOfGoRoutines := 100000
	offerId := "10000000-0000-0000-0000-000000000022"
	ctx := context.Background()
	s := storageImpl{}
	wg := sync.WaitGroup{}
	wg.Add(numberOfGoRoutines)

	err := s.ResetOfferRetries(ctx, offerId)
	if err != nil {
		t.Error("failed to reset offer retries", err)
	}

	for i := 0; i < numberOfGoRoutines; i++ {
		_, err := s.IncreaseOfferRetries(ctx, offerId)
		if err != nil {
			t.Error("failed to increase offer retries", err)
		}
		wg.Done()
	}
	wg.Wait()

	retries, err := s.GetOfferRetries(ctx, offerId)
	assert.Equal(t, numberOfGoRoutines, retries)
	assert.Nil(t, err)
}

func BenchmarkStorageImplIncreaseOfferRetriesWithLock(b *testing.B) {
	lockEnabled = true
	runStorageImplIncreaseOfferRatesBench(b)
}

func BenchmarkStorageImplIncreaseOfferRetriesWithoutLock(b *testing.B) {
	lockEnabled = false
	runStorageImplIncreaseOfferRatesBench(b)
}

func runStorageImplIncreaseOfferRatesBench(b *testing.B) {
	numberOfGoRoutines := 50
	offerId := "10000000-0000-0000-0000-000000000022"
	ctx := context.Background()
	s := storageImpl{}
	wg := sync.WaitGroup{}

	for i := 0; i < b.N; i++ {
		wg.Add(numberOfGoRoutines)
		err := s.ResetOfferRetries(ctx, offerId)
		if err != nil {
			b.Error("failed to reset offer retries", err)
		}

		for n := 0; n < numberOfGoRoutines; n++ {
			go func() {
				_, err := s.IncreaseOfferRetries(ctx, offerId)
				if err != nil {
					b.Error("failed to increase offer retries", err)
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}
