package models

// reputation configuration structs
type ReputationSettings struct {
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	ProxyPassword      string `json:"proxypassword,omitempty"`
	ProxyPort          int    `json:"proxyport,omitempty"`
	ProxyServer        string `json:"proxyserver,omitempty"`
	ProxyUsername      string `json:"proxyusername,omitempty"`
}
