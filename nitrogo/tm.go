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
	tmFormSSOActionURL                             = "/nitro/v1/config/tmformssoaction"
	tmGlobalAuditNSLogPolicyBindingURL             = "/nitro/v1/config/tmglobal_auditnslogpolicy_binding"
	tmGlobalAuditSyslogPolicyBindingURL            = "/nitro/v1/config/tmglobal_auditsyslogpolicy_binding"
	tmGlobalBindingURL                             = "/nitro/v1/config/tmglobal_binding"
	tmGlobalTMSessionPolicyBindingURL              = "/nitro/v1/config/tmglobal_tmsessionpolicy_binding"
	tmGlobalTMTrafficPolicyBindingURL              = "/nitro/v1/config/tmglobal_tmtrafficpolicy_binding"
	tmSAMLSSOProfileURL                            = "/nitro/v1/config/tmsamlssoprofile"
	tmSessionActionURL                             = "/nitro/v1/config/tmsessionaction"
	tmSessionParameterURL                          = "/nitro/v1/config/tmsessionparameter"
	tmSessionPolicyURL                             = "/nitro/v1/config/tmsessionpolicy"
	tmSessionPolicyAAAGroupBindingURL              = "/nitro/v1/config/tmsessionpolicy_aaagroup_binding"
	tmSessionPolicyAAAUserBindingURL               = "/nitro/v1/config/tmsessionpolicy_aaauser_binding"
	tmSessionPolicyAuthenticationVServerBindingURL = "/nitro/v1/config/tmsessionpolicy_authenticationvserver_binding"
	tmSessionPolicyBindingURL                      = "/nitro/v1/config/tmsessionpolicy_binding"
	tmSessionPolicyTMGlobalBindingURL              = "/nitro/v1/config/tmsessionpolicy_tmglobal_binding"
	tmTrafficActionURL                             = "/nitro/v1/config/tmtrafficaction"
	tmTrafficPolicyURL                             = "/nitro/v1/config/tmtrafficpolicy"
	tmTrafficPolicyBindingURL                      = "/nitro/v1/config/tmtrafficpolicy_binding"
	tmTrafficPolicyCSVServerBindingURL             = "/nitro/v1/config/tmtrafficpolicy_csvserver_binding"
	tmTrafficPolicyLBVServerBindingURL             = "/nitro/v1/config/tmtrafficpolicy_lbvserver_binding"
	tmTrafficPolicyTMGlobalBindingURL              = "/nitro/v1/config/tmtrafficpolicy_tmglobal_binding"
)

// Traffic Management
// TM session/policy configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/tm/tm
type TMService struct {
	client *Client
}

// tmformssoaction
func (s *TMService) AddTMFormSSOAction(action models.TMFormSSOAction) error {
	payload := map[string]any{"tmformssoaction": action}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, tmFormSSOActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TMService) DeleteTMFormSSOAction(name string) error {
	reqURL := fmt.Sprintf("%s/%s", tmFormSSOActionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TMService) UpdateTMFormSSOAction(action models.TMFormSSOAction) error {
	payload := map[string]any{"tmformssoaction": action}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, tmFormSSOActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TMService) UnsetTMFormSSOAction(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"tmformssoaction": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, tmFormSSOActionURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TMService) GetAllTMFormSSOAction() ([]models.TMFormSSOAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, tmFormSSOActionURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Actions []models.TMFormSSOAction `json:"tmformssoaction"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Actions, nil
}

func (s *TMService) GetTMFormSSOAction(name string) ([]models.TMFormSSOAction, error) {
	reqURL := fmt.Sprintf("%s/%s", tmFormSSOActionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Actions []models.TMFormSSOAction `json:"tmformssoaction"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Actions, nil
}

func (s *TMService) CountTMFormSSOAction() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, tmFormSSOActionURL+"?count=yes", nil)
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
		} `json:"tmformssoaction"`
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

// tmglobal_auditnslogpolicy_binding
func (s *TMService) AddTMGlobalAuditNSLogPolicyBinding(binding models.TMGlobalAuditNSLogPolicyBinding) error {
	payload := map[string]any{"tmglobal_auditnslogpolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, tmGlobalAuditNSLogPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TMService) DeleteTMGlobalAuditNSLogPolicyBinding(policyName string) error {
	reqURL := fmt.Sprintf("%s?args=policyname:%s", tmGlobalAuditNSLogPolicyBindingURL, url.QueryEscape(policyName))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TMService) GetTMGlobalAuditNSLogPolicyBinding() ([]models.TMGlobalAuditNSLogPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, tmGlobalAuditNSLogPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TMGlobalAuditNSLogPolicyBinding `json:"tmglobal_auditnslogpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *TMService) CountTMGlobalAuditNSLogPolicyBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, tmGlobalAuditNSLogPolicyBindingURL+"?count=yes", nil)
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
		} `json:"tmglobal_auditnslogpolicy_binding"`
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

// tmglobal_auditsyslogpolicy_binding
func (s *TMService) AddTMGlobalAuditSyslogPolicyBinding(binding models.TMGlobalAuditSyslogPolicyBinding) error {
	payload := map[string]any{"tmglobal_auditsyslogpolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, tmGlobalAuditSyslogPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TMService) DeleteTMGlobalAuditSyslogPolicyBinding(policyName string) error {
	reqURL := fmt.Sprintf("%s?args=policyname:%s", tmGlobalAuditSyslogPolicyBindingURL, url.QueryEscape(policyName))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TMService) GetTMGlobalAuditSyslogPolicyBinding() ([]models.TMGlobalAuditSyslogPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, tmGlobalAuditSyslogPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TMGlobalAuditSyslogPolicyBinding `json:"tmglobal_auditsyslogpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *TMService) CountTMGlobalAuditSyslogPolicyBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, tmGlobalAuditSyslogPolicyBindingURL+"?count=yes", nil)
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
		} `json:"tmglobal_auditsyslogpolicy_binding"`
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

// tmglobal_binding
func (s *TMService) GetTMGlobalBinding() ([]models.TMGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, tmGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TMGlobalBinding `json:"tmglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// tmglobal_tmsessionpolicy_binding
func (s *TMService) AddTMGlobalTMSessionPolicyBinding(binding models.TMGlobalTMSessionPolicyBinding) error {
	payload := map[string]any{"tmglobal_tmsessionpolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, tmGlobalTMSessionPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TMService) DeleteTMGlobalTMSessionPolicyBinding(policyName string) error {
	reqURL := fmt.Sprintf("%s?args=policyname:%s", tmGlobalTMSessionPolicyBindingURL, url.QueryEscape(policyName))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TMService) GetTMGlobalTMSessionPolicyBinding() ([]models.TMGlobalTMSessionPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, tmGlobalTMSessionPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TMGlobalTMSessionPolicyBinding `json:"tmglobal_tmsessionpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *TMService) CountTMGlobalTMSessionPolicyBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, tmGlobalTMSessionPolicyBindingURL+"?count=yes", nil)
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
		} `json:"tmglobal_tmsessionpolicy_binding"`
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

// tmglobal_tmtrafficpolicy_binding
func (s *TMService) AddTMGlobalTMTrafficPolicyBinding(binding models.TMGlobalTMTrafficPolicyBinding) error {
	payload := map[string]any{"tmglobal_tmtrafficpolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, tmGlobalTMTrafficPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TMService) DeleteTMGlobalTMTrafficPolicyBinding(policyName string) error {
	reqURL := fmt.Sprintf("%s?args=policyname:%s", tmGlobalTMTrafficPolicyBindingURL, url.QueryEscape(policyName))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TMService) GetTMGlobalTMTrafficPolicyBinding() ([]models.TMGlobalTMTrafficPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, tmGlobalTMTrafficPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TMGlobalTMTrafficPolicyBinding `json:"tmglobal_tmtrafficpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *TMService) CountTMGlobalTMTrafficPolicyBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, tmGlobalTMTrafficPolicyBindingURL+"?count=yes", nil)
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
		} `json:"tmglobal_tmtrafficpolicy_binding"`
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

// tmsamlssoprofile
func (s *TMService) AddTMSAMLSSOProfile(profile models.TMSAMLSSOProfile) error {
	payload := map[string]any{"tmsamlssoprofile": profile}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, tmSAMLSSOProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TMService) DeleteTMSAMLSSOProfile(name string) error {
	reqURL := fmt.Sprintf("%s/%s", tmSAMLSSOProfileURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TMService) UpdateTMSAMLSSOProfile(profile models.TMSAMLSSOProfile) error {
	payload := map[string]any{"tmsamlssoprofile": profile}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, tmSAMLSSOProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TMService) UnsetTMSAMLSSOProfile(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"tmsamlssoprofile": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, tmSAMLSSOProfileURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TMService) GetAllTMSAMLSSOProfile() ([]models.TMSAMLSSOProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, tmSAMLSSOProfileURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Profiles []models.TMSAMLSSOProfile `json:"tmsamlssoprofile"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Profiles, nil
}

func (s *TMService) GetTMSAMLSSOProfile(name string) ([]models.TMSAMLSSOProfile, error) {
	reqURL := fmt.Sprintf("%s/%s", tmSAMLSSOProfileURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Profiles []models.TMSAMLSSOProfile `json:"tmsamlssoprofile"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Profiles, nil
}

func (s *TMService) CountTMSAMLSSOProfile() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, tmSAMLSSOProfileURL+"?count=yes", nil)
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
		} `json:"tmsamlssoprofile"`
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

// tmsessionaction
func (s *TMService) AddTMSessionAction(action models.TMSessionAction) error {
	payload := map[string]any{"tmsessionaction": action}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, tmSessionActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TMService) DeleteTMSessionAction(name string) error {
	reqURL := fmt.Sprintf("%s/%s", tmSessionActionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TMService) UpdateTMSessionAction(action models.TMSessionAction) error {
	payload := map[string]any{"tmsessionaction": action}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, tmSessionActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TMService) UnsetTMSessionAction(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"tmsessionaction": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, tmSessionActionURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TMService) GetAllTMSessionAction() ([]models.TMSessionAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, tmSessionActionURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Actions []models.TMSessionAction `json:"tmsessionaction"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Actions, nil
}

func (s *TMService) GetTMSessionAction(name string) ([]models.TMSessionAction, error) {
	reqURL := fmt.Sprintf("%s/%s", tmSessionActionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Actions []models.TMSessionAction `json:"tmsessionaction"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Actions, nil
}

func (s *TMService) CountTMSessionAction() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, tmSessionActionURL+"?count=yes", nil)
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
		} `json:"tmsessionaction"`
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

// tmsessionparameter
func (s *TMService) UpdateTMSessionParameter(param models.TMSessionParameter) error {
	payload := map[string]any{"tmsessionparameter": param}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, tmSessionParameterURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TMService) UnsetTMSessionParameter(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"tmsessionparameter": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, tmSessionParameterURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TMService) GetAllTMSessionParameter() (models.TMSessionParameter, error) {
	req, err := s.client.NewRequest(http.MethodGet, tmSessionParameterURL, nil)
	if err != nil {
		return models.TMSessionParameter{}, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return models.TMSessionParameter{}, err
	}
	var result struct {
		Params []models.TMSessionParameter `json:"tmsessionparameter"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.TMSessionParameter{}, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Params) > 0 {
		return result.Params[0], nil
	}
	return models.TMSessionParameter{}, nil
}

// tmsessionpolicy
func (s *TMService) AddTMSessionPolicy(policy models.TMSessionPolicy) error {
	payload := map[string]any{"tmsessionpolicy": policy}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, tmSessionPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TMService) DeleteTMSessionPolicy(name string) error {
	reqURL := fmt.Sprintf("%s/%s", tmSessionPolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TMService) UpdateTMSessionPolicy(policy models.TMSessionPolicy) error {
	payload := map[string]any{"tmsessionpolicy": policy}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, tmSessionPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TMService) UnsetTMSessionPolicy(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"tmsessionpolicy": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, tmSessionPolicyURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TMService) GetAllTMSessionPolicy() ([]models.TMSessionPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, tmSessionPolicyURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Policies []models.TMSessionPolicy `json:"tmsessionpolicy"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Policies, nil
}

func (s *TMService) GetTMSessionPolicy(name string) ([]models.TMSessionPolicy, error) {
	reqURL := fmt.Sprintf("%s/%s", tmSessionPolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Policies []models.TMSessionPolicy `json:"tmsessionpolicy"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Policies, nil
}

func (s *TMService) CountTMSessionPolicy() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, tmSessionPolicyURL+"?count=yes", nil)
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
		} `json:"tmsessionpolicy"`
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

// tmsessionpolicy_aaagroup_binding
func (s *TMService) GetAllTMSessionPolicyAAAGroupBinding() ([]models.TMSessionPolicyAAAGroupBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, tmSessionPolicyAAAGroupBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TMSessionPolicyAAAGroupBinding `json:"tmsessionpolicy_aaagroup_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *TMService) GetTMSessionPolicyAAAGroupBinding(name string) ([]models.TMSessionPolicyAAAGroupBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", tmSessionPolicyAAAGroupBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TMSessionPolicyAAAGroupBinding `json:"tmsessionpolicy_aaagroup_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *TMService) CountTMSessionPolicyAAAGroupBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", tmSessionPolicyAAAGroupBindingURL, url.QueryEscape(name))
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
		} `json:"tmsessionpolicy_aaagroup_binding"`
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

// tmsessionpolicy_aaauser_binding
func (s *TMService) GetAllTMSessionPolicyAAAUserBinding() ([]models.TMSessionPolicyAAAUserBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, tmSessionPolicyAAAUserBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TMSessionPolicyAAAUserBinding `json:"tmsessionpolicy_aaauser_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *TMService) GetTMSessionPolicyAAAUserBinding(name string) ([]models.TMSessionPolicyAAAUserBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", tmSessionPolicyAAAUserBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TMSessionPolicyAAAUserBinding `json:"tmsessionpolicy_aaauser_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *TMService) CountTMSessionPolicyAAAUserBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", tmSessionPolicyAAAUserBindingURL, url.QueryEscape(name))
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
		} `json:"tmsessionpolicy_aaauser_binding"`
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

// tmsessionpolicy_authenticationvserver_binding
func (s *TMService) GetAllTMSessionPolicyAuthenticationVServerBinding() ([]models.TMSessionPolicyAuthenticationVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, tmSessionPolicyAuthenticationVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TMSessionPolicyAuthenticationVServerBinding `json:"tmsessionpolicy_authenticationvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *TMService) GetTMSessionPolicyAuthenticationVServerBinding(name string) ([]models.TMSessionPolicyAuthenticationVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", tmSessionPolicyAuthenticationVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TMSessionPolicyAuthenticationVServerBinding `json:"tmsessionpolicy_authenticationvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *TMService) CountTMSessionPolicyAuthenticationVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", tmSessionPolicyAuthenticationVServerBindingURL, url.QueryEscape(name))
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
		} `json:"tmsessionpolicy_authenticationvserver_binding"`
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

// tmsessionpolicy_binding
func (s *TMService) GetAllTMSessionPolicyBinding() ([]models.TMSessionPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, tmSessionPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TMSessionPolicyBinding `json:"tmsessionpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *TMService) GetTMSessionPolicyBinding(name string) ([]models.TMSessionPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", tmSessionPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TMSessionPolicyBinding `json:"tmsessionpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// tmsessionpolicy_tmglobal_binding
func (s *TMService) GetAllTMSessionPolicyTMGlobalBinding() ([]models.TMSessionPolicyTMGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, tmSessionPolicyTMGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TMSessionPolicyTMGlobalBinding `json:"tmsessionpolicy_tmglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *TMService) GetTMSessionPolicyTMGlobalBinding(name string) ([]models.TMSessionPolicyTMGlobalBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", tmSessionPolicyTMGlobalBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TMSessionPolicyTMGlobalBinding `json:"tmsessionpolicy_tmglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *TMService) CountTMSessionPolicyTMGlobalBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", tmSessionPolicyTMGlobalBindingURL, url.QueryEscape(name))
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
		} `json:"tmsessionpolicy_tmglobal_binding"`
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

// tmtrafficaction
func (s *TMService) AddTMTrafficAction(action models.TMTrafficAction) error {
	payload := map[string]any{"tmtrafficaction": action}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, tmTrafficActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TMService) DeleteTMTrafficAction(name string) error {
	reqURL := fmt.Sprintf("%s/%s", tmTrafficActionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TMService) UpdateTMTrafficAction(action models.TMTrafficAction) error {
	payload := map[string]any{"tmtrafficaction": action}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, tmTrafficActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TMService) UnsetTMTrafficAction(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"tmtrafficaction": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, tmTrafficActionURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TMService) GetAllTMTrafficAction() ([]models.TMTrafficAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, tmTrafficActionURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Actions []models.TMTrafficAction `json:"tmtrafficaction"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Actions, nil
}

func (s *TMService) GetTMTrafficAction(name string) ([]models.TMTrafficAction, error) {
	reqURL := fmt.Sprintf("%s/%s", tmTrafficActionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Actions []models.TMTrafficAction `json:"tmtrafficaction"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Actions, nil
}

func (s *TMService) CountTMTrafficAction() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, tmTrafficActionURL+"?count=yes", nil)
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
		} `json:"tmtrafficaction"`
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

// tmtrafficpolicy
func (s *TMService) AddTMTrafficPolicy(policy models.TMTrafficPolicy) error {
	payload := map[string]any{"tmtrafficpolicy": policy}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, tmTrafficPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TMService) DeleteTMTrafficPolicy(name string) error {
	reqURL := fmt.Sprintf("%s/%s", tmTrafficPolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TMService) UpdateTMTrafficPolicy(policy models.TMTrafficPolicy) error {
	payload := map[string]any{"tmtrafficpolicy": policy}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, tmTrafficPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TMService) UnsetTMTrafficPolicy(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"tmtrafficpolicy": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, tmTrafficPolicyURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *TMService) GetAllTMTrafficPolicy() ([]models.TMTrafficPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, tmTrafficPolicyURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Policies []models.TMTrafficPolicy `json:"tmtrafficpolicy"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Policies, nil
}

func (s *TMService) GetTMTrafficPolicy(name string) ([]models.TMTrafficPolicy, error) {
	reqURL := fmt.Sprintf("%s/%s", tmTrafficPolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Policies []models.TMTrafficPolicy `json:"tmtrafficpolicy"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Policies, nil
}

func (s *TMService) CountTMTrafficPolicy() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, tmTrafficPolicyURL+"?count=yes", nil)
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
		} `json:"tmtrafficpolicy"`
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

// tmtrafficpolicy_binding
func (s *TMService) GetAllTMTrafficPolicyBinding() ([]models.TMTrafficPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, tmTrafficPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TMTrafficPolicyBinding `json:"tmtrafficpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *TMService) GetTMTrafficPolicyBinding(name string) ([]models.TMTrafficPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", tmTrafficPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TMTrafficPolicyBinding `json:"tmtrafficpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// tmtrafficpolicy_csvserver_binding
func (s *TMService) GetAllTMTrafficPolicyCSVServerBinding() ([]models.TMTrafficPolicyCSVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, tmTrafficPolicyCSVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TMTrafficPolicyCSVServerBinding `json:"tmtrafficpolicy_csvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *TMService) GetTMTrafficPolicyCSVServerBinding(name string) ([]models.TMTrafficPolicyCSVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", tmTrafficPolicyCSVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TMTrafficPolicyCSVServerBinding `json:"tmtrafficpolicy_csvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *TMService) CountTMTrafficPolicyCSVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", tmTrafficPolicyCSVServerBindingURL, url.QueryEscape(name))
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
		} `json:"tmtrafficpolicy_csvserver_binding"`
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

// tmtrafficpolicy_lbvserver_binding
func (s *TMService) GetAllTMTrafficPolicyLBVServerBinding() ([]models.TMTrafficPolicyLBVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, tmTrafficPolicyLBVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TMTrafficPolicyLBVServerBinding `json:"tmtrafficpolicy_lbvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *TMService) GetTMTrafficPolicyLBVServerBinding(name string) ([]models.TMTrafficPolicyLBVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", tmTrafficPolicyLBVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TMTrafficPolicyLBVServerBinding `json:"tmtrafficpolicy_lbvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *TMService) CountTMTrafficPolicyLBVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", tmTrafficPolicyLBVServerBindingURL, url.QueryEscape(name))
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
		} `json:"tmtrafficpolicy_lbvserver_binding"`
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

// tmtrafficpolicy_tmglobal_binding
func (s *TMService) GetAllTMTrafficPolicyTMGlobalBinding() ([]models.TMTrafficPolicyTMGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, tmTrafficPolicyTMGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TMTrafficPolicyTMGlobalBinding `json:"tmtrafficpolicy_tmglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *TMService) GetTMTrafficPolicyTMGlobalBinding(name string) ([]models.TMTrafficPolicyTMGlobalBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", tmTrafficPolicyTMGlobalBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.TMTrafficPolicyTMGlobalBinding `json:"tmtrafficpolicy_tmglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *TMService) CountTMTrafficPolicyTMGlobalBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", tmTrafficPolicyTMGlobalBindingURL, url.QueryEscape(name))
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
		} `json:"tmtrafficpolicy_tmglobal_binding"`
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
