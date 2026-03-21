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
	cacheContentGroupURL                    = "/nitro/v1/config/cachecontentgroup"
	cacheForwardProxyURL                    = "/nitro/v1/config/cacheforwardproxy"
	cacheGlobalBindingURL                   = "/nitro/v1/config/cacheglobal_binding"
	cacheGlobalCachePolicyBindingURL        = "/nitro/v1/config/cacheglobal_cachepolicy_binding"
	cacheObjectURL                          = "/nitro/v1/config/cacheobject"
	cacheParameterURL                       = "/nitro/v1/config/cacheparameter"
	cachePolicyURL                          = "/nitro/v1/config/cachepolicy"
	cachePolicyBindingURL                   = "/nitro/v1/config/cachepolicy_binding"
	cachePolicyCacheGlobalBindingURL        = "/nitro/v1/config/cachepolicy_cacheglobal_binding"
	cachePolicyCachePolicyLabelBindingURL   = "/nitro/v1/config/cachepolicy_cachepolicylabel_binding"
	cachePolicyCSVServerBindingURL          = "/nitro/v1/config/cachepolicy_csvserver_binding"
	cachePolicyLBVServerBindingURL          = "/nitro/v1/config/cachepolicy_lbvserver_binding"
	cachePolicyLabelURL                     = "/nitro/v1/config/cachepolicylabel"
	cachePolicyLabelBindingURL              = "/nitro/v1/config/cachepolicylabel_binding"
	cachePolicyLabelCachePolicyBindingURL   = "/nitro/v1/config/cachepolicylabel_cachepolicy_binding"
	cachePolicyLabelPolicyBindingBindingURL = "/nitro/v1/config/cachepolicylabel_policybinding_binding"
	cacheSelectorURL                        = "/nitro/v1/config/cacheselector"
)

// Integrated Caching Configuration. The Integrated Caching feature is used to cache static and dynamic web application data.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/cache/cache
type CacheService struct {
	client *Client
}

// cachecontentgroup

func (s *CacheService) AddCacheContentGroup(group models.CacheContentGroup) error {
	payload := map[string]any{
		"cachecontentgroup": group,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, cacheContentGroupURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CacheService) DeleteCacheContentGroup(name string) error {
	reqURL := fmt.Sprintf("%s/%s", cacheContentGroupURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CacheService) UpdateCacheContentGroup(group models.CacheContentGroup) error {
	payload := map[string]any{
		"cachecontentgroup": group,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, cacheContentGroupURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CacheService) UnsetCacheContentGroup(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{
		"cachecontentgroup": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, cacheContentGroupURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CacheService) GetAllCacheContentGroup() ([]models.CacheContentGroup, error) {
	req, err := s.client.NewRequest(http.MethodGet, cacheContentGroupURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Groups []models.CacheContentGroup `json:"cachecontentgroup"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Groups, nil
}

func (s *CacheService) GetCacheContentGroup(name string) ([]models.CacheContentGroup, error) {
	reqURL := fmt.Sprintf("%s/%s", cacheContentGroupURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Groups []models.CacheContentGroup `json:"cachecontentgroup"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Groups, nil
}

func (s *CacheService) CountCacheContentGroup() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, cacheContentGroupURL+"?count=yes", nil)
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
		} `json:"cachecontentgroup"`
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

func (s *CacheService) ExpireCacheContentGroup(group models.CacheContentGroup) error {
	payload := map[string]any{
		"cachecontentgroup": group,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, cacheContentGroupURL+"?action=expire", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CacheService) FlushCacheContentGroup(group models.CacheContentGroup) error {
	payload := map[string]any{
		"cachecontentgroup": group,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, cacheContentGroupURL+"?action=flush", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CacheService) SaveCacheContentGroup(group models.CacheContentGroup) error {
	payload := map[string]any{
		"cachecontentgroup": group,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, cacheContentGroupURL+"?action=save", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

// cacheforwardproxy

func (s *CacheService) AddCacheForwardProxy(proxy models.CacheForwardProxy) error {
	payload := map[string]any{
		"cacheforwardproxy": proxy,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, cacheForwardProxyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CacheService) DeleteCacheForwardProxy(ipaddress string, port int) error {
	reqURL := fmt.Sprintf("%s/%s?args=port:%d", cacheForwardProxyURL, url.QueryEscape(ipaddress), port)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CacheService) GetAllCacheForwardProxy() ([]models.CacheForwardProxy, error) {
	req, err := s.client.NewRequest(http.MethodGet, cacheForwardProxyURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Proxies []models.CacheForwardProxy `json:"cacheforwardproxy"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Proxies, nil
}

func (s *CacheService) CountCacheForwardProxy() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, cacheForwardProxyURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Proxies []struct {
			Count float64 `json:"__count"`
		} `json:"cacheforwardproxy"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Proxies) > 0 {
		return result.Proxies[0].Count, nil
	}
	return 0, nil
}

// cacheglobal_binding

func (s *CacheService) GetCacheGlobalBinding() ([]models.CacheGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, cacheGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CacheGlobalBinding `json:"cacheglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// cacheglobal_cachepolicy_binding

func (s *CacheService) AddCacheGlobalCachePolicyBinding(binding models.CacheGlobalCachePolicyBinding) error {
	payload := map[string]any{
		"cacheglobal_cachepolicy_binding": binding,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, cacheGlobalCachePolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CacheService) DeleteCacheGlobalCachePolicyBinding(policyname string) error {
	reqURL := fmt.Sprintf("%s?args=policy:%s", cacheGlobalCachePolicyBindingURL, url.QueryEscape(policyname))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CacheService) GetCacheGlobalCachePolicyBinding() ([]models.CacheGlobalCachePolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, cacheGlobalCachePolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CacheGlobalCachePolicyBinding `json:"cacheglobal_cachepolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CacheService) CountCacheGlobalCachePolicyBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, cacheGlobalCachePolicyBindingURL+"?count=yes", nil)
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
		} `json:"cacheglobal_cachepolicy_binding"`
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

// cacheobject

func (s *CacheService) GetAllCacheObject() ([]models.CacheObject, error) {
	req, err := s.client.NewRequest(http.MethodGet, cacheObjectURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Objects []models.CacheObject `json:"cacheobject"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Objects, nil
}

func (s *CacheService) CountCacheObject() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, cacheObjectURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Objects []struct {
			Count float64 `json:"__count"`
		} `json:"cacheobject"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Objects) > 0 {
		return result.Objects[0].Count, nil
	}
	return 0, nil
}

func (s *CacheService) FlushCacheObject(object models.CacheObject) error {
	payload := map[string]any{
		"cacheobject": object,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, cacheObjectURL+"?action=flush", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CacheService) ExpireCacheObject(object models.CacheObject) error {
	payload := map[string]any{
		"cacheobject": object,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, cacheObjectURL+"?action=expire", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CacheService) SaveCacheObject(object models.CacheObject) error {
	payload := map[string]any{
		"cacheobject": object,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, cacheObjectURL+"?action=save", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

// cacheparameter

func (s *CacheService) UpdateCacheParameter(param models.CacheParameter) error {
	payload := map[string]any{
		"cacheparameter": param,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, cacheParameterURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CacheService) UnsetCacheParameter(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{
		"cacheparameter": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, cacheParameterURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CacheService) GetAllCacheParameter() (models.CacheParameter, error) {
	req, err := s.client.NewRequest(http.MethodGet, cacheParameterURL, nil)
	if err != nil {
		return models.CacheParameter{}, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return models.CacheParameter{}, err
	}
	var result struct {
		Params []models.CacheParameter `json:"cacheparameter"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.CacheParameter{}, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Params) > 0 {
		return result.Params[0], nil
	}
	return models.CacheParameter{}, nil
}

// cachepolicy

func (s *CacheService) AddCachePolicy(policy models.CachePolicy) error {
	payload := map[string]any{
		"cachepolicy": policy,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, cachePolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CacheService) DeleteCachePolicy(name string) error {
	reqURL := fmt.Sprintf("%s/%s", cachePolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CacheService) UpdateCachePolicy(policy models.CachePolicy) error {
	payload := map[string]any{
		"cachepolicy": policy,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, cachePolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CacheService) UnsetCachePolicy(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{
		"cachepolicy": unsetMap,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, cachePolicyURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CacheService) GetAllCachePolicy() ([]models.CachePolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, cachePolicyURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Policies []models.CachePolicy `json:"cachepolicy"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Policies, nil
}

func (s *CacheService) GetCachePolicy(name string) ([]models.CachePolicy, error) {
	reqURL := fmt.Sprintf("%s/%s", cachePolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Policies []models.CachePolicy `json:"cachepolicy"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Policies, nil
}

func (s *CacheService) CountCachePolicy() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, cachePolicyURL+"?count=yes", nil)
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
		} `json:"cachepolicy"`
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

func (s *CacheService) RenameCachePolicy(name, newname string) error {
	payload := map[string]any{
		"cachepolicy": map[string]string{
			"policyname": name,
			"newname":    newname,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, cachePolicyURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

// cachepolicylabel

func (s *CacheService) AddCachePolicyLabel(label models.CachePolicyLabel) error {
	payload := map[string]any{
		"cachepolicylabel": label,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, cachePolicyLabelURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CacheService) DeleteCachePolicyLabel(labelname string) error {
	reqURL := fmt.Sprintf("%s/%s", cachePolicyLabelURL, url.QueryEscape(labelname))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CacheService) GetAllCachePolicyLabel() ([]models.CachePolicyLabel, error) {
	req, err := s.client.NewRequest(http.MethodGet, cachePolicyLabelURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Labels []models.CachePolicyLabel `json:"cachepolicylabel"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Labels, nil
}

func (s *CacheService) GetCachePolicyLabel(labelname string) ([]models.CachePolicyLabel, error) {
	reqURL := fmt.Sprintf("%s/%s", cachePolicyLabelURL, url.QueryEscape(labelname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Labels []models.CachePolicyLabel `json:"cachepolicylabel"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Labels, nil
}

func (s *CacheService) CountCachePolicyLabel() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, cachePolicyLabelURL+"?count=yes", nil)
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
		} `json:"cachepolicylabel"`
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

func (s *CacheService) RenameCachePolicyLabel(labelname, newname string) error {
	payload := map[string]any{
		"cachepolicylabel": map[string]string{
			"labelname": labelname,
			"newname":   newname,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, cachePolicyLabelURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

// cachepolicylabel_binding

func (s *CacheService) GetAllCachePolicyLabelBinding() ([]models.CachePolicyLabelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, cachePolicyLabelBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CachePolicyLabelBinding `json:"cachepolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CacheService) GetCachePolicyLabelBinding(labelname string) ([]models.CachePolicyLabelBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", cachePolicyLabelBindingURL, url.QueryEscape(labelname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CachePolicyLabelBinding `json:"cachepolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// cachepolicylabel_cachepolicy_binding

func (s *CacheService) AddCachePolicyLabelCachePolicyBinding(binding models.CachePolicyLabelCachePolicyBinding) error {
	payload := map[string]any{
		"cachepolicylabel_cachepolicy_binding": binding,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, cachePolicyLabelCachePolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CacheService) DeleteCachePolicyLabelCachePolicyBinding(labelname, policyname string) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s", cachePolicyLabelCachePolicyBindingURL, url.QueryEscape(labelname), url.QueryEscape(policyname))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CacheService) GetAllCachePolicyLabelCachePolicyBinding() ([]models.CachePolicyLabelCachePolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, cachePolicyLabelCachePolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CachePolicyLabelCachePolicyBinding `json:"cachepolicylabel_cachepolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CacheService) GetCachePolicyLabelCachePolicyBinding(labelname string) ([]models.CachePolicyLabelCachePolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", cachePolicyLabelCachePolicyBindingURL, url.QueryEscape(labelname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CachePolicyLabelCachePolicyBinding `json:"cachepolicylabel_cachepolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CacheService) CountCachePolicyLabelCachePolicyBinding(labelname string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", cachePolicyLabelCachePolicyBindingURL, url.QueryEscape(labelname))
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
		} `json:"cachepolicylabel_cachepolicy_binding"`
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

// cachepolicylabel_policybinding_binding

func (s *CacheService) GetAllCachePolicyLabelPolicyBindingBinding() ([]models.CachePolicyLabelPolicyBindingBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, cachePolicyLabelPolicyBindingBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CachePolicyLabelPolicyBindingBinding `json:"cachepolicylabel_policybinding_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CacheService) GetCachePolicyLabelPolicyBindingBinding(labelname string) ([]models.CachePolicyLabelPolicyBindingBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", cachePolicyLabelPolicyBindingBindingURL, url.QueryEscape(labelname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CachePolicyLabelPolicyBindingBinding `json:"cachepolicylabel_policybinding_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CacheService) CountCachePolicyLabelPolicyBindingBinding(labelname string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", cachePolicyLabelPolicyBindingBindingURL, url.QueryEscape(labelname))
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
		} `json:"cachepolicylabel_policybinding_binding"`
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

// cachepolicy_binding

func (s *CacheService) GetAllCachePolicyBinding() ([]models.CachePolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, cachePolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CachePolicyBinding `json:"cachepolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CacheService) GetCachePolicyBinding(policyname string) ([]models.CachePolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", cachePolicyBindingURL, url.QueryEscape(policyname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CachePolicyBinding `json:"cachepolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

// cachepolicy_cacheglobal_binding

func (s *CacheService) GetAllCachePolicyCacheGlobalBinding() ([]models.CachePolicyCacheGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, cachePolicyCacheGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CachePolicyCacheGlobalBinding `json:"cachepolicy_cacheglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CacheService) GetCachePolicyCacheGlobalBinding(policyname string) ([]models.CachePolicyCacheGlobalBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", cachePolicyCacheGlobalBindingURL, url.QueryEscape(policyname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CachePolicyCacheGlobalBinding `json:"cachepolicy_cacheglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CacheService) CountCachePolicyCacheGlobalBinding(policyname string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", cachePolicyCacheGlobalBindingURL, url.QueryEscape(policyname))
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
		} `json:"cachepolicy_cacheglobal_binding"`
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

// cachepolicy_cachepolicylabel_binding

func (s *CacheService) GetAllCachePolicyCachePolicyLabelBinding() ([]models.CachePolicyCachePolicyLabelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, cachePolicyCachePolicyLabelBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CachePolicyCachePolicyLabelBinding `json:"cachepolicy_cachepolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CacheService) GetCachePolicyCachePolicyLabelBinding(policyname string) ([]models.CachePolicyCachePolicyLabelBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", cachePolicyCachePolicyLabelBindingURL, url.QueryEscape(policyname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CachePolicyCachePolicyLabelBinding `json:"cachepolicy_cachepolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CacheService) CountCachePolicyCachePolicyLabelBinding(policyname string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", cachePolicyCachePolicyLabelBindingURL, url.QueryEscape(policyname))
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
		} `json:"cachepolicy_cachepolicylabel_binding"`
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

// cachepolicy_csvserver_binding

func (s *CacheService) GetAllCachePolicyCSVServerBinding() ([]models.CachePolicyCSVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, cachePolicyCSVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CachePolicyCSVServerBinding `json:"cachepolicy_csvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CacheService) GetCachePolicyCSVServerBinding(policyname string) ([]models.CachePolicyCSVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", cachePolicyCSVServerBindingURL, url.QueryEscape(policyname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CachePolicyCSVServerBinding `json:"cachepolicy_csvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CacheService) CountCachePolicyCSVServerBinding(policyname string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", cachePolicyCSVServerBindingURL, url.QueryEscape(policyname))
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
		} `json:"cachepolicy_csvserver_binding"`
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

// cachepolicy_lbvserver_binding

func (s *CacheService) GetAllCachePolicyLBVServerBinding() ([]models.CachePolicyLBVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, cachePolicyLBVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CachePolicyLBVServerBinding `json:"cachepolicy_lbvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CacheService) GetCachePolicyLBVServerBinding(policyname string) ([]models.CachePolicyLBVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", cachePolicyLBVServerBindingURL, url.QueryEscape(policyname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Bindings []models.CachePolicyLBVServerBinding `json:"cachepolicy_lbvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Bindings, nil
}

func (s *CacheService) CountCachePolicyLBVServerBinding(policyname string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", cachePolicyLBVServerBindingURL, url.QueryEscape(policyname))
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
		} `json:"cachepolicy_lbvserver_binding"`
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

// cacheselector

func (s *CacheService) AddCacheSelector(selector models.CacheSelector) error {
	payload := map[string]any{
		"cacheselector": selector,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, cacheSelectorURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CacheService) DeleteCacheSelector(selectorname string) error {
	reqURL := fmt.Sprintf("%s/%s", cacheSelectorURL, url.QueryEscape(selectorname))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CacheService) UpdateCacheSelector(selector models.CacheSelector) error {
	payload := map[string]any{
		"cacheselector": selector,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, cacheSelectorURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *CacheService) GetAllCacheSelector() ([]models.CacheSelector, error) {
	req, err := s.client.NewRequest(http.MethodGet, cacheSelectorURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Selectors []models.CacheSelector `json:"cacheselector"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Selectors, nil
}

func (s *CacheService) GetCacheSelector(selectorname string) ([]models.CacheSelector, error) {
	reqURL := fmt.Sprintf("%s/%s", cacheSelectorURL, url.QueryEscape(selectorname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Selectors []models.CacheSelector `json:"cacheselector"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Selectors, nil
}

func (s *CacheService) CountCacheSelector() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, cacheSelectorURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Selectors []struct {
			Count float64 `json:"__count"`
		} `json:"cacheselector"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Selectors) > 0 {
		return result.Selectors[0].Count, nil
	}
	return 0, nil
}
