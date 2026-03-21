package nitrogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AdamJCrawford/NitroGo/nitrogo/models"
)

const (
	reputationSettingsURL = "/nitro/v1/config/reputationsettings"
)

// Reputation services configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/reputation/reputation
type ReputationService struct {
	client *Client
}

// reputationsettings
// Configuration for Reputation service settings resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/reputation/reputationsettings

func (s *ReputationService) UpdateReputationSettings(settings models.ReputationSettings) error {
	payload := map[string]any{
		"reputationsettings": settings,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, reputationSettingsURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ReputationService) UnsetReputationSettings(args []string) error {
	unsetMap := make(map[string]bool)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"reputationsettings": unsetMap,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, reputationSettingsURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ReputationService) GetAllReputationSettings() (models.ReputationSettings, error) {
	req, err := s.client.NewRequest(http.MethodGet, reputationSettingsURL, nil)
	if err != nil {
		return models.ReputationSettings{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.ReputationSettings{}, err
	}

	var result struct {
		Settings []models.ReputationSettings `json:"reputationsettings"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.ReputationSettings{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Settings) > 0 {
		return result.Settings[0], nil
	}

	return models.ReputationSettings{}, nil
}
