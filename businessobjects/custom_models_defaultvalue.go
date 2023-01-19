package businessobjects

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type DefaultClient struct {
	DefaultSystemBaseUriFromContext func(ctx context.Context) (string, error)
	DefaultAuthSessionIdFromContext func(ctx context.Context) (string, error)
	HttpClient                      *http.Client
}

func NewClient() DefaultClient {
	c := DefaultClient{
		HttpClient: &http.Client{Timeout: 10 * time.Second},
	}

	return c
}

type GetCustomModelsRequest struct {
	SystemBaseUri string
	AuthSessionId string
}

type CustomModelList struct {
	Items []CustomModel `json:"value"`
}

func (c *DefaultClient) GetCustomModels(ctx context.Context, request GetCustomModelsRequest) ([]CustomModel, error) {
	baseUri, authSessionId, err := c.getContextValues(ctx, request.SystemBaseUri, request.AuthSessionId)
	if err != nil {
		return nil, err
	}

	uri, err := url.Parse(baseUri + "/businessobjects/core/models/customModels")
	if err != nil {
		return nil, fmt.Errorf("error parsing raw url with base uri '%v' - error: %v", baseUri, err.Error())
	}

	req := &http.Request{
		Method: http.MethodGet,
		URL:    uri,
		Header: map[string][]string{
			"Accept":        {"application/json"},
			"Authorization": {"Bearer " + authSessionId},
		},
	}

	response, err := c.HttpClient.Do(req)
	if err != nil {
		log.Fatalf("error sending request to business objects. %+v", err)
	}
	if !isSuccessStatusCode(response.StatusCode) {
		return nil, errors.New("http request failed with status code: " + strconv.Itoa(response.StatusCode))
	}
	defer response.Body.Close()

	modelList := CustomModelList{}
	err = json.NewDecoder(response.Body).Decode(&modelList)
	if err != nil {
		return nil, fmt.Errorf("error parsing response from business objects list custom models - error: %v", err.Error())
	}

	return modelList.Items, nil
}

type GetCustomModelRequest struct {
	SystemBaseUri string
	AuthSessionId string
	ModelId       string
}

type CustomModel struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	State       string `json:"state"`
	Description string `json:"description"`
}

func (c *DefaultClient) GetCustomModel(ctx context.Context, request GetCustomModelRequest) (CustomModel, error) {
	baseUri, authSessionId, err := c.getContextValues(ctx, request.SystemBaseUri, request.AuthSessionId)
	if err != nil {
		return CustomModel{}, err
	}

	uri, err := url.Parse(fmt.Sprintf("%v/businessobjects/core/models/customModels(%v)", baseUri, request.ModelId))
	if err != nil {
		return CustomModel{}, fmt.Errorf("error parsing raw url with base uri '%v' - error: %v", baseUri, err.Error())
	}

	req := &http.Request{
		Method: http.MethodGet,
		URL:    uri,
		Header: map[string][]string{
			"Accept":        {"application/json"},
			"Authorization": {"Bearer " + authSessionId},
		},
	}

	response, err := c.HttpClient.Do(req)
	if err != nil {
		log.Fatalf("error sending request to business objects. %+v", err)
	}
	if !isSuccessStatusCode(response.StatusCode) {
		return CustomModel{}, errors.New("http request failed with status code: " + strconv.Itoa(response.StatusCode))
	}
	defer response.Body.Close()

	model := CustomModel{}
	err = json.NewDecoder(response.Body).Decode(&model)
	if err != nil {
		return CustomModel{}, fmt.Errorf("error parsing response from business objects get custom models - error: %v", err.Error())
	}

	return model, nil
}

type CreateCustomModelRequest struct {
	SystemBaseUri string `json:"-"`
	AuthSessionId string `json:"-"`
	Name          string `json:"name"`
	State         string `json:"state"`
	Description   string `json:"description"`
}

func (c *DefaultClient) CreateCustomModel(ctx context.Context, request CreateCustomModelRequest) (string, error) {
	baseUri, authSessionId, err := c.getContextValues(ctx, request.SystemBaseUri, request.AuthSessionId)
	if err != nil {
		return "", err
	}

	uri, err := url.Parse(fmt.Sprintf("%v/businessobjects/core/models/customModels", baseUri))
	if err != nil {
		return "", fmt.Errorf("error parsing raw url with base uri '%v' - error: %v", baseUri, err.Error())
	}

	requestBody, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("error marshalling body: %+v", request)
	}

	req := &http.Request{
		Method: http.MethodPost,
		URL:    uri,
		Header: map[string][]string{
			"Accept":        {"application/json"},
			"Authorization": {"Bearer " + authSessionId},
			"Origin":        {baseUri},
			"Content-Type":  {"application/json"},
		},
		Body: io.NopCloser(bytes.NewReader(requestBody)),
	}

	response, err := c.HttpClient.Do(req)
	if err != nil {
		return "", err
	}
	if !isSuccessStatusCode(response.StatusCode) {
		return "", errors.New("http request failed with status code: " + strconv.Itoa(response.StatusCode))
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	stringBody := string(body)
	fmt.Println(stringBody)

	model := CustomModel{}
	err = json.Unmarshal(body, &model)
	if err != nil {
		return "", fmt.Errorf("error parsing response from business objects create custom models - error: %v", err.Error())
	}

	return model.Id, nil
}

type UpdateCustomModelRequest struct {
	SystemBaseUri string `json:"-"`
	AuthSessionId string `json:"-"`
	ModelId       string `json:"id"`
	Name          string `json:"name"`
	State         string `json:"state"`
	Description   string `json:"description"`
}

func (c *DefaultClient) UpdateCustomModel(ctx context.Context, request UpdateCustomModelRequest) error {
	baseUri, authSessionId, err := c.getContextValues(ctx, request.SystemBaseUri, request.AuthSessionId)
	if err != nil {
		return err
	}

	uri, err := url.Parse(fmt.Sprintf("%v/businessobjects/core/models/customModels(%v)", baseUri, request.ModelId))
	if err != nil {
		return fmt.Errorf("error parsing raw url with base uri '%v' - error: %v", baseUri, err.Error())
	}

	requestBody, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("error marshalling body: %+v", request)
	}

	req := &http.Request{
		Method: http.MethodPut,
		URL:    uri,
		Header: map[string][]string{
			"Accept":        {"application/json"},
			"Authorization": {"Bearer " + authSessionId},
			"Origin":        {baseUri},
			"Content-Type":  {"application/json"},
		},
		Body: io.NopCloser(bytes.NewReader(requestBody)),
	}

	response, err := c.HttpClient.Do(req)
	if err != nil {
		return err
	}
	if !isSuccessStatusCode(response.StatusCode) {
		return errors.New("http request failed with status code: " + strconv.Itoa(response.StatusCode))
	}

	return nil
}

type PartiallyUpdateCustomModelRequest struct {
	SystemBaseUri string `json:"-"`
	AuthSessionId string `json:"-"`
	ModelId       string `json:"id"`
	Name          string `json:"name,omitempty"`
	State         string `json:"state,omitempty"`
	Description   string `json:"description,omitempty"`
}

func (c *DefaultClient) PartiallyUpdateCustomModel(ctx context.Context, request PartiallyUpdateCustomModelRequest) error {
	baseUri, authSessionId, err := c.getContextValues(ctx, request.SystemBaseUri, request.AuthSessionId)
	if err != nil {
		return err
	}

	uri, err := url.Parse(fmt.Sprintf("%v/businessobjects/core/models/customModels(%v)", baseUri, request.ModelId))
	if err != nil {
		return fmt.Errorf("error parsing raw url with base uri '%v' - error: %v", baseUri, err.Error())
	}

	requestBody, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("error marshalling body: %+v", request)
	}

	req := &http.Request{
		Method: http.MethodPatch,
		URL:    uri,
		Header: map[string][]string{
			"Accept":        {"application/json"},
			"Authorization": {"Bearer " + authSessionId},
			"Origin":        {baseUri},
			"Content-Type":  {"application/json"},
		},
		Body: io.NopCloser(bytes.NewReader(requestBody)),
	}

	response, err := c.HttpClient.Do(req)
	if err != nil {
		return err
	}
	if !isSuccessStatusCode(response.StatusCode) {
		return errors.New("http request failed with status code: " + strconv.Itoa(response.StatusCode))
	}

	return nil
}

type DeleteCustomModelRequest struct {
	SystemBaseUri string
	AuthSessionId string
	ModelId       string
}

func (c *DefaultClient) DeleteCustomModel(ctx context.Context, request DeleteCustomModelRequest) error {
	baseUri, authSessionId, err := c.getContextValues(ctx, request.SystemBaseUri, request.AuthSessionId)
	if err != nil {
		return err
	}

	uri, err := url.Parse(fmt.Sprintf("%v/businessobjects/core/models/customModels(%v)", baseUri, request.ModelId))
	if err != nil {
		return fmt.Errorf("error parsing raw url with base uri '%v' - error: %v", baseUri, err.Error())
	}

	req := &http.Request{
		Method: http.MethodDelete,
		URL:    uri,
		Header: map[string][]string{
			"Accept":        {"application/json"},
			"Authorization": {"Bearer " + authSessionId},
			"Origin":        {baseUri},
		},
	}

	response, err := c.HttpClient.Do(req)
	if err != nil {
		return err
	}
	if !isSuccessStatusCode(response.StatusCode) {
		return errors.New("http request failed with status code: " + strconv.Itoa(response.StatusCode))
	}

	return nil
}

func (c *DefaultClient) getContextValues(ctx context.Context, systemBaseUriFromRequest string, authSessionIdFromRequest string) (string, string, error) {
	var systemBaseUri string
	var authSessionId string
	var err error

	if systemBaseUriFromRequest != "" {
		systemBaseUri = systemBaseUriFromRequest
	} else if c.DefaultSystemBaseUriFromContext != nil {
		systemBaseUri, err = c.DefaultSystemBaseUriFromContext(ctx)

		if err != nil {
			return "", "", err
		}
	} else {
		return "", "", errors.New("missing SystemBaseUri")
	}

	if authSessionIdFromRequest != "" {
		authSessionId = authSessionIdFromRequest
	} else if c.DefaultAuthSessionIdFromContext != nil {
		authSessionId, err = c.DefaultAuthSessionIdFromContext(ctx)

		if err != nil {
			return "", "", err
		}
	}

	return systemBaseUri, authSessionId, nil
}

func isSuccessStatusCode(statusCode int) bool {
	return statusCode >= 200 && statusCode < 300
}
