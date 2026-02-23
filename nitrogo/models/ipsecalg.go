package models

// ipsecalg configuration structs
type IPSECALGSession struct {
	Count              float64 `json:"__count,omitempty"`
	DestIP             string  `json:"destip,omitempty"`
	DestIPAlg          string  `json:"destip_alg,omitempty"`
	NATIP              string  `json:"natip,omitempty"`
	NATIPAlg           string  `json:"natip_alg,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	SourceIP           string  `json:"sourceip,omitempty"`
	SourceIPAlg        string  `json:"sourceip_alg,omitempty"`
	SPIIn              int     `json:"spiin,omitempty"`
	SPIOut             int     `json:"spiout,omitempty"`
}

type IPSECALGProfile struct {
	ConnFailover       string  `json:"connfailover,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	ESPGateTimeout     int     `json:"espgatetimeout,omitempty"`
	ESPSessionTimeout  int     `json:"espsessiontimeout,omitempty"`
	IKESessionTimeout  int     `json:"ikesessiontimeout,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
}
