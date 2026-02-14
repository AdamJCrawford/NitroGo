// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Application struct {
	Apptemplatefilename string `json:"apptemplatefilename,omitempty"`
	Appname             string `json:"appname,omitempty"`
	Deploymentfilename  string `json:"deploymentfilename,omitempty"`
}
