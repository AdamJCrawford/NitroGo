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
	crActionURL                         = "/nitro/v1/config/craction"
	crPolicyURL                         = "/nitro/v1/config/crpolicy"
	crPolicyBindingURL                  = "/nitro/v1/config/crpolicy_binding"
	crPolicyCRVServerBindingURL         = "/nitro/v1/config/crpolicy_crvserver_binding"
	crVServerURL                        = "/nitro/v1/config/crvserver"
	crVServerAnalyticsProfileBindingURL = "/nitro/v1/config/crvserver_analyticsprofile_binding"
	crVServerAppFlowPolicyBindingURL    = "/nitro/v1/config/crvserver_appflowpolicy_binding"
	crVServerAppFWPolicyBindingURL      = "/nitro/v1/config/crvserver_appfwpolicy_binding"
	crVServerAppQOEPolicyBindingURL     = "/nitro/v1/config/crvserver_appqoepolicy_binding"
	crVServerBindingURL                 = "/nitro/v1/config/crvserver_binding"
	crVServerCachePolicyBindingURL      = "/nitro/v1/config/crvserver_cachepolicy_binding"
	crVServerCMPPolicyBindingURL        = "/nitro/v1/config/crvserver_cmppolicy_binding"
	crVServerCRPolicyBindingURL         = "/nitro/v1/config/crvserver_crpolicy_binding"
	crVServerCSPolicyBindingURL         = "/nitro/v1/config/crvserver_cspolicy_binding"
	crVServerFEOPolicyBindingURL        = "/nitro/v1/config/crvserver_feopolicy_binding"
	crVServerICAPolicyBindingURL        = "/nitro/v1/config/crvserver_icapolicy_binding"
	crVServerLBVServerBindingURL        = "/nitro/v1/config/crvserver_lbvserver_binding"
	crVServerPolicyMapBindingURL        = "/nitro/v1/config/crvserver_policymap_binding"
	crVServerResponderPolicyBindingURL  = "/nitro/v1/config/crvserver_responderpolicy_binding"
	crVServerRewritePolicyBindingURL    = "/nitro/v1/config/crvserver_rewritepolicy_binding"
	crVServerSpilloverPolicyBindingURL  = "/nitro/v1/config/crvserver_spilloverpolicy_binding"
)

// Cache redirection configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/cr/cr
type CRService struct {
	client *Client
}

// craction
func (s *CRService) GetAllCRAction() ([]models.CRAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, crActionURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Actions []models.CRAction `json:"craction"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Actions, nil
}

func (s *CRService) GetCRAction(name string) ([]models.CRAction, error) {
	reqURL := fmt.Sprintf("%s/%s", crActionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Actions []models.CRAction `json:"craction"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Actions, nil
}

func (s *CRService) CountCRAction() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, crActionURL+"?count=yes", nil)
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
		} `json:"craction"`
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

// crpolicy
func (s *CRService) AddCRPolicy(policy models.CRPolicy) error {
	payload := map[string]any{"crpolicy": policy}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, crPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) DeleteCRPolicy(policyName string) error {
	reqURL := fmt.Sprintf("%s/%s", crPolicyURL, url.QueryEscape(policyName))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) UpdateCRPolicy(policy models.CRPolicy) error {
	payload := map[string]any{"crpolicy": policy}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, crPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) UnsetCRPolicy(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"crpolicy": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, crPolicyURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) RenameCRPolicy(name string, newName string) error {
	payload := map[string]any{
		"crpolicy": map[string]string{
			"policyname": name,
			"newname":    newName,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, crPolicyURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) GetAllCRPolicy() ([]models.CRPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, crPolicyURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Policies []models.CRPolicy `json:"crpolicy"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Policies, nil
}

func (s *CRService) GetCRPolicy(policyName string) ([]models.CRPolicy, error) {
	reqURL := fmt.Sprintf("%s/%s", crPolicyURL, url.QueryEscape(policyName))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Policies []models.CRPolicy `json:"crpolicy"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Policies, nil
}

func (s *CRService) CountCRPolicy() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, crPolicyURL+"?count=yes", nil)
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
		} `json:"crpolicy"`
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

// crpolicy_binding
func (s *CRService) GetAllCRPolicyBinding() ([]models.CRPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, crPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CRPolicyBinding `json:"crpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CRService) GetCRPolicyBinding(policyName string) ([]models.CRPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", crPolicyBindingURL, url.QueryEscape(policyName))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CRPolicyBinding `json:"crpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// crpolicy_crvserver_binding
func (s *CRService) GetAllCRPolicyCRVServerBinding() ([]models.CRPolicyCRVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, crPolicyCRVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CRPolicyCRVServerBinding `json:"crpolicy_crvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CRService) GetCRPolicyCRVServerBinding(policyName string) ([]models.CRPolicyCRVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", crPolicyCRVServerBindingURL, url.QueryEscape(policyName))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CRPolicyCRVServerBinding `json:"crpolicy_crvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CRService) CountCRPolicyCRVServerBinding(policyName string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", crPolicyCRVServerBindingURL, url.QueryEscape(policyName))
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
		} `json:"crpolicy_crvserver_binding"`
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

// crvserver
func (s *CRService) AddCRVServer(vserver models.CRVServer) error {
	payload := map[string]any{"crvserver": vserver}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, crVServerURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) DeleteCRVServer(name string) error {
	reqURL := fmt.Sprintf("%s/%s", crVServerURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) UpdateCRVServer(vserver models.CRVServer) error {
	payload := map[string]any{"crvserver": vserver}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, crVServerURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) UnsetCRVServer(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"crvserver": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, crVServerURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) EnableCRVServer(name string) error {
	payload := map[string]any{
		"crvserver": map[string]string{"name": name},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, crVServerURL+"?action=enable", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) DisableCRVServer(name string) error {
	payload := map[string]any{
		"crvserver": map[string]string{"name": name},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, crVServerURL+"?action=disable", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) RenameCRVServer(name string, newName string) error {
	payload := map[string]any{
		"crvserver": map[string]string{
			"name":    name,
			"newname": newName,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, crVServerURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) GetAllCRVServer() ([]models.CRVServer, error) {
	req, err := s.client.NewRequest(http.MethodGet, crVServerURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		VServers []models.CRVServer `json:"crvserver"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.VServers, nil
}

func (s *CRService) GetCRVServer(name string) ([]models.CRVServer, error) {
	reqURL := fmt.Sprintf("%s/%s", crVServerURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		VServers []models.CRVServer `json:"crvserver"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.VServers, nil
}

func (s *CRService) CountCRVServer() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, crVServerURL+"?count=yes", nil)
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
		} `json:"crvserver"`
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

// crvserver_analyticsprofile_binding
func (s *CRService) AddCRVServerAnalyticsProfileBinding(binding models.CRVServerAnalyticsProfileBinding) error {
	payload := map[string]any{"crvserver_analyticsprofile_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, crVServerAnalyticsProfileBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) DeleteCRVServerAnalyticsProfileBinding(name string, analyticsProfile string) error {
	reqURL := fmt.Sprintf("%s/%s?args=analyticsprofile:%s", crVServerAnalyticsProfileBindingURL, url.QueryEscape(name), url.QueryEscape(analyticsProfile))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) GetAllCRVServerAnalyticsProfileBinding() ([]models.CRVServerAnalyticsProfileBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, crVServerAnalyticsProfileBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CRVServerAnalyticsProfileBinding `json:"crvserver_analyticsprofile_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CRService) GetCRVServerAnalyticsProfileBinding(name string) ([]models.CRVServerAnalyticsProfileBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", crVServerAnalyticsProfileBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CRVServerAnalyticsProfileBinding `json:"crvserver_analyticsprofile_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CRService) CountCRVServerAnalyticsProfileBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", crVServerAnalyticsProfileBindingURL, url.QueryEscape(name))
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
		} `json:"crvserver_analyticsprofile_binding"`
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

// crvserver_appflowpolicy_binding
func (s *CRService) AddCRVServerAppFlowPolicyBinding(binding models.CRVServerAppFlowPolicyBinding) error {
	payload := map[string]any{"crvserver_appflowpolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, crVServerAppFlowPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) DeleteCRVServerAppFlowPolicyBinding(name string, policyName string, bindPoint string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s,bindpoint:%s,priority:%d", crVServerAppFlowPolicyBindingURL, url.QueryEscape(name), url.QueryEscape(policyName), url.QueryEscape(bindPoint), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) GetAllCRVServerAppFlowPolicyBinding() ([]models.CRVServerAppFlowPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, crVServerAppFlowPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CRVServerAppFlowPolicyBinding `json:"crvserver_appflowpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CRService) GetCRVServerAppFlowPolicyBinding(name string) ([]models.CRVServerAppFlowPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", crVServerAppFlowPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CRVServerAppFlowPolicyBinding `json:"crvserver_appflowpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CRService) CountCRVServerAppFlowPolicyBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", crVServerAppFlowPolicyBindingURL, url.QueryEscape(name))
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
		} `json:"crvserver_appflowpolicy_binding"`
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

// crvserver_appfwpolicy_binding
func (s *CRService) AddCRVServerAppFWPolicyBinding(binding models.CRVServerAppFWPolicyBinding) error {
	payload := map[string]any{"crvserver_appfwpolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, crVServerAppFWPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) DeleteCRVServerAppFWPolicyBinding(name string, policyName string, bindPoint string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s,bindpoint:%s,priority:%d", crVServerAppFWPolicyBindingURL, url.QueryEscape(name), url.QueryEscape(policyName), url.QueryEscape(bindPoint), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) GetAllCRVServerAppFWPolicyBinding() ([]models.CRVServerAppFWPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, crVServerAppFWPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CRVServerAppFWPolicyBinding `json:"crvserver_appfwpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CRService) GetCRVServerAppFWPolicyBinding(name string) ([]models.CRVServerAppFWPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", crVServerAppFWPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CRVServerAppFWPolicyBinding `json:"crvserver_appfwpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CRService) CountCRVServerAppFWPolicyBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", crVServerAppFWPolicyBindingURL, url.QueryEscape(name))
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
		} `json:"crvserver_appfwpolicy_binding"`
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

// crvserver_appqoepolicy_binding
func (s *CRService) AddCRVServerAppQOEPolicyBinding(binding models.CRVServerAppQOEPolicyBinding) error {
	payload := map[string]any{"crvserver_appqoepolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, crVServerAppQOEPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) DeleteCRVServerAppQOEPolicyBinding(name string, policyName string, bindPoint string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s,bindpoint:%s,priority:%d", crVServerAppQOEPolicyBindingURL, url.QueryEscape(name), url.QueryEscape(policyName), url.QueryEscape(bindPoint), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) GetAllCRVServerAppQOEPolicyBinding() ([]models.CRVServerAppQOEPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, crVServerAppQOEPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CRVServerAppQOEPolicyBinding `json:"crvserver_appqoepolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CRService) GetCRVServerAppQOEPolicyBinding(name string) ([]models.CRVServerAppQOEPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", crVServerAppQOEPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CRVServerAppQOEPolicyBinding `json:"crvserver_appqoepolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CRService) CountCRVServerAppQOEPolicyBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", crVServerAppQOEPolicyBindingURL, url.QueryEscape(name))
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
		} `json:"crvserver_appqoepolicy_binding"`
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

// crvserver_binding
func (s *CRService) GetAllCRVServerBinding() ([]models.CRVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, crVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CRVServerBinding `json:"crvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CRService) GetCRVServerBinding(name string) ([]models.CRVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", crVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CRVServerBinding `json:"crvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// crvserver_cachepolicy_binding
func (s *CRService) AddCRVServerCachePolicyBinding(binding models.CRVServerCachePolicyBinding) error {
	payload := map[string]any{"crvserver_cachepolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, crVServerCachePolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) DeleteCRVServerCachePolicyBinding(name string, policyName string, bindPoint string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s,bindpoint:%s,priority:%d", crVServerCachePolicyBindingURL, url.QueryEscape(name), url.QueryEscape(policyName), url.QueryEscape(bindPoint), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) GetAllCRVServerCachePolicyBinding() ([]models.CRVServerCachePolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, crVServerCachePolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CRVServerCachePolicyBinding `json:"crvserver_cachepolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CRService) GetCRVServerCachePolicyBinding(name string) ([]models.CRVServerCachePolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", crVServerCachePolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CRVServerCachePolicyBinding `json:"crvserver_cachepolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CRService) CountCRVServerCachePolicyBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", crVServerCachePolicyBindingURL, url.QueryEscape(name))
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
		} `json:"crvserver_cachepolicy_binding"`
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

// crvserver_cmppolicy_binding
func (s *CRService) AddCRVServerCMPPolicyBinding(binding models.CRVServerCMPPolicyBinding) error {
	payload := map[string]any{"crvserver_cmppolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, crVServerCMPPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) DeleteCRVServerCMPPolicyBinding(name string, policyName string, bindPoint string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s,bindpoint:%s,priority:%d", crVServerCMPPolicyBindingURL, url.QueryEscape(name), url.QueryEscape(policyName), url.QueryEscape(bindPoint), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) GetAllCRVServerCMPPolicyBinding() ([]models.CRVServerCMPPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, crVServerCMPPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CRVServerCMPPolicyBinding `json:"crvserver_cmppolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CRService) GetCRVServerCMPPolicyBinding(name string) ([]models.CRVServerCMPPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", crVServerCMPPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CRVServerCMPPolicyBinding `json:"crvserver_cmppolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CRService) CountCRVServerCMPPolicyBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", crVServerCMPPolicyBindingURL, url.QueryEscape(name))
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
		} `json:"crvserver_cmppolicy_binding"`
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

// crvserver_crpolicy_binding
func (s *CRService) AddCRVServerCRPolicyBinding(binding models.CRVServerCRPolicyBinding) error {
	payload := map[string]any{"crvserver_crpolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, crVServerCRPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) DeleteCRVServerCRPolicyBinding(name string, policyName string, bindPoint string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s,bindpoint:%s,priority:%d", crVServerCRPolicyBindingURL, url.QueryEscape(name), url.QueryEscape(policyName), url.QueryEscape(bindPoint), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) GetAllCRVServerCRPolicyBinding() ([]models.CRVServerCRPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, crVServerCRPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CRVServerCRPolicyBinding `json:"crvserver_crpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CRService) GetCRVServerCRPolicyBinding(name string) ([]models.CRVServerCRPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", crVServerCRPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CRVServerCRPolicyBinding `json:"crvserver_crpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CRService) CountCRVServerCRPolicyBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", crVServerCRPolicyBindingURL, url.QueryEscape(name))
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
		} `json:"crvserver_crpolicy_binding"`
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

// crvserver_cspolicy_binding
func (s *CRService) AddCRVServerCSPolicyBinding(binding models.CRVServerCSPolicyBinding) error {
	payload := map[string]any{"crvserver_cspolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, crVServerCSPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) DeleteCRVServerCSPolicyBinding(name string, policyName string, bindPoint string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s,bindpoint:%s,priority:%d", crVServerCSPolicyBindingURL, url.QueryEscape(name), url.QueryEscape(policyName), url.QueryEscape(bindPoint), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) GetAllCRVServerCSPolicyBinding() ([]models.CRVServerCSPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, crVServerCSPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CRVServerCSPolicyBinding `json:"crvserver_cspolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CRService) GetCRVServerCSPolicyBinding(name string) ([]models.CRVServerCSPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", crVServerCSPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CRVServerCSPolicyBinding `json:"crvserver_cspolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CRService) CountCRVServerCSPolicyBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", crVServerCSPolicyBindingURL, url.QueryEscape(name))
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
		} `json:"crvserver_cspolicy_binding"`
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

// crvserver_feopolicy_binding
func (s *CRService) AddCRVServerFEOPolicyBinding(binding models.CRVServerFEOPolicyBinding) error {
	payload := map[string]any{"crvserver_feopolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, crVServerFEOPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) DeleteCRVServerFEOPolicyBinding(name string, policyName string, bindPoint string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s,bindpoint:%s,priority:%d", crVServerFEOPolicyBindingURL, url.QueryEscape(name), url.QueryEscape(policyName), url.QueryEscape(bindPoint), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) GetAllCRVServerFEOPolicyBinding() ([]models.CRVServerFEOPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, crVServerFEOPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CRVServerFEOPolicyBinding `json:"crvserver_feopolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CRService) GetCRVServerFEOPolicyBinding(name string) ([]models.CRVServerFEOPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", crVServerFEOPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CRVServerFEOPolicyBinding `json:"crvserver_feopolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CRService) CountCRVServerFEOPolicyBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", crVServerFEOPolicyBindingURL, url.QueryEscape(name))
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
		} `json:"crvserver_feopolicy_binding"`
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

// crvserver_icapolicy_binding
func (s *CRService) AddCRVServerICAPolicyBinding(binding models.CRVServerICAPolicyBinding) error {
	payload := map[string]any{"crvserver_icapolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, crVServerICAPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) DeleteCRVServerICAPolicyBinding(name string, policyName string, bindPoint string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s,bindpoint:%s,priority:%d", crVServerICAPolicyBindingURL, url.QueryEscape(name), url.QueryEscape(policyName), url.QueryEscape(bindPoint), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) GetAllCRVServerICAPolicyBinding() ([]models.CRVServerICAPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, crVServerICAPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CRVServerICAPolicyBinding `json:"crvserver_icapolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CRService) GetCRVServerICAPolicyBinding(name string) ([]models.CRVServerICAPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", crVServerICAPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CRVServerICAPolicyBinding `json:"crvserver_icapolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CRService) CountCRVServerICAPolicyBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", crVServerICAPolicyBindingURL, url.QueryEscape(name))
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
		} `json:"crvserver_icapolicy_binding"`
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

// crvserver_lbvserver_binding
func (s *CRService) AddCRVServerLBVServerBinding(binding models.CRVServerLBVServerBinding) error {
	payload := map[string]any{"crvserver_lbvserver_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, crVServerLBVServerBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) DeleteCRVServerLBVServerBinding(name string, lbvserver string) error {
	reqURL := fmt.Sprintf("%s/%s?args=lbvserver:%s", crVServerLBVServerBindingURL, url.QueryEscape(name), url.QueryEscape(lbvserver))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) GetAllCRVServerLBVServerBinding() ([]models.CRVServerLBVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, crVServerLBVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CRVServerLBVServerBinding `json:"crvserver_lbvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CRService) GetCRVServerLBVServerBinding(name string) ([]models.CRVServerLBVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", crVServerLBVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CRVServerLBVServerBinding `json:"crvserver_lbvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CRService) CountCRVServerLBVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", crVServerLBVServerBindingURL, url.QueryEscape(name))
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
		} `json:"crvserver_lbvserver_binding"`
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

// crvserver_policymap_binding
func (s *CRService) AddCRVServerPolicyMapBinding(binding models.CRVServerPolicyMapBinding) error {
	payload := map[string]any{"crvserver_policymap_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, crVServerPolicyMapBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) DeleteCRVServerPolicyMapBinding(name string, policyName string, bindPoint string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s,bindpoint:%s,priority:%d", crVServerPolicyMapBindingURL, url.QueryEscape(name), url.QueryEscape(policyName), url.QueryEscape(bindPoint), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) GetAllCRVServerPolicyMapBinding() ([]models.CRVServerPolicyMapBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, crVServerPolicyMapBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CRVServerPolicyMapBinding `json:"crvserver_policymap_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CRService) GetCRVServerPolicyMapBinding(name string) ([]models.CRVServerPolicyMapBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", crVServerPolicyMapBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CRVServerPolicyMapBinding `json:"crvserver_policymap_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CRService) CountCRVServerPolicyMapBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", crVServerPolicyMapBindingURL, url.QueryEscape(name))
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
		} `json:"crvserver_policymap_binding"`
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

// crvserver_responderpolicy_binding
func (s *CRService) AddCRVServerResponderPolicyBinding(binding models.CRVServerResponderPolicyBinding) error {
	payload := map[string]any{"crvserver_responderpolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, crVServerResponderPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) DeleteCRVServerResponderPolicyBinding(name string, policyName string, bindPoint string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s,bindpoint:%s,priority:%d", crVServerResponderPolicyBindingURL, url.QueryEscape(name), url.QueryEscape(policyName), url.QueryEscape(bindPoint), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) GetAllCRVServerResponderPolicyBinding() ([]models.CRVServerResponderPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, crVServerResponderPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CRVServerResponderPolicyBinding `json:"crvserver_responderpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CRService) GetCRVServerResponderPolicyBinding(name string) ([]models.CRVServerResponderPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", crVServerResponderPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CRVServerResponderPolicyBinding `json:"crvserver_responderpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CRService) CountCRVServerResponderPolicyBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", crVServerResponderPolicyBindingURL, url.QueryEscape(name))
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
		} `json:"crvserver_responderpolicy_binding"`
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

// crvserver_rewritepolicy_binding
func (s *CRService) AddCRVServerRewritePolicyBinding(binding models.CRVServerRewritePolicyBinding) error {
	payload := map[string]any{"crvserver_rewritepolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, crVServerRewritePolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) DeleteCRVServerRewritePolicyBinding(name string, policyName string, bindPoint string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s,bindpoint:%s,priority:%d", crVServerRewritePolicyBindingURL, url.QueryEscape(name), url.QueryEscape(policyName), url.QueryEscape(bindPoint), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) GetAllCRVServerRewritePolicyBinding() ([]models.CRVServerRewritePolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, crVServerRewritePolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CRVServerRewritePolicyBinding `json:"crvserver_rewritepolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CRService) GetCRVServerRewritePolicyBinding(name string) ([]models.CRVServerRewritePolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", crVServerRewritePolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CRVServerRewritePolicyBinding `json:"crvserver_rewritepolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CRService) CountCRVServerRewritePolicyBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", crVServerRewritePolicyBindingURL, url.QueryEscape(name))
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
		} `json:"crvserver_rewritepolicy_binding"`
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

// crvserver_spilloverpolicy_binding
func (s *CRService) AddCRVServerSpilloverPolicyBinding(binding models.CRVServerSpilloverPolicyBinding) error {
	payload := map[string]any{"crvserver_spilloverpolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, crVServerSpilloverPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) DeleteCRVServerSpilloverPolicyBinding(name string, policyName string, bindPoint string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s,bindpoint:%s,priority:%d", crVServerSpilloverPolicyBindingURL, url.QueryEscape(name), url.QueryEscape(policyName), url.QueryEscape(bindPoint), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CRService) GetAllCRVServerSpilloverPolicyBinding() ([]models.CRVServerSpilloverPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, crVServerSpilloverPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CRVServerSpilloverPolicyBinding `json:"crvserver_spilloverpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CRService) GetCRVServerSpilloverPolicyBinding(name string) ([]models.CRVServerSpilloverPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", crVServerSpilloverPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CRVServerSpilloverPolicyBinding `json:"crvserver_spilloverpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CRService) CountCRVServerSpilloverPolicyBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", crVServerSpilloverPolicyBindingURL, url.QueryEscape(name))
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
		} `json:"crvserver_spilloverpolicy_binding"`
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
