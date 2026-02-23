package models

// bfd configuration structs
type BFDSession struct {
	AdminDown                         bool    `json:"admindown,omitempty"`
	Count                             float64 `json:"__count,omitempty"`
	CurrentOwnerPE                    int     `json:"currentownerpe,omitempty"`
	LocalDiagnotic                    int     `json:"localdiagnotic,omitempty"`
	LocalDiscriminator                int     `json:"localdiscriminator,omitempty"`
	LocalIP                           string  `json:"localip,omitempty"`
	LocalPort                         int     `json:"localport,omitempty"`
	MinimumReceiveInterval            int     `json:"minimumreceiveinterval,omitempty"`
	MinimumTransmitInterval           int     `json:"minimumtransmitinterval,omitempty"`
	MultiHop                          bool    `json:"multihop,omitempty"`
	Multiplier                        int     `json:"multiplier,omitempty"`
	NegotiatedMinimumReceiveInterval  int     `json:"negotiatedminimumreceiveinterval,omitempty"`
	NegotiatedMinimumTransmitInterval int     `json:"negotiatedminimumtransmitinterval,omitempty"`
	NextGenAPIResource                string  `json:"_nextgenapiresource,omitempty"`
	OriginalOwnerPE                   int     `json:"originalownerpe,omitempty"`
	OwnerNode                         int     `json:"ownernode,omitempty"`
	Passive                           bool    `json:"passive,omitempty"`
	RemoteDiscriminator               int     `json:"remotediscriminator,omitempty"`
	RemoteIP                          string  `json:"remoteip,omitempty"`
	RemoteMultiplier                  int     `json:"remotemultiplier,omitempty"`
	RemotePort                        int     `json:"remoteport,omitempty"`
	State                             string  `json:"state,omitempty"`
	VLAN                              int     `json:"vlan,omitempty"`
}
