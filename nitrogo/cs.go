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
	csActionURL                                = "/nitro/v1/config/csaction"
	csParameterURL                             = "/nitro/v1/config/csparameter"
	csPolicyURL                                = "/nitro/v1/config/cspolicy"
	csPolicyLabelURL                           = "/nitro/v1/config/cspolicylabel"
	csPolicyLabelBindingURL                    = "/nitro/v1/config/cspolicylabel_binding"
	csPolicyLabelCSPolicyBindingURL            = "/nitro/v1/config/cspolicylabel_cspolicy_binding"
	csPolicyBindingURL                         = "/nitro/v1/config/cspolicy_binding"
	csPolicyCRVServerBindingURL                = "/nitro/v1/config/cspolicy_crvserver_binding"
	csPolicyCSPolicyLabelBindingURL            = "/nitro/v1/config/cspolicy_cspolicylabel_binding"
	csPolicyCSVServerBindingURL                = "/nitro/v1/config/cspolicy_csvserver_binding"
	csVServerURL                               = "/nitro/v1/config/csvserver"
	csVServerAnalyticsProfileBindingURL        = "/nitro/v1/config/csvserver_analyticsprofile_binding"
	csVServerAppFlowPolicyBindingURL           = "/nitro/v1/config/csvserver_appflowpolicy_binding"
	csVServerAppFWPolicyBindingURL             = "/nitro/v1/config/csvserver_appfwpolicy_binding"
	csVServerAppQOEPolicyBindingURL            = "/nitro/v1/config/csvserver_appqoepolicy_binding"
	csVServerAuditNSLogPolicyBindingURL        = "/nitro/v1/config/csvserver_auditnslogpolicy_binding"
	csVServerAuditSyslogPolicyBindingURL       = "/nitro/v1/config/csvserver_auditsyslogpolicy_binding"
	csVServerAuthorizationPolicyBindingURL     = "/nitro/v1/config/csvserver_authorizationpolicy_binding"
	csVServerBindingURL                        = "/nitro/v1/config/csvserver_binding"
	csVServerBotPolicyBindingURL               = "/nitro/v1/config/csvserver_botpolicy_binding"
	csVServerCachePolicyBindingURL             = "/nitro/v1/config/csvserver_cachepolicy_binding"
	csVServerCMPPolicyBindingURL               = "/nitro/v1/config/csvserver_cmppolicy_binding"
	csVServerContentInspectionPolicyBindingURL = "/nitro/v1/config/csvserver_contentinspectionpolicy_binding"
	csVServerCSPolicyBindingURL                = "/nitro/v1/config/csvserver_cspolicy_binding"
	csVServerDomainBindingURL                  = "/nitro/v1/config/csvserver_domain_binding"
	csVServerFEOPolicyBindingURL               = "/nitro/v1/config/csvserver_feopolicy_binding"
	csVServerGSLBVServerBindingURL             = "/nitro/v1/config/csvserver_gslbvserver_binding"
	csVServerLBVServerBindingURL               = "/nitro/v1/config/csvserver_lbvserver_binding"
	csVServerResponderPolicyBindingURL         = "/nitro/v1/config/csvserver_responderpolicy_binding"
	csVServerRewritePolicyBindingURL           = "/nitro/v1/config/csvserver_rewritepolicy_binding"
	csVServerSpilloverPolicyBindingURL         = "/nitro/v1/config/csvserver_spilloverpolicy_binding"
	csVServerTMTrafficPolicyBindingURL         = "/nitro/v1/config/csvserver_tmtrafficpolicy_binding"
	csVServerTransformPolicyBindingURL         = "/nitro/v1/config/csvserver_transformpolicy_binding"
	csVServerVPNVServerBindingURL              = "/nitro/v1/config/csvserver_vpnvserver_binding"
)

// Content Switching configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/cs/cs
type CSService struct {
	client *Client
}

// csaction
func (s *CSService) AddCSAction(action models.CSAction) error {
	payload := map[string]any{"csaction": action}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, csActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) DeleteCSAction(name string) error {
	reqURL := fmt.Sprintf("%s/%s", csActionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) UpdateCSAction(action models.CSAction) error {
	payload := map[string]any{"csaction": action}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, csActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) UnsetCSAction(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"csaction": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, csActionURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) RenameCSAction(name string, newName string) error {
	payload := map[string]any{
		"csaction": map[string]string{
			"name":    name,
			"newname": newName,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, csActionURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) GetAllCSAction() ([]models.CSAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, csActionURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Actions []models.CSAction `json:"csaction"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Actions, nil
}

func (s *CSService) GetCSAction(name string) ([]models.CSAction, error) {
	reqURL := fmt.Sprintf("%s/%s", csActionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Actions []models.CSAction `json:"csaction"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Actions, nil
}

func (s *CSService) CountCSAction() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, csActionURL+"?count=yes", nil)
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
		} `json:"csaction"`
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

// csparameter
func (s *CSService) UpdateCSParameter(param models.CSParameter) error {
	payload := map[string]any{"csparameter": param}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, csParameterURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) UnsetCSParameter(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"csparameter": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, csParameterURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) GetAllCSParameter() (models.CSParameter, error) {
	req, err := s.client.NewRequest(http.MethodGet, csParameterURL, nil)
	if err != nil {
		return models.CSParameter{}, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return models.CSParameter{}, err
	}
	var result struct {
		Params []models.CSParameter `json:"csparameter"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.CSParameter{}, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Params) > 0 {
		return result.Params[0], nil
	}
	return models.CSParameter{}, nil
}

// cspolicy
func (s *CSService) AddCSPolicy(policy models.CSPolicy) error {
	payload := map[string]any{"cspolicy": policy}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, csPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) DeleteCSPolicy(policyName string) error {
	reqURL := fmt.Sprintf("%s/%s", csPolicyURL, url.QueryEscape(policyName))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) UpdateCSPolicy(policy models.CSPolicy) error {
	payload := map[string]any{"cspolicy": policy}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, csPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) UnsetCSPolicy(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"cspolicy": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, csPolicyURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) RenameCSPolicy(policyName string, newName string) error {
	payload := map[string]any{
		"cspolicy": map[string]string{
			"policyname": policyName,
			"newname":    newName,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, csPolicyURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) GetAllCSPolicy() ([]models.CSPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, csPolicyURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Policies []models.CSPolicy `json:"cspolicy"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Policies, nil
}

func (s *CSService) GetCSPolicy(policyName string) ([]models.CSPolicy, error) {
	reqURL := fmt.Sprintf("%s/%s", csPolicyURL, url.QueryEscape(policyName))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Policies []models.CSPolicy `json:"cspolicy"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Policies, nil
}

func (s *CSService) CountCSPolicy() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, csPolicyURL+"?count=yes", nil)
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
		} `json:"cspolicy"`
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

// cspolicylabel
func (s *CSService) AddCSPolicyLabel(label models.CSPolicyLabel) error {
	payload := map[string]any{"cspolicylabel": label}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, csPolicyLabelURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) DeleteCSPolicyLabel(labelName string) error {
	reqURL := fmt.Sprintf("%s/%s", csPolicyLabelURL, url.QueryEscape(labelName))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) RenameCSPolicyLabel(labelName string, newName string) error {
	payload := map[string]any{
		"cspolicylabel": map[string]string{
			"labelname": labelName,
			"newname":   newName,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, csPolicyLabelURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) GetAllCSPolicyLabel() ([]models.CSPolicyLabel, error) {
	req, err := s.client.NewRequest(http.MethodGet, csPolicyLabelURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Labels []models.CSPolicyLabel `json:"cspolicylabel"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Labels, nil
}

func (s *CSService) GetCSPolicyLabel(labelName string) ([]models.CSPolicyLabel, error) {
	reqURL := fmt.Sprintf("%s/%s", csPolicyLabelURL, url.QueryEscape(labelName))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Labels []models.CSPolicyLabel `json:"cspolicylabel"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Labels, nil
}

func (s *CSService) CountCSPolicyLabel() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, csPolicyLabelURL+"?count=yes", nil)
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
		} `json:"cspolicylabel"`
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

// cspolicylabel_binding
func (s *CSService) GetAllCSPolicyLabelBinding() ([]models.CSPolicyLabelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, csPolicyLabelBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSPolicyLabelBinding `json:"cspolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) GetCSPolicyLabelBinding(labelName string) ([]models.CSPolicyLabelBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", csPolicyLabelBindingURL, url.QueryEscape(labelName))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSPolicyLabelBinding `json:"cspolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// cspolicylabel_cspolicy_binding
func (s *CSService) AddCSPolicyLabelCSPolicyBinding(binding models.CSPolicyLabelCSPolicyBinding) error {
	payload := map[string]any{"cspolicylabel_cspolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, csPolicyLabelCSPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) DeleteCSPolicyLabelCSPolicyBinding(labelName string, policyName string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s,priority:%d", csPolicyLabelCSPolicyBindingURL, url.QueryEscape(labelName), url.QueryEscape(policyName), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) GetAllCSPolicyLabelCSPolicyBinding() ([]models.CSPolicyLabelCSPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, csPolicyLabelCSPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSPolicyLabelCSPolicyBinding `json:"cspolicylabel_cspolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) GetCSPolicyLabelCSPolicyBinding(labelName string) ([]models.CSPolicyLabelCSPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", csPolicyLabelCSPolicyBindingURL, url.QueryEscape(labelName))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSPolicyLabelCSPolicyBinding `json:"cspolicylabel_cspolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) CountCSPolicyLabelCSPolicyBinding(labelName string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", csPolicyLabelCSPolicyBindingURL, url.QueryEscape(labelName))
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
		} `json:"cspolicylabel_cspolicy_binding"`
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

// cspolicy_binding
func (s *CSService) GetAllCSPolicyBinding() ([]models.CSPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, csPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSPolicyBinding `json:"cspolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) GetCSPolicyBinding(policyName string) ([]models.CSPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", csPolicyBindingURL, url.QueryEscape(policyName))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSPolicyBinding `json:"cspolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// cspolicy_crvserver_binding
func (s *CSService) GetAllCSPolicyCRVServerBinding() ([]models.CSPolicyCRVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, csPolicyCRVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSPolicyCRVServerBinding `json:"cspolicy_crvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) GetCSPolicyCRVServerBinding(policyName string) ([]models.CSPolicyCRVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", csPolicyCRVServerBindingURL, url.QueryEscape(policyName))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSPolicyCRVServerBinding `json:"cspolicy_crvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) CountCSPolicyCRVServerBinding(policyName string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", csPolicyCRVServerBindingURL, url.QueryEscape(policyName))
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
		} `json:"cspolicy_crvserver_binding"`
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

// cspolicy_cspolicylabel_binding
func (s *CSService) GetAllCSPolicyCSPolicyLabelBinding() ([]models.CSPolicyCSPolicyLabelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, csPolicyCSPolicyLabelBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSPolicyCSPolicyLabelBinding `json:"cspolicy_cspolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) GetCSPolicyCSPolicyLabelBinding(policyName string) ([]models.CSPolicyCSPolicyLabelBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", csPolicyCSPolicyLabelBindingURL, url.QueryEscape(policyName))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSPolicyCSPolicyLabelBinding `json:"cspolicy_cspolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) CountCSPolicyCSPolicyLabelBinding(policyName string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", csPolicyCSPolicyLabelBindingURL, url.QueryEscape(policyName))
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
		} `json:"cspolicy_cspolicylabel_binding"`
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

// cspolicy_csvserver_binding
func (s *CSService) GetAllCSPolicyCSVServerBinding() ([]models.CSPolicyCSVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, csPolicyCSVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSPolicyCSVServerBinding `json:"cspolicy_csvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) GetCSPolicyCSVServerBinding(policyName string) ([]models.CSPolicyCSVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", csPolicyCSVServerBindingURL, url.QueryEscape(policyName))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSPolicyCSVServerBinding `json:"cspolicy_csvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) CountCSPolicyCSVServerBinding(policyName string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", csPolicyCSVServerBindingURL, url.QueryEscape(policyName))
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
		} `json:"cspolicy_csvserver_binding"`
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

// csvserver
func (s *CSService) AddCSVServer(vserver models.CSVServer) error {
	payload := map[string]any{"csvserver": vserver}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, csVServerURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) DeleteCSVServer(name string) error {
	reqURL := fmt.Sprintf("%s/%s", csVServerURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) UpdateCSVServer(vserver models.CSVServer) error {
	payload := map[string]any{"csvserver": vserver}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, csVServerURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) UnsetCSVServer(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"csvserver": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, csVServerURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) EnableCSVServer(name string) error {
	payload := map[string]any{
		"csvserver": map[string]string{"name": name},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, csVServerURL+"?action=enable", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) DisableCSVServer(name string) error {
	payload := map[string]any{
		"csvserver": map[string]string{"name": name},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, csVServerURL+"?action=disable", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) RenameCSVServer(name string, newName string) error {
	payload := map[string]any{
		"csvserver": map[string]string{
			"name":    name,
			"newname": newName,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, csVServerURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) GetAllCSVServer() ([]models.CSVServer, error) {
	req, err := s.client.NewRequest(http.MethodGet, csVServerURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		VServers []models.CSVServer `json:"csvserver"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.VServers, nil
}

func (s *CSService) GetCSVServer(name string) ([]models.CSVServer, error) {
	reqURL := fmt.Sprintf("%s/%s", csVServerURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		VServers []models.CSVServer `json:"csvserver"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.VServers, nil
}

func (s *CSService) CountCSVServer() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, csVServerURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		VServers []struct {
			Count float64 `json:"__count"`
		} `json:"csvserver"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.VServers) > 0 {
		return result.VServers[0].Count, nil
	}
	return 0, nil
}

// csvserver_analyticsprofile_binding
func (s *CSService) AddCSVServerAnalyticsProfileBinding(binding models.CSVServerAnalyticsProfileBinding) error {
	payload := map[string]any{"csvserver_analyticsprofile_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, csVServerAnalyticsProfileBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) DeleteCSVServerAnalyticsProfileBinding(name string, analyticsProfile string) error {
	reqURL := fmt.Sprintf("%s/%s?args=analyticsprofile:%s", csVServerAnalyticsProfileBindingURL, url.QueryEscape(name), url.QueryEscape(analyticsProfile))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) GetAllCSVServerAnalyticsProfileBinding() ([]models.CSVServerAnalyticsProfileBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, csVServerAnalyticsProfileBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerAnalyticsProfileBinding `json:"csvserver_analyticsprofile_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) GetCSVServerAnalyticsProfileBinding(name string) ([]models.CSVServerAnalyticsProfileBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", csVServerAnalyticsProfileBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerAnalyticsProfileBinding `json:"csvserver_analyticsprofile_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) CountCSVServerAnalyticsProfileBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", csVServerAnalyticsProfileBindingURL, url.QueryEscape(name))
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
		} `json:"csvserver_analyticsprofile_binding"`
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

// csvserver_appflowpolicy_binding
func (s *CSService) AddCSVServerAppFlowPolicyBinding(binding models.CSVServerAppFlowPolicyBinding) error {
	payload := map[string]any{"csvserver_appflowpolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, csVServerAppFlowPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) DeleteCSVServerAppFlowPolicyBinding(name string, policyName string, bindPoint string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s,bindpoint:%s,priority:%d", csVServerAppFlowPolicyBindingURL, url.QueryEscape(name), url.QueryEscape(policyName), url.QueryEscape(bindPoint), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) GetAllCSVServerAppFlowPolicyBinding() ([]models.CSVServerAppFlowPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, csVServerAppFlowPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerAppFlowPolicyBinding `json:"csvserver_appflowpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) GetCSVServerAppFlowPolicyBinding(name string) ([]models.CSVServerAppFlowPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", csVServerAppFlowPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerAppFlowPolicyBinding `json:"csvserver_appflowpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) CountCSVServerAppFlowPolicyBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", csVServerAppFlowPolicyBindingURL, url.QueryEscape(name))
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
		} `json:"csvserver_appflowpolicy_binding"`
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

// csvserver_appfwpolicy_binding
func (s *CSService) AddCSVServerAppFWPolicyBinding(binding models.CSVServerAppFWPolicyBinding) error {
	payload := map[string]any{"csvserver_appfwpolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, csVServerAppFWPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) DeleteCSVServerAppFWPolicyBinding(name string, policyName string, bindPoint string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s,bindpoint:%s,priority:%d", csVServerAppFWPolicyBindingURL, url.QueryEscape(name), url.QueryEscape(policyName), url.QueryEscape(bindPoint), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) GetAllCSVServerAppFWPolicyBinding() ([]models.CSVServerAppFWPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, csVServerAppFWPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerAppFWPolicyBinding `json:"csvserver_appfwpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) GetCSVServerAppFWPolicyBinding(name string) ([]models.CSVServerAppFWPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", csVServerAppFWPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerAppFWPolicyBinding `json:"csvserver_appfwpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) CountCSVServerAppFWPolicyBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", csVServerAppFWPolicyBindingURL, url.QueryEscape(name))
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
		} `json:"csvserver_appfwpolicy_binding"`
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

// csvserver_appqoepolicy_binding
func (s *CSService) AddCSVServerAppQOEPolicyBinding(binding models.CSVServerAppQOEPolicyBinding) error {
	payload := map[string]any{"csvserver_appqoepolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, csVServerAppQOEPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) DeleteCSVServerAppQOEPolicyBinding(name string, policyName string, bindPoint string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s,bindpoint:%s,priority:%d", csVServerAppQOEPolicyBindingURL, url.QueryEscape(name), url.QueryEscape(policyName), url.QueryEscape(bindPoint), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) GetAllCSVServerAppQOEPolicyBinding() ([]models.CSVServerAppQOEPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, csVServerAppQOEPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerAppQOEPolicyBinding `json:"csvserver_appqoepolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) GetCSVServerAppQOEPolicyBinding(name string) ([]models.CSVServerAppQOEPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", csVServerAppQOEPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerAppQOEPolicyBinding `json:"csvserver_appqoepolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) CountCSVServerAppQOEPolicyBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", csVServerAppQOEPolicyBindingURL, url.QueryEscape(name))
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
		} `json:"csvserver_appqoepolicy_binding"`
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

// csvserver_auditnslogpolicy_binding
func (s *CSService) AddCSVServerAuditNSLogPolicyBinding(binding models.CSVServerAuditNSLogPolicyBinding) error {
	payload := map[string]any{"csvserver_auditnslogpolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, csVServerAuditNSLogPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) DeleteCSVServerAuditNSLogPolicyBinding(name string, policyName string, bindPoint string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s,bindpoint:%s,priority:%d", csVServerAuditNSLogPolicyBindingURL, url.QueryEscape(name), url.QueryEscape(policyName), url.QueryEscape(bindPoint), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) GetAllCSVServerAuditNSLogPolicyBinding() ([]models.CSVServerAuditNSLogPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, csVServerAuditNSLogPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerAuditNSLogPolicyBinding `json:"csvserver_auditnslogpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) GetCSVServerAuditNSLogPolicyBinding(name string) ([]models.CSVServerAuditNSLogPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", csVServerAuditNSLogPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerAuditNSLogPolicyBinding `json:"csvserver_auditnslogpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) CountCSVServerAuditNSLogPolicyBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", csVServerAuditNSLogPolicyBindingURL, url.QueryEscape(name))
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
		} `json:"csvserver_auditnslogpolicy_binding"`
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

// csvserver_auditsyslogpolicy_binding
func (s *CSService) AddCSVServerAuditSyslogPolicyBinding(binding models.CSVServerAuditSyslogPolicyBinding) error {
	payload := map[string]any{"csvserver_auditsyslogpolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, csVServerAuditSyslogPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) DeleteCSVServerAuditSyslogPolicyBinding(name string, policyName string, bindPoint string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s,bindpoint:%s,priority:%d", csVServerAuditSyslogPolicyBindingURL, url.QueryEscape(name), url.QueryEscape(policyName), url.QueryEscape(bindPoint), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) GetAllCSVServerAuditSyslogPolicyBinding() ([]models.CSVServerAuditSyslogPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, csVServerAuditSyslogPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerAuditSyslogPolicyBinding `json:"csvserver_auditsyslogpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) GetCSVServerAuditSyslogPolicyBinding(name string) ([]models.CSVServerAuditSyslogPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", csVServerAuditSyslogPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerAuditSyslogPolicyBinding `json:"csvserver_auditsyslogpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) CountCSVServerAuditSyslogPolicyBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", csVServerAuditSyslogPolicyBindingURL, url.QueryEscape(name))
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
		} `json:"csvserver_auditsyslogpolicy_binding"`
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

// csvserver_authorizationpolicy_binding
func (s *CSService) AddCSVServerAuthorizationPolicyBinding(binding models.CSVServerAuthorizationPolicyBinding) error {
	payload := map[string]any{"csvserver_authorizationpolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, csVServerAuthorizationPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) DeleteCSVServerAuthorizationPolicyBinding(name string, policyName string, bindPoint string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s,bindpoint:%s,priority:%d", csVServerAuthorizationPolicyBindingURL, url.QueryEscape(name), url.QueryEscape(policyName), url.QueryEscape(bindPoint), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) GetAllCSVServerAuthorizationPolicyBinding() ([]models.CSVServerAuthorizationPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, csVServerAuthorizationPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerAuthorizationPolicyBinding `json:"csvserver_authorizationpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) GetCSVServerAuthorizationPolicyBinding(name string) ([]models.CSVServerAuthorizationPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", csVServerAuthorizationPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerAuthorizationPolicyBinding `json:"csvserver_authorizationpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) CountCSVServerAuthorizationPolicyBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", csVServerAuthorizationPolicyBindingURL, url.QueryEscape(name))
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
		} `json:"csvserver_authorizationpolicy_binding"`
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

// csvserver_binding
func (s *CSService) GetAllCSVServerBinding() ([]models.CSVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, csVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerBinding `json:"csvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) GetCSVServerBinding(name string) ([]models.CSVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", csVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerBinding `json:"csvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// csvserver_botpolicy_binding
func (s *CSService) AddCSVServerBotPolicyBinding(binding models.CSVServerBotPolicyBinding) error {
	payload := map[string]any{"csvserver_botpolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, csVServerBotPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) DeleteCSVServerBotPolicyBinding(name string, policyName string, bindPoint string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s,bindpoint:%s,priority:%d", csVServerBotPolicyBindingURL, url.QueryEscape(name), url.QueryEscape(policyName), url.QueryEscape(bindPoint), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) GetAllCSVServerBotPolicyBinding() ([]models.CSVServerBotPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, csVServerBotPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerBotPolicyBinding `json:"csvserver_botpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) GetCSVServerBotPolicyBinding(name string) ([]models.CSVServerBotPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", csVServerBotPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerBotPolicyBinding `json:"csvserver_botpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) CountCSVServerBotPolicyBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", csVServerBotPolicyBindingURL, url.QueryEscape(name))
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
		} `json:"csvserver_botpolicy_binding"`
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

// csvserver_cachepolicy_binding
func (s *CSService) AddCSVServerCachePolicyBinding(binding models.CSVServerCachePolicyBinding) error {
	payload := map[string]any{"csvserver_cachepolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, csVServerCachePolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) DeleteCSVServerCachePolicyBinding(name string, policyName string, bindPoint string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s,bindpoint:%s,priority:%d", csVServerCachePolicyBindingURL, url.QueryEscape(name), url.QueryEscape(policyName), url.QueryEscape(bindPoint), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) GetAllCSVServerCachePolicyBinding() ([]models.CSVServerCachePolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, csVServerCachePolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerCachePolicyBinding `json:"csvserver_cachepolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) GetCSVServerCachePolicyBinding(name string) ([]models.CSVServerCachePolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", csVServerCachePolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerCachePolicyBinding `json:"csvserver_cachepolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) CountCSVServerCachePolicyBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", csVServerCachePolicyBindingURL, url.QueryEscape(name))
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
		} `json:"csvserver_cachepolicy_binding"`
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

// csvserver_cmppolicy_binding
func (s *CSService) AddCSVServerCMPPolicyBinding(binding models.CSVServerCMPPolicyBinding) error {
	payload := map[string]any{"csvserver_cmppolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, csVServerCMPPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) DeleteCSVServerCMPPolicyBinding(name string, policyName string, bindPoint string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s,bindpoint:%s,priority:%d", csVServerCMPPolicyBindingURL, url.QueryEscape(name), url.QueryEscape(policyName), url.QueryEscape(bindPoint), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) GetAllCSVServerCMPPolicyBinding() ([]models.CSVServerCMPPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, csVServerCMPPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerCMPPolicyBinding `json:"csvserver_cmppolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) GetCSVServerCMPPolicyBinding(name string) ([]models.CSVServerCMPPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", csVServerCMPPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerCMPPolicyBinding `json:"csvserver_cmppolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) CountCSVServerCMPPolicyBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", csVServerCMPPolicyBindingURL, url.QueryEscape(name))
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
		} `json:"csvserver_cmppolicy_binding"`
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

// csvserver_contentinspectionpolicy_binding
func (s *CSService) AddCSVServerContentInspectionPolicyBinding(binding models.CSVServerContentInspectionPolicyBinding) error {
	payload := map[string]any{"csvserver_contentinspectionpolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, csVServerContentInspectionPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) DeleteCSVServerContentInspectionPolicyBinding(name string, policyName string, bindPoint string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s,bindpoint:%s,priority:%d", csVServerContentInspectionPolicyBindingURL, url.QueryEscape(name), url.QueryEscape(policyName), url.QueryEscape(bindPoint), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) GetAllCSVServerContentInspectionPolicyBinding() ([]models.CSVServerContentInspectionPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, csVServerContentInspectionPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerContentInspectionPolicyBinding `json:"csvserver_contentinspectionpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) GetCSVServerContentInspectionPolicyBinding(name string) ([]models.CSVServerContentInspectionPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", csVServerContentInspectionPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerContentInspectionPolicyBinding `json:"csvserver_contentinspectionpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) CountCSVServerContentInspectionPolicyBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", csVServerContentInspectionPolicyBindingURL, url.QueryEscape(name))
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
		} `json:"csvserver_contentinspectionpolicy_binding"`
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

// csvserver_cspolicy_binding
func (s *CSService) AddCSVServerCSPolicyBinding(binding models.CSVServerCSPolicyBinding) error {
	payload := map[string]any{"csvserver_cspolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, csVServerCSPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) DeleteCSVServerCSPolicyBinding(name string, policyName string, bindPoint string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s,bindpoint:%s,priority:%d", csVServerCSPolicyBindingURL, url.QueryEscape(name), url.QueryEscape(policyName), url.QueryEscape(bindPoint), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) GetAllCSVServerCSPolicyBinding() ([]models.CSVServerCSPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, csVServerCSPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerCSPolicyBinding `json:"csvserver_cspolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) GetCSVServerCSPolicyBinding(name string) ([]models.CSVServerCSPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", csVServerCSPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerCSPolicyBinding `json:"csvserver_cspolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) CountCSVServerCSPolicyBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", csVServerCSPolicyBindingURL, url.QueryEscape(name))
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
		} `json:"csvserver_cspolicy_binding"`
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

// csvserver_domain_binding
func (s *CSService) AddCSVServerDomainBinding(binding models.CSVServerDomainBinding) error {
	payload := map[string]any{"csvserver_domain_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, csVServerDomainBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) DeleteCSVServerDomainBinding(name string, domainName string) error {
	reqURL := fmt.Sprintf("%s/%s?args=domainname:%s", csVServerDomainBindingURL, url.QueryEscape(name), url.QueryEscape(domainName))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) GetAllCSVServerDomainBinding() ([]models.CSVServerDomainBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, csVServerDomainBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerDomainBinding `json:"csvserver_domain_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) GetCSVServerDomainBinding(name string) ([]models.CSVServerDomainBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", csVServerDomainBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerDomainBinding `json:"csvserver_domain_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) CountCSVServerDomainBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", csVServerDomainBindingURL, url.QueryEscape(name))
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
		} `json:"csvserver_domain_binding"`
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

// csvserver_feopolicy_binding
func (s *CSService) AddCSVServerFEOPolicyBinding(binding models.CSVServerFEOPolicyBinding) error {
	payload := map[string]any{"csvserver_feopolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, csVServerFEOPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) DeleteCSVServerFEOPolicyBinding(name string, policyName string, bindPoint string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s,bindpoint:%s,priority:%d", csVServerFEOPolicyBindingURL, url.QueryEscape(name), url.QueryEscape(policyName), url.QueryEscape(bindPoint), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) GetAllCSVServerFEOPolicyBinding() ([]models.CSVServerFEOPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, csVServerFEOPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerFEOPolicyBinding `json:"csvserver_feopolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) GetCSVServerFEOPolicyBinding(name string) ([]models.CSVServerFEOPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", csVServerFEOPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerFEOPolicyBinding `json:"csvserver_feopolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) CountCSVServerFEOPolicyBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", csVServerFEOPolicyBindingURL, url.QueryEscape(name))
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
		} `json:"csvserver_feopolicy_binding"`
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

// csvserver_gslbvserver_binding
func (s *CSService) AddCSVServerGSLBVServerBinding(binding models.CSVServerGSLBVServerBinding) error {
	payload := map[string]any{"csvserver_gslbvserver_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, csVServerGSLBVServerBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) DeleteCSVServerGSLBVServerBinding(name string, vserverName string) error {
	reqURL := fmt.Sprintf("%s/%s?args=vserver:%s", csVServerGSLBVServerBindingURL, url.QueryEscape(name), url.QueryEscape(vserverName))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) GetAllCSVServerGSLBVServerBinding() ([]models.CSVServerGSLBVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, csVServerGSLBVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerGSLBVServerBinding `json:"csvserver_gslbvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) GetCSVServerGSLBVServerBinding(name string) ([]models.CSVServerGSLBVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", csVServerGSLBVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerGSLBVServerBinding `json:"csvserver_gslbvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) CountCSVServerGSLBVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", csVServerGSLBVServerBindingURL, url.QueryEscape(name))
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
		} `json:"csvserver_gslbvserver_binding"`
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

// csvserver_lbvserver_binding
func (s *CSService) AddCSVServerLBVServerBinding(binding models.CSVServerLBVServerBinding) error {
	payload := map[string]any{"csvserver_lbvserver_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, csVServerLBVServerBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) DeleteCSVServerLBVServerBinding(name string, lbvserver string) error {
	reqURL := fmt.Sprintf("%s/%s?args=lbvserver:%s", csVServerLBVServerBindingURL, url.QueryEscape(name), url.QueryEscape(lbvserver))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) GetAllCSVServerLBVServerBinding() ([]models.CSVServerLBVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, csVServerLBVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerLBVServerBinding `json:"csvserver_lbvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) GetCSVServerLBVServerBinding(name string) ([]models.CSVServerLBVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", csVServerLBVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerLBVServerBinding `json:"csvserver_lbvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) CountCSVServerLBVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", csVServerLBVServerBindingURL, url.QueryEscape(name))
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
		} `json:"csvserver_lbvserver_binding"`
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

// csvserver_responderpolicy_binding
func (s *CSService) AddCSVServerResponderPolicyBinding(binding models.CSVServerResponderPolicyBinding) error {
	payload := map[string]any{"csvserver_responderpolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, csVServerResponderPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) DeleteCSVServerResponderPolicyBinding(name string, policyName string, bindPoint string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s,bindpoint:%s,priority:%d", csVServerResponderPolicyBindingURL, url.QueryEscape(name), url.QueryEscape(policyName), url.QueryEscape(bindPoint), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) GetAllCSVServerResponderPolicyBinding() ([]models.CSVServerResponderPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, csVServerResponderPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerResponderPolicyBinding `json:"csvserver_responderpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) GetCSVServerResponderPolicyBinding(name string) ([]models.CSVServerResponderPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", csVServerResponderPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerResponderPolicyBinding `json:"csvserver_responderpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) CountCSVServerResponderPolicyBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", csVServerResponderPolicyBindingURL, url.QueryEscape(name))
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
		} `json:"csvserver_responderpolicy_binding"`
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

// csvserver_rewritepolicy_binding
func (s *CSService) AddCSVServerRewritePolicyBinding(binding models.CSVServerRewritePolicyBinding) error {
	payload := map[string]any{"csvserver_rewritepolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, csVServerRewritePolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) DeleteCSVServerRewritePolicyBinding(name string, policyName string, bindPoint string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s,bindpoint:%s,priority:%d", csVServerRewritePolicyBindingURL, url.QueryEscape(name), url.QueryEscape(policyName), url.QueryEscape(bindPoint), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) GetAllCSVServerRewritePolicyBinding() ([]models.CSVServerRewritePolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, csVServerRewritePolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerRewritePolicyBinding `json:"csvserver_rewritepolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) GetCSVServerRewritePolicyBinding(name string) ([]models.CSVServerRewritePolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", csVServerRewritePolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerRewritePolicyBinding `json:"csvserver_rewritepolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) CountCSVServerRewritePolicyBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", csVServerRewritePolicyBindingURL, url.QueryEscape(name))
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
		} `json:"csvserver_rewritepolicy_binding"`
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

// csvserver_spilloverpolicy_binding
func (s *CSService) AddCSVServerSpilloverPolicyBinding(binding models.CSVServerSpilloverPolicyBinding) error {
	payload := map[string]any{"csvserver_spilloverpolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, csVServerSpilloverPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) DeleteCSVServerSpilloverPolicyBinding(name string, policyName string, bindPoint string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s,bindpoint:%s,priority:%d", csVServerSpilloverPolicyBindingURL, url.QueryEscape(name), url.QueryEscape(policyName), url.QueryEscape(bindPoint), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) GetAllCSVServerSpilloverPolicyBinding() ([]models.CSVServerSpilloverPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, csVServerSpilloverPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerSpilloverPolicyBinding `json:"csvserver_spilloverpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) GetCSVServerSpilloverPolicyBinding(name string) ([]models.CSVServerSpilloverPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", csVServerSpilloverPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerSpilloverPolicyBinding `json:"csvserver_spilloverpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) CountCSVServerSpilloverPolicyBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", csVServerSpilloverPolicyBindingURL, url.QueryEscape(name))
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
		} `json:"csvserver_spilloverpolicy_binding"`
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

// csvserver_tmtrafficpolicy_binding
func (s *CSService) AddCSVServerTMTrafficPolicyBinding(binding models.CSVServerTMTrafficPolicyBinding) error {
	payload := map[string]any{"csvserver_tmtrafficpolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, csVServerTMTrafficPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) DeleteCSVServerTMTrafficPolicyBinding(name string, policyName string, bindPoint string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s,bindpoint:%s,priority:%d", csVServerTMTrafficPolicyBindingURL, url.QueryEscape(name), url.QueryEscape(policyName), url.QueryEscape(bindPoint), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) GetAllCSVServerTMTrafficPolicyBinding() ([]models.CSVServerTMTrafficPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, csVServerTMTrafficPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerTMTrafficPolicyBinding `json:"csvserver_tmtrafficpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) GetCSVServerTMTrafficPolicyBinding(name string) ([]models.CSVServerTMTrafficPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", csVServerTMTrafficPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerTMTrafficPolicyBinding `json:"csvserver_tmtrafficpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) CountCSVServerTMTrafficPolicyBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", csVServerTMTrafficPolicyBindingURL, url.QueryEscape(name))
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
		} `json:"csvserver_tmtrafficpolicy_binding"`
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

// csvserver_transformpolicy_binding
func (s *CSService) AddCSVServerTransformPolicyBinding(binding models.CSVServerTransformPolicyBinding) error {
	payload := map[string]any{"csvserver_transformpolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, csVServerTransformPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) DeleteCSVServerTransformPolicyBinding(name string, policyName string, bindPoint string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s,bindpoint:%s,priority:%d", csVServerTransformPolicyBindingURL, url.QueryEscape(name), url.QueryEscape(policyName), url.QueryEscape(bindPoint), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) GetAllCSVServerTransformPolicyBinding() ([]models.CSVServerTransformPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, csVServerTransformPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerTransformPolicyBinding `json:"csvserver_transformpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) GetCSVServerTransformPolicyBinding(name string) ([]models.CSVServerTransformPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", csVServerTransformPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerTransformPolicyBinding `json:"csvserver_transformpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) CountCSVServerTransformPolicyBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", csVServerTransformPolicyBindingURL, url.QueryEscape(name))
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
		} `json:"csvserver_transformpolicy_binding"`
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

// csvserver_vpnvserver_binding
func (s *CSService) AddCSVServerVPNVServerBinding(binding models.CSVServerVPNVServerBinding) error {
	payload := map[string]any{"csvserver_vpnvserver_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, csVServerVPNVServerBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) DeleteCSVServerVPNVServerBinding(name string, vserver string) error {
	reqURL := fmt.Sprintf("%s/%s?args=vserver:%s", csVServerVPNVServerBindingURL, url.QueryEscape(name), url.QueryEscape(vserver))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CSService) GetAllCSVServerVPNVServerBinding() ([]models.CSVServerVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, csVServerVPNVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerVPNVServerBinding `json:"csvserver_vpnvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) GetCSVServerVPNVServerBinding(name string) ([]models.CSVServerVPNVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", csVServerVPNVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CSVServerVPNVServerBinding `json:"csvserver_vpnvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CSService) CountCSVServerVPNVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", csVServerVPNVServerBindingURL, url.QueryEscape(name))
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
		} `json:"csvserver_vpnvserver_binding"`
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
