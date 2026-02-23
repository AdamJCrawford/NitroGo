package models

// lsn configuration structs
type LSNGroupLSNLogProfileBinding struct {
	GroupName      string `json:"groupname,omitempty"`
	LogProfileName string `json:"logprofilename,omitempty"`
}

type LSNClientNSACLBinding struct {
	ACLName    string `json:"aclname,omitempty"`
	ClientName string `json:"clientname,omitempty"`
	TD         int    `json:"td,omitempty"`
}

type LSNIP6Profile struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NatPrefix          string  `json:"natprefix,omitempty"`
	Network6           string  `json:"network6,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	TypeField          string  `json:"type,omitempty"`
}

type LSNGroupLSNHTTPHdrLogProfileBinding struct {
	GroupName             string `json:"groupname,omitempty"`
	HTTPHdrLogProfileName string `json:"httphdrlogprofilename,omitempty"`
}

type LSNGroupLSNTransportProfileBinding struct {
	GroupName            string `json:"groupname,omitempty"`
	TransportProfileName string `json:"transportprofilename,omitempty"`
}

type LSNAppsAttributes struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Port               string  `json:"port,omitempty"`
	SessionTimeout     int     `json:"sessiontimeout,omitempty"`
	TransportProtocol  string  `json:"transportprotocol,omitempty"`
}

type LSNGroupLSNRTSPALGProfileBinding struct {
	GroupName          string `json:"groupname,omitempty"`
	RTSPALGProfileName string `json:"rtspalgprofilename,omitempty"`
}

type LSNClient struct {
	ClientName         string  `json:"clientname,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
}

type LSNClientNSACL6Binding struct {
	ACL6Name   string `json:"acl6name,omitempty"`
	ClientName string `json:"clientname,omitempty"`
	TD         int    `json:"td,omitempty"`
}

type LSNParameter struct {
	MaxMemLimit          int    `json:"maxmemlimit,omitempty"`
	MemLimit             int    `json:"memlimit,omitempty"`
	MemLimitActive       int    `json:"memlimitactive,omitempty"`
	NextGenAPIResource   string `json:"_nextgenapiresource,omitempty"`
	SessionSync          string `json:"sessionsync,omitempty"`
	SubscrSessionRemoval string `json:"subscrsessionremoval,omitempty"`
}

type LSNClientNetworkBinding struct {
	ClientName string `json:"clientname,omitempty"`
	Netmask    string `json:"netmask,omitempty"`
	Network    string `json:"network,omitempty"`
	TD         int    `json:"td,omitempty"`
}

type LSNClientNetwork6Binding struct {
	ClientName string `json:"clientname,omitempty"`
	Netmask    string `json:"netmask,omitempty"`
	Network    string `json:"network,omitempty"`
	Network6   string `json:"network6,omitempty"`
	TD         int    `json:"td,omitempty"`
}

type LSNAppsProfilePortBinding struct {
	AppsProfileName string `json:"appsprofilename,omitempty"`
	LSNPort         string `json:"lsnport,omitempty"`
}

type LSNSession struct {
	ClientName         string  `json:"clientname,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	DestIP             string  `json:"destip,omitempty"`
	DestPort           int     `json:"destport,omitempty"`
	DstTD              int     `json:"dsttd,omitempty"`
	IPv6Address        string  `json:"ipv6address,omitempty"`
	NatIP              string  `json:"natip,omitempty"`
	NatPort            int     `json:"natport,omitempty"`
	NatPort2           int     `json:"natport2,omitempty"`
	NatPrefix          string  `json:"natprefix,omitempty"`
	NatType            string  `json:"nattype,omitempty"`
	Netmask            string  `json:"netmask,omitempty"`
	Network            string  `json:"network,omitempty"`
	Network6           string  `json:"network6,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NodeID             int     `json:"nodeid,omitempty"`
	SessionEstDir      string  `json:"sessionestdir,omitempty"`
	SrcTD              int     `json:"srctd,omitempty"`
	SubscrIP           string  `json:"subscrip,omitempty"`
	SubscrPort         int     `json:"subscrport,omitempty"`
	TD                 int     `json:"td,omitempty"`
	TransportProtocol  string  `json:"transportprotocol,omitempty"`
}

type LSNSIPALGProfile struct {
	Count                  float64 `json:"__count,omitempty"`
	DataSessionIdleTimeout int     `json:"datasessionidletimeout,omitempty"`
	NextGenAPIResource     string  `json:"_nextgenapiresource,omitempty"`
	OpenContactPinhole     string  `json:"opencontactpinhole,omitempty"`
	OpenRecordRoutePinhole string  `json:"openrecordroutepinhole,omitempty"`
	OpenRegisterPinhole    string  `json:"openregisterpinhole,omitempty"`
	OpenRoutePinhole       string  `json:"openroutepinhole,omitempty"`
	OpenViaPinhole         string  `json:"openviapinhole,omitempty"`
	RegistrationTimeout    int     `json:"registrationtimeout,omitempty"`
	RPort                  string  `json:"rport,omitempty"`
	SIPALGProfileName      string  `json:"sipalgprofilename,omitempty"`
	SIPDstPortRange        string  `json:"sipdstportrange,omitempty"`
	SIPSessionTimeout      int     `json:"sipsessiontimeout,omitempty"`
	SIPSrcPortRange        string  `json:"sipsrcportrange,omitempty"`
	SIPTransportProtocol   string  `json:"siptransportprotocol,omitempty"`
}

type LSNRTSPALGSessionBinding struct {
	LSNRTSPALGSessionDataChannelBinding []interface{} `json:"lsnrtspalgsession_datachannel_binding,omitempty"`
	SessionID                           string        `json:"sessionid,omitempty"`
}

type LSNLogProfile struct {
	AnalyticsProfile   string  `json:"analyticsprofile,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	LogCompact         string  `json:"logcompact,omitempty"`
	LogIPFIX           string  `json:"logipfix,omitempty"`
	LogProfileName     string  `json:"logprofilename,omitempty"`
	LogSessDeletion    string  `json:"logsessdeletion,omitempty"`
	LogSubscrInfo      string  `json:"logsubscrinfo,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
}

type LSNGroupIPSECALGProfileBinding struct {
	GroupName       string `json:"groupname,omitempty"`
	IPSECALGProfile string `json:"ipsecalgprofile,omitempty"`
}

type LSNGroup struct {
	AllocPolicy        string  `json:"allocpolicy,omitempty"`
	ClientName         string  `json:"clientname,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	FTP                string  `json:"ftp,omitempty"`
	FTPCM              string  `json:"ftpcm,omitempty"`
	GroupID            int     `json:"groupid,omitempty"`
	GroupName          string  `json:"groupname,omitempty"`
	IP6Profile         string  `json:"ip6profile,omitempty"`
	Logging            string  `json:"logging,omitempty"`
	NatType            string  `json:"nattype,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	PortBlockSize      int     `json:"portblocksize,omitempty"`
	PPTP               string  `json:"pptp,omitempty"`
	RTSPALG            string  `json:"rtspalg,omitempty"`
	SessionLogging     string  `json:"sessionlogging,omitempty"`
	SessionSync        string  `json:"sessionsync,omitempty"`
	SIPALG             string  `json:"sipalg,omitempty"`
	SNMPTrapLimit      int     `json:"snmptraplimit,omitempty"`
}

type LSNStatic struct {
	Count              float64 `json:"__count,omitempty"`
	DestIP             string  `json:"destip,omitempty"`
	DstTD              int     `json:"dsttd,omitempty"`
	Name               string  `json:"name,omitempty"`
	NatIP              string  `json:"natip,omitempty"`
	NatPort            int     `json:"natport,omitempty"`
	NatType            string  `json:"nattype,omitempty"`
	Network6           string  `json:"network6,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Status             string  `json:"status,omitempty"`
	SubscrIP           string  `json:"subscrip,omitempty"`
	SubscrPort         int     `json:"subscrport,omitempty"`
	TD                 int     `json:"td,omitempty"`
	TransportProtocol  string  `json:"transportprotocol,omitempty"`
}

type LSNTransportProfile struct {
	Count                float64 `json:"__count,omitempty"`
	FinRstTimeout        int     `json:"finrsttimeout,omitempty"`
	GroupSessionLimit    int     `json:"groupsessionlimit,omitempty"`
	NextGenAPIResource   string  `json:"_nextgenapiresource,omitempty"`
	PortPreserveParity   string  `json:"portpreserveparity,omitempty"`
	PortPreserveRange    string  `json:"portpreserverange,omitempty"`
	PortQuota            int     `json:"portquota,omitempty"`
	SessionQuota         int     `json:"sessionquota,omitempty"`
	SessionTimeout       int     `json:"sessiontimeout,omitempty"`
	STUNTimeout          int     `json:"stuntimeout,omitempty"`
	SynCheck             string  `json:"syncheck,omitempty"`
	SynIdleTimeout       int     `json:"synidletimeout,omitempty"`
	TransportProfileName string  `json:"transportprofilename,omitempty"`
	TransportProtocol    string  `json:"transportprotocol,omitempty"`
}

type LSNSIPALGCallDataChannelBinding struct {
	CallID          string `json:"callid,omitempty"`
	ChannelFlags    int    `json:"channelflags,omitempty"`
	ChannelIP       string `json:"channelip,omitempty"`
	ChannelNatIP    string `json:"channelnatip,omitempty"`
	ChannelNatPort  int    `json:"channelnatport,omitempty"`
	ChannelPort     int    `json:"channelport,omitempty"`
	ChannelProtocol string `json:"channelprotocol,omitempty"`
	ChannelTimeout  int    `json:"channeltimeout,omitempty"`
}

type LSNAppsProfileBinding struct {
	AppsProfileName                        string        `json:"appsprofilename,omitempty"`
	LSNAppsProfileLSNAppsAttributesBinding []interface{} `json:"lsnappsprofile_lsnappsattributes_binding,omitempty"`
	LSNAppsProfilePortBinding              []interface{} `json:"lsnappsprofile_port_binding,omitempty"`
}

type LSNPool struct {
	Count               float64 `json:"__count,omitempty"`
	MaxPortReallocTMQ   int     `json:"maxportrealloctmq,omitempty"`
	NatType             string  `json:"nattype,omitempty"`
	NextGenAPIResource  string  `json:"_nextgenapiresource,omitempty"`
	PoolName            string  `json:"poolname,omitempty"`
	PortBlockAllocation string  `json:"portblockallocation,omitempty"`
	PortReallocTimeout  int     `json:"portrealloctimeout,omitempty"`
}

type LSNAppsProfileLSNAppsAttributesBinding struct {
	AppsAttributesName string `json:"appsattributesname,omitempty"`
	AppsProfileName    string `json:"appsprofilename,omitempty"`
}

type LSNGroupBinding struct {
	GroupName                           string        `json:"groupname,omitempty"`
	LSNGroupIPSECALGProfileBinding      []interface{} `json:"lsngroup_ipsecalgprofile_binding,omitempty"`
	LSNGroupLSNAppsProfileBinding       []interface{} `json:"lsngroup_lsnappsprofile_binding,omitempty"`
	LSNGroupLSNHTTPHdrLogProfileBinding []interface{} `json:"lsngroup_lsnhttphdrlogprofile_binding,omitempty"`
	LSNGroupLSNLogProfileBinding        []interface{} `json:"lsngroup_lsnlogprofile_binding,omitempty"`
	LSNGroupLSNPoolBinding              []interface{} `json:"lsngroup_lsnpool_binding,omitempty"`
	LSNGroupLSNRTSPALGProfileBinding    []interface{} `json:"lsngroup_lsnrtspalgprofile_binding,omitempty"`
	LSNGroupLSNSIPALGProfileBinding     []interface{} `json:"lsngroup_lsnsipalgprofile_binding,omitempty"`
	LSNGroupLSNTransportProfileBinding  []interface{} `json:"lsngroup_lsntransportprofile_binding,omitempty"`
	LSNGroupPCPServerBinding            []interface{} `json:"lsngroup_pcpserver_binding,omitempty"`
}

type LSNPoolBinding struct {
	LSNPoolLSNIPBinding []interface{} `json:"lsnpool_lsnip_binding,omitempty"`
	PoolName            string        `json:"poolname,omitempty"`
}

type LSNDeterministicNAT struct {
	ClientName         string  `json:"clientname,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	FirstPort          int     `json:"firstport,omitempty"`
	LastPort           int     `json:"lastport,omitempty"`
	NatIP              string  `json:"natip,omitempty"`
	NatIP2             string  `json:"natip2,omitempty"`
	NatPrefix          string  `json:"natprefix,omitempty"`
	NatType            string  `json:"nattype,omitempty"`
	Network6           string  `json:"network6,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	SrcTD              int     `json:"srctd,omitempty"`
	SubscrIP           string  `json:"subscrip,omitempty"`
	SubscrIP2          string  `json:"subscrip2,omitempty"`
	TD                 int     `json:"td,omitempty"`
}

type LSNGroupLSNAppsProfileBinding struct {
	AppsProfileName string `json:"appsprofilename,omitempty"`
	GroupName       string `json:"groupname,omitempty"`
}

type LSNSIPALGCallBinding struct {
	CallID                             string        `json:"callid,omitempty"`
	LSNSIPALGCallControlChannelBinding []interface{} `json:"lsnsipalgcall_controlchannel_binding,omitempty"`
	LSNSIPALGCallDataChannelBinding    []interface{} `json:"lsnsipalgcall_datachannel_binding,omitempty"`
}

type LSNAppsProfile struct {
	AppsProfileName    string  `json:"appsprofilename,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Filtering          string  `json:"filtering,omitempty"`
	IPPooling          string  `json:"ippooling,omitempty"`
	L2Info             string  `json:"l2info,omitempty"`
	Mapping            string  `json:"mapping,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	TCPProxy           string  `json:"tcpproxy,omitempty"`
	TD                 int     `json:"td,omitempty"`
	TransportProtocol  string  `json:"transportprotocol,omitempty"`
}

type LSNHTTPHdrLogProfile struct {
	Count                 float64 `json:"__count,omitempty"`
	HTTPHdrLogProfileName string  `json:"httphdrlogprofilename,omitempty"`
	LogHost               string  `json:"loghost,omitempty"`
	LogMethod             string  `json:"logmethod,omitempty"`
	LogURL                string  `json:"logurl,omitempty"`
	LogVersion            string  `json:"logversion,omitempty"`
	NextGenAPIResource    string  `json:"_nextgenapiresource,omitempty"`
}

type LSNSIPALGCallControlChannelBinding struct {
	CallID          string `json:"callid,omitempty"`
	ChannelFlags    int    `json:"channelflags,omitempty"`
	ChannelIP       string `json:"channelip,omitempty"`
	ChannelNatIP    string `json:"channelnatip,omitempty"`
	ChannelNatPort  int    `json:"channelnatport,omitempty"`
	ChannelPort     int    `json:"channelport,omitempty"`
	ChannelProtocol string `json:"channelprotocol,omitempty"`
	ChannelTimeout  int    `json:"channeltimeout,omitempty"`
}

type LSNGroupLSNPoolBinding struct {
	GroupName string `json:"groupname,omitempty"`
	PoolName  string `json:"poolname,omitempty"`
}

type LSNSIPALGCall struct {
	CallFlags          int     `json:"callflags,omitempty"`
	CallID             string  `json:"callid,omitempty"`
	CallRefCount       int     `json:"callrefcount,omitempty"`
	CallTimer          int     `json:"calltimer,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NodeID             int     `json:"nodeid,omitempty"`
	XlatIP             string  `json:"xlatip,omitempty"`
}

type LSNPoolLSNIPBinding struct {
	LSNIP     string `json:"lsnip,omitempty"`
	OwnerNode int    `json:"ownernode,omitempty"`
	PoolName  string `json:"poolname,omitempty"`
}

type LSNRTSPALGSession struct {
	CallFlags          int     `json:"callflags,omitempty"`
	CallRefCount       int     `json:"callrefcount,omitempty"`
	CallTimer          int     `json:"calltimer,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NodeID             int     `json:"nodeid,omitempty"`
	SessionID          string  `json:"sessionid,omitempty"`
	XlatIP             string  `json:"xlatip,omitempty"`
}

type LSNRTSPALGProfile struct {
	Count                 float64 `json:"__count,omitempty"`
	NextGenAPIResource    string  `json:"_nextgenapiresource,omitempty"`
	RTSPALGProfileName    string  `json:"rtspalgprofilename,omitempty"`
	RTSPIdleTimeout       int     `json:"rtspidletimeout,omitempty"`
	RTSPPortRange         string  `json:"rtspportrange,omitempty"`
	RTSPTransportProtocol string  `json:"rtsptransportprotocol,omitempty"`
}

type LSNGroupPCPServerBinding struct {
	GroupName string `json:"groupname,omitempty"`
	PCPServer string `json:"pcpserver,omitempty"`
}

type LSNClientBinding struct {
	ClientName               string        `json:"clientname,omitempty"`
	LSNClientNetwork6Binding []interface{} `json:"lsnclient_network6_binding,omitempty"`
	LSNClientNetworkBinding  []interface{} `json:"lsnclient_network_binding,omitempty"`
	LSNClientNSACL6Binding   []interface{} `json:"lsnclient_nsacl6_binding,omitempty"`
	LSNClientNSACLBinding    []interface{} `json:"lsnclient_nsacl_binding,omitempty"`
}

type LSNRTSPALGSessionDataChannelBinding struct {
	ChannelFlags    int    `json:"channelflags,omitempty"`
	ChannelIP       string `json:"channelip,omitempty"`
	ChannelNatIP    string `json:"channelnatip,omitempty"`
	ChannelNatPort  int    `json:"channelnatport,omitempty"`
	ChannelPort     int    `json:"channelport,omitempty"`
	ChannelProtocol string `json:"channelprotocol,omitempty"`
	ChannelTimeout  int    `json:"channeltimeout,omitempty"`
	SessionID       string `json:"sessionid,omitempty"`
}

type LSNGroupLSNSIPALGProfileBinding struct {
	GroupName         string `json:"groupname,omitempty"`
	SIPALGProfileName string `json:"sipalgprofilename,omitempty"`
}
