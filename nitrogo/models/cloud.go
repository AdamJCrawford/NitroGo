package models

// cloud configuration structs
type Cloudcredential struct {
	Applicationid      string `json:"applicationid,omitempty"`
	Applicationsecret  string `json:"applicationsecret,omitempty"`
	Isset              int    `json:"isset,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Tenantidentifier   string `json:"tenantidentifier,omitempty"`
}

type Cloudngsparameter struct {
	Allowdtls12                string `json:"allowdtls12,omitempty"`
	Allowedudtversion          string `json:"allowedudtversion,omitempty"`
	Blockonallowedngstktprof   string `json:"blockonallowedngstktprof,omitempty"`
	Csvserverticketingdecouple string `json:"csvserverticketingdecouple,omitempty"`
	Nextgenapiresource         string `json:"_nextgenapiresource,omitempty"`
}

type Cloudawsparam struct {
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Rolearn            string `json:"rolearn,omitempty"`
}

type Cloudparameter struct {
	Activationcode          string `json:"activationcode,omitempty"`
	Connectorresidence      string `json:"connectorresidence,omitempty"`
	Controlconnectionstatus string `json:"controlconnectionstatus,omitempty"`
	Controllerfqdn          string `json:"controllerfqdn,omitempty"`
	Controllerport          int    `json:"controllerport,omitempty"`
	Customerid              string `json:"customerid,omitempty"`
	Deployment              string `json:"deployment,omitempty"`
	Instanceid              string `json:"instanceid,omitempty"`
	Nextgenapiresource      string `json:"_nextgenapiresource,omitempty"`
	Resourcelocation        string `json:"resourcelocation,omitempty"`
}

type Cloudallowedngsticketprofile struct {
	Count              float64 `json:"__count,omitempty"`
	Creator            string  `json:"creator,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
}

type Cloudservice struct {
	Response string `json:"response,omitempty"`
}

type Cloudparaminternal struct {
	Count              float64 `json:"__count,omitempty"`
	Iamperm            string  `json:"iamperm,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Nonftumode         string  `json:"nonftumode,omitempty"`
}

type Cloudprofile struct {
	Azurepollperiod          int     `json:"azurepollperiod,omitempty"`
	Azuretagname             string  `json:"azuretagname,omitempty"`
	Azuretagvalue            string  `json:"azuretagvalue,omitempty"`
	Boundservicegroupsvctype string  `json:"boundservicegroupsvctype,omitempty"`
	Count                    float64 `json:"__count,omitempty"`
	Delay                    int     `json:"delay,omitempty"`
	Graceful                 string  `json:"graceful,omitempty"`
	Ipaddress                string  `json:"ipaddress,omitempty"`
	Name                     string  `json:"name,omitempty"`
	Nextgenapiresource       string  `json:"_nextgenapiresource,omitempty"`
	Port                     int     `json:"port,omitempty"`
	Servicegroupname         string  `json:"servicegroupname,omitempty"`
	Servicetype              string  `json:"servicetype,omitempty"`
	TypeField                string  `json:"type,omitempty"`
	Vservername              string  `json:"vservername,omitempty"`
	Vsvrbindsvcport          int     `json:"vsvrbindsvcport,omitempty"`
}

type Cloudautoscalegroup struct {
	Azcount            int      `json:"azcount,omitempty"`
	Aznames            []string `json:"aznames,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Graceful           string   `json:"graceful,omitempty"`
	Name               string   `json:"name,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
}

type Cloudvserverip struct {
	Count              float64 `json:"__count,omitempty"`
	Ipaddress          string  `json:"ipaddress,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
}
