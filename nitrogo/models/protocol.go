package models

// protocol configuration structs
type ProtocolHTTPBand struct {
	AccessCount        []interface{} `json:"accesscount,omitempty"`
	AccessRatio        []interface{} `json:"accessratio,omitempty"`
	AccessRatioNew     []interface{} `json:"accessrationew,omitempty"`
	AvgBandSize        []interface{} `json:"avgbandsize,omitempty"`
	AvgBandSizeNew     []interface{} `json:"avgbandsizenew,omitempty"`
	BandData           []interface{} `json:"banddata,omitempty"`
	BandDataNew        []interface{} `json:"banddatanew,omitempty"`
	BandRange          int           `json:"bandrange,omitempty"`
	NextGenAPIResource string        `json:"_nextgenapiresource,omitempty"`
	NodeID             int           `json:"nodeid,omitempty"`
	NumberOfBands      int           `json:"numberofbands,omitempty"`
	ReqBandSize        int           `json:"reqbandsize,omitempty"`
	RespBandSize       int           `json:"respbandsize,omitempty"`
	TotalBandSize      []interface{} `json:"totalbandsize,omitempty"`
	Totals             []interface{} `json:"totals,omitempty"`
	TypeField          string        `json:"type,omitempty"`
}
