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
	botGlobalBindingURL                   = "/nitro/v1/config/botglobal_binding"
	botGlobalBotPolicyBindingURL          = "/nitro/v1/config/botglobal_botpolicy_binding"
	botPolicyURL                          = "/nitro/v1/config/botpolicy"
	botPolicyBindingURL                   = "/nitro/v1/config/botpolicy_binding"
	botPolicyBotGlobalBindingURL          = "/nitro/v1/config/botpolicy_botglobal_binding"
	botPolicyBotPolicyLabelBindingURL     = "/nitro/v1/config/botpolicy_botpolicylabel_binding"
	botPolicyCSVServerBindingURL          = "/nitro/v1/config/botpolicy_csvserver_binding"
	botPolicyLBVServerBindingURL          = "/nitro/v1/config/botpolicy_lbvserver_binding"
	botPolicyLabelURL                     = "/nitro/v1/config/botpolicylabel"
	botPolicyLabelBindingURL              = "/nitro/v1/config/botpolicylabel_binding"
	botPolicyLabelBotPolicyBindingURL     = "/nitro/v1/config/botpolicylabel_botpolicy_binding"
	botPolicyLabelPolicyBindingBindingURL = "/nitro/v1/config/botpolicylabel_policybinding_binding"
	botProfileURL                         = "/nitro/v1/config/botprofile"
	botProfileBindingURL                  = "/nitro/v1/config/botprofile_binding"
	botProfileBlacklistBindingURL         = "/nitro/v1/config/botprofile_blacklist_binding"
	botProfileCaptchaBindingURL           = "/nitro/v1/config/botprofile_captcha_binding"
	botProfileIPReputationBindingURL      = "/nitro/v1/config/botprofile_ipreputation_binding"
	botProfileRatelimitBindingURL         = "/nitro/v1/config/botprofile_ratelimit_binding"
	botProfileTPSBindingURL               = "/nitro/v1/config/botprofile_tps_binding"
	botProfileWhitelistBindingURL         = "/nitro/v1/config/botprofile_whitelist_binding"
	botSettingsURL                        = "/nitro/v1/config/botsettings"
	botSignatureURL                       = "/nitro/v1/config/botsignature"
)

// Bot Management.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/bot/bot
type BotService struct {
	client *Client
}

// botglobal_binding

func (s *BotService) GetBotGlobalBinding() ([]models.BotGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, botGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.BotGlobalBinding `json:"botglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// botglobal_botpolicy_binding

func (s *BotService) AddBotGlobalBotPolicyBinding(binding models.BotGlobalBotPolicyBinding) error {
	payload := map[string]any{
		"botglobal_botpolicy_binding": binding,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, botGlobalBotPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *BotService) DeleteBotGlobalBotPolicyBinding(policyname, typeField string, priority int) error {
	reqURL := fmt.Sprintf("%s?args=policyname:%s,type:%s,priority:%d", botGlobalBotPolicyBindingURL, url.QueryEscape(policyname), url.QueryEscape(typeField), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *BotService) GetBotGlobalBotPolicyBinding() ([]models.BotGlobalBotPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, botGlobalBotPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.BotGlobalBotPolicyBinding `json:"botglobal_botpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *BotService) CountBotGlobalBotPolicyBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, botGlobalBotPolicyBindingURL+"?count=yes", nil)
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
		} `json:"botglobal_botpolicy_binding"`
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

// botpolicy

func (s *BotService) AddBotPolicy(policy models.BotPolicy) error {
	payload := map[string]any{
		"botpolicy": policy,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, botPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *BotService) DeleteBotPolicy(name string) error {
	reqURL := fmt.Sprintf("%s/%s", botPolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *BotService) UpdateBotPolicy(policy models.BotPolicy) error {
	payload := map[string]any{
		"botpolicy": policy,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, botPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *BotService) UnsetBotPolicy(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{
		"botpolicy": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, botPolicyURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *BotService) RenameBotPolicy(name, newname string) error {
	payload := map[string]any{
		"botpolicy": map[string]string{
			"name":    name,
			"newname": newname,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, botPolicyURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *BotService) GetAllBotPolicy() ([]models.BotPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, botPolicyURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Policies []models.BotPolicy `json:"botpolicy"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Policies, nil
}

func (s *BotService) GetBotPolicy(name string) ([]models.BotPolicy, error) {
	reqURL := fmt.Sprintf("%s/%s", botPolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Policies []models.BotPolicy `json:"botpolicy"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Policies, nil
}

func (s *BotService) CountBotPolicy() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, botPolicyURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Policies []struct {
			Count float64 `json:"__count"`
		} `json:"botpolicy"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Policies) > 0 {
		return result.Policies[0].Count, nil
	}
	return 0, nil
}

// botpolicylabel

func (s *BotService) AddBotPolicyLabel(label models.BotPolicyLabel) error {
	payload := map[string]any{
		"botpolicylabel": label,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, botPolicyLabelURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *BotService) DeleteBotPolicyLabel(labelname string) error {
	reqURL := fmt.Sprintf("%s/%s", botPolicyLabelURL, url.QueryEscape(labelname))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *BotService) RenameBotPolicyLabel(labelname, newname string) error {
	payload := map[string]any{
		"botpolicylabel": map[string]string{
			"labelname": labelname,
			"newname":   newname,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, botPolicyLabelURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *BotService) GetAllBotPolicyLabel() ([]models.BotPolicyLabel, error) {
	req, err := s.client.NewRequest(http.MethodGet, botPolicyLabelURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Labels []models.BotPolicyLabel `json:"botpolicylabel"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Labels, nil
}

func (s *BotService) GetBotPolicyLabel(labelname string) ([]models.BotPolicyLabel, error) {
	reqURL := fmt.Sprintf("%s/%s", botPolicyLabelURL, url.QueryEscape(labelname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Labels []models.BotPolicyLabel `json:"botpolicylabel"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Labels, nil
}

func (s *BotService) CountBotPolicyLabel() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, botPolicyLabelURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Labels []struct {
			Count float64 `json:"__count"`
		} `json:"botpolicylabel"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Labels) > 0 {
		return result.Labels[0].Count, nil
	}
	return 0, nil
}

// botpolicylabel_binding

func (s *BotService) GetAllBotPolicyLabelBinding() ([]models.BotPolicyLabelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, botPolicyLabelBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.BotPolicyLabelBinding `json:"botpolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *BotService) GetBotPolicyLabelBinding(labelname string) ([]models.BotPolicyLabelBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", botPolicyLabelBindingURL, url.QueryEscape(labelname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.BotPolicyLabelBinding `json:"botpolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// botpolicylabel_botpolicy_binding

func (s *BotService) AddBotPolicyLabelBotPolicyBinding(binding models.BotPolicyLabelBotPolicyBinding) error {
	payload := map[string]any{
		"botpolicylabel_botpolicy_binding": binding,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, botPolicyLabelBotPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *BotService) DeleteBotPolicyLabelBotPolicyBinding(labelname, policyname string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s,priority:%d", botPolicyLabelBotPolicyBindingURL, url.QueryEscape(labelname), url.QueryEscape(policyname), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *BotService) GetAllBotPolicyLabelBotPolicyBinding() ([]models.BotPolicyLabelBotPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, botPolicyLabelBotPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.BotPolicyLabelBotPolicyBinding `json:"botpolicylabel_botpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *BotService) GetBotPolicyLabelBotPolicyBinding(labelname string) ([]models.BotPolicyLabelBotPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", botPolicyLabelBotPolicyBindingURL, url.QueryEscape(labelname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.BotPolicyLabelBotPolicyBinding `json:"botpolicylabel_botpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *BotService) CountBotPolicyLabelBotPolicyBinding(labelname string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", botPolicyLabelBotPolicyBindingURL, url.QueryEscape(labelname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
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
		} `json:"botpolicylabel_botpolicy_binding"`
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

// botpolicylabel_policybinding_binding

func (s *BotService) GetAllBotPolicyLabelPolicyBindingBinding() ([]models.BotPolicyLabelPolicyBindingBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, botPolicyLabelPolicyBindingBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.BotPolicyLabelPolicyBindingBinding `json:"botpolicylabel_policybinding_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *BotService) GetBotPolicyLabelPolicyBindingBinding(labelname string) ([]models.BotPolicyLabelPolicyBindingBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", botPolicyLabelPolicyBindingBindingURL, url.QueryEscape(labelname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.BotPolicyLabelPolicyBindingBinding `json:"botpolicylabel_policybinding_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *BotService) CountBotPolicyLabelPolicyBindingBinding(labelname string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", botPolicyLabelPolicyBindingBindingURL, url.QueryEscape(labelname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
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
		} `json:"botpolicylabel_policybinding_binding"`
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

// botpolicy_binding

func (s *BotService) GetAllBotPolicyBinding() ([]models.BotPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, botPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.BotPolicyBinding `json:"botpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *BotService) GetBotPolicyBinding(name string) ([]models.BotPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", botPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.BotPolicyBinding `json:"botpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// botpolicy_botglobal_binding

func (s *BotService) GetAllBotPolicyBotGlobalBinding() ([]models.BotPolicyBotGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, botPolicyBotGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.BotPolicyBotGlobalBinding `json:"botpolicy_botglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *BotService) GetBotPolicyBotGlobalBinding(name string) ([]models.BotPolicyBotGlobalBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", botPolicyBotGlobalBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.BotPolicyBotGlobalBinding `json:"botpolicy_botglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *BotService) CountBotPolicyBotGlobalBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", botPolicyBotGlobalBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
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
		} `json:"botpolicy_botglobal_binding"`
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

// botpolicy_botpolicylabel_binding

func (s *BotService) GetAllBotPolicyBotPolicyLabelBinding() ([]models.BotPolicyBotPolicyLabelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, botPolicyBotPolicyLabelBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.BotPolicyBotPolicyLabelBinding `json:"botpolicy_botpolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *BotService) GetBotPolicyBotPolicyLabelBinding(name string) ([]models.BotPolicyBotPolicyLabelBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", botPolicyBotPolicyLabelBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.BotPolicyBotPolicyLabelBinding `json:"botpolicy_botpolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *BotService) CountBotPolicyBotPolicyLabelBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", botPolicyBotPolicyLabelBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
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
		} `json:"botpolicy_botpolicylabel_binding"`
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

// botpolicy_csvserver_binding

func (s *BotService) GetAllBotPolicyCSVServerBinding() ([]models.BotPolicyCSVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, botPolicyCSVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.BotPolicyCSVServerBinding `json:"botpolicy_csvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *BotService) GetBotPolicyCSVServerBinding(name string) ([]models.BotPolicyCSVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", botPolicyCSVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.BotPolicyCSVServerBinding `json:"botpolicy_csvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *BotService) CountBotPolicyCSVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", botPolicyCSVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
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
		} `json:"botpolicy_csvserver_binding"`
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

// botpolicy_lbvserver_binding

func (s *BotService) GetAllBotPolicyLBVServerBinding() ([]models.BotPolicyLBVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, botPolicyLBVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.BotPolicyLBVServerBinding `json:"botpolicy_lbvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *BotService) GetBotPolicyLBVServerBinding(name string) ([]models.BotPolicyLBVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", botPolicyLBVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.BotPolicyLBVServerBinding `json:"botpolicy_lbvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *BotService) CountBotPolicyLBVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", botPolicyLBVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
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
		} `json:"botpolicy_lbvserver_binding"`
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

// botprofile

func (s *BotService) AddBotProfile(profile models.BotProfile) error {
	payload := map[string]any{
		"botprofile": profile,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, botProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *BotService) DeleteBotProfile(name string) error {
	reqURL := fmt.Sprintf("%s/%s", botProfileURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *BotService) UpdateBotProfile(profile models.BotProfile) error {
	payload := map[string]any{
		"botprofile": profile,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, botProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *BotService) UnsetBotProfile(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{
		"botprofile": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, botProfileURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *BotService) GetAllBotProfile() ([]models.BotProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, botProfileURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Profiles []models.BotProfile `json:"botprofile"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Profiles, nil
}

func (s *BotService) GetBotProfile(name string) ([]models.BotProfile, error) {
	reqURL := fmt.Sprintf("%s/%s", botProfileURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Profiles []models.BotProfile `json:"botprofile"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Profiles, nil
}

func (s *BotService) CountBotProfile() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, botProfileURL+"?count=yes", nil)
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
		} `json:"botprofile"`
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

// botprofile_binding

func (s *BotService) GetAllBotProfileBinding() ([]models.BotProfileBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, botProfileBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.BotProfileBinding `json:"botprofile_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *BotService) GetBotProfileBinding(name string) ([]models.BotProfileBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", botProfileBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.BotProfileBinding `json:"botprofile_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// botprofile_blacklist_binding

func (s *BotService) AddBotProfileBlacklistBinding(binding models.BotProfileBlacklistBinding) error {
	payload := map[string]any{
		"botprofile_blacklist_binding": binding,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, botProfileBlacklistBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *BotService) DeleteBotProfileBlacklistBinding(name, blacklistType, blacklistValue string) error {
	reqURL := fmt.Sprintf("%s/%s?args=bot_blacklist_type:%s,bot_blacklist_value:%s", botProfileBlacklistBindingURL, url.QueryEscape(name), url.QueryEscape(blacklistType), url.QueryEscape(blacklistValue))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *BotService) GetAllBotProfileBlacklistBinding() ([]models.BotProfileBlacklistBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, botProfileBlacklistBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.BotProfileBlacklistBinding `json:"botprofile_blacklist_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *BotService) GetBotProfileBlacklistBinding(name string) ([]models.BotProfileBlacklistBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", botProfileBlacklistBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.BotProfileBlacklistBinding `json:"botprofile_blacklist_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *BotService) CountBotProfileBlacklistBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", botProfileBlacklistBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
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
		} `json:"botprofile_blacklist_binding"`
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

// botprofile_captcha_binding

func (s *BotService) AddBotProfileCaptchaBinding(binding models.BotProfileCaptchaBinding) error {
	payload := map[string]any{
		"botprofile_captcha_binding": binding,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, botProfileCaptchaBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *BotService) DeleteBotProfileCaptchaBinding(name, captchaURL string) error {
	reqURL := fmt.Sprintf("%s/%s?args=bot_captcha_url:%s", botProfileCaptchaBindingURL, url.QueryEscape(name), url.QueryEscape(captchaURL))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *BotService) GetAllBotProfileCaptchaBinding() ([]models.BotProfileCaptchaBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, botProfileCaptchaBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.BotProfileCaptchaBinding `json:"botprofile_captcha_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *BotService) GetBotProfileCaptchaBinding(name string) ([]models.BotProfileCaptchaBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", botProfileCaptchaBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.BotProfileCaptchaBinding `json:"botprofile_captcha_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *BotService) CountBotProfileCaptchaBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", botProfileCaptchaBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
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
		} `json:"botprofile_captcha_binding"`
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

// botprofile_ipreputation_binding

func (s *BotService) AddBotProfileIPReputationBinding(binding models.BotProfileIPReputationBinding) error {
	payload := map[string]any{
		"botprofile_ipreputation_binding": binding,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, botProfileIPReputationBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *BotService) DeleteBotProfileIPReputationBinding(name, category string) error {
	reqURL := fmt.Sprintf("%s/%s?args=category:%s", botProfileIPReputationBindingURL, url.QueryEscape(name), url.QueryEscape(category))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *BotService) GetAllBotProfileIPReputationBinding() ([]models.BotProfileIPReputationBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, botProfileIPReputationBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.BotProfileIPReputationBinding `json:"botprofile_ipreputation_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *BotService) GetBotProfileIPReputationBinding(name string) ([]models.BotProfileIPReputationBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", botProfileIPReputationBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.BotProfileIPReputationBinding `json:"botprofile_ipreputation_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *BotService) CountBotProfileIPReputationBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", botProfileIPReputationBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
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
		} `json:"botprofile_ipreputation_binding"`
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

// botprofile_ratelimit_binding

func (s *BotService) AddBotProfileRatelimitBinding(binding models.BotProfileRateLimitBinding) error {
	payload := map[string]any{
		"botprofile_ratelimit_binding": binding,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, botProfileRatelimitBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *BotService) DeleteBotProfileRatelimitBinding(name, limitType, botRateLimitURL string) error {
	reqURL := fmt.Sprintf("%s/%s?args=bot_rate_limit_type:%s,bot_rate_limit_url:%s", botProfileRatelimitBindingURL, url.QueryEscape(name), url.QueryEscape(limitType), url.QueryEscape(botRateLimitURL))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *BotService) GetAllBotProfileRatelimitBinding() ([]models.BotProfileRateLimitBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, botProfileRatelimitBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.BotProfileRateLimitBinding `json:"botprofile_ratelimit_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *BotService) GetBotProfileRatelimitBinding(name string) ([]models.BotProfileRateLimitBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", botProfileRatelimitBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.BotProfileRateLimitBinding `json:"botprofile_ratelimit_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *BotService) CountBotProfileRatelimitBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", botProfileRatelimitBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
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
		} `json:"botprofile_ratelimit_binding"`
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

// botprofile_tps_binding

func (s *BotService) AddBotProfileTPSBinding(binding models.BotProfileTPSBinding) error {
	payload := map[string]any{
		"botprofile_tps_binding": binding,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, botProfileTPSBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *BotService) DeleteBotProfileTPSBinding(name, tpsType string) error {
	reqURL := fmt.Sprintf("%s/%s?args=bot_tps_type:%s", botProfileTPSBindingURL, url.QueryEscape(name), url.QueryEscape(tpsType))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *BotService) GetAllBotProfileTPSBinding() ([]models.BotProfileTPSBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, botProfileTPSBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.BotProfileTPSBinding `json:"botprofile_tps_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *BotService) GetBotProfileTPSBinding(name string) ([]models.BotProfileTPSBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", botProfileTPSBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.BotProfileTPSBinding `json:"botprofile_tps_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *BotService) CountBotProfileTPSBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", botProfileTPSBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
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
		} `json:"botprofile_tps_binding"`
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

// botprofile_whitelist_binding

func (s *BotService) AddBotProfileWhitelistBinding(binding models.BotProfileWhitelistBinding) error {
	payload := map[string]any{
		"botprofile_whitelist_binding": binding,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, botProfileWhitelistBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *BotService) DeleteBotProfileWhitelistBinding(name, whitelistType, whitelistValue string) error {
	reqURL := fmt.Sprintf("%s/%s?args=bot_whitelist_type:%s,bot_whitelist_value:%s", botProfileWhitelistBindingURL, url.QueryEscape(name), url.QueryEscape(whitelistType), url.QueryEscape(whitelistValue))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *BotService) GetAllBotProfileWhitelistBinding() ([]models.BotProfileWhitelistBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, botProfileWhitelistBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.BotProfileWhitelistBinding `json:"botprofile_whitelist_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *BotService) GetBotProfileWhitelistBinding(name string) ([]models.BotProfileWhitelistBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", botProfileWhitelistBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.BotProfileWhitelistBinding `json:"botprofile_whitelist_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *BotService) CountBotProfileWhitelistBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", botProfileWhitelistBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
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
		} `json:"botprofile_whitelist_binding"`
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

// botsettings

func (s *BotService) UpdateBotSettings(settings models.BotSettings) error {
	payload := map[string]any{
		"botsettings": settings,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, botSettingsURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *BotService) UnsetBotSettings(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{
		"botsettings": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, botSettingsURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *BotService) GetAllBotSettings() (models.BotSettings, error) {
	req, err := s.client.NewRequest(http.MethodGet, botSettingsURL, nil)
	if err != nil {
		return models.BotSettings{}, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return models.BotSettings{}, err
	}
	var result struct {
		Settings []models.BotSettings `json:"botsettings"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.BotSettings{}, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Settings) > 0 {
		return result.Settings[0], nil
	}
	return models.BotSettings{}, nil
}

// botsignature

func (s *BotService) ImportBotSignature(signature models.BotSignature) error {
	payload := map[string]any{
		"botsignature": signature,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, botSignatureURL+"?action=Import", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *BotService) DeleteBotSignature(name string) error {
	reqURL := fmt.Sprintf("%s/%s", botSignatureURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *BotService) ChangeBotSignature(signature models.BotSignature) error {
	payload := map[string]any{
		"botsignature": signature,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, botSignatureURL+"?action=update", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *BotService) GetAllBotSignature() ([]models.BotSignature, error) {
	req, err := s.client.NewRequest(http.MethodGet, botSignatureURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Signatures []models.BotSignature `json:"botsignature"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Signatures, nil
}

func (s *BotService) GetBotSignature(name string) ([]models.BotSignature, error) {
	reqURL := fmt.Sprintf("%s/%s", botSignatureURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Signatures []models.BotSignature `json:"botsignature"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Signatures, nil
}
