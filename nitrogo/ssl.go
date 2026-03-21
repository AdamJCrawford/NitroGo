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
	sslActionURL                            = "/nitro/v1/config/sslaction"
	sslCACertGroupURL                       = "/nitro/v1/config/sslcacertgroup"
	sslCACertGroupBindingURL                = "/nitro/v1/config/sslcacertgroup_binding"
	sslCACertGroupSSLCertKeyBindingURL      = "/nitro/v1/config/sslcacertgroup_sslcertkey_binding"
	sslCertURL                              = "/nitro/v1/config/sslcert"
	sslCertBundleURL                        = "/nitro/v1/config/sslcertbundle"
	sslCertChainURL                         = "/nitro/v1/config/sslcertchain"
	sslCertChainBindingURL                  = "/nitro/v1/config/sslcertchain_binding"
	sslCertChainSSLCertKeyBindingURL        = "/nitro/v1/config/sslcertchain_sslcertkey_binding"
	sslCertFileURL                          = "/nitro/v1/config/sslcertfile"
	sslCertificateChainURL                  = "/nitro/v1/config/sslcertificatechain"
	sslCertKeyURL                           = "/nitro/v1/config/sslcertkey"
	sslCertKeyBindingURL                    = "/nitro/v1/config/sslcertkey_binding"
	sslCertKeyCRLDistributionBindingURL     = "/nitro/v1/config/sslcertkey_crldistribution_binding"
	sslCertKeyServiceBindingURL             = "/nitro/v1/config/sslcertkey_service_binding"
	sslCertKeySSLOCSPResponderBindingURL    = "/nitro/v1/config/sslcertkey_sslocspresponder_binding"
	sslCertKeySSLProfileBindingURL          = "/nitro/v1/config/sslcertkey_sslprofile_binding"
	sslCertKeySSLVServerBindingURL          = "/nitro/v1/config/sslcertkey_sslvserver_binding"
	sslCertLinkURL                          = "/nitro/v1/config/sslcertlink"
	sslCertReqURL                           = "/nitro/v1/config/sslcertreq"
	sslCipherURL                            = "/nitro/v1/config/sslcipher"
	sslCipherSuiteURL                       = "/nitro/v1/config/sslciphersuite"
	sslCipherBindingURL                     = "/nitro/v1/config/sslcipher_binding"
	sslCipherIndividualCipherBindingURL     = "/nitro/v1/config/sslcipher_individualcipher_binding"
	sslCipherServiceGroupBindingURL         = "/nitro/v1/config/sslcipher_servicegroup_binding"
	sslCipherServiceBindingURL              = "/nitro/v1/config/sslcipher_service_binding"
	sslCipherSSLCipherSuiteBindingURL       = "/nitro/v1/config/sslcipher_sslciphersuite_binding"
	sslCipherSSLProfileBindingURL           = "/nitro/v1/config/sslcipher_sslprofile_binding"
	sslCipherSSLVServerBindingURL           = "/nitro/v1/config/sslcipher_sslvserver_binding"
	sslCRLURL                               = "/nitro/v1/config/sslcrl"
	sslCRLFileURL                           = "/nitro/v1/config/sslcrlfile"
	sslCRLBindingURL                        = "/nitro/v1/config/sslcrl_binding"
	sslCRLSerialNumberBindingURL            = "/nitro/v1/config/sslcrl_serialnumber_binding"
	sslDHFileURL                            = "/nitro/v1/config/ssldhfile"
	sslDHParamURL                           = "/nitro/v1/config/ssldhparam"
	sslDTLSProfileURL                       = "/nitro/v1/config/ssldtlsprofile"
	sslECDSAKeyURL                          = "/nitro/v1/config/sslecdsakey"
	sslFIPSURL                              = "/nitro/v1/config/sslfips"
	sslFIPSKeyURL                           = "/nitro/v1/config/sslfipskey"
	sslFIPSSIMSourceURL                     = "/nitro/v1/config/sslfipssimsource"
	sslFIPSSIMTargetURL                     = "/nitro/v1/config/sslfipssimtarget"
	sslGlobalBindingURL                     = "/nitro/v1/config/sslglobal_binding"
	sslGlobalSSLPolicyBindingURL            = "/nitro/v1/config/sslglobal_sslpolicy_binding"
	sslHSMKeyURL                            = "/nitro/v1/config/sslhsmkey"
	sslKeyFileURL                           = "/nitro/v1/config/sslkeyfile"
	sslLogProfileURL                        = "/nitro/v1/config/ssllogprofile"
	sslOCSPResponderURL                     = "/nitro/v1/config/sslocspresponder"
	sslParameterURL                         = "/nitro/v1/config/sslparameter"
	sslPKCS12URL                            = "/nitro/v1/config/sslpkcs12"
	sslPKCS8URL                             = "/nitro/v1/config/sslpkcs8"
	sslPolicyURL                            = "/nitro/v1/config/sslpolicy"
	sslPolicyLabelURL                       = "/nitro/v1/config/sslpolicylabel"
	sslPolicyLabelBindingURL                = "/nitro/v1/config/sslpolicylabel_binding"
	sslPolicyLabelSSLPolicyBindingURL       = "/nitro/v1/config/sslpolicylabel_sslpolicy_binding"
	sslPolicyBindingURL                     = "/nitro/v1/config/sslpolicy_binding"
	sslPolicyCSVServerBindingURL            = "/nitro/v1/config/sslpolicy_csvserver_binding"
	sslPolicyLBVServerBindingURL            = "/nitro/v1/config/sslpolicy_lbvserver_binding"
	sslPolicySSLGlobalBindingURL            = "/nitro/v1/config/sslpolicy_sslglobal_binding"
	sslPolicySSLPolicyLabelBindingURL       = "/nitro/v1/config/sslpolicy_sslpolicylabel_binding"
	sslPolicySSLServiceBindingURL           = "/nitro/v1/config/sslpolicy_sslservice_binding"
	sslPolicySSLVServerBindingURL           = "/nitro/v1/config/sslpolicy_sslvserver_binding"
	sslProfileURL                           = "/nitro/v1/config/sslprofile"
	sslProfileBindingURL                    = "/nitro/v1/config/sslprofile_binding"
	sslProfileECCCurveBindingURL            = "/nitro/v1/config/sslprofile_ecccurve_binding"
	sslProfileSSLCertKeyBindingURL          = "/nitro/v1/config/sslprofile_sslcertkey_binding"
	sslProfileSSLCipherSuiteBindingURL      = "/nitro/v1/config/sslprofile_sslciphersuite_binding"
	sslProfileSSLCipherBindingURL           = "/nitro/v1/config/sslprofile_sslcipher_binding"
	sslProfileSSLVServerBindingURL          = "/nitro/v1/config/sslprofile_sslvserver_binding"
	sslRSAKeyURL                            = "/nitro/v1/config/sslrsakey"
	sslServiceURL                           = "/nitro/v1/config/sslservice"
	sslServiceGroupURL                      = "/nitro/v1/config/sslservicegroup"
	sslServiceGroupBindingURL               = "/nitro/v1/config/sslservicegroup_binding"
	sslServiceGroupECCCurveBindingURL       = "/nitro/v1/config/sslservicegroup_ecccurve_binding"
	sslServiceGroupSSLCertKeyBindingURL     = "/nitro/v1/config/sslservicegroup_sslcertkey_binding"
	sslServiceGroupSSLCipherSuiteBindingURL = "/nitro/v1/config/sslservicegroup_sslciphersuite_binding"
	sslServiceGroupSSLCipherBindingURL      = "/nitro/v1/config/sslservicegroup_sslcipher_binding"
	sslServiceBindingURL                    = "/nitro/v1/config/sslservice_binding"
	sslServiceECCCurveBindingURL            = "/nitro/v1/config/sslservice_ecccurve_binding"
	sslServiceSSLCertKeyBindingURL          = "/nitro/v1/config/sslservice_sslcertkey_binding"
	sslServiceSSLCipherSuiteBindingURL      = "/nitro/v1/config/sslservice_sslciphersuite_binding"
	sslServiceSSLCipherBindingURL           = "/nitro/v1/config/sslservice_sslcipher_binding"
	sslServiceSSLPolicyBindingURL           = "/nitro/v1/config/sslservice_sslpolicy_binding"
	sslVServerURL                           = "/nitro/v1/config/sslvserver"
	sslVServerBindingURL                    = "/nitro/v1/config/sslvserver_binding"
	sslVServerECCCurveBindingURL            = "/nitro/v1/config/sslvserver_ecccurve_binding"
	sslVServerSSLCertKeyBindingURL          = "/nitro/v1/config/sslvserver_sslcertkey_binding"
	sslVServerSSLCipherSuiteBindingURL      = "/nitro/v1/config/sslvserver_sslciphersuite_binding"
	sslVServerSSLCipherBindingURL           = "/nitro/v1/config/sslvserver_sslcipher_binding"
	sslVServerSSLPolicyBindingURL           = "/nitro/v1/config/sslvserver_sslpolicy_binding"
	sslWrapKeyURL                           = "/nitro/v1/config/sslwrapkey"
)

// SSL
// SSL Configuration
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/ssl/ssl
type SSLService struct {
	client *Client
}

// sslaction
func (s *SSLService) AddSSLAction(action models.SSLAction) error {
	payload := map[string]any{"sslaction": action}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslActionURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLAction(name string) error {
	reqURL := fmt.Sprintf("%s/%s", sslActionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLAction() ([]models.SSLAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslActionURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLAction `json:"sslaction"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLAction(name string) ([]models.SSLAction, error) {
	reqURL := fmt.Sprintf("%s/%s", sslActionURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLAction `json:"sslaction"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLAction() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslActionURL+"?count=yes", nil)
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
		} `json:"sslaction"`
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

// sslcacertgroup
func (s *SSLService) AddSSLCACertGroup(group models.SSLCACertGroup) error {
	payload := map[string]any{"sslcacertgroup": group}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslCACertGroupURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLCACertGroup(cacertgroupname string) error {
	reqURL := fmt.Sprintf("%s/%s", sslCACertGroupURL, url.QueryEscape(cacertgroupname))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLCACertGroup() ([]models.SSLCACertGroup, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslCACertGroupURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCACertGroup `json:"sslcacertgroup"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLCACertGroup(cacertgroupname string) ([]models.SSLCACertGroup, error) {
	reqURL := fmt.Sprintf("%s/%s", sslCACertGroupURL, url.QueryEscape(cacertgroupname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCACertGroup `json:"sslcacertgroup"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLCACertGroup() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslCACertGroupURL+"?count=yes", nil)
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
		} `json:"sslcacertgroup"`
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

// sslcacertgroup_binding
func (s *SSLService) GetAllSSLCACertGroupBinding() ([]models.SSLCACertGroupBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslCACertGroupBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCACertGroupBinding `json:"sslcacertgroup_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLCACertGroupBinding(cacertgroupname string) ([]models.SSLCACertGroupBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslCACertGroupBindingURL, url.QueryEscape(cacertgroupname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCACertGroupBinding `json:"sslcacertgroup_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

// sslcacertgroup_sslcertkey_binding
func (s *SSLService) AddSSLCACertGroupSSLCertKeyBinding(binding models.SSLCACertGroupSSLCertKeyBinding) error {
	payload := map[string]any{"sslcacertgroup_sslcertkey_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, sslCACertGroupSSLCertKeyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLCACertGroupSSLCertKeyBinding(cacertgroupname string, certkeyname string) error {
	reqURL := fmt.Sprintf("%s/%s?args=certkeyname:%s", sslCACertGroupSSLCertKeyBindingURL, url.QueryEscape(cacertgroupname), url.QueryEscape(certkeyname))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLCACertGroupSSLCertKeyBinding() ([]models.SSLCACertGroupSSLCertKeyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslCACertGroupSSLCertKeyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCACertGroupSSLCertKeyBinding `json:"sslcacertgroup_sslcertkey_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLCACertGroupSSLCertKeyBinding(cacertgroupname string) ([]models.SSLCACertGroupSSLCertKeyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslCACertGroupSSLCertKeyBindingURL, url.QueryEscape(cacertgroupname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCACertGroupSSLCertKeyBinding `json:"sslcacertgroup_sslcertkey_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLCACertGroupSSLCertKeyBinding(cacertgroupname string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", sslCACertGroupSSLCertKeyBindingURL, url.QueryEscape(cacertgroupname))
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
		} `json:"sslcacertgroup_sslcertkey_binding"`
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

// sslcert
func (s *SSLService) CreateSSLCert(cert models.SSLCert) error {
	payload := map[string]any{"sslcert": cert}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslCertURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

// sslcertbundle
func (s *SSLService) ImportSSLCertBundle(bundle models.SSLCertBundle) error {
	payload := map[string]any{"sslcertbundle": bundle}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslCertBundleURL+"?action=import", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLCertBundle(name string) error {
	reqURL := fmt.Sprintf("%s/%s", sslCertBundleURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) ApplySSLCertBundle() error {
	req, err := s.client.NewRequest(http.MethodPost, sslCertBundleURL+"?action=apply", nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) ExportSSLCertBundle(bundle models.SSLCertBundle) error {
	payload := map[string]any{"sslcertbundle": bundle}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslCertBundleURL+"?action=export", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLCertBundle() ([]models.SSLCertBundle, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslCertBundleURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCertBundle `json:"sslcertbundle"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLCertBundle() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslCertBundleURL+"?count=yes", nil)
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
		} `json:"sslcertbundle"`
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

// sslcertchain
func (s *SSLService) GetAllSSLCertChain() ([]models.SSLCertChain, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslCertChainURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCertChain `json:"sslcertchain"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLCertChain(certkeyname string) ([]models.SSLCertChain, error) {
	reqURL := fmt.Sprintf("%s/%s", sslCertChainURL, url.QueryEscape(certkeyname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCertChain `json:"sslcertchain"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLCertChain() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslCertChainURL+"?count=yes", nil)
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
		} `json:"sslcertchain"`
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

// sslcertchain_binding
func (s *SSLService) GetSSLCertChainBinding(certkeyname string) ([]models.SSLCertChainBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslCertChainBindingURL, url.QueryEscape(certkeyname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCertChainBinding `json:"sslcertchain_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

// sslcertchain_sslcertkey_binding
func (s *SSLService) GetSSLCertChainSSLCertKeyBinding(certkeyname string) ([]models.SSLCertChainSSLCertKeyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslCertChainSSLCertKeyBindingURL, url.QueryEscape(certkeyname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCertChainSSLCertKeyBinding `json:"sslcertchain_sslcertkey_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLCertChainSSLCertKeyBinding(certkeyname string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", sslCertChainSSLCertKeyBindingURL, url.QueryEscape(certkeyname))
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
		} `json:"sslcertchain_sslcertkey_binding"`
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

// sslcertfile
func (s *SSLService) ImportSSLCertFile(file models.SSLCertFile) error {
	payload := map[string]any{"sslcertfile": file}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslCertFileURL+"?action=import", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLCertFile(name string) error {
	reqURL := fmt.Sprintf("%s/%s", sslCertFileURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLCertFile() ([]models.SSLCertFile, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslCertFileURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCertFile `json:"sslcertfile"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLCertFile() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslCertFileURL+"?count=yes", nil)
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
		} `json:"sslcertfile"`
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

// sslcertificatechain
func (s *SSLService) AddSSLCertificateChain(chain models.SSLCertificateChain) error {
	payload := map[string]any{"sslcertificatechain": chain}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslCertificateChainURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLCertificateChain() ([]models.SSLCertificateChain, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslCertificateChainURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCertificateChain `json:"sslcertificatechain"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLCertificateChain(certkeyname string) ([]models.SSLCertificateChain, error) {
	reqURL := fmt.Sprintf("%s/%s", sslCertificateChainURL, url.QueryEscape(certkeyname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCertificateChain `json:"sslcertificatechain"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLCertificateChain() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslCertificateChainURL+"?count=yes", nil)
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
		} `json:"sslcertificatechain"`
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

// sslcertkey
func (s *SSLService) AddSSLCertKey(certkey models.SSLCertKey) error {
	payload := map[string]any{"sslcertkey": certkey}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslCertKeyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLCertKey(certkey string) error {
	reqURL := fmt.Sprintf("%s/%s", sslCertKeyURL, url.QueryEscape(certkey))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) UpdateSSLCertKey(certkey models.SSLCertKey) error {
	payload := map[string]any{"sslcertkey": certkey}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, sslCertKeyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) UnsetSSLCertKey(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"sslcertkey": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslCertKeyURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) LinkSSLCertKey(certkey models.SSLCertKey) error {
	payload := map[string]any{"sslcertkey": certkey}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslCertKeyURL+"?action=link", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) UnlinkSSLCertKey(certkey models.SSLCertKey) error {
	payload := map[string]any{"sslcertkey": certkey}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslCertKeyURL+"?action=unlink", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) ChangeSSLCertKey(certkey models.SSLCertKey) error {
	payload := map[string]any{"sslcertkey": certkey}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslCertKeyURL+"?action=change", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) ClearSSLCertKey(certkey models.SSLCertKey) error {
	payload := map[string]any{"sslcertkey": certkey}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslCertKeyURL+"?action=clear", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLCertKey() ([]models.SSLCertKey, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslCertKeyURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCertKey `json:"sslcertkey"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLCertKey(certkey string) ([]models.SSLCertKey, error) {
	reqURL := fmt.Sprintf("%s/%s", sslCertKeyURL, url.QueryEscape(certkey))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCertKey `json:"sslcertkey"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLCertKey() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslCertKeyURL+"?count=yes", nil)
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
		} `json:"sslcertkey"`
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

// sslcertkey_binding
func (s *SSLService) GetAllSSLCertKeyBinding() ([]models.SSLCertKeyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslCertKeyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCertKeyBinding `json:"sslcertkey_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLCertKeyBinding(certkey string) ([]models.SSLCertKeyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslCertKeyBindingURL, url.QueryEscape(certkey))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCertKeyBinding `json:"sslcertkey_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

// sslcertkey_crldistribution_binding
func (s *SSLService) GetAllSSLCertKeyCRLDistributionBinding() ([]models.SSLCertKeyCRLDistributionBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslCertKeyCRLDistributionBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCertKeyCRLDistributionBinding `json:"sslcertkey_crldistribution_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLCertKeyCRLDistributionBinding(certkey string) ([]models.SSLCertKeyCRLDistributionBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslCertKeyCRLDistributionBindingURL, url.QueryEscape(certkey))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCertKeyCRLDistributionBinding `json:"sslcertkey_crldistribution_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLCertKeyCRLDistributionBinding(certkey string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", sslCertKeyCRLDistributionBindingURL, url.QueryEscape(certkey))
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
		} `json:"sslcertkey_crldistribution_binding"`
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

// sslcertkey_service_binding
func (s *SSLService) GetAllSSLCertKeyServiceBinding() ([]models.SSLCertKeyServiceBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslCertKeyServiceBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCertKeyServiceBinding `json:"sslcertkey_service_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLCertKeyServiceBinding(certkey string) ([]models.SSLCertKeyServiceBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslCertKeyServiceBindingURL, url.QueryEscape(certkey))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCertKeyServiceBinding `json:"sslcertkey_service_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLCertKeyServiceBinding(certkey string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", sslCertKeyServiceBindingURL, url.QueryEscape(certkey))
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
		} `json:"sslcertkey_service_binding"`
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

// sslcertkey_sslocspresponder_binding
func (s *SSLService) AddSSLCertKeySSLOCSPResponderBinding(binding models.SSLCertKeySSLOCSPResponderBinding) error {
	payload := map[string]any{"sslcertkey_sslocspresponder_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, sslCertKeySSLOCSPResponderBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLCertKeySSLOCSPResponderBinding(certkey string, ocspresponder string) error {
	reqURL := fmt.Sprintf("%s/%s?args=ocspresponder:%s", sslCertKeySSLOCSPResponderBindingURL, url.QueryEscape(certkey), url.QueryEscape(ocspresponder))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLCertKeySSLOCSPResponderBinding() ([]models.SSLCertKeySSLOCSPResponderBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslCertKeySSLOCSPResponderBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCertKeySSLOCSPResponderBinding `json:"sslcertkey_sslocspresponder_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLCertKeySSLOCSPResponderBinding(certkey string) ([]models.SSLCertKeySSLOCSPResponderBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslCertKeySSLOCSPResponderBindingURL, url.QueryEscape(certkey))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCertKeySSLOCSPResponderBinding `json:"sslcertkey_sslocspresponder_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLCertKeySSLOCSPResponderBinding(certkey string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", sslCertKeySSLOCSPResponderBindingURL, url.QueryEscape(certkey))
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
		} `json:"sslcertkey_sslocspresponder_binding"`
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

// sslcertkey_sslprofile_binding
func (s *SSLService) GetAllSSLCertKeySSLProfileBinding() ([]models.SSLCertKeySSLProfileBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslCertKeySSLProfileBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCertKeySSLProfileBinding `json:"sslcertkey_sslprofile_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLCertKeySSLProfileBinding(certkey string) ([]models.SSLCertKeySSLProfileBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslCertKeySSLProfileBindingURL, url.QueryEscape(certkey))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCertKeySSLProfileBinding `json:"sslcertkey_sslprofile_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLCertKeySSLProfileBinding(certkey string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", sslCertKeySSLProfileBindingURL, url.QueryEscape(certkey))
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
		} `json:"sslcertkey_sslprofile_binding"`
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

// sslcertkey_sslvserver_binding
func (s *SSLService) GetAllSSLCertKeySSLVServerBinding() ([]models.SSLCertKeySSLVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslCertKeySSLVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCertKeySSLVServerBinding `json:"sslcertkey_sslvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLCertKeySSLVServerBinding(certkey string) ([]models.SSLCertKeySSLVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslCertKeySSLVServerBindingURL, url.QueryEscape(certkey))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCertKeySSLVServerBinding `json:"sslcertkey_sslvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLCertKeySSLVServerBinding(certkey string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", sslCertKeySSLVServerBindingURL, url.QueryEscape(certkey))
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
		} `json:"sslcertkey_sslvserver_binding"`
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

// sslcertlink
func (s *SSLService) GetAllSSLCertLink() ([]models.SSLCertLink, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslCertLinkURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCertLink `json:"sslcertlink"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLCertLink() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslCertLinkURL+"?count=yes", nil)
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
		} `json:"sslcertlink"`
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

// sslcertreq
func (s *SSLService) CreateSSLCertReq(certreq models.SSLCertReq) error {
	payload := map[string]any{"sslcertreq": certreq}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslCertReqURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

// sslcipher
func (s *SSLService) AddSSLCipher(cipher models.SSLCipher) error {
	payload := map[string]any{"sslcipher": cipher}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslCipherURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLCipher(ciphergroupname string) error {
	reqURL := fmt.Sprintf("%s/%s", sslCipherURL, url.QueryEscape(ciphergroupname))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) UpdateSSLCipher(cipher models.SSLCipher) error {
	payload := map[string]any{"sslcipher": cipher}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, sslCipherURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) UnsetSSLCipher(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"sslcipher": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslCipherURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLCipher() ([]models.SSLCipher, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslCipherURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCipher `json:"sslcipher"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLCipher(ciphergroupname string) ([]models.SSLCipher, error) {
	reqURL := fmt.Sprintf("%s/%s", sslCipherURL, url.QueryEscape(ciphergroupname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCipher `json:"sslcipher"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLCipher() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslCipherURL+"?count=yes", nil)
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
		} `json:"sslcipher"`
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

// sslciphersuite
func (s *SSLService) GetAllSSLCipherSuite() ([]models.SSLCipherSuite, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslCipherSuiteURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCipherSuite `json:"sslciphersuite"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLCipherSuite(ciphername string) ([]models.SSLCipherSuite, error) {
	reqURL := fmt.Sprintf("%s/%s", sslCipherSuiteURL, url.QueryEscape(ciphername))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCipherSuite `json:"sslciphersuite"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLCipherSuite() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslCipherSuiteURL+"?count=yes", nil)
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
		} `json:"sslciphersuite"`
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

// sslcipher_binding
func (s *SSLService) GetAllSSLCipherBinding() ([]models.SSLCipherBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslCipherBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCipherBinding `json:"sslcipher_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLCipherBinding(ciphergroupname string) ([]models.SSLCipherBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslCipherBindingURL, url.QueryEscape(ciphergroupname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCipherBinding `json:"sslcipher_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

// sslcipher_individualcipher_binding
func (s *SSLService) GetAllSSLCipherIndividualCipherBinding() ([]models.SSLCipherIndividualCipherBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslCipherIndividualCipherBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCipherIndividualCipherBinding `json:"sslcipher_individualcipher_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLCipherIndividualCipherBinding(ciphergroupname string) ([]models.SSLCipherIndividualCipherBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslCipherIndividualCipherBindingURL, url.QueryEscape(ciphergroupname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCipherIndividualCipherBinding `json:"sslcipher_individualcipher_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLCipherIndividualCipherBinding(ciphergroupname string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", sslCipherIndividualCipherBindingURL, url.QueryEscape(ciphergroupname))
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
		} `json:"sslcipher_individualcipher_binding"`
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

// sslcipher_servicegroup_binding
// This object has no operations documented. These are guesses and need to be tested before they will be uncommented.
// func (s *SSLService) GetAllSSLCipherServiceGroupBinding() {}
// func (s *SSLService) GetSSLCipherServiceGroupBinding()    {}
// func (s *SSLService) CountSSLCipherServiceGroupBinding()  {}

// sslcipher_service_binding
// This object has no operations documented. These are guesses and need to be tested before they will be uncommented.
// func (s *SSLService) GetAllSSLCipherServiceBinding() {}
// func (s *SSLService) GetSSLCipherServiceBinding()    {}
// func (s *SSLService) CountSSLCipherServiceBinding()  {}

// sslcipher_sslciphersuite_binding
func (s *SSLService) AddSSLCipherSSLCipherSuiteBinding(binding models.SSLCipherSSLCipherSuiteBinding) error {
	payload := map[string]any{"sslcipher_sslciphersuite_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, sslCipherSSLCipherSuiteBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLCipherSSLCipherSuiteBinding(ciphergroupname string, ciphername string) error {
	reqURL := fmt.Sprintf("%s/%s?args=ciphername:%s", sslCipherSSLCipherSuiteBindingURL, url.QueryEscape(ciphergroupname), url.QueryEscape(ciphername))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLCipherSSLCipherSuiteBinding() ([]models.SSLCipherSSLCipherSuiteBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslCipherSSLCipherSuiteBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCipherSSLCipherSuiteBinding `json:"sslcipher_sslciphersuite_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLCipherSSLCipherSuiteBinding(ciphergroupname string) ([]models.SSLCipherSSLCipherSuiteBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslCipherSSLCipherSuiteBindingURL, url.QueryEscape(ciphergroupname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCipherSSLCipherSuiteBinding `json:"sslcipher_sslciphersuite_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLCipherSSLCipherSuiteBinding(ciphergroupname string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", sslCipherSSLCipherSuiteBindingURL, url.QueryEscape(ciphergroupname))
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
		} `json:"sslcipher_sslciphersuite_binding"`
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

// sslcipher_sslprofile_binding
func (s *SSLService) GetAllSSLCipherSSLProfileBinding() ([]models.SSLCipherSSLProfileBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslCipherSSLProfileBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCipherSSLProfileBinding `json:"sslcipher_sslprofile_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLCipherSSLProfileBinding(ciphergroupname string) ([]models.SSLCipherSSLProfileBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslCipherSSLProfileBindingURL, url.QueryEscape(ciphergroupname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCipherSSLProfileBinding `json:"sslcipher_sslprofile_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLCipherSSLProfileBinding(ciphergroupname string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", sslCipherSSLProfileBindingURL, url.QueryEscape(ciphergroupname))
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
		} `json:"sslcipher_sslprofile_binding"`
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

// sslcipher_sslvserver_binding
// This object has no operations documented. These are guesses and need to be tested before they will be uncommented.
// func (s *SSLService) GetAllSSLCipherSSLVServerBinding() {}
// func (s *SSLService) GetSSLCipherSSLVServerBinding()    {}
// func (s *SSLService) CountSSLCipherSSLVServerBinding()  {}

// sslcrl
func (s *SSLService) AddSSLCRL(crl models.SSLCRL) error {
	payload := map[string]any{"sslcrl": crl}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslCRLURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLCRL(crlname string) error {
	reqURL := fmt.Sprintf("%s/%s", sslCRLURL, url.QueryEscape(crlname))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) UpdateSSLCRL(crl models.SSLCRL) error {
	payload := map[string]any{"sslcrl": crl}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, sslCRLURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) UnsetSSLCRL(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"sslcrl": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslCRLURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) CreateSSLCRL(crl models.SSLCRL) error {
	payload := map[string]any{"sslcrl": crl}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslCRLURL+"?action=create", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLCRL() ([]models.SSLCRL, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslCRLURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCRL `json:"sslcrl"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLCRL(crlname string) ([]models.SSLCRL, error) {
	reqURL := fmt.Sprintf("%s/%s", sslCRLURL, url.QueryEscape(crlname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCRL `json:"sslcrl"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLCRL() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslCRLURL+"?count=yes", nil)
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
		} `json:"sslcrl"`
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

// sslcrlfile
func (s *SSLService) ImportSSLCRLFile(file models.SSLCRLFile) error {
	payload := map[string]any{"sslcrlfile": file}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslCRLFileURL+"?action=import", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLCRLFile(name string) error {
	reqURL := fmt.Sprintf("%s/%s", sslCRLFileURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLCRLFile() ([]models.SSLCRLFile, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslCRLFileURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCRLFile `json:"sslcrlfile"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLCRLFile() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslCRLFileURL+"?count=yes", nil)
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
		} `json:"sslcrlfile"`
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

// sslcrl_binding
func (s *SSLService) GetAllSSLCRLBinding() ([]models.SSLCRLBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslCRLBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCRLBinding `json:"sslcrl_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLCRLBinding(crlname string) ([]models.SSLCRLBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslCRLBindingURL, url.QueryEscape(crlname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCRLBinding `json:"sslcrl_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

// sslcrl_serialnumber_binding
func (s *SSLService) GetAllSSLCRLSerialNumberBinding() ([]models.SSLCRLSerialNumberBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslCRLSerialNumberBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCRLSerialNumberBinding `json:"sslcrl_serialnumber_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLCRLSerialNumberBinding(crlname string) ([]models.SSLCRLSerialNumberBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslCRLSerialNumberBindingURL, url.QueryEscape(crlname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLCRLSerialNumberBinding `json:"sslcrl_serialnumber_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLCRLSerialNumberBinding(crlname string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", sslCRLSerialNumberBindingURL, url.QueryEscape(crlname))
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
		} `json:"sslcrl_serialnumber_binding"`
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

// ssldhfile
func (s *SSLService) ImportSSLDHFile(file models.SSLDHFile) error {
	payload := map[string]any{"ssldhfile": file}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslDHFileURL+"?action=import", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLDHFile(name string) error {
	reqURL := fmt.Sprintf("%s/%s", sslDHFileURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLDHFile() ([]models.SSLDHFile, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslDHFileURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLDHFile `json:"ssldhfile"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLDHFile() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslDHFileURL+"?count=yes", nil)
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
		} `json:"ssldhfile"`
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

// ssldhparam
func (s *SSLService) ParamSSLDHParam(param models.SSLDHParam) error {
	payload := map[string]any{"ssldhparam": param}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslDHParamURL+"?action=param", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

// ssldtlsprofile
func (s *SSLService) AddSSLDTLSProfile(profile models.SSLDTLSProfile) error {
	payload := map[string]any{"ssldtlsprofile": profile}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslDTLSProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLDTLSProfile(name string) error {
	reqURL := fmt.Sprintf("%s/%s", sslDTLSProfileURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) UpdateSSLDTLSProfile(profile models.SSLDTLSProfile) error {
	payload := map[string]any{"ssldtlsprofile": profile}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, sslDTLSProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) UnsetSSLDTLSProfile(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"ssldtlsprofile": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslDTLSProfileURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLDTLSProfile() ([]models.SSLDTLSProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslDTLSProfileURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLDTLSProfile `json:"ssldtlsprofile"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLDTLSProfile(name string) ([]models.SSLDTLSProfile, error) {
	reqURL := fmt.Sprintf("%s/%s", sslDTLSProfileURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLDTLSProfile `json:"ssldtlsprofile"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLDTLSProfile() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslDTLSProfileURL+"?count=yes", nil)
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
		} `json:"ssldtlsprofile"`
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

// sslecdsakey
func (s *SSLService) CreateSSLECDSAKey(key models.SSLECDSAKey) error {
	payload := map[string]any{"sslecdsakey": key}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslECDSAKeyURL+"?action=create", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

// sslfips
func (s *SSLService) UpdateSSLFIPS(fips models.SSLFIPS) error {
	payload := map[string]any{"sslfips": fips}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, sslFIPSURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) UnsetSSLFIPS(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"sslfips": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslFIPSURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) ChangeSSLFIPS(fips models.SSLFIPS) error {
	payload := map[string]any{"sslfips": fips}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslFIPSURL+"?action=change", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) ResetSSLFIPS(fips models.SSLFIPS) error {
	payload := map[string]any{"sslfips": fips}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslFIPSURL+"?action=reset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLFIPS() (models.SSLFIPS, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslFIPSURL, nil)
	if err != nil {
		return models.SSLFIPS{}, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return models.SSLFIPS{}, err
	}
	var result struct {
		Items []models.SSLFIPS `json:"sslfips"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.SSLFIPS{}, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Items) > 0 {
		return result.Items[0], nil
	}
	return models.SSLFIPS{}, nil
}

// sslfipskey
func (s *SSLService) CreateSSLFIPSKey(key models.SSLFIPSKey) error {
	payload := map[string]any{"sslfipskey": key}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslFIPSKeyURL+"?action=create", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLFIPSKey(name string) error {
	reqURL := fmt.Sprintf("%s/%s", sslFIPSKeyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) ImportSSLFIPSKey(key models.SSLFIPSKey) error {
	payload := map[string]any{"sslfipskey": key}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslFIPSKeyURL+"?action=import", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) ExportSSLFIPSKey(key models.SSLFIPSKey) error {
	payload := map[string]any{"sslfipskey": key}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslFIPSKeyURL+"?action=export", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLFIPSKey() ([]models.SSLFIPSKey, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslFIPSKeyURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLFIPSKey `json:"sslfipskey"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLFIPSKey(name string) ([]models.SSLFIPSKey, error) {
	reqURL := fmt.Sprintf("%s/%s", sslFIPSKeyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLFIPSKey `json:"sslfipskey"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLFIPSKey() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslFIPSKeyURL+"?count=yes", nil)
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
		} `json:"sslfipskey"`
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

// sslfipssimsource
func (s *SSLService) EnableSSLFIPSSimSource(source models.SSLFIPSSimSource) error {
	payload := map[string]any{"sslfipssimsource": source}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslFIPSSIMSourceURL+"?action=enable", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) InitSSLFIPSSimSource(source models.SSLFIPSSimSource) error {
	payload := map[string]any{"sslfipssimsource": source}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslFIPSSIMSourceURL+"?action=init", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

// sslfipssimtarget
func (s *SSLService) EnableSSLFIPSSimTarget(target models.SSLFIPSSimTarget) error {
	payload := map[string]any{"sslfipssimtarget": target}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslFIPSSIMTargetURL+"?action=enable", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) InitSSLFIPSSimTarget(target models.SSLFIPSSimTarget) error {
	payload := map[string]any{"sslfipssimtarget": target}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslFIPSSIMTargetURL+"?action=init", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

// sslglobal_binding
func (s *SSLService) GetSSLGlobalBinding() ([]models.SSLGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLGlobalBinding `json:"sslglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

// sslglobal_sslpolicy_binding
func (s *SSLService) AddSSLGlobalSSLPolicyBinding(binding models.SSLGlobalSSLPolicyBinding) error {
	payload := map[string]any{"sslglobal_sslpolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, sslGlobalSSLPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLGlobalSSLPolicyBinding(policyname string, typefield string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=type:%s,priority:%d", sslGlobalSSLPolicyBindingURL, url.QueryEscape(policyname), url.QueryEscape(typefield), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetSSLGlobalSSLPolicyBinding() ([]models.SSLGlobalSSLPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslGlobalSSLPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLGlobalSSLPolicyBinding `json:"sslglobal_sslpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLGlobalSSLPolicyBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslGlobalSSLPolicyBindingURL+"?count=yes", nil)
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
		} `json:"sslglobal_sslpolicy_binding"`
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

// sslhsmkey
func (s *SSLService) AddSSLHSMKey(key models.SSLHSMKey) error {
	payload := map[string]any{"sslhsmkey": key}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslHSMKeyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLHSMKey(name string) error {
	reqURL := fmt.Sprintf("%s/%s", sslHSMKeyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLHSMKey() ([]models.SSLHSMKey, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslHSMKeyURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLHSMKey `json:"sslhsmkey"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLHSMKey(name string) ([]models.SSLHSMKey, error) {
	reqURL := fmt.Sprintf("%s/%s", sslHSMKeyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLHSMKey `json:"sslhsmkey"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLHSMKey() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslHSMKeyURL+"?count=yes", nil)
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
		} `json:"sslhsmkey"`
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

// sslkeyfile
func (s *SSLService) ImportSSLKeyFile(file models.SSLKeyFile) error {
	payload := map[string]any{"sslkeyfile": file}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslKeyFileURL+"?action=import", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLKeyFile(name string) error {
	reqURL := fmt.Sprintf("%s/%s", sslKeyFileURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLKeyFile() ([]models.SSLKeyFile, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslKeyFileURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLKeyFile `json:"sslkeyfile"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLKeyFile() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslKeyFileURL+"?count=yes", nil)
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
		} `json:"sslkeyfile"`
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

// ssllogprofile
func (s *SSLService) AddSSLLogProfile(profile models.SSLLogProfile) error {
	payload := map[string]any{"ssllogprofile": profile}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslLogProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLLogProfile(name string) error {
	reqURL := fmt.Sprintf("%s/%s", sslLogProfileURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) UpdateSSLLogProfile(profile models.SSLLogProfile) error {
	payload := map[string]any{"ssllogprofile": profile}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, sslLogProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) UnsetSSLLogProfile(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"ssllogprofile": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslLogProfileURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLLogProfile() ([]models.SSLLogProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslLogProfileURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLLogProfile `json:"ssllogprofile"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLLogProfile(name string) ([]models.SSLLogProfile, error) {
	reqURL := fmt.Sprintf("%s/%s", sslLogProfileURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLLogProfile `json:"ssllogprofile"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLLogProfile() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslLogProfileURL+"?count=yes", nil)
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
		} `json:"ssllogprofile"`
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

// sslocspresponder
func (s *SSLService) AddSSLOCSPResponder(responder models.SSLOCSPResponder) error {
	payload := map[string]any{"sslocspresponder": responder}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslOCSPResponderURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLOCSPResponder(name string) error {
	reqURL := fmt.Sprintf("%s/%s", sslOCSPResponderURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) UpdateSSLOCSPResponder(responder models.SSLOCSPResponder) error {
	payload := map[string]any{"sslocspresponder": responder}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, sslOCSPResponderURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) UnsetSSLOCSPResponder(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"sslocspresponder": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslOCSPResponderURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLOCSPResponder() ([]models.SSLOCSPResponder, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslOCSPResponderURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLOCSPResponder `json:"sslocspresponder"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLOCSPResponder(name string) ([]models.SSLOCSPResponder, error) {
	reqURL := fmt.Sprintf("%s/%s", sslOCSPResponderURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLOCSPResponder `json:"sslocspresponder"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLOCSPResponder() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslOCSPResponderURL+"?count=yes", nil)
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
		} `json:"sslocspresponder"`
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

// sslparameter
func (s *SSLService) UpdateSSLParameter(param models.SSLParameter) error {
	payload := map[string]any{"sslparameter": param}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, sslParameterURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) UnsetSSLParameter(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"sslparameter": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslParameterURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLParameter() (models.SSLParameter, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslParameterURL, nil)
	if err != nil {
		return models.SSLParameter{}, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return models.SSLParameter{}, err
	}
	var result struct {
		Items []models.SSLParameter `json:"sslparameter"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.SSLParameter{}, fmt.Errorf("failed to unmarshal: %w", err)
	}
	if len(result.Items) > 0 {
		return result.Items[0], nil
	}
	return models.SSLParameter{}, nil
}

// sslpkcs12
func (s *SSLService) ConvertSSLPKCS12(pkcs12 models.SSLPKCS12) error {
	payload := map[string]any{"sslpkcs12": pkcs12}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslPKCS12URL+"?action=convert", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

// sslpkcs8
func (s *SSLService) ConvertSSLPKCS8(pkcs8 models.SSLPKCS8) error {
	payload := map[string]any{"sslpkcs8": pkcs8}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslPKCS8URL+"?action=convert", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

// sslpolicy
func (s *SSLService) AddSSLPolicy(policy models.SSLPolicy) error {
	payload := map[string]any{"sslpolicy": policy}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLPolicy(name string) error {
	reqURL := fmt.Sprintf("%s/%s", sslPolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) UpdateSSLPolicy(policy models.SSLPolicy) error {
	payload := map[string]any{"sslpolicy": policy}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, sslPolicyURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) UnsetSSLPolicy(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"sslpolicy": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslPolicyURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLPolicy() ([]models.SSLPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslPolicyURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLPolicy `json:"sslpolicy"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLPolicy(name string) ([]models.SSLPolicy, error) {
	reqURL := fmt.Sprintf("%s/%s", sslPolicyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLPolicy `json:"sslpolicy"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLPolicy() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslPolicyURL+"?count=yes", nil)
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
		} `json:"sslpolicy"`
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

// sslpolicylabel
func (s *SSLService) AddSSLPolicyLabel(label models.SSLPolicyLabel) error {
	payload := map[string]any{"sslpolicylabel": label}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslPolicyLabelURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLPolicyLabel(name string) error {
	reqURL := fmt.Sprintf("%s/%s", sslPolicyLabelURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLPolicyLabel() ([]models.SSLPolicyLabel, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslPolicyLabelURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLPolicyLabel `json:"sslpolicylabel"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLPolicyLabel(name string) ([]models.SSLPolicyLabel, error) {
	reqURL := fmt.Sprintf("%s/%s", sslPolicyLabelURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLPolicyLabel `json:"sslpolicylabel"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLPolicyLabel() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslPolicyLabelURL+"?count=yes", nil)
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
		} `json:"sslpolicylabel"`
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

// sslpolicylabel_binding
func (s *SSLService) GetAllSSLPolicyLabelBinding() ([]models.SSLPolicyLabelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslPolicyLabelBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLPolicyLabelBinding `json:"sslpolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLPolicyLabelBinding(name string) ([]models.SSLPolicyLabelBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslPolicyLabelBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLPolicyLabelBinding `json:"sslpolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

// sslpolicylabel_sslpolicy_binding
func (s *SSLService) AddSSLpolicyLabelSSLPolicyBinding(binding models.SSLPolicyLabelSSLPolicyBinding) error {
	payload := map[string]any{"sslpolicylabel_sslpolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, sslPolicyLabelSSLPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLpolicyLabelSSLPolicyBinding(labelname string, policyname string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s,priority:%d", sslPolicyLabelSSLPolicyBindingURL, url.QueryEscape(labelname), url.QueryEscape(policyname), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLpolicyLabelSSLPolicyBinding() ([]models.SSLPolicyLabelSSLPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslPolicyLabelSSLPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLPolicyLabelSSLPolicyBinding `json:"sslpolicylabel_sslpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLpolicyLabelSSLPolicyBinding(labelname string) ([]models.SSLPolicyLabelSSLPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslPolicyLabelSSLPolicyBindingURL, url.QueryEscape(labelname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLPolicyLabelSSLPolicyBinding `json:"sslpolicylabel_sslpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLpolicyLabelSSLPolicyBinding(labelname string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", sslPolicyLabelSSLPolicyBindingURL, url.QueryEscape(labelname))
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
		} `json:"sslpolicylabel_sslpolicy_binding"`
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

// sslpolicy_binding
func (s *SSLService) GetAllSSLPolicyBinding() ([]models.SSLPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLPolicyBinding `json:"sslpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLPolicyBinding(name string) ([]models.SSLPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslPolicyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLPolicyBinding `json:"sslpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

// sslpolicy_csvserver_binding
func (s *SSLService) GetAllSSLPolicyCSVServerBinding() ([]models.SSLPolicyCSVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslPolicyCSVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLPolicyCSVServerBinding `json:"sslpolicy_csvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLPolicyCSVServerBinding(name string) ([]models.SSLPolicyCSVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslPolicyCSVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLPolicyCSVServerBinding `json:"sslpolicy_csvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLPolicyCSVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", sslPolicyCSVServerBindingURL, url.QueryEscape(name))
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
		} `json:"sslpolicy_csvserver_binding"`
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

// sslpolicy_lbvserver_binding
func (s *SSLService) GetAllSSLPolicyLBVServerBinding() ([]models.SSLPolicyLBVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslPolicyLBVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLPolicyLBVServerBinding `json:"sslpolicy_lbvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLPolicyLBVServerBinding(name string) ([]models.SSLPolicyLBVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslPolicyLBVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLPolicyLBVServerBinding `json:"sslpolicy_lbvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLPolicyLBVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", sslPolicyLBVServerBindingURL, url.QueryEscape(name))
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
		} `json:"sslpolicy_lbvserver_binding"`
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

// sslpolicy_sslglobal_binding
func (s *SSLService) GetAllSSLPolicySSLGlobalBinding() ([]models.SSLPolicySSLGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslPolicySSLGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLPolicySSLGlobalBinding `json:"sslpolicy_sslglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLPolicySSLGlobalBinding(name string) ([]models.SSLPolicySSLGlobalBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslPolicySSLGlobalBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLPolicySSLGlobalBinding `json:"sslpolicy_sslglobal_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLPolicySSLGlobalBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", sslPolicySSLGlobalBindingURL, url.QueryEscape(name))
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
		} `json:"sslpolicy_sslglobal_binding"`
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

// sslpolicy_sslpolicylabel_binding
func (s *SSLService) GetAllSSLPolicySSLPolicyLabelBinding() ([]models.SSLPolicySSLPolicyLabelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslPolicySSLPolicyLabelBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLPolicySSLPolicyLabelBinding `json:"sslpolicy_sslpolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLPolicySSLPolicyLabelBinding(name string) ([]models.SSLPolicySSLPolicyLabelBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslPolicySSLPolicyLabelBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLPolicySSLPolicyLabelBinding `json:"sslpolicy_sslpolicylabel_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLPolicySSLPolicyLabelBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", sslPolicySSLPolicyLabelBindingURL, url.QueryEscape(name))
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
		} `json:"sslpolicy_sslpolicylabel_binding"`
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

// sslpolicy_sslservice_binding
func (s *SSLService) GetAllSSLPolicySSLServiceBinding() ([]models.SSLPolicySSLServiceBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslPolicySSLServiceBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLPolicySSLServiceBinding `json:"sslpolicy_sslservice_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLPolicySSLServiceBinding(name string) ([]models.SSLPolicySSLServiceBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslPolicySSLServiceBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLPolicySSLServiceBinding `json:"sslpolicy_sslservice_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLPolicySSLServiceBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", sslPolicySSLServiceBindingURL, url.QueryEscape(name))
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
		} `json:"sslpolicy_sslservice_binding"`
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

// sslpolicy_sslvserver_binding
func (s *SSLService) GetAllSSLPolicySSLVServerBinding() ([]models.SSLPolicySSLVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslPolicySSLVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLPolicySSLVServerBinding `json:"sslpolicy_sslvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLPolicySSLVServerBinding(name string) ([]models.SSLPolicySSLVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslPolicySSLVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLPolicySSLVServerBinding `json:"sslpolicy_sslvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLPolicySSLVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", sslPolicySSLVServerBindingURL, url.QueryEscape(name))
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
		} `json:"sslpolicy_sslvserver_binding"`
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

// sslprofile
func (s *SSLService) AddSSLProfile(profile models.SSLProfile) error {
	payload := map[string]any{"sslprofile": profile}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLProfile(name string) error {
	reqURL := fmt.Sprintf("%s/%s", sslProfileURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) UpdateSSLProfile(profile models.SSLProfile) error {
	payload := map[string]any{"sslprofile": profile}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, sslProfileURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) UnsetSSLProfile(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"sslprofile": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslProfileURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLProfile() ([]models.SSLProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslProfileURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLProfile `json:"sslprofile"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLProfile(name string) ([]models.SSLProfile, error) {
	reqURL := fmt.Sprintf("%s/%s", sslProfileURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLProfile `json:"sslprofile"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLProfile() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslProfileURL+"?count=yes", nil)
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
		} `json:"sslprofile"`
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

// sslprofile_binding
func (s *SSLService) GetAllSSLProfileBinding() ([]models.SSLProfileBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslProfileBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLProfileBinding `json:"sslprofile_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLProfileBinding(name string) ([]models.SSLProfileBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslProfileBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLProfileBinding `json:"sslprofile_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

// sslprofile_eccurve_binding
func (s *SSLService) AddSSLProfileECCCurveBinding(binding models.SSLProfileECCCurveBinding) error {
	payload := map[string]any{"sslprofile_ecccurve_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, sslProfileECCCurveBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLProfileECCCurveBinding(name string, ecccurvename string) error {
	reqURL := fmt.Sprintf("%s/%s?args=ecccurvename:%s", sslProfileECCCurveBindingURL, url.QueryEscape(name), url.QueryEscape(ecccurvename))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLProfileECCCurveBinding() ([]models.SSLProfileECCCurveBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslProfileECCCurveBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLProfileECCCurveBinding `json:"sslprofile_ecccurve_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLProfileECCCurveBinding(name string) ([]models.SSLProfileECCCurveBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslProfileECCCurveBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLProfileECCCurveBinding `json:"sslprofile_ecccurve_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLProfileECCCurveBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", sslProfileECCCurveBindingURL, url.QueryEscape(name))
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
		} `json:"sslprofile_ecccurve_binding"`
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

// sslprofile_sslcertkey_binding
func (s *SSLService) AddSSLProfileSSLCertKeyBinding(binding models.SSLProfileSSLCertKeyBinding) error {
	payload := map[string]any{"sslprofile_sslcertkey_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, sslProfileSSLCertKeyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLProfileSSLCertKeyBinding(name string, certkeyname string) error {
	reqURL := fmt.Sprintf("%s/%s?args=certkeyname:%s", sslProfileSSLCertKeyBindingURL, url.QueryEscape(name), url.QueryEscape(certkeyname))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLProfileSSLCertKeyBinding() ([]models.SSLProfileSSLCertKeyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslProfileSSLCertKeyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLProfileSSLCertKeyBinding `json:"sslprofile_sslcertkey_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLProfileSSLCertKeyBinding(name string) ([]models.SSLProfileSSLCertKeyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslProfileSSLCertKeyBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLProfileSSLCertKeyBinding `json:"sslprofile_sslcertkey_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLProfileSSLCertKeyBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", sslProfileSSLCertKeyBindingURL, url.QueryEscape(name))
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
		} `json:"sslprofile_sslcertkey_binding"`
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

// sslprofile_sslciphersuite_binding
func (s *SSLService) AddSSLProfileSSLCipherSuiteBinding(binding models.SSLProfileSSLCipherSuiteBinding) error {
	payload := map[string]any{"sslprofile_sslciphersuite_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, sslProfileSSLCipherSuiteBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLProfileSSLCipherSuiteBinding(name string, ciphername string) error {
	reqURL := fmt.Sprintf("%s/%s?args=ciphername:%s", sslProfileSSLCipherSuiteBindingURL, url.QueryEscape(name), url.QueryEscape(ciphername))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLProfileSSLCipherSuiteBinding() ([]models.SSLProfileSSLCipherSuiteBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslProfileSSLCipherSuiteBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLProfileSSLCipherSuiteBinding `json:"sslprofile_sslciphersuite_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLProfileSSLCipherSuiteBinding(name string) ([]models.SSLProfileSSLCipherSuiteBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslProfileSSLCipherSuiteBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLProfileSSLCipherSuiteBinding `json:"sslprofile_sslciphersuite_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLProfileSSLCipherSuiteBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", sslProfileSSLCipherSuiteBindingURL, url.QueryEscape(name))
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
		} `json:"sslprofile_sslciphersuite_binding"`
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

// sslprofile_sslcipher_binding
func (s *SSLService) AddSSLProfileSSLCipherBinding(binding models.SSLProfileSSLCipherBinding) error {
	payload := map[string]any{"sslprofile_sslcipher_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, sslProfileSSLCipherBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLProfileSSLCipherBinding(name string, ciphername string) error {
	reqURL := fmt.Sprintf("%s/%s?args=ciphername:%s", sslProfileSSLCipherBindingURL, url.QueryEscape(name), url.QueryEscape(ciphername))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLProfileSSLCipherBinding() ([]models.SSLProfileSSLCipherBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslProfileSSLCipherBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLProfileSSLCipherBinding `json:"sslprofile_sslcipher_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLProfileSSLCipherBinding(name string) ([]models.SSLProfileSSLCipherBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslProfileSSLCipherBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLProfileSSLCipherBinding `json:"sslprofile_sslcipher_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLProfileSSLCipherBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", sslProfileSSLCipherBindingURL, url.QueryEscape(name))
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
		} `json:"sslprofile_sslcipher_binding"`
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

// sslprofile_sslvserver_binding
func (s *SSLService) GetAllSSLProfileSSLVServerBinding() ([]models.SSLProfileSSLVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslProfileSSLVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLProfileSSLVServerBinding `json:"sslprofile_sslvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLProfileSSLVServerBinding(name string) ([]models.SSLProfileSSLVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslProfileSSLVServerBindingURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLProfileSSLVServerBinding `json:"sslprofile_sslvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLProfileSSLVServerBinding(name string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", sslProfileSSLVServerBindingURL, url.QueryEscape(name))
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
		} `json:"sslprofile_sslvserver_binding"`
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

// sslrsakey
func (s *SSLService) CreateSSLRSAKey(key models.SSLRSAKey) error {
	payload := map[string]any{"sslrsakey": key}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslRSAKeyURL+"?action=create", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

// sslservice
func (s *SSLService) UpdateSSLService(service models.SSLService) error {
	payload := map[string]any{"sslservice": service}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, sslServiceURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) UnsetSSLService(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"sslservice": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslServiceURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLService() ([]models.SSLService, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslServiceURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLService `json:"sslservice"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLService(servicename string) ([]models.SSLService, error) {
	reqURL := fmt.Sprintf("%s/%s", sslServiceURL, url.QueryEscape(servicename))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLService `json:"sslservice"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLService() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslServiceURL+"?count=yes", nil)
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
		} `json:"sslservice"`
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

// sslservicegroup
func (s *SSLService) UpdateSSLServiceGroup(group models.SSLServiceGroup) error {
	payload := map[string]any{"sslservicegroup": group}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, sslServiceGroupURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) UnsetSSLServiceGroup(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"sslservicegroup": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslServiceGroupURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLServiceGroup() ([]models.SSLServiceGroup, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslServiceGroupURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLServiceGroup `json:"sslservicegroup"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLServiceGroup(servicegroupname string) ([]models.SSLServiceGroup, error) {
	reqURL := fmt.Sprintf("%s/%s", sslServiceGroupURL, url.QueryEscape(servicegroupname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLServiceGroup `json:"sslservicegroup"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLServiceGroup() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslServiceGroupURL+"?count=yes", nil)
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
		} `json:"sslservicegroup"`
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

// sslservicegroup_binding
func (s *SSLService) GetAllSSLServiceGroupBinding() ([]models.SSLServiceGroupBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslServiceGroupBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLServiceGroupBinding `json:"sslservicegroup_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLServiceGroupBinding(servicegroupname string) ([]models.SSLServiceGroupBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslServiceGroupBindingURL, url.QueryEscape(servicegroupname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLServiceGroupBinding `json:"sslservicegroup_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

// sslservicegroup_ecccurve_binding
func (s *SSLService) AddSSLServiceGroupECCCurveBinding(binding models.SSLServiceGroupECCCurveBinding) error {
	payload := map[string]any{"sslservicegroup_ecccurve_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, sslServiceGroupECCCurveBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLServiceGroupECCCurveBinding(servicegroupname string, ecccurvename string) error {
	reqURL := fmt.Sprintf("%s/%s?args=ecccurvename:%s", sslServiceGroupECCCurveBindingURL, url.QueryEscape(servicegroupname), url.QueryEscape(ecccurvename))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLServiceGroupECCCurveBinding() ([]models.SSLServiceGroupECCCurveBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslServiceGroupECCCurveBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLServiceGroupECCCurveBinding `json:"sslservicegroup_ecccurve_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLServiceGroupECCCurveBinding(servicegroupname string) ([]models.SSLServiceGroupECCCurveBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslServiceGroupECCCurveBindingURL, url.QueryEscape(servicegroupname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLServiceGroupECCCurveBinding `json:"sslservicegroup_ecccurve_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLServiceGroupECCCurveBinding(servicegroupname string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", sslServiceGroupECCCurveBindingURL, url.QueryEscape(servicegroupname))
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
		} `json:"sslservicegroup_ecccurve_binding"`
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

// sslservicegroup_sslcertkey_binding
func (s *SSLService) AddSSLServiceGroupSSLCertKeyBinding(binding models.SSLServiceGroupSSLCertKeyBinding) error {
	payload := map[string]any{"sslservicegroup_sslcertkey_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, sslServiceGroupSSLCertKeyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLServiceGroupSSLCertKeyBinding(servicegroupname string, certkeyname string) error {
	reqURL := fmt.Sprintf("%s/%s?args=certkeyname:%s", sslServiceGroupSSLCertKeyBindingURL, url.QueryEscape(servicegroupname), url.QueryEscape(certkeyname))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLServiceGroupSSLCertKeyBinding() ([]models.SSLServiceGroupSSLCertKeyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslServiceGroupSSLCertKeyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLServiceGroupSSLCertKeyBinding `json:"sslservicegroup_sslcertkey_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLServiceGroupSSLCertKeyBinding(servicegroupname string) ([]models.SSLServiceGroupSSLCertKeyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslServiceGroupSSLCertKeyBindingURL, url.QueryEscape(servicegroupname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLServiceGroupSSLCertKeyBinding `json:"sslservicegroup_sslcertkey_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLServiceGroupSSLCertKeyBinding(servicegroupname string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", sslServiceGroupSSLCertKeyBindingURL, url.QueryEscape(servicegroupname))
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
		} `json:"sslservicegroup_sslcertkey_binding"`
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

// sslservicegroup_sslciphersuite_binding
func (s *SSLService) AddSSLServiceGroupSSLCipherSuiteBinding(binding models.SSLServiceGroupSSLCipherSuiteBinding) error {
	payload := map[string]any{"sslservicegroup_sslciphersuite_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, sslServiceGroupSSLCipherSuiteBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLServiceGroupSSLCipherSuiteBinding(servicegroupname string, ciphername string) error {
	reqURL := fmt.Sprintf("%s/%s?args=ciphername:%s", sslServiceGroupSSLCipherSuiteBindingURL, url.QueryEscape(servicegroupname), url.QueryEscape(ciphername))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLServiceGroupSSLCipherSuiteBinding() ([]models.SSLServiceGroupSSLCipherSuiteBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslServiceGroupSSLCipherSuiteBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLServiceGroupSSLCipherSuiteBinding `json:"sslservicegroup_sslciphersuite_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLServiceGroupSSLCipherSuiteBinding(servicegroupname string) ([]models.SSLServiceGroupSSLCipherSuiteBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslServiceGroupSSLCipherSuiteBindingURL, url.QueryEscape(servicegroupname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLServiceGroupSSLCipherSuiteBinding `json:"sslservicegroup_sslciphersuite_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLServiceGroupSSLCipherSuiteBinding(servicegroupname string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", sslServiceGroupSSLCipherSuiteBindingURL, url.QueryEscape(servicegroupname))
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
		} `json:"sslservicegroup_sslciphersuite_binding"`
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

// sslservicegroup_sslcipher_binding
func (s *SSLService) AddSSLServiceGroupSSLCipherBinding(binding models.SSLServiceGroupSSLCipherBinding) error {
	payload := map[string]any{"sslservicegroup_sslcipher_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, sslServiceGroupSSLCipherBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLServiceGroupSSLCipherBinding(servicegroupname string, ciphername string) error {
	reqURL := fmt.Sprintf("%s/%s?args=ciphername:%s", sslServiceGroupSSLCipherBindingURL, url.QueryEscape(servicegroupname), url.QueryEscape(ciphername))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLServiceGroupSSLCipherBinding() ([]models.SSLServiceGroupSSLCipherBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslServiceGroupSSLCipherBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLServiceGroupSSLCipherBinding `json:"sslservicegroup_sslcipher_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLServiceGroupSSLCipherBinding(servicegroupname string) ([]models.SSLServiceGroupSSLCipherBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslServiceGroupSSLCipherBindingURL, url.QueryEscape(servicegroupname))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLServiceGroupSSLCipherBinding `json:"sslservicegroup_sslcipher_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLServiceGroupSSLCipherBinding(servicegroupname string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", sslServiceGroupSSLCipherBindingURL, url.QueryEscape(servicegroupname))
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
		} `json:"sslservicegroup_sslcipher_binding"`
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

// sslservice_binding
func (s *SSLService) GetAllSSLServiceBinding() ([]models.SSLServiceBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslServiceBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLServiceBinding `json:"sslservice_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLServiceBinding(servicename string) ([]models.SSLServiceBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslServiceBindingURL, url.QueryEscape(servicename))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLServiceBinding `json:"sslservice_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

// sslservice_ecccurve_binding
func (s *SSLService) AddSSLServiceECCCurveBinding(binding models.SSLServiceECCCurveBinding) error {
	payload := map[string]any{"sslservice_ecccurve_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, sslServiceECCCurveBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLServiceECCCurveBinding(servicename string, ecccurvename string) error {
	reqURL := fmt.Sprintf("%s/%s?args=ecccurvename:%s", sslServiceECCCurveBindingURL, url.QueryEscape(servicename), url.QueryEscape(ecccurvename))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLServiceECCCurveBinding() ([]models.SSLServiceECCCurveBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslServiceECCCurveBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLServiceECCCurveBinding `json:"sslservice_ecccurve_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLServiceECCCurveBinding(servicename string) ([]models.SSLServiceECCCurveBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslServiceECCCurveBindingURL, url.QueryEscape(servicename))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLServiceECCCurveBinding `json:"sslservice_ecccurve_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLServiceECCCurveBinding(servicename string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", sslServiceECCCurveBindingURL, url.QueryEscape(servicename))
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
		} `json:"sslservice_ecccurve_binding"`
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

// sslservice_sslcertkey_binding
func (s *SSLService) AddSSLServiceSSLCertKeyBinding(binding models.SSLServiceSSLCertKeyBinding) error {
	payload := map[string]any{"sslservice_sslcertkey_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, sslServiceSSLCertKeyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLServiceSSLCertKeyBinding(servicename string, certkeyname string) error {
	reqURL := fmt.Sprintf("%s/%s?args=certkeyname:%s", sslServiceSSLCertKeyBindingURL, url.QueryEscape(servicename), url.QueryEscape(certkeyname))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLServiceSSLCertKeyBinding() ([]models.SSLServiceSSLCertKeyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslServiceSSLCertKeyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLServiceSSLCertKeyBinding `json:"sslservice_sslcertkey_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLServiceSSLCertKeyBinding(servicename string) ([]models.SSLServiceSSLCertKeyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslServiceSSLCertKeyBindingURL, url.QueryEscape(servicename))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLServiceSSLCertKeyBinding `json:"sslservice_sslcertkey_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLServiceSSLCertKeyBinding(servicename string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", sslServiceSSLCertKeyBindingURL, url.QueryEscape(servicename))
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
		} `json:"sslservice_sslcertkey_binding"`
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

// sslservice_sslciphersuite_binding
func (s *SSLService) AddSSLServiceSSLCipherSuiteBinding(binding models.SSLServiceSSLCipherSuiteBinding) error {
	payload := map[string]any{"sslservice_sslciphersuite_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, sslServiceSSLCipherSuiteBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLServiceSSLCipherSuiteBinding(servicename string, ciphername string) error {
	reqURL := fmt.Sprintf("%s/%s?args=ciphername:%s", sslServiceSSLCipherSuiteBindingURL, url.QueryEscape(servicename), url.QueryEscape(ciphername))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLServiceSSLCipherSuiteBinding() ([]models.SSLServiceSSLCipherSuiteBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslServiceSSLCipherSuiteBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLServiceSSLCipherSuiteBinding `json:"sslservice_sslciphersuite_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLServiceSSLCipherSuiteBinding(servicename string) ([]models.SSLServiceSSLCipherSuiteBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslServiceSSLCipherSuiteBindingURL, url.QueryEscape(servicename))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLServiceSSLCipherSuiteBinding `json:"sslservice_sslciphersuite_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLServiceSSLCipherSuiteBinding(servicename string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", sslServiceSSLCipherSuiteBindingURL, url.QueryEscape(servicename))
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
		} `json:"sslservice_sslciphersuite_binding"`
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

// sslservice_sslcipher_binding
func (s *SSLService) AddSSLServiceSSLCipherBinding(binding models.SSLServiceSSLCipherBinding) error {
	payload := map[string]any{"sslservice_sslcipher_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, sslServiceSSLCipherBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLServiceSSLCipherBinding(servicename string, ciphername string) error {
	reqURL := fmt.Sprintf("%s/%s?args=ciphername:%s", sslServiceSSLCipherBindingURL, url.QueryEscape(servicename), url.QueryEscape(ciphername))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLServiceSSLCipherBinding() ([]models.SSLServiceSSLCipherBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslServiceSSLCipherBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLServiceSSLCipherBinding `json:"sslservice_sslcipher_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLServiceSSLCipherBinding(servicename string) ([]models.SSLServiceSSLCipherBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslServiceSSLCipherBindingURL, url.QueryEscape(servicename))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLServiceSSLCipherBinding `json:"sslservice_sslcipher_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLServiceSSLCipherBinding(servicename string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", sslServiceSSLCipherBindingURL, url.QueryEscape(servicename))
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
		} `json:"sslservice_sslcipher_binding"`
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

// sslservice_sslpolicy_binding
func (s *SSLService) AddSSLServiceSSLPolicyBinding(binding models.SSLServiceSSLPolicyBinding) error {
	payload := map[string]any{"sslservice_sslpolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, sslServiceSSLPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLServiceSSLPolicyBinding(servicename string, policyname string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s,priority:%d", sslServiceSSLPolicyBindingURL, url.QueryEscape(servicename), url.QueryEscape(policyname), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLServiceSSLPolicyBinding() ([]models.SSLServiceSSLPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslServiceSSLPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLServiceSSLPolicyBinding `json:"sslservice_sslpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLServiceSSLPolicyBinding(servicename string) ([]models.SSLServiceSSLPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslServiceSSLPolicyBindingURL, url.QueryEscape(servicename))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLServiceSSLPolicyBinding `json:"sslservice_sslpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLServiceSSLPolicyBinding(servicename string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", sslServiceSSLPolicyBindingURL, url.QueryEscape(servicename))
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
		} `json:"sslservice_sslpolicy_binding"`
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

// sslvserver
func (s *SSLService) UpdateSSLVServer(vserver models.SSLVServer) error {
	payload := map[string]any{"sslvserver": vserver}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, sslVServerURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) UnsetSSLVServer(args []string) error {
	unsetMap := make(map[string]any)
	for _, arg := range args {
		unsetMap[arg] = true
	}
	payload := map[string]any{"sslvserver": unsetMap}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslVServerURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLVServer() ([]models.SSLVServer, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslVServerURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLVServer `json:"sslvserver"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLVServer(vservername string) ([]models.SSLVServer, error) {
	reqURL := fmt.Sprintf("%s/%s", sslVServerURL, url.QueryEscape(vservername))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLVServer `json:"sslvserver"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLVServer() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslVServerURL+"?count=yes", nil)
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
		} `json:"sslvserver"`
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

// sslvserver_binding
func (s *SSLService) GetAllSSLVServerBinding() ([]models.SSLVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLVServerBinding `json:"sslvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLVServerBinding(vservername string) ([]models.SSLVServerBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslVServerBindingURL, url.QueryEscape(vservername))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLVServerBinding `json:"sslvserver_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

// sslvserver_ecccurve_binding
func (s *SSLService) AddSSLVServerECCCurveBinding(binding models.SSLVServerECCCurveBinding) error {
	payload := map[string]any{"sslvserver_ecccurve_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, sslVServerECCCurveBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLVServerECCCurveBinding(vservername string, ecccurvename string) error {
	reqURL := fmt.Sprintf("%s/%s?args=ecccurvename:%s", sslVServerECCCurveBindingURL, url.QueryEscape(vservername), url.QueryEscape(ecccurvename))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLVServerECCCurveBinding() ([]models.SSLVServerECCCurveBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslVServerECCCurveBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLVServerECCCurveBinding `json:"sslvserver_ecccurve_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLVServerECCCurveBinding(vservername string) ([]models.SSLVServerECCCurveBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslVServerECCCurveBindingURL, url.QueryEscape(vservername))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLVServerECCCurveBinding `json:"sslvserver_ecccurve_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLVServerECCCurveBinding(vservername string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", sslVServerECCCurveBindingURL, url.QueryEscape(vservername))
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
		} `json:"sslvserver_ecccurve_binding"`
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

// sslvserver_sslcertkey_binding
func (s *SSLService) AddSSLVServerSSLCertKeyBinding(binding models.SSLVServerSSLCertKeyBinding) error {
	payload := map[string]any{"sslvserver_sslcertkey_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, sslVServerSSLCertKeyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLVServerSSLCertKeyBinding(vservername string, certkeyname string) error {
	reqURL := fmt.Sprintf("%s/%s?args=certkeyname:%s", sslVServerSSLCertKeyBindingURL, url.QueryEscape(vservername), url.QueryEscape(certkeyname))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLVServerSSLCertKeyBinding() ([]models.SSLVServerSSLCertKeyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslVServerSSLCertKeyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLVServerSSLCertKeyBinding `json:"sslvserver_sslcertkey_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLVServerSSLCertKeyBinding(vservername string) ([]models.SSLVServerSSLCertKeyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslVServerSSLCertKeyBindingURL, url.QueryEscape(vservername))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLVServerSSLCertKeyBinding `json:"sslvserver_sslcertkey_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLVServerSSLCertKeyBinding(vservername string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", sslVServerSSLCertKeyBindingURL, url.QueryEscape(vservername))
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
		} `json:"sslvserver_sslcertkey_binding"`
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

// sslvserver_sslciphersuite_binding
func (s *SSLService) AddSSLVServerSSLCipherSuiteBinding(binding models.SSLVServerSSLCipherSuiteBinding) error {
	payload := map[string]any{"sslvserver_sslciphersuite_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, sslVServerSSLCipherSuiteBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLVServerSSLCipherSuiteBinding(vservername string, ciphername string) error {
	reqURL := fmt.Sprintf("%s/%s?args=ciphername:%s", sslVServerSSLCipherSuiteBindingURL, url.QueryEscape(vservername), url.QueryEscape(ciphername))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLVServerSSLCipherSuiteBinding() ([]models.SSLVServerSSLCipherSuiteBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslVServerSSLCipherSuiteBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLVServerSSLCipherSuiteBinding `json:"sslvserver_sslciphersuite_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLVServerSSLCipherSuiteBinding(vservername string) ([]models.SSLVServerSSLCipherSuiteBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslVServerSSLCipherSuiteBindingURL, url.QueryEscape(vservername))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLVServerSSLCipherSuiteBinding `json:"sslvserver_sslciphersuite_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLVServerSSLCipherSuiteBinding(vservername string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", sslVServerSSLCipherSuiteBindingURL, url.QueryEscape(vservername))
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
		} `json:"sslvserver_sslciphersuite_binding"`
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

// sslvserver_sslcipher_binding
func (s *SSLService) AddSSLVServerSSLCipherBinding(binding models.SSLVServerSSLCipherBinding) error {
	payload := map[string]any{"sslvserver_sslcipher_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, sslVServerSSLCipherBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLVServerSSLCipherBinding(vservername string, ciphername string) error {
	reqURL := fmt.Sprintf("%s/%s?args=ciphername:%s", sslVServerSSLCipherBindingURL, url.QueryEscape(vservername), url.QueryEscape(ciphername))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLVServerSSLCipherBinding() ([]models.SSLVServerSSLCipherBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslVServerSSLCipherBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLVServerSSLCipherBinding `json:"sslvserver_sslcipher_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLVServerSSLCipherBinding(vservername string) ([]models.SSLVServerSSLCipherBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslVServerSSLCipherBindingURL, url.QueryEscape(vservername))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLVServerSSLCipherBinding `json:"sslvserver_sslcipher_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLVServerSSLCipherBinding(vservername string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", sslVServerSSLCipherBindingURL, url.QueryEscape(vservername))
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
		} `json:"sslvserver_sslcipher_binding"`
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

// sslvserver_sslpolicy_binding
func (s *SSLService) AddSSLVServerSSLPolicyBinding(binding models.SSLVServerSSLPolicyBinding) error {
	payload := map[string]any{"sslvserver_sslpolicy_binding": binding}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPut, sslVServerSSLPolicyBindingURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLVServerSSLPolicyBinding(vservername string, policyname string, typefield string, priority int) error {
	reqURL := fmt.Sprintf("%s/%s?args=policyname:%s,type:%s,priority:%d", sslVServerSSLPolicyBindingURL, url.QueryEscape(vservername), url.QueryEscape(policyname), url.QueryEscape(typefield), priority)
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLVServerSSLPolicyBinding() ([]models.SSLVServerSSLPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslVServerSSLPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLVServerSSLPolicyBinding `json:"sslvserver_sslpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) GetSSLVServerSSLPolicyBinding(vservername string) ([]models.SSLVServerSSLPolicyBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", sslVServerSSLPolicyBindingURL, url.QueryEscape(vservername))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLVServerSSLPolicyBinding `json:"sslvserver_sslpolicy_binding"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLVServerSSLPolicyBinding(vservername string) (float64, error) {
	reqURL := fmt.Sprintf("%s/%s?count=yes", sslVServerSSLPolicyBindingURL, url.QueryEscape(vservername))
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
		} `json:"sslvserver_sslpolicy_binding"`
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

// sslwrapkey
func (s *SSLService) CreateSSLWrapKey(key models.SSLWrapKey) error {
	payload := map[string]any{"sslwrapkey": key}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	req, err := s.client.NewRequest(http.MethodPost, sslWrapKeyURL+"?action=create", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) DeleteSSLWrapKey(name string) error {
	reqURL := fmt.Sprintf("%s/%s", sslWrapKeyURL, url.QueryEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(req)
	return err
}

func (s *SSLService) GetAllSSLWrapKey() ([]models.SSLWrapKey, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslWrapKeyURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Items []models.SSLWrapKey `json:"sslwrapkey"`
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}
	return result.Items, nil
}

func (s *SSLService) CountSSLWrapKey() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, sslWrapKeyURL+"?count=yes", nil)
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
		} `json:"sslwrapkey"`
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
