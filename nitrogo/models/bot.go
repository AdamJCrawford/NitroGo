package models

// bot configuration structs
type Botprofile struct {
	Addcookieflags                         string   `json:"addcookieflags,omitempty"`
	BotEnableBlackList                     string   `json:"bot_enable_black_list,omitempty"`
	BotEnableIpReputation                  string   `json:"bot_enable_ip_reputation,omitempty"`
	BotEnableRateLimit                     string   `json:"bot_enable_rate_limit,omitempty"`
	BotEnableTps                           string   `json:"bot_enable_tps,omitempty"`
	BotEnableWhiteList                     string   `json:"bot_enable_white_list,omitempty"`
	Builtin                                []string `json:"builtin,omitempty"`
	Clientipexpression                     string   `json:"clientipexpression,omitempty"`
	Comment                                string   `json:"comment,omitempty"`
	Count                                  float64  `json:"__count,omitempty"`
	Devicefingerprint                      string   `json:"devicefingerprint,omitempty"`
	Devicefingerprintaction                []string `json:"devicefingerprintaction,omitempty"`
	Devicefingerprintmobile                []string `json:"devicefingerprintmobile,omitempty"`
	Dfprequestlimit                        int      `json:"dfprequestlimit,omitempty"`
	Errorurl                               string   `json:"errorurl,omitempty"`
	Feature                                string   `json:"feature,omitempty"`
	Headlessbrowserdetection               string   `json:"headlessbrowserdetection,omitempty"`
	Kmdetection                            string   `json:"kmdetection,omitempty"`
	Kmeventspostbodylimit                  int      `json:"kmeventspostbodylimit,omitempty"`
	Kmjavascriptname                       string   `json:"kmjavascriptname,omitempty"`
	Name                                   string   `json:"name,omitempty"`
	Nextgenapiresource                     string   `json:"_nextgenapiresource,omitempty"`
	Sessioncookiename                      string   `json:"sessioncookiename,omitempty"`
	Sessiontimeout                         int      `json:"sessiontimeout,omitempty"`
	Signature                              string   `json:"signature,omitempty"`
	Signaturemultipleuseragentheaderaction []string `json:"signaturemultipleuseragentheaderaction,omitempty"`
	Signaturenouseragentheaderaction       []string `json:"signaturenouseragentheaderaction,omitempty"`
	Spoofedreqaction                       []string `json:"spoofedreqaction,omitempty"`
	Trap                                   string   `json:"trap,omitempty"`
	Trapaction                             []string `json:"trapaction,omitempty"`
	Trapurl                                string   `json:"trapurl,omitempty"`
	Verboseloglevel                        string   `json:"verboseloglevel,omitempty"`
}

type BotpolicyCsvserverBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type BotprofileKmdetectionexprBinding struct {
	BotBindComment        string `json:"bot_bind_comment,omitempty"`
	BotKmDetectionEnabled string `json:"bot_km_detection_enabled,omitempty"`
	BotKmExpressionName   string `json:"bot_km_expression_name,omitempty"`
	BotKmExpressionValue  string `json:"bot_km_expression_value,omitempty"`
	Kmdetectionexpr       bool   `json:"kmdetectionexpr,omitempty"`
	Logmessage            string `json:"logmessage,omitempty"`
	Name                  string `json:"name,omitempty"`
}

type BotprofileIpreputationBinding struct {
	BotBindComment  string   `json:"bot_bind_comment,omitempty"`
	BotIprepAction  []string `json:"bot_iprep_action,omitempty"`
	BotIprepEnabled string   `json:"bot_iprep_enabled,omitempty"`
	BotIpreputation bool     `json:"bot_ipreputation,omitempty"`
	Category        string   `json:"category,omitempty"`
	Logmessage      string   `json:"logmessage,omitempty"`
	Name            string   `json:"name,omitempty"`
}

type Botsignature struct {
	Comment            string `json:"comment,omitempty"`
	Name               string `json:"name,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Overwrite          bool   `json:"overwrite,omitempty"`
	Response           string `json:"response,omitempty"`
	Src                string `json:"src,omitempty"`
}

type BotpolicyLbvserverBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Botsettings struct {
	Builtin                    []string `json:"builtin,omitempty"`
	Defaultnonintrusiveprofile string   `json:"defaultnonintrusiveprofile,omitempty"`
	Defaultprofile             string   `json:"defaultprofile,omitempty"`
	Dfprequestlimit            int      `json:"dfprequestlimit,omitempty"`
	Feature                    string   `json:"feature,omitempty"`
	Javascriptname             string   `json:"javascriptname,omitempty"`
	Nextgenapiresource         string   `json:"_nextgenapiresource,omitempty"`
	Proxypassword              string   `json:"proxypassword,omitempty"`
	Proxyport                  int      `json:"proxyport,omitempty"`
	Proxyserver                string   `json:"proxyserver,omitempty"`
	Proxyusername              string   `json:"proxyusername,omitempty"`
	Sessioncookiename          string   `json:"sessioncookiename,omitempty"`
	Sessiontimeout             int      `json:"sessiontimeout,omitempty"`
	Signatureautoupdate        string   `json:"signatureautoupdate,omitempty"`
	Signatureurl               string   `json:"signatureurl,omitempty"`
	Trapurlautogenerate        string   `json:"trapurlautogenerate,omitempty"`
	Trapurlinterval            int      `json:"trapurlinterval,omitempty"`
	Trapurllength              int      `json:"trapurllength,omitempty"`
}

type BotpolicylabelBinding struct {
	BotpolicylabelBotpolicyBinding     []interface{} `json:"botpolicylabel_botpolicy_binding,omitempty"`
	BotpolicylabelPolicybindingBinding []interface{} `json:"botpolicylabel_policybinding_binding,omitempty"`
	Labelname                          string        `json:"labelname,omitempty"`
}

type BotglobalBotpolicyBinding struct {
	Flowtype               int    `json:"flowtype,omitempty"`
	Globalbindtype         string `json:"globalbindtype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Numpol                 int    `json:"numpol,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
}

type BotpolicyBinding struct {
	BotpolicyBotglobalBinding      []interface{} `json:"botpolicy_botglobal_binding,omitempty"`
	BotpolicyBotpolicylabelBinding []interface{} `json:"botpolicy_botpolicylabel_binding,omitempty"`
	BotpolicyCsvserverBinding      []interface{} `json:"botpolicy_csvserver_binding,omitempty"`
	BotpolicyLbvserverBinding      []interface{} `json:"botpolicy_lbvserver_binding,omitempty"`
	Name                           string        `json:"name,omitempty"`
}

type BotprofileTrapinsertionurlBinding struct {
	BotBindComment             string `json:"bot_bind_comment,omitempty"`
	BotTrapUrl                 string `json:"bot_trap_url,omitempty"`
	BotTrapUrlInsertionEnabled string `json:"bot_trap_url_insertion_enabled,omitempty"`
	Logmessage                 string `json:"logmessage,omitempty"`
	Name                       string `json:"name,omitempty"`
	Trapinsertionurl           bool   `json:"trapinsertionurl,omitempty"`
}

type BotprofileBinding struct {
	BotprofileBlacklistBinding        []interface{} `json:"botprofile_blacklist_binding,omitempty"`
	BotprofileCaptchaBinding          []interface{} `json:"botprofile_captcha_binding,omitempty"`
	BotprofileIpreputationBinding     []interface{} `json:"botprofile_ipreputation_binding,omitempty"`
	BotprofileKmdetectionexprBinding  []interface{} `json:"botprofile_kmdetectionexpr_binding,omitempty"`
	BotprofileLogexpressionBinding    []interface{} `json:"botprofile_logexpression_binding,omitempty"`
	BotprofileRatelimitBinding        []interface{} `json:"botprofile_ratelimit_binding,omitempty"`
	BotprofileTpsBinding              []interface{} `json:"botprofile_tps_binding,omitempty"`
	BotprofileTrapinsertionurlBinding []interface{} `json:"botprofile_trapinsertionurl_binding,omitempty"`
	BotprofileWhitelistBinding        []interface{} `json:"botprofile_whitelist_binding,omitempty"`
	Name                              string        `json:"name,omitempty"`
}

type Botpolicy struct {
	Builtin            []string `json:"builtin,omitempty"`
	Comment            string   `json:"comment,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	Logaction          string   `json:"logaction,omitempty"`
	Name               string   `json:"name,omitempty"`
	Newname            string   `json:"newname,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Profilename        string   `json:"profilename,omitempty"`
	Rule               string   `json:"rule,omitempty"`
	Undefaction        string   `json:"undefaction,omitempty"`
	Undefhits          int      `json:"undefhits,omitempty"`
}

type BotpolicylabelBotpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type BotprofileWhitelistBinding struct {
	BotBindComment      string `json:"bot_bind_comment,omitempty"`
	BotWhitelist        bool   `json:"bot_whitelist,omitempty"`
	BotWhitelistEnabled string `json:"bot_whitelist_enabled,omitempty"`
	BotWhitelistType    string `json:"bot_whitelist_type,omitempty"`
	BotWhitelistValue   string `json:"bot_whitelist_value,omitempty"`
	Log                 string `json:"log,omitempty"`
	Logmessage          string `json:"logmessage,omitempty"`
	Name                string `json:"name,omitempty"`
}

type BotglobalBinding struct {
	BotglobalBotpolicyBinding []interface{} `json:"botglobal_botpolicy_binding,omitempty"`
}

type BotprofileLogexpressionBinding struct {
	BotBindComment          string `json:"bot_bind_comment,omitempty"`
	BotLogExpressionEnabled string `json:"bot_log_expression_enabled,omitempty"`
	BotLogExpressionName    string `json:"bot_log_expression_name,omitempty"`
	BotLogExpressionValue   string `json:"bot_log_expression_value,omitempty"`
	Logexpression           bool   `json:"logexpression,omitempty"`
	Logmessage              string `json:"logmessage,omitempty"`
	Name                    string `json:"name,omitempty"`
}

type Botpolicylabel struct {
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Hits               int     `json:"hits,omitempty"`
	Labelname          string  `json:"labelname,omitempty"`
	Newname            string  `json:"newname,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Numpol             int     `json:"numpol,omitempty"`
}

type BotprofileCaptchaBinding struct {
	BotBindComment    string   `json:"bot_bind_comment,omitempty"`
	BotCaptchaAction  []string `json:"bot_captcha_action,omitempty"`
	BotCaptchaEnabled string   `json:"bot_captcha_enabled,omitempty"`
	BotCaptchaUrl     string   `json:"bot_captcha_url,omitempty"`
	Captcharesource   bool     `json:"captcharesource,omitempty"`
	Graceperiod       int      `json:"graceperiod,omitempty"`
	Logmessage        string   `json:"logmessage,omitempty"`
	Muteperiod        int      `json:"muteperiod,omitempty"`
	Name              string   `json:"name,omitempty"`
	Requestsizelimit  int      `json:"requestsizelimit,omitempty"`
	Retryattempts     int      `json:"retryattempts,omitempty"`
	Waittime          int      `json:"waittime,omitempty"`
}

type BotpolicyBotglobalBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type BotpolicyBotpolicylabelBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type BotpolicylabelPolicybindingBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type BotprofileTpsBinding struct {
	BotBindComment string   `json:"bot_bind_comment,omitempty"`
	BotTps         bool     `json:"bot_tps,omitempty"`
	BotTpsAction   []string `json:"bot_tps_action,omitempty"`
	BotTpsEnabled  string   `json:"bot_tps_enabled,omitempty"`
	BotTpsType     string   `json:"bot_tps_type,omitempty"`
	Logmessage     string   `json:"logmessage,omitempty"`
	Name           string   `json:"name,omitempty"`
	Percentage     int      `json:"percentage,omitempty"`
	Threshold      int      `json:"threshold,omitempty"`
}

type BotprofileRatelimitBinding struct {
	BotBindComment      string   `json:"bot_bind_comment,omitempty"`
	BotRateLimitAction  []string `json:"bot_rate_limit_action,omitempty"`
	BotRateLimitEnabled string   `json:"bot_rate_limit_enabled,omitempty"`
	BotRateLimitType    string   `json:"bot_rate_limit_type,omitempty"`
	BotRateLimitUrl     string   `json:"bot_rate_limit_url,omitempty"`
	BotRatelimit        bool     `json:"bot_ratelimit,omitempty"`
	Condition           string   `json:"condition,omitempty"`
	Cookiename          string   `json:"cookiename,omitempty"`
	Countrycode         string   `json:"countrycode,omitempty"`
	Limittype           string   `json:"limittype,omitempty"`
	Logmessage          string   `json:"logmessage,omitempty"`
	Name                string   `json:"name,omitempty"`
	Rate                int      `json:"rate,omitempty"`
	Timeslice           int      `json:"timeslice,omitempty"`
}

type BotprofileBlacklistBinding struct {
	BotBindComment      string   `json:"bot_bind_comment,omitempty"`
	BotBlacklist        bool     `json:"bot_blacklist,omitempty"`
	BotBlacklistAction  []string `json:"bot_blacklist_action,omitempty"`
	BotBlacklistEnabled string   `json:"bot_blacklist_enabled,omitempty"`
	BotBlacklistType    string   `json:"bot_blacklist_type,omitempty"`
	BotBlacklistValue   string   `json:"bot_blacklist_value,omitempty"`
	Logmessage          string   `json:"logmessage,omitempty"`
	Name                string   `json:"name,omitempty"`
}
