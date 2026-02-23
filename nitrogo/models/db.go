package models

// db configuration structs
type DBUser struct {
	Count              float64 `json:"__count,omitempty"`
	LoggedIn           bool    `json:"loggedin,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Password           string  `json:"password,omitempty"`
	Username           string  `json:"username,omitempty"`
}

type DBDBProfile struct {
	ConMultiplex           string  `json:"conmultiplex,omitempty"`
	Count                  float64 `json:"__count,omitempty"`
	EnableCachingConMuxOff string  `json:"enablecachingconmuxoff,omitempty"`
	InterpretQuery         string  `json:"interpretquery,omitempty"`
	KCDAccount             string  `json:"kcdaccount,omitempty"`
	Name                   string  `json:"name,omitempty"`
	NextGenAPIResource     string  `json:"_nextgenapiresource,omitempty"`
	RefCnt                 int     `json:"refcnt,omitempty"`
	Stickiness             string  `json:"stickiness,omitempty"`
}
