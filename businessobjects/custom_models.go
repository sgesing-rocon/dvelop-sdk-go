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

type Key struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	State       string `json:"state"`
	Type        string `json:"type"`
}

type Property struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	State       string `json:"state"`
	Type        string `json:"type"`
	Required    bool   `json:"required"`
	Indexed     bool   `json:"indexed"`
}

type EntityType struct {
	PluralName  string     `json:"pluralName"`
	Id          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	State       string     `json:"state"`
	Key         Key        `json:"key"`
	Properties  []Property `json:"properties"`
}

type CustomModel struct {
	Id          string       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	State       string       `json:"state"`
	EntityTypes []EntityType `json:"entityTypes"`
}

type getCustomModelsResponseDto struct {
	Items []CustomModel `json:"value"`
}

type GetCustomModelRequest struct {
	SystemBaseUri string
	AuthSessionId string
	ModelId       string
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

type ListCustomModelsRequest struct {
	SystemBaseUri string
	AuthSessionId string
}

func (c *DefaultClient) ListCustomModels(ctx context.Context, request ListCustomModelsRequest) ([]CustomModel, error) {
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

	modelList := getCustomModelsResponseDto{}
	err = json.NewDecoder(response.Body).Decode(&modelList)
	if err != nil {
		return nil, fmt.Errorf("error parsing response from business objects list custom models - error: %v", err.Error())
	}

	return modelList.Items, nil
}

type CreateCustomModelDtoProperty struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
	Required    bool   `json:"required"`
	Indexed     bool   `json:"indexed"`
}

type CreateCustomModelDtoKey struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
}

type CreateCustomModelDtoEntityType struct {
	Name        string                         `json:"name,omitempty"`
	Description string                         `json:"description,omitempty"`
	PluralName  string                         `json:"pluralName,omitempty"`
	Key         CreateCustomModelDtoKey        `json:"key,omitempty"`
	Properties  []CreateCustomModelDtoProperty `json:"properties,omitempty"`
}

type CreateCustomModelDto struct {
	Name        string                           `json:"name,omitempty"`
	Description string                           `json:"description,omitempty"`
	State       string                           `json:"state,omitempty"`
	EntityTypes []CreateCustomModelDtoEntityType `json:"entityTypes,omitempty"`
}

type CreateCustomModelRequest struct {
	SystemBaseUri string
	AuthSessionId string
	CustomModel   CreateCustomModelDto
}

func (c *DefaultClient) CreateCustomModel(ctx context.Context, request CreateCustomModelRequest) (CustomModel, error) {
	baseUri, authSessionId, err := c.getContextValues(ctx, request.SystemBaseUri, request.AuthSessionId)
	if err != nil {
		return CustomModel{}, err
	}

	uri, err := url.Parse(fmt.Sprintf("%v/businessobjects/core/models/customModels", baseUri))
	if err != nil {
		return CustomModel{}, fmt.Errorf("error parsing raw url with base uri '%v' - error: %v", baseUri, err.Error())
	}

	requestBody, err := json.Marshal(request.CustomModel)
	if err != nil {
		return CustomModel{}, fmt.Errorf("error marshalling body: %+v", request.CustomModel)
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
	if err != nil { // network,timout,etc - not bad status code
		return CustomModel{}, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return CustomModel{}, err
	}
	defer response.Body.Close()

	stringBody := string(body)
	fmt.Println(stringBody)

	if !isSuccessStatusCode(response.StatusCode) {
		errResponseModel := BusinessObjectsErrorResponse{}
		err = json.Unmarshal(body, &errResponseModel)
		if err != nil {
			return CustomModel{}, fmt.Errorf("error parsing error-response from business objects - error: %v", err.Error())
		}

		return CustomModel{}, errResponseModel.Error
	}

	model := CustomModel{}
	err = json.Unmarshal(body, &model)
	if err != nil {
		return CustomModel{}, fmt.Errorf("error parsing response from business objects create custom models - error: %v", err.Error())
	}

	return model, nil
}

type UpdateCustomModelDtoProperty struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
	Required    bool   `json:"required"`
	Indexed     bool   `json:"indexed"`
}

type UpdateCustomModelDtoKey struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
}

type UpdateCustomModelDtoEntityType struct {
	Id          string                         `json:"id,omitempty"`
	Name        string                         `json:"name,omitempty"`
	Description string                         `json:"description,omitempty"`
	PluralName  string                         `json:"pluralName,omitempty"`
	Key         UpdateCustomModelDtoKey        `json:"key,omitempty"`
	Properties  []UpdateCustomModelDtoProperty `json:"properties,omitempty"`
}

type UpdateCustomModelDto struct {
	Id          string                           `json:"id,omitempty"`
	Name        string                           `json:"name,omitempty"`
	Description string                           `json:"description,omitempty"`
	State       string                           `json:"state,omitempty"`
	EntityTypes []UpdateCustomModelDtoEntityType `json:"entityTypes,omitempty"`
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
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		errResponseModel := BusinessObjectsErrorResponse{}
		err = json.Unmarshal(body, &errResponseModel)
		if err != nil {
			return fmt.Errorf("error parsing error-response from business objects - error: %v", err.Error())
		}

		return errResponseModel.Error
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
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		errResponseModel := BusinessObjectsErrorResponse{}
		err = json.Unmarshal(body, &errResponseModel)
		if err != nil {
			return fmt.Errorf("error parsing error-response from business objects - error: %v", err.Error())
		}

		return errResponseModel.Error
	}

	return nil
}
