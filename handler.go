package pipe

import (
	"context"
	"net/http"
)

// Client has the information needed to work as a pipehub pipe.
type Client struct{}

// Close the client.
func (Client) Close(ctx context.Context) error {
	return nil
}

// Default is a sample HTTP handler is just a dump proxy, it does nothing.
func (Client) Default(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// NewClient return a initialized client.
func NewClient() (Client, error) {
	return Client{}, nil
}
