package models

// vpn configuration structs
type VPNClientlessAccessPolicyVPNVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type VPNGlobalAuditNSLogPolicyBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNICADTLSConnection struct {
	ChannelNumber      int     `json:"channelnumber,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	DestIP             string  `json:"destip,omitempty"`
	DestPort           int     `json:"destport,omitempty"`
	Domain             string  `json:"domain,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NodeID             int     `json:"nodeid,omitempty"`
	PeID               int     `json:"peid,omitempty"`
	SrcIP              string  `json:"srcip,omitempty"`
	SrcPort            int     `json:"srcport,omitempty"`
	Username           string  `json:"username,omitempty"`
}

type VPNStoreInfo struct {
	Count              float64 `json:"__count,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	StoreAPISupport    string  `json:"storeapisupport,omitempty"`
	StoreList          string  `json:"storelist,omitempty"`
	StoreServerIsSF    string  `json:"storeserverissf,omitempty"`
	StoreServerStatus  string  `json:"storeserverstatus,omitempty"`
	StoreStatus        string  `json:"storestatus,omitempty"`
	URL                string  `json:"url,omitempty"`
}

type VPNGlobalIntranetIPBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	IntranetIP             string `json:"intranetip,omitempty"`
	Netmask                string `json:"netmask,omitempty"`
}

type VPNURLPolicyVPNVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type VPNURLPolicyAAAUserBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type VPNEULA struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
}

type VPNSessionPolicyVPNVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type VPNSessionPolicyBinding struct {
	Name                              string        `json:"name,omitempty"`
	VPNSessionPolicyAAAGroupBinding   []interface{} `json:"vpnsessionpolicy_aaagroup_binding,omitempty"`
	VPNSessionPolicyAAAUserBinding    []interface{} `json:"vpnsessionpolicy_aaauser_binding,omitempty"`
	VPNSessionPolicyVPNGlobalBinding  []interface{} `json:"vpnsessionpolicy_vpnglobal_binding,omitempty"`
	VPNSessionPolicyVPNVServerBinding []interface{} `json:"vpnsessionpolicy_vpnvserver_binding,omitempty"`
}

type VPNVServerCSPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNVServerAuditSyslogPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNGlobalVPNClientlessAccessPolicyBinding struct {
	Builtin                []string `json:"builtin,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	GlobalBindType         string   `json:"globalbindtype,omitempty"`
	GotoPriorityExpression string   `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool     `json:"groupextraction,omitempty"`
	PolicyName             string   `json:"policyname,omitempty"`
	Priority               int      `json:"priority,omitempty"`
	Secondary              bool     `json:"secondary,omitempty"`
	TypeField              string   `json:"type,omitempty"`
}

type VPNGlobalSecurePrivateAccessURLBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	SecurePrivateAccessURL string `json:"secureprivateaccessurl,omitempty"`
}

type VPNVServerAppControllerBinding struct {
	ActType       int    `json:"acttype,omitempty"`
	AppController string `json:"appcontroller,omitempty"`
	Name          string `json:"name,omitempty"`
}

type VPNGlobalVPNSecurePrivateAccessProfileBinding struct {
	GotoPriorityExpression     string `json:"gotopriorityexpression,omitempty"`
	SecurePrivateAccessProfile string `json:"secureprivateaccessprofile,omitempty"`
}

type VPNURLPolicyAAAGroupBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type VPNGlobalAuthenticationCertPolicyBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNVServerAnalyticsProfileBinding struct {
	AnalyticsProfile string `json:"analyticsprofile,omitempty"`
	Name             string `json:"name,omitempty"`
}

type VPNIntranetApplication struct {
	ClientApplication   []string `json:"clientapplication,omitempty"`
	Count               float64  `json:"__count,omitempty"`
	DestIP              string   `json:"destip,omitempty"`
	DestPort            string   `json:"destport,omitempty"`
	Hostname            string   `json:"hostname,omitempty"`
	Interception        string   `json:"interception,omitempty"`
	IntranetApplication string   `json:"intranetapplication,omitempty"`
	IPAddress           string   `json:"ipaddress,omitempty"`
	IPRange             string   `json:"iprange,omitempty"`
	Netmask             string   `json:"netmask,omitempty"`
	NextGenAPIResource  string   `json:"_nextgenapiresource,omitempty"`
	Protocol            string   `json:"protocol,omitempty"`
	SpoofIIP            string   `json:"spoofiip,omitempty"`
	SrcIP               string   `json:"srcip,omitempty"`
	SrcPort             int      `json:"srcport,omitempty"`
}

type VPNGlobalVPNTrafficPolicyBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNURLAction struct {
	ActualURL          string  `json:"actualurl,omitempty"`
	ApplicationType    string  `json:"applicationtype,omitempty"`
	ClientlessAccess   string  `json:"clientlessaccess,omitempty"`
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	IconURL            string  `json:"iconurl,omitempty"`
	LinkName           string  `json:"linkname,omitempty"`
	Name               string  `json:"name,omitempty"`
	NewName            string  `json:"newname,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	SAMLSSOProfile     string  `json:"samlssoprofile,omitempty"`
	SSOType            string  `json:"ssotype,omitempty"`
	VServerName        string  `json:"vservername,omitempty"`
}

type VPNPCoIPConnection struct {
	All                bool    `json:"all,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	DestIP             string  `json:"destip,omitempty"`
	DestPort           int     `json:"destport,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NodeID             int     `json:"nodeid,omitempty"`
	PeID               int     `json:"peid,omitempty"`
	SrcIP              string  `json:"srcip,omitempty"`
	SrcPort            int     `json:"srcport,omitempty"`
	Username           string  `json:"username,omitempty"`
}

type VPNVServerVPNEULABinding struct {
	ActType int    `json:"acttype,omitempty"`
	EULA    string `json:"eula,omitempty"`
	Name    string `json:"name,omitempty"`
}

type VPNVServerAuthenticationNegotiatePolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNSFConfig struct {
	Count              float64  `json:"__count,omitempty"`
	Filename           string   `json:"filename,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	Vserver            []string `json:"vserver,omitempty"`
}

type VPNPortalTheme struct {
	BaseTheme          string  `json:"basetheme,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Feature            string  `json:"feature,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
}

type VPNTrafficPolicyVPNVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type VPNGlobalAuthenticationLDAPPolicyBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNGlobalVPNEULABinding struct {
	EULA                   string `json:"eula,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
}

type VPNParameter struct {
	AccessRestrictedPageRedirect string        `json:"accessrestrictedpageredirect,omitempty"`
	AdvancedClientlessVPNMode    string        `json:"advancedclientlessvpnmode,omitempty"`
	AllowedLoginGroups           string        `json:"allowedlogingroups,omitempty"`
	AllProtocolProxy             string        `json:"allprotocolproxy,omitempty"`
	AlwaysOnProfileName          string        `json:"alwaysonprofilename,omitempty"`
	AppTokenTimeout              int           `json:"apptokentimeout,omitempty"`
	AuthorizationGroup           string        `json:"authorizationgroup,omitempty"`
	AutoProxyURL                 string        `json:"autoproxyurl,omitempty"`
	BackendCertValidation        string        `json:"backendcertvalidation,omitempty"`
	BackendDTLS12                string        `json:"backenddtls12,omitempty"`
	BackendServerSNI             string        `json:"backendserversni,omitempty"`
	CitrixReceiverHome           string        `json:"citrixreceiverhome,omitempty"`
	ClientChoices                string        `json:"clientchoices,omitempty"`
	ClientCleanupPrompt          string        `json:"clientcleanupprompt,omitempty"`
	ClientConfiguration          []string      `json:"clientconfiguration,omitempty"`
	ClientDebug                  string        `json:"clientdebug,omitempty"`
	ClientIDleTimeout            int           `json:"clientidletimeout,omitempty"`
	ClientIDleTimeoutWarning     int           `json:"clientidletimeoutwarning,omitempty"`
	ClientlessModeURLEncoding    string        `json:"clientlessmodeurlencoding,omitempty"`
	ClientlessPersistentCookie   string        `json:"clientlesspersistentcookie,omitempty"`
	ClientlessVPNMode            string        `json:"clientlessvpnmode,omitempty"`
	ClientOptions                []string      `json:"clientoptions,omitempty"`
	ClientSecurity               string        `json:"clientsecurity,omitempty"`
	ClientSecurityGroup          string        `json:"clientsecuritygroup,omitempty"`
	ClientSecurityLog            string        `json:"clientsecuritylog,omitempty"`
	ClientSecurityMessage        string        `json:"clientsecuritymessage,omitempty"`
	ClientVersions               string        `json:"clientversions,omitempty"`
	DefaultAuthorizationAction   string        `json:"defaultauthorizationaction,omitempty"`
	DevicePosture                string        `json:"deviceposture,omitempty"`
	DNSVServerName               string        `json:"dnsvservername,omitempty"`
	EmailHome                    string        `json:"emailhome,omitempty"`
	EncryptCSECExp               string        `json:"encryptcsecexp,omitempty"`
	EPAClientType                string        `json:"epaclienttype,omitempty"`
	ForceCleanup                 []string      `json:"forcecleanup,omitempty"`
	ForcedTimeout                int           `json:"forcedtimeout,omitempty"`
	ForcedTimeoutWarning         int           `json:"forcedtimeoutwarning,omitempty"`
	FQDNSpoofedIP                string        `json:"fqdnspoofedip,omitempty"`
	FTPProxy                     string        `json:"ftpproxy,omitempty"`
	GopherProxy                  string        `json:"gopherproxy,omitempty"`
	HomePage                     string        `json:"homepage,omitempty"`
	HTTPPort                     []interface{} `json:"httpport,omitempty"`
	HTTPProxy                    string        `json:"httpproxy,omitempty"`
	HTTPTrackConnProxy           string        `json:"httptrackconnproxy,omitempty"`
	ICAProxy                     string        `json:"icaproxy,omitempty"`
	ICASessionTimeout            string        `json:"icasessiontimeout,omitempty"`
	ICAUserAccounting            string        `json:"icauseraccounting,omitempty"`
	IconWithReceiver             string        `json:"iconwithreceiver,omitempty"`
	IIPDNSSuffix                 string        `json:"iipdnssuffix,omitempty"`
	KCDAccount                   string        `json:"kcdaccount,omitempty"`
	KillConnections              string        `json:"killconnections,omitempty"`
	LinuxPluginUpgrade           string        `json:"linuxpluginupgrade,omitempty"`
	LocalLANAccess               string        `json:"locallanaccess,omitempty"`
	LoginScript                  string        `json:"loginscript,omitempty"`
	LogoutScript                 string        `json:"logoutscript,omitempty"`
	MacPluginUpgrade             string        `json:"macpluginupgrade,omitempty"`
	MaxIIPPerUser                int           `json:"maxiipperuser,omitempty"`
	MDXTokenTimeout              int           `json:"mdxtokentimeout,omitempty"`
	Name                         string        `json:"name,omitempty"`
	Netmask                      string        `json:"netmask,omitempty"`
	NextGenAPIResource           string        `json:"_nextgenapiresource,omitempty"`
	NTDomain                     string        `json:"ntdomain,omitempty"`
	PCoIPProfileName             string        `json:"pcoipprofilename,omitempty"`
	Proxy                        string        `json:"proxy,omitempty"`
	ProxyException               string        `json:"proxyexception,omitempty"`
	ProxyLocalBypass             string        `json:"proxylocalbypass,omitempty"`
	RDPClientProfileName         string        `json:"rdpclientprofilename,omitempty"`
	RFC1918                      string        `json:"rfc1918,omitempty"`
	SameSite                     string        `json:"samesite,omitempty"`
	SecureBrowse                 string        `json:"securebrowse,omitempty"`
	SecurePrivateAccess          string        `json:"secureprivateaccess,omitempty"`
	SecurePrivateAccessProfile   string        `json:"secureprivateaccessprofile,omitempty"`
	SessTimeout                  int           `json:"sesstimeout,omitempty"`
	SmartGroup                   string        `json:"smartgroup,omitempty"`
	SocksProxy                   string        `json:"socksproxy,omitempty"`
	SplitDNS                     string        `json:"splitdns,omitempty"`
	SplitTunnel                  string        `json:"splittunnel,omitempty"`
	SpoofIIP                     string        `json:"spoofiip,omitempty"`
	SSLProxy                     string        `json:"sslproxy,omitempty"`
	SSO                          string        `json:"sso,omitempty"`
	SSOCredential                string        `json:"ssocredential,omitempty"`
	StoreFrontURL                string        `json:"storefronturl,omitempty"`
	TransparentInterception      string        `json:"transparentinterception,omitempty"`
	UITheme                      string        `json:"uitheme,omitempty"`
	UseIIP                       string        `json:"useiip,omitempty"`
	UseMIP                       string        `json:"usemip,omitempty"`
	UserDomains                  string        `json:"userdomains,omitempty"`
	VPNSessionPolicyBindType     string        `json:"vpnsessionpolicybindtype,omitempty"`
	VPNSessionPolicyCount        int           `json:"vpnsessionpolicycount,omitempty"`
	WIHome                       string        `json:"wihome,omitempty"`
	WIHomeAddressType            string        `json:"wihomeaddresstype,omitempty"`
	WindowsAutoLogon             string        `json:"windowsautologon,omitempty"`
	WindowsClientType            string        `json:"windowsclienttype,omitempty"`
	WindowsPluginUpgrade         string        `json:"windowspluginupgrade,omitempty"`
	Winsip                       string        `json:"winsip,omitempty"`
	WIPortalMode                 string        `json:"wiportalmode,omitempty"`
}

type VPNClientlessAccessProfile struct {
	Builtin                        []string `json:"builtin,omitempty"`
	ClientConsumedCookies          string   `json:"clientconsumedcookies,omitempty"`
	Count                          float64  `json:"__count,omitempty"`
	CSSRewritePolicyLabel          string   `json:"cssrewritepolicylabel,omitempty"`
	Description                    string   `json:"description,omitempty"`
	Feature                        string   `json:"feature,omitempty"`
	IsDefault                      bool     `json:"isdefault,omitempty"`
	JavascriptRewritePolicyLabel   string   `json:"javascriptrewritepolicylabel,omitempty"`
	NextGenAPIResource             string   `json:"_nextgenapiresource,omitempty"`
	ProfileName                    string   `json:"profilename,omitempty"`
	RegexForFindingCustomURLs      string   `json:"regexforfindingcustomurls,omitempty"`
	RegexForFindingURLInCSS        string   `json:"regexforfindingurlincss,omitempty"`
	RegexForFindingURLInJavascript string   `json:"regexforfindingurlinjavascript,omitempty"`
	RegexForFindingURLInXComponent string   `json:"regexforfindingurlinxcomponent,omitempty"`
	RegexForFindingURLInXML        string   `json:"regexforfindingurlinxml,omitempty"`
	ReqHdrRewritePolicyLabel       string   `json:"reqhdrrewritepolicylabel,omitempty"`
	RequirePersistentCookie        string   `json:"requirepersistentcookie,omitempty"`
	ResHdrRewritePolicyLabel       string   `json:"reshdrrewritepolicylabel,omitempty"`
	URLRewritePolicyLabel          string   `json:"urlrewritepolicylabel,omitempty"`
	XComponentRewritePolicyLabel   string   `json:"xcomponentrewritepolicylabel,omitempty"`
	XMLRewritePolicyLabel          string   `json:"xmlrewritepolicylabel,omitempty"`
}

type VPNSAMLSSOProfile struct {
	AssertionConsumerServiceURL string  `json:"assertionconsumerserviceurl,omitempty"`
	Attribute1                  string  `json:"attribute1,omitempty"`
	Attribute10                 string  `json:"attribute10,omitempty"`
	Attribute10Expr             string  `json:"attribute10expr,omitempty"`
	Attribute10Format           string  `json:"attribute10format,omitempty"`
	Attribute10FriendlyName     string  `json:"attribute10friendlyname,omitempty"`
	Attribute11                 string  `json:"attribute11,omitempty"`
	Attribute11Expr             string  `json:"attribute11expr,omitempty"`
	Attribute11Format           string  `json:"attribute11format,omitempty"`
	Attribute11FriendlyName     string  `json:"attribute11friendlyname,omitempty"`
	Attribute12                 string  `json:"attribute12,omitempty"`
	Attribute12Expr             string  `json:"attribute12expr,omitempty"`
	Attribute12Format           string  `json:"attribute12format,omitempty"`
	Attribute12FriendlyName     string  `json:"attribute12friendlyname,omitempty"`
	Attribute13                 string  `json:"attribute13,omitempty"`
	Attribute13Expr             string  `json:"attribute13expr,omitempty"`
	Attribute13Format           string  `json:"attribute13format,omitempty"`
	Attribute13FriendlyName     string  `json:"attribute13friendlyname,omitempty"`
	Attribute14                 string  `json:"attribute14,omitempty"`
	Attribute14Expr             string  `json:"attribute14expr,omitempty"`
	Attribute14Format           string  `json:"attribute14format,omitempty"`
	Attribute14FriendlyName     string  `json:"attribute14friendlyname,omitempty"`
	Attribute15                 string  `json:"attribute15,omitempty"`
	Attribute15Expr             string  `json:"attribute15expr,omitempty"`
	Attribute15Format           string  `json:"attribute15format,omitempty"`
	Attribute15FriendlyName     string  `json:"attribute15friendlyname,omitempty"`
	Attribute16                 string  `json:"attribute16,omitempty"`
	Attribute16Expr             string  `json:"attribute16expr,omitempty"`
	Attribute16Format           string  `json:"attribute16format,omitempty"`
	Attribute16FriendlyName     string  `json:"attribute16friendlyname,omitempty"`
	Attribute1Expr              string  `json:"attribute1expr,omitempty"`
	Attribute1Format            string  `json:"attribute1format,omitempty"`
	Attribute1FriendlyName      string  `json:"attribute1friendlyname,omitempty"`
	Attribute2                  string  `json:"attribute2,omitempty"`
	Attribute2Expr              string  `json:"attribute2expr,omitempty"`
	Attribute2Format            string  `json:"attribute2format,omitempty"`
	Attribute2FriendlyName      string  `json:"attribute2friendlyname,omitempty"`
	Attribute3                  string  `json:"attribute3,omitempty"`
	Attribute3Expr              string  `json:"attribute3expr,omitempty"`
	Attribute3Format            string  `json:"attribute3format,omitempty"`
	Attribute3FriendlyName      string  `json:"attribute3friendlyname,omitempty"`
	Attribute4                  string  `json:"attribute4,omitempty"`
	Attribute4Expr              string  `json:"attribute4expr,omitempty"`
	Attribute4Format            string  `json:"attribute4format,omitempty"`
	Attribute4FriendlyName      string  `json:"attribute4friendlyname,omitempty"`
	Attribute5                  string  `json:"attribute5,omitempty"`
	Attribute5Expr              string  `json:"attribute5expr,omitempty"`
	Attribute5Format            string  `json:"attribute5format,omitempty"`
	Attribute5FriendlyName      string  `json:"attribute5friendlyname,omitempty"`
	Attribute6                  string  `json:"attribute6,omitempty"`
	Attribute6Expr              string  `json:"attribute6expr,omitempty"`
	Attribute6Format            string  `json:"attribute6format,omitempty"`
	Attribute6FriendlyName      string  `json:"attribute6friendlyname,omitempty"`
	Attribute7                  string  `json:"attribute7,omitempty"`
	Attribute7Expr              string  `json:"attribute7expr,omitempty"`
	Attribute7Format            string  `json:"attribute7format,omitempty"`
	Attribute7FriendlyName      string  `json:"attribute7friendlyname,omitempty"`
	Attribute8                  string  `json:"attribute8,omitempty"`
	Attribute8Expr              string  `json:"attribute8expr,omitempty"`
	Attribute8Format            string  `json:"attribute8format,omitempty"`
	Attribute8FriendlyName      string  `json:"attribute8friendlyname,omitempty"`
	Attribute9                  string  `json:"attribute9,omitempty"`
	Attribute9Expr              string  `json:"attribute9expr,omitempty"`
	Attribute9Format            string  `json:"attribute9format,omitempty"`
	Attribute9FriendlyName      string  `json:"attribute9friendlyname,omitempty"`
	Audience                    string  `json:"audience,omitempty"`
	Count                       float64 `json:"__count,omitempty"`
	DigestMethod                string  `json:"digestmethod,omitempty"`
	EncryptAssertion            string  `json:"encryptassertion,omitempty"`
	EncryptionAlgorithm         string  `json:"encryptionalgorithm,omitempty"`
	Name                        string  `json:"name,omitempty"`
	NameIDExpr                  string  `json:"nameidexpr,omitempty"`
	NameIDFormat                string  `json:"nameidformat,omitempty"`
	NextGenAPIResource          string  `json:"_nextgenapiresource,omitempty"`
	RelayStateRule              string  `json:"relaystaterule,omitempty"`
	SAMLIssuerName              string  `json:"samlissuername,omitempty"`
	SAMLSigningCertName         string  `json:"samlsigningcertname,omitempty"`
	SAMLSPCertName              string  `json:"samlspcertname,omitempty"`
	SendPassword                string  `json:"sendpassword,omitempty"`
	SignAssertion               string  `json:"signassertion,omitempty"`
	SignatureAlg                string  `json:"signaturealg,omitempty"`
	SignatureService            string  `json:"signatureservice,omitempty"`
	SkewTime                    int     `json:"skewtime,omitempty"`
}

type VPNVServerAuditNSLogPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNGlobalVPNSessionPolicyBinding struct {
	Builtin                []string `json:"builtin,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	GotoPriorityExpression string   `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool     `json:"groupextraction,omitempty"`
	PolicyName             string   `json:"policyname,omitempty"`
	Priority               int      `json:"priority,omitempty"`
	Secondary              bool     `json:"secondary,omitempty"`
}

type VPNPCoIPVServerProfile struct {
	Count              float64 `json:"__count,omitempty"`
	LoginDomain        string  `json:"logindomain,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	UDPPort            int     `json:"udpport,omitempty"`
}

type VPNSessionPolicyVPNGlobalBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type VPNVServerIntranetIP6Binding struct {
	ActType     int    `json:"acttype,omitempty"`
	IntranetIP6 string `json:"intranetip6,omitempty"`
	Name        string `json:"name,omitempty"`
	NumAddr     int    `json:"numaddr,omitempty"`
}

type VPNVServerFEOPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNGlobalAuditSyslogPolicyBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNVServerVPNSessionPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNVServerVPNURLBinding struct {
	ActType int    `json:"acttype,omitempty"`
	Name    string `json:"name,omitempty"`
	URLName string `json:"urlname,omitempty"`
}

type VPNVServerVPNURLPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNVServerAuthenticationWebAuthPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNClientlessAccessPolicy struct {
	Builtin            []string `json:"builtin,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Description        string   `json:"description,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	IsDefault          bool     `json:"isdefault,omitempty"`
	Name               string   `json:"name,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	ProfileName        string   `json:"profilename,omitempty"`
	Rule               string   `json:"rule,omitempty"`
	UndefAction        string   `json:"undefaction,omitempty"`
	UndefHits          int      `json:"undefhits,omitempty"`
}

type VPNVServerAAAPreauthenticationPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNTrafficPolicyBinding struct {
	Name                              string        `json:"name,omitempty"`
	VPNTrafficPolicyAAAGroupBinding   []interface{} `json:"vpntrafficpolicy_aaagroup_binding,omitempty"`
	VPNTrafficPolicyAAAUserBinding    []interface{} `json:"vpntrafficpolicy_aaauser_binding,omitempty"`
	VPNTrafficPolicyVPNGlobalBinding  []interface{} `json:"vpntrafficpolicy_vpnglobal_binding,omitempty"`
	VPNTrafficPolicyVPNVServerBinding []interface{} `json:"vpntrafficpolicy_vpnvserver_binding,omitempty"`
}

type VPNVServerAuthenticationSAMLIDPPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNNextHopServer struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NextHopFQDN        string  `json:"nexthopfqdn,omitempty"`
	NextHopIP          string  `json:"nexthopip,omitempty"`
	NextHopPort        int     `json:"nexthopport,omitempty"`
	ResAddressType     string  `json:"resaddresstype,omitempty"`
	Secure             string  `json:"secure,omitempty"`
}

type VPNURLPolicy struct {
	Action             string   `json:"action,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	Comment            string   `json:"comment,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	LogAction          string   `json:"logaction,omitempty"`
	Name               string   `json:"name,omitempty"`
	NewName            string   `json:"newname,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	Rule               string   `json:"rule,omitempty"`
	UndefHits          int      `json:"undefhits,omitempty"`
}

type VPNGlobalVPNURLPolicyBinding struct {
	Builtin                []string `json:"builtin,omitempty"`
	GotoPriorityExpression string   `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool     `json:"groupextraction,omitempty"`
	PolicyName             string   `json:"policyname,omitempty"`
	Priority               int      `json:"priority,omitempty"`
	Secondary              bool     `json:"secondary,omitempty"`
}

type VPNSessionPolicy struct {
	Action             string   `json:"action,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	ExpressionType     string   `json:"expressiontype,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	Name               string   `json:"name,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	Rule               string   `json:"rule,omitempty"`
}

type VPNURL struct {
	ActualURL          string  `json:"actualurl,omitempty"`
	AppJSON            string  `json:"appjson,omitempty"`
	ApplicationType    string  `json:"applicationtype,omitempty"`
	ClientlessAccess   string  `json:"clientlessaccess,omitempty"`
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	IconURL            string  `json:"iconurl,omitempty"`
	LinkName           string  `json:"linkname,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	SAMLSSOProfile     string  `json:"samlssoprofile,omitempty"`
	SSOType            string  `json:"ssotype,omitempty"`
	URLName            string  `json:"urlname,omitempty"`
	VServerName        string  `json:"vservername,omitempty"`
}

type VPNGlobalAuthenticationPolicyBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNVServerAuthenticationTACACSPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNGlobalVPNNextHopServerBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	NextHopServer          string `json:"nexthopserver,omitempty"`
}

type VPNVServerICAPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNVServerAuthenticationCertPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNVServerAuthenticationSAMLPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNURLPolicyVPNGlobalBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type VPNTrafficPolicyAAAGroupBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type VPNVServerCachePolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNGlobalVPNPortalThemeBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	PortalTheme            string `json:"portaltheme,omitempty"`
}

type VPNVServerShareFileServerBinding struct {
	ActType   int    `json:"acttype,omitempty"`
	Name      string `json:"name,omitempty"`
	ShareFile string `json:"sharefile,omitempty"`
}

type VPNGlobalAuthenticationSAMLPolicyBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNTrafficPolicy struct {
	Action             string  `json:"action,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	ExpressionType     string  `json:"expressiontype,omitempty"`
	Hits               int     `json:"hits,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Rule               string  `json:"rule,omitempty"`
}

type VPNVServerAuthenticationLoginSchemaPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNVServerVPNTrafficPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNSessionPolicyAAAUserBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type VPNVServerVPNIntranetApplicationBinding struct {
	ActType             int    `json:"acttype,omitempty"`
	IntranetApplication string `json:"intranetapplication,omitempty"`
	Name                string `json:"name,omitempty"`
}

type VPNVServerAppFlowPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNGlobalAuthenticationRADIUSPolicyBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNClientlessAccessPolicyVPNGlobalBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type VPNGlobalAuthenticationLocalPolicyBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNGlobalVPNURLBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	URLName                string `json:"urlname,omitempty"`
}

type VPNSessionPolicyAAAGroupBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type VPNVServerBinding struct {
	Name                                             string        `json:"name,omitempty"`
	VPNVServerAAAPreauthenticationPolicyBinding      []interface{} `json:"vpnvserver_aaapreauthenticationpolicy_binding,omitempty"`
	VPNVServerAnalyticsProfileBinding                []interface{} `json:"vpnvserver_analyticsprofile_binding,omitempty"`
	VPNVServerAppControllerBinding                   []interface{} `json:"vpnvserver_appcontroller_binding,omitempty"`
	VPNVServerAppFlowPolicyBinding                   []interface{} `json:"vpnvserver_appflowpolicy_binding,omitempty"`
	VPNVServerAppFWPolicyBinding                     []interface{} `json:"vpnvserver_appfwpolicy_binding,omitempty"`
	VPNVServerAuditNSLogPolicyBinding                []interface{} `json:"vpnvserver_auditnslogpolicy_binding,omitempty"`
	VPNVServerAuditSyslogPolicyBinding               []interface{} `json:"vpnvserver_auditsyslogpolicy_binding,omitempty"`
	VPNVServerAuthenticationCertPolicyBinding        []interface{} `json:"vpnvserver_authenticationcertpolicy_binding,omitempty"`
	VPNVServerAuthenticationDFAPolicyBinding         []interface{} `json:"vpnvserver_authenticationdfapolicy_binding,omitempty"`
	VPNVServerAuthenticationLDAPPolicyBinding        []interface{} `json:"vpnvserver_authenticationldappolicy_binding,omitempty"`
	VPNVServerAuthenticationLocalPolicyBinding       []interface{} `json:"vpnvserver_authenticationlocalpolicy_binding,omitempty"`
	VPNVServerAuthenticationLoginSchemaPolicyBinding []interface{} `json:"vpnvserver_authenticationloginschemapolicy_binding,omitempty"`
	VPNVServerAuthenticationNegotiatePolicyBinding   []interface{} `json:"vpnvserver_authenticationnegotiatepolicy_binding,omitempty"`
	VPNVServerAuthenticationOAuthIDPPolicyBinding    []interface{} `json:"vpnvserver_authenticationoauthidppolicy_binding,omitempty"`
	VPNVServerAuthenticationPolicyBinding            []interface{} `json:"vpnvserver_authenticationpolicy_binding,omitempty"`
	VPNVServerAuthenticationRADIUSPolicyBinding      []interface{} `json:"vpnvserver_authenticationradiuspolicy_binding,omitempty"`
	VPNVServerAuthenticationSAMLIDPPolicyBinding     []interface{} `json:"vpnvserver_authenticationsamlidppolicy_binding,omitempty"`
	VPNVServerAuthenticationSAMLPolicyBinding        []interface{} `json:"vpnvserver_authenticationsamlpolicy_binding,omitempty"`
	VPNVServerAuthenticationTACACSPolicyBinding      []interface{} `json:"vpnvserver_authenticationtacacspolicy_binding,omitempty"`
	VPNVServerAuthenticationWebAuthPolicyBinding     []interface{} `json:"vpnvserver_authenticationwebauthpolicy_binding,omitempty"`
	VPNVServerCachePolicyBinding                     []interface{} `json:"vpnvserver_cachepolicy_binding,omitempty"`
	VPNVServerCSPolicyBinding                        []interface{} `json:"vpnvserver_cspolicy_binding,omitempty"`
	VPNVServerFEOPolicyBinding                       []interface{} `json:"vpnvserver_feopolicy_binding,omitempty"`
	VPNVServerICAPolicyBinding                       []interface{} `json:"vpnvserver_icapolicy_binding,omitempty"`
	VPNVServerIntranetIP6Binding                     []interface{} `json:"vpnvserver_intranetip6_binding,omitempty"`
	VPNVServerIntranetIPBinding                      []interface{} `json:"vpnvserver_intranetip_binding,omitempty"`
	VPNVServerResponderPolicyBinding                 []interface{} `json:"vpnvserver_responderpolicy_binding,omitempty"`
	VPNVServerRewritePolicyBinding                   []interface{} `json:"vpnvserver_rewritepolicy_binding,omitempty"`
	VPNVServerSecurePrivateAccessURLBinding          []interface{} `json:"vpnvserver_secureprivateaccessurl_binding,omitempty"`
	VPNVServerShareFileServerBinding                 []interface{} `json:"vpnvserver_sharefileserver_binding,omitempty"`
	VPNVServerSTAServerBinding                       []interface{} `json:"vpnvserver_staserver_binding,omitempty"`
	VPNVServerVPNClientlessAccessPolicyBinding       []interface{} `json:"vpnvserver_vpnclientlessaccesspolicy_binding,omitempty"`
	VPNVServerVPNEPAProfileBinding                   []interface{} `json:"vpnvserver_vpnepaprofile_binding,omitempty"`
	VPNVServerVPNEULABinding                         []interface{} `json:"vpnvserver_vpneula_binding,omitempty"`
	VPNVServerVPNIntranetApplicationBinding          []interface{} `json:"vpnvserver_vpnintranetapplication_binding,omitempty"`
	VPNVServerVPNNextHopServerBinding                []interface{} `json:"vpnvserver_vpnnexthopserver_binding,omitempty"`
	VPNVServerVPNPortalThemeBinding                  []interface{} `json:"vpnvserver_vpnportaltheme_binding,omitempty"`
	VPNVServerVPNSecurePrivateAccessProfileBinding   []interface{} `json:"vpnvserver_vpnsecureprivateaccessprofile_binding,omitempty"`
	VPNVServerVPNSessionPolicyBinding                []interface{} `json:"vpnvserver_vpnsessionpolicy_binding,omitempty"`
	VPNVServerVPNTrafficPolicyBinding                []interface{} `json:"vpnvserver_vpntrafficpolicy_binding,omitempty"`
	VPNVServerVPNURLBinding                          []interface{} `json:"vpnvserver_vpnurl_binding,omitempty"`
	VPNVServerVPNURLPolicyBinding                    []interface{} `json:"vpnvserver_vpnurlpolicy_binding,omitempty"`
}

type VPNSessionAction struct {
	AdvancedClientlessVPNMode  string        `json:"advancedclientlessvpnmode,omitempty"`
	AllowedLoginGroups         string        `json:"allowedlogingroups,omitempty"`
	AllProtocolProxy           string        `json:"allprotocolproxy,omitempty"`
	AlwaysOnProfileName        string        `json:"alwaysonprofilename,omitempty"`
	AuthorizationGroup         string        `json:"authorizationgroup,omitempty"`
	AutoProxyURL               string        `json:"autoproxyurl,omitempty"`
	Builtin                    []string      `json:"builtin,omitempty"`
	CitrixReceiverHome           string        `json:"citrixreceiverhome,omitempty"`
	ClientChoices              string        `json:"clientchoices,omitempty"`
	ClientCleanupPrompt          string        `json:"clientcleanupprompt,omitempty"`
	ClientConfiguration        []string      `json:"clientconfiguration,omitempty"`
	ClientDebug                string        `json:"clientdebug,omitempty"`
	ClientIDleTimeout          int           `json:"clientidletimeout,omitempty"`
	ClientIDleTimeoutWarning   int           `json:"clientidletimeoutwarning,omitempty"`
	ClientlessModeURLEncoding  string        `json:"clientlessmodeurlencoding,omitempty"`
	ClientlessPersistentCookie string        `json:"clientlesspersistentcookie,omitempty"`
	ClientlessVPNMode          string        `json:"clientlessvpnmode,omitempty"`
	ClientOptions              string        `json:"clientoptions,omitempty"`
	ClientSecurity             string        `json:"clientsecurity,omitempty"`
	ClientSecurityGroup        string        `json:"clientsecuritygroup,omitempty"`
	ClientSecurityLog          string        `json:"clientsecuritylog,omitempty"`
	ClientSecurityMessage      string        `json:"clientsecuritymessage,omitempty"`
	Count                      float64       `json:"__count,omitempty"`
	DefaultAuthorizationAction string        `json:"defaultauthorizationaction,omitempty"`
	DNSVServerName             string        `json:"dnsvservername,omitempty"`
	EmailHome                  string        `json:"emailhome,omitempty"`
	EPAClientType              string        `json:"epaclienttype,omitempty"`
	Feature                    string        `json:"feature,omitempty"`
	ForceCleanup               []string      `json:"forcecleanup,omitempty"`
	ForcedTimeout              int           `json:"forcedtimeout,omitempty"`
	ForcedTimeoutWarning       int           `json:"forcedtimeoutwarning,omitempty"`
	FQDNSpoofedIP              string        `json:"fqdnspoofedip,omitempty"`
	FTPProxy                   string        `json:"ftpproxy,omitempty"`
	GopherProxy                string        `json:"gopherproxy,omitempty"`
	HomePage                   string        `json:"homepage,omitempty"`
	HTTPPort                   []interface{} `json:"httpport,omitempty"`
	HTTPProxy                  string        `json:"httpproxy,omitempty"`
	ICAProxy                   string        `json:"icaproxy,omitempty"`
	IconWithReceiver             string        `json:"iconwithreceiver,omitempty"`
	IIPDNSSuffix               string        `json:"iipdnssuffix,omitempty"`
	KCDAccount                 string        `json:"kcdaccount,omitempty"`
	KillConnections              string        `json:"killconnections,omitempty"`
	LinuxPluginUpgrade           string        `json:"linuxpluginupgrade,omitempty"`
	LocalLANAccess               string        `json:"locallanaccess,omitempty"`
	LoginScript                string        `json:"loginscript,omitempty"`
	LogoutScript               string        `json:"logoutscript,omitempty"`
	MacPluginUpgrade             string        `json:"macpluginupgrade,omitempty"`
	Name                       string        `json:"name,omitempty"`
	Netmask                    string        `json:"netmask,omitempty"`
	NextGenAPIResource         string        `json:"_nextgenapiresource,omitempty"`
	NTDomain                   string        `json:"ntdomain,omitempty"`
	PCoIPProfileName           string        `json:"pcoipprofilename,omitempty"`
	Proxy                      string        `json:"proxy,omitempty"`
	ProxyException             string        `json:"proxyexception,omitempty"`
	ProxyLocalBypass           string        `json:"proxylocalbypass,omitempty"`
	RDPClientProfileName       string        `json:"rdpclientprofilename,omitempty"`
	RFC1918                    string        `json:"rfc1918,omitempty"`
	SecureBrowse               string        `json:"securebrowse,omitempty"`
	SessTimeout                  int           `json:"sesstimeout,omitempty"`
	SFGatewayAuthType          string        `json:"sfgatewayauthtype,omitempty"`
	SmartGroup                   string        `json:"smartgroup,omitempty"`
	SocksProxy                   string        `json:"socksproxy,omitempty"`
	SplitDNS                   string        `json:"splitdns,omitempty"`
	SplitTunnel                  string        `json:"splittunnel,omitempty"`
	SpoofIIP                   string        `json:"spoofiip,omitempty"`
	SSLProxy                     string        `json:"sslproxy,omitempty"`
	SSO                        string        `json:"sso,omitempty"`
	SSOCredential                string        `json:"ssocredential,omitempty"`
	StoreFrontURL                string        `json:"storefronturl,omitempty"`
	TransparentInterception    string        `json:"transparentinterception,omitempty"`
	UseIIP                       string        `json:"useiip,omitempty"`
	UseMIP                       string        `json:"usemip,omitempty"`
	UserAccounting             string        `json:"useraccounting,omitempty"`
	WIHome                       string        `json:"wihome,omitempty"`
	WIHomeAddressType            string        `json:"wihomeaddresstype,omitempty"`
	WindowsAutoLogon             string        `json:"windowsautologon,omitempty"`
	WindowsClientType            string        `json:"windowsclienttype,omitempty"`
	WindowsPluginUpgrade         string        `json:"windowspluginupgrade,omitempty"`
	Winsip                       string        `json:"winsip,omitempty"`
	WIPortalMode                 string        `json:"wiportalmode,omitempty"`
}

type VPNVServerAuthenticationLDAPPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNTrafficAction struct {
	AppTimeout         int     `json:"apptimeout,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	FormSSOAction      string  `json:"formssoaction,omitempty"`
	Fta                string  `json:"fta,omitempty"`
	Hdx                string  `json:"hdx,omitempty"`
	KCDAccount         string  `json:"kcdaccount,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	PasswdExpression   string  `json:"passwdexpression,omitempty"`
	Proxy              string  `json:"proxy,omitempty"`
	Qual               string  `json:"qual,omitempty"`
	SAMLSSOProfile     string  `json:"samlssoprofile,omitempty"`
	SSO                string  `json:"sso,omitempty"`
	UserExpression     string  `json:"userexpression,omitempty"`
	Wanscaler          string  `json:"wanscaler,omitempty"`
}

type VPNVServerVPNSecurePrivateAccessProfileBinding struct {
	ActType                    int    `json:"acttype,omitempty"`
	Name                       string `json:"name,omitempty"`
	SecurePrivateAccessProfile string `json:"secureprivateaccessprofile,omitempty"`
}

type VPNVServerVPNEPAProfileBinding struct {
	ActType            int    `json:"acttype,omitempty"`
	EPAProfile         string `json:"epaprofile,omitempty"`
	EPAProfileOptional bool   `json:"epaprofileoptional,omitempty"`
	Name               string `json:"name,omitempty"`
}

type VPNFormSSOAction struct {
	ActionURL          string  `json:"actionurl,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NameValuePair      string  `json:"namevaluepair,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NvType             string  `json:"nvtype,omitempty"`
	PasswdField        string  `json:"passwdfield,omitempty"`
	ResponseSize       int     `json:"responsesize,omitempty"`
	SSOSuccessRule     string  `json:"ssosuccessrule,omitempty"`
	SubmitMethod       string  `json:"submitmethod,omitempty"`
	UserField          string  `json:"userfield,omitempty"`
}

type VPNVServerAuthenticationRADIUSPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNVServerAuthenticationLocalPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNTrafficPolicyVPNGlobalBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type VPNURLPolicyBinding struct {
	Name                          string        `json:"name,omitempty"`
	VPNURLPolicyAAAGroupBinding   []interface{} `json:"vpnurlpolicy_aaagroup_binding,omitempty"`
	VPNURLPolicyAAAUserBinding    []interface{} `json:"vpnurlpolicy_aaauser_binding,omitempty"`
	VPNURLPolicyVPNGlobalBinding  []interface{} `json:"vpnurlpolicy_vpnglobal_binding,omitempty"`
	VPNURLPolicyVPNVServerBinding []interface{} `json:"vpnurlpolicy_vpnvserver_binding,omitempty"`
}

type VPNVServerResponderPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNVServerSecurePrivateAccessURLBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	Name                   string `json:"name,omitempty"`
	SecurePrivateAccessURL string `json:"secureprivateaccessurl,omitempty"`
}

type VPNVServerAuthenticationPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNGlobalBinding struct {
	VPNGlobalAppControllerBinding                 []interface{} `json:"vpnglobal_appcontroller_binding,omitempty"`
	VPNGlobalAppFWPolicyBinding                   []interface{} `json:"vpnglobal_appfwpolicy_binding,omitempty"`
	VPNGlobalAuditNSLogPolicyBinding              []interface{} `json:"vpnglobal_auditnslogpolicy_binding,omitempty"`
	VPNGlobalAuditSyslogPolicyBinding             []interface{} `json:"vpnglobal_auditsyslogpolicy_binding,omitempty"`
	VPNGlobalAuthenticationCertPolicyBinding      []interface{} `json:"vpnglobal_authenticationcertpolicy_binding,omitempty"`
	VPNGlobalAuthenticationLDAPPolicyBinding      []interface{} `json:"vpnglobal_authenticationldappolicy_binding,omitempty"`
	VPNGlobalAuthenticationLocalPolicyBinding     []interface{} `json:"vpnglobal_authenticationlocalpolicy_binding,omitempty"`
	VPNGlobalAuthenticationNegotiatePolicyBinding []interface{} `json:"vpnglobal_authenticationnegotiatepolicy_binding,omitempty"`
	VPNGlobalAuthenticationPolicyBinding          []interface{} `json:"vpnglobal_authenticationpolicy_binding,omitempty"`
	VPNGlobalAuthenticationRADIUSPolicyBinding    []interface{} `json:"vpnglobal_authenticationradiuspolicy_binding,omitempty"`
	VPNGlobalAuthenticationSAMLPolicyBinding      []interface{} `json:"vpnglobal_authenticationsamlpolicy_binding,omitempty"`
	VPNGlobalAuthenticationTACACSPolicyBinding    []interface{} `json:"vpnglobal_authenticationtacacspolicy_binding,omitempty"`
	VPNGlobalGSLBDomainBinding                    []interface{} `json:"vpnglobal_gslbdomain_binding,omitempty"`
	VPNGlobalIntranetIP6Binding                   []interface{} `json:"vpnglobal_intranetip6_binding,omitempty"`
	VPNGlobalIntranetIPBinding                    []interface{} `json:"vpnglobal_intranetip_binding,omitempty"`
	VPNGlobalSecurePrivateAccessURLBinding        []interface{} `json:"vpnglobal_secureprivateaccessurl_binding,omitempty"`
	VPNGlobalShareFileServerBinding               []interface{} `json:"vpnglobal_sharefileserver_binding,omitempty"`
	VPNGlobalSSLCertKeyBinding                    []interface{} `json:"vpnglobal_sslcertkey_binding,omitempty"`
	VPNGlobalSTAServerBinding                     []interface{} `json:"vpnglobal_staserver_binding,omitempty"`
	VPNGlobalVPNClientlessAccessPolicyBinding     []interface{} `json:"vpnglobal_vpnclientlessaccesspolicy_binding,omitempty"`
	VPNGlobalVPNEULABinding                       []interface{} `json:"vpnglobal_vpneula_binding,omitempty"`
	VPNGlobalVPNIntranetApplicationBinding        []interface{} `json:"vpnglobal_vpnintranetapplication_binding,omitempty"`
	VPNGlobalVPNNextHopServerBinding              []interface{} `json:"vpnglobal_vpnnexthopserver_binding,omitempty"`
	VPNGlobalVPNPortalThemeBinding                []interface{} `json:"vpnglobal_vpnportaltheme_binding,omitempty"`
	VPNGlobalVPNSecurePrivateAccessProfileBinding []interface{} `json:"vpnglobal_vpnsecureprivateaccessprofile_binding,omitempty"`
	VPNGlobalVPNSessionPolicyBinding              []interface{} `json:"vpnglobal_vpnsessionpolicy_binding,omitempty"`
	VPNGlobalVPNTrafficPolicyBinding              []interface{} `json:"vpnglobal_vpntrafficpolicy_binding,omitempty"`
	VPNGlobalVPNURLBinding                        []interface{} `json:"vpnglobal_vpnurl_binding,omitempty"`
	VPNGlobalVPNURLPolicyBinding                  []interface{} `json:"vpnglobal_vpnurlpolicy_binding,omitempty"`
}

type VPNGlobalAppControllerBinding struct {
	AppController          string `json:"appcontroller,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
}

type VPNVServerAuthenticationDFAPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNPCoIPProfile struct {
	ConServerURL       string  `json:"conserverurl,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	IcvVerification    string  `json:"icvverification,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	SessionIdleTimeout int     `json:"sessionidletimeout,omitempty"`
}

type VPNGlobalShareFileServerBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	ShareFile              string `json:"sharefile,omitempty"`
}

type VPNTrafficPolicyAAAUserBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type VPNAlwaysOnProfile struct {
	ClientControl             string  `json:"clientcontrol,omitempty"`
	Count                     float64 `json:"__count,omitempty"`
	LocationBasedVPN          string  `json:"locationbasedvpn,omitempty"`
	Name                      string  `json:"name,omitempty"`
	NetworkAccessOnVPNFailure string  `json:"networkaccessonvpnfailure,omitempty"`
	NextGenAPIResource        string  `json:"_nextgenapiresource,omitempty"`
}

type VPNVServerVPNNextHopServerBinding struct {
	ActType       int    `json:"acttype,omitempty"`
	Name          string `json:"name,omitempty"`
	NextHopServer string `json:"nexthopserver,omitempty"`
}

type VPNGlobalAppFWPolicyBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNVServerSTAServerBinding struct {
	ActType        int    `json:"acttype,omitempty"`
	Name           string `json:"name,omitempty"`
	STAAddressType string `json:"staaddresstype,omitempty"`
	STAAuthID      string `json:"staauthid,omitempty"`
	STAServer      string `json:"staserver,omitempty"`
	STAState       string `json:"stastate,omitempty"`
}

type VPNGlobalVPNIntranetApplicationBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	IntranetApplication    string `json:"intranetapplication,omitempty"`
}

type VPNVServerIntranetIPBinding struct {
	ActType    int    `json:"acttype,omitempty"`
	IntranetIP string `json:"intranetip,omitempty"`
	MapField   string `json:"map,omitempty"`
	Name       string `json:"name,omitempty"`
	Netmask    string `json:"netmask,omitempty"`
}

type VPNClientlessAccessPolicyBinding struct {
	Name                                       string        `json:"name,omitempty"`
	VPNClientlessAccessPolicyVPNGlobalBinding  []interface{} `json:"vpnclientlessaccesspolicy_vpnglobal_binding,omitempty"`
	VPNClientlessAccessPolicyVPNVServerBinding []interface{} `json:"vpnclientlessaccesspolicy_vpnvserver_binding,omitempty"`
}

type VPNGlobalAuthenticationNegotiatePolicyBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNGlobalAuthenticationTACACSPolicyBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNGlobalIntranetIP6Binding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	IntranetIP6            string `json:"intranetip6,omitempty"`
	NumAddr                int    `json:"numaddr,omitempty"`
}

type VPNVServer struct {
	AccessRestrictedPageRedirect string  `json:"accessrestrictedpageredirect,omitempty"`
	AdvancedEPA                  string  `json:"advancedepa,omitempty"`
	AppFlowLog                   string  `json:"appflowlog,omitempty"`
	Authentication               string  `json:"authentication,omitempty"`
	AuthnProfile                 string  `json:"authnprofile,omitempty"`
	BackupVServer                string  `json:"backupvserver,omitempty"`
	BindPoint                    string  `json:"bindpoint,omitempty"`
	CacheType                    string  `json:"cachetype,omitempty"`
	CacheVServer                 string  `json:"cachevserver,omitempty"`
	CertKeyNames                 string  `json:"certkeynames,omitempty"`
	CGInfraHomePageRedirect      string  `json:"cginfrahomepageredirect,omitempty"`
	CltTimeout                   int     `json:"clttimeout,omitempty"`
	Comment                      string  `json:"comment,omitempty"`
	Count                        float64 `json:"__count,omitempty"`
	CSVServer                    string  `json:"csvserver,omitempty"`
	CurAAAUsers                  int     `json:"curaaausers,omitempty"`
	CurState                     string  `json:"curstate,omitempty"`
	CurTotalUsers                int     `json:"curtotalusers,omitempty"`
	DeploymentType               string  `json:"deploymenttype,omitempty"`
	DeviceCert                   string  `json:"devicecert,omitempty"`
	DevicePosture                string  `json:"deviceposture,omitempty"`
	DisablePrimaryOnDown         string  `json:"disableprimaryondown,omitempty"`
	Domain                       string  `json:"domain,omitempty"`
	DoubleHop                    string  `json:"doublehop,omitempty"`
	DownStateFlush               string  `json:"downstateflush,omitempty"`
	DTLS                         string  `json:"dtls,omitempty"`
	EPAProfileOptional           bool    `json:"epaprofileoptional,omitempty"`
	FailedLoginTimeout           int     `json:"failedlogintimeout,omitempty"`
	GroupExtraction              bool    `json:"groupextraction,omitempty"`
	HTTPProfileName              string  `json:"httpprofilename,omitempty"`
	ICAOnly                      string  `json:"icaonly,omitempty"`
	ICAProxySessionMigration     string  `json:"icaproxysessionmigration,omitempty"`
	ICMPVSRResponse              string  `json:"icmpvsrresponse,omitempty"`
	IP                           string  `json:"ip,omitempty"`
	IPSet                        string  `json:"ipset,omitempty"`
	IPv46                        string  `json:"ipv46,omitempty"`
	L2Conn                       string  `json:"l2conn,omitempty"`
	LinuxEPAPluginUpgrade        string  `json:"linuxepapluginupgrade,omitempty"`
	ListenPolicy                 string  `json:"listenpolicy,omitempty"`
	ListenPriority               int     `json:"listenpriority,omitempty"`
	LoginOnce                    string  `json:"loginonce,omitempty"`
	LogoutOnSmartcardRemoval     string  `json:"logoutonsmartcardremoval,omitempty"`
	MacEPAPluginUpgrade          string  `json:"macepapluginupgrade,omitempty"`
	MapField                     string  `json:"map,omitempty"`
	MaxAAAUsers                  int     `json:"maxaaausers,omitempty"`
	MaxLoginAttempts             int     `json:"maxloginattempts,omitempty"`
	Name                         string  `json:"name,omitempty"`
	NetProfile                   string  `json:"netprofile,omitempty"`
	NewName                      string  `json:"newname,omitempty"`
	NextGenAPIResource           string  `json:"_nextgenapiresource,omitempty"`
	NGName                       string  `json:"ngname,omitempty"`
	NoDefaultBindings            string  `json:"nodefaultbindings,omitempty"`
	PCoIPVServerProfileName      string  `json:"pcoipvserverprofilename,omitempty"`
	Port                         int     `json:"port,omitempty"`
	Precedence                   string  `json:"precedence,omitempty"`
	QUICProfileName              string  `json:"quicprofilename,omitempty"`
	Range                        int     `json:"range,omitempty"`
	RDPServerProfileName         string  `json:"rdpserverprofilename,omitempty"`
	Redirect                     string  `json:"redirect,omitempty"`
	RedirectURL                  string  `json:"redirecturl,omitempty"`
	Response                     string  `json:"response,omitempty"`
	RHIState                     string  `json:"rhistate,omitempty"`
	Rule                         string  `json:"rule,omitempty"`
	SameSite                     string  `json:"samesite,omitempty"`
	Secondary                    bool    `json:"secondary,omitempty"`
	SecurePrivateAccess          string  `json:"secureprivateaccess,omitempty"`
	ServiceName                  string  `json:"servicename,omitempty"`
	ServiceType                  string  `json:"servicetype,omitempty"`
	SOMethod                     string  `json:"somethod,omitempty"`
	SOPersistence                string  `json:"sopersistence,omitempty"`
	SOPersistenceTimeout         int     `json:"sopersistencetimeout,omitempty"`
	SOThreshold                  int     `json:"sothreshold,omitempty"`
	State                        string  `json:"state,omitempty"`
	Status                       int     `json:"status,omitempty"`
	TCPProfileName               string  `json:"tcpprofilename,omitempty"`
	TypeField                    string  `json:"type,omitempty"`
	UseMIP                       string  `json:"usemip,omitempty"`
	UserDomains                  string  `json:"userdomains,omitempty"`
	Value                        string  `json:"value,omitempty"`
	VServerFQDN                  string  `json:"vserverfqdn,omitempty"`
	Weight                       int     `json:"weight,omitempty"`
	WindowsEPAPluginUpgrade      string  `json:"windowsepapluginupgrade,omitempty"`
}

type VPNGlobalSSLCertKeyBinding struct {
	CACert                 string `json:"cacert,omitempty"`
	CertKeyName            string `json:"certkeyname,omitempty"`
	CRLCheck               string `json:"crlcheck,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	OCSPCheck              string `json:"ocspcheck,omitempty"`
	UserDataEncryptionKey  string `json:"userdataencryptionkey,omitempty"`
}

type VPNVServerAppFWPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNICAConnection struct {
	All                bool    `json:"all,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	DestIP             string  `json:"destip,omitempty"`
	DestPort           int     `json:"destport,omitempty"`
	Domain             string  `json:"domain,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NodeID             int     `json:"nodeid,omitempty"`
	PeID               int     `json:"peid,omitempty"`
	ProductName        string  `json:"productname,omitempty"`
	SrcIP              string  `json:"srcip,omitempty"`
	SrcPort            int     `json:"srcport,omitempty"`
	TenantName         string  `json:"tenantname,omitempty"`
	TransProto         string  `json:"transproto,omitempty"`
	Username           string  `json:"username,omitempty"`
}

type VPNVServerRewritePolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNVServerAuthenticationOAuthIDPPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VPNGlobalSTAServerBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	STAAddressType         string `json:"staaddresstype,omitempty"`
	STAAuthID              string `json:"staauthid,omitempty"`
	STAServer              string `json:"staserver,omitempty"`
	STAState               string `json:"stastate,omitempty"`
}

type VPNSecurePrivateAccessProfile struct {
	AccessRestrictedPageRedirect string  `json:"accessrestrictedpageredirect,omitempty"`
	ChromeEnterprisePremiumMode  string  `json:"chromeenterprisepremiummode,omitempty"`
	CloudDeployment              string  `json:"clouddeployment,omitempty"`
	Count                        float64 `json:"__count,omitempty"`
	CustomerID                   string  `json:"customerid,omitempty"`
	Name                         string  `json:"name,omitempty"`
	NextGenAPIResource           string  `json:"_nextgenapiresource,omitempty"`
	URL                          string  `json:"url,omitempty"`
}

type VPNEPAProfile struct {
	Count              float64 `json:"__count,omitempty"`
	Data               string  `json:"data,omitempty"`
	Filename           string  `json:"filename,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
}

type VPNVServerVPNPortalThemeBinding struct {
	ActType     int    `json:"acttype,omitempty"`
	Name        string `json:"name,omitempty"`
	PortalTheme string `json:"portaltheme,omitempty"`
}

type VPNGlobalDomainBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	IntranetDomain         string `json:"intranetdomain,omitempty"`
}

type VPNVServerVPNClientlessAccessPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}
