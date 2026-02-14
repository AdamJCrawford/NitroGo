// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Rdpclientprofile struct {
	Name                 string `json:"name,omitempty"`
	Rdpurloverride       string `json:"rdpurloverride,omitempty"`
	Redirectclipboard    string `json:"redirectclipboard,omitempty"`
	Redirectdrives       string `json:"redirectdrives,omitempty"`
	Redirectprinters     string `json:"redirectprinters,omitempty"`
	Redirectcomports     string `json:"redirectcomports,omitempty"`
	Redirectpnpdevices   string `json:"redirectpnpdevices,omitempty"`
	Keyboardhook         string `json:"keyboardhook,omitempty"`
	Audiocapturemode     string `json:"audiocapturemode,omitempty"`
	Videoplaybackmode    string `json:"videoplaybackmode,omitempty"`
	Multimonitorsupport  string `json:"multimonitorsupport,omitempty"`
	Rdpcookievalidity    int    `json:"rdpcookievalidity,omitempty"`
	Addusernameinrdpfile string `json:"addusernameinrdpfile,omitempty"`
	Rdpfilename          string `json:"rdpfilename,omitempty"`
	Rdphost              string `json:"rdphost,omitempty"`
	Rdplistener          string `json:"rdplistener,omitempty"`
	Rdpcustomparams      string `json:"rdpcustomparams,omitempty"`
	Psk                  string `json:"psk,omitempty"`
	Randomizerdpfilename string `json:"randomizerdpfilename,omitempty"`
	Rdplinkattribute     string `json:"rdplinkattribute,omitempty"`
	Rdpvalidateclientip  string `json:"rdpvalidateclientip,omitempty"`
	Builtin              string `json:"builtin,omitempty"`
	Feature              string `json:"feature,omitempty"`
	Nextgenapiresource   string `json:"_nextgenapiresource,omitempty"`
}

type Rdpconnections struct {
	Username           string `json:"username,omitempty"`
	All                bool   `json:"all,omitempty"`
	Endpointip         string `json:"endpointip,omitempty"`
	Endpointport       string `json:"endpointport,omitempty"`
	Targetip           string `json:"targetip,omitempty"`
	Targetport         string `json:"targetport,omitempty"`
	Peid               string `json:"peid,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}

type Rdpserverprofile struct {
	Name               string `json:"name,omitempty"`
	Rdpip              string `json:"rdpip,omitempty"`
	Rdpport            int    `json:"rdpport,omitempty"`
	Psk                string `json:"psk,omitempty"`
	Rdpredirection     string `json:"rdpredirection,omitempty"`
	Builtin            string `json:"builtin,omitempty"`
	Feature            string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}
