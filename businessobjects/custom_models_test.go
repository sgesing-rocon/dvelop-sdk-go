package businessobjects

import (
	"context"
	"net/http"
	"testing"
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

func TestBOReturns200_GetCustomModels_ReturnsModels(t *testing.T) {
	c := New(validUriFromContext, validTokenFromContext)

	got, err := c.GetCustomModels(context.Background())
	if err != nil {
		t.Error(err)
	}

	if len(got) == 0 {
		t.Errorf("no models returned")
	}
}

func TestBOReturns201_CreateCustomModel_ReturnsId(t *testing.T) {
	c := New(validUriFromContext, validTokenFromContext)

	newModel := CreateCustomModelParams{
		Name:        "TestfromSDK",
		State:       stateInitial,
		Description: "test description",
	}
	testId := "meineId"

	id, err := c.CreateCustomModel(context.Background(), newModel)
	if err != nil {
		t.Error(err)
		return
	}
	if id != testId {
		t.Errorf("Wrong id returned - expected: %v got: %v", testId, id)
	}
}

func TestBOReturns201_UpdateCustomModel_ReturnsId(t *testing.T) {
	c := New(validUriFromContext, validTokenFromContext)

	newModel := CreateCustomModelParams{
		Name:        "TestfromSDK",
		State:       stateInitial,
		Description: "test description",
	}
	testId := "meineId"

	id, err := c.CreateCustomModel(context.Background(), newModel)
	if err != nil {
		t.Error(err)
		return
	}
	if id != testId {
		t.Errorf("Wrong id returned - expected: %v got: %v", testId, id)
	}
}
