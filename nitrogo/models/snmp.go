package models

// snmp configuration structs
type SNMPManager struct {
	Count              float64 `json:"__count,omitempty"`
	Domain             string  `json:"domain,omitempty"`
	DomainResolveRetry int     `json:"domainresolveretry,omitempty"`
	IP                 string  `json:"ip,omitempty"`
	IPAddress          string  `json:"ipaddress,omitempty"`
	Netmask            string  `json:"netmask,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
}

type SNMPAlarm struct {
	Count              float64 `json:"__count,omitempty"`
	Logging            string  `json:"logging,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NormalValue        int     `json:"normalvalue,omitempty"`
	Severity           string  `json:"severity,omitempty"`
	State              string  `json:"state,omitempty"`
	ThresholdValue     int     `json:"thresholdvalue,omitempty"`
	Time               int     `json:"time,omitempty"`
	Timeout            int     `json:"timeout,omitempty"`
	TrapName           string  `json:"trapname,omitempty"`
}

type SNMPOption struct {
	CustomTrap           string `json:"customtrap,omitempty"`
	NextGenAPIResource   string `json:"_nextgenapiresource,omitempty"`
	PartitionNameInTrap  string `json:"partitionnameintrap,omitempty"`
	SeverityInfoInTrap   string `json:"severityinfointrap,omitempty"`
	SNMPSet              string `json:"snmpset,omitempty"`
	SNMPTrapLogging      string `json:"snmptraplogging,omitempty"`
	SNMPTrapLoggingLevel string `json:"snmptraplogginglevel,omitempty"`
}

type SNMPEngineID struct {
	Count              float64 `json:"__count,omitempty"`
	DefaultEngineID    string  `json:"defaultengineid,omitempty"`
	EngineID           string  `json:"engineid,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	OwnerNode          int     `json:"ownernode,omitempty"`
}

type SNMPMIB struct {
	Contact            string  `json:"contact,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	CustomID           string  `json:"customid,omitempty"`
	Location           string  `json:"location,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	OwnerNode          int     `json:"ownernode,omitempty"`
	SysDesc            string  `json:"sysdesc,omitempty"`
	SysOID             string  `json:"sysoid,omitempty"`
	SysServices        int     `json:"sysservices,omitempty"`
	SysUptime          int     `json:"sysuptime,omitempty"`
}

type SNMPCommunity struct {
	CommunityName      string  `json:"communityname,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Permissions        string  `json:"permissions,omitempty"`
}

type SNMPTrapBinding struct {
	SNMPTrapSNMPUserBinding []interface{} `json:"snmptrap_snmpuser_binding,omitempty"`
	TD                      int           `json:"td,omitempty"`
	TrapClass               string        `json:"trapclass,omitempty"`
	TrapDestination         string        `json:"trapdestination,omitempty"`
	Version                 string        `json:"version,omitempty"`
}

type SNMPUser struct {
	AuthPasswd         string  `json:"authpasswd,omitempty"`
	AuthType           string  `json:"authtype,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	EngineID           string  `json:"engineid,omitempty"`
	Group              string  `json:"group,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	PrivPasswd         string  `json:"privpasswd,omitempty"`
	PrivType           string  `json:"privtype,omitempty"`
	Status             string  `json:"status,omitempty"`
	StorageType        string  `json:"storagetype,omitempty"`
}

type SNMPGroup struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	ReadViewName       string  `json:"readviewname,omitempty"`
	SecurityLevel      string  `json:"securitylevel,omitempty"`
	Status             string  `json:"status,omitempty"`
	StorageType        string  `json:"storagetype,omitempty"`
}

type SNMPTrap struct {
	AllPartitions      string  `json:"allpartitions,omitempty"`
	CommunityName      string  `json:"communityname,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	DestPort           int     `json:"destport,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Severity           string  `json:"severity,omitempty"`
	SrcIP              string  `json:"srcip,omitempty"`
	TD                 int     `json:"td,omitempty"`
	TrapClass          string  `json:"trapclass,omitempty"`
	TrapDestination    string  `json:"trapdestination,omitempty"`
	Version            string  `json:"version,omitempty"`
}

type SNMPTrapSNMPUserBinding struct {
	SecurityLevel   string `json:"securitylevel,omitempty"`
	TD              int    `json:"td,omitempty"`
	TrapClass       string `json:"trapclass,omitempty"`
	TrapDestination string `json:"trapdestination,omitempty"`
	Username        string `json:"username,omitempty"`
	Version         string `json:"version,omitempty"`
}

type SNMPOID struct {
	Count              float64 `json:"__count,omitempty"`
	EntityType         string  `json:"entitytype,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	SNMPOID            string  `json:"Snmpoid,omitempty"`
}

type SNMPView struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Status             string  `json:"status,omitempty"`
	StorageType        string  `json:"storagetype,omitempty"`
	Subtree            string  `json:"subtree,omitempty"`
	TypeField          string  `json:"type,omitempty"`
}
