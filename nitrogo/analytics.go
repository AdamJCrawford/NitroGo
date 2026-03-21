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
	analyticsProfileURL                       = "/nitro/v1/config/analyticsprofile"
	analyticsGlobalAnalyticsProfileBindingURL = "/nitro/v1/config/analyticsglobal_analyticsprofile_binding"
	analyticsGlobalBindingURL                 = "/nitro/v1/config/analyticsglobal_binding"
)

// Analytics configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/analytics/analytics
type AnalyticsService struct {
	client *Client
}

// analyticsglobal_analyticsprofile_binding
// Binding object showing the analyticsprofile that can be bound to analyticsglobal.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/analytics/analyticsglobal_analyticsprofile_binding

func (s *AnalyticsService) AddAnalyticsGlobalAnalyticsProfileBinding(binding models.AnalyticsGlobalAnalyticsProfileBinding) error {
	payload := map[string]any{
		"analyticsglobal_analyticsprofile_binding": binding,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, analyticsGlobalAnalyticsProfileBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AnalyticsService) DeleteAnalyticsGlobalAnalyticsProfileBinding(analyticsprofile string) error {
	reqURL := fmt.Sprintf("%s?args=analyticsprofile:%s", analyticsGlobalAnalyticsProfileBindingURL, url.QueryEscape(analyticsprofile))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AnalyticsService) GetAnalyticsGlobalAnalyticsProfileBinding() ([]models.AnalyticsGlobalAnalyticsProfileBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, analyticsGlobalAnalyticsProfileBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AnalyticsGlobalAnalyticsProfileBinding `json:"analyticsglobal_analyticsprofile_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AnalyticsService) CountAnalyticsGlobalAnalyticsProfileBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, analyticsGlobalAnalyticsProfileBindingURL+"?count=yes", nil)
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
		} `json:"analyticsglobal_analyticsprofile_binding"`
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

// analyticsglobal_binding
// Binding object which returns the resources bound to analyticsglobal.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/analytics/analyticsglobal_binding

func (s *AnalyticsService) AnalyticsGlobalBinding() (models.AnalyticsGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, analyticsGlobalBindingURL, nil)
	if err != nil {
		return models.AnalyticsGlobalBinding{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.AnalyticsGlobalBinding{}, err
	}

	var result struct {
		Bindings []models.AnalyticsGlobalBinding `json:"analyticsglobal_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.AnalyticsGlobalBinding{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0], nil
	}

	return models.AnalyticsGlobalBinding{}, nil
}

// analyticsprofile
// Configuration for Analytics profile resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/analytics/analyticsprofile

func (s *AnalyticsService) AddAnalyticsProfile(profile models.AnalyticsProfile) error {
	payload := map[string]any{
		"analyticsprofile": profile,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, analyticsProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AnalyticsService) UpdateAnalyticsProfile(profile models.AnalyticsProfile) error {
	payload := map[string]any{
		"analyticsprofile": profile,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, analyticsProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AnalyticsService) UnsetAnalyticsProfile(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"analyticsprofile": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, analyticsProfileURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AnalyticsService) DeleteAnalyticsProfile(name string) error {
	url := fmt.Sprintf("%s/%s", analyticsProfileURL, name)
	req, err := s.client.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AnalyticsService) GetAllAnalyticsProfile() ([]models.AnalyticsProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, analyticsProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		AnalyticsProfiles []models.AnalyticsProfile `json:"analyticsprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.AnalyticsProfiles, nil
}

func (s *AnalyticsService) GetAnalyticsProfile(name string) ([]models.AnalyticsProfile, error) {
	url := fmt.Sprintf("%s/%s", analyticsProfileURL, name)
	req, err := s.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		AnalyticsProfiles []models.AnalyticsProfile `json:"analyticsprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.AnalyticsProfiles, nil
}

func (s *AnalyticsService) CheckAnalyticsProfile() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, analyticsProfileURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		AnalyticsProfiles []struct {
			Count float64 `json:"__count"`
		} `json:"analyticsprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.AnalyticsProfiles) > 0 {
		return result.AnalyticsProfiles[0].Count, nil
	}

	return 0, nil
}

func (s *AnalyticsService) ChangeAnalyticsProfile(profile models.AnalyticsProfile) error {
	payload := map[string]any{
		"analyticsprofile": profile,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, analyticsProfileURL+"?action=update", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}
