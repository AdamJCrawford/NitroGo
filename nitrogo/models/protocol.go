package models

// protocol configuration structs
type Protocolhttpband struct {
	Accesscount        []interface{} `json:"accesscount,omitempty"`
	Accessratio        []interface{} `json:"accessratio,omitempty"`
	Accessrationew     []interface{} `json:"accessrationew,omitempty"`
	Avgbandsize        []interface{} `json:"avgbandsize,omitempty"`
	Avgbandsizenew     []interface{} `json:"avgbandsizenew,omitempty"`
	Banddata           []interface{} `json:"banddata,omitempty"`
	Banddatanew        []interface{} `json:"banddatanew,omitempty"`
	Bandrange          int           `json:"bandrange,omitempty"`
	Nextgenapiresource string        `json:"_nextgenapiresource,omitempty"`
	Nodeid             int           `json:"nodeid,omitempty"`
	Numberofbands      int           `json:"numberofbands,omitempty"`
	Reqbandsize        int           `json:"reqbandsize,omitempty"`
	Respbandsize       int           `json:"respbandsize,omitempty"`
	Totalbandsize      []interface{} `json:"totalbandsize,omitempty"`
	Totals             []interface{} `json:"totals,omitempty"`
	TypeField          string        `json:"type,omitempty"`
}
