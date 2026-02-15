package models

// reputation configuration structs
type Reputationsettings struct {
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Proxypassword      string `json:"proxypassword,omitempty"`
	Proxyport          int    `json:"proxyport,omitempty"`
	Proxyserver        string `json:"proxyserver,omitempty"`
	Proxyusername      string `json:"proxyusername,omitempty"`
}
