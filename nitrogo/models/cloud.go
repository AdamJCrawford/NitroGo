// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Cloudservice struct {
	Response string `json:"response,omitempty"`
}

type Cloudvserverip struct {
	Ipaddress          string `json:"ipaddress,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Cloudautoscalegroup struct {
	Name               string `json:"name,omitempty"`
	Azcount            string `json:"azcount,omitempty"`
	Aznames            string `json:"aznames,omitempty"`
	Graceful           string `json:"graceful,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Cloudcredential struct {
	Tenantidentifier   string `json:"tenantidentifier,omitempty"`
	Applicationid      string `json:"applicationid,omitempty"`
	Applicationsecret  string `json:"applicationsecret,omitempty"`
	Isset              string `json:"isset,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Cloudngsparameter struct {
	Blockonallowedngstktprof   string `json:"blockonallowedngstktprof,omitempty"`
	Allowedudtversion          string `json:"allowedudtversion,omitempty"`
	Csvserverticketingdecouple string `json:"csvserverticketingdecouple,omitempty"`
	Allowdtls12                string `json:"allowdtls12,omitempty"`
	Nextgenapiresource         string `json:"_nextgenapiresource,omitempty"`
}

type Cloudparaminternal struct {
	Nonftumode         string `json:"nonftumode,omitempty"`
	Iamperm            string `json:"iamperm,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Cloudallowedngsticketprofile struct {
	Name               string `json:"name,omitempty"`
	Creator            string `json:"creator,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Cloudawsparam struct {
	Rolearn            string `json:"rolearn,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Cloudparameter struct {
	Controllerfqdn          string `json:"controllerfqdn,omitempty"`
	Controllerport          int    `json:"controllerport,omitempty"`
	Instanceid              string `json:"instanceid,omitempty"`
	Customerid              string `json:"customerid,omitempty"`
	Resourcelocation        string `json:"resourcelocation,omitempty"`
	Activationcode          string `json:"activationcode,omitempty"`
	Deployment              string `json:"deployment,omitempty"`
	Connectorresidence      string `json:"connectorresidence,omitempty"`
	Controlconnectionstatus string `json:"controlconnectionstatus,omitempty"`
	Nextgenapiresource      string `json:"_nextgenapiresource,omitempty"`
}

type Cloudprofile struct {
	Name                     string `json:"name,omitempty"`
	Type                     string `json:"type,omitempty"`
	Vservername              string `json:"vservername,omitempty"`
	Servicetype              string `json:"servicetype,omitempty"`
	Ipaddress                string `json:"ipaddress,omitempty"`
	Port                     int    `json:"port,omitempty"`
	Servicegroupname         string `json:"servicegroupname,omitempty"`
	Boundservicegroupsvctype string `json:"boundservicegroupsvctype,omitempty"`
	Vsvrbindsvcport          int    `json:"vsvrbindsvcport,omitempty"`
	Graceful                 string `json:"graceful,omitempty"`
	Delay                    int    `json:"delay,omitempty"`
	Azuretagname             string `json:"azuretagname,omitempty"`
	Azuretagvalue            string `json:"azuretagvalue,omitempty"`
	Azurepollperiod          int    `json:"azurepollperiod,omitempty"`
	Nextgenapiresource       string `json:"_nextgenapiresource,omitempty"`
}
