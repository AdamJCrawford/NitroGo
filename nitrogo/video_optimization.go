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
	videoOptimizationDetectionActionURL                                             = "/nitro/v1/config/videooptimizationdetectionaction"
	videoOptimizationDetectionPolicyURL                                             = "/nitro/v1/config/videooptimizationdetectionpolicy"
	videoOptimizationDetectionPolicyLabelURL                                        = "/nitro/v1/config/videooptimizationdetectionpolicylabel"
	videoOptimizationDetectionPolicyLabelBindingURL                                 = "/nitro/v1/config/videooptimizationdetectionpolicylabel_binding"
	videoOptimizationDetectionPolicyLabelPolicyBindingBindingURL                    = "/nitro/v1/config/videooptimizationdetectionpolicylabel_policybinding_binding"
	videoOptimizationDetectionPolicyLabelVideoOptimizationDetectionPolicyBindingURL = "/nitro/v1/config/videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding"
	videoOptimizationDetectionPolicyBindingURL                                      = "/nitro/v1/config/videooptimizationdetectionpolicy_binding"
	videoOptimizationDetectionPolicyLBVServerBindingURL                             = "/nitro/v1/config/videooptimizationdetectionpolicy_lbvserver_binding"
	videoOptimizationDetectionPolicyVideoOptimizationGlobalDetectionBindingURL      = "/nitro/v1/config/videooptimizationdetectionpolicy_videooptimizationglobaldetection_binding"
	videoOptimizationGlobalDetectionBindingURL                                      = "/nitro/v1/config/videooptimizationglobaldetection_binding"
	videoOptimizationGlobalDetectionVideoOptimizationDetectionPolicyBindingURL      = "/nitro/v1/config/videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding"
	videoOptimizationGlobalPacingBindingURL                                         = "/nitro/v1/config/videooptimizationglobalpacing_binding"
	videoOptimizationGlobalPacingVideoOptimizationPacingPolicyBindingURL            = "/nitro/v1/config/videooptimizationglobalpacing_videooptimizationpacingpolicy_binding"
	videoOptimizationPacingActionURL                                                = "/nitro/v1/config/videooptimizationpacingaction"
	videoOptimizationPacingPolicyURL                                                = "/nitro/v1/config/videooptimizationpacingpolicy"
	videoOptimizationPacingPolicyLabelURL                                           = "/nitro/v1/config/videooptimizationpacingpolicylabel"
	videoOptimizationPacingPolicyLabelBindingURL                                    = "/nitro/v1/config/videooptimizationpacingpolicylabel_binding"
	videoOptimizationPacingPolicyLabelPolicyBindingBindingURL                       = "/nitro/v1/config/videooptimizationpacingpolicylabel_policybinding_binding"
	videoOptimizationPacingPolicyLabelVideoOptimizationPacingPolicyBindingURL       = "/nitro/v1/config/videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding"
	videoOptimizationPacingPolicyBindingURL                                         = "/nitro/v1/config/videooptimizationpacingpolicy_binding"
	videoOptimizationPacingPolicyLBVServerBindingURL                                = "/nitro/v1/config/videooptimizationpacingpolicy_lbvserver_binding"
	videoOptimizationPacingPolicyVideoOptimizationGlobalPacingBindingURL            = "/nitro/v1/config/videooptimizationpacingpolicy_videooptimizationglobalpacing_binding"
	videoOptimizationParameterURL                                                   = "/nitro/v1/config/videooptimizationparameter"
)

// VideoOptimization
// Video optimization feature is used to show (i) the stats of different media types that are being served by the Citrix ADC and (ii) the details of optimization applied on ABR videos
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/videooptimization/videooptimization
type VideoOptimizationService struct {
	client *Client
}

// videooptimizationdetectionaction
func (s *VideoOptimizationService) AddVideoOptimizationDetectionAction(resource models.VideoOptimizationDetectionAction) error {
	payload := map[string]any{
		"videooptimizationdetectionaction": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, videoOptimizationDetectionActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VideoOptimizationService) DeleteVideoOptimizationDetectionAction(name string) error {
	reqURL := fmt.Sprintf("%s/%s", videoOptimizationDetectionActionURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VideoOptimizationService) UpdateVideoOptimizationDetectionAction(resource models.VideoOptimizationDetectionAction) error {
	payload := map[string]any{
		"videooptimizationdetectionaction": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, videoOptimizationDetectionActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VideoOptimizationService) UnsetVideoOptimizationDetectionAction(resource models.VideoOptimizationDetectionAction) error {
	payload := map[string]any{
		"videooptimizationdetectionaction": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, videoOptimizationDetectionActionURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VideoOptimizationService) RenameVideoOptimizationDetectionAction(resource models.VideoOptimizationDetectionAction) error {
	payload := map[string]any{
		"videooptimizationdetectionaction": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, videoOptimizationDetectionActionURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VideoOptimizationService) GetAllVideoOptimizationDetectionAction() ([]models.VideoOptimizationDetectionAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, videoOptimizationDetectionActionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Actions []models.VideoOptimizationDetectionAction `json:"videooptimizationdetectionaction"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Actions, nil
}

func (s *VideoOptimizationService) GetVideoOptimizationDetectionAction(name string) (*models.VideoOptimizationDetectionAction, error) {
	reqURL := fmt.Sprintf("%s/%s", videoOptimizationDetectionActionURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Actions []models.VideoOptimizationDetectionAction `json:"videooptimizationdetectionaction"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Actions) == 0 {
		return nil, fmt.Errorf("videooptimizationdetectionaction not found")
	}

	return &result.Actions[0], nil
}

func (s *VideoOptimizationService) CountVideoOptimizationDetectionAction() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, videoOptimizationDetectionActionURL+"?count=yes", nil)
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
		} `json:"videooptimizationdetectionaction"`
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

// videooptimizationdetectionpolicy
func (s *VideoOptimizationService) AddVideoOptimizationDetectionPolicy(resource models.VideoOptimizationDetectionPolicy) error {
	payload := map[string]any{
		"videooptimizationdetectionpolicy": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, videoOptimizationDetectionPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VideoOptimizationService) DeleteVideoOptimizationDetectionPolicy(name string) error {
	reqURL := fmt.Sprintf("%s/%s", videoOptimizationDetectionPolicyURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VideoOptimizationService) UpdateVideoOptimizationDetectionPolicy(resource models.VideoOptimizationDetectionPolicy) error {
	payload := map[string]any{
		"videooptimizationdetectionpolicy": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, videoOptimizationDetectionPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VideoOptimizationService) UnsetVideoOptimizationDetectionPolicy(resource models.VideoOptimizationDetectionPolicy) error {
	payload := map[string]any{
		"videooptimizationdetectionpolicy": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, videoOptimizationDetectionPolicyURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VideoOptimizationService) RenameVideoOptimizationDetectionPolicy(resource models.VideoOptimizationDetectionPolicy) error {
	payload := map[string]any{
		"videooptimizationdetectionpolicy": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, videoOptimizationDetectionPolicyURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VideoOptimizationService) GetAllVideoOptimizationDetectionPolicy() ([]models.VideoOptimizationDetectionPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, videoOptimizationDetectionPolicyURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Policies []models.VideoOptimizationDetectionPolicy `json:"videooptimizationdetectionpolicy"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Policies, nil
}

func (s *VideoOptimizationService) GetVideoOptimizationDetectionPolicy(name string) (*models.VideoOptimizationDetectionPolicy, error) {
	reqURL := fmt.Sprintf("%s/%s", videoOptimizationDetectionPolicyURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Policies []models.VideoOptimizationDetectionPolicy `json:"videooptimizationdetectionpolicy"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Policies) == 0 {
		return nil, fmt.Errorf("videooptimizationdetectionpolicy not found")
	}

	return &result.Policies[0], nil
}

func (s *VideoOptimizationService) CountVideoOptimizationDetectionPolicy() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, videoOptimizationDetectionPolicyURL+"?count=yes", nil)
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
		} `json:"videooptimizationdetectionpolicy"`
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

// videooptimizationdetectionpolicylabel
func (s *VideoOptimizationService) AddVideoOptimizationDetectionPolicyLabel(resource models.VideoOptimizationDetectionPolicyLabel) error {
	payload := map[string]any{
		"videooptimizationdetectionpolicylabel": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, videoOptimizationDetectionPolicyLabelURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VideoOptimizationService) DeleteVideoOptimizationDetectionPolicyLabel(name string) error {
	reqURL := fmt.Sprintf("%s/%s", videoOptimizationDetectionPolicyLabelURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VideoOptimizationService) RenameVideoOptimizationDetectionPolicyLabel(resource models.VideoOptimizationDetectionPolicyLabel) error {
	payload := map[string]any{
		"videooptimizationdetectionpolicylabel": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, videoOptimizationDetectionPolicyLabelURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VideoOptimizationService) GetAllVideoOptimizationDetectionPolicyLabel() ([]models.VideoOptimizationDetectionPolicyLabel, error) {
	req, err := s.client.NewRequest(http.MethodGet, videoOptimizationDetectionPolicyLabelURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Labels []models.VideoOptimizationDetectionPolicyLabel `json:"videooptimizationdetectionpolicylabel"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Labels, nil
}

func (s *VideoOptimizationService) GetVideoOptimizationDetectionPolicyLabel(name string) (*models.VideoOptimizationDetectionPolicyLabel, error) {
	reqURL := fmt.Sprintf("%s/%s", videoOptimizationDetectionPolicyLabelURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Labels []models.VideoOptimizationDetectionPolicyLabel `json:"videooptimizationdetectionpolicylabel"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Labels) == 0 {
		return nil, fmt.Errorf("videooptimizationdetectionpolicylabel not found")
	}

	return &result.Labels[0], nil
}

func (s *VideoOptimizationService) CountVideoOptimizationDetectionPolicyLabel() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, videoOptimizationDetectionPolicyLabelURL+"?count=yes", nil)
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
		} `json:"videooptimizationdetectionpolicylabel"`
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

// videooptimizationdetectionpolicylabel_binding
func (s *VideoOptimizationService) GetAllVideoOptimizationDetectionPolicyLabelBinding() ([]models.VideoOptimizationDetectionPolicyLabelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, videoOptimizationDetectionPolicyLabelBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.VideoOptimizationDetectionPolicyLabelBinding `json:"videooptimizationdetectionpolicylabel_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *VideoOptimizationService) GetVideoOptimizationDetectionPolicyLabelBinding(name string) (*models.VideoOptimizationDetectionPolicyLabelBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", videoOptimizationDetectionPolicyLabelBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.VideoOptimizationDetectionPolicyLabelBinding `json:"videooptimizationdetectionpolicylabel_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) == 0 {
		return nil, fmt.Errorf("videooptimizationdetectionpolicylabel_binding not found")
	}

	return &result.Bindings[0], nil
}

// videooptimizationdetectionpolicylabel_policybinding_binding
func (s *VideoOptimizationService) GetAllVideoOptimizationDetectionPolicyLabelPolicyBindingBinding() ([]models.VideoOptimizationDetectionPolicyLabelPolicyBindingBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, videoOptimizationDetectionPolicyLabelPolicyBindingBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.VideoOptimizationDetectionPolicyLabelPolicyBindingBinding `json:"videooptimizationdetectionpolicylabel_policybinding_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *VideoOptimizationService) GetVideoOptimizationDetectionPolicyLabelPolicyBindingBinding(name string) ([]models.VideoOptimizationDetectionPolicyLabelPolicyBindingBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", videoOptimizationDetectionPolicyLabelPolicyBindingBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.VideoOptimizationDetectionPolicyLabelPolicyBindingBinding `json:"videooptimizationdetectionpolicylabel_policybinding_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *VideoOptimizationService) CountVideoOptimizationDetectionPolicyLabelPolicyBindingBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", videoOptimizationDetectionPolicyLabelPolicyBindingBindingURL, url.PathEscape(name))
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
		} `json:"videooptimizationdetectionpolicylabel_policybinding_binding"`
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

// videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding
func (s *VideoOptimizationService) AddVideoOptimizationDetectionPolicyLabelVideoOptimizationDetectionPolicyBinding(resource models.VideoOptimizationDetectionPolicyLabelVideoOptimizationDetectionPolicyBinding) error {
	payload := map[string]any{
		"videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, videoOptimizationDetectionPolicyLabelVideoOptimizationDetectionPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VideoOptimizationService) DeleteVideoOptimizationDetectionPolicyLabelVideoOptimizationDetectionPolicyBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", videoOptimizationDetectionPolicyLabelVideoOptimizationDetectionPolicyBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VideoOptimizationService) GetAllVideoOptimizationDetectionPolicyLabelVideoOptimizationDetectionPolicyBinding() ([]models.VideoOptimizationDetectionPolicyLabelVideoOptimizationDetectionPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, videoOptimizationDetectionPolicyLabelVideoOptimizationDetectionPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.VideoOptimizationDetectionPolicyLabelVideoOptimizationDetectionPolicyBinding `json:"videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *VideoOptimizationService) GetVideoOptimizationDetectionPolicyLabelVideoOptimizationDetectionPolicyBinding(name string) ([]models.VideoOptimizationDetectionPolicyLabelVideoOptimizationDetectionPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", videoOptimizationDetectionPolicyLabelVideoOptimizationDetectionPolicyBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.VideoOptimizationDetectionPolicyLabelVideoOptimizationDetectionPolicyBinding `json:"videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *VideoOptimizationService) CountVideoOptimizationDetectionPolicyLabelVideoOptimizationDetectionPolicyBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", videoOptimizationDetectionPolicyLabelVideoOptimizationDetectionPolicyBindingURL, url.PathEscape(name))
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
		} `json:"videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding"`
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

// videooptimizationdetectionpolicy_binding
func (s *VideoOptimizationService) GetAllVideoOptimizationDetectionPolicyBinding() ([]models.VideoOptimizationDetectionPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, videoOptimizationDetectionPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.VideoOptimizationDetectionPolicyBinding `json:"videooptimizationdetectionpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *VideoOptimizationService) GetVideoOptimizationDetectionPolicyBinding(name string) (*models.VideoOptimizationDetectionPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", videoOptimizationDetectionPolicyBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.VideoOptimizationDetectionPolicyBinding `json:"videooptimizationdetectionpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) == 0 {
		return nil, fmt.Errorf("videooptimizationdetectionpolicy_binding not found")
	}

	return &result.Bindings[0], nil
}

// videooptimizationdetectionpolicy_lbvserver_binding
func (s *VideoOptimizationService) GetAllVideoOptimizationDetectionPolicyLBVServerBinding() ([]models.VideoOptimizationDetectionPolicyLBVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, videoOptimizationDetectionPolicyLBVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.VideoOptimizationDetectionPolicyLBVServerBinding `json:"videooptimizationdetectionpolicy_lbvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *VideoOptimizationService) GetVideoOptimizationDetectionPolicyLBVServerBinding(name string) ([]models.VideoOptimizationDetectionPolicyLBVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", videoOptimizationDetectionPolicyLBVServerBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.VideoOptimizationDetectionPolicyLBVServerBinding `json:"videooptimizationdetectionpolicy_lbvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *VideoOptimizationService) CountVideoOptimizationDetectionPolicyLBVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", videoOptimizationDetectionPolicyLBVServerBindingURL, url.PathEscape(name))
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
		} `json:"videooptimizationdetectionpolicy_lbvserver_binding"`
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

// videooptimizationdetectionpolicy_videooptimizationglobaldetection_binding
func (s *VideoOptimizationService) GetAllVideoOptimizationDetectionPolicyVideoOptimizationGlobalDetectionBinding() ([]models.VideoOptimizationDetectionPolicyVideoOptimizationGlobalDetectionBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, videoOptimizationDetectionPolicyVideoOptimizationGlobalDetectionBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.VideoOptimizationDetectionPolicyVideoOptimizationGlobalDetectionBinding `json:"videooptimizationdetectionpolicy_videooptimizationglobaldetection_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *VideoOptimizationService) GetVideoOptimizationDetectionPolicyVideoOptimizationGlobalDetectionBinding(name string) ([]models.VideoOptimizationDetectionPolicyVideoOptimizationGlobalDetectionBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", videoOptimizationDetectionPolicyVideoOptimizationGlobalDetectionBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.VideoOptimizationDetectionPolicyVideoOptimizationGlobalDetectionBinding `json:"videooptimizationdetectionpolicy_videooptimizationglobaldetection_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *VideoOptimizationService) CountVideoOptimizationDetectionPolicyVideoOptimizationGlobalDetectionBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", videoOptimizationDetectionPolicyVideoOptimizationGlobalDetectionBindingURL, url.PathEscape(name))
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
		} `json:"videooptimizationdetectionpolicy_videooptimizationglobaldetection_binding"`
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

// videooptimizationglobaldetection_binding
func (s *VideoOptimizationService) GetVideoOptimizationGlobalDetectionBinding() (*models.VideoOptimizationGlobalDetectionBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, videoOptimizationGlobalDetectionBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.VideoOptimizationGlobalDetectionBinding `json:"videooptimizationglobaldetection_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) == 0 {
		return nil, fmt.Errorf("videooptimizationglobaldetection_binding not found")
	}

	return &result.Bindings[0], nil
}

// videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding
func (s *VideoOptimizationService) AddVideoOptimizationGlobalDetectionVideoOptimizationDetectionPolicyBinding(resource models.VideoOptimizationGlobalDetectionVideoOptimizationDetectionPolicyBinding) error {
	payload := map[string]any{
		"videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, videoOptimizationGlobalDetectionVideoOptimizationDetectionPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VideoOptimizationService) DeleteVideoOptimizationGlobalDetectionVideoOptimizationDetectionPolicyBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", videoOptimizationGlobalDetectionVideoOptimizationDetectionPolicyBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VideoOptimizationService) GetVideoOptimizationGlobalDetectionVideoOptimizationDetectionPolicyBinding() ([]models.VideoOptimizationGlobalDetectionVideoOptimizationDetectionPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, videoOptimizationGlobalDetectionVideoOptimizationDetectionPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.VideoOptimizationGlobalDetectionVideoOptimizationDetectionPolicyBinding `json:"videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *VideoOptimizationService) CountVideoOptimizationGlobalDetectionVideoOptimizationDetectionPolicyBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, videoOptimizationGlobalDetectionVideoOptimizationDetectionPolicyBindingURL+"?count=yes", nil)
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
		} `json:"videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding"`
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

// videooptimizationglobalpacing_binding
func (s *VideoOptimizationService) GetVideoOptimizationGlobalPacingBinding() (*models.VideoOptimizationGlobalPacingBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, videoOptimizationGlobalPacingBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.VideoOptimizationGlobalPacingBinding `json:"videooptimizationglobalpacing_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) == 0 {
		return nil, fmt.Errorf("videooptimizationglobalpacing_binding not found")
	}

	return &result.Bindings[0], nil
}

// videooptimizationglobalpacing_videooptimizationpacingpolicy_binding
func (s *VideoOptimizationService) AddVideoOptimizationGlobalPacingVideoOptimizationPacingPolicyBinding(resource models.VideoOptimizationGlobalPacingVideoOptimizationPacingPolicyBinding) error {
	payload := map[string]any{
		"videooptimizationglobalpacing_videooptimizationpacingpolicy_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, videoOptimizationGlobalPacingVideoOptimizationPacingPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VideoOptimizationService) DeleteVideoOptimizationGlobalPacingVideoOptimizationPacingPolicyBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", videoOptimizationGlobalPacingVideoOptimizationPacingPolicyBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VideoOptimizationService) GetVideoOptimizationGlobalPacingVideoOptimizationPacingPolicyBinding() ([]models.VideoOptimizationGlobalPacingVideoOptimizationPacingPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, videoOptimizationGlobalPacingVideoOptimizationPacingPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.VideoOptimizationGlobalPacingVideoOptimizationPacingPolicyBinding `json:"videooptimizationglobalpacing_videooptimizationpacingpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *VideoOptimizationService) CountVideoOptimizationGlobalPacingVideoOptimizationPacingPolicyBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, videoOptimizationGlobalPacingVideoOptimizationPacingPolicyBindingURL+"?count=yes", nil)
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
		} `json:"videooptimizationglobalpacing_videooptimizationpacingpolicy_binding"`
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

// videooptimizationpacingaction
func (s *VideoOptimizationService) AddVideoOptimizationPacingAction(resource models.VideoOptimizationPacingAction) error {
	payload := map[string]any{
		"videooptimizationpacingaction": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, videoOptimizationPacingActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VideoOptimizationService) DeleteVideoOptimizationPacingAction(name string) error {
	reqURL := fmt.Sprintf("%s/%s", videoOptimizationPacingActionURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VideoOptimizationService) UpdateVideoOptimizationPacingAction(resource models.VideoOptimizationPacingAction) error {
	payload := map[string]any{
		"videooptimizationpacingaction": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, videoOptimizationPacingActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VideoOptimizationService) UnsetVideoOptimizationPacingAction(resource models.VideoOptimizationPacingAction) error {
	payload := map[string]any{
		"videooptimizationpacingaction": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, videoOptimizationPacingActionURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VideoOptimizationService) RenameVideoOptimizationPacingAction(resource models.VideoOptimizationPacingAction) error {
	payload := map[string]any{
		"videooptimizationpacingaction": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, videoOptimizationPacingActionURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VideoOptimizationService) GetAllVideoOptimizationPacingAction() ([]models.VideoOptimizationPacingAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, videoOptimizationPacingActionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Actions []models.VideoOptimizationPacingAction `json:"videooptimizationpacingaction"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Actions, nil
}

func (s *VideoOptimizationService) GetVideoOptimizationPacingAction(name string) (*models.VideoOptimizationPacingAction, error) {
	reqURL := fmt.Sprintf("%s/%s", videoOptimizationPacingActionURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Actions []models.VideoOptimizationPacingAction `json:"videooptimizationpacingaction"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Actions) == 0 {
		return nil, fmt.Errorf("videooptimizationpacingaction not found")
	}

	return &result.Actions[0], nil
}

func (s *VideoOptimizationService) CountVideoOptimizationPacingAction() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, videoOptimizationPacingActionURL+"?count=yes", nil)
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
		} `json:"videooptimizationpacingaction"`
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

// videooptimizationpacingpolicy
func (s *VideoOptimizationService) AddVideoOptimizationPacingPolicy(resource models.VideoOptimizationPacingPolicy) error {
	payload := map[string]any{
		"videooptimizationpacingpolicy": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, videoOptimizationPacingPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VideoOptimizationService) DeleteVideoOptimizationPacingPolicy(name string) error {
	reqURL := fmt.Sprintf("%s/%s", videoOptimizationPacingPolicyURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VideoOptimizationService) UpdateVideoOptimizationPacingPolicy(resource models.VideoOptimizationPacingPolicy) error {
	payload := map[string]any{
		"videooptimizationpacingpolicy": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, videoOptimizationPacingPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VideoOptimizationService) UnsetVideoOptimizationPacingPolicy(resource models.VideoOptimizationPacingPolicy) error {
	payload := map[string]any{
		"videooptimizationpacingpolicy": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, videoOptimizationPacingPolicyURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VideoOptimizationService) RenameVideoOptimizationPacingPolicy(resource models.VideoOptimizationPacingPolicy) error {
	payload := map[string]any{
		"videooptimizationpacingpolicy": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, videoOptimizationPacingPolicyURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VideoOptimizationService) GetAllVideoOptimizationPacingPolicy() ([]models.VideoOptimizationPacingPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, videoOptimizationPacingPolicyURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Policies []models.VideoOptimizationPacingPolicy `json:"videooptimizationpacingpolicy"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Policies, nil
}

func (s *VideoOptimizationService) GetVideoOptimizationPacingPolicy(name string) (*models.VideoOptimizationPacingPolicy, error) {
	reqURL := fmt.Sprintf("%s/%s", videoOptimizationPacingPolicyURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Policies []models.VideoOptimizationPacingPolicy `json:"videooptimizationpacingpolicy"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Policies) == 0 {
		return nil, fmt.Errorf("videooptimizationpacingpolicy not found")
	}

	return &result.Policies[0], nil
}

func (s *VideoOptimizationService) CountVideoOptimizationPacingPolicy() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, videoOptimizationPacingPolicyURL+"?count=yes", nil)
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
		} `json:"videooptimizationpacingpolicy"`
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

// videooptimizationpacingpolicylabel
func (s *VideoOptimizationService) AddVideoOptimizationPacingPolicyLabel(resource models.VideoOptimizationPacingPolicyLabel) error {
	payload := map[string]any{
		"videooptimizationpacingpolicylabel": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, videoOptimizationPacingPolicyLabelURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VideoOptimizationService) DeleteVideoOptimizationPacingPolicyLabel(name string) error {
	reqURL := fmt.Sprintf("%s/%s", videoOptimizationPacingPolicyLabelURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VideoOptimizationService) RenameVideoOptimizationPacingPolicyLabel(resource models.VideoOptimizationPacingPolicyLabel) error {
	payload := map[string]any{
		"videooptimizationpacingpolicylabel": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, videoOptimizationPacingPolicyLabelURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VideoOptimizationService) GetAllVideoOptimizationPacingPolicyLabel() ([]models.VideoOptimizationPacingPolicyLabel, error) {
	req, err := s.client.NewRequest(http.MethodGet, videoOptimizationPacingPolicyLabelURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Labels []models.VideoOptimizationPacingPolicyLabel `json:"videooptimizationpacingpolicylabel"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Labels, nil
}

func (s *VideoOptimizationService) GetVideoOptimizationPacingPolicyLabel(name string) (*models.VideoOptimizationPacingPolicyLabel, error) {
	reqURL := fmt.Sprintf("%s/%s", videoOptimizationPacingPolicyLabelURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Labels []models.VideoOptimizationPacingPolicyLabel `json:"videooptimizationpacingpolicylabel"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Labels) == 0 {
		return nil, fmt.Errorf("videooptimizationpacingpolicylabel not found")
	}

	return &result.Labels[0], nil
}

func (s *VideoOptimizationService) CountVideoOptimizationPacingPolicyLabel() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, videoOptimizationPacingPolicyLabelURL+"?count=yes", nil)
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
		} `json:"videooptimizationpacingpolicylabel"`
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

// videooptimizationpacingpolicylabel_binding
func (s *VideoOptimizationService) GetAllVideoOptimizationPacingPolicyLabelBinding() ([]models.VideoOptimizationPacingPolicyLabelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, videoOptimizationPacingPolicyLabelBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.VideoOptimizationPacingPolicyLabelBinding `json:"videooptimizationpacingpolicylabel_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *VideoOptimizationService) GetVideoOptimizationPacingPolicyLabelBinding(name string) (*models.VideoOptimizationPacingPolicyLabelBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", videoOptimizationPacingPolicyLabelBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.VideoOptimizationPacingPolicyLabelBinding `json:"videooptimizationpacingpolicylabel_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) == 0 {
		return nil, fmt.Errorf("videooptimizationpacingpolicylabel_binding not found")
	}

	return &result.Bindings[0], nil
}

// videooptimizationpacingpolicylabel_policybinding_binding
func (s *VideoOptimizationService) GetAllVideoOptimizationPacingPolicyLabelPolicyBindingBinding() ([]models.VideoOptimizationPacingPolicyLabelPolicyBindingBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, videoOptimizationPacingPolicyLabelPolicyBindingBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.VideoOptimizationPacingPolicyLabelPolicyBindingBinding `json:"videooptimizationpacingpolicylabel_policybinding_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *VideoOptimizationService) GetVideoOptimizationPacingPolicyLabelPolicyBindingBinding(name string) ([]models.VideoOptimizationPacingPolicyLabelPolicyBindingBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", videoOptimizationPacingPolicyLabelPolicyBindingBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.VideoOptimizationPacingPolicyLabelPolicyBindingBinding `json:"videooptimizationpacingpolicylabel_policybinding_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *VideoOptimizationService) CountVideoOptimizationPacingPolicyLabelPolicyBindingBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", videoOptimizationPacingPolicyLabelPolicyBindingBindingURL, url.PathEscape(name))
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
		} `json:"videooptimizationpacingpolicylabel_policybinding_binding"`
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

// videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding
func (s *VideoOptimizationService) AddVideoOptimizationPacingPolicyLabelVideoOptimizationPacingPolicyBinding(resource models.VideoOptimizationPacingPolicyLabelVideoOptimizationPacingPolicyBinding) error {
	payload := map[string]any{
		"videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, videoOptimizationPacingPolicyLabelVideoOptimizationPacingPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VideoOptimizationService) DeleteVideoOptimizationPacingPolicyLabelVideoOptimizationPacingPolicyBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", videoOptimizationPacingPolicyLabelVideoOptimizationPacingPolicyBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VideoOptimizationService) GetAllVideoOptimizationPacingPolicyLabelVideoOptimizationPacingPolicyBinding() ([]models.VideoOptimizationPacingPolicyLabelVideoOptimizationPacingPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, videoOptimizationPacingPolicyLabelVideoOptimizationPacingPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.VideoOptimizationPacingPolicyLabelVideoOptimizationPacingPolicyBinding `json:"videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *VideoOptimizationService) GetVideoOptimizationPacingPolicyLabelVideoOptimizationPacingPolicyBinding(name string) ([]models.VideoOptimizationPacingPolicyLabelVideoOptimizationPacingPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", videoOptimizationPacingPolicyLabelVideoOptimizationPacingPolicyBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.VideoOptimizationPacingPolicyLabelVideoOptimizationPacingPolicyBinding `json:"videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *VideoOptimizationService) CountVideoOptimizationPacingPolicyLabelVideoOptimizationPacingPolicyBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", videoOptimizationPacingPolicyLabelVideoOptimizationPacingPolicyBindingURL, url.PathEscape(name))
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
		} `json:"videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding"`
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

// videooptimizationpacingpolicy_binding
func (s *VideoOptimizationService) GetAllVideoOptimizationPacingPolicyBinding() ([]models.VideoOptimizationPacingPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, videoOptimizationPacingPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.VideoOptimizationPacingPolicyBinding `json:"videooptimizationpacingpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *VideoOptimizationService) GetVideoOptimizationPacingPolicyBinding(name string) (*models.VideoOptimizationPacingPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", videoOptimizationPacingPolicyBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.VideoOptimizationPacingPolicyBinding `json:"videooptimizationpacingpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) == 0 {
		return nil, fmt.Errorf("videooptimizationpacingpolicy_binding not found")
	}

	return &result.Bindings[0], nil
}

// videooptimizationpacingpolicy_lbvserver_binding
func (s *VideoOptimizationService) GetAllVideoOptimizationPacingPolicyLBVServerBinding() ([]models.VideoOptimizationPacingPolicyLBVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, videoOptimizationPacingPolicyLBVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.VideoOptimizationPacingPolicyLBVServerBinding `json:"videooptimizationpacingpolicy_lbvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *VideoOptimizationService) GetVideoOptimizationPacingPolicyLBVServerBinding(name string) ([]models.VideoOptimizationPacingPolicyLBVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", videoOptimizationPacingPolicyLBVServerBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.VideoOptimizationPacingPolicyLBVServerBinding `json:"videooptimizationpacingpolicy_lbvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *VideoOptimizationService) CountVideoOptimizationPacingPolicyLBVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", videoOptimizationPacingPolicyLBVServerBindingURL, url.PathEscape(name))
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
		} `json:"videooptimizationpacingpolicy_lbvserver_binding"`
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

// videooptimizationpacingpolicy_videooptimizationglobalpacing_binding
func (s *VideoOptimizationService) GetAllVideoOptimizationPacingPolicyVideoOptimizationGlobalPacingBinding() ([]models.VideoOptimizationPacingPolicyVideoOptimizationGlobalPacingBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, videoOptimizationPacingPolicyVideoOptimizationGlobalPacingBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.VideoOptimizationPacingPolicyVideoOptimizationGlobalPacingBinding `json:"videooptimizationpacingpolicy_videooptimizationglobalpacing_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *VideoOptimizationService) GetVideoOptimizationPacingPolicyVideoOptimizationGlobalPacingBinding(name string) ([]models.VideoOptimizationPacingPolicyVideoOptimizationGlobalPacingBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", videoOptimizationPacingPolicyVideoOptimizationGlobalPacingBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.VideoOptimizationPacingPolicyVideoOptimizationGlobalPacingBinding `json:"videooptimizationpacingpolicy_videooptimizationglobalpacing_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *VideoOptimizationService) CountVideoOptimizationPacingPolicyVideoOptimizationGlobalPacingBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", videoOptimizationPacingPolicyVideoOptimizationGlobalPacingBindingURL, url.PathEscape(name))
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
		} `json:"videooptimizationpacingpolicy_videooptimizationglobalpacing_binding"`
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

// videooptimizationparameter
func (s *VideoOptimizationService) UpdateVideoOptimizationParameter(resource models.VideoOptimizationParameter) error {
	payload := map[string]any{
		"videooptimizationparameter": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, videoOptimizationParameterURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VideoOptimizationService) UnsetVideoOptimizationParameter(resource models.VideoOptimizationParameter) error {
	payload := map[string]any{
		"videooptimizationparameter": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, videoOptimizationParameterURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VideoOptimizationService) GetAllVideoOptimizationParameter() (*models.VideoOptimizationParameter, error) {
	req, err := s.client.NewRequest(http.MethodGet, videoOptimizationParameterURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Parameters []models.VideoOptimizationParameter `json:"videooptimizationparameter"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Parameters) == 0 {
		return nil, fmt.Errorf("videooptimizationparameter not found")
	}

	return &result.Parameters[0], nil
}
