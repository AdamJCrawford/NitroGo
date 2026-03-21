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
	rewriteActionURL                          = "/nitro/v1/config/rewriteaction"
	rewriteGlobalBindingURL                   = "/nitro/v1/config/rewriteglobal_binding"
	rewriteGlobalRewritePolicyBindingURL      = "/nitro/v1/config/rewriteglobal_rewritepolicy_binding"
	rewriteParamURL                           = "/nitro/v1/config/rewriteparam"
	rewritePolicyURL                          = "/nitro/v1/config/rewritepolicy"
	rewritePolicyLabelURL                     = "/nitro/v1/config/rewritepolicylabel"
	rewritePolicyLabelBindingURL              = "/nitro/v1/config/rewritepolicylabel_binding"
	rewritePolicyLabelPolicyBindingBindingURL = "/nitro/v1/config/rewritepolicylabel_policybinding_binding"
	rewritePolicyLabelRewritePolicyBindingURL = "/nitro/v1/config/rewritepolicylabel_rewritepolicy_binding"
	rewritePolicyBindingURL                   = "/nitro/v1/config/rewritepolicy_binding"
	rewritePolicyCSVServerBindingURL          = "/nitro/v1/config/rewritepolicy_csvserver_binding"
	rewritePolicyLBVServerBindingURL          = "/nitro/v1/config/rewritepolicy_lbvserver_binding"
	rewritePolicyRewriteGlobalBindingURL      = "/nitro/v1/config/rewritepolicy_rewriteglobal_binding"
	rewritePolicyRewritePolicyLabelBindingURL = "/nitro/v1/config/rewritepolicy_rewritepolicylabel_binding"
	rewritePolicyVPNVServerBindingURL         = "/nitro/v1/config/rewritepolicy_vpnvserver_binding"
)

// Rewrite configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/rewrite/rewrite
type RewriteService struct {
	client *Client
}

// rewriteaction

func (s *RewriteService) AddRewriteAction(action models.RewriteAction) error {
	payload := map[string]any{
		"rewriteaction": action,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, rewriteActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *RewriteService) DeleteRewriteAction(name string) error {
	reqURL := fmt.Sprintf("%s/%s", rewriteActionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *RewriteService) UpdateRewriteAction(action models.RewriteAction) error {
	payload := map[string]any{
		"rewriteaction": action,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, rewriteActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *RewriteService) UnsetRewriteAction(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{
		"rewriteaction": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, rewriteActionURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *RewriteService) GetAllRewriteAction() ([]models.RewriteAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, rewriteActionURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Actions []models.RewriteAction `json:"rewriteaction"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Actions, nil
}

func (s *RewriteService) GetRewriteAction(name string) ([]models.RewriteAction, error) {
	reqURL := fmt.Sprintf("%s/%s", rewriteActionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Actions []models.RewriteAction `json:"rewriteaction"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Actions, nil
}

func (s *RewriteService) CountRewriteAction() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, rewriteActionURL+"?count=yes", nil)
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
		} `json:"rewriteaction"`
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

func (s *RewriteService) RenameRewriteAction(name, newname string) error {
	payload := map[string]any{
		"rewriteaction": map[string]string{
			"name":    name,
			"newname": newname,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, rewriteActionURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

// rewriteglobal_binding

func (s *RewriteService) GetRewriteGlobalBinding() ([]models.RewriteGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, rewriteGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.RewriteGlobalBinding `json:"rewriteglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// rewriteglobal_rewritepolicy_binding

func (s *RewriteService) AddRewriteGlobalRewritePolicyBinding(binding models.RewriteGlobalRewritePolicyBinding) error {
	payload := map[string]any{
		"rewriteglobal_rewritepolicy_binding": binding,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, rewriteGlobalRewritePolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *RewriteService) DeleteRewriteGlobalRewritePolicyBinding(policyname string) error {
	reqURL := fmt.Sprintf("%s?args=policyname:%s", rewriteGlobalRewritePolicyBindingURL, url.QueryEscape(policyname))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *RewriteService) GetRewriteGlobalRewritePolicyBinding() ([]models.RewriteGlobalRewritePolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, rewriteGlobalRewritePolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.RewriteGlobalRewritePolicyBinding `json:"rewriteglobal_rewritepolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *RewriteService) CountRewriteGlobalRewritePolicyBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, rewriteGlobalRewritePolicyBindingURL+"?count=yes", nil)
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
		} `json:"rewriteglobal_rewritepolicy_binding"`
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

// rewriteparam

func (s *RewriteService) UpdateRewriteParam(param models.RewriteParam) error {
	payload := map[string]any{
		"rewriteparam": param,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, rewriteParamURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *RewriteService) UnsetRewriteParam(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{
		"rewriteparam": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, rewriteParamURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *RewriteService) GetAllRewriteParam() (models.RewriteParam, error) {
	req, err := s.client.NewRequest(http.MethodGet, rewriteParamURL, nil)
	if err != nil {
		return models.RewriteParam{}, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return models.RewriteParam{}, err
	}
	var result struct {
		Params []models.RewriteParam `json:"rewriteparam"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.RewriteParam{}, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Params) > 0 {
		return result.Params[0], nil
	}
	return models.RewriteParam{}, nil
}

// rewritepolicy

func (s *RewriteService) AddRewritePolicy(policy models.RewritePolicy) error {
	payload := map[string]any{
		"rewritepolicy": policy,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, rewritePolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *RewriteService) DeleteRewritePolicy(name string) error {
	reqURL := fmt.Sprintf("%s/%s", rewritePolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *RewriteService) UpdateRewritePolicy(policy models.RewritePolicy) error {
	payload := map[string]any{
		"rewritepolicy": policy,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, rewritePolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *RewriteService) UnsetRewritePolicy(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{
		"rewritepolicy": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, rewritePolicyURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *RewriteService) GetAllRewritePolicy() ([]models.RewritePolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, rewritePolicyURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Policies []models.RewritePolicy `json:"rewritepolicy"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Policies, nil
}

func (s *RewriteService) GetRewritePolicy(name string) ([]models.RewritePolicy, error) {
	reqURL := fmt.Sprintf("%s/%s", rewritePolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Policies []models.RewritePolicy `json:"rewritepolicy"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Policies, nil
}

func (s *RewriteService) CountRewritePolicy() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, rewritePolicyURL+"?count=yes", nil)
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
		} `json:"rewritepolicy"`
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

func (s *RewriteService) RenameRewritePolicy(name, newname string) error {
	payload := map[string]any{
		"rewritepolicy": map[string]string{
			"name":    name,
			"newname": newname,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, rewritePolicyURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

// rewritepolicylabel

func (s *RewriteService) AddRewritePolicyLabel(label models.RewritePolicyLabel) error {
	payload := map[string]any{
		"rewritepolicylabel": label,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, rewritePolicyLabelURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *RewriteService) DeleteRewritePolicyLabel(labelname string) error {
	reqURL := fmt.Sprintf("%s/%s", rewritePolicyLabelURL, url.QueryEscape(labelname))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *RewriteService) GetAllRewritePolicyLabel() ([]models.RewritePolicyLabel, error) {
	req, err := s.client.NewRequest(http.MethodGet, rewritePolicyLabelURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Labels []models.RewritePolicyLabel `json:"rewritepolicylabel"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Labels, nil
}

func (s *RewriteService) GetRewritePolicyLabel(labelname string) ([]models.RewritePolicyLabel, error) {
	reqURL := fmt.Sprintf("%s/%s", rewritePolicyLabelURL, url.QueryEscape(labelname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Labels []models.RewritePolicyLabel `json:"rewritepolicylabel"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Labels, nil
}

func (s *RewriteService) CountRewritePolicyLabel() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, rewritePolicyLabelURL+"?count=yes", nil)
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
		} `json:"rewritepolicylabel"`
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

func (s *RewriteService) RenameRewritePolicyLabel(labelname, newname string) error {
	payload := map[string]any{
		"rewritepolicylabel": map[string]string{
			"labelname": labelname,
			"newname":   newname,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, rewritePolicyLabelURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

// rewritepolicylabel_binding

func (s *RewriteService) GetAllRewritePolicyLabelBinding() ([]models.RewritePolicyLabelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, rewritePolicyLabelBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.RewritePolicyLabelBinding `json:"rewritepolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *RewriteService) GetRewritePolicyLabelBinding(labelname string) ([]models.RewritePolicyLabelBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", rewritePolicyLabelBindingURL, url.QueryEscape(labelname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.RewritePolicyLabelBinding `json:"rewritepolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// rewritepolicylabel_policybinding_binding

func (s *RewriteService) GetAllRewritePolicyLabelPolicyBindingBinding() ([]models.RewritePolicyLabelPolicyBindingBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, rewritePolicyLabelPolicyBindingBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.RewritePolicyLabelPolicyBindingBinding `json:"rewritepolicylabel_policybinding_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *RewriteService) GetRewritePolicyLabelPolicyBindingBinding(labelname string) ([]models.RewritePolicyLabelPolicyBindingBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", rewritePolicyLabelPolicyBindingBindingURL, url.QueryEscape(labelname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.RewritePolicyLabelPolicyBindingBinding `json:"rewritepolicylabel_policybinding_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *RewriteService) CountRewritePolicyLabelPolicyBindingBinding(labelname string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", rewritePolicyLabelPolicyBindingBindingURL, url.QueryEscape(labelname))
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
		} `json:"rewritepolicylabel_policybinding_binding"`
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

// rewritepolicylabel_rewritepolicy_binding

func (s *RewriteService) AddRewritePolicyLabelRewritePolicyBinding(binding models.RewritePolicyLabelRewritePolicyBinding) error {
	payload := map[string]any{
		"rewritepolicylabel_rewritepolicy_binding": binding,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, rewritePolicyLabelRewritePolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *RewriteService) DeleteRewritePolicyLabelRewritePolicyBinding(labelname, policyname string) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s", rewritePolicyLabelRewritePolicyBindingURL, url.QueryEscape(labelname), url.QueryEscape(policyname))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *RewriteService) GetAllRewritePolicyLabelRewritePolicyBinding() ([]models.RewritePolicyLabelRewritePolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, rewritePolicyLabelRewritePolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.RewritePolicyLabelRewritePolicyBinding `json:"rewritepolicylabel_rewritepolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *RewriteService) GetRewritePolicyLabelRewritePolicyBinding(labelname string) ([]models.RewritePolicyLabelRewritePolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", rewritePolicyLabelRewritePolicyBindingURL, url.QueryEscape(labelname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.RewritePolicyLabelRewritePolicyBinding `json:"rewritepolicylabel_rewritepolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *RewriteService) CountRewritePolicyLabelRewritePolicyBinding(labelname string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", rewritePolicyLabelRewritePolicyBindingURL, url.QueryEscape(labelname))
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
		} `json:"rewritepolicylabel_rewritepolicy_binding"`
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

// rewritepolicy_binding

func (s *RewriteService) GetAllRewritePolicyBinding() ([]models.RewritePolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, rewritePolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.RewritePolicyBinding `json:"rewritepolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *RewriteService) GetRewritePolicyBinding(name string) ([]models.RewritePolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", rewritePolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.RewritePolicyBinding `json:"rewritepolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// rewritepolicy_csvserver_binding

func (s *RewriteService) GetAllRewritePolicyCSVServerBinding() ([]models.RewritePolicyCSVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, rewritePolicyCSVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.RewritePolicyCSVServerBinding `json:"rewritepolicy_csvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *RewriteService) GetRewritePolicyCSVServerBinding(name string) ([]models.RewritePolicyCSVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", rewritePolicyCSVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.RewritePolicyCSVServerBinding `json:"rewritepolicy_csvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *RewriteService) CountRewritePolicyCSVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", rewritePolicyCSVServerBindingURL, url.QueryEscape(name))
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
		} `json:"rewritepolicy_csvserver_binding"`
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

// rewritepolicy_lbvserver_binding

func (s *RewriteService) GetAllRewritePolicyLBVServerBinding() ([]models.RewritePolicyLBVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, rewritePolicyLBVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.RewritePolicyLBVServerBinding `json:"rewritepolicy_lbvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *RewriteService) GetRewritePolicyLBVServerBinding(name string) ([]models.RewritePolicyLBVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", rewritePolicyLBVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.RewritePolicyLBVServerBinding `json:"rewritepolicy_lbvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *RewriteService) CountRewritePolicyLBVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", rewritePolicyLBVServerBindingURL, url.QueryEscape(name))
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
		} `json:"rewritepolicy_lbvserver_binding"`
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

// rewritepolicy_rewriteglobal_binding

func (s *RewriteService) GetAllRewritePolicyRewriteGlobalBinding() ([]models.RewritePolicyRewriteGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, rewritePolicyRewriteGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.RewritePolicyRewriteGlobalBinding `json:"rewritepolicy_rewriteglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *RewriteService) GetRewritePolicyRewriteGlobalBinding(name string) ([]models.RewritePolicyRewriteGlobalBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", rewritePolicyRewriteGlobalBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.RewritePolicyRewriteGlobalBinding `json:"rewritepolicy_rewriteglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *RewriteService) CountRewritePolicyRewriteGlobalBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", rewritePolicyRewriteGlobalBindingURL, url.QueryEscape(name))
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
		} `json:"rewritepolicy_rewriteglobal_binding"`
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

// rewritepolicy_rewritepolicylabel_binding

func (s *RewriteService) GetAllRewritePolicyRewritePolicyLabelBinding() ([]models.RewritePolicyRewritePolicyLabelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, rewritePolicyRewritePolicyLabelBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.RewritePolicyRewritePolicyLabelBinding `json:"rewritepolicy_rewritepolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *RewriteService) GetRewritePolicyRewritePolicyLabelBinding(name string) ([]models.RewritePolicyRewritePolicyLabelBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", rewritePolicyRewritePolicyLabelBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.RewritePolicyRewritePolicyLabelBinding `json:"rewritepolicy_rewritepolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *RewriteService) CountRewritePolicyRewritePolicyLabelBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", rewritePolicyRewritePolicyLabelBindingURL, url.QueryEscape(name))
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
		} `json:"rewritepolicy_rewritepolicylabel_binding"`
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

// rewritepolicy_vpnvserver_binding

func (s *RewriteService) GetAllRewritePolicyVPNVServerBinding() ([]models.RewritePolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, rewritePolicyVPNVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.RewritePolicyVPNVServerBinding `json:"rewritepolicy_vpnvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *RewriteService) GetRewritePolicyVPNVServerBinding(name string) ([]models.RewritePolicyVPNVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", rewritePolicyVPNVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.RewritePolicyVPNVServerBinding `json:"rewritepolicy_vpnvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *RewriteService) CountRewritePolicyVPNVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", rewritePolicyVPNVServerBindingURL, url.QueryEscape(name))
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
		} `json:"rewritepolicy_vpnvserver_binding"`
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
