package models

// cs configuration structs
type CSVServerTransformPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TargetLBVServer        string `json:"targetlbvserver,omitempty"`
}

type CSVServerDomainBinding struct {
	AppFlowLog    string `json:"appflowlog,omitempty"`
	BackupIP      string `json:"backupip,omitempty"`
	CookieDomain  string `json:"cookiedomain,omitempty"`
	CookieTimeout int    `json:"cookietimeout,omitempty"`
	DomainName    string `json:"domainname,omitempty"`
	Name          string `json:"name,omitempty"`
	SiteDomainTTL int    `json:"sitedomainttl,omitempty"`
	TTL           int    `json:"ttl,omitempty"`
}

type CSVServerBotPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TargetLBVServer        string `json:"targetlbvserver,omitempty"`
}

type CSParameter struct {
	Builtin            []string `json:"builtin,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	StateUpdate        string   `json:"stateupdate,omitempty"`
}

type CSVServerVPNVServerBinding struct {
	Hits    int    `json:"hits,omitempty"`
	Name    string `json:"name,omitempty"`
	VServer string `json:"vserver,omitempty"`
}

type CSVServerResponderPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TargetLBVServer        string `json:"targetlbvserver,omitempty"`
}

type CSPolicy struct {
	Action             string  `json:"action,omitempty"`
	ActivePolicy       bool    `json:"activepolicy,omitempty"`
	BoundTo            string  `json:"boundto,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Hits               int     `json:"hits,omitempty"`
	LabelName          string  `json:"labelname,omitempty"`
	LabelType          string  `json:"labeltype,omitempty"`
	LogAction          string  `json:"logaction,omitempty"`
	NewName            string  `json:"newname,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	PolicyName         string  `json:"policyname,omitempty"`
	Priority           int     `json:"priority,omitempty"`
	Rule               string  `json:"rule,omitempty"`
	VSType             int     `json:"vstype,omitempty"`
}

type CSVServerBinding struct {
	CSVServerAnalyticsProfileBinding        []any  `json:"csvserver_analyticsprofile_binding,omitempty"`
	CSVServerAppFlowPolicyBinding           []any  `json:"csvserver_appflowpolicy_binding,omitempty"`
	CSVServerAppFWPolicyBinding             []any  `json:"csvserver_appfwpolicy_binding,omitempty"`
	CSVServerAppQOEPolicyBinding            []any  `json:"csvserver_appqoepolicy_binding,omitempty"`
	CSVServerAuditNSLogPolicyBinding        []any  `json:"csvserver_auditnslogpolicy_binding,omitempty"`
	CSVServerAuditSyslogPolicyBinding       []any  `json:"csvserver_auditsyslogpolicy_binding,omitempty"`
	CSVServerAuthorizationPolicyBinding     []any  `json:"csvserver_authorizationpolicy_binding,omitempty"`
	CSVServerBotPolicyBinding               []any  `json:"csvserver_botpolicy_binding,omitempty"`
	CSVServerCachePolicyBinding             []any  `json:"csvserver_cachepolicy_binding,omitempty"`
	CSVServerCMPPolicyBinding               []any  `json:"csvserver_cmppolicy_binding,omitempty"`
	CSVServerContentInspectionPolicyBinding []any  `json:"csvserver_contentinspectionpolicy_binding,omitempty"`
	CSVServerCSPolicyBinding                []any  `json:"csvserver_cspolicy_binding,omitempty"`
	CSVServerFEOPolicyBinding               []any  `json:"csvserver_feopolicy_binding,omitempty"`
	CSVServerGSLBDomainBinding              []any  `json:"csvserver_gslbdomain_binding,omitempty"`
	CSVServerGSLBVServerBinding             []any  `json:"csvserver_gslbvserver_binding,omitempty"`
	CSVServerLBVServerBinding               []any  `json:"csvserver_lbvserver_binding,omitempty"`
	CSVServerResponderPolicyBinding         []any  `json:"csvserver_responderpolicy_binding,omitempty"`
	CSVServerRewritePolicyBinding           []any  `json:"csvserver_rewritepolicy_binding,omitempty"`
	CSVServerSpilloverPolicyBinding         []any  `json:"csvserver_spilloverpolicy_binding,omitempty"`
	CSVServerTMTrafficPolicyBinding         []any  `json:"csvserver_tmtrafficpolicy_binding,omitempty"`
	CSVServerTransformPolicyBinding         []any  `json:"csvserver_transformpolicy_binding,omitempty"`
	CSVServerVPNVServerBinding              []any  `json:"csvserver_vpnvserver_binding,omitempty"`
	Name                                    string `json:"name,omitempty"`
}

type CSPolicyCSVServerBinding struct {
	Action                 string `json:"action,omitempty"`
	BindHits               int    `json:"bindhits,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Hits                   int    `json:"hits,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type CSPolicyCSPolicyLabelBinding struct {
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Hits                   int    `json:"hits,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type CSVServerAuditNSLogPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TargetLBVServer        string `json:"targetlbvserver,omitempty"`
}

type CSVServerAppFlowPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TargetLBVServer        string `json:"targetlbvserver,omitempty"`
}

type CSVServerAuditSyslogPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TargetLBVServer        string `json:"targetlbvserver,omitempty"`
}

type CSVServerTMTrafficPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TargetLBVServer        string `json:"targetlbvserver,omitempty"`
}

type CSVServerCachePolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TargetLBVServer        string `json:"targetlbvserver,omitempty"`
}

type CSVServerAppFWPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TargetLBVServer        string `json:"targetlbvserver,omitempty"`
}

type CSVServerCSPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	CookieIPPort           string `json:"cookieipport,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Hits                   int    `json:"hits,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	PIPolicyHits           int    `json:"pipolicyhits,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Rule                   string `json:"rule,omitempty"`
	TargetLBVServer        string `json:"targetlbvserver,omitempty"`
	VServerID              string `json:"vserverid,omitempty"`
}

type CSVServerAppQOEPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TargetLBVServer        string `json:"targetlbvserver,omitempty"`
}

type CSVServerSpilloverPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TargetLBVServer        string `json:"targetlbvserver,omitempty"`
}

type CSVServerAuthorizationPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TargetLBVServer        string `json:"targetlbvserver,omitempty"`
}

type CSVServerAnalyticsProfileBinding struct {
	AnalyticsProfile string `json:"analyticsprofile,omitempty"`
	Name             string `json:"name,omitempty"`
}

type CSPolicyCRVServerBinding struct {
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Hits                   int    `json:"hits,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type CSPolicyLabelBinding struct {
	CSPolicyLabelCSPolicyBinding []any  `json:"cspolicylabel_cspolicy_binding,omitempty"`
	LabelName                    string `json:"labelname,omitempty"`
}

type CSVServerContentInspectionPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TargetLBVServer        string `json:"targetlbvserver,omitempty"`
}

type CSPolicyLabel struct {
	Count                  float64 `json:"__count,omitempty"`
	CSPolicyLabelType      string  `json:"cspolicylabeltype,omitempty"`
	GotoPriorityExpression string  `json:"gotopriorityexpression,omitempty"`
	Hits                   int     `json:"hits,omitempty"`
	InvokeLabelName        string  `json:"invoke_labelname,omitempty"`
	LabelName              string  `json:"labelname,omitempty"`
	LabelType              string  `json:"labeltype,omitempty"`
	NewName                string  `json:"newname,omitempty"`
	NextGenAPIResource     string  `json:"_nextgenapiresource,omitempty"`
	NumPol                 int     `json:"numpol,omitempty"`
	PolicyName             string  `json:"policyname,omitempty"`
	Priority               int     `json:"priority,omitempty"`
	TargetVServer          string  `json:"targetvserver,omitempty"`
}

type CSVServerGSLBVServerBinding struct {
	Hits    int    `json:"hits,omitempty"`
	Name    string `json:"name,omitempty"`
	VServer string `json:"vserver,omitempty"`
}

type CSVServerFEOPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TargetLBVServer        string `json:"targetlbvserver,omitempty"`
}

type CSVServerRewritePolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TargetLBVServer        string `json:"targetlbvserver,omitempty"`
}

type CSVServerCMPPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TargetLBVServer        string `json:"targetlbvserver,omitempty"`
}

type CSAction struct {
	Builtin            []string `json:"builtin,omitempty"`
	Comment            string   `json:"comment,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	Name               string   `json:"name,omitempty"`
	NewName            string   `json:"newname,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	ReferenceCount     int      `json:"referencecount,omitempty"`
	TargetLBVServer    string   `json:"targetlbvserver,omitempty"`
	TargetVServer      string   `json:"targetvserver,omitempty"`
	TargetVServerExpr  string   `json:"targetvserverexpr,omitempty"`
	UndefHits          int      `json:"undefhits,omitempty"`
}

type CSPolicyLabelCSPolicyBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelName        string `json:"invoke_labelname,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TargetVServer          string `json:"targetvserver,omitempty"`
}

type CSVServer struct {
	APIProfile                string  `json:"apiprofile,omitempty"`
	AppFlowLog                string  `json:"appflowlog,omitempty"`
	Authentication            string  `json:"authentication,omitempty"`
	AuthenticationHost        string  `json:"authenticationhost,omitempty"`
	Authn401                  string  `json:"authn401,omitempty"`
	AuthnProfile              string  `json:"authnprofile,omitempty"`
	AuthnVSName               string  `json:"authnvsname,omitempty"`
	BackupIP                  string  `json:"backupip,omitempty"`
	BackupPersistenceTimeout  int     `json:"backuppersistencetimeout,omitempty"`
	BackupVServer             string  `json:"backupvserver,omitempty"`
	BindPoint                 string  `json:"bindpoint,omitempty"`
	Cacheable                 string  `json:"cacheable,omitempty"`
	CacheType                 string  `json:"cachetype,omitempty"`
	CacheVServer              string  `json:"cachevserver,omitempty"`
	CaseSensitive             string  `json:"casesensitive,omitempty"`
	CltTimeout                int     `json:"clttimeout,omitempty"`
	Comment                   string  `json:"comment,omitempty"`
	CookieDomain              string  `json:"cookiedomain,omitempty"`
	CookieName                string  `json:"cookiename,omitempty"`
	CookieTimeout             int     `json:"cookietimeout,omitempty"`
	Count                     float64 `json:"__count,omitempty"`
	CurState                  string  `json:"curstate,omitempty"`
	DBProfileName             string  `json:"dbprofilename,omitempty"`
	DisablePrimaryOnDown      string  `json:"disableprimaryondown,omitempty"`
	DNSOverHTTPS              string  `json:"dnsoverhttps,omitempty"`
	DNSProfileName            string  `json:"dnsprofilename,omitempty"`
	DNSRecordType             string  `json:"dnsrecordtype,omitempty"`
	DNSVServerName            string  `json:"dnsvservername,omitempty"`
	Domain                    string  `json:"domain,omitempty"`
	DomainName                string  `json:"domainname,omitempty"`
	DownStateFlush            string  `json:"downstateflush,omitempty"`
	DTLS                      string  `json:"dtls,omitempty"`
	Gt2gb                     string  `json:"gt2gb,omitempty"`
	HomePage                  string  `json:"homepage,omitempty"`
	HTTPProfileName           string  `json:"httpprofilename,omitempty"`
	HTTPSRedirectURL          string  `json:"httpsredirecturl,omitempty"`
	ICMPVsrResponse           string  `json:"icmpvsrresponse,omitempty"`
	InsertVServerIPPort       string  `json:"insertvserveripport,omitempty"`
	IP                        string  `json:"ip,omitempty"`
	IPMask                    string  `json:"ipmask,omitempty"`
	IPPattern                 string  `json:"ippattern,omitempty"`
	IPSet                     string  `json:"ipset,omitempty"`
	IPv46                     string  `json:"ipv46,omitempty"`
	L2Conn                    string  `json:"l2conn,omitempty"`
	LBVServer                 string  `json:"lbvserver,omitempty"`
	ListenPolicy              string  `json:"listenpolicy,omitempty"`
	ListenPriority            int     `json:"listenpriority,omitempty"`
	MSSQLServerVersion        string  `json:"mssqlserverversion,omitempty"`
	MySQLCharacterSet         int     `json:"mysqlcharacterset,omitempty"`
	MySQLProtocolVersion      int     `json:"mysqlprotocolversion,omitempty"`
	MySQLServerCapabilities   int     `json:"mysqlservercapabilities,omitempty"`
	MySQLServerVersion        string  `json:"mysqlserverversion,omitempty"`
	Name                      string  `json:"name,omitempty"`
	NetProfile                string  `json:"netprofile,omitempty"`
	NewName                   string  `json:"newname,omitempty"`
	NextGenAPIResource        string  `json:"_nextgenapiresource,omitempty"`
	NgName                    string  `json:"ngname,omitempty"`
	NoDefaultBindings         string  `json:"nodefaultbindings,omitempty"`
	OracleServerVersion       string  `json:"oracleserverversion,omitempty"`
	PersistenceBackup         string  `json:"persistencebackup,omitempty"`
	PersistenceID             int     `json:"persistenceid,omitempty"`
	PersistenceType           string  `json:"persistencetype,omitempty"`
	PersistMask               string  `json:"persistmask,omitempty"`
	Port                      int     `json:"port,omitempty"`
	Precedence                string  `json:"precedence,omitempty"`
	ProbePort                 int     `json:"probeport,omitempty"`
	ProbeProtocol             string  `json:"probeprotocol,omitempty"`
	ProbeSuccessResponseCode  string  `json:"probesuccessresponsecode,omitempty"`
	Push                      string  `json:"push,omitempty"`
	PushLabel                 string  `json:"pushlabel,omitempty"`
	PushMultiClients          string  `json:"pushmulticlients,omitempty"`
	PushVServer               string  `json:"pushvserver,omitempty"`
	QUICProfileName           string  `json:"quicprofilename,omitempty"`
	Range                     int     `json:"range,omitempty"`
	Redirect                  string  `json:"redirect,omitempty"`
	RedirectFromPort          int     `json:"redirectfromport,omitempty"`
	RedirectPortRewrite       string  `json:"redirectportrewrite,omitempty"`
	RedirectURL               string  `json:"redirecturl,omitempty"`
	RHIState                  string  `json:"rhistate,omitempty"`
	RTSPNAT                   string  `json:"rtspnat,omitempty"`
	RuleType                  int     `json:"ruletype,omitempty"`
	ServiceName               string  `json:"servicename,omitempty"`
	ServiceType               string  `json:"servicetype,omitempty"`
	SiteDomainTTL             int     `json:"sitedomainttl,omitempty"`
	SOBackupAction            string  `json:"sobackupaction,omitempty"`
	SOMethod                  string  `json:"somethod,omitempty"`
	SOPersistence             string  `json:"sopersistence,omitempty"`
	SOPersistenceTimeout      int     `json:"sopersistencetimeout,omitempty"`
	SOThreshold               int     `json:"sothreshold,omitempty"`
	State                     string  `json:"state,omitempty"`
	StateChangeTimeMsec       int     `json:"statechangetimemsec,omitempty"`
	StateChangeTimeSec        string  `json:"statechangetimesec,omitempty"`
	StateUpdate               string  `json:"stateupdate,omitempty"`
	Status                    int     `json:"status,omitempty"`
	TargetLBVServer           string  `json:"targetlbvserver,omitempty"`
	TargetType                string  `json:"targettype,omitempty"`
	TargetVServer             string  `json:"targetvserver,omitempty"`
	TCPProbePort              int     `json:"tcpprobeport,omitempty"`
	TCPProfileName            string  `json:"tcpprofilename,omitempty"`
	TD                        int     `json:"td,omitempty"`
	TicksSinceLastStateChange int     `json:"tickssincelaststatechange,omitempty"`
	Timeout                   int     `json:"timeout,omitempty"`
	TTL                       int     `json:"ttl,omitempty"`
	TypeField                 string  `json:"type,omitempty"`
	URL                       string  `json:"url,omitempty"`
	V6PersistMaskLen          int     `json:"v6persistmasklen,omitempty"`
	Value                     string  `json:"value,omitempty"`
	Version                   int     `json:"version,omitempty"`
	VIPHeader                 string  `json:"vipheader,omitempty"`
	Weight                    int     `json:"weight,omitempty"`
}

type CSPolicyBinding struct {
	CSPolicyCRVServerBinding     []any  `json:"cspolicy_crvserver_binding,omitempty"`
	CSPolicyCSPolicyLabelBinding []any  `json:"cspolicy_cspolicylabel_binding,omitempty"`
	CSPolicyCSVServerBinding     []any  `json:"cspolicy_csvserver_binding,omitempty"`
	PolicyName                   string `json:"policyname,omitempty"`
}

type CSVServerLBVServerBinding struct {
	CookieIPPort  string `json:"cookieipport,omitempty"`
	Hits          int    `json:"hits,omitempty"`
	LBVServer     string `json:"lbvserver,omitempty"`
	Name          string `json:"name,omitempty"`
	TargetVServer string `json:"targetvserver,omitempty"`
	VServerID     string `json:"vserverid,omitempty"`
}
