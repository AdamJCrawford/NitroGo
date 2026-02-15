package models

// lb configuration structs
type LbvserverRewritepolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LbvserverFeopolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Lbmonbindings struct {
	Boundservicegroupsvrstate string  `json:"boundservicegroupsvrstate,omitempty"`
	Count                     float64 `json:"__count,omitempty"`
	Monitorname               string  `json:"monitorname,omitempty"`
	Nextgenapiresource        string  `json:"_nextgenapiresource,omitempty"`
	State                     string  `json:"state,omitempty"`
	TypeField                 string  `json:"type,omitempty"`
}

type LbvserverCsvserverBinding struct {
	Cachetype     string `json:"cachetype,omitempty"`
	Cachevserver  string `json:"cachevserver,omitempty"`
	Hits          int    `json:"hits,omitempty"`
	Name          string `json:"name,omitempty"`
	Order         int    `json:"order,omitempty"`
	Pipolicyhits  int    `json:"pipolicyhits,omitempty"`
	Policyname    string `json:"policyname,omitempty"`
	Policysubtype int    `json:"policysubtype,omitempty"`
	Priority      int    `json:"priority,omitempty"`
}

type LbmonitorBinding struct {
	LbmonitorMetricBinding     []interface{} `json:"lbmonitor_metric_binding,omitempty"`
	LbmonitorSslcertkeyBinding []interface{} `json:"lbmonitor_sslcertkey_binding,omitempty"`
	Monitorname                string        `json:"monitorname,omitempty"`
}

type LbmonbindingsBinding struct {
	LbmonbindingsGslbservicegroupBinding []interface{} `json:"lbmonbindings_gslbservicegroup_binding,omitempty"`
	LbmonbindingsServiceBinding          []interface{} `json:"lbmonbindings_service_binding,omitempty"`
	LbmonbindingsServicegroupBinding     []interface{} `json:"lbmonbindings_servicegroup_binding,omitempty"`
	Monitorname                          string        `json:"monitorname,omitempty"`
}

type LbvserverBotpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Lbroute6 struct {
	Count              float64 `json:"__count,omitempty"`
	Flags              string  `json:"flags,omitempty"`
	Gatewayname        string  `json:"gatewayname,omitempty"`
	Network            string  `json:"network,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Td                 int     `json:"td,omitempty"`
}

type LbpolicyGslbvserverBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LbvserverResponderpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LbmonbindingsServiceBinding struct {
	Ipaddress   string `json:"ipaddress,omitempty"`
	Monitorname string `json:"monitorname,omitempty"`
	Monsvcstate string `json:"monsvcstate,omitempty"`
	Port        int    `json:"port,omitempty"`
	Servicename string `json:"servicename,omitempty"`
	Servicetype string `json:"servicetype,omitempty"`
	Svrstate    string `json:"svrstate,omitempty"`
}

type LbvserverTransformpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Lbaction struct {
	Builtin            []string      `json:"builtin,omitempty"`
	Comment            string        `json:"comment,omitempty"`
	Count              float64       `json:"__count,omitempty"`
	Feature            string        `json:"feature,omitempty"`
	Hits               int           `json:"hits,omitempty"`
	Name               string        `json:"name,omitempty"`
	Newname            string        `json:"newname,omitempty"`
	Nextgenapiresource string        `json:"_nextgenapiresource,omitempty"`
	Referencecount     int           `json:"referencecount,omitempty"`
	TypeField          string        `json:"type,omitempty"`
	Undefhits          int           `json:"undefhits,omitempty"`
	Value              []interface{} `json:"value,omitempty"`
}

type LbvserverAuditsyslogpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LbglobalBinding struct {
	LbglobalLbpolicyBinding []interface{} `json:"lbglobal_lbpolicy_binding,omitempty"`
}

type LbvserverAppflowpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Lbroute struct {
	Count              float64 `json:"__count,omitempty"`
	Flags              string  `json:"flags,omitempty"`
	Gatewayname        string  `json:"gatewayname,omitempty"`
	Netmask            string  `json:"netmask,omitempty"`
	Network            string  `json:"network,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Td                 int     `json:"td,omitempty"`
}

type LbmonitorServiceBinding struct {
	DupState         string `json:"dup_state,omitempty"`
	DupWeight        int    `json:"dup_weight,omitempty"`
	Monitorname      string `json:"monitorname,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Servicename      string `json:"servicename,omitempty"`
	State            string `json:"state,omitempty"`
	Weight           int    `json:"weight,omitempty"`
}

type LbvserverAppfwpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LbvserverCachepolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LbvserverBinding struct {
	LbvserverAnalyticsprofileBinding                 []interface{} `json:"lbvserver_analyticsprofile_binding,omitempty"`
	LbvserverAppflowpolicyBinding                    []interface{} `json:"lbvserver_appflowpolicy_binding,omitempty"`
	LbvserverAppfwpolicyBinding                      []interface{} `json:"lbvserver_appfwpolicy_binding,omitempty"`
	LbvserverAppqoepolicyBinding                     []interface{} `json:"lbvserver_appqoepolicy_binding,omitempty"`
	LbvserverAuditnslogpolicyBinding                 []interface{} `json:"lbvserver_auditnslogpolicy_binding,omitempty"`
	LbvserverAuditsyslogpolicyBinding                []interface{} `json:"lbvserver_auditsyslogpolicy_binding,omitempty"`
	LbvserverAuthorizationpolicyBinding              []interface{} `json:"lbvserver_authorizationpolicy_binding,omitempty"`
	LbvserverBotpolicyBinding                        []interface{} `json:"lbvserver_botpolicy_binding,omitempty"`
	LbvserverCachepolicyBinding                      []interface{} `json:"lbvserver_cachepolicy_binding,omitempty"`
	LbvserverCmppolicyBinding                        []interface{} `json:"lbvserver_cmppolicy_binding,omitempty"`
	LbvserverContentinspectionpolicyBinding          []interface{} `json:"lbvserver_contentinspectionpolicy_binding,omitempty"`
	LbvserverCsvserverBinding                        []interface{} `json:"lbvserver_csvserver_binding,omitempty"`
	LbvserverDnspolicy64Binding                      []interface{} `json:"lbvserver_dnspolicy64_binding,omitempty"`
	LbvserverFeopolicyBinding                        []interface{} `json:"lbvserver_feopolicy_binding,omitempty"`
	LbvserverLbpolicyBinding                         []interface{} `json:"lbvserver_lbpolicy_binding,omitempty"`
	LbvserverResponderpolicyBinding                  []interface{} `json:"lbvserver_responderpolicy_binding,omitempty"`
	LbvserverRewritepolicyBinding                    []interface{} `json:"lbvserver_rewritepolicy_binding,omitempty"`
	LbvserverServiceBinding                          []interface{} `json:"lbvserver_service_binding,omitempty"`
	LbvserverServicegroupBinding                     []interface{} `json:"lbvserver_servicegroup_binding,omitempty"`
	LbvserverServicegroupmemberBinding               []interface{} `json:"lbvserver_servicegroupmember_binding,omitempty"`
	LbvserverSpilloverpolicyBinding                  []interface{} `json:"lbvserver_spilloverpolicy_binding,omitempty"`
	LbvserverTmtrafficpolicyBinding                  []interface{} `json:"lbvserver_tmtrafficpolicy_binding,omitempty"`
	LbvserverTransformpolicyBinding                  []interface{} `json:"lbvserver_transformpolicy_binding,omitempty"`
	LbvserverVideooptimizationdetectionpolicyBinding []interface{} `json:"lbvserver_videooptimizationdetectionpolicy_binding,omitempty"`
	LbvserverVideooptimizationpacingpolicyBinding    []interface{} `json:"lbvserver_videooptimizationpacingpolicy_binding,omitempty"`
	Name                                             string        `json:"name,omitempty"`
}

type LbpolicylabelPolicybindingBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LbvserverContentinspectionpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LbglobalLbpolicyBinding struct {
	Flowtype               int    `json:"flowtype,omitempty"`
	Globalbindtype         string `json:"globalbindtype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Numpol                 int    `json:"numpol,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
}

type LbvserverServiceBinding struct {
	Cookieipport      string `json:"cookieipport,omitempty"`
	Cookiename        string `json:"cookiename,omitempty"`
	Curstate          string `json:"curstate,omitempty"`
	Dynamicweight     int    `json:"dynamicweight,omitempty"`
	Ipv46             string `json:"ipv46,omitempty"`
	Name              string `json:"name,omitempty"`
	Order             int    `json:"order,omitempty"`
	Orderstr          string `json:"orderstr,omitempty"`
	Port              int    `json:"port,omitempty"`
	Preferredlocation string `json:"preferredlocation,omitempty"`
	Servicegroupname  string `json:"servicegroupname,omitempty"`
	Servicename       string `json:"servicename,omitempty"`
	Servicetype       string `json:"servicetype,omitempty"`
	Vserverid         string `json:"vserverid,omitempty"`
	Vsvrbindsvcip     string `json:"vsvrbindsvcip,omitempty"`
	Vsvrbindsvcport   int    `json:"vsvrbindsvcport,omitempty"`
	Weight            int    `json:"weight,omitempty"`
}

type LbvserverAppqoepolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LbpolicyLbglobalBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Lbgroup struct {
	Backuppersistencetimeout int     `json:"backuppersistencetimeout,omitempty"`
	Cookiedomain             string  `json:"cookiedomain,omitempty"`
	Cookiename               string  `json:"cookiename,omitempty"`
	Count                    float64 `json:"__count,omitempty"`
	Mastervserver            string  `json:"mastervserver,omitempty"`
	Name                     string  `json:"name,omitempty"`
	Newname                  string  `json:"newname,omitempty"`
	Nextgenapiresource       string  `json:"_nextgenapiresource,omitempty"`
	Persistencebackup        string  `json:"persistencebackup,omitempty"`
	Persistencetype          string  `json:"persistencetype,omitempty"`
	Persistmask              string  `json:"persistmask,omitempty"`
	Rule                     string  `json:"rule,omitempty"`
	Td                       int     `json:"td,omitempty"`
	Timeout                  int     `json:"timeout,omitempty"`
	Usevserverpersistency    string  `json:"usevserverpersistency,omitempty"`
	V6persistmasklen         int     `json:"v6persistmasklen,omitempty"`
}

type LbpolicylabelLbpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Lbsipparameters struct {
	Addrportvip         string   `json:"addrportvip,omitempty"`
	Builtin             []string `json:"builtin,omitempty"`
	Feature             string   `json:"feature,omitempty"`
	Nextgenapiresource  string   `json:"_nextgenapiresource,omitempty"`
	Retrydur            int      `json:"retrydur,omitempty"`
	Rnatdstport         int      `json:"rnatdstport,omitempty"`
	Rnatsecuredstport   int      `json:"rnatsecuredstport,omitempty"`
	Rnatsecuresrcport   int      `json:"rnatsecuresrcport,omitempty"`
	Rnatsrcport         int      `json:"rnatsrcport,omitempty"`
	Sip503ratethreshold int      `json:"sip503ratethreshold,omitempty"`
}

type LbvserverAnalyticsprofileBinding struct {
	Analyticsprofile string `json:"analyticsprofile,omitempty"`
	Name             string `json:"name,omitempty"`
	Order            int    `json:"order,omitempty"`
}

type LbwlmLbvserverBinding struct {
	Vservername string `json:"vservername,omitempty"`
	Wlmname     string `json:"wlmname,omitempty"`
}

type LbvserverCmppolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LBVServer struct {
	Activeservices                     int           `json:"activeservices,omitempty"`
	Adfsproxyprofile                   string        `json:"adfsproxyprofile,omitempty"`
	Apiprofile                         string        `json:"apiprofile,omitempty"`
	Appflowlog                         string        `json:"appflowlog,omitempty"`
	Authentication                     string        `json:"authentication,omitempty"`
	Authenticationhost                 string        `json:"authenticationhost,omitempty"`
	Authn401                           string        `json:"authn401,omitempty"`
	Authnprofile                       string        `json:"authnprofile,omitempty"`
	Authnvsname                        string        `json:"authnvsname,omitempty"`
	Backuplbmethod                     string        `json:"backuplbmethod,omitempty"`
	Backuppersistencetimeout           int           `json:"backuppersistencetimeout,omitempty"`
	Backupvserver                      string        `json:"backupvserver,omitempty"`
	Backupvserverstatus                string        `json:"backupvserverstatus,omitempty"`
	Bindpoint                          string        `json:"bindpoint,omitempty"`
	Bypassaaaa                         string        `json:"bypassaaaa,omitempty"`
	Cacheable                          string        `json:"cacheable,omitempty"`
	Cachevserver                       string        `json:"cachevserver,omitempty"`
	Clttimeout                         int           `json:"clttimeout,omitempty"`
	Comment                            string        `json:"comment,omitempty"`
	Connfailover                       string        `json:"connfailover,omitempty"`
	Consolidatedlconn                  string        `json:"consolidatedlconn,omitempty"`
	Consolidatedlconngbl               string        `json:"consolidatedlconngbl,omitempty"`
	Cookiedomain                       string        `json:"cookiedomain,omitempty"`
	Cookiename                         string        `json:"cookiename,omitempty"`
	Count                              float64       `json:"__count,omitempty"`
	Currentactiveorder                 string        `json:"currentactiveorder,omitempty"`
	Curstate                           string        `json:"curstate,omitempty"`
	Datalength                         int           `json:"datalength,omitempty"`
	Dataoffset                         int           `json:"dataoffset,omitempty"`
	Dbprofilename                      string        `json:"dbprofilename,omitempty"`
	Dbslb                              string        `json:"dbslb,omitempty"`
	Disableprimaryondown               string        `json:"disableprimaryondown,omitempty"`
	Dns64                              string        `json:"dns64,omitempty"`
	Dnsoverhttps                       string        `json:"dnsoverhttps,omitempty"`
	Dnsprofilename                     string        `json:"dnsprofilename,omitempty"`
	Dnsvservername                     string        `json:"dnsvservername,omitempty"`
	Domain                             string        `json:"domain,omitempty"`
	Downstateflush                     string        `json:"downstateflush,omitempty"`
	Effectivestate                     string        `json:"effectivestate,omitempty"`
	Groupname                          string        `json:"groupname,omitempty"`
	Gt2gb                              string        `json:"gt2gb,omitempty"`
	Hashlength                         int           `json:"hashlength,omitempty"`
	Health                             int           `json:"health,omitempty"`
	Healththreshold                    int           `json:"healththreshold,omitempty"`
	Homepage                           string        `json:"homepage,omitempty"`
	Httpprofilename                    string        `json:"httpprofilename,omitempty"`
	Httpsredirecturl                   string        `json:"httpsredirecturl,omitempty"`
	Icmpvsrresponse                    string        `json:"icmpvsrresponse,omitempty"`
	Insertvserveripport                string        `json:"insertvserveripport,omitempty"`
	Ipmapping                          string        `json:"ipmapping,omitempty"`
	Ipmask                             string        `json:"ipmask,omitempty"`
	Ippattern                          string        `json:"ippattern,omitempty"`
	Ipset                              string        `json:"ipset,omitempty"`
	Ipv46                              string        `json:"ipv46,omitempty"`
	Isgslb                             bool          `json:"isgslb,omitempty"`
	L2conn                             string        `json:"l2conn,omitempty"`
	Lbmethod                           string        `json:"lbmethod,omitempty"`
	Lbprofilename                      string        `json:"lbprofilename,omitempty"`
	Lbrrreason                         int           `json:"lbrrreason,omitempty"`
	Listenpolicy                       string        `json:"listenpolicy,omitempty"`
	Listenpriority                     int           `json:"listenpriority,omitempty"`
	M                                  string        `json:"m,omitempty"`
	Macmoderetainvlan                  string        `json:"macmoderetainvlan,omitempty"`
	MapField                           string        `json:"map,omitempty"`
	Maxautoscalemembers                int           `json:"maxautoscalemembers,omitempty"`
	Minautoscalemembers                int           `json:"minautoscalemembers,omitempty"`
	Mssqlserverversion                 string        `json:"mssqlserverversion,omitempty"`
	Mysqlcharacterset                  int           `json:"mysqlcharacterset,omitempty"`
	Mysqlprotocolversion               int           `json:"mysqlprotocolversion,omitempty"`
	Mysqlservercapabilities            int           `json:"mysqlservercapabilities,omitempty"`
	Mysqlserverversion                 string        `json:"mysqlserverversion,omitempty"`
	Name                               string        `json:"name,omitempty"`
	Netmask                            string        `json:"netmask,omitempty"`
	Netprofile                         string        `json:"netprofile,omitempty"`
	Newname                            string        `json:"newname,omitempty"`
	Newservicerequest                  int           `json:"newservicerequest,omitempty"`
	Newservicerequestincrementinterval int           `json:"newservicerequestincrementinterval,omitempty"`
	Newservicerequestunit              string        `json:"newservicerequestunit,omitempty"`
	Nextgenapiresource                 string        `json:"_nextgenapiresource,omitempty"`
	Ngname                             string        `json:"ngname,omitempty"`
	Nodefaultbindings                  string        `json:"nodefaultbindings,omitempty"`
	Oracleserverversion                string        `json:"oracleserverversion,omitempty"`
	Order                              int           `json:"order,omitempty"`
	Orderthreshold                     int           `json:"orderthreshold,omitempty"`
	Persistavpno                       []interface{} `json:"persistavpno,omitempty"`
	Persistencebackup                  string        `json:"persistencebackup,omitempty"`
	Persistencetype                    string        `json:"persistencetype,omitempty"`
	Persistmask                        string        `json:"persistmask,omitempty"`
	Port                               int           `json:"port,omitempty"`
	Precedence                         string        `json:"precedence,omitempty"`
	Probeport                          int           `json:"probeport,omitempty"`
	Probeprotocol                      string        `json:"probeprotocol,omitempty"`
	Probesuccessresponsecode           string        `json:"probesuccessresponsecode,omitempty"`
	Processlocal                       string        `json:"processlocal,omitempty"`
	Push                               string        `json:"push,omitempty"`
	Pushlabel                          string        `json:"pushlabel,omitempty"`
	Pushmulticlients                   string        `json:"pushmulticlients,omitempty"`
	Pushvserver                        string        `json:"pushvserver,omitempty"`
	Quicbridgeprofilename              string        `json:"quicbridgeprofilename,omitempty"`
	Quicprofilename                    string        `json:"quicprofilename,omitempty"`
	Range                              int           `json:"range,omitempty"`
	Recursionavailable                 string        `json:"recursionavailable,omitempty"`
	Redirect                           string        `json:"redirect,omitempty"`
	Redirectfromport                   int           `json:"redirectfromport,omitempty"`
	Redirectportrewrite                string        `json:"redirectportrewrite,omitempty"`
	Redirurl                           string        `json:"redirurl,omitempty"`
	Redirurlflags                      bool          `json:"redirurlflags,omitempty"`
	Resrule                            string        `json:"resrule,omitempty"`
	Retainconnectionsoncluster         string        `json:"retainconnectionsoncluster,omitempty"`
	Rhistate                           string        `json:"rhistate,omitempty"`
	Rtspnat                            string        `json:"rtspnat,omitempty"`
	Rule                               string        `json:"rule,omitempty"`
	Ruletype                           int           `json:"ruletype,omitempty"`
	Servicename                        string        `json:"servicename,omitempty"`
	Servicetype                        string        `json:"servicetype,omitempty"`
	Sessionless                        string        `json:"sessionless,omitempty"`
	Skippersistency                    string        `json:"skippersistency,omitempty"`
	Sobackupaction                     string        `json:"sobackupaction,omitempty"`
	Somethod                           string        `json:"somethod,omitempty"`
	Sopersistence                      string        `json:"sopersistence,omitempty"`
	Sopersistencetimeout               int           `json:"sopersistencetimeout,omitempty"`
	Sothreshold                        int           `json:"sothreshold,omitempty"`
	State                              string        `json:"state,omitempty"`
	Statechangetimemsec                int           `json:"statechangetimemsec,omitempty"`
	Statechangetimesec                 string        `json:"statechangetimesec,omitempty"`
	Statechangetimeseconds             int           `json:"statechangetimeseconds,omitempty"`
	Status                             int           `json:"status,omitempty"`
	Tcpprobeport                       int           `json:"tcpprobeport,omitempty"`
	Tcpprofilename                     string        `json:"tcpprofilename,omitempty"`
	Td                                 int           `json:"td,omitempty"`
	Thresholdvalue                     int           `json:"thresholdvalue,omitempty"`
	Tickssincelaststatechange          int           `json:"tickssincelaststatechange,omitempty"`
	Timeout                            int           `json:"timeout,omitempty"`
	Toggleorder                        string        `json:"toggleorder,omitempty"`
	Tosid                              int           `json:"tosid,omitempty"`
	Totalservices                      int           `json:"totalservices,omitempty"`
	Trofspersistence                   string        `json:"trofspersistence,omitempty"`
	TypeField                          string        `json:"type,omitempty"`
	V6netmasklen                       int           `json:"v6netmasklen,omitempty"`
	V6persistmasklen                   int           `json:"v6persistmasklen,omitempty"`
	Value                              string        `json:"value,omitempty"`
	Version                            int           `json:"version,omitempty"`
	Vipheader                          string        `json:"vipheader,omitempty"`
	Vsvrdynconnsothreshold             int           `json:"vsvrdynconnsothreshold,omitempty"`
	Weight                             int           `json:"weight,omitempty"`
}

type LbvserverTmtrafficpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Lbmonitor struct {
	Acctapplicationid                []interface{} `json:"acctapplicationid,omitempty"`
	Action                           string        `json:"action,omitempty"`
	Alertretries                     int           `json:"alertretries,omitempty"`
	Application                      string        `json:"application,omitempty"`
	Attribute                        string        `json:"attribute,omitempty"`
	Authapplicationid                []interface{} `json:"authapplicationid,omitempty"`
	Basedn                           string        `json:"basedn,omitempty"`
	Binddn                           string        `json:"binddn,omitempty"`
	Count                            float64       `json:"__count,omitempty"`
	Customheaders                    string        `json:"customheaders,omitempty"`
	Database                         string        `json:"database,omitempty"`
	Destip                           string        `json:"destip,omitempty"`
	Destport                         int           `json:"destport,omitempty"`
	Deviation                        int           `json:"deviation,omitempty"`
	Dispatcherip                     string        `json:"dispatcherip,omitempty"`
	Dispatcherport                   int           `json:"dispatcherport,omitempty"`
	Domain                           string        `json:"domain,omitempty"`
	Downtime                         int           `json:"downtime,omitempty"`
	DupState                         string        `json:"dup_state,omitempty"`
	DupWeight                        int           `json:"dup_weight,omitempty"`
	Dynamicinterval                  int           `json:"dynamicinterval,omitempty"`
	Dynamicresponsetimeout           int           `json:"dynamicresponsetimeout,omitempty"`
	Evalrule                         string        `json:"evalrule,omitempty"`
	Failureretries                   int           `json:"failureretries,omitempty"`
	Filename                         string        `json:"filename,omitempty"`
	Filter                           string        `json:"filter,omitempty"`
	Firmwarerevision                 int           `json:"firmwarerevision,omitempty"`
	Group                            string        `json:"group,omitempty"`
	Grpchealthcheck                  string        `json:"grpchealthcheck,omitempty"`
	Grpcservicename                  string        `json:"grpcservicename,omitempty"`
	Grpcstatuscode                   []interface{} `json:"grpcstatuscode,omitempty"`
	Hostipaddress                    string        `json:"hostipaddress,omitempty"`
	Hostname                         string        `json:"hostname,omitempty"`
	Httprequest                      string        `json:"httprequest,omitempty"`
	Inbandsecurityid                 string        `json:"inbandsecurityid,omitempty"`
	Interval                         int           `json:"interval,omitempty"`
	Ipaddress                        []string      `json:"ipaddress,omitempty"`
	Iptunnel                         string        `json:"iptunnel,omitempty"`
	Kcdaccount                       string        `json:"kcdaccount,omitempty"`
	Lasversion                       string        `json:"lasversion,omitempty"`
	Logonpointname                   string        `json:"logonpointname,omitempty"`
	Lrtm                             string        `json:"lrtm,omitempty"`
	Lrtmconf                         int           `json:"lrtmconf,omitempty"`
	Lrtmconfstr                      string        `json:"lrtmconfstr,omitempty"`
	Maxforwards                      int           `json:"maxforwards,omitempty"`
	Metric                           string        `json:"metric,omitempty"`
	Metrictable                      string        `json:"metrictable,omitempty"`
	Metricthreshold                  int           `json:"metricthreshold,omitempty"`
	Metricweight                     int           `json:"metricweight,omitempty"`
	Monitorname                      string        `json:"monitorname,omitempty"`
	Mqttclientidentifier             string        `json:"mqttclientidentifier,omitempty"`
	Mqttversion                      int           `json:"mqttversion,omitempty"`
	Mssqlprotocolversion             string        `json:"mssqlprotocolversion,omitempty"`
	Multimetrictable                 []string      `json:"multimetrictable,omitempty"`
	Netprofile                       string        `json:"netprofile,omitempty"`
	Nextgenapiresource               string        `json:"_nextgenapiresource,omitempty"`
	Oraclesid                        string        `json:"oraclesid,omitempty"`
	Originhost                       string        `json:"originhost,omitempty"`
	Originrealm                      string        `json:"originrealm,omitempty"`
	Password                         string        `json:"password,omitempty"`
	Productname                      string        `json:"productname,omitempty"`
	Query                            string        `json:"query,omitempty"`
	Querytype                        string        `json:"querytype,omitempty"`
	Radaccountsession                string        `json:"radaccountsession,omitempty"`
	Radaccounttype                   int           `json:"radaccounttype,omitempty"`
	Radapn                           string        `json:"radapn,omitempty"`
	Radframedip                      string        `json:"radframedip,omitempty"`
	Radkey                           string        `json:"radkey,omitempty"`
	Radmsisdn                        string        `json:"radmsisdn,omitempty"`
	Radnasid                         string        `json:"radnasid,omitempty"`
	Radnasip                         string        `json:"radnasip,omitempty"`
	Recv                             string        `json:"recv,omitempty"`
	Respcode                         []string      `json:"respcode,omitempty"`
	Resptimeout                      int           `json:"resptimeout,omitempty"`
	Resptimeoutthresh                int           `json:"resptimeoutthresh,omitempty"`
	Retries                          int           `json:"retries,omitempty"`
	Reverse                          string        `json:"reverse,omitempty"`
	Rtsprequest                      string        `json:"rtsprequest,omitempty"`
	Scriptargs                       string        `json:"scriptargs,omitempty"`
	Scriptname                       string        `json:"scriptname,omitempty"`
	Secondarypassword                string        `json:"secondarypassword,omitempty"`
	Secure                           string        `json:"secure,omitempty"`
	Secureargs                       string        `json:"secureargs,omitempty"`
	Send                             string        `json:"send,omitempty"`
	Servicegroupname                 string        `json:"servicegroupname,omitempty"`
	Servicename                      string        `json:"servicename,omitempty"`
	Sipmethod                        string        `json:"sipmethod,omitempty"`
	Sipreguri                        string        `json:"sipreguri,omitempty"`
	Sipuri                           string        `json:"sipuri,omitempty"`
	Sitepath                         string        `json:"sitepath,omitempty"`
	Snmpcommunity                    string        `json:"snmpcommunity,omitempty"`
	Snmpoid                          string        `json:"Snmpoid,omitempty"`
	Snmpthreshold                    string        `json:"snmpthreshold,omitempty"`
	Snmpversion                      string        `json:"snmpversion,omitempty"`
	Sqlquery                         string        `json:"sqlquery,omitempty"`
	Sslprofile                       string        `json:"sslprofile,omitempty"`
	State                            string        `json:"state,omitempty"`
	Storedb                          string        `json:"storedb,omitempty"`
	Storefrontacctservice            string        `json:"storefrontacctservice,omitempty"`
	Storefrontcheckbackendservices   string        `json:"storefrontcheckbackendservices,omitempty"`
	Storename                        string        `json:"storename,omitempty"`
	Successretries                   int           `json:"successretries,omitempty"`
	Supportedvendorids               []interface{} `json:"supportedvendorids,omitempty"`
	Tos                              string        `json:"tos,omitempty"`
	Tosid                            int           `json:"tosid,omitempty"`
	Transparent                      string        `json:"transparent,omitempty"`
	Trofscode                        int           `json:"trofscode,omitempty"`
	Trofsstring                      string        `json:"trofsstring,omitempty"`
	TypeField                        string        `json:"type,omitempty"`
	Units1                           string        `json:"units1,omitempty"`
	Units2                           string        `json:"units2,omitempty"`
	Units3                           string        `json:"units3,omitempty"`
	Units4                           string        `json:"units4,omitempty"`
	Username                         string        `json:"username,omitempty"`
	Validatecred                     string        `json:"validatecred,omitempty"`
	Vendorid                         int           `json:"vendorid,omitempty"`
	Vendorspecificacctapplicationids []interface{} `json:"vendorspecificacctapplicationids,omitempty"`
	Vendorspecificauthapplicationids []interface{} `json:"vendorspecificauthapplicationids,omitempty"`
	Vendorspecificvendorid           int           `json:"vendorspecificvendorid,omitempty"`
	Weight                           int           `json:"weight,omitempty"`
}

type Lbprofile struct {
	Adccookieattributewarningmsg  string  `json:"adccookieattributewarningmsg,omitempty"`
	Computedadccookieattribute    string  `json:"computedadccookieattribute,omitempty"`
	Cookiepassphrase              string  `json:"cookiepassphrase,omitempty"`
	Count                         float64 `json:"__count,omitempty"`
	Dbslb                         string  `json:"dbslb,omitempty"`
	Httponlycookieflag            string  `json:"httponlycookieflag,omitempty"`
	Lbhashalgorithm               string  `json:"lbhashalgorithm,omitempty"`
	Lbhashalgowinsize             int     `json:"lbhashalgowinsize,omitempty"`
	Lbhashfingers                 int     `json:"lbhashfingers,omitempty"`
	Lbprofilename                 string  `json:"lbprofilename,omitempty"`
	Literaladccookieattribute     string  `json:"literaladccookieattribute,omitempty"`
	Nextgenapiresource            string  `json:"_nextgenapiresource,omitempty"`
	Processlocal                  string  `json:"processlocal,omitempty"`
	Proximityfromself             string  `json:"proximityfromself,omitempty"`
	Storemqttclientidandusername  string  `json:"storemqttclientidandusername,omitempty"`
	Useencryptedpersistencecookie string  `json:"useencryptedpersistencecookie,omitempty"`
	Usesecuredpersistencecookie   string  `json:"usesecuredpersistencecookie,omitempty"`
	Vsvrcount                     int     `json:"vsvrcount,omitempty"`
}

type LbvserverServicegroupBinding struct {
	Name             string `json:"name,omitempty"`
	Order            int    `json:"order,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Servicename      string `json:"servicename,omitempty"`
	Weight           int    `json:"weight,omitempty"`
}

type LbwlmBinding struct {
	LbwlmLbvserverBinding []interface{} `json:"lbwlm_lbvserver_binding,omitempty"`
	Wlmname               string        `json:"wlmname,omitempty"`
}

type LbmonbindingsGslbservicegroupBinding struct {
	Boundservicegroupsvrstate string `json:"boundservicegroupsvrstate,omitempty"`
	Monitorname               string `json:"monitorname,omitempty"`
	Monstate                  string `json:"monstate,omitempty"`
	Servicegroupname          string `json:"servicegroupname,omitempty"`
	Servicetype               string `json:"servicetype,omitempty"`
}

type Lbpolicy struct {
	Action             string   `json:"action,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	Comment            string   `json:"comment,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	Logaction          string   `json:"logaction,omitempty"`
	Name               string   `json:"name,omitempty"`
	Newname            string   `json:"newname,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Rule               string   `json:"rule,omitempty"`
	Undefaction        string   `json:"undefaction,omitempty"`
	Undefhits          int      `json:"undefhits,omitempty"`
}

type LbmonitorServicegroupBinding struct {
	DupState         string `json:"dup_state,omitempty"`
	DupWeight        int    `json:"dup_weight,omitempty"`
	Monitorname      string `json:"monitorname,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Servicename      string `json:"servicename,omitempty"`
	State            string `json:"state,omitempty"`
	Weight           int    `json:"weight,omitempty"`
}

type LbvserverSpilloverpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Lbwlm struct {
	Count              float64 `json:"__count,omitempty"`
	Ipaddress          string  `json:"ipaddress,omitempty"`
	Katimeout          int     `json:"katimeout,omitempty"`
	Lbuid              string  `json:"lbuid,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Port               int     `json:"port,omitempty"`
	Secure             string  `json:"secure,omitempty"`
	State              string  `json:"state,omitempty"`
	Wlmname            string  `json:"wlmname,omitempty"`
}

type LbgroupLbvserverBinding struct {
	Name        string `json:"name,omitempty"`
	Vservername string `json:"vservername,omitempty"`
}

type LbmetrictableBinding struct {
	LbmetrictableMetricBinding []interface{} `json:"lbmetrictable_metric_binding,omitempty"`
	Metrictable                string        `json:"metrictable,omitempty"`
}

type Lbpersistentsessions struct {
	Cnamepersparam       string  `json:"cnamepersparam,omitempty"`
	Count                float64 `json:"__count,omitempty"`
	Destip               string  `json:"destip,omitempty"`
	Destipv6             string  `json:"destipv6,omitempty"`
	Destport             int     `json:"destport,omitempty"`
	Flags                bool    `json:"flags,omitempty"`
	Nextgenapiresource   string  `json:"_nextgenapiresource,omitempty"`
	Nodeid               int     `json:"nodeid,omitempty"`
	Persistenceparam     string  `json:"persistenceparam,omitempty"`
	Persistenceparameter string  `json:"persistenceparameter,omitempty"`
	Referencecount       int     `json:"referencecount,omitempty"`
	Srcip                string  `json:"srcip,omitempty"`
	Srcipv6              string  `json:"srcipv6,omitempty"`
	Timeout              int     `json:"timeout,omitempty"`
	TypeField            int     `json:"type,omitempty"`
	Typestring           string  `json:"typestring,omitempty"`
	Vserver              string  `json:"vserver,omitempty"`
	Vservername          string  `json:"vservername,omitempty"`
}

type LbpolicyLbvserverBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LbvserverAuditnslogpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LbpolicylabelBinding struct {
	Labelname                         string        `json:"labelname,omitempty"`
	LbpolicylabelLbpolicyBinding      []interface{} `json:"lbpolicylabel_lbpolicy_binding,omitempty"`
	LbpolicylabelPolicybindingBinding []interface{} `json:"lbpolicylabel_policybinding_binding,omitempty"`
}

type LbmetrictableMetricBinding struct {
	Metric      string `json:"metric,omitempty"`
	Metrictable string `json:"metrictable,omitempty"`
	Metrictype  string `json:"metrictype,omitempty"`
	Snmpoid     string `json:"Snmpoid,omitempty"`
}

type LbvserverDnspolicy64Binding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LbvserverVideooptimizationpacingpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LbpolicyLbpolicylabelBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LbmonbindingsServicegroupBinding struct {
	Boundservicegroupsvrstate string `json:"boundservicegroupsvrstate,omitempty"`
	Monitorname               string `json:"monitorname,omitempty"`
	Monstate                  string `json:"monstate,omitempty"`
	Servicegroupname          string `json:"servicegroupname,omitempty"`
	Servicetype               string `json:"servicetype,omitempty"`
}

type LbmonitorSslcertkeyBinding struct {
	Ca          bool   `json:"ca,omitempty"`
	Certkeyname string `json:"certkeyname,omitempty"`
	Crlcheck    string `json:"crlcheck,omitempty"`
	Monitorname string `json:"monitorname,omitempty"`
	Ocspcheck   string `json:"ocspcheck,omitempty"`
}

type LbvserverAuthorizationpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LbvserverVideooptimizationdetectionpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Lbpolicylabel struct {
	Comment                string  `json:"comment,omitempty"`
	Count                  float64 `json:"__count,omitempty"`
	Gotopriorityexpression string  `json:"gotopriorityexpression,omitempty"`
	Hits                   int     `json:"hits,omitempty"`
	InvokeLabelname        string  `json:"invoke_labelname,omitempty"`
	Labelname              string  `json:"labelname,omitempty"`
	Labeltype              string  `json:"labeltype,omitempty"`
	Newname                string  `json:"newname,omitempty"`
	Nextgenapiresource     string  `json:"_nextgenapiresource,omitempty"`
	Numpol                 int     `json:"numpol,omitempty"`
	Policylabeltype        string  `json:"policylabeltype,omitempty"`
}

type LbpolicyBinding struct {
	LbpolicyGslbvserverBinding   []interface{} `json:"lbpolicy_gslbvserver_binding,omitempty"`
	LbpolicyLbglobalBinding      []interface{} `json:"lbpolicy_lbglobal_binding,omitempty"`
	LbpolicyLbpolicylabelBinding []interface{} `json:"lbpolicy_lbpolicylabel_binding,omitempty"`
	LbpolicyLbvserverBinding     []interface{} `json:"lbpolicy_lbvserver_binding,omitempty"`
	Name                         string        `json:"name,omitempty"`
}

type LbmonitorMetricBinding struct {
	Metric          string `json:"metric,omitempty"`
	MetricUnit      string `json:"metric_unit,omitempty"`
	Metrictable     string `json:"metrictable,omitempty"`
	Metricthreshold int    `json:"metricthreshold,omitempty"`
	Metricweight    int    `json:"metricweight,omitempty"`
	Monitorname     string `json:"monitorname,omitempty"`
}

type LbvserverLbpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type LbgroupBinding struct {
	LbgroupLbvserverBinding []interface{} `json:"lbgroup_lbvserver_binding,omitempty"`
	Name                    string        `json:"name,omitempty"`
}

type Lbparameter struct {
	Adccookieattributewarningmsg  string   `json:"adccookieattributewarningmsg,omitempty"`
	Allowboundsvcremoval          string   `json:"allowboundsvcremoval,omitempty"`
	Builtin                       []string `json:"builtin,omitempty"`
	Computedadccookieattribute    string   `json:"computedadccookieattribute,omitempty"`
	Consolidatedlconn             string   `json:"consolidatedlconn,omitempty"`
	Cookiepassphrase              string   `json:"cookiepassphrase,omitempty"`
	Dbsttl                        int      `json:"dbsttl,omitempty"`
	Dropmqttjumbomessage          string   `json:"dropmqttjumbomessage,omitempty"`
	Feature                       string   `json:"feature,omitempty"`
	Httponlycookieflag            string   `json:"httponlycookieflag,omitempty"`
	Lbhashalgorithm               string   `json:"lbhashalgorithm,omitempty"`
	Lbhashalgowinsize             int      `json:"lbhashalgowinsize,omitempty"`
	Lbhashfingers                 int      `json:"lbhashfingers,omitempty"`
	Literaladccookieattribute     string   `json:"literaladccookieattribute,omitempty"`
	Maxpipelinenat                int      `json:"maxpipelinenat,omitempty"`
	Monitorconnectionclose        string   `json:"monitorconnectionclose,omitempty"`
	Monitorskipmaxclient          string   `json:"monitorskipmaxclient,omitempty"`
	Nextgenapiresource            string   `json:"_nextgenapiresource,omitempty"`
	Overridepersistencyfororder   string   `json:"overridepersistencyfororder,omitempty"`
	Preferdirectroute             string   `json:"preferdirectroute,omitempty"`
	Proximityfromself             string   `json:"proximityfromself,omitempty"`
	Radiusmessageauthenticator    string   `json:"radiusmessageauthenticator,omitempty"`
	Retainservicestate            string   `json:"retainservicestate,omitempty"`
	Sessionsthreshold             int      `json:"sessionsthreshold,omitempty"`
	Startuprrfactor               int      `json:"startuprrfactor,omitempty"`
	Storemqttclientidandusername  string   `json:"storemqttclientidandusername,omitempty"`
	Undefaction                   string   `json:"undefaction,omitempty"`
	Useencryptedpersistencecookie string   `json:"useencryptedpersistencecookie,omitempty"`
	Useportforhashlb              string   `json:"useportforhashlb,omitempty"`
	Usesecuredpersistencecookie   string   `json:"usesecuredpersistencecookie,omitempty"`
	Vserverspecificmac            string   `json:"vserverspecificmac,omitempty"`
}

type LBVServerServiceGroupMemberBinding struct {
	Cookieipport      string `json:"cookieipport,omitempty"`
	Cookiename        string `json:"cookiename,omitempty"`
	Curstate          string `json:"curstate,omitempty"`
	Dynamicweight     int    `json:"dynamicweight,omitempty"`
	Ipv46             string `json:"ipv46,omitempty"`
	Name              string `json:"name,omitempty"`
	Order             int    `json:"order,omitempty"`
	Orderstr          string `json:"orderstr,omitempty"`
	Port              int    `json:"port,omitempty"`
	Preferredlocation string `json:"preferredlocation,omitempty"`
	Servicegroupname  string `json:"servicegroupname,omitempty"`
	Servicetype       string `json:"servicetype,omitempty"`
	Vserverid         string `json:"vserverid,omitempty"`
	Weight            int    `json:"weight,omitempty"`
}

type Lbmetrictable struct {
	Builtin            []string `json:"builtin,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Metric             string   `json:"metric,omitempty"`
	Metrictable        string   `json:"metrictable,omitempty"`
	Metrictype         string   `json:"metrictype,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Snmpoid            string   `json:"Snmpoid,omitempty"`
}
