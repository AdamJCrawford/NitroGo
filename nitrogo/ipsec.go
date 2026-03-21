package nitrogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AdamJCrawford/NitroGo/nitrogo/models"
)

const (
	ipsecParameterURL = "/nitro/v1/config/ipsecparameter"
	ipsecProfileURL   = "/nitro/v1/config/ipsecprofile"
)

// IPSEC configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/ipsec/ipsec
type IPSECService struct {
	client *Client
}

// ipsecparameter
// Configuration for IPSEC paramter resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/ipsec/ipsecparameter

func (s *IPSECService) UpdateIPSECParamter(param models.IPSECParameter) error {
	payload := map[string]any{
		"ipsecparameter": param,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, ipsecParameterURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *IPSECService) UnsetIPSECParamter(args []string) error {
	unsetMap := make(map[string]bool)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"ipsecparameter": unsetMap,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, ipsecParameterURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *IPSECService) GetAllIPSECParamter() (models.IPSECParameter, error) {
	req, err := s.client.NewRequest(http.MethodGet, ipsecParameterURL, nil)
	if err != nil {
		return models.IPSECParameter{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.IPSECParameter{}, err
	}

	var result struct {
		IPSECParameters []models.IPSECParameter `json:"ipsecparameter"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.IPSECParameter{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.IPSECParameters) > 0 {
		return result.IPSECParameters[0], nil
	}

	return models.IPSECParameter{}, nil
}

// ipsecprofile
// Configuration for IPSEC profile resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/ipsec/ipsecprofile

func (s *IPSECService) AddIPSECProfile(profile models.IPSECProfile) error {
	payload := map[string]any{
		"ipsecprofile": profile,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, ipsecProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *IPSECService) DeleteIPSECProfile(name string) error {
	url := fmt.Sprintf("%s/%s", ipsecProfileURL, name)
	req, err := s.client.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *IPSECService) GetAllIPSECProfile() ([]models.IPSECProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, ipsecProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		IPSECProfiles []models.IPSECProfile `json:"ipsecprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.IPSECProfiles, nil
}

func (s *IPSECService) GetIPSECProfile(name string) ([]models.IPSECProfile, error) {
	url := fmt.Sprintf("%s/%s", ipsecProfileURL, name)
	req, err := s.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		IPSECProfiles []models.IPSECProfile `json:"ipsecprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.IPSECProfiles, nil
}

func (s *IPSECService) CountIPSECProfile() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, ipsecProfileURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		IPSECProfiles []struct {
			Count float64 `json:"__count"`
		} `json:"ipsecprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.IPSECProfiles) > 0 {
		return result.IPSECProfiles[0].Count, nil
	}

	return 0, nil
}
