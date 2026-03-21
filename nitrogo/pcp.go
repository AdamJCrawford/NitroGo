package nitrogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AdamJCrawford/NitroGo/nitrogo/models"
)

const (
	pcpMapURL     = "/nitro/v1/config/pcpmap"
	pcpProfileURL = "/nitro/v1/config/pcpprofile"
	pcpServerURL  = "/nitro/v1/config/pcpserver"
)

// Port Control Protocol Configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/pcp/pcp
type PCPService struct {
	client *Client
}

// pcpmap
// Configuration for server resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/pcp/pcpmap

func (s *PCPService) GetAllPCPMap() ([]models.PCPMap, error) {
	req, err := s.client.NewRequest(http.MethodGet, pcpMapURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		PCPMaps []models.PCPMap `json:"pcpmap"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.PCPMaps, nil
}

func (s *PCPService) CountPCPMap() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, pcpMapURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		PCPMaps []struct {
			Count float64 `json:"__count"`
		} `json:"pcpmap"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.PCPMaps) > 0 {
		return result.PCPMaps[0].Count, nil
	}

	return 0, nil
}

// pcpprofile
// Configuration for PCP Profile resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/pcp/pcpprofile

func (s *PCPService) AddPCPProfile(profile models.PCPProfile) error {
	payload := map[string]any{
		"pcpprofile": profile,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, pcpProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *PCPService) DeletePCPProfile(name string) error {
	url := fmt.Sprintf("%s/%s", pcpProfileURL, name)
	req, err := s.client.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *PCPService) UpdatePCPProfile(profile models.PCPProfile) error {
	payload := map[string]any{
		"pcpprofile": profile,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, pcpProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *PCPService) UnsetPCPProfile(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"pcpprofile": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, pcpProfileURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *PCPService) GetAllPCPProfile() ([]models.PCPProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, pcpProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		PCPProfiles []models.PCPProfile `json:"pcpprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.PCPProfiles, nil
}

func (s *PCPService) GetPCPProfile(name string) ([]models.PCPProfile, error) {
	url := fmt.Sprintf("%s/%s", pcpProfileURL, name)
	req, err := s.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		PCPProfiles []models.PCPProfile `json:"pcpprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.PCPProfiles, nil
}

func (s *PCPService) CountPCPProfile() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, pcpProfileURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		PCPProfiles []struct {
			Count float64 `json:"__count"`
		} `json:"pcpprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.PCPProfiles) > 0 {
		return result.PCPProfiles[0].Count, nil
	}

	return 0, nil
}

// pcpserver
// Configuration for server resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/pcp/pcpserver

func (s *PCPService) AddPCPServer(server models.PCPServer) error {
	payload := map[string]any{
		"pcpserver": server,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, pcpServerURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *PCPService) DeletePCPServer(name string) error {
	url := fmt.Sprintf("%s/%s", pcpServerURL, name)
	req, err := s.client.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *PCPService) UpdatePCPServer(server models.PCPServer) error {
	payload := map[string]any{
		"pcpserver": server,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, pcpServerURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *PCPService) UnsetPCPServer(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"pcpserver": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, pcpServerURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *PCPService) GetAllPCPServer() ([]models.PCPServer, error) {
	req, err := s.client.NewRequest(http.MethodGet, pcpServerURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		PCPServers []models.PCPServer `json:"pcpserver"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.PCPServers, nil
}

func (s *PCPService) GetPCPServer(name string) ([]models.PCPServer, error) {
	url := fmt.Sprintf("%s/%s", pcpServerURL, name)
	req, err := s.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		PCPServers []models.PCPServer `json:"pcpserver"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.PCPServers, nil
}

func (s *PCPService) CountPCPServer() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, pcpServerURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		PCPServers []struct {
			Count float64 `json:"__count"`
		} `json:"pcpserver"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.PCPServers) > 0 {
		return result.PCPServers[0].Count, nil
	}

	return 0, nil
}
