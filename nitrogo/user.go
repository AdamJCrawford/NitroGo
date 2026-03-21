package nitrogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AdamJCrawford/NitroGo/nitrogo/models"
)

const (
	userProtocolURL = "/nitro/v1/config/userprotocol"
	userVServerURL  = "/nitro/v1/config/uservserver"
)

// User protocol configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/user/user
type UserService struct {
	client *Client
}

// userprotocol
// Configuration for user protocol resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/user/userprotocol

func (s *UserService) AddUserProtocol(protocol models.UserProtocol) error {
	payload := map[string]any{
		"userprotocol": protocol,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, userProtocolURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *UserService) DeleteUserProtocol(name string) error {
	url := fmt.Sprintf("%s/%s", userProtocolURL, name)
	req, err := s.client.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *UserService) UpdateUserProtocol(protocol models.UserProtocol) error {
	payload := map[string]any{
		"userprotocol": protocol,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, userProtocolURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *UserService) UnsetUserProtocol(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"userprotocol": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, userProtocolURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *UserService) GetAllUserProtocol() ([]models.UserProtocol, error) {
	req, err := s.client.NewRequest(http.MethodGet, userProtocolURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		UserProtocols []models.UserProtocol `json:"userprotocol"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.UserProtocols, nil
}

func (s *UserService) GetUserProtocol(name string) ([]models.UserProtocol, error) {
	url := fmt.Sprintf("%s/%s", userProtocolURL, name)
	req, err := s.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		UserProtocols []models.UserProtocol `json:"userprotocol"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.UserProtocols, nil
}

func (s *UserService) CountUserProtocol() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, userProtocolURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		UserProtocols []struct {
			Count float64 `json:"__count"`
		} `json:"userprotocol"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.UserProtocols) > 0 {
		return result.UserProtocols[0].Count, nil
	}

	return 0, nil
}

// uservserver
// Configuration for virtual server resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/user/uservserver

func (s *UserService) AddUserVServer(vserver models.UserVServer) error {
	payload := map[string]any{
		"uservserver": vserver,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, userVServerURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *UserService) DeleteUserVServer(name string) error {
	url := fmt.Sprintf("%s/%s", userVServerURL, name)
	req, err := s.client.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *UserService) UpdateUserVServer(vserver models.UserVServer) error {
	payload := map[string]any{
		"uservserver": vserver,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, userVServerURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *UserService) UnsetUserVServer(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"uservserver": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, userVServerURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *UserService) EnableUserVServer(name string) error {
	payload := map[string]any{
		"uservserver": map[string]any{
			"name": name,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, userVServerURL+"?action=enable", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *UserService) GetAllUserVServer() ([]models.UserVServer, error) {
	req, err := s.client.NewRequest(http.MethodGet, userVServerURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		UserVServers []models.UserVServer `json:"uservserver"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.UserVServers, nil
}

func (s *UserService) GetUserVServer(name string) ([]models.UserVServer, error) {
	url := fmt.Sprintf("%s/%s", userVServerURL, name)
	req, err := s.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		UserVServers []models.UserVServer `json:"uservserver"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.UserVServers, nil
}

func (s *UserService) CountUserVServer() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, userVServerURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		UserVServers []struct {
			Count float64 `json:"__count"`
		} `json:"uservserver"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.UserVServers) > 0 {
		return result.UserVServers[0].Count, nil
	}

	return 0, nil
}
