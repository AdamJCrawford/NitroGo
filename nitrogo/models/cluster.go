package models

// cluster configuration structs
type ClusternodeBinding struct {
	ClusternodeRoutemonitorBinding []interface{} `json:"clusternode_routemonitor_binding,omitempty"`
	Nodeid                         int           `json:"nodeid,omitempty"`
}

type Clusternode struct {
	Backplane                  string        `json:"backplane,omitempty"`
	Cfgflags                   int           `json:"cfgflags,omitempty"`
	Clearnodegroupconfig       string        `json:"clearnodegroupconfig,omitempty"`
	Clusterhealth              string        `json:"clusterhealth,omitempty"`
	Count                      float64       `json:"__count,omitempty"`
	Delay                      int           `json:"delay,omitempty"`
	Disabledifaces             string        `json:"disabledifaces,omitempty"`
	Effectivestate             string        `json:"effectivestate,omitempty"`
	Enabledifaces              string        `json:"enabledifaces,omitempty"`
	Force                      bool          `json:"force,omitempty"`
	Hamonifaces                string        `json:"hamonifaces,omitempty"`
	Health                     string        `json:"health,omitempty"`
	Ifaceslist                 []string      `json:"ifaceslist,omitempty"`
	Ipaddress                  string        `json:"ipaddress,omitempty"`
	Isconfigurationcoordinator bool          `json:"isconfigurationcoordinator,omitempty"`
	Islocalnode                bool          `json:"islocalnode,omitempty"`
	Masterstate                string        `json:"masterstate,omitempty"`
	Name                       string        `json:"name,omitempty"`
	Netmask                    string        `json:"netmask,omitempty"`
	Nextgenapiresource         string        `json:"_nextgenapiresource,omitempty"`
	Nodegroup                  string        `json:"nodegroup,omitempty"`
	Nodeid                     int           `json:"nodeid,omitempty"`
	Nodejumbonotsupported      bool          `json:"nodejumbonotsupported,omitempty"`
	Nodelicensemismatch        bool          `json:"nodelicensemismatch,omitempty"`
	Nodelist                   []interface{} `json:"nodelist,omitempty"`
	Nodersskeymismatch         bool          `json:"nodersskeymismatch,omitempty"`
	Operationalsyncstate       string        `json:"operationalsyncstate,omitempty"`
	Partialfailifaces          string        `json:"partialfailifaces,omitempty"`
	Priority                   int           `json:"priority,omitempty"`
	Routemonitor               string        `json:"routemonitor,omitempty"`
	State                      string        `json:"state,omitempty"`
	Syncfailurereason          string        `json:"syncfailurereason,omitempty"`
	Syncstate                  string        `json:"syncstate,omitempty"`
	Tunnelmode                 string        `json:"tunnelmode,omitempty"`
}

type ClusternodegroupStreamidentifierBinding struct {
	Identifiername string `json:"identifiername,omitempty"`
	Name           string `json:"name,omitempty"`
}

type Cluster struct {
	Clip     string `json:"clip,omitempty"`
	Password string `json:"password,omitempty"`
}

type Clusterinstance struct {
	Adminstate                 string  `json:"adminstate,omitempty"`
	Backplanebasedview         string  `json:"backplanebasedview,omitempty"`
	Clid                       int     `json:"clid,omitempty"`
	Clusterclipfailure         bool    `json:"clusterclipfailure,omitempty"`
	Clusterhbhmacerrordetected bool    `json:"clusterhbhmacerrordetected,omitempty"`
	Clusternoheartbeatonnode   bool    `json:"clusternoheartbeatonnode,omitempty"`
	Clusternolinksetmbf        bool    `json:"clusternolinksetmbf,omitempty"`
	Clusternospottedip         bool    `json:"clusternospottedip,omitempty"`
	Clusterproxyarp            string  `json:"clusterproxyarp,omitempty"`
	Clustertunnelmodemismatch  bool    `json:"clustertunnelmodemismatch,omitempty"`
	Count                      float64 `json:"__count,omitempty"`
	Deadinterval               int     `json:"deadinterval,omitempty"`
	Dfdretainl2params          string  `json:"dfdretainl2params,omitempty"`
	Hellointerval              int     `json:"hellointerval,omitempty"`
	Heterogeneousflag          string  `json:"heterogeneousflag,omitempty"`
	Inc                        string  `json:"inc,omitempty"`
	Jumbonotsupported          bool    `json:"jumbonotsupported,omitempty"`
	Licensemismatch            bool    `json:"licensemismatch,omitempty"`
	Nextgenapiresource         string  `json:"_nextgenapiresource,omitempty"`
	Nodegroup                  string  `json:"nodegroup,omitempty"`
	Nodegroupstatewarning      bool    `json:"nodegroupstatewarning,omitempty"`
	Nodepenummismatch          bool    `json:"nodepenummismatch,omitempty"`
	Operationalpropstate       string  `json:"operationalpropstate,omitempty"`
	Operationalstate           string  `json:"operationalstate,omitempty"`
	Penummismatch              bool    `json:"penummismatch,omitempty"`
	Preemption                 string  `json:"preemption,omitempty"`
	Processlocal               string  `json:"processlocal,omitempty"`
	Propstate                  string  `json:"propstate,omitempty"`
	Quorumtype                 string  `json:"quorumtype,omitempty"`
	Retainconnectionsoncluster string  `json:"retainconnectionsoncluster,omitempty"`
	Rsskeymismatch             bool    `json:"rsskeymismatch,omitempty"`
	Secureheartbeats           string  `json:"secureheartbeats,omitempty"`
	Status                     string  `json:"status,omitempty"`
	Syncstatusstrictmode       string  `json:"syncstatusstrictmode,omitempty"`
	Validmtu                   int     `json:"validmtu,omitempty"`
}

type ClusternodeRoutemonitorBinding struct {
	Netmask       string `json:"netmask,omitempty"`
	Nodeid        int    `json:"nodeid,omitempty"`
	Routemonitor  string `json:"routemonitor,omitempty"`
	Routemonstate int    `json:"routemonstate,omitempty"`
}

type ClusternodegroupGslbvserverBinding struct {
	Name    string `json:"name,omitempty"`
	Vserver string `json:"vserver,omitempty"`
}

type ClusternodegroupBinding struct {
	ClusternodegroupAuthenticationvserverBinding []interface{} `json:"clusternodegroup_authenticationvserver_binding,omitempty"`
	ClusternodegroupClusternodeBinding           []interface{} `json:"clusternodegroup_clusternode_binding,omitempty"`
	ClusternodegroupCrvserverBinding             []interface{} `json:"clusternodegroup_crvserver_binding,omitempty"`
	ClusternodegroupCsvserverBinding             []interface{} `json:"clusternodegroup_csvserver_binding,omitempty"`
	ClusternodegroupGslbsiteBinding              []interface{} `json:"clusternodegroup_gslbsite_binding,omitempty"`
	ClusternodegroupGslbvserverBinding           []interface{} `json:"clusternodegroup_gslbvserver_binding,omitempty"`
	ClusternodegroupLbvserverBinding             []interface{} `json:"clusternodegroup_lbvserver_binding,omitempty"`
	ClusternodegroupNslimitidentifierBinding     []interface{} `json:"clusternodegroup_nslimitidentifier_binding,omitempty"`
	ClusternodegroupServiceBinding               []interface{} `json:"clusternodegroup_service_binding,omitempty"`
	ClusternodegroupStreamidentifierBinding      []interface{} `json:"clusternodegroup_streamidentifier_binding,omitempty"`
	ClusternodegroupVpnvserverBinding            []interface{} `json:"clusternodegroup_vpnvserver_binding,omitempty"`
	Name                                         string        `json:"name,omitempty"`
}

type ClusternodegroupAuthenticationvserverBinding struct {
	Name    string `json:"name,omitempty"`
	Vserver string `json:"vserver,omitempty"`
}

type ClusternodegroupLbvserverBinding struct {
	Name    string `json:"name,omitempty"`
	Vserver string `json:"vserver,omitempty"`
}

type ClusternodegroupCrvserverBinding struct {
	Name    string `json:"name,omitempty"`
	Vserver string `json:"vserver,omitempty"`
}

type ClusternodegroupNslimitidentifierBinding struct {
	Identifiername string `json:"identifiername,omitempty"`
	Name           string `json:"name,omitempty"`
}

type Clustersync struct {
}

type ClusternodegroupGslbsiteBinding struct {
	Gslbsite string `json:"gslbsite,omitempty"`
	Name     string `json:"name,omitempty"`
}

type ClusternodegroupServiceBinding struct {
	Name    string `json:"name,omitempty"`
	Service string `json:"service,omitempty"`
}

type Clusternodegroup struct {
	Activelist               []interface{} `json:"activelist,omitempty"`
	Backuplist               []interface{} `json:"backuplist,omitempty"`
	Backupnodemask           int           `json:"backupnodemask,omitempty"`
	Boundedentitiescntfrompe int           `json:"boundedentitiescntfrompe,omitempty"`
	Count                    float64       `json:"__count,omitempty"`
	Currentnodemask          int           `json:"currentnodemask,omitempty"`
	Name                     string        `json:"name,omitempty"`
	Nextgenapiresource       string        `json:"_nextgenapiresource,omitempty"`
	Priority                 int           `json:"priority,omitempty"`
	State                    string        `json:"state,omitempty"`
	Sticky                   string        `json:"sticky,omitempty"`
	Strict                   string        `json:"strict,omitempty"`
}

type ClusternodegroupClusternodeBinding struct {
	Name string `json:"name,omitempty"`
	Node int    `json:"node,omitempty"`
}

type Clustersyncfailures struct {
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Clusterpropstatus struct {
	Cmdstrs            string  `json:"cmdstrs,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Nodeid             int     `json:"nodeid,omitempty"`
	Numpropcmdfailed   int     `json:"numpropcmdfailed,omitempty"`
}

type ClusterinstanceBinding struct {
	Clid                              int           `json:"clid,omitempty"`
	ClusterinstanceClusternodeBinding []interface{} `json:"clusterinstance_clusternode_binding,omitempty"`
}

type Clusterfiles struct {
	Mode []string `json:"mode,omitempty"`
}

type ClusterinstanceClusternodeBinding struct {
	Clid                       int    `json:"clid,omitempty"`
	Clusterhealth              string `json:"clusterhealth,omitempty"`
	Effectivestate             string `json:"effectivestate,omitempty"`
	Health                     string `json:"health,omitempty"`
	Ipaddress                  string `json:"ipaddress,omitempty"`
	Isconfigurationcoordinator bool   `json:"isconfigurationcoordinator,omitempty"`
	Islocalnode                bool   `json:"islocalnode,omitempty"`
	Masterstate                string `json:"masterstate,omitempty"`
	Nodeid                     int    `json:"nodeid,omitempty"`
	Nodejumbonotsupported      bool   `json:"nodejumbonotsupported,omitempty"`
	Nodelicensemismatch        bool   `json:"nodelicensemismatch,omitempty"`
	Nodersskeymismatch         bool   `json:"nodersskeymismatch,omitempty"`
	State                      string `json:"state,omitempty"`
}

type ClusternodegroupVpnvserverBinding struct {
	Name    string `json:"name,omitempty"`
	Vserver string `json:"vserver,omitempty"`
}

type ClusternodegroupCsvserverBinding struct {
	Name    string `json:"name,omitempty"`
	Vserver string `json:"vserver,omitempty"`
}
