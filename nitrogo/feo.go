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
	feoActionURL                 = "/nitro/v1/config/feoaction"
	feoGlobalBindingURL          = "/nitro/v1/config/feoglobal_binding"
	feoGlobalFEOPolicyBindingURL = "/nitro/v1/config/feoglobal_feopolicy_binding"
	feoParameterURL              = "/nitro/v1/config/feoparameter"
	feoPolicyURL                 = "/nitro/v1/config/feopolicy"
	feoPolicyBindingURL          = "/nitro/v1/config/feopolicy_binding"
	feoPolicyCSVServerBindingURL = "/nitro/v1/config/feopolicy_csvserver_binding"
	feoPolicyFEOGlobalBindingURL = "/nitro/v1/config/feopolicy_feoglobal_binding"
	feoPolicyLBVServerBindingURL = "/nitro/v1/config/feopolicy_lbvserver_binding"
)

// Front end optimization configuration. The system’s feature to optimize Web content for performance.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/feo/feo
type FEOService struct {
	client *Client
}

// feoaction
// Configuration for Front end optimization action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/feo/feoaction

func (s *FEOService) AddFEOAction(action models.FEOAction) error {
	payload := map[string]any{
		"feoaction": action,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, feoActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *FEOService) DeleteFEOAction(name string) error {
	urlReq := fmt.Sprintf("%s/%s", feoActionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, urlReq, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *FEOService) UpdateFEOAction(action models.FEOAction) error {
	payload := map[string]any{
		"feoaction": action,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, feoActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *FEOService) UnsetFEOAction(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"feoaction": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, feoActionURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *FEOService) GetAllFEOAction() ([]models.FEOAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, feoActionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Actions []models.FEOAction `json:"feoaction"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Actions, nil
}

func (s *FEOService) GetFEOAction(name string) ([]models.FEOAction, error) {
	urlReq := fmt.Sprintf("%s/%s", feoActionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Actions []models.FEOAction `json:"feoaction"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Actions, nil
}

func (s *FEOService) CountFEOAction() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, feoActionURL+"?count=yes", nil)
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
		} `json:"feoaction"`
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

// feoglobal_binding
// Binding object which returns the resources bound to feoglobal.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/feo/feoglobal_binding

func (s *FEOService) GetFEOGlobalBinding() ([]models.FEOGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, feoGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.FEOGlobalBinding `json:"feoglobal_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

// feoglobal_feopolicy_binding
// Binding object showing the feopolicy that can be bound to feoglobal.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/feo/feoglobal_feopolicy_binding

func (s *FEOService) AddFEOGlobalFEOPolicyBinding(binding models.FEOGlobalFEOPolicyBinding) error {
	payload := map[string]any{
		"feoglobal_feopolicy_binding": binding,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, feoGlobalFEOPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *FEOService) DeleteFEOGlobalFEOPolicyBinding(policyname string) error {
	urlReq := fmt.Sprintf("%s?args=policyname:%s", feoGlobalFEOPolicyBindingURL, url.QueryEscape(policyname))
	req, err := s.client.NewRequest(http.MethodDelete, urlReq, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *FEOService) GetFEOGlobalFEOPolicyBinding() ([]models.FEOGlobalFEOPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, feoGlobalFEOPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.FEOGlobalFEOPolicyBinding `json:"feoglobal_feopolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *FEOService) CountFEOGlobalFEOPolicyBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, feoGlobalFEOPolicyBindingURL+"?count=yes", nil)
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
		} `json:"feoglobal_feopolicy_binding"`
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

// feoparameter
// Configuration for FEO parameter resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/feo/feoparameter

func (s *FEOService) UpdateFEOParameter(param models.FEOParameter) error {
	payload := map[string]any{
		"feoparameter": param,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, feoParameterURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *FEOService) UnsetFEOParameter(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"feoparameter": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, feoParameterURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *FEOService) GetAllFEOParameter() (models.FEOParameter, error) {
	req, err := s.client.NewRequest(http.MethodGet, feoParameterURL, nil)
	if err != nil {
		return models.FEOParameter{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.FEOParameter{}, err
	}

	var result struct {
		Parameters []models.FEOParameter `json:"feoparameter"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.FEOParameter{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Parameters) > 0 {
		return result.Parameters[0], nil
	}

	return models.FEOParameter{}, nil
}

// feopolicy
// Configuration for Front end optimization policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/feo/feopolicy

func (s *FEOService) AddFEOPolicy(policy models.FEOPolicy) error {
	payload := map[string]any{
		"feopolicy": policy,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, feoPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *FEOService) DeleteFEOPolicy(name string) error {
	urlReq := fmt.Sprintf("%s/%s", feoPolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, urlReq, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *FEOService) UpdateFEOPolicy(policy models.FEOPolicy) error {
	payload := map[string]any{
		"feopolicy": policy,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, feoPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *FEOService) UnsetFEOPolicy(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"feopolicy": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, feoPolicyURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *FEOService) GetAllFEOPolicy() ([]models.FEOPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, feoPolicyURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Policies []models.FEOPolicy `json:"feopolicy"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Policies, nil
}

func (s *FEOService) GetFEOPolicy(name string) ([]models.FEOPolicy, error) {
	urlReq := fmt.Sprintf("%s/%s", feoPolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Policies []models.FEOPolicy `json:"feopolicy"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Policies, nil
}

func (s *FEOService) CountFEOPolicy() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, feoPolicyURL+"?count=yes", nil)
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
		} `json:"feopolicy"`
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

// feopolicy_binding
// Binding object which returns the resources bound to feopolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/feo/feopolicy_binding

func (s *FEOService) GetAllFEOPolicyBinding() ([]models.FEOPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, feoPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.FEOPolicyBinding `json:"feopolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *FEOService) GetFEOPolicyBinding(name string) ([]models.FEOPolicyBinding, error) {
	urlReq := fmt.Sprintf("%s/%s", feoPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.FEOPolicyBinding `json:"feopolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

// feopolicy_csvserver_binding
// Binding object showing the csvserver that can be bound to feopolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/feo/feopolicy_csvserver_binding

func (s *FEOService) GetAllFEOPolicyCSVServerBinding() ([]models.FEOPolicyCSVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, feoPolicyCSVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.FEOPolicyCSVServerBinding `json:"feopolicy_csvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *FEOService) GetFEOPolicyCSVServerBinding(name string) ([]models.FEOPolicyCSVServerBinding, error) {
	urlReq := fmt.Sprintf("%s/%s", feoPolicyCSVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.FEOPolicyCSVServerBinding `json:"feopolicy_csvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *FEOService) CountFEOPolicyCSVServerBinding(name string) (float64, error) {
	urlReq := fmt.Sprintf("%s/%s?count=yes", feoPolicyCSVServerBindingURL, url.QueryEscape(name))
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
		} `json:"feopolicy_csvserver_binding"`
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

// feopolicy_feoglobal_binding
// Binding object showing the feoglobal that can be bound to feopolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/feo/feopolicy_feoglobal_binding

func (s *FEOService) GetAllFEOPolicyFEOGlobalBinding() ([]models.FEOPolicyFEOGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, feoPolicyFEOGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.FEOPolicyFEOGlobalBinding `json:"feopolicy_feoglobal_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *FEOService) GetFEOPolicyFEOGlobalBinding(name string) ([]models.FEOPolicyFEOGlobalBinding, error) {
	urlReq := fmt.Sprintf("%s/%s", feoPolicyFEOGlobalBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.FEOPolicyFEOGlobalBinding `json:"feopolicy_feoglobal_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *FEOService) CountFEOPolicyFEOGlobalBinding(name string) (float64, error) {
	urlReq := fmt.Sprintf("%s/%s?count=yes", feoPolicyFEOGlobalBindingURL, url.QueryEscape(name))
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
		} `json:"feopolicy_feoglobal_binding"`
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

// feopolicy_lbvserver_binding
// Binding object showing the lbvserver that can be bound to feopolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/feo/feopolicy_lbvserver_binding

func (s *FEOService) GetAllFEOPolicyLBVServerBinding() ([]models.FEOPolicyLBVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, feoPolicyLBVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.FEOPolicyLBVServerBinding `json:"feopolicy_lbvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *FEOService) GetFEOPolicyLBVServerBinding(name string) ([]models.FEOPolicyLBVServerBinding, error) {
	urlReq := fmt.Sprintf("%s/%s", feoPolicyLBVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.FEOPolicyLBVServerBinding `json:"feopolicy_lbvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *FEOService) CountFEOPolicyLBVServerBinding(name string) (float64, error) {
	urlReq := fmt.Sprintf("%s/%s?count=yes", feoPolicyLBVServerBindingURL, url.QueryEscape(name))
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
		} `json:"feopolicy_lbvserver_binding"`
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
