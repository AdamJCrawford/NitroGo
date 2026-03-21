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
	spilloverActionURL                   = "/nitro/v1/config/spilloveraction"
	spilloverPolicyURL                   = "/nitro/v1/config/spilloverpolicy"
	spilloverPolicyBindingURL            = "/nitro/v1/config/spilloverpolicy_binding"
	spilloverPolicyCSVServerBindingURL   = "/nitro/v1/config/spilloverpolicy_csvserver_binding"
	spilloverPolicyGSLBVServerBindingURL = "/nitro/v1/config/spilloverpolicy_gslbvserver_binding"
	spilloverPolicyLBVServerBindingURL   = "/nitro/v1/config/spilloverpolicy_lbvserver_binding"
)

// Spillover policies and actions.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/spillover/spillover
type SpilloverService struct {
	client *Client
}

// spilloveraction
// Configuration for Spillover action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/spillover/spilloveraction

func (s *SpilloverService) AddSpilloverAction(action models.SpilloverAction) error {
	payload := map[string]any{
		"spilloveraction": action,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, spilloverActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SpilloverService) DeleteSpilloverAction(name string) error {
	urlReq := fmt.Sprintf("%s/%s", spilloverActionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, urlReq, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SpilloverService) GetAllSpilloverAction() ([]models.SpilloverAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, spilloverActionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Actions []models.SpilloverAction `json:"spilloveraction"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Actions, nil
}

func (s *SpilloverService) GetSpilloverAction(name string) ([]models.SpilloverAction, error) {
	urlReq := fmt.Sprintf("%s/%s", spilloverActionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Actions []models.SpilloverAction `json:"spilloveraction"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Actions, nil
}

func (s *SpilloverService) CountSpilloverAction() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, spilloverActionURL+"?count=yes", nil)
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
		} `json:"spilloveraction"`
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

func (s *SpilloverService) RenameSpilloverAction(name string, newname string) error {
	payload := map[string]any{
		"spilloveraction": map[string]string{
			"name":    name,
			"newname": newname,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, spilloverActionURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// spilloverpolicy
// Configuration for Spillover policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/spillover/spilloverpolicy

func (s *SpilloverService) AddSpilloverPolicy(policy models.SpilloverPolicy) error {
	payload := map[string]any{
		"spilloverpolicy": policy,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, spilloverPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SpilloverService) DeleteSpilloverPolicy(name string) error {
	urlReq := fmt.Sprintf("%s/%s", spilloverPolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, urlReq, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SpilloverService) UpdateSpilloverPolicy(policy models.SpilloverPolicy) error {
	payload := map[string]any{
		"spilloverpolicy": policy,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, spilloverPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SpilloverService) UnsetSpilloverPolicy(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"spilloverpolicy": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, spilloverPolicyURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SpilloverService) GetAllSpilloverPolicy() ([]models.SpilloverPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, spilloverPolicyURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Policies []models.SpilloverPolicy `json:"spilloverpolicy"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Policies, nil
}

func (s *SpilloverService) GetSpilloverPolicy(name string) ([]models.SpilloverPolicy, error) {
	urlReq := fmt.Sprintf("%s/%s", spilloverPolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Policies []models.SpilloverPolicy `json:"spilloverpolicy"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Policies, nil
}

func (s *SpilloverService) CountSpilloverPolicy() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, spilloverPolicyURL+"?count=yes", nil)
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
		} `json:"spilloverpolicy"`
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

func (s *SpilloverService) RenameSpilloverPolicy(name string, newname string) error {
	payload := map[string]any{
		"spilloverpolicy": map[string]string{
			"name":    name,
			"newname": newname,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, spilloverPolicyURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// spilloverpolicy_binding
// Binding object which returns the resources bound to spilloverpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/spillover/spilloverpolicy_binding

func (s *SpilloverService) GetAllSpilloverPolicyBinding() ([]models.SpilloverPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, spilloverPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.SpilloverPolicyBinding `json:"spilloverpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *SpilloverService) GetSpilloverPolicyBinding(name string) ([]models.SpilloverPolicyBinding, error) {
	urlReq := fmt.Sprintf("%s/%s", spilloverPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.SpilloverPolicyBinding `json:"spilloverpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

// spilloverpolicy_csvserver_binding
// Binding object showing the csvserver that can be bound to spilloverpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/spillover/spilloverpolicy_csvserver_binding

func (s *SpilloverService) GetAllSpilloverPolicyCSVServerBinding() ([]models.SpilloverPolicyCSVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, spilloverPolicyCSVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.SpilloverPolicyCSVServerBinding `json:"spilloverpolicy_csvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *SpilloverService) GetSpilloverPolicyCSVServerBinding(name string) ([]models.SpilloverPolicyCSVServerBinding, error) {
	urlReq := fmt.Sprintf("%s/%s", spilloverPolicyCSVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.SpilloverPolicyCSVServerBinding `json:"spilloverpolicy_csvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *SpilloverService) CountSpilloverPolicyCSVServerBinding(name string) (float64, error) {
	urlReq := fmt.Sprintf("%s/%s?count=yes", spilloverPolicyCSVServerBindingURL, url.QueryEscape(name))
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
		} `json:"spilloverpolicy_csvserver_binding"`
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

// spilloverpolicy_gslbvserver_binding
// Binding object showing the gslbvserver that can be bound to spilloverpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/spillover/spilloverpolicy_gslbvserver_binding

func (s *SpilloverService) GetAllSpilloverPolicyGSLBVServerBinding() ([]models.SpilloverPolicyGSLBVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, spilloverPolicyGSLBVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.SpilloverPolicyGSLBVServerBinding `json:"spilloverpolicy_gslbvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *SpilloverService) GetSpilloverPolicyGSLBVServerBinding(name string) ([]models.SpilloverPolicyGSLBVServerBinding, error) {
	urlReq := fmt.Sprintf("%s/%s", spilloverPolicyGSLBVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.SpilloverPolicyGSLBVServerBinding `json:"spilloverpolicy_gslbvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *SpilloverService) CountSpilloverPolicyGSLBVServerBinding(name string) (float64, error) {
	urlReq := fmt.Sprintf("%s/%s?count=yes", spilloverPolicyGSLBVServerBindingURL, url.QueryEscape(name))
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
		} `json:"spilloverpolicy_gslbvserver_binding"`
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

// spilloverpolicy_lbvserver_binding
// Binding object showing the lbvserver that can be bound to spilloverpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/spillover/spilloverpolicy_lbvserver_binding

func (s *SpilloverService) GetAllSpilloverPolicyLBVServerBinding() ([]models.SpilloverPolicyLBVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, spilloverPolicyLBVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.SpilloverPolicyLBVServerBinding `json:"spilloverpolicy_lbvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *SpilloverService) GetSpilloverPolicyLBVServerBinding(name string) ([]models.SpilloverPolicyLBVServerBinding, error) {
	urlReq := fmt.Sprintf("%s/%s", spilloverPolicyLBVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.SpilloverPolicyLBVServerBinding `json:"spilloverpolicy_lbvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *SpilloverService) CountSpilloverPolicyLBVServerBinding(name string) (float64, error) {
	urlReq := fmt.Sprintf("%s/%s?count=yes", spilloverPolicyLBVServerBindingURL, url.QueryEscape(name))
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
		} `json:"spilloverpolicy_lbvserver_binding"`
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
