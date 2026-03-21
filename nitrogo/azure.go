package nitrogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AdamJCrawford/NitroGo/nitrogo/models"
)

const (
	azureApplicationURL = "/nitro/v1/config/azureapplication"
	azureKeyVaultURL    = "/nitro/v1/config/azurekeyvault"
)

// Azure configuration
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/azure/azure
type AzureService struct {
	client *Client
}

// azureapplication
// Configuration for Azure Application resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/azure/azureapplication

func (s *AzureService) AddAzureApplication(application models.AzureApplication) error {
	payload := map[string]any{
		"azureapplication": application,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, azureApplicationURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AzureService) DeleteAzureApplication(name string) error {
	url := fmt.Sprintf("%s/%s", azureApplicationURL, name)
	req, err := s.client.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AzureService) GetAllAzureApplication() ([]models.AzureApplication, error) {
	req, err := s.client.NewRequest(http.MethodGet, azureApplicationURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		AzureApplications []models.AzureApplication `json:"azureapplication"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.AzureApplications, nil
}

func (s *AzureService) GetAzureApplication(name string) ([]models.AzureApplication, error) {
	url := fmt.Sprintf("%s/%s", azureApplicationURL, name)
	req, err := s.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		AzureApplications []models.AzureApplication `json:"azureapplication"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.AzureApplications, nil
}

func (s *AzureService) CountAzureApplication() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, azureApplicationURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		AzureApplications []struct {
			Count float64 `json:"__count"`
		} `json:"azureapplication"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.AzureApplications) > 0 {
		return result.AzureApplications[0].Count, nil
	}

	return 0, nil
}

// azurekeyvault
// Configuration for Azure Key Vault entity resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/azure/azurekeyvault

func (s *AzureService) AddAzureKeyVault(keyVault models.AzureKeyVault) error {
	payload := map[string]any{
		"azurekeyvault": keyVault,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, azureKeyVaultURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AzureService) DeleteAzureKeyVault(name string) error {
	url := fmt.Sprintf("%s/%s", azureKeyVaultURL, name)
	req, err := s.client.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AzureService) GetAllAzureKeyVault() ([]models.AzureKeyVault, error) {
	req, err := s.client.NewRequest(http.MethodGet, azureKeyVaultURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		AzureKeyVaults []models.AzureKeyVault `json:"azurekeyvault"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.AzureKeyVaults, nil
}

func (s *AzureService) GetAzureKeyVault(name string) ([]models.AzureKeyVault, error) {
	url := fmt.Sprintf("%s/%s", azureKeyVaultURL, name)
	req, err := s.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		AzureKeyVaults []models.AzureKeyVault `json:"azurekeyvault"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.AzureKeyVaults, nil
}

func (s *AzureService) CountAzureKeyVault() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, azureKeyVaultURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		AzureKeyVaults []struct {
			Count float64 `json:"__count"`
		} `json:"azurekeyvault"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.AzureKeyVaults) > 0 {
		return result.AzureKeyVaults[0].Count, nil
	}

	return 0, nil
}
