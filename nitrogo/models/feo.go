// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Feoglobalpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Type                   string `json:"type,omitempty"`
	Numpol                 uint32 `json:"numpol,omitempty"`
	Globalbindtype         string `json:"globalbindtype,omitempty"`
}

type Feoparameter struct {
	Jpegqualitypercent int    `json:"jpegqualitypercent"`
	Cssinlinethressize int    `json:"cssinlinethressize,omitempty"`
	Jsinlinethressize  int    `json:"jsinlinethressize,omitempty"`
	Imginlinethressize int    `json:"imginlinethressize,omitempty"`
	Builtin            string `json:"builtin,omitempty"`
	Feature            string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Feopolicy struct {
	Name               string `json:"name,omitempty"`
	Rule               string `json:"rule,omitempty"`
	Action             string `json:"action,omitempty"`
	Builtin            string `json:"builtin,omitempty"`
	Feature            string `json:"feature,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Undefhits          string `json:"undefhits,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Feopolicyfeoglobalbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Feoglobalfeopolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Type                   string `json:"type,omitempty"`
	Numpol                 int    `json:"numpol,omitempty"`
	Globalbindtype         string `json:"globalbindtype,omitempty"`
}

type Feopolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Feopolicycsvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Feopolicyglobalbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Activepolicy           int32  `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Feopolicylbvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Feopolicyvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Activepolicy           int32  `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Feoaction struct {
	Name                   string   `json:"name,omitempty"`
	Pageextendcache        bool     `json:"pageextendcache,omitempty"`
	Cachemaxage            int      `json:"cachemaxage"`
	Imgshrinktoattrib      bool     `json:"imgshrinktoattrib,omitempty"`
	Imggiftopng            bool     `json:"imggiftopng,omitempty"`
	Imgtowebp              bool     `json:"imgtowebp,omitempty"`
	Imgtojpegxr            bool     `json:"imgtojpegxr,omitempty"`
	Imginline              bool     `json:"imginline,omitempty"`
	Cssimginline           bool     `json:"cssimginline,omitempty"`
	Jpgoptimize            bool     `json:"jpgoptimize,omitempty"`
	Imglazyload            bool     `json:"imglazyload,omitempty"`
	Cssminify              bool     `json:"cssminify,omitempty"`
	Cssinline              bool     `json:"cssinline,omitempty"`
	Csscombine             bool     `json:"csscombine,omitempty"`
	Convertimporttolink    bool     `json:"convertimporttolink,omitempty"`
	Jsminify               bool     `json:"jsminify,omitempty"`
	Jsinline               bool     `json:"jsinline,omitempty"`
	Htmlminify             bool     `json:"htmlminify,omitempty"`
	Cssmovetohead          bool     `json:"cssmovetohead,omitempty"`
	Jsmovetoend            bool     `json:"jsmovetoend,omitempty"`
	Domainsharding         string   `json:"domainsharding,omitempty"`
	Dnsshards              []string `json:"dnsshards,omitempty"`
	Clientsidemeasurements bool     `json:"clientsidemeasurements,omitempty"`
	Imgadddimensions       string   `json:"imgadddimensions,omitempty"`
	Imgshrinkformobile     string   `json:"imgshrinkformobile,omitempty"`
	Imgweaken              string   `json:"imgweaken,omitempty"`
	Jpgprogressive         string   `json:"jpgprogressive,omitempty"`
	Cssflattenimports      string   `json:"cssflattenimports,omitempty"`
	Jscombine              string   `json:"jscombine,omitempty"`
	Htmlrmdefaultattribs   string   `json:"htmlrmdefaultattribs,omitempty"`
	Htmlrmattribquotes     string   `json:"htmlrmattribquotes,omitempty"`
	Htmltrimurls           string   `json:"htmltrimurls,omitempty"`
	Hits                   string   `json:"hits,omitempty"`
	Undefhits              string   `json:"undefhits,omitempty"`
	Builtin                string   `json:"builtin,omitempty"`
	Feature                string   `json:"feature,omitempty"`
	Nextgenapiresource     string   `json:"_nextgenapiresource,omitempty"`
}

type Feoglobalbinding struct {
}
