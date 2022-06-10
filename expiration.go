package main

import (
	"crypto/tls"
	"time"
)

func getExpirationDate(url string, port string) (string, error) {
	conn, err := tls.Dial("tcp", url+":"+port, nil)
	if err != nil {
		panic("SSL certificate err: " + err.Error())
	}
	expiry := conn.ConnectionState().PeerCertificates[0].NotAfter
	return expiry.Format(time.RFC3339), err
}
