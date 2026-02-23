// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type URLFilteringParameter struct {
	HoursBetweenDBUpdates     int    `json:"hoursbetweendbupdates,omitempty"`
	TimeOfDayToUpdateDB       string `json:"timeofdaytoupdatedb,omitempty"`
	LocalDatabaseThreads      int    `json:"localdatabasethreads,omitempty"`
	CloudHost                 string `json:"cloudhost,omitempty"`
	SeedDBPath                string `json:"seeddbpath,omitempty"`
	MaxNumberOfCloudThreads   string `json:"maxnumberofcloudthreads,omitempty"`
	CloudKeepAliveTimeout     string `json:"cloudkeepalivetimeout,omitempty"`
	CloudServerConnectTimeout string `json:"cloudserverconnecttimeout,omitempty"`
	CloudDBLookupTimeout      string `json:"clouddblookuptimeout,omitempty"`
	ProxyHostIP               string `json:"proxyhostip,omitempty"`
	ProxyPort                 string `json:"proxyport,omitempty"`
	ProxyUsername             string `json:"proxyusername,omitempty"`
	ProxyPassword             string `json:"proxypassword,omitempty"`
	SeedDBSizeLevel           string `json:"seeddbsizelevel,omitempty"`
}

type URLFilteringCategories struct {
	Group      string `json:"group,omitempty"`
	Categories string `json:"categories,omitempty"`
}

type URLFilteringCategorization struct {
	URL            string `json:"url,omitempty"`
	Categorization string `json:"categorization,omitempty"`
}

type URLFilteringCategoryGroups struct {
	CategoryGroups string `json:"categorygroups,omitempty"`
}
