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

type GetEntityTypesRequest struct {
	SystemBaseUri string
	AuthSessionId string
	ModelId       string
}

func (c *DefaultClient) GetEntityTypes(ctx context.Context, request GetEntityTypesRequest) ([]EntityType, error) {
	baseUri, authSessionId, err := c.getContextValues(ctx, request.SystemBaseUri, request.AuthSessionId)
	if err != nil {
		return nil, err
	}

	uri, err := url.Parse(fmt.Sprintf("%v/businessobjects/core/models/customModels(%v)/entityTypes", baseUri, request.ModelId))
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

	modelList := EntityTypeList{}
	err = json.NewDecoder(response.Body).Decode(&modelList)
	if err != nil {
		return nil, fmt.Errorf("error parsing response from business objects get entity types - error: %v", err.Error())
	}

	return modelList.Items, nil
}

type GetEntityTypeRequest struct {
	SystemBaseUri string
	AuthSessionId string
	ModelId       string
	EntityId      string
}

func (c *DefaultClient) GetEntityType(ctx context.Context, request GetEntityTypeRequest) (EntityType, error) {
	baseUri, authSessionId, err := c.getContextValues(ctx, request.SystemBaseUri, request.AuthSessionId)
	if err != nil {
		return EntityType{}, err
	}

	uri, err := url.Parse(fmt.Sprintf("%v/businessobjects/core/models/customModels(%v)/entityTypes(%v)", baseUri, request.ModelId, request.EntityId))
	if err != nil {
		return EntityType{}, fmt.Errorf("error parsing raw url with base uri '%v' - error: %v", baseUri, err.Error())
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
		return EntityType{}, errors.New("http request failed with status code: " + strconv.Itoa(response.StatusCode))
	}
	defer response.Body.Close()

	model := EntityType{}
	err = json.NewDecoder(response.Body).Decode(&model)
	if err != nil {
		return EntityType{}, fmt.Errorf("error parsing response from business objects get entity type - error: %v", err.Error())
	}

	return model, nil
}

type CreateEntityTypeRequest struct {
	SystemBaseUri string
	AuthSessionId string
	ModelId       string
	Data          EntityType
}

func (c *DefaultClient) CreateEntityType(ctx context.Context, request CreateEntityTypeRequest) (string, error) {
	baseUri, authSessionId, err := c.getContextValues(ctx, request.SystemBaseUri, request.AuthSessionId)
	if err != nil {
		return "", err
	}

	uri, err := url.Parse(fmt.Sprintf("%v/businessobjects/core/models/customModels(%v)/entityTypes", baseUri, request.ModelId))
	if err != nil {
		return "", fmt.Errorf("error parsing raw url with base uri '%v' - error: %v", baseUri, err.Error())
	}

	requestBody, err := json.Marshal(request.Data)
	if err != nil {
		return "", fmt.Errorf("error marshalling body: %+v", request.Data)
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

	model := EntityType{}
	err = json.Unmarshal(body, &model)
	if err != nil {
		return "", fmt.Errorf("error parsing response from business objects create entity type - error: %v", err.Error())
	}

	return model.Id, nil
}

type ReplaceEntityTypeRequest struct {
	SystemBaseUri string
	AuthSessionId string
	ModelId       string
	EntityId      string
	Data          EntityType
}

func (c *DefaultClient) ReplaceEntityType(ctx context.Context, request ReplaceEntityTypeRequest) error {
	baseUri, authSessionId, err := c.getContextValues(ctx, request.SystemBaseUri, request.AuthSessionId)
	if err != nil {
		return err
	}

	uri, err := url.Parse(fmt.Sprintf("%v/businessobjects/core/models/customModels(%v)/entityTypes(%v)", baseUri, request.ModelId, request.EntityId))
	if err != nil {
		return fmt.Errorf("error parsing raw url with base uri '%v' - error: %v", baseUri, err.Error())
	}

	requestBody, err := json.Marshal(request.Data)
	if err != nil {
		return fmt.Errorf("error marshalling body: %+v", request.Data)
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

type PartiallyUpdateEntityTypeRequest struct {
	SystemBaseUri string
	AuthSessionId string
	ModelId       string
	EntityId      string
	Data          EntityType
}

func (c *DefaultClient) PartiallyUpdateEntityType(ctx context.Context, request PartiallyUpdateEntityTypeRequest) error {
	baseUri, authSessionId, err := c.getContextValues(ctx, request.SystemBaseUri, request.AuthSessionId)
	if err != nil {
		return err
	}

	uri, err := url.Parse(fmt.Sprintf("%v/businessobjects/core/models/customModels(%v)/entityTypes(%v)", baseUri, request.ModelId, request.EntityId))
	if err != nil {
		return fmt.Errorf("error parsing raw url with base uri '%v' - error: %v", baseUri, err.Error())
	}

	requestBody, err := json.Marshal(request.Data)
	if err != nil {
		return fmt.Errorf("error marshalling body: %+v", request.Data)
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

type DeleteEntityTypeRequest struct {
	SystemBaseUri string
	AuthSessionId string
	ModelId       string
	EntityId      string
}

func (c *DefaultClient) DeleteEntityType(ctx context.Context, request DeleteEntityTypeRequest) error {
	baseUri, authSessionId, err := c.getContextValues(ctx, request.SystemBaseUri, request.AuthSessionId)
	if err != nil {
		return err
	}

	uri, err := url.Parse(fmt.Sprintf("%v/businessobjects/core/models/customModels(%v)/entityTypes(%v)", baseUri, request.ModelId, request.EntityId))
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
