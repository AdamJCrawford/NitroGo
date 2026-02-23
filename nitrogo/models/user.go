package models

// user configuration structs
type UserVServer struct {
	Comment                   string  `json:"comment,omitempty"`
	Count                     float64 `json:"__count,omitempty"`
	CurState                  string  `json:"curstate,omitempty"`
	DefaultLB                 string  `json:"defaultlb,omitempty"`
	IPAddress                 string  `json:"ipaddress,omitempty"`
	Name                      string  `json:"name,omitempty"`
	NextGenAPIResource        string  `json:"_nextgenapiresource,omitempty"`
	NoDefaultBindings         string  `json:"nodefaultbindings,omitempty"`
	Params                    string  `json:"Params,omitempty"`
	Port                      int     `json:"port,omitempty"`
	State                     string  `json:"state,omitempty"`
	StateChangeTimeMsec       int     `json:"statechangetimemsec,omitempty"`
	StateChangeTimeSec        string  `json:"statechangetimesec,omitempty"`
	TicksSinceLastStateChange int     `json:"tickssincelaststatechange,omitempty"`
	UserProtocol              string  `json:"userprotocol,omitempty"`
	Value                     string  `json:"value,omitempty"`
}

type UserProtocol struct {
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Extension          string  `json:"extension,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Transport          string  `json:"transport,omitempty"`
}
