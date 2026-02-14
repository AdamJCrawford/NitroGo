// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Snmpview struct {
	Name               string `json:"name,omitempty"`
	Subtree            string `json:"subtree,omitempty"`
	Type               string `json:"type,omitempty"`
	Storagetype        string `json:"storagetype,omitempty"`
	Status             string `json:"status,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Snmpengineid struct {
	Engineid           string `json:"engineid,omitempty"`
	Ownernode          int    `json:"ownernode,omitempty"`
	Defaultengineid    string `json:"defaultengineid,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Snmpgroup struct {
	Name               string `json:"name,omitempty"`
	Securitylevel      string `json:"securitylevel,omitempty"`
	Readviewname       string `json:"readviewname,omitempty"`
	Storagetype        string `json:"storagetype,omitempty"`
	Status             string `json:"status,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Snmpmanager struct {
	Ipaddress          string `json:"ipaddress,omitempty"`
	Netmask            string `json:"netmask,omitempty"`
	Domainresolveretry int    `json:"domainresolveretry,omitempty"`
	Ip                 string `json:"ip,omitempty"`
	Domain             string `json:"domain,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Snmpoid struct {
	Entitytype         string `json:"entitytype,omitempty"`
	Name               string `json:"name,omitempty"`
	Snmpoid            string `json:"Snmpoid,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Snmpoption struct {
	Snmpset              string `json:"snmpset,omitempty"`
	Snmptraplogging      string `json:"snmptraplogging,omitempty"`
	Partitionnameintrap  string `json:"partitionnameintrap,omitempty"`
	Snmptraplogginglevel string `json:"snmptraplogginglevel,omitempty"`
	Severityinfointrap   string `json:"severityinfointrap,omitempty"`
	Nextgenapiresource   string `json:"_nextgenapiresource,omitempty"`
}

type Snmptrap struct {
	Trapclass          string `json:"trapclass,omitempty"`
	Trapdestination    string `json:"trapdestination,omitempty"`
	Version            string `json:"version,omitempty"`
	Td                 int    `json:"td,omitempty"`
	Destport           int    `json:"destport,omitempty"`
	Communityname      string `json:"communityname,omitempty"`
	Srcip              string `json:"srcip,omitempty"`
	Severity           string `json:"severity,omitempty"`
	Allpartitions      string `json:"allpartitions,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Snmptrapbinding struct {
	Trapclass       string `json:"trapclass,omitempty"`
	Trapdestination string `json:"trapdestination,omitempty"`
	Version         string `json:"version,omitempty"`
	Td              int    `json:"td,omitempty"`
}

type Snmpuser struct {
	Name               string `json:"name,omitempty"`
	Group              string `json:"group,omitempty"`
	Authtype           string `json:"authtype,omitempty"`
	Authpasswd         string `json:"authpasswd,omitempty"`
	Privtype           string `json:"privtype,omitempty"`
	Privpasswd         string `json:"privpasswd,omitempty"`
	Engineid           string `json:"engineid,omitempty"`
	Storagetype        string `json:"storagetype,omitempty"`
	Status             string `json:"status,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Snmpalarm struct {
	Trapname           string `json:"trapname,omitempty"`
	Thresholdvalue     int    `json:"thresholdvalue,omitempty"`
	Normalvalue        int    `json:"normalvalue,omitempty"`
	Time               int    `json:"time,omitempty"`
	State              string `json:"state,omitempty"`
	Severity           string `json:"severity,omitempty"`
	Logging            string `json:"logging,omitempty"`
	Timeout            string `json:"timeout,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Snmpcommunity struct {
	Communityname      string `json:"communityname,omitempty"`
	Permissions        string `json:"permissions,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Snmpmib struct {
	Contact            string `json:"contact,omitempty"`
	Name               string `json:"name,omitempty"`
	Location           string `json:"location,omitempty"`
	Customid           string `json:"customid,omitempty"`
	Ownernode          int    `json:"ownernode,omitempty"`
	Sysdesc            string `json:"sysdesc,omitempty"`
	Sysuptime          string `json:"sysuptime,omitempty"`
	Sysservices        string `json:"sysservices,omitempty"`
	Sysoid             string `json:"sysoid,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Snmptrapsnmpuserbinding struct {
	Username        string `json:"username,omitempty"`
	Securitylevel   string `json:"securitylevel,omitempty"`
	Trapclass       string `json:"trapclass,omitempty"`
	Trapdestination string `json:"trapdestination,omitempty"`
	Td              int    `json:"td,omitempty"`
	Version         string `json:"version,omitempty"`
}

type Snmptrapuserbinding struct {
	Username        string `json:"username,omitempty"`
	Securitylevel   string `json:"securitylevel,omitempty"`
	Trapclass       string `json:"trapclass,omitempty"`
	Trapdestination string `json:"trapdestination,omitempty"`
	Td              uint32 `json:"td,omitempty"`
	Version         string `json:"version,omitempty"`
}
