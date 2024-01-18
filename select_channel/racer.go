package select_channel

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondsTimeout = 10 * time.Second

func Racer(urlA, urlB string) (winner string, error error) {
	return ConfigurableRacer(urlA, urlB, tenSecondsTimeout)
}

func ConfigurableRacer(urlA, urlB string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(urlA):
		return urlA, nil
	case <-ping(urlB):
		return urlB, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timeout while waiting fot %s and %s", urlA, urlB)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
