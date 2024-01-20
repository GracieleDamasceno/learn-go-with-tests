package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertEquals(counter.Value(), 3, t)
	})

	t.Run("running safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()
		assertEquals(counter.Value(), wantedCount, t)
	})
}

func assertEquals(got int, want int, t *testing.T) {
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
