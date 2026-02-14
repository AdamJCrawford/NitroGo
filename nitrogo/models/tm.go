// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Tmsessionpolicyauthenticationvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Tmsessionpolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Tmtrafficpolicycsvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Tmglobalauditsyslogpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Bindpolicytype         int    `json:"bindpolicytype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Tmglobalnslogpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Bindpolicytype         uint32 `json:"bindpolicytype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Tmglobaltrafficpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Bindpolicytype         uint32 `json:"bindpolicytype,omitempty"`
	Type                   string `json:"type,omitempty"`
	Globalbindtype         string `json:"globalbindtype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Tmsessionaction struct {
	Name                       string `json:"name,omitempty"`
	Sesstimeout                int    `json:"sesstimeout,omitempty"`
	Defaultauthorizationaction string `json:"defaultauthorizationaction,omitempty"`
	Sso                        string `json:"sso,omitempty"`
	Ssocredential              string `json:"ssocredential,omitempty"`
	Ssodomain                  string `json:"ssodomain,omitempty"`
	Httponlycookie             string `json:"httponlycookie,omitempty"`
	Kcdaccount                 string `json:"kcdaccount,omitempty"`
	Persistentcookie           string `json:"persistentcookie,omitempty"`
	Persistentcookievalidity   int    `json:"persistentcookievalidity,omitempty"`
	Homepage                   string `json:"homepage,omitempty"`
	Builtin                    string `json:"builtin,omitempty"`
	Feature                    string `json:"feature,omitempty"`
	Nextgenapiresource         string `json:"_nextgenapiresource,omitempty"`
}

type Tmtrafficaction struct {
	Name               string `json:"name,omitempty"`
	Apptimeout         int    `json:"apptimeout,omitempty"`
	Sso                string `json:"sso,omitempty"`
	Formssoaction      string `json:"formssoaction,omitempty"`
	Persistentcookie   string `json:"persistentcookie,omitempty"`
	Initiatelogout     string `json:"initiatelogout,omitempty"`
	Kcdaccount         string `json:"kcdaccount,omitempty"`
	Samlssoprofile     string `json:"samlssoprofile,omitempty"`
	Forcedtimeout      string `json:"forcedtimeout,omitempty"`
	Forcedtimeoutval   int    `json:"forcedtimeoutval,omitempty"`
	Userexpression     string `json:"userexpression,omitempty"`
	Passwdexpression   string `json:"passwdexpression,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Tmtrafficpolicytmglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Tmformssoaction struct {
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

type Tmglobalsyslogpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Bindpolicytype         uint32 `json:"bindpolicytype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Tmsessionpolicyglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Tmsessionpolicytmglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Tmsessionpolicy struct {
	Name                   string `json:"name,omitempty"`
	Rule                   string `json:"rule,omitempty"`
	Action                 string `json:"action,omitempty"`
	Builtin                string `json:"builtin,omitempty"`
	Feature                string `json:"feature,omitempty"`
	Expressiontype         string `json:"expressiontype,omitempty"`
	Hits                   string `json:"hits,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Nextgenapiresource     string `json:"_nextgenapiresource,omitempty"`
}

type Tmsessionpolicyaaauserbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Tmsessionpolicyuserbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Tmtrafficpolicyvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Tmglobalauditnslogpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Bindpolicytype         int    `json:"bindpolicytype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Tmsessionparameter struct {
	Sesstimeout                int    `json:"sesstimeout,omitempty"`
	Defaultauthorizationaction string `json:"defaultauthorizationaction,omitempty"`
	Sso                        string `json:"sso,omitempty"`
	Ssocredential              string `json:"ssocredential,omitempty"`
	Ssodomain                  string `json:"ssodomain,omitempty"`
	Kcdaccount                 string `json:"kcdaccount,omitempty"`
	Httponlycookie             string `json:"httponlycookie,omitempty"`
	Persistentcookie           string `json:"persistentcookie,omitempty"`
	Persistentcookievalidity   int    `json:"persistentcookievalidity,omitempty"`
	Homepage                   string `json:"homepage,omitempty"`
	Name                       string `json:"name,omitempty"`
	Tmsessionpolicybindtype    string `json:"tmsessionpolicybindtype,omitempty"`
	Tmsessionpolicycount       string `json:"tmsessionpolicycount,omitempty"`
	Nextgenapiresource         string `json:"_nextgenapiresource,omitempty"`
}

type Tmtrafficpolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Tmtrafficpolicylbvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Tmglobaltmsessionpolicybinding struct {
	Policyname             string   `json:"policyname,omitempty"`
	Priority               int      `json:"priority,omitempty"`
	Bindpolicytype         int      `json:"bindpolicytype,omitempty"`
	Builtin                []string `json:"builtin,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	Gotopriorityexpression string   `json:"gotopriorityexpression,omitempty"`
}

type Tmsamlssoprofile struct {
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
	Nextgenapiresource          string `json:"_nextgenapiresource,omitempty"`
}

type Tmsessionpolicyaaagroupbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     int    `json:"priority,omitempty"`
	Activepolicy int    `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Tmglobalbinding struct {
}

type Tmglobalsessionpolicybinding struct {
	Policyname             string   `json:"policyname,omitempty"`
	Priority               uint32   `json:"priority,omitempty"`
	Bindpolicytype         uint32   `json:"bindpolicytype,omitempty"`
	Builtin                []string `json:"builtin,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	Gotopriorityexpression string   `json:"gotopriorityexpression,omitempty"`
}

type Tmsessionpolicygroupbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Tmtrafficpolicyglobalbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Tmglobaltmtrafficpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Bindpolicytype         int    `json:"bindpolicytype,omitempty"`
	Type                   string `json:"type,omitempty"`
	Globalbindtype         string `json:"globalbindtype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
}

type Tmsessionpolicyvserverbinding struct {
	Boundto      string `json:"boundto,omitempty"`
	Priority     uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Tmtrafficpolicy struct {
	Name               string `json:"name,omitempty"`
	Rule               string `json:"rule,omitempty"`
	Action             string `json:"action,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}
