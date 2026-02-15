package models

// ntp configuration structs
type Ntpserver struct {
	Autokey            bool    `json:"autokey,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Key                int     `json:"key,omitempty"`
	Maxpoll            int     `json:"maxpoll,omitempty"`
	Minpoll            int     `json:"minpoll,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Preferredntpserver string  `json:"preferredntpserver,omitempty"`
	Serverip           string  `json:"serverip,omitempty"`
	Servername         string  `json:"servername,omitempty"`
}

type Ntpparam struct {
	Authentication     string        `json:"authentication,omitempty"`
	Autokeylogsec      int           `json:"autokeylogsec,omitempty"`
	Nextgenapiresource string        `json:"_nextgenapiresource,omitempty"`
	Revokelogsec       int           `json:"revokelogsec,omitempty"`
	Trustedkey         []interface{} `json:"trustedkey,omitempty"`
}

type Ntpsync struct {
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	State              string `json:"state,omitempty"`
}

type Ntpstatus struct {
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Response           string `json:"response,omitempty"`
}
