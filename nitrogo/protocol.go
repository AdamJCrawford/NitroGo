package nitrogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AdamJCrawford/NitroGo/nitrogo/models"
)

const (
	protocolHTTPBandURL = "/nitro/v1/config/protocolhttpband"
)

// Protocol Configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/protocol/protocol
type ProtocolService struct {
	client *Client
}

// protocolhttpband
// Configuration for HTTP request/response band resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/protocol/protocolhttpband

func (s *ProtocolService) UpdateProtocolHTTPBand(band models.ProtocolHTTPBand) error {
	payload := map[string]any{
		"protocolhttpband": band,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, protocolHTTPBandURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ProtocolService) UnsetProtocolHTTPBand(args []string) error {
	unsetMap := make(map[string]bool)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"protocolhttpband": unsetMap,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, protocolHTTPBandURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ProtocolService) GetAllProtocolHTTPBand() ([]models.ProtocolHTTPBand, error) {
	req, err := s.client.NewRequest(http.MethodGet, protocolHTTPBandURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bands []models.ProtocolHTTPBand `json:"protocolhttpband"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bands, nil
}

func (s *ProtocolService) ClearProtocolHTTPBand(band models.ProtocolHTTPBand) error {
	payload := map[string]any{
		"protocolhttpband": band,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, protocolHTTPBandURL+"?action=clear", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}
