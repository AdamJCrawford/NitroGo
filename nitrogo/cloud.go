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
	cloudAllowedNGSTicketProfileURL = "/nitro/v1/config/cloudallowedngsticketprofile"
	cloudAutoscaleGroupURL          = "/nitro/v1/config/cloudautoscalegroup"
	cloudCredentialURL              = "/nitro/v1/config/cloudcredential"
	cloudParameterURL               = "/nitro/v1/config/cloudparameter"
	cloudParamInternalURL           = "/nitro/v1/config/cloudparaminternal"
	cloudProfileURL                 = "/nitro/v1/config/cloudprofile"
	cloudServiceURL                 = "/nitro/v1/config/cloudservice"
	cloudVServerIPURL               = "/nitro/v1/config/cloudvserverip"
)

// Citrix ADC as SD proxy Configuration and cloud discovery commands.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/cloud/cloud
type CloudService struct {
	client *Client
}

// cloudallowedngsticketprofile
// Configuration for Allowed ticket profile for NGS resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/cloud/cloudallowedngsticketprofile

func (s *CloudService) AddCloudAllowedNGSTicketProfile(profile models.CloudAllowedNGSTicketProfile) error {
	payload := map[string]any{
		"cloudallowedngsticketprofile": profile,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, cloudAllowedNGSTicketProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *CloudService) DeleteCloudAllowedNGSTicketProfile(name string) error {
	urlReq := fmt.Sprintf("%s/%s", cloudAllowedNGSTicketProfileURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, urlReq, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *CloudService) UpdateCloudAllowedNGSTicketProfile(profile models.CloudAllowedNGSTicketProfile) error {
	payload := map[string]any{
		"cloudallowedngsticketprofile": profile,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, cloudAllowedNGSTicketProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *CloudService) GetAllCloudAllowedNGSTicketProfile() ([]models.CloudAllowedNGSTicketProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, cloudAllowedNGSTicketProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Profiles []models.CloudAllowedNGSTicketProfile `json:"cloudallowedngsticketprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Profiles, nil
}

func (s *CloudService) GetCloudAllowedNGSTicketProfile(name string) ([]models.CloudAllowedNGSTicketProfile, error) {
	urlReq := fmt.Sprintf("%s/%s", cloudAllowedNGSTicketProfileURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Profiles []models.CloudAllowedNGSTicketProfile `json:"cloudallowedngsticketprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Profiles, nil
}

func (s *CloudService) CountCloudAllowedNGSTicketProfile() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, cloudAllowedNGSTicketProfileURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Profiles []struct {
			Count float64 `json:"__count"`
		} `json:"cloudallowedngsticketprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Profiles) > 0 {
		return result.Profiles[0].Count, nil
	}

	return 0, nil
}

// cloudautoscalegroup
// Configuration for Cloud Autoscale Group resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/cloud/cloudautoscalegroup

func (s *CloudService) GetAllCloudAutoscaleGroup() ([]models.CloudAutoscaleGroup, error) {
	req, err := s.client.NewRequest(http.MethodGet, cloudAutoscaleGroupURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Groups []models.CloudAutoscaleGroup `json:"cloudautoscalegroup"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Groups, nil
}

func (s *CloudService) GetCloudAutoscaleGroup(name string) ([]models.CloudAutoscaleGroup, error) {
	urlReq := fmt.Sprintf("%s/%s", cloudAutoscaleGroupURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Groups []models.CloudAutoscaleGroup `json:"cloudautoscalegroup"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Groups, nil
}

func (s *CloudService) CountCloudAutoscaleGroup() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, cloudAutoscaleGroupURL+"?count=yes", nil)
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
		} `json:"cloudautoscalegroup"`
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

// cloudcredential
// Configuration for cloud credentials resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/cloud/cloudcredential

func (s *CloudService) GetAllCloudCredential() ([]models.CloudCredential, error) {
	req, err := s.client.NewRequest(http.MethodGet, cloudCredentialURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Credentials []models.CloudCredential `json:"cloudcredential"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Credentials, nil
}

func (s *CloudService) UpdateCloudCredential(credential models.CloudCredential) error {
	payload := map[string]any{
		"cloudcredential": credential,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, cloudCredentialURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// cloudparameter
// Configuration for cloud parameter resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/cloud/cloudparameter

func (s *CloudService) UpdateCloudParameter(param models.CloudParameter) error {
	payload := map[string]any{
		"cloudparameter": param,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, cloudParameterURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *CloudService) UnsetCloudParameter(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"cloudparameter": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, cloudParameterURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *CloudService) GetAllCloudParameter() (models.CloudParameter, error) {
	req, err := s.client.NewRequest(http.MethodGet, cloudParameterURL, nil)
	if err != nil {
		return models.CloudParameter{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.CloudParameter{}, err
	}

	var result struct {
		Parameters []models.CloudParameter `json:"cloudparameter"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.CloudParameter{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Parameters) > 0 {
		return result.Parameters[0], nil
	}

	return models.CloudParameter{}, nil
}

// cloudparaminternal
// Configuration for cloud paramInternal resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/cloud/cloudparaminternal

func (s *CloudService) GetAllCloudParamInternal() (models.CloudParamInternal, error) {
	req, err := s.client.NewRequest(http.MethodGet, cloudParamInternalURL, nil)
	if err != nil {
		return models.CloudParamInternal{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.CloudParamInternal{}, err
	}

	var result struct {
		Parameters []models.CloudParamInternal `json:"cloudparaminternal"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.CloudParamInternal{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Parameters) > 0 {
		return result.Parameters[0], nil
	}

	return models.CloudParamInternal{}, nil
}

func (s *CloudService) CountCloudParamInternal() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, cloudParamInternalURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Parameters []struct {
			Count float64 `json:"__count"`
		} `json:"cloudparaminternal"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Parameters) > 0 {
		return result.Parameters[0].Count, nil
	}

	return 0, nil
}

func (s *CloudService) UpdateCloudParamInternal(param models.CloudParamInternal) error {
	payload := map[string]any{
		"cloudparaminternal": param,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, cloudParamInternalURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// cloudprofile
// Configuration for cloud profile resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/cloud/cloudprofile

func (s *CloudService) AddCloudProfile(profile models.CloudProfile) error {
	payload := map[string]any{
		"cloudprofile": profile,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, cloudProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *CloudService) DeleteCloudProfile(name string) error {
	urlReq := fmt.Sprintf("%s/%s", cloudProfileURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, urlReq, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *CloudService) GetAllCloudProfile() ([]models.CloudProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, cloudProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Profiles []models.CloudProfile `json:"cloudprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Profiles, nil
}

func (s *CloudService) GetCloudProfile(name string) ([]models.CloudProfile, error) {
	urlReq := fmt.Sprintf("%s/%s", cloudProfileURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, urlReq, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Profiles []models.CloudProfile `json:"cloudprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Profiles, nil
}

func (s *CloudService) CountCloudProfile() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, cloudProfileURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Profiles []struct {
			Count float64 `json:"__count"`
		} `json:"cloudprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Profiles) > 0 {
		return result.Profiles[0].Count, nil
	}

	return 0, nil
}

// cloudservice
// Configuration for cloud service resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/cloud/cloudservice

func (s *CloudService) CheckCloudService(service models.CloudService) error {
	payload := map[string]any{
		"cloudservice": service,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, cloudServiceURL+"?action=check", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// cloudvserverip
// Configuration for Cloud virtual server IPs resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/cloud/cloudvserverip

func (s *CloudService) GetAllCloudVServerIP() ([]models.CloudVServerIP, error) {
	req, err := s.client.NewRequest(http.MethodGet, cloudVServerIPURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		VServerIPs []models.CloudVServerIP `json:"cloudvserverip"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.VServerIPs, nil
}

func (s *CloudService) CountCloudVServerIP() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, cloudVServerIPURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		VServerIPs []struct {
			Count float64 `json:"__count"`
		} `json:"cloudvserverip"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.VServerIPs) > 0 {
		return result.VServerIPs[0].Count, nil
	}

	return 0, nil
}
