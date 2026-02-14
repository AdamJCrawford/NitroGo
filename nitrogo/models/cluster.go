// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Clusternodegroup struct {
	Name                     string `json:"name,omitempty"`
	Strict                   string `json:"strict,omitempty"`
	Sticky                   string `json:"sticky,omitempty"`
	State                    string `json:"state,omitempty"`
	Priority                 int    `json:"priority,omitempty"`
	Currentnodemask          string `json:"currentnodemask,omitempty"`
	Backupnodemask           string `json:"backupnodemask,omitempty"`
	Boundedentitiescntfrompe string `json:"boundedentitiescntfrompe,omitempty"`
	Activelist               string `json:"activelist,omitempty"`
	Backuplist               string `json:"backuplist,omitempty"`
	Nextgenapiresource       string `json:"_nextgenapiresource,omitempty"`
}

type Clusternodegroupauthenticationvserverbinding struct {
	Vserver string `json:"vserver,omitempty"`
	Name    string `json:"name,omitempty"`
}

type Clusternodegroupcsvserverbinding struct {
	Vserver string `json:"vserver,omitempty"`
	Name    string `json:"name,omitempty"`
}

type Clusternodegrouplimitidentifierbinding struct {
	Identifiername string `json:"identifiername,omitempty"`
	Name           string `json:"name,omitempty"`
}

type Clusternodegroupvserverbinding struct {
	Vserver string `json:"vserver,omitempty"`
	Name    string `json:"name,omitempty"`
}

type Clustersyncfailures struct {
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Clusterinstance struct {
	Clid                       int    `json:"clid,omitempty"`
	Deadinterval               int    `json:"deadinterval,omitempty"`
	Hellointerval              int    `json:"hellointerval,omitempty"`
	Preemption                 string `json:"preemption,omitempty"`
	Quorumtype                 string `json:"quorumtype,omitempty"`
	Inc                        string `json:"inc,omitempty"`
	Processlocal               string `json:"processlocal,omitempty"`
	Retainconnectionsoncluster string `json:"retainconnectionsoncluster,omitempty"`
	Backplanebasedview         string `json:"backplanebasedview,omitempty"`
	Syncstatusstrictmode       string `json:"syncstatusstrictmode,omitempty"`
	Dfdretainl2params          string `json:"dfdretainl2params,omitempty"`
	Clusterproxyarp            string `json:"clusterproxyarp,omitempty"`
	Secureheartbeats           string `json:"secureheartbeats,omitempty"`
	Nodegroup                  string `json:"nodegroup,omitempty"`
	Adminstate                 string `json:"adminstate,omitempty"`
	Propstate                  string `json:"propstate,omitempty"`
	Validmtu                   string `json:"validmtu,omitempty"`
	Heterogeneousflag          string `json:"heterogeneousflag,omitempty"`
	Operationalstate           string `json:"operationalstate,omitempty"`
	Status                     string `json:"status,omitempty"`
	Rsskeymismatch             string `json:"rsskeymismatch,omitempty"`
	Penummismatch              string `json:"penummismatch,omitempty"`
	Nodegroupstatewarning      string `json:"nodegroupstatewarning,omitempty"`
	Licensemismatch            string `json:"licensemismatch,omitempty"`
	Jumbonotsupported          string `json:"jumbonotsupported,omitempty"`
	Clustertunnelmodemismatch  string `json:"clustertunnelmodemismatch,omitempty"`
	Clusternoheartbeatonnode   string `json:"clusternoheartbeatonnode,omitempty"`
	Clusternolinksetmbf        string `json:"clusternolinksetmbf,omitempty"`
	Clusternospottedip         string `json:"clusternospottedip,omitempty"`
	Clusterclipfailure         string `json:"clusterclipfailure,omitempty"`
	Clusterhbhmacerrordetected string `json:"clusterhbhmacerrordetected,omitempty"`
	Nodepenummismatch          string `json:"nodepenummismatch,omitempty"`
	Operationalpropstate       string `json:"operationalpropstate,omitempty"`
	Nextgenapiresource         string `json:"_nextgenapiresource,omitempty"`
}

type Clusterinstancenodebinding struct {
	Nodeid                     uint32 `json:"nodeid,omitempty"`
	Ipaddress                  string `json:"ipaddress,omitempty"`
	Health                     string `json:"health,omitempty"`
	Clusterhealth              string `json:"clusterhealth,omitempty"`
	Effectivestate             string `json:"effectivestate,omitempty"`
	Masterstate                string `json:"masterstate,omitempty"`
	State                      string `json:"state,omitempty"`
	Isconfigurationcoordinator bool   `json:"isconfigurationcoordinator,omitempty"`
	Islocalnode                bool   `json:"islocalnode,omitempty"`
	Nodersskeymismatch         bool   `json:"nodersskeymismatch,omitempty"`
	Nodelicensemismatch        bool   `json:"nodelicensemismatch,omitempty"`
	Nodejumbonotsupported      bool   `json:"nodejumbonotsupported,omitempty"`
	Clid                       uint32 `json:"clid,omitempty"`
}

type Clusternodegroupcrvserverbinding struct {
	Vserver string `json:"vserver,omitempty"`
	Name    string `json:"name,omitempty"`
}

type Clusternodegroupnslimitidentifierbinding struct {
	Identifiername string `json:"identifiername,omitempty"`
	Name           string `json:"name,omitempty"`
}

type Clusterfiles struct {
	Mode []string `json:"mode,omitempty"`
}

type Clusterinstancebinding struct {
	Clid int `json:"clid,omitempty"`
}

type Clusternodegroupnodebinding struct {
	Node uint32 `json:"node,omitempty"`
	Name string `json:"name,omitempty"`
}

type Clusternodegroupgslbvserverbinding struct {
	Vserver string `json:"vserver,omitempty"`
	Name    string `json:"name,omitempty"`
}

type Clusternodegrouplbvserverbinding struct {
	Vserver string `json:"vserver,omitempty"`
	Name    string `json:"name,omitempty"`
}

type Clusternodegroupvpnvserverbinding struct {
	Vserver string `json:"vserver,omitempty"`
	Name    string `json:"name,omitempty"`
}

type Cluster struct {
	Clip     string `json:"clip,omitempty"`
	Password string `json:"password,omitempty"`
}

type Clusternode struct {
	Nodeid                     int    `json:"nodeid"`
	Ipaddress                  string `json:"ipaddress,omitempty"`
	State                      string `json:"state,omitempty"`
	Backplane                  string `json:"backplane,omitempty"`
	Priority                   int    `json:"priority,omitempty"`
	Nodegroup                  string `json:"nodegroup,omitempty"`
	Delay                      int    `json:"delay,omitempty"`
	Tunnelmode                 string `json:"tunnelmode,omitempty"`
	Clearnodegroupconfig       string `json:"clearnodegroupconfig,omitempty"`
	Force                      bool   `json:"force,omitempty"`
	Clusterhealth              string `json:"clusterhealth,omitempty"`
	Effectivestate             string `json:"effectivestate,omitempty"`
	Operationalsyncstate       string `json:"operationalsyncstate,omitempty"`
	Syncfailurereason          string `json:"syncfailurereason,omitempty"`
	Masterstate                string `json:"masterstate,omitempty"`
	Health                     string `json:"health,omitempty"`
	Syncstate                  string `json:"syncstate,omitempty"`
	Isconfigurationcoordinator string `json:"isconfigurationcoordinator,omitempty"`
	Islocalnode                string `json:"islocalnode,omitempty"`
	Nodersskeymismatch         string `json:"nodersskeymismatch,omitempty"`
	Nodelicensemismatch        string `json:"nodelicensemismatch,omitempty"`
	Nodejumbonotsupported      string `json:"nodejumbonotsupported,omitempty"`
	Nodelist                   string `json:"nodelist,omitempty"`
	Ifaceslist                 string `json:"ifaceslist,omitempty"`
	Enabledifaces              string `json:"enabledifaces,omitempty"`
	Disabledifaces             string `json:"disabledifaces,omitempty"`
	Partialfailifaces          string `json:"partialfailifaces,omitempty"`
	Hamonifaces                string `json:"hamonifaces,omitempty"`
	Name                       string `json:"name,omitempty"`
	Cfgflags                   string `json:"cfgflags,omitempty"`
	Routemonitor               string `json:"routemonitor,omitempty"`
	Netmask                    string `json:"netmask,omitempty"`
	Nextgenapiresource         string `json:"_nextgenapiresource,omitempty"`
}

type Clusternodegroupgslbsitebinding struct {
	Gslbsite string `json:"gslbsite,omitempty"`
	Name     string `json:"name,omitempty"`
}

type Clusternodegroupservicebinding struct {
	Service string `json:"service,omitempty"`
	Name    string `json:"name,omitempty"`
}

type Clustersync struct {
}

type Clusterinstanceclusternodebinding struct {
	Nodeid                     int    `json:"nodeid,omitempty"`
	Ipaddress                  string `json:"ipaddress,omitempty"`
	Health                     string `json:"health,omitempty"`
	Clusterhealth              string `json:"clusterhealth,omitempty"`
	Effectivestate             string `json:"effectivestate,omitempty"`
	Masterstate                string `json:"masterstate,omitempty"`
	State                      string `json:"state,omitempty"`
	Isconfigurationcoordinator bool   `json:"isconfigurationcoordinator,omitempty"`
	Islocalnode                bool   `json:"islocalnode,omitempty"`
	Nodersskeymismatch         bool   `json:"nodersskeymismatch,omitempty"`
	Nodelicensemismatch        bool   `json:"nodelicensemismatch,omitempty"`
	Nodejumbonotsupported      bool   `json:"nodejumbonotsupported,omitempty"`
	Clid                       int    `json:"clid,omitempty"`
}

type Clusternodegroupclusternodebinding struct {
	Node int    `json:"node,omitempty"`
	Name string `json:"name,omitempty"`
}

type Clusternodegroupidentifierbinding struct {
	Identifiername string `json:"identifiername,omitempty"`
	Name           string `json:"name,omitempty"`
}

type Clusternodegroupsitebinding struct {
	Gslbsite string `json:"gslbsite,omitempty"`
	Name     string `json:"name,omitempty"`
}

type Clusternodegroupstreamidentifierbinding struct {
	Identifiername string `json:"identifiername,omitempty"`
	Name           string `json:"name,omitempty"`
}

type Clusterpropstatus struct {
	Nodeid             int    `json:"nodeid,omitempty"`
	Numpropcmdfailed   string `json:"numpropcmdfailed,omitempty"`
	Cmdstrs            string `json:"cmdstrs,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Clusternodebinding struct {
	Nodeid int `json:"nodeid,omitempty"`
}

type Clusternoderoutemonitorbinding struct {
	Routemonitor  string `json:"routemonitor,omitempty"`
	Netmask       string `json:"netmask,omitempty"`
	Routemonstate int    `json:"routemonstate,omitempty"`
	Nodeid        int    `json:"nodeid,omitempty"`
}

type Clusternodegroupbinding struct {
	Name string `json:"name,omitempty"`
}
