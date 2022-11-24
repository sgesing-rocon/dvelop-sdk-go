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
)

type CustomModelList struct {
	Items []CustomModel `json:"value"`
}

type CustomModel struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	State       string `json:"state"`
	Description string `json:"description"`
}

type CreateCustomModelParams struct {
	Name        string `json:"name"`
	State       string `json:"state"`
	Description string `json:"description"`
}

type EntityType struct {
	Id         string     `json:"id"`
	Name       string     `json:"name"`
	PluralName string     `json:"pluralName"`
	State      string     `json:"state"`
	Key        Key        `json:"key"`
	Properties []Property `json:"properties"`
}

type Property struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Required bool   `json:"required"`
	Indexed  bool   `json:"indexed"`
	Type     string `json:"type"`
	State    string `json:"state"`
}

type Key struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	State string `json:"state"`
}

func (c *client) GetCustomModels(ctx context.Context) ([]CustomModel, error) {
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

func (c *client) GetCustomModel(ctx context.Context, id string) (CustomModel, error) {
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

func (c *client) CreateCustomModel(ctx context.Context, params CreateCustomModelParams) (string, error) {
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

func (c *client) UpdateCustomModel(ctx context.Context, params CustomModel) error {
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

func (c *client) PartiallyUpdateCustomModel(ctx context.Context, params CustomModel) error {
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

func (c *client) DeleteCustomModel(ctx context.Context, id string) error {
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

func (c *client) getContextValues(ctx context.Context) (string, string, error) {
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
