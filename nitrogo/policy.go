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
	policyDatasetURL                 = "/nitro/v1/config/policydataset"
	policyDatasetBindingURL          = "/nitro/v1/config/policydataset_binding"
	policyDatasetValueBindingURL     = "/nitro/v1/config/policydataset_value_binding"
	policyEvaluationURL              = "/nitro/v1/config/policyevaluation"
	policyExpressionURL              = "/nitro/v1/config/policyexpression"
	policyHTTPCalloutURL             = "/nitro/v1/config/policyhttpcallout"
	policyMapURL                     = "/nitro/v1/config/policymap"
	policyParamURL                   = "/nitro/v1/config/policyparam"
	policyPATSetURL                  = "/nitro/v1/config/policypatset"
	policyPATSetBindingURL           = "/nitro/v1/config/policypatset_binding"
	policyPATSetPatternBindingURL    = "/nitro/v1/config/policypatset_pattern_binding"
	policyStringMapURL               = "/nitro/v1/config/policystringmap"
	policyStringMapBindingURL        = "/nitro/v1/config/policystringmap_binding"
	policyStringMapPatternBindingURL = "/nitro/v1/config/policystringmap_pattern_binding"
	policyURLSetURL                  = "/nitro/v1/config/policyurlset"
)

// Policy configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/policy/policy
type PolicyService struct {
	client *Client
}

// policydataset

func (s *PolicyService) AddPolicyDataset(dataset models.PolicyDataset) error {
	payload := map[string]any{
		"policydataset": dataset,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, policyDatasetURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *PolicyService) DeletePolicyDataset(name string) error {
	reqURL := fmt.Sprintf("%s/%s", policyDatasetURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *PolicyService) UpdatePolicyDataset(dataset models.PolicyDataset) error {
	payload := map[string]any{
		"policydataset": dataset,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, policyDatasetURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *PolicyService) UnsetPolicyDataset(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{
		"policydataset": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, policyDatasetURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *PolicyService) GetAllPolicyDataset() ([]models.PolicyDataset, error) {
	req, err := s.client.NewRequest(http.MethodGet, policyDatasetURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Datasets []models.PolicyDataset `json:"policydataset"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Datasets, nil
}

func (s *PolicyService) GetPolicyDataset(name string) ([]models.PolicyDataset, error) {
	reqURL := fmt.Sprintf("%s/%s", policyDatasetURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Datasets []models.PolicyDataset `json:"policydataset"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Datasets, nil
}

func (s *PolicyService) CountPolicyDataset() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, policyDatasetURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Datasets []struct {
			Count float64 `json:"__count"`
		} `json:"policydataset"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Datasets) > 0 {
		return result.Datasets[0].Count, nil
	}
	return 0, nil
}

// policydataset_binding

func (s *PolicyService) GetAllPolicyDatasetBinding() ([]models.PolicyDatasetBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, policyDatasetBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.PolicyDatasetBinding `json:"policydataset_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *PolicyService) GetPolicyDatasetBinding(name string) ([]models.PolicyDatasetBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", policyDatasetBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.PolicyDatasetBinding `json:"policydataset_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// policydataset_value_binding

func (s *PolicyService) AddPolicyDatasetValueBinding(binding models.PolicyDatasetValueBinding) error {
	payload := map[string]any{
		"policydataset_value_binding": binding,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, policyDatasetValueBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *PolicyService) DeletePolicyDatasetValueBinding(name, value string) error {
	reqURL := fmt.Sprintf("%s/%s?args=value:%s", policyDatasetValueBindingURL, url.QueryEscape(name), url.QueryEscape(value))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *PolicyService) GetAllPolicyDatasetValueBinding() ([]models.PolicyDatasetValueBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, policyDatasetValueBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.PolicyDatasetValueBinding `json:"policydataset_value_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *PolicyService) GetPolicyDatasetValueBinding(name string) ([]models.PolicyDatasetValueBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", policyDatasetValueBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.PolicyDatasetValueBinding `json:"policydataset_value_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *PolicyService) CountPolicyDatasetValueBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", policyDatasetValueBindingURL, url.QueryEscape(name))
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
		} `json:"policydataset_value_binding"`
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

// policyevaluation

func (s *PolicyService) GetAllPolicyEvaluation() ([]models.PolicyEvaluation, error) {
	req, err := s.client.NewRequest(http.MethodGet, policyEvaluationURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Evals []models.PolicyEvaluation `json:"policyevaluation"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Evals, nil
}

func (s *PolicyService) CountPolicyEvaluation() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, policyEvaluationURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Evals []struct {
			Count float64 `json:"__count"`
		} `json:"policyevaluation"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Evals) > 0 {
		return result.Evals[0].Count, nil
	}
	return 0, nil
}

// policyexpression

func (s *PolicyService) AddPolicyExpression(expression models.PolicyExpression) error {
	payload := map[string]any{
		"policyexpression": expression,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, policyExpressionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *PolicyService) DeletePolicyExpression(name string) error {
	reqURL := fmt.Sprintf("%s/%s", policyExpressionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *PolicyService) UpdatePolicyExpression(expression models.PolicyExpression) error {
	payload := map[string]any{
		"policyexpression": expression,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, policyExpressionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *PolicyService) UnsetPolicyExpression(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{
		"policyexpression": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, policyExpressionURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *PolicyService) GetAllPolicyExpression() ([]models.PolicyExpression, error) {
	req, err := s.client.NewRequest(http.MethodGet, policyExpressionURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Exprs []models.PolicyExpression `json:"policyexpression"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Exprs, nil
}

func (s *PolicyService) GetPolicyExpression(name string) ([]models.PolicyExpression, error) {
	reqURL := fmt.Sprintf("%s/%s", policyExpressionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Exprs []models.PolicyExpression `json:"policyexpression"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Exprs, nil
}

func (s *PolicyService) CountPolicyExpression() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, policyExpressionURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Exprs []struct {
			Count float64 `json:"__count"`
		} `json:"policyexpression"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Exprs) > 0 {
		return result.Exprs[0].Count, nil
	}
	return 0, nil
}

// policyhttpcallout

func (s *PolicyService) AddPolicyHTTPCallout(callout models.PolicyHTTPCallout) error {
	payload := map[string]any{
		"policyhttpcallout": callout,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, policyHTTPCalloutURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *PolicyService) DeletePolicyHTTPCallout(name string) error {
	reqURL := fmt.Sprintf("%s/%s", policyHTTPCalloutURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *PolicyService) UpdatePolicyHTTPCallout(callout models.PolicyHTTPCallout) error {
	payload := map[string]any{
		"policyhttpcallout": callout,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, policyHTTPCalloutURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *PolicyService) UnsetPolicyHTTPCallout(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{
		"policyhttpcallout": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, policyHTTPCalloutURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *PolicyService) GetAllPolicyHTTPCallout() ([]models.PolicyHTTPCallout, error) {
	req, err := s.client.NewRequest(http.MethodGet, policyHTTPCalloutURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Callouts []models.PolicyHTTPCallout `json:"policyhttpcallout"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Callouts, nil
}

func (s *PolicyService) GetPolicyHTTPCallout(name string) ([]models.PolicyHTTPCallout, error) {
	reqURL := fmt.Sprintf("%s/%s", policyHTTPCalloutURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Callouts []models.PolicyHTTPCallout `json:"policyhttpcallout"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Callouts, nil
}

func (s *PolicyService) CountPolicyHTTPCallout() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, policyHTTPCalloutURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Callouts []struct {
			Count float64 `json:"__count"`
		} `json:"policyhttpcallout"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Callouts) > 0 {
		return result.Callouts[0].Count, nil
	}
	return 0, nil
}

// policymap

func (s *PolicyService) AddPolicyMap(policymap models.PolicyMap) error {
	payload := map[string]any{
		"policymap": policymap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, policyMapURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *PolicyService) DeletePolicyMap(mappolicyname string) error {
	reqURL := fmt.Sprintf("%s/%s", policyMapURL, url.QueryEscape(mappolicyname))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *PolicyService) GetAllPolicyMap() ([]models.PolicyMap, error) {
	req, err := s.client.NewRequest(http.MethodGet, policyMapURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Maps []models.PolicyMap `json:"policymap"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Maps, nil
}

func (s *PolicyService) GetPolicyMap(mappolicyname string) ([]models.PolicyMap, error) {
	reqURL := fmt.Sprintf("%s/%s", policyMapURL, url.QueryEscape(mappolicyname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Maps []models.PolicyMap `json:"policymap"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Maps, nil
}

func (s *PolicyService) CountPolicyMap() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, policyMapURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Maps []struct {
			Count float64 `json:"__count"`
		} `json:"policymap"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Maps) > 0 {
		return result.Maps[0].Count, nil
	}
	return 0, nil
}

// policyparam

func (s *PolicyService) UpdatePolicyParam(param models.PolicyParam) error {
	payload := map[string]any{
		"policyparam": param,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, policyParamURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *PolicyService) UnsetPolicyParam(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{
		"policyparam": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, policyParamURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *PolicyService) GetAllPolicyParam() (models.PolicyParam, error) {
	req, err := s.client.NewRequest(http.MethodGet, policyParamURL, nil)
	if err != nil {
		return models.PolicyParam{}, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return models.PolicyParam{}, err
	}
	var result struct {
		Params []models.PolicyParam `json:"policyparam"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.PolicyParam{}, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Params) > 0 {
		return result.Params[0], nil
	}
	return models.PolicyParam{}, nil
}

// policypatset

func (s *PolicyService) AddPolicyPATSet(patset models.PolicyPatset) error {
	payload := map[string]any{
		"policypatset": patset,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, policyPATSetURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *PolicyService) DeletePolicyPATSet(name string) error {
	reqURL := fmt.Sprintf("%s/%s", policyPATSetURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *PolicyService) UpdatePolicyPATSet(patset models.PolicyPatset) error {
	payload := map[string]any{
		"policypatset": patset,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, policyPATSetURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *PolicyService) UnsetPolicyPATSet(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{
		"policypatset": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, policyPATSetURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *PolicyService) GetAllPolicyPATSet() ([]models.PolicyPatset, error) {
	req, err := s.client.NewRequest(http.MethodGet, policyPATSetURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Pats []models.PolicyPatset `json:"policypatset"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Pats, nil
}

func (s *PolicyService) GetPolicyPATSet(name string) ([]models.PolicyPatset, error) {
	reqURL := fmt.Sprintf("%s/%s", policyPATSetURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Pats []models.PolicyPatset `json:"policypatset"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Pats, nil
}

func (s *PolicyService) CountPolicyPATSet() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, policyPATSetURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Pats []struct {
			Count float64 `json:"__count"`
		} `json:"policypatset"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Pats) > 0 {
		return result.Pats[0].Count, nil
	}
	return 0, nil
}

// policypatset_binding

func (s *PolicyService) GetAllPolicyPATSetBinding() ([]models.PolicyPatsetBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, policyPATSetBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.PolicyPatsetBinding `json:"policypatset_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *PolicyService) GetPolicyPATSetBinding(name string) ([]models.PolicyPatsetBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", policyPATSetBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.PolicyPatsetBinding `json:"policypatset_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// policypatset_pattern_binding

func (s *PolicyService) AddPolicyPATSetPatternBinding(binding models.PolicyPatsetPatternBinding) error {
	payload := map[string]any{
		"policypatset_pattern_binding": binding,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, policyPATSetPatternBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *PolicyService) DeletePolicyPATSetPatternBinding(name, pattern string) error {
	reqURL := fmt.Sprintf("%s/%s?args=String:%s", policyPATSetPatternBindingURL, url.QueryEscape(name), url.QueryEscape(pattern))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *PolicyService) GetAllPolicyPATSetPatternBinding() ([]models.PolicyPatsetPatternBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, policyPATSetPatternBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.PolicyPatsetPatternBinding `json:"policypatset_pattern_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *PolicyService) GetPolicyPATSetPatternBinding(name string) ([]models.PolicyPatsetPatternBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", policyPATSetPatternBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.PolicyPatsetPatternBinding `json:"policypatset_pattern_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *PolicyService) CountPolicyPATSetPatternBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", policyPATSetPatternBindingURL, url.QueryEscape(name))
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
		} `json:"policypatset_pattern_binding"`
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

// policystringmap

func (s *PolicyService) AddPolicyStringMap(stringmap models.PolicyStringMap) error {
	payload := map[string]any{
		"policystringmap": stringmap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, policyStringMapURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *PolicyService) DeletePolicyStringMap(name string) error {
	reqURL := fmt.Sprintf("%s/%s", policyStringMapURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *PolicyService) UpdatePolicyStringMap(stringmap models.PolicyStringMap) error {
	payload := map[string]any{
		"policystringmap": stringmap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, policyStringMapURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *PolicyService) UnsetPolicyStringMap(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{
		"policystringmap": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, policyStringMapURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *PolicyService) GetAllPolicyStringMap() ([]models.PolicyStringMap, error) {
	req, err := s.client.NewRequest(http.MethodGet, policyStringMapURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Maps []models.PolicyStringMap `json:"policystringmap"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Maps, nil
}

func (s *PolicyService) GetPolicyStringMap(name string) ([]models.PolicyStringMap, error) {
	reqURL := fmt.Sprintf("%s/%s", policyStringMapURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Maps []models.PolicyStringMap `json:"policystringmap"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Maps, nil
}

func (s *PolicyService) CountPolicyStringMap() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, policyStringMapURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Maps []struct {
			Count float64 `json:"__count"`
		} `json:"policystringmap"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Maps) > 0 {
		return result.Maps[0].Count, nil
	}
	return 0, nil
}

// policystringmap_binding

func (s *PolicyService) GetAllPolicyStringMapBinding() ([]models.PolicyStringMapBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, policyStringMapBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.PolicyStringMapBinding `json:"policystringmap_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *PolicyService) GetPolicyStringMapBinding(name string) ([]models.PolicyStringMapBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", policyStringMapBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.PolicyStringMapBinding `json:"policystringmap_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// policystringmap_pattern_binding

func (s *PolicyService) AddPolicyStringMapPatternBinding(binding models.PolicyStringMapPatternBinding) error {
	payload := map[string]any{
		"policystringmap_pattern_binding": binding,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, policyStringMapPatternBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *PolicyService) DeletePolicyStringMapPatternBinding(name, key string) error {
	reqURL := fmt.Sprintf("%s/%s?args=key:%s", policyStringMapPatternBindingURL, url.QueryEscape(name), url.QueryEscape(key))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *PolicyService) GetAllPolicyStringMapPatternBinding() ([]models.PolicyStringMapPatternBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, policyStringMapPatternBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.PolicyStringMapPatternBinding `json:"policystringmap_pattern_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *PolicyService) GetPolicyStringMapPatternBinding(name string) ([]models.PolicyStringMapPatternBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", policyStringMapPatternBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.PolicyStringMapPatternBinding `json:"policystringmap_pattern_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *PolicyService) CountPolicyStringMapPatternBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", policyStringMapPatternBindingURL, url.QueryEscape(name))
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
		} `json:"policystringmap_pattern_binding"`
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

// policyurlset

func (s *PolicyService) AddPolicyURLSet(urlset models.PolicyURLSet) error {
	payload := map[string]any{
		"policyurlset": urlset,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, policyURLSetURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *PolicyService) DeletePolicyURLSet(name string) error {
	reqURL := fmt.Sprintf("%s/%s", policyURLSetURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *PolicyService) GetAllPolicyURLSet() ([]models.PolicyURLSet, error) {
	req, err := s.client.NewRequest(http.MethodGet, policyURLSetURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Sets []models.PolicyURLSet `json:"policyurlset"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Sets, nil
}

func (s *PolicyService) GetPolicyURLSet(name string) ([]models.PolicyURLSet, error) {
	reqURL := fmt.Sprintf("%s/%s", policyURLSetURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Sets []models.PolicyURLSet `json:"policyurlset"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Sets, nil
}

func (s *PolicyService) CountPolicyURLSet() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, policyURLSetURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Sets []struct {
			Count float64 `json:"__count"`
		} `json:"policyurlset"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Sets) > 0 {
		return result.Sets[0].Count, nil
	}
	return 0, nil
}

func (s *PolicyService) ImportPolicyURLSet(urlset models.PolicyURLSet) error {
	payload := map[string]any{
		"policyurlset": urlset,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, policyURLSetURL+"?action=Import", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *PolicyService) ExportPolicyURLSet(urlset models.PolicyURLSet) error {
	payload := map[string]any{
		"policyurlset": urlset,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, policyURLSetURL+"?action=export", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *PolicyService) ChangePolicyURLSet(urlset models.PolicyURLSet) error {
	payload := map[string]any{
		"policyurlset": urlset,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, policyURLSetURL+"?action=change", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}
