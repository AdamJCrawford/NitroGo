package models

// contentinspection configuration structs
type ContentInspectionPolicyContentInspectionGlobalBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type ContentInspectionPolicyLabel struct {
	Comment                string  `json:"comment,omitempty"`
	Count                  float64 `json:"__count,omitempty"`
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
	TypeField              string  `json:"type,omitempty"`
}

type ContentInspectionWasmProfile struct {
	AnomalousDataSize  int     `json:"anomalousdatasize,omitempty"`
	AnomalousTTFBTime  int     `json:"anomalousttfbtime,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	MaxBodyLen         int     `json:"maxbodylen,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Timeout            int     `json:"timeout,omitempty"`
	TimeoutAction      string  `json:"timeoutaction,omitempty"`
	WasmModule         string  `json:"wasmmodule,omitempty"`
}

type ContentInspectionPolicyLabelPolicyBindingBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelName        string `json:"invoke_labelname,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type ContentInspectionPolicyCSVServerBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type ContentInspectionCallout struct {
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Hits               int     `json:"hits,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	ProfileName        string  `json:"profilename,omitempty"`
	ResultExpr         string  `json:"resultexpr,omitempty"`
	ReturnType         string  `json:"returntype,omitempty"`
	ServerIP           string  `json:"serverip,omitempty"`
	ServerName         string  `json:"servername,omitempty"`
	ServerPort         int     `json:"serverport,omitempty"`
	TypeField          string  `json:"type,omitempty"`
	UndefHits          int     `json:"undefhits,omitempty"`
	UndefReason        string  `json:"undefreason,omitempty"`
}

type ContentInspectionProfile struct {
	Count              float64 `json:"__count,omitempty"`
	EgressInterface    string  `json:"egressinterface,omitempty"`
	EgressVLAN         int     `json:"egressvlan,omitempty"`
	IngressInterface   string  `json:"ingressinterface,omitempty"`
	IngressVLAN        int     `json:"ingressvlan,omitempty"`
	IPTunnel           string  `json:"iptunnel,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	TypeField          string  `json:"type,omitempty"`
}

type ContentInspectionGlobalContentInspectionPolicyBinding struct {
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

type ContentInspectionParameter struct {
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	UndefAction        string `json:"undefaction,omitempty"`
}

type ContentInspectionPolicyLabelBinding struct {
	ContentInspectionPolicyLabelContentInspectionPolicyBinding []any  `json:"contentinspectionpolicylabel_contentinspectionpolicy_binding,omitempty"`
	ContentInspectionPolicyLabelPolicyBindingBinding           []any  `json:"contentinspectionpolicylabel_policybinding_binding,omitempty"`
	LabelName                                                  string `json:"labelname,omitempty"`
}

type ContentInspectionAction struct {
	Builtin            []string `json:"builtin,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	ICAPProfileName    string   `json:"icapprofilename,omitempty"`
	IfServerDown       string   `json:"ifserverdown,omitempty"`
	Name               string   `json:"name,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	ReferenceCount     int      `json:"referencecount,omitempty"`
	ReqTimeout         int      `json:"reqtimeout,omitempty"`
	ReqTimeoutAction   string   `json:"reqtimeoutaction,omitempty"`
	ServerIP           string   `json:"serverip,omitempty"`
	ServerName         string   `json:"servername,omitempty"`
	ServerPort         int      `json:"serverport,omitempty"`
	TypeField          string   `json:"type,omitempty"`
	UndefHits          int      `json:"undefhits,omitempty"`
	WasmProfileName    string   `json:"wasmprofilename,omitempty"`
}

type ContentInspectionPolicyLBVServerBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type ContentInspectionGlobalBinding struct {
	ContentInspectionGlobalContentInspectionPolicyBinding []any `json:"contentinspectionglobal_contentinspectionpolicy_binding,omitempty"`
}

type ContentInspectionPolicyContentInspectionPolicyLabelBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type ContentInspectionPolicyBinding struct {
	ContentInspectionPolicyContentInspectionGlobalBinding      []any  `json:"contentinspectionpolicy_contentinspectionglobal_binding,omitempty"`
	ContentInspectionPolicyContentInspectionPolicyLabelBinding []any  `json:"contentinspectionpolicy_contentinspectionpolicylabel_binding,omitempty"`
	ContentInspectionPolicyCSVServerBinding                    []any  `json:"contentinspectionpolicy_csvserver_binding,omitempty"`
	ContentInspectionPolicyLBVServerBinding                    []any  `json:"contentinspectionpolicy_lbvserver_binding,omitempty"`
	Name                                                       string `json:"name,omitempty"`
}

type ContentInspectionPolicyLabelContentInspectionPolicyBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelName        string `json:"invoke_labelname,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type ContentInspectionPolicy struct {
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
	UndefAction        string   `json:"undefaction,omitempty"`
	UndefHits          int      `json:"undefhits,omitempty"`
}
