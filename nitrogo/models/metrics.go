package models

// metrics configuration structs
type MetricsProfileAuthenticationVServerBinding struct {
	EntityName string `json:"entityname,omitempty"`
	EntityType string `json:"entitytype,omitempty"`
	Name       string `json:"name,omitempty"`
}

type MetricsProfileVPNVServerBinding struct {
	EntityName string `json:"entityname,omitempty"`
	EntityType string `json:"entitytype,omitempty"`
	Name       string `json:"name,omitempty"`
}

type MetricsProfileGSLBVServerBinding struct {
	EntityName string `json:"entityname,omitempty"`
	EntityType string `json:"entitytype,omitempty"`
	Name       string `json:"name,omitempty"`
}

type MetricsProfileBinding struct {
	MetricsProfileAuthenticationVServerBinding []interface{} `json:"metricsprofile_authenticationvserver_binding,omitempty"`
	MetricsProfileCRVServerBinding             []interface{} `json:"metricsprofile_crvserver_binding,omitempty"`
	MetricsProfileCSVServerBinding             []interface{} `json:"metricsprofile_csvserver_binding,omitempty"`
	MetricsProfileGSLBVServerBinding           []interface{} `json:"metricsprofile_gslbvserver_binding,omitempty"`
	MetricsProfileLBVServerBinding             []interface{} `json:"metricsprofile_lbvserver_binding,omitempty"`
	MetricsProfileServiceBinding               []interface{} `json:"metricsprofile_service_binding,omitempty"`
	MetricsProfileServiceGroupBinding          []interface{} `json:"metricsprofile_servicegroup_binding,omitempty"`
	MetricsProfileUserVServerBinding           []interface{} `json:"metricsprofile_uservserver_binding,omitempty"`
	MetricsProfileVPNVServerBinding            []interface{} `json:"metricsprofile_vpnvserver_binding,omitempty"`
	Name                                       string        `json:"name,omitempty"`
}

type MetricsProfile struct {
	Collector              string  `json:"collector,omitempty"`
	Count                  float64 `json:"__count,omitempty"`
	Metrics                string  `json:"metrics,omitempty"`
	MetricsAuthToken       string  `json:"metricsauthtoken,omitempty"`
	MetricsEndpointURL     string  `json:"metricsendpointurl,omitempty"`
	MetricsExportFrequency int     `json:"metricsexportfrequency,omitempty"`
	Name                   string  `json:"name,omitempty"`
	NextGenAPIResource     string  `json:"_nextgenapiresource,omitempty"`
	OutputMode             string  `json:"outputmode,omitempty"`
	RefCnt                 int     `json:"refcnt,omitempty"`
	SchemaFile             string  `json:"schemafile,omitempty"`
	ServeMode              string  `json:"servemode,omitempty"`
}

type MetricsProfileCRVServerBinding struct {
	EntityName string `json:"entityname,omitempty"`
	EntityType string `json:"entitytype,omitempty"`
	Name       string `json:"name,omitempty"`
}

type MetricsProfileLBVServerBinding struct {
	EntityName string `json:"entityname,omitempty"`
	EntityType string `json:"entitytype,omitempty"`
	Name       string `json:"name,omitempty"`
}

type MetricsProfileCSVServerBinding struct {
	EntityName string `json:"entityname,omitempty"`
	EntityType string `json:"entitytype,omitempty"`
	Name       string `json:"name,omitempty"`
}

type MetricsProfileServiceGroupBinding struct {
	EntityName string `json:"entityname,omitempty"`
	EntityType string `json:"entitytype,omitempty"`
	Name       string `json:"name,omitempty"`
}

type MetricsProfileServiceBinding struct {
	EntityName string `json:"entityname,omitempty"`
	EntityType string `json:"entitytype,omitempty"`
	Name       string `json:"name,omitempty"`
}

type MetricsProfileUserVServerBinding struct {
	EntityName string `json:"entityname,omitempty"`
	EntityType string `json:"entitytype,omitempty"`
	Name       string `json:"name,omitempty"`
}
