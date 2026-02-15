package models

// ica configuration structs
type IcapolicyCrvserverBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type IcaglobalIcapolicyBinding struct {
	Flowtype               int    `json:"flowtype,omitempty"`
	Globalbindtype         string `json:"globalbindtype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Numpol                 int    `json:"numpol,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	TypeField              string `json:"type,omitempty"`
}

type Icalatencyprofile struct {
	Builtin                  []string `json:"builtin,omitempty"`
	Count                    float64  `json:"__count,omitempty"`
	Feature                  string   `json:"feature,omitempty"`
	Isdefault                bool     `json:"isdefault,omitempty"`
	L7latencymaxnotifycount  int      `json:"l7latencymaxnotifycount,omitempty"`
	L7latencymonitoring      string   `json:"l7latencymonitoring,omitempty"`
	L7latencynotifyinterval  int      `json:"l7latencynotifyinterval,omitempty"`
	L7latencythresholdfactor int      `json:"l7latencythresholdfactor,omitempty"`
	L7latencywaittime        int      `json:"l7latencywaittime,omitempty"`
	Name                     string   `json:"name,omitempty"`
	Nextgenapiresource       string   `json:"_nextgenapiresource,omitempty"`
	Refcnt                   int      `json:"refcnt,omitempty"`
}

type Icaaction struct {
	Accessprofilename  string   `json:"accessprofilename,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	Isdefault          bool     `json:"isdefault,omitempty"`
	Latencyprofilename string   `json:"latencyprofilename,omitempty"`
	Name               string   `json:"name,omitempty"`
	Newname            string   `json:"newname,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Referencecount     int      `json:"referencecount,omitempty"`
	Undefhits          int      `json:"undefhits,omitempty"`
}

type IcapolicyVpnvserverBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type IcaglobalBinding struct {
	IcaglobalIcapolicyBinding []interface{} `json:"icaglobal_icapolicy_binding,omitempty"`
}

type Icaparameter struct {
	Builtin               []string `json:"builtin,omitempty"`
	Dfpersistence         string   `json:"dfpersistence,omitempty"`
	Edtlosstolerant       string   `json:"edtlosstolerant,omitempty"`
	Edtpmtuddf            string   `json:"edtpmtuddf,omitempty"`
	Edtpmtuddftimeout     int      `json:"edtpmtuddftimeout,omitempty"`
	Edtpmtudrediscovery   string   `json:"edtpmtudrediscovery,omitempty"`
	Enablesronhafailover  string   `json:"enablesronhafailover,omitempty"`
	Hdxinsightnonnsap     string   `json:"hdxinsightnonnsap,omitempty"`
	Insightonlytodirector string   `json:"insightonlytodirector,omitempty"`
	L7latencyfrequency    int      `json:"l7latencyfrequency,omitempty"`
	Nextgenapiresource    string   `json:"_nextgenapiresource,omitempty"`
}

type IcapolicyIcaglobalBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Icapolicy struct {
	Action             string   `json:"action,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	Comment            string   `json:"comment,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	Logaction          string   `json:"logaction,omitempty"`
	Name               string   `json:"name,omitempty"`
	Newname            string   `json:"newname,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Rule               string   `json:"rule,omitempty"`
	Undefhits          int      `json:"undefhits,omitempty"`
}

type IcapolicyBinding struct {
	IcapolicyCrvserverBinding  []interface{} `json:"icapolicy_crvserver_binding,omitempty"`
	IcapolicyIcaglobalBinding  []interface{} `json:"icapolicy_icaglobal_binding,omitempty"`
	IcapolicyVpnvserverBinding []interface{} `json:"icapolicy_vpnvserver_binding,omitempty"`
	Name                       string        `json:"name,omitempty"`
}

type Icaaccessprofile struct {
	Builtin                      []string `json:"builtin,omitempty"`
	Clientaudioredirection       string   `json:"clientaudioredirection,omitempty"`
	Clientclipboardredirection   string   `json:"clientclipboardredirection,omitempty"`
	Clientcomportredirection     string   `json:"clientcomportredirection,omitempty"`
	Clientdriveredirection       string   `json:"clientdriveredirection,omitempty"`
	Clientprinterredirection     string   `json:"clientprinterredirection,omitempty"`
	Clienttwaindeviceredirection string   `json:"clienttwaindeviceredirection,omitempty"`
	Clientusbdriveredirection    string   `json:"clientusbdriveredirection,omitempty"`
	Connectclientlptports        string   `json:"connectclientlptports,omitempty"`
	Count                        float64  `json:"__count,omitempty"`
	Draganddrop                  string   `json:"draganddrop,omitempty"`
	Feature                      string   `json:"feature,omitempty"`
	Fido2redirection             string   `json:"fido2redirection,omitempty"`
	Isdefault                    bool     `json:"isdefault,omitempty"`
	Localremotedatasharing       string   `json:"localremotedatasharing,omitempty"`
	Multistream                  string   `json:"multistream,omitempty"`
	Name                         string   `json:"name,omitempty"`
	Nextgenapiresource           string   `json:"_nextgenapiresource,omitempty"`
	Refcnt                       int      `json:"refcnt,omitempty"`
	Smartcardredirection         string   `json:"smartcardredirection,omitempty"`
	Wiaredirection               string   `json:"wiaredirection,omitempty"`
}
