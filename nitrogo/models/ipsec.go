package models

// ipsec configuration structs
type IPSECParameter struct {
	EncAlgo               []string `json:"encalgo,omitempty"`
	HashAlgo              []string `json:"hashalgo,omitempty"`
	IKERetryInterval      int      `json:"ikeretryinterval,omitempty"`
	IKEVersion            string   `json:"ikeversion,omitempty"`
	Lifetime              int      `json:"lifetime,omitempty"`
	LivenessCheckInterval int      `json:"livenesscheckinterval,omitempty"`
	NextGenAPIResource    string   `json:"_nextgenapiresource,omitempty"`
	PerfectForwardSecrecy string   `json:"perfectforwardsecrecy,omitempty"`
	ReplayWindowSize      int      `json:"replaywindowsize,omitempty"`
	ResponderOnly         string   `json:"responderonly,omitempty"`
	RetransmissionTime    int      `json:"retransmissiontime,omitempty"`
}

type IPSECProfile struct {
	Builtin               []string `json:"builtin,omitempty"`
	Count                 float64  `json:"__count,omitempty"`
	EncAlgo               []string `json:"encalgo,omitempty"`
	Feature               string   `json:"feature,omitempty"`
	HashAlgo              []string `json:"hashalgo,omitempty"`
	IKERetryInterval      int      `json:"ikeretryinterval,omitempty"`
	IKEVersion            string   `json:"ikeversion,omitempty"`
	Lifetime              int      `json:"lifetime,omitempty"`
	LivenessCheckInterval int      `json:"livenesscheckinterval,omitempty"`
	Name                  string   `json:"name,omitempty"`
	NextGenAPIResource    string   `json:"_nextgenapiresource,omitempty"`
	PeerPublicKey         string   `json:"peerpublickey,omitempty"`
	PerfectForwardSecrecy string   `json:"perfectforwardsecrecy,omitempty"`
	PrivateKey            string   `json:"privatekey,omitempty"`
	PSK                   string   `json:"psk,omitempty"`
	PublicKey             string   `json:"publickey,omitempty"`
	ReplayWindowSize      int      `json:"replaywindowsize,omitempty"`
	ResponderOnly         string   `json:"responderonly,omitempty"`
	RetransmissionTime    int      `json:"retransmissiontime,omitempty"`
}
