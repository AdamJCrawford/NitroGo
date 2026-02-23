package models

// transform configuration structs
type TransformPolicyCSVServerBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type TransformProfileBinding struct {
	Name                                   string        `json:"name,omitempty"`
	TransformProfileTransformActionBinding []interface{} `json:"transformprofile_transformaction_binding,omitempty"`
}

type TransformPolicyBinding struct {
	Name                                       string        `json:"name,omitempty"`
	TransformPolicyCSVServerBinding            []interface{} `json:"transformpolicy_csvserver_binding,omitempty"`
	TransformPolicyLBVServerBinding            []interface{} `json:"transformpolicy_lbvserver_binding,omitempty"`
	TransformPolicyTransformGlobalBinding      []interface{} `json:"transformpolicy_transformglobal_binding,omitempty"`
	TransformPolicyTransformPolicyLabelBinding []interface{} `json:"transformpolicy_transformpolicylabel_binding,omitempty"`
}

type TransformPolicyLabel struct {
	Count                  float64 `json:"__count,omitempty"`
	Description            string  `json:"description,omitempty"`
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

type TransformPolicyTransformPolicyLabelBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type TransformPolicyLabelTransformPolicyBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelName        string `json:"invoke_labelname,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type TransformPolicy struct {
	Builtin            []string `json:"builtin,omitempty"`
	Comment            string   `json:"comment,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	IsDefault          bool     `json:"isdefault,omitempty"`
	LogAction          string   `json:"logaction,omitempty"`
	Name               string   `json:"name,omitempty"`
	NewName            string   `json:"newname,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	ProfileName        string   `json:"profilename,omitempty"`
	Rule               string   `json:"rule,omitempty"`
}

type TransformProfile struct {
	AdditionalReqHeadersList       string  `json:"additionalreqheaderslist,omitempty"`
	AdditionalRespHeadersList      string  `json:"additionalrespheaderslist,omitempty"`
	Comment                        string  `json:"comment,omitempty"`
	Count                          float64 `json:"__count,omitempty"`
	Name                           string  `json:"name,omitempty"`
	NextGenAPIResource             string  `json:"_nextgenapiresource,omitempty"`
	OnlyTransformAbsURLInBody      string  `json:"onlytransformabsurlinbody,omitempty"`
	RegexForFindingURLInCSS        string  `json:"regexforfindingurlincss,omitempty"`
	RegexForFindingURLInJavascript string  `json:"regexforfindingurlinjavascript,omitempty"`
	RegexForFindingURLInXComponent string  `json:"regexforfindingurlinxcomponent,omitempty"`
	RegexForFindingURLInXML        string  `json:"regexforfindingurlinxml,omitempty"`
	TypeField                      string  `json:"type,omitempty"`
}

type TransformGlobalTransformPolicyBinding struct {
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

type TransformPolicyLBVServerBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type TransformAction struct {
	Comment            string  `json:"comment,omitempty"`
	ContinueMatching   string  `json:"continuematching,omitempty"`
	CookieDomainFrom   string  `json:"cookiedomainfrom,omitempty"`
	CookieDomainInto   string  `json:"cookiedomaininto,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Priority           int     `json:"priority,omitempty"`
	ProfileName        string  `json:"profilename,omitempty"`
	ReqURLFrom         string  `json:"requrlfrom,omitempty"`
	ReqURLInto         string  `json:"requrlinto,omitempty"`
	ResURLFrom         string  `json:"resurlfrom,omitempty"`
	ResURLInto         string  `json:"resurlinto,omitempty"`
	State              string  `json:"state,omitempty"`
}

type TransformPolicyLabelBinding struct {
	LabelName                                  string        `json:"labelname,omitempty"`
	TransformPolicyLabelPolicyBindingBinding   []interface{} `json:"transformpolicylabel_policybinding_binding,omitempty"`
	TransformPolicyLabelTransformPolicyBinding []interface{} `json:"transformpolicylabel_transformpolicy_binding,omitempty"`
}

type TransformGlobalBinding struct {
	TransformGlobalTransformPolicyBinding []interface{} `json:"transformglobal_transformpolicy_binding,omitempty"`
}

type TransformProfileTransformActionBinding struct {
	ActionComment    string `json:"actioncomment,omitempty"`
	ActionName       string `json:"actionname,omitempty"`
	CookieDomainFrom string `json:"cookiedomainfrom,omitempty"`
	CookieDomainInto string `json:"cookiedomaininto,omitempty"`
	Name             string `json:"name,omitempty"`
	Priority         int    `json:"priority,omitempty"`
	ProfileName      string `json:"profilename,omitempty"`
	ReqURLFrom       string `json:"requrlfrom,omitempty"`
	ReqURLInto       string `json:"requrlinto,omitempty"`
	ResURLFrom       string `json:"resurlfrom,omitempty"`
	ResURLInto       string `json:"resurlinto,omitempty"`
	State            string `json:"state,omitempty"`
}

type TransformPolicyLabelPolicyBindingBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelName        string `json:"invoke_labelname,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type TransformPolicyTransformGlobalBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}
