package models

// authentication configuration structs
type AuthenticationdfapolicyBinding struct {
	AuthenticationdfapolicyVpnvserverBinding []interface{} `json:"authenticationdfapolicy_vpnvserver_binding,omitempty"`
	Name                                     string        `json:"name,omitempty"`
}

type Authenticationwebauthpolicy struct {
	Action             string  `json:"action,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Rule               string  `json:"rule,omitempty"`
}

type AuthenticationldappolicyBinding struct {
	AuthenticationldappolicyAuthenticationvserverBinding []interface{} `json:"authenticationldappolicy_authenticationvserver_binding,omitempty"`
	AuthenticationldappolicySystemglobalBinding          []interface{} `json:"authenticationldappolicy_systemglobal_binding,omitempty"`
	AuthenticationldappolicyVpnglobalBinding             []interface{} `json:"authenticationldappolicy_vpnglobal_binding,omitempty"`
	AuthenticationldappolicyVpnvserverBinding            []interface{} `json:"authenticationldappolicy_vpnvserver_binding,omitempty"`
	Name                                                 string        `json:"name,omitempty"`
}

type AuthenticationwebauthpolicyVpnvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationpolicySystemglobalBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Authenticationnegotiatepolicy struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Reqaction          string  `json:"reqaction,omitempty"`
	Rule               string  `json:"rule,omitempty"`
}

type AuthenticationcertpolicyAuthenticationvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type Authenticationldappolicy struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Reqaction          string  `json:"reqaction,omitempty"`
	Rule               string  `json:"rule,omitempty"`
}

type Authenticationsmartaccessprofile struct {
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Tags               string  `json:"tags,omitempty"`
}

type AuthenticationlocalpolicySystemglobalBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type Authenticationpushservice struct {
	Certendpoint          string  `json:"certendpoint,omitempty"`
	Clientid              string  `json:"clientid,omitempty"`
	Clientsecret          string  `json:"clientsecret,omitempty"`
	Count                 float64 `json:"__count,omitempty"`
	Customerid            string  `json:"customerid,omitempty"`
	Hubname               string  `json:"hubname,omitempty"`
	Name                  string  `json:"name,omitempty"`
	Namespace             string  `json:"Namespace,omitempty"`
	Nextgenapiresource    string  `json:"_nextgenapiresource,omitempty"`
	Pushcloudserverstatus string  `json:"pushcloudserverstatus,omitempty"`
	Pushservicestatus     string  `json:"pushservicestatus,omitempty"`
	Refreshinterval       int     `json:"refreshinterval,omitempty"`
	Servicekey            string  `json:"servicekey,omitempty"`
	Servicekeyname        string  `json:"servicekeyname,omitempty"`
	Signingkey            string  `json:"signingkey,omitempty"`
	Signingkeyname        string  `json:"signingkeyname,omitempty"`
	Trustservice          string  `json:"trustservice,omitempty"`
}

type Authenticationstorefrontauthaction struct {
	Count                      float64 `json:"__count,omitempty"`
	Defaultauthenticationgroup string  `json:"defaultauthenticationgroup,omitempty"`
	Domain                     string  `json:"domain,omitempty"`
	Failure                    int     `json:"failure,omitempty"`
	Name                       string  `json:"name,omitempty"`
	Nextgenapiresource         string  `json:"_nextgenapiresource,omitempty"`
	Serverurl                  string  `json:"serverurl,omitempty"`
	Success                    int     `json:"success,omitempty"`
}

type AuthenticationwebauthpolicyBinding struct {
	AuthenticationwebauthpolicyAuthenticationvserverBinding []interface{} `json:"authenticationwebauthpolicy_authenticationvserver_binding,omitempty"`
	AuthenticationwebauthpolicySystemglobalBinding          []interface{} `json:"authenticationwebauthpolicy_systemglobal_binding,omitempty"`
	AuthenticationwebauthpolicyVpnglobalBinding             []interface{} `json:"authenticationwebauthpolicy_vpnglobal_binding,omitempty"`
	AuthenticationwebauthpolicyVpnvserverBinding            []interface{} `json:"authenticationwebauthpolicy_vpnvserver_binding,omitempty"`
	Name                                                    string        `json:"name,omitempty"`
}

type AuthenticationlocalpolicyAuthenticationvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationwebauthpolicyAuthenticationvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type Authenticationsmartaccesspolicy struct {
	Action             string  `json:"action,omitempty"`
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Hits               int     `json:"hits,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Rule               string  `json:"rule,omitempty"`
}

type AuthenticationsamlidppolicyBinding struct {
	AuthenticationsamlidppolicyAuthenticationvserverBinding []interface{} `json:"authenticationsamlidppolicy_authenticationvserver_binding,omitempty"`
	AuthenticationsamlidppolicyVpnvserverBinding            []interface{} `json:"authenticationsamlidppolicy_vpnvserver_binding,omitempty"`
	Name                                                    string        `json:"name,omitempty"`
}

type Authenticationoauthidppolicy struct {
	Action                 string  `json:"action,omitempty"`
	Comment                string  `json:"comment,omitempty"`
	Count                  float64 `json:"__count,omitempty"`
	Gotopriorityexpression string  `json:"gotopriorityexpression,omitempty"`
	Hits                   int     `json:"hits,omitempty"`
	Logaction              string  `json:"logaction,omitempty"`
	Name                   string  `json:"name,omitempty"`
	Newname                string  `json:"newname,omitempty"`
	Nextgenapiresource     string  `json:"_nextgenapiresource,omitempty"`
	Rule                   string  `json:"rule,omitempty"`
	Undefaction            string  `json:"undefaction,omitempty"`
}

type AuthenticationvserverBinding struct {
	AuthenticationvserverAuditnslogpolicyBinding                []interface{} `json:"authenticationvserver_auditnslogpolicy_binding,omitempty"`
	AuthenticationvserverAuditsyslogpolicyBinding               []interface{} `json:"authenticationvserver_auditsyslogpolicy_binding,omitempty"`
	AuthenticationvserverAuthenticationcertpolicyBinding        []interface{} `json:"authenticationvserver_authenticationcertpolicy_binding,omitempty"`
	AuthenticationvserverAuthenticationldappolicyBinding        []interface{} `json:"authenticationvserver_authenticationldappolicy_binding,omitempty"`
	AuthenticationvserverAuthenticationlocalpolicyBinding       []interface{} `json:"authenticationvserver_authenticationlocalpolicy_binding,omitempty"`
	AuthenticationvserverAuthenticationloginschemapolicyBinding []interface{} `json:"authenticationvserver_authenticationloginschemapolicy_binding,omitempty"`
	AuthenticationvserverAuthenticationnegotiatepolicyBinding   []interface{} `json:"authenticationvserver_authenticationnegotiatepolicy_binding,omitempty"`
	AuthenticationvserverAuthenticationoauthidppolicyBinding    []interface{} `json:"authenticationvserver_authenticationoauthidppolicy_binding,omitempty"`
	AuthenticationvserverAuthenticationpolicyBinding            []interface{} `json:"authenticationvserver_authenticationpolicy_binding,omitempty"`
	AuthenticationvserverAuthenticationradiuspolicyBinding      []interface{} `json:"authenticationvserver_authenticationradiuspolicy_binding,omitempty"`
	AuthenticationvserverAuthenticationsamlidppolicyBinding     []interface{} `json:"authenticationvserver_authenticationsamlidppolicy_binding,omitempty"`
	AuthenticationvserverAuthenticationsamlpolicyBinding        []interface{} `json:"authenticationvserver_authenticationsamlpolicy_binding,omitempty"`
	AuthenticationvserverAuthenticationsmartaccesspolicyBinding []interface{} `json:"authenticationvserver_authenticationsmartaccesspolicy_binding,omitempty"`
	AuthenticationvserverAuthenticationtacacspolicyBinding      []interface{} `json:"authenticationvserver_authenticationtacacspolicy_binding,omitempty"`
	AuthenticationvserverAuthenticationwebauthpolicyBinding     []interface{} `json:"authenticationvserver_authenticationwebauthpolicy_binding,omitempty"`
	AuthenticationvserverCachepolicyBinding                     []interface{} `json:"authenticationvserver_cachepolicy_binding,omitempty"`
	AuthenticationvserverCspolicyBinding                        []interface{} `json:"authenticationvserver_cspolicy_binding,omitempty"`
	AuthenticationvserverResponderpolicyBinding                 []interface{} `json:"authenticationvserver_responderpolicy_binding,omitempty"`
	AuthenticationvserverRewritepolicyBinding                   []interface{} `json:"authenticationvserver_rewritepolicy_binding,omitempty"`
	AuthenticationvserverTmsessionpolicyBinding                 []interface{} `json:"authenticationvserver_tmsessionpolicy_binding,omitempty"`
	AuthenticationvserverVpnportalthemeBinding                  []interface{} `json:"authenticationvserver_vpnportaltheme_binding,omitempty"`
	Name                                                        string        `json:"name,omitempty"`
}

type AuthenticationlocalpolicyBinding struct {
	AuthenticationlocalpolicyAuthenticationvserverBinding []interface{} `json:"authenticationlocalpolicy_authenticationvserver_binding,omitempty"`
	AuthenticationlocalpolicySystemglobalBinding          []interface{} `json:"authenticationlocalpolicy_systemglobal_binding,omitempty"`
	AuthenticationlocalpolicyVpnglobalBinding             []interface{} `json:"authenticationlocalpolicy_vpnglobal_binding,omitempty"`
	AuthenticationlocalpolicyVpnvserverBinding            []interface{} `json:"authenticationlocalpolicy_vpnvserver_binding,omitempty"`
	Name                                                  string        `json:"name,omitempty"`
}

type AuthenticationsmartaccesspolicyBinding struct {
	AuthenticationsmartaccesspolicyAuthenticationvserverBinding []interface{} `json:"authenticationsmartaccesspolicy_authenticationvserver_binding,omitempty"`
	Name                                                        string        `json:"name,omitempty"`
}

type AuthenticationvserverAuthenticationwebauthpolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type AuthenticationvserverAuditsyslogpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type AuthenticationradiuspolicyAuthenticationvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type Authenticationsamlidppolicy struct {
	Action                 string  `json:"action,omitempty"`
	Comment                string  `json:"comment,omitempty"`
	Count                  float64 `json:"__count,omitempty"`
	Gotopriorityexpression string  `json:"gotopriorityexpression,omitempty"`
	Hits                   int     `json:"hits,omitempty"`
	Logaction              string  `json:"logaction,omitempty"`
	Name                   string  `json:"name,omitempty"`
	Newname                string  `json:"newname,omitempty"`
	Nextgenapiresource     string  `json:"_nextgenapiresource,omitempty"`
	Rule                   string  `json:"rule,omitempty"`
	Undefaction            string  `json:"undefaction,omitempty"`
}

type Authenticationradiusaction struct {
	Accounting                 string  `json:"accounting,omitempty"`
	Authentication             string  `json:"authentication,omitempty"`
	Authservretry              int     `json:"authservretry,omitempty"`
	Authtimeout                int     `json:"authtimeout,omitempty"`
	Callingstationid           string  `json:"callingstationid,omitempty"`
	Count                      float64 `json:"__count,omitempty"`
	Defaultauthenticationgroup string  `json:"defaultauthenticationgroup,omitempty"`
	Failure                    int     `json:"failure,omitempty"`
	Ipaddress                  string  `json:"ipaddress,omitempty"`
	Ipattributetype            int     `json:"ipattributetype,omitempty"`
	Ipvendorid                 int     `json:"ipvendorid,omitempty"`
	Messageauthenticator       string  `json:"messageauthenticator,omitempty"`
	Name                       string  `json:"name,omitempty"`
	Nextgenapiresource         string  `json:"_nextgenapiresource,omitempty"`
	Passencoding               string  `json:"passencoding,omitempty"`
	Pwdattributetype           int     `json:"pwdattributetype,omitempty"`
	Pwdvendorid                int     `json:"pwdvendorid,omitempty"`
	Radattributetype           int     `json:"radattributetype,omitempty"`
	Radgroupseparator          string  `json:"radgroupseparator,omitempty"`
	Radgroupsprefix            string  `json:"radgroupsprefix,omitempty"`
	Radkey                     string  `json:"radkey,omitempty"`
	Radnasid                   string  `json:"radnasid,omitempty"`
	Radnasip                   string  `json:"radnasip,omitempty"`
	Radvendorid                int     `json:"radvendorid,omitempty"`
	Serverip                   string  `json:"serverip,omitempty"`
	Servername                 string  `json:"servername,omitempty"`
	Serverport                 int     `json:"serverport,omitempty"`
	Success                    int     `json:"success,omitempty"`
	Targetlbvserver            string  `json:"targetlbvserver,omitempty"`
	Transport                  string  `json:"transport,omitempty"`
	Tunnelendpointclientip     string  `json:"tunnelendpointclientip,omitempty"`
}

type Authenticationazurekeyvault struct {
	Authentication             string  `json:"authentication,omitempty"`
	Clientid                   string  `json:"clientid,omitempty"`
	Clientsecret               string  `json:"clientsecret,omitempty"`
	Count                      float64 `json:"__count,omitempty"`
	Defaultauthenticationgroup string  `json:"defaultauthenticationgroup,omitempty"`
	Name                       string  `json:"name,omitempty"`
	Nextgenapiresource         string  `json:"_nextgenapiresource,omitempty"`
	Pushservice                string  `json:"pushservice,omitempty"`
	Refreshinterval            int     `json:"refreshinterval,omitempty"`
	Servicekeyname             string  `json:"servicekeyname,omitempty"`
	Signaturealg               string  `json:"signaturealg,omitempty"`
	Tenantid                   string  `json:"tenantid,omitempty"`
	Tokenendpoint              string  `json:"tokenendpoint,omitempty"`
	Vaultname                  string  `json:"vaultname,omitempty"`
}

type Authenticationemailaction struct {
	Content                    string  `json:"content,omitempty"`
	Count                      float64 `json:"__count,omitempty"`
	Defaultauthenticationgroup string  `json:"defaultauthenticationgroup,omitempty"`
	Emailaddress               string  `json:"emailaddress,omitempty"`
	Name                       string  `json:"name,omitempty"`
	Nextgenapiresource         string  `json:"_nextgenapiresource,omitempty"`
	Password                   string  `json:"password,omitempty"`
	Serverurl                  string  `json:"serverurl,omitempty"`
	Timeout                    int     `json:"timeout,omitempty"`
	TypeField                  string  `json:"type,omitempty"`
	Username                   string  `json:"username,omitempty"`
}

type AuthenticationvserverTmsessionpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type Authenticationtacacspolicy struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Reqaction          string  `json:"reqaction,omitempty"`
	Rule               string  `json:"rule,omitempty"`
}

type AuthenticationldappolicyAuthenticationvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationnegotiatepolicyBinding struct {
	AuthenticationnegotiatepolicyAuthenticationvserverBinding []interface{} `json:"authenticationnegotiatepolicy_authenticationvserver_binding,omitempty"`
	AuthenticationnegotiatepolicyVpnglobalBinding             []interface{} `json:"authenticationnegotiatepolicy_vpnglobal_binding,omitempty"`
	AuthenticationnegotiatepolicyVpnvserverBinding            []interface{} `json:"authenticationnegotiatepolicy_vpnvserver_binding,omitempty"`
	Name                                                      string        `json:"name,omitempty"`
}

type AuthenticationsmartaccesspolicyAuthenticationvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationloginschemapolicyVpnvserverBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type AuthenticationnegotiatepolicyAuthenticationvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type Authenticationsamlaction struct {
	Artifactresolutionserviceurl   string   `json:"artifactresolutionserviceurl,omitempty"`
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
	Attributeconsumingserviceindex int      `json:"attributeconsumingserviceindex,omitempty"`
	Attributes                     string   `json:"attributes,omitempty"`
	Audience                       string   `json:"audience,omitempty"`
	Authnctxclassref               []string `json:"authnctxclassref,omitempty"`
	Count                          float64  `json:"__count,omitempty"`
	Customauthnctxclassref         string   `json:"customauthnctxclassref,omitempty"`
	Defaultauthenticationgroup     string   `json:"defaultauthenticationgroup,omitempty"`
	Digestmethod                   string   `json:"digestmethod,omitempty"`
	Enforceusername                string   `json:"enforceusername,omitempty"`
	Forceauthn                     string   `json:"forceauthn,omitempty"`
	Groupnamefield                 string   `json:"groupnamefield,omitempty"`
	Logoutbinding                  string   `json:"logoutbinding,omitempty"`
	Logouturl                      string   `json:"logouturl,omitempty"`
	Metadataimportstatus           string   `json:"metadataimportstatus,omitempty"`
	Metadatarefreshinterval        int      `json:"metadatarefreshinterval,omitempty"`
	Metadataurl                    string   `json:"metadataurl,omitempty"`
	Name                           string   `json:"name,omitempty"`
	Nextgenapiresource             string   `json:"_nextgenapiresource,omitempty"`
	Preferredbindtype              []string `json:"preferredbindtype,omitempty"`
	Relaystaterule                 string   `json:"relaystaterule,omitempty"`
	Requestedauthncontext          string   `json:"requestedauthncontext,omitempty"`
	Samlacsindex                   int      `json:"samlacsindex,omitempty"`
	Samlbinding                    string   `json:"samlbinding,omitempty"`
	Samlidpcertname                string   `json:"samlidpcertname,omitempty"`
	Samlissuername                 string   `json:"samlissuername,omitempty"`
	Samlredirecturl                string   `json:"samlredirecturl,omitempty"`
	Samlrejectunsignedassertion    string   `json:"samlrejectunsignedassertion,omitempty"`
	Samlsigningcertname            string   `json:"samlsigningcertname,omitempty"`
	Samltwofactor                  string   `json:"samltwofactor,omitempty"`
	Samluserfield                  string   `json:"samluserfield,omitempty"`
	Sendthumbprint                 string   `json:"sendthumbprint,omitempty"`
	Signaturealg                   string   `json:"signaturealg,omitempty"`
	Skewtime                       int      `json:"skewtime,omitempty"`
	Statechecks                    string   `json:"statechecks,omitempty"`
	Storesamlresponse              string   `json:"storesamlresponse,omitempty"`
}

type AuthenticationvserverVpnportalthemeBinding struct {
	Acttype     int    `json:"acttype,omitempty"`
	Name        string `json:"name,omitempty"`
	Portaltheme string `json:"portaltheme,omitempty"`
}

type AuthenticationlocalpolicyVpnglobalBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type Authenticationloginschemapolicy struct {
	Action             string   `json:"action,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	Comment            string   `json:"comment,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	Logaction          string   `json:"logaction,omitempty"`
	Name               string   `json:"name,omitempty"`
	Newname            string   `json:"newname,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Rule               string   `json:"rule,omitempty"`
	Undefaction        string   `json:"undefaction,omitempty"`
	Undefhits          int      `json:"undefhits,omitempty"`
}

type AuthenticationloginschemapolicyBinding struct {
	AuthenticationloginschemapolicyAuthenticationvserverBinding []interface{} `json:"authenticationloginschemapolicy_authenticationvserver_binding,omitempty"`
	AuthenticationloginschemapolicyVpnvserverBinding            []interface{} `json:"authenticationloginschemapolicy_vpnvserver_binding,omitempty"`
	Name                                                        string        `json:"name,omitempty"`
}

type Authenticationprotecteduseraction struct {
	Count              float64 `json:"__count,omitempty"`
	Maxconcurrentusers int     `json:"maxconcurrentusers,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Realmstr           string  `json:"realmstr,omitempty"`
}

type AuthenticationpolicyAuthenticationvserverBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type AuthenticationsamlpolicyVpnglobalBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type Authenticationtacacsaction struct {
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
	Auditfailedcmds            string  `json:"auditfailedcmds,omitempty"`
	Authorization              string  `json:"authorization,omitempty"`
	Authtimeout                int     `json:"authtimeout,omitempty"`
	Count                      float64 `json:"__count,omitempty"`
	Defaultauthenticationgroup string  `json:"defaultauthenticationgroup,omitempty"`
	Failure                    int     `json:"failure,omitempty"`
	Groupattrname              string  `json:"groupattrname,omitempty"`
	Name                       string  `json:"name,omitempty"`
	Nextgenapiresource         string  `json:"_nextgenapiresource,omitempty"`
	Serverip                   string  `json:"serverip,omitempty"`
	Serverport                 int     `json:"serverport,omitempty"`
	Success                    int     `json:"success,omitempty"`
	Tacacssecret               string  `json:"tacacssecret,omitempty"`
}

type AuthenticationcertpolicyVpnvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationvserverAuthenticationloginschemapolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type AuthenticationsamlidppolicyAuthenticationvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationnegotiatepolicyVpnglobalBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationsamlpolicyBinding struct {
	AuthenticationsamlpolicyAuthenticationvserverBinding []interface{} `json:"authenticationsamlpolicy_authenticationvserver_binding,omitempty"`
	AuthenticationsamlpolicyVpnglobalBinding             []interface{} `json:"authenticationsamlpolicy_vpnglobal_binding,omitempty"`
	AuthenticationsamlpolicyVpnvserverBinding            []interface{} `json:"authenticationsamlpolicy_vpnvserver_binding,omitempty"`
	Name                                                 string        `json:"name,omitempty"`
}

type AuthenticationvserverAuthenticationlocalpolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type AuthenticationradiuspolicySystemglobalBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type Authenticationvserver struct {
	Appflowlog           string  `json:"appflowlog,omitempty"`
	Authentication       string  `json:"authentication,omitempty"`
	Authenticationdomain string  `json:"authenticationdomain,omitempty"`
	Backupvserver        string  `json:"backupvserver,omitempty"`
	Bindpoint            string  `json:"bindpoint,omitempty"`
	Cachetype            string  `json:"cachetype,omitempty"`
	Cachevserver         string  `json:"cachevserver,omitempty"`
	Certkeynames         string  `json:"certkeynames,omitempty"`
	Clttimeout           int     `json:"clttimeout,omitempty"`
	Comment              string  `json:"comment,omitempty"`
	Count                float64 `json:"__count,omitempty"`
	Curaaausers          int     `json:"curaaausers,omitempty"`
	Curstate             string  `json:"curstate,omitempty"`
	Disableprimaryondown string  `json:"disableprimaryondown,omitempty"`
	Downstateflush       string  `json:"downstateflush,omitempty"`
	Failedlogintimeout   int     `json:"failedlogintimeout,omitempty"`
	Groupextraction      bool    `json:"groupextraction,omitempty"`
	Httpprofilename      string  `json:"httpprofilename,omitempty"`
	Ip                   string  `json:"ip,omitempty"`
	Ipv46                string  `json:"ipv46,omitempty"`
	Listenpolicy         string  `json:"listenpolicy,omitempty"`
	Listenpriority       int     `json:"listenpriority,omitempty"`
	Maxloginattempts     int     `json:"maxloginattempts,omitempty"`
	Name                 string  `json:"name,omitempty"`
	Newname              string  `json:"newname,omitempty"`
	Nextgenapiresource   string  `json:"_nextgenapiresource,omitempty"`
	Ngname               string  `json:"ngname,omitempty"`
	Policy               string  `json:"policy,omitempty"`
	Port                 int     `json:"port,omitempty"`
	Precedence           string  `json:"precedence,omitempty"`
	Priority             int     `json:"priority,omitempty"`
	Range                int     `json:"range,omitempty"`
	Redirect             string  `json:"redirect,omitempty"`
	Redirecturl          string  `json:"redirecturl,omitempty"`
	Samesite             string  `json:"samesite,omitempty"`
	Secondary            bool    `json:"secondary,omitempty"`
	Servicename          string  `json:"servicename,omitempty"`
	Servicetype          string  `json:"servicetype,omitempty"`
	Somethod             string  `json:"somethod,omitempty"`
	Sopersistence        string  `json:"sopersistence,omitempty"`
	Sopersistencetimeout int     `json:"sopersistencetimeout,omitempty"`
	Sothreshold          int     `json:"sothreshold,omitempty"`
	State                string  `json:"state,omitempty"`
	Status               int     `json:"status,omitempty"`
	Tcpprofilename       string  `json:"tcpprofilename,omitempty"`
	Td                   int     `json:"td,omitempty"`
	TypeField            string  `json:"type,omitempty"`
	Value                string  `json:"value,omitempty"`
	Vstype               int     `json:"vstype,omitempty"`
	Weight               int     `json:"weight,omitempty"`
}

type Authenticationnoauthaction struct {
	Builtin                    []string `json:"builtin,omitempty"`
	Count                      float64  `json:"__count,omitempty"`
	Defaultauthenticationgroup string   `json:"defaultauthenticationgroup,omitempty"`
	Feature                    string   `json:"feature,omitempty"`
	Name                       string   `json:"name,omitempty"`
	Nextgenapiresource         string   `json:"_nextgenapiresource,omitempty"`
}

type AuthenticationnegotiatepolicyVpnvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationsamlpolicyVpnvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type Authenticationauthnprofile struct {
	Authenticationdomain string  `json:"authenticationdomain,omitempty"`
	Authenticationhost   string  `json:"authenticationhost,omitempty"`
	Authenticationlevel  int     `json:"authenticationlevel,omitempty"`
	Authnvsname          string  `json:"authnvsname,omitempty"`
	Count                float64 `json:"__count,omitempty"`
	Name                 string  `json:"name,omitempty"`
	Nextgenapiresource   string  `json:"_nextgenapiresource,omitempty"`
}

type Authenticationcaptchaaction struct {
	Count                      float64 `json:"__count,omitempty"`
	Defaultauthenticationgroup string  `json:"defaultauthenticationgroup,omitempty"`
	Name                       string  `json:"name,omitempty"`
	Nextgenapiresource         string  `json:"_nextgenapiresource,omitempty"`
	Scorethreshold             int     `json:"scorethreshold,omitempty"`
	Secretkey                  string  `json:"secretkey,omitempty"`
	Serverurl                  string  `json:"serverurl,omitempty"`
	Sitekey                    string  `json:"sitekey,omitempty"`
}

type AuthenticationradiuspolicyBinding struct {
	AuthenticationradiuspolicyAuthenticationvserverBinding []interface{} `json:"authenticationradiuspolicy_authenticationvserver_binding,omitempty"`
	AuthenticationradiuspolicySystemglobalBinding          []interface{} `json:"authenticationradiuspolicy_systemglobal_binding,omitempty"`
	AuthenticationradiuspolicyVpnglobalBinding             []interface{} `json:"authenticationradiuspolicy_vpnglobal_binding,omitempty"`
	AuthenticationradiuspolicyVpnvserverBinding            []interface{} `json:"authenticationradiuspolicy_vpnvserver_binding,omitempty"`
	Name                                                   string        `json:"name,omitempty"`
}

type AuthenticationtacacspolicyBinding struct {
	AuthenticationtacacspolicyAuthenticationvserverBinding []interface{} `json:"authenticationtacacspolicy_authenticationvserver_binding,omitempty"`
	AuthenticationtacacspolicySystemglobalBinding          []interface{} `json:"authenticationtacacspolicy_systemglobal_binding,omitempty"`
	AuthenticationtacacspolicyVpnglobalBinding             []interface{} `json:"authenticationtacacspolicy_vpnglobal_binding,omitempty"`
	AuthenticationtacacspolicyVpnvserverBinding            []interface{} `json:"authenticationtacacspolicy_vpnvserver_binding,omitempty"`
	Name                                                   string        `json:"name,omitempty"`
}

type Authenticationsamlpolicy struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Reqaction          string  `json:"reqaction,omitempty"`
	Rule               string  `json:"rule,omitempty"`
}

type AuthenticationldappolicyVpnvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationwebauthpolicyVpnglobalBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type Authenticationlocalpolicy struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Reqaction          string  `json:"reqaction,omitempty"`
	Rule               string  `json:"rule,omitempty"`
}

type AuthenticationvserverAuthenticationradiuspolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type AuthenticationloginschemapolicyAuthenticationvserverBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type AuthenticationlocalpolicyVpnvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationvserverAuthenticationsamlpolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type AuthenticationsamlidppolicyVpnvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationoauthidppolicyAuthenticationvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type Authenticationcitrixauthaction struct {
	Authentication     string  `json:"authentication,omitempty"`
	Authenticationtype string  `json:"authenticationtype,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
}

type AuthenticationvserverAuthenticationcertpolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type AuthenticationpolicylabelBinding struct {
	AuthenticationpolicylabelAuthenticationpolicyBinding []interface{} `json:"authenticationpolicylabel_authenticationpolicy_binding,omitempty"`
	Labelname                                            string        `json:"labelname,omitempty"`
}

type Authenticationdfapolicy struct {
	Action             string  `json:"action,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Rule               string  `json:"rule,omitempty"`
}

type AuthenticationtacacspolicyVpnvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationradiuspolicyVpnglobalBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationvserverAuthenticationldappolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type Authenticationoauthidpprofile struct {
	Attributes                 string  `json:"attributes,omitempty"`
	Audience                   string  `json:"audience,omitempty"`
	Clientid                   string  `json:"clientid,omitempty"`
	Clientsecret               string  `json:"clientsecret,omitempty"`
	Configservice              string  `json:"configservice,omitempty"`
	Count                      float64 `json:"__count,omitempty"`
	Defaultauthenticationgroup string  `json:"defaultauthenticationgroup,omitempty"`
	Encrypttoken               string  `json:"encrypttoken,omitempty"`
	Issuer                     string  `json:"issuer,omitempty"`
	Name                       string  `json:"name,omitempty"`
	Nextgenapiresource         string  `json:"_nextgenapiresource,omitempty"`
	Oauthstatus                string  `json:"oauthstatus,omitempty"`
	Redirecturl                string  `json:"redirecturl,omitempty"`
	Refreshinterval            int     `json:"refreshinterval,omitempty"`
	Relyingpartymetadataurl    string  `json:"relyingpartymetadataurl,omitempty"`
	Sendpassword               string  `json:"sendpassword,omitempty"`
	Signaturealg               string  `json:"signaturealg,omitempty"`
	Signatureservice           string  `json:"signatureservice,omitempty"`
	Skewtime                   int     `json:"skewtime,omitempty"`
}

type AuthenticationtacacspolicySystemglobalBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type Authenticationepaaction struct {
	Count              float64 `json:"__count,omitempty"`
	Csecexpr           string  `json:"csecexpr,omitempty"`
	Defaultepagroup    string  `json:"defaultepagroup,omitempty"`
	Deletefiles        string  `json:"deletefiles,omitempty"`
	Deviceposture      string  `json:"deviceposture,omitempty"`
	Killprocess        string  `json:"killprocess,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Quarantinegroup    string  `json:"quarantinegroup,omitempty"`
}

type AuthenticationpolicyBinding struct {
	AuthenticationpolicyAuthenticationpolicylabelBinding []interface{} `json:"authenticationpolicy_authenticationpolicylabel_binding,omitempty"`
	AuthenticationpolicyAuthenticationvserverBinding     []interface{} `json:"authenticationpolicy_authenticationvserver_binding,omitempty"`
	AuthenticationpolicySystemglobalBinding              []interface{} `json:"authenticationpolicy_systemglobal_binding,omitempty"`
	Name                                                 string        `json:"name,omitempty"`
}

type AuthenticationcertpolicyVpnglobalBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationdfapolicyVpnvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type Authenticationwebauthaction struct {
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
	Defaultauthenticationgroup string  `json:"defaultauthenticationgroup,omitempty"`
	Fullreqexpr                string  `json:"fullreqexpr,omitempty"`
	Name                       string  `json:"name,omitempty"`
	Nextgenapiresource         string  `json:"_nextgenapiresource,omitempty"`
	Scheme                     string  `json:"scheme,omitempty"`
	Serverip                   string  `json:"serverip,omitempty"`
	Serverport                 int     `json:"serverport,omitempty"`
	Successrule                string  `json:"successrule,omitempty"`
}

type Authenticationpolicylabel struct {
	Comment                string  `json:"comment,omitempty"`
	Count                  float64 `json:"__count,omitempty"`
	Description            string  `json:"description,omitempty"`
	Flowtype               int     `json:"flowtype,omitempty"`
	Gotopriorityexpression string  `json:"gotopriorityexpression,omitempty"`
	Hits                   int     `json:"hits,omitempty"`
	Labelname              string  `json:"labelname,omitempty"`
	Loginschema            string  `json:"loginschema,omitempty"`
	Newname                string  `json:"newname,omitempty"`
	Nextgenapiresource     string  `json:"_nextgenapiresource,omitempty"`
	Numpol                 int     `json:"numpol,omitempty"`
	Policyname             string  `json:"policyname,omitempty"`
	Priority               int     `json:"priority,omitempty"`
	TypeField              string  `json:"type,omitempty"`
}

type Authenticationradiuspolicy struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Reqaction          string  `json:"reqaction,omitempty"`
	Rule               string  `json:"rule,omitempty"`
}

type AuthenticationvserverAuthenticationoauthidppolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type AuthenticationvserverRewritepolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type Authenticationdfaaction struct {
	Clientid                   string  `json:"clientid,omitempty"`
	Count                      float64 `json:"__count,omitempty"`
	Defaultauthenticationgroup string  `json:"defaultauthenticationgroup,omitempty"`
	Failure                    int     `json:"failure,omitempty"`
	Name                       string  `json:"name,omitempty"`
	Nextgenapiresource         string  `json:"_nextgenapiresource,omitempty"`
	Passphrase                 string  `json:"passphrase,omitempty"`
	Serverurl                  string  `json:"serverurl,omitempty"`
	Success                    int     `json:"success,omitempty"`
}

type AuthenticationoauthidppolicyBinding struct {
	AuthenticationoauthidppolicyAuthenticationvserverBinding []interface{} `json:"authenticationoauthidppolicy_authenticationvserver_binding,omitempty"`
	AuthenticationoauthidppolicyVpnvserverBinding            []interface{} `json:"authenticationoauthidppolicy_vpnvserver_binding,omitempty"`
	Name                                                     string        `json:"name,omitempty"`
}

type AuthenticationvserverCachepolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type AuthenticationvserverAuthenticationnegotiatepolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type AuthenticationtacacspolicyAuthenticationvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationtacacspolicyVpnglobalBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationvserverAuthenticationsamlidppolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type AuthenticationvserverAuthenticationtacacspolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type Authenticationldapaction struct {
	Alternateemailattr         string  `json:"alternateemailattr,omitempty"`
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
	Authtimeout                int     `json:"authtimeout,omitempty"`
	Cloudattributes            string  `json:"cloudattributes,omitempty"`
	Count                      float64 `json:"__count,omitempty"`
	Defaultauthenticationgroup string  `json:"defaultauthenticationgroup,omitempty"`
	Email                      string  `json:"email,omitempty"`
	Failure                    int     `json:"failure,omitempty"`
	Followreferrals            string  `json:"followreferrals,omitempty"`
	Groupattrname              string  `json:"groupattrname,omitempty"`
	Groupnameidentifier        string  `json:"groupnameidentifier,omitempty"`
	Groupsearchattribute       string  `json:"groupsearchattribute,omitempty"`
	Groupsearchfilter          string  `json:"groupsearchfilter,omitempty"`
	Groupsearchsubattribute    string  `json:"groupsearchsubattribute,omitempty"`
	Kbattribute                string  `json:"kbattribute,omitempty"`
	Ldapbase                   string  `json:"ldapbase,omitempty"`
	Ldapbinddn                 string  `json:"ldapbinddn,omitempty"`
	Ldapbinddnpassword         string  `json:"ldapbinddnpassword,omitempty"`
	Ldaphostname               string  `json:"ldaphostname,omitempty"`
	Ldaploginname              string  `json:"ldaploginname,omitempty"`
	Maxldapreferrals           int     `json:"maxldapreferrals,omitempty"`
	Maxnestinglevel            int     `json:"maxnestinglevel,omitempty"`
	Mssrvrecordlocation        string  `json:"mssrvrecordlocation,omitempty"`
	Name                       string  `json:"name,omitempty"`
	Nestedgroupextraction      string  `json:"nestedgroupextraction,omitempty"`
	Nextgenapiresource         string  `json:"_nextgenapiresource,omitempty"`
	Otpsecret                  string  `json:"otpsecret,omitempty"`
	Passwdchange               string  `json:"passwdchange,omitempty"`
	Pushservice                string  `json:"pushservice,omitempty"`
	Referraldnslookup          string  `json:"referraldnslookup,omitempty"`
	Requireuser                string  `json:"requireuser,omitempty"`
	Searchfilter               string  `json:"searchfilter,omitempty"`
	Sectype                    string  `json:"sectype,omitempty"`
	Serverip                   string  `json:"serverip,omitempty"`
	Servername                 string  `json:"servername,omitempty"`
	Serverport                 int     `json:"serverport,omitempty"`
	Sshpublickey               string  `json:"sshpublickey,omitempty"`
	Ssonameattribute           string  `json:"ssonameattribute,omitempty"`
	Subattributename           string  `json:"subattributename,omitempty"`
	Success                    int     `json:"success,omitempty"`
	Svrtype                    string  `json:"svrtype,omitempty"`
	Validateservercert         string  `json:"validateservercert,omitempty"`
}

type AuthenticationvserverAuthenticationpolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type AuthenticationcertpolicyBinding struct {
	AuthenticationcertpolicyAuthenticationvserverBinding []interface{} `json:"authenticationcertpolicy_authenticationvserver_binding,omitempty"`
	AuthenticationcertpolicyVpnglobalBinding             []interface{} `json:"authenticationcertpolicy_vpnglobal_binding,omitempty"`
	AuthenticationcertpolicyVpnvserverBinding            []interface{} `json:"authenticationcertpolicy_vpnvserver_binding,omitempty"`
	Name                                                 string        `json:"name,omitempty"`
}

type AuthenticationwebauthpolicySystemglobalBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationvserverAuditnslogpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type AuthenticationsamlpolicyAuthenticationvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type Authenticationoauthaction struct {
	Allowedalgorithms          []string `json:"allowedalgorithms,omitempty"`
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
	Authorizationendpoint      string   `json:"authorizationendpoint,omitempty"`
	Certendpoint               string   `json:"certendpoint,omitempty"`
	Certfilepath               string   `json:"certfilepath,omitempty"`
	Clientid                   string   `json:"clientid,omitempty"`
	Clientsecret               string   `json:"clientsecret,omitempty"`
	Count                      float64  `json:"__count,omitempty"`
	Defaultauthenticationgroup string   `json:"defaultauthenticationgroup,omitempty"`
	Granttype                  string   `json:"granttype,omitempty"`
	Graphendpoint              string   `json:"graphendpoint,omitempty"`
	Idtokendecryptendpoint     string   `json:"idtokendecryptendpoint,omitempty"`
	Introspecturl              string   `json:"introspecturl,omitempty"`
	Intunedeviceidexpression   string   `json:"intunedeviceidexpression,omitempty"`
	Issuer                     string   `json:"issuer,omitempty"`
	Metadataurl                string   `json:"metadataurl,omitempty"`
	Name                       string   `json:"name,omitempty"`
	Nextgenapiresource         string   `json:"_nextgenapiresource,omitempty"`
	Oauthmiscflags             []string `json:"oauthmiscflags,omitempty"`
	Oauthstatus                string   `json:"oauthstatus,omitempty"`
	Oauthtype                  string   `json:"oauthtype,omitempty"`
	Pkce                       string   `json:"pkce,omitempty"`
	Refreshinterval            int      `json:"refreshinterval,omitempty"`
	Requestattribute           string   `json:"requestattribute,omitempty"`
	Resourceuri                string   `json:"resourceuri,omitempty"`
	Skewtime                   int      `json:"skewtime,omitempty"`
	Tenantid                   string   `json:"tenantid,omitempty"`
	Tokenendpoint              string   `json:"tokenendpoint,omitempty"`
	Tokenendpointauthmethod    string   `json:"tokenendpointauthmethod,omitempty"`
	Userinfourl                string   `json:"userinfourl,omitempty"`
	Usernamefield              string   `json:"usernamefield,omitempty"`
}

type Authenticationnegotiateaction struct {
	Count                      float64 `json:"__count,omitempty"`
	Defaultauthenticationgroup string  `json:"defaultauthenticationgroup,omitempty"`
	Domain                     string  `json:"domain,omitempty"`
	Domainuser                 string  `json:"domainuser,omitempty"`
	Domainuserpasswd           string  `json:"domainuserpasswd,omitempty"`
	Kcdspn                     string  `json:"kcdspn,omitempty"`
	Keytab                     string  `json:"keytab,omitempty"`
	Name                       string  `json:"name,omitempty"`
	Nextgenapiresource         string  `json:"_nextgenapiresource,omitempty"`
	Ntlmpath                   string  `json:"ntlmpath,omitempty"`
	Ou                         string  `json:"ou,omitempty"`
}

type AuthenticationldappolicySystemglobalBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type Authenticationloginschema struct {
	Authenticationschema    string   `json:"authenticationschema,omitempty"`
	Authenticationstrength  int      `json:"authenticationstrength,omitempty"`
	Builtin                 []string `json:"builtin,omitempty"`
	Count                   float64  `json:"__count,omitempty"`
	Feature                 string   `json:"feature,omitempty"`
	Name                    string   `json:"name,omitempty"`
	Nextgenapiresource      string   `json:"_nextgenapiresource,omitempty"`
	Passwdexpression        string   `json:"passwdexpression,omitempty"`
	Passwordcredentialindex int      `json:"passwordcredentialindex,omitempty"`
	Ssocredentials          string   `json:"ssocredentials,omitempty"`
	Usercredentialindex     int      `json:"usercredentialindex,omitempty"`
	Userexpression          string   `json:"userexpression,omitempty"`
}

type AuthenticationpolicylabelAuthenticationpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type AuthenticationradiuspolicyVpnvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type Authenticationcertaction struct {
	Count                      float64 `json:"__count,omitempty"`
	Defaultauthenticationgroup string  `json:"defaultauthenticationgroup,omitempty"`
	Groupnamefield             string  `json:"groupnamefield,omitempty"`
	Name                       string  `json:"name,omitempty"`
	Nextgenapiresource         string  `json:"_nextgenapiresource,omitempty"`
	Twofactor                  string  `json:"twofactor,omitempty"`
	Usernamefield              string  `json:"usernamefield,omitempty"`
}

type Authenticationadfsproxyprofile struct {
	Adfstruststatus    string  `json:"adfstruststatus,omitempty"`
	Certkeyname        string  `json:"certkeyname,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Password           string  `json:"password,omitempty"`
	Serverurl          string  `json:"serverurl,omitempty"`
	Username           string  `json:"username,omitempty"`
}

type AuthenticationvserverResponderpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type AuthenticationvserverCspolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type AuthenticationpolicyAuthenticationpolicylabelBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type AuthenticationoauthidppolicyVpnvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type Authenticationpolicy struct {
	Action             string  `json:"action,omitempty"`
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Description        string  `json:"description,omitempty"`
	Hits               int     `json:"hits,omitempty"`
	Logaction          string  `json:"logaction,omitempty"`
	Name               string  `json:"name,omitempty"`
	Newname            string  `json:"newname,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Policysubtype      string  `json:"policysubtype,omitempty"`
	Rule               string  `json:"rule,omitempty"`
	Undefaction        string  `json:"undefaction,omitempty"`
}

type Authenticationsamlidpprofile struct {
	Acsurlrule                  string  `json:"acsurlrule,omitempty"`
	Assertionconsumerserviceurl string  `json:"assertionconsumerserviceurl,omitempty"`
	Attribute1                  string  `json:"attribute1,omitempty"`
	Attribute10                 string  `json:"attribute10,omitempty"`
	Attribute10expr             string  `json:"attribute10expr,omitempty"`
	Attribute10format           string  `json:"attribute10format,omitempty"`
	Attribute10friendlyname     string  `json:"attribute10friendlyname,omitempty"`
	Attribute11                 string  `json:"attribute11,omitempty"`
	Attribute11expr             string  `json:"attribute11expr,omitempty"`
	Attribute11format           string  `json:"attribute11format,omitempty"`
	Attribute11friendlyname     string  `json:"attribute11friendlyname,omitempty"`
	Attribute12                 string  `json:"attribute12,omitempty"`
	Attribute12expr             string  `json:"attribute12expr,omitempty"`
	Attribute12format           string  `json:"attribute12format,omitempty"`
	Attribute12friendlyname     string  `json:"attribute12friendlyname,omitempty"`
	Attribute13                 string  `json:"attribute13,omitempty"`
	Attribute13expr             string  `json:"attribute13expr,omitempty"`
	Attribute13format           string  `json:"attribute13format,omitempty"`
	Attribute13friendlyname     string  `json:"attribute13friendlyname,omitempty"`
	Attribute14                 string  `json:"attribute14,omitempty"`
	Attribute14expr             string  `json:"attribute14expr,omitempty"`
	Attribute14format           string  `json:"attribute14format,omitempty"`
	Attribute14friendlyname     string  `json:"attribute14friendlyname,omitempty"`
	Attribute15                 string  `json:"attribute15,omitempty"`
	Attribute15expr             string  `json:"attribute15expr,omitempty"`
	Attribute15format           string  `json:"attribute15format,omitempty"`
	Attribute15friendlyname     string  `json:"attribute15friendlyname,omitempty"`
	Attribute16                 string  `json:"attribute16,omitempty"`
	Attribute16expr             string  `json:"attribute16expr,omitempty"`
	Attribute16format           string  `json:"attribute16format,omitempty"`
	Attribute16friendlyname     string  `json:"attribute16friendlyname,omitempty"`
	Attribute1expr              string  `json:"attribute1expr,omitempty"`
	Attribute1format            string  `json:"attribute1format,omitempty"`
	Attribute1friendlyname      string  `json:"attribute1friendlyname,omitempty"`
	Attribute2                  string  `json:"attribute2,omitempty"`
	Attribute2expr              string  `json:"attribute2expr,omitempty"`
	Attribute2format            string  `json:"attribute2format,omitempty"`
	Attribute2friendlyname      string  `json:"attribute2friendlyname,omitempty"`
	Attribute3                  string  `json:"attribute3,omitempty"`
	Attribute3expr              string  `json:"attribute3expr,omitempty"`
	Attribute3format            string  `json:"attribute3format,omitempty"`
	Attribute3friendlyname      string  `json:"attribute3friendlyname,omitempty"`
	Attribute4                  string  `json:"attribute4,omitempty"`
	Attribute4expr              string  `json:"attribute4expr,omitempty"`
	Attribute4format            string  `json:"attribute4format,omitempty"`
	Attribute4friendlyname      string  `json:"attribute4friendlyname,omitempty"`
	Attribute5                  string  `json:"attribute5,omitempty"`
	Attribute5expr              string  `json:"attribute5expr,omitempty"`
	Attribute5format            string  `json:"attribute5format,omitempty"`
	Attribute5friendlyname      string  `json:"attribute5friendlyname,omitempty"`
	Attribute6                  string  `json:"attribute6,omitempty"`
	Attribute6expr              string  `json:"attribute6expr,omitempty"`
	Attribute6format            string  `json:"attribute6format,omitempty"`
	Attribute6friendlyname      string  `json:"attribute6friendlyname,omitempty"`
	Attribute7                  string  `json:"attribute7,omitempty"`
	Attribute7expr              string  `json:"attribute7expr,omitempty"`
	Attribute7format            string  `json:"attribute7format,omitempty"`
	Attribute7friendlyname      string  `json:"attribute7friendlyname,omitempty"`
	Attribute8                  string  `json:"attribute8,omitempty"`
	Attribute8expr              string  `json:"attribute8expr,omitempty"`
	Attribute8format            string  `json:"attribute8format,omitempty"`
	Attribute8friendlyname      string  `json:"attribute8friendlyname,omitempty"`
	Attribute9                  string  `json:"attribute9,omitempty"`
	Attribute9expr              string  `json:"attribute9expr,omitempty"`
	Attribute9format            string  `json:"attribute9format,omitempty"`
	Attribute9friendlyname      string  `json:"attribute9friendlyname,omitempty"`
	Audience                    string  `json:"audience,omitempty"`
	Count                       float64 `json:"__count,omitempty"`
	Defaultauthenticationgroup  string  `json:"defaultauthenticationgroup,omitempty"`
	Digestmethod                string  `json:"digestmethod,omitempty"`
	Encryptassertion            string  `json:"encryptassertion,omitempty"`
	Encryptionalgorithm         string  `json:"encryptionalgorithm,omitempty"`
	Keytransportalg             string  `json:"keytransportalg,omitempty"`
	Logoutbinding               string  `json:"logoutbinding,omitempty"`
	Metadataimportstatus        string  `json:"metadataimportstatus,omitempty"`
	Metadatarefreshinterval     int     `json:"metadatarefreshinterval,omitempty"`
	Metadataurl                 string  `json:"metadataurl,omitempty"`
	Name                        string  `json:"name,omitempty"`
	Nameidexpr                  string  `json:"nameidexpr,omitempty"`
	Nameidformat                string  `json:"nameidformat,omitempty"`
	Nextgenapiresource          string  `json:"_nextgenapiresource,omitempty"`
	Rejectunsignedrequests      string  `json:"rejectunsignedrequests,omitempty"`
	Samlbinding                 string  `json:"samlbinding,omitempty"`
	Samlidpcertname             string  `json:"samlidpcertname,omitempty"`
	Samlissuername              string  `json:"samlissuername,omitempty"`
	Samlsigningcertversion      string  `json:"samlsigningcertversion,omitempty"`
	Samlspcertname              string  `json:"samlspcertname,omitempty"`
	Samlspcertversion           string  `json:"samlspcertversion,omitempty"`
	Sendpassword                string  `json:"sendpassword,omitempty"`
	Serviceproviderid           string  `json:"serviceproviderid,omitempty"`
	Signassertion               string  `json:"signassertion,omitempty"`
	Signaturealg                string  `json:"signaturealg,omitempty"`
	Signatureservice            string  `json:"signatureservice,omitempty"`
	Skewtime                    int     `json:"skewtime,omitempty"`
	Splogouturl                 string  `json:"splogouturl,omitempty"`
}

type AuthenticationldappolicyVpnglobalBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AuthenticationvserverAuthenticationsmartaccesspolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type Authenticationcertpolicy struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Reqaction          string  `json:"reqaction,omitempty"`
	Rule               string  `json:"rule,omitempty"`
}
