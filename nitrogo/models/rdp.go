package models

// rdp configuration structs
type RDPClientProfile struct {
	AddUsernameInRDPFile string   `json:"addusernameinrdpfile,omitempty"`
	AudioCaptureMode     string   `json:"audiocapturemode,omitempty"`
	Builtin              []string `json:"builtin,omitempty"`
	Count                float64  `json:"__count,omitempty"`
	Feature              string   `json:"feature,omitempty"`
	KeyboardHook         string   `json:"keyboardhook,omitempty"`
	MultiMonitorSupport  string   `json:"multimonitorsupport,omitempty"`
	Name                 string   `json:"name,omitempty"`
	NextGenAPIResource   string   `json:"_nextgenapiresource,omitempty"`
	PSK                  string   `json:"psk,omitempty"`
	RandomizeRDPFileName string   `json:"randomizerdpfilename,omitempty"`
	RDPCookieValidity    int      `json:"rdpcookievalidity,omitempty"`
	RDPCustomParams      string   `json:"rdpcustomparams,omitempty"`
	RDPFileName          string   `json:"rdpfilename,omitempty"`
	RDPHost              string   `json:"rdphost,omitempty"`
	RDPLinkAttribute     string   `json:"rdplinkattribute,omitempty"`
	RDPListener          string   `json:"rdplistener,omitempty"`
	RDPURLOverride       string   `json:"rdpurloverride,omitempty"`
	RDPValidateClientIP  string   `json:"rdpvalidateclientip,omitempty"`
	RedirectClipboard    string   `json:"redirectclipboard,omitempty"`
	RedirectComPorts     string   `json:"redirectcomports,omitempty"`
	RedirectDrives       string   `json:"redirectdrives,omitempty"`
	RedirectPnpDevices   string   `json:"redirectpnpdevices,omitempty"`
	RedirectPrinters     string   `json:"redirectprinters,omitempty"`
	VideoPlaybackMode    string   `json:"videoplaybackmode,omitempty"`
}

type RDPServerProfile struct {
	Builtin            []string `json:"builtin,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Name               string   `json:"name,omitempty"`
	NextGenAPIResource string   `json:"_nextgenapiresource,omitempty"`
	PSK                string   `json:"psk,omitempty"`
	RDPIP              string   `json:"rdpip,omitempty"`
	RDPPort            int      `json:"rdpport,omitempty"`
	RDPRedirection     string   `json:"rdpredirection,omitempty"`
}

type RDPConnections struct {
	All                bool    `json:"all,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	EndpointIP         string  `json:"endpointip,omitempty"`
	EndpointPort       int     `json:"endpointport,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	PEID               int     `json:"peid,omitempty"`
	TargetIP           string  `json:"targetip,omitempty"`
	TargetPort         int     `json:"targetport,omitempty"`
	Username           string  `json:"username,omitempty"`
}
