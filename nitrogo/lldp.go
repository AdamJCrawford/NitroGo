package nitrogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AdamJCrawford/NitroGo/nitrogo/models"
)

const (
	lldpNeighborsURL = "/nitro/v1/config/lldpneighbors"
	lldpParamURL     = "/nitro/v1/config/lldpparam"
)

// Link Layer Discovery Protocol.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/lldp/lldp
type LLDPService struct {
	client *Client
}

// lldpneighbors
// Configuration for lldp neighbors resource
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/lldp/lldpneighbors

func (s *LLDPService) GetAllLLDPNeighbors() ([]models.LLDPNeighbors, error) {
	req, err := s.client.NewRequest(http.MethodGet, lldpNeighborsURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		LLDPNeighbors []models.LLDPNeighbors `json:"lldpneighbors"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.LLDPNeighbors, nil
}

func (s *LLDPService) GetLLDPNeighbors(ifnum string) ([]models.LLDPNeighbors, error) {
	url := lldpNeighborsURL + "/" + ifnum
	req, err := s.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		LLDPNeighbors []models.LLDPNeighbors `json:"lldpneighbors"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.LLDPNeighbors, nil
}

func (s *LLDPService) CountLLDPNeighbors() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, lldpNeighborsURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		LLDPNeighbors []struct {
			Count float64 `json:"__count"`
		} `json:"lldpneighbors"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.LLDPNeighbors) > 0 {
		return result.LLDPNeighbors[0].Count, nil
	}

	return 0, nil
}

func (s *LLDPService) ClearLLDPNeighbors(neighbor models.LLDPNeighbors) error {
	payload := map[string]any{
		"lldpneighbors": neighbor,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, lldpNeighborsURL+"?action=clear", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// lldpparam
// Configuration for lldp params resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/lldp/lldpparam

func (s *LLDPService) UpdateLLDPParam(param models.LLDPParam) error {
	payload := map[string]any{
		"lldpparam": param,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, lldpParamURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LLDPService) UnsetLLDPParam(args []string) error {
	unsetMap := make(map[string]bool)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"lldpparam": unsetMap,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, lldpParamURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LLDPService) GetAllLLDPParam() (models.LLDPParam, error) {
	req, err := s.client.NewRequest(http.MethodGet, lldpParamURL, nil)
	if err != nil {
		return models.LLDPParam{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.LLDPParam{}, err
	}

	var result struct {
		LLDPParams []models.LLDPParam `json:"lldpparam"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.LLDPParam{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.LLDPParams) > 0 {
		return result.LLDPParams[0], nil
	}

	return models.LLDPParam{}, nil
}
