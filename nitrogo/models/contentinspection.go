package models

// contentinspection configuration structs
type ContentinspectionpolicyContentinspectionglobalBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Contentinspectionpolicylabel struct {
	Comment                string  `json:"comment,omitempty"`
	Count                  float64 `json:"__count,omitempty"`
	Flowtype               int     `json:"flowtype,omitempty"`
	Gotopriorityexpression string  `json:"gotopriorityexpression,omitempty"`
	Hits                   int     `json:"hits,omitempty"`
	InvokeLabelname        string  `json:"invoke_labelname,omitempty"`
	Isdefault              bool    `json:"isdefault,omitempty"`
	Labelname              string  `json:"labelname,omitempty"`
	Labeltype              string  `json:"labeltype,omitempty"`
	Newname                string  `json:"newname,omitempty"`
	Nextgenapiresource     string  `json:"_nextgenapiresource,omitempty"`
	Numpol                 int     `json:"numpol,omitempty"`
	Priority               int     `json:"priority,omitempty"`
	TypeField              string  `json:"type,omitempty"`
}

type Contentinspectionwasmprofile struct {
	Anomalousdatasize  int     `json:"anomalousdatasize,omitempty"`
	Anomalousttfbtime  int     `json:"anomalousttfbtime,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Maxbodylen         int     `json:"maxbodylen,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Timeout            int     `json:"timeout,omitempty"`
	Timeoutaction      string  `json:"timeoutaction,omitempty"`
	Wasmmodule         string  `json:"wasmmodule,omitempty"`
}

type ContentinspectionpolicylabelPolicybindingBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type ContentinspectionpolicyCsvserverBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Contentinspectioncallout struct {
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Hits               int     `json:"hits,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Profilename        string  `json:"profilename,omitempty"`
	Resultexpr         string  `json:"resultexpr,omitempty"`
	Returntype         string  `json:"returntype,omitempty"`
	Serverip           string  `json:"serverip,omitempty"`
	Servername         string  `json:"servername,omitempty"`
	Serverport         int     `json:"serverport,omitempty"`
	TypeField          string  `json:"type,omitempty"`
	Undefhits          int     `json:"undefhits,omitempty"`
	Undefreason        string  `json:"undefreason,omitempty"`
}

type Contentinspectionprofile struct {
	Count              float64 `json:"__count,omitempty"`
	Egressinterface    string  `json:"egressinterface,omitempty"`
	Egressvlan         int     `json:"egressvlan,omitempty"`
	Ingressinterface   string  `json:"ingressinterface,omitempty"`
	Ingressvlan        int     `json:"ingressvlan,omitempty"`
	Iptunnel           string  `json:"iptunnel,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	TypeField          string  `json:"type,omitempty"`
}

type ContentinspectionglobalContentinspectionpolicyBinding struct {
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

type Contentinspectionparameter struct {
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Undefaction        string `json:"undefaction,omitempty"`
}

type ContentinspectionpolicylabelBinding struct {
	ContentinspectionpolicylabelContentinspectionpolicyBinding []interface{} `json:"contentinspectionpolicylabel_contentinspectionpolicy_binding,omitempty"`
	ContentinspectionpolicylabelPolicybindingBinding           []interface{} `json:"contentinspectionpolicylabel_policybinding_binding,omitempty"`
	Labelname                                                  string        `json:"labelname,omitempty"`
}

type Contentinspectionaction struct {
	Builtin            []string `json:"builtin,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	Icapprofilename    string   `json:"icapprofilename,omitempty"`
	Ifserverdown       string   `json:"ifserverdown,omitempty"`
	Name               string   `json:"name,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Referencecount     int      `json:"referencecount,omitempty"`
	Reqtimeout         int      `json:"reqtimeout,omitempty"`
	Reqtimeoutaction   string   `json:"reqtimeoutaction,omitempty"`
	Serverip           string   `json:"serverip,omitempty"`
	Servername         string   `json:"servername,omitempty"`
	Serverport         int      `json:"serverport,omitempty"`
	TypeField          string   `json:"type,omitempty"`
	Undefhits          int      `json:"undefhits,omitempty"`
	Wasmprofilename    string   `json:"wasmprofilename,omitempty"`
}

type ContentinspectionpolicyLbvserverBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type ContentinspectionglobalBinding struct {
	ContentinspectionglobalContentinspectionpolicyBinding []interface{} `json:"contentinspectionglobal_contentinspectionpolicy_binding,omitempty"`
}

type ContentinspectionpolicyContentinspectionpolicylabelBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type ContentinspectionpolicyBinding struct {
	ContentinspectionpolicyContentinspectionglobalBinding      []interface{} `json:"contentinspectionpolicy_contentinspectionglobal_binding,omitempty"`
	ContentinspectionpolicyContentinspectionpolicylabelBinding []interface{} `json:"contentinspectionpolicy_contentinspectionpolicylabel_binding,omitempty"`
	ContentinspectionpolicyCsvserverBinding                    []interface{} `json:"contentinspectionpolicy_csvserver_binding,omitempty"`
	ContentinspectionpolicyLbvserverBinding                    []interface{} `json:"contentinspectionpolicy_lbvserver_binding,omitempty"`
	Name                                                       string        `json:"name,omitempty"`
}

type ContentinspectionpolicylabelContentinspectionpolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Contentinspectionpolicy struct {
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
	Undefaction        string   `json:"undefaction,omitempty"`
	Undefhits          int      `json:"undefhits,omitempty"`
}
