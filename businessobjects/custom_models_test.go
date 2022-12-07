package businessobjects_test

import (
	"context"
	"net/http"
)

const stateInitial = "initial"
const stateStaged = "staged"
const statePublished = "published"

type MockHttpClient struct {
	req  *http.Request
	resp *http.Response
}

func (c *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	c.req = req
	return c.resp, nil
}

func validTokenFromContext(ctx context.Context) (string, error) {
	return "", nil
}

func validUriFromContext(ctx context.Context) (string, error) {
	return "", nil
}
