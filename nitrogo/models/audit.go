// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Auditnslogpolicygroupbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Auditnslogpolicyuserbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Auditnslogpolicyvpnvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Auditsyslogpolicyauthenticationvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Auditsyslogpolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Auditsyslogpolicycsvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Auditsyslogpolicyglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Auditsyslogpolicylbvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Auditnslogpolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Auditmessages struct {
	Loglevel           []string `json:"loglevel,omitempty"`
	Numofmesgs         int      `json:"numofmesgs,omitempty"`
	Value              string   `json:"value,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
}

type Auditnslogglobalbinding struct {
}

type Auditnslogpolicytmglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Auditsyslogglobalauditsyslogpolicybinding struct {
	Policyname     string   `json:"policyname,omitempty"`
	Priority       int      `json:"priority,omitempty"`
	Numpol         int      `json:"numpol,omitempty"`
	Globalbindtype string   `json:"globalbindtype,omitempty"`
	Builtin        []string `json:"builtin,omitempty"`
	Feature        string   `json:"feature,omitempty"`
}

type Auditsyslogpolicyaaagroupbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Auditsyslogpolicyauditsyslogglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Auditsyslogpolicysyslogglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Auditnslogglobalnslogpolicybinding struct {
	Policyname     string   `json:"policyname,omitempty"`
	Priority       uint32   `json:"priority,omitempty"`
	Numpol         uint32   `json:"numpol,omitempty"`
	Globalbindtype string   `json:"globalbindtype,omitempty"`
	Builtin        []string `json:"builtin,omitempty"`
}

type Auditnslogpolicyvpnglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Auditsyslogglobalbinding struct {
}

type Auditsyslogpolicyaaauserbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Auditsyslogpolicygroupbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Auditsyslogpolicysystemglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Auditsyslogpolicytmglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Auditnslogpolicyaaagroupbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Auditsyslogaction struct {
	Name                 string   `json:"name,omitempty"`
	Serverip             string   `json:"serverip,omitempty"`
	Serverdomainname     string   `json:"serverdomainname,omitempty"`
	Domainresolveretry   int      `json:"domainresolveretry,omitempty"`
	Lbvservername        string   `json:"lbvservername,omitempty"`
	Serverport           int      `json:"serverport,omitempty"`
	Loglevel             []string `json:"loglevel,omitempty"`
	Managementlog        []string `json:"managementlog,omitempty"`
	Mgmtloglevel         []string `json:"mgmtloglevel,omitempty"`
	Syslogcompliance     string   `json:"syslogcompliance,omitempty"`
	Dateformat           string   `json:"dateformat,omitempty"`
	Logfacility          string   `json:"logfacility,omitempty"`
	Tcp                  string   `json:"tcp,omitempty"`
	Acl                  string   `json:"acl,omitempty"`
	Timezone             string   `json:"timezone,omitempty"`
	Userdefinedauditlog  string   `json:"userdefinedauditlog,omitempty"`
	Appflowexport        string   `json:"appflowexport,omitempty"`
	Lsn                  string   `json:"lsn,omitempty"`
	Alg                  string   `json:"alg,omitempty"`
	Subscriberlog        string   `json:"subscriberlog,omitempty"`
	Transport            string   `json:"transport,omitempty"`
	Httpauthtoken        string   `json:"httpauthtoken,omitempty"`
	Httpendpointurl      string   `json:"httpendpointurl,omitempty"`
	Tcpprofilename       string   `json:"tcpprofilename,omitempty"`
	Maxlogdatasizetohold int      `json:"maxlogdatasizetohold,omitempty"`
	Dns                  string   `json:"dns,omitempty"`
	Contentinspectionlog string   `json:"contentinspectionlog,omitempty"`
	Netprofile           string   `json:"netprofile,omitempty"`
	Sslinterception      string   `json:"sslinterception,omitempty"`
	Urlfiltering         string   `json:"urlfiltering,omitempty"`
	Streamanalytics      string   `json:"streamanalytics,omitempty"`
	Protocolviolations   string   `json:"protocolviolations,omitempty"`
	Domainresolvenow     bool     `json:"domainresolvenow,omitempty"`
	Ip                   string   `json:"ip,omitempty"`
	Builtin              string   `json:"builtin,omitempty"`
	Feature              string   `json:"feature,omitempty"`
	Nextgenapiresource   string   `json:"_nextgenapiresource,omitempty"`
}

type Auditnslogpolicyappfwglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Auditnslogpolicyauditnslogglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Auditnslogpolicyauthenticationvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Auditnslogpolicyglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Auditnslogpolicylbvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Auditnslogpolicynslogglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Auditsyslogpolicyrnatglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Auditsyslogpolicyvpnvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Auditnslogaction struct {
	Name                 string   `json:"name,omitempty"`
	Serverip             string   `json:"serverip,omitempty"`
	Serverdomainname     string   `json:"serverdomainname,omitempty"`
	Domainresolveretry   int      `json:"domainresolveretry,omitempty"`
	Serverport           int      `json:"serverport,omitempty"`
	Loglevel             []string `json:"loglevel,omitempty"`
	Dateformat           string   `json:"dateformat,omitempty"`
	Logfacility          string   `json:"logfacility,omitempty"`
	Tcp                  string   `json:"tcp,omitempty"`
	Acl                  string   `json:"acl,omitempty"`
	Timezone             string   `json:"timezone,omitempty"`
	Userdefinedauditlog  string   `json:"userdefinedauditlog,omitempty"`
	Appflowexport        string   `json:"appflowexport,omitempty"`
	Lsn                  string   `json:"lsn,omitempty"`
	Alg                  string   `json:"alg,omitempty"`
	Subscriberlog        string   `json:"subscriberlog,omitempty"`
	Sslinterception      string   `json:"sslinterception,omitempty"`
	Urlfiltering         string   `json:"urlfiltering,omitempty"`
	Contentinspectionlog string   `json:"contentinspectionlog,omitempty"`
	Protocolviolations   string   `json:"protocolviolations,omitempty"`
	Domainresolvenow     bool     `json:"domainresolvenow,omitempty"`
	Ip                   string   `json:"ip,omitempty"`
	Builtin              string   `json:"builtin,omitempty"`
	Feature              string   `json:"feature,omitempty"`
	Nextgenapiresource   string   `json:"_nextgenapiresource,omitempty"`
}

type Auditnslogpolicyaaauserbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Auditnslogpolicycsvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Auditnslogpolicysystemglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Auditsyslogglobalsyslogpolicybinding struct {
	Policyname     string   `json:"policyname,omitempty"`
	Priority       uint32   `json:"priority,omitempty"`
	Numpol         uint32   `json:"numpol,omitempty"`
	Globalbindtype string   `json:"globalbindtype,omitempty"`
	Builtin        []string `json:"builtin,omitempty"`
	Feature        string   `json:"feature,omitempty"`
}

type Auditsyslogparams struct {
	Serverip             string   `json:"serverip,omitempty"`
	Serverport           int      `json:"serverport,omitempty"`
	Dateformat           string   `json:"dateformat,omitempty"`
	Loglevel             []string `json:"loglevel,omitempty"`
	Logfacility          string   `json:"logfacility,omitempty"`
	Tcp                  string   `json:"tcp,omitempty"`
	Acl                  string   `json:"acl,omitempty"`
	Timezone             string   `json:"timezone,omitempty"`
	Userdefinedauditlog  string   `json:"userdefinedauditlog,omitempty"`
	Appflowexport        string   `json:"appflowexport,omitempty"`
	Lsn                  string   `json:"lsn,omitempty"`
	Alg                  string   `json:"alg,omitempty"`
	Subscriberlog        string   `json:"subscriberlog,omitempty"`
	Dns                  string   `json:"dns,omitempty"`
	Sslinterception      string   `json:"sslinterception,omitempty"`
	Urlfiltering         string   `json:"urlfiltering,omitempty"`
	Contentinspectionlog string   `json:"contentinspectionlog,omitempty"`
	Streamanalytics      string   `json:"streamanalytics,omitempty"`
	Protocolviolations   string   `json:"protocolviolations,omitempty"`
	Builtin              string   `json:"builtin,omitempty"`
	Feature              string   `json:"feature,omitempty"`
	Nextgenapiresource   string   `json:"_nextgenapiresource,omitempty"`
}

type Auditsyslogpolicy struct {
	Name               string `json:"name,omitempty"`
	Rule               string `json:"rule,omitempty"`
	Action             string `json:"action,omitempty"`
	Builtin            string `json:"builtin,omitempty"`
	Feature            string `json:"feature,omitempty"`
	Expressiontype     string `json:"expressiontype,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Auditmessageaction struct {
	Name               string `json:"name,omitempty"`
	Loglevel           string `json:"loglevel,omitempty"`
	Stringbuilderexpr  string `json:"stringbuilderexpr,omitempty"`
	Logtonewnslog      string `json:"logtonewnslog,omitempty"`
	Bypasssafetycheck  string `json:"bypasssafetycheck,omitempty"`
	Loglevel1          string `json:"loglevel1,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Undefhits          string `json:"undefhits,omitempty"`
	Referencecount     string `json:"referencecount,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Auditnslogpolicyvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Auditsyslogpolicyuserbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Auditsyslogpolicyvpnglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Auditsyslogpolicyvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Auditnslogglobalauditnslogpolicybinding struct {
	Policyname     string   `json:"policyname,omitempty"`
	Priority       int      `json:"priority,omitempty"`
	Numpol         int      `json:"numpol,omitempty"`
	Globalbindtype string   `json:"globalbindtype,omitempty"`
	Builtin        []string `json:"builtin,omitempty"`
}

type Auditnslogparams struct {
	Serverip             string   `json:"serverip,omitempty"`
	Serverport           int      `json:"serverport,omitempty"`
	Dateformat           string   `json:"dateformat,omitempty"`
	Loglevel             []string `json:"loglevel,omitempty"`
	Logfacility          string   `json:"logfacility,omitempty"`
	Tcp                  string   `json:"tcp,omitempty"`
	Acl                  string   `json:"acl,omitempty"`
	Timezone             string   `json:"timezone,omitempty"`
	Userdefinedauditlog  string   `json:"userdefinedauditlog,omitempty"`
	Appflowexport        string   `json:"appflowexport,omitempty"`
	Lsn                  string   `json:"lsn,omitempty"`
	Alg                  string   `json:"alg,omitempty"`
	Subscriberlog        string   `json:"subscriberlog,omitempty"`
	Sslinterception      string   `json:"sslinterception,omitempty"`
	Urlfiltering         string   `json:"urlfiltering,omitempty"`
	Contentinspectionlog string   `json:"contentinspectionlog,omitempty"`
	Protocolviolations   string   `json:"protocolviolations,omitempty"`
	Builtin              string   `json:"builtin,omitempty"`
	Feature              string   `json:"feature,omitempty"`
	Nextgenapiresource   string   `json:"_nextgenapiresource,omitempty"`
}

type Auditnslogpolicy struct {
	Name               string `json:"name,omitempty"`
	Rule               string `json:"rule,omitempty"`
	Action             string `json:"action,omitempty"`
	Builtin            string `json:"builtin,omitempty"`
	Feature            string `json:"feature,omitempty"`
	Expressiontype     string `json:"expressiontype,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}
