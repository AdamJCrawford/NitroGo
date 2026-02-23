package models

// quic configuration structs
type QUICParam struct {
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	QUICSecretTimeout  int    `json:"quicsecrettimeout,omitempty"`
}

type QUICProfile struct {
	AckDelayExponent               int      `json:"ackdelayexponent,omitempty"`
	ActiveConnectionIDLimit        int      `json:"activeconnectionidlimit,omitempty"`
	ActiveConnectionMigration      string   `json:"activeconnectionmigration,omitempty"`
	Builtin                        []string `json:"builtin,omitempty"`
	CongestionCtrlAlgorithm        string   `json:"congestionctrlalgorithm,omitempty"`
	Count                          float64  `json:"__count,omitempty"`
	Feature                        string   `json:"feature,omitempty"`
	InitialMaxData                 int      `json:"initialmaxdata,omitempty"`
	InitialMaxStreamDataBidiLocal  int      `json:"initialmaxstreamdatabidilocal,omitempty"`
	InitialMaxStreamDataBidiRemote int      `json:"initialmaxstreamdatabidiremote,omitempty"`
	InitialMaxStreamDataUni        int      `json:"initialmaxstreamdatauni,omitempty"`
	InitialMaxStreamsBidi          int      `json:"initialmaxstreamsbidi,omitempty"`
	InitialMaxStreamsUni           int      `json:"initialmaxstreamsuni,omitempty"`
	MaxAckDelay                    int      `json:"maxackdelay,omitempty"`
	MaxIdleTimeout                 int      `json:"maxidletimeout,omitempty"`
	MaxUDPDatagramsPerBurst        int      `json:"maxudpdatagramsperburst,omitempty"`
	MaxUDPPayloadSize              int      `json:"maxudppayloadsize,omitempty"`
	Name                           string   `json:"name,omitempty"`
	NewTokenValidityPeriod         int      `json:"newtokenvalidityperiod,omitempty"`
	NextGenAPIResource             string   `json:"_nextgenapiresource,omitempty"`
	RefCnt                         int      `json:"refcnt,omitempty"`
	RetryTokenValidityPeriod       int      `json:"retrytokenvalidityperiod,omitempty"`
	StatelessAddressValidation     string   `json:"statelessaddressvalidation,omitempty"`
}
