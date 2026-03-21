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
	cmpActionURL                          = "/nitro/v1/config/cmpaction"
	cmpGlobalBindingURL                   = "/nitro/v1/config/cmpglobal_binding"
	cmpGlobalCMPPolicyBindingURL          = "/nitro/v1/config/cmpglobal_cmppolicy_binding"
	cmpParameterURL                       = "/nitro/v1/config/cmpparameter"
	cmpPolicyURL                          = "/nitro/v1/config/cmppolicy"
	cmpPolicyLabelURL                     = "/nitro/v1/config/cmppolicylabel"
	cmpPolicyLabelBindingURL              = "/nitro/v1/config/cmppolicylabel_binding"
	cmpPolicyLabelCMPPolicyBindingURL     = "/nitro/v1/config/cmppolicylabel_cmppolicy_binding"
	cmpPolicyLabelPolicyBindingBindingURL = "/nitro/v1/config/cmppolicylabel_policybinding_binding"
	cmpPolicyBindingURL                   = "/nitro/v1/config/cmppolicy_binding"
	cmpPolicyCMPGlobalBindingURL          = "/nitro/v1/config/cmppolicy_cmpglobal_binding"
	cmpPolicyCMPPolicyLabelBindingURL     = "/nitro/v1/config/cmppolicy_cmppolicylabel_binding"
	cmpPolicyCRVServerBindingURL          = "/nitro/v1/config/cmppolicy_crvserver_binding"
	cmpPolicyCSVServerBindingURL          = "/nitro/v1/config/cmppolicy_csvserver_binding"
	cmpPolicyLBVServerBindingURL          = "/nitro/v1/config/cmppolicy_lbvserver_binding"
)

// Compression configuration. The system’s feature for compressing HTTP responses to compression-aware browsers.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/cmp/cmp
type CMPService struct {
	client *Client
}

// cmpaction

func (s *CMPService) AddCMPAction(action models.CMPAction) error {
	payload := map[string]any{
		"cmpaction": action,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, cmpActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CMPService) DeleteCMPAction(name string) error {
	reqURL := fmt.Sprintf("%s/%s", cmpActionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CMPService) UpdateCMPAction(action models.CMPAction) error {
	payload := map[string]any{
		"cmpaction": action,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, cmpActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CMPService) UnsetCMPAction(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{
		"cmpaction": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, cmpActionURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CMPService) GetAllCMPAction() ([]models.CMPAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, cmpActionURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Actions []models.CMPAction `json:"cmpaction"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Actions, nil
}

func (s *CMPService) GetCMPAction(name string) ([]models.CMPAction, error) {
	reqURL := fmt.Sprintf("%s/%s", cmpActionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Actions []models.CMPAction `json:"cmpaction"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Actions, nil
}

func (s *CMPService) CountCMPAction() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, cmpActionURL+"?count=yes", nil)
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
		} `json:"cmpaction"`
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

func (s *CMPService) RenameCMPAction(name, newname string) error {
	payload := map[string]any{
		"cmpaction": map[string]string{
			"name":    name,
			"newname": newname,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, cmpActionURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

// cmpglobal_binding

func (s *CMPService) GetCMPGlobalBinding() ([]models.CMPGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, cmpGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CMPGlobalBinding `json:"cmpglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// cmpglobal_cmppolicy_binding

func (s *CMPService) AddCMPGlobalCMPPolicyBinding(binding models.CMPGlobalCMPPolicyBinding) error {
	payload := map[string]any{
		"cmpglobal_cmppolicy_binding": binding,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, cmpGlobalCMPPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CMPService) DeleteCMPGlobalCMPPolicyBinding(policyname string) error {
	reqURL := fmt.Sprintf("%s?args=policyname:%s", cmpGlobalCMPPolicyBindingURL, url.QueryEscape(policyname))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CMPService) GetCMPGlobalCMPPolicyBinding() ([]models.CMPGlobalCMPPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, cmpGlobalCMPPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CMPGlobalCMPPolicyBinding `json:"cmpglobal_cmppolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CMPService) CountCMPGlobalCMPPolicyBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, cmpGlobalCMPPolicyBindingURL+"?count=yes", nil)
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
		} `json:"cmpglobal_cmppolicy_binding"`
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

// cmpparameter

func (s *CMPService) UpdateCMPParameter(param models.CMPParameter) error {
	payload := map[string]any{
		"cmpparameter": param,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, cmpParameterURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CMPService) UnsetCMPParameter(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{
		"cmpparameter": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, cmpParameterURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CMPService) GetAllCMPParameter() (models.CMPParameter, error) {
	req, err := s.client.NewRequest(http.MethodGet, cmpParameterURL, nil)
	if err != nil {
		return models.CMPParameter{}, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return models.CMPParameter{}, err
	}
	var result struct {
		Params []models.CMPParameter `json:"cmpparameter"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.CMPParameter{}, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Params) > 0 {
		return result.Params[0], nil
	}
	return models.CMPParameter{}, nil
}

// cmppolicy

func (s *CMPService) AddCOMPolicy(policy models.CMPPolicy) error {
	payload := map[string]any{
		"cmppolicy": policy,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, cmpPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CMPService) DeleteCOMPolicy(name string) error {
	reqURL := fmt.Sprintf("%s/%s", cmpPolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CMPService) UpdateCOMPolicy(policy models.CMPPolicy) error {
	payload := map[string]any{
		"cmppolicy": policy,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, cmpPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CMPService) GetAllCOMPolicy() ([]models.CMPPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, cmpPolicyURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Policies []models.CMPPolicy `json:"cmppolicy"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Policies, nil
}

func (s *CMPService) GetCOMPolicy(name string) ([]models.CMPPolicy, error) {
	reqURL := fmt.Sprintf("%s/%s", cmpPolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Policies []models.CMPPolicy `json:"cmppolicy"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Policies, nil
}

func (s *CMPService) CountCOMPolicy() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, cmpPolicyURL+"?count=yes", nil)
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
		} `json:"cmppolicy"`
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

func (s *CMPService) RenameCOMPolicy(name, newname string) error {
	payload := map[string]any{
		"cmppolicy": map[string]string{
			"name":    name,
			"newname": newname,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, cmpPolicyURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

// cmppolicylabel

func (s *CMPService) AddCMPPolicyLabel(label models.CMPPolicyLabel) error {
	payload := map[string]any{
		"cmppolicylabel": label,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, cmpPolicyLabelURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CMPService) DeleteCMPPolicyLabel(name string) error {
	reqURL := fmt.Sprintf("%s/%s", cmpPolicyLabelURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CMPService) GetAllCMPPolicyLabel() ([]models.CMPPolicyLabel, error) {
	req, err := s.client.NewRequest(http.MethodGet, cmpPolicyLabelURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Labels []models.CMPPolicyLabel `json:"cmppolicylabel"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Labels, nil
}

func (s *CMPService) GetCMPPolicyLabel(name string) ([]models.CMPPolicyLabel, error) {
	reqURL := fmt.Sprintf("%s/%s", cmpPolicyLabelURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Labels []models.CMPPolicyLabel `json:"cmppolicylabel"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Labels, nil
}

func (s *CMPService) CountCMPPolicyLabel() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, cmpPolicyLabelURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Labels []struct {
			Count float64 `json:"__count"`
		} `json:"cmppolicylabel"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Labels) > 0 {
		return result.Labels[0].Count, nil
	}
	return 0, nil
}

func (s *CMPService) RenameCMPPolicyLabel(name, newname string) error {
	payload := map[string]any{
		"cmppolicylabel": map[string]string{
			"labelname": name,
			"newname":   newname,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, cmpPolicyLabelURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

// cmppolicylabel_binding

func (s *CMPService) GetAllCMPPolicyLabelBinding() ([]models.CMPPolicyLabelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, cmpPolicyLabelBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CMPPolicyLabelBinding `json:"cmppolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CMPService) GetCMPPolicyLabelBinding(name string) ([]models.CMPPolicyLabelBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", cmpPolicyLabelBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CMPPolicyLabelBinding `json:"cmppolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// cmppolicylabel_cmppolicy_binding

func (s *CMPService) AddCMPPolicyLabelCMPPolicyBinding(binding models.CMPPolicyLabelCMPPolicyBinding) error {
	payload := map[string]any{
		"cmppolicylabel_cmppolicy_binding": binding,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, cmpPolicyLabelCMPPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CMPService) DeleteCMPPolicyLabelCMPPolicyBinding(labelname, policyname string) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s", cmpPolicyLabelCMPPolicyBindingURL, url.QueryEscape(labelname), url.QueryEscape(policyname))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CMPService) GetAllCMPPolicyLabelCMPPolicyBinding() ([]models.CMPPolicyLabelCMPPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, cmpPolicyLabelCMPPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CMPPolicyLabelCMPPolicyBinding `json:"cmppolicylabel_cmppolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CMPService) GetCMPPolicyLabelCMPPolicyBinding(name string) ([]models.CMPPolicyLabelCMPPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", cmpPolicyLabelCMPPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CMPPolicyLabelCMPPolicyBinding `json:"cmppolicylabel_cmppolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CMPService) CountCMPPolicyLabelCMPPolicyBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", cmpPolicyLabelCMPPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
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
		} `json:"cmppolicylabel_cmppolicy_binding"`
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

// cmppolicylabel_policybinding_binding

func (s *CMPService) GetAllCMPPolicyLabelPolicyBindingBinding() ([]models.CMPPolicyLabelPolicyBindingBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, cmpPolicyLabelPolicyBindingBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CMPPolicyLabelPolicyBindingBinding `json:"cmppolicylabel_policybinding_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CMPService) GetCMPPolicyLabelPolicyBindingBinding(name string) ([]models.CMPPolicyLabelPolicyBindingBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", cmpPolicyLabelPolicyBindingBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CMPPolicyLabelPolicyBindingBinding `json:"cmppolicylabel_policybinding_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CMPService) CountCMPPolicyLabelPolicyBindingBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", cmpPolicyLabelPolicyBindingBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
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
		} `json:"cmppolicylabel_policybinding_binding"`
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

// cmppolicy_binding

func (s *CMPService) GetAllCMPPolicyBinding() ([]models.CMPPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, cmpPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CMPPolicyBinding `json:"cmppolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CMPService) GetCMPPolicyBinding(name string) ([]models.CMPPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", cmpPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CMPPolicyBinding `json:"cmppolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// cmppolicy_cmpglobal_binding

func (s *CMPService) GetAllCMPPolicyCMPGlobalBinding() ([]models.CMPPolicyCMPGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, cmpPolicyCMPGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CMPPolicyCMPGlobalBinding `json:"cmppolicy_cmpglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CMPService) GetCMPPolicyCMPGlobalBinding(name string) ([]models.CMPPolicyCMPGlobalBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", cmpPolicyCMPGlobalBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CMPPolicyCMPGlobalBinding `json:"cmppolicy_cmpglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CMPService) CountCMPPolicyCMPGlobalBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", cmpPolicyCMPGlobalBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
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
		} `json:"cmppolicy_cmpglobal_binding"`
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

// cmppolicy_cmppolicylabel_binding

func (s *CMPService) GetAllCMPPolicyCMPPolicyLabelBinding() ([]models.CMPPolicyCMPPolicyLabelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, cmpPolicyCMPPolicyLabelBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CMPPolicyCMPPolicyLabelBinding `json:"cmppolicy_cmppolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CMPService) GetCMPPolicyCMPPolicyLabelBinding(name string) ([]models.CMPPolicyCMPPolicyLabelBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", cmpPolicyCMPPolicyLabelBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CMPPolicyCMPPolicyLabelBinding `json:"cmppolicy_cmppolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CMPService) CountCMPPolicyCMPPolicyLabelBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", cmpPolicyCMPPolicyLabelBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
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
		} `json:"cmppolicy_cmppolicylabel_binding"`
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

// cmppolicy_crvserver_binding

func (s *CMPService) GetAllCMPPolicyCRVserverBinding() ([]models.CMPPolicyCRVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, cmpPolicyCRVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CMPPolicyCRVServerBinding `json:"cmppolicy_crvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CMPService) GetCMPPolicyCRVserverBinding(name string) ([]models.CMPPolicyCRVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", cmpPolicyCRVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CMPPolicyCRVServerBinding `json:"cmppolicy_crvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CMPService) CountCMPPolicyCRVserverBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", cmpPolicyCRVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
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
		} `json:"cmppolicy_crvserver_binding"`
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

// cmppolicy_csvserver_binding

func (s *CMPService) GetAllCMPPolicyCSVServerBinding() ([]models.CMPPolicyCSVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, cmpPolicyCSVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CMPPolicyCSVServerBinding `json:"cmppolicy_csvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CMPService) GetCMPPolicyCSVServerBinding(name string) ([]models.CMPPolicyCSVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", cmpPolicyCSVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CMPPolicyCSVServerBinding `json:"cmppolicy_csvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CMPService) CountCMPPolicyCSVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", cmpPolicyCSVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
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
		} `json:"cmppolicy_csvserver_binding"`
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

// cmppolicy_lbvserver_binding

func (s *CMPService) GetAllCMPPolicyLBVServerBinding() ([]models.CMPPolicyLBVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, cmpPolicyLBVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CMPPolicyLBVServerBinding `json:"cmppolicy_lbvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CMPService) GetCMPPolicyLBVServerBinding(name string) ([]models.CMPPolicyLBVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", cmpPolicyLBVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CMPPolicyLBVServerBinding `json:"cmppolicy_lbvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CMPService) CountCMPPolicyLBVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", cmpPolicyLBVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
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
		} `json:"cmppolicy_lbvserver_binding"`
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
