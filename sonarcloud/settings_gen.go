package sonarcloud

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/form/v4"
	"github.com/m-yosefpor/go-sonarcloud/sonarcloud/settings"
	"strings"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Settings service

func (s *Settings) ListDefinitions(r settings.ListDefinitionsRequest) (*settings.ListDefinitionsResponse, error) {
	params := paramsFrom(r)

	req, err := s.client.GetRequest(fmt.Sprintf("%s/settings/list_definitions", API), params...)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %+v", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error trying to execute request: %+v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		if errorResponse, err := ErrorResponseFrom(resp); err != nil {
			return nil, fmt.Errorf("received non 2xx status code (%d), but could not decode error response: %+v", resp.StatusCode, err)
		} else {
			return nil, errorResponse
		}
	}

	response := &settings.ListDefinitionsResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}

func (s *Settings) Reset(r settings.ResetRequest) error {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.PostRequest(fmt.Sprintf("%s/settings/reset", API), strings.NewReader(values.Encode()))
	if err != nil {
		return fmt.Errorf("could not create request: %+v", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("error trying to execute request: %+v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		if errorResponse, err := ErrorResponseFrom(resp); err != nil {
			return fmt.Errorf("received non 2xx status code (%d), but could not decode error response: %+v", resp.StatusCode, err)
		} else {
			return errorResponse
		}
	}

	return nil
}

func (s *Settings) Set(r settings.SetRequest) error {
	encoder := form.NewEncoder()
	values, err := encoder.Encode(r)
	if err != nil {
		return fmt.Errorf("could not encode form values: %+v", err)
	}

	req, err := s.client.PostRequest(fmt.Sprintf("%s/settings/set", API), strings.NewReader(values.Encode()))
	if err != nil {
		return fmt.Errorf("could not create request: %+v", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("error trying to execute request: %+v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		if errorResponse, err := ErrorResponseFrom(resp); err != nil {
			return fmt.Errorf("received non 2xx status code (%d), but could not decode error response: %+v", resp.StatusCode, err)
		} else {
			return errorResponse
		}
	}

	return nil
}

func (s *Settings) Values(r settings.ValuesRequest) (*settings.ValuesResponse, error) {
	params := paramsFrom(r)

	req, err := s.client.GetRequest(fmt.Sprintf("%s/settings/values", API), params...)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %+v", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error trying to execute request: %+v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		if errorResponse, err := ErrorResponseFrom(resp); err != nil {
			return nil, fmt.Errorf("received non 2xx status code (%d), but could not decode error response: %+v", resp.StatusCode, err)
		} else {
			return nil, errorResponse
		}
	}

	response := &settings.ValuesResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}
