package models

// ntp configuration structs
type NTPServer struct {
	AutoKey            bool    `json:"autokey,omitempty"`
	Count              float64 `json:"__count,omitempty"`
	Key                int     `json:"key,omitempty"`
	MaxPoll            int     `json:"maxpoll,omitempty"`
	MinPoll            int     `json:"minpoll,omitempty"`
	NextGenAPIResource string  `json:"_nextgenapiresource,omitempty"`
	PreferredNTPServer string  `json:"preferredntpserver,omitempty"`
	ServerIP           string  `json:"serverip,omitempty"`
	ServerName         string  `json:"servername,omitempty"`
}

type NTPParam struct {
	Authentication     string `json:"authentication,omitempty"`
	AutoKeyLogSec      int    `json:"autokeylogsec,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	RevokeLogSec       int    `json:"revokelogsec,omitempty"`
	TrustedKey         []any  `json:"trustedkey,omitempty"`
}

type NTPSync struct {
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	State              string `json:"state,omitempty"`
}

type NTPStatus struct {
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
	Response           string `json:"response,omitempty"`
}
