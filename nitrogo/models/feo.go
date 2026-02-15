package models

// feo configuration structs
type FeopolicyCsvserverBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type FeoglobalBinding struct {
	FeoglobalFeopolicyBinding []interface{} `json:"feoglobal_feopolicy_binding,omitempty"`
}

type FeopolicyFeoglobalBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type FeopolicyLbvserverBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Feopolicy struct {
	Action             string   `json:"action,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	Name               string   `json:"name,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Rule               string   `json:"rule,omitempty"`
	Undefhits          int      `json:"undefhits,omitempty"`
}

type Feoaction struct {
	Builtin                []string `json:"builtin,omitempty"`
	Cachemaxage            int      `json:"cachemaxage,omitempty"`
	Clientsidemeasurements bool     `json:"clientsidemeasurements,omitempty"`
	Convertimporttolink    bool     `json:"convertimporttolink,omitempty"`
	Count                  float64  `json:"__count,omitempty"`
	Csscombine             bool     `json:"csscombine,omitempty"`
	Cssflattenimports      bool     `json:"cssflattenimports,omitempty"`
	Cssimginline           bool     `json:"cssimginline,omitempty"`
	Cssinline              bool     `json:"cssinline,omitempty"`
	Cssminify              bool     `json:"cssminify,omitempty"`
	Cssmovetohead          bool     `json:"cssmovetohead,omitempty"`
	Dnsshards              []string `json:"dnsshards,omitempty"`
	Domainsharding         string   `json:"domainsharding,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	Hits                   int      `json:"hits,omitempty"`
	Htmlminify             bool     `json:"htmlminify,omitempty"`
	Htmlrmattribquotes     bool     `json:"htmlrmattribquotes,omitempty"`
	Htmlrmdefaultattribs   bool     `json:"htmlrmdefaultattribs,omitempty"`
	Htmltrimurls           bool     `json:"htmltrimurls,omitempty"`
	Imgadddimensions       bool     `json:"imgadddimensions,omitempty"`
	Imggiftopng            bool     `json:"imggiftopng,omitempty"`
	Imginline              bool     `json:"imginline,omitempty"`
	Imglazyload            bool     `json:"imglazyload,omitempty"`
	Imgshrinkformobile     bool     `json:"imgshrinkformobile,omitempty"`
	Imgshrinktoattrib      bool     `json:"imgshrinktoattrib,omitempty"`
	Imgtojpegxr            bool     `json:"imgtojpegxr,omitempty"`
	Imgtowebp              bool     `json:"imgtowebp,omitempty"`
	Imgweaken              bool     `json:"imgweaken,omitempty"`
	Jpgoptimize            bool     `json:"jpgoptimize,omitempty"`
	Jpgprogressive         bool     `json:"jpgprogressive,omitempty"`
	Jscombine              bool     `json:"jscombine,omitempty"`
	Jsinline               bool     `json:"jsinline,omitempty"`
	Jsminify               bool     `json:"jsminify,omitempty"`
	Jsmovetoend            bool     `json:"jsmovetoend,omitempty"`
	Name                   string   `json:"name,omitempty"`
	Nextgenapiresource     string   `json:"_nextgenapiresource,omitempty"`
	Pageextendcache        bool     `json:"pageextendcache,omitempty"`
	Undefhits              int      `json:"undefhits,omitempty"`
}

type FeopolicyBinding struct {
	FeopolicyCsvserverBinding []interface{} `json:"feopolicy_csvserver_binding,omitempty"`
	FeopolicyFeoglobalBinding []interface{} `json:"feopolicy_feoglobal_binding,omitempty"`
	FeopolicyLbvserverBinding []interface{} `json:"feopolicy_lbvserver_binding,omitempty"`
	Name                      string        `json:"name,omitempty"`
}

type Feoparameter struct {
	Builtin            []string `json:"builtin,omitempty"`
	Cssinlinethressize int      `json:"cssinlinethressize,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Imginlinethressize int      `json:"imginlinethressize,omitempty"`
	Jpegqualitypercent int      `json:"jpegqualitypercent,omitempty"`
	Jsinlinethressize  int      `json:"jsinlinethressize,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
}

type FeoglobalFeopolicyBinding struct {
	Globalbindtype         string `json:"globalbindtype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Numpol                 int    `json:"numpol,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
}
