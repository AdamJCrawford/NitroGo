package nitrogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/AdamJCrawford/NitroGo/nitrogo/models"
)

const (
	aaaCertParamsURL                                 = "/nitro/v1/config/aaacertparams"
	aaaGlobalAAAPreauthenticationPolicyBindingURL    = "/nitro/v1/config/aaaglobal_aaapreauthenticationpolicy_binding"
	aaaGlobalAuthenticationNegotiateActionBindingURL = "/nitro/v1/config/aaaglobal_authenticationnegotiateaction_binding"
	aaaGlobalBindingURL                              = "/nitro/v1/config/aaaglobal_binding"
	aaaGroupURL                                      = "/nitro/v1/config/aaagroup"
	aaaGroupAAAUserBindingURL                        = "/nitro/v1/config/aaagroup_aaauser_binding"
	aaaGroupAuditnslogPolicyBindingURL               = "/nitro/v1/config/aaagroup_auditnslogpolicy_binding"
	aaaGroupAuditsyslogPolicyBindingURL              = "/nitro/v1/config/aaagroup_auditsyslogpolicy_binding"
	aaaGroupAuthorizationPolicyBindingURL            = "/nitro/v1/config/aaagroup_authorizationpolicy_binding"
	aaaGroupBindingURL                               = "/nitro/v1/config/aaagroup_binding"
	aaaGroupIntranetIP6BindingURL                    = "/nitro/v1/config/aaagroup_intranetip6_binding"
	aaaGroupIntranetIPBindingURL                     = "/nitro/v1/config/aaagroup_intranetip_binding"
	aaaGroupTMSessionPolicyBindingURL                = "/nitro/v1/config/aaagroup_tmsessionpolicy_binding"
	aaaGroupVPNIntranetApplicationBindingURL         = "/nitro/v1/config/aaagroup_vpnintranetapplication_binding"
	aaaGroupVPNSessionPolicyBindingURL               = "/nitro/v1/config/aaagroup_vpnsessionpolicy_binding"
	aaaGroupVPNTrafficPolicyBindingURL               = "/nitro/v1/config/aaagroup_vpntrafficpolicy_binding "
	aaaGroupVPNURLPolicyBindingURL                   = "/nitro/v1/config/aaagroup_vpnurlpolicy_binding"
	aaaGroupVPNURLBindingURL                         = "/nitro/v1/config/aaagroup_vpnurl_binding"
	aaaKCDAccountURL                                 = "/nitro/v1/config/aaakcdaccount"
	aaaLDAPParamsURL                                 = "/nitro/v1/config/aaaldapparams"
	aaaOTPParameterURL                               = "/nitro/v1/config/aaaotpparameter"
	aaaParameterURL                                  = "/nitro/v1/config/aaaparameter"
	aaapreauthenticationActionURL                    = "/nitro/v1/config/aaapreauthenticationaction"
	aaaPreauthenticationParameterURL                 = "/nitro/v1/config/aaapreauthenticationparameter"
	aaaPreauthenticationPolicyURL                    = "/nitro/v1/config/aaapreauthenticationpolicy"
	aaaPreauthenticationPolicyAAAGlobalBindingURL    = "/nitro/v1/config/aaapreauthenticationpolicy_aaaglobal_binding"
	aaaPreauthenticationPolicyBindingURL             = "/nitro/v1/config/aaapreauthenticationpolicy_binding"
	aaaPreauthenticationPolicyVPNVServerBindingURL   = "/nitro/v1/config/aaapreauthenticationpolicy_vpnvserver_binding"
	aaaRADIUSParamsURL                               = "/nitro/v1/config/aaaradiusparams"
	aaaSessionURL                                    = "/nitro/v1/config/aaasession"
	aaaSSOProfileURL                                 = "/nitro/v1/config/aaassoprofile"
	aaaTACACSParamsURL                               = "/nitro/v1/config/aaatacacsparams"
	aaaUserURL                                       = "/nitro/v1/config/aaauser"
	aaaUserAAAGroupBindingURL                        = "/nitro/v1/config/aaauser_aaagroup_binding"
	aaaUserAuditnslogPolicyBindingURL                = "/nitro/v1/config/aaauser_auditnslogpolicy_binding"
	aaaUserAuditsyslogPolicyBindingURL               = "/nitro/v1/config/aaauser_auditsyslogpolicy_binding"
	aaaUserAuthorizationPolicyBindingURL             = "/nitro/v1/config/aaauser_authorizationpolicy_binding"
	aaaUserBindingURL                                = "/nitro/v1/config/aaauser_binding"
	aaaUserIntranetIP6BindingURL                     = "/nitro/v1/config/aaauser_intranetip6_binding"
	aaaUserIntranetIPBindingURL                      = "/nitro/v1/config/aaauser_intranetip_binding"
	aaaUserTMSessionPolicyBindingURL                 = "/nitro/v1/config/aaauser_tmsessionpolicy_binding"
	aaaUserVPNIntranetApplicationBindingURL          = "/nitro/v1/config/aaauser_vpnintranetapplication_binding"
	aaaUserVPNSessionPolicyBindingURL                = "/nitro/v1/config/aaauser_vpnsessionpolicy_binding"
	aaaUserVPNTrafficPolicyBindingURL                = "/nitro/v1/config/aaauser_vpntrafficpolicy_binding"
	aaaUserVPNURLPolicyBindingURL                    = "/nitro/v1/config/aaauser_vpnurlpolicy_binding"
	aaaUserVPNURLBindingURL                          = "/nitro/v1/config/aaauser_vpnurl_binding"
)

// Authentication, authorization, and accounting service configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaa
type AAAService struct {
	client *Client
}

// aaacertparams
// Configuration for certificate parameter resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaacertparams
func (s *AAAService) UpdateAAACertParams(resource models.AAACertParams) error {
	payload := map[string]any{"aaacertparams": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, aaaCertParamsURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) UnsetAAACertParams(resource models.AAACertParams) error {
	payload := map[string]any{"aaacertparams": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", aaaCertParamsURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) GetAllAAACertParams() (models.AAACertParams, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaCertParamsURL, nil)
	if err != nil {
		return models.AAACertParams{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.AAACertParams{}, err
	}

	var result struct {
		Data models.AAACertParams `json:"aaacertparams"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.AAACertParams{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// aaaglobal_aaapreauthenticationpolicy_binding
// Binding object showing the aaapreauthenticationpolicy that can be bound to aaaglobal.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaaglobal_aaapreauthenticationpolicy_binding
func (s *AAAService) AddAAAGlobalAAAPreauthenticationPolicyBinding(resource models.AAAGlobalAAAPreAuthenticationPolicyBinding) error {
	payload := map[string]any{"aaaglobal_aaapreauthenticationpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, aaaGlobalAAAPreauthenticationPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) DeleteAAAGlobalAAAPreauthenticationPolicyBinding(policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", aaaGlobalAAAPreauthenticationPolicyBindingURL, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) GetAAAGlobalAAAPreauthenticationPolicyBinding() ([]models.AAAGlobalAAAPreAuthenticationPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaGlobalAAAPreauthenticationPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAGlobalAAAPreAuthenticationPolicyBinding `json:"aaaglobal_aaapreauthenticationpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) CountAAAGlobalAAAPreauthenticationPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", aaaGlobalAAAPreauthenticationPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AAAGlobalAAAPreAuthenticationPolicyBinding `json:"aaaglobal_aaapreauthenticationpolicy_binding"`
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

// aaaglobal_authenticationnegotiataction_binding
// Binding object showing the authenticationnegotiateaction that can be bound to aaaglobal.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaaglobal_authenticationnegotiateaction_binding
func (s *AAAService) AddAAAGlobalAuthenticationNegotiateActionBinding(resource models.AAAGlobalAuthenticationNegotiateActionBinding) error {
	payload := map[string]any{"aaaglobal_authenticationnegotiateaction_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, aaaGlobalAuthenticationNegotiateActionBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) DeleteAAAGlobalAuthenticationNegotiateActionBinding(policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", aaaGlobalAuthenticationNegotiateActionBindingURL, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) GetAAAGlobalAuthenticationNegotiateActionBinding() ([]models.AAAGlobalAuthenticationNegotiateActionBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaGlobalAuthenticationNegotiateActionBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAGlobalAuthenticationNegotiateActionBinding `json:"aaaglobal_authenticationnegotiateaction_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) CountAAAGlobalAuthenticationNegotiateActionBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", aaaGlobalAuthenticationNegotiateActionBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AAAGlobalAuthenticationNegotiateActionBinding `json:"aaaglobal_authenticationnegotiateaction_binding"`
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

// aaaglobal_binding
// Binding object which returns the resources bound to aaaglobal.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaaglobal_binding
func (s *AAAService) GetAAAGlobalBinding() (models.AAAGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaGlobalBindingURL, nil)
	if err != nil {
		return models.AAAGlobalBinding{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.AAAGlobalBinding{}, err
	}

	var result struct {
		Data []models.AAAGlobalBinding `json:"aaaglobal_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.AAAGlobalBinding{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.AAAGlobalBinding{}, fmt.Errorf("aaaglobal_binding not found")
	}

	return result.Data[0], nil
}

// aaagroup
// Configuration for AAA group resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaagroup
func (s *AAAService) AddAAAGroup(resource models.AAAGroup) error {
	payload := map[string]any{"aaagroup": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, aaaGroupURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) DeleteAAAGroup(groupname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", aaaGroupURL, groupname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) GetAllAAAGroup() ([]models.AAAGroup, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaGroupURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAGroup `json:"aaagroup"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) GetAAAGroup(groupname string) (models.AAAGroup, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", aaaGroupURL, groupname), nil)
	if err != nil {
		return models.AAAGroup{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.AAAGroup{}, err
	}

	var result struct {
		Data []models.AAAGroup `json:"aaagroup"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.AAAGroup{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.AAAGroup{}, fmt.Errorf("aaagroup %s not found", groupname)
	}

	return result.Data[0], nil
}

func (s *AAAService) CountAAAGroup() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", aaaGroupURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AAAGroup `json:"aaagroup"`
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

// aaagroup_aaauser_binding
// Binding object showing the aaauser that can be bound to aaagroup.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaagroup_aaauser_binding
func (s *AAAService) AddAAAGroupAAAUserBinding(resource models.AAAGroupAAAUserBinding) error {
	payload := map[string]any{"aaagroup_aaauser_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, aaaGroupAAAUserBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) DeleteAAAGroupAAAUserBinding(groupname string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", aaaGroupAAAUserBindingURL, groupname, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) GetAllAAAGroupAAAUserBinding() ([]models.AAAGroupAAAUserBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaGroupAAAUserBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAGroupAAAUserBinding `json:"aaagroup_aaauser_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) GetAAAGroupAAAUserBinding(groupname string) ([]models.AAAGroupAAAUserBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", aaaGroupAAAUserBindingURL, groupname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAGroupAAAUserBinding `json:"aaagroup_aaauser_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) CountAAAGroupAAAUserBinding(groupname string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", aaaGroupAAAUserBindingURL, groupname), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AAAGroupAAAUserBinding `json:"aaagroup_aaauser_binding"`
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

// aaagroup_auditnslogpolicy_binding
// Binding object showing the auditnslogpolicy that can be bound to aaagroup.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaagroup_auditnslogpolicy_binding
func (s *AAAService) AddAAAGroupAuditNSLogPolicyBinding(resource models.AAAGroupAuditNSLogPolicyBinding) error {
	payload := map[string]any{"aaagroup_auditnslogpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, aaaGroupAuditnslogPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) DeleteAAAGroupAuditNSLogPolicyBinding(groupname string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", aaaGroupAuditnslogPolicyBindingURL, groupname, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) GetAllAAAGroupAuditNSLogPolicyBinding() ([]models.AAAGroupAuditNSLogPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaGroupAuditnslogPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAGroupAuditNSLogPolicyBinding `json:"aaagroup_auditnslogpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) GetAAAGroupAuditNSLogPolicyBinding(groupname string) ([]models.AAAGroupAuditNSLogPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", aaaGroupAuditnslogPolicyBindingURL, groupname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAGroupAuditNSLogPolicyBinding `json:"aaagroup_auditnslogpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) CountAAAGroupAuditNSLogPolicyBinding(groupname string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", aaaGroupAuditnslogPolicyBindingURL, groupname), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AAAGroupAuditNSLogPolicyBinding `json:"aaagroup_auditnslogpolicy_binding"`
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

// aaagroup_auditsyslogpolicy_binding
// Binding object showing the auditsyslogpolicy that can be bound to aaagroup.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaagroup_auditsyslogpolicy_binding
func (s *AAAService) AddAAAGroupAuditSyslogPolicyBinding(resource models.AAAGroupAuditSyslogPolicyBinding) error {
	payload := map[string]any{"aaagroup_auditsyslogpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, aaaGroupAuditsyslogPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) DeleteAAAGroupAuditSyslogPolicyBinding(groupname string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", aaaGroupAuditsyslogPolicyBindingURL, groupname, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) GetAllAAAGroupAuditSyslogPolicyBinding() ([]models.AAAGroupAuditSyslogPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaGroupAuditsyslogPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAGroupAuditSyslogPolicyBinding `json:"aaagroup_auditsyslogpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) GetAAAGroupAuditSyslogPolicyBinding(groupname string) ([]models.AAAGroupAuditSyslogPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", aaaGroupAuditsyslogPolicyBindingURL, groupname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAGroupAuditSyslogPolicyBinding `json:"aaagroup_auditsyslogpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) CountAAAGroupAuditSyslogPolicyBinding(groupname string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", aaaGroupAuditsyslogPolicyBindingURL, groupname), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AAAGroupAuditSyslogPolicyBinding `json:"aaagroup_auditsyslogpolicy_binding"`
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

// aaagroup_authorizationpolicy_binding
// Binding object showing the authorizationpolicy that can be bound to aaagroup.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaagroup_authorizationpolicy_binding
func (s *AAAService) AddAAAGroupAuthorizationPolicyBinding(resource models.AAAGroupAuthorizationPolicyBinding) error {
	payload := map[string]any{"aaagroup_authorizationpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, aaaGroupAuthorizationPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) DeleteAAAGroupAuthorizationPolicyBinding(groupname string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", aaaGroupAuthorizationPolicyBindingURL, groupname, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) GetAllAAAGroupAuthorizationPolicyBinding() ([]models.AAAGroupAuthorizationPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaGroupAuthorizationPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAGroupAuthorizationPolicyBinding `json:"aaagroup_authorizationpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) GetAAAGroupAuthorizationPolicyBinding(groupname string) ([]models.AAAGroupAuthorizationPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", aaaGroupAuthorizationPolicyBindingURL, groupname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAGroupAuthorizationPolicyBinding `json:"aaagroup_authorizationpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) CountAAAGroupAuthorizationPolicyBinding(groupname string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", aaaGroupAuthorizationPolicyBindingURL, groupname), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AAAGroupAuthorizationPolicyBinding `json:"aaagroup_authorizationpolicy_binding"`
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

// aaagroup_binding
// Binding object which returns the resources bound to aaagroup.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaagroup_binding
func (s *AAAService) GetAllAAAGroupBinding() ([]models.AAAGroupBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaGroupBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAGroupBinding `json:"aaagroup_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) GetAAAGroupBinding(groupname string) (models.AAAGroupBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", aaaGroupBindingURL, groupname), nil)
	if err != nil {
		return models.AAAGroupBinding{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.AAAGroupBinding{}, err
	}

	var result struct {
		Data []models.AAAGroupBinding `json:"aaagroup_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.AAAGroupBinding{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.AAAGroupBinding{}, fmt.Errorf("aaagroup_binding %s not found", groupname)
	}

	return result.Data[0], nil
}

// aaagroup_intranetip6_binding
// Binding object showing the intranetip6 that can be bound to aaagroup.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaagroup_intranetip6_binding
func (s *AAAService) AddAAAGroupIntranetIP6Binding(resource models.AAAGroupIntranetIP6Binding) error {
	payload := map[string]any{"aaagroup_intranetip6_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, aaaGroupIntranetIP6BindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) DeleteAAAGroupIntranetIP6Binding(groupname string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", aaaGroupIntranetIP6BindingURL, groupname, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) GetAllAAAGroupIntranetIP6Binding() ([]models.AAAGroupIntranetIP6Binding, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaGroupIntranetIP6BindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAGroupIntranetIP6Binding `json:"aaagroup_intranetip6_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) GetAAAGroupIntranetIP6Binding(groupname string) ([]models.AAAGroupIntranetIP6Binding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", aaaGroupIntranetIP6BindingURL, groupname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAGroupIntranetIP6Binding `json:"aaagroup_intranetip6_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) CountAAAGroupIntranetIP6Binding(groupname string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", aaaGroupIntranetIP6BindingURL, groupname), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AAAGroupIntranetIP6Binding `json:"aaagroup_intranetip6_binding"`
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

// aaagroup_intranetip_binding
// Binding object showing the intranetip that can be bound to aaagroup.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaagroup_intranetip_binding
func (s *AAAService) AddAAAGroupIntranetIPBinding(resource models.AAAGroupIntranetIPBinding) error {
	payload := map[string]any{"aaagroup_intranetip_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, aaaGroupIntranetIPBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) DeleteAAAGroupIntranetIPBinding(groupname string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", aaaGroupIntranetIPBindingURL, groupname, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) GetAllAAAGroupIntranetIPBinding() ([]models.AAAGroupIntranetIPBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaGroupIntranetIPBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAGroupIntranetIPBinding `json:"aaagroup_intranetip_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) GetAAAGroupIntranetIPBinding(groupname string) ([]models.AAAGroupIntranetIPBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", aaaGroupIntranetIPBindingURL, groupname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAGroupIntranetIPBinding `json:"aaagroup_intranetip_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) CountAAAGroupIntranetIPBinding(groupname string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", aaaGroupIntranetIPBindingURL, groupname), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AAAGroupIntranetIPBinding `json:"aaagroup_intranetip_binding"`
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

// aaagroup_tmsessionpolicy_binding
// Binding object showing the tmsessionpolicy that can be bound to aaagroup.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaagroup_tmsessionpolicy_binding
func (s *AAAService) AddAAAGroupTMSessionPolicyBinding(resource models.AAAGroupTMSessionPolicyBinding) error {
	payload := map[string]any{"aaagroup_tmsessionpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, aaaGroupTMSessionPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) DeleteAAAGroupTMSessionPolicyBinding(groupname string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", aaaGroupTMSessionPolicyBindingURL, groupname, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) GetAllAAAGroupTMSessionPolicyBinding() ([]models.AAAGroupTMSessionPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaGroupTMSessionPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAGroupTMSessionPolicyBinding `json:"aaagroup_tmsessionpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) GetAAAGroupTMSessionPolicyBinding(groupname string) ([]models.AAAGroupTMSessionPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", aaaGroupTMSessionPolicyBindingURL, groupname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAGroupTMSessionPolicyBinding `json:"aaagroup_tmsessionpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) CountAAAGroupTMSessionPolicyBinding(groupname string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", aaaGroupTMSessionPolicyBindingURL, groupname), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AAAGroupTMSessionPolicyBinding `json:"aaagroup_tmsessionpolicy_binding"`
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

// aaagroup_vpnintranetapplication_binding
// Binding object showing the vpnintranetapplication that can be bound to aaagroup.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaagroup_vpnintranetapplication_binding
func (s *AAAService) AddAAAGroupVPNIntranetApplicationBinding(resource models.AAAGroupVPNIntranetApplicationBinding) error {
	payload := map[string]any{"aaagroup_vpnintranetapplication_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, aaaGroupVPNIntranetApplicationBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) DeleteAAAGroupVPNIntranetApplicationBinding(groupname string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", aaaGroupVPNIntranetApplicationBindingURL, groupname, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) GetAllAAAGroupVPNIntranetApplicationBinding() ([]models.AAAGroupVPNIntranetApplicationBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaGroupVPNIntranetApplicationBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAGroupVPNIntranetApplicationBinding `json:"aaagroup_vpnintranetapplication_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) GetAAAGroupVPNIntranetApplicationBinding(groupname string) ([]models.AAAGroupVPNIntranetApplicationBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", aaaGroupVPNIntranetApplicationBindingURL, groupname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAGroupVPNIntranetApplicationBinding `json:"aaagroup_vpnintranetapplication_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) CountAAAGroupVPNIntranetApplicationBinding(groupname string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", aaaGroupVPNIntranetApplicationBindingURL, groupname), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AAAGroupVPNIntranetApplicationBinding `json:"aaagroup_vpnintranetapplication_binding"`
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

// aaagroup_vpnsessionpolicy_binding
// Binding object showing the vpnsessionpolicy that can be bound to aaagroup.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaagroup_vpnsessionpolicy_binding
func (s *AAAService) AddAAAGroupVPNSessionPolicyBinding(resource models.AAAGroupVPNSessionPolicyBinding) error {
	payload := map[string]any{"aaagroup_vpnsessionpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, aaaGroupVPNSessionPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) DeleteAAAGroupVPNSessionPolicyBinding(groupname string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", aaaGroupVPNSessionPolicyBindingURL, groupname, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) GetAllAAAGroupVPNSessionPolicyBinding() ([]models.AAAGroupVPNSessionPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaGroupVPNSessionPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAGroupVPNSessionPolicyBinding `json:"aaagroup_vpnsessionpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) GetAAAGroupVPNSessionPolicyBinding(groupname string) ([]models.AAAGroupVPNSessionPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", aaaGroupVPNSessionPolicyBindingURL, groupname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAGroupVPNSessionPolicyBinding `json:"aaagroup_vpnsessionpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) CountAAAGroupVPNSessionPolicyBinding(groupname string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", aaaGroupVPNSessionPolicyBindingURL, groupname), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AAAGroupVPNSessionPolicyBinding `json:"aaagroup_vpnsessionpolicy_binding"`
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

// aaagroup_vpntrafficpolicy_binding
// Binding object showing the vpntrafficpolicy that can be bound to aaagroup.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaagroup_vpntrafficpolicy_binding
func (s *AAAService) AddAAAGroupVPNTrafficPolicyBinding(resource models.AAAGroupVPNTrafficPolicyBinding) error {
	payload := map[string]any{"aaagroup_vpntrafficpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, aaaGroupVPNTrafficPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) DeleteAAAGroupVPNTrafficPolicyBinding(groupname string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", aaaGroupVPNTrafficPolicyBindingURL, groupname, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) GetAllAAAGroupVPNTrafficPolicyBinding() ([]models.AAAGroupVPNTrafficPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaGroupVPNTrafficPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAGroupVPNTrafficPolicyBinding `json:"aaagroup_vpntrafficpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) GetAAAGroupVPNTrafficPolicyBinding(groupname string) ([]models.AAAGroupVPNTrafficPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", aaaGroupVPNTrafficPolicyBindingURL, groupname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAGroupVPNTrafficPolicyBinding `json:"aaagroup_vpntrafficpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) CountAAAGroupVPNTrafficPolicyBinding(groupname string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", aaaGroupVPNTrafficPolicyBindingURL, groupname), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AAAGroupVPNTrafficPolicyBinding `json:"aaagroup_vpntrafficpolicy_binding"`
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

// aaagroup_vpnurlpolicy_binding
// Binding object showing the vpnurlpolicy that can be bound to aaagroup.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaagroup_vpnurlpolicy_binding
func (s *AAAService) AddAAAGroupVPNURLPolicyBinding(resource models.AAAGroupVPNURLPolicyBinding) error {
	payload := map[string]any{"aaagroup_vpnurlpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, aaaGroupVPNURLPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) DeleteAAAGroupVPNURLPolicyBinding(groupname string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", aaaGroupVPNURLPolicyBindingURL, groupname, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) GetAllAAAGroupVPNURLPolicyBinding() ([]models.AAAGroupVPNURLPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaGroupVPNURLPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAGroupVPNURLPolicyBinding `json:"aaagroup_vpnurlpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) GetAAAGroupVPNURLPolicyBinding(groupname string) ([]models.AAAGroupVPNURLPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", aaaGroupVPNURLPolicyBindingURL, groupname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAGroupVPNURLPolicyBinding `json:"aaagroup_vpnurlpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) CountAAAGroupVPNURLPolicyBinding(groupname string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", aaaGroupVPNURLPolicyBindingURL, groupname), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AAAGroupVPNURLPolicyBinding `json:"aaagroup_vpnurlpolicy_binding"`
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

// aaagroup_vpnurl_binding
// Binding object showing the vpnurl that can be bound to aaagroup.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaagroup_vpnurl_binding
func (s *AAAService) AddAAAGroupVPNURLBinding(resource models.AAAGroupVPNURLBinding) error {
	payload := map[string]any{"aaagroup_vpnurl_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, aaaGroupVPNURLBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) DeleteAAAGroupVPNURLBinding(groupname string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", aaaGroupVPNURLBindingURL, groupname, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) GetAllAAAGroupVPNURLBinding() ([]models.AAAGroupVPNURLBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaGroupVPNURLBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAGroupVPNURLBinding `json:"aaagroup_vpnurl_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) GetAAAGroupVPNURLBinding(groupname string) ([]models.AAAGroupVPNURLBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", aaaGroupVPNURLBindingURL, groupname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAGroupVPNURLBinding `json:"aaagroup_vpnurl_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) CountAAAGroupVPNURLBinding(groupname string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", aaaGroupVPNURLBindingURL, groupname), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AAAGroupVPNURLBinding `json:"aaagroup_vpnurl_binding"`
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

// aaakcdaccount
// Configuration for Kerberos constrained delegation account resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaakcdaccount
func (s *AAAService) AddAAAKCDAccount(resource models.AAAKCDAccount) error {
	payload := map[string]any{"aaakcdaccount": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, aaaKCDAccountURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) DeleteAAAKCDAccount(kcdaccount string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", aaaKCDAccountURL, kcdaccount), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) UpdateAAAKCDAccount(resource models.AAAKCDAccount) error {
	payload := map[string]any{"aaakcdaccount": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, aaaKCDAccountURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) UnsetAAAKCDAccount(resource models.AAAKCDAccount) error {
	payload := map[string]any{"aaakcdaccount": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", aaaKCDAccountURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) GetAllAAAKCDAccount() ([]models.AAAKCDAccount, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaKCDAccountURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAKCDAccount `json:"aaakcdaccount"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) GetAAAKCDAccount(kcdaccount string) (models.AAAKCDAccount, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", aaaKCDAccountURL, kcdaccount), nil)
	if err != nil {
		return models.AAAKCDAccount{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.AAAKCDAccount{}, err
	}

	var result struct {
		Data []models.AAAKCDAccount `json:"aaakcdaccount"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.AAAKCDAccount{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.AAAKCDAccount{}, fmt.Errorf("aaakcdaccount %s not found", kcdaccount)
	}

	return result.Data[0], nil
}

func (s *AAAService) CountAAAKCDAccount() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", aaaKCDAccountURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AAAKCDAccount `json:"aaakcdaccount"`
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

func (s *AAAService) CheckAAAKCDAccount(resource models.AAAKCDAccount) error {
	payload := map[string]any{"aaakcdaccount": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=check", aaaKCDAccountURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// aaaldapparams
// Configuration for LDAP parameter resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaaldapparams
func (s *AAAService) UpdateAAALDAPParams(resource models.AAALDAPParams) error {
	payload := map[string]any{"aaaldapparams": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, aaaLDAPParamsURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) UnsetAAALDAPParams(resource models.AAALDAPParams) error {
	payload := map[string]any{"aaaldapparams": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", aaaLDAPParamsURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) GetAllAAALDAPParams() (models.AAALDAPParams, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaLDAPParamsURL, nil)
	if err != nil {
		return models.AAALDAPParams{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.AAALDAPParams{}, err
	}

	var result struct {
		Data models.AAALDAPParams `json:"aaaldapparams"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.AAALDAPParams{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// aaaotpparamter
// Configuration for AAA otpparameter resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaaotpparameter
func (s *AAAService) UpdateAAAOTPParameter(resource models.AAAOTPParameter) error {
	payload := map[string]any{"aaaotpparameter": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, aaaOTPParameterURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) UnsetAAAOTPParameter(resource models.AAAOTPParameter) error {
	payload := map[string]any{"aaaotpparameter": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", aaaOTPParameterURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) GetAllAAAOTPParameter() (models.AAAOTPParameter, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaOTPParameterURL, nil)
	if err != nil {
		return models.AAAOTPParameter{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.AAAOTPParameter{}, err
	}

	var result struct {
		Data models.AAAOTPParameter `json:"aaaotpparameter"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.AAAOTPParameter{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// aaaparameter
// Configuration for AAA parameter resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaaparameter
func (s *AAAService) UpdateAAAParameter(resource models.AAAParameter) error {
	payload := map[string]any{"aaaparameter": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, aaaParameterURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) UnsetAAAParameter(resource models.AAAParameter) error {
	payload := map[string]any{"aaaparameter": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", aaaParameterURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) GetAllAAAParameter() (models.AAAParameter, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaParameterURL, nil)
	if err != nil {
		return models.AAAParameter{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.AAAParameter{}, err
	}

	var result struct {
		Data models.AAAParameter `json:"aaaparameter"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.AAAParameter{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// aaapreauthenticationaction
// Configuration for pre authentication action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaapreauthenticationaction
func (s *AAAService) AddAAAPreauthenticationAction(resource models.AAAPreAuthenticationAction) error {
	payload := map[string]any{"aaapreauthenticationaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, aaapreauthenticationActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) DeleteAAAPreauthenticationAction(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", aaapreauthenticationActionURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) UpdateAAAPreauthenticationAction(resource models.AAAPreAuthenticationAction) error {
	payload := map[string]any{"aaapreauthenticationaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, aaapreauthenticationActionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) UnsetAAAPreauthenticationAction(resource models.AAAPreAuthenticationAction) error {
	payload := map[string]any{"aaapreauthenticationaction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", aaapreauthenticationActionURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) GetAllAAAPreauthenticationAction() ([]models.AAAPreAuthenticationAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaapreauthenticationActionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAPreAuthenticationAction `json:"aaapreauthenticationaction"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) GetAAAPreauthenticationAction(name string) (models.AAAPreAuthenticationAction, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", aaapreauthenticationActionURL, name), nil)
	if err != nil {
		return models.AAAPreAuthenticationAction{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.AAAPreAuthenticationAction{}, err
	}

	var result struct {
		Data []models.AAAPreAuthenticationAction `json:"aaapreauthenticationaction"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.AAAPreAuthenticationAction{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.AAAPreAuthenticationAction{}, fmt.Errorf("aaapreauthenticationaction %s not found", name)
	}

	return result.Data[0], nil
}

func (s *AAAService) CountAAAPreauthenticationAction() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", aaapreauthenticationActionURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AAAPreAuthenticationAction `json:"aaapreauthenticationaction"`
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

// aaapreauthenticationparameter
// Configuration for pre authentication parameter resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaapreauthenticationparameter
func (s *AAAService) UpdateAAAPreauthenticationParameter(resource models.AAAPreAuthenticationParameter) error {
	payload := map[string]any{"aaapreauthenticationparameter": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, aaaPreauthenticationParameterURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) UnsetAAAPreauthenticationParameter(resource models.AAAPreAuthenticationParameter) error {
	payload := map[string]any{"aaapreauthenticationparameter": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", aaaPreauthenticationParameterURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) GetAllAAAPreauthenticationParameter() (models.AAAPreAuthenticationParameter, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaPreauthenticationParameterURL, nil)
	if err != nil {
		return models.AAAPreAuthenticationParameter{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.AAAPreAuthenticationParameter{}, err
	}

	var result struct {
		Data models.AAAPreAuthenticationParameter `json:"aaapreauthenticationparameter"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.AAAPreAuthenticationParameter{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// aaapreauthenticationpolicy
// Configuration for pre authentication policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaapreauthenticationpolicy
func (s *AAAService) AddAAAPreauthenticationPolicy(resource models.AAAPreAuthenticationPolicy) error {
	payload := map[string]any{"aaapreauthenticationpolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, aaaPreauthenticationPolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) DeleteAAAPreauthenticationPolicy(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", aaaPreauthenticationPolicyURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) UpdateAAAPreauthenticationPolicy(resource models.AAAPreAuthenticationPolicy) error {
	payload := map[string]any{"aaapreauthenticationpolicy": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, aaaPreauthenticationPolicyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) GetAllAAAPreauthenticationPolicy() ([]models.AAAPreAuthenticationPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaPreauthenticationPolicyURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAPreAuthenticationPolicy `json:"aaapreauthenticationpolicy"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) GetAAAPreauthenticationPolicy(name string) (models.AAAPreAuthenticationPolicy, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", aaaPreauthenticationPolicyURL, name), nil)
	if err != nil {
		return models.AAAPreAuthenticationPolicy{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.AAAPreAuthenticationPolicy{}, err
	}

	var result struct {
		Data []models.AAAPreAuthenticationPolicy `json:"aaapreauthenticationpolicy"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.AAAPreAuthenticationPolicy{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.AAAPreAuthenticationPolicy{}, fmt.Errorf("aaapreauthenticationpolicy %s not found", name)
	}

	return result.Data[0], nil
}

func (s *AAAService) CountAAAPreauthenticationPolicy() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", aaaPreauthenticationPolicyURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AAAPreAuthenticationPolicy `json:"aaapreauthenticationpolicy"`
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

// aaapreauthenticationpolicy_aaaglobal_binding
// Binding object showing the aaaglobal that can be bound to aaapreauthenticationpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaapreauthenticationpolicy_aaaglobal_binding
func (s *AAAService) GetAllAAAPreauthenticationPolicyAAAGlobalBinding() ([]models.AAAPreAuthenticationPolicyAAAGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaPreauthenticationPolicyAAAGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAPreAuthenticationPolicyAAAGlobalBinding `json:"aaapreauthenticationpolicy_aaaglobal_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) GetAAAPreauthenticationPolicyAAAGlobalBinding(name string) (models.AAAPreAuthenticationPolicyAAAGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", aaaPreauthenticationPolicyAAAGlobalBindingURL, name), nil)
	if err != nil {
		return models.AAAPreAuthenticationPolicyAAAGlobalBinding{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.AAAPreAuthenticationPolicyAAAGlobalBinding{}, err
	}

	var result struct {
		Data []models.AAAPreAuthenticationPolicyAAAGlobalBinding `json:"aaapreauthenticationpolicy_aaaglobal_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.AAAPreAuthenticationPolicyAAAGlobalBinding{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.AAAPreAuthenticationPolicyAAAGlobalBinding{}, fmt.Errorf("aaapreauthenticationpolicy_aaaglobal_binding %s not found", name)
	}

	return result.Data[0], nil
}

func (s *AAAService) CountAAAPreauthenticationPolicyAAAGlobalBinding(name string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", aaaPreauthenticationPolicyAAAGlobalBindingURL, name), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AAAPreAuthenticationPolicyAAAGlobalBinding `json:"aaapreauthenticationpolicy_aaaglobal_binding"`
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

// aaapreauthenticationpolicy_binding
// Binding object which returns the resources bound to aaapreauthenticationpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaapreauthenticationpolicy_binding
func (s *AAAService) GetAllAAAPreauthenticationPolicyBinding() ([]models.AAAPreAuthenticationPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaPreauthenticationPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAPreAuthenticationPolicyBinding `json:"aaapreauthenticationpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) GetAAAPreauthenticationPolicyBinding(name string) (models.AAAPreAuthenticationPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", aaaPreauthenticationPolicyBindingURL, name), nil)
	if err != nil {
		return models.AAAPreAuthenticationPolicyBinding{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.AAAPreAuthenticationPolicyBinding{}, err
	}

	var result struct {
		Data []models.AAAPreAuthenticationPolicyBinding `json:"aaapreauthenticationpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.AAAPreAuthenticationPolicyBinding{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.AAAPreAuthenticationPolicyBinding{}, fmt.Errorf("aaapreauthenticationpolicy_binding %s not found", name)
	}

	return result.Data[0], nil
}

// aaapreauthentiucationpolicy_vpnvserver_binding
// Binding object showing the vpnvserver that can be bound to aaapreauthenticationpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaapreauthenticationpolicy_vpnvserver_binding
func (s *AAAService) GetAllAAAPreauthenticationPolicyVPNVServerBinding() ([]models.AAAPreAuthenticationPolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaPreauthenticationPolicyVPNVServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAPreAuthenticationPolicyVPNVServerBinding `json:"aaapreauthenticationpolicy_vpnvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) GetAAAPreauthenticationPolicyVPNVServerBinding(name string) (models.AAAPreAuthenticationPolicyVPNVServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", aaaPreauthenticationPolicyVPNVServerBindingURL, name), nil)
	if err != nil {
		return models.AAAPreAuthenticationPolicyVPNVServerBinding{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.AAAPreAuthenticationPolicyVPNVServerBinding{}, err
	}

	var result struct {
		Data []models.AAAPreAuthenticationPolicyVPNVServerBinding `json:"aaapreauthenticationpolicy_vpnvserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.AAAPreAuthenticationPolicyVPNVServerBinding{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.AAAPreAuthenticationPolicyVPNVServerBinding{}, fmt.Errorf("aaapreauthenticationpolicy_vpnvserver_binding %s not found", name)
	}

	return result.Data[0], nil
}

func (s *AAAService) CountAAAPreauthenticationPolicyVPNVServerBinding(name string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", aaaPreauthenticationPolicyVPNVServerBindingURL, name), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AAAPreAuthenticationPolicyVPNVServerBinding `json:"aaapreauthenticationpolicy_vpnvserver_binding"`
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

// aaaradiusparams
// Configuration for RADIUS parameter resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaaradiusparams
func (s *AAAService) UpdateAAARADIUSParams(resource models.AAARADIUSParams) error {
	payload := map[string]any{"aaaradiusparams": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, aaaRADIUSParamsURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) UnsetAAARADIUSParams(resource models.AAARADIUSParams) error {
	payload := map[string]any{"aaaradiusparams": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", aaaRADIUSParamsURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) GetAllAAARADIUSParams() (models.AAARADIUSParams, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaRADIUSParamsURL, nil)
	if err != nil {
		return models.AAARADIUSParams{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.AAARADIUSParams{}, err
	}

	var result struct {
		Data models.AAARADIUSParams `json:"aaaradiusparams"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.AAARADIUSParams{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// aaasession
// Configuration for active connection resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaasession
func (s *AAAService) GetAllAAASession() ([]models.AAASession, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaSessionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAASession `json:"aaasession"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) CountAAASession() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", aaaSessionURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AAASession `json:"aaasession"`
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

func (s *AAAService) KillAAASession(resource models.AAASession) error {
	payload := map[string]any{"aaasession": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=kill", aaaSessionURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// aaassoprofile
// Configuration for aaa sso profile resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaassoprofile
func (s *AAAService) AddAAASSOProfile(resource models.AAASSOProfile) error {
	payload := map[string]any{"aaassoprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, aaaSSOProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) DeleteAAASSOProfile(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", aaaSSOProfileURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) GetAllAAASSOProfile() ([]models.AAASSOProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaSSOProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAASSOProfile `json:"aaassoprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) GetAAASSOProfile(name string) (models.AAASSOProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", aaaSSOProfileURL, name), nil)
	if err != nil {
		return models.AAASSOProfile{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.AAASSOProfile{}, err
	}

	var result struct {
		Data []models.AAASSOProfile `json:"aaassoprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.AAASSOProfile{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.AAASSOProfile{}, fmt.Errorf("aaassoprofile %s not found", name)
	}

	return result.Data[0], nil
}

func (s *AAAService) CountAAASSOProfile() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", aaaSSOProfileURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AAASSOProfile `json:"aaassoprofile"`
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

func (s *AAAService) UpdateAAASSOProfile(resource models.AAASSOProfile) error {
	payload := map[string]any{"aaassoprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, aaaSSOProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// aaatacasparams
// Configuration for tacacs parameters resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaatacacsparams
func (s *AAAService) UpdateAAATACACSParams(resource models.AAATACACSParams) error {
	payload := map[string]any{"aaatacacsparams": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, aaaTACACSParamsURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) UnsetAAATACACSParams(resource models.AAATACACSParams) error {
	payload := map[string]any{"aaatacacsparams": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", aaaTACACSParamsURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) GetAllAAATACACSParams() (models.AAATACACSParams, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaTACACSParamsURL, nil)
	if err != nil {
		return models.AAATACACSParams{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.AAATACACSParams{}, err
	}

	var result struct {
		Data models.AAATACACSParams `json:"aaatacacsparams"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.AAATACACSParams{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// aaauser
// Configuration for AAA user resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaauser
func (s *AAAService) AddAAAUser(resource models.AAAUser) error {
	payload := map[string]any{"aaauser": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, aaaUserURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) DeleteAAAUser(username string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", aaaUserURL, username), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) UpdateAAAUser(resource models.AAAUser) error {
	payload := map[string]any{"aaauser": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, aaaUserURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) GetAllAAAUser() ([]models.AAAUser, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaUserURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAUser `json:"aaauser"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) GetAAAUser(username string) (models.AAAUser, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", aaaUserURL, username), nil)
	if err != nil {
		return models.AAAUser{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.AAAUser{}, err
	}

	var result struct {
		Data []models.AAAUser `json:"aaauser"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.AAAUser{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.AAAUser{}, fmt.Errorf("aaauser %s not found", username)
	}

	return result.Data[0], nil
}

func (s *AAAService) CountAAAUser() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", aaaUserURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AAAUser `json:"aaauser"`
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

func (s *AAAService) UnlockAAAUser(resource models.AAAUser) error {
	payload := map[string]any{"aaauser": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unlock", aaaUserURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// aaauser_aaagroup_binding
// Binding object showing the aaagroup that can be bound to aaauser.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaauser_aaagroup_binding
func (s *AAAService) GetAllAAAUserAAAGroupBinding() ([]models.AAAUserAAAGroupBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaUserAAAGroupBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAUserAAAGroupBinding `json:"aaauser_aaagroup_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) GetAAAUserAAAGroupBinding(username string) ([]models.AAAUserAAAGroupBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", aaaUserAAAGroupBindingURL, username), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAUserAAAGroupBinding `json:"aaauser_aaagroup_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) CountAAAUserAAAGroupBinding(username string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", aaaUserAAAGroupBindingURL, username), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AAAUserAAAGroupBinding `json:"aaauser_aaagroup_binding"`
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

// aaauser_auditnslogpolicy_binding
// Binding object showing the auditnslogpolicy that can be bound to aaauser.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaauser_auditnslogpolicy_binding
func (s *AAAService) AddAAAUserAuditNSLogPolicyBinding(resource models.AAAUserAuditNSLogPolicyBinding) error {
	payload := map[string]any{"aaauser_auditnslogpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, aaaUserAuditnslogPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) DeleteAAAUserAuditNSLogPolicyBinding(username string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", aaaUserAuditnslogPolicyBindingURL, username, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) GetAllAAAUserAuditNSLogPolicyBinding() ([]models.AAAUserAuditNSLogPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaUserAuditnslogPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAUserAuditNSLogPolicyBinding `json:"aaauser_auditnslogpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) GetAAAUserAuditNSLogPolicyBinding(username string) ([]models.AAAUserAuditNSLogPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", aaaUserAuditnslogPolicyBindingURL, username), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAUserAuditNSLogPolicyBinding `json:"aaauser_auditnslogpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) CountAAAUserAuditNSLogPolicyBinding(username string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", aaaUserAuditnslogPolicyBindingURL, username), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AAAUserAuditNSLogPolicyBinding `json:"aaauser_auditnslogpolicy_binding"`
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

// aaauser_auditsyslogpolicy_binding
// Binding object showing the auditsyslogpolicy that can be bound to aaauser.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaauser_auditsyslogpolicy_binding
func (s *AAAService) AddAAAUserAuditSyslogPolicyBinding(resource models.AAAUserAuditSyslogPolicyBinding) error {
	payload := map[string]any{"aaauser_auditsyslogpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, aaaUserAuditsyslogPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) DeleteAAAUserAuditSyslogPolicyBinding(username string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", aaaUserAuditsyslogPolicyBindingURL, username, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) GetAllAAAUserAuditSyslogPolicyBinding() ([]models.AAAUserAuditSyslogPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaUserAuditsyslogPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAUserAuditSyslogPolicyBinding `json:"aaauser_auditsyslogpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) GetAAAUserAuditSyslogPolicyBinding(username string) ([]models.AAAUserAuditSyslogPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", aaaUserAuditsyslogPolicyBindingURL, username), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAUserAuditSyslogPolicyBinding `json:"aaauser_auditsyslogpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) CountAAAUserAuditSyslogPolicyBinding(username string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", aaaUserAuditsyslogPolicyBindingURL, username), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AAAUserAuditSyslogPolicyBinding `json:"aaauser_auditsyslogpolicy_binding"`
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

// aaauser_authorizationpolicy_binding
// Binding object showing the authorizationpolicy that can be bound to aaauser.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaauser_authorizationpolicy_binding
func (s *AAAService) AddAAAUserAuthorizationPolicyBinding(resource models.AAAUserAuthorizationPolicyBinding) error {
	payload := map[string]any{"aaauser_authorizationpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, aaaUserAuthorizationPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) DeleteAAAUserAuthorizationPolicyBinding(username string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", aaaUserAuthorizationPolicyBindingURL, username, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) GetAllAAAUserAuthorizationPolicyBinding() ([]models.AAAUserAuthorizationPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaUserAuthorizationPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAUserAuthorizationPolicyBinding `json:"aaauser_authorizationpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) GetAAAUserAuthorizationPolicyBinding(username string) ([]models.AAAUserAuthorizationPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", aaaUserAuthorizationPolicyBindingURL, username), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAUserAuthorizationPolicyBinding `json:"aaauser_authorizationpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) CountAAAUserAuthorizationPolicyBinding(username string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", aaaUserAuthorizationPolicyBindingURL, username), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AAAUserAuthorizationPolicyBinding `json:"aaauser_authorizationpolicy_binding"`
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

// aaauser_binding
// Binding object which returns the resources bound to aaauser.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaauser_binding
func (s *AAAService) GetAllAAAUserBinding() ([]models.AAAUserBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaUserBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAUserBinding `json:"aaauser_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) GetAAAUserBinding(username string) (models.AAAUserBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", aaaUserBindingURL, username), nil)
	if err != nil {
		return models.AAAUserBinding{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.AAAUserBinding{}, err
	}

	var result struct {
		Data []models.AAAUserBinding `json:"aaauser_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.AAAUserBinding{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.AAAUserBinding{}, fmt.Errorf("aaauser_binding %s not found", username)
	}

	return result.Data[0], nil
}

// aaauser_intranetip6_binding
// Binding object showing the intranetip6 that can be bound to aaauser.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaauser_intranetip6_binding
func (s *AAAService) AddAAAUserIntranetIP6Binding(resource models.AAAUserIntranetIP6Binding) error {
	payload := map[string]any{"aaauser_intranetip6_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, aaaUserIntranetIP6BindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) DeleteAAAUserIntranetIP6Binding(username string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", aaaUserIntranetIP6BindingURL, username, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) GetAllAAAUserIntranetIP6Binding() ([]models.AAAUserIntranetIP6Binding, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaUserIntranetIP6BindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAUserIntranetIP6Binding `json:"aaauser_intranetip6_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) GetAAAUserIntranetIP6Binding(username string) ([]models.AAAUserIntranetIP6Binding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", aaaUserIntranetIP6BindingURL, username), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAUserIntranetIP6Binding `json:"aaauser_intranetip6_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) CountAAAUserIntranetIP6Binding(username string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", aaaUserIntranetIP6BindingURL, username), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AAAUserIntranetIP6Binding `json:"aaauser_intranetip6_binding"`
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

// aaauser_intranetip_binding
// Binding object showing the intranetip that can be bound to aaauser.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaauser_intranetip_binding
func (s *AAAService) AddAAAUserIntranetIPBinding(resource models.AAAUserIntranetIPBinding) error {
	payload := map[string]any{"aaauser_intranetip_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, aaaUserIntranetIPBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) DeleteAAAUserIntranetIPBinding(username string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", aaaUserIntranetIPBindingURL, username, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) GetAllAAAUserIntranetIPBinding() ([]models.AAAUserIntranetIPBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaUserIntranetIPBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAUserIntranetIPBinding `json:"aaauser_intranetip_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) GetAAAUserIntranetIPBinding(username string) ([]models.AAAUserIntranetIPBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", aaaUserIntranetIPBindingURL, username), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAUserIntranetIPBinding `json:"aaauser_intranetip_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) CountAAAUserIntranetIPBinding(username string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", aaaUserIntranetIPBindingURL, username), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AAAUserIntranetIPBinding `json:"aaauser_intranetip_binding"`
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

// aaauser_tmsessionpolicy_binding
// Binding object showing the tmsessionpolicy that can be bound to aaauser.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaauser_tmsessionpolicy_binding
func (s *AAAService) AddAAAUserTMSessionPolicyBinding(resource models.AAAUserTMSessionPolicyBinding) error {
	payload := map[string]any{"aaauser_tmsessionpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, aaaUserTMSessionPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) DeleteAAAUserTMSessionPolicyBinding(username string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", aaaUserTMSessionPolicyBindingURL, username, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) GetAllAAAUserTMSessionPolicyBinding() ([]models.AAAUserTMSessionPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaUserTMSessionPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAUserTMSessionPolicyBinding `json:"aaauser_tmsessionpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) GetAAAUserTMSessionPolicyBinding(username string) ([]models.AAAUserTMSessionPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", aaaUserTMSessionPolicyBindingURL, username), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAUserTMSessionPolicyBinding `json:"aaauser_tmsessionpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) CountAAAUserTMSessionPolicyBinding(username string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", aaaUserTMSessionPolicyBindingURL, username), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AAAUserTMSessionPolicyBinding `json:"aaauser_tmsessionpolicy_binding"`
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

// aaauser_vpnintranetapplication_binding
// Binding object showing the vpnintranetapplication that can be bound to aaauser.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaauser_vpnintranetapplication_binding
func (s *AAAService) AddAAAUserVPNIntranetApplicationBinding(resource models.AAAUserVPNIntranetApplicationBinding) error {
	payload := map[string]any{"aaauser_vpnintranetapplication_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, aaaUserVPNIntranetApplicationBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) DeleteAAAUserVPNIntranetApplicationBinding(username string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", aaaUserVPNIntranetApplicationBindingURL, username, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) GetAllAAAUserVPNIntranetApplicationBinding() ([]models.AAAUserVPNIntranetApplicationBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaUserVPNIntranetApplicationBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAUserVPNIntranetApplicationBinding `json:"aaauser_vpnintranetapplication_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) GetAAAUserVPNIntranetApplicationBinding(username string) ([]models.AAAUserVPNIntranetApplicationBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", aaaUserVPNIntranetApplicationBindingURL, username), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAUserVPNIntranetApplicationBinding `json:"aaauser_vpnintranetapplication_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) CountAAAUserVPNIntranetApplicationBinding(username string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", aaaUserVPNIntranetApplicationBindingURL, username), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AAAUserVPNIntranetApplicationBinding `json:"aaauser_vpnintranetapplication_binding"`
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

// aaauser_vpnsessionpolicy_binding
// Binding object showing the vpnsessionpolicy that can be bound to aaauser.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaauser_vpnsessionpolicy_binding
func (s *AAAService) AddAAAUserVPNSessionPolicyBinding(resource models.AAAUserVPNSessionPolicyBinding) error {
	payload := map[string]any{"aaauser_vpnsessionpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, aaaUserVPNSessionPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) DeleteAAAUserVPNSessionPolicyBinding(username string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", aaaUserVPNSessionPolicyBindingURL, username, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) GetAllAAAUserVPNSessionPolicyBinding() ([]models.AAAUserVPNSessionPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaUserVPNSessionPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAUserVPNSessionPolicyBinding `json:"aaauser_vpnsessionpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) GetAAAUserVPNSessionPolicyBinding(username string) ([]models.AAAUserVPNSessionPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", aaaUserVPNSessionPolicyBindingURL, username), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAUserVPNSessionPolicyBinding `json:"aaauser_vpnsessionpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) CountAAAUserVPNSessionPolicyBinding(username string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", aaaUserVPNSessionPolicyBindingURL, username), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AAAUserVPNSessionPolicyBinding `json:"aaauser_vpnsessionpolicy_binding"`
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

// aaauser_vpntrafficpolicy_binding
// Binding object showing the vpntrafficpolicy that can be bound to aaauser.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaauser_vpntrafficpolicy_binding
func (s *AAAService) AddAAAUserVPNTrafficPolicyBinding(resource models.AAAUserVPNTrafficPolicyBinding) error {
	payload := map[string]any{"aaauser_vpntrafficpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, aaaUserVPNTrafficPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) DeleteAAAUserVPNTrafficPolicyBinding(username string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", aaaUserVPNTrafficPolicyBindingURL, username, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) GetAllAAAUserVPNTrafficPolicyBinding() ([]models.AAAUserVPNTrafficPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaUserVPNTrafficPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAUserVPNTrafficPolicyBinding `json:"aaauser_vpntrafficpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) GetAAAUserVPNTrafficPolicyBinding(username string) ([]models.AAAUserVPNTrafficPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", aaaUserVPNTrafficPolicyBindingURL, username), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAUserVPNTrafficPolicyBinding `json:"aaauser_vpntrafficpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) CountAAAUserVPNTrafficPolicyBinding(username string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", aaaUserVPNTrafficPolicyBindingURL, username), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AAAUserVPNTrafficPolicyBinding `json:"aaauser_vpntrafficpolicy_binding"`
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

// aaauser_vpnurlpolicy_binding
// Binding object showing the vpnurlpolicy that can be bound to aaauser.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaauser_vpnurlpolicy_binding
func (s *AAAService) AddAAAUserVPNURLPolicyBinding(resource models.AAAUserVPNURLPolicyBinding) error {
	payload := map[string]any{"aaauser_vpnurlpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, aaaUserVPNURLPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) DeleteAAAUserVPNURLPolicyBinding(username string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", aaaUserVPNURLPolicyBindingURL, username, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) GetAllAAAUserVPNURLPolicyBinding() ([]models.AAAUserVPNURLPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaUserVPNURLPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAUserVPNURLPolicyBinding `json:"aaauser_vpnurlpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) GetAAAUserVPNURLPolicyBinding(username string) ([]models.AAAUserVPNURLPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", aaaUserVPNURLPolicyBindingURL, username), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAUserVPNURLPolicyBinding `json:"aaauser_vpnurlpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) CountAAAUserVPNURLPolicyBinding(username string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", aaaUserVPNURLPolicyBindingURL, username), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AAAUserVPNURLPolicyBinding `json:"aaauser_vpnurlpolicy_binding"`
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

// aaauser_vpnurl_binding
// Binding object showing the vpnurl that can be bound to aaauser.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/aaa/aaauser_vpnurl_binding
func (s *AAAService) AddAAAUserVPNURLBinding(resource models.AAAUserVPNURLBinding) error {
	payload := map[string]any{"aaauser_vpnurl_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, aaaUserVPNURLBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) DeleteAAAUserVPNURLBinding(username string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", aaaUserVPNURLBindingURL, username, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *AAAService) GetAllAAAUserVPNURLBinding() ([]models.AAAUserVPNURLBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, aaaUserVPNURLBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAUserVPNURLBinding `json:"aaauser_vpnurl_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) GetAAAUserVPNURLBinding(username string) ([]models.AAAUserVPNURLBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", aaaUserVPNURLBindingURL, username), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AAAUserVPNURLBinding `json:"aaauser_vpnurl_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *AAAService) CountAAAUserVPNURLBinding(username string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", aaaUserVPNURLBindingURL, username), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.AAAUserVPNURLBinding `json:"aaauser_vpnurl_binding"`
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
