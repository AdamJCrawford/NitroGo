package nitrogo

type AAAService struct {
	client *Client
}

// aaacertparams
func (s *AAAService) UpdateAAACertParams() {}
func (s *AAAService) UnsetAAACertParams()  {}
func (s *AAAService) GetAllAAACertParams() {}

// aaaglobal_aaapreauthenticationpolicy_binding
func (s *AAAService) AddAAAGlobaAAAPreauthenticationPolicyBinding()    {}
func (s *AAAService) DeleteAAAGlobaAAAPreauthenticationPolicyBinding() {}
func (s *AAAService) GetAAAGlobaAAAPreauthenticationPolicyBinding()    {}
func (s *AAAService) CountAAAGlobaAAAPreauthenticationPolicyBinding()  {}

// aaaglobal_authenticationnegotiataction_binding
func (s *AAAService) AddAAAGlobalAuthenticationNegotiateActionBinding()    {}
func (s *AAAService) DeleteAAAGlobalAuthenticationNegotiateActionBinding() {}
func (s *AAAService) GetAAAGlobalAuthenticationNegotiateActionBinding()    {}
func (s *AAAService) CountAAAGlobalAuthenticationNegotiateActionBinding()  {}

// aaaglobal_binding
func (s *AAAService) GetAAAGlobalBinding() {}

// aaagroup
func (s *AAAService) AddAAAGroup()    {}
func (s *AAAService) DeleteAAAGroup() {}
func (s *AAAService) GetAllAAAGroup() {}
func (s *AAAService) GetAAAGroup()    {}
func (s *AAAService) CountAAAGroup()  {}

// aaagroup_aaauser_binding
func (s *AAAService) AddAAAGroupAAAUserBinding()    {}
func (s *AAAService) DeleteAAAGroupAAAUserBinding() {}
func (s *AAAService) GetAllAAAGroupAAAUserBinding() {}
func (s *AAAService) GetAAAGroupAAAUserBinding()    {}
func (s *AAAService) CountAAAGroupAAAUserBinding()  {}

// aaagroup_auditnslogpolicy_binding
func (s *AAAService) AddAAAGroupAuditNSLogPolicyBinding()    {}
func (s *AAAService) DeleteAAAGroupAuditNSLogPolicyBinding() {}
func (s *AAAService) GetAllAAAGroupAuditNSLogPolicyBinding() {}
func (s *AAAService) GetAAAGroupAuditNSLogPolicyBinding()    {}
func (s *AAAService) CountAAAGroupAuditNSLogPolicyBinding()  {}

// aaagroup_auditsyslogpolicy_binding
func (s *AAAService) AddAAAGroupAuditSyslogPolicyBinding()    {}
func (s *AAAService) DeleteAAAGroupAuditSyslogPolicyBinding() {}
func (s *AAAService) GetAllAAAGroupAuditSyslogPolicyBinding() {}
func (s *AAAService) GetAAAGroupAuditSyslogPolicyBinding()    {}
func (s *AAAService) CountAAAGroupAuditSyslogPolicyBinding()  {}

// aaagroup_authorizationpolicy_binding
func (s *AAAService) AddAAAGroupAuthorizationPolicyBinding()    {}
func (s *AAAService) DeleteAAAGroupAuthorizationPolicyBinding() {}
func (s *AAAService) GetAllAAAGroupAuthorizationPolicyBinding() {}
func (s *AAAService) GetAAAGroupAuthorizationPolicyBinding()    {}
func (s *AAAService) CountAAAGroupAuthorizationPolicyBinding()  {}

// aaagroup_binding
func (s *AAAService) GetAllAAAGroupBinding() {}
func (s *AAAService) GetAAAGroupBinding()    {}

// aaagroup_intranetip6_binding
func (s *AAAService) AddAAAGroupIntranetIP6Binding()    {}
func (s *AAAService) DeleteAAAGroupIntranetIP6Binding() {}
func (s *AAAService) GetAllAAAGroupIntranetIP6Binding() {}
func (s *AAAService) GetAAAGroupIntranetIP6Binding()    {}
func (s *AAAService) CountAAAGroupIntranetIP6Binding()  {}

// aaagroup_intranetip_binding
func (s *AAAService) AddAAAGroupIntranetIPBinding()    {}
func (s *AAAService) DeleteAAAGroupIntranetIPBinding() {}
func (s *AAAService) GetAllAAAGroupIntranetIPBinding() {}
func (s *AAAService) GetAAAGroupIntranetIPBinding()    {}
func (s *AAAService) CountAAAGroupIntranetIPBinding()  {}

// aaagroup_tmsessionpolicy_binding
func (s *AAAService) AddAAAGroupTMSessionPolicyBinding()    {}
func (s *AAAService) DeleteAAAGroupTMSessionPolicyBinding() {}
func (s *AAAService) GetAllAAAGroupTMSessionPolicyBinding() {}
func (s *AAAService) GetAAAGroupTMSessionPolicyBinding()    {}
func (s *AAAService) CountAAAGroupTMSessionPolicyBinding()  {}

// aaagroup_vpnintranetapplication_binding
func (s *AAAService) AddAAAGroupVPNIntranetApplicationBinding()    {}
func (s *AAAService) DeleteAAAGroupVPNIntranetApplicationBinding() {}
func (s *AAAService) GetAllAAAGroupVPNIntranetApplicationBinding() {}
func (s *AAAService) GetAAAGroupVPNIntranetApplicationBinding()    {}
func (s *AAAService) CountAAAGroupVPNIntranetApplicationBinding()  {}

// aaagroup_vpnsessionpolicy_binding
func (s *AAAService) AddAAAGroupVPNSessionPolicyBinding()    {}
func (s *AAAService) DeleteAAAGroupVPNSessionPolicyBinding() {}
func (s *AAAService) GetAllAAAGroupVPNSessionPolicyBinding() {}
func (s *AAAService) GetAAAGroupVPNSessionPolicyBinding()    {}
func (s *AAAService) CountAAAGroupVPNSessionPolicyBinding()  {}

// aaagroup_vpntrafficpolicy_binding
func (s *AAAService) AddAAAGroupVPNTrafficPolicyBinding()    {}
func (s *AAAService) DeleteAAAGroupVPNTrafficPolicyBinding() {}
func (s *AAAService) GetAllAAAGroupVPNTrafficPolicyBinding() {}
func (s *AAAService) GetAAAGroupVPNTrafficPolicyBinding()    {}
func (s *AAAService) CountAAAGroupVPNTrafficPolicyBinding()  {}

// aaagroup_vpnurlpolicy_binding
func (s *AAAService) AddAAAGroupVPNURLPolicyBinding()    {}
func (s *AAAService) DeleteAAAGroupVPNURLPolicyBinding() {}
func (s *AAAService) GetAllAAAGroupVPNURLPolicyBinding() {}
func (s *AAAService) GetAAAGroupVPNURLPolicyBinding()    {}
func (s *AAAService) CountAAAGroupVPNURLPolicyBinding()  {}

// aaagroup_vpnurl_binding
func (s *AAAService) AddAAAGroupVPNURLBinding()    {}
func (s *AAAService) DeleteAAAGroupVPNURLBinding() {}
func (s *AAAService) GetAllAAAGroupVPNURLBinding() {}
func (s *AAAService) GetAAAGroupVPNURLBinding()    {}
func (s *AAAService) CountAAAGroupVPNURLBinding()  {}

// aaakcdaccount
func (s *AAAService) AddAAAKCDAccount()    {}
func (s *AAAService) DeleteAAAKCDAccount() {}
func (s *AAAService) UpdateAAAKCDAccount() {}
func (s *AAAService) UpsetAAAKCDAccount()  {}
func (s *AAAService) GetAllAAAKCDAccount() {}
func (s *AAAService) GetAAAKCDAccount()    {}
func (s *AAAService) CountAAAKCDAccount()  {}
func (s *AAAService) CheckAAAKCDAccount()  {}

// aaaldapparams
func (s *AAAService) UpdateAAALDAPParams() {}
func (s *AAAService) UnsetAAALDAPParams()  {}
func (s *AAAService) GetAllAAALDAPParams() {}

// aaaotpparamter
func (s *AAAService) UpdateAAAOTPParameter() {}
func (s *AAAService) UnsetAAAOTPParameter()  {}
func (s *AAAService) GetAllAAAOTPParameter() {}

// aaaparameter
func (s *AAAService) UpdateAAAParameter() {}
func (s *AAAService) UnsetAAAParameter()  {}
func (s *AAAService) GetAllAAAParameter() {}

// aaapreauthenticationaction
func (s *AAAService) AddAAAPreauthenticationAction()    {}
func (s *AAAService) DeleteAAAPreauthenticationAction() {}
func (s *AAAService) UpdateAAAPreauthenticationAction() {}
func (s *AAAService) UnsetAAAPreauthenticationAction()  {}
func (s *AAAService) GetAllAAAPreauthenticationAction() {}
func (s *AAAService) GetAAAPreauthenticationAction()    {}
func (s *AAAService) CountAAAPreauthenticationAction()  {}

// aaapreauthenticationparameter
func (s *AAAService) UpdateAAAPreauthenticationParamter() {}
func (s *AAAService) UnsetAAAPreauthenticationParamter()  {}
func (s *AAAService) GetAllAAAPreauthenticationParamter() {}

// aaapreauthenticationpolicy
func (s *AAAService) AddAAAPreauthenticationPolicy()    {}
func (s *AAAService) DeleteAAAPreauthenticationPolicy() {}
func (s *AAAService) UpdateAAAPreauthenticationPolicy() {}
func (s *AAAService) GetAllAAAPreauthenticationPolicy() {}
func (s *AAAService) GetAAAPreauthenticationPolicy()    {}
func (s *AAAService) CountAAAPreauthenticationPolicy()  {}

// aaapreauthenticationpolicy_aaaglobal_binding
func (s *AAAService) GetAllAAAPreauthenticationPolicyAAAGlobalBinding() {}
func (s *AAAService) GetAAAPreauthenticationPolicyAAAGlobalBinding()    {}
func (s *AAAService) CountAAAPreauthenticationPolicyAAAGlobalBinding()  {}

// aaapreauthenticationpolicy_binding
func (s *AAAService) GetAllAAAPreauthenticationPolicyBinding() {}
func (s *AAAService) GetAAAPreauthenticationPolicyBinding()    {}

// aaapreauthentiucationpolicy_vpnvserver_binding
func (s *AAAService) GetAllAAAPreauthenticationPolicyVPNVServerBinding() {}
func (s *AAAService) GetAAAPreauthenticationPolicyVPNVServerBinding()    {}
func (s *AAAService) CountAAAPreauthenticationPolicyVPNVServerBinding()  {}

// aaaradiusparams
func (s *AAAService) UpdateAAARADIUSParams() {}
func (s *AAAService) UnsetAAARADIUSParams()  {}
func (s *AAAService) GetAllAAARADIUSParams() {}

// aaasession
func (s *AAAService) GetAllAAASession() {}
func (s *AAAService) CountAAASession()  {}
func (s *AAAService) KillAAASession()   {}

// aaassoprofile
func (s *AAAService) AddAAASSOProfile()    {}
func (s *AAAService) DeleteAAASSOProfile() {}
func (s *AAAService) GetAllAAASSOProfile() {}
func (s *AAAService) GetAAASSOProfile()    {}
func (s *AAAService) CountAAASSOProfile()  {}
func (s *AAAService) UpdateAAASSOProfile() {}

// aaatacasparams
func (s *AAAService) UpdateAAATACACSParams() {}
func (s *AAAService) UnsetAAATACACSParams()  {}
func (s *AAAService) GetAllAAATACACSParams() {}

// aaauser
func (s *AAAService) AddAAAUser()    {}
func (s *AAAService) DeleteAAAUser() {}
func (s *AAAService) UpdateAAAUser() {}
func (s *AAAService) GetAllAAAUser() {}
func (s *AAAService) GetAAAUser()    {}
func (s *AAAService) CountAAAUser()  {}
func (s *AAAService) UnlockAAAUser() {}

// aaauser_aaagroup_binding
func (s *AAAService) GetAllAAAUserAAAGroupBinding() {}
func (s *AAAService) GetAAAUserAAAGroupBinding()    {}
func (s *AAAService) CountAAAUserAAAGroupBinding()  {}

// aaauser_auditnslogpolicy_binding
func (s *AAAService) AddAAAUserAuditNSLogPolicyBinding()    {}
func (s *AAAService) DeleteAAAUserAuditNSLogPolicyBinding() {}
func (s *AAAService) GetAllAAAUserAuditNSLogPolicyBinding() {}
func (s *AAAService) GetAAAUserAuditNSLogPolicyBinding()    {}
func (s *AAAService) CountAAAUserAuditNSLogPolicyBinding()  {}

// aaauser_auditsyslogpolicy_binding
func (s *AAAService) AddAAAUserAuditSyslogPolicyBinding()    {}
func (s *AAAService) DeleteAAAUserAuditSyslogPolicyBinding() {}
func (s *AAAService) GetAllAAAUserAuditSyslogPolicyBinding() {}
func (s *AAAService) GetAAAUserAuditSyslogPolicyBinding()    {}
func (s *AAAService) CountAAAUserAuditSyslogPolicyBinding()  {}

// aaauser_authorizationpolicy_binding
func (s *AAAService) AddAAAUserAuthorizationPolicyBinding()    {}
func (s *AAAService) DeleteAAAUserAuthorizationPolicyBinding() {}
func (s *AAAService) GetAllAAAUserAuthorizationPolicyBinding() {}
func (s *AAAService) GetAAAUserAuthorizationPolicyBinding()    {}
func (s *AAAService) CountAAAUserAuthorizationPolicyBinding()  {}

// aaauser_binding
func (s *AAAService) GetAllAAAUserBinding() {}
func (s *AAAService) GetAAAUserBinding()    {}

// aaauser_intranetip6_binding
func (s *AAAService) AddAAAUserIntranetIP6_binding()    {}
func (s *AAAService) DeleteAAAUserIntranetIP6_binding() {}
func (s *AAAService) GetAllAAAUserIntranetIP6_binding() {}
func (s *AAAService) GetAAAUserIntranetIP6_binding()    {}
func (s *AAAService) CountAAAUserIntranetIP6_binding()  {}

// aaauser_intranetip_binding
func (s *AAAService) AddAAAUserIntranetIPBinding()    {}
func (s *AAAService) DeleteAAAUserIntranetIPBinding() {}
func (s *AAAService) GetAllAAAUserIntranetIPBinding() {}
func (s *AAAService) GetAAAUserIntranetIPBinding()    {}
func (s *AAAService) CountAAAUserIntranetIPBinding()  {}

// aaauser_tmsessionpolicy_binding
func (s *AAAService) AddAAAUserTMSessionPolicyBinding()    {}
func (s *AAAService) DeleteAAAUserTMSessionPolicyBinding() {}
func (s *AAAService) GetAllAAAUserTMSessionPolicyBinding() {}
func (s *AAAService) GetAAAUserTMSessionPolicyBinding()    {}
func (s *AAAService) CountAAAUserTMSessionPolicyBinding()  {}

// aaauser_vpnintranetapplication_binding
func (s *AAAService) AddAAAUserVPNIntranetApplicationBinding()    {}
func (s *AAAService) DeleteAAAUserVPNIntranetApplicationBinding() {}
func (s *AAAService) GetAllAAAUserVPNIntranetApplicationBinding() {}
func (s *AAAService) GetAAAUserVPNIntranetApplicationBinding()    {}
func (s *AAAService) CountAAAUserVPNIntranetApplicationBinding()  {}

// aaauser_vpnsessionpolicy_binding
func (s *AAAService) AddAAAUserVPNSessionPolicyBinding()    {}
func (s *AAAService) DeleteAAAUserVPNSessionPolicyBinding() {}
func (s *AAAService) GetAllAAAUserVPNSessionPolicyBinding() {}
func (s *AAAService) GetAAAUserVPNSessionPolicyBinding()    {}
func (s *AAAService) CountAAAUserVPNSessionPolicyBinding()  {}

// aaauser_vpntrafficpolicy_binding
func (s *AAAService) AddAAAUserVPNTrafficPolicyBinding()    {}
func (s *AAAService) DeleteAAAUserVPNTrafficPolicyBinding() {}
func (s *AAAService) GetAllAAAUserVPNTrafficPolicyBinding() {}
func (s *AAAService) GetAAAUserVPNTrafficPolicyBinding()    {}
func (s *AAAService) CountAAAUserVPNTrafficPolicyBinding()  {}

// aaauser_vpnurlpolicy_binding
func (s *AAAService) AddAAAUserVPNURLPolicyBinding()    {}
func (s *AAAService) DeleteAAAUserVPNURLPolicyBinding() {}
func (s *AAAService) GetAllAAAUserVPNURLPolicyBinding() {}
func (s *AAAService) GetAAAUserVPNURLPolicyBinding()    {}
func (s *AAAService) CountAAAUserVPNURLPolicyBinding()  {}

// aaauser_vpnurl_binding
func (s *AAAService) AddAAAUserVPNURLBinding()    {}
func (s *AAAService) DeleteAAAUserVPNURLBinding() {}
func (s *AAAService) GetAllAAAUserVPNURLBinding() {}
func (s *AAAService) GetAAAUserVPNURLBinding()    {}
func (s *AAAService) CountAAAUserVPNURLBinding()  {}
