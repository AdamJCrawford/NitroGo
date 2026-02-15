package models

// appqoe configuration structs
type Appqoecustomresp struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Src                string  `json:"src,omitempty"`
}

type AppqoepolicyLbvserverBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Bindpriority           int    `json:"bindpriority,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Appqoeparameter struct {
	Avgwaitingclient    int    `json:"avgwaitingclient,omitempty"`
	Dosattackthresh     int    `json:"dosattackthresh,omitempty"`
	Maxaltrespbandwidth int    `json:"maxaltrespbandwidth,omitempty"`
	Nextgenapiresource  string `json:"_nextgenapiresource,omitempty"`
	Sessionlife         int    `json:"sessionlife,omitempty"`
}

type Appqoeaction struct {
	Altcontentpath     string  `json:"altcontentpath,omitempty"`
	Altcontentsvcname  string  `json:"altcontentsvcname,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Customfile         string  `json:"customfile,omitempty"`
	Delay              int     `json:"delay,omitempty"`
	Dosaction          string  `json:"dosaction,omitempty"`
	Dostrigexpression  string  `json:"dostrigexpression,omitempty"`
	Hits               int     `json:"hits,omitempty"`
	Maxconn            int     `json:"maxconn,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Numretries         int     `json:"numretries,omitempty"`
	Polqdepth          int     `json:"polqdepth,omitempty"`
	Priority           string  `json:"priority,omitempty"`
	Priqdepth          int     `json:"priqdepth,omitempty"`
	Respondwith        string  `json:"respondwith,omitempty"`
	Retryonreset       string  `json:"retryonreset,omitempty"`
	Retryontimeout     int     `json:"retryontimeout,omitempty"`
	Tcpprofile         string  `json:"tcpprofile,omitempty"`
}

type AppqoepolicyBinding struct {
	AppqoepolicyLbvserverBinding []interface{} `json:"appqoepolicy_lbvserver_binding,omitempty"`
	Name                         string        `json:"name,omitempty"`
}

type Appqoepolicy struct {
	Action             string  `json:"action,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Hits               int     `json:"hits,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Rule               string  `json:"rule,omitempty"`
}
