package models

// transform configuration structs
type TransformpolicyCsvserverBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type TransformprofileBinding struct {
	Name                                   string        `json:"name,omitempty"`
	TransformprofileTransformactionBinding []interface{} `json:"transformprofile_transformaction_binding,omitempty"`
}

type TransformpolicyBinding struct {
	Name                                       string        `json:"name,omitempty"`
	TransformpolicyCsvserverBinding            []interface{} `json:"transformpolicy_csvserver_binding,omitempty"`
	TransformpolicyLbvserverBinding            []interface{} `json:"transformpolicy_lbvserver_binding,omitempty"`
	TransformpolicyTransformglobalBinding      []interface{} `json:"transformpolicy_transformglobal_binding,omitempty"`
	TransformpolicyTransformpolicylabelBinding []interface{} `json:"transformpolicy_transformpolicylabel_binding,omitempty"`
}

type Transformpolicylabel struct {
	Count                  float64 `json:"__count,omitempty"`
	Description            string  `json:"description,omitempty"`
	Gotopriorityexpression string  `json:"gotopriorityexpression,omitempty"`
	Hits                   int     `json:"hits,omitempty"`
	InvokeLabelname        string  `json:"invoke_labelname,omitempty"`
	Labelname              string  `json:"labelname,omitempty"`
	Labeltype              string  `json:"labeltype,omitempty"`
	Newname                string  `json:"newname,omitempty"`
	Nextgenapiresource     string  `json:"_nextgenapiresource,omitempty"`
	Numpol                 int     `json:"numpol,omitempty"`
	Policylabeltype        string  `json:"policylabeltype,omitempty"`
	Priority               int     `json:"priority,omitempty"`
}

type TransformpolicyTransformpolicylabelBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type TransformpolicylabelTransformpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Transformpolicy struct {
	Builtin            []string `json:"builtin,omitempty"`
	Comment            string   `json:"comment,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	Isdefault          bool     `json:"isdefault,omitempty"`
	Logaction          string   `json:"logaction,omitempty"`
	Name               string   `json:"name,omitempty"`
	Newname            string   `json:"newname,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Profilename        string   `json:"profilename,omitempty"`
	Rule               string   `json:"rule,omitempty"`
}

type Transformprofile struct {
	Additionalreqheaderslist       string  `json:"additionalreqheaderslist,omitempty"`
	Additionalrespheaderslist      string  `json:"additionalrespheaderslist,omitempty"`
	Comment                        string  `json:"comment,omitempty"`
	Count                          float64 `json:"__count,omitempty"`
	Name                           string  `json:"name,omitempty"`
	Nextgenapiresource             string  `json:"_nextgenapiresource,omitempty"`
	Onlytransformabsurlinbody      string  `json:"onlytransformabsurlinbody,omitempty"`
	Regexforfindingurlincss        string  `json:"regexforfindingurlincss,omitempty"`
	Regexforfindingurlinjavascript string  `json:"regexforfindingurlinjavascript,omitempty"`
	Regexforfindingurlinxcomponent string  `json:"regexforfindingurlinxcomponent,omitempty"`
	Regexforfindingurlinxml        string  `json:"regexforfindingurlinxml,omitempty"`
	TypeField                      string  `json:"type,omitempty"`
}

type TransformglobalTransformpolicyBinding struct {
	Flowtype               int    `json:"flowtype,omitempty"`
	Globalbindtype         string `json:"globalbindtype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Numpol                 int    `json:"numpol,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
}

type TransformpolicyLbvserverBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Transformaction struct {
	Comment            string  `json:"comment,omitempty"`
	Continuematching   string  `json:"continuematching,omitempty"`
	Cookiedomainfrom   string  `json:"cookiedomainfrom,omitempty"`
	Cookiedomaininto   string  `json:"cookiedomaininto,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Priority           int     `json:"priority,omitempty"`
	Profilename        string  `json:"profilename,omitempty"`
	Requrlfrom         string  `json:"requrlfrom,omitempty"`
	Requrlinto         string  `json:"requrlinto,omitempty"`
	Resurlfrom         string  `json:"resurlfrom,omitempty"`
	Resurlinto         string  `json:"resurlinto,omitempty"`
	State              string  `json:"state,omitempty"`
}

type TransformpolicylabelBinding struct {
	Labelname                                  string        `json:"labelname,omitempty"`
	TransformpolicylabelPolicybindingBinding   []interface{} `json:"transformpolicylabel_policybinding_binding,omitempty"`
	TransformpolicylabelTransformpolicyBinding []interface{} `json:"transformpolicylabel_transformpolicy_binding,omitempty"`
}

type TransformglobalBinding struct {
	TransformglobalTransformpolicyBinding []interface{} `json:"transformglobal_transformpolicy_binding,omitempty"`
}

type TransformprofileTransformactionBinding struct {
	Actioncomment    string `json:"actioncomment,omitempty"`
	Actionname       string `json:"actionname,omitempty"`
	Cookiedomainfrom string `json:"cookiedomainfrom,omitempty"`
	Cookiedomaininto string `json:"cookiedomaininto,omitempty"`
	Name             string `json:"name,omitempty"`
	Priority         int    `json:"priority,omitempty"`
	Profilename      string `json:"profilename,omitempty"`
	Requrlfrom       string `json:"requrlfrom,omitempty"`
	Requrlinto       string `json:"requrlinto,omitempty"`
	Resurlfrom       string `json:"resurlfrom,omitempty"`
	Resurlinto       string `json:"resurlinto,omitempty"`
	State            string `json:"state,omitempty"`
}

type TransformpolicylabelPolicybindingBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type TransformpolicyTransformglobalBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}
