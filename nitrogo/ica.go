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
	icaAccessProfileURL           = "/nitro/v1/config/icaaccessprofile"
	icaActionURL                  = "/nitro/v1/config/icaaction"
	icaGlobalBindingURL           = "/nitro/v1/config/icaglobal_binding"
	icaGlobalICAPolicyBindingURL  = "/nitro/v1/config/icaglobal_icapolicy_binding"
	icaLatencyProfileURL          = "/nitro/v1/config/icalatencyprofile"
	icaParameterURL               = "/nitro/v1/config/icaparameter"
	icaPolicyURL                  = "/nitro/v1/config/icapolicy"
	icaPolicyBindingURL           = "/nitro/v1/config/icapolicy_binding"
	icaPolicyCRVServerBindingURL  = "/nitro/v1/config/icapolicy_crvserver_binding"
	icaPolicyICAGlobalBindingURL  = "/nitro/v1/config/icapolicy_icaglobal_binding"
	icaPolicyVPNVServerBindingURL = "/nitro/v1/config/icapolicy_vpnvserver_binding"
)

// ICA configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/ica/ica
type ICAService struct {
	client *Client
}

// icaaccessprofile

func (s *ICAService) AddICAAccessProfile(profile models.ICAAccessProfile) error {
	payload := map[string]any{
		"icaaccessprofile": profile,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, icaAccessProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ICAService) DeleteICAAccessProfile(name string) error {
	urlReq := fmt.Sprintf("%s/%s", icaAccessProfileURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, urlReq, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ICAService) UpdateICAAccessProfile(profile models.ICAAccessProfile) error {
	payload := map[string]any{
		"icaaccessprofile": profile,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, icaAccessProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ICAService) UnsetICAAccessProfile(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"icaaccessprofile": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, icaAccessProfileURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ICAService) GetAllICAAccessProfile() ([]models.ICAAccessProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, icaAccessProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Profiles []models.ICAAccessProfile `json:"icaaccessprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Profiles, nil
}

func (s *ICAService) GetICAAccessProfile(name string) ([]models.ICAAccessProfile, error) {
	urlReq := fmt.Sprintf("%s/%s", icaAccessProfileURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Profiles []models.ICAAccessProfile `json:"icaaccessprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Profiles, nil
}

func (s *ICAService) CountICAAccessProfile() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, icaAccessProfileURL+"?count=yes", nil)
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
		} `json:"icaaccessprofile"`
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

// icaaction

func (s *ICAService) AddICAAction(action models.ICAAction) error {
	payload := map[string]any{
		"icaaction": action,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, icaActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ICAService) DeleteICAAction(name string) error {
	urlReq := fmt.Sprintf("%s/%s", icaActionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, urlReq, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ICAService) UpdateICAAction(action models.ICAAction) error {
	payload := map[string]any{
		"icaaction": action,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, icaActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ICAService) UnsetICAAction(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"icaaction": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, icaActionURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ICAService) GetAllICAAction() ([]models.ICAAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, icaActionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Actions []models.ICAAction `json:"icaaction"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Actions, nil
}

func (s *ICAService) GetICAAction(name string) ([]models.ICAAction, error) {
	urlReq := fmt.Sprintf("%s/%s", icaActionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Actions []models.ICAAction `json:"icaaction"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Actions, nil
}

func (s *ICAService) CountICAAction() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, icaActionURL+"?count=yes", nil)
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
		} `json:"icaaction"`
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

func (s *ICAService) RenameICAAction(name string, newname string) error {
	payload := map[string]any{
		"icaaction": map[string]string{
			"name":    name,
			"newname": newname,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, icaActionURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// icaglobal_binding

func (s *ICAService) GetICAGlobalBinding() ([]models.ICAGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, icaGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.ICAGlobalBinding `json:"icaglobal_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

// icaglobal_icapolicy_binding

func (s *ICAService) AddICAGlobalICAPolicyBinding(binding models.ICAGlobalICAPolicyBinding) error {
	payload := map[string]any{
		"icaglobal_icapolicy_binding": binding,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, icaGlobalICAPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ICAService) DeleteICAGlobalICAPolicyBinding(policyname string) error {
	urlReq := fmt.Sprintf("%s?args=policyname:%s", icaGlobalICAPolicyBindingURL, url.QueryEscape(policyname))
	req, err := s.client.NewRequest(http.MethodDelete, urlReq, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ICAService) GetICAGlobalICAPolicyBinding() ([]models.ICAGlobalICAPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, icaGlobalICAPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.ICAGlobalICAPolicyBinding `json:"icaglobal_icapolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *ICAService) CountICAGlobalICAPolicyBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, icaGlobalICAPolicyBindingURL+"?count=yes", nil)
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
		} `json:"icaglobal_icapolicy_binding"`
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

// icalatencyprofile

func (s *ICAService) AddICALatencyProfile(profile models.ICALatencyProfile) error {
	payload := map[string]any{
		"icalatencyprofile": profile,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, icaLatencyProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ICAService) DeleteICALatencyProfile(name string) error {
	urlReq := fmt.Sprintf("%s/%s", icaLatencyProfileURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, urlReq, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ICAService) UpdateICALatencyProfile(profile models.ICALatencyProfile) error {
	payload := map[string]any{
		"icalatencyprofile": profile,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, icaLatencyProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ICAService) UnsetICALatencyProfile(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"icalatencyprofile": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, icaLatencyProfileURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ICAService) GetAllICALatencyProfile() ([]models.ICALatencyProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, icaLatencyProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Profiles []models.ICALatencyProfile `json:"icalatencyprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Profiles, nil
}

func (s *ICAService) GetICALatencyProfile(name string) ([]models.ICALatencyProfile, error) {
	urlReq := fmt.Sprintf("%s/%s", icaLatencyProfileURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Profiles []models.ICALatencyProfile `json:"icalatencyprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Profiles, nil
}

func (s *ICAService) CountICALatencyProfile() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, icaLatencyProfileURL+"?count=yes", nil)
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
		} `json:"icalatencyprofile"`
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

// icaparameter

func (s *ICAService) UpdateICAParameter(param models.ICAParameter) error {
	payload := map[string]any{
		"icaparameter": param,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, icaParameterURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ICAService) UnsetICAParameter(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"icaparameter": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, icaParameterURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ICAService) GetAllICAParameter() (models.ICAParameter, error) {
	req, err := s.client.NewRequest(http.MethodGet, icaParameterURL, nil)
	if err != nil {
		return models.ICAParameter{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.ICAParameter{}, err
	}

	var result struct {
		Parameters []models.ICAParameter `json:"icaparameter"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.ICAParameter{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Parameters) > 0 {
		return result.Parameters[0], nil
	}

	return models.ICAParameter{}, nil
}

// icapolicy

func (s *ICAService) AddICAPolicy(policy models.ICAPolicy) error {
	payload := map[string]any{
		"icapolicy": policy,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, icaPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ICAService) DeleteICAPolicy(name string) error {
	urlReq := fmt.Sprintf("%s/%s", icaPolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, urlReq, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ICAService) UpdateICAPolicy(policy models.ICAPolicy) error {
	payload := map[string]any{
		"icapolicy": policy,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, icaPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ICAService) UnsetICAPolicy(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"icapolicy": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, icaPolicyURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ICAService) GetAllICAPolicy() ([]models.ICAPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, icaPolicyURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Policies []models.ICAPolicy `json:"icapolicy"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Policies, nil
}

func (s *ICAService) GetICAPolicy(name string) ([]models.ICAPolicy, error) {
	urlReq := fmt.Sprintf("%s/%s", icaPolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Policies []models.ICAPolicy `json:"icapolicy"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Policies, nil
}

func (s *ICAService) CountICAPolicy() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, icaPolicyURL+"?count=yes", nil)
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
		} `json:"icapolicy"`
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

func (s *ICAService) RenameICAPolicy(name string, newname string) error {
	payload := map[string]any{
		"icapolicy": map[string]string{
			"name":    name,
			"newname": newname,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, icaPolicyURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// icapolicy_binding

func (s *ICAService) GetAllICAPolicyBinding() ([]models.ICAPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, icaPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.ICAPolicyBinding `json:"icapolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *ICAService) GetICAPolicyBinding(name string) ([]models.ICAPolicyBinding, error) {
	urlReq := fmt.Sprintf("%s/%s", icaPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.ICAPolicyBinding `json:"icapolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

// icapolicy_crvserver_binding

func (s *ICAService) GetAllICAPolicyCRVServerBinding() ([]models.ICAPolicyCRVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, icaPolicyCRVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.ICAPolicyCRVServerBinding `json:"icapolicy_crvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *ICAService) GetICAPolicyCRVServerBinding(name string) ([]models.ICAPolicyCRVServerBinding, error) {
	urlReq := fmt.Sprintf("%s/%s", icaPolicyCRVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.ICAPolicyCRVServerBinding `json:"icapolicy_crvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *ICAService) CountICAPolicyCRVServerBinding(name string) (float64, error) {
	urlReq := fmt.Sprintf("%s/%s?count=yes", icaPolicyCRVServerBindingURL, url.QueryEscape(name))
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
		} `json:"icapolicy_crvserver_binding"`
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

// icapolicy_icaglobal_binding

func (s *ICAService) GetAllICAPolicyICAGlobalBinding() ([]models.ICAPolicyICAGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, icaPolicyICAGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.ICAPolicyICAGlobalBinding `json:"icapolicy_icaglobal_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *ICAService) GetICAPolicyICAGlobalBinding(name string) ([]models.ICAPolicyICAGlobalBinding, error) {
	urlReq := fmt.Sprintf("%s/%s", icaPolicyICAGlobalBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.ICAPolicyICAGlobalBinding `json:"icapolicy_icaglobal_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *ICAService) CountICAPolicyICAGlobalBinding(name string) (float64, error) {
	urlReq := fmt.Sprintf("%s/%s?count=yes", icaPolicyICAGlobalBindingURL, url.QueryEscape(name))
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
		} `json:"icapolicy_icaglobal_binding"`
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

// icapolicy_vpnvserver_binding

func (s *ICAService) GetAllICAPolicyVPNVServerBinding() ([]models.ICAPolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, icaPolicyVPNVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.ICAPolicyVPNVServerBinding `json:"icapolicy_vpnvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *ICAService) GetICAPolicyVPNVServerBinding(name string) ([]models.ICAPolicyVPNVServerBinding, error) {
	urlReq := fmt.Sprintf("%s/%s", icaPolicyVPNVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.ICAPolicyVPNVServerBinding `json:"icapolicy_vpnvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *ICAService) CountICAPolicyVPNVServerBinding(name string) (float64, error) {
	urlReq := fmt.Sprintf("%s/%s?count=yes", icaPolicyVPNVServerBindingURL, url.QueryEscape(name))
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
		} `json:"icapolicy_vpnvserver_binding"`
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
