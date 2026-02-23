package models

// azure configuration structs
type AzureKeyVault struct {
	AzureApplication   string  `json:"azureapplication,omitempty"`
	AzureVaultName     string  `json:"azurevaultname,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	State              string  `json:"state,omitempty"`
}

type AzureApplication struct {
	ClientID           string  `json:"clientid,omitempty"`
	ClientSecret       string  `json:"clientsecret,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	TenantID           string  `json:"tenantid,omitempty"`
	TokenEndpoint      string  `json:"tokenendpoint,omitempty"`
	VaultResource      string  `json:"vaultresource,omitempty"`
}
