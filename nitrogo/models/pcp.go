package models

// pcp configuration structs
type Pcpserver struct {
	Count              float64 `json:"__count,omitempty"`
	Ipaddress          string  `json:"ipaddress,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Pcpprofile         string  `json:"pcpprofile,omitempty"`
	Port               int     `json:"port,omitempty"`
}

type Pcpmap struct {
	Count              float64 `json:"__count,omitempty"`
	Nattype            string  `json:"nattype,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Pcpaddr            int     `json:"pcpaddr,omitempty"`
	Pcpdstip           string  `json:"pcpdstip,omitempty"`
	Pcpdstport         int     `json:"pcpdstport,omitempty"`
	Pcplifetime        int     `json:"pcplifetime,omitempty"`
	Pcpnatip           string  `json:"pcpnatip,omitempty"`
	Pcpnatport         int     `json:"pcpnatport,omitempty"`
	Pcpnounce          int     `json:"pcpnounce,omitempty"`
	Pcpprotocol        string  `json:"pcpprotocol,omitempty"`
	Pcprefcnt          int     `json:"pcprefcnt,omitempty"`
	Pcpsrcip           string  `json:"pcpsrcip,omitempty"`
	Pcpsrcport         int     `json:"pcpsrcport,omitempty"`
	Subscrip           string  `json:"subscrip,omitempty"`
}

type Pcpprofile struct {
	Announcemulticount int     `json:"announcemulticount,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Mapping            string  `json:"mapping,omitempty"`
	Maxmaplife         int     `json:"maxmaplife,omitempty"`
	Minmaplife         int     `json:"minmaplife,omitempty"`
	Name               string  `json:"name,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Peer               string  `json:"peer,omitempty"`
	Thirdparty         string  `json:"thirdparty,omitempty"`
}
