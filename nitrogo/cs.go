package nitrogo

const (
	csActionURL                                = "/nitro/v1/config/csaction"
	csParameterURL                             = "/nitro/v1/config/csparameter"
	csPolicyURL                                = "/nitro/v1/config/cspolicy"
	csPolicyLabelURL                           = "/nitro/v1/config/cspolicylabel"
	csPolicyLabelBindingURL                    = "/nitro/v1/config/cspolicylabel_binding"
	csPolicyLabelCSPolicyBindingURL            = "/nitro/v1/config/cspolicylabel_cspolicy_binding"
	csPolicyBindingURL                         = "/nitro/v1/config/cspolicy_binding"
	csPolicyCRVServerBindingURL                = "/nitro/v1/config/cspolicy_crvserver_binding"
	csPolicyCSPolicyLabelBindingURL            = "/nitro/v1/config/cspolicy_cspolicylabel_binding"
	csPolicyCSVServerBindingURL                = "/nitro/v1/config/cspolicy_csvserver_binding"
	csVServerURL                               = "/nitro/v1/config/csvserver"
	csVServerAnalyticsProfileBindingURL        = "/nitro/v1/config/csvserver_analyticsprofile_binding"
	csVServerAppFlowPolicyBindingURL           = "/nitro/v1/config/csvserver_appflowpolicy_binding"
	csVServerAppFWPolicyBindingURL             = "/nitro/v1/config/csvserver_appfwpolicy_binding"
	csVServerAppQOEPolicyBindingURL            = "/nitro/v1/config/csvserver_appqoepolicy_binding"
	csVServerAuditNSLogPolicyBindingURL        = "/nitro/v1/config/csvserver_auditnslogpolicy_binding"
	csVServerAuditSyslogPolicyBindingURL       = "/nitro/v1/config/csvserver_auditsyslogpolicy_binding"
	csVServerAuthorizationPolicyBindingURL     = "/nitro/v1/config/csvserver_authorizationpolicy_binding"
	csVServerBindingURL                        = "/nitro/v1/config/csvserver_binding"
	csVServerBotPolicyBindingURL               = "/nitro/v1/config/csvserver_botpolicy_binding"
	csVServerCachePolicyBindingURL             = "/nitro/v1/config/csvserver_cachepolicy_binding"
	csVServerCMPPolicyBindingURL               = "/nitro/v1/config/csvserver_cmppolicy_binding"
	csVServerContentInspectionPolicyBindingURL = "/nitro/v1/config/csvserver_contentinspectionpolicy_binding"
	csVServerCSPolicyBindingURL                = "/nitro/v1/config/csvserver_cspolicy_binding"
	csVServerDomainBindingURL                  = "/nitro/v1/config/csvserver_domain_binding"
	csVServerFEOPolicyBindingURL               = "/nitro/v1/config/csvserver_feopolicy_binding"
	csVServerGSLBVServerBindingURL             = "/nitro/v1/config/csvserver_gslbvserver_binding"
	csVServerLBVServerBindingURL               = "/nitro/v1/config/csvserver_lbvserver_binding"
	csVServerResponderPolicyBindingURL         = "/nitro/v1/config/csvserver_responderpolicy_binding"
	csVServerRewritePolicyBindingURL           = "/nitro/v1/config/csvserver_rewritepolicy_binding"
	csVServerSpilloverPolicyBindingURL         = "/nitro/v1/config/csvserver_spilloverpolicy_binding"
	csVServerTMTrafficPolicyBindingURL         = "/nitro/v1/config/csvserver_tmtrafficpolicy_binding"
	csVServerTransformPolicyBindingURL         = "/nitro/v1/config/csvserver_transformpolicy_binding"
	csVServerVPNVServerBindingURL              = "/nitro/v1/config/csvserver_vpnvserver_binding"
)

// Content Switching configuration. Content Switching feature that enables you to direct traffic to servers on the basis of content.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/cs/cs
type CSService struct {
	client *Client
}

// csaction
func (s *CSService) AddCSAction()    {}
func (s *CSService) DeleteCSAction() {}
func (s *CSService) UpdateCSAction() {}
func (s *CSService) UnsetCSAction()  {}
func (s *CSService) GetAllCSAction() {}
func (s *CSService) GetCSAction()    {}
func (s *CSService) CountCSAction()  {}
func (s *CSService) RenameCSAction() {}

// csparameter
func (s *CSService) UpdateCSParameter() {}
func (s *CSService) UnsetCSParameter()  {}
func (s *CSService) GetAllCSParameter() {}

// cspolicy
func (s *CSService) AddCSPolicy()    {}
func (s *CSService) DeleteCSPolicy() {}
func (s *CSService) UpdateCSPolicy() {}
func (s *CSService) UnsetCSPolicy()  {}
func (s *CSService) GetAllCSPolicy() {}
func (s *CSService) GetCSPolicy()    {}
func (s *CSService) CountCSPolicy()  {}
func (s *CSService) RenameCSPolicy() {}

// cspolicylabel
func (s *CSService) AddCSPolicyLabel()    {}
func (s *CSService) DeleteCSPolicyLabel() {}
func (s *CSService) GetAllCSPolicyLabel() {}
func (s *CSService) GetCSPolicyLabel()    {}
func (s *CSService) CountCSPolicyLabel()  {}
func (s *CSService) RenameCSPolicyLabel() {}

// cspolicylabel_binding
func (s *CSService) GetAllCSPolicyLabelBinding() {}
func (s *CSService) GetCSPolicyLabelBinding()    {}

// cspolicylabel_cspolicy_binding
func (s *CSService) AddCSPolicyLabelCSPolicyBinding()    {}
func (s *CSService) DeleteCSPolicyLabelCSPolicyBinding() {}
func (s *CSService) GetAllCSPolicyLabelCSPolicyBinding() {}
func (s *CSService) GetCSPolicyLabelCSPolicyBinding()    {}
func (s *CSService) CountCSPolicyLabelCSPolicyBinding()  {}

// cspolicy_binding
func (s *CSService) GetAllCSPolicyBinding() {}
func (s *CSService) GetCSPolicyBinding()    {}

// cspolicy_crvserver_binding
func (s *CSService) GetAllCSPolicyCRVServerBinding() {}
func (s *CSService) GetCSPolicyCRVServerBinding()    {}
func (s *CSService) CountCSPolicyCRVServerBinding()  {}

// cspolicy_cspolicylabel_binding
func (s *CSService) GetAllCSPolicyCSPolicyLabelBinding() {}
func (s *CSService) GetCSPolicyCSPolicyLabelBinding()    {}
func (s *CSService) CountCSPolicyCSPolicyLabelBinding()  {}

// cspolicy_csvserver_binding
func (s *CSService) GetAllCSPolicyCSVServerBinding() {}
func (s *CSService) GetCSPolicyCSVServerBinding()    {}
func (s *CSService) CountCSPolicyCSVServerBinding()  {}

// csvserver
func (s *CSService) AddCSVServer()     {}
func (s *CSService) DeleteCSVServer()  {}
func (s *CSService) UpdateCSVServer()  {}
func (s *CSService) UnsetCSVServer()   {}
func (s *CSService) EnableCSVServer()  {}
func (s *CSService) DisableCSVServer() {}
func (s *CSService) GetAllCSVServer()  {}
func (s *CSService) GetCSVServer()     {}
func (s *CSService) CountCSVServer()   {}
func (s *CSService) RenameCSVServer()  {}

// csvserver_analyticsprofile_binding
func (s *CSService) AddCSVServerAnalyticsProfileBinding()    {}
func (s *CSService) DeleteCSVServerAnalyticsProfileBinding() {}
func (s *CSService) GetAllCSVServerAnalyticsProfileBinding() {}
func (s *CSService) GetCSVServerAnalyticsProfileBinding()    {}
func (s *CSService) CountCSVServerAnalyticsProfileBinding()  {}

// csvserver_appflowpolicy_binding
func (s *CSService) AddCSVServerAppFlowPolicyBinding()    {}
func (s *CSService) DeleteCSVServerAppFlowPolicyBinding() {}
func (s *CSService) GetAllCSVServerAppFlowPolicyBinding() {}
func (s *CSService) GetCSVServerAppFlowPolicyBinding()    {}
func (s *CSService) CountCSVServerAppFlowPolicyBinding()  {}

// csvserver_appfwpolicy_binding
func (s *CSService) AddCSVServerAppFWPolicyBinding()    {}
func (s *CSService) DeleteCSVServerAppFWPolicyBinding() {}
func (s *CSService) GetAllCSVServerAppFWPolicyBinding() {}
func (s *CSService) GetCSVServerAppFWPolicyBinding()    {}
func (s *CSService) CountCSVServerAppFWPolicyBinding()  {}

// csvserver_appqoepolicy_binding
func (s *CSService) AddCSVServerAppQOEPolicyBinding()    {}
func (s *CSService) DeleteCSVServerAppQOEPolicyBinding() {}
func (s *CSService) GetAllCSVServerAppQOEPolicyBinding() {}
func (s *CSService) GetCSVServerAppQOEPolicyBinding()    {}
func (s *CSService) CountCSVServerAppQOEPolicyBinding()  {}

// csvserver_auditnslogpolicy_binding
func (s *CSService) AddCSVServerAuditNSLogPolicyBinding()    {}
func (s *CSService) DeleteCSVServerAuditNSLogPolicyBinding() {}
func (s *CSService) GetAllCSVServerAuditNSLogPolicyBinding() {}
func (s *CSService) GetCSVServerAuditNSLogPolicyBinding()    {}
func (s *CSService) CountCSVServerAuditNSLogPolicyBinding()  {}

// csvserver_auditsyslogpolicy_binding
func (s *CSService) AddCSVServerAuditSyslogPolicyBinding()    {}
func (s *CSService) DeleteCSVServerAuditSyslogPolicyBinding() {}
func (s *CSService) GetAllCSVServerAuditSyslogPolicyBinding() {}
func (s *CSService) GetCSVServerAuditSyslogPolicyBinding()    {}
func (s *CSService) CountCSVServerAuditSyslogPolicyBinding()  {}

// csvserver_authorizationpolicy_binding
func (s *CSService) AddCSVServerAuthorizationPolicyBinding()    {}
func (s *CSService) DeleteCSVServerAuthorizationPolicyBinding() {}
func (s *CSService) GetAllCSVServerAuthorizationPolicyBinding() {}
func (s *CSService) GetCSVServerAuthorizationPolicyBinding()    {}
func (s *CSService) CountCSVServerAuthorizationPolicyBinding()  {}

// csvserver_binding
func (s *CSService) GetAllCSVServerBinding() {}
func (s *CSService) GetCSVServerBinding()    {}

// csvserver_botpolicy_binding
func (s *CSService) AddCSVServerBotPolicyBinding()    {}
func (s *CSService) DeleteCSVServerBotPolicyBinding() {}
func (s *CSService) GetAllCSVServerBotPolicyBinding() {}
func (s *CSService) GetCSVServerBotPolicyBinding()    {}
func (s *CSService) CountCSVServerBotPolicyBinding()  {}

// csvserver_cachepolicy_binding
func (s *CSService) AddCSVServerCachePolicyBinding()    {}
func (s *CSService) DeleteCSVServerCachePolicyBinding() {}
func (s *CSService) GetAllCSVServerCachePolicyBinding() {}
func (s *CSService) GetCSVServerCachePolicyBinding()    {}
func (s *CSService) CountCSVServerCachePolicyBinding()  {}

// csvserver_cmppolicy_binding
func (s *CSService) AddCSVServerCMPPolicyBinding()    {}
func (s *CSService) DeleteCSVServerCMPPolicyBinding() {}
func (s *CSService) GetAllCSVServerCMPPolicyBinding() {}
func (s *CSService) GetCSVServerCMPPolicyBinding()    {}
func (s *CSService) CountCSVServerCMPPolicyBinding()  {}

// csvserver_contentinspectionpolicy_binding
func (s *CSService) AddCSVServerContentInspectionPolicyBinding()    {}
func (s *CSService) DeleteCSVServerContentInspectionPolicyBinding() {}
func (s *CSService) GetAllCSVServerContentInspectionPolicyBinding() {}
func (s *CSService) GetCSVServerContentInspectionPolicyBinding()    {}
func (s *CSService) CountCSVServerContentInspectionPolicyBinding()  {}

// csvserver_cspolicy_binding
func (s *CSService) AddCSVServerCSPolicyBinding()    {}
func (s *CSService) DeleteCSVServerCSPolicyBinding() {}
func (s *CSService) GetAllCSVServerCSPolicyBinding() {}
func (s *CSService) GetCSVServerCSPolicyBinding()    {}
func (s *CSService) CountCSVServerCSPolicyBinding()  {}

// csvserver_domain_binding
func (s *CSService) AddCSVServerDomainBinding()    {}
func (s *CSService) DeleteCSVServerDomainBinding() {}
func (s *CSService) GetAllCSVServerDomainBinding() {}
func (s *CSService) GetCSVServerDomainBinding()    {}
func (s *CSService) CountCSVServerDomainBinding()  {}

// csvserver_feopolicy_binding
func (s *CSService) AddCSVServerFEOPolicyBinding()    {}
func (s *CSService) DeleteCSVServerFEOPolicyBinding() {}
func (s *CSService) GetAllCSVServerFEOPolicyBinding() {}
func (s *CSService) GetCSVServerFEOPolicyBinding()    {}
func (s *CSService) CountCSVServerFEOPolicyBinding()  {}

// csvserver_gslbvserver_binding
func (s *CSService) AddCSVServerGSLBVServerBinding()    {}
func (s *CSService) DeleteCSVServerGSLBVServerBinding() {}
func (s *CSService) GetAllCSVServerGSLBVServerBinding() {}
func (s *CSService) GetCSVServerGSLBVServerBinding()    {}
func (s *CSService) CountCSVServerGSLBVServerBinding()  {}

// csvserver_lbvserver_binding
func (s *CSService) AddCSVServerLBVServerBinding()    {}
func (s *CSService) DeleteCSVServerLBVServerBinding() {}
func (s *CSService) GetAllCSVServerLBVServerBinding() {}
func (s *CSService) GetCSVServerLBVServerBinding()    {}
func (s *CSService) CountCSVServerLBVServerBinding()  {}

// csvserver_responderpolicy_binding
func (s *CSService) AddCSVServerResponderPolicyBinding()    {}
func (s *CSService) DeleteCSVServerResponderPolicyBinding() {}
func (s *CSService) GetAllCSVServerResponderPolicyBinding() {}
func (s *CSService) GetCSVServerResponderPolicyBinding()    {}
func (s *CSService) CountCSVServerResponderPolicyBinding()  {}

// csvserver_rewritepolicy_binding
func (s *CSService) AddCSVServerRewritePolicyBinding()    {}
func (s *CSService) DeleteCSVServerRewritePolicyBinding() {}
func (s *CSService) GetAllCSVServerRewritePolicyBinding() {}
func (s *CSService) GetCSVServerRewritePolicyBinding()    {}
func (s *CSService) CountCSVServerRewritePolicyBinding()  {}

// csvserver_spilloverpolicy_binding
func (s *CSService) AddCSVServerSpilloverPolicyBinding()    {}
func (s *CSService) DeleteCSVServerSpilloverPolicyBinding() {}
func (s *CSService) GetAllCSVServerSpilloverPolicyBinding() {}
func (s *CSService) GetCSVServerSpilloverPolicyBinding()    {}
func (s *CSService) CountCSVServerSpilloverPolicyBinding()  {}

// csvserver_tmtrafficpolicy_binding
func (s *CSService) AddCSVServerTMTrafficPolicyBinding()    {}
func (s *CSService) DeleteCSVServerTMTrafficPolicyBinding() {}
func (s *CSService) GetAllCSVServerTMTrafficPolicyBinding() {}
func (s *CSService) GetCSVServerTMTrafficPolicyBinding()    {}
func (s *CSService) CountCSVServerTMTrafficPolicyBinding()  {}

// csvserver_transformpolicy_binding
func (s *CSService) AddCSVServerTransformPolicyBinding()    {}
func (s *CSService) DeleteCSVServerTransformPolicyBinding() {}
func (s *CSService) GetAllCSVServerTransformPolicyBinding() {}
func (s *CSService) GetCSVServerTransformPolicyBinding()    {}
func (s *CSService) CountCSVServerTransformPolicyBinding()  {}

// csvserver_vpnserver_binding
func (s *CSService) AddCSVServerVPNVServerBinding()    {}
func (s *CSService) DeleteCSVServerVPNVServerBinding() {}
func (s *CSService) GetAllCSVServerVPNVServerBinding() {}
func (s *CSService) GetCSVServerVPNVServerBinding()    {}
func (s *CSService) CountCSVServerVPNVServerBinding()  {}
