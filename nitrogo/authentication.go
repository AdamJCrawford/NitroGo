package nitrogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AdamJCrawford/NitroGo/nitrogo/models"
)

const (
	authenticationADFSProxyProfileURL                              = "/nitro/v1/config/authenticationadfsproxyprofile"
	authenticationAuthnProfileURL                                  = "/nitro/v1/config/authenticationauthnprofile"
	authenticationAzureKeyVaultURL                                 = "/nitro/v1/config/authenticationazurekeyvault"
	authenticationCaptchaActionURL                                 = "/nitro/v1/config/authenticationcaptchaaction"
	authenticationCertActionURL                                    = "/nitro/v1/config/authenticationcertaction"
	authenticationCertPolicyURL                                    = "/nitro/v1/config/authenticationcertpolicy"
	authenticationCertPolicyAuthenticationVServerBindingURL        = "/nitro/v1/config/authenticationcertpolicy_authenticationvserver_binding"
	authenticationCertPolicyBindingURL                             = "/nitro/v1/config/authenticationcertpolicy_binding"
	authenticationCertPolicyVPNGlobalBindingURL                    = "/nitro/v1/config/authenticationcertpolicy_vpnglobal_binding"
	authenticationCertPolicyVPNVServerBindingURL                   = "/nitro/v1/config/authenticationcertpolicy_vpnvserver_binding"
	authenticationCitrixAuthActionURL                              = "/nitro/v1/config/authenticationcitrixauthaction"
	authenticationDFAActionURL                                     = "/nitro/v1/config/authenticationdfaaction"
	authenticationDFAPolicyURL                                     = "/nitro/v1/config/authenticationdfapolicy"
	authenticationDFAPolicyBindingURL                              = "/nitro/v1/config/authenticationdfapolicy_binding"
	authenticationDFAPolicyVPNVServerBindingURL                    = "/nitro/v1/config/authenticationdfapolicy_vpnvserver_binding"
	authenticationEmailActionURL                                   = "/nitro/v1/config/authenticationemailaction"
	authenticationEPAActionURL                                     = "/nitro/v1/config/authenticationepaaction"
	authenticationLDAPActionURL                                    = "/nitro/v1/config/authenticationldapaction"
	authenticationLDAPPolicyURL                                    = "/nitro/v1/config/authenticationldappolicy"
	authenticationLDAPPolicyAuthenticationVServerBindingURL        = "/nitro/v1/config/authenticationldappolicy_authenticationvserver_binding"
	authenticationLDAPPolicyBindingURL                             = "/nitro/v1/config/authenticationldappolicy_binding"
	authenticationLDAPPolicySystemGlobalBindingURL                 = "/nitro/v1/config/authenticationldappolicy_systemglobal_binding"
	authenticationLDAPPolicyVPNGlobalBindingURL                    = "/nitro/v1/config/authenticationldappolicy_vpnglobal_binding"
	authenticationLDAPPolicyVPNVServerBindingURL                   = "/nitro/v1/config/authenticationldappolicy_vpnvserver_binding"
	authenticationLocalPolicyURL                                   = "/nitro/v1/config/authenticationlocalpolicy"
	authenticationLocalPolicyAuthenticationVServerBindingURL       = "/nitro/v1/config/authenticationlocalpolicy_authenticationvserver_binding"
	authenticationLocalPolicyBindingURL                            = "/nitro/v1/config/authenticationlocalpolicy_binding"
	authenticationLocalPolicySystemGlobalBindingURL                = "/nitro/v1/config/authenticationlocalpolicy_systemglobal_binding"
	authenticationLocalPolicyVPNGlobalBindingURL                   = "/nitro/v1/config/authenticationlocalpolicy_vpnglobal_binding"
	authenticationLocalPolicyVPNVServerBindingURL                  = "/nitro/v1/config/authenticationlocalpolicy_vpnvserver_binding"
	authenticationLoginSchemaURL                                   = "/nitro/v1/config/authenticationloginschema"
	authenticationLoginSchemaPolicyURL                             = "/nitro/v1/config/authenticationloginschemapolicy"
	authenticationLoginSchemaPolicyAuthenticationVServerBindingURL = "/nitro/v1/config/authenticationloginschemapolicy_authenticationvserver_binding"
	authenticationLoginSchemaPolicyBindingURL                      = "/nitro/v1/config/authenticationloginschemapolicy_binding"
	authenticationLoginSchemaPolicyVPNVServerBindingURL            = "/nitro/v1/config/authenticationloginschemapolicy_vpnvserver_binding"
	authenticationNegotiateActionURL                               = "/nitro/v1/config/authenticationnegotiateaction"
	authenticationNegotiatePolicyURL                               = "/nitro/v1/config/authenticationnegotiatepolicy"
	authenticationNegotiatePolicyAuthenticationVServerBindingURL   = "/nitro/v1/config/authenticationnegotiatepolicy_authenticationvserver_binding"
	authenticationNegotiatePolicyBindingURL                        = "/nitro/v1/config/authenticationnegotiatepolicy_binding"
	authenticationNegotiatePolicyVPNGlobalBindingURL               = "/nitro/v1/config/authenticationnegotiatepolicy_vpnglobal_binding"
	authenticationNegotiatePolicyVPNVServerBindingURL              = "/nitro/v1/config/authenticationnegotiatepolicy_vpnvserver_binding"
	authenticationNoAuthActionURL                                  = "/nitro/v1/config/authenticationnoauthaction"
	authenticationOAuthActionURL                                   = "/nitro/v1/config/authenticationoauthaction"
	authenticationOAuthIdPPolicyURL                                = "/nitro/v1/config/authenticationoauthidppolicy"
	authenticationOAuthIdPPolicyAuthenticationVServerBindingURL    = "/nitro/v1/config/authenticationoauthidppolicy_authenticationvserver_binding"
	authenticationOAuthIdPPolicyBindingURL                         = "/nitro/v1/config/authenticationoauthidppolicy_binding"
	authenticationOAuthIdPPolicyVPNVServerBindingURL               = "/nitro/v1/config/authenticationoauthidppolicy_vpnvserver_binding"
	authenticationOAuthIdPProfileURL                               = "/nitro/v1/config/authenticationoauthidpprofile"
	authenticationPolicyURL                                        = "/nitro/v1/config/authenticationpolicy"
	authenticationPolicyLabelURL                                   = "/nitro/v1/config/authenticationpolicylabel"
	authenticationPolicyLabelAuthenticationPolicyBindingURL        = "/nitro/v1/config/authenticationpolicylabel_authenticationpolicy_binding"
	authenticationPolicyLabelBindingURL                            = "/nitro/v1/config/authenticationpolicylabel_binding"
	authenticationPolicyAuthenticationPolicyLabelBindingURL        = "/nitro/v1/config/authenticationpolicy_authenticationpolicylabel_binding"
	authenticationPolicyAuthenticationVServerBindingURL            = "/nitro/v1/config/authenticationpolicy_authenticationvserver_binding"
	authenticationPolicyBindingURL                                 = "/nitro/v1/config/authenticationpolicy_binding"
	authenticationPolicySystemGlobalBindingURL                     = "/nitro/v1/config/authenticationpolicy_systemglobal_binding"
	authenticationPushServiceURL                                   = "/nitro/v1/config/authenticationpushservice"
	authenticationRADIUSActionURL                                  = "/nitro/v1/config/authenticationradiusaction"
	authenticationRADIUSPolicyURL                                  = "/nitro/v1/config/authenticationradiuspolicy"
	authenticationRADIUSPolicyAuthenticationVServerBindingURL      = "/nitro/v1/config/authenticationradiuspolicy_authenticationvserver_binding"
	authenticationRADIUSPolicyBindingURL                           = "/nitro/v1/config/authenticationradiuspolicy_binding"
	authenticationRADIUSPolicySystemGlobalBindingURL               = "/nitro/v1/config/authenticationradiuspolicy_systemglobal_binding"
	authenticationRADIUSPolicyVPNGlobalBindingURL                  = "/nitro/v1/config/authenticationradiuspolicy_vpnglobal_binding"
	authenticationRADIUSPolicyVPNVServerBindingURL                 = "/nitro/v1/config/authenticationradiuspolicy_vpnvserver_binding"
	authenticationSAMLActionURL                                    = "/nitro/v1/config/authenticationsamlaction"
	authenticationSAMLIdPPolicyURL                                 = "/nitro/v1/config/authenticationsamlidppolicy"
	authenticationSAMLIdPPolicyAuthenticationVServerBindingURL     = "/nitro/v1/config/authenticationsamlidppolicy_authenticationvserver_binding"
	authenticationSAMLIdPPolicyBindingURL                          = "/nitro/v1/config/authenticationsamlidppolicy_binding"
	authenticationSAMLIdPPolicyVPNVServerBindingURL                = "/nitro/v1/config/authenticationsamlidppolicy_vpnvserver_binding"
	authenticationSAMLIdPProfileURL                                = "/nitro/v1/config/authenticationsamlidpprofile"
	authenticationSAMLPolicyURL                                    = "/nitro/v1/config/authenticationsamlpolicy"
	authenticationSAMLPolicyAuthenticationVServerBindingURL        = "/nitro/v1/config/authenticationsamlpolicy_authenticationvserver_binding"
	authenticationSAMLPolicyBindingURL                             = "/nitro/v1/config/authenticationsamlpolicy_binding"
	authenticationSAMLPolicyVPNGlobalBindingURL                    = "/nitro/v1/config/authenticationsamlpolicy_vpnglobal_binding"
	authenticationSAMLPolicyVPNVServerBindingURL                   = "/nitro/v1/config/authenticationsamlpolicy_vpnvserver_binding"
	authenticationStoreFrontAuthActionURL                          = "/nitro/v1/config/authenticationstorefrontauthaction"
	authenticationTACACSActionURL                                  = "/nitro/v1/config/authenticationtacacsaction"
	authenticationTACACSPolicyURL                                  = "/nitro/v1/config/authenticationtacacspolicy"
	authenticationTACACSPolicyAuthenticationVServerBindingURL      = "/nitro/v1/config/authenticationtacacspolicy_authenticationvserver_binding"
	authenticationTACACSPolicyBindingURL                           = "/nitro/v1/config/authenticationtacacspolicy_binding"
	authenticationTACACSPolicySystemGlobalBindingURL               = "/nitro/v1/config/authenticationtacacspolicy_systemglobal_binding"
	authenticationTACACSPolicyVPNGlobalBindingURL                  = "/nitro/v1/config/authenticationtacacspolicy_vpnglobal_binding"
	authenticationTACACSPolicyVPNVServerBindingURL                 = "/nitro/v1/config/authenticationtacacspolicy_vpnvserver_binding"
	authenticationVServerURL                                       = "/nitro/v1/config/authenticationvserver"
	authenticationVServerAuditNSLogPolicyBindingURL                = "/nitro/v1/config/authenticationvserver_auditnslogpolicy_binding"
	authenticationVServerAuditSyslogPolicyBindingURL               = "/nitro/v1/config/authenticationvserver_auditsyslogpolicy_binding"
	authenticationVServerAuthenticationCertPolicyBindingURL        = "/nitro/v1/config/authenticationvserver_authenticationcertpolicy_binding"
	authenticationVServerAuthenticationLDAPPolicyBindingURL        = "/nitro/v1/config/authenticationvserver_authenticationldappolicy_binding"
	authenticationVServerAuthenticationLocalPolicyBindingURL       = "/nitro/v1/config/authenticationvserver_authenticationlocalpolicy_binding"
	authenticationVServerAuthenticationLoginSchemaPolicyBindingURL = "/nitro/v1/config/authenticationvserver_authenticationloginschemapolicy_binding"
	authenticationVServerAuthenticationNegotiatePolicyBindingURL   = "/nitro/v1/config/authenticationvserver_authenticationnegotiatepolicy_binding"
	authenticationVServerAuthenticationOAuthIdPPolicyBindingURL    = "/nitro/v1/config/authenticationvserver_authenticationoauthidppolicy_binding"
	authenticationVServerAuthenticationPolicyBindingURL            = "/nitro/v1/config/authenticationvserver_authenticationpolicy_binding"
	authenticationVServerAuthenticationRADIUSPolicyBindingURL      = "/nitro/v1/config/authenticationvserver_authenticationradiuspolicy_binding"
	authenticationVServerAuthenticationSAMLIdPPolicyBindingURL     = "/nitro/v1/config/authenticationvserver_authenticationsamlidppolicy_binding"
	authenticationVServerAuthenticationSAMLPolicyBindingURL        = "/nitro/v1/config/authenticationvserver_authenticationsamlpolicy_binding"
	authenticationVServerAuthenticationTACACSPolicyBindingURL      = "/nitro/v1/config/authenticationvserver_authenticationtacacspolicy_binding"
	authenticationVServerAuthenticationWebAuthPolicyBindingURL     = "/nitro/v1/config/authenticationvserver_authenticationwebauthpolicy_binding"
	authenticationVServerBindingURL                                = "/nitro/v1/config/authenticationvserver_binding"
	authenticationVServerCachePolicyBindingURL                     = "/nitro/v1/config/authenticationvserver_cachepolicy_binding"
	authenticationVServerCSPolicyBindingURL                        = "/nitro/v1/config/authenticationvserver_cspolicy_binding"
	authenticationVServerResponderPolicyBindingURL                 = "/nitro/v1/config/authenticationvserver_responderpolicy_binding"
	authenticationVServerTMSessionPolicyBindingURL                 = "/nitro/v1/config/authenticationvserver_tmsessionpolicy_binding"
	authenticationVServerVPNPortalThemeBindingURL                  = "/nitro/v1/config/authenticationvserver_vpnportaltheme_binding"
	authenticationWebAuthActionURL                                 = "/nitro/v1/config/authenticationwebauthaction"
	authenticationWebAuthPolicyURL                                 = "/nitro/v1/config/authenticationwebauthpolicy"
	authenticationWebAuthPolicyAuthenticationVServerBindingURL     = "/nitro/v1/config/authenticationwebauthpolicy_authenticationvserver_binding"
	authenticationWebAuthPolicyBindingURL                          = "/nitro/v1/config/authenticationwebauthpolicy_binding"
	authenticationWebAuthPolicySystemGlobalBindingURL              = "/nitro/v1/config/authenticationwebauthpolicy_systemglobal_binding"
	authenticationWebAuthPolicyVPNGlobalBindingURL                 = "/nitro/v1/config/authenticationwebauthpolicy_vpnglobal_binding"
	authenticationWebAuthPolicyVPNVServerBindingURL                = "/nitro/v1/config/authenticationwebauthpolicy_vpnvserver_binding"
)

// Authentication configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authentication
type AuthenticationService struct {
	client *Client
}

// authenticationadfsproxyprofile
// Configuration for ADFSProxy Profile resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationadfsproxyprofile
func (s *AuthenticationService) AddAuthenticationADFSProxyProfile(resource models.AuthenticationADFSProxyProfile) error {
	payload := map[string]any{"authenticationadfsproxyprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationADFSProxyProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationADFSProxyProfile(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", authenticationADFSProxyProfileURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UpdateAuthenticationADFSProxyProfile(resource models.AuthenticationADFSProxyProfile) error {
	payload := map[string]any{"authenticationadfsproxyprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, authenticationADFSProxyProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationADFSProxyProfile() ([]models.AuthenticationADFSProxyProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationADFSProxyProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationADFSProxyProfile `json:"authenticationadfsproxyprofile"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationADFSProxyProfile(name string) (*models.AuthenticationADFSProxyProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationADFSProxyProfileURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationADFSProxyProfile `json:"authenticationadfsproxyprofile"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationadfsproxyprofile %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationADFSProxyProfile() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationADFSProxyProfileURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationADFSProxyProfile `json:"authenticationadfsproxyprofile"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationauthnprofile
// Configuration for Authentication profile resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationauthnprofile
func (s *AuthenticationService) AddAuthenticationAuthnProfile(resource models.AuthenticationAuthNProfile) error {
	payload := map[string]any{"authenticationauthnprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationAuthnProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationAuthnProfile(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", authenticationAuthnProfileURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UpdateAuthenticationAuthnProfile(resource models.AuthenticationAuthNProfile) error {
	payload := map[string]any{"authenticationauthnprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, authenticationAuthnProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UnsetAuthenticationAuthnProfile(resource models.AuthenticationAuthNProfile) error {
	payload := map[string]any{"authenticationauthnprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", authenticationAuthnProfileURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationAuthnProfile() ([]models.AuthenticationAuthNProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationAuthnProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationAuthNProfile `json:"authenticationauthnprofile"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationAuthnProfile(name string) (*models.AuthenticationAuthNProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationAuthnProfileURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationAuthNProfile `json:"authenticationauthnprofile"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationauthnprofile %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationAuthnProfile() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationAuthnProfileURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationAuthNProfile `json:"authenticationauthnprofile"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationazurekeyvault
// Configuration for Azure Key Vault entity resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationazurekeyvault
func (s *AuthenticationService) AddAuthenticationAzureKeyVault(resource models.AuthenticationAzureKeyVault) error {
	payload := map[string]any{"authenticationazurekeyvault": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationAzureKeyVaultURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationAzureKeyVault(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", authenticationAzureKeyVaultURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UpdateAuthenticationAzureKeyVault(resource models.AuthenticationAzureKeyVault) error {
	payload := map[string]any{"authenticationazurekeyvault": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, authenticationAzureKeyVaultURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UnsetAuthenticationAzureKeyVault(resource models.AuthenticationAzureKeyVault) error {
	payload := map[string]any{"authenticationazurekeyvault": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", authenticationAzureKeyVaultURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationAzureKeyVault() ([]models.AuthenticationAzureKeyVault, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationAzureKeyVaultURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationAzureKeyVault `json:"authenticationazurekeyvault"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationAzureKeyVault(name string) (*models.AuthenticationAzureKeyVault, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationAzureKeyVaultURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationAzureKeyVault `json:"authenticationazurekeyvault"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationazurekeyvault %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationAzureKeyVault() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationAzureKeyVaultURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationAzureKeyVault `json:"authenticationazurekeyvault"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationcaptchaaction
// Configuration for Captcha Action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationcaptchaaction
func (s *AuthenticationService) AddAuthenticationCaptchaAction(resource models.AuthenticationCaptchaAction) error {
	payload := map[string]any{"authenticationcaptchaaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationCaptchaActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationCaptchaAction(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", authenticationCaptchaActionURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UpdateAuthenticationCaptchaAction(resource models.AuthenticationCaptchaAction) error {
	payload := map[string]any{"authenticationcaptchaaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, authenticationCaptchaActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UnsetAuthenticationCaptchaAction(resource models.AuthenticationCaptchaAction) error {
	payload := map[string]any{"authenticationcaptchaaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", authenticationCaptchaActionURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationCaptchaAction() ([]models.AuthenticationCaptchaAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationCaptchaActionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationCaptchaAction `json:"authenticationcaptchaaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationCaptchaAction(name string) (*models.AuthenticationCaptchaAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationCaptchaActionURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationCaptchaAction `json:"authenticationcaptchaaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationcaptchaaction %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationCaptchaAction() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationCaptchaActionURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationCaptchaAction `json:"authenticationcaptchaaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationcertaction
// Configuration for CERT action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationcertaction
func (s *AuthenticationService) AddAuthenticationCertAction(resource models.AuthenticationCertAction) error {
	payload := map[string]any{"authenticationcertaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationCertActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationCertAction(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", authenticationCertActionURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UpdateAuthenticationCertAction(resource models.AuthenticationCertAction) error {
	payload := map[string]any{"authenticationcertaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, authenticationCertActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UnsetAuthenticationCertAction(resource models.AuthenticationCertAction) error {
	payload := map[string]any{"authenticationcertaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", authenticationCertActionURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationCertAction() ([]models.AuthenticationCertAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationCertActionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationCertAction `json:"authenticationcertaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationCertAction(name string) (*models.AuthenticationCertAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationCertActionURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationCertAction `json:"authenticationcertaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationcertaction %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationCertAction() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationCertActionURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationCertAction `json:"authenticationcertaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationcertpolicy
// Configuration for CERT policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationcertpolicy
func (s *AuthenticationService) AddAuthenticationCertPolicy(resource models.AuthenticationCertPolicy) error {
	payload := map[string]any{"authenticationcertpolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationCertPolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationCertPolicy(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", authenticationCertPolicyURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UpdateAuthenticationCertPolicy(resource models.AuthenticationCertPolicy) error {
	payload := map[string]any{"authenticationcertpolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, authenticationCertPolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UnsetAuthenticationCertPolicy(resource models.AuthenticationCertPolicy) error {
	payload := map[string]any{"authenticationcertpolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", authenticationCertPolicyURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationCertPolicy() ([]models.AuthenticationCertPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationCertPolicyURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationCertPolicy `json:"authenticationcertpolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationCertPolicy(name string) (*models.AuthenticationCertPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationCertPolicyURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationCertPolicy `json:"authenticationcertpolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationcertpolicy %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationCertPolicy() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationCertPolicyURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationCertPolicy `json:"authenticationcertpolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationcertpolicy_authenticationvserver_binding
// Binding object showing the authenticationvserver that can be bound to authenticationcertpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationcertpolicy_authenticationvserver_binding
func (s *AuthenticationService) GetAllAuthenticationCertPolicyAuthenticationVServerBinding() ([]models.AuthenticationCertPolicyAuthenticationVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationCertPolicyAuthenticationVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationCertPolicyAuthenticationVServerBinding `json:"authenticationcertpolicy_authenticationvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationCertPolicyAuthenticationVServerBinding(name string) (*models.AuthenticationCertPolicyAuthenticationVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationCertPolicyAuthenticationVServerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationCertPolicyAuthenticationVServerBinding `json:"authenticationcertpolicy_authenticationvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationcertpolicy_authenticationvserver_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationCertPolicyAuthenticationVServerBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationCertPolicyAuthenticationVServerBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationCertPolicyAuthenticationVServerBinding `json:"authenticationcertpolicy_authenticationvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationcertpolicy_binding
// Binding object which returns the resources bound to authenticationcertpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationcertpolicy_binding
func (s *AuthenticationService) GetAllAuthenticationCertPolicyBinding() ([]models.AuthenticationCertPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationCertPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationCertPolicyBinding `json:"authenticationcertpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationCertPolicyBinding(name string) (*models.AuthenticationCertPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationCertPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationCertPolicyBinding `json:"authenticationcertpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationcertpolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

// authenticationcertpolicy_vpnglobal_binding
// Binding object showing the vpnglobal that can be bound to authenticationcertpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationcertpolicy_vpnglobal_binding
func (s *AuthenticationService) GetAllAuthenticationCertPolicyVPNGlobalBinding() ([]models.AuthenticationCertPolicyVPNGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationCertPolicyVPNGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationCertPolicyVPNGlobalBinding `json:"authenticationcertpolicy_vpnglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationCertPolicyVPNGlobalBinding(name string) (*models.AuthenticationCertPolicyVPNGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationCertPolicyVPNGlobalBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationCertPolicyVPNGlobalBinding `json:"authenticationcertpolicy_vpnglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationcertpolicy_vpnglobal_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationCertPolicyVPNGlobalBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationCertPolicyVPNGlobalBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationCertPolicyVPNGlobalBinding `json:"authenticationcertpolicy_vpnglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationcertpolicy_vpnvserver_binding
// Binding object showing the vpnvserver that can be bound to authenticationcertpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationcertpolicy_vpnvserver_binding
func (s *AuthenticationService) GetAllAuthenticationCertPolicyVPNVServerBinding() ([]models.AuthenticationCertPolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationCertPolicyVPNVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationCertPolicyVPNVServerBinding `json:"authenticationcertpolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationCertPolicyVPNVServerBinding(name string) (*models.AuthenticationCertPolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationCertPolicyVPNVServerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationCertPolicyVPNVServerBinding `json:"authenticationcertpolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationcertpolicy_vpnvserver_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationCertPolicyVPNVServerBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationCertPolicyVPNVServerBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationCertPolicyVPNVServerBinding `json:"authenticationcertpolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationcitrixauthaction
// Configuration for Citrix Authentication action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationcitrixauthaction
func (s *AuthenticationService) AddAuthenticationCitrixAuthAction(resource models.AuthenticationCitrixAuthAction) error {
	payload := map[string]any{"authenticationcitrixauthaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationCitrixAuthActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationCitrixAuthAction(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", authenticationCitrixAuthActionURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UpdateAuthenticationCitrixAuthAction(resource models.AuthenticationCitrixAuthAction) error {
	payload := map[string]any{"authenticationcitrixauthaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, authenticationCitrixAuthActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UnsetAuthenticationCitrixAuthAction(resource models.AuthenticationCitrixAuthAction) error {
	payload := map[string]any{"authenticationcitrixauthaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", authenticationCitrixAuthActionURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationCitrixAuthAction() ([]models.AuthenticationCitrixAuthAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationCitrixAuthActionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationCitrixAuthAction `json:"authenticationcitrixauthaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationCitrixAuthAction(name string) (*models.AuthenticationCitrixAuthAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationCitrixAuthActionURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationCitrixAuthAction `json:"authenticationcitrixauthaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationcitrixauthaction %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationCitrixAuthAction() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationCitrixAuthActionURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationCitrixAuthAction `json:"authenticationcitrixauthaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationdfaaction
// Configuration for Dfa authentication action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationdfaaction
func (s *AuthenticationService) AddAuthenticationDFAAction(resource models.AuthenticationDFAAction) error {
	payload := map[string]any{"authenticationdfaaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationDFAActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationDFAAction(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", authenticationDFAActionURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UpdateAuthenticationDFAAction(resource models.AuthenticationDFAAction) error {
	payload := map[string]any{"authenticationdfaaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, authenticationDFAActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UnsetAuthenticationDFAAction(resource models.AuthenticationDFAAction) error {
	payload := map[string]any{"authenticationdfaaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", authenticationDFAActionURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationDFAAction() ([]models.AuthenticationDFAAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationDFAActionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationDFAAction `json:"authenticationdfaaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationDFAAction(name string) (*models.AuthenticationDFAAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationDFAActionURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationDFAAction `json:"authenticationdfaaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationdfaaction %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationDFAAction() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationDFAActionURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationDFAAction `json:"authenticationdfaaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationdfapolicy
// Configuration for Dfa authentication policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationdfapolicy
func (s *AuthenticationService) AddAuthenticationDFAPolicy(resource models.AuthenticationDFAPolicy) error {
	payload := map[string]any{"authenticationdfapolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationDFAPolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationDFAPolicy(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", authenticationDFAPolicyURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UpdateAuthenticationDFAPolicy(resource models.AuthenticationDFAPolicy) error {
	payload := map[string]any{"authenticationdfapolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, authenticationDFAPolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationDFAPolicy() ([]models.AuthenticationDFAPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationDFAPolicyURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationDFAPolicy `json:"authenticationdfapolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationDFAPolicy(name string) (*models.AuthenticationDFAPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationDFAPolicyURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationDFAPolicy `json:"authenticationdfapolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationdfapolicy %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationDFAPolicy() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationDFAPolicyURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationDFAPolicy `json:"authenticationdfapolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationdfapolicy_binding
// Binding object which returns the resources bound to authenticationdfapolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationdfapolicy_binding
func (s *AuthenticationService) GetAllAuthenticationDFAPolicyBinding() ([]models.AuthenticationDFAPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationDFAPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationDFAPolicyBinding `json:"authenticationdfapolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationDFAPolicyBinding(name string) (*models.AuthenticationDFAPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationDFAPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationDFAPolicyBinding `json:"authenticationdfapolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationdfapolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

// authenticationdfapolicy_vpnvserver_binding
// Binding object showing the vpnvserver that can be bound to authenticationdfapolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationdfapolicy_vpnvserver_binding
func (s *AuthenticationService) GetAllAuthenticationDFAPolicyVPNVServerBinding() ([]models.AuthenticationDFAPolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationDFAPolicyVPNVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationDFAPolicyVPNVServerBinding `json:"authenticationdfapolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationDFAPolicyVPNVServerBinding(name string) (*models.AuthenticationDFAPolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationDFAPolicyVPNVServerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationDFAPolicyVPNVServerBinding `json:"authenticationdfapolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationdfapolicy_vpnvserver_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationDFAPolicyVPNVServerBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationDFAPolicyVPNVServerBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationDFAPolicyVPNVServerBinding `json:"authenticationdfapolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationemailaction
// Configuration for Email entity resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationemailaction
func (s *AuthenticationService) AddAuthenticationEmailAction(resource models.AuthenticationEmailAction) error {
	payload := map[string]any{"authenticationemailaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationEmailActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationEmailAction(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", authenticationEmailActionURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UpdateAuthenticationEmailAction(resource models.AuthenticationEmailAction) error {
	payload := map[string]any{"authenticationemailaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, authenticationEmailActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UnsetAuthenticationEmailAction(resource models.AuthenticationEmailAction) error {
	payload := map[string]any{"authenticationemailaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", authenticationEmailActionURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationEmailAction() ([]models.AuthenticationEmailAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationEmailActionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationEmailAction `json:"authenticationemailaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationEmailAction(name string) (*models.AuthenticationEmailAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationEmailActionURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationEmailAction `json:"authenticationemailaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationemailaction %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationEmailAction() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationEmailActionURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationEmailAction `json:"authenticationemailaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationepaaction
// Configuration for epa action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationepaaction
func (s *AuthenticationService) AddAuthenticationEPAAction(resource models.AuthenticationEPAAction) error {
	payload := map[string]any{"authenticationepaaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationEPAActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationEPAAction(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", authenticationEPAActionURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UpdateAuthenticationEPAAction(resource models.AuthenticationEPAAction) error {
	payload := map[string]any{"authenticationepaaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, authenticationEPAActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UnsetAuthenticationEPAAction(resource models.AuthenticationEPAAction) error {
	payload := map[string]any{"authenticationepaaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", authenticationEPAActionURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationEPAAction() ([]models.AuthenticationEPAAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationEPAActionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationEPAAction `json:"authenticationepaaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationEPAAction(name string) (*models.AuthenticationEPAAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationEPAActionURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationEPAAction `json:"authenticationepaaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationepaaction %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationEPAAction() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationEPAActionURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationEPAAction `json:"authenticationepaaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationldapaction
// Configuration for LDAP action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationldapaction
func (s *AuthenticationService) AddAuthenticationLDAPAction(resource models.AuthenticationLDAPAction) error {
	payload := map[string]any{"authenticationldapaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationLDAPActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationLDAPAction(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", authenticationLDAPActionURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UpdateAuthenticationLDAPAction(resource models.AuthenticationLDAPAction) error {
	payload := map[string]any{"authenticationldapaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, authenticationLDAPActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UnsetAuthenticationLDAPAction(resource models.AuthenticationLDAPAction) error {
	payload := map[string]any{"authenticationldapaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", authenticationLDAPActionURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationLDAPAction() ([]models.AuthenticationLDAPAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationLDAPActionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationLDAPAction `json:"authenticationldapaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationLDAPAction(name string) (*models.AuthenticationLDAPAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationLDAPActionURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationLDAPAction `json:"authenticationldapaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationldapaction %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationLDAPAction() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationLDAPActionURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationLDAPAction `json:"authenticationldapaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationldappolicy
// Configuration for LDAP policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationldappolicy
func (s *AuthenticationService) AddAuthenticationLDAPPolicy(resource models.AuthenticationLDAPPolicy) error {
	payload := map[string]any{"authenticationldappolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationLDAPPolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationLDAPPolicy(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", authenticationLDAPPolicyURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UpdateAuthenticationLDAPPolicy(resource models.AuthenticationLDAPPolicy) error {
	payload := map[string]any{"authenticationldappolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, authenticationLDAPPolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UnsetAuthenticationLDAPPolicy(resource models.AuthenticationLDAPPolicy) error {
	payload := map[string]any{"authenticationldappolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", authenticationLDAPPolicyURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationLDAPPolicy() ([]models.AuthenticationLDAPPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationLDAPPolicyURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationLDAPPolicy `json:"authenticationldappolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationLDAPPolicy(name string) (*models.AuthenticationLDAPPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationLDAPPolicyURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationLDAPPolicy `json:"authenticationldappolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationldappolicy %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationLDAPPolicy() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationLDAPPolicyURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationLDAPPolicy `json:"authenticationldappolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationldappolicy_authenticationvserver_binding
// Binding object showing the authenticationvserver that can be bound to authenticationldappolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationldappolicy_authenticationvserver_binding
func (s *AuthenticationService) GetAllAuthenticationLDAPPolicyAuthenticationVServerBinding() ([]models.AuthenticationLDAPPolicyAuthenticationVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationLDAPPolicyAuthenticationVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationLDAPPolicyAuthenticationVServerBinding `json:"authenticationldappolicy_authenticationvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationLDAPPolicyAuthenticationVServerBinding(name string) (*models.AuthenticationLDAPPolicyAuthenticationVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationLDAPPolicyAuthenticationVServerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationLDAPPolicyAuthenticationVServerBinding `json:"authenticationldappolicy_authenticationvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationldappolicy_authenticationvserver_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationLDAPPolicyAuthenticationVServerBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationLDAPPolicyAuthenticationVServerBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationLDAPPolicyAuthenticationVServerBinding `json:"authenticationldappolicy_authenticationvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationldappolicy_binding
// Binding object which returns the resources bound to authenticationldappolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationldappolicy_binding
func (s *AuthenticationService) GetAllAuthenticationLDAPPolicyBinding() ([]models.AuthenticationLDAPPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationLDAPPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationLDAPPolicyBinding `json:"authenticationldappolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationLDAPPolicyBinding(name string) (*models.AuthenticationLDAPPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationLDAPPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationLDAPPolicyBinding `json:"authenticationldappolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationldappolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

// authenticationldappolicy_systemglobal_binding
// Binding object showing the systemglobal that can be bound to authenticationldappolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationldappolicy_systemglobal_binding
func (s *AuthenticationService) GetAllAuthenticationLDAPPolicySystemGlobalBinding() ([]models.AuthenticationLDAPPolicySystemGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationLDAPPolicySystemGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationLDAPPolicySystemGlobalBinding `json:"authenticationldappolicy_systemglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationLDAPPolicySystemGlobalBinding(name string) (*models.AuthenticationLDAPPolicySystemGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationLDAPPolicySystemGlobalBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationLDAPPolicySystemGlobalBinding `json:"authenticationldappolicy_systemglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationldappolicy_systemglobal_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationLDAPPolicySystemGlobalBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationLDAPPolicySystemGlobalBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationLDAPPolicySystemGlobalBinding `json:"authenticationldappolicy_systemglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationldappolicy_vpnglobal_binding
// Binding object showing the vpnglobal that can be bound to authenticationldappolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationldappolicy_vpnglobal_binding
func (s *AuthenticationService) GetAllAuthenticationLDAPPolicyVPNGlobalBinding() ([]models.AuthenticationLDAPPolicyVPNGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationLDAPPolicyVPNGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationLDAPPolicyVPNGlobalBinding `json:"authenticationldappolicy_vpnglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationLDAPPolicyVPNGlobalBinding(name string) (*models.AuthenticationLDAPPolicyVPNGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationLDAPPolicyVPNGlobalBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationLDAPPolicyVPNGlobalBinding `json:"authenticationldappolicy_vpnglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationldappolicy_vpnglobal_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationLDAPPolicyVPNGlobalBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationLDAPPolicyVPNGlobalBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationLDAPPolicyVPNGlobalBinding `json:"authenticationldappolicy_vpnglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationldappolicy_vpnvserver_binding
// Binding object showing the vpnvserver that can be bound to authenticationldappolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationldappolicy_vpnvserver_binding
func (s *AuthenticationService) GetAllAuthenticationLDAPPolicyVPNVServerBinding() ([]models.AuthenticationLDAPPolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationLDAPPolicyVPNVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationLDAPPolicyVPNVServerBinding `json:"authenticationldappolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationLDAPPolicyVPNVServerBinding(name string) (*models.AuthenticationLDAPPolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationLDAPPolicyVPNVServerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationLDAPPolicyVPNVServerBinding `json:"authenticationldappolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationldappolicy_vpnvserver_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationLDAPPolicyVPNVServerBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationLDAPPolicyVPNVServerBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationLDAPPolicyVPNVServerBinding `json:"authenticationldappolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationlocalpolicy
// Configuration for LOCAL policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationlocalpolicy
func (s *AuthenticationService) AddAuthenticationLocalPolicy(resource models.AuthenticationLocalPolicy) error {
	payload := map[string]any{"authenticationlocalpolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationLocalPolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationLocalPolicy(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", authenticationLocalPolicyURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UpdateAuthenticationLocalPolicy(resource models.AuthenticationLocalPolicy) error {
	payload := map[string]any{"authenticationlocalpolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, authenticationLocalPolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationLocalPolicy() ([]models.AuthenticationLocalPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationLocalPolicyURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationLocalPolicy `json:"authenticationlocalpolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationLocalPolicy(name string) (*models.AuthenticationLocalPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationLocalPolicyURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationLocalPolicy `json:"authenticationlocalpolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationlocalpolicy %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationLocalPolicy() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationLocalPolicyURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationLocalPolicy `json:"authenticationlocalpolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationlocalpolicy_authenticationvserver_binding
// Binding object showing the authenticationvserver that can be bound to authenticationlocalpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationlocalpolicy_authenticationvserver_binding
func (s *AuthenticationService) GetAllAuthenticationLocalPolicyAuthenticationVServerBinding() ([]models.AuthenticationLocalPolicyAuthenticationVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationLocalPolicyAuthenticationVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationLocalPolicyAuthenticationVServerBinding `json:"authenticationlocalpolicy_authenticationvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationLocalPolicyAuthenticationVServerBinding(name string) (*models.AuthenticationLocalPolicyAuthenticationVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationLocalPolicyAuthenticationVServerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationLocalPolicyAuthenticationVServerBinding `json:"authenticationlocalpolicy_authenticationvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationlocalpolicy_authenticationvserver_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationLocalPolicyAuthenticationVServerBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationLocalPolicyAuthenticationVServerBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationLocalPolicyAuthenticationVServerBinding `json:"authenticationlocalpolicy_authenticationvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationlocalpolicy_binding
// Binding object which returns the resources bound to authenticationlocalpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationlocalpolicy_binding
func (s *AuthenticationService) GetAllAuthenticationLocalPolicyBinding() ([]models.AuthenticationLocalPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationLocalPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationLocalPolicyBinding `json:"authenticationlocalpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationLocalPolicyBinding(name string) (*models.AuthenticationLocalPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationLocalPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationLocalPolicyBinding `json:"authenticationlocalpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationlocalpolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

// authenticationlocalpolicy_systemglobal_binding
// Binding object showing the systemglobal that can be bound to authenticationlocalpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationlocalpolicy_systemglobal_binding
func (s *AuthenticationService) GetAllAuthenticationLocalPolicySystemGlobalBinding() ([]models.AuthenticationLocalPolicySystemGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationLocalPolicySystemGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationLocalPolicySystemGlobalBinding `json:"authenticationlocalpolicy_systemglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationLocalPolicySystemGlobalBinding(name string) (*models.AuthenticationLocalPolicySystemGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationLocalPolicySystemGlobalBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationLocalPolicySystemGlobalBinding `json:"authenticationlocalpolicy_systemglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationlocalpolicy_systemglobal_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationLocalPolicySystemGlobalBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationLocalPolicySystemGlobalBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationLocalPolicySystemGlobalBinding `json:"authenticationlocalpolicy_systemglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationlocalpolicy_vpnglobal_binding
// Binding object showing the vpnglobal that can be bound to authenticationlocalpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationlocalpolicy_vpnglobal_binding
func (s *AuthenticationService) GetAllAuthenticationLocalPolicyVPNGlobalBinding() ([]models.AuthenticationLocalPolicyVPNGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationLocalPolicyVPNGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationLocalPolicyVPNGlobalBinding `json:"authenticationlocalpolicy_vpnglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationLocalPolicyVPNGlobalBinding(name string) (*models.AuthenticationLocalPolicyVPNGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationLocalPolicyVPNGlobalBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationLocalPolicyVPNGlobalBinding `json:"authenticationlocalpolicy_vpnglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationlocalpolicy_vpnglobal_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationLocalPolicyVPNGlobalBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationLocalPolicyVPNGlobalBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationLocalPolicyVPNGlobalBinding `json:"authenticationlocalpolicy_vpnglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationlocalpolicy_vpnvserver_binding
// Binding object showing the vpnvserver that can be bound to authenticationlocalpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationlocalpolicy_vpnvserver_binding
func (s *AuthenticationService) GetAllAuthenticationLocalPolicyVPNVServerBinding() ([]models.AuthenticationLocalPolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationLocalPolicyVPNVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationLocalPolicyVPNVServerBinding `json:"authenticationlocalpolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationLocalPolicyVPNVServerBinding(name string) (*models.AuthenticationLocalPolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationLocalPolicyVPNVServerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationLocalPolicyVPNVServerBinding `json:"authenticationlocalpolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationlocalpolicy_vpnvserver_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationLocalPolicyVPNVServerBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationLocalPolicyVPNVServerBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationLocalPolicyVPNVServerBinding `json:"authenticationlocalpolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationloginschema
// Configuration for Login Schema resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationloginschema
func (s *AuthenticationService) AddAuthenticationLoginSchema(resource models.AuthenticationLoginSchema) error {
	payload := map[string]any{"authenticationloginschema": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationLoginSchemaURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationLoginSchema(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", authenticationLoginSchemaURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UpdateAuthenticationLoginSchema(resource models.AuthenticationLoginSchema) error {
	payload := map[string]any{"authenticationloginschema": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, authenticationLoginSchemaURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UnsetAuthenticationLoginSchema(resource models.AuthenticationLoginSchema) error {
	payload := map[string]any{"authenticationloginschema": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", authenticationLoginSchemaURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationLoginSchema() ([]models.AuthenticationLoginSchema, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationLoginSchemaURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationLoginSchema `json:"authenticationloginschema"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationLoginSchema(name string) (*models.AuthenticationLoginSchema, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationLoginSchemaURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationLoginSchema `json:"authenticationloginschema"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationloginschema %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationLoginSchema() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationLoginSchemaURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationLoginSchema `json:"authenticationloginschema"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationloginschemapolicy
// Configuration for Login Schema Policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationloginschemapolicy
func (s *AuthenticationService) AddAuthenticationLoginSchemaPolicy(resource models.AuthenticationLoginSchemaPolicy) error {
	payload := map[string]any{"authenticationloginschemapolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationLoginSchemaPolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationLoginSchemaPolicy(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", authenticationLoginSchemaPolicyURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UpdateAuthenticationLoginSchemaPolicy(resource models.AuthenticationLoginSchemaPolicy) error {
	payload := map[string]any{"authenticationloginschemapolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, authenticationLoginSchemaPolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UnsetAuthenticationLoginSchemaPolicy(resource models.AuthenticationLoginSchemaPolicy) error {
	payload := map[string]any{"authenticationloginschemapolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", authenticationLoginSchemaPolicyURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationLoginSchemaPolicy() ([]models.AuthenticationLoginSchemaPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationLoginSchemaPolicyURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationLoginSchemaPolicy `json:"authenticationloginschemapolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationLoginSchemaPolicy(name string) (*models.AuthenticationLoginSchemaPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationLoginSchemaPolicyURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationLoginSchemaPolicy `json:"authenticationloginschemapolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationloginschemapolicy %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationLoginSchemaPolicy() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationLoginSchemaPolicyURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationLoginSchemaPolicy `json:"authenticationloginschemapolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

func (s *AuthenticationService) RenameAuthenticationLoginSchemaPolicy(name string, newname string) error {
	payload := map[string]any{
		"authenticationloginschemapolicy": map[string]any{
			"name":    name,
			"newname": newname,
		},
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=rename", authenticationLoginSchemaPolicyURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// authenticationloginschemapolicy_authenticationvserver_binding
// Binding object showing the authenticationvserver that can be bound to authenticationloginschemapolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationloginschemapolicy_authenticationvserver_binding
func (s *AuthenticationService) GetAllAuthenticationLoginSchemaPolicyAuthenticationVServerBinding() ([]models.AuthenticationLoginSchemaPolicyAuthenticationVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationLoginSchemaPolicyAuthenticationVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationLoginSchemaPolicyAuthenticationVServerBinding `json:"authenticationloginschemapolicy_authenticationvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationLoginSchemaPolicyAuthenticationVServerBinding(name string) (*models.AuthenticationLoginSchemaPolicyAuthenticationVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationLoginSchemaPolicyAuthenticationVServerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationLoginSchemaPolicyAuthenticationVServerBinding `json:"authenticationloginschemapolicy_authenticationvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationloginschemapolicy_authenticationvserver_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationLoginSchemaPolicyAuthenticationVServerBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationLoginSchemaPolicyAuthenticationVServerBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationLoginSchemaPolicyAuthenticationVServerBinding `json:"authenticationloginschemapolicy_authenticationvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationloginschemapolicy_binding
// Binding object which returns the resources bound to authenticationloginschemapolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationloginschemapolicy_binding
func (s *AuthenticationService) GetAllAuthenticationLoginSchemaPolicyBinding() ([]models.AuthenticationLoginSchemaPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationLoginSchemaPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationLoginSchemaPolicyBinding `json:"authenticationloginschemapolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationLoginSchemaPolicyBinding(name string) (*models.AuthenticationLoginSchemaPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationLoginSchemaPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationLoginSchemaPolicyBinding `json:"authenticationloginschemapolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationloginschemapolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

// authenticationloginschemapolicy_vpnvserver_binding
// Binding object showing the vpnvserver that can be bound to authenticationloginschemapolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationloginschemapolicy_vpnvserver_binding
func (s *AuthenticationService) GetAllAuthenticationLoginSchemaPolicyVPNVServerBinding() ([]models.AuthenticationLoginSchemaPolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationLoginSchemaPolicyVPNVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationLoginSchemaPolicyVPNVServerBinding `json:"authenticationloginschemapolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationLoginSchemaPolicyVPNVServerBinding(name string) (*models.AuthenticationLoginSchemaPolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationLoginSchemaPolicyVPNVServerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationLoginSchemaPolicyVPNVServerBinding `json:"authenticationloginschemapolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationloginschemapolicy_vpnvserver_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationLoginSchemaPolicyVPNVServerBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationLoginSchemaPolicyVPNVServerBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationLoginSchemaPolicyVPNVServerBinding `json:"authenticationloginschemapolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationnegotiataction
// Configuration for Negotiate action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationnegotiateaction
func (s *AuthenticationService) AddAuthenticationNegotiateAction(resource models.AuthenticationNegotiateAction) error {
	payload := map[string]any{"authenticationnegotiateaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationNegotiateActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationNegotiateAction(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", authenticationNegotiateActionURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UpdateAuthenticationNegotiateAction(resource models.AuthenticationNegotiateAction) error {
	payload := map[string]any{"authenticationnegotiateaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, authenticationNegotiateActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UnsetAuthenticationNegotiateAction(resource models.AuthenticationNegotiateAction) error {
	payload := map[string]any{"authenticationnegotiateaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", authenticationNegotiateActionURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationNegotiateAction() ([]models.AuthenticationNegotiateAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationNegotiateActionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationNegotiateAction `json:"authenticationnegotiateaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationNegotiateAction(name string) (*models.AuthenticationNegotiateAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationNegotiateActionURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationNegotiateAction `json:"authenticationnegotiateaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationnegotiateaction %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationNegotiateAction() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationNegotiateActionURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationNegotiateAction `json:"authenticationnegotiateaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationnegotiatpolicy
// Configuration for Negotiate Policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationnegotiatepolicy
func (s *AuthenticationService) AddAuthenticationNegotiatePolicy(resource models.AuthenticationNegotiatePolicy) error {
	payload := map[string]any{"authenticationnegotiatepolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationNegotiatePolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationNegotiatePolicy(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", authenticationNegotiatePolicyURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UpdateAuthenticationNegotiatePolicy(resource models.AuthenticationNegotiatePolicy) error {
	payload := map[string]any{"authenticationnegotiatepolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, authenticationNegotiatePolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UnsetAuthenticationNegotiatePolicy(resource models.AuthenticationNegotiatePolicy) error {
	payload := map[string]any{"authenticationnegotiatepolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", authenticationNegotiatePolicyURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationNegotiatePolicy() ([]models.AuthenticationNegotiatePolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationNegotiatePolicyURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationNegotiatePolicy `json:"authenticationnegotiatepolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationNegotiatePolicy(name string) (*models.AuthenticationNegotiatePolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationNegotiatePolicyURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationNegotiatePolicy `json:"authenticationnegotiatepolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationnegotiatepolicy %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationNegotiatePolicy() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationNegotiatePolicyURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationNegotiatePolicy `json:"authenticationnegotiatepolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationnegotiatepolicy_authenticationvserver_binding
// Binding object showing the authenticationvserver that can be bound to authenticationnegotiatepolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationnegotiatepolicy_authenticationvserver_binding
func (s *AuthenticationService) GetAllAuthenticationNegotiatePolicyAuthenticationVServerBinding() ([]models.AuthenticationNegotiatePolicyAuthenticationVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationNegotiatePolicyAuthenticationVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationNegotiatePolicyAuthenticationVServerBinding `json:"authenticationnegotiatepolicy_authenticationvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationNegotiatePolicyAuthenticationVServerBinding(name string) (*models.AuthenticationNegotiatePolicyAuthenticationVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationNegotiatePolicyAuthenticationVServerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationNegotiatePolicyAuthenticationVServerBinding `json:"authenticationnegotiatepolicy_authenticationvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationnegotiatepolicy_authenticationvserver_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationNegotiatePolicyAuthenticationVServerBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationNegotiatePolicyAuthenticationVServerBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationNegotiatePolicyAuthenticationVServerBinding `json:"authenticationnegotiatepolicy_authenticationvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationnegotiatepolicy_binding
// Binding object which returns the resources bound to authenticationnegotiatepolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationnegotiatepolicy_binding
func (s *AuthenticationService) GetAllAuthenticationNegotiatePolicyBinding() ([]models.AuthenticationNegotiatePolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationNegotiatePolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationNegotiatePolicyBinding `json:"authenticationnegotiatepolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationNegotiatePolicyBinding(name string) (*models.AuthenticationNegotiatePolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationNegotiatePolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationNegotiatePolicyBinding `json:"authenticationnegotiatepolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationnegotiatepolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

// authenticationnegotiatepolicy_vpnglobal_binding
// Binding object showing the vpnglobal that can be bound to authenticationnegotiatepolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationnegotiatepolicy_vpnglobal_binding
func (s *AuthenticationService) GetAllAuthenticationNegotiatePolicyVPNGlobalBinding() ([]models.AuthenticationNegotiatePolicyVPNGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationNegotiatePolicyVPNGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationNegotiatePolicyVPNGlobalBinding `json:"authenticationnegotiatepolicy_vpnglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationNegotiatePolicyVPNGlobalBinding(name string) (*models.AuthenticationNegotiatePolicyVPNGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationNegotiatePolicyVPNGlobalBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationNegotiatePolicyVPNGlobalBinding `json:"authenticationnegotiatepolicy_vpnglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationnegotiatepolicy_vpnglobal_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationNegotiatePolicyVPNGlobalBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationNegotiatePolicyVPNGlobalBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationNegotiatePolicyVPNGlobalBinding `json:"authenticationnegotiatepolicy_vpnglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationnegotiatepolicy_vpnvserver_binding
// Binding object showing the vpnvserver that can be bound to authenticationnegotiatepolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationnegotiatepolicy_vpnvserver_binding
func (s *AuthenticationService) GetAllAuthenticationNegotiatePolicyVPNVServerBinding() ([]models.AuthenticationNegotiatePolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationNegotiatePolicyVPNVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationNegotiatePolicyVPNVServerBinding `json:"authenticationnegotiatepolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationNegotiatePolicyVPNVServerBinding(name string) (*models.AuthenticationNegotiatePolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationNegotiatePolicyVPNVServerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationNegotiatePolicyVPNVServerBinding `json:"authenticationnegotiatepolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationnegotiatepolicy_vpnvserver_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationNegotiatePolicyVPNVServerBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationNegotiatePolicyVPNVServerBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationNegotiatePolicyVPNVServerBinding `json:"authenticationnegotiatepolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationnoauthaction
// Configuration for no authentication action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationnoauthaction
func (s *AuthenticationService) AddAuthenticationNoAuthAction(resource models.AuthenticationNoAuthAction) error {
	payload := map[string]any{"authenticationnoauthaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationNoAuthActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationNoAuthAction(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", authenticationNoAuthActionURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UpdateAuthenticationNoAuthAction(resource models.AuthenticationNoAuthAction) error {
	payload := map[string]any{"authenticationnoauthaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, authenticationNoAuthActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UnsetAuthenticationNoAuthAction(resource models.AuthenticationNoAuthAction) error {
	payload := map[string]any{"authenticationnoauthaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", authenticationNoAuthActionURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationNoAuthAction() ([]models.AuthenticationNoAuthAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationNoAuthActionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationNoAuthAction `json:"authenticationnoauthaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationNoAuthAction(name string) (*models.AuthenticationNoAuthAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationNoAuthActionURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationNoAuthAction `json:"authenticationnoauthaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationnoauthaction %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationNoAuthAction() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationNoAuthActionURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationNoAuthAction `json:"authenticationnoauthaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationoauthaction
// Configuration for OAuth authentication action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationoauthaction
func (s *AuthenticationService) AddAuthenticationOAuthAction(resource models.AuthenticationOAuthAction) error {
	payload := map[string]any{"authenticationoauthaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationOAuthActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationOAuthAction(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", authenticationOAuthActionURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UpdateAuthenticationOAuthAction(resource models.AuthenticationOAuthAction) error {
	payload := map[string]any{"authenticationoauthaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, authenticationOAuthActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UnsetAuthenticationOAuthAction(resource models.AuthenticationOAuthAction) error {
	payload := map[string]any{"authenticationoauthaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", authenticationOAuthActionURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationOAuthAction() ([]models.AuthenticationOAuthAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationOAuthActionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationOAuthAction `json:"authenticationoauthaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationOAuthAction(name string) (*models.AuthenticationOAuthAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationOAuthActionURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationOAuthAction `json:"authenticationoauthaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationoauthaction %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationOAuthAction() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationOAuthActionURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationOAuthAction `json:"authenticationoauthaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationoauthidppolicy
// Configuration for AAA OAuth IdentityProvider (IdP) policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationoauthidppolicy
func (s *AuthenticationService) AddAuthenticationOAuthIdPPolicy(resource models.AuthenticationOAuthIDPPolicy) error {
	payload := map[string]any{"authenticationoauthidppolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationOAuthIdPPolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationOAuthIdPPolicy(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", authenticationOAuthIdPPolicyURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UpdateAuthenticationOAuthIdPPolicy(resource models.AuthenticationOAuthIDPPolicy) error {
	payload := map[string]any{"authenticationoauthidppolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, authenticationOAuthIdPPolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UnsetAuthenticationOAuthIdPPolicy(resource models.AuthenticationOAuthIDPPolicy) error {
	payload := map[string]any{"authenticationoauthidppolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", authenticationOAuthIdPPolicyURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationOAuthIdPPolicy() ([]models.AuthenticationOAuthIDPPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationOAuthIdPPolicyURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationOAuthIDPPolicy `json:"authenticationoauthidppolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationOAuthIdPPolicy(name string) (*models.AuthenticationOAuthIDPPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationOAuthIdPPolicyURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationOAuthIDPPolicy `json:"authenticationoauthidppolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationoauthidppolicy %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationOAuthIdPPolicy() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationOAuthIdPPolicyURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationOAuthIDPPolicy `json:"authenticationoauthidppolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

func (s *AuthenticationService) RenameAuthenticationOAuthIdPPolicy(name string, newname string) error {
	payload := map[string]any{
		"authenticationoauthidppolicy": map[string]any{
			"name":    name,
			"newname": newname,
		},
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=rename", authenticationOAuthIdPPolicyURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// authenticationoauthidppolicy_authenticationvserver_binding
// Binding object showing the authenticationvserver that can be bound to authenticationoauthidppolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationoauthidppolicy_authenticationvserver_binding
func (s *AuthenticationService) GetAllAuthenticationOAuthIdPPolicyAuthenticationVServerBinding() ([]models.AuthenticationOAuthIDPPolicyAuthenticationVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationOAuthIdPPolicyAuthenticationVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationOAuthIDPPolicyAuthenticationVServerBinding `json:"authenticationoauthidppolicy_authenticationvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationOAuthIdPPolicyAuthenticationVServerBinding(name string) (*models.AuthenticationOAuthIDPPolicyAuthenticationVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationOAuthIdPPolicyAuthenticationVServerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationOAuthIDPPolicyAuthenticationVServerBinding `json:"authenticationoauthidppolicy_authenticationvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationoauthidppolicy_authenticationvserver_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationOAuthIdPPolicyAuthenticationVServerBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationOAuthIdPPolicyAuthenticationVServerBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationOAuthIDPPolicyAuthenticationVServerBinding `json:"authenticationoauthidppolicy_authenticationvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationoauthidppolicy_binding
// Binding object which returns the resources bound to authenticationoauthidppolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationoauthidppolicy_binding
func (s *AuthenticationService) GetAllAuthenticationOAuthIdPPolicyBinding() ([]models.AuthenticationOAuthIDPPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationOAuthIdPPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationOAuthIDPPolicyBinding `json:"authenticationoauthidppolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationOAuthIdPPolicyBinding(name string) (*models.AuthenticationOAuthIDPPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationOAuthIdPPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationOAuthIDPPolicyBinding `json:"authenticationoauthidppolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationoauthidppolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

// authenticationoauthidppolicy_vpnvserver_binding
// Binding object showing the vpnvserver that can be bound to authenticationoauthidppolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationoauthidppolicy_vpnvserver_binding
func (s *AuthenticationService) GetAllAuthenticationOAuthIdPPolicyVPNVServerBinding() ([]models.AuthenticationOAuthIDPPolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationOAuthIdPPolicyVPNVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationOAuthIDPPolicyVPNVServerBinding `json:"authenticationoauthidppolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationOAuthIdPPolicyVPNVServerBinding(name string) (*models.AuthenticationOAuthIDPPolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationOAuthIdPPolicyVPNVServerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationOAuthIDPPolicyVPNVServerBinding `json:"authenticationoauthidppolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationoauthidppolicy_vpnvserver_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationOAuthIdPPolicyVPNVServerBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationOAuthIdPPolicyVPNVServerBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationOAuthIDPPolicyVPNVServerBinding `json:"authenticationoauthidppolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationoauthidpprofile
// Configuration for OAuth Identity Provider (IdP) profile resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationoauthidpprofile
func (s *AuthenticationService) AddAuthenticationOAuthIdPProfile(resource models.AuthenticationOAuthIDPProfile) error {
	payload := map[string]any{"authenticationoauthidpprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationOAuthIdPProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationOAuthIdPProfile(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", authenticationOAuthIdPProfileURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UpdateAuthenticationOAuthIdPProfile(resource models.AuthenticationOAuthIDPProfile) error {
	payload := map[string]any{"authenticationoauthidpprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, authenticationOAuthIdPProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UnsetAuthenticationOAuthIdPProfile(resource models.AuthenticationOAuthIDPProfile) error {
	payload := map[string]any{"authenticationoauthidpprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", authenticationOAuthIdPProfileURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationOAuthIdPProfile() ([]models.AuthenticationOAuthIDPProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationOAuthIdPProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationOAuthIDPProfile `json:"authenticationoauthidpprofile"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationOAuthIdPProfile(name string) (*models.AuthenticationOAuthIDPProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationOAuthIdPProfileURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationOAuthIDPProfile `json:"authenticationoauthidpprofile"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationoauthidpprofile %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationOAuthIdPProfile() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationOAuthIdPProfileURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationOAuthIDPProfile `json:"authenticationoauthidpprofile"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationpolicy
// Configuration for Authentication Policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationpolicy
func (s *AuthenticationService) AddAuthenticationPolicy(resource models.AuthenticationPolicy) error {
	payload := map[string]any{"authenticationpolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationPolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationPolicy(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", authenticationPolicyURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UpdateAuthenticationPolicy(resource models.AuthenticationPolicy) error {
	payload := map[string]any{"authenticationpolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, authenticationPolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UnsetAuthenticationPolicy(resource models.AuthenticationPolicy) error {
	payload := map[string]any{"authenticationpolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", authenticationPolicyURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationPolicy() ([]models.AuthenticationPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationPolicyURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationPolicy `json:"authenticationpolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationPolicy(name string) (*models.AuthenticationPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationPolicyURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationPolicy `json:"authenticationpolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationpolicy %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationPolicy() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationPolicyURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationPolicy `json:"authenticationpolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

func (s *AuthenticationService) RenameAuthenticationPolicy(name string, newname string) error {
	payload := map[string]any{
		"authenticationpolicy": map[string]any{
			"name":    name,
			"newname": newname,
		},
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=rename", authenticationPolicyURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// authenticationpolicylabel
// Configuration for authentication policy label resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationpolicylabel
func (s *AuthenticationService) AddAuthenticationPolicyLabel(resource models.AuthenticationPolicyLabel) error {
	payload := map[string]any{"authenticationpolicylabel": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationPolicyLabelURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationPolicyLabel(labelname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", authenticationPolicyLabelURL, labelname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationPolicyLabel() ([]models.AuthenticationPolicyLabel, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationPolicyLabelURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationPolicyLabel `json:"authenticationpolicylabel"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationPolicyLabel(labelname string) (*models.AuthenticationPolicyLabel, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationPolicyLabelURL, labelname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationPolicyLabel `json:"authenticationpolicylabel"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationpolicylabel %s not found", labelname)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationPolicyLabel() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationPolicyLabelURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationPolicyLabel `json:"authenticationpolicylabel"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

func (s *AuthenticationService) RenameAuthenticationPolicyLabel(labelname string, newname string) error {
	payload := map[string]any{
		"authenticationpolicylabel": map[string]any{
			"labelname": labelname,
			"newname":   newname,
		},
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=rename", authenticationPolicyLabelURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// authenticationpolicylabel_authenticationpolicy_binding
// Binding object showing the authenticationpolicy that can be bound to authenticationpolicylabel.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationpolicylabel_authenticationpolicy_binding
func (s *AuthenticationService) AddAuthenticationPolicyLabelAuthenticationPolicyBinding(resource models.AuthenticationPolicyLabelAuthenticationPolicyBinding) error {
	payload := map[string]any{"authenticationpolicylabel_authenticationpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationPolicyLabelAuthenticationPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationPolicyLabelAuthenticationPolicyBinding(labelname string, policyname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policyname:%s", authenticationPolicyLabelAuthenticationPolicyBindingURL, labelname, policyname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationPolicyLabelAuthenticationPolicyBinding() ([]models.AuthenticationPolicyLabelAuthenticationPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationPolicyLabelAuthenticationPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationPolicyLabelAuthenticationPolicyBinding `json:"authenticationpolicylabel_authenticationpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationPolicyLabelAuthenticationPolicyBinding(labelname string) (*models.AuthenticationPolicyLabelAuthenticationPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationPolicyLabelAuthenticationPolicyBindingURL, labelname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationPolicyLabelAuthenticationPolicyBinding `json:"authenticationpolicylabel_authenticationpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationpolicylabel_authenticationpolicy_binding %s not found", labelname)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationPolicyLabelAuthenticationPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationPolicyLabelAuthenticationPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationPolicyLabelAuthenticationPolicyBinding `json:"authenticationpolicylabel_authenticationpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticaitonpolicylabel_binding
// Binding object which returns the resources bound to authenticationpolicylabel.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationpolicylabel_binding
func (s *AuthenticationService) GetAllAuthenticationPolicyLabelBinding() ([]models.AuthenticationPolicyLabelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationPolicyLabelBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationPolicyLabelBinding `json:"authenticationpolicylabel_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationPolicyLabelBinding(labelname string) (*models.AuthenticationPolicyLabelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationPolicyLabelBindingURL, labelname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationPolicyLabelBinding `json:"authenticationpolicylabel_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationpolicylabel_binding %s not found", labelname)
	}

	return &result.Data[0], nil
}

// authenticationpolicy_authenticationpolicylabel_binding
// Binding object showing the authenticationpolicylabel that can be bound to authenticationpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationpolicy_authenticationpolicylabel_binding
func (s *AuthenticationService) GetAllAuthenticationPolicyAuthenticationPolicyLabelBinding() ([]models.AuthenticationPolicyAuthenticationPolicyLabelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationPolicyAuthenticationPolicyLabelBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationPolicyAuthenticationPolicyLabelBinding `json:"authenticationpolicy_authenticationpolicylabel_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationPolicyAuthenticationPolicyLabelBinding(name string) (*models.AuthenticationPolicyAuthenticationPolicyLabelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationPolicyAuthenticationPolicyLabelBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationPolicyAuthenticationPolicyLabelBinding `json:"authenticationpolicy_authenticationpolicylabel_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationpolicy_authenticationpolicylabel_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationPolicyAuthenticationPolicyLabelBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationPolicyAuthenticationPolicyLabelBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationPolicyAuthenticationPolicyLabelBinding `json:"authenticationpolicy_authenticationpolicylabel_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationpolicy_authenticationvserver_binding
// Binding object showing the authenticationvserver that can be bound to authenticationpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationpolicy_authenticationvserver_binding
func (s *AuthenticationService) GetAllAuthenticationPolicyAuthenticationVServerBinding() ([]models.AuthenticationPolicyAuthenticationVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationPolicyAuthenticationVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationPolicyAuthenticationVServerBinding `json:"authenticationpolicy_authenticationvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationPolicyAuthenticationVServerBinding(name string) (*models.AuthenticationPolicyAuthenticationVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationPolicyAuthenticationVServerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationPolicyAuthenticationVServerBinding `json:"authenticationpolicy_authenticationvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationpolicy_authenticationvserver_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationPolicyAuthenticationVServerBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationPolicyAuthenticationVServerBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationPolicyAuthenticationVServerBinding `json:"authenticationpolicy_authenticationvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationpolicy_binding
// Binding object which returns the resources bound to authenticationpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationpolicy_binding
func (s *AuthenticationService) GetAllAuthenticationPolicyBinding() ([]models.AuthenticationPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationPolicyBinding `json:"authenticationpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationPolicyBinding(name string) (*models.AuthenticationPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationPolicyBinding `json:"authenticationpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationpolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

// authenicationpolicy_systemglobal_binding
// Binding object showing the systemglobal that can be bound to authenticationpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationpolicy_systemglobal_binding
func (s *AuthenticationService) GetAllAuthenticationPolicySystemGlobalBinding() ([]models.AuthenticationPolicySystemGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationPolicySystemGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationPolicySystemGlobalBinding `json:"authenticationpolicy_systemglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationPolicySystemGlobalBinding(name string) (*models.AuthenticationPolicySystemGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationPolicySystemGlobalBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationPolicySystemGlobalBinding `json:"authenticationpolicy_systemglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationpolicy_systemglobal_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationPolicySystemGlobalBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationPolicySystemGlobalBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationPolicySystemGlobalBinding `json:"authenticationpolicy_systemglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationpushservice
// Configuration for Service details for sending push notifications resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationpushservice
func (s *AuthenticationService) AddAuthenticationPushService(resource models.AuthenticationPushService) error {
	payload := map[string]any{"authenticationpushservice": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationPushServiceURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationPushService(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", authenticationPushServiceURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UpdateAuthenticationPushService(resource models.AuthenticationPushService) error {
	payload := map[string]any{"authenticationpushservice": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, authenticationPushServiceURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UnsetAuthenticationPushService(resource models.AuthenticationPushService) error {
	payload := map[string]any{"authenticationpushservice": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", authenticationPushServiceURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationPushService() ([]models.AuthenticationPushService, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationPushServiceURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationPushService `json:"authenticationpushservice"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationPushService(name string) (*models.AuthenticationPushService, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationPushServiceURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationPushService `json:"authenticationpushservice"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationpushservice %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationPushService() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationPushServiceURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationPushService `json:"authenticationpushservice"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationradiusaction
// Configuration for RADIUS action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationradiusaction
func (s *AuthenticationService) AddAuthenticationRADIUSAction(resource models.AuthenticationRADIUSAction) error {
	payload := map[string]any{"authenticationradiusaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationRADIUSActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationRADIUSAction(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", authenticationRADIUSActionURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UpdateAuthenticationRADIUSAction(resource models.AuthenticationRADIUSAction) error {
	payload := map[string]any{"authenticationradiusaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, authenticationRADIUSActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UnsetAuthenticationRADIUSAction(resource models.AuthenticationRADIUSAction) error {
	payload := map[string]any{"authenticationradiusaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", authenticationRADIUSActionURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationRADIUSAction() ([]models.AuthenticationRADIUSAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationRADIUSActionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationRADIUSAction `json:"authenticationradiusaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationRADIUSAction(name string) (*models.AuthenticationRADIUSAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationRADIUSActionURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationRADIUSAction `json:"authenticationradiusaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationradiusaction %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationRADIUSAction() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationRADIUSActionURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationRADIUSAction `json:"authenticationradiusaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationradiuspolicy
// Configuration for RADIUS policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationradiuspolicy
func (s *AuthenticationService) AddAuthenticationRADIUSPolicy(resource models.AuthenticationRADIUSPolicy) error {
	payload := map[string]any{"authenticationradiuspolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationRADIUSPolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationRADIUSPolicy(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", authenticationRADIUSPolicyURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UpdateAuthenticationRADIUSPolicy(resource models.AuthenticationRADIUSPolicy) error {
	payload := map[string]any{"authenticationradiuspolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, authenticationRADIUSPolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UnsetAuthenticationRADIUSPolicy(resource models.AuthenticationRADIUSPolicy) error {
	payload := map[string]any{"authenticationradiuspolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", authenticationRADIUSPolicyURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationRADIUSPolicy() ([]models.AuthenticationRADIUSPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationRADIUSPolicyURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationRADIUSPolicy `json:"authenticationradiuspolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationRADIUSPolicy(name string) (*models.AuthenticationRADIUSPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationRADIUSPolicyURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationRADIUSPolicy `json:"authenticationradiuspolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationradiuspolicy %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationRADIUSPolicy() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationRADIUSPolicyURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationRADIUSPolicy `json:"authenticationradiuspolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationradiuspolicy_authenticationvserver_binding
// Binding object showing the authenticationvserver that can be bound to authenticationradiuspolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationradiuspolicy_authenticationvserver_binding
func (s *AuthenticationService) GetAllAuthenticationRADIUSPolicyAuthenticationVServerBinding() ([]models.AuthenticationRADIUSPolicyAuthenticationVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationRADIUSPolicyAuthenticationVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationRADIUSPolicyAuthenticationVServerBinding `json:"authenticationradiuspolicy_authenticationvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationRADIUSPolicyAuthenticationVServerBinding(name string) (*models.AuthenticationRADIUSPolicyAuthenticationVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationRADIUSPolicyAuthenticationVServerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationRADIUSPolicyAuthenticationVServerBinding `json:"authenticationradiuspolicy_authenticationvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationradiuspolicy_authenticationvserver_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationRADIUSPolicyAuthenticationVServerBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationRADIUSPolicyAuthenticationVServerBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationRADIUSPolicyAuthenticationVServerBinding `json:"authenticationradiuspolicy_authenticationvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationradiuspolicy_binding
// Binding object which returns the resources bound to authenticationradiuspolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationradiuspolicy_binding
func (s *AuthenticationService) GetAllAuthenticationRADIUSPolicyBinding() ([]models.AuthenticationRADIUSPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationRADIUSPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationRADIUSPolicyBinding `json:"authenticationradiuspolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationRADIUSPolicyBinding(name string) (*models.AuthenticationRADIUSPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationRADIUSPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationRADIUSPolicyBinding `json:"authenticationradiuspolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationradiuspolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

// authenticationradiuspolicy_systemglobal_binding
// Binding object showing the systemglobal that can be bound to authenticationradiuspolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationradiuspolicy_systemglobal_binding
func (s *AuthenticationService) GetAllAuthenticationRADIUSPolicySystemGlobalBinding() ([]models.AuthenticationRADIUSPolicySystemGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationRADIUSPolicySystemGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationRADIUSPolicySystemGlobalBinding `json:"authenticationradiuspolicy_systemglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationRADIUSPolicySystemGlobalBinding(name string) (*models.AuthenticationRADIUSPolicySystemGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationRADIUSPolicySystemGlobalBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationRADIUSPolicySystemGlobalBinding `json:"authenticationradiuspolicy_systemglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationradiuspolicy_systemglobal_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationRADIUSPolicySystemGlobalBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationRADIUSPolicySystemGlobalBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationRADIUSPolicySystemGlobalBinding `json:"authenticationradiuspolicy_systemglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationradiuspolicy_vpnglobal_binding
// Binding object showing the vpnglobal that can be bound to authenticationradiuspolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationradiuspolicy_vpnglobal_binding
func (s *AuthenticationService) GetAllAuthenticationRADIUSPolicyVPNGlobalBinding() ([]models.AuthenticationRADIUSPolicyVPNGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationRADIUSPolicyVPNGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationRADIUSPolicyVPNGlobalBinding `json:"authenticationradiuspolicy_vpnglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationRADIUSPolicyVPNGlobalBinding(name string) (*models.AuthenticationRADIUSPolicyVPNGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationRADIUSPolicyVPNGlobalBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationRADIUSPolicyVPNGlobalBinding `json:"authenticationradiuspolicy_vpnglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationradiuspolicy_vpnglobal_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationRADIUSPolicyVPNGlobalBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationRADIUSPolicyVPNGlobalBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationRADIUSPolicyVPNGlobalBinding `json:"authenticationradiuspolicy_vpnglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationradiuspolicy_vpnvserver_binding
// Binding object showing the vpnvserver that can be bound to authenticationradiuspolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationradiuspolicy_vpnvserver_binding
func (s *AuthenticationService) GetAllAuthenticationRADIUSPolicyVPNVServerBinding() ([]models.AuthenticationRADIUSPolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationRADIUSPolicyVPNVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationRADIUSPolicyVPNVServerBinding `json:"authenticationradiuspolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationRADIUSPolicyVPNVServerBinding(name string) (*models.AuthenticationRADIUSPolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationRADIUSPolicyVPNVServerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationRADIUSPolicyVPNVServerBinding `json:"authenticationradiuspolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationradiuspolicy_vpnvserver_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationRADIUSPolicyVPNVServerBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationRADIUSPolicyVPNVServerBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationRADIUSPolicyVPNVServerBinding `json:"authenticationradiuspolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationsamlaction
// Configuration for AAA Saml action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationsamlaction
func (s *AuthenticationService) AddAuthenticationSAMLAction(resource models.AuthenticationSAMLAction) error {
	payload := map[string]any{"authenticationsamlaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationSAMLActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationSAMLAction(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", authenticationSAMLActionURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UpdateAuthenticationSAMLAction(resource models.AuthenticationSAMLAction) error {
	payload := map[string]any{"authenticationsamlaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, authenticationSAMLActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UnsetAuthenticationSAMLAction(resource models.AuthenticationSAMLAction) error {
	payload := map[string]any{"authenticationsamlaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", authenticationSAMLActionURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationSAMLAction() ([]models.AuthenticationSAMLAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationSAMLActionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationSAMLAction `json:"authenticationsamlaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationSAMLAction(name string) (*models.AuthenticationSAMLAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationSAMLActionURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationSAMLAction `json:"authenticationsamlaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationsamlaction %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationSAMLAction() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationSAMLActionURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationSAMLAction `json:"authenticationsamlaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationsamlidppolicy
// Configuration for AAA Saml IdentityProvider (IdP) policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationsamlidppolicy
func (s *AuthenticationService) AddAuthenticationSAMLIdPPolicy(resource models.AuthenticationSAMLIDPPolicy) error {
	payload := map[string]any{"authenticationsamlidppolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationSAMLIdPPolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationSAMLIdPPolicy(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", authenticationSAMLIdPPolicyURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UpdateAuthenticationSAMLIdPPolicy(resource models.AuthenticationSAMLIDPPolicy) error {
	payload := map[string]any{"authenticationsamlidppolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, authenticationSAMLIdPPolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UnsetAuthenticationSAMLIdPPolicy(resource models.AuthenticationSAMLIDPPolicy) error {
	payload := map[string]any{"authenticationsamlidppolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", authenticationSAMLIdPPolicyURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationSAMLIdPPolicy() ([]models.AuthenticationSAMLIDPPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationSAMLIdPPolicyURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationSAMLIDPPolicy `json:"authenticationsamlidppolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationSAMLIdPPolicy(name string) (*models.AuthenticationSAMLIDPPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationSAMLIdPPolicyURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationSAMLIDPPolicy `json:"authenticationsamlidppolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationsamlidppolicy %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationSAMLIdPPolicy() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationSAMLIdPPolicyURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationSAMLIDPPolicy `json:"authenticationsamlidppolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

func (s *AuthenticationService) RenameAuthenticationSAMLIdPPolicy(name string, newname string) error {
	payload := map[string]any{
		"authenticationsamlidppolicy": map[string]any{
			"name":    name,
			"newname": newname,
		},
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=rename", authenticationSAMLIdPPolicyURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// authenticationsamlidppolicy_authenticationvserver_binding
// Binding object showing the authenticationvserver that can be bound to authenticationsamlidppolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationsamlidppolicy_authenticationvserver_binding
func (s *AuthenticationService) GetAllAuthenticationSAMLIdPPolicyAuthenticationVServerBinding() ([]models.AuthenticationSAMLIDPPolicyAuthenticationVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationSAMLIdPPolicyAuthenticationVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationSAMLIDPPolicyAuthenticationVServerBinding `json:"authenticationsamlidppolicy_authenticationvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationSAMLIdPPolicyAuthenticationVServerBinding(name string) (*models.AuthenticationSAMLIDPPolicyAuthenticationVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationSAMLIdPPolicyAuthenticationVServerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationSAMLIDPPolicyAuthenticationVServerBinding `json:"authenticationsamlidppolicy_authenticationvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationsamlidppolicy_authenticationvserver_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationSAMLIdPPolicyAuthenticationVServerBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationSAMLIdPPolicyAuthenticationVServerBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationSAMLIDPPolicyAuthenticationVServerBinding `json:"authenticationsamlidppolicy_authenticationvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationsamlidppolicy_binding
// Binding object which returns the resources bound to authenticationsamlidppolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationsamlidppolicy_binding
func (s *AuthenticationService) GetAllAuthenticationSAMLIdPPolicyBinding() ([]models.AuthenticationSAMLIDPPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationSAMLIdPPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationSAMLIDPPolicyBinding `json:"authenticationsamlidppolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationSAMLIdPPolicyBinding(name string) (*models.AuthenticationSAMLIDPPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationSAMLIdPPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationSAMLIDPPolicyBinding `json:"authenticationsamlidppolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationsamlidppolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

// authenticationsamlidppolicy_vpnvserver_binding
// Binding object showing the vpnvserver that can be bound to authenticationsamlidppolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationsamlidppolicy_vpnvserver_binding
func (s *AuthenticationService) GetAllAuthenticationSAMLIdPPolicyVPNVServerBinding() ([]models.AuthenticationSAMLIDPPolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationSAMLIdPPolicyVPNVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationSAMLIDPPolicyVPNVServerBinding `json:"authenticationsamlidppolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationSAMLIdPPolicyVPNVServerBinding(name string) (*models.AuthenticationSAMLIDPPolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationSAMLIdPPolicyVPNVServerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationSAMLIDPPolicyVPNVServerBinding `json:"authenticationsamlidppolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationsamlidppolicy_vpnvserver_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationSAMLIdPPolicyVPNVServerBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationSAMLIdPPolicyVPNVServerBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationSAMLIDPPolicyVPNVServerBinding `json:"authenticationsamlidppolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationsamlidpprofile
// Configuration for AAA Saml IdentityProvider (IdP) profile resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationsamlidpprofile
func (s *AuthenticationService) AddAuthenticationSAMLIdPProfile(resource models.AuthenticationSAMLIDPProfile) error {
	payload := map[string]any{"authenticationsamlidpprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationSAMLIdPProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationSAMLIdPProfile(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", authenticationSAMLIdPProfileURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UpdateAuthenticationSAMLIdPProfile(resource models.AuthenticationSAMLIDPProfile) error {
	payload := map[string]any{"authenticationsamlidpprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, authenticationSAMLIdPProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UnsetAuthenticationSAMLIdPProfile(resource models.AuthenticationSAMLIDPProfile) error {
	payload := map[string]any{"authenticationsamlidpprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", authenticationSAMLIdPProfileURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationSAMLIdPProfile() ([]models.AuthenticationSAMLIDPProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationSAMLIdPProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationSAMLIDPProfile `json:"authenticationsamlidpprofile"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationSAMLIdPProfile(name string) (*models.AuthenticationSAMLIDPProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationSAMLIdPProfileURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationSAMLIDPProfile `json:"authenticationsamlidpprofile"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationsamlidpprofile %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationSAMLIdPProfile() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationSAMLIdPProfileURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationSAMLIDPProfile `json:"authenticationsamlidpprofile"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticaitonsamlpolicy
// Configuration for AAA Saml policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationsamlpolicy
func (s *AuthenticationService) AddAuthenticationSAMLPolicy(resource models.AuthenticationSAMLPolicy) error {
	payload := map[string]any{"authenticationsamlpolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationSAMLPolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationSAMLPolicy(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", authenticationSAMLPolicyURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UpdateAuthenticationSAMLPolicy(resource models.AuthenticationSAMLPolicy) error {
	payload := map[string]any{"authenticationsamlpolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, authenticationSAMLPolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UnsetAuthenticationSAMLPolicy(resource models.AuthenticationSAMLPolicy) error {
	payload := map[string]any{"authenticationsamlpolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", authenticationSAMLPolicyURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationSAMLPolicy() ([]models.AuthenticationSAMLPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationSAMLPolicyURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationSAMLPolicy `json:"authenticationsamlpolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationSAMLPolicy(name string) (*models.AuthenticationSAMLPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationSAMLPolicyURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationSAMLPolicy `json:"authenticationsamlpolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationsamlpolicy %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationSAMLPolicy() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationSAMLPolicyURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationSAMLPolicy `json:"authenticationsamlpolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationsamlpolicy_authenticationvserver_binding
// Binding object showing the authenticationvserver that can be bound to authenticationsamlpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationsamlpolicy_authenticationvserver_binding
func (s *AuthenticationService) GetAllAuthenticationSAMLPolicyAuthenticationVServerBinding() ([]models.AuthenticationSAMLPolicyAuthenticationVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationSAMLPolicyAuthenticationVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationSAMLPolicyAuthenticationVServerBinding `json:"authenticationsamlpolicy_authenticationvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationSAMLPolicyAuthenticationVServerBinding(name string) (*models.AuthenticationSAMLPolicyAuthenticationVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationSAMLPolicyAuthenticationVServerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationSAMLPolicyAuthenticationVServerBinding `json:"authenticationsamlpolicy_authenticationvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationsamlpolicy_authenticationvserver_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationSAMLPolicyAuthenticationVServerBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationSAMLPolicyAuthenticationVServerBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationSAMLPolicyAuthenticationVServerBinding `json:"authenticationsamlpolicy_authenticationvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationsamlpolicy_binding
// Binding object which returns the resources bound to authenticationsamlpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationsamlpolicy_binding
func (s *AuthenticationService) GetAllAuthenticationSAMLPolicyBinding() ([]models.AuthenticationSAMLPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationSAMLPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationSAMLPolicyBinding `json:"authenticationsamlpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationSAMLPolicyBinding(name string) (*models.AuthenticationSAMLPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationSAMLPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationSAMLPolicyBinding `json:"authenticationsamlpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationsamlpolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

// authenticationsamlpolicy_vpnglobal_binding
// Binding object showing the vpnglobal that can be bound to authenticationsamlpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationsamlpolicy_vpnglobal_binding
func (s *AuthenticationService) GetAllAuthenticationSAMLPolicyVPNGlobalBinding() ([]models.AuthenticationSAMLPolicyVPNGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationSAMLPolicyVPNGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationSAMLPolicyVPNGlobalBinding `json:"authenticationsamlpolicy_vpnglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationSAMLPolicyVPNGlobalBinding(name string) (*models.AuthenticationSAMLPolicyVPNGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationSAMLPolicyVPNGlobalBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationSAMLPolicyVPNGlobalBinding `json:"authenticationsamlpolicy_vpnglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationsamlpolicy_vpnglobal_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationSAMLPolicyVPNGlobalBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationSAMLPolicyVPNGlobalBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationSAMLPolicyVPNGlobalBinding `json:"authenticationsamlpolicy_vpnglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationsamlpolicy_vpnvserver_binding
// Binding object showing the vpnvserver that can be bound to authenticationsamlpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationsamlpolicy_vpnvserver_binding
func (s *AuthenticationService) GetAllAuthenticationSAMLPolicyVPNVServerBinding() ([]models.AuthenticationSAMLPolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationSAMLPolicyVPNVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationSAMLPolicyVPNVServerBinding `json:"authenticationsamlpolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationSAMLPolicyVPNVServerBinding(name string) (*models.AuthenticationSAMLPolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationSAMLPolicyVPNVServerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationSAMLPolicyVPNVServerBinding `json:"authenticationsamlpolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationsamlpolicy_vpnvserver_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationSAMLPolicyVPNVServerBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationSAMLPolicyVPNVServerBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationSAMLPolicyVPNVServerBinding `json:"authenticationsamlpolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationstorefrontauthaction
// Configuration for Storefront authentication action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationstorefrontauthaction
func (s *AuthenticationService) AddAuthenticationStoreFrontAuthAction(resource models.AuthenticationStoreFrontAuthAction) error {
	payload := map[string]any{"authenticationstorefrontauthaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationStoreFrontAuthActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationStoreFrontAuthAction(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", authenticationStoreFrontAuthActionURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UpdateAuthenticationStoreFrontAuthAction(resource models.AuthenticationStoreFrontAuthAction) error {
	payload := map[string]any{"authenticationstorefrontauthaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, authenticationStoreFrontAuthActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UnsetAuthenticationStoreFrontAuthAction(resource models.AuthenticationStoreFrontAuthAction) error {
	payload := map[string]any{"authenticationstorefrontauthaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", authenticationStoreFrontAuthActionURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationStoreFrontAuthAction() ([]models.AuthenticationStoreFrontAuthAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationStoreFrontAuthActionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationStoreFrontAuthAction `json:"authenticationstorefrontauthaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationStoreFrontAuthAction(name string) (*models.AuthenticationStoreFrontAuthAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationStoreFrontAuthActionURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationStoreFrontAuthAction `json:"authenticationstorefrontauthaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationstorefrontauthaction %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationStoreFrontAuthAction() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationStoreFrontAuthActionURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationStoreFrontAuthAction `json:"authenticationstorefrontauthaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationtacacsaction
// Configuration for TACACS action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationtacacsaction
func (s *AuthenticationService) AddAuthenticationTACACSAction(resource models.AuthenticationTACACSAction) error {
	payload := map[string]any{"authenticationtacacsaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationTACACSActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationTACACSAction(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", authenticationTACACSActionURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UpdateAuthenticationTACACSAction(resource models.AuthenticationTACACSAction) error {
	payload := map[string]any{"authenticationtacacsaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, authenticationTACACSActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UnsetAuthenticationTACACSAction(resource models.AuthenticationTACACSAction) error {
	payload := map[string]any{"authenticationtacacsaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", authenticationTACACSActionURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationTACACSAction() ([]models.AuthenticationTACACSAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationTACACSActionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationTACACSAction `json:"authenticationtacacsaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationTACACSAction(name string) (*models.AuthenticationTACACSAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationTACACSActionURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationTACACSAction `json:"authenticationtacacsaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationtacacsaction %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationTACACSAction() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationTACACSActionURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationTACACSAction `json:"authenticationtacacsaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationtacacspolicy
// Configuration for TACACS policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationtacacspolicy
func (s *AuthenticationService) AddAuthenticationTACACSPolicy(resource models.AuthenticationTACACSPolicy) error {
	payload := map[string]any{"authenticationtacacspolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationTACACSPolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationTACACSPolicy(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", authenticationTACACSPolicyURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UpdateAuthenticationTACACSPolicy(resource models.AuthenticationTACACSPolicy) error {
	payload := map[string]any{"authenticationtacacspolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, authenticationTACACSPolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UnsetAuthenticationTACACSPolicy(resource models.AuthenticationTACACSPolicy) error {
	payload := map[string]any{"authenticationtacacspolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", authenticationTACACSPolicyURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationTACACSPolicy() ([]models.AuthenticationTACACSPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationTACACSPolicyURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationTACACSPolicy `json:"authenticationtacacspolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationTACACSPolicy(name string) (*models.AuthenticationTACACSPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationTACACSPolicyURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationTACACSPolicy `json:"authenticationtacacspolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationtacacspolicy %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationTACACSPolicy() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationTACACSPolicyURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationTACACSPolicy `json:"authenticationtacacspolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationtacacspolicy_authenticationvserver_binding
// Binding object showing the authenticationvserver that can be bound to authenticationtacacspolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationtacacspolicy_authenticationvserver_binding
func (s *AuthenticationService) GetAllAuthenticationTACACSPolicyAuthenticationVServerBinding() ([]models.AuthenticationTACACSPolicyAuthenticationVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationTACACSPolicyAuthenticationVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationTACACSPolicyAuthenticationVServerBinding `json:"authenticationtacacspolicy_authenticationvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationTACACSPolicyAuthenticationVServerBinding(name string) (*models.AuthenticationTACACSPolicyAuthenticationVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationTACACSPolicyAuthenticationVServerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationTACACSPolicyAuthenticationVServerBinding `json:"authenticationtacacspolicy_authenticationvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationtacacspolicy_authenticationvserver_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationTACACSPolicyAuthenticationVServerBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationTACACSPolicyAuthenticationVServerBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationTACACSPolicyAuthenticationVServerBinding `json:"authenticationtacacspolicy_authenticationvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationtacacspolicy_binding
// Binding object which returns the resources bound to authenticationtacacspolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationtacacspolicy_binding
func (s *AuthenticationService) GetAllAuthenticationTACACSPolicyBinding() ([]models.AuthenticationTACACSPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationTACACSPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationTACACSPolicyBinding `json:"authenticationtacacspolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationTACACSPolicyBinding(name string) (*models.AuthenticationTACACSPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationTACACSPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationTACACSPolicyBinding `json:"authenticationtacacspolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationtacacspolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

// authenticationtacacspolicy_systemglobal_binding
// Binding object showing the systemglobal that can be bound to authenticationtacacspolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationtacacspolicy_systemglobal_binding
func (s *AuthenticationService) GetAllAuthenticationTACACSPolicySystemGlobalBinding() ([]models.AuthenticationTACACSPolicySystemGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationTACACSPolicySystemGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationTACACSPolicySystemGlobalBinding `json:"authenticationtacacspolicy_systemglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationTACACSPolicySystemGlobalBinding(name string) (*models.AuthenticationTACACSPolicySystemGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationTACACSPolicySystemGlobalBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationTACACSPolicySystemGlobalBinding `json:"authenticationtacacspolicy_systemglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationtacacspolicy_systemglobal_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationTACACSPolicySystemGlobalBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationTACACSPolicySystemGlobalBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationTACACSPolicySystemGlobalBinding `json:"authenticationtacacspolicy_systemglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationtacacspolicy_vpnglobal_binding
// Binding object showing the vpnglobal that can be bound to authenticationtacacspolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationtacacspolicy_vpnglobal_binding
func (s *AuthenticationService) GetAllAuthenticationTACACSPolicyVPNGlobalBinding() ([]models.AuthenticationTACACSPolicyVPNGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationTACACSPolicyVPNGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationTACACSPolicyVPNGlobalBinding `json:"authenticationtacacspolicy_vpnglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationTACACSPolicyVPNGlobalBinding(name string) (*models.AuthenticationTACACSPolicyVPNGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationTACACSPolicyVPNGlobalBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationTACACSPolicyVPNGlobalBinding `json:"authenticationtacacspolicy_vpnglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationtacacspolicy_vpnglobal_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationTACACSPolicyVPNGlobalBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationTACACSPolicyVPNGlobalBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationTACACSPolicyVPNGlobalBinding `json:"authenticationtacacspolicy_vpnglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationtacacspolicy_vpnvserver_binding
// Binding object showing the vpnvserver that can be bound to authenticationtacacspolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationtacacspolicy_vpnvserver_binding
func (s *AuthenticationService) GetAllAuthenticationTACACSPolicyVPNVServerBinding() ([]models.AuthenticationTACACSPolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationTACACSPolicyVPNVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationTACACSPolicyVPNVServerBinding `json:"authenticationtacacspolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationTACACSPolicyVPNVServerBinding(name string) (*models.AuthenticationTACACSPolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationTACACSPolicyVPNVServerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationTACACSPolicyVPNVServerBinding `json:"authenticationtacacspolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationtacacspolicy_vpnvserver_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationTACACSPolicyVPNVServerBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationTACACSPolicyVPNVServerBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationTACACSPolicyVPNVServerBinding `json:"authenticationtacacspolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticaitonvserver
// Configuration for authentication virtual server resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver
func (s *AuthenticationService) AddAuthenticationVServer(resource models.AuthenticationVServer) error {
	payload := map[string]any{"authenticationvserver": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationVServerURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationVServer(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", authenticationVServerURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UpdateAuthenticationVServer(resource models.AuthenticationVServer) error {
	payload := map[string]any{"authenticationvserver": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, authenticationVServerURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UnsetAuthenticationVServer(resource models.AuthenticationVServer) error {
	payload := map[string]any{"authenticationvserver": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", authenticationVServerURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) EnableAuthenticationVServer(name string) error {
	payload := map[string]any{"authenticationvserver": map[string]string{"name": name}}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=enable", authenticationVServerURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DisableAuthenticationVServer(name string) error {
	payload := map[string]any{"authenticationvserver": map[string]string{"name": name}}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=disable", authenticationVServerURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationVServer() ([]models.AuthenticationVServer, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationVServerURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServer `json:"authenticationvserver"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationVServer(name string) (*models.AuthenticationVServer, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationVServerURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServer `json:"authenticationvserver"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationvserver %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationVServer() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationVServerURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationVServer `json:"authenticationvserver"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

func (s *AuthenticationService) RenameAuthenticationVServer(name string, newname string) error {
	payload := map[string]any{
		"authenticationvserver": map[string]any{
			"name":    name,
			"newname": newname,
		},
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=rename", authenticationVServerURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// authenticationvserver_auditnslogpolicy_binding
// Binding object showing the auditnslogpolicy that can be bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_auditnslogpolicy_binding
func (s *AuthenticationService) AddAuthenticationVServerAuditNSLogPolicyBinding(resource models.AuthenticationVServerAuditNSLogPolicyBinding) error {
	payload := map[string]any{"authenticationvserver_auditnslogpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationVServerAuditNSLogPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationVServerAuditNSLogPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", authenticationVServerAuditNSLogPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationVServerAuditNSLogPolicyBinding() ([]models.AuthenticationVServerAuditNSLogPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationVServerAuditNSLogPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuditNSLogPolicyBinding `json:"authenticationvserver_auditnslogpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationVServerAuditNSLogPolicyBinding(name string) (*models.AuthenticationVServerAuditNSLogPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationVServerAuditNSLogPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuditNSLogPolicyBinding `json:"authenticationvserver_auditnslogpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationvserver_auditnslogpolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationVServerAuditNSLogPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationVServerAuditNSLogPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuditNSLogPolicyBinding `json:"authenticationvserver_auditnslogpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationvserver_auditsyslogpolicy_binding
// Binding object showing the auditsyslogpolicy that can be bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_auditsyslogpolicy_binding
func (s *AuthenticationService) AddAuthenticationVServerAuditSyslogPolicyBinding(resource models.AuthenticationVServerAuditSyslogPolicyBinding) error {
	payload := map[string]any{"authenticationvserver_auditsyslogpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationVServerAuditSyslogPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationVServerAuditSyslogPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", authenticationVServerAuditSyslogPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationVServerAuditSyslogPolicyBinding() ([]models.AuthenticationVServerAuditSyslogPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationVServerAuditSyslogPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuditSyslogPolicyBinding `json:"authenticationvserver_auditsyslogpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationVServerAuditSyslogPolicyBinding(name string) (*models.AuthenticationVServerAuditSyslogPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationVServerAuditSyslogPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuditSyslogPolicyBinding `json:"authenticationvserver_auditsyslogpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationvserver_auditsyslogpolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationVServerAuditSyslogPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationVServerAuditSyslogPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuditSyslogPolicyBinding `json:"authenticationvserver_auditsyslogpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationvserver_authenticationcertpolicy_binding
// Binding object showing the authenticationcertpolicy that can be bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_authenticationcertpolicy_binding
func (s *AuthenticationService) AddAuthenticationVServerAuthenticationCertPolicyBinding(resource models.AuthenticationVServerAuthenticationCertPolicyBinding) error {
	payload := map[string]any{"authenticationvserver_authenticationcertpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationVServerAuthenticationCertPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationVServerAuthenticationCertPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", authenticationVServerAuthenticationCertPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationVServerAuthenticationCertPolicyBinding() ([]models.AuthenticationVServerAuthenticationCertPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationVServerAuthenticationCertPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuthenticationCertPolicyBinding `json:"authenticationvserver_authenticationcertpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationVServerAuthenticationCertPolicyBinding(name string) (*models.AuthenticationVServerAuthenticationCertPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationVServerAuthenticationCertPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuthenticationCertPolicyBinding `json:"authenticationvserver_authenticationcertpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationvserver_authenticationcertpolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationVServerAuthenticationCertPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationVServerAuthenticationCertPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuthenticationCertPolicyBinding `json:"authenticationvserver_authenticationcertpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationvserver_authenticationldappolicy_binding
// Binding object showing the authenticationldappolicy that can be bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_authenticationldappolicy_binding
func (s *AuthenticationService) AddAuthenticationVServerAuthenticationLDAPPolicyBinding(resource models.AuthenticationVServerAuthenticationLDAPPolicyBinding) error {
	payload := map[string]any{"authenticationvserver_authenticationldappolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationVServerAuthenticationLDAPPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationVServerAuthenticationLDAPPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", authenticationVServerAuthenticationLDAPPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationVServerAuthenticationLDAPPolicyBinding() ([]models.AuthenticationVServerAuthenticationLDAPPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationVServerAuthenticationLDAPPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuthenticationLDAPPolicyBinding `json:"authenticationvserver_authenticationldappolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationVServerAuthenticationLDAPPolicyBinding(name string) (*models.AuthenticationVServerAuthenticationLDAPPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationVServerAuthenticationLDAPPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuthenticationLDAPPolicyBinding `json:"authenticationvserver_authenticationldappolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationvserver_authenticationldappolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationVServerAuthenticationLDAPPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationVServerAuthenticationLDAPPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuthenticationLDAPPolicyBinding `json:"authenticationvserver_authenticationldappolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationvserver_authenticationlocalpolicy_binding
// Binding object showing the authenticationlocalpolicy that can be bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_authenticationlocalpolicy_binding
func (s *AuthenticationService) AddAuthenticationVServerAuthenticationLocalPolicyBinding(resource models.AuthenticationVServerAuthenticationLocalPolicyBinding) error {
	payload := map[string]any{"authenticationvserver_authenticationlocalpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationVServerAuthenticationLocalPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationVServerAuthenticationLocalPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", authenticationVServerAuthenticationLocalPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationVServerAuthenticationLocalPolicyBinding() ([]models.AuthenticationVServerAuthenticationLocalPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationVServerAuthenticationLocalPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuthenticationLocalPolicyBinding `json:"authenticationvserver_authenticationlocalpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationVServerAuthenticationLocalPolicyBinding(name string) (*models.AuthenticationVServerAuthenticationLocalPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationVServerAuthenticationLocalPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuthenticationLocalPolicyBinding `json:"authenticationvserver_authenticationlocalpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationvserver_authenticationlocalpolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationVServerAuthenticationLocalPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationVServerAuthenticationLocalPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuthenticationLocalPolicyBinding `json:"authenticationvserver_authenticationlocalpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationvserver_authenticationloginschemapolicy_binding
// Binding object showing the authenticationloginschemapolicy that can be bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_authenticationloginschemapolicy_binding
func (s *AuthenticationService) AddAuthenticationVServerAuthenticationLoginSchemaPolicyBinding(resource models.AuthenticationVServerAuthenticationLoginSchemaPolicyBinding) error {
	payload := map[string]any{"authenticationvserver_authenticationloginschemapolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationVServerAuthenticationLoginSchemaPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationVServerAuthenticationLoginSchemaPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", authenticationVServerAuthenticationLoginSchemaPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationVServerAuthenticationLoginSchemaPolicyBinding() ([]models.AuthenticationVServerAuthenticationLoginSchemaPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationVServerAuthenticationLoginSchemaPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuthenticationLoginSchemaPolicyBinding `json:"authenticationvserver_authenticationloginschemapolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationVServerAuthenticationLoginSchemaPolicyBinding(name string) (*models.AuthenticationVServerAuthenticationLoginSchemaPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationVServerAuthenticationLoginSchemaPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuthenticationLoginSchemaPolicyBinding `json:"authenticationvserver_authenticationloginschemapolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationvserver_authenticationloginschemapolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationVServerAuthenticationLoginSchemaPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationVServerAuthenticationLoginSchemaPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuthenticationLoginSchemaPolicyBinding `json:"authenticationvserver_authenticationloginschemapolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationvserver_authenticationnegotiatepolicy_binding
// Binding object showing the authenticationnegotiatepolicy that can be bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_authenticationnegotiatepolicy_binding
func (s *AuthenticationService) AddAuthenticationVServerAuthenticationNegotiatePolicyBinding(resource models.AuthenticationVServerAuthenticationNegotiatePolicyBinding) error {
	payload := map[string]any{"authenticationvserver_authenticationnegotiatepolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationVServerAuthenticationNegotiatePolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationVServerAuthenticationNegotiatePolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", authenticationVServerAuthenticationNegotiatePolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationVServerAuthenticationNegotiatePolicyBinding() ([]models.AuthenticationVServerAuthenticationNegotiatePolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationVServerAuthenticationNegotiatePolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuthenticationNegotiatePolicyBinding `json:"authenticationvserver_authenticationnegotiatepolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationVServerAuthenticationNegotiatePolicyBinding(name string) (*models.AuthenticationVServerAuthenticationNegotiatePolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationVServerAuthenticationNegotiatePolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuthenticationNegotiatePolicyBinding `json:"authenticationvserver_authenticationnegotiatepolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationvserver_authenticationnegotiatepolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationVServerAuthenticationNegotiatePolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationVServerAuthenticationNegotiatePolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuthenticationNegotiatePolicyBinding `json:"authenticationvserver_authenticationnegotiatepolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationvserver_authenticationoauthidppolicy_binding
// Binding object showing the authenticationoauthidppolicy that can be bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_authenticationoauthidppolicy_binding
func (s *AuthenticationService) AddAuthenticationVServerAuthenticationOAuthIdPPolicyBinding(resource models.AuthenticationVServerAuthenticationOAuthIDPPolicyBinding) error {
	payload := map[string]any{"authenticationvserver_authenticationoauthidppolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationVServerAuthenticationOAuthIdPPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationVServerAuthenticationOAuthIdPPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", authenticationVServerAuthenticationOAuthIdPPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationVServerAuthenticationOAuthIdPPolicyBinding() ([]models.AuthenticationVServerAuthenticationOAuthIDPPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationVServerAuthenticationOAuthIdPPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuthenticationOAuthIDPPolicyBinding `json:"authenticationvserver_authenticationoauthidppolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationVServerAuthenticationOAuthIdPPolicyBinding(name string) (*models.AuthenticationVServerAuthenticationOAuthIDPPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationVServerAuthenticationOAuthIdPPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuthenticationOAuthIDPPolicyBinding `json:"authenticationvserver_authenticationoauthidppolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationvserver_authenticationoauthidppolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationVServerAuthenticationOAuthIdPPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationVServerAuthenticationOAuthIdPPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuthenticationOAuthIDPPolicyBinding `json:"authenticationvserver_authenticationoauthidppolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationvserver_authenticationpolicy_binding
// Binding object showing the authenticationpolicy that can be bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_authenticationpolicy_binding
func (s *AuthenticationService) AddAuthenticationVServerAuthenticationPolicyBinding(resource models.AuthenticationVServerAuthenticationPolicyBinding) error {
	payload := map[string]any{"authenticationvserver_authenticationpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationVServerAuthenticationPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationVServerAuthenticationPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", authenticationVServerAuthenticationPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationVServerAuthenticationPolicyBinding() ([]models.AuthenticationVServerAuthenticationPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationVServerAuthenticationPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuthenticationPolicyBinding `json:"authenticationvserver_authenticationpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationVServerAuthenticationPolicyBinding(name string) (*models.AuthenticationVServerAuthenticationPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationVServerAuthenticationPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuthenticationPolicyBinding `json:"authenticationvserver_authenticationpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationvserver_authenticationpolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationVServerAuthenticationPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationVServerAuthenticationPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuthenticationPolicyBinding `json:"authenticationvserver_authenticationpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationvserver_authenticationradiuspolicy_binding
// Binding object showing the authenticationradiuspolicy that can be bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_authenticationradiuspolicy_binding
func (s *AuthenticationService) AddAuthenticationVServerAuthenticationRADIUSPolicyBinding(resource models.AuthenticationVServerAuthenticationRADIUSPolicyBinding) error {
	payload := map[string]any{"authenticationvserver_authenticationradiuspolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationVServerAuthenticationRADIUSPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationVServerAuthenticationRADIUSPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", authenticationVServerAuthenticationRADIUSPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationVServerAuthenticationRADIUSPolicyBinding() ([]models.AuthenticationVServerAuthenticationRADIUSPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationVServerAuthenticationRADIUSPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuthenticationRADIUSPolicyBinding `json:"authenticationvserver_authenticationradiuspolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationVServerAuthenticationRADIUSPolicyBinding(name string) (*models.AuthenticationVServerAuthenticationRADIUSPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationVServerAuthenticationRADIUSPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuthenticationRADIUSPolicyBinding `json:"authenticationvserver_authenticationradiuspolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationvserver_authenticationradiuspolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationVServerAuthenticationRADIUSPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationVServerAuthenticationRADIUSPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuthenticationRADIUSPolicyBinding `json:"authenticationvserver_authenticationradiuspolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationvserver_authenticationsamlidppolicy_binding
// Binding object showing the authenticationsamlidppolicy that can be bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_authenticationsamlidppolicy_binding
func (s *AuthenticationService) AddAuthenticationVServerAuthenticationSAMLIdPPolicyBinding(resource models.AuthenticationVServerAuthenticationSAMLIDPPolicyBinding) error {
	payload := map[string]any{"authenticationvserver_authenticationsamlidppolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationVServerAuthenticationSAMLIdPPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationVServerAuthenticationSAMLIdPPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", authenticationVServerAuthenticationSAMLIdPPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationVServerAuthenticationSAMLIdPPolicyBinding() ([]models.AuthenticationVServerAuthenticationSAMLIDPPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationVServerAuthenticationSAMLIdPPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuthenticationSAMLIDPPolicyBinding `json:"authenticationvserver_authenticationsamlidppolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationVServerAuthenticationSAMLIdPPolicyBinding(name string) (*models.AuthenticationVServerAuthenticationSAMLIDPPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationVServerAuthenticationSAMLIdPPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuthenticationSAMLIDPPolicyBinding `json:"authenticationvserver_authenticationsamlidppolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationvserver_authenticationsamlidppolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationVServerAuthenticationSAMLIdPPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationVServerAuthenticationSAMLIdPPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuthenticationSAMLIDPPolicyBinding `json:"authenticationvserver_authenticationsamlidppolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationvserver_authenticationsamlpolicy_binding
// Binding object showing the authenticationsamlpolicy that can be bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_authenticationsamlpolicy_binding
func (s *AuthenticationService) AddAuthenticationVServerAuthenticationSAMLPolicyBinding(resource models.AuthenticationVServerAuthenticationSAMLPolicyBinding) error {
	payload := map[string]any{"authenticationvserver_authenticationsamlpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationVServerAuthenticationSAMLPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationVServerAuthenticationSAMLPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", authenticationVServerAuthenticationSAMLPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationVServerAuthenticationSAMLPolicyBinding() ([]models.AuthenticationVServerAuthenticationSAMLPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationVServerAuthenticationSAMLPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuthenticationSAMLPolicyBinding `json:"authenticationvserver_authenticationsamlpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationVServerAuthenticationSAMLPolicyBinding(name string) (*models.AuthenticationVServerAuthenticationSAMLPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationVServerAuthenticationSAMLPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuthenticationSAMLPolicyBinding `json:"authenticationvserver_authenticationsamlpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationvserver_authenticationsamlpolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationVServerAuthenticationSAMLPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationVServerAuthenticationSAMLPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuthenticationSAMLPolicyBinding `json:"authenticationvserver_authenticationsamlpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationvserver_authenticationtacacspolicy_binding
// Binding object showing the authenticationtacacspolicy that can be bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_authenticationtacacspolicy_binding
func (s *AuthenticationService) AddAuthenticationVServerAuthenticationTACACSPolicyBinding(resource models.AuthenticationVServerAuthenticationTACACSPolicyBinding) error {
	payload := map[string]any{"authenticationvserver_authenticationtacacspolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationVServerAuthenticationTACACSPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationVServerAuthenticationTACACSPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", authenticationVServerAuthenticationTACACSPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationVServerAuthenticationTACACSPolicyBinding() ([]models.AuthenticationVServerAuthenticationTACACSPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationVServerAuthenticationTACACSPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuthenticationTACACSPolicyBinding `json:"authenticationvserver_authenticationtacacspolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationVServerAuthenticationTACACSPolicyBinding(name string) (*models.AuthenticationVServerAuthenticationTACACSPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationVServerAuthenticationTACACSPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuthenticationTACACSPolicyBinding `json:"authenticationvserver_authenticationtacacspolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationvserver_authenticationtacacspolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationVServerAuthenticationTACACSPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationVServerAuthenticationTACACSPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuthenticationTACACSPolicyBinding `json:"authenticationvserver_authenticationtacacspolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationvserver_authenticationwebauthpolicy_binding
// Binding object showing the authenticationwebauthpolicy that can be bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_authenticationwebauthpolicy_binding
func (s *AuthenticationService) AddAuthenticationVServerAuthenticationWebAuthPolicyBinding(resource models.AuthenticationVServerAuthenticationWebAuthPolicyBinding) error {
	payload := map[string]any{"authenticationvserver_authenticationwebauthpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationVServerAuthenticationWebAuthPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationVServerAuthenticationWebAuthPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", authenticationVServerAuthenticationWebAuthPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationVServerAuthenticationWebAuthPolicyBinding() ([]models.AuthenticationVServerAuthenticationWebAuthPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationVServerAuthenticationWebAuthPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuthenticationWebAuthPolicyBinding `json:"authenticationvserver_authenticationwebauthpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationVServerAuthenticationWebAuthPolicyBinding(name string) (*models.AuthenticationVServerAuthenticationWebAuthPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationVServerAuthenticationWebAuthPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuthenticationWebAuthPolicyBinding `json:"authenticationvserver_authenticationwebauthpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationvserver_authenticationwebauthpolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationVServerAuthenticationWebAuthPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationVServerAuthenticationWebAuthPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationVServerAuthenticationWebAuthPolicyBinding `json:"authenticationvserver_authenticationwebauthpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationvserver_binding
// Binding object which returns the resources bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_binding
func (s *AuthenticationService) GetAllAuthenticationVServerBinding() ([]models.AuthenticationVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerBinding `json:"authenticationvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationVServerBinding(name string) (*models.AuthenticationVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationVServerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerBinding `json:"authenticationvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationvserver_binding %s not found", name)
	}

	return &result.Data[0], nil
}

// authenticationvserver_cachepolicy_binding
// Binding object showing the cachepolicy that can be bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_cachepolicy_binding
func (s *AuthenticationService) AddAuthenticationVServerCachePolicyBinding(resource models.AuthenticationVServerCachePolicyBinding) error {
	payload := map[string]any{"authenticationvserver_cachepolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationVServerCachePolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationVServerCachePolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", authenticationVServerCachePolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationVServerCachePolicyBinding() ([]models.AuthenticationVServerCachePolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationVServerCachePolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerCachePolicyBinding `json:"authenticationvserver_cachepolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationVServerCachePolicyBinding(name string) (*models.AuthenticationVServerCachePolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationVServerCachePolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerCachePolicyBinding `json:"authenticationvserver_cachepolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationvserver_cachepolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationVServerCachePolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationVServerCachePolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationVServerCachePolicyBinding `json:"authenticationvserver_cachepolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationvserver_cspolicy_binding
// Binding object showing the cspolicy that can be bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_cspolicy_binding
func (s *AuthenticationService) AddAuthenticationVServerCSPolicyBinding(resource models.AuthenticationVServerCSPolicyBinding) error {
	payload := map[string]any{"authenticationvserver_cspolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationVServerCSPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationVServerCSPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", authenticationVServerCSPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationVServerCSPolicyBinding() ([]models.AuthenticationVServerCSPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationVServerCSPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerCSPolicyBinding `json:"authenticationvserver_cspolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationVServerCSPolicyBinding(name string) (*models.AuthenticationVServerCSPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationVServerCSPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerCSPolicyBinding `json:"authenticationvserver_cspolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationvserver_cspolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationVServerCSPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationVServerCSPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationVServerCSPolicyBinding `json:"authenticationvserver_cspolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationvserver_responderpolicy_binding
// Binding object showing the responderpolicy that can be bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_responderpolicy_binding
func (s *AuthenticationService) AddAuthenticationVServerResponderPolicyBinding(resource models.AuthenticationVServerResponderPolicyBinding) error {
	payload := map[string]any{"authenticationvserver_responderpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationVServerResponderPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationVServerResponderPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", authenticationVServerResponderPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationVServerResponderPolicyBinding() ([]models.AuthenticationVServerResponderPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationVServerResponderPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerResponderPolicyBinding `json:"authenticationvserver_responderpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationVServerResponderPolicyBinding(name string) (*models.AuthenticationVServerResponderPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationVServerResponderPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerResponderPolicyBinding `json:"authenticationvserver_responderpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationvserver_responderpolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationVServerResponderPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationVServerResponderPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationVServerResponderPolicyBinding `json:"authenticationvserver_responderpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationvserver_tmsessionpolicy_binding
// Binding object showing the tmsessionpolicy that can be bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_tmsessionpolicy_binding
func (s *AuthenticationService) AddAuthenticationVServerTMSessionPolicyBinding(resource models.AuthenticationVServerTMSessionPolicyBinding) error {
	payload := map[string]any{"authenticationvserver_tmsessionpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationVServerTMSessionPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationVServerTMSessionPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", authenticationVServerTMSessionPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationVServerTMSessionPolicyBinding() ([]models.AuthenticationVServerTMSessionPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationVServerTMSessionPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerTMSessionPolicyBinding `json:"authenticationvserver_tmsessionpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationVServerTMSessionPolicyBinding(name string) (*models.AuthenticationVServerTMSessionPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationVServerTMSessionPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerTMSessionPolicyBinding `json:"authenticationvserver_tmsessionpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationvserver_tmsessionpolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationVServerTMSessionPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationVServerTMSessionPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationVServerTMSessionPolicyBinding `json:"authenticationvserver_tmsessionpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationvserver_vpnportaltheme_binding
// Binding object showing the vpnportaltheme that can be bound to authenticationvserver.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationvserver_vpnportaltheme_binding
func (s *AuthenticationService) AddAuthenticationVServerVPNPortalThemeBinding(resource models.AuthenticationVServerVPNPortalThemeBinding) error {
	payload := map[string]any{"authenticationvserver_vpnportaltheme_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationVServerVPNPortalThemeBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationVServerVPNPortalThemeBinding(name string, portaltheme string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=portaltheme:%s", authenticationVServerVPNPortalThemeBindingURL, name, portaltheme), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationVServerVPNPortalThemeBinding() ([]models.AuthenticationVServerVPNPortalThemeBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationVServerVPNPortalThemeBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerVPNPortalThemeBinding `json:"authenticationvserver_vpnportaltheme_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationVServerVPNPortalThemeBinding(name string) (*models.AuthenticationVServerVPNPortalThemeBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationVServerVPNPortalThemeBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationVServerVPNPortalThemeBinding `json:"authenticationvserver_vpnportaltheme_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationvserver_vpnportaltheme_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationVServerVPNPortalThemeBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationVServerVPNPortalThemeBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationVServerVPNPortalThemeBinding `json:"authenticationvserver_vpnportaltheme_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationwebauthaction
// Configuration for Web authentication action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationwebauthaction
func (s *AuthenticationService) AddAuthenticationWebAuthAction(resource models.AuthenticationWebAuthAction) error {
	payload := map[string]any{"authenticationwebauthaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationWebAuthActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationWebAuthAction(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", authenticationWebAuthActionURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UpdateAuthenticationWebAuthAction(resource models.AuthenticationWebAuthAction) error {
	payload := map[string]any{"authenticationwebauthaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, authenticationWebAuthActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UnsetAuthenticationWebAuthAction(resource models.AuthenticationWebAuthAction) error {
	payload := map[string]any{"authenticationwebauthaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", authenticationWebAuthActionURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationWebAuthAction() ([]models.AuthenticationWebAuthAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationWebAuthActionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationWebAuthAction `json:"authenticationwebauthaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationWebAuthAction(name string) (*models.AuthenticationWebAuthAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationWebAuthActionURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationWebAuthAction `json:"authenticationwebauthaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationwebauthaction %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationWebAuthAction() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationWebAuthActionURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationWebAuthAction `json:"authenticationwebauthaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationwebauthpolicy
// Configuration for Web authentication policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationwebauthpolicy
func (s *AuthenticationService) AddAuthenticationWebAuthPolicy(resource models.AuthenticationWebAuthPolicy) error {
	payload := map[string]any{"authenticationwebauthpolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, authenticationWebAuthPolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) DeleteAuthenticationWebAuthPolicy(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", authenticationWebAuthPolicyURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) UpdateAuthenticationWebAuthPolicy(resource models.AuthenticationWebAuthPolicy) error {
	payload := map[string]any{"authenticationwebauthpolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, authenticationWebAuthPolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AuthenticationService) GetAllAuthenticationWebAuthPolicy() ([]models.AuthenticationWebAuthPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationWebAuthPolicyURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationWebAuthPolicy `json:"authenticationwebauthpolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationWebAuthPolicy(name string) (*models.AuthenticationWebAuthPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationWebAuthPolicyURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationWebAuthPolicy `json:"authenticationwebauthpolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationwebauthpolicy %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationWebAuthPolicy() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationWebAuthPolicyURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationWebAuthPolicy `json:"authenticationwebauthpolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationwebauthpolicy_authenticationvserver_binding
// Binding object showing the authenticationvserver that can be bound to authenticationwebauthpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationwebauthpolicy_authenticationvserver_binding
func (s *AuthenticationService) GetAllAuthenticationWebAuthPolicyAuthenticationVServerBinding() ([]models.AuthenticationWebAuthPolicyAuthenticationVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationWebAuthPolicyAuthenticationVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationWebAuthPolicyAuthenticationVServerBinding `json:"authenticationwebauthpolicy_authenticationvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationWebAuthPolicyAuthenticationVServerBinding(name string) (*models.AuthenticationWebAuthPolicyAuthenticationVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationWebAuthPolicyAuthenticationVServerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationWebAuthPolicyAuthenticationVServerBinding `json:"authenticationwebauthpolicy_authenticationvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationwebauthpolicy_authenticationvserver_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationWebAuthPolicyAuthenticationVServerBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationWebAuthPolicyAuthenticationVServerBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationWebAuthPolicyAuthenticationVServerBinding `json:"authenticationwebauthpolicy_authenticationvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationwebauthpolicy_binding
// Binding object which returns the resources bound to authenticationwebauthpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationwebauthpolicy_binding
func (s *AuthenticationService) GetAllAuthenticationWebAuthPolicyBinding() ([]models.AuthenticationWebAuthPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationWebAuthPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationWebAuthPolicyBinding `json:"authenticationwebauthpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationWebAuthPolicyBinding(name string) (*models.AuthenticationWebAuthPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationWebAuthPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationWebAuthPolicyBinding `json:"authenticationwebauthpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationwebauthpolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

// authenticationwebauthpolicy_systemglobal_binding
// Binding object showing the systemglobal that can be bound to authenticationwebauthpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationwebauthpolicy_systemglobal_binding
func (s *AuthenticationService) GetAllAuthenticationWebAuthPolicySystemGlobalBinding() ([]models.AuthenticationWebAuthPolicySystemGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationWebAuthPolicySystemGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationWebAuthPolicySystemGlobalBinding `json:"authenticationwebauthpolicy_systemglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationWebAuthPolicySystemGlobalBinding(name string) (*models.AuthenticationWebAuthPolicySystemGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationWebAuthPolicySystemGlobalBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationWebAuthPolicySystemGlobalBinding `json:"authenticationwebauthpolicy_systemglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationwebauthpolicy_systemglobal_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationWebAuthPolicySystemGlobalBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationWebAuthPolicySystemGlobalBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationWebAuthPolicySystemGlobalBinding `json:"authenticationwebauthpolicy_systemglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationwebauthpolicy_vpnglobal_binding
// Binding object showing the vpnglobal that can be bound to authenticationwebauthpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationwebauthpolicy_vpnglobal_binding
func (s *AuthenticationService) GetAllAuthenticationWebAuthPolicyVPNGlobalBinding() ([]models.AuthenticationWebAuthPolicyVPNGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationWebAuthPolicyVPNGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationWebAuthPolicyVPNGlobalBinding `json:"authenticationwebauthpolicy_vpnglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationWebAuthPolicyVPNGlobalBinding(name string) (*models.AuthenticationWebAuthPolicyVPNGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationWebAuthPolicyVPNGlobalBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationWebAuthPolicyVPNGlobalBinding `json:"authenticationwebauthpolicy_vpnglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationwebauthpolicy_vpnglobal_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationWebAuthPolicyVPNGlobalBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationWebAuthPolicyVPNGlobalBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationWebAuthPolicyVPNGlobalBinding `json:"authenticationwebauthpolicy_vpnglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// authenticationwebauthpolicy_vpnvserver_binding
// Binding object showing the vpnvserver that can be bound to authenticationwebauthpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authentication/authenticationwebauthpolicy_vpnvserver_binding
func (s *AuthenticationService) GetAllAuthenticationWebAuthPolicyVPNVServerBinding() ([]models.AuthenticationWebAuthPolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, authenticationWebAuthPolicyVPNVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationWebAuthPolicyVPNVServerBinding `json:"authenticationwebauthpolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *AuthenticationService) GetAuthenticationWebAuthPolicyVPNVServerBinding(name string) (*models.AuthenticationWebAuthPolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", authenticationWebAuthPolicyVPNVServerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AuthenticationWebAuthPolicyVPNVServerBinding `json:"authenticationwebauthpolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("authenticationwebauthpolicy_vpnvserver_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *AuthenticationService) CountAuthenticationWebAuthPolicyVPNVServerBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", authenticationWebAuthPolicyVPNVServerBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AuthenticationWebAuthPolicyVPNVServerBinding `json:"authenticationwebauthpolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}
