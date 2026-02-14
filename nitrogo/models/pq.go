// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Pqbinding struct {
	Vservername string `json:"vservername,omitempty"`
	Policyname  string `json:"policyname,omitempty"`
	Rule        string `json:"rule,omitempty"`
	Priority    string `json:"priority,omitempty"`
	Weight      string `json:"weight,omitempty"`
	Qdepth      string `json:"qdepth,omitempty"`
	Polqdepth   string `json:"polqdepth,omitempty"`
	Hits        string `json:"hits,omitempty"`
}

type Pqpolicy struct {
	Policyname string `json:"policyname,omitempty"`
	Rule       string `json:"rule,omitempty"`
	Priority   int    `json:"priority,omitempty"`
	Weight     int    `json:"weight,omitempty"`
	Qdepth     int    `json:"qdepth,omitempty"`
	Polqdepth  int    `json:"polqdepth,omitempty"`
	Hits       string `json:"hits,omitempty"`
}
