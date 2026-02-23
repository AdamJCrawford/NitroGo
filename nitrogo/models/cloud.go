package models

// cloud configuration structs
type CloudCredential struct {
	ApplicationID      string `json:"applicationid,omitempty"`
	ApplicationSecret  string `json:"applicationsecret,omitempty"`
	IsSet              int    `json:"isset,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	TenantIdentifier   string `json:"tenantidentifier,omitempty"`
}

type CloudNGSParameter struct {
	AllowDTLS12                string `json:"allowdtls12,omitempty"`
	AllowedUDTVersion          string `json:"allowedudtversion,omitempty"`
	BlockOnAllowedNGSTktProf   string `json:"blockonallowedngstktprof,omitempty"`
	CSVServerTicketingDecouple string `json:"csvserverticketingdecouple,omitempty"`
	NextGenAPIResource         string `json:"_nextgenapiresource,omitempty"`
}

type CloudAWSParam struct {
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	RoleARN            string `json:"rolearn,omitempty"`
}

type CloudParameter struct {
	ActivationCode          string `json:"activationcode,omitempty"`
	ConnectorResidence      string `json:"connectorresidence,omitempty"`
	ControlConnectionStatus string `json:"controlconnectionstatus,omitempty"`
	ControllerFQDN          string `json:"controllerfqdn,omitempty"`
	ControllerPort          int    `json:"controllerport,omitempty"`
	CustomerID              string `json:"customerid,omitempty"`
	Deployment              string `json:"deployment,omitempty"`
	InstanceID              string `json:"instanceid,omitempty"`
	NextGenAPIResource      string `json:"_nextgenapiresource,omitempty"`
	ResourceLocation        string `json:"resourcelocation,omitempty"`
}

type CloudAllowedNGSTicketProfile struct {
	Count              float64 `json:"__count,omitempty"`
	Creator            string  `json:"creator,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
}

type CloudService struct {
	Response string `json:"response,omitempty"`
}

type CloudParamInternal struct {
	Count              float64 `json:"__count,omitempty"`
	IAMPerm            string  `json:"iamperm,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NonFTUMode         string  `json:"nonftumode,omitempty"`
}

type CloudProfile struct {
	AzurePollPeriod          int     `json:"azurepollperiod,omitempty"`
	AzureTagName             string  `json:"azuretagname,omitempty"`
	AzureTagValue            string  `json:"azuretagvalue,omitempty"`
	BoundServiceGroupSvcType string  `json:"boundservicegroupsvctype,omitempty"`
	Count                    float64 `json:"__count,omitempty"`
	Delay                    int     `json:"delay,omitempty"`
	Graceful                 string  `json:"graceful,omitempty"`
	IPAddress                string  `json:"ipaddress,omitempty"`
	Name                     string  `json:"name,omitempty"`
	NextGenAPIResource       string  `json:"_nextgenapiresource,omitempty"`
	Port                     int     `json:"port,omitempty"`
	ServiceGroupName         string  `json:"servicegroupname,omitempty"`
	ServiceType              string  `json:"servicetype,omitempty"`
	TypeField                string  `json:"type,omitempty"`
	VServerName              string  `json:"vservername,omitempty"`
	VSvrBindSvcPort          int     `json:"vsvrbindsvcport,omitempty"`
}

type CloudAutoscaleGroup struct {
	AZCount            int      `json:"azcount,omitempty"`
	AZNames            []string `json:"aznames,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Graceful           string   `json:"graceful,omitempty"`
	Name               string   `json:"name,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
}

type CloudVServerIP struct {
	Count              float64 `json:"__count,omitempty"`
	IPAddress          string  `json:"ipaddress,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
}
