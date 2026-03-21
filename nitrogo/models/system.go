package models

// system configuration structs
type SystemGlobalAuthenticationLDAPPolicyBinding struct {
	Builtin                []string `json:"builtin,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	GlobalBindType         string   `json:"globalbindtype,omitempty"`
	GotoPriorityExpression string   `json:"gotopriorityexpression,omitempty"`
	NextFactor             string   `json:"nextfactor,omitempty"`
	PolicyName             string   `json:"policyname,omitempty"`
	Priority               int      `json:"priority,omitempty"`
}

type SystemUserSystemCmdPolicyBinding struct {
	PolicyName string `json:"policyname,omitempty"`
	Priority   int    `json:"priority,omitempty"`
	Username   string `json:"username,omitempty"`
}

type SystemGlobalAuditSyslogPolicyBinding struct {
	Builtin                []string `json:"builtin,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	GlobalBindType         string   `json:"globalbindtype,omitempty"`
	GotoPriorityExpression string   `json:"gotopriorityexpression,omitempty"`
	NextFactor             string   `json:"nextfactor,omitempty"`
	PolicyName             string   `json:"policyname,omitempty"`
	Priority               int      `json:"priority,omitempty"`
}

type SystemHWError struct {
	DiskCheck    bool   `json:"diskcheck,omitempty"`
	HWErrorCount int    `json:"hwerrorcount,omitempty"`
	Response     string `json:"response,omitempty"`
}

type SystemParameter struct {
	AllowDefaultPartition   string   `json:"allowdefaultpartition,omitempty"`
	BasicAuth               string   `json:"basicauth,omitempty"`
	CLILogLevel             string   `json:"cliloglevel,omitempty"`
	DaysToExpire            int      `json:"daystoexpire,omitempty"`
	Doppler                 string   `json:"doppler,omitempty"`
	FIPSUserMode            string   `json:"fipsusermode,omitempty"`
	ForcePasswordChange     string   `json:"forcepasswordchange,omitempty"`
	GoogleAnalytics         string   `json:"googleanalytics,omitempty"`
	LocalAuth               string   `json:"localauth,omitempty"`
	MaxClient               int      `json:"maxclient,omitempty"`
	MaxSessionPerUser       int      `json:"maxsessionperuser,omitempty"`
	MinPasswordLen          int      `json:"minpasswordlen,omitempty"`
	NATPCBForceFlushLimit   int      `json:"natpcbforceflushlimit,omitempty"`
	NATPCBRstOnTimeout      string   `json:"natpcbrstontimeout,omitempty"`
	NextGenAPIResource      string   `json:"_nextgenapiresource,omitempty"`
	PasswordHistoryControl  string   `json:"passwordhistorycontrol,omitempty"`
	PromptString            string   `json:"promptstring,omitempty"`
	PwdHistoryCount         int      `json:"pwdhistorycount,omitempty"`
	RBAOnResponse           string   `json:"rbaonresponse,omitempty"`
	ReAuthOnAuthParamChange string   `json:"reauthonauthparamchange,omitempty"`
	RemoveSensitiveFiles    string   `json:"removesensitivefiles,omitempty"`
	RestrictedTimeout       string   `json:"restrictedtimeout,omitempty"`
	StrongPassword          string   `json:"strongpassword,omitempty"`
	Timeout                 int      `json:"timeout,omitempty"`
	TotalAuthTimeout        int      `json:"totalauthtimeout,omitempty"`
	WAFProtection           []string `json:"wafprotection,omitempty"`
	WarnPriorNDays          int      `json:"warnpriorndays,omitempty"`
}

type SystemGroupSystemCmdPolicyBinding struct {
	GroupName  string `json:"groupname,omitempty"`
	PolicyName string `json:"policyname,omitempty"`
	Priority   int    `json:"priority,omitempty"`
}

type SystemGlobalAuthenticationLocalPolicyBinding struct {
	Builtin                []string `json:"builtin,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	GlobalBindType         string   `json:"globalbindtype,omitempty"`
	GotoPriorityExpression string   `json:"gotopriorityexpression,omitempty"`
	NextFactor             string   `json:"nextfactor,omitempty"`
	PolicyName             string   `json:"policyname,omitempty"`
	Priority               int      `json:"priority,omitempty"`
}

type SystemADMUserInfo struct {
	Username string `json:"username,omitempty"`
}

type SystemSession struct {
	All                   bool    `json:"all,omitempty"`
	ClientIPAddress       string  `json:"clientipaddress,omitempty"`
	ClientType            string  `json:"clienttype,omitempty"`
	Count                 float64 `json:"__count,omitempty"`
	CurrentConn           bool    `json:"currentconn,omitempty"`
	ExpiryTime            int     `json:"expirytime,omitempty"`
	LastActivityTime      string  `json:"lastactivitytime,omitempty"`
	LastActivityTimeLocal string  `json:"lastactivitytimelocal,omitempty"`
	LoginTime             string  `json:"logintime,omitempty"`
	LoginTimeLocal        string  `json:"logintimelocal,omitempty"`
	NextGenAPIResource    string  `json:"_nextgenapiresource,omitempty"`
	NumOfConnections      int     `json:"numofconnections,omitempty"`
	PartitionName         string  `json:"partitionname,omitempty"`
	SID                   int     `json:"sid,omitempty"`
	Username              string  `json:"username,omitempty"`
}

type SystemAutoRestoreFeature struct {
}

type SystemGlobalBinding struct {
	SystemGlobalAuditNSLogPolicyBinding           []any `json:"systemglobal_auditnslogpolicy_binding,omitempty"`
	SystemGlobalAuditSyslogPolicyBinding          []any `json:"systemglobal_auditsyslogpolicy_binding,omitempty"`
	SystemGlobalAuthenticationLDAPPolicyBinding   []any `json:"systemglobal_authenticationldappolicy_binding,omitempty"`
	SystemGlobalAuthenticationLocalPolicyBinding  []any `json:"systemglobal_authenticationlocalpolicy_binding,omitempty"`
	SystemGlobalAuthenticationPolicyBinding       []any `json:"systemglobal_authenticationpolicy_binding,omitempty"`
	SystemGlobalAuthenticationRADIUSPolicyBinding []any `json:"systemglobal_authenticationradiuspolicy_binding,omitempty"`
	SystemGlobalAuthenticationTACACSPolicyBinding []any `json:"systemglobal_authenticationtacacspolicy_binding,omitempty"`
}

type SystemKEK struct {
	Level string `json:"level,omitempty"`
}

type SystemCPUParam struct {
	Count              float64 `json:"__count,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	PEMode             string  `json:"pemode,omitempty"`
}

type SystemRestorePoint struct {
	BackupFileName     string  `json:"backupfilename,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	CreatedBy          string  `json:"createdby,omitempty"`
	CreationTime       string  `json:"creationtime,omitempty"`
	FileName           string  `json:"filename,omitempty"`
	IPAddress          string  `json:"ipaddress,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	TechSuprtName      string  `json:"techsuprtname,omitempty"`
	Version            string  `json:"version,omitempty"`
}

type SystemGlobalAuditNSLogPolicyBinding struct {
	Builtin                []string `json:"builtin,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	GlobalBindType         string   `json:"globalbindtype,omitempty"`
	GotoPriorityExpression string   `json:"gotopriorityexpression,omitempty"`
	NextFactor             string   `json:"nextfactor,omitempty"`
	PolicyName             string   `json:"policyname,omitempty"`
	Priority               int      `json:"priority,omitempty"`
}

type SystemGroup struct {
	AllowedManagementInterface []string `json:"allowedmanagementinterface,omitempty"`
	Count                      float64  `json:"__count,omitempty"`
	DaysToExpire               int      `json:"daystoexpire,omitempty"`
	GroupName                  string   `json:"groupname,omitempty"`
	NextGenAPIResource         string   `json:"_nextgenapiresource,omitempty"`
	PromptString               string   `json:"promptstring,omitempty"`
	Timeout                    int      `json:"timeout,omitempty"`
	WarnPriorNDays             int      `json:"warnpriorndays,omitempty"`
}

type SystemGlobalAuthenticationRADIUSPolicyBinding struct {
	Builtin                []string `json:"builtin,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	GlobalBindType         string   `json:"globalbindtype,omitempty"`
	GotoPriorityExpression string   `json:"gotopriorityexpression,omitempty"`
	NextFactor             string   `json:"nextfactor,omitempty"`
	PolicyName             string   `json:"policyname,omitempty"`
	Priority               int      `json:"priority,omitempty"`
}

type SystemGroupSystemUserBinding struct {
	GroupName string `json:"groupname,omitempty"`
	Username  string `json:"username,omitempty"`
}

type SystemGlobalAuthenticationTACACSPolicyBinding struct {
	Builtin                []string `json:"builtin,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	GlobalBindType         string   `json:"globalbindtype,omitempty"`
	GotoPriorityExpression string   `json:"gotopriorityexpression,omitempty"`
	NextFactor             string   `json:"nextfactor,omitempty"`
	PolicyName             string   `json:"policyname,omitempty"`
	Priority               int      `json:"priority,omitempty"`
}

type SystemUserSystemGroupBinding struct {
	GroupName string `json:"groupname,omitempty"`
	Username  string `json:"username,omitempty"`
}

type SystemFIPSStatus struct {
	FIPSStatus                                       string `json:"fipsstatus,omitempty"`
	IntelHWCryptographicAcceleratorVersion           string `json:"intelhwcryptographicacceleratorversion,omitempty"`
	NetScalerControlPlaneCryptographicLibraryVersion string `json:"netscalercontrolplanecryptographiclibraryversion,omitempty"`
	NetScalerCrytographicModuleVersion               string `json:"netscalercrytographicmoduleversion,omitempty"`
	NetScalerDataPlaneCryptographicLibraryVersion    string `json:"netscalerdataplanecryptographiclibraryversion,omitempty"`
	NetScalerJitterEntropySourceVersion              string `json:"netscalerjitterentropysourceversion,omitempty"`
	NextGenAPIResource                               string `json:"_nextgenapiresource,omitempty"`
}

type SystemSignedExeReport struct {
	Message string `json:"message,omitempty"`
}

type SystemNSBTracing struct {
	ConfiguredState    string `json:"configuredstate,omitempty"`
	EffectiveState     string `json:"effectivestate,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	NodeID             int    `json:"nodeid,omitempty"`
}

type SystemCmdPolicy struct {
	Action             string   `json:"action,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	CmdSpec            string   `json:"cmdspec,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	PolicyName         string   `json:"policyname,omitempty"`
}

type SystemGroupNSPartitionBinding struct {
	GroupName     string `json:"groupname,omitempty"`
	PartitionName string `json:"partitionname,omitempty"`
}

type SystemGlobalAuthenticationPolicyBinding struct {
	Builtin                []string `json:"builtin,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	GlobalBindType         string   `json:"globalbindtype,omitempty"`
	GotoPriorityExpression string   `json:"gotopriorityexpression,omitempty"`
	NextFactor             string   `json:"nextfactor,omitempty"`
	PolicyName             string   `json:"policyname,omitempty"`
	Priority               int      `json:"priority,omitempty"`
}

type SystemGroupBinding struct {
	GroupName                         string `json:"groupname,omitempty"`
	SystemGroupNSPartitionBinding     []any  `json:"systemgroup_nspartition_binding,omitempty"`
	SystemGroupSystemCmdPolicyBinding []any  `json:"systemgroup_systemcmdpolicy_binding,omitempty"`
	SystemGroupSystemUserBinding      []any  `json:"systemgroup_systemuser_binding,omitempty"`
}

type SystemBackup struct {
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	CreatedBy          string  `json:"createdby,omitempty"`
	CreationTime       string  `json:"creationtime,omitempty"`
	FileName           string  `json:"filename,omitempty"`
	IncludeKernel      string  `json:"includekernel,omitempty"`
	IPAddress          string  `json:"ipaddress,omitempty"`
	Level              string  `json:"level,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Size               int     `json:"size,omitempty"`
	SkipBackup         bool    `json:"skipbackup,omitempty"`
	UseLocalTimezone   bool    `json:"uselocaltimezone,omitempty"`
	Version            string  `json:"version,omitempty"`
}

type SystemSSHKey struct {
	Name               string `json:"name,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	Src                string `json:"src,omitempty"`
	SSHKeyType         string `json:"sshkeytype,omitempty"`
}

type SystemUser struct {
	AllowedManagementInterface     []string `json:"allowedmanagementinterface,omitempty"`
	AllowedManagementInterfaceKind string   `json:"allowedmanagementinterfacekind,omitempty"`
	Count                          float64  `json:"__count,omitempty"`
	DaysToExpireKind               string   `json:"daystoexpirekind,omitempty"`
	Encrypted                      bool     `json:"encrypted,omitempty"`
	ExternalAuth                   string   `json:"externalauth,omitempty"`
	HashMethod                     string   `json:"hashmethod,omitempty"`
	LastPwdChangeTimestamp         int      `json:"lastpwdchangetimestamp,omitempty"`
	Logging                        string   `json:"logging,omitempty"`
	MaxSession                     int      `json:"maxsession,omitempty"`
	NextGenAPIResource             string   `json:"_nextgenapiresource,omitempty"`
	Password                       string   `json:"password,omitempty"`
	PromptInheritedFrom            string   `json:"promptinheritedfrom,omitempty"`
	PromptString                   string   `json:"promptstring,omitempty"`
	Timeout                        int      `json:"timeout,omitempty"`
	TimeoutKind                    string   `json:"timeoutkind,omitempty"`
	Username                       string   `json:"username,omitempty"`
}

type SystemUserBinding struct {
	SystemUserNSPartitionBinding     []any  `json:"systemuser_nspartition_binding,omitempty"`
	SystemUserSystemCmdPolicyBinding []any  `json:"systemuser_systemcmdpolicy_binding,omitempty"`
	SystemUserSystemGroupBinding     []any  `json:"systemuser_systemgroup_binding,omitempty"`
	Username                         string `json:"username,omitempty"`
}

type SystemFile struct {
	Count              float64  `json:"__count,omitempty"`
	FileAccessTime     string   `json:"fileaccesstime,omitempty"`
	FileContent        string   `json:"filecontent,omitempty"`
	FileEncoding       string   `json:"fileencoding,omitempty"`
	FileLocation       string   `json:"filelocation,omitempty"`
	FileMode           []string `json:"filemode,omitempty"`
	FileModifiedTime   string   `json:"filemodifiedtime,omitempty"`
	FileName           string   `json:"filename,omitempty"`
	FileSize           int      `json:"filesize,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
}

type SystemUserNSPartitionBinding struct {
	PartitionName string `json:"partitionname,omitempty"`
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
	CPU0Temp             int     `json:"cpu0temp,omitempty"`
	CPU1Temp             int     `json:"cpu1temp,omitempty"`
	CPUFan0Speed         int     `json:"cpufan0speed,omitempty"`
	CPUFan1Speed         int     `json:"cpufan1speed,omitempty"`
	CPUUsage             string  `json:"cpuusage,omitempty"`
	CPUUsagePcnt         float64 `json:"cpuusagepcnt,omitempty"`
	Disk0Avail           int     `json:"disk0avail,omitempty"`
	Disk0PerUsage        int     `json:"disk0perusage,omitempty"`
	Disk0Size            int     `json:"disk0size,omitempty"`
	Disk0Used            int     `json:"disk0used,omitempty"`
	Disk1Avail           int     `json:"disk1avail,omitempty"`
	Disk1PerUsage        int     `json:"disk1perusage,omitempty"`
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
	MgmtCPUUsagePcnt     float64 `json:"mgmtcpuusagepcnt,omitempty"`
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
	VoltageV12N          float64 `json:"voltagev12n,omitempty"`
	VoltageV12P          float64 `json:"voltagev12p,omitempty"`
	VoltageV33Main       float64 `json:"voltagev33main,omitempty"`
	VoltageV33Stby       float64 `json:"voltagev33stby,omitempty"`
	VoltageV5N           float64 `json:"voltagev5n,omitempty"`
	VoltageV5P           float64 `json:"voltagev5p,omitempty"`
	VoltageV5SB          float64 `json:"voltagev5sb,omitempty"`
	VoltageVBat          float64 `json:"voltagevbat,omitempty"`
	VoltageVCC0          float64 `json:"voltagevcc0,omitempty"`
	VoltageVCC1          float64 `json:"voltagevcc1,omitempty"`
	VoltageVSen2         float64 `json:"voltagevsen2,omitempty"`
	VoltageVTT           float64 `json:"voltagevtt,omitempty"`
}

type SystemExtraMgmtCPU struct {
	NodeID             int    `json:"nodeid,omitempty"`
	ConfiguredState    string `json:"configuredstate,omitempty"`
	EffectiveState     string `json:"effectivestate,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
}

type SystemCollectionParam struct {
	CommunityName string `json:"communityname,omitempty"`
	LogLevel      string `json:"loglevel,omitempty"`
	DataPath      string `json:"datapath,omitempty"`
}

type SystemCore struct {
	Datasource string `json:"datasource,omitempty"`
	Response   string `json:"response,omitempty"`
}

type SystemCounterGroup struct {
	Datasource string `json:"datasource,omitempty"`
	Response   string `json:"response,omitempty"`
}

type SystemCounters struct {
	CounterGroup string `json:"countergroup,omitempty"`
	Datasource   string `json:"datasource,omitempty"`
	Response     string `json:"response,omitempty"`
}

type SystemDataSource struct {
	Datasource string `json:"datasource,omitempty"`
	Response   string `json:"response,omitempty"`
}

type SystemEntity struct {
	Type       string `json:"type,omitempty"`
	Datasource string `json:"datasource,omitempty"`
	Core       int    `json:"core,omitempty"`
	Response   string `json:"response,omitempty"`
}

type SystemEntityData struct {
	Type        string `json:"type,omitempty"`
	Name        string `json:"name,omitempty"`
	AllDeleted  string `json:"alldeleted,omitempty"`
	AllInactive string `json:"allinactive,omitempty"`
	Datasource  string `json:"datasource,omitempty"`
	Core        int    `json:"core,omitempty"`
	Counters    string `json:"counters,omitempty"`
	StartTime   string `json:"starttime,omitempty"`
	EndTime     string `json:"endtime,omitempty"`
	Last        int    `json:"last,omitempty"`
	Unit        string `json:"unit,omitempty"`
	Response    string `json:"response,omitempty"`
	StartUpdate string `json:"startupdate,omitempty"`
	LastUpdate  string `json:"lastupdate,omitempty"`
}

type SystemEntityType struct {
	Datasource string `json:"datasource,omitempty"`
	Response   string `json:"response,omitempty"`
}

type SystemEventHistory struct {
	StartTime  string `json:"starttime,omitempty"`
	EndTime    string `json:"endtime,omitempty"`
	Last       int    `json:"last,omitempty"`
	Unit       string `json:"unit,omitempty"`
	Datasource string `json:"datasource,omitempty"`
	Response   string `json:"response,omitempty"`
}

type SystemGlobalData struct {
	Counters     string  `json:"counters,omitempty"`
	CounterGroup string  `json:"countergroup,omitempty"`
	StartTime    string  `json:"starttime,omitempty"`
	EndTime      string  `json:"endtime,omitempty"`
	Last         int     `json:"last,omitempty"`
	Unit         string  `json:"unit,omitempty"`
	Datasource   string  `json:"datasource,omitempty"`
	Core         int     `json:"core,omitempty"`
	Response     string  `json:"response,omitempty"`
	StartUpdate  float64 `json:"startupdate,omitempty"`
	LastUpdate   float64 `json:"lastupdate,omitempty"`
}
