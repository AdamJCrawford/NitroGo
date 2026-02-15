package models

// lsn configuration structs
type LsngroupLsnlogprofileBinding struct {
	Groupname      string `json:"groupname,omitempty"`
	Logprofilename string `json:"logprofilename,omitempty"`
}

type LsnclientNsaclBinding struct {
	Aclname    string `json:"aclname,omitempty"`
	Clientname string `json:"clientname,omitempty"`
	Td         int    `json:"td,omitempty"`
}

type Lsnip6profile struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Natprefix          string  `json:"natprefix,omitempty"`
	Network6           string  `json:"network6,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	TypeField          string  `json:"type,omitempty"`
}

type LsngroupLsnhttphdrlogprofileBinding struct {
	Groupname             string `json:"groupname,omitempty"`
	Httphdrlogprofilename string `json:"httphdrlogprofilename,omitempty"`
}

type LsngroupLsntransportprofileBinding struct {
	Groupname            string `json:"groupname,omitempty"`
	Transportprofilename string `json:"transportprofilename,omitempty"`
}

type Lsnappsattributes struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Port               string  `json:"port,omitempty"`
	Sessiontimeout     int     `json:"sessiontimeout,omitempty"`
	Transportprotocol  string  `json:"transportprotocol,omitempty"`
}

type LsngroupLsnrtspalgprofileBinding struct {
	Groupname          string `json:"groupname,omitempty"`
	Rtspalgprofilename string `json:"rtspalgprofilename,omitempty"`
}

type Lsnclient struct {
	Clientname         string  `json:"clientname,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
}

type LsnclientNsacl6Binding struct {
	Acl6name   string `json:"acl6name,omitempty"`
	Clientname string `json:"clientname,omitempty"`
	Td         int    `json:"td,omitempty"`
}

type Lsnparameter struct {
	Maxmemlimit          int    `json:"maxmemlimit,omitempty"`
	Memlimit             int    `json:"memlimit,omitempty"`
	Memlimitactive       int    `json:"memlimitactive,omitempty"`
	Nextgenapiresource   string `json:"_nextgenapiresource,omitempty"`
	Sessionsync          string `json:"sessionsync,omitempty"`
	Subscrsessionremoval string `json:"subscrsessionremoval,omitempty"`
}

type LsnclientNetworkBinding struct {
	Clientname string `json:"clientname,omitempty"`
	Netmask    string `json:"netmask,omitempty"`
	Network    string `json:"network,omitempty"`
	Td         int    `json:"td,omitempty"`
}

type LsnclientNetwork6Binding struct {
	Clientname string `json:"clientname,omitempty"`
	Netmask    string `json:"netmask,omitempty"`
	Network    string `json:"network,omitempty"`
	Network6   string `json:"network6,omitempty"`
	Td         int    `json:"td,omitempty"`
}

type LsnappsprofilePortBinding struct {
	Appsprofilename string `json:"appsprofilename,omitempty"`
	Lsnport         string `json:"lsnport,omitempty"`
}

type Lsnsession struct {
	Clientname         string  `json:"clientname,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Destip             string  `json:"destip,omitempty"`
	Destport           int     `json:"destport,omitempty"`
	Dsttd              int     `json:"dsttd,omitempty"`
	Ipv6address        string  `json:"ipv6address,omitempty"`
	Natip              string  `json:"natip,omitempty"`
	Natport            int     `json:"natport,omitempty"`
	Natport2           int     `json:"natport2,omitempty"`
	Natprefix          string  `json:"natprefix,omitempty"`
	Nattype            string  `json:"nattype,omitempty"`
	Netmask            string  `json:"netmask,omitempty"`
	Network            string  `json:"network,omitempty"`
	Network6           string  `json:"network6,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Nodeid             int     `json:"nodeid,omitempty"`
	Sessionestdir      string  `json:"sessionestdir,omitempty"`
	Srctd              int     `json:"srctd,omitempty"`
	Subscrip           string  `json:"subscrip,omitempty"`
	Subscrport         int     `json:"subscrport,omitempty"`
	Td                 int     `json:"td,omitempty"`
	Transportprotocol  string  `json:"transportprotocol,omitempty"`
}

type Lsnsipalgprofile struct {
	Count                  float64 `json:"__count,omitempty"`
	Datasessionidletimeout int     `json:"datasessionidletimeout,omitempty"`
	Nextgenapiresource     string  `json:"_nextgenapiresource,omitempty"`
	Opencontactpinhole     string  `json:"opencontactpinhole,omitempty"`
	Openrecordroutepinhole string  `json:"openrecordroutepinhole,omitempty"`
	Openregisterpinhole    string  `json:"openregisterpinhole,omitempty"`
	Openroutepinhole       string  `json:"openroutepinhole,omitempty"`
	Openviapinhole         string  `json:"openviapinhole,omitempty"`
	Registrationtimeout    int     `json:"registrationtimeout,omitempty"`
	Rport                  string  `json:"rport,omitempty"`
	Sipalgprofilename      string  `json:"sipalgprofilename,omitempty"`
	Sipdstportrange        string  `json:"sipdstportrange,omitempty"`
	Sipsessiontimeout      int     `json:"sipsessiontimeout,omitempty"`
	Sipsrcportrange        string  `json:"sipsrcportrange,omitempty"`
	Siptransportprotocol   string  `json:"siptransportprotocol,omitempty"`
}

type LsnrtspalgsessionBinding struct {
	LsnrtspalgsessionDatachannelBinding []interface{} `json:"lsnrtspalgsession_datachannel_binding,omitempty"`
	Sessionid                           string        `json:"sessionid,omitempty"`
}

type Lsnlogprofile struct {
	Analyticsprofile   string  `json:"analyticsprofile,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Logcompact         string  `json:"logcompact,omitempty"`
	Logipfix           string  `json:"logipfix,omitempty"`
	Logprofilename     string  `json:"logprofilename,omitempty"`
	Logsessdeletion    string  `json:"logsessdeletion,omitempty"`
	Logsubscrinfo      string  `json:"logsubscrinfo,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
}

type LsngroupIpsecalgprofileBinding struct {
	Groupname       string `json:"groupname,omitempty"`
	Ipsecalgprofile string `json:"ipsecalgprofile,omitempty"`
}

type Lsngroup struct {
	Allocpolicy        string  `json:"allocpolicy,omitempty"`
	Clientname         string  `json:"clientname,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Ftp                string  `json:"ftp,omitempty"`
	Ftpcm              string  `json:"ftpcm,omitempty"`
	Groupid            int     `json:"groupid,omitempty"`
	Groupname          string  `json:"groupname,omitempty"`
	Ip6profile         string  `json:"ip6profile,omitempty"`
	Logging            string  `json:"logging,omitempty"`
	Nattype            string  `json:"nattype,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Portblocksize      int     `json:"portblocksize,omitempty"`
	Pptp               string  `json:"pptp,omitempty"`
	Rtspalg            string  `json:"rtspalg,omitempty"`
	Sessionlogging     string  `json:"sessionlogging,omitempty"`
	Sessionsync        string  `json:"sessionsync,omitempty"`
	Sipalg             string  `json:"sipalg,omitempty"`
	Snmptraplimit      int     `json:"snmptraplimit,omitempty"`
}

type Lsnstatic struct {
	Count              float64 `json:"__count,omitempty"`
	Destip             string  `json:"destip,omitempty"`
	Dsttd              int     `json:"dsttd,omitempty"`
	Name               string  `json:"name,omitempty"`
	Natip              string  `json:"natip,omitempty"`
	Natport            int     `json:"natport,omitempty"`
	Nattype            string  `json:"nattype,omitempty"`
	Network6           string  `json:"network6,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Status             string  `json:"status,omitempty"`
	Subscrip           string  `json:"subscrip,omitempty"`
	Subscrport         int     `json:"subscrport,omitempty"`
	Td                 int     `json:"td,omitempty"`
	Transportprotocol  string  `json:"transportprotocol,omitempty"`
}

type Lsntransportprofile struct {
	Count                float64 `json:"__count,omitempty"`
	Finrsttimeout        int     `json:"finrsttimeout,omitempty"`
	Groupsessionlimit    int     `json:"groupsessionlimit,omitempty"`
	Nextgenapiresource   string  `json:"_nextgenapiresource,omitempty"`
	Portpreserveparity   string  `json:"portpreserveparity,omitempty"`
	Portpreserverange    string  `json:"portpreserverange,omitempty"`
	Portquota            int     `json:"portquota,omitempty"`
	Sessionquota         int     `json:"sessionquota,omitempty"`
	Sessiontimeout       int     `json:"sessiontimeout,omitempty"`
	Stuntimeout          int     `json:"stuntimeout,omitempty"`
	Syncheck             string  `json:"syncheck,omitempty"`
	Synidletimeout       int     `json:"synidletimeout,omitempty"`
	Transportprofilename string  `json:"transportprofilename,omitempty"`
	Transportprotocol    string  `json:"transportprotocol,omitempty"`
}

type LsnsipalgcallDatachannelBinding struct {
	Callid          string `json:"callid,omitempty"`
	Channelflags    int    `json:"channelflags,omitempty"`
	Channelip       string `json:"channelip,omitempty"`
	Channelnatip    string `json:"channelnatip,omitempty"`
	Channelnatport  int    `json:"channelnatport,omitempty"`
	Channelport     int    `json:"channelport,omitempty"`
	Channelprotocol string `json:"channelprotocol,omitempty"`
	Channeltimeout  int    `json:"channeltimeout,omitempty"`
}

type LsnappsprofileBinding struct {
	Appsprofilename                        string        `json:"appsprofilename,omitempty"`
	LsnappsprofileLsnappsattributesBinding []interface{} `json:"lsnappsprofile_lsnappsattributes_binding,omitempty"`
	LsnappsprofilePortBinding              []interface{} `json:"lsnappsprofile_port_binding,omitempty"`
}

type Lsnpool struct {
	Count               float64 `json:"__count,omitempty"`
	Maxportrealloctmq   int     `json:"maxportrealloctmq,omitempty"`
	Nattype             string  `json:"nattype,omitempty"`
	Nextgenapiresource  string  `json:"_nextgenapiresource,omitempty"`
	Poolname            string  `json:"poolname,omitempty"`
	Portblockallocation string  `json:"portblockallocation,omitempty"`
	Portrealloctimeout  int     `json:"portrealloctimeout,omitempty"`
}

type LsnappsprofileLsnappsattributesBinding struct {
	Appsattributesname string `json:"appsattributesname,omitempty"`
	Appsprofilename    string `json:"appsprofilename,omitempty"`
}

type LsngroupBinding struct {
	Groupname                           string        `json:"groupname,omitempty"`
	LsngroupIpsecalgprofileBinding      []interface{} `json:"lsngroup_ipsecalgprofile_binding,omitempty"`
	LsngroupLsnappsprofileBinding       []interface{} `json:"lsngroup_lsnappsprofile_binding,omitempty"`
	LsngroupLsnhttphdrlogprofileBinding []interface{} `json:"lsngroup_lsnhttphdrlogprofile_binding,omitempty"`
	LsngroupLsnlogprofileBinding        []interface{} `json:"lsngroup_lsnlogprofile_binding,omitempty"`
	LsngroupLsnpoolBinding              []interface{} `json:"lsngroup_lsnpool_binding,omitempty"`
	LsngroupLsnrtspalgprofileBinding    []interface{} `json:"lsngroup_lsnrtspalgprofile_binding,omitempty"`
	LsngroupLsnsipalgprofileBinding     []interface{} `json:"lsngroup_lsnsipalgprofile_binding,omitempty"`
	LsngroupLsntransportprofileBinding  []interface{} `json:"lsngroup_lsntransportprofile_binding,omitempty"`
	LsngroupPcpserverBinding            []interface{} `json:"lsngroup_pcpserver_binding,omitempty"`
}

type LsnpoolBinding struct {
	LsnpoolLsnipBinding []interface{} `json:"lsnpool_lsnip_binding,omitempty"`
	Poolname            string        `json:"poolname,omitempty"`
}

type Lsndeterministicnat struct {
	Clientname         string  `json:"clientname,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Firstport          int     `json:"firstport,omitempty"`
	Lastport           int     `json:"lastport,omitempty"`
	Natip              string  `json:"natip,omitempty"`
	Natip2             string  `json:"natip2,omitempty"`
	Natprefix          string  `json:"natprefix,omitempty"`
	Nattype            string  `json:"nattype,omitempty"`
	Network6           string  `json:"network6,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Srctd              int     `json:"srctd,omitempty"`
	Subscrip           string  `json:"subscrip,omitempty"`
	Subscrip2          string  `json:"subscrip2,omitempty"`
	Td                 int     `json:"td,omitempty"`
}

type LsngroupLsnappsprofileBinding struct {
	Appsprofilename string `json:"appsprofilename,omitempty"`
	Groupname       string `json:"groupname,omitempty"`
}

type LsnsipalgcallBinding struct {
	Callid                             string        `json:"callid,omitempty"`
	LsnsipalgcallControlchannelBinding []interface{} `json:"lsnsipalgcall_controlchannel_binding,omitempty"`
	LsnsipalgcallDatachannelBinding    []interface{} `json:"lsnsipalgcall_datachannel_binding,omitempty"`
}

type Lsnappsprofile struct {
	Appsprofilename    string  `json:"appsprofilename,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Filtering          string  `json:"filtering,omitempty"`
	Ippooling          string  `json:"ippooling,omitempty"`
	L2info             string  `json:"l2info,omitempty"`
	Mapping            string  `json:"mapping,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Tcpproxy           string  `json:"tcpproxy,omitempty"`
	Td                 int     `json:"td,omitempty"`
	Transportprotocol  string  `json:"transportprotocol,omitempty"`
}

type Lsnhttphdrlogprofile struct {
	Count                 float64 `json:"__count,omitempty"`
	Httphdrlogprofilename string  `json:"httphdrlogprofilename,omitempty"`
	Loghost               string  `json:"loghost,omitempty"`
	Logmethod             string  `json:"logmethod,omitempty"`
	Logurl                string  `json:"logurl,omitempty"`
	Logversion            string  `json:"logversion,omitempty"`
	Nextgenapiresource    string  `json:"_nextgenapiresource,omitempty"`
}

type LsnsipalgcallControlchannelBinding struct {
	Callid          string `json:"callid,omitempty"`
	Channelflags    int    `json:"channelflags,omitempty"`
	Channelip       string `json:"channelip,omitempty"`
	Channelnatip    string `json:"channelnatip,omitempty"`
	Channelnatport  int    `json:"channelnatport,omitempty"`
	Channelport     int    `json:"channelport,omitempty"`
	Channelprotocol string `json:"channelprotocol,omitempty"`
	Channeltimeout  int    `json:"channeltimeout,omitempty"`
}

type LsngroupLsnpoolBinding struct {
	Groupname string `json:"groupname,omitempty"`
	Poolname  string `json:"poolname,omitempty"`
}

type Lsnsipalgcall struct {
	Callflags          int     `json:"callflags,omitempty"`
	Callid             string  `json:"callid,omitempty"`
	Callrefcount       int     `json:"callrefcount,omitempty"`
	Calltimer          int     `json:"calltimer,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Nodeid             int     `json:"nodeid,omitempty"`
	Xlatip             string  `json:"xlatip,omitempty"`
}

type LsnpoolLsnipBinding struct {
	Lsnip     string `json:"lsnip,omitempty"`
	Ownernode int    `json:"ownernode,omitempty"`
	Poolname  string `json:"poolname,omitempty"`
}

type Lsnrtspalgsession struct {
	Callflags          int     `json:"callflags,omitempty"`
	Callrefcount       int     `json:"callrefcount,omitempty"`
	Calltimer          int     `json:"calltimer,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Nodeid             int     `json:"nodeid,omitempty"`
	Sessionid          string  `json:"sessionid,omitempty"`
	Xlatip             string  `json:"xlatip,omitempty"`
}

type Lsnrtspalgprofile struct {
	Count                 float64 `json:"__count,omitempty"`
	Nextgenapiresource    string  `json:"_nextgenapiresource,omitempty"`
	Rtspalgprofilename    string  `json:"rtspalgprofilename,omitempty"`
	Rtspidletimeout       int     `json:"rtspidletimeout,omitempty"`
	Rtspportrange         string  `json:"rtspportrange,omitempty"`
	Rtsptransportprotocol string  `json:"rtsptransportprotocol,omitempty"`
}

type LsngroupPcpserverBinding struct {
	Groupname string `json:"groupname,omitempty"`
	Pcpserver string `json:"pcpserver,omitempty"`
}

type LsnclientBinding struct {
	Clientname               string        `json:"clientname,omitempty"`
	LsnclientNetwork6Binding []interface{} `json:"lsnclient_network6_binding,omitempty"`
	LsnclientNetworkBinding  []interface{} `json:"lsnclient_network_binding,omitempty"`
	LsnclientNsacl6Binding   []interface{} `json:"lsnclient_nsacl6_binding,omitempty"`
	LsnclientNsaclBinding    []interface{} `json:"lsnclient_nsacl_binding,omitempty"`
}

type LsnrtspalgsessionDatachannelBinding struct {
	Channelflags    int    `json:"channelflags,omitempty"`
	Channelip       string `json:"channelip,omitempty"`
	Channelnatip    string `json:"channelnatip,omitempty"`
	Channelnatport  int    `json:"channelnatport,omitempty"`
	Channelport     int    `json:"channelport,omitempty"`
	Channelprotocol string `json:"channelprotocol,omitempty"`
	Channeltimeout  int    `json:"channeltimeout,omitempty"`
	Sessionid       string `json:"sessionid,omitempty"`
}

type LsngroupLsnsipalgprofileBinding struct {
	Groupname         string `json:"groupname,omitempty"`
	Sipalgprofilename string `json:"sipalgprofilename,omitempty"`
}
