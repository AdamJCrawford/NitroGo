package models

// cr configuration structs
type CrvserverAppqoepolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Targetvserver          string `json:"targetvserver,omitempty"`
}

type CrvserverCrpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Hits                   int    `json:"hits,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Pipolicyhits           int    `json:"pipolicyhits,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Targetvserver          string `json:"targetvserver,omitempty"`
}

type CrvserverCachepolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Targetvserver          string `json:"targetvserver,omitempty"`
}

type CrpolicyBinding struct {
	CrpolicyCrvserverBinding []interface{} `json:"crpolicy_crvserver_binding,omitempty"`
	Policyname               string        `json:"policyname,omitempty"`
}

type CrvserverCmppolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Inherited              string `json:"inherited,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Targetvserver          string `json:"targetvserver,omitempty"`
}

type CrvserverAnalyticsprofileBinding struct {
	Analyticsprofile string `json:"analyticsprofile,omitempty"`
	Name             string `json:"name,omitempty"`
}

type CrvserverLbvserverBinding struct {
	Hits      int    `json:"hits,omitempty"`
	Lbvserver string `json:"lbvserver,omitempty"`
	Name      string `json:"name,omitempty"`
}

type CrvserverAppflowpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Targetvserver          string `json:"targetvserver,omitempty"`
}

type CrvserverFeopolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Targetvserver          string `json:"targetvserver,omitempty"`
}

type CrvserverIcapolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Targetvserver          string `json:"targetvserver,omitempty"`
}

type CrvserverResponderpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Targetvserver          string `json:"targetvserver,omitempty"`
}

type Craction struct {
	Builtin            []string `json:"builtin,omitempty"`
	Comment            string   `json:"comment,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Crtype             string   `json:"crtype,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	Isdefault          bool     `json:"isdefault,omitempty"`
	Name               string   `json:"name,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Referencecount     int      `json:"referencecount,omitempty"`
	Undefhits          int      `json:"undefhits,omitempty"`
}

type CrvserverAppfwpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Targetvserver          string `json:"targetvserver,omitempty"`
}

type CrvserverPolicymapBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Targetvserver          string `json:"targetvserver,omitempty"`
}

type CrvserverRewritepolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Targetvserver          string `json:"targetvserver,omitempty"`
}

type CrvserverSpilloverpolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Targetvserver          string `json:"targetvserver,omitempty"`
}

type Crpolicy struct {
	Action             string   `json:"action,omitempty"`
	Activepolicy       bool     `json:"activepolicy,omitempty"`
	Boundto            string   `json:"boundto,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	Isdefault          bool     `json:"isdefault,omitempty"`
	Labelname          string   `json:"labelname,omitempty"`
	Labeltype          string   `json:"labeltype,omitempty"`
	Logaction          string   `json:"logaction,omitempty"`
	Newname            string   `json:"newname,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Policyname         string   `json:"policyname,omitempty"`
	Priority           int      `json:"priority,omitempty"`
	Rule               string   `json:"rule,omitempty"`
	Vstype             int      `json:"vstype,omitempty"`
}

type CrvserverBinding struct {
	CrvserverAnalyticsprofileBinding []interface{} `json:"crvserver_analyticsprofile_binding,omitempty"`
	CrvserverAppflowpolicyBinding    []interface{} `json:"crvserver_appflowpolicy_binding,omitempty"`
	CrvserverAppfwpolicyBinding      []interface{} `json:"crvserver_appfwpolicy_binding,omitempty"`
	CrvserverAppqoepolicyBinding     []interface{} `json:"crvserver_appqoepolicy_binding,omitempty"`
	CrvserverCachepolicyBinding      []interface{} `json:"crvserver_cachepolicy_binding,omitempty"`
	CrvserverCmppolicyBinding        []interface{} `json:"crvserver_cmppolicy_binding,omitempty"`
	CrvserverCrpolicyBinding         []interface{} `json:"crvserver_crpolicy_binding,omitempty"`
	CrvserverCspolicyBinding         []interface{} `json:"crvserver_cspolicy_binding,omitempty"`
	CrvserverFeopolicyBinding        []interface{} `json:"crvserver_feopolicy_binding,omitempty"`
	CrvserverIcapolicyBinding        []interface{} `json:"crvserver_icapolicy_binding,omitempty"`
	CrvserverLbvserverBinding        []interface{} `json:"crvserver_lbvserver_binding,omitempty"`
	CrvserverPolicymapBinding        []interface{} `json:"crvserver_policymap_binding,omitempty"`
	CrvserverResponderpolicyBinding  []interface{} `json:"crvserver_responderpolicy_binding,omitempty"`
	CrvserverRewritepolicyBinding    []interface{} `json:"crvserver_rewritepolicy_binding,omitempty"`
	CrvserverSpilloverpolicyBinding  []interface{} `json:"crvserver_spilloverpolicy_binding,omitempty"`
	Name                             string        `json:"name,omitempty"`
}

type Crvserver struct {
	Appflowlog               string  `json:"appflowlog,omitempty"`
	Arp                      string  `json:"arp,omitempty"`
	Authentication           string  `json:"authentication,omitempty"`
	Backendssl               string  `json:"backendssl,omitempty"`
	Backupvserver            string  `json:"backupvserver,omitempty"`
	Bindpoint                string  `json:"bindpoint,omitempty"`
	Cachetype                string  `json:"cachetype,omitempty"`
	Cachevserver             string  `json:"cachevserver,omitempty"`
	Clttimeout               int     `json:"clttimeout,omitempty"`
	Comment                  string  `json:"comment,omitempty"`
	Count                    float64 `json:"__count,omitempty"`
	Curstate                 string  `json:"curstate,omitempty"`
	Destinationvserver       string  `json:"destinationvserver,omitempty"`
	Disableprimaryondown     string  `json:"disableprimaryondown,omitempty"`
	Disallowserviceaccess    string  `json:"disallowserviceaccess,omitempty"`
	Dnsvservername           string  `json:"dnsvservername,omitempty"`
	Domain                   string  `json:"domain,omitempty"`
	Downstateflush           string  `json:"downstateflush,omitempty"`
	Format                   string  `json:"format,omitempty"`
	Ghost                    string  `json:"ghost,omitempty"`
	Gotopriorityexpression   string  `json:"gotopriorityexpression,omitempty"`
	Homepage                 string  `json:"homepage,omitempty"`
	Httpprofilename          string  `json:"httpprofilename,omitempty"`
	Icmpvsrresponse          string  `json:"icmpvsrresponse,omitempty"`
	Invoke                   bool    `json:"invoke,omitempty"`
	Ip                       string  `json:"ip,omitempty"`
	Ipset                    string  `json:"ipset,omitempty"`
	Ipv46                    string  `json:"ipv46,omitempty"`
	L2conn                   string  `json:"l2conn,omitempty"`
	Labelname                string  `json:"labelname,omitempty"`
	Labeltype                string  `json:"labeltype,omitempty"`
	Lbvserver                string  `json:"lbvserver,omitempty"`
	Listenpolicy             string  `json:"listenpolicy,omitempty"`
	Listenpriority           int     `json:"listenpriority,omitempty"`
	MapField                 string  `json:"map,omitempty"`
	Name                     string  `json:"name,omitempty"`
	Netprofile               string  `json:"netprofile,omitempty"`
	Newname                  string  `json:"newname,omitempty"`
	Nextgenapiresource       string  `json:"_nextgenapiresource,omitempty"`
	Ngname                   string  `json:"ngname,omitempty"`
	Nodefaultbindings        string  `json:"nodefaultbindings,omitempty"`
	Onpolicymatch            string  `json:"onpolicymatch,omitempty"`
	Originusip               string  `json:"originusip,omitempty"`
	Pipolicyhits             int     `json:"pipolicyhits,omitempty"`
	Policyname               string  `json:"policyname,omitempty"`
	Port                     int     `json:"port,omitempty"`
	Precedence               string  `json:"precedence,omitempty"`
	Priority                 int     `json:"priority,omitempty"`
	Probeport                int     `json:"probeport,omitempty"`
	Probeprotocol            string  `json:"probeprotocol,omitempty"`
	Probesuccessresponsecode string  `json:"probesuccessresponsecode,omitempty"`
	Range                    int     `json:"range,omitempty"`
	Redirect                 string  `json:"redirect,omitempty"`
	Redirecturl              string  `json:"redirecturl,omitempty"`
	Reuse                    string  `json:"reuse,omitempty"`
	Rhistate                 string  `json:"rhistate,omitempty"`
	Rule                     string  `json:"rule,omitempty"`
	Servicename              string  `json:"servicename,omitempty"`
	Servicetype              string  `json:"servicetype,omitempty"`
	Somethod                 string  `json:"somethod,omitempty"`
	Sopersistence            string  `json:"sopersistence,omitempty"`
	Sopersistencetimeout     int     `json:"sopersistencetimeout,omitempty"`
	Sothreshold              int     `json:"sothreshold,omitempty"`
	Srcipexpr                string  `json:"srcipexpr,omitempty"`
	State                    string  `json:"state,omitempty"`
	Status                   int     `json:"status,omitempty"`
	Targetvserver            string  `json:"targetvserver,omitempty"`
	Tcpprobeport             int     `json:"tcpprobeport,omitempty"`
	Tcpprofilename           string  `json:"tcpprofilename,omitempty"`
	Td                       int     `json:"td,omitempty"`
	TypeField                string  `json:"type,omitempty"`
	Useoriginipportforcache  string  `json:"useoriginipportforcache,omitempty"`
	Useportrange             string  `json:"useportrange,omitempty"`
	Value                    string  `json:"value,omitempty"`
	Via                      string  `json:"via,omitempty"`
	Weight                   int     `json:"weight,omitempty"`
}

type CrpolicyCrvserverBinding struct {
	Bindhits               int    `json:"bindhits,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Hits                   int    `json:"hits,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type CrvserverCspolicyBinding struct {
	Bindpoint              string `json:"bindpoint,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Hits                   int    `json:"hits,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Pipolicyhits           int    `json:"pipolicyhits,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Targetvserver          string `json:"targetvserver,omitempty"`
}
