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
	contentInspectionActionURL                                    = "/nitro/v1/config/contentinspectionaction"
	contentInspectionCalloutURL                                   = "/nitro/v1/config/contentinspectioncallout"
	contentInspectionGlobalBindingURL                             = "/nitro/v1/config/contentinspectionglobal_binding"
	contentInspectionGlobalContentInspectionPolicyBindingURL      = "/nitro/v1/config/contentinspectionglobal_contentinspectionpolicy_binding"
	contentInspectionParameterURL                                 = "/nitro/v1/config/contentinspectionparameter"
	contentInspectionPolicyURL                                    = "/nitro/v1/config/contentinspectionpolicy"
	contentInspectionPolicyLabelURL                               = "/nitro/v1/config/contentinspectionpolicylabel"
	contentInspectionPolicyLabelBindingURL                        = "/nitro/v1/config/contentinspectionpolicylabel_binding"
	contentInspectionPolicyLabelContentInspectionPolicyBindingURL = "/nitro/v1/config/contentinspectionpolicylabel_contentinspectionpolicy_binding"
	contentInspectionPolicyLabelPolicyBindingBindingURL           = "/nitro/v1/config/contentinspectionpolicylabel_policybinding_binding"
	contentInspectionPolicyBindingURL                             = "/nitro/v1/config/contentinspectionpolicy_binding"
	contentInspectionPolicyContentInspectionGlobalBindingURL      = "/nitro/v1/config/contentinspectionpolicy_contentinspectionglobal_binding"
	contentInspectionPolicyContentInspectionPolicyLabelBindingURL = "/nitro/v1/config/contentinspectionpolicy_contentinspectionpolicylabel_binding"
	contentInspectionPolicyCSVServerBindingURL                    = "/nitro/v1/config/contentinspectionpolicy_csvserver_binding"
	contentInspectionPolicyLBVServerBindingURL                    = "/nitro/v1/config/contentinspectionpolicy_lbvserver_binding"
	contentInspectionProfileURL                                   = "/nitro/v1/config/contentinspectionprofile"
)

// Content Inspection
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/contentinspection/contentinspection
type ContentInspectionService struct {
	client *Client
}

// contentinspectionaction
func (s *ContentInspectionService) AddContentInspectionAction(resource models.ContentInspectionAction) error {
	payload := map[string]any{"contentinspectionaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, contentInspectionActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ContentInspectionService) DeleteContentInspectionAction(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", contentInspectionActionURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ContentInspectionService) UpdateContentInspectionAction(resource models.ContentInspectionAction) error {
	payload := map[string]any{"contentinspectionaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, contentInspectionActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ContentInspectionService) UnsetContentInspectionAction(resource models.ContentInspectionAction) error {
	payload := map[string]any{"contentinspectionaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", contentInspectionActionURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ContentInspectionService) GetAllContentInspectionAction() ([]models.ContentInspectionAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, contentInspectionActionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.ContentInspectionAction `json:"contentinspectionaction"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *ContentInspectionService) GetContentInspectionAction(name string) (models.ContentInspectionAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", contentInspectionActionURL, name), nil)
	if err != nil {
		return models.ContentInspectionAction{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.ContentInspectionAction{}, err
	}

	var result struct {
		Data []models.ContentInspectionAction `json:"contentinspectionaction"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.ContentInspectionAction{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.ContentInspectionAction{}, fmt.Errorf("contentinspectionaction %s not found", name)
	}

	return result.Data[0], nil
}

func (s *ContentInspectionService) CountContentInspectionAction() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", contentInspectionActionURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.ContentInspectionAction `json:"contentinspectionaction"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) > 0 {
		return int(result.Data[0].Count), nil
	}
	return 0, fmt.Errorf("failed to parse count")
}

// contentinspectioncallout
func (s *ContentInspectionService) AddContentInspectionCallout(resource models.ContentInspectionCallout) error {
	payload := map[string]any{"contentinspectioncallout": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, contentInspectionCalloutURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ContentInspectionService) DeleteContentInspectionCallout(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", contentInspectionCalloutURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ContentInspectionService) UpdateContentInspectionCallout(resource models.ContentInspectionCallout) error {
	payload := map[string]any{"contentinspectioncallout": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, contentInspectionCalloutURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ContentInspectionService) UnsetContentInspectionCallout(resource models.ContentInspectionCallout) error {
	payload := map[string]any{"contentinspectioncallout": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", contentInspectionCalloutURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ContentInspectionService) GetAllContentInspectionCallout() ([]models.ContentInspectionCallout, error) {
	req, err := s.client.NewRequest(http.MethodGet, contentInspectionCalloutURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.ContentInspectionCallout `json:"contentinspectioncallout"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *ContentInspectionService) GetContentInspectionCallout(name string) (models.ContentInspectionCallout, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", contentInspectionCalloutURL, name), nil)
	if err != nil {
		return models.ContentInspectionCallout{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.ContentInspectionCallout{}, err
	}

	var result struct {
		Data []models.ContentInspectionCallout `json:"contentinspectioncallout"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.ContentInspectionCallout{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.ContentInspectionCallout{}, fmt.Errorf("contentinspectioncallout %s not found", name)
	}

	return result.Data[0], nil
}

func (s *ContentInspectionService) CountContentInspectionCallout() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", contentInspectionCalloutURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.ContentInspectionCallout `json:"contentinspectioncallout"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) > 0 {
		return int(result.Data[0].Count), nil
	}
	return 0, fmt.Errorf("failed to parse count")
}

// contentinspectionglobal_binding
func (s *ContentInspectionService) GetContentInspectionGlobalBinding() (models.ContentInspectionGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, contentInspectionGlobalBindingURL, nil)
	if err != nil {
		return models.ContentInspectionGlobalBinding{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.ContentInspectionGlobalBinding{}, err
	}

	var result struct {
		Data []models.ContentInspectionGlobalBinding `json:"contentinspectionglobal_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.ContentInspectionGlobalBinding{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.ContentInspectionGlobalBinding{}, fmt.Errorf("contentinspectionglobal_binding not found")
	}

	return result.Data[0], nil
}

// contentinspectionglobal_contentinspectionpolicy_binding
func (s *ContentInspectionService) AddContentInspectionGlobalContentInspectionPolicyBinding(resource models.ContentInspectionGlobalContentInspectionPolicyBinding) error {
	payload := map[string]any{"contentinspectionglobal_contentinspectionpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, contentInspectionGlobalContentInspectionPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ContentInspectionService) DeleteContentInspectionGlobalContentInspectionPolicyBinding(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", contentInspectionGlobalContentInspectionPolicyBindingURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ContentInspectionService) GetContentInspectionGlobalContentInspectionPolicyBinding() ([]models.ContentInspectionGlobalContentInspectionPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, contentInspectionGlobalContentInspectionPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.ContentInspectionGlobalContentInspectionPolicyBinding `json:"contentinspectionglobal_contentinspectionpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// Struct lacks Count field, leaving for later
func (s *ContentInspectionService) CountContentInspectionGlobalContentInspectionPolicyBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, contentInspectionGlobalContentInspectionPolicyBindingURL+"?count=yes", nil)
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
		} `json:"contentinspectionglobal_contentinspectionpolicy_binding"`
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

// contentinspectionparameter
func (s *ContentInspectionService) UpdateContentInspectionParameter(resource models.ContentInspectionParameter) error {
	payload := map[string]any{"contentinspectionparameter": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, contentInspectionParameterURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ContentInspectionService) UnsetContentInspectionParameter(resource models.ContentInspectionParameter) error {
	payload := map[string]any{"contentinspectionparameter": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", contentInspectionParameterURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ContentInspectionService) GetAllContentInspectionParameter() (models.ContentInspectionParameter, error) {
	req, err := s.client.NewRequest(http.MethodGet, contentInspectionParameterURL, nil)
	if err != nil {
		return models.ContentInspectionParameter{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.ContentInspectionParameter{}, err
	}

	var result struct {
		Data models.ContentInspectionParameter `json:"contentinspectionparameter"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.ContentInspectionParameter{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// contentinspectionpolicy
func (s *ContentInspectionService) AddContentInspectionPolicy(resource models.ContentInspectionPolicy) error {
	payload := map[string]any{"contentinspectionpolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, contentInspectionPolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ContentInspectionService) DeleteContentInspectionPolicy(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", contentInspectionPolicyURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ContentInspectionService) UpdateContentInspectionPolicy(resource models.ContentInspectionPolicy) error {
	payload := map[string]any{"contentinspectionpolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, contentInspectionPolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ContentInspectionService) UnsetContentInspectionPolicy(resource models.ContentInspectionPolicy) error {
	payload := map[string]any{"contentinspectionpolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", contentInspectionPolicyURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ContentInspectionService) GetAllContentInspectionPolicy() ([]models.ContentInspectionPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, contentInspectionPolicyURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.ContentInspectionPolicy `json:"contentinspectionpolicy"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *ContentInspectionService) GetContentInspectionPolicy(name string) (models.ContentInspectionPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", contentInspectionPolicyURL, name), nil)
	if err != nil {
		return models.ContentInspectionPolicy{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.ContentInspectionPolicy{}, err
	}

	var result struct {
		Data []models.ContentInspectionPolicy `json:"contentinspectionpolicy"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.ContentInspectionPolicy{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.ContentInspectionPolicy{}, fmt.Errorf("contentinspectionpolicy %s not found", name)
	}

	return result.Data[0], nil
}

func (s *ContentInspectionService) CountContentInspectionPolicy() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", contentInspectionPolicyURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.ContentInspectionPolicy `json:"contentinspectionpolicy"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) > 0 {
		return int(result.Data[0].Count), nil
	}
	return 0, fmt.Errorf("failed to parse count")
}

func (s *ContentInspectionService) RenameContentInspectionPolicy(resource models.ContentInspectionPolicy) error {
	payload := map[string]any{"contentinspectionpolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=rename", contentInspectionPolicyURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// contentinspectionpolicylabel
func (s *ContentInspectionService) AddContentInspectionPolicyLabel(resource models.ContentInspectionPolicyLabel) error {
	payload := map[string]any{"contentinspectionpolicylabel": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, contentInspectionPolicyLabelURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ContentInspectionService) DeleteContentInspectionPolicyLabel(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", contentInspectionPolicyLabelURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ContentInspectionService) GetAllContentInspectionPolicyLabel() ([]models.ContentInspectionPolicyLabel, error) {
	req, err := s.client.NewRequest(http.MethodGet, contentInspectionPolicyLabelURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.ContentInspectionPolicyLabel `json:"contentinspectionpolicylabel"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *ContentInspectionService) GetContentInspectionPolicyLabel(name string) (models.ContentInspectionPolicyLabel, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", contentInspectionPolicyLabelURL, name), nil)
	if err != nil {
		return models.ContentInspectionPolicyLabel{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.ContentInspectionPolicyLabel{}, err
	}

	var result struct {
		Data []models.ContentInspectionPolicyLabel `json:"contentinspectionpolicylabel"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.ContentInspectionPolicyLabel{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.ContentInspectionPolicyLabel{}, fmt.Errorf("contentinspectionpolicylabel %s not found", name)
	}

	return result.Data[0], nil
}

func (s *ContentInspectionService) CountContentInspectionPolicyLabel() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", contentInspectionPolicyLabelURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.ContentInspectionPolicyLabel `json:"contentinspectionpolicylabel"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) > 0 {
		return int(result.Data[0].Count), nil
	}
	return 0, fmt.Errorf("failed to parse count")
}

func (s *ContentInspectionService) RenameContentInspectionPolicyLabel(resource models.ContentInspectionPolicyLabel) error {
	payload := map[string]any{"contentinspectionpolicylabel": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=rename", contentInspectionPolicyLabelURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// contentinspectionpolicylabel_binding
func (s *ContentInspectionService) GetAllContentInspectionPolicyLabelBinding() ([]models.ContentInspectionPolicyLabelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, contentInspectionPolicyLabelBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.ContentInspectionPolicyLabelBinding `json:"contentinspectionpolicylabel_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *ContentInspectionService) GetContentInspectionPolicyLabelBinding(name string) (models.ContentInspectionPolicyLabelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", contentInspectionPolicyLabelBindingURL, name), nil)
	if err != nil {
		return models.ContentInspectionPolicyLabelBinding{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.ContentInspectionPolicyLabelBinding{}, err
	}

	var result struct {
		Data []models.ContentInspectionPolicyLabelBinding `json:"contentinspectionpolicylabel_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.ContentInspectionPolicyLabelBinding{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.ContentInspectionPolicyLabelBinding{}, fmt.Errorf("contentinspectionpolicylabel_binding %s not found", name)
	}

	return result.Data[0], nil
}

// contentinspectionpolicylabel_contentinspectionpolicy_binding
func (s *ContentInspectionService) AddContentInspectionPolicyLabelContentInspectionPolicyBinding(resource models.ContentInspectionPolicyLabelContentInspectionPolicyBinding) error {
	payload := map[string]any{"contentinspectionpolicylabel_contentinspectionpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, contentInspectionPolicyLabelContentInspectionPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ContentInspectionService) DeleteContentInspectionPolicyLabelContentInspectionPolicyBinding(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", contentInspectionPolicyLabelContentInspectionPolicyBindingURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ContentInspectionService) GetAllContentInspectionPolicyLabelContentInspectionPolicyBinding() ([]models.ContentInspectionPolicyLabelContentInspectionPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, contentInspectionPolicyLabelContentInspectionPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.ContentInspectionPolicyLabelContentInspectionPolicyBinding `json:"contentinspectionpolicylabel_contentinspectionpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *ContentInspectionService) GetContentInspectionPolicyLabelContentInspectionPolicyBinding(name string) (models.ContentInspectionPolicyLabelContentInspectionPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", contentInspectionPolicyLabelContentInspectionPolicyBindingURL, name), nil)
	if err != nil {
		return models.ContentInspectionPolicyLabelContentInspectionPolicyBinding{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.ContentInspectionPolicyLabelContentInspectionPolicyBinding{}, err
	}

	var result struct {
		Data []models.ContentInspectionPolicyLabelContentInspectionPolicyBinding `json:"contentinspectionpolicylabel_contentinspectionpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.ContentInspectionPolicyLabelContentInspectionPolicyBinding{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.ContentInspectionPolicyLabelContentInspectionPolicyBinding{}, fmt.Errorf("contentinspectionpolicylabel_contentinspectionpolicy_binding %s not found", name)
	}

	return result.Data[0], nil
}

// Struct lacks Count field, leaving for later
func (s *ContentInspectionService) CountContentInspectionPolicyLabelContentInspectionPolicyBinding() {
}

// contentinspectionpolicylabel_policybinding_binding
func (s *ContentInspectionService) GetAllContentInspectionPolicyLabelPolicyBindingBinding() ([]models.ContentInspectionPolicyLabelPolicyBindingBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, contentInspectionPolicyLabelPolicyBindingBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.ContentInspectionPolicyLabelPolicyBindingBinding `json:"contentinspectionpolicylabel_policybinding_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *ContentInspectionService) GetContentInspectionPolicyLabelPolicyBindingBinding(name string) (models.ContentInspectionPolicyLabelPolicyBindingBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", contentInspectionPolicyLabelPolicyBindingBindingURL, name), nil)
	if err != nil {
		return models.ContentInspectionPolicyLabelPolicyBindingBinding{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.ContentInspectionPolicyLabelPolicyBindingBinding{}, err
	}

	var result struct {
		Data []models.ContentInspectionPolicyLabelPolicyBindingBinding `json:"contentinspectionpolicylabel_policybinding_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.ContentInspectionPolicyLabelPolicyBindingBinding{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.ContentInspectionPolicyLabelPolicyBindingBinding{}, fmt.Errorf("contentinspectionpolicylabel_policybinding_binding %s not found", name)
	}

	return result.Data[0], nil
}

// Struct lacks Count field, leaving for later
func (s *ContentInspectionService) CountContentInspectionPolicyLabelPolicyBindingBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, contentInspectionPolicyLabelPolicyBindingBindingURL+"?count=yes", nil)
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
		} `json:"contentinspectionpolicylabel_policybinding_binding"`
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

// contentinspectionpolicy_binding
func (s *ContentInspectionService) GetAllContentInspectionPolicyBinding() ([]models.ContentInspectionPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, contentInspectionPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.ContentInspectionPolicyBinding `json:"contentinspectionpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *ContentInspectionService) GetContentInspectionPolicyBinding(name string) (models.ContentInspectionPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", contentInspectionPolicyBindingURL, name), nil)
	if err != nil {
		return models.ContentInspectionPolicyBinding{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.ContentInspectionPolicyBinding{}, err
	}

	var result struct {
		Data []models.ContentInspectionPolicyBinding `json:"contentinspectionpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.ContentInspectionPolicyBinding{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.ContentInspectionPolicyBinding{}, fmt.Errorf("contentinspectionpolicy_binding %s not found", name)
	}

	return result.Data[0], nil
}

// contentinspectionpolicy_contentinspectionglobal_binding
func (s *ContentInspectionService) GetAllContentInspectionPolcyContentInspectionGlobalBinding() ([]models.ContentInspectionPolicyContentInspectionGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, contentInspectionPolicyContentInspectionGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.ContentInspectionPolicyContentInspectionGlobalBinding `json:"contentinspectionpolicy_contentinspectionglobal_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *ContentInspectionService) GetContentInspectionPolcyContentInspectionGlobalBinding(name string) (models.ContentInspectionPolicyContentInspectionGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", contentInspectionPolicyContentInspectionGlobalBindingURL, name), nil)
	if err != nil {
		return models.ContentInspectionPolicyContentInspectionGlobalBinding{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.ContentInspectionPolicyContentInspectionGlobalBinding{}, err
	}

	var result struct {
		Data []models.ContentInspectionPolicyContentInspectionGlobalBinding `json:"contentinspectionpolicy_contentinspectionglobal_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.ContentInspectionPolicyContentInspectionGlobalBinding{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.ContentInspectionPolicyContentInspectionGlobalBinding{}, fmt.Errorf("contentinspectionpolicy_contentinspectionglobal_binding %s not found", name)
	}

	return result.Data[0], nil
}

// Struct lacks Count field, leaving for later
func (s *ContentInspectionService) CountContentInspectionPolcyContentInspectionGlobalBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", contentInspectionPolicyContentInspectionGlobalBindingURL, url.PathEscape(name))
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
		} `json:"contentinspectionpolicy_contentinspectionglobal_binding"`
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

// contentinspectionpolicy_contentinspectionpolicylabel_binding
func (s *ContentInspectionService) GetAllContentInspectionPolicyContentInspectionPolicyLabelBinding() ([]models.ContentInspectionPolicyContentInspectionPolicyLabelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, contentInspectionPolicyContentInspectionPolicyLabelBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.ContentInspectionPolicyContentInspectionPolicyLabelBinding `json:"contentinspectionpolicy_contentinspectionpolicylabel_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *ContentInspectionService) GetContentInspectionPolicyContentInspectionPolicyLabelBinding(name string) (models.ContentInspectionPolicyContentInspectionPolicyLabelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", contentInspectionPolicyContentInspectionPolicyLabelBindingURL, name), nil)
	if err != nil {
		return models.ContentInspectionPolicyContentInspectionPolicyLabelBinding{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.ContentInspectionPolicyContentInspectionPolicyLabelBinding{}, err
	}

	var result struct {
		Data []models.ContentInspectionPolicyContentInspectionPolicyLabelBinding `json:"contentinspectionpolicy_contentinspectionpolicylabel_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.ContentInspectionPolicyContentInspectionPolicyLabelBinding{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.ContentInspectionPolicyContentInspectionPolicyLabelBinding{}, fmt.Errorf("contentinspectionpolicy_contentinspectionpolicylabel_binding %s not found", name)
	}

	return result.Data[0], nil
}

// Struct lacks Count field, leaving for later
func (s *ContentInspectionService) CountContentInspectionPolicyContentInspectionPolicyLabelBinding() {
}

// contentinspectionpolicy_csvserver_binding
func (s *ContentInspectionService) GetAllContentInspectionPolicyCSVServerBinding() ([]models.ContentInspectionPolicyCSVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, contentInspectionPolicyCSVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.ContentInspectionPolicyCSVServerBinding `json:"contentinspectionpolicy_csvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *ContentInspectionService) GetContentInspectionPolicyCSVServerBinding(name string) (models.ContentInspectionPolicyCSVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", contentInspectionPolicyCSVServerBindingURL, name), nil)
	if err != nil {
		return models.ContentInspectionPolicyCSVServerBinding{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.ContentInspectionPolicyCSVServerBinding{}, err
	}

	var result struct {
		Data []models.ContentInspectionPolicyCSVServerBinding `json:"contentinspectionpolicy_csvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.ContentInspectionPolicyCSVServerBinding{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.ContentInspectionPolicyCSVServerBinding{}, fmt.Errorf("contentinspectionpolicy_csvserver_binding %s not found", name)
	}

	return result.Data[0], nil
}

// Struct lacks Count field, leaving for later
func (s *ContentInspectionService) CountContentInspectionPolicyCSVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", contentInspectionPolicyCSVServerBindingURL, url.PathEscape(name))
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
		} `json:"contentinspectionpolicy_csvserver_binding"`
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

// contentinspectionpolicy_lbvserver_binding
func (s *ContentInspectionService) GetAllContentInspectionPolcyLBVServerBinding() ([]models.ContentInspectionPolicyLBVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, contentInspectionPolicyLBVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.ContentInspectionPolicyLBVServerBinding `json:"contentinspectionpolicy_lbvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *ContentInspectionService) GetContentInspectionPolcyLBVServerBinding(name string) (models.ContentInspectionPolicyLBVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", contentInspectionPolicyLBVServerBindingURL, name), nil)
	if err != nil {
		return models.ContentInspectionPolicyLBVServerBinding{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.ContentInspectionPolicyLBVServerBinding{}, err
	}

	var result struct {
		Data []models.ContentInspectionPolicyLBVServerBinding `json:"contentinspectionpolicy_lbvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.ContentInspectionPolicyLBVServerBinding{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.ContentInspectionPolicyLBVServerBinding{}, fmt.Errorf("contentinspectionpolicy_lbvserver_binding %s not found", name)
	}

	return result.Data[0], nil
}

// Struct lacks Count field, leaving for later
func (s *ContentInspectionService) CountContentInspectionPolcyLBVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", contentInspectionPolicyLBVServerBindingURL, url.PathEscape(name))
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
		} `json:"contentinspectionpolicy_lbvserver_binding"`
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

// contentinspectionprofile
func (s *ContentInspectionService) AddContentInspectionProfile(resource models.ContentInspectionProfile) error {
	payload := map[string]any{"contentinspectionprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, contentInspectionProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ContentInspectionService) DeleteContentInspectionProfile(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", contentInspectionProfileURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ContentInspectionService) UpdateContentInspectionProfile(resource models.ContentInspectionProfile) error {
	payload := map[string]any{"contentinspectionprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, contentInspectionProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ContentInspectionService) UnsetContentInspectionProfile(resource models.ContentInspectionProfile) error {
	payload := map[string]any{"contentinspectionprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", contentInspectionProfileURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *ContentInspectionService) GetAllContentInspectionProfile() ([]models.ContentInspectionProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, contentInspectionProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.ContentInspectionProfile `json:"contentinspectionprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *ContentInspectionService) GetContentInspectionProfile(name string) (models.ContentInspectionProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", contentInspectionProfileURL, name), nil)
	if err != nil {
		return models.ContentInspectionProfile{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.ContentInspectionProfile{}, err
	}

	var result struct {
		Data []models.ContentInspectionProfile `json:"contentinspectionprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.ContentInspectionProfile{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.ContentInspectionProfile{}, fmt.Errorf("contentinspectionprofile %s not found", name)
	}

	return result.Data[0], nil
}

func (s *ContentInspectionService) CountContentInspectionProfile() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", contentInspectionProfileURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.ContentInspectionProfile `json:"contentinspectionprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) > 0 {
		return int(result.Data[0].Count), nil
	}
	return 0, fmt.Errorf("failed to parse count")
}
