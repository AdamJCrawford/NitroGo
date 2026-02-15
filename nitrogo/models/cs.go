package models

// cs configuration structs
type CsvserverTransformpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
}

type CsvserverDomainBinding struct {
	Appflowlog    string `json:"appflowlog,omitempty"`
	Backupip      string `json:"backupip,omitempty"`
	Cookiedomain  string `json:"cookiedomain,omitempty"`
	Cookietimeout int    `json:"cookietimeout,omitempty"`
	Domainname    string `json:"domainname,omitempty"`
	Name          string `json:"name,omitempty"`
	Sitedomainttl int    `json:"sitedomainttl,omitempty"`
	Ttl           int    `json:"ttl,omitempty"`
}

type CsvserverBotpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
}

type Csparameter struct {
	Builtin            []string `json:"builtin,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Stateupdate        string   `json:"stateupdate,omitempty"`
}

type CsvserverVpnvserverBinding struct {
	Hits    int    `json:"hits,omitempty"`
	Name    string `json:"name,omitempty"`
	Vserver string `json:"vserver,omitempty"`
}

type CsvserverResponderpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
}

type Cspolicy struct {
	Action             string  `json:"action,omitempty"`
	Activepolicy       bool    `json:"activepolicy,omitempty"`
	Boundto            string  `json:"boundto,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Hits               int     `json:"hits,omitempty"`
	Labelname          string  `json:"labelname,omitempty"`
	Labeltype          string  `json:"labeltype,omitempty"`
	Logaction          string  `json:"logaction,omitempty"`
	Newname            string  `json:"newname,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Policyname         string  `json:"policyname,omitempty"`
	Priority           int     `json:"priority,omitempty"`
	Rule               string  `json:"rule,omitempty"`
	Vstype             int     `json:"vstype,omitempty"`
}

type CsvserverBinding struct {
	CsvserverAnalyticsprofileBinding        []interface{} `json:"csvserver_analyticsprofile_binding,omitempty"`
	CsvserverAppflowpolicyBinding           []interface{} `json:"csvserver_appflowpolicy_binding,omitempty"`
	CsvserverAppfwpolicyBinding             []interface{} `json:"csvserver_appfwpolicy_binding,omitempty"`
	CsvserverAppqoepolicyBinding            []interface{} `json:"csvserver_appqoepolicy_binding,omitempty"`
	CsvserverAuditnslogpolicyBinding        []interface{} `json:"csvserver_auditnslogpolicy_binding,omitempty"`
	CsvserverAuditsyslogpolicyBinding       []interface{} `json:"csvserver_auditsyslogpolicy_binding,omitempty"`
	CsvserverAuthorizationpolicyBinding     []interface{} `json:"csvserver_authorizationpolicy_binding,omitempty"`
	CsvserverBotpolicyBinding               []interface{} `json:"csvserver_botpolicy_binding,omitempty"`
	CsvserverCachepolicyBinding             []interface{} `json:"csvserver_cachepolicy_binding,omitempty"`
	CsvserverCmppolicyBinding               []interface{} `json:"csvserver_cmppolicy_binding,omitempty"`
	CsvserverContentinspectionpolicyBinding []interface{} `json:"csvserver_contentinspectionpolicy_binding,omitempty"`
	CsvserverCspolicyBinding                []interface{} `json:"csvserver_cspolicy_binding,omitempty"`
	CsvserverFeopolicyBinding               []interface{} `json:"csvserver_feopolicy_binding,omitempty"`
	CsvserverGslbdomainBinding              []interface{} `json:"csvserver_gslbdomain_binding,omitempty"`
	CsvserverGslbvserverBinding             []interface{} `json:"csvserver_gslbvserver_binding,omitempty"`
	CsvserverLbvserverBinding               []interface{} `json:"csvserver_lbvserver_binding,omitempty"`
	CsvserverResponderpolicyBinding         []interface{} `json:"csvserver_responderpolicy_binding,omitempty"`
	CsvserverRewritepolicyBinding           []interface{} `json:"csvserver_rewritepolicy_binding,omitempty"`
	CsvserverSpilloverpolicyBinding         []interface{} `json:"csvserver_spilloverpolicy_binding,omitempty"`
	CsvserverTmtrafficpolicyBinding         []interface{} `json:"csvserver_tmtrafficpolicy_binding,omitempty"`
	CsvserverTransformpolicyBinding         []interface{} `json:"csvserver_transformpolicy_binding,omitempty"`
	CsvserverVpnvserverBinding              []interface{} `json:"csvserver_vpnvserver_binding,omitempty"`
	Name                                    string        `json:"name,omitempty"`
}

type CspolicyCsvserverBinding struct {
	Action                 string `json:"action,omitempty"`
	Bindhits               int    `json:"bindhits,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Hits                   int    `json:"hits,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type CspolicyCspolicylabelBinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Hits                   int    `json:"hits,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type CsvserverAuditnslogpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
}

type CsvserverAppflowpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
}

type CsvserverAuditsyslogpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
}

type CsvserverTmtrafficpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
}

type CsvserverCachepolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
}

type CsvserverAppfwpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
}

type CsvserverCspolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Cookieipport           string `json:"cookieipport,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Hits                   int    `json:"hits,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Pipolicyhits           int    `json:"pipolicyhits,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Rule                   string `json:"rule,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
	Vserverid              string `json:"vserverid,omitempty"`
}

type CsvserverAppqoepolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
}

type CsvserverSpilloverpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
}

type CsvserverAuthorizationpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
}

type CsvserverAnalyticsprofileBinding struct {
	Analyticsprofile string `json:"analyticsprofile,omitempty"`
	Name             string `json:"name,omitempty"`
}

type CspolicyCrvserverBinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Hits                   int    `json:"hits,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type CspolicylabelBinding struct {
	CspolicylabelCspolicyBinding []interface{} `json:"cspolicylabel_cspolicy_binding,omitempty"`
	Labelname                    string        `json:"labelname,omitempty"`
}

type CsvserverContentinspectionpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
}

type Cspolicylabel struct {
	Count                  float64 `json:"__count,omitempty"`
	Cspolicylabeltype      string  `json:"cspolicylabeltype,omitempty"`
	Gotopriorityexpression string  `json:"gotopriorityexpression,omitempty"`
	Hits                   int     `json:"hits,omitempty"`
	InvokeLabelname        string  `json:"invoke_labelname,omitempty"`
	Labelname              string  `json:"labelname,omitempty"`
	Labeltype              string  `json:"labeltype,omitempty"`
	Newname                string  `json:"newname,omitempty"`
	Nextgenapiresource     string  `json:"_nextgenapiresource,omitempty"`
	Numpol                 int     `json:"numpol,omitempty"`
	Policyname             string  `json:"policyname,omitempty"`
	Priority               int     `json:"priority,omitempty"`
	Targetvserver          string  `json:"targetvserver,omitempty"`
}

type CsvserverGslbvserverBinding struct {
	Hits    int    `json:"hits,omitempty"`
	Name    string `json:"name,omitempty"`
	Vserver string `json:"vserver,omitempty"`
}

type CsvserverFeopolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
}

type CsvserverRewritepolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
}

type CsvserverCmppolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
}

type Csaction struct {
	Builtin            []string `json:"builtin,omitempty"`
	Comment            string   `json:"comment,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	Name               string   `json:"name,omitempty"`
	Newname            string   `json:"newname,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Referencecount     int      `json:"referencecount,omitempty"`
	Targetlbvserver    string   `json:"targetlbvserver,omitempty"`
	Targetvserver      string   `json:"targetvserver,omitempty"`
	Targetvserverexpr  string   `json:"targetvserverexpr,omitempty"`
	Undefhits          int      `json:"undefhits,omitempty"`
}

type CspolicylabelCspolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Targetvserver          string `json:"targetvserver,omitempty"`
}

type Csvserver struct {
	Apiprofile                string  `json:"apiprofile,omitempty"`
	Appflowlog                string  `json:"appflowlog,omitempty"`
	Authentication            string  `json:"authentication,omitempty"`
	Authenticationhost        string  `json:"authenticationhost,omitempty"`
	Authn401                  string  `json:"authn401,omitempty"`
	Authnprofile              string  `json:"authnprofile,omitempty"`
	Authnvsname               string  `json:"authnvsname,omitempty"`
	Backupip                  string  `json:"backupip,omitempty"`
	Backuppersistencetimeout  int     `json:"backuppersistencetimeout,omitempty"`
	Backupvserver             string  `json:"backupvserver,omitempty"`
	Bindpoint                 string  `json:"bindpoint,omitempty"`
	Cacheable                 string  `json:"cacheable,omitempty"`
	Cachetype                 string  `json:"cachetype,omitempty"`
	Cachevserver              string  `json:"cachevserver,omitempty"`
	Casesensitive             string  `json:"casesensitive,omitempty"`
	Clttimeout                int     `json:"clttimeout,omitempty"`
	Comment                   string  `json:"comment,omitempty"`
	Cookiedomain              string  `json:"cookiedomain,omitempty"`
	Cookiename                string  `json:"cookiename,omitempty"`
	Cookietimeout             int     `json:"cookietimeout,omitempty"`
	Count                     float64 `json:"__count,omitempty"`
	Curstate                  string  `json:"curstate,omitempty"`
	Dbprofilename             string  `json:"dbprofilename,omitempty"`
	Disableprimaryondown      string  `json:"disableprimaryondown,omitempty"`
	Dnsoverhttps              string  `json:"dnsoverhttps,omitempty"`
	Dnsprofilename            string  `json:"dnsprofilename,omitempty"`
	Dnsrecordtype             string  `json:"dnsrecordtype,omitempty"`
	Dnsvservername            string  `json:"dnsvservername,omitempty"`
	Domain                    string  `json:"domain,omitempty"`
	Domainname                string  `json:"domainname,omitempty"`
	Downstateflush            string  `json:"downstateflush,omitempty"`
	Dtls                      string  `json:"dtls,omitempty"`
	Gt2gb                     string  `json:"gt2gb,omitempty"`
	Homepage                  string  `json:"homepage,omitempty"`
	Httpprofilename           string  `json:"httpprofilename,omitempty"`
	Httpsredirecturl          string  `json:"httpsredirecturl,omitempty"`
	Icmpvsrresponse           string  `json:"icmpvsrresponse,omitempty"`
	Insertvserveripport       string  `json:"insertvserveripport,omitempty"`
	Ip                        string  `json:"ip,omitempty"`
	Ipmask                    string  `json:"ipmask,omitempty"`
	Ippattern                 string  `json:"ippattern,omitempty"`
	Ipset                     string  `json:"ipset,omitempty"`
	Ipv46                     string  `json:"ipv46,omitempty"`
	L2conn                    string  `json:"l2conn,omitempty"`
	Lbvserver                 string  `json:"lbvserver,omitempty"`
	Listenpolicy              string  `json:"listenpolicy,omitempty"`
	Listenpriority            int     `json:"listenpriority,omitempty"`
	Mssqlserverversion        string  `json:"mssqlserverversion,omitempty"`
	Mysqlcharacterset         int     `json:"mysqlcharacterset,omitempty"`
	Mysqlprotocolversion      int     `json:"mysqlprotocolversion,omitempty"`
	Mysqlservercapabilities   int     `json:"mysqlservercapabilities,omitempty"`
	Mysqlserverversion        string  `json:"mysqlserverversion,omitempty"`
	Name                      string  `json:"name,omitempty"`
	Netprofile                string  `json:"netprofile,omitempty"`
	Newname                   string  `json:"newname,omitempty"`
	Nextgenapiresource        string  `json:"_nextgenapiresource,omitempty"`
	Ngname                    string  `json:"ngname,omitempty"`
	Nodefaultbindings         string  `json:"nodefaultbindings,omitempty"`
	Oracleserverversion       string  `json:"oracleserverversion,omitempty"`
	Persistencebackup         string  `json:"persistencebackup,omitempty"`
	Persistenceid             int     `json:"persistenceid,omitempty"`
	Persistencetype           string  `json:"persistencetype,omitempty"`
	Persistmask               string  `json:"persistmask,omitempty"`
	Port                      int     `json:"port,omitempty"`
	Precedence                string  `json:"precedence,omitempty"`
	Probeport                 int     `json:"probeport,omitempty"`
	Probeprotocol             string  `json:"probeprotocol,omitempty"`
	Probesuccessresponsecode  string  `json:"probesuccessresponsecode,omitempty"`
	Push                      string  `json:"push,omitempty"`
	Pushlabel                 string  `json:"pushlabel,omitempty"`
	Pushmulticlients          string  `json:"pushmulticlients,omitempty"`
	Pushvserver               string  `json:"pushvserver,omitempty"`
	Quicprofilename           string  `json:"quicprofilename,omitempty"`
	Range                     int     `json:"range,omitempty"`
	Redirect                  string  `json:"redirect,omitempty"`
	Redirectfromport          int     `json:"redirectfromport,omitempty"`
	Redirectportrewrite       string  `json:"redirectportrewrite,omitempty"`
	Redirecturl               string  `json:"redirecturl,omitempty"`
	Rhistate                  string  `json:"rhistate,omitempty"`
	Rtspnat                   string  `json:"rtspnat,omitempty"`
	Ruletype                  int     `json:"ruletype,omitempty"`
	Servicename               string  `json:"servicename,omitempty"`
	Servicetype               string  `json:"servicetype,omitempty"`
	Sitedomainttl             int     `json:"sitedomainttl,omitempty"`
	Sobackupaction            string  `json:"sobackupaction,omitempty"`
	Somethod                  string  `json:"somethod,omitempty"`
	Sopersistence             string  `json:"sopersistence,omitempty"`
	Sopersistencetimeout      int     `json:"sopersistencetimeout,omitempty"`
	Sothreshold               int     `json:"sothreshold,omitempty"`
	State                     string  `json:"state,omitempty"`
	Statechangetimemsec       int     `json:"statechangetimemsec,omitempty"`
	Statechangetimesec        string  `json:"statechangetimesec,omitempty"`
	Stateupdate               string  `json:"stateupdate,omitempty"`
	Status                    int     `json:"status,omitempty"`
	Targetlbvserver           string  `json:"targetlbvserver,omitempty"`
	Targettype                string  `json:"targettype,omitempty"`
	Targetvserver             string  `json:"targetvserver,omitempty"`
	Tcpprobeport              int     `json:"tcpprobeport,omitempty"`
	Tcpprofilename            string  `json:"tcpprofilename,omitempty"`
	Td                        int     `json:"td,omitempty"`
	Tickssincelaststatechange int     `json:"tickssincelaststatechange,omitempty"`
	Timeout                   int     `json:"timeout,omitempty"`
	Ttl                       int     `json:"ttl,omitempty"`
	TypeField                 string  `json:"type,omitempty"`
	Url                       string  `json:"url,omitempty"`
	V6persistmasklen          int     `json:"v6persistmasklen,omitempty"`
	Value                     string  `json:"value,omitempty"`
	Version                   int     `json:"version,omitempty"`
	Vipheader                 string  `json:"vipheader,omitempty"`
	Weight                    int     `json:"weight,omitempty"`
}

type CspolicyBinding struct {
	CspolicyCrvserverBinding     []interface{} `json:"cspolicy_crvserver_binding,omitempty"`
	CspolicyCspolicylabelBinding []interface{} `json:"cspolicy_cspolicylabel_binding,omitempty"`
	CspolicyCsvserverBinding     []interface{} `json:"cspolicy_csvserver_binding,omitempty"`
	Policyname                   string        `json:"policyname,omitempty"`
}

type CsvserverLbvserverBinding struct {
	Cookieipport  string `json:"cookieipport,omitempty"`
	Hits          int    `json:"hits,omitempty"`
	Lbvserver     string `json:"lbvserver,omitempty"`
	Name          string `json:"name,omitempty"`
	Targetvserver string `json:"targetvserver,omitempty"`
	Vserverid     string `json:"vserverid,omitempty"`
}
