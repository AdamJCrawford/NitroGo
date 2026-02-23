package models

// gslb configuration structs
type GSLBServiceGroupBinding struct {
	GSLBServiceGroupGSLBServiceGroupMemberBinding        []interface{} `json:"gslbservicegroup_gslbservicegroupmember_binding,omitempty"`
	GSLBServiceGroupLBMonitorBinding                     []interface{} `json:"gslbservicegroup_lbmonitor_binding,omitempty"`
	GSLBServiceGroupServiceGroupEntityMonBindingsBinding []interface{} `json:"gslbservicegroup_servicegroupentitymonbindings_binding,omitempty"`
	ServiceGroupName                                     string        `json:"servicegroupname,omitempty"`
}

type GSLBServiceGroup struct {
	AppFlowLog                 string  `json:"appflowlog,omitempty"`
	AutoDelayedTROFS           string  `json:"autodelayedtrofs,omitempty"`
	AutoScale                  string  `json:"autoscale,omitempty"`
	CIP                        string  `json:"cip,omitempty"`
	CIPHeader                  string  `json:"cipheader,omitempty"`
	ClMonOwner                 int     `json:"clmonowner,omitempty"`
	ClMonView                  int     `json:"clmonview,omitempty"`
	CltTimeout                 int     `json:"clttimeout,omitempty"`
	Comment                    string  `json:"comment,omitempty"`
	Count                      float64 `json:"__count,omitempty"`
	Delay                      int     `json:"delay,omitempty"`
	DownStateFlush             string  `json:"downstateflush,omitempty"`
	DupWeight                  int     `json:"dup_weight,omitempty"`
	Graceful                   string  `json:"graceful,omitempty"`
	GroupCount                 int     `json:"groupcount,omitempty"`
	GSLB                       string  `json:"gslb,omitempty"`
	HashID                     int     `json:"hashid,omitempty"`
	HealthMonitor              string  `json:"healthmonitor,omitempty"`
	IncludeMembers             bool    `json:"includemembers,omitempty"`
	IP                         string  `json:"ip,omitempty"`
	MaxBandwidth               int     `json:"maxbandwidth,omitempty"`
	MaxClient                  int     `json:"maxclient,omitempty"`
	MonitorNameSvc             string  `json:"monitor_name_svc,omitempty"`
	MonStatCode                int     `json:"monstatcode,omitempty"`
	MonStatParam1              int     `json:"monstatparam1,omitempty"`
	MonStatParam2              int     `json:"monstatparam2,omitempty"`
	MonStatParam3              int     `json:"monstatparam3,omitempty"`
	MonThreshold               int     `json:"monthreshold,omitempty"`
	NewName                    string  `json:"newname,omitempty"`
	NextGenAPIResource         string  `json:"_nextgenapiresource,omitempty"`
	NoDefaultBindings          string  `json:"nodefaultbindings,omitempty"`
	NumOfConnections           int     `json:"numofconnections,omitempty"`
	Order                      int     `json:"order,omitempty"`
	Port                       int     `json:"port,omitempty"`
	PublicIP                   string  `json:"publicip,omitempty"`
	PublicPort                 int     `json:"publicport,omitempty"`
	ServerName                 string  `json:"servername,omitempty"`
	ServiceConfType            bool    `json:"serviceconftype,omitempty"`
	ServiceGroupEffectiveState string  `json:"servicegroupeffectivestate,omitempty"`
	ServiceGroupName           string  `json:"servicegroupname,omitempty"`
	ServiceIPStr               string  `json:"serviceipstr,omitempty"`
	ServiceType                string  `json:"servicetype,omitempty"`
	SiteName                   string  `json:"sitename,omitempty"`
	SitePersistence            string  `json:"sitepersistence,omitempty"`
	SitePrefix                 string  `json:"siteprefix,omitempty"`
	State                      string  `json:"state,omitempty"`
	StateChangeTimeMsec        int     `json:"statechangetimemsec,omitempty"`
	StateUpdateReason          int     `json:"stateupdatereason,omitempty"`
	SvrEffGSLBState            string  `json:"svreffgslbstate,omitempty"`
	SvrState                   string  `json:"svrstate,omitempty"`
	SvrTimeout                 int     `json:"svrtimeout,omitempty"`
	Value                      string  `json:"value,omitempty"`
	Weight                     int     `json:"weight,omitempty"`
}

type GSLBLDNSEntries struct {
	Count              float64       `json:"__count,omitempty"`
	IPAddress          string        `json:"ipaddress,omitempty"`
	Name               string        `json:"name,omitempty"`
	NextGenAPIResource string        `json:"_nextgenapiresource,omitempty"`
	NodeID             int           `json:"nodeid,omitempty"`
	NumSites           int           `json:"numsites,omitempty"`
	RTT                []interface{} `json:"rtt,omitempty"`
	SiteName           string        `json:"sitename,omitempty"`
	TTL                int           `json:"ttl,omitempty"`
}

type GSLBServiceGroupServiceGroupEntityMonBindingsBinding struct {
	HashID                     int    `json:"hashid,omitempty"`
	LastResponse               string `json:"lastresponse,omitempty"`
	MonitorName                string `json:"monitor_name,omitempty"`
	MonitorState               string `json:"monitor_state,omitempty"`
	MonitorCurrentFailedProbes int    `json:"monitorcurrentfailedprobes,omitempty"`
	MonitorTotalFailedProbes   int    `json:"monitortotalfailedprobes,omitempty"`
	MonitorTotalProbes         int    `json:"monitortotalprobes,omitempty"`
	Order                      int    `json:"order,omitempty"`
	Passive                    bool   `json:"passive,omitempty"`
	Port                       int    `json:"port,omitempty"`
	PublicIP                   string `json:"publicip,omitempty"`
	PublicPort                 int    `json:"publicport,omitempty"`
	ServiceGroupEntName2       string `json:"servicegroupentname2,omitempty"`
	ServiceGroupName           string `json:"servicegroupname,omitempty"`
	SitePrefix                 string `json:"siteprefix,omitempty"`
	State                      string `json:"state,omitempty"`
	Weight                     int    `json:"weight,omitempty"`
}

type GSLBVServerLBPolicyBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
}

type GSLBVServer struct {
	ActiveServices            int     `json:"activeservices,omitempty"`
	AppFlowLog                string  `json:"appflowlog,omitempty"`
	BackupIP                  string  `json:"backupip,omitempty"`
	BackupLBMethod            string  `json:"backuplbmethod,omitempty"`
	BackupSessionTimeout      int     `json:"backupsessiontimeout,omitempty"`
	BackupVServer             string  `json:"backupvserver,omitempty"`
	Comment                   string  `json:"comment,omitempty"`
	ConsiderEffectiveState    string  `json:"considereffectivestate,omitempty"`
	CookieDomain              string  `json:"cookie_domain,omitempty"`
	CookieTimeout             int     `json:"cookietimeout,omitempty"`
	Count                     float64 `json:"__count,omitempty"`
	CurrentActiveOrder        string  `json:"currentactiveorder,omitempty"`
	CurState                  string  `json:"curstate,omitempty"`
	DisablePrimaryOnDown      string  `json:"disableprimaryondown,omitempty"`
	DNSRecordType             string  `json:"dnsrecordtype,omitempty"`
	DomainName                string  `json:"domainname,omitempty"`
	DynamicWeight             string  `json:"dynamicweight,omitempty"`
	ECS                       string  `json:"ecs,omitempty"`
	ECSAddrValidation         string  `json:"ecsaddrvalidation,omitempty"`
	EDR                       string  `json:"edr,omitempty"`
	GotoPriorityExpression    string  `json:"gotopriorityexpression,omitempty"`
	Health                    int     `json:"health,omitempty"`
	IPType                    string  `json:"iptype,omitempty"`
	ISCName                   string  `json:"iscname,omitempty"`
	LBMethod                  string  `json:"lbmethod,omitempty"`
	LBRRReason                int     `json:"lbrrreason,omitempty"`
	MIR                       string  `json:"mir,omitempty"`
	Name                      string  `json:"name,omitempty"`
	Netmask                   string  `json:"netmask,omitempty"`
	NewName                   string  `json:"newname,omitempty"`
	NextGenAPIResource        string  `json:"_nextgenapiresource,omitempty"`
	NoDefaultBindings         string  `json:"nodefaultbindings,omitempty"`
	Order                     int     `json:"order,omitempty"`
	OrderThreshold            int     `json:"orderthreshold,omitempty"`
	PersistenceID             int     `json:"persistenceid,omitempty"`
	PersistenceType           string  `json:"persistencetype,omitempty"`
	PersistMask               string  `json:"persistmask,omitempty"`
	PolicyName                string  `json:"policyname,omitempty"`
	Priority                  int     `json:"priority,omitempty"`
	Rule                      string  `json:"rule,omitempty"`
	ServerName                string  `json:"servername,omitempty"`
	ServiceGroupName          string  `json:"servicegroupname,omitempty"`
	ServiceName               string  `json:"servicename,omitempty"`
	ServiceType               string  `json:"servicetype,omitempty"`
	SiteDomainTTL             int     `json:"sitedomainttl,omitempty"`
	SitePersistence           string  `json:"sitepersistence,omitempty"`
	SOBackupAction            string  `json:"sobackupaction,omitempty"`
	SOMethod                  string  `json:"somethod,omitempty"`
	SOPersistence             string  `json:"sopersistence,omitempty"`
	SOPersistenceTimeout      int     `json:"sopersistencetimeout,omitempty"`
	SOThreshold               int     `json:"sothreshold,omitempty"`
	State                     string  `json:"state,omitempty"`
	StateChangeTimeMsec       int     `json:"statechangetimemsec,omitempty"`
	StateChangeTimeSec        string  `json:"statechangetimesec,omitempty"`
	Status                    int     `json:"status,omitempty"`
	TicksSinceLastStateChange int     `json:"tickssincelaststatechange,omitempty"`
	Timeout                   int     `json:"timeout,omitempty"`
	ToggleOrder               string  `json:"toggleorder,omitempty"`
	Tolerance                 int     `json:"tolerance,omitempty"`
	TotalServices             int     `json:"totalservices,omitempty"`
	TTL                       int     `json:"ttl,omitempty"`
	TypeField                 string  `json:"type,omitempty"`
	V6NetmaskLen              int     `json:"v6netmasklen,omitempty"`
	V6PersistMaskLen          int     `json:"v6persistmasklen,omitempty"`
	VSvrBindSvcIP             string  `json:"vsvrbindsvcip,omitempty"`
	VSvrBindSvcPort           int     `json:"vsvrbindsvcport,omitempty"`
	Weight                    int     `json:"weight,omitempty"`
}

type GSLBDomainLBMonitorBinding struct {
	CustomHeaders              string `json:"customheaders,omitempty"`
	GRPCHealthCheck            string `json:"grpchealthcheck,omitempty"`
	GRPCServiceName            string `json:"grpcservicename,omitempty"`
	GRPCStatusCode             int    `json:"grpcstatuscode,omitempty"`
	HTTPRequest                string `json:"httprequest,omitempty"`
	IPTunnel                   string `json:"iptunnel,omitempty"`
	LastResponse               string `json:"lastresponse,omitempty"`
	MonitorCurrentFailedProbes int    `json:"monitorcurrentfailedprobes,omitempty"`
	MonitorName                string `json:"monitorname,omitempty"`
	MonitorTotalFailedProbes   int    `json:"monitortotalfailedprobes,omitempty"`
	MonitorTotalProbes         int    `json:"monitortotalprobes,omitempty"`
	MonStatCode                int    `json:"monstatcode,omitempty"`
	MonState                   string `json:"monstate,omitempty"`
	Name                       string `json:"name,omitempty"`
	RespCode                   string `json:"respcode,omitempty"`
	ResponseTime               int    `json:"responsetime,omitempty"`
	ServiceName                string `json:"servicename,omitempty"`
	VServerName                string `json:"vservername,omitempty"`
}

type GSLBParameter struct {
	AutomaticConfigSync         string   `json:"automaticconfigsync,omitempty"`
	Builtin                     []string `json:"builtin,omitempty"`
	DropLDNSReq                 string   `json:"dropldnsreq,omitempty"`
	Feature                     string   `json:"feature,omitempty"`
	Flags                       int      `json:"flags,omitempty"`
	GSLBConfigSyncMonitor       string   `json:"gslbconfigsyncmonitor,omitempty"`
	GSLBSvcStateDelayTime       int      `json:"gslbsvcstatedelaytime,omitempty"`
	GSLBSyncInterval            int      `json:"gslbsyncinterval,omitempty"`
	GSLBSyncLocFiles            string   `json:"gslbsynclocfiles,omitempty"`
	GSLBSyncMode                string   `json:"gslbsyncmode,omitempty"`
	GSLBSyncSaveConfigCommand   string   `json:"gslbsyncsaveconfigcommand,omitempty"`
	Incarnation                 int      `json:"incarnation,omitempty"`
	LDNSEntryTimeout            int      `json:"ldnsentrytimeout,omitempty"`
	LDNSMask                    string   `json:"ldnsmask,omitempty"`
	LDNSProbeOrder              []string `json:"ldnsprobeorder,omitempty"`
	MEPKeepaliveTimeout         int      `json:"mepkeepalivetimeout,omitempty"`
	NextGenAPIResource          string   `json:"_nextgenapiresource,omitempty"`
	OverridePersistencyForOrder string   `json:"overridepersistencyfororder,omitempty"`
	RTTTolerance                int      `json:"rtttolerance,omitempty"`
	SvcStateLearningTime        int      `json:"svcstatelearningtime,omitempty"`
	UndefAction                 string   `json:"undefaction,omitempty"`
	V6LDNSMaskLen               int      `json:"v6ldnsmasklen,omitempty"`
}

type GSLBVServerSpilloverPolicyBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
}

type GSLBVServerBinding struct {
	GSLBVServerGSLBDomainBinding             []interface{} `json:"gslbvserver_gslbdomain_binding,omitempty"`
	GSLBVServerGSLBServiceBinding            []interface{} `json:"gslbvserver_gslbservice_binding,omitempty"`
	GSLBVServerGSLBServiceGroupBinding       []interface{} `json:"gslbvserver_gslbservicegroup_binding,omitempty"`
	GSLBVServerGSLBServiceGroupMemberBinding []interface{} `json:"gslbvserver_gslbservicegroupmember_binding,omitempty"`
	GSLBVServerLBPolicyBinding               []interface{} `json:"gslbvserver_lbpolicy_binding,omitempty"`
	GSLBVServerSpilloverPolicyBinding        []interface{} `json:"gslbvserver_spilloverpolicy_binding,omitempty"`
	Name                                     string        `json:"name,omitempty"`
}

type GSLBServiceDNSViewBinding struct {
	ServiceName string `json:"servicename,omitempty"`
	ViewIP      string `json:"viewip,omitempty"`
	ViewName    string `json:"viewname,omitempty"`
}

type GSLBServiceGroupGSLBServiceGroupMemberBinding struct {
	Delay                     int    `json:"delay,omitempty"`
	Graceful                  string `json:"graceful,omitempty"`
	GSLBThreshold             int    `json:"gslbthreshold,omitempty"`
	HashID                    int    `json:"hashid,omitempty"`
	IP                        string `json:"ip,omitempty"`
	Order                     int    `json:"order,omitempty"`
	OrderStr                  string `json:"orderstr,omitempty"`
	Port                      int    `json:"port,omitempty"`
	PreferredLocation         string `json:"preferredlocation,omitempty"`
	PublicIP                  string `json:"publicip,omitempty"`
	PublicPort                int    `json:"publicport,omitempty"`
	ServerName                string `json:"servername,omitempty"`
	ServiceGroupName          string `json:"servicegroupname,omitempty"`
	SitePrefix                string `json:"siteprefix,omitempty"`
	State                     string `json:"state,omitempty"`
	StateChangeTimeSec        string `json:"statechangetimesec,omitempty"`
	SvrState                  string `json:"svrstate,omitempty"`
	Threshold                 string `json:"threshold,omitempty"`
	TicksSinceLastStateChange int    `json:"tickssincelaststatechange,omitempty"`
	TROFSDelay                int    `json:"trofsdelay,omitempty"`
	Weight                    int    `json:"weight,omitempty"`
}

type GSLBDomainGSLBServiceGroupBinding struct {
	Name             string `json:"name,omitempty"`
	Order            int    `json:"order,omitempty"`
	ServiceGroupName string `json:"servicegroupname,omitempty"`
}

type GSLBDomain struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
}

type GSLBVServerGSLBServiceGroupMemberBinding struct {
	CurState           string `json:"curstate,omitempty"`
	DynamicWeight      string `json:"dynamicweight,omitempty"`
	GSLBThreshold      int    `json:"gslbthreshold,omitempty"`
	IPAddress          string `json:"ipaddress,omitempty"`
	Name               string `json:"name,omitempty"`
	Order              int    `json:"order,omitempty"`
	OrderStr           string `json:"orderstr,omitempty"`
	Port               int    `json:"port,omitempty"`
	PreferredLocation  string `json:"preferredlocation,omitempty"`
	ServiceGroupName   string `json:"servicegroupname,omitempty"`
	ServiceType        string `json:"servicetype,omitempty"`
	SitePersistCookie  string `json:"sitepersistcookie,omitempty"`
	SvcSitePersistence string `json:"svcsitepersistence,omitempty"`
	SvrEffGSLBState    string `json:"svreffgslbstate,omitempty"`
	ThresholdValue     int    `json:"thresholdvalue,omitempty"`
	Weight             int    `json:"weight,omitempty"`
}

type GSLBSiteGSLBServiceBinding struct {
	CnameEntry  string `json:"cnameentry,omitempty"`
	IPAddress   string `json:"ipaddress,omitempty"`
	Port        int    `json:"port,omitempty"`
	ServiceName string `json:"servicename,omitempty"`
	ServiceType string `json:"servicetype,omitempty"`
	SiteName    string `json:"sitename,omitempty"`
	State       string `json:"state,omitempty"`
}

type GSLBSiteGSLBServiceGroupMemberBinding struct {
	IPAddress        string `json:"ipaddress,omitempty"`
	Port             int    `json:"port,omitempty"`
	ServiceGroupName string `json:"servicegroupname,omitempty"`
	ServiceType      string `json:"servicetype,omitempty"`
	SiteName         string `json:"sitename,omitempty"`
	State            string `json:"state,omitempty"`
}

type GSLBSiteBinding struct {
	GSLBSiteGSLBServiceBinding            []interface{} `json:"gslbsite_gslbservice_binding,omitempty"`
	GSLBSiteGSLBServiceGroupBinding       []interface{} `json:"gslbsite_gslbservicegroup_binding,omitempty"`
	GSLBSiteGSLBServiceGroupMemberBinding []interface{} `json:"gslbsite_gslbservicegroupmember_binding,omitempty"`
	SiteName                              string        `json:"sitename,omitempty"`
}

type GSLBDomainGSLBServiceBinding struct {
	CnameEntry       string `json:"cnameentry,omitempty"`
	CumulativeWeight int    `json:"cumulativeweight,omitempty"`
	DynamicConfWt    int    `json:"dynamicconfwt,omitempty"`
	GSLBThreshold    int    `json:"gslbthreshold,omitempty"`
	IPAddress        string `json:"ipaddress,omitempty"`
	Name             string `json:"name,omitempty"`
	Order            int    `json:"order,omitempty"`
	Port             int    `json:"port,omitempty"`
	ServiceName      string `json:"servicename,omitempty"`
	ServiceType      string `json:"servicetype,omitempty"`
	State            string `json:"state,omitempty"`
	SvrEffGSLBState  string `json:"svreffgslbstate,omitempty"`
	VServerName      string `json:"vservername,omitempty"`
	Weight           int    `json:"weight,omitempty"`
}

type GSLBVServerGSLBServiceGroupBinding struct {
	Name             string `json:"name,omitempty"`
	Order            int    `json:"order,omitempty"`
	ServiceGroupName string `json:"servicegroupname,omitempty"`
}

type GSLBSiteGSLBServiceGroupBinding struct {
	ServiceGroupName string `json:"servicegroupname,omitempty"`
	ServiceType      string `json:"servicetype,omitempty"`
	SiteName         string `json:"sitename,omitempty"`
}

type GSLBServiceGroupLBMonitorBinding struct {
	HashID           int    `json:"hashid,omitempty"`
	MonitorName      string `json:"monitor_name,omitempty"`
	MonState         string `json:"monstate,omitempty"`
	MonWeight        int    `json:"monweight,omitempty"`
	Order            int    `json:"order,omitempty"`
	Passive          bool   `json:"passive,omitempty"`
	Port             int    `json:"port,omitempty"`
	PublicIP         string `json:"publicip,omitempty"`
	PublicPort       int    `json:"publicport,omitempty"`
	ServiceGroupName string `json:"servicegroupname,omitempty"`
	SitePrefix       string `json:"siteprefix,omitempty"`
	State            string `json:"state,omitempty"`
	Weight           int    `json:"weight,omitempty"`
}

type GSLBVServerDomainBinding struct {
	BackupIP         string `json:"backupip,omitempty"`
	BackupIPFlag     bool   `json:"backupipflag,omitempty"`
	CookieDomain     string `json:"cookie_domain,omitempty"`
	CookieDomainFlag bool   `json:"cookie_domainflag,omitempty"`
	CookieTimeout    int    `json:"cookietimeout,omitempty"`
	DomainName       string `json:"domainname,omitempty"`
	Name             string `json:"name,omitempty"`
	Order            int    `json:"order,omitempty"`
	SiteDomainTTL    int    `json:"sitedomainttl,omitempty"`
	TTL              int    `json:"ttl,omitempty"`
}

type GSLBDomainBinding struct {
	GSLBDomainGSLBServiceBinding            []interface{} `json:"gslbdomain_gslbservice_binding,omitempty"`
	GSLBDomainGSLBServiceGroupBinding       []interface{} `json:"gslbdomain_gslbservicegroup_binding,omitempty"`
	GSLBDomainGSLBServiceGroupMemberBinding []interface{} `json:"gslbdomain_gslbservicegroupmember_binding,omitempty"`
	GSLBDomainGSLBVServerBinding            []interface{} `json:"gslbdomain_gslbvserver_binding,omitempty"`
	GSLBDomainLBMonitorBinding              []interface{} `json:"gslbdomain_lbmonitor_binding,omitempty"`
	Name                                    string        `json:"name,omitempty"`
}

type GSLBServiceBinding struct {
	GSLBServiceDNSViewBinding   []interface{} `json:"gslbservice_dnsview_binding,omitempty"`
	GSLBServiceLBMonitorBinding []interface{} `json:"gslbservice_lbmonitor_binding,omitempty"`
	ServiceName                 string        `json:"servicename,omitempty"`
}

type GSLBServiceLBMonitorBinding struct {
	FailedProbes               int    `json:"failedprobes,omitempty"`
	LastResponse               string `json:"lastresponse,omitempty"`
	MonitorName                string `json:"monitor_name,omitempty"`
	MonitorState               string `json:"monitor_state,omitempty"`
	MonitorCurrentFailedProbes int    `json:"monitorcurrentfailedprobes,omitempty"`
	MonitorTotalFailedProbes   int    `json:"monitortotalfailedprobes,omitempty"`
	MonitorTotalProbes         int    `json:"monitortotalprobes,omitempty"`
	MonStatCode                int    `json:"monstatcode,omitempty"`
	MonState                   string `json:"monstate,omitempty"`
	MonStatParam1              int    `json:"monstatparam1,omitempty"`
	MonStatParam2              int    `json:"monstatparam2,omitempty"`
	MonStatParam3              int    `json:"monstatparam3,omitempty"`
	ResponseTime               int    `json:"responsetime,omitempty"`
	ServiceName                string `json:"servicename,omitempty"`
	TotalFailedProbes          int    `json:"totalfailedprobes,omitempty"`
	Weight                     int    `json:"weight,omitempty"`
}

type GSLBDomainGSLBServiceGroupMemberBinding struct {
	GSLBThreshold    int    `json:"gslbthreshold,omitempty"`
	IPAddress        string `json:"ipaddress,omitempty"`
	Name             string `json:"name,omitempty"`
	Order            int    `json:"order,omitempty"`
	Port             int    `json:"port,omitempty"`
	ServiceGroupName string `json:"servicegroupname,omitempty"`
	ServiceType      string `json:"servicetype,omitempty"`
	SvrEffGSLBState  string `json:"svreffgslbstate,omitempty"`
	Weight           int    `json:"weight,omitempty"`
}

type GSLBService struct {
	AppFlowLog                string  `json:"appflowlog,omitempty"`
	CIP                       string  `json:"cip,omitempty"`
	CIPHeader                 string  `json:"cipheader,omitempty"`
	ClMonOwner                int     `json:"clmonowner,omitempty"`
	ClMonView                 int     `json:"clmonview,omitempty"`
	CltTimeout                int     `json:"clttimeout,omitempty"`
	CnameEntry                string  `json:"cnameentry,omitempty"`
	Comment                   string  `json:"comment,omitempty"`
	CookieTimeout             int     `json:"cookietimeout,omitempty"`
	Count                     float64 `json:"__count,omitempty"`
	DownStateFlush            string  `json:"downstateflush,omitempty"`
	GLSBSvcHealthDescr        string  `json:"glsbsvchealthdescr,omitempty"`
	GSLB                      string  `json:"gslb,omitempty"`
	GSLBSvcHealth             int     `json:"gslbsvchealth,omitempty"`
	GSLBSvcStats              int     `json:"gslbsvcstats,omitempty"`
	GSLBThreshold             int     `json:"gslbthreshold,omitempty"`
	HashID                    int     `json:"hashid,omitempty"`
	HealthMonitor             string  `json:"healthmonitor,omitempty"`
	IP                        string  `json:"ip,omitempty"`
	IPAddress                 string  `json:"ipaddress,omitempty"`
	MaxAAAUsers               int     `json:"maxaaausers,omitempty"`
	MaxBandwidth              int     `json:"maxbandwidth,omitempty"`
	MaxClient                 int     `json:"maxclient,omitempty"`
	MonitorNameSvc            string  `json:"monitor_name_svc,omitempty"`
	MonitorState              string  `json:"monitor_state,omitempty"`
	MonState                  string  `json:"monstate,omitempty"`
	MonThreshold              int     `json:"monthreshold,omitempty"`
	NAPTRDomainTTL            int     `json:"naptrdomainttl,omitempty"`
	NAPTROrder                int     `json:"naptrorder,omitempty"`
	NAPTRPreference           int     `json:"naptrpreference,omitempty"`
	NAPTRReplacement          string  `json:"naptrreplacement,omitempty"`
	NAPTRServices             string  `json:"naptrservices,omitempty"`
	NewName                   string  `json:"newname,omitempty"`
	NextGenAPIResource        string  `json:"_nextgenapiresource,omitempty"`
	NoDefaultBindings         string  `json:"nodefaultbindings,omitempty"`
	Port                      int     `json:"port,omitempty"`
	PreferredLocation         string  `json:"preferredlocation,omitempty"`
	PublicIP                  string  `json:"publicip,omitempty"`
	PublicPort                int     `json:"publicport,omitempty"`
	ServerName                string  `json:"servername,omitempty"`
	ServiceName               string  `json:"servicename,omitempty"`
	ServiceType               string  `json:"servicetype,omitempty"`
	SiteName                  string  `json:"sitename,omitempty"`
	SitePersistence           string  `json:"sitepersistence,omitempty"`
	SitePrefix                string  `json:"siteprefix,omitempty"`
	State                     string  `json:"state,omitempty"`
	StateChangeTimeSec        string  `json:"statechangetimesec,omitempty"`
	SvrEffGSLBState           string  `json:"svreffgslbstate,omitempty"`
	SvrState                  string  `json:"svrstate,omitempty"`
	SvrTimeout                int     `json:"svrtimeout,omitempty"`
	Threshold                 string  `json:"threshold,omitempty"`
	TicksSinceLastStateChange int     `json:"tickssincelaststatechange,omitempty"`
	ViewIP                    string  `json:"viewip,omitempty"`
	ViewName                  string  `json:"viewname,omitempty"`
	Weight                    int     `json:"weight,omitempty"`
}

type GSLBConfig struct {
	Command    string `json:"command,omitempty"`
	Debug      bool   `json:"debug,omitempty"`
	ForceSync  string `json:"forcesync,omitempty"`
	NoWarn     bool   `json:"nowarn,omitempty"`
	Preview    bool   `json:"preview,omitempty"`
	SaveConfig bool   `json:"saveconfig,omitempty"`
}

type GSLBRunningConfig struct {
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	Response           string `json:"response,omitempty"`
}

type GSLBSite struct {
	BackupParentList       []string `json:"backupparentlist,omitempty"`
	CLIP                   string   `json:"clip,omitempty"`
	Count                  float64  `json:"__count,omitempty"`
	CurBackupParentIP      string   `json:"curbackupparentip,omitempty"`
	MetricExchange         string   `json:"metricexchange,omitempty"`
	NAPTRReplacementSuffix string   `json:"naptrreplacementsuffix,omitempty"`
	NewName                string   `json:"newname,omitempty"`
	NextGenAPIResource     string   `json:"_nextgenapiresource,omitempty"`
	NWMetricExchange       string   `json:"nwmetricexchange,omitempty"`
	OldName                string   `json:"oldname,omitempty"`
	ParentSite             string   `json:"parentsite,omitempty"`
	PersistenceMEPStatus   string   `json:"persistencemepstatus,omitempty"`
	PublicCLIP             string   `json:"publicclip,omitempty"`
	PublicIP               string   `json:"publicip,omitempty"`
	SessionExchange        string   `json:"sessionexchange,omitempty"`
	SiteIPAddress          string   `json:"siteipaddress,omitempty"`
	SiteName               string   `json:"sitename,omitempty"`
	SitePassword           string   `json:"sitepassword,omitempty"`
	SiteState              string   `json:"sitestate,omitempty"`
	SiteType               string   `json:"sitetype,omitempty"`
	Status                 string   `json:"status,omitempty"`
	TriggerMonitor         string   `json:"triggermonitor,omitempty"`
	Version                int      `json:"version,omitempty"`
}

type GSLBSyncStatus struct {
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	Response           string `json:"response,omitempty"`
	Summary            bool   `json:"summary,omitempty"`
}

type GSLBLDNSEntry struct {
	IPAddress string `json:"ipaddress,omitempty"`
}

type GSLBVServerGSLBServiceBinding struct {
	CnameEntry         string `json:"cnameentry,omitempty"`
	CumulativeWeight   int    `json:"cumulativeweight,omitempty"`
	CurState           string `json:"curstate,omitempty"`
	DomainName         string `json:"domainname,omitempty"`
	DynamicConfWt      int    `json:"dynamicconfwt,omitempty"`
	GSLBBoundSvcType   string `json:"gslbboundsvctype,omitempty"`
	GSLBThreshold      int    `json:"gslbthreshold,omitempty"`
	IPAddress          string `json:"ipaddress,omitempty"`
	ISCName            string `json:"iscname,omitempty"`
	Name               string `json:"name,omitempty"`
	Order              int    `json:"order,omitempty"`
	OrderStr           string `json:"orderstr,omitempty"`
	Port               int    `json:"port,omitempty"`
	PreferredLocation  string `json:"preferredlocation,omitempty"`
	ServiceName        string `json:"servicename,omitempty"`
	SitePersistCookie  string `json:"sitepersistcookie,omitempty"`
	SvcSitePersistence string `json:"svcsitepersistence,omitempty"`
	SvrEffGSLBState    string `json:"svreffgslbstate,omitempty"`
	ThresholdValue     int    `json:"thresholdvalue,omitempty"`
	Weight             int    `json:"weight,omitempty"`
}

type GSLBDomainGSLBVServerBinding struct {
	BackupLBMethod     string `json:"backuplbmethod,omitempty"`
	CIP                string `json:"cip,omitempty"`
	CustomHeaders      string `json:"customheaders,omitempty"`
	DNSRecordType      string `json:"dnsrecordtype,omitempty"`
	DynamicWeight      string `json:"dynamicweight,omitempty"`
	EDR                string `json:"edr,omitempty"`
	LBMethod           string `json:"lbmethod,omitempty"`
	MIR                string `json:"mir,omitempty"`
	Name               string `json:"name,omitempty"`
	Netmask            string `json:"netmask,omitempty"`
	PersistenceID      int    `json:"persistenceid,omitempty"`
	PersistenceType    string `json:"persistencetype,omitempty"`
	PersistMask        string `json:"persistmask,omitempty"`
	ServiceType        string `json:"servicetype,omitempty"`
	SiteName           string `json:"sitename,omitempty"`
	SitePersistence    string `json:"sitepersistence,omitempty"`
	SitePrefix         string `json:"siteprefix,omitempty"`
	State              string `json:"state,omitempty"`
	StateChangeTimeSec string `json:"statechangetimesec,omitempty"`
	V6NetmaskLen       int    `json:"v6netmasklen,omitempty"`
	V6PersistMaskLen   int    `json:"v6persistmasklen,omitempty"`
	VServerName        string `json:"vservername,omitempty"`
}
