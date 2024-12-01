package nitrogo

// Application Firewall configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfw
type AppFWService struct {
	client *Client
}

// appfwarchive
// Configuration for archive resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwarchive
func (s *AppFWService) DeleteAppFWArchive() {}
func (s *AppFWService) ExportAppFWArchive() {}
func (s *AppFWService) ImportAppFWArchive() {}
func (s *AppFWService) GetAllAppFWArchive() {}

// appfwconfidfield
// Configuration for configured confidential form fields resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwconfidfield
func (s *AppFWService) AddAppFWConfigField()    {}
func (s *AppFWService) DeleteAppFWConfigField() {}
func (s *AppFWService) UpdateAppFWConfigField() {}
func (s *AppFWService) UnsetAppFWConfigField()  {}
func (s *AppFWService) GetAllAppFWConfigField() {}
func (s *AppFWService) CountAppFWConfigField()  {}

// appfwcustomsettings
// Configuration for application firewall custom settings XML configuration resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwcustomsettings
func (s *AppFWService) ExportAppFWCustomSettings() {}

// appfwfieldtype
// Configuration for application firewall form field type resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwfieldtype
func (s *AppFWService) AddAppFWFieldType()    {}
func (s *AppFWService) DeleteAppFWFieldType() {}
func (s *AppFWService) UpdateAppFWFieldType() {}
func (s *AppFWService) GetAllAppFWFieldType() {}
func (s *AppFWService) GetAppFWFieldType()    {}
func (s *AppFWService) CountAppFWFieldType()  {}

// appfwglobal_appfwpolicy_binding
// Binding object showing the appfwpolicy that can be bound to appfwglobal.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwglobal_appfwpolicy_binding
func (s *AppFWService) AddAppFWGlobalAppFWPolicyBinding()    {}
func (s *AppFWService) DeleteAppFWGlobalAppFWPolicyBinding() {}
func (s *AppFWService) GetAppFWGlobalAppFWPolicyBinding()    {}
func (s *AppFWService) CountAppFWGlobalAppFWPolicyBinding()  {}

// appfwglobal_auditnslogpolicy_binding
// Binding object showing the auditnslogpolicy that can be bound to appfwglobal.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwglobal_auditnslogpolicy_binding
func (s *AppFWService) AddAppFWGlobalAuditNSLogPolicyBinding()    {}
func (s *AppFWService) DeleteAppFWGlobalAuditNSLogPolicyBinding() {}
func (s *AppFWService) GetAppFWGlobalAuditNSLogPolicyBinding()    {}
func (s *AppFWService) CountAppFWGlobalAuditNSLogPolicyBinding()  {}

// appfwglobal_audtisyslogpolicy_binding
// Binding object showing the auditsyslogpolicy that can be bound to appfwglobal.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwglobal_auditsyslogpolicy_binding
func (s *AppFWService) AddAppFWGlobalAuditSyslogPolicyBinding()    {}
func (s *AppFWService) DeleteAppFWGlobalAuditSyslogPolicyBinding() {}
func (s *AppFWService) GetAppFWGlobalAuditSyslogPolicyBinding()    {}
func (s *AppFWService) CountAppFWGlobalAuditSyslogPolicyBinding()  {}

// appfwglobal_binding
// Binding object which returns the resources bound to appfwglobal.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwglobal_binding
func (s *AppFWService) GetAppFWGlobalBinding() {}

// appfwhtmlerrorpage
// Configuration for HTML error page resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwhtmlerrorpage
func (s *AppFWService) DeleteAppFWHTMLErrorPage() {}
func (s *AppFWService) GetAllAppFWHTMLErrorPage() {}
func (s *AppFWService) GetAppFWHTMLErrorPage()    {}
func (s *AppFWService) ImportAppFWHTMLErrorPage() {}
func (s *AppFWService) ChangeAppFWHTMLErrorPage() {}

// appfwjsoncontenttype
// Configuration for JSON content type resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwjsoncontenttype
func (s *AppFWService) AddAppFWJSONContentType()    {}
func (s *AppFWService) DeleteAppFWJSONContentType() {}
func (s *AppFWService) GetAllAppFWJSONContentType() {}
func (s *AppFWService) GetAppFWJSONContentType()    {}
func (s *AppFWService) CountAppFWJSONContentType()  {}

// appfwjsonerrorpage
// Configuration for JSON error page resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwjsonerrorpage
func (s *AppFWService) DeleteAppFWJSONErrorPage() {}
func (s *AppFWService) GetAllAppFWJSONErrorPage() {}
func (s *AppFWService) GetAppFWJSONErrorPage()    {}
func (s *AppFWService) ImportAppFWJSONErrorPage() {}
func (s *AppFWService) ChangeAppFWJSONErrorPage() {}

// appfwlearningdata
// Configuration for learning data resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwlearningdata
func (s *AppFWService) DeleteAppFWLearningData() {}
func (s *AppFWService) GetAllAppFWLearningData() {}
func (s *AppFWService) CountAppFWLearningData()  {}
func (s *AppFWService) ResetAppFWLearningData()  {}
func (s *AppFWService) ExportAppFWLearningData() {}

// appfwlearningsettings
// Configuration for learning settings resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwlearningsettings
func (s *AppFWService) UpdateAppFWLearningSettings() {}
func (s *AppFWService) UnsetAppFWLearningSettings()  {}
func (s *AppFWService) GetAllAppFWLearningSettings() {}
func (s *AppFWService) GetAppFWLearningSettings()    {}
func (s *AppFWService) CountAppFWLearningSettings()  {}

// appfwmultipartformcontenttype
// Configuration for Multipart form content type resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwmultipartformcontenttype
func (s *AppFWService) AddAppFWMultipartFormContentType()    {}
func (s *AppFWService) DeleteAppFWMultipartFormContentType() {}
func (s *AppFWService) GetAllAppFWMultipartFormContentType() {}
func (s *AppFWService) GetAppFWMultipartFormContentType()    {}
func (s *AppFWService) CountAppFWMultipartFormContentType()  {}

// appfwpolicy
// Configuration for application firewall policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwpolicy
func (s *AppFWService) AddAppFWPolicy()    {}
func (s *AppFWService) DeleteAppFWPolicy() {}
func (s *AppFWService) UpdateAppFWPolicy() {}
func (s *AppFWService) UnsetAppFWPolicy()  {}
func (s *AppFWService) GetAllAppFWPolicy() {}
func (s *AppFWService) GetAppFWPolicy()    {}
func (s *AppFWService) CountAppFWPolicy()  {}
func (s *AppFWService) RenameAppFWPolicy() {}

// appfwpolicylabel
// Configuration for application firewall policy label resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwpolicylabel
func (s *AppFWService) AddAppFWPolicyLabel()    {}
func (s *AppFWService) DeleteAppFWPolicyLabel() {}
func (s *AppFWService) GetAllAppFWPolicyLabel() {}
func (s *AppFWService) GetAppFWPolicyLabel()    {}
func (s *AppFWService) CountAppFWPolicyLabel()  {}
func (s *AppFWService) RenameAppFWPolicyLabel() {}

// appfwpolicylabel_appfwpolicy_binding
// Binding object showing the appfwpolicy that can be bound to appfwpolicylabel.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwpolicylabel_appfwpolicy_binding
func (s *AppFWService) AddAppFWPolicyLabelAppFWPolicyBinding()    {}
func (s *AppFWService) DeleteAppFWPolicyLabelAppFWPolicyBinding() {}
func (s *AppFWService) GetAllAppFWPolicyLabelAppFWPolicyBinding() {}
func (s *AppFWService) GetAppFWPolicyLabelAppFWPolicyBinding()    {}
func (s *AppFWService) CountAppFWPolicyLabelAppFWPolicyBinding()  {}

// appfwpolicylabel_binding
// Binding object which returns the resources bound to appfwpolicylabel.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwpolicylabel_binding
func (s *AppFWService) GetAllAppFWPolicyLabelBinding() {}
func (s *AppFWService) GetAppFWPolicyLabelBinding()    {}

// appfwpolicylabel_policybinding_binding
// Binding object showing the policybinding that can be bound to appfwpolicylabel.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwpolicylabel_policybinding_binding
func (s *AppFWService) GetAllAppFWPolicyLabelPolicyBindingBinding() {}
func (s *AppFWService) GetAppFWPolicyLabelPolicyBindingBinding()    {}
func (s *AppFWService) CountAppFWPolicyLabelPolicyBindingBinding()  {}

// appfwpolicy_appfwglobal_binding
// Binding object showing the appfwglobal that can be bound to appfwpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwpolicy_appfwglobal_binding
func (s *AppFWService) GetAllAppFWPolicyAppFWGlobalBinding() {}
func (s *AppFWService) GetAppFWPolicyAppFWGlobalBinding()    {}
func (s *AppFWService) CountAppFWPolicyAppFWGlobalBinding()  {}

// appfwpolicy_appfwpolicylabel_binding
// Binding object showing the appfwpolicylabel that can be bound to appfwpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwpolicy_appfwpolicylabel_binding
func (s *AppFWService) GetAllAppFWPolicyAppFWPolicyLabelBinding() {}
func (s *AppFWService) GetAppFWPolicyAppFWPolicyLabelBinding()    {}
func (s *AppFWService) CountAppFWPolicyAppFWPolicyLabelBinding()  {}

// appfwpolicy_binding
// Binding object which returns the resources bound to appfwpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwpolicy_binding
func (s *AppFWService) GetAllAppFWPolcyBinding() {}
func (s *AppFWService) GetAppFWPolcyBinding()    {}

// appfwpolicy_csvserver_binding
// Binding object showing the csvserver that can be bound to appfwpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwpolicy_csvserver_binding
func (s *AppFWService) GetAllAppFWPolicyCSVServerBinding() {}
func (s *AppFWService) GetAppFWPolicyCSVServerBinding()    {}
func (s *AppFWService) CountAppFWPolicyCSVServerBinding()  {}

// appfwpolicy_lbvserver_binding
// Binding object showing the lbvserver that can be bound to appfwpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwpolicy_lbvserver_binding
func (s *AppFWService) GetAllAppFWPolcyLBVServerBinding() {}
func (s *AppFWService) GetAppFWPolcyLBVServerBinding()    {}
func (s *AppFWService) CountAppFWPolcyLBVServerBinding()  {}

// appfwprofile
// Configuration for application firewall profile resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile
func (s *AppFWService) AddAppFWProfile()     {}
func (s *AppFWService) DeleteAppFWProfile()  {}
func (s *AppFWService) UpdateAppFWProfile()  {}
func (s *AppFWService) UnsetAppFWProfile()   {}
func (s *AppFWService) GetAllAppFWProfile()  {}
func (s *AppFWService) GetAppFWProfile()     {}
func (s *AppFWService) CountAppFWProfile()   {}
func (s *AppFWService) RestoreAppFWProfile() {}

// appfwprofile_binding
// Binding object which returns the resources bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_binding
func (s *AppFWService) GetAllAppFWProfileBinding() {}
func (s *AppFWService) GetAppFWProfileBinding()    {}

// appfwprofile_cmdinjection_binding
// Binding object showing the cmdinjection that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_cmdinjection_binding
func (s *AppFWService) AddAppFWProfileCMDInjectionBinding()    {}
func (s *AppFWService) DeleteAppFWProfileCMDInjectionBinding() {}
func (s *AppFWService) GetAllAppFWProfileCMDInjectionBinding() {}
func (s *AppFWService) GetAppFWProfileCMDInjectionBinding()    {}
func (s *AppFWService) CountAppFWProfileCMDInjectionBinding()  {}

// appfwprofile_contenttype_binding
// Binding object showing the contenttype that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_contenttype_binding
func (s *AppFWService) AddAppFWProfileContentypeBinding()    {}
func (s *AppFWService) DeleteAppFWProfileContentypeBinding() {}
func (s *AppFWService) GetAllAppFWProfileContentypeBinding() {}
func (s *AppFWService) GetAppFWProfileContentypeBinding()    {}
func (s *AppFWService) CountAppFWProfileContentypeBinding()  {}

// appfwprofile_cookieconsistency_binding
// Binding object showing the cookieconsistency that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_cookieconsistency_binding
func (s *AppFWService) AddAppFWProfileCookieConsistencyBinding()    {}
func (s *AppFWService) DeleteAppFWProfileCookieConsistencyBinding() {}
func (s *AppFWService) GetAllAppFWProfileCookieConsistencyBinding() {}
func (s *AppFWService) GetAppFWProfileCookieConsistencyBinding()    {}
func (s *AppFWService) CountAppFWProfileCookieConsistencyBinding()  {}

// appfwprofile_creditcardnumber_binding
// Binding object showing the creditcardnumber that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_creditcardnumber_binding
func (s *AppFWService) AddAppFWProfileCreditCardNumberBinding()    {}
func (s *AppFWService) DeleteAppFWProfileCreditCardNumberBinding() {}
func (s *AppFWService) GetAllAppFWProfileCreditCardNumberBinding() {}
func (s *AppFWService) GetAppFWProfileCreditCardNumberBinding()    {}
func (s *AppFWService) CountAppFWProfileCreditCardNumberBinding()  {}

// appfwprofile_crosssitescripting_binding
// Binding object showing the crosssitescripting that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_crosssitescripting_binding
func (s *AppFWService) AddAppFWProfileCrossSiteScriptingBinding()    {}
func (s *AppFWService) DeleteAppFWProfileCrossSiteScriptingBinding() {}
func (s *AppFWService) GetAllAppFWProfileCrossSiteScriptingBinding() {}
func (s *AppFWService) GetAppFWProfileCrossSiteScriptingBinding()    {}
func (s *AppFWService) CountAppFWProfileCrossSiteScriptingBinding()  {}

// appfwprofile_csrftag_binding
// Binding object showing the csrftag that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_csrftag_binding
func (s *AppFWService) AddAppFWProfileCSRFTagBinding()    {}
func (s *AppFWService) DeleteAppFWProfileCSRFTagBinding() {}
func (s *AppFWService) GetAllAppFWProfileCSRFTagBinding() {}
func (s *AppFWService) GetAppFWProfileCSRFTagBinding()    {}
func (s *AppFWService) CountAppFWProfileCSRFTagBinding()  {}

// appfwprofile_denyurl_binding
// Binding object showing the denyurl that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_denyurl_binding
func (s *AppFWService) AddAppFWProfileDenyURLBinding()    {}
func (s *AppFWService) DeleteAppFWProfileDenyURLBinding() {}
func (s *AppFWService) GetAllAppFWProfileDenyURLBinding() {}
func (s *AppFWService) GetAppFWProfileDenyURLBinding()    {}
func (s *AppFWService) CountAppFWProfileDenyURLBinding()  {}

// appfwprofile_excluderescontenttype_binding
// Binding object showing the excluderescontenttype that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_excluderescontenttype_binding
func (s *AppFWService) AddAppFWProfileExcludeRESContentTypeBinding()    {}
func (s *AppFWService) DeleteAppFWProfileExcludeRESContentTypeBinding() {}
func (s *AppFWService) GetAllAppFWProfileExcludeRESContentTypeBinding() {}
func (s *AppFWService) GetAppFWProfileExcludeRESContentTypeBinding()    {}
func (s *AppFWService) CountAppFWProfileExcludeRESContentTypeBinding()  {}

// appfwprofile_fieldconsistency_binding
// Binding object showing the fieldconsistency that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_fieldconsistency_binding
func (s *AppFWService) AddAppFWProfileFieldConsistencyBinding()    {}
func (s *AppFWService) DeleteAppFWProfileFieldConsistencyBinding() {}
func (s *AppFWService) GetAllAppFWProfileFieldConsistencyBinding() {}
func (s *AppFWService) GetAppFWProfileFieldConsistencyBinding()    {}
func (s *AppFWService) CountAppFWProfileFieldConsistencyBinding()  {}

// appfwprofile_fieldformat_binding
// Binding object showing the fieldformat that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_fieldformat_binding
func (s *AppFWService) AddAppFWProfileFieldFormatBinding()    {}
func (s *AppFWService) DeleteAppFWProfileFieldFormatBinding() {}
func (s *AppFWService) GetAllAppFWProfileFieldFormatBinding() {}
func (s *AppFWService) GetAppFWProfileFieldFormatBinding()    {}
func (s *AppFWService) CountAppFWProfileFieldFormatBinding()  {}

// appfwprofile_fileuploadtype_binding
// Binding object showing the fileuploadtype that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_fileuploadtype_binding
func (s *AppFWService) AddAppFWProfileFielUploadTypeBinding()    {}
func (s *AppFWService) DeleteAppFWProfileFielUploadTypeBinding() {}
func (s *AppFWService) GetAllAppFWProfileFielUploadTypeBinding() {}
func (s *AppFWService) GetAppFWProfileFielUploadTypeBinding()    {}
func (s *AppFWService) CountAppFWProfileFielUploadTypeBinding()  {}

// appfwprofile_jsondosurl_binding
// Binding object showing the jsondosurl that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_jsondosurl_binding
func (s *AppFWService) AddAppFWProfileJSONDOSURLBinding()    {}
func (s *AppFWService) DeleteAppFWProfileJSONDOSURLBinding() {}
func (s *AppFWService) GetAllAppFWProfileJSONDOSURLBinding() {}
func (s *AppFWService) GetAppFWProfileJSONDOSURLBinding()    {}
func (s *AppFWService) CountAppFWProfileJSONDOSURLBinding()  {}

// appfwprofile_jsonsqlurl_binding
// Binding object showing the jsonsqlurl that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_jsonsqlurl_binding
func (s *AppFWService) AddAppFWProfileJSONSQLURLBinding()    {}
func (s *AppFWService) DeleteAppFWProfileJSONSQLURLBinding() {}
func (s *AppFWService) GetAllAppFWProfileJSONSQLURLBinding() {}
func (s *AppFWService) GetAppFWProfileJSONSQLURLBinding()    {}
func (s *AppFWService) CountAppFWProfileJSONSQLURLBinding()  {}

// appfwprofile_jsonxssurl_binding
// Binding object showing the jsonxssurl that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_jsonxssurl_binding
func (s *AppFWService) AddAppFWProfileJSONXSSURLBinding()    {}
func (s *AppFWService) DeleteAppFWProfileJSONXSSURLBinding() {}
func (s *AppFWService) GetAllAppFWProfileJSONXSSURLBinding() {}
func (s *AppFWService) GetAppFWProfileJSONXSSURLBinding()    {}
func (s *AppFWService) CountAppFWProfileJSONXSSURLBinding()  {}

// appfwprofile_logexpression_binding
// Binding object showing the logexpression that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_logexpression_binding
func (s *AppFWService) AddAppFWProfileLogExpressionBinding()    {}
func (s *AppFWService) DeleteAppFWProfileLogExpressionBinding() {}
func (s *AppFWService) GetAllAppFWProfileLogExpressionBinding() {}
func (s *AppFWService) GetAppFWProfileLogExpressionBinding()    {}
func (s *AppFWService) CountAppFWProfileLogExpressionBinding()  {}

// appfwprofile_safeobject_binding
// Binding object showing the safeobject that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_safeobject_binding
func (s *AppFWService) AddAppFWProfileSafeObjectBinding()    {}
func (s *AppFWService) DeleteAppFWProfileSafeObjectBinding() {}
func (s *AppFWService) GetAllAppFWProfileSafeObjectBinding() {}
func (s *AppFWService) GetAppFWProfileSafeObjectBinding()    {}
func (s *AppFWService) CountAppFWProfileSafeObjectBinding()  {}

// appfwprofile_sqlinjection_binding
// Binding object showing the sqlinjection that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_sqlinjection_binding
func (s *AppFWService) AddAppFWProfileSQLInjectionBinding()    {}
func (s *AppFWService) DeleteAppFWProfileSQLInjectionBinding() {}
func (s *AppFWService) GetAllAppFWProfileSQLInjectionBinding() {}
func (s *AppFWService) GetAppFWProfileSQLInjectionBinding()    {}
func (s *AppFWService) CountAppFWProfileSQLInjectionBinding()  {}

// appfwprofile_starturl_binding
// Binding object showing the starturl that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_starturl_binding
func (s *AppFWService) AddAppFWProfileStartURLBinding()    {}
func (s *AppFWService) DeleteAppFWProfileStartURLBinding() {}
func (s *AppFWService) GetAllAppFWProfileStartURLBinding() {}
func (s *AppFWService) GetAppFWProfileStartURLBinding()    {}
func (s *AppFWService) CountAppFWProfileStartURLBinding()  {}

// appfwprofile_trustedlearningclients_binding
// Binding object showing the trustedlearningclients that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_trustedlearningclients_binding
func (s *AppFWService) AddAppFWProfileTrustedLearningClientsBinding()    {}
func (s *AppFWService) DeleteAppFWProfileTrustedLearningClientsBinding() {}
func (s *AppFWService) GetAllAppFWProfileTrustedLearningClientsBinding() {}
func (s *AppFWService) GetAppFWProfileTrustedLearningClientsBinding()    {}
func (s *AppFWService) CountAppFWProfileTrustedLearningClientsBinding()  {}

// appfwprofile_xmlattachmenturl_binding
// Binding object showing the xmlattachmenturl that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_xmlattachmenturl_binding
func (s *AppFWService) AddAppFWProfileXMLAttachmentURLBinding()    {}
func (s *AppFWService) DeleteAppFWProfileXMLAttachmentURLBinding() {}
func (s *AppFWService) GetAllAppFWProfileXMLAttachmentURLBinding() {}
func (s *AppFWService) GetAppFWProfileXMLAttachmentURLBinding()    {}
func (s *AppFWService) CountAppFWProfileXMLAttachmentURLBinding()  {}

// appfwprofile_xmldosurl_binding
// Binding object showing the xmldosurl that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_xmldosurl_binding
func (s *AppFWService) AddAppFWProfileXMLDOSURLBinding()    {}
func (s *AppFWService) DeleteAppFWProfileXMLDOSURLBinding() {}
func (s *AppFWService) GetAllAppFWProfileXMLDOSURLBinding() {}
func (s *AppFWService) GetAppFWProfileXMLDOSURLBinding()    {}
func (s *AppFWService) CountAppFWProfileXMLDOSURLBinding()  {}

// appfwprofile_xmlsqlinjection_binding
// Binding object showing the xmlsqlinjection that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_xmlsqlinjection_binding
func (s *AppFWService) AddAppFWProfileXMLSQLInjectionBinding()    {}
func (s *AppFWService) DeleteAppFWProfileXMLSQLInjectionBinding() {}
func (s *AppFWService) GetAllAppFWProfileXMLSQLInjectionBinding() {}
func (s *AppFWService) GetAppFWProfileXMLSQLInjectionBinding()    {}
func (s *AppFWService) CountAppFWProfileXMLSQLInjectionBinding()  {}

// appfwprofile_xmlvalidationurl_binding
// Binding object showing the xmlvalidationurl that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_xmlvalidationurl_binding
func (s *AppFWService) AddAppFWProfileXMLValidationURLBinding()    {}
func (s *AppFWService) DeleteAppFWProfileXMLValidationURLBinding() {}
func (s *AppFWService) GetAllAppFWProfileXMLValidationURLBinding() {}
func (s *AppFWService) GetAppFWProfileXMLValidationURLBinding()    {}
func (s *AppFWService) CountAppFWProfileXMLValidationURLBinding()  {}

// appfwprofile_xmlwsiurl_binding
// Binding object showing the xmlwsiurl that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_xmlwsiurl_binding
func (s *AppFWService) AddAppFWProfileXMLWSIURLBinding()    {}
func (s *AppFWService) DeleteAppFWProfileXMLWSIURLBinding() {}
func (s *AppFWService) GetAllAppFWProfileXMLWSIURLBinding() {}
func (s *AppFWService) GetAppFWProfileXMLWSIURLBinding()    {}
func (s *AppFWService) CountAppFWProfileXMLWSIURLBinding()  {}

// appfwprofile_xmlxss_binding
// Binding object showing the xmlxss that can be bound to appfwprofile.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwprofile_xmlxss_binding
func (s *AppFWService) AddAppFWProfileXMLXSSBinding()    {}
func (s *AppFWService) DeleteAppFWProfileXMLXSSBinding() {}
func (s *AppFWService) GetAllAppFWProfileXMLXSSBinding() {}
func (s *AppFWService) GetAppFWProfileXMLXSSBinding()    {}
func (s *AppFWService) CountAppFWProfileXMLXSSBinding()  {}

// appfwsettings
// Configuration for AS settings resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwsettings
func (s *AppFWService) UpdateAppFWSettings() {}
func (s *AppFWService) UnsetAppFWSettings()  {}
func (s *AppFWService) GetAllAppFWSettings() {}

// appfwsignatures
// Configuration for application firewall signatures XML configuration resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwsignatures
func (s *AppFWService) DeleteAppFWSignatures() {}
func (s *AppFWService) GetAllAppFWSignatures() {}
func (s *AppFWService) GetAppFWSignatures()    {}
func (s *AppFWService) ImportAppFWSignatures() {}
func (s *AppFWService) ChangeAppFWSignatures() {}

// appfwtransactionrecords
// Configuration for Application firewall transaction record resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwtransactionrecords
func (s *AppFWService) GetAllAppFWTransactionRecords() {}
func (s *AppFWService) CountAppFWTransactionRecords()  {}

// appfwurlencodedformcontenttype
// Configuration for Urlencoded form content type resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwurlencodedformcontenttype
func (s *AppFWService) AddAppFWURLEncodedFormContentType()    {}
func (s *AppFWService) DeleteAppFWURLEncodedFormContentType() {}
func (s *AppFWService) GetAllAppFWURLEncodedFormContentType() {}
func (s *AppFWService) GetAppFWURLEncodedFormContentType()    {}
func (s *AppFWService) CountAppFWURLEncodedFormContentType()  {}

// appfwwsdl
// Configuration for WSDL file resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwwsdl
func (s *AppFWService) DeleteAppFWWSDL() {}
func (s *AppFWService) GetAllAppFWWSDL() {}
func (s *AppFWService) GetAppFWWSDL()    {}
func (s *AppFWService) ImportAppFWWSDL() {}

// appfwxmlcontenttype
// Configuration for XML Content type resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwxmlcontenttype
func (s *AppFWService) AddAppFWXMLContentType()    {}
func (s *AppFWService) DeleteAppFWXMLContentType() {}
func (s *AppFWService) GetAllAppFWXMLContentType() {}
func (s *AppFWService) GetAppFWXMLContentType()    {}
func (s *AppFWService) CountAppFWXMLContentType()  {}

// appfwxmlerrorpage
// Configuration for xml error page resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwxmlerrorpage
func (s *AppFWService) DeleteAppFWXMLErrorPage() {}
func (s *AppFWService) GetAllAppFWXMLErrorPage() {}
func (s *AppFWService) GetAppFWXMLErrorPage()    {}
func (s *AppFWService) ImportAppFWXMLErrorPage() {}
func (s *AppFWService) ChangeAppFWXMLErrorPage() {}

// appfwxmlscheme
// Configuration for XML schema resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appfw/appfwxmlschema
func (s *AppFWService) DeleteAppFWXMLScheme() {}
func (s *AppFWService) GetAllAppFWXMLScheme() {}
func (s *AppFWService) GetAppFWXMLScheme()    {}
func (s *AppFWService) ImportAppFWXMLScheme() {}
