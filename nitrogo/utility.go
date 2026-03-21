package nitrogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AdamJCrawford/NitroGo/nitrogo/models"
)

const (
	callHomeURL    = "/nitro/v1/config/callhome"
	installURL     = "/nitro/v1/config/install"
	pingURL        = "/nitro/v1/config/ping"
	ping6URL       = "/nitro/v1/config/ping6"
	raidURL        = "/nitro/v1/config/raid"
	techSupportURL = "/nitro/v1/config/techsupport"
	traceRouteURL  = "/nitro/v1/config/traceroute"
	traceRoute6URL = "/nitro/v1/config/traceroute6"
)

// Utilities.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/utility/utility
type UtilityService struct {
	client *Client
}

// callhome

func (s *UtilityService) UpdateCallHome(callhome models.CallHome) error {
	payload := map[string]any{
		"callhome": callhome,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, callHomeURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *UtilityService) UnsetCallHome(args []string) error {
	unsetMap := make(map[string]bool)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"callhome": unsetMap,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, callHomeURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *UtilityService) GetAllCallHome() (models.CallHome, error) {
	req, err := s.client.NewRequest(http.MethodGet, callHomeURL, nil)
	if err != nil {
		return models.CallHome{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.CallHome{}, err
	}

	var result struct {
		CallHomes []models.CallHome `json:"callhome"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.CallHome{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.CallHomes) > 0 {
		return result.CallHomes[0], nil
	}

	return models.CallHome{}, nil
}

// install

func (s *UtilityService) InstallInstall(install models.Install) error {
	payload := map[string]any{
		"install": install,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, installURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// ping

func (s *UtilityService) PingPing(ping models.Ping) (models.Ping, error) {
	payload := map[string]any{
		"ping": ping,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return models.Ping{}, fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, pingURL, bytes.NewBuffer(body))
	if err != nil {
		return models.Ping{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.Ping{}, err
	}

	var result struct {
		Pings []models.Ping `json:"ping"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.Ping{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Pings) > 0 {
		return result.Pings[0], nil
	}

	return models.Ping{}, nil
}

// ping6

func (s *UtilityService) Ping6Ping6(ping6 models.Ping6) (models.Ping6, error) {
	payload := map[string]any{
		"ping6": ping6,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return models.Ping6{}, fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, ping6URL, bytes.NewBuffer(body))
	if err != nil {
		return models.Ping6{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.Ping6{}, err
	}

	var result struct {
		Ping6s []models.Ping6 `json:"ping6"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.Ping6{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Ping6s) > 0 {
		return result.Ping6s[0], nil
	}

	return models.Ping6{}, nil
}

// raid

func (s *UtilityService) GetAllRAID() ([]models.Raid, error) {
	req, err := s.client.NewRequest(http.MethodGet, raidURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Raids []models.Raid `json:"raid"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Raids, nil
}

// techsupport

func (s *UtilityService) GetAllTechSupport() ([]models.TechSupport, error) {
	req, err := s.client.NewRequest(http.MethodGet, techSupportURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		TechSupports []models.TechSupport `json:"techsupport"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.TechSupports, nil
}

// traceroute

func (s *UtilityService) TraceRouteTraceRoute(traceroute models.Traceroute) (models.Traceroute, error) {
	payload := map[string]any{
		"traceroute": traceroute,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return models.Traceroute{}, fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, traceRouteURL, bytes.NewBuffer(body))
	if err != nil {
		return models.Traceroute{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.Traceroute{}, err
	}

	var result struct {
		Traceroutes []models.Traceroute `json:"traceroute"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.Traceroute{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Traceroutes) > 0 {
		return result.Traceroutes[0], nil
	}

	return models.Traceroute{}, nil
}

// traceroute6

func (s *UtilityService) TraceRoute6TraceRoute6(traceroute6 models.Traceroute6) (models.Traceroute6, error) {
	payload := map[string]any{
		"traceroute6": traceroute6,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return models.Traceroute6{}, fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, traceRoute6URL, bytes.NewBuffer(body))
	if err != nil {
		return models.Traceroute6{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.Traceroute6{}, err
	}

	var result struct {
		Traceroute6s []models.Traceroute6 `json:"traceroute6"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.Traceroute6{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Traceroute6s) > 0 {
		return result.Traceroute6s[0], nil
	}

	return models.Traceroute6{}, nil
}
