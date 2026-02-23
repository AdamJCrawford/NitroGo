package models

// autoscale configuration structs
type AutoscaleProfile struct {
	APIKey             string  `json:"apikey,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	SharedSecret       string  `json:"sharedsecret,omitempty"`
	TypeField          string  `json:"type,omitempty"`
	URL                string  `json:"url,omitempty"`
}

type AutoscaleAction struct {
	Count                float64 `json:"__count,omitempty"`
	Name                 string  `json:"name,omitempty"`
	NextGenAPIResource   string  `json:"_nextgenapiresource,omitempty"`
	Parameters           string  `json:"parameters,omitempty"`
	ProfileName          string  `json:"profilename,omitempty"`
	QuietTime            int     `json:"quiettime,omitempty"`
	TypeField            string  `json:"type,omitempty"`
	VMDestroyGracePeriod int     `json:"vmdestroygraceperiod,omitempty"`
	VServer              string  `json:"vserver,omitempty"`
}

type AutoscalePolicyNSTimerBinding struct {
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type AutoscalePolicyBinding struct {
	AutoscalePolicyNSTimerBinding []interface{} `json:"autoscalepolicy_nstimer_binding,omitempty"`
	Name                          string        `json:"name,omitempty"`
}

type AutoscalePolicy struct {
	Action             string  `json:"action,omitempty"`
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Hits               int     `json:"hits,omitempty"`
	LogAction          string  `json:"logaction,omitempty"`
	Name               string  `json:"name,omitempty"`
	NewName            string  `json:"newname,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Priority           int     `json:"priority,omitempty"`
	Rule               string  `json:"rule,omitempty"`
	UndefHits          int     `json:"undefhits,omitempty"`
}
