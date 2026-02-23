package models

// network configuration structs
type ChannelInterfaceBinding struct {
	ID           string   `json:"id,omitempty"`
	IFNum        []string `json:"ifnum,omitempty"`
	LAMode       string   `json:"lamode,omitempty"`
	LRActiveIntf int      `json:"lractiveintf,omitempty"`
	SlaveDuplex  int      `json:"slaveduplex,omitempty"`
	SlaveFlowCtl int      `json:"slaveflowctl,omitempty"`
	SlaveMedia   int      `json:"slavemedia,omitempty"`
	SlaveSpeed   int      `json:"slavespeed,omitempty"`
	SlaveState   int      `json:"slavestate,omitempty"`
	SlaveTime    int      `json:"slavetime,omitempty"`
	SVMCmd       int      `json:"svmcmd,omitempty"`
}

type MapDMR struct {
	BRIPv6Prefix       string  `json:"bripv6prefix,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
}

type RNAT6Binding struct {
	Name              string        `json:"name,omitempty"`
	RNAT6NSIP6Binding []interface{} `json:"rnat6_nsip6_binding,omitempty"`
}

type VLANNSIP6Binding struct {
	ID         int    `json:"id,omitempty"`
	IPAddress  string `json:"ipaddress,omitempty"`
	Netmask    string `json:"netmask,omitempty"`
	OwnerGroup string `json:"ownergroup,omitempty"`
	TD         int    `json:"td,omitempty"`
}

type ND6RAVariablesOnLinkIPv6PrefixBinding struct {
	IPv6Prefix string `json:"ipv6prefix,omitempty"`
	VLAN       int    `json:"vlan,omitempty"`
}

type PTP struct {
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	State              string `json:"state,omitempty"`
}

type RNATNSIPBinding struct {
	Name       string `json:"name,omitempty"`
	NatIP      string `json:"natip,omitempty"`
	OwnerGroup string `json:"ownergroup,omitempty"`
	TD         int    `json:"td,omitempty"`
}

type VRID6 struct {
	All                  bool    `json:"all,omitempty"`
	Count                float64 `json:"__count,omitempty"`
	EffectivePriority    int     `json:"effectivepriority,omitempty"`
	Flags                int     `json:"flags,omitempty"`
	ID                   int     `json:"id,omitempty"`
	Ifaces               string  `json:"ifaces,omitempty"`
	IFNum                string  `json:"ifnum,omitempty"`
	IPAddress            string  `json:"ipaddress,omitempty"`
	NextGenAPIResource   string  `json:"_nextgenapiresource,omitempty"`
	OperationalOwnerNode int     `json:"operationalownernode,omitempty"`
	OwnerNode            int     `json:"ownernode,omitempty"`
	Preemption           string  `json:"preemption,omitempty"`
	PreemptionDelayTimer int     `json:"preemptiondelaytimer,omitempty"`
	Priority             int     `json:"priority,omitempty"`
	Sharing              string  `json:"sharing,omitempty"`
	State                int     `json:"state,omitempty"`
	TrackIFNumPriority   int     `json:"trackifnumpriority,omitempty"`
	Tracking             string  `json:"tracking,omitempty"`
	TypeField            string  `json:"type,omitempty"`
}

type NetProfileNATRuleBinding struct {
	Name      string `json:"name,omitempty"`
	NATRule   string `json:"natrule,omitempty"`
	Netmask   string `json:"netmask,omitempty"`
	RewriteIP string `json:"rewriteip,omitempty"`
}

type IPTunnel struct {
	Channel            int      `json:"channel,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	DestPort           int      `json:"destport,omitempty"`
	EncapIP            string   `json:"encapip,omitempty"`
	GREPayload         string   `json:"grepayload,omitempty"`
	IPSECProfileName   string   `json:"ipsecprofilename,omitempty"`
	IPSECTunnelStatus  string   `json:"ipsectunnelstatus,omitempty"`
	Local              string   `json:"local,omitempty"`
	Name               string   `json:"name,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	OwnerGroup         string   `json:"ownergroup,omitempty"`
	Protocol           string   `json:"protocol,omitempty"`
	RefCnt             int      `json:"refcnt,omitempty"`
	Remote             string   `json:"remote,omitempty"`
	RemoteSubnetMask   string   `json:"remotesubnetmask,omitempty"`
	SysName            string   `json:"sysname,omitempty"`
	TOSInherit         string   `json:"tosinherit,omitempty"`
	TunnelType         []string `json:"tunneltype,omitempty"`
	TypeField          int      `json:"type,omitempty"`
	VLAN               int      `json:"vlan,omitempty"`
	VLANTagging        string   `json:"vlantagging,omitempty"`
	VNID               int      `json:"vnid,omitempty"`
}

type NetProfileSrcPortSetBinding struct {
	Name         string `json:"name,omitempty"`
	SrcPortRange string `json:"srcportrange,omitempty"`
}

type L2Param struct {
	BdgGrpProxyARP          string `json:"bdggrpproxyarp,omitempty"`
	BdgSetting              string `json:"bdgsetting,omitempty"`
	BridgeAgeTimeout        int    `json:"bridgeagetimeout,omitempty"`
	GARPOnVRIDIntf          string `json:"garponvridintf,omitempty"`
	GARPReply               string `json:"garpreply,omitempty"`
	MacModeFwdMyPkt         string `json:"macmodefwdmypkt,omitempty"`
	MaxBridgeCollision      int    `json:"maxbridgecollision,omitempty"`
	MBFInstLearning         string `json:"mbfinstlearning,omitempty"`
	MBFPeerMACUpdate        int    `json:"mbfpeermacupdate,omitempty"`
	NextGenAPIResource      string `json:"_nextgenapiresource,omitempty"`
	ProxyARP                string `json:"proxyarp,omitempty"`
	ReturnToEthernetSender  string `json:"returntoethernetsender,omitempty"`
	RstIntfOnHAFO           string `json:"rstintfonhafo,omitempty"`
	SkipProxyingBSDTraffic  string `json:"skipproxyingbsdtraffic,omitempty"`
	StopMACMoveUpdate       string `json:"stopmacmoveupdate,omitempty"`
	UseMyMAC                string `json:"usemymac,omitempty"`
	UseNetProfileBSDTraffic string `json:"usenetprofilebsdtraffic,omitempty"`
}

type ND6RAVariablesBinding struct {
	ND6RAVariablesOnLinkIPv6PrefixBinding []interface{} `json:"nd6ravariables_onlinkipv6prefix_binding,omitempty"`
	VLAN                                  int           `json:"vlan,omitempty"`
}

type BridgeGroup struct {
	Count              float64 `json:"__count,omitempty"`
	DynamicRouting     string  `json:"dynamicrouting,omitempty"`
	Flags              bool    `json:"flags,omitempty"`
	ID                 int     `json:"id,omitempty"`
	Ifaces             string  `json:"ifaces,omitempty"`
	IPv6DynamicRouting string  `json:"ipv6dynamicrouting,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	PartitionName      string  `json:"partitionname,omitempty"`
	PortBitmap         int     `json:"portbitmap,omitempty"`
	RNAT               bool    `json:"rnat,omitempty"`
	TagBitmap          int     `json:"tagbitmap,omitempty"`
	TagIfaces          string  `json:"tagifaces,omitempty"`
}

type IPSetNSIP6Binding struct {
	IPAddress string `json:"ipaddress,omitempty"`
	Name      string `json:"name,omitempty"`
}

type VRID6ChannelBinding struct {
	Flags int    `json:"flags,omitempty"`
	ID    int    `json:"id,omitempty"`
	IFNum string `json:"ifnum,omitempty"`
	VLAN  int    `json:"vlan,omitempty"`
}

type ChannelBinding struct {
	ChannelInterfaceBinding []interface{} `json:"channel_interface_binding,omitempty"`
	ID                      string        `json:"id,omitempty"`
}

type RNAT6 struct {
	ACL6Name           string  `json:"acl6name,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Network            string  `json:"network,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	OwnerGroup         string  `json:"ownergroup,omitempty"`
	RedirectPort       int     `json:"redirectport,omitempty"`
	SrcIPPersistency   string  `json:"srcippersistency,omitempty"`
	TD                 int     `json:"td,omitempty"`
}

type BridgeGroupNSIP6Binding struct {
	ID         int    `json:"id,omitempty"`
	IPAddress  string `json:"ipaddress,omitempty"`
	Netmask    string `json:"netmask,omitempty"`
	OwnerGroup string `json:"ownergroup,omitempty"`
	RNAT       bool   `json:"rnat,omitempty"`
	TD         int    `json:"td,omitempty"`
}

type BridgeGroupVLANBinding struct {
	ID   int  `json:"id,omitempty"`
	RNAT bool `json:"rnat,omitempty"`
	VLAN int  `json:"vlan,omitempty"`
}

type ND6RAVariables struct {
	CeaseRouterAdv           string  `json:"ceaserouteradv,omitempty"`
	Count                    float64 `json:"__count,omitempty"`
	CurrHopLimit             int     `json:"currhoplimit,omitempty"`
	DefaultLifetime          int     `json:"defaultlifetime,omitempty"`
	LastRtAdvTime            int     `json:"lastrtadvtime,omitempty"`
	LinkMTU                  int     `json:"linkmtu,omitempty"`
	ManagedAddrConfig        string  `json:"managedaddrconfig,omitempty"`
	MaxRtAdvInterval         int     `json:"maxrtadvinterval,omitempty"`
	MinRtAdvInterval         int     `json:"minrtadvinterval,omitempty"`
	NextGenAPIResource       string  `json:"_nextgenapiresource,omitempty"`
	NextRtAdvDelay           int     `json:"nextrtadvdelay,omitempty"`
	OnlyUnicastRtAdvResponse string  `json:"onlyunicastrtadvresponse,omitempty"`
	OtherAddrConfig          string  `json:"otheraddrconfig,omitempty"`
	ReachableTime            int     `json:"reachabletime,omitempty"`
	RetransTime              int     `json:"retranstime,omitempty"`
	SendRouterAdv            string  `json:"sendrouteradv,omitempty"`
	SrcLinkLayerAddrOption   string  `json:"srclinklayeraddroption,omitempty"`
	VLAN                     int     `json:"vlan,omitempty"`
}

type VRIDNSIPBinding struct {
	Flags     int    `json:"flags,omitempty"`
	ID        int    `json:"id,omitempty"`
	IPAddress string `json:"ipaddress,omitempty"`
}

type VLANInterfaceBinding struct {
	ID         int    `json:"id,omitempty"`
	IFNum      string `json:"ifnum,omitempty"`
	OwnerGroup string `json:"ownergroup,omitempty"`
	Tagged     bool   `json:"tagged,omitempty"`
}

type IP6TunnelParam struct {
	DropFrag             string `json:"dropfrag,omitempty"`
	DropFragCPUThreshold int    `json:"dropfragcputhreshold,omitempty"`
	NextGenAPIResource   string `json:"_nextgenapiresource,omitempty"`
	SrcIP                string `json:"srcip,omitempty"`
	SrcIPRoundRobin      string `json:"srciproundrobin,omitempty"`
	UseClientSourceIP    string `json:"useclientsourceipv6,omitempty"`
}

type VRID6TrackInterfaceBinding struct {
	Flags      int    `json:"flags,omitempty"`
	ID         int    `json:"id,omitempty"`
	TrackIFNum string `json:"trackifnum,omitempty"`
}

type RNATGlobalAuditSyslogPolicyBinding struct {
	All      bool   `json:"all,omitempty"`
	Policy   string `json:"policy,omitempty"`
	Priority int    `json:"priority,omitempty"`
}

type ARPParam struct {
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	SpoofValidation    string `json:"spoofvalidation,omitempty"`
	Timeout            int    `json:"timeout,omitempty"`
}

type Channel struct {
	ActFlowCtl                string   `json:"actflowctl,omitempty"`
	ActSpeed                  string   `json:"actspeed,omitempty"`
	ActThroughput             int      `json:"actthroughput,omitempty"`
	ActualMTU                 int      `json:"actualmtu,omitempty"`
	Autoneg                   int      `json:"autoneg,omitempty"`
	AutonegResult             int      `json:"autonegresult,omitempty"`
	Backplane                 string   `json:"backplane,omitempty"`
	BandwidthHigh             int      `json:"bandwidthhigh,omitempty"`
	BandwidthNormal           int      `json:"bandwidthnormal,omitempty"`
	BdgMuted                  int      `json:"bdgmuted,omitempty"`
	ClearTime                 int      `json:"cleartime,omitempty"`
	ConnDistr                 string   `json:"conndistr,omitempty"`
	Count                     float64  `json:"__count,omitempty"`
	Description               string   `json:"description,omitempty"`
	DeviceName                string   `json:"devicename,omitempty"`
	Downtime                  int      `json:"downtime,omitempty"`
	Duplex                    string   `json:"duplex,omitempty"`
	FCtls                     int      `json:"fctls,omitempty"`
	Flags                     int      `json:"flags,omitempty"`
	FlowCtl                   string   `json:"flowctl,omitempty"`
	HAHeartbeat               string   `json:"haheartbeat,omitempty"`
	HAMonitor                 string   `json:"hamonitor,omitempty"`
	HangDetect                int      `json:"hangdetect,omitempty"`
	HangReset                 int      `json:"hangreset,omitempty"`
	Hangs                     int      `json:"hangs,omitempty"`
	ID                        string   `json:"id,omitempty"`
	IFAlias                   string   `json:"ifalias,omitempty"`
	IFNum                     []string `json:"ifnum,omitempty"`
	InDisc                    int      `json:"indisc,omitempty"`
	IntfState                 int      `json:"intfstate,omitempty"`
	LACPActorAggregation      string   `json:"lacpactoraggregation,omitempty"`
	LACPActorCollecting       string   `json:"lacpactorcollecting,omitempty"`
	LACPActorDistributing     string   `json:"lacpactordistributing,omitempty"`
	LACPActorInSync           string   `json:"lacpactorinsync,omitempty"`
	LACPActorPortNo           int      `json:"lacpactorportno,omitempty"`
	LACPActorPriority         int      `json:"lacpactorpriority,omitempty"`
	LACPMode                  string   `json:"lacpmode,omitempty"`
	LACPPartnerAggregation    string   `json:"lacppartneraggregation,omitempty"`
	LACPPartnerCollecting     string   `json:"lacppartnercollecting,omitempty"`
	LACPPartnerDefaulted      string   `json:"lacppartnerdefaulted,omitempty"`
	LACPPartnerDistributing   string   `json:"lacppartnerdistributing,omitempty"`
	LACPPartnerExpired        string   `json:"lacppartnerexpired,omitempty"`
	LACPPartnerInSync         string   `json:"lacppartnerinsync,omitempty"`
	LACPPartnerKey            int      `json:"lacppartnerkey,omitempty"`
	LACPPartnerPortNo         int      `json:"lacppartnerportno,omitempty"`
	LACPPartnerPriority       int      `json:"lacppartnerpriority,omitempty"`
	LACPPartnerState          string   `json:"lacppartnerstate,omitempty"`
	LACPPartnerSystemMAC      string   `json:"lacppartnersystemmac,omitempty"`
	LACPPartnerSystemPriority int      `json:"lacppartnersystempriority,omitempty"`
	LACPPartnerTimeout        string   `json:"lacppartnertimeout,omitempty"`
	LACPPortMuxState          string   `json:"lacpportmuxstate,omitempty"`
	LACPPortRxStat            string   `json:"lacpportrxstat,omitempty"`
	LACPPortSelectState       string   `json:"lacpportselectstate,omitempty"`
	LACPTimeout               string   `json:"lacptimeout,omitempty"`
	LAMAC                     string   `json:"lamac,omitempty"`
	LAMode                    string   `json:"lamode,omitempty"`
	LinkRedundancy            string   `json:"linkredundancy,omitempty"`
	LinkState                 int      `json:"linkstate,omitempty"`
	LLDPMode                  string   `json:"lldpmode,omitempty"`
	LRMinThroughput           int      `json:"lrminthroughput,omitempty"`
	MAC                       string   `json:"mac,omitempty"`
	MACDistr                  string   `json:"macdistr,omitempty"`
	Media                     string   `json:"media,omitempty"`
	Mode                      string   `json:"mode,omitempty"`
	MTU                       int      `json:"mtu,omitempty"`
	NextGenAPIResource        string   `json:"_nextgenapiresource,omitempty"`
	OutDisc                   int      `json:"outdisc,omitempty"`
	ReqDuplex                 string   `json:"reqduplex,omitempty"`
	ReqFlowControl            string   `json:"reqflowcontrol,omitempty"`
	ReqMedia                  string   `json:"reqmedia,omitempty"`
	ReqSpeed                  string   `json:"reqspeed,omitempty"`
	ReqThroughput             int      `json:"reqthroughput,omitempty"`
	RxBytes                   int      `json:"rxbytes,omitempty"`
	RxDrops                   int      `json:"rxdrops,omitempty"`
	RxErrors                  int      `json:"rxerrors,omitempty"`
	RxPackets                 int      `json:"rxpackets,omitempty"`
	RxStalls                  int      `json:"rxstalls,omitempty"`
	Speed                     string   `json:"speed,omitempty"`
	State                     string   `json:"state,omitempty"`
	StsStalls                 int      `json:"stsstalls,omitempty"`
	TagAll                    string   `json:"tagall,omitempty"`
	Tagged                    int      `json:"tagged,omitempty"`
	TaggedAny                 int      `json:"taggedany,omitempty"`
	TaggedAutoLearn           int      `json:"taggedautolearn,omitempty"`
	Throughput                int      `json:"throughput,omitempty"`
	Trunk                     string   `json:"trunk,omitempty"`
	TxBytes                   int      `json:"txbytes,omitempty"`
	TxDrops                   int      `json:"txdrops,omitempty"`
	TxErrors                  int      `json:"txerrors,omitempty"`
	TxPackets                 int      `json:"txpackets,omitempty"`
	TxStalls                  int      `json:"txstalls,omitempty"`
	Unit                      int      `json:"unit,omitempty"`
	Uptime                    int      `json:"uptime,omitempty"`
	VLAN                      int      `json:"vlan,omitempty"`
	VMAC                      string   `json:"vmac,omitempty"`
	VMAC6                     string   `json:"vmac6,omitempty"`
}

type NetProfile struct {
	BadIPActionThreshold           int      `json:"badipactionthreshold,omitempty"`
	Count                          float64  `json:"__count,omitempty"`
	MBF                            string   `json:"mbf,omitempty"`
	Name                           string   `json:"name,omitempty"`
	NextGenAPIResource             string   `json:"_nextgenapiresource,omitempty"`
	OverrideLSN                    string   `json:"overridelsn,omitempty"`
	ProxyProtocol                  string   `json:"proxyprotocol,omitempty"`
	ProxyProtocolAfterTLSHandshake string   `json:"proxyprotocolaftertlshandshake,omitempty"`
	ProxyProtocolTLVOptions        []string `json:"proxyprotocoltlvoptions,omitempty"`
	ProxyProtocolTxVersion         string   `json:"proxyprotocoltxversion,omitempty"`
	SrcIP                          string   `json:"srcip,omitempty"`
	SrcIPPersistency               string   `json:"srcippersistency,omitempty"`
	TD                             int      `json:"td,omitempty"`
}

type VXLANNSIPBinding struct {
	ID        int    `json:"id,omitempty"`
	IPAddress string `json:"ipaddress,omitempty"`
	Netmask   string `json:"netmask,omitempty"`
}

type RNAT struct {
	ACLName            string  `json:"aclname,omitempty"`
	ConnFailover       string  `json:"connfailover,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NatIP              string  `json:"natip,omitempty"`
	Netmask            string  `json:"netmask,omitempty"`
	Network            string  `json:"network,omitempty"`
	NewName            string  `json:"newname,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	OwnerGroup         string  `json:"ownergroup,omitempty"`
	RedirectPort       int     `json:"redirectport,omitempty"`
	SrcIPPersistency   string  `json:"srcippersistency,omitempty"`
	TD                 int     `json:"td,omitempty"`
	UseProxyPort       string  `json:"useproxyport,omitempty"`
}

type AppALGParam struct {
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	PPTPGREIdleTimeout int    `json:"pptpgreidletimeout,omitempty"`
}

type IPTunnelParam struct {
	DropFrag             string `json:"dropfrag,omitempty"`
	DropFragCPUThreshold int    `json:"dropfragcputhreshold,omitempty"`
	EnableStrictRx       string `json:"enablestrictrx,omitempty"`
	EnableStrictTx       string `json:"enablestricttx,omitempty"`
	MAC                  string `json:"mac,omitempty"`
	NextGenAPIResource   string `json:"_nextgenapiresource,omitempty"`
	SrcIP                string `json:"srcip,omitempty"`
	SrcIPRoundRobin      string `json:"srciproundrobin,omitempty"`
	UseClientSourceIP    string `json:"useclientsourceip,omitempty"`
}

type VRIDNSIP6Binding struct {
	Flags     int    `json:"flags,omitempty"`
	ID        int    `json:"id,omitempty"`
	IPAddress string `json:"ipaddress,omitempty"`
}

type BridgeGroupBinding struct {
	BridgeGroupNSIP6Binding []interface{} `json:"bridgegroup_nsip6_binding,omitempty"`
	BridgeGroupNSIPBinding  []interface{} `json:"bridgegroup_nsip_binding,omitempty"`
	BridgeGroupVLANBinding  []interface{} `json:"bridgegroup_vlan_binding,omitempty"`
	ID                      int           `json:"id,omitempty"`
}

type IPSet struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	TD                 int     `json:"td,omitempty"`
}

type ARP struct {
	All                bool    `json:"all,omitempty"`
	Channel            int     `json:"channel,omitempty"`
	ControlPlane       bool    `json:"controlplane,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Flags              int     `json:"flags,omitempty"`
	IFNum              string  `json:"ifnum,omitempty"`
	IPAddress          string  `json:"ipaddress,omitempty"`
	MAC                string  `json:"mac,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NodeID             int     `json:"nodeid,omitempty"`
	OwnerNode          int     `json:"ownernode,omitempty"`
	State              int     `json:"state,omitempty"`
	TD                 int     `json:"td,omitempty"`
	Timeout            int     `json:"timeout,omitempty"`
	TypeField          string  `json:"type,omitempty"`
	VLAN               int     `json:"vlan,omitempty"`
	VTEP               string  `json:"vtep,omitempty"`
	VXLAN              int     `json:"vxlan,omitempty"`
}

type INATParam struct {
	Count              float64 `json:"__count,omitempty"`
	Nat46FragHeader    string  `json:"nat46fragheader,omitempty"`
	Nat46IgnoreTOS     string  `json:"nat46ignoretos,omitempty"`
	Nat46V6MTU         int     `json:"nat46v6mtu,omitempty"`
	Nat46V6Prefix      string  `json:"nat46v6prefix,omitempty"`
	Nat46ZeroChecksum  string  `json:"nat46zerochecksum,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	TD                 int     `json:"td,omitempty"`
}

type L3Param struct {
	ACLLogTime           int    `json:"acllogtime,omitempty"`
	AllowClassEIPv4      string `json:"allowclasseipv4,omitempty"`
	DropDFFlag           string `json:"dropdfflag,omitempty"`
	DropIPFragments      string `json:"dropipfragments,omitempty"`
	DynamicRouting       string `json:"dynamicrouting,omitempty"`
	ExternalLoopback     string `json:"externalloopback,omitempty"`
	ForwardICMPFragments string `json:"forwardicmpfragments,omitempty"`
	ICMPGenRateThreshold int    `json:"icmpgenratethreshold,omitempty"`
	ImplicitACLAllow     string `json:"implicitaclallow,omitempty"`
	ImplicitPBR          string `json:"implicitpbr,omitempty"`
	IPv6DynamicRouting   string `json:"ipv6dynamicrouting,omitempty"`
	MIPRoundRobin        string `json:"miproundrobin,omitempty"`
	NextGenAPIResource   string `json:"_nextgenapiresource,omitempty"`
	OverrideRNAT         string `json:"overridernat,omitempty"`
	SrcNat               string `json:"srcnat,omitempty"`
	TNLPmtuWOConn        string `json:"tnlpmtuwoconn,omitempty"`
	USIPServerStrayPkt   string `json:"usipserverstraypkt,omitempty"`
}

type NAT64Param struct {
	Count              float64 `json:"__count,omitempty"`
	Nat64FragHeader    string  `json:"nat64fragheader,omitempty"`
	Nat64IgnoreTOS     string  `json:"nat64ignoretos,omitempty"`
	Nat64V6MTU         int     `json:"nat64v6mtu,omitempty"`
	Nat64ZeroChecksum  string  `json:"nat64zerochecksum,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	TD                 int     `json:"td,omitempty"`
}

type LACP struct {
	ClusterMAC         string  `json:"clustermac,omitempty"`
	ClusterSysPriority int     `json:"clustersyspriority,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	DeviceName         string  `json:"devicename,omitempty"`
	Flags              int     `json:"flags,omitempty"`
	LACPKey            int     `json:"lacpkey,omitempty"`
	MAC                string  `json:"mac,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	OwnerNode          int     `json:"ownernode,omitempty"`
	SysPriority        int     `json:"syspriority,omitempty"`
}

type MapDomain struct {
	Count              float64 `json:"__count,omitempty"`
	MapDMRName         string  `json:"mapdmrname,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
}

type ForwardingSession struct {
	ACL6Name           string  `json:"acl6name,omitempty"`
	ACLName            string  `json:"aclname,omitempty"`
	ConnFailover       string  `json:"connfailover,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Netmask            string  `json:"netmask,omitempty"`
	Network            string  `json:"network,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	ProcessLocal       string  `json:"processlocal,omitempty"`
	SourceRouteCache   string  `json:"sourceroutecache,omitempty"`
	TD                 int     `json:"td,omitempty"`
}

type Route6 struct {
	Active             bool     `json:"active,omitempty"`
	Advertise          string   `json:"advertise,omitempty"`
	BGP                bool     `json:"bgp,omitempty"`
	Connected          bool     `json:"connected,omitempty"`
	Cost               int      `json:"cost,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Data               bool     `json:"data,omitempty"`
	Data1              string   `json:"data1,omitempty"`
	Detail             bool     `json:"detail,omitempty"`
	Distance           int      `json:"distance,omitempty"`
	Dynamic            bool     `json:"dynamic,omitempty"`
	FailedProbes       int      `json:"failedprobes,omitempty"`
	Flags              bool     `json:"flags,omitempty"`
	Gateway            string   `json:"gateway,omitempty"`
	GatewayName        string   `json:"gatewayname,omitempty"`
	ISIS               bool     `json:"isis,omitempty"`
	Mgmt               bool     `json:"mgmt,omitempty"`
	Monitor            string   `json:"monitor,omitempty"`
	MonStatCode        int      `json:"monstatcode,omitempty"`
	MonStatParam1      int      `json:"monstatparam1,omitempty"`
	MonStatParam2      int      `json:"monstatparam2,omitempty"`
	MonStatParam3      int      `json:"monstatparam3,omitempty"`
	MSR                string   `json:"msr,omitempty"`
	Network            string   `json:"network,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	OSPFv3             bool     `json:"ospfv3,omitempty"`
	OwnerGroup         string   `json:"ownergroup,omitempty"`
	Permanent          bool     `json:"permanent,omitempty"`
	RARoute            bool     `json:"raroute,omitempty"`
	Retain             int      `json:"retain,omitempty"`
	RIP                bool     `json:"rip,omitempty"`
	RouteOwners        []string `json:"routeowners,omitempty"`
	RouteType          string   `json:"routetype,omitempty"`
	State              int      `json:"state,omitempty"`
	Static             bool     `json:"Static,omitempty"`
	TD                 int      `json:"td,omitempty"`
	TotalFailedProbes  int      `json:"totalfailedprobes,omitempty"`
	TotalProbes        int      `json:"totalprobes,omitempty"`
	TypeField          bool     `json:"type,omitempty"`
	VLAN               int      `json:"vlan,omitempty"`
	VXLAN              int      `json:"vxlan,omitempty"`
	Weight             int      `json:"weight,omitempty"`
}

type VLANLinkSetBinding struct {
	ID         int    `json:"id,omitempty"`
	IFNum      string `json:"ifnum,omitempty"`
	OwnerGroup string `json:"ownergroup,omitempty"`
	Tagged     bool   `json:"tagged,omitempty"`
}

type VLANChannelBinding struct {
	ID         int    `json:"id,omitempty"`
	IFNum      string `json:"ifnum,omitempty"`
	OwnerGroup string `json:"ownergroup,omitempty"`
	Tagged     bool   `json:"tagged,omitempty"`
}

type FISBinding struct {
	FISChannelBinding []interface{} `json:"fis_channel_binding,omitempty"`
	Name              string        `json:"name,omitempty"`
}

type VXLANIPTunnelBinding struct {
	ID     int    `json:"id,omitempty"`
	Tunnel string `json:"tunnel,omitempty"`
}

type NetProfileBinding struct {
	Name                        string        `json:"name,omitempty"`
	NetProfileNATRuleBinding    []interface{} `json:"netprofile_natrule_binding,omitempty"`
	NetProfileSrcPortSetBinding []interface{} `json:"netprofile_srcportset_binding,omitempty"`
}

type VXLANVLANMapVXLANBinding struct {
	Name  string   `json:"name,omitempty"`
	VLAN  []string `json:"vlan,omitempty"`
	VXLAN int      `json:"vxlan,omitempty"`
}

type NetBridge struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	VXLANVLANMap       string  `json:"vxlanvlanmap,omitempty"`
}

type VRIDParam struct {
	DeadInterval       int    `json:"deadinterval,omitempty"`
	HelloInterval      int    `json:"hellointerval,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	SendToMaster       string `json:"sendtomaster,omitempty"`
}

type VXLANSrcIPBinding struct {
	ID    int    `json:"id,omitempty"`
	SrcIP string `json:"srcip,omitempty"`
}

type IPSetBinding struct {
	IPSetNSIP6Binding []interface{} `json:"ipset_nsip6_binding,omitempty"`
	IPSetNSIPBinding  []interface{} `json:"ipset_nsip_binding,omitempty"`
	Name              string        `json:"name,omitempty"`
}

type NAT64 struct {
	ACL6Name           string  `json:"acl6name,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NetProfile         string  `json:"netprofile,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
}

type RNATSession struct {
	ACLName string `json:"aclname,omitempty"`
	NatIP   string `json:"natip,omitempty"`
	Netmask string `json:"netmask,omitempty"`
	Network string `json:"network,omitempty"`
}

type VXLAN struct {
	Count              float64 `json:"__count,omitempty"`
	DynamicRouting     string  `json:"dynamicrouting,omitempty"`
	ID                 int     `json:"id,omitempty"`
	InnerVLANTagging   string  `json:"innervlantagging,omitempty"`
	IPv6DynamicRouting string  `json:"ipv6dynamicrouting,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	PartitionName      string  `json:"partitionname,omitempty"`
	Port               int     `json:"port,omitempty"`
	Protocol           string  `json:"protocol,omitempty"`
	TD                 int     `json:"td,omitempty"`
	TypeField          string  `json:"type,omitempty"`
	VLAN               int     `json:"vlan,omitempty"`
}

type FISInterfaceBinding struct {
	IFNum     string `json:"ifnum,omitempty"`
	Name      string `json:"name,omitempty"`
	OwnerNode int    `json:"ownernode,omitempty"`
}

type VRIDChannelBinding struct {
	Flags int    `json:"flags,omitempty"`
	ID    int    `json:"id,omitempty"`
	IFNum string `json:"ifnum,omitempty"`
	VLAN  int    `json:"vlan,omitempty"`
}

type VRIDBinding struct {
	ID                        int           `json:"id,omitempty"`
	VRIDChannelBinding        []interface{} `json:"vrid_channel_binding,omitempty"`
	VRIDInterfaceBinding      []interface{} `json:"vrid_interface_binding,omitempty"`
	VRIDNSIP6Binding          []interface{} `json:"vrid_nsip6_binding,omitempty"`
	VRIDNSIPBinding           []interface{} `json:"vrid_nsip_binding,omitempty"`
	VRIDTrackInterfaceBinding []interface{} `json:"vrid_trackinterface_binding,omitempty"`
}

type NetBridgeIPTunnelBinding struct {
	Name   string `json:"name,omitempty"`
	Tunnel string `json:"tunnel,omitempty"`
}

type IPSetNSIPBinding struct {
	IPAddress string `json:"ipaddress,omitempty"`
	Name      string `json:"name,omitempty"`
}

type VLANNSIPBinding struct {
	ID         int    `json:"id,omitempty"`
	IPAddress  string `json:"ipaddress,omitempty"`
	Netmask    string `json:"netmask,omitempty"`
	OwnerGroup string `json:"ownergroup,omitempty"`
	TD         int    `json:"td,omitempty"`
}

type LinkSetInterfaceBinding struct {
	ID    string `json:"id,omitempty"`
	IFNum string `json:"ifnum,omitempty"`
}

type VRIDTrackInterfaceBinding struct {
	Flags      int    `json:"flags,omitempty"`
	ID         int    `json:"id,omitempty"`
	TrackIFNum string `json:"trackifnum,omitempty"`
}

type VRID6Binding struct {
	ID                         int           `json:"id,omitempty"`
	VRID6ChannelBinding        []interface{} `json:"vrid6_channel_binding,omitempty"`
	VRID6InterfaceBinding      []interface{} `json:"vrid6_interface_binding,omitempty"`
	VRID6NSIP6Binding          []interface{} `json:"vrid6_nsip6_binding,omitempty"`
	VRID6NSIPBinding           []interface{} `json:"vrid6_nsip_binding,omitempty"`
	VRID6TrackInterfaceBinding []interface{} `json:"vrid6_trackinterface_binding,omitempty"`
}

type NetBridgeNSIPBinding struct {
	IPAddress string `json:"ipaddress,omitempty"`
	Name      string `json:"name,omitempty"`
	Netmask   string `json:"netmask,omitempty"`
}

type RSSKeyType struct {
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	RSSType            string `json:"rsstype,omitempty"`
}

type MapBMRBMRV4NetworkBinding struct {
	Name    string `json:"name,omitempty"`
	Netmask string `json:"netmask,omitempty"`
	Network string `json:"network,omitempty"`
}

type RNAT6NSIP6Binding struct {
	Name       string `json:"name,omitempty"`
	NatIP6     string `json:"natip6,omitempty"`
	OwnerGroup string `json:"ownergroup,omitempty"`
	TD         int    `json:"td,omitempty"`
}

type Route struct {
	Adv                bool     `json:"adv,omitempty"`
	AdvBGP             bool     `json:"advbgp,omitempty"`
	Advertise          string   `json:"advertise,omitempty"`
	AdvISIS            bool     `json:"advisis,omitempty"`
	AdvOSPF            bool     `json:"advospf,omitempty"`
	AdvRIP             bool     `json:"advrip,omitempty"`
	BGP                bool     `json:"bgp,omitempty"`
	Cost               int      `json:"cost,omitempty"`
	Cost1              int      `json:"cost1,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Data               bool     `json:"data,omitempty"`
	Data0              bool     `json:"data0,omitempty"`
	Detail             bool     `json:"detail,omitempty"`
	DHCP               bool     `json:"dhcp,omitempty"`
	Direct             bool     `json:"direct,omitempty"`
	Distance           int      `json:"distance,omitempty"`
	Dynamic            bool     `json:"dynamic,omitempty"`
	FailedProbes       int      `json:"failedprobes,omitempty"`
	Flags              bool     `json:"flags,omitempty"`
	Gateway            string   `json:"gateway,omitempty"`
	GatewayName        string   `json:"gatewayname,omitempty"`
	ISIS               bool     `json:"isis,omitempty"`
	LBRoute            bool     `json:"lbroute,omitempty"`
	Mgmt               bool     `json:"mgmt,omitempty"`
	Monitor            string   `json:"monitor,omitempty"`
	MonStatCode        int      `json:"monstatcode,omitempty"`
	MonStatParam1      int      `json:"monstatparam1,omitempty"`
	MonStatParam2      int      `json:"monstatparam2,omitempty"`
	MonStatParam3      int      `json:"monstatparam3,omitempty"`
	MSR                string   `json:"msr,omitempty"`
	NAT                bool     `json:"nat,omitempty"`
	Netmask            string   `json:"netmask,omitempty"`
	Network            string   `json:"network,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	OSPF               bool     `json:"ospf,omitempty"`
	OwnerGroup         string   `json:"ownergroup,omitempty"`
	Permanent          bool     `json:"permanent,omitempty"`
	Protocol           []string `json:"protocol,omitempty"`
	Retain             int      `json:"retain,omitempty"`
	RIP                bool     `json:"rip,omitempty"`
	RouteOwners        []string `json:"routeowners,omitempty"`
	RouteType          string   `json:"routetype,omitempty"`
	State              int      `json:"state,omitempty"`
	Static             bool     `json:"Static,omitempty"`
	TD                 int      `json:"td,omitempty"`
	TotalFailedProbes  int      `json:"totalfailedprobes,omitempty"`
	TotalProbes        int      `json:"totalprobes,omitempty"`
	Tunnel             bool     `json:"tunnel,omitempty"`
	TypeField          bool     `json:"type,omitempty"`
	VLAN               int      `json:"vlan,omitempty"`
	Weight             int      `json:"weight,omitempty"`
}

type NetBridgeBinding struct {
	Name                     string        `json:"name,omitempty"`
	NetBridgeIPTunnelBinding []interface{} `json:"netbridge_iptunnel_binding,omitempty"`
	NetBridgeNSIP6Binding    []interface{} `json:"netbridge_nsip6_binding,omitempty"`
	NetBridgeNSIPBinding     []interface{} `json:"netbridge_nsip_binding,omitempty"`
	NetBridgeVLANBinding     []interface{} `json:"netbridge_vlan_binding,omitempty"`
}

type LinkSetChannelBinding struct {
	ID    string `json:"id,omitempty"`
	IFNum string `json:"ifnum,omitempty"`
}

type IP6Tunnel struct {
	Count              float64 `json:"__count,omitempty"`
	EncapIP            string  `json:"encapip,omitempty"`
	Local              string  `json:"local,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	OwnerGroup         string  `json:"ownergroup,omitempty"`
	Remote             string  `json:"remote,omitempty"`
	RemoteIP           string  `json:"remoteip,omitempty"`
	TypeField          int     `json:"type,omitempty"`
}

type L4Param struct {
	L2ConnMethod       string `json:"l2connmethod,omitempty"`
	L4Switch           string `json:"l4switch,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
}

type VXLANVLANMap struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
}

type MapBMRBinding struct {
	MapBMRBMRV4NetworkBinding []interface{} `json:"mapbmr_bmrv4network_binding,omitempty"`
	Name                      string        `json:"name,omitempty"`
}

type VRIDInterfaceBinding struct {
	Flags int    `json:"flags,omitempty"`
	ID    int    `json:"id,omitempty"`
	IFNum string `json:"ifnum,omitempty"`
	VLAN  int    `json:"vlan,omitempty"`
}

type BridgeTable struct {
	BridgeAge          int     `json:"bridgeage,omitempty"`
	Channel            int     `json:"channel,omitempty"`
	ControlPlane       bool    `json:"controlplane,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	DeviceVLAN         int     `json:"devicevlan,omitempty"`
	Flags              int     `json:"flags,omitempty"`
	IFNum              string  `json:"ifnum,omitempty"`
	MAC                string  `json:"mac,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NodeID             int     `json:"nodeid,omitempty"`
	TypeField          string  `json:"type,omitempty"`
	VLAN               int     `json:"vlan,omitempty"`
	VNI                int     `json:"vni,omitempty"`
	VTEP               string  `json:"vtep,omitempty"`
	VXLAN              int     `json:"vxlan,omitempty"`
}

type VXLANVLANMapBinding struct {
	Name                     string        `json:"name,omitempty"`
	VXLANVLANMapVXLANBinding []interface{} `json:"vxlanvlanmap_vxlan_binding,omitempty"`
}

type VRID struct {
	All                  bool    `json:"all,omitempty"`
	Count                float64 `json:"__count,omitempty"`
	EffectivePriority    int     `json:"effectivepriority,omitempty"`
	Flags                int     `json:"flags,omitempty"`
	ID                   int     `json:"id,omitempty"`
	Ifaces               string  `json:"ifaces,omitempty"`
	IPAddress            string  `json:"ipaddress,omitempty"`
	NextGenAPIResource   string  `json:"_nextgenapiresource,omitempty"`
	OperationalOwnerNode int     `json:"operationalownernode,omitempty"`
	OwnerNode            int     `json:"ownernode,omitempty"`
	Preemption           string  `json:"preemption,omitempty"`
	PreemptionDelayTimer int     `json:"preemptiondelaytimer,omitempty"`
	Priority             int     `json:"priority,omitempty"`
	Sharing              string  `json:"sharing,omitempty"`
	State                int     `json:"state,omitempty"`
	TrackIFNumPriority   int     `json:"trackifnumpriority,omitempty"`
	Tracking             string  `json:"tracking,omitempty"`
	TypeField            string  `json:"type,omitempty"`
}

type NetBridgeNSIP6Binding struct {
	IPAddress string `json:"ipaddress,omitempty"`
	Name      string `json:"name,omitempty"`
	Netmask   string `json:"netmask,omitempty"`
}

type INAT struct {
	ConnFailover       string  `json:"connfailover,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Flags              int     `json:"flags,omitempty"`
	FTP                string  `json:"ftp,omitempty"`
	Mode               string  `json:"mode,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	PrivateIP          string  `json:"privateip,omitempty"`
	ProxyIP            string  `json:"proxyip,omitempty"`
	PublicIP           string  `json:"publicip,omitempty"`
	TCPProxy           string  `json:"tcpproxy,omitempty"`
	TD                 int     `json:"td,omitempty"`
	TFTP               string  `json:"tftp,omitempty"`
	UseProxyPort       string  `json:"useproxyport,omitempty"`
	USIP               string  `json:"usip,omitempty"`
	USNIP              string  `json:"usnip,omitempty"`
}

type RNATRetainSourcePortSetBinding struct {
	Name                  string `json:"name,omitempty"`
	RetainSourcePortRange string `json:"retainsourceportrange,omitempty"`
}

type VLAN struct {
	AliasName          string  `json:"aliasname,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	DynamicRouting     string  `json:"dynamicrouting,omitempty"`
	ID                 int     `json:"id,omitempty"`
	Ifaces             string  `json:"ifaces,omitempty"`
	IFNum              string  `json:"ifnum,omitempty"`
	IPv6DynamicRouting string  `json:"ipv6dynamicrouting,omitempty"`
	LinkLocalIPv6Addr  string  `json:"linklocalipv6addr,omitempty"`
	LSBitmap           int     `json:"lsbitmap,omitempty"`
	LSTagBitmap        int     `json:"lstagbitmap,omitempty"`
	MTU                int     `json:"mtu,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	PartitionName      string  `json:"partitionname,omitempty"`
	PortBitmap         int     `json:"portbitmap,omitempty"`
	RNAT               bool    `json:"rnat,omitempty"`
	SDXVLAN            string  `json:"sdxvlan,omitempty"`
	Sharing            string  `json:"sharing,omitempty"`
	TagBitmap          int     `json:"tagbitmap,omitempty"`
	Tagged             bool    `json:"tagged,omitempty"`
	TagIfaces          string  `json:"tagifaces,omitempty"`
	VLANTD             int     `json:"vlantd,omitempty"`
	VXLAN              int     `json:"vxlan,omitempty"`
}

type PortAllocation struct {
	Count              float64 `json:"__count,omitempty"`
	DestIP             string  `json:"destip,omitempty"`
	DestPort           int     `json:"destport,omitempty"`
	FreePorts          int     `json:"freeports,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NodeID             int     `json:"nodeid,omitempty"`
	Protocol           int     `json:"protocol,omitempty"`
	SrcIP              string  `json:"srcip,omitempty"`
}

type RNATParam struct {
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	SrcIPPersistency   string `json:"srcippersistency,omitempty"`
	TCPProxy           string `json:"tcpproxy,omitempty"`
}

type VRID6NSIP6Binding struct {
	Flags     int    `json:"flags,omitempty"`
	ID        int    `json:"id,omitempty"`
	IPAddress string `json:"ipaddress,omitempty"`
}

type VXLANNSIP6Binding struct {
	ID        int    `json:"id,omitempty"`
	IPAddress string `json:"ipaddress,omitempty"`
	Netmask   string `json:"netmask,omitempty"`
}

type VLANBinding struct {
	ID                   int           `json:"id,omitempty"`
	VLANChannelBinding   []interface{} `json:"vlan_channel_binding,omitempty"`
	VLANInterfaceBinding []interface{} `json:"vlan_interface_binding,omitempty"`
	VLANLinkSetBinding   []interface{} `json:"vlan_linkset_binding,omitempty"`
	VLANNSIP6Binding     []interface{} `json:"vlan_nsip6_binding,omitempty"`
	VLANNSIPBinding      []interface{} `json:"vlan_nsip_binding,omitempty"`
}

type RNATBinding struct {
	Name                           string        `json:"name,omitempty"`
	RNATNSIPBinding                []interface{} `json:"rnat_nsip_binding,omitempty"`
	RNATRetainSourcePortSetBinding []interface{} `json:"rnat_retainsourceportset_binding,omitempty"`
}

type NetBridgeVLANBinding struct {
	Name string `json:"name,omitempty"`
	VLAN int    `json:"vlan,omitempty"`
}

type Interface struct {
	ActDuplex                 string   `json:"actduplex,omitempty"`
	ActFlowCtl                string   `json:"actflowctl,omitempty"`
	ActMedia                  string   `json:"actmedia,omitempty"`
	ActSpeed                  string   `json:"actspeed,omitempty"`
	ActThroughput             int      `json:"actthroughput,omitempty"`
	ActualMTU                 int      `json:"actualmtu,omitempty"`
	ActualRingSize            int      `json:"actualringsize,omitempty"`
	Autoneg                   string   `json:"autoneg,omitempty"`
	AutonegResult             int      `json:"autonegresult,omitempty"`
	Backplane                 string   `json:"backplane,omitempty"`
	BandwidthHigh             int      `json:"bandwidthhigh,omitempty"`
	BandwidthNormal           int      `json:"bandwidthnormal,omitempty"`
	BdgMacMoved               int      `json:"bdgmacmoved,omitempty"`
	BdgMuted                  int      `json:"bdgmuted,omitempty"`
	ClearTime                 int      `json:"cleartime,omitempty"`
	Count                     float64  `json:"__count,omitempty"`
	Description               string   `json:"description,omitempty"`
	DeviceName                string   `json:"devicename,omitempty"`
	Downtime                  int      `json:"downtime,omitempty"`
	Duplex                    string   `json:"duplex,omitempty"`
	FCtls                     int      `json:"fctls,omitempty"`
	Flags                     int      `json:"flags,omitempty"`
	FlowCtl                   string   `json:"flowctl,omitempty"`
	HAHeartbeat               string   `json:"haheartbeat,omitempty"`
	HAMonitor                 string   `json:"hamonitor,omitempty"`
	HangDetect                int      `json:"hangdetect,omitempty"`
	HangReset                 int      `json:"hangreset,omitempty"`
	Hangs                     int      `json:"hangs,omitempty"`
	ID                        string   `json:"id,omitempty"`
	IFAlias                   string   `json:"ifalias,omitempty"`
	IFNum                     []string `json:"ifnum,omitempty"`
	InDisc                    int      `json:"indisc,omitempty"`
	IntfState                 int      `json:"intfstate,omitempty"`
	IntfType                  string   `json:"intftype,omitempty"`
	LACPActorAggregation      string   `json:"lacpactoraggregation,omitempty"`
	LACPActorCollecting       string   `json:"lacpactorcollecting,omitempty"`
	LACPActorDistributing     string   `json:"lacpactordistributing,omitempty"`
	LACPActorInSync           string   `json:"lacpactorinsync,omitempty"`
	LACPActorMode             string   `json:"lacpactormode,omitempty"`
	LACPActorPortNo           int      `json:"lacpactorportno,omitempty"`
	LACPActorPriority         int      `json:"lacpactorpriority,omitempty"`
	LACPActorTimeout          string   `json:"lacpactortimeout,omitempty"`
	LACPKey                   int      `json:"lacpkey,omitempty"`
	LACPMode                  string   `json:"lacpmode,omitempty"`
	LACPPartnerAggregation    string   `json:"lacppartneraggregation,omitempty"`
	LACPPartnerCollecting     string   `json:"lacppartnercollecting,omitempty"`
	LACPPartnerDefaulted      string   `json:"lacppartnerdefaulted,omitempty"`
	LACPPartnerDistributing   string   `json:"lacppartnerdistributing,omitempty"`
	LACPPartnerExpired        string   `json:"lacppartnerexpired,omitempty"`
	LACPPartnerInSync         string   `json:"lacppartnerinsync,omitempty"`
	LACPPartnerKey            int      `json:"lacppartnerkey,omitempty"`
	LACPPartnerPortNo         int      `json:"lacppartnerportno,omitempty"`
	LACPPartnerPriority       int      `json:"lacppartnerpriority,omitempty"`
	LACPPartnerState          string   `json:"lacppartnerstate,omitempty"`
	LACPPartnerSystemMAC      string   `json:"lacppartnersystemmac,omitempty"`
	LACPPartnerSystemPriority int      `json:"lacppartnersystempriority,omitempty"`
	LACPPartnerTimeout        string   `json:"lacppartnertimeout,omitempty"`
	LACPPortMuxState          string   `json:"lacpportmuxstate,omitempty"`
	LACPPortRxStat            string   `json:"lacpportrxstat,omitempty"`
	LACPPortSelectState       string   `json:"lacpportselectstate,omitempty"`
	LACPPriority              int      `json:"lacppriority,omitempty"`
	LACPTimeout               string   `json:"lacptimeout,omitempty"`
	LAGType                   string   `json:"lagtype,omitempty"`
	LinkRedundancy            string   `json:"linkredundancy,omitempty"`
	LinkState                 int      `json:"linkstate,omitempty"`
	LLDPMode                  string   `json:"lldpmode,omitempty"`
	LRActiveIntf              int      `json:"lractiveintf,omitempty"`
	LRSetPriority             int      `json:"lrsetpriority,omitempty"`
	MAC                       string   `json:"mac,omitempty"`
	MACDistr                  string   `json:"macdistr,omitempty"`
	Media                     string   `json:"media,omitempty"`
	Mode                      string   `json:"mode,omitempty"`
	MTU                       int      `json:"mtu,omitempty"`
	NextGenAPIResource        string   `json:"_nextgenapiresource,omitempty"`
	OutDisc                   int      `json:"outdisc,omitempty"`
	ReqDuplex                 string   `json:"reqduplex,omitempty"`
	ReqFlowControl            string   `json:"reqflowcontrol,omitempty"`
	ReqMedia                  string   `json:"reqmedia,omitempty"`
	ReqSpeed                  string   `json:"reqspeed,omitempty"`
	ReqThroughput             int      `json:"reqthroughput,omitempty"`
	RingSize                  int      `json:"ringsize,omitempty"`
	RingType                  string   `json:"ringtype,omitempty"`
	RxBytes                   int      `json:"rxbytes,omitempty"`
	RxDrops                   int      `json:"rxdrops,omitempty"`
	RxErrors                  int      `json:"rxerrors,omitempty"`
	RxPackets                 int      `json:"rxpackets,omitempty"`
	RxStalls                  int      `json:"rxstalls,omitempty"`
	SlaveDuplex               int      `json:"slaveduplex,omitempty"`
	SlaveFlowCtl              int      `json:"slaveflowctl,omitempty"`
	SlaveMedia                int      `json:"slavemedia,omitempty"`
	SlaveSpeed                int      `json:"slavespeed,omitempty"`
	SlaveState                int      `json:"slavestate,omitempty"`
	SlaveTime                 int      `json:"slavetime,omitempty"`
	Speed                     string   `json:"speed,omitempty"`
	State                     string   `json:"state,omitempty"`
	StsStalls                 int      `json:"stsstalls,omitempty"`
	SVMCmd                    int      `json:"svmcmd,omitempty"`
	TagAll                    string   `json:"tagall,omitempty"`
	Tagged                    int      `json:"tagged,omitempty"`
	TaggedAny                 int      `json:"taggedany,omitempty"`
	TaggedAutoLearn           int      `json:"taggedautolearn,omitempty"`
	Throughput                int      `json:"throughput,omitempty"`
	Trunk                     string   `json:"trunk,omitempty"`
	TrunkAllowedVLAN          []string `json:"trunkallowedvlan,omitempty"`
	TrunkMode                 string   `json:"trunkmode,omitempty"`
	TxBytes                   int      `json:"txbytes,omitempty"`
	TxDrops                   int      `json:"txdrops,omitempty"`
	TxErrors                  int      `json:"txerrors,omitempty"`
	TxPackets                 int      `json:"txpackets,omitempty"`
	TxStalls                  int      `json:"txstalls,omitempty"`
	Unit                      int      `json:"unit,omitempty"`
	Uptime                    int      `json:"uptime,omitempty"`
	VLAN                      int      `json:"vlan,omitempty"`
	VMAC                      string   `json:"vmac,omitempty"`
	VMAC6                     string   `json:"vmac6,omitempty"`
}

type ND6 struct {
	Channel            int     `json:"channel,omitempty"`
	ControlPlane       bool    `json:"controlplane,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Flags              int     `json:"flags,omitempty"`
	IFNum              string  `json:"ifnum,omitempty"`
	MAC                string  `json:"mac,omitempty"`
	Neighbor           string  `json:"neighbor,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NodeID             int     `json:"nodeid,omitempty"`
	State              string  `json:"state,omitempty"`
	TD                 int     `json:"td,omitempty"`
	Timeout            int     `json:"timeout,omitempty"`
	VLAN               int     `json:"vlan,omitempty"`
	VTEP               string  `json:"vtep,omitempty"`
	VXLAN              int     `json:"vxlan,omitempty"`
}

type InterfacePair struct {
	Count              float64  `json:"__count,omitempty"`
	ID                 int      `json:"id,omitempty"`
	Ifaces             string   `json:"ifaces,omitempty"`
	IFNum              []string `json:"ifnum,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
}

type MapBMR struct {
	Count              float64 `json:"__count,omitempty"`
	EABitLength        int     `json:"eabitlength,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	PSIDLength         int     `json:"psidlength,omitempty"`
	PSIDOffset         int     `json:"psidoffset,omitempty"`
	RuleIPv6Prefix     string  `json:"ruleipv6prefix,omitempty"`
}

type VxlanBinding struct {
	ID                   int           `json:"id,omitempty"`
	VxlanIPTunnelBinding []interface{} `json:"vxlan_iptunnel_binding,omitempty"`
	VxlanNSIP6Binding    []interface{} `json:"vxlan_nsip6_binding,omitempty"`
	VxlanNSIPBinding     []interface{} `json:"vxlan_nsip_binding,omitempty"`
	VxlanSrcIPBinding    []interface{} `json:"vxlan_srcip_binding,omitempty"`
}

type FIS struct {
	Count              float64 `json:"__count,omitempty"`
	Ifaces             string  `json:"ifaces,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	OwnerNode          int     `json:"ownernode,omitempty"`
}

type LinkSet struct {
	Count              float64 `json:"__count,omitempty"`
	ID                 string  `json:"id,omitempty"`
	IFNum              string  `json:"ifnum,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
}

type BridgeGroupNSIPBinding struct {
	ID         int    `json:"id,omitempty"`
	IPAddress  string `json:"ipaddress,omitempty"`
	Netmask    string `json:"netmask,omitempty"`
	OwnerGroup string `json:"ownergroup,omitempty"`
	RNAT       bool   `json:"rnat,omitempty"`
	TD         int    `json:"td,omitempty"`
}

type OnLinkIPv6Prefix struct {
	AutonomusPrefix          string  `json:"autonomusprefix,omitempty"`
	Count                    float64 `json:"__count,omitempty"`
	DecrementPrefixLifetimes string  `json:"decrementprefixlifetimes,omitempty"`
	DepricatePrefix          string  `json:"depricateprefix,omitempty"`
	IPv6Prefix               string  `json:"ipv6prefix,omitempty"`
	NextGenAPIResource       string  `json:"_nextgenapiresource,omitempty"`
	OnLinkPrefix             string  `json:"onlinkprefix,omitempty"`
	PrefixCurrPreferredLft   int     `json:"prefixcurrpreferredlft,omitempty"`
	PrefixCurrValidLft       int     `json:"prefixcurrvalidelft,omitempty"`
	PrefixPreferredLifetime  int     `json:"prefixpreferredlifetime,omitempty"`
	PrefixValidLifetime      int     `json:"prefixvalidelifetime,omitempty"`
}

type VRID6NSIPBinding struct {
	Flags     int    `json:"flags,omitempty"`
	ID        int    `json:"id,omitempty"`
	IPAddress string `json:"ipaddress,omitempty"`
}

type CI struct {
	Count              float64 `json:"__count,omitempty"`
	Ifaces             string  `json:"ifaces,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
}

type IPv6 struct {
	BaseReachTime        int     `json:"basereachtime,omitempty"`
	Count                float64 `json:"__count,omitempty"`
	DODAD                string  `json:"dodad,omitempty"`
	NatPrefix            string  `json:"natprefix,omitempty"`
	NDBaseReachTime      int     `json:"ndbasereachtime,omitempty"`
	NDReachTime          int     `json:"ndreachtime,omitempty"`
	NDRetransmissionTime int     `json:"ndretransmissiontime,omitempty"`
	NextGenAPIResource   string  `json:"_nextgenapiresource,omitempty"`
	RALearning           string  `json:"ralearning,omitempty"`
	ReachTime            int     `json:"reachtime,omitempty"`
	RetransmissionTime   int     `json:"retransmissiontime,omitempty"`
	RouterRedirection    string  `json:"routerredirection,omitempty"`
	TD                   int     `json:"td,omitempty"`
	USIPNatPrefix        string  `json:"usipnatprefix,omitempty"`
}

type LinkSetBinding struct {
	ID                      string        `json:"id,omitempty"`
	LinkSetChannelBinding   []interface{} `json:"linkset_channel_binding,omitempty"`
	LinkSetInterfaceBinding []interface{} `json:"linkset_interface_binding,omitempty"`
}

type VRID6InterfaceBinding struct {
	Flags int    `json:"flags,omitempty"`
	ID    int    `json:"id,omitempty"`
	IFNum string `json:"ifnum,omitempty"`
	VLAN  int    `json:"vlan,omitempty"`
}

type RNATGlobalBinding struct {
	RNATGlobalAuditSyslogPolicyBinding []interface{} `json:"rnatglobal_auditsyslogpolicy_binding,omitempty"`
}

type FISChannelBinding struct {
	IFNum     string `json:"ifnum,omitempty"`
	Name      string `json:"name,omitempty"`
	OwnerNode int    `json:"ownernode,omitempty"`
}

type MapDomainBinding struct {
	MapDomainMapBMRBinding []interface{} `json:"mapdomain_mapbmr_binding,omitempty"`
	Name                   string        `json:"name,omitempty"`
}

type MapDomainMapBMRBinding struct {
	MapBMRName string `json:"mapbmrname,omitempty"`
	Name       string `json:"name,omitempty"`
}
