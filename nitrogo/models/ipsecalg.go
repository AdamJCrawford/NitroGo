package models

// ipsecalg configuration structs
type Ipsecalgsession struct {
	Count              float64 `json:"__count,omitempty"`
	Destip             string  `json:"destip,omitempty"`
	DestipAlg          string  `json:"destip_alg,omitempty"`
	Natip              string  `json:"natip,omitempty"`
	NatipAlg           string  `json:"natip_alg,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Sourceip           string  `json:"sourceip,omitempty"`
	SourceipAlg        string  `json:"sourceip_alg,omitempty"`
	Spiin              int     `json:"spiin,omitempty"`
	Spiout             int     `json:"spiout,omitempty"`
}

type Ipsecalgprofile struct {
	Connfailover       string  `json:"connfailover,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Espgatetimeout     int     `json:"espgatetimeout,omitempty"`
	Espsessiontimeout  int     `json:"espsessiontimeout,omitempty"`
	Ikesessiontimeout  int     `json:"ikesessiontimeout,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
}
