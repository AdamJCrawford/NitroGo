package models

// smpp configuration structs
type SMPPParam struct {
	AddrNPI            int    `json:"addrnpi,omitempty"`
	AddrRange          string `json:"addrrange,omitempty"`
	AddrTON            int    `json:"addrton,omitempty"`
	ClientMode         string `json:"clientmode,omitempty"`
	MsgQueue           string `json:"msgqueue,omitempty"`
	MsgQueueSize       int    `json:"msgqueuesize,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
}

type SMPPUser struct {
	Count              float64 `json:"__count,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Password           string  `json:"password,omitempty"`
	Username           string  `json:"username,omitempty"`
}
