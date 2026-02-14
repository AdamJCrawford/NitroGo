// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Botprofileblacklistbinding struct {
	Botblacklist        bool     `json:"bot_blacklist,omitempty"`
	Botblacklisttype    string   `json:"bot_blacklist_type,omitempty"`
	Botblacklistenabled string   `json:"bot_blacklist_enabled,omitempty"`
	Botblacklistvalue   string   `json:"bot_blacklist_value,omitempty"`
	Botblacklistaction  []string `json:"bot_blacklist_action,omitempty"`
	Logmessage          string   `json:"logmessage,omitempty"`
	Botbindcomment      string   `json:"bot_bind_comment,omitempty"`
	Name                string   `json:"name,omitempty"`
}

type Botsignature struct {
	Src                string `json:"src,omitempty"`
	Name               string `json:"name,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Overwrite          bool   `json:"overwrite,omitempty"`
	Response           string `json:"response,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Botpolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Botpolicyglobalbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Activepolicy           int32  `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Botprofilecaptchabinding struct {
	Captcharesource   bool     `json:"captcharesource,omitempty"`
	Botcaptchaurl     string   `json:"bot_captcha_url,omitempty"`
	Botcaptchaenabled string   `json:"bot_captcha_enabled,omitempty"`
	Waittime          int      `json:"waittime,omitempty"`
	Graceperiod       int      `json:"graceperiod,omitempty"`
	Muteperiod        int      `json:"muteperiod,omitempty"`
	Requestsizelimit  int      `json:"requestsizelimit,omitempty"`
	Retryattempts     int      `json:"retryattempts,omitempty"`
	Botcaptchaaction  []string `json:"bot_captcha_action,omitempty"`
	Logmessage        string   `json:"logmessage,omitempty"`
	Botbindcomment    string   `json:"bot_bind_comment,omitempty"`
	Name              string   `json:"name,omitempty"`
}

type Botprofiletpsbinding struct {
	Bottps         bool     `json:"bot_tps,omitempty"`
	Bottpstype     string   `json:"bot_tps_type,omitempty"`
	Threshold      int      `json:"threshold,omitempty"`
	Percentage     int      `json:"percentage,omitempty"`
	Bottpsaction   []string `json:"bot_tps_action,omitempty"`
	Logmessage     string   `json:"logmessage,omitempty"`
	Botbindcomment string   `json:"bot_bind_comment,omitempty"`
	Bottpsenabled  string   `json:"bot_tps_enabled,omitempty"`
	Name           string   `json:"name,omitempty"`
}

type Botpolicybotpolicylabelbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Botpolicyvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Activepolicy           int32  `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Botpolicylabelbotpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Botglobalpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Type                   string `json:"type,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Numpol                 uint32 `json:"numpol,omitempty"`
	Flowtype               uint32 `json:"flowtype,omitempty"`
	Globalbindtype         string `json:"globalbindtype,omitempty"`
}

type Botpolicycsvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Botprofilelogexpressionbinding struct {
	Logexpression           bool   `json:"logexpression,omitempty"`
	Botlogexpressionname    string `json:"bot_log_expression_name,omitempty"`
	Botlogexpressionvalue   string `json:"bot_log_expression_value,omitempty"`
	Botlogexpressionenabled string `json:"bot_log_expression_enabled,omitempty"`
	Botbindcomment          string `json:"bot_bind_comment,omitempty"`
	Name                    string `json:"name,omitempty"`
	Logmessage              string `json:"logmessage,omitempty"`
}

type Botpolicylabelpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Botpolicylabelpolicybindingbinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Botprofileratelimitbinding struct {
	Botratelimit        bool     `json:"bot_ratelimit,omitempty"`
	Botratelimittype    string   `json:"bot_rate_limit_type,omitempty"`
	Botratelimitenabled string   `json:"bot_rate_limit_enabled,omitempty"`
	Botratelimiturl     string   `json:"bot_rate_limit_url,omitempty"`
	Cookiename          string   `json:"cookiename,omitempty"`
	Countrycode         string   `json:"countrycode,omitempty"`
	Rate                int      `json:"rate,omitempty"`
	Timeslice           int      `json:"timeslice,omitempty"`
	Limittype           string   `json:"limittype,omitempty"`
	Condition           string   `json:"condition,omitempty"`
	Botratelimitaction  []string `json:"bot_rate_limit_action,omitempty"`
	Logmessage          string   `json:"logmessage,omitempty"`
	Botbindcomment      string   `json:"bot_bind_comment,omitempty"`
	Name                string   `json:"name,omitempty"`
}

type Botsettings struct {
	Defaultprofile             string `json:"defaultprofile,omitempty"`
	Defaultnonintrusiveprofile string `json:"defaultnonintrusiveprofile,omitempty"`
	Javascriptname             string `json:"javascriptname,omitempty"`
	Sessiontimeout             int    `json:"sessiontimeout,omitempty"`
	Sessioncookiename          string `json:"sessioncookiename,omitempty"`
	Dfprequestlimit            int    `json:"dfprequestlimit,omitempty"`
	Signatureautoupdate        string `json:"signatureautoupdate,omitempty"`
	Signatureurl               string `json:"signatureurl,omitempty"`
	Proxyserver                string `json:"proxyserver,omitempty"`
	Proxyport                  int    `json:"proxyport,omitempty"`
	Trapurlautogenerate        string `json:"trapurlautogenerate,omitempty"`
	Trapurlinterval            int    `json:"trapurlinterval,omitempty"`
	Trapurllength              int    `json:"trapurllength,omitempty"`
	Proxyusername              string `json:"proxyusername,omitempty"`
	Proxypassword              string `json:"proxypassword,omitempty"`
	Builtin                    string `json:"builtin,omitempty"`
	Feature                    string `json:"feature,omitempty"`
	Nextgenapiresource         string `json:"_nextgenapiresource,omitempty"`
}

type Botpolicylabel struct {
	Labelname          string `json:"labelname,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Newname            string `json:"newname,omitempty"`
	Numpol             string `json:"numpol,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Botpolicylbvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Botprofilebinding struct {
	Name string `json:"name,omitempty"`
}

type Botprofileipreputationbinding struct {
	Botipreputation bool     `json:"bot_ipreputation,omitempty"`
	Category        string   `json:"category,omitempty"`
	Botiprepenabled string   `json:"bot_iprep_enabled,omitempty"`
	Botiprepaction  []string `json:"bot_iprep_action,omitempty"`
	Logmessage      string   `json:"logmessage,omitempty"`
	Botbindcomment  string   `json:"bot_bind_comment,omitempty"`
	Name            string   `json:"name,omitempty"`
}

type Botprofiletrapinsertionurlbinding struct {
	Trapinsertionurl           bool   `json:"trapinsertionurl,omitempty"`
	Bottrapurl                 string `json:"bot_trap_url,omitempty"`
	Bottrapurlinsertionenabled string `json:"bot_trap_url_insertion_enabled,omitempty"`
	Botbindcomment             string `json:"bot_bind_comment,omitempty"`
	Name                       string `json:"name,omitempty"`
	Logmessage                 string `json:"logmessage,omitempty"`
}

type Botprofilewhitelistbinding struct {
	Botwhitelist        bool   `json:"bot_whitelist,omitempty"`
	Botwhitelisttype    string `json:"bot_whitelist_type,omitempty"`
	Botwhitelistenabled string `json:"bot_whitelist_enabled,omitempty"`
	Botwhitelistvalue   string `json:"bot_whitelist_value,omitempty"`
	Log                 string `json:"log,omitempty"`
	Logmessage          string `json:"logmessage,omitempty"`
	Botbindcomment      string `json:"bot_bind_comment,omitempty"`
	Name                string `json:"name,omitempty"`
}

type Botpolicylabelbinding struct {
	Labelname string `json:"labelname,omitempty"`
}

type Botglobalbinding struct {
}

type Botglobalbotpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Type                   string `json:"type,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Numpol                 int    `json:"numpol,omitempty"`
	Flowtype               int    `json:"flowtype,omitempty"`
	Globalbindtype         string `json:"globalbindtype,omitempty"`
}

type Botpolicy struct {
	Name               string `json:"name,omitempty"`
	Rule               string `json:"rule,omitempty"`
	Profilename        string `json:"profilename,omitempty"`
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

type Botprofilekmdetectionexprbinding struct {
	Kmdetectionexpr       bool   `json:"kmdetectionexpr,omitempty"`
	Botkmexpressionname   string `json:"bot_km_expression_name,omitempty"`
	Botkmexpressionvalue  string `json:"bot_km_expression_value,omitempty"`
	Botkmdetectionenabled string `json:"bot_km_detection_enabled,omitempty"`
	Botbindcomment        string `json:"bot_bind_comment,omitempty"`
	Name                  string `json:"name,omitempty"`
	Logmessage            string `json:"logmessage,omitempty"`
}

type Botpolicybotglobalbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Botpolicypolicylabelbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Activepolicy           int32  `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Botprofile struct {
	Name                                   string   `json:"name,omitempty"`
	Signature                              string   `json:"signature,omitempty"`
	Errorurl                               string   `json:"errorurl,omitempty"`
	Trapurl                                string   `json:"trapurl,omitempty"`
	Comment                                string   `json:"comment,omitempty"`
	Botenablewhitelist                     string   `json:"bot_enable_white_list,omitempty"`
	Botenableblacklist                     string   `json:"bot_enable_black_list,omitempty"`
	Botenableratelimit                     string   `json:"bot_enable_rate_limit,omitempty"`
	Devicefingerprint                      string   `json:"devicefingerprint,omitempty"`
	Devicefingerprintaction                []string `json:"devicefingerprintaction,omitempty"`
	Botenableipreputation                  string   `json:"bot_enable_ip_reputation,omitempty"`
	Trap                                   string   `json:"trap,omitempty"`
	Trapaction                             []string `json:"trapaction,omitempty"`
	Signaturenouseragentheaderaction       []string `json:"signaturenouseragentheaderaction,omitempty"`
	Signaturemultipleuseragentheaderaction []string `json:"signaturemultipleuseragentheaderaction,omitempty"`
	Botenabletps                           string   `json:"bot_enable_tps,omitempty"`
	Devicefingerprintmobile                []string `json:"devicefingerprintmobile,omitempty"`
	Headlessbrowserdetection               string   `json:"headlessbrowserdetection,omitempty"`
	Clientipexpression                     string   `json:"clientipexpression,omitempty"`
	Kmjavascriptname                       string   `json:"kmjavascriptname,omitempty"`
	Kmdetection                            string   `json:"kmdetection,omitempty"`
	Kmeventspostbodylimit                  int      `json:"kmeventspostbodylimit,omitempty"`
	Verboseloglevel                        string   `json:"verboseloglevel,omitempty"`
	Spoofedreqaction                       []string `json:"spoofedreqaction,omitempty"`
	Dfprequestlimit                        int      `json:"dfprequestlimit,omitempty"`
	Sessioncookiename                      string   `json:"sessioncookiename,omitempty"`
	Sessiontimeout                         int      `json:"sessiontimeout,omitempty"`
	Addcookieflags                         string   `json:"addcookieflags,omitempty"`
	Builtin                                string   `json:"builtin,omitempty"`
	Feature                                string   `json:"feature,omitempty"`
	Nextgenapiresource                     string   `json:"_nextgenapiresource,omitempty"`
}
