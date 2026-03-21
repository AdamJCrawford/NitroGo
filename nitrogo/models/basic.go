package models

// basic configuration structs
type ServerGSLBServiceGroupBinding struct {
	AppFlowLog       string `json:"appflowlog,omitempty"`
	BoundTD          int    `json:"boundtd,omitempty"`
	CIP              string `json:"cip,omitempty"`
	CIPHeader        string `json:"cipheader,omitempty"`
	CltTimeout       int    `json:"clttimeout,omitempty"`
	CustomServerID   string `json:"customserverid,omitempty"`
	DownStateFlush   string `json:"downstateflush,omitempty"`
	DupPort          int    `json:"dup_port,omitempty"`
	DupSvcType       string `json:"dup_svctype,omitempty"`
	MaxBandwidth     int    `json:"maxbandwidth,omitempty"`
	MaxClient        int    `json:"maxclient,omitempty"`
	MaxReq           int    `json:"maxreq,omitempty"`
	MonThreshold     int    `json:"monthreshold,omitempty"`
	Name             string `json:"name,omitempty"`
	Port             int    `json:"port,omitempty"`
	ServiceGroupName string `json:"servicegroupname,omitempty"`
	ServiceIPAddress string `json:"serviceipaddress,omitempty"`
	ServiceIPStr     string `json:"serviceipstr,omitempty"`
	SvcType          string `json:"svctype,omitempty"`
	SvrCfgFlags      int    `json:"svrcfgflags,omitempty"`
	SvrState         string `json:"svrstate,omitempty"`
	SvrTimeout       int    `json:"svrtimeout,omitempty"`
}

type ServerServiceBinding struct {
	Name             string `json:"name,omitempty"`
	Port             int    `json:"port,omitempty"`
	ServiceIPAddress string `json:"serviceipaddress,omitempty"`
	ServiceIPStr     string `json:"serviceipstr,omitempty"`
	ServiceName      string `json:"servicename,omitempty"`
	SvcType          string `json:"svctype,omitempty"`
	SvrState         string `json:"svrstate,omitempty"`
}

type ServiceGroupBindings struct {
	Count              float64 `json:"__count,omitempty"`
	IPAddress          string  `json:"ipaddress,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Port               int     `json:"port,omitempty"`
	ServiceGroupName   string  `json:"servicegroupname,omitempty"`
	State              string  `json:"state,omitempty"`
	SvrState           string  `json:"svrstate,omitempty"`
	VServerName        string  `json:"vservername,omitempty"`
}

type LocationParameter struct {
	Builtin            []string `json:"builtin,omitempty"`
	Context            string   `json:"context,omitempty"`
	Custom             int      `json:"custom,omitempty"`
	Custom6            int      `json:"custom6,omitempty"`
	DatabaseMode       string   `json:"databasemode,omitempty"`
	Entries            int      `json:"entries,omitempty"`
	Entries6           int      `json:"entries6,omitempty"`
	Errors             int      `json:"errors,omitempty"`
	Errors6            int      `json:"errors6,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Flags              int      `json:"flags,omitempty"`
	Flushing           string   `json:"flushing,omitempty"`
	Format             string   `json:"format,omitempty"`
	Format6            string   `json:"format6,omitempty"`
	Lines              int      `json:"lines,omitempty"`
	Lines6             int      `json:"lines6,omitempty"`
	Loading            string   `json:"loading,omitempty"`
	LocationFile       string   `json:"Locationfile,omitempty"`
	LocationFile6      string   `json:"locationfile6,omitempty"`
	MatchWildcardToAny string   `json:"matchwildcardtoany,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	Q1Label            string   `json:"q1label,omitempty"`
	Q2Label            string   `json:"q2label,omitempty"`
	Q3Label            string   `json:"q3label,omitempty"`
	Q4Label            string   `json:"q4label,omitempty"`
	Q5Label            string   `json:"q5label,omitempty"`
	Q6Label            string   `json:"q6label,omitempty"`
	Static             int      `json:"Static,omitempty"`
	Static6            int      `json:"static6,omitempty"`
	Status             int      `json:"status,omitempty"`
	Warnings           int      `json:"warnings,omitempty"`
	Warnings6          int      `json:"warnings6,omitempty"`
}

type LocationData struct {
}

type ServiceGroupServiceGroupEntityMonBindingsBinding struct {
	CustomServerID             string `json:"customserverid,omitempty"`
	DBSTTL                     int    `json:"dbsttl,omitempty"`
	HashID                     int    `json:"hashid,omitempty"`
	LastResponse               string `json:"lastresponse,omitempty"`
	MonitorName                string `json:"monitor_name,omitempty"`
	MonitorState               string `json:"monitor_state,omitempty"`
	MonitorCurrentFailedProbes int    `json:"monitorcurrentfailedprobes,omitempty"`
	MonitorTotalFailedProbes   int    `json:"monitortotalfailedprobes,omitempty"`
	MonitorTotalProbes         int    `json:"monitortotalprobes,omitempty"`
	NameServer                 string `json:"nameserver,omitempty"`
	Order                      int    `json:"order,omitempty"`
	Passive                    bool   `json:"passive,omitempty"`
	Port                       int    `json:"port,omitempty"`
	ResponseTime               int    `json:"responsetime,omitempty"`
	ServerID                   int    `json:"serverid,omitempty"`
	ServiceGroupEntName2       string `json:"servicegroupentname2,omitempty"`
	ServiceGroupName           string `json:"servicegroupname,omitempty"`
	State                      string `json:"state,omitempty"`
	Weight                     int    `json:"weight,omitempty"`
}

type ServiceGroupBinding struct {
	ServiceGroupLBMonitorBinding                     []any  `json:"servicegroup_lbmonitor_binding,omitempty"`
	ServiceGroupServiceGroupEntityMonBindingsBinding []any  `json:"servicegroup_servicegroupentitymonbindings_binding,omitempty"`
	ServiceGroupServiceGroupMemberBinding            []any  `json:"servicegroup_servicegroupmember_binding,omitempty"`
	ServiceGroupName                                 string `json:"servicegroupname,omitempty"`
}

type ServiceGroupLBMonitorBinding struct {
	CustomServerID   string `json:"customserverid,omitempty"`
	DBSTTL           int    `json:"dbsttl,omitempty"`
	HashID           int    `json:"hashid,omitempty"`
	MonitorName      string `json:"monitor_name,omitempty"`
	MonState         string `json:"monstate,omitempty"`
	MonWeight        int    `json:"monweight,omitempty"`
	NameServer       string `json:"nameserver,omitempty"`
	Order            int    `json:"order,omitempty"`
	Passive          bool   `json:"passive,omitempty"`
	Port             int    `json:"port,omitempty"`
	ServerID         int    `json:"serverid,omitempty"`
	ServiceGroupName string `json:"servicegroupname,omitempty"`
	State            string `json:"state,omitempty"`
	Weight           int    `json:"weight,omitempty"`
}

type ServiceBinding struct {
	Name                    string `json:"name,omitempty"`
	ServiceLBMonitorBinding []any  `json:"service_lbmonitor_binding,omitempty"`
}

type RADIUSNode struct {
	Count              float64 `json:"__count,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NodePrefix         string  `json:"nodeprefix,omitempty"`
	RADKey             string  `json:"radkey,omitempty"`
}

type SVCBindings struct {
	Count              float64 `json:"__count,omitempty"`
	IPAddress          string  `json:"ipaddress,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Port               int     `json:"port,omitempty"`
	ServiceName        string  `json:"servicename,omitempty"`
	SvrState           string  `json:"svrstate,omitempty"`
	VServerName        string  `json:"vservername,omitempty"`
}

type ServiceGroupServiceGroupMemberBinding struct {
	CustomServerID            string `json:"customserverid,omitempty"`
	DBSTTL                    int    `json:"dbsttl,omitempty"`
	Delay                     int    `json:"delay,omitempty"`
	Graceful                  string `json:"graceful,omitempty"`
	HashID                    int    `json:"hashid,omitempty"`
	IP                        string `json:"ip,omitempty"`
	NameServer                string `json:"nameserver,omitempty"`
	Order                     int    `json:"order,omitempty"`
	OrderStr                  string `json:"orderstr,omitempty"`
	Port                      int    `json:"port,omitempty"`
	ServerID                  int    `json:"serverid,omitempty"`
	ServerName                string `json:"servername,omitempty"`
	ServiceGroupName          string `json:"servicegroupname,omitempty"`
	State                     string `json:"state,omitempty"`
	StateChangeTimeSec        string `json:"statechangetimesec,omitempty"`
	SvcItmPriority            int    `json:"svcitmpriority,omitempty"`
	SvrState                  string `json:"svrstate,omitempty"`
	TicksSinceLastStateChange int    `json:"tickssincelaststatechange,omitempty"`
	TROFSDelay                int    `json:"trofsdelay,omitempty"`
	TROFSReason               string `json:"trofsreason,omitempty"`
	Weight                    int    `json:"weight,omitempty"`
}

type VServer struct {
	BackupVServer        string `json:"backupvserver,omitempty"`
	Cacheable            string `json:"cacheable,omitempty"`
	CltTimeout           int    `json:"clttimeout,omitempty"`
	Name                 string `json:"name,omitempty"`
	PushVServer          string `json:"pushvserver,omitempty"`
	RedirectURL          string `json:"redirecturl,omitempty"`
	SOMethod             string `json:"somethod,omitempty"`
	SOPersistence        string `json:"sopersistence,omitempty"`
	SOPersistenceTimeout int    `json:"sopersistencetimeout,omitempty"`
	SOThreshold          int    `json:"sothreshold,omitempty"`
}

type NSTrace struct {
	CapDropPkt         string   `json:"capdroppkt,omitempty"`
	CapSSLKeys         string   `json:"capsslkeys,omitempty"`
	DoRuntimeCleanup   string   `json:"doruntimecleanup,omitempty"`
	FileID             string   `json:"fileid,omitempty"`
	FileName           string   `json:"filename,omitempty"`
	FileSize           int      `json:"filesize,omitempty"`
	Filter             string   `json:"filter,omitempty"`
	InMemoryTrace      string   `json:"inmemorytrace,omitempty"`
	Link               string   `json:"link,omitempty"`
	Merge              string   `json:"merge,omitempty"`
	Mode               []string `json:"mode,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	NF                 int      `json:"nf,omitempty"`
	NodeID             int      `json:"nodeid,omitempty"`
	Nodes              []any    `json:"nodes,omitempty"`
	PerNIC             string   `json:"pernic,omitempty"`
	Scope              string   `json:"scope,omitempty"`
	Size               int      `json:"size,omitempty"`
	SkipLocalSSH       string   `json:"skiplocalssh,omitempty"`
	SkipRPC            string   `json:"skiprpc,omitempty"`
	State              string   `json:"state,omitempty"`
	Time               int      `json:"time,omitempty"`
	TraceBuffers       int      `json:"tracebuffers,omitempty"`
	TraceFormat        string   `json:"traceformat,omitempty"`
	TraceLocation      string   `json:"tracelocation,omitempty"`
}

type LocationFile struct {
	CurLocFileStatus   string `json:"curlocfilestatus,omitempty"`
	Format             string `json:"format,omitempty"`
	LocationFile       string `json:"Locationfile,omitempty"`
	LocFileStatusStr   string `json:"locfilestatusstr,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	PrevLocationFile   string `json:"prevlocationfile,omitempty"`
	PrevLocFileFormat  string `json:"prevlocfileformat,omitempty"`
	PrevLocFileStatus  string `json:"prevlocfilestatus,omitempty"`
	Src                string `json:"src,omitempty"`
}

type ServerGSLBServiceBinding struct {
	Name             string `json:"name,omitempty"`
	Port             int    `json:"port,omitempty"`
	ServiceIPAddress string `json:"serviceipaddress,omitempty"`
	ServiceIPStr     string `json:"serviceipstr,omitempty"`
	ServiceName      string `json:"servicename,omitempty"`
	SvcType          string `json:"svctype,omitempty"`
	SvrState         string `json:"svrstate,omitempty"`
}

type ExtendedMemoryParam struct {
	MaxMemLimit        int    `json:"maxmemlimit,omitempty"`
	MemLimit           int    `json:"memlimit,omitempty"`
	MemLimitActive     int    `json:"memlimitactive,omitempty"`
	MinRequiredMemory  int    `json:"minrequiredmemory,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
}

type Service struct {
	AccessDown                   string   `json:"accessdown,omitempty"`
	All                          bool     `json:"all,omitempty"`
	AppFlowLog                   string   `json:"appflowlog,omitempty"`
	Builtin                      []string `json:"builtin,omitempty"`
	Cacheable                    string   `json:"cacheable,omitempty"`
	CacheType                    string   `json:"cachetype,omitempty"`
	CIP                          string   `json:"cip,omitempty"`
	CIPHeader                    string   `json:"cipheader,omitempty"`
	CKA                          string   `json:"cka,omitempty"`
	ClearTextPort                int      `json:"cleartextport,omitempty"`
	ClMonOwner                   int      `json:"clmonowner,omitempty"`
	ClMonView                    int      `json:"clmonview,omitempty"`
	CltTimeout                   int      `json:"clttimeout,omitempty"`
	CMP                          string   `json:"cmp,omitempty"`
	Comment                      string   `json:"comment,omitempty"`
	ContentInspectionProfileName string   `json:"contentinspectionprofilename,omitempty"`
	Count                        float64  `json:"__count,omitempty"`
	CustomServerID               string   `json:"customserverid,omitempty"`
	Delay                        int      `json:"delay,omitempty"`
	DNSProfileName               string   `json:"dnsprofilename,omitempty"`
	DownStateFlush               string   `json:"downstateflush,omitempty"`
	DupState                     string   `json:"dup_state,omitempty"`
	Feature                      string   `json:"feature,omitempty"`
	Graceful                     string   `json:"graceful,omitempty"`
	GSLB                         string   `json:"gslb,omitempty"`
	HashID                       int      `json:"hashid,omitempty"`
	HealthMonitor                string   `json:"healthmonitor,omitempty"`
	HTTPProfileName              string   `json:"httpprofilename,omitempty"`
	Internal                     bool     `json:"Internal,omitempty"`
	IP                           string   `json:"ip,omitempty"`
	IPAddress                    string   `json:"ipaddress,omitempty"`
	LastResponse                 string   `json:"lastresponse,omitempty"`
	MaxBandwidth                 int      `json:"maxbandwidth,omitempty"`
	MaxClient                    int      `json:"maxclient,omitempty"`
	MaxReq                       int      `json:"maxreq,omitempty"`
	MonConnectionClose           string   `json:"monconnectionclose,omitempty"`
	MonitorNameSvc               string   `json:"monitor_name_svc,omitempty"`
	MonitorState                 string   `json:"monitor_state,omitempty"`
	MonStatCode                  int      `json:"monstatcode,omitempty"`
	MonStatParam1                int      `json:"monstatparam1,omitempty"`
	MonStatParam2                int      `json:"monstatparam2,omitempty"`
	MonStatParam3                int      `json:"monstatparam3,omitempty"`
	MonThreshold                 int      `json:"monthreshold,omitempty"`
	MonUserStatusMesg            string   `json:"monuserstatusmesg,omitempty"`
	Name                         string   `json:"name,omitempty"`
	NetProfile                   string   `json:"netprofile,omitempty"`
	NewName                      string   `json:"newname,omitempty"`
	NextGenAPIResource           string   `json:"_nextgenapiresource,omitempty"`
	NoDefaultBindings            string   `json:"nodefaultbindings,omitempty"`
	NumOfConnections             int      `json:"numofconnections,omitempty"`
	OracleServerVersion          string   `json:"oracleserverversion,omitempty"`
	PathMonitor                  string   `json:"pathmonitor,omitempty"`
	PathMonitorIndv              string   `json:"pathmonitorindv,omitempty"`
	PolicyName                   string   `json:"policyname,omitempty"`
	Port                         int      `json:"port,omitempty"`
	ProcessLocal                 string   `json:"processlocal,omitempty"`
	PublicIP                     string   `json:"publicip,omitempty"`
	PublicPort                   int      `json:"publicport,omitempty"`
	QUICProfileName              string   `json:"quicprofilename,omitempty"`
	ResponseTime                 int      `json:"responsetime,omitempty"`
	RTSPSessionIDRemap           string   `json:"rtspsessionidremap,omitempty"`
	ServerID                     int      `json:"serverid,omitempty"`
	ServerName                   string   `json:"servername,omitempty"`
	ServiceConfType              bool     `json:"serviceconftype,omitempty"`
	ServiceConfType2             string   `json:"serviceconftype2,omitempty"`
	ServiceIPStr                 string   `json:"serviceipstr,omitempty"`
	ServiceType                  string   `json:"servicetype,omitempty"`
	SP                           string   `json:"sp,omitempty"`
	State                        string   `json:"state,omitempty"`
	StateChangeTimeMsec          int      `json:"statechangetimemsec,omitempty"`
	StateChangeTimeSec           string   `json:"statechangetimesec,omitempty"`
	StateUpdateReason            int      `json:"stateupdatereason,omitempty"`
	SvrState                     string   `json:"svrstate,omitempty"`
	SvrTimeout                   int      `json:"svrtimeout,omitempty"`
	TCPB                         string   `json:"tcpb,omitempty"`
	TCPProfileName               string   `json:"tcpprofilename,omitempty"`
	TD                           int      `json:"td,omitempty"`
	TicksSinceLastStateChange    int      `json:"tickssincelaststatechange,omitempty"`
	UseProxyPort                 string   `json:"useproxyport,omitempty"`
	USIP                         string   `json:"usip,omitempty"`
	Value                        string   `json:"value,omitempty"`
	Weight                       int      `json:"weight,omitempty"`
}

type ServerServiceGroupBinding struct {
	AppFlowLog           string `json:"appflowlog,omitempty"`
	BoundTD              int    `json:"boundtd,omitempty"`
	Cacheable            string `json:"cacheable,omitempty"`
	CIP                  string `json:"cip,omitempty"`
	CIPHeader            string `json:"cipheader,omitempty"`
	CKA                  string `json:"cka,omitempty"`
	CltTimeout           int    `json:"clttimeout,omitempty"`
	CMP                  string `json:"cmp,omitempty"`
	CustomServerID       string `json:"customserverid,omitempty"`
	DownStateFlush       string `json:"downstateflush,omitempty"`
	DupPort              int    `json:"dup_port,omitempty"`
	DupSvcType           string `json:"dup_svctype,omitempty"`
	MaxBandwidth         int    `json:"maxbandwidth,omitempty"`
	MaxClient            int    `json:"maxclient,omitempty"`
	MaxReq               int    `json:"maxreq,omitempty"`
	MonThreshold         int    `json:"monthreshold,omitempty"`
	Name                 string `json:"name,omitempty"`
	Port                 int    `json:"port,omitempty"`
	ServiceGroupEntName2 string `json:"servicegroupentname2,omitempty"`
	ServiceGroupName     string `json:"servicegroupname,omitempty"`
	ServiceIPAddress     string `json:"serviceipaddress,omitempty"`
	ServiceIPStr         string `json:"serviceipstr,omitempty"`
	SP                   string `json:"sp,omitempty"`
	SvcItmActSvcs        int    `json:"svcitmactsvcs,omitempty"`
	SvcItmBoundSvcs      int    `json:"svcitmboundsvcs,omitempty"`
	SvcItmPriority       int    `json:"svcitmpriority,omitempty"`
	SvcType              string `json:"svctype,omitempty"`
	SvrCfgFlags          int    `json:"svrcfgflags,omitempty"`
	SvrState             string `json:"svrstate,omitempty"`
	SvrTimeout           int    `json:"svrtimeout,omitempty"`
	TCPB                 string `json:"tcpb,omitempty"`
	USIP                 string `json:"usip,omitempty"`
	Weight               int    `json:"weight,omitempty"`
}

type DBSMonitors struct {
}

type LocationFile6 struct {
	Count              float64 `json:"__count,omitempty"`
	CurLocFileStatus   string  `json:"curlocfilestatus,omitempty"`
	Format             string  `json:"format,omitempty"`
	LocationFile       string  `json:"Locationfile,omitempty"`
	LocFileStatusStr   string  `json:"locfilestatusstr,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	PrevLocationFile   string  `json:"prevlocationfile,omitempty"`
	PrevLocFileFormat  string  `json:"prevlocfileformat,omitempty"`
	PrevLocFileStatus  string  `json:"prevlocfilestatus,omitempty"`
	Src                string  `json:"src,omitempty"`
}

type ServiceGroup struct {
	AppFlowLog                 string  `json:"appflowlog,omitempty"`
	AutoDelayedTROFS           string  `json:"autodelayedtrofs,omitempty"`
	AutoDisableDelay           int     `json:"autodisabledelay,omitempty"`
	AutoDisableGraceful        string  `json:"autodisablegraceful,omitempty"`
	AutoScale                  string  `json:"autoscale,omitempty"`
	Bootstrap                  string  `json:"bootstrap,omitempty"`
	Cacheable                  string  `json:"cacheable,omitempty"`
	CacheType                  string  `json:"cachetype,omitempty"`
	CIP                        string  `json:"cip,omitempty"`
	CIPHeader                  string  `json:"cipheader,omitempty"`
	CKA                        string  `json:"cka,omitempty"`
	ClMonOwner                 int     `json:"clmonowner,omitempty"`
	ClMonView                  int     `json:"clmonview,omitempty"`
	CltTimeout                 int     `json:"clttimeout,omitempty"`
	CMP                        string  `json:"cmp,omitempty"`
	Comment                    string  `json:"comment,omitempty"`
	Count                      float64 `json:"__count,omitempty"`
	CustomServerID             string  `json:"customserverid,omitempty"`
	DBSTTL                     int     `json:"dbsttl,omitempty"`
	Delay                      int     `json:"delay,omitempty"`
	DownStateFlush             string  `json:"downstateflush,omitempty"`
	DupWeight                  int     `json:"dup_weight,omitempty"`
	Graceful                   string  `json:"graceful,omitempty"`
	GroupCount                 int     `json:"groupcount,omitempty"`
	HashID                     int     `json:"hashid,omitempty"`
	HealthMonitor              string  `json:"healthmonitor,omitempty"`
	HTTPProfileName            string  `json:"httpprofilename,omitempty"`
	IncludeMembers             bool    `json:"includemembers,omitempty"`
	IP                         string  `json:"ip,omitempty"`
	MaxBandwidth               int     `json:"maxbandwidth,omitempty"`
	MaxClient                  int     `json:"maxclient,omitempty"`
	MaxReq                     int     `json:"maxreq,omitempty"`
	MemberPort                 int     `json:"memberport,omitempty"`
	MonConnectionClose         string  `json:"monconnectionclose,omitempty"`
	MonitorNameSvc             string  `json:"monitor_name_svc,omitempty"`
	MonStatCode                int     `json:"monstatcode,omitempty"`
	MonStatParam1              int     `json:"monstatparam1,omitempty"`
	MonStatParam2              int     `json:"monstatparam2,omitempty"`
	MonStatParam3              int     `json:"monstatparam3,omitempty"`
	MonThreshold               int     `json:"monthreshold,omitempty"`
	MonUserStatusMesg          string  `json:"monuserstatusmesg,omitempty"`
	NameServer                 string  `json:"nameserver,omitempty"`
	NetProfile                 string  `json:"netprofile,omitempty"`
	NewName                    string  `json:"newname,omitempty"`
	NextGenAPIResource         string  `json:"_nextgenapiresource,omitempty"`
	NoDefaultBindings          string  `json:"nodefaultbindings,omitempty"`
	NumOfConnections           int     `json:"numofconnections,omitempty"`
	Order                      int     `json:"order,omitempty"`
	PathMonitor                string  `json:"pathmonitor,omitempty"`
	PathMonitorIndv            string  `json:"pathmonitorindv,omitempty"`
	Port                       int     `json:"port,omitempty"`
	QUICProfileName            string  `json:"quicprofilename,omitempty"`
	RTSPSessionIDRemap         string  `json:"rtspsessionidremap,omitempty"`
	ServerID                   int     `json:"serverid,omitempty"`
	ServerName                 string  `json:"servername,omitempty"`
	ServiceConfType            bool    `json:"serviceconftype,omitempty"`
	ServiceGroupEffectiveState string  `json:"servicegroupeffectivestate,omitempty"`
	ServiceGroupName           string  `json:"servicegroupname,omitempty"`
	ServiceIPStr               string  `json:"serviceipstr,omitempty"`
	ServiceType                string  `json:"servicetype,omitempty"`
	SP                         string  `json:"sp,omitempty"`
	State                      string  `json:"state,omitempty"`
	StateChangeTimeMsec        int     `json:"statechangetimemsec,omitempty"`
	StateUpdateReason          int     `json:"stateupdatereason,omitempty"`
	SvcItmActSvcs              int     `json:"svcitmactsvcs,omitempty"`
	SvcItmBoundSvcs            int     `json:"svcitmboundsvcs,omitempty"`
	SvrState                   string  `json:"svrstate,omitempty"`
	SvrTimeout                 int     `json:"svrtimeout,omitempty"`
	TCPB                       string  `json:"tcpb,omitempty"`
	TCPProfileName             string  `json:"tcpprofilename,omitempty"`
	TD                         int     `json:"td,omitempty"`
	TopicName                  string  `json:"topicname,omitempty"`
	UseProxyPort               string  `json:"useproxyport,omitempty"`
	USIP                       string  `json:"usip,omitempty"`
	Value                      string  `json:"value,omitempty"`
	Weight                     int     `json:"weight,omitempty"`
}

type ServerBinding struct {
	Name                          string `json:"name,omitempty"`
	ServerGSLBServiceBinding      []any  `json:"server_gslbservice_binding,omitempty"`
	ServerGSLBServiceGroupBinding []any  `json:"server_gslbservicegroup_binding,omitempty"`
	ServerServiceBinding          []any  `json:"server_service_binding,omitempty"`
	ServerServiceGroupBinding     []any  `json:"server_servicegroup_binding,omitempty"`
}

type ServiceLBMonitorBinding struct {
	DupState                   string `json:"dup_state,omitempty"`
	DupWeight                  int    `json:"dup_weight,omitempty"`
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
	Name                       string `json:"name,omitempty"`
	Passive                    bool   `json:"passive,omitempty"`
	ResponseTime               int    `json:"responsetime,omitempty"`
	TotalFailedProbes          int    `json:"totalfailedprobes,omitempty"`
	TotalProbes                int    `json:"totalprobes,omitempty"`
	Weight                     int    `json:"weight,omitempty"`
}

type ServiceGroupServiceGroupMemberListBinding struct {
	FailedMembers    []any  `json:"failedmembers,omitempty"`
	Members          []any  `json:"members,omitempty"`
	ServiceGroupName string `json:"servicegroupname,omitempty"`
}

type Server struct {
	AutoScale                 string  `json:"autoscale,omitempty"`
	Cacheable                 string  `json:"cacheable,omitempty"`
	CKA                       string  `json:"cka,omitempty"`
	CMP                       string  `json:"cmp,omitempty"`
	Comment                   string  `json:"comment,omitempty"`
	Count                     float64 `json:"__count,omitempty"`
	Delay                     int     `json:"delay,omitempty"`
	Domain                    string  `json:"domain,omitempty"`
	DomainResolveNow          bool    `json:"domainresolvenow,omitempty"`
	DomainResolveRetry        int     `json:"domainresolveretry,omitempty"`
	Graceful                  string  `json:"graceful,omitempty"`
	Internal                  bool    `json:"Internal,omitempty"`
	IPAddress                 string  `json:"ipaddress,omitempty"`
	IPv6Address               string  `json:"ipv6address,omitempty"`
	Name                      string  `json:"name,omitempty"`
	NewName                   string  `json:"newname,omitempty"`
	NextGenAPIResource        string  `json:"_nextgenapiresource,omitempty"`
	QueryType                 string  `json:"querytype,omitempty"`
	SP                        string  `json:"sp,omitempty"`
	State                     string  `json:"state,omitempty"`
	StateChangeTimeSec        string  `json:"statechangetimesec,omitempty"`
	TCPB                      string  `json:"tcpb,omitempty"`
	TD                        int     `json:"td,omitempty"`
	TicksSinceLastStateChange int     `json:"tickssincelaststatechange,omitempty"`
	TranslationIP             string  `json:"translationip,omitempty"`
	TranslationMask           string  `json:"translationmask,omitempty"`
	USIP                      string  `json:"usip,omitempty"`
}

type Location struct {
	Count              float64 `json:"__count,omitempty"`
	IPFrom             string  `json:"ipfrom,omitempty"`
	IPTo               string  `json:"ipto,omitempty"`
	Latitude           int     `json:"latitude,omitempty"`
	Longitude          int     `json:"longitude,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	PreferredLocation  string  `json:"preferredlocation,omitempty"`
	Q1Label            string  `json:"q1label,omitempty"`
	Q2Label            string  `json:"q2label,omitempty"`
	Q3Label            string  `json:"q3label,omitempty"`
	Q4Label            string  `json:"q4label,omitempty"`
	Q5Label            string  `json:"q5label,omitempty"`
	Q6Label            string  `json:"q6label,omitempty"`
}

type Reporting struct {
	State string `json:"state,omitempty"`
}
