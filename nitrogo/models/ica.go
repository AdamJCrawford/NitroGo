package models

// ica configuration structs
type ICAPolicyCRVServerBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type ICAGlobalICAPolicyBinding struct {
	FlowType               int    `json:"flowtype,omitempty"`
	GlobalBindType         string `json:"globalbindtype,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	NumPol                 int    `json:"numpol,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
}

type ICALatencyProfile struct {
	Builtin                  []string `json:"builtin,omitempty"`
	Count                    float64  `json:"__count,omitempty"`
	Feature                  string   `json:"feature,omitempty"`
	IsDefault                bool     `json:"isdefault,omitempty"`
	L7LatencyMaxNotifyCount  int      `json:"l7latencymaxnotifycount,omitempty"`
	L7LatencyMonitoring      string   `json:"l7latencymonitoring,omitempty"`
	L7LatencyNotifyInterval  int      `json:"l7latencynotifyinterval,omitempty"`
	L7LatencyThresholdFactor int      `json:"l7latencythresholdfactor,omitempty"`
	L7LatencyWaitTime        int      `json:"l7latencywaittime,omitempty"`
	Name                     string   `json:"name,omitempty"`
	NextGenAPIResource       string   `json:"_nextgenapiresource,omitempty"`
	RefCnt                   int      `json:"refcnt,omitempty"`
}

type ICAAction struct {
	AccessProfileName  string   `json:"accessprofilename,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	IsDefault          bool     `json:"isdefault,omitempty"`
	LatencyProfileName string   `json:"latencyprofilename,omitempty"`
	Name               string   `json:"name,omitempty"`
	NewName            string   `json:"newname,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	ReferenceCount     int      `json:"referencecount,omitempty"`
	UndefHits          int      `json:"undefhits,omitempty"`
}

type ICAPolicyVPNVServerBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type ICAGlobalBinding struct {
	ICAGlobalICAPolicyBinding []interface{} `json:"icaglobal_icapolicy_binding,omitempty"`
}

type ICAParameter struct {
	Builtin               []string `json:"builtin,omitempty"`
	DFPersistence         string   `json:"dfpersistence,omitempty"`
	EDTLossTolerant       string   `json:"edtlosstolerant,omitempty"`
	EDTPMTUDDF            string   `json:"edtpmtuddf,omitempty"`
	EDTPMTUDDFTimeout     int      `json:"edtpmtuddftimeout,omitempty"`
	EDTPMTUDRediscovery   string   `json:"edtpmtudrediscovery,omitempty"`
	EnableSROnHAFailover  string   `json:"enablesronhafailover,omitempty"`
	HDXInsightNonNSAP     string   `json:"hdxinsightnonnsap,omitempty"`
	InsightOnlyToDirector string   `json:"insightonlytodirector,omitempty"`
	L7LatencyFrequency    int      `json:"l7latencyfrequency,omitempty"`
	NextGenAPIResource    string   `json:"_nextgenapiresource,omitempty"`
}

type ICAPolicyICAGlobalBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type ICAPolicy struct {
	Action             string   `json:"action,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	Comment            string   `json:"comment,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	LogAction          string   `json:"logaction,omitempty"`
	Name               string   `json:"name,omitempty"`
	NewName            string   `json:"newname,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	Rule               string   `json:"rule,omitempty"`
	UndefHits          int      `json:"undefhits,omitempty"`
}

type ICAPolicyBinding struct {
	ICAPolicyCRVServerBinding  []interface{} `json:"icapolicy_crvserver_binding,omitempty"`
	ICAPolicyICAGlobalBinding  []interface{} `json:"icapolicy_icaglobal_binding,omitempty"`
	ICAPolicyVPNVServerBinding []interface{} `json:"icapolicy_vpnvserver_binding,omitempty"`
	Name                       string        `json:"name,omitempty"`
}

type ICAAccessProfile struct {
	Builtin                      []string `json:"builtin,omitempty"`
	ClientAudioRedirection       string   `json:"clientaudioredirection,omitempty"`
	ClientClipboardRedirection   string   `json:"clientclipboardredirection,omitempty"`
	ClientCOMPortRedirection     string   `json:"clientcomportredirection,omitempty"`
	ClientDriveRedirection       string   `json:"clientdriveredirection,omitempty"`
	ClientPrinterRedirection     string   `json:"clientprinterredirection,omitempty"`
	ClientTwainDeviceRedirection string   `json:"clienttwaindeviceredirection,omitempty"`
	ClientUSBDriveRedirection    string   `json:"clientusbdriveredirection,omitempty"`
	ConnectClientLPTPorts        string   `json:"connectclientlptports,omitempty"`
	Count                        float64  `json:"__count,omitempty"`
	DragAndDrop                  string   `json:"draganddrop,omitempty"`
	Feature                      string   `json:"feature,omitempty"`
	FIDO2Redirection             string   `json:"fido2redirection,omitempty"`
	IsDefault                    bool     `json:"isdefault,omitempty"`
	LocalRemoteDataSharing       string   `json:"localremotedatasharing,omitempty"`
	MultiStream                  string   `json:"multistream,omitempty"`
	Name                         string   `json:"name,omitempty"`
	NextGenAPIResource           string   `json:"_nextgenapiresource,omitempty"`
	RefCnt                       int      `json:"refcnt,omitempty"`
	SmartCardRedirection         string   `json:"smartcardredirection,omitempty"`
	WIARedirection               string   `json:"wiaredirection,omitempty"`
}
