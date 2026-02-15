package models

// vpn configuration structs
type VpnclientlessaccesspolicyVpnvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type VpnglobalAuditnslogpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type Vpnicadtlsconnection struct {
	Channelnumber      int     `json:"channelnumber,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Destip             string  `json:"destip,omitempty"`
	Destport           int     `json:"destport,omitempty"`
	Domain             string  `json:"domain,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Nodeid             int     `json:"nodeid,omitempty"`
	Peid               int     `json:"peid,omitempty"`
	Srcip              string  `json:"srcip,omitempty"`
	Srcport            int     `json:"srcport,omitempty"`
	Username           string  `json:"username,omitempty"`
}

type Vpnstoreinfo struct {
	Count              float64 `json:"__count,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Storeapisupport    string  `json:"storeapisupport,omitempty"`
	Storelist          string  `json:"storelist,omitempty"`
	Storeserverissf    string  `json:"storeserverissf,omitempty"`
	Storeserverstatus  string  `json:"storeserverstatus,omitempty"`
	Storestatus        string  `json:"storestatus,omitempty"`
	Url                string  `json:"url,omitempty"`
}

type VpnglobalIntranetipBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Intranetip             string `json:"intranetip,omitempty"`
	Netmask                string `json:"netmask,omitempty"`
}

type VpnurlpolicyVpnvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type VpnurlpolicyAaauserBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type Vpneula struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
}

type VpnsessionpolicyVpnvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type VpnsessionpolicyBinding struct {
	Name                              string        `json:"name,omitempty"`
	VpnsessionpolicyAaagroupBinding   []interface{} `json:"vpnsessionpolicy_aaagroup_binding,omitempty"`
	VpnsessionpolicyAaauserBinding    []interface{} `json:"vpnsessionpolicy_aaauser_binding,omitempty"`
	VpnsessionpolicyVpnglobalBinding  []interface{} `json:"vpnsessionpolicy_vpnglobal_binding,omitempty"`
	VpnsessionpolicyVpnvserverBinding []interface{} `json:"vpnsessionpolicy_vpnvserver_binding,omitempty"`
}

type VpnvserverCspolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VpnvserverAuditsyslogpolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VpnglobalVpnclientlessaccesspolicyBinding struct {
	Builtin                []string `json:"builtin,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	Globalbindtype         string   `json:"globalbindtype,omitempty"`
	Gotopriorityexpression string   `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool     `json:"groupextraction,omitempty"`
	Policyname             string   `json:"policyname,omitempty"`
	Priority               int      `json:"priority,omitempty"`
	Secondary              bool     `json:"secondary,omitempty"`
	TypeField              string   `json:"type,omitempty"`
}

type VpnglobalSecureprivateaccessurlBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Secureprivateaccessurl string `json:"secureprivateaccessurl,omitempty"`
}

type VpnvserverAppcontrollerBinding struct {
	Acttype       int    `json:"acttype,omitempty"`
	Appcontroller string `json:"appcontroller,omitempty"`
	Name          string `json:"name,omitempty"`
}

type VpnglobalVpnsecureprivateaccessprofileBinding struct {
	Gotopriorityexpression     string `json:"gotopriorityexpression,omitempty"`
	Secureprivateaccessprofile string `json:"secureprivateaccessprofile,omitempty"`
}

type VpnurlpolicyAaagroupBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type VpnglobalAuthenticationcertpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VpnvserverAnalyticsprofileBinding struct {
	Analyticsprofile string `json:"analyticsprofile,omitempty"`
	Name             string `json:"name,omitempty"`
}

type Vpnintranetapplication struct {
	Clientapplication   []string `json:"clientapplication,omitempty"`
	Count               float64  `json:"__count,omitempty"`
	Destip              string   `json:"destip,omitempty"`
	Destport            string   `json:"destport,omitempty"`
	Hostname            string   `json:"hostname,omitempty"`
	Interception        string   `json:"interception,omitempty"`
	Intranetapplication string   `json:"intranetapplication,omitempty"`
	Ipaddress           string   `json:"ipaddress,omitempty"`
	Iprange             string   `json:"iprange,omitempty"`
	Netmask             string   `json:"netmask,omitempty"`
	Nextgenapiresource  string   `json:"_nextgenapiresource,omitempty"`
	Protocol            string   `json:"protocol,omitempty"`
	Spoofiip            string   `json:"spoofiip,omitempty"`
	Srcip               string   `json:"srcip,omitempty"`
	Srcport             int      `json:"srcport,omitempty"`
}

type VpnglobalVpntrafficpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type Vpnurlaction struct {
	Actualurl          string  `json:"actualurl,omitempty"`
	Applicationtype    string  `json:"applicationtype,omitempty"`
	Clientlessaccess   string  `json:"clientlessaccess,omitempty"`
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Iconurl            string  `json:"iconurl,omitempty"`
	Linkname           string  `json:"linkname,omitempty"`
	Name               string  `json:"name,omitempty"`
	Newname            string  `json:"newname,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Samlssoprofile     string  `json:"samlssoprofile,omitempty"`
	Ssotype            string  `json:"ssotype,omitempty"`
	Vservername        string  `json:"vservername,omitempty"`
}

type Vpnpcoipconnection struct {
	All                bool    `json:"all,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Destip             string  `json:"destip,omitempty"`
	Destport           int     `json:"destport,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Nodeid             int     `json:"nodeid,omitempty"`
	Peid               int     `json:"peid,omitempty"`
	Srcip              string  `json:"srcip,omitempty"`
	Srcport            int     `json:"srcport,omitempty"`
	Username           string  `json:"username,omitempty"`
}

type VpnvserverVpneulaBinding struct {
	Acttype int    `json:"acttype,omitempty"`
	Eula    string `json:"eula,omitempty"`
	Name    string `json:"name,omitempty"`
}

type VpnvserverAuthenticationnegotiatepolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type Vpnsfconfig struct {
	Count              float64  `json:"__count,omitempty"`
	Filename           string   `json:"filename,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Vserver            []string `json:"vserver,omitempty"`
}

type Vpnportaltheme struct {
	Basetheme          string  `json:"basetheme,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Feature            string  `json:"feature,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
}

type VpntrafficpolicyVpnvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type VpnglobalAuthenticationldappolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VpnglobalVpneulaBinding struct {
	Eula                   string `json:"eula,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpnparameter struct {
	Accessrestrictedpageredirect string        `json:"accessrestrictedpageredirect,omitempty"`
	Advancedclientlessvpnmode    string        `json:"advancedclientlessvpnmode,omitempty"`
	Allowedlogingroups           string        `json:"allowedlogingroups,omitempty"`
	Allprotocolproxy             string        `json:"allprotocolproxy,omitempty"`
	Alwaysonprofilename          string        `json:"alwaysonprofilename,omitempty"`
	Apptokentimeout              int           `json:"apptokentimeout,omitempty"`
	Authorizationgroup           string        `json:"authorizationgroup,omitempty"`
	Autoproxyurl                 string        `json:"autoproxyurl,omitempty"`
	Backendcertvalidation        string        `json:"backendcertvalidation,omitempty"`
	Backenddtls12                string        `json:"backenddtls12,omitempty"`
	Backendserversni             string        `json:"backendserversni,omitempty"`
	Citrixreceiverhome           string        `json:"citrixreceiverhome,omitempty"`
	Clientchoices                string        `json:"clientchoices,omitempty"`
	Clientcleanupprompt          string        `json:"clientcleanupprompt,omitempty"`
	Clientconfiguration          []string      `json:"clientconfiguration,omitempty"`
	Clientdebug                  string        `json:"clientdebug,omitempty"`
	Clientidletimeout            int           `json:"clientidletimeout,omitempty"`
	Clientidletimeoutwarning     int           `json:"clientidletimeoutwarning,omitempty"`
	Clientlessmodeurlencoding    string        `json:"clientlessmodeurlencoding,omitempty"`
	Clientlesspersistentcookie   string        `json:"clientlesspersistentcookie,omitempty"`
	Clientlessvpnmode            string        `json:"clientlessvpnmode,omitempty"`
	Clientoptions                []string      `json:"clientoptions,omitempty"`
	Clientsecurity               string        `json:"clientsecurity,omitempty"`
	Clientsecuritygroup          string        `json:"clientsecuritygroup,omitempty"`
	Clientsecuritylog            string        `json:"clientsecuritylog,omitempty"`
	Clientsecuritymessage        string        `json:"clientsecuritymessage,omitempty"`
	Clientversions               string        `json:"clientversions,omitempty"`
	Defaultauthorizationaction   string        `json:"defaultauthorizationaction,omitempty"`
	Deviceposture                string        `json:"deviceposture,omitempty"`
	Dnsvservername               string        `json:"dnsvservername,omitempty"`
	Emailhome                    string        `json:"emailhome,omitempty"`
	Encryptcsecexp               string        `json:"encryptcsecexp,omitempty"`
	Epaclienttype                string        `json:"epaclienttype,omitempty"`
	Forcecleanup                 []string      `json:"forcecleanup,omitempty"`
	Forcedtimeout                int           `json:"forcedtimeout,omitempty"`
	Forcedtimeoutwarning         int           `json:"forcedtimeoutwarning,omitempty"`
	Fqdnspoofedip                string        `json:"fqdnspoofedip,omitempty"`
	Ftpproxy                     string        `json:"ftpproxy,omitempty"`
	Gopherproxy                  string        `json:"gopherproxy,omitempty"`
	Homepage                     string        `json:"homepage,omitempty"`
	Httpport                     []interface{} `json:"httpport,omitempty"`
	Httpproxy                    string        `json:"httpproxy,omitempty"`
	Httptrackconnproxy           string        `json:"httptrackconnproxy,omitempty"`
	Icaproxy                     string        `json:"icaproxy,omitempty"`
	Icasessiontimeout            string        `json:"icasessiontimeout,omitempty"`
	Icauseraccounting            string        `json:"icauseraccounting,omitempty"`
	Iconwithreceiver             string        `json:"iconwithreceiver,omitempty"`
	Iipdnssuffix                 string        `json:"iipdnssuffix,omitempty"`
	Kcdaccount                   string        `json:"kcdaccount,omitempty"`
	Killconnections              string        `json:"killconnections,omitempty"`
	Linuxpluginupgrade           string        `json:"linuxpluginupgrade,omitempty"`
	Locallanaccess               string        `json:"locallanaccess,omitempty"`
	Loginscript                  string        `json:"loginscript,omitempty"`
	Logoutscript                 string        `json:"logoutscript,omitempty"`
	Macpluginupgrade             string        `json:"macpluginupgrade,omitempty"`
	Maxiipperuser                int           `json:"maxiipperuser,omitempty"`
	Mdxtokentimeout              int           `json:"mdxtokentimeout,omitempty"`
	Name                         string        `json:"name,omitempty"`
	Netmask                      string        `json:"netmask,omitempty"`
	Nextgenapiresource           string        `json:"_nextgenapiresource,omitempty"`
	Ntdomain                     string        `json:"ntdomain,omitempty"`
	Pcoipprofilename             string        `json:"pcoipprofilename,omitempty"`
	Proxy                        string        `json:"proxy,omitempty"`
	Proxyexception               string        `json:"proxyexception,omitempty"`
	Proxylocalbypass             string        `json:"proxylocalbypass,omitempty"`
	Rdpclientprofilename         string        `json:"rdpclientprofilename,omitempty"`
	Rfc1918                      string        `json:"rfc1918,omitempty"`
	Samesite                     string        `json:"samesite,omitempty"`
	Securebrowse                 string        `json:"securebrowse,omitempty"`
	Secureprivateaccess          string        `json:"secureprivateaccess,omitempty"`
	Secureprivateaccessprofile   string        `json:"secureprivateaccessprofile,omitempty"`
	Sesstimeout                  int           `json:"sesstimeout,omitempty"`
	Smartgroup                   string        `json:"smartgroup,omitempty"`
	Socksproxy                   string        `json:"socksproxy,omitempty"`
	Splitdns                     string        `json:"splitdns,omitempty"`
	Splittunnel                  string        `json:"splittunnel,omitempty"`
	Spoofiip                     string        `json:"spoofiip,omitempty"`
	Sslproxy                     string        `json:"sslproxy,omitempty"`
	Sso                          string        `json:"sso,omitempty"`
	Ssocredential                string        `json:"ssocredential,omitempty"`
	Storefronturl                string        `json:"storefronturl,omitempty"`
	Transparentinterception      string        `json:"transparentinterception,omitempty"`
	Uitheme                      string        `json:"uitheme,omitempty"`
	Useiip                       string        `json:"useiip,omitempty"`
	Usemip                       string        `json:"usemip,omitempty"`
	Userdomains                  string        `json:"userdomains,omitempty"`
	Vpnsessionpolicybindtype     string        `json:"vpnsessionpolicybindtype,omitempty"`
	Vpnsessionpolicycount        int           `json:"vpnsessionpolicycount,omitempty"`
	Wihome                       string        `json:"wihome,omitempty"`
	Wihomeaddresstype            string        `json:"wihomeaddresstype,omitempty"`
	Windowsautologon             string        `json:"windowsautologon,omitempty"`
	Windowsclienttype            string        `json:"windowsclienttype,omitempty"`
	Windowspluginupgrade         string        `json:"windowspluginupgrade,omitempty"`
	Winsip                       string        `json:"winsip,omitempty"`
	Wiportalmode                 string        `json:"wiportalmode,omitempty"`
}

type Vpnclientlessaccessprofile struct {
	Builtin                        []string `json:"builtin,omitempty"`
	Clientconsumedcookies          string   `json:"clientconsumedcookies,omitempty"`
	Count                          float64  `json:"__count,omitempty"`
	Cssrewritepolicylabel          string   `json:"cssrewritepolicylabel,omitempty"`
	Description                    string   `json:"description,omitempty"`
	Feature                        string   `json:"feature,omitempty"`
	Isdefault                      bool     `json:"isdefault,omitempty"`
	Javascriptrewritepolicylabel   string   `json:"javascriptrewritepolicylabel,omitempty"`
	Nextgenapiresource             string   `json:"_nextgenapiresource,omitempty"`
	Profilename                    string   `json:"profilename,omitempty"`
	Regexforfindingcustomurls      string   `json:"regexforfindingcustomurls,omitempty"`
	Regexforfindingurlincss        string   `json:"regexforfindingurlincss,omitempty"`
	Regexforfindingurlinjavascript string   `json:"regexforfindingurlinjavascript,omitempty"`
	Regexforfindingurlinxcomponent string   `json:"regexforfindingurlinxcomponent,omitempty"`
	Regexforfindingurlinxml        string   `json:"regexforfindingurlinxml,omitempty"`
	Reqhdrrewritepolicylabel       string   `json:"reqhdrrewritepolicylabel,omitempty"`
	Requirepersistentcookie        string   `json:"requirepersistentcookie,omitempty"`
	Reshdrrewritepolicylabel       string   `json:"reshdrrewritepolicylabel,omitempty"`
	Urlrewritepolicylabel          string   `json:"urlrewritepolicylabel,omitempty"`
	Xcomponentrewritepolicylabel   string   `json:"xcomponentrewritepolicylabel,omitempty"`
	Xmlrewritepolicylabel          string   `json:"xmlrewritepolicylabel,omitempty"`
}

type Vpnsamlssoprofile struct {
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
	Digestmethod                string  `json:"digestmethod,omitempty"`
	Encryptassertion            string  `json:"encryptassertion,omitempty"`
	Encryptionalgorithm         string  `json:"encryptionalgorithm,omitempty"`
	Name                        string  `json:"name,omitempty"`
	Nameidexpr                  string  `json:"nameidexpr,omitempty"`
	Nameidformat                string  `json:"nameidformat,omitempty"`
	Nextgenapiresource          string  `json:"_nextgenapiresource,omitempty"`
	Relaystaterule              string  `json:"relaystaterule,omitempty"`
	Samlissuername              string  `json:"samlissuername,omitempty"`
	Samlsigningcertname         string  `json:"samlsigningcertname,omitempty"`
	Samlspcertname              string  `json:"samlspcertname,omitempty"`
	Sendpassword                string  `json:"sendpassword,omitempty"`
	Signassertion               string  `json:"signassertion,omitempty"`
	Signaturealg                string  `json:"signaturealg,omitempty"`
	Signatureservice            string  `json:"signatureservice,omitempty"`
	Skewtime                    int     `json:"skewtime,omitempty"`
}

type VpnvserverAuditnslogpolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VpnglobalVpnsessionpolicyBinding struct {
	Builtin                []string `json:"builtin,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	Gotopriorityexpression string   `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool     `json:"groupextraction,omitempty"`
	Policyname             string   `json:"policyname,omitempty"`
	Priority               int      `json:"priority,omitempty"`
	Secondary              bool     `json:"secondary,omitempty"`
}

type Vpnpcoipvserverprofile struct {
	Count              float64 `json:"__count,omitempty"`
	Logindomain        string  `json:"logindomain,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Udpport            int     `json:"udpport,omitempty"`
}

type VpnsessionpolicyVpnglobalBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type VpnvserverIntranetip6Binding struct {
	Acttype     int    `json:"acttype,omitempty"`
	Intranetip6 string `json:"intranetip6,omitempty"`
	Name        string `json:"name,omitempty"`
	Numaddr     int    `json:"numaddr,omitempty"`
}

type VpnvserverFeopolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VpnglobalAuditsyslogpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VpnvserverVpnsessionpolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VpnvserverVpnurlBinding struct {
	Acttype int    `json:"acttype,omitempty"`
	Name    string `json:"name,omitempty"`
	Urlname string `json:"urlname,omitempty"`
}

type VpnvserverVpnurlpolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VpnvserverAuthenticationwebauthpolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type Vpnclientlessaccesspolicy struct {
	Builtin            []string `json:"builtin,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Description        string   `json:"description,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	Isdefault          bool     `json:"isdefault,omitempty"`
	Name               string   `json:"name,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Profilename        string   `json:"profilename,omitempty"`
	Rule               string   `json:"rule,omitempty"`
	Undefaction        string   `json:"undefaction,omitempty"`
	Undefhits          int      `json:"undefhits,omitempty"`
}

type VpnvserverAaapreauthenticationpolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VpntrafficpolicyBinding struct {
	Name                              string        `json:"name,omitempty"`
	VpntrafficpolicyAaagroupBinding   []interface{} `json:"vpntrafficpolicy_aaagroup_binding,omitempty"`
	VpntrafficpolicyAaauserBinding    []interface{} `json:"vpntrafficpolicy_aaauser_binding,omitempty"`
	VpntrafficpolicyVpnglobalBinding  []interface{} `json:"vpntrafficpolicy_vpnglobal_binding,omitempty"`
	VpntrafficpolicyVpnvserverBinding []interface{} `json:"vpntrafficpolicy_vpnvserver_binding,omitempty"`
}

type VpnvserverAuthenticationsamlidppolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type Vpnnexthopserver struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Nexthopfqdn        string  `json:"nexthopfqdn,omitempty"`
	Nexthopip          string  `json:"nexthopip,omitempty"`
	Nexthopport        int     `json:"nexthopport,omitempty"`
	Resaddresstype     string  `json:"resaddresstype,omitempty"`
	Secure             string  `json:"secure,omitempty"`
}

type Vpnurlpolicy struct {
	Action             string   `json:"action,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	Comment            string   `json:"comment,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Logaction          string   `json:"logaction,omitempty"`
	Name               string   `json:"name,omitempty"`
	Newname            string   `json:"newname,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Rule               string   `json:"rule,omitempty"`
	Undefhits          int      `json:"undefhits,omitempty"`
}

type VpnglobalVpnurlpolicyBinding struct {
	Builtin                []string `json:"builtin,omitempty"`
	Gotopriorityexpression string   `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool     `json:"groupextraction,omitempty"`
	Policyname             string   `json:"policyname,omitempty"`
	Priority               int      `json:"priority,omitempty"`
	Secondary              bool     `json:"secondary,omitempty"`
}

type Vpnsessionpolicy struct {
	Action             string   `json:"action,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Expressiontype     string   `json:"expressiontype,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	Name               string   `json:"name,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Rule               string   `json:"rule,omitempty"`
}

type Vpnurl struct {
	Actualurl          string  `json:"actualurl,omitempty"`
	Appjson            string  `json:"appjson,omitempty"`
	Applicationtype    string  `json:"applicationtype,omitempty"`
	Clientlessaccess   string  `json:"clientlessaccess,omitempty"`
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Iconurl            string  `json:"iconurl,omitempty"`
	Linkname           string  `json:"linkname,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Samlssoprofile     string  `json:"samlssoprofile,omitempty"`
	Ssotype            string  `json:"ssotype,omitempty"`
	Urlname            string  `json:"urlname,omitempty"`
	Vservername        string  `json:"vservername,omitempty"`
}

type VpnglobalAuthenticationpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VpnvserverAuthenticationtacacspolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VpnglobalVpnnexthopserverBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Nexthopserver          string `json:"nexthopserver,omitempty"`
}

type VpnvserverIcapolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VpnvserverAuthenticationcertpolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VpnvserverAuthenticationsamlpolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VpnurlpolicyVpnglobalBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type VpntrafficpolicyAaagroupBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type VpnvserverCachepolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VpnglobalVpnportalthemeBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Portaltheme            string `json:"portaltheme,omitempty"`
}

type VpnvserverSharefileserverBinding struct {
	Acttype   int    `json:"acttype,omitempty"`
	Name      string `json:"name,omitempty"`
	Sharefile string `json:"sharefile,omitempty"`
}

type VpnglobalAuthenticationsamlpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type Vpntrafficpolicy struct {
	Action             string  `json:"action,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Expressiontype     string  `json:"expressiontype,omitempty"`
	Hits               int     `json:"hits,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Rule               string  `json:"rule,omitempty"`
}

type VpnvserverAuthenticationloginschemapolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VpnvserverVpntrafficpolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VpnsessionpolicyAaauserBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type VpnvserverVpnintranetapplicationBinding struct {
	Acttype             int    `json:"acttype,omitempty"`
	Intranetapplication string `json:"intranetapplication,omitempty"`
	Name                string `json:"name,omitempty"`
}

type VpnvserverAppflowpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VpnglobalAuthenticationradiuspolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VpnclientlessaccesspolicyVpnglobalBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type VpnglobalAuthenticationlocalpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VpnglobalVpnurlBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Urlname                string `json:"urlname,omitempty"`
}

type VpnsessionpolicyAaagroupBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type VpnvserverBinding struct {
	Name                                             string        `json:"name,omitempty"`
	VpnvserverAaapreauthenticationpolicyBinding      []interface{} `json:"vpnvserver_aaapreauthenticationpolicy_binding,omitempty"`
	VpnvserverAnalyticsprofileBinding                []interface{} `json:"vpnvserver_analyticsprofile_binding,omitempty"`
	VpnvserverAppcontrollerBinding                   []interface{} `json:"vpnvserver_appcontroller_binding,omitempty"`
	VpnvserverAppflowpolicyBinding                   []interface{} `json:"vpnvserver_appflowpolicy_binding,omitempty"`
	VpnvserverAppfwpolicyBinding                     []interface{} `json:"vpnvserver_appfwpolicy_binding,omitempty"`
	VpnvserverAuditnslogpolicyBinding                []interface{} `json:"vpnvserver_auditnslogpolicy_binding,omitempty"`
	VpnvserverAuditsyslogpolicyBinding               []interface{} `json:"vpnvserver_auditsyslogpolicy_binding,omitempty"`
	VpnvserverAuthenticationcertpolicyBinding        []interface{} `json:"vpnvserver_authenticationcertpolicy_binding,omitempty"`
	VpnvserverAuthenticationdfapolicyBinding         []interface{} `json:"vpnvserver_authenticationdfapolicy_binding,omitempty"`
	VpnvserverAuthenticationldappolicyBinding        []interface{} `json:"vpnvserver_authenticationldappolicy_binding,omitempty"`
	VpnvserverAuthenticationlocalpolicyBinding       []interface{} `json:"vpnvserver_authenticationlocalpolicy_binding,omitempty"`
	VpnvserverAuthenticationloginschemapolicyBinding []interface{} `json:"vpnvserver_authenticationloginschemapolicy_binding,omitempty"`
	VpnvserverAuthenticationnegotiatepolicyBinding   []interface{} `json:"vpnvserver_authenticationnegotiatepolicy_binding,omitempty"`
	VpnvserverAuthenticationoauthidppolicyBinding    []interface{} `json:"vpnvserver_authenticationoauthidppolicy_binding,omitempty"`
	VpnvserverAuthenticationpolicyBinding            []interface{} `json:"vpnvserver_authenticationpolicy_binding,omitempty"`
	VpnvserverAuthenticationradiuspolicyBinding      []interface{} `json:"vpnvserver_authenticationradiuspolicy_binding,omitempty"`
	VpnvserverAuthenticationsamlidppolicyBinding     []interface{} `json:"vpnvserver_authenticationsamlidppolicy_binding,omitempty"`
	VpnvserverAuthenticationsamlpolicyBinding        []interface{} `json:"vpnvserver_authenticationsamlpolicy_binding,omitempty"`
	VpnvserverAuthenticationtacacspolicyBinding      []interface{} `json:"vpnvserver_authenticationtacacspolicy_binding,omitempty"`
	VpnvserverAuthenticationwebauthpolicyBinding     []interface{} `json:"vpnvserver_authenticationwebauthpolicy_binding,omitempty"`
	VpnvserverCachepolicyBinding                     []interface{} `json:"vpnvserver_cachepolicy_binding,omitempty"`
	VpnvserverCspolicyBinding                        []interface{} `json:"vpnvserver_cspolicy_binding,omitempty"`
	VpnvserverFeopolicyBinding                       []interface{} `json:"vpnvserver_feopolicy_binding,omitempty"`
	VpnvserverIcapolicyBinding                       []interface{} `json:"vpnvserver_icapolicy_binding,omitempty"`
	VpnvserverIntranetip6Binding                     []interface{} `json:"vpnvserver_intranetip6_binding,omitempty"`
	VpnvserverIntranetipBinding                      []interface{} `json:"vpnvserver_intranetip_binding,omitempty"`
	VpnvserverResponderpolicyBinding                 []interface{} `json:"vpnvserver_responderpolicy_binding,omitempty"`
	VpnvserverRewritepolicyBinding                   []interface{} `json:"vpnvserver_rewritepolicy_binding,omitempty"`
	VpnvserverSecureprivateaccessurlBinding          []interface{} `json:"vpnvserver_secureprivateaccessurl_binding,omitempty"`
	VpnvserverSharefileserverBinding                 []interface{} `json:"vpnvserver_sharefileserver_binding,omitempty"`
	VpnvserverStaserverBinding                       []interface{} `json:"vpnvserver_staserver_binding,omitempty"`
	VpnvserverVpnclientlessaccesspolicyBinding       []interface{} `json:"vpnvserver_vpnclientlessaccesspolicy_binding,omitempty"`
	VpnvserverVpnepaprofileBinding                   []interface{} `json:"vpnvserver_vpnepaprofile_binding,omitempty"`
	VpnvserverVpneulaBinding                         []interface{} `json:"vpnvserver_vpneula_binding,omitempty"`
	VpnvserverVpnintranetapplicationBinding          []interface{} `json:"vpnvserver_vpnintranetapplication_binding,omitempty"`
	VpnvserverVpnnexthopserverBinding                []interface{} `json:"vpnvserver_vpnnexthopserver_binding,omitempty"`
	VpnvserverVpnportalthemeBinding                  []interface{} `json:"vpnvserver_vpnportaltheme_binding,omitempty"`
	VpnvserverVpnsecureprivateaccessprofileBinding   []interface{} `json:"vpnvserver_vpnsecureprivateaccessprofile_binding,omitempty"`
	VpnvserverVpnsessionpolicyBinding                []interface{} `json:"vpnvserver_vpnsessionpolicy_binding,omitempty"`
	VpnvserverVpntrafficpolicyBinding                []interface{} `json:"vpnvserver_vpntrafficpolicy_binding,omitempty"`
	VpnvserverVpnurlBinding                          []interface{} `json:"vpnvserver_vpnurl_binding,omitempty"`
	VpnvserverVpnurlpolicyBinding                    []interface{} `json:"vpnvserver_vpnurlpolicy_binding,omitempty"`
}

type Vpnsessionaction struct {
	Advancedclientlessvpnmode  string        `json:"advancedclientlessvpnmode,omitempty"`
	Allowedlogingroups         string        `json:"allowedlogingroups,omitempty"`
	Allprotocolproxy           string        `json:"allprotocolproxy,omitempty"`
	Alwaysonprofilename        string        `json:"alwaysonprofilename,omitempty"`
	Authorizationgroup         string        `json:"authorizationgroup,omitempty"`
	Autoproxyurl               string        `json:"autoproxyurl,omitempty"`
	Builtin                    []string      `json:"builtin,omitempty"`
	Citrixreceiverhome         string        `json:"citrixreceiverhome,omitempty"`
	Clientchoices              string        `json:"clientchoices,omitempty"`
	Clientcleanupprompt        string        `json:"clientcleanupprompt,omitempty"`
	Clientconfiguration        []string      `json:"clientconfiguration,omitempty"`
	Clientdebug                string        `json:"clientdebug,omitempty"`
	Clientidletimeout          int           `json:"clientidletimeout,omitempty"`
	Clientidletimeoutwarning   int           `json:"clientidletimeoutwarning,omitempty"`
	Clientlessmodeurlencoding  string        `json:"clientlessmodeurlencoding,omitempty"`
	Clientlesspersistentcookie string        `json:"clientlesspersistentcookie,omitempty"`
	Clientlessvpnmode          string        `json:"clientlessvpnmode,omitempty"`
	Clientoptions              string        `json:"clientoptions,omitempty"`
	Clientsecurity             string        `json:"clientsecurity,omitempty"`
	Clientsecuritygroup        string        `json:"clientsecuritygroup,omitempty"`
	Clientsecuritylog          string        `json:"clientsecuritylog,omitempty"`
	Clientsecuritymessage      string        `json:"clientsecuritymessage,omitempty"`
	Count                      float64       `json:"__count,omitempty"`
	Defaultauthorizationaction string        `json:"defaultauthorizationaction,omitempty"`
	Dnsvservername             string        `json:"dnsvservername,omitempty"`
	Emailhome                  string        `json:"emailhome,omitempty"`
	Epaclienttype              string        `json:"epaclienttype,omitempty"`
	Feature                    string        `json:"feature,omitempty"`
	Forcecleanup               []string      `json:"forcecleanup,omitempty"`
	Forcedtimeout              int           `json:"forcedtimeout,omitempty"`
	Forcedtimeoutwarning       int           `json:"forcedtimeoutwarning,omitempty"`
	Fqdnspoofedip              string        `json:"fqdnspoofedip,omitempty"`
	Ftpproxy                   string        `json:"ftpproxy,omitempty"`
	Gopherproxy                string        `json:"gopherproxy,omitempty"`
	Homepage                   string        `json:"homepage,omitempty"`
	Httpport                   []interface{} `json:"httpport,omitempty"`
	Httpproxy                  string        `json:"httpproxy,omitempty"`
	Icaproxy                   string        `json:"icaproxy,omitempty"`
	Iconwithreceiver           string        `json:"iconwithreceiver,omitempty"`
	Iipdnssuffix               string        `json:"iipdnssuffix,omitempty"`
	Kcdaccount                 string        `json:"kcdaccount,omitempty"`
	Killconnections            string        `json:"killconnections,omitempty"`
	Linuxpluginupgrade         string        `json:"linuxpluginupgrade,omitempty"`
	Locallanaccess             string        `json:"locallanaccess,omitempty"`
	Loginscript                string        `json:"loginscript,omitempty"`
	Logoutscript               string        `json:"logoutscript,omitempty"`
	Macpluginupgrade           string        `json:"macpluginupgrade,omitempty"`
	Name                       string        `json:"name,omitempty"`
	Netmask                    string        `json:"netmask,omitempty"`
	Nextgenapiresource         string        `json:"_nextgenapiresource,omitempty"`
	Ntdomain                   string        `json:"ntdomain,omitempty"`
	Pcoipprofilename           string        `json:"pcoipprofilename,omitempty"`
	Proxy                      string        `json:"proxy,omitempty"`
	Proxyexception             string        `json:"proxyexception,omitempty"`
	Proxylocalbypass           string        `json:"proxylocalbypass,omitempty"`
	Rdpclientprofilename       string        `json:"rdpclientprofilename,omitempty"`
	Rfc1918                    string        `json:"rfc1918,omitempty"`
	Securebrowse               string        `json:"securebrowse,omitempty"`
	Sesstimeout                int           `json:"sesstimeout,omitempty"`
	Sfgatewayauthtype          string        `json:"sfgatewayauthtype,omitempty"`
	Smartgroup                 string        `json:"smartgroup,omitempty"`
	Socksproxy                 string        `json:"socksproxy,omitempty"`
	Splitdns                   string        `json:"splitdns,omitempty"`
	Splittunnel                string        `json:"splittunnel,omitempty"`
	Spoofiip                   string        `json:"spoofiip,omitempty"`
	Sslproxy                   string        `json:"sslproxy,omitempty"`
	Sso                        string        `json:"sso,omitempty"`
	Ssocredential              string        `json:"ssocredential,omitempty"`
	Storefronturl              string        `json:"storefronturl,omitempty"`
	Transparentinterception    string        `json:"transparentinterception,omitempty"`
	Useiip                     string        `json:"useiip,omitempty"`
	Usemip                     string        `json:"usemip,omitempty"`
	Useraccounting             string        `json:"useraccounting,omitempty"`
	Wihome                     string        `json:"wihome,omitempty"`
	Wihomeaddresstype          string        `json:"wihomeaddresstype,omitempty"`
	Windowsautologon           string        `json:"windowsautologon,omitempty"`
	Windowsclienttype          string        `json:"windowsclienttype,omitempty"`
	Windowspluginupgrade       string        `json:"windowspluginupgrade,omitempty"`
	Winsip                     string        `json:"winsip,omitempty"`
	Wiportalmode               string        `json:"wiportalmode,omitempty"`
}

type VpnvserverAuthenticationldappolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type Vpntrafficaction struct {
	Apptimeout         int     `json:"apptimeout,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Formssoaction      string  `json:"formssoaction,omitempty"`
	Fta                string  `json:"fta,omitempty"`
	Hdx                string  `json:"hdx,omitempty"`
	Kcdaccount         string  `json:"kcdaccount,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Passwdexpression   string  `json:"passwdexpression,omitempty"`
	Proxy              string  `json:"proxy,omitempty"`
	Qual               string  `json:"qual,omitempty"`
	Samlssoprofile     string  `json:"samlssoprofile,omitempty"`
	Sso                string  `json:"sso,omitempty"`
	Userexpression     string  `json:"userexpression,omitempty"`
	Wanscaler          string  `json:"wanscaler,omitempty"`
}

type VpnvserverVpnsecureprivateaccessprofileBinding struct {
	Acttype                    int    `json:"acttype,omitempty"`
	Name                       string `json:"name,omitempty"`
	Secureprivateaccessprofile string `json:"secureprivateaccessprofile,omitempty"`
}

type VpnvserverVpnepaprofileBinding struct {
	Acttype            int    `json:"acttype,omitempty"`
	Epaprofile         string `json:"epaprofile,omitempty"`
	Epaprofileoptional bool   `json:"epaprofileoptional,omitempty"`
	Name               string `json:"name,omitempty"`
}

type Vpnformssoaction struct {
	Actionurl          string  `json:"actionurl,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Namevaluepair      string  `json:"namevaluepair,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Nvtype             string  `json:"nvtype,omitempty"`
	Passwdfield        string  `json:"passwdfield,omitempty"`
	Responsesize       int     `json:"responsesize,omitempty"`
	Ssosuccessrule     string  `json:"ssosuccessrule,omitempty"`
	Submitmethod       string  `json:"submitmethod,omitempty"`
	Userfield          string  `json:"userfield,omitempty"`
}

type VpnvserverAuthenticationradiuspolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VpnvserverAuthenticationlocalpolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VpntrafficpolicyVpnglobalBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type VpnurlpolicyBinding struct {
	Name                          string        `json:"name,omitempty"`
	VpnurlpolicyAaagroupBinding   []interface{} `json:"vpnurlpolicy_aaagroup_binding,omitempty"`
	VpnurlpolicyAaauserBinding    []interface{} `json:"vpnurlpolicy_aaauser_binding,omitempty"`
	VpnurlpolicyVpnglobalBinding  []interface{} `json:"vpnurlpolicy_vpnglobal_binding,omitempty"`
	VpnurlpolicyVpnvserverBinding []interface{} `json:"vpnurlpolicy_vpnvserver_binding,omitempty"`
}

type VpnvserverResponderpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VpnvserverSecureprivateaccessurlBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secureprivateaccessurl string `json:"secureprivateaccessurl,omitempty"`
}

type VpnvserverAuthenticationpolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VpnglobalBinding struct {
	VpnglobalAppcontrollerBinding                 []interface{} `json:"vpnglobal_appcontroller_binding,omitempty"`
	VpnglobalAppfwpolicyBinding                   []interface{} `json:"vpnglobal_appfwpolicy_binding,omitempty"`
	VpnglobalAuditnslogpolicyBinding              []interface{} `json:"vpnglobal_auditnslogpolicy_binding,omitempty"`
	VpnglobalAuditsyslogpolicyBinding             []interface{} `json:"vpnglobal_auditsyslogpolicy_binding,omitempty"`
	VpnglobalAuthenticationcertpolicyBinding      []interface{} `json:"vpnglobal_authenticationcertpolicy_binding,omitempty"`
	VpnglobalAuthenticationldappolicyBinding      []interface{} `json:"vpnglobal_authenticationldappolicy_binding,omitempty"`
	VpnglobalAuthenticationlocalpolicyBinding     []interface{} `json:"vpnglobal_authenticationlocalpolicy_binding,omitempty"`
	VpnglobalAuthenticationnegotiatepolicyBinding []interface{} `json:"vpnglobal_authenticationnegotiatepolicy_binding,omitempty"`
	VpnglobalAuthenticationpolicyBinding          []interface{} `json:"vpnglobal_authenticationpolicy_binding,omitempty"`
	VpnglobalAuthenticationradiuspolicyBinding    []interface{} `json:"vpnglobal_authenticationradiuspolicy_binding,omitempty"`
	VpnglobalAuthenticationsamlpolicyBinding      []interface{} `json:"vpnglobal_authenticationsamlpolicy_binding,omitempty"`
	VpnglobalAuthenticationtacacspolicyBinding    []interface{} `json:"vpnglobal_authenticationtacacspolicy_binding,omitempty"`
	VpnglobalGslbdomainBinding                    []interface{} `json:"vpnglobal_gslbdomain_binding,omitempty"`
	VpnglobalIntranetip6Binding                   []interface{} `json:"vpnglobal_intranetip6_binding,omitempty"`
	VpnglobalIntranetipBinding                    []interface{} `json:"vpnglobal_intranetip_binding,omitempty"`
	VpnglobalSecureprivateaccessurlBinding        []interface{} `json:"vpnglobal_secureprivateaccessurl_binding,omitempty"`
	VpnglobalSharefileserverBinding               []interface{} `json:"vpnglobal_sharefileserver_binding,omitempty"`
	VpnglobalSslcertkeyBinding                    []interface{} `json:"vpnglobal_sslcertkey_binding,omitempty"`
	VpnglobalStaserverBinding                     []interface{} `json:"vpnglobal_staserver_binding,omitempty"`
	VpnglobalVpnclientlessaccesspolicyBinding     []interface{} `json:"vpnglobal_vpnclientlessaccesspolicy_binding,omitempty"`
	VpnglobalVpneulaBinding                       []interface{} `json:"vpnglobal_vpneula_binding,omitempty"`
	VpnglobalVpnintranetapplicationBinding        []interface{} `json:"vpnglobal_vpnintranetapplication_binding,omitempty"`
	VpnglobalVpnnexthopserverBinding              []interface{} `json:"vpnglobal_vpnnexthopserver_binding,omitempty"`
	VpnglobalVpnportalthemeBinding                []interface{} `json:"vpnglobal_vpnportaltheme_binding,omitempty"`
	VpnglobalVpnsecureprivateaccessprofileBinding []interface{} `json:"vpnglobal_vpnsecureprivateaccessprofile_binding,omitempty"`
	VpnglobalVpnsessionpolicyBinding              []interface{} `json:"vpnglobal_vpnsessionpolicy_binding,omitempty"`
	VpnglobalVpntrafficpolicyBinding              []interface{} `json:"vpnglobal_vpntrafficpolicy_binding,omitempty"`
	VpnglobalVpnurlBinding                        []interface{} `json:"vpnglobal_vpnurl_binding,omitempty"`
	VpnglobalVpnurlpolicyBinding                  []interface{} `json:"vpnglobal_vpnurlpolicy_binding,omitempty"`
}

type VpnglobalAppcontrollerBinding struct {
	Appcontroller          string `json:"appcontroller,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type VpnvserverAuthenticationdfapolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type Vpnpcoipprofile struct {
	Conserverurl       string  `json:"conserverurl,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Icvverification    string  `json:"icvverification,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Sessionidletimeout int     `json:"sessionidletimeout,omitempty"`
}

type VpnglobalSharefileserverBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Sharefile              string `json:"sharefile,omitempty"`
}

type VpntrafficpolicyAaauserBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type Vpnalwaysonprofile struct {
	Clientcontrol             string  `json:"clientcontrol,omitempty"`
	Count                     float64 `json:"__count,omitempty"`
	Locationbasedvpn          string  `json:"locationbasedvpn,omitempty"`
	Name                      string  `json:"name,omitempty"`
	Networkaccessonvpnfailure string  `json:"networkaccessonvpnfailure,omitempty"`
	Nextgenapiresource        string  `json:"_nextgenapiresource,omitempty"`
}

type VpnvserverVpnnexthopserverBinding struct {
	Acttype       int    `json:"acttype,omitempty"`
	Name          string `json:"name,omitempty"`
	Nexthopserver string `json:"nexthopserver,omitempty"`
}

type VpnglobalAppfwpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VpnvserverStaserverBinding struct {
	Acttype        int    `json:"acttype,omitempty"`
	Name           string `json:"name,omitempty"`
	Staaddresstype string `json:"staaddresstype,omitempty"`
	Staauthid      string `json:"staauthid,omitempty"`
	Staserver      string `json:"staserver,omitempty"`
	Stastate       string `json:"stastate,omitempty"`
}

type VpnglobalVpnintranetapplicationBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Intranetapplication    string `json:"intranetapplication,omitempty"`
}

type VpnvserverIntranetipBinding struct {
	Acttype    int    `json:"acttype,omitempty"`
	Intranetip string `json:"intranetip,omitempty"`
	MapField   string `json:"map,omitempty"`
	Name       string `json:"name,omitempty"`
	Netmask    string `json:"netmask,omitempty"`
}

type VpnclientlessaccesspolicyBinding struct {
	Name                                       string        `json:"name,omitempty"`
	VpnclientlessaccesspolicyVpnglobalBinding  []interface{} `json:"vpnclientlessaccesspolicy_vpnglobal_binding,omitempty"`
	VpnclientlessaccesspolicyVpnvserverBinding []interface{} `json:"vpnclientlessaccesspolicy_vpnvserver_binding,omitempty"`
}

type VpnglobalAuthenticationnegotiatepolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VpnglobalAuthenticationtacacspolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VpnglobalIntranetip6Binding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Intranetip6            string `json:"intranetip6,omitempty"`
	Numaddr                int    `json:"numaddr,omitempty"`
}

type Vpnvserver struct {
	Accessrestrictedpageredirect string  `json:"accessrestrictedpageredirect,omitempty"`
	Advancedepa                  string  `json:"advancedepa,omitempty"`
	Appflowlog                   string  `json:"appflowlog,omitempty"`
	Authentication               string  `json:"authentication,omitempty"`
	Authnprofile                 string  `json:"authnprofile,omitempty"`
	Backupvserver                string  `json:"backupvserver,omitempty"`
	Bindpoint                    string  `json:"bindpoint,omitempty"`
	Cachetype                    string  `json:"cachetype,omitempty"`
	Cachevserver                 string  `json:"cachevserver,omitempty"`
	Certkeynames                 string  `json:"certkeynames,omitempty"`
	Cginfrahomepageredirect      string  `json:"cginfrahomepageredirect,omitempty"`
	Clttimeout                   int     `json:"clttimeout,omitempty"`
	Comment                      string  `json:"comment,omitempty"`
	Count                        float64 `json:"__count,omitempty"`
	Csvserver                    string  `json:"csvserver,omitempty"`
	Curaaausers                  int     `json:"curaaausers,omitempty"`
	Curstate                     string  `json:"curstate,omitempty"`
	Curtotalusers                int     `json:"curtotalusers,omitempty"`
	Deploymenttype               string  `json:"deploymenttype,omitempty"`
	Devicecert                   string  `json:"devicecert,omitempty"`
	Deviceposture                string  `json:"deviceposture,omitempty"`
	Disableprimaryondown         string  `json:"disableprimaryondown,omitempty"`
	Domain                       string  `json:"domain,omitempty"`
	Doublehop                    string  `json:"doublehop,omitempty"`
	Downstateflush               string  `json:"downstateflush,omitempty"`
	Dtls                         string  `json:"dtls,omitempty"`
	Epaprofileoptional           bool    `json:"epaprofileoptional,omitempty"`
	Failedlogintimeout           int     `json:"failedlogintimeout,omitempty"`
	Groupextraction              bool    `json:"groupextraction,omitempty"`
	Httpprofilename              string  `json:"httpprofilename,omitempty"`
	Icaonly                      string  `json:"icaonly,omitempty"`
	Icaproxysessionmigration     string  `json:"icaproxysessionmigration,omitempty"`
	Icmpvsrresponse              string  `json:"icmpvsrresponse,omitempty"`
	Ip                           string  `json:"ip,omitempty"`
	Ipset                        string  `json:"ipset,omitempty"`
	Ipv46                        string  `json:"ipv46,omitempty"`
	L2conn                       string  `json:"l2conn,omitempty"`
	Linuxepapluginupgrade        string  `json:"linuxepapluginupgrade,omitempty"`
	Listenpolicy                 string  `json:"listenpolicy,omitempty"`
	Listenpriority               int     `json:"listenpriority,omitempty"`
	Loginonce                    string  `json:"loginonce,omitempty"`
	Logoutonsmartcardremoval     string  `json:"logoutonsmartcardremoval,omitempty"`
	Macepapluginupgrade          string  `json:"macepapluginupgrade,omitempty"`
	MapField                     string  `json:"map,omitempty"`
	Maxaaausers                  int     `json:"maxaaausers,omitempty"`
	Maxloginattempts             int     `json:"maxloginattempts,omitempty"`
	Name                         string  `json:"name,omitempty"`
	Netprofile                   string  `json:"netprofile,omitempty"`
	Newname                      string  `json:"newname,omitempty"`
	Nextgenapiresource           string  `json:"_nextgenapiresource,omitempty"`
	Ngname                       string  `json:"ngname,omitempty"`
	Nodefaultbindings            string  `json:"nodefaultbindings,omitempty"`
	Pcoipvserverprofilename      string  `json:"pcoipvserverprofilename,omitempty"`
	Port                         int     `json:"port,omitempty"`
	Precedence                   string  `json:"precedence,omitempty"`
	Quicprofilename              string  `json:"quicprofilename,omitempty"`
	Range                        int     `json:"range,omitempty"`
	Rdpserverprofilename         string  `json:"rdpserverprofilename,omitempty"`
	Redirect                     string  `json:"redirect,omitempty"`
	Redirecturl                  string  `json:"redirecturl,omitempty"`
	Response                     string  `json:"response,omitempty"`
	Rhistate                     string  `json:"rhistate,omitempty"`
	Rule                         string  `json:"rule,omitempty"`
	Samesite                     string  `json:"samesite,omitempty"`
	Secondary                    bool    `json:"secondary,omitempty"`
	Secureprivateaccess          string  `json:"secureprivateaccess,omitempty"`
	Servicename                  string  `json:"servicename,omitempty"`
	Servicetype                  string  `json:"servicetype,omitempty"`
	Somethod                     string  `json:"somethod,omitempty"`
	Sopersistence                string  `json:"sopersistence,omitempty"`
	Sopersistencetimeout         int     `json:"sopersistencetimeout,omitempty"`
	Sothreshold                  int     `json:"sothreshold,omitempty"`
	State                        string  `json:"state,omitempty"`
	Status                       int     `json:"status,omitempty"`
	Tcpprofilename               string  `json:"tcpprofilename,omitempty"`
	TypeField                    string  `json:"type,omitempty"`
	Usemip                       string  `json:"usemip,omitempty"`
	Userdomains                  string  `json:"userdomains,omitempty"`
	Value                        string  `json:"value,omitempty"`
	Vserverfqdn                  string  `json:"vserverfqdn,omitempty"`
	Weight                       int     `json:"weight,omitempty"`
	Windowsepapluginupgrade      string  `json:"windowsepapluginupgrade,omitempty"`
}

type VpnglobalSslcertkeyBinding struct {
	Cacert                 string `json:"cacert,omitempty"`
	Certkeyname            string `json:"certkeyname,omitempty"`
	Crlcheck               string `json:"crlcheck,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Ocspcheck              string `json:"ocspcheck,omitempty"`
	Userdataencryptionkey  string `json:"userdataencryptionkey,omitempty"`
}

type VpnvserverAppfwpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type Vpnicaconnection struct {
	All                bool    `json:"all,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Destip             string  `json:"destip,omitempty"`
	Destport           int     `json:"destport,omitempty"`
	Domain             string  `json:"domain,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Nodeid             int     `json:"nodeid,omitempty"`
	Peid               int     `json:"peid,omitempty"`
	Productname        string  `json:"productname,omitempty"`
	Srcip              string  `json:"srcip,omitempty"`
	Srcport            int     `json:"srcport,omitempty"`
	Tenantname         string  `json:"tenantname,omitempty"`
	Transproto         string  `json:"transproto,omitempty"`
	Username           string  `json:"username,omitempty"`
}

type VpnvserverRewritepolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VpnvserverAuthenticationoauthidppolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}

type VpnglobalStaserverBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Staaddresstype         string `json:"staaddresstype,omitempty"`
	Staauthid              string `json:"staauthid,omitempty"`
	Staserver              string `json:"staserver,omitempty"`
	Stastate               string `json:"stastate,omitempty"`
}

type Vpnsecureprivateaccessprofile struct {
	Accessrestrictedpageredirect string  `json:"accessrestrictedpageredirect,omitempty"`
	Chromeenterprisepremiummode  string  `json:"chromeenterprisepremiummode,omitempty"`
	Clouddeployment              string  `json:"clouddeployment,omitempty"`
	Count                        float64 `json:"__count,omitempty"`
	Customerid                   string  `json:"customerid,omitempty"`
	Name                         string  `json:"name,omitempty"`
	Nextgenapiresource           string  `json:"_nextgenapiresource,omitempty"`
	Url                          string  `json:"url,omitempty"`
}

type Vpnepaprofile struct {
	Count              float64 `json:"__count,omitempty"`
	Data               string  `json:"data,omitempty"`
	Filename           string  `json:"filename,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
}

type VpnvserverVpnportalthemeBinding struct {
	Acttype     int    `json:"acttype,omitempty"`
	Name        string `json:"name,omitempty"`
	Portaltheme string `json:"portaltheme,omitempty"`
}

type VpnglobalDomainBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Intranetdomain         string `json:"intranetdomain,omitempty"`
}

type VpnvserverVpnclientlessaccesspolicyBinding struct {
	Acttype                int    `json:"acttype,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
}
