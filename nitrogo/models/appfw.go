// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Appfwlearningdata struct {
	Profilename            string `json:"profilename,omitempty"`
	Starturl               string `json:"starturl,omitempty"`
	Cookieconsistency      string `json:"cookieconsistency,omitempty"`
	Fieldconsistency       string `json:"fieldconsistency,omitempty"`
	Formactionurlffc       string `json:"formactionurl_ffc,omitempty"`
	Contenttype            string `json:"contenttype,omitempty"`
	Crosssitescripting     string `json:"crosssitescripting,omitempty"`
	Formactionurlxss       string `json:"formactionurl_xss,omitempty"`
	Asscanlocationxss      string `json:"as_scan_location_xss,omitempty"`
	Asvaluetypexss         string `json:"as_value_type_xss,omitempty"`
	Asvalueexprxss         string `json:"as_value_expr_xss,omitempty"`
	Sqlinjection           string `json:"sqlinjection,omitempty"`
	Formactionurlsql       string `json:"formactionurl_sql,omitempty"`
	Asscanlocationsql      string `json:"as_scan_location_sql,omitempty"`
	Asvaluetypesql         string `json:"as_value_type_sql,omitempty"`
	Asvalueexprsql         string `json:"as_value_expr_sql,omitempty"`
	Fieldformat            string `json:"fieldformat,omitempty"`
	Formactionurlff        string `json:"formactionurl_ff,omitempty"`
	Csrftag                string `json:"csrftag,omitempty"`
	Csrfformoriginurl      string `json:"csrfformoriginurl,omitempty"`
	Creditcardnumber       string `json:"creditcardnumber,omitempty"`
	Creditcardnumberurl    string `json:"creditcardnumberurl,omitempty"`
	Xmldoscheck            string `json:"xmldoscheck,omitempty"`
	Xmlwsicheck            string `json:"xmlwsicheck,omitempty"`
	Xmlattachmentcheck     string `json:"xmlattachmentcheck,omitempty"`
	Totalxmlrequests       bool   `json:"totalxmlrequests,omitempty"`
	Securitycheck          string `json:"securitycheck,omitempty"`
	Target                 string `json:"target,omitempty"`
	Url                    string `json:"url,omitempty"`
	Name                   string `json:"name,omitempty"`
	Fieldtype              string `json:"fieldtype,omitempty"`
	Fieldformatminlength   string `json:"fieldformatminlength,omitempty"`
	Fieldformatmaxlength   string `json:"fieldformatmaxlength,omitempty"`
	Fieldformatcharmappcre string `json:"fieldformatcharmappcre,omitempty"`
	Valuetype              string `json:"value_type,omitempty"`
	Value                  string `json:"value,omitempty"`
	Hits                   string `json:"hits,omitempty"`
	Data                   string `json:"data,omitempty"`
	Nextgenapiresource     string `json:"_nextgenapiresource,omitempty"`
}

type Appfwprofilefileuploadtypebinding struct {
	Fileuploadtype            string   `json:"fileuploadtype,omitempty"`
	Asfileuploadtypesurl      string   `json:"as_fileuploadtypes_url,omitempty"`
	Isnameregex               string   `json:"isnameregex,omitempty"`
	State                     string   `json:"state,omitempty"`
	Comment                   string   `json:"comment,omitempty"`
	Filetype                  []string `json:"filetype,omitempty"`
	Isautodeployed            string   `json:"isautodeployed,omitempty"`
	Alertonly                 string   `json:"alertonly,omitempty"`
	Resourceid                string   `json:"resourceid,omitempty"`
	Name                      string   `json:"name,omitempty"`
	Isregexfileuploadtypesurl string   `json:"isregex_fileuploadtypes_url,omitempty"`
	Ruletype                  string   `json:"ruletype,omitempty"`
}

type Appfwprofilerestvalidationbinding struct {
	Restvalidation       string `json:"restvalidation,omitempty"`
	Restvalidationaction string `json:"rest_validation_action,omitempty"`
	State                string `json:"state,omitempty"`
	Comment              string `json:"comment,omitempty"`
	Isautodeployed       string `json:"isautodeployed,omitempty"`
	Alertonly            string `json:"alertonly,omitempty"`
	Resourceid           string `json:"resourceid,omitempty"`
	Name                 string `json:"name,omitempty"`
	Ruletype             string `json:"ruletype,omitempty"`
}

type Appfwxmlcontenttype struct {
	Xmlcontenttypevalue string `json:"xmlcontenttypevalue,omitempty"`
	Isregex             string `json:"isregex,omitempty"`
	Builtin             string `json:"builtin,omitempty"`
	Feature             string `json:"feature,omitempty"`
	Nextgenapiresource  string `json:"_nextgenapiresource,omitempty"`
}

type Appfwxmlschema struct {
	Name               string `json:"name,omitempty"`
	Src                string `json:"src,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Overwrite          bool   `json:"overwrite,omitempty"`
	Response           string `json:"response,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Appfwprofilecontenttypebinding struct {
	Contenttype    string `json:"contenttype,omitempty"`
	State          string `json:"state,omitempty"`
	Comment        string `json:"comment,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	Alertonly      string `json:"alertonly,omitempty"`
	Resourceid     string `json:"resourceid,omitempty"`
	Name           string `json:"name,omitempty"`
	Ruletype       string `json:"ruletype,omitempty"`
}

type Appfwprofiledenyurlbinding struct {
	Denyurl        string `json:"denyurl,omitempty"`
	State          string `json:"state,omitempty"`
	Comment        string `json:"comment,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	Alertonly      string `json:"alertonly,omitempty"`
	Resourceid     string `json:"resourceid,omitempty"`
	Name           string `json:"name,omitempty"`
	Ruletype       string `json:"ruletype,omitempty"`
}

type Appfwmultipartformcontenttype struct {
	Multipartformcontenttypevalue string `json:"multipartformcontenttypevalue,omitempty"`
	Isregex                       string `json:"isregex,omitempty"`
	Builtin                       string `json:"builtin,omitempty"`
	Feature                       string `json:"feature,omitempty"`
	Nextgenapiresource            string `json:"_nextgenapiresource,omitempty"`
}

type Appfwpolicypolicylabelbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Activepolicy           int32  `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Appfwcustomsettings struct {
	Name   string `json:"name,omitempty"`
	Target string `json:"target,omitempty"`
}

type Appfwgrpcwebtextcontenttype struct {
	Grpcwebtextcontenttypevalue string `json:"grpcwebtextcontenttypevalue,omitempty"`
	Isregex                     string `json:"isregex,omitempty"`
	Nextgenapiresource          string `json:"_nextgenapiresource,omitempty"`
}

type Appfwpolicyvpnvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Appfwprofile struct {
	Name                                       string   `json:"name,omitempty"`
	Defaults                                   string   `json:"defaults,omitempty"`
	Starturlaction                             []string `json:"starturlaction,omitempty"`
	Infercontenttypexmlpayloadaction           []string `json:"infercontenttypexmlpayloadaction,omitempty"`
	Contenttypeaction                          []string `json:"contenttypeaction,omitempty"`
	Inspectcontenttypes                        []string `json:"inspectcontenttypes,omitempty"`
	Starturlclosure                            string   `json:"starturlclosure,omitempty"`
	Denyurlaction                              []string `json:"denyurlaction,omitempty"`
	Refererheadercheck                         string   `json:"refererheadercheck,omitempty"`
	Cookieconsistencyaction                    []string `json:"cookieconsistencyaction,omitempty"`
	Cookiehijackingaction                      []string `json:"cookiehijackingaction,omitempty"`
	Cookietransforms                           string   `json:"cookietransforms,omitempty"`
	Cookieencryption                           string   `json:"cookieencryption,omitempty"`
	Cookieproxying                             string   `json:"cookieproxying,omitempty"`
	Addcookieflags                             string   `json:"addcookieflags,omitempty"`
	Fieldconsistencyaction                     []string `json:"fieldconsistencyaction,omitempty"`
	Csrftagaction                              []string `json:"csrftagaction,omitempty"`
	Crosssitescriptingaction                   []string `json:"crosssitescriptingaction,omitempty"`
	Crosssitescriptingtransformunsafehtml      string   `json:"crosssitescriptingtransformunsafehtml,omitempty"`
	Crosssitescriptingcheckcompleteurls        string   `json:"crosssitescriptingcheckcompleteurls,omitempty"`
	Sqlinjectionaction                         []string `json:"sqlinjectionaction,omitempty"`
	Cmdinjectionaction                         []string `json:"cmdinjectionaction,omitempty"`
	Cmdinjectiontype                           string   `json:"cmdinjectiontype,omitempty"`
	Sqlinjectiongrammar                        string   `json:"sqlinjectiongrammar,omitempty"`
	Cmdinjectiongrammar                        string   `json:"cmdinjectiongrammar,omitempty"`
	Fieldscan                                  string   `json:"fieldscan,omitempty"`
	Fieldscanlimit                             int      `json:"fieldscanlimit,omitempty"`
	Jsonfieldscan                              string   `json:"jsonfieldscan,omitempty"`
	Jsonfieldscanlimit                         int      `json:"jsonfieldscanlimit,omitempty"`
	Messagescan                                string   `json:"messagescan,omitempty"`
	Messagescanlimit                           int      `json:"messagescanlimit,omitempty"`
	Jsonmessagescan                            string   `json:"jsonmessagescan,omitempty"`
	Jsonmessagescanlimit                       int      `json:"jsonmessagescanlimit,omitempty"`
	Messagescanlimitcontenttypes               []string `json:"messagescanlimitcontenttypes,omitempty"`
	Sqlinjectiontransformspecialchars          string   `json:"sqlinjectiontransformspecialchars,omitempty"`
	Sqlinjectiononlycheckfieldswithsqlchars    string   `json:"sqlinjectiononlycheckfieldswithsqlchars,omitempty"`
	Sqlinjectiontype                           string   `json:"sqlinjectiontype,omitempty"`
	Sqlinjectionchecksqlwildchars              string   `json:"sqlinjectionchecksqlwildchars,omitempty"`
	Fieldformataction                          []string `json:"fieldformataction,omitempty"`
	Defaultfieldformattype                     string   `json:"defaultfieldformattype,omitempty"`
	Defaultfieldformatminlength                int      `json:"defaultfieldformatminlength,omitempty"`
	Defaultfieldformatmaxlength                int      `json:"defaultfieldformatmaxlength,omitempty"`
	Defaultfieldformatmaxoccurrences           int      `json:"defaultfieldformatmaxoccurrences,omitempty"`
	Bufferoverflowaction                       []string `json:"bufferoverflowaction,omitempty"`
	Grpcaction                                 []string `json:"grpcaction,omitempty"`
	Restaction                                 []string `json:"restaction,omitempty"`
	Bufferoverflowmaxurllength                 int      `json:"bufferoverflowmaxurllength,omitempty"`
	Bufferoverflowmaxheaderlength              int      `json:"bufferoverflowmaxheaderlength,omitempty"`
	Bufferoverflowmaxcookielength              int      `json:"bufferoverflowmaxcookielength,omitempty"`
	Bufferoverflowmaxquerylength               int      `json:"bufferoverflowmaxquerylength,omitempty"`
	Bufferoverflowmaxtotalheaderlength         int      `json:"bufferoverflowmaxtotalheaderlength,omitempty"`
	Creditcardaction                           []string `json:"creditcardaction,omitempty"`
	Creditcard                                 []string `json:"creditcard,omitempty"`
	Creditcardmaxallowed                       int      `json:"creditcardmaxallowed,omitempty"`
	Creditcardxout                             string   `json:"creditcardxout,omitempty"`
	Dosecurecreditcardlogging                  string   `json:"dosecurecreditcardlogging,omitempty"`
	Streaming                                  string   `json:"streaming,omitempty"`
	Trace                                      string   `json:"trace,omitempty"`
	Requestcontenttype                         string   `json:"requestcontenttype,omitempty"`
	Responsecontenttype                        string   `json:"responsecontenttype,omitempty"`
	Jsonerrorobject                            string   `json:"jsonerrorobject,omitempty"`
	Apispec                                    string   `json:"apispec,omitempty"`
	Protofileobject                            string   `json:"protofileobject,omitempty"`
	Jsonerrorstatuscode                        int      `json:"jsonerrorstatuscode,omitempty"`
	Jsonerrorstatusmessage                     string   `json:"jsonerrorstatusmessage,omitempty"`
	Jsondosaction                              []string `json:"jsondosaction,omitempty"`
	Jsonsqlinjectionaction                     []string `json:"jsonsqlinjectionaction,omitempty"`
	Jsonsqlinjectiontype                       string   `json:"jsonsqlinjectiontype,omitempty"`
	Jsonsqlinjectiongrammar                    string   `json:"jsonsqlinjectiongrammar,omitempty"`
	Jsoncmdinjectionaction                     []string `json:"jsoncmdinjectionaction,omitempty"`
	Jsoncmdinjectiontype                       string   `json:"jsoncmdinjectiontype,omitempty"`
	Jsoncmdinjectiongrammar                    string   `json:"jsoncmdinjectiongrammar,omitempty"`
	Jsonxssaction                              []string `json:"jsonxssaction,omitempty"`
	Xmldosaction                               []string `json:"xmldosaction,omitempty"`
	Xmlformataction                            []string `json:"xmlformataction,omitempty"`
	Xmlsqlinjectionaction                      []string `json:"xmlsqlinjectionaction,omitempty"`
	Xmlsqlinjectiononlycheckfieldswithsqlchars string   `json:"xmlsqlinjectiononlycheckfieldswithsqlchars,omitempty"`
	Xmlsqlinjectiontype                        string   `json:"xmlsqlinjectiontype,omitempty"`
	Xmlsqlinjectionchecksqlwildchars           string   `json:"xmlsqlinjectionchecksqlwildchars,omitempty"`
	Xmlsqlinjectionparsecomments               string   `json:"xmlsqlinjectionparsecomments,omitempty"`
	Xmlxssaction                               []string `json:"xmlxssaction,omitempty"`
	Xmlwsiaction                               []string `json:"xmlwsiaction,omitempty"`
	Xmlattachmentaction                        []string `json:"xmlattachmentaction,omitempty"`
	Xmlvalidationaction                        []string `json:"xmlvalidationaction,omitempty"`
	Xmlerrorobject                             string   `json:"xmlerrorobject,omitempty"`
	Xmlerrorstatuscode                         int      `json:"xmlerrorstatuscode,omitempty"`
	Xmlerrorstatusmessage                      string   `json:"xmlerrorstatusmessage,omitempty"`
	Customsettings                             string   `json:"customsettings,omitempty"`
	Signatures                                 string   `json:"signatures,omitempty"`
	Xmlsoapfaultaction                         []string `json:"xmlsoapfaultaction,omitempty"`
	Usehtmlerrorobject                         string   `json:"usehtmlerrorobject,omitempty"`
	Errorurl                                   string   `json:"errorurl,omitempty"`
	Htmlerrorobject                            string   `json:"htmlerrorobject,omitempty"`
	Htmlerrorstatuscode                        int      `json:"htmlerrorstatuscode,omitempty"`
	Htmlerrorstatusmessage                     string   `json:"htmlerrorstatusmessage,omitempty"`
	Logeverypolicyhit                          string   `json:"logeverypolicyhit,omitempty"`
	Stripcomments                              string   `json:"stripcomments,omitempty"`
	Striphtmlcomments                          string   `json:"striphtmlcomments,omitempty"`
	Stripxmlcomments                           string   `json:"stripxmlcomments,omitempty"`
	Exemptclosureurlsfromsecuritychecks        string   `json:"exemptclosureurlsfromsecuritychecks,omitempty"`
	Defaultcharset                             string   `json:"defaultcharset,omitempty"`
	Clientipexpression                         string   `json:"clientipexpression,omitempty"`
	Dynamiclearning                            []string `json:"dynamiclearning,omitempty"`
	Postbodylimit                              int      `json:"postbodylimit,omitempty"`
	Postbodylimitaction                        []string `json:"postbodylimitaction,omitempty"`
	Postbodylimitsignature                     int      `json:"postbodylimitsignature,omitempty"`
	Fileuploadmaxnum                           int      `json:"fileuploadmaxnum,omitempty"`
	Canonicalizehtmlresponse                   string   `json:"canonicalizehtmlresponse,omitempty"`
	Enableformtagging                          string   `json:"enableformtagging,omitempty"`
	Sessionlessfieldconsistency                string   `json:"sessionlessfieldconsistency,omitempty"`
	Sessionlessurlclosure                      string   `json:"sessionlessurlclosure,omitempty"`
	Semicolonfieldseparator                    string   `json:"semicolonfieldseparator,omitempty"`
	Excludefileuploadfromchecks                string   `json:"excludefileuploadfromchecks,omitempty"`
	Sqlinjectionparsecomments                  string   `json:"sqlinjectionparsecomments,omitempty"`
	Invalidpercenthandling                     string   `json:"invalidpercenthandling,omitempty"`
	Type                                       []string `json:"type,omitempty"`
	Checkrequestheaders                        string   `json:"checkrequestheaders,omitempty"`
	Inspectquerycontenttypes                   []string `json:"inspectquerycontenttypes,omitempty"`
	Optimizepartialreqs                        string   `json:"optimizepartialreqs,omitempty"`
	Urldecoderequestcookies                    string   `json:"urldecoderequestcookies,omitempty"`
	Comment                                    string   `json:"comment,omitempty"`
	Percentdecoderecursively                   string   `json:"percentdecoderecursively,omitempty"`
	Multipleheaderaction                       []string `json:"multipleheaderaction,omitempty"`
	Rfcprofile                                 string   `json:"rfcprofile,omitempty"`
	Fileuploadtypesaction                      []string `json:"fileuploadtypesaction,omitempty"`
	Verboseloglevel                            string   `json:"verboseloglevel,omitempty"`
	Insertcookiesamesiteattribute              string   `json:"insertcookiesamesiteattribute,omitempty"`
	Cookiesamesiteattribute                    string   `json:"cookiesamesiteattribute,omitempty"`
	Sqlinjectionruletype                       string   `json:"sqlinjectionruletype,omitempty"`
	Fakeaccountdetection                       string   `json:"fakeaccountdetection,omitempty"`
	Geolocationlogging                         string   `json:"geolocationlogging,omitempty"`
	Ceflogging                                 string   `json:"ceflogging,omitempty"`
	Blockkeywordaction                         []string `json:"blockkeywordaction,omitempty"`
	Jsonblockkeywordaction                     []string `json:"jsonblockkeywordaction,omitempty"`
	Asprofbypasslistenable                     string   `json:"as_prof_bypass_list_enable,omitempty"`
	Asprofdenylistenable                       string   `json:"as_prof_deny_list_enable,omitempty"`
	Sessioncookiename                          string   `json:"sessioncookiename,omitempty"`
	Archivename                                string   `json:"archivename,omitempty"`
	Relaxationrules                            bool     `json:"relaxationrules,omitempty"`
	Importprofilename                          string   `json:"importprofilename,omitempty"`
	Matchurlstring                             string   `json:"matchurlstring,omitempty"`
	Replaceurlstring                           string   `json:"replaceurlstring,omitempty"`
	Overwrite                                  bool     `json:"overwrite,omitempty"`
	Augment                                    bool     `json:"augment,omitempty"`
	State                                      string   `json:"state,omitempty"`
	Learning                                   string   `json:"learning,omitempty"`
	Csrftag                                    string   `json:"csrftag,omitempty"`
	Builtin                                    string   `json:"builtin,omitempty"`
	Nextgenapiresource                         string   `json:"_nextgenapiresource,omitempty"`
}

type Appfwglobalauditsyslogpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	State                  string `json:"state,omitempty"`
	Type                   string `json:"type,omitempty"`
	Policytype             string `json:"policytype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Appfwglobalbinding struct {
}

type Appfwpolicylabelpolicybindingbinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Appfwprofileblockkeywordbinding struct {
	Blockkeyword                   string `json:"blockkeyword,omitempty"`
	Asblockkeywordformurl          string `json:"as_blockkeyword_formurl,omitempty"`
	Fieldname                      string `json:"fieldname,omitempty"`
	Asfieldnameisregexblockkeyword string `json:"as_fieldname_isregex_blockkeyword,omitempty"`
	Blockkeywordtype               string `json:"blockkeywordtype,omitempty"`
	State                          string `json:"state,omitempty"`
	Comment                        string `json:"comment,omitempty"`
	Ruletype                       string `json:"ruletype,omitempty"`
	Isautodeployed                 string `json:"isautodeployed,omitempty"`
	Alertonly                      string `json:"alertonly,omitempty"`
	Resourceid                     string `json:"resourceid,omitempty"`
	Name                           string `json:"name,omitempty"`
}

type Appfwglobalauditnslogpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	State                  string `json:"state,omitempty"`
	Type                   string `json:"type,omitempty"`
	Policytype             string `json:"policytype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Appfwprofilexmlvalidationurlbinding struct {
	Xmlvalidationurl         string `json:"xmlvalidationurl,omitempty"`
	Xmlvalidateresponse      string `json:"xmlvalidateresponse,omitempty"`
	Xmlwsdl                  string `json:"xmlwsdl,omitempty"`
	Xmladditionalsoapheaders string `json:"xmladditionalsoapheaders,omitempty"`
	Xmlendpointcheck         string `json:"xmlendpointcheck,omitempty"`
	Xmlrequestschema         string `json:"xmlrequestschema,omitempty"`
	Xmlresponseschema        string `json:"xmlresponseschema,omitempty"`
	Xmlvalidatesoapenvelope  string `json:"xmlvalidatesoapenvelope,omitempty"`
	State                    string `json:"state,omitempty"`
	Comment                  string `json:"comment,omitempty"`
	Isautodeployed           string `json:"isautodeployed,omitempty"`
	Alertonly                string `json:"alertonly,omitempty"`
	Name                     string `json:"name,omitempty"`
	Resourceid               string `json:"resourceid,omitempty"`
	Ruletype                 string `json:"ruletype,omitempty"`
}

type Appfwprofilexmlwsiurlbinding struct {
	Xmlwsiurl      string `json:"xmlwsiurl,omitempty"`
	Xmlwsichecks   string `json:"xmlwsichecks,omitempty"`
	State          string `json:"state,omitempty"`
	Comment        string `json:"comment,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	Alertonly      string `json:"alertonly,omitempty"`
	Name           string `json:"name,omitempty"`
	Resourceid     string `json:"resourceid,omitempty"`
	Ruletype       string `json:"ruletype,omitempty"`
}

type Appfwfieldtype struct {
	Name               string `json:"name,omitempty"`
	Regex              string `json:"regex,omitempty"`
	Priority           int    `json:"priority,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Nocharmaps         bool   `json:"nocharmaps,omitempty"`
	Builtin            string `json:"builtin,omitempty"`
	Feature            string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Appfwglobalpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	State                  string `json:"state,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Numpol                 uint32 `json:"numpol,omitempty"`
	Flowtype               uint32 `json:"flowtype,omitempty"`
	Type                   string `json:"type,omitempty"`
	Policytype             string `json:"policytype,omitempty"`
	Globalbindtype         string `json:"globalbindtype,omitempty"`
}

type Appfwpolicylabelappfwpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Appfwprofilecreditcardnumberbinding struct {
	Creditcardnumber    string `json:"creditcardnumber,omitempty"`
	Creditcardnumberurl string `json:"creditcardnumberurl,omitempty"`
	State               string `json:"state,omitempty"`
	Comment             string `json:"comment,omitempty"`
	Isautodeployed      string `json:"isautodeployed,omitempty"`
	Alertonly           string `json:"alertonly,omitempty"`
	Name                string `json:"name,omitempty"`
	Resourceid          string `json:"resourceid,omitempty"`
	Ruletype            string `json:"ruletype,omitempty"`
}

type Appfwpolicy struct {
	Name               string `json:"name,omitempty"`
	Rule               string `json:"rule,omitempty"`
	Profilename        string `json:"profilename,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Logaction          string `json:"logaction,omitempty"`
	Newname            string `json:"newname,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Undefhits          string `json:"undefhits,omitempty"`
	Policytype         string `json:"policytype,omitempty"`
	Builtin            string `json:"builtin,omitempty"`
	Feature            string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Appfwhtmlerrorpage struct {
	Name               string `json:"name,omitempty"`
	Src                string `json:"src,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Overwrite          bool   `json:"overwrite,omitempty"`
	Response           string `json:"response,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Appfwprofileappfwconfidfieldbinding struct {
	Confidfield    string `json:"confidfield,omitempty"`
	Isregexcffield string `json:"isregex_cffield,omitempty"`
	Cffieldurl     string `json:"cffield_url,omitempty"`
	State          string `json:"state,omitempty"`
	Comment        string `json:"comment,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	Alertonly      string `json:"alertonly,omitempty"`
	Resourceid     string `json:"resourceid,omitempty"`
	Name           string `json:"name,omitempty"`
	Ruletype       string `json:"ruletype,omitempty"`
}

type Appfwprofilexmlxssbinding struct {
	Xmlxss               string `json:"xmlxss,omitempty"`
	Isregexxmlxss        string `json:"isregex_xmlxss,omitempty"`
	Asscanlocationxmlxss string `json:"as_scan_location_xmlxss,omitempty"`
	State                string `json:"state,omitempty"`
	Comment              string `json:"comment,omitempty"`
	Isautodeployed       string `json:"isautodeployed,omitempty"`
	Alertonly            string `json:"alertonly,omitempty"`
	Name                 string `json:"name,omitempty"`
	Resourceid           string `json:"resourceid,omitempty"`
	Ruletype             string `json:"ruletype,omitempty"`
}

type Appfwurlencodedformcontenttype struct {
	Urlencodedformcontenttypevalue string `json:"urlencodedformcontenttypevalue,omitempty"`
	Isregex                        string `json:"isregex,omitempty"`
	Builtin                        string `json:"builtin,omitempty"`
	Feature                        string `json:"feature,omitempty"`
	Nextgenapiresource             string `json:"_nextgenapiresource,omitempty"`
}

type Appfwprofilecookieconsistencybinding struct {
	Cookieconsistency string `json:"cookieconsistency,omitempty"`
	Isregex           string `json:"isregex,omitempty"`
	State             string `json:"state,omitempty"`
	Comment           string `json:"comment,omitempty"`
	Isautodeployed    string `json:"isautodeployed,omitempty"`
	Alertonly         string `json:"alertonly,omitempty"`
	Resourceid        string `json:"resourceid,omitempty"`
	Name              string `json:"name,omitempty"`
	Ruletype          string `json:"ruletype,omitempty"`
}

type Appfwprofilecrosssitescriptingbinding struct {
	Crosssitescripting string `json:"crosssitescripting,omitempty"`
	Isregexxss         string `json:"isregex_xss,omitempty"`
	Formactionurlxss   string `json:"formactionurl_xss,omitempty"`
	Asscanlocationxss  string `json:"as_scan_location_xss,omitempty"`
	Asvaluetypexss     string `json:"as_value_type_xss,omitempty"`
	Asvalueexprxss     string `json:"as_value_expr_xss,omitempty"`
	Isvalueregexxss    string `json:"isvalueregex_xss,omitempty"`
	State              string `json:"state,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Isautodeployed     string `json:"isautodeployed,omitempty"`
	Alertonly          string `json:"alertonly,omitempty"`
	Resourceid         string `json:"resourceid,omitempty"`
	Name               string `json:"name,omitempty"`
	Ruletype           string `json:"ruletype,omitempty"`
}

type Appfwprofilexmlsqlinjectionbinding struct {
	Xmlsqlinjection      string `json:"xmlsqlinjection,omitempty"`
	Isregexxmlsql        string `json:"isregex_xmlsql,omitempty"`
	Asscanlocationxmlsql string `json:"as_scan_location_xmlsql,omitempty"`
	State                string `json:"state,omitempty"`
	Comment              string `json:"comment,omitempty"`
	Isautodeployed       string `json:"isautodeployed,omitempty"`
	Alertonly            string `json:"alertonly,omitempty"`
	Name                 string `json:"name,omitempty"`
	Resourceid           string `json:"resourceid,omitempty"`
	Ruletype             string `json:"ruletype,omitempty"`
}

type Appfwsignatures struct {
	Name                    string   `json:"name,omitempty"`
	Src                     string   `json:"src,omitempty"`
	Xslt                    string   `json:"xslt,omitempty"`
	Comment                 string   `json:"comment,omitempty"`
	Overwrite               bool     `json:"overwrite,omitempty"`
	Merge                   bool     `json:"merge,omitempty"`
	Preservedefactions      bool     `json:"preservedefactions,omitempty"`
	Sha1                    string   `json:"sha1,omitempty"`
	Vendortype              string   `json:"vendortype,omitempty"`
	Autoenablenewsignatures string   `json:"autoenablenewsignatures,omitempty"`
	Ruleid                  []int    `json:"ruleid,omitempty"`
	Category                string   `json:"category,omitempty"`
	Enabled                 string   `json:"enabled,omitempty"`
	Action                  []string `json:"action,omitempty"`
	Mergedefault            bool     `json:"mergedefault,omitempty"`
	Response                string   `json:"response,omitempty"`
	Encryptedversion        string   `json:"encryptedversion,omitempty"`
	Nextgenapiresource      string   `json:"_nextgenapiresource,omitempty"`
}

type Appfwpolicylbvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Appfwpolicylabelbinding struct {
	Labelname string `json:"labelname,omitempty"`
}

type Appfwprofileexcluderescontenttypebinding struct {
	Excluderescontenttype string `json:"excluderescontenttype,omitempty"`
	State                 string `json:"state,omitempty"`
	Comment               string `json:"comment,omitempty"`
	Isautodeployed        string `json:"isautodeployed,omitempty"`
	Alertonly             string `json:"alertonly,omitempty"`
	Resourceid            string `json:"resourceid,omitempty"`
	Name                  string `json:"name,omitempty"`
	Ruletype              string `json:"ruletype,omitempty"`
}

type Appfwprofilefakeaccountbinding struct {
	Fakeaccount      string `json:"fakeaccount,omitempty"`
	Isfieldnameregex string `json:"isfieldnameregex,omitempty"`
	Formurlfad       string `json:"formurl_fad,omitempty"`
	Formexpression   string `json:"formexpression,omitempty"`
	Tag              string `json:"tag,omitempty"`
	State            string `json:"state,omitempty"`
	Comment          string `json:"comment,omitempty"`
	Isautodeployed   string `json:"isautodeployed,omitempty"`
	Alertonly        string `json:"alertonly,omitempty"`
	Name             string `json:"name,omitempty"`
	Resourceid       string `json:"resourceid,omitempty"`
	Ruletype         string `json:"ruletype,omitempty"`
}

type Appfwprofilefieldconsistencybinding struct {
	Fieldconsistency string `json:"fieldconsistency,omitempty"`
	Isregexffc       string `json:"isregex_ffc,omitempty"`
	Formactionurlffc string `json:"formactionurl_ffc,omitempty"`
	State            string `json:"state,omitempty"`
	Comment          string `json:"comment,omitempty"`
	Isautodeployed   string `json:"isautodeployed,omitempty"`
	Alertonly        string `json:"alertonly,omitempty"`
	Resourceid       string `json:"resourceid,omitempty"`
	Name             string `json:"name,omitempty"`
	Ruletype         string `json:"ruletype,omitempty"`
}

type Appfwprofilejsonxssurlbinding struct {
	Jsonxssurl          string `json:"jsonxssurl,omitempty"`
	State               string `json:"state,omitempty"`
	Comment             string `json:"comment,omitempty"`
	Iskeyregexjsonxss   string `json:"iskeyregex_json_xss,omitempty"`
	Keynamejsonxss      string `json:"keyname_json_xss,omitempty"`
	Asvaluetypejsonxss  string `json:"as_value_type_json_xss,omitempty"`
	Asvalueexprjsonxss  string `json:"as_value_expr_json_xss,omitempty"`
	Isvalueregexjsonxss string `json:"isvalueregex_json_xss,omitempty"`
	Isautodeployed      string `json:"isautodeployed,omitempty"`
	Alertonly           string `json:"alertonly,omitempty"`
	Resourceid          string `json:"resourceid,omitempty"`
	Name                string `json:"name,omitempty"`
	Ruletype            string `json:"ruletype,omitempty"`
}

type Appfwprofilexmlattachmenturlbinding struct {
	Xmlattachmenturl              string `json:"xmlattachmenturl,omitempty"`
	Xmlmaxattachmentsizecheck     string `json:"xmlmaxattachmentsizecheck,omitempty"`
	Xmlmaxattachmentsize          int    `json:"xmlmaxattachmentsize,omitempty"`
	Xmlattachmentcontenttypecheck string `json:"xmlattachmentcontenttypecheck,omitempty"`
	Xmlattachmentcontenttype      string `json:"xmlattachmentcontenttype,omitempty"`
	State                         string `json:"state,omitempty"`
	Comment                       string `json:"comment,omitempty"`
	Isautodeployed                string `json:"isautodeployed,omitempty"`
	Alertonly                     string `json:"alertonly,omitempty"`
	Name                          string `json:"name,omitempty"`
	Resourceid                    string `json:"resourceid,omitempty"`
	Ruletype                      string `json:"ruletype,omitempty"`
}

type Appfwprofiledenylistbinding struct {
	Asdenylist          string   `json:"as_deny_list,omitempty"`
	Asdenylistvaluetype string   `json:"as_deny_list_value_type,omitempty"`
	Asdenylistaction    []string `json:"as_deny_list_action,omitempty"`
	Asdenylistlocation  string   `json:"as_deny_list_location,omitempty"`
	State               string   `json:"state,omitempty"`
	Comment             string   `json:"comment,omitempty"`
	Isautodeployed      string   `json:"isautodeployed,omitempty"`
	Alertonly           string   `json:"alertonly,omitempty"`
	Resourceid          string   `json:"resourceid,omitempty"`
	Name                string   `json:"name,omitempty"`
	Ruletype            string   `json:"ruletype,omitempty"`
}

type Appfwprofilejsonblockkeywordbinding struct {
	Jsonblockkeyword           string `json:"jsonblockkeyword,omitempty"`
	Jsonblockkeywordurl        string `json:"jsonblockkeywordurl,omitempty"`
	Keynamejsonblockkeyword    string `json:"keyname_json_blockkeyword,omitempty"`
	Iskeyregexjsonblockkeyword string `json:"iskeyregex_json_blockkeyword,omitempty"`
	Jsonblockkeywordtype       string `json:"jsonblockkeywordtype,omitempty"`
	State                      string `json:"state,omitempty"`
	Comment                    string `json:"comment,omitempty"`
	Ruletype                   string `json:"ruletype,omitempty"`
	Isautodeployed             string `json:"isautodeployed,omitempty"`
	Alertonly                  string `json:"alertonly,omitempty"`
	Resourceid                 string `json:"resourceid,omitempty"`
	Name                       string `json:"name,omitempty"`
}

type Appfwglobalsyslogpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	State                  string `json:"state,omitempty"`
	Type                   string `json:"type,omitempty"`
	Policytype             string `json:"policytype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Appfwpolicyvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Activepolicy           int32  `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Appfwpolicylabel struct {
	Labelname              string `json:"labelname,omitempty"`
	Policylabeltype        string `json:"policylabeltype,omitempty"`
	Newname                string `json:"newname,omitempty"`
	Numpol                 string `json:"numpol,omitempty"`
	Hits                   string `json:"hits,omitempty"`
	Priority               string `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Description            string `json:"description,omitempty"`
	Policytype             string `json:"policytype,omitempty"`
	Nextgenapiresource     string `json:"_nextgenapiresource,omitempty"`
}

type Appfwpolicylabelpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Appfwprofilecmdinjectionbinding struct {
	Cmdinjection      string `json:"cmdinjection,omitempty"`
	Isregexcmd        string `json:"isregex_cmd,omitempty"`
	Formactionurlcmd  string `json:"formactionurl_cmd,omitempty"`
	Asscanlocationcmd string `json:"as_scan_location_cmd,omitempty"`
	Asvaluetypecmd    string `json:"as_value_type_cmd,omitempty"`
	Asvalueexprcmd    string `json:"as_value_expr_cmd,omitempty"`
	Isvalueregexcmd   string `json:"isvalueregex_cmd,omitempty"`
	State             string `json:"state,omitempty"`
	Comment           string `json:"comment,omitempty"`
	Isautodeployed    string `json:"isautodeployed,omitempty"`
	Alertonly         string `json:"alertonly,omitempty"`
	Resourceid        string `json:"resourceid,omitempty"`
	Name              string `json:"name,omitempty"`
	Ruletype          string `json:"ruletype,omitempty"`
}

type Appfwprofilelogexpressionbinding struct {
	Logexpression   string `json:"logexpression,omitempty"`
	Aslogexpression string `json:"as_logexpression,omitempty"`
	State           string `json:"state,omitempty"`
	Comment         string `json:"comment,omitempty"`
	Isautodeployed  string `json:"isautodeployed,omitempty"`
	Alertonly       string `json:"alertonly,omitempty"`
	Name            string `json:"name,omitempty"`
	Resourceid      string `json:"resourceid,omitempty"`
	Ruletype        string `json:"ruletype,omitempty"`
}

type Appfwconfidfield struct {
	Fieldname          string `json:"fieldname,omitempty"`
	Url                string `json:"url,omitempty"`
	Isregex            string `json:"isregex,omitempty"`
	Comment            string `json:"comment,omitempty"`
	State              string `json:"state,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Appfwpolicyappfwglobalbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Appfwprofilebinding struct {
	Name string `json:"name,omitempty"`
}

type Appfwprofilebypasslistbinding struct {
	Asbypasslist          string `json:"as_bypass_list,omitempty"`
	Asbypasslistvaluetype string `json:"as_bypass_list_value_type,omitempty"`
	Asbypasslistaction    string `json:"as_bypass_list_action,omitempty"`
	Asbypasslistlocation  string `json:"as_bypass_list_location,omitempty"`
	State                 string `json:"state,omitempty"`
	Comment               string `json:"comment,omitempty"`
	Isautodeployed        string `json:"isautodeployed,omitempty"`
	Alertonly             string `json:"alertonly,omitempty"`
	Resourceid            string `json:"resourceid,omitempty"`
	Name                  string `json:"name,omitempty"`
	Ruletype              string `json:"ruletype,omitempty"`
}

type Appfwprofilecsrftagbinding struct {
	Csrftag           string `json:"csrftag,omitempty"`
	Csrfformactionurl string `json:"csrfformactionurl,omitempty"`
	State             string `json:"state,omitempty"`
	Comment           string `json:"comment,omitempty"`
	Isautodeployed    string `json:"isautodeployed,omitempty"`
	Alertonly         string `json:"alertonly,omitempty"`
	Resourceid        string `json:"resourceid,omitempty"`
	Name              string `json:"name,omitempty"`
	Ruletype          string `json:"ruletype,omitempty"`
}

type Appfwtransactionrecords struct {
	Nodeid                    int    `json:"nodeid,omitempty"`
	Httptransactionid         string `json:"httptransactionid,omitempty"`
	Packetengineid            string `json:"packetengineid,omitempty"`
	Appfwsessionid            string `json:"appfwsessionid,omitempty"`
	Profilename               string `json:"profilename,omitempty"`
	Url                       string `json:"url,omitempty"`
	Clientip                  string `json:"clientip,omitempty"`
	Destip                    string `json:"destip,omitempty"`
	Starttime                 string `json:"starttime,omitempty"`
	Endtime                   string `json:"endtime,omitempty"`
	Requestcontentlength      string `json:"requestcontentlength,omitempty"`
	Requestyields             string `json:"requestyields,omitempty"`
	Requestmaxprocessingtime  string `json:"requestmaxprocessingtime,omitempty"`
	Responsecontentlength     string `json:"responsecontentlength,omitempty"`
	Responseyields            string `json:"responseyields,omitempty"`
	Responsemaxprocessingtime string `json:"responsemaxprocessingtime,omitempty"`
	Nextgenapiresource        string `json:"_nextgenapiresource,omitempty"`
}

type Appfwgrpcwebjsoncontenttype struct {
	Grpcwebjsoncontenttypevalue string `json:"grpcwebjsoncontenttypevalue,omitempty"`
	Isregex                     string `json:"isregex,omitempty"`
	Nextgenapiresource          string `json:"_nextgenapiresource,omitempty"`
}

type Appfwpolicyappfwpolicylabelbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Appfwprofilejsoncmdurlbinding struct {
	Jsoncmdurl          string `json:"jsoncmdurl,omitempty"`
	State               string `json:"state,omitempty"`
	Comment             string `json:"comment,omitempty"`
	Iskeyregexjsoncmd   string `json:"iskeyregex_json_cmd,omitempty"`
	Keynamejsoncmd      string `json:"keyname_json_cmd,omitempty"`
	Asvaluetypejsoncmd  string `json:"as_value_type_json_cmd,omitempty"`
	Asvalueexprjsoncmd  string `json:"as_value_expr_json_cmd,omitempty"`
	Isvalueregexjsoncmd string `json:"isvalueregex_json_cmd,omitempty"`
	Isautodeployed      string `json:"isautodeployed,omitempty"`
	Alertonly           string `json:"alertonly,omitempty"`
	Resourceid          string `json:"resourceid,omitempty"`
	Name                string `json:"name,omitempty"`
	Ruletype            string `json:"ruletype,omitempty"`
}

type Appfwprofilejsonsqlurlbinding struct {
	Jsonsqlurl          string `json:"jsonsqlurl,omitempty"`
	State               string `json:"state,omitempty"`
	Comment             string `json:"comment,omitempty"`
	Iskeyregexjsonsql   string `json:"iskeyregex_json_sql,omitempty"`
	Keynamejsonsql      string `json:"keyname_json_sql,omitempty"`
	Asvaluetypejsonsql  string `json:"as_value_type_json_sql,omitempty"`
	Asvalueexprjsonsql  string `json:"as_value_expr_json_sql,omitempty"`
	Isvalueregexjsonsql string `json:"isvalueregex_json_sql,omitempty"`
	Isautodeployed      string `json:"isautodeployed,omitempty"`
	Alertonly           string `json:"alertonly,omitempty"`
	Resourceid          string `json:"resourceid,omitempty"`
	Name                string `json:"name,omitempty"`
	Ruletype            string `json:"ruletype,omitempty"`
}

type Appfwprotofile struct {
	Name               string `json:"name,omitempty"`
	Src                string `json:"src,omitempty"`
	Overwrite          bool   `json:"overwrite,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Response           string `json:"response,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Appfwsettings struct {
	Defaultprofile           string   `json:"defaultprofile,omitempty"`
	Undefaction              string   `json:"undefaction,omitempty"`
	Sessiontimeout           int      `json:"sessiontimeout,omitempty"`
	Learnratelimit           int      `json:"learnratelimit,omitempty"`
	Sessionlifetime          int      `json:"sessionlifetime"`
	Sessioncookiename        string   `json:"sessioncookiename,omitempty"`
	Clientiploggingheader    string   `json:"clientiploggingheader,omitempty"`
	Importsizelimit          int      `json:"importsizelimit,omitempty"`
	Signatureautoupdate      string   `json:"signatureautoupdate,omitempty"`
	Signatureurl             string   `json:"signatureurl,omitempty"`
	Cookiepostencryptprefix  string   `json:"cookiepostencryptprefix,omitempty"`
	Logmalformedreq          string   `json:"logmalformedreq,omitempty"`
	Geolocationlogging       string   `json:"geolocationlogging,omitempty"`
	Ceflogging               string   `json:"ceflogging,omitempty"`
	Entitydecoding           string   `json:"entitydecoding,omitempty"`
	Useconfigurablesecretkey string   `json:"useconfigurablesecretkey,omitempty"`
	Sessionlimit             int      `json:"sessionlimit"`
	Malformedreqaction       []string `json:"malformedreqaction,omitempty"`
	Centralizedlearning      string   `json:"centralizedlearning,omitempty"`
	Proxyserver              string   `json:"proxyserver,omitempty"`
	Proxyport                int      `json:"proxyport,omitempty"`
	Proxyusername            string   `json:"proxyusername,omitempty"`
	Proxypassword            string   `json:"proxypassword,omitempty"`
	Cookieflags              string   `json:"cookieflags,omitempty"`
	Learning                 string   `json:"learning,omitempty"`
	Builtin                  string   `json:"builtin,omitempty"`
	Feature                  string   `json:"feature,omitempty"`
	Nextgenapiresource       string   `json:"_nextgenapiresource,omitempty"`
}

type Appfwglobalnslogpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	State                  string `json:"state,omitempty"`
	Type                   string `json:"type,omitempty"`
	Policytype             string `json:"policytype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Appfwjsoncontenttype struct {
	Jsoncontenttypevalue string `json:"jsoncontenttypevalue,omitempty"`
	Isregex              string `json:"isregex,omitempty"`
	Builtin              string `json:"builtin,omitempty"`
	Feature              string `json:"feature,omitempty"`
	Nextgenapiresource   string `json:"_nextgenapiresource,omitempty"`
}

type Appfwpolicycsvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Appfwprofilestarturlbinding struct {
	Starturl       string `json:"starturl,omitempty"`
	State          string `json:"state,omitempty"`
	Comment        string `json:"comment,omitempty"`
	Isautodeployed string `json:"isautodeployed,omitempty"`
	Alertonly      string `json:"alertonly,omitempty"`
	Resourceid     string `json:"resourceid,omitempty"`
	Name           string `json:"name,omitempty"`
	Ruletype       string `json:"ruletype,omitempty"`
}

type Appfwpolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Appfwarchive struct {
	Name               string `json:"name,omitempty"`
	Target             string `json:"target,omitempty"`
	Src                string `json:"src,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Response           string `json:"response,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Appfwglobalappfwpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	State                  string `json:"state,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Numpol                 int    `json:"numpol,omitempty"`
	Flowtype               int    `json:"flowtype,omitempty"`
	Type                   string `json:"type,omitempty"`
	Policytype             string `json:"policytype,omitempty"`
	Globalbindtype         string `json:"globalbindtype,omitempty"`
}

type Appfwgrpccontenttype struct {
	Grpccontenttypevalue string `json:"grpccontenttypevalue,omitempty"`
	Isregex              string `json:"isregex,omitempty"`
	Builtin              string `json:"builtin,omitempty"`
	Feature              string `json:"feature,omitempty"`
	Nextgenapiresource   string `json:"_nextgenapiresource,omitempty"`
}

type Appfwjsonerrorpage struct {
	Name               string `json:"name,omitempty"`
	Src                string `json:"src,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Overwrite          bool   `json:"overwrite,omitempty"`
	Response           string `json:"response,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Appfwprofiletrustedlearningclientsbinding struct {
	Trustedlearningclients string `json:"trustedlearningclients,omitempty"`
	State                  string `json:"state,omitempty"`
	Comment                string `json:"comment,omitempty"`
	Isautodeployed         string `json:"isautodeployed,omitempty"`
	Alertonly              string `json:"alertonly,omitempty"`
	Name                   string `json:"name,omitempty"`
	Resourceid             string `json:"resourceid,omitempty"`
	Ruletype               string `json:"ruletype,omitempty"`
}

type Appfwprofilexmldosurlbinding struct {
	Xmldosurl                       string `json:"xmldosurl,omitempty"`
	Xmlmaxelementdepthcheck         string `json:"xmlmaxelementdepthcheck,omitempty"`
	Xmlmaxelementdepth              int    `json:"xmlmaxelementdepth,omitempty"`
	Xmlmaxelementnamelengthcheck    string `json:"xmlmaxelementnamelengthcheck,omitempty"`
	Xmlmaxelementnamelength         int    `json:"xmlmaxelementnamelength,omitempty"`
	Xmlmaxelementscheck             string `json:"xmlmaxelementscheck,omitempty"`
	Xmlmaxelements                  int    `json:"xmlmaxelements,omitempty"`
	Xmlmaxelementchildrencheck      string `json:"xmlmaxelementchildrencheck,omitempty"`
	Xmlmaxelementchildren           int    `json:"xmlmaxelementchildren,omitempty"`
	Xmlmaxnodescheck                string `json:"xmlmaxnodescheck,omitempty"`
	Xmlmaxnodes                     int    `json:"xmlmaxnodes,omitempty"`
	Xmlmaxentityexpansionscheck     string `json:"xmlmaxentityexpansionscheck,omitempty"`
	Xmlmaxentityexpansions          int    `json:"xmlmaxentityexpansions,omitempty"`
	Xmlmaxentityexpansiondepthcheck string `json:"xmlmaxentityexpansiondepthcheck,omitempty"`
	Xmlmaxentityexpansiondepth      int    `json:"xmlmaxentityexpansiondepth,omitempty"`
	Xmlmaxattributescheck           string `json:"xmlmaxattributescheck,omitempty"`
	Xmlmaxattributes                int    `json:"xmlmaxattributes,omitempty"`
	Xmlmaxattributenamelengthcheck  string `json:"xmlmaxattributenamelengthcheck,omitempty"`
	Xmlmaxattributenamelength       int    `json:"xmlmaxattributenamelength,omitempty"`
	Xmlmaxattributevaluelengthcheck string `json:"xmlmaxattributevaluelengthcheck,omitempty"`
	Xmlmaxattributevaluelength      int    `json:"xmlmaxattributevaluelength,omitempty"`
	Xmlmaxnamespacescheck           string `json:"xmlmaxnamespacescheck,omitempty"`
	Xmlmaxnamespaces                int    `json:"xmlmaxnamespaces,omitempty"`
	Xmlmaxnamespaceurilengthcheck   string `json:"xmlmaxnamespaceurilengthcheck,omitempty"`
	Xmlmaxnamespaceurilength        int    `json:"xmlmaxnamespaceurilength,omitempty"`
	Xmlmaxchardatalengthcheck       string `json:"xmlmaxchardatalengthcheck,omitempty"`
	Xmlmaxchardatalength            int    `json:"xmlmaxchardatalength,omitempty"`
	Xmlmaxfilesizecheck             string `json:"xmlmaxfilesizecheck,omitempty"`
	Xmlmaxfilesize                  int    `json:"xmlmaxfilesize,omitempty"`
	Xmlminfilesizecheck             string `json:"xmlminfilesizecheck,omitempty"`
	Xmlminfilesize                  int    `json:"xmlminfilesize,omitempty"`
	Xmlblockpi                      string `json:"xmlblockpi,omitempty"`
	Xmlblockdtd                     string `json:"xmlblockdtd,omitempty"`
	Xmlblockexternalentities        string `json:"xmlblockexternalentities,omitempty"`
	Xmlsoaparraycheck               string `json:"xmlsoaparraycheck,omitempty"`
	Xmlmaxsoaparraysize             int    `json:"xmlmaxsoaparraysize,omitempty"`
	Xmlmaxsoaparrayrank             int    `json:"xmlmaxsoaparrayrank,omitempty"`
	State                           string `json:"state,omitempty"`
	Comment                         string `json:"comment,omitempty"`
	Isautodeployed                  string `json:"isautodeployed,omitempty"`
	Alertonly                       string `json:"alertonly,omitempty"`
	Name                            string `json:"name,omitempty"`
	Resourceid                      string `json:"resourceid,omitempty"`
	Ruletype                        string `json:"ruletype,omitempty"`
}

type Appfwprofilefieldformatbinding struct {
	Fieldformat          string `json:"fieldformat,omitempty"`
	Isregexff            string `json:"isregex_ff,omitempty"`
	Formactionurlff      string `json:"formactionurl_ff,omitempty"`
	Fieldtype            string `json:"fieldtype,omitempty"`
	Fieldformatminlength int    `json:"fieldformatminlength,omitempty"`
	Fieldformatmaxlength int    `json:"fieldformatmaxlength,omitempty"`
	State                string `json:"state,omitempty"`
	Comment              string `json:"comment,omitempty"`
	Isautodeployed       string `json:"isautodeployed,omitempty"`
	Alertonly            string `json:"alertonly,omitempty"`
	Resourceid           string `json:"resourceid,omitempty"`
	Name                 string `json:"name,omitempty"`
	Ruletype             string `json:"ruletype,omitempty"`
}

type Appfwprofilesafeobjectbinding struct {
	Safeobject     string   `json:"safeobject,omitempty"`
	Asexpression   string   `json:"as_expression,omitempty"`
	Maxmatchlength int      `json:"maxmatchlength,omitempty"`
	Action         []string `json:"action,omitempty"`
	State          string   `json:"state,omitempty"`
	Comment        string   `json:"comment,omitempty"`
	Isautodeployed string   `json:"isautodeployed,omitempty"`
	Alertonly      string   `json:"alertonly,omitempty"`
	Name           string   `json:"name,omitempty"`
	Resourceid     string   `json:"resourceid,omitempty"`
	Ruletype       string   `json:"ruletype,omitempty"`
}

type Appfwprofilesqlinjectionbinding struct {
	Sqlinjection      string `json:"sqlinjection,omitempty"`
	Isregexsql        string `json:"isregex_sql,omitempty"`
	Formactionurlsql  string `json:"formactionurl_sql,omitempty"`
	Asscanlocationsql string `json:"as_scan_location_sql,omitempty"`
	Asvaluetypesql    string `json:"as_value_type_sql,omitempty"`
	Asvalueexprsql    string `json:"as_value_expr_sql,omitempty"`
	Isvalueregexsql   string `json:"isvalueregex_sql,omitempty"`
	Ruletype          string `json:"ruletype,omitempty"`
	State             string `json:"state,omitempty"`
	Comment           string `json:"comment,omitempty"`
	Isautodeployed    string `json:"isautodeployed,omitempty"`
	Alertonly         string `json:"alertonly,omitempty"`
	Resourceid        string `json:"resourceid,omitempty"`
	Name              string `json:"name,omitempty"`
}

type Appfwlearningsettings struct {
	Profilename                             string `json:"profilename,omitempty"`
	Starturlminthreshold                    int    `json:"starturlminthreshold,omitempty"`
	Starturlpercentthreshold                int    `json:"starturlpercentthreshold,omitempty"`
	Cookieconsistencyminthreshold           int    `json:"cookieconsistencyminthreshold,omitempty"`
	Cookieconsistencypercentthreshold       int    `json:"cookieconsistencypercentthreshold,omitempty"`
	Csrftagminthreshold                     int    `json:"csrftagminthreshold,omitempty"`
	Csrftagpercentthreshold                 int    `json:"csrftagpercentthreshold,omitempty"`
	Fieldconsistencyminthreshold            int    `json:"fieldconsistencyminthreshold,omitempty"`
	Fieldconsistencypercentthreshold        int    `json:"fieldconsistencypercentthreshold,omitempty"`
	Crosssitescriptingminthreshold          int    `json:"crosssitescriptingminthreshold,omitempty"`
	Crosssitescriptingpercentthreshold      int    `json:"crosssitescriptingpercentthreshold,omitempty"`
	Sqlinjectionminthreshold                int    `json:"sqlinjectionminthreshold,omitempty"`
	Sqlinjectionpercentthreshold            int    `json:"sqlinjectionpercentthreshold,omitempty"`
	Fieldformatminthreshold                 int    `json:"fieldformatminthreshold,omitempty"`
	Fieldformatpercentthreshold             int    `json:"fieldformatpercentthreshold,omitempty"`
	Creditcardnumberminthreshold            int    `json:"creditcardnumberminthreshold,omitempty"`
	Creditcardnumberpercentthreshold        int    `json:"creditcardnumberpercentthreshold,omitempty"`
	Contenttypeminthreshold                 int    `json:"contenttypeminthreshold,omitempty"`
	Contenttypepercentthreshold             int    `json:"contenttypepercentthreshold,omitempty"`
	Xmlwsiminthreshold                      int    `json:"xmlwsiminthreshold,omitempty"`
	Xmlwsipercentthreshold                  int    `json:"xmlwsipercentthreshold,omitempty"`
	Xmlattachmentminthreshold               int    `json:"xmlattachmentminthreshold,omitempty"`
	Xmlattachmentpercentthreshold           int    `json:"xmlattachmentpercentthreshold,omitempty"`
	Fieldformatautodeploygraceperiod        int    `json:"fieldformatautodeploygraceperiod,omitempty"`
	Sqlinjectionautodeploygraceperiod       int    `json:"sqlinjectionautodeploygraceperiod,omitempty"`
	Crosssitescriptingautodeploygraceperiod int    `json:"crosssitescriptingautodeploygraceperiod,omitempty"`
	Starturlautodeploygraceperiod           int    `json:"starturlautodeploygraceperiod,omitempty"`
	Cookieconsistencyautodeploygraceperiod  int    `json:"cookieconsistencyautodeploygraceperiod,omitempty"`
	Csrftagautodeploygraceperiod            int    `json:"csrftagautodeploygraceperiod,omitempty"`
	Fieldconsistencyautodeploygraceperiod   int    `json:"fieldconsistencyautodeploygraceperiod,omitempty"`
	Contenttypeautodeploygraceperiod        int    `json:"contenttypeautodeploygraceperiod,omitempty"`
	Nextgenapiresource                      string `json:"_nextgenapiresource,omitempty"`
}

type Appfwpolicyglobalbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Activepolicy           int32  `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Appfwprofilegrpcvalidationbinding struct {
	Grpcvalidation            string `json:"grpcvalidation,omitempty"`
	Grpcrelaxvalidationaction string `json:"grpc_relax_validation_action,omitempty"`
	State                     string `json:"state,omitempty"`
	Comment                   string `json:"comment,omitempty"`
	Isautodeployed            string `json:"isautodeployed,omitempty"`
	Alertonly                 string `json:"alertonly,omitempty"`
	Resourceid                string `json:"resourceid,omitempty"`
	Name                      string `json:"name,omitempty"`
	Ruletype                  string `json:"ruletype,omitempty"`
}

type Appfwprofilejsondosurlbinding struct {
	Jsondosurl                  string `json:"jsondosurl,omitempty"`
	Jsonmaxdocumentlengthcheck  string `json:"jsonmaxdocumentlengthcheck,omitempty"`
	Jsonmaxdocumentlength       int    `json:"jsonmaxdocumentlength,omitempty"`
	Jsonmaxcontainerdepthcheck  string `json:"jsonmaxcontainerdepthcheck,omitempty"`
	Jsonmaxcontainerdepth       int    `json:"jsonmaxcontainerdepth,omitempty"`
	Jsonmaxobjectkeycountcheck  string `json:"jsonmaxobjectkeycountcheck,omitempty"`
	Jsonmaxobjectkeycount       int    `json:"jsonmaxobjectkeycount,omitempty"`
	Jsonmaxobjectkeylengthcheck string `json:"jsonmaxobjectkeylengthcheck,omitempty"`
	Jsonmaxobjectkeylength      int    `json:"jsonmaxobjectkeylength,omitempty"`
	Jsonmaxarraylengthcheck     string `json:"jsonmaxarraylengthcheck,omitempty"`
	Jsonmaxarraylength          int    `json:"jsonmaxarraylength,omitempty"`
	Jsonmaxstringlengthcheck    string `json:"jsonmaxstringlengthcheck,omitempty"`
	Jsonmaxstringlength         int    `json:"jsonmaxstringlength,omitempty"`
	State                       string `json:"state,omitempty"`
	Comment                     string `json:"comment,omitempty"`
	Isautodeployed              string `json:"isautodeployed,omitempty"`
	Alertonly                   string `json:"alertonly,omitempty"`
	Resourceid                  string `json:"resourceid,omitempty"`
	Name                        string `json:"name,omitempty"`
	Ruletype                    string `json:"ruletype,omitempty"`
}

type Appfwwsdl struct {
	Name               string `json:"name,omitempty"`
	Src                string `json:"src,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Overwrite          bool   `json:"overwrite,omitempty"`
	Response           string `json:"response,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Appfwxmlerrorpage struct {
	Name               string `json:"name,omitempty"`
	Src                string `json:"src,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Overwrite          bool   `json:"overwrite,omitempty"`
	Response           string `json:"response,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}
