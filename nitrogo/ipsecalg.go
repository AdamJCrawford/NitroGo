package nitrogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AdamJCrawford/NitroGo/nitrogo/models"
)

const (
	ipsecALGProfileURL = "/nitro/v1/config/ipsecalgprofile"
	ipsecALGSessionURL = "/nitro/v1/config/ipsecalgsession"
)

// IPSECALG configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/ipsecalg/ipsecalg
type IPSECALGService struct {
	client *Client
}

// ipsecalgprofile
// Configuration for IPSEC ALG profile resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/ipsecalg/ipsecalgprofile

func (s *IPSECALGService) AddIPSECALGProfile(profile models.IPSECALGProfile) error {
	payload := map[string]any{
		"ipsecalgprofile": profile,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, ipsecALGProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *IPSECALGService) DeleteIPSECALGProfile(name string) error {
	url := fmt.Sprintf("%s/%s", ipsecALGProfileURL, name)
	req, err := s.client.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *IPSECALGService) UpdateIPSECALGProfile(profile models.IPSECALGProfile) error {
	payload := map[string]any{
		"ipsecalgprofile": profile,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, ipsecALGProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *IPSECALGService) UnsetIPSECALGProfile(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"ipsecalgprofile": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, ipsecALGProfileURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *IPSECALGService) GetAllIPSECALGProfile() ([]models.IPSECALGProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, ipsecALGProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		IPSECALGProfiles []models.IPSECALGProfile `json:"ipsecalgprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.IPSECALGProfiles, nil
}

func (s *IPSECALGService) GetIPSECALGProfile(name string) ([]models.IPSECALGProfile, error) {
	url := fmt.Sprintf("%s/%s", ipsecALGProfileURL, name)
	req, err := s.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		IPSECALGProfiles []models.IPSECALGProfile `json:"ipsecalgprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.IPSECALGProfiles, nil
}

func (s *IPSECALGService) CountIPSECALGProfile() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, ipsecALGProfileURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		IPSECALGProfiles []struct {
			Count float64 `json:"__count"`
		} `json:"ipsecalgprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.IPSECALGProfiles) > 0 {
		return result.IPSECALGProfiles[0].Count, nil
	}

	return 0, nil
}

// ipsecalgsession
// Configuration for IPSEC ALG session resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/ipsecalg/ipsecalgsession

func (s *IPSECALGService) GetAllIPSECALGSession() ([]models.IPSECALGSession, error) {
	req, err := s.client.NewRequest(http.MethodGet, ipsecALGSessionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		IPSECALGSessions []models.IPSECALGSession `json:"ipsecalgsession"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.IPSECALGSessions, nil
}

func (s *IPSECALGService) CountIPSECALGSession() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, ipsecALGSessionURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		IPSECALGSessions []struct {
			Count float64 `json:"__count"`
		} `json:"ipsecalgsession"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.IPSECALGSessions) > 0 {
		return result.IPSECALGSessions[0].Count, nil
	}

	return 0, nil
}

func (s *IPSECALGService) FlushIPSECALGSession(session models.IPSECALGSession) error {
	payload := map[string]any{
		"ipsecalgsession": session,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, ipsecALGSessionURL+"?action=flush", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}
