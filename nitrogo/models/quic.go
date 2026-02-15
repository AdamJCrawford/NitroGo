package models

// quic configuration structs
type Quicparam struct {
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Quicsecrettimeout  int    `json:"quicsecrettimeout,omitempty"`
}

type Quicprofile struct {
	Ackdelayexponent               int      `json:"ackdelayexponent,omitempty"`
	Activeconnectionidlimit        int      `json:"activeconnectionidlimit,omitempty"`
	Activeconnectionmigration      string   `json:"activeconnectionmigration,omitempty"`
	Builtin                        []string `json:"builtin,omitempty"`
	Congestionctrlalgorithm        string   `json:"congestionctrlalgorithm,omitempty"`
	Count                          float64  `json:"__count,omitempty"`
	Feature                        string   `json:"feature,omitempty"`
	Initialmaxdata                 int      `json:"initialmaxdata,omitempty"`
	Initialmaxstreamdatabidilocal  int      `json:"initialmaxstreamdatabidilocal,omitempty"`
	Initialmaxstreamdatabidiremote int      `json:"initialmaxstreamdatabidiremote,omitempty"`
	Initialmaxstreamdatauni        int      `json:"initialmaxstreamdatauni,omitempty"`
	Initialmaxstreamsbidi          int      `json:"initialmaxstreamsbidi,omitempty"`
	Initialmaxstreamsuni           int      `json:"initialmaxstreamsuni,omitempty"`
	Maxackdelay                    int      `json:"maxackdelay,omitempty"`
	Maxidletimeout                 int      `json:"maxidletimeout,omitempty"`
	Maxudpdatagramsperburst        int      `json:"maxudpdatagramsperburst,omitempty"`
	Maxudppayloadsize              int      `json:"maxudppayloadsize,omitempty"`
	Name                           string   `json:"name,omitempty"`
	Newtokenvalidityperiod         int      `json:"newtokenvalidityperiod,omitempty"`
	Nextgenapiresource             string   `json:"_nextgenapiresource,omitempty"`
	Refcnt                         int      `json:"refcnt,omitempty"`
	Retrytokenvalidityperiod       int      `json:"retrytokenvalidityperiod,omitempty"`
	Statelessaddressvalidation     string   `json:"statelessaddressvalidation,omitempty"`
}
