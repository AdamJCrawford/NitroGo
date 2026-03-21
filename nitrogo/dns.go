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
	dnsAAAARecURL                         = "/nitro/v1/config/dnsaaaarec"
	dnsActionURL                          = "/nitro/v1/config/dnsaction"
	dnsAction64URL                        = "/nitro/v1/config/dnsaction64"
	dnsAddRecURL                          = "/nitro/v1/config/dnsaddrec"
	dnsCNameRecURL                        = "/nitro/v1/config/dnscnamerec"
	dnsGlobalBindingURL                   = "/nitro/v1/config/dnsglobal_binding"
	dnsGlobalDNSPolicyBindingURL          = "/nitro/v1/config/dnsglobal_dnspolicy_binding"
	dnsKeyURL                             = "/nitro/v1/config/dnskey"
	dnsMXRecURL                           = "/nitro/v1/config/dnsmxrec"
	dnsNameServerURL                      = "/nitro/v1/config/dnsnameserver"
	dnsNAPTRRecURL                        = "/nitro/v1/config/dnsnaptrrec"
	dnsNSECRecURL                         = "/nitro/v1/config/dnsnsecrec"
	dnsNSRecURL                           = "/nitro/v1/config/dnsnsrec"
	dnsParameterURL                       = "/nitro/v1/config/dnsparameter"
	dnsPolicyURL                          = "/nitro/v1/config/dnspolicy"
	dnsPolicy64URL                        = "/nitro/v1/config/dnspolicy64"
	dnsPolicy64BindingURL                 = "/nitro/v1/config/dnspolicy64_binding"
	dnsPolicy64LBVServerBindingURL        = "/nitro/v1/config/dnspolicy64_lbvserver_binding"
	dnsPolicyLabelURL                     = "/nitro/v1/config/dnspolicylabel"
	dnsPolicyLabelBindingURL              = "/nitro/v1/config/dnspolicylabel_binding"
	dnsPolicyLabelDNSPolicyBindingURL     = "/nitro/v1/config/dnspolicylabel_dnspolicy_binding"
	dnsPolicyLabelPolicyBindingBindingURL = "/nitro/v1/config/dnspolicylabel_policybinding_binding"
	dnsPolicyBindingURL                   = "/nitro/v1/config/dnspolicy_binding"
	dnsPolicyDNSGlobalBindingURL          = "/nitro/v1/config/dnspolicy_dnsglobal_binding"
	dnsPolicyDNSPolicyLabelBindingURL     = "/nitro/v1/config/dnspolicy_dnspolicylabel_binding"
	dnsProfileURL                         = "/nitro/v1/config/dnsprofile"
	dnsProxyRecordsURL                    = "/nitro/v1/config/dnsproxyrecords"
	dnsPTRRecURL                          = "/nitro/v1/config/dnsptrrec"
	dnsSOARecURL                          = "/nitro/v1/config/dnssoarec"
	dnsSRVRecURL                          = "/nitro/v1/config/dnssrvrec"
	dnsSubnetCacheURL                     = "/nitro/v1/config/dnssubnetcache"
	dnsSuffixURL                          = "/nitro/v1/config/dnssuffix"
	dnsTXTRecURL                          = "/nitro/v1/config/dnstxtrec"
	dnsViewURL                            = "/nitro/v1/config/dnsview"
	dnsViewBindingURL                     = "/nitro/v1/config/dnsview_binding"
	dnsViewDNSPolicyBindingURL            = "/nitro/v1/config/dnsview_dnspolicy_binding"
	dnsViewGSLBServiceBindingURL          = "/nitro/v1/config/dnsview_gslbservice_binding"
	dnsZoneURL                            = "/nitro/v1/config/dnszone"
	dnsZoneBindingURL                     = "/nitro/v1/config/dnszone_binding"
	dnsZoneDNSKeyBindingURL               = "/nitro/v1/config/dnszone_dnskey_binding"
	dnsZoneDomainBindingURL               = "/nitro/v1/config/dnszone_domain_binding"
)

// Domain Name Service(DNS) configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/dns/dns
type DNSService struct {
	client *Client
}

// dnsaaaarec
func (s *DNSService) AddDNSAAAARec(rec models.DNSAAAARec) error {
	payload := map[string]any{"dnsaaaarec": rec}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsAAAARecURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) DeleteDNSAAAARec(hostname string) error {
	reqURL := fmt.Sprintf("%s/%s", dnsAAAARecURL, url.QueryEscape(hostname))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) GetAllDNSAAAARec() ([]models.DNSAAAARec, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsAAAARecURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSAAAARec `json:"dnsaaaarec"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) CountDNSAAAARec() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsAAAARecURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Items []struct {
			Count float64 `json:"__count"`
		} `json:"dnsaaaarec"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Items) > 0 {
		return result.Items[0].Count, nil
	}
	return 0, nil
}

// dnsaction
func (s *DNSService) AddDNSAction(action models.DNSAction) error {
	payload := map[string]any{"dnsaction": action}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) DeleteDNSAction(actionName string) error {
	reqURL := fmt.Sprintf("%s/%s", dnsActionURL, url.QueryEscape(actionName))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) UpdateDNSAction(action models.DNSAction) error {
	payload := map[string]any{"dnsaction": action}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, dnsActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) UnsetDNSAction(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"dnsaction": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsActionURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) GetAllDNSAction() ([]models.DNSAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsActionURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSAction `json:"dnsaction"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) GetDNSAction(actionName string) ([]models.DNSAction, error) {
	reqURL := fmt.Sprintf("%s/%s", dnsActionURL, url.QueryEscape(actionName))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSAction `json:"dnsaction"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) CountDNSAction() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsActionURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Items []struct {
			Count float64 `json:"__count"`
		} `json:"dnsaction"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Items) > 0 {
		return result.Items[0].Count, nil
	}
	return 0, nil
}

// dnsaction64
func (s *DNSService) AddDNSAction64(action models.DNSAction64) error {
	payload := map[string]any{"dnsaction64": action}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsAction64URL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) DeleteDNSAction64(actionName string) error {
	reqURL := fmt.Sprintf("%s/%s", dnsAction64URL, url.QueryEscape(actionName))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) UpdateDNSAction64(action models.DNSAction64) error {
	payload := map[string]any{"dnsaction64": action}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, dnsAction64URL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) UnsetDNSAction64(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"dnsaction64": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsAction64URL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) GetAllDNSAction64() ([]models.DNSAction64, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsAction64URL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSAction64 `json:"dnsaction64"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) GetDNSAction64(actionName string) ([]models.DNSAction64, error) {
	reqURL := fmt.Sprintf("%s/%s", dnsAction64URL, url.QueryEscape(actionName))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSAction64 `json:"dnsaction64"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) CountDNSAction64() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsAction64URL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Items []struct {
			Count float64 `json:"__count"`
		} `json:"dnsaction64"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Items) > 0 {
		return result.Items[0].Count, nil
	}
	return 0, nil
}

// dnsaddrec
func (s *DNSService) AddDNSAddRec(rec models.DNSAddRec) error {
	payload := map[string]any{"dnsaddrec": rec}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsAddRecURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) DeleteDNSAddRec(hostname string) error {
	reqURL := fmt.Sprintf("%s/%s", dnsAddRecURL, url.QueryEscape(hostname))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) GetAllDNSAddRec() ([]models.DNSAddRec, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsAddRecURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSAddRec `json:"dnsaddrec"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) GetDNSAddRec(hostname string) ([]models.DNSAddRec, error) {
	reqURL := fmt.Sprintf("%s/%s", dnsAddRecURL, url.QueryEscape(hostname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSAddRec `json:"dnsaddrec"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) CountDNSAddRec() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsAddRecURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Items []struct {
			Count float64 `json:"__count"`
		} `json:"dnsaddrec"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Items) > 0 {
		return result.Items[0].Count, nil
	}
	return 0, nil
}

// dnscnamerec
func (s *DNSService) AddDNSCNameRec(rec models.DNSCnameRec) error {
	payload := map[string]any{"dnscnamerec": rec}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsCNameRecURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) DeleteDNSCNameRec(aliasName string) error {
	reqURL := fmt.Sprintf("%s/%s", dnsCNameRecURL, url.QueryEscape(aliasName))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) GetAllDNSCNameRec() ([]models.DNSCnameRec, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsCNameRecURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSCnameRec `json:"dnscnamerec"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) GetDNSCNameRec(aliasName string) ([]models.DNSCnameRec, error) {
	reqURL := fmt.Sprintf("%s/%s", dnsCNameRecURL, url.QueryEscape(aliasName))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSCnameRec `json:"dnscnamerec"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) CountDNSCNameRec() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsCNameRecURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Items []struct {
			Count float64 `json:"__count"`
		} `json:"dnscnamerec"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Items) > 0 {
		return result.Items[0].Count, nil
	}
	return 0, nil
}

// dnsglobal_binding
func (s *DNSService) GetDNSGlobalBinding() ([]models.DNSGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSGlobalBinding `json:"dnsglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

// dnsglobal_dnspolicy_binding
func (s *DNSService) AddDNSGlobalDNSPolicyBinding(binding models.DNSGlobalDNSPolicyBinding) error {
	payload := map[string]any{"dnsglobal_dnspolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, dnsGlobalDNSPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) DeleteDNSGlobalDNSPolicyBinding(policyname string, typefield string) error {
	reqURL := fmt.Sprintf("%s/%s?args=type:%s", dnsGlobalDNSPolicyBindingURL, url.QueryEscape(policyname), url.QueryEscape(typefield))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) GetDNSGlobalDNSPolicyBinding() ([]models.DNSGlobalDNSPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsGlobalDNSPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSGlobalDNSPolicyBinding `json:"dnsglobal_dnspolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) CountDNSGlobalDNSPolicyBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsGlobalDNSPolicyBindingURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Items []struct {
			Count float64 `json:"__count"`
		} `json:"dnsglobal_dnspolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Items) > 0 {
		return result.Items[0].Count, nil
	}
	return 0, nil
}

// dnskey
func (s *DNSService) AddDNSKey(key models.DNSKey) error {
	payload := map[string]any{"dnskey": key}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsKeyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) DeleteDNSKey(keyName string) error {
	reqURL := fmt.Sprintf("%s/%s", dnsKeyURL, url.QueryEscape(keyName))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) UpdateDNSKey(key models.DNSKey) error {
	payload := map[string]any{"dnskey": key}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, dnsKeyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) UnsetDNSKey(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"dnskey": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsKeyURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) GetAllDNSKey() ([]models.DNSKey, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsKeyURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSKey `json:"dnskey"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) GetDNSKey(keyName string) ([]models.DNSKey, error) {
	reqURL := fmt.Sprintf("%s/%s", dnsKeyURL, url.QueryEscape(keyName))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSKey `json:"dnskey"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) CountDNSKey() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsKeyURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Items []struct {
			Count float64 `json:"__count"`
		} `json:"dnskey"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Items) > 0 {
		return result.Items[0].Count, nil
	}
	return 0, nil
}

func (s *DNSService) CreateDNSKey(key models.DNSKey) error {
	payload := map[string]any{"dnskey": key}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsKeyURL+"?action=create", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) ImportDNSKey(key models.DNSKey) error {
	payload := map[string]any{"dnskey": key}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsKeyURL+"?action=import", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

// dnsmxrec
func (s *DNSService) AddDNSMXRec(rec models.DNSMXRec) error {
	payload := map[string]any{"dnsmxrec": rec}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsMXRecURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) DeleteDNSMXRec(domain string) error {
	reqURL := fmt.Sprintf("%s/%s", dnsMXRecURL, url.QueryEscape(domain))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) UpdateDNSMXRec(rec models.DNSMXRec) error {
	payload := map[string]any{"dnsmxrec": rec}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, dnsMXRecURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) UnsetDNSMXRec(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"dnsmxrec": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsMXRecURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) GetAllDNSMXRec() ([]models.DNSMXRec, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsMXRecURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSMXRec `json:"dnsmxrec"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) GetDNSMXRec(domain string) ([]models.DNSMXRec, error) {
	reqURL := fmt.Sprintf("%s/%s", dnsMXRecURL, url.QueryEscape(domain))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSMXRec `json:"dnsmxrec"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) CountDNSMXRec() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsMXRecURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Items []struct {
			Count float64 `json:"__count"`
		} `json:"dnsmxrec"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Items) > 0 {
		return result.Items[0].Count, nil
	}
	return 0, nil
}

// dnsnameserver
func (s *DNSService) AddDNSNameServer(server models.DNSNameServer) error {
	payload := map[string]any{"dnsnameserver": server}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsNameServerURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) DeleteDNSNameServer(ip string) error {
	reqURL := fmt.Sprintf("%s/%s", dnsNameServerURL, url.QueryEscape(ip))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) UpdateDNSNameServer(server models.DNSNameServer) error {
	payload := map[string]any{"dnsnameserver": server}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, dnsNameServerURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) UnsetDNSNameServer(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"dnsnameserver": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsNameServerURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) EnableDNSNameServer(ip string) error {
	payload := map[string]any{
		"dnsnameserver": map[string]string{"ip": ip},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsNameServerURL+"?action=enable", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) DisableDNSNameServer(ip string) error {
	payload := map[string]any{
		"dnsnameserver": map[string]string{"ip": ip},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsNameServerURL+"?action=disable", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) GetAllDNSNameServer() ([]models.DNSNameServer, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsNameServerURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSNameServer `json:"dnsnameserver"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) CountDNSNameServer() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsNameServerURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Items []struct {
			Count float64 `json:"__count"`
		} `json:"dnsnameserver"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Items) > 0 {
		return result.Items[0].Count, nil
	}
	return 0, nil
}

// dnsnaptrrec
func (s *DNSService) AddDNSNAPTRRec(rec models.DNSNaptrRec) error {
	payload := map[string]any{"dnsnaptrrec": rec}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsNAPTRRecURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) DeleteDNSNAPTRRec(domain string) error {
	reqURL := fmt.Sprintf("%s/%s", dnsNAPTRRecURL, url.QueryEscape(domain))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) GetAllDNSNAPTRRec() ([]models.DNSNaptrRec, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsNAPTRRecURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSNaptrRec `json:"dnsnaptrrec"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) GetDNSNAPTRRec(domain string) ([]models.DNSNaptrRec, error) {
	reqURL := fmt.Sprintf("%s/%s", dnsNAPTRRecURL, url.QueryEscape(domain))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSNaptrRec `json:"dnsnaptrrec"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) CountDNSNAPTRRec() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsNAPTRRecURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Items []struct {
			Count float64 `json:"__count"`
		} `json:"dnsnaptrrec"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Items) > 0 {
		return result.Items[0].Count, nil
	}
	return 0, nil
}

// dnsnsecrec
func (s *DNSService) GetAllDNSNSECRec() ([]models.DNSNSECRec, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsNSECRecURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSNSECRec `json:"dnsnsecrec"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) GetDNSNSEcRec(hostname string) ([]models.DNSNSECRec, error) {
	reqURL := fmt.Sprintf("%s/%s", dnsNSECRecURL, url.QueryEscape(hostname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSNSECRec `json:"dnsnsecrec"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) CountDNSNSECRec() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsNSECRecURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Items []struct {
			Count float64 `json:"__count"`
		} `json:"dnsnsecrec"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Items) > 0 {
		return result.Items[0].Count, nil
	}
	return 0, nil
}

// dnsnsrec
func (s *DNSService) AddDNSNSRec(rec models.DNSNSRec) error {
	payload := map[string]any{"dnsnsrec": rec}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsNSRecURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) DeleteDNSNSRec(domain string) error {
	reqURL := fmt.Sprintf("%s/%s", dnsNSRecURL, url.QueryEscape(domain))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) GetAllDNSNSRec() ([]models.DNSNSRec, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsNSRecURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSNSRec `json:"dnsnsrec"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) GetDNSNSRec(domain string) ([]models.DNSNSRec, error) {
	reqURL := fmt.Sprintf("%s/%s", dnsNSRecURL, url.QueryEscape(domain))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSNSRec `json:"dnsnsrec"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) CountDNSNSRec() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsNSRecURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Items []struct {
			Count float64 `json:"__count"`
		} `json:"dnsnsrec"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Items) > 0 {
		return result.Items[0].Count, nil
	}
	return 0, nil
}

// dnsparameter
func (s *DNSService) UpdateDNSParameter(param models.DNSParameter) error {
	payload := map[string]any{"dnsparameter": param}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, dnsParameterURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) UnsetDNSParameter(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"dnsparameter": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsParameterURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) GetAllDNSParameter() (models.DNSParameter, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsParameterURL, nil)
	if err != nil {
		return models.DNSParameter{}, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return models.DNSParameter{}, err
	}
	var result struct {
		Items []models.DNSParameter `json:"dnsparameter"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.DNSParameter{}, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Items) > 0 {
		return result.Items[0], nil
	}
	return models.DNSParameter{}, nil
}

// dnspolicy
func (s *DNSService) AddDNSPolicy(policy models.DNSPolicy) error {
	payload := map[string]any{"dnspolicy": policy}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) DeleteDNSPolicy(name string) error {
	reqURL := fmt.Sprintf("%s/%s", dnsPolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) UpdateDNSPolicy(policy models.DNSPolicy) error {
	payload := map[string]any{"dnspolicy": policy}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, dnsPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) UnsetDNSPolicy(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"dnspolicy": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsPolicyURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) GetAllDNSPolicy() ([]models.DNSPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsPolicyURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSPolicy `json:"dnspolicy"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) GetDNSPolicy(name string) ([]models.DNSPolicy, error) {
	reqURL := fmt.Sprintf("%s/%s", dnsPolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSPolicy `json:"dnspolicy"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) CountDNSPolicy() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsPolicyURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Items []struct {
			Count float64 `json:"__count"`
		} `json:"dnspolicy"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Items) > 0 {
		return result.Items[0].Count, nil
	}
	return 0, nil
}

// dnspolicy64
func (s *DNSService) AddDNSPolicy64(policy models.DNSPolicy64) error {
	payload := map[string]any{"dnspolicy64": policy}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsPolicy64URL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) DeleteDNSPolicy64(name string) error {
	reqURL := fmt.Sprintf("%s/%s", dnsPolicy64URL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) UpdateDNSPolicy64(policy models.DNSPolicy64) error {
	payload := map[string]any{"dnspolicy64": policy}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, dnsPolicy64URL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) GetAllDNSPolicy64() ([]models.DNSPolicy64, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsPolicy64URL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSPolicy64 `json:"dnspolicy64"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) GetDNSPolicy64(name string) ([]models.DNSPolicy64, error) {
	reqURL := fmt.Sprintf("%s/%s", dnsPolicy64URL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSPolicy64 `json:"dnspolicy64"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) CountDNSPolicy64() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsPolicy64URL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Items []struct {
			Count float64 `json:"__count"`
		} `json:"dnspolicy64"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Items) > 0 {
		return result.Items[0].Count, nil
	}
	return 0, nil
}

// dnspolicy64_binding
func (s *DNSService) GetAllDNSPolicy64Binding() ([]models.DNSPolicy64Binding, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsPolicy64BindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSPolicy64Binding `json:"dnspolicy64_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) GetDNSPolicy64Binding(name string) ([]models.DNSPolicy64Binding, error) {
	reqURL := fmt.Sprintf("%s/%s", dnsPolicy64BindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSPolicy64Binding `json:"dnspolicy64_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

// dnspolicy64_lbvserver_binding
func (s *DNSService) GetAllDNSPolicy64LBVServerBinding() ([]models.DNSPolicy64LBVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsPolicy64LBVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSPolicy64LBVServerBinding `json:"dnspolicy64_lbvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) GetDNSPolicy64LBVServerBinding(name string) ([]models.DNSPolicy64LBVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", dnsPolicy64LBVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSPolicy64LBVServerBinding `json:"dnspolicy64_lbvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) CountDNSPolicy64LBVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", dnsPolicy64LBVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Items []struct {
			Count float64 `json:"__count"`
		} `json:"dnspolicy64_lbvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Items) > 0 {
		return result.Items[0].Count, nil
	}
	return 0, nil
}

// dnspolicylabel
func (s *DNSService) AddDNSPolicyLabel(label models.DNSPolicyLabel) error {
	payload := map[string]any{"dnspolicylabel": label}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsPolicyLabelURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) DeleteDNSPolicyLabel(labelName string) error {
	reqURL := fmt.Sprintf("%s/%s", dnsPolicyLabelURL, url.QueryEscape(labelName))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) GetAllDNSPolicyLabel() ([]models.DNSPolicyLabel, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsPolicyLabelURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSPolicyLabel `json:"dnspolicylabel"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) GetDNSPolicyLabel(labelName string) ([]models.DNSPolicyLabel, error) {
	reqURL := fmt.Sprintf("%s/%s", dnsPolicyLabelURL, url.QueryEscape(labelName))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSPolicyLabel `json:"dnspolicylabel"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) CountDNSPolicyLabel() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsPolicyLabelURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Items []struct {
			Count float64 `json:"__count"`
		} `json:"dnspolicylabel"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Items) > 0 {
		return result.Items[0].Count, nil
	}
	return 0, nil
}

func (s *DNSService) RenameDNSPolicyLabel(name string, newName string) error {
	payload := map[string]any{
		"dnspolicylabel": map[string]string{
			"labelname": name,
			"newname":   newName,
		},
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsPolicyLabelURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

// dnspolicylabel_binding
func (s *DNSService) GetAllDNSPolicyLabelBinding() ([]models.DNSPolicyLabelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsPolicyLabelBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSPolicyLabelBinding `json:"dnspolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) GetDNSPolicyLabelBinding(labelName string) ([]models.DNSPolicyLabelBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", dnsPolicyLabelBindingURL, url.QueryEscape(labelName))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSPolicyLabelBinding `json:"dnspolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

// dnspolicylabel_dnspolicy_binding
func (s *DNSService) AddDNSPolicyLabelDNSPolicyBinding(binding models.DNSPolicyLabelDNSPolicyBinding) error {
	payload := map[string]any{"dnspolicylabel_dnspolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, dnsPolicyLabelDNSPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) DeleteDNSPolicyLabelDNSPolicyBinding(labelName string, policyName string) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s", dnsPolicyLabelDNSPolicyBindingURL, url.QueryEscape(labelName), url.QueryEscape(policyName))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) GetAllDNSPolicyLabelDNSPolicyBinding() ([]models.DNSPolicyLabelDNSPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsPolicyLabelDNSPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSPolicyLabelDNSPolicyBinding `json:"dnspolicylabel_dnspolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) GetDNSPolicyLabelDNSPolicyBinding(labelName string) ([]models.DNSPolicyLabelDNSPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", dnsPolicyLabelDNSPolicyBindingURL, url.QueryEscape(labelName))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSPolicyLabelDNSPolicyBinding `json:"dnspolicylabel_dnspolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) CountDNSPolicyLabelDNSPolicyBinding(labelName string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", dnsPolicyLabelDNSPolicyBindingURL, url.QueryEscape(labelName))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Items []struct {
			Count float64 `json:"__count"`
		} `json:"dnspolicylabel_dnspolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Items) > 0 {
		return result.Items[0].Count, nil
	}
	return 0, nil
}

// dnspolicylabel_policybinding_binding
func (s *DNSService) GetAllDNSPolicyLabelPolicyBindingBinding() ([]models.DNSPolicyLabelPolicyBindingBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsPolicyLabelPolicyBindingBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSPolicyLabelPolicyBindingBinding `json:"dnspolicylabel_policybinding_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) GetDNSPolicyLabelPolicyBindingBinding(labelName string) ([]models.DNSPolicyLabelPolicyBindingBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", dnsPolicyLabelPolicyBindingBindingURL, url.QueryEscape(labelName))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSPolicyLabelPolicyBindingBinding `json:"dnspolicylabel_policybinding_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) CountDNSPolicyLabelPolicyBindingBinding(labelName string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", dnsPolicyLabelPolicyBindingBindingURL, url.QueryEscape(labelName))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Items []struct {
			Count float64 `json:"__count"`
		} `json:"dnspolicylabel_policybinding_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Items) > 0 {
		return result.Items[0].Count, nil
	}
	return 0, nil
}

// dnspolicy_binding
func (s *DNSService) GetAllDNSPolicyBinding() ([]models.DNSPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSPolicyBinding `json:"dnspolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) GetDNSPolicyBinding(name string) ([]models.DNSPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", dnsPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSPolicyBinding `json:"dnspolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

// dnspolicy_dnsglobal_binding
func (s *DNSService) GetAllDNSPolicyDNSGlobalBinding() ([]models.DNSPolicyDNSGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsPolicyDNSGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSPolicyDNSGlobalBinding `json:"dnspolicy_dnsglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) GetDNSPolicyDNSGlobalBinding(name string) ([]models.DNSPolicyDNSGlobalBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", dnsPolicyDNSGlobalBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSPolicyDNSGlobalBinding `json:"dnspolicy_dnsglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) CountDNSPolicyDNSGlobalBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", dnsPolicyDNSGlobalBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Items []struct {
			Count float64 `json:"__count"`
		} `json:"dnspolicy_dnsglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Items) > 0 {
		return result.Items[0].Count, nil
	}
	return 0, nil
}

// dnspolicy_dnspolicylabel_binding
func (s *DNSService) GetAllDNSPolicyDNSPolicyLabelBinding() ([]models.DNSPolicyDNSPolicyLabelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsPolicyDNSPolicyLabelBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSPolicyDNSPolicyLabelBinding `json:"dnspolicy_dnspolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) GetDNSPolicyDNSPolicyLabelBinding(name string) ([]models.DNSPolicyDNSPolicyLabelBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", dnsPolicyDNSPolicyLabelBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSPolicyDNSPolicyLabelBinding `json:"dnspolicy_dnspolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) CountDNSPolicyDNSPolicyLabelBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", dnsPolicyDNSPolicyLabelBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Items []struct {
			Count float64 `json:"__count"`
		} `json:"dnspolicy_dnspolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Items) > 0 {
		return result.Items[0].Count, nil
	}
	return 0, nil
}

// dnsprofile
func (s *DNSService) AddDNSProfile(profile models.DNSProfile) error {
	payload := map[string]any{"dnsprofile": profile}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) DeleteDNSProfile(name string) error {
	reqURL := fmt.Sprintf("%s/%s", dnsProfileURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) UpdateDNSProfile(profile models.DNSProfile) error {
	payload := map[string]any{"dnsprofile": profile}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, dnsProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) UnsetDNSProfile(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"dnsprofile": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsProfileURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) GetAllDNSProfile() ([]models.DNSProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsProfileURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSProfile `json:"dnsprofile"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) GetDNSProfile(name string) ([]models.DNSProfile, error) {
	reqURL := fmt.Sprintf("%s/%s", dnsProfileURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSProfile `json:"dnsprofile"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) CountDNSProfile() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsProfileURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Items []struct {
			Count float64 `json:"__count"`
		} `json:"dnsprofile"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Items) > 0 {
		return result.Items[0].Count, nil
	}
	return 0, nil
}

// dnsproxyrecords
func (s *DNSService) FlushDNSProxyRecords() error {
	req, err := s.client.NewRequest(http.MethodPost, dnsProxyRecordsURL+"?action=flush", nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

// dnsptrrec
func (s *DNSService) AddDNSPTRRec(rec models.DNSPtrRec) error {
	payload := map[string]any{"dnsptrrec": rec}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsPTRRecURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) DeleteDNSPTRRec(reverseDomain string) error {
	reqURL := fmt.Sprintf("%s/%s", dnsPTRRecURL, url.QueryEscape(reverseDomain))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) GetAllDNSPTRRec() ([]models.DNSPtrRec, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsPTRRecURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSPtrRec `json:"dnsptrrec"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) GetDNSPTRRec(reverseDomain string) ([]models.DNSPtrRec, error) {
	reqURL := fmt.Sprintf("%s/%s", dnsPTRRecURL, url.QueryEscape(reverseDomain))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSPtrRec `json:"dnsptrrec"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) CountDNSPTRRec() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsPTRRecURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Items []struct {
			Count float64 `json:"__count"`
		} `json:"dnsptrrec"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Items) > 0 {
		return result.Items[0].Count, nil
	}
	return 0, nil
}

// dnssoarec
func (s *DNSService) AddDNSSOARec(rec models.DNSSOARec) error {
	payload := map[string]any{"dnssoarec": rec}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsSOARecURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) DeleteDNSSOARec(domain string) error {
	reqURL := fmt.Sprintf("%s/%s", dnsSOARecURL, url.QueryEscape(domain))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) UpdateDNSSOARec(rec models.DNSSOARec) error {
	payload := map[string]any{"dnssoarec": rec}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, dnsSOARecURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) UnsetDNSSOARec(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"dnssoarec": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsSOARecURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) GetAllDNSSOARec() ([]models.DNSSOARec, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsSOARecURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSSOARec `json:"dnssoarec"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) GetDNSSOARec(domain string) ([]models.DNSSOARec, error) {
	reqURL := fmt.Sprintf("%s/%s", dnsSOARecURL, url.QueryEscape(domain))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSSOARec `json:"dnssoarec"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) CountDNSSOARec() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsSOARecURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Items []struct {
			Count float64 `json:"__count"`
		} `json:"dnssoarec"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Items) > 0 {
		return result.Items[0].Count, nil
	}
	return 0, nil
}

// dnssrvrec
func (s *DNSService) AddDNSSRVRec(rec models.DNSSrvRec) error {
	payload := map[string]any{"dnssrvrec": rec}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsSRVRecURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) DeleteDNSSRVRec(domain string) error {
	reqURL := fmt.Sprintf("%s/%s", dnsSRVRecURL, url.QueryEscape(domain))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) UpdateDNSSRVRec(rec models.DNSSrvRec) error {
	payload := map[string]any{"dnssrvrec": rec}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, dnsSRVRecURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) UnsetDNSSRVRec(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"dnssrvrec": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsSRVRecURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) GetAllDNSSRVRec() ([]models.DNSSrvRec, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsSRVRecURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSSrvRec `json:"dnssrvrec"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) CountDNSSRVRec() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsSRVRecURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Items []struct {
			Count float64 `json:"__count"`
		} `json:"dnssrvrec"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Items) > 0 {
		return result.Items[0].Count, nil
	}
	return 0, nil
}

// dnssubnetcache
func (s *DNSService) FlushDNSSubnetCache() error {
	req, err := s.client.NewRequest(http.MethodPost, dnsSubnetCacheURL+"?action=flush", nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) GetAllDNSSubnetCache() ([]models.DNSSubnetCache, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsSubnetCacheURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSSubnetCache `json:"dnssubnetcache"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) GetDNSSubnetCache(hostname string) ([]models.DNSSubnetCache, error) {
	reqURL := fmt.Sprintf("%s/%s", dnsSubnetCacheURL, url.QueryEscape(hostname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSSubnetCache `json:"dnssubnetcache"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) CountDNSSubnetCache() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsSubnetCacheURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Items []struct {
			Count float64 `json:"__count"`
		} `json:"dnssubnetcache"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Items) > 0 {
		return result.Items[0].Count, nil
	}
	return 0, nil
}

// dnssuffix
func (s *DNSService) AddDNSSuffix(suffix models.DNSSuffix) error {
	payload := map[string]any{"dnssuffix": suffix}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsSuffixURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) DeleteDNSSuffix(suffix string) error {
	reqURL := fmt.Sprintf("%s/%s", dnsSuffixURL, url.QueryEscape(suffix))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) GetAllDNSSuffix() ([]models.DNSSuffix, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsSuffixURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSSuffix `json:"dnssuffix"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) GetDNSSuffix(suffix string) ([]models.DNSSuffix, error) {
	reqURL := fmt.Sprintf("%s/%s", dnsSuffixURL, url.QueryEscape(suffix))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSSuffix `json:"dnssuffix"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) CountDNSSuffix() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsSuffixURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Items []struct {
			Count float64 `json:"__count"`
		} `json:"dnssuffix"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Items) > 0 {
		return result.Items[0].Count, nil
	}
	return 0, nil
}

// dnstxtrec
func (s *DNSService) AddDNSTXTRec(rec models.DNSTxtRec) error {
	payload := map[string]any{"dnstxtrec": rec}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsTXTRecURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) DeleteDNSTXTRec(domain string) error {
	reqURL := fmt.Sprintf("%s/%s", dnsTXTRecURL, url.QueryEscape(domain))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) GetAllDNSTXTRec() ([]models.DNSTxtRec, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsTXTRecURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSTxtRec `json:"dnstxtrec"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) GetDNSTXTRec(domain string) ([]models.DNSTxtRec, error) {
	reqURL := fmt.Sprintf("%s/%s", dnsTXTRecURL, url.QueryEscape(domain))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSTxtRec `json:"dnstxtrec"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) CountDNSTXTRec() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsTXTRecURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Items []struct {
			Count float64 `json:"__count"`
		} `json:"dnstxtrec"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Items) > 0 {
		return result.Items[0].Count, nil
	}
	return 0, nil
}

// dnsview
func (s *DNSService) AddDNSView(view models.DNSView) error {
	payload := map[string]any{"dnsview": view}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsViewURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) DeleteDNSView(viewName string) error {
	reqURL := fmt.Sprintf("%s/%s", dnsViewURL, url.QueryEscape(viewName))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) GetAllDNSView() ([]models.DNSView, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsViewURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSView `json:"dnsview"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) GetDNSView(viewName string) ([]models.DNSView, error) {
	reqURL := fmt.Sprintf("%s/%s", dnsViewURL, url.QueryEscape(viewName))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSView `json:"dnsview"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) CountDNSView() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsViewURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Items []struct {
			Count float64 `json:"__count"`
		} `json:"dnsview"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Items) > 0 {
		return result.Items[0].Count, nil
	}
	return 0, nil
}

// dnsview_binding
func (s *DNSService) GetAllDNSViewBinding() ([]models.DNSViewBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsViewBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSViewBinding `json:"dnsview_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) GetDNSViewBinding(viewName string) ([]models.DNSViewBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", dnsViewBindingURL, url.QueryEscape(viewName))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSViewBinding `json:"dnsview_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

// dnsview_dnspolicy_binding
func (s *DNSService) GetAllDNSViewDNSPolicyBinding() ([]models.DNSViewDNSPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsViewDNSPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSViewDNSPolicyBinding `json:"dnsview_dnspolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) GetDNSViewDNSPolicyBinding(viewName string) ([]models.DNSViewDNSPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", dnsViewDNSPolicyBindingURL, url.QueryEscape(viewName))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSViewDNSPolicyBinding `json:"dnsview_dnspolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) CountDNSViewDNSPolicyBinding(viewName string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", dnsViewDNSPolicyBindingURL, url.QueryEscape(viewName))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Items []struct {
			Count float64 `json:"__count"`
		} `json:"dnsview_dnspolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Items) > 0 {
		return result.Items[0].Count, nil
	}
	return 0, nil
}

// dnsview_gslbservice_binding
func (s *DNSService) GetAllDNSViewGSLBServiceBinding() ([]models.DNSViewGSLBServiceBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsViewGSLBServiceBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSViewGSLBServiceBinding `json:"dnsview_gslbservice_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) GetDNSViewGSLBServiceBinding(viewName string) ([]models.DNSViewGSLBServiceBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", dnsViewGSLBServiceBindingURL, url.QueryEscape(viewName))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSViewGSLBServiceBinding `json:"dnsview_gslbservice_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) CountDNSViewGSLBServiceBinding(viewName string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", dnsViewGSLBServiceBindingURL, url.QueryEscape(viewName))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Items []struct {
			Count float64 `json:"__count"`
		} `json:"dnsview_gslbservice_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Items) > 0 {
		return result.Items[0].Count, nil
	}
	return 0, nil
}

// dnszone
func (s *DNSService) AddDNSZone(zone models.DNSZone) error {
	payload := map[string]any{"dnszone": zone}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsZoneURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) DeleteDNSZone(zoneName string) error {
	reqURL := fmt.Sprintf("%s/%s", dnsZoneURL, url.QueryEscape(zoneName))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) UpdateDNSZone(zone models.DNSZone) error {
	payload := map[string]any{"dnszone": zone}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, dnsZoneURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) UnsetDNSZone(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"dnszone": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsZoneURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) SignDNSZone(zone models.DNSZone) error {
	payload := map[string]any{"dnszone": zone}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsZoneURL+"?action=sign", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) UnsignDNSZone(zone models.DNSZone) error {
	payload := map[string]any{"dnszone": zone}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, dnsZoneURL+"?action=unsign", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *DNSService) GetAllDNSZone() ([]models.DNSZone, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsZoneURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSZone `json:"dnszone"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) GetDNSZone(zoneName string) ([]models.DNSZone, error) {
	reqURL := fmt.Sprintf("%s/%s", dnsZoneURL, url.QueryEscape(zoneName))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSZone `json:"dnszone"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) CountDNSZone() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsZoneURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Items []struct {
			Count float64 `json:"__count"`
		} `json:"dnszone"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Items) > 0 {
		return result.Items[0].Count, nil
	}
	return 0, nil
}

// dnszone_binding
func (s *DNSService) GetAllDNSZoneBinding() ([]models.DNSZoneBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsZoneBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSZoneBinding `json:"dnszone_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) GetDNSZoneBinding(zoneName string) ([]models.DNSZoneBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", dnsZoneBindingURL, url.QueryEscape(zoneName))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSZoneBinding `json:"dnszone_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

// dnszone_dnskey_binding
func (s *DNSService) GetAllDNSZoneDNSKeyBinding() ([]models.DNSZoneDNSKeyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsZoneDNSKeyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSZoneDNSKeyBinding `json:"dnszone_dnskey_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) GetDNSZoneDNSKeyBinding(zoneName string) ([]models.DNSZoneDNSKeyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", dnsZoneDNSKeyBindingURL, url.QueryEscape(zoneName))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSZoneDNSKeyBinding `json:"dnszone_dnskey_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) CountDNSZoneDNSKeyBinding(zoneName string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", dnsZoneDNSKeyBindingURL, url.QueryEscape(zoneName))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Items []struct {
			Count float64 `json:"__count"`
		} `json:"dnszone_dnskey_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Items) > 0 {
		return result.Items[0].Count, nil
	}
	return 0, nil
}

// dnszone_domain_binding
func (s *DNSService) GetAllDNSZoneDomainBinding() ([]models.DNSZoneDomainBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, dnsZoneDomainBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSZoneDomainBinding `json:"dnszone_domain_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) GetDNSZoneDomainBinding(zoneName string) ([]models.DNSZoneDomainBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", dnsZoneDomainBindingURL, url.QueryEscape(zoneName))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.DNSZoneDomainBinding `json:"dnszone_domain_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *DNSService) CountDNSZoneDomainBinding(zoneName string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", dnsZoneDomainBindingURL, url.QueryEscape(zoneName))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	var result struct {
		Items []struct {
			Count float64 `json:"__count"`
		} `json:"dnszone_domain_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Items) > 0 {
		return result.Items[0].Count, nil
	}
	return 0, nil
}
