package main

import (
	"regexp"
	"testing"
)

func TestGetExpirationDate(t *testing.T) {
	const url = "www.google.com.ar"
	const port = "443"

	want := regexp.MustCompile(`^((?:(\d{4}-\d{2}-\d{2})T(\d{2}:\d{2}:\d{2}(?:\.\d+)?))(Z|[\+-]\d{2}:\d{2})?)$`)
	expirationDate := getExpirationDate(url, port)
	if !want.MatchString(expirationDate) {
		t.Fatalf(`getExpirationDate(url, port) = %q, want match for %#q, nil`, expirationDate, want)
	}
}
