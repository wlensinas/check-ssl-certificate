package main

import (
	"regexp"
	"testing"
)

func TestGetExpirationDate(t *testing.T) {
	var url string = "www.google.com.ar"
	var port string = "443"
	want := regexp.MustCompile(`^((?:(\d{4}-\d{2}-\d{2})T(\d{2}:\d{2}:\d{2}(?:\.\d+)?))(Z|[\+-]\d{2}:\d{2})?)$`)
	expiration_date, err := getExpirationDate(url, port)
	if !want.MatchString(expiration_date) || err != nil {
		t.Fatalf(`getExpirationDate(url, port) = %q, %v, want match for %#q, nil`, expiration_date, err, want)
	}
}
