package models

// audit configuration structs
type AuditSyslogAction struct {
	ACL                  string   `json:"acl,omitempty"`
	ALG                  string   `json:"alg,omitempty"`
	AppFlowExport        string   `json:"appflowexport,omitempty"`
	Builtin              []string `json:"builtin,omitempty"`
	ContentInspectionLog string   `json:"contentinspectionlog,omitempty"`
	Count                float64  `json:"__count,omitempty"`
	DateFormat           string   `json:"dateformat,omitempty"`
	DNS                  string   `json:"dns,omitempty"`
	DomainResolveNow     bool     `json:"domainresolvenow,omitempty"`
	DomainResolveRetry   int      `json:"domainresolveretry,omitempty"`
	Feature              string   `json:"feature,omitempty"`
	HTTPAuthToken        string   `json:"httpauthtoken,omitempty"`
	HTTPEndpointURL      string   `json:"httpendpointurl,omitempty"`
	IP                   string   `json:"ip,omitempty"`
	LBVServerName        string   `json:"lbvservername,omitempty"`
	LogFacility          string   `json:"logfacility,omitempty"`
	LogLevel             []string `json:"loglevel,omitempty"`
	LSN                  string   `json:"lsn,omitempty"`
	ManagementLog        []string `json:"managementlog,omitempty"`
	MaxLogDataSizeToHold int      `json:"maxlogdatasizetohold,omitempty"`
	MgmtLogLevel         []string `json:"mgmtloglevel,omitempty"`
	Name                 string   `json:"name,omitempty"`
	NetProfile           string   `json:"netprofile,omitempty"`
	NextGenAPIResource   string   `json:"_nextgenapiresource,omitempty"`
	ProtocolViolations   string   `json:"protocolviolations,omitempty"`
	ServerDomainName     string   `json:"serverdomainname,omitempty"`
	ServerIP             string   `json:"serverip,omitempty"`
	ServerPort           int      `json:"serverport,omitempty"`
	SSLInterception      string   `json:"sslinterception,omitempty"`
	StreamAnalytics      string   `json:"streamanalytics,omitempty"`
	SubscriberLog        string   `json:"subscriberlog,omitempty"`
	SyslogCompliance     string   `json:"syslogcompliance,omitempty"`
	TCP                  string   `json:"tcp,omitempty"`
	TCPProfileName       string   `json:"tcpprofilename,omitempty"`
	TimeZone             string   `json:"timezone,omitempty"`
	Transport            string   `json:"transport,omitempty"`
	URLFiltering         string   `json:"urlfiltering,omitempty"`
	UserDefinedAuditLog  string   `json:"userdefinedauditlog,omitempty"`
}

type AuditSyslogParams struct {
	ACL                  string   `json:"acl,omitempty"`
	ALG                  string   `json:"alg,omitempty"`
	AppFlowExport        string   `json:"appflowexport,omitempty"`
	Builtin              []string `json:"builtin,omitempty"`
	ContentInspectionLog string   `json:"contentinspectionlog,omitempty"`
	DateFormat           string   `json:"dateformat,omitempty"`
	DNS                  string   `json:"dns,omitempty"`
	Feature              string   `json:"feature,omitempty"`
	LogFacility          string   `json:"logfacility,omitempty"`
	LogLevel             []string `json:"loglevel,omitempty"`
	LSN                  string   `json:"lsn,omitempty"`
	NextGenAPIResource   string   `json:"_nextgenapiresource,omitempty"`
	ProtocolViolations   string   `json:"protocolviolations,omitempty"`
	ServerIP             string   `json:"serverip,omitempty"`
	ServerPort           int      `json:"serverport,omitempty"`
	SSLInterception      string   `json:"sslinterception,omitempty"`
	StreamAnalytics      string   `json:"streamanalytics,omitempty"`
	SubscriberLog        string   `json:"subscriberlog,omitempty"`
	TCP                  string   `json:"tcp,omitempty"`
	TimeZone             string   `json:"timezone,omitempty"`
	URLFiltering         string   `json:"urlfiltering,omitempty"`
	UserDefinedAuditLog  string   `json:"userdefinedauditlog,omitempty"`
}

type AuditSyslogPolicyVPNVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditNSLogPolicyAuditNSLogGlobalBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditSyslogPolicyAAAGroupBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditSyslogPolicyRNATGlobalBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditNSLogPolicyVPNVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditNSLogPolicyAuthenticationVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditNSLogPolicyVPNGlobalBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditNSLogPolicyTMGlobalBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditNSLogPolicyLBVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditNSLogAction struct {
	ACL                  string   `json:"acl,omitempty"`
	ALG                  string   `json:"alg,omitempty"`
	AppFlowExport        string   `json:"appflowexport,omitempty"`
	Builtin              []string `json:"builtin,omitempty"`
	ContentInspectionLog string   `json:"contentinspectionlog,omitempty"`
	Count                float64  `json:"__count,omitempty"`
	DateFormat           string   `json:"dateformat,omitempty"`
	DomainResolveNow     bool     `json:"domainresolvenow,omitempty"`
	DomainResolveRetry   int      `json:"domainresolveretry,omitempty"`
	Feature              string   `json:"feature,omitempty"`
	IP                   string   `json:"ip,omitempty"`
	LogFacility          string   `json:"logfacility,omitempty"`
	LogLevel             []string `json:"loglevel,omitempty"`
	LSN                  string   `json:"lsn,omitempty"`
	Name                 string   `json:"name,omitempty"`
	NextGenAPIResource   string   `json:"_nextgenapiresource,omitempty"`
	ProtocolViolations   string   `json:"protocolviolations,omitempty"`
	ServerDomainName     string   `json:"serverdomainname,omitempty"`
	ServerIP             string   `json:"serverip,omitempty"`
	ServerPort           int      `json:"serverport,omitempty"`
	SSLInterception      string   `json:"sslinterception,omitempty"`
	SubscriberLog        string   `json:"subscriberlog,omitempty"`
	TCP                  string   `json:"tcp,omitempty"`
	TimeZone             string   `json:"timezone,omitempty"`
	URLFiltering         string   `json:"urlfiltering,omitempty"`
	UserDefinedAuditLog  string   `json:"userdefinedauditlog,omitempty"`
}

type AuditSyslogPolicy struct {
	Action             string   `json:"action,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	ExpressionType     string   `json:"expressiontype,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Name               string   `json:"name,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	Rule               string   `json:"rule,omitempty"`
}

type AuditMessageAction struct {
	BypassSafetyCheck  string  `json:"bypasssafetycheck,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Hits               int     `json:"hits,omitempty"`
	LogLevel           string  `json:"loglevel,omitempty"`
	LogLevel1          string  `json:"loglevel1,omitempty"`
	LogToNewNSLog      string  `json:"logtonewnslog,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	ReferenceCount     int     `json:"referencecount,omitempty"`
	StringBuilderExpr  string  `json:"stringbuilderexpr,omitempty"`
	UndefHits          int     `json:"undefhits,omitempty"`
}

type AuditSyslogPolicyVPNGlobalBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditMessages struct {
	Count              float64  `json:"__count,omitempty"`
	LogLevel           []string `json:"loglevel,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	NumOfMesgs         int      `json:"numofmesgs,omitempty"`
	Value              string   `json:"value,omitempty"`
}

type AuditNSLogPolicy struct {
	Action             string   `json:"action,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	ExpressionType     string   `json:"expressiontype,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Name               string   `json:"name,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	Rule               string   `json:"rule,omitempty"`
}

type AuditNSLogGlobalAuditNSLogPolicyBinding struct {
	Builtin        []string `json:"builtin,omitempty"`
	GlobalBindType string   `json:"globalbindtype,omitempty"`
	NumPol         int      `json:"numpol,omitempty"`
	PolicyName     string   `json:"policyname,omitempty"`
	Priority       int      `json:"priority,omitempty"`
}

type AuditNSLogPolicySystemGlobalBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditSyslogPolicySystemGlobalBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditNSLogPolicyAppFWGlobalBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditSyslogPolicyTMGlobalBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditNSLogPolicyBinding struct {
	AuditNSLogPolicyAAAGroupBinding              []interface{} `json:"auditnslogpolicy_aaagroup_binding,omitempty"`
	AuditNSLogPolicyAAAUserBinding               []interface{} `json:"auditnslogpolicy_aaauser_binding,omitempty"`
	AuditNSLogPolicyAppFWGlobalBinding           []interface{} `json:"auditnslogpolicy_appfwglobal_binding,omitempty"`
	AuditNSLogPolicyAuditNSLogGlobalBinding      []interface{} `json:"auditnslogpolicy_auditnslogglobal_binding,omitempty"`
	AuditNSLogPolicyAuthenticationVServerBinding []interface{} `json:"auditnslogpolicy_authenticationvserver_binding,omitempty"`
	AuditNSLogPolicyCSVServerBinding             []interface{} `json:"auditnslogpolicy_csvserver_binding,omitempty"`
	AuditNSLogPolicyLBVServerBinding             []interface{} `json:"auditnslogpolicy_lbvserver_binding,omitempty"`
	AuditNSLogPolicySystemGlobalBinding          []interface{} `json:"auditnslogpolicy_systemglobal_binding,omitempty"`
	AuditNSLogPolicyTMGlobalBinding              []interface{} `json:"auditnslogpolicy_tmglobal_binding,omitempty"`
	AuditNSLogPolicyVPNGlobalBinding             []interface{} `json:"auditnslogpolicy_vpnglobal_binding,omitempty"`
	AuditNSLogPolicyVPNVServerBinding            []interface{} `json:"auditnslogpolicy_vpnvserver_binding,omitempty"`
	Name                                         string        `json:"name,omitempty"`
}

type AuditSyslogPolicyBinding struct {
	AuditSyslogPolicyAAAGroupBinding              []interface{} `json:"auditsyslogpolicy_aaagroup_binding,omitempty"`
	AuditSyslogPolicyAAAUserBinding               []interface{} `json:"auditsyslogpolicy_aaauser_binding,omitempty"`
	AuditSyslogPolicyAuditSyslogGlobalBinding     []interface{} `json:"auditsyslogpolicy_auditsyslogglobal_binding,omitempty"`
	AuditSyslogPolicyAuthenticationVServerBinding []interface{} `json:"auditsyslogpolicy_authenticationvserver_binding,omitempty"`
	AuditSyslogPolicyCSVServerBinding             []interface{} `json:"auditsyslogpolicy_csvserver_binding,omitempty"`
	AuditSyslogPolicyLBVServerBinding             []interface{} `json:"auditsyslogpolicy_lbvserver_binding,omitempty"`
	AuditSyslogPolicyRNATGlobalBinding            []interface{} `json:"auditsyslogpolicy_rnatglobal_binding,omitempty"`
	AuditSyslogPolicySystemGlobalBinding          []interface{} `json:"auditsyslogpolicy_systemglobal_binding,omitempty"`
	AuditSyslogPolicyTMGlobalBinding              []interface{} `json:"auditsyslogpolicy_tmglobal_binding,omitempty"`
	AuditSyslogPolicyVPNGlobalBinding             []interface{} `json:"auditsyslogpolicy_vpnglobal_binding,omitempty"`
	AuditSyslogPolicyVPNVServerBinding            []interface{} `json:"auditsyslogpolicy_vpnvserver_binding,omitempty"`
	Name                                          string        `json:"name,omitempty"`
}

type AuditSyslogGlobalAuditSyslogPolicyBinding struct {
	Builtin        []string `json:"builtin,omitempty"`
	Feature        string   `json:"feature,omitempty"`
	GlobalBindType string   `json:"globalbindtype,omitempty"`
	NumPol         int      `json:"numpol,omitempty"`
	PolicyName     string   `json:"policyname,omitempty"`
	Priority       int      `json:"priority,omitempty"`
}

type AuditSyslogGlobalBinding struct {
	AuditSyslogGlobalAuditSyslogPolicyBinding []interface{} `json:"auditsyslogglobal_auditsyslogpolicy_binding,omitempty"`
}

type AuditSyslogPolicyCSVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditNSLogPolicyAAAGroupBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditSyslogPolicyLBVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditNSLogPolicyAAAUserBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditSyslogPolicyAuditSyslogGlobalBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditNSLogParams struct {
	ACL                  string   `json:"acl,omitempty"`
	ALG                  string   `json:"alg,omitempty"`
	AppFlowExport        string   `json:"appflowexport,omitempty"`
	Builtin              []string `json:"builtin,omitempty"`
	ContentInspectionLog string   `json:"contentinspectionlog,omitempty"`
	DateFormat           string   `json:"dateformat,omitempty"`
	Feature              string   `json:"feature,omitempty"`
	LogFacility          string   `json:"logfacility,omitempty"`
	LogLevel             []string `json:"loglevel,omitempty"`
	LSN                  string   `json:"lsn,omitempty"`
	NextGenAPIResource   string   `json:"_nextgenapiresource,omitempty"`
	ProtocolViolations   string   `json:"protocolviolations,omitempty"`
	ServerIP             string   `json:"serverip,omitempty"`
	ServerPort           int      `json:"serverport,omitempty"`
	SSLInterception      string   `json:"sslinterception,omitempty"`
	SubscriberLog        string   `json:"subscriberlog,omitempty"`
	TCP                  string   `json:"tcp,omitempty"`
	TimeZone             string   `json:"timezone,omitempty"`
	URLFiltering         string   `json:"urlfiltering,omitempty"`
	UserDefinedAuditLog  string   `json:"userdefinedauditlog,omitempty"`
}

type AuditNSLogPolicyCSVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditSyslogPolicyAAAUserBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditNSLogGlobalBinding struct {
	AuditNSLogGlobalAuditNSLogPolicyBinding []interface{} `json:"auditnslogglobal_auditnslogpolicy_binding,omitempty"`
}

type AuditSyslogPolicyAuthenticationVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}
