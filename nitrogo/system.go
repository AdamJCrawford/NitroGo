package nitrogo

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AdamJCrawford/NitroGo/nitrogo/models"
)

// System
// System configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/system/system
type SystemService struct {
	client *Client
}

// systemautorestorefeature
func (s *SystemService) EnableSystemAutoRestoreFeature()  {}
func (s *SystemService) DisableSystemAutoRestoreFeature() {}

// systembackup
func (s *SystemService) AddSystemBackup()     {}
func (s *SystemService) DeleteSystemBackup()  {}
func (s *SystemService) CreateSystemBackup()  {}
func (s *SystemService) RestoreSystemBackup() {}
func (s *SystemService) GetAllSystemBackup()  {}
func (s *SystemService) GetSystemBackup()     {}
func (s *SystemService) CountSystemBackup()   {}

// systemcmdpolicy
func (s *SystemService) AddSystemCmdPolicy()    {}
func (s *SystemService) DeleteSystemCmdPolicy() {}
func (s *SystemService) UpdateSystemCmdPolicy() {}
func (s *SystemService) GetAllSystemCmdPolicy() {}
func (s *SystemService) GetSystemCmdPolicy()    {}
func (s *SystemService) CountSystemCmdPolicy()  {}

// systemcollectionparam
func (s *SystemService) UpdateSystemCollectionParam() {}
func (s *SystemService) UnsetSystemCollectionParam()  {}
func (s *SystemService) GetAllSystemCollectionParam() {}

// systemcore
func (s *SystemService) GetAllSystemCore() {}

// systemcountergroup
func (s *SystemService) GetAllSystemCounterGroup() {}

// systemcounters
func (s *SystemService) GetAllSystemCounters() {}

// systemdatasource
func (s *SystemService) GetAllSystemDataSource() {}

// systementity
func (s *SystemService) GetAllSystemEntity() {}

// systementitydata
func (s *SystemService) DeleteSystemEntityData() {}
func (s *SystemService) GetAllSystemEntityData() {}

// systementitytype
func (s *SystemService) GetAllSystemEntityType() {}

// systemeventhistory
func (s *SystemService) GetAllSystemEventHistory() {}

// systemextramgmtcpu
func (s *SystemService) EnableSystemExtraMgmtCPU()  {}
func (s *SystemService) DisableSystemExtraMgmtCPU() {}
func (s *SystemService) GetAllSystemExtraMgmtCPU() (models.SystemExtraMgmtCPU, error) {
	u := "nitro/v1/config/systemextramgmtcpu"
	v := models.SystemExtraMgmtCPU{}

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return v, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return v, err
	}

	data, err := json.Marshal(resp["systemextramgmtcpu"])
	if err != nil {
		return v, err
	}

	err = json.Unmarshal(data, &v)
	if err != nil {
		return v, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return v, nil
}

// systemfile
func (s *SystemService) AddSystemFile()    {}
func (s *SystemService) DeleteSystemFile() {}
func (s *SystemService) GetAllSystemFile() {}
func (s *SystemService) CountSystemFile()  {}

// systemglobaldata
func (s *SystemService) GetAllSystemGlobalData() {}

// systemglobal_auditglobal_auditnslogpolicy_binding
func (s *SystemService) AddSystemGlobalAuditNSLogPolicyBinding()    {}
func (s *SystemService) DeleteSystemGlobalAuditNSLogPolicyBinding() {}
func (s *SystemService) GetSystemGlobalAuditNSLogPolicyBinding()    {}
func (s *SystemService) CountSystemGlobalAuditNSLogPolicyBinding()  {}

// systemglobal_auditsyslogpolicy_binding
func (s *SystemService) AddSystemGlobalAuditSyslogPolicyBinding()    {}
func (s *SystemService) DeleteSystemGlobalAuditSyslogPolicyBinding() {}
func (s *SystemService) GetSystemGlobalAuditSyslogPolicyBinding()    {}
func (s *SystemService) CountSystemGlobalAuditSyslogPolicyBinding()  {}

// systemglobal_authenticationldappolicy_binding
func (s *SystemService) AddSystemGlobalAuthenticationLDAPPolicyBinding()    {}
func (s *SystemService) DeleteSystemGlobalAuthenticationLDAPPolicyBinding() {}
func (s *SystemService) GetSystemGlobalAuthenticationLDAPPolicyBinding()    {}
func (s *SystemService) CountSystemGlobalAuthenticationLDAPPolicyBinding()  {}

// systemglobal_authenticationlocalpolicy_binding
func (s *SystemService) AddSystemGlobalAuthenticationLocalPolicyBinding()    {}
func (s *SystemService) DeleteSystemGlobalAuthenticationLocalPolicyBinding() {}
func (s *SystemService) GetSystemGlobalAuthenticationLocalPolicyBinding()    {}
func (s *SystemService) CountSystemGlobalAuthenticationLocalPolicyBinding()  {}

// systemglobal_authenticationpolicy_binding
func (s *SystemService) AddSystemGlobalAuthenticationPolicyBinding()    {}
func (s *SystemService) DeleteSystemGlobalAuthenticationPolicyBinding() {}
func (s *SystemService) GetSystemGlobalAuthenticationPolicyBinding()    {}
func (s *SystemService) CountSystemGlobalAuthenticationPolicyBinding()  {}

// systemglobal_authenticationradiuspolicy_binding
func (s *SystemService) AddSystemGlobalAuthenticationRADIUSPolicyBinding()    {}
func (s *SystemService) DeleteSystemGlobalAuthenticationRADIUSPolicyBinding() {}
func (s *SystemService) GetSystemGlobalAuthenticationRADIUSPolicyBinding()    {}
func (s *SystemService) CountSystemGlobalAuthenticationRADIUSPolicyBinding()  {}

// systemglobal_authenticationtacacspolicy_binding
func (s *SystemService) AddSystemGlobalAuthenticationTACACSPolicyBinding()    {}
func (s *SystemService) DeleteSystemGlobalAuthenticationTACACSPolicyBinding() {}
func (s *SystemService) GetSystemGlobalAuthenticationTACACSPolicyBinding()    {}
func (s *SystemService) CountSystemGlobalAuthenticationTACACSPolicyBinding()  {}

// systemglobal_binding
func (s *SystemService) GetSystemGlobalBinding() {}

// systemgroup
func (s *SystemService) AddSystemGroup()    {}
func (s *SystemService) DeleteSystemGroup() {}
func (s *SystemService) UpdateSystemGroup() {}
func (s *SystemService) UnsetSystemGroup()  {}
func (s *SystemService) GetAllSystemGroup() {}
func (s *SystemService) GetSystemGroup()    {}
func (s *SystemService) CountSystemGroup()  {}

// systemgroup_binding
func (s *SystemService) GetAllSystemGroupBinding() {}
func (s *SystemService) GetSystemGroupBinding()    {}

// systemgroup_nspartition_binding
func (s *SystemService) AddSystemGroupNSPartitionBinding()    {}
func (s *SystemService) DeleteSystemGroupNSPartitionBinding() {}
func (s *SystemService) GetAllSystemGroupNSPartitionBinding() {}
func (s *SystemService) GetSystemGroupNSPartitionBinding()    {}
func (s *SystemService) CountSystemGroupNSPartitionBinding()  {}

// systemgroup_systemcmdpolicy_binding
func (s *SystemService) AddSystemGroupSystemCmdPolicyBinding()    {}
func (s *SystemService) DeleteSystemGroupSystemCmdPolicyBinding() {}
func (s *SystemService) GetAllSystemGroupSystemCmdPolicyBinding() {}
func (s *SystemService) GetSystemGroupSystemCmdPolicyBinding()    {}
func (s *SystemService) CountSystemGroupSystemCmdPolicyBinding()  {}

// systemgroup_systemuser_binding
func (s *SystemService) AddSystemGroupSystemUserBinding()    {}
func (s *SystemService) DeleteSystemGroupSystemUserBinding() {}
func (s *SystemService) GetAllSystemGroupSystemUserBinding() {}
func (s *SystemService) GetSystemGroupSystemUserBinding()    {}
func (s *SystemService) CountSystemGroupSystemUserBinding()  {}

// systemhwerror
func (s *SystemService) CheckSystemHWError() {}

// systemkek
func (s *SystemService) ChangeSystemKEK() {}

// systemparameter
func (s *SystemService) UpdateSystemParameter() {}
func (s *SystemService) UnsetSystemParameter()  {}
func (s *SystemService) GetAllSystemParameter() {}

// systemrestorepoint
func (s *SystemService) CreateSystemRestorePoint() {}
func (s *SystemService) DeleteSystemRestorePoint() {}
func (s *SystemService) GetAllSystemRestorePoint() {}
func (s *SystemService) GetSystemRestorePoint()    {}
func (s *SystemService) CountSystemRestorePoint()  {}

// systemsession
func (s *SystemService) KillSystemSession()   {}
func (s *SystemService) GetAllSystemSession() {}
func (s *SystemService) GetSystemSession()    {}
func (s *SystemService) CountSystemSession()  {}

// systemsshkey
func (s *SystemService) ImportSystemSSHKey() {}
func (s *SystemService) DeleteSystemSSHKey() {}
func (s *SystemService) GetAllSystemSSHKey() {}

// systemuser
func (s *SystemService) AddSystemUser()    {}
func (s *SystemService) DeleteSystemUser() {}
func (s *SystemService) UpdateSystemUser() {}
func (s *SystemService) UnsetSystemUser()  {}
func (s *SystemService) GetAllSystemUser() {}
func (s *SystemService) GetSystemUser()    {}
func (s *SystemService) CountSystemUser()  {}

// systemuser_binding
func (s *SystemService) GetAllSystemUserBinding() {}
func (s *SystemService) GetSystemUserBinding()    {}

// systemuser_nspartition_binding
func (s *SystemService) AddSystemUserNSPartitionBinding()    {}
func (s *SystemService) DeleteSystemUserNSPartitionBinding() {}
func (s *SystemService) GetAllSystemUserNSPartitionBinding() {}
func (s *SystemService) GetSystemUserNSPartitionBinding()    {}
func (s *SystemService) CountSystemUserNSPartitionBinding()  {}

// systemuser_systemcmdpolicy_binding
func (s *SystemService) AddSystemUserSystemCmdPolicyBinding()    {}
func (s *SystemService) DeleteSystemUserSystemCmdPolicyBinding() {}
func (s *SystemService) GetAllSystemUserSystemCmdPolicyBinding() {}
func (s *SystemService) GetSystemUserSystemCmdPolicyBinding()    {}
func (s *SystemService) CountSystemUserSystemCmdPolicyBinding()  {}

// systemuser_systemgroup_binding
func (s *SystemService) GetAllSystemUserSystemGroupBinding() {}
func (s *SystemService) GetSystemUserSystemGroupBinding()    {}
func (s *SystemService) CountSystemUserSystemGroupBinding()  {}

// statistics

// system
func (s *SystemService) GetAllSystemStats() (models.SystemStatus, error) {
	u := "nitro/v1/stat/system"
	v := models.SystemStatus{}

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return v, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return v, err
	}

	data, err := json.Marshal(resp["system"])
	if err != nil {
		return v, err
	}

	err = json.Unmarshal(data, &v)
	if err != nil {
		return v, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return v, nil
}
