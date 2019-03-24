package pipe

import (
	"context"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

// Client has the information needed to work as a pipehub pipe.
type Client struct {
	url *url.URL
}

func (c *Client) init(config map[string]interface{}) error {
	rawHost, ok := config["host"].(string)
	if !ok {
		return errors.New("casting host to string error")
	}

	var err error
	c.url, err = url.Parse(rawHost)
	if err != nil {
		return errors.Wrapf(err, "parse url '%s' error", rawHost)
	}

	return nil
}

// Close the client.
func (Client) Close(ctx context.Context) error {
	return nil
}

// Default is a sample HTTP handler is just a dump proxy, it does nothing.
func (c Client) Default(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		r.Host = c.url.Host
		r.URL.Host = c.url.Host
		r.URL.Scheme = c.url.Scheme

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// NewClient return a initialized client.
func NewClient(config map[string]interface{}) (Client, error) {
	var c Client
	if err := c.init(config); err != nil {
		return c, errors.Wrap(err, "initialization error")
	}
	return c, nil
}
