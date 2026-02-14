// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Ping6 struct {
	B        int    `json:"b,omitempty"`
	C        int    `json:"c,omitempty"`
	I        int    `json:"i,omitempty"`
	I_       string `json:"I,omitempty"`
	M        bool   `json:"m,omitempty"`
	N        bool   `json:"n,omitempty"`
	P        string `json:"p,omitempty"`
	Q        bool   `json:"q,omitempty"`
	S        int    `json:"s,omitempty"`
	V        int    `json:"V,omitempty"`
	S_       string `json:"S,omitempty"`
	T        int    `json:"T,omitempty"`
	T_       int    `json:"t,omitempty"`
	HostName string `json:"hostName,omitempty"`
	Response string `json:"response,omitempty"`
}

type Raid struct {
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Techsupport struct {
	Scope              string `json:"scope,omitempty"`
	Partitionname      string `json:"partitionname,omitempty"`
	Upload             bool   `json:"upload,omitempty"`
	Proxy              string `json:"proxy,omitempty"`
	Casenumber         string `json:"casenumber,omitempty"`
	File               string `json:"file,omitempty"`
	Description        string `json:"description,omitempty"`
	Authtoken          string `json:"authtoken,omitempty"`
	Time               string `json:"time,omitempty"`
	Adss               bool   `json:"adss,omitempty"`
	Nodes              []int  `json:"nodes,omitempty"`
	Response           string `json:"response,omitempty"`
	Servername         string `json:"servername,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Traceroute struct {
	S         bool   `json:"S,omitempty"`
	N         bool   `json:"n,omitempty"`
	R         bool   `json:"r,omitempty"`
	V         bool   `json:"v,omitempty"`
	M         int    `json:"M,omitempty"`
	M_        int    `json:"m,omitempty"`
	P         string `json:"P,omitempty"`
	P_        int    `json:"p,omitempty"`
	Q         int    `json:"q,omitempty"`
	S_        string `json:"s,omitempty"`
	T         int    `json:"T,omitempty"`
	T_        int    `json:"t,omitempty"`
	W         int    `json:"w,omitempty"`
	Host      string `json:"host,omitempty"`
	Packetlen int    `json:"packetlen,omitempty"`
	Response  string `json:"response,omitempty"`
}

type Traceroute6 struct {
	N         bool   `json:"n,omitempty"`
	I         bool   `json:"I,omitempty"`
	R         bool   `json:"r,omitempty"`
	V         bool   `json:"v,omitempty"`
	M         int    `json:"m,omitempty"`
	P         int    `json:"p,omitempty"`
	Q         int    `json:"q,omitempty"`
	S         string `json:"s,omitempty"`
	T         int    `json:"T,omitempty"`
	W         int    `json:"w,omitempty"`
	Host      string `json:"host,omitempty"`
	Packetlen int    `json:"packetlen,omitempty"`
	Response  string `json:"response,omitempty"`
}

type Callhome struct {
	Nodeid               int    `json:"nodeid,omitempty"`
	Mode                 string `json:"mode,omitempty"`
	Emailaddress         string `json:"emailaddress,omitempty"`
	Hbcustominterval     int    `json:"hbcustominterval,omitempty"`
	Proxymode            string `json:"proxymode,omitempty"`
	Ipaddress            string `json:"ipaddress,omitempty"`
	Proxyauthservice     string `json:"proxyauthservice,omitempty"`
	Port                 int    `json:"port,omitempty"`
	Sslcardfirstfailure  string `json:"sslcardfirstfailure,omitempty"`
	Sslcardlatestfailure string `json:"sslcardlatestfailure,omitempty"`
	Powfirstfail         string `json:"powfirstfail,omitempty"`
	Powlatestfailure     string `json:"powlatestfailure,omitempty"`
	Hddfirstfail         string `json:"hddfirstfail,omitempty"`
	Hddlatestfailure     string `json:"hddlatestfailure,omitempty"`
	Flashfirstfail       string `json:"flashfirstfail,omitempty"`
	Flashlatestfailure   string `json:"flashlatestfailure,omitempty"`
	Rlfirsthighdrop      string `json:"rlfirsthighdrop,omitempty"`
	Rllatesthighdrop     string `json:"rllatesthighdrop,omitempty"`
	Restartlatestfail    string `json:"restartlatestfail,omitempty"`
	Memthrefirstanomaly  string `json:"memthrefirstanomaly,omitempty"`
	Memthrelatestanomaly string `json:"memthrelatestanomaly,omitempty"`
	Callhomestatus       string `json:"callhomestatus,omitempty"`
	Anomalydetection     string `json:"anomalydetection,omitempty"`
	Nextgenapiresource   string `json:"_nextgenapiresource,omitempty"`
}

type Filesystemencryption struct {
	Ntimes0flash       int    `json:"ntimes0flash,omitempty"`
	Ntimes0var         int    `json:"ntimes0var,omitempty"`
	Passphrase         string `json:"passphrase,omitempty"`
	Nodeid             int    `json:"nodeid,omitempty"`
	Supportedstate     string `json:"supportedstate,omitempty"`
	Effectivestate     string `json:"effectivestate,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Install struct {
	Url             string `json:"url,omitempty"`
	Y               bool   `json:"y,omitempty"`
	L               bool   `json:"l,omitempty"`
	A               bool   `json:"a,omitempty"`
	Enhancedupgrade bool   `json:"enhancedupgrade,omitempty"`
	Resizeswapvar   bool   `json:"resizeswapvar,omitempty"`
	Async           bool   `json:"Async,omitempty"`
	Id              string `json:"id,omitempty"`
}

type Ping struct {
	C        int    `json:"c,omitempty"`
	I        int    `json:"i,omitempty"`
	I_       string `json:"I,omitempty"`
	N        bool   `json:"n,omitempty"`
	P        string `json:"p,omitempty"`
	Q        bool   `json:"q,omitempty"`
	S        int    `json:"s,omitempty"`
	S_       string `json:"S,omitempty"`
	T        int    `json:"T,omitempty"`
	T_       int    `json:"t,omitempty"`
	HostName string `json:"hostName,omitempty"`
	Response string `json:"response,omitempty"`
}
