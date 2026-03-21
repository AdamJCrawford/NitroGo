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
	routerDynamicRoutingURL = "/nitro/v1/config/routerdynamicrouting"
)

// Router configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/router/router
type RouterService struct {
	client *Client
}

// routerdynamicrouting
// Configuration for dynamic routing config resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/router/routerdynamicrouting

func (s *RouterService) AddRouterDynamicRouting(routing models.RouterDynamicRouting) error {
	payload := map[string]any{
		"routerdynamicrouting": routing,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, routerDynamicRoutingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *RouterService) DeleteRouterDynamicRouting(commandString string) error {
	// The DELETE for routerdynamicrouting has URL parameters args=commandstring:<String_value>
	reqURL := fmt.Sprintf("%s?args=commandstring:%s", routerDynamicRoutingURL, url.QueryEscape(commandString))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *RouterService) UpdateRouterDynamicRouting(routing models.RouterDynamicRouting) error {
	payload := map[string]any{
		"routerdynamicrouting": routing,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, routerDynamicRoutingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *RouterService) UnsetRouterDynamicRouting(args []string) error {
	unsetMap := make(map[string]bool)
	for _, arg := range args {
		unsetMap[arg] = true
	}

	payload := map[string]any{
		"routerdynamicrouting": unsetMap,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, routerDynamicRoutingURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *RouterService) GetAllRouterDynamicRouting() ([]models.RouterDynamicRouting, error) {
	req, err := s.client.NewRequest(http.MethodGet, routerDynamicRoutingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		RouterDynamicRoutings []models.RouterDynamicRouting `json:"routerdynamicrouting"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.RouterDynamicRoutings, nil
}

func (s *RouterService) CountRouterDynamicRouting() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, routerDynamicRoutingURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		RouterDynamicRoutings []struct {
			Count float64 `json:"__count"`
		} `json:"routerdynamicrouting"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.RouterDynamicRoutings) > 0 {
		return result.RouterDynamicRoutings[0].Count, nil
	}

	return 0, nil
}

func (s *RouterService) ApplyRouterDynamicRouting(routing models.RouterDynamicRouting) error {
	payload := map[string]any{
		"routerdynamicrouting": routing,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, routerDynamicRoutingURL+"?action=apply", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}
