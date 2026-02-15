package models

// audit configuration structs
type Auditsyslogaction struct {
	Acl                  string   `json:"acl,omitempty"`
	Alg                  string   `json:"alg,omitempty"`
	Appflowexport        string   `json:"appflowexport,omitempty"`
	Builtin              []string `json:"builtin,omitempty"`
	Contentinspectionlog string   `json:"contentinspectionlog,omitempty"`
	Count                float64  `json:"__count,omitempty"`
	Dateformat           string   `json:"dateformat,omitempty"`
	Dns                  string   `json:"dns,omitempty"`
	Domainresolvenow     bool     `json:"domainresolvenow,omitempty"`
	Domainresolveretry   int      `json:"domainresolveretry,omitempty"`
	Feature              string   `json:"feature,omitempty"`
	Httpauthtoken        string   `json:"httpauthtoken,omitempty"`
	Httpendpointurl      string   `json:"httpendpointurl,omitempty"`
	Ip                   string   `json:"ip,omitempty"`
	Lbvservername        string   `json:"lbvservername,omitempty"`
	Logfacility          string   `json:"logfacility,omitempty"`
	Loglevel             []string `json:"loglevel,omitempty"`
	Lsn                  string   `json:"lsn,omitempty"`
	Managementlog        []string `json:"managementlog,omitempty"`
	Maxlogdatasizetohold int      `json:"maxlogdatasizetohold,omitempty"`
	Mgmtloglevel         []string `json:"mgmtloglevel,omitempty"`
	Name                 string   `json:"name,omitempty"`
	Netprofile           string   `json:"netprofile,omitempty"`
	Nextgenapiresource   string   `json:"_nextgenapiresource,omitempty"`
	Protocolviolations   string   `json:"protocolviolations,omitempty"`
	Serverdomainname     string   `json:"serverdomainname,omitempty"`
	Serverip             string   `json:"serverip,omitempty"`
	Serverport           int      `json:"serverport,omitempty"`
	Sslinterception      string   `json:"sslinterception,omitempty"`
	Streamanalytics      string   `json:"streamanalytics,omitempty"`
	Subscriberlog        string   `json:"subscriberlog,omitempty"`
	Syslogcompliance     string   `json:"syslogcompliance,omitempty"`
	Tcp                  string   `json:"tcp,omitempty"`
	Tcpprofilename       string   `json:"tcpprofilename,omitempty"`
	Timezone             string   `json:"timezone,omitempty"`
	Transport            string   `json:"transport,omitempty"`
	Urlfiltering         string   `json:"urlfiltering,omitempty"`
	Userdefinedauditlog  string   `json:"userdefinedauditlog,omitempty"`
}

type Auditsyslogparams struct {
	Acl                  string   `json:"acl,omitempty"`
	Alg                  string   `json:"alg,omitempty"`
	Appflowexport        string   `json:"appflowexport,omitempty"`
	Builtin              []string `json:"builtin,omitempty"`
	Contentinspectionlog string   `json:"contentinspectionlog,omitempty"`
	Dateformat           string   `json:"dateformat,omitempty"`
	Dns                  string   `json:"dns,omitempty"`
	Feature              string   `json:"feature,omitempty"`
	Logfacility          string   `json:"logfacility,omitempty"`
	Loglevel             []string `json:"loglevel,omitempty"`
	Lsn                  string   `json:"lsn,omitempty"`
	Nextgenapiresource   string   `json:"_nextgenapiresource,omitempty"`
	Protocolviolations   string   `json:"protocolviolations,omitempty"`
	Serverip             string   `json:"serverip,omitempty"`
	Serverport           int      `json:"serverport,omitempty"`
	Sslinterception      string   `json:"sslinterception,omitempty"`
	Streamanalytics      string   `json:"streamanalytics,omitempty"`
	Subscriberlog        string   `json:"subscriberlog,omitempty"`
	Tcp                  string   `json:"tcp,omitempty"`
	Timezone             string   `json:"timezone,omitempty"`
	Urlfiltering         string   `json:"urlfiltering,omitempty"`
	Userdefinedauditlog  string   `json:"userdefinedauditlog,omitempty"`
}

type AuditsyslogpolicyVpnvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditnslogpolicyAuditnslogglobalBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditsyslogpolicyAaagroupBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditsyslogpolicyRnatglobalBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditnslogpolicyVpnvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditnslogpolicyAuthenticationvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditnslogpolicyVpnglobalBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditnslogpolicyTmglobalBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditnslogpolicyLbvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type Auditnslogaction struct {
	Acl                  string   `json:"acl,omitempty"`
	Alg                  string   `json:"alg,omitempty"`
	Appflowexport        string   `json:"appflowexport,omitempty"`
	Builtin              []string `json:"builtin,omitempty"`
	Contentinspectionlog string   `json:"contentinspectionlog,omitempty"`
	Count                float64  `json:"__count,omitempty"`
	Dateformat           string   `json:"dateformat,omitempty"`
	Domainresolvenow     bool     `json:"domainresolvenow,omitempty"`
	Domainresolveretry   int      `json:"domainresolveretry,omitempty"`
	Feature              string   `json:"feature,omitempty"`
	Ip                   string   `json:"ip,omitempty"`
	Logfacility          string   `json:"logfacility,omitempty"`
	Loglevel             []string `json:"loglevel,omitempty"`
	Lsn                  string   `json:"lsn,omitempty"`
	Name                 string   `json:"name,omitempty"`
	Nextgenapiresource   string   `json:"_nextgenapiresource,omitempty"`
	Protocolviolations   string   `json:"protocolviolations,omitempty"`
	Serverdomainname     string   `json:"serverdomainname,omitempty"`
	Serverip             string   `json:"serverip,omitempty"`
	Serverport           int      `json:"serverport,omitempty"`
	Sslinterception      string   `json:"sslinterception,omitempty"`
	Subscriberlog        string   `json:"subscriberlog,omitempty"`
	Tcp                  string   `json:"tcp,omitempty"`
	Timezone             string   `json:"timezone,omitempty"`
	Urlfiltering         string   `json:"urlfiltering,omitempty"`
	Userdefinedauditlog  string   `json:"userdefinedauditlog,omitempty"`
}

type Auditsyslogpolicy struct {
	Action             string   `json:"action,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Expressiontype     string   `json:"expressiontype,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Name               string   `json:"name,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Rule               string   `json:"rule,omitempty"`
}

type Auditmessageaction struct {
	Bypasssafetycheck  string  `json:"bypasssafetycheck,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Hits               int     `json:"hits,omitempty"`
	Loglevel           string  `json:"loglevel,omitempty"`
	Loglevel1          string  `json:"loglevel1,omitempty"`
	Logtonewnslog      string  `json:"logtonewnslog,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Referencecount     int     `json:"referencecount,omitempty"`
	Stringbuilderexpr  string  `json:"stringbuilderexpr,omitempty"`
	Undefhits          int     `json:"undefhits,omitempty"`
}

type AuditsyslogpolicyVpnglobalBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type Auditmessages struct {
	Count              float64  `json:"__count,omitempty"`
	Loglevel           []string `json:"loglevel,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Numofmesgs         int      `json:"numofmesgs,omitempty"`
	Value              string   `json:"value,omitempty"`
}

type Auditnslogpolicy struct {
	Action             string   `json:"action,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Expressiontype     string   `json:"expressiontype,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Name               string   `json:"name,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Rule               string   `json:"rule,omitempty"`
}

type AuditnslogglobalAuditnslogpolicyBinding struct {
	Builtin        []string `json:"builtin,omitempty"`
	Globalbindtype string   `json:"globalbindtype,omitempty"`
	Numpol         int      `json:"numpol,omitempty"`
	Policyname     string   `json:"policyname,omitempty"`
	Priority       int      `json:"priority,omitempty"`
}

type AuditnslogpolicySystemglobalBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditsyslogpolicySystemglobalBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditnslogpolicyAppfwglobalBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditsyslogpolicyTmglobalBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditnslogpolicyBinding struct {
	AuditnslogpolicyAaagroupBinding              []interface{} `json:"auditnslogpolicy_aaagroup_binding,omitempty"`
	AuditnslogpolicyAaauserBinding               []interface{} `json:"auditnslogpolicy_aaauser_binding,omitempty"`
	AuditnslogpolicyAppfwglobalBinding           []interface{} `json:"auditnslogpolicy_appfwglobal_binding,omitempty"`
	AuditnslogpolicyAuditnslogglobalBinding      []interface{} `json:"auditnslogpolicy_auditnslogglobal_binding,omitempty"`
	AuditnslogpolicyAuthenticationvserverBinding []interface{} `json:"auditnslogpolicy_authenticationvserver_binding,omitempty"`
	AuditnslogpolicyCsvserverBinding             []interface{} `json:"auditnslogpolicy_csvserver_binding,omitempty"`
	AuditnslogpolicyLbvserverBinding             []interface{} `json:"auditnslogpolicy_lbvserver_binding,omitempty"`
	AuditnslogpolicySystemglobalBinding          []interface{} `json:"auditnslogpolicy_systemglobal_binding,omitempty"`
	AuditnslogpolicyTmglobalBinding              []interface{} `json:"auditnslogpolicy_tmglobal_binding,omitempty"`
	AuditnslogpolicyVpnglobalBinding             []interface{} `json:"auditnslogpolicy_vpnglobal_binding,omitempty"`
	AuditnslogpolicyVpnvserverBinding            []interface{} `json:"auditnslogpolicy_vpnvserver_binding,omitempty"`
	Name                                         string        `json:"name,omitempty"`
}

type AuditsyslogpolicyBinding struct {
	AuditsyslogpolicyAaagroupBinding              []interface{} `json:"auditsyslogpolicy_aaagroup_binding,omitempty"`
	AuditsyslogpolicyAaauserBinding               []interface{} `json:"auditsyslogpolicy_aaauser_binding,omitempty"`
	AuditsyslogpolicyAuditsyslogglobalBinding     []interface{} `json:"auditsyslogpolicy_auditsyslogglobal_binding,omitempty"`
	AuditsyslogpolicyAuthenticationvserverBinding []interface{} `json:"auditsyslogpolicy_authenticationvserver_binding,omitempty"`
	AuditsyslogpolicyCsvserverBinding             []interface{} `json:"auditsyslogpolicy_csvserver_binding,omitempty"`
	AuditsyslogpolicyLbvserverBinding             []interface{} `json:"auditsyslogpolicy_lbvserver_binding,omitempty"`
	AuditsyslogpolicyRnatglobalBinding            []interface{} `json:"auditsyslogpolicy_rnatglobal_binding,omitempty"`
	AuditsyslogpolicySystemglobalBinding          []interface{} `json:"auditsyslogpolicy_systemglobal_binding,omitempty"`
	AuditsyslogpolicyTmglobalBinding              []interface{} `json:"auditsyslogpolicy_tmglobal_binding,omitempty"`
	AuditsyslogpolicyVpnglobalBinding             []interface{} `json:"auditsyslogpolicy_vpnglobal_binding,omitempty"`
	AuditsyslogpolicyVpnvserverBinding            []interface{} `json:"auditsyslogpolicy_vpnvserver_binding,omitempty"`
	Name                                          string        `json:"name,omitempty"`
}

type AuditsyslogglobalAuditsyslogpolicyBinding struct {
	Builtin        []string `json:"builtin,omitempty"`
	Feature        string   `json:"feature,omitempty"`
	Globalbindtype string   `json:"globalbindtype,omitempty"`
	Numpol         int      `json:"numpol,omitempty"`
	Policyname     string   `json:"policyname,omitempty"`
	Priority       int      `json:"priority,omitempty"`
}

type AuditsyslogglobalBinding struct {
	AuditsyslogglobalAuditsyslogpolicyBinding []interface{} `json:"auditsyslogglobal_auditsyslogpolicy_binding,omitempty"`
}

type AuditsyslogpolicyCsvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditnslogpolicyAaagroupBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditsyslogpolicyLbvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditnslogpolicyAaauserBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditsyslogpolicyAuditsyslogglobalBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type Auditnslogparams struct {
	Acl                  string   `json:"acl,omitempty"`
	Alg                  string   `json:"alg,omitempty"`
	Appflowexport        string   `json:"appflowexport,omitempty"`
	Builtin              []string `json:"builtin,omitempty"`
	Contentinspectionlog string   `json:"contentinspectionlog,omitempty"`
	Dateformat           string   `json:"dateformat,omitempty"`
	Feature              string   `json:"feature,omitempty"`
	Logfacility          string   `json:"logfacility,omitempty"`
	Loglevel             []string `json:"loglevel,omitempty"`
	Lsn                  string   `json:"lsn,omitempty"`
	Nextgenapiresource   string   `json:"_nextgenapiresource,omitempty"`
	Protocolviolations   string   `json:"protocolviolations,omitempty"`
	Serverip             string   `json:"serverip,omitempty"`
	Serverport           int      `json:"serverport,omitempty"`
	Sslinterception      string   `json:"sslinterception,omitempty"`
	Subscriberlog        string   `json:"subscriberlog,omitempty"`
	Tcp                  string   `json:"tcp,omitempty"`
	Timezone             string   `json:"timezone,omitempty"`
	Urlfiltering         string   `json:"urlfiltering,omitempty"`
	Userdefinedauditlog  string   `json:"userdefinedauditlog,omitempty"`
}

type AuditnslogpolicyCsvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditsyslogpolicyAaauserBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuditnslogglobalBinding struct {
	AuditnslogglobalAuditnslogpolicyBinding []interface{} `json:"auditnslogglobal_auditnslogpolicy_binding,omitempty"`
}

type AuditsyslogpolicyAuthenticationvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}
