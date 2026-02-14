// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Azureapplication struct {
	Name               string `json:"name,omitempty"`
	Clientid           string `json:"clientid,omitempty"`
	Clientsecret       string `json:"clientsecret,omitempty"`
	Tenantid           string `json:"tenantid,omitempty"`
	Vaultresource      string `json:"vaultresource,omitempty"`
	Tokenendpoint      string `json:"tokenendpoint,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Azurekeyvault struct {
	Name               string `json:"name,omitempty"`
	Azurevaultname     string `json:"azurevaultname,omitempty"`
	Azureapplication   string `json:"azureapplication,omitempty"`
	State              string `json:"state,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}
