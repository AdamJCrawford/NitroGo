// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Lbvserverfilterpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Sc                     string `json:"sc,omitempty"`
	Name                   string `json:"name,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Lbpolicygslbvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Lbvserverprofilebinding struct {
	Analyticsprofile string `json:"analyticsprofile,omitempty"`
	Name             string `json:"name,omitempty"`
}

type Lbvserverservicebinding struct {
	Servicename       string `json:"servicename,omitempty"`
	Ipv46             string `json:"ipv46,omitempty"`
	Port              int    `json:"port,omitempty"`
	Servicetype       string `json:"servicetype,omitempty"`
	Curstate          string `json:"curstate,omitempty"`
	Weight            int    `json:"weight,omitempty"`
	Dynamicweight     int    `json:"dynamicweight,omitempty"`
	Cookieipport      string `json:"cookieipport,omitempty"`
	Vserverid         string `json:"vserverid,omitempty"`
	Vsvrbindsvcip     string `json:"vsvrbindsvcip,omitempty"`
	Vsvrbindsvcport   int    `json:"vsvrbindsvcport,omitempty"`
	Preferredlocation string `json:"preferredlocation,omitempty"`
	Order             int    `json:"order,omitempty"`
	Orderstr          string `json:"orderstr,omitempty"`
	Name              string `json:"name,omitempty"`
	Servicegroupname  string `json:"servicegroupname,omitempty"`
}

type Lbvservertransformpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
}

type Lbvservervideooptimizationpacingpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
}

type Lbmetrictablebinding struct {
	Metrictable string `json:"metrictable,omitempty"`
}

type Lbmonbindingsgslbservicegroupbinding struct {
	Servicegroupname          string `json:"servicegroupname,omitempty"`
	Servicetype               string `json:"servicetype,omitempty"`
	Boundservicegroupsvrstate string `json:"boundservicegroupsvrstate,omitempty"`
	Monstate                  string `json:"monstate,omitempty"`
	Monitorname               string `json:"monitorname,omitempty"`
}

type Lbmonitorservicegroupbinding struct {
	Monitorname      string `json:"monitorname,omitempty"`
	Servicename      string `json:"servicename,omitempty"`
	Dupstate         string `json:"dup_state,omitempty"`
	Dupweight        int    `json:"dup_weight,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	State            string `json:"state,omitempty"`
	Weight           int    `json:"weight,omitempty"`
}

type Lbsipparameters struct {
	Rnatsrcport         int    `json:"rnatsrcport,omitempty"`
	Rnatdstport         int    `json:"rnatdstport,omitempty"`
	Retrydur            int    `json:"retrydur,omitempty"`
	Addrportvip         string `json:"addrportvip,omitempty"`
	Sip503ratethreshold int    `json:"sip503ratethreshold,omitempty"`
	Rnatsecuresrcport   int    `json:"rnatsecuresrcport,omitempty"`
	Rnatsecuredstport   int    `json:"rnatsecuredstport,omitempty"`
	Builtin             string `json:"builtin,omitempty"`
	Feature             string `json:"feature,omitempty"`
	Nextgenapiresource  string `json:"_nextgenapiresource,omitempty"`
}

type Lbmonbindingsservicebinding struct {
	Servicename string `json:"servicename,omitempty"`
	Ipaddress   string `json:"ipaddress,omitempty"`
	Port        int    `json:"port,omitempty"`
	Servicetype string `json:"servicetype,omitempty"`
	Svrstate    string `json:"svrstate,omitempty"`
	Monsvcstate string `json:"monsvcstate,omitempty"`
	Monitorname string `json:"monitorname,omitempty"`
}

type Lbvserverauditnslogpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Name                   string `json:"name,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Order                  int    `json:"order,omitempty"`
}

type Lbmetrictable struct {
	Metrictable        string `json:"metrictable,omitempty"`
	Metric             string `json:"metric,omitempty"`
	Snmpoid            string `json:"Snmpoid,omitempty"`
	Metrictype         string `json:"metrictype,omitempty"`
	Builtin            string `json:"builtin,omitempty"`
	Feature            string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Lbparameter struct {
	Httponlycookieflag            string `json:"httponlycookieflag,omitempty"`
	Usesecuredpersistencecookie   string `json:"usesecuredpersistencecookie,omitempty"`
	Useencryptedpersistencecookie string `json:"useencryptedpersistencecookie,omitempty"`
	Cookiepassphrase              string `json:"cookiepassphrase,omitempty"`
	Consolidatedlconn             string `json:"consolidatedlconn,omitempty"`
	Useportforhashlb              string `json:"useportforhashlb,omitempty"`
	Preferdirectroute             string `json:"preferdirectroute,omitempty"`
	Startuprrfactor               int    `json:"startuprrfactor,omitempty"`
	Monitorskipmaxclient          string `json:"monitorskipmaxclient,omitempty"`
	Monitorconnectionclose        string `json:"monitorconnectionclose,omitempty"`
	Vserverspecificmac            string `json:"vserverspecificmac,omitempty"`
	Allowboundsvcremoval          string `json:"allowboundsvcremoval,omitempty"`
	Retainservicestate            string `json:"retainservicestate,omitempty"`
	Dbsttl                        int    `json:"dbsttl,omitempty"`
	Maxpipelinenat                int    `json:"maxpipelinenat,omitempty"`
	Literaladccookieattribute     string `json:"literaladccookieattribute,omitempty"`
	Computedadccookieattribute    string `json:"computedadccookieattribute,omitempty"`
	Storemqttclientidandusername  string `json:"storemqttclientidandusername,omitempty"`
	Dropmqttjumbomessage          string `json:"dropmqttjumbomessage,omitempty"`
	Lbhashalgorithm               string `json:"lbhashalgorithm,omitempty"`
	Lbhashfingers                 int    `json:"lbhashfingers,omitempty"`
	Undefaction                   string `json:"undefaction,omitempty"`
	Proximityfromself             string `json:"proximityfromself,omitempty"`
	Sessionsthreshold             string `json:"sessionsthreshold,omitempty"`
	Builtin                       string `json:"builtin,omitempty"`
	Feature                       string `json:"feature,omitempty"`
	Adccookieattributewarningmsg  string `json:"adccookieattributewarningmsg,omitempty"`
	Lbhashalgowinsize             string `json:"lbhashalgowinsize,omitempty"`
	Overridepersistencyfororder   string `json:"overridepersistencyfororder,omitempty"`
	Nextgenapiresource            string `json:"_nextgenapiresource,omitempty"`
}

type LBVServer struct {
	ADFSProxyProfile                   string `json:"adfsproxyprofile,omitempty"`
	ActiveServices                     string `json:"activeservices,omitempty"`
	APIProfile                         string `json:"apiprofile,omitempty"`
	AppFlowLog                         string `json:"appflowlog,omitempty"`
	Authentication                     string `json:"authentication,omitempty"`
	AuthenticationHost                 string `json:"authenticationhost,omitempty"`
	Authn401                           string `json:"authn401,omitempty"`
	AuthnProfile                       string `json:"authnprofile,omitempty"`
	AuthnVSName                        string `json:"authnvsname,omitempty"`
	BackupLBMethod                     string `json:"backuplbmethod,omitempty"`
	BackupPersistenceTimeout           int    `json:"backuppersistencetimeout,omitempty"`
	BackupVServer                      string `json:"backupvserver,omitempty"`
	BackupvServerStatus                string `json:"backupvserverstatus,omitempty"`
	BindPoint                          string `json:"bindpoint,omitempty"`
	BypassAAAA                         string `json:"bypassaaaa,omitempty"`
	Cacheable                          string `json:"cacheable,omitempty"`
	CacheVServer                       string `json:"cachevserver,omitempty"`
	CltTimeOut                         string `json:"clttimeout,omitempty"`
	Comment                            string `json:"comment,omitempty"`
	ConnFailOver                       string `json:"connfailover,omitempty"`
	ConsolidatedLConn                  string `json:"consolidatedlconn,omitempty"`
	ConsolidatedLConngbl               string `json:"consolidatedlconngbl,omitempty"`
	CookieDomain                       string `json:"cookiedomain,omitempty"`
	CookieName                         string `json:"cookiename,omitempty"`
	CurrentactiveOrder                 string `json:"currentactiveorder,omitempty"`
	CurState                           string `json:"curstate,omitempty"`
	DataLength                         string `json:"datalength,omitempty"`
	DataOffset                         string `json:"dataoffset,omitempty"`
	DbProfileName                      string `json:"dbprofilename,omitempty"`
	DbsLB                              string `json:"dbslb,omitempty"`
	DisablePrimaryOnDown               string `json:"disableprimaryondown,omitempty"`
	DNS64                              string `json:"dns64,omitempty"`
	DNSOverHTTPS                       string `json:"dnsoverhttps,omitempty"`
	DNSProfileName                     string `json:"dnsprofilename,omitempty"`
	DNSVServerName                     string `json:"dnsvservername,omitempty"`
	Domain                             string `json:"domain,omitempty"`
	DownStateFlush                     string `json:"downstateflush,omitempty"`
	EffectiveState                     string `json:"effectivestate,omitempty"`
	GroupName                          string `json:"groupname,omitempty"`
	Gt2GB                              string `json:"gt2gb,omitempty"`
	HashLength                         int    `json:"hashlength,omitempty"`
	Health                             string `json:"health,omitempty"`
	HealthThreshold                    string `json:"healththreshold,omitempty"`
	Homepage                           string `json:"homepage,omitempty"`
	HTTPProfileName                    string `json:"httpprofilename,omitempty"`
	HTTPSRedirectURL                   string `json:"httpsredirecturl,omitempty"`
	ICMPVSRResponse                    string `json:"icmpvsrresponse,omitempty"`
	InsertVServerIPPort                string `json:"insertvserveripport,omitempty"`
	IPMapping                          string `json:"ipmapping,omitempty"`
	IPMask                             string `json:"ipmask,omitempty"`
	IPPattern                          string `json:"ippattern,omitempty"`
	IPSet                              string `json:"ipset,omitempty"`
	IPv46                              string `json:"ipv46,omitempty"`
	IsGSLB                             bool   `json:"isgslb,omitempty"`
	L2Conn                             string `json:"l2conn,omitempty"`
	LBMethod                           string `json:"lbmethod,omitempty"`
	LBProfileName                      string `json:"lbprofilename,omitempty"`
	LBRRReason                         int    `json:"lbrrreason,omitempty"`
	ListenPolicy                       string `json:"listenpolicy,omitempty"`
	ListenPriority                     string `json:"listenpriority,omitempty"`
	M                                  string `json:"m,omitempty"`
	MACModeRetainVLAN                  string `json:"macmoderetainvlan,omitempty"`
	Map                                string `json:"map,omitempty"`
	MaxAutoScaleMembers                string `json:"maxautoscalemembers,omitempty"`
	MinAutoScaleMembers                string `json:"minautoscalemembers,omitempty"`
	MSSQLServerVersion                 string `json:"mssqlserverversion,omitempty"`
	MySQLCharacterSet                  int    `json:"mysqlcharacterset,omitempty"`
	MySQLProtocolVersion               int    `json:"mysqlprotocolversion,omitempty"`
	MySQLServerCapabilities            int    `json:"mysqlservercapabilities,omitempty"`
	MySQLServerVersion                 string `json:"mysqlserverversion,omitempty"`
	Name                               string `json:"name,omitempty"`
	Netmask                            string `json:"netmask,omitempty"`
	NetProfile                         string `json:"netprofile,omitempty"`
	NewName                            string `json:"newname,omitempty"`
	NewServiceRequest                  int    `json:"newservicerequest,omitempty"`
	NewServiceRequestIncrementInterval int    `json:"newservicerequestincrementinterval,omitempty"`
	NewserviceRequestUnit              string `json:"newservicerequestunit,omitempty"`
	NextGenAPIResource                 string `json:"_nextgenapiresource,omitempty"`
	NGName                             string `json:"ngname,omitempty"`
	NodeFaultBindings                  string `json:"nodefaultbindings,omitempty"`
	OracleServerVersion                string `json:"oracleserverversion,omitempty"`
	Order                              int    `json:"order,omitempty"`
	OrderThreshold                     string `json:"orderthreshold,omitempty"`
	PersistAVPNO                       []int  `json:"persistavpno,omitempty"`
	PersistenceBackup                  string `json:"persistencebackup,omitempty"`
	PersistenceType                    string `json:"persistencetype,omitempty"`
	PersistMask                        string `json:"persistmask,omitempty"`
	Port                               int    `json:"port,omitempty"`
	Precedence                         string `json:"precedence,omitempty"`
	ProbePort                          int    `json:"probeport,omitempty"`
	ProbeProtocol                      string `json:"probeprotocol,omitempty"`
	ProbeSuccessResponsecode           string `json:"probesuccessresponsecode,omitempty"`
	ProcessLocal                       string `json:"processlocal,omitempty"`
	Push                               string `json:"push,omitempty"`
	PushLabel                          string `json:"pushlabel,omitempty"`
	PushMultiClients                   string `json:"pushmulticlients,omitempty"`
	PushVServer                        string `json:"pushvserver,omitempty"`
	QUICBridgeProfileName              string `json:"quicbridgeprofilename,omitempty"`
	QUICProfileName                    string `json:"quicprofilename,omitempty"`
	Range                              string `json:"range,omitempty"`
	RecursionAvailable                 string `json:"recursionavailable,omitempty"`
	Redirect                           string `json:"redirect,omitempty"`
	RedirectFromPort                   int    `json:"redirectfromport,omitempty"`
	RedirectPortRewrite                string `json:"redirectportrewrite,omitempty"`
	RedirURL                           string `json:"redirurl,omitempty"`
	RedirURLFlags                      bool   `json:"redirurlflags,omitempty"`
	Resrule                            string `json:"resrule,omitempty"`
	RetainConnectionsonCluster         string `json:"retainconnectionsoncluster,omitempty"`
	RHIState                           string `json:"rhistate,omitempty"`
	RTSPNAT                            string `json:"rtspnat,omitempty"`
	Rule                               string `json:"rule,omitempty"`
	Ruletype                           string `json:"ruletype,omitempty"`
	ServiceName                        string `json:"servicename,omitempty"`
	ServiceType                        string `json:"servicetype,omitempty"`
	Sessionless                        string `json:"sessionless,omitempty"`
	SkipPersistency                    string `json:"skippersistency,omitempty"`
	SoBackupAction                     string `json:"sobackupaction,omitempty"`
	SoMethod                           string `json:"somethod,omitempty"`
	SoPersistence                      string `json:"sopersistence,omitempty"`
	SoPersistenceTimeout               string `json:"sopersistencetimeout,omitempty"`
	SoThreshold                        string `json:"sothreshold,omitempty"`
	State                              string `json:"state,omitempty"`
	StateChangeTimeMsec                string `json:"statechangetimemsec,omitempty"`
	StateChangeTimeSec                 string `json:"statechangetimesec,omitempty"`
	StateChangeTimeSeconds             string `json:"statechangetimeseconds,omitempty"`
	Status                             int    `json:"status,omitempty"`
	TCPProbePort                       int    `json:"tcpprobeport,omitempty"`
	TCPProfileName                     string `json:"tcpprofilename,omitempty"`
	Td                                 string `json:"td,omitempty"`
	ThresholdValue                     int    `json:"thresholdvalue,omitempty"`
	TicksSinceLastStateChange          string `json:"tickssincelaststatechange,omitempty"`
	Timeout                            int    `json:"timeout,omitempty"`
	ToggleOrder                        string `json:"toggleorder,omitempty"`
	TOSID                              int    `json:"tosid,omitempty"`
	TotalServices                      string `json:"totalservices,omitempty"`
	TROFSPersistence                   string `json:"trofspersistence,omitempty"`
	Type                               string `json:"type,omitempty"`
	V6NetMaskLen                       int    `json:"v6netmasklen,omitempty"`
	V6PersistMaskLen                   string `json:"v6persistmasklen,omitempty"`
	Value                              string `json:"value,omitempty"`
	Version                            int    `json:"version,omitempty"`
	VIPHeader                          string `json:"vipheader,omitempty"`
	VSvrDynConnsoThreshold             string `json:"vsvrdynconnsothreshold,omitempty"`
	Weight                             int    `json:"weight,omitempty"`
}

type Lbvserverauditsyslogpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Name                   string `json:"name,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Order                  int    `json:"order,omitempty"`
}

type Lbpolicy struct {
	Name               string `json:"name,omitempty"`
	Rule               string `json:"rule,omitempty"`
	Action             string `json:"action,omitempty"`
	Undefaction        string `json:"undefaction,omitempty"`
	Logaction          string `json:"logaction,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Newname            string `json:"newname,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Undefhits          string `json:"undefhits,omitempty"`
	Feature            string `json:"feature,omitempty"`
	Builtin            string `json:"builtin,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Lbgroupbinding struct {
	Name string `json:"name,omitempty"`
}

type Lbpolicylabellbpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Lbvserverappqoepolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Name                   string `json:"name,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Order                  int    `json:"order,omitempty"`
}

type Lbvserverdnspolicy64binding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Order                  int    `json:"order,omitempty"`
}

type Lbvserverscpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Name                   string `json:"name,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Lbvserversyslogpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Sc                     string `json:"sc,omitempty"`
	Name                   string `json:"name,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Lbmonitorcertkeybinding struct {
	Certkeyname string `json:"certkeyname,omitempty"`
	Ca          bool   `json:"ca,omitempty"`
	Crlcheck    string `json:"crlcheck,omitempty"`
	Ocspcheck   string `json:"ocspcheck,omitempty"`
	Monitorname string `json:"monitorname,omitempty"`
}

type Lbpolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Lbwlmbinding struct {
	Wlmname string `json:"wlmname,omitempty"`
}

type Lbvservercachepolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
}

type Lbmonitorservicebinding struct {
	Monitorname      string `json:"monitorname,omitempty"`
	Servicename      string `json:"servicename,omitempty"`
	Dupstate         string `json:"dup_state,omitempty"`
	Dupweight        int    `json:"dup_weight,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	State            string `json:"state,omitempty"`
	Weight           int    `json:"weight,omitempty"`
}

type Lbpolicylbglobalbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Lbprofile struct {
	Lbprofilename                 string `json:"lbprofilename,omitempty"`
	Dbslb                         string `json:"dbslb,omitempty"`
	Processlocal                  string `json:"processlocal,omitempty"`
	Httponlycookieflag            string `json:"httponlycookieflag,omitempty"`
	Cookiepassphrase              string `json:"cookiepassphrase,omitempty"`
	Usesecuredpersistencecookie   string `json:"usesecuredpersistencecookie,omitempty"`
	Useencryptedpersistencecookie string `json:"useencryptedpersistencecookie,omitempty"`
	Literaladccookieattribute     string `json:"literaladccookieattribute,omitempty"`
	Computedadccookieattribute    string `json:"computedadccookieattribute,omitempty"`
	Storemqttclientidandusername  string `json:"storemqttclientidandusername,omitempty"`
	Lbhashalgorithm               string `json:"lbhashalgorithm,omitempty"`
	Lbhashfingers                 int    `json:"lbhashfingers,omitempty"`
	Proximityfromself             string `json:"proximityfromself,omitempty"`
	Vsvrcount                     string `json:"vsvrcount,omitempty"`
	Adccookieattributewarningmsg  string `json:"adccookieattributewarningmsg,omitempty"`
	Lbhashalgowinsize             string `json:"lbhashalgowinsize,omitempty"`
	Nextgenapiresource            string `json:"_nextgenapiresource,omitempty"`
}

type Lbvserveranalyticsprofilebinding struct {
	Analyticsprofile string `json:"analyticsprofile,omitempty"`
	Name             string `json:"name,omitempty"`
	Order            int    `json:"order,omitempty"`
}

type Lbvservercsvserverbinding struct {
	Cachevserver  string `json:"cachevserver,omitempty"`
	Policyname    string `json:"policyname,omitempty"`
	Cachetype     string `json:"cachetype,omitempty"`
	Priority      int    `json:"priority,omitempty"`
	Hits          int    `json:"hits,omitempty"`
	Pipolicyhits  int    `json:"pipolicyhits,omitempty"`
	Policysubtype int    `json:"policysubtype,omitempty"`
	Name          string `json:"name,omitempty"`
	Order         int    `json:"order,omitempty"`
}

type Lbvserverdetectionpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Lbvserverdospolicybinding struct {
	Policyname string `json:"policyname,omitempty"`
	Priority   int    `json:"priority,omitempty"`
	Name       string `json:"name,omitempty"`
}

type Lbvserverspilloverpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Name                   string `json:"name,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Order                  int    `json:"order,omitempty"`
}

type Lbwlm struct {
	Wlmname            string `json:"wlmname,omitempty"`
	Ipaddress          string `json:"ipaddress,omitempty"`
	Port               int    `json:"port,omitempty"`
	Lbuid              string `json:"lbuid,omitempty"`
	Katimeout          int    `json:"katimeout,omitempty"`
	Secure             string `json:"secure,omitempty"`
	State              string `json:"state,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Lbgloballbpolicybinding struct {
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

type Lbmetrictablemetricbinding struct {
	Metric      string `json:"metric,omitempty"`
	Snmpoid     string `json:"Snmpoid,omitempty"`
	Metrictype  string `json:"metrictype,omitempty"`
	Metrictable string `json:"metrictable,omitempty"`
}

type Lbmonitormetricbinding struct {
	Metric          string `json:"metric,omitempty"`
	Metrictable     string `json:"metrictable,omitempty"`
	Metricunit      string `json:"metric_unit,omitempty"`
	Metricweight    int    `json:"metricweight,omitempty"`
	Metricthreshold int    `json:"metricthreshold,omitempty"`
	Monitorname     string `json:"monitorname,omitempty"`
}

type Lbpersistentsessions struct {
	Vserver              string `json:"vserver,omitempty"`
	Nodeid               int    `json:"nodeid,omitempty"`
	Persistenceparameter string `json:"persistenceparameter,omitempty"`
	Type                 string `json:"type,omitempty"`
	Typestring           string `json:"typestring,omitempty"`
	Srcip                string `json:"srcip,omitempty"`
	Srcipv6              string `json:"srcipv6,omitempty"`
	Destip               string `json:"destip,omitempty"`
	Destipv6             string `json:"destipv6,omitempty"`
	Flags                string `json:"flags,omitempty"`
	Destport             string `json:"destport,omitempty"`
	Vservername          string `json:"vservername,omitempty"`
	Timeout              string `json:"timeout,omitempty"`
	Referencecount       string `json:"referencecount,omitempty"`
	Persistenceparam     string `json:"persistenceparam,omitempty"`
	Cnamepersparam       string `json:"cnamepersparam,omitempty"`
	Nextgenapiresource   string `json:"_nextgenapiresource,omitempty"`
}

type Lbroute6 struct {
	Network            string `json:"network,omitempty"`
	Gatewayname        string `json:"gatewayname,omitempty"`
	Td                 int    `json:"td,omitempty"`
	Flags              string `json:"flags,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Lbvserverpacingpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Lbvserverpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Sc                     string `json:"sc,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Lbvserverresponderpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
}

type Lbgroup struct {
	Name                     string `json:"name,omitempty"`
	Persistencetype          string `json:"persistencetype,omitempty"`
	Persistencebackup        string `json:"persistencebackup,omitempty"`
	Backuppersistencetimeout int    `json:"backuppersistencetimeout,omitempty"`
	Persistmask              string `json:"persistmask,omitempty"`
	Cookiename               string `json:"cookiename,omitempty"`
	V6persistmasklen         int    `json:"v6persistmasklen,omitempty"`
	Cookiedomain             string `json:"cookiedomain,omitempty"`
	Timeout                  int    `json:"timeout,omitempty"`
	Rule                     string `json:"rule,omitempty"`
	Mastervserver            string `json:"mastervserver,omitempty"`
	Usevserverpersistency    string `json:"usevserverpersistency,omitempty"`
	Newname                  string `json:"newname,omitempty"`
	Td                       string `json:"td,omitempty"`
	Nextgenapiresource       string `json:"_nextgenapiresource,omitempty"`
}

type Lbvserverservicegroupbinding struct {
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Servicename      string `json:"servicename,omitempty"`
	Order            int    `json:"order,omitempty"`
	Name             string `json:"name,omitempty"`
	Weight           int    `json:"weight,omitempty"`
}

type Lbvservertrafficpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Name                   string `json:"name,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Lbvservervideooptimizationdetectionpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
}

type Lbwlmlbvserverbinding struct {
	Vservername string `json:"vservername,omitempty"`
	Wlmname     string `json:"wlmname,omitempty"`
}

type Lbwlmvserverbinding struct {
	Vservername string `json:"vservername,omitempty"`
	Wlmname     string `json:"wlmname,omitempty"`
}

type Lbvservercmppolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
}

type Lbvserverpolicy64binding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Lbpolicylabelpolicybindingbinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Lbvserverappflowpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
}

type Lbvserverauthorizationpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
}

type Lbvservernslogpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Sc                     string `json:"sc,omitempty"`
	Name                   string `json:"name,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Lbvservertmtrafficpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Name                   string `json:"name,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Order                  int    `json:"order,omitempty"`
}

type Lbvservervserverbinding struct {
	Cachevserver  string `json:"cachevserver,omitempty"`
	Policyname    string `json:"policyname,omitempty"`
	Cachetype     string `json:"cachetype,omitempty"`
	Priority      uint32 `json:"priority,omitempty"`
	Hits          uint32 `json:"hits,omitempty"`
	Pipolicyhits  uint32 `json:"pipolicyhits,omitempty"`
	Policysubtype uint32 `json:"policysubtype,omitempty"`
	Name          string `json:"name,omitempty"`
	Labelname     string `json:"labelname,omitempty"`
}

type Lbaction struct {
	Name               string `json:"name,omitempty"`
	Type               string `json:"type,omitempty"`
	Value              []int  `json:"value,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Newname            string `json:"newname,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Referencecount     string `json:"referencecount,omitempty"`
	Undefhits          string `json:"undefhits,omitempty"`
	Feature            string `json:"feature,omitempty"`
	Builtin            string `json:"builtin,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Lbglobalbinding struct {
}

type Lbgroupvserverbinding struct {
	Vservername string `json:"vservername,omitempty"`
	Name        string `json:"name,omitempty"`
}

type Lbmonitor struct {
	Monitorname                      string   `json:"monitorname,omitempty"`
	Type                             string   `json:"type,omitempty"`
	Action                           string   `json:"action,omitempty"`
	Respcode                         []string `json:"respcode,omitempty"`
	Httprequest                      string   `json:"httprequest,omitempty"`
	Rtsprequest                      string   `json:"rtsprequest,omitempty"`
	Customheaders                    string   `json:"customheaders,omitempty"`
	Maxforwards                      int      `json:"maxforwards,omitempty"`
	Sipmethod                        string   `json:"sipmethod,omitempty"`
	Sipuri                           string   `json:"sipuri,omitempty"`
	Sipreguri                        string   `json:"sipreguri,omitempty"`
	Send                             string   `json:"send,omitempty"`
	Recv                             string   `json:"recv,omitempty"`
	Query                            string   `json:"query,omitempty"`
	Querytype                        string   `json:"querytype,omitempty"`
	Scriptname                       string   `json:"scriptname,omitempty"`
	Scriptargs                       string   `json:"scriptargs,omitempty"`
	Secureargs                       string   `json:"secureargs,omitempty"`
	Dispatcherip                     string   `json:"dispatcherip,omitempty"`
	Dispatcherport                   int      `json:"dispatcherport,omitempty"`
	Username                         string   `json:"username,omitempty"`
	Password                         string   `json:"password,omitempty"`
	Secondarypassword                string   `json:"secondarypassword,omitempty"`
	Logonpointname                   string   `json:"logonpointname,omitempty"`
	Lasversion                       string   `json:"lasversion,omitempty"`
	Radkey                           string   `json:"radkey,omitempty"`
	Radnasid                         string   `json:"radnasid,omitempty"`
	Radnasip                         string   `json:"radnasip,omitempty"`
	Radaccounttype                   int      `json:"radaccounttype,omitempty"`
	Radframedip                      string   `json:"radframedip,omitempty"`
	Radapn                           string   `json:"radapn,omitempty"`
	Radmsisdn                        string   `json:"radmsisdn,omitempty"`
	Radaccountsession                string   `json:"radaccountsession,omitempty"`
	Lrtm                             string   `json:"lrtm,omitempty"`
	Deviation                        int      `json:"deviation"`
	Units1                           string   `json:"units1,omitempty"`
	Interval                         int      `json:"interval,omitempty"`
	Units3                           string   `json:"units3,omitempty"`
	Resptimeout                      int      `json:"resptimeout,omitempty"`
	Units4                           string   `json:"units4,omitempty"`
	Resptimeoutthresh                int      `json:"resptimeoutthresh,omitempty"`
	Retries                          int      `json:"retries,omitempty"`
	Failureretries                   int      `json:"failureretries,omitempty"`
	Alertretries                     int      `json:"alertretries,omitempty"`
	Successretries                   int      `json:"successretries,omitempty"`
	Downtime                         int      `json:"downtime,omitempty"`
	Units2                           string   `json:"units2,omitempty"`
	Destip                           string   `json:"destip,omitempty"`
	Destport                         int      `json:"destport,omitempty"`
	State                            string   `json:"state,omitempty"`
	Reverse                          string   `json:"reverse,omitempty"`
	Transparent                      string   `json:"transparent,omitempty"`
	Iptunnel                         string   `json:"iptunnel,omitempty"`
	Tos                              string   `json:"tos,omitempty"`
	Tosid                            int      `json:"tosid,omitempty"`
	Secure                           string   `json:"secure,omitempty"`
	Validatecred                     string   `json:"validatecred,omitempty"`
	Domain                           string   `json:"domain,omitempty"`
	Ipaddress                        []string `json:"ipaddress,omitempty"`
	Group                            string   `json:"group,omitempty"`
	Filename                         string   `json:"filename,omitempty"`
	Basedn                           string   `json:"basedn,omitempty"`
	Binddn                           string   `json:"binddn,omitempty"`
	Filter                           string   `json:"filter,omitempty"`
	Attribute                        string   `json:"attribute,omitempty"`
	Database                         string   `json:"database,omitempty"`
	Oraclesid                        string   `json:"oraclesid,omitempty"`
	Sqlquery                         string   `json:"sqlquery,omitempty"`
	Evalrule                         string   `json:"evalrule,omitempty"`
	Mssqlprotocolversion             string   `json:"mssqlprotocolversion,omitempty"`
	Snmpoid                          string   `json:"Snmpoid,omitempty"`
	Snmpcommunity                    string   `json:"snmpcommunity,omitempty"`
	Snmpthreshold                    string   `json:"snmpthreshold,omitempty"`
	Snmpversion                      string   `json:"snmpversion,omitempty"`
	Metrictable                      string   `json:"metrictable,omitempty"`
	Application                      string   `json:"application,omitempty"`
	Sitepath                         string   `json:"sitepath,omitempty"`
	Storename                        string   `json:"storename,omitempty"`
	Storefrontacctservice            string   `json:"storefrontacctservice,omitempty"`
	Hostname                         string   `json:"hostname,omitempty"`
	Netprofile                       string   `json:"netprofile,omitempty"`
	Originhost                       string   `json:"originhost,omitempty"`
	Originrealm                      string   `json:"originrealm,omitempty"`
	Hostipaddress                    string   `json:"hostipaddress,omitempty"`
	Vendorid                         int      `json:"vendorid,omitempty"`
	Productname                      string   `json:"productname,omitempty"`
	Firmwarerevision                 int      `json:"firmwarerevision,omitempty"`
	Authapplicationid                []int    `json:"authapplicationid,omitempty"`
	Acctapplicationid                []int    `json:"acctapplicationid,omitempty"`
	Inbandsecurityid                 string   `json:"inbandsecurityid,omitempty"`
	Supportedvendorids               []int    `json:"supportedvendorids,omitempty"`
	Vendorspecificvendorid           int      `json:"vendorspecificvendorid,omitempty"`
	Vendorspecificauthapplicationids []int    `json:"vendorspecificauthapplicationids,omitempty"`
	Vendorspecificacctapplicationids []int    `json:"vendorspecificacctapplicationids,omitempty"`
	Kcdaccount                       string   `json:"kcdaccount,omitempty"`
	Storedb                          string   `json:"storedb,omitempty"`
	Storefrontcheckbackendservices   string   `json:"storefrontcheckbackendservices,omitempty"`
	Trofscode                        int      `json:"trofscode,omitempty"`
	Trofsstring                      string   `json:"trofsstring,omitempty"`
	Sslprofile                       string   `json:"sslprofile,omitempty"`
	Mqttclientidentifier             string   `json:"mqttclientidentifier,omitempty"`
	Mqttversion                      int      `json:"mqttversion,omitempty"`
	Grpchealthcheck                  string   `json:"grpchealthcheck,omitempty"`
	Grpcstatuscode                   []int    `json:"grpcstatuscode,omitempty"`
	Grpcservicename                  string   `json:"grpcservicename,omitempty"`
	Metric                           string   `json:"metric,omitempty"`
	Metricthreshold                  int      `json:"metricthreshold,omitempty"`
	Metricweight                     int      `json:"metricweight,omitempty"`
	Servicename                      string   `json:"servicename,omitempty"`
	Servicegroupname                 string   `json:"servicegroupname,omitempty"`
	Lrtmconf                         string   `json:"lrtmconf,omitempty"`
	Lrtmconfstr                      string   `json:"lrtmconfstr,omitempty"`
	Dynamicresponsetimeout           string   `json:"dynamicresponsetimeout,omitempty"`
	Dynamicinterval                  string   `json:"dynamicinterval,omitempty"`
	Multimetrictable                 string   `json:"multimetrictable,omitempty"`
	Dupstate                         string   `json:"dup_state,omitempty"`
	Dupweight                        string   `json:"dup_weight,omitempty"`
	Weight                           string   `json:"weight,omitempty"`
	Nextgenapiresource               string   `json:"_nextgenapiresource,omitempty"`
}

type Lbmonitorbinding struct {
	Monitorname string `json:"monitorname,omitempty"`
}

type Lbmonitorsslcertkeybinding struct {
	Certkeyname string `json:"certkeyname,omitempty"`
	Ca          bool   `json:"ca,omitempty"`
	Crlcheck    string `json:"crlcheck,omitempty"`
	Ocspcheck   string `json:"ocspcheck,omitempty"`
	Monitorname string `json:"monitorname,omitempty"`
}

type Lbpolicylabelbinding struct {
	Labelname string `json:"labelname,omitempty"`
}

type Lbvserverappfwpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
}

type Lbpolicylbpolicylabelbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Lbvserverlbpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
}

type Lbvserverrewritepolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
}

type LBVServerServiceGroupMemberBinding struct {
	CookieIPPort      string `json:"cookieipport,omitempty"`
	CookieName        string `json:"cookiename,omitempty"`
	CurState          string `json:"curstate,omitempty"`
	DynamicWeight     string `json:"dynamicweight,omitempty"`
	Ipv46             string `json:"ipv46,omitempty"`
	Name              string `json:"name,omitempty"`
	Order             string `json:"order,omitempty"`
	OrderStr          string `json:"orderstr,omitempty"`
	Port              int    `json:"port,omitempty"`
	PreferredLocation string `json:"preferredlocation,omitempty"`
	ServiceGroupName  string `json:"servicegroupname,omitempty"`
	ServiceType       string `json:"servicetype,omitempty"`
	VServerid         string `json:"vserverid,omitempty"`
	Weight            string `json:"weight,omitempty"`
}

type Lbmonbindingsbinding struct {
	Monitorname string `json:"monitorname,omitempty"`
}

type Lbmonbindingsservicegroupbinding struct {
	Servicegroupname          string `json:"servicegroupname,omitempty"`
	Servicetype               string `json:"servicetype,omitempty"`
	Boundservicegroupsvrstate string `json:"boundservicegroupsvrstate,omitempty"`
	Monstate                  string `json:"monstate,omitempty"`
	Monitorname               string `json:"monitorname,omitempty"`
}

type Lbroute struct {
	Network            string `json:"network,omitempty"`
	Netmask            string `json:"netmask,omitempty"`
	Gatewayname        string `json:"gatewayname,omitempty"`
	Td                 int    `json:"td,omitempty"`
	Flags              string `json:"flags,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Lbvserverbinding struct {
	Name string `json:"name,omitempty"`
}

type Lbvserverpqpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Lbgrouplbvserverbinding struct {
	Vservername string `json:"vservername,omitempty"`
	Name        string `json:"name,omitempty"`
}

type Lbmonbindings struct {
	Monitorname               string `json:"monitorname,omitempty"`
	Type                      string `json:"type,omitempty"`
	State                     string `json:"state,omitempty"`
	Boundservicegroupsvrstate string `json:"boundservicegroupsvrstate,omitempty"`
	Nextgenapiresource        string `json:"_nextgenapiresource,omitempty"`
}

type Lbpolicylbvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Lbpolicylabel struct {
	Labelname              string `json:"labelname,omitempty"`
	Policylabeltype        string `json:"policylabeltype,omitempty"`
	Comment                string `json:"comment,omitempty"`
	Newname                string `json:"newname,omitempty"`
	Numpol                 string `json:"numpol,omitempty"`
	Hits                   string `json:"hits,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Nextgenapiresource     string `json:"_nextgenapiresource,omitempty"`
}

type Lbvserverbotpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
}

type Lbvservercontentinspectionpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
}

type Lbvserverfeopolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Name                   string `json:"name,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Order                  int    `json:"order,omitempty"`
}
