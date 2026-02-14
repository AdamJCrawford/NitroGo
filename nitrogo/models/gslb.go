// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Gslbservicegrouplbmonitorbinding struct {
	Monitorname      string `json:"monitor_name,omitempty"`
	Monweight        int    `json:"monweight,omitempty"`
	Monstate         string `json:"monstate,omitempty"`
	Weight           int    `json:"weight,omitempty"`
	Passive          bool   `json:"passive,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Port             int    `json:"port,omitempty"`
	State            string `json:"state,omitempty"`
	Hashid           int    `json:"hashid,omitempty"`
	Publicip         string `json:"publicip,omitempty"`
	Publicport       int    `json:"publicport,omitempty"`
	Siteprefix       string `json:"siteprefix,omitempty"`
	Order            int    `json:"order,omitempty"`
}

type Gslbconfig struct {
	Preview    bool   `json:"preview,omitempty"`
	Debug      bool   `json:"debug,omitempty"`
	Forcesync  string `json:"forcesync,omitempty"`
	Nowarn     bool   `json:"nowarn,omitempty"`
	Saveconfig bool   `json:"saveconfig,omitempty"`
	Command    string `json:"command,omitempty"`
}

type Gslbdomain struct {
	Name               string `json:"name,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Gslbdomainservicebinding struct {
	Servicename      string `json:"servicename,omitempty"`
	Servicetype      string `json:"servicetype,omitempty"`
	Vservername      string `json:"vservername,omitempty"`
	Ipaddress        string `json:"ipaddress,omitempty"`
	Port             int32  `json:"port,omitempty"`
	State            string `json:"state,omitempty"`
	Weight           uint32 `json:"weight,omitempty"`
	Dynamicconfwt    uint32 `json:"dynamicconfwt,omitempty"`
	Cumulativeweight uint32 `json:"cumulativeweight,omitempty"`
	Svreffgslbstate  string `json:"svreffgslbstate,omitempty"`
	Gslbthreshold    int32  `json:"gslbthreshold,omitempty"`
	Cnameentry       string `json:"cnameentry,omitempty"`
	Name             string `json:"name,omitempty"`
}

type Gslbservicelbmonitorbinding struct {
	Monitorname                string `json:"monitor_name,omitempty"`
	Monstate                   string `json:"monstate,omitempty"`
	Monitorstate               string `json:"monitor_state,omitempty"`
	Weight                     int    `json:"weight,omitempty"`
	Totalfailedprobes          int    `json:"totalfailedprobes,omitempty"`
	Failedprobes               int    `json:"failedprobes,omitempty"`
	Monstatcode                int    `json:"monstatcode,omitempty"`
	Monstatparam1              int    `json:"monstatparam1,omitempty"`
	Monstatparam2              int    `json:"monstatparam2,omitempty"`
	Monstatparam3              int    `json:"monstatparam3,omitempty"`
	Responsetime               int    `json:"responsetime,omitempty"`
	Monitortotalprobes         int    `json:"monitortotalprobes,omitempty"`
	Monitortotalfailedprobes   int    `json:"monitortotalfailedprobes,omitempty"`
	Monitorcurrentfailedprobes int    `json:"monitorcurrentfailedprobes,omitempty"`
	Lastresponse               string `json:"lastresponse,omitempty"`
	Servicename                string `json:"servicename,omitempty"`
}

type Gslbvservergslbservicegroupmemberbinding struct {
	Servicegroupname   string `json:"servicegroupname,omitempty"`
	Ipaddress          string `json:"ipaddress,omitempty"`
	Port               int    `json:"port,omitempty"`
	Servicetype        string `json:"servicetype,omitempty"`
	Curstate           string `json:"curstate,omitempty"`
	Weight             int    `json:"weight,omitempty"`
	Dynamicweight      string `json:"dynamicweight,omitempty"`
	Preferredlocation  string `json:"preferredlocation,omitempty"`
	Svreffgslbstate    string `json:"svreffgslbstate,omitempty"`
	Thresholdvalue     int    `json:"thresholdvalue,omitempty"`
	Gslbthreshold      int    `json:"gslbthreshold,omitempty"`
	Sitepersistcookie  string `json:"sitepersistcookie,omitempty"`
	Svcsitepersistence string `json:"svcsitepersistence,omitempty"`
	Order              int    `json:"order,omitempty"`
	Orderstr           string `json:"orderstr,omitempty"`
	Name               string `json:"name,omitempty"`
}

type Gslbdomaingslbservicebinding struct {
	Servicename      string `json:"servicename,omitempty"`
	Servicetype      string `json:"servicetype,omitempty"`
	Vservername      string `json:"vservername,omitempty"`
	Ipaddress        string `json:"ipaddress,omitempty"`
	Port             int    `json:"port,omitempty"`
	State            string `json:"state,omitempty"`
	Weight           int    `json:"weight,omitempty"`
	Dynamicconfwt    int    `json:"dynamicconfwt,omitempty"`
	Cumulativeweight int    `json:"cumulativeweight,omitempty"`
	Svreffgslbstate  string `json:"svreffgslbstate,omitempty"`
	Gslbthreshold    int    `json:"gslbthreshold,omitempty"`
	Cnameentry       string `json:"cnameentry,omitempty"`
	Order            int    `json:"order,omitempty"`
	Name             string `json:"name,omitempty"`
}

type Gslbdomaingslbservicegroupmemberbinding struct {
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Ipaddress        string `json:"ipaddress,omitempty"`
	Port             int    `json:"port,omitempty"`
	Servicetype      string `json:"servicetype,omitempty"`
	Weight           int    `json:"weight,omitempty"`
	Svreffgslbstate  string `json:"svreffgslbstate,omitempty"`
	Gslbthreshold    int    `json:"gslbthreshold,omitempty"`
	Order            int    `json:"order,omitempty"`
	Name             string `json:"name,omitempty"`
}

type Gslbrunningconfig struct {
	Response           string `json:"response,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Gslbservicebinding struct {
	Servicename string `json:"servicename,omitempty"`
}

type Gslbsitegslbservicegroupmemberbinding struct {
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Ipaddress        string `json:"ipaddress,omitempty"`
	Port             int    `json:"port,omitempty"`
	Servicetype      string `json:"servicetype,omitempty"`
	State            string `json:"state,omitempty"`
	Sitename         string `json:"sitename,omitempty"`
}

type Gslbvservergslbservicegroupbinding struct {
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Order            int    `json:"order,omitempty"`
	Name             string `json:"name,omitempty"`
}

type Gslbvserverservicegroupbinding struct {
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Name             string `json:"name,omitempty"`
}

type Gslbdomaingslbvserverbinding struct {
	Vservername        string `json:"vservername,omitempty"`
	Servicetype        string `json:"servicetype,omitempty"`
	State              string `json:"state,omitempty"`
	Lbmethod           string `json:"lbmethod,omitempty"`
	Dnsrecordtype      string `json:"dnsrecordtype,omitempty"`
	Backuplbmethod     string `json:"backuplbmethod,omitempty"`
	Persistencetype    string `json:"persistencetype,omitempty"`
	Edr                string `json:"edr,omitempty"`
	Mir                string `json:"mir,omitempty"`
	Dynamicweight      string `json:"dynamicweight,omitempty"`
	Statechangetimesec string `json:"statechangetimesec,omitempty"`
	Cip                string `json:"cip,omitempty"`
	Persistenceid      int    `json:"persistenceid,omitempty"`
	Netmask            string `json:"netmask,omitempty"`
	V6netmasklen       int    `json:"v6netmasklen,omitempty"`
	Sitename           string `json:"sitename,omitempty"`
	Sitepersistence    string `json:"sitepersistence,omitempty"`
	Siteprefix         string `json:"siteprefix,omitempty"`
	Customheaders      string `json:"customheaders,omitempty"`
	Persistmask        string `json:"persistmask,omitempty"`
	V6persistmasklen   int    `json:"v6persistmasklen,omitempty"`
	Name               string `json:"name,omitempty"`
}

type Gslbldnsentry struct {
	Ipaddress string `json:"ipaddress,omitempty"`
}

type Gslbservice struct {
	Servicename               string `json:"servicename,omitempty"`
	Cnameentry                string `json:"cnameentry,omitempty"`
	Ip                        string `json:"ip,omitempty"`
	Servername                string `json:"servername,omitempty"`
	Servicetype               string `json:"servicetype,omitempty"`
	Port                      int    `json:"port,omitempty"`
	Publicip                  string `json:"publicip,omitempty"`
	Publicport                int    `json:"publicport,omitempty"`
	Maxclient                 int    `json:"maxclient,omitempty"`
	Healthmonitor             string `json:"healthmonitor,omitempty"`
	Sitename                  string `json:"sitename,omitempty"`
	State                     string `json:"state,omitempty"`
	Cip                       string `json:"cip,omitempty"`
	Cipheader                 string `json:"cipheader,omitempty"`
	Sitepersistence           string `json:"sitepersistence,omitempty"`
	Cookietimeout             int    `json:"cookietimeout,omitempty"`
	Siteprefix                string `json:"siteprefix,omitempty"`
	Clttimeout                int    `json:"clttimeout,omitempty"`
	Svrtimeout                int    `json:"svrtimeout,omitempty"`
	Maxbandwidth              int    `json:"maxbandwidth,omitempty"`
	Downstateflush            string `json:"downstateflush,omitempty"`
	Maxaaausers               int    `json:"maxaaausers,omitempty"`
	Monthreshold              int    `json:"monthreshold,omitempty"`
	Hashid                    int    `json:"hashid,omitempty"`
	Comment                   string `json:"comment,omitempty"`
	Appflowlog                string `json:"appflowlog,omitempty"`
	Naptrreplacement          string `json:"naptrreplacement,omitempty"`
	Naptrorder                int    `json:"naptrorder,omitempty"`
	Naptrservices             string `json:"naptrservices,omitempty"`
	Naptrdomainttl            int    `json:"naptrdomainttl,omitempty"`
	Naptrpreference           int    `json:"naptrpreference,omitempty"`
	Ipaddress                 string `json:"ipaddress,omitempty"`
	Viewname                  string `json:"viewname,omitempty"`
	Viewip                    string `json:"viewip,omitempty"`
	Weight                    int    `json:"weight,omitempty"`
	Monitornamesvc            string `json:"monitor_name_svc,omitempty"`
	Newname                   string `json:"newname,omitempty"`
	Gslb                      string `json:"gslb,omitempty"`
	Svrstate                  string `json:"svrstate,omitempty"`
	Svreffgslbstate           string `json:"svreffgslbstate,omitempty"`
	Gslbthreshold             string `json:"gslbthreshold,omitempty"`
	Gslbsvcstats              string `json:"gslbsvcstats,omitempty"`
	Monstate                  string `json:"monstate,omitempty"`
	Preferredlocation         string `json:"preferredlocation,omitempty"`
	Monitorstate              string `json:"monitor_state,omitempty"`
	Statechangetimesec        string `json:"statechangetimesec,omitempty"`
	Tickssincelaststatechange string `json:"tickssincelaststatechange,omitempty"`
	Threshold                 string `json:"threshold,omitempty"`
	Clmonowner                string `json:"clmonowner,omitempty"`
	Clmonview                 string `json:"clmonview,omitempty"`
	Gslbsvchealth             string `json:"gslbsvchealth,omitempty"`
	Glsbsvchealthdescr        string `json:"glsbsvchealthdescr,omitempty"`
	Nodefaultbindings         string `json:"nodefaultbindings,omitempty"`
	Nextgenapiresource        string `json:"_nextgenapiresource,omitempty"`
}

type Gslbservicegroupbinding struct {
	Servicegroupname string `json:"servicegroupname,omitempty"`
}

type Gslbvserverservicegroupmemberbinding struct {
	Servicegroupname   string `json:"servicegroupname,omitempty"`
	Ipaddress          string `json:"ipaddress,omitempty"`
	Port               int32  `json:"port,omitempty"`
	Servicetype        string `json:"servicetype,omitempty"`
	Curstate           string `json:"curstate,omitempty"`
	Weight             uint32 `json:"weight,omitempty"`
	Dynamicweight      string `json:"dynamicweight,omitempty"`
	Preferredlocation  string `json:"preferredlocation,omitempty"`
	Svreffgslbstate    string `json:"svreffgslbstate,omitempty"`
	Thresholdvalue     int32  `json:"thresholdvalue,omitempty"`
	Gslbthreshold      int32  `json:"gslbthreshold,omitempty"`
	Sitepersistcookie  string `json:"sitepersistcookie,omitempty"`
	Svcsitepersistence string `json:"svcsitepersistence,omitempty"`
	Name               string `json:"name,omitempty"`
}

type Gslbdomainservicegroupbinding struct {
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Name             string `json:"name,omitempty"`
}

type Gslbdomainvserverbinding struct {
	Vservername        string `json:"vservername,omitempty"`
	Servicetype        string `json:"servicetype,omitempty"`
	State              string `json:"state,omitempty"`
	Lbmethod           string `json:"lbmethod,omitempty"`
	Dnsrecordtype      string `json:"dnsrecordtype,omitempty"`
	Backuplbmethod     string `json:"backuplbmethod,omitempty"`
	Persistencetype    string `json:"persistencetype,omitempty"`
	Edr                string `json:"edr,omitempty"`
	Mir                string `json:"mir,omitempty"`
	Dynamicweight      string `json:"dynamicweight,omitempty"`
	Statechangetimesec string `json:"statechangetimesec,omitempty"`
	Cip                string `json:"cip,omitempty"`
	Persistenceid      uint32 `json:"persistenceid,omitempty"`
	Netmask            string `json:"netmask,omitempty"`
	V6netmasklen       uint32 `json:"v6netmasklen,omitempty"`
	Sitename           string `json:"sitename,omitempty"`
	Sitepersistence    string `json:"sitepersistence,omitempty"`
	Siteprefix         string `json:"siteprefix,omitempty"`
	Customheaders      string `json:"customheaders,omitempty"`
	Persistmask        string `json:"persistmask,omitempty"`
	V6persistmasklen   uint32 `json:"v6persistmasklen,omitempty"`
	Name               string `json:"name,omitempty"`
}

type Gslbservicemonitorbinding struct {
	Monitorname                string `json:"monitor_name,omitempty"`
	Monstate                   string `json:"monstate,omitempty"`
	Monitorstate               string `json:"monitor_state,omitempty"`
	Weight                     uint32 `json:"weight,omitempty"`
	Totalfailedprobes          uint32 `json:"totalfailedprobes,omitempty"`
	Failedprobes               uint32 `json:"failedprobes,omitempty"`
	Monstatcode                int32  `json:"monstatcode,omitempty"`
	Monstatparam1              int32  `json:"monstatparam1,omitempty"`
	Monstatparam2              int32  `json:"monstatparam2,omitempty"`
	Monstatparam3              int32  `json:"monstatparam3,omitempty"`
	Responsetime               uint64 `json:"responsetime,omitempty"`
	Monitortotalprobes         uint32 `json:"monitortotalprobes,omitempty"`
	Monitortotalfailedprobes   uint32 `json:"monitortotalfailedprobes,omitempty"`
	Monitorcurrentfailedprobes uint32 `json:"monitorcurrentfailedprobes,omitempty"`
	Lastresponse               string `json:"lastresponse,omitempty"`
	Servicename                string `json:"servicename,omitempty"`
}

type Gslbservicegroup struct {
	Servicegroupname           string `json:"servicegroupname,omitempty"`
	Servicetype                string `json:"servicetype,omitempty"`
	Maxclient                  int    `json:"maxclient,omitempty"`
	Cip                        string `json:"cip,omitempty"`
	Cipheader                  string `json:"cipheader,omitempty"`
	Healthmonitor              string `json:"healthmonitor,omitempty"`
	Clttimeout                 int    `json:"clttimeout,omitempty"`
	Svrtimeout                 int    `json:"svrtimeout,omitempty"`
	Maxbandwidth               int    `json:"maxbandwidth,omitempty"`
	Monthreshold               int    `json:"monthreshold,omitempty"`
	State                      string `json:"state,omitempty"`
	Downstateflush             string `json:"downstateflush,omitempty"`
	Comment                    string `json:"comment,omitempty"`
	Appflowlog                 string `json:"appflowlog,omitempty"`
	Autoscale                  string `json:"autoscale,omitempty"`
	Autodelayedtrofs           string `json:"autodelayedtrofs,omitempty"`
	Sitename                   string `json:"sitename,omitempty"`
	Sitepersistence            string `json:"sitepersistence,omitempty"`
	Servername                 string `json:"servername,omitempty"`
	Port                       int    `json:"port,omitempty"`
	Weight                     int    `json:"weight,omitempty"`
	Hashid                     int    `json:"hashid,omitempty"`
	Publicip                   string `json:"publicip,omitempty"`
	Publicport                 int    `json:"publicport,omitempty"`
	Siteprefix                 string `json:"siteprefix,omitempty"`
	Order                      int    `json:"order,omitempty"`
	Monitornamesvc             string `json:"monitor_name_svc,omitempty"`
	Dupweight                  int    `json:"dup_weight,omitempty"`
	Delay                      int    `json:"delay,omitempty"`
	Graceful                   string `json:"graceful,omitempty"`
	Includemembers             bool   `json:"includemembers,omitempty"`
	Newname                    string `json:"newname,omitempty"`
	Numofconnections           string `json:"numofconnections,omitempty"`
	Serviceconftype            string `json:"serviceconftype,omitempty"`
	Value                      string `json:"value,omitempty"`
	Svrstate                   string `json:"svrstate,omitempty"`
	Ip                         string `json:"ip,omitempty"`
	Monstatcode                string `json:"monstatcode,omitempty"`
	Monstatparam1              string `json:"monstatparam1,omitempty"`
	Monstatparam2              string `json:"monstatparam2,omitempty"`
	Monstatparam3              string `json:"monstatparam3,omitempty"`
	Statechangetimemsec        string `json:"statechangetimemsec,omitempty"`
	Stateupdatereason          string `json:"stateupdatereason,omitempty"`
	Clmonowner                 string `json:"clmonowner,omitempty"`
	Clmonview                  string `json:"clmonview,omitempty"`
	Groupcount                 string `json:"groupcount,omitempty"`
	Serviceipstr               string `json:"serviceipstr,omitempty"`
	Servicegroupeffectivestate string `json:"servicegroupeffectivestate,omitempty"`
	Gslb                       string `json:"gslb,omitempty"`
	Svreffgslbstate            string `json:"svreffgslbstate,omitempty"`
	Nodefaultbindings          string `json:"nodefaultbindings,omitempty"`
	Nextgenapiresource         string `json:"_nextgenapiresource,omitempty"`
}

type Gslbservicegroupservicegroupmemberbinding struct {
	Ip                        string `json:"ip,omitempty"`
	Port                      int32  `json:"port,omitempty"`
	Svrstate                  string `json:"svrstate,omitempty"`
	Statechangetimesec        string `json:"statechangetimesec,omitempty"`
	Tickssincelaststatechange uint32 `json:"tickssincelaststatechange,omitempty"`
	Weight                    uint32 `json:"weight,omitempty"`
	Servername                string `json:"servername,omitempty"`
	State                     string `json:"state,omitempty"`
	Hashid                    uint32 `json:"hashid,omitempty"`
	Graceful                  string `json:"graceful,omitempty"`
	Delay                     uint64 `json:"delay,omitempty"`
	Publicip                  string `json:"publicip,omitempty"`
	Publicport                int32  `json:"publicport,omitempty"`
	Gslbthreshold             int32  `json:"gslbthreshold,omitempty"`
	Threshold                 string `json:"threshold,omitempty"`
	Preferredlocation         string `json:"preferredlocation,omitempty"`
	Siteprefix                string `json:"siteprefix,omitempty"`
	Servicegroupname          string `json:"servicegroupname,omitempty"`
}

type Gslbsitebinding struct {
	Sitename string `json:"sitename,omitempty"`
}

type Gslbsiteservicegroupmemberbinding struct {
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Ipaddress        string `json:"ipaddress,omitempty"`
	Port             int32  `json:"port,omitempty"`
	Servicetype      string `json:"servicetype,omitempty"`
	State            string `json:"state,omitempty"`
	Sitename         string `json:"sitename,omitempty"`
}

type Gslbvserverbinding struct {
	Name string `json:"name,omitempty"`
}

type Gslbdomainlbmonitorbinding struct {
	Monitorname                string `json:"monitorname,omitempty"`
	Servicename                string `json:"servicename,omitempty"`
	Vservername                string `json:"vservername,omitempty"`
	Monstate                   string `json:"monstate,omitempty"`
	Httprequest                string `json:"httprequest,omitempty"`
	Iptunnel                   string `json:"iptunnel,omitempty"`
	Customheaders              string `json:"customheaders,omitempty"`
	Respcode                   string `json:"respcode,omitempty"`
	Monitortotalprobes         int    `json:"monitortotalprobes,omitempty"`
	Monitortotalfailedprobes   int    `json:"monitortotalfailedprobes,omitempty"`
	Monitorcurrentfailedprobes int    `json:"monitorcurrentfailedprobes,omitempty"`
	Responsetime               int    `json:"responsetime,omitempty"`
	Monstatcode                int    `json:"monstatcode,omitempty"`
	Lastresponse               string `json:"lastresponse,omitempty"`
	Grpchealthcheck            string `json:"grpchealthcheck,omitempty"`
	Grpcstatuscode             int    `json:"grpcstatuscode,omitempty"`
	Grpcservicename            string `json:"grpcservicename,omitempty"`
	Name                       string `json:"name,omitempty"`
}

type Gslbldnsentries struct {
	Nodeid             int    `json:"nodeid,omitempty"`
	Sitename           string `json:"sitename,omitempty"`
	Numsites           string `json:"numsites,omitempty"`
	Ipaddress          string `json:"ipaddress,omitempty"`
	Ttl                string `json:"ttl,omitempty"`
	Name               string `json:"name,omitempty"`
	Rtt                string `json:"rtt,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Gslbservicednsviewbinding struct {
	Viewname    string `json:"viewname,omitempty"`
	Viewip      string `json:"viewip,omitempty"`
	Servicename string `json:"servicename,omitempty"`
}

type Gslbserviceviewbinding struct {
	Viewname    string `json:"viewname,omitempty"`
	Viewip      string `json:"viewip,omitempty"`
	Servicename string `json:"servicename,omitempty"`
}

type Gslbsitegslbservicebinding struct {
	Servicename string `json:"servicename,omitempty"`
	Cnameentry  string `json:"cnameentry,omitempty"`
	Ipaddress   string `json:"ipaddress,omitempty"`
	Port        int    `json:"port,omitempty"`
	Servicetype string `json:"servicetype,omitempty"`
	State       string `json:"state,omitempty"`
	Sitename    string `json:"sitename,omitempty"`
}

type Gslbsyncstatus struct {
	Summary            bool   `json:"summary,omitempty"`
	Response           string `json:"response,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Gslbvserverdomainbinding struct {
	Domainname       string `json:"domainname,omitempty"`
	Ttl              int    `json:"ttl,omitempty"`
	Backupip         string `json:"backupip,omitempty"`
	Cookiedomain     string `json:"cookie_domain,omitempty"`
	Cookietimeout    int    `json:"cookietimeout,omitempty"`
	Sitedomainttl    int    `json:"sitedomainttl,omitempty"`
	Name             string `json:"name,omitempty"`
	Order            int    `json:"order,omitempty"`
	Backupipflag     bool   `json:"backupipflag,omitempty"`
	Cookiedomainflag bool   `json:"cookie_domainflag,omitempty"`
}

type Gslbvservergslbservicebinding struct {
	Servicename        string `json:"servicename,omitempty"`
	Weight             int    `json:"weight,omitempty"`
	Cnameentry         string `json:"cnameentry,omitempty"`
	Ipaddress          string `json:"ipaddress,omitempty"`
	Port               int    `json:"port,omitempty"`
	Gslbboundsvctype   string `json:"gslbboundsvctype,omitempty"`
	Curstate           string `json:"curstate,omitempty"`
	Dynamicconfwt      int    `json:"dynamicconfwt,omitempty"`
	Cumulativeweight   int    `json:"cumulativeweight,omitempty"`
	Svreffgslbstate    string `json:"svreffgslbstate,omitempty"`
	Gslbthreshold      int    `json:"gslbthreshold,omitempty"`
	Preferredlocation  string `json:"preferredlocation,omitempty"`
	Thresholdvalue     int    `json:"thresholdvalue,omitempty"`
	Iscname            string `json:"iscname,omitempty"`
	Domainname         string `json:"domainname,omitempty"`
	Sitepersistcookie  string `json:"sitepersistcookie,omitempty"`
	Svcsitepersistence string `json:"svcsitepersistence,omitempty"`
	Order              int    `json:"order,omitempty"`
	Orderstr           string `json:"orderstr,omitempty"`
	Name               string `json:"name,omitempty"`
}

type Gslbdomainbinding struct {
	Name string `json:"name,omitempty"`
}

type Gslbdomainservicegroupmemberbinding struct {
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Ipaddress        string `json:"ipaddress,omitempty"`
	Port             int32  `json:"port,omitempty"`
	Servicetype      string `json:"servicetype,omitempty"`
	Weight           uint32 `json:"weight,omitempty"`
	Svreffgslbstate  string `json:"svreffgslbstate,omitempty"`
	Gslbthreshold    int32  `json:"gslbthreshold,omitempty"`
	Name             string `json:"name,omitempty"`
}

type Gslbparameter struct {
	Ldnsentrytimeout            int      `json:"ldnsentrytimeout,omitempty"`
	Rtttolerance                int      `json:"rtttolerance,omitempty"`
	Ldnsmask                    string   `json:"ldnsmask,omitempty"`
	V6ldnsmasklen               int      `json:"v6ldnsmasklen,omitempty"`
	Ldnsprobeorder              []string `json:"ldnsprobeorder,omitempty"`
	Dropldnsreq                 string   `json:"dropldnsreq,omitempty"`
	Gslbsvcstatedelaytime       int      `json:"gslbsvcstatedelaytime,omitempty"`
	Svcstatelearningtime        int      `json:"svcstatelearningtime,omitempty"`
	Automaticconfigsync         string   `json:"automaticconfigsync,omitempty"`
	Mepkeepalivetimeout         int      `json:"mepkeepalivetimeout,omitempty"`
	Gslbsyncinterval            int      `json:"gslbsyncinterval,omitempty"`
	Gslbsyncmode                string   `json:"gslbsyncmode,omitempty"`
	Gslbsynclocfiles            string   `json:"gslbsynclocfiles,omitempty"`
	Gslbconfigsyncmonitor       string   `json:"gslbconfigsyncmonitor,omitempty"`
	Gslbsyncsaveconfigcommand   string   `json:"gslbsyncsaveconfigcommand,omitempty"`
	Undefaction                 string   `json:"undefaction,omitempty"`
	Flags                       string   `json:"flags,omitempty"`
	Builtin                     string   `json:"builtin,omitempty"`
	Feature                     string   `json:"feature,omitempty"`
	Incarnation                 string   `json:"incarnation,omitempty"`
	Overridepersistencyfororder string   `json:"overridepersistencyfororder,omitempty"`
	Nextgenapiresource          string   `json:"_nextgenapiresource,omitempty"`
}

type Gslbservicegroupgslbservicegroupmemberbinding struct {
	Ip                        string `json:"ip,omitempty"`
	Port                      int    `json:"port,omitempty"`
	Svrstate                  string `json:"svrstate,omitempty"`
	Statechangetimesec        string `json:"statechangetimesec,omitempty"`
	Tickssincelaststatechange int    `json:"tickssincelaststatechange,omitempty"`
	Weight                    int    `json:"weight,omitempty"`
	Servername                string `json:"servername,omitempty"`
	State                     string `json:"state,omitempty"`
	Hashid                    int    `json:"hashid,omitempty"`
	Graceful                  string `json:"graceful,omitempty"`
	Delay                     int    `json:"delay,omitempty"`
	Publicip                  string `json:"publicip,omitempty"`
	Publicport                int    `json:"publicport,omitempty"`
	Gslbthreshold             int    `json:"gslbthreshold,omitempty"`
	Threshold                 string `json:"threshold,omitempty"`
	Preferredlocation         string `json:"preferredlocation,omitempty"`
	Siteprefix                string `json:"siteprefix,omitempty"`
	Order                     int    `json:"order,omitempty"`
	Orderstr                  string `json:"orderstr,omitempty"`
	Trofsdelay                int    `json:"trofsdelay,omitempty"`
	Servicegroupname          string `json:"servicegroupname,omitempty"`
}

type Gslbsite struct {
	Sitename               string   `json:"sitename,omitempty"`
	Sitetype               string   `json:"sitetype,omitempty"`
	Siteipaddress          string   `json:"siteipaddress,omitempty"`
	Publicip               string   `json:"publicip,omitempty"`
	Metricexchange         string   `json:"metricexchange,omitempty"`
	Nwmetricexchange       string   `json:"nwmetricexchange,omitempty"`
	Sessionexchange        string   `json:"sessionexchange,omitempty"`
	Triggermonitor         string   `json:"triggermonitor,omitempty"`
	Parentsite             string   `json:"parentsite,omitempty"`
	Clip                   string   `json:"clip,omitempty"`
	Publicclip             string   `json:"publicclip,omitempty"`
	Naptrreplacementsuffix string   `json:"naptrreplacementsuffix,omitempty"`
	Backupparentlist       []string `json:"backupparentlist,omitempty"`
	Sitepassword           string   `json:"sitepassword,omitempty"`
	Newname                string   `json:"newname,omitempty"`
	Status                 string   `json:"status,omitempty"`
	Persistencemepstatus   string   `json:"persistencemepstatus,omitempty"`
	Version                string   `json:"version,omitempty"`
	Curbackupparentip      string   `json:"curbackupparentip,omitempty"`
	Sitestate              string   `json:"sitestate,omitempty"`
	Oldname                string   `json:"oldname,omitempty"`
	Nextgenapiresource     string   `json:"_nextgenapiresource,omitempty"`
}

type Gslbsitegslbservicegroupbinding struct {
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Servicetype      string `json:"servicetype,omitempty"`
	Sitename         string `json:"sitename,omitempty"`
}

type Gslbvserver struct {
	Name                      string `json:"name,omitempty"`
	Servicetype               string `json:"servicetype,omitempty"`
	Iptype                    string `json:"iptype,omitempty"`
	Dnsrecordtype             string `json:"dnsrecordtype,omitempty"`
	Lbmethod                  string `json:"lbmethod,omitempty"`
	Backupsessiontimeout      int    `json:"backupsessiontimeout,omitempty"`
	Backuplbmethod            string `json:"backuplbmethod,omitempty"`
	Netmask                   string `json:"netmask,omitempty"`
	V6netmasklen              int    `json:"v6netmasklen,omitempty"`
	Rule                      string `json:"rule,omitempty"`
	Tolerance                 int    `json:"tolerance,omitempty"`
	Persistencetype           string `json:"persistencetype,omitempty"`
	Persistenceid             int    `json:"persistenceid,omitempty"`
	Persistmask               string `json:"persistmask,omitempty"`
	V6persistmasklen          int    `json:"v6persistmasklen,omitempty"`
	Timeout                   int    `json:"timeout,omitempty"`
	Edr                       string `json:"edr,omitempty"`
	Ecs                       string `json:"ecs,omitempty"`
	Ecsaddrvalidation         string `json:"ecsaddrvalidation,omitempty"`
	Mir                       string `json:"mir,omitempty"`
	Disableprimaryondown      string `json:"disableprimaryondown,omitempty"`
	Dynamicweight             string `json:"dynamicweight,omitempty"`
	State                     string `json:"state,omitempty"`
	Considereffectivestate    string `json:"considereffectivestate,omitempty"`
	Comment                   string `json:"comment,omitempty"`
	Somethod                  string `json:"somethod,omitempty"`
	Sopersistence             string `json:"sopersistence,omitempty"`
	Sopersistencetimeout      int    `json:"sopersistencetimeout,omitempty"`
	Sothreshold               int    `json:"sothreshold,omitempty"`
	Sobackupaction            string `json:"sobackupaction,omitempty"`
	Appflowlog                string `json:"appflowlog,omitempty"`
	Toggleorder               string `json:"toggleorder,omitempty"`
	Orderthreshold            int    `json:"orderthreshold,omitempty"`
	Backupvserver             string `json:"backupvserver,omitempty"`
	Servicename               string `json:"servicename,omitempty"`
	Weight                    int    `json:"weight,omitempty"`
	Servicegroupname          string `json:"servicegroupname,omitempty"`
	Domainname                string `json:"domainname,omitempty"`
	Ttl                       int    `json:"ttl,omitempty"`
	Backupip                  string `json:"backupip,omitempty"`
	Cookiedomain              string `json:"cookie_domain,omitempty"`
	Cookietimeout             int    `json:"cookietimeout,omitempty"`
	Sitedomainttl             int    `json:"sitedomainttl,omitempty"`
	Order                     int    `json:"order,omitempty"`
	Newname                   string `json:"newname,omitempty"`
	Curstate                  string `json:"curstate,omitempty"`
	Status                    string `json:"status,omitempty"`
	Lbrrreason                string `json:"lbrrreason,omitempty"`
	Iscname                   string `json:"iscname,omitempty"`
	Sitepersistence           string `json:"sitepersistence,omitempty"`
	Totalservices             string `json:"totalservices,omitempty"`
	Activeservices            string `json:"activeservices,omitempty"`
	Statechangetimesec        string `json:"statechangetimesec,omitempty"`
	Statechangetimemsec       string `json:"statechangetimemsec,omitempty"`
	Tickssincelaststatechange string `json:"tickssincelaststatechange,omitempty"`
	Health                    string `json:"health,omitempty"`
	Policyname                string `json:"policyname,omitempty"`
	Priority                  string `json:"priority,omitempty"`
	Gotopriorityexpression    string `json:"gotopriorityexpression,omitempty"`
	Type                      string `json:"type,omitempty"`
	Vsvrbindsvcip             string `json:"vsvrbindsvcip,omitempty"`
	Vsvrbindsvcport           string `json:"vsvrbindsvcport,omitempty"`
	Servername                string `json:"servername,omitempty"`
	Nodefaultbindings         string `json:"nodefaultbindings,omitempty"`
	Currentactiveorder        string `json:"currentactiveorder,omitempty"`
	Nextgenapiresource        string `json:"_nextgenapiresource,omitempty"`
}

type Gslbvserverlbpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Type                   string `json:"type,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
}

type Gslbservicegroupmonitorbinding struct {
	Monitorname      string `json:"monitor_name,omitempty"`
	Monweight        uint32 `json:"monweight,omitempty"`
	Monstate         string `json:"monstate,omitempty"`
	Weight           uint32 `json:"weight,omitempty"`
	Passive          bool   `json:"passive,omitempty"`
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Port             int32  `json:"port,omitempty"`
	State            string `json:"state,omitempty"`
	Hashid           uint32 `json:"hashid,omitempty"`
	Publicip         string `json:"publicip,omitempty"`
	Publicport       int32  `json:"publicport,omitempty"`
	Siteprefix       string `json:"siteprefix,omitempty"`
}

type Gslbsiteservicebinding struct {
	Servicename string `json:"servicename,omitempty"`
	Cnameentry  string `json:"cnameentry,omitempty"`
	Ipaddress   string `json:"ipaddress,omitempty"`
	Port        int32  `json:"port,omitempty"`
	Servicetype string `json:"servicetype,omitempty"`
	State       string `json:"state,omitempty"`
	Sitename    string `json:"sitename,omitempty"`
}

type Gslbvserverservicebinding struct {
	Servicename        string `json:"servicename,omitempty"`
	Weight             uint32 `json:"weight,omitempty"`
	Cnameentry         string `json:"cnameentry,omitempty"`
	Ipaddress          string `json:"ipaddress,omitempty"`
	Port               int32  `json:"port,omitempty"`
	Gslbboundsvctype   string `json:"gslbboundsvctype,omitempty"`
	Curstate           string `json:"curstate,omitempty"`
	Dynamicconfwt      uint32 `json:"dynamicconfwt,omitempty"`
	Cumulativeweight   uint32 `json:"cumulativeweight,omitempty"`
	Svreffgslbstate    string `json:"svreffgslbstate,omitempty"`
	Gslbthreshold      int32  `json:"gslbthreshold,omitempty"`
	Preferredlocation  string `json:"preferredlocation,omitempty"`
	Thresholdvalue     int32  `json:"thresholdvalue,omitempty"`
	Iscname            string `json:"iscname,omitempty"`
	Domainname         string `json:"domainname,omitempty"`
	Sitepersistcookie  string `json:"sitepersistcookie,omitempty"`
	Svcsitepersistence string `json:"svcsitepersistence,omitempty"`
	Name               string `json:"name,omitempty"`
}

type Gslbvserverspilloverpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Type                   string `json:"type,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
}

type Gslbdomaingslbservicegroupbinding struct {
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Order            int    `json:"order,omitempty"`
	Name             string `json:"name,omitempty"`
}

type Gslbdomainmonitorbinding struct {
	Monitorname                string `json:"monitorname,omitempty"`
	Servicename                string `json:"servicename,omitempty"`
	Vservername                string `json:"vservername,omitempty"`
	Monstate                   string `json:"monstate,omitempty"`
	Httprequest                string `json:"httprequest,omitempty"`
	Iptunnel                   string `json:"iptunnel,omitempty"`
	Customheaders              string `json:"customheaders,omitempty"`
	Respcode                   string `json:"respcode,omitempty"`
	Monitortotalprobes         uint32 `json:"monitortotalprobes,omitempty"`
	Monitortotalfailedprobes   uint32 `json:"monitortotalfailedprobes,omitempty"`
	Monitorcurrentfailedprobes uint32 `json:"monitorcurrentfailedprobes,omitempty"`
	Responsetime               uint64 `json:"responsetime,omitempty"`
	Monstatcode                int32  `json:"monstatcode,omitempty"`
	Lastresponse               string `json:"lastresponse,omitempty"`
	Name                       string `json:"name,omitempty"`
}

type Gslbservicegroupservicegroupentitymonbindingsbinding struct {
	Servicegroupentname2       string `json:"servicegroupentname2,omitempty"`
	Monitorname                string `json:"monitor_name,omitempty"`
	Monitorstate               string `json:"monitor_state,omitempty"`
	Passive                    bool   `json:"passive,omitempty"`
	Monitortotalprobes         int    `json:"monitortotalprobes,omitempty"`
	Monitortotalfailedprobes   int    `json:"monitortotalfailedprobes,omitempty"`
	Monitorcurrentfailedprobes int    `json:"monitorcurrentfailedprobes,omitempty"`
	Lastresponse               string `json:"lastresponse,omitempty"`
	Servicegroupname           string `json:"servicegroupname,omitempty"`
	Port                       int    `json:"port,omitempty"`
	Weight                     int    `json:"weight,omitempty"`
	State                      string `json:"state,omitempty"`
	Hashid                     int    `json:"hashid,omitempty"`
	Publicip                   string `json:"publicip,omitempty"`
	Publicport                 int    `json:"publicport,omitempty"`
	Siteprefix                 string `json:"siteprefix,omitempty"`
	Order                      int    `json:"order,omitempty"`
}

type Gslbsiteservicegroupbinding struct {
	Servicegroupname string `json:"servicegroupname,omitempty"`
	Servicetype      string `json:"servicetype,omitempty"`
	Sitename         string `json:"sitename,omitempty"`
}

type Gslbvserverpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Type                   string `json:"type,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Name                   string `json:"name,omitempty"`
}
