// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Appflowcollector struct {
	Name               string `json:"name,omitempty"`
	Ipaddress          string `json:"ipaddress,omitempty"`
	Port               int    `json:"port,omitempty"`
	Netprofile         string `json:"netprofile,omitempty"`
	Transport          string `json:"transport,omitempty"`
	Newname            string `json:"newname,omitempty"`
	State              string `json:"state,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Appflowglobalbinding struct {
}

type Appflowpolicy struct {
	Name               string `json:"name,omitempty"`
	Rule               string `json:"rule,omitempty"`
	Action             string `json:"action,omitempty"`
	Undefaction        string `json:"undefaction,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Newname            string `json:"newname,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Undefhits          string `json:"undefhits,omitempty"`
	Description        string `json:"description,omitempty"`
	Builtin            string `json:"builtin,omitempty"`
	Feature            string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Appflowpolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Appflowpolicycsvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Appflowpolicyappflowglobalbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Appflowpolicyappflowpolicylabelbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Appflowpolicypolicylabelbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Activepolicy           int32  `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Appflowpolicylabel struct {
	Labelname              string `json:"labelname,omitempty"`
	Policylabeltype        string `json:"policylabeltype,omitempty"`
	Newname                string `json:"newname,omitempty"`
	Numpol                 string `json:"numpol,omitempty"`
	Hits                   string `json:"hits,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               string `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Flowtype               string `json:"flowtype,omitempty"`
	Description            string `json:"description,omitempty"`
	Nextgenapiresource     string `json:"_nextgenapiresource,omitempty"`
}

type Appflowactionbinding struct {
	Name string `json:"name,omitempty"`
}

type Appflowparam struct {
	Templaterefresh                     int    `json:"templaterefresh,omitempty"`
	Appnamerefresh                      int    `json:"appnamerefresh,omitempty"`
	Flowrecordinterval                  int    `json:"flowrecordinterval,omitempty"`
	Securityinsightrecordinterval       int    `json:"securityinsightrecordinterval,omitempty"`
	Udppmtu                             int    `json:"udppmtu,omitempty"`
	Httpurl                             string `json:"httpurl,omitempty"`
	Aaausername                         string `json:"aaausername,omitempty"`
	Httpcookie                          string `json:"httpcookie,omitempty"`
	Httpreferer                         string `json:"httpreferer,omitempty"`
	Httpmethod                          string `json:"httpmethod,omitempty"`
	Httphost                            string `json:"httphost,omitempty"`
	Httpuseragent                       string `json:"httpuseragent,omitempty"`
	Clienttrafficonly                   string `json:"clienttrafficonly,omitempty"`
	Httpcontenttype                     string `json:"httpcontenttype,omitempty"`
	Httpauthorization                   string `json:"httpauthorization,omitempty"`
	Httpvia                             string `json:"httpvia,omitempty"`
	Httpxforwardedfor                   string `json:"httpxforwardedfor,omitempty"`
	Httplocation                        string `json:"httplocation,omitempty"`
	Httpsetcookie                       string `json:"httpsetcookie,omitempty"`
	Httpsetcookie2                      string `json:"httpsetcookie2,omitempty"`
	Connectionchaining                  string `json:"connectionchaining,omitempty"`
	Httpdomain                          string `json:"httpdomain,omitempty"`
	Skipcacheredirectionhttptransaction string `json:"skipcacheredirectionhttptransaction,omitempty"`
	Identifiername                      string `json:"identifiername,omitempty"`
	Identifiersessionname               string `json:"identifiersessionname,omitempty"`
	Observationdomainid                 int    `json:"observationdomainid,omitempty"`
	Observationdomainname               string `json:"observationdomainname,omitempty"`
	Subscriberawareness                 string `json:"subscriberawareness,omitempty"`
	Subscriberidobfuscation             string `json:"subscriberidobfuscation,omitempty"`
	Subscriberidobfuscationalgo         string `json:"subscriberidobfuscationalgo,omitempty"`
	Gxsessionreporting                  string `json:"gxsessionreporting,omitempty"`
	Securityinsighttraffic              string `json:"securityinsighttraffic,omitempty"`
	Cacheinsight                        string `json:"cacheinsight,omitempty"`
	Videoinsight                        string `json:"videoinsight,omitempty"`
	Httpquerywithurl                    string `json:"httpquerywithurl,omitempty"`
	Urlcategory                         string `json:"urlcategory,omitempty"`
	Lsnlogging                          string `json:"lsnlogging,omitempty"`
	Cqareporting                        string `json:"cqareporting,omitempty"`
	Emailaddress                        string `json:"emailaddress,omitempty"`
	Usagerecordinterval                 int    `json:"usagerecordinterval,omitempty"`
	Websaasappusagereporting            string `json:"websaasappusagereporting,omitempty"`
	Metrics                             string `json:"metrics,omitempty"`
	Events                              string `json:"events,omitempty"`
	Auditlogs                           string `json:"auditlogs,omitempty"`
	Observationpointid                  int    `json:"observationpointid,omitempty"`
	Distributedtracing                  string `json:"distributedtracing,omitempty"`
	Disttracingsamplingrate             int    `json:"disttracingsamplingrate,omitempty"`
	Tcpattackcounterinterval            int    `json:"tcpattackcounterinterval,omitempty"`
	Logstreamovernsip                   string `json:"logstreamovernsip,omitempty"`
	Analyticsauthtoken                  string `json:"analyticsauthtoken,omitempty"`
	Timeseriesovernsip                  string `json:"timeseriesovernsip,omitempty"`
	Builtin                             string `json:"builtin,omitempty"`
	Feature                             string `json:"feature,omitempty"`
	Tcpburstreporting                   string `json:"tcpburstreporting,omitempty"`
	Tcpburstreportingthreshold          string `json:"tcpburstreportingthreshold,omitempty"`
	Nextgenapiresource                  string `json:"_nextgenapiresource,omitempty"`
}

type Appflowpolicyglobalbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Activepolicy           int32  `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Appflowpolicylbvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Appflowpolicyvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Activepolicy           int32  `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Appflowpolicylabelappflowpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Appflowpolicylabelbinding struct {
	Labelname string `json:"labelname,omitempty"`
}

type Appflowaction struct {
	Name                   string   `json:"name,omitempty"`
	Collectors             []string `json:"collectors,omitempty"`
	Clientsidemeasurements string   `json:"clientsidemeasurements,omitempty"`
	Pagetracking           string   `json:"pagetracking,omitempty"`
	Webinsight             string   `json:"webinsight,omitempty"`
	Securityinsight        string   `json:"securityinsight,omitempty"`
	Botinsight             string   `json:"botinsight,omitempty"`
	Ciinsight              string   `json:"ciinsight,omitempty"`
	Videoanalytics         string   `json:"videoanalytics,omitempty"`
	Distributionalgorithm  string   `json:"distributionalgorithm,omitempty"`
	Metricslog             bool     `json:"metricslog,omitempty"`
	Transactionlog         string   `json:"transactionlog,omitempty"`
	Comment                string   `json:"comment,omitempty"`
	Newname                string   `json:"newname,omitempty"`
	Hits                   string   `json:"hits,omitempty"`
	Referencecount         string   `json:"referencecount,omitempty"`
	Description            string   `json:"description,omitempty"`
	Nextgenapiresource     string   `json:"_nextgenapiresource,omitempty"`
}

type Appflowactionanalyticsprofilebinding struct {
	Analyticsprofile string `json:"analyticsprofile,omitempty"`
	Name             string `json:"name,omitempty"`
}

type Appflowglobalappflowpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Type                   string `json:"type,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Numpol                 int    `json:"numpol,omitempty"`
	Flowtype               int    `json:"flowtype,omitempty"`
	Globalbindtype         string `json:"globalbindtype,omitempty"`
}

type Appflowglobalpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Type                   string `json:"type,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Numpol                 uint32 `json:"numpol,omitempty"`
	Flowtype               uint32 `json:"flowtype,omitempty"`
	Globalbindtype         string `json:"globalbindtype,omitempty"`
}

type Appflowpolicyvpnvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Appflowpolicylabelpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Appflowactionprofilebinding struct {
	Analyticsprofile string `json:"analyticsprofile,omitempty"`
	Name             string `json:"name,omitempty"`
}
