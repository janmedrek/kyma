package testkit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type RegistryClient struct {
	url        string
	httpClient *http.Client
}

func NewRegistryClient(url string, httpClient *http.Client) *RegistryClient {
	return &RegistryClient{
		url:        url,
		httpClient: httpClient,
	}
}

func (rc *RegistryClient) RegisterService(service *ServiceDetails) (string, error) {
	body, err := json.Marshal(service)
	if err != nil {
		return "", err
	}

	request, err := http.NewRequest(http.MethodPost, rc.url, bytes.NewReader(body))
	if err != nil {
		return "", err
	}

	request.Header.Add("Content-Type", "application/json")

	response, err := rc.httpClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		err := parseErrorResponse(response)
		return "", err
	}

	registerServiceResponse := &RegisterServiceResponse{}

	err = json.NewDecoder(response.Body).Decode(registerServiceResponse)
	if err != nil {
		return "", err
	}

	return registerServiceResponse.ID, nil
}

func (rc *RegistryClient) GetService(id string) (*ServiceDetails, error) {
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/%s", rc.url, id), nil)
	if err != nil {
		return nil, err
	}

	response, err := rc.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		err := parseErrorResponse(response)
		return nil, err
	}

	serviceDetails := &ServiceDetails{}

	err = json.NewDecoder(response.Body).Decode(serviceDetails)
	if err != nil {
		return nil, err
	}

	return serviceDetails, nil
}

func (rc *RegistryClient) DeleteService(id string) error {
	request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", rc.url, id), nil)
	if err != nil {
		return err
	}

	response, err := rc.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusNoContent {
		err := parseErrorResponse(response)
		return err
	}

	return nil
}
