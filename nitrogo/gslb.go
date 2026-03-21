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
	gslbConfigURL                                           = "/nitro/v1/config/gslbconfig"
	gslbDomainURL                                           = "/nitro/v1/config/gslbdomain"
	gslbDomainBindingURL                                    = "/nitro/v1/config/gslbdomain_binding"
	gslbDomainGSLBServiceGroupMemberBindingURL              = "/nitro/v1/config/gslbdomain_gslbservicegroupmember_binding"
	gslbDomainGSLBServiceGroupBindingURL                    = "/nitro/v1/config/gslbdomain_gslbservicegroup_binding"
	gslbDomainGSLBServiceBindingURL                         = "/nitro/v1/config/gslbdomain_gslbservice_binding"
	gslbDomainGSLBVServerBindingURL                         = "/nitro/v1/config/gslbdomain_gslbvserver_binding"
	gslbDomainLBMonitorBindingURL                           = "/nitro/v1/config/gslbdomain_lbmonitor_binding"
	gslbDNSEntriesURL                                       = "/nitro/v1/config/gslbdnsentries"
	gslbDNSEntryURL                                         = "/nitro/v1/config/gslbdnsentry"
	gslbParameterURL                                        = "/nitro/v1/config/gslbparameter"
	gslbRunningConfigURL                                    = "/nitro/v1/config/gslbrunningconfig"
	gslbServiceURL                                          = "/nitro/v1/config/gslbservice"
	gslbServiceGroupURL                                     = "/nitro/v1/config/gslbservicegroup"
	gslbServiceGroupBindingURL                              = "/nitro/v1/config/gslbservicegroup_binding"
	gslbServiceGroupGSLBServiceGroupMemberBindingURL        = "/nitro/v1/config/gslbservicegroup_gslbservicegroupmember_binding"
	gslbServiceGroupLBMonitorBindingURL                     = "/nitro/v1/config/gslbservicegroup_lbmonitor_binding"
	gslbServiceGroupServiceGroupEntityMonBindingsBindingURL = "/nitro/v1/config/gslbservicegroup_servicegroupentitymonbindings_binding"
	gslbServiceBindingURL                                   = "/nitro/v1/config/gslbservice_binding"
	gslbServiceDNSViewBindingURL                            = "/nitro/v1/config/gslbservice_dnsview_binding"
	gslbServiceLBMonitorBindingURL                          = "/nitro/v1/config/gslbservice_lbmonitor_binding"
	gslbSiteURL                                             = "/nitro/v1/config/gslbsite"
	gslbSiteBindingURL                                      = "/nitro/v1/config/gslbsite_binding"
	gslbSiteGSLBServiceGroupMemberBindingURL                = "/nitro/v1/config/gslbsite_gslbservicegroupmember_binding"
	gslbSiteGSLBServiceGroupBindingURL                      = "/nitro/v1/config/gslbsite_gslbservicegroup_binding"
	gslbSiteGSLBServiceBindingURL                           = "/nitro/v1/config/gslbsite_gslbservice_binding"
	gslbSyncStatusURL                                       = "/nitro/v1/config/gslbsyncstatus"
	gslbVServerURL                                          = "/nitro/v1/config/gslbvserver"
	gslbVServerBindingURL                                   = "/nitro/v1/config/gslbvserver_binding"
	gslbVServerDomainBindingURL                             = "/nitro/v1/config/gslbvserver_domain_binding"
	gslbVServerGSLBServiceGroupMemberBindingURL             = "/nitro/v1/config/gslbvserver_gslbservicegroupmember_binding"
	gslbVServerGSLBServiceGroupBindingURL                   = "/nitro/v1/config/gslbvserver_gslbservicegroup_binding"
	gslbVServerGSLBServiceBindingURL                        = "/nitro/v1/config/gslbvserver_gslbservice_binding"
	gslbVServerSpilloverPolicyBindingURL                    = "/nitro/v1/config/gslbvserver_spilloverpolicy_binding"
)

// Global Server Load Balancing (GSLB) configuration. GSLB feature ensures that client requests are
// directed to a best performing site available in a global enterprise and distributed Internet environment.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/gslb/gslb
type GSLBService struct {
	client *Client
}

// gslbconfig
func (s *GSLBService) SyncGSLBConfig(config models.GSLBConfig) error {
	payload := map[string]any{
		"gslbconfig": config,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, gslbConfigURL+"?action=sync", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// gslbdomain
func (s *GSLBService) GetAllGSLBDomain() ([]models.GSLBDomain, error) {
	req, err := s.client.NewRequest(http.MethodGet, gslbDomainURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		GSLBDomains []models.GSLBDomain `json:"gslbdomain"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.GSLBDomains, nil
}

func (s *GSLBService) GetGSLBDomain(name string) (*models.GSLBDomain, error) {
	reqURL := fmt.Sprintf("%s/%s", gslbDomainURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		GSLBDomains []models.GSLBDomain `json:"gslbdomain"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.GSLBDomains) == 0 {
		return nil, fmt.Errorf("gslbdomain not found")
	}

	return &result.GSLBDomains[0], nil
}

func (s *GSLBService) CountGSLBDomain() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, gslbDomainURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		GSLBDomains []struct {
			Count float64 `json:"__count"`
		} `json:"gslbdomain"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.GSLBDomains) > 0 {
		return result.GSLBDomains[0].Count, nil
	}

	return 0, nil
}

// gslbdomain_binding
func (s *GSLBService) GetAllGSLBDomainBinding() ([]models.GSLBDomainBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, gslbDomainBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBDomainBinding `json:"gslbdomain_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) GetGSLBDomainBinding(name string) (*models.GSLBDomainBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", gslbDomainBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBDomainBinding `json:"gslbdomain_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) == 0 {
		return nil, fmt.Errorf("gslbdomain_binding not found")
	}

	return &result.Bindings[0], nil
}

// gslbdomain_gslbservicegroupmember_binding
func (s *GSLBService) GetAllGSLBDomainGSLBServiceGroupMemberBinding() ([]models.GSLBDomainGSLBServiceGroupMemberBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, gslbDomainGSLBServiceGroupMemberBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBDomainGSLBServiceGroupMemberBinding `json:"gslbdomain_gslbservicegroupmember_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) GetGSLBDomainGSLBServiceGroupMemberBinding(name string) ([]models.GSLBDomainGSLBServiceGroupMemberBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", gslbDomainGSLBServiceGroupMemberBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBDomainGSLBServiceGroupMemberBinding `json:"gslbdomain_gslbservicegroupmember_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) CountGSLBDomainGSLBServiceGroupMemberBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", gslbDomainGSLBServiceGroupMemberBindingURL, url.PathEscape(name))
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
		} `json:"gslbdomain_gslbservicegroupmember_binding"`
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

// gslbdomain_gslbservicegroup_binding
func (s *GSLBService) GetAllGSLBDomainGSLBServiceGroupBinding() ([]models.GSLBDomainGSLBServiceGroupBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, gslbDomainGSLBServiceGroupBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBDomainGSLBServiceGroupBinding `json:"gslbdomain_gslbservicegroup_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) GetGSLBDomainGSLBServiceGroupBinding(name string) ([]models.GSLBDomainGSLBServiceGroupBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", gslbDomainGSLBServiceGroupBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBDomainGSLBServiceGroupBinding `json:"gslbdomain_gslbservicegroup_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) CountGSLBDomainGSLBServiceGroupBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", gslbDomainGSLBServiceGroupBindingURL, url.PathEscape(name))
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
		} `json:"gslbdomain_gslbservicegroup_binding"`
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

// gslbdomain_gslbservice_binding
func (s *GSLBService) GetAllGSLBDomainGSLBServiceBinding() ([]models.GSLBDomainGSLBServiceBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, gslbDomainGSLBServiceBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBDomainGSLBServiceBinding `json:"gslbdomain_gslbservice_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) GetGSLBDomainGSLBServiceBinding(name string) ([]models.GSLBDomainGSLBServiceBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", gslbDomainGSLBServiceBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBDomainGSLBServiceBinding `json:"gslbdomain_gslbservice_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) CountGSLBDomainGSLBServiceBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", gslbDomainGSLBServiceBindingURL, url.PathEscape(name))
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
		} `json:"gslbdomain_gslbservice_binding"`
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

// gslbdomain_gslbvserver_binding
func (s *GSLBService) GetAllGSLBDomainGSLBVServerBinding() ([]models.GSLBDomainGSLBVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, gslbDomainGSLBVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBDomainGSLBVServerBinding `json:"gslbdomain_gslbvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) GetGSLBDomainGSLBVServerBinding(name string) ([]models.GSLBDomainGSLBVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", gslbDomainGSLBVServerBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBDomainGSLBVServerBinding `json:"gslbdomain_gslbvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) CountGSLBDomainGSLBVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", gslbDomainGSLBVServerBindingURL, url.PathEscape(name))
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
		} `json:"gslbdomain_gslbvserver_binding"`
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

// gslbdomain_lbmonitor_binding
func (s *GSLBService) GetAllGSLBDomainLBMonitorBinding() ([]models.GSLBDomainLBMonitorBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, gslbDomainLBMonitorBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBDomainLBMonitorBinding `json:"gslbdomain_lbmonitor_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) GetGSLBDomainLBMonitorBinding(name string) ([]models.GSLBDomainLBMonitorBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", gslbDomainLBMonitorBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBDomainLBMonitorBinding `json:"gslbdomain_lbmonitor_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) CountGSLBDomainLBMonitorBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", gslbDomainLBMonitorBindingURL, url.PathEscape(name))
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
		} `json:"gslbdomain_lbmonitor_binding"`
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

// gslbdnsentries
func (s *GSLBService) ClearGSLBDNSEntries(resource models.GSLBLDNSEntries) error {
	payload := map[string]any{
		"gslbdnsentries": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, gslbDNSEntriesURL+"?action=clear", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) GetAllGSLBDNSEntries() ([]models.GSLBLDNSEntries, error) {
	req, err := s.client.NewRequest(http.MethodGet, gslbDNSEntriesURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Entries []models.GSLBLDNSEntries `json:"gslbdnsentries"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Entries, nil
}

func (s *GSLBService) CountGSLBDNSEntries() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, gslbDNSEntriesURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Entries []struct {
			Count float64 `json:"__count"`
		} `json:"gslbdnsentries"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Entries) > 0 {
		return result.Entries[0].Count, nil
	}

	return 0, nil
}

// gslbdnsentry
func (s *GSLBService) DeleteGSLBDNSEntry(name string) error {
	reqURL := fmt.Sprintf("%s/%s", gslbDNSEntryURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// gslbparameter
func (s *GSLBService) UpdateGSLBParameter(resource models.GSLBParameter) error {
	payload := map[string]any{
		"gslbparameter": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, gslbParameterURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) UnsetGSLBParameter(resource models.GSLBParameter) error {
	payload := map[string]any{
		"gslbparameter": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, gslbParameterURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) GetAllGSLBParameter() (*models.GSLBParameter, error) {
	req, err := s.client.NewRequest(http.MethodGet, gslbParameterURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Parameter []models.GSLBParameter `json:"gslbparameter"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Parameter) == 0 {
		return nil, fmt.Errorf("gslbparameter not found")
	}

	return &result.Parameter[0], nil
}

// gslbrunningconfig
func (s *GSLBService) GetAllGSLBRunningConfig() ([]models.GSLBRunningConfig, error) {
	req, err := s.client.NewRequest(http.MethodGet, gslbRunningConfigURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		RunningConfig []models.GSLBRunningConfig `json:"gslbrunningconfig"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.RunningConfig, nil
}

// gslbservice
func (s *GSLBService) AddGSLBService(resource models.GSLBService) error {
	payload := map[string]any{
		"gslbservice": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, gslbServiceURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) DeleteGSLBService(name string) error {
	reqURL := fmt.Sprintf("%s/%s", gslbServiceURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) UpdateGSLBService(resource models.GSLBService) error {
	payload := map[string]any{
		"gslbservice": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, gslbServiceURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) UnsetGSLBService(resource models.GSLBService) error {
	payload := map[string]any{
		"gslbservice": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, gslbServiceURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) GetAllGSLBService() ([]models.GSLBService, error) {
	req, err := s.client.NewRequest(http.MethodGet, gslbServiceURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Services []models.GSLBService `json:"gslbservice"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Services, nil
}

func (s *GSLBService) GetGSLBService(name string) (*models.GSLBService, error) {
	reqURL := fmt.Sprintf("%s/%s", gslbServiceURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Services []models.GSLBService `json:"gslbservice"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Services) == 0 {
		return nil, fmt.Errorf("gslbservice not found")
	}

	return &result.Services[0], nil
}

func (s *GSLBService) CountGSLBService() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, gslbServiceURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Services []struct {
			Count float64 `json:"__count"`
		} `json:"gslbservice"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Services) > 0 {
		return result.Services[0].Count, nil
	}

	return 0, nil
}

func (s *GSLBService) RenameGSLBService(resource models.GSLBService) error {
	payload := map[string]any{
		"gslbservice": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, gslbServiceURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// gslbservicegroup
func (s *GSLBService) AddGSLBServiceGroup(resource models.GSLBServiceGroup) error {
	payload := map[string]any{
		"gslbservicegroup": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, gslbServiceGroupURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) DeleteGSLBServiceGroup(name string) error {
	reqURL := fmt.Sprintf("%s/%s", gslbServiceGroupURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) UpdateGSLBServiceGroup(resource models.GSLBServiceGroup) error {
	payload := map[string]any{
		"gslbservicegroup": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, gslbServiceGroupURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) UnsetGSLBServiceGroup(resource models.GSLBServiceGroup) error {
	payload := map[string]any{
		"gslbservicegroup": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, gslbServiceGroupURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) EnableGSLBServiceGroup(resource models.GSLBServiceGroup) error {
	payload := map[string]any{
		"gslbservicegroup": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, gslbServiceGroupURL+"?action=enable", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) DisableGSLBServiceGroup(resource models.GSLBServiceGroup) error {
	payload := map[string]any{
		"gslbservicegroup": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, gslbServiceGroupURL+"?action=disable", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) GetAllGSLBServiceGroup() ([]models.GSLBServiceGroup, error) {
	req, err := s.client.NewRequest(http.MethodGet, gslbServiceGroupURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Groups []models.GSLBServiceGroup `json:"gslbservicegroup"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Groups, nil
}

func (s *GSLBService) GetGSLBServiceGroup(name string) (*models.GSLBServiceGroup, error) {
	reqURL := fmt.Sprintf("%s/%s", gslbServiceGroupURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Groups []models.GSLBServiceGroup `json:"gslbservicegroup"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Groups) == 0 {
		return nil, fmt.Errorf("gslbservicegroup not found")
	}

	return &result.Groups[0], nil
}

func (s *GSLBService) CountGSLBServiceGroup() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, gslbServiceGroupURL+"?count=yes", nil)
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
		} `json:"gslbservicegroup"`
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

func (s *GSLBService) RenameGSLBServiceGroup(resource models.GSLBServiceGroup) error {
	payload := map[string]any{
		"gslbservicegroup": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, gslbServiceGroupURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// gslbservicegroup_binding
func (s *GSLBService) GetAllGSLBServiceGroupBinding() ([]models.GSLBServiceGroupBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, gslbServiceGroupBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBServiceGroupBinding `json:"gslbservicegroup_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) GetGSLBServiceGroupBinding(name string) (*models.GSLBServiceGroupBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", gslbServiceGroupBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBServiceGroupBinding `json:"gslbservicegroup_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) == 0 {
		return nil, fmt.Errorf("gslbservicegroup_binding not found")
	}

	return &result.Bindings[0], nil
}

// gslbservicegroup_gslbservicegroupmember_binding
func (s *GSLBService) AddGSLBServiceGroupGSLBServiceGroupMemberBinding(resource models.GSLBServiceGroupGSLBServiceGroupMemberBinding) error {
	payload := map[string]any{
		"gslbservicegroup_gslbservicegroupmember_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, gslbServiceGroupGSLBServiceGroupMemberBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) DeleteGSLBServiceGroupGSLBServiceGroupMemberBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", gslbServiceGroupGSLBServiceGroupMemberBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) GetAllGSLBServiceGroupGSLBServiceGroupMemberBinding() ([]models.GSLBServiceGroupGSLBServiceGroupMemberBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, gslbServiceGroupGSLBServiceGroupMemberBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBServiceGroupGSLBServiceGroupMemberBinding `json:"gslbservicegroup_gslbservicegroupmember_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) GetGSLBServiceGroupGSLBServiceGroupMemberBinding(name string) ([]models.GSLBServiceGroupGSLBServiceGroupMemberBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", gslbServiceGroupGSLBServiceGroupMemberBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBServiceGroupGSLBServiceGroupMemberBinding `json:"gslbservicegroup_gslbservicegroupmember_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) CountGSLBServiceGroupGSLBServiceGroupMemberBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", gslbServiceGroupGSLBServiceGroupMemberBindingURL, url.PathEscape(name))
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
		} `json:"gslbservicegroup_gslbservicegroupmember_binding"`
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

// gslbservicegroup_lbmonitor_binding
func (s *GSLBService) AddGSLBServiceGroupLBMonitorBinding(resource models.GSLBServiceGroupLBMonitorBinding) error {
	payload := map[string]any{
		"gslbservicegroup_lbmonitor_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, gslbServiceGroupLBMonitorBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) DeleteGSLBServiceGroupLBMonitorBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", gslbServiceGroupLBMonitorBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) GetAllGSLBServiceGroupLBMonitorBinding() ([]models.GSLBServiceGroupLBMonitorBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, gslbServiceGroupLBMonitorBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBServiceGroupLBMonitorBinding `json:"gslbservicegroup_lbmonitor_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) GetGSLBServiceGroupLBMonitorBinding(name string) ([]models.GSLBServiceGroupLBMonitorBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", gslbServiceGroupLBMonitorBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBServiceGroupLBMonitorBinding `json:"gslbservicegroup_lbmonitor_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) CountGSLBServiceGroupLBMonitorBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", gslbServiceGroupLBMonitorBindingURL, url.PathEscape(name))
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
		} `json:"gslbservicegroup_lbmonitor_binding"`
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

// gslbservicegroup_servicegroupentitymonbindings_binding
func (s *GSLBService) GetAllGSLBServiceGroupServiceGroupEntityMonBindingsBinding() ([]models.GSLBServiceGroupServiceGroupEntityMonBindingsBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, gslbServiceGroupServiceGroupEntityMonBindingsBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBServiceGroupServiceGroupEntityMonBindingsBinding `json:"gslbservicegroup_servicegroupentitymonbindings_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) GetGSLBServiceGroupServiceGroupEntityMonBindingsBinding(name string) ([]models.GSLBServiceGroupServiceGroupEntityMonBindingsBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", gslbServiceGroupServiceGroupEntityMonBindingsBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBServiceGroupServiceGroupEntityMonBindingsBinding `json:"gslbservicegroup_servicegroupentitymonbindings_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) CountGSLBServiceGroupServiceGroupEntityMonBindingsBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", gslbServiceGroupServiceGroupEntityMonBindingsBindingURL, url.PathEscape(name))
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
		} `json:"gslbservicegroup_servicegroupentitymonbindings_binding"`
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

// gslbservice_binding
func (s *GSLBService) GetAllGSLBServiceBinding() ([]models.GSLBServiceBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, gslbServiceBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBServiceBinding `json:"gslbservice_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) GetGSLBServiceBinding(name string) (*models.GSLBServiceBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", gslbServiceBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBServiceBinding `json:"gslbservice_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) == 0 {
		return nil, fmt.Errorf("gslbservice_binding not found")
	}

	return &result.Bindings[0], nil
}

// gslbservice_dnsview_binding
func (s *GSLBService) AddGSLBServiceDNSViewBinding(resource models.GSLBServiceDNSViewBinding) error {
	payload := map[string]any{
		"gslbservice_dnsview_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, gslbServiceDNSViewBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) DeleteGSLBServiceDNSViewBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", gslbServiceDNSViewBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) GetAllGSLBServiceDNSViewBinding() ([]models.GSLBServiceDNSViewBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, gslbServiceDNSViewBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBServiceDNSViewBinding `json:"gslbservice_dnsview_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) GetGSLBServiceDNSViewBinding(name string) ([]models.GSLBServiceDNSViewBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", gslbServiceDNSViewBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBServiceDNSViewBinding `json:"gslbservice_dnsview_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) CountGSLBServiceDNSViewBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", gslbServiceDNSViewBindingURL, url.PathEscape(name))
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
		} `json:"gslbservice_dnsview_binding"`
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

// gslbservice_lbmonitor_binding
func (s *GSLBService) AddGSLBServiceLBMonitorBinding(resource models.GSLBServiceLBMonitorBinding) error {
	payload := map[string]any{
		"gslbservice_lbmonitor_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, gslbServiceLBMonitorBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) DeleteGSLBServiceLBMonitorBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", gslbServiceLBMonitorBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) GetAllGSLBServiceLBMonitorBinding() ([]models.GSLBServiceLBMonitorBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, gslbServiceLBMonitorBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBServiceLBMonitorBinding `json:"gslbservice_lbmonitor_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) GetGSLBServiceLBMonitorBinding(name string) ([]models.GSLBServiceLBMonitorBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", gslbServiceLBMonitorBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBServiceLBMonitorBinding `json:"gslbservice_lbmonitor_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) CountGSLBServiceLBMonitorBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", gslbServiceLBMonitorBindingURL, url.PathEscape(name))
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
		} `json:"gslbservice_lbmonitor_binding"`
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

// gslbsite
func (s *GSLBService) AddGSLBSite(resource models.GSLBSite) error {
	payload := map[string]any{
		"gslbsite": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, gslbSiteURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) DeleteGSLBSite(name string) error {
	reqURL := fmt.Sprintf("%s/%s", gslbSiteURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) UpdateGSLBSite(resource models.GSLBSite) error {
	payload := map[string]any{
		"gslbsite": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, gslbSiteURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) UnsetGSLBSite(resource models.GSLBSite) error {
	payload := map[string]any{
		"gslbsite": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, gslbSiteURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) GetAllGSLBSite() ([]models.GSLBSite, error) {
	req, err := s.client.NewRequest(http.MethodGet, gslbSiteURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Sites []models.GSLBSite `json:"gslbsite"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Sites, nil
}

func (s *GSLBService) GetGSLBSite(name string) (*models.GSLBSite, error) {
	reqURL := fmt.Sprintf("%s/%s", gslbSiteURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Sites []models.GSLBSite `json:"gslbsite"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Sites) == 0 {
		return nil, fmt.Errorf("gslbsite not found")
	}

	return &result.Sites[0], nil
}

func (s *GSLBService) CountGSLBSite() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, gslbSiteURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Sites []struct {
			Count float64 `json:"__count"`
		} `json:"gslbsite"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Sites) > 0 {
		return result.Sites[0].Count, nil
	}

	return 0, nil
}

func (s *GSLBService) RenameGSLBSite(resource models.GSLBSite) error {
	payload := map[string]any{
		"gslbsite": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, gslbSiteURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// gslbsite_binding
func (s *GSLBService) GetAllGSLBSiteBinding() ([]models.GSLBSiteBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, gslbSiteBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBSiteBinding `json:"gslbsite_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) GetGSLBSiteBinding(name string) (*models.GSLBSiteBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", gslbSiteBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBSiteBinding `json:"gslbsite_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) == 0 {
		return nil, fmt.Errorf("gslbsite_binding not found")
	}

	return &result.Bindings[0], nil
}

// gslbsite_gslbservicegroupmember_binding
func (s *GSLBService) GetAllGSLBSiteGSLBServiceGroupMemberBinding() ([]models.GSLBSiteGSLBServiceGroupMemberBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, gslbSiteGSLBServiceGroupMemberBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBSiteGSLBServiceGroupMemberBinding `json:"gslbsite_gslbservicegroupmember_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) GetGSLBSiteGSLBServiceGroupMemberBinding(name string) ([]models.GSLBSiteGSLBServiceGroupMemberBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", gslbSiteGSLBServiceGroupMemberBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBSiteGSLBServiceGroupMemberBinding `json:"gslbsite_gslbservicegroupmember_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) CountGSLBSiteGSLBServiceGroupMemberBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", gslbSiteGSLBServiceGroupMemberBindingURL, url.PathEscape(name))
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
		} `json:"gslbsite_gslbservicegroupmember_binding"`
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

// gslbsite_gslbservicegroup_binding
func (s *GSLBService) GetAllGSLBSiteGSLBServiceGroupBinding() ([]models.GSLBSiteGSLBServiceGroupBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, gslbSiteGSLBServiceGroupBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBSiteGSLBServiceGroupBinding `json:"gslbsite_gslbservicegroup_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) GetGSLBSiteGSLBServiceGroupBinding(name string) ([]models.GSLBSiteGSLBServiceGroupBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", gslbSiteGSLBServiceGroupBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBSiteGSLBServiceGroupBinding `json:"gslbsite_gslbservicegroup_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) CountGSLBSiteGSLBServiceGroupBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", gslbSiteGSLBServiceGroupBindingURL, url.PathEscape(name))
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
		} `json:"gslbsite_gslbservicegroup_binding"`
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

// gslbsite_gslbservice_binding
func (s *GSLBService) GetAllGSLBSiteGSLBServiceBinding() ([]models.GSLBSiteGSLBServiceBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, gslbSiteGSLBServiceBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBSiteGSLBServiceBinding `json:"gslbsite_gslbservice_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) GetGSLBSiteGSLBServiceBinding(name string) ([]models.GSLBSiteGSLBServiceBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", gslbSiteGSLBServiceBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBSiteGSLBServiceBinding `json:"gslbsite_gslbservice_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) CountGSLBSiteGSLBServiceBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", gslbSiteGSLBServiceBindingURL, url.PathEscape(name))
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
		} `json:"gslbsite_gslbservice_binding"`
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

// gslbsyncstatus
func (s *GSLBService) GetAllGSLBSyncStatus() ([]models.GSLBSyncStatus, error) {
	req, err := s.client.NewRequest(http.MethodGet, gslbSyncStatusURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Status []models.GSLBSyncStatus `json:"gslbsyncstatus"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Status, nil
}

// gslbvserver
func (s *GSLBService) AddGSLBVServer(resource models.GSLBVServer) error {
	payload := map[string]any{
		"gslbvserver": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, gslbVServerURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) DeleteGSLBVServer(name string) error {
	reqURL := fmt.Sprintf("%s/%s", gslbVServerURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) UpdateGSLBVServer(resource models.GSLBVServer) error {
	payload := map[string]any{
		"gslbvserver": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, gslbVServerURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) UnsetGSLBVServer(resource models.GSLBVServer) error {
	payload := map[string]any{
		"gslbvserver": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, gslbVServerURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) EnableGSLBVServer(resource models.GSLBVServer) error {
	payload := map[string]any{
		"gslbvserver": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, gslbVServerURL+"?action=enable", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) DisableGSLBVServer(resource models.GSLBVServer) error {
	payload := map[string]any{
		"gslbvserver": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, gslbVServerURL+"?action=disable", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) GetAllGSLBVServer() ([]models.GSLBVServer, error) {
	req, err := s.client.NewRequest(http.MethodGet, gslbVServerURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		VServers []models.GSLBVServer `json:"gslbvserver"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.VServers, nil
}

func (s *GSLBService) GetGSLBVServer(name string) (*models.GSLBVServer, error) {
	reqURL := fmt.Sprintf("%s/%s", gslbVServerURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		VServers []models.GSLBVServer `json:"gslbvserver"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.VServers) == 0 {
		return nil, fmt.Errorf("gslbvserver not found")
	}

	return &result.VServers[0], nil
}

func (s *GSLBService) CountGSLBVServer() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, gslbVServerURL+"?count=yes", nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		VServers []struct {
			Count float64 `json:"__count"`
		} `json:"gslbvserver"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.VServers) > 0 {
		return result.VServers[0].Count, nil
	}

	return 0, nil
}

func (s *GSLBService) RenameGSLBVServer(resource models.GSLBVServer) error {
	payload := map[string]any{
		"gslbvserver": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, gslbVServerURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// gslbvserver_binding
func (s *GSLBService) GetAllGSLBVServerBinding() ([]models.GSLBVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, gslbVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBVServerBinding `json:"gslbvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) GetGSLBVServerBinding(name string) (*models.GSLBVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", gslbVServerBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBVServerBinding `json:"gslbvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Bindings) == 0 {
		return nil, fmt.Errorf("gslbvserver_binding not found")
	}

	return &result.Bindings[0], nil
}

// gslbvserver_domain_binding
func (s *GSLBService) AddGSLBVServerDomainBinding(resource models.GSLBVServerDomainBinding) error {
	payload := map[string]any{
		"gslbvserver_domain_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, gslbVServerDomainBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) DeleteGSLBVServerDomainBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", gslbVServerDomainBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) GetAllGSLBVServerDomainBinding() ([]models.GSLBVServerDomainBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, gslbVServerDomainBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBVServerDomainBinding `json:"gslbvserver_domain_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) GetGSLBVServerDomainBinding(name string) ([]models.GSLBVServerDomainBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", gslbVServerDomainBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBVServerDomainBinding `json:"gslbvserver_domain_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) CountGSLBVServerDomainBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", gslbVServerDomainBindingURL, url.PathEscape(name))
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
		} `json:"gslbvserver_domain_binding"`
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

// gslbvserver_gslbservicegroupmember_binding
func (s *GSLBService) GetAllGSLBVServerGSLBServiceGroupMemberBinding() ([]models.GSLBVServerGSLBServiceGroupMemberBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, gslbVServerGSLBServiceGroupMemberBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBVServerGSLBServiceGroupMemberBinding `json:"gslbvserver_gslbservicegroupmember_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) GetGSLBVServerGSLBServiceGroupMemberBinding(name string) ([]models.GSLBVServerGSLBServiceGroupMemberBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", gslbVServerGSLBServiceGroupMemberBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBVServerGSLBServiceGroupMemberBinding `json:"gslbvserver_gslbservicegroupmember_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) CountGSLBVServerGSLBServiceGroupMemberBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", gslbVServerGSLBServiceGroupMemberBindingURL, url.PathEscape(name))
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
		} `json:"gslbvserver_gslbservicegroupmember_binding"`
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

// gslbvserver_gslbservicegroup_binding
func (s *GSLBService) AddGSLBVServerGSLBServiceGroupBinding(resource models.GSLBVServerGSLBServiceGroupBinding) error {
	payload := map[string]any{
		"gslbvserver_gslbservicegroup_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, gslbVServerGSLBServiceGroupBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) DeleteGSLBVServerGSLBServiceGroupBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", gslbVServerGSLBServiceGroupBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) GetAllGSLBVServerGSLBServiceGroupBinding() ([]models.GSLBVServerGSLBServiceGroupBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, gslbVServerGSLBServiceGroupBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBVServerGSLBServiceGroupBinding `json:"gslbvserver_gslbservicegroup_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) GetGSLBVServerGSLBServiceGroupBinding(name string) ([]models.GSLBVServerGSLBServiceGroupBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", gslbVServerGSLBServiceGroupBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBVServerGSLBServiceGroupBinding `json:"gslbvserver_gslbservicegroup_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) CountGSLBVServerGSLBServiceGroupBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", gslbVServerGSLBServiceGroupBindingURL, url.PathEscape(name))
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
		} `json:"gslbvserver_gslbservicegroup_binding"`
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

// gslbvserver_gslbservice_binding
func (s *GSLBService) AddGSLBVServerGSLBServiceBinding(resource models.GSLBVServerGSLBServiceBinding) error {
	payload := map[string]any{
		"gslbvserver_gslbservice_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, gslbVServerGSLBServiceBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) DeleteGSLBVServerGSLBServiceBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", gslbVServerGSLBServiceBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) GetAllGSLBVServerGSLBServiceBinding() ([]models.GSLBVServerGSLBServiceBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, gslbVServerGSLBServiceBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBVServerGSLBServiceBinding `json:"gslbvserver_gslbservice_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) GetGSLBVServerGSLBServiceBinding(name string) ([]models.GSLBVServerGSLBServiceBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", gslbVServerGSLBServiceBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBVServerGSLBServiceBinding `json:"gslbvserver_gslbservice_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) CountGSLBVServerGSLBServiceBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", gslbVServerGSLBServiceBindingURL, url.PathEscape(name))
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
		} `json:"gslbvserver_gslbservice_binding"`
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

// gslbvserver_spilloverpolicy_binding
func (s *GSLBService) AddGSLBVServerSpilloverPolicyBinding(resource models.GSLBVServerSpilloverPolicyBinding) error {
	payload := map[string]any{
		"gslbvserver_spilloverpolicy_binding": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, gslbVServerSpilloverPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) DeleteGSLBVServerSpilloverPolicyBinding(name string, args string) error {
	reqURL := fmt.Sprintf("%s/%s", gslbVServerSpilloverPolicyBindingURL, url.PathEscape(name))
	if args != "" {
		reqURL += "?args=" + url.QueryEscape(args)
	}
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *GSLBService) GetAllGSLBVServerSpilloverPolicyBinding() ([]models.GSLBVServerSpilloverPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, gslbVServerSpilloverPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBVServerSpilloverPolicyBinding `json:"gslbvserver_spilloverpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) GetGSLBVServerSpilloverPolicyBinding(name string) ([]models.GSLBVServerSpilloverPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", gslbVServerSpilloverPolicyBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Bindings []models.GSLBVServerSpilloverPolicyBinding `json:"gslbvserver_spilloverpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Bindings, nil
}

func (s *GSLBService) CountGSLBVServerSpilloverPolicyBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", gslbVServerSpilloverPolicyBindingURL, url.PathEscape(name))
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
		} `json:"gslbvserver_spilloverpolicy_binding"`
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
