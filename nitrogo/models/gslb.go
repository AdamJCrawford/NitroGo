package models

// gslb configuration structs
type GslbservicegroupBinding struct {
	GslbservicegroupGslbservicegroupmemberBinding        []interface{} `json:"gslbservicegroup_gslbservicegroupmember_binding,omitempty"`
	GslbservicegroupLbmonitorBinding                     []interface{} `json:"gslbservicegroup_lbmonitor_binding,omitempty"`
	GslbservicegroupServicegroupentitymonbindingsBinding []interface{} `json:"gslbservicegroup_servicegroupentitymonbindings_binding,omitempty"`
	Servicegroupname                                     string        `json:"servicegroupname,omitempty"`
}

type Gslbservicegroup struct {
	Appflowlog                 string  `json:"appflowlog,omitempty"`
	Autodelayedtrofs           string  `json:"autodelayedtrofs,omitempty"`
	Autoscale                  string  `json:"autoscale,omitempty"`
	Cip                        string  `json:"cip,omitempty"`
	Cipheader                  string  `json:"cipheader,omitempty"`
	Clmonowner                 int     `json:"clmonowner,omitempty"`
	Clmonview                  int     `json:"clmonview,omitempty"`
	Clttimeout                 int     `json:"clttimeout,omitempty"`
	Comment                    string  `json:"comment,omitempty"`
	Count                      float64 `json:"__count,omitempty"`
	Delay                      int     `json:"delay,omitempty"`
	Downstateflush             string  `json:"downstateflush,omitempty"`
	DupWeight                  int     `json:"dup_weight,omitempty"`
	Graceful                   string  `json:"graceful,omitempty"`
	Groupcount                 int     `json:"groupcount,omitempty"`
	Gslb                       string  `json:"gslb,omitempty"`
	Hashid                     int     `json:"hashid,omitempty"`
	Healthmonitor              string  `json:"healthmonitor,omitempty"`
	Includemembers             bool    `json:"includemembers,omitempty"`
	Ip                         string  `json:"ip,omitempty"`
	Maxbandwidth               int     `json:"maxbandwidth,omitempty"`
	Maxclient                  int     `json:"maxclient,omitempty"`
	MonitorNameSvc             string  `json:"monitor_name_svc,omitempty"`
	Monstatcode                int     `json:"monstatcode,omitempty"`
	Monstatparam1              int     `json:"monstatparam1,omitempty"`
	Monstatparam2              int     `json:"monstatparam2,omitempty"`
	Monstatparam3              int     `json:"monstatparam3,omitempty"`
	Monthreshold               int     `json:"monthreshold,omitempty"`
	Newname                    string  `json:"newname,omitempty"`
	Nextgenapiresource         string  `json:"_nextgenapiresource,omitempty"`
	Nodefaultbindings          string  `json:"nodefaultbindings,omitempty"`
	Numofconnections           int     `json:"numofconnections,omitempty"`
	Order                      int     `json:"order,omitempty"`
	Port                       int     `json:"port,omitempty"`
	Publicip                   string  `json:"publicip,omitempty"`
	Publicport                 int     `json:"publicport,omitempty"`
	Servername                 string  `json:"servername,omitempty"`
	Serviceconftype            bool    `json:"serviceconftype,omitempty"`
	Servicegroupeffectivestate string  `json:"servicegroupeffectivestate,omitempty"`
	Servicegroupname           string  `json:"servicegroupname,omitempty"`
	Serviceipstr               string  `json:"serviceipstr,omitempty"`
	Servicetype                string  `json:"servicetype,omitempty"`
	Sitename                   string  `json:"sitename,omitempty"`
	Sitepersistence            string  `json:"sitepersistence,omitempty"`
	Siteprefix                 string  `json:"siteprefix,omitempty"`
	State                      string  `json:"state,omitempty"`
	Statechangetimemsec        int     `json:"statechangetimemsec,omitempty"`
	Stateupdatereason          int     `json:"stateupdatereason,omitempty"`
	Svreffgslbstate            string  `json:"svreffgslbstate,omitempty"`
	Svrstate                   string  `json:"svrstate,omitempty"`
	Svrtimeout                 int     `json:"svrtimeout,omitempty"`
	Value                      string  `json:"value,omitempty"`
	Weight                     int     `json:"weight,omitempty"`
}

type Gslbldnsentries struct {
	Count              float64       `json:"__count,omitempty"`
	Ipaddress          string        `json:"ipaddress,omitempty"`
	Name               string        `json:"name,omitempty"`
	Nextgenapiresource string        `json:"_nextgenapiresource,omitempty"`
	Nodeid             int           `json:"nodeid,omitempty"`
	Numsites           int           `json:"numsites,omitempty"`
	Rtt                []interface{} `json:"rtt,omitempty"`
	Sitename           string        `json:"sitename,omitempty"`
	Ttl                int           `json:"ttl,omitempty"`
}

type GslbservicegroupServicegroupentitymonbindingsBinding struct {
	Hashid                     int    `json:"hashid,omitempty"`
	Lastresponse               string `json:"lastresponse,omitempty"`
	MonitorName                string `json:"monitor_name,omitempty"`
	MonitorState               string `json:"monitor_state,omitempty"`
	Monitorcurrentfailedprobes int    `json:"monitorcurrentfailedprobes,omitempty"`
	Monitortotalfailedprobes   int    `json:"monitortotalfailedprobes,omitempty"`
	Monitortotalprobes         int    `json:"monitortotalprobes,omitempty"`
	Order                      int    `json:"order,omitempty"`
	Passive                    bool   `json:"passive,omitempty"`
	Port                       int    `json:"port,omitempty"`
	Publicip                   string `json:"publicip,omitempty"`
	Publicport                 int    `json:"publicport,omitempty"`
	Servicegroupentname2       string `json:"servicegroupentname2,omitempty"`
	Servicegroupname           string `json:"servicegroupname,omitempty"`
	Siteprefix                 string `json:"siteprefix,omitempty"`
	State                      string `json:"state,omitempty"`
	Weight                     int    `json:"weight,omitempty"`
}

type GslbvserverLbpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
}

type Gslbvserver struct {
	Activeservices            int     `json:"activeservices,omitempty"`
	Appflowlog                string  `json:"appflowlog,omitempty"`
	Backupip                  string  `json:"backupip,omitempty"`
	Backuplbmethod            string  `json:"backuplbmethod,omitempty"`
	Backupsessiontimeout      int     `json:"backupsessiontimeout,omitempty"`
	Backupvserver             string  `json:"backupvserver,omitempty"`
	Comment                   string  `json:"comment,omitempty"`
	Considereffectivestate    string  `json:"considereffectivestate,omitempty"`
	CookieDomain              string  `json:"cookie_domain,omitempty"`
	Cookietimeout             int     `json:"cookietimeout,omitempty"`
	Count                     float64 `json:"__count,omitempty"`
	Currentactiveorder        string  `json:"currentactiveorder,omitempty"`
	Curstate                  string  `json:"curstate,omitempty"`
	Disableprimaryondown      string  `json:"disableprimaryondown,omitempty"`
	Dnsrecordtype             string  `json:"dnsrecordtype,omitempty"`
	Domainname                string  `json:"domainname,omitempty"`
	Dynamicweight             string  `json:"dynamicweight,omitempty"`
	Ecs                       string  `json:"ecs,omitempty"`
	Ecsaddrvalidation         string  `json:"ecsaddrvalidation,omitempty"`
	Edr                       string  `json:"edr,omitempty"`
	Gotopriorityexpression    string  `json:"gotopriorityexpression,omitempty"`
	Health                    int     `json:"health,omitempty"`
	Iptype                    string  `json:"iptype,omitempty"`
	Iscname                   string  `json:"iscname,omitempty"`
	Lbmethod                  string  `json:"lbmethod,omitempty"`
	Lbrrreason                int     `json:"lbrrreason,omitempty"`
	Mir                       string  `json:"mir,omitempty"`
	Name                      string  `json:"name,omitempty"`
	Netmask                   string  `json:"netmask,omitempty"`
	Newname                   string  `json:"newname,omitempty"`
	Nextgenapiresource        string  `json:"_nextgenapiresource,omitempty"`
	Nodefaultbindings         string  `json:"nodefaultbindings,omitempty"`
	Order                     int     `json:"order,omitempty"`
	Orderthreshold            int     `json:"orderthreshold,omitempty"`
	Persistenceid             int     `json:"persistenceid,omitempty"`
	Persistencetype           string  `json:"persistencetype,omitempty"`
	Persistmask               string  `json:"persistmask,omitempty"`
	Policyname                string  `json:"policyname,omitempty"`
	Priority                  int     `json:"priority,omitempty"`
	Rule                      string  `json:"rule,omitempty"`
	Servername                string  `json:"servername,omitempty"`
	Servicegroupname          string  `json:"servicegroupname,omitempty"`
	Servicename               string  `json:"servicename,omitempty"`
	Servicetype               string  `json:"servicetype,omitempty"`
	Sitedomainttl             int     `json:"sitedomainttl,omitempty"`
	Sitepersistence           string  `json:"sitepersistence,omitempty"`
	Sobackupaction            string  `json:"sobackupaction,omitempty"`
	Somethod                  string  `json:"somethod,omitempty"`
	Sopersistence             string  `json:"sopersistence,omitempty"`
	Sopersistencetimeout      int     `json:"sopersistencetimeout,omitempty"`
	Sothreshold               int     `json:"sothreshold,omitempty"`
	State                     string  `json:"state,omitempty"`
	Statechangetimemsec       int     `json:"statechangetimemsec,omitempty"`
	Statechangetimesec        string  `json:"statechangetimesec,omitempty"`
	Status                    int     `json:"status,omitempty"`
	Tickssincelaststatechange int     `json:"tickssincelaststatechange,omitempty"`
	Timeout                   int     `json:"timeout,omitempty"`
	Toggleorder               string  `json:"toggleorder,omitempty"`
	Tolerance                 int     `json:"tolerance,omitempty"`
	Totalservices             int     `json:"totalservices,omitempty"`
	Ttl                       int     `json:"ttl,omitempty"`
	TypeField                 string  `json:"type,omitempty"`
	V6netmasklen              int     `json:"v6netmasklen,omitempty"`
	V6persistmasklen          int     `json:"v6persistmasklen,omitempty"`
	Vsvrbindsvcip             string  `json:"vsvrbindsvcip,omitempty"`
	Vsvrbindsvcport           int     `json:"vsvrbindsvcport,omitempty"`
	Weight                    int     `json:"weight,omitempty"`
}

type GslbdomainLbmonitorBinding struct {
	Customheaders              string `json:"customheaders,omitempty"`
	Grpchealthcheck            string `json:"grpchealthcheck,omitempty"`
	Grpcservicename            string `json:"grpcservicename,omitempty"`
	Grpcstatuscode             int    `json:"grpcstatuscode,omitempty"`
	Httprequest                string `json:"httprequest,omitempty"`
	Iptunnel                   string `json:"iptunnel,omitempty"`
	Lastresponse               string `json:"lastresponse,omitempty"`
	Monitorcurrentfailedprobes int    `json:"monitorcurrentfailedprobes,omitempty"`
	Monitorname                string `json:"monitorname,omitempty"`
	Monitortotalfailedprobes   int    `json:"monitortotalfailedprobes,omitempty"`
	Monitortotalprobes         int    `json:"monitortotalprobes,omitempty"`
	Monstatcode                int    `json:"monstatcode,omitempty"`
	Monstate                   string `json:"monstate,omitempty"`
	Name                       string `json:"name,omitempty"`
	Respcode                   string `json:"respcode,omitempty"`
	Responsetime               int    `json:"responsetime,omitempty"`
	Servicename                string `json:"servicename,omitempty"`
	Vservername                string `json:"vservername,omitempty"`
}

type Gslbparameter struct {
	Automaticconfigsync         string   `json:"automaticconfigsync,omitempty"`
	Builtin                     []string `json:"builtin,omitempty"`
	Dropldnsreq                 string   `json:"dropldnsreq,omitempty"`
	Feature                     string   `json:"feature,omitempty"`
	Flags                       int      `json:"flags,omitempty"`
	Gslbconfigsyncmonitor       string   `json:"gslbconfigsyncmonitor,omitempty"`
	Gslbsvcstatedelaytime       int      `json:"gslbsvcstatedelaytime,omitempty"`
	Gslbsyncinterval            int      `json:"gslbsyncinterval,omitempty"`
	Gslbsynclocfiles            string   `json:"gslbsynclocfiles,omitempty"`
	Gslbsyncmode                string   `json:"gslbsyncmode,omitempty"`
	Gslbsyncsaveconfigcommand   string   `json:"gslbsyncsaveconfigcommand,omitempty"`
	Incarnation                 int      `json:"incarnation,omitempty"`
	Ldnsentrytimeout            int      `json:"ldnsentrytimeout,omitempty"`
	Ldnsmask                    string   `json:"ldnsmask,omitempty"`
	Ldnsprobeorder              []string `json:"ldnsprobeorder,omitempty"`
	Mepkeepalivetimeout         int      `json:"mepkeepalivetimeout,omitempty"`
	Nextgenapiresource          string   `json:"_nextgenapiresource,omitempty"`
	Overridepersistencyfororder string   `json:"overridepersistencyfororder,omitempty"`
	Rtttolerance                int      `json:"rtttolerance,omitempty"`
	Svcstatelearningtime        int      `json:"svcstatelearningtime,omitempty"`
	Undefaction                 string   `json:"undefaction,omitempty"`
	V6ldnsmasklen               int      `json:"v6ldnsmasklen,omitempty"`
}

type GslbvserverSpilloverpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
}

type GslbvserverBinding struct {
	GslbvserverGslbdomainBinding             []interface{} `json:"gslbvserver_gslbdomain_binding,omitempty"`
	GslbvserverGslbserviceBinding            []interface{} `json:"gslbvserver_gslbservice_binding,omitempty"`
	GslbvserverGslbservicegroupBinding       []interface{} `json:"gslbvserver_gslbservicegroup_binding,omitempty"`
	GslbvserverGslbservicegroupmemberBinding []interface{} `json:"gslbvserver_gslbservicegroupmember_binding,omitempty"`
	GslbvserverLbpolicyBinding               []interface{} `json:"gslbvserver_lbpolicy_binding,omitempty"`
	GslbvserverSpilloverpolicyBinding        []interface{} `json:"gslbvserver_spilloverpolicy_binding,omitempty"`
	Name                                     string        `json:"name,omitempty"`
}

type GslbserviceDnsviewBinding struct {
	Servicename string `json:"servicename,omitempty"`
	Viewip      string `json:"viewip,omitempty"`
	Viewname    string `json:"viewname,omitempty"`
}

type GslbservicegroupGslbservicegroupmemberBinding struct {
	Delay                     int    `json:"delay,omitempty"`
	Graceful                  string `json:"graceful,omitempty"`
	Gslbthreshold             int    `json:"gslbthreshold,omitempty"`
	Hashid                    int    `json:"hashid,omitempty"`
	Ip                        string `json:"ip,omitempty"`
	Order                     int    `json:"order,omitempty"`
	Orderstr                  string `json:"orderstr,omitempty"`
	Port                      int    `json:"port,omitempty"`
	Preferredlocation         string `json:"preferredlocation,omitempty"`
	Publicip                  string `json:"publicip,omitempty"`
	Publicport                int    `json:"publicport,omitempty"`
	Servername                string `json:"servername,omitempty"`
	Servicegroupname          string `json:"servicegroupname,omitempty"`
	Siteprefix                string `json:"siteprefix,omitempty"`
	State                     string `json:"state,omitempty"`
	Statechangetimesec        string `json:"statechangetimesec,omitempty"`
	Svrstate                  string `json:"svrstate,omitempty"`
	Threshold                 string `json:"threshold,omitempty"`
	Tickssincelaststatechange int    `json:"tickssincelaststatechange,omitempty"`
	Trofsdelay                int    `json:"trofsdelay,omitempty"`
	Weight                    int    `json:"weight,omitempty"`
}

type GslbdomainGslbservicegroupBinding struct {
	Name             string `json:"name,omitempty"`
	Order            int    `json:"order,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
}

type Gslbdomain struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
}

type GslbvserverGslbservicegroupmemberBinding struct {
	Curstate           string `json:"curstate,omitempty"`
	Dynamicweight      string `json:"dynamicweight,omitempty"`
	Gslbthreshold      int    `json:"gslbthreshold,omitempty"`
	Ipaddress          string `json:"ipaddress,omitempty"`
	Name               string `json:"name,omitempty"`
	Order              int    `json:"order,omitempty"`
	Orderstr           string `json:"orderstr,omitempty"`
	Port               int    `json:"port,omitempty"`
	Preferredlocation  string `json:"preferredlocation,omitempty"`
	Servicegroupname   string `json:"servicegroupname,omitempty"`
	Servicetype        string `json:"servicetype,omitempty"`
	Sitepersistcookie  string `json:"sitepersistcookie,omitempty"`
	Svcsitepersistence string `json:"svcsitepersistence,omitempty"`
	Svreffgslbstate    string `json:"svreffgslbstate,omitempty"`
	Thresholdvalue     int    `json:"thresholdvalue,omitempty"`
	Weight             int    `json:"weight,omitempty"`
}

type GslbsiteGslbserviceBinding struct {
	Cnameentry  string `json:"cnameentry,omitempty"`
	Ipaddress   string `json:"ipaddress,omitempty"`
	Port        int    `json:"port,omitempty"`
	Servicename string `json:"servicename,omitempty"`
	Servicetype string `json:"servicetype,omitempty"`
	Sitename    string `json:"sitename,omitempty"`
	State       string `json:"state,omitempty"`
}

type GslbsiteGslbservicegroupmemberBinding struct {
	Ipaddress        string `json:"ipaddress,omitempty"`
	Port             int    `json:"port,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Servicetype      string `json:"servicetype,omitempty"`
	Sitename         string `json:"sitename,omitempty"`
	State            string `json:"state,omitempty"`
}

type GslbsiteBinding struct {
	GslbsiteGslbserviceBinding            []interface{} `json:"gslbsite_gslbservice_binding,omitempty"`
	GslbsiteGslbservicegroupBinding       []interface{} `json:"gslbsite_gslbservicegroup_binding,omitempty"`
	GslbsiteGslbservicegroupmemberBinding []interface{} `json:"gslbsite_gslbservicegroupmember_binding,omitempty"`
	Sitename                              string        `json:"sitename,omitempty"`
}

type GslbdomainGslbserviceBinding struct {
	Cnameentry       string `json:"cnameentry,omitempty"`
	Cumulativeweight int    `json:"cumulativeweight,omitempty"`
	Dynamicconfwt    int    `json:"dynamicconfwt,omitempty"`
	Gslbthreshold    int    `json:"gslbthreshold,omitempty"`
	Ipaddress        string `json:"ipaddress,omitempty"`
	Name             string `json:"name,omitempty"`
	Order            int    `json:"order,omitempty"`
	Port             int    `json:"port,omitempty"`
	Servicename      string `json:"servicename,omitempty"`
	Servicetype      string `json:"servicetype,omitempty"`
	State            string `json:"state,omitempty"`
	Svreffgslbstate  string `json:"svreffgslbstate,omitempty"`
	Vservername      string `json:"vservername,omitempty"`
	Weight           int    `json:"weight,omitempty"`
}

type GslbvserverGslbservicegroupBinding struct {
	Name             string `json:"name,omitempty"`
	Order            int    `json:"order,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
}

type GslbsiteGslbservicegroupBinding struct {
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Servicetype      string `json:"servicetype,omitempty"`
	Sitename         string `json:"sitename,omitempty"`
}

type GslbservicegroupLbmonitorBinding struct {
	Hashid           int    `json:"hashid,omitempty"`
	MonitorName      string `json:"monitor_name,omitempty"`
	Monstate         string `json:"monstate,omitempty"`
	Monweight        int    `json:"monweight,omitempty"`
	Order            int    `json:"order,omitempty"`
	Passive          bool   `json:"passive,omitempty"`
	Port             int    `json:"port,omitempty"`
	Publicip         string `json:"publicip,omitempty"`
	Publicport       int    `json:"publicport,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Siteprefix       string `json:"siteprefix,omitempty"`
	State            string `json:"state,omitempty"`
	Weight           int    `json:"weight,omitempty"`
}

type GslbvserverDomainBinding struct {
	Backupip         string `json:"backupip,omitempty"`
	Backupipflag     bool   `json:"backupipflag,omitempty"`
	CookieDomain     string `json:"cookie_domain,omitempty"`
	CookieDomainflag bool   `json:"cookie_domainflag,omitempty"`
	Cookietimeout    int    `json:"cookietimeout,omitempty"`
	Domainname       string `json:"domainname,omitempty"`
	Name             string `json:"name,omitempty"`
	Order            int    `json:"order,omitempty"`
	Sitedomainttl    int    `json:"sitedomainttl,omitempty"`
	Ttl              int    `json:"ttl,omitempty"`
}

type GslbdomainBinding struct {
	GslbdomainGslbserviceBinding            []interface{} `json:"gslbdomain_gslbservice_binding,omitempty"`
	GslbdomainGslbservicegroupBinding       []interface{} `json:"gslbdomain_gslbservicegroup_binding,omitempty"`
	GslbdomainGslbservicegroupmemberBinding []interface{} `json:"gslbdomain_gslbservicegroupmember_binding,omitempty"`
	GslbdomainGslbvserverBinding            []interface{} `json:"gslbdomain_gslbvserver_binding,omitempty"`
	GslbdomainLbmonitorBinding              []interface{} `json:"gslbdomain_lbmonitor_binding,omitempty"`
	Name                                    string        `json:"name,omitempty"`
}

type GslbserviceBinding struct {
	GslbserviceDnsviewBinding   []interface{} `json:"gslbservice_dnsview_binding,omitempty"`
	GslbserviceLbmonitorBinding []interface{} `json:"gslbservice_lbmonitor_binding,omitempty"`
	Servicename                 string        `json:"servicename,omitempty"`
}

type GslbserviceLbmonitorBinding struct {
	Failedprobes               int    `json:"failedprobes,omitempty"`
	Lastresponse               string `json:"lastresponse,omitempty"`
	MonitorName                string `json:"monitor_name,omitempty"`
	MonitorState               string `json:"monitor_state,omitempty"`
	Monitorcurrentfailedprobes int    `json:"monitorcurrentfailedprobes,omitempty"`
	Monitortotalfailedprobes   int    `json:"monitortotalfailedprobes,omitempty"`
	Monitortotalprobes         int    `json:"monitortotalprobes,omitempty"`
	Monstatcode                int    `json:"monstatcode,omitempty"`
	Monstate                   string `json:"monstate,omitempty"`
	Monstatparam1              int    `json:"monstatparam1,omitempty"`
	Monstatparam2              int    `json:"monstatparam2,omitempty"`
	Monstatparam3              int    `json:"monstatparam3,omitempty"`
	Responsetime               int    `json:"responsetime,omitempty"`
	Servicename                string `json:"servicename,omitempty"`
	Totalfailedprobes          int    `json:"totalfailedprobes,omitempty"`
	Weight                     int    `json:"weight,omitempty"`
}

type GslbdomainGslbservicegroupmemberBinding struct {
	Gslbthreshold    int    `json:"gslbthreshold,omitempty"`
	Ipaddress        string `json:"ipaddress,omitempty"`
	Name             string `json:"name,omitempty"`
	Order            int    `json:"order,omitempty"`
	Port             int    `json:"port,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Servicetype      string `json:"servicetype,omitempty"`
	Svreffgslbstate  string `json:"svreffgslbstate,omitempty"`
	Weight           int    `json:"weight,omitempty"`
}

type Gslbservice struct {
	Appflowlog                string  `json:"appflowlog,omitempty"`
	Cip                       string  `json:"cip,omitempty"`
	Cipheader                 string  `json:"cipheader,omitempty"`
	Clmonowner                int     `json:"clmonowner,omitempty"`
	Clmonview                 int     `json:"clmonview,omitempty"`
	Clttimeout                int     `json:"clttimeout,omitempty"`
	Cnameentry                string  `json:"cnameentry,omitempty"`
	Comment                   string  `json:"comment,omitempty"`
	Cookietimeout             int     `json:"cookietimeout,omitempty"`
	Count                     float64 `json:"__count,omitempty"`
	Downstateflush            string  `json:"downstateflush,omitempty"`
	Glsbsvchealthdescr        string  `json:"glsbsvchealthdescr,omitempty"`
	Gslb                      string  `json:"gslb,omitempty"`
	Gslbsvchealth             int     `json:"gslbsvchealth,omitempty"`
	Gslbsvcstats              int     `json:"gslbsvcstats,omitempty"`
	Gslbthreshold             int     `json:"gslbthreshold,omitempty"`
	Hashid                    int     `json:"hashid,omitempty"`
	Healthmonitor             string  `json:"healthmonitor,omitempty"`
	Ip                        string  `json:"ip,omitempty"`
	Ipaddress                 string  `json:"ipaddress,omitempty"`
	Maxaaausers               int     `json:"maxaaausers,omitempty"`
	Maxbandwidth              int     `json:"maxbandwidth,omitempty"`
	Maxclient                 int     `json:"maxclient,omitempty"`
	MonitorNameSvc            string  `json:"monitor_name_svc,omitempty"`
	MonitorState              string  `json:"monitor_state,omitempty"`
	Monstate                  string  `json:"monstate,omitempty"`
	Monthreshold              int     `json:"monthreshold,omitempty"`
	Naptrdomainttl            int     `json:"naptrdomainttl,omitempty"`
	Naptrorder                int     `json:"naptrorder,omitempty"`
	Naptrpreference           int     `json:"naptrpreference,omitempty"`
	Naptrreplacement          string  `json:"naptrreplacement,omitempty"`
	Naptrservices             string  `json:"naptrservices,omitempty"`
	Newname                   string  `json:"newname,omitempty"`
	Nextgenapiresource        string  `json:"_nextgenapiresource,omitempty"`
	Nodefaultbindings         string  `json:"nodefaultbindings,omitempty"`
	Port                      int     `json:"port,omitempty"`
	Preferredlocation         string  `json:"preferredlocation,omitempty"`
	Publicip                  string  `json:"publicip,omitempty"`
	Publicport                int     `json:"publicport,omitempty"`
	Servername                string  `json:"servername,omitempty"`
	Servicename               string  `json:"servicename,omitempty"`
	Servicetype               string  `json:"servicetype,omitempty"`
	Sitename                  string  `json:"sitename,omitempty"`
	Sitepersistence           string  `json:"sitepersistence,omitempty"`
	Siteprefix                string  `json:"siteprefix,omitempty"`
	State                     string  `json:"state,omitempty"`
	Statechangetimesec        string  `json:"statechangetimesec,omitempty"`
	Svreffgslbstate           string  `json:"svreffgslbstate,omitempty"`
	Svrstate                  string  `json:"svrstate,omitempty"`
	Svrtimeout                int     `json:"svrtimeout,omitempty"`
	Threshold                 string  `json:"threshold,omitempty"`
	Tickssincelaststatechange int     `json:"tickssincelaststatechange,omitempty"`
	Viewip                    string  `json:"viewip,omitempty"`
	Viewname                  string  `json:"viewname,omitempty"`
	Weight                    int     `json:"weight,omitempty"`
}

type Gslbconfig struct {
	Command    string `json:"command,omitempty"`
	Debug      bool   `json:"debug,omitempty"`
	Forcesync  string `json:"forcesync,omitempty"`
	Nowarn     bool   `json:"nowarn,omitempty"`
	Preview    bool   `json:"preview,omitempty"`
	Saveconfig bool   `json:"saveconfig,omitempty"`
}

type Gslbrunningconfig struct {
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Response           string `json:"response,omitempty"`
}

type Gslbsite struct {
	Backupparentlist       []string `json:"backupparentlist,omitempty"`
	Clip                   string   `json:"clip,omitempty"`
	Count                  float64  `json:"__count,omitempty"`
	Curbackupparentip      string   `json:"curbackupparentip,omitempty"`
	Metricexchange         string   `json:"metricexchange,omitempty"`
	Naptrreplacementsuffix string   `json:"naptrreplacementsuffix,omitempty"`
	Newname                string   `json:"newname,omitempty"`
	Nextgenapiresource     string   `json:"_nextgenapiresource,omitempty"`
	Nwmetricexchange       string   `json:"nwmetricexchange,omitempty"`
	Oldname                string   `json:"oldname,omitempty"`
	Parentsite             string   `json:"parentsite,omitempty"`
	Persistencemepstatus   string   `json:"persistencemepstatus,omitempty"`
	Publicclip             string   `json:"publicclip,omitempty"`
	Publicip               string   `json:"publicip,omitempty"`
	Sessionexchange        string   `json:"sessionexchange,omitempty"`
	Siteipaddress          string   `json:"siteipaddress,omitempty"`
	Sitename               string   `json:"sitename,omitempty"`
	Sitepassword           string   `json:"sitepassword,omitempty"`
	Sitestate              string   `json:"sitestate,omitempty"`
	Sitetype               string   `json:"sitetype,omitempty"`
	Status                 string   `json:"status,omitempty"`
	Triggermonitor         string   `json:"triggermonitor,omitempty"`
	Version                int      `json:"version,omitempty"`
}

type Gslbsyncstatus struct {
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Response           string `json:"response,omitempty"`
	Summary            bool   `json:"summary,omitempty"`
}

type Gslbldnsentry struct {
	Ipaddress string `json:"ipaddress,omitempty"`
}

type GslbvserverGslbserviceBinding struct {
	Cnameentry         string `json:"cnameentry,omitempty"`
	Cumulativeweight   int    `json:"cumulativeweight,omitempty"`
	Curstate           string `json:"curstate,omitempty"`
	Domainname         string `json:"domainname,omitempty"`
	Dynamicconfwt      int    `json:"dynamicconfwt,omitempty"`
	Gslbboundsvctype   string `json:"gslbboundsvctype,omitempty"`
	Gslbthreshold      int    `json:"gslbthreshold,omitempty"`
	Ipaddress          string `json:"ipaddress,omitempty"`
	Iscname            string `json:"iscname,omitempty"`
	Name               string `json:"name,omitempty"`
	Order              int    `json:"order,omitempty"`
	Orderstr           string `json:"orderstr,omitempty"`
	Port               int    `json:"port,omitempty"`
	Preferredlocation  string `json:"preferredlocation,omitempty"`
	Servicename        string `json:"servicename,omitempty"`
	Sitepersistcookie  string `json:"sitepersistcookie,omitempty"`
	Svcsitepersistence string `json:"svcsitepersistence,omitempty"`
	Svreffgslbstate    string `json:"svreffgslbstate,omitempty"`
	Thresholdvalue     int    `json:"thresholdvalue,omitempty"`
	Weight             int    `json:"weight,omitempty"`
}

type GslbdomainGslbvserverBinding struct {
	Backuplbmethod     string `json:"backuplbmethod,omitempty"`
	Cip                string `json:"cip,omitempty"`
	Customheaders      string `json:"customheaders,omitempty"`
	Dnsrecordtype      string `json:"dnsrecordtype,omitempty"`
	Dynamicweight      string `json:"dynamicweight,omitempty"`
	Edr                string `json:"edr,omitempty"`
	Lbmethod           string `json:"lbmethod,omitempty"`
	Mir                string `json:"mir,omitempty"`
	Name               string `json:"name,omitempty"`
	Netmask            string `json:"netmask,omitempty"`
	Persistenceid      int    `json:"persistenceid,omitempty"`
	Persistencetype    string `json:"persistencetype,omitempty"`
	Persistmask        string `json:"persistmask,omitempty"`
	Servicetype        string `json:"servicetype,omitempty"`
	Sitename           string `json:"sitename,omitempty"`
	Sitepersistence    string `json:"sitepersistence,omitempty"`
	Siteprefix         string `json:"siteprefix,omitempty"`
	State              string `json:"state,omitempty"`
	Statechangetimesec string `json:"statechangetimesec,omitempty"`
	V6netmasklen       int    `json:"v6netmasklen,omitempty"`
	V6persistmasklen   int    `json:"v6persistmasklen,omitempty"`
	Vservername        string `json:"vservername,omitempty"`
}
