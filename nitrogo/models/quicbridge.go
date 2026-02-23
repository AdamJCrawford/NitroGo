package models

// quicbridge configuration structs
type QUICBridgeProfile struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	RefCnt             int     `json:"refcnt,omitempty"`
	RoutingAlgorithm   string  `json:"routingalgorithm,omitempty"`
	ServerIDLength     int     `json:"serveridlength,omitempty"`
}
