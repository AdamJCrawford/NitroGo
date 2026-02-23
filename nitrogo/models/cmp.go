package models

// cmp configuration structs
type CMPParameter struct {
	AddVaryHeader               string   `json:"addvaryheader,omitempty"`
	Builtin                     []string `json:"builtin,omitempty"`
	CMPBypassPct                int      `json:"cmpbypasspct,omitempty"`
	CMPLevel                    string   `json:"cmplevel,omitempty"`
	CMPOnPush                   string   `json:"cmponpush,omitempty"`
	ExternalCache               string   `json:"externalcache,omitempty"`
	Feature                     string   `json:"feature,omitempty"`
	HeurExpiry                  string   `json:"heurexpiry,omitempty"`
	HeurExpiryHistWt            int      `json:"heurexpiryhistwt,omitempty"`
	HeurExpiryThres             int      `json:"heurexpirythres,omitempty"`
	MinResSize                  int      `json:"minressize,omitempty"`
	NextGenAPIResource          string   `json:"_nextgenapiresource,omitempty"`
	PolicyType                  string   `json:"policytype,omitempty"`
	QuantumSize                 int      `json:"quantumsize,omitempty"`
	RandomGzipFilename          string   `json:"randomgzipfilename,omitempty"`
	RandomGzipFilenameMaxLength int      `json:"randomgzipfilenamemaxlength,omitempty"`
	RandomGzipFilenameMinLength int      `json:"randomgzipfilenameminlength,omitempty"`
	ServerCMP                   string   `json:"servercmp,omitempty"`
	VaryHeaderValue             string   `json:"varyheadervalue,omitempty"`
}

type CMPGlobalBinding struct {
	CMPGlobalCMPPolicyBinding []interface{} `json:"cmpglobal_cmppolicy_binding,omitempty"`
}

type CMPPolicyBinding struct {
	CMPPolicyCMPGlobalBinding      []interface{} `json:"cmppolicy_cmpglobal_binding,omitempty"`
	CMPPolicyCMPPolicyLabelBinding []interface{} `json:"cmppolicy_cmppolicylabel_binding,omitempty"`
	CMPPolicyCRVServerBinding      []interface{} `json:"cmppolicy_crvserver_binding,omitempty"`
	CMPPolicyCSVServerBinding      []interface{} `json:"cmppolicy_csvserver_binding,omitempty"`
	CMPPolicyLBVServerBinding      []interface{} `json:"cmppolicy_lbvserver_binding,omitempty"`
	Name                           string        `json:"name,omitempty"`
}

type CMPPolicyLabelPolicyBindingBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelName        string `json:"invoke_labelname,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type CMPPolicyCMPPolicyLabelBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type CMPGlobalCMPPolicyBinding struct {
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

type CMPPolicy struct {
	Builtin            []string `json:"builtin,omitempty"`
	ClientTransactions int      `json:"clienttransactions,omitempty"`
	ClientTTLB         int      `json:"clientttlb,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Description        string   `json:"description,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	IsDefault          bool     `json:"isdefault,omitempty"`
	Name               string   `json:"name,omitempty"`
	NewName            string   `json:"newname,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	ReqAction          string   `json:"reqaction,omitempty"`
	ResAction          string   `json:"resaction,omitempty"`
	Rule               string   `json:"rule,omitempty"`
	RxBytes            int      `json:"rxbytes,omitempty"`
	ServerTransactions int      `json:"servertransactions,omitempty"`
	ServerTTLB         int      `json:"serverttlb,omitempty"`
	TxBytes            int      `json:"txbytes,omitempty"`
}

type CMPAction struct {
	AddVaryHeader      string   `json:"addvaryheader,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	CMPType            string   `json:"cmptype,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	DeltaType          string   `json:"deltatype,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	IsDefault          bool     `json:"isdefault,omitempty"`
	Name               string   `json:"name,omitempty"`
	NewName            string   `json:"newname,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	VaryHeaderValue    string   `json:"varyheadervalue,omitempty"`
}

type CMPPolicyCMPGlobalBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type CMPPolicyLabel struct {
	Count                  float64 `json:"__count,omitempty"`
	Description            string  `json:"description,omitempty"`
	FlowType               int     `json:"flowtype,omitempty"`
	GotoPriorityExpression string  `json:"gotopriorityexpression,omitempty"`
	Hits                   int     `json:"hits,omitempty"`
	InvokeLabelName        string  `json:"invoke_labelname,omitempty"`
	LabelName              string  `json:"labelname,omitempty"`
	LabelType              string  `json:"labeltype,omitempty"`
	NewName                string  `json:"newname,omitempty"`
	NextGenAPIResource     string  `json:"_nextgenapiresource,omitempty"`
	NumPol                 int     `json:"numpol,omitempty"`
	Priority               int     `json:"priority,omitempty"`
	TypeField              string  `json:"type,omitempty"`
}

type CMPPolicyCRVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type CMPPolicyLBVServerBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type CMPPolicyLabelCMPPolicyBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelName        string `json:"invoke_labelname,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type CMPPolicyCSVServerBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type CMPPolicyLabelBinding struct {
	CMPPolicyLabelCMPPolicyBinding     []interface{} `json:"cmppolicylabel_cmppolicy_binding,omitempty"`
	CMPPolicyLabelPolicyBindingBinding []interface{} `json:"cmppolicylabel_policybinding_binding,omitempty"`
	LabelName                          string        `json:"labelname,omitempty"`
}
