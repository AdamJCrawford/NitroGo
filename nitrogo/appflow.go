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
	appFlowActionURL                          = "/nitro/v1/config/appflowaction"
	appFlowActionAnalyticsProfileBindingURL   = "/nitro/v1/config/appflowaction_analyticsprofile_binding"
	appFlowActionBindingURL                   = "/nitro/v1/config/appflowaction_binding"
	appFlowCollectorURL                       = "/nitro/v1/config/appflowcollector"
	appFlowGlobalAppFlowPolicyBindingURL      = "/nitro/v1/config/appflowglobal_appflowpolicy_binding"
	appFlowGlobalBindingURL                   = "/nitro/v1/config/appflowglobal_binding"
	appFlowParamURL                           = "/nitro/v1/config/appflowparam"
	appFlowPolicyURL                          = "/nitro/v1/config/appflowpolicy"
	appFlowPolicyAppFlowGlobalBindingURL      = "/nitro/v1/config/appflowpolicy_appflowglobal_binding"
	appFlowPolicyAppFlowPolicyLabelBindingURL = "/nitro/v1/config/appflowpolicy_appflowpolicylabel_binding"
	appFlowPolicyBindingURL                   = "/nitro/v1/config/appflowpolicy_binding"
	appFlowPolicyCSVServerBindingURL          = "/nitro/v1/config/appflowpolicy_csvserver_binding"
	appFlowPolicyLBVServerBindingURL          = "/nitro/v1/config/appflowpolicy_lbvserver_binding"
	appFlowPolicyVPNVServerBindingURL         = "/nitro/v1/config/appflowpolicy_vpnvserver_binding"
	appFlowPolicyLabelURL                     = "/nitro/v1/config/appflowpolicylabel"
	appFlowPolicyLabelAppFlowPolicyBindingURL = "/nitro/v1/config/appflowpolicylabel_appflowpolicy_binding"
	appFlowPolicyLabelBindingURL              = "/nitro/v1/config/appflowpolicylabel_binding"
)

// AppFlow configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appflow/appflow
type AppFlowService struct {
	client *Client
}

// appflowaction
// Configuration for AppFlow action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appflow/appflowaction

func (s *AppFlowService) AddAppFlowAction(action models.AppFlowAction) error {
	payload := map[string]any{
		"appflowaction": action,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, appFlowActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AppFlowService) DeleteAppFlowAction(name string) error {
	reqURL := fmt.Sprintf("%s/%s", appFlowActionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AppFlowService) UpdateAppFlowAction(action models.AppFlowAction) error {
	payload := map[string]any{
		"appflowaction": action,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, appFlowActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AppFlowService) UnsetAppFlowAction(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{
		"appflowaction": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, appFlowActionURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AppFlowService) RenameAppFlowAction(name, newname string) error {
	payload := map[string]any{
		"appflowaction": map[string]string{
			"name":    name,
			"newname": newname,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, appFlowActionURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AppFlowService) GetAllAppFlowAction() ([]models.AppFlowAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFlowActionURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		AppFlowActions []models.AppFlowAction `json:"appflowaction"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.AppFlowActions, nil
}

func (s *AppFlowService) GetAppFlowAction(name string) ([]models.AppFlowAction, error) {
	reqURL := fmt.Sprintf("%s/%s", appFlowActionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		AppFlowActions []models.AppFlowAction `json:"appflowaction"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.AppFlowActions, nil
}

func (s *AppFlowService) CountAppFlowAction() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFlowActionURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		AppFlowActions []struct {
			Count float64 `json:"__count"`
		} `json:"appflowaction"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.AppFlowActions) > 0 {
		return result.AppFlowActions[0].Count, nil
	}
	return 0, nil
}

// appflowaction_analyticsprofile_binding
// Binding object showing the analyticsprofile that can be bound to appflowaction.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appflow/appflowaction_analyticsprofile_binding

func (s *AppFlowService) AddAppFlowActionAnalyticsProfileBinding(binding models.AppFlowActionAnalyticsProfileBinding) error {
	payload := map[string]any{
		"appflowaction_analyticsprofile_binding": binding,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, appFlowActionAnalyticsProfileBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AppFlowService) DeleteAppFlowActionAnalyticsProfileBinding(name string) error {
	reqURL := fmt.Sprintf("%s/%s", appFlowActionAnalyticsProfileBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AppFlowService) GetAllAppFlowActionAnalyticsProfileBinding() ([]models.AppFlowActionAnalyticsProfileBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFlowActionAnalyticsProfileBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AppFlowActionAnalyticsProfileBinding `json:"appflowaction_analyticsprofile_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AppFlowService) GetAppFlowActionAnalyticsProfileBinding(name string) ([]models.AppFlowActionAnalyticsProfileBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFlowActionAnalyticsProfileBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AppFlowActionAnalyticsProfileBinding `json:"appflowaction_analyticsprofile_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AppFlowService) CountAppFlowActionAnalyticsProfileBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", appFlowActionAnalyticsProfileBindingURL, url.QueryEscape(name))
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
		} `json:"appflowaction_analyticsprofile_binding"`
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

// appflowaction_binding
// Binding object which returns the resources bound to appflowaction.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appflow/appflowaction_binding

func (s *AppFlowService) GetAllAppFlowActionBinding() ([]models.AppFlowActionBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFlowActionBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AppFlowActionBinding `json:"appflowaction_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AppFlowService) GetAppFlowActionBinding(name string) ([]models.AppFlowActionBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFlowActionBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AppFlowActionBinding `json:"appflowaction_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// appflowcollector
// Configuration for AppFlow collector resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appflow/appflowcollector

func (s *AppFlowService) AddAppFlowCollector(collector models.AppFlowCollector) error {
	payload := map[string]any{
		"appflowcollector": collector,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, appFlowCollectorURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AppFlowService) UpdateAppFlowCollector(collector models.AppFlowCollector) error {
	payload := map[string]any{
		"appflowcollector": collector,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, appFlowCollectorURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AppFlowService) UnsetAppFlowCollector(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{
		"appflowcollector": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, appFlowCollectorURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AppFlowService) DeleteAppFlowCollector(name string) error {
	reqURL := fmt.Sprintf("%s/%s", appFlowCollectorURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AppFlowService) RenameAppFlowCollector(name, newname string) error {
	payload := map[string]any{
		"appflowcollector": map[string]string{
			"name":    name,
			"newname": newname,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, appFlowCollectorURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AppFlowService) GetAllAppFlowCollector() ([]models.AppFlowCollector, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFlowCollectorURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		AppFlowCollectors []models.AppFlowCollector `json:"appflowcollector"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.AppFlowCollectors, nil
}

func (s *AppFlowService) GetAppFlowCollector(name string) ([]models.AppFlowCollector, error) {
	reqURL := fmt.Sprintf("%s/%s", appFlowCollectorURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		AppFlowCollectors []models.AppFlowCollector `json:"appflowcollector"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.AppFlowCollectors, nil
}

func (s *AppFlowService) CountAppFlowCollector() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFlowCollectorURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		AppFlowCollectors []struct {
			Count float64 `json:"__count"`
		} `json:"appflowcollector"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.AppFlowCollectors) > 0 {
		return result.AppFlowCollectors[0].Count, nil
	}
	return 0, nil
}

// appflowglobal_appflowpolicy_binding
// Binding object showing the appflowpolicy that can be bound to appflowglobal.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appflow/appflowglobal_appflowpolicy_binding

func (s *AppFlowService) AddAppFlowGlobalAppFlowPolicyBinding(binding models.AppFlowGlobalAppFlowPolicyBinding) error {
	payload := map[string]any{
		"appflowglobal_appflowpolicy_binding": binding,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, appFlowGlobalAppFlowPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AppFlowService) DeleteAppFlowGlobalAppFlowPolicyBinding(policyname string) error {
	reqURL := fmt.Sprintf("%s?args=policyname:%s", appFlowGlobalAppFlowPolicyBindingURL, url.QueryEscape(policyname))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AppFlowService) GetAppFlowGlobalAppFlowPolicyBinding() ([]models.AppFlowGlobalAppFlowPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFlowGlobalAppFlowPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AppFlowGlobalAppFlowPolicyBinding `json:"appflowglobal_appflowpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AppFlowService) CountAppFlowGlobalAppFlowPolicyBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFlowGlobalAppFlowPolicyBindingURL+"?count=yes", nil)
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
		} `json:"appflowglobal_appflowpolicy_binding"`
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

// appflowglobal_binding
// Binding object which returns the resources bound to appflowglobal.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appflow/appflowglobal_binding

func (s *AppFlowService) GetAppFlowGlobalBinding() ([]models.AppFlowGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFlowGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AppFlowGlobalBinding `json:"appflowglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// appflowparam
// Configuration for AppFlow parameter resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appflow/appflowparam

func (s *AppFlowService) UpdateAppFlowParam(param models.AppFlowParam) error {
	payload := map[string]any{
		"appflowparam": param,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, appFlowParamURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AppFlowService) UnsetAppFlowParam(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{
		"appflowparam": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, appFlowParamURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AppFlowService) GetAllAppFlowParam() (models.AppFlowParam, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFlowParamURL, nil)
	if err != nil {
		return models.AppFlowParam{}, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return models.AppFlowParam{}, err
	}
	var result struct {
		AppFlowParams []models.AppFlowParam `json:"appflowparam"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.AppFlowParam{}, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.AppFlowParams) > 0 {
		return result.AppFlowParams[0], nil
	}
	return models.AppFlowParam{}, nil
}

// appflowpolicy
// Configuration for AppFlow policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appflow/appflowpolicy

func (s *AppFlowService) AddAppFlowPolicy(policy models.AppFlowPolicy) error {
	payload := map[string]any{
		"appflowpolicy": policy,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, appFlowPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AppFlowService) DeleteAppFlowPolicy(name string) error {
	reqURL := fmt.Sprintf("%s/%s", appFlowPolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AppFlowService) UpdateAppFlowPolicy(policy models.AppFlowPolicy) error {
	payload := map[string]any{
		"appflowpolicy": policy,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, appFlowPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AppFlowService) UnsetAppFlowPolicy(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{
		"appflowpolicy": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, appFlowPolicyURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AppFlowService) RenameAppFlowPolicy(name, newname string) error {
	payload := map[string]any{
		"appflowpolicy": map[string]string{
			"name":    name,
			"newname": newname,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, appFlowPolicyURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AppFlowService) GetAllAppFlowPolicy() ([]models.AppFlowPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFlowPolicyURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		AppFlowPolicies []models.AppFlowPolicy `json:"appflowpolicy"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.AppFlowPolicies, nil
}

func (s *AppFlowService) GetAppFlowPolicy(name string) ([]models.AppFlowPolicy, error) {
	reqURL := fmt.Sprintf("%s/%s", appFlowPolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		AppFlowPolicies []models.AppFlowPolicy `json:"appflowpolicy"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.AppFlowPolicies, nil
}

func (s *AppFlowService) CountAppFlowPolicy() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFlowPolicyURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		AppFlowPolicies []struct {
			Count float64 `json:"__count"`
		} `json:"appflowpolicy"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.AppFlowPolicies) > 0 {
		return result.AppFlowPolicies[0].Count, nil
	}
	return 0, nil
}

// appflowpolicylabel
// Configuration for AppFlow policy label resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appflow/appflowpolicylabel

func (s *AppFlowService) AddAppFlowPolicyLabel(label models.AppFlowPolicyLabel) error {
	payload := map[string]any{
		"appflowpolicylabel": label,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, appFlowPolicyLabelURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AppFlowService) DeleteAppFlowPolicyLabel(labelname string) error {
	reqURL := fmt.Sprintf("%s/%s", appFlowPolicyLabelURL, url.QueryEscape(labelname))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AppFlowService) RenameAppFlowPolicyLabel(labelname, newname string) error {
	payload := map[string]any{
		"appflowpolicylabel": map[string]string{
			"labelname": labelname,
			"newname":   newname,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, appFlowPolicyLabelURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AppFlowService) GetAllAppFlowPolicyLabel() ([]models.AppFlowPolicyLabel, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFlowPolicyLabelURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		AppFlowPolicyLabels []models.AppFlowPolicyLabel `json:"appflowpolicylabel"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.AppFlowPolicyLabels, nil
}

func (s *AppFlowService) GetAppFlowPolicyLabel(labelname string) ([]models.AppFlowPolicyLabel, error) {
	reqURL := fmt.Sprintf("%s/%s", appFlowPolicyLabelURL, url.QueryEscape(labelname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		AppFlowPolicyLabels []models.AppFlowPolicyLabel `json:"appflowpolicylabel"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.AppFlowPolicyLabels, nil
}

func (s *AppFlowService) CountAppFlowPolicyLabel() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFlowPolicyLabelURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		AppFlowPolicyLabels []struct {
			Count float64 `json:"__count"`
		} `json:"appflowpolicylabel"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.AppFlowPolicyLabels) > 0 {
		return result.AppFlowPolicyLabels[0].Count, nil
	}
	return 0, nil
}

// appflowpolicylabel_appflowpolicy_binding
// Binding object showing the appflowpolicy that can be bound to appflowpolicylabel.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appflow/appflowpolicylabel_appflowpolicy_binding

func (s *AppFlowService) AddAppFlowPolicyLabelAppFlowPolicyBinding(binding models.AppFlowPolicyLabelAppFlowPolicyBinding) error {
	payload := map[string]any{
		"appflowpolicylabel_appflowpolicy_binding": binding,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, appFlowPolicyLabelAppFlowPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AppFlowService) DeleteAppFlowPolicyLabelAppFlowPolicyBinding(labelname string, policyname string) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s", appFlowPolicyLabelAppFlowPolicyBindingURL, url.QueryEscape(labelname), url.QueryEscape(policyname))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *AppFlowService) GetAllAppFlowPolicyLabelAppFlowPolicyBinding() ([]models.AppFlowPolicyLabelAppFlowPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFlowPolicyLabelAppFlowPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AppFlowPolicyLabelAppFlowPolicyBinding `json:"appflowpolicylabel_appflowpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AppFlowService) GetAppFlowPolicyLabelAppFlowPolicyBinding(labelname string) ([]models.AppFlowPolicyLabelAppFlowPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFlowPolicyLabelAppFlowPolicyBindingURL, url.QueryEscape(labelname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AppFlowPolicyLabelAppFlowPolicyBinding `json:"appflowpolicylabel_appflowpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AppFlowService) CountAppFlowPolicyLabelAppFlowPolicyBinding(labelname string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", appFlowPolicyLabelAppFlowPolicyBindingURL, url.QueryEscape(labelname))
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
		} `json:"appflowpolicylabel_appflowpolicy_binding"`
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

// appflowpolicylabel_binding
// Binding object which returns the resources bound to appflowpolicylabel.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appflow/appflowpolicylabel_binding

func (s *AppFlowService) GetAllAppFlowPolicyLabelBinding() ([]models.AppFlowPolicyLabelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFlowPolicyLabelBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AppFlowPolicyLabelBinding `json:"appflowpolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AppFlowService) GetAppFlowPolicyLabelBinding(labelname string) ([]models.AppFlowPolicyLabelBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFlowPolicyLabelBindingURL, url.QueryEscape(labelname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AppFlowPolicyLabelBinding `json:"appflowpolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// appflowpolicy_appflowglobal_binding
// Binding object showing the appflowglobal that can be bound to appflowpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appflow/appflowpolicy_appflowglobal_binding

func (s *AppFlowService) GetAllAppFlowPolicyAppFlowGlobalBinding() ([]models.AppFlowPolicyAppFlowGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFlowPolicyAppFlowGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AppFlowPolicyAppFlowGlobalBinding `json:"appflowpolicy_appflowglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AppFlowService) GetAppFlowPolicyAppFlowGlobalBinding(name string) ([]models.AppFlowPolicyAppFlowGlobalBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFlowPolicyAppFlowGlobalBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AppFlowPolicyAppFlowGlobalBinding `json:"appflowpolicy_appflowglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AppFlowService) CountAppFlowPolicyAppFlowGlobalBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", appFlowPolicyAppFlowGlobalBindingURL, url.QueryEscape(name))
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
		} `json:"appflowpolicy_appflowglobal_binding"`
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

// appflowpolicy_appflowpolicylabel_binding
// Binding object showing the appflowpolicylabel that can be bound to appflowpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appflow/appflowpolicy_appflowpolicylabel_binding

func (s *AppFlowService) GetAllAppFlowPolicyAppFlowPolicyLabelBinding() ([]models.AppFlowPolicyAppFlowPolicyLabelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFlowPolicyAppFlowPolicyLabelBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AppFlowPolicyAppFlowPolicyLabelBinding `json:"appflowpolicy_appflowpolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AppFlowService) GetAppFlowPolicyAppFlowPolicyLabelBinding(name string) ([]models.AppFlowPolicyAppFlowPolicyLabelBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFlowPolicyAppFlowPolicyLabelBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AppFlowPolicyAppFlowPolicyLabelBinding `json:"appflowpolicy_appflowpolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AppFlowService) CountAppFlowPolicyAppFlowPolicyLabelBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", appFlowPolicyAppFlowPolicyLabelBindingURL, url.QueryEscape(name))
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
		} `json:"appflowpolicy_appflowpolicylabel_binding"`
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

// appflowpolicy_binding
// Binding object which returns the resources bound to appflowpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appflow/appflowpolicy_binding

func (s *AppFlowService) GetAllAppFlowPolicyBinding() ([]models.AppFlowPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFlowPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AppFlowPolicyBinding `json:"appflowpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AppFlowService) GetAppFlowPolicyBinding(name string) ([]models.AppFlowPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFlowPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AppFlowPolicyBinding `json:"appflowpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// appflowpolicy_csvserver_binding
// Binding object showing the csvserver that can be bound to appflowpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appflow/appflowpolicy_csvserver_binding

func (s *AppFlowService) GetAllAppFlowPolicyCSVServerBinding() ([]models.AppFlowPolicyCSVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFlowPolicyCSVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AppFlowPolicyCSVServerBinding `json:"appflowpolicy_csvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AppFlowService) GetAppFlowPolicyCSVServerBinding(name string) ([]models.AppFlowPolicyCSVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFlowPolicyCSVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AppFlowPolicyCSVServerBinding `json:"appflowpolicy_csvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AppFlowService) CountAppFlowPolicyCSVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", appFlowPolicyCSVServerBindingURL, url.QueryEscape(name))
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
		} `json:"appflowpolicy_csvserver_binding"`
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

// appflowpolicy_lbvserver_binding
// Binding object showing the lbvserver that can be bound to appflowpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appflow/appflowpolicy_lbvserver_binding

func (s *AppFlowService) GetAllAppFlowPolicy_LBVServerBinding() ([]models.AppFlowPolicyLBVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFlowPolicyLBVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AppFlowPolicyLBVServerBinding `json:"appflowpolicy_lbvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AppFlowService) GetAppFlowPolicy_LBVServerBinding(name string) ([]models.AppFlowPolicyLBVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFlowPolicyLBVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AppFlowPolicyLBVServerBinding `json:"appflowpolicy_lbvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AppFlowService) CountAppFlowPolicy_LBVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", appFlowPolicyLBVServerBindingURL, url.QueryEscape(name))
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
		} `json:"appflowpolicy_lbvserver_binding"`
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

// appflowpolicy_vpnvserver_binding
// Binding object showing the vpnvserver that can be bound to appflowpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appflow/appflowpolicy_vpnvserver_binding

func (s *AppFlowService) GetAllAppFlowPolicyVPNVServerBinding() ([]models.AppFlowPolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, appFlowPolicyVPNVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AppFlowPolicyVPNVServerBinding `json:"appflowpolicy_vpnvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AppFlowService) GetAppFlowPolicyVPNVServerBinding(name string) ([]models.AppFlowPolicyVPNVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", appFlowPolicyVPNVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.AppFlowPolicyVPNVServerBinding `json:"appflowpolicy_vpnvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *AppFlowService) CountAppFlowPolicyVPNVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", appFlowPolicyVPNVServerBindingURL, url.QueryEscape(name))
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
		} `json:"appflowpolicy_vpnvserver_binding"`
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
