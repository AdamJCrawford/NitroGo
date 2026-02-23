package nitrogo

const (
	authenticationADFSProxyProfileURL                              = "/nitro/v1/config/authenticationadfsproxyprofile"
	authenticationAuthnProfileURL                                  = "/nitro/v1/config/authenticationauthnprofile"
	authenticationAzureKeyVaultURL                                 = "/nitro/v1/config/authenticationazurekeyvault"
	authenticationCaptchaActionURL                                 = "/nitro/v1/config/authenticationcaptchaaction"
	authenticationCertActionURL                                    = "/nitro/v1/config/authenticationcertaction"
	authenticationCertPolicyURL                                    = "/nitro/v1/config/authenticationcertpolicy"
	authenticationCertPolicyAuthenticationVServerBindingURL        = "/nitro/v1/config/authenticationcertpolicy_authenticationvserver_binding"
	authenticationCertPolicyBindingURL                             = "/nitro/v1/config/authenticationcertpolicy_binding"
	authenticationCertPolicyVPNGlobalBindingURL                    = "/nitro/v1/config/authenticationcertpolicy_vpnglobal_binding"
	authenticationCertPolicyVPNVServerBindingURL                   = "/nitro/v1/config/authenticationcertpolicy_vpnvserver_binding"
	authenticationCitrixAuthActionURL                              = "/nitro/v1/config/authenticationcitrixauthaction"
	authenticationDFAActionURL                                     = "/nitro/v1/config/authenticationdfaaction"
	authenticationDFAPolicyURL                                     = "/nitro/v1/config/authenticationdfapolicy"
	authenticationDFAPolicyBindingURL                              = "/nitro/v1/config/authenticationdfapolicy_binding"
	authenticationDFAPolicyVPNVServerBindingURL                    = "/nitro/v1/config/authenticationdfapolicy_vpnvserver_binding"
	authenticationEmailActionURL                                   = "/nitro/v1/config/authenticationemailaction"
	authenticationEPAActionURL                                     = "/nitro/v1/config/authenticationepaaction"
	authenticationLDAPActionURL                                    = "/nitro/v1/config/authenticationldapaction"
	authenticationLDAPPolicyURL                                    = "/nitro/v1/config/authenticationldappolicy"
	authenticationLDAPPolicyAuthenticationVServerBindingURL        = "/nitro/v1/config/authenticationldappolicy_authenticationvserver_binding"
	authenticationLDAPPolicyBindingURL                             = "/nitro/v1/config/authenticationldappolicy_binding"
	authenticationLDAPPolicySystemGlobalBindingURL                 = "/nitro/v1/config/authenticationldappolicy_systemglobal_binding"
	authenticationLDAPPolicyVPNGlobalBindingURL                    = "/nitro/v1/config/authenticationldappolicy_vpnglobal_binding"
	authenticationLDAPPolicyVPNVServerBindingURL                   = "/nitro/v1/config/authenticationldappolicy_vpnvserver_binding"
	authenticationLocalPolicyURL                                   = "/nitro/v1/config/authenticationlocalpolicy"
	authenticationLocalPolicyAuthenticationVServerBindingURL       = "/nitro/v1/config/authenticationlocalpolicy_authenticationvserver_binding"
	authenticationLocalPolicyBindingURL                            = "/nitro/v1/config/authenticationlocalpolicy_binding"
	authenticationLocalPolicySystemGlobalBindingURL                = "/nitro/v1/config/authenticationlocalpolicy_systemglobal_binding"
	authenticationLocalPolicyVPNGlobalBindingURL                   = "/nitro/v1/config/authenticationlocalpolicy_vpnglobal_binding"
	authenticationLocalPolicyVPNVServerBindingURL                  = "/nitro/v1/config/authenticationlocalpolicy_vpnvserver_binding"
	authenticationLoginSchemaURL                                   = "/nitro/v1/config/authenticationloginschema"
	authenticationLoginSchemaPolicyURL                             = "/nitro/v1/config/authenticationloginschemapolicy"
	authenticationLoginSchemaPolicyAuthenticationVServerBindingURL = "/nitro/v1/config/authenticationloginschemapolicy_authenticationvserver_binding"
	authenticationLoginSchemaPolicyBindingURL                      = "/nitro/v1/config/authenticationloginschemapolicy_binding"
	authenticationLoginSchemaPolicyVPNVServerBindingURL            = "/nitro/v1/config/authenticationloginschemapolicy_vpnvserver_binding"
	authenticationNegotiateActionURL                               = "/nitro/v1/config/authenticationnegotiateaction"
	authenticationNegotiatePolicyURL                               = "/nitro/v1/config/authenticationnegotiatepolicy"
	authenticationNegotiatePolicyAuthenticationVServerBindingURL   = "/nitro/v1/config/authenticationnegotiatepolicy_authenticationvserver_binding"
	authenticationNegotiatePolicyBindingURL                        = "/nitro/v1/config/authenticationnegotiatepolicy_binding"
	authenticationNegotiatePolicyVPNGlobalBindingURL               = "/nitro/v1/config/authenticationnegotiatepolicy_vpnglobal_binding"
	authenticationNegotiatePolicyVPNVServerBindingURL              = "/nitro/v1/config/authenticationnegotiatepolicy_vpnvserver_binding"
	authenticationNoAuthActionURL                                  = "/nitro/v1/config/authenticationnoauthaction"
	authenticationOAuthActionURL                                   = "/nitro/v1/config/authenticationoauthaction"
	authenticationOAuthIdPPolicyURL                                = "/nitro/v1/config/authenticationoauthidppolicy"
	authenticationOAuthIdPPolicyAuthenticationVServerBindingURL    = "/nitro/v1/config/authenticationoauthidppolicy_authenticationvserver_binding"
	authenticationOAuthIdPPolicyBindingURL                         = "/nitro/v1/config/authenticationoauthidppolicy_binding"
	authenticationOAuthIdPPolicyVPNVServerBindingURL               = "/nitro/v1/config/authenticationoauthidppolicy_vpnvserver_binding"
	authenticationOAuthIdPProfileURL                               = "/nitro/v1/config/authenticationoauthidpprofile"
	authenticationPolicyURL                                        = "/nitro/v1/config/authenticationpolicy"
	authenticationPolicyLabelURL                                   = "/nitro/v1/config/authenticationpolicylabel"
	authenticationPolicyLabelAuthenticationPolicyBindingURL        = "/nitro/v1/config/authenticationpolicylabel_authenticationpolicy_binding"
	authenticationPolicyLabelBindingURL                            = "/nitro/v1/config/authenticationpolicylabel_binding"
	authenticationPolicyAuthenticationPolicyLabelBindingURL        = "/nitro/v1/config/authenticationpolicy_authenticationpolicylabel_binding"
	authenticationPolicyAuthenticationVServerBindingURL            = "/nitro/v1/config/authenticationpolicy_authenticationvserver_binding"
	authenticationPolicyBindingURL                                 = "/nitro/v1/config/authenticationpolicy_binding"
	authenticationPolicySystemGlobalBindingURL                     = "/nitro/v1/config/authenticationpolicy_systemglobal_binding"
	authenticationPushServiceURL                                   = "/nitro/v1/config/authenticationpushservice"
	authenticationRADIUSActionURL                                  = "/nitro/v1/config/authenticationradiusaction"
	authenticationRADIUSPolicyURL                                  = "/nitro/v1/config/authenticationradiuspolicy"
	authenticationRADIUSPolicyAuthenticationVServerBindingURL      = "/nitro/v1/config/authenticationradiuspolicy_authenticationvserver_binding"
	authenticationRADIUSPolicyBindingURL                           = "/nitro/v1/config/authenticationradiuspolicy_binding"
	authenticationRADIUSPolicySystemGlobalBindingURL               = "/nitro/v1/config/authenticationradiuspolicy_systemglobal_binding"
	authenticationRADIUSPolicyVPNGlobalBindingURL                  = "/nitro/v1/config/authenticationradiuspolicy_vpnglobal_binding"
	authenticationRADIUSPolicyVPNVServerBindingURL                 = "/nitro/v1/config/authenticationradiuspolicy_vpnvserver_binding"
	authenticationSAMLActionURL                                    = "/nitro/v1/config/authenticationsamlaction"
	authenticationSAMLIdPPolicyURL                                 = "/nitro/v1/config/authenticationsamlidppolicy"
	authenticationSAMLIdPPolicyAuthenticationVServerBindingURL     = "/nitro/v1/config/authenticationsamlidppolicy_authenticationvserver_binding"
	authenticationSAMLIdPPolicyBindingURL                          = "/nitro/v1/config/authenticationsamlidppolicy_binding"
	authenticationSAMLIdPPolicyVPNVServerBindingURL                = "/nitro/v1/config/authenticationsamlidppolicy_vpnvserver_binding"
	authenticationSAMLIdPProfileURL                                = "/nitro/v1/config/authenticationsamlidpprofile"
	authenticationSAMLPolicyURL                                    = "/nitro/v1/config/authenticationsamlpolicy"
	authenticationSAMLPolicyAuthenticationVServerBindingURL        = "/nitro/v1/config/authenticationsamlpolicy_authenticationvserver_binding"
	authenticationSAMLPolicyBindingURL                             = "/nitro/v1/config/authenticationsamlpolicy_binding"
	authenticationSAMLPolicyVPNGlobalBindingURL                    = "/nitro/v1/config/authenticationsamlpolicy_vpnglobal_binding"
	authenticationSAMLPolicyVPNVServerBindingURL                   = "/nitro/v1/config/authenticationsamlpolicy_vpnvserver_binding"
	authenticationStoreFrontAuthActionURL                          = "/nitro/v1/config/authenticationstorefrontauthaction"
	authenticationTACACSActionURL                                  = "/nitro/v1/config/authenticationtacacsaction"
	authenticationTACACSPolicyURL                                  = "/nitro/v1/config/authenticationtacacspolicy"
	authenticationTACACSPolicyAuthenticationVServerBindingURL      = "/nitro/v1/config/authenticationtacacspolicy_authenticationvserver_binding"
	authenticationTACACSPolicyBindingURL                           = "/nitro/v1/config/authenticationtacacspolicy_binding"
	authenticationTACACSPolicySystemGlobalBindingURL               = "/nitro/v1/config/authenticationtacacspolicy_systemglobal_binding"
	authenticationTACACSPolicyVPNGlobalBindingURL                  = "/nitro/v1/config/authenticationtacacspolicy_vpnglobal_binding"
	authenticationTACACSPolicyVPNVServerBindingURL                 = "/nitro/v1/config/authenticationtacacspolicy_vpnvserver_binding"
	authenticationVServerURL                                       = "/nitro/v1/config/authenticationvserver"
	authenticationVServerAuditNSLogPolicyBindingURL                = "/nitro/v1/config/authenticationvserver_auditnslogpolicy_binding"
	authenticationVServerAuditSyslogPolicyBindingURL               = "/nitro/v1/config/authenticationvserver_auditsyslogpolicy_binding"
	authenticationVServerAuthenticationCertPolicyBindingURL        = "/nitro/v1/config/authenticationvserver_authenticationcertpolicy_binding"
	authenticationVServerAuthenticationLDAPPolicyBindingURL        = "/nitro/v1/config/authenticationvserver_authenticationldappolicy_binding"
	authenticationVServerAuthenticationLocalPolicyBindingURL       = "/nitro/v1/config/authenticationvserver_authenticationlocalpolicy_binding"
	authenticationVServerAuthenticationLoginSchemaPolicyBindingURL = "/nitro/v1/config/authenticationvserver_authenticationloginschemapolicy_binding"
	authenticationVServerAuthenticationNegotiatePolicyBindingURL   = "/nitro/v1/config/authenticationvserver_authenticationnegotiatepolicy_binding"
	authenticationVServerAuthenticationOAuthIdPPolicyBindingURL    = "/nitro/v1/config/authenticationvserver_authenticationoauthidppolicy_binding"
	authenticationVServerAuthenticationPolicyBindingURL            = "/nitro/v1/config/authenticationvserver_authenticationpolicy_binding"
	authenticationVServerAuthenticationRADIUSPolicyBindingURL      = "/nitro/v1/config/authenticationvserver_authenticationradiuspolicy_binding"
	authenticationVServerAuthenticationSAMLIdPPolicyBindingURL     = "/nitro/v1/config/authenticationvserver_authenticationsamlidppolicy_binding"
	authenticationVServerAuthenticationSAMLPolicyBindingURL        = "/nitro/v1/config/authenticationvserver_authenticationsamlpolicy_binding"
	authenticationVServerAuthenticationTACACSPolicyBindingURL      = "/nitro/v1/config/authenticationvserver_authenticationtacacspolicy_binding"
	authenticationVServerAuthenticationWebAuthPolicyBindingURL     = "/nitro/v1/config/authenticationvserver_authenticationwebauthpolicy_binding"
	authenticationVServerBindingURL                                = "/nitro/v1/config/authenticationvserver_binding"
	authenticationVServerCachePolicyBindingURL                     = "/nitro/v1/config/authenticationvserver_cachepolicy_binding"
	authenticationVServerCSPolicyBindingURL                        = "/nitro/v1/config/authenticationvserver_cspolicy_binding"
	authenticationVServerResponderPolicyBindingURL                 = "/nitro/v1/config/authenticationvserver_responderpolicy_binding"
	authenticationVServerTMSessionPolicyBindingURL                 = "/nitro/v1/config/authenticationvserver_tmsessionpolicy_binding"
	authenticationVServerVPNPortalThemeBindingURL                  = "/nitro/v1/config/authenticationvserver_vpnportaltheme_binding"
	authenticationWebAuthActionURL                                 = "/nitro/v1/config/authenticationwebauthaction"
	authenticationWebAuthPolicyURL                                 = "/nitro/v1/config/authenticationwebauthpolicy"
	authenticationWebAuthPolicyAuthenticationVServerBindingURL     = "/nitro/v1/config/authenticationwebauthpolicy_authenticationvserver_binding"
	authenticationWebAuthPolicyBindingURL                          = "/nitro/v1/config/authenticationwebauthpolicy_binding"
	authenticationWebAuthPolicySystemGlobalBindingURL              = "/nitro/v1/config/authenticationwebauthpolicy_systemglobal_binding"
	authenticationWebAuthPolicyVPNGlobalBindingURL                 = "/nitro/v1/config/authenticationwebauthpolicy_vpnglobal_binding"
	authenticationWebAuthPolicyVPNVServerBindingURL                = "/nitro/v1/config/authenticationwebauthpolicy_vpnvserver_binding"
)

// Authentication configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authentication
type AuthenticationService struct {
	client *Client
}

// authenticationadfsproxyprofile
// Configuration for ADFSProxy Profile resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationadfsproxyprofile
func (s *AuthenticationService) AddAuthenticationADFSProxyProfile()    {}
func (s *AuthenticationService) DeleteAuthenticationADFSProxyProfile() {}
func (s *AuthenticationService) UpdateAuthenticationADFSProxyProfile() {}
func (s *AuthenticationService) GetAllAuthenticationADFSProxyProfile() {}
func (s *AuthenticationService) GetAuthenticationADFSProxyProfile()    {}
func (s *AuthenticationService) CountAuthenticationADFSProxyProfile()  {}

// authenticationauthnprofile
// Configuration for Authentication profile resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationauthnprofile
func (s *AuthenticationService) AddAuthenticationAuthnProfile()    {}
func (s *AuthenticationService) DeleteAuthenticationAuthnProfile() {}
func (s *AuthenticationService) UpdateAuthenticationAuthnProfile() {}
func (s *AuthenticationService) UnsetAuthenticationAuthnProfile()  {}
func (s *AuthenticationService) GetAllAuthenticationAuthnProfile() {}
func (s *AuthenticationService) GetAuthenticationAuthnProfile()    {}
func (s *AuthenticationService) CountAuthenticationAuthnProfile()  {}

// authenticationazurekeyvault
// Configuration for Azure Key Vault entity resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationazurekeyvault
func (s *AuthenticationService) AddAuthenticationAzureKeyVault()    {}
func (s *AuthenticationService) DeleteAuthenticationAzureKeyVault() {}
func (s *AuthenticationService) UpdateAuthenticationAzureKeyVault() {}
func (s *AuthenticationService) UnsetAuthenticationAzureKeyVault()  {}
func (s *AuthenticationService) GetAllAuthenticationAzureKeyVault() {}
func (s *AuthenticationService) GetAuthenticationAzureKeyVault()    {}
func (s *AuthenticationService) CountAuthenticationAzureKeyVault()  {}

// authenticationcaptchaaction
// Configuration for Captcha Action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationcaptchaaction
func (s *AuthenticationService) AddAuthenticationCaptchaAction()    {}
func (s *AuthenticationService) DeleteAuthenticationCaptchaAction() {}
func (s *AuthenticationService) UpdateAuthenticationCaptchaAction() {}
func (s *AuthenticationService) UnsetAuthenticationCaptchaAction()  {}
func (s *AuthenticationService) GetAllAuthenticationCaptchaAction() {}
func (s *AuthenticationService) GetAuthenticationCaptchaAction()    {}
func (s *AuthenticationService) CountAuthenticationCaptchaAction()  {}

// authenticationcertaction
// Configuration for CERT action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationcertaction
func (s *AuthenticationService) AddAuthenticationCertAction()    {}
func (s *AuthenticationService) DeleteAuthenticationCertAction() {}
func (s *AuthenticationService) UpdateAuthenticationCertAction() {}
func (s *AuthenticationService) UnsetAuthenticationCertAction()  {}
func (s *AuthenticationService) GetAllAuthenticationCertAction() {}
func (s *AuthenticationService) GetAuthenticationCertAction()    {}
func (s *AuthenticationService) CountAuthenticationCertAction()  {}

// authenticationcertpolicy
// Configuration for CERT policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationcertpolicy
func (s *AuthenticationService) AddAuthenticationCertPolicy()    {}
func (s *AuthenticationService) DeleteAuthenticationCertPolicy() {}
func (s *AuthenticationService) UpdateAuthenticationCertPolicy() {}
func (s *AuthenticationService) UnsetAuthenticationCertPolicy()  {}
func (s *AuthenticationService) GetAllAuthenticationCertPolicy() {}
func (s *AuthenticationService) GetAuthenticationCertPolicy()    {}
func (s *AuthenticationService) CountAuthenticationCertPolicy()  {}

// authenticationcertpolicy_authenticationvserver_binding
// Binding object showing the authenticationvserver that can be bound to authenticationcertpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationcertpolicy_authenticationvserver_binding
func (s *AuthenticationService) GetAllAuthenticationCertPolicyAuthenticationVServerBinding() {}
func (s *AuthenticationService) GetAuthenticationCertPolicyAuthenticationVServerBinding()    {}
func (s *AuthenticationService) CountAuthenticationCertPolicyAuthenticationVServerBinding()  {}

// authenticationcertpolicy_binding
// Binding object which returns the resources bound to authenticationcertpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationcertpolicy_binding
func (s *AuthenticationService) GetAllAuthenticationCertPolicyBinding() {}
func (s *AuthenticationService) GetAuthenticationCertPolicyBinding()    {}

// authenticationcertpolicy_vpnglobal_binding
// Binding object showing the vpnglobal that can be bound to authenticationcertpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationcertpolicy_vpnglobal_binding
func (s *AuthenticationService) GetAllAuthenticationCertPolicyVPNGlobalBinding() {}
func (s *AuthenticationService) GetAuthenticationCertPolicyVPNGlobalBinding()    {}
func (s *AuthenticationService) CountAuthenticationCertPolicyVPNGlobalBinding()  {}

// authenticationcertpolicy_vpnvserver_binding
// Binding object showing the vpnvserver that can be bound to authenticationcertpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationcertpolicy_vpnvserver_binding
func (s *AuthenticationService) GetAllAuthenticationCertPolicyVPNVServerBinding() {}
func (s *AuthenticationService) GetAuthenticationCertPolicyVPNVServerBinding()    {}
func (s *AuthenticationService) CountAuthenticationCertPolicyVPNVServerBinding()  {}

// authenticationcitrixauthaction
// Configuration for Citrix Authentication action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationcitrixauthaction
func (s *AuthenticationService) AddAuthenticationCitrixAuthAction()    {}
func (s *AuthenticationService) DeleteAuthenticationCitrixAuthAction() {}
func (s *AuthenticationService) UpdateAuthenticationCitrixAuthAction() {}
func (s *AuthenticationService) UnsetAuthenticationCitrixAuthAction()  {}
func (s *AuthenticationService) GetAllAuthenticationCitrixAuthAction() {}
func (s *AuthenticationService) GetAuthenticationCitrixAuthAction()    {}
func (s *AuthenticationService) CountAuthenticationCitrixAuthAction()  {}

// authenticationdfaaction
// Configuration for Dfa authentication action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationdfaaction
func (s *AuthenticationService) AddAuthenticationDFAAction()    {}
func (s *AuthenticationService) DeleteAuthenticationDFAAction() {}
func (s *AuthenticationService) UpdateAuthenticationDFAAction() {}
func (s *AuthenticationService) UnsetAuthenticationDFAAction()  {}
func (s *AuthenticationService) GetAllAuthenticationDFAAction() {}
func (s *AuthenticationService) GetAuthenticationDFAAction()    {}
func (s *AuthenticationService) CountAuthenticationDFAAction()  {}

// authenticationdfapolicy
// Configuration for Dfa authentication policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationdfapolicy
func (s *AuthenticationService) AddAuthenticationDFAPolicy()    {}
func (s *AuthenticationService) DeleteAuthenticationDFAPolicy() {}
func (s *AuthenticationService) UpdateAuthenticationDFAPolicy() {}
func (s *AuthenticationService) GetAllAuthenticationDFAPolicy() {}
func (s *AuthenticationService) GetAuthenticationDFAPolicy()    {}
func (s *AuthenticationService) CountAuthenticationDFAPolicy()  {}

// authenticationdfapolicy_binding
// Binding object which returns the resources bound to authenticationdfapolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationdfapolicy_binding
func (s *AuthenticationService) GetAllAuthenticationDFAPolicyBinding() {}
func (s *AuthenticationService) GetAuthenticationDFAPolicyBinding()    {}

// authenticationdfapolicy_vpnvserver_binding
// Binding object showing the vpnvserver that can be bound to authenticationdfapolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationdfapolicy_vpnvserver_binding
func (s *AuthenticationService) GetAllAuthenticationDFAPolicyVPNVServerBinding() {}
func (s *AuthenticationService) GetAuthenticationDFAPolicyVPNVServerBinding()    {}
func (s *AuthenticationService) CountAuthenticationDFAPolicyVPNVServerBinding()  {}

// authenticationemailaction
// Configuration for Email entity resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationemailaction
func (s *AuthenticationService) AddAuthenticationEmailAction()    {}
func (s *AuthenticationService) DeleteAuthenticationEmailAction() {}
func (s *AuthenticationService) UpdateAuthenticationEmailAction() {}
func (s *AuthenticationService) UnsetAuthenticationEmailAction()  {}
func (s *AuthenticationService) GetAllAuthenticationEmailAction() {}
func (s *AuthenticationService) GetAuthenticationEmailAction()    {}
func (s *AuthenticationService) CountAuthenticationEmailAction()  {}

// authenticationepaaction
// Configuration for epa action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationepaaction
func (s *AuthenticationService) AddAuthenticationEPAAction()    {}
func (s *AuthenticationService) DeleteAuthenticationEPAAction() {}
func (s *AuthenticationService) UpdateAuthenticationEPAAction() {}
func (s *AuthenticationService) UnsetAuthenticationEPAAction()  {}
func (s *AuthenticationService) GetAllAuthenticationEPAAction() {}
func (s *AuthenticationService) GetAuthenticationEPAAction()    {}
func (s *AuthenticationService) CountAuthenticationEPAAction()  {}

// authenticationldapaction
// Configuration for LDAP action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationldapaction
func (s *AuthenticationService) AddAuthenticationLDAPAction()    {}
func (s *AuthenticationService) DeleteAuthenticationLDAPAction() {}
func (s *AuthenticationService) UpdateAuthenticationLDAPAction() {}
func (s *AuthenticationService) UnsetAuthenticationLDAPAction()  {}
func (s *AuthenticationService) GetAllAuthenticationLDAPAction() {}
func (s *AuthenticationService) GetAuthenticationLDAPAction()    {}
func (s *AuthenticationService) CountAuthenticationLDAPAction()  {}

// authenticationldappolicy
// Configuration for LDAP policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationldappolicy
func (s *AuthenticationService) AddAuthenticationLDAPPolicy()    {}
func (s *AuthenticationService) DeleteAuthenticationLDAPPolicy() {}
func (s *AuthenticationService) UpdateAuthenticationLDAPPolicy() {}
func (s *AuthenticationService) UnsetAuthenticationLDAPPolicy()  {}
func (s *AuthenticationService) GetAllAuthenticationLDAPPolicy() {}
func (s *AuthenticationService) GetAuthenticationLDAPPolicy()    {}
func (s *AuthenticationService) CountAuthenticationLDAPPolicy()  {}

// authenticationldappolicy_authenticationvserver_binding
// Binding object showing the authenticationvserver that can be bound to authenticationldappolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationldappolicy_authenticationvserver_binding
func (s *AuthenticationService) GetAllAuthenticationLDAPPolicyAuthenticationVServerBinding() {}
func (s *AuthenticationService) GetAuthenticationLDAPPolicyAuthenticationVServerBinding()    {}
func (s *AuthenticationService) CountAuthenticationLDAPPolicyAuthenticationVServerBinding()  {}

// authenticationldappolicy_binding
// Binding object which returns the resources bound to authenticationldappolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationldappolicy_binding
func (s *AuthenticationService) GetAllAuthenticationLDAPPolicyBinding() {}
func (s *AuthenticationService) GetAuthenticationLDAPPolicyBinding()    {}

// authenticationldappolicy_systemglobal_binding
// Binding object showing the systemglobal that can be bound to authenticationldappolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationldappolicy_systemglobal_binding
func (s *AuthenticationService) GetAllAuthenticationLDAPPolicySystemGlobalBinding() {}
func (s *AuthenticationService) GetAuthenticationLDAPPolicySystemGlobalBinding()    {}
func (s *AuthenticationService) CountAuthenticationLDAPPolicySystemGlobalBinding()  {}

// authenticationldappolicy_vpnglobal_binding
// Binding object showing the vpnglobal that can be bound to authenticationldappolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationldappolicy_vpnglobal_binding
func (s *AuthenticationService) GetAllAuthenticationLDAPPolicyVPNGlobalBinding() {}
func (s *AuthenticationService) GetAuthenticationLDAPPolicyVPNGlobalBinding()    {}
func (s *AuthenticationService) CountAuthenticationLDAPPolicyVPNGlobalBinding()  {}

// authenticationldappolicy_vpnvserver_binding
// Binding object showing the vpnvserver that can be bound to authenticationldappolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationldappolicy_vpnvserver_binding
func (s *AuthenticationService) GetAllAuthenticationLDAPPolicyVPNVServerBinding() {}
func (s *AuthenticationService) GetAuthenticationLDAPPolicyVPNVServerBinding()    {}
func (s *AuthenticationService) CountAuthenticationLDAPPolicyVPNVServerBinding()  {}

// authenticationlocalpolicy
// Configuration for LOCAL policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationlocalpolicy
func (s *AuthenticationService) AddAuthenticationLocalPolicy()    {}
func (s *AuthenticationService) DeleteAuthenticationLocalPolicy() {}
func (s *AuthenticationService) UpdateAuthenticationLocalPolicy() {}
func (s *AuthenticationService) GetAllAuthenticationLocalPolicy() {}
func (s *AuthenticationService) GetAuthenticationLocalPolicy()    {}
func (s *AuthenticationService) CountAuthenticationLocalPolicy()  {}

// authenticationlocalpolicy_authenticationvserver_binding
// Binding object showing the authenticationvserver that can be bound to authenticationlocalpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationlocalpolicy_authenticationvserver_binding
func (s *AuthenticationService) GetAllAuthenticationLocalPolicyAuthenticationVServerBinding() {}
func (s *AuthenticationService) GetAuthenticationLocalPolicyAuthenticationVServerBinding()    {}
func (s *AuthenticationService) CountAuthenticationLocalPolicyAuthenticationVServerBinding()  {}

// authenticationlocalpolicy_binding
// Binding object which returns the resources bound to authenticationlocalpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationlocalpolicy_binding
func (s *AuthenticationService) GetAllAuthenticationLocalPolicyBinding() {}
func (s *AuthenticationService) GetAuthenticationLocalPolicyBinding()    {}

// authenticationlocalpolicy_systemglobal_binding
// Binding object showing the systemglobal that can be bound to authenticationlocalpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationlocalpolicy_systemglobal_binding
func (s *AuthenticationService) GetAllAuthenticationLocalPolicySystemGlobalBinding() {}
func (s *AuthenticationService) GetAuthenticationLocalPolicySystemGlobalBinding()    {}
func (s *AuthenticationService) CountAuthenticationLocalPolicySystemGlobalBinding()  {}

// authenticationlocalpolicy_vpnglobal_binding
// Binding object showing the vpnglobal that can be bound to authenticationlocalpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationlocalpolicy_vpnglobal_binding
func (s *AuthenticationService) GetAllAuthenticationLocalPolicyVPNGlobalBinding() {}
func (s *AuthenticationService) GetAuthenticationLocalPolicyVPNGlobalBinding()    {}
func (s *AuthenticationService) CountAuthenticationLocalPolicyVPNGlobalBinding()  {}

// authenticationlocalpolicy_vpnvserver_binding
// Binding object showing the vpnvserver that can be bound to authenticationlocalpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationlocalpolicy_vpnvserver_binding
func (s *AuthenticationService) GetAllAuthenticationLocalPolicyVPNVServerBinding() {}
func (s *AuthenticationService) GetAuthenticationLocalPolicyVPNVServerBinding()    {}
func (s *AuthenticationService) CountAuthenticationLocalPolicyVPNVServerBinding()  {}

// authenticationloginschema
// Configuration for Login Schema resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationloginschema
func (s *AuthenticationService) AddAuthenticationLoginSchema()    {}
func (s *AuthenticationService) DeleteAuthenticationLoginSchema() {}
func (s *AuthenticationService) UpdateAuthenticationLoginSchema() {}
func (s *AuthenticationService) UnsetAuthenticationLoginSchema()  {}
func (s *AuthenticationService) GetAllAuthenticationLoginSchema() {}
func (s *AuthenticationService) GetAuthenticationLoginSchema()    {}
func (s *AuthenticationService) CountAuthenticationLoginSchema()  {}

// authenticationloginschemapolicy
// Configuration for Login Schema Policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationloginschemapolicy
func (s *AuthenticationService) AddAuthenticationLoginSchemaPolicy()    {}
func (s *AuthenticationService) DeleteAuthenticationLoginSchemaPolicy() {}
func (s *AuthenticationService) UpdateAuthenticationLoginSchemaPolicy() {}
func (s *AuthenticationService) UnsetAuthenticationLoginSchemaPolicy()  {}
func (s *AuthenticationService) GetAllAuthenticationLoginSchemaPolicy() {}
func (s *AuthenticationService) GetAuthenticationLoginSchemaPolicy()    {}
func (s *AuthenticationService) CountAuthenticationLoginSchemaPolicy()  {}
func (s *AuthenticationService) RenameAuthenticationLoginSchemaPolicy() {}

// authenticationloginschemapolicy_authenticationvserver_binding
// Binding object showing the authenticationvserver that can be bound to authenticationloginschemapolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationloginschemapolicy_authenticationvserver_binding
func (s *AuthenticationService) GetAllAuthenticationLoginSchemaPolicyAuthenticationVServerBinding() {}
func (s *AuthenticationService) GetAuthenticationLoginSchemaPolicyAuthenticationVServerBinding()    {}
func (s *AuthenticationService) CountAuthenticationLoginSchemaPolicyAuthenticationVServerBinding()  {}

// authenticationloginschemapolicy_binding
// Binding object which returns the resources bound to authenticationloginschemapolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationloginschemapolicy_binding
func (s *AuthenticationService) GetAllAuthenticationLoginSchemaPolicyBinding() {}
func (s *AuthenticationService) GetAuthenticationLoginSchemaPolicyBinding()    {}

// authenticationloginschemapolicy_vpnvserver_binding
// Binding object showing the vpnvserver that can be bound to authenticationloginschemapolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationloginschemapolicy_vpnvserver_binding
func (s *AuthenticationService) GetAllAuthenticationLoginSchemaPolicyVPNVServerBinding() {}
func (s *AuthenticationService) GetAuthenticationLoginSchemaPolicyVPNVServerBinding()    {}
func (s *AuthenticationService) CountAuthenticationLoginSchemaPolicyVPNVServerBinding()  {}

// authenticationnegotiataction
// Configuration for Negotiate action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationnegotiateaction
func (s *AuthenticationService) AddAuthenticationNegotiateAction()    {}
func (s *AuthenticationService) DeleteAuthenticationNegotiateAction() {}
func (s *AuthenticationService) UpdateAuthenticationNegotiateAction() {}
func (s *AuthenticationService) UnsetAuthenticationNegotiateAction()  {}
func (s *AuthenticationService) GetAllAuthenticationNegotiateAction() {}
func (s *AuthenticationService) GetAuthenticationNegotiateAction()    {}
func (s *AuthenticationService) CountAuthenticationNegotiateAction()  {}

// authenticationnegotiatpolicy
// Configuration for Negotiate Policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationnegotiatepolicy
func (s *AuthenticationService) AddAuthenticationNegotiatePolicy()    {}
func (s *AuthenticationService) DeleteAuthenticationNegotiatePolicy() {}
func (s *AuthenticationService) UpdateAuthenticationNegotiatePolicy() {}
func (s *AuthenticationService) UnsetAuthenticationNegotiatePolicy()  {}
func (s *AuthenticationService) GetAllAuthenticationNegotiatePolicy() {}
func (s *AuthenticationService) GetAuthenticationNegotiatePolicy()    {}
func (s *AuthenticationService) CountAuthenticationNegotiatePolicy()  {}

// authenticationnegotiatepolicy_authenticationvserver_binding
// Binding object showing the authenticationvserver that can be bound to authenticationnegotiatepolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationnegotiatepolicy_authenticationvserver_binding
func (s *AuthenticationService) GetAllAuthenticationNegotiatePolicyAuthenticationVServerBinding() {}
func (s *AuthenticationService) GetAuthenticationNegotiatePolicyAuthenticationVServerBinding()    {}
func (s *AuthenticationService) CountAuthenticationNegotiatePolicyAuthenticationVServerBinding()  {}

// authenticationnegotiatepolicy_binding
// Binding object which returns the resources bound to authenticationnegotiatepolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationnegotiatepolicy_binding
func (s *AuthenticationService) GetAllAuthenticationNegotiatePolicyBinding() {}
func (s *AuthenticationService) GetAuthenticationNegotiatePolicyBinding()    {}

// authenticationnegotiatepolicy_vpnglobal_binding
// Binding object showing the vpnglobal that can be bound to authenticationnegotiatepolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationnegotiatepolicy_vpnglobal_binding
func (s *AuthenticationService) GetAllAuthenticationNegotiatePolicyVPNGlobalBinding() {}
func (s *AuthenticationService) GetAuthenticationNegotiatePolicyVPNGlobalBinding()    {}
func (s *AuthenticationService) CountAuthenticationNegotiatePolicyVPNGlobalBinding()  {}

// authenticationnegotiatepolicy_vpnvserver_binding
// Binding object showing the vpnvserver that can be bound to authenticationnegotiatepolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationnegotiatepolicy_vpnvserver_binding
func (s *AuthenticationService) GetAllAuthenticationNegotiatePolicyVPNVServerBinding() {}
func (s *AuthenticationService) GetAuthenticationNegotiatePolicyVPNVServerBinding()    {}
func (s *AuthenticationService) CountAuthenticationNegotiatePolicyVPNVServerBinding()  {}

// authenticationnoauthaction
// Configuration for no authentication action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationnoauthaction
func (s *AuthenticationService) AddAuthenticationNoAuthAction()    {}
func (s *AuthenticationService) DeleteAuthenticationNoAuthAction() {}
func (s *AuthenticationService) UpdateAuthenticationNoAuthAction() {}
func (s *AuthenticationService) UnsetAuthenticationNoAuthAction()  {}
func (s *AuthenticationService) GetAllAuthenticationNoAuthAction() {}
func (s *AuthenticationService) GetAuthenticationNoAuthAction()    {}
func (s *AuthenticationService) CountAuthenticationNoAuthAction()  {}

// authenticationoauthaction
// Configuration for OAuth authentication action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationoauthaction
func (s *AuthenticationService) AddAuthenticationOAuthAction()    {}
func (s *AuthenticationService) DeleteAuthenticationOAuthAction() {}
func (s *AuthenticationService) UpdateAuthenticationOAuthAction() {}
func (s *AuthenticationService) UnsetAuthenticationOAuthAction()  {}
func (s *AuthenticationService) GetAllAuthenticationOAuthAction() {}
func (s *AuthenticationService) GetAuthenticationOAuthAction()    {}
func (s *AuthenticationService) CountAuthenticationOAuthAction()  {}

// authenticationoauthidppolicy
// Configuration for AAA OAuth IdentityProvider (IdP) policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationoauthidppolicy
func (s *AuthenticationService) AddAuthenticationOAuthIdPPolicy()    {}
func (s *AuthenticationService) DeleteAuthenticationOAuthIdPPolicy() {}
func (s *AuthenticationService) UpdateAuthenticationOAuthIdPPolicy() {}
func (s *AuthenticationService) UnsetAuthenticationOAuthIdPPolicy()  {}
func (s *AuthenticationService) GetAllAuthenticationOAuthIdPPolicy() {}
func (s *AuthenticationService) GetAuthenticationOAuthIdPPolicy()    {}
func (s *AuthenticationService) CountAuthenticationOAuthIdPPolicy()  {}
func (s *AuthenticationService) RenameAuthenticationOAuthIdPPolicy() {}

// authenticationoauthidppolicy_authenticationvserver_binding
// Binding object showing the authenticationvserver that can be bound to authenticationoauthidppolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationoauthidppolicy_authenticationvserver_binding
func (s *AuthenticationService) GetAllAuthenticationOAuthIdPPolicyAuthenticationVServerBinding() {}
func (s *AuthenticationService) GetAuthenticationOAuthIdPPolicyAuthenticationVServerBinding()    {}
func (s *AuthenticationService) CountAuthenticationOAuthIdPPolicyAuthenticationVServerBinding()  {}

// authenticationoauthidppolicy_binding
// Binding object which returns the resources bound to authenticationoauthidppolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationoauthidppolicy_binding
func (s *AuthenticationService) GetAllAuthenticationOAuthIdPPolicyBinding() {}
func (s *AuthenticationService) GetAuthenticationOAuthIdPPolicyBinding()    {}

// authenticationoauthidppolicy_vpnvserver_binding
// Binding object showing the vpnvserver that can be bound to authenticationoauthidppolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationoauthidppolicy_vpnvserver_binding
func (s *AuthenticationService) GetAllAuthenticationOAuthIdPPolicyVPNVServerBinding() {}
func (s *AuthenticationService) GetAuthenticationOAuthIdPPolicyVPNVServerBinding()    {}
func (s *AuthenticationService) CountAuthenticationOAuthIdPPolicyVPNVServerBinding()  {}

// authenticationoauthidpprofile
// Configuration for OAuth Identity Provider (IdP) profile resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationoauthidpprofile
func (s *AuthenticationService) AddAuthenticationOAuthIdPProfile()    {}
func (s *AuthenticationService) DeleteAuthenticationOAuthIdPProfile() {}
func (s *AuthenticationService) UpdateAuthenticationOAuthIdPProfile() {}
func (s *AuthenticationService) UnsetAuthenticationOAuthIdPProfile()  {}
func (s *AuthenticationService) GetAllAuthenticationOAuthIdPProfile() {}
func (s *AuthenticationService) GetAuthenticationOAuthIdPProfile()    {}
func (s *AuthenticationService) CountAuthenticationOAuthIdPProfile()  {}

// authenticationpolicy
// Configuration for Authentication Policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationpolicy
func (s *AuthenticationService) AddAuthenticationPolicy()    {}
func (s *AuthenticationService) DeleteAuthenticationPolicy() {}
func (s *AuthenticationService) UpdateAuthenticationPolicy() {}
func (s *AuthenticationService) UnsetAuthenticationPolicy()  {}
func (s *AuthenticationService) GetAllAuthenticationPolicy() {}
func (s *AuthenticationService) GetAuthenticationPolicy()    {}
func (s *AuthenticationService) CountAuthenticationPolicy()  {}
func (s *AuthenticationService) RenameAuthenticationPolicy() {}

// authenticationpolicylabel
// Configuration for authentication policy label resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationpolicylabel
func (s *AuthenticationService) AddAuthenticationPolicyLabel()    {}
func (s *AuthenticationService) DeleteAuthenticationPolicyLabel() {}
func (s *AuthenticationService) GetAllAuthenticationPolicyLabel() {}
func (s *AuthenticationService) GetAuthenticationPolicyLabel()    {}
func (s *AuthenticationService) CountAuthenticationPolicyLabel()  {}
func (s *AuthenticationService) RenameAuthenticationPolicyLabel() {}

// authenticationpolicylabel_authenticationpolicy_binding
// Binding object showing the authenticationpolicy that can be bound to authenticationpolicylabel.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationpolicylabel_authenticationpolicy_binding
func (s *AuthenticationService) AddAuthenticationPolicyLabelAuthenticationPolicyBinding()    {}
func (s *AuthenticationService) DeleteAuthenticationPolicyLabelAuthenticationPolicyBinding() {}
func (s *AuthenticationService) GetAllAuthenticationPolicyLabelAuthenticationPolicyBinding() {}
func (s *AuthenticationService) GetAuthenticationPolicyLabelAuthenticationPolicyBinding()    {}
func (s *AuthenticationService) CountAuthenticationPolicyLabelAuthenticationPolicyBinding()  {}

// authenticaitonpolicylabel_binding
// Binding object which returns the resources bound to authenticationpolicylabel.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationpolicylabel_binding
func (s *AuthenticationService) GetAllAuthenticationPolicyLabelBinding() {}
func (s *AuthenticationService) GetAuthenticationPolicyLabelBinding()    {}

// authenticationpolicy_authenticationpolicylabel_binding
// Binding object showing the authenticationpolicylabel that can be bound to authenticationpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationpolicy_authenticationpolicylabel_binding
func (s *AuthenticationService) GetAllAuthenticationPolicyAuthenticationPolicyLabelBinding() {}
func (s *AuthenticationService) GetAuthenticationPolicyAuthenticationPolicyLabelBinding()    {}
func (s *AuthenticationService) CountAuthenticationPolicyAuthenticationPolicyLabelBinding()  {}

// authenticationpolicy_authenticationvserver_binding
// Binding object showing the authenticationvserver that can be bound to authenticationpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationpolicy_authenticationvserver_binding
func (s *AuthenticationService) GetAllAuthenticationPolicyAuthenticationVServerBinding() {}
func (s *AuthenticationService) GetAuthenticationPolicyAuthenticationVServerBinding()    {}
func (s *AuthenticationService) CountAuthenticationPolicyAuthenticationVServerBinding()  {}

// authenticationpolicy_binding
// Binding object which returns the resources bound to authenticationpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationpolicy_binding
func (s *AuthenticationService) GetAllAuthenticationPolicyBinding() {}
func (s *AuthenticationService) GetAuthenticationPolicyBinding()    {}

// authenicationpolicy_systemglobal_binding
// Binding object showing the systemglobal that can be bound to authenticationpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationpolicy_systemglobal_binding
func (s *AuthenticationService) GetAllAuthenticationPolicySystemGlobalBinding() {}
func (s *AuthenticationService) GetAuthenticationPolicySystemGlobalBinding()    {}
func (s *AuthenticationService) CountAuthenticationPolicySystemGlobalBinding()  {}

// authenticationpushservice
// Configuration for Service details for sending push notifications resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationpushservice
func (s *AuthenticationService) AddAuthenticationPushService()    {}
func (s *AuthenticationService) DeleteAuthenticationPushService() {}
func (s *AuthenticationService) UpdateAuthenticationPushService() {}
func (s *AuthenticationService) UnsetAuthenticationPushService()  {}
func (s *AuthenticationService) GetAllAuthenticationPushService() {}
func (s *AuthenticationService) GetAuthenticationPushService()    {}
func (s *AuthenticationService) CountAuthenticationPushService()  {}

// authenticationradiusaction
// Configuration for RADIUS action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationradiusaction
func (s *AuthenticationService) AddAuthenticationRADIUSAction()    {}
func (s *AuthenticationService) DeleteAuthenticationRADIUSAction() {}
func (s *AuthenticationService) UpdateAuthenticationRADIUSAction() {}
func (s *AuthenticationService) UnsetAuthenticationRADIUSAction()  {}
func (s *AuthenticationService) GetAllAuthenticationRADIUSAction() {}
func (s *AuthenticationService) GetAuthenticationRADIUSAction()    {}
func (s *AuthenticationService) CountAuthenticationRADIUSAction()  {}

// authenticationradiuspolicy
// Configuration for RADIUS policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationradiuspolicy
func (s *AuthenticationService) AddAuthenticationRADIUSPolicy()    {}
func (s *AuthenticationService) DeleteAuthenticationRADIUSPolicy() {}
func (s *AuthenticationService) UpdateAuthenticationRADIUSPolicy() {}
func (s *AuthenticationService) UnsetAuthenticationRADIUSPolicy()  {}
func (s *AuthenticationService) GetAllAuthenticationRADIUSPolicy() {}
func (s *AuthenticationService) GetAuthenticationRADIUSPolicy()    {}
func (s *AuthenticationService) CountAuthenticationRADIUSPolicy()  {}

// authenticationradiuspolicy_authenticationvserver_binding
// Binding object showing the authenticationvserver that can be bound to authenticationradiuspolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationradiuspolicy_authenticationvserver_binding
func (s *AuthenticationService) GetAllAuthenticationRADIUSPolicyAuthenticationVServerBinding() {}
func (s *AuthenticationService) GetAuthenticationRADIUSPolicyAuthenticationVServerBinding()    {}
func (s *AuthenticationService) CountAuthenticationRADIUSPolicyAuthenticationVServerBinding()  {}

// authenticationradiuspolicy_binding
// Binding object which returns the resources bound to authenticationradiuspolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationradiuspolicy_binding
func (s *AuthenticationService) GetAllAuthenticationRADIUSPolicyBinding() {}
func (s *AuthenticationService) GetAuthenticationRADIUSPolicyBinding()    {}

// authenticationradiuspolicy_systemglobal_binding
// Binding object showing the systemglobal that can be bound to authenticationradiuspolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationradiuspolicy_systemglobal_binding
func (s *AuthenticationService) GetAllAuthenticationRADIUSPolicySystemGlobalBinding() {}
func (s *AuthenticationService) GetAuthenticationRADIUSPolicySystemGlobalBinding()    {}
func (s *AuthenticationService) CountAuthenticationRADIUSPolicySystemGlobalBinding()  {}

// authenticationradiuspolicy_vpnglobal_binding
// Binding object showing the vpnglobal that can be bound to authenticationradiuspolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationradiuspolicy_vpnglobal_binding
func (s *AuthenticationService) GetAllAuthenticationRADIUSPolicyVPNGlobalBinding() {}
func (s *AuthenticationService) GetAuthenticationRADIUSPolicyVPNGlobalBinding()    {}
func (s *AuthenticationService) CountAuthenticationRADIUSPolicyVPNGlobalBinding()  {}

// authenticationradiuspolicy_vpnvserver_binding
// Binding object showing the vpnvserver that can be bound to authenticationradiuspolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationradiuspolicy_vpnvserver_binding
func (s *AuthenticationService) GetAllAuthenticationRADIUSPolicyVPNVServerBinding() {}
func (s *AuthenticationService) GetAuthenticationRADIUSPolicyVPNVServerBinding()    {}
func (s *AuthenticationService) CountAuthenticationRADIUSPolicyVPNVServerBinding()  {}

// authenticationsamlaction
// Configuration for AAA Saml action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationsamlaction
func (s *AuthenticationService) AddAuthenticationSAMLAction()    {}
func (s *AuthenticationService) DeleteAuthenticationSAMLAction() {}
func (s *AuthenticationService) UpdateAuthenticationSAMLAction() {}
func (s *AuthenticationService) UnsetAuthenticationSAMLAction()  {}
func (s *AuthenticationService) GetAllAuthenticationSAMLAction() {}
func (s *AuthenticationService) GetAuthenticationSAMLAction()    {}
func (s *AuthenticationService) CountAuthenticationSAMLAction()  {}

// authenticationsamlidppolicy
// Configuration for AAA Saml IdentityProvider (IdP) policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationsamlidppolicy
func (s *AuthenticationService) AddAuthenticationSAMLIdPPolicy()    {}
func (s *AuthenticationService) DeleteAuthenticationSAMLIdPPolicy() {}
func (s *AuthenticationService) UpdateAuthenticationSAMLIdPPolicy() {}
func (s *AuthenticationService) UnsetAuthenticationSAMLIdPPolicy()  {}
func (s *AuthenticationService) GetAllAuthenticationSAMLIdPPolicy() {}
func (s *AuthenticationService) GetAuthenticationSAMLIdPPolicy()    {}
func (s *AuthenticationService) CountAuthenticationSAMLIdPPolicy()  {}
func (s *AuthenticationService) RenameAuthenticationSAMLIdPPolicy() {}

// authenticationsamlidppolicy_authenticationvserver_binding
// Binding object showing the authenticationvserver that can be bound to authenticationsamlidppolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationsamlidppolicy_authenticationvserver_binding
func (s *AuthenticationService) GetAllAuthenticationSAMLIdPPolicyAuthenticationVServerBinding() {}
func (s *AuthenticationService) GetAuthenticationSAMLIdPPolicyAuthenticationVServerBinding()    {}
func (s *AuthenticationService) CountAuthenticationSAMLIdPPolicyAuthenticationVServerBinding()  {}

// authenticationsamlidppolicy_binding
// Binding object which returns the resources bound to authenticationsamlidppolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationsamlidppolicy_binding
func (s *AuthenticationService) GetAllAuthenticationSAMLIdPPolicyBinding() {}
func (s *AuthenticationService) GetAuthenticationSAMLIdPPolicyBinding()    {}

// authenticationsamlidppolicy_vpnvserver_binding
// Binding object showing the vpnvserver that can be bound to authenticationsamlidppolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationsamlidppolicy_vpnvserver_binding
func (s *AuthenticationService) GetAllAuthenticationSAMLIdPPolicyVPNVServerBinding() {}
func (s *AuthenticationService) GetAuthenticationSAMLIdPPolicyVPNVServerBinding()    {}
func (s *AuthenticationService) CountAuthenticationSAMLIdPPolicyVPNVServerBinding()  {}

// authenticationsamlidpprofile
// Configuration for AAA Saml IdentityProvider (IdP) profile resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationsamlidpprofile
func (s *AuthenticationService) AddAuthenticationSAMLIdPProfile()    {}
func (s *AuthenticationService) DeleteAuthenticationSAMLIdPProfile() {}
func (s *AuthenticationService) UpdateAuthenticationSAMLIdPProfile() {}
func (s *AuthenticationService) UnsetAuthenticationSAMLIdPProfile()  {}
func (s *AuthenticationService) GetAllAuthenticationSAMLIdPProfile() {}
func (s *AuthenticationService) GetAuthenticationSAMLIdPProfile()    {}
func (s *AuthenticationService) CountAuthenticationSAMLIdPProfile()  {}

// authenticaitonsamlpolicy
// Configuration for AAA Saml policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationsamlpolicy
func (s *AuthenticationService) AddAuthenticationSAMLPolicy()    {}
func (s *AuthenticationService) DeleteAuthenticationSAMLPolicy() {}
func (s *AuthenticationService) UpdateAuthenticationSAMLPolicy() {}
func (s *AuthenticationService) UnsetAuthenticationSAMLPolicy()  {}
func (s *AuthenticationService) GetAllAuthenticationSAMLPolicy() {}
func (s *AuthenticationService) GetAuthenticationSAMLPolicy()    {}
func (s *AuthenticationService) CountAuthenticationSAMLPolicy()  {}

// authenticationsamlpolicy_authenticationvserver_binding
// Binding object showing the authenticationvserver that can be bound to authenticationsamlpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationsamlpolicy_authenticationvserver_binding
func (s *AuthenticationService) GetAllAuthenticationSAMLPolicyAuthenticationVServerBinding() {}
func (s *AuthenticationService) GetAuthenticationSAMLPolicyAuthenticationVServerBinding()    {}
func (s *AuthenticationService) CountAuthenticationSAMLPolicyAuthenticationVServerBinding()  {}

// authenticationsamlpolicy_binding
// Binding object which returns the resources bound to authenticationsamlpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationsamlpolicy_binding
func (s *AuthenticationService) GetAllAuthenticationSAMLPolicyBinding() {}
func (s *AuthenticationService) GetAuthenticationSAMLPolicyBinding()    {}

// authenticationsamlpolicy_vpnglobal_binding
// Binding object showing the vpnglobal that can be bound to authenticationsamlpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationsamlpolicy_vpnglobal_binding
func (s *AuthenticationService) GetAllAuthenticationSAMLPolicyVPNGlobalBinding() {}
func (s *AuthenticationService) GetAuthenticationSAMLPolicyVPNGlobalBinding()    {}
func (s *AuthenticationService) CountAuthenticationSAMLPolicyVPNGlobalBinding()  {}

// authenticationsamlpolicy_vpnvserver_binding
// Binding object showing the vpnvserver that can be bound to authenticationsamlpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationsamlpolicy_vpnvserver_binding
func (s *AuthenticationService) GetAllAuthenticationSAMLPolicyVPNVServerBinding() {}
func (s *AuthenticationService) GetAuthenticationSAMLPolicyVPNVServerBinding()    {}
func (s *AuthenticationService) CountAuthenticationSAMLPolicyVPNVServerBinding()  {}

// authenticationstorefrontauthaction
// Configuration for Storefront authentication action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationstorefrontauthaction
func (s *AuthenticationService) AddAuthenticationStoreFrontAuthAction()    {}
func (s *AuthenticationService) DeleteAuthenticationStoreFrontAuthAction() {}
func (s *AuthenticationService) UpdateAuthenticationStoreFrontAuthAction() {}
func (s *AuthenticationService) UnsetAuthenticationStoreFrontAuthAction()  {}
func (s *AuthenticationService) GetAllAuthenticationStoreFrontAuthAction() {}
func (s *AuthenticationService) GetAuthenticationStoreFrontAuthAction()    {}
func (s *AuthenticationService) CountAuthenticationStoreFrontAuthAction()  {}

// authenticationtacacsaction
// Configuration for TACACS action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationtacacsaction
func (s *AuthenticationService) AddAuthenticationTACACSAction()    {}
func (s *AuthenticationService) DeleteAuthenticationTACACSAction() {}
func (s *AuthenticationService) UpdateAuthenticationTACACSAction() {}
func (s *AuthenticationService) UnsetAuthenticationTACACSAction()  {}
func (s *AuthenticationService) GetAllAuthenticationTACACSAction() {}
func (s *AuthenticationService) GetAuthenticationTACACSAction()    {}
func (s *AuthenticationService) CountAuthenticationTACACSAction()  {}

// authenticationtacacspolicy
// Configuration for TACACS policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationtacacspolicy
func (s *AuthenticationService) AddAuthenticationTACACSPolicy()    {}
func (s *AuthenticationService) DeleteAuthenticationTACACSPolicy() {}
func (s *AuthenticationService) UpdateAuthenticationTACACSPolicy() {}
func (s *AuthenticationService) UnsetAuthenticationTACACSPolicy()  {}
func (s *AuthenticationService) GetAllAuthenticationTACACSPolicy() {}
func (s *AuthenticationService) GetAuthenticationTACACSPolicy()    {}
func (s *AuthenticationService) CountAuthenticationTACACSPolicy()  {}

// authenticationtacacspolicy_authenticationvserver_binding
// Binding object showing the authenticationvserver that can be bound to authenticationtacacspolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationtacacspolicy_authenticationvserver_binding
func (s *AuthenticationService) GetAllAuthenticationTACACSPolicyAuthenticationVServerBinding() {}
func (s *AuthenticationService) GetAuthenticationTACACSPolicyAuthenticationVServerBinding()    {}
func (s *AuthenticationService) CountAuthenticationTACACSPolicyAuthenticationVServerBinding()  {}

// authenticationtacacspolicy_binding
// Binding object which returns the resources bound to authenticationtacacspolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationtacacspolicy_binding
func (s *AuthenticationService) GetAllAuthenticationTACACSPolicyBinding() {}
func (s *AuthenticationService) GetAuthenticationTACACSPolicyBinding()    {}

// authenticationtacacspolicy_systemglobal_binding
// Binding object showing the systemglobal that can be bound to authenticationtacacspolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationtacacspolicy_systemglobal_binding
func (s *AuthenticationService) GetAllAuthenticationTACACSPolicySystemGlobalBinding() {}
func (s *AuthenticationService) GetAuthenticationTACACSPolicySystemGlobalBinding()    {}
func (s *AuthenticationService) CountAuthenticationTACACSPolicySystemGlobalBinding()  {}

// authenticationtacacspolicy_vpnglobal_binding
// Binding object showing the vpnglobal that can be bound to authenticationtacacspolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationtacacspolicy_vpnglobal_binding
func (s *AuthenticationService) GetAllAuthenticationTACACSPolicyVPNGlobalBinding() {}
func (s *AuthenticationService) GetAuthenticationTACACSPolicyVPNGlobalBinding()    {}
func (s *AuthenticationService) CountAuthenticationTACACSPolicyVPNGlobalBinding()  {}

// authenticationtacacspolicy_vpnvserver_binding
// Binding object showing the vpnvserver that can be bound to authenticationtacacspolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationtacacspolicy_vpnvserver_binding
func (s *AuthenticationService) GetAllAuthenticationTACACSPolicyVPNVServerBinding() {}
func (s *AuthenticationService) GetAuthenticationTACACSPolicyVPNVServerBinding()    {}
func (s *AuthenticationService) CountAuthenticationTACACSPolicyVPNVServerBinding()  {}

// authenticaitonvserver
// Configuration for authentication virtual server resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver
func (s *AuthenticationService) AddAuthenticationVServer()     {}
func (s *AuthenticationService) DeleteAuthenticationVServer()  {}
func (s *AuthenticationService) UpdateAuthenticationVServer()  {}
func (s *AuthenticationService) UnsetAuthenticationVServer()   {}
func (s *AuthenticationService) EnableAuthenticationVServer()  {}
func (s *AuthenticationService) DisableAuthenticationVServer() {}
func (s *AuthenticationService) GetAllAuthenticationVServer()  {}
func (s *AuthenticationService) GetAuthenticationVServer()     {}
func (s *AuthenticationService) CountAuthenticationVServer()   {}
func (s *AuthenticationService) RenameAuthenticationVServer()  {}

// authenticationvserver_auditnslogpolicy_binding
// Binding object showing the auditnslogpolicy that can be bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_auditnslogpolicy_binding
func (s *AuthenticationService) AddAuthenticationVServerAuditNSLogPolicyBinding()    {}
func (s *AuthenticationService) DeleteAuthenticationVServerAuditNSLogPolicyBinding() {}
func (s *AuthenticationService) GetAllAuthenticationVServerAuditNSLogPolicyBinding() {}
func (s *AuthenticationService) GetAuthenticationVServerAuditNSLogPolicyBinding()    {}
func (s *AuthenticationService) CountAuthenticationVServerAuditNSLogPolicyBinding()  {}

// authenticationvserver_auditsyslogpolicy_binding
// Binding object showing the auditsyslogpolicy that can be bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_auditsyslogpolicy_binding
func (s *AuthenticationService) AddAuthenticationVServerAuditSyslogPolicyBinding()    {}
func (s *AuthenticationService) DeleteAuthenticationVServerAuditSyslogPolicyBinding() {}
func (s *AuthenticationService) GetAllAuthenticationVServerAuditSyslogPolicyBinding() {}
func (s *AuthenticationService) GetAuthenticationVServerAuditSyslogPolicyBinding()    {}
func (s *AuthenticationService) CountAuthenticationVServerAuditSyslogPolicyBinding()  {}

// authenticationvserver_authenticationcertpolicy_binding
// Binding object showing the authenticationcertpolicy that can be bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_authenticationcertpolicy_binding
func (s *AuthenticationService) AddAuthenticationVServerAuthenticationCertPolicyBinding()    {}
func (s *AuthenticationService) DeleteAuthenticationVServerAuthenticationCertPolicyBinding() {}
func (s *AuthenticationService) GetAllAuthenticationVServerAuthenticationCertPolicyBinding() {}
func (s *AuthenticationService) GetAuthenticationVServerAuthenticationCertPolicyBinding()    {}
func (s *AuthenticationService) CountAuthenticationVServerAuthenticationCertPolicyBinding()  {}

// authenticationvserver_authenticationldappolicy_binding
// Binding object showing the authenticationldappolicy that can be bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_authenticationldappolicy_binding
func (s *AuthenticationService) AddAuthenticationVServerAuthenticationLDAPPolicyBinding()    {}
func (s *AuthenticationService) DeleteAuthenticationVServerAuthenticationLDAPPolicyBinding() {}
func (s *AuthenticationService) GetAllAuthenticationVServerAuthenticationLDAPPolicyBinding() {}
func (s *AuthenticationService) GetAuthenticationVServerAuthenticationLDAPPolicyBinding()    {}
func (s *AuthenticationService) CountAuthenticationVServerAuthenticationLDAPPolicyBinding()  {}

// authenticationvserver_authenticationlocalpolicy_binding
// Binding object showing the authenticationlocalpolicy that can be bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_authenticationlocalpolicy_binding
func (s *AuthenticationService) AddAuthenticationVServerAuthenticationLocalPolicyBinding()    {}
func (s *AuthenticationService) DeleteAuthenticationVServerAuthenticationLocalPolicyBinding() {}
func (s *AuthenticationService) GetAllAuthenticationVServerAuthenticationLocalPolicyBinding() {}
func (s *AuthenticationService) GetAuthenticationVServerAuthenticationLocalPolicyBinding()    {}
func (s *AuthenticationService) CountAuthenticationVServerAuthenticationLocalPolicyBinding()  {}

// authenticationvserver_authenticationloginschemapolicy_binding
// Binding object showing the authenticationloginschemapolicy that can be bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_authenticationloginschemapolicy_binding
func (s *AuthenticationService) AddAuthenticationVServerAuthenticationLoginSchemaPolicyBinding()    {}
func (s *AuthenticationService) DeleteAuthenticationVServerAuthenticationLoginSchemaPolicyBinding() {}
func (s *AuthenticationService) GetAllAuthenticationVServerAuthenticationLoginSchemaPolicyBinding() {}
func (s *AuthenticationService) GetAuthenticationVServerAuthenticationLoginSchemaPolicyBinding()    {}
func (s *AuthenticationService) CountAuthenticationVServerAuthenticationLoginSchemaPolicyBinding()  {}

// authenticationvserver_authenticationnegotiatepolicy_binding
// Binding object showing the authenticationnegotiatepolicy that can be bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_authenticationnegotiatepolicy_binding
func (s *AuthenticationService) AddAuthenticationVServerAuthenticationNegotiatePolicyBinding()    {}
func (s *AuthenticationService) DeleteAuthenticationVServerAuthenticationNegotiatePolicyBinding() {}
func (s *AuthenticationService) GetAllAuthenticationVServerAuthenticationNegotiatePolicyBinding() {}
func (s *AuthenticationService) GetAuthenticationVServerAuthenticationNegotiatePolicyBinding()    {}
func (s *AuthenticationService) CountAuthenticationVServerAuthenticationNegotiatePolicyBinding()  {}

// authenticationvserver_authenticationoauthidppolicy_binding
// Binding object showing the authenticationoauthidppolicy that can be bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_authenticationoauthidppolicy_binding
func (s *AuthenticationService) AddAuthenticationVServerAuthenticationOAuthIdPPolicyBinding()    {}
func (s *AuthenticationService) DeleteAuthenticationVServerAuthenticationOAuthIdPPolicyBinding() {}
func (s *AuthenticationService) GetAllAuthenticationVServerAuthenticationOAuthIdPPolicyBinding() {}
func (s *AuthenticationService) GetAuthenticationVServerAuthenticationOAuthIdPPolicyBinding()    {}
func (s *AuthenticationService) CountAuthenticationVServerAuthenticationOAuthIdPPolicyBinding()  {}

// authenticationvserver_authenticationpolicy_binding
// Binding object showing the authenticationpolicy that can be bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_authenticationpolicy_binding
func (s *AuthenticationService) AddAuthenticationVServerAuthenticationPolicyBinding()    {}
func (s *AuthenticationService) DeleteAuthenticationVServerAuthenticationPolicyBinding() {}
func (s *AuthenticationService) GetAllAuthenticationVServerAuthenticationPolicyBinding() {}
func (s *AuthenticationService) GetAuthenticationVServerAuthenticationPolicyBinding()    {}
func (s *AuthenticationService) CountAuthenticationVServerAuthenticationPolicyBinding()  {}

// authenticationvserver_authenticationradiuspolicy_binding
// Binding object showing the authenticationradiuspolicy that can be bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_authenticationradiuspolicy_binding
func (s *AuthenticationService) AddAuthenticationVServerAuthenticationRADIUSPolicyBinding()    {}
func (s *AuthenticationService) DeleteAuthenticationVServerAuthenticationRADIUSPolicyBinding() {}
func (s *AuthenticationService) GetAllAuthenticationVServerAuthenticationRADIUSPolicyBinding() {}
func (s *AuthenticationService) GetAuthenticationVServerAuthenticationRADIUSPolicyBinding()    {}
func (s *AuthenticationService) CountAuthenticationVServerAuthenticationRADIUSPolicyBinding()  {}

// authenticationvserver_authenticationsamlidppolicy_binding
// Binding object showing the authenticationsamlidppolicy that can be bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_authenticationsamlidppolicy_binding
func (s *AuthenticationService) AddAuthenticationVServerAuthenticationSAMLIdPPolicyBinding()    {}
func (s *AuthenticationService) DeleteAuthenticationVServerAuthenticationSAMLIdPPolicyBinding() {}
func (s *AuthenticationService) GetAllAuthenticationVServerAuthenticationSAMLIdPPolicyBinding() {}
func (s *AuthenticationService) GetAuthenticationVServerAuthenticationSAMLIdPPolicyBinding()    {}
func (s *AuthenticationService) CountAuthenticationVServerAuthenticationSAMLIdPPolicyBinding()  {}

// authenticationvserver_authenticationsamlpolicy_binding
// Binding object showing the authenticationsamlpolicy that can be bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_authenticationsamlpolicy_binding
func (s *AuthenticationService) AddAuthenticationVServerAuthenticationSAMLPolicyBinding()    {}
func (s *AuthenticationService) DeleteAuthenticationVServerAuthenticationSAMLPolicyBinding() {}
func (s *AuthenticationService) GetAllAuthenticationVServerAuthenticationSAMLPolicyBinding() {}
func (s *AuthenticationService) GetAuthenticationVServerAuthenticationSAMLPolicyBinding()    {}
func (s *AuthenticationService) CountAuthenticationVServerAuthenticationSAMLPolicyBinding()  {}

// authenticationvserver_authenticationtacacspolicy_binding
// Binding object showing the authenticationtacacspolicy that can be bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_authenticationtacacspolicy_binding
func (s *AuthenticationService) AddAuthenticationVServerAuthenticationTACACSPolicyBinding()    {}
func (s *AuthenticationService) DeleteAuthenticationVServerAuthenticationTACACSPolicyBinding() {}
func (s *AuthenticationService) GetAllAuthenticationVServerAuthenticationTACACSPolicyBinding() {}
func (s *AuthenticationService) GetAuthenticationVServerAuthenticationTACACSPolicyBinding()    {}
func (s *AuthenticationService) CountAuthenticationVServerAuthenticationTACACSPolicyBinding()  {}

// authenticationvserver_authenticationwebauthpolicy_binding
// Binding object showing the authenticationwebauthpolicy that can be bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_authenticationwebauthpolicy_binding
func (s *AuthenticationService) AddAuthenticationVServerAuthenticationWebAuthPolicyBinding()    {}
func (s *AuthenticationService) DeleteAuthenticationVServerAuthenticationWebAuthPolicyBinding() {}
func (s *AuthenticationService) GetAllAuthenticationVServerAuthenticationWebAuthPolicyBinding() {}
func (s *AuthenticationService) GetAuthenticationVServerAuthenticationWebAuthPolicyBinding()    {}
func (s *AuthenticationService) CountAuthenticationVServerAuthenticationWebAuthPolicyBinding()  {}

// authenticationvserver_binding
// Binding object which returns the resources bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_binding
func (s *AuthenticationService) GetAllAuthenticationVServerBinding() {}
func (s *AuthenticationService) GetAuthenticationVServerBinding()    {}

// authenticationvserver_cachepolicy_binding
// Binding object showing the cachepolicy that can be bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_cachepolicy_binding
func (s *AuthenticationService) AddAuthenticationVServerCachePolicyBinding()    {}
func (s *AuthenticationService) DeleteAuthenticationVServerCachePolicyBinding() {}
func (s *AuthenticationService) GetAllAuthenticationVServerCachePolicyBinding() {}
func (s *AuthenticationService) GetAuthenticationVServerCachePolicyBinding()    {}
func (s *AuthenticationService) CountAuthenticationVServerCachePolicyBinding()  {}

// authenticationvserver_cspolicy_binding
// Binding object showing the cspolicy that can be bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_cspolicy_binding
func (s *AuthenticationService) AddAuthenticationVServerCSPolicyBinding()    {}
func (s *AuthenticationService) DeleteAuthenticationVServerCSPolicyBinding() {}
func (s *AuthenticationService) GetAllAuthenticationVServerCSPolicyBinding() {}
func (s *AuthenticationService) GetAuthenticationVServerCSPolicyBinding()    {}
func (s *AuthenticationService) CountAuthenticationVServerCSPolicyBinding()  {}

// authenticationvserver_responderpolicy_binding
// Binding object showing the responderpolicy that can be bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_responderpolicy_binding
func (s *AuthenticationService) AddAuthenticationVServerResponderPolicyBinding()    {}
func (s *AuthenticationService) DeleteAuthenticationVServerResponderPolicyBinding() {}
func (s *AuthenticationService) GetAllAuthenticationVServerResponderPolicyBinding() {}
func (s *AuthenticationService) GetAuthenticationVServerResponderPolicyBinding()    {}
func (s *AuthenticationService) CountAuthenticationVServerResponderPolicyBinding()  {}

// authenticationvserver_tmsessionpolicy_binding
// Binding object showing the tmsessionpolicy that can be bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_tmsessionpolicy_binding
func (s *AuthenticationService) AddAuthenticationVServerTMSessionPolicyBinding()    {}
func (s *AuthenticationService) DeleteAuthenticationVServerTMSessionPolicyBinding() {}
func (s *AuthenticationService) GetAllAuthenticationVServerTMSessionPolicyBinding() {}
func (s *AuthenticationService) GetAuthenticationVServerTMSessionPolicyBinding()    {}
func (s *AuthenticationService) CountAuthenticationVServerTMSessionPolicyBinding()  {}

// authenticationvserver_vpnportaltheme_binding
// Binding object showing the vpnportaltheme that can be bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_vpnportaltheme_binding
func (s *AuthenticationService) AddAuthenticationVServerVPNPortalThemeBinding()    {}
func (s *AuthenticationService) DeleteAuthenticationVServerVPNPortalThemeBinding() {}
func (s *AuthenticationService) GetAllAuthenticationVServerVPNPortalThemeBinding() {}
func (s *AuthenticationService) GetAuthenticationVServerVPNPortalThemeBinding()    {}
func (s *AuthenticationService) CountAuthenticationVServerVPNPortalThemeBinding()  {}

// authenticationwebauthaction
// Configuration for Web authentication action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationwebauthaction
func (s *AuthenticationService) AddAuthenticationWebAuthAction()    {}
func (s *AuthenticationService) DeleteAuthenticationWebAuthAction() {}
func (s *AuthenticationService) UpdateAuthenticationWebAuthAction() {}
func (s *AuthenticationService) UnsetAuthenticationWebAuthAction()  {}
func (s *AuthenticationService) GetAllAuthenticationWebAuthAction() {}
func (s *AuthenticationService) GetAuthenticationWebAuthAction()    {}
func (s *AuthenticationService) CountAuthenticationWebAuthAction()  {}

// authenticationwebauthpolicy
// Configuration for Web authentication policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationwebauthpolicy
func (s *AuthenticationService) AddAuthenticationWebAuthPolicy()    {}
func (s *AuthenticationService) DeleteAuthenticationWebAuthPolicy() {}
func (s *AuthenticationService) UpdateAuthenticationWebAuthPolicy() {}
func (s *AuthenticationService) GetAllAuthenticationWebAuthPolicy() {}
func (s *AuthenticationService) GetAuthenticationWebAuthPolicy()    {}
func (s *AuthenticationService) CountAuthenticationWebAuthPolicy()  {}

// authenticationwebauthpolicy_authenticationvserver_binding
// Binding object showing the authenticationvserver that can be bound to authenticationwebauthpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationwebauthpolicy_authenticationvserver_binding
func (s *AuthenticationService) GetAllAuthenticationWebAuthPolicyAuthenticationVServerBinding() {}
func (s *AuthenticationService) GetAuthenticationWebAuthPolicyAuthenticationVServerBinding()    {}
func (s *AuthenticationService) CountAuthenticationWebAuthPolicyAuthenticationVServerBinding()  {}

// authenticationwebauthpolicy_binding
// Binding object which returns the resources bound to authenticationwebauthpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationwebauthpolicy_binding
func (s *AuthenticationService) GetAllAuthenticationWebAuthPolicyBinding() {}
func (s *AuthenticationService) GetAuthenticationWebAuthPolicyBinding()    {}

// authenticationwebauthpolicy_systemglobal_binding
// Binding object showing the systemglobal that can be bound to authenticationwebauthpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationwebauthpolicy_systemglobal_binding
func (s *AuthenticationService) GetAllAuthenticationWebAuthPolicySystemGlobalBinding() {}
func (s *AuthenticationService) GetAuthenticationWebAuthPolicySystemGlobalBinding()    {}
func (s *AuthenticationService) CountAuthenticationWebAuthPolicySystemGlobalBinding()  {}

// authenticationwebauthpolicy_vpnglobal_binding
// Binding object showing the vpnglobal that can be bound to authenticationwebauthpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationwebauthpolicy_vpnglobal_binding
func (s *AuthenticationService) GetAllAuthenticationWebAuthPolicyVPNGlobalBinding() {}
func (s *AuthenticationService) GetAuthenticationWebAuthPolicyVPNGlobalBinding()    {}
func (s *AuthenticationService) CountAuthenticationWebAuthPolicyVPNGlobalBinding()  {}

// authenticationwebauthpolicy_vpnvserver_binding
// Binding object showing the vpnvserver that can be bound to authenticationwebauthpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationwebauthpolicy_vpnvserver_binding
func (s *AuthenticationService) GetAllAuthenticationWebAuthPolicyVPNVServerBinding() {}
func (s *AuthenticationService) GetAuthenticationWebAuthPolicyVPNVServerBinding()    {}
func (s *AuthenticationService) CountAuthenticationWebAuthPolicyVPNVServerBinding()  {}
