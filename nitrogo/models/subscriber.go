// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Subscribergxinterface struct {
	Vserver                   string `json:"vserver,omitempty"`
	Service                   string `json:"service,omitempty"`
	Pcrfrealm                 string `json:"pcrfrealm,omitempty"`
	Holdonsubscriberabsence   string `json:"holdonsubscriberabsence,omitempty"`
	Requesttimeout            int    `json:"requesttimeout,omitempty"`
	Requestretryattempts      int    `json:"requestretryattempts,omitempty"`
	Idlettl                   int    `json:"idlettl,omitempty"`
	Revalidationtimeout       int    `json:"revalidationtimeout,omitempty"`
	Healthcheck               string `json:"healthcheck,omitempty"`
	Healthcheckttl            int    `json:"healthcheckttl,omitempty"`
	Cerrequesttimeout         int    `json:"cerrequesttimeout,omitempty"`
	Negativettl               int    `json:"negativettl,omitempty"`
	Negativettllimitedsuccess string `json:"negativettllimitedsuccess,omitempty"`
	Purgesdbongxfailure       string `json:"purgesdbongxfailure,omitempty"`
	Servicepathavp            []int  `json:"servicepathavp,omitempty"`
	Servicepathvendorid       int    `json:"servicepathvendorid,omitempty"`
	Nodeid                    int    `json:"nodeid,omitempty"`
	Svrstate                  string `json:"svrstate,omitempty"`
	Identity                  string `json:"identity,omitempty"`
	Realm                     string `json:"realm,omitempty"`
	Status                    string `json:"status,omitempty"`
	Servicepathinfomode       string `json:"servicepathinfomode,omitempty"`
	Gxreportingavp1           string `json:"gxreportingavp1,omitempty"`
	Gxreportingavp1vendorid   string `json:"gxreportingavp1vendorid,omitempty"`
	Gxreportingavp1type       string `json:"gxreportingavp1type,omitempty"`
	Gxreportingavp2           string `json:"gxreportingavp2,omitempty"`
	Gxreportingavp2vendorid   string `json:"gxreportingavp2vendorid,omitempty"`
	Gxreportingavp2type       string `json:"gxreportingavp2type,omitempty"`
	Gxreportingavp3           string `json:"gxreportingavp3,omitempty"`
	Gxreportingavp3vendorid   string `json:"gxreportingavp3vendorid,omitempty"`
	Gxreportingavp3type       string `json:"gxreportingavp3type,omitempty"`
	Gxreportingavp4           string `json:"gxreportingavp4,omitempty"`
	Gxreportingavp4vendorid   string `json:"gxreportingavp4vendorid,omitempty"`
	Gxreportingavp4type       string `json:"gxreportingavp4type,omitempty"`
	Gxreportingavp5           string `json:"gxreportingavp5,omitempty"`
	Gxreportingavp5vendorid   string `json:"gxreportingavp5vendorid,omitempty"`
	Gxreportingavp5type       string `json:"gxreportingavp5type,omitempty"`
	Nextgenapiresource        string `json:"_nextgenapiresource,omitempty"`
}

type Subscriberparam struct {
	Keytype              string `json:"keytype,omitempty"`
	Interfacetype        string `json:"interfacetype,omitempty"`
	Idlettl              int    `json:"idlettl,omitempty"`
	Idleaction           string `json:"idleaction,omitempty"`
	Ipv6prefixlookuplist []int  `json:"ipv6prefixlookuplist,omitempty"`
	Builtin              string `json:"builtin,omitempty"`
	Feature              string `json:"feature,omitempty"`
	Nextgenapiresource   string `json:"_nextgenapiresource,omitempty"`
}

type Subscriberprofile struct {
	Ip                  string   `json:"ip,omitempty"`
	Vlan                int      `json:"vlan"`
	Subscriberrules     []string `json:"subscriberrules,omitempty"`
	Subscriptionidtype  string   `json:"subscriptionidtype,omitempty"`
	Subscriptionidvalue string   `json:"subscriptionidvalue,omitempty"`
	Servicepath         string   `json:"servicepath,omitempty"`
	Flags               string   `json:"flags,omitempty"`
	Ttl                 string   `json:"ttl,omitempty"`
	Avpdisplaybuffer    string   `json:"avpdisplaybuffer,omitempty"`
	Nextgenapiresource  string   `json:"_nextgenapiresource,omitempty"`
}

type Subscriberradiusinterface struct {
	Listeningservice     string `json:"listeningservice,omitempty"`
	Radiusinterimasstart string `json:"radiusinterimasstart,omitempty"`
	Svrstate             string `json:"svrstate,omitempty"`
	Nextgenapiresource   string `json:"_nextgenapiresource,omitempty"`
}

type Subscribersessions struct {
	Ip                  string `json:"ip,omitempty"`
	Vlan                int    `json:"vlan,omitempty"`
	Nodeid              int    `json:"nodeid,omitempty"`
	Subscriptionidtype  string `json:"subscriptionidtype,omitempty"`
	Subscriptionidvalue string `json:"subscriptionidvalue,omitempty"`
	Subscriberrules     string `json:"subscriberrules,omitempty"`
	Flags               string `json:"flags,omitempty"`
	Ttl                 string `json:"ttl,omitempty"`
	Idlettl             string `json:"idlettl,omitempty"`
	Avpdisplaybuffer    string `json:"avpdisplaybuffer,omitempty"`
	Servicepath         string `json:"servicepath,omitempty"`
	Nextgenapiresource  string `json:"_nextgenapiresource,omitempty"`
}
