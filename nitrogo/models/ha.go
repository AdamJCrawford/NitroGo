package models

// ha configuration structs
type HASyncFailures struct {
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	Response           string `json:"response,omitempty"`
}

type HANodeRouteMonitor6Binding struct {
	Flags             int    `json:"flags,omitempty"`
	ID                int    `json:"id,omitempty"`
	Netmask           string `json:"netmask,omitempty"`
	RouteMonitor      string `json:"routemonitor,omitempty"`
	RouteMonitorState string `json:"routemonitorstate,omitempty"`
}

type HANodeRouteMonitorBinding struct {
	Flags             int    `json:"flags,omitempty"`
	ID                int    `json:"id,omitempty"`
	Netmask           string `json:"netmask,omitempty"`
	RouteMonitor      string `json:"routemonitor,omitempty"`
	RouteMonitorState string `json:"routemonitorstate,omitempty"`
}

type HAFiles struct {
	Mode []string `json:"mode,omitempty"`
}

type HANodeCIBinding struct {
	EnaIfaces    string `json:"enaifaces,omitempty"`
	ID           int    `json:"id,omitempty"`
	RouteMonitor string `json:"routemonitor,omitempty"`
}

type HAFailover struct {
	Force bool `json:"force,omitempty"`
}

type HANode struct {
	CompletedFlipTime    int     `json:"completedfliptime,omitempty"`
	Count                float64 `json:"__count,omitempty"`
	CurFlips             int     `json:"curflips,omitempty"`
	DeadInterval         int     `json:"deadinterval,omitempty"`
	DisIfaces            string  `json:"disifaces,omitempty"`
	EnaIfaces            string  `json:"enaifaces,omitempty"`
	FailSafe             string  `json:"failsafe,omitempty"`
	Flags                int     `json:"flags,omitempty"`
	HAHeartbeatIfaces    string  `json:"haheartbeatifaces,omitempty"`
	HAMonIfaces          string  `json:"hamonifaces,omitempty"`
	HAProp               string  `json:"haprop,omitempty"`
	HAStatus             string  `json:"hastatus,omitempty"`
	HASync               string  `json:"hasync,omitempty"`
	HASyncFailureReason  string  `json:"hasyncfailurereason,omitempty"`
	HelloInterval        int     `json:"hellointerval,omitempty"`
	ID                   int     `json:"id,omitempty"`
	Ifaces               string  `json:"ifaces,omitempty"`
	Inc                  string  `json:"inc,omitempty"`
	IPAddress            string  `json:"ipaddress,omitempty"`
	MasterStateTime      int     `json:"masterstatetime,omitempty"`
	MaxFlips             int     `json:"maxflips,omitempty"`
	MaxFlipTime          int     `json:"maxfliptime,omitempty"`
	Name                 string  `json:"name,omitempty"`
	Netmask              string  `json:"netmask,omitempty"`
	NextGenAPIResource   string  `json:"_nextgenapiresource,omitempty"`
	PFIfaces             string  `json:"pfifaces,omitempty"`
	RouteMonitor         string  `json:"routemonitor,omitempty"`
	RouteMonitorState    string  `json:"routemonitorstate,omitempty"`
	RPCNodePassword      string  `json:"rpcnodepassword,omitempty"`
	SSL2                 string  `json:"ssl2,omitempty"`
	State                string  `json:"state,omitempty"`
	SyncStatusStrictMode string  `json:"syncstatusstrictmode,omitempty"`
	SyncVLAN             int     `json:"syncvlan,omitempty"`
}

type HANodeFISBinding struct {
	EnaIfaces    string `json:"enaifaces,omitempty"`
	ID           int    `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	RouteMonitor string `json:"routemonitor,omitempty"`
}

type HANodeBinding struct {
	HANodeCIBinding                       []any `json:"hanode_ci_binding,omitempty"`
	HANodeFISBinding                      []any `json:"hanode_fis_binding,omitempty"`
	HANodePartialFailureInterfacesBinding []any `json:"hanode_partialfailureinterfaces_binding,omitempty"`
	HANodeRouteMonitor6Binding            []any `json:"hanode_routemonitor6_binding,omitempty"`
	HANodeRouteMonitorBinding             []any `json:"hanode_routemonitor_binding,omitempty"`
	ID                                    int   `json:"id,omitempty"`
}

type HANodePartialFailureInterfacesBinding struct {
	ID           int    `json:"id,omitempty"`
	PFIfaces     string `json:"pfifaces,omitempty"`
	RouteMonitor string `json:"routemonitor,omitempty"`
}

type HASync struct {
	Force bool   `json:"force,omitempty"`
	Save  string `json:"save,omitempty"`
}
