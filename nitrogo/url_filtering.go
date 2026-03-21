package nitrogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AdamJCrawford/NitroGo/nitrogo/models"
)

const (
	urlFilteringCategoriesURL     = "/nitro/v1/config/urlfilteringcategories"
	urlFilteringCategorizationURL = "/nitro/v1/config/urlfilteringcategorization"
	urlFilteringCategoryGroupsURL = "/nitro/v1/config/urlfilteringcategorygroups"
	urlFilteringParameterURL      = "/nitro/v1/config/urlfilteringparameter"
)

// URL Filtering feature is used control access to webpages based on category.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/urlfiltering/urlfiltering
type URLFilteringService struct {
	client *Client
}

// urlfilteringcategories
// Configuration for Categories resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/urlfiltering/urlfilteringcategories

func (s *URLFilteringService) GetAllURLFilteringCategories() ([]models.URLFilteringCategories, error) {
	req, err := s.client.NewRequest(http.MethodGet, urlFilteringCategoriesURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Categories []models.URLFilteringCategories `json:"urlfilteringcategories"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Categories, nil
}

// urlfilteringcategorization
// Configuration for Categorization resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/urlfiltering/urlfilteringcategorization

func (s *URLFilteringService) AddURLFilteringCategorization(categorization models.URLFilteringCategorization) error {
	payload := map[string]any{
		"urlfilteringcategorization": categorization,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, urlFilteringCategorizationURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *URLFilteringService) ClearURLFilteringCategorization(categorization models.URLFilteringCategorization) error {
	payload := map[string]any{
		"urlfilteringcategorization": categorization,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, urlFilteringCategorizationURL+"?action=clear", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *URLFilteringService) GetAllURLFilteringCategorization() ([]models.URLFilteringCategorization, error) {
	req, err := s.client.NewRequest(http.MethodGet, urlFilteringCategorizationURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Categorization []models.URLFilteringCategorization `json:"urlfilteringcategorization"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Categorization, nil
}

func (s *URLFilteringService) CountURLFilteringCategorization() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, urlFilteringCategorizationURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Categorization []struct {
			Count float64 `json:"__count"`
		} `json:"urlfilteringcategorization"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Categorization) > 0 {
		return result.Categorization[0].Count, nil
	}

	return 0, nil
}

// urlfilteringcategorygroups
// Configuration for Category Groups resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/urlfiltering/urlfilteringcategorygroups

func (s *URLFilteringService) GetAllURLFilteringCategoryGroups() ([]models.URLFilteringCategoryGroups, error) {
	req, err := s.client.NewRequest(http.MethodGet, urlFilteringCategoryGroupsURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		CategoryGroups []models.URLFilteringCategoryGroups `json:"urlfilteringcategorygroups"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.CategoryGroups, nil
}

// urlfilteringparameter
// Configuration for URLFILTERING paramter resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/urlfiltering/urlfilteringparameter

func (s *URLFilteringService) UpdateURLFilteringParameter(param models.URLFilteringParameter) error {
	payload := map[string]any{
		"urlfilteringparameter": param,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, urlFilteringParameterURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *URLFilteringService) UnsetURLFilteringParameter(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"urlfilteringparameter": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, urlFilteringParameterURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *URLFilteringService) GetAllURLFilteringParameter() (models.URLFilteringParameter, error) {
	req, err := s.client.NewRequest(http.MethodGet, urlFilteringParameterURL, nil)
	if err != nil {
		return models.URLFilteringParameter{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.URLFilteringParameter{}, err
	}

	var result struct {
		URLFilteringParameters []models.URLFilteringParameter `json:"urlfilteringparameter"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.URLFilteringParameter{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.URLFilteringParameters) > 0 {
		return result.URLFilteringParameters[0], nil
	}

	return models.URLFilteringParameter{}, nil
}
