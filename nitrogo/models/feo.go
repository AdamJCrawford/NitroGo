package models

// feo configuration structs
type FEOPolicyCSVServerBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type FEOGlobalBinding struct {
	FEOGlobalFEOPolicyBinding []interface{} `json:"feoglobal_feopolicy_binding,omitempty"`
}

type FEOPolicyFEOGlobalBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type FEOPolicyLBVServerBinding struct {
	ActivePolicy           int    `json:"activepolicy,omitempty"`
	BoundTo                string `json:"boundto,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type FEOPolicy struct {
	Action             string   `json:"action,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	Name               string   `json:"name,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	Rule               string   `json:"rule,omitempty"`
	UndefHits          int      `json:"undefhits,omitempty"`
}

type FEOAction struct {
	Builtin                []string `json:"builtin,omitempty"`
	CacheMaxAge            int      `json:"cachemaxage,omitempty"`
	ClientSideMeasurements bool     `json:"clientsidemeasurements,omitempty"`
	ConvertImportToLink    bool     `json:"convertimporttolink,omitempty"`
	Count                  float64  `json:"__count,omitempty"`
	CSSCombine             bool     `json:"csscombine,omitempty"`
	CSSFlattenImports      bool     `json:"cssflattenimports,omitempty"`
	CSSImgInline           bool     `json:"cssimginline,omitempty"`
	CSSInline              bool     `json:"cssinline,omitempty"`
	CSSMinify              bool     `json:"cssminify,omitempty"`
	CSSMoveToHead          bool     `json:"cssmovetohead,omitempty"`
	DNSShards              []string `json:"dnsshards,omitempty"`
	DomainSharding         string   `json:"domainsharding,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	Hits                   int      `json:"hits,omitempty"`
	HTMLMinify             bool     `json:"htmlminify,omitempty"`
	HTMLRmAttribQuotes     bool     `json:"htmlrmattribquotes,omitempty"`
	HTMLRmDefaultAttribs   bool     `json:"htmlrmdefaultattribs,omitempty"`
	HTMLTrimURLs           bool     `json:"htmltrimurls,omitempty"`
	ImgAddDimensions       bool     `json:"imgadddimensions,omitempty"`
	ImgGifToPNG            bool     `json:"imggiftopng,omitempty"`
	ImgInline              bool     `json:"imginline,omitempty"`
	ImgLazyLoad            bool     `json:"imglazyload,omitempty"`
	ImgShrinkForMobile     bool     `json:"imgshrinkformobile,omitempty"`
	ImgShrinkToAttrib      bool     `json:"imgshrinktoattrib,omitempty"`
	ImgToJPEGXR            bool     `json:"imgtojpegxr,omitempty"`
	ImgToWebP              bool     `json:"imgtowebp,omitempty"`
	ImgWeaken              bool     `json:"imgweaken,omitempty"`
	JPGOptimize            bool     `json:"jpgoptimize,omitempty"`
	JPGProgressive         bool     `json:"jpgprogressive,omitempty"`
	JSCombine              bool     `json:"jscombine,omitempty"`
	JSInline               bool     `json:"jsinline,omitempty"`
	JSMinify               bool     `json:"jsminify,omitempty"`
	JSMoveToEnd            bool     `json:"jsmovetoend,omitempty"`
	Name                   string   `json:"name,omitempty"`
	NextGenAPIResource     string   `json:"_nextgenapiresource,omitempty"`
	PageExtendCache        bool     `json:"pageextendcache,omitempty"`
	UndefHits              int      `json:"undefhits,omitempty"`
}

type FEOPolicyBinding struct {
	FEOPolicyCSVServerBinding []interface{} `json:"feopolicy_csvserver_binding,omitempty"`
	FEOPolicyFEOGlobalBinding []interface{} `json:"feopolicy_feoglobal_binding,omitempty"`
	FEOPolicyLBVServerBinding []interface{} `json:"feopolicy_lbvserver_binding,omitempty"`
	Name                      string        `json:"name,omitempty"`
}

type FEOParameter struct {
	Builtin            []string `json:"builtin,omitempty"`
	CSSInlineThresSize int      `json:"cssinlinethressize,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	ImgInlineThresSize int      `json:"imginlinethressize,omitempty"`
	JPEGQualityPercent int      `json:"jpegqualitypercent,omitempty"`
	JSInlineThresSize  int      `json:"jsinlinethressize,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
}

type FEOGlobalFEOPolicyBinding struct {
	GlobalBindType         string `json:"globalbindtype,omitempty"`
	GotoPriorityExpression string `json:"gotopriorityexpression,omitempty"`
	NumPol                 int    `json:"numpol,omitempty"`
	PolicyName             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
}
