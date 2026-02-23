package nitrogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/AdamJCrawford/NitroGo/nitrogo/models"
)

const (
	lbActionURL                                         = "/nitro/v1/config/lbaction"
	lbGlobalBindingURL                                  = "/nitro/v1/config/lbglobal_binding"
	lbGlobalLBPolicyBindingURL                          = "/nitro/v1/config/lbglobal_lbpolicy_binding"
	lbGroupURL                                          = "/nitro/v1/config/lbgroup"
	lbGroupBindingURL                                   = "/nitro/v1/config/lbgroup_binding"
	lbGroupLBVServerBindingURL                          = "/nitro/v1/config/lbgroup_lbvserver_binding"
	lbMetricTableURL                                    = "/nitro/v1/config/lbmetrictable"
	lbMetricTableBindingURL                             = "/nitro/v1/config/lbmetrictable_binding"
	lbMetricTableMetricBindingURL                       = "/nitro/v1/config/lbmetrictable_metric_binding"
	lbMonBindingsURL                                    = "/nitro/v1/config/lbmonbindings"
	lbMonBindingsBindingURL                             = "/nitro/v1/config/lbmonbindings_binding"
	lbMonBindingsGSLBServiceGroupBindingURL             = "/nitro/v1/config/lbmonbindings_gslbservicegroup_binding"
	lbMonBindingsServiceBindingURL                      = "/nitro/v1/config/lbmonbindings_service_binding"
	lbMonBindingsServiceGroupBindingURL                 = "/nitro/v1/config/lbmonbindings_servicegroup_binding"
	lbMonitorURL                                        = "/nitro/v1/config/lbmonitor"
	lbMonitorBindingURL                                 = "/nitro/v1/config/lbmonitor_binding"
	lbMonitorMetricBindingURL                           = "/nitro/v1/config/lbmonitor_metric_binding"
	lbMonitorServiceBindingURL                          = "/nitro/v1/config/lbmonitor_service_binding"
	lbMonitorServiceGroupBindingURL                     = "/nitro/v1/config/lbmonitor_servicegroup_binding"
	lbMonitorSSLCertKeyBindingURL                       = "/nitro/v1/config/lbmonitor_sslcertkey_binding"
	lbParameterURL                                      = "/nitro/v1/config/lbparameter"
	lbPersistentSessionsURL                             = "/nitro/v1/config/lbpersistentsessions"
	lbProfileURL                                        = "/nitro/v1/config/lbprofile"
	lbRouteURL                                          = "/nitro/v1/config/lbroute"
	lbRoute6URL                                         = "/nitro/v1/config/lbroute6"
	lbSIPParametersURL                                  = "/nitro/v1/config/lbsipparameters"
	lbVServerURL                                        = "/nitro/v1/config/lbvserver"
	lbVServerAnalyticsProfileBindingURL                 = "/nitro/v1/config/lbvserver_analyticsprofile_binding"
	lbVServerAppFlowPolicyBindingURL                    = "/nitro/v1/config/lbvserver_appflowpolicy_binding"
	lbVServerAppFWPolicyBindingURL                      = "/nitro/v1/config/lbvserver_appfwpolicy_binding"
	lbVServerAppQOEPolicyBindingURL                     = "/nitro/v1/config/lbvserver_appqoepolicy_binding"
	lbVServerAuditNSLogPolicyBindingURL                 = "/nitro/v1/config/lbvserver_auditnslogpolicy_binding"
	lbVServerAuditSyslogPolicyBindingURL                = "/nitro/v1/config/lbvserver_auditsyslogpolicy_binding"
	lbVServerAuthorizationPolicyBindingURL              = "/nitro/v1/config/lbvserver_authorizationpolicy_binding"
	lbVServerBindingURL                                 = "/nitro/v1/config/lbvserver_binding"
	lbVServerBotPolicyBindingURL                        = "/nitro/v1/config/lbvserver_botpolicy_binding"
	lbVServerCachePolicyBindingURL                      = "/nitro/v1/config/lbvserver_cachepolicy_binding"
	lbVServerCMPPolicyBindingURL                        = "/nitro/v1/config/lbvserver_cmppolicy_binding"
	lbVServerContentInspectionPolicyBindingURL          = "/nitro/v1/config/lbvserver_contentinspectionpolicy_binding"
	lbVServerCSVServerBindingURL                        = "/nitro/v1/config/lbvserver_csvserver_binding"
	lbVServerDNSPolicy64BindingURL                      = "/nitro/v1/config/lbvserver_dnspolicy64_binding"
	lbVServerFEOPolicyBindingURL                        = "/nitro/v1/config/lbvserver_feopolicy_binding"
	lbVServerResponderPolicyBindingURL                  = "/nitro/v1/config/lbvserver_responderpolicy_binding"
	lbVServerRewritePolicyBindingURL                    = "/nitro/v1/config/lbvserver_rewritepolicy_binding"
	lbVServerServiceBindingURL                          = "/nitro/v1/config/lbvserver_service_binding"
	lbVServerServiceGroupBindingURL                     = "/nitro/v1/config/lbvserver_servicegroup_binding"
	lbVServerServiceGroupMemberBindingURL               = "/nitro/v1/config/lbvserver_servicegroupmember_binding"
	lbVServerSpilloverPolicyBindingURL                  = "/nitro/v1/config/lbvserver_spilloverpolicy_binding"
	lbVServerTMTrafficPolicyBindingURL                  = "/nitro/v1/config/lbvserver_tmtrafficpolicy_binding"
	lbVServerTransformPolicyBindingURL                  = "/nitro/v1/config/lbvserver_transformpolicy_binding"
	lbVServerVideoOptimizationDetectionPolicyBindingURL = "/nitro/v1/config/lbvserver_videooptimizationdetectionpolicy_binding"
	lbVServerVideoOptimizationPacingPolicyBindingURL    = "/nitro/v1/config/lbvserver_videooptimizationpacingpolicy_binding"
	lbWLMURL                                            = "/nitro/v1/config/lbwlm"
	lbWLMBindingURL                                     = "/nitro/v1/config/lbwlm_binding"
	lbWLMLBVServerBindingURL                            = "/nitro/v1/config/lbwlm_lbvserver_binding"
)

// Load Balancing configuration. The load balancing methods manage the selection of an appropriate physical server in a server farm.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/lb/lb
type LBService struct {
	client *Client
}

// lbaction
// Configuration for lb action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/lb/lbaction
func (s *LBService) AddLBAction(resource map[string]any) error {
	payload := map[string]any{"lbaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBAction(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", lbActionURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) UpdateLBAction(resource map[string]any) error {
	payload := map[string]any{"lbaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, lbActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) UnsetLBAction(resource map[string]any) error {
	payload := map[string]any{"lbaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", lbActionURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetAllLBAction() ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, lbActionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbaction"], nil
}

func (s *LBService) GetLBAction(name string) (map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbActionURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbaction"]) == 0 {
		return nil, fmt.Errorf("lbaction %s not found", name)
	}

	return result["lbaction"][0], nil
}

func (s *LBService) CountLBAction() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", lbActionURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbaction"]) > 0 {
		if count, ok := result["lbaction"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

func (s *LBService) RenameLBAction(name, newName string) error {
	payload := map[string]any{"lbaction": map[string]string{"name": name, "newname": newName}}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=rename", lbActionURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// lbglobal_binding
func (s *LBService) GetLBGlobalBinding() (map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, lbGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if val, ok := result["lbglobal_binding"]; ok {
		return val.(map[string]any), nil
	}
	return nil, fmt.Errorf("lbglobal_binding not found")
}

// lbglobal_lbpolicy_binding
func (s *LBService) AddLBGlobalLBPolicyBinding(resource map[string]any) error {
	payload := map[string]any{"lbglobal_lbpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbGlobalLBPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBGlobalLBPolicyBinding(args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s%s", lbGlobalLBPolicyBindingURL, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetLBGlobalLBPolicyBinding() ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, lbGlobalLBPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbglobal_lbpolicy_binding"], nil
}

func (s *LBService) CountLBGlobalLBPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", lbGlobalLBPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbglobal_lbpolicy_binding"]) > 0 {
		if count, ok := result["lbglobal_lbpolicy_binding"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbgroup
func (s *LBService) AddLBGroup(resource map[string]any) error {
	payload := map[string]any{"lbgroup": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbGroupURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBGroup(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", lbGroupURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) UpdateLBGroup(resource map[string]any) error {
	payload := map[string]any{"lbgroup": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, lbGroupURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) UnsetLBGroup(resource map[string]any) error {
	payload := map[string]any{"lbgroup": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", lbGroupURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetAllLBGroup() ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, lbGroupURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbgroup"], nil
}

func (s *LBService) GetLBGroup(name string) (map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbGroupURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbgroup"]) == 0 {
		return nil, fmt.Errorf("lbgroup %s not found", name)
	}

	return result["lbgroup"][0], nil
}

func (s *LBService) CountLBGroup() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", lbGroupURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbgroup"]) > 0 {
		if count, ok := result["lbgroup"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

func (s *LBService) RenameLBGroup(name, newName string) error {
	payload := map[string]any{"lbgroup": map[string]string{"name": name, "newname": newName}}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=rename", lbGroupURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// lbgroup_binding
func (s *LBService) GetLBGroupBinding(name string) (map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbGroupBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbgroup_binding"]) == 0 {
		return nil, fmt.Errorf("lbgroup_binding %s not found", name)
	}

	return result["lbgroup_binding"][0], nil
}

// lbgroup_lbvserver_binding
func (s *LBService) AddLBGroupLBVServerBinding(resource map[string]any) error {
	payload := map[string]any{"lbgroup_lbvserver_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbGroupLBVServerBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBGroupLBVServerBinding(name string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lbGroupLBVServerBindingURL, name, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetLBGroupLBVServerBinding(name string) ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbGroupLBVServerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbgroup_lbvserver_binding"], nil
}

func (s *LBService) CountLBGroupLBVServerBinding(name string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lbGroupLBVServerBindingURL, name), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbgroup_lbvserver_binding"]) > 0 {
		if count, ok := result["lbgroup_lbvserver_binding"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbmetrictable
func (s *LBService) AddLBMetricTable(resource map[string]any) error {
	payload := map[string]any{"lbmetrictable": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbMetricTableURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBMetricTable(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", lbMetricTableURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) UpdateLBMetricTable(resource map[string]any) error {
	payload := map[string]any{"lbmetrictable": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, lbMetricTableURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetAllLBMetricTable() ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, lbMetricTableURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbmetrictable"], nil
}

func (s *LBService) GetLBMetricTable(name string) (map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbMetricTableURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbmetrictable"]) == 0 {
		return nil, fmt.Errorf("lbmetrictable %s not found", name)
	}

	return result["lbmetrictable"][0], nil
}

func (s *LBService) CountLBMetricTable() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", lbMetricTableURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbmetrictable"]) > 0 {
		if count, ok := result["lbmetrictable"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbmetrictable_binding
func (s *LBService) GetLBMetricTableBinding(name string) (map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbMetricTableBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbmetrictable_binding"]) == 0 {
		return nil, fmt.Errorf("lbmetrictable_binding %s not found", name)
	}

	return result["lbmetrictable_binding"][0], nil
}

// lbmetrictable_metric_binding
func (s *LBService) AddLBMetricTableMetricBinding(resource map[string]any) error {
	payload := map[string]any{"lbmetrictable_metric_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbMetricTableMetricBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBMetricTableMetricBinding(name string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lbMetricTableMetricBindingURL, name, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetLBMetricTableMetricBinding(name string) ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbMetricTableMetricBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbmetrictable_metric_binding"], nil
}

func (s *LBService) CountLBMetricTableMetricBinding(name string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lbMetricTableMetricBindingURL, name), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbmetrictable_metric_binding"]) > 0 {
		if count, ok := result["lbmetrictable_metric_binding"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbmonbindings
func (s *LBService) GetAllLBMonBindings() ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, lbMonBindingsURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbmonbindings"], nil
}

func (s *LBService) GetLBMonBindings(name string) (map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbMonBindingsURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbmonbindings"]) == 0 {
		return nil, fmt.Errorf("lbmonbindings %s not found", name)
	}

	return result["lbmonbindings"][0], nil
}

func (s *LBService) CountLBMonBindings() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", lbMonBindingsURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbmonbindings"]) > 0 {
		if count, ok := result["lbmonbindings"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbmonbindings_binding
func (s *LBService) GetLBMonBindingsBinding(name string) (map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbMonBindingsBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbmonbindings_binding"]) == 0 {
		return nil, fmt.Errorf("lbmonbindings_binding %s not found", name)
	}

	return result["lbmonbindings_binding"][0], nil
}

// lbmonbindings_gslbservicegroup_binding
func (s *LBService) GetLBMonBindingsGSLBServiceGroupBinding(name string) ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbMonBindingsGSLBServiceGroupBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbmonbindings_gslbservicegroup_binding"], nil
}

func (s *LBService) CountLBMonBindingsGSLBServiceGroupBinding(name string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lbMonBindingsGSLBServiceGroupBindingURL, name), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbmonbindings_gslbservicegroup_binding"]) > 0 {
		if count, ok := result["lbmonbindings_gslbservicegroup_binding"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbmonbindings_servicegroup_binding
func (s *LBService) GetLBMonBindingsServiceGroupBinding(name string) ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbMonBindingsServiceGroupBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbmonbindings_servicegroup_binding"], nil
}

func (s *LBService) CountLBMonBindingsServiceGroupBinding(name string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lbMonBindingsServiceGroupBindingURL, name), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbmonbindings_servicegroup_binding"]) > 0 {
		if count, ok := result["lbmonbindings_servicegroup_binding"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbmonbindings_service_binding
func (s *LBService) GetLBMonBindingsServiceBinding(name string) ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbMonBindingsServiceBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbmonbindings_service_binding"], nil
}

func (s *LBService) CountLBMonBindingsServiceBinding(name string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lbMonBindingsServiceBindingURL, name), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbmonbindings_service_binding"]) > 0 {
		if count, ok := result["lbmonbindings_service_binding"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbmonitor
func (s *LBService) AddLBMonitor(resource map[string]any) error {
	payload := map[string]any{"lbmonitor": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbMonitorURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBMonitor(name string, resource map[string]any) error {
	// lbmonitor delete might require type, so we allow resource map for query params or payload if needed.
	// Standard delete is by name.
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", lbMonitorURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) UpdateLBMonitor(resource map[string]any) error {
	payload := map[string]any{"lbmonitor": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, lbMonitorURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) UnsetLBMonitor(resource map[string]any) error {
	payload := map[string]any{"lbmonitor": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", lbMonitorURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) EnableLBMonitor(name string) error {
	payload := map[string]any{"lbmonitor": map[string]string{"monitorname": name}}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=enable", lbMonitorURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DisableLBMonitor(name string) error {
	payload := map[string]any{"lbmonitor": map[string]string{"monitorname": name}}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=disable", lbMonitorURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetAllLBMonitor() ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, lbMonitorURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbmonitor"], nil
}

func (s *LBService) GetLBMonitor(name string) (map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbMonitorURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbmonitor"]) == 0 {
		return nil, fmt.Errorf("lbmonitor %s not found", name)
	}

	return result["lbmonitor"][0], nil
}

func (s *LBService) CountLBMonitor() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", lbMonitorURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbmonitor"]) > 0 {
		if count, ok := result["lbmonitor"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbmonitor_binding
func (s *LBService) GetLBMonitorBinding(name string) (map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbMonitorBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbmonitor_binding"]) == 0 {
		return nil, fmt.Errorf("lbmonitor_binding %s not found", name)
	}

	return result["lbmonitor_binding"][0], nil
}

// lbmonitor_metric_binding
func (s *LBService) AddLBMonitorMetricBinding(resource map[string]any) error {
	payload := map[string]any{"lbmonitor_metric_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbMonitorMetricBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBMonitorMetricBinding(name string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lbMonitorMetricBindingURL, name, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetLBMonitorMetricBinding(name string) ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbMonitorMetricBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbmonitor_metric_binding"], nil
}

func (s *LBService) CountLBMonitorMetricBinding(name string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lbMonitorMetricBindingURL, name), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbmonitor_metric_binding"]) > 0 {
		if count, ok := result["lbmonitor_metric_binding"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbmonitor_servicegroup_binding
func (s *LBService) AddLBMonitorServiceGroupBinding(resource map[string]any) error {
	payload := map[string]any{"lbmonitor_servicegroup_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbMonitorServiceGroupBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBMonitorServiceGroupBinding(name string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lbMonitorServiceGroupBindingURL, name, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// lbmonitor_service_binding
func (s *LBService) AddLBMonitorServiceBinding(resource map[string]any) error {
	payload := map[string]any{"lbmonitor_service_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbMonitorServiceBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBMonitorServiceBinding(name string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lbMonitorServiceBindingURL, name, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// lbmonitor_sslcertkey_binding
func (s *LBService) AddLBMonitorSSLCertKeyBinding(resource map[string]any) error {
	payload := map[string]any{"lbmonitor_sslcertkey_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbMonitorSSLCertKeyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBMonitorSSLCertKeyBinding(name string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lbMonitorSSLCertKeyBindingURL, name, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetLBMonitorSSLCertKeyBinding(name string) ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbMonitorSSLCertKeyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbmonitor_sslcertkey_binding"], nil
}

func (s *LBService) CountLBMonitorSSLCertKeyBinding(name string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lbMonitorSSLCertKeyBindingURL, name), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbmonitor_sslcertkey_binding"]) > 0 {
		if count, ok := result["lbmonitor_sslcertkey_binding"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbparameter
func (s *LBService) UpdateLBParameter(resource map[string]any) error {
	payload := map[string]any{"lbparameter": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, lbParameterURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) UnsetLBParameter(resource map[string]any) error {
	payload := map[string]any{"lbparameter": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", lbParameterURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetAllLBParameter() (map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, lbParameterURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if val, ok := result["lbparameter"]; ok {
		return val.(map[string]any), nil
	}
	return nil, fmt.Errorf("lbparameter not found in response")
}

// lbpersistentsessions
func (s *LBService) GetAllLBPersistentSessions() ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, lbPersistentSessionsURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbpersistentsessions"], nil
}

func (s *LBService) CountLBPersistentSessions() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", lbPersistentSessionsURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbpersistentsessions"]) > 0 {
		if count, ok := result["lbpersistentsessions"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

func (s *LBService) ClearLBPersistentSessions(resource map[string]any) error {
	payload := map[string]any{"lbpersistentsessions": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=clear", lbPersistentSessionsURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// lbprofile
func (s *LBService) AddLBProfile(resource map[string]any) error {
	payload := map[string]any{"lbprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBProfile(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", lbProfileURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) UpdateLBProfile(resource map[string]any) error {
	payload := map[string]any{"lbprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, lbProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) UnsetLBProfile(resource map[string]any) error {
	payload := map[string]any{"lbprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", lbProfileURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetAllLBProfile() ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, lbProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbprofile"], nil
}

func (s *LBService) GetLBProfile(name string) (map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbProfileURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbprofile"]) == 0 {
		return nil, fmt.Errorf("lbprofile %s not found", name)
	}

	return result["lbprofile"][0], nil
}

func (s *LBService) CountLBProfile() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", lbProfileURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbprofile"]) > 0 {
		if count, ok := result["lbprofile"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbroute
func (s *LBService) AddLBRoute(resource map[string]any) error {
	payload := map[string]any{"lbroute": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbRouteURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBRoute(network string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lbRouteURL, network, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetAllLBRoute() ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, lbRouteURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbroute"], nil
}

func (s *LBService) CountLBRoute() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", lbRouteURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbroute"]) > 0 {
		if count, ok := result["lbroute"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbroute6
func (s *LBService) AddLBRoute6(resource map[string]any) error {
	payload := map[string]any{"lbroute6": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbRoute6URL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBRoute6(network string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lbRoute6URL, network, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetAllLBRoute6() ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, lbRoute6URL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbroute6"], nil
}

func (s *LBService) CountLBRoute6() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", lbRoute6URL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbroute6"]) > 0 {
		if count, ok := result["lbroute6"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbsipparameters
func (s *LBService) UpdateLBSIPParameters(resource map[string]any) error {
	payload := map[string]any{"lbsipparameters": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, lbSIPParametersURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) UnsetLBSIPParameters(resource map[string]any) error {
	payload := map[string]any{"lbsipparameters": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", lbSIPParametersURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetAllLBSIPParameters() (map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, lbSIPParametersURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if val, ok := result["lbsipparameters"]; ok {
		return val.(map[string]any), nil
	}
	return nil, fmt.Errorf("lbsipparameters not found in response")
}

// lbvserver
func (s *LBService) AddLBVServer(resource map[string]any) error {
	payload := map[string]any{"lbvserver": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbVServerURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBVServer(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", lbVServerURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) UpdateLBVServer(resource map[string]any) error {
	payload := map[string]any{"lbvserver": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, lbVServerURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) UnsetLBVServer(resource map[string]any) error {
	payload := map[string]any{"lbvserver": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", lbVServerURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) EnableLBVServer(name string) error {
	payload := map[string]any{"lbvserver": map[string]string{"name": name}}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=enable", lbVServerURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DisableLBVServer(name string) error {
	payload := map[string]any{"lbvserver": map[string]string{"name": name}}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=disable", lbVServerURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetAllLBVServer() ([]models.LBVServer, error) {
	req, err := s.client.NewRequest(http.MethodGet, lbVServerURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		LBVServer []models.LBVServer `json:"lbvserver"`
	}

	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.LBVServer, nil
}

func (s *LBService) GetLBVServer(name string) (models.LBVServer, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbVServerURL, name), nil)
	if err != nil {
		return models.LBVServer{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.LBVServer{}, err
	}

	var result struct {
		LBVServer []models.LBVServer `json:"lbvserver"`
	}

	if err := json.Unmarshal(resp, &result); err != nil {
		return models.LBVServer{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.LBVServer) == 0 {
		return models.LBVServer{}, fmt.Errorf("lbvserver %s not found", name)
	}

	return result.LBVServer[0], nil
}

func (s *LBService) CountLBVServer() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", lbVServerURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbvserver"]) > 0 {
		if count, ok := result["lbvserver"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

func (s *LBService) RenameLBVServer(name, newName string) error {
	payload := map[string]any{"lbvserver": map[string]string{"name": name, "newname": newName}}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=rename", lbVServerURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// lbvserver_analyticsprofile_binding
func (s *LBService) AddLBVServerAnalyticsProfileBinding(resource map[string]any) error {
	payload := map[string]any{"lbvserver_analyticsprofile_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbVServerAnalyticsProfileBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBVServerAnalyticsProfileBinding(name string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lbVServerAnalyticsProfileBindingURL, name, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetLBVServerAnalyticsProfileBinding(name string) ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbVServerAnalyticsProfileBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbvserver_analyticsprofile_binding"], nil
}

func (s *LBService) CountLBVServerAnalyticsProfileBinding(name string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lbVServerAnalyticsProfileBindingURL, name), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbvserver_analyticsprofile_binding"]) > 0 {
		if count, ok := result["lbvserver_analyticsprofile_binding"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbvserver_appflowpolicy_binding
func (s *LBService) AddLBVServerAppFlowPolicyBinding(resource map[string]any) error {
	payload := map[string]any{"lbvserver_appflowpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbVServerAppFlowPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBVServerAppFlowPolicyBinding(name string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lbVServerAppFlowPolicyBindingURL, name, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetLBVServerAppFlowPolicyBinding(name string) ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbVServerAppFlowPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbvserver_appflowpolicy_binding"], nil
}

func (s *LBService) CountLBVServerAppFlowPolicyBinding(name string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lbVServerAppFlowPolicyBindingURL, name), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbvserver_appflowpolicy_binding"]) > 0 {
		if count, ok := result["lbvserver_appflowpolicy_binding"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbvserver_appfwpolicy_binding
func (s *LBService) AddLBVServerAppFWPolicyBinding(resource map[string]any) error {
	payload := map[string]any{"lbvserver_appfwpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbVServerAppFWPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBVServerAppFWPolicyBinding(name string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lbVServerAppFWPolicyBindingURL, name, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetLBVServerAppFWPolicyBinding(name string) ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbVServerAppFWPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbvserver_appfwpolicy_binding"], nil
}

func (s *LBService) CountLBVServerAppFWPolicyBinding(name string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lbVServerAppFWPolicyBindingURL, name), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbvserver_appfwpolicy_binding"]) > 0 {
		if count, ok := result["lbvserver_appfwpolicy_binding"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbvserver_appqoepolicy_binding
func (s *LBService) AddLBVServerAppQOEPolicyBinding(resource map[string]any) error {
	payload := map[string]any{"lbvserver_appqoepolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbVServerAppQOEPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBVServerAppQOEPolicyBinding(name string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lbVServerAppQOEPolicyBindingURL, name, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetLBVServerAppQOEPolicyBinding(name string) ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbVServerAppQOEPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbvserver_appqoepolicy_binding"], nil
}

func (s *LBService) CountLBVServerAppQOEPolicyBinding(name string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lbVServerAppQOEPolicyBindingURL, name), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbvserver_appqoepolicy_binding"]) > 0 {
		if count, ok := result["lbvserver_appqoepolicy_binding"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbvserver_auditnslogpolicy_binding
func (s *LBService) AddLBVServerAuditNSLogPolicyBinding(resource map[string]any) error {
	payload := map[string]any{"lbvserver_auditnslogpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbVServerAuditNSLogPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBVServerAuditNSLogPolicyBinding(name string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lbVServerAuditNSLogPolicyBindingURL, name, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetLBVServerAuditNSLogPolicyBinding(name string) ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbVServerAuditNSLogPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbvserver_auditnslogpolicy_binding"], nil
}

func (s *LBService) CountLBVServerAuditNSLogPolicyBinding(name string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lbVServerAuditNSLogPolicyBindingURL, name), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbvserver_auditnslogpolicy_binding"]) > 0 {
		if count, ok := result["lbvserver_auditnslogpolicy_binding"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbvserver_auditsyslogpolicy_binding
func (s *LBService) AddLBVServerAuditSyslogPolicyBinding(resource map[string]any) error {
	payload := map[string]any{"lbvserver_auditsyslogpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbVServerAuditSyslogPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBVServerAuditSyslogPolicyBinding(name string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lbVServerAuditSyslogPolicyBindingURL, name, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetLBVServerAuditSyslogPolicyBinding(name string) ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbVServerAuditSyslogPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbvserver_auditsyslogpolicy_binding"], nil
}

func (s *LBService) CountLBVServerAuditSyslogPolicyBinding(name string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lbVServerAuditSyslogPolicyBindingURL, name), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbvserver_auditsyslogpolicy_binding"]) > 0 {
		if count, ok := result["lbvserver_auditsyslogpolicy_binding"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbvserver_authorizationpolicy_binding
func (s *LBService) AddLBVServerAuthorizationPolicyBinding(resource map[string]any) error {
	payload := map[string]any{"lbvserver_authorizationpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbVServerAuthorizationPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBVServerAuthorizationPolicyBinding(name string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lbVServerAuthorizationPolicyBindingURL, name, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetLBVServerAuthorizationPolicyBinding(name string) ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbVServerAuthorizationPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbvserver_authorizationpolicy_binding"], nil
}

func (s *LBService) CountLBVServerAuthorizationPolicyBinding(name string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lbVServerAuthorizationPolicyBindingURL, name), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbvserver_authorizationpolicy_binding"]) > 0 {
		if count, ok := result["lbvserver_authorizationpolicy_binding"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbvserver_binding
func (s *LBService) GetLBVServerBinding(name string) (map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbVServerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbvserver_binding"]) == 0 {
		return nil, fmt.Errorf("lbvserver_binding %s not found", name)
	}

	return result["lbvserver_binding"][0], nil
}

// lbvserver_botpolicy_binding
func (s *LBService) AddLBVServerBotPolicyBinding(resource map[string]any) error {
	payload := map[string]any{"lbvserver_botpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbVServerBotPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBVServerBotPolicyBinding(name string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lbVServerBotPolicyBindingURL, name, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetLBVServerBotPolicyBinding(name string) ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbVServerBotPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbvserver_botpolicy_binding"], nil
}

func (s *LBService) CountLBVServerBotPolicyBinding(name string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lbVServerBotPolicyBindingURL, name), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbvserver_botpolicy_binding"]) > 0 {
		if count, ok := result["lbvserver_botpolicy_binding"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbvserver_cachepolicy_binding
func (s *LBService) AddLBVServerCachePolicyBinding(resource map[string]any) error {
	payload := map[string]any{"lbvserver_cachepolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbVServerCachePolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBVServerCachePolicyBinding(name string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lbVServerCachePolicyBindingURL, name, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetLBVServerCachePolicyBinding(name string) ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbVServerCachePolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbvserver_cachepolicy_binding"], nil
}

func (s *LBService) CountLBVServerCachePolicyBinding(name string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lbVServerCachePolicyBindingURL, name), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbvserver_cachepolicy_binding"]) > 0 {
		if count, ok := result["lbvserver_cachepolicy_binding"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbvserver_cmppolicy_binding
func (s *LBService) AddLBVServerCMPPolicyBinding(resource map[string]any) error {
	payload := map[string]any{"lbvserver_cmppolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbVServerCMPPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBVServerCMPPolicyBinding(name string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lbVServerCMPPolicyBindingURL, name, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetLBVServerCMPPolicyBinding(name string) ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbVServerCMPPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbvserver_cmppolicy_binding"], nil
}

func (s *LBService) CountLBVServerCMPPolicyBinding(name string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lbVServerCMPPolicyBindingURL, name), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbvserver_cmppolicy_binding"]) > 0 {
		if count, ok := result["lbvserver_cmppolicy_binding"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbvserver_contentinspectionpolicy_binding
func (s *LBService) AddLBVServerContentInspectionPolicyBinding(resource map[string]any) error {
	payload := map[string]any{"lbvserver_contentinspectionpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbVServerContentInspectionPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBVServerContentInspectionPolicyBinding(name string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lbVServerContentInspectionPolicyBindingURL, name, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetLBVServerContentInspectionPolicyBinding(name string) ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbVServerContentInspectionPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbvserver_contentinspectionpolicy_binding"], nil
}

func (s *LBService) CountLBVServerContentInspectionPolicyBinding(name string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lbVServerContentInspectionPolicyBindingURL, name), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbvserver_contentinspectionpolicy_binding"]) > 0 {
		if count, ok := result["lbvserver_contentinspectionpolicy_binding"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbvserver_csvserver_binding
func (s *LBService) GetLBVServerCSVServerBinding(name string) ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbVServerCSVServerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbvserver_csvserver_binding"], nil
}

func (s *LBService) CountLBVServerCSVServerBinding(name string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lbVServerCSVServerBindingURL, name), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbvserver_csvserver_binding"]) > 0 {
		if count, ok := result["lbvserver_csvserver_binding"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbvserver_dnspolicy64_binding
func (s *LBService) AddLBVServerDNSPolicy64Binding(resource map[string]any) error {
	payload := map[string]any{"lbvserver_dnspolicy64_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbVServerDNSPolicy64BindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBVServerDNSPolicy64Binding(name string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lbVServerDNSPolicy64BindingURL, name, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetLBVServerDNSPolicy64Binding(name string) ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbVServerDNSPolicy64BindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbvserver_dnspolicy64_binding"], nil
}

func (s *LBService) CountLBVServerDNSPolicy64Binding(name string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lbVServerDNSPolicy64BindingURL, name), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbvserver_dnspolicy64_binding"]) > 0 {
		if count, ok := result["lbvserver_dnspolicy64_binding"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbvserver_feopolicy_binding
func (s *LBService) AddLBVServerFEOPolicyBinding(resource map[string]any) error {
	payload := map[string]any{"lbvserver_feopolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbVServerFEOPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBVServerFEOPolicyBinding(name string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lbVServerFEOPolicyBindingURL, name, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetLBVServerFEOPolicyBinding(name string) ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbVServerFEOPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbvserver_feopolicy_binding"], nil
}

func (s *LBService) CountLBVServerFEOPolicyBinding(name string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lbVServerFEOPolicyBindingURL, name), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbvserver_feopolicy_binding"]) > 0 {
		if count, ok := result["lbvserver_feopolicy_binding"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbvserver_responderpolicy_binding
func (s *LBService) AddLBVServerResponderPolicyBinding(resource map[string]any) error {
	payload := map[string]any{"lbvserver_responderpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbVServerResponderPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBVServerResponderPolicyBinding(name string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lbVServerResponderPolicyBindingURL, name, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetLBVServerResponderPolicyBinding(name string) ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbVServerResponderPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbvserver_responderpolicy_binding"], nil
}

func (s *LBService) CountLBVServerResponderPolicyBinding(name string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lbVServerResponderPolicyBindingURL, name), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbvserver_responderpolicy_binding"]) > 0 {
		if count, ok := result["lbvserver_responderpolicy_binding"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbvserver_rewritepolicy_binding
func (s *LBService) AddLBVServerRewritePolicyBinding(resource map[string]any) error {
	payload := map[string]any{"lbvserver_rewritepolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbVServerRewritePolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBVServerRewritePolicyBinding(name string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lbVServerRewritePolicyBindingURL, name, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetLBVServerRewritePolicyBinding(name string) ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbVServerRewritePolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbvserver_rewritepolicy_binding"], nil
}

func (s *LBService) CountLBVServerRewritePolicyBinding(name string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lbVServerRewritePolicyBindingURL, name), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbvserver_rewritepolicy_binding"]) > 0 {
		if count, ok := result["lbvserver_rewritepolicy_binding"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbvserver_servicegroupmember_binding
func (s *LBService) GetLBVServerServiceGroupMemberBinding(lbvserver string) ([]models.LBVServerServiceGroupMemberBinding, error) {
	u := fmt.Sprintf("%s/%s", lbVServerServiceGroupMemberBindingURL, lbvserver)

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Binding []models.LBVServerServiceGroupMemberBinding `json:"lbvserver_servicegroupmember_binding"`
	}

	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Binding, nil
}

func (s *LBService) CountLBVServerServiceGroupMemberBinding(name string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lbVServerServiceGroupMemberBindingURL, name), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbvserver_servicegroupmember_binding"]) > 0 {
		if count, ok := result["lbvserver_servicegroupmember_binding"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbvserver_servicegroup_binding
func (s *LBService) AddLBVServerServiceGroupBinding(resource map[string]any) error {
	payload := map[string]any{"lbvserver_servicegroup_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbVServerServiceGroupBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBVServerServiceGroupBinding(name string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lbVServerServiceGroupBindingURL, name, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetLBVServerServiceGroupBinding(name string) ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbVServerServiceGroupBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbvserver_servicegroup_binding"], nil
}

func (s *LBService) CountLBVServerServiceGroupBinding(name string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lbVServerServiceGroupBindingURL, name), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbvserver_servicegroup_binding"]) > 0 {
		if count, ok := result["lbvserver_servicegroup_binding"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbvserver_service_binding
func (s *LBService) AddLBVServerServiceBinding(resource map[string]any) error {
	payload := map[string]any{"lbvserver_service_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbVServerServiceBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBVServerServiceBinding(name string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lbVServerServiceBindingURL, name, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetLBVServerServiceBinding(name string) ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbVServerServiceBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbvserver_service_binding"], nil
}

func (s *LBService) CountLBVServerServiceBinding(name string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lbVServerServiceBindingURL, name), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbvserver_service_binding"]) > 0 {
		if count, ok := result["lbvserver_service_binding"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbvserver_spilloverpolicy_binding
func (s *LBService) AddLBVServerSpilloverPolicyBinding(resource map[string]any) error {
	payload := map[string]any{"lbvserver_spilloverpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbVServerSpilloverPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBVServerSpilloverPolicyBinding(name string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lbVServerSpilloverPolicyBindingURL, name, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetLBVServerSpilloverPolicyBinding(name string) ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbVServerSpilloverPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbvserver_spilloverpolicy_binding"], nil
}

func (s *LBService) CountLBVServerSpilloverPolicyBinding(name string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lbVServerSpilloverPolicyBindingURL, name), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbvserver_spilloverpolicy_binding"]) > 0 {
		if count, ok := result["lbvserver_spilloverpolicy_binding"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbvserver_tmtrafficpolicy_binding
func (s *LBService) AddLBVServerTMTrafficPolicyBinding(resource map[string]any) error {
	payload := map[string]any{"lbvserver_tmtrafficpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbVServerTMTrafficPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBVServerTMTrafficPolicyBinding(name string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lbVServerTMTrafficPolicyBindingURL, name, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetLBVServerTMTrafficPolicyBinding(name string) ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbVServerTMTrafficPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbvserver_tmtrafficpolicy_binding"], nil
}

func (s *LBService) CountLBVServerTMTrafficPolicyBinding(name string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lbVServerTMTrafficPolicyBindingURL, name), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbvserver_tmtrafficpolicy_binding"]) > 0 {
		if count, ok := result["lbvserver_tmtrafficpolicy_binding"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbvserver_transformpolicy_binding
func (s *LBService) AddLBVServerTransformPolicyBinding(resource map[string]any) error {
	payload := map[string]any{"lbvserver_transformpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbVServerTransformPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBVServerTransformPolicyBinding(name string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lbVServerTransformPolicyBindingURL, name, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetLBVServerTransformPolicyBinding(name string) ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbVServerTransformPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbvserver_transformpolicy_binding"], nil
}

func (s *LBService) CountLBVServerTransformPolicyBinding(name string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lbVServerTransformPolicyBindingURL, name), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbvserver_transformpolicy_binding"]) > 0 {
		if count, ok := result["lbvserver_transformpolicy_binding"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbvserver_videooptimizationdetectionpolicy_binding
func (s *LBService) AddLBVServerVideoOptimizationDetectionPolicyBinding(resource map[string]any) error {
	payload := map[string]any{"lbvserver_videooptimizationdetectionpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbVServerVideoOptimizationDetectionPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBVServerVideoOptimizationDetectionPolicyBinding(name string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lbVServerVideoOptimizationDetectionPolicyBindingURL, name, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetLBVServerVideoOptimizationDetectionPolicyBinding(name string) ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbVServerVideoOptimizationDetectionPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbvserver_videooptimizationdetectionpolicy_binding"], nil
}

func (s *LBService) CountLBVServerVideoOptimizationDetectionPolicyBinding(name string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lbVServerVideoOptimizationDetectionPolicyBindingURL, name), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbvserver_videooptimizationdetectionpolicy_binding"]) > 0 {
		if count, ok := result["lbvserver_videooptimizationdetectionpolicy_binding"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbvserver_videooptimizationpacingpolicy_binding
func (s *LBService) AddLBVServerVideoOptimizationPacingPolicyBinding(resource map[string]any) error {
	payload := map[string]any{"lbvserver_videooptimizationpacingpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbVServerVideoOptimizationPacingPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBVServerVideoOptimizationPacingPolicyBinding(name string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lbVServerVideoOptimizationPacingPolicyBindingURL, name, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetLBVServerVideoOptimizationPacingPolicyBinding(name string) ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbVServerVideoOptimizationPacingPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbvserver_videooptimizationpacingpolicy_binding"], nil
}

func (s *LBService) CountLBVServerVideoOptimizationPacingPolicyBinding(name string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lbVServerVideoOptimizationPacingPolicyBindingURL, name), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbvserver_videooptimizationpacingpolicy_binding"]) > 0 {
		if count, ok := result["lbvserver_videooptimizationpacingpolicy_binding"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbwlm
func (s *LBService) AddLBWLM(resource map[string]any) error {
	payload := map[string]any{"lbwlm": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbWLMURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBWLM(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", lbWLMURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) UpdateLBWLM(resource map[string]any) error {
	payload := map[string]any{"lbwlm": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, lbWLMURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) UnsetLBWLM(resource map[string]any) error {
	payload := map[string]any{"lbwlm": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", lbWLMURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetAllLBWLM() ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, lbWLMURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbwlm"], nil
}

func (s *LBService) GetLBWLM(name string) (map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbWLMURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbwlm"]) == 0 {
		return nil, fmt.Errorf("lbwlm %s not found", name)
	}

	return result["lbwlm"][0], nil
}

func (s *LBService) CountLBWLM() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", lbWLMURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbwlm"]) > 0 {
		if count, ok := result["lbwlm"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}

// lbwlm_binding
func (s *LBService) GetLBWLMBinding(name string) (map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbWLMBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbwlm_binding"]) == 0 {
		return nil, fmt.Errorf("lbwlm_binding %s not found", name)
	}

	return result["lbwlm_binding"][0], nil
}

// lbwlm_lbvserver_binding
func (s *LBService) AddLBWLMLBVServerBinding(resource map[string]any) error {
	payload := map[string]any{"lbwlm_lbvserver_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lbWLMLBVServerBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) DeleteLBWLMLBVServerBinding(name string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lbWLMLBVServerBindingURL, name, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LBService) GetLBWLMLBVServerBinding(name string) ([]map[string]any, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lbWLMLBVServerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result["lbwlm_lbvserver_binding"], nil
}

func (s *LBService) CountLBWLMLBVServerBinding(name string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lbWLMLBVServerBindingURL, name), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string][]map[string]any
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result["lbwlm_lbvserver_binding"]) > 0 {
		if count, ok := result["lbwlm_lbvserver_binding"][0]["__count"].(float64); ok {
			return int(count), nil
		}
	}
	return 0, fmt.Errorf("failed to parse count")
}
