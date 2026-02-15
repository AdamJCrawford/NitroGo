package models

// network configuration structs
type ChannelInterfaceBinding struct {
	Id           string   `json:"id,omitempty"`
	Ifnum        []string `json:"ifnum,omitempty"`
	Lamode       string   `json:"lamode,omitempty"`
	Lractiveintf int      `json:"lractiveintf,omitempty"`
	Slaveduplex  int      `json:"slaveduplex,omitempty"`
	Slaveflowctl int      `json:"slaveflowctl,omitempty"`
	Slavemedia   int      `json:"slavemedia,omitempty"`
	Slavespeed   int      `json:"slavespeed,omitempty"`
	Slavestate   int      `json:"slavestate,omitempty"`
	Slavetime    int      `json:"slavetime,omitempty"`
	Svmcmd       int      `json:"svmcmd,omitempty"`
}

type Mapdmr struct {
	Bripv6prefix       string  `json:"bripv6prefix,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
}

type Rnat6Binding struct {
	Name              string        `json:"name,omitempty"`
	Rnat6Nsip6Binding []interface{} `json:"rnat6_nsip6_binding,omitempty"`
}

type VlanNsip6Binding struct {
	Id         int    `json:"id,omitempty"`
	Ipaddress  string `json:"ipaddress,omitempty"`
	Netmask    string `json:"netmask,omitempty"`
	Ownergroup string `json:"ownergroup,omitempty"`
	Td         int    `json:"td,omitempty"`
}

type Nd6ravariablesOnlinkipv6prefixBinding struct {
	Ipv6prefix string `json:"ipv6prefix,omitempty"`
	Vlan       int    `json:"vlan,omitempty"`
}

type Ptp struct {
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	State              string `json:"state,omitempty"`
}

type RnatNsipBinding struct {
	Name       string `json:"name,omitempty"`
	Natip      string `json:"natip,omitempty"`
	Ownergroup string `json:"ownergroup,omitempty"`
	Td         int    `json:"td,omitempty"`
}

type Vrid6 struct {
	All                  bool    `json:"all,omitempty"`
	Count                float64 `json:"__count,omitempty"`
	Effectivepriority    int     `json:"effectivepriority,omitempty"`
	Flags                int     `json:"flags,omitempty"`
	Id                   int     `json:"id,omitempty"`
	Ifaces               string  `json:"ifaces,omitempty"`
	Ifnum                string  `json:"ifnum,omitempty"`
	Ipaddress            string  `json:"ipaddress,omitempty"`
	Nextgenapiresource   string  `json:"_nextgenapiresource,omitempty"`
	Operationalownernode int     `json:"operationalownernode,omitempty"`
	Ownernode            int     `json:"ownernode,omitempty"`
	Preemption           string  `json:"preemption,omitempty"`
	Preemptiondelaytimer int     `json:"preemptiondelaytimer,omitempty"`
	Priority             int     `json:"priority,omitempty"`
	Sharing              string  `json:"sharing,omitempty"`
	State                int     `json:"state,omitempty"`
	Trackifnumpriority   int     `json:"trackifnumpriority,omitempty"`
	Tracking             string  `json:"tracking,omitempty"`
	TypeField            string  `json:"type,omitempty"`
}

type NetprofileNatruleBinding struct {
	Name      string `json:"name,omitempty"`
	Natrule   string `json:"natrule,omitempty"`
	Netmask   string `json:"netmask,omitempty"`
	Rewriteip string `json:"rewriteip,omitempty"`
}

type Iptunnel struct {
	Channel            int      `json:"channel,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Destport           int      `json:"destport,omitempty"`
	Encapip            string   `json:"encapip,omitempty"`
	Grepayload         string   `json:"grepayload,omitempty"`
	Ipsecprofilename   string   `json:"ipsecprofilename,omitempty"`
	Ipsectunnelstatus  string   `json:"ipsectunnelstatus,omitempty"`
	Local              string   `json:"local,omitempty"`
	Name               string   `json:"name,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Ownergroup         string   `json:"ownergroup,omitempty"`
	Protocol           string   `json:"protocol,omitempty"`
	Refcnt             int      `json:"refcnt,omitempty"`
	Remote             string   `json:"remote,omitempty"`
	Remotesubnetmask   string   `json:"remotesubnetmask,omitempty"`
	Sysname            string   `json:"sysname,omitempty"`
	Tosinherit         string   `json:"tosinherit,omitempty"`
	Tunneltype         []string `json:"tunneltype,omitempty"`
	TypeField          int      `json:"type,omitempty"`
	Vlan               int      `json:"vlan,omitempty"`
	Vlantagging        string   `json:"vlantagging,omitempty"`
	Vnid               int      `json:"vnid,omitempty"`
}

type NetprofileSrcportsetBinding struct {
	Name         string `json:"name,omitempty"`
	Srcportrange string `json:"srcportrange,omitempty"`
}

type L2param struct {
	Bdggrpproxyarp          string `json:"bdggrpproxyarp,omitempty"`
	Bdgsetting              string `json:"bdgsetting,omitempty"`
	Bridgeagetimeout        int    `json:"bridgeagetimeout,omitempty"`
	Garponvridintf          string `json:"garponvridintf,omitempty"`
	Garpreply               string `json:"garpreply,omitempty"`
	Macmodefwdmypkt         string `json:"macmodefwdmypkt,omitempty"`
	Maxbridgecollision      int    `json:"maxbridgecollision,omitempty"`
	Mbfinstlearning         string `json:"mbfinstlearning,omitempty"`
	Mbfpeermacupdate        int    `json:"mbfpeermacupdate,omitempty"`
	Nextgenapiresource      string `json:"_nextgenapiresource,omitempty"`
	Proxyarp                string `json:"proxyarp,omitempty"`
	Returntoethernetsender  string `json:"returntoethernetsender,omitempty"`
	Rstintfonhafo           string `json:"rstintfonhafo,omitempty"`
	Skipproxyingbsdtraffic  string `json:"skipproxyingbsdtraffic,omitempty"`
	Stopmacmoveupdate       string `json:"stopmacmoveupdate,omitempty"`
	Usemymac                string `json:"usemymac,omitempty"`
	Usenetprofilebsdtraffic string `json:"usenetprofilebsdtraffic,omitempty"`
}

type Nd6ravariablesBinding struct {
	Nd6ravariablesOnlinkipv6prefixBinding []interface{} `json:"nd6ravariables_onlinkipv6prefix_binding,omitempty"`
	Vlan                                  int           `json:"vlan,omitempty"`
}

type Bridgegroup struct {
	Count              float64 `json:"__count,omitempty"`
	Dynamicrouting     string  `json:"dynamicrouting,omitempty"`
	Flags              bool    `json:"flags,omitempty"`
	Id                 int     `json:"id,omitempty"`
	Ifaces             string  `json:"ifaces,omitempty"`
	Ipv6dynamicrouting string  `json:"ipv6dynamicrouting,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Partitionname      string  `json:"partitionname,omitempty"`
	Portbitmap         int     `json:"portbitmap,omitempty"`
	Rnat               bool    `json:"rnat,omitempty"`
	Tagbitmap          int     `json:"tagbitmap,omitempty"`
	Tagifaces          string  `json:"tagifaces,omitempty"`
}

type IpsetNsip6Binding struct {
	Ipaddress string `json:"ipaddress,omitempty"`
	Name      string `json:"name,omitempty"`
}

type Vrid6ChannelBinding struct {
	Flags int    `json:"flags,omitempty"`
	Id    int    `json:"id,omitempty"`
	Ifnum string `json:"ifnum,omitempty"`
	Vlan  int    `json:"vlan,omitempty"`
}

type ChannelBinding struct {
	ChannelInterfaceBinding []interface{} `json:"channel_interface_binding,omitempty"`
	Id                      string        `json:"id,omitempty"`
}

type Rnat6 struct {
	Acl6name           string  `json:"acl6name,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Network            string  `json:"network,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Ownergroup         string  `json:"ownergroup,omitempty"`
	Redirectport       int     `json:"redirectport,omitempty"`
	Srcippersistency   string  `json:"srcippersistency,omitempty"`
	Td                 int     `json:"td,omitempty"`
}

type BridgegroupNsip6Binding struct {
	Id         int    `json:"id,omitempty"`
	Ipaddress  string `json:"ipaddress,omitempty"`
	Netmask    string `json:"netmask,omitempty"`
	Ownergroup string `json:"ownergroup,omitempty"`
	Rnat       bool   `json:"rnat,omitempty"`
	Td         int    `json:"td,omitempty"`
}

type BridgegroupVlanBinding struct {
	Id   int  `json:"id,omitempty"`
	Rnat bool `json:"rnat,omitempty"`
	Vlan int  `json:"vlan,omitempty"`
}

type Nd6ravariables struct {
	Ceaserouteradv           string  `json:"ceaserouteradv,omitempty"`
	Count                    float64 `json:"__count,omitempty"`
	Currhoplimit             int     `json:"currhoplimit,omitempty"`
	Defaultlifetime          int     `json:"defaultlifetime,omitempty"`
	Lastrtadvtime            int     `json:"lastrtadvtime,omitempty"`
	Linkmtu                  int     `json:"linkmtu,omitempty"`
	Managedaddrconfig        string  `json:"managedaddrconfig,omitempty"`
	Maxrtadvinterval         int     `json:"maxrtadvinterval,omitempty"`
	Minrtadvinterval         int     `json:"minrtadvinterval,omitempty"`
	Nextgenapiresource       string  `json:"_nextgenapiresource,omitempty"`
	Nextrtadvdelay           int     `json:"nextrtadvdelay,omitempty"`
	Onlyunicastrtadvresponse string  `json:"onlyunicastrtadvresponse,omitempty"`
	Otheraddrconfig          string  `json:"otheraddrconfig,omitempty"`
	Reachabletime            int     `json:"reachabletime,omitempty"`
	Retranstime              int     `json:"retranstime,omitempty"`
	Sendrouteradv            string  `json:"sendrouteradv,omitempty"`
	Srclinklayeraddroption   string  `json:"srclinklayeraddroption,omitempty"`
	Vlan                     int     `json:"vlan,omitempty"`
}

type VridNsipBinding struct {
	Flags     int    `json:"flags,omitempty"`
	Id        int    `json:"id,omitempty"`
	Ipaddress string `json:"ipaddress,omitempty"`
}

type VlanInterfaceBinding struct {
	Id         int    `json:"id,omitempty"`
	Ifnum      string `json:"ifnum,omitempty"`
	Ownergroup string `json:"ownergroup,omitempty"`
	Tagged     bool   `json:"tagged,omitempty"`
}

type Ip6tunnelparam struct {
	Dropfrag             string `json:"dropfrag,omitempty"`
	Dropfragcputhreshold int    `json:"dropfragcputhreshold,omitempty"`
	Nextgenapiresource   string `json:"_nextgenapiresource,omitempty"`
	Srcip                string `json:"srcip,omitempty"`
	Srciproundrobin      string `json:"srciproundrobin,omitempty"`
	Useclientsourceipv6  string `json:"useclientsourceipv6,omitempty"`
}

type Vrid6TrackinterfaceBinding struct {
	Flags      int    `json:"flags,omitempty"`
	Id         int    `json:"id,omitempty"`
	Trackifnum string `json:"trackifnum,omitempty"`
}

type RnatglobalAuditsyslogpolicyBinding struct {
	All      bool   `json:"all,omitempty"`
	Policy   string `json:"policy,omitempty"`
	Priority int    `json:"priority,omitempty"`
}

type Arpparam struct {
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Spoofvalidation    string `json:"spoofvalidation,omitempty"`
	Timeout            int    `json:"timeout,omitempty"`
}

type Channel struct {
	Actflowctl                string   `json:"actflowctl,omitempty"`
	Actspeed                  string   `json:"actspeed,omitempty"`
	Actthroughput             int      `json:"actthroughput,omitempty"`
	Actualmtu                 int      `json:"actualmtu,omitempty"`
	Autoneg                   int      `json:"autoneg,omitempty"`
	Autonegresult             int      `json:"autonegresult,omitempty"`
	Backplane                 string   `json:"backplane,omitempty"`
	Bandwidthhigh             int      `json:"bandwidthhigh,omitempty"`
	Bandwidthnormal           int      `json:"bandwidthnormal,omitempty"`
	Bdgmuted                  int      `json:"bdgmuted,omitempty"`
	Cleartime                 int      `json:"cleartime,omitempty"`
	Conndistr                 string   `json:"conndistr,omitempty"`
	Count                     float64  `json:"__count,omitempty"`
	Description               string   `json:"description,omitempty"`
	Devicename                string   `json:"devicename,omitempty"`
	Downtime                  int      `json:"downtime,omitempty"`
	Duplex                    string   `json:"duplex,omitempty"`
	Fctls                     int      `json:"fctls,omitempty"`
	Flags                     int      `json:"flags,omitempty"`
	Flowctl                   string   `json:"flowctl,omitempty"`
	Haheartbeat               string   `json:"haheartbeat,omitempty"`
	Hamonitor                 string   `json:"hamonitor,omitempty"`
	Hangdetect                int      `json:"hangdetect,omitempty"`
	Hangreset                 int      `json:"hangreset,omitempty"`
	Hangs                     int      `json:"hangs,omitempty"`
	Id                        string   `json:"id,omitempty"`
	Ifalias                   string   `json:"ifalias,omitempty"`
	Ifnum                     []string `json:"ifnum,omitempty"`
	Indisc                    int      `json:"indisc,omitempty"`
	Intfstate                 int      `json:"intfstate,omitempty"`
	Lacpactoraggregation      string   `json:"lacpactoraggregation,omitempty"`
	Lacpactorcollecting       string   `json:"lacpactorcollecting,omitempty"`
	Lacpactordistributing     string   `json:"lacpactordistributing,omitempty"`
	Lacpactorinsync           string   `json:"lacpactorinsync,omitempty"`
	Lacpactorportno           int      `json:"lacpactorportno,omitempty"`
	Lacpactorpriority         int      `json:"lacpactorpriority,omitempty"`
	Lacpmode                  string   `json:"lacpmode,omitempty"`
	Lacppartneraggregation    string   `json:"lacppartneraggregation,omitempty"`
	Lacppartnercollecting     string   `json:"lacppartnercollecting,omitempty"`
	Lacppartnerdefaulted      string   `json:"lacppartnerdefaulted,omitempty"`
	Lacppartnerdistributing   string   `json:"lacppartnerdistributing,omitempty"`
	Lacppartnerexpired        string   `json:"lacppartnerexpired,omitempty"`
	Lacppartnerinsync         string   `json:"lacppartnerinsync,omitempty"`
	Lacppartnerkey            int      `json:"lacppartnerkey,omitempty"`
	Lacppartnerportno         int      `json:"lacppartnerportno,omitempty"`
	Lacppartnerpriority       int      `json:"lacppartnerpriority,omitempty"`
	Lacppartnerstate          string   `json:"lacppartnerstate,omitempty"`
	Lacppartnersystemmac      string   `json:"lacppartnersystemmac,omitempty"`
	Lacppartnersystempriority int      `json:"lacppartnersystempriority,omitempty"`
	Lacppartnertimeout        string   `json:"lacppartnertimeout,omitempty"`
	Lacpportmuxstate          string   `json:"lacpportmuxstate,omitempty"`
	Lacpportrxstat            string   `json:"lacpportrxstat,omitempty"`
	Lacpportselectstate       string   `json:"lacpportselectstate,omitempty"`
	Lacptimeout               string   `json:"lacptimeout,omitempty"`
	Lamac                     string   `json:"lamac,omitempty"`
	Lamode                    string   `json:"lamode,omitempty"`
	Linkredundancy            string   `json:"linkredundancy,omitempty"`
	Linkstate                 int      `json:"linkstate,omitempty"`
	Lldpmode                  string   `json:"lldpmode,omitempty"`
	Lrminthroughput           int      `json:"lrminthroughput,omitempty"`
	Mac                       string   `json:"mac,omitempty"`
	Macdistr                  string   `json:"macdistr,omitempty"`
	Media                     string   `json:"media,omitempty"`
	Mode                      string   `json:"mode,omitempty"`
	Mtu                       int      `json:"mtu,omitempty"`
	Nextgenapiresource        string   `json:"_nextgenapiresource,omitempty"`
	Outdisc                   int      `json:"outdisc,omitempty"`
	Reqduplex                 string   `json:"reqduplex,omitempty"`
	Reqflowcontrol            string   `json:"reqflowcontrol,omitempty"`
	Reqmedia                  string   `json:"reqmedia,omitempty"`
	Reqspeed                  string   `json:"reqspeed,omitempty"`
	Reqthroughput             int      `json:"reqthroughput,omitempty"`
	Rxbytes                   int      `json:"rxbytes,omitempty"`
	Rxdrops                   int      `json:"rxdrops,omitempty"`
	Rxerrors                  int      `json:"rxerrors,omitempty"`
	Rxpackets                 int      `json:"rxpackets,omitempty"`
	Rxstalls                  int      `json:"rxstalls,omitempty"`
	Speed                     string   `json:"speed,omitempty"`
	State                     string   `json:"state,omitempty"`
	Stsstalls                 int      `json:"stsstalls,omitempty"`
	Tagall                    string   `json:"tagall,omitempty"`
	Tagged                    int      `json:"tagged,omitempty"`
	Taggedany                 int      `json:"taggedany,omitempty"`
	Taggedautolearn           int      `json:"taggedautolearn,omitempty"`
	Throughput                int      `json:"throughput,omitempty"`
	Trunk                     string   `json:"trunk,omitempty"`
	Txbytes                   int      `json:"txbytes,omitempty"`
	Txdrops                   int      `json:"txdrops,omitempty"`
	Txerrors                  int      `json:"txerrors,omitempty"`
	Txpackets                 int      `json:"txpackets,omitempty"`
	Txstalls                  int      `json:"txstalls,omitempty"`
	Unit                      int      `json:"unit,omitempty"`
	Uptime                    int      `json:"uptime,omitempty"`
	Vlan                      int      `json:"vlan,omitempty"`
	Vmac                      string   `json:"vmac,omitempty"`
	Vmac6                     string   `json:"vmac6,omitempty"`
}

type Netprofile struct {
	Badipactionthreshold           int      `json:"badipactionthreshold,omitempty"`
	Count                          float64  `json:"__count,omitempty"`
	Mbf                            string   `json:"mbf,omitempty"`
	Name                           string   `json:"name,omitempty"`
	Nextgenapiresource             string   `json:"_nextgenapiresource,omitempty"`
	Overridelsn                    string   `json:"overridelsn,omitempty"`
	Proxyprotocol                  string   `json:"proxyprotocol,omitempty"`
	Proxyprotocolaftertlshandshake string   `json:"proxyprotocolaftertlshandshake,omitempty"`
	Proxyprotocoltlvoptions        []string `json:"proxyprotocoltlvoptions,omitempty"`
	Proxyprotocoltxversion         string   `json:"proxyprotocoltxversion,omitempty"`
	Srcip                          string   `json:"srcip,omitempty"`
	Srcippersistency               string   `json:"srcippersistency,omitempty"`
	Td                             int      `json:"td,omitempty"`
}

type VxlanNsipBinding struct {
	Id        int    `json:"id,omitempty"`
	Ipaddress string `json:"ipaddress,omitempty"`
	Netmask   string `json:"netmask,omitempty"`
}

type Rnat struct {
	Aclname            string  `json:"aclname,omitempty"`
	Connfailover       string  `json:"connfailover,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Natip              string  `json:"natip,omitempty"`
	Netmask            string  `json:"netmask,omitempty"`
	Network            string  `json:"network,omitempty"`
	Newname            string  `json:"newname,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Ownergroup         string  `json:"ownergroup,omitempty"`
	Redirectport       int     `json:"redirectport,omitempty"`
	Srcippersistency   string  `json:"srcippersistency,omitempty"`
	Td                 int     `json:"td,omitempty"`
	Useproxyport       string  `json:"useproxyport,omitempty"`
}

type Appalgparam struct {
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Pptpgreidletimeout int    `json:"pptpgreidletimeout,omitempty"`
}

type Iptunnelparam struct {
	Dropfrag             string `json:"dropfrag,omitempty"`
	Dropfragcputhreshold int    `json:"dropfragcputhreshold,omitempty"`
	Enablestrictrx       string `json:"enablestrictrx,omitempty"`
	Enablestricttx       string `json:"enablestricttx,omitempty"`
	Mac                  string `json:"mac,omitempty"`
	Nextgenapiresource   string `json:"_nextgenapiresource,omitempty"`
	Srcip                string `json:"srcip,omitempty"`
	Srciproundrobin      string `json:"srciproundrobin,omitempty"`
	Useclientsourceip    string `json:"useclientsourceip,omitempty"`
}

type VridNsip6Binding struct {
	Flags     int    `json:"flags,omitempty"`
	Id        int    `json:"id,omitempty"`
	Ipaddress string `json:"ipaddress,omitempty"`
}

type BridgegroupBinding struct {
	BridgegroupNsip6Binding []interface{} `json:"bridgegroup_nsip6_binding,omitempty"`
	BridgegroupNsipBinding  []interface{} `json:"bridgegroup_nsip_binding,omitempty"`
	BridgegroupVlanBinding  []interface{} `json:"bridgegroup_vlan_binding,omitempty"`
	Id                      int           `json:"id,omitempty"`
}

type Ipset struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Td                 int     `json:"td,omitempty"`
}

type Arp struct {
	All                bool    `json:"all,omitempty"`
	Channel            int     `json:"channel,omitempty"`
	Controlplane       bool    `json:"controlplane,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Flags              int     `json:"flags,omitempty"`
	Ifnum              string  `json:"ifnum,omitempty"`
	Ipaddress          string  `json:"ipaddress,omitempty"`
	Mac                string  `json:"mac,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Nodeid             int     `json:"nodeid,omitempty"`
	Ownernode          int     `json:"ownernode,omitempty"`
	State              int     `json:"state,omitempty"`
	Td                 int     `json:"td,omitempty"`
	Timeout            int     `json:"timeout,omitempty"`
	TypeField          string  `json:"type,omitempty"`
	Vlan               int     `json:"vlan,omitempty"`
	Vtep               string  `json:"vtep,omitempty"`
	Vxlan              int     `json:"vxlan,omitempty"`
}

type Inatparam struct {
	Count              float64 `json:"__count,omitempty"`
	Nat46fragheader    string  `json:"nat46fragheader,omitempty"`
	Nat46ignoretos     string  `json:"nat46ignoretos,omitempty"`
	Nat46v6mtu         int     `json:"nat46v6mtu,omitempty"`
	Nat46v6prefix      string  `json:"nat46v6prefix,omitempty"`
	Nat46zerochecksum  string  `json:"nat46zerochecksum,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Td                 int     `json:"td,omitempty"`
}

type L3param struct {
	Acllogtime           int    `json:"acllogtime,omitempty"`
	Allowclasseipv4      string `json:"allowclasseipv4,omitempty"`
	Dropdfflag           string `json:"dropdfflag,omitempty"`
	Dropipfragments      string `json:"dropipfragments,omitempty"`
	Dynamicrouting       string `json:"dynamicrouting,omitempty"`
	Externalloopback     string `json:"externalloopback,omitempty"`
	Forwardicmpfragments string `json:"forwardicmpfragments,omitempty"`
	Icmpgenratethreshold int    `json:"icmpgenratethreshold,omitempty"`
	Implicitaclallow     string `json:"implicitaclallow,omitempty"`
	Implicitpbr          string `json:"implicitpbr,omitempty"`
	Ipv6dynamicrouting   string `json:"ipv6dynamicrouting,omitempty"`
	Miproundrobin        string `json:"miproundrobin,omitempty"`
	Nextgenapiresource   string `json:"_nextgenapiresource,omitempty"`
	Overridernat         string `json:"overridernat,omitempty"`
	Srcnat               string `json:"srcnat,omitempty"`
	Tnlpmtuwoconn        string `json:"tnlpmtuwoconn,omitempty"`
	Usipserverstraypkt   string `json:"usipserverstraypkt,omitempty"`
}

type Nat64param struct {
	Count              float64 `json:"__count,omitempty"`
	Nat64fragheader    string  `json:"nat64fragheader,omitempty"`
	Nat64ignoretos     string  `json:"nat64ignoretos,omitempty"`
	Nat64v6mtu         int     `json:"nat64v6mtu,omitempty"`
	Nat64zerochecksum  string  `json:"nat64zerochecksum,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Td                 int     `json:"td,omitempty"`
}

type Lacp struct {
	Clustermac         string  `json:"clustermac,omitempty"`
	Clustersyspriority int     `json:"clustersyspriority,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Devicename         string  `json:"devicename,omitempty"`
	Flags              int     `json:"flags,omitempty"`
	Lacpkey            int     `json:"lacpkey,omitempty"`
	Mac                string  `json:"mac,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Ownernode          int     `json:"ownernode,omitempty"`
	Syspriority        int     `json:"syspriority,omitempty"`
}

type Mapdomain struct {
	Count              float64 `json:"__count,omitempty"`
	Mapdmrname         string  `json:"mapdmrname,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
}

type Forwardingsession struct {
	Acl6name           string  `json:"acl6name,omitempty"`
	Aclname            string  `json:"aclname,omitempty"`
	Connfailover       string  `json:"connfailover,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Netmask            string  `json:"netmask,omitempty"`
	Network            string  `json:"network,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Processlocal       string  `json:"processlocal,omitempty"`
	Sourceroutecache   string  `json:"sourceroutecache,omitempty"`
	Td                 int     `json:"td,omitempty"`
}

type Route6 struct {
	Active             bool     `json:"active,omitempty"`
	Advertise          string   `json:"advertise,omitempty"`
	Bgp                bool     `json:"bgp,omitempty"`
	Connected          bool     `json:"connected,omitempty"`
	Cost               int      `json:"cost,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Data               bool     `json:"data,omitempty"`
	Data1              string   `json:"data1,omitempty"`
	Detail             bool     `json:"detail,omitempty"`
	Distance           int      `json:"distance,omitempty"`
	Dynamic            bool     `json:"dynamic,omitempty"`
	Failedprobes       int      `json:"failedprobes,omitempty"`
	Flags              bool     `json:"flags,omitempty"`
	Gateway            string   `json:"gateway,omitempty"`
	Gatewayname        string   `json:"gatewayname,omitempty"`
	Isis               bool     `json:"isis,omitempty"`
	Mgmt               bool     `json:"mgmt,omitempty"`
	Monitor            string   `json:"monitor,omitempty"`
	Monstatcode        int      `json:"monstatcode,omitempty"`
	Monstatparam1      int      `json:"monstatparam1,omitempty"`
	Monstatparam2      int      `json:"monstatparam2,omitempty"`
	Monstatparam3      int      `json:"monstatparam3,omitempty"`
	Msr                string   `json:"msr,omitempty"`
	Network            string   `json:"network,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Ospfv3             bool     `json:"ospfv3,omitempty"`
	Ownergroup         string   `json:"ownergroup,omitempty"`
	Permanent          bool     `json:"permanent,omitempty"`
	Raroute            bool     `json:"raroute,omitempty"`
	Retain             int      `json:"retain,omitempty"`
	Rip                bool     `json:"rip,omitempty"`
	Routeowners        []string `json:"routeowners,omitempty"`
	Routetype          string   `json:"routetype,omitempty"`
	State              int      `json:"state,omitempty"`
	Static             bool     `json:"Static,omitempty"`
	Td                 int      `json:"td,omitempty"`
	Totalfailedprobes  int      `json:"totalfailedprobes,omitempty"`
	Totalprobes        int      `json:"totalprobes,omitempty"`
	TypeField          bool     `json:"type,omitempty"`
	Vlan               int      `json:"vlan,omitempty"`
	Vxlan              int      `json:"vxlan,omitempty"`
	Weight             int      `json:"weight,omitempty"`
}

type VlanLinksetBinding struct {
	Id         int    `json:"id,omitempty"`
	Ifnum      string `json:"ifnum,omitempty"`
	Ownergroup string `json:"ownergroup,omitempty"`
	Tagged     bool   `json:"tagged,omitempty"`
}

type VlanChannelBinding struct {
	Id         int    `json:"id,omitempty"`
	Ifnum      string `json:"ifnum,omitempty"`
	Ownergroup string `json:"ownergroup,omitempty"`
	Tagged     bool   `json:"tagged,omitempty"`
}

type FisBinding struct {
	FisChannelBinding []interface{} `json:"fis_channel_binding,omitempty"`
	Name              string        `json:"name,omitempty"`
}

type VxlanIptunnelBinding struct {
	Id     int    `json:"id,omitempty"`
	Tunnel string `json:"tunnel,omitempty"`
}

type NetprofileBinding struct {
	Name                        string        `json:"name,omitempty"`
	NetprofileNatruleBinding    []interface{} `json:"netprofile_natrule_binding,omitempty"`
	NetprofileSrcportsetBinding []interface{} `json:"netprofile_srcportset_binding,omitempty"`
}

type VxlanvlanmapVxlanBinding struct {
	Name  string   `json:"name,omitempty"`
	Vlan  []string `json:"vlan,omitempty"`
	Vxlan int      `json:"vxlan,omitempty"`
}

type Netbridge struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Vxlanvlanmap       string  `json:"vxlanvlanmap,omitempty"`
}

type Vridparam struct {
	Deadinterval       int    `json:"deadinterval,omitempty"`
	Hellointerval      int    `json:"hellointerval,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Sendtomaster       string `json:"sendtomaster,omitempty"`
}

type VxlanSrcipBinding struct {
	Id    int    `json:"id,omitempty"`
	Srcip string `json:"srcip,omitempty"`
}

type IpsetBinding struct {
	IpsetNsip6Binding []interface{} `json:"ipset_nsip6_binding,omitempty"`
	IpsetNsipBinding  []interface{} `json:"ipset_nsip_binding,omitempty"`
	Name              string        `json:"name,omitempty"`
}

type Nat64 struct {
	Acl6name           string  `json:"acl6name,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Netprofile         string  `json:"netprofile,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
}

type Rnatsession struct {
	Aclname string `json:"aclname,omitempty"`
	Natip   string `json:"natip,omitempty"`
	Netmask string `json:"netmask,omitempty"`
	Network string `json:"network,omitempty"`
}

type Vxlan struct {
	Count              float64 `json:"__count,omitempty"`
	Dynamicrouting     string  `json:"dynamicrouting,omitempty"`
	Id                 int     `json:"id,omitempty"`
	Innervlantagging   string  `json:"innervlantagging,omitempty"`
	Ipv6dynamicrouting string  `json:"ipv6dynamicrouting,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Partitionname      string  `json:"partitionname,omitempty"`
	Port               int     `json:"port,omitempty"`
	Protocol           string  `json:"protocol,omitempty"`
	Td                 int     `json:"td,omitempty"`
	TypeField          string  `json:"type,omitempty"`
	Vlan               int     `json:"vlan,omitempty"`
}

type FisInterfaceBinding struct {
	Ifnum     string `json:"ifnum,omitempty"`
	Name      string `json:"name,omitempty"`
	Ownernode int    `json:"ownernode,omitempty"`
}

type VridChannelBinding struct {
	Flags int    `json:"flags,omitempty"`
	Id    int    `json:"id,omitempty"`
	Ifnum string `json:"ifnum,omitempty"`
	Vlan  int    `json:"vlan,omitempty"`
}

type VridBinding struct {
	Id                        int           `json:"id,omitempty"`
	VridChannelBinding        []interface{} `json:"vrid_channel_binding,omitempty"`
	VridInterfaceBinding      []interface{} `json:"vrid_interface_binding,omitempty"`
	VridNsip6Binding          []interface{} `json:"vrid_nsip6_binding,omitempty"`
	VridNsipBinding           []interface{} `json:"vrid_nsip_binding,omitempty"`
	VridTrackinterfaceBinding []interface{} `json:"vrid_trackinterface_binding,omitempty"`
}

type NetbridgeIptunnelBinding struct {
	Name   string `json:"name,omitempty"`
	Tunnel string `json:"tunnel,omitempty"`
}

type IpsetNsipBinding struct {
	Ipaddress string `json:"ipaddress,omitempty"`
	Name      string `json:"name,omitempty"`
}

type VlanNsipBinding struct {
	Id         int    `json:"id,omitempty"`
	Ipaddress  string `json:"ipaddress,omitempty"`
	Netmask    string `json:"netmask,omitempty"`
	Ownergroup string `json:"ownergroup,omitempty"`
	Td         int    `json:"td,omitempty"`
}

type LinksetInterfaceBinding struct {
	Id    string `json:"id,omitempty"`
	Ifnum string `json:"ifnum,omitempty"`
}

type VridTrackinterfaceBinding struct {
	Flags      int    `json:"flags,omitempty"`
	Id         int    `json:"id,omitempty"`
	Trackifnum string `json:"trackifnum,omitempty"`
}

type Vrid6Binding struct {
	Id                         int           `json:"id,omitempty"`
	Vrid6ChannelBinding        []interface{} `json:"vrid6_channel_binding,omitempty"`
	Vrid6InterfaceBinding      []interface{} `json:"vrid6_interface_binding,omitempty"`
	Vrid6Nsip6Binding          []interface{} `json:"vrid6_nsip6_binding,omitempty"`
	Vrid6NsipBinding           []interface{} `json:"vrid6_nsip_binding,omitempty"`
	Vrid6TrackinterfaceBinding []interface{} `json:"vrid6_trackinterface_binding,omitempty"`
}

type NetbridgeNsipBinding struct {
	Ipaddress string `json:"ipaddress,omitempty"`
	Name      string `json:"name,omitempty"`
	Netmask   string `json:"netmask,omitempty"`
}

type Rsskeytype struct {
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Rsstype            string `json:"rsstype,omitempty"`
}

type MapbmrBmrv4networkBinding struct {
	Name    string `json:"name,omitempty"`
	Netmask string `json:"netmask,omitempty"`
	Network string `json:"network,omitempty"`
}

type Rnat6Nsip6Binding struct {
	Name       string `json:"name,omitempty"`
	Natip6     string `json:"natip6,omitempty"`
	Ownergroup string `json:"ownergroup,omitempty"`
	Td         int    `json:"td,omitempty"`
}

type Route struct {
	Adv                bool     `json:"adv,omitempty"`
	Advbgp             bool     `json:"advbgp,omitempty"`
	Advertise          string   `json:"advertise,omitempty"`
	Advisis            bool     `json:"advisis,omitempty"`
	Advospf            bool     `json:"advospf,omitempty"`
	Advrip             bool     `json:"advrip,omitempty"`
	Bgp                bool     `json:"bgp,omitempty"`
	Cost               int      `json:"cost,omitempty"`
	Cost1              int      `json:"cost1,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Data               bool     `json:"data,omitempty"`
	Data0              bool     `json:"data0,omitempty"`
	Detail             bool     `json:"detail,omitempty"`
	Dhcp               bool     `json:"dhcp,omitempty"`
	Direct             bool     `json:"direct,omitempty"`
	Distance           int      `json:"distance,omitempty"`
	Dynamic            bool     `json:"dynamic,omitempty"`
	Failedprobes       int      `json:"failedprobes,omitempty"`
	Flags              bool     `json:"flags,omitempty"`
	Gateway            string   `json:"gateway,omitempty"`
	Gatewayname        string   `json:"gatewayname,omitempty"`
	Isis               bool     `json:"isis,omitempty"`
	Lbroute            bool     `json:"lbroute,omitempty"`
	Mgmt               bool     `json:"mgmt,omitempty"`
	Monitor            string   `json:"monitor,omitempty"`
	Monstatcode        int      `json:"monstatcode,omitempty"`
	Monstatparam1      int      `json:"monstatparam1,omitempty"`
	Monstatparam2      int      `json:"monstatparam2,omitempty"`
	Monstatparam3      int      `json:"monstatparam3,omitempty"`
	Msr                string   `json:"msr,omitempty"`
	Nat                bool     `json:"nat,omitempty"`
	Netmask            string   `json:"netmask,omitempty"`
	Network            string   `json:"network,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Ospf               bool     `json:"ospf,omitempty"`
	Ownergroup         string   `json:"ownergroup,omitempty"`
	Permanent          bool     `json:"permanent,omitempty"`
	Protocol           []string `json:"protocol,omitempty"`
	Retain             int      `json:"retain,omitempty"`
	Rip                bool     `json:"rip,omitempty"`
	Routeowners        []string `json:"routeowners,omitempty"`
	Routetype          string   `json:"routetype,omitempty"`
	State              int      `json:"state,omitempty"`
	Static             bool     `json:"Static,omitempty"`
	Td                 int      `json:"td,omitempty"`
	Totalfailedprobes  int      `json:"totalfailedprobes,omitempty"`
	Totalprobes        int      `json:"totalprobes,omitempty"`
	Tunnel             bool     `json:"tunnel,omitempty"`
	TypeField          bool     `json:"type,omitempty"`
	Vlan               int      `json:"vlan,omitempty"`
	Weight             int      `json:"weight,omitempty"`
}

type NetbridgeBinding struct {
	Name                     string        `json:"name,omitempty"`
	NetbridgeIptunnelBinding []interface{} `json:"netbridge_iptunnel_binding,omitempty"`
	NetbridgeNsip6Binding    []interface{} `json:"netbridge_nsip6_binding,omitempty"`
	NetbridgeNsipBinding     []interface{} `json:"netbridge_nsip_binding,omitempty"`
	NetbridgeVlanBinding     []interface{} `json:"netbridge_vlan_binding,omitempty"`
}

type LinksetChannelBinding struct {
	Id    string `json:"id,omitempty"`
	Ifnum string `json:"ifnum,omitempty"`
}

type Ip6tunnel struct {
	Count              float64 `json:"__count,omitempty"`
	Encapip            string  `json:"encapip,omitempty"`
	Local              string  `json:"local,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Ownergroup         string  `json:"ownergroup,omitempty"`
	Remote             string  `json:"remote,omitempty"`
	Remoteip           string  `json:"remoteip,omitempty"`
	TypeField          int     `json:"type,omitempty"`
}

type L4param struct {
	L2connmethod       string `json:"l2connmethod,omitempty"`
	L4switch           string `json:"l4switch,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Vxlanvlanmap struct {
	Count              float64 `json:"__count,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
}

type MapbmrBinding struct {
	MapbmrBmrv4networkBinding []interface{} `json:"mapbmr_bmrv4network_binding,omitempty"`
	Name                      string        `json:"name,omitempty"`
}

type VridInterfaceBinding struct {
	Flags int    `json:"flags,omitempty"`
	Id    int    `json:"id,omitempty"`
	Ifnum string `json:"ifnum,omitempty"`
	Vlan  int    `json:"vlan,omitempty"`
}

type Bridgetable struct {
	Bridgeage          int     `json:"bridgeage,omitempty"`
	Channel            int     `json:"channel,omitempty"`
	Controlplane       bool    `json:"controlplane,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Devicevlan         int     `json:"devicevlan,omitempty"`
	Flags              int     `json:"flags,omitempty"`
	Ifnum              string  `json:"ifnum,omitempty"`
	Mac                string  `json:"mac,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Nodeid             int     `json:"nodeid,omitempty"`
	TypeField          string  `json:"type,omitempty"`
	Vlan               int     `json:"vlan,omitempty"`
	Vni                int     `json:"vni,omitempty"`
	Vtep               string  `json:"vtep,omitempty"`
	Vxlan              int     `json:"vxlan,omitempty"`
}

type VxlanvlanmapBinding struct {
	Name                     string        `json:"name,omitempty"`
	VxlanvlanmapVxlanBinding []interface{} `json:"vxlanvlanmap_vxlan_binding,omitempty"`
}

type Vrid struct {
	All                  bool    `json:"all,omitempty"`
	Count                float64 `json:"__count,omitempty"`
	Effectivepriority    int     `json:"effectivepriority,omitempty"`
	Flags                int     `json:"flags,omitempty"`
	Id                   int     `json:"id,omitempty"`
	Ifaces               string  `json:"ifaces,omitempty"`
	Ipaddress            string  `json:"ipaddress,omitempty"`
	Nextgenapiresource   string  `json:"_nextgenapiresource,omitempty"`
	Operationalownernode int     `json:"operationalownernode,omitempty"`
	Ownernode            int     `json:"ownernode,omitempty"`
	Preemption           string  `json:"preemption,omitempty"`
	Preemptiondelaytimer int     `json:"preemptiondelaytimer,omitempty"`
	Priority             int     `json:"priority,omitempty"`
	Sharing              string  `json:"sharing,omitempty"`
	State                int     `json:"state,omitempty"`
	Trackifnumpriority   int     `json:"trackifnumpriority,omitempty"`
	Tracking             string  `json:"tracking,omitempty"`
	TypeField            string  `json:"type,omitempty"`
}

type NetbridgeNsip6Binding struct {
	Ipaddress string `json:"ipaddress,omitempty"`
	Name      string `json:"name,omitempty"`
	Netmask   string `json:"netmask,omitempty"`
}

type Inat struct {
	Connfailover       string  `json:"connfailover,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Flags              int     `json:"flags,omitempty"`
	Ftp                string  `json:"ftp,omitempty"`
	Mode               string  `json:"mode,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Privateip          string  `json:"privateip,omitempty"`
	Proxyip            string  `json:"proxyip,omitempty"`
	Publicip           string  `json:"publicip,omitempty"`
	Tcpproxy           string  `json:"tcpproxy,omitempty"`
	Td                 int     `json:"td,omitempty"`
	Tftp               string  `json:"tftp,omitempty"`
	Useproxyport       string  `json:"useproxyport,omitempty"`
	Usip               string  `json:"usip,omitempty"`
	Usnip              string  `json:"usnip,omitempty"`
}

type RnatRetainsourceportsetBinding struct {
	Name                  string `json:"name,omitempty"`
	Retainsourceportrange string `json:"retainsourceportrange,omitempty"`
}

type Vlan struct {
	Aliasname          string  `json:"aliasname,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Dynamicrouting     string  `json:"dynamicrouting,omitempty"`
	Id                 int     `json:"id,omitempty"`
	Ifaces             string  `json:"ifaces,omitempty"`
	Ifnum              string  `json:"ifnum,omitempty"`
	Ipv6dynamicrouting string  `json:"ipv6dynamicrouting,omitempty"`
	Linklocalipv6addr  string  `json:"linklocalipv6addr,omitempty"`
	Lsbitmap           int     `json:"lsbitmap,omitempty"`
	Lstagbitmap        int     `json:"lstagbitmap,omitempty"`
	Mtu                int     `json:"mtu,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Partitionname      string  `json:"partitionname,omitempty"`
	Portbitmap         int     `json:"portbitmap,omitempty"`
	Rnat               bool    `json:"rnat,omitempty"`
	Sdxvlan            string  `json:"sdxvlan,omitempty"`
	Sharing            string  `json:"sharing,omitempty"`
	Tagbitmap          int     `json:"tagbitmap,omitempty"`
	Tagged             bool    `json:"tagged,omitempty"`
	Tagifaces          string  `json:"tagifaces,omitempty"`
	Vlantd             int     `json:"vlantd,omitempty"`
	Vxlan              int     `json:"vxlan,omitempty"`
}

type Portallocation struct {
	Count              float64 `json:"__count,omitempty"`
	Destip             string  `json:"destip,omitempty"`
	Destport           int     `json:"destport,omitempty"`
	Freeports          int     `json:"freeports,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Nodeid             int     `json:"nodeid,omitempty"`
	Protocol           int     `json:"protocol,omitempty"`
	Srcip              string  `json:"srcip,omitempty"`
}

type Rnatparam struct {
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Srcippersistency   string `json:"srcippersistency,omitempty"`
	Tcpproxy           string `json:"tcpproxy,omitempty"`
}

type Vrid6Nsip6Binding struct {
	Flags     int    `json:"flags,omitempty"`
	Id        int    `json:"id,omitempty"`
	Ipaddress string `json:"ipaddress,omitempty"`
}

type VxlanNsip6Binding struct {
	Id        int    `json:"id,omitempty"`
	Ipaddress string `json:"ipaddress,omitempty"`
	Netmask   string `json:"netmask,omitempty"`
}

type VlanBinding struct {
	Id                   int           `json:"id,omitempty"`
	VlanChannelBinding   []interface{} `json:"vlan_channel_binding,omitempty"`
	VlanInterfaceBinding []interface{} `json:"vlan_interface_binding,omitempty"`
	VlanLinksetBinding   []interface{} `json:"vlan_linkset_binding,omitempty"`
	VlanNsip6Binding     []interface{} `json:"vlan_nsip6_binding,omitempty"`
	VlanNsipBinding      []interface{} `json:"vlan_nsip_binding,omitempty"`
}

type RnatBinding struct {
	Name                           string        `json:"name,omitempty"`
	RnatNsipBinding                []interface{} `json:"rnat_nsip_binding,omitempty"`
	RnatRetainsourceportsetBinding []interface{} `json:"rnat_retainsourceportset_binding,omitempty"`
}

type NetbridgeVlanBinding struct {
	Name string `json:"name,omitempty"`
	Vlan int    `json:"vlan,omitempty"`
}

type Interface struct {
	Actduplex                 string   `json:"actduplex,omitempty"`
	Actflowctl                string   `json:"actflowctl,omitempty"`
	Actmedia                  string   `json:"actmedia,omitempty"`
	Actspeed                  string   `json:"actspeed,omitempty"`
	Actthroughput             int      `json:"actthroughput,omitempty"`
	Actualmtu                 int      `json:"actualmtu,omitempty"`
	Actualringsize            int      `json:"actualringsize,omitempty"`
	Autoneg                   string   `json:"autoneg,omitempty"`
	Autonegresult             int      `json:"autonegresult,omitempty"`
	Backplane                 string   `json:"backplane,omitempty"`
	Bandwidthhigh             int      `json:"bandwidthhigh,omitempty"`
	Bandwidthnormal           int      `json:"bandwidthnormal,omitempty"`
	Bdgmacmoved               int      `json:"bdgmacmoved,omitempty"`
	Bdgmuted                  int      `json:"bdgmuted,omitempty"`
	Cleartime                 int      `json:"cleartime,omitempty"`
	Count                     float64  `json:"__count,omitempty"`
	Description               string   `json:"description,omitempty"`
	Devicename                string   `json:"devicename,omitempty"`
	Downtime                  int      `json:"downtime,omitempty"`
	Duplex                    string   `json:"duplex,omitempty"`
	Fctls                     int      `json:"fctls,omitempty"`
	Flags                     int      `json:"flags,omitempty"`
	Flowctl                   string   `json:"flowctl,omitempty"`
	Haheartbeat               string   `json:"haheartbeat,omitempty"`
	Hamonitor                 string   `json:"hamonitor,omitempty"`
	Hangdetect                int      `json:"hangdetect,omitempty"`
	Hangreset                 int      `json:"hangreset,omitempty"`
	Hangs                     int      `json:"hangs,omitempty"`
	Id                        string   `json:"id,omitempty"`
	Ifalias                   string   `json:"ifalias,omitempty"`
	Ifnum                     []string `json:"ifnum,omitempty"`
	Indisc                    int      `json:"indisc,omitempty"`
	Intfstate                 int      `json:"intfstate,omitempty"`
	Intftype                  string   `json:"intftype,omitempty"`
	Lacpactoraggregation      string   `json:"lacpactoraggregation,omitempty"`
	Lacpactorcollecting       string   `json:"lacpactorcollecting,omitempty"`
	Lacpactordistributing     string   `json:"lacpactordistributing,omitempty"`
	Lacpactorinsync           string   `json:"lacpactorinsync,omitempty"`
	Lacpactormode             string   `json:"lacpactormode,omitempty"`
	Lacpactorportno           int      `json:"lacpactorportno,omitempty"`
	Lacpactorpriority         int      `json:"lacpactorpriority,omitempty"`
	Lacpactortimeout          string   `json:"lacpactortimeout,omitempty"`
	Lacpkey                   int      `json:"lacpkey,omitempty"`
	Lacpmode                  string   `json:"lacpmode,omitempty"`
	Lacppartneraggregation    string   `json:"lacppartneraggregation,omitempty"`
	Lacppartnercollecting     string   `json:"lacppartnercollecting,omitempty"`
	Lacppartnerdefaulted      string   `json:"lacppartnerdefaulted,omitempty"`
	Lacppartnerdistributing   string   `json:"lacppartnerdistributing,omitempty"`
	Lacppartnerexpired        string   `json:"lacppartnerexpired,omitempty"`
	Lacppartnerinsync         string   `json:"lacppartnerinsync,omitempty"`
	Lacppartnerkey            int      `json:"lacppartnerkey,omitempty"`
	Lacppartnerportno         int      `json:"lacppartnerportno,omitempty"`
	Lacppartnerpriority       int      `json:"lacppartnerpriority,omitempty"`
	Lacppartnerstate          string   `json:"lacppartnerstate,omitempty"`
	Lacppartnersystemmac      string   `json:"lacppartnersystemmac,omitempty"`
	Lacppartnersystempriority int      `json:"lacppartnersystempriority,omitempty"`
	Lacppartnertimeout        string   `json:"lacppartnertimeout,omitempty"`
	Lacpportmuxstate          string   `json:"lacpportmuxstate,omitempty"`
	Lacpportrxstat            string   `json:"lacpportrxstat,omitempty"`
	Lacpportselectstate       string   `json:"lacpportselectstate,omitempty"`
	Lacppriority              int      `json:"lacppriority,omitempty"`
	Lacptimeout               string   `json:"lacptimeout,omitempty"`
	Lagtype                   string   `json:"lagtype,omitempty"`
	Linkredundancy            string   `json:"linkredundancy,omitempty"`
	Linkstate                 int      `json:"linkstate,omitempty"`
	Lldpmode                  string   `json:"lldpmode,omitempty"`
	Lractiveintf              int      `json:"lractiveintf,omitempty"`
	Lrsetpriority             int      `json:"lrsetpriority,omitempty"`
	Mac                       string   `json:"mac,omitempty"`
	Mode                      string   `json:"mode,omitempty"`
	Mtu                       int      `json:"mtu,omitempty"`
	Nextgenapiresource        string   `json:"_nextgenapiresource,omitempty"`
	Outdisc                   int      `json:"outdisc,omitempty"`
	Reqduplex                 string   `json:"reqduplex,omitempty"`
	Reqflowcontrol            string   `json:"reqflowcontrol,omitempty"`
	Reqmedia                  string   `json:"reqmedia,omitempty"`
	Reqspeed                  string   `json:"reqspeed,omitempty"`
	Reqthroughput             int      `json:"reqthroughput,omitempty"`
	Ringsize                  int      `json:"ringsize,omitempty"`
	Ringtype                  string   `json:"ringtype,omitempty"`
	Rxbytes                   int      `json:"rxbytes,omitempty"`
	Rxdrops                   int      `json:"rxdrops,omitempty"`
	Rxerrors                  int      `json:"rxerrors,omitempty"`
	Rxpackets                 int      `json:"rxpackets,omitempty"`
	Rxstalls                  int      `json:"rxstalls,omitempty"`
	Slaveduplex               int      `json:"slaveduplex,omitempty"`
	Slaveflowctl              int      `json:"slaveflowctl,omitempty"`
	Slavemedia                int      `json:"slavemedia,omitempty"`
	Slavespeed                int      `json:"slavespeed,omitempty"`
	Slavestate                int      `json:"slavestate,omitempty"`
	Slavetime                 int      `json:"slavetime,omitempty"`
	Speed                     string   `json:"speed,omitempty"`
	State                     string   `json:"state,omitempty"`
	Stsstalls                 int      `json:"stsstalls,omitempty"`
	Svmcmd                    int      `json:"svmcmd,omitempty"`
	Tagall                    string   `json:"tagall,omitempty"`
	Tagged                    int      `json:"tagged,omitempty"`
	Taggedany                 int      `json:"taggedany,omitempty"`
	Taggedautolearn           int      `json:"taggedautolearn,omitempty"`
	Throughput                int      `json:"throughput,omitempty"`
	Trunk                     string   `json:"trunk,omitempty"`
	Trunkallowedvlan          []string `json:"trunkallowedvlan,omitempty"`
	Trunkmode                 string   `json:"trunkmode,omitempty"`
	Txbytes                   int      `json:"txbytes,omitempty"`
	Txdrops                   int      `json:"txdrops,omitempty"`
	Txerrors                  int      `json:"txerrors,omitempty"`
	Txpackets                 int      `json:"txpackets,omitempty"`
	Txstalls                  int      `json:"txstalls,omitempty"`
	Unit                      int      `json:"unit,omitempty"`
	Uptime                    int      `json:"uptime,omitempty"`
	Vlan                      int      `json:"vlan,omitempty"`
	Vmac                      string   `json:"vmac,omitempty"`
	Vmac6                     string   `json:"vmac6,omitempty"`
}

type Nd6 struct {
	Channel            int     `json:"channel,omitempty"`
	Controlplane       bool    `json:"controlplane,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Flags              int     `json:"flags,omitempty"`
	Ifnum              string  `json:"ifnum,omitempty"`
	Mac                string  `json:"mac,omitempty"`
	Neighbor           string  `json:"neighbor,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Nodeid             int     `json:"nodeid,omitempty"`
	State              string  `json:"state,omitempty"`
	Td                 int     `json:"td,omitempty"`
	Timeout            int     `json:"timeout,omitempty"`
	Vlan               int     `json:"vlan,omitempty"`
	Vtep               string  `json:"vtep,omitempty"`
	Vxlan              int     `json:"vxlan,omitempty"`
}

type Interfacepair struct {
	Count              float64  `json:"__count,omitempty"`
	Id                 int      `json:"id,omitempty"`
	Ifaces             string   `json:"ifaces,omitempty"`
	Ifnum              []string `json:"ifnum,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
}

type Mapbmr struct {
	Count              float64 `json:"__count,omitempty"`
	Eabitlength        int     `json:"eabitlength,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Psidlength         int     `json:"psidlength,omitempty"`
	Psidoffset         int     `json:"psidoffset,omitempty"`
	Ruleipv6prefix     string  `json:"ruleipv6prefix,omitempty"`
}

type VxlanBinding struct {
	Id                   int           `json:"id,omitempty"`
	VxlanIptunnelBinding []interface{} `json:"vxlan_iptunnel_binding,omitempty"`
	VxlanNsip6Binding    []interface{} `json:"vxlan_nsip6_binding,omitempty"`
	VxlanNsipBinding     []interface{} `json:"vxlan_nsip_binding,omitempty"`
	VxlanSrcipBinding    []interface{} `json:"vxlan_srcip_binding,omitempty"`
}

type Fis struct {
	Count              float64 `json:"__count,omitempty"`
	Ifaces             string  `json:"ifaces,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Ownernode          int     `json:"ownernode,omitempty"`
}

type Linkset struct {
	Count              float64 `json:"__count,omitempty"`
	Id                 string  `json:"id,omitempty"`
	Ifnum              string  `json:"ifnum,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
}

type BridgegroupNsipBinding struct {
	Id         int    `json:"id,omitempty"`
	Ipaddress  string `json:"ipaddress,omitempty"`
	Netmask    string `json:"netmask,omitempty"`
	Ownergroup string `json:"ownergroup,omitempty"`
	Rnat       bool   `json:"rnat,omitempty"`
	Td         int    `json:"td,omitempty"`
}

type Onlinkipv6prefix struct {
	Autonomusprefix          string  `json:"autonomusprefix,omitempty"`
	Count                    float64 `json:"__count,omitempty"`
	Decrementprefixlifetimes string  `json:"decrementprefixlifetimes,omitempty"`
	Depricateprefix          string  `json:"depricateprefix,omitempty"`
	Ipv6prefix               string  `json:"ipv6prefix,omitempty"`
	Nextgenapiresource       string  `json:"_nextgenapiresource,omitempty"`
	Onlinkprefix             string  `json:"onlinkprefix,omitempty"`
	Prefixcurrpreferredlft   int     `json:"prefixcurrpreferredlft,omitempty"`
	Prefixcurrvalidelft      int     `json:"prefixcurrvalidelft,omitempty"`
	Prefixpreferredlifetime  int     `json:"prefixpreferredlifetime,omitempty"`
	Prefixvalidelifetime     int     `json:"prefixvalidelifetime,omitempty"`
}

type Vrid6NsipBinding struct {
	Flags     int    `json:"flags,omitempty"`
	Id        int    `json:"id,omitempty"`
	Ipaddress string `json:"ipaddress,omitempty"`
}

type Ci struct {
	Count              float64 `json:"__count,omitempty"`
	Ifaces             string  `json:"ifaces,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
}

type Ipv6 struct {
	Basereachtime        int     `json:"basereachtime,omitempty"`
	Count                float64 `json:"__count,omitempty"`
	Dodad                string  `json:"dodad,omitempty"`
	Natprefix            string  `json:"natprefix,omitempty"`
	Ndbasereachtime      int     `json:"ndbasereachtime,omitempty"`
	Ndreachtime          int     `json:"ndreachtime,omitempty"`
	Ndretransmissiontime int     `json:"ndretransmissiontime,omitempty"`
	Nextgenapiresource   string  `json:"_nextgenapiresource,omitempty"`
	Ralearning           string  `json:"ralearning,omitempty"`
	Reachtime            int     `json:"reachtime,omitempty"`
	Retransmissiontime   int     `json:"retransmissiontime,omitempty"`
	Routerredirection    string  `json:"routerredirection,omitempty"`
	Td                   int     `json:"td,omitempty"`
	Usipnatprefix        string  `json:"usipnatprefix,omitempty"`
}

type LinksetBinding struct {
	Id                      string        `json:"id,omitempty"`
	LinksetChannelBinding   []interface{} `json:"linkset_channel_binding,omitempty"`
	LinksetInterfaceBinding []interface{} `json:"linkset_interface_binding,omitempty"`
}

type Vrid6InterfaceBinding struct {
	Flags int    `json:"flags,omitempty"`
	Id    int    `json:"id,omitempty"`
	Ifnum string `json:"ifnum,omitempty"`
	Vlan  int    `json:"vlan,omitempty"`
}

type RnatglobalBinding struct {
	RnatglobalAuditsyslogpolicyBinding []interface{} `json:"rnatglobal_auditsyslogpolicy_binding,omitempty"`
}

type FisChannelBinding struct {
	Ifnum     string `json:"ifnum,omitempty"`
	Name      string `json:"name,omitempty"`
	Ownernode int    `json:"ownernode,omitempty"`
}

type MapdomainBinding struct {
	MapdomainMapbmrBinding []interface{} `json:"mapdomain_mapbmr_binding,omitempty"`
	Name                   string        `json:"name,omitempty"`
}

type MapdomainMapbmrBinding struct {
	Mapbmrname string `json:"mapbmrname,omitempty"`
	Name       string `json:"name,omitempty"`
}
