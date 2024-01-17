package select_channel

import (
	"net/http"
	"time"
)

func Racer(urlA, urlB string) (winner string) {
	durationA := measureResponseTimeURL(urlA)

	durationB := measureResponseTimeURL(urlB)

	if durationA < durationB {
		return urlA
	}

	return urlB
}

func measureResponseTimeURL(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
