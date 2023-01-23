package businessobjects_test

// import (
// 	"context"
// )

// const stateInitial = "initial"
// const stateStaged = "staged"
// const statePublished = "published"

// func validTokenFromContext(ctx context.Context) (string, error) {
// 	return "abc", nil
// }

// func validUriFromContext(ctx context.Context) (string, error) {
// 	return "abc", nil
// }

// import (
// 	"context"
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/d-velop/dvelop-sdk-go/businessobjects"
// )

// func getValidMockClient(t *testing.T) businessobjects.DefaultClient {

// 	client := businessobjects.NewClient(validUriFromContext)

// 	ts := httptest.NewServer(
// 		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			fmt.Fprintln(w, "Hello, client")
// 		}))

// 	// directly on client
// 	client.DefaultAuthSessionIdFromContext = validTokenFromContext
// 	client.HttpClient = ts.Client()

// 	// or using setter functions, allowing validation
// 	err := client.SetAuthSessionFromContextFunction(validTokenFromContext)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	err = client.SetHttpClient(http.DefaultClient)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	return client
// }

// func TestDefaultValueTestBOReturns200_GetCustomModels_ReturnsModels(t *testing.T) {
// 	ts := httptest.NewServer(
// 		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			w.WriteHeader(http.StatusTeapot)
// 			fmt.Fprintln(w, `{
// 									"value": [
// 										{
// 											"id": "2cee33e2-6f7c-4a03-a1c2-e7e5433c5127",
// 											"name": "TestfromSDK",
// 											"state": "initial",
// 											"description": "test description",
// 											"entityTypes": []
// 										}
// 									]
// 					}`)
// 		}))
// 	defer ts.Close()

// 	urlFunc := func(ctx context.Context) (string, error) {
// 		return ts.URL, nil
// 	}

// 	c := businessobjects.NewClient(urlFunc)

// 	// directly on client
// 	c.DefaultAuthSessionIdFromContext = validTokenFromContext
// 	c.HttpClient = ts.Client()

// 	got, err := c.GetCustomModels(context.Background())
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if len(got) == 0 {
// 		t.Errorf("no models returned")
// 	}
// }

// func getMockedServer(t *testing.T, validateRequest func(*http.Request) error, mockedHttpResponse http.Response) (httptest.Server, error) {
// 	httptest.NewServer(
// 		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 			validationError := validateRequest(r)
// 			if validationError != nil {
// 				t.Fatal("request validation failed", validationError)
// 			}
// 			// TODO: return Response
// 		}))

// 	ts := httptest.NewServer(
// 		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			w.WriteHeader(http.StatusTeapot)
// 			fmt.Fprintln(w, `{
// 									"value": [
// 										{
// 											"id": "2cee33e2-6f7c-4a03-a1c2-e7e5433c5127",
// 											"name": "TestfromSDK",
// 											"state": "initial",
// 											"description": "test description",
// 											"entityTypes": []
// 										}
// 									]
// 					}`)
// 		}))
// }

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
