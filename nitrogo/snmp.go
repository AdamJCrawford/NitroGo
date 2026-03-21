package nitrogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/AdamJCrawford/NitroGo/nitrogo/models"
)

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

func (s *SNMPService) UpdateSNMPAlarm(alarm models.SNMPAlarm) error {
	payload := map[string]any{
		"snmpalarm": alarm,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, snmpAlarmURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SNMPService) UnsetSNMPAlarm(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"snmpalarm": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, snmpAlarmURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SNMPService) EnableSNMPAlarm(trapname string) error {
	payload := map[string]any{
		"snmpalarm": map[string]string{
			"trapname": trapname,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, snmpAlarmURL+"?action=enable", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SNMPService) DisableSNMPAlarm(trapname string) error {
	payload := map[string]any{
		"snmpalarm": map[string]string{
			"trapname": trapname,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, snmpAlarmURL+"?action=disable", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SNMPService) GetAllSNMPAlarm() ([]models.SNMPAlarm, error) {
	req, err := s.client.NewRequest(http.MethodGet, snmpAlarmURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Alarms []models.SNMPAlarm `json:"snmpalarm"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Alarms, nil
}

func (s *SNMPService) GetSNMPAlarm(trapname string) ([]models.SNMPAlarm, error) {
	urlReq := fmt.Sprintf("%s/%s", snmpAlarmURL, url.QueryEscape(trapname))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Alarms []models.SNMPAlarm `json:"snmpalarm"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Alarms, nil
}

func (s *SNMPService) CountSNMPAlarm() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, snmpAlarmURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Alarms []struct {
			Count float64 `json:"__count"`
		} `json:"snmpalarm"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Alarms) > 0 {
		return result.Alarms[0].Count, nil
	}

	return 0, nil
}

// snmpcommunity

func (s *SNMPService) AddSNMPCommunity(community models.SNMPCommunity) error {
	payload := map[string]any{
		"snmpcommunity": community,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, snmpCommunityURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SNMPService) DeleteSNMPCommunity(communityname string) error {
	urlReq := fmt.Sprintf("%s/%s", snmpCommunityURL, url.QueryEscape(communityname))
	req, err := s.client.NewRequest(http.MethodDelete, urlReq, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SNMPService) GetAllSNMPCommunity() ([]models.SNMPCommunity, error) {
	req, err := s.client.NewRequest(http.MethodGet, snmpCommunityURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Communities []models.SNMPCommunity `json:"snmpcommunity"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Communities, nil
}

func (s *SNMPService) GetSNMPCommunity(communityname string) ([]models.SNMPCommunity, error) {
	urlReq := fmt.Sprintf("%s/%s", snmpCommunityURL, url.QueryEscape(communityname))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Communities []models.SNMPCommunity `json:"snmpcommunity"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Communities, nil
}

func (s *SNMPService) CountSNMPCommunity() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, snmpCommunityURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Communities []struct {
			Count float64 `json:"__count"`
		} `json:"snmpcommunity"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Communities) > 0 {
		return result.Communities[0].Count, nil
	}

	return 0, nil
}

// snmpengineid

func (s *SNMPService) UpdateSNMPEngineId(engineid models.SNMPEngineID) error {
	payload := map[string]any{
		"snmpengineid": engineid,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, snmpEngineIdURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SNMPService) UnsetSNMPEngineId(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"snmpengineid": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, snmpEngineIdURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SNMPService) GetAllSNMPEngineId() ([]models.SNMPEngineID, error) {
	req, err := s.client.NewRequest(http.MethodGet, snmpEngineIdURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		EngineIDs []models.SNMPEngineID `json:"snmpengineid"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.EngineIDs, nil
}

func (s *SNMPService) GetSNMPEngineId(ownernode string) ([]models.SNMPEngineID, error) {
	urlReq := fmt.Sprintf("%s/%s", snmpEngineIdURL, url.QueryEscape(ownernode))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		EngineIDs []models.SNMPEngineID `json:"snmpengineid"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.EngineIDs, nil
}

func (s *SNMPService) CountSNMPEngineId() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, snmpEngineIdURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		EngineIDs []struct {
			Count float64 `json:"__count"`
		} `json:"snmpengineid"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.EngineIDs) > 0 {
		return result.EngineIDs[0].Count, nil
	}

	return 0, nil
}

// snmpgroup

func (s *SNMPService) AddSNMPGroup(group models.SNMPGroup) error {
	payload := map[string]any{
		"snmpgroup": group,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, snmpGroupURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SNMPService) DeleteSNMPGroup(name string, securitylevel string) error {
	urlReq := fmt.Sprintf("%s/%s?args=securitylevel:%s", snmpGroupURL, url.QueryEscape(name), url.QueryEscape(securitylevel))
	req, err := s.client.NewRequest(http.MethodDelete, urlReq, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SNMPService) GetAllSNMPGroup() ([]models.SNMPGroup, error) {
	req, err := s.client.NewRequest(http.MethodGet, snmpGroupURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Groups []models.SNMPGroup `json:"snmpgroup"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Groups, nil
}

func (s *SNMPService) GetSNMPGroup(name string) ([]models.SNMPGroup, error) {
	urlReq := fmt.Sprintf("%s/%s", snmpGroupURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Groups []models.SNMPGroup `json:"snmpgroup"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Groups, nil
}

func (s *SNMPService) CountSNMPGroup() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, snmpGroupURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Groups []struct {
			Count float64 `json:"__count"`
		} `json:"snmpgroup"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Groups) > 0 {
		return result.Groups[0].Count, nil
	}

	return 0, nil
}

// snmpmanager

func (s *SNMPService) AddSNMPManager(manager models.SNMPManager) error {
	payload := map[string]any{
		"snmpmanager": manager,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, snmpManagerURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SNMPService) DeleteSNMPManager(ip string) error {
	urlReq := fmt.Sprintf("%s/%s", snmpManagerURL, url.QueryEscape(ip))
	req, err := s.client.NewRequest(http.MethodDelete, urlReq, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SNMPService) UpdateSNMPManager(manager models.SNMPManager) error {
	payload := map[string]any{
		"snmpmanager": manager,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, snmpManagerURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SNMPService) UnsetSNMPManager(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"snmpmanager": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, snmpManagerURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SNMPService) GetAllSNMPManager() ([]models.SNMPManager, error) {
	req, err := s.client.NewRequest(http.MethodGet, snmpManagerURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Managers []models.SNMPManager `json:"snmpmanager"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Managers, nil
}

func (s *SNMPService) CountSNMPManager() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, snmpManagerURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Managers []struct {
			Count float64 `json:"__count"`
		} `json:"snmpmanager"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Managers) > 0 {
		return result.Managers[0].Count, nil
	}

	return 0, nil
}

// snmpmib

func (s *SNMPService) UpdateSNMPMIB(mib models.SNMPMIB) error {
	payload := map[string]any{
		"snmpmib": mib,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, snmpMIBURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SNMPService) UnsetSNMPMIB(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"snmpmib": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, snmpMIBURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SNMPService) GetAllSNMPMIB() ([]models.SNMPMIB, error) {
	req, err := s.client.NewRequest(http.MethodGet, snmpMIBURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		MIBs []models.SNMPMIB `json:"snmpmib"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.MIBs, nil
}

func (s *SNMPService) GetSNMPMIB(ownernode string) ([]models.SNMPMIB, error) {
	urlReq := fmt.Sprintf("%s/%s", snmpMIBURL, url.QueryEscape(ownernode))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		MIBs []models.SNMPMIB `json:"snmpmib"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.MIBs, nil
}

func (s *SNMPService) CountSNMPMIB() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, snmpMIBURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		MIBs []struct {
			Count float64 `json:"__count"`
		} `json:"snmpmib"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.MIBs) > 0 {
		return result.MIBs[0].Count, nil
	}

	return 0, nil
}

// snmpoid

func (s *SNMPService) GetAllSNMPOId() ([]models.SNMPOID, error) {
	req, err := s.client.NewRequest(http.MethodGet, snmpOIDURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		OIDs []models.SNMPOID `json:"snmpoid"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.OIDs, nil
}

func (s *SNMPService) CountSNMPOId() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, snmpOIDURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		OIDs []struct {
			Count float64 `json:"__count"`
		} `json:"snmpoid"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.OIDs) > 0 {
		return result.OIDs[0].Count, nil
	}

	return 0, nil
}

// snmpoption

func (s *SNMPService) UpdateSNMPOption(option models.SNMPOption) error {
	payload := map[string]any{
		"snmpoption": option,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, snmpOptionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SNMPService) UnsetSNMPOption(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"snmpoption": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, snmpOptionURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SNMPService) GetAllSNMPOption() (models.SNMPOption, error) {
	req, err := s.client.NewRequest(http.MethodGet, snmpOptionURL, nil)
	if err != nil {
		return models.SNMPOption{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.SNMPOption{}, err
	}

	var result struct {
		Options []models.SNMPOption `json:"snmpoption"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.SNMPOption{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Options) > 0 {
		return result.Options[0], nil
	}

	return models.SNMPOption{}, nil
}

// snmptrap

func (s *SNMPService) AddSNMPTrap(trap models.SNMPTrap) error {
	payload := map[string]any{
		"snmptrap": trap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, snmpTrapURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SNMPService) DeleteSNMPTrap(trapclass string, trapdestination string) error {
	urlReq := fmt.Sprintf("%s/%s?args=trapdestination:%s", snmpTrapURL, url.QueryEscape(trapclass), url.QueryEscape(trapdestination))
	req, err := s.client.NewRequest(http.MethodDelete, urlReq, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SNMPService) UpdateSNMPTrap(trap models.SNMPTrap) error {
	payload := map[string]any{
		"snmptrap": trap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, snmpTrapURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SNMPService) UnsetSNMPTrap(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"snmptrap": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, snmpTrapURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SNMPService) GetAllSNMPTrap() ([]models.SNMPTrap, error) {
	req, err := s.client.NewRequest(http.MethodGet, snmpTrapURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Traps []models.SNMPTrap `json:"snmptrap"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Traps, nil
}

func (s *SNMPService) CountSNMPTrap() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, snmpTrapURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Traps []struct {
			Count float64 `json:"__count"`
		} `json:"snmptrap"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Traps) > 0 {
		return result.Traps[0].Count, nil
	}

	return 0, nil
}

// snmptrap_binding

func (s *SNMPService) GetAllSNMPTrapBinding() ([]models.SNMPTrapBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, snmpTrapBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.SNMPTrapBinding `json:"snmptrap_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *SNMPService) GetSNMPTrapBinding(trapclass string) ([]models.SNMPTrapBinding, error) {
	urlReq := fmt.Sprintf("%s/%s", snmpTrapBindingURL, url.QueryEscape(trapclass))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.SNMPTrapBinding `json:"snmptrap_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

// snmptrap_snmpuser_binding

func (s *SNMPService) AddSNMPTrapSNMPUserBinding(binding models.SNMPTrapSNMPUserBinding) error {
	payload := map[string]any{
		"snmptrap_snmpuser_binding": binding,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, snmpTrapSNMPUserBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SNMPService) DeleteSNMPTrapSNMPUserBinding(trapclass string, trapdestination string) error {
	urlReq := fmt.Sprintf("%s/%s?args=trapdestination:%s", snmpTrapSNMPUserBindingURL, url.QueryEscape(trapclass), url.QueryEscape(trapdestination))
	req, err := s.client.NewRequest(http.MethodDelete, urlReq, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SNMPService) GetAllSNMPTrapSNMPUserBinding() ([]models.SNMPTrapSNMPUserBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, snmpTrapSNMPUserBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.SNMPTrapSNMPUserBinding `json:"snmptrap_snmpuser_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *SNMPService) GetSNMPTrapSNMPUserBinding(trapclass string) ([]models.SNMPTrapSNMPUserBinding, error) {
	urlReq := fmt.Sprintf("%s/%s", snmpTrapSNMPUserBindingURL, url.QueryEscape(trapclass))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.SNMPTrapSNMPUserBinding `json:"snmptrap_snmpuser_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *SNMPService) CountSNMPTrapSNMPUserBinding(trapclass string) (float64, error) {
	urlReq := fmt.Sprintf("%s/%s?count=yes", snmpTrapSNMPUserBindingURL, url.QueryEscape(trapclass))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
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
		} `json:"snmptrap_snmpuser_binding"`
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

// snmpuser

func (s *SNMPService) AddSNMPUser(user models.SNMPUser) error {
	payload := map[string]any{
		"snmpuser": user,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, snmpUserURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SNMPService) DeleteSNMPUser(name string) error {
	urlReq := fmt.Sprintf("%s/%s", snmpUserURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, urlReq, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SNMPService) UpdateSNMPUser(user models.SNMPUser) error {
	payload := map[string]any{
		"snmpuser": user,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, snmpUserURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SNMPService) UnsetSNMPUser(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"snmpuser": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, snmpUserURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SNMPService) GetAllSNMPUser() ([]models.SNMPUser, error) {
	req, err := s.client.NewRequest(http.MethodGet, snmpUserURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Users []models.SNMPUser `json:"snmpuser"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Users, nil
}

func (s *SNMPService) GetSNMPUser(name string) ([]models.SNMPUser, error) {
	urlReq := fmt.Sprintf("%s/%s", snmpUserURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Users []models.SNMPUser `json:"snmpuser"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Users, nil
}

func (s *SNMPService) CountSNMPUser() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, snmpUserURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Users []struct {
			Count float64 `json:"__count"`
		} `json:"snmpuser"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Users) > 0 {
		return result.Users[0].Count, nil
	}

	return 0, nil
}

// snmpview

func (s *SNMPService) AddSNMPView(view models.SNMPView) error {
	payload := map[string]any{
		"snmpview": view,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, snmpViewURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SNMPService) DeleteSNMPView(name string, subtree string) error {
	urlReq := fmt.Sprintf("%s/%s?args=subtree:%s", snmpViewURL, url.QueryEscape(name), url.QueryEscape(subtree))
	req, err := s.client.NewRequest(http.MethodDelete, urlReq, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SNMPService) UpdateSNMPView(view models.SNMPView) error {
	payload := map[string]any{
		"snmpview": view,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, snmpViewURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SNMPService) GetAllSNMPView() ([]models.SNMPView, error) {
	req, err := s.client.NewRequest(http.MethodGet, snmpViewURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Views []models.SNMPView `json:"snmpview"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Views, nil
}

func (s *SNMPService) CountSNMPView() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, snmpViewURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Views []struct {
			Count float64 `json:"__count"`
		} `json:"snmpview"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Views) > 0 {
		return result.Views[0].Count, nil
	}

	return 0, nil
}
