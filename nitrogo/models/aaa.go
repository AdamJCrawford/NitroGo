package models

// aaa configuration structs
type AaagroupVpnurlpolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupname              string `json:"groupname,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
}

type AaagroupVpnsessionpolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupname              string `json:"groupname,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
}

type AaagroupVpnurlBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupname              string `json:"groupname,omitempty"`
	Urlname                string `json:"urlname,omitempty"`
}

type Aaakcdaccount struct {
	Cacert             string  `json:"cacert,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Delegateduser      string  `json:"delegateduser,omitempty"`
	Enterpriserealm    string  `json:"enterpriserealm,omitempty"`
	Kcdaccount         string  `json:"kcdaccount,omitempty"`
	Kcdpassword        string  `json:"kcdpassword,omitempty"`
	Kcdspn             string  `json:"kcdspn,omitempty"`
	Keytab             string  `json:"keytab,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Principle          string  `json:"principle,omitempty"`
	Realmstr           string  `json:"realmstr,omitempty"`
	Saltexpression     string  `json:"saltexpression,omitempty"`
	Servicespn         string  `json:"servicespn,omitempty"`
	Usercert           string  `json:"usercert,omitempty"`
	Userrealm          string  `json:"userrealm,omitempty"`
}

type AaagroupAuditnslogpolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupname              string `json:"groupname,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
}

type AaaglobalBinding struct {
	AaaglobalAaapreauthenticationpolicyBinding    []interface{} `json:"aaaglobal_aaapreauthenticationpolicy_binding,omitempty"`
	AaaglobalAuthenticationnegotiateactionBinding []interface{} `json:"aaaglobal_authenticationnegotiateaction_binding,omitempty"`
}

type AaagroupVpnintranetapplicationBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupname              string `json:"groupname,omitempty"`
	Intranetapplication    string `json:"intranetapplication,omitempty"`
}

type AaauserVpntrafficpolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
	Username               string `json:"username,omitempty"`
}

type AaauserAuditsyslogpolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
	Username               string `json:"username,omitempty"`
}

type Aaassoprofile struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Password           string  `json:"password,omitempty"`
	Username           string  `json:"username,omitempty"`
}

type Aaapreauthenticationparameter struct {
	Builtin                 []string `json:"builtin,omitempty"`
	Deletefiles             string   `json:"deletefiles,omitempty"`
	Feature                 string   `json:"feature,omitempty"`
	Killprocess             string   `json:"killprocess,omitempty"`
	Nextgenapiresource      string   `json:"_nextgenapiresource,omitempty"`
	Preauthenticationaction string   `json:"preauthenticationaction,omitempty"`
	Rule                    string   `json:"rule,omitempty"`
}

type AaauserAuthorizationpolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
	Username               string `json:"username,omitempty"`
}

type AaauserIntranetipBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Intranetip             string `json:"intranetip,omitempty"`
	Netmask                string `json:"netmask,omitempty"`
	Username               string `json:"username,omitempty"`
}

type AaagroupIntranetipBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupname              string `json:"groupname,omitempty"`
	Intranetip             string `json:"intranetip,omitempty"`
	Netmask                string `json:"netmask,omitempty"`
}

type AaagroupVpntrafficpolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupname              string `json:"groupname,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
}

type AaagroupVpnsecureprivateaccessprofileBinding struct {
	Acttype                    int    `json:"acttype,omitempty"`
	Gotopriorityexpression     string `json:"gotopriorityexpression,omitempty"`
	Groupname                  string `json:"groupname,omitempty"`
	Secureprivateaccessprofile string `json:"secureprivateaccessprofile,omitempty"`
}

type AaagroupAuditsyslogpolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupname              string `json:"groupname,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
}

type AaauserVpnurlpolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
	Username               string `json:"username,omitempty"`
}

type AaauserVpnurlBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Urlname                string `json:"urlname,omitempty"`
	Username               string `json:"username,omitempty"`
}

type AaagroupIntranetip6Binding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupname              string `json:"groupname,omitempty"`
	Intranetip6            string `json:"intranetip6,omitempty"`
	Numaddr                int    `json:"numaddr,omitempty"`
}

type AaagroupAaauserBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupname              string `json:"groupname,omitempty"`
	Username               string `json:"username,omitempty"`
}

type AaapreauthenticationpolicyBinding struct {
	AaapreauthenticationpolicyAaaglobalBinding  []interface{} `json:"aaapreauthenticationpolicy_aaaglobal_binding,omitempty"`
	AaapreauthenticationpolicyVpnvserverBinding []interface{} `json:"aaapreauthenticationpolicy_vpnvserver_binding,omitempty"`
	Name                                        string        `json:"name,omitempty"`
}

type Aaacertparams struct {
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	Groupnamefield             string `json:"groupnamefield,omitempty"`
	Nextgenapiresource         string `json:"_nextgenapiresource,omitempty"`
	Twofactor                  string `json:"twofactor,omitempty"`
	Usernamefield              string `json:"usernamefield,omitempty"`
}

type AaauserVpnsessionpolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
	Username               string `json:"username,omitempty"`
}

type Aaaotpparameter struct {
	Encryption         string `json:"encryption,omitempty"`
	Maxotpdevices      int    `json:"maxotpdevices,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Aaaradiusparams struct {
	Accounting                 string   `json:"accounting,omitempty"`
	Authentication             string   `json:"authentication,omitempty"`
	Authservretry              int      `json:"authservretry,omitempty"`
	Authtimeout                int      `json:"authtimeout,omitempty"`
	Builtin                    []string `json:"builtin,omitempty"`
	Callingstationid           string   `json:"callingstationid,omitempty"`
	Defaultauthenticationgroup string   `json:"defaultauthenticationgroup,omitempty"`
	Feature                    string   `json:"feature,omitempty"`
	Groupauthname              string   `json:"groupauthname,omitempty"`
	Ipaddress                  string   `json:"ipaddress,omitempty"`
	Ipattributetype            int      `json:"ipattributetype,omitempty"`
	Ipvendorid                 int      `json:"ipvendorid,omitempty"`
	Messageauthenticator       string   `json:"messageauthenticator,omitempty"`
	Nextgenapiresource         string   `json:"_nextgenapiresource,omitempty"`
	Passencoding               string   `json:"passencoding,omitempty"`
	Pwdattributetype           int      `json:"pwdattributetype,omitempty"`
	Pwdvendorid                int      `json:"pwdvendorid,omitempty"`
	Radattributetype           int      `json:"radattributetype,omitempty"`
	Radgroupseparator          string   `json:"radgroupseparator,omitempty"`
	Radgroupsprefix            string   `json:"radgroupsprefix,omitempty"`
	Radkey                     string   `json:"radkey,omitempty"`
	Radnasid                   string   `json:"radnasid,omitempty"`
	Radnasip                   string   `json:"radnasip,omitempty"`
	Radvendorid                int      `json:"radvendorid,omitempty"`
	Serverip                   string   `json:"serverip,omitempty"`
	Serverport                 int      `json:"serverport,omitempty"`
	Tunnelendpointclientip     string   `json:"tunnelendpointclientip,omitempty"`
}

type AaapreauthenticationpolicyAaaglobalBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type AaaglobalAuthenticationnegotiateactionBinding struct {
	Windowsprofile string `json:"windowsprofile,omitempty"`
}

type AaaglobalAaapreauthenticationpolicyBinding struct {
	Bindpolicytype int      `json:"bindpolicytype,omitempty"`
	Builtin        []string `json:"builtin,omitempty"`
	Policy         string   `json:"policy,omitempty"`
	Priority       int      `json:"priority,omitempty"`
}

type AaagroupAuthorizationpolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupname              string `json:"groupname,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
}

type AaauserVpnintranetapplicationBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Intranetapplication    string `json:"intranetapplication,omitempty"`
	Username               string `json:"username,omitempty"`
}

type Aaaparameter struct {
	Aaadloglevel               string   `json:"aaadloglevel,omitempty"`
	Aaadnatip                  string   `json:"aaadnatip,omitempty"`
	Aaasessionloglevel         string   `json:"aaasessionloglevel,omitempty"`
	Apitokencache              string   `json:"apitokencache,omitempty"`
	Builtin                    []string `json:"builtin,omitempty"`
	Classicendpoints           string   `json:"classicendpoints,omitempty"`
	Defaultauthtype            string   `json:"defaultauthtype,omitempty"`
	Defaultcspheader           string   `json:"defaultcspheader,omitempty"`
	Dynaddr                    string   `json:"dynaddr,omitempty"`
	Enableenhancedauthfeedback string   `json:"enableenhancedauthfeedback,omitempty"`
	Enablesessionstickiness    string   `json:"enablesessionstickiness,omitempty"`
	Enablestaticpagecaching    string   `json:"enablestaticpagecaching,omitempty"`
	Enhancedepa                string   `json:"enhancedepa,omitempty"`
	Failedlogintimeout         int      `json:"failedlogintimeout,omitempty"`
	Feature                    string   `json:"feature,omitempty"`
	Ftmode                     string   `json:"ftmode,omitempty"`
	Httponlycookie             string   `json:"httponlycookie,omitempty"`
	Loginencryption            string   `json:"loginencryption,omitempty"`
	Maxaaausers                int      `json:"maxaaausers,omitempty"`
	Maxkbquestions             int      `json:"maxkbquestions,omitempty"`
	Maxloginattempts           int      `json:"maxloginattempts,omitempty"`
	Maxsamldeflatesize         int      `json:"maxsamldeflatesize,omitempty"`
	Nextgenapiresource         string   `json:"_nextgenapiresource,omitempty"`
	Persistentloginattempts    string   `json:"persistentloginattempts,omitempty"`
	Pwdexpirynotificationdays  int      `json:"pwdexpirynotificationdays,omitempty"`
	Samesite                   string   `json:"samesite,omitempty"`
	Securityinsights           string   `json:"securityinsights,omitempty"`
	Tokenintrospectioninterval int      `json:"tokenintrospectioninterval,omitempty"`
	Wafprotection              []string `json:"wafprotection,omitempty"`
	Webviewendpoints           string   `json:"webviewendpoints,omitempty"`
}

type AaauserVpnsecureprivateaccessprofileBinding struct {
	Acttype                    int    `json:"acttype,omitempty"`
	Gotopriorityexpression     string `json:"gotopriorityexpression,omitempty"`
	Secureprivateaccessprofile string `json:"secureprivateaccessprofile,omitempty"`
	Username                   string `json:"username,omitempty"`
}

type AaauserAuditnslogpolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
	Username               string `json:"username,omitempty"`
}

type AaauserBinding struct {
	AaauserAaagroupBinding                      []interface{} `json:"aaauser_aaagroup_binding,omitempty"`
	AaauserAuditnslogpolicyBinding              []interface{} `json:"aaauser_auditnslogpolicy_binding,omitempty"`
	AaauserAuditsyslogpolicyBinding             []interface{} `json:"aaauser_auditsyslogpolicy_binding,omitempty"`
	AaauserAuthorizationpolicyBinding           []interface{} `json:"aaauser_authorizationpolicy_binding,omitempty"`
	AaauserIntranetip6Binding                   []interface{} `json:"aaauser_intranetip6_binding,omitempty"`
	AaauserIntranetipBinding                    []interface{} `json:"aaauser_intranetip_binding,omitempty"`
	AaauserTmsessionpolicyBinding               []interface{} `json:"aaauser_tmsessionpolicy_binding,omitempty"`
	AaauserVpnintranetapplicationBinding        []interface{} `json:"aaauser_vpnintranetapplication_binding,omitempty"`
	AaauserVpnsecureprivateaccessprofileBinding []interface{} `json:"aaauser_vpnsecureprivateaccessprofile_binding,omitempty"`
	AaauserVpnsessionpolicyBinding              []interface{} `json:"aaauser_vpnsessionpolicy_binding,omitempty"`
	AaauserVpntrafficpolicyBinding              []interface{} `json:"aaauser_vpntrafficpolicy_binding,omitempty"`
	AaauserVpnurlBinding                        []interface{} `json:"aaauser_vpnurl_binding,omitempty"`
	AaauserVpnurlpolicyBinding                  []interface{} `json:"aaauser_vpnurlpolicy_binding,omitempty"`
	Username                                    string        `json:"username,omitempty"`
}

type AaagroupTmsessionpolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupname              string `json:"groupname,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
}

type Aaapreauthenticationpolicy struct {
	Builtin            []string `json:"builtin,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	Name               string   `json:"name,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Reqaction          string   `json:"reqaction,omitempty"`
	Rule               string   `json:"rule,omitempty"`
}

type Aaatacacsparams struct {
	Accounting                 string   `json:"accounting,omitempty"`
	Auditfailedcmds            string   `json:"auditfailedcmds,omitempty"`
	Authorization              string   `json:"authorization,omitempty"`
	Authtimeout                int      `json:"authtimeout,omitempty"`
	Builtin                    []string `json:"builtin,omitempty"`
	Defaultauthenticationgroup string   `json:"defaultauthenticationgroup,omitempty"`
	Feature                    string   `json:"feature,omitempty"`
	Groupattrname              string   `json:"groupattrname,omitempty"`
	Nextgenapiresource         string   `json:"_nextgenapiresource,omitempty"`
	Serverip                   string   `json:"serverip,omitempty"`
	Serverport                 int      `json:"serverport,omitempty"`
	Tacacssecret               string   `json:"tacacssecret,omitempty"`
}

type Aaagroup struct {
	Count              float64 `json:"__count,omitempty"`
	Groupname          string  `json:"groupname,omitempty"`
	Loggedin           bool    `json:"loggedin,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Weight             int     `json:"weight,omitempty"`
}

type AaauserTmsessionpolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
	Username               string `json:"username,omitempty"`
}

type AaagroupBinding struct {
	AaagroupAaauserBinding                       []interface{} `json:"aaagroup_aaauser_binding,omitempty"`
	AaagroupAuditnslogpolicyBinding              []interface{} `json:"aaagroup_auditnslogpolicy_binding,omitempty"`
	AaagroupAuditsyslogpolicyBinding             []interface{} `json:"aaagroup_auditsyslogpolicy_binding,omitempty"`
	AaagroupAuthorizationpolicyBinding           []interface{} `json:"aaagroup_authorizationpolicy_binding,omitempty"`
	AaagroupIntranetip6Binding                   []interface{} `json:"aaagroup_intranetip6_binding,omitempty"`
	AaagroupIntranetipBinding                    []interface{} `json:"aaagroup_intranetip_binding,omitempty"`
	AaagroupTmsessionpolicyBinding               []interface{} `json:"aaagroup_tmsessionpolicy_binding,omitempty"`
	AaagroupVpnintranetapplicationBinding        []interface{} `json:"aaagroup_vpnintranetapplication_binding,omitempty"`
	AaagroupVpnsecureprivateaccessprofileBinding []interface{} `json:"aaagroup_vpnsecureprivateaccessprofile_binding,omitempty"`
	AaagroupVpnsessionpolicyBinding              []interface{} `json:"aaagroup_vpnsessionpolicy_binding,omitempty"`
	AaagroupVpntrafficpolicyBinding              []interface{} `json:"aaagroup_vpntrafficpolicy_binding,omitempty"`
	AaagroupVpnurlBinding                        []interface{} `json:"aaagroup_vpnurl_binding,omitempty"`
	AaagroupVpnurlpolicyBinding                  []interface{} `json:"aaagroup_vpnurlpolicy_binding,omitempty"`
	Groupname                                    string        `json:"groupname,omitempty"`
}

type Aaasession struct {
	All                bool    `json:"all,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Destip             string  `json:"destip,omitempty"`
	Destport           int     `json:"destport,omitempty"`
	Groupname          string  `json:"groupname,omitempty"`
	Iip                string  `json:"iip,omitempty"`
	Intranetip         string  `json:"intranetip,omitempty"`
	Intranetip6        string  `json:"intranetip6,omitempty"`
	Ipaddress          string  `json:"ipaddress,omitempty"`
	Netmask            string  `json:"netmask,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Nodeid             int     `json:"nodeid,omitempty"`
	Peid               int     `json:"peid,omitempty"`
	Port               int     `json:"port,omitempty"`
	Privateip          string  `json:"privateip,omitempty"`
	Privateport        int     `json:"privateport,omitempty"`
	Publicip           string  `json:"publicip,omitempty"`
	Publicport         int     `json:"publicport,omitempty"`
	Sessionkey         string  `json:"sessionkey,omitempty"`
	Username           string  `json:"username,omitempty"`
}

type Aaapreauthenticationaction struct {
	Builtin                 []string `json:"builtin,omitempty"`
	Count                   float64  `json:"__count,omitempty"`
	Defaultepagroup         string   `json:"defaultepagroup,omitempty"`
	Deletefiles             string   `json:"deletefiles,omitempty"`
	Feature                 string   `json:"feature,omitempty"`
	Killprocess             string   `json:"killprocess,omitempty"`
	Name                    string   `json:"name,omitempty"`
	Nextgenapiresource      string   `json:"_nextgenapiresource,omitempty"`
	Preauthenticationaction string   `json:"preauthenticationaction,omitempty"`
}

type AaauserAaagroupBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupname              string `json:"groupname,omitempty"`
	Username               string `json:"username,omitempty"`
}

type Aaaldapparams struct {
	Authtimeout                int      `json:"authtimeout,omitempty"`
	Builtin                    []string `json:"builtin,omitempty"`
	Defaultauthenticationgroup string   `json:"defaultauthenticationgroup,omitempty"`
	Feature                    string   `json:"feature,omitempty"`
	Groupattrname              string   `json:"groupattrname,omitempty"`
	Groupauthname              string   `json:"groupauthname,omitempty"`
	Groupnameidentifier        string   `json:"groupnameidentifier,omitempty"`
	Groupsearchattribute       string   `json:"groupsearchattribute,omitempty"`
	Groupsearchfilter          string   `json:"groupsearchfilter,omitempty"`
	Groupsearchsubattribute    string   `json:"groupsearchsubattribute,omitempty"`
	Ldapbase                   string   `json:"ldapbase,omitempty"`
	Ldapbinddn                 string   `json:"ldapbinddn,omitempty"`
	Ldapbinddnpassword         string   `json:"ldapbinddnpassword,omitempty"`
	Ldaploginname              string   `json:"ldaploginname,omitempty"`
	Maxnestinglevel            int      `json:"maxnestinglevel,omitempty"`
	Nestedgroupextraction      string   `json:"nestedgroupextraction,omitempty"`
	Nextgenapiresource         string   `json:"_nextgenapiresource,omitempty"`
	Passwdchange               string   `json:"passwdchange,omitempty"`
	Searchfilter               string   `json:"searchfilter,omitempty"`
	Sectype                    string   `json:"sectype,omitempty"`
	Serverip                   string   `json:"serverip,omitempty"`
	Serverport                 int      `json:"serverport,omitempty"`
	Ssonameattribute           string   `json:"ssonameattribute,omitempty"`
	Subattributename           string   `json:"subattributename,omitempty"`
	Svrtype                    string   `json:"svrtype,omitempty"`
}

type Aaauser struct {
	Count              float64 `json:"__count,omitempty"`
	Loggedin           bool    `json:"loggedin,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Password           string  `json:"password,omitempty"`
	Username           string  `json:"username,omitempty"`
}

type AaauserIntranetip6Binding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Intranetip6            string `json:"intranetip6,omitempty"`
	Numaddr                int    `json:"numaddr,omitempty"`
	Username               string `json:"username,omitempty"`
}

type AaapreauthenticationpolicyVpnvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}
