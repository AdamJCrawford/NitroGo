package nitrogo

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AdamJCrawford/NitroGo/nitrogo/models"
)

const (
	bfdSessionURL = "/nitro/v1/config/bfdsession"
)

// Bidirectional Forwarding Detection
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/bfd/bfd
type BFDService struct {
	client *Client
}

// bfdsession
// Configuration for BFD configuration resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/bfd/bfdsession

func (s *BFDService) GetAllBFDSession() ([]models.BFDSession, error) {
	req, err := s.client.NewRequest(http.MethodGet, bfdSessionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		BFDSessions []models.BFDSession `json:"bfdsession"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.BFDSessions, nil
}

func (s *BFDService) CountBFDSession() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, bfdSessionURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		BFDSessions []struct {
			Count float64 `json:"__count"`
		} `json:"bfdsession"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.BFDSessions) > 0 {
		return result.BFDSessions[0].Count, nil
	}

	return 0, nil
}
