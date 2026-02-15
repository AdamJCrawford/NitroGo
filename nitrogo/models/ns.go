package models

// ns configuration structs
type Nspbr struct {
	Action             string  `json:"action,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Curstate           int     `json:"curstate,omitempty"`
	Data               bool    `json:"data,omitempty"`
	Destip             bool    `json:"destip,omitempty"`
	Destipop           string  `json:"destipop,omitempty"`
	Destipval          string  `json:"destipval,omitempty"`
	Destport           bool    `json:"destport,omitempty"`
	Destportop         string  `json:"destportop,omitempty"`
	Destportval        string  `json:"destportval,omitempty"`
	Detail             bool    `json:"detail,omitempty"`
	Failedprobes       int     `json:"failedprobes,omitempty"`
	Hits               int     `json:"hits,omitempty"`
	InterfaceField     string  `json:"Interface,omitempty"`
	Iptunnel           bool    `json:"iptunnel,omitempty"`
	Iptunnelname       string  `json:"iptunnelname,omitempty"`
	Kernelstate        string  `json:"kernelstate,omitempty"`
	Monitor            string  `json:"monitor,omitempty"`
	Monstatcode        int     `json:"monstatcode,omitempty"`
	Monstatparam1      int     `json:"monstatparam1,omitempty"`
	Monstatparam2      int     `json:"monstatparam2,omitempty"`
	Monstatparam3      int     `json:"monstatparam3,omitempty"`
	Msr                string  `json:"msr,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Nexthop            bool    `json:"nexthop,omitempty"`
	Nexthopval         string  `json:"nexthopval,omitempty"`
	Ownergroup         string  `json:"ownergroup,omitempty"`
	Priority           int     `json:"priority,omitempty"`
	Protocol           string  `json:"protocol,omitempty"`
	Protocolnumber     int     `json:"protocolnumber,omitempty"`
	Srcip              bool    `json:"srcip,omitempty"`
	Srcipop            string  `json:"srcipop,omitempty"`
	Srcipval           string  `json:"srcipval,omitempty"`
	Srcmac             string  `json:"srcmac,omitempty"`
	Srcmacmask         string  `json:"srcmacmask,omitempty"`
	Srcport            bool    `json:"srcport,omitempty"`
	Srcportop          string  `json:"srcportop,omitempty"`
	Srcportval         string  `json:"srcportval,omitempty"`
	State              string  `json:"state,omitempty"`
	Targettd           int     `json:"targettd,omitempty"`
	Td                 int     `json:"td,omitempty"`
	Totalfailedprobes  int     `json:"totalfailedprobes,omitempty"`
	Totalprobes        int     `json:"totalprobes,omitempty"`
	Vlan               int     `json:"vlan,omitempty"`
	Vxlan              int     `json:"vxlan,omitempty"`
	Vxlanvlanmap       string  `json:"vxlanvlanmap,omitempty"`
}

type NstrafficdomainVxlanBinding struct {
	Td    int `json:"td,omitempty"`
	Vxlan int `json:"vxlan,omitempty"`
}

type Nspartition struct {
	Count              float64 `json:"__count,omitempty"`
	Force              bool    `json:"force,omitempty"`
	Maxbandwidth       int     `json:"maxbandwidth,omitempty"`
	Maxconn            int     `json:"maxconn,omitempty"`
	Maxmemlimit        int     `json:"maxmemlimit,omitempty"`
	Minbandwidth       int     `json:"minbandwidth,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Partitionid        int     `json:"partitionid,omitempty"`
	Partitionmac       string  `json:"partitionmac,omitempty"`
	Partitionname      string  `json:"partitionname,omitempty"`
	Partitiontype      string  `json:"partitiontype,omitempty"`
	Pmacinternal       bool    `json:"pmacinternal,omitempty"`
	Save               bool    `json:"save,omitempty"`
}

type NsextensionBinding struct {
	Name                                string        `json:"name,omitempty"`
	NsextensionExtensionfunctionBinding []interface{} `json:"nsextension_extensionfunction_binding,omitempty"`
}

type Nslicense struct {
	Aaa                     bool   `json:"aaa,omitempty"`
	Adaptivetcp             bool   `json:"adaptivetcp,omitempty"`
	Agee                    bool   `json:"agee,omitempty"`
	Apigateway              bool   `json:"apigateway,omitempty"`
	Appflow                 bool   `json:"appflow,omitempty"`
	Appflowica              bool   `json:"appflowica,omitempty"`
	Appfw                   bool   `json:"appfw,omitempty"`
	Appqoe                  bool   `json:"appqoe,omitempty"`
	Bgp                     bool   `json:"bgp,omitempty"`
	Bot                     bool   `json:"bot,omitempty"`
	Cf                      bool   `json:"cf,omitempty"`
	Ch                      bool   `json:"ch,omitempty"`
	Cloudbridge             bool   `json:"cloudbridge,omitempty"`
	Cloudbridgeappliance    bool   `json:"cloudbridgeappliance,omitempty"`
	Cloudextenderappliance  bool   `json:"cloudextenderappliance,omitempty"`
	Cloudsubscriptionimage  string `json:"cloudsubscriptionimage,omitempty"`
	Cluster                 bool   `json:"cluster,omitempty"`
	Cmp                     bool   `json:"cmp,omitempty"`
	Contentaccelerator      bool   `json:"contentaccelerator,omitempty"`
	Cqa                     bool   `json:"cqa,omitempty"`
	Cr                      bool   `json:"cr,omitempty"`
	Cs                      bool   `json:"cs,omitempty"`
	Daystoexpiration        int    `json:"daystoexpiration,omitempty"`
	Daystolasenforcement    int    `json:"daystolasenforcement,omitempty"`
	Delta                   bool   `json:"delta,omitempty"`
	FIcaUsers               int    `json:"f_ica_users,omitempty"`
	FSslvpnUsers            int    `json:"f_sslvpn_users,omitempty"`
	Feo                     bool   `json:"feo,omitempty"`
	Forwardproxy            bool   `json:"forwardproxy,omitempty"`
	Gslb                    bool   `json:"gslb,omitempty"`
	Gslbp                   bool   `json:"gslbp,omitempty"`
	Ic                      bool   `json:"ic,omitempty"`
	Ipv6pt                  bool   `json:"ipv6pt,omitempty"`
	Isenterpriselic         bool   `json:"isenterpriselic,omitempty"`
	Isis                    bool   `json:"isis,omitempty"`
	Isplatinumlic           bool   `json:"isplatinumlic,omitempty"`
	Issgwylic               bool   `json:"issgwylic,omitempty"`
	Isstandardlic           bool   `json:"isstandardlic,omitempty"`
	Isswglic                bool   `json:"isswglic,omitempty"`
	Lb                      bool   `json:"lb,omitempty"`
	Licensingmode           string `json:"licensingmode,omitempty"`
	Lsn                     bool   `json:"lsn,omitempty"`
	Modelid                 int    `json:"modelid,omitempty"`
	Nextgenapiresource      string `json:"_nextgenapiresource,omitempty"`
	Nsxn                    bool   `json:"nsxn,omitempty"`
	Ospf                    bool   `json:"ospf,omitempty"`
	Push                    bool   `json:"push,omitempty"`
	Rdpproxy                bool   `json:"rdpproxy,omitempty"`
	Remotecontentinspection bool   `json:"remotecontentinspection,omitempty"`
	Rep                     bool   `json:"rep,omitempty"`
	Responder               bool   `json:"responder,omitempty"`
	Rewrite                 bool   `json:"rewrite,omitempty"`
	Rip                     bool   `json:"rip,omitempty"`
	Routing                 bool   `json:"routing,omitempty"`
	Sp                      bool   `json:"sp,omitempty"`
	Ssl                     bool   `json:"ssl,omitempty"`
	Sslinterception         bool   `json:"sslinterception,omitempty"`
	Sslvpn                  bool   `json:"sslvpn,omitempty"`
	Urlfiltering            bool   `json:"urlfiltering,omitempty"`
	Videooptimization       bool   `json:"videooptimization,omitempty"`
	Wl                      bool   `json:"wl,omitempty"`
}

type Nstestlicense struct {
	Aaa                     bool   `json:"aaa,omitempty"`
	Adaptivetcp             bool   `json:"adaptivetcp,omitempty"`
	Agee                    bool   `json:"agee,omitempty"`
	Apigateway              bool   `json:"apigateway,omitempty"`
	Appflow                 bool   `json:"appflow,omitempty"`
	Appflowica              bool   `json:"appflowica,omitempty"`
	Appfw                   bool   `json:"appfw,omitempty"`
	Appqoe                  bool   `json:"appqoe,omitempty"`
	Bgp                     bool   `json:"bgp,omitempty"`
	Bot                     bool   `json:"bot,omitempty"`
	Cf                      bool   `json:"cf,omitempty"`
	Ch                      bool   `json:"ch,omitempty"`
	Cloudbridge             bool   `json:"cloudbridge,omitempty"`
	Cloudbridgeappliance    bool   `json:"cloudbridgeappliance,omitempty"`
	Cloudextenderappliance  bool   `json:"cloudextenderappliance,omitempty"`
	Cluster                 bool   `json:"cluster,omitempty"`
	Cmp                     bool   `json:"cmp,omitempty"`
	Contentaccelerator      bool   `json:"contentaccelerator,omitempty"`
	Cqa                     bool   `json:"cqa,omitempty"`
	Cr                      bool   `json:"cr,omitempty"`
	Cs                      bool   `json:"cs,omitempty"`
	Daystoexpiration        int    `json:"daystoexpiration,omitempty"`
	Delta                   bool   `json:"delta,omitempty"`
	FIcaUsers               int    `json:"f_ica_users,omitempty"`
	FSslvpnUsers            int    `json:"f_sslvpn_users,omitempty"`
	Feo                     bool   `json:"feo,omitempty"`
	Forwardproxy            bool   `json:"forwardproxy,omitempty"`
	Gslb                    bool   `json:"gslb,omitempty"`
	Gslbp                   bool   `json:"gslbp,omitempty"`
	Ic                      bool   `json:"ic,omitempty"`
	Ipv6pt                  bool   `json:"ipv6pt,omitempty"`
	Isenterpriselic         bool   `json:"isenterpriselic,omitempty"`
	Isis                    bool   `json:"isis,omitempty"`
	Isplatinumlic           bool   `json:"isplatinumlic,omitempty"`
	Issgwylic               bool   `json:"issgwylic,omitempty"`
	Isstandardlic           bool   `json:"isstandardlic,omitempty"`
	Isswglic                bool   `json:"isswglic,omitempty"`
	Lb                      bool   `json:"lb,omitempty"`
	Licensingmode           string `json:"licensingmode,omitempty"`
	Lsn                     bool   `json:"lsn,omitempty"`
	Modelid                 int    `json:"modelid,omitempty"`
	Nextgenapiresource      string `json:"_nextgenapiresource,omitempty"`
	Nsxn                    bool   `json:"nsxn,omitempty"`
	Ospf                    bool   `json:"ospf,omitempty"`
	Push                    bool   `json:"push,omitempty"`
	Rdpproxy                bool   `json:"rdpproxy,omitempty"`
	Remotecontentinspection bool   `json:"remotecontentinspection,omitempty"`
	Rep                     bool   `json:"rep,omitempty"`
	Responder               bool   `json:"responder,omitempty"`
	Rewrite                 bool   `json:"rewrite,omitempty"`
	Rip                     bool   `json:"rip,omitempty"`
	Routing                 bool   `json:"routing,omitempty"`
	Sp                      bool   `json:"sp,omitempty"`
	Ssl                     bool   `json:"ssl,omitempty"`
	Sslinterception         bool   `json:"sslinterception,omitempty"`
	Sslvpn                  bool   `json:"sslvpn,omitempty"`
	Urlfiltering            bool   `json:"urlfiltering,omitempty"`
	Videooptimization       bool   `json:"videooptimization,omitempty"`
	Wl                      bool   `json:"wl,omitempty"`
}

type Nslicenseserverpool struct {
	Cpxinstanceavailable         int    `json:"cpxinstanceavailable,omitempty"`
	Cpxinstancetotal             int    `json:"cpxinstancetotal,omitempty"`
	Enterprisebandwidthavailable int    `json:"enterprisebandwidthavailable,omitempty"`
	Enterprisebandwidthtotal     int    `json:"enterprisebandwidthtotal,omitempty"`
	Enterprisecpuavailable       int    `json:"enterprisecpuavailable,omitempty"`
	Enterprisecputotal           int    `json:"enterprisecputotal,omitempty"`
	Getalllicenses               bool   `json:"getalllicenses,omitempty"`
	Instanceavailable            int    `json:"instanceavailable,omitempty"`
	Instancetotal                int    `json:"instancetotal,omitempty"`
	Licensemode                  string `json:"licensemode,omitempty"`
	Nextgenapiresource           string `json:"_nextgenapiresource,omitempty"`
	Platinumbandwidthavailable   int    `json:"platinumbandwidthavailable,omitempty"`
	Platinumbandwidthtotal       int    `json:"platinumbandwidthtotal,omitempty"`
	Platinumcpuavailable         int    `json:"platinumcpuavailable,omitempty"`
	Platinumcputotal             int    `json:"platinumcputotal,omitempty"`
	Standardbandwidthavailable   int    `json:"standardbandwidthavailable,omitempty"`
	Standardbandwidthtotal       int    `json:"standardbandwidthtotal,omitempty"`
	Standardcpuavailable         int    `json:"standardcpuavailable,omitempty"`
	Standardcputotal             int    `json:"standardcputotal,omitempty"`
	Vpx100000eavailable          int    `json:"vpx100000eavailable,omitempty"`
	Vpx100000etotal              int    `json:"vpx100000etotal,omitempty"`
	Vpx100000pavailable          int    `json:"vpx100000pavailable,omitempty"`
	Vpx100000ptotal              int    `json:"vpx100000ptotal,omitempty"`
	Vpx100000savailable          int    `json:"vpx100000savailable,omitempty"`
	Vpx100000stotal              int    `json:"vpx100000stotal,omitempty"`
	Vpx10000eavailable           int    `json:"vpx10000eavailable,omitempty"`
	Vpx10000etotal               int    `json:"vpx10000etotal,omitempty"`
	Vpx10000pavailable           int    `json:"vpx10000pavailable,omitempty"`
	Vpx10000ptotal               int    `json:"vpx10000ptotal,omitempty"`
	Vpx10000savailable           int    `json:"vpx10000savailable,omitempty"`
	Vpx10000stotal               int    `json:"vpx10000stotal,omitempty"`
	Vpx1000eavailable            int    `json:"vpx1000eavailable,omitempty"`
	Vpx1000etotal                int    `json:"vpx1000etotal,omitempty"`
	Vpx1000pavailable            int    `json:"vpx1000pavailable,omitempty"`
	Vpx1000ptotal                int    `json:"vpx1000ptotal,omitempty"`
	Vpx1000savailable            int    `json:"vpx1000savailable,omitempty"`
	Vpx1000stotal                int    `json:"vpx1000stotal,omitempty"`
	Vpx100eavailable             int    `json:"vpx100eavailable,omitempty"`
	Vpx100etotal                 int    `json:"vpx100etotal,omitempty"`
	Vpx100pavailable             int    `json:"vpx100pavailable,omitempty"`
	Vpx100ptotal                 int    `json:"vpx100ptotal,omitempty"`
	Vpx100savailable             int    `json:"vpx100savailable,omitempty"`
	Vpx100stotal                 int    `json:"vpx100stotal,omitempty"`
	Vpx10eavailable              int    `json:"vpx10eavailable,omitempty"`
	Vpx10etotal                  int    `json:"vpx10etotal,omitempty"`
	Vpx10pavailable              int    `json:"vpx10pavailable,omitempty"`
	Vpx10ptotal                  int    `json:"vpx10ptotal,omitempty"`
	Vpx10savailable              int    `json:"vpx10savailable,omitempty"`
	Vpx10stotal                  int    `json:"vpx10stotal,omitempty"`
	Vpx15000eavailable           int    `json:"vpx15000eavailable,omitempty"`
	Vpx15000etotal               int    `json:"vpx15000etotal,omitempty"`
	Vpx15000pavailable           int    `json:"vpx15000pavailable,omitempty"`
	Vpx15000ptotal               int    `json:"vpx15000ptotal,omitempty"`
	Vpx15000savailable           int    `json:"vpx15000savailable,omitempty"`
	Vpx15000stotal               int    `json:"vpx15000stotal,omitempty"`
	Vpx1pavailable               int    `json:"vpx1pavailable,omitempty"`
	Vpx1ptotal                   int    `json:"vpx1ptotal,omitempty"`
	Vpx1savailable               int    `json:"vpx1savailable,omitempty"`
	Vpx1stotal                   int    `json:"vpx1stotal,omitempty"`
	Vpx2000pavailable            int    `json:"vpx2000pavailable,omitempty"`
	Vpx2000ptotal                int    `json:"vpx2000ptotal,omitempty"`
	Vpx200eavailable             int    `json:"vpx200eavailable,omitempty"`
	Vpx200etotal                 int    `json:"vpx200etotal,omitempty"`
	Vpx200pavailable             int    `json:"vpx200pavailable,omitempty"`
	Vpx200ptotal                 int    `json:"vpx200ptotal,omitempty"`
	Vpx200savailable             int    `json:"vpx200savailable,omitempty"`
	Vpx200stotal                 int    `json:"vpx200stotal,omitempty"`
	Vpx25000eavailable           int    `json:"vpx25000eavailable,omitempty"`
	Vpx25000etotal               int    `json:"vpx25000etotal,omitempty"`
	Vpx25000pavailable           int    `json:"vpx25000pavailable,omitempty"`
	Vpx25000ptotal               int    `json:"vpx25000ptotal,omitempty"`
	Vpx25000savailable           int    `json:"vpx25000savailable,omitempty"`
	Vpx25000stotal               int    `json:"vpx25000stotal,omitempty"`
	Vpx25eavailable              int    `json:"vpx25eavailable,omitempty"`
	Vpx25etotal                  int    `json:"vpx25etotal,omitempty"`
	Vpx25pavailable              int    `json:"vpx25pavailable,omitempty"`
	Vpx25ptotal                  int    `json:"vpx25ptotal,omitempty"`
	Vpx25savailable              int    `json:"vpx25savailable,omitempty"`
	Vpx25stotal                  int    `json:"vpx25stotal,omitempty"`
	Vpx3000eavailable            int    `json:"vpx3000eavailable,omitempty"`
	Vpx3000etotal                int    `json:"vpx3000etotal,omitempty"`
	Vpx3000pavailable            int    `json:"vpx3000pavailable,omitempty"`
	Vpx3000ptotal                int    `json:"vpx3000ptotal,omitempty"`
	Vpx3000savailable            int    `json:"vpx3000savailable,omitempty"`
	Vpx3000stotal                int    `json:"vpx3000stotal,omitempty"`
	Vpx40000eavailable           int    `json:"vpx40000eavailable,omitempty"`
	Vpx40000etotal               int    `json:"vpx40000etotal,omitempty"`
	Vpx40000pavailable           int    `json:"vpx40000pavailable,omitempty"`
	Vpx40000ptotal               int    `json:"vpx40000ptotal,omitempty"`
	Vpx40000savailable           int    `json:"vpx40000savailable,omitempty"`
	Vpx40000stotal               int    `json:"vpx40000stotal,omitempty"`
	Vpx4000pavailable            int    `json:"vpx4000pavailable,omitempty"`
	Vpx4000ptotal                int    `json:"vpx4000ptotal,omitempty"`
	Vpx5000eavailable            int    `json:"vpx5000eavailable,omitempty"`
	Vpx5000etotal                int    `json:"vpx5000etotal,omitempty"`
	Vpx5000pavailable            int    `json:"vpx5000pavailable,omitempty"`
	Vpx5000ptotal                int    `json:"vpx5000ptotal,omitempty"`
	Vpx5000savailable            int    `json:"vpx5000savailable,omitempty"`
	Vpx5000stotal                int    `json:"vpx5000stotal,omitempty"`
	Vpx500eavailable             int    `json:"vpx500eavailable,omitempty"`
	Vpx500etotal                 int    `json:"vpx500etotal,omitempty"`
	Vpx500pavailable             int    `json:"vpx500pavailable,omitempty"`
	Vpx500ptotal                 int    `json:"vpx500ptotal,omitempty"`
	Vpx500savailable             int    `json:"vpx500savailable,omitempty"`
	Vpx500stotal                 int    `json:"vpx500stotal,omitempty"`
	Vpx50eavailable              int    `json:"vpx50eavailable,omitempty"`
	Vpx50etotal                  int    `json:"vpx50etotal,omitempty"`
	Vpx50pavailable              int    `json:"vpx50pavailable,omitempty"`
	Vpx50ptotal                  int    `json:"vpx50ptotal,omitempty"`
	Vpx50savailable              int    `json:"vpx50savailable,omitempty"`
	Vpx50stotal                  int    `json:"vpx50stotal,omitempty"`
	Vpx5pavailable               int    `json:"vpx5pavailable,omitempty"`
	Vpx5ptotal                   int    `json:"vpx5ptotal,omitempty"`
	Vpx5savailable               int    `json:"vpx5savailable,omitempty"`
	Vpx5stotal                   int    `json:"vpx5stotal,omitempty"`
	Vpx8000eavailable            int    `json:"vpx8000eavailable,omitempty"`
	Vpx8000etotal                int    `json:"vpx8000etotal,omitempty"`
	Vpx8000pavailable            int    `json:"vpx8000pavailable,omitempty"`
	Vpx8000ptotal                int    `json:"vpx8000ptotal,omitempty"`
	Vpx8000savailable            int    `json:"vpx8000savailable,omitempty"`
	Vpx8000stotal                int    `json:"vpx8000stotal,omitempty"`
}

type Nsconfigview struct {
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	State              string `json:"state,omitempty"`
}

type NstimerBinding struct {
	Name                          string        `json:"name,omitempty"`
	NstimerAutoscalepolicyBinding []interface{} `json:"nstimer_autoscalepolicy_binding,omitempty"`
}

type Nsdiameter struct {
	Count                  float64 `json:"__count,omitempty"`
	Identity               string  `json:"identity,omitempty"`
	Nextgenapiresource     string  `json:"_nextgenapiresource,omitempty"`
	Ownernode              int     `json:"ownernode,omitempty"`
	Realm                  string  `json:"realm,omitempty"`
	Serverclosepropagation string  `json:"serverclosepropagation,omitempty"`
}

type NstrafficdomainBridgegroupBinding struct {
	Bridgegroup int `json:"bridgegroup,omitempty"`
	Td          int `json:"td,omitempty"`
}

type Nsrollbackcmd struct {
	Filename           string `json:"filename,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Outtype            string `json:"outtype,omitempty"`
}

type Nsvpxparam struct {
	Cloudproductcode    string  `json:"cloudproductcode,omitempty"`
	Count               float64 `json:"__count,omitempty"`
	Cpuyield            string  `json:"cpuyield,omitempty"`
	Kvmvirtiomultiqueue string  `json:"kvmvirtiomultiqueue,omitempty"`
	Masterclockcpu1     string  `json:"masterclockcpu1,omitempty"`
	Memorystatus        string  `json:"memorystatus,omitempty"`
	Nextgenapiresource  string  `json:"_nextgenapiresource,omitempty"`
	Ownernode           int     `json:"ownernode,omitempty"`
	Technicalsupportpin string  `json:"technicalsupportpin,omitempty"`
	Vpxenvironment      string  `json:"vpxenvironment,omitempty"`
	Vpxoemcode          int     `json:"vpxoemcode,omitempty"`
}

type Nslicenseproxyserver struct {
	Count              float64 `json:"__count,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Port               int     `json:"port,omitempty"`
	Serverip           string  `json:"serverip,omitempty"`
	Servername         string  `json:"servername,omitempty"`
}

type NslimitidentifierNslimitsessionsBinding struct {
	Limitidentifier string `json:"limitidentifier,omitempty"`
}

type Nspartitionmac struct {
	Count              float64 `json:"__count,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Partitionmac       string  `json:"partitionmac,omitempty"`
	Partitionname      string  `json:"partitionname,omitempty"`
}

type Nsservicepath struct {
	Count              float64 `json:"__count,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Servicepathname    string  `json:"servicepathname,omitempty"`
}

type NspartitionBinding struct {
	NspartitionBridgegroupBinding []interface{} `json:"nspartition_bridgegroup_binding,omitempty"`
	NspartitionVlanBinding        []interface{} `json:"nspartition_vlan_binding,omitempty"`
	NspartitionVxlanBinding       []interface{} `json:"nspartition_vxlan_binding,omitempty"`
	Partitionname                 string        `json:"partitionname,omitempty"`
}

type NspartitionBridgegroupBinding struct {
	Bridgegroup   int    `json:"bridgegroup,omitempty"`
	Partitionname string `json:"partitionname,omitempty"`
}

type Nsconnectiontable struct {
	Adaptivetcpprofname    string   `json:"adaptivetcpprofname,omitempty"`
	Advwnd                 int      `json:"advwnd,omitempty"`
	Burstratecontrol       string   `json:"burstratecontrol,omitempty"`
	Bwestimate             int      `json:"bwestimate,omitempty"`
	Channelidnnm           int      `json:"channelidnnm,omitempty"`
	Congstate              string   `json:"congstate,omitempty"`
	Connid                 int      `json:"connid,omitempty"`
	Connproperties         []string `json:"connproperties,omitempty"`
	Count                  float64  `json:"__count,omitempty"`
	Cqabifavg              int      `json:"cqabifavg,omitempty"`
	Cqaccl                 int      `json:"cqaccl,omitempty"`
	Cqacsq                 int      `json:"cqacsq,omitempty"`
	Cqaiai1mspct           int      `json:"cqaiai1mspct,omitempty"`
	Cqaiai2mspct           int      `json:"cqaiai2mspct,omitempty"`
	Cqaiaiavg              int      `json:"cqaiaiavg,omitempty"`
	Cqaiaisamples          int      `json:"cqaiaisamples,omitempty"`
	Cqaisiavg              int      `json:"cqaisiavg,omitempty"`
	Cqaloaddelayavg        int      `json:"cqaloaddelayavg,omitempty"`
	Cqanetclass            string   `json:"cqanetclass,omitempty"`
	Cqanoisedelayavg       int      `json:"cqanoisedelayavg,omitempty"`
	Cqarcvwndavg           int      `json:"cqarcvwndavg,omitempty"`
	Cqarcvwndmin           int      `json:"cqarcvwndmin,omitempty"`
	Cqaretxcong            int      `json:"cqaretxcong,omitempty"`
	Cqaretxcorr            int      `json:"cqaretxcorr,omitempty"`
	Cqaretxpackets         int      `json:"cqaretxpackets,omitempty"`
	Cqarttavg              int      `json:"cqarttavg,omitempty"`
	Cqarttmax              int      `json:"cqarttmax,omitempty"`
	Cqarttmin              int      `json:"cqarttmin,omitempty"`
	Cqasamples             int      `json:"cqasamples,omitempty"`
	Cqathruputavg          int      `json:"cqathruputavg,omitempty"`
	Creditsinbytes         int      `json:"creditsinbytes,omitempty"`
	Destip                 string   `json:"destip,omitempty"`
	Destport               int      `json:"destport,omitempty"`
	Entityname             string   `json:"entityname,omitempty"`
	Filterexpression       string   `json:"filterexpression,omitempty"`
	Filtername             bool     `json:"filtername,omitempty"`
	Flavor                 string   `json:"flavor,omitempty"`
	Httpendseq             int      `json:"httpendseq,omitempty"`
	Httprequest            string   `json:"httprequest,omitempty"`
	Httpreqver             string   `json:"httpreqver,omitempty"`
	Httprspcode            int      `json:"httprspcode,omitempty"`
	Httpstate              string   `json:"httpstate,omitempty"`
	Idletime               int      `json:"idletime,omitempty"`
	Irs                    int      `json:"irs,omitempty"`
	Iss                    int      `json:"iss,omitempty"`
	Link                   bool     `json:"link,omitempty"`
	Linkburstratecontrol   string   `json:"linkburstratecontrol,omitempty"`
	Linkbwestimate         int      `json:"linkbwestimate,omitempty"`
	Linkcongstate          string   `json:"linkcongstate,omitempty"`
	Linkconnid             int      `json:"linkconnid,omitempty"`
	Linkcredits            int      `json:"linkcredits,omitempty"`
	Linkdestip             string   `json:"linkdestip,omitempty"`
	Linkdestport           int      `json:"linkdestport,omitempty"`
	Linkentityname         string   `json:"linkentityname,omitempty"`
	Linkflavor             string   `json:"linkflavor,omitempty"`
	Linkidletime           int      `json:"linkidletime,omitempty"`
	Linkmaxrcvbuf          int      `json:"linkmaxrcvbuf,omitempty"`
	Linkmaxsndbuf          int      `json:"linkmaxsndbuf,omitempty"`
	Linkname               string   `json:"linkname,omitempty"`
	Linknsbretxq           int      `json:"linknsbretxq,omitempty"`
	Linknsbtcpwaitq        int      `json:"linknsbtcpwaitq,omitempty"`
	Linknswsvalue          int      `json:"linknswsvalue,omitempty"`
	Linkoptionflag         []string `json:"linkoptionflag,omitempty"`
	Linkpeerwsvalue        int      `json:"linkpeerwsvalue,omitempty"`
	Linkrateinbytes        int      `json:"linkrateinbytes,omitempty"`
	Linkrateschedulerqueue int      `json:"linkrateschedulerqueue,omitempty"`
	Linkrealtimertt        int      `json:"linkrealtimertt,omitempty"`
	Linkrttmin             int      `json:"linkrttmin,omitempty"`
	Linkrxqsize            int      `json:"linkrxqsize,omitempty"`
	Linksackblocks         int      `json:"linksackblocks,omitempty"`
	Linkservicetype        string   `json:"linkservicetype,omitempty"`
	Linksndbuf             int      `json:"linksndbuf,omitempty"`
	Linksndrecoverle       int      `json:"linksndrecoverle,omitempty"`
	Linksourceip           string   `json:"linksourceip,omitempty"`
	Linksourceport         int      `json:"linksourceport,omitempty"`
	Linkstate              string   `json:"linkstate,omitempty"`
	Linktcpmode            string   `json:"linktcpmode,omitempty"`
	Linktxqsize            int      `json:"linktxqsize,omitempty"`
	Listen                 bool     `json:"listen,omitempty"`
	Maxack                 int      `json:"maxack,omitempty"`
	Maxrcvbuf              int      `json:"maxrcvbuf,omitempty"`
	Maxsndbuf              int      `json:"maxsndbuf,omitempty"`
	Msgversionnnm          int      `json:"msgversionnnm,omitempty"`
	Mss                    int      `json:"mss,omitempty"`
	Name                   string   `json:"name,omitempty"`
	Nextgenapiresource     string   `json:"_nextgenapiresource,omitempty"`
	Nodeid                 int      `json:"nodeid,omitempty"`
	Nsbretxq               int      `json:"nsbretxq,omitempty"`
	Nsbtcpwaitq            int      `json:"nsbtcpwaitq,omitempty"`
	Nswsvalue              int      `json:"nswsvalue,omitempty"`
	Optionflags            []string `json:"optionflags,omitempty"`
	Outoforderblocks       int      `json:"outoforderblocks,omitempty"`
	Outoforderbytes        int      `json:"outoforderbytes,omitempty"`
	Outoforderflushedcount int      `json:"outoforderflushedcount,omitempty"`
	Outoforderpkts         int      `json:"outoforderpkts,omitempty"`
	Peerwsvalue            int      `json:"peerwsvalue,omitempty"`
	Priority               string   `json:"priority,omitempty"`
	Rateinbytes            int      `json:"rateinbytes,omitempty"`
	Rateschedulerqueue     int      `json:"rateschedulerqueue,omitempty"`
	Rcvnxt                 int      `json:"rcvnxt,omitempty"`
	Rcvwnd                 int      `json:"rcvwnd,omitempty"`
	Realtimertt            int      `json:"realtimertt,omitempty"`
	Retxretrycnt           int      `json:"retxretrycnt,omitempty"`
	Rttmin                 int      `json:"rttmin,omitempty"`
	Rttsmoothed            int      `json:"rttsmoothed,omitempty"`
	Rttvariance            int      `json:"rttvariance,omitempty"`
	Rxqsize                int      `json:"rxqsize,omitempty"`
	Sackblocks             int      `json:"sackblocks,omitempty"`
	Sndbuf                 int      `json:"sndbuf,omitempty"`
	Sndcwnd                int      `json:"sndcwnd,omitempty"`
	Sndnxt                 int      `json:"sndnxt,omitempty"`
	Sndrecoverle           int      `json:"sndrecoverle,omitempty"`
	Sndunack               int      `json:"sndunack,omitempty"`
	Sourceip               string   `json:"sourceip,omitempty"`
	Sourcenodeidnnm        int      `json:"sourcenodeidnnm,omitempty"`
	Sourceport             int      `json:"sourceport,omitempty"`
	State                  string   `json:"state,omitempty"`
	Svctype                string   `json:"svctype,omitempty"`
	Targetnodeidnnm        int      `json:"targetnodeidnnm,omitempty"`
	Tcpmode                string   `json:"tcpmode,omitempty"`
	Td                     int      `json:"td,omitempty"`
	Trcount                int      `json:"trcount,omitempty"`
	Txqsize                int      `json:"txqsize,omitempty"`
}

type NsservicepathNsservicefunctionBinding struct {
	Index           int    `json:"index,omitempty"`
	Servicefunction string `json:"servicefunction,omitempty"`
	Servicepathname string `json:"servicepathname,omitempty"`
}

type Nsparam struct {
	Advancedanalyticsstats    string        `json:"advancedanalyticsstats,omitempty"`
	Aftpallowrandomsourceport string        `json:"aftpallowrandomsourceport,omitempty"`
	Autoscaleoption           int           `json:"autoscaleoption,omitempty"`
	Cip                       string        `json:"cip,omitempty"`
	Cipheader                 string        `json:"cipheader,omitempty"`
	Cookieversion             string        `json:"cookieversion,omitempty"`
	Crportrange               string        `json:"crportrange,omitempty"`
	Exclusivequotamaxclient   int           `json:"exclusivequotamaxclient,omitempty"`
	Exclusivequotaspillover   int           `json:"exclusivequotaspillover,omitempty"`
	Ftpportrange              string        `json:"ftpportrange,omitempty"`
	Grantquotamaxclient       int           `json:"grantquotamaxclient,omitempty"`
	Grantquotaspillover       int           `json:"grantquotaspillover,omitempty"`
	Httpport                  []interface{} `json:"httpport,omitempty"`
	Icaports                  []interface{} `json:"icaports,omitempty"`
	Internaluserlogin         string        `json:"internaluserlogin,omitempty"`
	Ipttl                     int           `json:"ipttl,omitempty"`
	Maxconn                   int           `json:"maxconn,omitempty"`
	Maxreq                    int           `json:"maxreq,omitempty"`
	Mgmthttpport              int           `json:"mgmthttpport,omitempty"`
	Mgmthttpsport             int           `json:"mgmthttpsport,omitempty"`
	Nextgenapiresource        string        `json:"_nextgenapiresource,omitempty"`
	Pmtumin                   int           `json:"pmtumin,omitempty"`
	Pmtutimeout               int           `json:"pmtutimeout,omitempty"`
	Proxyprotocol             string        `json:"proxyprotocol,omitempty"`
	Securecookie              string        `json:"securecookie,omitempty"`
	Secureicaports            []interface{} `json:"secureicaports,omitempty"`
	Servicepathingressvlan    int           `json:"servicepathingressvlan,omitempty"`
	Tcpcip                    string        `json:"tcpcip,omitempty"`
	Timezone                  string        `json:"timezone,omitempty"`
	Useproxyport              string        `json:"useproxyport,omitempty"`
}

type NsextensionExtensionfunctionBinding struct {
	Activeextensionfunction         int      `json:"activeextensionfunction,omitempty"`
	Extensionfuncdescription        string   `json:"extensionfuncdescription,omitempty"`
	Extensionfunctionallparams      []string `json:"extensionfunctionallparams,omitempty"`
	Extensionfunctionallparamscount int      `json:"extensionfunctionallparamscount,omitempty"`
	Extensionfunctionargcount       int      `json:"extensionfunctionargcount,omitempty"`
	Extensionfunctionargtype        []string `json:"extensionfunctionargtype,omitempty"`
	Extensionfunctionclasses        []string `json:"extensionfunctionclasses,omitempty"`
	Extensionfunctionclassescount   int      `json:"extensionfunctionclassescount,omitempty"`
	Extensionfunctionclasstype      string   `json:"extensionfunctionclasstype,omitempty"`
	Extensionfunctionlinenumber     int      `json:"extensionfunctionlinenumber,omitempty"`
	Extensionfunctionname           string   `json:"extensionfunctionname,omitempty"`
	Extensionfunctionreturntype     string   `json:"extensionfunctionreturntype,omitempty"`
	Name                            string   `json:"name,omitempty"`
}

type Nsversion struct {
	Installedversion   bool   `json:"installedversion,omitempty"`
	Mode               int    `json:"mode,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Version            string `json:"version,omitempty"`
}

type Nshttpparam struct {
	Builtin                   []string `json:"builtin,omitempty"`
	Conmultiplex              string   `json:"conmultiplex,omitempty"`
	Dropinvalreqs             string   `json:"dropinvalreqs,omitempty"`
	Feature                   string   `json:"feature,omitempty"`
	Http2serverside           string   `json:"http2serverside,omitempty"`
	Ignoreconnectcodingscheme string   `json:"ignoreconnectcodingscheme,omitempty"`
	Insnssrvrhdr              string   `json:"insnssrvrhdr,omitempty"`
	Logerrresp                string   `json:"logerrresp,omitempty"`
	Markconnreqinval          string   `json:"markconnreqinval,omitempty"`
	Markhttp09inval           string   `json:"markhttp09inval,omitempty"`
	Maxreusepool              int      `json:"maxreusepool,omitempty"`
	Nextgenapiresource        string   `json:"_nextgenapiresource,omitempty"`
	Nssrvrhdr                 string   `json:"nssrvrhdr,omitempty"`
}

type Nsip6 struct {
	Advertiseondefaultpartition string   `json:"advertiseondefaultpartition,omitempty"`
	Count                       float64  `json:"__count,omitempty"`
	Curstate                    string   `json:"curstate,omitempty"`
	Decrementhoplimit           string   `json:"decrementhoplimit,omitempty"`
	Dynamicrouting              string   `json:"dynamicrouting,omitempty"`
	Ftp                         string   `json:"ftp,omitempty"`
	Gui                         string   `json:"gui,omitempty"`
	Hostroute                   string   `json:"hostroute,omitempty"`
	Icmp                        string   `json:"icmp,omitempty"`
	Icmpresponse                string   `json:"icmpresponse,omitempty"`
	Ip6hostrtgw                 string   `json:"ip6hostrtgw,omitempty"`
	Iptype                      []string `json:"iptype,omitempty"`
	Ipv6address                 string   `json:"ipv6address,omitempty"`
	MapField                    string   `json:"map,omitempty"`
	Metric                      int      `json:"metric,omitempty"`
	Mgmtaccess                  string   `json:"mgmtaccess,omitempty"`
	Mptcpadvertise              string   `json:"mptcpadvertise,omitempty"`
	Nd                          string   `json:"nd,omitempty"`
	Ndowner                     int      `json:"ndowner,omitempty"`
	Networkroute                string   `json:"networkroute,omitempty"`
	Nextgenapiresource          string   `json:"_nextgenapiresource,omitempty"`
	Operationalndowner          int      `json:"operationalndowner,omitempty"`
	Ospf6lsatype                string   `json:"ospf6lsatype,omitempty"`
	Ospfarea                    int      `json:"ospfarea,omitempty"`
	Ownerdownresponse           string   `json:"ownerdownresponse,omitempty"`
	Ownernode                   int      `json:"ownernode,omitempty"`
	Restrictaccess              string   `json:"restrictaccess,omitempty"`
	Scope                       string   `json:"scope,omitempty"`
	Snmp                        string   `json:"snmp,omitempty"`
	Ssh                         string   `json:"ssh,omitempty"`
	State                       string   `json:"state,omitempty"`
	Systemtype                  string   `json:"systemtype,omitempty"`
	Tag                         int      `json:"tag,omitempty"`
	Td                          int      `json:"td,omitempty"`
	Telnet                      string   `json:"telnet,omitempty"`
	TypeField                   string   `json:"type,omitempty"`
	Viprtadv2bsd                bool     `json:"viprtadv2bsd,omitempty"`
	Vipvsercount                int      `json:"vipvsercount,omitempty"`
	Vipvserdowncount            int      `json:"vipvserdowncount,omitempty"`
	Vlan                        int      `json:"vlan,omitempty"`
	Vrid6                       int      `json:"vrid6,omitempty"`
	Vserver                     string   `json:"vserver,omitempty"`
	Vserverrhilevel             string   `json:"vserverrhilevel,omitempty"`
}

type NspartitionVxlanBinding struct {
	Partitionname string `json:"partitionname,omitempty"`
	Vxlan         int    `json:"vxlan,omitempty"`
}

type Nsip struct {
	Advertiseondefaultpartition string   `json:"advertiseondefaultpartition,omitempty"`
	Arp                         string   `json:"arp,omitempty"`
	Arpowner                    int      `json:"arpowner,omitempty"`
	Arpresponse                 string   `json:"arpresponse,omitempty"`
	Bgp                         string   `json:"bgp,omitempty"`
	Count                       float64  `json:"__count,omitempty"`
	Decrementttl                string   `json:"decrementttl,omitempty"`
	Dynamicrouting              string   `json:"dynamicrouting,omitempty"`
	Flags                       int      `json:"flags,omitempty"`
	Freeports                   int      `json:"freeports,omitempty"`
	Ftp                         string   `json:"ftp,omitempty"`
	Gui                         string   `json:"gui,omitempty"`
	Hostroute                   string   `json:"hostroute,omitempty"`
	Hostrtgw                    string   `json:"hostrtgw,omitempty"`
	Hostrtgwact                 string   `json:"hostrtgwact,omitempty"`
	Icmp                        string   `json:"icmp,omitempty"`
	Icmpresponse                string   `json:"icmpresponse,omitempty"`
	Ipaddress                   string   `json:"ipaddress,omitempty"`
	Iptype                      []string `json:"iptype,omitempty"`
	Metric                      int      `json:"metric,omitempty"`
	Mgmtaccess                  string   `json:"mgmtaccess,omitempty"`
	Mptcpadvertise              string   `json:"mptcpadvertise,omitempty"`
	Netmask                     string   `json:"netmask,omitempty"`
	Networkroute                string   `json:"networkroute,omitempty"`
	Nextgenapiresource          string   `json:"_nextgenapiresource,omitempty"`
	Operationalarpowner         int      `json:"operationalarpowner,omitempty"`
	Ospf                        string   `json:"ospf,omitempty"`
	Ospfarea                    int      `json:"ospfarea,omitempty"`
	Ospfareaval                 int      `json:"ospfareaval,omitempty"`
	Ospflsatype                 string   `json:"ospflsatype,omitempty"`
	Ownerdownresponse           string   `json:"ownerdownresponse,omitempty"`
	Ownernode                   int      `json:"ownernode,omitempty"`
	Restrictaccess              string   `json:"restrictaccess,omitempty"`
	Rip                         string   `json:"rip,omitempty"`
	Snmp                        string   `json:"snmp,omitempty"`
	Ssh                         string   `json:"ssh,omitempty"`
	State                       string   `json:"state,omitempty"`
	Tag                         int      `json:"tag,omitempty"`
	Td                          int      `json:"td,omitempty"`
	Telnet                      string   `json:"telnet,omitempty"`
	TypeField                   string   `json:"type,omitempty"`
	Viprtadv2bsd                bool     `json:"viprtadv2bsd,omitempty"`
	Vipvsercount                int      `json:"vipvsercount,omitempty"`
	Vipvserdowncount            int      `json:"vipvserdowncount,omitempty"`
	Vipvsrvrrhiactivecount      int      `json:"vipvsrvrrhiactivecount,omitempty"`
	Vipvsrvrrhiactiveupcount    int      `json:"vipvsrvrrhiactiveupcount,omitempty"`
	Vrid                        int      `json:"vrid,omitempty"`
	Vserver                     string   `json:"vserver,omitempty"`
	Vserverrhilevel             string   `json:"vserverrhilevel,omitempty"`
}

type Nsextension struct {
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Detail             string  `json:"detail,omitempty"`
	Functionhaltcount  int     `json:"functionhaltcount,omitempty"`
	Functionhits       int     `json:"functionhits,omitempty"`
	Functionundefhits  int     `json:"functionundefhits,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Overwrite          bool    `json:"overwrite,omitempty"`
	Src                string  `json:"src,omitempty"`
	Trace              string  `json:"trace,omitempty"`
	Tracefunctions     string  `json:"tracefunctions,omitempty"`
	Tracevariables     string  `json:"tracevariables,omitempty"`
	TypeField          string  `json:"type,omitempty"`
}

type Shutdown struct {
}

type Nscapacity struct {
	Actualbandwidth    int    `json:"actualbandwidth,omitempty"`
	Bandwidth          int    `json:"bandwidth,omitempty"`
	Daystoexpiration   int    `json:"daystoexpiration,omitempty"`
	Edition            string `json:"edition,omitempty"`
	Instancecount      int    `json:"instancecount,omitempty"`
	Maxbandwidth       int    `json:"maxbandwidth,omitempty"`
	Maxvcpucount       int    `json:"maxvcpucount,omitempty"`
	Minbandwidth       int    `json:"minbandwidth,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Nodeid             int    `json:"nodeid,omitempty"`
	Password           string `json:"password,omitempty"`
	Platform           string `json:"platform,omitempty"`
	Unit               string `json:"unit,omitempty"`
	Username           string `json:"username,omitempty"`
	Vcpu               bool   `json:"vcpu,omitempty"`
	Vcpucount          int    `json:"vcpucount,omitempty"`
}

type Nssavedconfig struct {
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Textblob           string `json:"textblob,omitempty"`
}

type Nsrpcnode struct {
	Count              float64 `json:"__count,omitempty"`
	Ipaddress          string  `json:"ipaddress,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Password           string  `json:"password,omitempty"`
	Secure             string  `json:"secure,omitempty"`
	Srcip              string  `json:"srcip,omitempty"`
	Validatecert       string  `json:"validatecert,omitempty"`
}

type Nsassignment struct {
	Add                string  `json:"Add,omitempty"`
	Append             string  `json:"append,omitempty"`
	Clear              bool    `json:"clear,omitempty"`
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Hits               int     `json:"hits,omitempty"`
	Name               string  `json:"name,omitempty"`
	Newname            string  `json:"newname,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Referencecount     int     `json:"referencecount,omitempty"`
	Set                string  `json:"set,omitempty"`
	Sub                string  `json:"sub,omitempty"`
	Undefhits          int     `json:"undefhits,omitempty"`
	Variable           string  `json:"variable,omitempty"`
}

type Nstcpbufparam struct {
	Builtin            []string `json:"builtin,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Memlimit           int      `json:"memlimit,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Size               int      `json:"size,omitempty"`
}

type Nsratecontrol struct {
	Icmpthreshold      int    `json:"icmpthreshold,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Tcprstthreshold    int    `json:"tcprstthreshold,omitempty"`
	Tcpthreshold       int    `json:"tcpthreshold,omitempty"`
	Udpthreshold       int    `json:"udpthreshold,omitempty"`
}

type Nslimitsessions struct {
	Count              float64       `json:"__count,omitempty"`
	Detail             bool          `json:"detail,omitempty"`
	Drop               int           `json:"drop,omitempty"`
	Flag               int           `json:"flag,omitempty"`
	Flags              int           `json:"flags,omitempty"`
	Hits               int           `json:"hits,omitempty"`
	Limitidentifier    string        `json:"limitidentifier,omitempty"`
	Maxbandwidth       int           `json:"maxbandwidth,omitempty"`
	Name               string        `json:"name,omitempty"`
	Nextgenapiresource string        `json:"_nextgenapiresource,omitempty"`
	Number             []interface{} `json:"number,omitempty"`
	Referencecount     int           `json:"referencecount,omitempty"`
	Selectoripv61      string        `json:"selectoripv61,omitempty"`
	Selectoripv62      string        `json:"selectoripv62,omitempty"`
	Timeout            int           `json:"timeout,omitempty"`
	Unit               int           `json:"unit,omitempty"`
}

type NstrafficdomainVlanBinding struct {
	Td   int `json:"td,omitempty"`
	Vlan int `json:"vlan,omitempty"`
}

type Nsspparams struct {
	Basethreshold      int           `json:"basethreshold,omitempty"`
	Builtin            []string      `json:"builtin,omitempty"`
	Feature            string        `json:"feature,omitempty"`
	Nextgenapiresource string        `json:"_nextgenapiresource,omitempty"`
	Table0             []interface{} `json:"table0,omitempty"`
	Throttle           string        `json:"throttle,omitempty"`
}

type Nsacls6 struct {
	TypeField string `json:"type,omitempty"`
}

type Nsrunningconfig struct {
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Response           string `json:"response,omitempty"`
	Withdefaults       bool   `json:"withdefaults,omitempty"`
}

type Reboot struct {
	Warm bool `json:"warm,omitempty"`
}

type Nstimer struct {
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Interval           int     `json:"interval,omitempty"`
	Name               string  `json:"name,omitempty"`
	Newname            string  `json:"newname,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Unit               string  `json:"unit,omitempty"`
}

type Nscentralmanagementserver struct {
	Activationcode             string  `json:"activationcode,omitempty"`
	Adcpassword                string  `json:"adcpassword,omitempty"`
	Adcusername                string  `json:"adcusername,omitempty"`
	Admserviceconnectionstatus string  `json:"admserviceconnectionstatus,omitempty"`
	Admserviceenvironment      string  `json:"admserviceenvironment,omitempty"`
	Count                      float64 `json:"__count,omitempty"`
	Customerid                 string  `json:"customerid,omitempty"`
	Deviceprofilename          string  `json:"deviceprofilename,omitempty"`
	Instanceid                 string  `json:"instanceid,omitempty"`
	Ipaddress                  string  `json:"ipaddress,omitempty"`
	Nextgenapiresource         string  `json:"_nextgenapiresource,omitempty"`
	Password                   string  `json:"password,omitempty"`
	Servername                 string  `json:"servername,omitempty"`
	TypeField                  string  `json:"type,omitempty"`
	Username                   string  `json:"username,omitempty"`
	Validatecert               string  `json:"validatecert,omitempty"`
}

type Nsencryptionparams struct {
	Keyvalue           string `json:"keyvalue,omitempty"`
	Method             string `json:"method,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nssourceroutecachetable struct {
	Count              float64 `json:"__count,omitempty"`
	InterfaceField     string  `json:"Interface,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Sourceip           string  `json:"sourceip,omitempty"`
	Sourcemac          string  `json:"sourcemac,omitempty"`
	Vlan               int     `json:"vlan,omitempty"`
}

type Nstimezone struct {
	Count              float64 `json:"__count,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Value              string  `json:"value,omitempty"`
}

type Nsacl6 struct {
	Acl6action         string   `json:"acl6action,omitempty"`
	Acl6name           string   `json:"acl6name,omitempty"`
	Aclaction          string   `json:"aclaction,omitempty"`
	Aclassociate       []string `json:"aclassociate,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Destipop           string   `json:"destipop,omitempty"`
	Destipv6           bool     `json:"destipv6,omitempty"`
	Destipv6val        string   `json:"destipv6val,omitempty"`
	Destport           bool     `json:"destport,omitempty"`
	Destportop         string   `json:"destportop,omitempty"`
	Destportval        string   `json:"destportval,omitempty"`
	Dfdhash            string   `json:"dfdhash,omitempty"`
	Dfdprefix          int      `json:"dfdprefix,omitempty"`
	Established        bool     `json:"established,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	Icmpcode           int      `json:"icmpcode,omitempty"`
	Icmptype           int      `json:"icmptype,omitempty"`
	InterfaceField     string   `json:"Interface,omitempty"`
	Kernelstate        string   `json:"kernelstate,omitempty"`
	Logstate           string   `json:"logstate,omitempty"`
	Newname            string   `json:"newname,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Nodeid             int      `json:"nodeid,omitempty"`
	Priority           int      `json:"priority,omitempty"`
	Protocol           string   `json:"protocol,omitempty"`
	Protocolnumber     int      `json:"protocolnumber,omitempty"`
	Ratelimit          int      `json:"ratelimit,omitempty"`
	Srcipop            string   `json:"srcipop,omitempty"`
	Srcipv6            bool     `json:"srcipv6,omitempty"`
	Srcipv6val         string   `json:"srcipv6val,omitempty"`
	Srcmac             string   `json:"srcmac,omitempty"`
	Srcmacmask         string   `json:"srcmacmask,omitempty"`
	Srcport            bool     `json:"srcport,omitempty"`
	Srcportop          string   `json:"srcportop,omitempty"`
	Srcportval         string   `json:"srcportval,omitempty"`
	State              string   `json:"state,omitempty"`
	Stateful           string   `json:"stateful,omitempty"`
	Td                 int      `json:"td,omitempty"`
	Ttl                int      `json:"ttl,omitempty"`
	TypeField          string   `json:"type,omitempty"`
	Vlan               int      `json:"vlan,omitempty"`
	Vxlan              int      `json:"vxlan,omitempty"`
}

type Nsappflowparam struct {
	Clienttrafficonly  string `json:"clienttrafficonly,omitempty"`
	Httpcookie         string `json:"httpcookie,omitempty"`
	Httphost           string `json:"httphost,omitempty"`
	Httpmethod         string `json:"httpmethod,omitempty"`
	Httpreferer        string `json:"httpreferer,omitempty"`
	Httpurl            string `json:"httpurl,omitempty"`
	Httpuseragent      string `json:"httpuseragent,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Templaterefresh    int    `json:"templaterefresh,omitempty"`
	Udppmtu            int    `json:"udppmtu,omitempty"`
}

type Nsvariable struct {
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Expires            int     `json:"expires,omitempty"`
	Iffull             string  `json:"iffull,omitempty"`
	Ifnovalue          string  `json:"ifnovalue,omitempty"`
	Ifvaluetoobig      string  `json:"ifvaluetoobig,omitempty"`
	Init               string  `json:"init,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Referencecount     int     `json:"referencecount,omitempty"`
	Scope              string  `json:"scope,omitempty"`
	TypeField          string  `json:"type,omitempty"`
}

type Nsconsoleloginprompt struct {
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Promptstring       string `json:"promptstring,omitempty"`
}

type Nshostname struct {
	Count              float64 `json:"__count,omitempty"`
	Hostname           string  `json:"hostname,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Ownernode          int     `json:"ownernode,omitempty"`
}

type Nsjob struct {
	Count              float64 `json:"__count,omitempty"`
	Errorcode          int     `json:"errorcode,omitempty"`
	Id                 int     `json:"id,omitempty"`
	Message            string  `json:"message,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Progress           string  `json:"progress,omitempty"`
	Response           string  `json:"response,omitempty"`
	Status             string  `json:"status,omitempty"`
	Timeelapsed        int     `json:"timeelapsed,omitempty"`
}

type Nsappflowcollector struct {
	Count              float64 `json:"__count,omitempty"`
	Ipaddress          string  `json:"ipaddress,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Port               int     `json:"port,omitempty"`
}

type Nslimitselector struct {
	Count              float64  `json:"__count,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Rule               []string `json:"rule,omitempty"`
	Selectorname       string   `json:"selectorname,omitempty"`
}

type Nstimeout struct {
	Anyclient          int    `json:"anyclient,omitempty"`
	Anyserver          int    `json:"anyserver,omitempty"`
	Anytcpclient       int    `json:"anytcpclient,omitempty"`
	Anytcpserver       int    `json:"anytcpserver,omitempty"`
	Client             int    `json:"client,omitempty"`
	Halfclose          int    `json:"halfclose,omitempty"`
	Httpclient         int    `json:"httpclient,omitempty"`
	Httpserver         int    `json:"httpserver,omitempty"`
	Newconnidletimeout int    `json:"newconnidletimeout,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Nontcpzombie       int    `json:"nontcpzombie,omitempty"`
	Reducedfintimeout  int    `json:"reducedfintimeout,omitempty"`
	Reducedrsttimeout  int    `json:"reducedrsttimeout,omitempty"`
	Server             int    `json:"server,omitempty"`
	Tcpclient          int    `json:"tcpclient,omitempty"`
	Tcpserver          int    `json:"tcpserver,omitempty"`
	Zombie             int    `json:"zombie,omitempty"`
}

type Nstcpparam struct {
	Ackonpush                           string   `json:"ackonpush,omitempty"`
	Autosyncookietimeout                int      `json:"autosyncookietimeout,omitempty"`
	Builtin                             []string `json:"builtin,omitempty"`
	Compacttcpoptionnoop                string   `json:"compacttcpoptionnoop,omitempty"`
	Connflushifnomem                    string   `json:"connflushifnomem,omitempty"`
	Connflushthres                      int      `json:"connflushthres,omitempty"`
	Delayedack                          int      `json:"delayedack,omitempty"`
	Delinkclientserveronrst             string   `json:"delinkclientserveronrst,omitempty"`
	Downstaterst                        string   `json:"downstaterst,omitempty"`
	Enhancedisngeneration               string   `json:"enhancedisngeneration,omitempty"`
	Feature                             string   `json:"feature,omitempty"`
	Initialcwnd                         int      `json:"initialcwnd,omitempty"`
	Kaprobeupdatelastactivity           string   `json:"kaprobeupdatelastactivity,omitempty"`
	Learnvsvrmss                        string   `json:"learnvsvrmss,omitempty"`
	Limitedpersist                      string   `json:"limitedpersist,omitempty"`
	Maxburst                            int      `json:"maxburst,omitempty"`
	Maxdynserverprobes                  int      `json:"maxdynserverprobes,omitempty"`
	Maxpktpermss                        int      `json:"maxpktpermss,omitempty"`
	Maxsynackretx                       int      `json:"maxsynackretx,omitempty"`
	Maxsynhold                          int      `json:"maxsynhold,omitempty"`
	Maxsynholdperprobe                  int      `json:"maxsynholdperprobe,omitempty"`
	Maxtimewaitconn                     int      `json:"maxtimewaitconn,omitempty"`
	Minrto                              int      `json:"minrto,omitempty"`
	Mptcpchecksum                       string   `json:"mptcpchecksum,omitempty"`
	Mptcpclosemptcpsessiononlastsfclose string   `json:"mptcpclosemptcpsessiononlastsfclose,omitempty"`
	Mptcpconcloseonpassivesf            string   `json:"mptcpconcloseonpassivesf,omitempty"`
	Mptcpfastcloseoption                string   `json:"mptcpfastcloseoption,omitempty"`
	Mptcpimmediatesfcloseonfin          string   `json:"mptcpimmediatesfcloseonfin,omitempty"`
	Mptcpmaxpendingsf                   int      `json:"mptcpmaxpendingsf,omitempty"`
	Mptcpmaxsf                          int      `json:"mptcpmaxsf,omitempty"`
	Mptcppendingjointhreshold           int      `json:"mptcppendingjointhreshold,omitempty"`
	Mptcpreliableaddaddr                string   `json:"mptcpreliableaddaddr,omitempty"`
	Mptcprtostoswitchsf                 int      `json:"mptcprtostoswitchsf,omitempty"`
	Mptcpsendsfresetoption              string   `json:"mptcpsendsfresetoption,omitempty"`
	Mptcpsfreplacetimeout               int      `json:"mptcpsfreplacetimeout,omitempty"`
	Mptcpsftimeout                      int      `json:"mptcpsftimeout,omitempty"`
	Mptcpusebackupondss                 string   `json:"mptcpusebackupondss,omitempty"`
	Msslearndelay                       int      `json:"msslearndelay,omitempty"`
	Msslearninterval                    int      `json:"msslearninterval,omitempty"`
	Nagle                               string   `json:"nagle,omitempty"`
	Nextgenapiresource                  string   `json:"_nextgenapiresource,omitempty"`
	Oooqsize                            int      `json:"oooqsize,omitempty"`
	Pktperretx                          int      `json:"pktperretx,omitempty"`
	Recvbuffsize                        int      `json:"recvbuffsize,omitempty"`
	Rfc5961chlgacklimit                 int      `json:"rfc5961chlgacklimit,omitempty"`
	Sack                                string   `json:"sack,omitempty"`
	Sendresetreasoncode                 string   `json:"sendresetreasoncode,omitempty"`
	Slowstartincr                       int      `json:"slowstartincr,omitempty"`
	Synattackdetection                  string   `json:"synattackdetection,omitempty"`
	Synholdfastgiveup                   int      `json:"synholdfastgiveup,omitempty"`
	Tcpfastopencookietimeout            int      `json:"tcpfastopencookietimeout,omitempty"`
	Tcpfintimeout                       int      `json:"tcpfintimeout,omitempty"`
	Tcpmaxretries                       int      `json:"tcpmaxretries,omitempty"`
	Ws                                  string   `json:"ws,omitempty"`
	Wsval                               int      `json:"wsval,omitempty"`
}

type Nsservicefunction struct {
	Count               float64 `json:"__count,omitempty"`
	Ingressvlan         int     `json:"ingressvlan,omitempty"`
	Nextgenapiresource  string  `json:"_nextgenapiresource,omitempty"`
	Servicefunctionname string  `json:"servicefunctionname,omitempty"`
}

type Nschannelparam struct {
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Vfautorecover      string `json:"vfautorecover,omitempty"`
}

type Nsencryptionkey struct {
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Iv                 string  `json:"iv,omitempty"`
	Keyvalue           string  `json:"keyvalue,omitempty"`
	Method             string  `json:"method,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Padding            string  `json:"padding,omitempty"`
}

type Nskeymanagerproxy struct {
	Count              float64 `json:"__count,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Nodeid             int     `json:"nodeid,omitempty"`
	Port               int     `json:"port,omitempty"`
	Serverip           string  `json:"serverip,omitempty"`
	Servername         string  `json:"servername,omitempty"`
	Status             int     `json:"status,omitempty"`
}

type Nsicapprofile struct {
	Allow204            string  `json:"allow204,omitempty"`
	Connectionkeepalive string  `json:"connectionkeepalive,omitempty"`
	Count               float64 `json:"__count,omitempty"`
	Hostheader          string  `json:"hostheader,omitempty"`
	Inserthttprequest   string  `json:"inserthttprequest,omitempty"`
	Inserticapheaders   string  `json:"inserticapheaders,omitempty"`
	Inspecthttp2        string  `json:"inspecthttp2,omitempty"`
	Logaction           string  `json:"logaction,omitempty"`
	Mode                string  `json:"mode,omitempty"`
	Name                string  `json:"name,omitempty"`
	Nextgenapiresource  string  `json:"_nextgenapiresource,omitempty"`
	Preview             string  `json:"preview,omitempty"`
	Previewlength       int     `json:"previewlength,omitempty"`
	Queryparams         string  `json:"queryparams,omitempty"`
	Reqtimeout          int     `json:"reqtimeout,omitempty"`
	Reqtimeoutaction    string  `json:"reqtimeoutaction,omitempty"`
	Uri                 string  `json:"uri,omitempty"`
	Useragent           string  `json:"useragent,omitempty"`
}

type Nssimpleacl struct {
	Aclaction          string  `json:"aclaction,omitempty"`
	Aclname            string  `json:"aclname,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Destport           int     `json:"destport,omitempty"`
	Estsessions        bool    `json:"estsessions,omitempty"`
	Hits               int     `json:"hits,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Protocol           string  `json:"protocol,omitempty"`
	Srcip              string  `json:"srcip,omitempty"`
	Td                 int     `json:"td,omitempty"`
	Ttl                int     `json:"ttl,omitempty"`
}

type Nsvariablevalues struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Nodeid             int     `json:"nodeid,omitempty"`
	Variabledata       string  `json:"variabledata,omitempty"`
	Variablekey        string  `json:"variablekey,omitempty"`
	Variablevalue      string  `json:"variablevalue,omitempty"`
}

type Nspbrs struct {
}

type Nsweblogparam struct {
	Buffersizemb       int      `json:"buffersizemb,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	Customreqhdrs      []string `json:"customreqhdrs,omitempty"`
	Customrsphdrs      []string `json:"customrsphdrs,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
}

type Nslicenseactivationdata struct {
	Filename           string `json:"filename,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nsevents struct {
	Count              float64 `json:"__count,omitempty"`
	Data0              int     `json:"data0,omitempty"`
	Data1              int     `json:"data1,omitempty"`
	Data2              int     `json:"data2,omitempty"`
	Data3              int     `json:"data3,omitempty"`
	Devid              int     `json:"devid,omitempty"`
	Devname            string  `json:"devname,omitempty"`
	Eventcode          int     `json:"eventcode,omitempty"`
	Eventno            int     `json:"eventno,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Text               string  `json:"text,omitempty"`
	Time               int     `json:"time,omitempty"`
}

type Nsmigration struct {
	Count                           float64 `json:"__count,omitempty"`
	Destip                          string  `json:"destip,omitempty"`
	Destport                        int     `json:"destport,omitempty"`
	Dumpsession                     string  `json:"dumpsession,omitempty"`
	Migdfdsessionsactive            int     `json:"migdfdsessionsactive,omitempty"`
	Migdfdsessionsactiverollback    int     `json:"migdfdsessionsactiverollback,omitempty"`
	Migdfdsessionsallocated         int     `json:"migdfdsessionsallocated,omitempty"`
	Migdfdsessionsallocatedrollback int     `json:"migdfdsessionsallocatedrollback,omitempty"`
	Mighastateflag                  int     `json:"mighastateflag,omitempty"`
	Migl4sessionsactive             int     `json:"migl4sessionsactive,omitempty"`
	Migl4sessionsactiverollback     int     `json:"migl4sessionsactiverollback,omitempty"`
	Migl4sessionsallocated          int     `json:"migl4sessionsallocated,omitempty"`
	Migl4sessionsallocatedrollback  int     `json:"migl4sessionsallocatedrollback,omitempty"`
	Migrationendtime                string  `json:"migrationendtime,omitempty"`
	Migrationrollbackstarttime      string  `json:"migrationrollbackstarttime,omitempty"`
	Migrationstarttime              string  `json:"migrationstarttime,omitempty"`
	Migrationstatus                 string  `json:"migrationstatus,omitempty"`
	Nextgenapiresource              string  `json:"_nextgenapiresource,omitempty"`
	Srcip                           string  `json:"srcip,omitempty"`
	Srcport                         int     `json:"srcport,omitempty"`
	Timeout                         int     `json:"timeout,omitempty"`
}

type Nshmackey struct {
	Comment            string  `json:"comment,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Digest             string  `json:"digest,omitempty"`
	Keyvalue           string  `json:"keyvalue,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
}

type Nspbr6 struct {
	Action             string  `json:"action,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Curstate           int     `json:"curstate,omitempty"`
	Data               bool    `json:"data,omitempty"`
	Destipop           string  `json:"destipop,omitempty"`
	Destipv6           bool    `json:"destipv6,omitempty"`
	Destipv6val        string  `json:"destipv6val,omitempty"`
	Destport           bool    `json:"destport,omitempty"`
	Destportop         string  `json:"destportop,omitempty"`
	Destportval        string  `json:"destportval,omitempty"`
	Detail             bool    `json:"detail,omitempty"`
	Failedprobes       int     `json:"failedprobes,omitempty"`
	Hits               int     `json:"hits,omitempty"`
	InterfaceField     string  `json:"Interface,omitempty"`
	Iptunnel           string  `json:"iptunnel,omitempty"`
	Kernelstate        string  `json:"kernelstate,omitempty"`
	Monitor            string  `json:"monitor,omitempty"`
	Monstatcode        int     `json:"monstatcode,omitempty"`
	Monstatparam1      int     `json:"monstatparam1,omitempty"`
	Monstatparam2      int     `json:"monstatparam2,omitempty"`
	Monstatparam3      int     `json:"monstatparam3,omitempty"`
	Msr                string  `json:"msr,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Nexthop            bool    `json:"nexthop,omitempty"`
	Nexthopval         string  `json:"nexthopval,omitempty"`
	Nexthopvlan        int     `json:"nexthopvlan,omitempty"`
	Ownergroup         string  `json:"ownergroup,omitempty"`
	Priority           int     `json:"priority,omitempty"`
	Protocol           string  `json:"protocol,omitempty"`
	Protocolnumber     int     `json:"protocolnumber,omitempty"`
	Srcipop            string  `json:"srcipop,omitempty"`
	Srcipv6            bool    `json:"srcipv6,omitempty"`
	Srcipv6val         string  `json:"srcipv6val,omitempty"`
	Srcmac             string  `json:"srcmac,omitempty"`
	Srcmacmask         string  `json:"srcmacmask,omitempty"`
	Srcport            bool    `json:"srcport,omitempty"`
	Srcportop          string  `json:"srcportop,omitempty"`
	Srcportval         string  `json:"srcportval,omitempty"`
	State              string  `json:"state,omitempty"`
	Td                 int     `json:"td,omitempty"`
	Totalfailedprobes  int     `json:"totalfailedprobes,omitempty"`
	Totalprobes        int     `json:"totalprobes,omitempty"`
	Vlan               int     `json:"vlan,omitempty"`
	Vxlan              int     `json:"vxlan,omitempty"`
	Vxlanvlanmap       string  `json:"vxlanvlanmap,omitempty"`
}

type Nsacls struct {
	TypeField string `json:"type,omitempty"`
}

type Nsdhcpparams struct {
	Dhcpclient         string `json:"dhcpclient,omitempty"`
	Hostrtgw           string `json:"hostrtgw,omitempty"`
	Ipaddress          string `json:"ipaddress,omitempty"`
	Netmask            string `json:"netmask,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Running            bool   `json:"running,omitempty"`
	Saveroute          string `json:"saveroute,omitempty"`
}

type Nssurgeq struct {
	Name       string `json:"name,omitempty"`
	Port       int    `json:"port,omitempty"`
	Servername string `json:"servername,omitempty"`
}

type Nsnextgenapi struct {
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	State              string `json:"state,omitempty"`
}

type Nsxmlnamespace struct {
	Count              float64 `json:"__count,omitempty"`
	Description        string  `json:"description,omitempty"`
	Namespace          string  `json:"Namespace,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Prefix             string  `json:"prefix,omitempty"`
}

type Nsacl struct {
	Aclaction          string   `json:"aclaction,omitempty"`
	Aclassociate       []string `json:"aclassociate,omitempty"`
	Aclchildcount      int      `json:"aclchildcount,omitempty"`
	Aclname            string   `json:"aclname,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Destip             bool     `json:"destip,omitempty"`
	Destipdataset      string   `json:"destipdataset,omitempty"`
	Destipop           string   `json:"destipop,omitempty"`
	Destipval          string   `json:"destipval,omitempty"`
	Destport           bool     `json:"destport,omitempty"`
	Destportdataset    string   `json:"destportdataset,omitempty"`
	Destportop         string   `json:"destportop,omitempty"`
	Destportval        string   `json:"destportval,omitempty"`
	Dfdhash            string   `json:"dfdhash,omitempty"`
	Established        bool     `json:"established,omitempty"`
	Hits               int      `json:"hits,omitempty"`
	Icmpcode           int      `json:"icmpcode,omitempty"`
	Icmptype           int      `json:"icmptype,omitempty"`
	InterfaceField     string   `json:"Interface,omitempty"`
	Kernelstate        string   `json:"kernelstate,omitempty"`
	Logstate           string   `json:"logstate,omitempty"`
	Newname            string   `json:"newname,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Nodeid             int      `json:"nodeid,omitempty"`
	Priority           int      `json:"priority,omitempty"`
	Protocol           string   `json:"protocol,omitempty"`
	Protocolnumber     int      `json:"protocolnumber,omitempty"`
	Ratelimit          int      `json:"ratelimit,omitempty"`
	Srcip              bool     `json:"srcip,omitempty"`
	Srcipdataset       string   `json:"srcipdataset,omitempty"`
	Srcipop            string   `json:"srcipop,omitempty"`
	Srcipval           string   `json:"srcipval,omitempty"`
	Srcmac             string   `json:"srcmac,omitempty"`
	Srcmacmask         string   `json:"srcmacmask,omitempty"`
	Srcport            bool     `json:"srcport,omitempty"`
	Srcportdataset     string   `json:"srcportdataset,omitempty"`
	Srcportop          string   `json:"srcportop,omitempty"`
	Srcportval         string   `json:"srcportval,omitempty"`
	State              string   `json:"state,omitempty"`
	Stateful           string   `json:"stateful,omitempty"`
	Td                 int      `json:"td,omitempty"`
	Ttl                int      `json:"ttl,omitempty"`
	TypeField          string   `json:"type,omitempty"`
	Vlan               int      `json:"vlan,omitempty"`
	Vxlan              int      `json:"vxlan,omitempty"`
}

type Nslicenseserver struct {
	Count              float64 `json:"__count,omitempty"`
	Deviceprofilename  string  `json:"deviceprofilename,omitempty"`
	Forceupdateip      bool    `json:"forceupdateip,omitempty"`
	Gptimeleft         int     `json:"gptimeleft,omitempty"`
	Grace              int     `json:"grace,omitempty"`
	Licensemode        string  `json:"licensemode,omitempty"`
	Licenseserverip    string  `json:"licenseserverip,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Nodeid             int     `json:"nodeid,omitempty"`
	Password           string  `json:"password,omitempty"`
	Port               int     `json:"port,omitempty"`
	Servername         string  `json:"servername,omitempty"`
	Status             int     `json:"status,omitempty"`
	TypeField          string  `json:"type,omitempty"`
	Username           string  `json:"username,omitempty"`
}

type Nsmgmtparam struct {
	Httpdmaxclients    int    `json:"httpdmaxclients,omitempty"`
	Httpdmaxreqworkers int    `json:"httpdmaxreqworkers,omitempty"`
	Mgmthttpport       int    `json:"mgmthttpport,omitempty"`
	Mgmthttpsport      int    `json:"mgmthttpsport,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nsdhcpip struct {
}

type Nstcpprofile struct {
	Ackaggregation              string   `json:"ackaggregation,omitempty"`
	Ackonpush                   string   `json:"ackonpush,omitempty"`
	Applyadaptivetcp            string   `json:"applyadaptivetcp,omitempty"`
	Buffersize                  int      `json:"buffersize,omitempty"`
	Builtin                     []string `json:"builtin,omitempty"`
	Burstratecontrol            string   `json:"burstratecontrol,omitempty"`
	Clientiptcpoption           string   `json:"clientiptcpoption,omitempty"`
	Clientiptcpoptionnumber     int      `json:"clientiptcpoptionnumber,omitempty"`
	Count                       float64  `json:"__count,omitempty"`
	Delayedack                  int      `json:"delayedack,omitempty"`
	Dropestconnontimeout        string   `json:"dropestconnontimeout,omitempty"`
	Drophalfclosedconnontimeout string   `json:"drophalfclosedconnontimeout,omitempty"`
	Dsack                       string   `json:"dsack,omitempty"`
	Dupackthresh                int      `json:"dupackthresh,omitempty"`
	Dynamicreceivebuffering     string   `json:"dynamicreceivebuffering,omitempty"`
	Ecn                         string   `json:"ecn,omitempty"`
	Establishclientconn         string   `json:"establishclientconn,omitempty"`
	Fack                        string   `json:"fack,omitempty"`
	Feature                     string   `json:"feature,omitempty"`
	Flavor                      string   `json:"flavor,omitempty"`
	Frto                        string   `json:"frto,omitempty"`
	Hystart                     string   `json:"hystart,omitempty"`
	Initialcwnd                 int      `json:"initialcwnd,omitempty"`
	Ka                          string   `json:"ka,omitempty"`
	Kaconnidletime              int      `json:"kaconnidletime,omitempty"`
	Kamaxprobes                 int      `json:"kamaxprobes,omitempty"`
	Kaprobeinterval             int      `json:"kaprobeinterval,omitempty"`
	Kaprobeupdatelastactivity   string   `json:"kaprobeupdatelastactivity,omitempty"`
	Maxburst                    int      `json:"maxburst,omitempty"`
	Maxcwnd                     int      `json:"maxcwnd,omitempty"`
	Maxpktpermss                int      `json:"maxpktpermss,omitempty"`
	Minrto                      int      `json:"minrto,omitempty"`
	Mpcapablecbit               string   `json:"mpcapablecbit,omitempty"`
	Mptcp                       string   `json:"mptcp,omitempty"`
	Mptcpdropdataonpreestsf     string   `json:"mptcpdropdataonpreestsf,omitempty"`
	Mptcpfastopen               string   `json:"mptcpfastopen,omitempty"`
	Mptcpsessiontimeout         int      `json:"mptcpsessiontimeout,omitempty"`
	Mss                         int      `json:"mss,omitempty"`
	Nagle                       string   `json:"nagle,omitempty"`
	Name                        string   `json:"name,omitempty"`
	Nextgenapiresource          string   `json:"_nextgenapiresource,omitempty"`
	Oooqsize                    int      `json:"oooqsize,omitempty"`
	Pktperretx                  int      `json:"pktperretx,omitempty"`
	Rateqmax                    int      `json:"rateqmax,omitempty"`
	Refcnt                      int      `json:"refcnt,omitempty"`
	Rfc5961compliance           string   `json:"rfc5961compliance,omitempty"`
	Rstmaxack                   string   `json:"rstmaxack,omitempty"`
	Rstwindowattenuate          string   `json:"rstwindowattenuate,omitempty"`
	Sack                        string   `json:"sack,omitempty"`
	Sendbuffsize                int      `json:"sendbuffsize,omitempty"`
	Sendclientportintcpoption   string   `json:"sendclientportintcpoption,omitempty"`
	Slowstartincr               int      `json:"slowstartincr,omitempty"`
	Slowstartthreshold          int      `json:"slowstartthreshold,omitempty"`
	Spoofsyndrop                string   `json:"spoofsyndrop,omitempty"`
	Syncookie                   string   `json:"syncookie,omitempty"`
	Taillossprobe               string   `json:"taillossprobe,omitempty"`
	Tcpfastopen                 string   `json:"tcpfastopen,omitempty"`
	Tcpfastopencookiesize       int      `json:"tcpfastopencookiesize,omitempty"`
	Tcpmode                     string   `json:"tcpmode,omitempty"`
	Tcprate                     int      `json:"tcprate,omitempty"`
	Tcpsegoffload               string   `json:"tcpsegoffload,omitempty"`
	Timestamp                   string   `json:"timestamp,omitempty"`
	Ws                          string   `json:"ws,omitempty"`
	Wsval                       int      `json:"wsval,omitempty"`
}

type NstrafficdomainBinding struct {
	NstrafficdomainBridgegroupBinding []interface{} `json:"nstrafficdomain_bridgegroup_binding,omitempty"`
	NstrafficdomainVlanBinding        []interface{} `json:"nstrafficdomain_vlan_binding,omitempty"`
	NstrafficdomainVxlanBinding       []interface{} `json:"nstrafficdomain_vxlan_binding,omitempty"`
	Td                                int           `json:"td,omitempty"`
}

type NstimerAutoscalepolicyBinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Name                   string `json:"name,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Samplesize             int    `json:"samplesize,omitempty"`
	Threshold              int    `json:"threshold,omitempty"`
	Vserver                string `json:"vserver,omitempty"`
}

type Nsmode struct {
	Bridgebpdus         bool     `json:"bridgebpdus,omitempty"`
	Cka                 bool     `json:"cka,omitempty"`
	Dradv               bool     `json:"dradv,omitempty"`
	Dradv6              bool     `json:"dradv6,omitempty"`
	Edge                bool     `json:"edge,omitempty"`
	Fr                  bool     `json:"fr,omitempty"`
	Iradv               bool     `json:"iradv,omitempty"`
	L2                  bool     `json:"l2,omitempty"`
	L3                  bool     `json:"l3,omitempty"`
	Mbf                 bool     `json:"mbf,omitempty"`
	Mediaclassification bool     `json:"mediaclassification,omitempty"`
	Mode                []string `json:"mode,omitempty"`
	Nextgenapiresource  string   `json:"_nextgenapiresource,omitempty"`
	Pmtud               bool     `json:"pmtud,omitempty"`
	SingleIp            bool     `json:"single_ip,omitempty"`
	Sradv               bool     `json:"sradv,omitempty"`
	Sradv6              bool     `json:"sradv6,omitempty"`
	Tcpb                bool     `json:"tcpb,omitempty"`
	Ulfd                bool     `json:"ulfd,omitempty"`
	Usip                bool     `json:"usip,omitempty"`
	Usnip               bool     `json:"usnip,omitempty"`
}

type Nssimpleacl6 struct {
	Aclaction          string  `json:"aclaction,omitempty"`
	Aclname            string  `json:"aclname,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Destport           int     `json:"destport,omitempty"`
	Estsessions        bool    `json:"estsessions,omitempty"`
	Hits               int     `json:"hits,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Protocol           string  `json:"protocol,omitempty"`
	Srcipv6            string  `json:"srcipv6,omitempty"`
	Td                 int     `json:"td,omitempty"`
	Ttl                int     `json:"ttl,omitempty"`
}

type Nslaslicense struct {
	Daystoexpiration   int    `json:"daystoexpiration,omitempty"`
	Filelocation       string `json:"filelocation,omitempty"`
	Filename           string `json:"filename,omitempty"`
	Fixedbandwidth     bool   `json:"fixedbandwidth,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Renewalnext        int    `json:"renewalnext,omitempty"`
	Renewalnextdate    string `json:"renewalnextdate,omitempty"`
	Renewalprev        int    `json:"renewalprev,omitempty"`
	Renewalprevdate    string `json:"renewalprevdate,omitempty"`
	Status             string `json:"status,omitempty"`
}

type Nsconfig struct {
	All                     bool          `json:"all,omitempty"`
	Async                   bool          `json:"Async,omitempty"`
	Changedpassword         bool          `json:"changedpassword,omitempty"`
	Cip                     string        `json:"cip,omitempty"`
	Cipheader               string        `json:"cipheader,omitempty"`
	Config                  string        `json:"config,omitempty"`
	Config1                 string        `json:"config1,omitempty"`
	Config2                 string        `json:"config2,omitempty"`
	Configchanged           bool          `json:"configchanged,omitempty"`
	Configfile              string        `json:"configfile,omitempty"`
	Cookieversion           string        `json:"cookieversion,omitempty"`
	Crportrange             string        `json:"crportrange,omitempty"`
	Currentsytemtime        string        `json:"currentsytemtime,omitempty"`
	Exclusivequotamaxclient int           `json:"exclusivequotamaxclient,omitempty"`
	Exclusivequotaspillover int           `json:"exclusivequotaspillover,omitempty"`
	Flags                   int           `json:"flags,omitempty"`
	Force                   bool          `json:"force,omitempty"`
	Ftpportrange            string        `json:"ftpportrange,omitempty"`
	Grantquotamaxclient     int           `json:"grantquotamaxclient,omitempty"`
	Grantquotaspillover     int           `json:"grantquotaspillover,omitempty"`
	Httpport                []interface{} `json:"httpport,omitempty"`
	Id                      int           `json:"id,omitempty"`
	Ifnum                   []string      `json:"ifnum,omitempty"`
	Ignoredevicespecific    bool          `json:"ignoredevicespecific,omitempty"`
	Ipaddress               string        `json:"ipaddress,omitempty"`
	Lastconfigchangedtime   string        `json:"lastconfigchangedtime,omitempty"`
	Lastconfigsavetime      string        `json:"lastconfigsavetime,omitempty"`
	Level                   string        `json:"level,omitempty"`
	Mappedip                string        `json:"mappedip,omitempty"`
	Maxconn                 int           `json:"maxconn,omitempty"`
	Maxreq                  int           `json:"maxreq,omitempty"`
	Message                 string        `json:"message,omitempty"`
	Netmask                 string        `json:"netmask,omitempty"`
	Nextgenapiresource      string        `json:"_nextgenapiresource,omitempty"`
	Nsvlan                  int           `json:"nsvlan,omitempty"`
	Outtype                 string        `json:"outtype,omitempty"`
	Pmtumin                 int           `json:"pmtumin,omitempty"`
	Pmtutimeout             int           `json:"pmtutimeout,omitempty"`
	Primaryip               string        `json:"primaryip,omitempty"`
	Primaryip6              string        `json:"primaryip6,omitempty"`
	Range                   int           `json:"range,omitempty"`
	Rbaconfig               string        `json:"rbaconfig,omitempty"`
	Response                string        `json:"response,omitempty"`
	Responsefile            string        `json:"responsefile,omitempty"`
	Securecookie            string        `json:"securecookie,omitempty"`
	Securemanagementtd      int           `json:"securemanagementtd,omitempty"`
	Securemanagementtraffic string        `json:"securemanagementtraffic,omitempty"`
	Svmcmd                  int           `json:"svmcmd,omitempty"`
	Systemtime              int           `json:"systemtime,omitempty"`
	Systemtype              string        `json:"systemtype,omitempty"`
	Tagged                  string        `json:"tagged,omitempty"`
	Template                bool          `json:"template,omitempty"`
	Timezone                string        `json:"timezone,omitempty"`
	Weakpassword            bool          `json:"weakpassword,omitempty"`
}

type Nsaptlicense struct {
	Bindtype           string   `json:"bindtype,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Countavailable     string   `json:"countavailable,omitempty"`
	Counttotal         string   `json:"counttotal,omitempty"`
	Dateexp            string   `json:"dateexp,omitempty"`
	Datepurchased      string   `json:"datepurchased,omitempty"`
	Datesa             string   `json:"datesa,omitempty"`
	Features           []string `json:"features,omitempty"`
	Id                 string   `json:"id,omitempty"`
	Licensedir         string   `json:"licensedir,omitempty"`
	Name               string   `json:"name,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Relevance          string   `json:"relevance,omitempty"`
	Response           string   `json:"response,omitempty"`
	Serialno           string   `json:"serialno,omitempty"`
	Sessionid          string   `json:"sessionid,omitempty"`
	Useproxy           string   `json:"useproxy,omitempty"`
}

type Nshttpprofile struct {
	Adpttimeout                      string   `json:"adpttimeout,omitempty"`
	Allowonlywordcharactersandhyphen string   `json:"allowonlywordcharactersandhyphen,omitempty"`
	Altsvc                           string   `json:"altsvc,omitempty"`
	Altsvcvalue                      string   `json:"altsvcvalue,omitempty"`
	Apdexcltresptimethreshold        int      `json:"apdexcltresptimethreshold,omitempty"`
	Apdexsvrresptimethreshold        int      `json:"apdexsvrresptimethreshold,omitempty"`
	Builtin                          []string `json:"builtin,omitempty"`
	Clientiphdrexpr                  string   `json:"clientiphdrexpr,omitempty"`
	Cmponpush                        string   `json:"cmponpush,omitempty"`
	Conmultiplex                     string   `json:"conmultiplex,omitempty"`
	Count                            float64  `json:"__count,omitempty"`
	Dropextracrlf                    string   `json:"dropextracrlf,omitempty"`
	Dropextradata                    string   `json:"dropextradata,omitempty"`
	Dropinvalreqs                    string   `json:"dropinvalreqs,omitempty"`
	Dropinvalreqswarning             string   `json:"dropinvalreqswarning,omitempty"`
	Feature                          string   `json:"feature,omitempty"`
	Grpcholdlimit                    int      `json:"grpcholdlimit,omitempty"`
	Grpcholdtimeout                  int      `json:"grpcholdtimeout,omitempty"`
	Grpclengthdelimitation           string   `json:"grpclengthdelimitation,omitempty"`
	Hostheadervalidation             string   `json:"hostheadervalidation,omitempty"`
	Http2                            string   `json:"http2,omitempty"`
	Http2altsvcframe                 string   `json:"http2altsvcframe,omitempty"`
	Http2direct                      string   `json:"http2direct,omitempty"`
	Http2extendedconnect             string   `json:"http2extendedconnect,omitempty"`
	Http2headertablesize             int      `json:"http2headertablesize,omitempty"`
	Http2initialconnwindowsize       int      `json:"http2initialconnwindowsize,omitempty"`
	Http2initialwindowsize           int      `json:"http2initialwindowsize,omitempty"`
	Http2maxconcurrentstreams        int      `json:"http2maxconcurrentstreams,omitempty"`
	Http2maxemptyframespermin        int      `json:"http2maxemptyframespermin,omitempty"`
	Http2maxframesize                int      `json:"http2maxframesize,omitempty"`
	Http2maxheaderlistsize           int      `json:"http2maxheaderlistsize,omitempty"`
	Http2maxpingframespermin         int      `json:"http2maxpingframespermin,omitempty"`
	Http2maxresetframespermin        int      `json:"http2maxresetframespermin,omitempty"`
	Http2maxrxresetframespermin      int      `json:"http2maxrxresetframespermin,omitempty"`
	Http2maxsettingsframespermin     int      `json:"http2maxsettingsframespermin,omitempty"`
	Http2minseverconn                int      `json:"http2minseverconn,omitempty"`
	Http2strictcipher                string   `json:"http2strictcipher,omitempty"`
	Http3                            string   `json:"http3,omitempty"`
	Http3maxheaderblockedstreams     int      `json:"http3maxheaderblockedstreams,omitempty"`
	Http3maxheaderfieldsectionsize   int      `json:"http3maxheaderfieldsectionsize,omitempty"`
	Http3maxheadertablesize          int      `json:"http3maxheadertablesize,omitempty"`
	Http3minseverconn                int      `json:"http3minseverconn,omitempty"`
	Http3webtransport                string   `json:"http3webtransport,omitempty"`
	Httppipelinebuffsize             int      `json:"httppipelinebuffsize,omitempty"`
	Incomphdrdelay                   int      `json:"incomphdrdelay,omitempty"`
	Markconnreqinval                 string   `json:"markconnreqinval,omitempty"`
	Markhttp09inval                  string   `json:"markhttp09inval,omitempty"`
	Markhttpheaderextrawserror       string   `json:"markhttpheaderextrawserror,omitempty"`
	Markrfc7230noncompliantinval     string   `json:"markrfc7230noncompliantinval,omitempty"`
	Marktracereqinval                string   `json:"marktracereqinval,omitempty"`
	Maxduplicateheaderfields         int      `json:"maxduplicateheaderfields,omitempty"`
	Maxheaderfieldlen                int      `json:"maxheaderfieldlen,omitempty"`
	Maxheaderlen                     int      `json:"maxheaderlen,omitempty"`
	Maxreq                           int      `json:"maxreq,omitempty"`
	Maxreusepool                     int      `json:"maxreusepool,omitempty"`
	Minreusepool                     int      `json:"minreusepool,omitempty"`
	Name                             string   `json:"name,omitempty"`
	Nextgenapiresource               string   `json:"_nextgenapiresource,omitempty"`
	Passprotocolupgrade              string   `json:"passprotocolupgrade,omitempty"`
	Persistentetag                   string   `json:"persistentetag,omitempty"`
	Refcnt                           int      `json:"refcnt,omitempty"`
	Reqtimeout                       int      `json:"reqtimeout,omitempty"`
	Reqtimeoutaction                 string   `json:"reqtimeoutaction,omitempty"`
	Reusepooltimeout                 int      `json:"reusepooltimeout,omitempty"`
	Rtsptunnel                       string   `json:"rtsptunnel,omitempty"`
	Weblog                           string   `json:"weblog,omitempty"`
	Websocket                        string   `json:"websocket,omitempty"`
}

type Nsfeature struct {
	Aaa                bool     `json:"aaa,omitempty"`
	Adaptivetcp        bool     `json:"adaptivetcp,omitempty"`
	Apigateway         bool     `json:"apigateway,omitempty"`
	Appflow            bool     `json:"appflow,omitempty"`
	Appfw              bool     `json:"appfw,omitempty"`
	Appqoe             bool     `json:"appqoe,omitempty"`
	Bgp                bool     `json:"bgp,omitempty"`
	Bot                bool     `json:"bot,omitempty"`
	Cf                 bool     `json:"cf,omitempty"`
	Ch                 bool     `json:"ch,omitempty"`
	Ci                 bool     `json:"ci,omitempty"`
	Cloudbridge        bool     `json:"cloudbridge,omitempty"`
	Cmp                bool     `json:"cmp,omitempty"`
	Contentaccelerator bool     `json:"contentaccelerator,omitempty"`
	Cqa                bool     `json:"cqa,omitempty"`
	Cr                 bool     `json:"cr,omitempty"`
	Cs                 bool     `json:"cs,omitempty"`
	Feature            []string `json:"feature,omitempty"`
	Feo                bool     `json:"feo,omitempty"`
	Forwardproxy       bool     `json:"forwardproxy,omitempty"`
	Gslb               bool     `json:"gslb,omitempty"`
	Ic                 bool     `json:"ic,omitempty"`
	Ipv6pt             bool     `json:"ipv6pt,omitempty"`
	Isis               bool     `json:"isis,omitempty"`
	Lb                 bool     `json:"lb,omitempty"`
	Lsn                bool     `json:"lsn,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Ospf               bool     `json:"ospf,omitempty"`
	Push               bool     `json:"push,omitempty"`
	Rdpproxy           bool     `json:"rdpproxy,omitempty"`
	Rep                bool     `json:"rep,omitempty"`
	Responder          bool     `json:"responder,omitempty"`
	Rewrite            bool     `json:"rewrite,omitempty"`
	Rip                bool     `json:"rip,omitempty"`
	Sp                 bool     `json:"sp,omitempty"`
	Ssl                bool     `json:"ssl,omitempty"`
	Sslinterception    bool     `json:"sslinterception,omitempty"`
	Sslvpn             bool     `json:"sslvpn,omitempty"`
	Videooptimization  bool     `json:"videooptimization,omitempty"`
	Wl                 bool     `json:"wl,omitempty"`
}

type NsservicepathBinding struct {
	NsservicepathNsservicefunctionBinding []interface{} `json:"nsservicepath_nsservicefunction_binding,omitempty"`
	Servicepathname                       string        `json:"servicepathname,omitempty"`
}

type Nslicenseparameters struct {
	Alert1gracetimeout       int    `json:"alert1gracetimeout,omitempty"`
	Alert2gracetimeout       int    `json:"alert2gracetimeout,omitempty"`
	Heartbeatinterval        int    `json:"heartbeatinterval,omitempty"`
	Inventoryrefreshinterval int    `json:"inventoryrefreshinterval,omitempty"`
	Licenseexpiryalerttime   int    `json:"licenseexpiryalerttime,omitempty"`
	Nextgenapiresource       string `json:"_nextgenapiresource,omitempty"`
}

type Nslimitidentifier struct {
	Computedtraptimeslice    int      `json:"computedtraptimeslice,omitempty"`
	Count                    float64  `json:"__count,omitempty"`
	Drop                     int      `json:"drop,omitempty"`
	Hits                     int      `json:"hits,omitempty"`
	Limitidentifier          string   `json:"limitidentifier,omitempty"`
	Limittype                string   `json:"limittype,omitempty"`
	Maxbandwidth             int      `json:"maxbandwidth,omitempty"`
	Mode                     string   `json:"mode,omitempty"`
	Nextgenapiresource       string   `json:"_nextgenapiresource,omitempty"`
	Ngname                   string   `json:"ngname,omitempty"`
	Referencecount           int      `json:"referencecount,omitempty"`
	Rule                     []string `json:"rule,omitempty"`
	Selectorname             string   `json:"selectorname,omitempty"`
	Threshold                int      `json:"threshold,omitempty"`
	Time                     int      `json:"time,omitempty"`
	Timeslice                int      `json:"timeslice,omitempty"`
	Total                    int      `json:"total,omitempty"`
	Trapscomputedintimeslice int      `json:"trapscomputedintimeslice,omitempty"`
	Trapsintimeslice         int      `json:"trapsintimeslice,omitempty"`
}

type Nshardware struct {
	Bmcrevision        string `json:"bmcrevision,omitempty"`
	Cpufrequncy        int    `json:"cpufrequncy,omitempty"`
	Encodedserialno    string `json:"encodedserialno,omitempty"`
	Host               string `json:"host,omitempty"`
	Hostid             int    `json:"hostid,omitempty"`
	Hwdescription      string `json:"hwdescription,omitempty"`
	Manufactureday     int    `json:"manufactureday,omitempty"`
	Manufacturemonth   int    `json:"manufacturemonth,omitempty"`
	Manufactureyear    int    `json:"manufactureyear,omitempty"`
	Netscaleruuid      string `json:"netscaleruuid,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Serialno           string `json:"serialno,omitempty"`
	Sysid              int    `json:"sysid,omitempty"`
}

type Nstrafficdomain struct {
	Aliasname          string  `json:"aliasname,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	State              string  `json:"state,omitempty"`
	Td                 int     `json:"td,omitempty"`
	Vmac               string  `json:"vmac,omitempty"`
}

type NspartitionVlanBinding struct {
	Partitionname string `json:"partitionname,omitempty"`
	Vlan          int    `json:"vlan,omitempty"`
}

type Nscqaparam struct {
	Harqretxdelay      int     `json:"harqretxdelay,omitempty"`
	Lr1coeflist        string  `json:"lr1coeflist,omitempty"`
	Lr1probthresh      float64 `json:"lr1probthresh,omitempty"`
	Lr2coeflist        string  `json:"lr2coeflist,omitempty"`
	Lr2probthresh      float64 `json:"lr2probthresh,omitempty"`
	Minrttnet1         int     `json:"minrttnet1,omitempty"`
	Minrttnet2         int     `json:"minrttnet2,omitempty"`
	Minrttnet3         int     `json:"minrttnet3,omitempty"`
	Net1cclscale       string  `json:"net1cclscale,omitempty"`
	Net1csqscale       string  `json:"net1csqscale,omitempty"`
	Net1label          string  `json:"net1label,omitempty"`
	Net1logcoef        string  `json:"net1logcoef,omitempty"`
	Net2cclscale       string  `json:"net2cclscale,omitempty"`
	Net2csqscale       string  `json:"net2csqscale,omitempty"`
	Net2label          string  `json:"net2label,omitempty"`
	Net2logcoef        string  `json:"net2logcoef,omitempty"`
	Net3cclscale       string  `json:"net3cclscale,omitempty"`
	Net3csqscale       string  `json:"net3csqscale,omitempty"`
	Net3label          string  `json:"net3label,omitempty"`
	Net3logcoef        string  `json:"net3logcoef,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
}

type Nsstats struct {
	Cleanuplevel string `json:"cleanuplevel,omitempty"`
}

type NslimitidentifierBinding struct {
	Limitidentifier                         string        `json:"limitidentifier,omitempty"`
	NslimitidentifierNslimitsessionsBinding []interface{} `json:"nslimitidentifier_nslimitsessions_binding,omitempty"`
}
