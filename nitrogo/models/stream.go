package models

// stream configuration structs
type Streamsession struct {
	Name string `json:"name,omitempty"`
}

type Streamidentifier struct {
	Acceptancethreshold     string   `json:"acceptancethreshold,omitempty"`
	Appflowlog              string   `json:"appflowlog,omitempty"`
	Breachthreshold         int      `json:"breachthreshold,omitempty"`
	Count                   float64  `json:"__count,omitempty"`
	Interval                int      `json:"interval,omitempty"`
	Log                     string   `json:"log,omitempty"`
	Loginterval             int      `json:"loginterval,omitempty"`
	Loglimit                int      `json:"loglimit,omitempty"`
	Maxtransactionthreshold int      `json:"maxtransactionthreshold,omitempty"`
	Mintransactionthreshold int      `json:"mintransactionthreshold,omitempty"`
	Name                    string   `json:"name,omitempty"`
	Nextgenapiresource      string   `json:"_nextgenapiresource,omitempty"`
	Rule                    []string `json:"rule,omitempty"`
	Samplecount             int      `json:"samplecount,omitempty"`
	Selectorname            string   `json:"selectorname,omitempty"`
	Snmptrap                string   `json:"snmptrap,omitempty"`
	Sort                    string   `json:"sort,omitempty"`
	Trackackonlypackets     string   `json:"trackackonlypackets,omitempty"`
	Tracktransactions       string   `json:"tracktransactions,omitempty"`
}

type StreamidentifierBinding struct {
	Name                                    string        `json:"name,omitempty"`
	StreamidentifierAnalyticsprofileBinding []interface{} `json:"streamidentifier_analyticsprofile_binding,omitempty"`
	StreamidentifierStreamsessionBinding    []interface{} `json:"streamidentifier_streamsession_binding,omitempty"`
}

type StreamidentifierAnalyticsprofileBinding struct {
	Analyticsprofile string `json:"analyticsprofile,omitempty"`
	Name             string `json:"name,omitempty"`
}

type Streamselector struct {
	Count              float64  `json:"__count,omitempty"`
	Name               string   `json:"name,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Rule               []string `json:"rule,omitempty"`
}

type StreamidentifierStreamsessionBinding struct {
	Analyticsprofile string `json:"analyticsprofile,omitempty"`
	Name             string `json:"name,omitempty"`
}
