package models

// authorization configuration structs
type AuthorizationpolicyAuthorizationpolicylabelBinding struct {
	Boundto  string `json:"boundto,omitempty"`
	Name     string `json:"name,omitempty"`
	Priority int    `json:"priority,omitempty"`
}

type AuthorizationpolicylabelAuthorizationpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type AuthorizationpolicyCsvserverBinding struct {
	Boundto  string `json:"boundto,omitempty"`
	Name     string `json:"name,omitempty"`
	Priority int    `json:"priority,omitempty"`
}

type AuthorizationpolicyBinding struct {
	AuthorizationpolicyAaagroupBinding                 []interface{} `json:"authorizationpolicy_aaagroup_binding,omitempty"`
	AuthorizationpolicyAaauserBinding                  []interface{} `json:"authorizationpolicy_aaauser_binding,omitempty"`
	AuthorizationpolicyAuthorizationpolicylabelBinding []interface{} `json:"authorizationpolicy_authorizationpolicylabel_binding,omitempty"`
	AuthorizationpolicyCsvserverBinding                []interface{} `json:"authorizationpolicy_csvserver_binding,omitempty"`
	AuthorizationpolicyLbvserverBinding                []interface{} `json:"authorizationpolicy_lbvserver_binding,omitempty"`
	Name                                               string        `json:"name,omitempty"`
}

type AuthorizationpolicyLbvserverBinding struct {
	Boundto  string `json:"boundto,omitempty"`
	Name     string `json:"name,omitempty"`
	Priority int    `json:"priority,omitempty"`
}

type Authorizationpolicylabel struct {
	Count                  float64 `json:"__count,omitempty"`
	Description            string  `json:"description,omitempty"`
	Flowtype               int     `json:"flowtype,omitempty"`
	Gotopriorityexpression string  `json:"gotopriorityexpression,omitempty"`
	Hits                   int     `json:"hits,omitempty"`
	InvokeLabelname        string  `json:"invoke_labelname,omitempty"`
	Labelname              string  `json:"labelname,omitempty"`
	Labeltype              string  `json:"labeltype,omitempty"`
	Newname                string  `json:"newname,omitempty"`
	Nextgenapiresource     string  `json:"_nextgenapiresource,omitempty"`
	Numpol                 int     `json:"numpol,omitempty"`
	Policyname             string  `json:"policyname,omitempty"`
	Priority               int     `json:"priority,omitempty"`
}

type AuthorizationpolicyAaagroupBinding struct {
	Boundto  string `json:"boundto,omitempty"`
	Name     string `json:"name,omitempty"`
	Priority int    `json:"priority,omitempty"`
}

type Authorizationpolicy struct {
	Action             string  `json:"action,omitempty"`
	Activepolicy       int     `json:"activepolicy,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Expressiontype     string  `json:"expressiontype,omitempty"`
	Hits               int     `json:"hits,omitempty"`
	Name               string  `json:"name,omitempty"`
	Newname            string  `json:"newname,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Rule               string  `json:"rule,omitempty"`
}

type Authorizationaction struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
}

type AuthorizationpolicylabelBinding struct {
	AuthorizationpolicylabelAuthorizationpolicyBinding []interface{} `json:"authorizationpolicylabel_authorizationpolicy_binding,omitempty"`
	Labelname                                          string        `json:"labelname,omitempty"`
}

type AuthorizationpolicyAaauserBinding struct {
	Boundto  string `json:"boundto,omitempty"`
	Name     string `json:"name,omitempty"`
	Priority int    `json:"priority,omitempty"`
}
