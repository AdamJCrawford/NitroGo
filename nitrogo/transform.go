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
	transformActionURL                            = "/nitro/v1/config/transformaction"
	transformGlobalBindingURL                     = "/nitro/v1/config/transformglobal_binding"
	transformGlobalTransformPolicyBindingURL      = "/nitro/v1/config/transformglobal_transformpolicy_binding"
	transformPolicyURL                            = "/nitro/v1/config/transformpolicy"
	transformPolicyLabelURL                       = "/nitro/v1/config/transformpolicylabel"
	transformPolicyLabelBindingURL                = "/nitro/v1/config/transformpolicylabel_binding"
	transformPolicyLabelPolicyBindingBindingURL   = "/nitro/v1/config/transformpolicylabel_policybinding_binding"
	transformPolicyLabelTransformPolicyBindingURL = "/nitro/v1/config/transformpolicylabel_transformpolicy_binding"
	transformPolicyBindingURL                     = "/nitro/v1/config/transformpolicy_binding"
	transformPolicyCSVServerBindingURL            = "/nitro/v1/config/transformpolicy_csvserver_binding"
	transformPolicyLBVServerBindingURL            = "/nitro/v1/config/transformpolicy_lbvserver_binding"
	transformPolicyTransformGlobalBindingURL      = "/nitro/v1/config/transformpolicy_transformglobal_binding"
	transformPolicyTransformPolicyLabelBindingURL = "/nitro/v1/config/transformpolicy_transformpolicylabel_binding"
	transformProfileURL                           = "/nitro/v1/config/transformprofile"
	transformProfileBindingURL                    = "/nitro/v1/config/transformprofile_binding"
	transformProfileTransformActionBindingURL     = "/nitro/v1/config/transformprofile_transformaction_binding"
)

// Transform
// URL Transformation configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/transform/transform
type TransformService struct {
	client *Client
}

// transformaction

func (s *TransformService) AddTransformAction(action models.TransformAction) error {
	payload := map[string]any{
		"transformaction": action,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, transformActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TransformService) DeleteTransformAction(name string) error {
	reqURL := fmt.Sprintf("%s/%s", transformActionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TransformService) UpdateTransformAction(action models.TransformAction) error {
	payload := map[string]any{
		"transformaction": action,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, transformActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TransformService) UnsetTransformAction(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{
		"transformaction": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, transformActionURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TransformService) GetAllTransformAction() ([]models.TransformAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, transformActionURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Actions []models.TransformAction `json:"transformaction"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Actions, nil
}

func (s *TransformService) GetTransformAction(name string) ([]models.TransformAction, error) {
	reqURL := fmt.Sprintf("%s/%s", transformActionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Actions []models.TransformAction `json:"transformaction"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Actions, nil
}

func (s *TransformService) CountTransformAction() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, transformActionURL+"?count=yes", nil)
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
		} `json:"transformaction"`
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

// transformglobal_binding

func (s *TransformService) GetTransformGlobalBinding() ([]models.TransformGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, transformGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TransformGlobalBinding `json:"transformglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// transformglobal_transformpolicy_binding

func (s *TransformService) AddTransformGlobalTransformPolicyBinding(binding models.TransformGlobalTransformPolicyBinding) error {
	payload := map[string]any{
		"transformglobal_transformpolicy_binding": binding,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, transformGlobalTransformPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TransformService) DeleteTransformGlobalTransformPolicyBinding(policyname string, typeField string, priority int) error {
	reqURL := fmt.Sprintf("%s?args=policyname:%s,type:%s,priority:%d", transformGlobalTransformPolicyBindingURL, url.QueryEscape(policyname), url.QueryEscape(typeField), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TransformService) GetTransformGlobalTransformPolicyBinding() ([]models.TransformGlobalTransformPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, transformGlobalTransformPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TransformGlobalTransformPolicyBinding `json:"transformglobal_transformpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *TransformService) CountTransformGlobalTransformPolicyBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, transformGlobalTransformPolicyBindingURL+"?count=yes", nil)
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
		} `json:"transformglobal_transformpolicy_binding"`
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

// transformpolicy

func (s *TransformService) AddTransformPolicy(policy models.TransformPolicy) error {
	payload := map[string]any{
		"transformpolicy": policy,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, transformPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TransformService) DeleteTransformPolicy(name string) error {
	reqURL := fmt.Sprintf("%s/%s", transformPolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TransformService) UpdateTransformPolicy(policy models.TransformPolicy) error {
	payload := map[string]any{
		"transformpolicy": policy,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, transformPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TransformService) UnsetTransformPolicy(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{
		"transformpolicy": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, transformPolicyURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TransformService) RenameTransformPolicy(name, newname string) error {
	payload := map[string]any{
		"transformpolicy": map[string]string{
			"name":    name,
			"newname": newname,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, transformPolicyURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TransformService) GetAllTransformPolicy() ([]models.TransformPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, transformPolicyURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Policies []models.TransformPolicy `json:"transformpolicy"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Policies, nil
}

func (s *TransformService) GetTransformPolicy(name string) ([]models.TransformPolicy, error) {
	reqURL := fmt.Sprintf("%s/%s", transformPolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Policies []models.TransformPolicy `json:"transformpolicy"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Policies, nil
}

func (s *TransformService) CountTransformPolicy() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, transformPolicyURL+"?count=yes", nil)
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
		} `json:"transformpolicy"`
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

// transformpolicylabel

func (s *TransformService) AddTransformPolicyLabel(label models.TransformPolicyLabel) error {
	payload := map[string]any{
		"transformpolicylabel": label,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, transformPolicyLabelURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TransformService) DeleteTransformPolicyLabel(labelname string) error {
	reqURL := fmt.Sprintf("%s/%s", transformPolicyLabelURL, url.QueryEscape(labelname))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TransformService) RenameTransformPolicyLabel(labelname, newname string) error {
	payload := map[string]any{
		"transformpolicylabel": map[string]string{
			"labelname": labelname,
			"newname":   newname,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, transformPolicyLabelURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TransformService) GetAllTransformPolicyLabel() ([]models.TransformPolicyLabel, error) {
	req, err := s.client.NewRequest(http.MethodGet, transformPolicyLabelURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Labels []models.TransformPolicyLabel `json:"transformpolicylabel"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Labels, nil
}

func (s *TransformService) GetTransformPolicyLabel(labelname string) ([]models.TransformPolicyLabel, error) {
	reqURL := fmt.Sprintf("%s/%s", transformPolicyLabelURL, url.QueryEscape(labelname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Labels []models.TransformPolicyLabel `json:"transformpolicylabel"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Labels, nil
}

func (s *TransformService) CountTransformPolicyLabel() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, transformPolicyLabelURL+"?count=yes", nil)
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
		} `json:"transformpolicylabel"`
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

// transformpolicylabel_binding

func (s *TransformService) GetAllTransformPolicyLabelBinding() ([]models.TransformPolicyLabelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, transformPolicyLabelBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TransformPolicyLabelBinding `json:"transformpolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *TransformService) GetTransformPolicyLabelBinding(labelname string) ([]models.TransformPolicyLabelBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", transformPolicyLabelBindingURL, url.QueryEscape(labelname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TransformPolicyLabelBinding `json:"transformpolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// transformpolicylabel_policybinding_binding

func (s *TransformService) GetAllTransformPolicyLabelPolicyBindingBinding() ([]models.TransformPolicyLabelPolicyBindingBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, transformPolicyLabelPolicyBindingBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TransformPolicyLabelPolicyBindingBinding `json:"transformpolicylabel_policybinding_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *TransformService) GetTransformPolicyLabelPolicyBindingBinding(labelname string) ([]models.TransformPolicyLabelPolicyBindingBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", transformPolicyLabelPolicyBindingBindingURL, url.QueryEscape(labelname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TransformPolicyLabelPolicyBindingBinding `json:"transformpolicylabel_policybinding_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *TransformService) CountTransformPolicyLabelPolicyBindingBinding(labelname string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", transformPolicyLabelPolicyBindingBindingURL, url.QueryEscape(labelname))
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
		} `json:"transformpolicylabel_policybinding_binding"`
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

// transformpolicylabel_transformpolicy_binding

func (s *TransformService) AddTransformPolicyLabelTransformPolicyBinding(binding models.TransformPolicyLabelTransformPolicyBinding) error {
	payload := map[string]any{
		"transformpolicylabel_transformpolicy_binding": binding,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, transformPolicyLabelTransformPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TransformService) DeleteTransformPolicyLabelTransformPolicyBinding(labelname, policyname string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s,priority:%d", transformPolicyLabelTransformPolicyBindingURL, url.QueryEscape(labelname), url.QueryEscape(policyname), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TransformService) GetAllTransformPolicyLabelTransformPolicyBinding() ([]models.TransformPolicyLabelTransformPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, transformPolicyLabelTransformPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TransformPolicyLabelTransformPolicyBinding `json:"transformpolicylabel_transformpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *TransformService) GetTransformPolicyLabelTransformPolicyBinding(labelname string) ([]models.TransformPolicyLabelTransformPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", transformPolicyLabelTransformPolicyBindingURL, url.QueryEscape(labelname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TransformPolicyLabelTransformPolicyBinding `json:"transformpolicylabel_transformpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *TransformService) CountTransformPolicyLabelTransformPolicyBinding(labelname string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", transformPolicyLabelTransformPolicyBindingURL, url.QueryEscape(labelname))
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
		} `json:"transformpolicylabel_transformpolicy_binding"`
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

// transformpolicy_binding

func (s *TransformService) GetAllTransformPolicyBinding() ([]models.TransformPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, transformPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TransformPolicyBinding `json:"transformpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *TransformService) GetTransformPolicyBinding(name string) ([]models.TransformPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", transformPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TransformPolicyBinding `json:"transformpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// transformpolicy_csvserver_binding

func (s *TransformService) GetAllTransformPolicyCSVServerBinding() ([]models.TransformPolicyCSVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, transformPolicyCSVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TransformPolicyCSVServerBinding `json:"transformpolicy_csvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *TransformService) GetTransformPolicyCSVServerBinding(name string) ([]models.TransformPolicyCSVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", transformPolicyCSVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TransformPolicyCSVServerBinding `json:"transformpolicy_csvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *TransformService) CountTransformPolicyCSVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", transformPolicyCSVServerBindingURL, url.QueryEscape(name))
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
		} `json:"transformpolicy_csvserver_binding"`
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

// transformpolicy_lbvserver_binding

func (s *TransformService) GetAllTransformPolicyLBVServerBinding() ([]models.TransformPolicyLBVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, transformPolicyLBVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TransformPolicyLBVServerBinding `json:"transformpolicy_lbvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *TransformService) GetTransformPolicyLBVServerBinding(name string) ([]models.TransformPolicyLBVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", transformPolicyLBVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TransformPolicyLBVServerBinding `json:"transformpolicy_lbvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *TransformService) CountTransformPolicyLBVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", transformPolicyLBVServerBindingURL, url.QueryEscape(name))
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
		} `json:"transformpolicy_lbvserver_binding"`
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

// transformpolicy_transformglobal_binding

func (s *TransformService) GetAllTransformPolicyTransformGlobalBinding() ([]models.TransformPolicyTransformGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, transformPolicyTransformGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TransformPolicyTransformGlobalBinding `json:"transformpolicy_transformglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *TransformService) GetTransformPolicyTransformGlobalBinding(name string) ([]models.TransformPolicyTransformGlobalBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", transformPolicyTransformGlobalBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TransformPolicyTransformGlobalBinding `json:"transformpolicy_transformglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *TransformService) CountTransformPolicyTransformGlobalBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", transformPolicyTransformGlobalBindingURL, url.QueryEscape(name))
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
		} `json:"transformpolicy_transformglobal_binding"`
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

// transformpolicy_transformpolicylabel_binding

func (s *TransformService) GetAllTransformPolicyTransformPolicyLabelBinding() ([]models.TransformPolicyTransformPolicyLabelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, transformPolicyTransformPolicyLabelBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TransformPolicyTransformPolicyLabelBinding `json:"transformpolicy_transformpolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *TransformService) GetTransformPolicyTransformPolicyLabelBinding(name string) ([]models.TransformPolicyTransformPolicyLabelBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", transformPolicyTransformPolicyLabelBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TransformPolicyTransformPolicyLabelBinding `json:"transformpolicy_transformpolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *TransformService) CountTransformPolicyTransformPolicyLabelBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", transformPolicyTransformPolicyLabelBindingURL, url.QueryEscape(name))
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
		} `json:"transformpolicy_transformpolicylabel_binding"`
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

// transformprofile

func (s *TransformService) AddTransformProfile(profile models.TransformProfile) error {
	payload := map[string]any{
		"transformprofile": profile,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, transformProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TransformService) DeleteTransformProfile(name string) error {
	reqURL := fmt.Sprintf("%s/%s", transformProfileURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TransformService) UpdateTransformProfile(profile models.TransformProfile) error {
	payload := map[string]any{
		"transformprofile": profile,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, transformProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TransformService) UnsetTransformProfile(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{
		"transformprofile": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, transformProfileURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TransformService) GetAllTransformProfile() ([]models.TransformProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, transformProfileURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Profiles []models.TransformProfile `json:"transformprofile"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Profiles, nil
}

func (s *TransformService) GetTransformProfile(name string) ([]models.TransformProfile, error) {
	reqURL := fmt.Sprintf("%s/%s", transformProfileURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Profiles []models.TransformProfile `json:"transformprofile"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Profiles, nil
}

func (s *TransformService) CountTransformProfile() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, transformProfileURL+"?count=yes", nil)
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
		} `json:"transformprofile"`
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

// transformprofile_binding

func (s *TransformService) GetAllTransformProfileBinding() ([]models.TransformProfileBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, transformProfileBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TransformProfileBinding `json:"transformprofile_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *TransformService) GetTransformProfileBinding(name string) ([]models.TransformProfileBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", transformProfileBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TransformProfileBinding `json:"transformprofile_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// transformprofile_transformaction_binding

func (s *TransformService) GetAllTransformProfileTransformActionBinding() ([]models.TransformProfileTransformActionBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, transformProfileTransformActionBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TransformProfileTransformActionBinding `json:"transformprofile_transformaction_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *TransformService) GetTransformProfileTransformActionBinding(name string) ([]models.TransformProfileTransformActionBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", transformProfileTransformActionBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TransformProfileTransformActionBinding `json:"transformprofile_transformaction_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *TransformService) CountTransformProfileTransformActionBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", transformProfileTransformActionBindingURL, url.QueryEscape(name))
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
		} `json:"transformprofile_transformaction_binding"`
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
