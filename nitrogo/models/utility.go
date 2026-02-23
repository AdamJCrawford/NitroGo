package models

// utility configuration structs
type Raid struct {
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
}

type Install struct {
	A               bool   `json:"a,omitempty"`
	Async           bool   `json:"Async,omitempty"`
	EnhancedUpgrade bool   `json:"enhancedupgrade,omitempty"`
	ID              int    `json:"id,omitempty"`
	L               bool   `json:"l,omitempty"`
	ResizeSwapVar   bool   `json:"resizeswapvar,omitempty"`
	URL             string `json:"url,omitempty"`
	Y               bool   `json:"y,omitempty"`
}

type Traceroute6 struct {
	Host      string `json:"host,omitempty"`
	I         bool   `json:"I,omitempty"`
	M         int    `json:"m,omitempty"`
	N         bool   `json:"n,omitempty"`
	P         int    `json:"p,omitempty"`
	PacketLen int    `json:"packetlen,omitempty"`
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

type TechSupport struct {
	ADSS               bool          `json:"adss,omitempty"`
	AuthToken          string        `json:"authtoken,omitempty"`
	CaseNumber         string        `json:"casenumber,omitempty"`
	Description        string        `json:"description,omitempty"`
	File               string        `json:"file,omitempty"`
	NextGenAPIResource string        `json:"_nextgenapiresource,omitempty"`
	Nodes              []interface{} `json:"nodes,omitempty"`
	PartitionName      string        `json:"partitionname,omitempty"`
	Proxy              string        `json:"proxy,omitempty"`
	Response           string        `json:"response,omitempty"`
	Scope              string        `json:"scope,omitempty"`
	ServerName         string        `json:"servername,omitempty"`
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
	PacketLen int    `json:"packetlen,omitempty"`
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

type CallHome struct {
	AnomalyDetection     string   `json:"anomalydetection,omitempty"`
	CallHomeStatus       []string `json:"callhomestatus,omitempty"`
	EmailAddress         string   `json:"emailaddress,omitempty"`
	FlashFirstFail       string   `json:"flashfirstfail,omitempty"`
	FlashLatestFailure   string   `json:"flashlatestfailure,omitempty"`
	HBCustomInterval     int      `json:"hbcustominterval,omitempty"`
	HDDFirstFail         string   `json:"hddfirstfail,omitempty"`
	HDDLatestFailure     string   `json:"hddlatestfailure,omitempty"`
	IPAddress            string   `json:"ipaddress,omitempty"`
	MemThreFirstAnomaly  string   `json:"memthrefirstanomaly,omitempty"`
	MemThreLatestAnomaly string   `json:"memthrelatestanomaly,omitempty"`
	Mode                 string   `json:"mode,omitempty"`
	NextGenAPIResource   string   `json:"_nextgenapiresource,omitempty"`
	NodeID               int      `json:"nodeid,omitempty"`
	Port                 int      `json:"port,omitempty"`
	PowFirstFail         string   `json:"powfirstfail,omitempty"`
	PowLatestFailure     string   `json:"powlatestfailure,omitempty"`
	ProxyAuthService     string   `json:"proxyauthservice,omitempty"`
	ProxyMode            string   `json:"proxymode,omitempty"`
	RestartLatestFail    string   `json:"restartlatestfail,omitempty"`
	RLFirstHighDrop      string   `json:"rlfirsthighdrop,omitempty"`
	RLLatestHighDrop     string   `json:"rllatesthighdrop,omitempty"`
	SSLCardFirstFailure  string   `json:"sslcardfirstfailure,omitempty"`
	SSLCardLatestFailure string   `json:"sslcardlatestfailure,omitempty"`
}

type FileSystemEncryption struct {
	EffectiveState     string `json:"effectivestate,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	NodeID             int    `json:"nodeid,omitempty"`
	NTimes0Flash       int    `json:"ntimes0flash,omitempty"`
	NTimes0Var         int    `json:"ntimes0var,omitempty"`
	Passphrase         string `json:"passphrase,omitempty"`
	SupportedState     string `json:"supportedstate,omitempty"`
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
