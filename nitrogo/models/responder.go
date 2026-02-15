package models

// responder configuration structs
type ResponderglobalResponderpolicyBinding struct {
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

type ResponderpolicyCsvserverBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type ResponderpolicyResponderglobalBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type ResponderpolicyResponderpolicylabelBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type ResponderpolicyCrvserverBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type ResponderpolicylabelPolicybindingBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Responderhtmlpage struct {
	Cacertfile         string `json:"cacertfile,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Name               string `json:"name,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Overwrite          bool   `json:"overwrite,omitempty"`
	Response           string `json:"response,omitempty"`
	Src                string `json:"src,omitempty"`
}

type Responderaction struct {
	Builtin            []string `json:"builtin,omitempty"`
	Bypasssafetycheck  string   `json:"bypasssafetycheck,omitempty"`
	Comment            string   `json:"comment,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Headers            []string `json:"headers,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	Htmlpage           string   `json:"htmlpage,omitempty"`
	Name               string   `json:"name,omitempty"`
	Newname            string   `json:"newname,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Reasonphrase       string   `json:"reasonphrase,omitempty"`
	Referencecount     int      `json:"referencecount,omitempty"`
	Responsestatuscode int      `json:"responsestatuscode,omitempty"`
	Target             string   `json:"target,omitempty"`
	TypeField          string   `json:"type,omitempty"`
	Undefhits          int      `json:"undefhits,omitempty"`
}

type Responderpolicylabel struct {
	Comment                string  `json:"comment,omitempty"`
	Count                  float64 `json:"__count,omitempty"`
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

type ResponderpolicyLbvserverBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type ResponderpolicylabelResponderpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type ResponderpolicylabelBinding struct {
	Labelname                                  string        `json:"labelname,omitempty"`
	ResponderpolicylabelPolicybindingBinding   []interface{} `json:"responderpolicylabel_policybinding_binding,omitempty"`
	ResponderpolicylabelResponderpolicyBinding []interface{} `json:"responderpolicylabel_responderpolicy_binding,omitempty"`
}

type Responderpolicy struct {
	Action             string   `json:"action,omitempty"`
	Appflowaction      string   `json:"appflowaction,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	Comment            string   `json:"comment,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	Logaction          string   `json:"logaction,omitempty"`
	Name               string   `json:"name,omitempty"`
	Newname            string   `json:"newname,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Rule               string   `json:"rule,omitempty"`
	Undefaction        string   `json:"undefaction,omitempty"`
	Undefhits          int      `json:"undefhits,omitempty"`
}

type ResponderpolicyVpnvserverBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Responderparam struct {
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Timeout            int    `json:"timeout,omitempty"`
	Undefaction        string `json:"undefaction,omitempty"`
}

type ResponderpolicyBinding struct {
	Name                                       string        `json:"name,omitempty"`
	ResponderpolicyCrvserverBinding            []interface{} `json:"responderpolicy_crvserver_binding,omitempty"`
	ResponderpolicyCsvserverBinding            []interface{} `json:"responderpolicy_csvserver_binding,omitempty"`
	ResponderpolicyLbvserverBinding            []interface{} `json:"responderpolicy_lbvserver_binding,omitempty"`
	ResponderpolicyResponderglobalBinding      []interface{} `json:"responderpolicy_responderglobal_binding,omitempty"`
	ResponderpolicyResponderpolicylabelBinding []interface{} `json:"responderpolicy_responderpolicylabel_binding,omitempty"`
	ResponderpolicyVpnvserverBinding           []interface{} `json:"responderpolicy_vpnvserver_binding,omitempty"`
}

type ResponderglobalBinding struct {
	ResponderglobalResponderpolicyBinding []interface{} `json:"responderglobal_responderpolicy_binding,omitempty"`
}
