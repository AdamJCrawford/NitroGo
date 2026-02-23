package models

// cr configuration structs
type CRVServerAppQOEPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TargetVServer          string `json:"targetvserver,omitempty"`
}

type CRVServerCRPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Hits                   int    `json:"hits,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	PIPolicyHits           int    `json:"pipolicyhits,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TargetVServer          string `json:"targetvserver,omitempty"`
}

type CRVServerCachePolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TargetVServer          string `json:"targetvserver,omitempty"`
}

type CRPolicyBinding struct {
	CRPolicyCRVServerBinding []interface{} `json:"crpolicy_crvserver_binding,omitempty"`
	PolicyName               string        `json:"policyname,omitempty"`
}

type CRVServerCMPPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Inherited              string `json:"inherited,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TargetVServer          string `json:"targetvserver,omitempty"`
}

type CRVServerAnalyticsProfileBinding struct {
	AnalyticsProfile string `json:"analyticsprofile,omitempty"`
	Name             string `json:"name,omitempty"`
}

type CRVServerLBVServerBinding struct {
	Hits      int    `json:"hits,omitempty"`
	LBVServer string `json:"lbvserver,omitempty"`
	Name      string `json:"name,omitempty"`
}

type CRVServerAppFlowPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TargetVServer          string `json:"targetvserver,omitempty"`
}

type CRVServerFEOPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TargetVServer          string `json:"targetvserver,omitempty"`
}

type CRVServerICAPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TargetVServer          string `json:"targetvserver,omitempty"`
}

type CRVServerResponderPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TargetVServer          string `json:"targetvserver,omitempty"`
}

type CRAction struct {
	Builtin            []string `json:"builtin,omitempty"`
	Comment            string   `json:"comment,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	CRType             string   `json:"crtype,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	IsDefault          bool     `json:"isdefault,omitempty"`
	Name               string   `json:"name,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	ReferenceCount     int      `json:"referencecount,omitempty"`
	UndefHits          int      `json:"undefhits,omitempty"`
}

type CRVServerAppFWPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TargetVServer          string `json:"targetvserver,omitempty"`
}

type CRVServerPolicyMapBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TargetVServer          string `json:"targetvserver,omitempty"`
}

type CRVServerRewritePolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TargetVServer          string `json:"targetvserver,omitempty"`
}

type CRVServerSpilloverPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TargetVServer          string `json:"targetvserver,omitempty"`
}

type CRPolicy struct {
	Action             string   `json:"action,omitempty"`
	ActivePolicy       bool     `json:"activepolicy,omitempty"`
	BoundTo            string   `json:"boundto,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	IsDefault          bool     `json:"isdefault,omitempty"`
	LabelName          string   `json:"labelname,omitempty"`
	LabelType          string   `json:"labeltype,omitempty"`
	LogAction          string   `json:"logaction,omitempty"`
	NewName            string   `json:"newname,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	PolicyName         string   `json:"policyname,omitempty"`
	Priority           int      `json:"priority,omitempty"`
	Rule               string   `json:"rule,omitempty"`
	VSType             int      `json:"vstype,omitempty"`
}

type CRVServerBinding struct {
	CRVServerAnalyticsProfileBinding []interface{} `json:"crvserver_analyticsprofile_binding,omitempty"`
	CRVServerAppFlowPolicyBinding    []interface{} `json:"crvserver_appflowpolicy_binding,omitempty"`
	CRVServerAppFWPolicyBinding      []interface{} `json:"crvserver_appfwpolicy_binding,omitempty"`
	CRVServerAppQOEPolicyBinding     []interface{} `json:"crvserver_appqoepolicy_binding,omitempty"`
	CRVServerCachePolicyBinding      []interface{} `json:"crvserver_cachepolicy_binding,omitempty"`
	CRVServerCMPPolicyBinding        []interface{} `json:"crvserver_cmppolicy_binding,omitempty"`
	CRVServerCRPolicyBinding         []interface{} `json:"crvserver_crpolicy_binding,omitempty"`
	CRVServerCSPolicyBinding         []interface{} `json:"crvserver_cspolicy_binding,omitempty"`
	CRVServerFEOPolicyBinding        []interface{} `json:"crvserver_feopolicy_binding,omitempty"`
	CRVServerICAPolicyBinding        []interface{} `json:"crvserver_icapolicy_binding,omitempty"`
	CRVServerLBVServerBinding        []interface{} `json:"crvserver_lbvserver_binding,omitempty"`
	CRVServerPolicyMapBinding        []interface{} `json:"crvserver_policymap_binding,omitempty"`
	CRVServerResponderPolicyBinding  []interface{} `json:"crvserver_responderpolicy_binding,omitempty"`
	CRVServerRewritePolicyBinding    []interface{} `json:"crvserver_rewritepolicy_binding,omitempty"`
	CRVServerSpilloverPolicyBinding  []interface{} `json:"crvserver_spilloverpolicy_binding,omitempty"`
	Name                             string        `json:"name,omitempty"`
}

type CRVServer struct {
	AppFlowLog               string  `json:"appflowlog,omitempty"`
	ARP                      string  `json:"arp,omitempty"`
	Authentication           string  `json:"authentication,omitempty"`
	BackendSSL               string  `json:"backendssl,omitempty"`
	BackupVServer            string  `json:"backupvserver,omitempty"`
	BindPoint                string  `json:"bindpoint,omitempty"`
	CacheType                string  `json:"cachetype,omitempty"`
	CacheVServer             string  `json:"cachevserver,omitempty"`
	CltTimeout               int     `json:"clttimeout,omitempty"`
	Comment                  string  `json:"comment,omitempty"`
	Count                    float64 `json:"__count,omitempty"`
	CurState                 string  `json:"curstate,omitempty"`
	DestinationVServer       string  `json:"destinationvserver,omitempty"`
	DisablePrimaryOnDown     string  `json:"disableprimaryondown,omitempty"`
	DisallowServiceAccess    string  `json:"disallowserviceaccess,omitempty"`
	DNSVServerName           string  `json:"dnsvservername,omitempty"`
	Domain                   string  `json:"domain,omitempty"`
	DownStateFlush           string  `json:"downstateflush,omitempty"`
	Format                   string  `json:"format,omitempty"`
	Ghost                    string  `json:"ghost,omitempty"`
	GotoPriorityExpression   string  `json:"gotopriorityexpression,omitempty"`
	HomePage                 string  `json:"homepage,omitempty"`
	HTTPProfileName          string  `json:"httpprofilename,omitempty"`
	ICMPVSRResponse          string  `json:"icmpvsrresponse,omitempty"`
	Invoke                   bool    `json:"invoke,omitempty"`
	IP                       string  `json:"ip,omitempty"`
	IPSet                    string  `json:"ipset,omitempty"`
	IPv46                    string  `json:"ipv46,omitempty"`
	L2Conn                   string  `json:"l2conn,omitempty"`
	LabelName                string  `json:"labelname,omitempty"`
	LabelType                string  `json:"labeltype,omitempty"`
	LBVServer                string  `json:"lbvserver,omitempty"`
	ListenPolicy             string  `json:"listenpolicy,omitempty"`
	ListenPriority           int     `json:"listenpriority,omitempty"`
	MapField                 string  `json:"map,omitempty"`
	Name                     string  `json:"name,omitempty"`
	NetProfile               string  `json:"netprofile,omitempty"`
	NewName                  string  `json:"newname,omitempty"`
	NextGenAPIResource       string  `json:"_nextgenapiresource,omitempty"`
	NGName                   string  `json:"ngname,omitempty"`
	NoDefaultBindings        string  `json:"nodefaultbindings,omitempty"`
	OnPolicyMatch            string  `json:"onpolicymatch,omitempty"`
	OriginUSIP               string  `json:"originusip,omitempty"`
	PIPolicyHits             int     `json:"pipolicyhits,omitempty"`
	PolicyName               string  `json:"policyname,omitempty"`
	Port                     int     `json:"port,omitempty"`
	Precedence               string  `json:"precedence,omitempty"`
	Priority                 int     `json:"priority,omitempty"`
	ProbePort                int     `json:"probeport,omitempty"`
	ProbeProtocol            string  `json:"probeprotocol,omitempty"`
	ProbeSuccessResponseCode string  `json:"probesuccessresponsecode,omitempty"`
	Range                    int     `json:"range,omitempty"`
	Redirect                 string  `json:"redirect,omitempty"`
	RedirectURL              string  `json:"redirecturl,omitempty"`
	Reuse                    string  `json:"reuse,omitempty"`
	RHIState                 string  `json:"rhistate,omitempty"`
	Rule                     string  `json:"rule,omitempty"`
	ServiceName              string  `json:"servicename,omitempty"`
	ServiceType              string  `json:"servicetype,omitempty"`
	SOMethod                 string  `json:"somethod,omitempty"`
	SOPersistence            string  `json:"sopersistence,omitempty"`
	SOPersistenceTimeout     int     `json:"sopersistencetimeout,omitempty"`
	SOThreshold              int     `json:"sothreshold,omitempty"`
	SrcIPExpr                string  `json:"srcipexpr,omitempty"`
	State                    string  `json:"state,omitempty"`
	Status                   int     `json:"status,omitempty"`
	TargetVServer            string  `json:"targetvserver,omitempty"`
	TCPProbePort             int     `json:"tcpprobeport,omitempty"`
	TCPProfileName           string  `json:"tcpprofilename,omitempty"`
	TD                       int     `json:"td,omitempty"`
	TypeField                string  `json:"type,omitempty"`
	UseOriginIPPortForCache  string  `json:"useoriginipportforcache,omitempty"`
	UsePortRange             string  `json:"useportrange,omitempty"`
	Value                    string  `json:"value,omitempty"`
	Via                      string  `json:"via,omitempty"`
	Weight                   int     `json:"weight,omitempty"`
}

type CRPolicyCRVServerBinding struct {
	BindHits               int    `json:"bindhits,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Hits                   int    `json:"hits,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type CRVServerCSPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Hits                   int    `json:"hits,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	PIPolicyHits           int    `json:"pipolicyhits,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TargetVServer          string `json:"targetvserver,omitempty"`
}
