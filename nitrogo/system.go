package nitrogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/AdamJCrawford/NitroGo/nitrogo/models"
)

const (
	systemAutoRestoreFeatureURL                      = "/nitro/v1/config/systemautorestorefeature"
	systemBackupURL                                  = "/nitro/v1/config/systembackup"
	systemCmdPolicyURL                               = "/nitro/v1/config/systemcmdpolicy"
	systemCollectionParamURL                         = "/nitro/v1/config/systemcollectionparam"
	systemCoreURL                                    = "/nitro/v1/config/systemcore"
	systemCounterGroupURL                            = "/nitro/v1/config/systemcountergroup"
	systemCountersURL                                = "/nitro/v1/config/systemcounters"
	systemDataSourceURL                              = "/nitro/v1/config/systemdatasource"
	systemEntityURL                                  = "/nitro/v1/config/systementity"
	systemEntityDataURL                              = "/nitro/v1/config/systementitydata"
	systemEntityTypeURL                              = "/nitro/v1/config/systementitytype"
	systemEventHistoryURL                            = "/nitro/v1/config/systemeventhistory"
	systemExtraMgmtCPUURL                            = "/nitro/v1/config/systemextramgmtcpu"
	systemFileURL                                    = "/nitro/v1/config/systemfile"
	systemGlobalDataURL                              = "/nitro/v1/config/systemglobaldata"
	systemGlobalAuditNSLogPolicyBindingURL           = "/nitro/v1/config/systemglobal_auditnslogpolicy_binding"
	systemGlobalAuditSyslogPolicyBindingURL          = "/nitro/v1/config/systemglobal_auditsyslogpolicy_binding"
	systemGlobalAuthenticationLDAPPolicyBindingURL   = "/nitro/v1/config/systemglobal_authenticationldappolicy_binding"
	systemGlobalAuthenticationLocalPolicyBindingURL  = "/nitro/v1/config/systemglobal_authenticationlocalpolicy_binding"
	systemGlobalAuthenticationPolicyBindingURL       = "/nitro/v1/config/systemglobal_authenticationpolicy_binding"
	systemGlobalAuthenticationRADIUSPolicyBindingURL = "/nitro/v1/config/systemglobal_authenticationradiuspolicy_binding"
	systemGlobalAuthenticationTACACSPolicyBindingURL = "/nitro/v1/config/systemglobal_authenticationtacacspolicy_binding"
	systemGlobalBindingURL                           = "/nitro/v1/config/systemglobal_binding"
	systemGroupURL                                   = "/nitro/v1/config/systemgroup"
	systemGroupBindingURL                            = "/nitro/v1/config/systemgroup_binding"
	systemGroupNSPartitionBindingURL                 = "/nitro/v1/config/systemgroup_nspartition_binding"
	systemGroupSystemCmdPolicyBindingURL             = "/nitro/v1/config/systemgroup_systemcmdpolicy_binding"
	systemGroupSystemUserBindingURL                  = "/nitro/v1/config/systemgroup_systemuser_binding"
	systemHWErrorURL                                 = "/nitro/v1/config/systemhwerror"
	systemKEKURL                                     = "/nitro/v1/config/systemkek"
	systemParameterURL                               = "/nitro/v1/config/systemparameter"
	systemRestorePointURL                            = "/nitro/v1/config/systemrestorepoint"
	systemSessionURL                                 = "/nitro/v1/config/systemsession"
	systemSSHKeyURL                                  = "/nitro/v1/config/systemsshkey"
	systemUserURL                                    = "/nitro/v1/config/systemuser"
	systemUserBindingURL                             = "/nitro/v1/config/systemuser_binding"
	systemUserNSPartitionBindingURL                  = "/nitro/v1/config/systemuser_nspartition_binding"
	systemUserSystemCmdPolicyBindingURL              = "/nitro/v1/config/systemuser_systemcmdpolicy_binding"
	systemUserSystemGroupBindingURL                  = "/nitro/v1/config/systemuser_systemgroup_binding"
	systemStatsURL                                   = "/nitro/v1/stat/system"
)

// System
// System configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/system/system
type SystemService struct {
	client *Client
}

// systemautorestorefeature
func (s *SystemService) EnableSystemAutoRestoreFeature() error {
	payload := map[string]any{"systemautorestorefeature": map[string]any{}}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=enable", systemAutoRestoreFeatureURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) DisableSystemAutoRestoreFeature() error {
	payload := map[string]any{"systemautorestorefeature": map[string]any{}}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=disable", systemAutoRestoreFeatureURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// systembackup
func (s *SystemService) AddSystemBackup(resource models.SystemBackup) error {
	payload := map[string]any{"systembackup": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, systemBackupURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) DeleteSystemBackup(filename string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", systemBackupURL, filename), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) CreateSystemBackup(resource models.SystemBackup) error {
	payload := map[string]any{"systembackup": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=create", systemBackupURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) RestoreSystemBackup(resource models.SystemBackup) error {
	payload := map[string]any{"systembackup": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=restore", systemBackupURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) GetAllSystemBackup() ([]models.SystemBackup, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemBackupURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		SystemBackup []models.SystemBackup `json:"systembackup"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.SystemBackup, nil
}

func (s *SystemService) GetSystemBackup(filename string) (models.SystemBackup, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", systemBackupURL, filename), nil)
	if err != nil {
		return models.SystemBackup{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.SystemBackup{}, err
	}

	var result struct {
		SystemBackup []models.SystemBackup `json:"systembackup"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.SystemBackup{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.SystemBackup) == 0 {
		return models.SystemBackup{}, fmt.Errorf("systembackup %s not found", filename)
	}

	return result.SystemBackup[0], nil
}

func (s *SystemService) CountSystemBackup() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", systemBackupURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		SystemBackup []models.SystemBackup `json:"systembackup"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.SystemBackup) > 0 {
		return int(result.SystemBackup[0].Count), nil
	}
	return 0, fmt.Errorf("failed to parse count")
}

// systemcmdpolicy
func (s *SystemService) AddSystemCmdPolicy(resource models.SystemCmdPolicy) error {
	payload := map[string]any{"systemcmdpolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, systemCmdPolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) DeleteSystemCmdPolicy(policyname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", systemCmdPolicyURL, policyname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) UpdateSystemCmdPolicy(resource models.SystemCmdPolicy) error {
	payload := map[string]any{"systemcmdpolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, systemCmdPolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) GetAllSystemCmdPolicy() ([]models.SystemCmdPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemCmdPolicyURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		SystemCmdPolicy []models.SystemCmdPolicy `json:"systemcmdpolicy"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.SystemCmdPolicy, nil
}

func (s *SystemService) GetSystemCmdPolicy(policyname string) (models.SystemCmdPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", systemCmdPolicyURL, policyname), nil)
	if err != nil {
		return models.SystemCmdPolicy{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.SystemCmdPolicy{}, err
	}

	var result struct {
		SystemCmdPolicy []models.SystemCmdPolicy `json:"systemcmdpolicy"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.SystemCmdPolicy{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.SystemCmdPolicy) == 0 {
		return models.SystemCmdPolicy{}, fmt.Errorf("systemcmdpolicy %s not found", policyname)
	}

	return result.SystemCmdPolicy[0], nil
}

func (s *SystemService) CountSystemCmdPolicy() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", systemCmdPolicyURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		SystemCmdPolicy []models.SystemCmdPolicy `json:"systemcmdpolicy"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.SystemCmdPolicy) > 0 {
		return int(result.SystemCmdPolicy[0].Count), nil
	}
	return 0, fmt.Errorf("failed to parse count")
}

// systemcollectionparam
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/system/systemcollectionparam
func (s *SystemService) UpdateSystemCollectionParam(resource models.SystemCollectionParam) error {
	payload := map[string]any{"systemcollectionparam": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, systemCollectionParamURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) UnsetSystemCollectionParam(resource models.SystemCollectionParam) error {
	payload := map[string]any{"systemcollectionparam": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", systemCollectionParamURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) GetAllSystemCollectionParam() (models.SystemCollectionParam, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemCollectionParamURL, nil)
	if err != nil {
		return models.SystemCollectionParam{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.SystemCollectionParam{}, err
	}

	var result struct {
		Data models.SystemCollectionParam `json:"systemcollectionparam"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.SystemCollectionParam{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// systemcore
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/system/systemcore
func (s *SystemService) GetAllSystemCore(args map[string]string) ([]models.SystemCore, error) {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", systemCoreURL, argsStr), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.SystemCore `json:"systemcore"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// systemcountergroup
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/system/systemcountergroup
func (s *SystemService) GetAllSystemCounterGroup(args map[string]string) ([]models.SystemCounterGroup, error) {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", systemCounterGroupURL, argsStr), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.SystemCounterGroup `json:"systemcountergroup"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// systemcounters
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/system/systemcounters
func (s *SystemService) GetAllSystemCounters(args map[string]string) ([]models.SystemCounters, error) {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", systemCountersURL, argsStr), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.SystemCounters `json:"systemcounters"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// systemdatasource
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/system/systemdatasource
func (s *SystemService) GetAllSystemDataSource(args map[string]string) ([]models.SystemDataSource, error) {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", systemDataSourceURL, argsStr), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.SystemDataSource `json:"systemdatasource"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// systementity
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/system/systementity
func (s *SystemService) GetAllSystemEntity(args map[string]string) ([]models.SystemEntity, error) {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", systemEntityURL, argsStr), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.SystemEntity `json:"systementity"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// systementitydata
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/system/systementitydata
func (s *SystemService) DeleteSystemEntityData(args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s%s", systemEntityDataURL, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) GetAllSystemEntityData(args map[string]string) ([]models.SystemEntityData, error) {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", systemEntityDataURL, argsStr), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.SystemEntityData `json:"systementitydata"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// systementitytype
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/system/systementitytype
func (s *SystemService) GetAllSystemEntityType(args map[string]string) ([]models.SystemEntityType, error) {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", systemEntityTypeURL, argsStr), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.SystemEntityType `json:"systementitytype"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// systemeventhistory
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/system/systemeventhistory
func (s *SystemService) GetAllSystemEventHistory(args map[string]string) ([]models.SystemEventHistory, error) {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", systemEventHistoryURL, argsStr), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.SystemEventHistory `json:"systemeventhistory"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// systemextramgmtcpu
func (s *SystemService) EnableSystemExtraMgmtCPU() error {
	payload := map[string]any{"systemextramgmtcpu": map[string]any{}}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=enable", systemExtraMgmtCPUURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) DisableSystemExtraMgmtCPU() error {
	payload := map[string]any{"systemextramgmtcpu": map[string]any{}}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=disable", systemExtraMgmtCPUURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) GetAllSystemExtraMgmtCPU() (models.SystemExtraMgmtCPU, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemExtraMgmtCPUURL, nil)
	if err != nil {
		return models.SystemExtraMgmtCPU{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.SystemExtraMgmtCPU{}, err
	}

	var result struct {
		ExtraMgmtCPU models.SystemExtraMgmtCPU `json:"systemextramgmtcpu"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.SystemExtraMgmtCPU{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.ExtraMgmtCPU, nil
}

// systemfile
func (s *SystemService) AddSystemFile(resource models.SystemFile) error {
	payload := map[string]any{"systemfile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, systemFileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) DeleteSystemFile(filename string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", systemFileURL, filename, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) GetAllSystemFile() ([]models.SystemFile, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemFileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		SystemFile []models.SystemFile `json:"systemfile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.SystemFile, nil
}

func (s *SystemService) CountSystemFile() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", systemFileURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		SystemFile []models.SystemFile `json:"systemfile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.SystemFile) > 0 {
		return int(result.SystemFile[0].Count), nil
	}
	return 0, fmt.Errorf("failed to parse count")
}

// systemglobaldata
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/system/systemglobaldata
func (s *SystemService) GetAllSystemGlobalData(args map[string]string) ([]models.SystemGlobalData, error) {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", systemGlobalDataURL, argsStr), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.SystemGlobalData `json:"systemglobaldata"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// systemglobal_auditnslogpolicy_binding
func (s *SystemService) AddSystemGlobalAuditNSLogPolicyBinding(resource models.SystemGlobalAuditNSLogPolicyBinding) error {
	payload := map[string]any{"systemglobal_auditnslogpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, systemGlobalAuditNSLogPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) DeleteSystemGlobalAuditNSLogPolicyBinding(policyname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", systemGlobalAuditNSLogPolicyBindingURL, policyname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) GetSystemGlobalAuditNSLogPolicyBinding() ([]models.SystemGlobalAuditNSLogPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemGlobalAuditNSLogPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		SystemGlobalAuditNSLogPolicyBinding []models.SystemGlobalAuditNSLogPolicyBinding `json:"systemglobal_auditnslogpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.SystemGlobalAuditNSLogPolicyBinding, nil
}

func (s *SystemService) CountSystemGlobalAuditNSLogPolicyBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemGlobalAuditNSLogPolicyBindingURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"systemglobal_auditnslogpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// systemglobal_auditsyslogpolicy_binding
func (s *SystemService) AddSystemGlobalAuditSyslogPolicyBinding(resource models.SystemGlobalAuditSyslogPolicyBinding) error {
	payload := map[string]any{"systemglobal_auditsyslogpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, systemGlobalAuditSyslogPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) DeleteSystemGlobalAuditSyslogPolicyBinding(policyname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", systemGlobalAuditSyslogPolicyBindingURL, policyname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) GetSystemGlobalAuditSyslogPolicyBinding() ([]models.SystemGlobalAuditSyslogPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemGlobalAuditSyslogPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		SystemGlobalAuditSyslogPolicyBinding []models.SystemGlobalAuditSyslogPolicyBinding `json:"systemglobal_auditsyslogpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.SystemGlobalAuditSyslogPolicyBinding, nil
}

// Struct lacks Count field, leaving for later
func (s *SystemService) CountSystemGlobalAuditSyslogPolicyBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemGlobalAuditSyslogPolicyBindingURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"systemglobal_auditsyslogpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// systemglobal_authenticationldappolicy_binding
func (s *SystemService) AddSystemGlobalAuthenticationLDAPPolicyBinding(resource models.SystemGlobalAuthenticationLDAPPolicyBinding) error {
	payload := map[string]any{"systemglobal_authenticationldappolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, systemGlobalAuthenticationLDAPPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) DeleteSystemGlobalAuthenticationLDAPPolicyBinding(policyname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", systemGlobalAuthenticationLDAPPolicyBindingURL, policyname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) GetSystemGlobalAuthenticationLDAPPolicyBinding() ([]models.SystemGlobalAuthenticationLDAPPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemGlobalAuthenticationLDAPPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		SystemGlobalAuthenticationLDAPPolicyBinding []models.SystemGlobalAuthenticationLDAPPolicyBinding `json:"systemglobal_authenticationldappolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.SystemGlobalAuthenticationLDAPPolicyBinding, nil
}

// Struct lacks Count field, leaving for later
func (s *SystemService) CountSystemGlobalAuthenticationLDAPPolicyBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemGlobalAuthenticationLDAPPolicyBindingURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"systemglobal_authenticationldappolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// systemglobal_authenticationlocalpolicy_binding
func (s *SystemService) AddSystemGlobalAuthenticationLocalPolicyBinding(resource models.SystemGlobalAuthenticationLocalPolicyBinding) error {
	payload := map[string]any{"systemglobal_authenticationlocalpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, systemGlobalAuthenticationLocalPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) DeleteSystemGlobalAuthenticationLocalPolicyBinding(policyname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", systemGlobalAuthenticationLocalPolicyBindingURL, policyname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) GetSystemGlobalAuthenticationLocalPolicyBinding() ([]models.SystemGlobalAuthenticationLocalPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemGlobalAuthenticationLocalPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		SystemGlobalAuthenticationLocalPolicyBinding []models.SystemGlobalAuthenticationLocalPolicyBinding `json:"systemglobal_authenticationlocalpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.SystemGlobalAuthenticationLocalPolicyBinding, nil
}

// Struct lacks Count field, leaving for later
func (s *SystemService) CountSystemGlobalAuthenticationLocalPolicyBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemGlobalAuthenticationLocalPolicyBindingURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"systemglobal_authenticationlocalpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// systemglobal_authenticationpolicy_binding
func (s *SystemService) AddSystemGlobalAuthenticationPolicyBinding(resource models.SystemGlobalAuthenticationPolicyBinding) error {
	payload := map[string]any{"systemglobal_authenticationpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, systemGlobalAuthenticationPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) DeleteSystemGlobalAuthenticationPolicyBinding(policyname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", systemGlobalAuthenticationPolicyBindingURL, policyname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) GetSystemGlobalAuthenticationPolicyBinding() ([]models.SystemGlobalAuthenticationPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemGlobalAuthenticationPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		SystemGlobalAuthenticationPolicyBinding []models.SystemGlobalAuthenticationPolicyBinding `json:"systemglobal_authenticationpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.SystemGlobalAuthenticationPolicyBinding, nil
}

// Struct lacks Count field, leaving for later
func (s *SystemService) CountSystemGlobalAuthenticationPolicyBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemGlobalAuthenticationPolicyBindingURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"systemglobal_authenticationpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// systemglobal_authenticationradiuspolicy_binding
func (s *SystemService) AddSystemGlobalAuthenticationRADIUSPolicyBinding(resource models.SystemGlobalAuthenticationRADIUSPolicyBinding) error {
	payload := map[string]any{"systemglobal_authenticationradiuspolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, systemGlobalAuthenticationRADIUSPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) DeleteSystemGlobalAuthenticationRADIUSPolicyBinding(policyname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", systemGlobalAuthenticationRADIUSPolicyBindingURL, policyname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) GetSystemGlobalAuthenticationRADIUSPolicyBinding() ([]models.SystemGlobalAuthenticationRADIUSPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemGlobalAuthenticationRADIUSPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		SystemGlobalAuthenticationRADIUSPolicyBinding []models.SystemGlobalAuthenticationRADIUSPolicyBinding `json:"systemglobal_authenticationradiuspolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.SystemGlobalAuthenticationRADIUSPolicyBinding, nil
}

// Struct lacks Count field, leaving for later
func (s *SystemService) CountSystemGlobalAuthenticationRADIUSPolicyBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemGlobalAuthenticationRADIUSPolicyBindingURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"systemglobal_authenticationradiuspolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// systemglobal_authenticationtacacspolicy_binding
func (s *SystemService) AddSystemGlobalAuthenticationTACACSPolicyBinding(resource models.SystemGlobalAuthenticationTACACSPolicyBinding) error {
	payload := map[string]any{"systemglobal_authenticationtacacspolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, systemGlobalAuthenticationTACACSPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) DeleteSystemGlobalAuthenticationTACACSPolicyBinding(policyname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", systemGlobalAuthenticationTACACSPolicyBindingURL, policyname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) GetSystemGlobalAuthenticationTACACSPolicyBinding() ([]models.SystemGlobalAuthenticationTACACSPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemGlobalAuthenticationTACACSPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		SystemGlobalAuthenticationTACACSPolicyBinding []models.SystemGlobalAuthenticationTACACSPolicyBinding `json:"systemglobal_authenticationtacacspolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.SystemGlobalAuthenticationTACACSPolicyBinding, nil
}

// Struct lacks Count field, leaving for later
func (s *SystemService) CountSystemGlobalAuthenticationTACACSPolicyBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemGlobalAuthenticationTACACSPolicyBindingURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"systemglobal_authenticationtacacspolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// systemglobal_binding
func (s *SystemService) GetSystemGlobalBinding() (models.SystemGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemGlobalBindingURL, nil)
	if err != nil {
		return models.SystemGlobalBinding{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.SystemGlobalBinding{}, err
	}

	var result struct {
		SystemGlobalBinding []models.SystemGlobalBinding `json:"systemglobal_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.SystemGlobalBinding{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.SystemGlobalBinding) == 0 {
		return models.SystemGlobalBinding{}, fmt.Errorf("systemglobal_binding not found")
	}

	return result.SystemGlobalBinding[0], nil
}

// systemgroup
func (s *SystemService) AddSystemGroup(resource models.SystemGroup) error {
	payload := map[string]any{"systemgroup": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, systemGroupURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) DeleteSystemGroup(groupname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", systemGroupURL, groupname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) UpdateSystemGroup(resource models.SystemGroup) error {
	payload := map[string]any{"systemgroup": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, systemGroupURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) UnsetSystemGroup(resource models.SystemGroup) error {
	payload := map[string]any{"systemgroup": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", systemGroupURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) GetAllSystemGroup() ([]models.SystemGroup, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemGroupURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		SystemGroup []models.SystemGroup `json:"systemgroup"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.SystemGroup, nil
}

func (s *SystemService) GetSystemGroup(groupname string) (models.SystemGroup, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", systemGroupURL, groupname), nil)
	if err != nil {
		return models.SystemGroup{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.SystemGroup{}, err
	}

	var result struct {
		SystemGroup []models.SystemGroup `json:"systemgroup"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.SystemGroup{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.SystemGroup) == 0 {
		return models.SystemGroup{}, fmt.Errorf("systemgroup %s not found", groupname)
	}

	return result.SystemGroup[0], nil
}

func (s *SystemService) CountSystemGroup() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", systemGroupURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		SystemGroup []models.SystemGroup `json:"systemgroup"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.SystemGroup) > 0 {
		return int(result.SystemGroup[0].Count), nil
	}
	return 0, fmt.Errorf("failed to parse count")
}

// systemgroup_binding
func (s *SystemService) GetAllSystemGroupBinding() ([]models.SystemGroupBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemGroupBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		SystemGroupBinding []models.SystemGroupBinding `json:"systemgroup_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.SystemGroupBinding, nil
}

func (s *SystemService) GetSystemGroupBinding(groupname string) (models.SystemGroupBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", systemGroupBindingURL, groupname), nil)
	if err != nil {
		return models.SystemGroupBinding{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.SystemGroupBinding{}, err
	}

	var result struct {
		SystemGroupBinding []models.SystemGroupBinding `json:"systemgroup_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.SystemGroupBinding{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.SystemGroupBinding) == 0 {
		return models.SystemGroupBinding{}, fmt.Errorf("systemgroup_binding %s not found", groupname)
	}

	return result.SystemGroupBinding[0], nil
}

// systemgroup_nspartition_binding
func (s *SystemService) AddSystemGroupNSPartitionBinding(resource models.SystemGroupNSPartitionBinding) error {
	payload := map[string]any{"systemgroup_nspartition_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, systemGroupNSPartitionBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) DeleteSystemGroupNSPartitionBinding(groupname string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", systemGroupNSPartitionBindingURL, groupname, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) GetAllSystemGroupNSPartitionBinding() ([]models.SystemGroupNSPartitionBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemGroupNSPartitionBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		SystemGroupNSPartitionBinding []models.SystemGroupNSPartitionBinding `json:"systemgroup_nspartition_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.SystemGroupNSPartitionBinding, nil
}

func (s *SystemService) GetSystemGroupNSPartitionBinding(groupname string) ([]models.SystemGroupNSPartitionBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", systemGroupNSPartitionBindingURL, groupname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		SystemGroupNSPartitionBinding []models.SystemGroupNSPartitionBinding `json:"systemgroup_nspartition_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.SystemGroupNSPartitionBinding, nil
}

// Struct lacks Count field, leaving for later
func (s *SystemService) CountSystemGroupNSPartitionBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemGroupNSPartitionBindingURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"systemgroup_nspartition_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// systemgroup_systemcmdpolicy_binding
func (s *SystemService) AddSystemGroupSystemCmdPolicyBinding(resource models.SystemGroupSystemCmdPolicyBinding) error {
	payload := map[string]any{"systemgroup_systemcmdpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, systemGroupSystemCmdPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) DeleteSystemGroupSystemCmdPolicyBinding(groupname string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", systemGroupSystemCmdPolicyBindingURL, groupname, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) GetAllSystemGroupSystemCmdPolicyBinding() ([]models.SystemGroupSystemCmdPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemGroupSystemCmdPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		SystemGroupSystemCmdPolicyBinding []models.SystemGroupSystemCmdPolicyBinding `json:"systemgroup_systemcmdpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.SystemGroupSystemCmdPolicyBinding, nil
}

func (s *SystemService) GetSystemGroupSystemCmdPolicyBinding(groupname string) ([]models.SystemGroupSystemCmdPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", systemGroupSystemCmdPolicyBindingURL, groupname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		SystemGroupSystemCmdPolicyBinding []models.SystemGroupSystemCmdPolicyBinding `json:"systemgroup_systemcmdpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.SystemGroupSystemCmdPolicyBinding, nil
}

// Struct lacks Count field, leaving for later
func (s *SystemService) CountSystemGroupSystemCmdPolicyBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemGroupSystemCmdPolicyBindingURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"systemgroup_systemcmdpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// systemgroup_systemuser_binding
func (s *SystemService) AddSystemGroupSystemUserBinding(resource models.SystemGroupSystemUserBinding) error {
	payload := map[string]any{"systemgroup_systemuser_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, systemGroupSystemUserBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) DeleteSystemGroupSystemUserBinding(groupname string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", systemGroupSystemUserBindingURL, groupname, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) GetAllSystemGroupSystemUserBinding() ([]models.SystemGroupSystemUserBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemGroupSystemUserBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		SystemGroupSystemUserBinding []models.SystemGroupSystemUserBinding `json:"systemgroup_systemuser_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.SystemGroupSystemUserBinding, nil
}

func (s *SystemService) GetSystemGroupSystemUserBinding(groupname string) ([]models.SystemGroupSystemUserBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", systemGroupSystemUserBindingURL, groupname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		SystemGroupSystemUserBinding []models.SystemGroupSystemUserBinding `json:"systemgroup_systemuser_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.SystemGroupSystemUserBinding, nil
}

// Struct lacks Count field, leaving for later
func (s *SystemService) CountSystemGroupSystemUserBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemGroupSystemUserBindingURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"systemgroup_systemuser_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// systemhwerror
func (s *SystemService) CheckSystemHWError() (models.SystemHWError, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemHWErrorURL, nil)
	if err != nil {
		return models.SystemHWError{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.SystemHWError{}, err
	}

	var result struct {
		SystemHWError []models.SystemHWError `json:"systemhwerror"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.SystemHWError{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.SystemHWError) == 0 {
		return models.SystemHWError{}, fmt.Errorf("systemhwerror not found")
	}

	return result.SystemHWError[0], nil
}

// systemkek
func (s *SystemService) ChangeSystemKEK(resource models.SystemKEK) error {
	payload := map[string]any{"systemkek": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, systemKEKURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// systemparameter
func (s *SystemService) UpdateSystemParameter(resource models.SystemParameter) error {
	payload := map[string]any{"systemparameter": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, systemParameterURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) UnsetSystemParameter(resource models.SystemParameter) error {
	payload := map[string]any{"systemparameter": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", systemParameterURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) GetAllSystemParameter() (models.SystemParameter, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemParameterURL, nil)
	if err != nil {
		return models.SystemParameter{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.SystemParameter{}, err
	}

	var result struct {
		SystemParameter models.SystemParameter `json:"systemparameter"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.SystemParameter{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.SystemParameter, nil
}

// systemrestorepoint
func (s *SystemService) CreateSystemRestorePoint(resource models.SystemRestorePoint) error {
	payload := map[string]any{"systemrestorepoint": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=create", systemRestorePointURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) DeleteSystemRestorePoint(filename string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", systemRestorePointURL, filename), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) GetAllSystemRestorePoint() ([]models.SystemRestorePoint, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemRestorePointURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		SystemRestorePoint []models.SystemRestorePoint `json:"systemrestorepoint"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.SystemRestorePoint, nil
}

func (s *SystemService) GetSystemRestorePoint(filename string) (models.SystemRestorePoint, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", systemRestorePointURL, filename), nil)
	if err != nil {
		return models.SystemRestorePoint{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.SystemRestorePoint{}, err
	}

	var result struct {
		SystemRestorePoint []models.SystemRestorePoint `json:"systemrestorepoint"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.SystemRestorePoint{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.SystemRestorePoint) == 0 {
		return models.SystemRestorePoint{}, fmt.Errorf("systemrestorepoint %s not found", filename)
	}

	return result.SystemRestorePoint[0], nil
}

func (s *SystemService) CountSystemRestorePoint() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", systemRestorePointURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		SystemRestorePoint []models.SystemRestorePoint `json:"systemrestorepoint"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.SystemRestorePoint) > 0 {
		return int(result.SystemRestorePoint[0].Count), nil
	}
	return 0, fmt.Errorf("failed to parse count")
}

// systemsession
func (s *SystemService) KillSystemSession(resource models.SystemSession) error {
	payload := map[string]any{"systemsession": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=kill", systemSessionURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) GetAllSystemSession() ([]models.SystemSession, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemSessionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		SystemSession []models.SystemSession `json:"systemsession"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.SystemSession, nil
}

func (s *SystemService) GetSystemSession(sid string) (models.SystemSession, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", systemSessionURL, sid), nil)
	if err != nil {
		return models.SystemSession{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.SystemSession{}, err
	}

	var result struct {
		SystemSession []models.SystemSession `json:"systemsession"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.SystemSession{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.SystemSession) == 0 {
		return models.SystemSession{}, fmt.Errorf("systemsession %s not found", sid)
	}

	return result.SystemSession[0], nil
}

func (s *SystemService) CountSystemSession() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", systemSessionURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		SystemSession []models.SystemSession `json:"systemsession"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.SystemSession) > 0 {
		return int(result.SystemSession[0].Count), nil
	}
	return 0, fmt.Errorf("failed to parse count")
}

// systemsshkey
func (s *SystemService) ImportSystemSSHKey(resource models.SystemSSHKey) error {
	payload := map[string]any{"systemsshkey": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=import", systemSSHKeyURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) DeleteSystemSSHKey(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", systemSSHKeyURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) GetAllSystemSSHKey() ([]models.SystemSSHKey, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemSSHKeyURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		SystemSSHKey []models.SystemSSHKey `json:"systemsshkey"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.SystemSSHKey, nil
}

// systemuser
func (s *SystemService) AddSystemUser(resource models.SystemUser) error {
	payload := map[string]any{"systemuser": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, systemUserURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) DeleteSystemUser(username string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", systemUserURL, username), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) UpdateSystemUser(resource models.SystemUser) error {
	payload := map[string]any{"systemuser": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, systemUserURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) UnsetSystemUser(resource models.SystemUser) error {
	payload := map[string]any{"systemuser": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", systemUserURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) GetAllSystemUser() ([]models.SystemUser, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemUserURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		SystemUser []models.SystemUser `json:"systemuser"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.SystemUser, nil
}

func (s *SystemService) GetSystemUser(username string) (models.SystemUser, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", systemUserURL, username), nil)
	if err != nil {
		return models.SystemUser{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.SystemUser{}, err
	}

	var result struct {
		SystemUser []models.SystemUser `json:"systemuser"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.SystemUser{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.SystemUser) == 0 {
		return models.SystemUser{}, fmt.Errorf("systemuser %s not found", username)
	}

	return result.SystemUser[0], nil
}

func (s *SystemService) CountSystemUser() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", systemUserURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		SystemUser []models.SystemUser `json:"systemuser"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.SystemUser) > 0 {
		return int(result.SystemUser[0].Count), nil
	}
	return 0, fmt.Errorf("failed to parse count")
}

// systemuser_binding
func (s *SystemService) GetAllSystemUserBinding() ([]models.SystemUserBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemUserBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		SystemUserBinding []models.SystemUserBinding `json:"systemuser_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.SystemUserBinding, nil
}

func (s *SystemService) GetSystemUserBinding(username string) (models.SystemUserBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", systemUserBindingURL, username), nil)
	if err != nil {
		return models.SystemUserBinding{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.SystemUserBinding{}, err
	}

	var result struct {
		SystemUserBinding []models.SystemUserBinding `json:"systemuser_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.SystemUserBinding{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.SystemUserBinding) == 0 {
		return models.SystemUserBinding{}, fmt.Errorf("systemuser_binding %s not found", username)
	}

	return result.SystemUserBinding[0], nil
}

// systemuser_nspartition_binding
func (s *SystemService) AddSystemUserNSPartitionBinding(resource models.SystemUserNSPartitionBinding) error {
	payload := map[string]any{"systemuser_nspartition_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, systemUserNSPartitionBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) DeleteSystemUserNSPartitionBinding(username string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", systemUserNSPartitionBindingURL, username, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) GetAllSystemUserNSPartitionBinding() ([]models.SystemUserNSPartitionBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemUserNSPartitionBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		SystemUserNSPartitionBinding []models.SystemUserNSPartitionBinding `json:"systemuser_nspartition_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.SystemUserNSPartitionBinding, nil
}

func (s *SystemService) GetSystemUserNSPartitionBinding(username string) ([]models.SystemUserNSPartitionBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", systemUserNSPartitionBindingURL, username), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		SystemUserNSPartitionBinding []models.SystemUserNSPartitionBinding `json:"systemuser_nspartition_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.SystemUserNSPartitionBinding, nil
}

// Struct lacks Count field, leaving for later
func (s *SystemService) CountSystemUserNSPartitionBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemUserNSPartitionBindingURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"systemuser_nspartition_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// systemuser_systemcmdpolicy_binding
func (s *SystemService) AddSystemUserSystemCmdPolicyBinding(resource models.SystemUserSystemCmdPolicyBinding) error {
	payload := map[string]any{"systemuser_systemcmdpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, systemUserSystemCmdPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) DeleteSystemUserSystemCmdPolicyBinding(username string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", systemUserSystemCmdPolicyBindingURL, username, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SystemService) GetAllSystemUserSystemCmdPolicyBinding() ([]models.SystemUserSystemCmdPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemUserSystemCmdPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		SystemUserSystemCmdPolicyBinding []models.SystemUserSystemCmdPolicyBinding `json:"systemuser_systemcmdpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.SystemUserSystemCmdPolicyBinding, nil
}

func (s *SystemService) GetSystemUserSystemCmdPolicyBinding(username string) ([]models.SystemUserSystemCmdPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", systemUserSystemCmdPolicyBindingURL, username), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		SystemUserSystemCmdPolicyBinding []models.SystemUserSystemCmdPolicyBinding `json:"systemuser_systemcmdpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.SystemUserSystemCmdPolicyBinding, nil
}

// Struct lacks Count field, leaving for later
func (s *SystemService) CountSystemUserSystemCmdPolicyBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemUserSystemCmdPolicyBindingURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"systemuser_systemcmdpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

// systemuser_systemgroup_binding
func (s *SystemService) GetAllSystemUserSystemGroupBinding() ([]models.SystemUserSystemGroupBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemUserSystemGroupBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		SystemUserSystemGroupBinding []models.SystemUserSystemGroupBinding `json:"systemuser_systemgroup_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.SystemUserSystemGroupBinding, nil
}

func (s *SystemService) GetSystemUserSystemGroupBinding(username string) ([]models.SystemUserSystemGroupBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", systemUserSystemGroupBindingURL, username), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		SystemUserSystemGroupBinding []models.SystemUserSystemGroupBinding `json:"systemuser_systemgroup_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.SystemUserSystemGroupBinding, nil
}

// Struct lacks Count field, leaving for later
func (s *SystemService) CountSystemUserSystemGroupBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemUserSystemGroupBindingURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Bindings []struct {
			Count float64 `json:"__count"`
		} `json:"systemuser_systemgroup_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) > 0 {
		return result.Bindings[0].Count, nil
	}

	return 0, nil
}

////////////////
// statistics //
////////////////

// system
func (s *SystemService) GetAllSystemStats() (models.SystemStatus, error) {
	req, err := s.client.NewRequest(http.MethodGet, systemStatsURL, nil)
	if err != nil {
		return models.SystemStatus{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.SystemStatus{}, err
	}

	var result struct {
		SystemStatus models.SystemStatus `json:"system"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.SystemStatus{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.SystemStatus, nil
}
