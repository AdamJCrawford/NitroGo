// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2026

package models

type Hafailover struct {
	Force bool `json:"force,omitempty"`
}

type Hanode struct {
	Id                   int    `json:"id"`
	Ipaddress            string `json:"ipaddress,omitempty"`
	Inc                  string `json:"inc,omitempty"`
	Rpcnodepassword      string `json:"rpcnodepassword,omitempty"`
	Hastatus             string `json:"hastatus,omitempty"`
	Hasync               string `json:"hasync,omitempty"`
	Haprop               string `json:"haprop,omitempty"`
	Hellointerval        int    `json:"hellointerval,omitempty"`
	Deadinterval         int    `json:"deadinterval,omitempty"`
	Failsafe             string `json:"failsafe,omitempty"`
	Maxflips             int    `json:"maxflips,omitempty"`
	Maxfliptime          int    `json:"maxfliptime,omitempty"`
	Syncvlan             int    `json:"syncvlan,omitempty"`
	Syncstatusstrictmode string `json:"syncstatusstrictmode,omitempty"`
	Name                 string `json:"name,omitempty"`
	Flags                string `json:"flags,omitempty"`
	State                string `json:"state,omitempty"`
	Enaifaces            string `json:"enaifaces,omitempty"`
	Disifaces            string `json:"disifaces,omitempty"`
	Hamonifaces          string `json:"hamonifaces,omitempty"`
	Haheartbeatifaces    string `json:"haheartbeatifaces,omitempty"`
	Pfifaces             string `json:"pfifaces,omitempty"`
	Ifaces               string `json:"ifaces,omitempty"`
	Netmask              string `json:"netmask,omitempty"`
	Ssl2                 string `json:"ssl2,omitempty"`
	Masterstatetime      string `json:"masterstatetime,omitempty"`
	Routemonitor         string `json:"routemonitor,omitempty"`
	Curflips             string `json:"curflips,omitempty"`
	Completedfliptime    string `json:"completedfliptime,omitempty"`
	Routemonitorstate    string `json:"routemonitorstate,omitempty"`
	Hasyncfailurereason  string `json:"hasyncfailurereason,omitempty"`
	Nextgenapiresource   string `json:"_nextgenapiresource,omitempty"`
}

type Hanodepartialfailureinterfacesbinding struct {
	Pfifaces     string `json:"pfifaces,omitempty"`
	Id           int    `json:"id,omitempty"`
	Routemonitor string `json:"routemonitor,omitempty"`
}

type Hanoderoutemonitor6binding struct {
	Routemonitor      string `json:"routemonitor,omitempty"`
	Netmask           string `json:"netmask,omitempty"`
	Flags             int    `json:"flags,omitempty"`
	Routemonitorstate string `json:"routemonitorstate,omitempty"`
	Id                int    `json:"id"`
}

type Hanoderoutemonitorbinding struct {
	Routemonitor      string `json:"routemonitor,omitempty"`
	Netmask           string `json:"netmask,omitempty"`
	Flags             int    `json:"flags,omitempty"`
	Routemonitorstate string `json:"routemonitorstate,omitempty"`
	Id                int    `json:"id"`
}

type Hasync struct {
	Force bool   `json:"force,omitempty"`
	Save  string `json:"save,omitempty"`
}

type Hafiles struct {
	Mode []string `json:"mode,omitempty"`
}

type Hanodebinding struct {
	Id int `json:"id,omitempty"`
}

type Hanodecibinding struct {
	Enaifaces    string `json:"enaifaces,omitempty"`
	Id           int    `json:"id,omitempty"`
	Routemonitor string `json:"routemonitor,omitempty"`
}

type Hanodefisbinding struct {
	Name         string `json:"name,omitempty"`
	Enaifaces    string `json:"enaifaces,omitempty"`
	Id           int    `json:"id,omitempty"`
	Routemonitor string `json:"routemonitor,omitempty"`
}

type Hasyncfailures struct {
	Response           string `json:"response,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}
