package models

// basic configuration structs
type ServerGslbservicegroupBinding struct {
	Appflowlog       string `json:"appflowlog,omitempty"`
	Boundtd          int    `json:"boundtd,omitempty"`
	Cip              string `json:"cip,omitempty"`
	Cipheader        string `json:"cipheader,omitempty"`
	Clttimeout       int    `json:"clttimeout,omitempty"`
	Customserverid   string `json:"customserverid,omitempty"`
	Downstateflush   string `json:"downstateflush,omitempty"`
	DupPort          int    `json:"dup_port,omitempty"`
	DupSvctype       string `json:"dup_svctype,omitempty"`
	Maxbandwidth     int    `json:"maxbandwidth,omitempty"`
	Maxclient        int    `json:"maxclient,omitempty"`
	Maxreq           int    `json:"maxreq,omitempty"`
	Monthreshold     int    `json:"monthreshold,omitempty"`
	Name             string `json:"name,omitempty"`
	Port             int    `json:"port,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Serviceipaddress string `json:"serviceipaddress,omitempty"`
	Serviceipstr     string `json:"serviceipstr,omitempty"`
	Svctype          string `json:"svctype,omitempty"`
	Svrcfgflags      int    `json:"svrcfgflags,omitempty"`
	Svrstate         string `json:"svrstate,omitempty"`
	Svrtimeout       int    `json:"svrtimeout,omitempty"`
}

type ServerServiceBinding struct {
	Name             string `json:"name,omitempty"`
	Port             int    `json:"port,omitempty"`
	Serviceipaddress string `json:"serviceipaddress,omitempty"`
	Serviceipstr     string `json:"serviceipstr,omitempty"`
	Servicename      string `json:"servicename,omitempty"`
	Svctype          string `json:"svctype,omitempty"`
	Svrstate         string `json:"svrstate,omitempty"`
}

type Servicegroupbindings struct {
	Count              float64 `json:"__count,omitempty"`
	Ipaddress          string  `json:"ipaddress,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Port               int     `json:"port,omitempty"`
	Servicegroupname   string  `json:"servicegroupname,omitempty"`
	State              string  `json:"state,omitempty"`
	Svrstate           string  `json:"svrstate,omitempty"`
	Vservername        string  `json:"vservername,omitempty"`
}

type Locationparameter struct {
	Builtin            []string `json:"builtin,omitempty"`
	Context            string   `json:"context,omitempty"`
	Custom             int      `json:"custom,omitempty"`
	Custom6            int      `json:"custom6,omitempty"`
	Databasemode       string   `json:"databasemode,omitempty"`
	Entries            int      `json:"entries,omitempty"`
	Entries6           int      `json:"entries6,omitempty"`
	Errors             int      `json:"errors,omitempty"`
	Errors6            int      `json:"errors6,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Flags              int      `json:"flags,omitempty"`
	Flushing           string   `json:"flushing,omitempty"`
	Format             string   `json:"format,omitempty"`
	Format6            string   `json:"format6,omitempty"`
	Lines              int      `json:"lines,omitempty"`
	Lines6             int      `json:"lines6,omitempty"`
	Loading            string   `json:"loading,omitempty"`
	Locationfile       string   `json:"Locationfile,omitempty"`
	Locationfile6      string   `json:"locationfile6,omitempty"`
	Matchwildcardtoany string   `json:"matchwildcardtoany,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Q1label            string   `json:"q1label,omitempty"`
	Q2label            string   `json:"q2label,omitempty"`
	Q3label            string   `json:"q3label,omitempty"`
	Q4label            string   `json:"q4label,omitempty"`
	Q5label            string   `json:"q5label,omitempty"`
	Q6label            string   `json:"q6label,omitempty"`
	Static             int      `json:"Static,omitempty"`
	Static6            int      `json:"static6,omitempty"`
	Status             int      `json:"status,omitempty"`
	Warnings           int      `json:"warnings,omitempty"`
	Warnings6          int      `json:"warnings6,omitempty"`
}

type Locationdata struct {
}

type ServicegroupServicegroupentitymonbindingsBinding struct {
	Customserverid             string `json:"customserverid,omitempty"`
	Dbsttl                     int    `json:"dbsttl,omitempty"`
	Hashid                     int    `json:"hashid,omitempty"`
	Lastresponse               string `json:"lastresponse,omitempty"`
	MonitorName                string `json:"monitor_name,omitempty"`
	MonitorState               string `json:"monitor_state,omitempty"`
	Monitorcurrentfailedprobes int    `json:"monitorcurrentfailedprobes,omitempty"`
	Monitortotalfailedprobes   int    `json:"monitortotalfailedprobes,omitempty"`
	Monitortotalprobes         int    `json:"monitortotalprobes,omitempty"`
	Nameserver                 string `json:"nameserver,omitempty"`
	Order                      int    `json:"order,omitempty"`
	Passive                    bool   `json:"passive,omitempty"`
	Port                       int    `json:"port,omitempty"`
	Responsetime               int    `json:"responsetime,omitempty"`
	Serverid                   int    `json:"serverid,omitempty"`
	Servicegroupentname2       string `json:"servicegroupentname2,omitempty"`
	Servicegroupname           string `json:"servicegroupname,omitempty"`
	State                      string `json:"state,omitempty"`
	Weight                     int    `json:"weight,omitempty"`
}

type ServiceGroupBinding struct {
	ServicegroupLbmonitorBinding                     []interface{} `json:"servicegroup_lbmonitor_binding,omitempty"`
	ServicegroupServicegroupentitymonbindingsBinding []interface{} `json:"servicegroup_servicegroupentitymonbindings_binding,omitempty"`
	ServicegroupServicegroupmemberBinding            []interface{} `json:"servicegroup_servicegroupmember_binding,omitempty"`
	Servicegroupname                                 string        `json:"servicegroupname,omitempty"`
}

type ServicegroupLbmonitorBinding struct {
	Customserverid   string `json:"customserverid,omitempty"`
	Dbsttl           int    `json:"dbsttl,omitempty"`
	Hashid           int    `json:"hashid,omitempty"`
	MonitorName      string `json:"monitor_name,omitempty"`
	Monstate         string `json:"monstate,omitempty"`
	Monweight        int    `json:"monweight,omitempty"`
	Nameserver       string `json:"nameserver,omitempty"`
	Order            int    `json:"order,omitempty"`
	Passive          bool   `json:"passive,omitempty"`
	Port             int    `json:"port,omitempty"`
	Serverid         int    `json:"serverid,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	State            string `json:"state,omitempty"`
	Weight           int    `json:"weight,omitempty"`
}

type ServiceBinding struct {
	Name                    string        `json:"name,omitempty"`
	ServiceLbmonitorBinding []interface{} `json:"service_lbmonitor_binding,omitempty"`
}

type Radiusnode struct {
	Count              float64 `json:"__count,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Nodeprefix         string  `json:"nodeprefix,omitempty"`
	Radkey             string  `json:"radkey,omitempty"`
}

type Svcbindings struct {
	Count              float64 `json:"__count,omitempty"`
	Ipaddress          string  `json:"ipaddress,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Port               int     `json:"port,omitempty"`
	Servicename        string  `json:"servicename,omitempty"`
	Svrstate           string  `json:"svrstate,omitempty"`
	Vservername        string  `json:"vservername,omitempty"`
}

type ServicegroupServicegroupmemberBinding struct {
	Customserverid            string `json:"customserverid,omitempty"`
	Dbsttl                    int    `json:"dbsttl,omitempty"`
	Delay                     int    `json:"delay,omitempty"`
	Graceful                  string `json:"graceful,omitempty"`
	Hashid                    int    `json:"hashid,omitempty"`
	Ip                        string `json:"ip,omitempty"`
	Nameserver                string `json:"nameserver,omitempty"`
	Order                     int    `json:"order,omitempty"`
	Orderstr                  string `json:"orderstr,omitempty"`
	Port                      int    `json:"port,omitempty"`
	Serverid                  int    `json:"serverid,omitempty"`
	Servername                string `json:"servername,omitempty"`
	Servicegroupname          string `json:"servicegroupname,omitempty"`
	State                     string `json:"state,omitempty"`
	Statechangetimesec        string `json:"statechangetimesec,omitempty"`
	Svcitmpriority            int    `json:"svcitmpriority,omitempty"`
	Svrstate                  string `json:"svrstate,omitempty"`
	Tickssincelaststatechange int    `json:"tickssincelaststatechange,omitempty"`
	Trofsdelay                int    `json:"trofsdelay,omitempty"`
	Trofsreason               string `json:"trofsreason,omitempty"`
	Weight                    int    `json:"weight,omitempty"`
}

type Vserver struct {
	Backupvserver        string `json:"backupvserver,omitempty"`
	Cacheable            string `json:"cacheable,omitempty"`
	Clttimeout           int    `json:"clttimeout,omitempty"`
	Name                 string `json:"name,omitempty"`
	Pushvserver          string `json:"pushvserver,omitempty"`
	Redirecturl          string `json:"redirecturl,omitempty"`
	Somethod             string `json:"somethod,omitempty"`
	Sopersistence        string `json:"sopersistence,omitempty"`
	Sopersistencetimeout int    `json:"sopersistencetimeout,omitempty"`
	Sothreshold          int    `json:"sothreshold,omitempty"`
}

type Nstrace struct {
	Capdroppkt         string        `json:"capdroppkt,omitempty"`
	Capsslkeys         string        `json:"capsslkeys,omitempty"`
	Doruntimecleanup   string        `json:"doruntimecleanup,omitempty"`
	Fileid             string        `json:"fileid,omitempty"`
	Filename           string        `json:"filename,omitempty"`
	Filesize           int           `json:"filesize,omitempty"`
	Filter             string        `json:"filter,omitempty"`
	Inmemorytrace      string        `json:"inmemorytrace,omitempty"`
	Link               string        `json:"link,omitempty"`
	Merge              string        `json:"merge,omitempty"`
	Mode               []string      `json:"mode,omitempty"`
	Nextgenapiresource string        `json:"_nextgenapiresource,omitempty"`
	Nf                 int           `json:"nf,omitempty"`
	Nodeid             int           `json:"nodeid,omitempty"`
	Nodes              []interface{} `json:"nodes,omitempty"`
	Pernic             string        `json:"pernic,omitempty"`
	Scope              string        `json:"scope,omitempty"`
	Size               int           `json:"size,omitempty"`
	Skiplocalssh       string        `json:"skiplocalssh,omitempty"`
	Skiprpc            string        `json:"skiprpc,omitempty"`
	State              string        `json:"state,omitempty"`
	Time               int           `json:"time,omitempty"`
	Tracebuffers       int           `json:"tracebuffers,omitempty"`
	Traceformat        string        `json:"traceformat,omitempty"`
	Tracelocation      string        `json:"tracelocation,omitempty"`
}

type Locationfile struct {
	Curlocfilestatus   string `json:"curlocfilestatus,omitempty"`
	Format             string `json:"format,omitempty"`
	Locationfile       string `json:"Locationfile,omitempty"`
	Locfilestatusstr   string `json:"locfilestatusstr,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Prevlocationfile   string `json:"prevlocationfile,omitempty"`
	Prevlocfileformat  string `json:"prevlocfileformat,omitempty"`
	Prevlocfilestatus  string `json:"prevlocfilestatus,omitempty"`
	Src                string `json:"src,omitempty"`
}

type ServerGslbserviceBinding struct {
	Name             string `json:"name,omitempty"`
	Port             int    `json:"port,omitempty"`
	Serviceipaddress string `json:"serviceipaddress,omitempty"`
	Serviceipstr     string `json:"serviceipstr,omitempty"`
	Servicename      string `json:"servicename,omitempty"`
	Svctype          string `json:"svctype,omitempty"`
	Svrstate         string `json:"svrstate,omitempty"`
}

type Extendedmemoryparam struct {
	Maxmemlimit        int    `json:"maxmemlimit,omitempty"`
	Memlimit           int    `json:"memlimit,omitempty"`
	Memlimitactive     int    `json:"memlimitactive,omitempty"`
	Minrequiredmemory  int    `json:"minrequiredmemory,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Service struct {
	Accessdown                   string   `json:"accessdown,omitempty"`
	All                          bool     `json:"all,omitempty"`
	Appflowlog                   string   `json:"appflowlog,omitempty"`
	Builtin                      []string `json:"builtin,omitempty"`
	Cacheable                    string   `json:"cacheable,omitempty"`
	Cachetype                    string   `json:"cachetype,omitempty"`
	Cip                          string   `json:"cip,omitempty"`
	Cipheader                    string   `json:"cipheader,omitempty"`
	Cka                          string   `json:"cka,omitempty"`
	Cleartextport                int      `json:"cleartextport,omitempty"`
	Clmonowner                   int      `json:"clmonowner,omitempty"`
	Clmonview                    int      `json:"clmonview,omitempty"`
	Clttimeout                   int      `json:"clttimeout,omitempty"`
	Cmp                          string   `json:"cmp,omitempty"`
	Comment                      string   `json:"comment,omitempty"`
	Contentinspectionprofilename string   `json:"contentinspectionprofilename,omitempty"`
	Count                        float64  `json:"__count,omitempty"`
	Customserverid               string   `json:"customserverid,omitempty"`
	Delay                        int      `json:"delay,omitempty"`
	Dnsprofilename               string   `json:"dnsprofilename,omitempty"`
	Downstateflush               string   `json:"downstateflush,omitempty"`
	DupState                     string   `json:"dup_state,omitempty"`
	Feature                      string   `json:"feature,omitempty"`
	Graceful                     string   `json:"graceful,omitempty"`
	Gslb                         string   `json:"gslb,omitempty"`
	Hashid                       int      `json:"hashid,omitempty"`
	Healthmonitor                string   `json:"healthmonitor,omitempty"`
	Httpprofilename              string   `json:"httpprofilename,omitempty"`
	Internal                     bool     `json:"Internal,omitempty"`
	Ip                           string   `json:"ip,omitempty"`
	Ipaddress                    string   `json:"ipaddress,omitempty"`
	Lastresponse                 string   `json:"lastresponse,omitempty"`
	Maxbandwidth                 int      `json:"maxbandwidth,omitempty"`
	Maxclient                    int      `json:"maxclient,omitempty"`
	Maxreq                       int      `json:"maxreq,omitempty"`
	Monconnectionclose           string   `json:"monconnectionclose,omitempty"`
	MonitorNameSvc               string   `json:"monitor_name_svc,omitempty"`
	MonitorState                 string   `json:"monitor_state,omitempty"`
	Monstatcode                  int      `json:"monstatcode,omitempty"`
	Monstatparam1                int      `json:"monstatparam1,omitempty"`
	Monstatparam2                int      `json:"monstatparam2,omitempty"`
	Monstatparam3                int      `json:"monstatparam3,omitempty"`
	Monthreshold                 int      `json:"monthreshold,omitempty"`
	Monuserstatusmesg            string   `json:"monuserstatusmesg,omitempty"`
	Name                         string   `json:"name,omitempty"`
	Netprofile                   string   `json:"netprofile,omitempty"`
	Newname                      string   `json:"newname,omitempty"`
	Nextgenapiresource           string   `json:"_nextgenapiresource,omitempty"`
	Nodefaultbindings            string   `json:"nodefaultbindings,omitempty"`
	Numofconnections             int      `json:"numofconnections,omitempty"`
	Oracleserverversion          string   `json:"oracleserverversion,omitempty"`
	Pathmonitor                  string   `json:"pathmonitor,omitempty"`
	Pathmonitorindv              string   `json:"pathmonitorindv,omitempty"`
	Policyname                   string   `json:"policyname,omitempty"`
	Port                         int      `json:"port,omitempty"`
	Processlocal                 string   `json:"processlocal,omitempty"`
	Publicip                     string   `json:"publicip,omitempty"`
	Publicport                   int      `json:"publicport,omitempty"`
	Quicprofilename              string   `json:"quicprofilename,omitempty"`
	Responsetime                 int      `json:"responsetime,omitempty"`
	Rtspsessionidremap           string   `json:"rtspsessionidremap,omitempty"`
	Serverid                     int      `json:"serverid,omitempty"`
	Servername                   string   `json:"servername,omitempty"`
	Serviceconftype              bool     `json:"serviceconftype,omitempty"`
	Serviceconftype2             string   `json:"serviceconftype2,omitempty"`
	Serviceipstr                 string   `json:"serviceipstr,omitempty"`
	Servicetype                  string   `json:"servicetype,omitempty"`
	Sp                           string   `json:"sp,omitempty"`
	State                        string   `json:"state,omitempty"`
	Statechangetimemsec          int      `json:"statechangetimemsec,omitempty"`
	Statechangetimesec           string   `json:"statechangetimesec,omitempty"`
	Stateupdatereason            int      `json:"stateupdatereason,omitempty"`
	Svrstate                     string   `json:"svrstate,omitempty"`
	Svrtimeout                   int      `json:"svrtimeout,omitempty"`
	Tcpb                         string   `json:"tcpb,omitempty"`
	Tcpprofilename               string   `json:"tcpprofilename,omitempty"`
	Td                           int      `json:"td,omitempty"`
	Tickssincelaststatechange    int      `json:"tickssincelaststatechange,omitempty"`
	Useproxyport                 string   `json:"useproxyport,omitempty"`
	Usip                         string   `json:"usip,omitempty"`
	Value                        string   `json:"value,omitempty"`
	Weight                       int      `json:"weight,omitempty"`
}

type ServerServicegroupBinding struct {
	Appflowlog           string `json:"appflowlog,omitempty"`
	Boundtd              int    `json:"boundtd,omitempty"`
	Cacheable            string `json:"cacheable,omitempty"`
	Cip                  string `json:"cip,omitempty"`
	Cipheader            string `json:"cipheader,omitempty"`
	Cka                  string `json:"cka,omitempty"`
	Clttimeout           int    `json:"clttimeout,omitempty"`
	Cmp                  string `json:"cmp,omitempty"`
	Customserverid       string `json:"customserverid,omitempty"`
	Downstateflush       string `json:"downstateflush,omitempty"`
	DupPort              int    `json:"dup_port,omitempty"`
	DupSvctype           string `json:"dup_svctype,omitempty"`
	Maxbandwidth         int    `json:"maxbandwidth,omitempty"`
	Maxclient            int    `json:"maxclient,omitempty"`
	Maxreq               int    `json:"maxreq,omitempty"`
	Monthreshold         int    `json:"monthreshold,omitempty"`
	Name                 string `json:"name,omitempty"`
	Port                 int    `json:"port,omitempty"`
	Servicegroupentname2 string `json:"servicegroupentname2,omitempty"`
	Servicegroupname     string `json:"servicegroupname,omitempty"`
	Serviceipaddress     string `json:"serviceipaddress,omitempty"`
	Serviceipstr         string `json:"serviceipstr,omitempty"`
	Sp                   string `json:"sp,omitempty"`
	Svcitmactsvcs        int    `json:"svcitmactsvcs,omitempty"`
	Svcitmboundsvcs      int    `json:"svcitmboundsvcs,omitempty"`
	Svcitmpriority       int    `json:"svcitmpriority,omitempty"`
	Svctype              string `json:"svctype,omitempty"`
	Svrcfgflags          int    `json:"svrcfgflags,omitempty"`
	Svrstate             string `json:"svrstate,omitempty"`
	Svrtimeout           int    `json:"svrtimeout,omitempty"`
	Tcpb                 string `json:"tcpb,omitempty"`
	Usip                 string `json:"usip,omitempty"`
	Weight               int    `json:"weight,omitempty"`
}

type Dbsmonitors struct {
}

type Locationfile6 struct {
	Count              float64 `json:"__count,omitempty"`
	Curlocfilestatus   string  `json:"curlocfilestatus,omitempty"`
	Format             string  `json:"format,omitempty"`
	Locationfile       string  `json:"Locationfile,omitempty"`
	Locfilestatusstr   string  `json:"locfilestatusstr,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Prevlocationfile   string  `json:"prevlocationfile,omitempty"`
	Prevlocfileformat  string  `json:"prevlocfileformat,omitempty"`
	Prevlocfilestatus  string  `json:"prevlocfilestatus,omitempty"`
	Src                string  `json:"src,omitempty"`
}

type Servicegroup struct {
	Appflowlog                 string  `json:"appflowlog,omitempty"`
	Autodelayedtrofs           string  `json:"autodelayedtrofs,omitempty"`
	Autodisabledelay           int     `json:"autodisabledelay,omitempty"`
	Autodisablegraceful        string  `json:"autodisablegraceful,omitempty"`
	Autoscale                  string  `json:"autoscale,omitempty"`
	Bootstrap                  string  `json:"bootstrap,omitempty"`
	Cacheable                  string  `json:"cacheable,omitempty"`
	Cachetype                  string  `json:"cachetype,omitempty"`
	Cip                        string  `json:"cip,omitempty"`
	Cipheader                  string  `json:"cipheader,omitempty"`
	Cka                        string  `json:"cka,omitempty"`
	Clmonowner                 int     `json:"clmonowner,omitempty"`
	Clmonview                  int     `json:"clmonview,omitempty"`
	Clttimeout                 int     `json:"clttimeout,omitempty"`
	Cmp                        string  `json:"cmp,omitempty"`
	Comment                    string  `json:"comment,omitempty"`
	Count                      float64 `json:"__count,omitempty"`
	Customserverid             string  `json:"customserverid,omitempty"`
	Dbsttl                     int     `json:"dbsttl,omitempty"`
	Delay                      int     `json:"delay,omitempty"`
	Downstateflush             string  `json:"downstateflush,omitempty"`
	DupWeight                  int     `json:"dup_weight,omitempty"`
	Graceful                   string  `json:"graceful,omitempty"`
	Groupcount                 int     `json:"groupcount,omitempty"`
	Hashid                     int     `json:"hashid,omitempty"`
	Healthmonitor              string  `json:"healthmonitor,omitempty"`
	Httpprofilename            string  `json:"httpprofilename,omitempty"`
	Includemembers             bool    `json:"includemembers,omitempty"`
	Ip                         string  `json:"ip,omitempty"`
	Maxbandwidth               int     `json:"maxbandwidth,omitempty"`
	Maxclient                  int     `json:"maxclient,omitempty"`
	Maxreq                     int     `json:"maxreq,omitempty"`
	Memberport                 int     `json:"memberport,omitempty"`
	Monconnectionclose         string  `json:"monconnectionclose,omitempty"`
	MonitorNameSvc             string  `json:"monitor_name_svc,omitempty"`
	Monstatcode                int     `json:"monstatcode,omitempty"`
	Monstatparam1              int     `json:"monstatparam1,omitempty"`
	Monstatparam2              int     `json:"monstatparam2,omitempty"`
	Monstatparam3              int     `json:"monstatparam3,omitempty"`
	Monthreshold               int     `json:"monthreshold,omitempty"`
	Monuserstatusmesg          string  `json:"monuserstatusmesg,omitempty"`
	Nameserver                 string  `json:"nameserver,omitempty"`
	Netprofile                 string  `json:"netprofile,omitempty"`
	Newname                    string  `json:"newname,omitempty"`
	Nextgenapiresource         string  `json:"_nextgenapiresource,omitempty"`
	Nodefaultbindings          string  `json:"nodefaultbindings,omitempty"`
	Numofconnections           int     `json:"numofconnections,omitempty"`
	Order                      int     `json:"order,omitempty"`
	Pathmonitor                string  `json:"pathmonitor,omitempty"`
	Pathmonitorindv            string  `json:"pathmonitorindv,omitempty"`
	Port                       int     `json:"port,omitempty"`
	Quicprofilename            string  `json:"quicprofilename,omitempty"`
	Rtspsessionidremap         string  `json:"rtspsessionidremap,omitempty"`
	Serverid                   int     `json:"serverid,omitempty"`
	Servername                 string  `json:"servername,omitempty"`
	Serviceconftype            bool    `json:"serviceconftype,omitempty"`
	Servicegroupeffectivestate string  `json:"servicegroupeffectivestate,omitempty"`
	Servicegroupname           string  `json:"servicegroupname,omitempty"`
	Serviceipstr               string  `json:"serviceipstr,omitempty"`
	Servicetype                string  `json:"servicetype,omitempty"`
	Sp                         string  `json:"sp,omitempty"`
	State                      string  `json:"state,omitempty"`
	Statechangetimemsec        int     `json:"statechangetimemsec,omitempty"`
	Stateupdatereason          int     `json:"stateupdatereason,omitempty"`
	Svcitmactsvcs              int     `json:"svcitmactsvcs,omitempty"`
	Svcitmboundsvcs            int     `json:"svcitmboundsvcs,omitempty"`
	Svrstate                   string  `json:"svrstate,omitempty"`
	Svrtimeout                 int     `json:"svrtimeout,omitempty"`
	Tcpb                       string  `json:"tcpb,omitempty"`
	Tcpprofilename             string  `json:"tcpprofilename,omitempty"`
	Td                         int     `json:"td,omitempty"`
	Topicname                  string  `json:"topicname,omitempty"`
	Useproxyport               string  `json:"useproxyport,omitempty"`
	Usip                       string  `json:"usip,omitempty"`
	Value                      string  `json:"value,omitempty"`
	Weight                     int     `json:"weight,omitempty"`
}

type ServerBinding struct {
	Name                          string        `json:"name,omitempty"`
	ServerGslbserviceBinding      []interface{} `json:"server_gslbservice_binding,omitempty"`
	ServerGslbservicegroupBinding []interface{} `json:"server_gslbservicegroup_binding,omitempty"`
	ServerServiceBinding          []interface{} `json:"server_service_binding,omitempty"`
	ServerServicegroupBinding     []interface{} `json:"server_servicegroup_binding,omitempty"`
}

type ServiceLbmonitorBinding struct {
	DupState                   string `json:"dup_state,omitempty"`
	DupWeight                  int    `json:"dup_weight,omitempty"`
	Failedprobes               int    `json:"failedprobes,omitempty"`
	Lastresponse               string `json:"lastresponse,omitempty"`
	MonitorName                string `json:"monitor_name,omitempty"`
	MonitorState               string `json:"monitor_state,omitempty"`
	Monitorcurrentfailedprobes int    `json:"monitorcurrentfailedprobes,omitempty"`
	Monitortotalfailedprobes   int    `json:"monitortotalfailedprobes,omitempty"`
	Monitortotalprobes         int    `json:"monitortotalprobes,omitempty"`
	Monstatcode                int    `json:"monstatcode,omitempty"`
	Monstate                   string `json:"monstate,omitempty"`
	Monstatparam1              int    `json:"monstatparam1,omitempty"`
	Monstatparam2              int    `json:"monstatparam2,omitempty"`
	Monstatparam3              int    `json:"monstatparam3,omitempty"`
	Name                       string `json:"name,omitempty"`
	Passive                    bool   `json:"passive,omitempty"`
	Responsetime               int    `json:"responsetime,omitempty"`
	Totalfailedprobes          int    `json:"totalfailedprobes,omitempty"`
	Totalprobes                int    `json:"totalprobes,omitempty"`
	Weight                     int    `json:"weight,omitempty"`
}

type ServicegroupServicegroupmemberlistBinding struct {
	Failedmembers    []interface{} `json:"failedmembers,omitempty"`
	Members          []interface{} `json:"members,omitempty"`
	Servicegroupname string        `json:"servicegroupname,omitempty"`
}

type Server struct {
	Autoscale                 string  `json:"autoscale,omitempty"`
	Cacheable                 string  `json:"cacheable,omitempty"`
	Cka                       string  `json:"cka,omitempty"`
	Cmp                       string  `json:"cmp,omitempty"`
	Comment                   string  `json:"comment,omitempty"`
	Count                     float64 `json:"__count,omitempty"`
	Delay                     int     `json:"delay,omitempty"`
	Domain                    string  `json:"domain,omitempty"`
	Domainresolvenow          bool    `json:"domainresolvenow,omitempty"`
	Domainresolveretry        int     `json:"domainresolveretry,omitempty"`
	Graceful                  string  `json:"graceful,omitempty"`
	Internal                  bool    `json:"Internal,omitempty"`
	Ipaddress                 string  `json:"ipaddress,omitempty"`
	Ipv6address               string  `json:"ipv6address,omitempty"`
	Name                      string  `json:"name,omitempty"`
	Newname                   string  `json:"newname,omitempty"`
	Nextgenapiresource        string  `json:"_nextgenapiresource,omitempty"`
	Querytype                 string  `json:"querytype,omitempty"`
	Sp                        string  `json:"sp,omitempty"`
	State                     string  `json:"state,omitempty"`
	Statechangetimesec        string  `json:"statechangetimesec,omitempty"`
	Tcpb                      string  `json:"tcpb,omitempty"`
	Td                        int     `json:"td,omitempty"`
	Tickssincelaststatechange int     `json:"tickssincelaststatechange,omitempty"`
	Translationip             string  `json:"translationip,omitempty"`
	Translationmask           string  `json:"translationmask,omitempty"`
	Usip                      string  `json:"usip,omitempty"`
}

type Location struct {
	Count              float64 `json:"__count,omitempty"`
	Ipfrom             string  `json:"ipfrom,omitempty"`
	Ipto               string  `json:"ipto,omitempty"`
	Latitude           int     `json:"latitude,omitempty"`
	Longitude          int     `json:"longitude,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Preferredlocation  string  `json:"preferredlocation,omitempty"`
	Q1label            string  `json:"q1label,omitempty"`
	Q2label            string  `json:"q2label,omitempty"`
	Q3label            string  `json:"q3label,omitempty"`
	Q4label            string  `json:"q4label,omitempty"`
	Q5label            string  `json:"q5label,omitempty"`
	Q6label            string  `json:"q6label,omitempty"`
}
