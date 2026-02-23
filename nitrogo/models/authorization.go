package models

// authorization configuration structs
type AuthorizationPolicyAuthorizationPolicyLabelBinding struct {
	BoundTo  string `json:"boundto,omitempty"`
	Name     string `json:"name,omitempty"`
	Priority int    `json:"priority,omitempty"`
}

type AuthorizationPolicyLabelAuthorizationPolicyBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelName        string `json:"invoke_labelname,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type AuthorizationPolicyCSVServerBinding struct {
	BoundTo  string `json:"boundto,omitempty"`
	Name     string `json:"name,omitempty"`
	Priority int    `json:"priority,omitempty"`
}

type AuthorizationPolicyBinding struct {
	AuthorizationPolicyAAAGroupBinding                 []interface{} `json:"authorizationpolicy_aaagroup_binding,omitempty"`
	AuthorizationPolicyAAAUserBinding                  []interface{} `json:"authorizationpolicy_aaauser_binding,omitempty"`
	AuthorizationPolicyAuthorizationPolicyLabelBinding []interface{} `json:"authorizationpolicy_authorizationpolicylabel_binding,omitempty"`
	AuthorizationPolicyCSVServerBinding                []interface{} `json:"authorizationpolicy_csvserver_binding,omitempty"`
	AuthorizationPolicyLBVServerBinding                []interface{} `json:"authorizationpolicy_lbvserver_binding,omitempty"`
	Name                                               string        `json:"name,omitempty"`
}

type AuthorizationPolicyLBVServerBinding struct {
	BoundTo  string `json:"boundto,omitempty"`
	Name     string `json:"name,omitempty"`
	Priority int    `json:"priority,omitempty"`
}

type AuthorizationPolicyLabel struct {
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
	PolicyName             string  `json:"policyname,omitempty"`
	Priority               int     `json:"priority,omitempty"`
}

type AuthorizationPolicyAAAGroupBinding struct {
	BoundTo  string `json:"boundto,omitempty"`
	Name     string `json:"name,omitempty"`
	Priority int    `json:"priority,omitempty"`
}

type AuthorizationPolicy struct {
	Action             string  `json:"action,omitempty"`
	ActivePolicy       int     `json:"activepolicy,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	ExpressionType     string  `json:"expressiontype,omitempty"`
	Hits               int     `json:"hits,omitempty"`
	Name               string  `json:"name,omitempty"`
	NewName            string  `json:"newname,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Rule               string  `json:"rule,omitempty"`
}

type AuthorizationAction struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
}

type AuthorizationPolicyLabelBinding struct {
	AuthorizationPolicyLabelAuthorizationPolicyBinding []interface{} `json:"authorizationpolicylabel_authorizationpolicy_binding,omitempty"`
	LabelName                                          string        `json:"labelname,omitempty"`
}

type AuthorizationPolicyAAAUserBinding struct {
	BoundTo  string `json:"boundto,omitempty"`
	Name     string `json:"name,omitempty"`
	Priority int    `json:"priority,omitempty"`
}
