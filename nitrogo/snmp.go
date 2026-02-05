package nitrogo

const (
	snmpAlarmURL               = "/nitro/v1/config/snmpalarm"
	snmpCommunityURL           = "/nitro/v1/config/snmpcommunity"
	snmpEngineIdURL            = "/nitro/v1/config/snmpengineid"
	snmpGroupURL               = "/nitro/v1/config/snmpgroup"
	snmpManagerURL             = "/nitro/v1/config/snmpmanager"
	snmpMIBURL                 = "/nitro/v1/config/snmpmib"
	snmpOIDURL                 = "/nitro/v1/config/snmpoid"
	snmpOptionURL              = "/nitro/v1/config/snmpoption"
	snmpTrapURL                = "/nitro/v1/config/snmptrap"
	snmpTrapBindingURL         = "/nitro/v1/config/snmptrap_binding"
	snmpTrapSNMPUserBindingURL = "/nitro/v1/config/snmptrap_snmpuser_binding"
	snmpUserURL                = "/nitro/v1/config/snmpuser"
	snmpViewURL                = "/nitro/v1/config/snmpview"
)

// SNMP(Simple Network Management Protocol) configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/snmp/snmp
type SNMPService struct {
	client *Client
}

// snmpalarm
func (s *SNMPService) UpdateSNMPAlarm()  {}
func (s *SNMPService) UnsetSNMPAlarm()   {}
func (s *SNMPService) EnableSNMPAlarm()  {}
func (s *SNMPService) DisableSNMPAlarm() {}
func (s *SNMPService) GetAllSNMPAlarm()  {}
func (s *SNMPService) GetSNMPAlarm()     {}
func (s *SNMPService) CountSNMPAlarm()   {}

// snmpcommunity
func (s *SNMPService) AddSNMPCommunity()    {}
func (s *SNMPService) DeleteSNMPCommunity() {}
func (s *SNMPService) GetAllSNMPCommunity() {}
func (s *SNMPService) GetSNMPCommunity()    {}
func (s *SNMPService) CountSNMPCommunity()  {}

// snmpengineid
func (s *SNMPService) UpdateSNMPEngineId() {}
func (s *SNMPService) UnsetSNMPEngineId()  {}
func (s *SNMPService) GetAllSNMPEngineId() {}
func (s *SNMPService) GetSNMPEngineId()    {}
func (s *SNMPService) CountSNMPEngineId()  {}

// snmpgroup
func (s *SNMPService) AddSNMPGroup()    {}
func (s *SNMPService) DeleteSNMPGroup() {}
func (s *SNMPService) GetAllSNMPGroup() {}
func (s *SNMPService) GetSNMPGroup()    {}
func (s *SNMPService) CountSNMPGroup()  {}

// snmpmanager
func (s *SNMPService) AddSNMPManager()    {}
func (s *SNMPService) DeleteSNMPManager() {}
func (s *SNMPService) UpdateSNMPManager() {}
func (s *SNMPService) UnsetSNMPManager()  {}
func (s *SNMPService) GetAllSNMPManager() {}
func (s *SNMPService) CountSNMPManager()  {}

// snmpmib
func (s *SNMPService) UpdateSNMPMIB() {}
func (s *SNMPService) UnsetSNMPMIB()  {}
func (s *SNMPService) GetAllSNMPMIB() {}
func (s *SNMPService) GetSNMPMIB()    {}
func (s *SNMPService) CountSNMPMIB()  {}

// snmpoid
func (s *SNMPService) GetAllSNMPOId() {}
func (s *SNMPService) CountSNMPOId()  {}

// snmpoption
func (s *SNMPService) UpdateSNMPOption() {}
func (s *SNMPService) UnsetSNMPOption()  {}
func (s *SNMPService) GetAllSNMPOption() {}

// snmptrap
func (s *SNMPService) AddSNMPTrap()    {}
func (s *SNMPService) DeleteSNMPTrap() {}
func (s *SNMPService) UpdateSNMPTrap() {}
func (s *SNMPService) UnsetSNMPTrap()  {}
func (s *SNMPService) GetAllSNMPTrap() {}
func (s *SNMPService) CountSNMPTrap()  {}

// snmptrap_binding
func (s *SNMPService) GetAllSNMPTrapBinding() {}
func (s *SNMPService) GetSNMPTrapBinding()    {}

// snmptrap_snmpuser_binding
func (s *SNMPService) AddSNMPTrapSNMPUserBinding()    {}
func (s *SNMPService) DeleteSNMPTrapSNMPUserBinding() {}
func (s *SNMPService) GetAllSNMPTrapSNMPUserBinding() {}
func (s *SNMPService) GetSNMPTrapSNMPUserBinding()    {}
func (s *SNMPService) CountSNMPTrapSNMPUserBinding()  {}

// snmpuser
func (s *SNMPService) AddSNMPUser()    {}
func (s *SNMPService) DeleteSNMPUser() {}
func (s *SNMPService) UpdateSNMPUser() {}
func (s *SNMPService) UnsetSNMPUser()  {}
func (s *SNMPService) GetAllSNMPUser() {}
func (s *SNMPService) GetSNMPUser()    {}
func (s *SNMPService) CountSNMPUser()  {}

// snmpview
func (s *SNMPService) AddSNMPView()    {}
func (s *SNMPService) DeleteSNMPView() {}
func (s *SNMPService) UpdateSNMPView() {}
func (s *SNMPService) GetAllSNMPView() {}
func (s *SNMPService) CountSNMPView()  {}
