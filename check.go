package main

import (
	"fmt"
	"net"
	"time"
)

// Check tests TCP connectivity to a given destination and port.
// It returns a string describing whether the host is reachable.
//
// Parameters:
//   - destination: The domain or IP address to check.
//   - port: The TCP port as a string.
//
// The function attempts to connect within a 5-second timeout.
// If successful, it reports "[UP]" along with local and remote addresses.
// If unsuccessful, it reports "[DOWN]" with the error message.
func Check(destination, port string) string {
	address := destination + ":" + port       // Combine domain and port into an address https://example.com:80
	timeout := time.Duration(5 * time.Second) // Set a timeout for the connection

	conn, err := net.DialTimeout("tcp", address, timeout)

	var status string

	if err != nil {
		status = fmt.Sprintf("[DOWN] %v is unreachable: \n Error: %v", destination, err)
	} else {
		status = fmt.Sprintf(
			"[UP] %v is reachable. \n From: %v \n To: %v",
			destination,
			conn.LocalAddr(),
			conn.RemoteAddr(),
		)
	}

	return status
}
