// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Lsnsipalgcall struct {
	Callid             string `json:"callid,omitempty"`
	Nodeid             int    `json:"nodeid,omitempty"`
	Callflags          string `json:"callflags,omitempty"`
	Xlatip             string `json:"xlatip,omitempty"`
	Callrefcount       string `json:"callrefcount,omitempty"`
	Calltimer          string `json:"calltimer,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Lsnsipalgcallbinding struct {
	Callid string `json:"callid,omitempty"`
}

type Lsngroupbinding struct {
	Groupname string `json:"groupname,omitempty"`
}

type Lsngroupipsecalgprofilebinding struct {
	Ipsecalgprofile string `json:"ipsecalgprofile,omitempty"`
	Groupname       string `json:"groupname,omitempty"`
}

type Lsngrouplogprofilebinding struct {
	Logprofilename string `json:"logprofilename,omitempty"`
	Groupname      string `json:"groupname,omitempty"`
}

type Lsngrouplsnrtspalgprofilebinding struct {
	Rtspalgprofilename string `json:"rtspalgprofilename,omitempty"`
	Groupname          string `json:"groupname,omitempty"`
}

type Lsngrouplsnsipalgprofilebinding struct {
	Sipalgprofilename string `json:"sipalgprofilename,omitempty"`
	Groupname         string `json:"groupname,omitempty"`
}

type Lsngroupprofilebinding struct {
	Ipsecalgprofile string `json:"ipsecalgprofile,omitempty"`
	Groupname       string `json:"groupname,omitempty"`
}

type Lsnlogprofile struct {
	Logprofilename     string `json:"logprofilename,omitempty"`
	Logsubscrinfo      string `json:"logsubscrinfo,omitempty"`
	Logcompact         string `json:"logcompact,omitempty"`
	Logipfix           string `json:"logipfix,omitempty"`
	Analyticsprofile   string `json:"analyticsprofile,omitempty"`
	Logsessdeletion    string `json:"logsessdeletion,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Lsnsipalgcallcontrolchannelbinding struct {
	Channelip       string `json:"channelip,omitempty"`
	Channelnatip    string `json:"channelnatip,omitempty"`
	Channelport     int    `json:"channelport,omitempty"`
	Channelnatport  int    `json:"channelnatport,omitempty"`
	Channelprotocol string `json:"channelprotocol,omitempty"`
	Channelflags    int    `json:"channelflags,omitempty"`
	Channeltimeout  int    `json:"channeltimeout,omitempty"`
	Callid          string `json:"callid,omitempty"`
}

type Lsnclientnsacl6binding struct {
	Acl6name   string `json:"acl6name,omitempty"`
	Td         int    `json:"td,omitempty"`
	Clientname string `json:"clientname,omitempty"`
}

type Lsnclientbinding struct {
	Clientname string `json:"clientname,omitempty"`
}

type Lsngrouppoolbinding struct {
	Poolname  string `json:"poolname,omitempty"`
	Groupname string `json:"groupname,omitempty"`
}

type Lsngroupserverbinding struct {
	Pcpserver string `json:"pcpserver,omitempty"`
	Groupname string `json:"groupname,omitempty"`
}

type Lsnrtspalgsessionbinding struct {
	Sessionid string `json:"sessionid,omitempty"`
}

type Lsnstatic struct {
	Name               string `json:"name,omitempty"`
	Transportprotocol  string `json:"transportprotocol,omitempty"`
	Subscrip           string `json:"subscrip,omitempty"`
	Subscrport         int    `json:"subscrport,omitempty"`
	Network6           string `json:"network6,omitempty"`
	Td                 int    `json:"td"`
	Natip              string `json:"natip,omitempty"`
	Natport            int    `json:"natport,omitempty"`
	Destip             string `json:"destip,omitempty"`
	Dsttd              int    `json:"dsttd,omitempty"`
	Nattype            string `json:"nattype,omitempty"`
	Status             string `json:"status,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Lsntransportprofile struct {
	Transportprofilename string `json:"transportprofilename,omitempty"`
	Transportprotocol    string `json:"transportprotocol,omitempty"`
	Sessiontimeout       int    `json:"sessiontimeout,omitempty"`
	Finrsttimeout        int    `json:"finrsttimeout,omitempty"`
	Stuntimeout          int    `json:"stuntimeout,omitempty"`
	Synidletimeout       int    `json:"synidletimeout,omitempty"`
	Portquota            int    `json:"portquota"`
	Sessionquota         int    `json:"sessionquota"`
	Groupsessionlimit    int    `json:"groupsessionlimit"`
	Portpreserveparity   string `json:"portpreserveparity,omitempty"`
	Portpreserverange    string `json:"portpreserverange,omitempty"`
	Syncheck             string `json:"syncheck,omitempty"`
	Nextgenapiresource   string `json:"_nextgenapiresource,omitempty"`
}

type Lsnappsprofilebinding struct {
	Appsprofilename string `json:"appsprofilename,omitempty"`
}

type Lsnclientnetworkbinding struct {
	Network    string `json:"network,omitempty"`
	Netmask    string `json:"netmask,omitempty"`
	Td         int    `json:"td,omitempty"`
	Clientname string `json:"clientname,omitempty"`
}

type Lsngrouphttphdrlogprofilebinding struct {
	Httphdrlogprofilename string `json:"httphdrlogprofilename,omitempty"`
	Groupname             string `json:"groupname,omitempty"`
}

type Lsngrouppcpserverbinding struct {
	Pcpserver string `json:"pcpserver,omitempty"`
	Groupname string `json:"groupname,omitempty"`
}

type Lsngrouprtspalgprofilebinding struct {
	Rtspalgprofilename string `json:"rtspalgprofilename,omitempty"`
	Groupname          string `json:"groupname,omitempty"`
}

type Lsnpoollsnipbinding struct {
	Lsnip     string `json:"lsnip,omitempty"`
	Ownernode int    `json:"ownernode,omitempty"`
	Poolname  string `json:"poolname,omitempty"`
}

type Lsnsipalgprofile struct {
	Sipalgprofilename      string `json:"sipalgprofilename,omitempty"`
	Datasessionidletimeout int    `json:"datasessionidletimeout,omitempty"`
	Sipsessiontimeout      int    `json:"sipsessiontimeout,omitempty"`
	Registrationtimeout    int    `json:"registrationtimeout,omitempty"`
	Sipsrcportrange        string `json:"sipsrcportrange,omitempty"`
	Sipdstportrange        string `json:"sipdstportrange,omitempty"`
	Openregisterpinhole    string `json:"openregisterpinhole,omitempty"`
	Opencontactpinhole     string `json:"opencontactpinhole,omitempty"`
	Openviapinhole         string `json:"openviapinhole,omitempty"`
	Openrecordroutepinhole string `json:"openrecordroutepinhole,omitempty"`
	Siptransportprotocol   string `json:"siptransportprotocol,omitempty"`
	Openroutepinhole       string `json:"openroutepinhole,omitempty"`
	Rport                  string `json:"rport,omitempty"`
	Nextgenapiresource     string `json:"_nextgenapiresource,omitempty"`
}

type Lsngrouplsnappsprofilebinding struct {
	Appsprofilename string `json:"appsprofilename,omitempty"`
	Groupname       string `json:"groupname,omitempty"`
}

type Lsngroup struct {
	Groupname          string `json:"groupname,omitempty"`
	Clientname         string `json:"clientname,omitempty"`
	Nattype            string `json:"nattype,omitempty"`
	Allocpolicy        string `json:"allocpolicy,omitempty"`
	Portblocksize      int    `json:"portblocksize"`
	Logging            string `json:"logging,omitempty"`
	Sessionlogging     string `json:"sessionlogging,omitempty"`
	Sessionsync        string `json:"sessionsync,omitempty"`
	Snmptraplimit      int    `json:"snmptraplimit"`
	Ftp                string `json:"ftp,omitempty"`
	Pptp               string `json:"pptp,omitempty"`
	Sipalg             string `json:"sipalg,omitempty"`
	Rtspalg            string `json:"rtspalg,omitempty"`
	Ip6profile         string `json:"ip6profile,omitempty"`
	Ftpcm              string `json:"ftpcm,omitempty"`
	Groupid            string `json:"groupid,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Lsngroupappsprofilebinding struct {
	Appsprofilename string `json:"appsprofilename,omitempty"`
	Groupname       string `json:"groupname,omitempty"`
}

type Lsngroupsipalgprofilebinding struct {
	Sipalgprofilename string `json:"sipalgprofilename,omitempty"`
	Groupname         string `json:"groupname,omitempty"`
}

type Lsnrtspalgsession struct {
	Sessionid          string `json:"sessionid,omitempty"`
	Nodeid             int    `json:"nodeid,omitempty"`
	Callflags          string `json:"callflags,omitempty"`
	Xlatip             string `json:"xlatip,omitempty"`
	Callrefcount       string `json:"callrefcount,omitempty"`
	Calltimer          string `json:"calltimer,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Lsnappsprofile struct {
	Appsprofilename    string `json:"appsprofilename,omitempty"`
	Transportprotocol  string `json:"transportprotocol,omitempty"`
	Ippooling          string `json:"ippooling,omitempty"`
	Mapping            string `json:"mapping,omitempty"`
	Filtering          string `json:"filtering,omitempty"`
	Tcpproxy           string `json:"tcpproxy,omitempty"`
	Td                 int    `json:"td,omitempty"`
	L2info             string `json:"l2info,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Lsngrouplsnhttphdrlogprofilebinding struct {
	Httphdrlogprofilename string `json:"httphdrlogprofilename,omitempty"`
	Groupname             string `json:"groupname,omitempty"`
}

type Lsngrouptransportprofilebinding struct {
	Transportprofilename string `json:"transportprofilename,omitempty"`
	Groupname            string `json:"groupname,omitempty"`
}

type Lsnrtspalgprofile struct {
	Rtspalgprofilename    string `json:"rtspalgprofilename,omitempty"`
	Rtspidletimeout       int    `json:"rtspidletimeout,omitempty"`
	Rtspportrange         string `json:"rtspportrange,omitempty"`
	Rtsptransportprotocol string `json:"rtsptransportprotocol,omitempty"`
	Nextgenapiresource    string `json:"_nextgenapiresource,omitempty"`
}

type Lsnsipalgcalldatachannelbinding struct {
	Channelip       string `json:"channelip,omitempty"`
	Channelnatip    string `json:"channelnatip,omitempty"`
	Channelport     int    `json:"channelport,omitempty"`
	Channelnatport  int    `json:"channelnatport,omitempty"`
	Channelprotocol string `json:"channelprotocol,omitempty"`
	Channelflags    int    `json:"channelflags,omitempty"`
	Channeltimeout  int    `json:"channeltimeout,omitempty"`
	Callid          string `json:"callid,omitempty"`
}

type Lsnappsprofileappsattributesbinding struct {
	Appsattributesname string `json:"appsattributesname,omitempty"`
	Appsprofilename    string `json:"appsprofilename,omitempty"`
}

type Lsnappsprofilelsnappsattributesbinding struct {
	Appsattributesname string `json:"appsattributesname,omitempty"`
	Appsprofilename    string `json:"appsprofilename,omitempty"`
}

type Lsndeterministicnat struct {
	Clientname         string `json:"clientname,omitempty"`
	Network6           string `json:"network6,omitempty"`
	Subscrip           string `json:"subscrip,omitempty"`
	Td                 int    `json:"td,omitempty"`
	Natip              string `json:"natip,omitempty"`
	Natprefix          string `json:"natprefix,omitempty"`
	Subscrip2          string `json:"subscrip2,omitempty"`
	Natip2             string `json:"natip2,omitempty"`
	Firstport          string `json:"firstport,omitempty"`
	Lastport           string `json:"lastport,omitempty"`
	Srctd              string `json:"srctd,omitempty"`
	Nattype            string `json:"nattype,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Lsngrouplsnlogprofilebinding struct {
	Logprofilename string `json:"logprofilename,omitempty"`
	Groupname      string `json:"groupname,omitempty"`
}

type Lsnpool struct {
	Poolname            string `json:"poolname,omitempty"`
	Nattype             string `json:"nattype,omitempty"`
	Portblockallocation string `json:"portblockallocation,omitempty"`
	Portrealloctimeout  int    `json:"portrealloctimeout"`
	Maxportrealloctmq   int    `json:"maxportrealloctmq"`
	Nextgenapiresource  string `json:"_nextgenapiresource,omitempty"`
}

type Lsnpoolbinding struct {
	Poolname string `json:"poolname,omitempty"`
}

type Lsnclient struct {
	Clientname         string `json:"clientname,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Lsngrouplsnpoolbinding struct {
	Poolname  string `json:"poolname,omitempty"`
	Groupname string `json:"groupname,omitempty"`
}

type Lsngrouplsntransportprofilebinding struct {
	Transportprofilename string `json:"transportprofilename,omitempty"`
	Groupname            string `json:"groupname,omitempty"`
}

type Lsnhttphdrlogprofile struct {
	Httphdrlogprofilename string `json:"httphdrlogprofilename,omitempty"`
	Logurl                string `json:"logurl,omitempty"`
	Logmethod             string `json:"logmethod,omitempty"`
	Logversion            string `json:"logversion,omitempty"`
	Loghost               string `json:"loghost,omitempty"`
	Nextgenapiresource    string `json:"_nextgenapiresource,omitempty"`
}

type Lsnip6profile struct {
	Name               string `json:"name,omitempty"`
	Type               string `json:"type,omitempty"`
	Natprefix          string `json:"natprefix,omitempty"`
	Network6           string `json:"network6,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Lsnparameter struct {
	Memlimit             int    `json:"memlimit,omitempty"`
	Sessionsync          string `json:"sessionsync,omitempty"`
	Subscrsessionremoval string `json:"subscrsessionremoval,omitempty"`
	Memlimitactive       string `json:"memlimitactive,omitempty"`
	Maxmemlimit          string `json:"maxmemlimit,omitempty"`
	Nextgenapiresource   string `json:"_nextgenapiresource,omitempty"`
}

type Lsnrtspalgsessiondatachannelbinding struct {
	Channelip       string `json:"channelip,omitempty"`
	Channelnatip    string `json:"channelnatip,omitempty"`
	Channelport     int    `json:"channelport,omitempty"`
	Channelnatport  int    `json:"channelnatport,omitempty"`
	Channelprotocol string `json:"channelprotocol,omitempty"`
	Channelflags    int    `json:"channelflags,omitempty"`
	Channeltimeout  int    `json:"channeltimeout,omitempty"`
	Sessionid       string `json:"sessionid,omitempty"`
}

type Lsnsession struct {
	Nattype            string `json:"nattype,omitempty"`
	Clientname         string `json:"clientname,omitempty"`
	Network            string `json:"network,omitempty"`
	Netmask            string `json:"netmask,omitempty"`
	Network6           string `json:"network6,omitempty"`
	Td                 int    `json:"td,omitempty"`
	Natip              string `json:"natip,omitempty"`
	Nodeid             int    `json:"nodeid,omitempty"`
	Natport2           int    `json:"natport2,omitempty"`
	Natprefix          string `json:"natprefix,omitempty"`
	Subscrip           string `json:"subscrip,omitempty"`
	Subscrport         string `json:"subscrport,omitempty"`
	Destip             string `json:"destip,omitempty"`
	Destport           string `json:"destport,omitempty"`
	Natport            string `json:"natport,omitempty"`
	Transportprotocol  string `json:"transportprotocol,omitempty"`
	Sessionestdir      string `json:"sessionestdir,omitempty"`
	Dsttd              string `json:"dsttd,omitempty"`
	Srctd              string `json:"srctd,omitempty"`
	Ipv6address        string `json:"ipv6address,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Lsnclientnsaclbinding struct {
	Aclname    string `json:"aclname,omitempty"`
	Td         int    `json:"td,omitempty"`
	Clientname string `json:"clientname,omitempty"`
}

type Lsnappsattributes struct {
	Name               string `json:"name,omitempty"`
	Transportprotocol  string `json:"transportprotocol,omitempty"`
	Port               string `json:"port,omitempty"`
	Sessiontimeout     int    `json:"sessiontimeout,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Lsnappsprofileportbinding struct {
	Lsnport         string `json:"lsnport,omitempty"`
	Appsprofilename string `json:"appsprofilename,omitempty"`
}

type Lsnclientacl6binding struct {
	Acl6name   string `json:"acl6name,omitempty"`
	Td         uint32 `json:"td,omitempty"`
	Clientname string `json:"clientname,omitempty"`
}

type Lsnclientaclbinding struct {
	Aclname    string `json:"aclname,omitempty"`
	Td         uint32 `json:"td,omitempty"`
	Clientname string `json:"clientname,omitempty"`
}

type Lsnclientnetwork6binding struct {
	Network6   string `json:"network6,omitempty"`
	Network    string `json:"network,omitempty"`
	Netmask    string `json:"netmask,omitempty"`
	Td         int    `json:"td,omitempty"`
	Clientname string `json:"clientname,omitempty"`
}
