package main

import (
	"net/http"
	"strings"
)

type RoundTripCounter struct {
	count int
}

func main() {
	var tripper http.RoundTripper = &RoundTripCounter{}
	r, _ := http.NewRequest(http.MethodGet, "http://devto", strings.NewReader("test"))
	_, _ = tripper.RoundTrip(r)
}
func (rt *RoundTripCounter) RoundTrip(*http.Request) (*http.Response, error) {
	rt.count += 1
	return nil, nil
}
