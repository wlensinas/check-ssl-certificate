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

func getExpirationDate(url, port string) (string, error) {
	connection, err := tls.Dial("tcp", buildURL(url, port), nil)
	if err != nil {
		return "", err
	}

	expiry := getExpiryFromConnection(connection)

	return expiry.Format(time.RFC3339), nil
}
