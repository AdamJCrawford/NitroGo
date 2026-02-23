package models

// lldp configuration structs
type LLDPParam struct {
	HoldTimeTxMult     int    `json:"holdtimetxmult,omitempty"`
	Mode               string `json:"mode,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	Timer              int    `json:"timer,omitempty"`
}

type LLDPNeighbors struct {
	AutonegAdvertised  string  `json:"autonegadvertised,omitempty"`
	AutonegEnabled     string  `json:"autonegenabled,omitempty"`
	AutonegMauType     string  `json:"autonegmautype,omitempty"`
	AutonegSupport     string  `json:"autonegsupport,omitempty"`
	ChassisID          string  `json:"chassisid,omitempty"`
	ChassisIDSubtype   string  `json:"chassisidsubtype,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Flag               int     `json:"flag,omitempty"`
	IFNum              string  `json:"ifnum,omitempty"`
	IFNumber           int     `json:"ifnumber,omitempty"`
	IFType             string  `json:"iftype,omitempty"`
	LinkAggrCapable    string  `json:"linkaggrcapable,omitempty"`
	LinkAggrEnabled    string  `json:"linkaggrenabled,omitempty"`
	LinkAggrID         int     `json:"linkaggrid,omitempty"`
	MgmtAddress        string  `json:"mgmtaddress,omitempty"`
	MgmtAddressSubtype string  `json:"mgmtaddresssubtype,omitempty"`
	MTU                int     `json:"mtu,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	NodeID             int     `json:"nodeid,omitempty"`
	PortDescription    string  `json:"portdescription,omitempty"`
	PortID             string  `json:"portid,omitempty"`
	PortIDSubtype      string  `json:"portidsubtype,omitempty"`
	PortProtoEnabled   int     `json:"portprotoenabled,omitempty"`
	PortProtoID        int     `json:"portprotoid,omitempty"`
	PortProtoSupported int     `json:"portprotosupported,omitempty"`
	PortVLANID         int     `json:"portvlanid,omitempty"`
	ProtocolID         string  `json:"protocolid,omitempty"`
	Sys                string  `json:"sys,omitempty"`
	SysCapabilities    string  `json:"syscapabilities,omitempty"`
	SysCapEnabled      string  `json:"syscapenabled,omitempty"`
	SysDesc            string  `json:"sysdesc,omitempty"`
	TTL                int     `json:"ttl,omitempty"`
	VLAN               string  `json:"vlan,omitempty"`
	VLANID             int     `json:"vlanid,omitempty"`
}
