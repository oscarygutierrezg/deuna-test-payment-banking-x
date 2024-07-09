package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"payment-banking-x/internal/config"
)

type RestClient interface {
	Post(endpoint string, req interface{}, res interface{}) (interface{}, error)
}

func NewRestClient(cfg *config.Config) *restClient {
	return &restClient{
		baseUrl:     cfg.BaseUrl,
		contentType: cfg.ContentType,
	}
}

type restClient struct {
	baseUrl     string
	contentType string
}

func (s *restClient) Post(endpoint string, req interface{}, res interface{}) (interface{}, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error encoding request: %v", err)
	}

	resp, err := http.Post(fmt.Sprintf("%s/%s", s.baseUrl, endpoint), s.contentType, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("error calling endpoint: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errorDTO ErrorDTO
		err = json.NewDecoder(resp.Body).Decode(&errorDTO)
		if err != nil {
			return nil, fmt.Errorf("error calling endpoint: %v", err)
		}
		return errorDTO, fmt.Errorf("error calling endpoint: %v", err)
	}

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}
	return res, nil
}
