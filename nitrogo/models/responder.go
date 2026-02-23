package models

// responder configuration structs
type ResponderGlobalResponderPolicyBinding struct {
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

type ResponderPolicyCSVServerBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type ResponderPolicyResponderGlobalBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type ResponderPolicyResponderPolicyLabelBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type ResponderPolicyCRVServerBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type ResponderPolicyLabelPolicyBindingBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelName        string `json:"invoke_labelname,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type ResponderHTMLPage struct {
	CACertFile         string `json:"cacertfile,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Name               string `json:"name,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	Overwrite          bool   `json:"overwrite,omitempty"`
	Response           string `json:"response,omitempty"`
	Src                string `json:"src,omitempty"`
}

type ResponderAction struct {
	Builtin            []string `json:"builtin,omitempty"`
	BypassSafetyCheck  string   `json:"bypasssafetycheck,omitempty"`
	Comment            string   `json:"comment,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Headers            []string `json:"headers,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	HTMLPage           string   `json:"htmlpage,omitempty"`
	Name               string   `json:"name,omitempty"`
	NewName            string   `json:"newname,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	ReasonPhrase       string   `json:"reasonphrase,omitempty"`
	ReferenceCount     int      `json:"referencecount,omitempty"`
	ResponseStatusCode int      `json:"responsestatuscode,omitempty"`
	Target             string   `json:"target,omitempty"`
	TypeField          string   `json:"type,omitempty"`
	UndefHits          int      `json:"undefhits,omitempty"`
}

type ResponderPolicyLabel struct {
	Comment                string  `json:"comment,omitempty"`
	Count                  float64 `json:"__count,omitempty"`
	GotoPriorityExpression string  `json:"gotopriorityexpression,omitempty"`
	Hits                   int     `json:"hits,omitempty"`
	InvokeLabelName        string  `json:"invoke_labelname,omitempty"`
	LabelName              string  `json:"labelname,omitempty"`
	LabelType              string  `json:"labeltype,omitempty"`
	NewName                string  `json:"newname,omitempty"`
	NextGenAPIResource     string  `json:"_nextgenapiresource,omitempty"`
	NumPol                 int     `json:"numpol,omitempty"`
	PolicyLabelType        string  `json:"policylabeltype,omitempty"`
	Priority               int     `json:"priority,omitempty"`
}

type ResponderPolicyLBVServerBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type ResponderPolicyLabelResponderPolicyBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelName        string `json:"invoke_labelname,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type ResponderPolicyLabelBinding struct {
	LabelName                                  string        `json:"labelname,omitempty"`
	ResponderPolicyLabelPolicyBindingBinding   []interface{} `json:"responderpolicylabel_policybinding_binding,omitempty"`
	ResponderPolicyLabelResponderPolicyBinding []interface{} `json:"responderpolicylabel_responderpolicy_binding,omitempty"`
}

type ResponderPolicy struct {
	Action             string   `json:"action,omitempty"`
	AppFlowAction      string   `json:"appflowaction,omitempty"`
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

type ResponderPolicyVPNVServerBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type ResponderParam struct {
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	Timeout            int    `json:"timeout,omitempty"`
	UndefAction        string `json:"undefaction,omitempty"`
}

type ResponderPolicyBinding struct {
	Name                                       string        `json:"name,omitempty"`
	ResponderPolicyCRVServerBinding            []interface{} `json:"responderpolicy_crvserver_binding,omitempty"`
	ResponderPolicyCSVServerBinding            []interface{} `json:"responderpolicy_csvserver_binding,omitempty"`
	ResponderPolicyLBVServerBinding            []interface{} `json:"responderpolicy_lbvserver_binding,omitempty"`
	ResponderPolicyResponderGlobalBinding      []interface{} `json:"responderpolicy_responderglobal_binding,omitempty"`
	ResponderPolicyResponderPolicyLabelBinding []interface{} `json:"responderpolicy_responderpolicylabel_binding,omitempty"`
	ResponderPolicyVPNVServerBinding           []interface{} `json:"responderpolicy_vpnvserver_binding,omitempty"`
}

type ResponderGlobalBinding struct {
	ResponderGlobalResponderPolicyBinding []interface{} `json:"responderglobal_responderpolicy_binding,omitempty"`
}
