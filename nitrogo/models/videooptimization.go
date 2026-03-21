package models

// videooptimization configuration structs
type VideoOptimizationGlobalPacingBinding struct {
	VideoOptimizationGlobalPacingVideoOptimizationPacingPolicyBinding []any `json:"videooptimizationglobalpacing_videooptimizationpacingpolicy_binding,omitempty"`
}

type VideoOptimizationParameter struct {
	NextGenAPIResource       string  `json:"_nextgenapiresource,omitempty"`
	QUICPacingRate           int     `json:"quicpacingrate,omitempty"`
	RandomSamplingPercentage float64 `json:"randomsamplingpercentage,omitempty"`
}

type VideoOptimizationPacingPolicyLabelBinding struct {
	LabelName                                                              string `json:"labelname,omitempty"`
	VideoOptimizationPacingPolicyLabelPolicyBindingBinding                 []any  `json:"videooptimizationpacingpolicylabel_policybinding_binding,omitempty"`
	VideoOptimizationPacingPolicyLabelVideoOptimizationPacingPolicyBinding []any  `json:"videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding,omitempty"`
}

type VideoOptimizationPacingAction struct {
	Builtin            []string `json:"builtin,omitempty"`
	Comment            string   `json:"comment,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	Name               string   `json:"name,omitempty"`
	NewName            string   `json:"newname,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	Rate               int      `json:"rate,omitempty"`
	ReferenceCount     int      `json:"referencecount,omitempty"`
	UndefHits          int      `json:"undefhits,omitempty"`
}

type VideoOptimizationDetectionPolicyLabel struct {
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

type VideoOptimizationGlobalDetectionVideoOptimizationDetectionPolicyBinding struct {
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

type VideoOptimizationDetectionPolicyLBVServerBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type VideoOptimizationPacingPolicyLBVServerBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type VideoOptimizationPacingPolicy struct {
	Action             string   `json:"action,omitempty"`
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

type VideoOptimizationDetectionPolicyLabelPolicyBindingBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelName        string `json:"invoke_labelname,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type VideoOptimizationDetectionPolicyBinding struct {
	Name                                                                    string `json:"name,omitempty"`
	VideoOptimizationDetectionPolicyLBVServerBinding                        []any  `json:"videooptimizationdetectionpolicy_lbvserver_binding,omitempty"`
	VideoOptimizationDetectionPolicyVideoOptimizationGlobalDetectionBinding []any  `json:"videooptimizationdetectionpolicy_videooptimizationglobaldetection_binding,omitempty"`
}

type VideoOptimizationGlobalPacingVideoOptimizationPacingPolicyBinding struct {
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

type VideoOptimizationPacingPolicyBinding struct {
	Name                                                              string `json:"name,omitempty"`
	VideoOptimizationPacingPolicyLBVServerBinding                     []any  `json:"videooptimizationpacingpolicy_lbvserver_binding,omitempty"`
	VideoOptimizationPacingPolicyVideoOptimizationGlobalPacingBinding []any  `json:"videooptimizationpacingpolicy_videooptimizationglobalpacing_binding,omitempty"`
}

type VideoOptimizationDetectionPolicyLabelBinding struct {
	LabelName                                                                    string `json:"labelname,omitempty"`
	VideoOptimizationDetectionPolicyLabelPolicyBindingBinding                    []any  `json:"videooptimizationdetectionpolicylabel_policybinding_binding,omitempty"`
	VideoOptimizationDetectionPolicyLabelVideoOptimizationDetectionPolicyBinding []any  `json:"videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding,omitempty"`
}

type VideoOptimizationGlobalDetectionBinding struct {
	VideoOptimizationGlobalDetectionVideoOptimizationDetectionPolicyBinding []any `json:"videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding,omitempty"`
}

type VideoOptimizationPacingPolicyLabelPolicyBindingBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelName        string `json:"invoke_labelname,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type VideoOptimizationDetectionPolicyLabelVideoOptimizationDetectionPolicyBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelName        string `json:"invoke_labelname,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type VideoOptimizationDetectionAction struct {
	Builtin            []string `json:"builtin,omitempty"`
	Comment            string   `json:"comment,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	Name               string   `json:"name,omitempty"`
	NewName            string   `json:"newname,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	ReferenceCount     int      `json:"referencecount,omitempty"`
	TypeField          string   `json:"type,omitempty"`
	UndefHits          int      `json:"undefhits,omitempty"`
}

type VideoOptimizationPacingPolicyVideoOptimizationGlobalPacingBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type VideoOptimizationPacingPolicyLabelVideoOptimizationPacingPolicyBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelName        string `json:"invoke_labelname,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type VideoOptimizationDetectionPolicy struct {
	Action             string   `json:"action,omitempty"`
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

type VideoOptimizationPacingPolicyLabel struct {
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

type VideoOptimizationDetectionPolicyVideoOptimizationGlobalDetectionBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}
