package models

// cloudtunnel configuration structs
type CloudTunnelParameter struct {
	ControllerFQDN                 string `json:"controllerfqdn,omitempty"`
	FQDN                           string `json:"fqdn,omitempty"`
	NextGenAPIResource             string `json:"_nextgenapiresource,omitempty"`
	ResourceLocation               string `json:"resourcelocation,omitempty"`
	SubnetResourceLocationMappings string `json:"subnetresourcelocationmappings,omitempty"`
}

type CloudTunnelVServer struct {
	CacheType          string  `json:"cachetype,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	EffectiveState     string  `json:"effectivestate,omitempty"`
	IP                 string  `json:"ip,omitempty"`
	IPPattern          string  `json:"ippattern,omitempty"`
	IPv46              string  `json:"ipv46,omitempty"`
	ListenPolicy       string  `json:"listenpolicy,omitempty"`
	ListenPriority     int     `json:"listenpriority,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Port               int     `json:"port,omitempty"`
	Range              int     `json:"range,omitempty"`
	ServiceType        string  `json:"servicetype,omitempty"`
	State              string  `json:"state,omitempty"`
	TypeField          string  `json:"type,omitempty"`
}
