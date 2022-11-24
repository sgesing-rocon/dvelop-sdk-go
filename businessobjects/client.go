package businessobjects

import (
	"context"
	"log"
	"net/http"
	"time"
)

type client struct {
	baseUriFromContext   func(ctx context.Context) (string, error)
	authSessionIdFromCtx func(ctx context.Context) (string, error)
	httpClient           *http.Client
}

func New(baseUriFromContext, authSessionIdFromCtx func(ctx context.Context) (string, error)) *client {
	c := &client{
		authSessionIdFromCtx: authSessionIdFromCtx,
		baseUriFromContext:   baseUriFromContext,
	}
	c = c.WithHttpClient(nil)

	log.Println()

	return c
}

func (c *client) WithHttpClient(httpClient *http.Client) *client {
	if httpClient != nil {
		c.httpClient = httpClient
	} else {
		c.httpClient = &http.Client{
			Timeout: 10 * time.Second,
		}
	}

	return c
}
