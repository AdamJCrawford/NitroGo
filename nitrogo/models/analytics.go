// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Analyticsglobalbinding struct {
}

type Analyticsglobalprofilebinding struct {
	Analyticsprofile string `json:"analyticsprofile,omitempty"`
}

type Analyticsprofile struct {
	Name                         string   `json:"name,omitempty"`
	Collectors                   string   `json:"collectors,omitempty"`
	Type                         string   `json:"type,omitempty"`
	Httpclientsidemeasurements   string   `json:"httpclientsidemeasurements,omitempty"`
	Httppagetracking             string   `json:"httppagetracking,omitempty"`
	Httpurl                      string   `json:"httpurl,omitempty"`
	Httphost                     string   `json:"httphost,omitempty"`
	Httpmethod                   string   `json:"httpmethod,omitempty"`
	Httpreferer                  string   `json:"httpreferer,omitempty"`
	Httpuseragent                string   `json:"httpuseragent,omitempty"`
	Httpcookie                   string   `json:"httpcookie,omitempty"`
	Httplocation                 string   `json:"httplocation,omitempty"`
	Urlcategory                  string   `json:"urlcategory,omitempty"`
	Allhttpheaders               string   `json:"allhttpheaders,omitempty"`
	Httpcontenttype              string   `json:"httpcontenttype,omitempty"`
	Httpauthentication           string   `json:"httpauthentication,omitempty"`
	Httpvia                      string   `json:"httpvia,omitempty"`
	Httpxforwardedforheader      string   `json:"httpxforwardedforheader,omitempty"`
	Httpsetcookie                string   `json:"httpsetcookie,omitempty"`
	Httpsetcookie2               string   `json:"httpsetcookie2,omitempty"`
	Httpdomainname               string   `json:"httpdomainname,omitempty"`
	Httpurlquery                 string   `json:"httpurlquery,omitempty"`
	Tcpburstreporting            string   `json:"tcpburstreporting,omitempty"`
	Cqareporting                 string   `json:"cqareporting,omitempty"`
	Integratedcache              string   `json:"integratedcache,omitempty"`
	Grpcstatus                   string   `json:"grpcstatus,omitempty"`
	Outputmode                   string   `json:"outputmode,omitempty"`
	Metrics                      string   `json:"metrics,omitempty"`
	Events                       string   `json:"events,omitempty"`
	Auditlogs                    string   `json:"auditlogs,omitempty"`
	Servemode                    string   `json:"servemode,omitempty"`
	Schemafile                   string   `json:"schemafile,omitempty"`
	Metricsexportfrequency       int      `json:"metricsexportfrequency,omitempty"`
	Analyticsendpointmetadata    string   `json:"analyticsendpointmetadata,omitempty"`
	Dataformatfile               string   `json:"dataformatfile,omitempty"`
	Topn                         string   `json:"topn,omitempty"`
	Httpcustomheaders            []string `json:"httpcustomheaders,omitempty"`
	Managementlog                []string `json:"managementlog,omitempty"`
	Analyticsauthtoken           string   `json:"analyticsauthtoken,omitempty"`
	Analyticsendpointurl         string   `json:"analyticsendpointurl,omitempty"`
	Analyticsendpointcontenttype string   `json:"analyticsendpointcontenttype,omitempty"`
	Refcnt                       string   `json:"refcnt,omitempty"`
	Nextgenapiresource           string   `json:"_nextgenapiresource,omitempty"`
}

type Analyticsglobalanalyticsprofilebinding struct {
	Analyticsprofile string `json:"analyticsprofile,omitempty"`
}
