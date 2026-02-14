// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Dnsaddrec struct {
	Hostname           string `json:"hostname,omitempty"`
	Ipaddress          string `json:"ipaddress,omitempty"`
	Ttl                int    `json:"ttl,omitempty"`
	Ecssubnet          string `json:"ecssubnet,omitempty"`
	Type               string `json:"type,omitempty"`
	Nodeid             int    `json:"nodeid,omitempty"`
	Vservername        string `json:"vservername,omitempty"`
	Authtype           string `json:"authtype,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Dnscnamerec struct {
	Aliasname          string `json:"aliasname,omitempty"`
	Canonicalname      string `json:"canonicalname,omitempty"`
	Ttl                int    `json:"ttl,omitempty"`
	Ecssubnet          string `json:"ecssubnet,omitempty"`
	Type               string `json:"type,omitempty"`
	Nodeid             int    `json:"nodeid,omitempty"`
	Vservername        string `json:"vservername,omitempty"`
	Authtype           string `json:"authtype,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Dnsglobaldnspolicybinding struct {
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

type Dnskey struct {
	Keyname            string `json:"keyname,omitempty"`
	Publickey          string `json:"publickey,omitempty"`
	Privatekey         string `json:"privatekey,omitempty"`
	Expires            int    `json:"expires,omitempty"`
	Units1             string `json:"units1,omitempty"`
	Notificationperiod int    `json:"notificationperiod,omitempty"`
	Units2             string `json:"units2,omitempty"`
	Ttl                int    `json:"ttl,omitempty"`
	Password           string `json:"password,omitempty"`
	Autorollover       string `json:"autorollover,omitempty"`
	Rollovermethod     string `json:"rollovermethod,omitempty"`
	Revoke             bool   `json:"revoke,omitempty"`
	Zonename           string `json:"zonename,omitempty"`
	Keytype            string `json:"keytype,omitempty"`
	Algorithm          string `json:"algorithm,omitempty"`
	Keysize            int    `json:"keysize,omitempty"`
	Filenameprefix     string `json:"filenameprefix,omitempty"`
	Src                string `json:"src,omitempty"`
	State              string `json:"state,omitempty"`
	Type               string `json:"type,omitempty"`
	Tag                string `json:"tag,omitempty"`
	Createtimestr      string `json:"createtimestr,omitempty"`
	Activationtimestr  string `json:"activationtimestr,omitempty"`
	Expirytimestr      string `json:"expirytimestr,omitempty"`
	Deletiontimestr    string `json:"deletiontimestr,omitempty"`
	Rolloverfailrc     string `json:"rolloverfailrc,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Dnspolicydnspolicylabelbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Dnspolicylabel struct {
	Labelname              string `json:"labelname,omitempty"`
	Transform              string `json:"transform,omitempty"`
	Newname                string `json:"newname,omitempty"`
	Numpol                 string `json:"numpol,omitempty"`
	Hits                   string `json:"hits,omitempty"`
	Priority               string `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Flowtype               string `json:"flowtype,omitempty"`
	Description            string `json:"description,omitempty"`
	Isdefault              string `json:"isdefault,omitempty"`
	Nextgenapiresource     string `json:"_nextgenapiresource,omitempty"`
}

type Dnsprofile struct {
	Dnsprofilename               string `json:"dnsprofilename,omitempty"`
	Recursiveresolution          string `json:"recursiveresolution,omitempty"`
	Dnsquerylogging              string `json:"dnsquerylogging,omitempty"`
	Dnsanswerseclogging          string `json:"dnsanswerseclogging,omitempty"`
	Dnsextendedlogging           string `json:"dnsextendedlogging,omitempty"`
	Dnserrorlogging              string `json:"dnserrorlogging,omitempty"`
	Cacherecords                 string `json:"cacherecords,omitempty"`
	Cachenegativeresponses       string `json:"cachenegativeresponses,omitempty"`
	Dropmultiqueryrequest        string `json:"dropmultiqueryrequest,omitempty"`
	Cacheecsresponses            string `json:"cacheecsresponses,omitempty"`
	Insertecs                    string `json:"insertecs,omitempty"`
	Replaceecs                   string `json:"replaceecs,omitempty"`
	Maxcacheableecsprefixlength  int    `json:"maxcacheableecsprefixlength,omitempty"`
	Maxcacheableecsprefixlength6 int    `json:"maxcacheableecsprefixlength6,omitempty"`
	Referencecount               string `json:"referencecount,omitempty"`
	Nextgenapiresource           string `json:"_nextgenapiresource,omitempty"`
}

type Dnsviewpolicybinding struct {
	Dnspolicyname string `json:"dnspolicyname,omitempty"`
	Viewname      string `json:"viewname,omitempty"`
}

type Dnsaction struct {
	Actionname         string   `json:"actionname,omitempty"`
	Actiontype         string   `json:"actiontype,omitempty"`
	Ipaddress          []string `json:"ipaddress,omitempty"`
	Ttl                int      `json:"ttl,omitempty"`
	Viewname           string   `json:"viewname,omitempty"`
	Preferredloclist   []string `json:"preferredloclist,omitempty"`
	Dnsprofilename     string   `json:"dnsprofilename,omitempty"`
	Drop               string   `json:"drop,omitempty"`
	Cachebypass        string   `json:"cachebypass,omitempty"`
	Builtin            string   `json:"builtin,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
}

type Dnsaction64 struct {
	Actionname         string `json:"actionname,omitempty"`
	Prefix             string `json:"prefix,omitempty"`
	Mappedrule         string `json:"mappedrule,omitempty"`
	Excluderule        string `json:"excluderule,omitempty"`
	Builtin            string `json:"builtin,omitempty"`
	Feature            string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Dnsviewbinding struct {
	Viewname string `json:"viewname,omitempty"`
}

type Dnszonebinding struct {
	Zonename string `json:"zonename,omitempty"`
}

type Dnspolicybinding struct {
	Name string `json:"name,omitempty"`
}

type Dnspolicyglobalbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Activepolicy           int32  `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Dnspolicylabelbinding struct {
	Labelname string `json:"labelname,omitempty"`
}

type Dnsptrrec struct {
	Reversedomain      string `json:"reversedomain,omitempty"`
	Domain             string `json:"domain,omitempty"`
	Ttl                int    `json:"ttl,omitempty"`
	Ecssubnet          string `json:"ecssubnet,omitempty"`
	Type               string `json:"type,omitempty"`
	Nodeid             int    `json:"nodeid,omitempty"`
	Authtype           string `json:"authtype,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Dnssubnetcache struct {
	Ecssubnet          string `json:"ecssubnet,omitempty"`
	All                bool   `json:"all,omitempty"`
	Nodeid             int    `json:"nodeid,omitempty"`
	Hostname           string `json:"hostname,omitempty"`
	Nextrecs           string `json:"nextrecs,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Dnstxtrec struct {
	Domain             string   `json:"domain,omitempty"`
	String             []string `json:"String,omitempty"`
	Ttl                int      `json:"ttl,omitempty"`
	Recordid           int      `json:"recordid,omitempty"`
	Ecssubnet          string   `json:"ecssubnet,omitempty"`
	Type               string   `json:"type,omitempty"`
	Nodeid             int      `json:"nodeid,omitempty"`
	Authtype           string   `json:"authtype,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
}

type Dnszone struct {
	Zonename           string   `json:"zonename,omitempty"`
	Proxymode          string   `json:"proxymode,omitempty"`
	Dnssecoffload      string   `json:"dnssecoffload,omitempty"`
	Nsec               string   `json:"nsec,omitempty"`
	Keyname            []string `json:"keyname,omitempty"`
	Type               string   `json:"type,omitempty"`
	Flags              string   `json:"flags,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
}

type Dnszonedomainbinding struct {
	Domain   string   `json:"domain,omitempty"`
	Nextrecs []string `json:"nextrecs,omitempty"`
	Zonename string   `json:"zonename,omitempty"`
}

type Dnsnsrec struct {
	Domain             string `json:"domain,omitempty"`
	Nameserver         string `json:"nameserver,omitempty"`
	Ttl                int    `json:"ttl,omitempty"`
	Ecssubnet          string `json:"ecssubnet,omitempty"`
	Type               string `json:"type,omitempty"`
	Nodeid             int    `json:"nodeid,omitempty"`
	Authtype           string `json:"authtype,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Dnsparameter struct {
	Retries                      int    `json:"retries,omitempty"`
	Minttl                       int    `json:"minttl"`
	Maxttl                       int    `json:"maxttl,omitempty"`
	Cacherecords                 string `json:"cacherecords,omitempty"`
	Namelookuppriority           string `json:"namelookuppriority,omitempty"`
	Recursion                    string `json:"recursion,omitempty"`
	Resolutionorder              string `json:"resolutionorder,omitempty"`
	Dnssec                       string `json:"dnssec,omitempty"`
	Maxpipeline                  int    `json:"maxpipeline,omitempty"`
	Dnsrootreferral              string `json:"dnsrootreferral,omitempty"`
	Dns64timeout                 int    `json:"dns64timeout"`
	Ecsmaxsubnets                int    `json:"ecsmaxsubnets"`
	Maxnegcachettl               int    `json:"maxnegcachettl,omitempty"`
	Cachehitbypass               string `json:"cachehitbypass,omitempty"`
	Maxcachesize                 int    `json:"maxcachesize"`
	Resolvermaxactiveresolutions int    `json:"resolvermaxactiveresolutions,omitempty"`
	Resolvermaxtcpconnections    int    `json:"resolvermaxtcpconnections,omitempty"`
	Resolvermaxtcptimeout        int    `json:"resolvermaxtcptimeout,omitempty"`
	Maxnegativecachesize         int    `json:"maxnegativecachesize"`
	Cachenoexpire                string `json:"cachenoexpire,omitempty"`
	Splitpktqueryprocessing      string `json:"splitpktqueryprocessing,omitempty"`
	Cacheecszeroprefix           string `json:"cacheecszeroprefix,omitempty"`
	Maxudppacketsize             int    `json:"maxudppacketsize,omitempty"`
	Zonetransfer                 string `json:"zonetransfer,omitempty"`
	Autosavekeyops               string `json:"autosavekeyops,omitempty"`
	Nxdomainratelimitthreshold   int    `json:"nxdomainratelimitthreshold"`
	Builtin                      string `json:"builtin,omitempty"`
	Feature                      string `json:"feature,omitempty"`
	Nxdomainthresholdcrossed     string `json:"nxdomainthresholdcrossed,omitempty"`
	Nextgenapiresource           string `json:"_nextgenapiresource,omitempty"`
}

type Dnspolicy struct {
	Name               string   `json:"name,omitempty"`
	Rule               string   `json:"rule,omitempty"`
	Viewname           string   `json:"viewname,omitempty"`
	Preferredlocation  string   `json:"preferredlocation,omitempty"`
	Preferredloclist   []string `json:"preferredloclist,omitempty"`
	Drop               string   `json:"drop,omitempty"`
	Cachebypass        string   `json:"cachebypass,omitempty"`
	Actionname         string   `json:"actionname,omitempty"`
	Logaction          string   `json:"logaction,omitempty"`
	Hits               string   `json:"hits,omitempty"`
	Undefhits          string   `json:"undefhits,omitempty"`
	Description        string   `json:"description,omitempty"`
	Builtin            string   `json:"builtin,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
}

type Dnsnameserver struct {
	Ip                 string `json:"ip,omitempty"`
	Dnsvservername     string `json:"dnsvservername,omitempty"`
	Local              bool   `json:"local,omitempty"`
	State              string `json:"state,omitempty"`
	Type               string `json:"type,omitempty"`
	Dnsprofilename     string `json:"dnsprofilename,omitempty"`
	Servicename        string `json:"servicename,omitempty"`
	Port               string `json:"port,omitempty"`
	Nameserverstate    string `json:"nameserverstate,omitempty"`
	Clmonowner         string `json:"clmonowner,omitempty"`
	Clmonview          string `json:"clmonview,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Dnsnaptrrec struct {
	Domain             string `json:"domain,omitempty"`
	Order              int    `json:"order,omitempty"`
	Preference         int    `json:"preference,omitempty"`
	Flags              string `json:"flags,omitempty"`
	Services           string `json:"services,omitempty"`
	Regexp             string `json:"regexp,omitempty"`
	Replacement        string `json:"replacement,omitempty"`
	Ttl                int    `json:"ttl,omitempty"`
	Recordid           int    `json:"recordid,omitempty"`
	Ecssubnet          string `json:"ecssubnet,omitempty"`
	Type               string `json:"type,omitempty"`
	Nodeid             int    `json:"nodeid,omitempty"`
	Authtype           string `json:"authtype,omitempty"`
	Vservername        string `json:"vservername,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Dnsmxrec struct {
	Domain             string `json:"domain,omitempty"`
	Mx                 string `json:"mx,omitempty"`
	Pref               int    `json:"pref,omitempty"`
	Ttl                int    `json:"ttl,omitempty"`
	Ecssubnet          string `json:"ecssubnet,omitempty"`
	Type               string `json:"type,omitempty"`
	Nodeid             int    `json:"nodeid,omitempty"`
	Authtype           string `json:"authtype,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Dnspolicy64vserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Dnspolicydnsglobalbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Activepolicy           int    `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Dnspolicylabelpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Dnszonekeybinding struct {
	Keyname          []string `json:"keyname,omitempty"`
	Siginceptiontime []uint32 `json:"siginceptiontime,omitempty"`
	Signed           uint32   `json:"signed,omitempty"`
	Expires          uint32   `json:"expires,omitempty"`
	Zonename         string   `json:"zonename,omitempty"`
}

type Dnscaarec struct {
	Domain             string `json:"domain,omitempty"`
	Valuestring        string `json:"valuestring,omitempty"`
	Tag                string `json:"tag,omitempty"`
	Flag               string `json:"flag,omitempty"`
	Ttl                int    `json:"ttl,omitempty"`
	Recordid           int    `json:"recordid,omitempty"`
	Ecssubnet          string `json:"ecssubnet,omitempty"`
	Type               string `json:"type,omitempty"`
	Nodeid             int    `json:"nodeid,omitempty"`
	Authtype           string `json:"authtype,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Dnsnegativecacherecords struct {
	Nodeid             int    `json:"nodeid,omitempty"`
	Hostname           string `json:"hostname,omitempty"`
	Negcachetype       string `json:"negcachetype,omitempty"`
	Rdclient           string `json:"rdclient,omitempty"`
	Querytype          string `json:"querytype,omitempty"`
	Ttl                string `json:"ttl,omitempty"`
	Ecssubnet          string `json:"ecssubnet,omitempty"`
	Peid               string `json:"peid,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Dnspolicylabeldnspolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Dnssuffix struct {
	Dnssuffix          string `json:"Dnssuffix,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Dnsviewdnspolicybinding struct {
	Dnspolicyname string `json:"dnspolicyname,omitempty"`
	Viewname      string `json:"viewname,omitempty"`
}

type Dnszonednskeybinding struct {
	Keyname          []string `json:"keyname,omitempty"`
	Siginceptiontime []int    `json:"siginceptiontime,omitempty"`
	Signed           int      `json:"signed,omitempty"`
	Expires          int      `json:"expires,omitempty"`
	Zonename         string   `json:"zonename,omitempty"`
}

type Dnsaaaarec struct {
	Hostname           string `json:"hostname,omitempty"`
	Ipv6address        string `json:"ipv6address,omitempty"`
	Ttl                int    `json:"ttl,omitempty"`
	Ecssubnet          string `json:"ecssubnet,omitempty"`
	Type               string `json:"type,omitempty"`
	Nodeid             int    `json:"nodeid,omitempty"`
	Vservername        string `json:"vservername,omitempty"`
	Authtype           string `json:"authtype,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Dnsdsfile struct {
	Keyname            string `json:"keyname,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Dnsglobalpolicybinding struct {
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

type Dnsproxyrecords struct {
	Type       string `json:"type,omitempty"`
	Negrectype string `json:"negrectype,omitempty"`
}

type Dnssoarec struct {
	Domain             string `json:"domain,omitempty"`
	Originserver       string `json:"originserver,omitempty"`
	Contact            string `json:"contact,omitempty"`
	Serial             int    `json:"serial,omitempty"`
	Refresh            int    `json:"refresh,omitempty"`
	Retry              int    `json:"retry,omitempty"`
	Expire             int    `json:"expire,omitempty"`
	Minimum            int    `json:"minimum,omitempty"`
	Ttl                int    `json:"ttl,omitempty"`
	Ecssubnet          string `json:"ecssubnet,omitempty"`
	Type               string `json:"type,omitempty"`
	Nodeid             int    `json:"nodeid,omitempty"`
	Authtype           string `json:"authtype,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Dnssrvrec struct {
	Domain             string `json:"domain,omitempty"`
	Target             string `json:"target,omitempty"`
	Priority           int    `json:"priority,omitempty"`
	Weight             int    `json:"weight,omitempty"`
	Port               int    `json:"port,omitempty"`
	Ttl                int    `json:"ttl,omitempty"`
	Ecssubnet          string `json:"ecssubnet,omitempty"`
	Type               string `json:"type,omitempty"`
	Nodeid             int    `json:"nodeid,omitempty"`
	Authtype           string `json:"authtype,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Dnsviewgslbservicebinding struct {
	Gslbservicename string `json:"gslbservicename,omitempty"`
	Ipaddress       string `json:"ipaddress,omitempty"`
	Viewname        string `json:"viewname,omitempty"`
}

type Dnsviewservicebinding struct {
	Gslbservicename string `json:"gslbservicename,omitempty"`
	Ipaddress       string `json:"ipaddress,omitempty"`
	Viewname        string `json:"viewname,omitempty"`
}

type Dnsglobalbinding struct {
}

type Dnsnsecrec struct {
	Hostname           string `json:"hostname,omitempty"`
	Type               string `json:"type,omitempty"`
	Nextnsec           string `json:"nextnsec,omitempty"`
	Nextrecs           string `json:"nextrecs,omitempty"`
	Ttl                string `json:"ttl,omitempty"`
	Ecssubnet          string `json:"ecssubnet,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Dnspolicy64 struct {
	Name               string `json:"name,omitempty"`
	Rule               string `json:"rule,omitempty"`
	Action             string `json:"action,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Labeltype          string `json:"labeltype,omitempty"`
	Labelname          string `json:"labelname,omitempty"`
	Undefhits          string `json:"undefhits,omitempty"`
	Description        string `json:"description,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Dnspolicy64binding struct {
	Name string `json:"name,omitempty"`
}

type Dnspolicy64lbvserverbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Dnspolicypolicylabelbinding struct {
	Boundto                string `json:"boundto,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Activepolicy           int32  `json:"activepolicy,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Dnspolicylabelpolicybindingbinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
}

type Dnsview struct {
	Viewname           string `json:"viewname,omitempty"`
	Flags              string `json:"flags,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}
