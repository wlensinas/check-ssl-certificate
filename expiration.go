package main

import (
	"crypto/tls"
	"fmt"
	"time"
)

func buildURL(url, port string) string {
	return fmt.Sprintf("%s:%s", url, port)
}

func getExpiryFromConnection(connection *tls.Conn) time.Time {
	return connection.ConnectionState().PeerCertificates[0].NotAfter
}

func getExpirationDate(url, port string) string {
	connection, err := tls.Dial("tcp", buildURL(url, port), nil)
	if err != nil {
		return ""
	}

	expiry := getExpiryFromConnection(connection)

	return expiry.Format(time.RFC3339)
}

func getDurationFromDates(actualDate, certificateSSLDate string) (time.Duration, error) {
	d1, err := time.Parse(time.RFC3339, actualDate)
	if err != nil {
		fmt.Println("Error to parse actualDate:", err)
		return 0, err
	}

	d2, err := time.Parse(time.RFC3339, certificateSSLDate)
	if err != nil {
		fmt.Println("Error to parse certificateSSLDate:", err)
		return 0, err
	}
	return d2.Sub(d1), nil
}
