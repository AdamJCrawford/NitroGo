package nitrogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AdamJCrawford/NitroGo/nitrogo/models"
)

const (
	dbDBProfileURL = "/nitro/v1/config/dbdbprofile"
	dbUserURL      = "/nitro/v1/config/dbuser"
)

// All the commands associated with database user
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/db/db
type DBService struct {
	client *Client
}

// dbdbprofile
// Configuration for DB profile resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/db/dbdbprofile

func (s *DBService) AddDBDBProfile(profile models.DBDBProfile) error {
	payload := map[string]any{
		"dbdbprofile": profile,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, dbDBProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *DBService) DeleteDBDBProfile(name string) error {
	url := fmt.Sprintf("%s/%s", dbDBProfileURL, name)
	req, err := s.client.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *DBService) UpdateDBDBProfile(profile models.DBDBProfile) error {
	payload := map[string]any{
		"dbdbprofile": profile,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, dbDBProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *DBService) UnsetDBDBProfile(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"dbdbprofile": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, dbDBProfileURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *DBService) GetAllDBDBProfile() ([]models.DBDBProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, dbDBProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		DBDBProfiles []models.DBDBProfile `json:"dbdbprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.DBDBProfiles, nil
}

func (s *DBService) GetDBDBProfile(name string) ([]models.DBDBProfile, error) {
	url := fmt.Sprintf("%s/%s", dbDBProfileURL, name)
	req, err := s.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		DBDBProfiles []models.DBDBProfile `json:"dbdbprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.DBDBProfiles, nil
}

func (s *DBService) CountDBDBProfile() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, dbDBProfileURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		DBDBProfiles []struct {
			Count float64 `json:"__count"`
		} `json:"dbdbprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.DBDBProfiles) > 0 {
		return result.DBDBProfiles[0].Count, nil
	}

	return 0, nil
}

// dbuser
// Configuration for DB user resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/db/dbuser

func (s *DBService) AddDBUser(user models.DBUser) error {
	payload := map[string]any{
		"dbuser": user,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, dbUserURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *DBService) DeleteDBUser(username string) error {
	url := fmt.Sprintf("%s/%s", dbUserURL, username)
	req, err := s.client.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *DBService) UpdateDBUser(user models.DBUser) error {
	payload := map[string]any{
		"dbuser": user,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, dbUserURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *DBService) GetAllDBUser() ([]models.DBUser, error) {
	req, err := s.client.NewRequest(http.MethodGet, dbUserURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		DBUsers []models.DBUser `json:"dbuser"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.DBUsers, nil
}

func (s *DBService) GetDBUser(username string) ([]models.DBUser, error) {
	url := fmt.Sprintf("%s/%s", dbUserURL, username)
	req, err := s.client.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		DBUsers []models.DBUser `json:"dbuser"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.DBUsers, nil
}

func (s *DBService) CountDBUser() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, dbUserURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		DBUsers []struct {
			Count float64 `json:"__count"`
		} `json:"dbuser"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.DBUsers) > 0 {
		return result.DBUsers[0].Count, nil
	}

	return 0, nil
}
