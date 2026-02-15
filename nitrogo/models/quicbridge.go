package models

// quicbridge configuration structs
type Quicbridgeprofile struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Refcnt             int     `json:"refcnt,omitempty"`
	Routingalgorithm   string  `json:"routingalgorithm,omitempty"`
	Serveridlength     int     `json:"serveridlength,omitempty"`
}
