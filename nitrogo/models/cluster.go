package models

// cluster configuration structs
type ClusterNodeBinding struct {
	ClusterNodeRouteMonitorBinding []interface{} `json:"clusternode_routemonitor_binding,omitempty"`
	NodeID                         int           `json:"nodeid,omitempty"`
}

type ClusterNode struct {
	Backplane                  string        `json:"backplane,omitempty"`
	CfgFlags                   int           `json:"cfgflags,omitempty"`
	ClearNodeGroupConfig       string        `json:"clearnodegroupconfig,omitempty"`
	ClusterHealth              string        `json:"clusterhealth,omitempty"`
	Count                      float64       `json:"__count,omitempty"`
	Delay                      int           `json:"delay,omitempty"`
	DisabledIfaces             string        `json:"disabledifaces,omitempty"`
	EffectiveState             string        `json:"effectivestate,omitempty"`
	EnabledIfaces              string        `json:"enabledifaces,omitempty"`
	Force                      bool          `json:"force,omitempty"`
	HAMonIfaces                string        `json:"hamonifaces,omitempty"`
	Health                     string        `json:"health,omitempty"`
	IfacesList                 []string      `json:"ifaceslist,omitempty"`
	IPAddress                  string        `json:"ipaddress,omitempty"`
	IsConfigurationCoordinator bool          `json:"isconfigurationcoordinator,omitempty"`
	IsLocalNode                bool          `json:"islocalnode,omitempty"`
	MasterState                string        `json:"masterstate,omitempty"`
	Name                       string        `json:"name,omitempty"`
	Netmask                    string        `json:"netmask,omitempty"`
	NextGenAPIResource         string        `json:"_nextgenapiresource,omitempty"`
	NodeGroup                  string        `json:"nodegroup,omitempty"`
	NodeID                     int           `json:"nodeid,omitempty"`
	NodeJumboNotSupported      bool          `json:"nodejumbonotsupported,omitempty"`
	NodeLicenseMismatch        bool          `json:"nodelicensemismatch,omitempty"`
	NodeList                   []interface{} `json:"nodelist,omitempty"`
	NodeRSSKeyMismatch         bool          `json:"nodersskeymismatch,omitempty"`
	OperationalSyncState       string        `json:"operationalsyncstate,omitempty"`
	PartialFailIfaces          string        `json:"partialfailifaces,omitempty"`
	Priority                   int           `json:"priority,omitempty"`
	RouteMonitor               string        `json:"routemonitor,omitempty"`
	State                      string        `json:"state,omitempty"`
	SyncFailureReason          string        `json:"syncfailurereason,omitempty"`
	SyncState                  string        `json:"syncstate,omitempty"`
	TunnelMode                 string        `json:"tunnelmode,omitempty"`
}

type ClusterNodeGroupStreamIdentifierBinding struct {
	IdentifierName string `json:"identifiername,omitempty"`
	Name           string `json:"name,omitempty"`
}

type Cluster struct {
	Clip     string `json:"clip,omitempty"`
	Password string `json:"password,omitempty"`
}

type ClusterInstance struct {
	AdminState                 string  `json:"adminstate,omitempty"`
	BackplaneBasedView         string  `json:"backplanebasedview,omitempty"`
	CLID                       int     `json:"clid,omitempty"`
	ClusterCLIPFailure         bool    `json:"clusterclipfailure,omitempty"`
	ClusterHBHMacErrorDetected bool    `json:"clusterhbhmacerrordetected,omitempty"`
	ClusterNoHeartbeatOnNode   bool    `json:"clusternoheartbeatonnode,omitempty"`
	ClusterNoLinksetMBF        bool    `json:"clusternolinksetmbf,omitempty"`
	ClusterNoSpottedIP         bool    `json:"clusternospottedip,omitempty"`
	ClusterProxyARP            string  `json:"clusterproxyarp,omitempty"`
	ClusterTunnelModeMismatch  bool    `json:"clustertunnelmodemismatch,omitempty"`
	Count                      float64 `json:"__count,omitempty"`
	DeadInterval               int     `json:"deadinterval,omitempty"`
	DFDRetainL2Params          string  `json:"dfdretainl2params,omitempty"`
	HelloInterval              int     `json:"hellointerval,omitempty"`
	HeterogeneousFlag          string  `json:"heterogeneousflag,omitempty"`
	Inc                        string  `json:"inc,omitempty"`
	JumboNotSupported          bool    `json:"jumbonotsupported,omitempty"`
	LicenseMismatch            bool    `json:"licensemismatch,omitempty"`
	NextGenAPIResource         string  `json:"_nextgenapiresource,omitempty"`
	NodeGroup                  string  `json:"nodegroup,omitempty"`
	NodeGroupStateWarning      bool    `json:"nodegroupstatewarning,omitempty"`
	NodePENumMismatch          bool    `json:"nodepenummismatch,omitempty"`
	OperationalPropState       string  `json:"operationalpropstate,omitempty"`
	OperationalState           string  `json:"operationalstate,omitempty"`
	PENumMismatch              bool    `json:"penummismatch,omitempty"`
	Preemption                 string  `json:"preemption,omitempty"`
	ProcessLocal               string  `json:"processlocal,omitempty"`
	PropState                  string  `json:"propstate,omitempty"`
	QuorumType                 string  `json:"quorumtype,omitempty"`
	RetainConnectionsOnCluster string  `json:"retainconnectionsoncluster,omitempty"`
	RSSKeyMismatch             bool    `json:"rsskeymismatch,omitempty"`
	SecureHeartbeats           string  `json:"secureheartbeats,omitempty"`
	Status                     string  `json:"status,omitempty"`
	SyncStatusStrictMode       string  `json:"syncstatusstrictmode,omitempty"`
	ValidMTU                   int     `json:"validmtu,omitempty"`
}

type ClusterNodeRouteMonitorBinding struct {
	Netmask       string `json:"netmask,omitempty"`
	NodeID        int    `json:"nodeid,omitempty"`
	RouteMonitor  string `json:"routemonitor,omitempty"`
	RouteMonState int    `json:"routemonstate,omitempty"`
}

type ClusterNodeGroupGSLBVServerBinding struct {
	Name    string `json:"name,omitempty"`
	VServer string `json:"vserver,omitempty"`
}

type ClusterNodeGroupBinding struct {
	ClusterNodeGroupAuthenticationVServerBinding []interface{} `json:"clusternodegroup_authenticationvserver_binding,omitempty"`
	ClusterNodeGroupClusterNodeBinding           []interface{} `json:"clusternodegroup_clusternode_binding,omitempty"`
	ClusterNodeGroupCRVServerBinding             []interface{} `json:"clusternodegroup_crvserver_binding,omitempty"`
	ClusterNodeGroupCSVServerBinding             []interface{} `json:"clusternodegroup_csvserver_binding,omitempty"`
	ClusterNodeGroupGSLBSiteBinding              []interface{} `json:"clusternodegroup_gslbsite_binding,omitempty"`
	ClusterNodeGroupGSLBVServerBinding           []interface{} `json:"clusternodegroup_gslbvserver_binding,omitempty"`
	ClusterNodeGroupLBVServerBinding             []interface{} `json:"clusternodegroup_lbvserver_binding,omitempty"`
	ClusterNodeGroupNSLimitIdentifierBinding     []interface{} `json:"clusternodegroup_nslimitidentifier_binding,omitempty"`
	ClusterNodeGroupServiceBinding               []interface{} `json:"clusternodegroup_service_binding,omitempty"`
	ClusterNodeGroupStreamIdentifierBinding      []interface{} `json:"clusternodegroup_streamidentifier_binding,omitempty"`
	ClusterNodeGroupVPNVServerBinding            []interface{} `json:"clusternodegroup_vpnvserver_binding,omitempty"`
	Name                                         string        `json:"name,omitempty"`
}

type ClusterNodeGroupAuthenticationVServerBinding struct {
	Name    string `json:"name,omitempty"`
	VServer string `json:"vserver,omitempty"`
}

type ClusterNodeGroupLBVServerBinding struct {
	Name    string `json:"name,omitempty"`
	VServer string `json:"vserver,omitempty"`
}

type ClusterNodeGroupCRVServerBinding struct {
	Name    string `json:"name,omitempty"`
	VServer string `json:"vserver,omitempty"`
}

type ClusterNodeGroupNSLimitIdentifierBinding struct {
	IdentifierName string `json:"identifiername,omitempty"`
	Name           string `json:"name,omitempty"`
}

type ClusterSync struct {
}

type ClusterNodeGroupGSLBSiteBinding struct {
	GSLBSite string `json:"gslbsite,omitempty"`
	Name     string `json:"name,omitempty"`
}

type ClusterNodeGroupServiceBinding struct {
	Name    string `json:"name,omitempty"`
	Service string `json:"service,omitempty"`
}

type ClusterNodeGroup struct {
	ActiveList               []interface{} `json:"activelist,omitempty"`
	BackupList               []interface{} `json:"backuplist,omitempty"`
	BackupNodeMask           int           `json:"backupnodemask,omitempty"`
	BoundedEntitiesCntFromPE int           `json:"boundedentitiescntfrompe,omitempty"`
	Count                    float64       `json:"__count,omitempty"`
	CurrentNodeMask          int           `json:"currentnodemask,omitempty"`
	Name                     string        `json:"name,omitempty"`
	NextGenAPIResource       string        `json:"_nextgenapiresource,omitempty"`
	Priority                 int           `json:"priority,omitempty"`
	State                    string        `json:"state,omitempty"`
	Sticky                   string        `json:"sticky,omitempty"`
	Strict                   string        `json:"strict,omitempty"`
}

type ClusterNodeGroupClusterNodeBinding struct {
	Name string `json:"name,omitempty"`
	Node int    `json:"node,omitempty"`
}

type ClusterSyncFailures struct {
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
}

type ClusterPropStatus struct {
	CmdStrs            string  `json:"cmdstrs,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NodeID             int     `json:"nodeid,omitempty"`
	NumPropCmdFailed   int     `json:"numpropcmdfailed,omitempty"`
}

type ClusterInstanceBinding struct {
	CLID                              int           `json:"clid,omitempty"`
	ClusterInstanceClusterNodeBinding []interface{} `json:"clusterinstance_clusternode_binding,omitempty"`
}

type ClusterFiles struct {
	Mode []string `json:"mode,omitempty"`
}

type ClusterInstanceClusterNodeBinding struct {
	CLID                       int    `json:"clid,omitempty"`
	ClusterHealth              string `json:"clusterhealth,omitempty"`
	EffectiveState             string `json:"effectivestate,omitempty"`
	Health                     string `json:"health,omitempty"`
	IPAddress                  string `json:"ipaddress,omitempty"`
	IsConfigurationCoordinator bool   `json:"isconfigurationcoordinator,omitempty"`
	IsLocalNode                bool   `json:"islocalnode,omitempty"`
	MasterState                string `json:"masterstate,omitempty"`
	NodeID                     int    `json:"nodeid,omitempty"`
	NodeJumboNotSupported      bool   `json:"nodejumbonotsupported,omitempty"`
	NodeLicenseMismatch        bool   `json:"nodelicensemismatch,omitempty"`
	NodeRSSKeyMismatch         bool   `json:"nodersskeymismatch,omitempty"`
	State                      string `json:"state,omitempty"`
}

type ClusterNodeGroupVPNVServerBinding struct {
	Name    string `json:"name,omitempty"`
	VServer string `json:"vserver,omitempty"`
}

type ClusterNodeGroupCSVServerBinding struct {
	Name    string `json:"name,omitempty"`
	VServer string `json:"vserver,omitempty"`
}
