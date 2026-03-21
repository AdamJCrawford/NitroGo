package nitrogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AdamJCrawford/NitroGo/nitrogo/models"
)

const (
	vpnAlwaysOnProfileURL                               = "/nitro/v1/config/vpnalwaysonprofile"
	vpnClientlessAccessPolicyURL                        = "/nitro/v1/config/vpnclientlessaccesspolicy"
	vpnClientlessAccessPolicyBindingURL                 = "/nitro/v1/config/vpnclientlessaccesspolicy_binding"
	vpnClientlessAccessPolicyVPNGlobalBindingURL        = "/nitro/v1/config/vpnclientlessaccesspolicy_vpnglobal_binding"
	vpnClientlessAccessPolicyVPNVServerBindingURL       = "/nitro/v1/config/vpnclientlessaccesspolicy_vpnvserver_binding"
	vpnClientlessAccessProfileURL                       = "/nitro/v1/config/vpnclientlessaccessprofile"
	vpnEPAProfileURL                                    = "/nitro/v1/config/vpnepaprofile"
	vpnEULAURL                                          = "/nitro/v1/config/vpneula"
	vpnFormSSOActionURL                                 = "/nitro/v1/config/vpnformssoaction"
	vpnFormSSOProfileURL                                = "/nitro/v1/config/vpnformssoprofile"
	vpnGlobalBindingURL                                 = "/nitro/v1/config/vpnglobal_binding"
	vpnGlobalAppControllerBindingURL                    = "/nitro/v1/config/vpnglobal_appcontroller_binding"
	vpnGlobalAppFlowPolicyBindingURL                    = "/nitro/v1/config/vpnglobal_appflowpolicy_binding"
	vpnGlobalAuditNSLogPolicyBindingURL                 = "/nitro/v1/config/vpnglobal_auditnslogpolicy_binding"
	vpnGlobalAuditSyslogPolicyBindingURL                = "/nitro/v1/config/vpnglobal_auditsyslogpolicy_binding"
	vpnGlobalAuthenticationCertPolicyBindingURL         = "/nitro/v1/config/vpnglobal_authenticationcertpolicy_binding"
	vpnGlobalAuthenticationLDAPPolicyBindingURL         = "/nitro/v1/config/vpnglobal_authenticationldappolicy_binding"
	vpnGlobalAuthenticationLocalPolicyBindingURL        = "/nitro/v1/config/vpnglobal_authenticationlocalpolicy_binding"
	vpnGlobalAuthenticationNegotiatePolicyBindingURL    = "/nitro/v1/config/vpnglobal_authenticationnegotiatepolicy_binding"
	vpnGlobalAuthenticationPolicyBindingURL             = "/nitro/v1/config/vpnglobal_authenticationpolicy_binding"
	vpnGlobalAuthenticationRADIUSPolicyBindingURL       = "/nitro/v1/config/vpnglobal_authenticationradiuspolicy_binding"
	vpnGlobalAuthenticationSAMLPolicyBindingURL         = "/nitro/v1/config/vpnglobal_authenticationsamlpolicy_binding"
	vpnGlobalAuthenticationTACACSPolicyBindingURL       = "/nitro/v1/config/vpnglobal_authenticationtacacspolicy_binding"
	vpnGlobalDomainBindingURL                           = "/nitro/v1/config/vpnglobal_domain_binding"
	vpnGlobalIntranetIP6BindingURL                      = "/nitro/v1/config/vpnglobal_intranetip6_binding"
	vpnGlobalIntranetIPBindingURL                       = "/nitro/v1/config/vpnglobal_intranetip_binding"
	vpnGlobalShareFileServerBindingURL                  = "/nitro/v1/config/vpnglobal_sharefileserver_binding"
	vpnGlobalSSLCertKeyBindingURL                       = "/nitro/v1/config/vpnglobal_sslcertkey_binding"
	vpnGlobalSTAServerBindingURL                        = "/nitro/v1/config/vpnglobal_staserver_binding"
	vpnGlobalVPNClientlessAccessPolicyBindingURL        = "/nitro/v1/config/vpnglobal_vpnclientlessaccesspolicy_binding"
	vpnGlobalVPNEULABindingURL                          = "/nitro/v1/config/vpnglobal_vpneula_binding"
	vpnGlobalVPNIntranetApplicationBindingURL           = "/nitro/v1/config/vpnglobal_vpnintranetapplication_binding"
	vpnGlobalVPNNextHopServerBindingURL                 = "/nitro/v1/config/vpnglobal_vpnnexthopserver_binding"
	vpnGlobalVPNPortalThemeBindingURL                   = "/nitro/v1/config/vpnglobal_vpnportaltheme_binding"
	vpnGlobalVPNSessionPolicyBindingURL                 = "/nitro/v1/config/vpnglobal_vpnsessionpolicy_binding"
	vpnGlobalVPNTrafficPolicyBindingURL                 = "/nitro/v1/config/vpnglobal_vpntrafficpolicy_binding"
	vpnGlobalVPNURLPolicyBindingURL                     = "/nitro/v1/config/vpnglobal_vpnurlpolicy_binding"
	vpnGlobalVPNURLBindingURL                           = "/nitro/v1/config/vpnglobal_vpnurl_binding"
	vpnICAConnectionURL                                 = "/nitro/v1/config/vpnicaconnection"
	vpnICADTLSConnectionURL                             = "/nitro/v1/config/vpnicadtlsconnection"
	vpnIntranetApplicationURL                           = "/nitro/v1/config/vpnintranetapplication"
	vpnIntranetApplicationBindingURL                    = "/nitro/v1/config/vpnintranetapplication_binding"
	vpnIntranetApplicationVPNURLBindingURL              = "/nitro/v1/config/vpnintranetapplication_vpnurl_binding"
	vpnNextHopServerURL                                 = "/nitro/v1/config/vpnnexthopserver"
	vpnNextGenVServerURL                                = "/nitro/v1/config/vpnnextgenvserver"
	vpnParameterURL                                     = "/nitro/v1/config/vpnparameter"
	vpnPCoIPConnectionURL                               = "/nitro/v1/config/vpnpcoipconnection"
	vpnPCoIPProfileURL                                  = "/nitro/v1/config/vpnpcoipprofile"
	vpnPCoIPVServerProfileURL                           = "/nitro/v1/config/vpnpcoipvserverprofile"
	vpnPortalThemeURL                                   = "/nitro/v1/config/vpnportaltheme"
	vpnSAMLSSOProfileURL                                = "/nitro/v1/config/vpnsamlssoprofile"
	vpnSessionActionURL                                 = "/nitro/v1/config/vpnsessionaction"
	vpnSessionPolicyURL                                 = "/nitro/v1/config/vpnsessionpolicy"
	vpnSessionPolicyBindingURL                          = "/nitro/v1/config/vpnsessionpolicy_binding"
	vpnSessionPolicyAAAGroupBindingURL                  = "/nitro/v1/config/vpnsessionpolicy_aaagroup_binding"
	vpnSessionPolicyAAAUserBindingURL                   = "/nitro/v1/config/vpnsessionpolicy_aaauser_binding"
	vpnSessionPolicyVPNGlobalBindingURL                 = "/nitro/v1/config/vpnsessionpolicy_vpnglobal_binding"
	vpnSessionPolicyVPNVServerBindingURL                = "/nitro/v1/config/vpnsessionpolicy_vpnvserver_binding"
	vpnSFConfigURL                                      = "/nitro/v1/config/vpnsfconfig"
	vpnStoreInfoURL                                     = "/nitro/v1/config/vpnstoreinfo"
	vpnTrafficActionURL                                 = "/nitro/v1/config/vpntrafficaction"
	vpnTrafficPolicyURL                                 = "/nitro/v1/config/vpntrafficpolicy"
	vpnTrafficPolicyBindingURL                          = "/nitro/v1/config/vpntrafficpolicy_binding"
	vpnTrafficPolicyAAAGroupBindingURL                  = "/nitro/v1/config/vpntrafficpolicy_aaagroup_binding"
	vpnTrafficPolicyAAAUserBindingURL                   = "/nitro/v1/config/vpntrafficpolicy_aaauser_binding"
	vpnTrafficPolicyVPNGlobalBindingURL                 = "/nitro/v1/config/vpntrafficpolicy_vpnglobal_binding"
	vpnTrafficPolicyVPNVServerBindingURL                = "/nitro/v1/config/vpntrafficpolicy_vpnvserver_binding"
	vpnURLURL                                           = "/nitro/v1/config/vpnurl"
	vpnURLActionURL                                     = "/nitro/v1/config/vpnurlaction"
	vpnURLBindingURL                                    = "/nitro/v1/config/vpnurl_binding"
	vpnURLMemberBindingURL                              = "/nitro/v1/config/vpnurl_member_binding"
	vpnURLVPNIntranetApplicationBindingURL              = "/nitro/v1/config/vpnurl_vpnintranetapplication_binding"
	vpnURLPolicyURL                                     = "/nitro/v1/config/vpnurlpolicy"
	vpnURLPolicyBindingURL                              = "/nitro/v1/config/vpnurlpolicy_binding"
	vpnURLPolicyAAAGroupBindingURL                      = "/nitro/v1/config/vpnurlpolicy_aaagroup_binding"
	vpnURLPolicyAAAUserBindingURL                       = "/nitro/v1/config/vpnurlpolicy_aaauser_binding"
	vpnURLPolicyVPNGlobalBindingURL                     = "/nitro/v1/config/vpnurlpolicy_vpnglobal_binding"
	vpnURLPolicyVPNVServerBindingURL                    = "/nitro/v1/config/vpnurlpolicy_vpnvserver_binding"
	vpnVServerURL                                       = "/nitro/v1/config/vpnvserver"
	vpnVServerBindingURL                                = "/nitro/v1/config/vpnvserver_binding"
	vpnVServerAAAGroupBindingURL                        = "/nitro/v1/config/vpnvserver_aaagroup_binding"
	vpnVServerAAAPreauthenticationPolicyBindingURL      = "/nitro/v1/config/vpnvserver_aaapreauthenticationpolicy_binding"
	vpnVServerAAAUserBindingURL                         = "/nitro/v1/config/vpnvserver_aaauser_binding"
	vpnVServerAnalyticsProfileBindingURL                = "/nitro/v1/config/vpnvserver_analyticsprofile_binding"
	vpnVServerAppControllerBindingURL                   = "/nitro/v1/config/vpnvserver_appcontroller_binding"
	vpnVServerAppFlowPolicyBindingURL                   = "/nitro/v1/config/vpnvserver_appflowpolicy_binding"
	vpnVServerAppQOEPolicyBindingURL                    = "/nitro/v1/config/vpnvserver_appqoepolicy_binding"
	vpnVServerAuditNSLogPolicyBindingURL                = "/nitro/v1/config/vpnvserver_auditnslogpolicy_binding"
	vpnVServerAuditSyslogPolicyBindingURL               = "/nitro/v1/config/vpnvserver_auditsyslogpolicy_binding"
	vpnVServerAuthenticationCertPolicyBindingURL        = "/nitro/v1/config/vpnvserver_authenticationcertpolicy_binding"
	vpnVServerAuthenticationFAPolicyBindingURL          = "/nitro/v1/config/vpnvserver_authenticationfapolicy_binding"
	vpnVServerAuthenticationLDAPPolicyBindingURL        = "/nitro/v1/config/vpnvserver_authenticationldappolicy_binding"
	vpnVServerAuthenticationLocalPolicyBindingURL       = "/nitro/v1/config/vpnvserver_authenticationlocalpolicy_binding"
	vpnVServerAuthenticationLoginSchemaPolicyBindingURL = "/nitro/v1/config/vpnvserver_authenticationloginschemapolicy_binding"
	vpnVServerAuthenticationNegotiatePolicyBindingURL   = "/nitro/v1/config/vpnvserver_authenticationnegotiatepolicy_binding"
	vpnVServerAuthenticationOAuthIDPPolicyBindingURL    = "/nitro/v1/config/vpnvserver_authenticationoauthidppolicy_binding"
	vpnVServerAuthenticationPolicyBindingURL            = "/nitro/v1/config/vpnvserver_authenticationpolicy_binding"
	vpnVServerAuthenticationRADIUSPolicyBindingURL      = "/nitro/v1/config/vpnvserver_authenticationradiuspolicy_binding"
	vpnVServerAuthenticationSAMLIDPPolicyBindingURL     = "/nitro/v1/config/vpnvserver_authenticationsamlidppolicy_binding"
	vpnVServerAuthenticationSAMLPolicyBindingURL        = "/nitro/v1/config/vpnvserver_authenticationsamlpolicy_binding"
	vpnVServerAuthenticationTACACSPolicyBindingURL      = "/nitro/v1/config/vpnvserver_authenticationtacacspolicy_binding"
	vpnVServerAuthenticationWebAuthPolicyBindingURL     = "/nitro/v1/config/vpnvserver_authenticationwebauthpolicy_binding"
	vpnVServerCachePolicyBindingURL                     = "/nitro/v1/config/vpnvserver_cachepolicy_binding"
	vpnVServerCMPPolicyBindingURL                       = "/nitro/v1/config/vpnvserver_cmppolicy_binding"
	vpnVServerCSPolicyBindingURL                        = "/nitro/v1/config/vpnvserver_cspolicy_binding"
	vpnVServerFEOPolicyBindingURL                       = "/nitro/v1/config/vpnvserver_feopolicy_binding"
	vpnVServerICAPolicyBindingURL                       = "/nitro/v1/config/vpnvserver_icapolicy_binding"
	vpnVServerIntranetIP6BindingURL                     = "/nitro/v1/config/vpnvserver_intranetip6_binding"
	vpnVServerIntranetIPBindingURL                      = "/nitro/v1/config/vpnvserver_intranetip_binding"
	vpnVServerLBVServerBindingURL                       = "/nitro/v1/config/vpnvserver_lbvserver_binding"
	vpnVServerNextHopVPNVServerBindingURL               = "/nitro/v1/config/vpnvserver_nexthopvpnvserver_binding"
	vpnVServerResponderPolicyBindingURL                 = "/nitro/v1/config/vpnvserver_responderpolicy_binding"
	vpnVServerRewritePolicyBindingURL                   = "/nitro/v1/config/vpnvserver_rewritepolicy_binding"
	vpnVServerShareFileServerBindingURL                 = "/nitro/v1/config/vpnvserver_sharefileserver_binding"
	vpnVServerSTAServerBindingURL                       = "/nitro/v1/config/vpnvserver_staserver_binding"
	vpnVServerTMTrafficPolicyBindingURL                 = "/nitro/v1/config/vpnvserver_tmtrafficpolicy_binding"
	vpnVServerURLFilteringPolicyBindingURL              = "/nitro/v1/config/vpnvserver_urlfilteringpolicy_binding"
	vpnVServerVPNClientlessAccessPolicyBindingURL       = "/nitro/v1/config/vpnvserver_vpnclientlessaccesspolicy_binding"
	vpnVServerVPNEPAProfileBindingURL                   = "/nitro/v1/config/vpnvserver_vpnepaprofile_binding"
	vpnVServerVPNEULABindingURL                         = "/nitro/v1/config/vpnvserver_vpneula_binding"
	vpnVServerVPNIntranetApplicationBindingURL          = "/nitro/v1/config/vpnvserver_vpnintranetapplication_binding"
	vpnVServerVPNNextHopServerBindingURL                = "/nitro/v1/config/vpnvserver_vpnnexthopserver_binding"
	vpnVServerVPNNextGenVServerBindingURL               = "/nitro/v1/config/vpnvserver_vpnnextgenvserver_binding"
	vpnVServerVPNPortalThemeBindingURL                  = "/nitro/v1/config/vpnvserver_vpnportaltheme_binding"
	vpnVServerVPNSessionPolicyBindingURL                = "/nitro/v1/config/vpnvserver_vpnsessionpolicy_binding"
	vpnVServerVPNTrafficPolicyBindingURL                = "/nitro/v1/config/vpnvserver_vpntrafficpolicy_binding"
	vpnVServerVPNURLBindingURL                          = "/nitro/v1/config/vpnvserver_vpnurl_binding"
	vpnVServerVPNURLPolicyBindingURL                    = "/nitro/v1/config/vpnvserver_vpnurlpolicy_binding"
)

// SSL VPN
// Virtual Private Network configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/vpn/vpn
type VPNService struct {
	client *Client
}

// vpnalwaysonprofile
func (s *VPNService) AddVPNAlwaysOnProfile(resource models.VPNAlwaysOnProfile) error {
	payload := map[string]any{"vpnalwaysonprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnAlwaysOnProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNAlwaysOnProfile(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnAlwaysOnProfileURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) UpdateVPNAlwaysOnProfile(resource models.VPNAlwaysOnProfile) error {
	payload := map[string]any{"vpnalwaysonprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, vpnAlwaysOnProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) UnsetVPNAlwaysOnProfile(resource models.VPNAlwaysOnProfile) error {
	payload := map[string]any{"vpnalwaysonprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", vpnAlwaysOnProfileURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNAlwaysOnProfile() ([]models.VPNAlwaysOnProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnAlwaysOnProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNAlwaysOnProfile `json:"vpnalwaysonprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNAlwaysOnProfile(name string) (*models.VPNAlwaysOnProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnAlwaysOnProfileURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNAlwaysOnProfile `json:"vpnalwaysonprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("resource not found")
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNAlwaysOnProfile() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnAlwaysOnProfileURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNAlwaysOnProfile `json:"vpnalwaysonprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) > 0 {
		return int(result.Data[0].Count), nil
	}
	return 0, fmt.Errorf("failed to parse count")
}

// vpnclientlessaccesspolicy
func (s *VPNService) AddVPNClientLessAccessPolicy(resource models.VPNClientlessAccessPolicy) error {
	payload := map[string]any{"vpnclientlessaccesspolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnClientlessAccessPolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNClientLessAccessPolicy(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnClientlessAccessPolicyURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) UpdateVPNClientLessAccessPolicy(resource models.VPNClientlessAccessPolicy) error {
	payload := map[string]any{"vpnclientlessaccesspolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, vpnClientlessAccessPolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNClientLessAccessPolicy() ([]models.VPNClientlessAccessPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnClientlessAccessPolicyURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNClientlessAccessPolicy `json:"vpnclientlessaccesspolicy"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNClientLessAccessPolicy(name string) (*models.VPNClientlessAccessPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnClientlessAccessPolicyURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNClientlessAccessPolicy `json:"vpnclientlessaccesspolicy"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("resource not found")
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNClientLessAccessPolicy() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnClientlessAccessPolicyURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNClientlessAccessPolicy `json:"vpnclientlessaccesspolicy"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) > 0 {
		return int(result.Data[0].Count), nil
	}
	return 0, fmt.Errorf("failed to parse count")
}

// vpnclientlessaccesspolicy_binding
func (s *VPNService) GetAllVPNClientLessAccessPolicyBinding() ([]models.VPNClientlessAccessPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnClientlessAccessPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNClientlessAccessPolicyBinding `json:"vpnclientlessaccesspolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNClientLessAccessPolicyBinding(name string) (*models.VPNClientlessAccessPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnClientlessAccessPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNClientlessAccessPolicyBinding `json:"vpnclientlessaccesspolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("resource not found")
	}

	return &result.Data[0], nil
}

// vpnclientlessaccesspolicy_vpnglobal_binding
func (s *VPNService) GetAllVPNClientLessAccessPolicyVPNGlobalBinding() ([]models.VPNClientlessAccessPolicyVPNGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnClientlessAccessPolicyVPNGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNClientlessAccessPolicyVPNGlobalBinding `json:"vpnclientlessaccesspolicy_vpnglobal_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNClientLessAccessPolicyVPNGlobalBinding(name string) (*models.VPNClientlessAccessPolicyVPNGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnClientlessAccessPolicyVPNGlobalBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNClientlessAccessPolicyVPNGlobalBinding `json:"vpnclientlessaccesspolicy_vpnglobal_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("resource not found")
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNClientLessAccessPolicyVPNGlobalBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnClientlessAccessPolicyVPNGlobalBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNClientlessAccessPolicyVPNGlobalBinding `json:"vpnclientlessaccesspolicy_vpnglobal_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) > 0 {
		return int(result.Data[0].Count), nil
	}
	return 0, fmt.Errorf("failed to parse count")
}

// vpnclientlessaccesspolicy_vpnvserver_binding
func (s *VPNService) GetAllVPNClientLessAccessPolicyVPNVServerBinding() ([]models.VPNClientlessAccessPolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnClientlessAccessPolicyVPNVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNClientlessAccessPolicyVPNVServerBinding `json:"vpnclientlessaccesspolicy_vpnvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNClientLessAccessPolicyVPNVServerBinding(name string) (*models.VPNClientlessAccessPolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnClientlessAccessPolicyVPNVServerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNClientlessAccessPolicyVPNVServerBinding `json:"vpnclientlessaccesspolicy_vpnvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("resource not found")
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNClientLessAccessPolicyVPNVServerBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnClientlessAccessPolicyVPNVServerBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNClientlessAccessPolicyVPNVServerBinding `json:"vpnclientlessaccesspolicy_vpnvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) > 0 {
		return int(result.Data[0].Count), nil
	}
	return 0, fmt.Errorf("failed to parse count")
}

// vpnclientlessaccessprofile
func (s *VPNService) AddVPNClientLessAccessProfile(resource models.VPNClientlessAccessProfile) error {
	payload := map[string]any{"vpnclientlessaccessprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnClientlessAccessProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNClientLessAccessProfile(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnClientlessAccessProfileURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) UpdateVPNClientLessAccessProfile(resource models.VPNClientlessAccessProfile) error {
	payload := map[string]any{"vpnclientlessaccessprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, vpnClientlessAccessProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) UnsetVPNClientLessAccessProfile(resource models.VPNClientlessAccessProfile) error {
	payload := map[string]any{"vpnclientlessaccessprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", vpnClientlessAccessProfileURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNClientLessAccessProfile() ([]models.VPNClientlessAccessProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnClientlessAccessProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNClientlessAccessProfile `json:"vpnclientlessaccessprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNClientLessAccessProfile(name string) (*models.VPNClientlessAccessProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnClientlessAccessProfileURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNClientlessAccessProfile `json:"vpnclientlessaccessprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("resource not found")
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNClientLessAccessProfile() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnClientlessAccessProfileURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNClientlessAccessProfile `json:"vpnclientlessaccessprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) > 0 {
		return int(result.Data[0].Count), nil
	}
	return 0, fmt.Errorf("failed to parse count")
}

// vpnepaprofile
func (s *VPNService) AddVPNEPAProfile(resource models.VPNEPAProfile) error {
	payload := map[string]any{"vpnepaprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnEPAProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNEPAProfile(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnEPAProfileURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNEPAProfile() ([]models.VPNEPAProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnEPAProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNEPAProfile `json:"vpnepaprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNEPAProfile(name string) (*models.VPNEPAProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnEPAProfileURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNEPAProfile `json:"vpnepaprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("resource not found")
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNEPAProfile() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnEPAProfileURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNEPAProfile `json:"vpnepaprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) > 0 {
		return int(result.Data[0].Count), nil
	}
	return 0, fmt.Errorf("failed to parse count")
}

// vpneula
func (s *VPNService) AddVPNEULA(resource models.VPNEULA) error {
	payload := map[string]any{"vpneula": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnEULAURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNEULA(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnEULAURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNEULA() ([]models.VPNEULA, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnEULAURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNEULA `json:"vpneula"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNEULA(name string) (*models.VPNEULA, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnEULAURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNEULA `json:"vpneula"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("resource not found")
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNEULA() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnEULAURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNEULA `json:"vpneula"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) > 0 {
		return int(result.Data[0].Count), nil
	}
	return 0, fmt.Errorf("failed to parse count")
}

// vpnformssoaction
func (s *VPNService) AddVPNFormSSOAction(resource models.VPNFormSSOAction) error {
	payload := map[string]any{"vpnformssoaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnFormSSOActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNFormSSOAction(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnFormSSOActionURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) UpdateVPNFormSSOAction(resource models.VPNFormSSOAction) error {
	payload := map[string]any{"vpnformssoaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, vpnFormSSOActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) UnsetVPNFormSSOAction(resource models.VPNFormSSOAction) error {
	payload := map[string]any{"vpnformssoaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", vpnFormSSOActionURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNFormSSOAction() ([]models.VPNFormSSOAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnFormSSOActionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNFormSSOAction `json:"vpnformssoaction"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNFormSSOAction(name string) (*models.VPNFormSSOAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnFormSSOActionURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNFormSSOAction `json:"vpnformssoaction"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("resource not found")
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNFormSSOAction() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnFormSSOActionURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNFormSSOAction `json:"vpnformssoaction"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) > 0 {
		return int(result.Data[0].Count), nil
	}
	return 0, fmt.Errorf("failed to parse count")
}

// vpnglobal_appcontroller_binding
func (s *VPNService) AddVPNGlobalAppControllerBinding(resource models.VPNGlobalAppControllerBinding) error {
	payload := map[string]any{"vpnglobal_appcontroller_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnGlobalAppControllerBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNGlobalAppControllerBinding(appcontroller string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnGlobalAppControllerBindingURL, appcontroller), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetVPNGlobalAppControllerBinding(name string) (*models.VPNGlobalAppControllerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnGlobalAppControllerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNGlobalAppControllerBinding `json:"vpnglobal_appcontroller_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("resource not found")
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNGlobalAppControllerBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnGlobalAppControllerBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNGlobalAppControllerBinding `json:"vpnglobal_appcontroller_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) > 0 {
		return int(result.Data[0].Count), nil
	}
	return 0, fmt.Errorf("failed to parse count")
}

// vpnglobal_auditnslogpolicy_binding
func (s *VPNService) AddVPNGlobalAuditNSLogPolicyBinding(resource models.VPNGlobalAuditNSLogPolicyBinding) error {
	payload := map[string]any{"vpnglobal_auditnslogpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnGlobalAuditNSLogPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNGlobalAuditNSLogPolicyBinding(policyname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnGlobalAuditNSLogPolicyBindingURL, policyname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetVPNGlobalAuditNSLogPolicyBinding(name string) (*models.VPNGlobalAuditNSLogPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnGlobalAuditNSLogPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNGlobalAuditNSLogPolicyBinding `json:"vpnglobal_auditnslogpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("resource not found")
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNGlobalAuditNSLogPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnGlobalAuditNSLogPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNGlobalAuditNSLogPolicyBinding `json:"vpnglobal_auditnslogpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) > 0 {
		return int(result.Data[0].Count), nil
	}
	return 0, fmt.Errorf("failed to parse count")
}

// vpnglobal_auditsyslogpolicy_binding
func (s *VPNService) AddVPNGlobalAuditSyslogPolicyBinding(resource models.VPNGlobalAuditSyslogPolicyBinding) error {
	payload := map[string]any{"vpnglobal_auditsyslogpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnGlobalAuditSyslogPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNGlobalAuditSyslogPolicyBinding(policyname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnGlobalAuditSyslogPolicyBindingURL, policyname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetVPNGlobalAuditSyslogPolicyBinding(name string) (*models.VPNGlobalAuditSyslogPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnGlobalAuditSyslogPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNGlobalAuditSyslogPolicyBinding `json:"vpnglobal_auditsyslogpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("resource not found")
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNGlobalAuditSyslogPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnGlobalAuditSyslogPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNGlobalAuditSyslogPolicyBinding `json:"vpnglobal_auditsyslogpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) > 0 {
		return int(result.Data[0].Count), nil
	}
	return 0, fmt.Errorf("failed to parse count")
}

// vpnglobal_authenticationcertpolicy_binding
func (s *VPNService) AddVPNGlobalAuthenticationCertPolicyBinding(resource models.VPNGlobalAuthenticationCertPolicyBinding) error {
	payload := map[string]any{"vpnglobal_authenticationcertpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnGlobalAuthenticationCertPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNGlobalAuthenticationCertPolicyBinding(policyname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnGlobalAuthenticationCertPolicyBindingURL, policyname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetVPNGlobalAuthenticationCertPolicyBinding() ([]models.VPNGlobalAuthenticationCertPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnGlobalAuthenticationCertPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNGlobalAuthenticationCertPolicyBinding `json:"vpnglobal_authenticationcertpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) CountVPNGlobalAuthenticationCertPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnGlobalAuthenticationCertPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNGlobalAuthenticationCertPolicyBinding `json:"vpnglobal_authenticationcertpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnglobal_authenticationldappolicy_binding
func (s *VPNService) AddVPNGlobalAuthenticationLDAPPolicyBinding(resource models.VPNGlobalAuthenticationLDAPPolicyBinding) error {
	payload := map[string]any{"vpnglobal_authenticationldappolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnGlobalAuthenticationLDAPPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNGlobalAuthenticationLDAPPolicyBinding(policyname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnGlobalAuthenticationLDAPPolicyBindingURL, policyname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetVPNGlobalAuthenticationLDAPPolicyBinding() ([]models.VPNGlobalAuthenticationLDAPPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnGlobalAuthenticationLDAPPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNGlobalAuthenticationLDAPPolicyBinding `json:"vpnglobal_authenticationldappolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) CountVPNGlobalAuthenticationLDAPPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnGlobalAuthenticationLDAPPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNGlobalAuthenticationLDAPPolicyBinding `json:"vpnglobal_authenticationldappolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnglobal_authenticationlocalpolicy_binding
func (s *VPNService) AddVPNGlobalAuthenticationLocalPolicyBinding(resource models.VPNGlobalAuthenticationLocalPolicyBinding) error {
	payload := map[string]any{"vpnglobal_authenticationlocalpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnGlobalAuthenticationLocalPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNGlobalAuthenticationLocalPolicyBinding(policyname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnGlobalAuthenticationLocalPolicyBindingURL, policyname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetVPNGlobalAuthenticationLocalPolicyBinding() ([]models.VPNGlobalAuthenticationLocalPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnGlobalAuthenticationLocalPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNGlobalAuthenticationLocalPolicyBinding `json:"vpnglobal_authenticationlocalpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) CountVPNGlobalAuthenticationLocalPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnGlobalAuthenticationLocalPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNGlobalAuthenticationLocalPolicyBinding `json:"vpnglobal_authenticationlocalpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnglobal_authenticationnegotiatepolicy_binding
func (s *VPNService) AddVPNGlobalAuthenticationNegotiatePolicyBinding(resource models.VPNGlobalAuthenticationNegotiatePolicyBinding) error {
	payload := map[string]any{"vpnglobal_authenticationnegotiatepolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnGlobalAuthenticationNegotiatePolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNGlobalAuthenticationNegotiatePolicyBinding(policyname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnGlobalAuthenticationNegotiatePolicyBindingURL, policyname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetVPNGlobalAuthenticationNegotiatePolicyBinding() ([]models.VPNGlobalAuthenticationNegotiatePolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnGlobalAuthenticationNegotiatePolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNGlobalAuthenticationNegotiatePolicyBinding `json:"vpnglobal_authenticationnegotiatepolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) CountVPNGlobalAuthenticationNegotiatePolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnGlobalAuthenticationNegotiatePolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNGlobalAuthenticationNegotiatePolicyBinding `json:"vpnglobal_authenticationnegotiatepolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnglobal_authenticationpolicy_binding
func (s *VPNService) AddVPNGlobalAuthenticationPolicyBinding(resource models.VPNGlobalAuthenticationPolicyBinding) error {
	payload := map[string]any{"vpnglobal_authenticationpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnGlobalAuthenticationPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNGlobalAuthenticationPolicyBinding(policyname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnGlobalAuthenticationPolicyBindingURL, policyname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetVPNGlobalAuthenticationPolicyBinding() ([]models.VPNGlobalAuthenticationPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnGlobalAuthenticationPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNGlobalAuthenticationPolicyBinding `json:"vpnglobal_authenticationpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) CountVPNGlobalAuthenticationPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnGlobalAuthenticationPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNGlobalAuthenticationPolicyBinding `json:"vpnglobal_authenticationpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnglobal_authenticationradiuspolicy_binding
func (s *VPNService) AddVPNGlobalAuthenticationRADIUSPolicyBinding(resource models.VPNGlobalAuthenticationRADIUSPolicyBinding) error {
	payload := map[string]any{"vpnglobal_authenticationradiuspolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnGlobalAuthenticationRADIUSPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNGlobalAuthenticationRADIUSPolicyBinding(policyname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnGlobalAuthenticationRADIUSPolicyBindingURL, policyname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetVPNGlobalAuthenticationRADIUSPolicyBinding() ([]models.VPNGlobalAuthenticationRADIUSPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnGlobalAuthenticationRADIUSPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNGlobalAuthenticationRADIUSPolicyBinding `json:"vpnglobal_authenticationradiuspolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) CountVPNGlobalAuthenticationRADIUSPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnGlobalAuthenticationRADIUSPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNGlobalAuthenticationRADIUSPolicyBinding `json:"vpnglobal_authenticationradiuspolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnglobal_authenticationsamlpolicy_binding
func (s *VPNService) AddVPNGlobalAuthenticationSAMLPolicyBinding(resource models.VPNGlobalAuthenticationSAMLPolicyBinding) error {
	payload := map[string]any{"vpnglobal_authenticationsamlpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnGlobalAuthenticationSAMLPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNGlobalAuthenticationSAMLPolicyBinding(policyname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnGlobalAuthenticationSAMLPolicyBindingURL, policyname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetVPNGlobalAuthenticationSAMLPolicyBinding() ([]models.VPNGlobalAuthenticationSAMLPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnGlobalAuthenticationSAMLPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNGlobalAuthenticationSAMLPolicyBinding `json:"vpnglobal_authenticationsamlpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) CountVPNGlobalAuthenticationSAMLPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnGlobalAuthenticationSAMLPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNGlobalAuthenticationSAMLPolicyBinding `json:"vpnglobal_authenticationsamlpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnglobal_authenticationtacacspolicy_binding
func (s *VPNService) AddVPNGlobalAuthenticationTACACSPolicy(resource models.VPNGlobalAuthenticationTACACSPolicyBinding) error {
	payload := map[string]any{"vpnglobal_authenticationtacacspolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnGlobalAuthenticationTACACSPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNGlobalAuthenticationTACACSPolicy(policyname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnGlobalAuthenticationTACACSPolicyBindingURL, policyname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetVPNGlobalAuthenticationTACACSPolicy() ([]models.VPNGlobalAuthenticationTACACSPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnGlobalAuthenticationTACACSPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNGlobalAuthenticationTACACSPolicyBinding `json:"vpnglobal_authenticationtacacspolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) CountVPNGlobalAuthenticationTACACSPolicy() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnGlobalAuthenticationTACACSPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNGlobalAuthenticationTACACSPolicyBinding `json:"vpnglobal_authenticationtacacspolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnglobal_binding
func (s *VPNService) GetVPNGlobalBinding() (*models.VPNGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNGlobalBinding `json:"vpnglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("no vpnglobal_binding found")
	}

	return &result.Data[0], nil
}

// vpnglobal_domain_binding
func (s *VPNService) AddVPNGlobalDomainBinding(resource models.VPNGlobalDomainBinding) error {
	payload := map[string]any{"vpnglobal_domain_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnGlobalDomainBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNGlobalDomainBinding(intranetdomain string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnGlobalDomainBindingURL, intranetdomain), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetVPNGlobalDomainBinding() ([]models.VPNGlobalDomainBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnGlobalDomainBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNGlobalDomainBinding `json:"vpnglobal_domain_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) CountVPNGlobalDomainBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnGlobalDomainBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNGlobalDomainBinding `json:"vpnglobal_domain_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnglobal_intranetip6_binding
func (s *VPNService) AddVPNGlobalIntranetIP6Binding(resource models.VPNGlobalIntranetIP6Binding) error {
	payload := map[string]any{"vpnglobal_intranetip6_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnGlobalIntranetIP6BindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNGlobalIntranetIP6Binding(intranetip6 string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnGlobalIntranetIP6BindingURL, intranetip6), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetVPNGlobalIntranetIP6Binding() ([]models.VPNGlobalIntranetIP6Binding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnGlobalIntranetIP6BindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNGlobalIntranetIP6Binding `json:"vpnglobal_intranetip6_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) CountVPNGlobalIntranetIP6Binding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnGlobalIntranetIP6BindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNGlobalIntranetIP6Binding `json:"vpnglobal_intranetip6_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnglobal_intranetip_binding
func (s *VPNService) AddVPNGlobalIntranetIPBinding(resource models.VPNGlobalIntranetIPBinding) error {
	payload := map[string]any{"vpnglobal_intranetip_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnGlobalIntranetIPBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNGlobalIntranetIPBinding(intranetip string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnGlobalIntranetIPBindingURL, intranetip), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetVPNGlobalIntranetIPBinding() ([]models.VPNGlobalIntranetIPBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnGlobalIntranetIPBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNGlobalIntranetIPBinding `json:"vpnglobal_intranetip_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) CountVPNGlobalIntranetIPBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnGlobalIntranetIPBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNGlobalIntranetIPBinding `json:"vpnglobal_intranetip_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnglobal_sharefileserver_binding
func (s *VPNService) AddVPNGlobalShareFileServerBinding(resource models.VPNGlobalShareFileServerBinding) error {
	payload := map[string]any{"vpnglobal_sharefileserver_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnGlobalShareFileServerBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNGlobalShareFileServerBinding(sharefile string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnGlobalShareFileServerBindingURL, sharefile), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetVPNGlobalShareFileServerBinding() ([]models.VPNGlobalShareFileServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnGlobalShareFileServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNGlobalShareFileServerBinding `json:"vpnglobal_sharefileserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) CountVPNGlobalShareFileServerBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnGlobalShareFileServerBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNGlobalShareFileServerBinding `json:"vpnglobal_sharefileserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnglobal_sslcertkey_binding
func (s *VPNService) AddVPNGlobalSSLCertKeyBinding(resource models.VPNGlobalSSLCertKeyBinding) error {
	payload := map[string]any{"vpnglobal_sslcertkey_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnGlobalSSLCertKeyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNGlobalSSLCertKeyBinding(certkeyname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnGlobalSSLCertKeyBindingURL, certkeyname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetVPNGlobalSSLCertKeyBinding() ([]models.VPNGlobalSSLCertKeyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnGlobalSSLCertKeyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNGlobalSSLCertKeyBinding `json:"vpnglobal_sslcertkey_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) CountVPNGlobalSSLCertKeyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnGlobalSSLCertKeyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNGlobalSSLCertKeyBinding `json:"vpnglobal_sslcertkey_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnglobal_staserver_binding
func (s *VPNService) AddVPNGlobalSTAServerBinding(resource models.VPNGlobalSTAServerBinding) error {
	payload := map[string]any{"vpnglobal_staserver_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnGlobalSTAServerBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNGlobalSTAServerBinding(staserver string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnGlobalSTAServerBindingURL, staserver), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetVPNGlobalSTAServerBinding() ([]models.VPNGlobalSTAServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnGlobalSTAServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNGlobalSTAServerBinding `json:"vpnglobal_staserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) CountVPNGlobalSTAServerBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnGlobalSTAServerBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNGlobalSTAServerBinding `json:"vpnglobal_staserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnglobal_vpnclientlessaccesspolicy_binding
func (s *VPNService) AddVPNGlobalVPNClientLessAccessPolicyBinding(resource models.VPNGlobalVPNClientlessAccessPolicyBinding) error {
	payload := map[string]any{"vpnglobal_vpnclientlessaccesspolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnGlobalVPNClientlessAccessPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNGlobalVPNClientLessAccessPolicyBinding(policyname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnGlobalVPNClientlessAccessPolicyBindingURL, policyname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetVPNGlobalVPNClientLessAccessPolicyBinding() ([]models.VPNGlobalVPNClientlessAccessPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnGlobalVPNClientlessAccessPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNGlobalVPNClientlessAccessPolicyBinding `json:"vpnglobal_vpnclientlessaccesspolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) CountVPNGlobalVPNClientLessAccessPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnGlobalVPNClientlessAccessPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNGlobalVPNClientlessAccessPolicyBinding `json:"vpnglobal_vpnclientlessaccesspolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnglobal_vpneula_binding
func (s *VPNService) AddVPNGlobalVPNEULABinding(resource models.VPNGlobalVPNEULABinding) error {
	payload := map[string]any{"vpnglobal_vpneula_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnGlobalVPNEULABindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNGlobalVPNEULABinding(eula string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnGlobalVPNEULABindingURL, eula), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetVPNGlobalVPNEULABinding() ([]models.VPNGlobalVPNEULABinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnGlobalVPNEULABindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNGlobalVPNEULABinding `json:"vpnglobal_vpneula_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) CountVPNGlobalVPNEULABinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnGlobalVPNEULABindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNGlobalVPNEULABinding `json:"vpnglobal_vpneula_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnglobal_vpnintranetapplication_binding
func (s *VPNService) AddVPNGlobalVPNIntranetApplicationBinding(resource models.VPNGlobalVPNIntranetApplicationBinding) error {
	payload := map[string]any{"vpnglobal_vpnintranetapplication_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnGlobalVPNIntranetApplicationBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNGlobalVPNIntranetApplicationBinding(intranetapplication string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnGlobalVPNIntranetApplicationBindingURL, intranetapplication), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetVPNGlobalVPNIntranetApplicationBinding() ([]models.VPNGlobalVPNIntranetApplicationBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnGlobalVPNIntranetApplicationBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNGlobalVPNIntranetApplicationBinding `json:"vpnglobal_vpnintranetapplication_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) CountVPNGlobalVPNIntranetApplicationBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnGlobalVPNIntranetApplicationBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNGlobalVPNIntranetApplicationBinding `json:"vpnglobal_vpnintranetapplication_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnglobal_vpnnexthopserver_binding
func (s *VPNService) AddVPNGlobalVPNNextHopServerBinding(resource models.VPNGlobalVPNNextHopServerBinding) error {
	payload := map[string]any{"vpnglobal_vpnnexthopserver_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnGlobalVPNNextHopServerBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNGlobalVPNNextHopServerBinding(nexthopserver string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnGlobalVPNNextHopServerBindingURL, nexthopserver), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetVPNGlobalVPNNextHopServerBinding() ([]models.VPNGlobalVPNNextHopServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnGlobalVPNNextHopServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNGlobalVPNNextHopServerBinding `json:"vpnglobal_vpnnexthopserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) CountVPNGlobalVPNNextHopServerBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnGlobalVPNNextHopServerBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNGlobalVPNNextHopServerBinding `json:"vpnglobal_vpnnexthopserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnglobal_vpnportaltheme_binding
func (s *VPNService) AddVPNGlobalVPNPortalThemeBinding(resource models.VPNGlobalVPNPortalThemeBinding) error {
	payload := map[string]any{"vpnglobal_vpnportaltheme_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnGlobalVPNPortalThemeBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNGlobalVPNPortalThemeBinding(portaltheme string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnGlobalVPNPortalThemeBindingURL, portaltheme), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetVPNGlobalVPNPortalThemeBinding() ([]models.VPNGlobalVPNPortalThemeBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnGlobalVPNPortalThemeBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNGlobalVPNPortalThemeBinding `json:"vpnglobal_vpnportaltheme_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) CountVPNGlobalVPNPortalThemeBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnGlobalVPNPortalThemeBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNGlobalVPNPortalThemeBinding `json:"vpnglobal_vpnportaltheme_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnglobal_vpnsessionpolcy_binding
func (s *VPNService) AddVPNGlobalVPNSessionPolicyBinding(resource models.VPNGlobalVPNSessionPolicyBinding) error {
	payload := map[string]any{"vpnglobal_vpnsessionpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnGlobalVPNSessionPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNGlobalVPNSessionPolicyBinding(policyname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnGlobalVPNSessionPolicyBindingURL, policyname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetVPNGlobalVPNSessionPolicyBinding() ([]models.VPNGlobalVPNSessionPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnGlobalVPNSessionPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNGlobalVPNSessionPolicyBinding `json:"vpnglobal_vpnsessionpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) CountVPNGlobalVPNSessionPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnGlobalVPNSessionPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNGlobalVPNSessionPolicyBinding `json:"vpnglobal_vpnsessionpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnglobal_vpntrafficpolicy_binding
func (s *VPNService) AddVPNGlobalVPNTrafficPolicyBinding(resource models.VPNGlobalVPNTrafficPolicyBinding) error {
	payload := map[string]any{"vpnglobal_vpntrafficpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnGlobalVPNTrafficPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNGlobalVPNTrafficPolicyBinding(policyname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnGlobalVPNTrafficPolicyBindingURL, policyname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetVPNGlobalVPNTrafficPolicyBinding() ([]models.VPNGlobalVPNTrafficPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnGlobalVPNTrafficPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNGlobalVPNTrafficPolicyBinding `json:"vpnglobal_vpntrafficpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) CountVPNGlobalVPNTrafficPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnGlobalVPNTrafficPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNGlobalVPNTrafficPolicyBinding `json:"vpnglobal_vpntrafficpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnglobal_vpnurlpolicy_binding
func (s *VPNService) AddVPNGlobalVPNURLPolicyBinding(resource models.VPNGlobalVPNURLPolicyBinding) error {
	payload := map[string]any{"vpnglobal_vpnurlpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnGlobalVPNURLPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNGlobalVPNURLPolicyBinding(policyname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnGlobalVPNURLPolicyBindingURL, policyname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetVPNGlobalVPNURLPolicyBinding() ([]models.VPNGlobalVPNURLPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnGlobalVPNURLPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNGlobalVPNURLPolicyBinding `json:"vpnglobal_vpnurlpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) CountVPNGlobalVPNURLPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnGlobalVPNURLPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNGlobalVPNURLPolicyBinding `json:"vpnglobal_vpnurlpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnglobal_vpnurl_binding
func (s *VPNService) AddVPNGlobalVPNURLBinding(resource models.VPNGlobalVPNURLBinding) error {
	payload := map[string]any{"vpnglobal_vpnurl_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnGlobalVPNURLBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNGlobalVPNURLBinding(urlname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnGlobalVPNURLBindingURL, urlname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetVPNGlobalVPNURLBinding() ([]models.VPNGlobalVPNURLBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnGlobalVPNURLBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNGlobalVPNURLBinding `json:"vpnglobal_vpnurl_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) CountVPNGlobalVPNURLBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnGlobalVPNURLBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNGlobalVPNURLBinding `json:"vpnglobal_vpnurl_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnicaconnection
func (s *VPNService) KillVPNICAConnection(resource models.VPNICAConnection) error {
	payload := map[string]any{"vpnicaconnection": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=kill", vpnICAConnectionURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNICAConnection() ([]models.VPNICAConnection, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnICAConnectionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNICAConnection `json:"vpnicaconnection"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) CountVPNICAConnection() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnICAConnectionURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNICAConnection `json:"vpnicaconnection"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnicadtlsconnection
func (s *VPNService) GetAllVPNICADTLSConnection() ([]models.VPNICADTLSConnection, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnICADTLSConnectionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNICADTLSConnection `json:"vpnicadtlsconnection"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) CountVPNICADTLSConnection() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnICADTLSConnectionURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNICADTLSConnection `json:"vpnicadtlsconnection"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnintranetapplication
func (s *VPNService) AddVPNIntranetApplication(resource models.VPNIntranetApplication) error {
	payload := map[string]any{"vpnintranetapplication": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnIntranetApplicationURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNIntranetApplication(intranetapplication string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnIntranetApplicationURL, intranetapplication), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNIntranetApplication() ([]models.VPNIntranetApplication, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnIntranetApplicationURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNIntranetApplication `json:"vpnintranetapplication"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNIntranetApplication(intranetapplication string) (*models.VPNIntranetApplication, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnIntranetApplicationURL, intranetapplication), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNIntranetApplication `json:"vpnintranetapplication"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnintranetapplication %s not found", intranetapplication)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNIntranetApplication() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnIntranetApplicationURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNIntranetApplication `json:"vpnintranetapplication"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnnexthopserver
func (s *VPNService) AddVPNNextHopServer(resource models.VPNNextHopServer) error {
	payload := map[string]any{"vpnnexthopserver": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnNextHopServerURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNNextHopServer(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnNextHopServerURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNNextHopServer() ([]models.VPNNextHopServer, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnNextHopServerURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNNextHopServer `json:"vpnnexthopserver"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNNextHopServer(name string) (*models.VPNNextHopServer, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnNextHopServerURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNNextHopServer `json:"vpnnexthopserver"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnnexthopserver %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNNextHopServer() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnNextHopServerURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNNextHopServer `json:"vpnnexthopserver"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnparameter
func (s *VPNService) UpdateVPNParameter(resource models.VPNParameter) error {
	payload := map[string]any{"vpnparameter": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, vpnParameterURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) UnsetVPNParameter(resource models.VPNParameter) error {
	payload := map[string]any{"vpnparameter": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", vpnParameterURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNParameter() (*models.VPNParameter, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnParameterURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNParameter `json:"vpnparameter"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnparameter not found")
	}

	return &result.Data[0], nil
}

// vpnpcoipconnection
func (s *VPNService) KillVPNPCoIPConnection(resource models.VPNPCoIPConnection) error {
	payload := map[string]any{"vpnpcoipconnection": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=kill", vpnPCoIPConnectionURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNPCoIPConnection() ([]models.VPNPCoIPConnection, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnPCoIPConnectionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNPCoIPConnection `json:"vpnpcoipconnection"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) CountVPNPCoIPConnection() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnPCoIPConnectionURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNPCoIPConnection `json:"vpnpcoipconnection"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnpcoipprofile
func (s *VPNService) AddVPNPCoIPProfile(resource models.VPNPCoIPProfile) error {
	payload := map[string]any{"vpnpcoipprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnPCoIPProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNPCoIPProfile(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnPCoIPProfileURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) UpdateVPNPCoIPProfile(resource models.VPNPCoIPProfile) error {
	payload := map[string]any{"vpnpcoipprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, vpnPCoIPProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) UnsetVPNPCoIPProfile(resource models.VPNPCoIPProfile) error {
	payload := map[string]any{"vpnpcoipprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", vpnPCoIPProfileURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNPCoIPProfile() ([]models.VPNPCoIPProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnPCoIPProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNPCoIPProfile `json:"vpnpcoipprofile"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNPCoIPProfile(name string) (*models.VPNPCoIPProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnPCoIPProfileURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNPCoIPProfile `json:"vpnpcoipprofile"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnpcoipprofile %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNPCoIPProfile() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnPCoIPProfileURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNPCoIPProfile `json:"vpnpcoipprofile"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnpcoipvserverprofile
func (s *VPNService) AddVPNPCoIPVServerProfile(resource models.VPNPCoIPVServerProfile) error {
	payload := map[string]any{"vpnpcoipvserverprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnPCoIPVServerProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNPCoIPVServerProfile(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnPCoIPVServerProfileURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) UpdateVPNPCoIPVServerProfile(resource models.VPNPCoIPVServerProfile) error {
	payload := map[string]any{"vpnpcoipvserverprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, vpnPCoIPVServerProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) UnsetVPNPCoIPVServerProfile(resource models.VPNPCoIPVServerProfile) error {
	payload := map[string]any{"vpnpcoipvserverprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", vpnPCoIPVServerProfileURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNPCoIPVServerProfile() ([]models.VPNPCoIPVServerProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnPCoIPVServerProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNPCoIPVServerProfile `json:"vpnpcoipvserverprofile"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNPCoIPVServerProfile(name string) (*models.VPNPCoIPVServerProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnPCoIPVServerProfileURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNPCoIPVServerProfile `json:"vpnpcoipvserverprofile"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnpcoipvserverprofile %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNPCoIPVServerProfile() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnPCoIPVServerProfileURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNPCoIPVServerProfile `json:"vpnpcoipvserverprofile"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnportaltheme
func (s *VPNService) AddVPNPortalTheme(resource models.VPNPortalTheme) error {
	payload := map[string]any{"vpnportaltheme": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnPortalThemeURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNPortalTheme(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnPortalThemeURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNPortalTheme() ([]models.VPNPortalTheme, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnPortalThemeURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNPortalTheme `json:"vpnportaltheme"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNPortalTheme(name string) (*models.VPNPortalTheme, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnPortalThemeURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNPortalTheme `json:"vpnportaltheme"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnportaltheme %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNPortalTheme() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnPortalThemeURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNPortalTheme `json:"vpnportaltheme"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnsamlssoprofile
func (s *VPNService) AddVPNSAMLSSOProfile(resource models.VPNSAMLSSOProfile) error {
	payload := map[string]any{"vpnsamlssoprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnSAMLSSOProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNSAMLSSOProfile(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnSAMLSSOProfileURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) UpdateVPNSAMLSSOProfile(resource models.VPNSAMLSSOProfile) error {
	payload := map[string]any{"vpnsamlssoprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, vpnSAMLSSOProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) UnsetVPNSAMLSSOProfile(resource models.VPNSAMLSSOProfile) error {
	payload := map[string]any{"vpnsamlssoprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", vpnSAMLSSOProfileURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNSAMLSSOProfile() ([]models.VPNSAMLSSOProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnSAMLSSOProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNSAMLSSOProfile `json:"vpnsamlssoprofile"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNSAMLSSOProfile(name string) (*models.VPNSAMLSSOProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnSAMLSSOProfileURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNSAMLSSOProfile `json:"vpnsamlssoprofile"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnsamlssoprofile %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNSAMLSSOProfile() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnSAMLSSOProfileURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNSAMLSSOProfile `json:"vpnsamlssoprofile"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnsessionaction
func (s *VPNService) AddVPNSessionAction(resource models.VPNSessionAction) error {
	payload := map[string]any{"vpnsessionaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnSessionActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNSessionAction(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnSessionActionURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) UpdateVPNSessionAction(resource models.VPNSessionAction) error {
	payload := map[string]any{"vpnsessionaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, vpnSessionActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) UnsetVPNSessionAction(resource models.VPNSessionAction) error {
	payload := map[string]any{"vpnsessionaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", vpnSessionActionURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNSessionAction() ([]models.VPNSessionAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnSessionActionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNSessionAction `json:"vpnsessionaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNSessionAction(name string) (*models.VPNSessionAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnSessionActionURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNSessionAction `json:"vpnsessionaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnsessionaction %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNSessionAction() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnSessionActionURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNSessionAction `json:"vpnsessionaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnsessionpolicy
func (s *VPNService) AddVPNSessionPolicy(resource models.VPNSessionPolicy) error {
	payload := map[string]any{"vpnsessionpolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnSessionPolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNSessionPolicy(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnSessionPolicyURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) UpdateVPNSessionPolicy(resource models.VPNSessionPolicy) error {
	payload := map[string]any{"vpnsessionpolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, vpnSessionPolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) UnsetVPNSessionPolicy(resource models.VPNSessionPolicy) error {
	payload := map[string]any{"vpnsessionpolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", vpnSessionPolicyURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNSessionPolicy() ([]models.VPNSessionPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnSessionPolicyURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNSessionPolicy `json:"vpnsessionpolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNSessionPolicy(name string) (*models.VPNSessionPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnSessionPolicyURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNSessionPolicy `json:"vpnsessionpolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnsessionpolicy %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNSessionPolicy() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnSessionPolicyURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNSessionPolicy `json:"vpnsessionpolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnsesisonpolicy_aaagroup_binding
func (s *VPNService) GetAllVPNSessionPolicyAAAGroupBinding() ([]models.VPNSessionPolicyAAAGroupBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnSessionPolicyAAAGroupBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNSessionPolicyAAAGroupBinding `json:"vpnsessionpolicy_aaagroup_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNSessionPolicyAAAGroupBinding(name string) (*models.VPNSessionPolicyAAAGroupBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnSessionPolicyAAAGroupBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNSessionPolicyAAAGroupBinding `json:"vpnsessionpolicy_aaagroup_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnsessionpolicy_aaagroup_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNSessionPolicyAAAGroupBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnSessionPolicyAAAGroupBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNSessionPolicyAAAGroupBinding `json:"vpnsessionpolicy_aaagroup_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnsessionpolicy_aaauser_binding
func (s *VPNService) GetAllVPNSessionPolicyAAAUserBinding() ([]models.VPNSessionPolicyAAAUserBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnSessionPolicyAAAUserBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNSessionPolicyAAAUserBinding `json:"vpnsessionpolicy_aaauser_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNSessionPolicyAAAUserBinding(name string) (*models.VPNSessionPolicyAAAUserBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnSessionPolicyAAAUserBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNSessionPolicyAAAUserBinding `json:"vpnsessionpolicy_aaauser_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnsessionpolicy_aaauser_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNSessionPolicyAAAUserBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnSessionPolicyAAAUserBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNSessionPolicyAAAUserBinding `json:"vpnsessionpolicy_aaauser_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnsessionpolicy_binding
func (s *VPNService) GetAllVPNSessionPolicyBinding() ([]models.VPNSessionPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnSessionPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNSessionPolicyBinding `json:"vpnsessionpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNSessionPolicyBinding(name string) (*models.VPNSessionPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnSessionPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNSessionPolicyBinding `json:"vpnsessionpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnsessionpolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

// vpnsessionpolicy_vpnglobal_binding
func (s *VPNService) GetAllVPNSessionPolicyVPNGlobalBinding() ([]models.VPNSessionPolicyVPNGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnSessionPolicyVPNGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNSessionPolicyVPNGlobalBinding `json:"vpnsessionpolicy_vpnglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNSessionPolicyVPNGlobalBinding(name string) (*models.VPNSessionPolicyVPNGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnSessionPolicyVPNGlobalBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNSessionPolicyVPNGlobalBinding `json:"vpnsessionpolicy_vpnglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnsessionpolicy_vpnglobal_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNSessionPolicyVPNGlobalBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnSessionPolicyVPNGlobalBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNSessionPolicyVPNGlobalBinding `json:"vpnsessionpolicy_vpnglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnsessionpolicy_vpnvserver_binding
func (s *VPNService) GetAllVPNSessionPolicyVPNVServerBinding() ([]models.VPNSessionPolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnSessionPolicyVPNVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNSessionPolicyVPNVServerBinding `json:"vpnsessionpolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNSessionPolicyVPNVServerBinding(name string) (*models.VPNSessionPolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnSessionPolicyVPNVServerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNSessionPolicyVPNVServerBinding `json:"vpnsessionpolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnsessionpolicy_vpnvserver_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNSessionPolicyVPNVServerBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnSessionPolicyVPNVServerBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNSessionPolicyVPNVServerBinding `json:"vpnsessionpolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnsfconfig
func (s *VPNService) GetAllVPNSFConfig() ([]models.VPNSFConfig, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnSFConfigURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNSFConfig `json:"vpnsfconfig"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) CountVPNSFConfig() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnSFConfigURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNSFConfig `json:"vpnsfconfig"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnstoreinfo
func (s *VPNService) GetAllVPNStoreInfo() ([]models.VPNStoreInfo, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnStoreInfoURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNStoreInfo `json:"vpnstoreinfo"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) CountVPNStoreInfo() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnStoreInfoURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNStoreInfo `json:"vpnstoreinfo"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpntrafficaction
func (s *VPNService) AddVPNTrafficAction(resource models.VPNTrafficAction) error {
	payload := map[string]any{"vpntrafficaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnTrafficActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNTrafficAction(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnTrafficActionURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) UpdateVPNTrafficAction(resource models.VPNTrafficAction) error {
	payload := map[string]any{"vpntrafficaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, vpnTrafficActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) UnsetVPNTrafficAction(resource models.VPNTrafficAction) error {
	payload := map[string]any{"vpntrafficaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", vpnTrafficActionURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNTrafficAction() ([]models.VPNTrafficAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnTrafficActionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNTrafficAction `json:"vpntrafficaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNTrafficAction(name string) (*models.VPNTrafficAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnTrafficActionURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNTrafficAction `json:"vpntrafficaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpntrafficaction %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNTrafficAction() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnTrafficActionURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNTrafficAction `json:"vpntrafficaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpntrafficpolicy
func (s *VPNService) AddVPNTrafficPolicy(resource models.VPNTrafficPolicy) error {
	payload := map[string]any{"vpntrafficpolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnTrafficPolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNTrafficPolicy(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnTrafficPolicyURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) UpdateVPNTrafficPolicy(resource models.VPNTrafficPolicy) error {
	payload := map[string]any{"vpntrafficpolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, vpnTrafficPolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) UnsetVPNTrafficPolicy(resource models.VPNTrafficPolicy) error {
	payload := map[string]any{"vpntrafficpolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", vpnTrafficPolicyURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNTrafficPolicy() ([]models.VPNTrafficPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnTrafficPolicyURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNTrafficPolicy `json:"vpntrafficpolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNTrafficPolicy(name string) (*models.VPNTrafficPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnTrafficPolicyURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNTrafficPolicy `json:"vpntrafficpolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpntrafficpolicy %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNTrafficPolicy() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnTrafficPolicyURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNTrafficPolicy `json:"vpntrafficpolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpntrafficpolicy_aaagroup_binding
func (s *VPNService) GetAllVPNTrafficPolicyAAAGroupBinding() ([]models.VPNTrafficPolicyAAAGroupBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnTrafficPolicyAAAGroupBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNTrafficPolicyAAAGroupBinding `json:"vpntrafficpolicy_aaagroup_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNTrafficPolicyAAAGroupBinding(name string) (*models.VPNTrafficPolicyAAAGroupBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnTrafficPolicyAAAGroupBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNTrafficPolicyAAAGroupBinding `json:"vpntrafficpolicy_aaagroup_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpntrafficpolicy_aaagroup_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNTrafficPolicyAAAGroupBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnTrafficPolicyAAAGroupBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNTrafficPolicyAAAGroupBinding `json:"vpntrafficpolicy_aaagroup_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpntrafficpolicy_aaauser_binding
func (s *VPNService) GetAllVPNTrafficPolicyAAAUserBinding() ([]models.VPNTrafficPolicyAAAUserBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnTrafficPolicyAAAUserBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNTrafficPolicyAAAUserBinding `json:"vpntrafficpolicy_aaauser_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNTrafficPolicyAAAUserBinding(name string) (*models.VPNTrafficPolicyAAAUserBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnTrafficPolicyAAAUserBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNTrafficPolicyAAAUserBinding `json:"vpntrafficpolicy_aaauser_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpntrafficpolicy_aaauser_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNTrafficPolicyAAAUserBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnTrafficPolicyAAAUserBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNTrafficPolicyAAAUserBinding `json:"vpntrafficpolicy_aaauser_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpntrafficpolicy_binding
func (s *VPNService) GetAllVPNTrafficPolicyBinding() ([]models.VPNTrafficPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnTrafficPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNTrafficPolicyBinding `json:"vpntrafficpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNTrafficPolicyBinding(name string) (*models.VPNTrafficPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnTrafficPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNTrafficPolicyBinding `json:"vpntrafficpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpntrafficpolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

// vpntrafficpolicy_vpnglobal_binding
func (s *VPNService) GetAllVPNTrafficPolicyVPNGlobalBinding() ([]models.VPNTrafficPolicyVPNGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnTrafficPolicyVPNGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNTrafficPolicyVPNGlobalBinding `json:"vpntrafficpolicy_vpnglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNTrafficPolicyVPNGlobalBinding(name string) (*models.VPNTrafficPolicyVPNGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnTrafficPolicyVPNGlobalBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNTrafficPolicyVPNGlobalBinding `json:"vpntrafficpolicy_vpnglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpntrafficpolicy_vpnglobal_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNTrafficPolicyVPNGlobalBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnTrafficPolicyVPNGlobalBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNTrafficPolicyVPNGlobalBinding `json:"vpntrafficpolicy_vpnglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpntrafficpolicy_vpnvserver_binding
func (s *VPNService) GetAllVPNTrafficPolicyVPNVServerBinding() ([]models.VPNTrafficPolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnTrafficPolicyVPNVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNTrafficPolicyVPNVServerBinding `json:"vpntrafficpolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNTrafficPolicyVPNVServerBinding(name string) (*models.VPNTrafficPolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnTrafficPolicyVPNVServerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNTrafficPolicyVPNVServerBinding `json:"vpntrafficpolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpntrafficpolicy_vpnvserver_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNTrafficPolicyVPNVServerBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnTrafficPolicyVPNVServerBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNTrafficPolicyVPNVServerBinding `json:"vpntrafficpolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnurl
func (s *VPNService) AddVPNURL(resource models.VPNURL) error {
	payload := map[string]any{"vpnurl": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnURLURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNURL(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnURLURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) UpdateVPNURL(resource models.VPNURL) error {
	payload := map[string]any{"vpnurl": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, vpnURLURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) UnsetVPNURL(resource models.VPNURL) error {
	payload := map[string]any{"vpnurl": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", vpnURLURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNURL() ([]models.VPNURL, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnURLURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNURL `json:"vpnurl"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNURL(name string) (*models.VPNURL, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnURLURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNURL `json:"vpnurl"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnurl %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNURL() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnURLURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNURL `json:"vpnurl"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnurlaction
func (s *VPNService) AddVPNURLAction(resource models.VPNURLAction) error {
	payload := map[string]any{"vpnurlaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnURLActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNURLAction(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnURLActionURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) UpdateVPNURLAction(resource models.VPNURLAction) error {
	payload := map[string]any{"vpnurlaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, vpnURLActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) UnsetVPNURLAction(resource models.VPNURLAction) error {
	payload := map[string]any{"vpnurlaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", vpnURLActionURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) RenameVPNURLAction(name string, newname string) error {
	payload := map[string]any{
		"vpnurlaction": map[string]any{
			"name":    name,
			"newname": newname,
		},
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=rename", vpnURLActionURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNURLAction() ([]models.VPNURLAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnURLActionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNURLAction `json:"vpnurlaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNURLAction(name string) (*models.VPNURLAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnURLActionURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNURLAction `json:"vpnurlaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnurlaction %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNURLAction() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnURLActionURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNURLAction `json:"vpnurlaction"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnurlpolicy
func (s *VPNService) AddVPNURLPolicy(resource models.VPNURLPolicy) error {
	payload := map[string]any{"vpnurlpolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnURLPolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNURLPolicy(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnURLPolicyURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) UpdateVPNURLPolicy(resource models.VPNURLPolicy) error {
	payload := map[string]any{"vpnurlpolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, vpnURLPolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) UnsetVPNURLPolicy(resource models.VPNURLPolicy) error {
	payload := map[string]any{"vpnurlpolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", vpnURLPolicyURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) RenameVPNURLPolicy(name string, newname string) error {
	payload := map[string]any{
		"vpnurlpolicy": map[string]any{
			"name":    name,
			"newname": newname,
		},
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=rename", vpnURLPolicyURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNURLPolicy() ([]models.VPNURLPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnURLPolicyURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNURLPolicy `json:"vpnurlpolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNURLPolicy(name string) (*models.VPNURLPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnURLPolicyURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNURLPolicy `json:"vpnurlpolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnurlpolicy %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNURLPolicy() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnURLPolicyURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNURLPolicy `json:"vpnurlpolicy"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnurlpolicy_aaagroup_binding
func (s *VPNService) GetAllVPNURLPolicyAAAGroupBinding() ([]models.VPNURLPolicyAAAGroupBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnURLPolicyAAAGroupBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNURLPolicyAAAGroupBinding `json:"vpnurlpolicy_aaagroup_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNURLPolicyAAAGroupBinding(name string) (*models.VPNURLPolicyAAAGroupBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnURLPolicyAAAGroupBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNURLPolicyAAAGroupBinding `json:"vpnurlpolicy_aaagroup_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnurlpolicy_aaagroup_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNURLPolicyAAAGroupBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnURLPolicyAAAGroupBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNURLPolicyAAAGroupBinding `json:"vpnurlpolicy_aaagroup_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnurlpolicy_aaauser_binding
func (s *VPNService) GetAllVPNURLPolicyAAAUserBinding() ([]models.VPNURLPolicyAAAUserBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnURLPolicyAAAUserBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNURLPolicyAAAUserBinding `json:"vpnurlpolicy_aaauser_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNURLPolicyAAAUserBinding(name string) (*models.VPNURLPolicyAAAUserBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnURLPolicyAAAUserBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNURLPolicyAAAUserBinding `json:"vpnurlpolicy_aaauser_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnurlpolicy_aaauser_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNURLPolicyAAAUserBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnURLPolicyAAAUserBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNURLPolicyAAAUserBinding `json:"vpnurlpolicy_aaauser_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnurlpolicy_binding
func (s *VPNService) GetAllVPNURLPolicyBinding() ([]models.VPNURLPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnURLPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNURLPolicyBinding `json:"vpnurlpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNURLPolicyBinding(name string) (*models.VPNURLPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnURLPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNURLPolicyBinding `json:"vpnurlpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnurlpolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

// vpnurlpolicy_vpnglobal_binding
func (s *VPNService) GetAllVPNURLPolicyVPNGlobalBinding() ([]models.VPNURLPolicyVPNGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnURLPolicyVPNGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNURLPolicyVPNGlobalBinding `json:"vpnurlpolicy_vpnglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNURLPolicyVPNGlobalBinding(name string) (*models.VPNURLPolicyVPNGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnURLPolicyVPNGlobalBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNURLPolicyVPNGlobalBinding `json:"vpnurlpolicy_vpnglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnurlpolicy_vpnglobal_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNURLPolicyVPNGlobalBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnURLPolicyVPNGlobalBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNURLPolicyVPNGlobalBinding `json:"vpnurlpolicy_vpnglobal_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnurlpolicy_vpnvserver_binding
func (s *VPNService) GetAllVPNURLPolicyVPNVServerBinding() ([]models.VPNURLPolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnURLPolicyVPNVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNURLPolicyVPNVServerBinding `json:"vpnurlpolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNURLPolicyVPNVServerBinding(name string) (*models.VPNURLPolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnURLPolicyVPNVServerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNURLPolicyVPNVServerBinding `json:"vpnurlpolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnurlpolicy_vpnvserver_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNURLPolicyVPNVServerBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnURLPolicyVPNVServerBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNURLPolicyVPNVServerBinding `json:"vpnurlpolicy_vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver
func (s *VPNService) AddVPNVServer(resource models.VPNVServer) error {
	payload := map[string]any{"vpnvserver": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServer(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vpnVServerURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) UpdateVPNVServer(resource models.VPNVServer) error {
	payload := map[string]any{"vpnvserver": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, vpnVServerURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) UnsetVPNVServer(resource models.VPNVServer) error {
	payload := map[string]any{"vpnvserver": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", vpnVServerURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) EnableVPNVServer(name string) error {
	payload := map[string]any{"vpnvserver": map[string]string{"name": name}}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=enable", vpnVServerURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DisableVPNVServer(name string) error {
	payload := map[string]any{"vpnvserver": map[string]string{"name": name}}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=disable", vpnVServerURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) RenameVPNVServer(name string, newname string) error {
	payload := map[string]any{
		"vpnvserver": map[string]any{
			"name":    name,
			"newname": newname,
		},
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=rename", vpnVServerURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) CheckVPNVServer(name string) error {
	payload := map[string]any{"vpnvserver": map[string]string{"name": name}}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=check", vpnVServerURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServer() ([]models.VPNVServer, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServer `json:"vpnvserver"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServer(name string) (*models.VPNVServer, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServer `json:"vpnvserver"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServer() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServer `json:"vpnvserver"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_aaapreauthenticationpolicy_binding
func (s *VPNService) AddVPNVServerAAAPreauthenticationPolicyBinding(resource models.VPNVServerAAAPreauthenticationPolicyBinding) error {
	payload := map[string]any{"vpnvserver_aaapreauthenticationpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerAAAPreauthenticationPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerAAAPreauthenticationPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", vpnVServerAAAPreauthenticationPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerAAAPreauthenticationPolicyBinding() ([]models.VPNVServerAAAPreauthenticationPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerAAAPreauthenticationPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerAAAPreauthenticationPolicyBinding `json:"vpnvserver_aaapreauthenticationpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerAAAPreauthenticationPolicyBinding(name string) (*models.VPNVServerAAAPreauthenticationPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerAAAPreauthenticationPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerAAAPreauthenticationPolicyBinding `json:"vpnvserver_aaapreauthenticationpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_aaapreauthenticationpolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServerAAAPreauthenticationPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerAAAPreauthenticationPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerAAAPreauthenticationPolicyBinding `json:"vpnvserver_aaapreauthenticationpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_analyticsprofile_binding
func (s *VPNService) AddVPNVServerAnalyticsProfileBinding(resource models.VPNVServerAnalyticsProfileBinding) error {
	payload := map[string]any{"vpnvserver_analyticsprofile_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerAnalyticsProfileBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerAnalyticsProfileBinding(name string, analyticsprofile string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=analyticsprofile:%s", vpnVServerAnalyticsProfileBindingURL, name, analyticsprofile), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerAnalyticsProfileBinding() ([]models.VPNVServerAnalyticsProfileBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerAnalyticsProfileBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerAnalyticsProfileBinding `json:"vpnvserver_analyticsprofile_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerAnalyticsProfileBinding(name string) (*models.VPNVServerAnalyticsProfileBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerAnalyticsProfileBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerAnalyticsProfileBinding `json:"vpnvserver_analyticsprofile_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_analyticsprofile_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServerAnalyticsProfileBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerAnalyticsProfileBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerAnalyticsProfileBinding `json:"vpnvserver_analyticsprofile_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_appcontroller_binding
func (s *VPNService) AddVPNVServerAppControllerBinding(resource models.VPNVServerAppControllerBinding) error {
	payload := map[string]any{"vpnvserver_appcontroller_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerAppControllerBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerAppControllerBinding(name string, appcontroller string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=appcontroller:%s", vpnVServerAppControllerBindingURL, name, appcontroller), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerAppControllerBinding() ([]models.VPNVServerAppControllerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerAppControllerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerAppControllerBinding `json:"vpnvserver_appcontroller_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerAppControllerBinding(name string) (*models.VPNVServerAppControllerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerAppControllerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerAppControllerBinding `json:"vpnvserver_appcontroller_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_appcontroller_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServerAppControllerBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerAppControllerBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerAppControllerBinding `json:"vpnvserver_appcontroller_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_appflowpolicy_binding
func (s *VPNService) AddVPNVServerAppFlowPolicyBinding(resource models.VPNVServerAppFlowPolicyBinding) error {
	payload := map[string]any{"vpnvserver_appflowpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerAppFlowPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerAppFlowPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", vpnVServerAppFlowPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerAppFlowPolicyBinding() ([]models.VPNVServerAppFlowPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerAppFlowPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerAppFlowPolicyBinding `json:"vpnvserver_appflowpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerAppFlowPolicyBinding(name string) (*models.VPNVServerAppFlowPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerAppFlowPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerAppFlowPolicyBinding `json:"vpnvserver_appflowpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_appflowpolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServerAppFlowPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerAppFlowPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerAppFlowPolicyBinding `json:"vpnvserver_appflowpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_auditnslogpolicy_binding
func (s *VPNService) AddVPNVServerAuditNSLogPolicyBinding(resource models.VPNVServerAuditNSLogPolicyBinding) error {
	payload := map[string]any{"vpnvserver_auditnslogpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerAuditNSLogPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerAuditNSLogPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", vpnVServerAuditNSLogPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerAuditNSLogPolicyBinding() ([]models.VPNVServerAuditNSLogPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerAuditNSLogPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerAuditNSLogPolicyBinding `json:"vpnvserver_auditnslogpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerAuditNSLogPolicyBinding(name string) (*models.VPNVServerAuditNSLogPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerAuditNSLogPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerAuditNSLogPolicyBinding `json:"vpnvserver_auditnslogpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_auditnslogpolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServerAuditNSLogPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerAuditNSLogPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerAuditNSLogPolicyBinding `json:"vpnvserver_auditnslogpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_auditsyslogpolicy_binding
func (s *VPNService) AddVPNVServerAuditSyslogPolicyBinding(resource models.VPNVServerAuditSyslogPolicyBinding) error {
	payload := map[string]any{"vpnvserver_auditsyslogpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerAuditSyslogPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerAuditSyslogPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", vpnVServerAuditSyslogPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerAuditSyslogPolicyBinding() ([]models.VPNVServerAuditSyslogPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerAuditSyslogPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerAuditSyslogPolicyBinding `json:"vpnvserver_auditsyslogpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerAuditSyslogPolicyBinding(name string) (*models.VPNVServerAuditSyslogPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerAuditSyslogPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerAuditSyslogPolicyBinding `json:"vpnvserver_auditsyslogpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_auditsyslogpolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServerAuditSyslogPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerAuditSyslogPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerAuditSyslogPolicyBinding `json:"vpnvserver_auditsyslogpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_authenticationcertpolicy_binding
func (s *VPNService) AddVPNVServerAuthenticationCertPolicyBinding(resource models.VPNVServerAuthenticationCertPolicyBinding) error {
	payload := map[string]any{"vpnvserver_authenticationcertpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerAuthenticationCertPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerAuthenticationCertPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", vpnVServerAuthenticationCertPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerAuthenticationCertPolicyBinding() ([]models.VPNVServerAuthenticationCertPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerAuthenticationCertPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationCertPolicyBinding `json:"vpnvserver_authenticationcertpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerAuthenticationCertPolicyBinding(name string) (*models.VPNVServerAuthenticationCertPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerAuthenticationCertPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationCertPolicyBinding `json:"vpnvserver_authenticationcertpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_authenticationcertpolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServerAuthenticationCertPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerAuthenticationCertPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationCertPolicyBinding `json:"vpnvserver_authenticationcertpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_authenticationfapolicy_binding
func (s *VPNService) AddVPNVServerAuthenticationFAPolicyBinding(resource models.VPNVServerAuthenticationDFAPolicyBinding) error {
	payload := map[string]any{"vpnvserver_authenticationdfapolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerAuthenticationFAPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerAuthenticationFAPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", vpnVServerAuthenticationFAPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerAuthenticationFAPolicyBinding() ([]models.VPNVServerAuthenticationDFAPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerAuthenticationFAPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationDFAPolicyBinding `json:"vpnvserver_authenticationdfapolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerAuthenticationFAPolicyBinding(name string) (*models.VPNVServerAuthenticationDFAPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerAuthenticationFAPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationDFAPolicyBinding `json:"vpnvserver_authenticationdfapolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_authenticationfapolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServerAuthenticationFAPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerAuthenticationFAPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationDFAPolicyBinding `json:"vpnvserver_authenticationdfapolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_authenticationldappolicy_binding
func (s *VPNService) AddVPNVServerAuthenticationLDAPPolicyBinding(resource models.VPNVServerAuthenticationLDAPPolicyBinding) error {
	payload := map[string]any{"vpnvserver_authenticationldappolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerAuthenticationLDAPPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerAuthenticationLDAPPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", vpnVServerAuthenticationLDAPPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerAuthenticationLDAPPolicyBinding() ([]models.VPNVServerAuthenticationLDAPPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerAuthenticationLDAPPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationLDAPPolicyBinding `json:"vpnvserver_authenticationldappolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerAuthenticationLDAPPolicyBinding(name string) (*models.VPNVServerAuthenticationLDAPPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerAuthenticationLDAPPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationLDAPPolicyBinding `json:"vpnvserver_authenticationldappolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_authenticationldappolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServerAuthenticationLDAPPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerAuthenticationLDAPPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationLDAPPolicyBinding `json:"vpnvserver_authenticationldappolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_authenticationlocalpolicy_binding
func (s *VPNService) AddVPNVServerAuthenticationLocalPolicyBinding(resource models.VPNVServerAuthenticationLocalPolicyBinding) error {
	payload := map[string]any{"vpnvserver_authenticationlocalpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerAuthenticationLocalPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerAuthenticationLocalPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", vpnVServerAuthenticationLocalPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerAuthenticationLocalPolicyBinding() ([]models.VPNVServerAuthenticationLocalPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerAuthenticationLocalPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationLocalPolicyBinding `json:"vpnvserver_authenticationlocalpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerAuthenticationLocalPolicyBinding(name string) (*models.VPNVServerAuthenticationLocalPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerAuthenticationLocalPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationLocalPolicyBinding `json:"vpnvserver_authenticationlocalpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_authenticationlocalpolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServerAuthenticationLocalPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerAuthenticationLocalPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationLocalPolicyBinding `json:"vpnvserver_authenticationlocalpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_authenticationloginschemapolicy_binding
func (s *VPNService) AddVPNVServerAuthenticationLoginSchemaPolicyBinding(resource models.VPNVServerAuthenticationLoginSchemaPolicyBinding) error {
	payload := map[string]any{"vpnvserver_authenticationloginschemapolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerAuthenticationLoginSchemaPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerAuthenticationLoginSchemaPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", vpnVServerAuthenticationLoginSchemaPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerAuthenticationLoginSchemaPolicyBinding() ([]models.VPNVServerAuthenticationLoginSchemaPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerAuthenticationLoginSchemaPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationLoginSchemaPolicyBinding `json:"vpnvserver_authenticationloginschemapolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerAuthenticationLoginSchemaPolicyBinding(name string) (*models.VPNVServerAuthenticationLoginSchemaPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerAuthenticationLoginSchemaPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationLoginSchemaPolicyBinding `json:"vpnvserver_authenticationloginschemapolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_authenticationloginschemapolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServerAuthenticationLoginSchemaPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerAuthenticationLoginSchemaPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationLoginSchemaPolicyBinding `json:"vpnvserver_authenticationloginschemapolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_authenticationnegotiatepolcy_binding
func (s *VPNService) AddVPNVServerAuthenticationNegotiatePolicyBinding(resource models.VPNVServerAuthenticationNegotiatePolicyBinding) error {
	payload := map[string]any{"vpnvserver_authenticationnegotiatepolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerAuthenticationNegotiatePolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerAuthenticationNegotiatePolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", vpnVServerAuthenticationNegotiatePolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerAuthenticationNegotiatePolicyBinding() ([]models.VPNVServerAuthenticationNegotiatePolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerAuthenticationNegotiatePolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationNegotiatePolicyBinding `json:"vpnvserver_authenticationnegotiatepolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerAuthenticationNegotiatePolicyBinding(name string) (*models.VPNVServerAuthenticationNegotiatePolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerAuthenticationNegotiatePolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationNegotiatePolicyBinding `json:"vpnvserver_authenticationnegotiatepolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_authenticationnegotiatepolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServerAuthenticationNegotiatePolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerAuthenticationNegotiatePolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationNegotiatePolicyBinding `json:"vpnvserver_authenticationnegotiatepolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_authenticationoauthidppolicy_binding
func (s *VPNService) AddVPNVServerAuthenticationOAuthIDPPolicyBinding(resource models.VPNVServerAuthenticationOAuthIDPPolicyBinding) error {
	payload := map[string]any{"vpnvserver_authenticationoauthidppolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerAuthenticationOAuthIDPPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerAuthenticationOAuthIDPPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", vpnVServerAuthenticationOAuthIDPPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerAuthenticationOAuthIDPPolicyBinding() ([]models.VPNVServerAuthenticationOAuthIDPPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerAuthenticationOAuthIDPPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationOAuthIDPPolicyBinding `json:"vpnvserver_authenticationoauthidppolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerAuthenticationOAuthIDPPolicyBinding(name string) (*models.VPNVServerAuthenticationOAuthIDPPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerAuthenticationOAuthIDPPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationOAuthIDPPolicyBinding `json:"vpnvserver_authenticationoauthidppolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_authenticationoauthidppolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServerAuthenticationOAuthIDPPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerAuthenticationOAuthIDPPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationOAuthIDPPolicyBinding `json:"vpnvserver_authenticationoauthidppolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_authenticationpolicy_binding
func (s *VPNService) AddVPNVServerAuthenticationPolicyBinding(resource models.VPNVServerAuthenticationPolicyBinding) error {
	payload := map[string]any{"vpnvserver_authenticationpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerAuthenticationPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerAuthenticationPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", vpnVServerAuthenticationPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerAuthenticationPolicyBinding() ([]models.VPNVServerAuthenticationPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerAuthenticationPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationPolicyBinding `json:"vpnvserver_authenticationpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerAuthenticationPolicyBinding(name string) (*models.VPNVServerAuthenticationPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerAuthenticationPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationPolicyBinding `json:"vpnvserver_authenticationpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_authenticationpolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServerAuthenticationPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerAuthenticationPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationPolicyBinding `json:"vpnvserver_authenticationpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_authenticationradiuspolicy_binding
func (s *VPNService) AddVPNVServerAuthenticationRADIUSPolicyBinding(resource models.VPNVServerAuthenticationRADIUSPolicyBinding) error {
	payload := map[string]any{"vpnvserver_authenticationradiuspolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerAuthenticationRADIUSPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerAuthenticationRADIUSPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", vpnVServerAuthenticationRADIUSPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerAuthenticationRADIUSPolicyBinding() ([]models.VPNVServerAuthenticationRADIUSPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerAuthenticationRADIUSPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationRADIUSPolicyBinding `json:"vpnvserver_authenticationradiuspolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerAuthenticationRADIUSPolicyBinding(name string) (*models.VPNVServerAuthenticationRADIUSPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerAuthenticationRADIUSPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationRADIUSPolicyBinding `json:"vpnvserver_authenticationradiuspolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_authenticationradiuspolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServerAuthenticationRADIUSPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerAuthenticationRADIUSPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationRADIUSPolicyBinding `json:"vpnvserver_authenticationradiuspolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_authenticationsamlidppolicy_binding
func (s *VPNService) AddVPNVServerAuthenticationSAMLIDPPolicyBinding(resource models.VPNVServerAuthenticationSAMLIDPPolicyBinding) error {
	payload := map[string]any{"vpnvserver_authenticationsamlidppolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerAuthenticationSAMLIDPPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerAuthenticationSAMLIDPPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", vpnVServerAuthenticationSAMLIDPPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerAuthenticationSAMLIDPPolicyBinding() ([]models.VPNVServerAuthenticationSAMLIDPPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerAuthenticationSAMLIDPPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationSAMLIDPPolicyBinding `json:"vpnvserver_authenticationsamlidppolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerAuthenticationSAMLIDPPolicyBinding(name string) (*models.VPNVServerAuthenticationSAMLIDPPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerAuthenticationSAMLIDPPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationSAMLIDPPolicyBinding `json:"vpnvserver_authenticationsamlidppolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_authenticationsamlidppolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServerAuthenticationSAMLIDPPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerAuthenticationSAMLIDPPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationSAMLIDPPolicyBinding `json:"vpnvserver_authenticationsamlidppolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_authenticationsamlpolicy_binding
func (s *VPNService) AddVPNVServerAuthenticationSAMLPolicyBinding(resource models.VPNVServerAuthenticationSAMLPolicyBinding) error {
	payload := map[string]any{"vpnvserver_authenticationsamlpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerAuthenticationSAMLPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerAuthenticationSAMLPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", vpnVServerAuthenticationSAMLPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerAuthenticationSAMLPolicyBinding() ([]models.VPNVServerAuthenticationSAMLPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerAuthenticationSAMLPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationSAMLPolicyBinding `json:"vpnvserver_authenticationsamlpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerAuthenticationSAMLPolicyBinding(name string) (*models.VPNVServerAuthenticationSAMLPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerAuthenticationSAMLPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationSAMLPolicyBinding `json:"vpnvserver_authenticationsamlpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_authenticationsamlpolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServerAuthenticationSAMLPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerAuthenticationSAMLPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationSAMLPolicyBinding `json:"vpnvserver_authenticationsamlpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_authenticationtacacspolicy_binding
func (s *VPNService) AddVPNVServerAuthenticationTACACSPolicyBinding(resource models.VPNVServerAuthenticationTACACSPolicyBinding) error {
	payload := map[string]any{"vpnvserver_authenticationtacacspolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerAuthenticationTACACSPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerAuthenticationTACACSPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", vpnVServerAuthenticationTACACSPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerAuthenticationTACACSPolicyBinding() ([]models.VPNVServerAuthenticationTACACSPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerAuthenticationTACACSPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationTACACSPolicyBinding `json:"vpnvserver_authenticationtacacspolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerAuthenticationTACACSPolicyBinding(name string) (*models.VPNVServerAuthenticationTACACSPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerAuthenticationTACACSPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationTACACSPolicyBinding `json:"vpnvserver_authenticationtacacspolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_authenticationtacacspolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServerAuthenticationTACACSPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerAuthenticationTACACSPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationTACACSPolicyBinding `json:"vpnvserver_authenticationtacacspolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_authenticationwebauthpolicy_binding
func (s *VPNService) AddVPNVServerAuthenticationWebAuthPolicyBinding(resource models.VPNVServerAuthenticationWebAuthPolicyBinding) error {
	payload := map[string]any{"vpnvserver_authenticationwebauthpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerAuthenticationWebAuthPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerAuthenticationWebAuthPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", vpnVServerAuthenticationWebAuthPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerAuthenticationWebAuthPolicyBinding() ([]models.VPNVServerAuthenticationWebAuthPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerAuthenticationWebAuthPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationWebAuthPolicyBinding `json:"vpnvserver_authenticationwebauthpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerAuthenticationWebAuthPolicyBinding(name string) (*models.VPNVServerAuthenticationWebAuthPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerAuthenticationWebAuthPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationWebAuthPolicyBinding `json:"vpnvserver_authenticationwebauthpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_authenticationwebauthpolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServerAuthenticationWebAuthPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerAuthenticationWebAuthPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerAuthenticationWebAuthPolicyBinding `json:"vpnvserver_authenticationwebauthpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_binding
func (s *VPNService) GetAllVPNVServerBinding() ([]models.VPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerBinding `json:"vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerBinding(name string) (*models.VPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerBinding `json:"vpnvserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_binding %s not found", name)
	}

	return &result.Data[0], nil
}

// vpnvserver_cachepolicy_binding
func (s *VPNService) AddVPNVServerCachePolicyBinding(resource models.VPNVServerCachePolicyBinding) error {
	payload := map[string]any{"vpnvserver_cachepolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerCachePolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerCachePolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", vpnVServerCachePolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerCachePolicyBinding() ([]models.VPNVServerCachePolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerCachePolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerCachePolicyBinding `json:"vpnvserver_cachepolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerCachePolicyBinding(name string) (*models.VPNVServerCachePolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerCachePolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerCachePolicyBinding `json:"vpnvserver_cachepolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_cachepolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServerCachePolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerCachePolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerCachePolicyBinding `json:"vpnvserver_cachepolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_cspolicy_binding
func (s *VPNService) AddVPNVServerCSPolicyBinding(resource models.VPNVServerCSPolicyBinding) error {
	payload := map[string]any{"vpnvserver_cspolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerCSPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerCSPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", vpnVServerCSPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerCSPolicyBinding() ([]models.VPNVServerCSPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerCSPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerCSPolicyBinding `json:"vpnvserver_cspolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerCSPolicyBinding(name string) (*models.VPNVServerCSPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerCSPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerCSPolicyBinding `json:"vpnvserver_cspolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_cspolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServerCSPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerCSPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerCSPolicyBinding `json:"vpnvserver_cspolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_feopolicy_binding
func (s *VPNService) AddVPNVServerFEOPolicyBinding(resource models.VPNVServerFEOPolicyBinding) error {
	payload := map[string]any{"vpnvserver_feopolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerFEOPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerFEOPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", vpnVServerFEOPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerFEOPolicyBinding() ([]models.VPNVServerFEOPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerFEOPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerFEOPolicyBinding `json:"vpnvserver_feopolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerFEOPolicyBinding(name string) (*models.VPNVServerFEOPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerFEOPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerFEOPolicyBinding `json:"vpnvserver_feopolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_feopolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}
func (s *VPNService) CountVPNVServerFEOPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerFEOPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerFEOPolicyBinding `json:"vpnvserver_feopolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_icapolicy_binding
func (s *VPNService) AddVPNVServerICAPolicyBinding(resource models.VPNVServerICAPolicyBinding) error {
	payload := map[string]any{"vpnvserver_icapolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerICAPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerICAPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", vpnVServerICAPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerICAPolicyBinding() ([]models.VPNVServerICAPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerICAPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerICAPolicyBinding `json:"vpnvserver_icapolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerICAPolicyBinding(name string) (*models.VPNVServerICAPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerICAPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerICAPolicyBinding `json:"vpnvserver_icapolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_icapolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServerICAPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerICAPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerICAPolicyBinding `json:"vpnvserver_icapolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_intranetip6_binding
func (s *VPNService) AddVPNVServerIntranetIP6Binding(resource models.VPNVServerIntranetIP6Binding) error {
	payload := map[string]any{"vpnvserver_intranetip6_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerIntranetIP6BindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerIntranetIP6Binding(name string, intranetip6 string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=intranetip6:%s", vpnVServerIntranetIP6BindingURL, name, intranetip6), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerIntranetIP6Binding() ([]models.VPNVServerIntranetIP6Binding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerIntranetIP6BindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerIntranetIP6Binding `json:"vpnvserver_intranetip6_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerIntranetIP6Binding(name string) (*models.VPNVServerIntranetIP6Binding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerIntranetIP6BindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerIntranetIP6Binding `json:"vpnvserver_intranetip6_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_intranetip6_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServerIntranetIP6Binding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerIntranetIP6BindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerIntranetIP6Binding `json:"vpnvserver_intranetip6_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_intranetip_binding
func (s *VPNService) AddVPNVServerIntranetIPBinding(resource models.VPNVServerIntranetIPBinding) error {
	payload := map[string]any{"vpnvserver_intranetip_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerIntranetIPBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerIntranetIPBinding(name string, intranetip string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=intranetip:%s", vpnVServerIntranetIPBindingURL, name, intranetip), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerIntranetIPBinding() ([]models.VPNVServerIntranetIPBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerIntranetIPBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerIntranetIPBinding `json:"vpnvserver_intranetip_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerIntranetIPBinding(name string) (*models.VPNVServerIntranetIPBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerIntranetIPBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerIntranetIPBinding `json:"vpnvserver_intranetip_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_intranetip_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServerIntranetIPBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerIntranetIPBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerIntranetIPBinding `json:"vpnvserver_intranetip_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_responderpolicy_binding
func (s *VPNService) AddVPNVServerResponderPolicyBinding(resource models.VPNVServerResponderPolicyBinding) error {
	payload := map[string]any{"vpnvserver_responderpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerResponderPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerResponderPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", vpnVServerResponderPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerResponderPolicyBinding() ([]models.VPNVServerResponderPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerResponderPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerResponderPolicyBinding `json:"vpnvserver_responderpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerResponderPolicyBinding(name string) (*models.VPNVServerResponderPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerResponderPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerResponderPolicyBinding `json:"vpnvserver_responderpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_responderpolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServerResponderPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerResponderPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerResponderPolicyBinding `json:"vpnvserver_responderpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_rewritepolicy_binding
func (s *VPNService) AddVPNVServerRewritePolicyBinding(resource models.VPNVServerRewritePolicyBinding) error {
	payload := map[string]any{"vpnvserver_rewritepolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerRewritePolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerRewritePolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", vpnVServerRewritePolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerRewritePolicyBinding() ([]models.VPNVServerRewritePolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerRewritePolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerRewritePolicyBinding `json:"vpnvserver_rewritepolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerRewritePolicyBinding(name string) (*models.VPNVServerRewritePolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerRewritePolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerRewritePolicyBinding `json:"vpnvserver_rewritepolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_rewritepolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServerRewritePolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerRewritePolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerRewritePolicyBinding `json:"vpnvserver_rewritepolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_sharefileserver_binding
func (s *VPNService) AddVPNVServerShareFileServerBinding(resource models.VPNVServerShareFileServerBinding) error {
	payload := map[string]any{"vpnvserver_sharefileserver_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerShareFileServerBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerShareFileServerBinding(name string, sharefile string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=sharefile:%s", vpnVServerShareFileServerBindingURL, name, sharefile), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerShareFileServerBinding() ([]models.VPNVServerShareFileServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerShareFileServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerShareFileServerBinding `json:"vpnvserver_sharefileserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerShareFileServerBinding(name string) (*models.VPNVServerShareFileServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerShareFileServerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerShareFileServerBinding `json:"vpnvserver_sharefileserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_sharefileserver_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServerShareFileServerBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerShareFileServerBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerShareFileServerBinding `json:"vpnvserver_sharefileserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_staserver_binding
func (s *VPNService) AddVPNVServerSTAServerBinding(resource models.VPNVServerSTAServerBinding) error {
	payload := map[string]any{"vpnvserver_staserver_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerSTAServerBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerSTAServerBinding(name string, staserver string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=staserver:%s", vpnVServerSTAServerBindingURL, name, staserver), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerSTAServerBinding() ([]models.VPNVServerSTAServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerSTAServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerSTAServerBinding `json:"vpnvserver_staserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerSTAServerBinding(name string) (*models.VPNVServerSTAServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerSTAServerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerSTAServerBinding `json:"vpnvserver_staserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_staserver_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServerSTAServerBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerSTAServerBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerSTAServerBinding `json:"vpnvserver_staserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_vpnclientlessaccesspolicy_binding
func (s *VPNService) AddVPNVServerVPNClientlessAccessPolicyBinding(resource models.VPNVServerVPNClientlessAccessPolicyBinding) error {
	payload := map[string]any{"vpnvserver_vpnclientlessaccesspolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerVPNClientlessAccessPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerVPNClientlessAccessPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", vpnVServerVPNClientlessAccessPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerVPNClientlessAccessPolicyBinding() ([]models.VPNVServerVPNClientlessAccessPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerVPNClientlessAccessPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerVPNClientlessAccessPolicyBinding `json:"vpnvserver_vpnclientlessaccesspolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerVPNClientlessAccessPolicyBinding(name string) (*models.VPNVServerVPNClientlessAccessPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerVPNClientlessAccessPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerVPNClientlessAccessPolicyBinding `json:"vpnvserver_vpnclientlessaccesspolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_vpnclientlessaccesspolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServerVPNClientlessAccessPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerVPNClientlessAccessPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerVPNClientlessAccessPolicyBinding `json:"vpnvserver_vpnclientlessaccesspolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_vpnepaprofile_binding
func (s *VPNService) AddVPNVServerVPNEPAProfileBinding(resource models.VPNVServerVPNEPAProfileBinding) error {
	payload := map[string]any{"vpnvserver_vpnepaprofile_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerVPNEPAProfileBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerVPNEPAProfileBinding(name string, epaprofile string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=epaprofile:%s", vpnVServerVPNEPAProfileBindingURL, name, epaprofile), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerVPNEPAProfileBinding() ([]models.VPNVServerVPNEPAProfileBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerVPNEPAProfileBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerVPNEPAProfileBinding `json:"vpnvserver_vpnepaprofile_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerVPNEPAProfileBinding(name string) (*models.VPNVServerVPNEPAProfileBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerVPNEPAProfileBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerVPNEPAProfileBinding `json:"vpnvserver_vpnepaprofile_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_vpnepaprofile_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServerVPNEPAProfileBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerVPNEPAProfileBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerVPNEPAProfileBinding `json:"vpnvserver_vpnepaprofile_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_vpneula_binding
func (s *VPNService) AddVPNVServerVPNEULABinding(resource models.VPNVServerVPNEULABinding) error {
	payload := map[string]any{"vpnvserver_vpneula_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerVPNEULABindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerVPNEULABinding(name string, eula string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=eula:%s", vpnVServerVPNEULABindingURL, name, eula), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerVPNEULABinding() ([]models.VPNVServerVPNEULABinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerVPNEULABindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerVPNEULABinding `json:"vpnvserver_vpneula_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerVPNEULABinding(name string) (*models.VPNVServerVPNEULABinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerVPNEULABindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerVPNEULABinding `json:"vpnvserver_vpneula_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_vpneula_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServerVPNEULABinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerVPNEULABindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerVPNEULABinding `json:"vpnvserver_vpneula_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_vpnintranetapplication_binding
func (s *VPNService) AddVPNVServerVPNIntranetApplicationBinding(resource models.VPNVServerVPNIntranetApplicationBinding) error {
	payload := map[string]any{"vpnvserver_vpnintranetapplication_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerVPNIntranetApplicationBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerVPNIntranetApplicationBinding(name string, intranetapplication string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=intranetapplication:%s", vpnVServerVPNIntranetApplicationBindingURL, name, intranetapplication), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerVPNIntranetApplicationBinding() ([]models.VPNVServerVPNIntranetApplicationBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerVPNIntranetApplicationBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerVPNIntranetApplicationBinding `json:"vpnvserver_vpnintranetapplication_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerVPNIntranetApplicationBinding(name string) (*models.VPNVServerVPNIntranetApplicationBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerVPNIntranetApplicationBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerVPNIntranetApplicationBinding `json:"vpnvserver_vpnintranetapplication_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_vpnintranetapplication_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServerVPNIntranetApplicationBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerVPNIntranetApplicationBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerVPNIntranetApplicationBinding `json:"vpnvserver_vpnintranetapplication_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_vpnnexthopserver_binding
func (s *VPNService) AddVPNVServerVPNNextHopServerBinding(resource models.VPNVServerVPNNextHopServerBinding) error {
	payload := map[string]any{"vpnvserver_vpnnexthopserver_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerVPNNextHopServerBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerVPNNextHopServerBinding(name string, nexthopserver string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=nexthopserver:%s", vpnVServerVPNNextHopServerBindingURL, name, nexthopserver), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerVPNNextHopServerBinding() ([]models.VPNVServerVPNNextHopServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerVPNNextHopServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerVPNNextHopServerBinding `json:"vpnvserver_vpnnexthopserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerVPNNextHopServerBinding(name string) (*models.VPNVServerVPNNextHopServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerVPNNextHopServerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerVPNNextHopServerBinding `json:"vpnvserver_vpnnexthopserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_vpnnexthopserver_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServerVPNNextHopServerBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerVPNNextHopServerBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerVPNNextHopServerBinding `json:"vpnvserver_vpnnexthopserver_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_vpnportaltheme_binding
func (s *VPNService) AddVPNVServerVPNPortalThemeBinding(resource models.VPNVServerVPNPortalThemeBinding) error {
	payload := map[string]any{"vpnvserver_vpnportaltheme_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerVPNPortalThemeBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerVPNPortalThemeBinding(name string, portaltheme string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=portaltheme:%s", vpnVServerVPNPortalThemeBindingURL, name, portaltheme), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerVPNPortalThemeBinding() ([]models.VPNVServerVPNPortalThemeBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerVPNPortalThemeBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerVPNPortalThemeBinding `json:"vpnvserver_vpnportaltheme_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerVPNPortalThemeBinding(name string) (*models.VPNVServerVPNPortalThemeBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerVPNPortalThemeBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerVPNPortalThemeBinding `json:"vpnvserver_vpnportaltheme_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_vpnportaltheme_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServerVPNPortalThemeBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerVPNPortalThemeBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerVPNPortalThemeBinding `json:"vpnvserver_vpnportaltheme_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_vpnsessionpolicy_binding
func (s *VPNService) AddVPNVServerVPNSessionPolicyBinding(resource models.VPNVServerVPNSessionPolicyBinding) error {
	payload := map[string]any{"vpnvserver_vpnsessionpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerVPNSessionPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerVPNSessionPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", vpnVServerVPNSessionPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerVPNSessionPolicyBinding() ([]models.VPNVServerVPNSessionPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerVPNSessionPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerVPNSessionPolicyBinding `json:"vpnvserver_vpnsessionpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerVPNSessionPolicyBinding(name string) (*models.VPNVServerVPNSessionPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerVPNSessionPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerVPNSessionPolicyBinding `json:"vpnvserver_vpnsessionpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_vpnsessionpolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServerVPNSessionPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerVPNSessionPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerVPNSessionPolicyBinding `json:"vpnvserver_vpnsessionpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_vpntrafficpolicy_binding
func (s *VPNService) AddVPNVServerVPNTrafficPolicyBinding(resource models.VPNVServerVPNTrafficPolicyBinding) error {
	payload := map[string]any{"vpnvserver_vpntrafficpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerVPNTrafficPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerVPNTrafficPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", vpnVServerVPNTrafficPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerVPNTrafficPolicyBinding() ([]models.VPNVServerVPNTrafficPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerVPNTrafficPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerVPNTrafficPolicyBinding `json:"vpnvserver_vpntrafficpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerVPNTrafficPolicyBinding(name string) (*models.VPNVServerVPNTrafficPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerVPNTrafficPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerVPNTrafficPolicyBinding `json:"vpnvserver_vpntrafficpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_vpntrafficpolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServerVPNTrafficPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerVPNTrafficPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerVPNTrafficPolicyBinding `json:"vpnvserver_vpntrafficpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_vpnurlpolicy_binding
func (s *VPNService) AddVPNVServerVPNURLPolicyBinding(resource models.VPNVServerVPNURLPolicyBinding) error {
	payload := map[string]any{"vpnvserver_vpnurlpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerVPNURLPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerVPNURLPolicyBinding(name string, policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=policy:%s", vpnVServerVPNURLPolicyBindingURL, name, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerVPNURLPolicyBinding() ([]models.VPNVServerVPNURLPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerVPNURLPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerVPNURLPolicyBinding `json:"vpnvserver_vpnurlpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerVPNURLPolicyBinding(name string) (*models.VPNVServerVPNURLPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerVPNURLPolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerVPNURLPolicyBinding `json:"vpnvserver_vpnurlpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_vpnurlpolicy_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServerVPNURLPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerVPNURLPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerVPNURLPolicyBinding `json:"vpnvserver_vpnurlpolicy_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}

// vpnvserver_vpnurl_binding
func (s *VPNService) AddVPNVServerVPNURLBinding(resource models.VPNVServerVPNURLBinding) error {
	payload := map[string]any{"vpnvserver_vpnurl_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vpnVServerVPNURLBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) DeleteVPNVServerVPNURLBinding(name string, urlname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s?args=urlname:%s", vpnVServerVPNURLBindingURL, name, urlname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *VPNService) GetAllVPNVServerVPNURLBinding() ([]models.VPNVServerVPNURLBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vpnVServerVPNURLBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerVPNURLBinding `json:"vpnvserver_vpnurl_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return result.Data, nil
}

func (s *VPNService) GetVPNVServerVPNURLBinding(name string) (*models.VPNVServerVPNURLBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vpnVServerVPNURLBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VPNVServerVPNURLBinding `json:"vpnvserver_vpnurl_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("vpnvserver_vpnurl_binding %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *VPNService) CountVPNVServerVPNURLBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vpnVServerVPNURLBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VPNVServerVPNURLBinding `json:"vpnvserver_vpnurl_binding"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return 0, err
	}

	if len(result.Data) == 0 {
		return 0, nil
	}

	return int(result.Data[0].Count), nil
}
