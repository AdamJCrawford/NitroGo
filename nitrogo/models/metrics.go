package models

// metrics configuration structs
type MetricsprofileAuthenticationvserverBinding struct {
	Entityname string `json:"entityname,omitempty"`
	Entitytype string `json:"entitytype,omitempty"`
	Name       string `json:"name,omitempty"`
}

type MetricsprofileVpnvserverBinding struct {
	Entityname string `json:"entityname,omitempty"`
	Entitytype string `json:"entitytype,omitempty"`
	Name       string `json:"name,omitempty"`
}

type MetricsprofileGslbvserverBinding struct {
	Entityname string `json:"entityname,omitempty"`
	Entitytype string `json:"entitytype,omitempty"`
	Name       string `json:"name,omitempty"`
}

type MetricsprofileBinding struct {
	MetricsprofileAuthenticationvserverBinding []interface{} `json:"metricsprofile_authenticationvserver_binding,omitempty"`
	MetricsprofileCrvserverBinding             []interface{} `json:"metricsprofile_crvserver_binding,omitempty"`
	MetricsprofileCsvserverBinding             []interface{} `json:"metricsprofile_csvserver_binding,omitempty"`
	MetricsprofileGslbvserverBinding           []interface{} `json:"metricsprofile_gslbvserver_binding,omitempty"`
	MetricsprofileLbvserverBinding             []interface{} `json:"metricsprofile_lbvserver_binding,omitempty"`
	MetricsprofileServiceBinding               []interface{} `json:"metricsprofile_service_binding,omitempty"`
	MetricsprofileServicegroupBinding          []interface{} `json:"metricsprofile_servicegroup_binding,omitempty"`
	MetricsprofileUservserverBinding           []interface{} `json:"metricsprofile_uservserver_binding,omitempty"`
	MetricsprofileVpnvserverBinding            []interface{} `json:"metricsprofile_vpnvserver_binding,omitempty"`
	Name                                       string        `json:"name,omitempty"`
}

type Metricsprofile struct {
	Collector              string  `json:"collector,omitempty"`
	Count                  float64 `json:"__count,omitempty"`
	Metrics                string  `json:"metrics,omitempty"`
	Metricsauthtoken       string  `json:"metricsauthtoken,omitempty"`
	Metricsendpointurl     string  `json:"metricsendpointurl,omitempty"`
	Metricsexportfrequency int     `json:"metricsexportfrequency,omitempty"`
	Name                   string  `json:"name,omitempty"`
	Nextgenapiresource     string  `json:"_nextgenapiresource,omitempty"`
	Outputmode             string  `json:"outputmode,omitempty"`
	Refcnt                 int     `json:"refcnt,omitempty"`
	Schemafile             string  `json:"schemafile,omitempty"`
	Servemode              string  `json:"servemode,omitempty"`
}

type MetricsprofileCrvserverBinding struct {
	Entityname string `json:"entityname,omitempty"`
	Entitytype string `json:"entitytype,omitempty"`
	Name       string `json:"name,omitempty"`
}

type MetricsprofileLbvserverBinding struct {
	Entityname string `json:"entityname,omitempty"`
	Entitytype string `json:"entitytype,omitempty"`
	Name       string `json:"name,omitempty"`
}

type MetricsprofileCsvserverBinding struct {
	Entityname string `json:"entityname,omitempty"`
	Entitytype string `json:"entitytype,omitempty"`
	Name       string `json:"name,omitempty"`
}

type MetricsprofileServicegroupBinding struct {
	Entityname string `json:"entityname,omitempty"`
	Entitytype string `json:"entitytype,omitempty"`
	Name       string `json:"name,omitempty"`
}

type MetricsprofileServiceBinding struct {
	Entityname string `json:"entityname,omitempty"`
	Entitytype string `json:"entitytype,omitempty"`
	Name       string `json:"name,omitempty"`
}

type MetricsprofileUservserverBinding struct {
	Entityname string `json:"entityname,omitempty"`
	Entitytype string `json:"entitytype,omitempty"`
	Name       string `json:"name,omitempty"`
}
