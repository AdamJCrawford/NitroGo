package models

// autoscale configuration structs
type Autoscaleprofile struct {
	Apikey             string  `json:"apikey,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Sharedsecret       string  `json:"sharedsecret,omitempty"`
	TypeField          string  `json:"type,omitempty"`
	Url                string  `json:"url,omitempty"`
}

type Autoscaleaction struct {
	Count                float64 `json:"__count,omitempty"`
	Name                 string  `json:"name,omitempty"`
	Nextgenapiresource   string  `json:"_nextgenapiresource,omitempty"`
	Parameters           string  `json:"parameters,omitempty"`
	Profilename          string  `json:"profilename,omitempty"`
	Quiettime            int     `json:"quiettime,omitempty"`
	TypeField            string  `json:"type,omitempty"`
	Vmdestroygraceperiod int     `json:"vmdestroygraceperiod,omitempty"`
	Vserver              string  `json:"vserver,omitempty"`
}

type AutoscalepolicyNstimerBinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type AutoscalepolicyBinding struct {
	AutoscalepolicyNstimerBinding []interface{} `json:"autoscalepolicy_nstimer_binding,omitempty"`
	Name                          string        `json:"name,omitempty"`
}

type Autoscalepolicy struct {
	Action             string  `json:"action,omitempty"`
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Hits               int     `json:"hits,omitempty"`
	Logaction          string  `json:"logaction,omitempty"`
	Name               string  `json:"name,omitempty"`
	Newname            string  `json:"newname,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Priority           int     `json:"priority,omitempty"`
	Rule               string  `json:"rule,omitempty"`
	Undefhits          int     `json:"undefhits,omitempty"`
}
