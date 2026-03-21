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
	tunnelGlobalBindingURL                    = "/nitro/v1/config/tunnelglobal_binding"
	tunnelGlobalTunnelTrafficPolicyBindingURL = "/nitro/v1/config/tunnelglobal_tunneltrafficpolicy_binding"
	tunnelTrafficPolicyURL                    = "/nitro/v1/config/tunneltrafficpolicy"
	tunnelTrafficPolicyBindingURL             = "/nitro/v1/config/tunneltrafficpolicy_binding"
	tunnelTrafficPolicyTunnelGlobalBindingURL = "/nitro/v1/config/tunneltrafficpolicy_tunnelglobal_binding"
)

// SSL VPN Tunnel Configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/tunnel/tunnel
type TunnelService struct {
	client *Client
}

// tunnelglobal_binding
// Binding object which returns the resources bound to tunnelglobal.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/tunnel/tunnelglobal_binding

func (s *TunnelService) GetTunnelGlobalBinding() ([]models.TunnelGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, tunnelGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.TunnelGlobalBinding `json:"tunnelglobal_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

// tunnelglobal_tunneltrafficpolicy_binding
// Binding object showing the tunneltrafficpolicy that can be bound to tunnelglobal.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/tunnel/tunnelglobal_tunneltrafficpolicy_binding

func (s *TunnelService) AddTunnelGlobalTunnelTrafficPolicyBinding(binding models.TunnelGlobalTunnelTrafficPolicyBinding) error {
	payload := map[string]any{
		"tunnelglobal_tunneltrafficpolicy_binding": binding,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, tunnelGlobalTunnelTrafficPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *TunnelService) DeleteTunnelGlobalTunnelTrafficPolicyBinding(policyname string) error {
	urlReq := fmt.Sprintf("%s?args=policyname:%s", tunnelGlobalTunnelTrafficPolicyBindingURL, url.QueryEscape(policyname))
	req, err := s.client.NewRequest(http.MethodDelete, urlReq, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *TunnelService) GetTunnelGlobalTunnelTrafficPolicyBinding() ([]models.TunnelGlobalTunnelTrafficPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, tunnelGlobalTunnelTrafficPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.TunnelGlobalTunnelTrafficPolicyBinding `json:"tunnelglobal_tunneltrafficpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *TunnelService) CountTunnelGlobalTunnelTrafficPolicyBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, tunnelGlobalTunnelTrafficPolicyBindingURL+"?count=yes", nil)
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
		} `json:"tunnelglobal_tunneltrafficpolicy_binding"`
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

// tunneltrafficpolicy
// Configuration for tunnel policy resource
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/tunnel/tunneltrafficpolicy

func (s *TunnelService) AddTunnelTrafficPolicy(policy models.TunnelTrafficPolicy) error {
	payload := map[string]any{
		"tunneltrafficpolicy": policy,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, tunnelTrafficPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *TunnelService) DeleteTunnelTrafficPolicy(name string) error {
	urlReq := fmt.Sprintf("%s/%s", tunnelTrafficPolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, urlReq, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *TunnelService) UpdateTunnelTrafficPolicy(policy models.TunnelTrafficPolicy) error {
	payload := map[string]any{
		"tunneltrafficpolicy": policy,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, tunnelTrafficPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *TunnelService) UnsetTunnelTrafficPolicy(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"tunneltrafficpolicy": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, tunnelTrafficPolicyURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *TunnelService) GetAllTunnelTrafficPolicy() ([]models.TunnelTrafficPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, tunnelTrafficPolicyURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Policies []models.TunnelTrafficPolicy `json:"tunneltrafficpolicy"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Policies, nil
}

func (s *TunnelService) GetTunnelTrafficPolicy(name string) ([]models.TunnelTrafficPolicy, error) {
	urlReq := fmt.Sprintf("%s/%s", tunnelTrafficPolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Policies []models.TunnelTrafficPolicy `json:"tunneltrafficpolicy"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Policies, nil
}

func (s *TunnelService) CountTunnelTrafficPolicy() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, tunnelTrafficPolicyURL+"?count=yes", nil)
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
		} `json:"tunneltrafficpolicy"`
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

func (s *TunnelService) RenameTunnelTrafficPolicy(name string, newname string) error {
	payload := map[string]any{
		"tunneltrafficpolicy": map[string]string{
			"name":    name,
			"newname": newname,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, tunnelTrafficPolicyURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// tunneltrafficpolicy_binding
// Binding object which returns the resources bound to tunneltrafficpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/tunnel/tunneltrafficpolicy_binding

func (s *TunnelService) GetAllTunnelTrafficPolicyBinding() ([]models.TunnelTrafficPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, tunnelTrafficPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.TunnelTrafficPolicyBinding `json:"tunneltrafficpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *TunnelService) GetTunnelTrafficPolicyBinding(name string) ([]models.TunnelTrafficPolicyBinding, error) {
	urlReq := fmt.Sprintf("%s/%s", tunnelTrafficPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.TunnelTrafficPolicyBinding `json:"tunneltrafficpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

// tunneltrafficpolicy_tunnelglobal_binding
// Binding object showing the tunnelglobal that can be bound to tunneltrafficpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/tunnel/tunneltrafficpolicy_tunnelglobal_binding

func (s *TunnelService) GetAllTunnelTrafficPolicyTunnelGlobalBinding() ([]models.TunnelTrafficPolicyTunnelGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, tunnelTrafficPolicyTunnelGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.TunnelTrafficPolicyTunnelGlobalBinding `json:"tunneltrafficpolicy_tunnelglobal_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *TunnelService) GetTunnelTrafficPolicyTunnelGlobalBinding(name string) ([]models.TunnelTrafficPolicyTunnelGlobalBinding, error) {
	urlReq := fmt.Sprintf("%s/%s", tunnelTrafficPolicyTunnelGlobalBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.TunnelTrafficPolicyTunnelGlobalBinding `json:"tunneltrafficpolicy_tunnelglobal_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *TunnelService) CountTunnelTrafficPolicyTunnelGlobalBinding(name string) (float64, error) {
	urlReq := fmt.Sprintf("%s/%s?count=yes", tunnelTrafficPolicyTunnelGlobalBindingURL, url.QueryEscape(name))
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
		} `json:"tunneltrafficpolicy_tunnelglobal_binding"`
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
