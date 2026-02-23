package models

// appflow configuration structs
type AppFlowParam struct {
	AAAUsername                         string   `json:"aaausername,omitempty"`
	AnalyticsAuthToken                  string   `json:"analyticsauthtoken,omitempty"`
	AppNameRefresh                      int      `json:"appnamerefresh,omitempty"`
	AuditLogs                           string   `json:"auditlogs,omitempty"`
	Builtin                             []string `json:"builtin,omitempty"`
	CacheInsight                        string   `json:"cacheinsight,omitempty"`
	ClientTrafficOnly                   string   `json:"clienttrafficonly,omitempty"`
	ConnectionChaining                  string   `json:"connectionchaining,omitempty"`
	CQAReporting                        string   `json:"cqareporting,omitempty"`
	DistributedTracing                  string   `json:"distributedtracing,omitempty"`
	DistTracingSamplingRate             int      `json:"disttracingsamplingrate,omitempty"`
	EmailAddress                        string   `json:"emailaddress,omitempty"`
	Events                              string   `json:"events,omitempty"`
	Feature                             string   `json:"feature,omitempty"`
	FlowRecordInterval                  int      `json:"flowrecordinterval,omitempty"`
	GXSessionReporting                  string   `json:"gxsessionreporting,omitempty"`
	HTTPAuthorization                   string   `json:"httpauthorization,omitempty"`
	HTTPContentType                     string   `json:"httpcontenttype,omitempty"`
	HTTPCookie                          string   `json:"httpcookie,omitempty"`
	HTTPDomain                          string   `json:"httpdomain,omitempty"`
	HTTPHost                            string   `json:"httphost,omitempty"`
	HTTPLocation                        string   `json:"httplocation,omitempty"`
	HTTPMethod                          string   `json:"httpmethod,omitempty"`
	HTTPQueryWithURL                    string   `json:"httpquerywithurl,omitempty"`
	HTTPReferer                         string   `json:"httpreferer,omitempty"`
	HTTPSetCookie                       string   `json:"httpsetcookie,omitempty"`
	HTTPSetCookie2                      string   `json:"httpsetcookie2,omitempty"`
	HTTPURL                             string   `json:"httpurl,omitempty"`
	HTTPUserAgent                       string   `json:"httpuseragent,omitempty"`
	HTTPVia                             string   `json:"httpvia,omitempty"`
	HTTPXForwardedFor                   string   `json:"httpxforwardedfor,omitempty"`
	IdentifierName                      string   `json:"identifiername,omitempty"`
	IdentifierSessionName               string   `json:"identifiersessionname,omitempty"`
	LogStreamOverNSIP                   string   `json:"logstreamovernsip,omitempty"`
	LSNLogging                          string   `json:"lsnlogging,omitempty"`
	Metrics                             string   `json:"metrics,omitempty"`
	NextGenAPIResource                  string   `json:"_nextgenapiresource,omitempty"`
	ObservationDomainID                 int      `json:"observationdomainid,omitempty"`
	ObservationDomainName               string   `json:"observationdomainname,omitempty"`
	ObservationPointID                  int      `json:"observationpointid,omitempty"`
	SecurityInsightRecordInterval       int      `json:"securityinsightrecordinterval,omitempty"`
	SecurityInsightTraffic              string   `json:"securityinsighttraffic,omitempty"`
	SkipCacheRedirectionHTTPTransaction string   `json:"skipcacheredirectionhttptransaction,omitempty"`
	SubscriberAwareness                 string   `json:"subscriberawareness,omitempty"`
	SubscriberIDObfuscation             string   `json:"subscriberidobfuscation,omitempty"`
	SubscriberIDObfuscationAlgo         string   `json:"subscriberidobfuscationalgo,omitempty"`
	TCPAttackCounterInterval            int      `json:"tcpattackcounterinterval,omitempty"`
	TCPBurstReporting                   string   `json:"tcpburstreporting,omitempty"`
	TCPBurstReportingThreshold          int      `json:"tcpburstreportingthreshold,omitempty"`
	TemplateRefresh                     int      `json:"templaterefresh,omitempty"`
	TimeSeriesOverNSIP                  string   `json:"timeseriesovernsip,omitempty"`
	UDPPMTU                             int      `json:"udppmtu,omitempty"`
	URLCategory                         string   `json:"urlcategory,omitempty"`
	UsageRecordInterval                 int      `json:"usagerecordinterval,omitempty"`
	VideoInsight                        string   `json:"videoinsight,omitempty"`
	WebSaaSAppUsageReporting            string   `json:"websaasappusagereporting,omitempty"`
}

type AppFlowGlobalBinding struct {
	AppFlowGlobalAppFlowPolicyBinding []interface{} `json:"appflowglobal_appflowpolicy_binding,omitempty"`
}

type AppFlowPolicyAppFlowGlobalBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type AppFlowPolicyLabel struct {
	Count                  float64 `json:"__count,omitempty"`
	Description            string  `json:"description,omitempty"`
	FlowType               int     `json:"flowtype,omitempty"`
	GotoPriorityExpression string  `json:"gotopriorityexpression,omitempty"`
	Hits                   int     `json:"hits,omitempty"`
	InvokeLabelName        string  `json:"invoke_labelname,omitempty"`
	LabelName              string  `json:"labelname,omitempty"`
	LabelType              string  `json:"labeltype,omitempty"`
	NewName                string  `json:"newname,omitempty"`
	NextGenAPIResource     string  `json:"_nextgenapiresource,omitempty"`
	NumPol                 int     `json:"numpol,omitempty"`
	PolicyLabelType        string  `json:"policylabeltype,omitempty"`
	PolicyName             string  `json:"policyname,omitempty"`
	Priority               int     `json:"priority,omitempty"`
}

type AppFlowPolicyLabelBinding struct {
	AppFlowPolicyLabelAppFlowPolicyBinding []interface{} `json:"appflowpolicylabel_appflowpolicy_binding,omitempty"`
	LabelName                              string        `json:"labelname,omitempty"`
}

type AppFlowPolicyVPNVServerBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type AppFlowGlobalAppFlowPolicyBinding struct {
	FlowType               int    `json:"flowtype,omitempty"`
	GlobalBindType         string `json:"globalbindtype,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	NumPol                 int    `json:"numpol,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
}

type AppFlowPolicyAppFlowPolicyLabelBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type AppFlowActionBinding struct {
	AppFlowActionAnalyticsProfileBinding []interface{} `json:"appflowaction_analyticsprofile_binding,omitempty"`
	Name                                 string        `json:"name,omitempty"`
}

type AppFlowPolicyLabelAppFlowPolicyBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelName        string `json:"invoke_labelname,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type AppFlowPolicyBinding struct {
	AppFlowPolicyAppFlowGlobalBinding      []interface{} `json:"appflowpolicy_appflowglobal_binding,omitempty"`
	AppFlowPolicyAppFlowPolicyLabelBinding []interface{} `json:"appflowpolicy_appflowpolicylabel_binding,omitempty"`
	AppFlowPolicyCSVServerBinding          []interface{} `json:"appflowpolicy_csvserver_binding,omitempty"`
	AppFlowPolicyLBVServerBinding          []interface{} `json:"appflowpolicy_lbvserver_binding,omitempty"`
	AppFlowPolicyVPNVServerBinding         []interface{} `json:"appflowpolicy_vpnvserver_binding,omitempty"`
	Name                                   string        `json:"name,omitempty"`
}

type AppFlowPolicyLBVServerBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type AppFlowAction struct {
	BotInsight             string   `json:"botinsight,omitempty"`
	CIInsight              string   `json:"ciinsight,omitempty"`
	ClientSideMeasurements string   `json:"clientsidemeasurements,omitempty"`
	Collectors             []string `json:"collectors,omitempty"`
	Comment                string   `json:"comment,omitempty"`
	Count                  float64  `json:"__count,omitempty"`
	Description            string   `json:"description,omitempty"`
	DistributionAlgorithm  string   `json:"distributionalgorithm,omitempty"`
	Hits                   int      `json:"hits,omitempty"`
	MetricsLog             bool     `json:"metricslog,omitempty"`
	Name                   string   `json:"name,omitempty"`
	NewName                string   `json:"newname,omitempty"`
	NextGenAPIResource     string   `json:"_nextgenapiresource,omitempty"`
	PageTracking           string   `json:"pagetracking,omitempty"`
	ReferenceCount         int      `json:"referencecount,omitempty"`
	SecurityInsight        string   `json:"securityinsight,omitempty"`
	TransactionLog         string   `json:"transactionlog,omitempty"`
	VideoAnalytics         string   `json:"videoanalytics,omitempty"`
	WebInsight             string   `json:"webinsight,omitempty"`
}

type AppFlowPolicy struct {
	Action             string   `json:"action,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	Comment            string   `json:"comment,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Description        string   `json:"description,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	Name               string   `json:"name,omitempty"`
	NewName            string   `json:"newname,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	Rule               string   `json:"rule,omitempty"`
	UndefAction        string   `json:"undefaction,omitempty"`
	UndefHits          int      `json:"undefhits,omitempty"`
}

type AppFlowPolicyCSVServerBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type AppFlowCollector struct {
	Count              float64 `json:"__count,omitempty"`
	IPAddress          string  `json:"ipaddress,omitempty"`
	Name               string  `json:"name,omitempty"`
	NetProfile         string  `json:"netprofile,omitempty"`
	NewName            string  `json:"newname,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Port               int     `json:"port,omitempty"`
	State              string  `json:"state,omitempty"`
	Transport          string  `json:"transport,omitempty"`
}

type AppFlowActionAnalyticsProfileBinding struct {
	AnalyticsProfile string `json:"analyticsprofile,omitempty"`
	Name             string `json:"name,omitempty"`
}
