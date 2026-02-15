package models

// appflow configuration structs
type Appflowparam struct {
	Aaausername                         string   `json:"aaausername,omitempty"`
	Analyticsauthtoken                  string   `json:"analyticsauthtoken,omitempty"`
	Appnamerefresh                      int      `json:"appnamerefresh,omitempty"`
	Auditlogs                           string   `json:"auditlogs,omitempty"`
	Builtin                             []string `json:"builtin,omitempty"`
	Cacheinsight                        string   `json:"cacheinsight,omitempty"`
	Clienttrafficonly                   string   `json:"clienttrafficonly,omitempty"`
	Connectionchaining                  string   `json:"connectionchaining,omitempty"`
	Cqareporting                        string   `json:"cqareporting,omitempty"`
	Distributedtracing                  string   `json:"distributedtracing,omitempty"`
	Disttracingsamplingrate             int      `json:"disttracingsamplingrate,omitempty"`
	Emailaddress                        string   `json:"emailaddress,omitempty"`
	Events                              string   `json:"events,omitempty"`
	Feature                             string   `json:"feature,omitempty"`
	Flowrecordinterval                  int      `json:"flowrecordinterval,omitempty"`
	Gxsessionreporting                  string   `json:"gxsessionreporting,omitempty"`
	Httpauthorization                   string   `json:"httpauthorization,omitempty"`
	Httpcontenttype                     string   `json:"httpcontenttype,omitempty"`
	Httpcookie                          string   `json:"httpcookie,omitempty"`
	Httpdomain                          string   `json:"httpdomain,omitempty"`
	Httphost                            string   `json:"httphost,omitempty"`
	Httplocation                        string   `json:"httplocation,omitempty"`
	Httpmethod                          string   `json:"httpmethod,omitempty"`
	Httpquerywithurl                    string   `json:"httpquerywithurl,omitempty"`
	Httpreferer                         string   `json:"httpreferer,omitempty"`
	Httpsetcookie                       string   `json:"httpsetcookie,omitempty"`
	Httpsetcookie2                      string   `json:"httpsetcookie2,omitempty"`
	Httpurl                             string   `json:"httpurl,omitempty"`
	Httpuseragent                       string   `json:"httpuseragent,omitempty"`
	Httpvia                             string   `json:"httpvia,omitempty"`
	Httpxforwardedfor                   string   `json:"httpxforwardedfor,omitempty"`
	Identifiername                      string   `json:"identifiername,omitempty"`
	Identifiersessionname               string   `json:"identifiersessionname,omitempty"`
	Logstreamovernsip                   string   `json:"logstreamovernsip,omitempty"`
	Lsnlogging                          string   `json:"lsnlogging,omitempty"`
	Metrics                             string   `json:"metrics,omitempty"`
	Nextgenapiresource                  string   `json:"_nextgenapiresource,omitempty"`
	Observationdomainid                 int      `json:"observationdomainid,omitempty"`
	Observationdomainname               string   `json:"observationdomainname,omitempty"`
	Observationpointid                  int      `json:"observationpointid,omitempty"`
	Securityinsightrecordinterval       int      `json:"securityinsightrecordinterval,omitempty"`
	Securityinsighttraffic              string   `json:"securityinsighttraffic,omitempty"`
	Skipcacheredirectionhttptransaction string   `json:"skipcacheredirectionhttptransaction,omitempty"`
	Subscriberawareness                 string   `json:"subscriberawareness,omitempty"`
	Subscriberidobfuscation             string   `json:"subscriberidobfuscation,omitempty"`
	Subscriberidobfuscationalgo         string   `json:"subscriberidobfuscationalgo,omitempty"`
	Tcpattackcounterinterval            int      `json:"tcpattackcounterinterval,omitempty"`
	Tcpburstreporting                   string   `json:"tcpburstreporting,omitempty"`
	Tcpburstreportingthreshold          int      `json:"tcpburstreportingthreshold,omitempty"`
	Templaterefresh                     int      `json:"templaterefresh,omitempty"`
	Timeseriesovernsip                  string   `json:"timeseriesovernsip,omitempty"`
	Udppmtu                             int      `json:"udppmtu,omitempty"`
	Urlcategory                         string   `json:"urlcategory,omitempty"`
	Usagerecordinterval                 int      `json:"usagerecordinterval,omitempty"`
	Videoinsight                        string   `json:"videoinsight,omitempty"`
	Websaasappusagereporting            string   `json:"websaasappusagereporting,omitempty"`
}

type AppflowglobalBinding struct {
	AppflowglobalAppflowpolicyBinding []interface{} `json:"appflowglobal_appflowpolicy_binding,omitempty"`
}

type AppflowpolicyAppflowglobalBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Appflowpolicylabel struct {
	Count                  float64 `json:"__count,omitempty"`
	Description            string  `json:"description,omitempty"`
	Flowtype               int     `json:"flowtype,omitempty"`
	Gotopriorityexpression string  `json:"gotopriorityexpression,omitempty"`
	Hits                   int     `json:"hits,omitempty"`
	InvokeLabelname        string  `json:"invoke_labelname,omitempty"`
	Labelname              string  `json:"labelname,omitempty"`
	Labeltype              string  `json:"labeltype,omitempty"`
	Newname                string  `json:"newname,omitempty"`
	Nextgenapiresource     string  `json:"_nextgenapiresource,omitempty"`
	Numpol                 int     `json:"numpol,omitempty"`
	Policylabeltype        string  `json:"policylabeltype,omitempty"`
	Policyname             string  `json:"policyname,omitempty"`
	Priority               int     `json:"priority,omitempty"`
}

type AppflowpolicylabelBinding struct {
	AppflowpolicylabelAppflowpolicyBinding []interface{} `json:"appflowpolicylabel_appflowpolicy_binding,omitempty"`
	Labelname                              string        `json:"labelname,omitempty"`
}

type AppflowpolicyVpnvserverBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type AppflowglobalAppflowpolicyBinding struct {
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

type AppflowpolicyAppflowpolicylabelBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type AppflowactionBinding struct {
	AppflowactionAnalyticsprofileBinding []interface{} `json:"appflowaction_analyticsprofile_binding,omitempty"`
	Name                                 string        `json:"name,omitempty"`
}

type AppflowpolicylabelAppflowpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type AppflowpolicyBinding struct {
	AppflowpolicyAppflowglobalBinding      []interface{} `json:"appflowpolicy_appflowglobal_binding,omitempty"`
	AppflowpolicyAppflowpolicylabelBinding []interface{} `json:"appflowpolicy_appflowpolicylabel_binding,omitempty"`
	AppflowpolicyCsvserverBinding          []interface{} `json:"appflowpolicy_csvserver_binding,omitempty"`
	AppflowpolicyLbvserverBinding          []interface{} `json:"appflowpolicy_lbvserver_binding,omitempty"`
	AppflowpolicyVpnvserverBinding         []interface{} `json:"appflowpolicy_vpnvserver_binding,omitempty"`
	Name                                   string        `json:"name,omitempty"`
}

type AppflowpolicyLbvserverBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Appflowaction struct {
	Botinsight             string   `json:"botinsight,omitempty"`
	Ciinsight              string   `json:"ciinsight,omitempty"`
	Clientsidemeasurements string   `json:"clientsidemeasurements,omitempty"`
	Collectors             []string `json:"collectors,omitempty"`
	Comment                string   `json:"comment,omitempty"`
	Count                  float64  `json:"__count,omitempty"`
	Description            string   `json:"description,omitempty"`
	Distributionalgorithm  string   `json:"distributionalgorithm,omitempty"`
	Hits                   int      `json:"hits,omitempty"`
	Metricslog             bool     `json:"metricslog,omitempty"`
	Name                   string   `json:"name,omitempty"`
	Newname                string   `json:"newname,omitempty"`
	Nextgenapiresource     string   `json:"_nextgenapiresource,omitempty"`
	Pagetracking           string   `json:"pagetracking,omitempty"`
	Referencecount         int      `json:"referencecount,omitempty"`
	Securityinsight        string   `json:"securityinsight,omitempty"`
	Transactionlog         string   `json:"transactionlog,omitempty"`
	Videoanalytics         string   `json:"videoanalytics,omitempty"`
	Webinsight             string   `json:"webinsight,omitempty"`
}

type Appflowpolicy struct {
	Action             string   `json:"action,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	Comment            string   `json:"comment,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Description        string   `json:"description,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	Name               string   `json:"name,omitempty"`
	Newname            string   `json:"newname,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Rule               string   `json:"rule,omitempty"`
	Undefaction        string   `json:"undefaction,omitempty"`
	Undefhits          int      `json:"undefhits,omitempty"`
}

type AppflowpolicyCsvserverBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Appflowcollector struct {
	Count              float64 `json:"__count,omitempty"`
	Ipaddress          string  `json:"ipaddress,omitempty"`
	Name               string  `json:"name,omitempty"`
	Netprofile         string  `json:"netprofile,omitempty"`
	Newname            string  `json:"newname,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Port               int     `json:"port,omitempty"`
	State              string  `json:"state,omitempty"`
	Transport          string  `json:"transport,omitempty"`
}

type AppflowactionAnalyticsprofileBinding struct {
	Analyticsprofile string `json:"analyticsprofile,omitempty"`
	Name             string `json:"name,omitempty"`
}
