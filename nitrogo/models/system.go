// Derived from Citrix ADC Nitro Go SDK (https://github.com/netscaler/adc-nitro-go)
// Originally licensed under the Apache License, Version 2.0
// Modifications by Adam Crawford, 2025

package models

type SystemStatus struct {
	AddiMgmtCPUUsagePcnt float64 `json:"addimgmtcpuusagepcnt,omitempty"`
	AuxTemp0             int     `json:"auxtemp0,omitempty"`
	AuxTemp1             int     `json:"auxtemp1,omitempty"`
	AuxTemp2             int     `json:"auxtemp2,omitempty"`
	AuxTemp3             int     `json:"auxtemp3,omitempty"`
	AuxVolt0             float64 `json:"auxvolt0,omitempty"`
	AuxVolt1             float64 `json:"auxvolt1,omitempty"`
	AuxVolt2             float64 `json:"auxvolt2,omitempty"`
	AuxVolt3             float64 `json:"auxvolt3,omitempty"`
	AuxVolt4             float64 `json:"auxvolt4,omitempty"`
	AuxVolt5             float64 `json:"auxvolt5,omitempty"`
	AuxVolt6             float64 `json:"auxvolt6,omitempty"`
	AuxVolt7             float64 `json:"auxvolt7,omitempty"`
	ClearStats           string  `json:"clearstats,omitempty"`
	Cpu0Temp             int     `json:"cpu0temp,omitempty"`
	Cpu1Temp             int     `json:"cpu1temp,omitempty"`
	CpuFan0Apeed         int     `json:"cpufan0speed,omitempty"`
	CpuFan1Apeed         int     `json:"cpufan1speed,omitempty"`
	CPUUsage             string  `json:"cpuusage,omitempty"`
	CPUUsagePcnt         float64 `json:"cpuusagepcnt,omitempty"`
	Disk0Avail           int     `json:"disk0avail,omitempty"`
	Disk0PerUsage        int     `json:"disk0perusage,omitempty"`
	Disk0Size            int     `json:"disk0size,omitempty"`
	Disk0Used            int     `json:"disk0used,omitempty"`
	Disk1Avail           int     `json:"disk1avail,omitempty"`
	Disk1perUsage        int     `json:"disk1perusage,omitempty"`
	Disk1Size            int     `json:"disk1size,omitempty"`
	Disk1Used            int     `json:"disk1used,omitempty"`
	Fan0Speed            int     `json:"fan0speed,omitempty"`
	Fan2Speed            int     `json:"fan2speed,omitempty"`
	Fan3Speed            int     `json:"fan3speed,omitempty"`
	Fan4Speed            int     `json:"fan4speed,omitempty"`
	Fan5Speed            int     `json:"fan5speed,omitempty"`
	FanSpeed             int     `json:"fanspeed,omitempty"`
	InternalTemp         int     `json:"internaltemp,omitempty"`
	MasterCPUUsage       string  `json:"mastercpuusage,omitempty"`
	MemSizeMB            string  `json:"memsizemb,omitempty"`
	MemUsagePcnt         float64 `json:"memusagepcnt,omitempty"`
	MemUseInMB           string  `json:"memuseinmb,omitempty"`
	MgmtCPU0UsagePcnt    float64 `json:"mgmtcpu0usagepcnt,omitempty"`
	MgmtCPUUsageCcnt     float64 `json:"mgmtcpuusagepcnt,omitempty"`
	NumCPUs              string  `json:"numcpus,omitempty"`
	PktCPUUsagePcnt      float64 `json:"pktcpuusagepcnt,omitempty"`
	PowerSupply1Status   string  `json:"powersupply1status,omitempty"`
	PowerSupply2Status   string  `json:"powersupply2status,omitempty"`
	PowerSupply3Status   string  `json:"powersupply3status,omitempty"`
	PowerSupply4Status   string  `json:"powersupply4status,omitempty"`
	ResCPUUsage          string  `json:"rescpuusage,omitempty"`
	ResCPUUsagePcnt      float64 `json:"rescpuusagepcnt,omitempty"`
	SlaveCPUUsage        string  `json:"slavecpuusage,omitempty"`
	StartTime            string  `json:"starttime,omitempty"`
	StartTimeLocal       string  `json:"starttimelocal,omitempty"`
	SystemFanSpeed       int     `json:"systemfanspeed,omitempty"`
	TimeSinceStart       string  `json:"timesincestart,omitempty"`
	VoltageV12n          float64 `json:"voltagev12n,omitempty"`
	VoltageV12p          float64 `json:"voltagev12p,omitempty"`
	VoltageV33Main       float64 `json:"voltagev33main,omitempty"`
	VoltageV33Stby       float64 `json:"voltagev33stby,omitempty"`
	VoltageV5n           float64 `json:"voltagev5n,omitempty"`
	VoltageV5p           float64 `json:"voltagev5p,omitempty"`
	VoltageV5sb          float64 `json:"voltagev5sb,omitempty"`
	VoltageVBat          float64 `json:"voltagevbat,omitempty"`
	VoltageVCC0          float64 `json:"voltagevcc0,omitempty"`
	VoltageVCC1          float64 `json:"voltagevcc1,omitempty"`
	VoltageVSen2         float64 `json:"voltagevsen2,omitempty"`
	VoltageVTT           float64 `json:"voltagevtt,omitempty"`
}

type SystemExtraMgmtCPU struct {
	Nodeid             int    `json:"nodeid,omitempty"`
	ConfiguredState    string `json:"configuredstate,omitempty"`
	EffectiveState     string `json:"effectivestate,omitempty"`
	NextGenAPIResource string `json:"_nextgenapiresource,omitempty"`
}
