package models

// dns configuration structs
type Dnsnsecrec struct {
	Count              float64  `json:"__count,omitempty"`
	Ecssubnet          string   `json:"ecssubnet,omitempty"`
	Hostname           string   `json:"hostname,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Nextnsec           string   `json:"nextnsec,omitempty"`
	Nextrecs           []string `json:"nextrecs,omitempty"`
	Ttl                int      `json:"ttl,omitempty"`
	TypeField          string   `json:"type,omitempty"`
}

type Dnstxtrec struct {
	Authtype           string   `json:"authtype,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Domain             string   `json:"domain,omitempty"`
	Ecssubnet          string   `json:"ecssubnet,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Nodeid             int      `json:"nodeid,omitempty"`
	Recordid           int      `json:"recordid,omitempty"`
	String             []string `json:"String,omitempty"`
	Ttl                int      `json:"ttl,omitempty"`
	TypeField          string   `json:"type,omitempty"`
}

type Dnszone struct {
	Count              float64  `json:"__count,omitempty"`
	Dnssecoffload      string   `json:"dnssecoffload,omitempty"`
	Flags              int      `json:"flags,omitempty"`
	Keyname            []string `json:"keyname,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Nsec               string   `json:"nsec,omitempty"`
	Proxymode          string   `json:"proxymode,omitempty"`
	TypeField          string   `json:"type,omitempty"`
	Zonename           string   `json:"zonename,omitempty"`
}

type DnspolicylabelBinding struct {
	DnspolicylabelDnspolicyBinding     []interface{} `json:"dnspolicylabel_dnspolicy_binding,omitempty"`
	DnspolicylabelPolicybindingBinding []interface{} `json:"dnspolicylabel_policybinding_binding,omitempty"`
	Labelname                          string        `json:"labelname,omitempty"`
}

type Dnsdsfile struct {
	Keyname            string `json:"keyname,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Dnsnegativecacherecords struct {
	Count              float64 `json:"__count,omitempty"`
	Ecssubnet          string  `json:"ecssubnet,omitempty"`
	Hostname           string  `json:"hostname,omitempty"`
	Negcachetype       string  `json:"negcachetype,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Nodeid             int     `json:"nodeid,omitempty"`
	Peid               int     `json:"peid,omitempty"`
	Querytype          string  `json:"querytype,omitempty"`
	Rdclient           string  `json:"rdclient,omitempty"`
	Ttl                int     `json:"ttl,omitempty"`
}

type DnspolicylabelDnspolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Dnsptrrec struct {
	Authtype           string  `json:"authtype,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Domain             string  `json:"domain,omitempty"`
	Ecssubnet          string  `json:"ecssubnet,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Nodeid             int     `json:"nodeid,omitempty"`
	Reversedomain      string  `json:"reversedomain,omitempty"`
	Ttl                int     `json:"ttl,omitempty"`
	TypeField          string  `json:"type,omitempty"`
}

type DnszoneDomainBinding struct {
	Domain   string   `json:"domain,omitempty"`
	Nextrecs []string `json:"nextrecs,omitempty"`
	Zonename string   `json:"zonename,omitempty"`
}

type DnszoneBinding struct {
	DnszoneDnskeyBinding     []interface{} `json:"dnszone_dnskey_binding,omitempty"`
	DnszoneGslbdomainBinding []interface{} `json:"dnszone_gslbdomain_binding,omitempty"`
	Zonename                 string        `json:"zonename,omitempty"`
}

type DnsglobalDnspolicyBinding struct {
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

type Dnsnameserver struct {
	Clmonowner         int     `json:"clmonowner,omitempty"`
	Clmonview          int     `json:"clmonview,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Dnsprofilename     string  `json:"dnsprofilename,omitempty"`
	Dnsvservername     string  `json:"dnsvservername,omitempty"`
	Ip                 string  `json:"ip,omitempty"`
	Local              bool    `json:"local,omitempty"`
	Nameserverstate    string  `json:"nameserverstate,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Port               int     `json:"port,omitempty"`
	Servicename        string  `json:"servicename,omitempty"`
	State              string  `json:"state,omitempty"`
	TypeField          string  `json:"type,omitempty"`
}

type DnspolicyDnspolicylabelBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Dnssrvrec struct {
	Authtype           string  `json:"authtype,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Domain             string  `json:"domain,omitempty"`
	Ecssubnet          string  `json:"ecssubnet,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Nodeid             int     `json:"nodeid,omitempty"`
	Port               int     `json:"port,omitempty"`
	Priority           int     `json:"priority,omitempty"`
	Target             string  `json:"target,omitempty"`
	Ttl                int     `json:"ttl,omitempty"`
	TypeField          string  `json:"type,omitempty"`
	Weight             int     `json:"weight,omitempty"`
}

type DnsviewBinding struct {
	DnsviewDnspolicyBinding   []interface{} `json:"dnsview_dnspolicy_binding,omitempty"`
	DnsviewGslbserviceBinding []interface{} `json:"dnsview_gslbservice_binding,omitempty"`
	Viewname                  string        `json:"viewname,omitempty"`
}

type Dnspolicylabel struct {
	Count                  float64 `json:"__count,omitempty"`
	Description            string  `json:"description,omitempty"`
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
	Transform              string  `json:"transform,omitempty"`
}

type Dnsproxyrecords struct {
	Negrectype string `json:"negrectype,omitempty"`
	TypeField  string `json:"type,omitempty"`
}

type Dnsnsrec struct {
	Authtype           string  `json:"authtype,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Domain             string  `json:"domain,omitempty"`
	Ecssubnet          string  `json:"ecssubnet,omitempty"`
	Nameserver         string  `json:"nameserver,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Nodeid             int     `json:"nodeid,omitempty"`
	Ttl                int     `json:"ttl,omitempty"`
	TypeField          string  `json:"type,omitempty"`
}

type DnspolicyDnsglobalBinding struct {
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Dnspolicy struct {
	Actionname         string   `json:"actionname,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	Cachebypass        string   `json:"cachebypass,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Description        string   `json:"description,omitempty"`
	Drop               string   `json:"drop,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	Logaction          string   `json:"logaction,omitempty"`
	Name               string   `json:"name,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Preferredlocation  string   `json:"preferredlocation,omitempty"`
	Preferredloclist   []string `json:"preferredloclist,omitempty"`
	Rule               string   `json:"rule,omitempty"`
	Undefhits          int      `json:"undefhits,omitempty"`
	Viewname           string   `json:"viewname,omitempty"`
}

type Dnsaction struct {
	Actionname         string   `json:"actionname,omitempty"`
	Actiontype         string   `json:"actiontype,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	Cachebypass        string   `json:"cachebypass,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Dnsprofilename     string   `json:"dnsprofilename,omitempty"`
	Drop               string   `json:"drop,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Ipaddress          []string `json:"ipaddress,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Preferredloclist   []string `json:"preferredloclist,omitempty"`
	Ttl                int      `json:"ttl,omitempty"`
	Viewname           string   `json:"viewname,omitempty"`
}

type DnsglobalBinding struct {
	DnsglobalDnspolicyBinding []interface{} `json:"dnsglobal_dnspolicy_binding,omitempty"`
}

type Dnssoarec struct {
	Authtype           string  `json:"authtype,omitempty"`
	Contact            string  `json:"contact,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Domain             string  `json:"domain,omitempty"`
	Ecssubnet          string  `json:"ecssubnet,omitempty"`
	Expire             int     `json:"expire,omitempty"`
	Minimum            int     `json:"minimum,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Nodeid             int     `json:"nodeid,omitempty"`
	Originserver       string  `json:"originserver,omitempty"`
	Refresh            int     `json:"refresh,omitempty"`
	Retry              int     `json:"retry,omitempty"`
	Serial             int     `json:"serial,omitempty"`
	Ttl                int     `json:"ttl,omitempty"`
	TypeField          string  `json:"type,omitempty"`
}

type Dnspolicy64Binding struct {
	Dnspolicy64LbvserverBinding []interface{} `json:"dnspolicy64_lbvserver_binding,omitempty"`
	Name                        string        `json:"name,omitempty"`
}

type DnspolicyBinding struct {
	DnspolicyDnsglobalBinding      []interface{} `json:"dnspolicy_dnsglobal_binding,omitempty"`
	DnspolicyDnspolicylabelBinding []interface{} `json:"dnspolicy_dnspolicylabel_binding,omitempty"`
	Name                           string        `json:"name,omitempty"`
}

type DnszoneDnskeyBinding struct {
	Expires          int           `json:"expires,omitempty"`
	Keyname          []string      `json:"keyname,omitempty"`
	Siginceptiontime []interface{} `json:"siginceptiontime,omitempty"`
	Signed           int           `json:"signed,omitempty"`
	Zonename         string        `json:"zonename,omitempty"`
}

type Dnspolicy64LbvserverBinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}

type Dnssuffix struct {
	Count              float64 `json:"__count,omitempty"`
	Dnssuffix          string  `json:"Dnssuffix,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
}

type Dnsparameter struct {
	Autosavekeyops               string   `json:"autosavekeyops,omitempty"`
	Builtin                      []string `json:"builtin,omitempty"`
	Cacheecszeroprefix           string   `json:"cacheecszeroprefix,omitempty"`
	Cachehitbypass               string   `json:"cachehitbypass,omitempty"`
	Cachenoexpire                string   `json:"cachenoexpire,omitempty"`
	Cacherecords                 string   `json:"cacherecords,omitempty"`
	Dns64timeout                 int      `json:"dns64timeout,omitempty"`
	Dnsrootreferral              string   `json:"dnsrootreferral,omitempty"`
	Dnssec                       string   `json:"dnssec,omitempty"`
	Ecsmaxsubnets                int      `json:"ecsmaxsubnets,omitempty"`
	Feature                      string   `json:"feature,omitempty"`
	Maxcachesize                 int      `json:"maxcachesize,omitempty"`
	Maxnegativecachesize         int      `json:"maxnegativecachesize,omitempty"`
	Maxnegcachettl               int      `json:"maxnegcachettl,omitempty"`
	Maxpipeline                  int      `json:"maxpipeline,omitempty"`
	Maxttl                       int      `json:"maxttl,omitempty"`
	Maxudppacketsize             int      `json:"maxudppacketsize,omitempty"`
	Minttl                       int      `json:"minttl,omitempty"`
	Namelookuppriority           string   `json:"namelookuppriority,omitempty"`
	Nextgenapiresource           string   `json:"_nextgenapiresource,omitempty"`
	Nxdomainratelimitthreshold   int      `json:"nxdomainratelimitthreshold,omitempty"`
	Nxdomainthresholdcrossed     int      `json:"nxdomainthresholdcrossed,omitempty"`
	Recursion                    string   `json:"recursion,omitempty"`
	Resolutionorder              string   `json:"resolutionorder,omitempty"`
	Resolvermaxactiveresolutions int      `json:"resolvermaxactiveresolutions,omitempty"`
	Resolvermaxtcpconnections    int      `json:"resolvermaxtcpconnections,omitempty"`
	Resolvermaxtcptimeout        int      `json:"resolvermaxtcptimeout,omitempty"`
	Retries                      int      `json:"retries,omitempty"`
	Splitpktqueryprocessing      string   `json:"splitpktqueryprocessing,omitempty"`
	Zonetransfer                 string   `json:"zonetransfer,omitempty"`
}

type Dnsaction64 struct {
	Actionname         string   `json:"actionname,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Excluderule        string   `json:"excluderule,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Mappedrule         string   `json:"mappedrule,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Prefix             string   `json:"prefix,omitempty"`
}

type Dnscnamerec struct {
	Aliasname          string  `json:"aliasname,omitempty"`
	Authtype           string  `json:"authtype,omitempty"`
	Canonicalname      string  `json:"canonicalname,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Ecssubnet          string  `json:"ecssubnet,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Nodeid             int     `json:"nodeid,omitempty"`
	Ttl                int     `json:"ttl,omitempty"`
	TypeField          string  `json:"type,omitempty"`
	Vservername        string  `json:"vservername,omitempty"`
}

type Dnspolicy64 struct {
	Action             string  `json:"action,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Description        string  `json:"description,omitempty"`
	Hits               int     `json:"hits,omitempty"`
	Labelname          string  `json:"labelname,omitempty"`
	Labeltype          string  `json:"labeltype,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Rule               string  `json:"rule,omitempty"`
	Undefhits          int     `json:"undefhits,omitempty"`
}

type Dnsprofile struct {
	Cacheecsresponses            string  `json:"cacheecsresponses,omitempty"`
	Cachenegativeresponses       string  `json:"cachenegativeresponses,omitempty"`
	Cacherecords                 string  `json:"cacherecords,omitempty"`
	Count                        float64 `json:"__count,omitempty"`
	Dnsanswerseclogging          string  `json:"dnsanswerseclogging,omitempty"`
	Dnserrorlogging              string  `json:"dnserrorlogging,omitempty"`
	Dnsextendedlogging           string  `json:"dnsextendedlogging,omitempty"`
	Dnsprofilename               string  `json:"dnsprofilename,omitempty"`
	Dnsquerylogging              string  `json:"dnsquerylogging,omitempty"`
	Dropmultiqueryrequest        string  `json:"dropmultiqueryrequest,omitempty"`
	Insertecs                    string  `json:"insertecs,omitempty"`
	Maxcacheableecsprefixlength  int     `json:"maxcacheableecsprefixlength,omitempty"`
	Maxcacheableecsprefixlength6 int     `json:"maxcacheableecsprefixlength6,omitempty"`
	Nextgenapiresource           string  `json:"_nextgenapiresource,omitempty"`
	Recursiveresolution          string  `json:"recursiveresolution,omitempty"`
	Referencecount               int     `json:"referencecount,omitempty"`
	Replaceecs                   string  `json:"replaceecs,omitempty"`
}

type DnsviewGslbserviceBinding struct {
	Gslbservicename string `json:"gslbservicename,omitempty"`
	Ipaddress       string `json:"ipaddress,omitempty"`
	Viewname        string `json:"viewname,omitempty"`
}

type Dnsaaaarec struct {
	Authtype           string  `json:"authtype,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Ecssubnet          string  `json:"ecssubnet,omitempty"`
	Hostname           string  `json:"hostname,omitempty"`
	Ipv6address        string  `json:"ipv6address,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Nodeid             int     `json:"nodeid,omitempty"`
	Ttl                int     `json:"ttl,omitempty"`
	TypeField          string  `json:"type,omitempty"`
	Vservername        string  `json:"vservername,omitempty"`
}

type Dnsview struct {
	Count              float64 `json:"__count,omitempty"`
	Flags              int     `json:"flags,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Viewname           string  `json:"viewname,omitempty"`
}

type Dnskey struct {
	Activationtimestr  string  `json:"activationtimestr,omitempty"`
	Algorithm          string  `json:"algorithm,omitempty"`
	Autorollover       string  `json:"autorollover,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Createtimestr      string  `json:"createtimestr,omitempty"`
	Deletiontimestr    string  `json:"deletiontimestr,omitempty"`
	Expires            int     `json:"expires,omitempty"`
	Expirytimestr      string  `json:"expirytimestr,omitempty"`
	Filenameprefix     string  `json:"filenameprefix,omitempty"`
	Keyname            string  `json:"keyname,omitempty"`
	Keysize            int     `json:"keysize,omitempty"`
	Keytype            string  `json:"keytype,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Notificationperiod int     `json:"notificationperiod,omitempty"`
	Password           string  `json:"password,omitempty"`
	Privatekey         string  `json:"privatekey,omitempty"`
	Publickey          string  `json:"publickey,omitempty"`
	Revoke             bool    `json:"revoke,omitempty"`
	Rolloverfailrc     int     `json:"rolloverfailrc,omitempty"`
	Rollovermethod     string  `json:"rollovermethod,omitempty"`
	Src                string  `json:"src,omitempty"`
	State              string  `json:"state,omitempty"`
	Tag                int     `json:"tag,omitempty"`
	Ttl                int     `json:"ttl,omitempty"`
	TypeField          string  `json:"type,omitempty"`
	Units1             string  `json:"units1,omitempty"`
	Units2             string  `json:"units2,omitempty"`
	Zonename           string  `json:"zonename,omitempty"`
}

type Dnsnaptrrec struct {
	Authtype           string  `json:"authtype,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Domain             string  `json:"domain,omitempty"`
	Ecssubnet          string  `json:"ecssubnet,omitempty"`
	Flags              string  `json:"flags,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Nodeid             int     `json:"nodeid,omitempty"`
	Order              int     `json:"order,omitempty"`
	Preference         int     `json:"preference,omitempty"`
	Recordid           int     `json:"recordid,omitempty"`
	Regexp             string  `json:"regexp,omitempty"`
	Replacement        string  `json:"replacement,omitempty"`
	Services           string  `json:"services,omitempty"`
	Ttl                int     `json:"ttl,omitempty"`
	TypeField          string  `json:"type,omitempty"`
	Vservername        string  `json:"vservername,omitempty"`
}

type DnsviewDnspolicyBinding struct {
	Dnspolicyname string `json:"dnspolicyname,omitempty"`
	Viewname      string `json:"viewname,omitempty"`
}

type Dnssubnetcache struct {
	All                bool     `json:"all,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Ecssubnet          string   `json:"ecssubnet,omitempty"`
	Hostname           string   `json:"hostname,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Nextrecs           []string `json:"nextrecs,omitempty"`
	Nodeid             int      `json:"nodeid,omitempty"`
}

type Dnsaddrec struct {
	Authtype           string  `json:"authtype,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Ecssubnet          string  `json:"ecssubnet,omitempty"`
	Hostname           string  `json:"hostname,omitempty"`
	Ipaddress          string  `json:"ipaddress,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Nodeid             int     `json:"nodeid,omitempty"`
	Ttl                int     `json:"ttl,omitempty"`
	TypeField          string  `json:"type,omitempty"`
	Vservername        string  `json:"vservername,omitempty"`
}

type Dnsmxrec struct {
	Authtype           string  `json:"authtype,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Domain             string  `json:"domain,omitempty"`
	Ecssubnet          string  `json:"ecssubnet,omitempty"`
	Mx                 string  `json:"mx,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Nodeid             int     `json:"nodeid,omitempty"`
	Pref               int     `json:"pref,omitempty"`
	Ttl                int     `json:"ttl,omitempty"`
	TypeField          string  `json:"type,omitempty"`
}

type Dnscaarec struct {
	Authtype           string  `json:"authtype,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Domain             string  `json:"domain,omitempty"`
	Ecssubnet          string  `json:"ecssubnet,omitempty"`
	Flag               string  `json:"flag,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Nodeid             int     `json:"nodeid,omitempty"`
	Recordid           int     `json:"recordid,omitempty"`
	Tag                string  `json:"tag,omitempty"`
	Ttl                int     `json:"ttl,omitempty"`
	TypeField          string  `json:"type,omitempty"`
	Valuestring        string  `json:"valuestring,omitempty"`
}

type DnspolicylabelPolicybindingBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	InvokeLabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
}
