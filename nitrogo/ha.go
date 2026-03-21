package nitrogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AdamJCrawford/NitroGo/nitrogo/models"
)

const (
	haFilesURL                               = "/nitro/v1/config/hafiles"
	haNodeURL                                = "/nitro/v1/config/hanode"
	haNodeBindingURL                         = "/nitro/v1/config/hanode_binding"
	haNodeCIBindingURL                       = "/nitro/v1/config/hanode_ci_binding"
	haNodeFISBindingURL                      = "/nitro/v1/config/hanode_fis_binding"
	haNodePartialFailureInterfacesBindingURL = "/nitro/v1/config/hanode_partialfailureinterfaces_binding"
	haNodeRouteMonitor6BindingURL            = "/nitro/v1/config/hanode_routemonitor6_binding"
	haNodeRouteMonitorBindingURL             = "/nitro/v1/config/hanode_routemonitor_binding"
	haSyncURL                                = "/nitro/v1/config/hasync"
	haSyncFailuresURL                        = "/nitro/v1/config/hasyncfailures"
)

// Configuration for failover resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/ha/hafailover
type HAService struct {
	client *Client
}

// hafiles

func (s *HAService) SyncHAFiles(files models.HAFiles) error {
	payload := map[string]any{
		"hafiles": files,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, haFilesURL+"?action=sync", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// hanode

func (s *HAService) AddHANode(node models.HANode) error {
	payload := map[string]any{
		"hanode": node,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, haNodeURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *HAService) DeleteHANode(id int) error {
	url := fmt.Sprintf("%s/%d", haNodeURL, id)
	req, err := s.client.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *HAService) UpdateHANode(node models.HANode) error {
	payload := map[string]any{
		"hanode": node,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, haNodeURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *HAService) UnsetHANode(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"hanode": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, haNodeURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *HAService) GetAllHANode() ([]models.HANode, error) {
	req, err := s.client.NewRequest(http.MethodGet, haNodeURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		HANodes []models.HANode `json:"hanode"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.HANodes, nil
}

func (s *HAService) GetHANode(id int) ([]models.HANode, error) {
	url := fmt.Sprintf("%s/%d", haNodeURL, id)
	req, err := s.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		HANodes []models.HANode `json:"hanode"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.HANodes, nil
}

func (s *HAService) CountHANode() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, haNodeURL+"?view=count", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		HANodes []struct {
			Count float64 `json:"__count"`
		} `json:"hanode"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.HANodes) > 0 {
		return result.HANodes[0].Count, nil
	}

	return 0, nil
}

// hanode_binding

func (s *HAService) GetAllHANodeBinding() ([]models.HANodeBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, haNodeBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.HANodeBinding `json:"hanode_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *HAService) GetHANodeBinding(id int) ([]models.HANodeBinding, error) {
	url := fmt.Sprintf("%s/%d", haNodeBindingURL, id)
	req, err := s.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.HANodeBinding `json:"hanode_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

// hanode_ci_binding

func (s *HAService) GetAllHANodeCIBinding() ([]models.HANodeCIBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, haNodeCIBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.HANodeCIBinding `json:"hanode_ci_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *HAService) GetHANodeCIBinding(id int) ([]models.HANodeCIBinding, error) {
	url := fmt.Sprintf("%s/%d", haNodeCIBindingURL, id)
	req, err := s.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.HANodeCIBinding `json:"hanode_ci_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *HAService) CountHANodeCIBinding(id int) (float64, error) {
	url := fmt.Sprintf("%s/%d?count=yes", haNodeCIBindingURL, id)
	req, err := s.client.NewRequest(http.MethodGet, url, nil)
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
		} `json:"hanode_ci_binding"`
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

// hanode_fis_binding

func (s *HAService) GetAllHANodeFISBinding() ([]models.HANodeFISBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, haNodeFISBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.HANodeFISBinding `json:"hanode_fis_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *HAService) GetHANodeFISBinding(id int) ([]models.HANodeFISBinding, error) {
	url := fmt.Sprintf("%s/%d", haNodeFISBindingURL, id)
	req, err := s.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.HANodeFISBinding `json:"hanode_fis_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *HAService) CountHANodeFISBinding(id int) (float64, error) {
	url := fmt.Sprintf("%s/%d?count=yes", haNodeFISBindingURL, id)
	req, err := s.client.NewRequest(http.MethodGet, url, nil)
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
		} `json:"hanode_fis_binding"`
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

// hanode_partialfailureinterfaces_binding

func (s *HAService) GetAllHANodePartialFailureInterfacesBinding() ([]models.HANodePartialFailureInterfacesBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, haNodePartialFailureInterfacesBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.HANodePartialFailureInterfacesBinding `json:"hanode_partialfailureinterfaces_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *HAService) GetHANodePartialFailureInterfacesBinding(id int) ([]models.HANodePartialFailureInterfacesBinding, error) {
	url := fmt.Sprintf("%s/%d", haNodePartialFailureInterfacesBindingURL, id)
	req, err := s.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.HANodePartialFailureInterfacesBinding `json:"hanode_partialfailureinterfaces_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *HAService) CountHANodePartialFailureInterfacesBinding(id int) (float64, error) {
	url := fmt.Sprintf("%s/%d?count=yes", haNodePartialFailureInterfacesBindingURL, id)
	req, err := s.client.NewRequest(http.MethodGet, url, nil)
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
		} `json:"hanode_partialfailureinterfaces_binding"`
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

// hanode_routemonitor6_binding

func (s *HAService) AddHANodeRouteMonitor6Binding(binding models.HANodeRouteMonitor6Binding) error {
	payload := map[string]any{
		"hanode_routemonitor6_binding": binding,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, haNodeRouteMonitor6BindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *HAService) DeleteHANodeRouteMonitor6Binding(id int, routemonitor string) error {
	urlReq := fmt.Sprintf("%s/%d?args=routemonitor:%s", haNodeRouteMonitor6BindingURL, id, routemonitor)
	req, err := s.client.NewRequest(http.MethodDelete, urlReq, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *HAService) GetAllHANodeRouteMonitor6Binding() ([]models.HANodeRouteMonitor6Binding, error) {
	req, err := s.client.NewRequest(http.MethodGet, haNodeRouteMonitor6BindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.HANodeRouteMonitor6Binding `json:"hanode_routemonitor6_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *HAService) GetHANodeRouteMonitor6Binding(id int) ([]models.HANodeRouteMonitor6Binding, error) {
	url := fmt.Sprintf("%s/%d", haNodeRouteMonitor6BindingURL, id)
	req, err := s.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.HANodeRouteMonitor6Binding `json:"hanode_routemonitor6_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *HAService) CountHANodeRouteMonitor6Binding(id int) (float64, error) {
	url := fmt.Sprintf("%s/%d?count=yes", haNodeRouteMonitor6BindingURL, id)
	req, err := s.client.NewRequest(http.MethodGet, url, nil)
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
		} `json:"hanode_routemonitor6_binding"`
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

// hanode_routemonitor_binding

func (s *HAService) AddHANodeRouteMonitorBinding(binding models.HANodeRouteMonitorBinding) error {
	payload := map[string]any{
		"hanode_routemonitor_binding": binding,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, haNodeRouteMonitorBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *HAService) DeleteHANodeRouteMonitorBinding(id int, routemonitor string) error {
	urlReq := fmt.Sprintf("%s/%d?args=routemonitor:%s", haNodeRouteMonitorBindingURL, id, routemonitor)
	req, err := s.client.NewRequest(http.MethodDelete, urlReq, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *HAService) GetAllHANodeRouteMonitorBinding() ([]models.HANodeRouteMonitorBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, haNodeRouteMonitorBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.HANodeRouteMonitorBinding `json:"hanode_routemonitor_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *HAService) GetHANodeRouteMonitorBinding(id int) ([]models.HANodeRouteMonitorBinding, error) {
	url := fmt.Sprintf("%s/%d", haNodeRouteMonitorBindingURL, id)
	req, err := s.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.HANodeRouteMonitorBinding `json:"hanode_routemonitor_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *HAService) CountHANodeRouteMonitorBinding(id int) (float64, error) {
	url := fmt.Sprintf("%s/%d?count=yes", haNodeRouteMonitorBindingURL, id)
	req, err := s.client.NewRequest(http.MethodGet, url, nil)
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
		} `json:"hanode_routemonitor_binding"`
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

// hasync

func (s *HAService) ForceHASync(sync models.HASync) error {
	payload := map[string]any{
		"hasync": sync,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, haSyncURL+"?action=Force", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// hasyncfailures

func (s *HAService) GetAllHASyncFailures() ([]models.HASyncFailures, error) {
	req, err := s.client.NewRequest(http.MethodGet, haSyncFailuresURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		HASyncFailures []models.HASyncFailures `json:"hasyncfailures"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.HASyncFailures, nil
}
