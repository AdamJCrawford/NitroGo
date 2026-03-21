package nitrogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/AdamJCrawford/NitroGo/nitrogo/models"
)

const (
	streamIdentifierURL                     = "/nitro/v1/config/streamidentifier"
	streamIdentifierBindingURL              = "/nitro/v1/config/streamidentifier_binding"
	streamIdentifierStreamSessionBindingURL = "/nitro/v1/config/streamidentifier_streamsession_binding"
	streamSelectorURL                       = "/nitro/v1/config/streamselector"
	streamSessionURL                        = "/nitro/v1/config/streamsession"
)

// Stream configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/stream/stream
type StreamService struct {
	client *Client
}

// streamidentifier
// Configuration for identifier resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/stream/streamidentifier

func (s *StreamService) AddStreamIdentifier(identifier models.StreamIdentifier) error {
	payload := map[string]any{
		"streamidentifier": identifier,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, streamIdentifierURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *StreamService) DeleteStreamIdentifier(name string) error {
	url := fmt.Sprintf("%s/%s", streamIdentifierURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *StreamService) UpdateStreamIdentifier(identifier models.StreamIdentifier) error {
	payload := map[string]any{
		"streamidentifier": identifier,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, streamIdentifierURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *StreamService) UnsetStreamIdentifier(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"streamidentifier": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, streamIdentifierURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *StreamService) GetAllStreamIdentifier() ([]models.StreamIdentifier, error) {
	req, err := s.client.NewRequest(http.MethodGet, streamIdentifierURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		StreamIdentifiers []models.StreamIdentifier `json:"streamidentifier"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.StreamIdentifiers, nil
}

func (s *StreamService) GetStreamIdentifier(name string) ([]models.StreamIdentifier, error) {
	url := fmt.Sprintf("%s/%s", streamIdentifierURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		StreamIdentifiers []models.StreamIdentifier `json:"streamidentifier"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.StreamIdentifiers, nil
}

func (s *StreamService) CountStreamIdentifier() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, streamIdentifierURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		StreamIdentifiers []struct {
			Count float64 `json:"__count"`
		} `json:"streamidentifier"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.StreamIdentifiers) > 0 {
		return result.StreamIdentifiers[0].Count, nil
	}

	return 0, nil
}

// streamidentifier_binding
// Binding object which returns the resources bound to streamidentifier.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/stream/streamidentifier_binding

func (s *StreamService) GetAllStreamIdentifierBinding() ([]models.StreamIdentifierBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, streamIdentifierBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.StreamIdentifierBinding `json:"streamidentifier_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *StreamService) GetStreamIdentifierBinding(name string) ([]models.StreamIdentifierBinding, error) {
	url := fmt.Sprintf("%s/%s", streamIdentifierBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.StreamIdentifierBinding `json:"streamidentifier_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

// streamidentifier_streamsession_binding
// Binding object showing the streamsession that can be bound to streamidentifier.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/stream/streamidentifier_streamsession_binding

func (s *StreamService) GetAllStreamIdentifierStreamSessionBinding() ([]models.StreamIdentifierStreamSessionBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, streamIdentifierStreamSessionBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.StreamIdentifierStreamSessionBinding `json:"streamidentifier_streamsession_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *StreamService) GetStreamIdentifierStreamSessionBinding(name string) ([]models.StreamIdentifierStreamSessionBinding, error) {
	url := fmt.Sprintf("%s/%s", streamIdentifierStreamSessionBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.StreamIdentifierStreamSessionBinding `json:"streamidentifier_streamsession_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *StreamService) CountStreamIdentifierStreamSessionBinding(name string) (float64, error) {
	url := fmt.Sprintf("%s/%s?count=yes", streamIdentifierStreamSessionBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"streamidentifier_streamsession_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// streamselector
// Configuration for selector resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/stream/streamselector

func (s *StreamService) AddStreamSelector(selector models.StreamSelector) error {
	payload := map[string]any{
		"streamselector": selector,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, streamSelectorURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *StreamService) DeleteStreamSelector(name string) error {
	url := fmt.Sprintf("%s/%s", streamSelectorURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *StreamService) UpdateStreamSelector(selector models.StreamSelector) error {
	payload := map[string]any{
		"streamselector": selector,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, streamSelectorURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *StreamService) GetAllStreamSelector() ([]models.StreamSelector, error) {
	req, err := s.client.NewRequest(http.MethodGet, streamSelectorURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		StreamSelectors []models.StreamSelector `json:"streamselector"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.StreamSelectors, nil
}

func (s *StreamService) GetStreamSelector(name string) ([]models.StreamSelector, error) {
	url := fmt.Sprintf("%s/%s", streamSelectorURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		StreamSelectors []models.StreamSelector `json:"streamselector"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.StreamSelectors, nil
}

func (s *StreamService) CountStreamSelector() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, streamSelectorURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		StreamSelectors []struct {
			Count float64 `json:"__count"`
		} `json:"streamselector"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.StreamSelectors) > 0 {
		return result.StreamSelectors[0].Count, nil
	}

	return 0, nil
}

// streamsession
// Configuration for active connection resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/stream/streamsession

func (s *StreamService) ClearStreamSession(session models.StreamSession) error {
	payload := map[string]any{
		"streamsession": session,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, streamSessionURL+"?action=clear", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}
