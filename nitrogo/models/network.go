// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Inatparam struct {
	Nat46v6prefix      string `json:"nat46v6prefix,omitempty"`
	Td                 int    `json:"td"`
	Nat46ignoretos     string `json:"nat46ignoretos,omitempty"`
	Nat46zerochecksum  string `json:"nat46zerochecksum,omitempty"`
	Nat46v6mtu         int    `json:"nat46v6mtu,omitempty"`
	Nat46fragheader    string `json:"nat46fragheader,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Ipsetnsipbinding struct {
	Ipaddress string `json:"ipaddress,omitempty"`
	Name      string `json:"name,omitempty"`
}

type Mapbmr struct {
	Name               string `json:"name,omitempty"`
	Ruleipv6prefix     string `json:"ruleipv6prefix,omitempty"`
	Psidoffset         int    `json:"psidoffset,omitempty"`
	Eabitlength        int    `json:"eabitlength,omitempty"`
	Psidlength         int    `json:"psidlength,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Netprofile struct {
	Name                           string `json:"name,omitempty"`
	Td                             int    `json:"td,omitempty"`
	Srcip                          string `json:"srcip,omitempty"`
	Srcippersistency               string `json:"srcippersistency,omitempty"`
	Overridelsn                    string `json:"overridelsn,omitempty"`
	Mbf                            string `json:"mbf,omitempty"`
	Proxyprotocol                  string `json:"proxyprotocol,omitempty"`
	Proxyprotocoltxversion         string `json:"proxyprotocoltxversion,omitempty"`
	Proxyprotocolaftertlshandshake string `json:"proxyprotocolaftertlshandshake,omitempty"`
	Proxyprotocoltlvoptions        string `json:"proxyprotocoltlvoptions,omitempty"`
	Nextgenapiresource             string `json:"_nextgenapiresource,omitempty"`
}

type Netprofilebinding struct {
	Name string `json:"name,omitempty"`
}

type Rnat6 struct {
	Name               string `json:"name,omitempty"`
	Network            string `json:"network,omitempty"`
	Acl6name           string `json:"acl6name,omitempty"`
	Redirectport       int    `json:"redirectport,omitempty"`
	Td                 int    `json:"td,omitempty"`
	Srcippersistency   string `json:"srcippersistency,omitempty"`
	Ownergroup         string `json:"ownergroup,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Rnatnsipbinding struct {
	Natip      string `json:"natip,omitempty"`
	Td         int    `json:"td,omitempty"`
	Ownergroup string `json:"ownergroup,omitempty"`
	Name       string `json:"name,omitempty"`
}

type Rnatglobalsyslogpolicybinding struct {
	Policy   string `json:"policy,omitempty"`
	Priority uint32 `json:"priority,omitempty"`
	All      bool   `json:"all,omitempty"`
}

type Ci struct {
	Ifaces             string `json:"ifaces,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Route6 struct {
	Network            string `json:"network,omitempty"`
	Gateway            string `json:"gateway,omitempty"`
	Vlan               int    `json:"vlan,omitempty"`
	Vxlan              int    `json:"vxlan,omitempty"`
	Weight             int    `json:"weight,omitempty"`
	Distance           int    `json:"distance,omitempty"`
	Cost               int    `json:"cost,omitempty"`
	Advertise          string `json:"advertise,omitempty"`
	Msr                string `json:"msr,omitempty"`
	Monitor            string `json:"monitor,omitempty"`
	Td                 int    `json:"td,omitempty"`
	Ownergroup         string `json:"ownergroup,omitempty"`
	Mgmt               bool   `json:"mgmt,omitempty"`
	Routetype          string `json:"routetype,omitempty"`
	Detail             bool   `json:"detail,omitempty"`
	Gatewayname        string `json:"gatewayname,omitempty"`
	Type               string `json:"type,omitempty"`
	Dynamic            string `json:"dynamic,omitempty"`
	Data               string `json:"data,omitempty"`
	Flags              string `json:"flags,omitempty"`
	State              string `json:"state,omitempty"`
	Totalprobes        string `json:"totalprobes,omitempty"`
	Totalfailedprobes  string `json:"totalfailedprobes,omitempty"`
	Failedprobes       string `json:"failedprobes,omitempty"`
	Monstatcode        string `json:"monstatcode,omitempty"`
	Monstatparam1      string `json:"monstatparam1,omitempty"`
	Monstatparam2      string `json:"monstatparam2,omitempty"`
	Monstatparam3      string `json:"monstatparam3,omitempty"`
	Data1              string `json:"data1,omitempty"`
	Routeowners        string `json:"routeowners,omitempty"`
	Retain             string `json:"retain,omitempty"`
	Static             string `json:"Static,omitempty"`
	Permanent          string `json:"permanent,omitempty"`
	Connected          string `json:"connected,omitempty"`
	Ospfv3             string `json:"ospfv3,omitempty"`
	Isis               string `json:"isis,omitempty"`
	Active             string `json:"active,omitempty"`
	Bgp                string `json:"bgp,omitempty"`
	Rip                string `json:"rip,omitempty"`
	Raroute            string `json:"raroute,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Vxlan struct {
	Id                 int    `json:"id,omitempty"`
	Vlan               int    `json:"vlan,omitempty"`
	Port               int    `json:"port,omitempty"`
	Dynamicrouting     string `json:"dynamicrouting,omitempty"`
	Ipv6dynamicrouting string `json:"ipv6dynamicrouting,omitempty"`
	Type               string `json:"type,omitempty"`
	Protocol           string `json:"protocol,omitempty"`
	Innervlantagging   string `json:"innervlantagging,omitempty"`
	Td                 string `json:"td,omitempty"`
	Partitionname      string `json:"partitionname,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Interfacepair struct {
	Id                 int      `json:"id,omitempty"`
	Ifnum              []string `json:"ifnum,omitempty"`
	Ifaces             string   `json:"ifaces,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
}

type Nd6 struct {
	Neighbor           string `json:"neighbor,omitempty"`
	Mac                string `json:"mac,omitempty"`
	Ifnum              string `json:"ifnum,omitempty"`
	Vlan               int    `json:"vlan,omitempty"`
	Vxlan              int    `json:"vxlan,omitempty"`
	Vtep               string `json:"vtep,omitempty"`
	Td                 int    `json:"td"`
	Nodeid             int    `json:"nodeid"`
	State              string `json:"state,omitempty"`
	Timeout            string `json:"timeout,omitempty"`
	Flags              string `json:"flags,omitempty"`
	Controlplane       string `json:"controlplane,omitempty"`
	Channel            string `json:"channel,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Netbridgeipbinding struct {
	Ipaddress string `json:"ipaddress,omitempty"`
	Netmask   string `json:"netmask,omitempty"`
	Name      string `json:"name,omitempty"`
}

type Vridinterfacebinding struct {
	Ifnum string `json:"ifnum,omitempty"`
	Vlan  int    `json:"vlan,omitempty"`
	Flags int    `json:"flags,omitempty"`
	Id    int    `json:"id,omitempty"`
}

type Vridnsipbinding struct {
	Ipaddress string `json:"ipaddress,omitempty"`
	Flags     int    `json:"flags,omitempty"`
	Id        int    `json:"id,omitempty"`
}

type Bridgegroupnsipbinding struct {
	Ipaddress  string `json:"ipaddress,omitempty"`
	Td         int    `json:"td,omitempty"`
	Netmask    string `json:"netmask,omitempty"`
	Rnat       bool   `json:"rnat,omitempty"`
	Ownergroup string `json:"ownergroup,omitempty"`
	Id         int    `json:"id,omitempty"`
}

type L2param struct {
	Mbfpeermacupdate        int    `json:"mbfpeermacupdate"`
	Maxbridgecollision      int    `json:"maxbridgecollision"`
	Bdggrpproxyarp          string `json:"bdggrpproxyarp,omitempty"`
	Bdgsetting              string `json:"bdgsetting,omitempty"`
	Garponvridintf          string `json:"garponvridintf,omitempty"`
	Macmodefwdmypkt         string `json:"macmodefwdmypkt,omitempty"`
	Usemymac                string `json:"usemymac,omitempty"`
	Proxyarp                string `json:"proxyarp,omitempty"`
	Garpreply               string `json:"garpreply,omitempty"`
	Mbfinstlearning         string `json:"mbfinstlearning,omitempty"`
	Rstintfonhafo           string `json:"rstintfonhafo,omitempty"`
	Skipproxyingbsdtraffic  string `json:"skipproxyingbsdtraffic,omitempty"`
	Returntoethernetsender  string `json:"returntoethernetsender,omitempty"`
	Stopmacmoveupdate       string `json:"stopmacmoveupdate,omitempty"`
	Bridgeagetimeout        int    `json:"bridgeagetimeout,omitempty"`
	Usenetprofilebsdtraffic string `json:"usenetprofilebsdtraffic,omitempty"`
	Nextgenapiresource      string `json:"_nextgenapiresource,omitempty"`
}

type L4param struct {
	L2connmethod       string `json:"l2connmethod,omitempty"`
	L4switch           string `json:"l4switch,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Ipsetip6binding struct {
	Ipaddress string `json:"ipaddress,omitempty"`
	Name      string `json:"name,omitempty"`
}

type Mapbmrbinding struct {
	Name string `json:"name,omitempty"`
}

type Nd6ravariables struct {
	Vlan                     int    `json:"vlan,omitempty"`
	Ceaserouteradv           string `json:"ceaserouteradv,omitempty"`
	Sendrouteradv            string `json:"sendrouteradv,omitempty"`
	Srclinklayeraddroption   string `json:"srclinklayeraddroption,omitempty"`
	Onlyunicastrtadvresponse string `json:"onlyunicastrtadvresponse,omitempty"`
	Managedaddrconfig        string `json:"managedaddrconfig,omitempty"`
	Otheraddrconfig          string `json:"otheraddrconfig,omitempty"`
	Currhoplimit             int    `json:"currhoplimit,omitempty"`
	Maxrtadvinterval         int    `json:"maxrtadvinterval,omitempty"`
	Minrtadvinterval         int    `json:"minrtadvinterval,omitempty"`
	Linkmtu                  int    `json:"linkmtu,omitempty"`
	Reachabletime            int    `json:"reachabletime,omitempty"`
	Retranstime              int    `json:"retranstime,omitempty"`
	Defaultlifetime          int    `json:"defaultlifetime,omitempty"`
	Lastrtadvtime            string `json:"lastrtadvtime,omitempty"`
	Nextrtadvdelay           string `json:"nextrtadvdelay,omitempty"`
	Nextgenapiresource       string `json:"_nextgenapiresource,omitempty"`
}

type Ip6tunnelparam struct {
	Srcip                string `json:"srcip,omitempty"`
	Dropfrag             string `json:"dropfrag,omitempty"`
	Dropfragcputhreshold int    `json:"dropfragcputhreshold,omitempty"`
	Srciproundrobin      string `json:"srciproundrobin,omitempty"`
	Useclientsourceipv6  string `json:"useclientsourceipv6,omitempty"`
	Nextgenapiresource   string `json:"_nextgenapiresource,omitempty"`
}

type Netprofilenatrulebinding struct {
	Natrule   string `json:"natrule,omitempty"`
	Netmask   string `json:"netmask,omitempty"`
	Rewriteip string `json:"rewriteip,omitempty"`
	Name      string `json:"name,omitempty"`
}

type Vlannsipbinding struct {
	Ipaddress  string `json:"ipaddress,omitempty"`
	Netmask    string `json:"netmask,omitempty"`
	Td         int    `json:"td,omitempty"`
	Ownergroup string `json:"ownergroup,omitempty"`
	Id         int    `json:"id,omitempty"`
}

type Vrid6ip6binding struct {
	Ipaddress string `json:"ipaddress,omitempty"`
	Flags     uint32 `json:"flags,omitempty"`
	Id        uint32 `json:"id,omitempty"`
}

type L3param struct {
	Srcnat               string `json:"srcnat,omitempty"`
	Icmpgenratethreshold int    `json:"icmpgenratethreshold,omitempty"`
	Overridernat         string `json:"overridernat,omitempty"`
	Dropdfflag           string `json:"dropdfflag,omitempty"`
	Miproundrobin        string `json:"miproundrobin,omitempty"`
	Externalloopback     string `json:"externalloopback,omitempty"`
	Tnlpmtuwoconn        string `json:"tnlpmtuwoconn,omitempty"`
	Usipserverstraypkt   string `json:"usipserverstraypkt,omitempty"`
	Forwardicmpfragments string `json:"forwardicmpfragments,omitempty"`
	Dropipfragments      string `json:"dropipfragments,omitempty"`
	Acllogtime           int    `json:"acllogtime,omitempty"`
	Implicitaclallow     string `json:"implicitaclallow,omitempty"`
	Dynamicrouting       string `json:"dynamicrouting,omitempty"`
	Ipv6dynamicrouting   string `json:"ipv6dynamicrouting,omitempty"`
	Allowclasseipv4      string `json:"allowclasseipv4,omitempty"`
	Implicitpbr          string `json:"implicitpbr,omitempty"`
	Nextgenapiresource   string `json:"_nextgenapiresource,omitempty"`
}

type Netbridge struct {
	Name               string `json:"name,omitempty"`
	Vxlanvlanmap       string `json:"vxlanvlanmap,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Vlanlinksetbinding struct {
	Ifnum      string `json:"ifnum,omitempty"`
	Tagged     bool   `json:"tagged,omitempty"`
	Id         int    `json:"id,omitempty"`
	Ownergroup string `json:"ownergroup,omitempty"`
}

type Vrid6channelbinding struct {
	Ifnum string `json:"ifnum,omitempty"`
	Vlan  int    `json:"vlan,omitempty"`
	Flags int    `json:"flags,omitempty"`
	Id    int    `json:"id,omitempty"`
}

type Ipsetnsip6binding struct {
	Ipaddress string `json:"ipaddress,omitempty"`
	Name      string `json:"name,omitempty"`
}

type Linkset struct {
	Id                 string `json:"id,omitempty"`
	Ifnum              string `json:"ifnum,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Linksetchannelbinding struct {
	Ifnum string `json:"ifnum,omitempty"`
	Id    string `json:"id,omitempty"`
}

type Fischannelbinding struct {
	Ifnum     string `json:"ifnum,omitempty"`
	Ownernode int    `json:"ownernode,omitempty"`
	Name      string `json:"name,omitempty"`
}

type Inat struct {
	Name               string `json:"name,omitempty"`
	Publicip           string `json:"publicip,omitempty"`
	Privateip          string `json:"privateip,omitempty"`
	Mode               string `json:"mode,omitempty"`
	Tcpproxy           string `json:"tcpproxy,omitempty"`
	Ftp                string `json:"ftp,omitempty"`
	Tftp               string `json:"tftp,omitempty"`
	Usip               string `json:"usip,omitempty"`
	Usnip              string `json:"usnip,omitempty"`
	Proxyip            string `json:"proxyip,omitempty"`
	Useproxyport       string `json:"useproxyport,omitempty"`
	Td                 int    `json:"td,omitempty"`
	Connfailover       string `json:"connfailover,omitempty"`
	Flags              string `json:"flags,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Netbridgevlanbinding struct {
	Vlan int    `json:"vlan,omitempty"`
	Name string `json:"name,omitempty"`
}

type Vrid6 struct {
	Id                   int    `json:"id,omitempty"`
	Priority             int    `json:"priority,omitempty"`
	Preemption           string `json:"preemption,omitempty"`
	Sharing              string `json:"sharing,omitempty"`
	Tracking             string `json:"tracking,omitempty"`
	Preemptiondelaytimer int    `json:"preemptiondelaytimer,omitempty"`
	Trackifnumpriority   int    `json:"trackifnumpriority,omitempty"`
	Ownernode            int    `json:"ownernode,omitempty"`
	All                  bool   `json:"all,omitempty"`
	Ifaces               string `json:"ifaces,omitempty"`
	Ifnum                string `json:"ifnum,omitempty"`
	Type                 string `json:"type,omitempty"`
	State                string `json:"state,omitempty"`
	Flags                string `json:"flags,omitempty"`
	Ipaddress            string `json:"ipaddress,omitempty"`
	Effectivepriority    string `json:"effectivepriority,omitempty"`
	Operationalownernode string `json:"operationalownernode,omitempty"`
	Nextgenapiresource   string `json:"_nextgenapiresource,omitempty"`
}

type Rnat6binding struct {
	Name string `json:"name,omitempty"`
}

type Vlanbinding struct {
	Id int `json:"id,omitempty"`
}

type Mapdomainmapbmrbinding struct {
	Mapbmrname string `json:"mapbmrname,omitempty"`
	Name       string `json:"name,omitempty"`
}

type Netbridgebinding struct {
	Name string `json:"name,omitempty"`
}

type Netbridgeiptunnelbinding struct {
	Tunnel string `json:"tunnel,omitempty"`
	Name   string `json:"name,omitempty"`
}

type Vxlanip6binding struct {
	Ipaddress string `json:"ipaddress,omitempty"`
	Id        uint32 `json:"id,omitempty"`
	Netmask   string `json:"netmask,omitempty"`
}

type Vxlanvlanmapbinding struct {
	Name string `json:"name,omitempty"`
}

type Fis struct {
	Name               string `json:"name,omitempty"`
	Ownernode          int    `json:"ownernode,omitempty"`
	Ifaces             string `json:"ifaces,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Vlannsip6binding struct {
	Ipaddress  string `json:"ipaddress,omitempty"`
	Td         int    `json:"td,omitempty"`
	Ownergroup string `json:"ownergroup,omitempty"`
	Id         int    `json:"id,omitempty"`
	Netmask    string `json:"netmask,omitempty"`
}

type Vrid6interfacebinding struct {
	Ifnum string `json:"ifnum,omitempty"`
	Vlan  int    `json:"vlan,omitempty"`
	Flags int    `json:"flags,omitempty"`
	Id    int    `json:"id,omitempty"`
}

type Vrid6nsip6binding struct {
	Ipaddress string `json:"ipaddress,omitempty"`
	Flags     int    `json:"flags,omitempty"`
	Id        int    `json:"id,omitempty"`
}

type Channel struct {
	Id                        string   `json:"id,omitempty"`
	Ifnum                     []string `json:"ifnum,omitempty"`
	State                     string   `json:"state,omitempty"`
	Mode                      string   `json:"mode,omitempty"`
	Conndistr                 string   `json:"conndistr,omitempty"`
	Macdistr                  string   `json:"macdistr,omitempty"`
	Lamac                     string   `json:"lamac,omitempty"`
	Speed                     string   `json:"speed,omitempty"`
	Flowctl                   string   `json:"flowctl,omitempty"`
	Hamonitor                 string   `json:"hamonitor,omitempty"`
	Haheartbeat               string   `json:"haheartbeat,omitempty"`
	Tagall                    string   `json:"tagall,omitempty"`
	Trunk                     string   `json:"trunk,omitempty"`
	Ifalias                   string   `json:"ifalias,omitempty"`
	Throughput                int      `json:"throughput,omitempty"`
	Bandwidthhigh             int      `json:"bandwidthhigh,omitempty"`
	Bandwidthnormal           int      `json:"bandwidthnormal,omitempty"`
	Mtu                       int      `json:"mtu,omitempty"`
	Lrminthroughput           int      `json:"lrminthroughput,omitempty"`
	Linkredundancy            string   `json:"linkredundancy,omitempty"`
	Devicename                string   `json:"devicename,omitempty"`
	Unit                      string   `json:"unit,omitempty"`
	Description               string   `json:"description,omitempty"`
	Flags                     string   `json:"flags,omitempty"`
	Actualmtu                 string   `json:"actualmtu,omitempty"`
	Vlan                      string   `json:"vlan,omitempty"`
	Mac                       string   `json:"mac,omitempty"`
	Uptime                    string   `json:"uptime,omitempty"`
	Downtime                  string   `json:"downtime,omitempty"`
	Reqmedia                  string   `json:"reqmedia,omitempty"`
	Reqspeed                  string   `json:"reqspeed,omitempty"`
	Reqduplex                 string   `json:"reqduplex,omitempty"`
	Reqflowcontrol            string   `json:"reqflowcontrol,omitempty"`
	Media                     string   `json:"media,omitempty"`
	Actspeed                  string   `json:"actspeed,omitempty"`
	Duplex                    string   `json:"duplex,omitempty"`
	Actflowctl                string   `json:"actflowctl,omitempty"`
	Lamode                    string   `json:"lamode,omitempty"`
	Autoneg                   string   `json:"autoneg,omitempty"`
	Autonegresult             string   `json:"autonegresult,omitempty"`
	Tagged                    string   `json:"tagged,omitempty"`
	Taggedany                 string   `json:"taggedany,omitempty"`
	Taggedautolearn           string   `json:"taggedautolearn,omitempty"`
	Hangdetect                string   `json:"hangdetect,omitempty"`
	Hangreset                 string   `json:"hangreset,omitempty"`
	Linkstate                 string   `json:"linkstate,omitempty"`
	Intfstate                 string   `json:"intfstate,omitempty"`
	Rxpackets                 string   `json:"rxpackets,omitempty"`
	Rxbytes                   string   `json:"rxbytes,omitempty"`
	Rxerrors                  string   `json:"rxerrors,omitempty"`
	Rxdrops                   string   `json:"rxdrops,omitempty"`
	Txpackets                 string   `json:"txpackets,omitempty"`
	Txbytes                   string   `json:"txbytes,omitempty"`
	Txerrors                  string   `json:"txerrors,omitempty"`
	Txdrops                   string   `json:"txdrops,omitempty"`
	Indisc                    string   `json:"indisc,omitempty"`
	Outdisc                   string   `json:"outdisc,omitempty"`
	Fctls                     string   `json:"fctls,omitempty"`
	Hangs                     string   `json:"hangs,omitempty"`
	Stsstalls                 string   `json:"stsstalls,omitempty"`
	Txstalls                  string   `json:"txstalls,omitempty"`
	Rxstalls                  string   `json:"rxstalls,omitempty"`
	Bdgmuted                  string   `json:"bdgmuted,omitempty"`
	Vmac                      string   `json:"vmac,omitempty"`
	Vmac6                     string   `json:"vmac6,omitempty"`
	Reqthroughput             string   `json:"reqthroughput,omitempty"`
	Actthroughput             string   `json:"actthroughput,omitempty"`
	Backplane                 string   `json:"backplane,omitempty"`
	Cleartime                 string   `json:"cleartime,omitempty"`
	Lacpmode                  string   `json:"lacpmode,omitempty"`
	Lacptimeout               string   `json:"lacptimeout,omitempty"`
	Lacpactorpriority         string   `json:"lacpactorpriority,omitempty"`
	Lacpactorportno           string   `json:"lacpactorportno,omitempty"`
	Lacppartnerstate          string   `json:"lacppartnerstate,omitempty"`
	Lacppartnertimeout        string   `json:"lacppartnertimeout,omitempty"`
	Lacppartneraggregation    string   `json:"lacppartneraggregation,omitempty"`
	Lacppartnerinsync         string   `json:"lacppartnerinsync,omitempty"`
	Lacppartnercollecting     string   `json:"lacppartnercollecting,omitempty"`
	Lacppartnerdistributing   string   `json:"lacppartnerdistributing,omitempty"`
	Lacppartnerdefaulted      string   `json:"lacppartnerdefaulted,omitempty"`
	Lacppartnerexpired        string   `json:"lacppartnerexpired,omitempty"`
	Lacppartnerpriority       string   `json:"lacppartnerpriority,omitempty"`
	Lacppartnersystemmac      string   `json:"lacppartnersystemmac,omitempty"`
	Lacppartnersystempriority string   `json:"lacppartnersystempriority,omitempty"`
	Lacppartnerportno         string   `json:"lacppartnerportno,omitempty"`
	Lacppartnerkey            string   `json:"lacppartnerkey,omitempty"`
	Lacpactoraggregation      string   `json:"lacpactoraggregation,omitempty"`
	Lacpactorinsync           string   `json:"lacpactorinsync,omitempty"`
	Lacpactorcollecting       string   `json:"lacpactorcollecting,omitempty"`
	Lacpactordistributing     string   `json:"lacpactordistributing,omitempty"`
	Lacpportmuxstate          string   `json:"lacpportmuxstate,omitempty"`
	Lacpportrxstat            string   `json:"lacpportrxstat,omitempty"`
	Lacpportselectstate       string   `json:"lacpportselectstate,omitempty"`
	Lldpmode                  string   `json:"lldpmode,omitempty"`
	Nextgenapiresource        string   `json:"_nextgenapiresource,omitempty"`
}

type Fisinterfacebinding struct {
	Ifnum     string `json:"ifnum,omitempty"`
	Ownernode int    `json:"ownernode,omitempty"`
	Name      string `json:"name,omitempty"`
}

type Nd6ravariablesbinding struct {
	Vlan int `json:"vlan,omitempty"`
}

type Netbridgensipbinding struct {
	Ipaddress string `json:"ipaddress,omitempty"`
	Netmask   string `json:"netmask,omitempty"`
	Name      string `json:"name,omitempty"`
}

type Vrid6ipbinding struct {
	Ipaddress string `json:"ipaddress,omitempty"`
	Flags     uint32 `json:"flags,omitempty"`
	Id        uint32 `json:"id,omitempty"`
}

type Vridip6binding struct {
	Ipaddress string `json:"ipaddress,omitempty"`
	Flags     uint32 `json:"flags,omitempty"`
	Id        uint32 `json:"id,omitempty"`
}

type Bridgetable struct {
	Mac                string `json:"mac,omitempty"`
	Vxlan              int    `json:"vxlan,omitempty"`
	Vtep               string `json:"vtep,omitempty"`
	Vni                int    `json:"vni,omitempty"`
	Devicevlan         int    `json:"devicevlan,omitempty"`
	Bridgeage          int    `json:"bridgeage,omitempty"`
	Nodeid             int    `json:"nodeid,omitempty"`
	Vlan               int    `json:"vlan,omitempty"`
	Ifnum              string `json:"ifnum,omitempty"`
	Flags              string `json:"flags,omitempty"`
	Type               string `json:"type,omitempty"`
	Channel            string `json:"channel,omitempty"`
	Controlplane       string `json:"controlplane,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Channelbinding struct {
	Id string `json:"id,omitempty"`
}

type Ipset struct {
	Name               string `json:"name,omitempty"`
	Td                 int    `json:"td,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Mapbmrbmrv4networkbinding struct {
	Network string `json:"network,omitempty"`
	Netmask string `json:"netmask,omitempty"`
	Name    string `json:"name,omitempty"`
}

type Vlan struct {
	Id                 int    `json:"id,omitempty"`
	Aliasname          string `json:"aliasname,omitempty"`
	Dynamicrouting     string `json:"dynamicrouting,omitempty"`
	Ipv6dynamicrouting string `json:"ipv6dynamicrouting,omitempty"`
	Mtu                int    `json:"mtu,omitempty"`
	Sharing            string `json:"sharing,omitempty"`
	Linklocalipv6addr  string `json:"linklocalipv6addr,omitempty"`
	Rnat               string `json:"rnat,omitempty"`
	Portbitmap         string `json:"portbitmap,omitempty"`
	Lsbitmap           string `json:"lsbitmap,omitempty"`
	Tagbitmap          string `json:"tagbitmap,omitempty"`
	Lstagbitmap        string `json:"lstagbitmap,omitempty"`
	Ifaces             string `json:"ifaces,omitempty"`
	Tagifaces          string `json:"tagifaces,omitempty"`
	Ifnum              string `json:"ifnum,omitempty"`
	Tagged             string `json:"tagged,omitempty"`
	Vlantd             string `json:"vlantd,omitempty"`
	Sdxvlan            string `json:"sdxvlan,omitempty"`
	Partitionname      string `json:"partitionname,omitempty"`
	Vxlan              string `json:"vxlan,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Vxlanbinding struct {
	Id int `json:"id,omitempty"`
}

type Bridgegroupbinding struct {
	Id int `json:"id,omitempty"`
}

type Linksetinterfacebinding struct {
	Ifnum string `json:"ifnum,omitempty"`
	Id    string `json:"id,omitempty"`
}

type Mapdomainbinding struct {
	Name string `json:"name,omitempty"`
}

type Linksetbinding struct {
	Id string `json:"id,omitempty"`
}

type Nat64 struct {
	Name               string `json:"name,omitempty"`
	Acl6name           string `json:"acl6name,omitempty"`
	Netprofile         string `json:"netprofile,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nat64param struct {
	Td                 int    `json:"td"`
	Nat64ignoretos     string `json:"nat64ignoretos,omitempty"`
	Nat64zerochecksum  string `json:"nat64zerochecksum,omitempty"`
	Nat64v6mtu         int    `json:"nat64v6mtu,omitempty"`
	Nat64fragheader    string `json:"nat64fragheader,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Rnat6nsip6binding struct {
	Natip6     string `json:"natip6,omitempty"`
	Td         int    `json:"td,omitempty"`
	Ownergroup string `json:"ownergroup,omitempty"`
	Name       string `json:"name,omitempty"`
}

type Rnatretainsourceportsetbinding struct {
	Retainsourceportrange string `json:"retainsourceportrange,omitempty"`
	Name                  string `json:"name,omitempty"`
}

type Vrid struct {
	Id                   int    `json:"id,omitempty"`
	Priority             int    `json:"priority,omitempty"`
	Preemption           string `json:"preemption,omitempty"`
	Sharing              string `json:"sharing,omitempty"`
	Tracking             string `json:"tracking,omitempty"`
	Ownernode            int    `json:"ownernode,omitempty"`
	Trackifnumpriority   int    `json:"trackifnumpriority,omitempty"`
	Preemptiondelaytimer int    `json:"preemptiondelaytimer,omitempty"`
	All                  bool   `json:"all,omitempty"`
	Ifaces               string `json:"ifaces,omitempty"`
	Type                 string `json:"type,omitempty"`
	Effectivepriority    string `json:"effectivepriority,omitempty"`
	Flags                string `json:"flags,omitempty"`
	Ipaddress            string `json:"ipaddress,omitempty"`
	State                string `json:"state,omitempty"`
	Operationalownernode string `json:"operationalownernode,omitempty"`
	Nextgenapiresource   string `json:"_nextgenapiresource,omitempty"`
}

type Forwardingsession struct {
	Name               string `json:"name,omitempty"`
	Network            string `json:"network,omitempty"`
	Netmask            string `json:"netmask,omitempty"`
	Acl6name           string `json:"acl6name,omitempty"`
	Aclname            string `json:"aclname,omitempty"`
	Td                 int    `json:"td,omitempty"`
	Connfailover       string `json:"connfailover,omitempty"`
	Sourceroutecache   string `json:"sourceroutecache,omitempty"`
	Processlocal       string `json:"processlocal,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Route struct {
	Network            string   `json:"network,omitempty"`
	Netmask            string   `json:"netmask,omitempty"`
	Gateway            string   `json:"gateway,omitempty"`
	Vlan               int      `json:"vlan,omitempty"`
	Cost               int      `json:"cost,omitempty"`
	Td                 int      `json:"td,omitempty"`
	Distance           int      `json:"distance,omitempty"`
	Cost1              int      `json:"cost1,omitempty"`
	Weight             int      `json:"weight,omitempty"`
	Advertise          string   `json:"advertise,omitempty"`
	Protocol           []string `json:"protocol,omitempty"`
	Msr                string   `json:"msr,omitempty"`
	Monitor            string   `json:"monitor,omitempty"`
	Ownergroup         string   `json:"ownergroup,omitempty"`
	Mgmt               bool     `json:"mgmt,omitempty"`
	Routetype          string   `json:"routetype,omitempty"`
	Detail             bool     `json:"detail,omitempty"`
	Gatewayname        string   `json:"gatewayname,omitempty"`
	Type               string   `json:"type,omitempty"`
	Dynamic            string   `json:"dynamic,omitempty"`
	Static             string   `json:"Static,omitempty"`
	Permanent          string   `json:"permanent,omitempty"`
	Direct             string   `json:"direct,omitempty"`
	Nat                string   `json:"nat,omitempty"`
	Lbroute            string   `json:"lbroute,omitempty"`
	Adv                string   `json:"adv,omitempty"`
	Tunnel             string   `json:"tunnel,omitempty"`
	Data               string   `json:"data,omitempty"`
	Data0              string   `json:"data0,omitempty"`
	Flags              string   `json:"flags,omitempty"`
	Routeowners        string   `json:"routeowners,omitempty"`
	Retain             string   `json:"retain,omitempty"`
	Ospf               string   `json:"ospf,omitempty"`
	Isis               string   `json:"isis,omitempty"`
	Rip                string   `json:"rip,omitempty"`
	Bgp                string   `json:"bgp,omitempty"`
	Dhcp               string   `json:"dhcp,omitempty"`
	Advospf            string   `json:"advospf,omitempty"`
	Advisis            string   `json:"advisis,omitempty"`
	Advrip             string   `json:"advrip,omitempty"`
	Advbgp             string   `json:"advbgp,omitempty"`
	State              string   `json:"state,omitempty"`
	Totalprobes        string   `json:"totalprobes,omitempty"`
	Totalfailedprobes  string   `json:"totalfailedprobes,omitempty"`
	Failedprobes       string   `json:"failedprobes,omitempty"`
	Monstatcode        string   `json:"monstatcode,omitempty"`
	Monstatparam1      string   `json:"monstatparam1,omitempty"`
	Monstatparam2      string   `json:"monstatparam2,omitempty"`
	Monstatparam3      string   `json:"monstatparam3,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
}

type Arpparam struct {
	Timeout            int    `json:"timeout,omitempty"`
	Spoofvalidation    string `json:"spoofvalidation,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Vridipbinding struct {
	Ipaddress string `json:"ipaddress,omitempty"`
	Flags     uint32 `json:"flags,omitempty"`
	Id        uint32 `json:"id,omitempty"`
}

type Ipsetbinding struct {
	Name string `json:"name,omitempty"`
}

type Interface struct {
	Id                        string   `json:"id,omitempty"`
	Speed                     string   `json:"speed,omitempty"`
	Duplex                    string   `json:"duplex,omitempty"`
	Flowctl                   string   `json:"flowctl,omitempty"`
	Autoneg                   string   `json:"autoneg,omitempty"`
	Hamonitor                 string   `json:"hamonitor,omitempty"`
	Haheartbeat               string   `json:"haheartbeat,omitempty"`
	Mtu                       int      `json:"mtu,omitempty"`
	Ringsize                  int      `json:"ringsize,omitempty"`
	Ringtype                  string   `json:"ringtype,omitempty"`
	Tagall                    string   `json:"tagall,omitempty"`
	Trunk                     string   `json:"trunk,omitempty"`
	Trunkmode                 string   `json:"trunkmode,omitempty"`
	Trunkallowedvlan          []string `json:"trunkallowedvlan,omitempty"`
	Lacpmode                  string   `json:"lacpmode,omitempty"`
	Lacpkey                   int      `json:"lacpkey,omitempty"`
	Lagtype                   string   `json:"lagtype,omitempty"`
	Lacppriority              int      `json:"lacppriority,omitempty"`
	Lacptimeout               string   `json:"lacptimeout,omitempty"`
	Ifalias                   string   `json:"ifalias,omitempty"`
	Throughput                int      `json:"throughput,omitempty"`
	Linkredundancy            string   `json:"linkredundancy,omitempty"`
	Bandwidthhigh             int      `json:"bandwidthhigh,omitempty"`
	Bandwidthnormal           int      `json:"bandwidthnormal,omitempty"`
	Lldpmode                  string   `json:"lldpmode,omitempty"`
	Lrsetpriority             int      `json:"lrsetpriority,omitempty"`
	Devicename                string   `json:"devicename,omitempty"`
	Unit                      string   `json:"unit,omitempty"`
	Description               string   `json:"description,omitempty"`
	Flags                     string   `json:"flags,omitempty"`
	Actualmtu                 string   `json:"actualmtu,omitempty"`
	Vlan                      string   `json:"vlan,omitempty"`
	Mac                       string   `json:"mac,omitempty"`
	Uptime                    string   `json:"uptime,omitempty"`
	Downtime                  string   `json:"downtime,omitempty"`
	Actualringsize            string   `json:"actualringsize,omitempty"`
	Reqmedia                  string   `json:"reqmedia,omitempty"`
	Reqspeed                  string   `json:"reqspeed,omitempty"`
	Reqduplex                 string   `json:"reqduplex,omitempty"`
	Reqflowcontrol            string   `json:"reqflowcontrol,omitempty"`
	Actmedia                  string   `json:"actmedia,omitempty"`
	Actspeed                  string   `json:"actspeed,omitempty"`
	Actduplex                 string   `json:"actduplex,omitempty"`
	Actflowctl                string   `json:"actflowctl,omitempty"`
	Mode                      string   `json:"mode,omitempty"`
	State                     string   `json:"state,omitempty"`
	Autonegresult             string   `json:"autonegresult,omitempty"`
	Tagged                    string   `json:"tagged,omitempty"`
	Taggedany                 string   `json:"taggedany,omitempty"`
	Taggedautolearn           string   `json:"taggedautolearn,omitempty"`
	Hangdetect                string   `json:"hangdetect,omitempty"`
	Hangreset                 string   `json:"hangreset,omitempty"`
	Linkstate                 string   `json:"linkstate,omitempty"`
	Intfstate                 string   `json:"intfstate,omitempty"`
	Rxpackets                 string   `json:"rxpackets,omitempty"`
	Rxbytes                   string   `json:"rxbytes,omitempty"`
	Rxerrors                  string   `json:"rxerrors,omitempty"`
	Rxdrops                   string   `json:"rxdrops,omitempty"`
	Txpackets                 string   `json:"txpackets,omitempty"`
	Txbytes                   string   `json:"txbytes,omitempty"`
	Txerrors                  string   `json:"txerrors,omitempty"`
	Txdrops                   string   `json:"txdrops,omitempty"`
	Indisc                    string   `json:"indisc,omitempty"`
	Outdisc                   string   `json:"outdisc,omitempty"`
	Fctls                     string   `json:"fctls,omitempty"`
	Hangs                     string   `json:"hangs,omitempty"`
	Stsstalls                 string   `json:"stsstalls,omitempty"`
	Txstalls                  string   `json:"txstalls,omitempty"`
	Rxstalls                  string   `json:"rxstalls,omitempty"`
	Bdgmacmoved               string   `json:"bdgmacmoved,omitempty"`
	Bdgmuted                  string   `json:"bdgmuted,omitempty"`
	Vmac                      string   `json:"vmac,omitempty"`
	Vmac6                     string   `json:"vmac6,omitempty"`
	Reqthroughput             string   `json:"reqthroughput,omitempty"`
	Actthroughput             string   `json:"actthroughput,omitempty"`
	Backplane                 string   `json:"backplane,omitempty"`
	Ifnum                     string   `json:"ifnum,omitempty"`
	Cleartime                 string   `json:"cleartime,omitempty"`
	Slavestate                string   `json:"slavestate,omitempty"`
	Slavemedia                string   `json:"slavemedia,omitempty"`
	Slavespeed                string   `json:"slavespeed,omitempty"`
	Slaveduplex               string   `json:"slaveduplex,omitempty"`
	Slaveflowctl              string   `json:"slaveflowctl,omitempty"`
	Slavetime                 string   `json:"slavetime,omitempty"`
	Intftype                  string   `json:"intftype,omitempty"`
	Svmcmd                    string   `json:"svmcmd,omitempty"`
	Lacpactormode             string   `json:"lacpactormode,omitempty"`
	Lacpactortimeout          string   `json:"lacpactortimeout,omitempty"`
	Lacpactorpriority         string   `json:"lacpactorpriority,omitempty"`
	Lacpactorportno           string   `json:"lacpactorportno,omitempty"`
	Lacppartnerstate          string   `json:"lacppartnerstate,omitempty"`
	Lacppartnertimeout        string   `json:"lacppartnertimeout,omitempty"`
	Lacppartneraggregation    string   `json:"lacppartneraggregation,omitempty"`
	Lacppartnerinsync         string   `json:"lacppartnerinsync,omitempty"`
	Lacppartnercollecting     string   `json:"lacppartnercollecting,omitempty"`
	Lacppartnerdistributing   string   `json:"lacppartnerdistributing,omitempty"`
	Lacppartnerdefaulted      string   `json:"lacppartnerdefaulted,omitempty"`
	Lacppartnerexpired        string   `json:"lacppartnerexpired,omitempty"`
	Lacppartnerpriority       string   `json:"lacppartnerpriority,omitempty"`
	Lacppartnersystemmac      string   `json:"lacppartnersystemmac,omitempty"`
	Lacppartnersystempriority string   `json:"lacppartnersystempriority,omitempty"`
	Lacppartnerportno         string   `json:"lacppartnerportno,omitempty"`
	Lacppartnerkey            string   `json:"lacppartnerkey,omitempty"`
	Lacpactoraggregation      string   `json:"lacpactoraggregation,omitempty"`
	Lacpactorinsync           string   `json:"lacpactorinsync,omitempty"`
	Lacpactorcollecting       string   `json:"lacpactorcollecting,omitempty"`
	Lacpactordistributing     string   `json:"lacpactordistributing,omitempty"`
	Lacpportmuxstate          string   `json:"lacpportmuxstate,omitempty"`
	Lacpportrxstat            string   `json:"lacpportrxstat,omitempty"`
	Lacpportselectstate       string   `json:"lacpportselectstate,omitempty"`
	Lractiveintf              string   `json:"lractiveintf,omitempty"`
	Nextgenapiresource        string   `json:"_nextgenapiresource,omitempty"`
}

type Rnat6ip6binding struct {
	Natip6     string `json:"natip6,omitempty"`
	Td         uint32 `json:"td,omitempty"`
	Ownergroup string `json:"ownergroup,omitempty"`
	Name       string `json:"name,omitempty"`
}

type Vrid6binding struct {
	Id int `json:"id,omitempty"`
}

type Vxlanipbinding struct {
	Ipaddress string `json:"ipaddress,omitempty"`
	Netmask   string `json:"netmask,omitempty"`
	Id        uint32 `json:"id,omitempty"`
}

type Vxlansrcipbinding struct {
	Srcip string `json:"srcip,omitempty"`
	Id    int    `json:"id,omitempty"`
}

type Vxlanvlanmap struct {
	Name               string `json:"name,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Ip6tunnel struct {
	Name               string `json:"name,omitempty"`
	Remote             string `json:"remote,omitempty"`
	Local              string `json:"local,omitempty"`
	Ownergroup         string `json:"ownergroup,omitempty"`
	Remoteip           string `json:"remoteip,omitempty"`
	Type               string `json:"type,omitempty"`
	Encapip            string `json:"encapip,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Rsskeytype struct {
	Rsstype            string `json:"rsstype,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Bridgegroupip6binding struct {
	Ipaddress  string `json:"ipaddress,omitempty"`
	Td         uint32 `json:"td,omitempty"`
	Rnat       bool   `json:"rnat,omitempty"`
	Ownergroup string `json:"ownergroup,omitempty"`
	Id         uint32 `json:"id,omitempty"`
	Netmask    string `json:"netmask,omitempty"`
}

type Vridparam struct {
	Sendtomaster       string `json:"sendtomaster,omitempty"`
	Hellointerval      int    `json:"hellointerval,omitempty"`
	Deadinterval       int    `json:"deadinterval,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Nd6ravariablesonlinkipv6prefixbinding struct {
	Ipv6prefix string `json:"ipv6prefix,omitempty"`
	Vlan       int    `json:"vlan,omitempty"`
}

type Bridgegroupvlanbinding struct {
	Vlan int  `json:"vlan,omitempty"`
	Rnat bool `json:"rnat,omitempty"`
	Id   int  `json:"id,omitempty"`
}

type Rnatglobalauditsyslogpolicybinding struct {
	Policy   string `json:"policy,omitempty"`
	Priority int    `json:"priority,omitempty"`
	All      bool   `json:"all,omitempty"`
}

type Vlanipbinding struct {
	Ipaddress  string `json:"ipaddress,omitempty"`
	Netmask    string `json:"netmask,omitempty"`
	Td         uint32 `json:"td,omitempty"`
	Ownergroup string `json:"ownergroup,omitempty"`
	Id         uint32 `json:"id,omitempty"`
}

type Iptunnelparam struct {
	Srcip                string `json:"srcip,omitempty"`
	Dropfrag             string `json:"dropfrag,omitempty"`
	Dropfragcputhreshold int    `json:"dropfragcputhreshold,omitempty"`
	Srciproundrobin      string `json:"srciproundrobin,omitempty"`
	Enablestrictrx       string `json:"enablestrictrx,omitempty"`
	Enablestricttx       string `json:"enablestricttx,omitempty"`
	Mac                  string `json:"mac,omitempty"`
	Useclientsourceip    string `json:"useclientsourceip,omitempty"`
	Nextgenapiresource   string `json:"_nextgenapiresource,omitempty"`
}

type Mapdmr struct {
	Name               string `json:"name,omitempty"`
	Bripv6prefix       string `json:"bripv6prefix,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Netbridgeip6binding struct {
	Ipaddress string `json:"ipaddress,omitempty"`
	Netmask   string `json:"netmask,omitempty"`
	Name      string `json:"name,omitempty"`
}

type Vrid6trackinterfacebinding struct {
	Trackifnum string `json:"trackifnum,omitempty"`
	Flags      int    `json:"flags,omitempty"`
	Id         int    `json:"id,omitempty"`
}

type Mapdomain struct {
	Name               string `json:"name,omitempty"`
	Mapdmrname         string `json:"mapdmrname,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Arp struct {
	Ipaddress          string `json:"ipaddress,omitempty"`
	Td                 int    `json:"td"`
	Mac                string `json:"mac,omitempty"`
	Ifnum              string `json:"ifnum,omitempty"`
	Vxlan              int    `json:"vxlan,omitempty"`
	Vtep               string `json:"vtep,omitempty"`
	Vlan               int    `json:"vlan,omitempty"`
	Ownernode          int    `json:"ownernode"`
	All                bool   `json:"all,omitempty"`
	Nodeid             int    `json:"nodeid"`
	Timeout            string `json:"timeout,omitempty"`
	State              string `json:"state,omitempty"`
	Flags              string `json:"flags,omitempty"`
	Type               string `json:"type,omitempty"`
	Channel            string `json:"channel,omitempty"`
	Controlplane       string `json:"controlplane,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Lacp struct {
	Syspriority        int    `json:"syspriority,omitempty"`
	Ownernode          int    `json:"ownernode"`
	Devicename         string `json:"devicename,omitempty"`
	Mac                string `json:"mac,omitempty"`
	Flags              string `json:"flags,omitempty"`
	Lacpkey            string `json:"lacpkey,omitempty"`
	Clustersyspriority string `json:"clustersyspriority,omitempty"`
	Clustermac         string `json:"clustermac,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Onlinkipv6prefix struct {
	Ipv6prefix               string `json:"ipv6prefix,omitempty"`
	Onlinkprefix             string `json:"onlinkprefix,omitempty"`
	Autonomusprefix          string `json:"autonomusprefix,omitempty"`
	Depricateprefix          string `json:"depricateprefix,omitempty"`
	Decrementprefixlifetimes string `json:"decrementprefixlifetimes,omitempty"`
	Prefixvalidelifetime     int    `json:"prefixvalidelifetime,omitempty"`
	Prefixpreferredlifetime  int    `json:"prefixpreferredlifetime,omitempty"`
	Prefixcurrvalidelft      string `json:"prefixcurrvalidelft,omitempty"`
	Prefixcurrpreferredlft   string `json:"prefixcurrpreferredlft,omitempty"`
	Nextgenapiresource       string `json:"_nextgenapiresource,omitempty"`
}

type Vlanchannelbinding struct {
	Ifnum      string `json:"ifnum,omitempty"`
	Tagged     bool   `json:"tagged,omitempty"`
	Id         int    `json:"id,omitempty"`
	Ownergroup string `json:"ownergroup,omitempty"`
}

type Fisbinding struct {
	Name string `json:"name,omitempty"`
}

type Ipsetipbinding struct {
	Ipaddress string `json:"ipaddress,omitempty"`
	Name      string `json:"name,omitempty"`
}

type Ipv6 struct {
	Ralearning           string `json:"ralearning,omitempty"`
	Routerredirection    string `json:"routerredirection,omitempty"`
	Ndbasereachtime      int    `json:"ndbasereachtime,omitempty"`
	Ndretransmissiontime int    `json:"ndretransmissiontime,omitempty"`
	Natprefix            string `json:"natprefix,omitempty"`
	Td                   int    `json:"td"`
	Dodad                string `json:"dodad,omitempty"`
	Usipnatprefix        string `json:"usipnatprefix,omitempty"`
	Basereachtime        string `json:"basereachtime,omitempty"`
	Reachtime            string `json:"reachtime,omitempty"`
	Ndreachtime          string `json:"ndreachtime,omitempty"`
	Retransmissiontime   string `json:"retransmissiontime,omitempty"`
	Nextgenapiresource   string `json:"_nextgenapiresource,omitempty"`
}

type Rnatbinding struct {
	Name string `json:"name,omitempty"`
}

type Vxlaniptunnelbinding struct {
	Id     int    `json:"id,omitempty"`
	Tunnel string `json:"tunnel,omitempty"`
}

type Rnatglobalbinding struct {
}

type Vlanip6binding struct {
	Ipaddress  string `json:"ipaddress,omitempty"`
	Td         uint32 `json:"td,omitempty"`
	Ownergroup string `json:"ownergroup,omitempty"`
	Id         uint32 `json:"id,omitempty"`
	Netmask    string `json:"netmask,omitempty"`
}

type Bridgegroupipbinding struct {
	Ipaddress  string `json:"ipaddress,omitempty"`
	Td         uint32 `json:"td,omitempty"`
	Netmask    string `json:"netmask,omitempty"`
	Rnat       bool   `json:"rnat,omitempty"`
	Ownergroup string `json:"ownergroup,omitempty"`
	Id         uint32 `json:"id,omitempty"`
}

type Iptunnel struct {
	Name               string `json:"name,omitempty"`
	Remote             string `json:"remote,omitempty"`
	Remotesubnetmask   string `json:"remotesubnetmask,omitempty"`
	Local              string `json:"local,omitempty"`
	Protocol           string `json:"protocol,omitempty"`
	Vnid               int    `json:"vnid,omitempty"`
	Vlantagging        string `json:"vlantagging,omitempty"`
	Destport           int    `json:"destport,omitempty"`
	Tosinherit         string `json:"tosinherit,omitempty"`
	Grepayload         string `json:"grepayload,omitempty"`
	Ipsecprofilename   string `json:"ipsecprofilename,omitempty"`
	Vlan               int    `json:"vlan,omitempty"`
	Ownergroup         string `json:"ownergroup,omitempty"`
	Sysname            string `json:"sysname,omitempty"`
	Type               string `json:"type,omitempty"`
	Encapip            string `json:"encapip,omitempty"`
	Channel            string `json:"channel,omitempty"`
	Tunneltype         string `json:"tunneltype,omitempty"`
	Ipsectunnelstatus  string `json:"ipsectunnelstatus,omitempty"`
	Refcnt             string `json:"refcnt,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Netbridgensip6binding struct {
	Ipaddress string `json:"ipaddress,omitempty"`
	Netmask   string `json:"netmask,omitempty"`
	Name      string `json:"name,omitempty"`
}

type Portallocation struct {
	Srcip              string `json:"srcip,omitempty"`
	Destip             string `json:"destip,omitempty"`
	Destport           int    `json:"destport,omitempty"`
	Protocol           int    `json:"protocol,omitempty"`
	Nodeid             int    `json:"nodeid,omitempty"`
	Freeports          string `json:"freeports,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Rnat struct {
	Network            string `json:"network,omitempty"`
	Netmask            string `json:"netmask,omitempty"`
	Aclname            string `json:"aclname,omitempty"`
	Td                 int    `json:"td"`
	Ownergroup         string `json:"ownergroup,omitempty"`
	Name               string `json:"name,omitempty"`
	Redirectport       int    `json:"redirectport,omitempty"`
	Natip              string `json:"natip,omitempty"`
	Srcippersistency   string `json:"srcippersistency,omitempty"`
	Useproxyport       string `json:"useproxyport,omitempty"`
	Connfailover       string `json:"connfailover,omitempty"`
	Newname            string `json:"newname,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Bridgegroup struct {
	Id                 int    `json:"id,omitempty"`
	Dynamicrouting     string `json:"dynamicrouting,omitempty"`
	Ipv6dynamicrouting string `json:"ipv6dynamicrouting,omitempty"`
	Flags              string `json:"flags,omitempty"`
	Portbitmap         string `json:"portbitmap,omitempty"`
	Tagbitmap          string `json:"tagbitmap,omitempty"`
	Ifaces             string `json:"ifaces,omitempty"`
	Tagifaces          string `json:"tagifaces,omitempty"`
	Rnat               string `json:"rnat,omitempty"`
	Partitionname      string `json:"partitionname,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Vxlanvlanmapvxlanbinding struct {
	Vxlan int      `json:"vxlan,omitempty"`
	Vlan  []string `json:"vlan,omitempty"`
	Name  string   `json:"name,omitempty"`
}

type Rnatsession struct {
	Network string `json:"network,omitempty"`
	Netmask string `json:"netmask,omitempty"`
	Natip   string `json:"natip,omitempty"`
	Aclname string `json:"aclname,omitempty"`
}

type Vlaninterfacebinding struct {
	Ifnum      string `json:"ifnum,omitempty"`
	Tagged     bool   `json:"tagged,omitempty"`
	Id         int    `json:"id,omitempty"`
	Ownergroup string `json:"ownergroup,omitempty"`
}

type Vrid6nsipbinding struct {
	Ipaddress string `json:"ipaddress,omitempty"`
	Flags     int    `json:"flags,omitempty"`
	Id        int    `json:"id,omitempty"`
}

type Vridbinding struct {
	Id int `json:"id,omitempty"`
}

type Vxlannsip6binding struct {
	Ipaddress string `json:"ipaddress,omitempty"`
	Id        int    `json:"id,omitempty"`
	Netmask   string `json:"netmask,omitempty"`
}

type Bridgegroupnsip6binding struct {
	Ipaddress  string `json:"ipaddress,omitempty"`
	Td         int    `json:"td,omitempty"`
	Rnat       bool   `json:"rnat,omitempty"`
	Ownergroup string `json:"ownergroup,omitempty"`
	Id         int    `json:"id,omitempty"`
	Netmask    string `json:"netmask,omitempty"`
}

type Vridchannelbinding struct {
	Ifnum string `json:"ifnum,omitempty"`
	Vlan  int    `json:"vlan,omitempty"`
	Flags int    `json:"flags,omitempty"`
	Id    int    `json:"id,omitempty"`
}

type Vxlannsipbinding struct {
	Ipaddress string `json:"ipaddress,omitempty"`
	Netmask   string `json:"netmask,omitempty"`
	Id        int    `json:"id,omitempty"`
}

type Channelinterfacebinding struct {
	Ifnum        []string `json:"ifnum,omitempty"`
	Lamode       string   `json:"lamode,omitempty"`
	Slavestate   int      `json:"slavestate,omitempty"`
	Slavemedia   int      `json:"slavemedia,omitempty"`
	Slavespeed   int      `json:"slavespeed,omitempty"`
	Slaveduplex  int      `json:"slaveduplex,omitempty"`
	Slaveflowctl int      `json:"slaveflowctl,omitempty"`
	Slavetime    int      `json:"slavetime,omitempty"`
	Lractiveintf int      `json:"lractiveintf,omitempty"`
	Svmcmd       int      `json:"svmcmd,omitempty"`
	Id           string   `json:"id,omitempty"`
}

type Netprofilesrcportsetbinding struct {
	Srcportrange string `json:"srcportrange,omitempty"`
	Name         string `json:"name,omitempty"`
}

type Rnatparam struct {
	Tcpproxy           string `json:"tcpproxy,omitempty"`
	Srcippersistency   string `json:"srcippersistency,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Vridnsip6binding struct {
	Ipaddress string `json:"ipaddress,omitempty"`
	Flags     int    `json:"flags,omitempty"`
	Id        int    `json:"id,omitempty"`
}

type Appalgparam struct {
	Pptpgreidletimeout int    `json:"pptpgreidletimeout,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Ptp struct {
	State              string `json:"state,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Rnatipbinding struct {
	Natip      string `json:"natip,omitempty"`
	Td         uint32 `json:"td,omitempty"`
	Ownergroup string `json:"ownergroup,omitempty"`
	Name       string `json:"name,omitempty"`
}

type Vridtrackinterfacebinding struct {
	Trackifnum string `json:"trackifnum,omitempty"`
	Flags      int    `json:"flags,omitempty"`
	Id         int    `json:"id,omitempty"`
}
