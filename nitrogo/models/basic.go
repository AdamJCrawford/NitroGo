// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2025

package models

type ServiceGroupBinding struct {
	ServiceGroupName                                 string                                             `json:"servicegroupname,omitempty"`
	ServiceGroupLBMonitorBinding                     []ServiceGroupLBMonitorBinding                     `json:"servicegroup_lbmonitor_binding,omitempty"`
	ServiceGroupServiceGroupEntityMonBindingsBinding []ServiceGroupServiceGroupEntityMonBindingsBinding `json:"servicegroup_servicegroupentitymonbindings_binding,omitempty"`
	ServiceGroupServiceGroupMemberBinding            []ServiceGroupServiceGroupMemberBinding            `json:"servicegroup_servicegroupmember_binding,omitempty"`
}

type ServiceGroupServiceGroupEntityMonBindingsBinding struct {
	CustomServerid             string `json:"customserverid,omitempty"`
	Dbsttl                     int    `json:"dbsttl,omitempty"`
	Hashid                     int    `json:"hashid,omitempty"`
	LastResponse               string `json:"lastresponse,omitempty"`
	MonitorName                string `json:"monitor_name,omitempty"`
	MonitorCurrentFailedProbes string `json:"monitorcurrentfailedprobes,omitempty"`
	MonitorState               string `json:"monitor_state,omitempty"`
	MonitorTotalFailedProbes   string `json:"monitortotalfailedprobes,omitempty"`
	MonitorTotalProbes         string `json:"monitortotalprobes,omitempty"`
	NameServer                 string `json:"nameserver,omitempty"`
	Order                      int    `json:"order,omitempty"`
	Passive                    bool   `json:"passive,omitempty"`
	Port                       int    `json:"port,omitempty"`
	ResponseTime               string `json:"responsetime,omitempty"`
	Serverid                   int    `json:"serverid,omitempty"`
	ServiceGroupEntName2       string `json:"servicegroupentname2,omitempty"`
	ServiceGroupName           string `json:"servicegroupname,omitempty"`
	State                      string `json:"state,omitempty"`
	Weight                     int    `json:"weight,omitempty"`
}

type ServiceGroupServiceGroupMemberBinding struct {
	CustomServerid            string `json:"customserverid,omitempty"`
	Dbsttl                    int    `json:"dbsttl,omitempty"`
	Delay                     int    `json:"delay,omitempty"`
	Graceful                  string `json:"graceful,omitempty"`
	Hashid                    string `json:"hashid,omitempty"`
	IP                        string `json:"ip,omitempty"`
	NameServer                string `json:"nameserver,omitempty"`
	Order                     int    `json:"order,omitempty"`
	OrderStr                  string `json:"orderstr,omitempty"`
	Port                      int    `json:"port,omitempty"`
	Serverid                  string `json:"serverid,omitempty"`
	ServerName                string `json:"servername,omitempty"`
	ServiceGroupName          string `json:"servicegroupname,omitempty"`
	State                     string `json:"state,omitempty"`
	StateChangeTimeSec        string `json:"statechangetimesec,omitempty"`
	SvcitmPriority            int    `json:"svcitmpriority,omitempty"`
	SvrState                  string `json:"svrstate,omitempty"`
	TicksSinceLastStateChange string `json:"tickssincelaststatechange,omitempty"`
	TROFSDelay                string `json:"trofsdelay,omitempty"`
	TROFSReason               string `json:"trofsreason,omitempty"`
	Weight                    string `json:"weight,omitempty"`
}

type ServiceGroupLBMonitorBinding struct {
	CustomServerid   string `json:"customserverid,omitempty"`
	Dbsttl           int    `json:"dbsttl,omitempty"`
	Hashid           int    `json:"hashid,omitempty"`
	MonitorName      string `json:"monitor_name,omitempty"`
	MonState         string `json:"monstate,omitempty"`
	MonWeight        string `json:"monweight,omitempty"`
	NameServer       string `json:"nameserver,omitempty"`
	Order            int    `json:"order,omitempty"`
	Passive          bool   `json:"passive,omitempty"`
	Port             int    `json:"port,omitempty"`
	Serverid         int    `json:"serverid,omitempty"`
	ServiceGroupName string `json:"servicegroupname,omitempty"`
	State            string `json:"state,omitempty"`
	Weight           string `json:"weight,omitempty"`
}
