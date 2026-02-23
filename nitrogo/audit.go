package nitrogo

const (
	auditMessageActionURL                            = "/nitro/v1/config/auditmessageaction"
	auditMessagesURL                                 = "/nitro/v1/config/auditmessages"
	auditNSLogActionURL                              = "/nitro/v1/config/auditnslogaction"
	auditNSLogGlobalAuditNSLogPolicyBindingURL       = "/nitro/v1/config/auditnslogglobal_auditnslogpolicy_binding"
	auditNSLogGlobalBindingURL                       = "/nitro/v1/config/auditnslogglobal_binding"
	auditNSLogParamsURL                              = "/nitro/v1/config/auditnslogparams"
	auditNSLogPolicyURL                              = "/nitro/v1/config/auditnslogpolicy"
	auditNSLogPolicyAAAGroupBindingURL               = "/nitro/v1/config/auditnslogpolicy_aaagroup_binding"
	auditNSLogPolicyAAAUserBindingURL                = "/nitro/v1/config/auditnslogpolicy_aaauser_binding"
	auditNSLogPolicyAppFWGlobalBindingURL            = "/nitro/v1/config/auditnslogpolicy_appfwglobal_binding"
	auditNSLogPolicyAuditNSLogGlobalBindingURL       = "/nitro/v1/config/auditnslogpolicy_auditnslogglobal_binding"
	auditNSLogPolicyAuthenticationVServerBindingURL  = "/nitro/v1/config/auditnslogpolicy_authenticationvserver_binding"
	auditNSLogPolicyBindingURL                       = "/nitro/v1/config/auditnslogpolicy_binding"
	auditNSLogPolicyCSVServerBindingURL              = "/nitro/v1/config/auditnslogpolicy_csvserver_binding"
	auditNSLogPolicyLBVServerBindingURL              = "/nitro/v1/config/auditnslogpolicy_lbvserver_binding"
	auditNSLogPolicySystemGlobalBindingURL           = "/nitro/v1/config/auditnslogpolicy_systemglobal_binding"
	auditNSLogPolicyTMGlobalBindingURL               = "/nitro/v1/config/auditnslogpolicy_tmglobal_binding"
	auditNSLogPolicyVPNGlobalBindingURL              = "/nitro/v1/config/auditnslogpolicy_vpnglobal_binding"
	auditNSLogPolicyVPNVServerBindingURL             = "/nitro/v1/config/auditnslogpolicy_vpnvserver_binding"
	auditSyslogActionURL                             = "/nitro/v1/config/auditsyslogaction"
	auditSyslogGlobalAuditSyslogPolicyBindingURL     = "/nitro/v1/config/auditsyslogglobal_auditsyslogpolicy_binding"
	auditSyslogGlobalBindingURL                      = "/nitro/v1/config/auditsyslogglobal_binding"
	auditSyslogParamsURL                             = "/nitro/v1/config/auditsyslogparams"
	auditSyslogPolicyURL                             = "/nitro/v1/config/auditsyslogpolicy"
	auditSyslogPolicyAAAGroupBindingURL              = "/nitro/v1/config/auditsyslogpolicy_aaagroup_binding"
	auditSyslogPolicyAAAUserBindingURL               = "/nitro/v1/config/auditsyslogpolicy_aaauser_binding"
	auditSyslogPolicyAuditSyslogGlobalBindingURL     = "/nitro/v1/config/auditsyslogpolicy_auditsyslogglobal_binding"
	auditSyslogPolicyAuthenticationVServerBindingURL = "/nitro/v1/config/auditsyslogpolicy_authenticationvserver_binding"
	auditSyslogPolicyBindingURL                      = "/nitro/v1/config/auditsyslogpolicy_binding"
	auditSyslogPolicyCSVServerBindingURL             = "/nitro/v1/config/auditsyslogpolicy_csvserver_binding"
	auditSyslogPolicyLBVServerBindingURL             = "/nitro/v1/config/auditsyslogpolicy_lbvserver_binding"
	auditSyslogPolicyRNATGlobalBindingURL            = "/nitro/v1/config/auditsyslogpolicy_rnatglobal_binding"
	auditSyslogPolicySystemGlobalBindingURL          = "/nitro/v1/config/auditsyslogpolicy_systemglobal_binding"
	auditSyslogPolicyTMGlobalBindingURL              = "/nitro/v1/config/auditsyslogpolicy_tmglobal_binding"
	auditSyslogPolicyVPNGlobalBindingURL             = "/nitro/v1/config/auditsyslogpolicy_vpnglobal_binding"
	auditSyslogPolicyVPNVServerBindingURL            = "/nitro/v1/config/auditsyslogpolicy_vpnvserver_binding"
)

// Audit configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/audit
type AuditService struct {
	client *Client
}

// auditmessageaction
// Configuration for message action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditmessageaction
func (s *AuditService) AddAuditMessageAction()    {}
func (s *AuditService) DeleteAuditMessageAction() {}
func (s *AuditService) UpdateAuditMessageAction() {}
func (s *AuditService) UnsetAuditMessageAction()  {}
func (s *AuditService) GetAllAuditMessageAction() {}
func (s *AuditService) GetAuditMessageAction()    {}
func (s *AuditService) CountAuditMessageAction()  {}

// auditmessages
// Configuration for audit message resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditmessages
func (s *AuditService) GetAllAuditMessages() {}
func (s *AuditService) CountAuditMessages()  {}

// auditnslogaction
// Configuration for ns log action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditnslogaction
func (s *AuditService) AddAuditNSLogAction()    {}
func (s *AuditService) DeleteAuditNSLogAction() {}
func (s *AuditService) UpdateAuditNSLogAction() {}
func (s *AuditService) UnsetAuditNSLogAction()  {}
func (s *AuditService) GetAllAuditNSLogAction() {}
func (s *AuditService) GetAuditNSLogAction()    {}
func (s *AuditService) CountAuditNSLogAction()  {}

// auditnslogglobal_auditnslogpolicy_binding
// Binding object showing the auditnslogpolicy that can be bound to auditnslogglobal.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditnslogglobal_auditnslogpolicy_binding
func (s *AuditService) AddAuditNSLogGlobalAuditNSLogPolicyBinding()    {}
func (s *AuditService) DeleteAuditNSLogGlobalAuditNSLogPolicyBinding() {}
func (s *AuditService) GetAuditNSLogGlobalAuditNSLogPolicyBinding()    {}
func (s *AuditService) CountAuditNSLogGlobalAuditNSLogPolicyBinding()  {}

// auditnslogglobal_binding
// Binding object which returns the resources bound to auditnslogglobal.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditnslogglobal_binding
func (s *AuditService) GetAuditNSLogGlobalBinding() {}

// auditnslogparams
// Configuration for ns log parameters resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditnslogparams
func (s *AuditService) UpdateAuditNSLogParams() {}
func (s *AuditService) UnsetAuditNSLogParams()  {}
func (s *AuditService) GetAllAuditNSLogParams() {}

// auditnslogpolicy
// Configuration for ns log policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditnslogpolicy
func (s *AuditService) AddAuditNSLogPolicy()    {}
func (s *AuditService) DeleteAuditNSLogPolicy() {}
func (s *AuditService) UpdateAuditNSLogPolicy() {}
func (s *AuditService) GetAllAuditNSLogPolicy() {}
func (s *AuditService) GetAuditNSLogPolicy()    {}
func (s *AuditService) CountAuditNSLogPolicy()  {}

// auditnslogpolicy_aaagroup_binding
// Binding object showing the aaagroup that can be bound to auditnslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditnslogpolicy_aaagroup_binding
func (s *AuditService) GetAllAuditNSLogPolicyAAAGroupBinding() {}
func (s *AuditService) GetAuditNSLogPolicyAAAGroupBinding()    {}
func (s *AuditService) CountAuditNSLogPolicyAAAGroupBinding()  {}

// auditnslogpolicy_aaauser_binding
// Binding object showing the aaauser that can be bound to auditnslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditnslogpolicy_aaauser_binding
func (s *AuditService) GetAllAuditNSLogPolicyAAAUserBinding() {}
func (s *AuditService) GetAuditNSLogPolicyAAAUserBinding()    {}
func (s *AuditService) CountAuditNSLogPolicyAAAUserBinding()  {}

// auditnslogpolicy_appfwglobal_binding
// Binding object showing the appfwglobal that can be bound to auditnslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditnslogpolicy_appfwglobal_binding
func (s *AuditService) GetAllAuditNSLogPolicyAppFWGlobalBinding() {}
func (s *AuditService) GetAuditNSLogPolicyAppFWGlobalBinding()    {}
func (s *AuditService) CountAuditNSLogPolicyAppFWGlobalBinding()  {}

// auditnslogpolicy_auditnslogglobal_binding
// Binding object showing the auditnslogglobal that can be bound to auditnslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditnslogpolicy_auditnslogglobal_binding
func (s *AuditService) GetAllAuditNSLogPolicyAuditNSLogGlobalBinding() {}
func (s *AuditService) GetAuditNSLogPolicyAuditNSLogGlobalBinding()    {}
func (s *AuditService) CountAuditNSLogPolicyAuditNSLogGlobalBinding()  {}

// auditnslogpolicy_authenticationvserver_binding
// Binding object showing the authenticationvserver that can be bound to auditnslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditnslogpolicy_authenticationvserver_binding
func (s *AuditService) GetAllAuditNSLogPolicyAuthenticationVServerBinding() {}
func (s *AuditService) GetAuditNSLogPolicyAuthenticationVServerBinding()    {}
func (s *AuditService) CountAuditNSLogPolicyAuthenticationVServerBinding()  {}

// auditnslogpolicy_binding
// Binding object which returns the resources bound to auditnslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditnslogpolicy_binding
func (s *AuditService) GetAllAuditNSLogPolicyBinding() {}
func (s *AuditService) GetAuditNSLogPolicyBinding()    {}

// auditnslogpolicy_csvserver_binding
// Binding object showing the csvserver that can be bound to auditnslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditnslogpolicy_csvserver_binding
func (s *AuditService) GetAllAuditNSLogPolicyCSVServerBinding() {}
func (s *AuditService) GetAuditNSLogPolicyCSVServerBinding()    {}
func (s *AuditService) CountAuditNSLogPolicyCSVServerBinding()  {}

// auditnslogpolicy_lbvserver_binding
// Binding object showing the lbvserver that can be bound to auditnslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditnslogpolicy_lbvserver_binding
func (s *AuditService) GetAllAuditNSLogPolicyLBVServerBinding() {}
func (s *AuditService) GetAuditNSLogPolicyLBVServerBinding()    {}
func (s *AuditService) CountAuditNSLogPolicyLBVServerBinding()  {}

// auditnslogpolicy_systemglobal_binding
// Binding object showing the systemglobal that can be bound to auditnslogpolicy
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditnslogpolicy_systemglobal_binding
func (s *AuditService) GetAllAuditNSLogPolicySystemGlobalBinding() {}
func (s *AuditService) GetAuditNSLogPolicySystemGlobalBinding()    {}
func (s *AuditService) CountAuditNSLogPolicySystemGlobalBinding()  {}

// auditnslogpolicy_tmglobal_binding
// Binding object showing the tmglobal that can be bound to auditnslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditnslogpolicy_tmglobal_binding
func (s *AuditService) GetAllAuditNSLogPolicyTMGlobalBinding() {}
func (s *AuditService) GetAuditNSLogPolicyTMGlobalBinding()    {}
func (s *AuditService) CountAuditNSLogPolicyTMGlobalBinding()  {}

// auditnslogpolicy_vpnglobal_binding
// Binding object showing the vpnglobal that can be bound to auditnslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditnslogpolicy_vpnglobal_binding
func (s *AuditService) GetAllAuditNSLogPolicyVPNGlobalBinding() {}
func (s *AuditService) GetAuditNSLogPolicyVPNGlobalBinding()    {}
func (s *AuditService) CountAuditNSLogPolicyVPNGlobalBinding()  {}

// auditnslogpolicy_vpnvserver_binding
// Binding object showing the vpnvserver that can be bound to auditnslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditnslogpolicy_vpnvserver_binding
func (s *AuditService) GetAllAuditNSLogPolicyVPNVServerBinding() {}
func (s *AuditService) GetAuditNSLogPolicyVPNVServerBinding()    {}
func (s *AuditService) CountAuditNSLogPolicyVPNVServerBinding()  {}

// auditsyslogaction
// Configuration for system log action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditsyslogaction
func (s *AuditService) AddAuditSyslogAction()    {}
func (s *AuditService) DeleteAuditSyslogAction() {}
func (s *AuditService) UpdateAuditSyslogAction() {}
func (s *AuditService) UnsetAuditSyslogAction()  {}
func (s *AuditService) GetAllAuditSyslogAction() {}
func (s *AuditService) GetAuditSyslogAction()    {}
func (s *AuditService) CountAuditSyslogAction()  {}

// auditsyslogglobal_auditsyslogpolicy_binding
// Binding object showing the auditsyslogpolicy that can be bound to auditsyslogglobal.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditsyslogglobal_auditsyslogpolicy_binding
func (s *AuditService) AddAuditSyslogGlobalAuditSyslogPolicyBinding()    {}
func (s *AuditService) DeleteAuditSyslogGlobalAuditSyslogPolicyBinding() {}
func (s *AuditService) GetAuditSyslogGlobalAuditSyslogPolicyBinding()    {}
func (s *AuditService) CountAuditSyslogGlobalAuditSyslogPolicyBinding()  {}

// auditsyslogglobal_binding
// Binding object which returns the resources bound to auditsyslogglobal.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditsyslogglobal_binding
func (s *AuditService) GetAuditSyslogGlobalBinding() {}

// auditsyslogparams
// Configuration for system log parameters resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditsyslogparams
func (s *AuditService) UpdateAuditSyslogParams() {}
func (s *AuditService) UnsetAuditSyslogParams()  {}
func (s *AuditService) GetAllAuditSyslogParams() {}

// auditsyslogpolicy
// Configuration for system log policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditsyslogpolicy
func (s *AuditService) AddAuditSyslogPolicy()    {}
func (s *AuditService) DeleteAuditSyslogPolicy() {}
func (s *AuditService) UpdateAuditSyslogPolicy() {}
func (s *AuditService) GetAllAuditSyslogPolicy() {}
func (s *AuditService) GetAuditSyslogPolicy()    {}
func (s *AuditService) CountAuditSyslogPolicy()  {}

// auditsyslogpolicy_aaagroup_binding
// Binding object showing the aaagroup that can be bound to auditsyslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditsyslogpolicy_aaagroup_binding
func (s *AuditService) GetAllAuditSyslogPolicyAAAGroupBinding() {}
func (s *AuditService) GetAuditSyslogPolicyAAAGroupBinding()    {}
func (s *AuditService) CountAuditSyslogPolicyAAAGroupBinding()  {}

// auditsyslogpolicy_aaauser_binding
// Binding object showing the aaauser that can be bound to auditsyslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditsyslogpolicy_aaauser_binding
func (s *AuditService) GetAllAuditSyslogPolicyAAAUserBinding() {}
func (s *AuditService) GetAuditSyslogPolicyAAAUserBinding()    {}
func (s *AuditService) CountAuditSyslogPolicyAAAUserBinding()  {}

// auditsyslogpolicy_auditsyslogglobal_binding
// Binding object showing the auditsyslogglobal that can be bound to auditsyslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditsyslogpolicy_auditsyslogglobal_binding
func (s *AuditService) GetAllAuditSyslogPolicyAuditSyslogGlobalBinding() {}
func (s *AuditService) GetAuditSyslogPolicyAuditSyslogGlobalBinding()    {}
func (s *AuditService) CountAuditSyslogPolicyAuditSyslogGlobalBinding()  {}

// auditsyslogpolicy_authenticationvserver_binding
// Binding object showing the authenticationvserver that can be bound to auditsyslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditsyslogpolicy_authenticationvserver_binding
func (s *AuditService) GetAllAuditSyslogPolicyAuthenticationVServerBinding() {}
func (s *AuditService) GetAuditSyslogPolicyAuthenticationVServerBinding()    {}
func (s *AuditService) CountAuditSyslogPolicyAuthenticationVServerBinding()  {}

// auditsyslogpolicy_binding
// Binding object which returns the resources bound to auditsyslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditsyslogpolicy_binding
func (s *AuditService) GetAllAuditSyslogPolicyBinding() {}
func (s *AuditService) GetAuditSyslogPolicyBinding()    {}

// auditsyslogpolicy_csvserver_binding
// Binding object showing the csvserver that can be bound to auditsyslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditsyslogpolicy_csvserver_binding
func (s *AuditService) GetAllAuditSyslogPolicyCSVServerBinding() {}
func (s *AuditService) GetAuditSyslogPolicyCSVServerBinding()    {}
func (s *AuditService) CountAuditSyslogPolicyCSVServerBinding()  {}

// auditsyslogpolicy_lbvserver_binding
// Binding object showing the lbvserver that can be bound to auditsyslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditsyslogpolicy_lbvserver_binding
func (s *AuditService) GetAllAuditSyslogPolicyLBVServerBinding() {}
func (s *AuditService) GetAuditSyslogPolicyLBVServerBinding()    {}
func (s *AuditService) CountAuditSyslogPolicyLBVServerBinding()  {}

// auditsyslogpolicy_rnatglobal_binding
// Binding object showing the rnatglobal that can be bound to auditsyslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditsyslogpolicy_rnatglobal_binding
func (s *AuditService) GetAllAuditSyslogPolicyRNATGlobalBinding() {}
func (s *AuditService) GetAuditSyslogPolicyRNATGlobalBinding()    {}
func (s *AuditService) CountAuditSyslogPolicyRNATGlobalBinding()  {}

// auditsyslogpolicy_systemglobal_binding
// Binding object showing the systemglobal that can be bound to auditsyslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditsyslogpolicy_systemglobal_binding
func (s *AuditService) GetAllAuditSyslogPolicySystemGlobalBinding() {}
func (s *AuditService) GetAuditSyslogPolicySystemGlobalBinding()    {}
func (s *AuditService) CountAuditSyslogPolicySystemGlobalBinding()  {}

// auditsyslogpolicy_tmpglobal_binding
// Binding object showing the tmglobal that can be bound to auditsyslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditsyslogpolicy_tmglobal_binding
func (s *AuditService) GetAllAuditSyslogPolicyTMGlobalBinding() {}
func (s *AuditService) GetAuditSyslogPolicyTMGlobalBinding()    {}
func (s *AuditService) CountAuditSyslogPolicyTMGlobalBinding()  {}

// auditsyslogpolicy_vpnglobal_binding
// Binding object showing the vpnglobal that can be bound to auditsyslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditsyslogpolicy_vpnglobal_binding
func (s *AuditService) GetAllAuditSyslogPolicyVPNGlobalBinding() {}
func (s *AuditService) GetAuditSyslogPolicyVPNGlobalBinding()    {}
func (s *AuditService) CountAuditSyslogPolicyVPNGlobalBinding()  {}

// auditsyslogpolicy_vpnvserver_binding
// Binding object showing the vpnvserver that can be bound to auditsyslogpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/audit/auditsyslogpolicy_vpnvserver_binding
func (s *AuditService) GetAllAuditSyslogPolicyVPNVServerBinding() {}
func (s *AuditService) GetAuditSyslogPolicyVPNVServerBinding()    {}
func (s *AuditService) CountAuditSyslogPolicyVPNVServerBinding()  {}
