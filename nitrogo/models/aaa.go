// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Aaaglobalauthenticationnegotiateactionbinding struct {
	Windowsprofile string `json:"windowsprofile,omitempty"`
}

type Aaagroupintranetapplicationbinding struct {
	Intranetapplication    string `json:"intranetapplication,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Groupname              string `json:"groupname,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Aaaldapparams struct {
	Serverip                   string `json:"serverip,omitempty"`
	Serverport                 int    `json:"serverport,omitempty"`
	Authtimeout                int    `json:"authtimeout,omitempty"`
	Ldapbase                   string `json:"ldapbase,omitempty"`
	Ldapbinddn                 string `json:"ldapbinddn,omitempty"`
	Ldapbinddnpassword         string `json:"ldapbinddnpassword,omitempty"`
	Ldaploginname              string `json:"ldaploginname,omitempty"`
	Searchfilter               string `json:"searchfilter,omitempty"`
	Groupattrname              string `json:"groupattrname,omitempty"`
	Subattributename           string `json:"subattributename,omitempty"`
	Sectype                    string `json:"sectype,omitempty"`
	Svrtype                    string `json:"svrtype,omitempty"`
	Ssonameattribute           string `json:"ssonameattribute,omitempty"`
	Passwdchange               string `json:"passwdchange,omitempty"`
	Nestedgroupextraction      string `json:"nestedgroupextraction,omitempty"`
	Maxnestinglevel            int    `json:"maxnestinglevel,omitempty"`
	Groupnameidentifier        string `json:"groupnameidentifier,omitempty"`
	Groupsearchattribute       string `json:"groupsearchattribute,omitempty"`
	Groupsearchsubattribute    string `json:"groupsearchsubattribute,omitempty"`
	Groupsearchfilter          string `json:"groupsearchfilter,omitempty"`
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	Groupauthname              string `json:"groupauthname,omitempty"`
	Builtin                    string `json:"builtin,omitempty"`
	Feature                    string `json:"feature,omitempty"`
	Nextgenapiresource         string `json:"_nextgenapiresource,omitempty"`
}

type Aaaotpparameter struct {
	Encryption         string `json:"encryption,omitempty"`
	Maxotpdevices      int    `json:"maxotpdevices,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Aaasession struct {
	Username           string `json:"username,omitempty"`
	Groupname          string `json:"groupname,omitempty"`
	Iip                string `json:"iip,omitempty"`
	Netmask            string `json:"netmask,omitempty"`
	Sessionkey         string `json:"sessionkey,omitempty"`
	Nodeid             int    `json:"nodeid,omitempty"`
	All                bool   `json:"all,omitempty"`
	Publicip           string `json:"publicip,omitempty"`
	Publicport         string `json:"publicport,omitempty"`
	Ipaddress          string `json:"ipaddress,omitempty"`
	Port               string `json:"port,omitempty"`
	Privateip          string `json:"privateip,omitempty"`
	Privateport        string `json:"privateport,omitempty"`
	Destip             string `json:"destip,omitempty"`
	Destport           string `json:"destport,omitempty"`
	Intranetip         string `json:"intranetip,omitempty"`
	Intranetip6        string `json:"intranetip6,omitempty"`
	Peid               string `json:"peid,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Aaauserauditsyslogpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Username               string `json:"username,omitempty"`
	Type                   string `json:"type,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Aaagroupauthorizationpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Type                   string `json:"type,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupname              string `json:"groupname,omitempty"`
}

type Aaagroupsessionpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupname              string `json:"groupname,omitempty"`
	Type                   string `json:"type,omitempty"`
}

type Aaaparameter struct {
	Enablestaticpagecaching    string   `json:"enablestaticpagecaching,omitempty"`
	Enableenhancedauthfeedback string   `json:"enableenhancedauthfeedback,omitempty"`
	Defaultauthtype            string   `json:"defaultauthtype,omitempty"`
	Maxaaausers                int      `json:"maxaaausers,omitempty"`
	Maxloginattempts           int      `json:"maxloginattempts,omitempty"`
	Failedlogintimeout         int      `json:"failedlogintimeout,omitempty"`
	Aaadnatip                  string   `json:"aaadnatip,omitempty"`
	Enablesessionstickiness    string   `json:"enablesessionstickiness,omitempty"`
	Aaasessionloglevel         string   `json:"aaasessionloglevel,omitempty"`
	Aaadloglevel               string   `json:"aaadloglevel,omitempty"`
	Dynaddr                    string   `json:"dynaddr,omitempty"`
	Ftmode                     string   `json:"ftmode,omitempty"`
	Maxsamldeflatesize         int      `json:"maxsamldeflatesize,omitempty"`
	Persistentloginattempts    string   `json:"persistentloginattempts,omitempty"`
	Pwdexpirynotificationdays  int      `json:"pwdexpirynotificationdays,omitempty"`
	Maxkbquestions             int      `json:"maxkbquestions,omitempty"`
	Loginencryption            string   `json:"loginencryption,omitempty"`
	Samesite                   string   `json:"samesite,omitempty"`
	Apitokencache              string   `json:"apitokencache,omitempty"`
	Tokenintrospectioninterval int      `json:"tokenintrospectioninterval,omitempty"`
	Defaultcspheader           string   `json:"defaultcspheader,omitempty"`
	Httponlycookie             string   `json:"httponlycookie,omitempty"`
	Enhancedepa                string   `json:"enhancedepa,omitempty"`
	Wafprotection              []string `json:"wafprotection,omitempty"`
	Securityinsights           string   `json:"securityinsights,omitempty"`
	Builtin                    string   `json:"builtin,omitempty"`
	Feature                    string   `json:"feature,omitempty"`
	Nextgenapiresource         string   `json:"_nextgenapiresource,omitempty"`
}

type Aaauserbinding struct {
	Username string `json:"username,omitempty"`
}

type Aaausersyslogpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Username               string `json:"username,omitempty"`
	Type                   string `json:"type,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Aaausertrafficpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Username               string `json:"username,omitempty"`
	Type                   string `json:"type,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Aaauservpnurlbinding struct {
	Urlname                string `json:"urlname,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Username               string `json:"username,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Aaagroupuserbinding struct {
	Username               string `json:"username,omitempty"`
	Groupname              string `json:"groupname,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Aaapreauthenticationpolicyaaaglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Aaaglobalbinding struct {
}

type Aaagroupintranetip6binding struct {
	Intranetip6            string `json:"intranetip6,omitempty"`
	Numaddr                int    `json:"numaddr,omitempty"`
	Groupname              string `json:"groupname,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Aaagrouppolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Type                   string `json:"type,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupname              string `json:"groupname,omitempty"`
}

type Aaagroupsyslogpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Groupname              string `json:"groupname,omitempty"`
	Type                   string `json:"type,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Aaagroupvpnurlbinding struct {
	Urlname                string `json:"urlname,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Groupname              string `json:"groupname,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Aaagroupvpnurlpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Groupname              string `json:"groupname,omitempty"`
	Type                   string `json:"type,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Aaakcdaccount struct {
	Kcdaccount         string `json:"kcdaccount,omitempty"`
	Keytab             string `json:"keytab,omitempty"`
	Realmstr           string `json:"realmstr,omitempty"`
	Delegateduser      string `json:"delegateduser,omitempty"`
	Kcdpassword        string `json:"kcdpassword,omitempty"`
	Usercert           string `json:"usercert,omitempty"`
	Cacert             string `json:"cacert,omitempty"`
	Userrealm          string `json:"userrealm,omitempty"`
	Enterpriserealm    string `json:"enterpriserealm,omitempty"`
	Servicespn         string `json:"servicespn,omitempty"`
	Principle          string `json:"principle,omitempty"`
	Kcdspn             string `json:"kcdspn,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Aaatacacsparams struct {
	Serverip                   string `json:"serverip,omitempty"`
	Serverport                 int    `json:"serverport,omitempty"`
	Authtimeout                int    `json:"authtimeout,omitempty"`
	Tacacssecret               string `json:"tacacssecret,omitempty"`
	Authorization              string `json:"authorization,omitempty"`
	Accounting                 string `json:"accounting,omitempty"`
	Auditfailedcmds            string `json:"auditfailedcmds,omitempty"`
	Groupattrname              string `json:"groupattrname,omitempty"`
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	Builtin                    string `json:"builtin,omitempty"`
	Feature                    string `json:"feature,omitempty"`
	Nextgenapiresource         string `json:"_nextgenapiresource,omitempty"`
}

type Aaauservpntrafficpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Username               string `json:"username,omitempty"`
	Type                   string `json:"type,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Aaacertparams struct {
	Usernamefield              string `json:"usernamefield,omitempty"`
	Groupnamefield             string `json:"groupnamefield,omitempty"`
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	Twofactor                  string `json:"twofactor,omitempty"`
	Nextgenapiresource         string `json:"_nextgenapiresource,omitempty"`
}

type Aaaglobalaaapreauthenticationpolicybinding struct {
	Policy         string   `json:"policy,omitempty"`
	Priority       int      `json:"priority,omitempty"`
	Bindpolicytype int      `json:"bindpolicytype,omitempty"`
	Builtin        []string `json:"builtin,omitempty"`
}

type Aaagroupauditsyslogpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Groupname              string `json:"groupname,omitempty"`
	Type                   string `json:"type,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Aaauservpnintranetapplicationbinding struct {
	Intranetapplication    string `json:"intranetapplication,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Username               string `json:"username,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Aaauservpnsessionpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Username               string `json:"username,omitempty"`
	Type                   string `json:"type,omitempty"`
}

type Aaagroup struct {
	Groupname          string `json:"groupname,omitempty"`
	Weight             int    `json:"weight,omitempty"`
	Loggedin           bool   `json:"loggedin,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Aaauserpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Type                   string `json:"type,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Username               string `json:"username,omitempty"`
}

type Aaauservpnurlpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Username               string `json:"username,omitempty"`
	Type                   string `json:"type,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Aaaglobalnegotiateactionbinding struct {
	Windowsprofile string `json:"windowsprofile,omitempty"`
}

type Aaagroupaaauserbinding struct {
	Username               string `json:"username,omitempty"`
	Groupname              string `json:"groupname,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Aaapreauthenticationparameter struct {
	Preauthenticationaction string `json:"preauthenticationaction,omitempty"`
	Rule                    string `json:"rule,omitempty"`
	Killprocess             string `json:"killprocess,omitempty"`
	Deletefiles             string `json:"deletefiles,omitempty"`
	Builtin                 string `json:"builtin,omitempty"`
	Feature                 string `json:"feature,omitempty"`
	Nextgenapiresource      string `json:"_nextgenapiresource,omitempty"`
}

type Aaassoprofile struct {
	Name               string `json:"name,omitempty"`
	Username           string `json:"username,omitempty"`
	Password           string `json:"password,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Aaauser struct {
	Username           string `json:"username,omitempty"`
	Password           string `json:"password,omitempty"`
	Loggedin           bool   `json:"loggedin,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Aaauserintranetip6binding struct {
	Intranetip6            string `json:"intranetip6,omitempty"`
	Numaddr                int    `json:"numaddr,omitempty"`
	Username               string `json:"username,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Aaagrouptmsessionpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupname              string `json:"groupname,omitempty"`
	Type                   string `json:"type,omitempty"`
}

type Aaauserintranetapplicationbinding struct {
	Intranetapplication    string `json:"intranetapplication,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Username               string `json:"username,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Aaauserurlbinding struct {
	Urlname                string `json:"urlname,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Username               string `json:"username,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Aaagroupauditnslogpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Groupname              string `json:"groupname,omitempty"`
	Type                   string `json:"type,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Aaapreauthenticationaction struct {
	Name                    string `json:"name,omitempty"`
	Preauthenticationaction string `json:"preauthenticationaction,omitempty"`
	Killprocess             string `json:"killprocess,omitempty"`
	Deletefiles             string `json:"deletefiles,omitempty"`
	Defaultepagroup         string `json:"defaultepagroup,omitempty"`
	Builtin                 string `json:"builtin,omitempty"`
	Feature                 string `json:"feature,omitempty"`
	Nextgenapiresource      string `json:"_nextgenapiresource,omitempty"`
}

type Aaaradiusparams struct {
	Serverip                   string `json:"serverip,omitempty"`
	Serverport                 int    `json:"serverport,omitempty"`
	Authtimeout                int    `json:"authtimeout,omitempty"`
	Radkey                     string `json:"radkey,omitempty"`
	Radnasip                   string `json:"radnasip,omitempty"`
	Radnasid                   string `json:"radnasid,omitempty"`
	Radvendorid                int    `json:"radvendorid,omitempty"`
	Radattributetype           int    `json:"radattributetype,omitempty"`
	Radgroupsprefix            string `json:"radgroupsprefix,omitempty"`
	Radgroupseparator          string `json:"radgroupseparator,omitempty"`
	Passencoding               string `json:"passencoding,omitempty"`
	Ipvendorid                 int    `json:"ipvendorid,omitempty"`
	Ipattributetype            int    `json:"ipattributetype,omitempty"`
	Accounting                 string `json:"accounting,omitempty"`
	Pwdvendorid                int    `json:"pwdvendorid,omitempty"`
	Pwdattributetype           int    `json:"pwdattributetype,omitempty"`
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	Callingstationid           string `json:"callingstationid,omitempty"`
	Authservretry              int    `json:"authservretry,omitempty"`
	Authentication             string `json:"authentication,omitempty"`
	Tunnelendpointclientip     string `json:"tunnelendpointclientip,omitempty"`
	Messageauthenticator       string `json:"messageauthenticator,omitempty"`
	Groupauthname              string `json:"groupauthname,omitempty"`
	Ipaddress                  string `json:"ipaddress,omitempty"`
	Builtin                    string `json:"builtin,omitempty"`
	Feature                    string `json:"feature,omitempty"`
	Nextgenapiresource         string `json:"_nextgenapiresource,omitempty"`
}

type Aaaglobalpreauthenticationpolicybinding struct {
	Policy         string   `json:"policy,omitempty"`
	Priority       uint32   `json:"priority,omitempty"`
	Bindpolicytype uint32   `json:"bindpolicytype,omitempty"`
	Builtin        []string `json:"builtin,omitempty"`
}

type Aaagroupbinding struct {
	Groupname string `json:"groupname,omitempty"`
}

type Aaagroupurlbinding struct {
	Urlname                string `json:"urlname,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Groupname              string `json:"groupname,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Aaagroupurlpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Groupname              string `json:"groupname,omitempty"`
	Type                   string `json:"type,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Aaapreauthenticationpolicy struct {
	Name               string `json:"name,omitempty"`
	Rule               string `json:"rule,omitempty"`
	Reqaction          string `json:"reqaction,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Builtin            string `json:"builtin,omitempty"`
	Feature            string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Aaauseraaagroupbinding struct {
	Groupname              string `json:"groupname,omitempty"`
	Username               string `json:"username,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Aaagrouptrafficpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Groupname              string `json:"groupname,omitempty"`
	Type                   string `json:"type,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Aaapreauthenticationpolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Aaagroupnslogpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Groupname              string `json:"groupname,omitempty"`
	Type                   string `json:"type,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Aaagroupvpntrafficpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Groupname              string `json:"groupname,omitempty"`
	Type                   string `json:"type,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Aaapreauthenticationpolicyvpnvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Aaauserauthorizationpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Type                   string `json:"type,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Username               string `json:"username,omitempty"`
}

type Aaausersessionpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Username               string `json:"username,omitempty"`
	Type                   string `json:"type,omitempty"`
}

type Aaagroupvpnsessionpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupname              string `json:"groupname,omitempty"`
	Type                   string `json:"type,omitempty"`
}

type Aaapreauthenticationpolicyglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Aaapreauthenticationpolicyvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Aaauserauditnslogpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Username               string `json:"username,omitempty"`
	Type                   string `json:"type,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Aaausernslogpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Username               string `json:"username,omitempty"`
	Type                   string `json:"type,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Aaausertmsessionpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Username               string `json:"username,omitempty"`
	Type                   string `json:"type,omitempty"`
}

type Aaauserurlpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Username               string `json:"username,omitempty"`
	Type                   string `json:"type,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Aaauserintranetipbinding struct {
	Intranetip             string `json:"intranetip,omitempty"`
	Netmask                string `json:"netmask,omitempty"`
	Username               string `json:"username,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Aaagroupintranetipbinding struct {
	Intranetip             string `json:"intranetip,omitempty"`
	Netmask                string `json:"netmask,omitempty"`
	Groupname              string `json:"groupname,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Aaagroupvpnintranetapplicationbinding struct {
	Intranetapplication    string `json:"intranetapplication,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Groupname              string `json:"groupname,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Aaausergroupbinding struct {
	Groupname              string `json:"groupname,omitempty"`
	Username               string `json:"username,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}
