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
	responderActionURL                            = "/nitro/v1/config/responderaction"
	responderGlobalBindingURL                     = "/nitro/v1/config/responderglobal_binding"
	responderGlobalResponderPolicyBindingURL      = "/nitro/v1/config/responderglobal_responderpolicy_binding"
	responderHTMLPageURL                          = "/nitro/v1/config/responderhtmlpage"
	responderParamURL                             = "/nitro/v1/config/responderparam"
	responderPolicyURL                            = "/nitro/v1/config/responderpolicy"
	responderPolicyLabelURL                       = "/nitro/v1/config/responderpolicylabel"
	responderPolicyLabelBindingURL                = "/nitro/v1/config/responderpolicylabel_binding"
	responderPolicyLabelPolicyBindingBindingURL   = "/nitro/v1/config/responderpolicylabel_policybinding_binding"
	responderPolicyLabelResponderPolicyBindingURL = "/nitro/v1/config/responderpolicylabel_responderpolicy_binding"
	responderPolicyBindingURL                     = "/nitro/v1/config/responderpolicy_binding"
	responderPolicyCRVServerBindingURL            = "/nitro/v1/config/responderpolicy_crvserver_binding"
	responderPolicyCSVServerBindingURL            = "/nitro/v1/config/responderpolicy_csvserver_binding"
	responderPolicyLBVServerBindingURL            = "/nitro/v1/config/responderpolicy_lbvserver_binding"
	responderPolicyResponderGlobalBindingURL      = "/nitro/v1/config/responderpolicy_responderglobal_binding"
	responderPolicyResponderPolicyLabelBindingURL = "/nitro/v1/config/responderpolicy_responderpolicylabel_binding"
	responderPolicyVPNVServerBindingURL           = "/nitro/v1/config/responderpolicy_vpnvserver_binding"
)

// Responder configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/responder/responder
type ResponderService struct {
	client *Client
}

// responderaction

func (s *ResponderService) AddResponderAction(action models.ResponderAction) error {
	payload := map[string]any{
		"responderaction": action,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, responderActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ResponderService) DeleteResponderAction(name string) error {
	reqURL := fmt.Sprintf("%s/%s", responderActionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ResponderService) UpdateResponderAction(action models.ResponderAction) error {
	payload := map[string]any{
		"responderaction": action,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, responderActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ResponderService) UnsetResponderAction(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{
		"responderaction": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, responderActionURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ResponderService) GetAllResponderAction() ([]models.ResponderAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, responderActionURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Actions []models.ResponderAction `json:"responderaction"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Actions, nil
}

func (s *ResponderService) GetResponderAction(name string) ([]models.ResponderAction, error) {
	reqURL := fmt.Sprintf("%s/%s", responderActionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Actions []models.ResponderAction `json:"responderaction"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Actions, nil
}

func (s *ResponderService) CountResponderAction() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, responderActionURL+"?count=yes", nil)
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
		} `json:"responderaction"`
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

func (s *ResponderService) RenameResponderAction(name, newname string) error {
	payload := map[string]any{
		"responderaction": map[string]string{
			"name":    name,
			"newname": newname,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, responderActionURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

// responderglobal_binding

func (s *ResponderService) GetResponderGlobalBinding() ([]models.ResponderGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, responderGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ResponderGlobalBinding `json:"responderglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// responderglobal_responderpolicy_binding

func (s *ResponderService) AddResponderGlobalResponderPolicyBinding(binding models.ResponderGlobalResponderPolicyBinding) error {
	payload := map[string]any{
		"responderglobal_responderpolicy_binding": binding,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, responderGlobalResponderPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ResponderService) DeleteResponderGlobalResponderPolicyBinding(policyname, typeField string, priority int) error {
	reqURL := fmt.Sprintf("%s?args=policyname:%s,type:%s,priority:%d", responderGlobalResponderPolicyBindingURL, url.QueryEscape(policyname), url.QueryEscape(typeField), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ResponderService) GetResponderGlobalResponderPolicyBinding() ([]models.ResponderGlobalResponderPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, responderGlobalResponderPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ResponderGlobalResponderPolicyBinding `json:"responderglobal_responderpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ResponderService) CountResponderGlobalResponderPolicyBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, responderGlobalResponderPolicyBindingURL+"?count=yes", nil)
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
		} `json:"responderglobal_responderpolicy_binding"`
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

// responderhtmlpage

func (s *ResponderService) ImportResponderHTMLPage(page models.ResponderHTMLPage) error {
	payload := map[string]any{
		"responderhtmlpage": page,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, responderHTMLPageURL+"?action=Import", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ResponderService) DeleteResponderHTMLPage(name string) error {
	reqURL := fmt.Sprintf("%s/%s", responderHTMLPageURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ResponderService) GetAllResponderHTMLPage() ([]models.ResponderHTMLPage, error) {
	req, err := s.client.NewRequest(http.MethodGet, responderHTMLPageURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Pages []models.ResponderHTMLPage `json:"responderhtmlpage"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Pages, nil
}

func (s *ResponderService) GetResponderHTMLPage(name string) ([]models.ResponderHTMLPage, error) {
	reqURL := fmt.Sprintf("%s/%s", responderHTMLPageURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Pages []models.ResponderHTMLPage `json:"responderhtmlpage"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Pages, nil
}

func (s *ResponderService) ChangeResponderHTMLPage(page models.ResponderHTMLPage) error {
	payload := map[string]any{
		"responderhtmlpage": page,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, responderHTMLPageURL+"?action=update", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

// responderparam

func (s *ResponderService) UpdateResponderParam(param models.ResponderParam) error {
	payload := map[string]any{
		"responderparam": param,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, responderParamURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ResponderService) UnsetResponderParam(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{
		"responderparam": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, responderParamURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ResponderService) GetAllResponderParam() (models.ResponderParam, error) {
	req, err := s.client.NewRequest(http.MethodGet, responderParamURL, nil)
	if err != nil {
		return models.ResponderParam{}, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return models.ResponderParam{}, err
	}
	var result struct {
		Params []models.ResponderParam `json:"responderparam"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.ResponderParam{}, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Params) > 0 {
		return result.Params[0], nil
	}
	return models.ResponderParam{}, nil
}

// responderpolicy

func (s *ResponderService) AddResponderPolicy(policy models.ResponderPolicy) error {
	payload := map[string]any{
		"responderpolicy": policy,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, responderPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ResponderService) DeleteResponderPolicy(name string) error {
	reqURL := fmt.Sprintf("%s/%s", responderPolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ResponderService) UpdateResponderPolicy(policy models.ResponderPolicy) error {
	payload := map[string]any{
		"responderpolicy": policy,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, responderPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ResponderService) UnsetResponderPolicy(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{
		"responderpolicy": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, responderPolicyURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ResponderService) GetAllResponderPolicy() ([]models.ResponderPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, responderPolicyURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Policies []models.ResponderPolicy `json:"responderpolicy"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Policies, nil
}

func (s *ResponderService) GetResponderPolicy(name string) ([]models.ResponderPolicy, error) {
	reqURL := fmt.Sprintf("%s/%s", responderPolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Policies []models.ResponderPolicy `json:"responderpolicy"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Policies, nil
}

func (s *ResponderService) CountResponderPolicy() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, responderPolicyURL+"?count=yes", nil)
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
		} `json:"responderpolicy"`
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

func (s *ResponderService) RenameResponderPolicy(name, newname string) error {
	payload := map[string]any{
		"responderpolicy": map[string]string{
			"name":    name,
			"newname": newname,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, responderPolicyURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

// responderpolicylabel

func (s *ResponderService) AddResponderPolicyLabel(label models.ResponderPolicyLabel) error {
	payload := map[string]any{
		"responderpolicylabel": label,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, responderPolicyLabelURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ResponderService) DeleteResponderPolicyLabel(labelname string) error {
	reqURL := fmt.Sprintf("%s/%s", responderPolicyLabelURL, url.QueryEscape(labelname))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ResponderService) GetAllResponderPolicyLabel() ([]models.ResponderPolicyLabel, error) {
	req, err := s.client.NewRequest(http.MethodGet, responderPolicyLabelURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Labels []models.ResponderPolicyLabel `json:"responderpolicylabel"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Labels, nil
}

func (s *ResponderService) GetResponderPolicyLabel(labelname string) ([]models.ResponderPolicyLabel, error) {
	reqURL := fmt.Sprintf("%s/%s", responderPolicyLabelURL, url.QueryEscape(labelname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Labels []models.ResponderPolicyLabel `json:"responderpolicylabel"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Labels, nil
}

func (s *ResponderService) CountResponderPolicyLabel() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, responderPolicyLabelURL+"?count=yes", nil)
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
		} `json:"responderpolicylabel"`
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

func (s *ResponderService) RenameResponderPolicyLabel(labelname, newname string) error {
	payload := map[string]any{
		"responderpolicylabel": map[string]string{
			"labelname": labelname,
			"newname":   newname,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, responderPolicyLabelURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

// responderpolicylabel_binding

func (s *ResponderService) GetAllResponderPolicyLabelBinding() ([]models.ResponderPolicyLabelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, responderPolicyLabelBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ResponderPolicyLabelBinding `json:"responderpolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ResponderService) GetResponderPolicyLabelBinding(labelname string) ([]models.ResponderPolicyLabelBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", responderPolicyLabelBindingURL, url.QueryEscape(labelname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ResponderPolicyLabelBinding `json:"responderpolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// respondrepolicylabel_policybinding_binding

func (s *ResponderService) GetAllResponderPolicyLabelPolicyBindingBinding() ([]models.ResponderPolicyLabelPolicyBindingBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, responderPolicyLabelPolicyBindingBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ResponderPolicyLabelPolicyBindingBinding `json:"responderpolicylabel_policybinding_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ResponderService) GetResponderPolicyLabelPolicyBindingBinding(labelname string) ([]models.ResponderPolicyLabelPolicyBindingBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", responderPolicyLabelPolicyBindingBindingURL, url.QueryEscape(labelname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ResponderPolicyLabelPolicyBindingBinding `json:"responderpolicylabel_policybinding_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ResponderService) CountResponderPolicyLabelPolicyBindingBinding(labelname string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", responderPolicyLabelPolicyBindingBindingURL, url.QueryEscape(labelname))
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
		} `json:"responderpolicylabel_policybinding_binding"`
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

// responderpolicylabel_responderpolicy_binding

func (s *ResponderService) AddResponderPolicyLabelResponderPolicyBinding(binding models.ResponderPolicyLabelResponderPolicyBinding) error {
	payload := map[string]any{
		"responderpolicylabel_responderpolicy_binding": binding,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, responderPolicyLabelResponderPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ResponderService) DeleteResponderPolicyLabelResponderPolicyBinding(labelname, policyname string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s,priority:%d", responderPolicyLabelResponderPolicyBindingURL, url.QueryEscape(labelname), url.QueryEscape(policyname), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ResponderService) GetAllResponderPolicyLabelResponderPolicyBinding() ([]models.ResponderPolicyLabelResponderPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, responderPolicyLabelResponderPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ResponderPolicyLabelResponderPolicyBinding `json:"responderpolicylabel_responderpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ResponderService) GetResponderPolicyLabelResponderPolicyBinding(labelname string) ([]models.ResponderPolicyLabelResponderPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", responderPolicyLabelResponderPolicyBindingURL, url.QueryEscape(labelname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ResponderPolicyLabelResponderPolicyBinding `json:"responderpolicylabel_responderpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ResponderService) CountResponderPolicyLabelResponderPolicyBinding(labelname string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", responderPolicyLabelResponderPolicyBindingURL, url.QueryEscape(labelname))
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
		} `json:"responderpolicylabel_responderpolicy_binding"`
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

// responderpolicy_binding

func (s *ResponderService) GetAllResponderPolicyBinding() ([]models.ResponderPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, responderPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ResponderPolicyBinding `json:"responderpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ResponderService) GetResponderPolicyBinding(name string) ([]models.ResponderPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", responderPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ResponderPolicyBinding `json:"responderpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// responderpolicy_crvserver_binding

func (s *ResponderService) GetAllResponderPolicyCRVServerBinding() ([]models.ResponderPolicyCRVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, responderPolicyCRVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ResponderPolicyCRVServerBinding `json:"responderpolicy_crvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ResponderService) GetResponderPolicyCRVServerBinding(name string) ([]models.ResponderPolicyCRVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", responderPolicyCRVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ResponderPolicyCRVServerBinding `json:"responderpolicy_crvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ResponderService) CountResponderPolicyCRVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", responderPolicyCRVServerBindingURL, url.QueryEscape(name))
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
		} `json:"responderpolicy_crvserver_binding"`
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

// responderpolicy_csvserver_binding

func (s *ResponderService) GetAllResponderPolicyCSVServerBinding() ([]models.ResponderPolicyCSVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, responderPolicyCSVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ResponderPolicyCSVServerBinding `json:"responderpolicy_csvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ResponderService) GetResponderPolicyCSVServerBinding(name string) ([]models.ResponderPolicyCSVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", responderPolicyCSVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ResponderPolicyCSVServerBinding `json:"responderpolicy_csvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ResponderService) CountResponderPolicyCSVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", responderPolicyCSVServerBindingURL, url.QueryEscape(name))
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
		} `json:"responderpolicy_csvserver_binding"`
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

// responderpolicy_lbvserver_binding

func (s *ResponderService) GetAllResponderPolicyLBVServerBinding() ([]models.ResponderPolicyLBVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, responderPolicyLBVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ResponderPolicyLBVServerBinding `json:"responderpolicy_lbvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ResponderService) GetResponderPolicyLBVServerBinding(name string) ([]models.ResponderPolicyLBVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", responderPolicyLBVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ResponderPolicyLBVServerBinding `json:"responderpolicy_lbvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ResponderService) CountResponderPolicyLBVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", responderPolicyLBVServerBindingURL, url.QueryEscape(name))
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
		} `json:"responderpolicy_lbvserver_binding"`
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

// responderpolicy_responderglobal_binding

func (s *ResponderService) GetAllResponderPolicyResponderGlobalBinding() ([]models.ResponderPolicyResponderGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, responderPolicyResponderGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ResponderPolicyResponderGlobalBinding `json:"responderpolicy_responderglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ResponderService) GetResponderPolicyResponderGlobalBinding(name string) ([]models.ResponderPolicyResponderGlobalBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", responderPolicyResponderGlobalBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ResponderPolicyResponderGlobalBinding `json:"responderpolicy_responderglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ResponderService) CountResponderPolicyResponderGlobalBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", responderPolicyResponderGlobalBindingURL, url.QueryEscape(name))
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
		} `json:"responderpolicy_responderglobal_binding"`
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

// responderpolicy_responderpolicylabel_binding

func (s *ResponderService) GetAllResponderPolicyResponderPolicyLabelBinding() ([]models.ResponderPolicyResponderPolicyLabelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, responderPolicyResponderPolicyLabelBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ResponderPolicyResponderPolicyLabelBinding `json:"responderpolicy_responderpolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ResponderService) GetResponderPolicyResponderPolicyLabelBinding(name string) ([]models.ResponderPolicyResponderPolicyLabelBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", responderPolicyResponderPolicyLabelBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ResponderPolicyResponderPolicyLabelBinding `json:"responderpolicy_responderpolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ResponderService) CountResponderPolicyResponderPolicyLabelBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", responderPolicyResponderPolicyLabelBindingURL, url.QueryEscape(name))
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
		} `json:"responderpolicy_responderpolicylabel_binding"`
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

// responderpolicy_vpnvserver_binding

func (s *ResponderService) GetAllResponderPolicyVPNVServerBinding() ([]models.ResponderPolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, responderPolicyVPNVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ResponderPolicyVPNVServerBinding `json:"responderpolicy_vpnvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ResponderService) GetResponderPolicyVPNVServerBinding(name string) ([]models.ResponderPolicyVPNVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", responderPolicyVPNVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ResponderPolicyVPNVServerBinding `json:"responderpolicy_vpnvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ResponderService) CountResponderPolicyVPNVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", responderPolicyVPNVServerBindingURL, url.QueryEscape(name))
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
		} `json:"responderpolicy_vpnvserver_binding"`
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
