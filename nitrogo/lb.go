package nitrogo

import (
	"encoding/json"
	"fmt"
	"net/http"

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
func (s *LBService) AddLBAction()    {}
func (s *LBService) DeleteLBAction() {}
func (s *LBService) UpdateLBAction() {}
func (s *LBService) UnsetLBAction()  {}
func (s *LBService) GetAllLBAction() {}
func (s *LBService) GetLBAction()    {}
func (s *LBService) CountLBAction()  {}
func (s *LBService) RenameLBAction() {}

// lbglobal_binding
func (s *LBService) GetLBGlobalBinding() {}

// lbglobal_lbpolicy_binding
func (s *LBService) AddLBGlobalLBPolicyBinding()    {}
func (s *LBService) DeleteLBGlobalLBPolicyBinding() {}
func (s *LBService) GetLBGlobalLBPolicyBinding()    {}
func (s *LBService) CountLBGlobalLBPolicyBinding()  {}

// lbgroup
func (s *LBService) AddLBGroup()    {}
func (s *LBService) DeleteLBGroup() {}
func (s *LBService) UpdateLBGroup() {}
func (s *LBService) UnsetLBGroup()  {}
func (s *LBService) GetAllLBGroup() {}
func (s *LBService) GetLBGroup()    {}
func (s *LBService) CountLBGroup()  {}
func (s *LBService) RenameLBGroup() {}

// lbgroup_binding
func (s *LBService) GetAllLBGroupBinding() {}
func (s *LBService) GetLBGroupBinding()    {}

// lbgroup_lbvserver_binding
func (s *LBService) AddLBGroupLBVServerBinding()    {}
func (s *LBService) DeleteLBGroupLBVServerBinding() {}
func (s *LBService) GetAllLBGroupLBVServerBinding() {}
func (s *LBService) GetLBGroupLBVServerBinding()    {}
func (s *LBService) CountLBGroupLBVServerBinding()  {}

// lbmetrictable
func (s *LBService) AddLBMetricTable()    {}
func (s *LBService) DeleteLBMetricTable() {}
func (s *LBService) UpdateLBMetricTable() {}
func (s *LBService) GetAllLBMetricTable() {}
func (s *LBService) GetLBMetricTable()    {}
func (s *LBService) CountLBMetricTable()  {}

// lbmetrictable_binding
func (s *LBService) GetAllLBMetricTableBinding() {}
func (s *LBService) GetLBMetricTableBinding()    {}

// lbmetrictable_metric_binding
func (s *LBService) AddLBMetricTableMetricBinding()    {}
func (s *LBService) DeleteLBMetricTableMetricBinding() {}
func (s *LBService) GetAllLBMetricTableMetricBinding() {}
func (s *LBService) GetLBMetricTableMetricBinding()    {}
func (s *LBService) CountLBMetricTableMetricBinding()  {}

// lbmonbindings
func (s *LBService) GetAllLBMonBindings() {}
func (s *LBService) GetLBMonBindings()    {}
func (s *LBService) CountLBMonBindings()  {}

// lbmonbindings_binding
func (s *LBService) GetLBMonBindingsBinding() {}

// lbmonbindings_gslbservicegroup_binding
func (s *LBService) GetLBMonBindingsGSLBServiceGroupBinding()   {}
func (s *LBService) CountLBMonBindingsGSLBServiceGroupBinding() {}

// lbmonbindings_servicegroup_binding
func (s *LBService) GetLBMonBindingsServiceGroupBinding()   {}
func (s *LBService) CountLBMonBindingsServiceGroupBinding() {}

// lbmonbindings_service_binding
func (s *LBService) GetLBMonBindingsServiceBinding()   {}
func (s *LBService) CountLBMonBindingsServiceBinding() {}

// lbmonitor
func (s *LBService) AddLBMonitor()     {}
func (s *LBService) DeleteLBMonitor()  {}
func (s *LBService) UpdateLBMonitor()  {}
func (s *LBService) UnsetLBMonitor()   {}
func (s *LBService) EnableLBMonitor()  {}
func (s *LBService) DisableLBMonitor() {}
func (s *LBService) GetAllLBMonitor()  {}
func (s *LBService) GetLBMonitor()     {}
func (s *LBService) CountLBMonitor()   {}

// lbmonitor_binding
func (s *LBService) GetAllLBMonitorBinding() {}
func (s *LBService) GetLBMonitorBinding()    {}

// lbmonitor_metric_binding
func (s *LBService) AddLBMonitorMetricBinding()    {}
func (s *LBService) DeleteLBMonitorMetricBinding() {}
func (s *LBService) GetAllLBMonitorMetricBinding() {}
func (s *LBService) GetLBMonitorMetricBinding()    {}
func (s *LBService) CountLBMonitorMetricBinding()  {}

// lbmonitor_servicegroup_binding
func (s *LBService) AddLBMonitorServiceGroupBinding()    {}
func (s *LBService) DeleteLBMonitorServiceGroupBinding() {}

// lbmonitor_service_binding
func (s *LBService) AddLBMonitorServiceBinding()    {}
func (s *LBService) DeleteLBMonitorServiceBinding() {}

// lbmonitor_sslcertkey_binding
func (s *LBService) AddLBMonitorSSLCertKeyBinding()    {}
func (s *LBService) DeleteLBMonitorSSLCertKeyBinding() {}
func (s *LBService) GetAllLBMonitorSSLCertKeyBinding() {}
func (s *LBService) GetLBMonitorSSLCertKeyBinding()    {}
func (s *LBService) CountLBMonitorSSLCertKeyBinding()  {}

// lbparameter
func (s *LBService) UpdateLBParameter() {}
func (s *LBService) UnsetLBParameter()  {}
func (s *LBService) GetAllLBParameter() {}

// lbpersistentsessions
func (s *LBService) GetAllLBPersistentSessions() {}
func (s *LBService) CountLBPersistentSessions()  {}
func (s *LBService) ClearLBPersistentSessions()  {}

// lbprofile
func (s *LBService) AddLBProfile()    {}
func (s *LBService) DeleteLBProfile() {}
func (s *LBService) UpdateLBProfile() {}
func (s *LBService) UnsetLBProfile()  {}
func (s *LBService) GetAllLBProfile() {}
func (s *LBService) GetLBProfile()    {}
func (s *LBService) CountLBProfile()  {}

// lbroute
func (s *LBService) AddLBRoute()    {}
func (s *LBService) DeleteLBRoute() {}
func (s *LBService) GetAllLBRoute() {}
func (s *LBService) CountLBRoute()  {}

// lbroute6
func (s *LBService) AddLBRoute6()    {}
func (s *LBService) DeleteLBRoute6() {}
func (s *LBService) GetAllLBRoute6() {}
func (s *LBService) CountLBRoute6()  {}

// lbsipparameters
func (s *LBService) UpdateLBSIPParameters() {}
func (s *LBService) UnsetLBSIPParameters()  {}
func (s *LBService) GetAllLBSIPParameters() {}

// lbvserver
func (s *LBService) AddLBVServer()     {}
func (s *LBService) DeleteLBVServer()  {}
func (s *LBService) UpdateLBVServer()  {}
func (s *LBService) UnsetLBVServer()   {}
func (s *LBService) EnableLBVServer()  {}
func (s *LBService) DisableLBVServer() {}
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
func (s *LBService) GetLBVServer()    {}
func (s *LBService) CountLBVServer()  {}
func (s *LBService) RenameLBVServer() {}

// lbvserver_analyticsprofile_binding
func (s *LBService) AddLBVServerAnalyticsProfileBinding()    {}
func (s *LBService) DeleteLBVServerAnalyticsProfileBinding() {}
func (s *LBService) GetAllLBVServerAnalyticsProfileBinding() {}
func (s *LBService) GetLBVServerAnalyticsProfileBinding()    {}
func (s *LBService) CountLBVServerAnalyticsProfileBinding()  {}

// lbvserver_appflowpolicy_binding
func (s *LBService) AddLBVServerAppFlowPolicyBinding()    {}
func (s *LBService) DeleteLBVServerAppFlowPolicyBinding() {}
func (s *LBService) GetAllLBVServerAppFlowPolicyBinding() {}
func (s *LBService) GetLBVServerAppFlowPolicyBinding()    {}
func (s *LBService) CountLBVServerAppFlowPolicyBinding()  {}

// lbvserver_appfwpolicy_binding
func (s *LBService) AddLBVServerAppFWPolicyBinding()    {}
func (s *LBService) DeleteLBVServerAppFWPolicyBinding() {}
func (s *LBService) GetAllLBVServerAppFWPolicyBinding() {}
func (s *LBService) GetLBVServerAppFWPolicyBinding()    {}
func (s *LBService) CountLBVServerAppFWPolicyBinding()  {}

// lbvserver_appqoepolicy_binding
func (s *LBService) AddLBVServerAppQOEPolicyBinding()    {}
func (s *LBService) DeleteLBVServerAppQOEPolicyBinding() {}
func (s *LBService) GetAllLBVServerAppQOEPolicyBinding() {}
func (s *LBService) GetLBVServerAppQOEPolicyBinding()    {}
func (s *LBService) CountLBVServerAppQOEPolicyBinding()  {}

// lbvserver_auditnslogpolicy_binding
func (s *LBService) AddLBVServerAuditNSLogPolicyBinding()    {}
func (s *LBService) DeleteLBVServerAuditNSLogPolicyBinding() {}
func (s *LBService) GetAllLBVServerAuditNSLogPolicyBinding() {}
func (s *LBService) GetLBVServerAuditNSLogPolicyBinding()    {}
func (s *LBService) CountLBVServerAuditNSLogPolicyBinding()  {}

// lbvserver_auditsyslogpolicy_binding
func (s *LBService) AddLBVServerAuditSyslogPolicyBinding()    {}
func (s *LBService) DeleteLBVServerAuditSyslogPolicyBinding() {}
func (s *LBService) GetAllLBVServerAuditSyslogPolicyBinding() {}
func (s *LBService) GetLBVServerAuditSyslogPolicyBinding()    {}
func (s *LBService) CountLBVServerAuditSyslogPolicyBinding()  {}

// lbvserver_authorizationpolicy_binding
func (s *LBService) AddLBVServerAuthorizationPolicyBinding()    {}
func (s *LBService) DeleteLBVServerAuthorizationPolicyBinding() {}
func (s *LBService) GetAllLBVServerAuthorizationPolicyBinding() {}
func (s *LBService) GetLBVServerAuthorizationPolicyBinding()    {}
func (s *LBService) CountLBVServerAuthorizationPolicyBinding()  {}

// lbvserver_binding
func (s *LBService) GetAllLBVServerBinding() {}
func (s *LBService) GetLBVServerBinding()    {}

// lbvserver_botpolicy_binding
func (s *LBService) AddLBVServerBotPolicyBinding()    {}
func (s *LBService) DeleteLBVServerBotPolicyBinding() {}
func (s *LBService) GetAllLBVServerBotPolicyBinding() {}
func (s *LBService) GetLBVServerBotPolicyBinding()    {}
func (s *LBService) CountLBVServerBotPolicyBinding()  {}

// lbvserver_cachepolicy_binding
func (s *LBService) AddLBVServerCachePolicyBinding()    {}
func (s *LBService) DeleteLBVServerCachePolicyBinding() {}
func (s *LBService) GetAllLBVServerCachePolicyBinding() {}
func (s *LBService) GetLBVServerCachePolicyBinding()    {}
func (s *LBService) CountLBVServerCachePolicyBinding()  {}

// lbvserver_cmppolicy_binding
func (s *LBService) AddLBVServerCMPPolicyBinding()    {}
func (s *LBService) DeleteLBVServerCMPPolicyBinding() {}
func (s *LBService) GetAllLBVServerCMPPolicyBinding() {}
func (s *LBService) GetLBVServerCMPPolicyBinding()    {}
func (s *LBService) CountLBVServerCMPPolicyBinding()  {}

// lbvserver_contentinspectionpolicy_binding
func (s *LBService) AddLBVServerContentInspectionPolicyBinding()    {}
func (s *LBService) DeleteLBVServerContentInspectionPolicyBinding() {}
func (s *LBService) GetAllLBVServerContentInspectionPolicyBinding() {}
func (s *LBService) GetLBVServerContentInspectionPolicyBinding()    {}
func (s *LBService) CountLBVServerContentInspectionPolicyBinding()  {}

// lbvserver_csvserver_binding
func (s *LBService) GetAllLBVServerCSVServerBinding() {}
func (s *LBService) GetLBVServerCSVServerBinding()    {}
func (s *LBService) CountLBVServerCSVServerBinding()  {}

// lbvserver_dnspolicy64_binding
func (s *LBService) AddLBVServerDNSPolicy64Binding()    {}
func (s *LBService) DeleteLBVServerDNSPolicy64Binding() {}
func (s *LBService) GetAllLBVServerDNSPolicy64Binding() {}
func (s *LBService) GetLBVServerDNSPolicy64Binding()    {}
func (s *LBService) CountLBVServerDNSPolicy64Binding()  {}

// lbvserver_feopolicy_binding
func (s *LBService) AddLBVServerFEOPolicyBinding()    {}
func (s *LBService) DeleteLBVServerFEOPolicyBinding() {}
func (s *LBService) GetAllLBVServerFEOPolicyBinding() {}
func (s *LBService) GetLBVServerFEOPolicyBinding()    {}
func (s *LBService) CountLBVServerFEOPolicyBinding()  {}

// lbvserver_responderpolicy_binding
func (s *LBService) AddLBVServerResponderPolicyBinding()    {}
func (s *LBService) DeleteLBVServerResponderPolicyBinding() {}
func (s *LBService) GetAllLBVServerResponderPolicyBinding() {}
func (s *LBService) GetLBVServerResponderPolicyBinding()    {}
func (s *LBService) CountLBVServerResponderPolicyBinding()  {}

// lbvserver_rewritepolicy_binding
func (s *LBService) AddLBVServerRewritePolicyBinding()    {}
func (s *LBService) DeleteLBVServerRewritePolicyBinding() {}
func (s *LBService) GetAllLBVServerRewritePolicyBinding() {}
func (s *LBService) GetLBVServerRewritePolicyBinding()    {}
func (s *LBService) CountLBVServerRewritePolicyBinding()  {}

// lbvserver_servicegroupmember_binding
func (s *LBService) GetAllLBVServerServiceGroupMemberBinding() {}
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
func (s *LBService) CountLBVServerServiceGroupMemberBinding() {}

// lbvserver_servicegroup_binding
func (s *LBService) AddLBVServerServiceGroupBinding()    {}
func (s *LBService) DeleteLBVServerServiceGroupBinding() {}
func (s *LBService) GetAllLBVServerServiceGroupBinding() {}
func (s *LBService) GetLBVServerServiceGroupBinding()    {}
func (s *LBService) CountLBVServerServiceGroupBinding()  {}

// lbvserver_service_binding
func (s *LBService) AddLBVServerServiceBinding()    {}
func (s *LBService) DeleteLBVServerServiceBinding() {}
func (s *LBService) GetAllLBVServerServiceBinding() {}
func (s *LBService) GetLBVServerServiceBinding()    {}
func (s *LBService) CountLBVServerServiceBinding()  {}

// lbvserver_spilloverpolicy_binding
func (s *LBService) AddLBVServerSpilloverPolicyBinding()    {}
func (s *LBService) DeleteLBVServerSpilloverPolicyBinding() {}
func (s *LBService) GetAllLBVServerSpilloverPolicyBinding() {}
func (s *LBService) GetLBVServerSpilloverPolicyBinding()    {}
func (s *LBService) CountLBVServerSpilloverPolicyBinding()  {}

// lbvserver_tmtrafficpolicy_binding
func (s *LBService) AddLBVServerTMTrafficPolicyBinding()    {}
func (s *LBService) DeleteLBVServerTMTrafficPolicyBinding() {}
func (s *LBService) GetAllLBVServerTMTrafficPolicyBinding() {}
func (s *LBService) GetLBVServerTMTrafficPolicyBinding()    {}
func (s *LBService) CountLBVServerTMTrafficPolicyBinding()  {}

// lbvserver_transformpolicy_binding
func (s *LBService) AddLBVServerTransformPolicyBinding()    {}
func (s *LBService) DeleteLBVServerTransformPolicyBinding() {}
func (s *LBService) GetAllLBVServerTransformPolicyBinding() {}
func (s *LBService) GetLBVServerTransformPolicyBinding()    {}
func (s *LBService) CountLBVServerTransformPolicyBinding()  {}

// lbvserver_videooptimizationdetectionpolicy_binding
func (s *LBService) AddLBVServerVideoOptimizationDetectionPolicyBinding()    {}
func (s *LBService) DeleteLBVServerVideoOptimizationDetectionPolicyBinding() {}
func (s *LBService) GetAllLBVServerVideoOptimizationDetectionPolicyBinding() {}
func (s *LBService) GetLBVServerVideoOptimizationDetectionPolicyBinding()    {}
func (s *LBService) CountLBVServerVideoOptimizationDetectionPolicyBinding()  {}

// lbvserver_videooptimizationpacingpolicy_binding
func (s *LBService) AddLBVServerVideoOptimizationPacingPolicyBinding()    {}
func (s *LBService) DeleteLBVServerVideoOptimizationPacingPolicyBinding() {}
func (s *LBService) GetAllLBVServerVideoOptimizationPacingPolicyBinding() {}
func (s *LBService) GetLBVServerVideoOptimizationPacingPolicyBinding()    {}
func (s *LBService) CountLBVServerVideoOptimizationPacingPolicyBinding()  {}

// lbwlm
func (s *LBService) AddLBWLM()    {}
func (s *LBService) DeleteLBWLM() {}
func (s *LBService) UpdateLBWLM() {}
func (s *LBService) UnsetLBWLM()  {}
func (s *LBService) GetAllLBWLM() {}
func (s *LBService) GetLBWLM()    {}
func (s *LBService) CountLBWLM()  {}

// lbwlm_binding
func (s *LBService) GetAllLBWLMBinding() {}
func (s *LBService) GetLBWLMBinding()    {}

// lbwlm_lbvserver_binding
func (s *LBService) AddLBWLMLBVServerBinding()    {}
func (s *LBService) DeleteLBWLMLBVServerBinding() {}
func (s *LBService) GetAllLBWLMLBVServerBinding() {}
func (s *LBService) GetLBWLMLBVServerBinding()    {}
func (s *LBService) CountLBWLMLBVServerBinding()  {}
