package models

// analytics configuration structs
type AnalyticsglobalAnalyticsprofileBinding struct {
	Analyticsprofile string `json:"analyticsprofile,omitempty"`
}

type AnalyticsglobalBinding struct {
	AnalyticsglobalAnalyticsprofileBinding []interface{} `json:"analyticsglobal_analyticsprofile_binding,omitempty"`
}

type Analyticsprofile struct {
	Allhttpheaders               string   `json:"allhttpheaders,omitempty"`
	Analyticsauthtoken           string   `json:"analyticsauthtoken,omitempty"`
	Analyticsendpointcontenttype string   `json:"analyticsendpointcontenttype,omitempty"`
	Analyticsendpointmetadata    string   `json:"analyticsendpointmetadata,omitempty"`
	Analyticsendpointurl         string   `json:"analyticsendpointurl,omitempty"`
	Auditlogs                    string   `json:"auditlogs,omitempty"`
	Collectors                   string   `json:"collectors,omitempty"`
	Count                        float64  `json:"__count,omitempty"`
	Cqareporting                 string   `json:"cqareporting,omitempty"`
	Dataformatfile               string   `json:"dataformatfile,omitempty"`
	Events                       string   `json:"events,omitempty"`
	Grpcstatus                   string   `json:"grpcstatus,omitempty"`
	Httpauthentication           string   `json:"httpauthentication,omitempty"`
	Httpclientsidemeasurements   string   `json:"httpclientsidemeasurements,omitempty"`
	Httpcontenttype              string   `json:"httpcontenttype,omitempty"`
	Httpcookie                   string   `json:"httpcookie,omitempty"`
	Httpcustomheaders            []string `json:"httpcustomheaders,omitempty"`
	Httpdomainname               string   `json:"httpdomainname,omitempty"`
	Httphost                     string   `json:"httphost,omitempty"`
	Httplocation                 string   `json:"httplocation,omitempty"`
	Httpmethod                   string   `json:"httpmethod,omitempty"`
	Httppagetracking             string   `json:"httppagetracking,omitempty"`
	Httpreferer                  string   `json:"httpreferer,omitempty"`
	Httpsetcookie                string   `json:"httpsetcookie,omitempty"`
	Httpsetcookie2               string   `json:"httpsetcookie2,omitempty"`
	Httpurl                      string   `json:"httpurl,omitempty"`
	Httpurlquery                 string   `json:"httpurlquery,omitempty"`
	Httpuseragent                string   `json:"httpuseragent,omitempty"`
	Httpvia                      string   `json:"httpvia,omitempty"`
	Httpxforwardedforheader      string   `json:"httpxforwardedforheader,omitempty"`
	Integratedcache              string   `json:"integratedcache,omitempty"`
	Managementlog                []string `json:"managementlog,omitempty"`
	Metrics                      string   `json:"metrics,omitempty"`
	Metricsexportfrequency       int      `json:"metricsexportfrequency,omitempty"`
	Name                         string   `json:"name,omitempty"`
	Nextgenapiresource           string   `json:"_nextgenapiresource,omitempty"`
	Outputmode                   string   `json:"outputmode,omitempty"`
	Refcnt                       int      `json:"refcnt,omitempty"`
	Schemafile                   string   `json:"schemafile,omitempty"`
	Servemode                    string   `json:"servemode,omitempty"`
	Tcpburstreporting            string   `json:"tcpburstreporting,omitempty"`
	Topn                         string   `json:"topn,omitempty"`
	TypeField                    string   `json:"type,omitempty"`
	Urlcategory                  string   `json:"urlcategory,omitempty"`
}
