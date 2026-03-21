package models

// dns configuration structs
type DNSNSECRec struct {
	Count              float64  `json:"__count,omitempty"`
	ECSSubnet          string   `json:"ecssubnet,omitempty"`
	Hostname           string   `json:"hostname,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	NextNSEC           string   `json:"nextnsec,omitempty"`
	NextRecs           []string `json:"nextrecs,omitempty"`
	TTL                int      `json:"ttl,omitempty"`
	TypeField          string   `json:"type,omitempty"`
}

type DNSTxtRec struct {
	AuthType           string   `json:"authtype,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Domain             string   `json:"domain,omitempty"`
	ECSSubnet          string   `json:"ecssubnet,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	NodeID             int      `json:"nodeid,omitempty"`
	RecordID           int      `json:"recordid,omitempty"`
	String             []string `json:"String,omitempty"`
	TTL                int      `json:"ttl,omitempty"`
	TypeField          string   `json:"type,omitempty"`
}

type DNSZone struct {
	Count              float64  `json:"__count,omitempty"`
	DNSSECOffload      string   `json:"dnssecoffload,omitempty"`
	Flags              int      `json:"flags,omitempty"`
	KeyName            []string `json:"keyname,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	NSEC               string   `json:"nsec,omitempty"`
	ProxyMode          string   `json:"proxymode,omitempty"`
	TypeField          string   `json:"type,omitempty"`
	ZoneName           string   `json:"zonename,omitempty"`
}

type DNSPolicyLabelBinding struct {
	DNSPolicyLabelDNSPolicyBinding     []any  `json:"dnspolicylabel_dnspolicy_binding,omitempty"`
	DNSPolicyLabelPolicyBindingBinding []any  `json:"dnspolicylabel_policybinding_binding,omitempty"`
	LabelName                          string `json:"labelname,omitempty"`
}

type DNSDSFile struct {
	KeyName            string `json:"keyname,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
}

type DNSNegativeCacheRecords struct {
	Count              float64 `json:"__count,omitempty"`
	ECSSubnet          string  `json:"ecssubnet,omitempty"`
	Hostname           string  `json:"hostname,omitempty"`
	NegCacheType       string  `json:"negcachetype,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NodeID             int     `json:"nodeid,omitempty"`
	PEID               int     `json:"peid,omitempty"`
	QueryType          string  `json:"querytype,omitempty"`
	RDClient           string  `json:"rdclient,omitempty"`
	TTL                int     `json:"ttl,omitempty"`
}

type DNSPolicyLabelDNSPolicyBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelName        string `json:"invoke_labelname,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type DNSPtrRec struct {
	AuthType           string  `json:"authtype,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Domain             string  `json:"domain,omitempty"`
	ECSSubnet          string  `json:"ecssubnet,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NodeID             int     `json:"nodeid,omitempty"`
	ReverseDomain      string  `json:"reversedomain,omitempty"`
	TTL                int     `json:"ttl,omitempty"`
	TypeField          string  `json:"type,omitempty"`
}

type DNSZoneDomainBinding struct {
	Domain   string   `json:"domain,omitempty"`
	NextRecs []string `json:"nextrecs,omitempty"`
	ZoneName string   `json:"zonename,omitempty"`
}

type DNSZoneBinding struct {
	DNSZoneDNSKeyBinding     []any  `json:"dnszone_dnskey_binding,omitempty"`
	DNSZoneGSLBDomainBinding []any  `json:"dnszone_gslbdomain_binding,omitempty"`
	ZoneName                 string `json:"zonename,omitempty"`
}

type DNSGlobalDNSPolicyBinding struct {
	FlowType               int    `json:"flowtype,omitempty"`
	GlobalBindType         string `json:"globalbindtype,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	NumPol                 int    `json:"numpol,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
}

type DNSNameServer struct {
	ClMonOwner         int     `json:"clmonowner,omitempty"`
	ClMonView          int     `json:"clmonview,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	DNSProfileName     string  `json:"dnsprofilename,omitempty"`
	DNSVServerName     string  `json:"dnsvservername,omitempty"`
	IP                 string  `json:"ip,omitempty"`
	Local              bool    `json:"local,omitempty"`
	NameServerState    string  `json:"nameserverstate,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Port               int     `json:"port,omitempty"`
	ServiceName        string  `json:"servicename,omitempty"`
	State              string  `json:"state,omitempty"`
	TypeField          string  `json:"type,omitempty"`
}

type DNSPolicyDNSPolicyLabelBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type DNSSrvRec struct {
	AuthType           string  `json:"authtype,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Domain             string  `json:"domain,omitempty"`
	ECSSubnet          string  `json:"ecssubnet,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NodeID             int     `json:"nodeid,omitempty"`
	Port               int     `json:"port,omitempty"`
	Priority           int     `json:"priority,omitempty"`
	Target             string  `json:"target,omitempty"`
	TTL                int     `json:"ttl,omitempty"`
	TypeField          string  `json:"type,omitempty"`
	Weight             int     `json:"weight,omitempty"`
}

type DNSViewBinding struct {
	DNSViewDNSPolicyBinding   []any  `json:"dnsview_dnspolicy_binding,omitempty"`
	DNSViewGSLBServiceBinding []any  `json:"dnsview_gslbservice_binding,omitempty"`
	ViewName                  string `json:"viewname,omitempty"`
}

type DNSPolicyLabel struct {
	Count                  float64 `json:"__count,omitempty"`
	Description            string  `json:"description,omitempty"`
	FlowType               int     `json:"flowtype,omitempty"`
	GotoPriorityExpression string  `json:"gotopriorityexpression,omitempty"`
	Hits                   int     `json:"hits,omitempty"`
	InvokeLabelName        string  `json:"invoke_labelname,omitempty"`
	IsDefault              bool    `json:"isdefault,omitempty"`
	LabelName              string  `json:"labelname,omitempty"`
	LabelType              string  `json:"labeltype,omitempty"`
	NewName                string  `json:"newname,omitempty"`
	NextGenAPIResource     string  `json:"_nextgenapiresource,omitempty"`
	NumPol                 int     `json:"numpol,omitempty"`
	Priority               int     `json:"priority,omitempty"`
	Transform              string  `json:"transform,omitempty"`
}

type DNSProxyRecords struct {
	NegRecType string `json:"negrectype,omitempty"`
	TypeField  string `json:"type,omitempty"`
}

type DNSNSRec struct {
	AuthType           string  `json:"authtype,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Domain             string  `json:"domain,omitempty"`
	ECSSubnet          string  `json:"ecssubnet,omitempty"`
	NameServer         string  `json:"nameserver,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NodeID             int     `json:"nodeid,omitempty"`
	TTL                int     `json:"ttl,omitempty"`
	TypeField          string  `json:"type,omitempty"`
}

type DNSPolicyDNSGlobalBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type DNSPolicy struct {
	ActionName         string   `json:"actionname,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	CacheBypass        string   `json:"cachebypass,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Description        string   `json:"description,omitempty"`
	Drop               string   `json:"drop,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	LogAction          string   `json:"logaction,omitempty"`
	Name               string   `json:"name,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	PreferredLocation  string   `json:"preferredlocation,omitempty"`
	PreferredLocList   []string `json:"preferredloclist,omitempty"`
	Rule               string   `json:"rule,omitempty"`
	UndefHits          int      `json:"undefhits,omitempty"`
	ViewName           string   `json:"viewname,omitempty"`
}

type DNSAction struct {
	ActionName         string   `json:"actionname,omitempty"`
	ActionType         string   `json:"actiontype,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	CacheBypass        string   `json:"cachebypass,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	DNSProfileName     string   `json:"dnsprofilename,omitempty"`
	Drop               string   `json:"drop,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	IPAddress          []string `json:"ipaddress,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	PreferredLocList   []string `json:"preferredloclist,omitempty"`
	TTL                int      `json:"ttl,omitempty"`
	ViewName           string   `json:"viewname,omitempty"`
}

type DNSGlobalBinding struct {
	DNSGlobalDNSPolicyBinding []any `json:"dnsglobal_dnspolicy_binding,omitempty"`
}

type DNSSOARec struct {
	AuthType           string  `json:"authtype,omitempty"`
	Contact            string  `json:"contact,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Domain             string  `json:"domain,omitempty"`
	ECSSubnet          string  `json:"ecssubnet,omitempty"`
	Expire             int     `json:"expire,omitempty"`
	Minimum            int     `json:"minimum,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NodeID             int     `json:"nodeid,omitempty"`
	OriginServer       string  `json:"originserver,omitempty"`
	Refresh            int     `json:"refresh,omitempty"`
	Retry              int     `json:"retry,omitempty"`
	Serial             int     `json:"serial,omitempty"`
	TTL                int     `json:"ttl,omitempty"`
	TypeField          string  `json:"type,omitempty"`
}

type DNSPolicy64Binding struct {
	DNSPolicy64LBVServerBinding []any  `json:"dnspolicy64_lbvserver_binding,omitempty"`
	Name                        string `json:"name,omitempty"`
}

type DNSPolicyBinding struct {
	DNSPolicyDNSGlobalBinding      []any  `json:"dnspolicy_dnsglobal_binding,omitempty"`
	DNSPolicyDNSPolicyLabelBinding []any  `json:"dnspolicy_dnspolicylabel_binding,omitempty"`
	Name                           string `json:"name,omitempty"`
}

type DNSZoneDNSKeyBinding struct {
	Expires          int      `json:"expires,omitempty"`
	KeyName          []string `json:"keyname,omitempty"`
	SigInceptionTime []any    `json:"siginceptiontime,omitempty"`
	Signed           int      `json:"signed,omitempty"`
	ZoneName         string   `json:"zonename,omitempty"`
}

type DNSPolicy64LBVServerBinding struct {
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type DNSSuffix struct {
	Count              float64 `json:"__count,omitempty"`
	DNSSuffix          string  `json:"Dnssuffix,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
}

type DNSParameter struct {
	AutoSaveKeyOps               string   `json:"autosavekeyops,omitempty"`
	Builtin                      []string `json:"builtin,omitempty"`
	CacheECSZeroPrefix           string   `json:"cacheecszeroprefix,omitempty"`
	CacheHitBypass               string   `json:"cachehitbypass,omitempty"`
	CacheNoExpire                string   `json:"cachenoexpire,omitempty"`
	CacheRecords                 string   `json:"cacherecords,omitempty"`
	DNS64Timeout                 int      `json:"dns64timeout,omitempty"`
	DNSRootReferral              string   `json:"dnsrootreferral,omitempty"`
	DNSSec                       string   `json:"dnssec,omitempty"`
	ECSMaxSubnets                int      `json:"ecsmaxsubnets,omitempty"`
	Feature                      string   `json:"feature,omitempty"`
	MaxCacheSize                 int      `json:"maxcachesize,omitempty"`
	MaxNegativeCacheSize         int      `json:"maxnegativecachesize,omitempty"`
	MaxNegCacheTTL               int      `json:"maxnegcachettl,omitempty"`
	MaxPipeline                  int      `json:"maxpipeline,omitempty"`
	MaxTTL                       int      `json:"maxttl,omitempty"`
	MaxUDPPacketSize             int      `json:"maxudppacketsize,omitempty"`
	MinTTL                       int      `json:"minttl,omitempty"`
	NameLookupPriority           string   `json:"namelookuppriority,omitempty"`
	NextGenAPIResource           string   `json:"_nextgenapiresource,omitempty"`
	NXDomainRateLimitThreshold   int      `json:"nxdomainratelimitthreshold,omitempty"`
	NXDomainThresholdCrossed     int      `json:"nxdomainthresholdcrossed,omitempty"`
	Recursion                    string   `json:"recursion,omitempty"`
	ResolutionOrder              string   `json:"resolutionorder,omitempty"`
	ResolverMaxActiveResolutions int      `json:"resolvermaxactiveresolutions,omitempty"`
	ResolverMaxTCPConnections    int      `json:"resolvermaxtcpconnections,omitempty"`
	ResolverMaxTCPTimeout        int      `json:"resolvermaxtcptimeout,omitempty"`
	Retries                      int      `json:"retries,omitempty"`
	SplitPktQueryProcessing      string   `json:"splitpktqueryprocessing,omitempty"`
	ZoneTransfer                 string   `json:"zonetransfer,omitempty"`
}

type DNSAction64 struct {
	ActionName         string   `json:"actionname,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	ExcludeRule        string   `json:"excluderule,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	MappedRule         string   `json:"mappedrule,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	Prefix             string   `json:"prefix,omitempty"`
}

type DNSCnameRec struct {
	AliasName          string  `json:"aliasname,omitempty"`
	AuthType           string  `json:"authtype,omitempty"`
	CanonicalName      string  `json:"canonicalname,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	ECSSubnet          string  `json:"ecssubnet,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NodeID             int     `json:"nodeid,omitempty"`
	TTL                int     `json:"ttl,omitempty"`
	TypeField          string  `json:"type,omitempty"`
	VServerName        string  `json:"vservername,omitempty"`
}

type DNSPolicy64 struct {
	Action             string  `json:"action,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Description        string  `json:"description,omitempty"`
	Hits               int     `json:"hits,omitempty"`
	LabelName          string  `json:"labelname,omitempty"`
	LabelType          string  `json:"labeltype,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Rule               string  `json:"rule,omitempty"`
	UndefHits          int     `json:"undefhits,omitempty"`
}

type DNSProfile struct {
	CacheECSResponses            string  `json:"cacheecsresponses,omitempty"`
	CacheNegativeResponses       string  `json:"cachenegativeresponses,omitempty"`
	CacheRecords                 string  `json:"cacherecords,omitempty"`
	Count                        float64 `json:"__count,omitempty"`
	DNSAnswerSecLogging          string  `json:"dnsanswerseclogging,omitempty"`
	DNSErrorLogging              string  `json:"dnserrorlogging,omitempty"`
	DNSExtendedLogging           string  `json:"dnsextendedlogging,omitempty"`
	DNSProfileName               string  `json:"dnsprofilename,omitempty"`
	DNSQueryLogging              string  `json:"dnsquerylogging,omitempty"`
	DropMultiQueryRequest        string  `json:"dropmultiqueryrequest,omitempty"`
	InsertECS                    string  `json:"insertecs,omitempty"`
	MaxCacheableECSPrefixLength  int     `json:"maxcacheableecsprefixlength,omitempty"`
	MaxCacheableECSPrefixLength6 int     `json:"maxcacheableecsprefixlength6,omitempty"`
	NextGenAPIResource           string  `json:"_nextgenapiresource,omitempty"`
	RecursiveResolution          string  `json:"recursiveresolution,omitempty"`
	ReferenceCount               int     `json:"referencecount,omitempty"`
	ReplaceECS                   string  `json:"replaceecs,omitempty"`
}

type DNSViewGSLBServiceBinding struct {
	GSLBServiceName string `json:"gslbservicename,omitempty"`
	IPAddress       string `json:"ipaddress,omitempty"`
	ViewName        string `json:"viewname,omitempty"`
}

type DNSAAAARec struct {
	AuthType           string  `json:"authtype,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	ECSSubnet          string  `json:"ecssubnet,omitempty"`
	Hostname           string  `json:"hostname,omitempty"`
	IPv6Address        string  `json:"ipv6address,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NodeID             int     `json:"nodeid,omitempty"`
	TTL                int     `json:"ttl,omitempty"`
	TypeField          string  `json:"type,omitempty"`
	VServerName        string  `json:"vservername,omitempty"`
}

type DNSView struct {
	Count              float64 `json:"__count,omitempty"`
	Flags              int     `json:"flags,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	ViewName           string  `json:"viewname,omitempty"`
}

type DNSKey struct {
	ActivationTimeStr  string  `json:"activationtimestr,omitempty"`
	Algorithm          string  `json:"algorithm,omitempty"`
	AutoRollover       string  `json:"autorollover,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	CreateTimeStr      string  `json:"createtimestr,omitempty"`
	DeletionTimeStr    string  `json:"deletiontimestr,omitempty"`
	Expires            int     `json:"expires,omitempty"`
	ExpiryTimeStr      string  `json:"expirytimestr,omitempty"`
	FilenamePrefix     string  `json:"filenameprefix,omitempty"`
	KeyName            string  `json:"keyname,omitempty"`
	KeySize            int     `json:"keysize,omitempty"`
	KeyType            string  `json:"keytype,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NotificationPeriod int     `json:"notificationperiod,omitempty"`
	Password           string  `json:"password,omitempty"`
	PrivateKey         string  `json:"privatekey,omitempty"`
	PublicKey          string  `json:"publickey,omitempty"`
	Revoke             bool    `json:"revoke,omitempty"`
	RolloverFailRC     int     `json:"rolloverfailrc,omitempty"`
	RolloverMethod     string  `json:"rollovermethod,omitempty"`
	Src                string  `json:"src,omitempty"`
	State              string  `json:"state,omitempty"`
	Tag                int     `json:"tag,omitempty"`
	TTL                int     `json:"ttl,omitempty"`
	TypeField          string  `json:"type,omitempty"`
	Units1             string  `json:"units1,omitempty"`
	Units2             string  `json:"units2,omitempty"`
	ZoneName           string  `json:"zonename,omitempty"`
}

type DNSNaptrRec struct {
	AuthType           string  `json:"authtype,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Domain             string  `json:"domain,omitempty"`
	ECSSubnet          string  `json:"ecssubnet,omitempty"`
	Flags              string  `json:"flags,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NodeID             int     `json:"nodeid,omitempty"`
	Order              int     `json:"order,omitempty"`
	Preference         int     `json:"preference,omitempty"`
	RecordID           int     `json:"recordid,omitempty"`
	Regexp             string  `json:"regexp,omitempty"`
	Replacement        string  `json:"replacement,omitempty"`
	Services           string  `json:"services,omitempty"`
	TTL                int     `json:"ttl,omitempty"`
	TypeField          string  `json:"type,omitempty"`
	VServerName        string  `json:"vservername,omitempty"`
}

type DNSViewDNSPolicyBinding struct {
	DNSPolicyName string `json:"dnspolicyname,omitempty"`
	ViewName      string `json:"viewname,omitempty"`
}

type DNSSubnetCache struct {
	All                bool     `json:"all,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	ECSSubnet          string   `json:"ecssubnet,omitempty"`
	Hostname           string   `json:"hostname,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	NextRecs           []string `json:"nextrecs,omitempty"`
	NodeID             int      `json:"nodeid,omitempty"`
}

type DNSAddRec struct {
	AuthType           string  `json:"authtype,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	ECSSubnet          string  `json:"ecssubnet,omitempty"`
	Hostname           string  `json:"hostname,omitempty"`
	IPAddress          string  `json:"ipaddress,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NodeID             int     `json:"nodeid,omitempty"`
	TTL                int     `json:"ttl,omitempty"`
	TypeField          string  `json:"type,omitempty"`
	VServerName        string  `json:"vservername,omitempty"`
}

type DNSMXRec struct {
	AuthType           string  `json:"authtype,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Domain             string  `json:"domain,omitempty"`
	ECSSubnet          string  `json:"ecssubnet,omitempty"`
	MX                 string  `json:"mx,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NodeID             int     `json:"nodeid,omitempty"`
	Pref               int     `json:"pref,omitempty"`
	TTL                int     `json:"ttl,omitempty"`
	TypeField          string  `json:"type,omitempty"`
}

type DNSCAARec struct {
	AuthType           string  `json:"authtype,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Domain             string  `json:"domain,omitempty"`
	ECSSubnet          string  `json:"ecssubnet,omitempty"`
	Flag               string  `json:"flag,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NodeID             int     `json:"nodeid,omitempty"`
	RecordID           int     `json:"recordid,omitempty"`
	Tag                string  `json:"tag,omitempty"`
	TTL                int     `json:"ttl,omitempty"`
	TypeField          string  `json:"type,omitempty"`
	ValueString        string  `json:"valuestring,omitempty"`
}

type DNSPolicyLabelPolicyBindingBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelName        string `json:"invoke_labelname,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}
