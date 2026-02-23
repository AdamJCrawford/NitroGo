package models

// lb configuration structs
type LBVServerRewritePolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LBVServerFEOPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LBMonBindings struct {
	BoundServiceGroupSvrState string  `json:"boundservicegroupsvrstate,omitempty"`
	Count                     float64 `json:"__count,omitempty"`
	MonitorName               string  `json:"monitorname,omitempty"`
	NextGenAPIResource        string  `json:"_nextgenapiresource,omitempty"`
	State                     string  `json:"state,omitempty"`
	TypeField                 string  `json:"type,omitempty"`
}

type LBVServerCSVServerBinding struct {
	CacheType     string `json:"cachetype,omitempty"`
	CacheVServer  string `json:"cachevserver,omitempty"`
	Hits          int    `json:"hits,omitempty"`
	Name          string `json:"name,omitempty"`
	Order         int    `json:"order,omitempty"`
	PIPolicyHits  int    `json:"pipolicyhits,omitempty"`
	PolicyName    string `json:"policyname,omitempty"`
	PolicySubType int    `json:"policysubtype,omitempty"`
	Priority      int    `json:"priority,omitempty"`
}

type LBMonitorBinding struct {
	LBMonitorMetricBinding     []interface{} `json:"lbmonitor_metric_binding,omitempty"`
	LBMonitorSSLCertKeyBinding []interface{} `json:"lbmonitor_sslcertkey_binding,omitempty"`
	MonitorName                string        `json:"monitorname,omitempty"`
}

type LBMonBindingsBinding struct {
	LBMonBindingsGSLBServiceGroupBinding []interface{} `json:"lbmonbindings_gslbservicegroup_binding,omitempty"`
	LBMonBindingsServiceBinding          []interface{} `json:"lbmonbindings_service_binding,omitempty"`
	LBMonBindingsServiceGroupBinding     []interface{} `json:"lbmonbindings_servicegroup_binding,omitempty"`
	MonitorName                          string        `json:"monitorname,omitempty"`
}

type LBVServerBotPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LBRoute6 struct {
	Count              float64 `json:"__count,omitempty"`
	Flags              string  `json:"flags,omitempty"`
	GatewayName        string  `json:"gatewayname,omitempty"`
	Network            string  `json:"network,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	TD                 int     `json:"td,omitempty"`
}

type LBPolicyGSLBVServerBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LBVServerResponderPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LBMonBindingsServiceBinding struct {
	IPAddress   string `json:"ipaddress,omitempty"`
	MonitorName string `json:"monitorname,omitempty"`
	MonSvcState string `json:"monsvcstate,omitempty"`
	Port        int    `json:"port,omitempty"`
	ServiceName string `json:"servicename,omitempty"`
	ServiceType string `json:"servicetype,omitempty"`
	SvrState    string `json:"svrstate,omitempty"`
}

type LBVServerTransformPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LBAction struct {
	Builtin            []string      `json:"builtin,omitempty"`
	Comment            string        `json:"comment,omitempty"`
	Count              float64       `json:"__count,omitempty"`
	Feature            string        `json:"feature,omitempty"`
	Hits               int           `json:"hits,omitempty"`
	Name               string        `json:"name,omitempty"`
	NewName            string        `json:"newname,omitempty"`
	NextGenAPIResource string        `json:"_nextgenapiresource,omitempty"`
	ReferenceCount     int           `json:"referencecount,omitempty"`
	TypeField          string        `json:"type,omitempty"`
	UndefHits          int           `json:"undefhits,omitempty"`
	Value              []interface{} `json:"value,omitempty"`
}

type LBVServerAuditSyslogPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LBGlobalBinding struct {
	LBGlobalLBPolicyBinding []interface{} `json:"lbglobal_lbpolicy_binding,omitempty"`
}

type LBVServerAppFlowPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LBRoute struct {
	Count              float64 `json:"__count,omitempty"`
	Flags              string  `json:"flags,omitempty"`
	GatewayName        string  `json:"gatewayname,omitempty"`
	Netmask            string  `json:"netmask,omitempty"`
	Network            string  `json:"network,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	TD                 int     `json:"td,omitempty"`
}

type LBMonitorServiceBinding struct {
	DupState         string `json:"dup_state,omitempty"`
	DupWeight        int    `json:"dup_weight,omitempty"`
	MonitorName      string `json:"monitorname,omitempty"`
	ServiceGroupName string `json:"servicegroupname,omitempty"`
	ServiceName      string `json:"servicename,omitempty"`
	State            string `json:"state,omitempty"`
	Weight           int    `json:"weight,omitempty"`
}

type LBVServerAppFWPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LBVServerCachePolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LBVServerBinding struct {
	LBVServerAnalyticsProfileBinding                 []interface{} `json:"lbvserver_analyticsprofile_binding,omitempty"`
	LBVServerAppFlowPolicyBinding                    []interface{} `json:"lbvserver_appflowpolicy_binding,omitempty"`
	LBVServerAppFWPolicyBinding                      []interface{} `json:"lbvserver_appfwpolicy_binding,omitempty"`
	LBVServerAppQOEPolicyBinding                     []interface{} `json:"lbvserver_appqoepolicy_binding,omitempty"`
	LBVServerAuditNSLogPolicyBinding                 []interface{} `json:"lbvserver_auditnslogpolicy_binding,omitempty"`
	LBVServerAuditSyslogPolicyBinding                []interface{} `json:"lbvserver_auditsyslogpolicy_binding,omitempty"`
	LBVServerAuthorizationPolicyBinding              []interface{} `json:"lbvserver_authorizationpolicy_binding,omitempty"`
	LBVServerBotPolicyBinding                        []interface{} `json:"lbvserver_botpolicy_binding,omitempty"`
	LBVServerCachePolicyBinding                      []interface{} `json:"lbvserver_cachepolicy_binding,omitempty"`
	LBVServerCMPPolicyBinding                        []interface{} `json:"lbvserver_cmppolicy_binding,omitempty"`
	LBVServerContentInspectionPolicyBinding          []interface{} `json:"lbvserver_contentinspectionpolicy_binding,omitempty"`
	LBVServerCSVServerBinding                        []interface{} `json:"lbvserver_csvserver_binding,omitempty"`
	LBVServerDNSPolicy64Binding                      []interface{} `json:"lbvserver_dnspolicy64_binding,omitempty"`
	LBVServerFEOPolicyBinding                        []interface{} `json:"lbvserver_feopolicy_binding,omitempty"`
	LBVServerLBPolicyBinding                         []interface{} `json:"lbvserver_lbpolicy_binding,omitempty"`
	LBVServerResponderPolicyBinding                  []interface{} `json:"lbvserver_responderpolicy_binding,omitempty"`
	LBVServerRewritePolicyBinding                    []interface{} `json:"lbvserver_rewritepolicy_binding,omitempty"`
	LBVServerServiceBinding                          []interface{} `json:"lbvserver_service_binding,omitempty"`
	LBVServerServiceGroupBinding                     []interface{} `json:"lbvserver_servicegroup_binding,omitempty"`
	LBVServerServiceGroupMemberBinding               []interface{} `json:"lbvserver_servicegroupmember_binding,omitempty"`
	LBVServerSpilloverPolicyBinding                  []interface{} `json:"lbvserver_spilloverpolicy_binding,omitempty"`
	LBVServerTMTrafficPolicyBinding                  []interface{} `json:"lbvserver_tmtrafficpolicy_binding,omitempty"`
	LBVServerTransformPolicyBinding                  []interface{} `json:"lbvserver_transformpolicy_binding,omitempty"`
	LBVServerVideoOptimizationDetectionPolicyBinding []interface{} `json:"lbvserver_videooptimizationdetectionpolicy_binding,omitempty"`
	LBVServerVideoOptimizationPacingPolicyBinding    []interface{} `json:"lbvserver_videooptimizationpacingpolicy_binding,omitempty"`
	Name                                             string        `json:"name,omitempty"`
}

type LBPolicyLabelPolicyBindingBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelName        string `json:"invoke_labelname,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LBVServerContentInspectionPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LBGlobalLBPolicyBinding struct {
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

type LBVServerServiceBinding struct {
	CookieIPPort      string `json:"cookieipport,omitempty"`
	CookieName        string `json:"cookiename,omitempty"`
	CurState          string `json:"curstate,omitempty"`
	DynamicWeight     int    `json:"dynamicweight,omitempty"`
	IPv46             string `json:"ipv46,omitempty"`
	Name              string `json:"name,omitempty"`
	Order             int    `json:"order,omitempty"`
	OrderStr          string `json:"orderstr,omitempty"`
	Port              int    `json:"port,omitempty"`
	PreferredLocation string `json:"preferredlocation,omitempty"`
	ServiceGroupName  string `json:"servicegroupname,omitempty"`
	ServiceName       string `json:"servicename,omitempty"`
	ServiceType       string `json:"servicetype,omitempty"`
	VServerID         string `json:"vserverid,omitempty"`
	VSvrBindSvcIP     string `json:"vsvrbindsvcip,omitempty"`
	VSvrBindSvcPort   int    `json:"vsvrbindsvcport,omitempty"`
	Weight            int    `json:"weight,omitempty"`
}

type LBVServerAppQOEPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LBPolicyLBGlobalBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LBGroup struct {
	BackupPersistenceTimeout int     `json:"backuppersistencetimeout,omitempty"`
	CookieDomain             string  `json:"cookiedomain,omitempty"`
	CookieName               string  `json:"cookiename,omitempty"`
	Count                    float64 `json:"__count,omitempty"`
	MasterVServer            string  `json:"mastervserver,omitempty"`
	Name                     string  `json:"name,omitempty"`
	NewName                  string  `json:"newname,omitempty"`
	NextGenAPIResource       string  `json:"_nextgenapiresource,omitempty"`
	PersistenceBackup        string  `json:"persistencebackup,omitempty"`
	PersistenceType          string  `json:"persistencetype,omitempty"`
	PersistMask              string  `json:"persistmask,omitempty"`
	Rule                     string  `json:"rule,omitempty"`
	TD                       int     `json:"td,omitempty"`
	Timeout                  int     `json:"timeout,omitempty"`
	UseVServerPersistency    string  `json:"usevserverpersistency,omitempty"`
	V6PersistMaskLen         int     `json:"v6persistmasklen,omitempty"`
}

type LBPolicyLabelLBPolicyBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelName        string `json:"invoke_labelname,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LBSIPParameters struct {
	AddrPortVIP         string   `json:"addrportvip,omitempty"`
	Builtin             []string `json:"builtin,omitempty"`
	Feature             string   `json:"feature,omitempty"`
	NextGenAPIResource  string   `json:"_nextgenapiresource,omitempty"`
	RetryDur            int      `json:"retrydur,omitempty"`
	RNATDstPort         int      `json:"rnatdstport,omitempty"`
	RNATSecureDstPort   int      `json:"rnatsecuredstport,omitempty"`
	RNATSecureSrcPort   int      `json:"rnatsecuresrcport,omitempty"`
	RNATSrcPort         int      `json:"rnatsrcport,omitempty"`
	SIP503RateThreshold int      `json:"sip503ratethreshold,omitempty"`
}

type LBVServerAnalyticsProfileBinding struct {
	AnalyticsProfile string `json:"analyticsprofile,omitempty"`
	Name             string `json:"name,omitempty"`
	Order            int    `json:"order,omitempty"`
}

type LBWLMLBVServerBinding struct {
	VServerName string `json:"vservername,omitempty"`
	WLMName     string `json:"wlmname,omitempty"`
}

type LBVServerCMPPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LBVServer struct {
	ActiveServices                     int           `json:"activeservices,omitempty"`
	ADFSProxyProfile                   string        `json:"adfsproxyprofile,omitempty"`
	APIProfile                         string        `json:"apiprofile,omitempty"`
	AppFlowLog                         string        `json:"appflowlog,omitempty"`
	Authentication                     string        `json:"authentication,omitempty"`
	AuthenticationHost                 string        `json:"authenticationhost,omitempty"`
	Authn401                           string        `json:"authn401,omitempty"`
	AuthnProfile                       string        `json:"authnprofile,omitempty"`
	AuthnVSName                        string        `json:"authnvsname,omitempty"`
	BackupLBMethod                     string        `json:"backuplbmethod,omitempty"`
	BackupPersistenceTimeout           int           `json:"backuppersistencetimeout,omitempty"`
	BackupVServer                      string        `json:"backupvserver,omitempty"`
	BackupVServerStatus                string        `json:"backupvserverstatus,omitempty"`
	BindPoint                          string        `json:"bindpoint,omitempty"`
	BypassAAAA                         string        `json:"bypassaaaa,omitempty"`
	Cacheable                          string        `json:"cacheable,omitempty"`
	CacheVServer                       string        `json:"cachevserver,omitempty"`
	CltTimeout                         int           `json:"clttimeout,omitempty"`
	Comment                            string        `json:"comment,omitempty"`
	ConnFailover                       string        `json:"connfailover,omitempty"`
	ConsolidatedLConn                  string        `json:"consolidatedlconn,omitempty"`
	ConsolidatedLConnGbl               string        `json:"consolidatedlconngbl,omitempty"`
	CookieDomain                       string        `json:"cookiedomain,omitempty"`
	CookieName                         string        `json:"cookiename,omitempty"`
	Count                              float64       `json:"__count,omitempty"`
	CurrentActiveOrder                 string        `json:"currentactiveorder,omitempty"`
	CurState                           string        `json:"curstate,omitempty"`
	DataLength                         int           `json:"datalength,omitempty"`
	DataOffset                         int           `json:"dataoffset,omitempty"`
	DBProfileName                      string        `json:"dbprofilename,omitempty"`
	DBSLB                              string        `json:"dbslb,omitempty"`
	DisablePrimaryOnDown               string        `json:"disableprimaryondown,omitempty"`
	DNS64                              string        `json:"dns64,omitempty"`
	DNSOverHTTPS                       string        `json:"dnsoverhttps,omitempty"`
	DNSProfileName                     string        `json:"dnsprofilename,omitempty"`
	DNSVServerName                     string        `json:"dnsvservername,omitempty"`
	Domain                             string        `json:"domain,omitempty"`
	DownStateFlush                     string        `json:"downstateflush,omitempty"`
	EffectiveState                     string        `json:"effectivestate,omitempty"`
	GroupName                          string        `json:"groupname,omitempty"`
	Gt2gb                              string        `json:"gt2gb,omitempty"`
	HashLength                         int           `json:"hashlength,omitempty"`
	Health                             int           `json:"health,omitempty"`
	HealthThreshold                    int           `json:"healththreshold,omitempty"`
	HomePage                           string        `json:"homepage,omitempty"`
	HTTPProfileName                    string        `json:"httpprofilename,omitempty"`
	HTTPSRedirectURL                   string        `json:"httpsredirecturl,omitempty"`
	ICMPVsrResponse                    string        `json:"icmpvsrresponse,omitempty"`
	InsertVServerIPPort                string        `json:"insertvserveripport,omitempty"`
	IPMapping                          string        `json:"ipmapping,omitempty"`
	IPMask                             string        `json:"ipmask,omitempty"`
	IPPattern                          string        `json:"ippattern,omitempty"`
	IPSet                              string        `json:"ipset,omitempty"`
	IPv46                              string        `json:"ipv46,omitempty"`
	IsGSLB                             bool          `json:"isgslb,omitempty"`
	L2Conn                             string        `json:"l2conn,omitempty"`
	LBMethod                           string        `json:"lbmethod,omitempty"`
	LBProfileName                      string        `json:"lbprofilename,omitempty"`
	LBRRReason                         int           `json:"lbrrreason,omitempty"`
	ListenPolicy                       string        `json:"listenpolicy,omitempty"`
	ListenPriority                     int           `json:"listenpriority,omitempty"`
	M                                  string        `json:"m,omitempty"`
	MACModeRetainVLAN                  string        `json:"macmoderetainvlan,omitempty"`
	MapField                           string        `json:"map,omitempty"`
	MaxAutoscaleMembers                int           `json:"maxautoscalemembers,omitempty"`
	MinAutoscaleMembers                int           `json:"minautoscalemembers,omitempty"`
	MSSQLServerVersion                 string        `json:"mssqlserverversion,omitempty"`
	MySQLCharacterSet                  int           `json:"mysqlcharacterset,omitempty"`
	MySQLProtocolVersion               int           `json:"mysqlprotocolversion,omitempty"`
	MySQLServerCapabilities            int           `json:"mysqlservercapabilities,omitempty"`
	MySQLServerVersion                 string        `json:"mysqlserverversion,omitempty"`
	Name                               string        `json:"name,omitempty"`
	Netmask                            string        `json:"netmask,omitempty"`
	NetProfile                         string        `json:"netprofile,omitempty"`
	NewName                            string        `json:"newname,omitempty"`
	NewServiceRequest                  int           `json:"newservicerequest,omitempty"`
	NewServiceRequestIncrementInterval int           `json:"newservicerequestincrementinterval,omitempty"`
	NewServiceRequestUnit              string        `json:"newservicerequestunit,omitempty"`
	NextGenAPIResource                 string        `json:"_nextgenapiresource,omitempty"`
	NgName                             string        `json:"ngname,omitempty"`
	NoDefaultBindings                  string        `json:"nodefaultbindings,omitempty"`
	OracleServerVersion                string        `json:"oracleserverversion,omitempty"`
	Order                              int           `json:"order,omitempty"`
	OrderThreshold                     int           `json:"orderthreshold,omitempty"`
	PersistAVPNo                       []interface{} `json:"persistavpno,omitempty"`
	PersistenceBackup                  string        `json:"persistencebackup,omitempty"`
	PersistenceType                    string        `json:"persistencetype,omitempty"`
	PersistMask                        string        `json:"persistmask,omitempty"`
	Port                               int           `json:"port,omitempty"`
	Precedence                         string        `json:"precedence,omitempty"`
	ProbePort                          int           `json:"probeport,omitempty"`
	ProbeProtocol                      string        `json:"probeprotocol,omitempty"`
	ProbeSuccessResponseCode           string        `json:"probesuccessresponsecode,omitempty"`
	ProcessLocal                       string        `json:"processlocal,omitempty"`
	Push                               string        `json:"push,omitempty"`
	PushLabel                          string        `json:"pushlabel,omitempty"`
	PushMultiClients                   string        `json:"pushmulticlients,omitempty"`
	PushVServer                        string        `json:"pushvserver,omitempty"`
	QUICBridgeProfileName              string        `json:"quicbridgeprofilename,omitempty"`
	QUICProfileName                    string        `json:"quicprofilename,omitempty"`
	Range                              int           `json:"range,omitempty"`
	RecursionAvailable                 string        `json:"recursionavailable,omitempty"`
	Redirect                           string        `json:"redirect,omitempty"`
	RedirectFromPort                   int           `json:"redirectfromport,omitempty"`
	RedirectPortRewrite                string        `json:"redirectportrewrite,omitempty"`
	RedirURL                           string        `json:"redirurl,omitempty"`
	RedirURLFlags                      bool          `json:"redirurlflags,omitempty"`
	ResRule                            string        `json:"resrule,omitempty"`
	RetainConnectionsOnCluster         string        `json:"retainconnectionsoncluster,omitempty"`
	RHIState                           string        `json:"rhistate,omitempty"`
	RTSPNAT                            string        `json:"rtspnat,omitempty"`
	Rule                               string        `json:"rule,omitempty"`
	RuleType                           int           `json:"ruletype,omitempty"`
	ServiceName                        string        `json:"servicename,omitempty"`
	ServiceType                        string        `json:"servicetype,omitempty"`
	Sessionless                        string        `json:"sessionless,omitempty"`
	SkipPersistency                    string        `json:"skippersistency,omitempty"`
	SOBackupAction                     string        `json:"sobackupaction,omitempty"`
	SOMethod                           string        `json:"somethod,omitempty"`
	SOPersistence                      string        `json:"sopersistence,omitempty"`
	SOPersistenceTimeout               int           `json:"sopersistencetimeout,omitempty"`
	SOThreshold                        int           `json:"sothreshold,omitempty"`
	State                              string        `json:"state,omitempty"`
	StateChangeTimeMsec                int           `json:"statechangetimemsec,omitempty"`
	StateChangeTimeSec                 string        `json:"statechangetimesec,omitempty"`
	StateChangeTimeSeconds             int           `json:"statechangetimeseconds,omitempty"`
	Status                             int           `json:"status,omitempty"`
	TCPProbePort                       int           `json:"tcpprobeport,omitempty"`
	TCPProfileName                     string        `json:"tcpprofilename,omitempty"`
	TD                                 int           `json:"td,omitempty"`
	ThresholdValue                     int           `json:"thresholdvalue,omitempty"`
	TicksSinceLastStateChange          int           `json:"tickssincelaststatechange,omitempty"`
	Timeout                            int           `json:"timeout,omitempty"`
	ToggleOrder                        string        `json:"toggleorder,omitempty"`
	TOSID                              int           `json:"tosid,omitempty"`
	TotalServices                      int           `json:"totalservices,omitempty"`
	TROFSPersistence                   string        `json:"trofspersistence,omitempty"`
	TypeField                          string        `json:"type,omitempty"`
	V6NetmaskLen                       int           `json:"v6netmasklen,omitempty"`
	V6PersistMaskLen                   int           `json:"v6persistmasklen,omitempty"`
	Value                              string        `json:"value,omitempty"`
	Version                            int           `json:"version,omitempty"`
	VIPHeader                          string        `json:"vipheader,omitempty"`
	VSvrDynConnSOThreshold             int           `json:"vsvrdynconnsothreshold,omitempty"`
	Weight                             int           `json:"weight,omitempty"`
}

type LBVServerTMTrafficPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LBMonitor struct {
	AcctApplicationID                []interface{} `json:"acctapplicationid,omitempty"`
	Action                           string        `json:"action,omitempty"`
	AlertRetries                     int           `json:"alertretries,omitempty"`
	Application                      string        `json:"application,omitempty"`
	Attribute                        string        `json:"attribute,omitempty"`
	AuthApplicationID                []interface{} `json:"authapplicationid,omitempty"`
	BaseDN                           string        `json:"basedn,omitempty"`
	BindDN                           string        `json:"binddn,omitempty"`
	Count                            float64       `json:"__count,omitempty"`
	CustomHeaders                    string        `json:"customheaders,omitempty"`
	Database                         string        `json:"database,omitempty"`
	DestIP                           string        `json:"destip,omitempty"`
	DestPort                         int           `json:"destport,omitempty"`
	Deviation                        int           `json:"deviation,omitempty"`
	DispatcherIP                     string        `json:"dispatcherip,omitempty"`
	DispatcherPort                   int           `json:"dispatcherport,omitempty"`
	Domain                           string        `json:"domain,omitempty"`
	Downtime                         int           `json:"downtime,omitempty"`
	DupState                         string        `json:"dup_state,omitempty"`
	DupWeight                        int           `json:"dup_weight,omitempty"`
	DynamicInterval                  int           `json:"dynamicinterval,omitempty"`
	DynamicResponseTimeout           int           `json:"dynamicresponsetimeout,omitempty"`
	EvalRule                         string        `json:"evalrule,omitempty"`
	FailureRetries                   int           `json:"failureretries,omitempty"`
	FileName                         string        `json:"filename,omitempty"`
	Filter                           string        `json:"filter,omitempty"`
	FirmwareRevision                 int           `json:"firmwarerevision,omitempty"`
	Group                            string        `json:"group,omitempty"`
	GRPCHealthCheck                  string        `json:"grpchealthcheck,omitempty"`
	GRPCServiceName                  string        `json:"grpcservicename,omitempty"`
	GRPCStatusCode                   []interface{} `json:"grpcstatuscode,omitempty"`
	HostIPAddress                    string        `json:"hostipaddress,omitempty"`
	HostName                         string        `json:"hostname,omitempty"`
	HTTPRequest                      string        `json:"httprequest,omitempty"`
	InBandSecurityID                 string        `json:"inbandsecurityid,omitempty"`
	Interval                         int           `json:"interval,omitempty"`
	IPAddress                        []string      `json:"ipaddress,omitempty"`
	IPTunnel                         string        `json:"iptunnel,omitempty"`
	KCDAccount                       string        `json:"kcdaccount,omitempty"`
	LASVersion                       string        `json:"lasversion,omitempty"`
	LogonPointName                   string        `json:"logonpointname,omitempty"`
	LRTM                             string        `json:"lrtm,omitempty"`
	LRTMConf                         int           `json:"lrtmconf,omitempty"`
	LRTMConfStr                      string        `json:"lrtmconfstr,omitempty"`
	MaxForwards                      int           `json:"maxforwards,omitempty"`
	Metric                           string        `json:"metric,omitempty"`
	MetricTable                      string        `json:"metrictable,omitempty"`
	MetricThreshold                  int           `json:"metricthreshold,omitempty"`
	MetricWeight                     int           `json:"metricweight,omitempty"`
	MonitorName                      string        `json:"monitorname,omitempty"`
	MQTTClientIdentifier             string        `json:"mqttclientidentifier,omitempty"`
	MQTTVersion                      int           `json:"mqttversion,omitempty"`
	MSSQLProtocolVersion             string        `json:"mssqlprotocolversion,omitempty"`
	MultiMetricTable                 []string      `json:"multimetrictable,omitempty"`
	NetProfile                       string        `json:"netprofile,omitempty"`
	NextGenAPIResource               string        `json:"_nextgenapiresource,omitempty"`
	OracleSID                        string        `json:"oraclesid,omitempty"`
	OriginHost                       string        `json:"originhost,omitempty"`
	OriginRealm                      string        `json:"originrealm,omitempty"`
	Password                         string        `json:"password,omitempty"`
	ProductName                      string        `json:"productname,omitempty"`
	Query                            string        `json:"query,omitempty"`
	QueryType                        string        `json:"querytype,omitempty"`
	RADAccountSession                string        `json:"radaccountsession,omitempty"`
	RADAccountType                   int           `json:"radaccounttype,omitempty"`
	RADAPN                           string        `json:"radapn,omitempty"`
	RADFramedIP                      string        `json:"radframedip,omitempty"`
	RADKey                           string        `json:"radkey,omitempty"`
	RADMSISDN                        string        `json:"radmsisdn,omitempty"`
	RADNASID                         string        `json:"radnasid,omitempty"`
	RADNASIP                         string        `json:"radnasip,omitempty"`
	Recv                             string        `json:"recv,omitempty"`
	RespCode                         []string      `json:"respcode,omitempty"`
	RespTimeout                      int           `json:"resptimeout,omitempty"`
	RespTimeoutThresh                int           `json:"resptimeoutthresh,omitempty"`
	Retries                          int           `json:"retries,omitempty"`
	Reverse                          string        `json:"reverse,omitempty"`
	RTSPRequest                      string        `json:"rtsprequest,omitempty"`
	ScriptArgs                       string        `json:"scriptargs,omitempty"`
	ScriptName                       string        `json:"scriptname,omitempty"`
	SecondaryPassword                string        `json:"secondarypassword,omitempty"`
	Secure                           string        `json:"secure,omitempty"`
	SecureArgs                       string        `json:"secureargs,omitempty"`
	Send                             string        `json:"send,omitempty"`
	ServiceGroupName                 string        `json:"servicegroupname,omitempty"`
	ServiceName                      string        `json:"servicename,omitempty"`
	SIPMethod                        string        `json:"sipmethod,omitempty"`
	SIPRegURI                        string        `json:"sipreguri,omitempty"`
	SIPURI                           string        `json:"sipuri,omitempty"`
	SitePath                         string        `json:"sitepath,omitempty"`
	SNMPCommunity                    string        `json:"snmpcommunity,omitempty"`
	SNMPOID                          string        `json:"Snmpoid,omitempty"`
	SNMPThreshold                    string        `json:"snmpthreshold,omitempty"`
	SNMPVersion                      string        `json:"snmpversion,omitempty"`
	SQLQuery                         string        `json:"sqlquery,omitempty"`
	SSLProfile                       string        `json:"sslprofile,omitempty"`
	State                            string        `json:"state,omitempty"`
	StoreDB                          string        `json:"storedb,omitempty"`
	StoreFrontAcctService            string        `json:"storefrontacctservice,omitempty"`
	StoreFrontCheckBackendServices   string        `json:"storefrontcheckbackendservices,omitempty"`
	StoreName                        string        `json:"storename,omitempty"`
	SuccessRetries                   int           `json:"successretries,omitempty"`
	SupportedVendorIDs               []interface{} `json:"supportedvendorids,omitempty"`
	TOS                              string        `json:"tos,omitempty"`
	TOSID                            int           `json:"tosid,omitempty"`
	Transparent                      string        `json:"transparent,omitempty"`
	TROFSCode                        int           `json:"trofscode,omitempty"`
	TROFSString                      string        `json:"trofsstring,omitempty"`
	TypeField                        string        `json:"type,omitempty"`
	Units1                           string        `json:"units1,omitempty"`
	Units2                           string        `json:"units2,omitempty"`
	Units3                           string        `json:"units3,omitempty"`
	Units4                           string        `json:"units4,omitempty"`
	Username                         string        `json:"username,omitempty"`
	ValidateCred                     string        `json:"validatecred,omitempty"`
	VendorID                         int           `json:"vendorid,omitempty"`
	VendorSpecificAcctApplicationIDs []interface{} `json:"vendorspecificacctapplicationids,omitempty"`
	VendorSpecificAuthApplicationIDs []interface{} `json:"vendorspecificauthapplicationids,omitempty"`
	VendorSpecificVendorID           int           `json:"vendorspecificvendorid,omitempty"`
	Weight                           int           `json:"weight,omitempty"`
}

type LBProfile struct {
	ADCCookieAttributeWarningMsg  string  `json:"adccookieattributewarningmsg,omitempty"`
	ComputedADCCookieAttribute    string  `json:"computedadccookieattribute,omitempty"`
	CookiePassphrase              string  `json:"cookiepassphrase,omitempty"`
	Count                         float64 `json:"__count,omitempty"`
	DBSLB                         string  `json:"dbslb,omitempty"`
	HTTPOnlyCookieFlag            string  `json:"httponlycookieflag,omitempty"`
	LBHashAlgorithm               string  `json:"lbhashalgorithm,omitempty"`
	LBHashAlgoWinSize             int     `json:"lbhashalgowinsize,omitempty"`
	LBHashFingers                 int     `json:"lbhashfingers,omitempty"`
	LBProfileName                 string  `json:"lbprofilename,omitempty"`
	LiteralADCCookieAttribute     string  `json:"literaladccookieattribute,omitempty"`
	NextGenAPIResource            string  `json:"_nextgenapiresource,omitempty"`
	ProcessLocal                  string  `json:"processlocal,omitempty"`
	ProximityFromSelf             string  `json:"proximityfromself,omitempty"`
	StoreMQTTClientIDAndUsername  string  `json:"storemqttclientidandusername,omitempty"`
	UseEncryptedPersistenceCookie string  `json:"useencryptedpersistencecookie,omitempty"`
	UseSecuredPersistenceCookie   string  `json:"usesecuredpersistencecookie,omitempty"`
	VSvrCount                     int     `json:"vsvrcount,omitempty"`
}

type LBVServerServiceGroupBinding struct {
	Name             string `json:"name,omitempty"`
	Order            int    `json:"order,omitempty"`
	ServiceGroupName string `json:"servicegroupname,omitempty"`
	ServiceName      string `json:"servicename,omitempty"`
	Weight           int    `json:"weight,omitempty"`
}

type LBWLMBinding struct {
	LBWLMLBVServerBinding []interface{} `json:"lbwlm_lbvserver_binding,omitempty"`
	WLMName               string        `json:"wlmname,omitempty"`
}

type LBMonBindingsGSLBServiceGroupBinding struct {
	BoundServiceGroupSvrState string `json:"boundservicegroupsvrstate,omitempty"`
	MonitorName               string `json:"monitorname,omitempty"`
	MonState                  string `json:"monstate,omitempty"`
	ServiceGroupName          string `json:"servicegroupname,omitempty"`
	ServiceType               string `json:"servicetype,omitempty"`
}

type LBPolicy struct {
	Action             string   `json:"action,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	Comment            string   `json:"comment,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	LogAction          string   `json:"logaction,omitempty"`
	Name               string   `json:"name,omitempty"`
	NewName            string   `json:"newname,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	Rule               string   `json:"rule,omitempty"`
	UndefAction        string   `json:"undefaction,omitempty"`
	UndefHits          int      `json:"undefhits,omitempty"`
}

type LBMonitorServiceGroupBinding struct {
	DupState         string `json:"dup_state,omitempty"`
	DupWeight        int    `json:"dup_weight,omitempty"`
	MonitorName      string `json:"monitorname,omitempty"`
	ServiceGroupName string `json:"servicegroupname,omitempty"`
	ServiceName      string `json:"servicename,omitempty"`
	State            string `json:"state,omitempty"`
	Weight           int    `json:"weight,omitempty"`
}

type LBVServerSpilloverPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LBWLM struct {
	Count              float64 `json:"__count,omitempty"`
	IPAddress          string  `json:"ipaddress,omitempty"`
	KATimeout          int     `json:"katimeout,omitempty"`
	LBUID              string  `json:"lbuid,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Port               int     `json:"port,omitempty"`
	Secure             string  `json:"secure,omitempty"`
	State              string  `json:"state,omitempty"`
	WLMName            string  `json:"wlmname,omitempty"`
}

type LBGroupLBVServerBinding struct {
	Name        string `json:"name,omitempty"`
	VServerName string `json:"vservername,omitempty"`
}

type LBMetricTableBinding struct {
	LBMetricTableMetricBinding []interface{} `json:"lbmetrictable_metric_binding,omitempty"`
	MetricTable                string        `json:"metrictable,omitempty"`
}

type LBPersistentSessions struct {
	CNamePersParam       string  `json:"cnamepersparam,omitempty"`
	Count                float64 `json:"__count,omitempty"`
	DestIP               string  `json:"destip,omitempty"`
	DestIPv6             string  `json:"destipv6,omitempty"`
	DestPort             int     `json:"destport,omitempty"`
	Flags                bool    `json:"flags,omitempty"`
	NextGenAPIResource   string  `json:"_nextgenapiresource,omitempty"`
	NodeID               int     `json:"nodeid,omitempty"`
	PersistenceParam     string  `json:"persistenceparam,omitempty"`
	PersistenceParameter string  `json:"persistenceparameter,omitempty"`
	ReferenceCount       int     `json:"referencecount,omitempty"`
	SrcIP                string  `json:"srcip,omitempty"`
	SrcIPv6              string  `json:"srcipv6,omitempty"`
	Timeout              int     `json:"timeout,omitempty"`
	TypeField            int     `json:"type,omitempty"`
	TypeString           string  `json:"typestring,omitempty"`
	VServer              string  `json:"vserver,omitempty"`
	VServerName          string  `json:"vservername,omitempty"`
}

type LBPolicyLBVServerBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LBVServerAuditNSLogPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LBPolicyLabelBinding struct {
	LabelName                         string        `json:"labelname,omitempty"`
	LBPolicyLabelLBPolicyBinding      []interface{} `json:"lbpolicylabel_lbpolicy_binding,omitempty"`
	LBPolicyLabelPolicyBindingBinding []interface{} `json:"lbpolicylabel_policybinding_binding,omitempty"`
}

type LBMetricTableMetricBinding struct {
	Metric      string `json:"metric,omitempty"`
	MetricTable string `json:"metrictable,omitempty"`
	MetricType  string `json:"metrictype,omitempty"`
	SNMPOID     string `json:"Snmpoid,omitempty"`
}

type LBVServerDNSPolicy64Binding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LBVServerVideoOptimizationPacingPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LBPolicyLBPolicyLabelBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LBMonBindingsServiceGroupBinding struct {
	BoundServiceGroupSvrState string `json:"boundservicegroupsvrstate,omitempty"`
	MonitorName               string `json:"monitorname,omitempty"`
	MonState                  string `json:"monstate,omitempty"`
	ServiceGroupName          string `json:"servicegroupname,omitempty"`
	ServiceType               string `json:"servicetype,omitempty"`
}

type LBMonitorSSLCertKeyBinding struct {
	CA          bool   `json:"ca,omitempty"`
	CertKeyName string `json:"certkeyname,omitempty"`
	CRLCheck    string `json:"crlcheck,omitempty"`
	MonitorName string `json:"monitorname,omitempty"`
	OCSPCheck   string `json:"ocspcheck,omitempty"`
}

type LBVServerAuthorizationPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LBVServerVideoOptimizationDetectionPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LBPolicyLabel struct {
	Comment                string  `json:"comment,omitempty"`
	Count                  float64 `json:"__count,omitempty"`
	GotoPriorityExpression string  `json:"gotopriorityexpression,omitempty"`
	Hits                   int     `json:"hits,omitempty"`
	InvokeLabelName        string  `json:"invoke_labelname,omitempty"`
	LabelName              string  `json:"labelname,omitempty"`
	LabelType              string  `json:"labeltype,omitempty"`
	NewName                string  `json:"newname,omitempty"`
	NextGenAPIResource     string  `json:"_nextgenapiresource,omitempty"`
	NumPol                 int     `json:"numpol,omitempty"`
	PolicyLabelType        string  `json:"policylabeltype,omitempty"`
}

type LBPolicyBinding struct {
	LBPolicyGSLBVServerBinding   []interface{} `json:"lbpolicy_gslbvserver_binding,omitempty"`
	LBPolicyLBGlobalBinding      []interface{} `json:"lbpolicy_lbglobal_binding,omitempty"`
	LBPolicyLBPolicyLabelBinding []interface{} `json:"lbpolicy_lbpolicylabel_binding,omitempty"`
	LBPolicyLBVServerBinding     []interface{} `json:"lbpolicy_lbvserver_binding,omitempty"`
	Name                         string        `json:"name,omitempty"`
}

type LBMonitorMetricBinding struct {
	Metric          string `json:"metric,omitempty"`
	MetricUnit      string `json:"metric_unit,omitempty"`
	MetricTable     string `json:"metrictable,omitempty"`
	MetricThreshold int    `json:"metricthreshold,omitempty"`
	MetricWeight    int    `json:"metricweight,omitempty"`
	MonitorName     string `json:"monitorname,omitempty"`
}

type LBVServerLBPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LBGroupBinding struct {
	LBGroupLBVServerBinding []interface{} `json:"lbgroup_lbvserver_binding,omitempty"`
	Name                    string        `json:"name,omitempty"`
}

type LBParameter struct {
	ADCCookieAttributeWarningMsg  string   `json:"adccookieattributewarningmsg,omitempty"`
	AllowBoundSvcRemoval          string   `json:"allowboundsvcremoval,omitempty"`
	Builtin                       []string `json:"builtin,omitempty"`
	ComputedADCCookieAttribute    string   `json:"computedadccookieattribute,omitempty"`
	ConsolidatedLConn             string   `json:"consolidatedlconn,omitempty"`
	CookiePassphrase              string   `json:"cookiepassphrase,omitempty"`
	DBSTTL                        int      `json:"dbsttl,omitempty"`
	DropMQTTJumboMessage          string   `json:"dropmqttjumbomessage,omitempty"`
	Feature                       string   `json:"feature,omitempty"`
	HTTPOnlyCookieFlag            string   `json:"httponlycookieflag,omitempty"`
	LBHashAlgorithm               string   `json:"lbhashalgorithm,omitempty"`
	LBHashAlgoWinSize             int      `json:"lbhashalgowinsize,omitempty"`
	LBHashFingers                 int      `json:"lbhashfingers,omitempty"`
	LiteralADCCookieAttribute     string   `json:"literaladccookieattribute,omitempty"`
	MaxPipelineNAT                int      `json:"maxpipelinenat,omitempty"`
	MonitorConnectionClose        string   `json:"monitorconnectionclose,omitempty"`
	MonitorSkipMaxClient          string   `json:"monitorskipmaxclient,omitempty"`
	NextGenAPIResource            string   `json:"_nextgenapiresource,omitempty"`
	OverridePersistencyForOrder   string   `json:"overridepersistencyfororder,omitempty"`
	PreferDirectRoute             string   `json:"preferdirectroute,omitempty"`
	ProximityFromSelf             string   `json:"proximityfromself,omitempty"`
	RADIUSMessageAuthenticator    string   `json:"radiusmessageauthenticator,omitempty"`
	RetainServiceState            string   `json:"retainservicestate,omitempty"`
	SessionsThreshold             int      `json:"sessionsthreshold,omitempty"`
	StartupRRFactor               int      `json:"startuprrfactor,omitempty"`
	StoreMQTTClientIDAndUsername  string   `json:"storemqttclientidandusername,omitempty"`
	UndefAction                   string   `json:"undefaction,omitempty"`
	UseEncryptedPersistenceCookie string   `json:"useencryptedpersistencecookie,omitempty"`
	UsePortForHashLB              string   `json:"useportforhashlb,omitempty"`
	UseSecuredPersistenceCookie   string   `json:"usesecuredpersistencecookie,omitempty"`
	VServerSpecificMAC            string   `json:"vserverspecificmac,omitempty"`
}

type LBVServerServiceGroupMemberBinding struct {
	CookieIPPort      string `json:"cookieipport,omitempty"`
	CookieName        string `json:"cookiename,omitempty"`
	CurState          string `json:"curstate,omitempty"`
	DynamicWeight     int    `json:"dynamicweight,omitempty"`
	IPv46             string `json:"ipv46,omitempty"`
	Name              string `json:"name,omitempty"`
	Order             int    `json:"order,omitempty"`
	OrderStr          string `json:"orderstr,omitempty"`
	Port              int    `json:"port,omitempty"`
	PreferredLocation string `json:"preferredlocation,omitempty"`
	ServiceGroupName  string `json:"servicegroupname,omitempty"`
	ServiceType       string `json:"servicetype,omitempty"`
	VServerID         string `json:"vserverid,omitempty"`
	Weight            int    `json:"weight,omitempty"`
}

type LBMetricTable struct {
	Builtin            []string `json:"builtin,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Metric             string   `json:"metric,omitempty"`
	MetricTable        string   `json:"metrictable,omitempty"`
	MetricType         string   `json:"metrictype,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	SNMPOID            string   `json:"Snmpoid,omitempty"`
}
