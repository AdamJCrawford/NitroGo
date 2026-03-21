package models

// tm configuration structs
type TMFormSSOAction struct {
	ActionURL          string  `json:"actionurl,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	NameValuePair      string  `json:"namevaluepair,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NVType             string  `json:"nvtype,omitempty"`
	PasswdField        string  `json:"passwdfield,omitempty"`
	ResponseSize       int     `json:"responsesize,omitempty"`
	SSOSuccessRule     string  `json:"ssosuccessrule,omitempty"`
	SubmitMethod       string  `json:"submitmethod,omitempty"`
	UserField          string  `json:"userfield,omitempty"`
}

type TMTrafficPolicyLBVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type TMTrafficPolicy struct {
	Action             string  `json:"action,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Hits               int     `json:"hits,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Rule               string  `json:"rule,omitempty"`
}

type TMGlobalTMTrafficPolicyBinding struct {
	BindPolicyType         int    `json:"bindpolicytype,omitempty"`
	GlobalBindType         string `json:"globalbindtype,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
}

type TMGlobalBinding struct {
	TMGlobalAuditNSLogPolicyBinding  []any `json:"tmglobal_auditnslogpolicy_binding,omitempty"`
	TMGlobalAuditSyslogPolicyBinding []any `json:"tmglobal_auditsyslogpolicy_binding,omitempty"`
	TMGlobalTMSessionPolicyBinding   []any `json:"tmglobal_tmsessionpolicy_binding,omitempty"`
	TMGlobalTMTrafficPolicyBinding   []any `json:"tmglobal_tmtrafficpolicy_binding,omitempty"`
}

type TMTrafficPolicyBinding struct {
	Name                            string `json:"name,omitempty"`
	TMTrafficPolicyCSVServerBinding []any  `json:"tmtrafficpolicy_csvserver_binding,omitempty"`
	TMTrafficPolicyLBVServerBinding []any  `json:"tmtrafficpolicy_lbvserver_binding,omitempty"`
	TMTrafficPolicyTMGlobalBinding  []any  `json:"tmtrafficpolicy_tmglobal_binding,omitempty"`
}

type TMSessionPolicy struct {
	Action                 string   `json:"action,omitempty"`
	Builtin                []string `json:"builtin,omitempty"`
	Count                  float64  `json:"__count,omitempty"`
	ExpressionType         string   `json:"expressiontype,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	GotoPriorityExpression string   `json:"gotopriorityexpression,omitempty"`
	Hits                   int      `json:"hits,omitempty"`
	Name                   string   `json:"name,omitempty"`
	NextGenAPIResource     string   `json:"_nextgenapiresource,omitempty"`
	Rule                   string   `json:"rule,omitempty"`
}

type TMTrafficPolicyCSVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type TMSessionPolicyTMGlobalBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type TMSessionParameter struct {
	DefaultAuthorizationAction string `json:"defaultauthorizationaction,omitempty"`
	Homepage                   string `json:"homepage,omitempty"`
	HTTPOnlyCookie             string `json:"httponlycookie,omitempty"`
	KCDAccount                 string `json:"kcdaccount,omitempty"`
	Name                       string `json:"name,omitempty"`
	NextGenAPIResource         string `json:"_nextgenapiresource,omitempty"`
	PersistentCookie           string `json:"persistentcookie,omitempty"`
	PersistentCookieValidity   int    `json:"persistentcookievalidity,omitempty"`
	SessTimeout                int    `json:"sesstimeout,omitempty"`
	SSO                        string `json:"sso,omitempty"`
	SSOCredential              string `json:"ssocredential,omitempty"`
	SSODomain                  string `json:"ssodomain,omitempty"`
	TMSessionPolicyBindType    string `json:"tmsessionpolicybindtype,omitempty"`
	TMSessionPolicyCount       int    `json:"tmsessionpolicycount,omitempty"`
}

type TMGlobalAuditNSLogPolicyBinding struct {
	BindPolicyType         int    `json:"bindpolicytype,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type TMTrafficAction struct {
	AppTimeout         int     `json:"apptimeout,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	ForcedTimeout      string  `json:"forcedtimeout,omitempty"`
	ForcedTimeoutVal   int     `json:"forcedtimeoutval,omitempty"`
	FormSSOAction      string  `json:"formssoaction,omitempty"`
	InitiateLogout     string  `json:"initiatelogout,omitempty"`
	KCDAccount         string  `json:"kcdaccount,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	PasswdExpression   string  `json:"passwdexpression,omitempty"`
	PersistentCookie   string  `json:"persistentcookie,omitempty"`
	SAMLSSOProfile     string  `json:"samlssoprofile,omitempty"`
	SSO                string  `json:"sso,omitempty"`
	UserExpression     string  `json:"userexpression,omitempty"`
}

type TMSessionPolicyBinding struct {
	Name                                        string `json:"name,omitempty"`
	TMSessionPolicyAAAGroupBinding              []any  `json:"tmsessionpolicy_aaagroup_binding,omitempty"`
	TMSessionPolicyAAAUserBinding               []any  `json:"tmsessionpolicy_aaauser_binding,omitempty"`
	TMSessionPolicyAuthenticationVServerBinding []any  `json:"tmsessionpolicy_authenticationvserver_binding,omitempty"`
	TMSessionPolicyTMGlobalBinding              []any  `json:"tmsessionpolicy_tmglobal_binding,omitempty"`
}

type TMSessionPolicyAAAGroupBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type TMTrafficPolicyTMGlobalBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type TMGlobalTMSessionPolicyBinding struct {
	BindPolicyType         int      `json:"bindpolicytype,omitempty"`
	Builtin                []string `json:"builtin,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	GotoPriorityExpression string   `json:"gotopriorityexpression,omitempty"`
	PolicyName             string   `json:"policyname,omitempty"`
	Priority               int      `json:"priority,omitempty"`
}

type TMGlobalAuditSyslogPolicyBinding struct {
	BindPolicyType         int    `json:"bindpolicytype,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type TMSAMLSSOProfile struct {
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
	SkewTime                    int     `json:"skewtime,omitempty"`
}

type TMSessionAction struct {
	Builtin                    []string `json:"builtin,omitempty"`
	Count                      float64  `json:"__count,omitempty"`
	DefaultAuthorizationAction string   `json:"defaultauthorizationaction,omitempty"`
	Feature                    string   `json:"feature,omitempty"`
	Homepage                   string   `json:"homepage,omitempty"`
	HTTPOnlyCookie             string   `json:"httponlycookie,omitempty"`
	KCDAccount                 string   `json:"kcdaccount,omitempty"`
	Name                       string   `json:"name,omitempty"`
	NextGenAPIResource         string   `json:"_nextgenapiresource,omitempty"`
	PersistentCookie           string   `json:"persistentcookie,omitempty"`
	PersistentCookieValidity   int      `json:"persistentcookievalidity,omitempty"`
	SessTimeout                int      `json:"sesstimeout,omitempty"`
	SSO                        string   `json:"sso,omitempty"`
	SSOCredential              string   `json:"ssocredential,omitempty"`
	SSODomain                  string   `json:"ssodomain,omitempty"`
}

type TMSessionPolicyAuthenticationVServerBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type TMSessionPolicyAAAUserBinding struct {
	ActivePolicy int    `json:"activepolicy,omitempty"`
	BoundTo      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}
