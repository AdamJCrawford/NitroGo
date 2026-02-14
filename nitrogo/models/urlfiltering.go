// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Urlfilteringparameter struct {
	Hoursbetweendbupdates     int    `json:"hoursbetweendbupdates,omitempty"`
	Timeofdaytoupdatedb       string `json:"timeofdaytoupdatedb,omitempty"`
	Localdatabasethreads      int    `json:"localdatabasethreads,omitempty"`
	Cloudhost                 string `json:"cloudhost,omitempty"`
	Seeddbpath                string `json:"seeddbpath,omitempty"`
	Maxnumberofcloudthreads   string `json:"maxnumberofcloudthreads,omitempty"`
	Cloudkeepalivetimeout     string `json:"cloudkeepalivetimeout,omitempty"`
	Cloudserverconnecttimeout string `json:"cloudserverconnecttimeout,omitempty"`
	Clouddblookuptimeout      string `json:"clouddblookuptimeout,omitempty"`
	Proxyhostip               string `json:"proxyhostip,omitempty"`
	Proxyport                 string `json:"proxyport,omitempty"`
	Proxyusername             string `json:"proxyusername,omitempty"`
	Proxypassword             string `json:"proxypassword,omitempty"`
	Seeddbsizelevel           string `json:"seeddbsizelevel,omitempty"`
}

type Urlfilteringcategories struct {
	Group      string `json:"group,omitempty"`
	Categories string `json:"categories,omitempty"`
}

type Urlfilteringcategorization struct {
	Url            string `json:"url,omitempty"`
	Categorization string `json:"categorization,omitempty"`
}

type Urlfilteringcategorygroups struct {
	Categorygroups string `json:"categorygroups,omitempty"`
}
