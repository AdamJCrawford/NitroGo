package models

// cmp configuration structs
type Cmpparameter struct {
	Addvaryheader               string   `json:"addvaryheader,omitempty"`
	Builtin                     []string `json:"builtin,omitempty"`
	Cmpbypasspct                int      `json:"cmpbypasspct,omitempty"`
	Cmplevel                    string   `json:"cmplevel,omitempty"`
	Cmponpush                   string   `json:"cmponpush,omitempty"`
	Externalcache               string   `json:"externalcache,omitempty"`
	Feature                     string   `json:"feature,omitempty"`
	Heurexpiry                  string   `json:"heurexpiry,omitempty"`
	Heurexpiryhistwt            int      `json:"heurexpiryhistwt,omitempty"`
	Heurexpirythres             int      `json:"heurexpirythres,omitempty"`
	Minressize                  int      `json:"minressize,omitempty"`
	Nextgenapiresource          string   `json:"_nextgenapiresource,omitempty"`
	Policytype                  string   `json:"policytype,omitempty"`
	Quantumsize                 int      `json:"quantumsize,omitempty"`
	Randomgzipfilename          string   `json:"randomgzipfilename,omitempty"`
	Randomgzipfilenamemaxlength int      `json:"randomgzipfilenamemaxlength,omitempty"`
	Randomgzipfilenameminlength int      `json:"randomgzipfilenameminlength,omitempty"`
	Servercmp                   string   `json:"servercmp,omitempty"`
	Varyheadervalue             string   `json:"varyheadervalue,omitempty"`
}

type CmpglobalBinding struct {
	CmpglobalCmppolicyBinding []interface{} `json:"cmpglobal_cmppolicy_binding,omitempty"`
}

type CmppolicyBinding struct {
	CmppolicyCmpglobalBinding      []interface{} `json:"cmppolicy_cmpglobal_binding,omitempty"`
	CmppolicyCmppolicylabelBinding []interface{} `json:"cmppolicy_cmppolicylabel_binding,omitempty"`
	CmppolicyCrvserverBinding      []interface{} `json:"cmppolicy_crvserver_binding,omitempty"`
	CmppolicyCsvserverBinding      []interface{} `json:"cmppolicy_csvserver_binding,omitempty"`
	CmppolicyLbvserverBinding      []interface{} `json:"cmppolicy_lbvserver_binding,omitempty"`
	Name                           string        `json:"name,omitempty"`
}

type CmppolicylabelPolicybindingBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type CmppolicyCmppolicylabelBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type CmpglobalCmppolicyBinding struct {
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

type Cmppolicy struct {
	Builtin            []string `json:"builtin,omitempty"`
	Clienttransactions int      `json:"clienttransactions,omitempty"`
	Clientttlb         int      `json:"clientttlb,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Description        string   `json:"description,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	Isdefault          bool     `json:"isdefault,omitempty"`
	Name               string   `json:"name,omitempty"`
	Newname            string   `json:"newname,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Reqaction          string   `json:"reqaction,omitempty"`
	Resaction          string   `json:"resaction,omitempty"`
	Rule               string   `json:"rule,omitempty"`
	Rxbytes            int      `json:"rxbytes,omitempty"`
	Servertransactions int      `json:"servertransactions,omitempty"`
	Serverttlb         int      `json:"serverttlb,omitempty"`
	Txbytes            int      `json:"txbytes,omitempty"`
}

type Cmpaction struct {
	Addvaryheader      string   `json:"addvaryheader,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	Cmptype            string   `json:"cmptype,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Deltatype          string   `json:"deltatype,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Isdefault          bool     `json:"isdefault,omitempty"`
	Name               string   `json:"name,omitempty"`
	Newname            string   `json:"newname,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Varyheadervalue    string   `json:"varyheadervalue,omitempty"`
}

type CmppolicyCmpglobalBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Cmppolicylabel struct {
	Count                  float64 `json:"__count,omitempty"`
	Description            string  `json:"description,omitempty"`
	Flowtype               int     `json:"flowtype,omitempty"`
	Gotopriorityexpression string  `json:"gotopriorityexpression,omitempty"`
	Hits                   int     `json:"hits,omitempty"`
	InvokeLabelname        string  `json:"invoke_labelname,omitempty"`
	Labelname              string  `json:"labelname,omitempty"`
	Labeltype              string  `json:"labeltype,omitempty"`
	Newname                string  `json:"newname,omitempty"`
	Nextgenapiresource     string  `json:"_nextgenapiresource,omitempty"`
	Numpol                 int     `json:"numpol,omitempty"`
	Priority               int     `json:"priority,omitempty"`
	TypeField              string  `json:"type,omitempty"`
}

type CmppolicyCrvserverBinding struct {
	Activepolicy int    `json:"activepolicy,omitempty"`
	Boundto      string `json:"boundto,omitempty"`
	Name         string `json:"name,omitempty"`
	Priority     int    `json:"priority,omitempty"`
}

type CmppolicyLbvserverBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type CmppolicylabelCmppolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type CmppolicyCsvserverBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type CmppolicylabelBinding struct {
	CmppolicylabelCmppolicyBinding     []interface{} `json:"cmppolicylabel_cmppolicy_binding,omitempty"`
	CmppolicylabelPolicybindingBinding []interface{} `json:"cmppolicylabel_policybinding_binding,omitempty"`
	Labelname                          string        `json:"labelname,omitempty"`
}
