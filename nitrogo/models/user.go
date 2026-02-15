package models

// user configuration structs
type Uservserver struct {
	Comment                   string  `json:"comment,omitempty"`
	Count                     float64 `json:"__count,omitempty"`
	Curstate                  string  `json:"curstate,omitempty"`
	Defaultlb                 string  `json:"defaultlb,omitempty"`
	Ipaddress                 string  `json:"ipaddress,omitempty"`
	Name                      string  `json:"name,omitempty"`
	Nextgenapiresource        string  `json:"_nextgenapiresource,omitempty"`
	Nodefaultbindings         string  `json:"nodefaultbindings,omitempty"`
	Params                    string  `json:"Params,omitempty"`
	Port                      int     `json:"port,omitempty"`
	State                     string  `json:"state,omitempty"`
	Statechangetimemsec       int     `json:"statechangetimemsec,omitempty"`
	Statechangetimesec        string  `json:"statechangetimesec,omitempty"`
	Tickssincelaststatechange int     `json:"tickssincelaststatechange,omitempty"`
	Userprotocol              string  `json:"userprotocol,omitempty"`
	Value                     string  `json:"value,omitempty"`
}

type Userprotocol struct {
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Extension          string  `json:"extension,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Transport          string  `json:"transport,omitempty"`
}
