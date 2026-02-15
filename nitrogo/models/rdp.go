package models

// rdp configuration structs
type Rdpclientprofile struct {
	Addusernameinrdpfile string   `json:"addusernameinrdpfile,omitempty"`
	Audiocapturemode     string   `json:"audiocapturemode,omitempty"`
	Builtin              []string `json:"builtin,omitempty"`
	Count                float64  `json:"__count,omitempty"`
	Feature              string   `json:"feature,omitempty"`
	Keyboardhook         string   `json:"keyboardhook,omitempty"`
	Multimonitorsupport  string   `json:"multimonitorsupport,omitempty"`
	Name                 string   `json:"name,omitempty"`
	Nextgenapiresource   string   `json:"_nextgenapiresource,omitempty"`
	Psk                  string   `json:"psk,omitempty"`
	Randomizerdpfilename string   `json:"randomizerdpfilename,omitempty"`
	Rdpcookievalidity    int      `json:"rdpcookievalidity,omitempty"`
	Rdpcustomparams      string   `json:"rdpcustomparams,omitempty"`
	Rdpfilename          string   `json:"rdpfilename,omitempty"`
	Rdphost              string   `json:"rdphost,omitempty"`
	Rdplinkattribute     string   `json:"rdplinkattribute,omitempty"`
	Rdplistener          string   `json:"rdplistener,omitempty"`
	Rdpurloverride       string   `json:"rdpurloverride,omitempty"`
	Rdpvalidateclientip  string   `json:"rdpvalidateclientip,omitempty"`
	Redirectclipboard    string   `json:"redirectclipboard,omitempty"`
	Redirectcomports     string   `json:"redirectcomports,omitempty"`
	Redirectdrives       string   `json:"redirectdrives,omitempty"`
	Redirectpnpdevices   string   `json:"redirectpnpdevices,omitempty"`
	Redirectprinters     string   `json:"redirectprinters,omitempty"`
	Videoplaybackmode    string   `json:"videoplaybackmode,omitempty"`
}

type Rdpserverprofile struct {
	Builtin            []string `json:"builtin,omitempty"`
	Count              float64  `json:"__count,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Name               string   `json:"name,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
	Psk                string   `json:"psk,omitempty"`
	Rdpip              string   `json:"rdpip,omitempty"`
	Rdpport            int      `json:"rdpport,omitempty"`
	Rdpredirection     string   `json:"rdpredirection,omitempty"`
}

type Rdpconnections struct {
	All                bool    `json:"all,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Endpointip         string  `json:"endpointip,omitempty"`
	Endpointport       int     `json:"endpointport,omitempty"`
	Nextgenapiresource string  `json:"_nextgenapiresource,omitempty"`
	Peid               int     `json:"peid,omitempty"`
	Targetip           string  `json:"targetip,omitempty"`
	Targetport         int     `json:"targetport,omitempty"`
	Username           string  `json:"username,omitempty"`
}
