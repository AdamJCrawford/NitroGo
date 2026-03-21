package nitrogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AdamJCrawford/NitroGo/nitrogo/models"
)

const (
	ulfdServerURL = "/nitro/v1/config/ulfdserver"
)

// ULFD server configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/ulfd/ulfd
type ULFDService struct {
	client *Client
}

// ulfdserver
// Configuration for ulfd server resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/ulfd/ulfdserver

func (s *ULFDService) AddULFDServer(server models.ULFDServer) error {
	payload := map[string]any{
		"ulfdserver": server,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, ulfdServerURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ULFDService) DeleteULFDServer(loggerip string) error {
	url := fmt.Sprintf("%s/%s", ulfdServerURL, loggerip)
	req, err := s.client.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ULFDService) GetAllULFDServer() ([]models.ULFDServer, error) {
	req, err := s.client.NewRequest(http.MethodGet, ulfdServerURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		ULFDServers []models.ULFDServer `json:"ulfdserver"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.ULFDServers, nil
}

func (s *ULFDService) GetULFDServer(loggerip string) ([]models.ULFDServer, error) {
	url := fmt.Sprintf("%s/%s", ulfdServerURL, loggerip)
	req, err := s.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		ULFDServers []models.ULFDServer `json:"ulfdserver"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.ULFDServers, nil
}

func (s *ULFDService) CountULFDServer() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, ulfdServerURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		ULFDServers []struct {
			Count float64 `json:"__count"`
		} `json:"ulfdserver"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.ULFDServers) > 0 {
		return result.ULFDServers[0].Count, nil
	}

	return 0, nil
}
