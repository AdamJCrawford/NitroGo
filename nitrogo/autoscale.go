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
	autoscaleActionURL               = "/nitro/v1/config/autoscaleaction"
	autoscalePolicyURL               = "/nitro/v1/config/autoscalepolicy"
	autoscalePolicyBindingURL        = "/nitro/v1/config/autoscalepolicy_binding"
	autoscalePolicyNSTimerBindingURL = "/nitro/v1/config/autoscalepolicy_nstimer_binding"
	autoscaleProfileURL              = "/nitro/v1/config/autoscaleprofile"
)

// AutoScale configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/autoscale/autoscale
type AutoscaleService struct {
	client *Client
}

// autoscaleaction
// Configuration for autoscale action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/autoscale/autoscaleaction

func (s *AutoscaleService) AddAutoscaleAction(action models.AutoscaleAction) error {
	payload := map[string]any{
		"autoscaleaction": action,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, autoscaleActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AutoscaleService) DeleteAutoscaleAction(name string) error {
	urlReq := fmt.Sprintf("%s/%s", autoscaleActionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, urlReq, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AutoscaleService) UpdateAutoscaleAction(action models.AutoscaleAction) error {
	payload := map[string]any{
		"autoscaleaction": action,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, autoscaleActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AutoscaleService) UnsetAutoscaleAction(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"autoscaleaction": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, autoscaleActionURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AutoscaleService) GetAllAutoscaleAction() ([]models.AutoscaleAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, autoscaleActionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Actions []models.AutoscaleAction `json:"autoscaleaction"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Actions, nil
}

func (s *AutoscaleService) GetAutoscaleAction(name string) ([]models.AutoscaleAction, error) {
	urlReq := fmt.Sprintf("%s/%s", autoscaleActionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Actions []models.AutoscaleAction `json:"autoscaleaction"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Actions, nil
}

func (s *AutoscaleService) CountAutoscaleAction() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, autoscaleActionURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Actions []struct {
			Count float64 `json:"__count"`
		} `json:"autoscaleaction"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Actions) > 0 {
		return result.Actions[0].Count, nil
	}

	return 0, nil
}

// autoscalepolicy
// Configuration for Autoscale policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/autoscale/autoscalepolicy

func (s *AutoscaleService) AddAutoscalePolicy(policy models.AutoscalePolicy) error {
	payload := map[string]any{
		"autoscalepolicy": policy,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, autoscalePolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AutoscaleService) DeleteAutoscalePolicy(name string) error {
	urlReq := fmt.Sprintf("%s/%s", autoscalePolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, urlReq, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AutoscaleService) UpdateAutoscalePolicy(policy models.AutoscalePolicy) error {
	payload := map[string]any{
		"autoscalepolicy": policy,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, autoscalePolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AutoscaleService) UnsetAutoscalePolicy(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"autoscalepolicy": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, autoscalePolicyURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AutoscaleService) GetAllAutoscalePolicy() ([]models.AutoscalePolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, autoscalePolicyURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Policies []models.AutoscalePolicy `json:"autoscalepolicy"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Policies, nil
}

func (s *AutoscaleService) GetAutoscalePolicy(name string) ([]models.AutoscalePolicy, error) {
	urlReq := fmt.Sprintf("%s/%s", autoscalePolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Policies []models.AutoscalePolicy `json:"autoscalepolicy"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Policies, nil
}

func (s *AutoscaleService) CountAutoscalePolicy() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, autoscalePolicyURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Policies []struct {
			Count float64 `json:"__count"`
		} `json:"autoscalepolicy"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Policies) > 0 {
		return result.Policies[0].Count, nil
	}

	return 0, nil
}

func (s *AutoscaleService) RenameAutoscalePolicy(name string, newname string) error {
	payload := map[string]any{
		"autoscalepolicy": map[string]string{
			"name":    name,
			"newname": newname,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, autoscalePolicyURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// autoscalepolicy_binding
// Binding object which returns the resources bound to autoscalepolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/autoscale/autoscalepolicy_binding

func (s *AutoscaleService) GetAllAutoscalePolicyBinding() ([]models.AutoscalePolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, autoscalePolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AutoscalePolicyBinding `json:"autoscalepolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AutoscaleService) GetAutoscalePolicyBinding(name string) ([]models.AutoscalePolicyBinding, error) {
	urlReq := fmt.Sprintf("%s/%s", autoscalePolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AutoscalePolicyBinding `json:"autoscalepolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

// autoscalepolicy_nstimer_binding
// Binding object showing the nstimer that can be bound to autoscalepolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/autoscale/autoscalepolicy_nstimer_binding

func (s *AutoscaleService) GetAllAutoscalePolicyNSTimerBinding() ([]models.AutoscalePolicyNSTimerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, autoscalePolicyNSTimerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AutoscalePolicyNSTimerBinding `json:"autoscalepolicy_nstimer_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AutoscaleService) GetAutoscalePolicyNSTimerBinding(name string) ([]models.AutoscalePolicyNSTimerBinding, error) {
	urlReq := fmt.Sprintf("%s/%s", autoscalePolicyNSTimerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AutoscalePolicyNSTimerBinding `json:"autoscalepolicy_nstimer_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AutoscaleService) CountAutoscalePolicyNSTimerBinding(name string) (float64, error) {
	urlReq := fmt.Sprintf("%s/%s?count=yes", autoscalePolicyNSTimerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
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
		} `json:"autoscalepolicy_nstimer_binding"`
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

// autoscaleprofile
// Configuration for autoscale profile resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/autoscale/autoscaleprofile

func (s *AutoscaleService) AddAutoscaleProfile(profile models.AutoscaleProfile) error {
	payload := map[string]any{
		"autoscaleprofile": profile,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, autoscaleProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AutoscaleService) DeleteAutoscaleProfile(name string) error {
	urlReq := fmt.Sprintf("%s/%s", autoscaleProfileURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, urlReq, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AutoscaleService) UpdateAutoscaleProfile(profile models.AutoscaleProfile) error {
	payload := map[string]any{
		"autoscaleprofile": profile,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, autoscaleProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AutoscaleService) GetAllAutoscaleProfile() ([]models.AutoscaleProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, autoscaleProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Profiles []models.AutoscaleProfile `json:"autoscaleprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Profiles, nil
}

func (s *AutoscaleService) GetAutoscaleProfile(name string) ([]models.AutoscaleProfile, error) {
	urlReq := fmt.Sprintf("%s/%s", autoscaleProfileURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Profiles []models.AutoscaleProfile `json:"autoscaleprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Profiles, nil
}

func (s *AutoscaleService) CountAutoscaleProfile() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, autoscaleProfileURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Profiles []struct {
			Count float64 `json:"__count"`
		} `json:"autoscaleprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Profiles) > 0 {
		return result.Profiles[0].Count, nil
	}

	return 0, nil
}
