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
	applicationURL = "/nitro/v1/config/application"
)

// Application
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/app/app
type AppService struct {
	client *Client
}

// application
// Configuration for application resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/app/application

func (s *AppService) ImportApplication(app models.Application) error {
	payload := map[string]any{
		"application": app,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, applicationURL+"?action=Import", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppService) ExportApplication(app models.Application) error {
	payload := map[string]any{
		"application": app,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, applicationURL+"?action=export", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AppService) DeleteApplication(appname string) error {
	reqURL := fmt.Sprintf("%s?args=appname:%s", applicationURL, url.QueryEscape(appname))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}
