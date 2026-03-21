package models

// appfw configuration structs
type AppFWXMLErrorPage struct {
	Comment            string `json:"comment,omitempty"`
	Name               string `json:"name,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	Overwrite          bool   `json:"overwrite,omitempty"`
	Response           string `json:"response,omitempty"`
	Src                string `json:"src,omitempty"`
}

type AppFWPolicyLBVServerBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type AppFWGRPCWebTextContentType struct {
	Count                       float64 `json:"__count,omitempty"`
	GRPCWebTextContentTypeValue string  `json:"grpcwebtextcontenttypevalue,omitempty"`
	IsRegex                     string  `json:"isregex,omitempty"`
	NextGenAPIResource          string  `json:"_nextgenapiresource,omitempty"`
}

type AppFWGlobalAppFWPolicyBinding struct {
	FlowType               int    `json:"flowtype,omitempty"`
	GlobalBindType         string `json:"globalbindtype,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	NumPol                 int    `json:"numpol,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	PolicyType             string `json:"policytype,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	State                  string `json:"state,omitempty"`
	TypeField              string `json:"type,omitempty"`
}

type AppFWProfileCookieConsistencyBinding struct {
	AlertOnly         string `json:"alertonly,omitempty"`
	Comment           string `json:"comment,omitempty"`
	CookieConsistency string `json:"cookieconsistency,omitempty"`
	IsAutoDeployed    string `json:"isautodeployed,omitempty"`
	IsRegex           string `json:"isregex,omitempty"`
	Name              string `json:"name,omitempty"`
	ResourceID        string `json:"resourceid,omitempty"`
	RuleType          string `json:"ruletype,omitempty"`
	State             string `json:"state,omitempty"`
}

type AppFWProfileXMLDOSURLBinding struct {
	AlertOnly                       string `json:"alertonly,omitempty"`
	Comment                         string `json:"comment,omitempty"`
	IsAutoDeployed                  string `json:"isautodeployed,omitempty"`
	Name                            string `json:"name,omitempty"`
	ResourceID                      string `json:"resourceid,omitempty"`
	RuleType                        string `json:"ruletype,omitempty"`
	State                           string `json:"state,omitempty"`
	XMLBlockDTD                     string `json:"xmlblockdtd,omitempty"`
	XMLBlockExternalEntities        string `json:"xmlblockexternalentities,omitempty"`
	XMLBlockPI                      string `json:"xmlblockpi,omitempty"`
	XMLDOSURL                       string `json:"xmldosurl,omitempty"`
	XMLMaxAttributeNameLength       int    `json:"xmlmaxattributenamelength,omitempty"`
	XMLMaxAttributeNameLengthCheck  string `json:"xmlmaxattributenamelengthcheck,omitempty"`
	XMLMaxAttributes                int    `json:"xmlmaxattributes,omitempty"`
	XMLMaxAttributesCheck           string `json:"xmlmaxattributescheck,omitempty"`
	XMLMaxAttributeValueLength      int    `json:"xmlmaxattributevaluelength,omitempty"`
	XMLMaxAttributeValueLengthCheck string `json:"xmlmaxattributevaluelengthcheck,omitempty"`
	XMLMaxCharDataLength            int    `json:"xmlmaxchardatalength,omitempty"`
	XMLMaxCharDataLengthCheck       string `json:"xmlmaxchardatalengthcheck,omitempty"`
	XMLMaxElementChildren           int    `json:"xmlmaxelementchildren,omitempty"`
	XMLMaxElementChildrenCheck      string `json:"xmlmaxelementchildrencheck,omitempty"`
	XMLMaxElementDepth              int    `json:"xmlmaxelementdepth,omitempty"`
	XMLMaxElementDepthCheck         string `json:"xmlmaxelementdepthcheck,omitempty"`
	XMLMaxElementNameLength         int    `json:"xmlmaxelementnamelength,omitempty"`
	XMLMaxElementNameLengthCheck    string `json:"xmlmaxelementnamelengthcheck,omitempty"`
	XMLMaxElements                  int    `json:"xmlmaxelements,omitempty"`
	XMLMaxElementsCheck             string `json:"xmlmaxelementscheck,omitempty"`
	XMLMaxEntityExpansionDepth      int    `json:"xmlmaxentityexpansiondepth,omitempty"`
	XMLMaxEntityExpansionDepthCheck string `json:"xmlmaxentityexpansiondepthcheck,omitempty"`
	XMLMaxEntityExpansions          int    `json:"xmlmaxentityexpansions,omitempty"`
	XMLMaxEntityExpansionsCheck     string `json:"xmlmaxentityexpansionscheck,omitempty"`
	XMLMaxFileSize                  int    `json:"xmlmaxfilesize,omitempty"`
	XMLMaxFileSizeCheck             string `json:"xmlmaxfilesizecheck,omitempty"`
	XMLMaxNamespaces                int    `json:"xmlmaxnamespaces,omitempty"`
	XMLMaxNamespacesCheck           string `json:"xmlmaxnamespacescheck,omitempty"`
	XMLMaxNamespaceURILength        int    `json:"xmlmaxnamespaceurilength,omitempty"`
	XMLMaxNamespaceURILengthCheck   string `json:"xmlmaxnamespaceurilengthcheck,omitempty"`
	XMLMaxNodes                     int    `json:"xmlmaxnodes,omitempty"`
	XMLMaxNodesCheck                string `json:"xmlmaxnodescheck,omitempty"`
	XMLMaxSOAPArrayRank             int    `json:"xmlmaxsoaparrayrank,omitempty"`
	XMLMaxSOAPArraySize             int    `json:"xmlmaxsoaparraysize,omitempty"`
	XMLMinFileSize                  int    `json:"xmlminfilesize,omitempty"`
	XMLMinFileSizeCheck             string `json:"xmlminfilesizecheck,omitempty"`
	XMLSOAPArrayCheck               string `json:"xmlsoaparraycheck,omitempty"`
}

type AppFWProfileDenyURLBinding struct {
	AlertOnly      string `json:"alertonly,omitempty"`
	Comment        string `json:"comment,omitempty"`
	DenyURL        string `json:"denyurl,omitempty"`
	IsAutoDeployed string `json:"isautodeployed,omitempty"`
	Name           string `json:"name,omitempty"`
	ResourceID     string `json:"resourceid,omitempty"`
	RuleType       string `json:"ruletype,omitempty"`
	State          string `json:"state,omitempty"`
}

type AppFWSignatures struct {
	Action                  []string `json:"action,omitempty"`
	AutoEnableNewSignatures string   `json:"autoenablenewsignatures,omitempty"`
	Category                string   `json:"category,omitempty"`
	Comment                 string   `json:"comment,omitempty"`
	Enabled                 string   `json:"enabled,omitempty"`
	EncryptedVersion        int      `json:"encryptedversion,omitempty"`
	Merge                   bool     `json:"merge,omitempty"`
	MergeDefault            bool     `json:"mergedefault,omitempty"`
	Name                    string   `json:"name,omitempty"`
	NextGenAPIResource      string   `json:"_nextgenapiresource,omitempty"`
	Overwrite               bool     `json:"overwrite,omitempty"`
	PreserveDefActions      bool     `json:"preservedefactions,omitempty"`
	Response                string   `json:"response,omitempty"`
	RuleID                  []any    `json:"ruleid,omitempty"`
	SHA1                    string   `json:"sha1,omitempty"`
	Src                     string   `json:"src,omitempty"`
	VendorType              string   `json:"vendortype,omitempty"`
	XSLT                    string   `json:"xslt,omitempty"`
}

type AppFWProfileFieldFormatBinding struct {
	AlertOnly            string `json:"alertonly,omitempty"`
	Comment              string `json:"comment,omitempty"`
	FieldFormat          string `json:"fieldformat,omitempty"`
	FieldFormatMaxLength int    `json:"fieldformatmaxlength,omitempty"`
	FieldFormatMinLength int    `json:"fieldformatminlength,omitempty"`
	FieldType            string `json:"fieldtype,omitempty"`
	FormActionURLFF      string `json:"formactionurl_ff,omitempty"`
	IsAutoDeployed       string `json:"isautodeployed,omitempty"`
	IsRegexFF            string `json:"isregex_ff,omitempty"`
	Name                 string `json:"name,omitempty"`
	ResourceID           string `json:"resourceid,omitempty"`
	RuleType             string `json:"ruletype,omitempty"`
	State                string `json:"state,omitempty"`
}

type AppFWProfileLogExpressionBinding struct {
	AlertOnly       string `json:"alertonly,omitempty"`
	ASLogExpression string `json:"as_logexpression,omitempty"`
	Comment         string `json:"comment,omitempty"`
	IsAutoDeployed  string `json:"isautodeployed,omitempty"`
	LogExpression   string `json:"logexpression,omitempty"`
	Name            string `json:"name,omitempty"`
	ResourceID      string `json:"resourceid,omitempty"`
	RuleType        string `json:"ruletype,omitempty"`
	State           string `json:"state,omitempty"`
}

type AppFWJSONErrorPage struct {
	Comment            string `json:"comment,omitempty"`
	Name               string `json:"name,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	Overwrite          bool   `json:"overwrite,omitempty"`
	Response           string `json:"response,omitempty"`
	Src                string `json:"src,omitempty"`
}

type AppFWProfileJSONDOSURLBinding struct {
	AlertOnly                   string `json:"alertonly,omitempty"`
	Comment                     string `json:"comment,omitempty"`
	IsAutoDeployed              string `json:"isautodeployed,omitempty"`
	JSONDOSURL                  string `json:"jsondosurl,omitempty"`
	JSONMaxArrayLength          int    `json:"jsonmaxarraylength,omitempty"`
	JSONMaxArrayLengthCheck     string `json:"jsonmaxarraylengthcheck,omitempty"`
	JSONMaxContainerDepth       int    `json:"jsonmaxcontainerdepth,omitempty"`
	JSONMaxContainerDepthCheck  string `json:"jsonmaxcontainerdepthcheck,omitempty"`
	JSONMaxDocumentLength       int    `json:"jsonmaxdocumentlength,omitempty"`
	JSONMaxDocumentLengthCheck  string `json:"jsonmaxdocumentlengthcheck,omitempty"`
	JSONMaxObjectKeyCount       int    `json:"jsonmaxobjectkeycount,omitempty"`
	JSONMaxObjectKeyCountCheck  string `json:"jsonmaxobjectkeycountcheck,omitempty"`
	JSONMaxObjectKeyLength      int    `json:"jsonmaxobjectkeylength,omitempty"`
	JSONMaxObjectKeyLengthCheck string `json:"jsonmaxobjectkeylengthcheck,omitempty"`
	JSONMaxStringLength         int    `json:"jsonmaxstringlength,omitempty"`
	JSONMaxStringLengthCheck    string `json:"jsonmaxstringlengthcheck,omitempty"`
	Name                        string `json:"name,omitempty"`
	ResourceID                  string `json:"resourceid,omitempty"`
	RuleType                    string `json:"ruletype,omitempty"`
	State                       string `json:"state,omitempty"`
}

type AppFWPolicy struct {
	Builtin            []string `json:"builtin,omitempty"`
	Comment            string   `json:"comment,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	LogAction          string   `json:"logaction,omitempty"`
	Name               string   `json:"name,omitempty"`
	NewName            string   `json:"newname,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	PolicyType         string   `json:"policytype,omitempty"`
	ProfileName        string   `json:"profilename,omitempty"`
	Rule               string   `json:"rule,omitempty"`
	UndefHits          int      `json:"undefhits,omitempty"`
}

type AppFWProfileFakeAccountBinding struct {
	AlertOnly        string `json:"alertonly,omitempty"`
	Comment          string `json:"comment,omitempty"`
	FakeAccount      string `json:"fakeaccount,omitempty"`
	FormExpression   string `json:"formexpression,omitempty"`
	FormURLFAD       string `json:"formurl_fad,omitempty"`
	IsAutoDeployed   string `json:"isautodeployed,omitempty"`
	IsFieldNameRegex string `json:"isfieldnameregex,omitempty"`
	Name             string `json:"name,omitempty"`
	ResourceID       string `json:"resourceid,omitempty"`
	RuleType         string `json:"ruletype,omitempty"`
	State            string `json:"state,omitempty"`
	Tag              string `json:"tag,omitempty"`
}

type AppFWProfileJSONCMDURLBinding struct {
	AlertOnly           string `json:"alertonly,omitempty"`
	ASValueExprJSONCMD  string `json:"as_value_expr_json_cmd,omitempty"`
	ASValueTypeJSONCMD  string `json:"as_value_type_json_cmd,omitempty"`
	Comment             string `json:"comment,omitempty"`
	IsAutoDeployed      string `json:"isautodeployed,omitempty"`
	IsKeyRegexJSONCMD   string `json:"iskeyregex_json_cmd,omitempty"`
	IsValueRegexJSONCMD string `json:"isvalueregex_json_cmd,omitempty"`
	JSONCMDURL          string `json:"jsoncmdurl,omitempty"`
	KeyNameJSONCMD      string `json:"keyname_json_cmd,omitempty"`
	Name                string `json:"name,omitempty"`
	ResourceID          string `json:"resourceid,omitempty"`
	RuleType            string `json:"ruletype,omitempty"`
	State               string `json:"state,omitempty"`
}

type AppFWConfidField struct {
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	FieldName          string  `json:"fieldname,omitempty"`
	IsRegex            string  `json:"isregex,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	State              string  `json:"state,omitempty"`
	URL                string  `json:"url,omitempty"`
}

type AppFWGlobalAuditNSLogPolicyBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	PolicyType             string `json:"policytype,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	State                  string `json:"state,omitempty"`
	TypeField              string `json:"type,omitempty"`
}

type AppFWProfileXMLValidationURLBinding struct {
	AlertOnly                string `json:"alertonly,omitempty"`
	Comment                  string `json:"comment,omitempty"`
	IsAutoDeployed           string `json:"isautodeployed,omitempty"`
	Name                     string `json:"name,omitempty"`
	ResourceID               string `json:"resourceid,omitempty"`
	RuleType                 string `json:"ruletype,omitempty"`
	State                    string `json:"state,omitempty"`
	XMLAdditionalSOAPHeaders string `json:"xmladditionalsoapheaders,omitempty"`
	XMLEndpointCheck         string `json:"xmlendpointcheck,omitempty"`
	XMLRequestSchema         string `json:"xmlrequestschema,omitempty"`
	XMLResponseSchema        string `json:"xmlresponseschema,omitempty"`
	XMLValidateResponse      string `json:"xmlvalidateresponse,omitempty"`
	XMLValidateSOAPEnvelope  string `json:"xmlvalidatesoapenvelope,omitempty"`
	XMLValidationURL         string `json:"xmlvalidationurl,omitempty"`
	XMLWSDL                  string `json:"xmlwsdl,omitempty"`
}

type AppFWProfileAppFWConfidFieldBinding struct {
	AlertOnly      string `json:"alertonly,omitempty"`
	CFFieldURL     string `json:"cffield_url,omitempty"`
	Comment        string `json:"comment,omitempty"`
	ConfidField    string `json:"confidfield,omitempty"`
	IsAutoDeployed string `json:"isautodeployed,omitempty"`
	IsRegexCFField string `json:"isregex_cffield,omitempty"`
	Name           string `json:"name,omitempty"`
	ResourceID     string `json:"resourceid,omitempty"`
	RuleType       string `json:"ruletype,omitempty"`
	State          string `json:"state,omitempty"`
}

type AppFWProfileCrossSiteScriptingBinding struct {
	AlertOnly          string `json:"alertonly,omitempty"`
	ASScanLocationXSS  string `json:"as_scan_location_xss,omitempty"`
	ASValueExprXSS     string `json:"as_value_expr_xss,omitempty"`
	ASValueTypeXSS     string `json:"as_value_type_xss,omitempty"`
	Comment            string `json:"comment,omitempty"`
	CrossSiteScripting string `json:"crosssitescripting,omitempty"`
	FormActionURLXSS   string `json:"formactionurl_xss,omitempty"`
	IsAutoDeployed     string `json:"isautodeployed,omitempty"`
	IsRegexXSS         string `json:"isregex_xss,omitempty"`
	IsValueRegexXSS    string `json:"isvalueregex_xss,omitempty"`
	Name               string `json:"name,omitempty"`
	ResourceID         string `json:"resourceid,omitempty"`
	RuleType           string `json:"ruletype,omitempty"`
	State              string `json:"state,omitempty"`
}

type AppFWTransactionRecords struct {
	AppFWSessionID            string  `json:"appfwsessionid,omitempty"`
	ClientIP                  string  `json:"clientip,omitempty"`
	Count                     float64 `json:"__count,omitempty"`
	DestIP                    string  `json:"destip,omitempty"`
	EndTime                   string  `json:"endtime,omitempty"`
	HTTPTransactionID         int     `json:"httptransactionid,omitempty"`
	NextGenAPIResource        string  `json:"_nextgenapiresource,omitempty"`
	NodeID                    int     `json:"nodeid,omitempty"`
	PacketEngineID            int     `json:"packetengineid,omitempty"`
	ProfileName               string  `json:"profilename,omitempty"`
	RequestContentLength      int     `json:"requestcontentlength,omitempty"`
	RequestMaxProcessingTime  int     `json:"requestmaxprocessingtime,omitempty"`
	RequestYields             int     `json:"requestyields,omitempty"`
	ResponseContentLength     int     `json:"responsecontentlength,omitempty"`
	ResponseMaxProcessingTime int     `json:"responsemaxprocessingtime,omitempty"`
	ResponseYields            int     `json:"responseyields,omitempty"`
	StartTime                 string  `json:"starttime,omitempty"`
	URL                       string  `json:"url,omitempty"`
}

type AppFWProfileBinding struct {
	AppFWProfileAppFWConfidFieldBinding       []any  `json:"appfwprofile_appfwconfidfield_binding,omitempty"`
	AppFWProfileBlockKeywordBinding           []any  `json:"appfwprofile_blockkeyword_binding,omitempty"`
	AppFWProfileBypassListBinding             []any  `json:"appfwprofile_bypasslist_binding,omitempty"`
	AppFWProfileCMDInjectionBinding           []any  `json:"appfwprofile_cmdinjection_binding,omitempty"`
	AppFWProfileContentTypeBinding            []any  `json:"appfwprofile_contenttype_binding,omitempty"`
	AppFWProfileCookieConsistencyBinding      []any  `json:"appfwprofile_cookieconsistency_binding,omitempty"`
	AppFWProfileCreditCardNumberBinding       []any  `json:"appfwprofile_creditcardnumber_binding,omitempty"`
	AppFWProfileCrossSiteScriptingBinding     []any  `json:"appfwprofile_crosssitescripting_binding,omitempty"`
	AppFWProfileCSRFTagBinding                []any  `json:"appfwprofile_csrftag_binding,omitempty"`
	AppFWProfileDenyListBinding               []any  `json:"appfwprofile_denylist_binding,omitempty"`
	AppFWProfileDenyURLBinding                []any  `json:"appfwprofile_denyurl_binding,omitempty"`
	AppFWProfileExcludeRESContentTypeBinding  []any  `json:"appfwprofile_excluderescontenttype_binding,omitempty"`
	AppFWProfileFakeAccountBinding            []any  `json:"appfwprofile_fakeaccount_binding,omitempty"`
	AppFWProfileFieldConsistencyBinding       []any  `json:"appfwprofile_fieldconsistency_binding,omitempty"`
	AppFWProfileFieldFormatBinding            []any  `json:"appfwprofile_fieldformat_binding,omitempty"`
	AppFWProfileFileUploadTypeBinding         []any  `json:"appfwprofile_fileuploadtype_binding,omitempty"`
	AppFWProfileGRPCValidationBinding         []any  `json:"appfwprofile_grpcvalidation_binding,omitempty"`
	AppFWProfileJSONBlockKeywordBinding       []any  `json:"appfwprofile_jsonblockkeyword_binding,omitempty"`
	AppFWProfileJSONCMDURLBinding             []any  `json:"appfwprofile_jsoncmdurl_binding,omitempty"`
	AppFWProfileJSONDOSURLBinding             []any  `json:"appfwprofile_jsondosurl_binding,omitempty"`
	AppFWProfileJSONSQLURLBinding             []any  `json:"appfwprofile_jsonsqlurl_binding,omitempty"`
	AppFWProfileJSONXSSURLBinding             []any  `json:"appfwprofile_jsonxssurl_binding,omitempty"`
	AppFWProfileLogExpressionBinding          []any  `json:"appfwprofile_logexpression_binding,omitempty"`
	AppFWProfileRESTValidationBinding         []any  `json:"appfwprofile_restvalidation_binding,omitempty"`
	AppFWProfileSafeObjectBinding             []any  `json:"appfwprofile_safeobject_binding,omitempty"`
	AppFWProfileSQLInjectionBinding           []any  `json:"appfwprofile_sqlinjection_binding,omitempty"`
	AppFWProfileStartURLBinding               []any  `json:"appfwprofile_starturl_binding,omitempty"`
	AppFWProfileTrustedLearningClientsBinding []any  `json:"appfwprofile_trustedlearningclients_binding,omitempty"`
	AppFWProfileXMLAttachmentURLBinding       []any  `json:"appfwprofile_xmlattachmenturl_binding,omitempty"`
	AppFWProfileXMLDOSURLBinding              []any  `json:"appfwprofile_xmldosurl_binding,omitempty"`
	AppFWProfileXMLSQLInjectionBinding        []any  `json:"appfwprofile_xmlsqlinjection_binding,omitempty"`
	AppFWProfileXMLValidationURLBinding       []any  `json:"appfwprofile_xmlvalidationurl_binding,omitempty"`
	AppFWProfileXMLWSIURLBinding              []any  `json:"appfwprofile_xmlwsiurl_binding,omitempty"`
	AppFWProfileXMLXSSBinding                 []any  `json:"appfwprofile_xmlxss_binding,omitempty"`
	Name                                      string `json:"name,omitempty"`
}

type AppFWProfileCSRFTagBinding struct {
	AlertOnly         string `json:"alertonly,omitempty"`
	Comment           string `json:"comment,omitempty"`
	CSRFFormActionURL string `json:"csrfformactionurl,omitempty"`
	CSRFTag           string `json:"csrftag,omitempty"`
	IsAutoDeployed    string `json:"isautodeployed,omitempty"`
	Name              string `json:"name,omitempty"`
	ResourceID        string `json:"resourceid,omitempty"`
	RuleType          string `json:"ruletype,omitempty"`
	State             string `json:"state,omitempty"`
}

type AppFWWSDL struct {
	Comment            string `json:"comment,omitempty"`
	Name               string `json:"name,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	Overwrite          bool   `json:"overwrite,omitempty"`
	Response           string `json:"response,omitempty"`
	Src                string `json:"src,omitempty"`
}

type AppFWPolicyLabel struct {
	Count                  float64 `json:"__count,omitempty"`
	Description            string  `json:"description,omitempty"`
	GotoPriorityExpression string  `json:"gotopriorityexpression,omitempty"`
	Hits                   int     `json:"hits,omitempty"`
	InvokeLabelName        string  `json:"invoke_labelname,omitempty"`
	LabelName              string  `json:"labelname,omitempty"`
	LabelType              string  `json:"labeltype,omitempty"`
	NewName                string  `json:"newname,omitempty"`
	NextGenAPIResource     string  `json:"_nextgenapiresource,omitempty"`
	NumPol                 int     `json:"numpol,omitempty"`
	PolicyLabelType        string  `json:"policylabeltype,omitempty"`
	PolicyType             string  `json:"policytype,omitempty"`
	Priority               int     `json:"priority,omitempty"`
}

type AppFWProtoFile struct {
	Comment            string `json:"comment,omitempty"`
	Name               string `json:"name,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	Overwrite          bool   `json:"overwrite,omitempty"`
	Response           string `json:"response,omitempty"`
	Src                string `json:"src,omitempty"`
}

type AppFWGlobalAuditSyslogPolicyBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	PolicyType             string `json:"policytype,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	State                  string `json:"state,omitempty"`
	TypeField              string `json:"type,omitempty"`
}

type AppFWGlobalBinding struct {
	AppFWGlobalAppFWPolicyBinding       []any `json:"appfwglobal_appfwpolicy_binding,omitempty"`
	AppFWGlobalAuditNSLogPolicyBinding  []any `json:"appfwglobal_auditnslogpolicy_binding,omitempty"`
	AppFWGlobalAuditSyslogPolicyBinding []any `json:"appfwglobal_auditsyslogpolicy_binding,omitempty"`
}

type AppFWProfileJSONXSSURLBinding struct {
	AlertOnly           string `json:"alertonly,omitempty"`
	ASValueExprJSONXSS  string `json:"as_value_expr_json_xss,omitempty"`
	ASValueTypeJSONXSS  string `json:"as_value_type_json_xss,omitempty"`
	Comment             string `json:"comment,omitempty"`
	IsAutoDeployed      string `json:"isautodeployed,omitempty"`
	IsKeyRegexJSONXSS   string `json:"iskeyregex_json_xss,omitempty"`
	IsValueRegexJSONXSS string `json:"isvalueregex_json_xss,omitempty"`
	JSONXSSURL          string `json:"jsonxssurl,omitempty"`
	KeyNameJSONXSS      string `json:"keyname_json_xss,omitempty"`
	Name                string `json:"name,omitempty"`
	ResourceID          string `json:"resourceid,omitempty"`
	RuleType            string `json:"ruletype,omitempty"`
	State               string `json:"state,omitempty"`
}

type AppFWProfileFileUploadTypeBinding struct {
	AlertOnly                 string   `json:"alertonly,omitempty"`
	ASFileUploadTypesURL      string   `json:"as_fileuploadtypes_url,omitempty"`
	Comment                   string   `json:"comment,omitempty"`
	FileType                  []string `json:"filetype,omitempty"`
	FileUploadType            string   `json:"fileuploadtype,omitempty"`
	IsAutoDeployed            string   `json:"isautodeployed,omitempty"`
	IsNameRegex               string   `json:"isnameregex,omitempty"`
	IsRegexFileUploadTypesURL string   `json:"isregex_fileuploadtypes_url,omitempty"`
	Name                      string   `json:"name,omitempty"`
	ResourceID                string   `json:"resourceid,omitempty"`
	RuleType                  string   `json:"ruletype,omitempty"`
	State                     string   `json:"state,omitempty"`
}

type AppFWProfileXMLAttachmentURLBinding struct {
	AlertOnly                     string `json:"alertonly,omitempty"`
	Comment                       string `json:"comment,omitempty"`
	IsAutoDeployed                string `json:"isautodeployed,omitempty"`
	Name                          string `json:"name,omitempty"`
	ResourceID                    string `json:"resourceid,omitempty"`
	RuleType                      string `json:"ruletype,omitempty"`
	State                         string `json:"state,omitempty"`
	XMLAttachmentContentType      string `json:"xmlattachmentcontenttype,omitempty"`
	XMLAttachmentContentTypeCheck string `json:"xmlattachmentcontenttypecheck,omitempty"`
	XMLAttachmentURL              string `json:"xmlattachmenturl,omitempty"`
	XMLMaxAttachmentSize          int    `json:"xmlmaxattachmentsize,omitempty"`
	XMLMaxAttachmentSizeCheck     string `json:"xmlmaxattachmentsizecheck,omitempty"`
}

type AppFWURLEncodedFormContentType struct {
	Builtin                        []string `json:"builtin,omitempty"`
	Count                          float64  `json:"__count,omitempty"`
	Feature                        string   `json:"feature,omitempty"`
	IsRegex                        string   `json:"isregex,omitempty"`
	NextGenAPIResource             string   `json:"_nextgenapiresource,omitempty"`
	URLEncodedFormContentTypeValue string   `json:"urlencodedformcontenttypevalue,omitempty"`
}

type AppFWPolicyBinding struct {
	AppFWPolicyAppFWGlobalBinding      []any  `json:"appfwpolicy_appfwglobal_binding,omitempty"`
	AppFWPolicyAppFWPolicyLabelBinding []any  `json:"appfwpolicy_appfwpolicylabel_binding,omitempty"`
	AppFWPolicyCSVServerBinding        []any  `json:"appfwpolicy_csvserver_binding,omitempty"`
	AppFWPolicyLBVServerBinding        []any  `json:"appfwpolicy_lbvserver_binding,omitempty"`
	AppFWPolicyVPNVServerBinding       []any  `json:"appfwpolicy_vpnvserver_binding,omitempty"`
	Name                               string `json:"name,omitempty"`
}

type AppFWProfileFieldConsistencyBinding struct {
	AlertOnly        string `json:"alertonly,omitempty"`
	Comment          string `json:"comment,omitempty"`
	FieldConsistency string `json:"fieldconsistency,omitempty"`
	FormActionURLFFC string `json:"formactionurl_ffc,omitempty"`
	IsAutoDeployed   string `json:"isautodeployed,omitempty"`
	IsRegexFFC       string `json:"isregex_ffc,omitempty"`
	Name             string `json:"name,omitempty"`
	ResourceID       string `json:"resourceid,omitempty"`
	RuleType         string `json:"ruletype,omitempty"`
	State            string `json:"state,omitempty"`
}

type AppFWXMLSchema struct {
	Comment            string `json:"comment,omitempty"`
	Name               string `json:"name,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	Overwrite          bool   `json:"overwrite,omitempty"`
	Response           string `json:"response,omitempty"`
	Src                string `json:"src,omitempty"`
}

type AppFWArchive struct {
	Comment            string `json:"comment,omitempty"`
	Name               string `json:"name,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	Response           string `json:"response,omitempty"`
	Src                string `json:"src,omitempty"`
	Target             string `json:"target,omitempty"`
}

type AppFWPolicyAppFWPolicyLabelBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type AppFWProfileBlockKeywordBinding struct {
	AlertOnly                      string `json:"alertonly,omitempty"`
	ASBlockKeywordFormURL          string `json:"as_blockkeyword_formurl,omitempty"`
	ASFieldNameIsRegexBlockKeyword string `json:"as_fieldname_isregex_blockkeyword,omitempty"`
	BlockKeyword                   string `json:"blockkeyword,omitempty"`
	BlockKeywordType               string `json:"blockkeywordtype,omitempty"`
	Comment                        string `json:"comment,omitempty"`
	FieldName                      string `json:"fieldname,omitempty"`
	IsAutoDeployed                 string `json:"isautodeployed,omitempty"`
	Name                           string `json:"name,omitempty"`
	ResourceID                     string `json:"resourceid,omitempty"`
	RuleType                       string `json:"ruletype,omitempty"`
	State                          string `json:"state,omitempty"`
}

type AppFWSettings struct {
	Builtin                  []string `json:"builtin,omitempty"`
	CEFLogging               string   `json:"ceflogging,omitempty"`
	CentralizedLearning      string   `json:"centralizedlearning,omitempty"`
	ClientIPLoggingHeader    string   `json:"clientiploggingheader,omitempty"`
	CookieFlags              string   `json:"cookieflags,omitempty"`
	CookiePostEncryptPrefix  string   `json:"cookiepostencryptprefix,omitempty"`
	DefaultProfile           string   `json:"defaultprofile,omitempty"`
	EntityDecoding           string   `json:"entitydecoding,omitempty"`
	Feature                  string   `json:"feature,omitempty"`
	GeoLocationLogging       string   `json:"geolocationlogging,omitempty"`
	ImportSizeLimit          int      `json:"importsizelimit,omitempty"`
	Learning                 string   `json:"learning,omitempty"`
	LearnRateLimit           int      `json:"learnratelimit,omitempty"`
	LogMalformedReq          string   `json:"logmalformedreq,omitempty"`
	MalformedReqAction       []string `json:"malformedreqaction,omitempty"`
	NextGenAPIResource       string   `json:"_nextgenapiresource,omitempty"`
	ProxyPassword            string   `json:"proxypassword,omitempty"`
	ProxyPort                int      `json:"proxyport,omitempty"`
	ProxyServer              string   `json:"proxyserver,omitempty"`
	ProxyUsername            string   `json:"proxyusername,omitempty"`
	SessionCookieName        string   `json:"sessioncookiename,omitempty"`
	SessionLifetime          int      `json:"sessionlifetime,omitempty"`
	SessionLimit             int      `json:"sessionlimit,omitempty"`
	SessionTimeout           int      `json:"sessiontimeout,omitempty"`
	SignatureAutoUpdate      string   `json:"signatureautoupdate,omitempty"`
	SignatureURL             string   `json:"signatureurl,omitempty"`
	UndefAction              string   `json:"undefaction,omitempty"`
	UseConfigurableSecretKey string   `json:"useconfigurablesecretkey,omitempty"`
}

type AppFWPolicyCSVServerBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type AppFWProfileCMDInjectionBinding struct {
	AlertOnly         string `json:"alertonly,omitempty"`
	ASScanLocationCMD string `json:"as_scan_location_cmd,omitempty"`
	ASValueExprCMD    string `json:"as_value_expr_cmd,omitempty"`
	ASValueTypeCMD    string `json:"as_value_type_cmd,omitempty"`
	CMDInjection      string `json:"cmdinjection,omitempty"`
	Comment           string `json:"comment,omitempty"`
	FormActionURLCMD  string `json:"formactionurl_cmd,omitempty"`
	IsAutoDeployed    string `json:"isautodeployed,omitempty"`
	IsRegexCMD        string `json:"isregex_cmd,omitempty"`
	IsValueRegexCMD   string `json:"isvalueregex_cmd,omitempty"`
	Name              string `json:"name,omitempty"`
	ResourceID        string `json:"resourceid,omitempty"`
	RuleType          string `json:"ruletype,omitempty"`
	State             string `json:"state,omitempty"`
}

type AppFWPolicyLabelPolicyBindingBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelName        string `json:"invoke_labelname,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type AppFWProfileStartURLBinding struct {
	AlertOnly      string `json:"alertonly,omitempty"`
	Comment        string `json:"comment,omitempty"`
	IsAutoDeployed string `json:"isautodeployed,omitempty"`
	Name           string `json:"name,omitempty"`
	ResourceID     string `json:"resourceid,omitempty"`
	RuleType       string `json:"ruletype,omitempty"`
	StartURL       string `json:"starturl,omitempty"`
	State          string `json:"state,omitempty"`
}

type AppFWProfileJSONBlockKeywordBinding struct {
	AlertOnly                  string `json:"alertonly,omitempty"`
	Comment                    string `json:"comment,omitempty"`
	IsAutoDeployed             string `json:"isautodeployed,omitempty"`
	IsKeyRegexJSONBlockKeyword string `json:"iskeyregex_json_blockkeyword,omitempty"`
	JSONBlockKeyword           string `json:"jsonblockkeyword,omitempty"`
	JSONBlockKeywordType       string `json:"jsonblockkeywordtype,omitempty"`
	JSONBlockKeywordURL        string `json:"jsonblockkeywordurl,omitempty"`
	KeyNameJSONBlockKeyword    string `json:"keyname_json_blockkeyword,omitempty"`
	Name                       string `json:"name,omitempty"`
	ResourceID                 string `json:"resourceid,omitempty"`
	RuleType                   string `json:"ruletype,omitempty"`
	State                      string `json:"state,omitempty"`
}

type AppFWFieldType struct {
	Builtin            []string `json:"builtin,omitempty"`
	Comment            string   `json:"comment,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Name               string   `json:"name,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	NoCharMaps         bool     `json:"nocharmaps,omitempty"`
	Priority           int      `json:"priority,omitempty"`
	Regex              string   `json:"regex,omitempty"`
}

type AppFWLearningData struct {
	ASScanLocationSQL      string  `json:"as_scan_location_sql,omitempty"`
	ASScanLocationXSS      string  `json:"as_scan_location_xss,omitempty"`
	ASValueExprSQL         string  `json:"as_value_expr_sql,omitempty"`
	ASValueExprXSS         string  `json:"as_value_expr_xss,omitempty"`
	ASValueTypeSQL         string  `json:"as_value_type_sql,omitempty"`
	ASValueTypeXSS         string  `json:"as_value_type_xss,omitempty"`
	ContentType            string  `json:"contenttype,omitempty"`
	CookieConsistency      string  `json:"cookieconsistency,omitempty"`
	Count                  float64 `json:"__count,omitempty"`
	CreditCardNumber       string  `json:"creditcardnumber,omitempty"`
	CreditCardNumberURL    string  `json:"creditcardnumberurl,omitempty"`
	CrossSiteScripting     string  `json:"crosssitescripting,omitempty"`
	CSRFFormOriginURL      string  `json:"csrfformoriginurl,omitempty"`
	CSRFTag                string  `json:"csrftag,omitempty"`
	Data                   string  `json:"data,omitempty"`
	FieldConsistency       string  `json:"fieldconsistency,omitempty"`
	FieldFormat            string  `json:"fieldformat,omitempty"`
	FieldFormatCharMapPCRE string  `json:"fieldformatcharmappcre,omitempty"`
	FieldFormatMaxLength   int     `json:"fieldformatmaxlength,omitempty"`
	FieldFormatMinLength   int     `json:"fieldformatminlength,omitempty"`
	FieldType              string  `json:"fieldtype,omitempty"`
	FormActionURLFF        string  `json:"formactionurl_ff,omitempty"`
	FormActionURLFFC       string  `json:"formactionurl_ffc,omitempty"`
	FormActionURLSQL       string  `json:"formactionurl_sql,omitempty"`
	FormActionURLXSS       string  `json:"formactionurl_xss,omitempty"`
	Hits                   int     `json:"hits,omitempty"`
	Name                   string  `json:"name,omitempty"`
	NextGenAPIResource     string  `json:"_nextgenapiresource,omitempty"`
	ProfileName            string  `json:"profilename,omitempty"`
	SecurityCheck          string  `json:"securitycheck,omitempty"`
	SQLInjection           string  `json:"sqlinjection,omitempty"`
	StartURL               string  `json:"starturl,omitempty"`
	Target                 string  `json:"target,omitempty"`
	TotalXMLRequests       bool    `json:"totalxmlrequests,omitempty"`
	URL                    string  `json:"url,omitempty"`
	Value                  string  `json:"value,omitempty"`
	ValueType              string  `json:"value_type,omitempty"`
	XMLAttachmentCheck     string  `json:"xmlattachmentcheck,omitempty"`
	XMLDOSCheck            string  `json:"xmldoscheck,omitempty"`
	XMLWSICheck            string  `json:"xmlwsicheck,omitempty"`
}

type AppFWPolicyLabelAppFWPolicyBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelName        string `json:"invoke_labelname,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type AppFWPolicyVPNVServerBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type AppFWProfileXMLSQLInjectionBinding struct {
	AlertOnly            string `json:"alertonly,omitempty"`
	ASScanLocationXMLSQL string `json:"as_scan_location_xmlsql,omitempty"`
	Comment              string `json:"comment,omitempty"`
	IsAutoDeployed       string `json:"isautodeployed,omitempty"`
	IsRegexXMLSQL        string `json:"isregex_xmlsql,omitempty"`
	Name                 string `json:"name,omitempty"`
	ResourceID           string `json:"resourceid,omitempty"`
	RuleType             string `json:"ruletype,omitempty"`
	State                string `json:"state,omitempty"`
	XMLSQLInjection      string `json:"xmlsqlinjection,omitempty"`
}

type AppFWGRPCWebJSONContentType struct {
	Count                       float64 `json:"__count,omitempty"`
	GRPCWebJSONContentTypeValue string  `json:"grpcwebjsoncontenttypevalue,omitempty"`
	IsRegex                     string  `json:"isregex,omitempty"`
	NextGenAPIResource          string  `json:"_nextgenapiresource,omitempty"`
}

type AppFWProfileTrustedLearningClientsBinding struct {
	AlertOnly              string `json:"alertonly,omitempty"`
	Comment                string `json:"comment,omitempty"`
	IsAutoDeployed         string `json:"isautodeployed,omitempty"`
	Name                   string `json:"name,omitempty"`
	ResourceID             string `json:"resourceid,omitempty"`
	RuleType               string `json:"ruletype,omitempty"`
	State                  string `json:"state,omitempty"`
	TrustedLearningClients string `json:"trustedlearningclients,omitempty"`
}

type AppFWProfileCreditCardNumberBinding struct {
	AlertOnly           string `json:"alertonly,omitempty"`
	Comment             string `json:"comment,omitempty"`
	CreditCardNumber    string `json:"creditcardnumber,omitempty"`
	CreditCardNumberURL string `json:"creditcardnumberurl,omitempty"`
	IsAutoDeployed      string `json:"isautodeployed,omitempty"`
	Name                string `json:"name,omitempty"`
	ResourceID          string `json:"resourceid,omitempty"`
	RuleType            string `json:"ruletype,omitempty"`
	State               string `json:"state,omitempty"`
}

type AppFWProfileExcludeRESContentTypeBinding struct {
	AlertOnly             string `json:"alertonly,omitempty"`
	Comment               string `json:"comment,omitempty"`
	ExcludeRESContentType string `json:"excluderescontenttype,omitempty"`
	IsAutoDeployed        string `json:"isautodeployed,omitempty"`
	Name                  string `json:"name,omitempty"`
	ResourceID            string `json:"resourceid,omitempty"`
	RuleType              string `json:"ruletype,omitempty"`
	State                 string `json:"state,omitempty"`
}

type AppFWJSONContentType struct {
	Builtin              []string `json:"builtin,omitempty"`
	Count                float64  `json:"__count,omitempty"`
	Feature              string   `json:"feature,omitempty"`
	IsRegex              string   `json:"isregex,omitempty"`
	JSONContentTypeValue string   `json:"jsoncontenttypevalue,omitempty"`
	NextGenAPIResource   string   `json:"_nextgenapiresource,omitempty"`
}

type AppFWProfileSQLInjectionBinding struct {
	AlertOnly         string `json:"alertonly,omitempty"`
	ASScanLocationSQL string `json:"as_scan_location_sql,omitempty"`
	ASValueExprSQL    string `json:"as_value_expr_sql,omitempty"`
	ASValueTypeSQL    string `json:"as_value_type_sql,omitempty"`
	Comment           string `json:"comment,omitempty"`
	FormActionURLSQL  string `json:"formactionurl_sql,omitempty"`
	IsAutoDeployed    string `json:"isautodeployed,omitempty"`
	IsRegexSQL        string `json:"isregex_sql,omitempty"`
	IsValueRegexSQL   string `json:"isvalueregex_sql,omitempty"`
	Name              string `json:"name,omitempty"`
	ResourceID        string `json:"resourceid,omitempty"`
	RuleType          string `json:"ruletype,omitempty"`
	SQLInjection      string `json:"sqlinjection,omitempty"`
	State             string `json:"state,omitempty"`
}

type AppFWProfileRESTValidationBinding struct {
	AlertOnly            string `json:"alertonly,omitempty"`
	Comment              string `json:"comment,omitempty"`
	IsAutoDeployed       string `json:"isautodeployed,omitempty"`
	Name                 string `json:"name,omitempty"`
	ResourceID           string `json:"resourceid,omitempty"`
	RESTValidationAction string `json:"rest_validation_action,omitempty"`
	RESTValidation       string `json:"restvalidation,omitempty"`
	RuleType             string `json:"ruletype,omitempty"`
	State                string `json:"state,omitempty"`
}

type AppFWGRPCContentType struct {
	Builtin              []string `json:"builtin,omitempty"`
	Count                float64  `json:"__count,omitempty"`
	Feature              string   `json:"feature,omitempty"`
	GRPCContentTypeValue string   `json:"grpccontenttypevalue,omitempty"`
	IsRegex              string   `json:"isregex,omitempty"`
	NextGenAPIResource   string   `json:"_nextgenapiresource,omitempty"`
}

type AppFWProfileBypassListBinding struct {
	AlertOnly             string `json:"alertonly,omitempty"`
	ASBypassList          string `json:"as_bypass_list,omitempty"`
	ASBypassListAction    string `json:"as_bypass_list_action,omitempty"`
	ASBypassListLocation  string `json:"as_bypass_list_location,omitempty"`
	ASBypassListValueType string `json:"as_bypass_list_value_type,omitempty"`
	Comment               string `json:"comment,omitempty"`
	IsAutoDeployed        string `json:"isautodeployed,omitempty"`
	Name                  string `json:"name,omitempty"`
	ResourceID            string `json:"resourceid,omitempty"`
	RuleType              string `json:"ruletype,omitempty"`
	State                 string `json:"state,omitempty"`
}

type AppFWHTMLErrorPage struct {
	Comment            string `json:"comment,omitempty"`
	Name               string `json:"name,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	Overwrite          bool   `json:"overwrite,omitempty"`
	Response           string `json:"response,omitempty"`
	Src                string `json:"src,omitempty"`
}

type AppFWProfileSafeObjectBinding struct {
	Action         []string `json:"action,omitempty"`
	AlertOnly      string   `json:"alertonly,omitempty"`
	ASExpression   string   `json:"as_expression,omitempty"`
	Comment        string   `json:"comment,omitempty"`
	IsAutoDeployed string   `json:"isautodeployed,omitempty"`
	MaxMatchLength int      `json:"maxmatchlength,omitempty"`
	Name           string   `json:"name,omitempty"`
	ResourceID     string   `json:"resourceid,omitempty"`
	RuleType       string   `json:"ruletype,omitempty"`
	SafeObject     string   `json:"safeobject,omitempty"`
	State          string   `json:"state,omitempty"`
}

type AppFWProfileGRPCValidationBinding struct {
	AlertOnly                 string `json:"alertonly,omitempty"`
	Comment                   string `json:"comment,omitempty"`
	GRPCRelaxValidationAction string `json:"grpc_relax_validation_action,omitempty"`
	GRPCValidation            string `json:"grpcvalidation,omitempty"`
	IsAutoDeployed            string `json:"isautodeployed,omitempty"`
	Name                      string `json:"name,omitempty"`
	ResourceID                string `json:"resourceid,omitempty"`
	RuleType                  string `json:"ruletype,omitempty"`
	State                     string `json:"state,omitempty"`
}

type AppFWProfileXMLWSIURLBinding struct {
	AlertOnly      string `json:"alertonly,omitempty"`
	Comment        string `json:"comment,omitempty"`
	IsAutoDeployed string `json:"isautodeployed,omitempty"`
	Name           string `json:"name,omitempty"`
	ResourceID     string `json:"resourceid,omitempty"`
	RuleType       string `json:"ruletype,omitempty"`
	State          string `json:"state,omitempty"`
	XMLWSIChecks   string `json:"xmlwsichecks,omitempty"`
	XMLWSIURL      string `json:"xmlwsiurl,omitempty"`
}

type AppFWPolicyAppFWGlobalBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type AppFWProfileDenyListBinding struct {
	AlertOnly           string   `json:"alertonly,omitempty"`
	ASDenyList          string   `json:"as_deny_list,omitempty"`
	ASDenyListAction    []string `json:"as_deny_list_action,omitempty"`
	ASDenyListLocation  string   `json:"as_deny_list_location,omitempty"`
	ASDenyListValueType string   `json:"as_deny_list_value_type,omitempty"`
	Comment             string   `json:"comment,omitempty"`
	IsAutoDeployed      string   `json:"isautodeployed,omitempty"`
	Name                string   `json:"name,omitempty"`
	ResourceID          string   `json:"resourceid,omitempty"`
	RuleType            string   `json:"ruletype,omitempty"`
	State               string   `json:"state,omitempty"`
}

type AppFWPolicyLabelBinding struct {
	AppFWPolicyLabelAppFWPolicyBinding   []any  `json:"appfwpolicylabel_appfwpolicy_binding,omitempty"`
	AppFWPolicyLabelPolicyBindingBinding []any  `json:"appfwpolicylabel_policybinding_binding,omitempty"`
	LabelName                            string `json:"labelname,omitempty"`
}

type AppFWProfile struct {
	AddCookieFlags                             string   `json:"addcookieflags,omitempty"`
	APISpec                                    string   `json:"apispec,omitempty"`
	ArchiveName                                string   `json:"archivename,omitempty"`
	ASProfBypassListEnable                     string   `json:"as_prof_bypass_list_enable,omitempty"`
	ASProfDenyListEnable                       string   `json:"as_prof_deny_list_enable,omitempty"`
	Augment                                    bool     `json:"augment,omitempty"`
	BlockKeywordAction                         []string `json:"blockkeywordaction,omitempty"`
	BufferOverflowAction                       []string `json:"bufferoverflowaction,omitempty"`
	BufferOverflowMaxCookieLength              int      `json:"bufferoverflowmaxcookielength,omitempty"`
	BufferOverflowMaxHeaderLength              int      `json:"bufferoverflowmaxheaderlength,omitempty"`
	BufferOverflowMaxQueryLength               int      `json:"bufferoverflowmaxquerylength,omitempty"`
	BufferOverflowMaxTotalHeaderLength         int      `json:"bufferoverflowmaxtotalheaderlength,omitempty"`
	BufferOverflowMaxURLLength                 int      `json:"bufferoverflowmaxurllength,omitempty"`
	Builtin                                    bool     `json:"builtin,omitempty"`
	CanonicalizeHTMLResponse                   string   `json:"canonicalizehtmlresponse,omitempty"`
	CEFLogging                                 string   `json:"ceflogging,omitempty"`
	CheckRequestHeaders                        string   `json:"checkrequestheaders,omitempty"`
	ClientIPExpression                         string   `json:"clientipexpression,omitempty"`
	CMDInjectionAction                         []string `json:"cmdinjectionaction,omitempty"`
	CMDInjectionGrammar                        string   `json:"cmdinjectiongrammar,omitempty"`
	CMDInjectionType                           string   `json:"cmdinjectiontype,omitempty"`
	Comment                                    string   `json:"comment,omitempty"`
	ContentTypeAction                          []string `json:"contenttypeaction,omitempty"`
	CookieConsistencyAction                    []string `json:"cookieconsistencyaction,omitempty"`
	CookieEncryption                           string   `json:"cookieencryption,omitempty"`
	CookieHijackingAction                      []string `json:"cookiehijackingaction,omitempty"`
	CookieProxying                             string   `json:"cookieproxying,omitempty"`
	CookieSameSiteAttribute                    string   `json:"cookiesamesiteattribute,omitempty"`
	CookieTransforms                           string   `json:"cookietransforms,omitempty"`
	Count                                      float64  `json:"__count,omitempty"`
	CreditCard                                 []string `json:"creditcard,omitempty"`
	CreditCardAction                           []string `json:"creditcardaction,omitempty"`
	CreditCardMaxAllowed                       int      `json:"creditcardmaxallowed,omitempty"`
	CreditCardXOut                             string   `json:"creditcardxout,omitempty"`
	CrossSiteScriptingAction                   []string `json:"crosssitescriptingaction,omitempty"`
	CrossSiteScriptingCheckCompleteURLs        string   `json:"crosssitescriptingcheckcompleteurls,omitempty"`
	CrossSiteScriptingTransformUnsafeHTML      string   `json:"crosssitescriptingtransformunsafehtml,omitempty"`
	CSRFTag                                    string   `json:"csrftag,omitempty"`
	CSRFTagAction                              []string `json:"csrftagaction,omitempty"`
	CustomSettings                             string   `json:"customsettings,omitempty"`
	DefaultCharset                             string   `json:"defaultcharset,omitempty"`
	DefaultFieldFormatMaxLength                int      `json:"defaultfieldformatmaxlength,omitempty"`
	DefaultFieldFormatMaxOccurrences           int      `json:"defaultfieldformatmaxoccurrences,omitempty"`
	DefaultFieldFormatMinLength                int      `json:"defaultfieldformatminlength,omitempty"`
	DefaultFieldFormatType                     string   `json:"defaultfieldformattype,omitempty"`
	Defaults                                   string   `json:"defaults,omitempty"`
	DenyURLAction                              []string `json:"denyurlaction,omitempty"`
	DoSecureCreditCardLogging                  string   `json:"dosecurecreditcardlogging,omitempty"`
	DynamicLearning                            []string `json:"dynamiclearning,omitempty"`
	EnableFormTagging                          string   `json:"enableformtagging,omitempty"`
	ErrorURL                                   string   `json:"errorurl,omitempty"`
	ExcludeFileUploadFromChecks                string   `json:"excludefileuploadfromchecks,omitempty"`
	ExemptClosureURLsFromSecurityChecks        string   `json:"exemptclosureurlsfromsecuritychecks,omitempty"`
	FakeAccountDetection                       string   `json:"fakeaccountdetection,omitempty"`
	FieldConsistencyAction                     []string `json:"fieldconsistencyaction,omitempty"`
	FieldFormatAction                          []string `json:"fieldformataction,omitempty"`
	FieldScan                                  string   `json:"fieldscan,omitempty"`
	FieldScanLimit                             int      `json:"fieldscanlimit,omitempty"`
	FileUploadMaxNum                           int      `json:"fileuploadmaxnum,omitempty"`
	FileUploadTypesAction                      []string `json:"fileuploadtypesaction,omitempty"`
	GeoLocationLogging                         string   `json:"geolocationlogging,omitempty"`
	GRPCAction                                 []string `json:"grpcaction,omitempty"`
	HTMLErrorObject                            string   `json:"htmlerrorobject,omitempty"`
	HTMLErrorStatusCode                        int      `json:"htmlerrorstatuscode,omitempty"`
	HTMLErrorStatusMessage                     string   `json:"htmlerrorstatusmessage,omitempty"`
	ImportProfileName                          string   `json:"importprofilename,omitempty"`
	InferContentTypeXMLPayloadAction           []string `json:"infercontenttypexmlpayloadaction,omitempty"`
	InsertCookieSameSiteAttribute              string   `json:"insertcookiesamesiteattribute,omitempty"`
	InspectContentTypes                        []string `json:"inspectcontenttypes,omitempty"`
	InspectQueryContentTypes                   []string `json:"inspectquerycontenttypes,omitempty"`
	InvalidPercentHandling                     string   `json:"invalidpercenthandling,omitempty"`
	JSONBlockKeywordAction                     []string `json:"jsonblockkeywordaction,omitempty"`
	JSONCMDInjectionAction                     []string `json:"jsoncmdinjectionaction,omitempty"`
	JSONCMDInjectionGrammar                    string   `json:"jsoncmdinjectiongrammar,omitempty"`
	JSONCMDInjectionType                       string   `json:"jsoncmdinjectiontype,omitempty"`
	JSONDOSAction                              []string `json:"jsondosaction,omitempty"`
	JSONErrorObject                            string   `json:"jsonerrorobject,omitempty"`
	JSONErrorStatusCode                        int      `json:"jsonerrorstatuscode,omitempty"`
	JSONErrorStatusMessage                     string   `json:"jsonerrorstatusmessage,omitempty"`
	JSONFieldScan                              string   `json:"jsonfieldscan,omitempty"`
	JSONFieldScanLimit                         int      `json:"jsonfieldscanlimit,omitempty"`
	JSONMessageScan                            string   `json:"jsonmessagescan,omitempty"`
	JSONMessageScanLimit                       int      `json:"jsonmessagescanlimit,omitempty"`
	JSONSQLInjectionAction                     []string `json:"jsonsqlinjectionaction,omitempty"`
	JSONSQLInjectionGrammar                    string   `json:"jsonsqlinjectiongrammar,omitempty"`
	JSONSQLInjectionType                       string   `json:"jsonsqlinjectiontype,omitempty"`
	JSONXSSAction                              []string `json:"jsonxssaction,omitempty"`
	Learning                                   string   `json:"learning,omitempty"`
	LogEveryPolicyHit                          string   `json:"logeverypolicyhit,omitempty"`
	MatchURLString                             string   `json:"matchurlstring,omitempty"`
	MessageScan                                string   `json:"messagescan,omitempty"`
	MessageScanLimit                           int      `json:"messagescanlimit,omitempty"`
	MessageScanLimitContentTypes               []string `json:"messagescanlimitcontenttypes,omitempty"`
	MultipleHeaderAction                       []string `json:"multipleheaderaction,omitempty"`
	Name                                       string   `json:"name,omitempty"`
	NextGenAPIResource                         string   `json:"_nextgenapiresource,omitempty"`
	OptimizePartialReqs                        string   `json:"optimizepartialreqs,omitempty"`
	Overwrite                                  bool     `json:"overwrite,omitempty"`
	PercentDecodeRecursively                   string   `json:"percentdecoderecursively,omitempty"`
	PostBodyLimit                              int      `json:"postbodylimit,omitempty"`
	PostBodyLimitAction                        []string `json:"postbodylimitaction,omitempty"`
	PostBodyLimitSignature                     int      `json:"postbodylimitsignature,omitempty"`
	ProtoFileObject                            string   `json:"protofileobject,omitempty"`
	RefererHeaderCheck                         string   `json:"refererheadercheck,omitempty"`
	RelaxationRules                            bool     `json:"relaxationrules,omitempty"`
	ReplaceURLString                           string   `json:"replaceurlstring,omitempty"`
	RequestContentType                         string   `json:"requestcontenttype,omitempty"`
	ResponseContentType                        string   `json:"responsecontenttype,omitempty"`
	RESTAction                                 []string `json:"restaction,omitempty"`
	RFCProfile                                 string   `json:"rfcprofile,omitempty"`
	SemicolonFieldSeparator                    string   `json:"semicolonfieldseparator,omitempty"`
	SessionCookieName                          string   `json:"sessioncookiename,omitempty"`
	SessionlessFieldConsistency                string   `json:"sessionlessfieldconsistency,omitempty"`
	SessionlessURLClosure                      string   `json:"sessionlessurlclosure,omitempty"`
	Signatures                                 string   `json:"signatures,omitempty"`
	SQLInjectionAction                         []string `json:"sqlinjectionaction,omitempty"`
	SQLInjectionCheckSQLWildChars              string   `json:"sqlinjectionchecksqlwildchars,omitempty"`
	SQLInjectionGrammar                        string   `json:"sqlinjectiongrammar,omitempty"`
	SQLInjectionOnlyCheckFieldsWithSQLChars    string   `json:"sqlinjectiononlycheckfieldswithsqlchars,omitempty"`
	SQLInjectionParseComments                  string   `json:"sqlinjectionparsecomments,omitempty"`
	SQLInjectionRuleType                       string   `json:"sqlinjectionruletype,omitempty"`
	SQLInjectionTransformSpecialChars          string   `json:"sqlinjectiontransformspecialchars,omitempty"`
	SQLInjectionType                           string   `json:"sqlinjectiontype,omitempty"`
	StartURLAction                             []string `json:"starturlaction,omitempty"`
	StartURLClosure                            string   `json:"starturlclosure,omitempty"`
	State                                      string   `json:"state,omitempty"`
	Streaming                                  string   `json:"streaming,omitempty"`
	StripComments                              string   `json:"stripcomments,omitempty"`
	StripHTMLComments                          string   `json:"striphtmlcomments,omitempty"`
	StripXMLComments                           string   `json:"stripxmlcomments,omitempty"`
	Trace                                      string   `json:"trace,omitempty"`
	TypeField                                  []string `json:"type,omitempty"`
	URLDecodeRequestCookies                    string   `json:"urldecoderequestcookies,omitempty"`
	UseHTMLErrorObject                         string   `json:"usehtmlerrorobject,omitempty"`
	VerboseLogLevel                            string   `json:"verboseloglevel,omitempty"`
	XMLAttachmentAction                        []string `json:"xmlattachmentaction,omitempty"`
	XMLDOSAction                               []string `json:"xmldosaction,omitempty"`
	XMLErrorObject                             string   `json:"xmlerrorobject,omitempty"`
	XMLErrorStatusCode                         int      `json:"xmlerrorstatuscode,omitempty"`
	XMLErrorStatusMessage                      string   `json:"xmlerrorstatusmessage,omitempty"`
	XMLFormatAction                            []string `json:"xmlformataction,omitempty"`
	XMLSOAPFaultAction                         []string `json:"xmlsoapfaultaction,omitempty"`
	XMLSQLInjectionAction                      []string `json:"xmlsqlinjectionaction,omitempty"`
	XMLSQLInjectionCheckSQLWildChars           string   `json:"xmlsqlinjectionchecksqlwildchars,omitempty"`
	XMLSQLInjectionOnlyCheckFieldsWithSQLChars string   `json:"xmlsqlinjectiononlycheckfieldswithsqlchars,omitempty"`
	XMLSQLInjectionParseComments               string   `json:"xmlsqlinjectionparsecomments,omitempty"`
	XMLSQLInjectionType                        string   `json:"xmlsqlinjectiontype,omitempty"`
	XMLValidationAction                        []string `json:"xmlvalidationaction,omitempty"`
	XMLWSIAction                               []string `json:"xmlwsiaction,omitempty"`
	XMLXSSAction                               []string `json:"xmlxssaction,omitempty"`
}

type AppFWXMLContentType struct {
	Builtin             []string `json:"builtin,omitempty"`
	Count               float64  `json:"__count,omitempty"`
	Feature             string   `json:"feature,omitempty"`
	IsRegex             string   `json:"isregex,omitempty"`
	NextGenAPIResource  string   `json:"_nextgenapiresource,omitempty"`
	XMLContentTypeValue string   `json:"xmlcontenttypevalue,omitempty"`
}

type AppFWMultipartFormContentType struct {
	Builtin                       []string `json:"builtin,omitempty"`
	Count                         float64  `json:"__count,omitempty"`
	Feature                       string   `json:"feature,omitempty"`
	IsRegex                       string   `json:"isregex,omitempty"`
	MultipartFormContentTypeValue string   `json:"multipartformcontenttypevalue,omitempty"`
	NextGenAPIResource            string   `json:"_nextgenapiresource,omitempty"`
}

type AppFWLearningSettings struct {
	ContentTypeAutoDeployGracePeriod        int     `json:"contenttypeautodeploygraceperiod,omitempty"`
	ContentTypeMinThreshold                 int     `json:"contenttypeminthreshold,omitempty"`
	ContentTypePercentThreshold             int     `json:"contenttypepercentthreshold,omitempty"`
	CookieConsistencyAutoDeployGracePeriod  int     `json:"cookieconsistencyautodeploygraceperiod,omitempty"`
	CookieConsistencyMinThreshold           int     `json:"cookieconsistencyminthreshold,omitempty"`
	CookieConsistencyPercentThreshold       int     `json:"cookieconsistencypercentthreshold,omitempty"`
	Count                                   float64 `json:"__count,omitempty"`
	CreditCardNumberMinThreshold            int     `json:"creditcardnumberminthreshold,omitempty"`
	CreditCardNumberPercentThreshold        int     `json:"creditcardnumberpercentthreshold,omitempty"`
	CrossSiteScriptingAutoDeployGracePeriod int     `json:"crosssitescriptingautodeploygraceperiod,omitempty"`
	CrossSiteScriptingMinThreshold          int     `json:"crosssitescriptingminthreshold,omitempty"`
	CrossSiteScriptingPercentThreshold      int     `json:"crosssitescriptingpercentthreshold,omitempty"`
	CSRFTagAutoDeployGracePeriod            int     `json:"csrftagautodeploygraceperiod,omitempty"`
	CSRFTagMinThreshold                     int     `json:"csrftagminthreshold,omitempty"`
	CSRFTagPercentThreshold                 int     `json:"csrftagpercentthreshold,omitempty"`
	FieldConsistencyAutoDeployGracePeriod   int     `json:"fieldconsistencyautodeploygraceperiod,omitempty"`
	FieldConsistencyMinThreshold            int     `json:"fieldconsistencyminthreshold,omitempty"`
	FieldConsistencyPercentThreshold        int     `json:"fieldconsistencypercentthreshold,omitempty"`
	FieldFormatAutoDeployGracePeriod        int     `json:"fieldformatautodeploygraceperiod,omitempty"`
	FieldFormatMinThreshold                 int     `json:"fieldformatminthreshold,omitempty"`
	FieldFormatPercentThreshold             int     `json:"fieldformatpercentthreshold,omitempty"`
	NextGenAPIResource                      string  `json:"_nextgenapiresource,omitempty"`
	ProfileName                             string  `json:"profilename,omitempty"`
	SQLInjectionAutoDeployGracePeriod       int     `json:"sqlinjectionautodeploygraceperiod,omitempty"`
	SQLInjectionMinThreshold                int     `json:"sqlinjectionminthreshold,omitempty"`
	SQLInjectionPercentThreshold            int     `json:"sqlinjectionpercentthreshold,omitempty"`
	StartURLAutoDeployGracePeriod           int     `json:"starturlautodeploygraceperiod,omitempty"`
	StartURLMinThreshold                    int     `json:"starturlminthreshold,omitempty"`
	StartURLPercentThreshold                int     `json:"starturlpercentthreshold,omitempty"`
	XMLAttachmentMinThreshold               int     `json:"xmlattachmentminthreshold,omitempty"`
	XMLAttachmentPercentThreshold           int     `json:"xmlattachmentpercentthreshold,omitempty"`
	XMLWSIMinThreshold                      int     `json:"xmlwsiminthreshold,omitempty"`
	XMLWSIPercentThreshold                  int     `json:"xmlwsipercentthreshold,omitempty"`
}

type AppFWProfileXMLXSSBinding struct {
	AlertOnly            string `json:"alertonly,omitempty"`
	ASScanLocationXMLXSS string `json:"as_scan_location_xmlxss,omitempty"`
	Comment              string `json:"comment,omitempty"`
	IsAutoDeployed       string `json:"isautodeployed,omitempty"`
	IsRegexXMLXSS        string `json:"isregex_xmlxss,omitempty"`
	Name                 string `json:"name,omitempty"`
	ResourceID           string `json:"resourceid,omitempty"`
	RuleType             string `json:"ruletype,omitempty"`
	State                string `json:"state,omitempty"`
	XMLXSS               string `json:"xmlxss,omitempty"`
}

type AppFWProfileContentTypeBinding struct {
	AlertOnly      string `json:"alertonly,omitempty"`
	Comment        string `json:"comment,omitempty"`
	ContentType    string `json:"contenttype,omitempty"`
	IsAutoDeployed string `json:"isautodeployed,omitempty"`
	Name           string `json:"name,omitempty"`
	ResourceID     string `json:"resourceid,omitempty"`
	RuleType       string `json:"ruletype,omitempty"`
	State          string `json:"state,omitempty"`
}

type AppFWProfileJSONSQLURLBinding struct {
	AlertOnly           string `json:"alertonly,omitempty"`
	ASValueExprJSONSQL  string `json:"as_value_expr_json_sql,omitempty"`
	ASValueTypeJSONSQL  string `json:"as_value_type_json_sql,omitempty"`
	Comment             string `json:"comment,omitempty"`
	IsAutoDeployed      string `json:"isautodeployed,omitempty"`
	IsKeyRegexJSONSQL   string `json:"iskeyregex_json_sql,omitempty"`
	IsValueRegexJSONSQL string `json:"isvalueregex_json_sql,omitempty"`
	JSONSQLURL          string `json:"jsonsqlurl,omitempty"`
	KeyNameJSONSQL      string `json:"keyname_json_sql,omitempty"`
	Name                string `json:"name,omitempty"`
	ResourceID          string `json:"resourceid,omitempty"`
	RuleType            string `json:"ruletype,omitempty"`
	State               string `json:"state,omitempty"`
}

type AppFWCustomSettings struct {
	Name   string `json:"name,omitempty"`
	Target string `json:"target,omitempty"`
}
