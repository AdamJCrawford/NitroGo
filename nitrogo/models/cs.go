// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Csvserverresponderpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Name                   string `json:"name,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
}

type Csvserversyslogpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Sc                     string `json:"sc,omitempty"`
	Name                   string `json:"name,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Cspolicy struct {
	Policyname         string `json:"policyname,omitempty"`
	Rule               string `json:"rule,omitempty"`
	Action             string `json:"action,omitempty"`
	Logaction          string `json:"logaction,omitempty"`
	Newname            string `json:"newname,omitempty"`
	Vstype             string `json:"vstype,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Labelname          string `json:"labelname,omitempty"`
	Labeltype          string `json:"labeltype,omitempty"`
	Priority           string `json:"priority,omitempty"`
	Activepolicy       string `json:"activepolicy,omitempty"`
	Boundto            string `json:"boundto,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Cspolicycrvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Hits                   int    `json:"hits,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
}

type Cspolicycsvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Action                 string `json:"action,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Hits                   int    `json:"hits,omitempty"`
	Bindhits               int    `json:"bindhits,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
}

type Cspolicypolicylabelbinding struct {
	Domain     string `json:"domain,omitempty"`
	Url        string `json:"url,omitempty"`
	Priority   uint32 `json:"priority,omitempty"`
	Hits       uint32 `json:"hits,omitempty"`
	Labeltype  string `json:"labeltype,omitempty"`
	Labelname  string `json:"labelname,omitempty"`
	Policyname string `json:"policyname,omitempty"`
}

type Csvserver struct {
	Name                      string `json:"name,omitempty"`
	Td                        int    `json:"td,omitempty"`
	Servicetype               string `json:"servicetype,omitempty"`
	Ipv46                     string `json:"ipv46,omitempty"`
	Targettype                string `json:"targettype,omitempty"`
	Dnsrecordtype             string `json:"dnsrecordtype,omitempty"`
	Persistenceid             int    `json:"persistenceid,omitempty"`
	Ippattern                 string `json:"ippattern,omitempty"`
	Ipmask                    string `json:"ipmask,omitempty"`
	Range                     int    `json:"range,omitempty"`
	Port                      int    `json:"port,omitempty"`
	Ipset                     string `json:"ipset,omitempty"`
	State                     string `json:"state,omitempty"`
	Stateupdate               string `json:"stateupdate,omitempty"`
	Cacheable                 string `json:"cacheable,omitempty"`
	Redirecturl               string `json:"redirecturl,omitempty"`
	Clttimeout                int    `json:"clttimeout,omitempty"`
	Precedence                string `json:"precedence,omitempty"`
	Casesensitive             string `json:"casesensitive,omitempty"`
	Somethod                  string `json:"somethod,omitempty"`
	Sopersistence             string `json:"sopersistence,omitempty"`
	Sopersistencetimeout      int    `json:"sopersistencetimeout,omitempty"`
	Sothreshold               int    `json:"sothreshold,omitempty"`
	Sobackupaction            string `json:"sobackupaction,omitempty"`
	Redirectportrewrite       string `json:"redirectportrewrite,omitempty"`
	Downstateflush            string `json:"downstateflush,omitempty"`
	Backupvserver             string `json:"backupvserver,omitempty"`
	Disableprimaryondown      string `json:"disableprimaryondown,omitempty"`
	Insertvserveripport       string `json:"insertvserveripport,omitempty"`
	Vipheader                 string `json:"vipheader,omitempty"`
	Rtspnat                   string `json:"rtspnat,omitempty"`
	Authenticationhost        string `json:"authenticationhost,omitempty"`
	Authentication            string `json:"authentication,omitempty"`
	Listenpolicy              string `json:"listenpolicy,omitempty"`
	Listenpriority            int    `json:"listenpriority,omitempty"`
	Authn401                  string `json:"authn401,omitempty"`
	Authnvsname               string `json:"authnvsname,omitempty"`
	Push                      string `json:"push,omitempty"`
	Pushvserver               string `json:"pushvserver,omitempty"`
	Pushlabel                 string `json:"pushlabel,omitempty"`
	Pushmulticlients          string `json:"pushmulticlients,omitempty"`
	Tcpprofilename            string `json:"tcpprofilename,omitempty"`
	Httpprofilename           string `json:"httpprofilename,omitempty"`
	Dbprofilename             string `json:"dbprofilename,omitempty"`
	Oracleserverversion       string `json:"oracleserverversion,omitempty"`
	Comment                   string `json:"comment,omitempty"`
	Mssqlserverversion        string `json:"mssqlserverversion,omitempty"`
	L2conn                    string `json:"l2conn,omitempty"`
	Mysqlprotocolversion      int    `json:"mysqlprotocolversion,omitempty"`
	Mysqlserverversion        string `json:"mysqlserverversion,omitempty"`
	Mysqlcharacterset         int    `json:"mysqlcharacterset,omitempty"`
	Mysqlservercapabilities   int    `json:"mysqlservercapabilities,omitempty"`
	Appflowlog                string `json:"appflowlog,omitempty"`
	Netprofile                string `json:"netprofile,omitempty"`
	Icmpvsrresponse           string `json:"icmpvsrresponse,omitempty"`
	Rhistate                  string `json:"rhistate,omitempty"`
	Authnprofile              string `json:"authnprofile,omitempty"`
	Dnsprofilename            string `json:"dnsprofilename,omitempty"`
	Dtls                      string `json:"dtls,omitempty"`
	Persistencetype           string `json:"persistencetype,omitempty"`
	Persistmask               string `json:"persistmask,omitempty"`
	V6persistmasklen          int    `json:"v6persistmasklen,omitempty"`
	Timeout                   int    `json:"timeout,omitempty"`
	Cookiename                string `json:"cookiename,omitempty"`
	Persistencebackup         string `json:"persistencebackup,omitempty"`
	Backuppersistencetimeout  int    `json:"backuppersistencetimeout,omitempty"`
	Tcpprobeport              int    `json:"tcpprobeport,omitempty"`
	Probeprotocol             string `json:"probeprotocol,omitempty"`
	Probesuccessresponsecode  string `json:"probesuccessresponsecode,omitempty"`
	Probeport                 int    `json:"probeport,omitempty"`
	Quicprofilename           string `json:"quicprofilename,omitempty"`
	Redirectfromport          int    `json:"redirectfromport,omitempty"`
	Dnsoverhttps              string `json:"dnsoverhttps,omitempty"`
	Httpsredirecturl          string `json:"httpsredirecturl,omitempty"`
	Apiprofile                string `json:"apiprofile,omitempty"`
	Domainname                string `json:"domainname,omitempty"`
	Ttl                       int    `json:"ttl,omitempty"`
	Backupip                  string `json:"backupip,omitempty"`
	Cookiedomain              string `json:"cookiedomain,omitempty"`
	Cookietimeout             int    `json:"cookietimeout,omitempty"`
	Sitedomainttl             int    `json:"sitedomainttl,omitempty"`
	Newname                   string `json:"newname,omitempty"`
	Ip                        string `json:"ip,omitempty"`
	Value                     string `json:"value,omitempty"`
	Ngname                    string `json:"ngname,omitempty"`
	Type                      string `json:"type,omitempty"`
	Curstate                  string `json:"curstate,omitempty"`
	Status                    string `json:"status,omitempty"`
	Cachetype                 string `json:"cachetype,omitempty"`
	Redirect                  string `json:"redirect,omitempty"`
	Homepage                  string `json:"homepage,omitempty"`
	Dnsvservername            string `json:"dnsvservername,omitempty"`
	Domain                    string `json:"domain,omitempty"`
	Servicename               string `json:"servicename,omitempty"`
	Weight                    string `json:"weight,omitempty"`
	Cachevserver              string `json:"cachevserver,omitempty"`
	Targetvserver             string `json:"targetvserver,omitempty"`
	Url                       string `json:"url,omitempty"`
	Bindpoint                 string `json:"bindpoint,omitempty"`
	Gt2gb                     string `json:"gt2gb,omitempty"`
	Statechangetimesec        string `json:"statechangetimesec,omitempty"`
	Statechangetimemsec       string `json:"statechangetimemsec,omitempty"`
	Tickssincelaststatechange string `json:"tickssincelaststatechange,omitempty"`
	Ruletype                  string `json:"ruletype,omitempty"`
	Lbvserver                 string `json:"lbvserver,omitempty"`
	Targetlbvserver           string `json:"targetlbvserver,omitempty"`
	Nodefaultbindings         string `json:"nodefaultbindings,omitempty"`
	Version                   string `json:"version,omitempty"`
	Nextgenapiresource        string `json:"_nextgenapiresource,omitempty"`
}

type Csvserveranalyticsprofilebinding struct {
	Analyticsprofile string `json:"analyticsprofile,omitempty"`
	Name             string `json:"name,omitempty"`
}

type Csvservercachepolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
}

type Csvserverdomainbinding struct {
	Domainname    string `json:"domainname,omitempty"`
	Ttl           int    `json:"ttl,omitempty"`
	Backupip      string `json:"backupip,omitempty"`
	Cookiedomain  string `json:"cookiedomain,omitempty"`
	Cookietimeout int    `json:"cookietimeout,omitempty"`
	Sitedomainttl int    `json:"sitedomainttl,omitempty"`
	Appflowlog    string `json:"appflowlog,omitempty"`
	Name          string `json:"name,omitempty"`
}

type Csparameter struct {
	Stateupdate        string `json:"stateupdate,omitempty"`
	Builtin            string `json:"builtin,omitempty"`
	Feature            string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Cspolicybinding struct {
	Policyname string `json:"policyname,omitempty"`
}

type Csvserverauditnslogpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Name                   string `json:"name,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Csvserverauthorizationpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
}

type Csvserverlbvserverbinding struct {
	Lbvserver     string `json:"lbvserver,omitempty"`
	Hits          int    `json:"hits,omitempty"`
	Vserverid     string `json:"vserverid,omitempty"`
	Cookieipport  string `json:"cookieipport,omitempty"`
	Name          string `json:"name,omitempty"`
	Targetvserver string `json:"targetvserver,omitempty"`
}

type Csvserverspilloverpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Name                   string `json:"name,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Csvservertmtrafficpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Name                   string `json:"name,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Cspolicyvserverbinding struct {
	Domain       string `json:"domain,omitempty"`
	Action       string `json:"action,omitempty"`
	Url          string `json:"url,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Hits         uint32 `json:"hits,omitempty"`
	Pihits       uint32 `json:"pihits,omitempty"`
	Pipolicyhits uint32 `json:"pipolicyhits,omitempty"`
	Labeltype    string `json:"labeltype,omitempty"`
	Labelname    string `json:"labelname,omitempty"`
	Policyname   string `json:"policyname,omitempty"`
}

type Csvservercmppolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
}

type Csvservernslogpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Sc                     string `json:"sc,omitempty"`
	Name                   string `json:"name,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Csvserverrewritepolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
}

type Csvservervpnvserverbinding struct {
	Vserver string `json:"vserver,omitempty"`
	Hits    int    `json:"hits,omitempty"`
	Name    string `json:"name,omitempty"`
}

type Cspolicycspolicylabelbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Hits                   int    `json:"hits,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
}

type Cspolicylabel struct {
	Labelname              string `json:"labelname,omitempty"`
	Cspolicylabeltype      string `json:"cspolicylabeltype,omitempty"`
	Newname                string `json:"newname,omitempty"`
	Numpol                 string `json:"numpol,omitempty"`
	Hits                   string `json:"hits,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               string `json:"priority,omitempty"`
	Targetvserver          string `json:"targetvserver,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Nextgenapiresource     string `json:"_nextgenapiresource,omitempty"`
}

type Cspolicylabelcspolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Targetvserver          string `json:"targetvserver,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Csvserverappfwpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
}

type Csvservervserverbinding struct {
	Vserver       string `json:"vserver,omitempty"`
	Lbvserver     string `json:"lbvserver,omitempty"`
	Hits          uint32 `json:"hits,omitempty"`
	Vserverid     string `json:"vserverid,omitempty"`
	Cookieipport  string `json:"cookieipport,omitempty"`
	Name          string `json:"name,omitempty"`
	Targetvserver string `json:"targetvserver,omitempty"`
}

type Csvserverappflowpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Csvserverappqoepolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
}

type Csvserverbotpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Csvserverfeopolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Name                   string `json:"name,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Csvserverfilterpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Sc                     string `json:"sc,omitempty"`
	Name                   string `json:"name,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Csvservertrafficpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Name                   string `json:"name,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Csaction struct {
	Name               string `json:"name,omitempty"`
	Targetlbvserver    string `json:"targetlbvserver,omitempty"`
	Targetvserver      string `json:"targetvserver,omitempty"`
	Targetvserverexpr  string `json:"targetvserverexpr,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Newname            string `json:"newname,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Referencecount     string `json:"referencecount,omitempty"`
	Undefhits          string `json:"undefhits,omitempty"`
	Builtin            string `json:"builtin,omitempty"`
	Feature            string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Csvserverprofilebinding struct {
	Analyticsprofile string `json:"analyticsprofile,omitempty"`
	Name             string `json:"name,omitempty"`
}

type Csvservertransformpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
}

type Cspolicylabelpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Targetvserver          string `json:"targetvserver,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Csvserverbinding struct {
	Name string `json:"name,omitempty"`
}

type Csvservercontentinspectionpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
}

type Cspolicylabelbinding struct {
	Labelname string `json:"labelname,omitempty"`
}

type Csvserverauditsyslogpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Name                   string `json:"name,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Csvservercspolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Hits                   int    `json:"hits,omitempty"`
	Pipolicyhits           int    `json:"pipolicyhits,omitempty"`
	Rule                   string `json:"rule,omitempty"`
	Vserverid              string `json:"vserverid,omitempty"`
	Cookieipport           string `json:"cookieipport,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Csvservergslbvserverbinding struct {
	Vserver string `json:"vserver,omitempty"`
	Hits    int    `json:"hits,omitempty"`
	Name    string `json:"name,omitempty"`
}

type Csvserverpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Sc                     string `json:"sc,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Targetlbvserver        string `json:"targetlbvserver,omitempty"`
	Hits                   uint32 `json:"hits,omitempty"`
	Pipolicyhits           uint32 `json:"pipolicyhits,omitempty"`
	Rule                   string `json:"rule,omitempty"`
	Vserverid              string `json:"vserverid,omitempty"`
	Cookieipport           string `json:"cookieipport,omitempty"`
	Name                   string `json:"name,omitempty"`
}
