package nitrogo

const (
	aaaCertParamsURL                                 = "/nitro/v1/config/aaacertparams"
	aaaGlobalAAAPreauthenticationPolicyBindingURL    = "/nitro/v1/config/aaaglobal_aaapreauthenticationpolicy_binding"
	aaaGlobalAuthenticationNegotiateActionBindingURL = "/nitro/v1/config/aaaglobal_authenticationnegotiateaction_binding"
	aaaGlobalBindingURL                              = "/nitro/v1/config/aaaglobal_binding"
	aaaGroupURL                                      = "/nitro/v1/config/aaagroup"
	aaaGroupAAAUserBindingURL                        = "/nitro/v1/config/aaagroup_aaauser_binding"
	aaaGroupIntranetIP6BindingURL                    = "/nitro/v1/config/aaagroup_intranetip6_binding"
	aaaGroupIntranetIPBindingURL                     = "/nitro/v1/config/aaagroup_intranetip_binding"
	aaaGroupTMSessionPolicyBindingURL                = "/nitro/v1/config/aaagroup_tmsessionpolicy_binding"
	aaaGroupVPNIntranetApplicationBindingURL         = "/nitro/v1/config/aaagroup_vpnintranetapplication_binding"
	aaaGroupVPNSessionPolicyBindingURL               = "/nitro/v1/config/aaagroup_vpnsessionpolicy_binding"
	aaaGroupVPNTrafficPolicyBindingURL               = "/nitro/v1/config/aaagroup_vpntrafficpolicy_binding "
	aaaGroupVPNURLPolicyBindingURL                   = "/nitro/v1/config/aaagroup_vpnurlpolicy_binding"
	aaaGroupVPNURLBindingURL                         = "/nitro/v1/config/aaagroup_vpnurl_binding"
	aaaKCDAccountURL                                 = "/nitro/v1/config/aaakcdaccount"
	aaaLDAPParamsURL                                 = "/nitro/v1/config/aaaldapparams"
	aaaOTPParameterURL                               = "/nitro/v1/config/aaaotpparameter"
	aaaParameterURL                                  = "/nitro/v1/config/aaaparameter"
	aaapreauthenticationActionURL                    = "/nitro/v1/config/aaapreauthenticationaction"
	aaaPreauthenticationParameterURL                 = "/nitro/v1/config/aaapreauthenticationparameter"
	aaaPreauthenticationPolicyURL                    = "/nitro/v1/config/aaapreauthenticationpolicy"
	aaaPreauthenticationPolicyAAAGlobalBindingURL    = "/nitro/v1/config/aaapreauthenticationpolicy_aaaglobal_binding"
	aaaPreauthenticationPolicyBindingURL             = "/nitro/v1/config/aaapreauthenticationpolicy_binding"
	aaaPreauthenticationPolicyVPNVServerBindingURL   = "/nitro/v1/config/aaapreauthenticationpolicy_vpnvserver_binding"
	aaaRADIUSParamsURL                               = "/nitro/v1/config/aaaradiusparams"
	aaaSessionURL                                    = "/nitro/v1/config/aaasession"
	aaaSSOProfileURL                                 = "/nitro/v1/config/aaassoprofile"
	aaaTACACSParamsURL                               = "/nitro/v1/config/aaatacacsparams"
	aaaUserURL                                       = "/nitro/v1/config/aaauser"
	aaaUserAAAGroupBindingURL                        = "/nitro/v1/config/aaauser_aaagroup_binding"
	aaaUserAuditnslogPolicyBindingURL                = "/nitro/v1/config/aaauser_auditnslogpolicy_binding"
	aaaUserAuditsyslogPolicyBindingURL               = "/nitro/v1/config/aaauser_auditsyslogpolicy_binding"
	aaaUserAuthorizationPolicyBindingURL             = "/nitro/v1/config/aaauser_authorizationpolicy_binding"
	aaaUserBindingURL                                = "/nitro/v1/config/aaauser_binding"
	aaaUserIntranetIP6BindingURL                     = "/nitro/v1/config/aaauser_intranetip6_binding"
	aaaUserIntranetIPBindingURL                      = "/nitro/v1/config/aaauser_intranetip_binding"
	aaaUserTMSessionPolicyBindingURL                 = "/nitro/v1/config/aaauser_tmsessionpolicy_binding"
	aaaUserVPNIntranetApplicationBindingURL          = "/nitro/v1/config/aaauser_vpnintranetapplication_binding"
	aaaUserVPNSessionPolicyBindingURL                = "/nitro/v1/config/aaauser_vpnsessionpolicy_binding"
	aaaUserVPNTrafficPolicyBindingURL                = "/nitro/v1/config/aaauser_vpntrafficpolicy_binding"
	aaaUserVPNURLPolicyBindingURL                    = "/nitro/v1/config/aaauser_vpnurlpolicy_binding"
	aaaUserVPNURLBindingURL                          = "/nitro/v1/config/aaauser_vpnurl_binding"
)

// Authentication, authorization, and accounting service configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaa
type AAAService struct {
	client *Client
}

// aaacertparams
// Configuration for certificate parameter resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaacertparams
func (s *AAAService) UpdateAAACertParams() {}
func (s *AAAService) UnsetAAACertParams()  {}
func (s *AAAService) GetAllAAACertParams() {}

// aaaglobal_aaapreauthenticationpolicy_binding
// Binding object showing the aaapreauthenticationpolicy that can be bound to aaaglobal.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaaglobal_aaapreauthenticationpolicy_binding
func (s *AAAService) AddAAAGlobalAAAPreauthenticationPolicyBinding()    {}
func (s *AAAService) DeleteAAAGlobalAAAPreauthenticationPolicyBinding() {}
func (s *AAAService) GetAAAGlobalAAAPreauthenticationPolicyBinding()    {}
func (s *AAAService) CountAAAGlobalAAAPreauthenticationPolicyBinding()  {}

// aaaglobal_authenticationnegotiataction_binding
// Binding object showing the authenticationnegotiateaction that can be bound to aaaglobal.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaaglobal_authenticationnegotiateaction_binding
func (s *AAAService) AddAAAGlobalAuthenticationNegotiateActionBinding()    {}
func (s *AAAService) DeleteAAAGlobalAuthenticationNegotiateActionBinding() {}
func (s *AAAService) GetAAAGlobalAuthenticationNegotiateActionBinding()    {}
func (s *AAAService) CountAAAGlobalAuthenticationNegotiateActionBinding()  {}

// aaaglobal_binding
// Binding object which returns the resources bound to aaaglobal.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaaglobal_binding
func (s *AAAService) GetAAAGlobalBinding() {}

// aaagroup
// Configuration for AAA group resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaagroup
func (s *AAAService) AddAAAGroup()    {}
func (s *AAAService) DeleteAAAGroup() {}
func (s *AAAService) GetAllAAAGroup() {}
func (s *AAAService) GetAAAGroup()    {}
func (s *AAAService) CountAAAGroup()  {}

// aaagroup_aaauser_binding
// Binding object showing the aaauser that can be bound to aaagroup.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaagroup_aaauser_binding
func (s *AAAService) AddAAAGroupAAAUserBinding()    {}
func (s *AAAService) DeleteAAAGroupAAAUserBinding() {}
func (s *AAAService) GetAllAAAGroupAAAUserBinding() {}
func (s *AAAService) GetAAAGroupAAAUserBinding()    {}
func (s *AAAService) CountAAAGroupAAAUserBinding()  {}

// aaagroup_auditnslogpolicy_binding
// Binding object showing the auditnslogpolicy that can be bound to aaagroup.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaagroup_auditnslogpolicy_binding
func (s *AAAService) AddAAAGroupAuditNSLogPolicyBinding()    {}
func (s *AAAService) DeleteAAAGroupAuditNSLogPolicyBinding() {}
func (s *AAAService) GetAllAAAGroupAuditNSLogPolicyBinding() {}
func (s *AAAService) GetAAAGroupAuditNSLogPolicyBinding()    {}
func (s *AAAService) CountAAAGroupAuditNSLogPolicyBinding()  {}

// aaagroup_auditsyslogpolicy_binding
// Binding object showing the auditsyslogpolicy that can be bound to aaagroup.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaagroup_auditsyslogpolicy_binding
func (s *AAAService) AddAAAGroupAuditSyslogPolicyBinding()    {}
func (s *AAAService) DeleteAAAGroupAuditSyslogPolicyBinding() {}
func (s *AAAService) GetAllAAAGroupAuditSyslogPolicyBinding() {}
func (s *AAAService) GetAAAGroupAuditSyslogPolicyBinding()    {}
func (s *AAAService) CountAAAGroupAuditSyslogPolicyBinding()  {}

// aaagroup_authorizationpolicy_binding
// Binding object showing the authorizationpolicy that can be bound to aaagroup.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaagroup_authorizationpolicy_binding
func (s *AAAService) AddAAAGroupAuthorizationPolicyBinding()    {}
func (s *AAAService) DeleteAAAGroupAuthorizationPolicyBinding() {}
func (s *AAAService) GetAllAAAGroupAuthorizationPolicyBinding() {}
func (s *AAAService) GetAAAGroupAuthorizationPolicyBinding()    {}
func (s *AAAService) CountAAAGroupAuthorizationPolicyBinding()  {}

// aaagroup_binding
// Binding object which returns the resources bound to aaagroup.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaagroup_binding
func (s *AAAService) GetAllAAAGroupBinding() {}
func (s *AAAService) GetAAAGroupBinding()    {}

// aaagroup_intranetip6_binding
// Binding object showing the intranetip6 that can be bound to aaagroup.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaagroup_intranetip6_binding
func (s *AAAService) AddAAAGroupIntranetIP6Binding()    {}
func (s *AAAService) DeleteAAAGroupIntranetIP6Binding() {}
func (s *AAAService) GetAllAAAGroupIntranetIP6Binding() {}
func (s *AAAService) GetAAAGroupIntranetIP6Binding()    {}
func (s *AAAService) CountAAAGroupIntranetIP6Binding()  {}

// aaagroup_intranetip_binding
// Binding object showing the intranetip that can be bound to aaagroup.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaagroup_intranetip_binding
func (s *AAAService) AddAAAGroupIntranetIPBinding()    {}
func (s *AAAService) DeleteAAAGroupIntranetIPBinding() {}
func (s *AAAService) GetAllAAAGroupIntranetIPBinding() {}
func (s *AAAService) GetAAAGroupIntranetIPBinding()    {}
func (s *AAAService) CountAAAGroupIntranetIPBinding()  {}

// aaagroup_tmsessionpolicy_binding
// Binding object showing the tmsessionpolicy that can be bound to aaagroup.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaagroup_tmsessionpolicy_binding
func (s *AAAService) AddAAAGroupTMSessionPolicyBinding()    {}
func (s *AAAService) DeleteAAAGroupTMSessionPolicyBinding() {}
func (s *AAAService) GetAllAAAGroupTMSessionPolicyBinding() {}
func (s *AAAService) GetAAAGroupTMSessionPolicyBinding()    {}
func (s *AAAService) CountAAAGroupTMSessionPolicyBinding()  {}

// aaagroup_vpnintranetapplication_binding
// Binding object showing the vpnintranetapplication that can be bound to aaagroup.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaagroup_vpnintranetapplication_binding
func (s *AAAService) AddAAAGroupVPNIntranetApplicationBinding()    {}
func (s *AAAService) DeleteAAAGroupVPNIntranetApplicationBinding() {}
func (s *AAAService) GetAllAAAGroupVPNIntranetApplicationBinding() {}
func (s *AAAService) GetAAAGroupVPNIntranetApplicationBinding()    {}
func (s *AAAService) CountAAAGroupVPNIntranetApplicationBinding()  {}

// aaagroup_vpnsessionpolicy_binding
// Binding object showing the vpnsessionpolicy that can be bound to aaagroup.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaagroup_vpnsessionpolicy_binding
func (s *AAAService) AddAAAGroupVPNSessionPolicyBinding()    {}
func (s *AAAService) DeleteAAAGroupVPNSessionPolicyBinding() {}
func (s *AAAService) GetAllAAAGroupVPNSessionPolicyBinding() {}
func (s *AAAService) GetAAAGroupVPNSessionPolicyBinding()    {}
func (s *AAAService) CountAAAGroupVPNSessionPolicyBinding()  {}

// aaagroup_vpntrafficpolicy_binding
// Binding object showing the vpntrafficpolicy that can be bound to aaagroup.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaagroup_vpntrafficpolicy_binding
func (s *AAAService) AddAAAGroupVPNTrafficPolicyBinding()    {}
func (s *AAAService) DeleteAAAGroupVPNTrafficPolicyBinding() {}
func (s *AAAService) GetAllAAAGroupVPNTrafficPolicyBinding() {}
func (s *AAAService) GetAAAGroupVPNTrafficPolicyBinding()    {}
func (s *AAAService) CountAAAGroupVPNTrafficPolicyBinding()  {}

// aaagroup_vpnurlpolicy_binding
// Binding object showing the vpnurlpolicy that can be bound to aaagroup.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaagroup_vpnurlpolicy_binding
func (s *AAAService) AddAAAGroupVPNURLPolicyBinding()    {}
func (s *AAAService) DeleteAAAGroupVPNURLPolicyBinding() {}
func (s *AAAService) GetAllAAAGroupVPNURLPolicyBinding() {}
func (s *AAAService) GetAAAGroupVPNURLPolicyBinding()    {}
func (s *AAAService) CountAAAGroupVPNURLPolicyBinding()  {}

// aaagroup_vpnurl_binding
// Binding object showing the vpnurl that can be bound to aaagroup.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaagroup_vpnurl_binding
func (s *AAAService) AddAAAGroupVPNURLBinding()    {}
func (s *AAAService) DeleteAAAGroupVPNURLBinding() {}
func (s *AAAService) GetAllAAAGroupVPNURLBinding() {}
func (s *AAAService) GetAAAGroupVPNURLBinding()    {}
func (s *AAAService) CountAAAGroupVPNURLBinding()  {}

// aaakcdaccount
// Configuration for Kerberos constrained delegation account resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaakcdaccount
func (s *AAAService) AddAAAKCDAccount()    {}
func (s *AAAService) DeleteAAAKCDAccount() {}
func (s *AAAService) UpdateAAAKCDAccount() {}
func (s *AAAService) UpsetAAAKCDAccount()  {}
func (s *AAAService) GetAllAAAKCDAccount() {}
func (s *AAAService) GetAAAKCDAccount()    {}
func (s *AAAService) CountAAAKCDAccount()  {}
func (s *AAAService) CheckAAAKCDAccount()  {}

// aaaldapparams
// Configuration for LDAP parameter resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaaldapparams
func (s *AAAService) UpdateAAALDAPParams() {}
func (s *AAAService) UnsetAAALDAPParams()  {}
func (s *AAAService) GetAllAAALDAPParams() {}

// aaaotpparamter
// Configuration for AAA otpparameter resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaaotpparameter
func (s *AAAService) UpdateAAAOTPParameter() {}
func (s *AAAService) UnsetAAAOTPParameter()  {}
func (s *AAAService) GetAllAAAOTPParameter() {}

// aaaparameter
// Configuration for AAA parameter resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaaparameter
func (s *AAAService) UpdateAAAParameter() {}
func (s *AAAService) UnsetAAAParameter()  {}
func (s *AAAService) GetAllAAAParameter() {}

// aaapreauthenticationaction
// Configuration for pre authentication action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaapreauthenticationaction
func (s *AAAService) AddAAAPreauthenticationAction()    {}
func (s *AAAService) DeleteAAAPreauthenticationAction() {}
func (s *AAAService) UpdateAAAPreauthenticationAction() {}
func (s *AAAService) UnsetAAAPreauthenticationAction()  {}
func (s *AAAService) GetAllAAAPreauthenticationAction() {}
func (s *AAAService) GetAAAPreauthenticationAction()    {}
func (s *AAAService) CountAAAPreauthenticationAction()  {}

// aaapreauthenticationparameter
// Configuration for pre authentication parameter resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaapreauthenticationparameter
func (s *AAAService) UpdateAAAPreauthenticationParamter() {}
func (s *AAAService) UnsetAAAPreauthenticationParamter()  {}
func (s *AAAService) GetAllAAAPreauthenticationParamter() {}

// aaapreauthenticationpolicy
// Configuration for pre authentication policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaapreauthenticationpolicy
func (s *AAAService) AddAAAPreauthenticationPolicy()    {}
func (s *AAAService) DeleteAAAPreauthenticationPolicy() {}
func (s *AAAService) UpdateAAAPreauthenticationPolicy() {}
func (s *AAAService) GetAllAAAPreauthenticationPolicy() {}
func (s *AAAService) GetAAAPreauthenticationPolicy()    {}
func (s *AAAService) CountAAAPreauthenticationPolicy()  {}

// aaapreauthenticationpolicy_aaaglobal_binding
// Binding object showing the aaaglobal that can be bound to aaapreauthenticationpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaapreauthenticationpolicy_aaaglobal_binding
func (s *AAAService) GetAllAAAPreauthenticationPolicyAAAGlobalBinding() {}
func (s *AAAService) GetAAAPreauthenticationPolicyAAAGlobalBinding()    {}
func (s *AAAService) CountAAAPreauthenticationPolicyAAAGlobalBinding()  {}

// aaapreauthenticationpolicy_binding
// Binding object which returns the resources bound to aaapreauthenticationpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaapreauthenticationpolicy_binding
func (s *AAAService) GetAllAAAPreauthenticationPolicyBinding() {}
func (s *AAAService) GetAAAPreauthenticationPolicyBinding()    {}

// aaapreauthentiucationpolicy_vpnvserver_binding
// Binding object showing the vpnvserver that can be bound to aaapreauthenticationpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaapreauthenticationpolicy_vpnvserver_binding
func (s *AAAService) GetAllAAAPreauthenticationPolicyVPNVServerBinding() {}
func (s *AAAService) GetAAAPreauthenticationPolicyVPNVServerBinding()    {}
func (s *AAAService) CountAAAPreauthenticationPolicyVPNVServerBinding()  {}

// aaaradiusparams
// Configuration for RADIUS parameter resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaaradiusparams
func (s *AAAService) UpdateAAARADIUSParams() {}
func (s *AAAService) UnsetAAARADIUSParams()  {}
func (s *AAAService) GetAllAAARADIUSParams() {}

// aaasession
// Configuration for active connection resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaasession
func (s *AAAService) GetAllAAASession() {}
func (s *AAAService) CountAAASession()  {}
func (s *AAAService) KillAAASession()   {}

// aaassoprofile
// Configuration for aaa sso profile resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaassoprofile
func (s *AAAService) AddAAASSOProfile()    {}
func (s *AAAService) DeleteAAASSOProfile() {}
func (s *AAAService) GetAllAAASSOProfile() {}
func (s *AAAService) GetAAASSOProfile()    {}
func (s *AAAService) CountAAASSOProfile()  {}
func (s *AAAService) UpdateAAASSOProfile() {}

// aaatacasparams
// Configuration for tacacs parameters resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaatacacsparams
func (s *AAAService) UpdateAAATACACSParams() {}
func (s *AAAService) UnsetAAATACACSParams()  {}
func (s *AAAService) GetAllAAATACACSParams() {}

// aaauser
// Configuration for AAA user resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaauser
func (s *AAAService) AddAAAUser()    {}
func (s *AAAService) DeleteAAAUser() {}
func (s *AAAService) UpdateAAAUser() {}
func (s *AAAService) GetAllAAAUser() {}
func (s *AAAService) GetAAAUser()    {}
func (s *AAAService) CountAAAUser()  {}
func (s *AAAService) UnlockAAAUser() {}

// aaauser_aaagroup_binding
// Binding object showing the aaagroup that can be bound to aaauser.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaauser_aaagroup_binding
func (s *AAAService) GetAllAAAUserAAAGroupBinding() {}
func (s *AAAService) GetAAAUserAAAGroupBinding()    {}
func (s *AAAService) CountAAAUserAAAGroupBinding()  {}

// aaauser_auditnslogpolicy_binding
// Binding object showing the auditnslogpolicy that can be bound to aaauser.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaauser_auditnslogpolicy_binding
func (s *AAAService) AddAAAUserAuditNSLogPolicyBinding()    {}
func (s *AAAService) DeleteAAAUserAuditNSLogPolicyBinding() {}
func (s *AAAService) GetAllAAAUserAuditNSLogPolicyBinding() {}
func (s *AAAService) GetAAAUserAuditNSLogPolicyBinding()    {}
func (s *AAAService) CountAAAUserAuditNSLogPolicyBinding()  {}

// aaauser_auditsyslogpolicy_binding
// Binding object showing the auditsyslogpolicy that can be bound to aaauser.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaauser_auditsyslogpolicy_binding
func (s *AAAService) AddAAAUserAuditSyslogPolicyBinding()    {}
func (s *AAAService) DeleteAAAUserAuditSyslogPolicyBinding() {}
func (s *AAAService) GetAllAAAUserAuditSyslogPolicyBinding() {}
func (s *AAAService) GetAAAUserAuditSyslogPolicyBinding()    {}
func (s *AAAService) CountAAAUserAuditSyslogPolicyBinding()  {}

// aaauser_authorizationpolicy_binding
// Binding object showing the authorizationpolicy that can be bound to aaauser.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaauser_authorizationpolicy_binding
func (s *AAAService) AddAAAUserAuthorizationPolicyBinding()    {}
func (s *AAAService) DeleteAAAUserAuthorizationPolicyBinding() {}
func (s *AAAService) GetAllAAAUserAuthorizationPolicyBinding() {}
func (s *AAAService) GetAAAUserAuthorizationPolicyBinding()    {}
func (s *AAAService) CountAAAUserAuthorizationPolicyBinding()  {}

// aaauser_binding
// Binding object which returns the resources bound to aaauser.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaauser_binding
func (s *AAAService) GetAllAAAUserBinding() {}
func (s *AAAService) GetAAAUserBinding()    {}

// aaauser_intranetip6_binding
// Binding object showing the intranetip6 that can be bound to aaauser.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaauser_intranetip6_binding
func (s *AAAService) AddAAAUserIntranetIP6Binding()    {}
func (s *AAAService) DeleteAAAUserIntranetIP6Binding() {}
func (s *AAAService) GetAllAAAUserIntranetIP6Binding() {}
func (s *AAAService) GetAAAUserIntranetIP6Binding()    {}
func (s *AAAService) CountAAAUserIntranetIP6Binding()  {}

// aaauser_intranetip_binding
// Binding object showing the intranetip that can be bound to aaauser.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaauser_intranetip_binding
func (s *AAAService) AddAAAUserIntranetIPBinding()    {}
func (s *AAAService) DeleteAAAUserIntranetIPBinding() {}
func (s *AAAService) GetAllAAAUserIntranetIPBinding() {}
func (s *AAAService) GetAAAUserIntranetIPBinding()    {}
func (s *AAAService) CountAAAUserIntranetIPBinding()  {}

// aaauser_tmsessionpolicy_binding
// Binding object showing the tmsessionpolicy that can be bound to aaauser.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaauser_tmsessionpolicy_binding
func (s *AAAService) AddAAAUserTMSessionPolicyBinding()    {}
func (s *AAAService) DeleteAAAUserTMSessionPolicyBinding() {}
func (s *AAAService) GetAllAAAUserTMSessionPolicyBinding() {}
func (s *AAAService) GetAAAUserTMSessionPolicyBinding()    {}
func (s *AAAService) CountAAAUserTMSessionPolicyBinding()  {}

// aaauser_vpnintranetapplication_binding
// Binding object showing the vpnintranetapplication that can be bound to aaauser.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaauser_vpnintranetapplication_binding
func (s *AAAService) AddAAAUserVPNIntranetApplicationBinding()    {}
func (s *AAAService) DeleteAAAUserVPNIntranetApplicationBinding() {}
func (s *AAAService) GetAllAAAUserVPNIntranetApplicationBinding() {}
func (s *AAAService) GetAAAUserVPNIntranetApplicationBinding()    {}
func (s *AAAService) CountAAAUserVPNIntranetApplicationBinding()  {}

// aaauser_vpnsessionpolicy_binding
// Binding object showing the vpnsessionpolicy that can be bound to aaauser.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaauser_vpnsessionpolicy_binding
func (s *AAAService) AddAAAUserVPNSessionPolicyBinding()    {}
func (s *AAAService) DeleteAAAUserVPNSessionPolicyBinding() {}
func (s *AAAService) GetAllAAAUserVPNSessionPolicyBinding() {}
func (s *AAAService) GetAAAUserVPNSessionPolicyBinding()    {}
func (s *AAAService) CountAAAUserVPNSessionPolicyBinding()  {}

// aaauser_vpntrafficpolicy_binding
// Binding object showing the vpntrafficpolicy that can be bound to aaauser.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaauser_vpntrafficpolicy_binding
func (s *AAAService) AddAAAUserVPNTrafficPolicyBinding()    {}
func (s *AAAService) DeleteAAAUserVPNTrafficPolicyBinding() {}
func (s *AAAService) GetAllAAAUserVPNTrafficPolicyBinding() {}
func (s *AAAService) GetAAAUserVPNTrafficPolicyBinding()    {}
func (s *AAAService) CountAAAUserVPNTrafficPolicyBinding()  {}

// aaauser_vpnurlpolicy_binding
// Binding object showing the vpnurlpolicy that can be bound to aaauser.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaauser_vpnurlpolicy_binding
func (s *AAAService) AddAAAUserVPNURLPolicyBinding()    {}
func (s *AAAService) DeleteAAAUserVPNURLPolicyBinding() {}
func (s *AAAService) GetAllAAAUserVPNURLPolicyBinding() {}
func (s *AAAService) GetAAAUserVPNURLPolicyBinding()    {}
func (s *AAAService) CountAAAUserVPNURLPolicyBinding()  {}

// aaauser_vpnurl_binding
// Binding object showing the vpnurl that can be bound to aaauser.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaauser_vpnurl_binding
func (s *AAAService) AddAAAUserVPNURLBinding()    {}
func (s *AAAService) DeleteAAAUserVPNURLBinding() {}
func (s *AAAService) GetAllAAAUserVPNURLBinding() {}
func (s *AAAService) GetAAAUserVPNURLBinding()    {}
func (s *AAAService) CountAAAUserVPNURLBinding()  {}
