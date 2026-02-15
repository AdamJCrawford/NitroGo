package models

// system configuration structs
type SystemglobalAuthenticationldappolicyBinding struct {
	Builtin                []string `json:"builtin,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	Globalbindtype         string   `json:"globalbindtype,omitempty"`
	Gotopriorityexpression string   `json:"gotopriorityexpression,omitempty"`
	Nextfactor             string   `json:"nextfactor,omitempty"`
	Policyname             string   `json:"policyname,omitempty"`
	Priority               int      `json:"priority,omitempty"`
}

type SystemuserSystemcmdpolicyBinding struct {
	Policyname string `json:"policyname,omitempty"`
	Priority   int    `json:"priority,omitempty"`
	Username   string `json:"username,omitempty"`
}

type SystemglobalAuditsyslogpolicyBinding struct {
	Builtin                []string `json:"builtin,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	Globalbindtype         string   `json:"globalbindtype,omitempty"`
	Gotopriorityexpression string   `json:"gotopriorityexpression,omitempty"`
	Nextfactor             string   `json:"nextfactor,omitempty"`
	Policyname             string   `json:"policyname,omitempty"`
	Priority               int      `json:"priority,omitempty"`
}

type Systemhwerror struct {
	Diskcheck    bool   `json:"diskcheck,omitempty"`
	Hwerrorcount int    `json:"hwerrorcount,omitempty"`
	Response     string `json:"response,omitempty"`
}

type Systemparameter struct {
	Allowdefaultpartition   string   `json:"allowdefaultpartition,omitempty"`
	Basicauth               string   `json:"basicauth,omitempty"`
	Cliloglevel             string   `json:"cliloglevel,omitempty"`
	Daystoexpire            int      `json:"daystoexpire,omitempty"`
	Doppler                 string   `json:"doppler,omitempty"`
	Fipsusermode            string   `json:"fipsusermode,omitempty"`
	Forcepasswordchange     string   `json:"forcepasswordchange,omitempty"`
	Googleanalytics         string   `json:"googleanalytics,omitempty"`
	Localauth               string   `json:"localauth,omitempty"`
	Maxclient               int      `json:"maxclient,omitempty"`
	Maxsessionperuser       int      `json:"maxsessionperuser,omitempty"`
	Minpasswordlen          int      `json:"minpasswordlen,omitempty"`
	Natpcbforceflushlimit   int      `json:"natpcbforceflushlimit,omitempty"`
	Natpcbrstontimeout      string   `json:"natpcbrstontimeout,omitempty"`
	Nextgenapiresource      string   `json:"_nextgenapiresource,omitempty"`
	Passwordhistorycontrol  string   `json:"passwordhistorycontrol,omitempty"`
	Promptstring            string   `json:"promptstring,omitempty"`
	Pwdhistorycount         int      `json:"pwdhistorycount,omitempty"`
	Rbaonresponse           string   `json:"rbaonresponse,omitempty"`
	Reauthonauthparamchange string   `json:"reauthonauthparamchange,omitempty"`
	Removesensitivefiles    string   `json:"removesensitivefiles,omitempty"`
	Restrictedtimeout       string   `json:"restrictedtimeout,omitempty"`
	Strongpassword          string   `json:"strongpassword,omitempty"`
	Timeout                 int      `json:"timeout,omitempty"`
	Totalauthtimeout        int      `json:"totalauthtimeout,omitempty"`
	Wafprotection           []string `json:"wafprotection,omitempty"`
	Warnpriorndays          int      `json:"warnpriorndays,omitempty"`
}

type SystemgroupSystemcmdpolicyBinding struct {
	Groupname  string `json:"groupname,omitempty"`
	Policyname string `json:"policyname,omitempty"`
	Priority   int    `json:"priority,omitempty"`
}

type SystemglobalAuthenticationlocalpolicyBinding struct {
	Builtin                []string `json:"builtin,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	Globalbindtype         string   `json:"globalbindtype,omitempty"`
	Gotopriorityexpression string   `json:"gotopriorityexpression,omitempty"`
	Nextfactor             string   `json:"nextfactor,omitempty"`
	Policyname             string   `json:"policyname,omitempty"`
	Priority               int      `json:"priority,omitempty"`
}

type Systemadmuserinfo struct {
	Username string `json:"username,omitempty"`
}

type Systemsession struct {
	All                   bool    `json:"all,omitempty"`
	Clientipaddress       string  `json:"clientipaddress,omitempty"`
	Clienttype            string  `json:"clienttype,omitempty"`
	Count                 float64 `json:"__count,omitempty"`
	Currentconn           bool    `json:"currentconn,omitempty"`
	Expirytime            int     `json:"expirytime,omitempty"`
	Lastactivitytime      string  `json:"lastactivitytime,omitempty"`
	Lastactivitytimelocal string  `json:"lastactivitytimelocal,omitempty"`
	Logintime             string  `json:"logintime,omitempty"`
	Logintimelocal        string  `json:"logintimelocal,omitempty"`
	Nextgenapiresource    string  `json:"_nextgenapiresource,omitempty"`
	Numofconnections      int     `json:"numofconnections,omitempty"`
	Partitionname         string  `json:"partitionname,omitempty"`
	Sid                   int     `json:"sid,omitempty"`
	Username              string  `json:"username,omitempty"`
}

type Systemautorestorefeature struct {
}

type SystemglobalBinding struct {
	SystemglobalAuditnslogpolicyBinding           []interface{} `json:"systemglobal_auditnslogpolicy_binding,omitempty"`
	SystemglobalAuditsyslogpolicyBinding          []interface{} `json:"systemglobal_auditsyslogpolicy_binding,omitempty"`
	SystemglobalAuthenticationldappolicyBinding   []interface{} `json:"systemglobal_authenticationldappolicy_binding,omitempty"`
	SystemglobalAuthenticationlocalpolicyBinding  []interface{} `json:"systemglobal_authenticationlocalpolicy_binding,omitempty"`
	SystemglobalAuthenticationpolicyBinding       []interface{} `json:"systemglobal_authenticationpolicy_binding,omitempty"`
	SystemglobalAuthenticationradiuspolicyBinding []interface{} `json:"systemglobal_authenticationradiuspolicy_binding,omitempty"`
	SystemglobalAuthenticationtacacspolicyBinding []interface{} `json:"systemglobal_authenticationtacacspolicy_binding,omitempty"`
}

type Systemkek struct {
	Level string `json:"level,omitempty"`
}

type Systemcpuparam struct {
	Count              float64 `json:"__count,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Pemode             string  `json:"pemode,omitempty"`
}

type Systemrestorepoint struct {
	Backupfilename     string  `json:"backupfilename,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Createdby          string  `json:"createdby,omitempty"`
	Creationtime       string  `json:"creationtime,omitempty"`
	Filename           string  `json:"filename,omitempty"`
	Ipaddress          string  `json:"ipaddress,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Techsuprtname      string  `json:"techsuprtname,omitempty"`
	Version            string  `json:"version,omitempty"`
}

type SystemglobalAuditnslogpolicyBinding struct {
	Builtin                []string `json:"builtin,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	Globalbindtype         string   `json:"globalbindtype,omitempty"`
	Gotopriorityexpression string   `json:"gotopriorityexpression,omitempty"`
	Nextfactor             string   `json:"nextfactor,omitempty"`
	Policyname             string   `json:"policyname,omitempty"`
	Priority               int      `json:"priority,omitempty"`
}

type Systemgroup struct {
	Allowedmanagementinterface []string `json:"allowedmanagementinterface,omitempty"`
	Count                      float64  `json:"__count,omitempty"`
	Daystoexpire               int      `json:"daystoexpire,omitempty"`
	Groupname                  string   `json:"groupname,omitempty"`
	Nextgenapiresource         string   `json:"_nextgenapiresource,omitempty"`
	Promptstring               string   `json:"promptstring,omitempty"`
	Timeout                    int      `json:"timeout,omitempty"`
	Warnpriorndays             int      `json:"warnpriorndays,omitempty"`
}

type SystemglobalAuthenticationradiuspolicyBinding struct {
	Builtin                []string `json:"builtin,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	Globalbindtype         string   `json:"globalbindtype,omitempty"`
	Gotopriorityexpression string   `json:"gotopriorityexpression,omitempty"`
	Nextfactor             string   `json:"nextfactor,omitempty"`
	Policyname             string   `json:"policyname,omitempty"`
	Priority               int      `json:"priority,omitempty"`
}

type SystemgroupSystemuserBinding struct {
	Groupname string `json:"groupname,omitempty"`
	Username  string `json:"username,omitempty"`
}

type SystemglobalAuthenticationtacacspolicyBinding struct {
	Builtin                []string `json:"builtin,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	Globalbindtype         string   `json:"globalbindtype,omitempty"`
	Gotopriorityexpression string   `json:"gotopriorityexpression,omitempty"`
	Nextfactor             string   `json:"nextfactor,omitempty"`
	Policyname             string   `json:"policyname,omitempty"`
	Priority               int      `json:"priority,omitempty"`
}

type SystemuserSystemgroupBinding struct {
	Groupname string `json:"groupname,omitempty"`
	Username  string `json:"username,omitempty"`
}

type Systemfipsstatus struct {
	Fipsstatus                                       string `json:"fipsstatus,omitempty"`
	Intelhwcryptographicacceleratorversion           string `json:"intelhwcryptographicacceleratorversion,omitempty"`
	Netscalercontrolplanecryptographiclibraryversion string `json:"netscalercontrolplanecryptographiclibraryversion,omitempty"`
	Netscalercrytographicmoduleversion               string `json:"netscalercrytographicmoduleversion,omitempty"`
	Netscalerdataplanecryptographiclibraryversion    string `json:"netscalerdataplanecryptographiclibraryversion,omitempty"`
	Netscalerjitterentropysourceversion              string `json:"netscalerjitterentropysourceversion,omitempty"`
	Nextgenapiresource                               string `json:"_nextgenapiresource,omitempty"`
}

type Systemextramgmtcpu struct {
	Configuredstate    string `json:"configuredstate,omitempty"`
	Effectivestate     string `json:"effectivestate,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Nodeid             int    `json:"nodeid,omitempty"`
}

type Systemsignedexereport struct {
	Message string `json:"message,omitempty"`
}

type Systemnsbtracing struct {
	Configuredstate    string `json:"configuredstate,omitempty"`
	Effectivestate     string `json:"effectivestate,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Nodeid             int    `json:"nodeid,omitempty"`
}

type Systemcmdpolicy struct {
	Action             string   `json:"action,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	Cmdspec            string   `json:"cmdspec,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Policyname         string   `json:"policyname,omitempty"`
}

type SystemgroupNspartitionBinding struct {
	Groupname     string `json:"groupname,omitempty"`
	Partitionname string `json:"partitionname,omitempty"`
}

type SystemglobalAuthenticationpolicyBinding struct {
	Builtin                []string `json:"builtin,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	Globalbindtype         string   `json:"globalbindtype,omitempty"`
	Gotopriorityexpression string   `json:"gotopriorityexpression,omitempty"`
	Nextfactor             string   `json:"nextfactor,omitempty"`
	Policyname             string   `json:"policyname,omitempty"`
	Priority               int      `json:"priority,omitempty"`
}

type SystemgroupBinding struct {
	Groupname                         string        `json:"groupname,omitempty"`
	SystemgroupNspartitionBinding     []interface{} `json:"systemgroup_nspartition_binding,omitempty"`
	SystemgroupSystemcmdpolicyBinding []interface{} `json:"systemgroup_systemcmdpolicy_binding,omitempty"`
	SystemgroupSystemuserBinding      []interface{} `json:"systemgroup_systemuser_binding,omitempty"`
}

type Systembackup struct {
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Createdby          string  `json:"createdby,omitempty"`
	Creationtime       string  `json:"creationtime,omitempty"`
	Filename           string  `json:"filename,omitempty"`
	Includekernel      string  `json:"includekernel,omitempty"`
	Ipaddress          string  `json:"ipaddress,omitempty"`
	Level              string  `json:"level,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Size               int     `json:"size,omitempty"`
	Skipbackup         bool    `json:"skipbackup,omitempty"`
	Uselocaltimezone   bool    `json:"uselocaltimezone,omitempty"`
	Version            string  `json:"version,omitempty"`
}

type Systemsshkey struct {
	Name               string `json:"name,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Src                string `json:"src,omitempty"`
	Sshkeytype         string `json:"sshkeytype,omitempty"`
}

type Systemuser struct {
	Allowedmanagementinterface     []string `json:"allowedmanagementinterface,omitempty"`
	Allowedmanagementinterfacekind string   `json:"allowedmanagementinterfacekind,omitempty"`
	Count                          float64  `json:"__count,omitempty"`
	Daystoexpirekind               string   `json:"daystoexpirekind,omitempty"`
	Encrypted                      bool     `json:"encrypted,omitempty"`
	Externalauth                   string   `json:"externalauth,omitempty"`
	Hashmethod                     string   `json:"hashmethod,omitempty"`
	Lastpwdchangetimestamp         int      `json:"lastpwdchangetimestamp,omitempty"`
	Logging                        string   `json:"logging,omitempty"`
	Maxsession                     int      `json:"maxsession,omitempty"`
	Nextgenapiresource             string   `json:"_nextgenapiresource,omitempty"`
	Password                       string   `json:"password,omitempty"`
	Promptinheritedfrom            string   `json:"promptinheritedfrom,omitempty"`
	Promptstring                   string   `json:"promptstring,omitempty"`
	Timeout                        int      `json:"timeout,omitempty"`
	Timeoutkind                    string   `json:"timeoutkind,omitempty"`
	Username                       string   `json:"username,omitempty"`
}

type SystemuserBinding struct {
	SystemuserNspartitionBinding     []interface{} `json:"systemuser_nspartition_binding,omitempty"`
	SystemuserSystemcmdpolicyBinding []interface{} `json:"systemuser_systemcmdpolicy_binding,omitempty"`
	SystemuserSystemgroupBinding     []interface{} `json:"systemuser_systemgroup_binding,omitempty"`
	Username                         string        `json:"username,omitempty"`
}

type Systemfile struct {
	Count              float64  `json:"__count,omitempty"`
	Fileaccesstime     string   `json:"fileaccesstime,omitempty"`
	Filecontent        string   `json:"filecontent,omitempty"`
	Fileencoding       string   `json:"fileencoding,omitempty"`
	Filelocation       string   `json:"filelocation,omitempty"`
	Filemode           []string `json:"filemode,omitempty"`
	Filemodifiedtime   string   `json:"filemodifiedtime,omitempty"`
	Filename           string   `json:"filename,omitempty"`
	Filesize           int      `json:"filesize,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
}

type SystemuserNspartitionBinding struct {
	Partitionname string `json:"partitionname,omitempty"`
	Username      string `json:"username,omitempty"`
}

type SystemStatus struct {
	AddiMgmtCPUUsagePcnt float64 `json:"addimgmtcpuusagepcnt,omitempty"`
	AuxTemp0             int     `json:"auxtemp0,omitempty"`
	AuxTemp1             int     `json:"auxtemp1,omitempty"`
	AuxTemp2             int     `json:"auxtemp2,omitempty"`
	AuxTemp3             int     `json:"auxtemp3,omitempty"`
	AuxVolt0             float64 `json:"auxvolt0,omitempty"`
	AuxVolt1             float64 `json:"auxvolt1,omitempty"`
	AuxVolt2             float64 `json:"auxvolt2,omitempty"`
	AuxVolt3             float64 `json:"auxvolt3,omitempty"`
	AuxVolt4             float64 `json:"auxvolt4,omitempty"`
	AuxVolt5             float64 `json:"auxvolt5,omitempty"`
	AuxVolt6             float64 `json:"auxvolt6,omitempty"`
	AuxVolt7             float64 `json:"auxvolt7,omitempty"`
	ClearStats           string  `json:"clearstats,omitempty"`
	Cpu0Temp             int     `json:"cpu0temp,omitempty"`
	Cpu1Temp             int     `json:"cpu1temp,omitempty"`
	CpuFan0Apeed         int     `json:"cpufan0speed,omitempty"`
	CpuFan1Apeed         int     `json:"cpufan1speed,omitempty"`
	CPUUsage             string  `json:"cpuusage,omitempty"`
	CPUUsagePcnt         float64 `json:"cpuusagepcnt,omitempty"`
	Disk0Avail           int     `json:"disk0avail,omitempty"`
	Disk0PerUsage        int     `json:"disk0perusage,omitempty"`
	Disk0Size            int     `json:"disk0size,omitempty"`
	Disk0Used            int     `json:"disk0used,omitempty"`
	Disk1Avail           int     `json:"disk1avail,omitempty"`
	Disk1perUsage        int     `json:"disk1perusage,omitempty"`
	Disk1Size            int     `json:"disk1size,omitempty"`
	Disk1Used            int     `json:"disk1used,omitempty"`
	Fan0Speed            int     `json:"fan0speed,omitempty"`
	Fan2Speed            int     `json:"fan2speed,omitempty"`
	Fan3Speed            int     `json:"fan3speed,omitempty"`
	Fan4Speed            int     `json:"fan4speed,omitempty"`
	Fan5Speed            int     `json:"fan5speed,omitempty"`
	FanSpeed             int     `json:"fanspeed,omitempty"`
	InternalTemp         int     `json:"internaltemp,omitempty"`
	MasterCPUUsage       string  `json:"mastercpuusage,omitempty"`
	MemSizeMB            string  `json:"memsizemb,omitempty"`
	MemUsagePcnt         float64 `json:"memusagepcnt,omitempty"`
	MemUseInMB           string  `json:"memuseinmb,omitempty"`
	MgmtCPU0UsagePcnt    float64 `json:"mgmtcpu0usagepcnt,omitempty"`
	MgmtCPUUsageCcnt     float64 `json:"mgmtcpuusagepcnt,omitempty"`
	NumCPUs              string  `json:"numcpus,omitempty"`
	PktCPUUsagePcnt      float64 `json:"pktcpuusagepcnt,omitempty"`
	PowerSupply1Status   string  `json:"powersupply1status,omitempty"`
	PowerSupply2Status   string  `json:"powersupply2status,omitempty"`
	PowerSupply3Status   string  `json:"powersupply3status,omitempty"`
	PowerSupply4Status   string  `json:"powersupply4status,omitempty"`
	ResCPUUsage          string  `json:"rescpuusage,omitempty"`
	ResCPUUsagePcnt      float64 `json:"rescpuusagepcnt,omitempty"`
	SlaveCPUUsage        string  `json:"slavecpuusage,omitempty"`
	StartTime            string  `json:"starttime,omitempty"`
	StartTimeLocal       string  `json:"starttimelocal,omitempty"`
	SystemFanSpeed       int     `json:"systemfanspeed,omitempty"`
	TimeSinceStart       string  `json:"timesincestart,omitempty"`
	VoltageV12n          float64 `json:"voltagev12n,omitempty"`
	VoltageV12p          float64 `json:"voltagev12p,omitempty"`
	VoltageV33Main       float64 `json:"voltagev33main,omitempty"`
	VoltageV33Stby       float64 `json:"voltagev33stby,omitempty"`
	VoltageV5n           float64 `json:"voltagev5n,omitempty"`
	VoltageV5p           float64 `json:"voltagev5p,omitempty"`
	VoltageV5sb          float64 `json:"voltagev5sb,omitempty"`
	VoltageVBat          float64 `json:"voltagevbat,omitempty"`
	VoltageVCC0          float64 `json:"voltagevcc0,omitempty"`
	VoltageVCC1          float64 `json:"voltagevcc1,omitempty"`
	VoltageVSen2         float64 `json:"voltagevsen2,omitempty"`
	VoltageVTT           float64 `json:"voltagevtt,omitempty"`
}

type SystemExtraMgmtCPU struct {
	Nodeid             int    `json:"nodeid,omitempty"`
	ConfiguredState    string `json:"configuredstate,omitempty"`
	EffectiveState     string `json:"effectivestate,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
}
