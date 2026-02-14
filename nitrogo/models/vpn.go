// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Vpnglobalauthenticationtacacspolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpnglobaleulabinding struct {
	Eula                   string `json:"eula,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpnglobalportalthemebinding struct {
	Portaltheme            string `json:"portaltheme,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpnglobalstaserverbinding struct {
	Staserver              string `json:"staserver,omitempty"`
	Staauthid              string `json:"staauthid,omitempty"`
	Stastate               string `json:"stastate,omitempty"`
	Staaddresstype         string `json:"staaddresstype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpnglobalurlpolicybinding struct {
	Policyname             string   `json:"policyname,omitempty"`
	Priority               uint32   `json:"priority,omitempty"`
	Gotopriorityexpression string   `json:"gotopriorityexpression,omitempty"`
	Builtin                []string `json:"builtin,omitempty"`
	Secondary              bool     `json:"secondary,omitempty"`
	Groupextraction        bool     `json:"groupextraction,omitempty"`
}

type Vpnglobaltrafficpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
}

type Vpnportaltheme struct {
	Name               string `json:"name,omitempty"`
	Basetheme          string `json:"basetheme,omitempty"`
	Feature            string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Vpnsessionpolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Vpnvserversamlidppolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnvservervpnepaprofilebinding struct {
	Epaprofile         string `json:"epaprofile,omitempty"`
	Acttype            int    `json:"acttype,omitempty"`
	Name               string `json:"name,omitempty"`
	Epaprofileoptional bool   `json:"epaprofileoptional,omitempty"`
}

type Vpnvservervpnurlbinding struct {
	Urlname string `json:"urlname,omitempty"`
	Acttype int    `json:"acttype,omitempty"`
	Name    string `json:"name,omitempty"`
}

type Vpnglobalappfwpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
}

type Vpnglobalnslogpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpnglobalsamlpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpnglobalsessionpolicybinding struct {
	Policyname             string   `json:"policyname,omitempty"`
	Priority               uint32   `json:"priority,omitempty"`
	Gotopriorityexpression string   `json:"gotopriorityexpression,omitempty"`
	Builtin                []string `json:"builtin,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	Secondary              bool     `json:"secondary,omitempty"`
	Groupextraction        bool     `json:"groupextraction,omitempty"`
}

type Vpnsessionpolicygroupbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Vpnvserverauthenticationcertpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Name                   string `json:"name,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnvserverprofilebinding struct {
	Analyticsprofile string `json:"analyticsprofile,omitempty"`
	Name             string `json:"name,omitempty"`
}

type Vpngloballocalpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpnsessionpolicyvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Vpntrafficpolicygroupbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Vpntrafficpolicyvpnglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Vpnvserverintranetipbinding struct {
	Intranetip string `json:"intranetip,omitempty"`
	Netmask    string `json:"netmask,omitempty"`
	Map        string `json:"map,omitempty"`
	Acttype    int    `json:"acttype,omitempty"`
	Name       string `json:"name,omitempty"`
}

type Vpnvserverldappolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnvservervpnintranetapplicationbinding struct {
	Intranetapplication string `json:"intranetapplication,omitempty"`
	Acttype             int    `json:"acttype,omitempty"`
	Name                string `json:"name,omitempty"`
}

type Vpnglobalintranetip6binding struct {
	Intranetip6            string `json:"intranetip6,omitempty"`
	Numaddr                int    `json:"numaddr,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpnvservercertpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Name                   string `json:"name,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnvservervpnnexthopserverbinding struct {
	Nexthopserver string `json:"nexthopserver,omitempty"`
	Acttype       int    `json:"acttype,omitempty"`
	Name          string `json:"name,omitempty"`
}

type Vpnclientlessaccesspolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Vpnglobalauthenticationpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpnpcoipconnection struct {
	Username           string `json:"username,omitempty"`
	Nodeid             int    `json:"nodeid,omitempty"`
	All                bool   `json:"all,omitempty"`
	Srcip              string `json:"srcip,omitempty"`
	Srcport            string `json:"srcport,omitempty"`
	Destip             string `json:"destip,omitempty"`
	Destport           string `json:"destport,omitempty"`
	Peid               string `json:"peid,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Vpnurlpolicyglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Vpnvserverappfwpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnurlpolicyvpnvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Vpnvserver struct {
	Name                         string `json:"name,omitempty"`
	Servicetype                  string `json:"servicetype,omitempty"`
	Ipv46                        string `json:"ipv46,omitempty"`
	Range                        int    `json:"range,omitempty"`
	Port                         int    `json:"port,omitempty"`
	Ipset                        string `json:"ipset,omitempty"`
	State                        string `json:"state,omitempty"`
	Authentication               string `json:"authentication,omitempty"`
	Doublehop                    string `json:"doublehop,omitempty"`
	Maxaaausers                  int    `json:"maxaaausers,omitempty"`
	Icaonly                      string `json:"icaonly,omitempty"`
	Icaproxysessionmigration     string `json:"icaproxysessionmigration,omitempty"`
	Dtls                         string `json:"dtls,omitempty"`
	Loginonce                    string `json:"loginonce,omitempty"`
	Advancedepa                  string `json:"advancedepa,omitempty"`
	Devicecert                   string `json:"devicecert,omitempty"`
	Certkeynames                 string `json:"certkeynames,omitempty"`
	Downstateflush               string `json:"downstateflush,omitempty"`
	Listenpolicy                 string `json:"listenpolicy,omitempty"`
	Listenpriority               int    `json:"listenpriority,omitempty"`
	Tcpprofilename               string `json:"tcpprofilename,omitempty"`
	Httpprofilename              string `json:"httpprofilename,omitempty"`
	Comment                      string `json:"comment,omitempty"`
	Appflowlog                   string `json:"appflowlog,omitempty"`
	Icmpvsrresponse              string `json:"icmpvsrresponse,omitempty"`
	Rhistate                     string `json:"rhistate,omitempty"`
	Netprofile                   string `json:"netprofile,omitempty"`
	Cginfrahomepageredirect      string `json:"cginfrahomepageredirect,omitempty"`
	Secureprivateaccess          string `json:"secureprivateaccess,omitempty"`
	Accessrestrictedpageredirect string `json:"accessrestrictedpageredirect,omitempty"`
	Maxloginattempts             int    `json:"maxloginattempts,omitempty"`
	Failedlogintimeout           int    `json:"failedlogintimeout,omitempty"`
	L2conn                       string `json:"l2conn,omitempty"`
	Deploymenttype               string `json:"deploymenttype,omitempty"`
	Rdpserverprofilename         string `json:"rdpserverprofilename,omitempty"`
	Windowsepapluginupgrade      string `json:"windowsepapluginupgrade,omitempty"`
	Linuxepapluginupgrade        string `json:"linuxepapluginupgrade,omitempty"`
	Macepapluginupgrade          string `json:"macepapluginupgrade,omitempty"`
	Logoutonsmartcardremoval     string `json:"logoutonsmartcardremoval,omitempty"`
	Userdomains                  string `json:"userdomains,omitempty"`
	Authnprofile                 string `json:"authnprofile,omitempty"`
	Vserverfqdn                  string `json:"vserverfqdn,omitempty"`
	Pcoipvserverprofilename      string `json:"pcoipvserverprofilename,omitempty"`
	Samesite                     string `json:"samesite,omitempty"`
	Quicprofilename              string `json:"quicprofilename,omitempty"`
	Deviceposture                string `json:"deviceposture,omitempty"`
	Newname                      string `json:"newname,omitempty"`
	Ip                           string `json:"ip,omitempty"`
	Value                        string `json:"value,omitempty"`
	Type                         string `json:"type,omitempty"`
	Curstate                     string `json:"curstate,omitempty"`
	Status                       string `json:"status,omitempty"`
	Cachetype                    string `json:"cachetype,omitempty"`
	Redirect                     string `json:"redirect,omitempty"`
	Precedence                   string `json:"precedence,omitempty"`
	Redirecturl                  string `json:"redirecturl,omitempty"`
	Curaaausers                  string `json:"curaaausers,omitempty"`
	Curtotalusers                string `json:"curtotalusers,omitempty"`
	Domain                       string `json:"domain,omitempty"`
	Rule                         string `json:"rule,omitempty"`
	Servicename                  string `json:"servicename,omitempty"`
	Weight                       string `json:"weight,omitempty"`
	Cachevserver                 string `json:"cachevserver,omitempty"`
	Backupvserver                string `json:"backupvserver,omitempty"`
	Clttimeout                   string `json:"clttimeout,omitempty"`
	Somethod                     string `json:"somethod,omitempty"`
	Sothreshold                  string `json:"sothreshold,omitempty"`
	Sopersistence                string `json:"sopersistence,omitempty"`
	Sopersistencetimeout         string `json:"sopersistencetimeout,omitempty"`
	Usemip                       string `json:"usemip,omitempty"`
	Map                          string `json:"map,omitempty"`
	Bindpoint                    string `json:"bindpoint,omitempty"`
	Disableprimaryondown         string `json:"disableprimaryondown,omitempty"`
	Secondary                    string `json:"secondary,omitempty"`
	Groupextraction              string `json:"groupextraction,omitempty"`
	Epaprofileoptional           string `json:"epaprofileoptional,omitempty"`
	Ngname                       string `json:"ngname,omitempty"`
	Csvserver                    string `json:"csvserver,omitempty"`
	Nodefaultbindings            string `json:"nodefaultbindings,omitempty"`
	Nextgenapiresource           string `json:"_nextgenapiresource,omitempty"`
	Response                     string `json:"response,omitempty"`
}

type Vpnvserverauthenticationdfapolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Name                   string `json:"name,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnclientlessaccesspolicyvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy int32  `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Vpnglobalintranetipbinding struct {
	Intranetip             string `json:"intranetip,omitempty"`
	Netmask                string `json:"netmask,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpnvserverauditsyslogpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnvservervpnportalthemebinding struct {
	Portaltheme string `json:"portaltheme,omitempty"`
	Acttype     int    `json:"acttype,omitempty"`
	Name        string `json:"name,omitempty"`
}

type Vpnclientlessaccessprofile struct {
	Profilename                    string `json:"profilename,omitempty"`
	Urlrewritepolicylabel          string `json:"urlrewritepolicylabel,omitempty"`
	Javascriptrewritepolicylabel   string `json:"javascriptrewritepolicylabel,omitempty"`
	Reqhdrrewritepolicylabel       string `json:"reqhdrrewritepolicylabel,omitempty"`
	Reshdrrewritepolicylabel       string `json:"reshdrrewritepolicylabel,omitempty"`
	Regexforfindingurlinjavascript string `json:"regexforfindingurlinjavascript,omitempty"`
	Regexforfindingurlincss        string `json:"regexforfindingurlincss,omitempty"`
	Regexforfindingurlinxcomponent string `json:"regexforfindingurlinxcomponent,omitempty"`
	Regexforfindingurlinxml        string `json:"regexforfindingurlinxml,omitempty"`
	Regexforfindingcustomurls      string `json:"regexforfindingcustomurls,omitempty"`
	Clientconsumedcookies          string `json:"clientconsumedcookies,omitempty"`
	Requirepersistentcookie        string `json:"requirepersistentcookie,omitempty"`
	Cssrewritepolicylabel          string `json:"cssrewritepolicylabel,omitempty"`
	Xmlrewritepolicylabel          string `json:"xmlrewritepolicylabel,omitempty"`
	Xcomponentrewritepolicylabel   string `json:"xcomponentrewritepolicylabel,omitempty"`
	Isdefault                      string `json:"isdefault,omitempty"`
	Description                    string `json:"description,omitempty"`
	Builtin                        string `json:"builtin,omitempty"`
	Feature                        string `json:"feature,omitempty"`
	Nextgenapiresource             string `json:"_nextgenapiresource,omitempty"`
}

type Vpngloballdappolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpnvserverauthenticationwebauthpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Name                   string `json:"name,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnvserverbinding struct {
	Name string `json:"name,omitempty"`
}

type Vpnvserverloginschemapolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnglobalvpnportalthemebinding struct {
	Portaltheme            string `json:"portaltheme,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpnalwaysonprofile struct {
	Name                      string `json:"name,omitempty"`
	Networkaccessonvpnfailure string `json:"networkaccessonvpnfailure,omitempty"`
	Clientcontrol             string `json:"clientcontrol,omitempty"`
	Locationbasedvpn          string `json:"locationbasedvpn,omitempty"`
	Nextgenapiresource        string `json:"_nextgenapiresource,omitempty"`
}

type Vpnglobalclientlessaccesspolicybinding struct {
	Policyname             string   `json:"policyname,omitempty"`
	Priority               uint32   `json:"priority,omitempty"`
	Type                   string   `json:"type,omitempty"`
	Gotopriorityexpression string   `json:"gotopriorityexpression,omitempty"`
	Builtin                []string `json:"builtin,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	Globalbindtype         string   `json:"globalbindtype,omitempty"`
	Secondary              bool     `json:"secondary,omitempty"`
	Groupextraction        bool     `json:"groupextraction,omitempty"`
}

type Vpnglobalvpnsessionpolicybinding struct {
	Policyname             string   `json:"policyname,omitempty"`
	Priority               int      `json:"priority,omitempty"`
	Gotopriorityexpression string   `json:"gotopriorityexpression,omitempty"`
	Builtin                []string `json:"builtin,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	Secondary              bool     `json:"secondary,omitempty"`
	Groupextraction        bool     `json:"groupextraction,omitempty"`
}

type Vpnurlaction struct {
	Name               string `json:"name,omitempty"`
	Linkname           string `json:"linkname,omitempty"`
	Actualurl          string `json:"actualurl,omitempty"`
	Vservername        string `json:"vservername,omitempty"`
	Clientlessaccess   string `json:"clientlessaccess,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Iconurl            string `json:"iconurl,omitempty"`
	Ssotype            string `json:"ssotype,omitempty"`
	Applicationtype    string `json:"applicationtype,omitempty"`
	Samlssoprofile     string `json:"samlssoprofile,omitempty"`
	Newname            string `json:"newname,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Vpnvserverrewritepolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
}

type Vpnglobalauthenticationlocalpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpnglobalbinding struct {
}

type Vpnintranetapplication struct {
	Intranetapplication string   `json:"intranetapplication,omitempty"`
	Protocol            string   `json:"protocol,omitempty"`
	Destip              string   `json:"destip,omitempty"`
	Netmask             string   `json:"netmask,omitempty"`
	Iprange             string   `json:"iprange,omitempty"`
	Hostname            string   `json:"hostname,omitempty"`
	Clientapplication   []string `json:"clientapplication,omitempty"`
	Spoofiip            string   `json:"spoofiip,omitempty"`
	Destport            string   `json:"destport,omitempty"`
	Interception        string   `json:"interception,omitempty"`
	Srcip               string   `json:"srcip,omitempty"`
	Srcport             int      `json:"srcport,omitempty"`
	Ipaddress           string   `json:"ipaddress,omitempty"`
	Nextgenapiresource  string   `json:"_nextgenapiresource,omitempty"`
}

type Vpnparameter struct {
	Httpport                     []int    `json:"httpport,omitempty"`
	Winsip                       string   `json:"winsip,omitempty"`
	Dnsvservername               string   `json:"dnsvservername,omitempty"`
	Splitdns                     string   `json:"splitdns,omitempty"`
	Icauseraccounting            string   `json:"icauseraccounting,omitempty"`
	Sesstimeout                  int      `json:"sesstimeout,omitempty"`
	Clientsecurity               string   `json:"clientsecurity,omitempty"`
	Clientsecuritygroup          string   `json:"clientsecuritygroup,omitempty"`
	Clientsecuritymessage        string   `json:"clientsecuritymessage,omitempty"`
	Clientsecuritylog            string   `json:"clientsecuritylog,omitempty"`
	Smartgroup                   string   `json:"smartgroup,omitempty"`
	Splittunnel                  string   `json:"splittunnel,omitempty"`
	Locallanaccess               string   `json:"locallanaccess,omitempty"`
	Rfc1918                      string   `json:"rfc1918,omitempty"`
	Spoofiip                     string   `json:"spoofiip,omitempty"`
	Killconnections              string   `json:"killconnections,omitempty"`
	Transparentinterception      string   `json:"transparentinterception,omitempty"`
	Windowsclienttype            string   `json:"windowsclienttype,omitempty"`
	Defaultauthorizationaction   string   `json:"defaultauthorizationaction,omitempty"`
	Authorizationgroup           string   `json:"authorizationgroup,omitempty"`
	Clientidletimeout            int      `json:"clientidletimeout,omitempty"`
	Proxy                        string   `json:"proxy,omitempty"`
	Allprotocolproxy             string   `json:"allprotocolproxy,omitempty"`
	Httpproxy                    string   `json:"httpproxy,omitempty"`
	Ftpproxy                     string   `json:"ftpproxy,omitempty"`
	Socksproxy                   string   `json:"socksproxy,omitempty"`
	Gopherproxy                  string   `json:"gopherproxy,omitempty"`
	Sslproxy                     string   `json:"sslproxy,omitempty"`
	Proxyexception               string   `json:"proxyexception,omitempty"`
	Proxylocalbypass             string   `json:"proxylocalbypass,omitempty"`
	Clientcleanupprompt          string   `json:"clientcleanupprompt,omitempty"`
	Forcecleanup                 []string `json:"forcecleanup,omitempty"`
	Clientoptions                []string `json:"clientoptions,omitempty"`
	Clientconfiguration          []string `json:"clientconfiguration,omitempty"`
	Sso                          string   `json:"sso,omitempty"`
	Ssocredential                string   `json:"ssocredential,omitempty"`
	Windowsautologon             string   `json:"windowsautologon,omitempty"`
	Usemip                       string   `json:"usemip,omitempty"`
	Useiip                       string   `json:"useiip,omitempty"`
	Clientdebug                  string   `json:"clientdebug,omitempty"`
	Loginscript                  string   `json:"loginscript,omitempty"`
	Logoutscript                 string   `json:"logoutscript,omitempty"`
	Homepage                     string   `json:"homepage,omitempty"`
	Icaproxy                     string   `json:"icaproxy,omitempty"`
	Wihome                       string   `json:"wihome,omitempty"`
	Wihomeaddresstype            string   `json:"wihomeaddresstype,omitempty"`
	Citrixreceiverhome           string   `json:"citrixreceiverhome,omitempty"`
	Wiportalmode                 string   `json:"wiportalmode,omitempty"`
	Clientchoices                string   `json:"clientchoices,omitempty"`
	Epaclienttype                string   `json:"epaclienttype,omitempty"`
	Iipdnssuffix                 string   `json:"iipdnssuffix,omitempty"`
	Forcedtimeout                int      `json:"forcedtimeout,omitempty"`
	Forcedtimeoutwarning         int      `json:"forcedtimeoutwarning,omitempty"`
	Ntdomain                     string   `json:"ntdomain,omitempty"`
	Clientlessvpnmode            string   `json:"clientlessvpnmode,omitempty"`
	Clientlessmodeurlencoding    string   `json:"clientlessmodeurlencoding,omitempty"`
	Clientlesspersistentcookie   string   `json:"clientlesspersistentcookie,omitempty"`
	Emailhome                    string   `json:"emailhome,omitempty"`
	Allowedlogingroups           string   `json:"allowedlogingroups,omitempty"`
	Encryptcsecexp               string   `json:"encryptcsecexp,omitempty"`
	Apptokentimeout              int      `json:"apptokentimeout,omitempty"`
	Mdxtokentimeout              int      `json:"mdxtokentimeout,omitempty"`
	Uitheme                      string   `json:"uitheme,omitempty"`
	Securebrowse                 string   `json:"securebrowse,omitempty"`
	Storefronturl                string   `json:"storefronturl,omitempty"`
	Kcdaccount                   string   `json:"kcdaccount,omitempty"`
	Clientversions               string   `json:"clientversions,omitempty"`
	Rdpclientprofilename         string   `json:"rdpclientprofilename,omitempty"`
	Windowspluginupgrade         string   `json:"windowspluginupgrade,omitempty"`
	Macpluginupgrade             string   `json:"macpluginupgrade,omitempty"`
	Linuxpluginupgrade           string   `json:"linuxpluginupgrade,omitempty"`
	Iconwithreceiver             string   `json:"iconwithreceiver,omitempty"`
	Userdomains                  string   `json:"userdomains,omitempty"`
	Icasessiontimeout            string   `json:"icasessiontimeout,omitempty"`
	Httptrackconnproxy           string   `json:"httptrackconnproxy,omitempty"`
	Alwaysonprofilename          string   `json:"alwaysonprofilename,omitempty"`
	Autoproxyurl                 string   `json:"autoproxyurl,omitempty"`
	Advancedclientlessvpnmode    string   `json:"advancedclientlessvpnmode,omitempty"`
	Pcoipprofilename             string   `json:"pcoipprofilename,omitempty"`
	Backendserversni             string   `json:"backendserversni,omitempty"`
	Backendcertvalidation        string   `json:"backendcertvalidation,omitempty"`
	Secureprivateaccess          string   `json:"secureprivateaccess,omitempty"`
	Accessrestrictedpageredirect string   `json:"accessrestrictedpageredirect,omitempty"`
	Fqdnspoofedip                string   `json:"fqdnspoofedip,omitempty"`
	Netmask                      string   `json:"netmask,omitempty"`
	Samesite                     string   `json:"samesite,omitempty"`
	Maxiipperuser                int      `json:"maxiipperuser,omitempty"`
	Deviceposture                string   `json:"deviceposture,omitempty"`
	Backenddtls12                string   `json:"backenddtls12,omitempty"`
	Name                         string   `json:"name,omitempty"`
	Clientidletimeoutwarning     string   `json:"clientidletimeoutwarning,omitempty"`
	Vpnsessionpolicybindtype     string   `json:"vpnsessionpolicybindtype,omitempty"`
	Vpnsessionpolicycount        string   `json:"vpnsessionpolicycount,omitempty"`
	Nextgenapiresource           string   `json:"_nextgenapiresource,omitempty"`
}

type Vpnvserverauthenticationnegotiatepolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Name                   string `json:"name,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnvserverauthenticationradiuspolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Name                   string `json:"name,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnvserverclientlessaccesspolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
}

type Vpnvservernslogpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnclientlessaccesspolicyvpnglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Vpnformssoaction struct {
	Name               string `json:"name,omitempty"`
	Actionurl          string `json:"actionurl,omitempty"`
	Userfield          string `json:"userfield,omitempty"`
	Passwdfield        string `json:"passwdfield,omitempty"`
	Ssosuccessrule     string `json:"ssosuccessrule,omitempty"`
	Namevaluepair      string `json:"namevaluepair,omitempty"`
	Responsesize       int    `json:"responsesize,omitempty"`
	Nvtype             string `json:"nvtype,omitempty"`
	Submitmethod       string `json:"submitmethod,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Vpnnexthopserver struct {
	Name               string `json:"name,omitempty"`
	Nexthopip          string `json:"nexthopip,omitempty"`
	Nexthopfqdn        string `json:"nexthopfqdn,omitempty"`
	Resaddresstype     string `json:"resaddresstype,omitempty"`
	Nexthopport        int    `json:"nexthopport,omitempty"`
	Secure             string `json:"secure,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Vpntrafficpolicy struct {
	Name               string `json:"name,omitempty"`
	Rule               string `json:"rule,omitempty"`
	Action             string `json:"action,omitempty"`
	Expressiontype     string `json:"expressiontype,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Vpntrafficpolicyaaauserbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Vpnurlpolicyvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Vpnvserverappflowpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
}

type Vpnvservericapolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnglobalsslcertkeybinding struct {
	Certkeyname            string `json:"certkeyname,omitempty"`
	Crlcheck               string `json:"crlcheck,omitempty"`
	Ocspcheck              string `json:"ocspcheck,omitempty"`
	Userdataencryptionkey  string `json:"userdataencryptionkey,omitempty"`
	Cacert                 string `json:"cacert,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpnglobalsyslogpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpnvserverlocalpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Name                   string `json:"name,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnvservervpneulabinding struct {
	Eula    string `json:"eula,omitempty"`
	Acttype int    `json:"acttype,omitempty"`
	Name    string `json:"name,omitempty"`
}

type Vpnpcoipvserverprofile struct {
	Name               string `json:"name,omitempty"`
	Logindomain        string `json:"logindomain,omitempty"`
	Udpport            int    `json:"udpport,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Vpnvservernexthopserverbinding struct {
	Nexthopserver string `json:"nexthopserver,omitempty"`
	Acttype       uint32 `json:"acttype,omitempty"`
	Name          string `json:"name,omitempty"`
}

type Vpnvserverstaserverbinding struct {
	Staserver      string `json:"staserver,omitempty"`
	Staauthid      string `json:"staauthid,omitempty"`
	Stastate       string `json:"stastate,omitempty"`
	Acttype        int    `json:"acttype,omitempty"`
	Staaddresstype string `json:"staaddresstype,omitempty"`
	Name           string `json:"name,omitempty"`
}

type Vpnvserversyslogpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnvserverintranetip6binding struct {
	Intranetip6 string `json:"intranetip6,omitempty"`
	Numaddr     int    `json:"numaddr,omitempty"`
	Acttype     int    `json:"acttype,omitempty"`
	Name        string `json:"name,omitempty"`
}

type Vpnvservernegotiatepolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Name                   string `json:"name,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnvserverurlpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnvserverwebauthpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Name                   string `json:"name,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnsessionpolicyglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Vpntrafficpolicyglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Vpntrafficpolicyvpnvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Vpnvserverauditnslogpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnvservercachepolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
}

type Vpnvserversharefileserverbinding struct {
	Sharefile string `json:"sharefile,omitempty"`
	Acttype   int    `json:"acttype,omitempty"`
	Name      string `json:"name,omitempty"`
}

type Vpnurlpolicy struct {
	Name               string `json:"name,omitempty"`
	Rule               string `json:"rule,omitempty"`
	Action             string `json:"action,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Logaction          string `json:"logaction,omitempty"`
	Newname            string `json:"newname,omitempty"`
	Builtin            string `json:"builtin,omitempty"`
	Undefhits          string `json:"undefhits,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Vpnurlpolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Vpnvserverappcontrollerbinding struct {
	Appcontroller string `json:"appcontroller,omitempty"`
	Acttype       int    `json:"acttype,omitempty"`
	Name          string `json:"name,omitempty"`
}

type Vpnglobalcertkeybinding struct {
	Certkeyname            string `json:"certkeyname,omitempty"`
	Crlcheck               string `json:"crlcheck,omitempty"`
	Ocspcheck              string `json:"ocspcheck,omitempty"`
	Userdataencryptionkey  string `json:"userdataencryptionkey,omitempty"`
	Cacert                 string `json:"cacert,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpnglobalvpntrafficpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
}

type Vpntrafficaction struct {
	Name               string `json:"name,omitempty"`
	Qual               string `json:"qual,omitempty"`
	Apptimeout         int    `json:"apptimeout,omitempty"`
	Sso                string `json:"sso,omitempty"`
	Hdx                string `json:"hdx,omitempty"`
	Formssoaction      string `json:"formssoaction,omitempty"`
	Fta                string `json:"fta,omitempty"`
	Wanscaler          string `json:"wanscaler,omitempty"`
	Kcdaccount         string `json:"kcdaccount,omitempty"`
	Samlssoprofile     string `json:"samlssoprofile,omitempty"`
	Proxy              string `json:"proxy,omitempty"`
	Userexpression     string `json:"userexpression,omitempty"`
	Passwdexpression   string `json:"passwdexpression,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Vpntrafficpolicyaaagroupbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Vpnvserverintranetapplicationbinding struct {
	Intranetapplication string `json:"intranetapplication,omitempty"`
	Acttype             uint32 `json:"acttype,omitempty"`
	Name                string `json:"name,omitempty"`
}

type Vpnicaconnection struct {
	Username           string `json:"username,omitempty"`
	Transproto         string `json:"transproto,omitempty"`
	Nodeid             int    `json:"nodeid,omitempty"`
	All                bool   `json:"all,omitempty"`
	Domain             string `json:"domain,omitempty"`
	Srcip              string `json:"srcip,omitempty"`
	Srcport            string `json:"srcport,omitempty"`
	Destip             string `json:"destip,omitempty"`
	Destport           string `json:"destport,omitempty"`
	Peid               string `json:"peid,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Vpnvserverauthenticationpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnvserverpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
}

type Vpnvservertrafficpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnglobalauthenticationradiuspolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpntrafficpolicyuserbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Vpnvserverauthenticationloginschemapolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnvserverportalthemebinding struct {
	Portaltheme string `json:"portaltheme,omitempty"`
	Acttype     uint32 `json:"acttype,omitempty"`
	Name        string `json:"name,omitempty"`
}

type Vpnglobalcertpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpnvserverauthenticationlocalpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Name                   string `json:"name,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnglobalvpnintranetapplicationbinding struct {
	Intranetapplication    string `json:"intranetapplication,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpnclientlessaccesspolicy struct {
	Name               string `json:"name,omitempty"`
	Rule               string `json:"rule,omitempty"`
	Profilename        string `json:"profilename,omitempty"`
	Undefaction        string `json:"undefaction,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Undefhits          string `json:"undefhits,omitempty"`
	Description        string `json:"description,omitempty"`
	Isdefault          string `json:"isdefault,omitempty"`
	Builtin            string `json:"builtin,omitempty"`
	Feature            string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Vpnclientlessaccesspolicyvpnvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Vpnurlpolicygroupbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Vpnurlpolicyuserbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Vpnvserverepaprofilebinding struct {
	Epaprofile         string `json:"epaprofile,omitempty"`
	Acttype            uint32 `json:"acttype,omitempty"`
	Name               string `json:"name,omitempty"`
	Epaprofileoptional bool   `json:"epaprofileoptional,omitempty"`
}

type Vpnvservertacacspolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Name                   string `json:"name,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnicadtlsconnection struct {
	Username           string `json:"username,omitempty"`
	Nodeid             int    `json:"nodeid,omitempty"`
	Domain             string `json:"domain,omitempty"`
	Srcip              string `json:"srcip,omitempty"`
	Srcport            string `json:"srcport,omitempty"`
	Destip             string `json:"destip,omitempty"`
	Destport           string `json:"destport,omitempty"`
	Channelnumber      string `json:"channelnumber,omitempty"`
	Peid               string `json:"peid,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Vpnvserveraaapreauthenticationpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnvserverdfapolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Name                   string `json:"name,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnvserverpreauthenticationpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnvservervpnclientlessaccesspolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
}

type Vpnvservervpntrafficpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnclientlessaccesspolicyglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy int32  `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Vpnglobaldomainbinding struct {
	Intranetdomain         string `json:"intranetdomain,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpnglobalvpnnexthopserverbinding struct {
	Nexthopserver          string `json:"nexthopserver,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpnsessionaction struct {
	Name                       string   `json:"name,omitempty"`
	Useraccounting             string   `json:"useraccounting,omitempty"`
	Httpport                   []int    `json:"httpport,omitempty"`
	Winsip                     string   `json:"winsip,omitempty"`
	Dnsvservername             string   `json:"dnsvservername,omitempty"`
	Splitdns                   string   `json:"splitdns,omitempty"`
	Sesstimeout                int      `json:"sesstimeout,omitempty"`
	Clientsecurity             string   `json:"clientsecurity,omitempty"`
	Clientsecuritygroup        string   `json:"clientsecuritygroup,omitempty"`
	Clientsecuritymessage      string   `json:"clientsecuritymessage,omitempty"`
	Clientsecuritylog          string   `json:"clientsecuritylog,omitempty"`
	Splittunnel                string   `json:"splittunnel,omitempty"`
	Locallanaccess             string   `json:"locallanaccess,omitempty"`
	Rfc1918                    string   `json:"rfc1918,omitempty"`
	Spoofiip                   string   `json:"spoofiip,omitempty"`
	Killconnections            string   `json:"killconnections,omitempty"`
	Transparentinterception    string   `json:"transparentinterception,omitempty"`
	Windowsclienttype          string   `json:"windowsclienttype,omitempty"`
	Defaultauthorizationaction string   `json:"defaultauthorizationaction,omitempty"`
	Authorizationgroup         string   `json:"authorizationgroup,omitempty"`
	Smartgroup                 string   `json:"smartgroup,omitempty"`
	Clientidletimeout          int      `json:"clientidletimeout,omitempty"`
	Proxy                      string   `json:"proxy,omitempty"`
	Allprotocolproxy           string   `json:"allprotocolproxy,omitempty"`
	Httpproxy                  string   `json:"httpproxy,omitempty"`
	Ftpproxy                   string   `json:"ftpproxy,omitempty"`
	Socksproxy                 string   `json:"socksproxy,omitempty"`
	Gopherproxy                string   `json:"gopherproxy,omitempty"`
	Sslproxy                   string   `json:"sslproxy,omitempty"`
	Proxyexception             string   `json:"proxyexception,omitempty"`
	Proxylocalbypass           string   `json:"proxylocalbypass,omitempty"`
	Clientcleanupprompt        string   `json:"clientcleanupprompt,omitempty"`
	Forcecleanup               []string `json:"forcecleanup,omitempty"`
	Clientoptions              string   `json:"clientoptions,omitempty"`
	Clientconfiguration        []string `json:"clientconfiguration,omitempty"`
	Sso                        string   `json:"sso,omitempty"`
	Ssocredential              string   `json:"ssocredential,omitempty"`
	Windowsautologon           string   `json:"windowsautologon,omitempty"`
	Usemip                     string   `json:"usemip,omitempty"`
	Useiip                     string   `json:"useiip,omitempty"`
	Clientdebug                string   `json:"clientdebug,omitempty"`
	Loginscript                string   `json:"loginscript,omitempty"`
	Logoutscript               string   `json:"logoutscript,omitempty"`
	Homepage                   string   `json:"homepage,omitempty"`
	Icaproxy                   string   `json:"icaproxy,omitempty"`
	Wihome                     string   `json:"wihome,omitempty"`
	Wihomeaddresstype          string   `json:"wihomeaddresstype,omitempty"`
	Citrixreceiverhome         string   `json:"citrixreceiverhome,omitempty"`
	Wiportalmode               string   `json:"wiportalmode,omitempty"`
	Clientchoices              string   `json:"clientchoices,omitempty"`
	Epaclienttype              string   `json:"epaclienttype,omitempty"`
	Iipdnssuffix               string   `json:"iipdnssuffix,omitempty"`
	Forcedtimeout              int      `json:"forcedtimeout,omitempty"`
	Forcedtimeoutwarning       int      `json:"forcedtimeoutwarning,omitempty"`
	Ntdomain                   string   `json:"ntdomain,omitempty"`
	Clientlessvpnmode          string   `json:"clientlessvpnmode,omitempty"`
	Emailhome                  string   `json:"emailhome,omitempty"`
	Clientlessmodeurlencoding  string   `json:"clientlessmodeurlencoding,omitempty"`
	Clientlesspersistentcookie string   `json:"clientlesspersistentcookie,omitempty"`
	Allowedlogingroups         string   `json:"allowedlogingroups,omitempty"`
	Securebrowse               string   `json:"securebrowse,omitempty"`
	Storefronturl              string   `json:"storefronturl,omitempty"`
	Sfgatewayauthtype          string   `json:"sfgatewayauthtype,omitempty"`
	Kcdaccount                 string   `json:"kcdaccount,omitempty"`
	Rdpclientprofilename       string   `json:"rdpclientprofilename,omitempty"`
	Windowspluginupgrade       string   `json:"windowspluginupgrade,omitempty"`
	Macpluginupgrade           string   `json:"macpluginupgrade,omitempty"`
	Linuxpluginupgrade         string   `json:"linuxpluginupgrade,omitempty"`
	Iconwithreceiver           string   `json:"iconwithreceiver,omitempty"`
	Alwaysonprofilename        string   `json:"alwaysonprofilename,omitempty"`
	Autoproxyurl               string   `json:"autoproxyurl,omitempty"`
	Advancedclientlessvpnmode  string   `json:"advancedclientlessvpnmode,omitempty"`
	Pcoipprofilename           string   `json:"pcoipprofilename,omitempty"`
	Fqdnspoofedip              string   `json:"fqdnspoofedip,omitempty"`
	Netmask                    string   `json:"netmask,omitempty"`
	Clientidletimeoutwarning   string   `json:"clientidletimeoutwarning,omitempty"`
	Builtin                    string   `json:"builtin,omitempty"`
	Feature                    string   `json:"feature,omitempty"`
	Nextgenapiresource         string   `json:"_nextgenapiresource,omitempty"`
}

type Vpnsessionpolicyaaagroupbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Vpnsessionpolicyuserbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Vpnurlpolicyaaauserbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Vpnvserverauthenticationtacacspolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Name                   string `json:"name,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnglobalauthenticationnegotiatepolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpnglobalauthenticationsamlpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpnglobalradiuspolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpnvservervpnsessionpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnglobalauthenticationcertpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpnsfconfig struct {
	Vserver            []string `json:"vserver,omitempty"`
	Filename           string   `json:"filename,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
}

type Vpnurlpolicyaaagroupbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Vpnvserversecureprivateaccessurlbinding struct {
	Secureprivateaccessurl string `json:"secureprivateaccessurl,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Vpnpcoipprofile struct {
	Name               string `json:"name,omitempty"`
	Conserverurl       string `json:"conserverurl,omitempty"`
	Icvverification    string `json:"icvverification,omitempty"`
	Sessionidletimeout int    `json:"sessionidletimeout,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Vpnsamlssoprofile struct {
	Name                        string `json:"name,omitempty"`
	Samlsigningcertname         string `json:"samlsigningcertname,omitempty"`
	Assertionconsumerserviceurl string `json:"assertionconsumerserviceurl,omitempty"`
	Relaystaterule              string `json:"relaystaterule,omitempty"`
	Sendpassword                string `json:"sendpassword,omitempty"`
	Samlissuername              string `json:"samlissuername,omitempty"`
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
	Samlspcertname              string `json:"samlspcertname,omitempty"`
	Encryptionalgorithm         string `json:"encryptionalgorithm,omitempty"`
	Skewtime                    int    `json:"skewtime,omitempty"`
	Signassertion               string `json:"signassertion,omitempty"`
	Signatureservice            string `json:"signatureservice,omitempty"`
	Nextgenapiresource          string `json:"_nextgenapiresource,omitempty"`
}

type Vpnurl struct {
	Urlname            string `json:"urlname,omitempty"`
	Linkname           string `json:"linkname,omitempty"`
	Actualurl          string `json:"actualurl,omitempty"`
	Vservername        string `json:"vservername,omitempty"`
	Clientlessaccess   string `json:"clientlessaccess,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Iconurl            string `json:"iconurl,omitempty"`
	Ssotype            string `json:"ssotype,omitempty"`
	Applicationtype    string `json:"applicationtype,omitempty"`
	Samlssoprofile     string `json:"samlssoprofile,omitempty"`
	Appjson            string `json:"appjson,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Vpnsessionpolicyvpnvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Vpntrafficpolicyvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Vpnvserverfeopolicybinding struct {
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policy                 string `json:"policy,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
}

type Vpnvserversessionpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnvserverurlbinding struct {
	Urlname string `json:"urlname,omitempty"`
	Acttype uint32 `json:"acttype,omitempty"`
	Name    string `json:"name,omitempty"`
}

type Vpnglobalappcontrollerbinding struct {
	Appcontroller          string `json:"appcontroller,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpnglobalsecureprivateaccessurlbinding struct {
	Secureprivateaccessurl string `json:"secureprivateaccessurl,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpnglobalvpnurlpolicybinding struct {
	Policyname             string   `json:"policyname,omitempty"`
	Priority               int      `json:"priority,omitempty"`
	Gotopriorityexpression string   `json:"gotopriorityexpression,omitempty"`
	Builtin                []string `json:"builtin,omitempty"`
	Secondary              bool     `json:"secondary,omitempty"`
	Groupextraction        bool     `json:"groupextraction,omitempty"`
}

type Vpnsessionpolicy struct {
	Name               string `json:"name,omitempty"`
	Rule               string `json:"rule,omitempty"`
	Action             string `json:"action,omitempty"`
	Builtin            string `json:"builtin,omitempty"`
	Feature            string `json:"feature,omitempty"`
	Expressiontype     string `json:"expressiontype,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Vpnsessionpolicyvpnglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Vpnvserveroauthidppolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnvserverradiuspolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Name                   string `json:"name,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnvserversamlpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Acttype                uint32 `json:"acttype,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Name                   string `json:"name,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnglobalauditnslogpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpnglobalintranetapplicationbinding struct {
	Intranetapplication    string `json:"intranetapplication,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpnglobalpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpnglobalurlbinding struct {
	Urlname                string `json:"urlname,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpnstoreinfo struct {
	Url                string `json:"url,omitempty"`
	Storeserverstatus  string `json:"storeserverstatus,omitempty"`
	Storeserverissf    string `json:"storeserverissf,omitempty"`
	Storeapisupport    string `json:"storeapisupport,omitempty"`
	Storelist          string `json:"storelist,omitempty"`
	Storestatus        string `json:"storestatus,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Vpnvserverresponderpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
}

type Vpnvservervpnurlpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnepaprofile struct {
	Name               string `json:"name,omitempty"`
	Filename           string `json:"filename,omitempty"`
	Data               string `json:"data,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Vpnglobalnegotiatepolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpnglobalnexthopserverbinding struct {
	Nexthopserver          string `json:"nexthopserver,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpnglobaltacacspolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpnglobalvpneulabinding struct {
	Eula                   string `json:"eula,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpnsessionpolicyaaauserbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Vpnvservercspolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnvservereulabinding struct {
	Eula    string `json:"eula,omitempty"`
	Acttype uint32 `json:"acttype,omitempty"`
	Name    string `json:"name,omitempty"`
}

type Vpnglobalsharefileserverbinding struct {
	Sharefile              string `json:"sharefile,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpnglobalvpnclientlessaccesspolicybinding struct {
	Policyname             string   `json:"policyname,omitempty"`
	Priority               int      `json:"priority,omitempty"`
	Type                   string   `json:"type,omitempty"`
	Gotopriorityexpression string   `json:"gotopriorityexpression,omitempty"`
	Builtin                []string `json:"builtin,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	Globalbindtype         string   `json:"globalbindtype,omitempty"`
	Secondary              bool     `json:"secondary,omitempty"`
	Groupextraction        bool     `json:"groupextraction,omitempty"`
}

type Vpnglobalvpnurlbinding struct {
	Urlname                string `json:"urlname,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpntrafficpolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Vpnvserverauthenticationldappolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Name                   string `json:"name,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnvserverauthenticationoauthidppolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnvserverauthenticationsamlidppolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnvserverauthenticationsamlpolicybinding struct {
	Policy                 string `json:"policy,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Acttype                int    `json:"acttype,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Name                   string `json:"name,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Bindpoint              string `json:"bindpoint,omitempty"`
}

type Vpnurlpolicyvpnglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Vpnvserveranalyticsprofilebinding struct {
	Analyticsprofile string `json:"analyticsprofile,omitempty"`
	Name             string `json:"name,omitempty"`
}

type Vpneula struct {
	Name               string `json:"name,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Vpnglobalauditsyslogpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Vpnglobalauthenticationldappolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Secondary              bool   `json:"secondary,omitempty"`
	Groupextraction        bool   `json:"groupextraction,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}
