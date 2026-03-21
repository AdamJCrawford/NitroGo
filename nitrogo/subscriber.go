package nitrogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AdamJCrawford/NitroGo/nitrogo/models"
)

const (
	subscriberGxInterfaceURL     = "/nitro/v1/config/subscribergxinterface"
	subscriberParamURL           = "/nitro/v1/config/subscriberparam"
	subscriberProfileURL         = "/nitro/v1/config/subscriberprofile"
	subscriberRADIUSInterfaceURL = "/nitro/v1/config/subscriberradiusinterface"
	subscriberSessionsURL        = "/nitro/v1/config/subscribersessions"
)

// Subscriber configuration
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/subscriber/subscriber
type SubscriberService struct {
	client *Client
}

// subscribergxinterface
// Configuration for Gx interface Parameters resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/subscriber/subscribergxinterface

func (s *SubscriberService) UpdateSubscriberGxInterface(param models.SubscriberGxInterface) error {
	payload := map[string]any{
		"subscribergxinterface": param,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, subscriberGxInterfaceURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SubscriberService) UnsetSubscriberGxInterface(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"subscribergxinterface": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, subscriberGxInterfaceURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SubscriberService) GetAllSubscriberGxInterface() (models.SubscriberGxInterface, error) {
	req, err := s.client.NewRequest(http.MethodGet, subscriberGxInterfaceURL, nil)
	if err != nil {
		return models.SubscriberGxInterface{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.SubscriberGxInterface{}, err
	}

	var result struct {
		SubscriberGxInterfaces []models.SubscriberGxInterface `json:"subscribergxinterface"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.SubscriberGxInterface{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.SubscriberGxInterfaces) > 0 {
		return result.SubscriberGxInterfaces[0], nil
	}

	return models.SubscriberGxInterface{}, nil
}

// subscriberparam
// Configuration for Subscriber Params resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/subscriber/subscriberparam

func (s *SubscriberService) UpdateSubscriberParam(param models.SubscriberParam) error {
	payload := map[string]any{
		"subscriberparam": param,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, subscriberParamURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SubscriberService) UnsetSubscriberParam(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"subscriberparam": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, subscriberParamURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SubscriberService) GetAllSubscriberParam() (models.SubscriberParam, error) {
	req, err := s.client.NewRequest(http.MethodGet, subscriberParamURL, nil)
	if err != nil {
		return models.SubscriberParam{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.SubscriberParam{}, err
	}

	var result struct {
		SubscriberParams []models.SubscriberParam `json:"subscriberparam"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.SubscriberParam{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.SubscriberParams) > 0 {
		return result.SubscriberParams[0], nil
	}

	return models.SubscriberParam{}, nil
}

// subscriberprofile
// Configuration for Subscriber Profile resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/subscriber/subscriberprofile

func (s *SubscriberService) AddSubscriberProfile(profile models.SubscriberProfile) error {
	payload := map[string]any{
		"subscriberprofile": profile,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, subscriberProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SubscriberService) DeleteSubscriberProfile(ip string) error {
	url := fmt.Sprintf("%s/%s", subscriberProfileURL, ip)
	req, err := s.client.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SubscriberService) UpdateSubscriberProfile(profile models.SubscriberProfile) error {
	payload := map[string]any{
		"subscriberprofile": profile,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, subscriberProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SubscriberService) UnsetSubscriberProfile(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"subscriberprofile": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, subscriberProfileURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SubscriberService) GetAllSubscriberProfile() ([]models.SubscriberProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, subscriberProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		SubscriberProfiles []models.SubscriberProfile `json:"subscriberprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.SubscriberProfiles, nil
}

func (s *SubscriberService) CountSubscriberProfile() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, subscriberProfileURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		SubscriberProfiles []struct {
			Count float64 `json:"__count"`
		} `json:"subscriberprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.SubscriberProfiles) > 0 {
		return result.SubscriberProfiles[0].Count, nil
	}

	return 0, nil
}

// subscriberradiusinterface
// Configuration for RADIUS interface Parameters resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/subscriber/subscriberradiusinterface

func (s *SubscriberService) UpdateSubscriberRADIUSInterface(param models.SubscriberRadiusInterface) error {
	payload := map[string]any{
		"subscriberradiusinterface": param,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, subscriberRADIUSInterfaceURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SubscriberService) UnsetSubscriberRADIUSInterface(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"subscriberradiusinterface": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, subscriberRADIUSInterfaceURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SubscriberService) GetAllSubscriberRADIUSInterface() (models.SubscriberRadiusInterface, error) {
	req, err := s.client.NewRequest(http.MethodGet, subscriberRADIUSInterfaceURL, nil)
	if err != nil {
		return models.SubscriberRadiusInterface{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.SubscriberRadiusInterface{}, err
	}

	var result struct {
		SubscriberRADIUSInterfaces []models.SubscriberRadiusInterface `json:"subscriberradiusinterface"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.SubscriberRadiusInterface{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.SubscriberRADIUSInterfaces) > 0 {
		return result.SubscriberRADIUSInterfaces[0], nil
	}

	return models.SubscriberRadiusInterface{}, nil
}

// subscribersessions
// Configuration for subscriber sesions resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/subscriber/subscribersessions

func (s *SubscriberService) ClearSubscriberSessions(session models.SubscriberSessions) error {
	payload := map[string]any{
		"subscribersessions": session,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, subscriberSessionsURL+"?action=clear", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *SubscriberService) GetAllSubscriberSessions() ([]models.SubscriberSessions, error) {
	req, err := s.client.NewRequest(http.MethodGet, subscriberSessionsURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		SubscriberSessions []models.SubscriberSessions `json:"subscribersessions"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.SubscriberSessions, nil
}

func (s *SubscriberService) CountSubscriberSessions() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, subscriberSessionsURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		SubscriberSessions []struct {
			Count float64 `json:"__count"`
		} `json:"subscribersessions"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.SubscriberSessions) > 0 {
		return result.SubscriberSessions[0].Count, nil
	}

	return 0, nil
}
