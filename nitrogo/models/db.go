package models

// db configuration structs
type Dbuser struct {
	Count              float64 `json:"__count,omitempty"`
	Loggedin           bool    `json:"loggedin,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Password           string  `json:"password,omitempty"`
	Username           string  `json:"username,omitempty"`
}

type Dbdbprofile struct {
	Conmultiplex           string  `json:"conmultiplex,omitempty"`
	Count                  float64 `json:"__count,omitempty"`
	Enablecachingconmuxoff string  `json:"enablecachingconmuxoff,omitempty"`
	Interpretquery         string  `json:"interpretquery,omitempty"`
	Kcdaccount             string  `json:"kcdaccount,omitempty"`
	Name                   string  `json:"name,omitempty"`
	Nextgenapiresource     string  `json:"_nextgenapiresource,omitempty"`
	Refcnt                 int     `json:"refcnt,omitempty"`
	Stickiness             string  `json:"stickiness,omitempty"`
}
