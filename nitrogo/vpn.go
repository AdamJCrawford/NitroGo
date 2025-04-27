package nitrogo

// SSL VPN
// Virtual Private Network configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/vpn/vpn
type VPNService struct {
	client *Client
}

// vpnalwaysonprofile
func (s *VPNService) AddVPNAlwaysOnProfile()    {}
func (s *VPNService) DeleteVPNAlwaysOnProfile() {}
func (s *VPNService) UpdateVPNAlwaysOnProfile() {}
func (s *VPNService) UnsetVPNAlwaysOnProfile()  {}
func (s *VPNService) GetAllVPNAlwaysOnProfile() {}
func (s *VPNService) GetVPNAlwaysOnProfile()    {}
func (s *VPNService) CountVPNAlwaysOnProfile()  {}

// vpnclientlessaccesspolicy
func (s *VPNService) AddVPNClientLessAccessPolicy()    {}
func (s *VPNService) DeleteVPNClientLessAccessPolicy() {}
func (s *VPNService) UpdateVPNClientLessAccessPolicy() {}
func (s *VPNService) GetAllVPNClientLessAccessPolicy() {}
func (s *VPNService) GetVPNClientLessAccessPolicy()    {}
func (s *VPNService) CountVPNClientLessAccessPolicy()  {}

// vpnclientlessaccesspolicy_binding
func (s *VPNService) GetAllVPNClientLessAccessPolicyBinding() {}
func (s *VPNService) GetVPNClientLessAccessPolicyBinding()    {}

// vpnclientlessaccesspolicy_vpnglobal_binding
func (s *VPNService) GetAllVPNClientLessAccessPolicyVPNGlobalBinding() {}
func (s *VPNService) GetVPNClientLessAccessPolicyVPNGlobalBinding()    {}
func (s *VPNService) CountVPNClientLessAccessPolicyVPNGlobalBinding()  {}

// vpnclientlessaccesspolicy_vpnvserver_binding
func (s *VPNService) GetAllVPNClientLessAccessPolicyVPNVServerBinding() {}
func (s *VPNService) GetVPNClientLessAccessPolicyVPNVServerBinding()    {}
func (s *VPNService) CountVPNClientLessAccessPolicyVPNVServerBinding()  {}

// vpnclientlessaccessprofile
func (s *VPNService) AddVPNClientLessAccessProfile()    {}
func (s *VPNService) DeleteVPNClientLessAccessProfile() {}
func (s *VPNService) UpdateVPNClientLessAccessProfile() {}
func (s *VPNService) UnsetVPNClientLessAccessProfile()  {}
func (s *VPNService) GetAllVPNClientLessAccessProfile() {}
func (s *VPNService) GetVPNClientLessAccessProfile()    {}
func (s *VPNService) CountVPNClientLessAccessProfile()  {}

// vpnepaprofile
func (s *VPNService) AddVPNEPAProfile()    {}
func (s *VPNService) DeleteVPNEPAProfile() {}
func (s *VPNService) GetAllVPNEPAProfile() {}
func (s *VPNService) GetVPNEPAProfile()    {}
func (s *VPNService) CountVPNEPAProfile()  {}

// vpneula
func (s *VPNService) AddVPNEULA()    {}
func (s *VPNService) DeleteVPNEULA() {}
func (s *VPNService) GetAllVPNEULA() {}
func (s *VPNService) GetVPNEULA()    {}
func (s *VPNService) CountVPNEULA()  {}

// vpnformssoaction
func (s *VPNService) AddVPNFormSSOAction()    {}
func (s *VPNService) DeleteVPNFormSSOAction() {}
func (s *VPNService) UpdateVPNFormSSOAction() {}
func (s *VPNService) UnsetVPNFormSSOAction()  {}
func (s *VPNService) GetAllVPNFormSSOAction() {}
func (s *VPNService) GetVPNFormSSOAction()    {}
func (s *VPNService) CountVPNFormSSOAction()  {}

// vpnglobal_appcontroller_binding
func (s *VPNService) AddVPNGlobalAppControllerBinding()    {}
func (s *VPNService) DeleteVPNGlobalAppControllerBinding() {}
func (s *VPNService) GetVPNGlobalAppControllerBinding()    {}
func (s *VPNService) CountVPNGlobalAppControllerBinding()  {}

// vpnglobal_auditnslogpolicy_binding
func (s *VPNService) AddVPNGlobalAuditNSLogPolicyBinding()    {}
func (s *VPNService) DeleteVPNGlobalAuditNSLogPolicyBinding() {}
func (s *VPNService) GetVPNGlobalAuditNSLogPolicyBinding()    {}
func (s *VPNService) CountVPNGlobalAuditNSLogPolicyBinding()  {}

// vpnglobal_auditsyslogpolicy_binding
func (s *VPNService) AddVPNGlobalAuditSyslogPolicyBinding()    {}
func (s *VPNService) DeleteVPNGlobalAuditSyslogPolicyBinding() {}
func (s *VPNService) GetVPNGlobalAuditSyslogPolicyBinding()    {}
func (s *VPNService) CountVPNGlobalAuditSyslogPolicyBinding()  {}

// vpnglobal_authenticationcertpolicy_binding
func (s *VPNService) AddVPNGlobalAuthenticationCertPolicyBinding()    {}
func (s *VPNService) DeleteVPNGlobalAuthenticationCertPolicyBinding() {}
func (s *VPNService) GetVPNGlobalAuthenticationCertPolicyBinding()    {}
func (s *VPNService) CountVPNGlobalAuthenticationCertPolicyBinding()  {}

// vpnglobal_authenticationldappolicy_binding
func (s *VPNService) AddVPNGlobalAuthenticationLDAPPolicyBinding()    {}
func (s *VPNService) DeleteVPNGlobalAuthenticationLDAPPolicyBinding() {}
func (s *VPNService) GetVPNGlobalAuthenticationLDAPPolicyBinding()    {}
func (s *VPNService) CountVPNGlobalAuthenticationLDAPPolicyBinding()  {}

// vpnglobal_authenticationlocalpolicy_binding
func (s *VPNService) AddVPNGlobalAuthenticationLocalPolicyBinding()    {}
func (s *VPNService) DeleteVPNGlobalAuthenticationLocalPolicyBinding() {}
func (s *VPNService) GetVPNGlobalAuthenticationLocalPolicyBinding()    {}
func (s *VPNService) CountVPNGlobalAuthenticationLocalPolicyBinding()  {}

// vpnglobal_authenticationnegotiatepolicy_binding
func (s *VPNService) AddVPNGlobalAuthenticationNegotiatePolicyBinding()    {}
func (s *VPNService) DeleteVPNGlobalAuthenticationNegotiatePolicyBinding() {}
func (s *VPNService) GetVPNGlobalAuthenticationNegotiatePolicyBinding()    {}
func (s *VPNService) CountVPNGlobalAuthenticationNegotiatePolicyBinding()  {}

// vpnglobal_authenticationpolicy_binding
func (s *VPNService) AddVPNGlobalAuthenticationPolicyBinding()    {}
func (s *VPNService) DeleteVPNGlobalAuthenticationPolicyBinding() {}
func (s *VPNService) GetVPNGlobalAuthenticationPolicyBinding()    {}
func (s *VPNService) CountVPNGlobalAuthenticationPolicyBinding()  {}

// vpnglobal_authenticationradiuspolicy_binding
func (s *VPNService) AddVPNGlobalAuthenticationRADIUSPolicyBinding()    {}
func (s *VPNService) DeleteVPNGlobalAuthenticationRADIUSPolicyBinding() {}
func (s *VPNService) GetVPNGlobalAuthenticationRADIUSPolicyBinding()    {}
func (s *VPNService) CountVPNGlobalAuthenticationRADIUSPolicyBinding()  {}

// vpnglobal_authenticationsamlpolicy_binding
func (s *VPNService) AddVPNGlobalAuthenticationSAMLPolicyBinding()    {}
func (s *VPNService) DeleteVPNGlobalAuthenticationSAMLPolicyBinding() {}
func (s *VPNService) GetVPNGlobalAuthenticationSAMLPolicyBinding()    {}
func (s *VPNService) CountVPNGlobalAuthenticationSAMLPolicyBinding()  {}

// vpnglobal_authenticationtacacspolicy_binding
func (s *VPNService) AddVPNGlobalAuthenticationTACACSPolicy()    {}
func (s *VPNService) DeleteVPNGlobalAuthenticationTACACSPolicy() {}
func (s *VPNService) GetVPNGlobalAuthenticationTACACSPolicy()    {}
func (s *VPNService) CountVPNGlobalAuthenticationTACACSPolicy()  {}

// vpnglobal_binding
func (s *VPNService) GetVPNGlobalBinding() {}

// vpnglobal_domain_binding
func (s *VPNService) AddVPNGlobalDomainBinding()    {}
func (s *VPNService) DeleteVPNGlobalDomainBinding() {}
func (s *VPNService) GetVPNGlobalDomainBinding()    {}
func (s *VPNService) CountVPNGlobalDomainBinding()  {}

// vpnglobal_intranetip6_binding
func (s *VPNService) AddVPNGlobalIntranetIP6Binding()    {}
func (s *VPNService) DeleteVPNGlobalIntranetIP6Binding() {}
func (s *VPNService) GetVPNGlobalIntranetIP6Binding()    {}
func (s *VPNService) CountVPNGlobalIntranetIP6Binding()  {}

// vpnglobal_intranetip_binding
func (s *VPNService) AddVPNGlobalIntranetIPBinding()    {}
func (s *VPNService) DeleteVPNGlobalIntranetIPBinding() {}
func (s *VPNService) GetVPNGlobalIntranetIPBinding()    {}
func (s *VPNService) CountVPNGlobalIntranetIPBinding()  {}

// vpnglobal_sharefileserver_binding
func (s *VPNService) AddVPNGlobalShareFileServerBinding()    {}
func (s *VPNService) DeleteVPNGlobalShareFileServerBinding() {}
func (s *VPNService) GetVPNGlobalShareFileServerBinding()    {}
func (s *VPNService) CountVPNGlobalShareFileServerBinding()  {}

// vpnglobal_sslcertkey_binding
func (s *VPNService) AddVPNGlobalSSLCertKeyBinding()    {}
func (s *VPNService) DeleteVPNGlobalSSLCertKeyBinding() {}
func (s *VPNService) GetVPNGlobalSSLCertKeyBinding()    {}
func (s *VPNService) CountVPNGlobalSSLCertKeyBinding()  {}

// vpnglobal_staserver_binding
func (s *VPNService) AddVPNGlobalSTAServerBinding()    {}
func (s *VPNService) DeleteVPNGlobalSTAServerBinding() {}
func (s *VPNService) GetVPNGlobalSTAServerBinding()    {}
func (s *VPNService) CountVPNGlobalSTAServerBinding()  {}

// vpnglobal_vpnclientlessaccesspolicy_binding
func (s *VPNService) AddVPNGlobalVPNClientLessAccessPolicyBinding()    {}
func (s *VPNService) DeleteVPNGlobalVPNClientLessAccessPolicyBinding() {}
func (s *VPNService) GetVPNGlobalVPNClientLessAccessPolicyBinding()    {}
func (s *VPNService) CountVPNGlobalVPNClientLessAccessPolicyBinding()  {}

// vpnglobal_vpneula_binding
func (s *VPNService) AddVPNGlobalVPNEULABinding()    {}
func (s *VPNService) DeleteVPNGlobalVPNEULABinding() {}
func (s *VPNService) GetVPNGlobalVPNEULABinding()    {}
func (s *VPNService) CountVPNGlobalVPNEULABinding()  {}

// vpnglobal_vpnintranetapplication_binding
func (s *VPNService) AddVPNGlobalVPNIntranetApplicationBinding()    {}
func (s *VPNService) DeleteVPNGlobalVPNIntranetApplicationBinding() {}
func (s *VPNService) GetVPNGlobalVPNIntranetApplicationBinding()    {}
func (s *VPNService) CountVPNGlobalVPNIntranetApplicationBinding()  {}

// vpnglobal_vpnnexthopserver_binding
func (s *VPNService) AddVPNGlobalVPNNextHopServerBinding()    {}
func (s *VPNService) DeleteVPNGlobalVPNNextHopServerBinding() {}
func (s *VPNService) GetVPNGlobalVPNNextHopServerBinding()    {}
func (s *VPNService) CountVPNGlobalVPNNextHopServerBinding()  {}

// vpnglobal_vpnportaltheme_binding
func (s *VPNService) AddVPNGlobalVPNPortalThemeBinding()    {}
func (s *VPNService) DeleteVPNGlobalVPNPortalThemeBinding() {}
func (s *VPNService) GetVPNGlobalVPNPortalThemeBinding()    {}
func (s *VPNService) CountVPNGlobalVPNPortalThemeBinding()  {}

// vpnglobal_vpnsessionpolcy_binding
func (s *VPNService) AddVPNGlobalVPNSessionPolicyBinding()    {}
func (s *VPNService) DeleteVPNGlobalVPNSessionPolicyBinding() {}
func (s *VPNService) GetVPNGlobalVPNSessionPolicyBinding()    {}
func (s *VPNService) CountVPNGlobalVPNSessionPolicyBinding()  {}

// vpnglobal_vpntrafficpolicy_binding
func (s *VPNService) AddVPNGlobalVPNTrafficPolicyBinding()    {}
func (s *VPNService) DeleteVPNGlobalVPNTrafficPolicyBinding() {}
func (s *VPNService) GetVPNGlobalVPNTrafficPolicyBinding()    {}
func (s *VPNService) CountVPNGlobalVPNTrafficPolicyBinding()  {}

// vpnglobal_vpnurlpolicy_binding
func (s *VPNService) AddVPNGlobalVPNURLPolicyBinding()    {}
func (s *VPNService) DeleteVPNGlobalVPNURLPolicyBinding() {}
func (s *VPNService) GetVPNGlobalVPNURLPolicyBinding()    {}
func (s *VPNService) CountVPNGlobalVPNURLPolicyBinding()  {}

// vpnglobal_vpnurl_binding
func (s *VPNService) AddVPNGlobalVPNURLBinding()    {}
func (s *VPNService) DeleteVPNGlobalVPNURLBinding() {}
func (s *VPNService) GetVPNGlobalVPNURLBinding()    {}
func (s *VPNService) CountVPNGlobalVPNURLBinding()  {}

// vpnicaconnection
func (s *VPNService) KillVPNICAConnection()   {}
func (s *VPNService) GetAllVPNICAConnection() {}
func (s *VPNService) CountVPNICAConnection()  {}

// vpnicadtlsconnection
func (s *VPNService) GetAllVPNICADTLSConnection() {}
func (s *VPNService) CountVPNICADTLSConnection()  {}

// vpnintranetapplication
func (s *VPNService) AddVPNIntranetApplication()    {}
func (s *VPNService) DeleteVPNIntranetApplication() {}
func (s *VPNService) GetAllVPNIntranetApplication() {}
func (s *VPNService) GetVPNIntranetApplication()    {}
func (s *VPNService) CountVPNIntranetApplication()  {}

// vpnnexthopserver
func (s *VPNService) AddVPNNextHopServer()    {}
func (s *VPNService) DeleteVPNNextHopServer() {}
func (s *VPNService) GetAllVPNNextHopServer() {}
func (s *VPNService) GetVPNNextHopServer()    {}
func (s *VPNService) CountVPNNextHopServer()  {}

// vpnparameter
func (s *VPNService) UpdateVPNParameter() {}
func (s *VPNService) UnsetVPNParameter()  {}
func (s *VPNService) GetAllVPNParameter() {}

// vpnpcoipconnection
func (s *VPNService) KillVPNPCoIPConnection()   {}
func (s *VPNService) GetAllVPNPCoIPConnection() {}
func (s *VPNService) CountVPNPCoIPConnection()  {}

// vpnpcoipprofile
func (s *VPNService) AddVPNCoIPProfile()    {}
func (s *VPNService) DeleteVPNCoIPProfile() {}
func (s *VPNService) UpdateVPNCoIPProfile() {}
func (s *VPNService) UnsetVPNCoIPProfile()  {}
func (s *VPNService) GetAllVPNCoIPProfile() {}
func (s *VPNService) GetVPNCoIPProfile()    {}
func (s *VPNService) CountVPNCoIPProfile()  {}

// vpnpcoipvserverprofile
func (s *VPNService) AddVPNPCoIPVServerProfile()    {}
func (s *VPNService) DeleteVPNPCoIPVServerProfile() {}
func (s *VPNService) UpdateVPNPCoIPVServerProfile() {}
func (s *VPNService) UnsetVPNPCoIPVServerProfile()  {}
func (s *VPNService) GetAllVPNPCoIPVServerProfile() {}
func (s *VPNService) GetVPNPCoIPVServerProfile()    {}
func (s *VPNService) CountVPNPCoIPVServerProfile()  {}

// vpnportaltheme
func (s *VPNService) AddVPNPortalTheme()    {}
func (s *VPNService) DeleteVPNPortalTheme() {}
func (s *VPNService) GetAllVPNPortalTheme() {}
func (s *VPNService) GetVPNPortalTheme()    {}
func (s *VPNService) CountVPNPortalTheme()  {}

// vpnsamlssoprofile
func (s *VPNService) AddVPNSAMLSSOProfile()    {}
func (s *VPNService) DeleteVPNSAMLSSOProfile() {}
func (s *VPNService) UpdateVPNSAMLSSOProfile() {}
func (s *VPNService) UnsetVPNSAMLSSOProfile()  {}
func (s *VPNService) GetAllVPNSAMLSSOProfile() {}
func (s *VPNService) GetVPNSAMLSSOProfile()    {}
func (s *VPNService) CountVPNSAMLSSOProfile()  {}

// vpnsessionaction
func (s *VPNService) AddVPNSessionAction()    {}
func (s *VPNService) DeleteVPNSessionAction() {}
func (s *VPNService) UpdateVPNSessionAction() {}
func (s *VPNService) UnsetVPNSessionAction()  {}
func (s *VPNService) GetAllVPNSessionAction() {}
func (s *VPNService) GetVPNSessionAction()    {}
func (s *VPNService) CountVPNSessionAction()  {}

// vpnsessionpolicy
func (s *VPNService) AddVPNSessionPolicy()    {}
func (s *VPNService) DeleteVPNSessionPolicy() {}
func (s *VPNService) UpdateVPNSessionPolicy() {}
func (s *VPNService) UnsetVPNSessionPolicy()  {}
func (s *VPNService) GetAllVPNSessionPolicy() {}
func (s *VPNService) GetVPNSessionPolicy()    {}
func (s *VPNService) CountVPNSessionPolicy()  {}

// vpnsesisonpolicy_aaagroup_binding
func (s *VPNService) GetAllVPNSessionPolicyAAAGroupBinding() {}
func (s *VPNService) GetVPNSessionPolicyAAAGroupBinding()    {}
func (s *VPNService) CountVPNSessionPolicyAAAGroupBinding()  {}

// vpnsessionpolicy_aaauser_binding
func (s *VPNService) GetAllVPNSessionPolicyAAAUserBinding() {}
func (s *VPNService) GetVPNSessionPolicyAAAUserBinding()    {}
func (s *VPNService) CountVPNSessionPolicyAAAUserBinding()  {}

// vpnsessionpolicy_binding
func (s *VPNService) GetAllVPNSessionPolicyBinding() {}
func (s *VPNService) GetVPNSessionPolicyBinding()    {}

// vpnsessionpolicy_vpnglobal_binding
func (s *VPNService) GetAllVPNSessionPolicyVPNGlobalBinding() {}
func (s *VPNService) GetVPNSessionPolicyVPNGlobalBinding()    {}
func (s *VPNService) CountVPNSessionPolicyVPNGlobalBinding()  {}

// vpnsessionpolicy_vpnvserver_binding
func (s *VPNService) GetAllVPNSessionPolicyVPNVServerBinding() {}
func (s *VPNService) GetVPNSessionPolicyVPNVServerBinding()    {}
func (s *VPNService) CountVPNSessionPolicyVPNVServerBinding()  {}

// vpnsfconfig
func (s *VPNService) GetAllVPNSFConfig() {}
func (s *VPNService) CountVPNSFConfig()  {}

// vpnstoreinfo
func (s *VPNService) GetAllVPNStoreInfo() {}
func (s *VPNService) CountVPNStoreInfo()  {}

// vpntrafficaction
func (s *VPNService) AddVPNTrafficAction()    {}
func (s *VPNService) DeleteVPNTrafficAction() {}
func (s *VPNService) UpdateVPNTrafficAction() {}
func (s *VPNService) UnsetVPNTrafficAction()  {}
func (s *VPNService) GetAllVPNTrafficAction() {}
func (s *VPNService) GetVPNTrafficAction()    {}
func (s *VPNService) CountVPNTrafficAction()  {}

// vpntrafficpolicy
func (s *VPNService) AddVPNTrafficPolicy()    {}
func (s *VPNService) DeleteVPNTrafficPolicy() {}
func (s *VPNService) UpdateVPNTrafficPolicy() {}
func (s *VPNService) UnsetVPNTrafficPolicy()  {}
func (s *VPNService) GetAllVPNTrafficPolicy() {}
func (s *VPNService) GetVPNTrafficPolicy()    {}
func (s *VPNService) CountVPNTrafficPolicy()  {}

// vpntrafficpolicy_aaagroup_binding
func (s *VPNService) GetAllVPNTrafficPolicyAAAGroupBinding() {}
func (s *VPNService) GetVPNTrafficPolicyAAAGroupBinding()    {}
func (s *VPNService) CountVPNTrafficPolicyAAAGroupBinding()  {}

// vpntrafficpolicy_aaauser_binding
func (s *VPNService) GetAllVPNTrafficPolicyAAAUserBinding() {}
func (s *VPNService) GetVPNTrafficPolicyAAAUserBinding()    {}
func (s *VPNService) CountVPNTrafficPolicyAAAUserBinding()  {}

// vpntrafficpolicy_binding
func (s *VPNService) GetAllVPNTrafficPolicyBinding() {}
func (s *VPNService) GetVPNTrafficPolicyBinding()    {}

// vpntrafficpolicy_vpnglobal_binding
func (s *VPNService) GetAllVPNTrafficPolicyVPNGlobalBinding() {}
func (s *VPNService) GetVPNTrafficPolicyVPNGlobalBinding()    {}
func (s *VPNService) CountVPNTrafficPolicyVPNGlobalBinding()  {}

// vpntrafficpolicy_vpnvserver_binding
func (s *VPNService) GetAllVPNTrafficPolicyVPNVServerBinding() {}
func (s *VPNService) GetVPNTrafficPolicyVPNVServerBinding()    {}
func (s *VPNService) CountVPNTrafficPolicyVPNVServerBinding()  {}

// vpnurl
func (s *VPNService) AddVPNURL()    {}
func (s *VPNService) DeleteVPNURL() {}
func (s *VPNService) UpdateVPNURL() {}
func (s *VPNService) UnsetVPNURL()  {}
func (s *VPNService) GetAllVPNURL() {}
func (s *VPNService) GetVPNURL()    {}
func (s *VPNService) CountVPNURL()  {}

// vpnurlaction
func (s *VPNService) AddVPNURLAction()    {}
func (s *VPNService) DeleteVPNURLAction() {}
func (s *VPNService) UpdateVPNURLAction() {}
func (s *VPNService) UnsetVPNURLAction()  {}
func (s *VPNService) RenameVPNURLAction() {}
func (s *VPNService) GetAllVPNURLAction() {}
func (s *VPNService) GetVPNURLAction()    {}
func (s *VPNService) CountVPNURLAction()  {}

// vpnurlpolicy
func (s *VPNService) AddVPNURLPolicy()    {}
func (s *VPNService) DeleteVPNURLPolicy() {}
func (s *VPNService) UpdateVPNURLPolicy() {}
func (s *VPNService) UnsetVPNURLPolicy()  {}
func (s *VPNService) RenameVPNURLPolicy() {}
func (s *VPNService) GetAllVPNURLPolicy() {}
func (s *VPNService) GetVPNURLPolicy()    {}
func (s *VPNService) CountVPNURLPolicy()  {}

// vpnurlpolicy_aaagroup_binding
func (s *VPNService) GetAllVPNURLPolicyAAAGroupBinding() {}
func (s *VPNService) GetVPNURLPolicyAAAGroupBinding()    {}
func (s *VPNService) CountVPNURLPolicyAAAGroupBinding()  {}

// vpnurlpolicy_aaauser_binding
func (s *VPNService) GetAllVPNURLPolicyAAAUserBinding() {}
func (s *VPNService) GetVPNURLPolicyAAAUserBinding()    {}
func (s *VPNService) CountVPNURLPolicyAAAUserBinding()  {}

// vpnurlpolicy_binding
func (s *VPNService) GetAllVPNURLPolicyBinding() {}
func (s *VPNService) GetVPNURLPolicyBinding()    {}

// vpnurlpolicy_vpnglobal_binding
func (s *VPNService) GetAllVPNURLPolicyVPNGlobalBinding() {}
func (s *VPNService) GetVPNURLPolicyVPNGlobalBinding()    {}
func (s *VPNService) CountVPNURLPolicyVPNGlobalBinding()  {}

// vpnurlpolicy_vpnvserver_binding
func (s *VPNService) GetAllVPNURLPolicyVPNVServerBinding() {}
func (s *VPNService) GetVPNURLPolicyVPNVServerBinding()    {}
func (s *VPNService) CountVPNURLPolicyVPNVServerBinding()  {}

// vpnvserver
func (s *VPNService) AddVPNVServer()     {}
func (s *VPNService) DeleteVPNVServer()  {}
func (s *VPNService) UpdateVPNVServer()  {}
func (s *VPNService) UnsetVPNVServer()   {}
func (s *VPNService) EnableVPNVServer()  {}
func (s *VPNService) DisableVPNVServer() {}
func (s *VPNService) RenameVPNVServer()  {}
func (s *VPNService) CheckVPNVServer()   {}
func (s *VPNService) GetAllVPNVServer()  {}
func (s *VPNService) GetVPNVServer()     {}
func (s *VPNService) CountVPNVServer()   {}

// vpnvserver_aaapreauthenticationpolicy_binding
func (s *VPNService) AddVPNVServerAAAPreauthenticationPolicyBinding()    {}
func (s *VPNService) DeleteVPNVServerAAAPreauthenticationPolicyBinding() {}
func (s *VPNService) GetAllVPNVServerAAAPreauthenticationPolicyBinding() {}
func (s *VPNService) GetVPNVServerAAAPreauthenticationPolicyBinding()    {}
func (s *VPNService) CountVPNVServerAAAPreauthenticationPolicyBinding()  {}

// vpnvserver_analyticsprofile_binding
func (s *VPNService) AddVPNVServerAnalyticsProfileBinding()    {}
func (s *VPNService) DeleteVPNVServerAnalyticsProfileBinding() {}
func (s *VPNService) GetAllVPNVServerAnalyticsProfileBinding() {}
func (s *VPNService) GetVPNVServerAnalyticsProfileBinding()    {}
func (s *VPNService) CountVPNVServerAnalyticsProfileBinding()  {}

// vpnvserver_appcontroller_binding
func (s *VPNService) AddVPNVServerAppControllerBinding()    {}
func (s *VPNService) DeleteVPNVServerAppControllerBinding() {}
func (s *VPNService) GetAllVPNVServerAppControllerBinding() {}
func (s *VPNService) GetVPNVServerAppControllerBinding()    {}
func (s *VPNService) CountVPNVServerAppControllerBinding()  {}

// vpnvserver_appflowpolicy_binding
func (s *VPNService) AddVPNVServerAppFlowPolicyBinding()    {}
func (s *VPNService) DeleteVPNVServerAppFlowPolicyBinding() {}
func (s *VPNService) GetAllVPNVServerAppFlowPolicyBinding() {}
func (s *VPNService) GetVPNVServerAppFlowPolicyBinding()    {}
func (s *VPNService) CountVPNVServerAppFlowPolicyBinding()  {}

// vpnvserver_auditnslogpolicy_binding
func (s *VPNService) AddVPNVServerAuditNSLogPolicyBinding()    {}
func (s *VPNService) DeleteVPNVServerAuditNSLogPolicyBinding() {}
func (s *VPNService) GetAllVPNVServerAuditNSLogPolicyBinding() {}
func (s *VPNService) VPNVServerAuditNSLogPolicyBinding()       {}
func (s *VPNService) CountVPNVServerAuditNSLogPolicyBinding()  {}

// vpnvserver_auditsyslogpolicy_binding
func (s *VPNService) AddVPNVServerAuditSyslogPolicyBinding()    {}
func (s *VPNService) DeleteVPNVServerAuditSyslogPolicyBinding() {}
func (s *VPNService) GetAllVPNVServerAuditSyslogPolicyBinding() {}
func (s *VPNService) GetVPNVServerAuditSyslogPolicyBinding()    {}
func (s *VPNService) CountVPNVServerAuditSyslogPolicyBinding()  {}

// vpnvserver_authenticationcertpolicy_binding
func (s *VPNService) AddVPNVServerAuthenticationCertPolicyBinding()    {}
func (s *VPNService) DeleteVPNVServerAuthenticationCertPolicyBinding() {}
func (s *VPNService) GetAllVPNVServerAuthenticationCertPolicyBinding() {}
func (s *VPNService) GetVPNVServerAuthenticationCertPolicyBinding()    {}
func (s *VPNService) CountVPNVServerAuthenticationCertPolicyBinding()  {}

// vpnvserver_authenticationfapolicy_binding
func (s *VPNService) AddVPNVServerAuthenticationFAPolicyBinding()    {}
func (s *VPNService) DeleteVPNVServerAuthenticationFAPolicyBinding() {}
func (s *VPNService) GetAllVPNVServerAuthenticationFAPolicyBinding() {}
func (s *VPNService) GetVPNVServerAuthenticationFAPolicyBinding()    {}
func (s *VPNService) CountVPNVServerAuthenticationFAPolicyBinding()  {}

// vpnvserver_authenticationldappolicy_binding
func (s *VPNService) AddVPNVServerAuthenticationLDAPPolicyBinding()    {}
func (s *VPNService) DeleteVPNVServerAuthenticationLDAPPolicyBinding() {}
func (s *VPNService) GetAllVPNVServerAuthenticationLDAPPolicyBinding() {}
func (s *VPNService) GetVPNVServerAuthenticationLDAPPolicyBinding()    {}
func (s *VPNService) CountVPNVServerAuthenticationLDAPPolicyBinding()  {}

// vpnvserver_authenticationlocalpolicy_binding
func (s *VPNService) AddVPNVServerAuthenticationLocalPolicyBinding()    {}
func (s *VPNService) DeleteVPNVServerAuthenticationLocalPolicyBinding() {}
func (s *VPNService) GetAllVPNVServerAuthenticationLocalPolicyBinding() {}
func (s *VPNService) GetVPNVServerAuthenticationLocalPolicyBinding()    {}
func (s *VPNService) CountVPNVServerAuthenticationLocalPolicyBinding()  {}

// vpnvserver_authenticationloginschemapolicy_binding
func (s *VPNService) AddVPNVServerAuthenticationLoginSchemaPolicyBinding()    {}
func (s *VPNService) DeleteVPNVServerAuthenticationLoginSchemaPolicyBinding() {}
func (s *VPNService) GetAllVPNVServerAuthenticationLoginSchemaPolicyBinding() {}
func (s *VPNService) GetVPNVServerAuthenticationLoginSchemaPolicyBinding()    {}
func (s *VPNService) CountVPNVServerAuthenticationLoginSchemaPolicyBinding()  {}

// vpnvserver_authenticationnegotiatepolcy_binding
func (s *VPNService) AddVPNVServerAuthenticationNegotiatePolicyBinding()    {}
func (s *VPNService) DeleteVPNVServerAuthenticationNegotiatePolicyBinding() {}
func (s *VPNService) GetAllVPNVServerAuthenticationNegotiatePolicyBinding() {}
func (s *VPNService) GetVPNVServerAuthenticationNegotiatePolicyBinding()    {}
func (s *VPNService) CountVPNVServerAuthenticationNegotiatePolicyBinding()  {}

// vpnvserver_authenticationoauthidppolicy_binding
func (s *VPNService) AddVPNVServerAuthenticationOAuthIDPPolicyBinding()    {}
func (s *VPNService) DeleteVPNVServerAuthenticationOAuthIDPPolicyBinding() {}
func (s *VPNService) GetAllVPNVServerAuthenticationOAuthIDPPolicyBinding() {}
func (s *VPNService) GetVPNVServerAuthenticationOAuthIDPPolicyBinding()    {}
func (s *VPNService) CountVPNVServerAuthenticationOAuthIDPPolicyBinding()  {}

// vpnvserver_authenticationpolicy_binding
func (s *VPNService) AddVPNVServerAuthenticationPolicyBinding()    {}
func (s *VPNService) DeleteVPNVServerAuthenticationPolicyBinding() {}
func (s *VPNService) GetAllVPNVServerAuthenticationPolicyBinding() {}
func (s *VPNService) GetVPNVServerAuthenticationPolicyBinding()    {}
func (s *VPNService) CountVPNVServerAuthenticationPolicyBinding()  {}

// vpnvserver_authenticationradiuspolicy_binding
func (s *VPNService) AddVPNVServerAuthenticationRADIUSPolicyBinding()    {}
func (s *VPNService) DeleteVPNVServerAuthenticationRADIUSPolicyBinding() {}
func (s *VPNService) GetAllVPNVServerAuthenticationRADIUSPolicyBinding() {}
func (s *VPNService) GetVPNVServerAuthenticationRADIUSPolicyBinding()    {}
func (s *VPNService) CountVPNVServerAuthenticationRADIUSPolicyBinding()  {}

// vpnvserver_authenticationsamlidppolicy_binding
func (s *VPNService) AddVPNVServerAuthenticationSAMLIDPPolicyBinding()    {}
func (s *VPNService) DeleteVPNVServerAuthenticationSAMLIDPPolicyBinding() {}
func (s *VPNService) GetAllVPNVServerAuthenticationSAMLIDPPolicyBinding() {}
func (s *VPNService) GetVPNVServerAuthenticationSAMLIDPPolicyBinding()    {}
func (s *VPNService) CountVPNVServerAuthenticationSAMLIDPPolicyBinding()  {}

// vpnvserver_authenticationsamlpolicy_binding
func (s *VPNService) AddVPNVServerAuthenticationSAMLPolicyBinding()    {}
func (s *VPNService) DeleteVPNVServerAuthenticationSAMLPolicyBinding() {}
func (s *VPNService) GetAllVPNVServerAuthenticationSAMLPolicyBinding() {}
func (s *VPNService) GetVPNVServerAuthenticationSAMLPolicyBinding()    {}
func (s *VPNService) CountVPNVServerAuthenticationSAMLPolicyBinding()  {}

// vpnvserver_authenticationtacacspolicy_binding
func (s *VPNService) AddVPNVServerAuthenticationTACACSPolicyBinding()    {}
func (s *VPNService) DeleteVPNVServerAuthenticationTACACSPolicyBinding() {}
func (s *VPNService) GetAllVPNVServerAuthenticationTACACSPolicyBinding() {}
func (s *VPNService) GetVPNVServerAuthenticationTACACSPolicyBinding()    {}
func (s *VPNService) CountVPNVServerAuthenticationTACACSPolicyBinding()  {}

// vpnvserver_authenticationwebauthpolicy_binding
func (s *VPNService) AddVPNVServerAuthenticationWebAuthPolicyBinding()    {}
func (s *VPNService) DeleteVPNVServerAuthenticationWebAuthPolicyBinding() {}
func (s *VPNService) GetAllVPNVServerAuthenticationWebAuthPolicyBinding() {}
func (s *VPNService) GetVPNVServerAuthenticationWebAuthPolicyBinding()    {}
func (s *VPNService) CountVPNVServerAuthenticationWebAuthPolicyBinding()  {}

// vpnvserver_binding
func (s *VPNService) GetAllPVNVServerBinding() {}
func (s *VPNService) GetPVNVServerBinding()    {}

// vpnvserver_cachepolicy_binding
func (s *VPNService) AddVPNVServerCachePolicyBinding()    {}
func (s *VPNService) DeleteVPNVServerCachePolicyBinding() {}
func (s *VPNService) GetAllVPNVServerCachePolicyBinding() {}
func (s *VPNService) GetVPNVServerCachePolicyBinding()    {}
func (s *VPNService) CountVPNVServerCachePolicyBinding()  {}

// vpnvserver_cspolicy_binding
func (s *VPNService) AddVPNVServerCSPolicyBinding()    {}
func (s *VPNService) DeleteVPNVServerCSPolicyBinding() {}
func (s *VPNService) GetAllVPNVServerCSPolicyBinding() {}
func (s *VPNService) GetVPNVServerCSPolicyBinding()    {}
func (s *VPNService) CountVPNVServerCSPolicyBinding()  {}

// vpnvserver_feopolicy_binding
func (s *VPNService) AddVPNVServerFEOPolicyBinding()    {}
func (s *VPNService) DeleteVPNVServerFEOPolicyBinding() {}
func (s *VPNService) GetAllVPNVServerFEOPolicyBinding() {}
func (s *VPNService) GetVPNVServerFEOPolicyBinding()    {}
func (s *VPNService) CountVPNVServerFEOPolicyBinding()  {}

// vpnvserver_icapolicy_binding
func (s *VPNService) AddVPNVServerICAPolicyBinding()    {}
func (s *VPNService) DeleteVPNVServerICAPolicyBinding() {}
func (s *VPNService) GetAllVPNVServerICAPolicyBinding() {}
func (s *VPNService) GetVPNVServerICAPolicyBinding()    {}
func (s *VPNService) CountVPNVServerICAPolicyBinding()  {}

// vpnvserver_intranetip6_binding
func (s *VPNService) AddVPNVServerIntranetIP6Binding()    {}
func (s *VPNService) DeleteVPNVServerIntranetIP6Binding() {}
func (s *VPNService) GetAllVPNVServerIntranetIP6Binding() {}
func (s *VPNService) GetVPNVServerIntranetIP6Binding()    {}
func (s *VPNService) CountVPNVServerIntranetIP6Binding()  {}

// vpnvserver_intranetip_binding
func (s *VPNService) AddVPNVServerIntranetIPBinding()    {}
func (s *VPNService) DeleteVPNVServerIntranetIPBinding() {}
func (s *VPNService) GetAllVPNVServerIntranetIPBinding() {}
func (s *VPNService) GetVPNVServerIntranetIPBinding()    {}
func (s *VPNService) CountVPNVServerIntranetIPBinding()  {}

// vpnvserver_responderpolicy_binding
func (s *VPNService) AddVPNVServerResponderPolicyBinding()    {}
func (s *VPNService) DeleteVPNVServerResponderPolicyBinding() {}
func (s *VPNService) GetAllVPNVServerResponderPolicyBinding() {}
func (s *VPNService) GetVPNVServerResponderPolicyBinding()    {}
func (s *VPNService) CountVPNVServerResponderPolicyBinding()  {}

// vpnvserver_rewritepolicy_binding
func (s *VPNService) AddVPNVServerRewritePolicyBinding()    {}
func (s *VPNService) DeleteVPNVServerRewritePolicyBinding() {}
func (s *VPNService) GetAllVPNVServerRewritePolicyBinding() {}
func (s *VPNService) GetVPNVServerRewritePolicyBinding()    {}
func (s *VPNService) CountVPNVServerRewritePolicyBinding()  {}

// vpnvserver_sharefileserver_binding
func (s *VPNService) AddVPNVServerShareFileServerBinding()    {}
func (s *VPNService) DeleteVPNVServerShareFileServerBinding() {}
func (s *VPNService) GetAllVPNVServerShareFileServerBinding() {}
func (s *VPNService) GetVPNVServerShareFileServerBinding()    {}
func (s *VPNService) CountVPNVServerShareFileServerBinding()  {}

// vpnvserver_staserver_binding
func (s *VPNService) AddVPNVServerSTAServerBinding()    {}
func (s *VPNService) DeleteVPNVServerSTAServerBinding() {}
func (s *VPNService) GetAllVPNVServerSTAServerBinding() {}
func (s *VPNService) GetVPNVServerSTAServerBinding()    {}
func (s *VPNService) CountVPNVServerSTAServerBinding()  {}

// vpnvserver_vpnclientlessaccesspolicy_binding
func (s *VPNService) AddVPNVServerVPNClientlessAccessPolicyBinding()    {}
func (s *VPNService) DeleteVPNVServerVPNClientlessAccessPolicyBinding() {}
func (s *VPNService) GetAllVPNVServerVPNClientlessAccessPolicyBinding() {}
func (s *VPNService) GetVPNVServerVPNClientlessAccessPolicyBinding()    {}
func (s *VPNService) CountVPNVServerVPNClientlessAccessPolicyBinding()  {}

// vpnvserver_vpnepaprofile_binding
func (s *VPNService) AddVPNVServerVPNEPAProfileBinding()    {}
func (s *VPNService) DeleteVPNVServerVPNEPAProfileBinding() {}
func (s *VPNService) GetAllVPNVServerVPNEPAProfileBinding() {}
func (s *VPNService) GetVPNVServerVPNEPAProfileBinding()    {}
func (s *VPNService) CountVPNVServerVPNEPAProfileBinding()  {}

// vpnvserver_vpneula_binding
func (s *VPNService) AddVPNVServerVPNEULABinding()    {}
func (s *VPNService) DeleteVPNVServerVPNEULABinding() {}
func (s *VPNService) GetAllVPNVServerVPNEULABinding() {}
func (s *VPNService) GetVPNVServerVPNEULABinding()    {}
func (s *VPNService) CountVPNVServerVPNEULABinding()  {}

// vpnvserver_vpnintranetapplication_binding
func (s *VPNService) AddVPNVServerVPNIntranetApplicationBinding()    {}
func (s *VPNService) DeleteVPNVServerVPNIntranetApplicationBinding() {}
func (s *VPNService) GetAllVPNVServerVPNIntranetApplicationBinding() {}
func (s *VPNService) GetVPNVServerVPNIntranetApplicationBinding()    {}
func (s *VPNService) CountVPNVServerVPNIntranetApplicationBinding()  {}

// vpnvserver_vpnnexthopserver_binding
func (s *VPNService) AddVPNVServerVPNNextHopServerBinding()    {}
func (s *VPNService) DeleteVPNVServerVPNNextHopServerBinding() {}
func (s *VPNService) GetAllVPNVServerVPNNextHopServerBinding() {}
func (s *VPNService) GetVPNVServerVPNNextHopServerBinding()    {}
func (s *VPNService) CountVPNVServerVPNNextHopServerBinding()  {}

// vpnvserver_vpnportaltheme_binding
func (s *VPNService) AddVPNVServerVPNPortalThemeBinding()    {}
func (s *VPNService) DeleteVPNVServerVPNPortalThemeBinding() {}
func (s *VPNService) GetAllVPNVServerVPNPortalThemeBinding() {}
func (s *VPNService) GetVPNVServerVPNPortalThemeBinding()    {}
func (s *VPNService) CountVPNVServerVPNPortalThemeBinding()  {}

// vpnvserver_vpnsessionpolicy_binding
func (s *VPNService) AddVPNVServerVPNSessionPolicyBinding()    {}
func (s *VPNService) DeleteVPNVServerVPNSessionPolicyBinding() {}
func (s *VPNService) GetAllVPNVServerVPNSessionPolicyBinding() {}
func (s *VPNService) GetVPNVServerVPNSessionPolicyBinding()    {}
func (s *VPNService) CountVPNVServerVPNSessionPolicyBinding()  {}

// vpnvserver_vpntrafficpolicy_binding
func (s *VPNService) AddVPNVServerVPNTrafficPolicyBinding()    {}
func (s *VPNService) DeleteVPNVServerVPNTrafficPolicyBinding() {}
func (s *VPNService) GetAllVPNVServerVPNTrafficPolicyBinding() {}
func (s *VPNService) GetVPNVServerVPNTrafficPolicyBinding()    {}
func (s *VPNService) CountVPNVServerVPNTrafficPolicyBinding()  {}

// vpnvserver_vpnurlpolicy_binding
func (s *VPNService) AddVPNVServerVPNURLPolicyBinding()    {}
func (s *VPNService) DeleteVPNVServerVPNURLPolicyBinding() {}
func (s *VPNService) GetAllVPNVServerVPNURLPolicyBinding() {}
func (s *VPNService) GetVPNVServerVPNURLPolicyBinding()    {}
func (s *VPNService) CountVPNVServerVPNURLPolicyBinding()  {}

// vpnvserver_vpnurl_binding
func (s *VPNService) AddVPNVServerVPNURLBinding()    {}
func (s *VPNService) DeleteVPNVServerVPNURLBinding() {}
func (s *VPNService) GetAllVPNVServerVPNURLBinding() {}
func (s *VPNService) GetVPNVServerVPNURLBinding()    {}
func (s *VPNService) CountVPNVServerVPNURLBinding()  {}
