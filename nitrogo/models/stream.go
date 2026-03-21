package models

// stream configuration structs
type StreamSession struct {
	Name string `json:"name,omitempty"`
}

type StreamIdentifier struct {
	AcceptanceThreshold     string   `json:"acceptancethreshold,omitempty"`
	AppFlowLog              string   `json:"appflowlog,omitempty"`
	BreachThreshold         int      `json:"breachthreshold,omitempty"`
	Count                   float64  `json:"__count,omitempty"`
	Interval                int      `json:"interval,omitempty"`
	Log                     string   `json:"log,omitempty"`
	LogInterval             int      `json:"loginterval,omitempty"`
	LogLimit                int      `json:"loglimit,omitempty"`
	MaxTransactionThreshold int      `json:"maxtransactionthreshold,omitempty"`
	MinTransactionThreshold int      `json:"mintransactionthreshold,omitempty"`
	Name                    string   `json:"name,omitempty"`
	NextGenAPIResource      string   `json:"_nextgenapiresource,omitempty"`
	Rule                    []string `json:"rule,omitempty"`
	SampleCount             int      `json:"samplecount,omitempty"`
	SelectorName            string   `json:"selectorname,omitempty"`
	SNMPTrap                string   `json:"snmptrap,omitempty"`
	Sort                    string   `json:"sort,omitempty"`
	TrackAckOnlyPackets     string   `json:"trackackonlypackets,omitempty"`
	TrackTransactions       string   `json:"tracktransactions,omitempty"`
}

type StreamIdentifierBinding struct {
	Name                                    string `json:"name,omitempty"`
	StreamIdentifierAnalyticsProfileBinding []any  `json:"streamidentifier_analyticsprofile_binding,omitempty"`
	StreamIdentifierStreamSessionBinding    []any  `json:"streamidentifier_streamsession_binding,omitempty"`
}

type StreamIdentifierAnalyticsProfileBinding struct {
	AnalyticsProfile string `json:"analyticsprofile,omitempty"`
	Name             string `json:"name,omitempty"`
}

type StreamSelector struct {
	Count              float64  `json:"__count,omitempty"`
	Name               string   `json:"name,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	Rule               []string `json:"rule,omitempty"`
}

type StreamIdentifierStreamSessionBinding struct {
	AnalyticsProfile string `json:"analyticsprofile,omitempty"`
	Name             string `json:"name,omitempty"`
}
