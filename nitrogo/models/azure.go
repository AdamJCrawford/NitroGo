package models

// azure configuration structs
type Azurekeyvault struct {
	Azureapplication   string  `json:"azureapplication,omitempty"`
	Azurevaultname     string  `json:"azurevaultname,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	State              string  `json:"state,omitempty"`
}

type Azureapplication struct {
	Clientid           string  `json:"clientid,omitempty"`
	Clientsecret       string  `json:"clientsecret,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Tenantid           string  `json:"tenantid,omitempty"`
	Tokenendpoint      string  `json:"tokenendpoint,omitempty"`
	Vaultresource      string  `json:"vaultresource,omitempty"`
}
