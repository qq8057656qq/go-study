package _select

import (
	"net/http"
	"time"
)

func Racer(a, b string) string {
	aDuration := measureResponseTime(a)
	bDurtaion := measureResponseTime(b)
	if aDuration < bDurtaion {
		return a
	}
	return b
}

func measureResponseTime(a string) time.Duration {
	startA := time.Now()
	http.Get(a)
	aDuration := time.Since(startA)
	return aDuration
}
