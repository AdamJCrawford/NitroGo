// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type PQBinding struct {
	VServerName string `json:"vservername,omitempty"`
	PolicyName  string `json:"policyname,omitempty"`
	Rule        string `json:"rule,omitempty"`
	Priority    string `json:"priority,omitempty"`
	Weight      string `json:"weight,omitempty"`
	QDepth      string `json:"qdepth,omitempty"`
	PolQDepth   string `json:"polqdepth,omitempty"`
	Hits        string `json:"hits,omitempty"`
}

type PQPolicy struct {
	PolicyName string `json:"policyname,omitempty"`
	Rule       string `json:"rule,omitempty"`
	Priority   int    `json:"priority,omitempty"`
	Weight     int    `json:"weight,omitempty"`
	QDepth     int    `json:"qdepth,omitempty"`
	PolQDepth  int    `json:"polqdepth,omitempty"`
	Hits       string `json:"hits,omitempty"`
}
