package models

// rewrite configuration structs
type RewriteglobalBinding struct {
	RewriteglobalRewritepolicyBinding []interface{} `json:"rewriteglobal_rewritepolicy_binding,omitempty"`
}

type Rewritepolicy struct {
	Action             string   `json:"action,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	Comment            string   `json:"comment,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Description        string   `json:"description,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	Isdefault          bool     `json:"isdefault,omitempty"`
	Logaction          string   `json:"logaction,omitempty"`
	Name               string   `json:"name,omitempty"`
	Newname            string   `json:"newname,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Rule               string   `json:"rule,omitempty"`
	Undefaction        string   `json:"undefaction,omitempty"`
	Undefhits          int      `json:"undefhits,omitempty"`
}

type Rewriteaction struct {
	Builtin            []string `json:"builtin,omitempty"`
	Comment            string   `json:"comment,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Description        string   `json:"description,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	Isdefault          bool     `json:"isdefault,omitempty"`
	Name               string   `json:"name,omitempty"`
	Newname            string   `json:"newname,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Referencecount     int      `json:"referencecount,omitempty"`
	Refinesearch       string   `json:"refinesearch,omitempty"`
	Search             string   `json:"search,omitempty"`
	Stringbuilderexpr  string   `json:"stringbuilderexpr,omitempty"`
	Target             string   `json:"target,omitempty"`
	TypeField          string   `json:"type,omitempty"`
	Undefhits          int      `json:"undefhits,omitempty"`
}

type RewritepolicyVpnvserverBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type RewritepolicyLbvserverBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type RewritepolicylabelRewritepolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type RewritepolicylabelBinding struct {
	Labelname                              string        `json:"labelname,omitempty"`
	RewritepolicylabelPolicybindingBinding []interface{} `json:"rewritepolicylabel_policybinding_binding,omitempty"`
	RewritepolicylabelRewritepolicyBinding []interface{} `json:"rewritepolicylabel_rewritepolicy_binding,omitempty"`
}

type RewritepolicyCsvserverBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type RewritepolicyRewriteglobalBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type RewritepolicyBinding struct {
	Name                                   string        `json:"name,omitempty"`
	RewritepolicyCsvserverBinding          []interface{} `json:"rewritepolicy_csvserver_binding,omitempty"`
	RewritepolicyLbvserverBinding          []interface{} `json:"rewritepolicy_lbvserver_binding,omitempty"`
	RewritepolicyRewriteglobalBinding      []interface{} `json:"rewritepolicy_rewriteglobal_binding,omitempty"`
	RewritepolicyRewritepolicylabelBinding []interface{} `json:"rewritepolicy_rewritepolicylabel_binding,omitempty"`
	RewritepolicyVpnvserverBinding         []interface{} `json:"rewritepolicy_vpnvserver_binding,omitempty"`
}

type Rewriteparam struct {
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Timeout            int    `json:"timeout,omitempty"`
	Undefaction        string `json:"undefaction,omitempty"`
}

type Rewritepolicylabel struct {
	Builtin                []string `json:"builtin,omitempty"`
	Comment                string   `json:"comment,omitempty"`
	Count                  float64  `json:"__count,omitempty"`
	Description            string   `json:"description,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	Flowtype               int      `json:"flowtype,omitempty"`
	Gotopriorityexpression string   `json:"gotopriorityexpression,omitempty"`
	Hits                   int      `json:"hits,omitempty"`
	InvokeLabelname        string   `json:"invoke_labelname,omitempty"`
	Isdefault              bool     `json:"isdefault,omitempty"`
	Labelname              string   `json:"labelname,omitempty"`
	Labeltype              string   `json:"labeltype,omitempty"`
	Newname                string   `json:"newname,omitempty"`
	Nextgenapiresource     string   `json:"_nextgenapiresource,omitempty"`
	Numpol                 int      `json:"numpol,omitempty"`
	Priority               int      `json:"priority,omitempty"`
	Transform              string   `json:"transform,omitempty"`
}

type RewritepolicyRewritepolicylabelBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type RewriteglobalRewritepolicyBinding struct {
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

type RewritepolicylabelPolicybindingBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}
