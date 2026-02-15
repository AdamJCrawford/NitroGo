package models

// cloudtunnel configuration structs
type Cloudtunnelparameter struct {
	Controllerfqdn                 string `json:"controllerfqdn,omitempty"`
	Fqdn                           string `json:"fqdn,omitempty"`
	Nextgenapiresource             string `json:"_nextgenapiresource,omitempty"`
	Resourcelocation               string `json:"resourcelocation,omitempty"`
	Subnetresourcelocationmappings string `json:"subnetresourcelocationmappings,omitempty"`
}

type Cloudtunnelvserver struct {
	Cachetype          string  `json:"cachetype,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Effectivestate     string  `json:"effectivestate,omitempty"`
	Ip                 string  `json:"ip,omitempty"`
	Ippattern          string  `json:"ippattern,omitempty"`
	Ipv46              string  `json:"ipv46,omitempty"`
	Listenpolicy       string  `json:"listenpolicy,omitempty"`
	Listenpriority     int     `json:"listenpriority,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Port               int     `json:"port,omitempty"`
	Range              int     `json:"range,omitempty"`
	Servicetype        string  `json:"servicetype,omitempty"`
	State              string  `json:"state,omitempty"`
	TypeField          string  `json:"type,omitempty"`
}
