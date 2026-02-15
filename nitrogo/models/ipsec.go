package models

// ipsec configuration structs
type Ipsecparameter struct {
	Encalgo               []string `json:"encalgo,omitempty"`
	Hashalgo              []string `json:"hashalgo,omitempty"`
	Ikeretryinterval      int      `json:"ikeretryinterval,omitempty"`
	Ikeversion            string   `json:"ikeversion,omitempty"`
	Lifetime              int      `json:"lifetime,omitempty"`
	Livenesscheckinterval int      `json:"livenesscheckinterval,omitempty"`
	Nextgenapiresource    string   `json:"_nextgenapiresource,omitempty"`
	Perfectforwardsecrecy string   `json:"perfectforwardsecrecy,omitempty"`
	Replaywindowsize      int      `json:"replaywindowsize,omitempty"`
	Responderonly         string   `json:"responderonly,omitempty"`
	Retransmissiontime    int      `json:"retransmissiontime,omitempty"`
}

type Ipsecprofile struct {
	Builtin               []string `json:"builtin,omitempty"`
	Count                 float64  `json:"__count,omitempty"`
	Encalgo               []string `json:"encalgo,omitempty"`
	Feature               string   `json:"feature,omitempty"`
	Hashalgo              []string `json:"hashalgo,omitempty"`
	Ikeretryinterval      int      `json:"ikeretryinterval,omitempty"`
	Ikeversion            string   `json:"ikeversion,omitempty"`
	Lifetime              int      `json:"lifetime,omitempty"`
	Livenesscheckinterval int      `json:"livenesscheckinterval,omitempty"`
	Name                  string   `json:"name,omitempty"`
	Nextgenapiresource    string   `json:"_nextgenapiresource,omitempty"`
	Peerpublickey         string   `json:"peerpublickey,omitempty"`
	Perfectforwardsecrecy string   `json:"perfectforwardsecrecy,omitempty"`
	Privatekey            string   `json:"privatekey,omitempty"`
	Psk                   string   `json:"psk,omitempty"`
	Publickey             string   `json:"publickey,omitempty"`
	Replaywindowsize      int      `json:"replaywindowsize,omitempty"`
	Responderonly         string   `json:"responderonly,omitempty"`
	Retransmissiontime    int      `json:"retransmissiontime,omitempty"`
}
