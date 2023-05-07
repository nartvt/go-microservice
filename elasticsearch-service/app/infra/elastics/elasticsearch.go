package elastics

import (
	"crypto/tls"
	"net/http"
)

func load() {
	// Enable HTTPS
	cfg := elasticsearch.Config{
		Addresses: []string{"https://localhost:9200"},
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{MinVersion: tls.VersionTLS12},
		},
	}

	// Configure credentials
	cfg.Username = "my-username"
	cfg.Password = "my-password"

	// Create the client
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		// Handle error
	}
}
