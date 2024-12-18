package sonarcloud

import (
	"encoding/json"
	"fmt"
	"github.com/m-yosefpor/go-sonarcloud/sonarcloud/timemachine"
)

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

type Timemachine service

func (s *Timemachine) Index(r timemachine.IndexRequest) (*timemachine.IndexResponse, error) {
	params := paramsFrom(r)

	req, err := s.client.GetRequest(fmt.Sprintf("%s/timemachine/index", API), params...)
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

	response := &timemachine.IndexResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %+v", err)
	}
	return response, nil
}
