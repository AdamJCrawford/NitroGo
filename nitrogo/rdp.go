package nitrogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AdamJCrawford/NitroGo/nitrogo/models"
)

const (
	rdpClientProfileURL = "/nitro/v1/config/rdpclientprofile"
	rdpConnectionsURL   = "/nitro/v1/config/rdpconnections"
	rdpServerProfileURL = "/nitro/v1/config/rdpserverprofile"
)

// RDP configuration
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/rdp/rdp
type RDPService struct {
	client *Client
}

// rdpclientprofile
// Configuration for RDP clientprofile resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/rdp/rdpclientprofile

func (s *RDPService) AddRDPClientProfile(profile models.RDPClientProfile) error {
	payload := map[string]any{
		"rdpclientprofile": profile,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, rdpClientProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *RDPService) DeleteRDPClientProfile(name string) error {
	url := fmt.Sprintf("%s/%s", rdpClientProfileURL, name)
	req, err := s.client.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *RDPService) UpdateRDPClientProfile(profile models.RDPClientProfile) error {
	payload := map[string]any{
		"rdpclientprofile": profile,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, rdpClientProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *RDPService) UnsetRDPClientProfile(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"rdpclientprofile": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, rdpClientProfileURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *RDPService) GetAllRDPClientProfile() ([]models.RDPClientProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, rdpClientProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		RDPClientProfiles []models.RDPClientProfile `json:"rdpclientprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.RDPClientProfiles, nil
}

func (s *RDPService) GetRDPClientProfile(name string) ([]models.RDPClientProfile, error) {
	url := fmt.Sprintf("%s/%s", rdpClientProfileURL, name)
	req, err := s.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		RDPClientProfiles []models.RDPClientProfile `json:"rdpclientprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.RDPClientProfiles, nil
}

func (s *RDPService) CountRDPClientProfile() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, rdpClientProfileURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		RDPClientProfiles []struct {
			Count float64 `json:"__count"`
		} `json:"rdpclientprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.RDPClientProfiles) > 0 {
		return result.RDPClientProfiles[0].Count, nil
	}

	return 0, nil
}

// rdpconnections
// Configuration for active rdp connections resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/rdp/rdpconnections

func (s *RDPService) GetAllRDPConnections() ([]models.RDPConnections, error) {
	req, err := s.client.NewRequest(http.MethodGet, rdpConnectionsURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		RDPConnections []models.RDPConnections `json:"rdpconnections"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.RDPConnections, nil
}

func (s *RDPService) CountRDPConnections() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, rdpConnectionsURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		RDPConnections []struct {
			Count float64 `json:"__count"`
		} `json:"rdpconnections"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.RDPConnections) > 0 {
		return result.RDPConnections[0].Count, nil
	}

	return 0, nil
}

func (s *RDPService) KillRDPConnections(connections models.RDPConnections) error {
	payload := map[string]any{
		"rdpconnections": connections,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, rdpConnectionsURL+"?action=kill", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// rspserverprofile
// Configuration for RDP serverprofile resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/rdp/rdpserverprofile

func (s *RDPService) AddRDPServerProfile(profile models.RDPServerProfile) error {
	payload := map[string]any{
		"rdpserverprofile": profile,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, rdpServerProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *RDPService) DeleteRDPServerProfile(name string) error {
	url := fmt.Sprintf("%s/%s", rdpServerProfileURL, name)
	req, err := s.client.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *RDPService) UpdateRDPServerProfile(profile models.RDPServerProfile) error {
	payload := map[string]any{
		"rdpserverprofile": profile,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, rdpServerProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *RDPService) UnsetRDPServerProfile(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"rdpserverprofile": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, rdpServerProfileURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *RDPService) GetAllRDPServerProfile() ([]models.RDPServerProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, rdpServerProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		RDPServerProfiles []models.RDPServerProfile `json:"rdpserverprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.RDPServerProfiles, nil
}

func (s *RDPService) GetRDPServerProfile(name string) ([]models.RDPServerProfile, error) {
	url := fmt.Sprintf("%s/%s", rdpServerProfileURL, name)
	req, err := s.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		RDPServerProfiles []models.RDPServerProfile `json:"rdpserverprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.RDPServerProfiles, nil
}

func (s *RDPService) CountRDPServerProfile() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, rdpServerProfileURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		RDPServerProfiles []struct {
			Count float64 `json:"__count"`
		} `json:"rdpserverprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.RDPServerProfiles) > 0 {
		return result.RDPServerProfiles[0].Count, nil
	}

	return 0, nil
}
