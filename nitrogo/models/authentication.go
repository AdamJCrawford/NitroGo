package models

// authentication configuration structs
type AuthenticationDFAPolicyBinding struct {
	AuthenticationDFAPolicyVPNVServerBinding []interface{} `json:"authenticationdfapolicy_vpnvserver_binding,omitempty"`
	Name                                     string        `json:"name,omitempty"`
}

type AuthenticationWebAuthPolicy struct {
	Action             string  `json:"action,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Rule               string  `json:"rule,omitempty"`
}

type AuthenticationLDAPPolicyBinding struct {
	AuthenticationLDAPPolicyAuthenticationVServerBinding []interface{} `json:"authenticationldappolicy_authenticationvserver_binding,omitempty"`
	AuthenticationLDAPPolicySystemGlobalBinding          []interface{} `json:"authenticationldappolicy_systemglobal_binding,omitempty"`
	AuthenticationLDAPPolicyVPNGlobalBinding             []interface{} `json:"authenticationldappolicy_vpnglobal_binding,omitempty"`
	AuthenticationLDAPPolicyVPNVServerBinding            []interface{} `json:"authenticationldappolicy_vpnvserver_binding,omitempty"`
	Name                                                 string        `json:"name,omitempty"`
}

type AuthenticationWebAuthPolicyVPNVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationPolicySystemGlobalBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	NextFactor             string `json:"nextfactor,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type AuthenticationNegotiatePolicy struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	ReqAction          string  `json:"reqaction,omitempty"`
	Rule               string  `json:"rule,omitempty"`
}

type AuthenticationCertPolicyAuthenticationVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationLDAPPolicy struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	ReqAction          string  `json:"reqaction,omitempty"`
	Rule               string  `json:"rule,omitempty"`
}

type AuthenticationSmartAccessProfile struct {
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Tags               string  `json:"tags,omitempty"`
}

type AuthenticationLocalPolicySystemGlobalBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationPushService struct {
	CertEndpoint          string  `json:"certendpoint,omitempty"`
	ClientID              string  `json:"clientid,omitempty"`
	ClientSecret          string  `json:"clientsecret,omitempty"`
	Count                 float64 `json:"__count,omitempty"`
	CustomerID            string  `json:"customerid,omitempty"`
	HubName               string  `json:"hubname,omitempty"`
	Name                  string  `json:"name,omitempty"`
	Namespace             string  `json:"Namespace,omitempty"`
	NextGenAPIResource    string  `json:"_nextgenapiresource,omitempty"`
	PushCloudServerStatus string  `json:"pushcloudserverstatus,omitempty"`
	PushServiceStatus     string  `json:"pushservicestatus,omitempty"`
	RefreshInterval       int     `json:"refreshinterval,omitempty"`
	ServiceKey            string  `json:"servicekey,omitempty"`
	ServiceKeyName        string  `json:"servicekeyname,omitempty"`
	SigningKey            string  `json:"signingkey,omitempty"`
	SigningKeyName        string  `json:"signingkeyname,omitempty"`
	TrustService          string  `json:"trustservice,omitempty"`
}

type AuthenticationStoreFrontAuthAction struct {
	Count                      float64 `json:"__count,omitempty"`
	DefaultAuthenticationGroup string  `json:"defaultauthenticationgroup,omitempty"`
	Domain                     string  `json:"domain,omitempty"`
	Failure                    int     `json:"failure,omitempty"`
	Name                       string  `json:"name,omitempty"`
	NextGenAPIResource         string  `json:"_nextgenapiresource,omitempty"`
	ServerURL                  string  `json:"serverurl,omitempty"`
	Success                    int     `json:"success,omitempty"`
}

type AuthenticationWebAuthPolicyBinding struct {
	AuthenticationWebAuthPolicyAuthenticationVServerBinding []interface{} `json:"authenticationwebauthpolicy_authenticationvserver_binding,omitempty"`
	AuthenticationWebAuthPolicySystemGlobalBinding          []interface{} `json:"authenticationwebauthpolicy_systemglobal_binding,omitempty"`
	AuthenticationWebAuthPolicyVPNGlobalBinding             []interface{} `json:"authenticationwebauthpolicy_vpnglobal_binding,omitempty"`
	AuthenticationWebAuthPolicyVPNVServerBinding            []interface{} `json:"authenticationwebauthpolicy_vpnvserver_binding,omitempty"`
	Name                                                    string        `json:"name,omitempty"`
}

type AuthenticationLocalPolicyAuthenticationVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationWebAuthPolicyAuthenticationVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationSmartAccessPolicy struct {
	Action             string  `json:"action,omitempty"`
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Hits               int     `json:"hits,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Rule               string  `json:"rule,omitempty"`
}

type AuthenticationSAMLIDPPolicyBinding struct {
	AuthenticationSAMLIDPPolicyAuthenticationVServerBinding []interface{} `json:"authenticationsamlidppolicy_authenticationvserver_binding,omitempty"`
	AuthenticationSAMLIDPPolicyVPNVServerBinding            []interface{} `json:"authenticationsamlidppolicy_vpnvserver_binding,omitempty"`
	Name                                                    string        `json:"name,omitempty"`
}

type AuthenticationOAuthIDPPolicy struct {
	Action                 string  `json:"action,omitempty"`
	Comment                string  `json:"comment,omitempty"`
	Count                  float64 `json:"__count,omitempty"`
	GotoPriorityExpression string  `json:"gotopriorityexpression,omitempty"`
	Hits                   int     `json:"hits,omitempty"`
	LogAction              string  `json:"logaction,omitempty"`
	Name                   string  `json:"name,omitempty"`
	NewName                string  `json:"newname,omitempty"`
	NextGenAPIResource     string  `json:"_nextgenapiresource,omitempty"`
	Rule                   string  `json:"rule,omitempty"`
	UndefAction            string  `json:"undefaction,omitempty"`
}

type AuthenticationVServerBinding struct {
	AuthenticationVServerAuditNSLogPolicyBinding                []interface{} `json:"authenticationvserver_auditnslogpolicy_binding,omitempty"`
	AuthenticationVServerAuditSyslogPolicyBinding               []interface{} `json:"authenticationvserver_auditsyslogpolicy_binding,omitempty"`
	AuthenticationVServerAuthenticationCertPolicyBinding        []interface{} `json:"authenticationvserver_authenticationcertpolicy_binding,omitempty"`
	AuthenticationVServerAuthenticationLDAPPolicyBinding        []interface{} `json:"authenticationvserver_authenticationldappolicy_binding,omitempty"`
	AuthenticationVServerAuthenticationLocalPolicyBinding       []interface{} `json:"authenticationvserver_authenticationlocalpolicy_binding,omitempty"`
	AuthenticationVServerAuthenticationLoginSchemaPolicyBinding []interface{} `json:"authenticationvserver_authenticationloginschemapolicy_binding,omitempty"`
	AuthenticationVServerAuthenticationNegotiatePolicyBinding   []interface{} `json:"authenticationvserver_authenticationnegotiatepolicy_binding,omitempty"`
	AuthenticationVServerAuthenticationOAuthIDPPolicyBinding    []interface{} `json:"authenticationvserver_authenticationoauthidppolicy_binding,omitempty"`
	AuthenticationVServerAuthenticationPolicyBinding            []interface{} `json:"authenticationvserver_authenticationpolicy_binding,omitempty"`
	AuthenticationVServerAuthenticationRADIUSPolicyBinding      []interface{} `json:"authenticationvserver_authenticationradiuspolicy_binding,omitempty"`
	AuthenticationVServerAuthenticationSAMLIDPPolicyBinding     []interface{} `json:"authenticationvserver_authenticationsamlidppolicy_binding,omitempty"`
	AuthenticationVServerAuthenticationSAMLPolicyBinding        []interface{} `json:"authenticationvserver_authenticationsamlpolicy_binding,omitempty"`
	AuthenticationVServerAuthenticationSmartAccessPolicyBinding []interface{} `json:"authenticationvserver_authenticationsmartaccesspolicy_binding,omitempty"`
	AuthenticationVServerAuthenticationTACACSPolicyBinding      []interface{} `json:"authenticationvserver_authenticationtacacspolicy_binding,omitempty"`
	AuthenticationVServerAuthenticationWebAuthPolicyBinding     []interface{} `json:"authenticationvserver_authenticationwebauthpolicy_binding,omitempty"`
	AuthenticationVServerCachePolicyBinding                     []interface{} `json:"authenticationvserver_cachepolicy_binding,omitempty"`
	AuthenticationVServerCSPolicyBinding                        []interface{} `json:"authenticationvserver_cspolicy_binding,omitempty"`
	AuthenticationVServerResponderPolicyBinding                 []interface{} `json:"authenticationvserver_responderpolicy_binding,omitempty"`
	AuthenticationVServerRewritePolicyBinding                   []interface{} `json:"authenticationvserver_rewritepolicy_binding,omitempty"`
	AuthenticationVServerTMSessionPolicyBinding                 []interface{} `json:"authenticationvserver_tmsessionpolicy_binding,omitempty"`
	AuthenticationVServerVPNPortalThemeBinding                  []interface{} `json:"authenticationvserver_vpnportaltheme_binding,omitempty"`
	Name                                                        string        `json:"name,omitempty"`
}

type AuthenticationLocalPolicyBinding struct {
	AuthenticationLocalPolicyAuthenticationVServerBinding []interface{} `json:"authenticationlocalpolicy_authenticationvserver_binding,omitempty"`
	AuthenticationLocalPolicySystemGlobalBinding          []interface{} `json:"authenticationlocalpolicy_systemglobal_binding,omitempty"`
	AuthenticationLocalPolicyVPNGlobalBinding             []interface{} `json:"authenticationlocalpolicy_vpnglobal_binding,omitempty"`
	AuthenticationLocalPolicyVPNVServerBinding            []interface{} `json:"authenticationlocalpolicy_vpnvserver_binding,omitempty"`
	Name                                                  string        `json:"name,omitempty"`
}

type AuthenticationSmartAccessPolicyBinding struct {
	AuthenticationSmartAccessPolicyAuthenticationVServerBinding []interface{} `json:"authenticationsmartaccesspolicy_authenticationvserver_binding,omitempty"`
	Name                                                        string        `json:"name,omitempty"`
}

type AuthenticationVServerAuthenticationWebAuthPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	NextFactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type AuthenticationVServerAuditSyslogPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	NextFactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type AuthenticationRADIUSPolicyAuthenticationVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationSAMLIDPPolicy struct {
	Action                 string  `json:"action,omitempty"`
	Comment                string  `json:"comment,omitempty"`
	Count                  float64 `json:"__count,omitempty"`
	GotoPriorityExpression string  `json:"gotopriorityexpression,omitempty"`
	Hits                   int     `json:"hits,omitempty"`
	LogAction              string  `json:"logaction,omitempty"`
	Name                   string  `json:"name,omitempty"`
	NewName                string  `json:"newname,omitempty"`
	NextGenAPIResource     string  `json:"_nextgenapiresource,omitempty"`
	Rule                   string  `json:"rule,omitempty"`
	UndefAction            string  `json:"undefaction,omitempty"`
}

type AuthenticationRADIUSAction struct {
	Accounting                 string  `json:"accounting,omitempty"`
	Authentication             string  `json:"authentication,omitempty"`
	AuthServRetry              int     `json:"authservretry,omitempty"`
	AuthTimeout                int     `json:"authtimeout,omitempty"`
	CallingStationID           string  `json:"callingstationid,omitempty"`
	Count                      float64 `json:"__count,omitempty"`
	DefaultAuthenticationGroup string  `json:"defaultauthenticationgroup,omitempty"`
	Failure                    int     `json:"failure,omitempty"`
	IPAddress                  string  `json:"ipaddress,omitempty"`
	IPAttributeType            int     `json:"ipattributetype,omitempty"`
	IPVendorID                 int     `json:"ipvendorid,omitempty"`
	MessageAuthenticator       string  `json:"messageauthenticator,omitempty"`
	Name                       string  `json:"name,omitempty"`
	NextGenAPIResource         string  `json:"_nextgenapiresource,omitempty"`
	PassEncoding               string  `json:"passencoding,omitempty"`
	PwdAttributeType           int     `json:"pwdattributetype,omitempty"`
	PwdVendorID                int     `json:"pwdvendorid,omitempty"`
	RadAttributeType           int     `json:"radattributetype,omitempty"`
	RadGroupSeparator          string  `json:"radgroupseparator,omitempty"`
	RadGroupsPrefix            string  `json:"radgroupsprefix,omitempty"`
	RadKey                     string  `json:"radkey,omitempty"`
	RadNASID                   string  `json:"radnasid,omitempty"`
	RadNASIP                   string  `json:"radnasip,omitempty"`
	RadVendorID                int     `json:"radvendorid,omitempty"`
	ServerIP                   string  `json:"serverip,omitempty"`
	ServerName                 string  `json:"servername,omitempty"`
	ServerPort                 int     `json:"serverport,omitempty"`
	Success                    int     `json:"success,omitempty"`
	TargetLBVServer            string  `json:"targetlbvserver,omitempty"`
	Transport                  string  `json:"transport,omitempty"`
	TunnelEndpointClientIP     string  `json:"tunnelendpointclientip,omitempty"`
}

type AuthenticationAzureKeyVault struct {
	Authentication             string  `json:"authentication,omitempty"`
	ClientID                   string  `json:"clientid,omitempty"`
	ClientSecret               string  `json:"clientsecret,omitempty"`
	Count                      float64 `json:"__count,omitempty"`
	DefaultAuthenticationGroup string  `json:"defaultauthenticationgroup,omitempty"`
	Name                       string  `json:"name,omitempty"`
	NextGenAPIResource         string  `json:"_nextgenapiresource,omitempty"`
	PushService                string  `json:"pushservice,omitempty"`
	RefreshInterval            int     `json:"refreshinterval,omitempty"`
	ServiceKeyName             string  `json:"servicekeyname,omitempty"`
	SignatureAlg               string  `json:"signaturealg,omitempty"`
	TenantID                   string  `json:"tenantid,omitempty"`
	TokenEndpoint              string  `json:"tokenendpoint,omitempty"`
	VaultName                  string  `json:"vaultname,omitempty"`
}

type AuthenticationEmailAction struct {
	Content                    string  `json:"content,omitempty"`
	Count                      float64 `json:"__count,omitempty"`
	DefaultAuthenticationGroup string  `json:"defaultauthenticationgroup,omitempty"`
	EmailAddress               string  `json:"emailaddress,omitempty"`
	Name                       string  `json:"name,omitempty"`
	NextGenAPIResource         string  `json:"_nextgenapiresource,omitempty"`
	Password                   string  `json:"password,omitempty"`
	ServerURL                  string  `json:"serverurl,omitempty"`
	Timeout                    int     `json:"timeout,omitempty"`
	TypeField                  string  `json:"type,omitempty"`
	Username                   string  `json:"username,omitempty"`
}

type AuthenticationVServerTMSessionPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	NextFactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type AuthenticationTACACSPolicy struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	ReqAction          string  `json:"reqaction,omitempty"`
	Rule               string  `json:"rule,omitempty"`
}

type AuthenticationLDAPPolicyAuthenticationVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationNegotiatePolicyBinding struct {
	AuthenticationNegotiatePolicyAuthenticationVServerBinding []interface{} `json:"authenticationnegotiatepolicy_authenticationvserver_binding,omitempty"`
	AuthenticationNegotiatePolicyVPNGlobalBinding             []interface{} `json:"authenticationnegotiatepolicy_vpnglobal_binding,omitempty"`
	AuthenticationNegotiatePolicyVPNVServerBinding            []interface{} `json:"authenticationnegotiatepolicy_vpnvserver_binding,omitempty"`
	Name                                                      string        `json:"name,omitempty"`
}

type AuthenticationSmartAccessPolicyAuthenticationVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationLoginSchemaPolicyVPNVServerBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type AuthenticationNegotiatePolicyAuthenticationVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationSAMLAction struct {
	ArtifactResolutionServiceURL   string   `json:"artifactresolutionserviceurl,omitempty"`
	Attribute1                     string   `json:"attribute1,omitempty"`
	Attribute10                    string   `json:"attribute10,omitempty"`
	Attribute11                    string   `json:"attribute11,omitempty"`
	Attribute12                    string   `json:"attribute12,omitempty"`
	Attribute13                    string   `json:"attribute13,omitempty"`
	Attribute14                    string   `json:"attribute14,omitempty"`
	Attribute15                    string   `json:"attribute15,omitempty"`
	Attribute16                    string   `json:"attribute16,omitempty"`
	Attribute2                     string   `json:"attribute2,omitempty"`
	Attribute3                     string   `json:"attribute3,omitempty"`
	Attribute4                     string   `json:"attribute4,omitempty"`
	Attribute5                     string   `json:"attribute5,omitempty"`
	Attribute6                     string   `json:"attribute6,omitempty"`
	Attribute7                     string   `json:"attribute7,omitempty"`
	Attribute8                     string   `json:"attribute8,omitempty"`
	Attribute9                     string   `json:"attribute9,omitempty"`
	AttributeConsumingServiceIndex int      `json:"attributeconsumingserviceindex,omitempty"`
	Attributes                     string   `json:"attributes,omitempty"`
	Audience                       string   `json:"audience,omitempty"`
	AuthnCtxClassRef               []string `json:"authnctxclassref,omitempty"`
	Count                          float64  `json:"__count,omitempty"`
	CustomAuthnCtxClassRef         string   `json:"customauthnctxclassref,omitempty"`
	DefaultAuthenticationGroup     string   `json:"defaultauthenticationgroup,omitempty"`
	DigestMethod                   string   `json:"digestmethod,omitempty"`
	EnforceUsername                string   `json:"enforceusername,omitempty"`
	ForceAuthn                     string   `json:"forceauthn,omitempty"`
	GroupNameField                 string   `json:"groupnamefield,omitempty"`
	LogoutBinding                  string   `json:"logoutbinding,omitempty"`
	LogoutURL                      string   `json:"logouturl,omitempty"`
	MetadataImportStatus           string   `json:"metadataimportstatus,omitempty"`
	MetadataRefreshInterval        int      `json:"metadatarefreshinterval,omitempty"`
	MetadataURL                    string   `json:"metadataurl,omitempty"`
	Name                           string   `json:"name,omitempty"`
	NextGenAPIResource             string   `json:"_nextgenapiresource,omitempty"`
	PreferredBindType              []string `json:"preferredbindtype,omitempty"`
	RelayStateRule                 string   `json:"relaystaterule,omitempty"`
	RequestedAuthnContext          string   `json:"requestedauthncontext,omitempty"`
	SAMLAcsIndex                   int      `json:"samlacsindex,omitempty"`
	SAMLBinding                    string   `json:"samlbinding,omitempty"`
	SAMLIDPCertName                string   `json:"samlidpcertname,omitempty"`
	SAMLIssuerName                 string   `json:"samlissuername,omitempty"`
	SAMLRedirectURL                string   `json:"samlredirecturl,omitempty"`
	SAMLRejectUnsignedAssertion    string   `json:"samlrejectunsignedassertion,omitempty"`
	SAMLSigningCertName            string   `json:"samlsigningcertname,omitempty"`
	SAMLTwoFactor                  string   `json:"samltwofactor,omitempty"`
	SAMLUserField                  string   `json:"samluserfield,omitempty"`
	SendThumbprint                 string   `json:"sendthumbprint,omitempty"`
	SignatureAlg                   string   `json:"signaturealg,omitempty"`
	SkewTime                       int      `json:"skewtime,omitempty"`
	StateChecks                    string   `json:"statechecks,omitempty"`
	StoreSAMLResponse              string   `json:"storesamlresponse,omitempty"`
}

type AuthenticationVServerVPNPortalThemeBinding struct {
	ActType     int    `json:"acttype,omitempty"`
	Name        string `json:"name,omitempty"`
	PortalTheme string `json:"portaltheme,omitempty"`
}

type AuthenticationLocalPolicyVPNGlobalBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationLoginSchemaPolicy struct {
	Action             string   `json:"action,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	Comment            string   `json:"comment,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	LogAction          string   `json:"logaction,omitempty"`
	Name               string   `json:"name,omitempty"`
	NewName            string   `json:"newname,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	Rule               string   `json:"rule,omitempty"`
	UndefAction        string   `json:"undefaction,omitempty"`
	UndefHits          int      `json:"undefhits,omitempty"`
}

type AuthenticationLoginSchemaPolicyBinding struct {
	AuthenticationLoginSchemaPolicyAuthenticationVServerBinding []interface{} `json:"authenticationloginschemapolicy_authenticationvserver_binding,omitempty"`
	AuthenticationLoginSchemaPolicyVPNVServerBinding            []interface{} `json:"authenticationloginschemapolicy_vpnvserver_binding,omitempty"`
	Name                                                        string        `json:"name,omitempty"`
}

type AuthenticationProtectedUserAction struct {
	Count              float64 `json:"__count,omitempty"`
	MaxConcurrentUsers int     `json:"maxconcurrentusers,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	RealmStr           string  `json:"realmstr,omitempty"`
}

type AuthenticationPolicyAuthenticationVServerBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	NextFactor             string `json:"nextfactor,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type AuthenticationSAMLPolicyVPNGlobalBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationTACACSAction struct {
	Accounting                 string  `json:"accounting,omitempty"`
	Attribute1                 string  `json:"attribute1,omitempty"`
	Attribute10                string  `json:"attribute10,omitempty"`
	Attribute11                string  `json:"attribute11,omitempty"`
	Attribute12                string  `json:"attribute12,omitempty"`
	Attribute13                string  `json:"attribute13,omitempty"`
	Attribute14                string  `json:"attribute14,omitempty"`
	Attribute15                string  `json:"attribute15,omitempty"`
	Attribute16                string  `json:"attribute16,omitempty"`
	Attribute2                 string  `json:"attribute2,omitempty"`
	Attribute3                 string  `json:"attribute3,omitempty"`
	Attribute4                 string  `json:"attribute4,omitempty"`
	Attribute5                 string  `json:"attribute5,omitempty"`
	Attribute6                 string  `json:"attribute6,omitempty"`
	Attribute7                 string  `json:"attribute7,omitempty"`
	Attribute8                 string  `json:"attribute8,omitempty"`
	Attribute9                 string  `json:"attribute9,omitempty"`
	Attributes                 string  `json:"attributes,omitempty"`
	AuditFailedCmds            string  `json:"auditfailedcmds,omitempty"`
	Authorization              string  `json:"authorization,omitempty"`
	AuthTimeout                int     `json:"authtimeout,omitempty"`
	Count                      float64 `json:"__count,omitempty"`
	DefaultAuthenticationGroup string  `json:"defaultauthenticationgroup,omitempty"`
	Failure                    int     `json:"failure,omitempty"`
	GroupAttrName              string  `json:"groupattrname,omitempty"`
	Name                       string  `json:"name,omitempty"`
	NextGenAPIResource         string  `json:"_nextgenapiresource,omitempty"`
	ServerIP                   string  `json:"serverip,omitempty"`
	ServerPort                 int     `json:"serverport,omitempty"`
	Success                    int     `json:"success,omitempty"`
	TACACSSecret               string  `json:"tacacssecret,omitempty"`
}

type AuthenticationCertPolicyVPNVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationVServerAuthenticationLoginSchemaPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	NextFactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type AuthenticationSAMLIDPPolicyAuthenticationVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationNegotiatePolicyVPNGlobalBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationSAMLPolicyBinding struct {
	AuthenticationSAMLPolicyAuthenticationVServerBinding []interface{} `json:"authenticationsamlpolicy_authenticationvserver_binding,omitempty"`
	AuthenticationSAMLPolicyVPNGlobalBinding             []interface{} `json:"authenticationsamlpolicy_vpnglobal_binding,omitempty"`
	AuthenticationSAMLPolicyVPNVServerBinding            []interface{} `json:"authenticationsamlpolicy_vpnvserver_binding,omitempty"`
	Name                                                 string        `json:"name,omitempty"`
}

type AuthenticationVServerAuthenticationLocalPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	NextFactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type AuthenticationRADIUSPolicySystemGlobalBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationVServer struct {
	AppFlowLog           string  `json:"appflowlog,omitempty"`
	Authentication       string  `json:"authentication,omitempty"`
	AuthenticationDomain string  `json:"authenticationdomain,omitempty"`
	BackupVServer        string  `json:"backupvserver,omitempty"`
	BindPoint            string  `json:"bindpoint,omitempty"`
	CacheType            string  `json:"cachetype,omitempty"`
	CacheVServer         string  `json:"cachevserver,omitempty"`
	CertKeyNames         string  `json:"certkeynames,omitempty"`
	CltTimeout           int     `json:"clttimeout,omitempty"`
	Comment              string  `json:"comment,omitempty"`
	Count                float64 `json:"__count,omitempty"`
	CurAAAUsers          int     `json:"curaaausers,omitempty"`
	CurState             string  `json:"curstate,omitempty"`
	DisablePrimaryOnDown string  `json:"disableprimaryondown,omitempty"`
	DownStateFlush       string  `json:"downstateflush,omitempty"`
	FailedLoginTimeout   int     `json:"failedlogintimeout,omitempty"`
	GroupExtraction      bool    `json:"groupextraction,omitempty"`
	HTTPProfileName      string  `json:"httpprofilename,omitempty"`
	IP                   string  `json:"ip,omitempty"`
	IPv46                string  `json:"ipv46,omitempty"`
	ListenPolicy         string  `json:"listenpolicy,omitempty"`
	ListenPriority       int     `json:"listenpriority,omitempty"`
	MaxLoginAttempts     int     `json:"maxloginattempts,omitempty"`
	Name                 string  `json:"name,omitempty"`
	NewName              string  `json:"newname,omitempty"`
	NextGenAPIResource   string  `json:"_nextgenapiresource,omitempty"`
	NGName               string  `json:"ngname,omitempty"`
	Policy               string  `json:"policy,omitempty"`
	Port                 int     `json:"port,omitempty"`
	Precedence           string  `json:"precedence,omitempty"`
	Priority             int     `json:"priority,omitempty"`
	Range                int     `json:"range,omitempty"`
	Redirect             string  `json:"redirect,omitempty"`
	RedirectURL          string  `json:"redirecturl,omitempty"`
	SameSite             string  `json:"samesite,omitempty"`
	Secondary            bool    `json:"secondary,omitempty"`
	ServiceName          string  `json:"servicename,omitempty"`
	ServiceType          string  `json:"servicetype,omitempty"`
	SOMethod             string  `json:"somethod,omitempty"`
	SOPersistence        string  `json:"sopersistence,omitempty"`
	SOPersistenceTimeout int     `json:"sopersistencetimeout,omitempty"`
	SOThreshold          int     `json:"sothreshold,omitempty"`
	State                string  `json:"state,omitempty"`
	Status               int     `json:"status,omitempty"`
	TCPProfileName       string  `json:"tcpprofilename,omitempty"`
	TD                   int     `json:"td,omitempty"`
	TypeField            string  `json:"type,omitempty"`
	Value                string  `json:"value,omitempty"`
	VSType               int     `json:"vstype,omitempty"`
	Weight               int     `json:"weight,omitempty"`
}

type AuthenticationNoAuthAction struct {
	Builtin                    []string `json:"builtin,omitempty"`
	Count                      float64  `json:"__count,omitempty"`
	DefaultAuthenticationGroup string   `json:"defaultauthenticationgroup,omitempty"`
	Feature                    string   `json:"feature,omitempty"`
	Name                       string   `json:"name,omitempty"`
	NextGenAPIResource         string   `json:"_nextgenapiresource,omitempty"`
}

type AuthenticationNegotiatePolicyVPNVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationSAMLPolicyVPNVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationAuthNProfile struct {
	AuthenticationDomain string  `json:"authenticationdomain,omitempty"`
	AuthenticationHost   string  `json:"authenticationhost,omitempty"`
	AuthenticationLevel  int     `json:"authenticationlevel,omitempty"`
	AuthNVSName          string  `json:"authnvsname,omitempty"`
	Count                float64 `json:"__count,omitempty"`
	Name                 string  `json:"name,omitempty"`
	NextGenAPIResource   string  `json:"_nextgenapiresource,omitempty"`
}

type AuthenticationCaptchaAction struct {
	Count                      float64 `json:"__count,omitempty"`
	DefaultAuthenticationGroup string  `json:"defaultauthenticationgroup,omitempty"`
	Name                       string  `json:"name,omitempty"`
	NextGenAPIResource         string  `json:"_nextgenapiresource,omitempty"`
	ScoreThreshold             int     `json:"scorethreshold,omitempty"`
	SecretKey                  string  `json:"secretkey,omitempty"`
	ServerURL                  string  `json:"serverurl,omitempty"`
	SiteKey                    string  `json:"sitekey,omitempty"`
}

type AuthenticationRADIUSPolicyBinding struct {
	AuthenticationRADIUSPolicyAuthenticationVServerBinding []interface{} `json:"authenticationradiuspolicy_authenticationvserver_binding,omitempty"`
	AuthenticationRADIUSPolicySystemGlobalBinding          []interface{} `json:"authenticationradiuspolicy_systemglobal_binding,omitempty"`
	AuthenticationRADIUSPolicyVPNGlobalBinding             []interface{} `json:"authenticationradiuspolicy_vpnglobal_binding,omitempty"`
	AuthenticationRADIUSPolicyVPNVServerBinding            []interface{} `json:"authenticationradiuspolicy_vpnvserver_binding,omitempty"`
	Name                                                   string        `json:"name,omitempty"`
}

type AuthenticationTACACSPolicyBinding struct {
	AuthenticationTACACSPolicyAuthenticationVServerBinding []interface{} `json:"authenticationtacacspolicy_authenticationvserver_binding,omitempty"`
	AuthenticationTACACSPolicySystemGlobalBinding          []interface{} `json:"authenticationtacacspolicy_systemglobal_binding,omitempty"`
	AuthenticationTACACSPolicyVPNGlobalBinding             []interface{} `json:"authenticationtacacspolicy_vpnglobal_binding,omitempty"`
	AuthenticationTACACSPolicyVPNVServerBinding            []interface{} `json:"authenticationtacacspolicy_vpnvserver_binding,omitempty"`
	Name                                                   string        `json:"name,omitempty"`
}

type AuthenticationSAMLPolicy struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	ReqAction          string  `json:"reqaction,omitempty"`
	Rule               string  `json:"rule,omitempty"`
}

type AuthenticationLDAPPolicyVPNVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationWebAuthPolicyVPNGlobalBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationLocalPolicy struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	ReqAction          string  `json:"reqaction,omitempty"`
	Rule               string  `json:"rule,omitempty"`
}

type AuthenticationVServerAuthenticationRADIUSPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	NextFactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type AuthenticationLoginSchemaPolicyAuthenticationVServerBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type AuthenticationLocalPolicyVPNVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationVServerAuthenticationSAMLPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	NextFactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type AuthenticationSAMLIDPPolicyVPNVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationOAuthIDPPolicyAuthenticationVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationCitrixAuthAction struct {
	Authentication     string  `json:"authentication,omitempty"`
	AuthenticationType string  `json:"authenticationtype,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
}

type AuthenticationVServerAuthenticationCertPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	NextFactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type AuthenticationPolicyLabelBinding struct {
	AuthenticationPolicyLabelAuthenticationPolicyBinding []interface{} `json:"authenticationpolicylabel_authenticationpolicy_binding,omitempty"`
	LabelName                                            string        `json:"labelname,omitempty"`
}

type AuthenticationDFAPolicy struct {
	Action             string  `json:"action,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Rule               string  `json:"rule,omitempty"`
}

type AuthenticationTACACSPolicyVPNVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationRADIUSPolicyVPNGlobalBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationVServerAuthenticationLDAPPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	NextFactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type AuthenticationOAuthIDPProfile struct {
	Attributes                 string  `json:"attributes,omitempty"`
	Audience                   string  `json:"audience,omitempty"`
	ClientID                   string  `json:"clientid,omitempty"`
	ClientSecret               string  `json:"clientsecret,omitempty"`
	ConfigService              string  `json:"configservice,omitempty"`
	Count                      float64 `json:"__count,omitempty"`
	DefaultAuthenticationGroup string  `json:"defaultauthenticationgroup,omitempty"`
	EncryptToken               string  `json:"encrypttoken,omitempty"`
	Issuer                     string  `json:"issuer,omitempty"`
	Name                       string  `json:"name,omitempty"`
	NextGenAPIResource         string  `json:"_nextgenapiresource,omitempty"`
	OAuthStatus                string  `json:"oauthstatus,omitempty"`
	RedirectURL                string  `json:"redirecturl,omitempty"`
	RefreshInterval            int     `json:"refreshinterval,omitempty"`
	RelyingPartyMetadataURL    string  `json:"relyingpartymetadataurl,omitempty"`
	SendPassword               string  `json:"sendpassword,omitempty"`
	SignatureAlg               string  `json:"signaturealg,omitempty"`
	SignatureService           string  `json:"signatureservice,omitempty"`
	SkewTime                   int     `json:"skewtime,omitempty"`
}

type AuthenticationTACACSPolicySystemGlobalBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationEPAAction struct {
	Count              float64 `json:"__count,omitempty"`
	CSecExpr           string  `json:"csecexpr,omitempty"`
	DefaultEPAGroup    string  `json:"defaultepagroup,omitempty"`
	DeleteFiles        string  `json:"deletefiles,omitempty"`
	DevicePosture      string  `json:"deviceposture,omitempty"`
	KillProcess        string  `json:"killprocess,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	QuarantineGroup    string  `json:"quarantinegroup,omitempty"`
}

type AuthenticationPolicyBinding struct {
	AuthenticationPolicyAuthenticationPolicyLabelBinding []interface{} `json:"authenticationpolicy_authenticationpolicylabel_binding,omitempty"`
	AuthenticationPolicyAuthenticationVServerBinding     []interface{} `json:"authenticationpolicy_authenticationvserver_binding,omitempty"`
	AuthenticationPolicySystemGlobalBinding              []interface{} `json:"authenticationpolicy_systemglobal_binding,omitempty"`
	Name                                                 string        `json:"name,omitempty"`
}

type AuthenticationCertPolicyVPNGlobalBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationDFAPolicyVPNVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationWebAuthAction struct {
	Attribute1                 string  `json:"attribute1,omitempty"`
	Attribute10                string  `json:"attribute10,omitempty"`
	Attribute11                string  `json:"attribute11,omitempty"`
	Attribute12                string  `json:"attribute12,omitempty"`
	Attribute13                string  `json:"attribute13,omitempty"`
	Attribute14                string  `json:"attribute14,omitempty"`
	Attribute15                string  `json:"attribute15,omitempty"`
	Attribute16                string  `json:"attribute16,omitempty"`
	Attribute2                 string  `json:"attribute2,omitempty"`
	Attribute3                 string  `json:"attribute3,omitempty"`
	Attribute4                 string  `json:"attribute4,omitempty"`
	Attribute5                 string  `json:"attribute5,omitempty"`
	Attribute6                 string  `json:"attribute6,omitempty"`
	Attribute7                 string  `json:"attribute7,omitempty"`
	Attribute8                 string  `json:"attribute8,omitempty"`
	Attribute9                 string  `json:"attribute9,omitempty"`
	Count                      float64 `json:"__count,omitempty"`
	DefaultAuthenticationGroup string  `json:"defaultauthenticationgroup,omitempty"`
	FullReqExpr                string  `json:"fullreqexpr,omitempty"`
	Name                       string  `json:"name,omitempty"`
	NextGenAPIResource         string  `json:"_nextgenapiresource,omitempty"`
	Scheme                     string  `json:"scheme,omitempty"`
	ServerIP                   string  `json:"serverip,omitempty"`
	ServerPort                 int     `json:"serverport,omitempty"`
	SuccessRule                string  `json:"successrule,omitempty"`
}

type AuthenticationPolicyLabel struct {
	Comment                string  `json:"comment,omitempty"`
	Count                  float64 `json:"__count,omitempty"`
	Description            string  `json:"description,omitempty"`
	FlowType               int     `json:"flowtype,omitempty"`
	GotoPriorityExpression string  `json:"gotopriorityexpression,omitempty"`
	Hits                   int     `json:"hits,omitempty"`
	LabelName              string  `json:"labelname,omitempty"`
	LoginSchema            string  `json:"loginschema,omitempty"`
	NewName                string  `json:"newname,omitempty"`
	NextGenAPIResource     string  `json:"_nextgenapiresource,omitempty"`
	NumPol                 int     `json:"numpol,omitempty"`
	PolicyName             string  `json:"policyname,omitempty"`
	Priority               int     `json:"priority,omitempty"`
	TypeField              string  `json:"type,omitempty"`
}

type AuthenticationRADIUSPolicy struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	ReqAction          string  `json:"reqaction,omitempty"`
	Rule               string  `json:"rule,omitempty"`
}

type AuthenticationVServerAuthenticationOAuthIDPPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	NextFactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type AuthenticationVServerRewritePolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	NextFactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type AuthenticationDFAAction struct {
	ClientID                   string  `json:"clientid,omitempty"`
	Count                      float64 `json:"__count,omitempty"`
	DefaultAuthenticationGroup string  `json:"defaultauthenticationgroup,omitempty"`
	Failure                    int     `json:"failure,omitempty"`
	Name                       string  `json:"name,omitempty"`
	NextGenAPIResource         string  `json:"_nextgenapiresource,omitempty"`
	Passphrase                 string  `json:"passphrase,omitempty"`
	ServerURL                  string  `json:"serverurl,omitempty"`
	Success                    int     `json:"success,omitempty"`
}

type AuthenticationOAuthIDPPolicyBinding struct {
	AuthenticationOAuthIDPPolicyAuthenticationVServerBinding []interface{} `json:"authenticationoauthidppolicy_authenticationvserver_binding,omitempty"`
	AuthenticationOAuthIDPPolicyVPNVServerBinding            []interface{} `json:"authenticationoauthidppolicy_vpnvserver_binding,omitempty"`
	Name                                                     string        `json:"name,omitempty"`
}

type AuthenticationVServerCachePolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	NextFactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type AuthenticationVServerAuthenticationNegotiatePolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	NextFactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type AuthenticationTACACSPolicyAuthenticationVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationTACACSPolicyVPNGlobalBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationVServerAuthenticationSAMLIDPPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	NextFactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type AuthenticationVServerAuthenticationTACACSPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	NextFactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type AuthenticationLDAPAction struct {
	AlternateEmailAttr         string  `json:"alternateemailattr,omitempty"`
	Attribute1                 string  `json:"attribute1,omitempty"`
	Attribute10                string  `json:"attribute10,omitempty"`
	Attribute11                string  `json:"attribute11,omitempty"`
	Attribute12                string  `json:"attribute12,omitempty"`
	Attribute13                string  `json:"attribute13,omitempty"`
	Attribute14                string  `json:"attribute14,omitempty"`
	Attribute15                string  `json:"attribute15,omitempty"`
	Attribute16                string  `json:"attribute16,omitempty"`
	Attribute2                 string  `json:"attribute2,omitempty"`
	Attribute3                 string  `json:"attribute3,omitempty"`
	Attribute4                 string  `json:"attribute4,omitempty"`
	Attribute5                 string  `json:"attribute5,omitempty"`
	Attribute6                 string  `json:"attribute6,omitempty"`
	Attribute7                 string  `json:"attribute7,omitempty"`
	Attribute8                 string  `json:"attribute8,omitempty"`
	Attribute9                 string  `json:"attribute9,omitempty"`
	Attributes                 string  `json:"attributes,omitempty"`
	Authentication             string  `json:"authentication,omitempty"`
	AuthTimeout                int     `json:"authtimeout,omitempty"`
	CloudAttributes            string  `json:"cloudattributes,omitempty"`
	Count                      float64 `json:"__count,omitempty"`
	DefaultAuthenticationGroup string  `json:"defaultauthenticationgroup,omitempty"`
	Email                      string  `json:"email,omitempty"`
	Failure                    int     `json:"failure,omitempty"`
	FollowReferrals            string  `json:"followreferrals,omitempty"`
	GroupAttrName              string  `json:"groupattrname,omitempty"`
	GroupNameIdentifier        string  `json:"groupnameidentifier,omitempty"`
	GroupSearchAttribute       string  `json:"groupsearchattribute,omitempty"`
	GroupSearchFilter          string  `json:"groupsearchfilter,omitempty"`
	GroupSearchSubAttribute    string  `json:"groupsearchsubattribute,omitempty"`
	KBAttribute                string  `json:"kbattribute,omitempty"`
	LDAPBase                   string  `json:"ldapbase,omitempty"`
	LDAPBindDN                 string  `json:"ldapbinddn,omitempty"`
	LDAPBindDNPassword         string  `json:"ldapbinddnpassword,omitempty"`
	LDAPHostname               string  `json:"ldaphostname,omitempty"`
	LDAPLoginName              string  `json:"ldaploginname,omitempty"`
	MaxLDAPReferrals           int     `json:"maxldapreferrals,omitempty"`
	MaxNestingLevel            int     `json:"maxnestinglevel,omitempty"`
	MSSRVRecordLocation        string  `json:"mssrvrecordlocation,omitempty"`
	Name                       string  `json:"name,omitempty"`
	NestedGroupExtraction      string  `json:"nestedgroupextraction,omitempty"`
	NextGenAPIResource         string  `json:"_nextgenapiresource,omitempty"`
	OTPSecret                  string  `json:"otpsecret,omitempty"`
	PasswdChange               string  `json:"passwdchange,omitempty"`
	PushService                string  `json:"pushservice,omitempty"`
	ReferralDNSLookup          string  `json:"referraldnslookup,omitempty"`
	RequireUser                string  `json:"requireuser,omitempty"`
	SearchFilter               string  `json:"searchfilter,omitempty"`
	SecType                    string  `json:"sectype,omitempty"`
	ServerIP                   string  `json:"serverip,omitempty"`
	ServerName                 string  `json:"servername,omitempty"`
	ServerPort                 int     `json:"serverport,omitempty"`
	SSHPublicKey               string  `json:"sshpublickey,omitempty"`
	SSONameAttribute           string  `json:"ssonameattribute,omitempty"`
	SubAttributeName           string  `json:"subattributename,omitempty"`
	Success                    int     `json:"success,omitempty"`
	SvrType                    string  `json:"svrtype,omitempty"`
	ValidateServerCert         string  `json:"validateservercert,omitempty"`
}

type AuthenticationVServerAuthenticationPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	NextFactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type AuthenticationCertPolicyBinding struct {
	AuthenticationCertPolicyAuthenticationVServerBinding []interface{} `json:"authenticationcertpolicy_authenticationvserver_binding,omitempty"`
	AuthenticationCertPolicyVPNGlobalBinding             []interface{} `json:"authenticationcertpolicy_vpnglobal_binding,omitempty"`
	AuthenticationCertPolicyVPNVServerBinding            []interface{} `json:"authenticationcertpolicy_vpnvserver_binding,omitempty"`
	Name                                                 string        `json:"name,omitempty"`
}

type AuthenticationWebAuthPolicySystemGlobalBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationVServerAuditNSLogPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	NextFactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type AuthenticationSAMLPolicyAuthenticationVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationOAuthAction struct {
	AllowedAlgorithms          []string `json:"allowedalgorithms,omitempty"`
	Attribute1                 string   `json:"attribute1,omitempty"`
	Attribute10                string   `json:"attribute10,omitempty"`
	Attribute11                string   `json:"attribute11,omitempty"`
	Attribute12                string   `json:"attribute12,omitempty"`
	Attribute13                string   `json:"attribute13,omitempty"`
	Attribute14                string   `json:"attribute14,omitempty"`
	Attribute15                string   `json:"attribute15,omitempty"`
	Attribute16                string   `json:"attribute16,omitempty"`
	Attribute2                 string   `json:"attribute2,omitempty"`
	Attribute3                 string   `json:"attribute3,omitempty"`
	Attribute4                 string   `json:"attribute4,omitempty"`
	Attribute5                 string   `json:"attribute5,omitempty"`
	Attribute6                 string   `json:"attribute6,omitempty"`
	Attribute7                 string   `json:"attribute7,omitempty"`
	Attribute8                 string   `json:"attribute8,omitempty"`
	Attribute9                 string   `json:"attribute9,omitempty"`
	Attributes                 string   `json:"attributes,omitempty"`
	Audience                   string   `json:"audience,omitempty"`
	Authentication             string   `json:"authentication,omitempty"`
	AuthorizationEndpoint      string   `json:"authorizationendpoint,omitempty"`
	CertEndpoint               string   `json:"certendpoint,omitempty"`
	CertFilePath               string   `json:"certfilepath,omitempty"`
	ClientID                   string   `json:"clientid,omitempty"`
	ClientSecret               string   `json:"clientsecret,omitempty"`
	Count                      float64  `json:"__count,omitempty"`
	DefaultAuthenticationGroup string   `json:"defaultauthenticationgroup,omitempty"`
	GrantType                  string   `json:"granttype,omitempty"`
	GraphEndpoint              string   `json:"graphendpoint,omitempty"`
	IDTokenDecryptEndpoint     string   `json:"idtokendecryptendpoint,omitempty"`
	IntrospectURL              string   `json:"introspecturl,omitempty"`
	IntuneDeviceIDExpression   string   `json:"intunedeviceidexpression,omitempty"`
	Issuer                     string   `json:"issuer,omitempty"`
	MetadataURL                string   `json:"metadataurl,omitempty"`
	Name                       string   `json:"name,omitempty"`
	NextGenAPIResource         string   `json:"_nextgenapiresource,omitempty"`
	OAuthMiscFlags             []string `json:"oauthmiscflags,omitempty"`
	OAuthStatus                string   `json:"oauthstatus,omitempty"`
	OAuthType                  string   `json:"oauthtype,omitempty"`
	PKCE                       string   `json:"pkce,omitempty"`
	RefreshInterval            int      `json:"refreshinterval,omitempty"`
	RequestAttribute           string   `json:"requestattribute,omitempty"`
	ResourceURI                string   `json:"resourceuri,omitempty"`
	SkewTime                   int      `json:"skewtime,omitempty"`
	TenantID                   string   `json:"tenantid,omitempty"`
	TokenEndpoint              string   `json:"tokenendpoint,omitempty"`
	TokenEndpointAuthMethod    string   `json:"tokenendpointauthmethod,omitempty"`
	UserInfoURL                string   `json:"userinfourl,omitempty"`
	UsernameField              string   `json:"usernamefield,omitempty"`
}

type AuthenticationNegotiateAction struct {
	Count                      float64 `json:"__count,omitempty"`
	DefaultAuthenticationGroup string  `json:"defaultauthenticationgroup,omitempty"`
	Domain                     string  `json:"domain,omitempty"`
	DomainUser                 string  `json:"domainuser,omitempty"`
	DomainUserPasswd           string  `json:"domainuserpasswd,omitempty"`
	KCDSPN                     string  `json:"kcdspn,omitempty"`
	KeyTab                     string  `json:"keytab,omitempty"`
	Name                       string  `json:"name,omitempty"`
	NextGenAPIResource         string  `json:"_nextgenapiresource,omitempty"`
	NTLMPath                   string  `json:"ntlmpath,omitempty"`
	OU                         string  `json:"ou,omitempty"`
}

type AuthenticationLDAPPolicySystemGlobalBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationLoginSchema struct {
	AuthenticationSchema    string   `json:"authenticationschema,omitempty"`
	AuthenticationStrength  int      `json:"authenticationstrength,omitempty"`
	Builtin                 []string `json:"builtin,omitempty"`
	Count                   float64  `json:"__count,omitempty"`
	Feature                 string   `json:"feature,omitempty"`
	Name                    string   `json:"name,omitempty"`
	NextGenAPIResource      string   `json:"_nextgenapiresource,omitempty"`
	PasswdExpression        string   `json:"passwdexpression,omitempty"`
	PasswordCredentialIndex int      `json:"passwordcredentialindex,omitempty"`
	SSOCredentials          string   `json:"ssocredentials,omitempty"`
	UserCredentialIndex     int      `json:"usercredentialindex,omitempty"`
	UserExpression          string   `json:"userexpression,omitempty"`
}

type AuthenticationPolicyLabelAuthenticationPolicyBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	NextFactor             string `json:"nextfactor,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type AuthenticationRADIUSPolicyVPNVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationCertAction struct {
	Count                      float64 `json:"__count,omitempty"`
	DefaultAuthenticationGroup string  `json:"defaultauthenticationgroup,omitempty"`
	GroupNameField             string  `json:"groupnamefield,omitempty"`
	Name                       string  `json:"name,omitempty"`
	NextGenAPIResource         string  `json:"_nextgenapiresource,omitempty"`
	TwoFactor                  string  `json:"twofactor,omitempty"`
	UsernameField              string  `json:"usernamefield,omitempty"`
}

type AuthenticationADFSProxyProfile struct {
	ADFSTrustStatus    string  `json:"adfstruststatus,omitempty"`
	CertKeyName        string  `json:"certkeyname,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Password           string  `json:"password,omitempty"`
	ServerURL          string  `json:"serverurl,omitempty"`
	Username           string  `json:"username,omitempty"`
}

type AuthenticationVServerResponderPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	NextFactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type AuthenticationVServerCSPolicyBinding struct {
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	NextFactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type AuthenticationPolicyAuthenticationPolicyLabelBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	NextFactor             string `json:"nextfactor,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type AuthenticationOAuthIDPPolicyVPNVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationPolicy struct {
	Action             string  `json:"action,omitempty"`
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Description        string  `json:"description,omitempty"`
	Hits               int     `json:"hits,omitempty"`
	LogAction          string  `json:"logaction,omitempty"`
	Name               string  `json:"name,omitempty"`
	NewName            string  `json:"newname,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	PolicySubType      string  `json:"policysubtype,omitempty"`
	Rule               string  `json:"rule,omitempty"`
	UndefAction        string  `json:"undefaction,omitempty"`
}

type AuthenticationSAMLIDPProfile struct {
	ACSURLRule                  string  `json:"acsurlrule,omitempty"`
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
	DefaultAuthenticationGroup  string  `json:"defaultauthenticationgroup,omitempty"`
	DigestMethod                string  `json:"digestmethod,omitempty"`
	EncryptAssertion            string  `json:"encryptassertion,omitempty"`
	EncryptionAlgorithm         string  `json:"encryptionalgorithm,omitempty"`
	KeyTransportAlg             string  `json:"keytransportalg,omitempty"`
	LogoutBinding               string  `json:"logoutbinding,omitempty"`
	MetadataImportStatus        string  `json:"metadataimportstatus,omitempty"`
	MetadataRefreshInterval     int     `json:"metadatarefreshinterval,omitempty"`
	MetadataURL                 string  `json:"metadataurl,omitempty"`
	Name                        string  `json:"name,omitempty"`
	NameIDExpr                  string  `json:"nameidexpr,omitempty"`
	NameIDFormat                string  `json:"nameidformat,omitempty"`
	NextGenAPIResource          string  `json:"_nextgenapiresource,omitempty"`
	RejectUnsignedRequests      string  `json:"rejectunsignedrequests,omitempty"`
	SAMLBinding                 string  `json:"samlbinding,omitempty"`
	SAMLIDPCertName             string  `json:"samlidpcertname,omitempty"`
	SAMLIssuerName              string  `json:"samlissuername,omitempty"`
	SAMLSigningCertVersion      string  `json:"samlsigningcertversion,omitempty"`
	SAMLSPCertName              string  `json:"samlspcertname,omitempty"`
	SAMLSPCertVersion           string  `json:"samlspcertversion,omitempty"`
	SendPassword                string  `json:"sendpassword,omitempty"`
	ServiceProviderID           string  `json:"serviceproviderid,omitempty"`
	SignAssertion               string  `json:"signassertion,omitempty"`
	SignatureAlg                string  `json:"signaturealg,omitempty"`
	SignatureService            string  `json:"signatureservice,omitempty"`
	SkewTime                    int     `json:"skewtime,omitempty"`
	SPLogoutURL                 string  `json:"splogouturl,omitempty"`
}

type AuthenticationLDAPPolicyVPNGlobalBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationVServerAuthenticationSmartAccessPolicyBinding struct {
	ActType                int    `json:"acttype,omitempty"`
	BindPoint              string `json:"bindpoint,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	GroupExtraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	NextFactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type AuthenticationCertPolicy struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	ReqAction          string  `json:"reqaction,omitempty"`
	Rule               string  `json:"rule,omitempty"`
}
