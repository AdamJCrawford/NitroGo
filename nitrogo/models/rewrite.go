package models

// rewrite configuration structs
type RewriteGlobalBinding struct {
	RewriteGlobalRewritePolicyBinding []interface{} `json:"rewriteglobal_rewritepolicy_binding,omitempty"`
}

type RewritePolicy struct {
	Action             string   `json:"action,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	Comment            string   `json:"comment,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Description        string   `json:"description,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	IsDefault          bool     `json:"isdefault,omitempty"`
	LogAction          string   `json:"logaction,omitempty"`
	Name               string   `json:"name,omitempty"`
	NewName            string   `json:"newname,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	Rule               string   `json:"rule,omitempty"`
	UndefAction        string   `json:"undefaction,omitempty"`
	UndefHits          int      `json:"undefhits,omitempty"`
}

type RewriteAction struct {
	Builtin            []string `json:"builtin,omitempty"`
	Comment            string   `json:"comment,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Description        string   `json:"description,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	IsDefault          bool     `json:"isdefault,omitempty"`
	Name               string   `json:"name,omitempty"`
	NewName            string   `json:"newname,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	ReferenceCount     int      `json:"referencecount,omitempty"`
	RefineSearch       string   `json:"refinesearch,omitempty"`
	Search             string   `json:"search,omitempty"`
	StringBuilderExpr  string   `json:"stringbuilderexpr,omitempty"`
	Target             string   `json:"target,omitempty"`
	TypeField          string   `json:"type,omitempty"`
	UndefHits          int      `json:"undefhits,omitempty"`
}

type RewritePolicyVPNVServerBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type RewritePolicyLBVServerBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type RewritePolicyLabelRewritePolicyBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelName        string `json:"invoke_labelname,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type RewritePolicyLabelBinding struct {
	LabelName                              string        `json:"labelname,omitempty"`
	RewritePolicyLabelPolicyBindingBinding []interface{} `json:"rewritepolicylabel_policybinding_binding,omitempty"`
	RewritePolicyLabelRewritePolicyBinding []interface{} `json:"rewritepolicylabel_rewritepolicy_binding,omitempty"`
}

type RewritePolicyCSVServerBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type RewritePolicyRewriteGlobalBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type RewritePolicyBinding struct {
	Name                                   string        `json:"name,omitempty"`
	RewritePolicyCSVServerBinding          []interface{} `json:"rewritepolicy_csvserver_binding,omitempty"`
	RewritePolicyLBVServerBinding          []interface{} `json:"rewritepolicy_lbvserver_binding,omitempty"`
	RewritePolicyRewriteGlobalBinding      []interface{} `json:"rewritepolicy_rewriteglobal_binding,omitempty"`
	RewritePolicyRewritePolicyLabelBinding []interface{} `json:"rewritepolicy_rewritepolicylabel_binding,omitempty"`
	RewritePolicyVPNVServerBinding         []interface{} `json:"rewritepolicy_vpnvserver_binding,omitempty"`
}

type RewriteParam struct {
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	Timeout            int    `json:"timeout,omitempty"`
	UndefAction        string `json:"undefaction,omitempty"`
}

type RewritePolicyLabel struct {
	Builtin                []string `json:"builtin,omitempty"`
	Comment                string   `json:"comment,omitempty"`
	Count                  float64  `json:"__count,omitempty"`
	Description            string   `json:"description,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	FlowType               int      `json:"flowtype,omitempty"`
	GotoPriorityExpression string   `json:"gotopriorityexpression,omitempty"`
	Hits                   int      `json:"hits,omitempty"`
	InvokeLabelName        string   `json:"invoke_labelname,omitempty"`
	IsDefault              bool     `json:"isdefault,omitempty"`
	LabelName              string   `json:"labelname,omitempty"`
	LabelType              string   `json:"labeltype,omitempty"`
	NewName                string   `json:"newname,omitempty"`
	NextGenAPIResource     string   `json:"_nextgenapiresource,omitempty"`
	NumPol                 int      `json:"numpol,omitempty"`
	Priority               int      `json:"priority,omitempty"`
	Transform              string   `json:"transform,omitempty"`
}

type RewritePolicyRewritePolicyLabelBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type RewriteGlobalRewritePolicyBinding struct {
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

type RewritePolicyLabelPolicyBindingBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelName        string `json:"invoke_labelname,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}
