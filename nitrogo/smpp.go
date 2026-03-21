package nitrogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AdamJCrawford/NitroGo/nitrogo/models"
)

const (
	smppParamURL = "/nitro/v1/config/smppparam"
	smppUserURL  = "/nitro/v1/config/smppuser"
)

// All the commands associated with SMPP.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/smpp/smpp
type SMPPService struct {
	client *Client
}

// smppparam
// Configuration for SMPP configuration parameters resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/smpp/smppparam

func (s *SMPPService) UpdateSMPPParam(param models.SMPPParam) error {
	payload := map[string]any{
		"smppparam": param,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, smppParamURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SMPPService) UnsetSMPPParam(args []string) error {
	unsetMap := make(map[string]bool)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"smppparam": unsetMap,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, smppParamURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SMPPService) GetAllSMPPParam() (models.SMPPParam, error) {
	req, err := s.client.NewRequest(http.MethodGet, smppParamURL, nil)
	if err != nil {
		return models.SMPPParam{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.SMPPParam{}, err
	}

	var result struct {
		SMPPParams []models.SMPPParam `json:"smppparam"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.SMPPParam{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.SMPPParams) > 0 {
		return result.SMPPParams[0], nil
	}

	return models.SMPPParam{}, nil
}

// smppuser
// Configuration for SMPP user resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/smpp/smppuser

func (s *SMPPService) AddSMPPUser(user models.SMPPUser) error {
	payload := map[string]any{
		"smppuser": user,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, smppUserURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SMPPService) DeleteSMPPUser(username string) error {
	url := fmt.Sprintf("%s/%s", smppUserURL, username)
	req, err := s.client.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SMPPService) UpdateSMPPUser(user models.SMPPUser) error {
	payload := map[string]any{
		"smppuser": user,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, smppUserURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SMPPService) GetAllSMPPUser() ([]models.SMPPUser, error) {
	req, err := s.client.NewRequest(http.MethodGet, smppUserURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		SMPPUsers []models.SMPPUser `json:"smppuser"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.SMPPUsers, nil
}

func (s *SMPPService) GetSMPPUser(username string) ([]models.SMPPUser, error) {
	url := fmt.Sprintf("%s/%s", smppUserURL, username)
	req, err := s.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		SMPPUsers []models.SMPPUser `json:"smppuser"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.SMPPUsers, nil
}

func (s *SMPPService) CountSMPPUser() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, smppUserURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		SMPPUsers []struct {
			Count float64 `json:"__count"`
		} `json:"smppuser"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.SMPPUsers) > 0 {
		return result.SMPPUsers[0].Count, nil
	}

	return 0, nil
}
