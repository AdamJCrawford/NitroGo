package models

// videooptimization configuration structs
type VideooptimizationglobalpacingBinding struct {
	VideooptimizationglobalpacingVideooptimizationpacingpolicyBinding []interface{} `json:"videooptimizationglobalpacing_videooptimizationpacingpolicy_binding,omitempty"`
}

type Videooptimizationparameter struct {
	Nextgenapiresource       string  `json:"_nextgenapiresource,omitempty"`
	Quicpacingrate           int     `json:"quicpacingrate,omitempty"`
	Randomsamplingpercentage float64 `json:"randomsamplingpercentage,omitempty"`
}

type VideooptimizationpacingpolicylabelBinding struct {
	Labelname                                                              string        `json:"labelname,omitempty"`
	VideooptimizationpacingpolicylabelPolicybindingBinding                 []interface{} `json:"videooptimizationpacingpolicylabel_policybinding_binding,omitempty"`
	VideooptimizationpacingpolicylabelVideooptimizationpacingpolicyBinding []interface{} `json:"videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding,omitempty"`
}

type Videooptimizationpacingaction struct {
	Builtin            []string `json:"builtin,omitempty"`
	Comment            string   `json:"comment,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	Name               string   `json:"name,omitempty"`
	Newname            string   `json:"newname,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Rate               int      `json:"rate,omitempty"`
	Referencecount     int      `json:"referencecount,omitempty"`
	Undefhits          int      `json:"undefhits,omitempty"`
}

type Videooptimizationdetectionpolicylabel struct {
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

type VideooptimizationglobaldetectionVideooptimizationdetectionpolicyBinding struct {
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

type VideooptimizationdetectionpolicyLbvserverBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type VideooptimizationpacingpolicyLbvserverBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Videooptimizationpacingpolicy struct {
	Action             string   `json:"action,omitempty"`
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

type VideooptimizationdetectionpolicylabelPolicybindingBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type VideooptimizationdetectionpolicyBinding struct {
	Name                                                                    string        `json:"name,omitempty"`
	VideooptimizationdetectionpolicyLbvserverBinding                        []interface{} `json:"videooptimizationdetectionpolicy_lbvserver_binding,omitempty"`
	VideooptimizationdetectionpolicyVideooptimizationglobaldetectionBinding []interface{} `json:"videooptimizationdetectionpolicy_videooptimizationglobaldetection_binding,omitempty"`
}

type VideooptimizationglobalpacingVideooptimizationpacingpolicyBinding struct {
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

type VideooptimizationpacingpolicyBinding struct {
	Name                                                              string        `json:"name,omitempty"`
	VideooptimizationpacingpolicyLbvserverBinding                     []interface{} `json:"videooptimizationpacingpolicy_lbvserver_binding,omitempty"`
	VideooptimizationpacingpolicyVideooptimizationglobalpacingBinding []interface{} `json:"videooptimizationpacingpolicy_videooptimizationglobalpacing_binding,omitempty"`
}

type VideooptimizationdetectionpolicylabelBinding struct {
	Labelname                                                                    string        `json:"labelname,omitempty"`
	VideooptimizationdetectionpolicylabelPolicybindingBinding                    []interface{} `json:"videooptimizationdetectionpolicylabel_policybinding_binding,omitempty"`
	VideooptimizationdetectionpolicylabelVideooptimizationdetectionpolicyBinding []interface{} `json:"videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding,omitempty"`
}

type VideooptimizationglobaldetectionBinding struct {
	VideooptimizationglobaldetectionVideooptimizationdetectionpolicyBinding []interface{} `json:"videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding,omitempty"`
}

type VideooptimizationpacingpolicylabelPolicybindingBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type VideooptimizationdetectionpolicylabelVideooptimizationdetectionpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Videooptimizationdetectionaction struct {
	Builtin            []string `json:"builtin,omitempty"`
	Comment            string   `json:"comment,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	Name               string   `json:"name,omitempty"`
	Newname            string   `json:"newname,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Referencecount     int      `json:"referencecount,omitempty"`
	TypeField          string   `json:"type,omitempty"`
	Undefhits          int      `json:"undefhits,omitempty"`
}

type VideooptimizationpacingpolicyVideooptimizationglobalpacingBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type VideooptimizationpacingpolicylabelVideooptimizationpacingpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Videooptimizationdetectionpolicy struct {
	Action             string   `json:"action,omitempty"`
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

type Videooptimizationpacingpolicylabel struct {
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

type VideooptimizationdetectionpolicyVideooptimizationglobaldetectionBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}
