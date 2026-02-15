package models

// tm configuration structs
type Tmformssoaction struct {
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

type TmtrafficpolicyLbvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type Tmtrafficpolicy struct {
	Action             string  `json:"action,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Hits               int     `json:"hits,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Rule               string  `json:"rule,omitempty"`
}

type TmglobalTmtrafficpolicyBinding struct {
	Bindpolicytype         int    `json:"bindpolicytype,omitempty"`
	Globalbindtype         string `json:"globalbindtype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
}

type TmglobalBinding struct {
	TmglobalAuditnslogpolicyBinding  []interface{} `json:"tmglobal_auditnslogpolicy_binding,omitempty"`
	TmglobalAuditsyslogpolicyBinding []interface{} `json:"tmglobal_auditsyslogpolicy_binding,omitempty"`
	TmglobalTmsessionpolicyBinding   []interface{} `json:"tmglobal_tmsessionpolicy_binding,omitempty"`
	TmglobalTmtrafficpolicyBinding   []interface{} `json:"tmglobal_tmtrafficpolicy_binding,omitempty"`
}

type TmtrafficpolicyBinding struct {
	Name                            string        `json:"name,omitempty"`
	TmtrafficpolicyCsvserverBinding []interface{} `json:"tmtrafficpolicy_csvserver_binding,omitempty"`
	TmtrafficpolicyLbvserverBinding []interface{} `json:"tmtrafficpolicy_lbvserver_binding,omitempty"`
	TmtrafficpolicyTmglobalBinding  []interface{} `json:"tmtrafficpolicy_tmglobal_binding,omitempty"`
}

type Tmsessionpolicy struct {
	Action                 string   `json:"action,omitempty"`
	Builtin                []string `json:"builtin,omitempty"`
	Count                  float64  `json:"__count,omitempty"`
	Expressiontype         string   `json:"expressiontype,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	Gotopriorityexpression string   `json:"gotopriorityexpression,omitempty"`
	Hits                   int      `json:"hits,omitempty"`
	Name                   string   `json:"name,omitempty"`
	Nextgenapiresource     string   `json:"_nextgenapiresource,omitempty"`
	Rule                   string   `json:"rule,omitempty"`
}

type TmtrafficpolicyCsvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type TmsessionpolicyTmglobalBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type Tmsessionparameter struct {
	Defaultauthorizationaction string `json:"defaultauthorizationaction,omitempty"`
	Homepage                   string `json:"homepage,omitempty"`
	Httponlycookie             string `json:"httponlycookie,omitempty"`
	Kcdaccount                 string `json:"kcdaccount,omitempty"`
	Name                       string `json:"name,omitempty"`
	Nextgenapiresource         string `json:"_nextgenapiresource,omitempty"`
	Persistentcookie           string `json:"persistentcookie,omitempty"`
	Persistentcookievalidity   int    `json:"persistentcookievalidity,omitempty"`
	Sesstimeout                int    `json:"sesstimeout,omitempty"`
	Sso                        string `json:"sso,omitempty"`
	Ssocredential              string `json:"ssocredential,omitempty"`
	Ssodomain                  string `json:"ssodomain,omitempty"`
	Tmsessionpolicybindtype    string `json:"tmsessionpolicybindtype,omitempty"`
	Tmsessionpolicycount       int    `json:"tmsessionpolicycount,omitempty"`
}

type TmglobalAuditnslogpolicyBinding struct {
	Bindpolicytype         int    `json:"bindpolicytype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Tmtrafficaction struct {
	Apptimeout         int     `json:"apptimeout,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Forcedtimeout      string  `json:"forcedtimeout,omitempty"`
	Forcedtimeoutval   int     `json:"forcedtimeoutval,omitempty"`
	Formssoaction      string  `json:"formssoaction,omitempty"`
	Initiatelogout     string  `json:"initiatelogout,omitempty"`
	Kcdaccount         string  `json:"kcdaccount,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Passwdexpression   string  `json:"passwdexpression,omitempty"`
	Persistentcookie   string  `json:"persistentcookie,omitempty"`
	Samlssoprofile     string  `json:"samlssoprofile,omitempty"`
	Sso                string  `json:"sso,omitempty"`
	Userexpression     string  `json:"userexpression,omitempty"`
}

type TmsessionpolicyBinding struct {
	Name                                        string        `json:"name,omitempty"`
	TmsessionpolicyAaagroupBinding              []interface{} `json:"tmsessionpolicy_aaagroup_binding,omitempty"`
	TmsessionpolicyAaauserBinding               []interface{} `json:"tmsessionpolicy_aaauser_binding,omitempty"`
	TmsessionpolicyAuthenticationvserverBinding []interface{} `json:"tmsessionpolicy_authenticationvserver_binding,omitempty"`
	TmsessionpolicyTmglobalBinding              []interface{} `json:"tmsessionpolicy_tmglobal_binding,omitempty"`
}

type TmsessionpolicyAaagroupBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type TmtrafficpolicyTmglobalBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type TmglobalTmsessionpolicyBinding struct {
	Bindpolicytype         int      `json:"bindpolicytype,omitempty"`
	Builtin                []string `json:"builtin,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	Gotopriorityexpression string   `json:"gotopriorityexpression,omitempty"`
	Policyname             string   `json:"policyname,omitempty"`
	Priority               int      `json:"priority,omitempty"`
}

type TmglobalAuditsyslogpolicyBinding struct {
	Bindpolicytype         int    `json:"bindpolicytype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Tmsamlssoprofile struct {
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
	Skewtime                    int     `json:"skewtime,omitempty"`
}

type Tmsessionaction struct {
	Builtin                    []string `json:"builtin,omitempty"`
	Count                      float64  `json:"__count,omitempty"`
	Defaultauthorizationaction string   `json:"defaultauthorizationaction,omitempty"`
	Feature                    string   `json:"feature,omitempty"`
	Homepage                   string   `json:"homepage,omitempty"`
	Httponlycookie             string   `json:"httponlycookie,omitempty"`
	Kcdaccount                 string   `json:"kcdaccount,omitempty"`
	Name                       string   `json:"name,omitempty"`
	Nextgenapiresource         string   `json:"_nextgenapiresource,omitempty"`
	Persistentcookie           string   `json:"persistentcookie,omitempty"`
	Persistentcookievalidity   int      `json:"persistentcookievalidity,omitempty"`
	Sesstimeout                int      `json:"sesstimeout,omitempty"`
	Sso                        string   `json:"sso,omitempty"`
	Ssocredential              string   `json:"ssocredential,omitempty"`
	Ssodomain                  string   `json:"ssodomain,omitempty"`
}

type TmsessionpolicyAuthenticationvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type TmsessionpolicyAaauserBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}
