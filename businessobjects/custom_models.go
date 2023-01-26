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
)

type FullCustomModelDto struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	State       string `json:"state"`
	EntityTypes []struct {
		PluralName  string `json:"pluralName"`
		Id          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		State       string `json:"state"`
		Key         struct {
			Id          string `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			State       string `json:"state"`
			Type        string `json:"type"`
		} `json:"key"`
		Properties []struct {
			Id          string `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			State       string `json:"state"`
			Type        string `json:"type"`
			Required    bool   `json:"required"`
			Indexed     bool   `json:"indexed"`
		} `json:"properties"`
	} `json:"entityTypes"`
}

type customModelList struct {
	Items []FullCustomModelDto `json:"value"`
}

type GetCustomModelsRequest struct {
	SystemBaseUri string
	AuthSessionId string
}

func (c *DefaultClient) GetCustomModels(ctx context.Context, request GetCustomModelsRequest) ([]FullCustomModelDto, error) {
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

	modelList := customModelList{}
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

func (c *DefaultClient) GetCustomModel(ctx context.Context, request GetCustomModelRequest) (FullCustomModelDto, error) {
	baseUri, authSessionId, err := c.getContextValues(ctx, request.SystemBaseUri, request.AuthSessionId)
	if err != nil {
		return FullCustomModelDto{}, err
	}

	uri, err := url.Parse(fmt.Sprintf("%v/businessobjects/core/models/customModels(%v)", baseUri, request.ModelId))
	if err != nil {
		return FullCustomModelDto{}, fmt.Errorf("error parsing raw url with base uri '%v' - error: %v", baseUri, err.Error())
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
		return FullCustomModelDto{}, errors.New("http request failed with status code: " + strconv.Itoa(response.StatusCode))
	}
	defer response.Body.Close()

	model := FullCustomModelDto{}
	err = json.NewDecoder(response.Body).Decode(&model)
	if err != nil {
		return FullCustomModelDto{}, fmt.Errorf("error parsing response from business objects get custom models - error: %v", err.Error())
	}

	return model, nil
}

type CreateCustomModelDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	State       string `json:"state"`
	EntityTypes []struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		PluralName  string `json:"pluralName"`
		Key         struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"key"`
		Properties []struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Type        string `json:"type"`
			Required    bool   `json:"required"`
			Indexed     bool   `json:"indexed"`
		} `json:"properties,omitempty"`
	} `json:"entityTypes,omitempty"`
}

type CreateCustomModelRequest struct {
	SystemBaseUri string
	AuthSessionId string
	CustomModel   CreateCustomModelDto
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

	requestBody, err := json.Marshal(request.CustomModel)
	if err != nil {
		return "", fmt.Errorf("error marshalling body: %+v", request.CustomModel)
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

	model := FullCustomModelDto{}
	err = json.Unmarshal(body, &model)
	if err != nil {
		return "", fmt.Errorf("error parsing response from business objects create custom models - error: %v", err.Error())
	}

	return model.Id, nil
}

type UpdateCustomModelDto struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	State       string `json:"state"`
	EntityTypes []struct {
		PluralName  string `json:"pluralName"`
		Id          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Key         struct {
			Id          string `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Type        string `json:"type"`
		} `json:"key"`
		Properties []struct {
			Id          string `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Type        string `json:"type"`
			Required    bool   `json:"required"`
			Indexed     bool   `json:"indexed"`
		} `json:"properties,omitempty"`
	} `json:"entityTypes,omitempty"`
}

type UpdateCustomModelRequest struct {
	SystemBaseUri string
	AuthSessionId string
	CustomModel   UpdateCustomModelDto
}

func (c *DefaultClient) UpdateCustomModel(ctx context.Context, request UpdateCustomModelRequest) error {
	baseUri, authSessionId, err := c.getContextValues(ctx, request.SystemBaseUri, request.AuthSessionId)
	if err != nil {
		return err
	}

	uri, err := url.Parse(fmt.Sprintf("%v/businessobjects/core/models/customModels(%v)", baseUri, request.CustomModel.Id))
	if err != nil {
		return fmt.Errorf("error parsing raw url with base uri '%v' - error: %v", baseUri, err.Error())
	}

	requestBody, err := json.Marshal(request.CustomModel)
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
