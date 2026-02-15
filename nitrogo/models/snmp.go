package models

// snmp configuration structs
type Snmpmanager struct {
	Count              float64 `json:"__count,omitempty"`
	Domain             string  `json:"domain,omitempty"`
	Domainresolveretry int     `json:"domainresolveretry,omitempty"`
	Ip                 string  `json:"ip,omitempty"`
	Ipaddress          string  `json:"ipaddress,omitempty"`
	Netmask            string  `json:"netmask,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
}

type Snmpalarm struct {
	Count              float64 `json:"__count,omitempty"`
	Logging            string  `json:"logging,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Normalvalue        int     `json:"normalvalue,omitempty"`
	Severity           string  `json:"severity,omitempty"`
	State              string  `json:"state,omitempty"`
	Thresholdvalue     int     `json:"thresholdvalue,omitempty"`
	Time               int     `json:"time,omitempty"`
	Timeout            int     `json:"timeout,omitempty"`
	Trapname           string  `json:"trapname,omitempty"`
}

type Snmpoption struct {
	Customtrap           string `json:"customtrap,omitempty"`
	Nextgenapiresource   string `json:"_nextgenapiresource,omitempty"`
	Partitionnameintrap  string `json:"partitionnameintrap,omitempty"`
	Severityinfointrap   string `json:"severityinfointrap,omitempty"`
	Snmpset              string `json:"snmpset,omitempty"`
	Snmptraplogging      string `json:"snmptraplogging,omitempty"`
	Snmptraplogginglevel string `json:"snmptraplogginglevel,omitempty"`
}

type Snmpengineid struct {
	Count              float64 `json:"__count,omitempty"`
	Defaultengineid    string  `json:"defaultengineid,omitempty"`
	Engineid           string  `json:"engineid,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Ownernode          int     `json:"ownernode,omitempty"`
}

type Snmpmib struct {
	Contact            string  `json:"contact,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Customid           string  `json:"customid,omitempty"`
	Location           string  `json:"location,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Ownernode          int     `json:"ownernode,omitempty"`
	Sysdesc            string  `json:"sysdesc,omitempty"`
	Sysoid             string  `json:"sysoid,omitempty"`
	Sysservices        int     `json:"sysservices,omitempty"`
	Sysuptime          int     `json:"sysuptime,omitempty"`
}

type Snmpcommunity struct {
	Communityname      string  `json:"communityname,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Permissions        string  `json:"permissions,omitempty"`
}

type SnmptrapBinding struct {
	SnmptrapSnmpuserBinding []interface{} `json:"snmptrap_snmpuser_binding,omitempty"`
	Td                      int           `json:"td,omitempty"`
	Trapclass               string        `json:"trapclass,omitempty"`
	Trapdestination         string        `json:"trapdestination,omitempty"`
	Version                 string        `json:"version,omitempty"`
}

type Snmpuser struct {
	Authpasswd         string  `json:"authpasswd,omitempty"`
	Authtype           string  `json:"authtype,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Engineid           string  `json:"engineid,omitempty"`
	Group              string  `json:"group,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Privpasswd         string  `json:"privpasswd,omitempty"`
	Privtype           string  `json:"privtype,omitempty"`
	Status             string  `json:"status,omitempty"`
	Storagetype        string  `json:"storagetype,omitempty"`
}

type Snmpgroup struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Readviewname       string  `json:"readviewname,omitempty"`
	Securitylevel      string  `json:"securitylevel,omitempty"`
	Status             string  `json:"status,omitempty"`
	Storagetype        string  `json:"storagetype,omitempty"`
}

type Snmptrap struct {
	Allpartitions      string  `json:"allpartitions,omitempty"`
	Communityname      string  `json:"communityname,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Destport           int     `json:"destport,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Severity           string  `json:"severity,omitempty"`
	Srcip              string  `json:"srcip,omitempty"`
	Td                 int     `json:"td,omitempty"`
	Trapclass          string  `json:"trapclass,omitempty"`
	Trapdestination    string  `json:"trapdestination,omitempty"`
	Version            string  `json:"version,omitempty"`
}

type SnmptrapSnmpuserBinding struct {
	Securitylevel   string `json:"securitylevel,omitempty"`
	Td              int    `json:"td,omitempty"`
	Trapclass       string `json:"trapclass,omitempty"`
	Trapdestination string `json:"trapdestination,omitempty"`
	Username        string `json:"username,omitempty"`
	Version         string `json:"version,omitempty"`
}

type Snmpoid struct {
	Count              float64 `json:"__count,omitempty"`
	Entitytype         string  `json:"entitytype,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Snmpoid            string  `json:"Snmpoid,omitempty"`
}

type Snmpview struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Status             string  `json:"status,omitempty"`
	Storagetype        string  `json:"storagetype,omitempty"`
	Subtree            string  `json:"subtree,omitempty"`
	TypeField          string  `json:"type,omitempty"`
}
