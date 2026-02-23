package models

// appqoe configuration structs
type AppQOECustomResp struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Src                string  `json:"src,omitempty"`
}

type AppQOEPolicyLBVServerBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BindPriority           int    `json:"bindpriority,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type AppQOEParameter struct {
	AvgWaitingClient    int    `json:"avgwaitingclient,omitempty"`
	DosAttackThresh     int    `json:"dosattackthresh,omitempty"`
	MaxAltRespBandwidth int    `json:"maxaltrespbandwidth,omitempty"`
	NextGenAPIResource  string `json:"_nextgenapiresource,omitempty"`
	SessionLife         int    `json:"sessionlife,omitempty"`
}

type AppQOEAction struct {
	AltContentPath     string  `json:"altcontentpath,omitempty"`
	AltContentSvcName  string  `json:"altcontentsvcname,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	CustomFile         string  `json:"customfile,omitempty"`
	Delay              int     `json:"delay,omitempty"`
	DosAction          string  `json:"dosaction,omitempty"`
	DosTrigExpression  string  `json:"dostrigexpression,omitempty"`
	Hits               int     `json:"hits,omitempty"`
	MaxConn            int     `json:"maxconn,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NumRetries         int     `json:"numretries,omitempty"`
	PolQDepth          int     `json:"polqdepth,omitempty"`
	Priority           string  `json:"priority,omitempty"`
	PriQDepth          int     `json:"priqdepth,omitempty"`
	RespondWith        string  `json:"respondwith,omitempty"`
	RetryOnReset       string  `json:"retryonreset,omitempty"`
	RetryOnTimeout     int     `json:"retryontimeout,omitempty"`
	TCPProfile         string  `json:"tcpprofile,omitempty"`
}

type AppQOEPolicyBinding struct {
	AppQOEPolicyLBVServerBinding []interface{} `json:"appqoepolicy_lbvserver_binding,omitempty"`
	Name                         string        `json:"name,omitempty"`
}

type AppQOEPolicy struct {
	Action             string  `json:"action,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Hits               int     `json:"hits,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Rule               string  `json:"rule,omitempty"`
}
