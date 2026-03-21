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
	auditMessageActionURL                            = "/nitro/v1/config/auditmessageaction"
	auditMessagesURL                                 = "/nitro/v1/config/auditmessages"
	auditNSLogActionURL                              = "/nitro/v1/config/auditnslogaction"
	auditNSLogGlobalAuditNSLogPolicyBindingURL       = "/nitro/v1/config/auditnslogglobal_auditnslogpolicy_binding"
	auditNSLogGlobalBindingURL                       = "/nitro/v1/config/auditnslogglobal_binding"
	auditNSLogParamsURL                              = "/nitro/v1/config/auditnslogparams"
	auditNSLogPolicyURL                              = "/nitro/v1/config/auditnslogpolicy"
	auditNSLogPolicyAAAGroupBindingURL               = "/nitro/v1/config/auditnslogpolicy_aaagroup_binding"
	auditNSLogPolicyAAAUserBindingURL                = "/nitro/v1/config/auditnslogpolicy_aaauser_binding"
	auditNSLogPolicyAppFWGlobalBindingURL            = "/nitro/v1/config/auditnslogpolicy_appfwglobal_binding"
	auditNSLogPolicyAuditNSLogGlobalBindingURL       = "/nitro/v1/config/auditnslogpolicy_auditnslogglobal_binding"
	auditNSLogPolicyAuthenticationVServerBindingURL  = "/nitro/v1/config/auditnslogpolicy_authenticationvserver_binding"
	auditNSLogPolicyBindingURL                       = "/nitro/v1/config/auditnslogpolicy_binding"
	auditNSLogPolicyCSVServerBindingURL              = "/nitro/v1/config/auditnslogpolicy_csvserver_binding"
	auditNSLogPolicyLBVServerBindingURL              = "/nitro/v1/config/auditnslogpolicy_lbvserver_binding"
	auditNSLogPolicySystemGlobalBindingURL           = "/nitro/v1/config/auditnslogpolicy_systemglobal_binding"
	auditNSLogPolicyTMGlobalBindingURL               = "/nitro/v1/config/auditnslogpolicy_tmglobal_binding"
	auditNSLogPolicyVPNGlobalBindingURL              = "/nitro/v1/config/auditnslogpolicy_vpnglobal_binding"
	auditNSLogPolicyVPNVServerBindingURL             = "/nitro/v1/config/auditnslogpolicy_vpnvserver_binding"
	auditSyslogActionURL                             = "/nitro/v1/config/auditsyslogaction"
	auditSyslogGlobalAuditSyslogPolicyBindingURL     = "/nitro/v1/config/auditsyslogglobal_auditsyslogpolicy_binding"
	auditSyslogGlobalBindingURL                      = "/nitro/v1/config/auditsyslogglobal_binding"
	auditSyslogParamsURL                             = "/nitro/v1/config/auditsyslogparams"
	auditSyslogPolicyURL                             = "/nitro/v1/config/auditsyslogpolicy"
	auditSyslogPolicyAAAGroupBindingURL              = "/nitro/v1/config/auditsyslogpolicy_aaagroup_binding"
	auditSyslogPolicyAAAUserBindingURL               = "/nitro/v1/config/auditsyslogpolicy_aaauser_binding"
	auditSyslogPolicyAuditSyslogGlobalBindingURL     = "/nitro/v1/config/auditsyslogpolicy_auditsyslogglobal_binding"
	auditSyslogPolicyAuthenticationVServerBindingURL = "/nitro/v1/config/auditsyslogpolicy_authenticationvserver_binding"
	auditSyslogPolicyBindingURL                      = "/nitro/v1/config/auditsyslogpolicy_binding"
	auditSyslogPolicyCSVServerBindingURL             = "/nitro/v1/config/auditsyslogpolicy_csvserver_binding"
	auditSyslogPolicyLBVServerBindingURL             = "/nitro/v1/config/auditsyslogpolicy_lbvserver_binding"
	auditSyslogPolicyRNATGlobalBindingURL            = "/nitro/v1/config/auditsyslogpolicy_rnatglobal_binding"
	auditSyslogPolicySystemGlobalBindingURL          = "/nitro/v1/config/auditsyslogpolicy_systemglobal_binding"
	auditSyslogPolicyTMGlobalBindingURL              = "/nitro/v1/config/auditsyslogpolicy_tmglobal_binding"
	auditSyslogPolicyVPNGlobalBindingURL             = "/nitro/v1/config/auditsyslogpolicy_vpnglobal_binding"
	auditSyslogPolicyVPNVServerBindingURL            = "/nitro/v1/config/auditsyslogpolicy_vpnvserver_binding"
)

// Audit configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/audit
type AuditService struct {
	client *Client
}

// auditmessageaction
// Configuration for message action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditmessageaction

func (s *AuditService) AddAuditMessageAction(action models.AuditMessageAction) error {
	payload := map[string]any{"auditmessageaction": action}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, auditMessageActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AuditService) DeleteAuditMessageAction(name string) error {
	reqURL := fmt.Sprintf("%s/%s", auditMessageActionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AuditService) UpdateAuditMessageAction(action models.AuditMessageAction) error {
	payload := map[string]any{"auditmessageaction": action}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, auditMessageActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AuditService) UnsetAuditMessageAction(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"auditmessageaction": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, auditMessageActionURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AuditService) GetAllAuditMessageAction() ([]models.AuditMessageAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditMessageActionURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Actions []models.AuditMessageAction `json:"auditmessageaction"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Actions, nil
}

func (s *AuditService) GetAuditMessageAction(name string) ([]models.AuditMessageAction, error) {
	reqURL := fmt.Sprintf("%s/%s", auditMessageActionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Actions []models.AuditMessageAction `json:"auditmessageaction"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Actions, nil
}

func (s *AuditService) CountAuditMessageAction() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditMessageActionURL+"?count=yes", nil)
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
		} `json:"auditmessageaction"`
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

// auditmessages
// Configuration for audit message resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditmessages

func (s *AuditService) GetAllAuditMessages() ([]models.AuditMessages, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditMessagesURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Messages []models.AuditMessages `json:"auditmessages"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Messages, nil
}

func (s *AuditService) CountAuditMessages() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditMessagesURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Messages []struct {
			Count float64 `json:"__count"`
		} `json:"auditmessages"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Messages) > 0 {
		return result.Messages[0].Count, nil
	}
	return 0, nil
}

// auditnslogaction
// Configuration for ns log action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditnslogaction

func (s *AuditService) AddAuditNSLogAction(action models.AuditNSLogAction) error {
	payload := map[string]any{"auditnslogaction": action}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, auditNSLogActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AuditService) DeleteAuditNSLogAction(name string) error {
	reqURL := fmt.Sprintf("%s/%s", auditNSLogActionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AuditService) UpdateAuditNSLogAction(action models.AuditNSLogAction) error {
	payload := map[string]any{"auditnslogaction": action}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, auditNSLogActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AuditService) UnsetAuditNSLogAction(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"auditnslogaction": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, auditNSLogActionURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AuditService) GetAllAuditNSLogAction() ([]models.AuditNSLogAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditNSLogActionURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Actions []models.AuditNSLogAction `json:"auditnslogaction"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Actions, nil
}

func (s *AuditService) GetAuditNSLogAction(name string) ([]models.AuditNSLogAction, error) {
	reqURL := fmt.Sprintf("%s/%s", auditNSLogActionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Actions []models.AuditNSLogAction `json:"auditnslogaction"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Actions, nil
}

func (s *AuditService) CountAuditNSLogAction() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditNSLogActionURL+"?count=yes", nil)
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
		} `json:"auditnslogaction"`
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

// auditnslogglobal_auditnslogpolicy_binding
// Binding object showing the auditnslogpolicy that can be bound to auditnslogglobal.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditnslogglobal_auditnslogpolicy_binding

func (s *AuditService) AddAuditNSLogGlobalAuditNSLogPolicyBinding(binding models.AuditNSLogGlobalAuditNSLogPolicyBinding) error {
	payload := map[string]any{"auditnslogglobal_auditnslogpolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, auditNSLogGlobalAuditNSLogPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AuditService) DeleteAuditNSLogGlobalAuditNSLogPolicyBinding(policyname string) error {
	reqURL := fmt.Sprintf("%s?args=policyname:%s", auditNSLogGlobalAuditNSLogPolicyBindingURL, url.QueryEscape(policyname))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AuditService) GetAuditNSLogGlobalAuditNSLogPolicyBinding() ([]models.AuditNSLogGlobalAuditNSLogPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditNSLogGlobalAuditNSLogPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditNSLogGlobalAuditNSLogPolicyBinding `json:"auditnslogglobal_auditnslogpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) CountAuditNSLogGlobalAuditNSLogPolicyBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditNSLogGlobalAuditNSLogPolicyBindingURL+"?count=yes", nil)
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
		} `json:"auditnslogglobal_auditnslogpolicy_binding"`
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

// auditnslogglobal_binding
// Binding object which returns the resources bound to auditnslogglobal.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditnslogglobal_binding

func (s *AuditService) GetAuditNSLogGlobalBinding() ([]models.AuditNSLogGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditNSLogGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditNSLogGlobalBinding `json:"auditnslogglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// auditnslogparams
// Configuration for ns log parameters resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditnslogparams

func (s *AuditService) UpdateAuditNSLogParams(params models.AuditNSLogParams) error {
	payload := map[string]any{"auditnslogparams": params}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, auditNSLogParamsURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AuditService) UnsetAuditNSLogParams(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"auditnslogparams": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, auditNSLogParamsURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AuditService) GetAllAuditNSLogParams() (models.AuditNSLogParams, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditNSLogParamsURL, nil)
	if err != nil {
		return models.AuditNSLogParams{}, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return models.AuditNSLogParams{}, err
	}
	var result struct {
		Params []models.AuditNSLogParams `json:"auditnslogparams"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.AuditNSLogParams{}, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Params) > 0 {
		return result.Params[0], nil
	}
	return models.AuditNSLogParams{}, nil
}

// auditnslogpolicy
// Configuration for ns log policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditnslogpolicy

func (s *AuditService) AddAuditNSLogPolicy(policy models.AuditNSLogPolicy) error {
	payload := map[string]any{"auditnslogpolicy": policy}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, auditNSLogPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AuditService) DeleteAuditNSLogPolicy(name string) error {
	reqURL := fmt.Sprintf("%s/%s", auditNSLogPolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AuditService) UpdateAuditNSLogPolicy(policy models.AuditNSLogPolicy) error {
	payload := map[string]any{"auditnslogpolicy": policy}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, auditNSLogPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AuditService) GetAllAuditNSLogPolicy() ([]models.AuditNSLogPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditNSLogPolicyURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Policies []models.AuditNSLogPolicy `json:"auditnslogpolicy"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Policies, nil
}

func (s *AuditService) GetAuditNSLogPolicy(name string) ([]models.AuditNSLogPolicy, error) {
	reqURL := fmt.Sprintf("%s/%s", auditNSLogPolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Policies []models.AuditNSLogPolicy `json:"auditnslogpolicy"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Policies, nil
}

func (s *AuditService) CountAuditNSLogPolicy() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditNSLogPolicyURL+"?count=yes", nil)
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
		} `json:"auditnslogpolicy"`
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

// auditnslogpolicy_aaagroup_binding
// Binding object showing the aaagroup that can be bound to auditnslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditnslogpolicy_aaagroup_binding

func (s *AuditService) GetAllAuditNSLogPolicyAAAGroupBinding() ([]models.AuditNSLogPolicyAAAGroupBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditNSLogPolicyAAAGroupBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditNSLogPolicyAAAGroupBinding `json:"auditnslogpolicy_aaagroup_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) GetAuditNSLogPolicyAAAGroupBinding(name string) ([]models.AuditNSLogPolicyAAAGroupBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", auditNSLogPolicyAAAGroupBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditNSLogPolicyAAAGroupBinding `json:"auditnslogpolicy_aaagroup_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) CountAuditNSLogPolicyAAAGroupBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", auditNSLogPolicyAAAGroupBindingURL, url.QueryEscape(name))
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
		} `json:"auditnslogpolicy_aaagroup_binding"`
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

// auditnslogpolicy_aaauser_binding
// Binding object showing the aaauser that can be bound to auditnslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditnslogpolicy_aaauser_binding

func (s *AuditService) GetAllAuditNSLogPolicyAAAUserBinding() ([]models.AuditNSLogPolicyAAAUserBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditNSLogPolicyAAAUserBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditNSLogPolicyAAAUserBinding `json:"auditnslogpolicy_aaauser_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) GetAuditNSLogPolicyAAAUserBinding(name string) ([]models.AuditNSLogPolicyAAAUserBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", auditNSLogPolicyAAAUserBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditNSLogPolicyAAAUserBinding `json:"auditnslogpolicy_aaauser_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) CountAuditNSLogPolicyAAAUserBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", auditNSLogPolicyAAAUserBindingURL, url.QueryEscape(name))
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
		} `json:"auditnslogpolicy_aaauser_binding"`
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

// auditnslogpolicy_appfwglobal_binding
// Binding object showing the appfwglobal that can be bound to auditnslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditnslogpolicy_appfwglobal_binding

func (s *AuditService) GetAllAuditNSLogPolicyAppFWGlobalBinding() ([]models.AuditNSLogPolicyAppFWGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditNSLogPolicyAppFWGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditNSLogPolicyAppFWGlobalBinding `json:"auditnslogpolicy_appfwglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) GetAuditNSLogPolicyAppFWGlobalBinding(name string) ([]models.AuditNSLogPolicyAppFWGlobalBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", auditNSLogPolicyAppFWGlobalBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditNSLogPolicyAppFWGlobalBinding `json:"auditnslogpolicy_appfwglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) CountAuditNSLogPolicyAppFWGlobalBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", auditNSLogPolicyAppFWGlobalBindingURL, url.QueryEscape(name))
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
		} `json:"auditnslogpolicy_appfwglobal_binding"`
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

// auditnslogpolicy_auditnslogglobal_binding
// Binding object showing the auditnslogglobal that can be bound to auditnslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditnslogpolicy_auditnslogglobal_binding

func (s *AuditService) GetAllAuditNSLogPolicyAuditNSLogGlobalBinding() ([]models.AuditNSLogPolicyAuditNSLogGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditNSLogPolicyAuditNSLogGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditNSLogPolicyAuditNSLogGlobalBinding `json:"auditnslogpolicy_auditnslogglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) GetAuditNSLogPolicyAuditNSLogGlobalBinding(name string) ([]models.AuditNSLogPolicyAuditNSLogGlobalBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", auditNSLogPolicyAuditNSLogGlobalBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditNSLogPolicyAuditNSLogGlobalBinding `json:"auditnslogpolicy_auditnslogglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) CountAuditNSLogPolicyAuditNSLogGlobalBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", auditNSLogPolicyAuditNSLogGlobalBindingURL, url.QueryEscape(name))
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
		} `json:"auditnslogpolicy_auditnslogglobal_binding"`
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

// auditnslogpolicy_authenticationvserver_binding
// Binding object showing the authenticationvserver that can be bound to auditnslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditnslogpolicy_authenticationvserver_binding

func (s *AuditService) GetAllAuditNSLogPolicyAuthenticationVServerBinding() ([]models.AuditNSLogPolicyAuthenticationVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditNSLogPolicyAuthenticationVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditNSLogPolicyAuthenticationVServerBinding `json:"auditnslogpolicy_authenticationvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) GetAuditNSLogPolicyAuthenticationVServerBinding(name string) ([]models.AuditNSLogPolicyAuthenticationVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", auditNSLogPolicyAuthenticationVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditNSLogPolicyAuthenticationVServerBinding `json:"auditnslogpolicy_authenticationvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) CountAuditNSLogPolicyAuthenticationVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", auditNSLogPolicyAuthenticationVServerBindingURL, url.QueryEscape(name))
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
		} `json:"auditnslogpolicy_authenticationvserver_binding"`
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

// auditnslogpolicy_binding
// Binding object which returns the resources bound to auditnslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditnslogpolicy_binding

func (s *AuditService) GetAllAuditNSLogPolicyBinding() ([]models.AuditNSLogPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditNSLogPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditNSLogPolicyBinding `json:"auditnslogpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) GetAuditNSLogPolicyBinding(name string) ([]models.AuditNSLogPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", auditNSLogPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditNSLogPolicyBinding `json:"auditnslogpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// auditnslogpolicy_csvserver_binding
// Binding object showing the csvserver that can be bound to auditnslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditnslogpolicy_csvserver_binding

func (s *AuditService) GetAllAuditNSLogPolicyCSVServerBinding() ([]models.AuditNSLogPolicyCSVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditNSLogPolicyCSVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditNSLogPolicyCSVServerBinding `json:"auditnslogpolicy_csvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) GetAuditNSLogPolicyCSVServerBinding(name string) ([]models.AuditNSLogPolicyCSVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", auditNSLogPolicyCSVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditNSLogPolicyCSVServerBinding `json:"auditnslogpolicy_csvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) CountAuditNSLogPolicyCSVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", auditNSLogPolicyCSVServerBindingURL, url.QueryEscape(name))
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
		} `json:"auditnslogpolicy_csvserver_binding"`
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

// auditnslogpolicy_lbvserver_binding
// Binding object showing the lbvserver that can be bound to auditnslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditnslogpolicy_lbvserver_binding

func (s *AuditService) GetAllAuditNSLogPolicyLBVServerBinding() ([]models.AuditNSLogPolicyLBVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditNSLogPolicyLBVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditNSLogPolicyLBVServerBinding `json:"auditnslogpolicy_lbvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) GetAuditNSLogPolicyLBVServerBinding(name string) ([]models.AuditNSLogPolicyLBVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", auditNSLogPolicyLBVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditNSLogPolicyLBVServerBinding `json:"auditnslogpolicy_lbvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) CountAuditNSLogPolicyLBVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", auditNSLogPolicyLBVServerBindingURL, url.QueryEscape(name))
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
		} `json:"auditnslogpolicy_lbvserver_binding"`
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

// auditnslogpolicy_systemglobal_binding
// Binding object showing the systemglobal that can be bound to auditnslogpolicy
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditnslogpolicy_systemglobal_binding

func (s *AuditService) GetAllAuditNSLogPolicySystemGlobalBinding() ([]models.AuditNSLogPolicySystemGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditNSLogPolicySystemGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditNSLogPolicySystemGlobalBinding `json:"auditnslogpolicy_systemglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) GetAuditNSLogPolicySystemGlobalBinding(name string) ([]models.AuditNSLogPolicySystemGlobalBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", auditNSLogPolicySystemGlobalBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditNSLogPolicySystemGlobalBinding `json:"auditnslogpolicy_systemglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) CountAuditNSLogPolicySystemGlobalBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", auditNSLogPolicySystemGlobalBindingURL, url.QueryEscape(name))
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
		} `json:"auditnslogpolicy_systemglobal_binding"`
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

// auditnslogpolicy_tmglobal_binding
// Binding object showing the tmglobal that can be bound to auditnslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditnslogpolicy_tmglobal_binding

func (s *AuditService) GetAllAuditNSLogPolicyTMGlobalBinding() ([]models.AuditNSLogPolicyTMGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditNSLogPolicyTMGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditNSLogPolicyTMGlobalBinding `json:"auditnslogpolicy_tmglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) GetAuditNSLogPolicyTMGlobalBinding(name string) ([]models.AuditNSLogPolicyTMGlobalBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", auditNSLogPolicyTMGlobalBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditNSLogPolicyTMGlobalBinding `json:"auditnslogpolicy_tmglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) CountAuditNSLogPolicyTMGlobalBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", auditNSLogPolicyTMGlobalBindingURL, url.QueryEscape(name))
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
		} `json:"auditnslogpolicy_tmglobal_binding"`
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

// auditnslogpolicy_vpnglobal_binding
// Binding object showing the vpnglobal that can be bound to auditnslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditnslogpolicy_vpnglobal_binding

func (s *AuditService) GetAllAuditNSLogPolicyVPNGlobalBinding() ([]models.AuditNSLogPolicyVPNGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditNSLogPolicyVPNGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditNSLogPolicyVPNGlobalBinding `json:"auditnslogpolicy_vpnglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) GetAuditNSLogPolicyVPNGlobalBinding(name string) ([]models.AuditNSLogPolicyVPNGlobalBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", auditNSLogPolicyVPNGlobalBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditNSLogPolicyVPNGlobalBinding `json:"auditnslogpolicy_vpnglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) CountAuditNSLogPolicyVPNGlobalBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", auditNSLogPolicyVPNGlobalBindingURL, url.QueryEscape(name))
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
		} `json:"auditnslogpolicy_vpnglobal_binding"`
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

// auditnslogpolicy_vpnvserver_binding
// Binding object showing the vpnvserver that can be bound to auditnslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditnslogpolicy_vpnvserver_binding

func (s *AuditService) GetAllAuditNSLogPolicyVPNVServerBinding() ([]models.AuditNSLogPolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditNSLogPolicyVPNVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditNSLogPolicyVPNVServerBinding `json:"auditnslogpolicy_vpnvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) GetAuditNSLogPolicyVPNVServerBinding(name string) ([]models.AuditNSLogPolicyVPNVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", auditNSLogPolicyVPNVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditNSLogPolicyVPNVServerBinding `json:"auditnslogpolicy_vpnvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) CountAuditNSLogPolicyVPNVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", auditNSLogPolicyVPNVServerBindingURL, url.QueryEscape(name))
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
		} `json:"auditnslogpolicy_vpnvserver_binding"`
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

// auditsyslogaction
// Configuration for system log action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditsyslogaction

func (s *AuditService) AddAuditSyslogAction(action models.AuditSyslogAction) error {
	payload := map[string]any{"auditsyslogaction": action}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, auditSyslogActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AuditService) DeleteAuditSyslogAction(name string) error {
	reqURL := fmt.Sprintf("%s/%s", auditSyslogActionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AuditService) UpdateAuditSyslogAction(action models.AuditSyslogAction) error {
	payload := map[string]any{"auditsyslogaction": action}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, auditSyslogActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AuditService) UnsetAuditSyslogAction(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"auditsyslogaction": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, auditSyslogActionURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AuditService) GetAllAuditSyslogAction() ([]models.AuditSyslogAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditSyslogActionURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Actions []models.AuditSyslogAction `json:"auditsyslogaction"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Actions, nil
}

func (s *AuditService) GetAuditSyslogAction(name string) ([]models.AuditSyslogAction, error) {
	reqURL := fmt.Sprintf("%s/%s", auditSyslogActionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Actions []models.AuditSyslogAction `json:"auditsyslogaction"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Actions, nil
}

func (s *AuditService) CountAuditSyslogAction() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditSyslogActionURL+"?count=yes", nil)
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
		} `json:"auditsyslogaction"`
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

// auditsyslogglobal_auditsyslogpolicy_binding
// Binding object showing the auditsyslogpolicy that can be bound to auditsyslogglobal.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditsyslogglobal_auditsyslogpolicy_binding

func (s *AuditService) AddAuditSyslogGlobalAuditSyslogPolicyBinding(binding models.AuditSyslogGlobalAuditSyslogPolicyBinding) error {
	payload := map[string]any{"auditsyslogglobal_auditsyslogpolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, auditSyslogGlobalAuditSyslogPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AuditService) DeleteAuditSyslogGlobalAuditSyslogPolicyBinding(policyname string) error {
	reqURL := fmt.Sprintf("%s?args=policyname:%s", auditSyslogGlobalAuditSyslogPolicyBindingURL, url.QueryEscape(policyname))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AuditService) GetAuditSyslogGlobalAuditSyslogPolicyBinding() ([]models.AuditSyslogGlobalAuditSyslogPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditSyslogGlobalAuditSyslogPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditSyslogGlobalAuditSyslogPolicyBinding `json:"auditsyslogglobal_auditsyslogpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) CountAuditSyslogGlobalAuditSyslogPolicyBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditSyslogGlobalAuditSyslogPolicyBindingURL+"?count=yes", nil)
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
		} `json:"auditsyslogglobal_auditsyslogpolicy_binding"`
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

// auditsyslogglobal_binding
// Binding object which returns the resources bound to auditsyslogglobal.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditsyslogglobal_binding

func (s *AuditService) GetAuditSyslogGlobalBinding() ([]models.AuditSyslogGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditSyslogGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditSyslogGlobalBinding `json:"auditsyslogglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// auditsyslogparams
// Configuration for system log parameters resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditsyslogparams

func (s *AuditService) UpdateAuditSyslogParams(params models.AuditSyslogParams) error {
	payload := map[string]any{"auditsyslogparams": params}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, auditSyslogParamsURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AuditService) UnsetAuditSyslogParams(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"auditsyslogparams": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, auditSyslogParamsURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AuditService) GetAllAuditSyslogParams() (models.AuditSyslogParams, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditSyslogParamsURL, nil)
	if err != nil {
		return models.AuditSyslogParams{}, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return models.AuditSyslogParams{}, err
	}
	var result struct {
		Params []models.AuditSyslogParams `json:"auditsyslogparams"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.AuditSyslogParams{}, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Params) > 0 {
		return result.Params[0], nil
	}
	return models.AuditSyslogParams{}, nil
}

// auditsyslogpolicy
// Configuration for system log policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditsyslogpolicy

func (s *AuditService) AddAuditSyslogPolicy(policy models.AuditSyslogPolicy) error {
	payload := map[string]any{"auditsyslogpolicy": policy}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, auditSyslogPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AuditService) DeleteAuditSyslogPolicy(name string) error {
	reqURL := fmt.Sprintf("%s/%s", auditSyslogPolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AuditService) UpdateAuditSyslogPolicy(policy models.AuditSyslogPolicy) error {
	payload := map[string]any{"auditsyslogpolicy": policy}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, auditSyslogPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AuditService) GetAllAuditSyslogPolicy() ([]models.AuditSyslogPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditSyslogPolicyURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Policies []models.AuditSyslogPolicy `json:"auditsyslogpolicy"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Policies, nil
}

func (s *AuditService) GetAuditSyslogPolicy(name string) ([]models.AuditSyslogPolicy, error) {
	reqURL := fmt.Sprintf("%s/%s", auditSyslogPolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Policies []models.AuditSyslogPolicy `json:"auditsyslogpolicy"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Policies, nil
}

func (s *AuditService) CountAuditSyslogPolicy() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditSyslogPolicyURL+"?count=yes", nil)
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
		} `json:"auditsyslogpolicy"`
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

// auditsyslogpolicy_aaagroup_binding
// Binding object showing the aaagroup that can be bound to auditsyslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditsyslogpolicy_aaagroup_binding

func (s *AuditService) GetAllAuditSyslogPolicyAAAGroupBinding() ([]models.AuditSyslogPolicyAAAGroupBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditSyslogPolicyAAAGroupBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditSyslogPolicyAAAGroupBinding `json:"auditsyslogpolicy_aaagroup_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) GetAuditSyslogPolicyAAAGroupBinding(name string) ([]models.AuditSyslogPolicyAAAGroupBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", auditSyslogPolicyAAAGroupBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditSyslogPolicyAAAGroupBinding `json:"auditsyslogpolicy_aaagroup_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) CountAuditSyslogPolicyAAAGroupBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", auditSyslogPolicyAAAGroupBindingURL, url.QueryEscape(name))
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
		} `json:"auditsyslogpolicy_aaagroup_binding"`
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

// auditsyslogpolicy_aaauser_binding
// Binding object showing the aaauser that can be bound to auditsyslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditsyslogpolicy_aaauser_binding

func (s *AuditService) GetAllAuditSyslogPolicyAAAUserBinding() ([]models.AuditSyslogPolicyAAAUserBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditSyslogPolicyAAAUserBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditSyslogPolicyAAAUserBinding `json:"auditsyslogpolicy_aaauser_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) GetAuditSyslogPolicyAAAUserBinding(name string) ([]models.AuditSyslogPolicyAAAUserBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", auditSyslogPolicyAAAUserBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditSyslogPolicyAAAUserBinding `json:"auditsyslogpolicy_aaauser_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) CountAuditSyslogPolicyAAAUserBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", auditSyslogPolicyAAAUserBindingURL, url.QueryEscape(name))
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
		} `json:"auditsyslogpolicy_aaauser_binding"`
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

// auditsyslogpolicy_auditsyslogglobal_binding
// Binding object showing the auditsyslogglobal that can be bound to auditsyslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditsyslogpolicy_auditsyslogglobal_binding

func (s *AuditService) GetAllAuditSyslogPolicyAuditSyslogGlobalBinding() ([]models.AuditSyslogPolicyAuditSyslogGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditSyslogPolicyAuditSyslogGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditSyslogPolicyAuditSyslogGlobalBinding `json:"auditsyslogpolicy_auditsyslogglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) GetAuditSyslogPolicyAuditSyslogGlobalBinding(name string) ([]models.AuditSyslogPolicyAuditSyslogGlobalBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", auditSyslogPolicyAuditSyslogGlobalBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditSyslogPolicyAuditSyslogGlobalBinding `json:"auditsyslogpolicy_auditsyslogglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) CountAuditSyslogPolicyAuditSyslogGlobalBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", auditSyslogPolicyAuditSyslogGlobalBindingURL, url.QueryEscape(name))
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
		} `json:"auditsyslogpolicy_auditsyslogglobal_binding"`
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

// auditsyslogpolicy_authenticationvserver_binding
// Binding object showing the authenticationvserver that can be bound to auditsyslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditsyslogpolicy_authenticationvserver_binding

func (s *AuditService) GetAllAuditSyslogPolicyAuthenticationVServerBinding() ([]models.AuditSyslogPolicyAuthenticationVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditSyslogPolicyAuthenticationVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditSyslogPolicyAuthenticationVServerBinding `json:"auditsyslogpolicy_authenticationvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) GetAuditSyslogPolicyAuthenticationVServerBinding(name string) ([]models.AuditSyslogPolicyAuthenticationVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", auditSyslogPolicyAuthenticationVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditSyslogPolicyAuthenticationVServerBinding `json:"auditsyslogpolicy_authenticationvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) CountAuditSyslogPolicyAuthenticationVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", auditSyslogPolicyAuthenticationVServerBindingURL, url.QueryEscape(name))
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
		} `json:"auditsyslogpolicy_authenticationvserver_binding"`
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

// auditsyslogpolicy_binding
// Binding object which returns the resources bound to auditsyslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditsyslogpolicy_binding

func (s *AuditService) GetAllAuditSyslogPolicyBinding() ([]models.AuditSyslogPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditSyslogPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditSyslogPolicyBinding `json:"auditsyslogpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) GetAuditSyslogPolicyBinding(name string) ([]models.AuditSyslogPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", auditSyslogPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditSyslogPolicyBinding `json:"auditsyslogpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// auditsyslogpolicy_csvserver_binding
// Binding object showing the csvserver that can be bound to auditsyslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditsyslogpolicy_csvserver_binding

func (s *AuditService) GetAllAuditSyslogPolicyCSVServerBinding() ([]models.AuditSyslogPolicyCSVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditSyslogPolicyCSVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditSyslogPolicyCSVServerBinding `json:"auditsyslogpolicy_csvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) GetAuditSyslogPolicyCSVServerBinding(name string) ([]models.AuditSyslogPolicyCSVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", auditSyslogPolicyCSVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditSyslogPolicyCSVServerBinding `json:"auditsyslogpolicy_csvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) CountAuditSyslogPolicyCSVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", auditSyslogPolicyCSVServerBindingURL, url.QueryEscape(name))
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
		} `json:"auditsyslogpolicy_csvserver_binding"`
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

// auditsyslogpolicy_lbvserver_binding
// Binding object showing the lbvserver that can be bound to auditsyslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditsyslogpolicy_lbvserver_binding

func (s *AuditService) GetAllAuditSyslogPolicyLBVServerBinding() ([]models.AuditSyslogPolicyLBVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditSyslogPolicyLBVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditSyslogPolicyLBVServerBinding `json:"auditsyslogpolicy_lbvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) GetAuditSyslogPolicyLBVServerBinding(name string) ([]models.AuditSyslogPolicyLBVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", auditSyslogPolicyLBVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditSyslogPolicyLBVServerBinding `json:"auditsyslogpolicy_lbvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) CountAuditSyslogPolicyLBVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", auditSyslogPolicyLBVServerBindingURL, url.QueryEscape(name))
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
		} `json:"auditsyslogpolicy_lbvserver_binding"`
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

// auditsyslogpolicy_rnatglobal_binding
// Binding object showing the rnatglobal that can be bound to auditsyslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditsyslogpolicy_rnatglobal_binding

func (s *AuditService) GetAllAuditSyslogPolicyRNATGlobalBinding() ([]models.AuditSyslogPolicyRNATGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditSyslogPolicyRNATGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditSyslogPolicyRNATGlobalBinding `json:"auditsyslogpolicy_rnatglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) GetAuditSyslogPolicyRNATGlobalBinding(name string) ([]models.AuditSyslogPolicyRNATGlobalBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", auditSyslogPolicyRNATGlobalBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditSyslogPolicyRNATGlobalBinding `json:"auditsyslogpolicy_rnatglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) CountAuditSyslogPolicyRNATGlobalBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", auditSyslogPolicyRNATGlobalBindingURL, url.QueryEscape(name))
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
		} `json:"auditsyslogpolicy_rnatglobal_binding"`
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

// auditsyslogpolicy_systemglobal_binding
// Binding object showing the systemglobal that can be bound to auditsyslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditsyslogpolicy_systemglobal_binding

func (s *AuditService) GetAllAuditSyslogPolicySystemGlobalBinding() ([]models.AuditSyslogPolicySystemGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditSyslogPolicySystemGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditSyslogPolicySystemGlobalBinding `json:"auditsyslogpolicy_systemglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) GetAuditSyslogPolicySystemGlobalBinding(name string) ([]models.AuditSyslogPolicySystemGlobalBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", auditSyslogPolicySystemGlobalBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditSyslogPolicySystemGlobalBinding `json:"auditsyslogpolicy_systemglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) CountAuditSyslogPolicySystemGlobalBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", auditSyslogPolicySystemGlobalBindingURL, url.QueryEscape(name))
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
		} `json:"auditsyslogpolicy_systemglobal_binding"`
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

// auditsyslogpolicy_tmpglobal_binding
// Binding object showing the tmglobal that can be bound to auditsyslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditsyslogpolicy_tmglobal_binding

func (s *AuditService) GetAllAuditSyslogPolicyTMGlobalBinding() ([]models.AuditSyslogPolicyTMGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditSyslogPolicyTMGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditSyslogPolicyTMGlobalBinding `json:"auditsyslogpolicy_tmglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) GetAuditSyslogPolicyTMGlobalBinding(name string) ([]models.AuditSyslogPolicyTMGlobalBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", auditSyslogPolicyTMGlobalBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditSyslogPolicyTMGlobalBinding `json:"auditsyslogpolicy_tmglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) CountAuditSyslogPolicyTMGlobalBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", auditSyslogPolicyTMGlobalBindingURL, url.QueryEscape(name))
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
		} `json:"auditsyslogpolicy_tmglobal_binding"`
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

// auditsyslogpolicy_vpnglobal_binding
// Binding object showing the vpnglobal that can be bound to auditsyslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditsyslogpolicy_vpnglobal_binding

func (s *AuditService) GetAllAuditSyslogPolicyVPNGlobalBinding() ([]models.AuditSyslogPolicyVPNGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditSyslogPolicyVPNGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditSyslogPolicyVPNGlobalBinding `json:"auditsyslogpolicy_vpnglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) GetAuditSyslogPolicyVPNGlobalBinding(name string) ([]models.AuditSyslogPolicyVPNGlobalBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", auditSyslogPolicyVPNGlobalBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditSyslogPolicyVPNGlobalBinding `json:"auditsyslogpolicy_vpnglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) CountAuditSyslogPolicyVPNGlobalBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", auditSyslogPolicyVPNGlobalBindingURL, url.QueryEscape(name))
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
		} `json:"auditsyslogpolicy_vpnglobal_binding"`
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

// auditsyslogpolicy_vpnvserver_binding
// Binding object showing the vpnvserver that can be bound to auditsyslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditsyslogpolicy_vpnvserver_binding

func (s *AuditService) GetAllAuditSyslogPolicyVPNVServerBinding() ([]models.AuditSyslogPolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, auditSyslogPolicyVPNVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditSyslogPolicyVPNVServerBinding `json:"auditsyslogpolicy_vpnvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) GetAuditSyslogPolicyVPNVServerBinding(name string) ([]models.AuditSyslogPolicyVPNVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", auditSyslogPolicyVPNVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AuditSyslogPolicyVPNVServerBinding `json:"auditsyslogpolicy_vpnvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AuditService) CountAuditSyslogPolicyVPNVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", auditSyslogPolicyVPNVServerBindingURL, url.QueryEscape(name))
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
		} `json:"auditsyslogpolicy_vpnvserver_binding"`
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
