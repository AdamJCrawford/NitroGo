package models

// smpp configuration structs
type Smppparam struct {
	Addrnpi            int    `json:"addrnpi,omitempty"`
	Addrrange          string `json:"addrrange,omitempty"`
	Addrton            int    `json:"addrton,omitempty"`
	Clientmode         string `json:"clientmode,omitempty"`
	Msgqueue           string `json:"msgqueue,omitempty"`
	Msgqueuesize       int    `json:"msgqueuesize,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Smppuser struct {
	Count              float64 `json:"__count,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Password           string  `json:"password,omitempty"`
	Username           string  `json:"username,omitempty"`
}
