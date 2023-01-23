package businessobjects

import (
	"context"
	"errors"
	"net/http"
	"time"
)

type CustomModelList struct {
	Items []CustomModel `json:"value"`
}

type CustomModel struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name"`
	State       string `json:"state"`
	Description string `json:"description"`
}

type EntityTypeList struct {
	Items []EntityType `json:"value"`
}

type EntityType struct {
	Id         string     `json:"id,omitempty"`
	Name       string     `json:"name"`
	PluralName string     `json:"pluralName"`
	State      string     `json:"state"`
	Key        EntityKey  `json:"key"`
	Properties []Property `json:"properties,omitempty"`
}

type PropertyList struct {
	Items []Property `json:"value"`
}

type Property struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name"`
	Required bool   `json:"required"`
	Indexed  bool   `json:"indexed"`
	Type     string `json:"type"`
	State    string `json:"state"`
}

type EntityKey struct {
	Id    string `json:"id,omitempty"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	State string `json:"state"`
}

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
