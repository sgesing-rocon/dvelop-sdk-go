package businessobjects_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/d-velop/dvelop-sdk-go/businessobjects"
)

func BuilderTestBOReturns200_GetCustomModels_ReturnsModels(t *testing.T) {
	c, err := businessobjects.NewBuilder(validUriFromContext).
		AddHttpClient(http.DefaultClient).
		AddAuthSessionFromContextFunction(validTokenFromContext).
		Build()
	if err != nil {
		t.Error(err)
	}

	got, err := c.GetCustomModels(context.Background())
	if err != nil {
		t.Error(err)
	}

	if len(got) == 0 {
		t.Errorf("no models returned")
	}
}

//func TestBOReturns201_CreateCustomModel_ReturnsId(t *testing.T) {
//	c := NewOptions(validUriFromContext, validTokenFromContext)
//
//	newModel := CreateCustomModelParams{
//		Name:        "TestfromSDK",
//		State:       stateInitial,
//		Description: "test description",
//	}
//	testId := "meineId"
//
//	id, err := c.CreateCustomModel(context.Background(), newModel)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	if id != testId {
//		t.Errorf("Wrong id returned - expected: %v got: %v", testId, id)
//	}
//}
//
//func TestBOReturns201_UpdateCustomModel_ReturnsId(t *testing.T) {
//	c := NewOptions(validUriFromContext, validTokenFromContext)
//
//	newModel := CreateCustomModelParams{
//		Name:        "TestfromSDK",
//		State:       stateInitial,
//		Description: "test description",
//	}
//	testId := "meineId"
//
//	id, err := c.CreateCustomModel(context.Background(), newModel)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	if id != testId {
//		t.Errorf("Wrong id returned - expected: %v got: %v", testId, id)
//	}
//}
