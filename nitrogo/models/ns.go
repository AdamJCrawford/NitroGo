package models

// ns configuration structs
type NSPBR struct {
	Action             string  `json:"action,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	CurState           int     `json:"curstate,omitempty"`
	Data               bool    `json:"data,omitempty"`
	DestIP             bool    `json:"destip,omitempty"`
	DestIPOp           string  `json:"destipop,omitempty"`
	DestIPVal          string  `json:"destipval,omitempty"`
	DestPort           bool    `json:"destport,omitempty"`
	DestPortOp         string  `json:"destportop,omitempty"`
	DestPortVal        string  `json:"destportval,omitempty"`
	Detail             bool    `json:"detail,omitempty"`
	FailedProbes       int     `json:"failedprobes,omitempty"`
	Hits               int     `json:"hits,omitempty"`
	Interface          string  `json:"Interface,omitempty"`
	IPTunnel           bool    `json:"iptunnel,omitempty"`
	IPTunnelName       string  `json:"iptunnelname,omitempty"`
	KernelState        string  `json:"kernelstate,omitempty"`
	Monitor            string  `json:"monitor,omitempty"`
	MonStatCode        int     `json:"monstatcode,omitempty"`
	MonStatParam1      int     `json:"monstatparam1,omitempty"`
	MonStatParam2      int     `json:"monstatparam2,omitempty"`
	MonStatParam3      int     `json:"monstatparam3,omitempty"`
	MSR                string  `json:"msr,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NextHop            bool    `json:"nexthop,omitempty"`
	NextHopVal         string  `json:"nexthopval,omitempty"`
	OwnerGroup         string  `json:"ownergroup,omitempty"`
	Priority           int     `json:"priority,omitempty"`
	Protocol           string  `json:"protocol,omitempty"`
	ProtocolNumber     int     `json:"protocolnumber,omitempty"`
	SrcIP              bool    `json:"srcip,omitempty"`
	SrcIPOp            string  `json:"srcipop,omitempty"`
	SrcIPVal           string  `json:"srcipval,omitempty"`
	SrcMac             string  `json:"srcmac,omitempty"`
	SrcMacMask         string  `json:"srcmacmask,omitempty"`
	SrcPort            bool    `json:"srcport,omitempty"`
	SrcPortOp          string  `json:"srcportop,omitempty"`
	SrcPortVal         string  `json:"srcportval,omitempty"`
	State              string  `json:"state,omitempty"`
	TargetTD           int     `json:"targettd,omitempty"`
	TD                 int     `json:"td,omitempty"`
	TotalFailedProbes  int     `json:"totalfailedprobes,omitempty"`
	TotalProbes        int     `json:"totalprobes,omitempty"`
	VLAN               int     `json:"vlan,omitempty"`
	VXLAN              int     `json:"vxlan,omitempty"`
	VXLANVLANMap       string  `json:"vxlanvlanmap,omitempty"`
}

type NSTrafficDomainVXLANBinding struct {
	TD    int `json:"td,omitempty"`
	VXLAN int `json:"vxlan,omitempty"`
}

type NSPartition struct {
	Count              float64 `json:"__count,omitempty"`
	Force              bool    `json:"force,omitempty"`
	MaxBandwidth       int     `json:"maxbandwidth,omitempty"`
	MaxConn            int     `json:"maxconn,omitempty"`
	MaxMemLimit        int     `json:"maxmemlimit,omitempty"`
	MinBandwidth       int     `json:"minbandwidth,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	PartitionID        int     `json:"partitionid,omitempty"`
	PartitionMAC       string  `json:"partitionmac,omitempty"`
	PartitionName      string  `json:"partitionname,omitempty"`
	PartitionType      string  `json:"partitiontype,omitempty"`
	PMACInternal       bool    `json:"pmacinternal,omitempty"`
	Save               bool    `json:"save,omitempty"`
}

type NSExtensionBinding struct {
	Name                                string        `json:"name,omitempty"`
	NSExtensionExtensionFunctionBinding []interface{} `json:"nsextension_extensionfunction_binding,omitempty"`
}

type NSLicense struct {
	AAA                     bool   `json:"aaa,omitempty"`
	AdaptiveTCP             bool   `json:"adaptivetcp,omitempty"`
	AGEE                    bool   `json:"agee,omitempty"`
	APIGateway              bool   `json:"apigateway,omitempty"`
	AppFlow                 bool   `json:"appflow,omitempty"`
	AppFlowICA              bool   `json:"appflowica,omitempty"`
	AppFW                   bool   `json:"appfw,omitempty"`
	AppQOE                  bool   `json:"appqoe,omitempty"`
	BGP                     bool   `json:"bgp,omitempty"`
	Bot                     bool   `json:"bot,omitempty"`
	CF                      bool   `json:"cf,omitempty"`
	CH                      bool   `json:"ch,omitempty"`
	CloudBridge             bool   `json:"cloudbridge,omitempty"`
	CloudBridgeAppliance    bool   `json:"cloudbridgeappliance,omitempty"`
	CloudExtenderAppliance  bool   `json:"cloudextenderappliance,omitempty"`
	CloudSubscriptionImage  string `json:"cloudsubscriptionimage,omitempty"`
	Cluster                 bool   `json:"cluster,omitempty"`
	CMP                     bool   `json:"cmp,omitempty"`
	ContentAccelerator      bool   `json:"contentaccelerator,omitempty"`
	CQA                     bool   `json:"cqa,omitempty"`
	CR                      bool   `json:"cr,omitempty"`
	CS                      bool   `json:"cs,omitempty"`
	DaysToExpiration        int    `json:"daystoexpiration,omitempty"`
	DaysToLASEnforcement    int    `json:"daystolasenforcement,omitempty"`
	Delta                   bool   `json:"delta,omitempty"`
	FICAUsers               int    `json:"f_ica_users,omitempty"`
	FSSLVPNUsers            int    `json:"f_sslvpn_users,omitempty"`
	FEO                     bool   `json:"feo,omitempty"`
	ForwardProxy            bool   `json:"forwardproxy,omitempty"`
	GSLB                    bool   `json:"gslb,omitempty"`
	GSLBP                   bool   `json:"gslbp,omitempty"`
	IC                      bool   `json:"ic,omitempty"`
	IPv6PT                  bool   `json:"ipv6pt,omitempty"`
	IsEnterpriseLic         bool   `json:"isenterpriselic,omitempty"`
	ISIS                    bool   `json:"isis,omitempty"`
	IsPlatinumLic           bool   `json:"isplatinumlic,omitempty"`
	IsSGWYLic               bool   `json:"issgwylic,omitempty"`
	IsStandardLic           bool   `json:"isstandardlic,omitempty"`
	IsSWGLic                bool   `json:"isswglic,omitempty"`
	LB                      bool   `json:"lb,omitempty"`
	LicensingMode           string `json:"licensingmode,omitempty"`
	LSN                     bool   `json:"lsn,omitempty"`
	ModelID                 int    `json:"modelid,omitempty"`
	NextGenAPIResource      string `json:"_nextgenapiresource,omitempty"`
	NSXN                    bool   `json:"nsxn,omitempty"`
	OSPF                    bool   `json:"ospf,omitempty"`
	Push                    bool   `json:"push,omitempty"`
	RDPProxy                bool   `json:"rdpproxy,omitempty"`
	RemoteContentInspection bool   `json:"remotecontentinspection,omitempty"`
	Rep                     bool   `json:"rep,omitempty"`
	Responder               bool   `json:"responder,omitempty"`
	Rewrite                 bool   `json:"rewrite,omitempty"`
	RIP                     bool   `json:"rip,omitempty"`
	Routing                 bool   `json:"routing,omitempty"`
	SP                      bool   `json:"sp,omitempty"`
	SSL                     bool   `json:"ssl,omitempty"`
	SSLInterception         bool   `json:"sslinterception,omitempty"`
	SSLVPN                  bool   `json:"sslvpn,omitempty"`
	URLFiltering            bool   `json:"urlfiltering,omitempty"`
	VideoOptimization       bool   `json:"videooptimization,omitempty"`
	WL                      bool   `json:"wl,omitempty"`
}

type NSTestLicense struct {
	AAA                     bool   `json:"aaa,omitempty"`
	AdaptiveTCP             bool   `json:"adaptivetcp,omitempty"`
	AGEE                    bool   `json:"agee,omitempty"`
	APIGateway              bool   `json:"apigateway,omitempty"`
	AppFlow                 bool   `json:"appflow,omitempty"`
	AppFlowICA              bool   `json:"appflowica,omitempty"`
	AppFW                   bool   `json:"appfw,omitempty"`
	AppQOE                  bool   `json:"appqoe,omitempty"`
	BGP                     bool   `json:"bgp,omitempty"`
	Bot                     bool   `json:"bot,omitempty"`
	CF                      bool   `json:"cf,omitempty"`
	CH                      bool   `json:"ch,omitempty"`
	CloudBridge             bool   `json:"cloudbridge,omitempty"`
	CloudBridgeAppliance    bool   `json:"cloudbridgeappliance,omitempty"`
	CloudExtenderAppliance  bool   `json:"cloudextenderappliance,omitempty"`
	Cluster                 bool   `json:"cluster,omitempty"`
	CMP                     bool   `json:"cmp,omitempty"`
	ContentAccelerator      bool   `json:"contentaccelerator,omitempty"`
	CQA                     bool   `json:"cqa,omitempty"`
	CR                      bool   `json:"cr,omitempty"`
	CS                      bool   `json:"cs,omitempty"`
	DaysToExpiration        int    `json:"daystoexpiration,omitempty"`
	Delta                   bool   `json:"delta,omitempty"`
	FICAUsers               int    `json:"f_ica_users,omitempty"`
	FSSLVPNUsers            int    `json:"f_sslvpn_users,omitempty"`
	FEO                     bool   `json:"feo,omitempty"`
	ForwardProxy            bool   `json:"forwardproxy,omitempty"`
	GSLB                    bool   `json:"gslb,omitempty"`
	GSLBP                   bool   `json:"gslbp,omitempty"`
	IC                      bool   `json:"ic,omitempty"`
	IPv6PT                  bool   `json:"ipv6pt,omitempty"`
	IsEnterpriseLic         bool   `json:"isenterpriselic,omitempty"`
	ISIS                    bool   `json:"isis,omitempty"`
	IsPlatinumLic           bool   `json:"isplatinumlic,omitempty"`
	IsSGWYLic               bool   `json:"issgwylic,omitempty"`
	IsStandardLic           bool   `json:"isstandardlic,omitempty"`
	IsSWGLic                bool   `json:"isswglic,omitempty"`
	LB                      bool   `json:"lb,omitempty"`
	LicensingMode           string `json:"licensingmode,omitempty"`
	LSN                     bool   `json:"lsn,omitempty"`
	ModelID                 int    `json:"modelid,omitempty"`
	NextGenAPIResource      string `json:"_nextgenapiresource,omitempty"`
	NSXN                    bool   `json:"nsxn,omitempty"`
	OSPF                    bool   `json:"ospf,omitempty"`
	Push                    bool   `json:"push,omitempty"`
	RDPProxy                bool   `json:"rdpproxy,omitempty"`
	RemoteContentInspection bool   `json:"remotecontentinspection,omitempty"`
	Rep                     bool   `json:"rep,omitempty"`
	Responder               bool   `json:"responder,omitempty"`
	Rewrite                 bool   `json:"rewrite,omitempty"`
	RIP                     bool   `json:"rip,omitempty"`
	Routing                 bool   `json:"routing,omitempty"`
	SP                      bool   `json:"sp,omitempty"`
	SSL                     bool   `json:"ssl,omitempty"`
	SSLInterception         bool   `json:"sslinterception,omitempty"`
	SSLVPN                  bool   `json:"sslvpn,omitempty"`
	URLFiltering            bool   `json:"urlfiltering,omitempty"`
	VideoOptimization       bool   `json:"videooptimization,omitempty"`
	WL                      bool   `json:"wl,omitempty"`
}

type NSLicenseServerPool struct {
	CPXInstanceAvailable         int    `json:"cpxinstanceavailable,omitempty"`
	CPXInstanceTotal             int    `json:"cpxinstancetotal,omitempty"`
	EnterpriseBandwidthAvailable int    `json:"enterprisebandwidthavailable,omitempty"`
	EnterpriseBandwidthTotal     int    `json:"enterprisebandwidthtotal,omitempty"`
	EnterpriseCPUAvailable       int    `json:"enterprisecpuavailable,omitempty"`
	EnterpriseCPUTotal           int    `json:"enterprisecputotal,omitempty"`
	GetAllLicenses               bool   `json:"getalllicenses,omitempty"`
	InstanceAvailable            int    `json:"instanceavailable,omitempty"`
	InstanceTotal                int    `json:"instancetotal,omitempty"`
	LicenseMode                  string `json:"licensemode,omitempty"`
	NextGenAPIResource           string `json:"_nextgenapiresource,omitempty"`
	PlatinumBandwidthAvailable   int    `json:"platinumbandwidthavailable,omitempty"`
	PlatinumBandwidthTotal       int    `json:"platinumbandwidthtotal,omitempty"`
	PlatinumCPUAvailable         int    `json:"platinumcpuavailable,omitempty"`
	PlatinumCPUTotal             int    `json:"platinumcputotal,omitempty"`
	StandardBandwidthAvailable   int    `json:"standardbandwidthavailable,omitempty"`
	StandardBandwidthTotal       int    `json:"standardbandwidthtotal,omitempty"`
	StandardCPUAvailable         int    `json:"standardcpuavailable,omitempty"`
	StandardCPUTotal             int    `json:"standardcputotal,omitempty"`
	VPX100000EAvailable          int    `json:"vpx100000eavailable,omitempty"`
	VPX100000ETotal              int    `json:"vpx100000etotal,omitempty"`
	VPX100000PAvailable          int    `json:"vpx100000pavailable,omitempty"`
	VPX100000PTotal              int    `json:"vpx100000ptotal,omitempty"`
	VPX100000SAvailable          int    `json:"vpx100000savailable,omitempty"`
	VPX100000STotal              int    `json:"vpx100000stotal,omitempty"`
	VPX10000EAvailable           int    `json:"vpx10000eavailable,omitempty"`
	VPX10000ETotal               int    `json:"vpx10000etotal,omitempty"`
	VPX10000PAvailable           int    `json:"vpx10000pavailable,omitempty"`
	VPX10000PTotal               int    `json:"vpx10000ptotal,omitempty"`
	VPX10000SAvailable           int    `json:"vpx10000savailable,omitempty"`
	VPX10000STotal               int    `json:"vpx10000stotal,omitempty"`
	VPX1000EAvailable            int    `json:"vpx1000eavailable,omitempty"`
	VPX1000ETotal                int    `json:"vpx1000etotal,omitempty"`
	VPX1000PAvailable            int    `json:"vpx1000pavailable,omitempty"`
	VPX1000PTotal                int    `json:"vpx1000ptotal,omitempty"`
	VPX1000SAvailable            int    `json:"vpx1000savailable,omitempty"`
	VPX1000STotal                int    `json:"vpx1000stotal,omitempty"`
	VPX100EAvailable             int    `json:"vpx100eavailable,omitempty"`
	VPX100ETotal                 int    `json:"vpx100etotal,omitempty"`
	VPX100PAvailable             int    `json:"vpx100pavailable,omitempty"`
	VPX100PTotal                 int    `json:"vpx100ptotal,omitempty"`
	VPX100SAvailable             int    `json:"vpx100savailable,omitempty"`
	VPX100STotal                 int    `json:"vpx100stotal,omitempty"`
	VPX10EAvailable              int    `json:"vpx10eavailable,omitempty"`
	VPX10ETotal                  int    `json:"vpx10etotal,omitempty"`
	VPX10PAvailable              int    `json:"vpx10pavailable,omitempty"`
	VPX10PTotal                  int    `json:"vpx10ptotal,omitempty"`
	VPX10SAvailable              int    `json:"vpx10savailable,omitempty"`
	VPX10STotal                  int    `json:"vpx10stotal,omitempty"`
	VPX15000EAvailable           int    `json:"vpx15000eavailable,omitempty"`
	VPX15000ETotal               int    `json:"vpx15000etotal,omitempty"`
	VPX15000PAvailable           int    `json:"vpx15000pavailable,omitempty"`
	VPX15000PTotal               int    `json:"vpx15000ptotal,omitempty"`
	VPX15000SAvailable           int    `json:"vpx15000savailable,omitempty"`
	VPX15000STotal               int    `json:"vpx15000stotal,omitempty"`
	VPX1PAvailable               int    `json:"vpx1pavailable,omitempty"`
	VPX1PTotal                   int    `json:"vpx1ptotal,omitempty"`
	VPX1SAvailable               int    `json:"vpx1savailable,omitempty"`
	VPX1STotal                   int    `json:"vpx1stotal,omitempty"`
	VPX2000PAvailable            int    `json:"vpx2000pavailable,omitempty"`
	VPX2000PTotal                int    `json:"vpx2000ptotal,omitempty"`
	VPX200EAvailable             int    `json:"vpx200eavailable,omitempty"`
	VPX200ETotal                 int    `json:"vpx200etotal,omitempty"`
	VPX200PAvailable             int    `json:"vpx200pavailable,omitempty"`
	VPX200PTotal                 int    `json:"vpx200ptotal,omitempty"`
	VPX200SAvailable             int    `json:"vpx200savailable,omitempty"`
	VPX200STotal                 int    `json:"vpx200stotal,omitempty"`
	VPX25000EAvailable           int    `json:"vpx25000eavailable,omitempty"`
	VPX25000ETotal               int    `json:"vpx25000etotal,omitempty"`
	VPX25000PAvailable           int    `json:"vpx25000pavailable,omitempty"`
	VPX25000PTotal               int    `json:"vpx25000ptotal,omitempty"`
	VPX25000SAvailable           int    `json:"vpx25000savailable,omitempty"`
	VPX25000STotal               int    `json:"vpx25000stotal,omitempty"`
	VPX25EAvailable              int    `json:"vpx25eavailable,omitempty"`
	VPX25ETotal                  int    `json:"vpx25etotal,omitempty"`
	VPX25PAvailable              int    `json:"vpx25pavailable,omitempty"`
	VPX25PTotal                  int    `json:"vpx25ptotal,omitempty"`
	VPX25SAvailable              int    `json:"vpx25savailable,omitempty"`
	VPX25STotal                  int    `json:"vpx25stotal,omitempty"`
	VPX3000EAvailable            int    `json:"vpx3000eavailable,omitempty"`
	VPX3000ETotal                int    `json:"vpx3000etotal,omitempty"`
	VPX3000PAvailable            int    `json:"vpx3000pavailable,omitempty"`
	VPX3000PTotal                int    `json:"vpx3000ptotal,omitempty"`
	VPX3000SAvailable            int    `json:"vpx3000savailable,omitempty"`
	VPX3000STotal                int    `json:"vpx3000stotal,omitempty"`
	VPX40000EAvailable           int    `json:"vpx40000eavailable,omitempty"`
	VPX40000ETotal               int    `json:"vpx40000etotal,omitempty"`
	VPX40000PAvailable           int    `json:"vpx40000pavailable,omitempty"`
	VPX40000PTotal               int    `json:"vpx40000ptotal,omitempty"`
	VPX40000SAvailable           int    `json:"vpx40000savailable,omitempty"`
	VPX40000STotal               int    `json:"vpx40000stotal,omitempty"`
	VPX4000PAvailable            int    `json:"vpx4000pavailable,omitempty"`
	VPX4000PTotal                int    `json:"vpx4000ptotal,omitempty"`
	VPX5000EAvailable            int    `json:"vpx5000eavailable,omitempty"`
	VPX5000ETotal                int    `json:"vpx5000etotal,omitempty"`
	VPX5000PAvailable            int    `json:"vpx5000pavailable,omitempty"`
	VPX5000PTotal                int    `json:"vpx5000ptotal,omitempty"`
	VPX5000SAvailable            int    `json:"vpx5000savailable,omitempty"`
	VPX5000STotal                int    `json:"vpx5000stotal,omitempty"`
	VPX500EAvailable             int    `json:"vpx500eavailable,omitempty"`
	VPX500ETotal                 int    `json:"vpx500etotal,omitempty"`
	VPX500PAvailable             int    `json:"vpx500pavailable,omitempty"`
	VPX500PTotal                 int    `json:"vpx500ptotal,omitempty"`
	VPX500SAvailable             int    `json:"vpx500savailable,omitempty"`
	VPX500STotal                 int    `json:"vpx500stotal,omitempty"`
	VPX50EAvailable              int    `json:"vpx50eavailable,omitempty"`
	VPX50ETotal                  int    `json:"vpx50etotal,omitempty"`
	VPX50PAvailable              int    `json:"vpx50pavailable,omitempty"`
	VPX50PTotal                  int    `json:"vpx50ptotal,omitempty"`
	VPX50SAvailable              int    `json:"vpx50savailable,omitempty"`
	VPX50STotal                  int    `json:"vpx50stotal,omitempty"`
	VPX5PAvailable               int    `json:"vpx5pavailable,omitempty"`
	VPX5PTotal                   int    `json:"vpx5ptotal,omitempty"`
	VPX5SAvailable               int    `json:"vpx5savailable,omitempty"`
	VPX5STotal                   int    `json:"vpx5stotal,omitempty"`
	VPX8000EAvailable            int    `json:"vpx8000eavailable,omitempty"`
	VPX8000ETotal                int    `json:"vpx8000etotal,omitempty"`
	VPX8000PAvailable            int    `json:"vpx8000pavailable,omitempty"`
	VPX8000PTotal                int    `json:"vpx8000ptotal,omitempty"`
	VPX8000SAvailable            int    `json:"vpx8000savailable,omitempty"`
	VPX8000STotal                int    `json:"vpx8000stotal,omitempty"`
}

type NSConfigView struct {
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	State              string `json:"state,omitempty"`
}

type NSTimerBinding struct {
	Name                          string        `json:"name,omitempty"`
	NSTimerAutoScalePolicyBinding []interface{} `json:"nstimer_autoscalepolicy_binding,omitempty"`
}

type NSDiameter struct {
	Count                  float64 `json:"__count,omitempty"`
	Identity               string  `json:"identity,omitempty"`
	NextGenAPIResource     string  `json:"_nextgenapiresource,omitempty"`
	OwnerNode              int     `json:"ownernode,omitempty"`
	Realm                  string  `json:"realm,omitempty"`
	ServerClosePropagation string  `json:"serverclosepropagation,omitempty"`
}

type NSTrafficDomainBridgeGroupBinding struct {
	BridgeGroup int `json:"bridgegroup,omitempty"`
	TD          int `json:"td,omitempty"`
}

type NSRollbackCmd struct {
	Filename           string `json:"filename,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	OutType            string `json:"outtype,omitempty"`
}

type NSVPXParam struct {
	CloudProductCode    string  `json:"cloudproductcode,omitempty"`
	Count               float64 `json:"__count,omitempty"`
	CPUYield            string  `json:"cpuyield,omitempty"`
	KVMVirtioMultiQueue string  `json:"kvmvirtiomultiqueue,omitempty"`
	MasterClockCPU1     string  `json:"masterclockcpu1,omitempty"`
	MemoryStatus        string  `json:"memorystatus,omitempty"`
	NextGenAPIResource  string  `json:"_nextgenapiresource,omitempty"`
	OwnerNode           int     `json:"ownernode,omitempty"`
	TechnicalSupportPIN string  `json:"technicalsupportpin,omitempty"`
	VPXEnvironment      string  `json:"vpxenvironment,omitempty"`
	VPXOEMCode          int     `json:"vpxoemcode,omitempty"`
}

type NSLicenseProxyServer struct {
	Count              float64 `json:"__count,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Port               int     `json:"port,omitempty"`
	ServerIP           string  `json:"serverip,omitempty"`
	ServerName         string  `json:"servername,omitempty"`
}

type NSLimitIdentifierNSLimitSessionsBinding struct {
	LimitIdentifier string `json:"limitidentifier,omitempty"`
}

type NSPartitionMac struct {
	Count              float64 `json:"__count,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	PartitionMAC       string  `json:"partitionmac,omitempty"`
	PartitionName      string  `json:"partitionname,omitempty"`
}

type NSServicePath struct {
	Count              float64 `json:"__count,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	ServicePathName    string  `json:"servicepathname,omitempty"`
}

type NSPartitionBinding struct {
	NSPartitionBridgeGroupBinding []interface{} `json:"nspartition_bridgegroup_binding,omitempty"`
	NSPartitionVlanBinding        []interface{} `json:"nspartition_vlan_binding,omitempty"`
	NSPartitionVxlanBinding       []interface{} `json:"nspartition_vxlan_binding,omitempty"`
	PartitionName                 string        `json:"partitionname,omitempty"`
}

type NSPartitionBridgeGroupBinding struct {
	BridgeGroup   int    `json:"bridgegroup,omitempty"`
	PartitionName string `json:"partitionname,omitempty"`
}

type NSConnectionTable struct {
	AdaptiveTCPProfName    string   `json:"adaptivetcpprofname,omitempty"`
	AdvWnd                 int      `json:"advwnd,omitempty"`
	BurstRateControl       string   `json:"burstratecontrol,omitempty"`
	BWEstimate             int      `json:"bwestimate,omitempty"`
	ChannelIDNNM           int      `json:"channelidnnm,omitempty"`
	CongState              string   `json:"congstate,omitempty"`
	ConnID                 int      `json:"connid,omitempty"`
	ConnProperties         []string `json:"connproperties,omitempty"`
	Count                  float64  `json:"__count,omitempty"`
	CQABifAvg              int      `json:"cqabifavg,omitempty"`
	CQACCL                 int      `json:"cqaccl,omitempty"`
	CQACSQ                 int      `json:"cqacsq,omitempty"`
	CQAAIAI1MSPct          int      `json:"cqaiai1mspct,omitempty"`
	CQAAIAI2MSPct          int      `json:"cqaiai2mspct,omitempty"`
	CQAAIAIAvg             int      `json:"cqaiaiavg,omitempty"`
	CQAAIAISamples         int      `json:"cqaiaisamples,omitempty"`
	CQAISIAvg              int      `json:"cqaisiavg,omitempty"`
	CQALoadDelayAvg        int      `json:"cqaloaddelayavg,omitempty"`
	CQANetClass            string   `json:"cqanetclass,omitempty"`
	CQANoiseDelayAvg       int      `json:"cqanoisedelayavg,omitempty"`
	CQARcvWndAvg           int      `json:"cqarcvwndavg,omitempty"`
	CQARcvWndMin           int      `json:"cqarcvwndmin,omitempty"`
	CQARetxCong            int      `json:"cqaretxcong,omitempty"`
	CQARetxCorr            int      `json:"cqaretxcorr,omitempty"`
	CQARetxPackets         int      `json:"cqaretxpackets,omitempty"`
	CQARTTAvg              int      `json:"cqarttavg,omitempty"`
	CQARTTMax              int      `json:"cqarttmax,omitempty"`
	CQARTTMin              int      `json:"cqarttmin,omitempty"`
	CQASamples             int      `json:"cqasamples,omitempty"`
	CQAThruputAvg          int      `json:"cqathruputavg,omitempty"`
	CreditsInBytes         int      `json:"creditsinbytes,omitempty"`
	DestIP                 string   `json:"destip,omitempty"`
	DestPort               int      `json:"destport,omitempty"`
	EntityName             string   `json:"entityname,omitempty"`
	FilterExpression       string   `json:"filterexpression,omitempty"`
	FilterName             bool     `json:"filtername,omitempty"`
	Flavor                 string   `json:"flavor,omitempty"`
	HTTPEndSeq             int      `json:"httpendseq,omitempty"`
	HTTPRequest            string   `json:"httprequest,omitempty"`
	HTTPReqVer             string   `json:"httpreqver,omitempty"`
	HTTPRspCode            int      `json:"httprspcode,omitempty"`
	HTTPState              string   `json:"httpstate,omitempty"`
	IdleTime               int      `json:"idletime,omitempty"`
	IRS                    int      `json:"irs,omitempty"`
	ISS                    int      `json:"iss,omitempty"`
	Link                   bool     `json:"link,omitempty"`
	LinkBurstRateControl   string   `json:"linkburstratecontrol,omitempty"`
	LinkBWEstimate         int      `json:"linkbwestimate,omitempty"`
	LinkCongState          string   `json:"linkcongstate,omitempty"`
	LinkConnID             int      `json:"linkconnid,omitempty"`
	LinkCredits            int      `json:"linkcredits,omitempty"`
	LinkDestIP             string   `json:"linkdestip,omitempty"`
	LinkDestPort           int      `json:"linkdestport,omitempty"`
	LinkEntityName         string   `json:"linkentityname,omitempty"`
	LinkFlavor             string   `json:"linkflavor,omitempty"`
	LinkIdleTime           int      `json:"linkidletime,omitempty"`
	LinkMaxRcvBuf          int      `json:"linkmaxrcvbuf,omitempty"`
	LinkMaxSndBuf          int      `json:"linkmaxsndbuf,omitempty"`
	LinkName               string   `json:"linkname,omitempty"`
	LinkNSBRetxQ           int      `json:"linknsbretxq,omitempty"`
	LinkNSBTCPWaitQ        int      `json:"linknsbtcpwaitq,omitempty"`
	LinkNSWSValue          int      `json:"linknswsvalue,omitempty"`
	LinkOptionFlag         []string `json:"linkoptionflag,omitempty"`
	LinkPeerWSValue        int      `json:"linkpeerwsvalue,omitempty"`
	LinkRateInBytes        int      `json:"linkrateinbytes,omitempty"`
	LinkRateSchedulerQueue int      `json:"linkrateschedulerqueue,omitempty"`
	LinkRealTimeRTT        int      `json:"linkrealtimertt,omitempty"`
	LinkRTTMin             int      `json:"linkrttmin,omitempty"`
	LinkRxQSize            int      `json:"linkrxqsize,omitempty"`
	LinkSackBlocks         int      `json:"linksackblocks,omitempty"`
	LinkServiceType        string   `json:"linkservicetype,omitempty"`
	LinkSndBuf             int      `json:"linksndbuf,omitempty"`
	LinkSndRecoverLe       int      `json:"linksndrecoverle,omitempty"`
	LinkSourceIP           string   `json:"linksourceip,omitempty"`
	LinkSourcePort         int      `json:"linksourceport,omitempty"`
	LinkState              string   `json:"linkstate,omitempty"`
	LinkTCPMode            string   `json:"linktcpmode,omitempty"`
	LinkTxQSize            int      `json:"linktxqsize,omitempty"`
	Listen                 bool     `json:"listen,omitempty"`
	MaxAck                 int      `json:"maxack,omitempty"`
	MaxRcvBuf              int      `json:"maxrcvbuf,omitempty"`
	MaxSndBuf              int      `json:"maxsndbuf,omitempty"`
	MsgVersionNNM          int      `json:"msgversionnnm,omitempty"`
	MSS                    int      `json:"mss,omitempty"`
	Name                   string   `json:"name,omitempty"`
	NextGenAPIResource     string   `json:"_nextgenapiresource,omitempty"`
	NodeID                 int      `json:"nodeid,omitempty"`
	NSBRetxQ               int      `json:"nsbretxq,omitempty"`
	NSBTCPWaitQ            int      `json:"nsbtcpwaitq,omitempty"`
	NSWSValue              int      `json:"nswsvalue,omitempty"`
	OptionFlags            []string `json:"optionflags,omitempty"`
	OutOfOrderBlocks       int      `json:"outoforderblocks,omitempty"`
	OutOfOrderBytes        int      `json:"outoforderbytes,omitempty"`
	OutOfOrderFlushedCount int      `json:"outoforderflushedcount,omitempty"`
	OutOfOrderPkts         int      `json:"outoforderpkts,omitempty"`
	PeerWSValue            int      `json:"peerwsvalue,omitempty"`
	Priority               string   `json:"priority,omitempty"`
	RateInBytes            int      `json:"rateinbytes,omitempty"`
	RateSchedulerQueue     int      `json:"rateschedulerqueue,omitempty"`
	RcvNxt                 int      `json:"rcvnxt,omitempty"`
	RcvWnd                 int      `json:"rcvwnd,omitempty"`
	RealTimeRTT            int      `json:"realtimertt,omitempty"`
	RetxRetryCnt           int      `json:"retxretrycnt,omitempty"`
	RTTMin                 int      `json:"rttmin,omitempty"`
	RTTSmoothed            int      `json:"rttsmoothed,omitempty"`
	RTTVariance            int      `json:"rttvariance,omitempty"`
	RxQSize                int      `json:"rxqsize,omitempty"`
	SackBlocks             int      `json:"sackblocks,omitempty"`
	SndBuf                 int      `json:"sndbuf,omitempty"`
	SndCwnd                int      `json:"sndcwnd,omitempty"`
	SndNxt                 int      `json:"sndnxt,omitempty"`
	SndRecoverLe           int      `json:"sndrecoverle,omitempty"`
	SndUnack               int      `json:"sndunack,omitempty"`
	SourceIP               string   `json:"sourceip,omitempty"`
	SourceNodeIDNNM        int      `json:"sourcenodeidnnm,omitempty"`
	SourcePort             int      `json:"sourceport,omitempty"`
	State                  string   `json:"state,omitempty"`
	SvcType                string   `json:"svctype,omitempty"`
	TargetNodeIDNNM        int      `json:"targetnodeidnnm,omitempty"`
	TCPMode                string   `json:"tcpmode,omitempty"`
	TD                     int      `json:"td,omitempty"`
	TRCount                int      `json:"trcount,omitempty"`
	TxQSize                int      `json:"txqsize,omitempty"`
}

type NSServicePathNSServiceFunctionBinding struct {
	Index           int    `json:"index,omitempty"`
	ServiceFunction string `json:"servicefunction,omitempty"`
	ServicePathName string `json:"servicepathname,omitempty"`
}

type NSParam struct {
	AdvancedAnalyticsStats    string        `json:"advancedanalyticsstats,omitempty"`
	AFTPAllowRandomSourcePort string        `json:"aftpallowrandomsourceport,omitempty"`
	AutoScaleOption           int           `json:"autoscaleoption,omitempty"`
	CIP                       string        `json:"cip,omitempty"`
	CIPHeader                 string        `json:"cipheader,omitempty"`
	CookieVersion             string        `json:"cookieversion,omitempty"`
	CRPortRange               string        `json:"crportrange,omitempty"`
	ExclusiveQuotaMaxClient   int           `json:"exclusivequotamaxclient,omitempty"`
	ExclusiveQuotaSpillover   int           `json:"exclusivequotaspillover,omitempty"`
	FTPPortRange              string        `json:"ftpportrange,omitempty"`
	GrantQuotaMaxClient       int           `json:"grantquotamaxclient,omitempty"`
	GrantQuotaSpillover       int           `json:"grantquotaspillover,omitempty"`
	HTTPPort                  []interface{} `json:"httpport,omitempty"`
	ICAPorts                  []interface{} `json:"icaports,omitempty"`
	InternalUserLogin         string        `json:"internaluserlogin,omitempty"`
	IPTTL                     int           `json:"ipttl,omitempty"`
	MaxConn                   int           `json:"maxconn,omitempty"`
	MaxReq                    int           `json:"maxreq,omitempty"`
	MgmtHTTPPort              int           `json:"mgmthttpport,omitempty"`
	MgmtHTTPSPort             int           `json:"mgmthttpsport,omitempty"`
	NextGenAPIResource        string        `json:"_nextgenapiresource,omitempty"`
	PMTUMin                   int           `json:"pmtumin,omitempty"`
	PMTUTimeout               int           `json:"pmtutimeout,omitempty"`
	ProxyProtocol             string        `json:"proxyprotocol,omitempty"`
	SecureCookie              string        `json:"securecookie,omitempty"`
	SecureICAPorts            []interface{} `json:"secureicaports,omitempty"`
	ServicePathIngressVlan    int           `json:"servicepathingressvlan,omitempty"`
	TCPCIP                    string        `json:"tcpcip,omitempty"`
	TimeZone                  string        `json:"timezone,omitempty"`
	UseProxyPort              string        `json:"useproxyport,omitempty"`
}

type NSExtensionExtensionFunctionBinding struct {
	ActiveExtensionFunction         int      `json:"activeextensionfunction,omitempty"`
	ExtensionFuncDescription        string   `json:"extensionfuncdescription,omitempty"`
	ExtensionFunctionAllParams      []string `json:"extensionfunctionallparams,omitempty"`
	ExtensionFunctionAllParamsCount int      `json:"extensionfunctionallparamscount,omitempty"`
	ExtensionFunctionArgCount       int      `json:"extensionfunctionargcount,omitempty"`
	ExtensionFunctionArgType        []string `json:"extensionfunctionargtype,omitempty"`
	ExtensionFunctionClasses        []string `json:"extensionfunctionclasses,omitempty"`
	ExtensionFunctionClassesCount   int      `json:"extensionfunctionclassescount,omitempty"`
	ExtensionFunctionClassType      string   `json:"extensionfunctionclasstype,omitempty"`
	ExtensionFunctionLineNumber     int      `json:"extensionfunctionlinenumber,omitempty"`
	ExtensionFunctionName           string   `json:"extensionfunctionname,omitempty"`
	ExtensionFunctionReturnType     string   `json:"extensionfunctionreturntype,omitempty"`
	Name                            string   `json:"name,omitempty"`
}

type NSVersion struct {
	InstalledVersion   bool   `json:"installedversion,omitempty"`
	Mode               int    `json:"mode,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	Version            string `json:"version,omitempty"`
}

type NSHTTPParam struct {
	Builtin                   []string `json:"builtin,omitempty"`
	ConMultiplex              string   `json:"conmultiplex,omitempty"`
	DropInvalReqs             string   `json:"dropinvalreqs,omitempty"`
	Feature                   string   `json:"feature,omitempty"`
	HTTP2ServerSide           string   `json:"http2serverside,omitempty"`
	IgnoreConnectCodingScheme string   `json:"ignoreconnectcodingscheme,omitempty"`
	InsNSSrvrHdr              string   `json:"insnssrvrhdr,omitempty"`
	LogErrResp                string   `json:"logerrresp,omitempty"`
	MarkConnReqInval          string   `json:"markconnreqinval,omitempty"`
	MarkHTTP09Inval           string   `json:"markhttp09inval,omitempty"`
	MaxReusePool              int      `json:"maxreusepool,omitempty"`
	NextGenAPIResource        string   `json:"_nextgenapiresource,omitempty"`
	NSSrvrHdr                 string   `json:"nssrvrhdr,omitempty"`
}

type NSIP6 struct {
	AdvertiseOnDefaultPartition string   `json:"advertiseondefaultpartition,omitempty"`
	Count                       float64  `json:"__count,omitempty"`
	CurState                    string   `json:"curstate,omitempty"`
	DecrementHopLimit           string   `json:"decrementhoplimit,omitempty"`
	DynamicRouting              string   `json:"dynamicrouting,omitempty"`
	FTP                         string   `json:"ftp,omitempty"`
	GUI                         string   `json:"gui,omitempty"`
	HostRoute                   string   `json:"hostroute,omitempty"`
	ICMP                        string   `json:"icmp,omitempty"`
	ICMPResponse                string   `json:"icmpresponse,omitempty"`
	IP6HostRtGw                 string   `json:"ip6hostrtgw,omitempty"`
	IPType                      []string `json:"iptype,omitempty"`
	IPv6Address                 string   `json:"ipv6address,omitempty"`
	MapField                    string   `json:"map,omitempty"`
	Metric                      int      `json:"metric,omitempty"`
	MgmtAccess                  string   `json:"mgmtaccess,omitempty"`
	MPTCPAdvertise              string   `json:"mptcpadvertise,omitempty"`
	ND                          string   `json:"nd,omitempty"`
	NDOwner                     int      `json:"ndowner,omitempty"`
	NetworkRoute                string   `json:"networkroute,omitempty"`
	NextGenAPIResource          string   `json:"_nextgenapiresource,omitempty"`
	OperationalNDOwner          int      `json:"operationalndowner,omitempty"`
	OSPF6LSAType                string   `json:"ospf6lsatype,omitempty"`
	OSPFArea                    int      `json:"ospfarea,omitempty"`
	OwnerDownResponse           string   `json:"ownerdownresponse,omitempty"`
	OwnerNode                   int      `json:"ownernode,omitempty"`
	RestrictAccess              string   `json:"restrictaccess,omitempty"`
	Scope                       string   `json:"scope,omitempty"`
	SNMP                        string   `json:"snmp,omitempty"`
	SSH                         string   `json:"ssh,omitempty"`
	State                       string   `json:"state,omitempty"`
	SystemType                  string   `json:"systemtype,omitempty"`
	Tag                         int      `json:"tag,omitempty"`
	TD                          int      `json:"td,omitempty"`
	Telnet                      string   `json:"telnet,omitempty"`
	TypeField                   string   `json:"type,omitempty"`
	VIPRtAdv2BSD                bool     `json:"viprtadv2bsd,omitempty"`
	VIPVSerCount                int      `json:"vipvsercount,omitempty"`
	VIPVSerDownCount            int      `json:"vipvserdowncount,omitempty"`
	VLAN                        int      `json:"vlan,omitempty"`
	VRID6                       int      `json:"vrid6,omitempty"`
	VServer                     string   `json:"vserver,omitempty"`
	VServerRHILevel             string   `json:"vserverrhilevel,omitempty"`
}

type NSPartitionVXLANBinding struct {
	PartitionName string `json:"partitionname,omitempty"`
	VXLAN         int    `json:"vxlan,omitempty"`
}

type NSIP struct {
	AdvertiseOnDefaultPartition string   `json:"advertiseondefaultpartition,omitempty"`
	ARP                         string   `json:"arp,omitempty"`
	ARPOwner                    int      `json:"arpowner,omitempty"`
	ARPResponse                 string   `json:"arpresponse,omitempty"`
	BGP                         string   `json:"bgp,omitempty"`
	Count                       float64  `json:"__count,omitempty"`
	DecrementTTL                string   `json:"decrementttl,omitempty"`
	DynamicRouting              string   `json:"dynamicrouting,omitempty"`
	Flags                       int      `json:"flags,omitempty"`
	FreePorts                   int      `json:"freeports,omitempty"`
	FTP                         string   `json:"ftp,omitempty"`
	GUI                         string   `json:"gui,omitempty"`
	HostRoute                   string   `json:"hostroute,omitempty"`
	HostRtGw                    string   `json:"hostrtgw,omitempty"`
	HostRtGwAct                 string   `json:"hostrtgwact,omitempty"`
	ICMP                        string   `json:"icmp,omitempty"`
	ICMPResponse                string   `json:"icmpresponse,omitempty"`
	IPAddress                   string   `json:"ipaddress,omitempty"`
	IPType                      []string `json:"iptype,omitempty"`
	Metric                      int      `json:"metric,omitempty"`
	MgmtAccess                  string   `json:"mgmtaccess,omitempty"`
	MPTCPAdvertise              string   `json:"mptcpadvertise,omitempty"`
	Netmask                     string   `json:"netmask,omitempty"`
	NetworkRoute                string   `json:"networkroute,omitempty"`
	NextGenAPIResource          string   `json:"_nextgenapiresource,omitempty"`
	OperationalARPOwner         int      `json:"operationalarpowner,omitempty"`
	OSPF                        string   `json:"ospf,omitempty"`
	OSPFArea                    int      `json:"ospfarea,omitempty"`
	OSPFAreaVal                 int      `json:"ospfareaval,omitempty"`
	OSPFLSAType                 string   `json:"ospflsatype,omitempty"`
	OwnerDownResponse           string   `json:"ownerdownresponse,omitempty"`
	OwnerNode                   int      `json:"ownernode,omitempty"`
	RestrictAccess              string   `json:"restrictaccess,omitempty"`
	RIP                         string   `json:"rip,omitempty"`
	SNMP                        string   `json:"snmp,omitempty"`
	SSH                         string   `json:"ssh,omitempty"`
	State                       string   `json:"state,omitempty"`
	Tag                         int      `json:"tag,omitempty"`
	TD                          int      `json:"td,omitempty"`
	Telnet                      string   `json:"telnet,omitempty"`
	TypeField                   string   `json:"type,omitempty"`
	VIPRtAdv2BSD                bool     `json:"viprtadv2bsd,omitempty"`
	VIPVSerCount                int      `json:"vipvsercount,omitempty"`
	VIPVSerDownCount            int      `json:"vipvserdowncount,omitempty"`
	VIPVSrvrRHIActiveCount      int      `json:"vipvsrvrrhiactivecount,omitempty"`
	VIPVSrvrRHIActiveUpCount    int      `json:"vipvsrvrrhiactiveupcount,omitempty"`
	VRID                        int      `json:"vrid,omitempty"`
	VServer                     string   `json:"vserver,omitempty"`
	VServerRHILevel             string   `json:"vserverrhilevel,omitempty"`
}

type NSExtension struct {
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Detail             string  `json:"detail,omitempty"`
	FunctionHaltCount  int     `json:"functionhaltcount,omitempty"`
	FunctionHits       int     `json:"functionhits,omitempty"`
	FunctionUndefHits  int     `json:"functionundefhits,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Overwrite          bool    `json:"overwrite,omitempty"`
	Src                string  `json:"src,omitempty"`
	Trace              string  `json:"trace,omitempty"`
	TraceFunctions     string  `json:"tracefunctions,omitempty"`
	TraceVariables     string  `json:"tracevariables,omitempty"`
	TypeField          string  `json:"type,omitempty"`
}

type Shutdown struct {
}

type NSCapacity struct {
	ActualBandwidth    int    `json:"actualbandwidth,omitempty"`
	Bandwidth          int    `json:"bandwidth,omitempty"`
	DaysToExpiration   int    `json:"daystoexpiration,omitempty"`
	Edition            string `json:"edition,omitempty"`
	InstanceCount      int    `json:"instancecount,omitempty"`
	MaxBandwidth       int    `json:"maxbandwidth,omitempty"`
	MaxVCPUCount       int    `json:"maxvcpucount,omitempty"`
	MinBandwidth       int    `json:"minbandwidth,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	NodeID             int    `json:"nodeid,omitempty"`
	Password           string `json:"password,omitempty"`
	Platform           string `json:"platform,omitempty"`
	Unit               string `json:"unit,omitempty"`
	Username           string `json:"username,omitempty"`
	VCPU               bool   `json:"vcpu,omitempty"`
	VCPUCount          int    `json:"vcpucount,omitempty"`
}

type NSSavedConfig struct {
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	TextBlob           string `json:"textblob,omitempty"`
}

type NSRPCNode struct {
	Count              float64 `json:"__count,omitempty"`
	IPAddress          string  `json:"ipaddress,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Password           string  `json:"password,omitempty"`
	Secure             string  `json:"secure,omitempty"`
	SrcIP              string  `json:"srcip,omitempty"`
	ValidateCert       string  `json:"validatecert,omitempty"`
}

type NSAssignment struct {
	Add                string  `json:"Add,omitempty"`
	Append             string  `json:"append,omitempty"`
	Clear              bool    `json:"clear,omitempty"`
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Hits               int     `json:"hits,omitempty"`
	Name               string  `json:"name,omitempty"`
	NewName            string  `json:"newname,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	ReferenceCount     int     `json:"referencecount,omitempty"`
	Set                string  `json:"set,omitempty"`
	Sub                string  `json:"sub,omitempty"`
	UndefHits          int     `json:"undefhits,omitempty"`
	Variable           string  `json:"variable,omitempty"`
}

type NSTCPBufParam struct {
	Builtin            []string `json:"builtin,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	MemLimit           int      `json:"memlimit,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	Size               int      `json:"size,omitempty"`
}

type NSRateControl struct {
	ICMPThreshold      int    `json:"icmpthreshold,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	TCPRstThreshold    int    `json:"tcprstthreshold,omitempty"`
	TCPThreshold       int    `json:"tcpthreshold,omitempty"`
	UDPThreshold       int    `json:"udpthreshold,omitempty"`
}

type NSLimitSessions struct {
	Count              float64       `json:"__count,omitempty"`
	Detail             bool          `json:"detail,omitempty"`
	Drop               int           `json:"drop,omitempty"`
	Flag               int           `json:"flag,omitempty"`
	Flags              int           `json:"flags,omitempty"`
	Hits               int           `json:"hits,omitempty"`
	LimitIdentifier    string        `json:"limitidentifier,omitempty"`
	MaxBandwidth       int           `json:"maxbandwidth,omitempty"`
	Name               string        `json:"name,omitempty"`
	NextGenAPIResource string        `json:"_nextgenapiresource,omitempty"`
	Number             []interface{} `json:"number,omitempty"`
	ReferenceCount     int           `json:"referencecount,omitempty"`
	SelectorIPv61      string        `json:"selectoripv61,omitempty"`
	SelectorIPv62      string        `json:"selectoripv62,omitempty"`
	Timeout            int           `json:"timeout,omitempty"`
	Unit               int           `json:"unit,omitempty"`
}

type NSTrafficDomainVLANBinding struct {
	TD   int `json:"td,omitempty"`
	VLAN int `json:"vlan,omitempty"`
}

type NSSPParams struct {
	BaseThreshold      int           `json:"basethreshold,omitempty"`
	Builtin            []string      `json:"builtin,omitempty"`
	Feature            string        `json:"feature,omitempty"`
	NextGenAPIResource string        `json:"_nextgenapiresource,omitempty"`
	Table0             []interface{} `json:"table0,omitempty"`
	Throttle           string        `json:"throttle,omitempty"`
}

type NSACLs6 struct {
	TypeField string `json:"type,omitempty"`
}

type NSRunningConfig struct {
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	Response           string `json:"response,omitempty"`
	WithDefaults       bool   `json:"withdefaults,omitempty"`
}

type Reboot struct {
	Warm bool `json:"warm,omitempty"`
}

type NSTimer struct {
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Interval           int     `json:"interval,omitempty"`
	Name               string  `json:"name,omitempty"`
	NewName            string  `json:"newname,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Unit               string  `json:"unit,omitempty"`
}

type NSCentralManagementServer struct {
	ActivationCode             string  `json:"activationcode,omitempty"`
	ADCPassword                string  `json:"adcpassword,omitempty"`
	ADCUsername                string  `json:"adcusername,omitempty"`
	ADMServiceConnectionStatus string  `json:"admserviceconnectionstatus,omitempty"`
	ADMServiceEnvironment      string  `json:"admserviceenvironment,omitempty"`
	Count                      float64 `json:"__count,omitempty"`
	CustomerID                 string  `json:"customerid,omitempty"`
	DeviceProfileName          string  `json:"deviceprofilename,omitempty"`
	InstanceID                 string  `json:"instanceid,omitempty"`
	IPAddress                  string  `json:"ipaddress,omitempty"`
	NextGenAPIResource         string  `json:"_nextgenapiresource,omitempty"`
	Password                   string  `json:"password,omitempty"`
	ServerName                 string  `json:"servername,omitempty"`
	TypeField                  string  `json:"type,omitempty"`
	Username                   string  `json:"username,omitempty"`
	ValidateCert               string  `json:"validatecert,omitempty"`
}

type NSEncryptionParams struct {
	KeyValue           string `json:"keyvalue,omitempty"`
	Method             string `json:"method,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
}

type NSSourceRouteCacheTable struct {
	Count              float64 `json:"__count,omitempty"`
	InterfaceField     string  `json:"Interface,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	SourceIP           string  `json:"sourceip,omitempty"`
	SourceMAC          string  `json:"sourcemac,omitempty"`
	VLAN               int     `json:"vlan,omitempty"`
}

type NSTimeZone struct {
	Count              float64 `json:"__count,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Value              string  `json:"value,omitempty"`
}

type NSACL6 struct {
	ACL6Action         string   `json:"acl6action,omitempty"`
	ACL6Name           string   `json:"acl6name,omitempty"`
	ACLAction          string   `json:"aclaction,omitempty"`
	ACLAssociate       []string `json:"aclassociate,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	DestIPOp           string   `json:"destipop,omitempty"`
	DestIPv6           bool     `json:"destipv6,omitempty"`
	DestIPv6Val        string   `json:"destipv6val,omitempty"`
	DestPort           bool     `json:"destport,omitempty"`
	DestPortOp         string   `json:"destportop,omitempty"`
	DestPortVal        string   `json:"destportval,omitempty"`
	DFDHash            string   `json:"dfdhash,omitempty"`
	DFDPrefix          int      `json:"dfdprefix,omitempty"`
	Established        bool     `json:"established,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	ICMPCode           int      `json:"icmpcode,omitempty"`
	ICMPType           int      `json:"icmptype,omitempty"`
	InterfaceField     string   `json:"Interface,omitempty"`
	KernelState        string   `json:"kernelstate,omitempty"`
	LogState           string   `json:"logstate,omitempty"`
	NewName            string   `json:"newname,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	NodeID             int      `json:"nodeid,omitempty"`
	Priority           int      `json:"priority,omitempty"`
	Protocol           string   `json:"protocol,omitempty"`
	ProtocolNumber     int      `json:"protocolnumber,omitempty"`
	RateLimit          int      `json:"ratelimit,omitempty"`
	SrcIPOp            string   `json:"srcipop,omitempty"`
	SrcIPv6            bool     `json:"srcipv6,omitempty"`
	SrcIPv6Val         string   `json:"srcipv6val,omitempty"`
	SrcMAC             string   `json:"srcmac,omitempty"`
	SrcMACMask         string   `json:"srcmacmask,omitempty"`
	SrcPort            bool     `json:"srcport,omitempty"`
	SrcPortOp          string   `json:"srcportop,omitempty"`
	SrcPortVal         string   `json:"srcportval,omitempty"`
	State              string   `json:"state,omitempty"`
	Stateful           string   `json:"stateful,omitempty"`
	TD                 int      `json:"td,omitempty"`
	TTL                int      `json:"ttl,omitempty"`
	TypeField          string   `json:"type,omitempty"`
	VLAN               int      `json:"vlan,omitempty"`
	VXLAN              int      `json:"vxlan,omitempty"`
}

type NSAppFlowParam struct {
	ClientTrafficOnly  string `json:"clienttrafficonly,omitempty"`
	HTTPCookie         string `json:"httpcookie,omitempty"`
	HTTPHost           string `json:"httphost,omitempty"`
	HTTPMethod         string `json:"httpmethod,omitempty"`
	HTTPReferer        string `json:"httpreferer,omitempty"`
	HTTPURL            string `json:"httpurl,omitempty"`
	HTTPUserAgent      string `json:"httpuseragent,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	TemplateRefresh    int    `json:"templaterefresh,omitempty"`
	UDPPMTU            int    `json:"udppmtu,omitempty"`
}

type NSVariable struct {
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Expires            int     `json:"expires,omitempty"`
	IfFull             string  `json:"iffull,omitempty"`
	IfNoValue          string  `json:"ifnovalue,omitempty"`
	IfValueTooBig      string  `json:"ifvaluetoobig,omitempty"`
	Init               string  `json:"init,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	ReferenceCount     int     `json:"referencecount,omitempty"`
	Scope              string  `json:"scope,omitempty"`
	TypeField          string  `json:"type,omitempty"`
}

type NSConsoleLoginPrompt struct {
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	PromptString       string `json:"promptstring,omitempty"`
}

type NSHostname struct {
	Count              float64 `json:"__count,omitempty"`
	Hostname           string  `json:"hostname,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	OwnerNode          int     `json:"ownernode,omitempty"`
}

type NSJob struct {
	Count              float64 `json:"__count,omitempty"`
	ErrorCode          int     `json:"errorcode,omitempty"`
	ID                 int     `json:"id,omitempty"`
	Message            string  `json:"message,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Progress           string  `json:"progress,omitempty"`
	Response           string  `json:"response,omitempty"`
	Status             string  `json:"status,omitempty"`
	TimeElapsed        int     `json:"timeelapsed,omitempty"`
}

type NSAppFlowCollector struct {
	Count              float64 `json:"__count,omitempty"`
	IPAddress          string  `json:"ipaddress,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Port               int     `json:"port,omitempty"`
}

type NSLimitSelector struct {
	Count              float64  `json:"__count,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	Rule               []string `json:"rule,omitempty"`
	SelectorName       string   `json:"selectorname,omitempty"`
}

type NSTimeout struct {
	AnyClient          int    `json:"anyclient,omitempty"`
	AnyServer          int    `json:"anyserver,omitempty"`
	AnyTCPClient       int    `json:"anytcpclient,omitempty"`
	AnyTCPServer       int    `json:"anytcpserver,omitempty"`
	Client             int    `json:"client,omitempty"`
	HalfClose          int    `json:"halfclose,omitempty"`
	HTTPClient         int    `json:"httpclient,omitempty"`
	HTTPServer         int    `json:"httpserver,omitempty"`
	NewConnIdleTimeout int    `json:"newconnidletimeout,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	NonTCPZombie       int    `json:"nontcpzombie,omitempty"`
	ReducedFinTimeout  int    `json:"reducedfintimeout,omitempty"`
	ReducedRstTimeout  int    `json:"reducedrsttimeout,omitempty"`
	Server             int    `json:"server,omitempty"`
	TCPClient          int    `json:"tcpclient,omitempty"`
	TCPServer          int    `json:"tcpserver,omitempty"`
	Zombie             int    `json:"zombie,omitempty"`
}

type NSTCPParam struct {
	AckOnPush                           string   `json:"ackonpush,omitempty"`
	AutoSynCookieTimeout                int      `json:"autosyncookietimeout,omitempty"`
	Builtin                             []string `json:"builtin,omitempty"`
	CompactTCPOptionNoOp                string   `json:"compacttcpoptionnoop,omitempty"`
	ConnFlushIfNoMem                    string   `json:"connflushifnomem,omitempty"`
	ConnFlushThres                      int      `json:"connflushthres,omitempty"`
	DelayedAck                          int      `json:"delayedack,omitempty"`
	DelinkClientServerOnRst             string   `json:"delinkclientserveronrst,omitempty"`
	DownStateRst                        string   `json:"downstaterst,omitempty"`
	EnhancedISNGeneration               string   `json:"enhancedisngeneration,omitempty"`
	Feature                             string   `json:"feature,omitempty"`
	InitialCwnd                         int      `json:"initialcwnd,omitempty"`
	KAProbeUpdateLastActivity           string   `json:"kaprobeupdatelastactivity,omitempty"`
	LearnVSvrMSS                        string   `json:"learnvsvrmss,omitempty"`
	LimitedPersist                      string   `json:"limitedpersist,omitempty"`
	MaxBurst                            int      `json:"maxburst,omitempty"`
	MaxDynServerProbes                  int      `json:"maxdynserverprobes,omitempty"`
	MaxPktPerMSS                        int      `json:"maxpktpermss,omitempty"`
	MaxSynAckRetx                       int      `json:"maxsynackretx,omitempty"`
	MaxSynHold                          int      `json:"maxsynhold,omitempty"`
	MaxSynHoldPerProbe                  int      `json:"maxsynholdperprobe,omitempty"`
	MaxTimeWaitConn                     int      `json:"maxtimewaitconn,omitempty"`
	MinRTO                              int      `json:"minrto,omitempty"`
	MPTCPChecksum                       string   `json:"mptcpchecksum,omitempty"`
	MPTCPCloseMPTCPSessionOnLastSFClose string   `json:"mptcpclosemptcpsessiononlastsfclose,omitempty"`
	MPTCPConCloseOnPassiveSF            string   `json:"mptcpconcloseonpassivesf,omitempty"`
	MPTCPFastCloseOption                string   `json:"mptcpfastcloseoption,omitempty"`
	MPTCPImmediateSFCloseOnFin          string   `json:"mptcpimmediatesfcloseonfin,omitempty"`
	MPTCPMaxPendingSF                   int      `json:"mptcpmaxpendingsf,omitempty"`
	MPTCPMaxSF                          int      `json:"mptcpmaxsf,omitempty"`
	MPTCPPendingJoinThreshold           int      `json:"mptcppendingjointhreshold,omitempty"`
	MPTCPReliableAddAddr                string   `json:"mptcpreliableaddaddr,omitempty"`
	MPTCPRTOsToSwitchSF                 int      `json:"mptcprtostoswitchsf,omitempty"`
	MPTCPSendSFResetOption              string   `json:"mptcpsendsfresetoption,omitempty"`
	MPTCPSFReplaceTimeout               int      `json:"mptcpsfreplacetimeout,omitempty"`
	MPTCPSFTimeout                      int      `json:"mptcpsftimeout,omitempty"`
	MPTCPUseBackupOnDSS                 string   `json:"mptcpusebackupondss,omitempty"`
	MSSLearnDelay                       int      `json:"msslearndelay,omitempty"`
	MSSLearnInterval                    int      `json:"msslearninterval,omitempty"`
	Nagle                               string   `json:"nagle,omitempty"`
	NextGenAPIResource                  string   `json:"_nextgenapiresource,omitempty"`
	OOOQSize                            int      `json:"oooqsize,omitempty"`
	PktPerRetx                          int      `json:"pktperretx,omitempty"`
	RecvBuffSize                        int      `json:"recvbuffsize,omitempty"`
	RFC5961ChlgAckLimit                 int      `json:"rfc5961chlgacklimit,omitempty"`
	Sack                                string   `json:"sack,omitempty"`
	SendResetReasonCode                 string   `json:"sendresetreasoncode,omitempty"`
	SlowStartIncr                       int      `json:"slowstartincr,omitempty"`
	SynAttackDetection                  string   `json:"synattackdetection,omitempty"`
	SynHoldFastGiveUp                   int      `json:"synholdfastgiveup,omitempty"`
	TCPFastOpenCookieTimeout            int      `json:"tcpfastopencookietimeout,omitempty"`
	TCPFinTimeout                       int      `json:"tcpfintimeout,omitempty"`
	TCPMaxRetries                       int      `json:"tcpmaxretries,omitempty"`
	Ws                                  string   `json:"ws,omitempty"`
	WSVal                               int      `json:"wsval,omitempty"`
}

type NSServiceFunction struct {
	Count               float64 `json:"__count,omitempty"`
	IngressVLAN         int     `json:"ingressvlan,omitempty"`
	NextGenAPIResource  string  `json:"_nextgenapiresource,omitempty"`
	ServiceFunctionName string  `json:"servicefunctionname,omitempty"`
}

type NSChannelParam struct {
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	VFAutoRecover      string `json:"vfautorecover,omitempty"`
}

type NSEncryptionKey struct {
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	IV                 string  `json:"iv,omitempty"`
	KeyValue           string  `json:"keyvalue,omitempty"`
	Method             string  `json:"method,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Padding            string  `json:"padding,omitempty"`
}

type NSKeyManagerProxy struct {
	Count              float64 `json:"__count,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NodeID             int     `json:"nodeid,omitempty"`
	Port               int     `json:"port,omitempty"`
	ServerIP           string  `json:"serverip,omitempty"`
	ServerName         string  `json:"servername,omitempty"`
	Status             int     `json:"status,omitempty"`
}

type NSICAPProfile struct {
	Allow204            string  `json:"allow204,omitempty"`
	ConnectionKeepAlive string  `json:"connectionkeepalive,omitempty"`
	Count               float64 `json:"__count,omitempty"`
	HostHeader          string  `json:"hostheader,omitempty"`
	InsertHTTPRequest   string  `json:"inserthttprequest,omitempty"`
	InsertICAPHeaders   string  `json:"inserticapheaders,omitempty"`
	InspectHTTP2        string  `json:"inspecthttp2,omitempty"`
	LogAction           string  `json:"logaction,omitempty"`
	Mode                string  `json:"mode,omitempty"`
	Name                string  `json:"name,omitempty"`
	NextGenAPIResource  string  `json:"_nextgenapiresource,omitempty"`
	Preview             string  `json:"preview,omitempty"`
	PreviewLength       int     `json:"previewlength,omitempty"`
	QueryParams         string  `json:"queryparams,omitempty"`
	ReqTimeout          int     `json:"reqtimeout,omitempty"`
	ReqTimeoutAction    string  `json:"reqtimeoutaction,omitempty"`
	URI                 string  `json:"uri,omitempty"`
	UserAgent           string  `json:"useragent,omitempty"`
}

type NSSimpleACL struct {
	ACLAction          string  `json:"aclaction,omitempty"`
	ACLName            string  `json:"aclname,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	DestPort           int     `json:"destport,omitempty"`
	EstSessions        bool    `json:"estsessions,omitempty"`
	Hits               int     `json:"hits,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Protocol           string  `json:"protocol,omitempty"`
	SrcIP              string  `json:"srcip,omitempty"`
	TD                 int     `json:"td,omitempty"`
	TTL                int     `json:"ttl,omitempty"`
}

type NSVariableValues struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NodeID             int     `json:"nodeid,omitempty"`
	VariableData       string  `json:"variabledata,omitempty"`
	VariableKey        string  `json:"variablekey,omitempty"`
	VariableValue      string  `json:"variablevalue,omitempty"`
}

type NSPBRs struct {
}

type NSWebLogParam struct {
	BufferSizeMB       int      `json:"buffersizemb,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	CustomReqHdrs      []string `json:"customreqhdrs,omitempty"`
	CustomRspHdrs      []string `json:"customrsphdrs,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
}

type NSLicenseActivationData struct {
	Filename           string `json:"filename,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
}

type NSEvents struct {
	Count              float64 `json:"__count,omitempty"`
	Data0              int     `json:"data0,omitempty"`
	Data1              int     `json:"data1,omitempty"`
	Data2              int     `json:"data2,omitempty"`
	Data3              int     `json:"data3,omitempty"`
	DevID              int     `json:"devid,omitempty"`
	DevName            string  `json:"devname,omitempty"`
	EventCode          int     `json:"eventcode,omitempty"`
	EventNo            int     `json:"eventno,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Text               string  `json:"text,omitempty"`
	Time               int     `json:"time,omitempty"`
}

type NSMigration struct {
	Count                           float64 `json:"__count,omitempty"`
	DestIP                          string  `json:"destip,omitempty"`
	DestPort                        int     `json:"destport,omitempty"`
	DumpSession                     string  `json:"dumpsession,omitempty"`
	MigDFDSessionsActive            int     `json:"migdfdsessionsactive,omitempty"`
	MigDFDSessionsActiveRollback    int     `json:"migdfdsessionsactiverollback,omitempty"`
	MigDFDSessionsAllocated         int     `json:"migdfdsessionsallocated,omitempty"`
	MigDFDSessionsAllocatedRollback int     `json:"migdfdsessionsallocatedrollback,omitempty"`
	MigHAStateFlag                  int     `json:"mighastateflag,omitempty"`
	MigL4SessionsActive             int     `json:"migl4sessionsactive,omitempty"`
	MigL4SessionsActiveRollback     int     `json:"migl4sessionsactiverollback,omitempty"`
	MigL4SessionsAllocated          int     `json:"migl4sessionsallocated,omitempty"`
	MigL4SessionsAllocatedRollback  int     `json:"migl4sessionsallocatedrollback,omitempty"`
	MigrationEndTime                string  `json:"migrationendtime,omitempty"`
	MigrationRollbackStartTime      string  `json:"migrationrollbackstarttime,omitempty"`
	MigrationStartTime              string  `json:"migrationstarttime,omitempty"`
	MigrationStatus                 string  `json:"migrationstatus,omitempty"`
	NextGenAPIResource              string  `json:"_nextgenapiresource,omitempty"`
	SrcIP                           string  `json:"srcip,omitempty"`
	SrcPort                         int     `json:"srcport,omitempty"`
	Timeout                         int     `json:"timeout,omitempty"`
}

type NSHMACKey struct {
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Digest             string  `json:"digest,omitempty"`
	KeyValue           string  `json:"keyvalue,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
}

type NSPBR6 struct {
	Action             string  `json:"action,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	CurState           int     `json:"curstate,omitempty"`
	Data               bool    `json:"data,omitempty"`
	DestIPOp           string  `json:"destipop,omitempty"`
	DestIPv6           bool    `json:"destipv6,omitempty"`
	DestIPv6Val        string  `json:"destipv6val,omitempty"`
	DestPort           bool    `json:"destport,omitempty"`
	DestPortOp         string  `json:"destportop,omitempty"`
	DestPortVal        string  `json:"destportval,omitempty"`
	Detail             bool    `json:"detail,omitempty"`
	FailedProbes       int     `json:"failedprobes,omitempty"`
	Hits               int     `json:"hits,omitempty"`
	Interface          string  `json:"Interface,omitempty"`
	IPTunnel           string  `json:"iptunnel,omitempty"`
	KernelState        string  `json:"kernelstate,omitempty"`
	Monitor            string  `json:"monitor,omitempty"`
	MonStatCode        int     `json:"monstatcode,omitempty"`
	MonStatParam1      int     `json:"monstatparam1,omitempty"`
	MonStatParam2      int     `json:"monstatparam2,omitempty"`
	MonStatParam3      int     `json:"monstatparam3,omitempty"`
	MSR                string  `json:"msr,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NextHop            bool    `json:"nexthop,omitempty"`
	NextHopVal         string  `json:"nexthopval,omitempty"`
	NextHopVLAN        int     `json:"nexthopvlan,omitempty"`
	OwnerGroup         string  `json:"ownergroup,omitempty"`
	Priority           int     `json:"priority,omitempty"`
	Protocol           string  `json:"protocol,omitempty"`
	ProtocolNumber     int     `json:"protocolnumber,omitempty"`
	SrcIPOp            string  `json:"srcipop,omitempty"`
	SrcIPv6            bool    `json:"srcipv6,omitempty"`
	SrcIPv6Val         string  `json:"srcipv6val,omitempty"`
	SrcMAC             string  `json:"srcmac,omitempty"`
	SrcMACMask         string  `json:"srcmacmask,omitempty"`
	SrcPort            bool    `json:"srcport,omitempty"`
	SrcPortOp          string  `json:"srcportop,omitempty"`
	SrcPortVal         string  `json:"srcportval,omitempty"`
	State              string  `json:"state,omitempty"`
	TD                 int     `json:"td,omitempty"`
	TotalFailedProbes  int     `json:"totalfailedprobes,omitempty"`
	TotalProbes        int     `json:"totalprobes,omitempty"`
	VLAN               int     `json:"vlan,omitempty"`
	VXLAN              int     `json:"vxlan,omitempty"`
	VXLANVLANMap       string  `json:"vxlanvlanmap,omitempty"`
}

type NSACLs struct {
	TypeField string `json:"type,omitempty"`
}

type NSDHCPParams struct {
	DHCPClient         string `json:"dhcpclient,omitempty"`
	HostRtGw           string `json:"hostrtgw,omitempty"`
	IPAddress          string `json:"ipaddress,omitempty"`
	Netmask            string `json:"netmask,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	Running            bool   `json:"running,omitempty"`
	SaveRoute          string `json:"saveroute,omitempty"`
}

type NSSurgeQ struct {
	Name       string `json:"name,omitempty"`
	Port       int    `json:"port,omitempty"`
	ServerName string `json:"servername,omitempty"`
}

type NSNextGenAPI struct {
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	State              string `json:"state,omitempty"`
}

type NSXMLNamespace struct {
	Count              float64 `json:"__count,omitempty"`
	Description        string  `json:"description,omitempty"`
	Namespace          string  `json:"Namespace,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Prefix             string  `json:"prefix,omitempty"`
}

type NSACL struct {
	ACLAction          string   `json:"aclaction,omitempty"`
	ACLAssociate       []string `json:"aclassociate,omitempty"`
	ACLChildCount      int      `json:"aclchildcount,omitempty"`
	ACLName            string   `json:"aclname,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	DestIP             bool     `json:"destip,omitempty"`
	DestIPDataset      string   `json:"destipdataset,omitempty"`
	DestIPOp           string   `json:"destipop,omitempty"`
	DestIPVal          string   `json:"destipval,omitempty"`
	DestPort           bool     `json:"destport,omitempty"`
	DestPortDataset    string   `json:"destportdataset,omitempty"`
	DestPortOp         string   `json:"destportop,omitempty"`
	DestPortVal        string   `json:"destportval,omitempty"`
	DFDHash            string   `json:"dfdhash,omitempty"`
	Established        bool     `json:"established,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	ICMPCode           int      `json:"icmpcode,omitempty"`
	ICMPType           int      `json:"icmptype,omitempty"`
	Interface          string   `json:"Interface,omitempty"`
	KernelState        string   `json:"kernelstate,omitempty"`
	LogState           string   `json:"logstate,omitempty"`
	NewName            string   `json:"newname,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	NodeID             int      `json:"nodeid,omitempty"`
	Priority           int      `json:"priority,omitempty"`
	Protocol           string   `json:"protocol,omitempty"`
	ProtocolNumber     int      `json:"protocolnumber,omitempty"`
	RateLimit          int      `json:"ratelimit,omitempty"`
	SrcIP              bool     `json:"srcip,omitempty"`
	SrcIPDataset       string   `json:"srcipdataset,omitempty"`
	SrcIPOp            string   `json:"srcipop,omitempty"`
	SrcIPVal           string   `json:"srcipval,omitempty"`
	SrcMAC             string   `json:"srcmac,omitempty"`
	SrcMACMask         string   `json:"srcmacmask,omitempty"`
	SrcPort            bool     `json:"srcport,omitempty"`
	SrcPortDataset     string   `json:"srcportdataset,omitempty"`
	SrcPortOp          string   `json:"srcportop,omitempty"`
	SrcPortVal         string   `json:"srcportval,omitempty"`
	State              string   `json:"state,omitempty"`
	Stateful           string   `json:"stateful,omitempty"`
	TD                 int      `json:"td,omitempty"`
	TTL                int      `json:"ttl,omitempty"`
	TypeField          string   `json:"type,omitempty"`
	VLAN               int      `json:"vlan,omitempty"`
	VXLAN              int      `json:"vxlan,omitempty"`
}

type NSLicenseServer struct {
	Count              float64 `json:"__count,omitempty"`
	DeviceProfileName  string  `json:"deviceprofilename,omitempty"`
	ForceUpdateIP      bool    `json:"forceupdateip,omitempty"`
	GPTimeLeft         int     `json:"gptimeleft,omitempty"`
	Grace              int     `json:"grace,omitempty"`
	LicenseMode        string  `json:"licensemode,omitempty"`
	LicenseServerIP    string  `json:"licenseserverip,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NodeID             int     `json:"nodeid,omitempty"`
	Password           string  `json:"password,omitempty"`
	Port               int     `json:"port,omitempty"`
	ServerName         string  `json:"servername,omitempty"`
	Status             int     `json:"status,omitempty"`
	TypeField          string  `json:"type,omitempty"`
	Username           string  `json:"username,omitempty"`
}

type NSMgmtParam struct {
	HTTPDMaxClients    int    `json:"httpdmaxclients,omitempty"`
	HTTPDMaxReqWorkers int    `json:"httpdmaxreqworkers,omitempty"`
	MgmtHTTPPort       int    `json:"mgmthttpport,omitempty"`
	MgmtHTTPSPort      int    `json:"mgmthttpsport,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
}

type NSDHCPIP struct {
}

type NSTCPProfile struct {
	AckAggregation              string   `json:"ackaggregation,omitempty"`
	AckOnPush                   string   `json:"ackonpush,omitempty"`
	ApplyAdaptiveTCP            string   `json:"applyadaptivetcp,omitempty"`
	BufferSize                  int      `json:"buffersize,omitempty"`
	Builtin                     []string `json:"builtin,omitempty"`
	BurstRateControl            string   `json:"burstratecontrol,omitempty"`
	ClientIPTCPOption           string   `json:"clientiptcpoption,omitempty"`
	ClientIPTCPOptionNumber     int      `json:"clientiptcpoptionnumber,omitempty"`
	Count                       float64  `json:"__count,omitempty"`
	DelayedAck                  int      `json:"delayedack,omitempty"`
	DropEstConnOnTimeout        string   `json:"dropestconnontimeout,omitempty"`
	DropHalfClosedConnOnTimeout string   `json:"drophalfclosedconnontimeout,omitempty"`
	DSack                       string   `json:"dsack,omitempty"`
	DupAckThresh                int      `json:"dupackthresh,omitempty"`
	DynamicReceiveBuffering     string   `json:"dynamicreceivebuffering,omitempty"`
	Ecn                         string   `json:"ecn,omitempty"`
	EstablishClientConn         string   `json:"establishclientconn,omitempty"`
	Fack                        string   `json:"fack,omitempty"`
	Feature                     string   `json:"feature,omitempty"`
	Flavor                      string   `json:"flavor,omitempty"`
	FRTO                        string   `json:"frto,omitempty"`
	Hystart                     string   `json:"hystart,omitempty"`
	InitialCwnd                 int      `json:"initialcwnd,omitempty"`
	KA                          string   `json:"ka,omitempty"`
	KAConnIdleTime              int      `json:"kaconnidletime,omitempty"`
	KAMaxProbes                 int      `json:"kamaxprobes,omitempty"`
	KAProbeInterval             int      `json:"kaprobeinterval,omitempty"`
	KAProbeUpdateLastActivity   string   `json:"kaprobeupdatelastactivity,omitempty"`
	MaxBurst                    int      `json:"maxburst,omitempty"`
	MaxCwnd                     int      `json:"maxcwnd,omitempty"`
	MaxPktPerMSS                int      `json:"maxpktpermss,omitempty"`
	MinRTO                      int      `json:"minrto,omitempty"`
	MPCapableCBit               string   `json:"mpcapablecbit,omitempty"`
	MPTCP                       string   `json:"mptcp,omitempty"`
	MPTCPDropDataOnPreEstSF     string   `json:"mptcpdropdataonpreestsf,omitempty"`
	MPTCPFastOpen               string   `json:"mptcpfastopen,omitempty"`
	MPTCPSessionTimeout         int      `json:"mptcpsessiontimeout,omitempty"`
	MSS                         int      `json:"mss,omitempty"`
	Nagle                       string   `json:"nagle,omitempty"`
	Name                        string   `json:"name,omitempty"`
	NextGenAPIResource          string   `json:"_nextgenapiresource,omitempty"`
	OOOQSize                    int      `json:"oooqsize,omitempty"`
	PktPerRetx                  int      `json:"pktperretx,omitempty"`
	RateQMax                    int      `json:"rateqmax,omitempty"`
	RefCnt                      int      `json:"refcnt,omitempty"`
	RFC5961Compliance           string   `json:"rfc5961compliance,omitempty"`
	RstMaxAck                   string   `json:"rstmaxack,omitempty"`
	RstWindowAttenuate          string   `json:"rstwindowattenuate,omitempty"`
	Sack                        string   `json:"sack,omitempty"`
	SendBuffSize                int      `json:"sendbuffsize,omitempty"`
	SendClientPortInTCPOption   string   `json:"sendclientportintcpoption,omitempty"`
	SlowStartIncr               int      `json:"slowstartincr,omitempty"`
	SlowStartThreshold          int      `json:"slowstartthreshold,omitempty"`
	SpoofSynDrop                string   `json:"spoofsyndrop,omitempty"`
	SynCookie                   string   `json:"syncookie,omitempty"`
	TailLossProbe               string   `json:"taillossprobe,omitempty"`
	TCPFastOpen                 string   `json:"tcpfastopen,omitempty"`
	TCPFastOpenCookieSize       int      `json:"tcpfastopencookiesize,omitempty"`
	TCPMode                     string   `json:"tcpmode,omitempty"`
	TCPRate                     int      `json:"tcprate,omitempty"`
	TCPSegOffload               string   `json:"tcpsegoffload,omitempty"`
	Timestamp                   string   `json:"timestamp,omitempty"`
	Ws                          string   `json:"ws,omitempty"`
	WSVal                       int      `json:"wsval,omitempty"`
}

type NSTrafficDomainBinding struct {
	NSTrafficDomainBridgeGroupBinding []interface{} `json:"nstrafficdomain_bridgegroup_binding,omitempty"`
	NSTrafficDomainVlanBinding        []interface{} `json:"nstrafficdomain_vlan_binding,omitempty"`
	NSTrafficDomainVxlanBinding       []interface{} `json:"nstrafficdomain_vxlan_binding,omitempty"`
	TD                                int           `json:"td,omitempty"`
}

type NSTimerAutoScalePolicyBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	SampleSize             int    `json:"samplesize,omitempty"`
	Threshold              int    `json:"threshold,omitempty"`
	VServer                string `json:"vserver,omitempty"`
}

type NSMode struct {
	BridgeBPDUs         bool     `json:"bridgebpdus,omitempty"`
	CKA                 bool     `json:"cka,omitempty"`
	DRAdv               bool     `json:"dradv,omitempty"`
	DRAdv6              bool     `json:"dradv6,omitempty"`
	Edge                bool     `json:"edge,omitempty"`
	FR                  bool     `json:"fr,omitempty"`
	IRAdv               bool     `json:"iradv,omitempty"`
	L2                  bool     `json:"l2,omitempty"`
	L3                  bool     `json:"l3,omitempty"`
	MBF                 bool     `json:"mbf,omitempty"`
	MediaClassification bool     `json:"mediaclassification,omitempty"`
	Mode                []string `json:"mode,omitempty"`
	NextGenAPIResource  string   `json:"_nextgenapiresource,omitempty"`
	PMTUD               bool     `json:"pmtud,omitempty"`
	SingleIP            bool     `json:"single_ip,omitempty"`
	SRAdv               bool     `json:"sradv,omitempty"`
	SRAdv6              bool     `json:"sradv6,omitempty"`
	TCPB                bool     `json:"tcpb,omitempty"`
	ULFD                bool     `json:"ulfd,omitempty"`
	USIP                bool     `json:"usip,omitempty"`
	USNIP               bool     `json:"usnip,omitempty"`
}

type NSSimpleACL6 struct {
	ACLAction          string  `json:"aclaction,omitempty"`
	ACLName            string  `json:"aclname,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	DestPort           int     `json:"destport,omitempty"`
	EstSessions        bool    `json:"estsessions,omitempty"`
	Hits               int     `json:"hits,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Protocol           string  `json:"protocol,omitempty"`
	SrcIPv6            string  `json:"srcipv6,omitempty"`
	TD                 int     `json:"td,omitempty"`
	TTL                int     `json:"ttl,omitempty"`
}

type NSLASLicense struct {
	DaysToExpiration   int    `json:"daystoexpiration,omitempty"`
	FileLocation       string `json:"filelocation,omitempty"`
	Filename           string `json:"filename,omitempty"`
	FixedBandwidth     bool   `json:"fixedbandwidth,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	RenewalNext        int    `json:"renewalnext,omitempty"`
	RenewalNextDate    string `json:"renewalnextdate,omitempty"`
	RenewalPrev        int    `json:"renewalprev,omitempty"`
	RenewalPrevDate    string `json:"renewalprevdate,omitempty"`
	Status             string `json:"status,omitempty"`
}

type NSConfig struct {
	All                     bool          `json:"all,omitempty"`
	Async                   bool          `json:"Async,omitempty"`
	ChangedPassword         bool          `json:"changedpassword,omitempty"`
	CIP                     string        `json:"cip,omitempty"`
	CIPHeader               string        `json:"cipheader,omitempty"`
	Config                  string        `json:"config,omitempty"`
	Config1                 string        `json:"config1,omitempty"`
	Config2                 string        `json:"config2,omitempty"`
	ConfigChanged           bool          `json:"configchanged,omitempty"`
	ConfigFile              string        `json:"configfile,omitempty"`
	CookieVersion           string        `json:"cookieversion,omitempty"`
	CRPortRange             string        `json:"crportrange,omitempty"`
	CurrentSystemTime       string        `json:"currentsytemtime,omitempty"`
	ExclusiveQuotaMaxClient int           `json:"exclusivequotamaxclient,omitempty"`
	ExclusiveQuotaSpillover int           `json:"exclusivequotaspillover,omitempty"`
	Flags                   int           `json:"flags,omitempty"`
	Force                   bool          `json:"force,omitempty"`
	FTPPortRange            string        `json:"ftpportrange,omitempty"`
	GrantQuotaMaxClient     int           `json:"grantquotamaxclient,omitempty"`
	GrantQuotaSpillover     int           `json:"grantquotaspillover,omitempty"`
	HTTPPort                []interface{} `json:"httpport,omitempty"`
	ID                      int           `json:"id,omitempty"`
	IFNum                   []string      `json:"ifnum,omitempty"`
	IgnoreDeviceSpecific    bool          `json:"ignoredevicespecific,omitempty"`
	IPAddress               string        `json:"ipaddress,omitempty"`
	LastConfigChangedTime   string        `json:"lastconfigchangedtime,omitempty"`
	LastConfigSaveTime      string        `json:"lastconfigsavetime,omitempty"`
	Level                   string        `json:"level,omitempty"`
	MappedIP                string        `json:"mappedip,omitempty"`
	MaxConn                 int           `json:"maxconn,omitempty"`
	MaxReq                  int           `json:"maxreq,omitempty"`
	Message                 string        `json:"message,omitempty"`
	Netmask                 string        `json:"netmask,omitempty"`
	NextGenAPIResource      string        `json:"_nextgenapiresource,omitempty"`
	NSVLAN                  int           `json:"nsvlan,omitempty"`
	OutType                 string        `json:"outtype,omitempty"`
	PMTUMin                 int           `json:"pmtumin,omitempty"`
	PMTUTimeout             int           `json:"pmtutimeout,omitempty"`
	PrimaryIP               string        `json:"primaryip,omitempty"`
	PrimaryIP6              string        `json:"primaryip6,omitempty"`
	Range                   int           `json:"range,omitempty"`
	RBAConfig               string        `json:"rbaconfig,omitempty"`
	Response                string        `json:"response,omitempty"`
	ResponseFile            string        `json:"responsefile,omitempty"`
	SecureCookie            string        `json:"securecookie,omitempty"`
	SecureManagementTD      int           `json:"securemanagementtd,omitempty"`
	SecureManagementTraffic string        `json:"securemanagementtraffic,omitempty"`
	SVMCmd                  int           `json:"svmcmd,omitempty"`
	SystemTime              int           `json:"systemtime,omitempty"`
	SystemType              string        `json:"systemtype,omitempty"`
	Tagged                  string        `json:"tagged,omitempty"`
	Template                bool          `json:"template,omitempty"`
	TimeZone                string        `json:"timezone,omitempty"`
	WeakPassword            bool          `json:"weakpassword,omitempty"`
}

type NSAPTLicense struct {
	BindType           string   `json:"bindtype,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	CountAvailable     string   `json:"countavailable,omitempty"`
	CountTotal         string   `json:"counttotal,omitempty"`
	DateExp            string   `json:"dateexp,omitempty"`
	DatePurchased      string   `json:"datepurchased,omitempty"`
	DateSA             string   `json:"datesa,omitempty"`
	Features           []string `json:"features,omitempty"`
	ID                 string   `json:"id,omitempty"`
	LicenseDir         string   `json:"licensedir,omitempty"`
	Name               string   `json:"name,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	Relevance          string   `json:"relevance,omitempty"`
	Response           string   `json:"response,omitempty"`
	SerialNo           string   `json:"serialno,omitempty"`
	SessionID          string   `json:"sessionid,omitempty"`
	UseProxy           string   `json:"useproxy,omitempty"`
}

type NSHTTPProfile struct {
	AdptTimeout                      string   `json:"adpttimeout,omitempty"`
	AllowOnlyWordCharactersAndHyphen string   `json:"allowonlywordcharactersandhyphen,omitempty"`
	AltSvc                           string   `json:"altsvc,omitempty"`
	AltSvcValue                      string   `json:"altsvcvalue,omitempty"`
	ApdexCltRespTimeThreshold        int      `json:"apdexcltresptimethreshold,omitempty"`
	ApdexSvrRespTimeThreshold        int      `json:"apdexsvrresptimethreshold,omitempty"`
	Builtin                          []string `json:"builtin,omitempty"`
	ClientIPHdrExpr                  string   `json:"clientiphdrexpr,omitempty"`
	CmpOnPush                        string   `json:"cmponpush,omitempty"`
	ConMultiplex                     string   `json:"conmultiplex,omitempty"`
	Count                            float64  `json:"__count,omitempty"`
	DropExtraCRLF                    string   `json:"dropextracrlf,omitempty"`
	DropExtraData                    string   `json:"dropextradata,omitempty"`
	DropInvalReqs                    string   `json:"dropinvalreqs,omitempty"`
	DropInvalReqsWarning             string   `json:"dropinvalreqswarning,omitempty"`
	Feature                          string   `json:"feature,omitempty"`
	GRPCHoldLimit                    int      `json:"grpcholdlimit,omitempty"`
	GRPCHoldTimeout                  int      `json:"grpcholdtimeout,omitempty"`
	GRPCLengthDelimitation           string   `json:"grpclengthdelimitation,omitempty"`
	HostHeaderValidation             string   `json:"hostheadervalidation,omitempty"`
	HTTP2                            string   `json:"http2,omitempty"`
	HTTP2AltSvcFrame                 string   `json:"http2altsvcframe,omitempty"`
	HTTP2Direct                      string   `json:"http2direct,omitempty"`
	HTTP2ExtendedConnect             string   `json:"http2extendedconnect,omitempty"`
	HTTP2HeaderTableSize             int      `json:"http2headertablesize,omitempty"`
	HTTP2InitialConnWindowSize       int      `json:"http2initialconnwindowsize,omitempty"`
	HTTP2InitialWindowSize           int      `json:"http2initialwindowsize,omitempty"`
	HTTP2MaxConcurrentStreams        int      `json:"http2maxconcurrentstreams,omitempty"`
	HTTP2MaxEmptyFramesPerMin        int      `json:"http2maxemptyframespermin,omitempty"`
	HTTP2MaxFrameSize                int      `json:"http2maxframesize,omitempty"`
	HTTP2MaxHeaderListSize           int      `json:"http2maxheaderlistsize,omitempty"`
	HTTP2MaxPingFramesPerMin         int      `json:"http2maxpingframespermin,omitempty"`
	HTTP2MaxResetFramesPerMin        int      `json:"http2maxresetframespermin,omitempty"`
	HTTP2MaxRxResetFramesPerMin      int      `json:"http2maxrxresetframespermin,omitempty"`
	HTTP2MaxSettingsFramesPerMin     int      `json:"http2maxsettingsframespermin,omitempty"`
	HTTP2MinSeverConn                int      `json:"http2minseverconn,omitempty"`
	HTTP2StrictCipher                string   `json:"http2strictcipher,omitempty"`
	HTTP3                            string   `json:"http3,omitempty"`
	HTTP3MaxHeaderBlockedStreams     int      `json:"http3maxheaderblockedstreams,omitempty"`
	HTTP3MaxHeaderFieldSectionSize   int      `json:"http3maxheaderfieldsectionsize,omitempty"`
	HTTP3MaxHeaderTableSize          int      `json:"http3maxheadertablesize,omitempty"`
	HTTP3MinSeverConn                int      `json:"http3minseverconn,omitempty"`
	HTTP3WebTransport                string   `json:"http3webtransport,omitempty"`
	HTTPPipelineBuffSize             int      `json:"httppipelinebuffsize,omitempty"`
	IncompHdrDelay                   int      `json:"incomphdrdelay,omitempty"`
	MarkConnReqInval                 string   `json:"markconnreqinval,omitempty"`
	MarkHTTP09Inval                  string   `json:"markhttp09inval,omitempty"`
	MarkHTTPHeaderExtraWSError       string   `json:"markhttpheaderextrawserror,omitempty"`
	MarkRFC7230NonCompliantInval     string   `json:"markrfc7230noncompliantinval,omitempty"`
	MarkTraceReqInval                string   `json:"marktracereqinval,omitempty"`
	MaxDuplicateHeaderFields         int      `json:"maxduplicateheaderfields,omitempty"`
	MaxHeaderFieldLen                int      `json:"maxheaderfieldlen,omitempty"`
	MaxHeaderLen                     int      `json:"maxheaderlen,omitempty"`
	MaxReq                           int      `json:"maxreq,omitempty"`
	MaxReusePool                     int      `json:"maxreusepool,omitempty"`
	MinReusePool                     int      `json:"minreusepool,omitempty"`
	Name                             string   `json:"name,omitempty"`
	NextGenAPIResource               string   `json:"_nextgenapiresource,omitempty"`
	PassProtocolUpgrade              string   `json:"passprotocolupgrade,omitempty"`
	PersistentETag                   string   `json:"persistentetag,omitempty"`
	RefCnt                           int      `json:"refcnt,omitempty"`
	ReqTimeout                       int      `json:"reqtimeout,omitempty"`
	ReqTimeoutAction                 string   `json:"reqtimeoutaction,omitempty"`
	ReusePoolTimeout                 int      `json:"reusepooltimeout,omitempty"`
	RTSPTunnel                       string   `json:"rtsptunnel,omitempty"`
	WebLog                           string   `json:"weblog,omitempty"`
	WebSocket                        string   `json:"websocket,omitempty"`
}

type NSFeature struct {
	AAA                bool     `json:"aaa,omitempty"`
	AdaptiveTCP        bool     `json:"adaptivetcp,omitempty"`
	APIGateway         bool     `json:"apigateway,omitempty"`
	AppFlow            bool     `json:"appflow,omitempty"`
	AppFW              bool     `json:"appfw,omitempty"`
	AppQOE             bool     `json:"appqoe,omitempty"`
	BGP                bool     `json:"bgp,omitempty"`
	Bot                bool     `json:"bot,omitempty"`
	CF                 bool     `json:"cf,omitempty"`
	CH                 bool     `json:"ch,omitempty"`
	CI                 bool     `json:"ci,omitempty"`
	CloudBridge        bool     `json:"cloudbridge,omitempty"`
	CMP                bool     `json:"cmp,omitempty"`
	ContentAccelerator bool     `json:"contentaccelerator,omitempty"`
	CQA                bool     `json:"cqa,omitempty"`
	CR                 bool     `json:"cr,omitempty"`
	CS                 bool     `json:"cs,omitempty"`
	Feature            []string `json:"feature,omitempty"`
	FEO                bool     `json:"feo,omitempty"`
	ForwardProxy       bool     `json:"forwardproxy,omitempty"`
	GSLB               bool     `json:"gslb,omitempty"`
	IC                 bool     `json:"ic,omitempty"`
	IPv6PT             bool     `json:"ipv6pt,omitempty"`
	ISIS               bool     `json:"isis,omitempty"`
	LB                 bool     `json:"lb,omitempty"`
	LSN                bool     `json:"lsn,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	OSPF               bool     `json:"ospf,omitempty"`
	Push               bool     `json:"push,omitempty"`
	RDPProxy           bool     `json:"rdpproxy,omitempty"`
	Rep                bool     `json:"rep,omitempty"`
	Responder          bool     `json:"responder,omitempty"`
	Rewrite            bool     `json:"rewrite,omitempty"`
	RIP                bool     `json:"rip,omitempty"`
	SP                 bool     `json:"sp,omitempty"`
	SSL                bool     `json:"ssl,omitempty"`
	SSLInterception    bool     `json:"sslinterception,omitempty"`
	SSLVPN             bool     `json:"sslvpn,omitempty"`
	VideoOptimization  bool     `json:"videooptimization,omitempty"`
	WL                 bool     `json:"wl,omitempty"`
}

type NSServicePathBinding struct {
	NSServicePathNSServiceFunctionBinding []interface{} `json:"nsservicepath_nsservicefunction_binding,omitempty"`
	ServicePathName                       string        `json:"servicepathname,omitempty"`
}

type NSLicenseParameters struct {
	Alert1GraceTimeout       int    `json:"alert1gracetimeout,omitempty"`
	Alert2GraceTimeout       int    `json:"alert2gracetimeout,omitempty"`
	HeartbeatInterval        int    `json:"heartbeatinterval,omitempty"`
	InventoryRefreshInterval int    `json:"inventoryrefreshinterval,omitempty"`
	LicenseExpiryAlertTime   int    `json:"licenseexpiryalerttime,omitempty"`
	NextGenAPIResource       string `json:"_nextgenapiresource,omitempty"`
}

type NSLimitIdentifier struct {
	ComputedTrapTimeSlice    int      `json:"computedtraptimeslice,omitempty"`
	Count                    float64  `json:"__count,omitempty"`
	Drop                     int      `json:"drop,omitempty"`
	Hits                     int      `json:"hits,omitempty"`
	LimitIdentifier          string   `json:"limitidentifier,omitempty"`
	LimitType                string   `json:"limittype,omitempty"`
	MaxBandwidth             int      `json:"maxbandwidth,omitempty"`
	Mode                     string   `json:"mode,omitempty"`
	NextGenAPIResource       string   `json:"_nextgenapiresource,omitempty"`
	NGName                   string   `json:"ngname,omitempty"`
	ReferenceCount           int      `json:"referencecount,omitempty"`
	Rule                     []string `json:"rule,omitempty"`
	SelectorName             string   `json:"selectorname,omitempty"`
	Threshold                int      `json:"threshold,omitempty"`
	Time                     int      `json:"time,omitempty"`
	TimeSlice                int      `json:"timeslice,omitempty"`
	Total                    int      `json:"total,omitempty"`
	TrapsComputedInTimeSlice int      `json:"trapscomputedintimeslice,omitempty"`
	TrapsInTimeSlice         int      `json:"trapsintimeslice,omitempty"`
}

type NSHardware struct {
	BMCRevision        string `json:"bmcrevision,omitempty"`
	CPUFrequency       int    `json:"cpufrequncy,omitempty"`
	EncodedSerialNo    string `json:"encodedserialno,omitempty"`
	Host               string `json:"host,omitempty"`
	HostID             int    `json:"hostid,omitempty"`
	HWDescription      string `json:"hwdescription,omitempty"`
	ManufactureDay     int    `json:"manufactureday,omitempty"`
	ManufactureMonth   int    `json:"manufacturemonth,omitempty"`
	ManufactureYear    int    `json:"manufactureyear,omitempty"`
	NetScalerUUID      string `json:"netscaleruuid,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	SerialNo           string `json:"serialno,omitempty"`
	SysID              int    `json:"sysid,omitempty"`
}

type NSTrafficDomain struct {
	AliasName          string  `json:"aliasname,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	State              string  `json:"state,omitempty"`
	TD                 int     `json:"td,omitempty"`
	VMAC               string  `json:"vmac,omitempty"`
}

type NSPartitionVlanBinding struct {
	PartitionName string `json:"partitionname,omitempty"`
	VLAN          int    `json:"vlan,omitempty"`
}

type NSCQAParam struct {
	HarqRetxDelay      int     `json:"harqretxdelay,omitempty"`
	LR1CoefList        string  `json:"lr1coeflist,omitempty"`
	LR1ProbThresh      float64 `json:"lr1probthresh,omitempty"`
	LR2CoefList        string  `json:"lr2coeflist,omitempty"`
	LR2ProbThresh      float64 `json:"lr2probthresh,omitempty"`
	MinRTTNet1         int     `json:"minrttnet1,omitempty"`
	MinRTTNet2         int     `json:"minrttnet2,omitempty"`
	MinRTTNet3         int     `json:"minrttnet3,omitempty"`
	Net1CCLScale       string  `json:"net1cclscale,omitempty"`
	Net1CSQScale       string  `json:"net1csqscale,omitempty"`
	Net1Label          string  `json:"net1label,omitempty"`
	Net1LogCoef        string  `json:"net1logcoef,omitempty"`
	Net2CCLScale       string  `json:"net2cclscale,omitempty"`
	Net2CSQScale       string  `json:"net2csqscale,omitempty"`
	Net2Label          string  `json:"net2label,omitempty"`
	Net2LogCoef        string  `json:"net2logcoef,omitempty"`
	Net3CCLScale       string  `json:"net3cclscale,omitempty"`
	Net3CSQScale       string  `json:"net3csqscale,omitempty"`
	Net3Label          string  `json:"net3label,omitempty"`
	Net3LogCoef        string  `json:"net3logcoef,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
}

type NSStats struct {
	CleanupLevel string `json:"cleanuplevel,omitempty"`
}

type NSLimitIdentifierBinding struct {
	LimitIdentifier                         string        `json:"limitidentifier,omitempty"`
	NSLimitIdentifierNSLimitSessionsBinding []interface{} `json:"nslimitidentifier_nslimitsessions_binding,omitempty"`
}
