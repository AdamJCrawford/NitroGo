// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Authenticationwebauthpolicyvpnglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationwebauthpolicyvpnvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationcertpolicyvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationvserverrewritepolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
}

type Authenticationvserversessionpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Authenticationcertpolicyvpnvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationldappolicysystemglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationloginschemapolicyauthenticationvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Authenticationsamlidppolicyvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationtacacspolicyvpnvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationradiuspolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Authenticationvserverauthenticationldappolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Authenticationldappolicy struct {
	Name               string `json:"name,omitempty"`
	Rule               string `json:"rule,omitempty"`
	Reqaction          string `json:"reqaction,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Authenticationlocalpolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Authenticationloginschemapolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Authenticationpolicypolicylabelbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Activepolicy           uint32 `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Authenticationtacacsaction struct {
	Name                       string `json:"name,omitempty"`
	Serverip                   string `json:"serverip,omitempty"`
	Serverport                 int    `json:"serverport,omitempty"`
	Authtimeout                int    `json:"authtimeout,omitempty"`
	Tacacssecret               string `json:"tacacssecret,omitempty"`
	Authorization              string `json:"authorization,omitempty"`
	Accounting                 string `json:"accounting,omitempty"`
	Auditfailedcmds            string `json:"auditfailedcmds,omitempty"`
	Groupattrname              string `json:"groupattrname,omitempty"`
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	Attribute1                 string `json:"attribute1,omitempty"`
	Attribute2                 string `json:"attribute2,omitempty"`
	Attribute3                 string `json:"attribute3,omitempty"`
	Attribute4                 string `json:"attribute4,omitempty"`
	Attribute5                 string `json:"attribute5,omitempty"`
	Attribute6                 string `json:"attribute6,omitempty"`
	Attribute7                 string `json:"attribute7,omitempty"`
	Attribute8                 string `json:"attribute8,omitempty"`
	Attribute9                 string `json:"attribute9,omitempty"`
	Attribute10                string `json:"attribute10,omitempty"`
	Attribute11                string `json:"attribute11,omitempty"`
	Attribute12                string `json:"attribute12,omitempty"`
	Attribute13                string `json:"attribute13,omitempty"`
	Attribute14                string `json:"attribute14,omitempty"`
	Attribute15                string `json:"attribute15,omitempty"`
	Attribute16                string `json:"attribute16,omitempty"`
	Attributes                 string `json:"attributes,omitempty"`
	Success                    string `json:"success,omitempty"`
	Failure                    string `json:"failure,omitempty"`
	Nextgenapiresource         string `json:"_nextgenapiresource,omitempty"`
}

type Authenticationtacacspolicyauthenticationvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationvserversamlidppolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Authenticationldapaction struct {
	Name                       string `json:"name,omitempty"`
	Serverip                   string `json:"serverip,omitempty"`
	Servername                 string `json:"servername,omitempty"`
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
	Authentication             string `json:"authentication,omitempty"`
	Requireuser                string `json:"requireuser,omitempty"`
	Passwdchange               string `json:"passwdchange,omitempty"`
	Nestedgroupextraction      string `json:"nestedgroupextraction,omitempty"`
	Maxnestinglevel            int    `json:"maxnestinglevel,omitempty"`
	Followreferrals            string `json:"followreferrals,omitempty"`
	Maxldapreferrals           int    `json:"maxldapreferrals,omitempty"`
	Referraldnslookup          string `json:"referraldnslookup,omitempty"`
	Mssrvrecordlocation        string `json:"mssrvrecordlocation,omitempty"`
	Validateservercert         string `json:"validateservercert,omitempty"`
	Ldaphostname               string `json:"ldaphostname,omitempty"`
	Groupnameidentifier        string `json:"groupnameidentifier,omitempty"`
	Groupsearchattribute       string `json:"groupsearchattribute,omitempty"`
	Groupsearchsubattribute    string `json:"groupsearchsubattribute,omitempty"`
	Groupsearchfilter          string `json:"groupsearchfilter,omitempty"`
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	Attribute1                 string `json:"attribute1,omitempty"`
	Attribute2                 string `json:"attribute2,omitempty"`
	Attribute3                 string `json:"attribute3,omitempty"`
	Attribute4                 string `json:"attribute4,omitempty"`
	Attribute5                 string `json:"attribute5,omitempty"`
	Attribute6                 string `json:"attribute6,omitempty"`
	Attribute7                 string `json:"attribute7,omitempty"`
	Attribute8                 string `json:"attribute8,omitempty"`
	Attribute9                 string `json:"attribute9,omitempty"`
	Attribute10                string `json:"attribute10,omitempty"`
	Attribute11                string `json:"attribute11,omitempty"`
	Attribute12                string `json:"attribute12,omitempty"`
	Attribute13                string `json:"attribute13,omitempty"`
	Attribute14                string `json:"attribute14,omitempty"`
	Attribute15                string `json:"attribute15,omitempty"`
	Attribute16                string `json:"attribute16,omitempty"`
	Attributes                 string `json:"attributes,omitempty"`
	Sshpublickey               string `json:"sshpublickey,omitempty"`
	Pushservice                string `json:"pushservice,omitempty"`
	Otpsecret                  string `json:"otpsecret,omitempty"`
	Email                      string `json:"email,omitempty"`
	Kbattribute                string `json:"kbattribute,omitempty"`
	Alternateemailattr         string `json:"alternateemailattr,omitempty"`
	Cloudattributes            string `json:"cloudattributes,omitempty"`
	Success                    string `json:"success,omitempty"`
	Failure                    string `json:"failure,omitempty"`
	Nextgenapiresource         string `json:"_nextgenapiresource,omitempty"`
}

type Authenticationldappolicyvpnvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationlocalpolicyvpnvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationoauthidppolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Authenticationtacacspolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Authenticationcaptchaaction struct {
	Name                       string `json:"name,omitempty"`
	Serverurl                  string `json:"serverurl,omitempty"`
	Secretkey                  string `json:"secretkey,omitempty"`
	Sitekey                    string `json:"sitekey,omitempty"`
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	Scorethreshold             int    `json:"scorethreshold,omitempty"`
	Nextgenapiresource         string `json:"_nextgenapiresource,omitempty"`
}

type Authenticationvserverauthenticationloginschemapolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Authenticationpolicyauthenticationvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Authenticationvserverloginschemapolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Authenticationcertpolicyauthenticationvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationoauthidpprofile struct {
	Name                       string `json:"name,omitempty"`
	Clientid                   string `json:"clientid,omitempty"`
	Clientsecret               string `json:"clientsecret,omitempty"`
	Redirecturl                string `json:"redirecturl,omitempty"`
	Issuer                     string `json:"issuer,omitempty"`
	Configservice              string `json:"configservice,omitempty"`
	Audience                   string `json:"audience,omitempty"`
	Skewtime                   int    `json:"skewtime,omitempty"`
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	Relyingpartymetadataurl    string `json:"relyingpartymetadataurl,omitempty"`
	Refreshinterval            int    `json:"refreshinterval,omitempty"`
	Encrypttoken               string `json:"encrypttoken,omitempty"`
	Signatureservice           string `json:"signatureservice,omitempty"`
	Signaturealg               string `json:"signaturealg,omitempty"`
	Attributes                 string `json:"attributes,omitempty"`
	Sendpassword               string `json:"sendpassword,omitempty"`
	Oauthstatus                string `json:"oauthstatus,omitempty"`
	Nextgenapiresource         string `json:"_nextgenapiresource,omitempty"`
}

type Authenticationradiuspolicyvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationsamlpolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Authenticationsmartaccesspolicy struct {
	Name               string `json:"name,omitempty"`
	Rule               string `json:"rule,omitempty"`
	Action             string `json:"action,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Authenticationsmartaccesspolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Authenticationvserverauditsyslogpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Authenticationvservervpnportalthemebinding struct {
	Portaltheme string `json:"portaltheme,omitempty"`
	Acttype     int    `json:"acttype,omitempty"`
	Name        string `json:"name,omitempty"`
}

type Authenticationazurekeyvault struct {
	Name                       string `json:"name,omitempty"`
	Vaultname                  string `json:"vaultname,omitempty"`
	Clientid                   string `json:"clientid,omitempty"`
	Clientsecret               string `json:"clientsecret,omitempty"`
	Servicekeyname             string `json:"servicekeyname,omitempty"`
	Signaturealg               string `json:"signaturealg,omitempty"`
	Tokenendpoint              string `json:"tokenendpoint,omitempty"`
	Pushservice                string `json:"pushservice,omitempty"`
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	Refreshinterval            int    `json:"refreshinterval,omitempty"`
	Tenantid                   string `json:"tenantid,omitempty"`
	Authentication             string `json:"authentication,omitempty"`
	Nextgenapiresource         string `json:"_nextgenapiresource,omitempty"`
}

type Authenticationldappolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Authenticationloginschemapolicy struct {
	Name               string `json:"name,omitempty"`
	Rule               string `json:"rule,omitempty"`
	Action             string `json:"action,omitempty"`
	Undefaction        string `json:"undefaction,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Logaction          string `json:"logaction,omitempty"`
	Newname            string `json:"newname,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Undefhits          string `json:"undefhits,omitempty"`
	Builtin            string `json:"builtin,omitempty"`
	Feature            string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Authenticationloginschemapolicyvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Activepolicy           uint32 `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Authenticationoauthidppolicyauthenticationvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationpolicysystemglobalbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Authenticationpolicylabelauthenticationpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Authenticationvserverldappolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Authenticationnegotiatepolicyvpnvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationsmartaccessprofile struct {
	Name               string `json:"name,omitempty"`
	Tags               string `json:"tags,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Authenticationvserverauthenticationwebauthpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Name                   string `json:"name,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Authenticationvservernegotiatepolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Name                   string `json:"name,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Authenticationsamlpolicyglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationvserverauthenticationsmartaccesspolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Authenticationvservertmsessionpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Authenticationvserverwebauthpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Name                   string `json:"name,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Authenticationadfsproxyprofile struct {
	Name               string `json:"name,omitempty"`
	Username           string `json:"username,omitempty"`
	Password           string `json:"password,omitempty"`
	Serverurl          string `json:"serverurl,omitempty"`
	Certkeyname        string `json:"certkeyname,omitempty"`
	Adfstruststatus    string `json:"adfstruststatus,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Authenticationcertaction struct {
	Name                       string `json:"name,omitempty"`
	Twofactor                  string `json:"twofactor,omitempty"`
	Usernamefield              string `json:"usernamefield,omitempty"`
	Groupnamefield             string `json:"groupnamefield,omitempty"`
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	Nextgenapiresource         string `json:"_nextgenapiresource,omitempty"`
}

type Authenticationoauthidppolicy struct {
	Name                   string `json:"name,omitempty"`
	Rule                   string `json:"rule,omitempty"`
	Action                 string `json:"action,omitempty"`
	Undefaction            string `json:"undefaction,omitempty"`
	Comment                string `json:"comment,omitempty"`
	Logaction              string `json:"logaction,omitempty"`
	Newname                string `json:"newname,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Hits                   string `json:"hits,omitempty"`
	Nextgenapiresource     string `json:"_nextgenapiresource,omitempty"`
}

type Authenticationvserverauthenticationcertpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Name                   string `json:"name,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Authenticationprotecteduseraction struct {
	Name               string `json:"name,omitempty"`
	Realmstr           string `json:"realmstr,omitempty"`
	Maxconcurrentusers int    `json:"maxconcurrentusers,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Authenticationdfapolicy struct {
	Name               string `json:"name,omitempty"`
	Rule               string `json:"rule,omitempty"`
	Action             string `json:"action,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Authenticationdfapolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Authenticationnegotiatepolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Authenticationpolicyglobalbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Activepolicy           uint32 `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Authenticationradiuspolicyauthenticationvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationvservertacacspolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Name                   string `json:"name,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Authenticationpolicy struct {
	Name               string `json:"name,omitempty"`
	Rule               string `json:"rule,omitempty"`
	Action             string `json:"action,omitempty"`
	Undefaction        string `json:"undefaction,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Logaction          string `json:"logaction,omitempty"`
	Newname            string `json:"newname,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Description        string `json:"description,omitempty"`
	Policysubtype      string `json:"policysubtype,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Authenticationpolicylabel struct {
	Labelname              string `json:"labelname,omitempty"`
	Type                   string `json:"type,omitempty"`
	Comment                string `json:"comment,omitempty"`
	Loginschema            string `json:"loginschema,omitempty"`
	Newname                string `json:"newname,omitempty"`
	Numpol                 string `json:"numpol,omitempty"`
	Hits                   string `json:"hits,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               string `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Flowtype               string `json:"flowtype,omitempty"`
	Description            string `json:"description,omitempty"`
	Nextgenapiresource     string `json:"_nextgenapiresource,omitempty"`
}

type Authenticationsamlidpprofile struct {
	Name                        string `json:"name,omitempty"`
	Samlspcertname              string `json:"samlspcertname,omitempty"`
	Samlidpcertname             string `json:"samlidpcertname,omitempty"`
	Assertionconsumerserviceurl string `json:"assertionconsumerserviceurl,omitempty"`
	Sendpassword                string `json:"sendpassword,omitempty"`
	Samlissuername              string `json:"samlissuername,omitempty"`
	Rejectunsignedrequests      string `json:"rejectunsignedrequests,omitempty"`
	Signaturealg                string `json:"signaturealg,omitempty"`
	Digestmethod                string `json:"digestmethod,omitempty"`
	Audience                    string `json:"audience,omitempty"`
	Nameidformat                string `json:"nameidformat,omitempty"`
	Nameidexpr                  string `json:"nameidexpr,omitempty"`
	Attribute1                  string `json:"attribute1,omitempty"`
	Attribute1expr              string `json:"attribute1expr,omitempty"`
	Attribute1friendlyname      string `json:"attribute1friendlyname,omitempty"`
	Attribute1format            string `json:"attribute1format,omitempty"`
	Attribute2                  string `json:"attribute2,omitempty"`
	Attribute2expr              string `json:"attribute2expr,omitempty"`
	Attribute2friendlyname      string `json:"attribute2friendlyname,omitempty"`
	Attribute2format            string `json:"attribute2format,omitempty"`
	Attribute3                  string `json:"attribute3,omitempty"`
	Attribute3expr              string `json:"attribute3expr,omitempty"`
	Attribute3friendlyname      string `json:"attribute3friendlyname,omitempty"`
	Attribute3format            string `json:"attribute3format,omitempty"`
	Attribute4                  string `json:"attribute4,omitempty"`
	Attribute4expr              string `json:"attribute4expr,omitempty"`
	Attribute4friendlyname      string `json:"attribute4friendlyname,omitempty"`
	Attribute4format            string `json:"attribute4format,omitempty"`
	Attribute5                  string `json:"attribute5,omitempty"`
	Attribute5expr              string `json:"attribute5expr,omitempty"`
	Attribute5friendlyname      string `json:"attribute5friendlyname,omitempty"`
	Attribute5format            string `json:"attribute5format,omitempty"`
	Attribute6                  string `json:"attribute6,omitempty"`
	Attribute6expr              string `json:"attribute6expr,omitempty"`
	Attribute6friendlyname      string `json:"attribute6friendlyname,omitempty"`
	Attribute6format            string `json:"attribute6format,omitempty"`
	Attribute7                  string `json:"attribute7,omitempty"`
	Attribute7expr              string `json:"attribute7expr,omitempty"`
	Attribute7friendlyname      string `json:"attribute7friendlyname,omitempty"`
	Attribute7format            string `json:"attribute7format,omitempty"`
	Attribute8                  string `json:"attribute8,omitempty"`
	Attribute8expr              string `json:"attribute8expr,omitempty"`
	Attribute8friendlyname      string `json:"attribute8friendlyname,omitempty"`
	Attribute8format            string `json:"attribute8format,omitempty"`
	Attribute9                  string `json:"attribute9,omitempty"`
	Attribute9expr              string `json:"attribute9expr,omitempty"`
	Attribute9friendlyname      string `json:"attribute9friendlyname,omitempty"`
	Attribute9format            string `json:"attribute9format,omitempty"`
	Attribute10                 string `json:"attribute10,omitempty"`
	Attribute10expr             string `json:"attribute10expr,omitempty"`
	Attribute10friendlyname     string `json:"attribute10friendlyname,omitempty"`
	Attribute10format           string `json:"attribute10format,omitempty"`
	Attribute11                 string `json:"attribute11,omitempty"`
	Attribute11expr             string `json:"attribute11expr,omitempty"`
	Attribute11friendlyname     string `json:"attribute11friendlyname,omitempty"`
	Attribute11format           string `json:"attribute11format,omitempty"`
	Attribute12                 string `json:"attribute12,omitempty"`
	Attribute12expr             string `json:"attribute12expr,omitempty"`
	Attribute12friendlyname     string `json:"attribute12friendlyname,omitempty"`
	Attribute12format           string `json:"attribute12format,omitempty"`
	Attribute13                 string `json:"attribute13,omitempty"`
	Attribute13expr             string `json:"attribute13expr,omitempty"`
	Attribute13friendlyname     string `json:"attribute13friendlyname,omitempty"`
	Attribute13format           string `json:"attribute13format,omitempty"`
	Attribute14                 string `json:"attribute14,omitempty"`
	Attribute14expr             string `json:"attribute14expr,omitempty"`
	Attribute14friendlyname     string `json:"attribute14friendlyname,omitempty"`
	Attribute14format           string `json:"attribute14format,omitempty"`
	Attribute15                 string `json:"attribute15,omitempty"`
	Attribute15expr             string `json:"attribute15expr,omitempty"`
	Attribute15friendlyname     string `json:"attribute15friendlyname,omitempty"`
	Attribute15format           string `json:"attribute15format,omitempty"`
	Attribute16                 string `json:"attribute16,omitempty"`
	Attribute16expr             string `json:"attribute16expr,omitempty"`
	Attribute16friendlyname     string `json:"attribute16friendlyname,omitempty"`
	Attribute16format           string `json:"attribute16format,omitempty"`
	Encryptassertion            string `json:"encryptassertion,omitempty"`
	Encryptionalgorithm         string `json:"encryptionalgorithm,omitempty"`
	Samlbinding                 string `json:"samlbinding,omitempty"`
	Skewtime                    int    `json:"skewtime,omitempty"`
	Serviceproviderid           string `json:"serviceproviderid,omitempty"`
	Signassertion               string `json:"signassertion,omitempty"`
	Keytransportalg             string `json:"keytransportalg,omitempty"`
	Splogouturl                 string `json:"splogouturl,omitempty"`
	Logoutbinding               string `json:"logoutbinding,omitempty"`
	Defaultauthenticationgroup  string `json:"defaultauthenticationgroup,omitempty"`
	Metadataurl                 string `json:"metadataurl,omitempty"`
	Metadatarefreshinterval     int    `json:"metadatarefreshinterval,omitempty"`
	Signatureservice            string `json:"signatureservice,omitempty"`
	Samlsigningcertversion      string `json:"samlsigningcertversion,omitempty"`
	Samlspcertversion           string `json:"samlspcertversion,omitempty"`
	Acsurlrule                  string `json:"acsurlrule,omitempty"`
	Metadataimportstatus        string `json:"metadataimportstatus,omitempty"`
	Nextgenapiresource          string `json:"_nextgenapiresource,omitempty"`
}

type Authenticationvserverbinding struct {
	Name string `json:"name,omitempty"`
}

type Authenticationoauthaction struct {
	Name                       string   `json:"name,omitempty"`
	Oauthtype                  string   `json:"oauthtype,omitempty"`
	Authorizationendpoint      string   `json:"authorizationendpoint,omitempty"`
	Tokenendpoint              string   `json:"tokenendpoint,omitempty"`
	Idtokendecryptendpoint     string   `json:"idtokendecryptendpoint,omitempty"`
	Clientid                   string   `json:"clientid,omitempty"`
	Clientsecret               string   `json:"clientsecret,omitempty"`
	Defaultauthenticationgroup string   `json:"defaultauthenticationgroup,omitempty"`
	Oauthmiscflags             []string `json:"oauthmiscflags,omitempty"`
	Attribute1                 string   `json:"attribute1,omitempty"`
	Attribute2                 string   `json:"attribute2,omitempty"`
	Attribute3                 string   `json:"attribute3,omitempty"`
	Attribute4                 string   `json:"attribute4,omitempty"`
	Attribute5                 string   `json:"attribute5,omitempty"`
	Attribute6                 string   `json:"attribute6,omitempty"`
	Attribute7                 string   `json:"attribute7,omitempty"`
	Attribute8                 string   `json:"attribute8,omitempty"`
	Attribute9                 string   `json:"attribute9,omitempty"`
	Attribute10                string   `json:"attribute10,omitempty"`
	Attribute11                string   `json:"attribute11,omitempty"`
	Attribute12                string   `json:"attribute12,omitempty"`
	Attribute13                string   `json:"attribute13,omitempty"`
	Attribute14                string   `json:"attribute14,omitempty"`
	Attribute15                string   `json:"attribute15,omitempty"`
	Attribute16                string   `json:"attribute16,omitempty"`
	Attributes                 string   `json:"attributes,omitempty"`
	Tenantid                   string   `json:"tenantid,omitempty"`
	Graphendpoint              string   `json:"graphendpoint,omitempty"`
	Refreshinterval            int      `json:"refreshinterval,omitempty"`
	Certendpoint               string   `json:"certendpoint,omitempty"`
	Audience                   string   `json:"audience,omitempty"`
	Usernamefield              string   `json:"usernamefield,omitempty"`
	Skewtime                   int      `json:"skewtime,omitempty"`
	Issuer                     string   `json:"issuer,omitempty"`
	Userinfourl                string   `json:"userinfourl,omitempty"`
	Certfilepath               string   `json:"certfilepath,omitempty"`
	Granttype                  string   `json:"granttype,omitempty"`
	Authentication             string   `json:"authentication,omitempty"`
	Introspecturl              string   `json:"introspecturl,omitempty"`
	Allowedalgorithms          []string `json:"allowedalgorithms,omitempty"`
	Pkce                       string   `json:"pkce,omitempty"`
	Tokenendpointauthmethod    string   `json:"tokenendpointauthmethod,omitempty"`
	Metadataurl                string   `json:"metadataurl,omitempty"`
	Resourceuri                string   `json:"resourceuri,omitempty"`
	Requestattribute           string   `json:"requestattribute,omitempty"`
	Intunedeviceidexpression   string   `json:"intunedeviceidexpression,omitempty"`
	Oauthstatus                string   `json:"oauthstatus,omitempty"`
	Nextgenapiresource         string   `json:"_nextgenapiresource,omitempty"`
}

type Authenticationtacacspolicyvpnglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationvservercspolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Authenticationvserverportalthemebinding struct {
	Portaltheme string `json:"portaltheme,omitempty"`
	Acttype     uint32 `json:"acttype,omitempty"`
	Name        string `json:"name,omitempty"`
}

type Authenticationwebauthpolicysystemglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationnegotiatepolicy struct {
	Name               string `json:"name,omitempty"`
	Rule               string `json:"rule,omitempty"`
	Reqaction          string `json:"reqaction,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Authenticationoauthidppolicyvpnvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationpolicyauthenticationpolicylabelbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Authenticationradiusaction struct {
	Name                       string `json:"name,omitempty"`
	Serverip                   string `json:"serverip,omitempty"`
	Servername                 string `json:"servername,omitempty"`
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
	Transport                  string `json:"transport,omitempty"`
	Targetlbvserver            string `json:"targetlbvserver,omitempty"`
	Messageauthenticator       string `json:"messageauthenticator,omitempty"`
	Ipaddress                  string `json:"ipaddress,omitempty"`
	Success                    string `json:"success,omitempty"`
	Failure                    string `json:"failure,omitempty"`
	Nextgenapiresource         string `json:"_nextgenapiresource,omitempty"`
}

type Authenticationsamlpolicyvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationtacacspolicyglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationtacacspolicysystemglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationdfapolicyvpnvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationlocalpolicyglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationnegotiatepolicyvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationpolicyvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Activepolicy           uint32 `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Authenticationsamlidppolicyvpnvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationwebauthpolicyvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationcertpolicy struct {
	Name               string `json:"name,omitempty"`
	Rule               string `json:"rule,omitempty"`
	Reqaction          string `json:"reqaction,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Authenticationlocalpolicysystemglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationoauthidppolicyvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationradiuspolicyglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationcitrixauthaction struct {
	Name               string `json:"name,omitempty"`
	Authenticationtype string `json:"authenticationtype,omitempty"`
	Authentication     string `json:"authentication,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Authenticationnegotiatepolicyglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationnegotiatepolicyvpnglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationpushservice struct {
	Name                  string `json:"name,omitempty"`
	Clientid              string `json:"clientid,omitempty"`
	Clientsecret          string `json:"clientsecret,omitempty"`
	Customerid            string `json:"customerid,omitempty"`
	Refreshinterval       int    `json:"refreshinterval,omitempty"`
	Namespace             string `json:"Namespace,omitempty"`
	Hubname               string `json:"hubname,omitempty"`
	Servicekey            string `json:"servicekey,omitempty"`
	Servicekeyname        string `json:"servicekeyname,omitempty"`
	Certendpoint          string `json:"certendpoint,omitempty"`
	Pushservicestatus     string `json:"pushservicestatus,omitempty"`
	Trustservice          string `json:"trustservice,omitempty"`
	Pushcloudserverstatus string `json:"pushcloudserverstatus,omitempty"`
	Signingkeyname        string `json:"signingkeyname,omitempty"`
	Signingkey            string `json:"signingkey,omitempty"`
	Nextgenapiresource    string `json:"_nextgenapiresource,omitempty"`
}

type Authenticationradiuspolicyvpnglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationsamlpolicyvpnvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationvserversamlpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Name                   string `json:"name,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Authenticationwebauthpolicyauthenticationvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationnegotiateaction struct {
	Name                       string `json:"name,omitempty"`
	Domain                     string `json:"domain,omitempty"`
	Domainuser                 string `json:"domainuser,omitempty"`
	Domainuserpasswd           string `json:"domainuserpasswd,omitempty"`
	Ou                         string `json:"ou,omitempty"`
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	Keytab                     string `json:"keytab,omitempty"`
	Ntlmpath                   string `json:"ntlmpath,omitempty"`
	Kcdspn                     string `json:"kcdspn,omitempty"`
	Nextgenapiresource         string `json:"_nextgenapiresource,omitempty"`
}

type Authenticationvserverauthenticationnegotiatepolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Name                   string `json:"name,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Authenticationepaaction struct {
	Name               string `json:"name,omitempty"`
	Csecexpr           string `json:"csecexpr,omitempty"`
	Killprocess        string `json:"killprocess,omitempty"`
	Deletefiles        string `json:"deletefiles,omitempty"`
	Defaultepagroup    string `json:"defaultepagroup,omitempty"`
	Quarantinegroup    string `json:"quarantinegroup,omitempty"`
	Deviceposture      string `json:"deviceposture,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Authenticationradiuspolicysystemglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationstorefrontauthaction struct {
	Name                       string `json:"name,omitempty"`
	Serverurl                  string `json:"serverurl,omitempty"`
	Domain                     string `json:"domain,omitempty"`
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	Success                    string `json:"success,omitempty"`
	Failure                    string `json:"failure,omitempty"`
	Nextgenapiresource         string `json:"_nextgenapiresource,omitempty"`
}

type Authenticationvserverauthenticationpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Authenticationvserverauthenticationsamlpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Name                   string `json:"name,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Authenticationvservercertpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Name                   string `json:"name,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Authenticationcertpolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Authenticationcertpolicyvpnglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationldappolicyvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationsamlpolicyauthenticationvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationvserversyslogpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Authenticationemailaction struct {
	Name                       string `json:"name,omitempty"`
	Username                   string `json:"username,omitempty"`
	Password                   string `json:"password,omitempty"`
	Serverurl                  string `json:"serverurl,omitempty"`
	Content                    string `json:"content,omitempty"`
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	Timeout                    int    `json:"timeout,omitempty"`
	Type                       string `json:"type,omitempty"`
	Emailaddress               string `json:"emailaddress,omitempty"`
	Nextgenapiresource         string `json:"_nextgenapiresource,omitempty"`
}

type Authenticationlocalpolicyvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationvserverpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
}

type Authenticationvserverresponderpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
}

type Authenticationsamlaction struct {
	Name                           string   `json:"name,omitempty"`
	Metadataurl                    string   `json:"metadataurl,omitempty"`
	Samlidpcertname                string   `json:"samlidpcertname,omitempty"`
	Samlsigningcertname            string   `json:"samlsigningcertname,omitempty"`
	Samlredirecturl                string   `json:"samlredirecturl,omitempty"`
	Samlacsindex                   int      `json:"samlacsindex,omitempty"`
	Samluserfield                  string   `json:"samluserfield,omitempty"`
	Samlrejectunsignedassertion    string   `json:"samlrejectunsignedassertion,omitempty"`
	Samlissuername                 string   `json:"samlissuername,omitempty"`
	Samltwofactor                  string   `json:"samltwofactor,omitempty"`
	Preferredbindtype              []string `json:"preferredbindtype,omitempty"`
	Defaultauthenticationgroup     string   `json:"defaultauthenticationgroup,omitempty"`
	Attribute1                     string   `json:"attribute1,omitempty"`
	Attribute2                     string   `json:"attribute2,omitempty"`
	Attribute3                     string   `json:"attribute3,omitempty"`
	Attribute4                     string   `json:"attribute4,omitempty"`
	Attribute5                     string   `json:"attribute5,omitempty"`
	Attribute6                     string   `json:"attribute6,omitempty"`
	Attribute7                     string   `json:"attribute7,omitempty"`
	Attribute8                     string   `json:"attribute8,omitempty"`
	Attribute9                     string   `json:"attribute9,omitempty"`
	Attribute10                    string   `json:"attribute10,omitempty"`
	Attribute11                    string   `json:"attribute11,omitempty"`
	Attribute12                    string   `json:"attribute12,omitempty"`
	Attribute13                    string   `json:"attribute13,omitempty"`
	Attribute14                    string   `json:"attribute14,omitempty"`
	Attribute15                    string   `json:"attribute15,omitempty"`
	Attribute16                    string   `json:"attribute16,omitempty"`
	Attributes                     string   `json:"attributes,omitempty"`
	Relaystaterule                 string   `json:"relaystaterule,omitempty"`
	Signaturealg                   string   `json:"signaturealg,omitempty"`
	Digestmethod                   string   `json:"digestmethod,omitempty"`
	Requestedauthncontext          string   `json:"requestedauthncontext,omitempty"`
	Authnctxclassref               []string `json:"authnctxclassref,omitempty"`
	Customauthnctxclassref         string   `json:"customauthnctxclassref,omitempty"`
	Samlbinding                    string   `json:"samlbinding,omitempty"`
	Attributeconsumingserviceindex int      `json:"attributeconsumingserviceindex,omitempty"`
	Sendthumbprint                 string   `json:"sendthumbprint,omitempty"`
	Enforceusername                string   `json:"enforceusername,omitempty"`
	Logouturl                      string   `json:"logouturl,omitempty"`
	Artifactresolutionserviceurl   string   `json:"artifactresolutionserviceurl,omitempty"`
	Skewtime                       int      `json:"skewtime,omitempty"`
	Logoutbinding                  string   `json:"logoutbinding,omitempty"`
	Forceauthn                     string   `json:"forceauthn,omitempty"`
	Groupnamefield                 string   `json:"groupnamefield,omitempty"`
	Audience                       string   `json:"audience,omitempty"`
	Metadatarefreshinterval        int      `json:"metadatarefreshinterval,omitempty"`
	Storesamlresponse              string   `json:"storesamlresponse,omitempty"`
	Statechecks                    string   `json:"statechecks,omitempty"`
	Metadataimportstatus           string   `json:"metadataimportstatus,omitempty"`
	Nextgenapiresource             string   `json:"_nextgenapiresource,omitempty"`
}

type Authenticationvserverauditnslogpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Authenticationvserverauthenticationtacacspolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Name                   string `json:"name,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Authenticationcertpolicyglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationdfapolicyvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationlocalpolicyauthenticationvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationloginschema struct {
	Name                    string `json:"name,omitempty"`
	Authenticationschema    string `json:"authenticationschema,omitempty"`
	Userexpression          string `json:"userexpression,omitempty"`
	Passwdexpression        string `json:"passwdexpression,omitempty"`
	Usercredentialindex     int    `json:"usercredentialindex,omitempty"`
	Passwordcredentialindex int    `json:"passwordcredentialindex,omitempty"`
	Authenticationstrength  int    `json:"authenticationstrength,omitempty"`
	Ssocredentials          string `json:"ssocredentials,omitempty"`
	Builtin                 string `json:"builtin,omitempty"`
	Feature                 string `json:"feature,omitempty"`
	Nextgenapiresource      string `json:"_nextgenapiresource,omitempty"`
}

type Authenticationsamlidppolicyauthenticationvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationvservernslogpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Authenticationwebauthpolicy struct {
	Name               string `json:"name,omitempty"`
	Rule               string `json:"rule,omitempty"`
	Action             string `json:"action,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Authenticationwebauthpolicyglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationlocalpolicyvpnglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationnegotiatepolicyauthenticationvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationradiuspolicy struct {
	Name               string `json:"name,omitempty"`
	Rule               string `json:"rule,omitempty"`
	Reqaction          string `json:"reqaction,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Authenticationradiuspolicyvpnvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationsamlidppolicy struct {
	Name                   string `json:"name,omitempty"`
	Rule                   string `json:"rule,omitempty"`
	Action                 string `json:"action,omitempty"`
	Undefaction            string `json:"undefaction,omitempty"`
	Comment                string `json:"comment,omitempty"`
	Logaction              string `json:"logaction,omitempty"`
	Newname                string `json:"newname,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Hits                   string `json:"hits,omitempty"`
	Nextgenapiresource     string `json:"_nextgenapiresource,omitempty"`
}

type Authenticationauthnprofile struct {
	Name                 string `json:"name,omitempty"`
	Authnvsname          string `json:"authnvsname,omitempty"`
	Authenticationhost   string `json:"authenticationhost,omitempty"`
	Authenticationdomain string `json:"authenticationdomain,omitempty"`
	Authenticationlevel  int    `json:"authenticationlevel,omitempty"`
	Nextgenapiresource   string `json:"_nextgenapiresource,omitempty"`
}

type Authenticationnoauthaction struct {
	Name                       string `json:"name,omitempty"`
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	Builtin                    string `json:"builtin,omitempty"`
	Feature                    string `json:"feature,omitempty"`
	Nextgenapiresource         string `json:"_nextgenapiresource,omitempty"`
}

type Authenticationpolicylabelbinding struct {
	Labelname string `json:"labelname,omitempty"`
}

type Authenticationsamlpolicyvpnglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationvserverauthenticationoauthidppolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Authenticationvserverauthenticationradiuspolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Name                   string `json:"name,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Authenticationvserverauthenticationsamlidppolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Authenticationvserverlocalpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Name                   string `json:"name,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Authenticationdfaaction struct {
	Name                       string `json:"name,omitempty"`
	Clientid                   string `json:"clientid,omitempty"`
	Serverurl                  string `json:"serverurl,omitempty"`
	Passphrase                 string `json:"passphrase,omitempty"`
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	Success                    string `json:"success,omitempty"`
	Failure                    string `json:"failure,omitempty"`
	Nextgenapiresource         string `json:"_nextgenapiresource,omitempty"`
}

type Authenticationldappolicyauthenticationvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationsamlidppolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Authenticationsamlpolicy struct {
	Name               string `json:"name,omitempty"`
	Rule               string `json:"rule,omitempty"`
	Reqaction          string `json:"reqaction,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Authenticationvserver struct {
	Name                 string `json:"name,omitempty"`
	Servicetype          string `json:"servicetype,omitempty"`
	Ipv46                string `json:"ipv46,omitempty"`
	Range                int    `json:"range,omitempty"`
	Port                 int    `json:"port,omitempty"`
	State                string `json:"state,omitempty"`
	Authentication       string `json:"authentication,omitempty"`
	Authenticationdomain string `json:"authenticationdomain,omitempty"`
	Comment              string `json:"comment,omitempty"`
	Td                   int    `json:"td,omitempty"`
	Appflowlog           string `json:"appflowlog,omitempty"`
	Maxloginattempts     int    `json:"maxloginattempts,omitempty"`
	Failedlogintimeout   int    `json:"failedlogintimeout,omitempty"`
	Certkeynames         string `json:"certkeynames,omitempty"`
	Samesite             string `json:"samesite,omitempty"`
	Newname              string `json:"newname,omitempty"`
	Ip                   string `json:"ip,omitempty"`
	Value                string `json:"value,omitempty"`
	Type                 string `json:"type,omitempty"`
	Curstate             string `json:"curstate,omitempty"`
	Status               string `json:"status,omitempty"`
	Cachetype            string `json:"cachetype,omitempty"`
	Redirect             string `json:"redirect,omitempty"`
	Precedence           string `json:"precedence,omitempty"`
	Redirecturl          string `json:"redirecturl,omitempty"`
	Curaaausers          string `json:"curaaausers,omitempty"`
	Policy               string `json:"policy,omitempty"`
	Servicename          string `json:"servicename,omitempty"`
	Weight               string `json:"weight,omitempty"`
	Cachevserver         string `json:"cachevserver,omitempty"`
	Backupvserver        string `json:"backupvserver,omitempty"`
	Clttimeout           string `json:"clttimeout,omitempty"`
	Somethod             string `json:"somethod,omitempty"`
	Sothreshold          string `json:"sothreshold,omitempty"`
	Sopersistence        string `json:"sopersistence,omitempty"`
	Sopersistencetimeout string `json:"sopersistencetimeout,omitempty"`
	Priority             string `json:"priority,omitempty"`
	Downstateflush       string `json:"downstateflush,omitempty"`
	Bindpoint            string `json:"bindpoint,omitempty"`
	Disableprimaryondown string `json:"disableprimaryondown,omitempty"`
	Listenpolicy         string `json:"listenpolicy,omitempty"`
	Listenpriority       string `json:"listenpriority,omitempty"`
	Tcpprofilename       string `json:"tcpprofilename,omitempty"`
	Httpprofilename      string `json:"httpprofilename,omitempty"`
	Vstype               string `json:"vstype,omitempty"`
	Ngname               string `json:"ngname,omitempty"`
	Secondary            string `json:"secondary,omitempty"`
	Groupextraction      string `json:"groupextraction,omitempty"`
	Nextgenapiresource   string `json:"_nextgenapiresource,omitempty"`
}

type Authenticationvserverradiuspolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Name                   string `json:"name,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Authenticationldappolicyvpnglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationpolicylabelpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Authenticationsmartaccesspolicyauthenticationvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationvserveroauthidppolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Authenticationldappolicyglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationlocalpolicy struct {
	Name               string `json:"name,omitempty"`
	Rule               string `json:"rule,omitempty"`
	Reqaction          string `json:"reqaction,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Authenticationpolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Authenticationtacacspolicy struct {
	Name               string `json:"name,omitempty"`
	Rule               string `json:"rule,omitempty"`
	Reqaction          string `json:"reqaction,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Authenticationvserverauthenticationlocalpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Name                   string `json:"name,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Authenticationwebauthaction struct {
	Name                       string `json:"name,omitempty"`
	Serverip                   string `json:"serverip,omitempty"`
	Serverport                 int    `json:"serverport,omitempty"`
	Fullreqexpr                string `json:"fullreqexpr,omitempty"`
	Scheme                     string `json:"scheme,omitempty"`
	Successrule                string `json:"successrule,omitempty"`
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	Attribute1                 string `json:"attribute1,omitempty"`
	Attribute2                 string `json:"attribute2,omitempty"`
	Attribute3                 string `json:"attribute3,omitempty"`
	Attribute4                 string `json:"attribute4,omitempty"`
	Attribute5                 string `json:"attribute5,omitempty"`
	Attribute6                 string `json:"attribute6,omitempty"`
	Attribute7                 string `json:"attribute7,omitempty"`
	Attribute8                 string `json:"attribute8,omitempty"`
	Attribute9                 string `json:"attribute9,omitempty"`
	Attribute10                string `json:"attribute10,omitempty"`
	Attribute11                string `json:"attribute11,omitempty"`
	Attribute12                string `json:"attribute12,omitempty"`
	Attribute13                string `json:"attribute13,omitempty"`
	Attribute14                string `json:"attribute14,omitempty"`
	Attribute15                string `json:"attribute15,omitempty"`
	Attribute16                string `json:"attribute16,omitempty"`
	Nextgenapiresource         string `json:"_nextgenapiresource,omitempty"`
}

type Authenticationwebauthpolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Authenticationloginschemapolicyvpnvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Authenticationtacacspolicyvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Authenticationvservercachepolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Nextfactor             string `json:"nextfactor,omitempty"`
}
