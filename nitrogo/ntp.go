package nitrogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/AdamJCrawford/NitroGo/nitrogo/models"
)

const (
	ntpParamURL  = "/nitro/v1/config/ntpparam"
	ntpServerURL = "/nitro/v1/config/ntpserver"
	ntpStatusURL = "/nitro/v1/config/ntpstatus"
	ntpSyncURL   = "/nitro/v1/config/ntpsync"
)

// NTP configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/ntp/ntp
type NTPService struct {
	client *Client
}

// ntpparam
// Configuration for NTP parameter resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/ntp/ntpparam

func (s *NTPService) UpdateNTPParam(param models.NTPParam) error {
	payload := map[string]any{
		"ntpparam": param,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, ntpParamURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NTPService) UnsetNTPParam(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"ntpparam": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, ntpParamURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NTPService) GetAllNTPParam() (models.NTPParam, error) {
	req, err := s.client.NewRequest(http.MethodGet, ntpParamURL, nil)
	if err != nil {
		return models.NTPParam{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.NTPParam{}, err
	}

	var result struct {
		NTPParams []models.NTPParam `json:"ntpparam"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.NTPParam{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.NTPParams) > 0 {
		return result.NTPParams[0], nil
	}

	return models.NTPParam{}, nil
}

// ntpserver
// Configuration for NTP server resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/ntp/ntpserver

func (s *NTPService) AddNTPServer(server models.NTPServer) error {
	payload := map[string]any{
		"ntpserver": server,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, ntpServerURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NTPService) DeleteNTPServer(serverip string, servername string) error {
	reqURL := fmt.Sprintf("%s/%s", ntpServerURL, serverip)
	if servername != "" {
		reqURL += "?args=servername:" + url.QueryEscape(servername)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NTPService) UpdateNTPServer(server models.NTPServer) error {
	payload := map[string]any{
		"ntpserver": server,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, ntpServerURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NTPService) UnsetNTPServer(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"ntpserver": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, ntpServerURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NTPService) GetAllNTPServer() ([]models.NTPServer, error) {
	req, err := s.client.NewRequest(http.MethodGet, ntpServerURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		NTPServers []models.NTPServer `json:"ntpserver"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.NTPServers, nil
}

func (s *NTPService) CountNTPServer() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, ntpServerURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		NTPServers []struct {
			Count float64 `json:"__count"`
		} `json:"ntpserver"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.NTPServers) > 0 {
		return result.NTPServers[0].Count, nil
	}

	return 0, nil
}

// ntpstatus
// Configuration for ntp status resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/ntp/ntpstatus

func (s *NTPService) GetAllNTPStatus() ([]models.NTPStatus, error) {
	req, err := s.client.NewRequest(http.MethodGet, ntpStatusURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		NTPStatus []models.NTPStatus `json:"ntpstatus"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.NTPStatus, nil
}

// ntpsync
// Configuration for NTP sync resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/ntp/ntpsync

func (s *NTPService) EnableNTPSync() error {
	payload := map[string]any{
		"ntpsync": map[string]any{},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, ntpSyncURL+"?action=enable", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NTPService) DisableNTPSync() error {
	payload := map[string]any{
		"ntpsync": map[string]any{},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, ntpSyncURL+"?action=disable", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NTPService) GetAllNTPSync() ([]models.NTPSync, error) {
	req, err := s.client.NewRequest(http.MethodGet, ntpSyncURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		NTPSync []models.NTPSync `json:"ntpsync"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.NTPSync, nil
}
