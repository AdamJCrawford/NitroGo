// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Nshmackey struct {
	Name               string `json:"name,omitempty"`
	Digest             string `json:"digest,omitempty"`
	Keyvalue           string `json:"keyvalue,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nsconsoleloginprompt struct {
	Promptstring       string `json:"promptstring,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nscqaparam struct {
	Harqretxdelay      int     `json:"harqretxdelay,omitempty"`
	Net1label          string  `json:"net1label,omitempty"`
	Minrttnet1         int     `json:"minrttnet1,omitempty"`
	Lr1coeflist        string  `json:"lr1coeflist,omitempty"`
	Lr1probthresh      float64 `json:"lr1probthresh,omitempty"`
	Net1cclscale       string  `json:"net1cclscale,omitempty"`
	Net1csqscale       string  `json:"net1csqscale,omitempty"`
	Net1logcoef        string  `json:"net1logcoef,omitempty"`
	Net2label          string  `json:"net2label,omitempty"`
	Minrttnet2         int     `json:"minrttnet2,omitempty"`
	Lr2coeflist        string  `json:"lr2coeflist,omitempty"`
	Lr2probthresh      float64 `json:"lr2probthresh,omitempty"`
	Net2cclscale       string  `json:"net2cclscale,omitempty"`
	Net2csqscale       string  `json:"net2csqscale,omitempty"`
	Net2logcoef        string  `json:"net2logcoef,omitempty"`
	Net3label          string  `json:"net3label,omitempty"`
	Minrttnet3         int     `json:"minrttnet3,omitempty"`
	Net3cclscale       string  `json:"net3cclscale,omitempty"`
	Net3csqscale       string  `json:"net3csqscale,omitempty"`
	Net3logcoef        string  `json:"net3logcoef,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
}

type Nsicapprofile struct {
	Name                string `json:"name,omitempty"`
	Preview             string `json:"preview,omitempty"`
	Previewlength       int    `json:"previewlength,omitempty"`
	Uri                 string `json:"uri,omitempty"`
	Hostheader          string `json:"hostheader,omitempty"`
	Useragent           string `json:"useragent,omitempty"`
	Mode                string `json:"mode,omitempty"`
	Queryparams         string `json:"queryparams,omitempty"`
	Connectionkeepalive string `json:"connectionkeepalive,omitempty"`
	Allow204            string `json:"allow204,omitempty"`
	Inserticapheaders   string `json:"inserticapheaders,omitempty"`
	Inserthttprequest   string `json:"inserthttprequest,omitempty"`
	Reqtimeout          int    `json:"reqtimeout,omitempty"`
	Reqtimeoutaction    string `json:"reqtimeoutaction,omitempty"`
	Logaction           string `json:"logaction,omitempty"`
	Nextgenapiresource  string `json:"_nextgenapiresource,omitempty"`
}

type Nspbrs struct {
}

type Nssimpleacl struct {
	Aclname            string `json:"aclname,omitempty"`
	Aclaction          string `json:"aclaction,omitempty"`
	Td                 int    `json:"td,omitempty"`
	Srcip              string `json:"srcip,omitempty"`
	Destport           int    `json:"destport,omitempty"`
	Protocol           string `json:"protocol,omitempty"`
	Ttl                int    `json:"ttl,omitempty"`
	Estsessions        bool   `json:"estsessions,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nsweblogparam struct {
	Buffersizemb       int      `json:"buffersizemb,omitempty"`
	Customreqhdrs      []string `json:"customreqhdrs,omitempty"`
	Customrsphdrs      []string `json:"customrsphdrs,omitempty"`
	Builtin            string   `json:"builtin,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
}

type Nsdiameter struct {
	Identity               string `json:"identity,omitempty"`
	Realm                  string `json:"realm,omitempty"`
	Serverclosepropagation string `json:"serverclosepropagation,omitempty"`
	Ownernode              int    `json:"ownernode,omitempty"`
	Nextgenapiresource     string `json:"_nextgenapiresource,omitempty"`
}

type Nsextensionbinding struct {
	Name string `json:"name,omitempty"`
}

type Nslicenseserver struct {
	Licenseserverip    string `json:"licenseserverip,omitempty"`
	Servername         string `json:"servername,omitempty"`
	Port               int    `json:"port,omitempty"`
	Forceupdateip      bool   `json:"forceupdateip,omitempty"`
	Licensemode        string `json:"licensemode,omitempty"`
	Username           string `json:"username,omitempty"`
	Password           string `json:"password,omitempty"`
	Deviceprofilename  string `json:"deviceprofilename,omitempty"`
	Nodeid             int    `json:"nodeid,omitempty"`
	Status             string `json:"status,omitempty"`
	Grace              string `json:"grace,omitempty"`
	Gptimeleft         string `json:"gptimeleft,omitempty"`
	Type               string `json:"type,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nslimitidentifiernslimitsessionsbinding struct {
	Limitidentifier string `json:"limitidentifier,omitempty"`
}

type Nsrpcnode struct {
	Ipaddress          string `json:"ipaddress,omitempty"`
	Password           string `json:"password,omitempty"`
	Srcip              string `json:"srcip,omitempty"`
	Secure             string `json:"secure,omitempty"`
	Validatecert       string `json:"validatecert,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nsspparams struct {
	Basethreshold      int    `json:"basethreshold,omitempty"`
	Throttle           string `json:"throttle,omitempty"`
	Table0             string `json:"table0,omitempty"`
	Builtin            string `json:"builtin,omitempty"`
	Feature            string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nsvariablevalues struct {
	Name               string `json:"name,omitempty"`
	Nodeid             int    `json:"nodeid,omitempty"`
	Variablekey        string `json:"variablekey,omitempty"`
	Variablevalue      string `json:"variablevalue,omitempty"`
	Variabledata       string `json:"variabledata,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nsappflowparam struct {
	Templaterefresh    int    `json:"templaterefresh,omitempty"`
	Udppmtu            int    `json:"udppmtu,omitempty"`
	Httpurl            string `json:"httpurl,omitempty"`
	Httpcookie         string `json:"httpcookie,omitempty"`
	Httpreferer        string `json:"httpreferer,omitempty"`
	Httpmethod         string `json:"httpmethod,omitempty"`
	Httphost           string `json:"httphost,omitempty"`
	Httpuseragent      string `json:"httpuseragent,omitempty"`
	Clienttrafficonly  string `json:"clienttrafficonly,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nsassignment struct {
	Name               string `json:"name,omitempty"`
	Variable           string `json:"variable,omitempty"`
	Set                string `json:"set,omitempty"`
	Add                string `json:"Add,omitempty"`
	Sub                string `json:"sub,omitempty"`
	Append             string `json:"append,omitempty"`
	Clear              bool   `json:"clear,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Newname            string `json:"newname,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Undefhits          string `json:"undefhits,omitempty"`
	Referencecount     string `json:"referencecount,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nshttpparam struct {
	Dropinvalreqs             string `json:"dropinvalreqs,omitempty"`
	Markhttp09inval           string `json:"markhttp09inval,omitempty"`
	Markconnreqinval          string `json:"markconnreqinval,omitempty"`
	Insnssrvrhdr              string `json:"insnssrvrhdr,omitempty"`
	Nssrvrhdr                 string `json:"nssrvrhdr,omitempty"`
	Logerrresp                string `json:"logerrresp,omitempty"`
	Conmultiplex              string `json:"conmultiplex,omitempty"`
	Maxreusepool              int    `json:"maxreusepool"`
	Http2serverside           string `json:"http2serverside,omitempty"`
	Ignoreconnectcodingscheme string `json:"ignoreconnectcodingscheme,omitempty"`
	Builtin                   string `json:"builtin,omitempty"`
	Feature                   string `json:"feature,omitempty"`
	Nextgenapiresource        string `json:"_nextgenapiresource,omitempty"`
}

type Nstimerautoscalepolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Vserver                string `json:"vserver,omitempty"`
	Samplesize             int    `json:"samplesize,omitempty"`
	Threshold              int    `json:"threshold,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Nstrafficdomainvlanbinding struct {
	Vlan int `json:"vlan,omitempty"`
	Td   int `json:"td,omitempty"`
}

type Nslicenseproxyserver struct {
	Serverip           string `json:"serverip,omitempty"`
	Servername         string `json:"servername,omitempty"`
	Port               int    `json:"port,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nsratecontrol struct {
	Tcpthreshold       int    `json:"tcpthreshold"`
	Udpthreshold       int    `json:"udpthreshold"`
	Icmpthreshold      int    `json:"icmpthreshold"`
	Tcprstthreshold    int    `json:"tcprstthreshold"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nsservicefunction struct {
	Servicefunctionname string `json:"servicefunctionname,omitempty"`
	Ingressvlan         int    `json:"ingressvlan,omitempty"`
	Nextgenapiresource  string `json:"_nextgenapiresource,omitempty"`
}

type Nstimer struct {
	Name               string `json:"name,omitempty"`
	Interval           int    `json:"interval,omitempty"`
	Unit               string `json:"unit,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Newname            string `json:"newname,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nsversion struct {
	Installedversion   bool   `json:"installedversion,omitempty"`
	Version            string `json:"version,omitempty"`
	Mode               string `json:"mode,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nscapacity struct {
	Bandwidth          int    `json:"bandwidth,omitempty"`
	Platform           string `json:"platform,omitempty"`
	Vcpu               bool   `json:"vcpu,omitempty"`
	Edition            string `json:"edition,omitempty"`
	Unit               string `json:"unit,omitempty"`
	Username           string `json:"username,omitempty"`
	Password           string `json:"password,omitempty"`
	Nodeid             int    `json:"nodeid,omitempty"`
	Actualbandwidth    string `json:"actualbandwidth,omitempty"`
	Vcpucount          string `json:"vcpucount,omitempty"`
	Maxvcpucount       string `json:"maxvcpucount,omitempty"`
	Maxbandwidth       string `json:"maxbandwidth,omitempty"`
	Minbandwidth       string `json:"minbandwidth,omitempty"`
	Instancecount      string `json:"instancecount,omitempty"`
	Daystoexpiration   string `json:"daystoexpiration,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nslimitidentifierbinding struct {
	Limitidentifier string `json:"limitidentifier,omitempty"`
}

type Nsmode struct {
	Mode                []string `json:"mode,omitempty"`
	Fr                  string   `json:"fr,omitempty"`
	L2                  string   `json:"l2,omitempty"`
	Usip                string   `json:"usip,omitempty"`
	Cka                 string   `json:"cka,omitempty"`
	Tcpb                string   `json:"tcpb,omitempty"`
	Mbf                 string   `json:"mbf,omitempty"`
	Edge                string   `json:"edge,omitempty"`
	Usnip               string   `json:"usnip,omitempty"`
	L3                  string   `json:"l3,omitempty"`
	Pmtud               string   `json:"pmtud,omitempty"`
	Mediaclassification string   `json:"mediaclassification,omitempty"`
	Sradv               string   `json:"sradv,omitempty"`
	Dradv               string   `json:"dradv,omitempty"`
	Iradv               string   `json:"iradv,omitempty"`
	Sradv6              string   `json:"sradv6,omitempty"`
	Dradv6              string   `json:"dradv6,omitempty"`
	Bridgebpdus         string   `json:"bridgebpdus,omitempty"`
	Singleip            string   `json:"single_ip,omitempty"`
	Ulfd                string   `json:"ulfd,omitempty"`
	Nextgenapiresource  string   `json:"_nextgenapiresource,omitempty"`
}

type Nspartitionbridgegroupbinding struct {
	Bridgegroup   int    `json:"bridgegroup,omitempty"`
	Partitionname string `json:"partitionname,omitempty"`
}

type Nspartitionvlanbinding struct {
	Vlan          int    `json:"vlan,omitempty"`
	Partitionname string `json:"partitionname,omitempty"`
}

type Nspartitionmac struct {
	Partitionmac       string `json:"partitionmac,omitempty"`
	Partitionname      string `json:"partitionname,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nsservicepathnsservicefunctionbinding struct {
	Servicefunction string `json:"servicefunction,omitempty"`
	Index           int    `json:"index,omitempty"`
	Servicepathname string `json:"servicepathname,omitempty"`
}

type Nsconfig struct {
	Force                   bool     `json:"force,omitempty"`
	Level                   string   `json:"level,omitempty"`
	Rbaconfig               string   `json:"rbaconfig,omitempty"`
	Ipaddress               string   `json:"ipaddress,omitempty"`
	Netmask                 string   `json:"netmask,omitempty"`
	Nsvlan                  int      `json:"nsvlan,omitempty"`
	Ifnum                   []string `json:"ifnum,omitempty"`
	Tagged                  string   `json:"tagged,omitempty"`
	Httpport                []int    `json:"httpport,omitempty"`
	Maxconn                 int      `json:"maxconn,omitempty"`
	Maxreq                  int      `json:"maxreq,omitempty"`
	Cip                     string   `json:"cip,omitempty"`
	Cipheader               string   `json:"cipheader,omitempty"`
	Cookieversion           string   `json:"cookieversion,omitempty"`
	Securecookie            string   `json:"securecookie,omitempty"`
	Pmtumin                 int      `json:"pmtumin,omitempty"`
	Pmtutimeout             int      `json:"pmtutimeout,omitempty"`
	Ftpportrange            string   `json:"ftpportrange,omitempty"`
	Crportrange             string   `json:"crportrange,omitempty"`
	Timezone                string   `json:"timezone,omitempty"`
	Grantquotamaxclient     int      `json:"grantquotamaxclient,omitempty"`
	Exclusivequotamaxclient int      `json:"exclusivequotamaxclient,omitempty"`
	Grantquotaspillover     int      `json:"grantquotaspillover,omitempty"`
	Exclusivequotaspillover int      `json:"exclusivequotaspillover,omitempty"`
	Securemanagementtraffic string   `json:"securemanagementtraffic,omitempty"`
	Securemanagementtd      int      `json:"securemanagementtd,omitempty"`
	All                     bool     `json:"all,omitempty"`
	Config1                 string   `json:"config1,omitempty"`
	Config2                 string   `json:"config2,omitempty"`
	Outtype                 string   `json:"outtype,omitempty"`
	Template                bool     `json:"template,omitempty"`
	Ignoredevicespecific    bool     `json:"ignoredevicespecific,omitempty"`
	Weakpassword            bool     `json:"weakpassword,omitempty"`
	Changedpassword         bool     `json:"changedpassword,omitempty"`
	Config                  string   `json:"config,omitempty"`
	Configfile              string   `json:"configfile,omitempty"`
	Responsefile            string   `json:"responsefile,omitempty"`
	Async                   bool     `json:"Async,omitempty"`
	Message                 string   `json:"message,omitempty"`
	Mappedip                string   `json:"mappedip,omitempty"`
	Range                   string   `json:"range,omitempty"`
	Svmcmd                  string   `json:"svmcmd,omitempty"`
	Systemtype              string   `json:"systemtype,omitempty"`
	Primaryip               string   `json:"primaryip,omitempty"`
	Primaryip6              string   `json:"primaryip6,omitempty"`
	Flags                   string   `json:"flags,omitempty"`
	Lastconfigchangedtime   string   `json:"lastconfigchangedtime,omitempty"`
	Lastconfigsavetime      string   `json:"lastconfigsavetime,omitempty"`
	Currentsytemtime        string   `json:"currentsytemtime,omitempty"`
	Systemtime              string   `json:"systemtime,omitempty"`
	Configchanged           string   `json:"configchanged,omitempty"`
	Nextgenapiresource      string   `json:"_nextgenapiresource,omitempty"`
	Response                string   `json:"response,omitempty"`
	Id                      string   `json:"id,omitempty"`
}

type Nsencryptionparams struct {
	Method             string `json:"method,omitempty"`
	Keyvalue           string `json:"keyvalue,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nsfeature struct {
	Feature            []string `json:"feature,omitempty"`
	Wl                 string   `json:"wl,omitempty"`
	Sp                 string   `json:"sp,omitempty"`
	Lb                 string   `json:"lb,omitempty"`
	Cs                 string   `json:"cs,omitempty"`
	Cr                 string   `json:"cr,omitempty"`
	Cmp                string   `json:"cmp,omitempty"`
	Ssl                string   `json:"ssl,omitempty"`
	Gslb               string   `json:"gslb,omitempty"`
	Cf                 string   `json:"cf,omitempty"`
	Ic                 string   `json:"ic,omitempty"`
	Sslvpn             string   `json:"sslvpn,omitempty"`
	Aaa                string   `json:"aaa,omitempty"`
	Ospf               string   `json:"ospf,omitempty"`
	Rip                string   `json:"rip,omitempty"`
	Bgp                string   `json:"bgp,omitempty"`
	Rewrite            string   `json:"rewrite,omitempty"`
	Ipv6pt             string   `json:"ipv6pt,omitempty"`
	Appfw              string   `json:"appfw,omitempty"`
	Responder          string   `json:"responder,omitempty"`
	Push               string   `json:"push,omitempty"`
	Appflow            string   `json:"appflow,omitempty"`
	Cloudbridge        string   `json:"cloudbridge,omitempty"`
	Isis               string   `json:"isis,omitempty"`
	Ch                 string   `json:"ch,omitempty"`
	Appqoe             string   `json:"appqoe,omitempty"`
	Contentaccelerator string   `json:"contentaccelerator,omitempty"`
	Feo                string   `json:"feo,omitempty"`
	Lsn                string   `json:"lsn,omitempty"`
	Rdpproxy           string   `json:"rdpproxy,omitempty"`
	Rep                string   `json:"rep,omitempty"`
	Videooptimization  string   `json:"videooptimization,omitempty"`
	Forwardproxy       string   `json:"forwardproxy,omitempty"`
	Sslinterception    string   `json:"sslinterception,omitempty"`
	Adaptivetcp        string   `json:"adaptivetcp,omitempty"`
	Cqa                string   `json:"cqa,omitempty"`
	Ci                 string   `json:"ci,omitempty"`
	Bot                string   `json:"bot,omitempty"`
	Apigateway         string   `json:"apigateway,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
}

type Nsip struct {
	Ipaddress                   string `json:"ipaddress,omitempty"`
	Netmask                     string `json:"netmask,omitempty"`
	Type                        string `json:"type,omitempty"`
	Arp                         string `json:"arp,omitempty"`
	Icmp                        string `json:"icmp,omitempty"`
	Vserver                     string `json:"vserver,omitempty"`
	Telnet                      string `json:"telnet,omitempty"`
	Ftp                         string `json:"ftp,omitempty"`
	Gui                         string `json:"gui,omitempty"`
	Ssh                         string `json:"ssh,omitempty"`
	Snmp                        string `json:"snmp,omitempty"`
	Mgmtaccess                  string `json:"mgmtaccess,omitempty"`
	Restrictaccess              string `json:"restrictaccess,omitempty"`
	Dynamicrouting              string `json:"dynamicrouting,omitempty"`
	Decrementttl                string `json:"decrementttl,omitempty"`
	Ospf                        string `json:"ospf,omitempty"`
	Bgp                         string `json:"bgp,omitempty"`
	Rip                         string `json:"rip,omitempty"`
	Hostroute                   string `json:"hostroute,omitempty"`
	Advertiseondefaultpartition string `json:"advertiseondefaultpartition,omitempty"`
	Networkroute                string `json:"networkroute,omitempty"`
	Tag                         int    `json:"tag,omitempty"`
	Hostrtgw                    string `json:"hostrtgw,omitempty"`
	Metric                      int    `json:"metric,omitempty"`
	Vserverrhilevel             string `json:"vserverrhilevel,omitempty"`
	Ospflsatype                 string `json:"ospflsatype,omitempty"`
	Ospfarea                    int    `json:"ospfarea,omitempty"`
	State                       string `json:"state,omitempty"`
	Vrid                        int    `json:"vrid,omitempty"`
	Icmpresponse                string `json:"icmpresponse,omitempty"`
	Ownernode                   int    `json:"ownernode,omitempty"`
	Arpresponse                 string `json:"arpresponse,omitempty"`
	Ownerdownresponse           string `json:"ownerdownresponse,omitempty"`
	Td                          int    `json:"td,omitempty"`
	Arpowner                    int    `json:"arpowner,omitempty"`
	Mptcpadvertise              string `json:"mptcpadvertise,omitempty"`
	Flags                       string `json:"flags,omitempty"`
	Hostrtgwact                 string `json:"hostrtgwact,omitempty"`
	Ospfareaval                 string `json:"ospfareaval,omitempty"`
	Viprtadv2bsd                string `json:"viprtadv2bsd,omitempty"`
	Vipvsercount                string `json:"vipvsercount,omitempty"`
	Vipvserdowncount            string `json:"vipvserdowncount,omitempty"`
	Vipvsrvrrhiactivecount      string `json:"vipvsrvrrhiactivecount,omitempty"`
	Vipvsrvrrhiactiveupcount    string `json:"vipvsrvrrhiactiveupcount,omitempty"`
	Freeports                   string `json:"freeports,omitempty"`
	Iptype                      string `json:"iptype,omitempty"`
	Operationalarpowner         string `json:"operationalarpowner,omitempty"`
	Nextgenapiresource          string `json:"_nextgenapiresource,omitempty"`
}

type Nslimitselector struct {
	Selectorname       string   `json:"selectorname,omitempty"`
	Rule               []string `json:"rule,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
}

type Nstimeout struct {
	Zombie             int    `json:"zombie,omitempty"`
	Client             int    `json:"client"`
	Server             int    `json:"server"`
	Httpclient         int    `json:"httpclient"`
	Httpserver         int    `json:"httpserver"`
	Tcpclient          int    `json:"tcpclient"`
	Tcpserver          int    `json:"tcpserver"`
	Anyclient          int    `json:"anyclient"`
	Anyserver          int    `json:"anyserver"`
	Anytcpclient       int    `json:"anytcpclient"`
	Anytcpserver       int    `json:"anytcpserver"`
	Halfclose          int    `json:"halfclose,omitempty"`
	Nontcpzombie       int    `json:"nontcpzombie,omitempty"`
	Reducedfintimeout  int    `json:"reducedfintimeout,omitempty"`
	Reducedrsttimeout  int    `json:"reducedrsttimeout"`
	Newconnidletimeout int    `json:"newconnidletimeout,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nstimerbinding struct {
	Name string `json:"name,omitempty"`
}

type Nsaptlicense struct {
	Serialno           string `json:"serialno,omitempty"`
	Useproxy           string `json:"useproxy,omitempty"`
	Id                 string `json:"id,omitempty"`
	Sessionid          string `json:"sessionid,omitempty"`
	Bindtype           string `json:"bindtype,omitempty"`
	Countavailable     string `json:"countavailable,omitempty"`
	Licensedir         string `json:"licensedir,omitempty"`
	Response           string `json:"response,omitempty"`
	Counttotal         string `json:"counttotal,omitempty"`
	Name               string `json:"name,omitempty"`
	Relevance          string `json:"relevance,omitempty"`
	Datepurchased      string `json:"datepurchased,omitempty"`
	Datesa             string `json:"datesa,omitempty"`
	Dateexp            string `json:"dateexp,omitempty"`
	Features           string `json:"features,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nschannelparam struct {
	Vfautorecover      string `json:"vfautorecover,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nsip6 struct {
	Ipv6address                 string `json:"ipv6address,omitempty"`
	Scope                       string `json:"scope,omitempty"`
	Type                        string `json:"type,omitempty"`
	Vlan                        int    `json:"vlan,omitempty"`
	Nd                          string `json:"nd,omitempty"`
	Icmp                        string `json:"icmp,omitempty"`
	Vserver                     string `json:"vserver,omitempty"`
	Telnet                      string `json:"telnet,omitempty"`
	Ftp                         string `json:"ftp,omitempty"`
	Gui                         string `json:"gui,omitempty"`
	Ssh                         string `json:"ssh,omitempty"`
	Snmp                        string `json:"snmp,omitempty"`
	Mgmtaccess                  string `json:"mgmtaccess,omitempty"`
	Restrictaccess              string `json:"restrictaccess,omitempty"`
	Dynamicrouting              string `json:"dynamicrouting,omitempty"`
	Decrementhoplimit           string `json:"decrementhoplimit,omitempty"`
	Hostroute                   string `json:"hostroute,omitempty"`
	Advertiseondefaultpartition string `json:"advertiseondefaultpartition,omitempty"`
	Networkroute                string `json:"networkroute,omitempty"`
	Tag                         int    `json:"tag,omitempty"`
	Ip6hostrtgw                 string `json:"ip6hostrtgw,omitempty"`
	Metric                      int    `json:"metric,omitempty"`
	Vserverrhilevel             string `json:"vserverrhilevel,omitempty"`
	Ospf6lsatype                string `json:"ospf6lsatype,omitempty"`
	Ospfarea                    int    `json:"ospfarea,omitempty"`
	State                       string `json:"state,omitempty"`
	Map                         string `json:"map,omitempty"`
	Vrid6                       int    `json:"vrid6,omitempty"`
	Ownernode                   int    `json:"ownernode,omitempty"`
	Ownerdownresponse           string `json:"ownerdownresponse,omitempty"`
	Td                          int    `json:"td,omitempty"`
	Ndowner                     int    `json:"ndowner,omitempty"`
	Mptcpadvertise              string `json:"mptcpadvertise,omitempty"`
	Icmpresponse                string `json:"icmpresponse,omitempty"`
	Iptype                      string `json:"iptype,omitempty"`
	Curstate                    string `json:"curstate,omitempty"`
	Viprtadv2bsd                string `json:"viprtadv2bsd,omitempty"`
	Vipvsercount                string `json:"vipvsercount,omitempty"`
	Vipvserdowncount            string `json:"vipvserdowncount,omitempty"`
	Systemtype                  string `json:"systemtype,omitempty"`
	Operationalndowner          string `json:"operationalndowner,omitempty"`
	Nextgenapiresource          string `json:"_nextgenapiresource,omitempty"`
}

type Nsjob struct {
	Id                 int    `json:"id,omitempty"`
	Name               string `json:"name,omitempty"`
	Status             string `json:"status,omitempty"`
	Progress           string `json:"progress,omitempty"`
	Timeelapsed        string `json:"timeelapsed,omitempty"`
	Errorcode          string `json:"errorcode,omitempty"`
	Message            string `json:"message,omitempty"`
	Response           string `json:"response,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nslicense struct {
	Wl                      string `json:"wl,omitempty"`
	Sp                      string `json:"sp,omitempty"`
	Lb                      string `json:"lb,omitempty"`
	Cs                      string `json:"cs,omitempty"`
	Cr                      string `json:"cr,omitempty"`
	Cmp                     string `json:"cmp,omitempty"`
	Delta                   string `json:"delta,omitempty"`
	Ssl                     string `json:"ssl,omitempty"`
	Gslb                    string `json:"gslb,omitempty"`
	Gslbp                   string `json:"gslbp,omitempty"`
	Routing                 string `json:"routing,omitempty"`
	Cf                      string `json:"cf,omitempty"`
	Contentaccelerator      string `json:"contentaccelerator,omitempty"`
	Ic                      string `json:"ic,omitempty"`
	Sslvpn                  string `json:"sslvpn,omitempty"`
	Fsslvpnusers            string `json:"f_sslvpn_users,omitempty"`
	Ficausers               string `json:"f_ica_users,omitempty"`
	Aaa                     string `json:"aaa,omitempty"`
	Ospf                    string `json:"ospf,omitempty"`
	Rip                     string `json:"rip,omitempty"`
	Bgp                     string `json:"bgp,omitempty"`
	Rewrite                 string `json:"rewrite,omitempty"`
	Ipv6pt                  string `json:"ipv6pt,omitempty"`
	Appfw                   string `json:"appfw,omitempty"`
	Responder               string `json:"responder,omitempty"`
	Agee                    string `json:"agee,omitempty"`
	Nsxn                    string `json:"nsxn,omitempty"`
	Modelid                 string `json:"modelid,omitempty"`
	Push                    string `json:"push,omitempty"`
	Appflow                 string `json:"appflow,omitempty"`
	Cloudbridge             string `json:"cloudbridge,omitempty"`
	Cloudbridgeappliance    string `json:"cloudbridgeappliance,omitempty"`
	Cloudextenderappliance  string `json:"cloudextenderappliance,omitempty"`
	Isis                    string `json:"isis,omitempty"`
	Cluster                 string `json:"cluster,omitempty"`
	Ch                      string `json:"ch,omitempty"`
	Appqoe                  string `json:"appqoe,omitempty"`
	Appflowica              string `json:"appflowica,omitempty"`
	Isstandardlic           string `json:"isstandardlic,omitempty"`
	Isenterpriselic         string `json:"isenterpriselic,omitempty"`
	Isplatinumlic           string `json:"isplatinumlic,omitempty"`
	Issgwylic               string `json:"issgwylic,omitempty"`
	Isswglic                string `json:"isswglic,omitempty"`
	Feo                     string `json:"feo,omitempty"`
	Lsn                     string `json:"lsn,omitempty"`
	Licensingmode           string `json:"licensingmode,omitempty"`
	Cloudsubscriptionimage  string `json:"cloudsubscriptionimage,omitempty"`
	Daystoexpiration        string `json:"daystoexpiration,omitempty"`
	Daystolasenforcement    string `json:"daystolasenforcement,omitempty"`
	Rdpproxy                string `json:"rdpproxy,omitempty"`
	Rep                     string `json:"rep,omitempty"`
	Urlfiltering            string `json:"urlfiltering,omitempty"`
	Videooptimization       string `json:"videooptimization,omitempty"`
	Forwardproxy            string `json:"forwardproxy,omitempty"`
	Sslinterception         string `json:"sslinterception,omitempty"`
	Remotecontentinspection string `json:"remotecontentinspection,omitempty"`
	Adaptivetcp             string `json:"adaptivetcp,omitempty"`
	Cqa                     string `json:"cqa,omitempty"`
	Bot                     string `json:"bot,omitempty"`
	Apigateway              string `json:"apigateway,omitempty"`
	Nextgenapiresource      string `json:"_nextgenapiresource,omitempty"`
}

type Nslicenseactivationdata struct {
	Filename           string `json:"filename,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nsmigration struct {
	Dumpsession                     string `json:"dumpsession,omitempty"`
	Migrationstatus                 string `json:"migrationstatus,omitempty"`
	Migrationstarttime              string `json:"migrationstarttime,omitempty"`
	Migrationendtime                string `json:"migrationendtime,omitempty"`
	Migrationrollbackstarttime      string `json:"migrationrollbackstarttime,omitempty"`
	Srcip                           string `json:"srcip,omitempty"`
	Srcport                         string `json:"srcport,omitempty"`
	Destip                          string `json:"destip,omitempty"`
	Destport                        string `json:"destport,omitempty"`
	Timeout                         string `json:"timeout,omitempty"`
	Migdfdsessionsallocated         string `json:"migdfdsessionsallocated,omitempty"`
	Migdfdsessionsactive            string `json:"migdfdsessionsactive,omitempty"`
	Migl4sessionsallocated          string `json:"migl4sessionsallocated,omitempty"`
	Migl4sessionsactive             string `json:"migl4sessionsactive,omitempty"`
	Migdfdsessionsallocatedrollback string `json:"migdfdsessionsallocatedrollback,omitempty"`
	Migdfdsessionsactiverollback    string `json:"migdfdsessionsactiverollback,omitempty"`
	Migl4sessionsallocatedrollback  string `json:"migl4sessionsallocatedrollback,omitempty"`
	Migl4sessionsactiverollback     string `json:"migl4sessionsactiverollback,omitempty"`
	Mighastateflag                  string `json:"mighastateflag,omitempty"`
	Nextgenapiresource              string `json:"_nextgenapiresource,omitempty"`
}

type Nsservicepath struct {
	Servicepathname    string `json:"servicepathname,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nsacl6 struct {
	Acl6name           string `json:"acl6name,omitempty"`
	Acl6action         string `json:"acl6action,omitempty"`
	Td                 int    `json:"td,omitempty"`
	Srcipv6            bool   `json:"srcipv6,omitempty"`
	Srcipop            string `json:"srcipop,omitempty"`
	Srcipv6val         string `json:"srcipv6val,omitempty"`
	Srcport            bool   `json:"srcport,omitempty"`
	Srcportop          string `json:"srcportop,omitempty"`
	Srcportval         string `json:"srcportval,omitempty"`
	Destipv6           bool   `json:"destipv6,omitempty"`
	Destipop           string `json:"destipop,omitempty"`
	Destipv6val        string `json:"destipv6val,omitempty"`
	Destport           bool   `json:"destport,omitempty"`
	Destportop         string `json:"destportop,omitempty"`
	Destportval        string `json:"destportval,omitempty"`
	Ttl                int    `json:"ttl,omitempty"`
	Srcmac             string `json:"srcmac,omitempty"`
	Srcmacmask         string `json:"srcmacmask,omitempty"`
	Protocol           string `json:"protocol,omitempty"`
	Protocolnumber     int    `json:"protocolnumber,omitempty"`
	Vlan               int    `json:"vlan,omitempty"`
	Vxlan              int    `json:"vxlan,omitempty"`
	Interface          string `json:"Interface,omitempty"`
	Established        bool   `json:"established,omitempty"`
	Icmptype           int    `json:"icmptype,omitempty"`
	Icmpcode           int    `json:"icmpcode,omitempty"`
	Priority           int    `json:"priority,omitempty"`
	State              string `json:"state,omitempty"`
	Type               string `json:"type,omitempty"`
	Dfdhash            string `json:"dfdhash,omitempty"`
	Dfdprefix          int    `json:"dfdprefix,omitempty"`
	Nodeid             int    `json:"nodeid,omitempty"`
	Stateful           string `json:"stateful,omitempty"`
	Logstate           string `json:"logstate,omitempty"`
	Ratelimit          int    `json:"ratelimit,omitempty"`
	Aclaction          string `json:"aclaction,omitempty"`
	Newname            string `json:"newname,omitempty"`
	Kernelstate        string `json:"kernelstate,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Aclassociate       string `json:"aclassociate,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nsappflowcollector struct {
	Name               string `json:"name,omitempty"`
	Ipaddress          string `json:"ipaddress,omitempty"`
	Port               int    `json:"port,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nsconfigview struct {
	State              string `json:"state,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nslaslicense struct {
	Filename           string `json:"filename,omitempty"`
	Filelocation       string `json:"filelocation,omitempty"`
	Fixedbandwidth     bool   `json:"fixedbandwidth,omitempty"`
	Status             string `json:"status,omitempty"`
	Daystoexpiration   string `json:"daystoexpiration,omitempty"`
	Renewalprev        string `json:"renewalprev,omitempty"`
	Renewalnext        string `json:"renewalnext,omitempty"`
	Renewalprevdate    string `json:"renewalprevdate,omitempty"`
	Renewalnextdate    string `json:"renewalnextdate,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nslicenseserverpool struct {
	Getalllicenses               bool   `json:"getalllicenses,omitempty"`
	Instancetotal                string `json:"instancetotal,omitempty"`
	Instanceavailable            string `json:"instanceavailable,omitempty"`
	Standardbandwidthtotal       string `json:"standardbandwidthtotal,omitempty"`
	Standardbandwidthavailable   string `json:"standardbandwidthavailable,omitempty"`
	Enterprisebandwidthtotal     string `json:"enterprisebandwidthtotal,omitempty"`
	Enterprisebandwidthavailable string `json:"enterprisebandwidthavailable,omitempty"`
	Platinumbandwidthtotal       string `json:"platinumbandwidthtotal,omitempty"`
	Platinumbandwidthavailable   string `json:"platinumbandwidthavailable,omitempty"`
	Standardcputotal             string `json:"standardcputotal,omitempty"`
	Standardcpuavailable         string `json:"standardcpuavailable,omitempty"`
	Enterprisecputotal           string `json:"enterprisecputotal,omitempty"`
	Enterprisecpuavailable       string `json:"enterprisecpuavailable,omitempty"`
	Platinumcputotal             string `json:"platinumcputotal,omitempty"`
	Platinumcpuavailable         string `json:"platinumcpuavailable,omitempty"`
	Cpxinstancetotal             string `json:"cpxinstancetotal,omitempty"`
	Cpxinstanceavailable         string `json:"cpxinstanceavailable,omitempty"`
	Vpx1stotal                   string `json:"vpx1stotal,omitempty"`
	Vpx1savailable               string `json:"vpx1savailable,omitempty"`
	Vpx1ptotal                   string `json:"vpx1ptotal,omitempty"`
	Vpx1pavailable               string `json:"vpx1pavailable,omitempty"`
	Vpx5stotal                   string `json:"vpx5stotal,omitempty"`
	Vpx5savailable               string `json:"vpx5savailable,omitempty"`
	Vpx5ptotal                   string `json:"vpx5ptotal,omitempty"`
	Vpx5pavailable               string `json:"vpx5pavailable,omitempty"`
	Vpx10stotal                  string `json:"vpx10stotal,omitempty"`
	Vpx10savailable              string `json:"vpx10savailable,omitempty"`
	Vpx10etotal                  string `json:"vpx10etotal,omitempty"`
	Vpx10eavailable              string `json:"vpx10eavailable,omitempty"`
	Vpx10ptotal                  string `json:"vpx10ptotal,omitempty"`
	Vpx10pavailable              string `json:"vpx10pavailable,omitempty"`
	Vpx25stotal                  string `json:"vpx25stotal,omitempty"`
	Vpx25savailable              string `json:"vpx25savailable,omitempty"`
	Vpx25etotal                  string `json:"vpx25etotal,omitempty"`
	Vpx25eavailable              string `json:"vpx25eavailable,omitempty"`
	Vpx25ptotal                  string `json:"vpx25ptotal,omitempty"`
	Vpx25pavailable              string `json:"vpx25pavailable,omitempty"`
	Vpx50stotal                  string `json:"vpx50stotal,omitempty"`
	Vpx50savailable              string `json:"vpx50savailable,omitempty"`
	Vpx50etotal                  string `json:"vpx50etotal,omitempty"`
	Vpx50eavailable              string `json:"vpx50eavailable,omitempty"`
	Vpx50ptotal                  string `json:"vpx50ptotal,omitempty"`
	Vpx50pavailable              string `json:"vpx50pavailable,omitempty"`
	Vpx100stotal                 string `json:"vpx100stotal,omitempty"`
	Vpx100savailable             string `json:"vpx100savailable,omitempty"`
	Vpx100etotal                 string `json:"vpx100etotal,omitempty"`
	Vpx100eavailable             string `json:"vpx100eavailable,omitempty"`
	Vpx100ptotal                 string `json:"vpx100ptotal,omitempty"`
	Vpx100pavailable             string `json:"vpx100pavailable,omitempty"`
	Vpx200stotal                 string `json:"vpx200stotal,omitempty"`
	Vpx200savailable             string `json:"vpx200savailable,omitempty"`
	Vpx200etotal                 string `json:"vpx200etotal,omitempty"`
	Vpx200eavailable             string `json:"vpx200eavailable,omitempty"`
	Vpx200ptotal                 string `json:"vpx200ptotal,omitempty"`
	Vpx200pavailable             string `json:"vpx200pavailable,omitempty"`
	Vpx500stotal                 string `json:"vpx500stotal,omitempty"`
	Vpx500savailable             string `json:"vpx500savailable,omitempty"`
	Vpx500etotal                 string `json:"vpx500etotal,omitempty"`
	Vpx500eavailable             string `json:"vpx500eavailable,omitempty"`
	Vpx500ptotal                 string `json:"vpx500ptotal,omitempty"`
	Vpx500pavailable             string `json:"vpx500pavailable,omitempty"`
	Vpx1000stotal                string `json:"vpx1000stotal,omitempty"`
	Vpx1000savailable            string `json:"vpx1000savailable,omitempty"`
	Vpx1000etotal                string `json:"vpx1000etotal,omitempty"`
	Vpx1000eavailable            string `json:"vpx1000eavailable,omitempty"`
	Vpx1000ptotal                string `json:"vpx1000ptotal,omitempty"`
	Vpx1000pavailable            string `json:"vpx1000pavailable,omitempty"`
	Vpx2000ptotal                string `json:"vpx2000ptotal,omitempty"`
	Vpx2000pavailable            string `json:"vpx2000pavailable,omitempty"`
	Vpx3000stotal                string `json:"vpx3000stotal,omitempty"`
	Vpx3000savailable            string `json:"vpx3000savailable,omitempty"`
	Vpx3000etotal                string `json:"vpx3000etotal,omitempty"`
	Vpx3000eavailable            string `json:"vpx3000eavailable,omitempty"`
	Vpx3000ptotal                string `json:"vpx3000ptotal,omitempty"`
	Vpx3000pavailable            string `json:"vpx3000pavailable,omitempty"`
	Vpx4000ptotal                string `json:"vpx4000ptotal,omitempty"`
	Vpx4000pavailable            string `json:"vpx4000pavailable,omitempty"`
	Vpx5000stotal                string `json:"vpx5000stotal,omitempty"`
	Vpx5000savailable            string `json:"vpx5000savailable,omitempty"`
	Vpx5000etotal                string `json:"vpx5000etotal,omitempty"`
	Vpx5000eavailable            string `json:"vpx5000eavailable,omitempty"`
	Vpx5000ptotal                string `json:"vpx5000ptotal,omitempty"`
	Vpx5000pavailable            string `json:"vpx5000pavailable,omitempty"`
	Vpx8000stotal                string `json:"vpx8000stotal,omitempty"`
	Vpx8000savailable            string `json:"vpx8000savailable,omitempty"`
	Vpx8000etotal                string `json:"vpx8000etotal,omitempty"`
	Vpx8000eavailable            string `json:"vpx8000eavailable,omitempty"`
	Vpx8000ptotal                string `json:"vpx8000ptotal,omitempty"`
	Vpx8000pavailable            string `json:"vpx8000pavailable,omitempty"`
	Vpx10000stotal               string `json:"vpx10000stotal,omitempty"`
	Vpx10000savailable           string `json:"vpx10000savailable,omitempty"`
	Vpx10000etotal               string `json:"vpx10000etotal,omitempty"`
	Vpx10000eavailable           string `json:"vpx10000eavailable,omitempty"`
	Vpx10000ptotal               string `json:"vpx10000ptotal,omitempty"`
	Vpx10000pavailable           string `json:"vpx10000pavailable,omitempty"`
	Vpx15000stotal               string `json:"vpx15000stotal,omitempty"`
	Vpx15000savailable           string `json:"vpx15000savailable,omitempty"`
	Vpx15000etotal               string `json:"vpx15000etotal,omitempty"`
	Vpx15000eavailable           string `json:"vpx15000eavailable,omitempty"`
	Vpx15000ptotal               string `json:"vpx15000ptotal,omitempty"`
	Vpx15000pavailable           string `json:"vpx15000pavailable,omitempty"`
	Vpx25000stotal               string `json:"vpx25000stotal,omitempty"`
	Vpx25000savailable           string `json:"vpx25000savailable,omitempty"`
	Vpx25000etotal               string `json:"vpx25000etotal,omitempty"`
	Vpx25000eavailable           string `json:"vpx25000eavailable,omitempty"`
	Vpx25000ptotal               string `json:"vpx25000ptotal,omitempty"`
	Vpx25000pavailable           string `json:"vpx25000pavailable,omitempty"`
	Vpx40000stotal               string `json:"vpx40000stotal,omitempty"`
	Vpx40000savailable           string `json:"vpx40000savailable,omitempty"`
	Vpx40000etotal               string `json:"vpx40000etotal,omitempty"`
	Vpx40000eavailable           string `json:"vpx40000eavailable,omitempty"`
	Vpx40000ptotal               string `json:"vpx40000ptotal,omitempty"`
	Vpx40000pavailable           string `json:"vpx40000pavailable,omitempty"`
	Vpx100000stotal              string `json:"vpx100000stotal,omitempty"`
	Vpx100000savailable          string `json:"vpx100000savailable,omitempty"`
	Vpx100000etotal              string `json:"vpx100000etotal,omitempty"`
	Vpx100000eavailable          string `json:"vpx100000eavailable,omitempty"`
	Vpx100000ptotal              string `json:"vpx100000ptotal,omitempty"`
	Vpx100000pavailable          string `json:"vpx100000pavailable,omitempty"`
	Licensemode                  string `json:"licensemode,omitempty"`
	Nextgenapiresource           string `json:"_nextgenapiresource,omitempty"`
}

type Nslimitidentifierlimitsessionsbinding struct {
	Limitidentifier string `json:"limitidentifier,omitempty"`
}

type Nspartition struct {
	Partitionname      string `json:"partitionname,omitempty"`
	Maxbandwidth       int    `json:"maxbandwidth"`
	Minbandwidth       int    `json:"minbandwidth"`
	Maxconn            int    `json:"maxconn"`
	Maxmemlimit        int    `json:"maxmemlimit"`
	Partitionmac       string `json:"partitionmac,omitempty"`
	Force              bool   `json:"force,omitempty"`
	Save               bool   `json:"save,omitempty"`
	Partitionid        string `json:"partitionid,omitempty"`
	Partitiontype      string `json:"partitiontype,omitempty"`
	Pmacinternal       string `json:"pmacinternal,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nsrunningconfig struct {
	Withdefaults       bool   `json:"withdefaults,omitempty"`
	Response           string `json:"response,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nshardware struct {
	Hwdescription      string `json:"hwdescription,omitempty"`
	Sysid              string `json:"sysid,omitempty"`
	Manufactureday     string `json:"manufactureday,omitempty"`
	Manufacturemonth   string `json:"manufacturemonth,omitempty"`
	Manufactureyear    string `json:"manufactureyear,omitempty"`
	Cpufrequncy        string `json:"cpufrequncy,omitempty"`
	Hostid             string `json:"hostid,omitempty"`
	Host               string `json:"host,omitempty"`
	Serialno           string `json:"serialno,omitempty"`
	Encodedserialno    string `json:"encodedserialno,omitempty"`
	Netscaleruuid      string `json:"netscaleruuid,omitempty"`
	Bmcrevision        string `json:"bmcrevision,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nshttpprofile struct {
	Name                             string `json:"name,omitempty"`
	Dropinvalreqs                    string `json:"dropinvalreqs,omitempty"`
	Markhttp09inval                  string `json:"markhttp09inval,omitempty"`
	Markconnreqinval                 string `json:"markconnreqinval,omitempty"`
	Marktracereqinval                string `json:"marktracereqinval,omitempty"`
	Markrfc7230noncompliantinval     string `json:"markrfc7230noncompliantinval,omitempty"`
	Markhttpheaderextrawserror       string `json:"markhttpheaderextrawserror,omitempty"`
	Cmponpush                        string `json:"cmponpush,omitempty"`
	Conmultiplex                     string `json:"conmultiplex,omitempty"`
	Maxreusepool                     int    `json:"maxreusepool,omitempty"`
	Dropextracrlf                    string `json:"dropextracrlf,omitempty"`
	Incomphdrdelay                   int    `json:"incomphdrdelay,omitempty"`
	Websocket                        string `json:"websocket,omitempty"`
	Rtsptunnel                       string `json:"rtsptunnel,omitempty"`
	Reqtimeout                       int    `json:"reqtimeout,omitempty"`
	Adpttimeout                      string `json:"adpttimeout,omitempty"`
	Reqtimeoutaction                 string `json:"reqtimeoutaction,omitempty"`
	Dropextradata                    string `json:"dropextradata,omitempty"`
	Weblog                           string `json:"weblog,omitempty"`
	Clientiphdrexpr                  string `json:"clientiphdrexpr,omitempty"`
	Maxreq                           int    `json:"maxreq,omitempty"`
	Persistentetag                   string `json:"persistentetag,omitempty"`
	Http2                            string `json:"http2,omitempty"`
	Http2direct                      string `json:"http2direct,omitempty"`
	Http2strictcipher                string `json:"http2strictcipher,omitempty"`
	Http2altsvcframe                 string `json:"http2altsvcframe,omitempty"`
	Altsvc                           string `json:"altsvc,omitempty"`
	Altsvcvalue                      string `json:"altsvcvalue,omitempty"`
	Reusepooltimeout                 int    `json:"reusepooltimeout,omitempty"`
	Maxheaderlen                     int    `json:"maxheaderlen,omitempty"`
	Maxheaderfieldlen                int    `json:"maxheaderfieldlen,omitempty"`
	Minreusepool                     int    `json:"minreusepool,omitempty"`
	Http2maxheaderlistsize           int    `json:"http2maxheaderlistsize,omitempty"`
	Http2maxframesize                int    `json:"http2maxframesize,omitempty"`
	Http2maxconcurrentstreams        int    `json:"http2maxconcurrentstreams,omitempty"`
	Http2initialconnwindowsize       int    `json:"http2initialconnwindowsize,omitempty"`
	Http2initialwindowsize           int    `json:"http2initialwindowsize,omitempty"`
	Http2headertablesize             int    `json:"http2headertablesize,omitempty"`
	Http2minseverconn                int    `json:"http2minseverconn,omitempty"`
	Http2maxpingframespermin         int    `json:"http2maxpingframespermin,omitempty"`
	Http2maxsettingsframespermin     int    `json:"http2maxsettingsframespermin,omitempty"`
	Http2maxresetframespermin        int    `json:"http2maxresetframespermin,omitempty"`
	Http2maxemptyframespermin        int    `json:"http2maxemptyframespermin,omitempty"`
	Http2maxrxresetframespermin      int    `json:"http2maxrxresetframespermin,omitempty"`
	Grpcholdlimit                    int    `json:"grpcholdlimit,omitempty"`
	Grpcholdtimeout                  int    `json:"grpcholdtimeout,omitempty"`
	Grpclengthdelimitation           string `json:"grpclengthdelimitation,omitempty"`
	Apdexcltresptimethreshold        int    `json:"apdexcltresptimethreshold,omitempty"`
	Http3                            string `json:"http3,omitempty"`
	Http3maxheaderfieldsectionsize   int    `json:"http3maxheaderfieldsectionsize,omitempty"`
	Http3maxheadertablesize          int    `json:"http3maxheadertablesize,omitempty"`
	Http3maxheaderblockedstreams     int    `json:"http3maxheaderblockedstreams,omitempty"`
	Http3webtransport                string `json:"http3webtransport,omitempty"`
	Http3minseverconn                int    `json:"http3minseverconn,omitempty"`
	Httppipelinebuffsize             int    `json:"httppipelinebuffsize,omitempty"`
	Allowonlywordcharactersandhyphen string `json:"allowonlywordcharactersandhyphen,omitempty"`
	Hostheadervalidation             string `json:"hostheadervalidation,omitempty"`
	Maxduplicateheaderfields         int    `json:"maxduplicateheaderfields,omitempty"`
	Passprotocolupgrade              string `json:"passprotocolupgrade,omitempty"`
	Http2extendedconnect             string `json:"http2extendedconnect,omitempty"`
	Refcnt                           string `json:"refcnt,omitempty"`
	Builtin                          string `json:"builtin,omitempty"`
	Apdexsvrresptimethreshold        string `json:"apdexsvrresptimethreshold,omitempty"`
	Dropinvalreqswarning             string `json:"dropinvalreqswarning,omitempty"`
	Feature                          string `json:"feature,omitempty"`
	Nextgenapiresource               string `json:"_nextgenapiresource,omitempty"`
}

type Nslimitidentifier struct {
	Limitidentifier          string `json:"limitidentifier,omitempty"`
	Threshold                int    `json:"threshold,omitempty"`
	Timeslice                int    `json:"timeslice,omitempty"`
	Mode                     string `json:"mode,omitempty"`
	Limittype                string `json:"limittype,omitempty"`
	Selectorname             string `json:"selectorname,omitempty"`
	Maxbandwidth             int    `json:"maxbandwidth,omitempty"`
	Trapsintimeslice         int    `json:"trapsintimeslice,omitempty"`
	Ngname                   string `json:"ngname,omitempty"`
	Hits                     string `json:"hits,omitempty"`
	Drop                     string `json:"drop,omitempty"`
	Rule                     string `json:"rule,omitempty"`
	Time                     string `json:"time,omitempty"`
	Total                    string `json:"total,omitempty"`
	Trapscomputedintimeslice string `json:"trapscomputedintimeslice,omitempty"`
	Computedtraptimeslice    string `json:"computedtraptimeslice,omitempty"`
	Referencecount           string `json:"referencecount,omitempty"`
	Nextgenapiresource       string `json:"_nextgenapiresource,omitempty"`
}

type Nspbr6 struct {
	Name               string `json:"name,omitempty"`
	Td                 int    `json:"td,omitempty"`
	Action             string `json:"action,omitempty"`
	Srcipv6            bool   `json:"srcipv6,omitempty"`
	Srcipop            string `json:"srcipop,omitempty"`
	Srcipv6val         string `json:"srcipv6val,omitempty"`
	Srcport            bool   `json:"srcport,omitempty"`
	Srcportop          string `json:"srcportop,omitempty"`
	Srcportval         string `json:"srcportval,omitempty"`
	Destipv6           bool   `json:"destipv6,omitempty"`
	Destipop           string `json:"destipop,omitempty"`
	Destipv6val        string `json:"destipv6val,omitempty"`
	Destport           bool   `json:"destport,omitempty"`
	Destportop         string `json:"destportop,omitempty"`
	Destportval        string `json:"destportval,omitempty"`
	Srcmac             string `json:"srcmac,omitempty"`
	Srcmacmask         string `json:"srcmacmask,omitempty"`
	Protocol           string `json:"protocol,omitempty"`
	Protocolnumber     int    `json:"protocolnumber,omitempty"`
	Vlan               int    `json:"vlan,omitempty"`
	Vxlan              int    `json:"vxlan,omitempty"`
	Interface          string `json:"Interface,omitempty"`
	Priority           int    `json:"priority,omitempty"`
	State              string `json:"state,omitempty"`
	Msr                string `json:"msr,omitempty"`
	Monitor            string `json:"monitor,omitempty"`
	Nexthop            bool   `json:"nexthop,omitempty"`
	Nexthopval         string `json:"nexthopval,omitempty"`
	Iptunnel           string `json:"iptunnel,omitempty"`
	Vxlanvlanmap       string `json:"vxlanvlanmap,omitempty"`
	Nexthopvlan        int    `json:"nexthopvlan,omitempty"`
	Ownergroup         string `json:"ownergroup,omitempty"`
	Detail             bool   `json:"detail,omitempty"`
	Kernelstate        string `json:"kernelstate,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Curstate           string `json:"curstate,omitempty"`
	Totalprobes        string `json:"totalprobes,omitempty"`
	Totalfailedprobes  string `json:"totalfailedprobes,omitempty"`
	Failedprobes       string `json:"failedprobes,omitempty"`
	Monstatcode        string `json:"monstatcode,omitempty"`
	Monstatparam1      string `json:"monstatparam1,omitempty"`
	Monstatparam2      string `json:"monstatparam2,omitempty"`
	Monstatparam3      string `json:"monstatparam3,omitempty"`
	Data               string `json:"data,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nssimpleacl6 struct {
	Aclname            string `json:"aclname,omitempty"`
	Td                 int    `json:"td,omitempty"`
	Aclaction          string `json:"aclaction,omitempty"`
	Srcipv6            string `json:"srcipv6,omitempty"`
	Destport           int    `json:"destport,omitempty"`
	Protocol           string `json:"protocol,omitempty"`
	Ttl                int    `json:"ttl,omitempty"`
	Estsessions        bool   `json:"estsessions,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nssourceroutecachetable struct {
	Sourceip           string `json:"sourceip,omitempty"`
	Sourcemac          string `json:"sourcemac,omitempty"`
	Vlan               string `json:"vlan,omitempty"`
	Interface          string `json:"Interface,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nstimezone struct {
	Value              string `json:"value,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nsconnectiontable struct {
	Filterexpression       string   `json:"filterexpression,omitempty"`
	Link                   bool     `json:"link,omitempty"`
	Filtername             bool     `json:"filtername,omitempty"`
	Detail                 []string `json:"detail,omitempty"`
	Listen                 bool     `json:"listen,omitempty"`
	Nodeid                 int      `json:"nodeid,omitempty"`
	Sourceip               string   `json:"sourceip,omitempty"`
	Sourceport             string   `json:"sourceport,omitempty"`
	Destip                 string   `json:"destip,omitempty"`
	Destport               string   `json:"destport,omitempty"`
	Svctype                string   `json:"svctype,omitempty"`
	Idletime               string   `json:"idletime,omitempty"`
	State                  string   `json:"state,omitempty"`
	Linksourceip           string   `json:"linksourceip,omitempty"`
	Linksourceport         string   `json:"linksourceport,omitempty"`
	Linkdestip             string   `json:"linkdestip,omitempty"`
	Linkdestport           string   `json:"linkdestport,omitempty"`
	Linkservicetype        string   `json:"linkservicetype,omitempty"`
	Linkidletime           string   `json:"linkidletime,omitempty"`
	Linkstate              string   `json:"linkstate,omitempty"`
	Entityname             string   `json:"entityname,omitempty"`
	Linkentityname         string   `json:"linkentityname,omitempty"`
	Connid                 string   `json:"connid,omitempty"`
	Linkconnid             string   `json:"linkconnid,omitempty"`
	Connproperties         string   `json:"connproperties,omitempty"`
	Optionflags            string   `json:"optionflags,omitempty"`
	Nswsvalue              string   `json:"nswsvalue,omitempty"`
	Peerwsvalue            string   `json:"peerwsvalue,omitempty"`
	Mss                    string   `json:"mss,omitempty"`
	Retxretrycnt           string   `json:"retxretrycnt,omitempty"`
	Rcvwnd                 string   `json:"rcvwnd,omitempty"`
	Advwnd                 string   `json:"advwnd,omitempty"`
	Sndcwnd                string   `json:"sndcwnd,omitempty"`
	Iss                    string   `json:"iss,omitempty"`
	Irs                    string   `json:"irs,omitempty"`
	Rcvnxt                 string   `json:"rcvnxt,omitempty"`
	Maxack                 string   `json:"maxack,omitempty"`
	Sndnxt                 string   `json:"sndnxt,omitempty"`
	Sndunack               string   `json:"sndunack,omitempty"`
	Httpendseq             string   `json:"httpendseq,omitempty"`
	Httpstate              string   `json:"httpstate,omitempty"`
	Trcount                string   `json:"trcount,omitempty"`
	Priority               string   `json:"priority,omitempty"`
	Httpreqver             string   `json:"httpreqver,omitempty"`
	Httprequest            string   `json:"httprequest,omitempty"`
	Httprspcode            string   `json:"httprspcode,omitempty"`
	Rttsmoothed            string   `json:"rttsmoothed,omitempty"`
	Rttvariance            string   `json:"rttvariance,omitempty"`
	Outoforderpkts         string   `json:"outoforderpkts,omitempty"`
	Linkoptionflag         string   `json:"linkoptionflag,omitempty"`
	Linknswsvalue          string   `json:"linknswsvalue,omitempty"`
	Linkpeerwsvalue        string   `json:"linkpeerwsvalue,omitempty"`
	Targetnodeidnnm        string   `json:"targetnodeidnnm,omitempty"`
	Sourcenodeidnnm        string   `json:"sourcenodeidnnm,omitempty"`
	Channelidnnm           string   `json:"channelidnnm,omitempty"`
	Msgversionnnm          string   `json:"msgversionnnm,omitempty"`
	Td                     string   `json:"td,omitempty"`
	Maxrcvbuf              string   `json:"maxrcvbuf,omitempty"`
	Linkmaxrcvbuf          string   `json:"linkmaxrcvbuf,omitempty"`
	Rxqsize                string   `json:"rxqsize,omitempty"`
	Linkrxqsize            string   `json:"linkrxqsize,omitempty"`
	Maxsndbuf              string   `json:"maxsndbuf,omitempty"`
	Linkmaxsndbuf          string   `json:"linkmaxsndbuf,omitempty"`
	Txqsize                string   `json:"txqsize,omitempty"`
	Linktxqsize            string   `json:"linktxqsize,omitempty"`
	Flavor                 string   `json:"flavor,omitempty"`
	Linkflavor             string   `json:"linkflavor,omitempty"`
	Bwestimate             string   `json:"bwestimate,omitempty"`
	Linkbwestimate         string   `json:"linkbwestimate,omitempty"`
	Rttmin                 string   `json:"rttmin,omitempty"`
	Linkrttmin             string   `json:"linkrttmin,omitempty"`
	Name                   string   `json:"name,omitempty"`
	Linkname               string   `json:"linkname,omitempty"`
	Tcpmode                string   `json:"tcpmode,omitempty"`
	Linktcpmode            string   `json:"linktcpmode,omitempty"`
	Realtimertt            string   `json:"realtimertt,omitempty"`
	Linkrealtimertt        string   `json:"linkrealtimertt,omitempty"`
	Sndbuf                 string   `json:"sndbuf,omitempty"`
	Linksndbuf             string   `json:"linksndbuf,omitempty"`
	Nsbtcpwaitq            string   `json:"nsbtcpwaitq,omitempty"`
	Linknsbtcpwaitq        string   `json:"linknsbtcpwaitq,omitempty"`
	Nsbretxq               string   `json:"nsbretxq,omitempty"`
	Linknsbretxq           string   `json:"linknsbretxq,omitempty"`
	Sackblocks             string   `json:"sackblocks,omitempty"`
	Linksackblocks         string   `json:"linksackblocks,omitempty"`
	Congstate              string   `json:"congstate,omitempty"`
	Linkcongstate          string   `json:"linkcongstate,omitempty"`
	Sndrecoverle           string   `json:"sndrecoverle,omitempty"`
	Linksndrecoverle       string   `json:"linksndrecoverle,omitempty"`
	Creditsinbytes         string   `json:"creditsinbytes,omitempty"`
	Linkcredits            string   `json:"linkcredits,omitempty"`
	Rateinbytes            string   `json:"rateinbytes,omitempty"`
	Linkrateinbytes        string   `json:"linkrateinbytes,omitempty"`
	Rateschedulerqueue     string   `json:"rateschedulerqueue,omitempty"`
	Linkrateschedulerqueue string   `json:"linkrateschedulerqueue,omitempty"`
	Burstratecontrol       string   `json:"burstratecontrol,omitempty"`
	Linkburstratecontrol   string   `json:"linkburstratecontrol,omitempty"`
	Cqabifavg              string   `json:"cqabifavg,omitempty"`
	Cqathruputavg          string   `json:"cqathruputavg,omitempty"`
	Cqarcvwndavg           string   `json:"cqarcvwndavg,omitempty"`
	Cqaiai1mspct           string   `json:"cqaiai1mspct,omitempty"`
	Cqaiai2mspct           string   `json:"cqaiai2mspct,omitempty"`
	Cqasamples             string   `json:"cqasamples,omitempty"`
	Cqaiaisamples          string   `json:"cqaiaisamples,omitempty"`
	Cqanetclass            string   `json:"cqanetclass,omitempty"`
	Cqaccl                 string   `json:"cqaccl,omitempty"`
	Cqacsq                 string   `json:"cqacsq,omitempty"`
	Cqaiaiavg              string   `json:"cqaiaiavg,omitempty"`
	Cqaisiavg              string   `json:"cqaisiavg,omitempty"`
	Cqarcvwndmin           string   `json:"cqarcvwndmin,omitempty"`
	Cqaretxcorr            string   `json:"cqaretxcorr,omitempty"`
	Cqaretxcong            string   `json:"cqaretxcong,omitempty"`
	Cqaretxpackets         string   `json:"cqaretxpackets,omitempty"`
	Cqaloaddelayavg        string   `json:"cqaloaddelayavg,omitempty"`
	Cqanoisedelayavg       string   `json:"cqanoisedelayavg,omitempty"`
	Cqarttmax              string   `json:"cqarttmax,omitempty"`
	Cqarttmin              string   `json:"cqarttmin,omitempty"`
	Cqarttavg              string   `json:"cqarttavg,omitempty"`
	Adaptivetcpprofname    string   `json:"adaptivetcpprofname,omitempty"`
	Outoforderblocks       string   `json:"outoforderblocks,omitempty"`
	Outoforderflushedcount string   `json:"outoforderflushedcount,omitempty"`
	Outoforderbytes        string   `json:"outoforderbytes,omitempty"`
	Nextgenapiresource     string   `json:"_nextgenapiresource,omitempty"`
}

type Nsevents struct {
	Eventno            int    `json:"eventno,omitempty"`
	Time               string `json:"time,omitempty"`
	Eventcode          string `json:"eventcode,omitempty"`
	Devid              string `json:"devid,omitempty"`
	Devname            string `json:"devname,omitempty"`
	Text               string `json:"text,omitempty"`
	Data0              string `json:"data0,omitempty"`
	Data1              string `json:"data1,omitempty"`
	Data2              string `json:"data2,omitempty"`
	Data3              string `json:"data3,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nshostname struct {
	Hostname           string `json:"hostname,omitempty"`
	Ownernode          int    `json:"ownernode,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nslimitsessions struct {
	Limitidentifier    string `json:"limitidentifier,omitempty"`
	Detail             bool   `json:"detail,omitempty"`
	Timeout            string `json:"timeout,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Drop               string `json:"drop,omitempty"`
	Number             string `json:"number,omitempty"`
	Name               string `json:"name,omitempty"`
	Unit               string `json:"unit,omitempty"`
	Flags              string `json:"flags,omitempty"`
	Referencecount     string `json:"referencecount,omitempty"`
	Maxbandwidth       string `json:"maxbandwidth,omitempty"`
	Selectoripv61      string `json:"selectoripv61,omitempty"`
	Selectoripv62      string `json:"selectoripv62,omitempty"`
	Flag               string `json:"flag,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nsmgmtparam struct {
	Mgmthttpport       int    `json:"mgmthttpport,omitempty"`
	Mgmthttpsport      int    `json:"mgmthttpsport,omitempty"`
	Httpdmaxclients    int    `json:"httpdmaxclients,omitempty"`
	Httpdmaxreqworkers int    `json:"httpdmaxreqworkers,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nssurgeq struct {
	Name       string `json:"name,omitempty"`
	Servername string `json:"servername,omitempty"`
	Port       int    `json:"port,omitempty"`
}

type Nstcpbufparam struct {
	Size               int    `json:"size,omitempty"`
	Memlimit           int    `json:"memlimit"`
	Builtin            string `json:"builtin,omitempty"`
	Feature            string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nsxmlnamespace struct {
	Prefix             string `json:"prefix,omitempty"`
	Namespace          string `json:"Namespace,omitempty"`
	Description        string `json:"description,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nssavedconfig struct {
	Textblob           string `json:"textblob,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Shutdown struct {
}

type Nsacls6 struct {
	Type string `json:"type,omitempty"`
}

type Nsdhcpparams struct {
	Dhcpclient         string `json:"dhcpclient,omitempty"`
	Saveroute          string `json:"saveroute,omitempty"`
	Ipaddress          string `json:"ipaddress,omitempty"`
	Netmask            string `json:"netmask,omitempty"`
	Hostrtgw           string `json:"hostrtgw,omitempty"`
	Running            string `json:"running,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nsextensionextensionfunctionbinding struct {
	Extensionfunctionname           string   `json:"extensionfunctionname,omitempty"`
	Extensionfunctionlinenumber     int      `json:"extensionfunctionlinenumber,omitempty"`
	Extensionfunctionclasstype      string   `json:"extensionfunctionclasstype,omitempty"`
	Extensionfunctionreturntype     string   `json:"extensionfunctionreturntype,omitempty"`
	Activeextensionfunction         int      `json:"activeextensionfunction,omitempty"`
	Extensionfunctionargtype        []string `json:"extensionfunctionargtype,omitempty"`
	Extensionfuncdescription        string   `json:"extensionfuncdescription,omitempty"`
	Extensionfunctionargcount       int      `json:"extensionfunctionargcount,omitempty"`
	Extensionfunctionclasses        []string `json:"extensionfunctionclasses,omitempty"`
	Extensionfunctionclassescount   int      `json:"extensionfunctionclassescount,omitempty"`
	Extensionfunctionallparams      []string `json:"extensionfunctionallparams,omitempty"`
	Extensionfunctionallparamscount int      `json:"extensionfunctionallparamscount,omitempty"`
	Name                            string   `json:"name,omitempty"`
}

type Nsnextgenapi struct {
	State              string `json:"state,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nsparam struct {
	Httpport                  []int  `json:"httpport,omitempty"`
	Maxconn                   int    `json:"maxconn"`
	Maxreq                    int    `json:"maxreq"`
	Cip                       string `json:"cip,omitempty"`
	Cipheader                 string `json:"cipheader,omitempty"`
	Cookieversion             string `json:"cookieversion"`
	Securecookie              string `json:"securecookie,omitempty"`
	Pmtumin                   int    `json:"pmtumin,omitempty"`
	Pmtutimeout               int    `json:"pmtutimeout,omitempty"`
	Ftpportrange              string `json:"ftpportrange,omitempty"`
	Crportrange               string `json:"crportrange,omitempty"`
	Timezone                  string `json:"timezone,omitempty"`
	Grantquotamaxclient       int    `json:"grantquotamaxclient"`
	Exclusivequotamaxclient   int    `json:"exclusivequotamaxclient"`
	Grantquotaspillover       int    `json:"grantquotaspillover,omitempty"`
	Exclusivequotaspillover   int    `json:"exclusivequotaspillover"`
	Useproxyport              string `json:"useproxyport,omitempty"`
	Internaluserlogin         string `json:"internaluserlogin,omitempty"`
	Aftpallowrandomsourceport string `json:"aftpallowrandomsourceport,omitempty"`
	Icaports                  []int  `json:"icaports,omitempty"`
	Tcpcip                    string `json:"tcpcip,omitempty"`
	Servicepathingressvlan    int    `json:"servicepathingressvlan,omitempty"`
	Secureicaports            []int  `json:"secureicaports,omitempty"`
	Mgmthttpport              int    `json:"mgmthttpport,omitempty"`
	Mgmthttpsport             int    `json:"mgmthttpsport,omitempty"`
	Proxyprotocol             string `json:"proxyprotocol,omitempty"`
	Advancedanalyticsstats    string `json:"advancedanalyticsstats,omitempty"`
	Ipttl                     int    `json:"ipttl,omitempty"`
	Autoscaleoption           string `json:"autoscaleoption,omitempty"`
	Nextgenapiresource        string `json:"_nextgenapiresource,omitempty"`
}

type Nspartitionbinding struct {
	Partitionname string `json:"partitionname,omitempty"`
}

type Nspartitionvxlanbinding struct {
	Vxlan         int    `json:"vxlan,omitempty"`
	Partitionname string `json:"partitionname,omitempty"`
}

type Nspbr struct {
	Name               string `json:"name,omitempty"`
	Action             string `json:"action,omitempty"`
	Td                 int    `json:"td,omitempty"`
	Srcip              bool   `json:"srcip,omitempty"`
	Srcipop            string `json:"srcipop,omitempty"`
	Srcipval           string `json:"srcipval,omitempty"`
	Srcport            bool   `json:"srcport,omitempty"`
	Srcportop          string `json:"srcportop,omitempty"`
	Srcportval         string `json:"srcportval,omitempty"`
	Destip             bool   `json:"destip,omitempty"`
	Destipop           string `json:"destipop,omitempty"`
	Destipval          string `json:"destipval,omitempty"`
	Destport           bool   `json:"destport,omitempty"`
	Destportop         string `json:"destportop,omitempty"`
	Destportval        string `json:"destportval,omitempty"`
	Nexthop            bool   `json:"nexthop,omitempty"`
	Nexthopval         string `json:"nexthopval,omitempty"`
	Iptunnel           bool   `json:"iptunnel,omitempty"`
	Iptunnelname       string `json:"iptunnelname,omitempty"`
	Vxlanvlanmap       string `json:"vxlanvlanmap,omitempty"`
	Targettd           int    `json:"targettd,omitempty"`
	Srcmac             string `json:"srcmac,omitempty"`
	Srcmacmask         string `json:"srcmacmask,omitempty"`
	Protocol           string `json:"protocol,omitempty"`
	Protocolnumber     int    `json:"protocolnumber,omitempty"`
	Vlan               int    `json:"vlan,omitempty"`
	Vxlan              int    `json:"vxlan,omitempty"`
	Interface          string `json:"Interface,omitempty"`
	Priority           int    `json:"priority,omitempty"`
	Msr                string `json:"msr,omitempty"`
	Monitor            string `json:"monitor,omitempty"`
	State              string `json:"state,omitempty"`
	Ownergroup         string `json:"ownergroup,omitempty"`
	Detail             bool   `json:"detail,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Kernelstate        string `json:"kernelstate,omitempty"`
	Curstate           string `json:"curstate,omitempty"`
	Totalprobes        string `json:"totalprobes,omitempty"`
	Totalfailedprobes  string `json:"totalfailedprobes,omitempty"`
	Failedprobes       string `json:"failedprobes,omitempty"`
	Monstatcode        string `json:"monstatcode,omitempty"`
	Monstatparam1      string `json:"monstatparam1,omitempty"`
	Monstatparam2      string `json:"monstatparam2,omitempty"`
	Monstatparam3      string `json:"monstatparam3,omitempty"`
	Data               string `json:"data,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nsacl struct {
	Aclname            string `json:"aclname,omitempty"`
	Aclaction          string `json:"aclaction,omitempty"`
	Td                 int    `json:"td,omitempty"`
	Srcip              bool   `json:"srcip,omitempty"`
	Srcipop            string `json:"srcipop,omitempty"`
	Srcipval           string `json:"srcipval,omitempty"`
	Srcipdataset       string `json:"srcipdataset,omitempty"`
	Srcport            bool   `json:"srcport,omitempty"`
	Srcportop          string `json:"srcportop,omitempty"`
	Srcportval         string `json:"srcportval,omitempty"`
	Srcportdataset     string `json:"srcportdataset,omitempty"`
	Destip             bool   `json:"destip,omitempty"`
	Destipop           string `json:"destipop,omitempty"`
	Destipval          string `json:"destipval,omitempty"`
	Destipdataset      string `json:"destipdataset,omitempty"`
	Destport           bool   `json:"destport,omitempty"`
	Destportop         string `json:"destportop,omitempty"`
	Destportval        string `json:"destportval,omitempty"`
	Destportdataset    string `json:"destportdataset,omitempty"`
	Ttl                int    `json:"ttl,omitempty"`
	Srcmac             string `json:"srcmac,omitempty"`
	Srcmacmask         string `json:"srcmacmask,omitempty"`
	Protocol           string `json:"protocol,omitempty"`
	Protocolnumber     int    `json:"protocolnumber,omitempty"`
	Vlan               int    `json:"vlan,omitempty"`
	Vxlan              int    `json:"vxlan,omitempty"`
	Interface          string `json:"Interface,omitempty"`
	Established        bool   `json:"established,omitempty"`
	Icmptype           int    `json:"icmptype,omitempty"`
	Icmpcode           int    `json:"icmpcode,omitempty"`
	Priority           int    `json:"priority,omitempty"`
	State              string `json:"state,omitempty"`
	Logstate           string `json:"logstate,omitempty"`
	Ratelimit          int    `json:"ratelimit,omitempty"`
	Type               string `json:"type,omitempty"`
	Dfdhash            string `json:"dfdhash,omitempty"`
	Nodeid             int    `json:"nodeid,omitempty"`
	Stateful           string `json:"stateful,omitempty"`
	Newname            string `json:"newname,omitempty"`
	Hits               string `json:"hits,omitempty"`
	Kernelstate        string `json:"kernelstate,omitempty"`
	Aclassociate       string `json:"aclassociate,omitempty"`
	Aclchildcount      string `json:"aclchildcount,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nsacls struct {
	Type string `json:"type,omitempty"`
}

type Nscentralmanagementserver struct {
	Type                       string `json:"type,omitempty"`
	Username                   string `json:"username,omitempty"`
	Password                   string `json:"password,omitempty"`
	Activationcode             string `json:"activationcode,omitempty"`
	Ipaddress                  string `json:"ipaddress,omitempty"`
	Servername                 string `json:"servername,omitempty"`
	Validatecert               string `json:"validatecert,omitempty"`
	Deviceprofilename          string `json:"deviceprofilename,omitempty"`
	Adcusername                string `json:"adcusername,omitempty"`
	Adcpassword                string `json:"adcpassword,omitempty"`
	Instanceid                 string `json:"instanceid,omitempty"`
	Customerid                 string `json:"customerid,omitempty"`
	Admserviceenvironment      string `json:"admserviceenvironment,omitempty"`
	Admserviceconnectionstatus string `json:"admserviceconnectionstatus,omitempty"`
	Nextgenapiresource         string `json:"_nextgenapiresource,omitempty"`
}

type Nsdhcpip struct {
}

type Nsrollbackcmd struct {
	Filename           string `json:"filename,omitempty"`
	Outtype            string `json:"outtype,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nsservicepathbinding struct {
	Servicepathname string `json:"servicepathname,omitempty"`
}

type Nstcpparam struct {
	Ws                                  string `json:"ws,omitempty"`
	Wsval                               int    `json:"wsval"`
	Sack                                string `json:"sack,omitempty"`
	Learnvsvrmss                        string `json:"learnvsvrmss,omitempty"`
	Maxburst                            int    `json:"maxburst,omitempty"`
	Initialcwnd                         int    `json:"initialcwnd,omitempty"`
	Recvbuffsize                        int    `json:"recvbuffsize,omitempty"`
	Delayedack                          int    `json:"delayedack,omitempty"`
	Downstaterst                        string `json:"downstaterst,omitempty"`
	Nagle                               string `json:"nagle,omitempty"`
	Limitedpersist                      string `json:"limitedpersist,omitempty"`
	Oooqsize                            int    `json:"oooqsize"`
	Ackonpush                           string `json:"ackonpush,omitempty"`
	Maxpktpermss                        int    `json:"maxpktpermss"`
	Pktperretx                          int    `json:"pktperretx,omitempty"`
	Minrto                              int    `json:"minrto,omitempty"`
	Slowstartincr                       int    `json:"slowstartincr,omitempty"`
	Maxdynserverprobes                  int    `json:"maxdynserverprobes,omitempty"`
	Synholdfastgiveup                   int    `json:"synholdfastgiveup,omitempty"`
	Maxsynholdperprobe                  int    `json:"maxsynholdperprobe,omitempty"`
	Maxsynhold                          int    `json:"maxsynhold,omitempty"`
	Msslearninterval                    int    `json:"msslearninterval,omitempty"`
	Msslearndelay                       int    `json:"msslearndelay,omitempty"`
	Maxtimewaitconn                     int    `json:"maxtimewaitconn,omitempty"`
	Kaprobeupdatelastactivity           string `json:"kaprobeupdatelastactivity,omitempty"`
	Maxsynackretx                       int    `json:"maxsynackretx,omitempty"`
	Synattackdetection                  string `json:"synattackdetection,omitempty"`
	Connflushifnomem                    string `json:"connflushifnomem,omitempty"`
	Connflushthres                      int    `json:"connflushthres,omitempty"`
	Mptcpconcloseonpassivesf            string `json:"mptcpconcloseonpassivesf,omitempty"`
	Mptcpchecksum                       string `json:"mptcpchecksum,omitempty"`
	Mptcpsftimeout                      int    `json:""`
	Mptcpsfreplacetimeout               int    `json:"mptcpsfreplacetimeout"`
	Mptcpmaxsf                          int    `json:"mptcpmaxsf,omitempty"`
	Mptcpmaxpendingsf                   int    `json:"mptcpmaxpendingsf"`
	Mptcppendingjointhreshold           int    `json:"mptcppendingjointhreshold"`
	Mptcprtostoswitchsf                 int    `json:"mptcprtostoswitchsf,omitempty"`
	Mptcpusebackupondss                 string `json:"mptcpusebackupondss,omitempty"`
	Tcpmaxretries                       int    `json:"tcpmaxretries,omitempty"`
	Mptcpimmediatesfcloseonfin          string `json:"mptcpimmediatesfcloseonfin,omitempty"`
	Mptcpclosemptcpsessiononlastsfclose string `json:"mptcpclosemptcpsessiononlastsfclose,omitempty"`
	Mptcpsendsfresetoption              string `json:"mptcpsendsfresetoption,omitempty"`
	Mptcpfastcloseoption                string `json:"mptcpfastcloseoption,omitempty"`
	Mptcpreliableaddaddr                string `json:"mptcpreliableaddaddr,omitempty"`
	Tcpfastopencookietimeout            int    `json:"tcpfastopencookietimeout"`
	Autosyncookietimeout                int    `json:"autosyncookietimeout,omitempty"`
	Tcpfintimeout                       int    `json:"tcpfintimeout,omitempty"`
	Compacttcpoptionnoop                string `json:"compacttcpoptionnoop,omitempty"`
	Delinkclientserveronrst             string `json:"delinkclientserveronrst,omitempty"`
	Rfc5961chlgacklimit                 int    `json:"rfc5961chlgacklimit,omitempty"`
	Enhancedisngeneration               string `json:"enhancedisngeneration,omitempty"`
	Builtin                             string `json:"builtin,omitempty"`
	Feature                             string `json:"feature,omitempty"`
	Nextgenapiresource                  string `json:"_nextgenapiresource,omitempty"`
}

type Nsvpxparam struct {
	Masterclockcpu1     string `json:"masterclockcpu1,omitempty"`
	Cpuyield            string `json:"cpuyield,omitempty"`
	Ownernode           int    `json:"ownernode"`
	Kvmvirtiomultiqueue string `json:"kvmvirtiomultiqueue,omitempty"`
	Vpxenvironment      string `json:"vpxenvironment,omitempty"`
	Memorystatus        string `json:"memorystatus,omitempty"`
	Cloudproductcode    string `json:"cloudproductcode,omitempty"`
	Vpxoemcode          string `json:"vpxoemcode,omitempty"`
	Technicalsupportpin string `json:"technicalsupportpin,omitempty"`
	Nextgenapiresource  string `json:"_nextgenapiresource,omitempty"`
}

type Nstcpprofile struct {
	Name                        string `json:"name,omitempty"`
	Ws                          string `json:"ws,omitempty"`
	Sack                        string `json:"sack,omitempty"`
	Wsval                       int    `json:"wsval,omitempty"`
	Nagle                       string `json:"nagle,omitempty"`
	Ackonpush                   string `json:"ackonpush,omitempty"`
	Mss                         int    `json:"mss,omitempty"`
	Maxburst                    int    `json:"maxburst,omitempty"`
	Initialcwnd                 int    `json:"initialcwnd,omitempty"`
	Delayedack                  int    `json:"delayedack,omitempty"`
	Oooqsize                    int    `json:"oooqsize,omitempty"`
	Maxpktpermss                int    `json:"maxpktpermss,omitempty"`
	Pktperretx                  int    `json:"pktperretx,omitempty"`
	Minrto                      int    `json:"minrto,omitempty"`
	Slowstartincr               int    `json:"slowstartincr,omitempty"`
	Buffersize                  int    `json:"buffersize,omitempty"`
	Syncookie                   string `json:"syncookie,omitempty"`
	Kaprobeupdatelastactivity   string `json:"kaprobeupdatelastactivity,omitempty"`
	Flavor                      string `json:"flavor,omitempty"`
	Dynamicreceivebuffering     string `json:"dynamicreceivebuffering,omitempty"`
	Ka                          string `json:"ka,omitempty"`
	Kaconnidletime              int    `json:"kaconnidletime,omitempty"`
	Kamaxprobes                 int    `json:"kamaxprobes,omitempty"`
	Kaprobeinterval             int    `json:"kaprobeinterval,omitempty"`
	Sendbuffsize                int    `json:"sendbuffsize,omitempty"`
	Mptcp                       string `json:"mptcp,omitempty"`
	Establishclientconn         string `json:"establishclientconn,omitempty"`
	Tcpsegoffload               string `json:"tcpsegoffload,omitempty"`
	Rfc5961compliance           string `json:"rfc5961compliance,omitempty"`
	Rstwindowattenuate          string `json:"rstwindowattenuate,omitempty"`
	Rstmaxack                   string `json:"rstmaxack,omitempty"`
	Spoofsyndrop                string `json:"spoofsyndrop,omitempty"`
	Ecn                         string `json:"ecn,omitempty"`
	Mptcpdropdataonpreestsf     string `json:"mptcpdropdataonpreestsf,omitempty"`
	Mptcpfastopen               string `json:"mptcpfastopen,omitempty"`
	Mptcpsessiontimeout         int    `json:"mptcpsessiontimeout,omitempty"`
	Timestamp                   string `json:"timestamp,omitempty"`
	Dsack                       string `json:"dsack,omitempty"`
	Ackaggregation              string `json:"ackaggregation,omitempty"`
	Frto                        string `json:"frto,omitempty"`
	Maxcwnd                     int    `json:"maxcwnd,omitempty"`
	Fack                        string `json:"fack,omitempty"`
	Tcpmode                     string `json:"tcpmode,omitempty"`
	Tcpfastopen                 string `json:"tcpfastopen,omitempty"`
	Hystart                     string `json:"hystart,omitempty"`
	Dupackthresh                int    `json:"dupackthresh,omitempty"`
	Burstratecontrol            string `json:"burstratecontrol,omitempty"`
	Tcprate                     int    `json:"tcprate,omitempty"`
	Rateqmax                    int    `json:"rateqmax,omitempty"`
	Drophalfclosedconnontimeout string `json:"drophalfclosedconnontimeout,omitempty"`
	Dropestconnontimeout        string `json:"dropestconnontimeout,omitempty"`
	Applyadaptivetcp            string `json:"applyadaptivetcp,omitempty"`
	Tcpfastopencookiesize       int    `json:"tcpfastopencookiesize,omitempty"`
	Taillossprobe               string `json:"taillossprobe,omitempty"`
	Clientiptcpoption           string `json:"clientiptcpoption,omitempty"`
	Clientiptcpoptionnumber     int    `json:"clientiptcpoptionnumber,omitempty"`
	Mpcapablecbit               string `json:"mpcapablecbit,omitempty"`
	Sendclientportintcpoption   string `json:"sendclientportintcpoption,omitempty"`
	Slowstartthreshold          int    `json:"slowstartthreshold,omitempty"`
	Refcnt                      string `json:"refcnt,omitempty"`
	Builtin                     string `json:"builtin,omitempty"`
	Feature                     string `json:"feature,omitempty"`
	Nextgenapiresource          string `json:"_nextgenapiresource,omitempty"`
}

type Nstimerpolicybinding struct {
	Policyname             string `json:"policyname,omitempty"`
	Priority               uint32 `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Vserver                string `json:"vserver,omitempty"`
	Samplesize             uint32 `json:"samplesize,omitempty"`
	Threshold              uint32 `json:"threshold,omitempty"`
	Name                   string `json:"name,omitempty"`
}

type Nstrafficdomainbridgegroupbinding struct {
	Bridgegroup int `json:"bridgegroup,omitempty"`
	Td          int `json:"td,omitempty"`
}

type Nsvariable struct {
	Name               string `json:"name,omitempty"`
	Type               string `json:"type,omitempty"`
	Scope              string `json:"scope,omitempty"`
	Iffull             string `json:"iffull,omitempty"`
	Ifvaluetoobig      string `json:"ifvaluetoobig,omitempty"`
	Ifnovalue          string `json:"ifnovalue,omitempty"`
	Init               string `json:"init,omitempty"`
	Expires            int    `json:"expires,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Referencecount     string `json:"referencecount,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Reboot struct {
	Warm bool `json:"warm,omitempty"`
}

type Nsencryptionkey struct {
	Name               string `json:"name,omitempty"`
	Method             string `json:"method,omitempty"`
	Keyvalue           string `json:"keyvalue,omitempty"`
	Padding            string `json:"padding,omitempty"`
	Iv                 string `json:"iv,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nsextension struct {
	Src                string `json:"src,omitempty"`
	Name               string `json:"name,omitempty"`
	Comment            string `json:"comment,omitempty"`
	Overwrite          bool   `json:"overwrite,omitempty"`
	Trace              string `json:"trace,omitempty"`
	Tracefunctions     string `json:"tracefunctions,omitempty"`
	Tracevariables     string `json:"tracevariables,omitempty"`
	Detail             string `json:"detail,omitempty"`
	Type               string `json:"type,omitempty"`
	Functionhits       string `json:"functionhits,omitempty"`
	Functionundefhits  string `json:"functionundefhits,omitempty"`
	Functionhaltcount  string `json:"functionhaltcount,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nskeymanagerproxy struct {
	Serverip           string `json:"serverip,omitempty"`
	Servername         string `json:"servername,omitempty"`
	Port               int    `json:"port,omitempty"`
	Nodeid             int    `json:"nodeid,omitempty"`
	Status             string `json:"status,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nsservicepathservicefunctionbinding struct {
	Servicefunction string `json:"servicefunction,omitempty"`
	Index           uint32 `json:"index,omitempty"`
	Servicepathname string `json:"servicepathname,omitempty"`
}

type Nstestlicense struct {
	Wl                      string `json:"wl,omitempty"`
	Sp                      string `json:"sp,omitempty"`
	Lb                      string `json:"lb,omitempty"`
	Cs                      string `json:"cs,omitempty"`
	Cr                      string `json:"cr,omitempty"`
	Cmp                     string `json:"cmp,omitempty"`
	Delta                   string `json:"delta,omitempty"`
	Ssl                     string `json:"ssl,omitempty"`
	Gslb                    string `json:"gslb,omitempty"`
	Gslbp                   string `json:"gslbp,omitempty"`
	Routing                 string `json:"routing,omitempty"`
	Cf                      string `json:"cf,omitempty"`
	Contentaccelerator      string `json:"contentaccelerator,omitempty"`
	Ic                      string `json:"ic,omitempty"`
	Sslvpn                  string `json:"sslvpn,omitempty"`
	Fsslvpnusers            string `json:"f_sslvpn_users,omitempty"`
	Ficausers               string `json:"f_ica_users,omitempty"`
	Aaa                     string `json:"aaa,omitempty"`
	Ospf                    string `json:"ospf,omitempty"`
	Rip                     string `json:"rip,omitempty"`
	Bgp                     string `json:"bgp,omitempty"`
	Rewrite                 string `json:"rewrite,omitempty"`
	Ipv6pt                  string `json:"ipv6pt,omitempty"`
	Appfw                   string `json:"appfw,omitempty"`
	Responder               string `json:"responder,omitempty"`
	Agee                    string `json:"agee,omitempty"`
	Nsxn                    string `json:"nsxn,omitempty"`
	Modelid                 string `json:"modelid,omitempty"`
	Push                    string `json:"push,omitempty"`
	Appflow                 string `json:"appflow,omitempty"`
	Cloudbridge             string `json:"cloudbridge,omitempty"`
	Cloudbridgeappliance    string `json:"cloudbridgeappliance,omitempty"`
	Cloudextenderappliance  string `json:"cloudextenderappliance,omitempty"`
	Isis                    string `json:"isis,omitempty"`
	Cluster                 string `json:"cluster,omitempty"`
	Ch                      string `json:"ch,omitempty"`
	Appqoe                  string `json:"appqoe,omitempty"`
	Appflowica              string `json:"appflowica,omitempty"`
	Isstandardlic           string `json:"isstandardlic,omitempty"`
	Isenterpriselic         string `json:"isenterpriselic,omitempty"`
	Isplatinumlic           string `json:"isplatinumlic,omitempty"`
	Issgwylic               string `json:"issgwylic,omitempty"`
	Isswglic                string `json:"isswglic,omitempty"`
	Feo                     string `json:"feo,omitempty"`
	Lsn                     string `json:"lsn,omitempty"`
	Licensingmode           string `json:"licensingmode,omitempty"`
	Daystoexpiration        string `json:"daystoexpiration,omitempty"`
	Rdpproxy                string `json:"rdpproxy,omitempty"`
	Rep                     string `json:"rep,omitempty"`
	Urlfiltering            string `json:"urlfiltering,omitempty"`
	Videooptimization       string `json:"videooptimization,omitempty"`
	Forwardproxy            string `json:"forwardproxy,omitempty"`
	Sslinterception         string `json:"sslinterception,omitempty"`
	Remotecontentinspection string `json:"remotecontentinspection,omitempty"`
	Adaptivetcp             string `json:"adaptivetcp,omitempty"`
	Cqa                     string `json:"cqa,omitempty"`
	Bot                     string `json:"bot,omitempty"`
	Apigateway              string `json:"apigateway,omitempty"`
	Nextgenapiresource      string `json:"_nextgenapiresource,omitempty"`
}

type Nstrafficdomainbinding struct {
	Td int `json:"td,omitempty"`
}

type Nslicenseparameters struct {
	Alert1gracetimeout       int    `json:"alert1gracetimeout"`
	Alert2gracetimeout       int    `json:"alert2gracetimeout,omitempty"`
	Licenseexpiryalerttime   int    `json:"licenseexpiryalerttime,omitempty"`
	Heartbeatinterval        int    `json:"heartbeatinterval,omitempty"`
	Inventoryrefreshinterval int    `json:"inventoryrefreshinterval,omitempty"`
	Nextgenapiresource       string `json:"_nextgenapiresource,omitempty"`
}

type Nsmemrecovery struct {
	Percentage int `json:"percentage,omitempty"`
}

type Nsstats struct {
	Cleanuplevel string `json:"cleanuplevel,omitempty"`
}

type Nstrafficdomain struct {
	Td                 int    `json:"td,omitempty"`
	Aliasname          string `json:"aliasname,omitempty"`
	Vmac               string `json:"vmac,omitempty"`
	State              string `json:"state,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nstrafficdomainvxlanbinding struct {
	Vxlan int `json:"vxlan,omitempty"`
	Td    int `json:"td,omitempty"`
}
