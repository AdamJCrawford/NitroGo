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
	authorizationActionURL                                = "/nitro/v1/config/authorizationaction"
	authorizationPolicyURL                                = "/nitro/v1/config/authorizationpolicy"
	authorizationPolicyAAAGroupBindingURL                 = "/nitro/v1/config/authorizationpolicy_aaagroup_binding"
	authorizationPolicyAAAUserBindingURL                  = "/nitro/v1/config/authorizationpolicy_aaauser_binding"
	authorizationPolicyLabelAuthorizationPolicyBindingURL = "/nitro/v1/config/authorizationpolicylabel_authorizationpolicy_binding"
	authorizationPolicyLabelBindingURL                    = "/nitro/v1/config/authorizationpolicylabel_binding"
	authorizationPolicyCSVServerBindingURL                = "/nitro/v1/config/authorizationpolicy_csvserver_binding"
	authorizationPolicyLBVServerBindingURL                = "/nitro/v1/config/authorizationpolicy_lbvserver_binding"
	authorizationPolicyLabelURL                           = "/nitro/v1/config/authorizationpolicylabel"
	authorizationPolicyAuthorizationPolicyLabelBindingURL = "/nitro/v1/config/authorizationpolicy_authorizationpolicylabel_binding"
	authorizationPolicyBindingURL                         = "/nitro/v1/config/authorizationpolicy_binding"
)

// Authorization configuration. Authorization services check which resources users are authorized to access, and grant permissions accordingly.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authorization/authorization
type AuthorizationService struct {
	client *Client
}

// authorizationaction

func (s *AuthorizationService) GetAllAuthorizationAction() ([]models.AuthorizationAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, authorizationActionURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Actions []models.AuthorizationAction `json:"authorizationaction"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Actions, nil
}

func (s *AuthorizationService) GetAuthorizationAction(name string) ([]models.AuthorizationAction, error) {
	reqURL := fmt.Sprintf("%s/%s", authorizationActionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Actions []models.AuthorizationAction `json:"authorizationaction"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Actions, nil
}

func (s *AuthorizationService) CountAuthorizationAction() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, authorizationActionURL+"?count=yes", nil)
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
		} `json:"authorizationaction"`
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

// authorizationpolicy

func (s *AuthorizationService) AddAuthorizationPolicy(policy models.AuthorizationPolicy) error {
	payload := map[string]any{
		"authorizationpolicy": policy,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, authorizationPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AuthorizationService) DeleteAuthorizationPolicy(name string) error {
	reqURL := fmt.Sprintf("%s/%s", authorizationPolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AuthorizationService) UpdateAuthorizationPolicy(policy models.AuthorizationPolicy) error {
	payload := map[string]any{
		"authorizationpolicy": policy,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, authorizationPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AuthorizationService) RenameAuthorizationPolicy(name, newname string) error {
	payload := map[string]any{
		"authorizationpolicy": map[string]string{
			"name":    name,
			"newname": newname,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, authorizationPolicyURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AuthorizationService) GetAllAuthorizationPolicy() ([]models.AuthorizationPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, authorizationPolicyURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Policies []models.AuthorizationPolicy `json:"authorizationpolicy"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Policies, nil
}

func (s *AuthorizationService) GetAuthorizationPolicy(name string) ([]models.AuthorizationPolicy, error) {
	reqURL := fmt.Sprintf("%s/%s", authorizationPolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Policies []models.AuthorizationPolicy `json:"authorizationpolicy"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Policies, nil
}

func (s *AuthorizationService) CountAuthorizationPolicy() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, authorizationPolicyURL+"?count=yes", nil)
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
		} `json:"authorizationpolicy"`
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

// authorizationpolicylabel

func (s *AuthorizationService) AddAuthorizationPolicyLabel(label models.AuthorizationPolicyLabel) error {
	payload := map[string]any{
		"authorizationpolicylabel": label,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, authorizationPolicyLabelURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AuthorizationService) DeleteAuthorizationPolicyLabel(labelname string) error {
	reqURL := fmt.Sprintf("%s/%s", authorizationPolicyLabelURL, url.QueryEscape(labelname))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AuthorizationService) RenameAuthorizationPolicyLabel(labelname, newname string) error {
	payload := map[string]any{
		"authorizationpolicylabel": map[string]string{
			"labelname": labelname,
			"newname":   newname,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, authorizationPolicyLabelURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AuthorizationService) GetAllAuthorizationPolicyLabel() ([]models.AuthorizationPolicyLabel, error) {
	req, err := s.client.NewRequest(http.MethodGet, authorizationPolicyLabelURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Labels []models.AuthorizationPolicyLabel `json:"authorizationpolicylabel"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Labels, nil
}

func (s *AuthorizationService) GetAuthorizationPolicyLabel(labelname string) ([]models.AuthorizationPolicyLabel, error) {
	reqURL := fmt.Sprintf("%s/%s", authorizationPolicyLabelURL, url.QueryEscape(labelname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Labels []models.AuthorizationPolicyLabel `json:"authorizationpolicylabel"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Labels, nil
}

func (s *AuthorizationService) CountAuthorizationPolicyLabel() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, authorizationPolicyLabelURL+"?count=yes", nil)
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
		} `json:"authorizationpolicylabel"`
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

// authorizationpolicylabel_authorizationpolicy_binding

func (s *AuthorizationService) AddAuthorizationPolicyLabelAuthorizationPolicyBinding(binding models.AuthorizationPolicyLabelAuthorizationPolicyBinding) error {
	payload := map[string]any{
		"authorizationpolicylabel_authorizationpolicy_binding": binding,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, authorizationPolicyLabelAuthorizationPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AuthorizationService) DeleteAuthorizationPolicyLabelAuthorizationPolicyBinding(labelname, policyname string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s,priority:%d", authorizationPolicyLabelAuthorizationPolicyBindingURL, url.QueryEscape(labelname), url.QueryEscape(policyname), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AuthorizationService) GetAllAuthorizationPolicyLabelAuthorizationPolicyBinding() ([]models.AuthorizationPolicyLabelAuthorizationPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authorizationPolicyLabelAuthorizationPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuthorizationPolicyLabelAuthorizationPolicyBinding `json:"authorizationpolicylabel_authorizationpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuthorizationService) GetAuthorizationPolicyLabelAuthorizationPolicyBinding(labelname string) ([]models.AuthorizationPolicyLabelAuthorizationPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", authorizationPolicyLabelAuthorizationPolicyBindingURL, url.QueryEscape(labelname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuthorizationPolicyLabelAuthorizationPolicyBinding `json:"authorizationpolicylabel_authorizationpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuthorizationService) CountAuthorizationPolicyLabelAuthorizationPolicyBinding(labelname string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", authorizationPolicyLabelAuthorizationPolicyBindingURL, url.QueryEscape(labelname))
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
		} `json:"authorizationpolicylabel_authorizationpolicy_binding"`
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

// authorizationpolicylabel_binding

func (s *AuthorizationService) GetAllAuthorizationPolicyLabelBinding() ([]models.AuthorizationPolicyLabelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authorizationPolicyLabelBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuthorizationPolicyLabelBinding `json:"authorizationpolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuthorizationService) GetAuthorizationPolicyLabelBinding(labelname string) ([]models.AuthorizationPolicyLabelBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", authorizationPolicyLabelBindingURL, url.QueryEscape(labelname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuthorizationPolicyLabelBinding `json:"authorizationpolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// authorizationpolicy_aaagroup_binding

func (s *AuthorizationService) GetAllAuthorizationPolicyAAAGroupBinding() ([]models.AuthorizationPolicyAAAGroupBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authorizationPolicyAAAGroupBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuthorizationPolicyAAAGroupBinding `json:"authorizationpolicy_aaagroup_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuthorizationService) GetAuthorizationPolicyAAAGroupBinding(name string) ([]models.AuthorizationPolicyAAAGroupBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", authorizationPolicyAAAGroupBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuthorizationPolicyAAAGroupBinding `json:"authorizationpolicy_aaagroup_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuthorizationService) CountAuthorizationPolicyAAAGroupBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", authorizationPolicyAAAGroupBindingURL, url.QueryEscape(name))
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
		} `json:"authorizationpolicy_aaagroup_binding"`
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

// authorizationpolicy_aaauser_binding

func (s *AuthorizationService) GetAllAuthorizationPolicyAAAUserBinding() ([]models.AuthorizationPolicyAAAUserBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authorizationPolicyAAAUserBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuthorizationPolicyAAAUserBinding `json:"authorizationpolicy_aaauser_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuthorizationService) GetAuthorizationPolicyAAAUserBinding(name string) ([]models.AuthorizationPolicyAAAUserBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", authorizationPolicyAAAUserBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuthorizationPolicyAAAUserBinding `json:"authorizationpolicy_aaauser_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuthorizationService) CountAuthorizationPolicyAAAUserBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", authorizationPolicyAAAUserBindingURL, url.QueryEscape(name))
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
		} `json:"authorizationpolicy_aaauser_binding"`
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

// authorizationpolicy_authorizationpolicylabel_binding

func (s *AuthorizationService) GetAllAuthorizationPolicyAuthorizationPolicyLabelBinding() ([]models.AuthorizationPolicyAuthorizationPolicyLabelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authorizationPolicyAuthorizationPolicyLabelBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuthorizationPolicyAuthorizationPolicyLabelBinding `json:"authorizationpolicy_authorizationpolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuthorizationService) GetAuthorizationPolicyAuthorizationPolicyLabelBinding(name string) ([]models.AuthorizationPolicyAuthorizationPolicyLabelBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", authorizationPolicyAuthorizationPolicyLabelBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuthorizationPolicyAuthorizationPolicyLabelBinding `json:"authorizationpolicy_authorizationpolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuthorizationService) CountAuthorizationPolicyAuthorizationPolicyLabelBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", authorizationPolicyAuthorizationPolicyLabelBindingURL, url.QueryEscape(name))
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
		} `json:"authorizationpolicy_authorizationpolicylabel_binding"`
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

// authorizationpolicy_binding

func (s *AuthorizationService) GetAllAuthorizationPolicyBinding() ([]models.AuthorizationPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authorizationPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuthorizationPolicyBinding `json:"authorizationpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuthorizationService) GetAuthorizationPolicyBinding(name string) ([]models.AuthorizationPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", authorizationPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuthorizationPolicyBinding `json:"authorizationpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// authorizationpolicy_csvserver_binding

func (s *AuthorizationService) GetAllAuthorizationPolicyCSVServerBinding() ([]models.AuthorizationPolicyCSVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authorizationPolicyCSVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuthorizationPolicyCSVServerBinding `json:"authorizationpolicy_csvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuthorizationService) GetAuthorizationPolicyCSVServerBinding(name string) ([]models.AuthorizationPolicyCSVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", authorizationPolicyCSVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuthorizationPolicyCSVServerBinding `json:"authorizationpolicy_csvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuthorizationService) CountAuthorizationPolicyCSVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", authorizationPolicyCSVServerBindingURL, url.QueryEscape(name))
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
		} `json:"authorizationpolicy_csvserver_binding"`
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

// authorizationpolicy_lbvserver_binding

func (s *AuthorizationService) GetAllAuthorizationPolicyLBVServerBinding() ([]models.AuthorizationPolicyLBVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authorizationPolicyLBVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuthorizationPolicyLBVServerBinding `json:"authorizationpolicy_lbvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuthorizationService) GetAuthorizationPolicyLBVServerBinding(name string) ([]models.AuthorizationPolicyLBVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", authorizationPolicyLBVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuthorizationPolicyLBVServerBinding `json:"authorizationpolicy_lbvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuthorizationService) CountAuthorizationPolicyLBVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", authorizationPolicyLBVServerBindingURL, url.QueryEscape(name))
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
		} `json:"authorizationpolicy_lbvserver_binding"`
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
