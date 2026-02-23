package models

// aaa configuration structs
type AAACertParams struct {
	DefaultAuthenticationGroup string `json:"defaultauthenticationgroup,omitempty"`
	GroupNameField             string `json:"groupnamefield,omitempty"`
	NextGenAPIResource         string `json:"_nextgenapiresource,omitempty"`
	TwoFactor                  string `json:"twofactor,omitempty"`
	UsernameField              string `json:"usernamefield,omitempty"`
}

type AAAGlobalAAAPreAuthenticationPolicyBinding struct {
	BindPolicyType int      `json:"bindpolicytype,omitempty"`
	Builtin        []string `json:"builtin,omitempty"`
	Policy         string   `json:"policy,omitempty"`
	Priority       int      `json:"priority,omitempty"`
}

type AAAGlobalAuthenticationNegotiateActionBinding struct {
	WindowsProfile string `json:"windowsprofile,omitempty"`
}

type AAAGlobalBinding struct {
	AAAGlobalAAAPreAuthenticationPolicyBinding    []interface{} `json:"aaaglobal_aaapreauthenticationpolicy_binding,omitempty"`
	AAAGlobalAuthenticationNegotiateActionBinding []interface{} `json:"aaaglobal_authenticationnegotiateaction_binding,omitempty"`
}

type AAAGroup struct {
	Count              float64 `json:"__count,omitempty"`
	GroupName          string  `json:"groupname,omitempty"`
	LoggedIn           bool    `json:"loggedin,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Weight             int     `json:"weight,omitempty"`
}

type AAAGroupAAAUserBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupName              string `json:"groupname,omitempty"`
	Username               string `json:"username,omitempty"`
}

type AAAGroupAuditNSLogPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupName              string `json:"groupname,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
}

type AAAGroupAuditSyslogPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupName              string `json:"groupname,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
}

type AAAGroupAuthorizationPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupName              string `json:"groupname,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
}

type AAAGroupBinding struct {
	AAAGroupAAAUserBinding                       []interface{} `json:"aaagroup_aaauser_binding,omitempty"`
	AAAGroupAuditNSLogPolicyBinding              []interface{} `json:"aaagroup_auditnslogpolicy_binding,omitempty"`
	AAAGroupAuditSyslogPolicyBinding             []interface{} `json:"aaagroup_auditsyslogpolicy_binding,omitempty"`
	AAAGroupAuthorizationPolicyBinding           []interface{} `json:"aaagroup_authorizationpolicy_binding,omitempty"`
	AAAGroupIntranetIP6Binding                   []interface{} `json:"aaagroup_intranetip6_binding,omitempty"`
	AAAGroupIntranetIPBinding                    []interface{} `json:"aaagroup_intranetip_binding,omitempty"`
	AAAGroupTMSessionPolicyBinding               []interface{} `json:"aaagroup_tmsessionpolicy_binding,omitempty"`
	AAAGroupVPNIntranetApplicationBinding        []interface{} `json:"aaagroup_vpnintranetapplication_binding,omitempty"`
	AAAGroupVPNSecurePrivateAccessProfileBinding []interface{} `json:"aaagroup_vpnsecureprivateaccessprofile_binding,omitempty"`
	AAAGroupVPNSessionPolicyBinding              []interface{} `json:"aaagroup_vpnsessionpolicy_binding,omitempty"`
	AAAGroupVPNTrafficPolicyBinding              []interface{} `json:"aaagroup_vpntrafficpolicy_binding,omitempty"`
	AAAGroupVPNURLBinding                        []interface{} `json:"aaagroup_vpnurl_binding,omitempty"`
	AAAGroupVPNURLPolicyBinding                  []interface{} `json:"aaagroup_vpnurlpolicy_binding,omitempty"`
	GroupName                                    string        `json:"groupname,omitempty"`
}

type AAAGroupIntranetIP6Binding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupName              string `json:"groupname,omitempty"`
	IntranetIP6            string `json:"intranetip6,omitempty"`
	NumAddr                int    `json:"numaddr,omitempty"`
}

type AAAGroupIntranetIPBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupName              string `json:"groupname,omitempty"`
	IntranetIP             string `json:"intranetip,omitempty"`
	Netmask                string `json:"netmask,omitempty"`
}

type AAAGroupTMSessionPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupName              string `json:"groupname,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
}

type AAAGroupVPNIntranetApplicationBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupName              string `json:"groupname,omitempty"`
	IntranetApplication    string `json:"intranetapplication,omitempty"`
}

type AAAGroupVPNSecurePrivateAccessProfileBinding struct {
	ActType                    int    `json:"acttype,omitempty"`
	GotoPriorityExpression     string `json:"gotopriorityexpression,omitempty"`
	GroupName                  string `json:"groupname,omitempty"`
	SecurePrivateAccessProfile string `json:"secureprivateaccessprofile,omitempty"`
}

type AAAGroupVPNSessionPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupName              string `json:"groupname,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
}

type AAAGroupVPNTrafficPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupName              string `json:"groupname,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
}

type AAAGroupVPNURLBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupName              string `json:"groupname,omitempty"`
	URLName                string `json:"urlname,omitempty"`
}

type AAAGroupVPNURLPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupName              string `json:"groupname,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
}

type AAAKCDAccount struct {
	CACert             string  `json:"cacert,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	DelegatedUser      string  `json:"delegateduser,omitempty"`
	EnterpriseRealm    string  `json:"enterpriserealm,omitempty"`
	KCDAccount         string  `json:"kcdaccount,omitempty"`
	KCDPassword        string  `json:"kcdpassword,omitempty"`
	KCDSPN             string  `json:"kcdspn,omitempty"`
	KeyTab             string  `json:"keytab,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Principle          string  `json:"principle,omitempty"`
	RealmStr           string  `json:"realmstr,omitempty"`
	SaltExpression     string  `json:"saltexpression,omitempty"`
	ServiceSPN         string  `json:"servicespn,omitempty"`
	UserCert           string  `json:"usercert,omitempty"`
	UserRealm          string  `json:"userrealm,omitempty"`
}

type AAALDAPParams struct {
	AuthTimeout                int      `json:"authtimeout,omitempty"`
	Builtin                    []string `json:"builtin,omitempty"`
	DefaultAuthenticationGroup string   `json:"defaultauthenticationgroup,omitempty"`
	Feature                    string   `json:"feature,omitempty"`
	GroupAttrName              string   `json:"groupattrname,omitempty"`
	GroupAuthName              string   `json:"groupauthname,omitempty"`
	GroupNameIdentifier        string   `json:"groupnameidentifier,omitempty"`
	GroupSearchAttribute       string   `json:"groupsearchattribute,omitempty"`
	GroupSearchFilter          string   `json:"groupsearchfilter,omitempty"`
	GroupSearchSubAttribute    string   `json:"groupsearchsubattribute,omitempty"`
	LDAPBase                   string   `json:"ldapbase,omitempty"`
	LDAPBindDN                 string   `json:"ldapbinddn,omitempty"`
	LDAPBindDNPassword         string   `json:"ldapbinddnpassword,omitempty"`
	LDAPLoginName              string   `json:"ldaploginname,omitempty"`
	MaxNestingLevel            int      `json:"maxnestinglevel,omitempty"`
	NestedGroupExtraction      string   `json:"nestedgroupextraction,omitempty"`
	NextGenAPIResource         string   `json:"_nextgenapiresource,omitempty"`
	PasswdChange               string   `json:"passwdchange,omitempty"`
	SearchFilter               string   `json:"searchfilter,omitempty"`
	SecType                    string   `json:"sectype,omitempty"`
	ServerIP                   string   `json:"serverip,omitempty"`
	ServerPort                 int      `json:"serverport,omitempty"`
	SSONameAttribute           string   `json:"ssonameattribute,omitempty"`
	SubAttributeName           string   `json:"subattributename,omitempty"`
	SvrType                    string   `json:"svrtype,omitempty"`
}

type AAAOTPParameter struct {
	Encryption         string `json:"encryption,omitempty"`
	MaxOTPDevices      int    `json:"maxotpdevices,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
}

type AAAParameter struct {
	AAADLogLevel               string   `json:"aaadloglevel,omitempty"`
	AAADNATIP                  string   `json:"aaadnatip,omitempty"`
	AAASessionLogLevel         string   `json:"aaasessionloglevel,omitempty"`
	APITokenCache              string   `json:"apitokencache,omitempty"`
	Builtin                    []string `json:"builtin,omitempty"`
	ClassicEndpoints           string   `json:"classicendpoints,omitempty"`
	DefaultAuthType            string   `json:"defaultauthtype,omitempty"`
	DefaultCSPHeader           string   `json:"defaultcspheader,omitempty"`
	DynAddr                    string   `json:"dynaddr,omitempty"`
	EnableEnhancedAuthFeedback string   `json:"enableenhancedauthfeedback,omitempty"`
	EnableSessionStickiness    string   `json:"enablesessionstickiness,omitempty"`
	EnableStaticPageCaching    string   `json:"enablestaticpagecaching,omitempty"`
	EnhancedEPA                string   `json:"enhancedepa,omitempty"`
	FailedLoginTimeout         int      `json:"failedlogintimeout,omitempty"`
	Feature                    string   `json:"feature,omitempty"`
	FTMode                     string   `json:"ftmode,omitempty"`
	HTTPOnlyCookie             string   `json:"httponlycookie,omitempty"`
	LoginEncryption            string   `json:"loginencryption,omitempty"`
	MaxAAAUsers                int      `json:"maxaaausers,omitempty"`
	MaxKBQuestions             int      `json:"maxkbquestions,omitempty"`
	MaxLoginAttempts           int      `json:"maxloginattempts,omitempty"`
	MaxSAMLDeflateSize         int      `json:"maxsamldeflatesize,omitempty"`
	NextGenAPIResource         string   `json:"_nextgenapiresource,omitempty"`
	PersistentLoginAttempts    string   `json:"persistentloginattempts,omitempty"`
	PwdExpiryNotificationDays  int      `json:"pwdexpirynotificationdays,omitempty"`
	SameSite                   string   `json:"samesite,omitempty"`
	SecurityInsights           string   `json:"securityinsights,omitempty"`
	TokenIntrospectionInterval int      `json:"tokenintrospectioninterval,omitempty"`
	WAFProtection              []string `json:"wafprotection,omitempty"`
	WebViewEndpoints           string   `json:"webviewendpoints,omitempty"`
}

type AAAPreAuthenticationAction struct {
	Builtin                 []string `json:"builtin,omitempty"`
	Count                   float64  `json:"__count,omitempty"`
	DefaultEPAGroup         string   `json:"defaultepagroup,omitempty"`
	DeleteFiles             string   `json:"deletefiles,omitempty"`
	Feature                 string   `json:"feature,omitempty"`
	KillProcess             string   `json:"killprocess,omitempty"`
	Name                    string   `json:"name,omitempty"`
	NextGenAPIResource      string   `json:"_nextgenapiresource,omitempty"`
	PreAuthenticationAction string   `json:"preauthenticationaction,omitempty"`
}

type AAAPreAuthenticationParameter struct {
	Builtin                 []string `json:"builtin,omitempty"`
	DeleteFiles             string   `json:"deletefiles,omitempty"`
	Feature                 string   `json:"feature,omitempty"`
	KillProcess             string   `json:"killprocess,omitempty"`
	NextGenAPIResource      string   `json:"_nextgenapiresource,omitempty"`
	PreAuthenticationAction string   `json:"preauthenticationaction,omitempty"`
	Rule                    string   `json:"rule,omitempty"`
}

type AAAPreAuthenticationPolicy struct {
	Builtin            []string `json:"builtin,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	Name               string   `json:"name,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	ReqAction          string   `json:"reqaction,omitempty"`
	Rule               string   `json:"rule,omitempty"`
}

type AAAPreAuthenticationPolicyAAAGlobalBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AAAPreAuthenticationPolicyBinding struct {
	AAAPreAuthenticationPolicyAAAGlobalBinding  []interface{} `json:"aaapreauthenticationpolicy_aaaglobal_binding,omitempty"`
	AAAPreAuthenticationPolicyVPNVServerBinding []interface{} `json:"aaapreauthenticationpolicy_vpnvserver_binding,omitempty"`
	Name                                        string        `json:"name,omitempty"`
}

type AAAPreAuthenticationPolicyVPNVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AAARADIUSParams struct {
	Accounting                 string   `json:"accounting,omitempty"`
	Authentication             string   `json:"authentication,omitempty"`
	AuthServRetry              int      `json:"authservretry,omitempty"`
	AuthTimeout                int      `json:"authtimeout,omitempty"`
	Builtin                    []string `json:"builtin,omitempty"`
	CallingStationID           string   `json:"callingstationid,omitempty"`
	DefaultAuthenticationGroup string   `json:"defaultauthenticationgroup,omitempty"`
	Feature                    string   `json:"feature,omitempty"`
	GroupAuthName              string   `json:"groupauthname,omitempty"`
	IPAddress                  string   `json:"ipaddress,omitempty"`
	IPAttributeType            int      `json:"ipattributetype,omitempty"`
	IPVendorID                 int      `json:"ipvendorid,omitempty"`
	MessageAuthenticator       string   `json:"messageauthenticator,omitempty"`
	NextGenAPIResource         string   `json:"_nextgenapiresource,omitempty"`
	PassEncoding               string   `json:"passencoding,omitempty"`
	PwdAttributeType           int      `json:"pwdattributetype,omitempty"`
	PwdVendorID                int      `json:"pwdvendorid,omitempty"`
	RadAttributeType           int      `json:"radattributetype,omitempty"`
	RadGroupSeparator          string   `json:"radgroupseparator,omitempty"`
	RadGroupsPrefix            string   `json:"radgroupsprefix,omitempty"`
	RadKey                     string   `json:"radkey,omitempty"`
	RadNASID                   string   `json:"radnasid,omitempty"`
	RadNASIP                   string   `json:"radnasip,omitempty"`
	RadVendorID                int      `json:"radvendorid,omitempty"`
	ServerIP                   string   `json:"serverip,omitempty"`
	ServerPort                 int      `json:"serverport,omitempty"`
	TunnelEndpointClientIP     string   `json:"tunnelendpointclientip,omitempty"`
}

type AAASession struct {
	All                bool    `json:"all,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	DestIP             string  `json:"destip,omitempty"`
	DestPort           int     `json:"destport,omitempty"`
	GroupName          string  `json:"groupname,omitempty"`
	IIP                string  `json:"iip,omitempty"`
	IntranetIP         string  `json:"intranetip,omitempty"`
	IntranetIP6        string  `json:"intranetip6,omitempty"`
	IPAddress          string  `json:"ipaddress,omitempty"`
	Netmask            string  `json:"netmask,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NodeID             int     `json:"nodeid,omitempty"`
	PEID               int     `json:"peid,omitempty"`
	Port               int     `json:"port,omitempty"`
	PrivateIP          string  `json:"privateip,omitempty"`
	PrivatePort        int     `json:"privateport,omitempty"`
	PublicIP           string  `json:"publicip,omitempty"`
	PublicPort         int     `json:"publicport,omitempty"`
	SessionKey         string  `json:"sessionkey,omitempty"`
	Username           string  `json:"username,omitempty"`
}

type AAASSOProfile struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Password           string  `json:"password,omitempty"`
	Username           string  `json:"username,omitempty"`
}

type AAATACACSParams struct {
	Accounting                 string   `json:"accounting,omitempty"`
	AuditFailedCmds            string   `json:"auditfailedcmds,omitempty"`
	Authorization              string   `json:"authorization,omitempty"`
	AuthTimeout                int      `json:"authtimeout,omitempty"`
	Builtin                    []string `json:"builtin,omitempty"`
	DefaultAuthenticationGroup string   `json:"defaultauthenticationgroup,omitempty"`
	Feature                    string   `json:"feature,omitempty"`
	GroupAttrName              string   `json:"groupattrname,omitempty"`
	NextGenAPIResource         string   `json:"_nextgenapiresource,omitempty"`
	ServerIP                   string   `json:"serverip,omitempty"`
	ServerPort                 int      `json:"serverport,omitempty"`
	TACACSSecret               string   `json:"tacacssecret,omitempty"`
}

type AAAUser struct {
	Count              float64 `json:"__count,omitempty"`
	LoggedIn           bool    `json:"loggedin,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Password           string  `json:"password,omitempty"`
	Username           string  `json:"username,omitempty"`
}

type AAAUserAAAGroupBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupName              string `json:"groupname,omitempty"`
	Username               string `json:"username,omitempty"`
}

type AAAUserAuditNSLogPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
	Username               string `json:"username,omitempty"`
}

type AAAUserAuditSyslogPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
	Username               string `json:"username,omitempty"`
}

type AAAUserAuthorizationPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
	Username               string `json:"username,omitempty"`
}

type AAAUserBinding struct {
	AAAUserAAAGroupBinding                      []interface{} `json:"aaauser_aaagroup_binding,omitempty"`
	AAAUserAuditNSLogPolicyBinding              []interface{} `json:"aaauser_auditnslogpolicy_binding,omitempty"`
	AAAUserAuditSyslogPolicyBinding             []interface{} `json:"aaauser_auditsyslogpolicy_binding,omitempty"`
	AAAUserAuthorizationPolicyBinding           []interface{} `json:"aaauser_authorizationpolicy_binding,omitempty"`
	AAAUserIntranetIP6Binding                   []interface{} `json:"aaauser_intranetip6_binding,omitempty"`
	AAAUserIntranetIPBinding                    []interface{} `json:"aaauser_intranetip_binding,omitempty"`
	AAAUserTMSessionPolicyBinding               []interface{} `json:"aaauser_tmsessionpolicy_binding,omitempty"`
	AAAUserVPNIntranetApplicationBinding        []interface{} `json:"aaauser_vpnintranetapplication_binding,omitempty"`
	AAAUserVPNSecurePrivateAccessProfileBinding []interface{} `json:"aaauser_vpnsecureprivateaccessprofile_binding,omitempty"`
	AAAUserVPNSessionPolicyBinding              []interface{} `json:"aaauser_vpnsessionpolicy_binding,omitempty"`
	AAAUserVPNTrafficPolicyBinding              []interface{} `json:"aaauser_vpntrafficpolicy_binding,omitempty"`
	AAAUserVPNURLBinding                        []interface{} `json:"aaauser_vpnurl_binding,omitempty"`
	AAAUserVPNURLPolicyBinding                  []interface{} `json:"aaauser_vpnurlpolicy_binding,omitempty"`
	Username                                    string        `json:"username,omitempty"`
}

type AAAUserIntranetIP6Binding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	IntranetIP6            string `json:"intranetip6,omitempty"`
	NumAddr                int    `json:"numaddr,omitempty"`
	Username               string `json:"username,omitempty"`
}

type AAAUserIntranetIPBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	IntranetIP             string `json:"intranetip,omitempty"`
	Netmask                string `json:"netmask,omitempty"`
	Username               string `json:"username,omitempty"`
}

type AAAUserTMSessionPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
	Username               string `json:"username,omitempty"`
}

type AAAUserVPNIntranetApplicationBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	IntranetApplication    string `json:"intranetapplication,omitempty"`
	Username               string `json:"username,omitempty"`
}

type AAAUserVPNSecurePrivateAccessProfileBinding struct {
	ActType                    int    `json:"acttype,omitempty"`
	GotoPriorityExpression     string `json:"gotopriorityexpression,omitempty"`
	SecurePrivateAccessProfile string `json:"secureprivateaccessprofile,omitempty"`
	Username                   string `json:"username,omitempty"`
}

type AAAUserVPNSessionPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
	Username               string `json:"username,omitempty"`
}

type AAAUserVPNTrafficPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
	Username               string `json:"username,omitempty"`
}

type AAAUserVPNURLBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	URLName                string `json:"urlname,omitempty"`
	Username               string `json:"username,omitempty"`
}

type AAAUserVPNURLPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
	Username               string `json:"username,omitempty"`
}
