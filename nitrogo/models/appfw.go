package models

// appfw configuration structs
type Appfwxmlerrorpage struct {
	Comment string `json:"comment,omitempty"`
	Name string `json:"name,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Overwrite bool `json:"overwrite,omitempty"`
	Response string `json:"response,omitempty"`
	Src string `json:"src,omitempty"`
}

type AppfwpolicyLbvserverBinding struct {
	Activepolicy int `json:"activepolicy,omitempty"`
	Boundto string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname string `json:"labelname,omitempty"`
	Labeltype string `json:"labeltype,omitempty"`
	Name string `json:"name,omitempty"`
	Priority int `json:"priority,omitempty"`
}

type Appfwgrpcwebtextcontenttype struct {
	Count float64 `json:"__count,omitempty"`
	Grpcwebtextcontenttypevalue string `json:"grpcwebtextcontenttypevalue,omitempty"`
	Isregex string `json:"isregex,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type AppfwglobalAppfwpolicyBinding struct {
	Flowtype int `json:"flowtype,omitempty"`
	Globalbindtype string `json:"globalbindtype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke bool `json:"invoke,omitempty"`
	Labelname string `json:"labelname,omitempty"`
	Labeltype string `json:"labeltype,omitempty"`
	Numpol int `json:"numpol,omitempty"`
	Policyname string `json:"policyname,omitempty"`
	Policytype string `json:"policytype,omitempty"`
	Priority int `json:"priority,omitempty"`
	State string `json:"state,omitempty"`
	TypeField string `json:"type,omitempty"`
}

type AppfwprofileCookieconsistencyBinding struct {
	Alertonly string `json:"alertonly,omitempty"`
	Comment string `json:"comment,omitempty"`
	Cookieconsistency string `json:"cookieconsistency,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	Isregex string `json:"isregex,omitempty"`
	Name string `json:"name,omitempty"`
	Resourceid string `json:"resourceid,omitempty"`
	Ruletype string `json:"ruletype,omitempty"`
	State string `json:"state,omitempty"`
}

type AppfwprofileXmldosurlBinding struct {
	Alertonly string `json:"alertonly,omitempty"`
	Comment string `json:"comment,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	Name string `json:"name,omitempty"`
	Resourceid string `json:"resourceid,omitempty"`
	Ruletype string `json:"ruletype,omitempty"`
	State string `json:"state,omitempty"`
	Xmlblockdtd string `json:"xmlblockdtd,omitempty"`
	Xmlblockexternalentities string `json:"xmlblockexternalentities,omitempty"`
	Xmlblockpi string `json:"xmlblockpi,omitempty"`
	Xmldosurl string `json:"xmldosurl,omitempty"`
	Xmlmaxattributenamelength int `json:"xmlmaxattributenamelength,omitempty"`
	Xmlmaxattributenamelengthcheck string `json:"xmlmaxattributenamelengthcheck,omitempty"`
	Xmlmaxattributes int `json:"xmlmaxattributes,omitempty"`
	Xmlmaxattributescheck string `json:"xmlmaxattributescheck,omitempty"`
	Xmlmaxattributevaluelength int `json:"xmlmaxattributevaluelength,omitempty"`
	Xmlmaxattributevaluelengthcheck string `json:"xmlmaxattributevaluelengthcheck,omitempty"`
	Xmlmaxchardatalength int `json:"xmlmaxchardatalength,omitempty"`
	Xmlmaxchardatalengthcheck string `json:"xmlmaxchardatalengthcheck,omitempty"`
	Xmlmaxelementchildren int `json:"xmlmaxelementchildren,omitempty"`
	Xmlmaxelementchildrencheck string `json:"xmlmaxelementchildrencheck,omitempty"`
	Xmlmaxelementdepth int `json:"xmlmaxelementdepth,omitempty"`
	Xmlmaxelementdepthcheck string `json:"xmlmaxelementdepthcheck,omitempty"`
	Xmlmaxelementnamelength int `json:"xmlmaxelementnamelength,omitempty"`
	Xmlmaxelementnamelengthcheck string `json:"xmlmaxelementnamelengthcheck,omitempty"`
	Xmlmaxelements int `json:"xmlmaxelements,omitempty"`
	Xmlmaxelementscheck string `json:"xmlmaxelementscheck,omitempty"`
	Xmlmaxentityexpansiondepth int `json:"xmlmaxentityexpansiondepth,omitempty"`
	Xmlmaxentityexpansiondepthcheck string `json:"xmlmaxentityexpansiondepthcheck,omitempty"`
	Xmlmaxentityexpansions int `json:"xmlmaxentityexpansions,omitempty"`
	Xmlmaxentityexpansionscheck string `json:"xmlmaxentityexpansionscheck,omitempty"`
	Xmlmaxfilesize int `json:"xmlmaxfilesize,omitempty"`
	Xmlmaxfilesizecheck string `json:"xmlmaxfilesizecheck,omitempty"`
	Xmlmaxnamespaces int `json:"xmlmaxnamespaces,omitempty"`
	Xmlmaxnamespacescheck string `json:"xmlmaxnamespacescheck,omitempty"`
	Xmlmaxnamespaceurilength int `json:"xmlmaxnamespaceurilength,omitempty"`
	Xmlmaxnamespaceurilengthcheck string `json:"xmlmaxnamespaceurilengthcheck,omitempty"`
	Xmlmaxnodes int `json:"xmlmaxnodes,omitempty"`
	Xmlmaxnodescheck string `json:"xmlmaxnodescheck,omitempty"`
	Xmlmaxsoaparrayrank int `json:"xmlmaxsoaparrayrank,omitempty"`
	Xmlmaxsoaparraysize int `json:"xmlmaxsoaparraysize,omitempty"`
	Xmlminfilesize int `json:"xmlminfilesize,omitempty"`
	Xmlminfilesizecheck string `json:"xmlminfilesizecheck,omitempty"`
	Xmlsoaparraycheck string `json:"xmlsoaparraycheck,omitempty"`
}

type AppfwprofileDenyurlBinding struct {
	Alertonly string `json:"alertonly,omitempty"`
	Comment string `json:"comment,omitempty"`
	Denyurl string `json:"denyurl,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	Name string `json:"name,omitempty"`
	Resourceid string `json:"resourceid,omitempty"`
	Ruletype string `json:"ruletype,omitempty"`
	State string `json:"state,omitempty"`
}

type Appfwsignatures struct {
	Action []string `json:"action,omitempty"`
	Autoenablenewsignatures string `json:"autoenablenewsignatures,omitempty"`
	Category string `json:"category,omitempty"`
	Comment string `json:"comment,omitempty"`
	Enabled string `json:"enabled,omitempty"`
	Encryptedversion int `json:"encryptedversion,omitempty"`
	Merge bool `json:"merge,omitempty"`
	Mergedefault bool `json:"mergedefault,omitempty"`
	Name string `json:"name,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Overwrite bool `json:"overwrite,omitempty"`
	Preservedefactions bool `json:"preservedefactions,omitempty"`
	Response string `json:"response,omitempty"`
	Ruleid []interface{} `json:"ruleid,omitempty"`
	Sha1 string `json:"sha1,omitempty"`
	Src string `json:"src,omitempty"`
	Vendortype string `json:"vendortype,omitempty"`
	Xslt string `json:"xslt,omitempty"`
}

type AppfwprofileFieldformatBinding struct {
	Alertonly string `json:"alertonly,omitempty"`
	Comment string `json:"comment,omitempty"`
	Fieldformat string `json:"fieldformat,omitempty"`
	Fieldformatmaxlength int `json:"fieldformatmaxlength,omitempty"`
	Fieldformatminlength int `json:"fieldformatminlength,omitempty"`
	Fieldtype string `json:"fieldtype,omitempty"`
	FormactionurlFf string `json:"formactionurl_ff,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	IsregexFf string `json:"isregex_ff,omitempty"`
	Name string `json:"name,omitempty"`
	Resourceid string `json:"resourceid,omitempty"`
	Ruletype string `json:"ruletype,omitempty"`
	State string `json:"state,omitempty"`
}

type AppfwprofileLogexpressionBinding struct {
	Alertonly string `json:"alertonly,omitempty"`
	AsLogexpression string `json:"as_logexpression,omitempty"`
	Comment string `json:"comment,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	Logexpression string `json:"logexpression,omitempty"`
	Name string `json:"name,omitempty"`
	Resourceid string `json:"resourceid,omitempty"`
	Ruletype string `json:"ruletype,omitempty"`
	State string `json:"state,omitempty"`
}

type Appfwjsonerrorpage struct {
	Comment string `json:"comment,omitempty"`
	Name string `json:"name,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Overwrite bool `json:"overwrite,omitempty"`
	Response string `json:"response,omitempty"`
	Src string `json:"src,omitempty"`
}

type AppfwprofileJsondosurlBinding struct {
	Alertonly string `json:"alertonly,omitempty"`
	Comment string `json:"comment,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	Jsondosurl string `json:"jsondosurl,omitempty"`
	Jsonmaxarraylength int `json:"jsonmaxarraylength,omitempty"`
	Jsonmaxarraylengthcheck string `json:"jsonmaxarraylengthcheck,omitempty"`
	Jsonmaxcontainerdepth int `json:"jsonmaxcontainerdepth,omitempty"`
	Jsonmaxcontainerdepthcheck string `json:"jsonmaxcontainerdepthcheck,omitempty"`
	Jsonmaxdocumentlength int `json:"jsonmaxdocumentlength,omitempty"`
	Jsonmaxdocumentlengthcheck string `json:"jsonmaxdocumentlengthcheck,omitempty"`
	Jsonmaxobjectkeycount int `json:"jsonmaxobjectkeycount,omitempty"`
	Jsonmaxobjectkeycountcheck string `json:"jsonmaxobjectkeycountcheck,omitempty"`
	Jsonmaxobjectkeylength int `json:"jsonmaxobjectkeylength,omitempty"`
	Jsonmaxobjectkeylengthcheck string `json:"jsonmaxobjectkeylengthcheck,omitempty"`
	Jsonmaxstringlength int `json:"jsonmaxstringlength,omitempty"`
	Jsonmaxstringlengthcheck string `json:"jsonmaxstringlengthcheck,omitempty"`
	Name string `json:"name,omitempty"`
	Resourceid string `json:"resourceid,omitempty"`
	Ruletype string `json:"ruletype,omitempty"`
	State string `json:"state,omitempty"`
}

type Appfwpolicy struct {
	Builtin []string `json:"builtin,omitempty"`
	Comment string `json:"comment,omitempty"`
	Count float64 `json:"__count,omitempty"`
	Feature string `json:"feature,omitempty"`
	Hits int `json:"hits,omitempty"`
	Logaction string `json:"logaction,omitempty"`
	Name string `json:"name,omitempty"`
	Newname string `json:"newname,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Policytype string `json:"policytype,omitempty"`
	Profilename string `json:"profilename,omitempty"`
	Rule string `json:"rule,omitempty"`
	Undefhits int `json:"undefhits,omitempty"`
}

type AppfwprofileFakeaccountBinding struct {
	Alertonly string `json:"alertonly,omitempty"`
	Comment string `json:"comment,omitempty"`
	Fakeaccount string `json:"fakeaccount,omitempty"`
	Formexpression string `json:"formexpression,omitempty"`
	FormurlFad string `json:"formurl_fad,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	Isfieldnameregex string `json:"isfieldnameregex,omitempty"`
	Name string `json:"name,omitempty"`
	Resourceid string `json:"resourceid,omitempty"`
	Ruletype string `json:"ruletype,omitempty"`
	State string `json:"state,omitempty"`
	Tag string `json:"tag,omitempty"`
}

type AppfwprofileJsoncmdurlBinding struct {
	Alertonly string `json:"alertonly,omitempty"`
	AsValueExprJsonCmd string `json:"as_value_expr_json_cmd,omitempty"`
	AsValueTypeJsonCmd string `json:"as_value_type_json_cmd,omitempty"`
	Comment string `json:"comment,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	IskeyregexJsonCmd string `json:"iskeyregex_json_cmd,omitempty"`
	IsvalueregexJsonCmd string `json:"isvalueregex_json_cmd,omitempty"`
	Jsoncmdurl string `json:"jsoncmdurl,omitempty"`
	KeynameJsonCmd string `json:"keyname_json_cmd,omitempty"`
	Name string `json:"name,omitempty"`
	Resourceid string `json:"resourceid,omitempty"`
	Ruletype string `json:"ruletype,omitempty"`
	State string `json:"state,omitempty"`
}

type Appfwconfidfield struct {
	Comment string `json:"comment,omitempty"`
	Count float64 `json:"__count,omitempty"`
	Fieldname string `json:"fieldname,omitempty"`
	Isregex string `json:"isregex,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	State string `json:"state,omitempty"`
	Url string `json:"url,omitempty"`
}

type AppfwglobalAuditnslogpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke bool `json:"invoke,omitempty"`
	Labelname string `json:"labelname,omitempty"`
	Labeltype string `json:"labeltype,omitempty"`
	Policyname string `json:"policyname,omitempty"`
	Policytype string `json:"policytype,omitempty"`
	Priority int `json:"priority,omitempty"`
	State string `json:"state,omitempty"`
	TypeField string `json:"type,omitempty"`
}

type AppfwprofileXmlvalidationurlBinding struct {
	Alertonly string `json:"alertonly,omitempty"`
	Comment string `json:"comment,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	Name string `json:"name,omitempty"`
	Resourceid string `json:"resourceid,omitempty"`
	Ruletype string `json:"ruletype,omitempty"`
	State string `json:"state,omitempty"`
	Xmladditionalsoapheaders string `json:"xmladditionalsoapheaders,omitempty"`
	Xmlendpointcheck string `json:"xmlendpointcheck,omitempty"`
	Xmlrequestschema string `json:"xmlrequestschema,omitempty"`
	Xmlresponseschema string `json:"xmlresponseschema,omitempty"`
	Xmlvalidateresponse string `json:"xmlvalidateresponse,omitempty"`
	Xmlvalidatesoapenvelope string `json:"xmlvalidatesoapenvelope,omitempty"`
	Xmlvalidationurl string `json:"xmlvalidationurl,omitempty"`
	Xmlwsdl string `json:"xmlwsdl,omitempty"`
}

type AppfwprofileAppfwconfidfieldBinding struct {
	Alertonly string `json:"alertonly,omitempty"`
	CffieldUrl string `json:"cffield_url,omitempty"`
	Comment string `json:"comment,omitempty"`
	Confidfield string `json:"confidfield,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	IsregexCffield string `json:"isregex_cffield,omitempty"`
	Name string `json:"name,omitempty"`
	Resourceid string `json:"resourceid,omitempty"`
	Ruletype string `json:"ruletype,omitempty"`
	State string `json:"state,omitempty"`
}

type AppfwprofileCrosssitescriptingBinding struct {
	Alertonly string `json:"alertonly,omitempty"`
	AsScanLocationXss string `json:"as_scan_location_xss,omitempty"`
	AsValueExprXss string `json:"as_value_expr_xss,omitempty"`
	AsValueTypeXss string `json:"as_value_type_xss,omitempty"`
	Comment string `json:"comment,omitempty"`
	Crosssitescripting string `json:"crosssitescripting,omitempty"`
	FormactionurlXss string `json:"formactionurl_xss,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	IsregexXss string `json:"isregex_xss,omitempty"`
	IsvalueregexXss string `json:"isvalueregex_xss,omitempty"`
	Name string `json:"name,omitempty"`
	Resourceid string `json:"resourceid,omitempty"`
	Ruletype string `json:"ruletype,omitempty"`
	State string `json:"state,omitempty"`
}

type Appfwtransactionrecords struct {
	Appfwsessionid string `json:"appfwsessionid,omitempty"`
	Clientip string `json:"clientip,omitempty"`
	Count float64 `json:"__count,omitempty"`
	Destip string `json:"destip,omitempty"`
	Endtime string `json:"endtime,omitempty"`
	Httptransactionid int `json:"httptransactionid,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Nodeid int `json:"nodeid,omitempty"`
	Packetengineid int `json:"packetengineid,omitempty"`
	Profilename string `json:"profilename,omitempty"`
	Requestcontentlength int `json:"requestcontentlength,omitempty"`
	Requestmaxprocessingtime int `json:"requestmaxprocessingtime,omitempty"`
	Requestyields int `json:"requestyields,omitempty"`
	Responsecontentlength int `json:"responsecontentlength,omitempty"`
	Responsemaxprocessingtime int `json:"responsemaxprocessingtime,omitempty"`
	Responseyields int `json:"responseyields,omitempty"`
	Starttime string `json:"starttime,omitempty"`
	Url string `json:"url,omitempty"`
}

type AppfwprofileBinding struct {
	AppfwprofileAppfwconfidfieldBinding []interface{} `json:"appfwprofile_appfwconfidfield_binding,omitempty"`
	AppfwprofileBlockkeywordBinding []interface{} `json:"appfwprofile_blockkeyword_binding,omitempty"`
	AppfwprofileBypasslistBinding []interface{} `json:"appfwprofile_bypasslist_binding,omitempty"`
	AppfwprofileCmdinjectionBinding []interface{} `json:"appfwprofile_cmdinjection_binding,omitempty"`
	AppfwprofileContenttypeBinding []interface{} `json:"appfwprofile_contenttype_binding,omitempty"`
	AppfwprofileCookieconsistencyBinding []interface{} `json:"appfwprofile_cookieconsistency_binding,omitempty"`
	AppfwprofileCreditcardnumberBinding []interface{} `json:"appfwprofile_creditcardnumber_binding,omitempty"`
	AppfwprofileCrosssitescriptingBinding []interface{} `json:"appfwprofile_crosssitescripting_binding,omitempty"`
	AppfwprofileCsrftagBinding []interface{} `json:"appfwprofile_csrftag_binding,omitempty"`
	AppfwprofileDenylistBinding []interface{} `json:"appfwprofile_denylist_binding,omitempty"`
	AppfwprofileDenyurlBinding []interface{} `json:"appfwprofile_denyurl_binding,omitempty"`
	AppfwprofileExcluderescontenttypeBinding []interface{} `json:"appfwprofile_excluderescontenttype_binding,omitempty"`
	AppfwprofileFakeaccountBinding []interface{} `json:"appfwprofile_fakeaccount_binding,omitempty"`
	AppfwprofileFieldconsistencyBinding []interface{} `json:"appfwprofile_fieldconsistency_binding,omitempty"`
	AppfwprofileFieldformatBinding []interface{} `json:"appfwprofile_fieldformat_binding,omitempty"`
	AppfwprofileFileuploadtypeBinding []interface{} `json:"appfwprofile_fileuploadtype_binding,omitempty"`
	AppfwprofileGrpcvalidationBinding []interface{} `json:"appfwprofile_grpcvalidation_binding,omitempty"`
	AppfwprofileJsonblockkeywordBinding []interface{} `json:"appfwprofile_jsonblockkeyword_binding,omitempty"`
	AppfwprofileJsoncmdurlBinding []interface{} `json:"appfwprofile_jsoncmdurl_binding,omitempty"`
	AppfwprofileJsondosurlBinding []interface{} `json:"appfwprofile_jsondosurl_binding,omitempty"`
	AppfwprofileJsonsqlurlBinding []interface{} `json:"appfwprofile_jsonsqlurl_binding,omitempty"`
	AppfwprofileJsonxssurlBinding []interface{} `json:"appfwprofile_jsonxssurl_binding,omitempty"`
	AppfwprofileLogexpressionBinding []interface{} `json:"appfwprofile_logexpression_binding,omitempty"`
	AppfwprofileRestvalidationBinding []interface{} `json:"appfwprofile_restvalidation_binding,omitempty"`
	AppfwprofileSafeobjectBinding []interface{} `json:"appfwprofile_safeobject_binding,omitempty"`
	AppfwprofileSqlinjectionBinding []interface{} `json:"appfwprofile_sqlinjection_binding,omitempty"`
	AppfwprofileStarturlBinding []interface{} `json:"appfwprofile_starturl_binding,omitempty"`
	AppfwprofileTrustedlearningclientsBinding []interface{} `json:"appfwprofile_trustedlearningclients_binding,omitempty"`
	AppfwprofileXmlattachmenturlBinding []interface{} `json:"appfwprofile_xmlattachmenturl_binding,omitempty"`
	AppfwprofileXmldosurlBinding []interface{} `json:"appfwprofile_xmldosurl_binding,omitempty"`
	AppfwprofileXmlsqlinjectionBinding []interface{} `json:"appfwprofile_xmlsqlinjection_binding,omitempty"`
	AppfwprofileXmlvalidationurlBinding []interface{} `json:"appfwprofile_xmlvalidationurl_binding,omitempty"`
	AppfwprofileXmlwsiurlBinding []interface{} `json:"appfwprofile_xmlwsiurl_binding,omitempty"`
	AppfwprofileXmlxssBinding []interface{} `json:"appfwprofile_xmlxss_binding,omitempty"`
	Name string `json:"name,omitempty"`
}

type AppfwprofileCsrftagBinding struct {
	Alertonly string `json:"alertonly,omitempty"`
	Comment string `json:"comment,omitempty"`
	Csrfformactionurl string `json:"csrfformactionurl,omitempty"`
	Csrftag string `json:"csrftag,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	Name string `json:"name,omitempty"`
	Resourceid string `json:"resourceid,omitempty"`
	Ruletype string `json:"ruletype,omitempty"`
	State string `json:"state,omitempty"`
}

type Appfwwsdl struct {
	Comment string `json:"comment,omitempty"`
	Name string `json:"name,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Overwrite bool `json:"overwrite,omitempty"`
	Response string `json:"response,omitempty"`
	Src string `json:"src,omitempty"`
}

type Appfwpolicylabel struct {
	Count float64 `json:"__count,omitempty"`
	Description string `json:"description,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Hits int `json:"hits,omitempty"`
	InvokeLabelname string `json:"invoke_labelname,omitempty"`
	Labelname string `json:"labelname,omitempty"`
	Labeltype string `json:"labeltype,omitempty"`
	Newname string `json:"newname,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Numpol int `json:"numpol,omitempty"`
	Policylabeltype string `json:"policylabeltype,omitempty"`
	Policytype string `json:"policytype,omitempty"`
	Priority int `json:"priority,omitempty"`
}

type Appfwprotofile struct {
	Comment string `json:"comment,omitempty"`
	Name string `json:"name,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Overwrite bool `json:"overwrite,omitempty"`
	Response string `json:"response,omitempty"`
	Src string `json:"src,omitempty"`
}

type AppfwglobalAuditsyslogpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke bool `json:"invoke,omitempty"`
	Labelname string `json:"labelname,omitempty"`
	Labeltype string `json:"labeltype,omitempty"`
	Policyname string `json:"policyname,omitempty"`
	Policytype string `json:"policytype,omitempty"`
	Priority int `json:"priority,omitempty"`
	State string `json:"state,omitempty"`
	TypeField string `json:"type,omitempty"`
}

type AppfwglobalBinding struct {
	AppfwglobalAppfwpolicyBinding []interface{} `json:"appfwglobal_appfwpolicy_binding,omitempty"`
	AppfwglobalAuditnslogpolicyBinding []interface{} `json:"appfwglobal_auditnslogpolicy_binding,omitempty"`
	AppfwglobalAuditsyslogpolicyBinding []interface{} `json:"appfwglobal_auditsyslogpolicy_binding,omitempty"`
}

type AppfwprofileJsonxssurlBinding struct {
	Alertonly string `json:"alertonly,omitempty"`
	AsValueExprJsonXss string `json:"as_value_expr_json_xss,omitempty"`
	AsValueTypeJsonXss string `json:"as_value_type_json_xss,omitempty"`
	Comment string `json:"comment,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	IskeyregexJsonXss string `json:"iskeyregex_json_xss,omitempty"`
	IsvalueregexJsonXss string `json:"isvalueregex_json_xss,omitempty"`
	Jsonxssurl string `json:"jsonxssurl,omitempty"`
	KeynameJsonXss string `json:"keyname_json_xss,omitempty"`
	Name string `json:"name,omitempty"`
	Resourceid string `json:"resourceid,omitempty"`
	Ruletype string `json:"ruletype,omitempty"`
	State string `json:"state,omitempty"`
}

type AppfwprofileFileuploadtypeBinding struct {
	Alertonly string `json:"alertonly,omitempty"`
	AsFileuploadtypesUrl string `json:"as_fileuploadtypes_url,omitempty"`
	Comment string `json:"comment,omitempty"`
	Filetype []string `json:"filetype,omitempty"`
	Fileuploadtype string `json:"fileuploadtype,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	Isnameregex string `json:"isnameregex,omitempty"`
	IsregexFileuploadtypesUrl string `json:"isregex_fileuploadtypes_url,omitempty"`
	Name string `json:"name,omitempty"`
	Resourceid string `json:"resourceid,omitempty"`
	Ruletype string `json:"ruletype,omitempty"`
	State string `json:"state,omitempty"`
}

type AppfwprofileXmlattachmenturlBinding struct {
	Alertonly string `json:"alertonly,omitempty"`
	Comment string `json:"comment,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	Name string `json:"name,omitempty"`
	Resourceid string `json:"resourceid,omitempty"`
	Ruletype string `json:"ruletype,omitempty"`
	State string `json:"state,omitempty"`
	Xmlattachmentcontenttype string `json:"xmlattachmentcontenttype,omitempty"`
	Xmlattachmentcontenttypecheck string `json:"xmlattachmentcontenttypecheck,omitempty"`
	Xmlattachmenturl string `json:"xmlattachmenturl,omitempty"`
	Xmlmaxattachmentsize int `json:"xmlmaxattachmentsize,omitempty"`
	Xmlmaxattachmentsizecheck string `json:"xmlmaxattachmentsizecheck,omitempty"`
}

type Appfwurlencodedformcontenttype struct {
	Builtin []string `json:"builtin,omitempty"`
	Count float64 `json:"__count,omitempty"`
	Feature string `json:"feature,omitempty"`
	Isregex string `json:"isregex,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Urlencodedformcontenttypevalue string `json:"urlencodedformcontenttypevalue,omitempty"`
}

type AppfwpolicyBinding struct {
	AppfwpolicyAppfwglobalBinding []interface{} `json:"appfwpolicy_appfwglobal_binding,omitempty"`
	AppfwpolicyAppfwpolicylabelBinding []interface{} `json:"appfwpolicy_appfwpolicylabel_binding,omitempty"`
	AppfwpolicyCsvserverBinding []interface{} `json:"appfwpolicy_csvserver_binding,omitempty"`
	AppfwpolicyLbvserverBinding []interface{} `json:"appfwpolicy_lbvserver_binding,omitempty"`
	AppfwpolicyVpnvserverBinding []interface{} `json:"appfwpolicy_vpnvserver_binding,omitempty"`
	Name string `json:"name,omitempty"`
}

type AppfwprofileFieldconsistencyBinding struct {
	Alertonly string `json:"alertonly,omitempty"`
	Comment string `json:"comment,omitempty"`
	Fieldconsistency string `json:"fieldconsistency,omitempty"`
	FormactionurlFfc string `json:"formactionurl_ffc,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	IsregexFfc string `json:"isregex_ffc,omitempty"`
	Name string `json:"name,omitempty"`
	Resourceid string `json:"resourceid,omitempty"`
	Ruletype string `json:"ruletype,omitempty"`
	State string `json:"state,omitempty"`
}

type Appfwxmlschema struct {
	Comment string `json:"comment,omitempty"`
	Name string `json:"name,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Overwrite bool `json:"overwrite,omitempty"`
	Response string `json:"response,omitempty"`
	Src string `json:"src,omitempty"`
}

type Appfwarchive struct {
	Comment string `json:"comment,omitempty"`
	Name string `json:"name,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Response string `json:"response,omitempty"`
	Src string `json:"src,omitempty"`
	Target string `json:"target,omitempty"`
}

type AppfwpolicyAppfwpolicylabelBinding struct {
	Activepolicy int `json:"activepolicy,omitempty"`
	Boundto string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname string `json:"labelname,omitempty"`
	Labeltype string `json:"labeltype,omitempty"`
	Name string `json:"name,omitempty"`
	Priority int `json:"priority,omitempty"`
}

type AppfwprofileBlockkeywordBinding struct {
	Alertonly string `json:"alertonly,omitempty"`
	AsBlockkeywordFormurl string `json:"as_blockkeyword_formurl,omitempty"`
	AsFieldnameIsregexBlockkeyword string `json:"as_fieldname_isregex_blockkeyword,omitempty"`
	Blockkeyword string `json:"blockkeyword,omitempty"`
	Blockkeywordtype string `json:"blockkeywordtype,omitempty"`
	Comment string `json:"comment,omitempty"`
	Fieldname string `json:"fieldname,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	Name string `json:"name,omitempty"`
	Resourceid string `json:"resourceid,omitempty"`
	Ruletype string `json:"ruletype,omitempty"`
	State string `json:"state,omitempty"`
}

type Appfwsettings struct {
	Builtin []string `json:"builtin,omitempty"`
	Ceflogging string `json:"ceflogging,omitempty"`
	Centralizedlearning string `json:"centralizedlearning,omitempty"`
	Clientiploggingheader string `json:"clientiploggingheader,omitempty"`
	Cookieflags string `json:"cookieflags,omitempty"`
	Cookiepostencryptprefix string `json:"cookiepostencryptprefix,omitempty"`
	Defaultprofile string `json:"defaultprofile,omitempty"`
	Entitydecoding string `json:"entitydecoding,omitempty"`
	Feature string `json:"feature,omitempty"`
	Geolocationlogging string `json:"geolocationlogging,omitempty"`
	Importsizelimit int `json:"importsizelimit,omitempty"`
	Learning string `json:"learning,omitempty"`
	Learnratelimit int `json:"learnratelimit,omitempty"`
	Logmalformedreq string `json:"logmalformedreq,omitempty"`
	Malformedreqaction []string `json:"malformedreqaction,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Proxypassword string `json:"proxypassword,omitempty"`
	Proxyport int `json:"proxyport,omitempty"`
	Proxyserver string `json:"proxyserver,omitempty"`
	Proxyusername string `json:"proxyusername,omitempty"`
	Sessioncookiename string `json:"sessioncookiename,omitempty"`
	Sessionlifetime int `json:"sessionlifetime,omitempty"`
	Sessionlimit int `json:"sessionlimit,omitempty"`
	Sessiontimeout int `json:"sessiontimeout,omitempty"`
	Signatureautoupdate string `json:"signatureautoupdate,omitempty"`
	Signatureurl string `json:"signatureurl,omitempty"`
	Undefaction string `json:"undefaction,omitempty"`
	Useconfigurablesecretkey string `json:"useconfigurablesecretkey,omitempty"`
}

type AppfwpolicyCsvserverBinding struct {
	Activepolicy int `json:"activepolicy,omitempty"`
	Boundto string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname string `json:"labelname,omitempty"`
	Labeltype string `json:"labeltype,omitempty"`
	Name string `json:"name,omitempty"`
	Priority int `json:"priority,omitempty"`
}

type AppfwprofileCmdinjectionBinding struct {
	Alertonly string `json:"alertonly,omitempty"`
	AsScanLocationCmd string `json:"as_scan_location_cmd,omitempty"`
	AsValueExprCmd string `json:"as_value_expr_cmd,omitempty"`
	AsValueTypeCmd string `json:"as_value_type_cmd,omitempty"`
	Cmdinjection string `json:"cmdinjection,omitempty"`
	Comment string `json:"comment,omitempty"`
	FormactionurlCmd string `json:"formactionurl_cmd,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	IsregexCmd string `json:"isregex_cmd,omitempty"`
	IsvalueregexCmd string `json:"isvalueregex_cmd,omitempty"`
	Name string `json:"name,omitempty"`
	Resourceid string `json:"resourceid,omitempty"`
	Ruletype string `json:"ruletype,omitempty"`
	State string `json:"state,omitempty"`
}

type AppfwpolicylabelPolicybindingBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke bool `json:"invoke,omitempty"`
	InvokeLabelname string `json:"invoke_labelname,omitempty"`
	Labelname string `json:"labelname,omitempty"`
	Labeltype string `json:"labeltype,omitempty"`
	Policyname string `json:"policyname,omitempty"`
	Priority int `json:"priority,omitempty"`
}

type AppfwprofileStarturlBinding struct {
	Alertonly string `json:"alertonly,omitempty"`
	Comment string `json:"comment,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	Name string `json:"name,omitempty"`
	Resourceid string `json:"resourceid,omitempty"`
	Ruletype string `json:"ruletype,omitempty"`
	Starturl string `json:"starturl,omitempty"`
	State string `json:"state,omitempty"`
}

type AppfwprofileJsonblockkeywordBinding struct {
	Alertonly string `json:"alertonly,omitempty"`
	Comment string `json:"comment,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	IskeyregexJsonBlockkeyword string `json:"iskeyregex_json_blockkeyword,omitempty"`
	Jsonblockkeyword string `json:"jsonblockkeyword,omitempty"`
	Jsonblockkeywordtype string `json:"jsonblockkeywordtype,omitempty"`
	Jsonblockkeywordurl string `json:"jsonblockkeywordurl,omitempty"`
	KeynameJsonBlockkeyword string `json:"keyname_json_blockkeyword,omitempty"`
	Name string `json:"name,omitempty"`
	Resourceid string `json:"resourceid,omitempty"`
	Ruletype string `json:"ruletype,omitempty"`
	State string `json:"state,omitempty"`
}

type Appfwfieldtype struct {
	Builtin []string `json:"builtin,omitempty"`
	Comment string `json:"comment,omitempty"`
	Count float64 `json:"__count,omitempty"`
	Feature string `json:"feature,omitempty"`
	Name string `json:"name,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Nocharmaps bool `json:"nocharmaps,omitempty"`
	Priority int `json:"priority,omitempty"`
	Regex string `json:"regex,omitempty"`
}

type Appfwlearningdata struct {
	AsScanLocationSql string `json:"as_scan_location_sql,omitempty"`
	AsScanLocationXss string `json:"as_scan_location_xss,omitempty"`
	AsValueExprSql string `json:"as_value_expr_sql,omitempty"`
	AsValueExprXss string `json:"as_value_expr_xss,omitempty"`
	AsValueTypeSql string `json:"as_value_type_sql,omitempty"`
	AsValueTypeXss string `json:"as_value_type_xss,omitempty"`
	Contenttype string `json:"contenttype,omitempty"`
	Cookieconsistency string `json:"cookieconsistency,omitempty"`
	Count float64 `json:"__count,omitempty"`
	Creditcardnumber string `json:"creditcardnumber,omitempty"`
	Creditcardnumberurl string `json:"creditcardnumberurl,omitempty"`
	Crosssitescripting string `json:"crosssitescripting,omitempty"`
	Csrfformoriginurl string `json:"csrfformoriginurl,omitempty"`
	Csrftag string `json:"csrftag,omitempty"`
	Data string `json:"data,omitempty"`
	Fieldconsistency string `json:"fieldconsistency,omitempty"`
	Fieldformat string `json:"fieldformat,omitempty"`
	Fieldformatcharmappcre string `json:"fieldformatcharmappcre,omitempty"`
	Fieldformatmaxlength int `json:"fieldformatmaxlength,omitempty"`
	Fieldformatminlength int `json:"fieldformatminlength,omitempty"`
	Fieldtype string `json:"fieldtype,omitempty"`
	FormactionurlFf string `json:"formactionurl_ff,omitempty"`
	FormactionurlFfc string `json:"formactionurl_ffc,omitempty"`
	FormactionurlSql string `json:"formactionurl_sql,omitempty"`
	FormactionurlXss string `json:"formactionurl_xss,omitempty"`
	Hits int `json:"hits,omitempty"`
	Name string `json:"name,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Profilename string `json:"profilename,omitempty"`
	Securitycheck string `json:"securitycheck,omitempty"`
	Sqlinjection string `json:"sqlinjection,omitempty"`
	Starturl string `json:"starturl,omitempty"`
	Target string `json:"target,omitempty"`
	Totalxmlrequests bool `json:"totalxmlrequests,omitempty"`
	Url string `json:"url,omitempty"`
	Value string `json:"value,omitempty"`
	ValueType string `json:"value_type,omitempty"`
	Xmlattachmentcheck string `json:"xmlattachmentcheck,omitempty"`
	Xmldoscheck string `json:"xmldoscheck,omitempty"`
	Xmlwsicheck string `json:"xmlwsicheck,omitempty"`
}

type AppfwpolicylabelAppfwpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke bool `json:"invoke,omitempty"`
	InvokeLabelname string `json:"invoke_labelname,omitempty"`
	Labelname string `json:"labelname,omitempty"`
	Labeltype string `json:"labeltype,omitempty"`
	Policyname string `json:"policyname,omitempty"`
	Priority int `json:"priority,omitempty"`
}

type AppfwpolicyVpnvserverBinding struct {
	Activepolicy int `json:"activepolicy,omitempty"`
	Boundto string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname string `json:"labelname,omitempty"`
	Labeltype string `json:"labeltype,omitempty"`
	Name string `json:"name,omitempty"`
	Priority int `json:"priority,omitempty"`
}

type AppfwprofileXmlsqlinjectionBinding struct {
	Alertonly string `json:"alertonly,omitempty"`
	AsScanLocationXmlsql string `json:"as_scan_location_xmlsql,omitempty"`
	Comment string `json:"comment,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	IsregexXmlsql string `json:"isregex_xmlsql,omitempty"`
	Name string `json:"name,omitempty"`
	Resourceid string `json:"resourceid,omitempty"`
	Ruletype string `json:"ruletype,omitempty"`
	State string `json:"state,omitempty"`
	Xmlsqlinjection string `json:"xmlsqlinjection,omitempty"`
}

type Appfwgrpcwebjsoncontenttype struct {
	Count float64 `json:"__count,omitempty"`
	Grpcwebjsoncontenttypevalue string `json:"grpcwebjsoncontenttypevalue,omitempty"`
	Isregex string `json:"isregex,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type AppfwprofileTrustedlearningclientsBinding struct {
	Alertonly string `json:"alertonly,omitempty"`
	Comment string `json:"comment,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	Name string `json:"name,omitempty"`
	Resourceid string `json:"resourceid,omitempty"`
	Ruletype string `json:"ruletype,omitempty"`
	State string `json:"state,omitempty"`
	Trustedlearningclients string `json:"trustedlearningclients,omitempty"`
}

type AppfwprofileCreditcardnumberBinding struct {
	Alertonly string `json:"alertonly,omitempty"`
	Comment string `json:"comment,omitempty"`
	Creditcardnumber string `json:"creditcardnumber,omitempty"`
	Creditcardnumberurl string `json:"creditcardnumberurl,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	Name string `json:"name,omitempty"`
	Resourceid string `json:"resourceid,omitempty"`
	Ruletype string `json:"ruletype,omitempty"`
	State string `json:"state,omitempty"`
}

type AppfwprofileExcluderescontenttypeBinding struct {
	Alertonly string `json:"alertonly,omitempty"`
	Comment string `json:"comment,omitempty"`
	Excluderescontenttype string `json:"excluderescontenttype,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	Name string `json:"name,omitempty"`
	Resourceid string `json:"resourceid,omitempty"`
	Ruletype string `json:"ruletype,omitempty"`
	State string `json:"state,omitempty"`
}

type Appfwjsoncontenttype struct {
	Builtin []string `json:"builtin,omitempty"`
	Count float64 `json:"__count,omitempty"`
	Feature string `json:"feature,omitempty"`
	Isregex string `json:"isregex,omitempty"`
	Jsoncontenttypevalue string `json:"jsoncontenttypevalue,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type AppfwprofileSqlinjectionBinding struct {
	Alertonly string `json:"alertonly,omitempty"`
	AsScanLocationSql string `json:"as_scan_location_sql,omitempty"`
	AsValueExprSql string `json:"as_value_expr_sql,omitempty"`
	AsValueTypeSql string `json:"as_value_type_sql,omitempty"`
	Comment string `json:"comment,omitempty"`
	FormactionurlSql string `json:"formactionurl_sql,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	IsregexSql string `json:"isregex_sql,omitempty"`
	IsvalueregexSql string `json:"isvalueregex_sql,omitempty"`
	Name string `json:"name,omitempty"`
	Resourceid string `json:"resourceid,omitempty"`
	Ruletype string `json:"ruletype,omitempty"`
	Sqlinjection string `json:"sqlinjection,omitempty"`
	State string `json:"state,omitempty"`
}

type AppfwprofileRestvalidationBinding struct {
	Alertonly string `json:"alertonly,omitempty"`
	Comment string `json:"comment,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	Name string `json:"name,omitempty"`
	Resourceid string `json:"resourceid,omitempty"`
	RestValidationAction string `json:"rest_validation_action,omitempty"`
	Restvalidation string `json:"restvalidation,omitempty"`
	Ruletype string `json:"ruletype,omitempty"`
	State string `json:"state,omitempty"`
}

type Appfwgrpccontenttype struct {
	Builtin []string `json:"builtin,omitempty"`
	Count float64 `json:"__count,omitempty"`
	Feature string `json:"feature,omitempty"`
	Grpccontenttypevalue string `json:"grpccontenttypevalue,omitempty"`
	Isregex string `json:"isregex,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type AppfwprofileBypasslistBinding struct {
	Alertonly string `json:"alertonly,omitempty"`
	AsBypassList string `json:"as_bypass_list,omitempty"`
	AsBypassListAction string `json:"as_bypass_list_action,omitempty"`
	AsBypassListLocation string `json:"as_bypass_list_location,omitempty"`
	AsBypassListValueType string `json:"as_bypass_list_value_type,omitempty"`
	Comment string `json:"comment,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	Name string `json:"name,omitempty"`
	Resourceid string `json:"resourceid,omitempty"`
	Ruletype string `json:"ruletype,omitempty"`
	State string `json:"state,omitempty"`
}

type Appfwhtmlerrorpage struct {
	Comment string `json:"comment,omitempty"`
	Name string `json:"name,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Overwrite bool `json:"overwrite,omitempty"`
	Response string `json:"response,omitempty"`
	Src string `json:"src,omitempty"`
}

type AppfwprofileSafeobjectBinding struct {
	Action []string `json:"action,omitempty"`
	Alertonly string `json:"alertonly,omitempty"`
	AsExpression string `json:"as_expression,omitempty"`
	Comment string `json:"comment,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	Maxmatchlength int `json:"maxmatchlength,omitempty"`
	Name string `json:"name,omitempty"`
	Resourceid string `json:"resourceid,omitempty"`
	Ruletype string `json:"ruletype,omitempty"`
	Safeobject string `json:"safeobject,omitempty"`
	State string `json:"state,omitempty"`
}

type AppfwprofileGrpcvalidationBinding struct {
	Alertonly string `json:"alertonly,omitempty"`
	Comment string `json:"comment,omitempty"`
	GrpcRelaxValidationAction string `json:"grpc_relax_validation_action,omitempty"`
	Grpcvalidation string `json:"grpcvalidation,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	Name string `json:"name,omitempty"`
	Resourceid string `json:"resourceid,omitempty"`
	Ruletype string `json:"ruletype,omitempty"`
	State string `json:"state,omitempty"`
}

type AppfwprofileXmlwsiurlBinding struct {
	Alertonly string `json:"alertonly,omitempty"`
	Comment string `json:"comment,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	Name string `json:"name,omitempty"`
	Resourceid string `json:"resourceid,omitempty"`
	Ruletype string `json:"ruletype,omitempty"`
	State string `json:"state,omitempty"`
	Xmlwsichecks string `json:"xmlwsichecks,omitempty"`
	Xmlwsiurl string `json:"xmlwsiurl,omitempty"`
}

type AppfwpolicyAppfwglobalBinding struct {
	Activepolicy int `json:"activepolicy,omitempty"`
	Boundto string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname string `json:"labelname,omitempty"`
	Labeltype string `json:"labeltype,omitempty"`
	Name string `json:"name,omitempty"`
	Priority int `json:"priority,omitempty"`
}

type AppfwprofileDenylistBinding struct {
	Alertonly string `json:"alertonly,omitempty"`
	AsDenyList string `json:"as_deny_list,omitempty"`
	AsDenyListAction []string `json:"as_deny_list_action,omitempty"`
	AsDenyListLocation string `json:"as_deny_list_location,omitempty"`
	AsDenyListValueType string `json:"as_deny_list_value_type,omitempty"`
	Comment string `json:"comment,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	Name string `json:"name,omitempty"`
	Resourceid string `json:"resourceid,omitempty"`
	Ruletype string `json:"ruletype,omitempty"`
	State string `json:"state,omitempty"`
}

type AppfwpolicylabelBinding struct {
	AppfwpolicylabelAppfwpolicyBinding []interface{} `json:"appfwpolicylabel_appfwpolicy_binding,omitempty"`
	AppfwpolicylabelPolicybindingBinding []interface{} `json:"appfwpolicylabel_policybinding_binding,omitempty"`
	Labelname string `json:"labelname,omitempty"`
}

type Appfwprofile struct {
	Addcookieflags string `json:"addcookieflags,omitempty"`
	Apispec string `json:"apispec,omitempty"`
	Archivename string `json:"archivename,omitempty"`
	AsProfBypassListEnable string `json:"as_prof_bypass_list_enable,omitempty"`
	AsProfDenyListEnable string `json:"as_prof_deny_list_enable,omitempty"`
	Augment bool `json:"augment,omitempty"`
	Blockkeywordaction []string `json:"blockkeywordaction,omitempty"`
	Bufferoverflowaction []string `json:"bufferoverflowaction,omitempty"`
	Bufferoverflowmaxcookielength int `json:"bufferoverflowmaxcookielength,omitempty"`
	Bufferoverflowmaxheaderlength int `json:"bufferoverflowmaxheaderlength,omitempty"`
	Bufferoverflowmaxquerylength int `json:"bufferoverflowmaxquerylength,omitempty"`
	Bufferoverflowmaxtotalheaderlength int `json:"bufferoverflowmaxtotalheaderlength,omitempty"`
	Bufferoverflowmaxurllength int `json:"bufferoverflowmaxurllength,omitempty"`
	Builtin bool `json:"builtin,omitempty"`
	Canonicalizehtmlresponse string `json:"canonicalizehtmlresponse,omitempty"`
	Ceflogging string `json:"ceflogging,omitempty"`
	Checkrequestheaders string `json:"checkrequestheaders,omitempty"`
	Clientipexpression string `json:"clientipexpression,omitempty"`
	Cmdinjectionaction []string `json:"cmdinjectionaction,omitempty"`
	Cmdinjectiongrammar string `json:"cmdinjectiongrammar,omitempty"`
	Cmdinjectiontype string `json:"cmdinjectiontype,omitempty"`
	Comment string `json:"comment,omitempty"`
	Contenttypeaction []string `json:"contenttypeaction,omitempty"`
	Cookieconsistencyaction []string `json:"cookieconsistencyaction,omitempty"`
	Cookieencryption string `json:"cookieencryption,omitempty"`
	Cookiehijackingaction []string `json:"cookiehijackingaction,omitempty"`
	Cookieproxying string `json:"cookieproxying,omitempty"`
	Cookiesamesiteattribute string `json:"cookiesamesiteattribute,omitempty"`
	Cookietransforms string `json:"cookietransforms,omitempty"`
	Count float64 `json:"__count,omitempty"`
	Creditcard []string `json:"creditcard,omitempty"`
	Creditcardaction []string `json:"creditcardaction,omitempty"`
	Creditcardmaxallowed int `json:"creditcardmaxallowed,omitempty"`
	Creditcardxout string `json:"creditcardxout,omitempty"`
	Crosssitescriptingaction []string `json:"crosssitescriptingaction,omitempty"`
	Crosssitescriptingcheckcompleteurls string `json:"crosssitescriptingcheckcompleteurls,omitempty"`
	Crosssitescriptingtransformunsafehtml string `json:"crosssitescriptingtransformunsafehtml,omitempty"`
	Csrftag string `json:"csrftag,omitempty"`
	Csrftagaction []string `json:"csrftagaction,omitempty"`
	Customsettings string `json:"customsettings,omitempty"`
	Defaultcharset string `json:"defaultcharset,omitempty"`
	Defaultfieldformatmaxlength int `json:"defaultfieldformatmaxlength,omitempty"`
	Defaultfieldformatmaxoccurrences int `json:"defaultfieldformatmaxoccurrences,omitempty"`
	Defaultfieldformatminlength int `json:"defaultfieldformatminlength,omitempty"`
	Defaultfieldformattype string `json:"defaultfieldformattype,omitempty"`
	Defaults string `json:"defaults,omitempty"`
	Denyurlaction []string `json:"denyurlaction,omitempty"`
	Dosecurecreditcardlogging string `json:"dosecurecreditcardlogging,omitempty"`
	Dynamiclearning []string `json:"dynamiclearning,omitempty"`
	Enableformtagging string `json:"enableformtagging,omitempty"`
	Errorurl string `json:"errorurl,omitempty"`
	Excludefileuploadfromchecks string `json:"excludefileuploadfromchecks,omitempty"`
	Exemptclosureurlsfromsecuritychecks string `json:"exemptclosureurlsfromsecuritychecks,omitempty"`
	Fakeaccountdetection string `json:"fakeaccountdetection,omitempty"`
	Fieldconsistencyaction []string `json:"fieldconsistencyaction,omitempty"`
	Fieldformataction []string `json:"fieldformataction,omitempty"`
	Fieldscan string `json:"fieldscan,omitempty"`
	Fieldscanlimit int `json:"fieldscanlimit,omitempty"`
	Fileuploadmaxnum int `json:"fileuploadmaxnum,omitempty"`
	Fileuploadtypesaction []string `json:"fileuploadtypesaction,omitempty"`
	Geolocationlogging string `json:"geolocationlogging,omitempty"`
	Grpcaction []string `json:"grpcaction,omitempty"`
	Htmlerrorobject string `json:"htmlerrorobject,omitempty"`
	Htmlerrorstatuscode int `json:"htmlerrorstatuscode,omitempty"`
	Htmlerrorstatusmessage string `json:"htmlerrorstatusmessage,omitempty"`
	Importprofilename string `json:"importprofilename,omitempty"`
	Infercontenttypexmlpayloadaction []string `json:"infercontenttypexmlpayloadaction,omitempty"`
	Insertcookiesamesiteattribute string `json:"insertcookiesamesiteattribute,omitempty"`
	Inspectcontenttypes []string `json:"inspectcontenttypes,omitempty"`
	Inspectquerycontenttypes []string `json:"inspectquerycontenttypes,omitempty"`
	Invalidpercenthandling string `json:"invalidpercenthandling,omitempty"`
	Jsonblockkeywordaction []string `json:"jsonblockkeywordaction,omitempty"`
	Jsoncmdinjectionaction []string `json:"jsoncmdinjectionaction,omitempty"`
	Jsoncmdinjectiongrammar string `json:"jsoncmdinjectiongrammar,omitempty"`
	Jsoncmdinjectiontype string `json:"jsoncmdinjectiontype,omitempty"`
	Jsondosaction []string `json:"jsondosaction,omitempty"`
	Jsonerrorobject string `json:"jsonerrorobject,omitempty"`
	Jsonerrorstatuscode int `json:"jsonerrorstatuscode,omitempty"`
	Jsonerrorstatusmessage string `json:"jsonerrorstatusmessage,omitempty"`
	Jsonfieldscan string `json:"jsonfieldscan,omitempty"`
	Jsonfieldscanlimit int `json:"jsonfieldscanlimit,omitempty"`
	Jsonmessagescan string `json:"jsonmessagescan,omitempty"`
	Jsonmessagescanlimit int `json:"jsonmessagescanlimit,omitempty"`
	Jsonsqlinjectionaction []string `json:"jsonsqlinjectionaction,omitempty"`
	Jsonsqlinjectiongrammar string `json:"jsonsqlinjectiongrammar,omitempty"`
	Jsonsqlinjectiontype string `json:"jsonsqlinjectiontype,omitempty"`
	Jsonxssaction []string `json:"jsonxssaction,omitempty"`
	Learning string `json:"learning,omitempty"`
	Logeverypolicyhit string `json:"logeverypolicyhit,omitempty"`
	Matchurlstring string `json:"matchurlstring,omitempty"`
	Messagescan string `json:"messagescan,omitempty"`
	Messagescanlimit int `json:"messagescanlimit,omitempty"`
	Messagescanlimitcontenttypes []string `json:"messagescanlimitcontenttypes,omitempty"`
	Multipleheaderaction []string `json:"multipleheaderaction,omitempty"`
	Name string `json:"name,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Optimizepartialreqs string `json:"optimizepartialreqs,omitempty"`
	Overwrite bool `json:"overwrite,omitempty"`
	Percentdecoderecursively string `json:"percentdecoderecursively,omitempty"`
	Postbodylimit int `json:"postbodylimit,omitempty"`
	Postbodylimitaction []string `json:"postbodylimitaction,omitempty"`
	Postbodylimitsignature int `json:"postbodylimitsignature,omitempty"`
	Protofileobject string `json:"protofileobject,omitempty"`
	Refererheadercheck string `json:"refererheadercheck,omitempty"`
	Relaxationrules bool `json:"relaxationrules,omitempty"`
	Replaceurlstring string `json:"replaceurlstring,omitempty"`
	Requestcontenttype string `json:"requestcontenttype,omitempty"`
	Responsecontenttype string `json:"responsecontenttype,omitempty"`
	Restaction []string `json:"restaction,omitempty"`
	Rfcprofile string `json:"rfcprofile,omitempty"`
	Semicolonfieldseparator string `json:"semicolonfieldseparator,omitempty"`
	Sessioncookiename string `json:"sessioncookiename,omitempty"`
	Sessionlessfieldconsistency string `json:"sessionlessfieldconsistency,omitempty"`
	Sessionlessurlclosure string `json:"sessionlessurlclosure,omitempty"`
	Signatures string `json:"signatures,omitempty"`
	Sqlinjectionaction []string `json:"sqlinjectionaction,omitempty"`
	Sqlinjectionchecksqlwildchars string `json:"sqlinjectionchecksqlwildchars,omitempty"`
	Sqlinjectiongrammar string `json:"sqlinjectiongrammar,omitempty"`
	Sqlinjectiononlycheckfieldswithsqlchars string `json:"sqlinjectiononlycheckfieldswithsqlchars,omitempty"`
	Sqlinjectionparsecomments string `json:"sqlinjectionparsecomments,omitempty"`
	Sqlinjectionruletype string `json:"sqlinjectionruletype,omitempty"`
	Sqlinjectiontransformspecialchars string `json:"sqlinjectiontransformspecialchars,omitempty"`
	Sqlinjectiontype string `json:"sqlinjectiontype,omitempty"`
	Starturlaction []string `json:"starturlaction,omitempty"`
	Starturlclosure string `json:"starturlclosure,omitempty"`
	State string `json:"state,omitempty"`
	Streaming string `json:"streaming,omitempty"`
	Stripcomments string `json:"stripcomments,omitempty"`
	Striphtmlcomments string `json:"striphtmlcomments,omitempty"`
	Stripxmlcomments string `json:"stripxmlcomments,omitempty"`
	Trace string `json:"trace,omitempty"`
	TypeField []string `json:"type,omitempty"`
	Urldecoderequestcookies string `json:"urldecoderequestcookies,omitempty"`
	Usehtmlerrorobject string `json:"usehtmlerrorobject,omitempty"`
	Verboseloglevel string `json:"verboseloglevel,omitempty"`
	Xmlattachmentaction []string `json:"xmlattachmentaction,omitempty"`
	Xmldosaction []string `json:"xmldosaction,omitempty"`
	Xmlerrorobject string `json:"xmlerrorobject,omitempty"`
	Xmlerrorstatuscode int `json:"xmlerrorstatuscode,omitempty"`
	Xmlerrorstatusmessage string `json:"xmlerrorstatusmessage,omitempty"`
	Xmlformataction []string `json:"xmlformataction,omitempty"`
	Xmlsoapfaultaction []string `json:"xmlsoapfaultaction,omitempty"`
	Xmlsqlinjectionaction []string `json:"xmlsqlinjectionaction,omitempty"`
	Xmlsqlinjectionchecksqlwildchars string `json:"xmlsqlinjectionchecksqlwildchars,omitempty"`
	Xmlsqlinjectiononlycheckfieldswithsqlchars string `json:"xmlsqlinjectiononlycheckfieldswithsqlchars,omitempty"`
	Xmlsqlinjectionparsecomments string `json:"xmlsqlinjectionparsecomments,omitempty"`
	Xmlsqlinjectiontype string `json:"xmlsqlinjectiontype,omitempty"`
	Xmlvalidationaction []string `json:"xmlvalidationaction,omitempty"`
	Xmlwsiaction []string `json:"xmlwsiaction,omitempty"`
	Xmlxssaction []string `json:"xmlxssaction,omitempty"`
}

type Appfwxmlcontenttype struct {
	Builtin []string `json:"builtin,omitempty"`
	Count float64 `json:"__count,omitempty"`
	Feature string `json:"feature,omitempty"`
	Isregex string `json:"isregex,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Xmlcontenttypevalue string `json:"xmlcontenttypevalue,omitempty"`
}

type Appfwmultipartformcontenttype struct {
	Builtin []string `json:"builtin,omitempty"`
	Count float64 `json:"__count,omitempty"`
	Feature string `json:"feature,omitempty"`
	Isregex string `json:"isregex,omitempty"`
	Multipartformcontenttypevalue string `json:"multipartformcontenttypevalue,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Appfwlearningsettings struct {
	Contenttypeautodeploygraceperiod int `json:"contenttypeautodeploygraceperiod,omitempty"`
	Contenttypeminthreshold int `json:"contenttypeminthreshold,omitempty"`
	Contenttypepercentthreshold int `json:"contenttypepercentthreshold,omitempty"`
	Cookieconsistencyautodeploygraceperiod int `json:"cookieconsistencyautodeploygraceperiod,omitempty"`
	Cookieconsistencyminthreshold int `json:"cookieconsistencyminthreshold,omitempty"`
	Cookieconsistencypercentthreshold int `json:"cookieconsistencypercentthreshold,omitempty"`
	Count float64 `json:"__count,omitempty"`
	Creditcardnumberminthreshold int `json:"creditcardnumberminthreshold,omitempty"`
	Creditcardnumberpercentthreshold int `json:"creditcardnumberpercentthreshold,omitempty"`
	Crosssitescriptingautodeploygraceperiod int `json:"crosssitescriptingautodeploygraceperiod,omitempty"`
	Crosssitescriptingminthreshold int `json:"crosssitescriptingminthreshold,omitempty"`
	Crosssitescriptingpercentthreshold int `json:"crosssitescriptingpercentthreshold,omitempty"`
	Csrftagautodeploygraceperiod int `json:"csrftagautodeploygraceperiod,omitempty"`
	Csrftagminthreshold int `json:"csrftagminthreshold,omitempty"`
	Csrftagpercentthreshold int `json:"csrftagpercentthreshold,omitempty"`
	Fieldconsistencyautodeploygraceperiod int `json:"fieldconsistencyautodeploygraceperiod,omitempty"`
	Fieldconsistencyminthreshold int `json:"fieldconsistencyminthreshold,omitempty"`
	Fieldconsistencypercentthreshold int `json:"fieldconsistencypercentthreshold,omitempty"`
	Fieldformatautodeploygraceperiod int `json:"fieldformatautodeploygraceperiod,omitempty"`
	Fieldformatminthreshold int `json:"fieldformatminthreshold,omitempty"`
	Fieldformatpercentthreshold int `json:"fieldformatpercentthreshold,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Profilename string `json:"profilename,omitempty"`
	Sqlinjectionautodeploygraceperiod int `json:"sqlinjectionautodeploygraceperiod,omitempty"`
	Sqlinjectionminthreshold int `json:"sqlinjectionminthreshold,omitempty"`
	Sqlinjectionpercentthreshold int `json:"sqlinjectionpercentthreshold,omitempty"`
	Starturlautodeploygraceperiod int `json:"starturlautodeploygraceperiod,omitempty"`
	Starturlminthreshold int `json:"starturlminthreshold,omitempty"`
	Starturlpercentthreshold int `json:"starturlpercentthreshold,omitempty"`
	Xmlattachmentminthreshold int `json:"xmlattachmentminthreshold,omitempty"`
	Xmlattachmentpercentthreshold int `json:"xmlattachmentpercentthreshold,omitempty"`
	Xmlwsiminthreshold int `json:"xmlwsiminthreshold,omitempty"`
	Xmlwsipercentthreshold int `json:"xmlwsipercentthreshold,omitempty"`
}

type AppfwprofileXmlxssBinding struct {
	Alertonly string `json:"alertonly,omitempty"`
	AsScanLocationXmlxss string `json:"as_scan_location_xmlxss,omitempty"`
	Comment string `json:"comment,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	IsregexXmlxss string `json:"isregex_xmlxss,omitempty"`
	Name string `json:"name,omitempty"`
	Resourceid string `json:"resourceid,omitempty"`
	Ruletype string `json:"ruletype,omitempty"`
	State string `json:"state,omitempty"`
	Xmlxss string `json:"xmlxss,omitempty"`
}

type AppfwprofileContenttypeBinding struct {
	Alertonly string `json:"alertonly,omitempty"`
	Comment string `json:"comment,omitempty"`
	Contenttype string `json:"contenttype,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	Name string `json:"name,omitempty"`
	Resourceid string `json:"resourceid,omitempty"`
	Ruletype string `json:"ruletype,omitempty"`
	State string `json:"state,omitempty"`
}

type AppfwprofileJsonsqlurlBinding struct {
	Alertonly string `json:"alertonly,omitempty"`
	AsValueExprJsonSql string `json:"as_value_expr_json_sql,omitempty"`
	AsValueTypeJsonSql string `json:"as_value_type_json_sql,omitempty"`
	Comment string `json:"comment,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	IskeyregexJsonSql string `json:"iskeyregex_json_sql,omitempty"`
	IsvalueregexJsonSql string `json:"isvalueregex_json_sql,omitempty"`
	Jsonsqlurl string `json:"jsonsqlurl,omitempty"`
	KeynameJsonSql string `json:"keyname_json_sql,omitempty"`
	Name string `json:"name,omitempty"`
	Resourceid string `json:"resourceid,omitempty"`
	Ruletype string `json:"ruletype,omitempty"`
	State string `json:"state,omitempty"`
}

type Appfwcustomsettings struct {
	Name string `json:"name,omitempty"`
	Target string `json:"target,omitempty"`
}
