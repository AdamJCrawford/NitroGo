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
	clusterFilesURL                                 = "/nitro/v1/config/clusterfiles"
	clusterInstanceURL                              = "/nitro/v1/config/clusterinstance"
	clusterInstanceBindingURL                       = "/nitro/v1/config/clusterinstance_binding"
	clusterInstanceClusterNodeBindingURL            = "/nitro/v1/config/clusterinstance_clusternode_binding"
	clusterNodeURL                                  = "/nitro/v1/config/clusternode"
	clusterNodeBindingURL                           = "/nitro/v1/config/clusternode_binding"
	clusterNodeRouteMonitorBindingURL               = "/nitro/v1/config/clusternode_routemonitor_binding"
	clusterNodeGroupURL                             = "/nitro/v1/config/clusternodegroup"
	clusterNodeGroupAuthenticationVServerBindingURL = "/nitro/v1/config/clusternodegroup_authenticationvserver_binding"
	clusterNodeGroupBindingURL                      = "/nitro/v1/config/clusternodegroup_binding"
	clusterNodeGroupClusterNodeBindingURL           = "/nitro/v1/config/clusternodegroup_clusternode_binding"
	clusterNodeGroupCRVServerBindingURL             = "/nitro/v1/config/clusternodegroup_crvserver_binding"
	clusterNodeGroupCSVServerBindingURL             = "/nitro/v1/config/clusternodegroup_csvserver_binding"
	clusterNodeGroupGSLBSiteBindingURL              = "/nitro/v1/config/clusternodegroup_gslbsite_binding"
	clusterNodeGroupGSLBVServerBindingURL           = "/nitro/v1/config/clusternodegroup_gslbvserver_binding"
	clusterNodeGroupLBVServerBindingURL             = "/nitro/v1/config/clusternodegroup_lbvserver_binding"
	clusterNodeGroupNSLimitIdentifierBindingURL     = "/nitro/v1/config/clusternodegroup_nslimitidentifier_binding"
	clusterNodeGroupServiceBindingURL               = "/nitro/v1/config/clusternodegroup_service_binding"
	clusterNodeGroupStreamIdentifierBindingURL      = "/nitro/v1/config/clusternodegroup_streamidentifier_binding"
	clusterNodeGroupVPNVServerBindingURL            = "/nitro/v1/config/clusternodegroup_vpnvserver_binding"
	clusterPropStatusURL                            = "/nitro/v1/config/clusterpropstatus"
	clusterSyncURL                                  = "/nitro/v1/config/clustersync"
)

// Configuration for cluster resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/cluster/cluster
type ClusterService struct {
	client *Client
}

// clusterfiles
func (s *ClusterService) SyncClusterFiles(files models.ClusterFiles) error {
	payload := map[string]any{"clusterfiles": files}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, clusterFilesURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

// clusterinstance
func (s *ClusterService) AddClusterInstance(instance models.ClusterInstance) error {
	payload := map[string]any{"clusterinstance": instance}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, clusterInstanceURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ClusterService) DeleteClusterInstance(clid int) error {
	reqURL := fmt.Sprintf("%s/%d", clusterInstanceURL, clid)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ClusterService) UpdateClusterInstance(instance models.ClusterInstance) error {
	payload := map[string]any{"clusterinstance": instance}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, clusterInstanceURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ClusterService) UnsetClusterInstance(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"clusterinstance": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, clusterInstanceURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ClusterService) EnableClusterInstance(clid int) error {
	payload := map[string]any{
		"clusterinstance": map[string]any{"clid": clid},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, clusterInstanceURL+"?action=enable", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ClusterService) DisableClusterInstance(clid int) error {
	payload := map[string]any{
		"clusterinstance": map[string]any{"clid": clid},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, clusterInstanceURL+"?action=disable", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ClusterService) GetAllClusterInstance() ([]models.ClusterInstance, error) {
	req, err := s.client.NewRequest(http.MethodGet, clusterInstanceURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Instances []models.ClusterInstance `json:"clusterinstance"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Instances, nil
}

func (s *ClusterService) GetClusterInstance(clid string) ([]models.ClusterInstance, error) {
	reqURL := fmt.Sprintf("%s/%s", clusterInstanceURL, url.QueryEscape(clid))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Instances []models.ClusterInstance `json:"clusterinstance"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Instances, nil
}

func (s *ClusterService) CountClusterInstance() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, clusterInstanceURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Instances []struct {
			Count float64 `json:"__count"`
		} `json:"clusterinstance"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Instances) > 0 {
		return result.Instances[0].Count, nil
	}
	return 0, nil
}

// clusterinstance_binding
func (s *ClusterService) GetAllClusterInstanceBinding() ([]models.ClusterInstanceBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, clusterInstanceBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ClusterInstanceBinding `json:"clusterinstance_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ClusterService) GetClusterInstanceBinding(clid string) ([]models.ClusterInstanceBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", clusterInstanceBindingURL, url.QueryEscape(clid))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ClusterInstanceBinding `json:"clusterinstance_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// clusterinstance_clusternode_binding
func (s *ClusterService) GetAllClusterInstanceClusterNodeBinding() ([]models.ClusterInstanceClusterNodeBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, clusterInstanceClusterNodeBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ClusterInstanceClusterNodeBinding `json:"clusterinstance_clusternode_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ClusterService) GetClusterInstanceClusterNodeBinding(clid string) ([]models.ClusterInstanceClusterNodeBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", clusterInstanceClusterNodeBindingURL, url.QueryEscape(clid))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ClusterInstanceClusterNodeBinding `json:"clusterinstance_clusternode_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ClusterService) CountClusterInstanceClusterNodeBinding(clid string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", clusterInstanceClusterNodeBindingURL, url.QueryEscape(clid))
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
		} `json:"clusterinstance_clusternode_binding"`
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

// clusternode
func (s *ClusterService) AddClusterNode(node models.ClusterNode) error {
	payload := map[string]any{"clusternode": node}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, clusterNodeURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ClusterService) DeleteClusterNode(nodeID int) error {
	reqURL := fmt.Sprintf("%s/%d", clusterNodeURL, nodeID)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ClusterService) UpdateClusterNode(node models.ClusterNode) error {
	payload := map[string]any{"clusternode": node}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, clusterNodeURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ClusterService) UnsetClusterNode(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"clusternode": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, clusterNodeURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ClusterService) GetAllClusterNode() ([]models.ClusterNode, error) {
	req, err := s.client.NewRequest(http.MethodGet, clusterNodeURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Nodes []models.ClusterNode `json:"clusternode"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Nodes, nil
}

func (s *ClusterService) GetClusterNode(nodeID string) ([]models.ClusterNode, error) {
	reqURL := fmt.Sprintf("%s/%s", clusterNodeURL, url.QueryEscape(nodeID))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Nodes []models.ClusterNode `json:"clusternode"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Nodes, nil
}

func (s *ClusterService) CountClusterNode() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, clusterNodeURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Nodes []struct {
			Count float64 `json:"__count"`
		} `json:"clusternode"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Nodes) > 0 {
		return result.Nodes[0].Count, nil
	}
	return 0, nil
}

// clusternodegroup
func (s *ClusterService) AddClusterNodeGroup(group models.ClusterNodeGroup) error {
	payload := map[string]any{"clusternodegroup": group}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, clusterNodeGroupURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ClusterService) DeleteClusterNodeGroup(name string) error {
	reqURL := fmt.Sprintf("%s/%s", clusterNodeGroupURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ClusterService) UpdateClusterNodeGroup(group models.ClusterNodeGroup) error {
	payload := map[string]any{"clusternodegroup": group}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, clusterNodeGroupURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ClusterService) UnsetClusterNodeGroup(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"clusternodegroup": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, clusterNodeGroupURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ClusterService) GetAllClusterNodeGroup() ([]models.ClusterNodeGroup, error) {
	req, err := s.client.NewRequest(http.MethodGet, clusterNodeGroupURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Groups []models.ClusterNodeGroup `json:"clusternodegroup"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Groups, nil
}

func (s *ClusterService) GetClusterNodeGroup(name string) ([]models.ClusterNodeGroup, error) {
	reqURL := fmt.Sprintf("%s/%s", clusterNodeGroupURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Groups []models.ClusterNodeGroup `json:"clusternodegroup"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Groups, nil
}

func (s *ClusterService) CountClusterNodeGroup() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, clusterNodeGroupURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Groups []struct {
			Count float64 `json:"__count"`
		} `json:"clusternodegroup"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Groups) > 0 {
		return result.Groups[0].Count, nil
	}
	return 0, nil
}

// clusternodegroup_authenticationvserver_binding
func (s *ClusterService) AddClusternodeGroupAuthenticationVServerBinding(binding models.ClusterNodeGroupAuthenticationVServerBinding) error {
	payload := map[string]any{"clusternodegroup_authenticationvserver_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, clusterNodeGroupAuthenticationVServerBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ClusterService) DeleteClusternodeGroupAuthenticationVServerBinding(name string, vserver string) error {
	reqURL := fmt.Sprintf("%s/%s?args=vserver:%s", clusterNodeGroupAuthenticationVServerBindingURL, url.QueryEscape(name), url.QueryEscape(vserver))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ClusterService) GetAllClusternodeGroupAuthenticationVServerBinding() ([]models.ClusterNodeGroupAuthenticationVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, clusterNodeGroupAuthenticationVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ClusterNodeGroupAuthenticationVServerBinding `json:"clusternodegroup_authenticationvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ClusterService) GetClusternodeGroupAuthenticationVServerBinding(name string) ([]models.ClusterNodeGroupAuthenticationVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", clusterNodeGroupAuthenticationVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ClusterNodeGroupAuthenticationVServerBinding `json:"clusternodegroup_authenticationvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ClusterService) CountClusternodeGroupAuthenticationVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", clusterNodeGroupAuthenticationVServerBindingURL, url.QueryEscape(name))
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
		} `json:"clusternodegroup_authenticationvserver_binding"`
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

// clusternodegroup_binding
func (s *ClusterService) GetAllClusterNodeGroupBinding() ([]models.ClusterNodeGroupBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, clusterNodeGroupBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ClusterNodeGroupBinding `json:"clusternodegroup_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ClusterService) GetClusterNodeGroupBinding(name string) ([]models.ClusterNodeGroupBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", clusterNodeGroupBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ClusterNodeGroupBinding `json:"clusternodegroup_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// clusternodegroup_clusternode_binding
func (s *ClusterService) AddClusterNodeGroupClusterNodeBinding(binding models.ClusterNodeGroupClusterNodeBinding) error {
	payload := map[string]any{"clusternodegroup_clusternode_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, clusterNodeGroupClusterNodeBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ClusterService) DeleteClusterNodeGroupClusterNodeBinding(name string, node int) error {
	reqURL := fmt.Sprintf("%s/%s?args=node:%d", clusterNodeGroupClusterNodeBindingURL, url.QueryEscape(name), node)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ClusterService) GetAllClusterNodeGroupClusterNodeBinding() ([]models.ClusterNodeGroupClusterNodeBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, clusterNodeGroupClusterNodeBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ClusterNodeGroupClusterNodeBinding `json:"clusternodegroup_clusternode_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ClusterService) GetClusterNodeGroupClusterNodeBinding(name string) ([]models.ClusterNodeGroupClusterNodeBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", clusterNodeGroupClusterNodeBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ClusterNodeGroupClusterNodeBinding `json:"clusternodegroup_clusternode_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ClusterService) CountClusterNodeGroupClusterNodeBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", clusterNodeGroupClusterNodeBindingURL, url.QueryEscape(name))
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
		} `json:"clusternodegroup_clusternode_binding"`
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

// clusternodegroup_crvserver_binding
func (s *ClusterService) AddClusterNodeGroupCRVServerBinding(binding models.ClusterNodeGroupCRVServerBinding) error {
	payload := map[string]any{"clusternodegroup_crvserver_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, clusterNodeGroupCRVServerBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ClusterService) DeleteClusterNodeGroupCRVServerBinding(name string, vserver string) error {
	reqURL := fmt.Sprintf("%s/%s?args=vserver:%s", clusterNodeGroupCRVServerBindingURL, url.QueryEscape(name), url.QueryEscape(vserver))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ClusterService) GetAllClusterNodeGroupCRVServerBinding() ([]models.ClusterNodeGroupCRVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, clusterNodeGroupCRVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ClusterNodeGroupCRVServerBinding `json:"clusternodegroup_crvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ClusterService) GetClusterNodeGroupCRVServerBinding(name string) ([]models.ClusterNodeGroupCRVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", clusterNodeGroupCRVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ClusterNodeGroupCRVServerBinding `json:"clusternodegroup_crvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ClusterService) CountClusterNodeGroupCRVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", clusterNodeGroupCRVServerBindingURL, url.QueryEscape(name))
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
		} `json:"clusternodegroup_crvserver_binding"`
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

// clusternodegroup_csvserver_binding
func (s *ClusterService) AddClusterNodeGroupCSVServerBinding(binding models.ClusterNodeGroupCSVServerBinding) error {
	payload := map[string]any{"clusternodegroup_csvserver_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, clusterNodeGroupCSVServerBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ClusterService) DeleteClusterNodeGroupCSVServerBinding(name string, vserver string) error {
	reqURL := fmt.Sprintf("%s/%s?args=vserver:%s", clusterNodeGroupCSVServerBindingURL, url.QueryEscape(name), url.QueryEscape(vserver))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ClusterService) GetAllClusterNodeGroupCSVServerBinding() ([]models.ClusterNodeGroupCSVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, clusterNodeGroupCSVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ClusterNodeGroupCSVServerBinding `json:"clusternodegroup_csvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ClusterService) GetClusterNodeGroupCSVServerBinding(name string) ([]models.ClusterNodeGroupCSVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", clusterNodeGroupCSVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ClusterNodeGroupCSVServerBinding `json:"clusternodegroup_csvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ClusterService) CountClusterNodeGroupCSVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", clusterNodeGroupCSVServerBindingURL, url.QueryEscape(name))
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
		} `json:"clusternodegroup_csvserver_binding"`
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

// clusternodegroup_gslbsite_binding
func (s *ClusterService) AddClusterNodeGroupGSLBSiteBinding(binding models.ClusterNodeGroupGSLBSiteBinding) error {
	payload := map[string]any{"clusternodegroup_gslbsite_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, clusterNodeGroupGSLBSiteBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ClusterService) DeleteClusterNodeGroupGSLBSiteBinding(name string, gslbsite string) error {
	reqURL := fmt.Sprintf("%s/%s?args=gslbsite:%s", clusterNodeGroupGSLBSiteBindingURL, url.QueryEscape(name), url.QueryEscape(gslbsite))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ClusterService) GetAllClusterNodeGroupGSLBSiteBinding() ([]models.ClusterNodeGroupGSLBSiteBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, clusterNodeGroupGSLBSiteBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ClusterNodeGroupGSLBSiteBinding `json:"clusternodegroup_gslbsite_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ClusterService) GetClusterNodeGroupGSLBSiteBinding(name string) ([]models.ClusterNodeGroupGSLBSiteBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", clusterNodeGroupGSLBSiteBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ClusterNodeGroupGSLBSiteBinding `json:"clusternodegroup_gslbsite_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ClusterService) CountClusterNodeGroupGSLBSiteBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", clusterNodeGroupGSLBSiteBindingURL, url.QueryEscape(name))
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
		} `json:"clusternodegroup_gslbsite_binding"`
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

// clusternodegroup_gslbvserver_binding
func (s *ClusterService) AddClusterNodeGroupGSLBVServerBinding(binding models.ClusterNodeGroupGSLBVServerBinding) error {
	payload := map[string]any{"clusternodegroup_gslbvserver_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, clusterNodeGroupGSLBVServerBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ClusterService) DeleteClusterNodeGroupGSLBVServerBinding(name string, vserver string) error {
	reqURL := fmt.Sprintf("%s/%s?args=vserver:%s", clusterNodeGroupGSLBVServerBindingURL, url.QueryEscape(name), url.QueryEscape(vserver))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ClusterService) GetAllClusterNodeGroupGSLBVServerBinding() ([]models.ClusterNodeGroupGSLBVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, clusterNodeGroupGSLBVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ClusterNodeGroupGSLBVServerBinding `json:"clusternodegroup_gslbvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ClusterService) GetClusterNodeGroupGSLBVServerBinding(name string) ([]models.ClusterNodeGroupGSLBVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", clusterNodeGroupGSLBVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ClusterNodeGroupGSLBVServerBinding `json:"clusternodegroup_gslbvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ClusterService) CountClusterNodeGroupGSLBVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", clusterNodeGroupGSLBVServerBindingURL, url.QueryEscape(name))
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
		} `json:"clusternodegroup_gslbvserver_binding"`
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

// clusternodegroup_lbvserver_binding
func (s *ClusterService) AddClusterNodeGroupLBVServerBinding(binding models.ClusterNodeGroupLBVServerBinding) error {
	payload := map[string]any{"clusternodegroup_lbvserver_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, clusterNodeGroupLBVServerBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ClusterService) DeleteClusterNodeGroupLBVServerBinding(name string, vserver string) error {
	reqURL := fmt.Sprintf("%s/%s?args=vserver:%s", clusterNodeGroupLBVServerBindingURL, url.QueryEscape(name), url.QueryEscape(vserver))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ClusterService) GetAllClusterNodeGroupLBVServerBinding() ([]models.ClusterNodeGroupLBVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, clusterNodeGroupLBVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ClusterNodeGroupLBVServerBinding `json:"clusternodegroup_lbvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ClusterService) GetClusterNodeGroupLBVServerBinding(name string) ([]models.ClusterNodeGroupLBVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", clusterNodeGroupLBVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ClusterNodeGroupLBVServerBinding `json:"clusternodegroup_lbvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ClusterService) CountClusterNodeGroupLBVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", clusterNodeGroupLBVServerBindingURL, url.QueryEscape(name))
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
		} `json:"clusternodegroup_lbvserver_binding"`
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

// clusternodegroup_nslimitidentifier_binding
func (s *ClusterService) AddClusterNodeGroupNSLimitIdentifierBinding(binding models.ClusterNodeGroupNSLimitIdentifierBinding) error {
	payload := map[string]any{"clusternodegroup_nslimitidentifier_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, clusterNodeGroupNSLimitIdentifierBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ClusterService) DeleteClusterNodeGroupNSLimitIdentifierBinding(name string, identifierName string) error {
	reqURL := fmt.Sprintf("%s/%s?args=identifiername:%s", clusterNodeGroupNSLimitIdentifierBindingURL, url.QueryEscape(name), url.QueryEscape(identifierName))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ClusterService) GetAllClusterNodeGroupNSLimitIdentifierBinding() ([]models.ClusterNodeGroupNSLimitIdentifierBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, clusterNodeGroupNSLimitIdentifierBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ClusterNodeGroupNSLimitIdentifierBinding `json:"clusternodegroup_nslimitidentifier_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ClusterService) GetClusterNodeGroupNSLimitIdentifierBinding(name string) ([]models.ClusterNodeGroupNSLimitIdentifierBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", clusterNodeGroupNSLimitIdentifierBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ClusterNodeGroupNSLimitIdentifierBinding `json:"clusternodegroup_nslimitidentifier_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ClusterService) CountClusterNodeGroupNSLimitIdentifierBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", clusterNodeGroupNSLimitIdentifierBindingURL, url.QueryEscape(name))
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
		} `json:"clusternodegroup_nslimitidentifier_binding"`
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

// clusternodegroup_service_binding
func (s *ClusterService) AddClusterNodeGroupServiceBinding(binding models.ClusterNodeGroupServiceBinding) error {
	payload := map[string]any{"clusternodegroup_service_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, clusterNodeGroupServiceBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ClusterService) DeleteClusterNodeGroupServiceBinding(name string, service string) error {
	reqURL := fmt.Sprintf("%s/%s?args=service:%s", clusterNodeGroupServiceBindingURL, url.QueryEscape(name), url.QueryEscape(service))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ClusterService) GetAllClusterNodeGroupServiceBinding() ([]models.ClusterNodeGroupServiceBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, clusterNodeGroupServiceBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ClusterNodeGroupServiceBinding `json:"clusternodegroup_service_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ClusterService) GetClusterNodeGroupServiceBinding(name string) ([]models.ClusterNodeGroupServiceBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", clusterNodeGroupServiceBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ClusterNodeGroupServiceBinding `json:"clusternodegroup_service_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ClusterService) CountClusterNodeGroupServiceBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", clusterNodeGroupServiceBindingURL, url.QueryEscape(name))
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
		} `json:"clusternodegroup_service_binding"`
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

// clusternodegroup_streamidentifier_binding
func (s *ClusterService) AddClusterNodeGroupStreamIdentifierBinding(binding models.ClusterNodeGroupStreamIdentifierBinding) error {
	payload := map[string]any{"clusternodegroup_streamidentifier_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, clusterNodeGroupStreamIdentifierBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ClusterService) DeleteClusterNodeGroupStreamIdentifierBinding(name string, identifierName string) error {
	reqURL := fmt.Sprintf("%s/%s?args=identifiername:%s", clusterNodeGroupStreamIdentifierBindingURL, url.QueryEscape(name), url.QueryEscape(identifierName))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ClusterService) GetAllClusterNodeGroupStreamIdentifierBinding() ([]models.ClusterNodeGroupStreamIdentifierBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, clusterNodeGroupStreamIdentifierBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ClusterNodeGroupStreamIdentifierBinding `json:"clusternodegroup_streamidentifier_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ClusterService) GetClusterNodeGroupStreamIdentifierBinding(name string) ([]models.ClusterNodeGroupStreamIdentifierBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", clusterNodeGroupStreamIdentifierBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ClusterNodeGroupStreamIdentifierBinding `json:"clusternodegroup_streamidentifier_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ClusterService) CountClusterNodeGroupStreamIdentifierBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", clusterNodeGroupStreamIdentifierBindingURL, url.QueryEscape(name))
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
		} `json:"clusternodegroup_streamidentifier_binding"`
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

// clusternodegroup_vpnvserver_binding
func (s *ClusterService) AddClusterNodeGroupVPNVServerBinding(binding models.ClusterNodeGroupVPNVServerBinding) error {
	payload := map[string]any{"clusternodegroup_vpnvserver_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, clusterNodeGroupVPNVServerBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ClusterService) DeleteClusterNodeGroupVPNVServerBinding(name string, vserver string) error {
	reqURL := fmt.Sprintf("%s/%s?args=vserver:%s", clusterNodeGroupVPNVServerBindingURL, url.QueryEscape(name), url.QueryEscape(vserver))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ClusterService) GetAllClusterNodeGroupVPNVServerBinding() ([]models.ClusterNodeGroupVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, clusterNodeGroupVPNVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ClusterNodeGroupVPNVServerBinding `json:"clusternodegroup_vpnvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ClusterService) GetClusterNodeGroupVPNVServerBinding(name string) ([]models.ClusterNodeGroupVPNVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", clusterNodeGroupVPNVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ClusterNodeGroupVPNVServerBinding `json:"clusternodegroup_vpnvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ClusterService) CountClusterNodeGroupVPNVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", clusterNodeGroupVPNVServerBindingURL, url.QueryEscape(name))
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
		} `json:"clusternodegroup_vpnvserver_binding"`
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

// clusternode_binding
func (s *ClusterService) GetAllClusterNodeBinding() ([]models.ClusterNodeBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, clusterNodeBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ClusterNodeBinding `json:"clusternode_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ClusterService) GetClusterNodeBinding(nodeID string) ([]models.ClusterNodeBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", clusterNodeBindingURL, url.QueryEscape(nodeID))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ClusterNodeBinding `json:"clusternode_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// clusternode_routemonitor_binding
func (s *ClusterService) AddClusterNodeRouteMonitorBinding(binding models.ClusterNodeRouteMonitorBinding) error {
	payload := map[string]any{"clusternode_routemonitor_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, clusterNodeRouteMonitorBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ClusterService) DeleteClusterNodeRouteMonitorBinding(nodeID int, routeMonitor string) error {
	reqURL := fmt.Sprintf("%s/%d?args=routemonitor:%s", clusterNodeRouteMonitorBindingURL, nodeID, url.QueryEscape(routeMonitor))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *ClusterService) GetAllClusterNodeRouteMonitorBinding() ([]models.ClusterNodeRouteMonitorBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, clusterNodeRouteMonitorBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ClusterNodeRouteMonitorBinding `json:"clusternode_routemonitor_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ClusterService) GetClusterNodeRouteMonitorBinding(nodeID string) ([]models.ClusterNodeRouteMonitorBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", clusterNodeRouteMonitorBindingURL, url.QueryEscape(nodeID))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.ClusterNodeRouteMonitorBinding `json:"clusternode_routemonitor_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *ClusterService) CountClusterNodeRouteMonitorBinding(nodeID string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", clusterNodeRouteMonitorBindingURL, url.QueryEscape(nodeID))
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
		} `json:"clusternode_routemonitor_binding"`
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

// clusterpropstatus
func (s *ClusterService) GetAllClusterPropStatus() ([]models.ClusterPropStatus, error) {
	req, err := s.client.NewRequest(http.MethodGet, clusterPropStatusURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Status []models.ClusterPropStatus `json:"clusterpropstatus"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Status, nil
}

func (s *ClusterService) CountClusterPropStatus() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, clusterPropStatusURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Status []struct {
			Count float64 `json:"__count"`
		} `json:"clusterpropstatus"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Status) > 0 {
		return result.Status[0].Count, nil
	}
	return 0, nil
}

func (s *ClusterService) ClearClusterPropStatus() error {
	payload := map[string]any{}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, clusterPropStatusURL+"?action=clear", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

// clustersync
func (s *ClusterService) ForceClusterSync() error {
	payload := map[string]any{}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, clusterSyncURL+"?action=force", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}
