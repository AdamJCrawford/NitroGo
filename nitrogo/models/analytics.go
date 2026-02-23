package models

// analytics configuration structs
type AnalyticsGlobalAnalyticsProfileBinding struct {
	AnalyticsProfile string `json:"analyticsprofile,omitempty"`
}

type AnalyticsGlobalBinding struct {
	AnalyticsGlobalAnalyticsProfileBinding []interface{} `json:"analyticsglobal_analyticsprofile_binding,omitempty"`
}

type AnalyticsProfile struct {
	AllHTTPHeaders               string   `json:"allhttpheaders,omitempty"`
	AnalyticsAuthToken           string   `json:"analyticsauthtoken,omitempty"`
	AnalyticsEndpointContentType string   `json:"analyticsendpointcontenttype,omitempty"`
	AnalyticsEndpointMetadata    string   `json:"analyticsendpointmetadata,omitempty"`
	AnalyticsEndpointURL         string   `json:"analyticsendpointurl,omitempty"`
	AuditLogs                    string   `json:"auditlogs,omitempty"`
	Collectors                   string   `json:"collectors,omitempty"`
	Count                        float64  `json:"__count,omitempty"`
	CQAReporting                 string   `json:"cqareporting,omitempty"`
	DataFormatFile               string   `json:"dataformatfile,omitempty"`
	Events                       string   `json:"events,omitempty"`
	GRPCStatus                   string   `json:"grpcstatus,omitempty"`
	HTTPAuthentication           string   `json:"httpauthentication,omitempty"`
	HTTPClientSideMeasurements   string   `json:"httpclientsidemeasurements,omitempty"`
	HTTPContentType              string   `json:"httpcontenttype,omitempty"`
	HTTPCookie                   string   `json:"httpcookie,omitempty"`
	HTTPCustomHeaders            []string `json:"httpcustomheaders,omitempty"`
	HTTPDomainName               string   `json:"httpdomainname,omitempty"`
	HTTPHost                     string   `json:"httphost,omitempty"`
	HTTPLocation                 string   `json:"httplocation,omitempty"`
	HTTPMethod                   string   `json:"httpmethod,omitempty"`
	HTTPPageTracking             string   `json:"httppagetracking,omitempty"`
	HTTPReferer                  string   `json:"httpreferer,omitempty"`
	HTTPSetCookie                string   `json:"httpsetcookie,omitempty"`
	HTTPSetCookie2               string   `json:"httpsetcookie2,omitempty"`
	HTTPURL                      string   `json:"httpurl,omitempty"`
	HTTPURLQuery                 string   `json:"httpurlquery,omitempty"`
	HTTPUserAgent                string   `json:"httpuseragent,omitempty"`
	HTTPVia                      string   `json:"httpvia,omitempty"`
	HTTPXForwardedForHeader      string   `json:"httpxforwardedforheader,omitempty"`
	IntegratedCache              string   `json:"integratedcache,omitempty"`
	ManagementLog                []string `json:"managementlog,omitempty"`
	Metrics                      string   `json:"metrics,omitempty"`
	MetricsExportFrequency       int      `json:"metricsexportfrequency,omitempty"`
	Name                         string   `json:"name,omitempty"`
	NextGenAPIResource           string   `json:"_nextgenapiresource,omitempty"`
	OutputMode                   string   `json:"outputmode,omitempty"`
	RefCnt                       int      `json:"refcnt,omitempty"`
	SchemaFile                   string   `json:"schemafile,omitempty"`
	ServeMode                    string   `json:"servemode,omitempty"`
	TCPBurstReporting            string   `json:"tcpburstreporting,omitempty"`
	TopN                         string   `json:"topn,omitempty"`
	TypeField                    string   `json:"type,omitempty"`
	URLCategory                  string   `json:"urlcategory,omitempty"`
}
