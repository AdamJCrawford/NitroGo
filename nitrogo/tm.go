package nitrogo

const (
	tmFormSSOActionURL                             = "/nitro/v1/config/tmformssoaction"
	tmGlobalAuditNSLogPolicyBindingURL             = "/nitro/v1/config/tmglobal_auditnslogpolicy_binding"
	tmGlobalAuditSyslogPolicyBindingURL            = "/nitro/v1/config/tmglobal_auditsyslogpolicy_binding"
	tmGlobalBindingURL                             = "/nitro/v1/config/tmglobal_binding"
	tmGlobalTMSessionPolicyBindingURL              = "/nitro/v1/config/tmglobal_tmsessionpolicy_binding"
	tmGlobalTMTrafficPolicyBindingURL              = "/nitro/v1/config/tmglobal_tmtrafficpolicy_binding"
	tmSAMLSSOProfileURL                            = "/nitro/v1/config/tmsamlssoprofile"
	tmSessionActionURL                             = "/nitro/v1/config/tmsessionaction"
	tmSessionParameterURL                          = "/nitro/v1/config/tmsessionparameter"
	tmSessionPolicyURL                             = "/nitro/v1/config/tmsessionpolicy"
	tmSessionPolicyAAAGroupBindingURL              = "/nitro/v1/config/tmsessionpolicy_aaagroup_binding"
	tmSessionPolicyAAAUserBindingURL               = "/nitro/v1/config/tmsessionpolicy_aaauser_binding"
	tmSessionPolicyAuthenticationVServerBindingURL = "/nitro/v1/config/tmsessionpolicy_authenticationvserver_binding"
	tmSessionPolicyBindingURL                      = "/nitro/v1/config/tmsessionpolicy_binding"
	tmSessionPolicyTMGlobalBindingURL              = "/nitro/v1/config/tmsessionpolicy_tmglobal_binding"
	tmTrafficActionURL                             = "/nitro/v1/config/tmtrafficaction"
	tmTrafficPolicyURL                             = "/nitro/v1/config/tmtrafficpolicy"
	tmTrafficPolicyBindingURL                      = "/nitro/v1/config/tmtrafficpolicy_binding"
	tmTrafficPolicyCSVServerBindingURL             = "/nitro/v1/config/tmtrafficpolicy_csvserver_binding"
	tmTrafficPolicyLBVServerBindingURL             = "/nitro/v1/config/tmtrafficpolicy_lbvserver_binding"
	tmTrafficPolicyTMGlobalBindingURL              = "/nitro/v1/config/tmtrafficpolicy_tmglobal_binding"
)

// Traffic Management
// TM session/policy configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/tm/tm
type TMService struct {
	client *Client
}

// tmformssoaction
func (s *TMService) AddTMFormSSOAction()    {}
func (s *TMService) DeleteTMFormSSOAction() {}
func (s *TMService) UpdateTMFormSSOAction() {}
func (s *TMService) UnsetTMFormSSOAction()  {}
func (s *TMService) GetAllTMFormSSOAction() {}
func (s *TMService) GetTMFormSSOAction()    {}
func (s *TMService) CountTMFormSSOAction()  {}

// tmglobal_auditnslogpolicy_binding
func (s *TMService) AddTMGlobalAuditNSLogPolicyBinding()    {}
func (s *TMService) DeleteTMGlobalAuditNSLogPolicyBinding() {}
func (s *TMService) GetTMGlobalAuditNSLogPolicyBinding()    {}
func (s *TMService) CountTMGlobalAuditNSLogPolicyBinding()  {}

// tmglobal_auditsyslogpolcy_binding
func (s *TMService) AddTMGlobalAuditSyslogPolicyBinding()    {}
func (s *TMService) DeleteTMGlobalAuditSyslogPolicyBinding() {}
func (s *TMService) GetTMGlobalAuditSyslogPolicyBinding()    {}
func (s *TMService) CountTMGlobalAuditSyslogPolicyBinding()  {}

// tmglobal_binding
func (s *TMService) GetTMGlobalBinding() {}

// tmglobal_tmsessionpolicy_binding
func (s *TMService) AddTMGlobalTMSessionPolicyBinding()    {}
func (s *TMService) DeleteTMGlobalTMSessionPolicyBinding() {}
func (s *TMService) GetTMGlobalTMSessionPolicyBinding()    {}
func (s *TMService) CountTMGlobalTMSessionPolicyBinding()  {}

// tmglobal_tmtrafficpolicy_binding
func (s *TMService) AddTMGlobalTMTrafficPolicyBinding()    {}
func (s *TMService) DeleteTMGlobalTMTrafficPolicyBinding() {}
func (s *TMService) GetTMGlobalTMTrafficPolicyBinding()    {}
func (s *TMService) CountTMGlobalTMTrafficPolicyBinding()  {}

// tmsamlssoprofile
func (s *TMService) AddTMSAMLSSOProfile()    {}
func (s *TMService) DeleteTMSAMLSSOProfile() {}
func (s *TMService) UpdateTMSAMLSSOProfile() {}
func (s *TMService) UnsetTMSAMLSSOProfile()  {}
func (s *TMService) GetAllTMSAMLSSOProfile() {}
func (s *TMService) GetTMSAMLSSOProfile()    {}
func (s *TMService) CountTMSAMLSSOProfile()  {}

// tmsessionaction
func (s *TMService) AddTMSessionAction()    {}
func (s *TMService) DeleteTMSessionAction() {}
func (s *TMService) UpdateTMSessionAction() {}
func (s *TMService) UnsetTMSessionAction()  {}
func (s *TMService) GetAllTMSessionAction() {}
func (s *TMService) GetTMSessionAction()    {}
func (s *TMService) CountTMSessionAction()  {}

// tmsessionparameter
func (s *TMService) UpdateTMSessonParameter() {}
func (s *TMService) UnsetTMSessonParameter()  {}
func (s *TMService) GetAllTMSessonParameter() {}

// tmsessionpolicy
func (s *TMService) AddTMSessionPolicy()    {}
func (s *TMService) DeleteTMSessionPolicy() {}
func (s *TMService) UpdateTMSessionPolicy() {}
func (s *TMService) UnsetTMSessionPolicy()  {}
func (s *TMService) GetAllTMSessionPolicy() {}
func (s *TMService) GetTMSessionPolicy()    {}
func (s *TMService) CountTMSessionPolicy()  {}

// tmsessionpolicy_aaagroup_binding
func (s *TMService) GetAllTMSessionPolicyAAAGroupBinding() {}
func (s *TMService) GetTMSessionPolicyAAAGroupBinding()    {}
func (s *TMService) CountTMSessionPolicyAAAGroupBinding()  {}

// tmsessionpolicy_aaauser_binding
func (s *TMService) GetAllTMSessionPolicyAAAUserBinding() {}
func (s *TMService) GetTMSessionPolicyAAAUserBinding()    {}
func (s *TMService) CountTMSessionPolicyAAAUserBinding()  {}

// tmsesisonpolicy_authenticationvserver_binding
func (s *TMService) GetAllTMSessionPolicyAuthenticationVServerBinding() {}
func (s *TMService) GetTMSessionPolicyAuthenticationVServerBinding()    {}
func (s *TMService) CountTMSessionPolicyAuthenticationVServerBinding()  {}

// tmsessionpolicy_binding
func (s *TMService) GetAllTMSessionPolicyBinding() {}
func (s *TMService) GetTMSessionPolicyBinding()    {}

// tmsessionpolicy_tmglobal_binding
func (s *TMService) GetAllTMSessionPolicyTMGlobalBinding() {}
func (s *TMService) GetTMSessionPolicyTMGlobalBinding()    {}
func (s *TMService) CountTMSessionPolicyTMGlobalBinding()  {}

// tmtrafficaction
func (s *TMService) AddTMTrafficAction()    {}
func (s *TMService) DeleteTMTrafficAction() {}
func (s *TMService) UpdateTMTrafficAction() {}
func (s *TMService) UnsetTMTrafficAction()  {}
func (s *TMService) GetAllTMTrafficAction() {}
func (s *TMService) GetTMTrafficAction()    {}
func (s *TMService) CountTMTrafficAction()  {}

// tmtrafficpolicy
func (s *TMService) AddTMTrafficPolicy()    {}
func (s *TMService) DeleteTMTrafficPolicy() {}
func (s *TMService) UpdateTMTrafficPolicy() {}
func (s *TMService) UnsetTMTrafficPolicy()  {}
func (s *TMService) GetAllTMTrafficPolicy() {}
func (s *TMService) GetTMTrafficPolicy()    {}
func (s *TMService) CountTMTrafficPolicy()  {}

// tmtrafficpolicy_binding
func (s *TMService) GetAllTMTrafficPolicyBinding() {}
func (s *TMService) GetTMTrafficPolicyBinding()    {}

// tmtrafficpolicy_csvserver_binding
func (s *TMService) GetAllTMTrafficPolicyCSVServerBinding() {}
func (s *TMService) GetTMTrafficPolicyCSVServerBinding()    {}
func (s *TMService) CountTMTrafficPolicyCSVServerBinding()  {}

// tmtrafficpolicy_lbvserver_binding
func (s *TMService) GetAllTMTrafficPolicyLBVServerBinding() {}
func (s *TMService) GetTMTrafficPolicyLBVServerBinding()    {}
func (s *TMService) CountTMTrafficPolicyLBVServerBinding()  {}

// tmtrafficpolicy_tmglobal_binding
func (s *TMService) GetAllTMTrafficPolicyTMGlobalBinding() {}
func (s *TMService) GetTMTrafficPolicyTMGlobalBinding()    {}
func (s *TMService) CountTMTrafficPolicyTMGlobalBinding()  {}
