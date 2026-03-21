package models

// protocol configuration structs
type ProtocolHTTPBand struct {
	AccessCount        []any  `json:"accesscount,omitempty"`
	AccessRatio        []any  `json:"accessratio,omitempty"`
	AccessRatioNew     []any  `json:"accessrationew,omitempty"`
	AvgBandSize        []any  `json:"avgbandsize,omitempty"`
	AvgBandSizeNew     []any  `json:"avgbandsizenew,omitempty"`
	BandData           []any  `json:"banddata,omitempty"`
	BandDataNew        []any  `json:"banddatanew,omitempty"`
	BandRange          int    `json:"bandrange,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	NodeID             int    `json:"nodeid,omitempty"`
	NumberOfBands      int    `json:"numberofbands,omitempty"`
	ReqBandSize        int    `json:"reqbandsize,omitempty"`
	RespBandSize       int    `json:"respbandsize,omitempty"`
	TotalBandSize      []any  `json:"totalbandsize,omitempty"`
	Totals             []any  `json:"totals,omitempty"`
	TypeField          string `json:"type,omitempty"`
}
