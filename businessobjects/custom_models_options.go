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
	"time"
)

type optionsClient struct {
	baseUriFromContext   func(ctx context.Context) (string, error)
	authSessionIdFromCtx func(ctx context.Context) (string, error)
	httpClient           *http.Client
}

type Option func(*optionsClient) error

func NewOptions(baseUriFromContext func(ctx context.Context) (string, error), options ...Option) (*optionsClient, error) {
	c := &optionsClient{
		baseUriFromContext: baseUriFromContext,
		httpClient:         &http.Client{Timeout: 10 * time.Second},
	}

	// Apply all the optional configs
	for _, applyOption := range options {
		err := applyOption(c)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

func WithHttpClient(httpClient *http.Client) Option {
	return func(c *optionsClient) error {
		if httpClient == nil {
			return errors.New("httpClient must not be nil")
		}

		c.httpClient = httpClient
		return nil
	}
}

func WithAuthSessionIdFromCtx(auth func(ctx context.Context) (string, error)) Option {
	return func(c *optionsClient) error {
		if auth == nil {
			return errors.New("function for auth session retrieval must not be nil")
		}

		c.authSessionIdFromCtx = auth
		return nil
	}
}

func (c *optionsClient) GetCustomModels(ctx context.Context) ([]CustomModel, error) {
	baseUri, authSessionId, err := c.getContextValues(ctx)
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

	response, err := c.httpClient.Do(req)
	if err != nil {
		log.Fatalf("error sending request to business objects. %+v", err)
	}
	defer response.Body.Close()

	modelList := CustomModelList{}
	err = json.NewDecoder(response.Body).Decode(&modelList)
	if err != nil {
		return nil, fmt.Errorf("error parsing response from business objects list custom models - error: %v", err.Error())
	}

	return modelList.Items, nil
}

func (c *optionsClient) GetCustomModel(ctx context.Context, id string) (CustomModel, error) {
	baseUri, authSessionId, err := c.getContextValues(ctx)
	if err != nil {
		return CustomModel{}, err
	}

	uri, err := url.Parse(fmt.Sprintf("%v/businessobjects/core/models/customModels(%v)", baseUri, id))
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

	response, err := c.httpClient.Do(req)
	if err != nil {
		log.Fatalf("error sending request to business objects. %+v", err)
	}
	defer response.Body.Close()

	model := CustomModel{}
	err = json.NewDecoder(response.Body).Decode(&model)
	if err != nil {
		return CustomModel{}, fmt.Errorf("error parsing response from business objects get custom models - error: %v", err.Error())
	}

	return model, nil
}

func (c *optionsClient) CreateCustomModel(ctx context.Context, params CreateCustomModelParams) (string, error) {
	baseUri, authSessionId, err := c.getContextValues(ctx)
	if err != nil {
		return "", err
	}

	uri, err := url.Parse(fmt.Sprintf("%v/businessobjects/core/models/customModels", baseUri))
	if err != nil {
		return "", fmt.Errorf("error parsing raw url with base uri '%v' - error: %v", baseUri, err.Error())
	}

	requestBody, err := json.Marshal(params)
	if err != nil {
		return "", fmt.Errorf("error marshalling body: %+v", params)
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

	response, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
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

func (c *optionsClient) UpdateCustomModel(ctx context.Context, params CustomModel) error {
	baseUri, authSessionId, err := c.getContextValues(ctx)
	if err != nil {
		return err
	}

	uri, err := url.Parse(fmt.Sprintf("%v/businessobjects/core/models/customModels(%v)", baseUri, params.Id))
	if err != nil {
		return fmt.Errorf("error parsing raw url with base uri '%v' - error: %v", baseUri, err.Error())
	}

	requestBody, err := json.Marshal(params)
	if err != nil {
		return fmt.Errorf("error marshalling body: %+v", params)
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

	response, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	stringBody := string(body)
	fmt.Println(stringBody)

	model := CustomModel{}
	err = json.Unmarshal(body, &model)
	if err != nil {
		return fmt.Errorf("error parsing response from business objects update custom models - error: %v", err.Error())
	}

	return nil
}

func (c *optionsClient) PartiallyUpdateCustomModel(ctx context.Context, params CustomModel) error {
	baseUri, authSessionId, err := c.getContextValues(ctx)
	if err != nil {
		return err
	}

	uri, err := url.Parse(fmt.Sprintf("%v/businessobjects/core/models/customModels(%v)", baseUri, params.Id))
	if err != nil {
		return fmt.Errorf("error parsing raw url with base uri '%v' - error: %v", baseUri, err.Error())
	}

	requestBody, err := json.Marshal(params)
	if err != nil {
		return fmt.Errorf("error marshalling body: %+v", params)
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

	response, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	stringBody := string(body)
	fmt.Println(stringBody)

	model := CustomModel{}
	err = json.Unmarshal(body, &model)
	if err != nil {
		return fmt.Errorf("error parsing response from business objects update custom models - error: %v", err.Error())
	}

	return nil
}

func (c *optionsClient) DeleteCustomModel(ctx context.Context, id string) error {
	baseUri, authSessionId, err := c.getContextValues(ctx)
	if err != nil {
		return err
	}

	uri, err := url.Parse(fmt.Sprintf("%v/businessobjects/core/models/customModels(%v)", baseUri, id))
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

	response, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	stringBody := string(body)
	fmt.Println(stringBody)
	return nil
}

func (c *optionsClient) getContextValues(ctx context.Context) (string, string, error) {
	baseUri, err := c.baseUriFromContext(ctx)
	if err != nil || baseUri == "" {
		return "", "", errors.New("missing base uri")
	}
	authSessionId, err := c.authSessionIdFromCtx(ctx)
	if err != nil || authSessionId == "" {
		return "", "", errors.New("missing authSessionId")
	}

	return baseUri, authSessionId, nil
}
