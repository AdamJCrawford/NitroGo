package models

// utility configuration structs
type Raid struct {
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Install struct {
	A               bool   `json:"a,omitempty"`
	Async           bool   `json:"Async,omitempty"`
	Enhancedupgrade bool   `json:"enhancedupgrade,omitempty"`
	Id              int    `json:"id,omitempty"`
	L               bool   `json:"l,omitempty"`
	Resizeswapvar   bool   `json:"resizeswapvar,omitempty"`
	Url             string `json:"url,omitempty"`
	Y               bool   `json:"y,omitempty"`
}

type Traceroute6 struct {
	Host      string `json:"host,omitempty"`
	I         bool   `json:"I,omitempty"`
	M         int    `json:"m,omitempty"`
	N         bool   `json:"n,omitempty"`
	P         int    `json:"p,omitempty"`
	Packetlen int    `json:"packetlen,omitempty"`
	Q         int    `json:"q,omitempty"`
	R         bool   `json:"r,omitempty"`
	Response  string `json:"response,omitempty"`
	S         string `json:"s,omitempty"`
	T         int    `json:"T,omitempty"`
	V         bool   `json:"v,omitempty"`
	W         int    `json:"w,omitempty"`
}

type Ping struct {
	C        int    `json:"c,omitempty"`
	HostName string `json:"hostName,omitempty"`
	I        string `json:"I,omitempty"`
	I1       int    `json:"i,omitempty"`
	N        bool   `json:"n,omitempty"`
	P        string `json:"p,omitempty"`
	Q        bool   `json:"q,omitempty"`
	Response string `json:"response,omitempty"`
	S        string `json:"S,omitempty"`
	S1       int    `json:"s,omitempty"`
	T        int    `json:"T,omitempty"`
	T1       int    `json:"t,omitempty"`
}

type Techsupport struct {
	Adss               bool          `json:"adss,omitempty"`
	Authtoken          string        `json:"authtoken,omitempty"`
	Casenumber         string        `json:"casenumber,omitempty"`
	Description        string        `json:"description,omitempty"`
	File               string        `json:"file,omitempty"`
	Nextgenapiresource string        `json:"_nextgenapiresource,omitempty"`
	Nodes              []interface{} `json:"nodes,omitempty"`
	Partitionname      string        `json:"partitionname,omitempty"`
	Proxy              string        `json:"proxy,omitempty"`
	Response           string        `json:"response,omitempty"`
	Scope              string        `json:"scope,omitempty"`
	Servername         string        `json:"servername,omitempty"`
	Time               string        `json:"time,omitempty"`
	Upload             bool          `json:"upload,omitempty"`
}

type Traceroute struct {
	Host      string `json:"host,omitempty"`
	M         int    `json:"M,omitempty"`
	M1        int    `json:"m,omitempty"`
	N         bool   `json:"n,omitempty"`
	P         string `json:"P,omitempty"`
	P1        int    `json:"p,omitempty"`
	Packetlen int    `json:"packetlen,omitempty"`
	Q         int    `json:"q,omitempty"`
	R         bool   `json:"r,omitempty"`
	Response  string `json:"response,omitempty"`
	S         bool   `json:"S,omitempty"`
	S1        string `json:"s,omitempty"`
	T         int    `json:"T,omitempty"`
	T1        int    `json:"t,omitempty"`
	V         bool   `json:"v,omitempty"`
	W         int    `json:"w,omitempty"`
}

type Callhome struct {
	Anomalydetection     string   `json:"anomalydetection,omitempty"`
	Callhomestatus       []string `json:"callhomestatus,omitempty"`
	Emailaddress         string   `json:"emailaddress,omitempty"`
	Flashfirstfail       string   `json:"flashfirstfail,omitempty"`
	Flashlatestfailure   string   `json:"flashlatestfailure,omitempty"`
	Hbcustominterval     int      `json:"hbcustominterval,omitempty"`
	Hddfirstfail         string   `json:"hddfirstfail,omitempty"`
	Hddlatestfailure     string   `json:"hddlatestfailure,omitempty"`
	Ipaddress            string   `json:"ipaddress,omitempty"`
	Memthrefirstanomaly  string   `json:"memthrefirstanomaly,omitempty"`
	Memthrelatestanomaly string   `json:"memthrelatestanomaly,omitempty"`
	Mode                 string   `json:"mode,omitempty"`
	Nextgenapiresource   string   `json:"_nextgenapiresource,omitempty"`
	Nodeid               int      `json:"nodeid,omitempty"`
	Port                 int      `json:"port,omitempty"`
	Powfirstfail         string   `json:"powfirstfail,omitempty"`
	Powlatestfailure     string   `json:"powlatestfailure,omitempty"`
	Proxyauthservice     string   `json:"proxyauthservice,omitempty"`
	Proxymode            string   `json:"proxymode,omitempty"`
	Restartlatestfail    string   `json:"restartlatestfail,omitempty"`
	Rlfirsthighdrop      string   `json:"rlfirsthighdrop,omitempty"`
	Rllatesthighdrop     string   `json:"rllatesthighdrop,omitempty"`
	Sslcardfirstfailure  string   `json:"sslcardfirstfailure,omitempty"`
	Sslcardlatestfailure string   `json:"sslcardlatestfailure,omitempty"`
}

type Filesystemencryption struct {
	Effectivestate     string `json:"effectivestate,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Nodeid             int    `json:"nodeid,omitempty"`
	Ntimes0flash       int    `json:"ntimes0flash,omitempty"`
	Ntimes0var         int    `json:"ntimes0var,omitempty"`
	Passphrase         string `json:"passphrase,omitempty"`
	Supportedstate     string `json:"supportedstate,omitempty"`
}

type Ping6 struct {
	B        int    `json:"b,omitempty"`
	C        int    `json:"c,omitempty"`
	HostName string `json:"hostName,omitempty"`
	I        string `json:"I,omitempty"`
	I1       int    `json:"i,omitempty"`
	M        bool   `json:"m,omitempty"`
	N        bool   `json:"n,omitempty"`
	P        string `json:"p,omitempty"`
	Q        bool   `json:"q,omitempty"`
	Response string `json:"response,omitempty"`
	S        string `json:"S,omitempty"`
	S1       int    `json:"s,omitempty"`
	T        int    `json:"T,omitempty"`
	T1       int    `json:"t,omitempty"`
	V        int    `json:"V,omitempty"`
}
