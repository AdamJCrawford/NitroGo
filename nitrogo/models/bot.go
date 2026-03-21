package models

// bot configuration structs
type BotProfile struct {
	AddCookieFlags                         string   `json:"addcookieflags,omitempty"`
	BotEnableBlackList                     string   `json:"bot_enable_black_list,omitempty"`
	BotEnableIPReputation                  string   `json:"bot_enable_ip_reputation,omitempty"`
	BotEnableRateLimit                     string   `json:"bot_enable_rate_limit,omitempty"`
	BotEnableTPS                           string   `json:"bot_enable_tps,omitempty"`
	BotEnableWhiteList                     string   `json:"bot_enable_white_list,omitempty"`
	Builtin                                []string `json:"builtin,omitempty"`
	ClientIPExpression                     string   `json:"clientipexpression,omitempty"`
	Comment                                string   `json:"comment,omitempty"`
	Count                                  float64  `json:"__count,omitempty"`
	DeviceFingerprint                      string   `json:"devicefingerprint,omitempty"`
	DeviceFingerprintAction                []string `json:"devicefingerprintaction,omitempty"`
	DeviceFingerprintMobile                []string `json:"devicefingerprintmobile,omitempty"`
	DFPRequestLimit                        int      `json:"dfprequestlimit,omitempty"`
	ErrorURL                               string   `json:"errorurl,omitempty"`
	Feature                                string   `json:"feature,omitempty"`
	HeadlessBrowserDetection               string   `json:"headlessbrowserdetection,omitempty"`
	KMDetection                            string   `json:"kmdetection,omitempty"`
	KMEventsPostBodyLimit                  int      `json:"kmeventspostbodylimit,omitempty"`
	KMJavascriptName                       string   `json:"kmjavascriptname,omitempty"`
	Name                                   string   `json:"name,omitempty"`
	NextGenAPIResource                     string   `json:"_nextgenapiresource,omitempty"`
	SessionCookieName                      string   `json:"sessioncookiename,omitempty"`
	SessionTimeout                         int      `json:"sessiontimeout,omitempty"`
	Signature                              string   `json:"signature,omitempty"`
	SignatureMultipleUserAgentHeaderAction []string `json:"signaturemultipleuseragentheaderaction,omitempty"`
	SignatureNoUserAgentHeaderAction       []string `json:"signaturenouseragentheaderaction,omitempty"`
	SpoofedReqAction                       []string `json:"spoofedreqaction,omitempty"`
	Trap                                   string   `json:"trap,omitempty"`
	TrapAction                             []string `json:"trapaction,omitempty"`
	TrapURL                                string   `json:"trapurl,omitempty"`
	VerboseLogLevel                        string   `json:"verboseloglevel,omitempty"`
}

type BotPolicyCSVServerBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type BotProfileKMDetectionExprBinding struct {
	BotBindComment        string `json:"bot_bind_comment,omitempty"`
	BotKMDetectionEnabled string `json:"bot_km_detection_enabled,omitempty"`
	BotKMExpressionName   string `json:"bot_km_expression_name,omitempty"`
	BotKMExpressionValue  string `json:"bot_km_expression_value,omitempty"`
	KMDetectionExpr       bool   `json:"kmdetectionexpr,omitempty"`
	LogMessage            string `json:"logmessage,omitempty"`
	Name                  string `json:"name,omitempty"`
}

type BotProfileIPReputationBinding struct {
	BotBindComment  string   `json:"bot_bind_comment,omitempty"`
	BotIPRepAction  []string `json:"bot_iprep_action,omitempty"`
	BotIPRepEnabled string   `json:"bot_iprep_enabled,omitempty"`
	BotIPReputation bool     `json:"bot_ipreputation,omitempty"`
	Category        string   `json:"category,omitempty"`
	LogMessage      string   `json:"logmessage,omitempty"`
	Name            string   `json:"name,omitempty"`
}

type BotSignature struct {
	Comment            string `json:"comment,omitempty"`
	Name               string `json:"name,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	Overwrite          bool   `json:"overwrite,omitempty"`
	Response           string `json:"response,omitempty"`
	Src                string `json:"src,omitempty"`
}

type BotPolicyLBVServerBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type BotSettings struct {
	Builtin                    []string `json:"builtin,omitempty"`
	DefaultNonIntrusiveProfile string   `json:"defaultnonintrusiveprofile,omitempty"`
	DefaultProfile             string   `json:"defaultprofile,omitempty"`
	DFPRequestLimit            int      `json:"dfprequestlimit,omitempty"`
	Feature                    string   `json:"feature,omitempty"`
	JavascriptName             string   `json:"javascriptname,omitempty"`
	NextGenAPIResource         string   `json:"_nextgenapiresource,omitempty"`
	ProxyPassword              string   `json:"proxypassword,omitempty"`
	ProxyPort                  int      `json:"proxyport,omitempty"`
	ProxyServer                string   `json:"proxyserver,omitempty"`
	ProxyUsername              string   `json:"proxyusername,omitempty"`
	SessionCookieName          string   `json:"sessioncookiename,omitempty"`
	SessionTimeout             int      `json:"sessiontimeout,omitempty"`
	SignatureAutoUpdate        string   `json:"signatureautoupdate,omitempty"`
	SignatureURL               string   `json:"signatureurl,omitempty"`
	TrapURLAutoGenerate        string   `json:"trapurlautogenerate,omitempty"`
	TrapURLInterval            int      `json:"trapurlinterval,omitempty"`
	TrapURLLength              int      `json:"trapurllength,omitempty"`
}

type BotPolicyLabelBinding struct {
	BotPolicyLabelBotPolicyBinding     []any  `json:"botpolicylabel_botpolicy_binding,omitempty"`
	BotPolicyLabelPolicyBindingBinding []any  `json:"botpolicylabel_policybinding_binding,omitempty"`
	LabelName                          string `json:"labelname,omitempty"`
}

type BotGlobalBotPolicyBinding struct {
	FlowType               int    `json:"flowtype,omitempty"`
	GlobalBindType         string `json:"globalbindtype,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	NumPol                 int    `json:"numpol,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
}

type BotPolicyBinding struct {
	BotPolicyBotGlobalBinding      []any  `json:"botpolicy_botglobal_binding,omitempty"`
	BotPolicyBotPolicyLabelBinding []any  `json:"botpolicy_botpolicylabel_binding,omitempty"`
	BotPolicyCSVServerBinding      []any  `json:"botpolicy_csvserver_binding,omitempty"`
	BotPolicyLBVServerBinding      []any  `json:"botpolicy_lbvserver_binding,omitempty"`
	Name                           string `json:"name,omitempty"`
}

type BotProfileTrapInsertionURLBinding struct {
	BotBindComment             string `json:"bot_bind_comment,omitempty"`
	BotTrapURL                 string `json:"bot_trap_url,omitempty"`
	BotTrapURLInsertionEnabled string `json:"bot_trap_url_insertion_enabled,omitempty"`
	LogMessage                 string `json:"logmessage,omitempty"`
	Name                       string `json:"name,omitempty"`
	TrapInsertionURL           bool   `json:"trapinsertionurl,omitempty"`
}

type BotProfileBinding struct {
	BotProfileBlacklistBinding        []any  `json:"botprofile_blacklist_binding,omitempty"`
	BotProfileCaptchaBinding          []any  `json:"botprofile_captcha_binding,omitempty"`
	BotProfileIPReputationBinding     []any  `json:"botprofile_ipreputation_binding,omitempty"`
	BotProfileKMDetectionExprBinding  []any  `json:"botprofile_kmdetectionexpr_binding,omitempty"`
	BotProfileLogExpressionBinding    []any  `json:"botprofile_logexpression_binding,omitempty"`
	BotProfileRateLimitBinding        []any  `json:"botprofile_ratelimit_binding,omitempty"`
	BotProfileTPSBinding              []any  `json:"botprofile_tps_binding,omitempty"`
	BotProfileTrapInsertionURLBinding []any  `json:"botprofile_trapinsertionurl_binding,omitempty"`
	BotProfileWhitelistBinding        []any  `json:"botprofile_whitelist_binding,omitempty"`
	Name                              string `json:"name,omitempty"`
}

type BotPolicy struct {
	Builtin            []string `json:"builtin,omitempty"`
	Comment            string   `json:"comment,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	LogAction          string   `json:"logaction,omitempty"`
	Name               string   `json:"name,omitempty"`
	NewName            string   `json:"newname,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	ProfileName        string   `json:"profilename,omitempty"`
	Rule               string   `json:"rule,omitempty"`
	UndefAction        string   `json:"undefaction,omitempty"`
	UndefHits          int      `json:"undefhits,omitempty"`
}

type BotPolicyLabelBotPolicyBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelName        string `json:"invoke_labelname,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type BotProfileWhitelistBinding struct {
	BotBindComment      string `json:"bot_bind_comment,omitempty"`
	BotWhitelist        bool   `json:"bot_whitelist,omitempty"`
	BotWhitelistEnabled string `json:"bot_whitelist_enabled,omitempty"`
	BotWhitelistType    string `json:"bot_whitelist_type,omitempty"`
	BotWhitelistValue   string `json:"bot_whitelist_value,omitempty"`
	Log                 string `json:"log,omitempty"`
	LogMessage          string `json:"logmessage,omitempty"`
	Name                string `json:"name,omitempty"`
}

type BotGlobalBinding struct {
	BotGlobalBotPolicyBinding []any `json:"botglobal_botpolicy_binding,omitempty"`
}

type BotProfileLogExpressionBinding struct {
	BotBindComment          string `json:"bot_bind_comment,omitempty"`
	BotLogExpressionEnabled string `json:"bot_log_expression_enabled,omitempty"`
	BotLogExpressionName    string `json:"bot_log_expression_name,omitempty"`
	BotLogExpressionValue   string `json:"bot_log_expression_value,omitempty"`
	LogExpression           bool   `json:"logexpression,omitempty"`
	LogMessage              string `json:"logmessage,omitempty"`
	Name                    string `json:"name,omitempty"`
}

type BotPolicyLabel struct {
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Hits               int     `json:"hits,omitempty"`
	LabelName          string  `json:"labelname,omitempty"`
	NewName            string  `json:"newname,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NumPol             int     `json:"numpol,omitempty"`
}

type BotProfileCaptchaBinding struct {
	BotBindComment    string   `json:"bot_bind_comment,omitempty"`
	BotCaptchaAction  []string `json:"bot_captcha_action,omitempty"`
	BotCaptchaEnabled string   `json:"bot_captcha_enabled,omitempty"`
	BotCaptchaURL     string   `json:"bot_captcha_url,omitempty"`
	CaptchaResource   bool     `json:"captcharesource,omitempty"`
	GracePeriod       int      `json:"graceperiod,omitempty"`
	LogMessage        string   `json:"logmessage,omitempty"`
	MutePeriod        int      `json:"muteperiod,omitempty"`
	Name              string   `json:"name,omitempty"`
	RequestSizeLimit  int      `json:"requestsizelimit,omitempty"`
	RetryAttempts     int      `json:"retryattempts,omitempty"`
	WaitTime          int      `json:"waittime,omitempty"`
}

type BotPolicyBotGlobalBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type BotPolicyBotPolicyLabelBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type BotPolicyLabelPolicyBindingBinding struct {
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelName        string `json:"invoke_labelname,omitempty"`
	LabelName              string `json:"labelname,omitempty"`
	LabelType              string `json:"labeltype,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type BotProfileTPSBinding struct {
	BotBindComment string   `json:"bot_bind_comment,omitempty"`
	BotTPS         bool     `json:"bot_tps,omitempty"`
	BotTPSAction   []string `json:"bot_tps_action,omitempty"`
	BotTPSEnabled  string   `json:"bot_tps_enabled,omitempty"`
	BotTPSType     string   `json:"bot_tps_type,omitempty"`
	LogMessage     string   `json:"logmessage,omitempty"`
	Name           string   `json:"name,omitempty"`
	Percentage     int      `json:"percentage,omitempty"`
	Threshold      int      `json:"threshold,omitempty"`
}

type BotProfileRateLimitBinding struct {
	BotBindComment      string   `json:"bot_bind_comment,omitempty"`
	BotRateLimitAction  []string `json:"bot_rate_limit_action,omitempty"`
	BotRateLimitEnabled string   `json:"bot_rate_limit_enabled,omitempty"`
	BotRateLimitType    string   `json:"bot_rate_limit_type,omitempty"`
	BotRateLimitURL     string   `json:"bot_rate_limit_url,omitempty"`
	BotRateLimit        bool     `json:"bot_ratelimit,omitempty"`
	Condition           string   `json:"condition,omitempty"`
	CookieName          string   `json:"cookiename,omitempty"`
	CountryCode         string   `json:"countrycode,omitempty"`
	LimitType           string   `json:"limittype,omitempty"`
	LogMessage          string   `json:"logmessage,omitempty"`
	Name                string   `json:"name,omitempty"`
	Rate                int      `json:"rate,omitempty"`
	TimeSlice           int      `json:"timeslice,omitempty"`
}

type BotProfileBlacklistBinding struct {
	BotBindComment      string   `json:"bot_bind_comment,omitempty"`
	BotBlacklist        bool     `json:"bot_blacklist,omitempty"`
	BotBlacklistAction  []string `json:"bot_blacklist_action,omitempty"`
	BotBlacklistEnabled string   `json:"bot_blacklist_enabled,omitempty"`
	BotBlacklistType    string   `json:"bot_blacklist_type,omitempty"`
	BotBlacklistValue   string   `json:"bot_blacklist_value,omitempty"`
	LogMessage          string   `json:"logmessage,omitempty"`
	Name                string   `json:"name,omitempty"`
}
