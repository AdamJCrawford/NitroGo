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
	appQOEActionURL                 = "/nitro/v1/config/appqoeaction"
	appQOECustomRespURL             = "/nitro/v1/config/appqoecustomresp"
	appQOEParameterURL              = "/nitro/v1/config/appqoeparameter"
	appQOEPolicyURL                 = "/nitro/v1/config/appqoepolicy"
	appQOEPolicyBindingURL          = "/nitro/v1/config/appqoepolicy_binding"
	appQOEPolicyLBVServerBindingURL = "/nitro/v1/config/appqoepolicy_lbvserver_binding"
)

// Application Level Quality of Experience configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appqoe/appqoe
type AppQOEService struct {
	client *Client
}

// appqoeaction
// Configuration for AppQoS action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appqoe/appqoeaction

func (s *AppQOEService) AddAppQOEAction(action models.AppQOEAction) error {
	payload := map[string]any{
		"appqoeaction": action,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appQOEActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppQOEService) DeleteAppQOEAction(name string) error {
	urlReq := fmt.Sprintf("%s/%s", appQOEActionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, urlReq, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppQOEService) UpdateAppQOEAction(action models.AppQOEAction) error {
	payload := map[string]any{
		"appqoeaction": action,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, appQOEActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppQOEService) UnsetAppQOEAction(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"appqoeaction": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appQOEActionURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppQOEService) GetAllAppQOEAction() ([]models.AppQOEAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, appQOEActionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Actions []models.AppQOEAction `json:"appqoeaction"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Actions, nil
}

func (s *AppQOEService) GetAppQOEAction(name string) ([]models.AppQOEAction, error) {
	urlReq := fmt.Sprintf("%s/%s", appQOEActionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Actions []models.AppQOEAction `json:"appqoeaction"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Actions, nil
}

func (s *AppQOEService) CountAppQOEAction() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, appQOEActionURL+"?count=yes", nil)
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
		} `json:"appqoeaction"`
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

// appqoecustomerresp
// Configuration for AppQoE custom response page resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appqoe/appqoecustomresp

func (s *AppQOEService) ImportAppQOECustomResp(resp models.AppQOECustomResp) error {
	payload := map[string]any{
		"appqoecustomresp": resp,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appQOECustomRespURL+"?action=Import", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppQOEService) DeleteAppQOECustomResp(name string) error {
	urlReq := fmt.Sprintf("%s/%s", appQOECustomRespURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, urlReq, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppQOEService) GetAllAppQOECustomResp() ([]models.AppQOECustomResp, error) {
	req, err := s.client.NewRequest(http.MethodGet, appQOECustomRespURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		CustomResps []models.AppQOECustomResp `json:"appqoecustomresp"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.CustomResps, nil
}

func (s *AppQOEService) CountAppQOECustomResp() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, appQOECustomRespURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		CustomResps []struct {
			Count float64 `json:"__count"`
		} `json:"appqoecustomresp"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.CustomResps) > 0 {
		return result.CustomResps[0].Count, nil
	}

	return 0, nil
}

func (s *AppQOEService) ChangeAppQOECustomResp(resp models.AppQOECustomResp) error {
	payload := map[string]any{
		"appqoecustomresp": resp,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appQOECustomRespURL+"?action=change", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// appqoeparameter
// Configuration for QOS parameter resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appqoe/appqoeparameter

func (s *AppQOEService) UpdateAppQOEParameter(param models.AppQOEParameter) error {
	payload := map[string]any{
		"appqoeparameter": param,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, appQOEParameterURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppQOEService) UnsetAppQOEParameter(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"appqoeparameter": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appQOEParameterURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppQOEService) GetAllAppQOEParameter() (models.AppQOEParameter, error) {
	req, err := s.client.NewRequest(http.MethodGet, appQOEParameterURL, nil)
	if err != nil {
		return models.AppQOEParameter{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.AppQOEParameter{}, err
	}

	var result struct {
		Parameters []models.AppQOEParameter `json:"appqoeparameter"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.AppQOEParameter{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Parameters) > 0 {
		return result.Parameters[0], nil
	}

	return models.AppQOEParameter{}, nil
}

// appqoepolicy
// Configuration for AppQoS policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appqoe/appqoepolicy

func (s *AppQOEService) AddAppQOEPolicy(policy models.AppQOEPolicy) error {
	payload := map[string]any{
		"appqoepolicy": policy,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, appQOEPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppQOEService) DeleteAppQOEPolicy(name string) error {
	urlReq := fmt.Sprintf("%s/%s", appQOEPolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, urlReq, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppQOEService) UpdateAppQOEPolicy(policy models.AppQOEPolicy) error {
	payload := map[string]any{
		"appqoepolicy": policy,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, appQOEPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppQOEService) GetAllAppQOEPolicy() ([]models.AppQOEPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, appQOEPolicyURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Policies []models.AppQOEPolicy `json:"appqoepolicy"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Policies, nil
}

func (s *AppQOEService) GetAppQOEPolicy(name string) ([]models.AppQOEPolicy, error) {
	urlReq := fmt.Sprintf("%s/%s", appQOEPolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Policies []models.AppQOEPolicy `json:"appqoepolicy"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Policies, nil
}

func (s *AppQOEService) CountAppQOEPolicy() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, appQOEPolicyURL+"?count=yes", nil)
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
		} `json:"appqoepolicy"`
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

// appqoepolicy_binding
// Binding object which returns the resources bound to appqoepolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appqoe/appqoepolicy_binding

func (s *AppQOEService) GetAllAppQOEPolicyBinding() ([]models.AppQOEPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appQOEPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppQOEPolicyBinding `json:"appqoepolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppQOEService) GetAppQOEPolicyBinding(name string) ([]models.AppQOEPolicyBinding, error) {
	urlReq := fmt.Sprintf("%s/%s", appQOEPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppQOEPolicyBinding `json:"appqoepolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

// appqoepolicy_lbvserver_binding
// Binding object showing the lbvserver that can be bound to appqoepolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appqoe/appqoepolicy_lbvserver_binding

func (s *AppQOEService) GetAllAppQOEPolicyLBVServerBinding() ([]models.AppQOEPolicyLBVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appQOEPolicyLBVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppQOEPolicyLBVServerBinding `json:"appqoepolicy_lbvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppQOEService) GetAppQOEPolicyLBVServerBinding(name string) ([]models.AppQOEPolicyLBVServerBinding, error) {
	urlReq := fmt.Sprintf("%s/%s", appQOEPolicyLBVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.AppQOEPolicyLBVServerBinding `json:"appqoepolicy_lbvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *AppQOEService) CountAppQOEPolicyLBVServerBinding(name string) (float64, error) {
	urlReq := fmt.Sprintf("%s/%s?count=yes", appQOEPolicyLBVServerBindingURL, url.QueryEscape(name))
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
		} `json:"appqoepolicy_lbvserver_binding"`
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
