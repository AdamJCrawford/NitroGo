package models

// pcp configuration structs
type PCPServer struct {
	Count              float64 `json:"__count,omitempty"`
	IPAddress          string  `json:"ipaddress,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	PCPProfile         string  `json:"pcpprofile,omitempty"`
	Port               int     `json:"port,omitempty"`
}

type PCPMap struct {
	Count              float64 `json:"__count,omitempty"`
	NatType            string  `json:"nattype,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	PCPAddr            int     `json:"pcpaddr,omitempty"`
	PCPDstIP           string  `json:"pcpdstip,omitempty"`
	PCPDstPort         int     `json:"pcpdstport,omitempty"`
	PCPLifetime        int     `json:"pcplifetime,omitempty"`
	PCPNatIP           string  `json:"pcpnatip,omitempty"`
	PCPNatPort         int     `json:"pcpnatport,omitempty"`
	PCPNounce          int     `json:"pcpnounce,omitempty"`
	PCPProtocol        string  `json:"pcpprotocol,omitempty"`
	PCPRefCnt          int     `json:"pcprefcnt,omitempty"`
	PCPSrcIP           string  `json:"pcpsrcip,omitempty"`
	PCPSrcPort         int     `json:"pcpsrcport,omitempty"`
	SubscrIP           string  `json:"subscrip,omitempty"`
}

type PCPProfile struct {
	AnnounceMultiCount int     `json:"announcemulticount,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Mapping            string  `json:"mapping,omitempty"`
	MaxMapLife         int     `json:"maxmaplife,omitempty"`
	MinMapLife         int     `json:"minmaplife,omitempty"`
	Name               string  `json:"name,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	Peer               string  `json:"peer,omitempty"`
	ThirdParty         string  `json:"thirdparty,omitempty"`
}
