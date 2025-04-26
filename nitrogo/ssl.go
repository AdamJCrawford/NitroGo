package nitrogo

// SSL
// SSL Configuration
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/ssl/ssl
type SSLService struct {
	client *Client
}

// sslaction
func (s *SSLService) AddSSLAction()    {}
func (s *SSLService) DeleteSSLAction() {}
func (s *SSLService) GetAllSSLAction() {}
func (s *SSLService) GetSSLAction()    {}
func (s *SSLService) CountSSLAction()  {}

// sslcacertgroup
func (s *SSLService) AddSSLCACertGroup()    {}
func (s *SSLService) DeleteSSLCACertGroup() {}
func (s *SSLService) GetAllSSLCACertGroup() {}
func (s *SSLService) GetSSLCACertGroup()    {}
func (s *SSLService) CountSSLCACertGroup()  {}

// sslcacertgroup_binding
func (s *SSLService) GetAllSSLCACertGroupBinding() {}
func (s *SSLService) GetSSLCACertGroupBinding()    {}

// sslcacertgroup_sslcertkey_binding
func (s *SSLService) AddSSLCACertGroupSSLCertKeyBinding()    {}
func (s *SSLService) DeleteSSLCACertGroupSSLCertKeyBinding() {}
func (s *SSLService) GetAllSSLCACertGroupSSLCertKeyBinding() {}
func (s *SSLService) GetSSLCACertGroupSSLCertKeyBinding()    {}
func (s *SSLService) CountSSLCACertGroupSSLCertKeyBinding()  {}

// sslcert
func (s *SSLService) CreateSSLCert() {}

// sslcertbundle
func (s *SSLService) ImportSSLCertBundle() {}
func (s *SSLService) DeleteSSLCertBundle() {}
func (s *SSLService) ApplySSLCertBundle()  {}
func (s *SSLService) ExportSSLCertBundle() {}
func (s *SSLService) GetAllSSLCertBundle() {}
func (s *SSLService) CountSSLCertBundle()  {}

// sslcertchain
func (s *SSLService) GetAllSSLCertChain() {}
func (s *SSLService) GetSSLCertChain()    {}
func (s *SSLService) CountSSLCertChain()  {}

// sslcertchain_binding
func (s *SSLService) GetSSLCertChainBinding() {}

// sslcertchain_sslcertkey_binding
func (s *SSLService) GetSSLCertChainSSLCertKeyBinding()   {}
func (s *SSLService) CountSSLCertChainSSLCertKeyBinding() {}

// sslcertfile
func (s *SSLService) ImportSSLCertFile() {}
func (s *SSLService) DeleteSSLCertFile() {}
func (s *SSLService) GetAllSSLCertFile() {}
func (s *SSLService) CountSSLCertFile()  {}

// sslcertificatechain
func (s *SSLService) AddSSLCertificateChain()    {}
func (s *SSLService) GetAllSSLCertificateChain() {}
func (s *SSLService) GetSSLCertificateChain()    {}
func (s *SSLService) CountSSLCertificateChain()  {}

// sslcertkey
func (s *SSLService) AddSSLCertKey()    {}
func (s *SSLService) DeleteSSLCertKey() {}
func (s *SSLService) UpdateSSLCertKey() {}
func (s *SSLService) UnsetSSLCertKey()  {}
func (s *SSLService) LinkSSLCertKey()   {}
func (s *SSLService) UnlinkSSLCertKey() {}
func (s *SSLService) ChangeSSLCertKey() {}
func (s *SSLService) ClearSSLCertKey()  {}
func (s *SSLService) GetAllSSLCertKey() {}
func (s *SSLService) GetSSLCertKey()    {}
func (s *SSLService) CountSSLCertKey()  {}

// sslcertkey_binding
func (s *SSLService) GetAllSSLCertKeyBinding() {}
func (s *SSLService) GetSSLCertKeyBinding()    {}

// sslcertkey_crldistribution_binding
func (s *SSLService) GetAllSSLCertKeyCRLDistributionBinding() {}
func (s *SSLService) GetSSLCertKeyCRLDistributionBinding()    {}
func (s *SSLService) CountSSLCertKeyCRLDistributionBinding()  {}

// sslcertkey_service_binding
func (s *SSLService) GetAllSSLCertKeyServiceBinding() {}
func (s *SSLService) GetSSLCertKeyServiceBinding()    {}
func (s *SSLService) CountSSLCertKeyServiceBinding()  {}

// sslcertkey_sslocspresponder_binding
func (s *SSLService) AddSSLCertKeySSLOCSPResponderBinding()    {}
func (s *SSLService) DeleteSSLCertKeySSLOCSPResponderBinding() {}
func (s *SSLService) GetAllSSLCertKeySSLOCSPResponderBinding() {}
func (s *SSLService) GetSSLCertKeySSLOCSPResponderBinding()    {}
func (s *SSLService) CountSSLCertKeySSLOCSPResponderBinding()  {}

// sslcertkey_sslprofile_binding
func (s *SSLService) GetAllSSLCertKeySSLProfileBinding() {}
func (s *SSLService) GetSSLCertKeySSLProfileBinding()    {}
func (s *SSLService) CountSSLCertKeySSLProfileBinding()  {}

// sslcertkey_sslvserver_binding
func (s *SSLService) GetAllSSLCertKeySSLVServerBinding() {}
func (s *SSLService) GetSSLCertKeySSLVServerBinding()    {}
func (s *SSLService) CountSSLCertKeySSLVServerBinding()  {}

// sslcertlink
func (s *SSLService) GetAllSSLCertLink() {}
func (s *SSLService) CountSSLCertLink()  {}

// sslcertreq
func (s *SSLService) CreateSSLCertReq() {}

// sslcipher
func (s *SSLService) AddSSLCipher()    {}
func (s *SSLService) DeleteSSLCipher() {}
func (s *SSLService) UpdateSSLCipher() {}
func (s *SSLService) UnsetSSLCipher()  {}
func (s *SSLService) GetAllSSLCipher() {}
func (s *SSLService) GetSSLCipher()    {}
func (s *SSLService) CountSSLCipher()  {}

// sslciphersuite
func (s *SSLService) GetAllSSLCipherSuite() {}
func (s *SSLService) GetSSLCipherSuite()    {}
func (s *SSLService) CountSSLCipherSuite()  {}

// sslcipher_binding
func (s *SSLService) GetAllSSLCipherBinding() {}
func (s *SSLService) GetSSLCipherBinding()    {}

// sslcipher_individualcipher_binding
func (s *SSLService) GetAllSSLCipherIndividualCipherBinding() {}
func (s *SSLService) GetSSLCipherIndividualCipherBinding()    {}
func (s *SSLService) CountSSLCipherIndividualCipherBinding()  {}

// sslcipher_servicegroup_binding
// This object has no operations documented. These are guesses and need to be tested before they will be uncommented.
// func (s *SSLService) GetAllSSLCipherServiceGroupBinding() {}
// func (s *SSLService) GetSSLCipherServiceGroupBinding()    {}
// func (s *SSLService) CountSSLCipherServiceGroupBinding()  {}

// sslcipher_service_binding
// This object has no operations documented. These are guesses and need to be tested before they will be uncommented.
// func (s *SSLService) GetAllSSLCipherServiceBinding() {}
// func (s *SSLService) GetSSLCipherServiceBinding()    {}
// func (s *SSLService) CountSSLCipherServiceBinding()  {}

// sslcipher_sslciphersuite_binding
func (s *SSLService) AddSSLCipherSSLCipherSuiteBinding()    {}
func (s *SSLService) DeleteSSLCipherSSLCipherSuiteBinding() {}
func (s *SSLService) GetAllSSLCipherSSLCipherSuiteBinding() {}
func (s *SSLService) GetSSLCipherSSLCipherSuiteBinding()    {}
func (s *SSLService) CountSSLCipherSSLCipherSuiteBinding()  {}

// sslcipher_sslprofile_binding
func (s *SSLService) GetAllSSLCipherSSLProfileBinding() {}
func (s *SSLService) GetSSLCipherSSLProfileBinding()    {}
func (s *SSLService) CountSSLCipherSSLProfileBinding()  {}

// sslcipher_sslvserver_binding
// This object has no operations documented. These are guesses and need to be tested before they will be uncommented.
// func (s *SSLService) GetAllSSLCipherSSLVServerBinding() {}
// func (s *SSLService) GetSSLCipherSSLVServerBinding()    {}
// func (s *SSLService) CountSSLCipherSSLVServerBinding()  {}

// sslcrl
func (s *SSLService) AddSSLCRL()    {}
func (s *SSLService) DeleteSSLCRL() {}
func (s *SSLService) UpdateSSLCRL() {}
func (s *SSLService) UnsetSSLCRL()  {}
func (s *SSLService) CreateSSLCRL() {}
func (s *SSLService) GetAllSSLCRL() {}
func (s *SSLService) GetSSLCRL()    {}
func (s *SSLService) CountSSLCRL()  {}

// sslcrlfile
func (s *SSLService) ImportSSLCRLFile() {}
func (s *SSLService) DeleteSSLCRLFile() {}
func (s *SSLService) GetAllSSLCRLFile() {}
func (s *SSLService) CountSSLCRLFile()  {}

// sslcrl_binding
func (s *SSLService) GetAllSSLCRLBinding() {}
func (s *SSLService) GetSSLCRLBinding()    {}

// sslcrl_serialnumber_binding
func (s *SSLService) GetAllSSLCRLSerialNumberBinding() {}
func (s *SSLService) GetSSLCRLSerialNumberBinding()    {}
func (s *SSLService) CountSSLCRLSerialNumberBinding()  {}

// ssldhfile
func (s *SSLService) ImportSSLDHFile() {}
func (s *SSLService) DeleteSSLDHFile() {}
func (s *SSLService) GetAllSSLDHFile() {}
func (s *SSLService) CountSSLDHFile()  {}

// ssldhparam
func (s *SSLService) ParamSSLDHParam() {}

// ssldtlsprofile
func (s *SSLService) AddSSLDTLSProfile()    {}
func (s *SSLService) DeleteSSLDTLSProfile() {}
func (s *SSLService) UpdateSSLDTLSProfile() {}
func (s *SSLService) UnsetSSLDTLSProfile()  {}
func (s *SSLService) GetAllSSLDTLSProfile() {}
func (s *SSLService) GetSSLDTLSProfile()    {}
func (s *SSLService) CountSSLDTLSProfile()  {}

// sslecdsakey
func (s *SSLService) CreateSSLECDSAKey() {}

// sslfips
func (s *SSLService) UpdateSSLFIPs() {}
func (s *SSLService) UnsetSSLFIPs()  {}
func (s *SSLService) ChangeSSLFIPs() {}
func (s *SSLService) ResetSSLFIPs()  {}
func (s *SSLService) GetAllSSLFIPs() {}

// sslfipskey
func (s *SSLService) CreateSSLFIPsKey() {}
func (s *SSLService) DeleteSSLFIPsKey() {}
func (s *SSLService) ImportSSLFIPsKey() {}
func (s *SSLService) ExportSSLFIPsKey() {}
func (s *SSLService) GetAllSSLFIPsKey() {}
func (s *SSLService) GetSSLFIPsKey()    {}
func (s *SSLService) CountSSLFIPsKey()  {}

// sslfipssimsource
func (s *SSLService) EnableSSLFIPsSIMSource() {}
func (s *SSLService) InitSSLFIPsSIMSource()   {}

// sslfipssimtarget
func (s *SSLService) EnableSSLFIPsSIMTarget() {}
func (s *SSLService) InitSSLFIPsSIMTarget()   {}

// sslglobal_binding
func (s *SSLService) GetSSLGLobalBinding() {}

// sslglobal_sslpolicy_binding
func (s *SSLService) AddSSLGlobalSSLPolicyBinding()    {}
func (s *SSLService) DeleteSSLGlobalSSLPolicyBinding() {}
func (s *SSLService) GetSSLGlobalSSLPolicyBinding()    {}
func (s *SSLService) CountSSLGlobalSSLPolicyBinding()  {}

// sslhsmkey
func (s *SSLService) AddSSLHSMKey()    {}
func (s *SSLService) DeleteSSLHSMKey() {}
func (s *SSLService) GetAllSSLHSMKey() {}
func (s *SSLService) GetSSLHSMKey()    {}
func (s *SSLService) CountSSLHSMKey()  {}

// sslkeyfile
func (s *SSLService) ImportSSLKeyFile() {}
func (s *SSLService) DeleteSSLKeyFile() {}
func (s *SSLService) GetAllSSLKeyFile() {}
func (s *SSLService) CountSSLKeyFile()  {}

// ssllogprofile
func (s *SSLService) AddSSLLogProfile()    {}
func (s *SSLService) DeleteSSLLogProfile() {}
func (s *SSLService) UpdateSSLLogProfile() {}
func (s *SSLService) UnsetSSLLogProfile()  {}
func (s *SSLService) GetAllSSLLogProfile() {}
func (s *SSLService) GetSSLLogProfile()    {}
func (s *SSLService) CountSSLLogProfile()  {}

// sslocspresponder
func (s *SSLService) AddSSLOCSPResponder()    {}
func (s *SSLService) DeleteSSLOCSPResponder() {}
func (s *SSLService) UpdateSSLOCSPResponder() {}
func (s *SSLService) UnsetSSLOCSPResponder()  {}
func (s *SSLService) GetAllSSLOCSPResponder() {}
func (s *SSLService) GetSSLOCSPResponder()    {}
func (s *SSLService) CountSSLOCSPResponder()  {}

// sslparameter
func (s *SSLService) UpdateSSLParameter() {}
func (s *SSLService) UnsetSSLParameter()  {}
func (s *SSLService) GetAllSSLParameter() {}

// sslpkcs12
func (s *SSLService) ConvertSSLPKCS12() {}

// sslpkcs8
func (s *SSLService) ConvertSSLPKCS8() {}

// sslpolicy
func (s *SSLService) AddSSLPolicy()    {}
func (s *SSLService) DeleteSSLPolicy() {}
func (s *SSLService) UpdateSSLPolicy() {}
func (s *SSLService) UnsetSSLPolicy()  {}
func (s *SSLService) GetAllSSLPolicy() {}
func (s *SSLService) GetSSLPolicy()    {}
func (s *SSLService) CountSSLPolicy()  {}

// sslpolicylabel
func (s *SSLService) AddSSLPolicyLabel()    {}
func (s *SSLService) DeleteSSLPolicyLabel() {}
func (s *SSLService) GetAllSSLPolicyLabel() {}
func (s *SSLService) GetSSLPolicyLabel()    {}
func (s *SSLService) CountSSLPolicyLabel()  {}

// sslpolicylabel_binding
func (s *SSLService) GetAllSSLPolicyLabelBinding() {}
func (s *SSLService) GetSSLPolicyLabelBinding()    {}

// sslpolicylabel_sslpolicy_binding
func (s *SSLService) AddSSLpolicyLabelSSLPolicyBinding()    {}
func (s *SSLService) DeleteSSLpolicyLabelSSLPolicyBinding() {}
func (s *SSLService) GetAllSSLpolicyLabelSSLPolicyBinding() {}
func (s *SSLService) GetSSLpolicyLabelSSLPolicyBinding()    {}
func (s *SSLService) CountSSLpolicyLabelSSLPolicyBinding()  {}

// sslpolicy_binding
func (s *SSLService) GetAllSSLPolicyBinding() {}
func (s *SSLService) GetSSLPolicyBinding()    {}

// sslpolicy_csvserver_binding
func (s *SSLService) GetAllSSLPolicyCSVServerBinding() {}
func (s *SSLService) GetSSLPolicyCSVServerBinding()    {}
func (s *SSLService) CountSSLPolicyCSVServerBinding()  {}

// sslpolicy_lbvserver_binding
func (s *SSLService) GetAllSSLPolicyLBVServerBinding() {}
func (s *SSLService) GetSSLPolicyLBVServerBinding()    {}
func (s *SSLService) CountSSLPolicyLBVServerBinding()  {}

// sslpolicy_sslglobal_binding
func (s *SSLService) GetAllSSLPolicySSLGlobalBinding() {}
func (s *SSLService) GetSSLPolicySSLGlobalBinding()    {}
func (s *SSLService) CountSSLPolicySSLGlobalBinding()  {}

// sslpolicy_sslpolicylabel_binding
func (s *SSLService) GetAllSSLPolicySSLPolicyLabelBinding() {}
func (s *SSLService) GetSSLPolicySSLPolicyLabelBinding()    {}
func (s *SSLService) CountSSLPolicySSLPolicyLabelBinding()  {}

// sslpolicy_sslservice_binding
func (s *SSLService) GetAllSSLPolicySSLServiceBinding() {}
func (s *SSLService) GetSSLPolicySSLServiceBinding()    {}
func (s *SSLService) CountSSLPolicySSLServiceBinding()  {}

// sslpolicy_sslvserver_binding
func (s *SSLService) GetAllSSLPolicySSLVServerBinding() {}
func (s *SSLService) GetSSLPolicySSLVServerBinding()    {}
func (s *SSLService) CountSSLPolicySSLVServerBinding()  {}

// sslprofile
func (s *SSLService) AddSSLProfile()    {}
func (s *SSLService) DeleteSSLProfile() {}
func (s *SSLService) UpdateSSLProfile() {}
func (s *SSLService) UnsetSSLProfile()  {}
func (s *SSLService) GetAllSSLProfile() {}
func (s *SSLService) GetSSLProfile()    {}
func (s *SSLService) CountSSLProfile()  {}

// sslprofile_binding
func (s *SSLService) GetAllSSLProfileBinding() {}
func (s *SSLService) GetSSLProfileBinding()    {}

// sslprofile_eccurve_binding
func (s *SSLService) AddSSLProfileECCCurveBinding()    {}
func (s *SSLService) DeleteSSLProfileECCCurveBinding() {}
func (s *SSLService) GetAllSSLProfileECCCurveBinding() {}
func (s *SSLService) GetSSLProfileECCCurveBinding()    {}
func (s *SSLService) CountSSLProfileECCCurveBinding()  {}

// sslprofile_sslcertkey_binding
func (s *SSLService) AddSSLProfileSSLCertKeyBinding()    {}
func (s *SSLService) DeleteSSLProfileSSLCertKeyBinding() {}
func (s *SSLService) GetAllSSLProfileSSLCertKeyBinding() {}
func (s *SSLService) GetSSLProfileSSLCertKeyBinding()    {}
func (s *SSLService) CountSSLProfileSSLCertKeyBinding()  {}

// sslprofile_sslciphersuite_binding
func (s *SSLService) AddSSLProfileSSLCipherSuiteBinding()    {}
func (s *SSLService) DeleteSSLProfileSSLCipherSuiteBinding() {}
func (s *SSLService) GetAllSSLProfileSSLCipherSuiteBinding() {}
func (s *SSLService) GetSSLProfileSSLCipherSuiteBinding()    {}
func (s *SSLService) CountSSLProfileSSLCipherSuiteBinding()  {}

// sslprofile_sslcipher_binding
func (s *SSLService) AddSSLProfileSSLCipherBinding()    {}
func (s *SSLService) DeleteSSLProfileSSLCipherBinding() {}
func (s *SSLService) GetAllSSLProfileSSLCipherBinding() {}
func (s *SSLService) GetSSLProfileSSLCipherBinding()    {}
func (s *SSLService) CountSSLProfileSSLCipherBinding()  {}

// sslprofile_sslvserver_binding
func (s *SSLService) GetAllSSLProfileSSLVServerBinding() {}
func (s *SSLService) GetSSLProfileSSLVServerBinding()    {}
func (s *SSLService) CountSSLProfileSSLVServerBinding()  {}

// sslrsakey
func (s *SSLService) CreateSSLRSAKey() {}

// sslservice
func (s *SSLService) UpdateSSLService() {}
func (s *SSLService) UnsetSSLService()  {}
func (s *SSLService) GetAllSSLService() {}
func (s *SSLService) GetSSLService()    {}
func (s *SSLService) CountSSLService()  {}

// sslservicegroup
func (s *SSLService) UpdateSSLServiceGroup() {}
func (s *SSLService) UnsetSSLServiceGroup()  {}
func (s *SSLService) GetAllSSLServiceGroup() {}
func (s *SSLService) GetSSLServiceGroup()    {}
func (s *SSLService) CountSSLServiceGroup()  {}

// sslservicegroup_binding
func (s *SSLService) GetAllSSLServiceGroupBinding() {}
func (s *SSLService) GetSSLServiceGroupBinding()    {}

// sslservicegroup_ecccurve_binding
func (s *SSLService) AddSSLServiceGroupECCCurveBinding()    {}
func (s *SSLService) DeleteSSLServiceGroupECCCurveBinding() {}
func (s *SSLService) GetAllSSLServiceGroupECCCurveBinding() {}
func (s *SSLService) GetSSLServiceGroupECCCurveBinding()    {}
func (s *SSLService) CountSSLServiceGroupECCCurveBinding()  {}

// sslservicegroup_sslcertkey_binding
func (s *SSLService) AddSSLServiceGroupSSLCertKeyBinding()    {}
func (s *SSLService) DeleteSSLServiceGroupSSLCertKeyBinding() {}
func (s *SSLService) GetAllSSLServiceGroupSSLCertKeyBinding() {}
func (s *SSLService) GetSSLServiceGroupSSLCertKeyBinding()    {}
func (s *SSLService) CountSSLServiceGroupSSLCertKeyBinding()  {}

// sslservicegroup_sslciphersuite_binding
func (s *SSLService) AddSSLServiceGroupSSLCipherSuiteBinding()    {}
func (s *SSLService) DeleteSSLServiceGroupSSLCipherSuiteBinding() {}
func (s *SSLService) GetAllSSLServiceGroupSSLCipherSuiteBinding() {}
func (s *SSLService) GetSSLServiceGroupSSLCipherSuiteBinding()    {}
func (s *SSLService) CountSSLServiceGroupSSLCipherSuiteBinding()  {}

// sslservicegroup_sslcipher_binding
func (s *SSLService) AddSSLServiceGroupSSLCipherBinding()    {}
func (s *SSLService) DeleteSSLServiceGroupSSLCipherBinding() {}
func (s *SSLService) GetAllSSLServiceGroupSSLCipherBinding() {}
func (s *SSLService) GetSSLServiceGroupSSLCipherBinding()    {}
func (s *SSLService) CountSSLServiceGroupSSLCipherBinding()  {}

// sslservice_binding
func (s *SSLService) GetAllSSLServiceBinding() {}
func (s *SSLService) GetSSLServiceBinding()    {}

// sslservice_ecccurve_binding
func (s *SSLService) SSLServiceECCCurveBinding()       {}
func (s *SSLService) DeleteSSLServiceECCCurveBinding() {}
func (s *SSLService) GetAllSSLServiceECCCurveBinding() {}
func (s *SSLService) GetSSLServiceECCCurveBinding()    {}
func (s *SSLService) CountSSLServiceECCCurveBinding()  {}

// sslservice_sslcertkey_binding
func (s *SSLService) AddSSLServiceSSLCertKeyBinding()    {}
func (s *SSLService) DeleteSSLServiceSSLCertKeyBinding() {}
func (s *SSLService) GetAllSSLServiceSSLCertKeyBinding() {}
func (s *SSLService) GetSSLServiceSSLCertKeyBinding()    {}
func (s *SSLService) CountSSLServiceSSLCertKeyBinding()  {}

// sslservice_sslciphersuite_binding
func (s *SSLService) AddSSLServiceSSLCipherSuiteBinding()    {}
func (s *SSLService) DeleteSSLServiceSSLCipherSuiteBinding() {}
func (s *SSLService) GetAllSSLServiceSSLCipherSuiteBinding() {}
func (s *SSLService) GetSSLServiceSSLCipherSuiteBinding()    {}
func (s *SSLService) CountSSLServiceSSLCipherSuiteBinding()  {}

// sslservice_sslcipher_binding
func (s *SSLService) AddSSLServiceSSLCipherBinding()    {}
func (s *SSLService) DeleteSSLServiceSSLCipherBinding() {}
func (s *SSLService) GetAllSSLServiceSSLCipherBinding() {}
func (s *SSLService) GetSSLServiceSSLCipherBinding()    {}
func (s *SSLService) CountSSLServiceSSLCipherBinding()  {}

// sslservice_sslpolicy_binding
func (s *SSLService) AddSSLServiceSSLPolicyBinding()    {}
func (s *SSLService) DeleteSSLServiceSSLPolicyBinding() {}
func (s *SSLService) GetAllSSLServiceSSLPolicyBinding() {}
func (s *SSLService) GetSSLServiceSSLPolicyBinding()    {}
func (s *SSLService) CountSSLServiceSSLPolicyBinding()  {}

// sslvserver
func (s *SSLService) UpdateSSLVServer() {}
func (s *SSLService) UnsetSSLVServer()  {}
func (s *SSLService) GetAllSSLVServer() {}
func (s *SSLService) GetSSLVServer()    {}
func (s *SSLService) CountSSLVServer()  {}

// sslvserver_binding
func (s *SSLService) GetAllSSLVServerBinding() {}
func (s *SSLService) GetSSLVServerBinding()    {}

// sslvserver_ecccurve_binding
func (s *SSLService) AddSSLVServerECCCurveBinding()    {}
func (s *SSLService) DeleteSSLVServerECCCurveBinding() {}
func (s *SSLService) GetAllSSLVServerECCCurveBinding() {}
func (s *SSLService) GetSSLVServerECCCurveBinding()    {}
func (s *SSLService) CountSSLVServerECCCurveBinding()  {}

// sslvserver_sslcertkey_binding
func (s *SSLService) AddSSLVServerSSLCertKeyBinding()    {}
func (s *SSLService) DeleteSSLVServerSSLCertKeyBinding() {}
func (s *SSLService) GetAllSSLVServerSSLCertKeyBinding() {}
func (s *SSLService) GetSSLVServerSSLCertKeyBinding()    {}
func (s *SSLService) CountSSLVServerSSLCertKeyBinding()  {}

// sslvserver_sslciphersuite_binding
func (s *SSLService) AddSSLVServerSSLCipherSuiteBinding()    {}
func (s *SSLService) DeleteSSLVServerSSLCipherSuiteBinding() {}
func (s *SSLService) GetAllSSLVServerSSLCipherSuiteBinding() {}
func (s *SSLService) GetSSLVServerSSLCipherSuiteBinding()    {}
func (s *SSLService) CountSSLVServerSSLCipherSuiteBinding()  {}

// sslvserver_sslcipher_binding
func (s *SSLService) AddSSLVServerSSLCipherBInding()    {}
func (s *SSLService) DeleteSSLVServerSSLCipherBInding() {}
func (s *SSLService) GetAllSSLVServerSSLCipherBInding() {}
func (s *SSLService) GetSSLVServerSSLCipherBInding()    {}
func (s *SSLService) CountSSLVServerSSLCipherBInding()  {}

// sslvserver_sslpolicy_binding
func (s *SSLService) AddSSLVServerSSLPolicyBinding()    {}
func (s *SSLService) DeleteSSLVServerSSLPolicyBinding() {}
func (s *SSLService) GetAllSSLVServerSSLPolicyBinding() {}
func (s *SSLService) GetSSLVServerSSLPolicyBinding()    {}
func (s *SSLService) CountSSLVServerSSLPolicyBinding()  {}

// sslwrapkey
func (s *SSLService) CreateSSLWrapKey() {}
func (s *SSLService) DeleteSSLWrapKey() {}
func (s *SSLService) GetAllSSLWrapKey() {}
func (s *SSLService) CountSSLWrapKey()  {}
