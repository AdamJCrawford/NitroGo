package nitrogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AdamJCrawford/NitroGo/nitrogo/models"
)

const (
	admParameterURL = "/nitro/v1/config/admparameter"
)

// ADM related configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/adm/adm
type ADMService struct {
	client *Client
}

// admparameter
// Configuration for ADM parameter resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/adm/admparameter

func (s *ADMService) UpdateADMParamter(param models.ADMParameter) error {
	payload := map[string]any{
		"admparameter": param,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, admParameterURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ADMService) UnsetADMParamter(args []string) error {
	unsetMap := make(map[string]bool)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"admparameter": unsetMap,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, admParameterURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ADMService) GetAllADMParamter() (models.ADMParameter, error) {
	req, err := s.client.NewRequest(http.MethodGet, admParameterURL, nil)
	if err != nil {
		return models.ADMParameter{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.ADMParameter{}, err
	}

	var result struct {
		ADMParameters []models.ADMParameter `json:"admparameter"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.ADMParameter{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.ADMParameters) > 0 {
		return result.ADMParameters[0], nil
	}

	return models.ADMParameter{}, nil
}
