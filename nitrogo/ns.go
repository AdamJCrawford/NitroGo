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
	nsACLURL                                   = "/nitro/v1/config/nsacl"
	nsACL6URL                                  = "/nitro/v1/config/nsacl6"
	nsACLSURL                                  = "/nitro/v1/config/nsacls"
	nsACLS6URL                                 = "/nitro/v1/config/nsacls6"
	nsAppFlowCollectorURL                      = "/nitro/v1/config/nsappflowcollector"
	nsAppFlowParamURL                          = "/nitro/v1/config/nsappflowparam"
	nsCapacityURL                              = "/nitro/v1/config/nscapacity"
	nsConfigURL                                = "/nitro/v1/config/nsconfig"
	nsConnectionTableURL                       = "/nitro/v1/config/nsconnectiontable"
	nsConsoleLoginPromptURL                    = "/nitro/v1/config/nsconsoleloginprompt"
	nsDHCPParamsURL                            = "/nitro/v1/config/nsdhcpparams"
	nsDialogURL                                = "/nitro/v1/config/nsdialog"
	nsEncryptionParamsURL                      = "/nitro/v1/config/nsencryptionparams"
	nsEventsURL                                = "/nitro/v1/config/nsevents"
	nsExtensionURL                             = "/nitro/v1/config/nsextension"
	nsExtensionBindingURL                      = "/nitro/v1/config/nsextension_binding"
	nsExtensionExtensionFunctionBindingURL     = "/nitro/v1/config/nsextension_extensionfunction_binding"
	nsFeatureURL                               = "/nitro/v1/config/nsfeature"
	nsHardwareURL                              = "/nitro/v1/config/nshardware"
	nsHostNameURL                              = "/nitro/v1/config/nshostname"
	nsHTTPParamURL                             = "/nitro/v1/config/nshttpparam"
	nsHTTPProfileURL                           = "/nitro/v1/config/nshttpprofile"
	nsICAPProfileURL                           = "/nitro/v1/config/nsicapprofile"
	nsIPURL                                    = "/nitro/v1/config/nsip"
	nsIP6URL                                   = "/nitro/v1/config/nsip6"
	nsLicenseURL                               = "/nitro/v1/config/nslicense"
	nsLicenseParametersURL                     = "/nitro/v1/config/nslicenseparameters"
	nsLicenseServerURL                         = "/nitro/v1/config/nslicenseserver"
	nsLimitIdentifierURL                       = "/nitro/v1/config/nslimitidentifier"
	nsLimitSelectorURL                         = "/nitro/v1/config/nslimitselector"
	nsLimitSessionsURL                         = "/nitro/v1/config/nslimitsessions"
	nsLogActionURL                             = "/nitro/v1/config/nslogaction"
	nsLogGlobalBindingURL                      = "/nitro/v1/config/nslogglobal_binding"
	nsLogGlobalAuditNSLogPolicyBindingURL      = "/nitro/v1/config/nslogglobal_auditnslogpolicy_binding"
	nsLogParamsURL                             = "/nitro/v1/config/nslogparams"
	nsLogPolicyURL                             = "/nitro/v1/config/nslogpolicy"
	nsLogPolicyBindingURL                      = "/nitro/v1/config/nslogpolicy_binding"
	nsModeURL                                  = "/nitro/v1/config/nsmode"
	nsParamURL                                 = "/nitro/v1/config/nsparam"
	nsPartitionURL                             = "/nitro/v1/config/nspartition"
	nsPartitionBindingURL                      = "/nitro/v1/config/nspartition_binding"
	nsPartitionBridgeGroupBindingURL           = "/nitro/v1/config/nspartition_bridgegroup_binding"
	nsPartitionVLANBindingURL                  = "/nitro/v1/config/nspartition_vlan_binding"
	nsPartitionVXLANBindingURL                 = "/nitro/v1/config/nspartition_vxlan_binding"
	nsRateControlURL                           = "/nitro/v1/config/nsratecontrol"
	nsRPCNodeURL                               = "/nitro/v1/config/nsrpcnode"
	nsRunningConfigURL                         = "/nitro/v1/config/nsrunningconfig"
	nsSavedConfigURL                           = "/nitro/v1/config/nssavedconfig"
	nsServicePathURL                           = "/nitro/v1/config/nsservicepath"
	nsSimpleACLURL                             = "/nitro/v1/config/nssimpleacl"
	nsSimpleACL6URL                            = "/nitro/v1/config/nssimpleacl6"
	nsSPParamsURL                              = "/nitro/v1/config/nsspparams"
	nsSurgeProtectionURL                       = "/nitro/v1/config/nssurgeprotection"
	nsTCPBufParamURL                           = "/nitro/v1/config/nstcpbufparam"
	nsTCPParamURL                              = "/nitro/v1/config/nstcpparam"
	nsTCPProfileURL                            = "/nitro/v1/config/nstcpprofile"
	nsTimeoutURL                               = "/nitro/v1/config/nstimeout"
	nsTimerURL                                 = "/nitro/v1/config/nstimer"
	nsTrafficDomainURL                         = "/nitro/v1/config/nstrafficdomain"
	nsTrafficDomainBindingURL                  = "/nitro/v1/config/nstrafficdomain_binding"
	nsTrafficDomainBridgeGroupBindingURL       = "/nitro/v1/config/nstrafficdomain_bridgegroup_binding"
	nsTrafficDomainVLANBindingURL              = "/nitro/v1/config/nstrafficdomain_vlan_binding"
	nsTrafficDomainVXLANBindingURL             = "/nitro/v1/config/nstrafficdomain_vxlan_binding"
	nsVariableURL                              = "/nitro/v1/config/nsvariable"
	nsVersionURL                               = "/nitro/v1/config/nsversion"
	nsWeblogParamURL                           = "/nitro/v1/config/nsweblogparam"
	nsXMLNamespaceURL                          = "/nitro/v1/config/nsxmlnamespace"
	nsXMLSchemaURL                             = "/nitro/v1/config/nsxmlschema"
	nsAptLicenseURL                            = "/nitro/v1/config/nsaptlicense"
	nsAssignmentURL                            = "/nitro/v1/config/nsassignment"
	nsCQAParamURL                              = "/nitro/v1/config/nscqaparam"
	nsDHCPIPURL                                = "/nitro/v1/config/nsdhcpip"
	nsDiameterURL                              = "/nitro/v1/config/nsdiameter"
	nsEncryptionKeyURL                         = "/nitro/v1/config/nsencryptionkey"
	nsHostnameURL                              = "/nitro/v1/config/nshostname"
	nsLicenseProxyServerURL                    = "/nitro/v1/config/nslicenseproxyserver"
	nsLicenseServerPoolURL                     = "/nitro/v1/config/nslicenseserverpool"
	nsMigrationURL                             = "/nitro/v1/config/nsmigration"
	nsPartitionMACURL                          = "/nitro/v1/config/nspartitionmac"
	nsPBRURL                                   = "/nitro/v1/config/nspbr"
	nsPBR6URL                                  = "/nitro/v1/config/nspbr6"
	nsPBRsURL                                  = "/nitro/v1/config/nspbrs"
	nsRollBackCMDURL                           = "/nitro/v1/config/nsrollbackcmd"
	nsServiceFunctionURL                       = "/nitro/v1/config/nsservicefunction"
	nsServicePathBindingURL                    = "/nitro/v1/config/nsservicepath_binding"
	nsServicePathNSServiceFunctionBindingURL   = "/nitro/v1/config/nsservicepath_nsservicefunction_binding"
	nsSourceRouteCacheTableURL                 = "/nitro/v1/config/nssourceroutecachetable"
	nsStatsURL                                 = "/nitro/v1/config/nsstats"
	nsSurgeQURL                                = "/nitro/v1/config/nssurgeq"
	nsTimerAutoScalePolicyBindingURL           = "/nitro/v1/config/nstimer_autoscalepolicy_binding"
	nsTimerBindingURL                          = "/nitro/v1/config/nstimer_binding"
	nsTimeZoneURL                              = "/nitro/v1/config/nstimezone"
	nsVPXParamURL                              = "/nitro/v1/config/nsvpxparam"
	nsCentralManagementServerURL               = "/nitro/v1/config/nscentralmanagementserver"
	nsHMACKeyURL                               = "/nitro/v1/config/nshmackey"
	nsLimitIdentifierBindingURL                = "/nitro/v1/config/nslimitidentifier_binding"
	nsLimitIdentifierNSLimitSessionsBindingURL = "/nitro/v1/config/nslimitidentifier_nslimitsessions_binding"
	nsPartitionVlanBindingURL                  = "/nitro/v1/config/nspartition_vlan_binding"
	rebootURL                                  = "/nitro/v1/config/reboot"
	shutdownURL                                = "/nitro/v1/config/shutdown"
)

// ns
// System/Global level configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/ns/ns
type NSService struct {
	client *Client
}

// nsacl
func (s *NSService) AddNSACL(resource models.NSACL) error {
	payload := map[string]any{"nsacl": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nsACLURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DeleteNSACL(aclname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nsACLURL, aclname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UpdateNSACL(resource models.NSACL) error {
	payload := map[string]any{"nsacl": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsACLURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UnsetNSACL(resource models.NSACL) error {
	payload := map[string]any{"nsacl": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", nsACLURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) EnableNSACL(resource models.NSACL) error {
	payload := map[string]any{"nsacl": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=enable", nsACLURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DisableNSACL(resource models.NSACL) error {
	payload := map[string]any{"nsacl": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=disable", nsACLURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) RenameNSACL(resource models.NSACL) error {
	payload := map[string]any{"nsacl": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=rename", nsACLURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSACL() ([]models.NSACL, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsACLURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSACL `json:"nsacl"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSACL(aclname string) (*models.NSACL, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsACLURL, aclname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSACL `json:"nsacl"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("nsacl %s not found", aclname)
	}

	return &result.Data[0], nil
}

func (s *NSService) CountNSACL() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsACLURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSACL `json:"nsacl"`
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

// nsacl6
func (s *NSService) AddNSACL6(resource models.NSACL6) error {
	payload := map[string]any{"nsacl6": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nsACL6URL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DeleteNSACL6(acl6name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nsACL6URL, acl6name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UpdateNSACL6(resource models.NSACL6) error {
	payload := map[string]any{"nsacl6": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsACL6URL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UnsetNSACL6(resource models.NSACL6) error {
	payload := map[string]any{"nsacl6": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", nsACL6URL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) EnableNSACL6(resource models.NSACL6) error {
	payload := map[string]any{"nsacl6": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=enable", nsACL6URL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DisableNSACL6(resource models.NSACL6) error {
	payload := map[string]any{"nsacl6": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=disable", nsACL6URL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) RenameNSACL6(resource models.NSACL6) error {
	payload := map[string]any{"nsacl6": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=rename", nsACL6URL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSACL6() ([]models.NSACL6, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsACL6URL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSACL6 `json:"nsacl6"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSACL6(acl6name string) (*models.NSACL6, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsACL6URL, acl6name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSACL6 `json:"nsacl6"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("nsacl6 %s not found", acl6name)
	}

	return &result.Data[0], nil
}

func (s *NSService) CountNSACL6() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsACL6URL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSACL6 `json:"nsacl6"`
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

// nsacls
func (s *NSService) RenumberNSACLS() error {
	payload := map[string]any{"nsacls": map[string]any{}}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=renumber", nsACLSURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) ClearNSACLS() error {
	payload := map[string]any{"nsacls": map[string]any{}}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=clear", nsACLSURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) ApplyNSACLS() error {
	payload := map[string]any{"nsacls": map[string]any{}}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=apply", nsACLSURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// nsacls6
func (s *NSService) RenumberNSACLS6() error {
	payload := map[string]any{"nsacls6": map[string]any{}}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=renumber", nsACLS6URL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) ClearNSACLS6() error {
	payload := map[string]any{"nsacls6": map[string]any{}}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=clear", nsACLS6URL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) ApplyNSACLS6() error {
	payload := map[string]any{"nsacls6": map[string]any{}}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=apply", nsACLS6URL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// nsappflowcollector
func (s *NSService) AddNSAppFlowCollector(resource models.NSAppFlowCollector) error {
	payload := map[string]any{"nsappflowcollector": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nsAppFlowCollectorURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DeleteNSAppFlowCollector(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nsAppFlowCollectorURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSAppFlowCollector() ([]models.NSAppFlowCollector, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsAppFlowCollectorURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSAppFlowCollector `json:"nsappflowcollector"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSAppFlowCollector(name string) (*models.NSAppFlowCollector, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsAppFlowCollectorURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSAppFlowCollector `json:"nsappflowcollector"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("nsappflowcollector %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *NSService) CountNSAppFlowCollector() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsAppFlowCollectorURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSAppFlowCollector `json:"nsappflowcollector"`
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

// nsappflowparam
func (s *NSService) UpdateNSAppFlowParam(resource models.NSAppFlowParam) error {
	payload := map[string]any{"nsappflowparam": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsAppFlowParamURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UnsetNSAppFlowParam(resource models.NSAppFlowParam) error {
	payload := map[string]any{"nsappflowparam": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", nsAppFlowParamURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSAppFlowParam() ([]models.NSAppFlowParam, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsAppFlowParamURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSAppFlowParam `json:"nsappflowparam"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// nsaptlicense
func (s *NSService) GetAllNSAptLicense() ([]models.NSAPTLicense, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsAptLicenseURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSAPTLicense `json:"nsaptlicense"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) CountNSAptLicense() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsAptLicenseURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSAPTLicense `json:"nsaptlicense"`
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

func (s *NSService) ChangeNSAptLicense(resource models.NSAPTLicense) error {
	payload := map[string]any{"nsaptlicense": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=change", nsAptLicenseURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// nsassignment
func (s *NSService) AddNSAssignment(resource models.NSAssignment) error {
	payload := map[string]any{"nsassignment": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nsAssignmentURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DeleteNSAssignment(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nsAssignmentURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UpdateNSAssignment(resource models.NSAssignment) error {
	payload := map[string]any{"nsassignment": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsAssignmentURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UnsetNSAssignment(resource models.NSAssignment) error {
	payload := map[string]any{"nsassignment": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", nsAssignmentURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSAssignment() ([]models.NSAssignment, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsAssignmentURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSAssignment `json:"nsassignment"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSAssignment(name string) (*models.NSAssignment, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsAssignmentURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSAssignment `json:"nsassignment"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("nsassignment %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *NSService) CountNSAssignment() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsAssignmentURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSAssignment `json:"nsassignment"`
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

func (s *NSService) RenameNSAssignment(resource models.NSAssignment) error {
	payload := map[string]any{"nsassignment": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=rename", nsAssignmentURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// nscapacity
func (s *NSService) UpdateNSCapacity(resource models.NSCapacity) error {
	payload := map[string]any{"nscapacity": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsCapacityURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UnsetNSCapacity(resource models.NSCapacity) error {
	payload := map[string]any{"nscapacity": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", nsCapacityURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSCapacity() ([]models.NSCapacity, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsCapacityURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSCapacity `json:"nscapacity"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// nscentralmanagementserver
func (s *NSService) AddNSCentralManagementServer(resource models.NSCentralManagementServer) error {
	payload := map[string]any{"nscentralmanagementserver": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nsCentralManagementServerURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DeleteNSCentralManagementServer(servername string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nsCentralManagementServerURL, servername), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSCentralManagementServer() ([]models.NSCentralManagementServer, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsCentralManagementServerURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSCentralManagementServer `json:"nscentralmanagementserver"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSCentralManagementServer(servername string) (*models.NSCentralManagementServer, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsCentralManagementServerURL, servername), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSCentralManagementServer `json:"nscentralmanagementserver"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("nscentralmanagementserver %s not found", servername)
	}

	return &result.Data[0], nil
}

func (s *NSService) CountNSCentralManagementServer() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsCentralManagementServerURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSCentralManagementServer `json:"nscentralmanagementserver"`
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

// nsconfig
func (s *NSService) ClearNSConfig(resource models.NSConfig) error {
	payload := map[string]any{"nsconfig": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=clear", nsConfigURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) SaveNSConfig(resource models.NSConfig) error {
	payload := map[string]any{"nsconfig": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=save", nsConfigURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DiffNSConfig(resource models.NSConfig) error {
	payload := map[string]any{"nsconfig": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=diff", nsConfigURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UpdateNSConfig(resource models.NSConfig) error {
	payload := map[string]any{"nsconfig": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsConfigURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UnsetNSConfig(resource models.NSConfig) error {
	payload := map[string]any{"nsconfig": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", nsConfigURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSConfig() ([]models.NSConfig, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsConfigURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSConfig `json:"nsconfig"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// nsconnectiontable
func (s *NSService) GetAllNSconnectionTable(args map[string]string) ([]models.NSConnectionTable, error) {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", nsConnectionTableURL, argsStr), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSConnectionTable `json:"nsconnectiontable"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) CountNSconnectionTable(args map[string]string) (int, error) {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",") + "&count=yes"
	} else {
		argsStr = "?count=yes"
	}
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", nsConnectionTableURL, argsStr), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSConnectionTable `json:"nsconnectiontable"`
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

// nsconsoleloginprompt
func (s *NSService) UpdateNSConsoleLoginPrompt(resource models.NSConsoleLoginPrompt) error {
	payload := map[string]any{"nsconsoleloginprompt": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsConsoleLoginPromptURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UnsetNSConsoleLoginPrompt(resource models.NSConsoleLoginPrompt) error {
	payload := map[string]any{"nsconsoleloginprompt": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", nsConsoleLoginPromptURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSConsoleLoginPrompt() ([]models.NSConsoleLoginPrompt, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsConsoleLoginPromptURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSConsoleLoginPrompt `json:"nsconsoleloginprompt"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// nscqaparam
func (s *NSService) UpdateNSCQAParam(resource models.NSCQAParam) error {
	payload := map[string]any{"nscqaparam": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsCQAParamURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UnsetNSCQAParam(resource models.NSCQAParam) error {
	payload := map[string]any{"nscqaparam": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", nsCQAParamURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSCQAParam() ([]models.NSCQAParam, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsCQAParamURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSCQAParam `json:"nscqaparam"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// nsdhcpip
func (s *NSService) ReleaseNSDHCPIP() error {
	payload := map[string]any{"nsdhcpip": map[string]any{}}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=release", nsDHCPIPURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// nsdhcpparams
func (s *NSService) UpdateNSDHCPParams(resource models.NSDHCPParams) error {
	payload := map[string]any{"nsdhcpparams": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsDHCPParamsURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UnsetNSDHCPParams(resource models.NSDHCPParams) error {
	payload := map[string]any{"nsdhcpparams": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", nsDHCPParamsURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSDHCPParams() ([]models.NSDHCPParams, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsDHCPParamsURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSDHCPParams `json:"nsdhcpparams"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// nsdiamter
func (s *NSService) UpdateNSDiameter(resource models.NSDiameter) error {
	payload := map[string]any{"nsdiameter": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsDiameterURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UnsetNSDiameter(resource models.NSDiameter) error {
	payload := map[string]any{"nsdiameter": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", nsDiameterURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSDiameter() ([]models.NSDiameter, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsDiameterURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSDiameter `json:"nsdiameter"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSDiameter() (*models.NSDiameter, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsDiameterURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSDiameter `json:"nsdiameter"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("nsdiameter not found")
	}

	return &result.Data[0], nil
}

func (s *NSService) CountNSDiameter() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsDiameterURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSDiameter `json:"nsdiameter"`
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

// nsencyrptionkey
func (s *NSService) AddNSEncryptionKey(resource models.NSEncryptionKey) error {
	payload := map[string]any{"nsencryptionkey": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nsEncryptionKeyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DeleteNSEncryptionKey(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nsEncryptionKeyURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UpdateNSEncryptionKey(resource models.NSEncryptionKey) error {
	payload := map[string]any{"nsencryptionkey": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsEncryptionKeyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UnsetNSEncryptionKey(resource models.NSEncryptionKey) error {
	payload := map[string]any{"nsencryptionkey": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", nsEncryptionKeyURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSEncryptionKey() ([]models.NSEncryptionKey, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsEncryptionKeyURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSEncryptionKey `json:"nsencryptionkey"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSEncryptionKey(name string) (*models.NSEncryptionKey, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsEncryptionKeyURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSEncryptionKey `json:"nsencryptionkey"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("nsencryptionkey %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *NSService) CountNSEncryptionKey() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsEncryptionKeyURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSEncryptionKey `json:"nsencryptionkey"`
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

// nsencryptionparams
func (s *NSService) UpdateNSEncryptionParams(resource models.NSEncryptionParams) error {
	payload := map[string]any{"nsencryptionparams": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsEncryptionParamsURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSEncryptionParams() ([]models.NSEncryptionParams, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsEncryptionParamsURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSEncryptionParams `json:"nsencryptionparams"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// nsevents
func (s *NSService) GetAllNSEvents() ([]models.NSEvents, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsEventsURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSEvents `json:"nsevents"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) CountNSEvents() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsEventsURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSEvents `json:"nsevents"`
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

// nsextension
func (s *NSService) ImportNSExtension(resource models.NSExtension) error {
	payload := map[string]any{"nsextension": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=import", nsExtensionURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) AddNSExtension(resource models.NSExtension) error {
	payload := map[string]any{"nsextension": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nsExtensionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DeleteNSExtension(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nsExtensionURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UpdateNSExtension(resource models.NSExtension) error {
	payload := map[string]any{"nsextension": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsExtensionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UnsetNSExtension(resource models.NSExtension) error {
	payload := map[string]any{"nsextension": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", nsExtensionURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSExtension() ([]models.NSExtension, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsExtensionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSExtension `json:"nsextension"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSExtension(name string) (*models.NSExtension, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsExtensionURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSExtension `json:"nsextension"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("nsextension %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *NSService) CountNSExtension() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsExtensionURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSExtension `json:"nsextension"`
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

func (s *NSService) ChangeNSExtension(resource models.NSExtension) error {
	payload := map[string]any{"nsextension": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=change", nsExtensionURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// nsextension_binding
func (s *NSService) GetAllNSExtenionBinding() ([]models.NSExtensionBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsExtensionBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSExtensionBinding `json:"nsextension_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSExtenionBinding(name string) (*models.NSExtensionBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsExtensionBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSExtensionBinding `json:"nsextension_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("nsextension_binding %s not found", name)
	}

	return &result.Data[0], nil
}

// nsextension_extensionfunction_binding
func (s *NSService) GetAllNSExtensionExtensionFunctionBinding() ([]models.NSExtensionExtensionFunctionBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsExtensionExtensionFunctionBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSExtensionExtensionFunctionBinding `json:"nsextension_extensionfunction_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSExtensionExtensionFunctionBinding(name string) ([]models.NSExtensionExtensionFunctionBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsExtensionExtensionFunctionBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSExtensionExtensionFunctionBinding `json:"nsextension_extensionfunction_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) CountNSExtensionExtensionFunctionBinding(name string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", nsExtensionExtensionFunctionBindingURL, name), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSExtensionExtensionFunctionBinding `json:"nsextension_extensionfunction_binding"`
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

// nsfeature
func (s *NSService) EnableNSFeature(resource models.NSFeature) error {
	payload := map[string]any{"nsfeature": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=enable", nsFeatureURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DisableNSFeature(resource models.NSFeature) error {
	payload := map[string]any{"nsfeature": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=disable", nsFeatureURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSFeature() ([]models.NSFeature, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsFeatureURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSFeature `json:"nsfeature"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// nshardware
func (s *NSService) GetAllNSHardware() ([]models.NSHardware, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsHardwareURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSHardware `json:"nshardware"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// nshmackey
func (s *NSService) AddNSHMACKey(resource models.NSHMACKey) error {
	payload := map[string]any{"nshmackey": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nsHMACKeyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DeleteNSHMACKey(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nsHMACKeyURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UpdateNSHMACKey(resource models.NSHMACKey) error {
	payload := map[string]any{"nshmackey": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsHMACKeyURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UnsetNSHMACKey(resource models.NSHMACKey) error {
	payload := map[string]any{"nshmackey": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", nsHMACKeyURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSHMACKey() ([]models.NSHMACKey, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsHMACKeyURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSHMACKey `json:"nshmackey"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSHMACKey(name string) (*models.NSHMACKey, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsHMACKeyURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSHMACKey `json:"nshmackey"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("nshmackey %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *NSService) CountNSHMACKey() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsHMACKeyURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSHMACKey `json:"nshmackey"`
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

// nshostname
func (s *NSService) UpdateNSHostname(resource models.NSHostname) error {
	payload := map[string]any{"nshostname": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsHostnameURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSHostname() ([]models.NSHostname, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsHostnameURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSHostname `json:"nshostname"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) CountNSHostname() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsHostnameURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSHostname `json:"nshostname"`
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

// nshttpparam
func (s *NSService) UpdateNSHTTPParam(resource models.NSHTTPParam) error {
	payload := map[string]any{"nshttpparam": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsHTTPParamURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSHTTPParam() ([]models.NSHTTPParam, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsHTTPParamURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSHTTPParam `json:"nshttpparam"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) CountNSHTTPParam() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsHTTPParamURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSHTTPParam `json:"nshttpparam"`
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

// nshttpprofile
func (s *NSService) AddNSHTTPProfile(resource models.NSHTTPProfile) error {
	payload := map[string]any{"nshttpprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nsHTTPProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DeleteNSHTTPProfile(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nsHTTPProfileURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UpdateNSHTTPProfile(resource models.NSHTTPProfile) error {
	payload := map[string]any{"nshttpprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsHTTPProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UnsetNSHTTPProfile(resource models.NSHTTPProfile) error {
	payload := map[string]any{"nshttpprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", nsHTTPProfileURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSHTTPProfile() ([]models.NSHTTPProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsHTTPProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSHTTPProfile `json:"nshttpprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSHTTPProfile(name string) (*models.NSHTTPProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsHTTPProfileURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSHTTPProfile `json:"nshttpprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("nshttpprofile %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *NSService) CountNSHTTPProfile() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsHTTPProfileURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSHTTPProfile `json:"nshttpprofile"`
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

// nsicapprofile
func (s *NSService) AddNSICAPProfile(resource models.NSICAPProfile) error {
	payload := map[string]any{"nsicapprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nsICAPProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DeleteNSICAPProfile(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nsICAPProfileURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UpdateNSICAPProfile(resource models.NSICAPProfile) error {
	payload := map[string]any{"nsicapprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsICAPProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UnsetNSICAPProfile(resource models.NSICAPProfile) error {
	payload := map[string]any{"nsicapprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", nsICAPProfileURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSICAPProfile() ([]models.NSICAPProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsICAPProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSICAPProfile `json:"nsicapprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSICAPProfile(name string) (*models.NSICAPProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsICAPProfileURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSICAPProfile `json:"nsicapprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("nsicapprofile %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *NSService) CountNSICAPProfile() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsICAPProfileURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSICAPProfile `json:"nsicapprofile"`
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

// nsip
func (s *NSService) AddNSIP(resource models.NSIP) error {
	payload := map[string]any{"nsip": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nsIPURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DeleteNSIP(ipaddress string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nsIPURL, ipaddress), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UpdateNSIP(resource models.NSIP) error {
	payload := map[string]any{"nsip": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsIPURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UnsetNSIP(resource models.NSIP) error {
	payload := map[string]any{"nsip": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", nsIPURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) EnableNSIP(resource models.NSIP) error {
	payload := map[string]any{"nsip": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=enable", nsIPURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DisableNSIP(resource models.NSIP) error {
	payload := map[string]any{"nsip": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=disable", nsIPURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSIP() ([]models.NSIP, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsIPURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSIP `json:"nsip"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) CountNSIP() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsIPURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSIP `json:"nsip"`
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

// nsip6
func (s *NSService) AddNSIP6(resource models.NSIP6) error {
	payload := map[string]any{"nsip6": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nsIP6URL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DeleteNSIP6(ipv6address string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nsIP6URL, ipv6address), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UpdateNSIP6(resource models.NSIP6) error {
	payload := map[string]any{"nsip6": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsIP6URL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UnsetNSIP6(resource models.NSIP6) error {
	payload := map[string]any{"nsip6": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", nsIP6URL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSIP6() ([]models.NSIP6, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsIP6URL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSIP6 `json:"nsip6"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) CountNSIP6() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsIP6URL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSIP6 `json:"nsip6"`
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

// nslicense
func (s *NSService) GetAllNSLicense() ([]models.NSLicense, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsLicenseURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSLicense `json:"nslicense"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// nslicenseproxyserver
func (s *NSService) AddNSLicenseProxyServer(resource models.NSLicenseProxyServer) error {
	payload := map[string]any{"nslicenseproxyserver": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nsLicenseProxyServerURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DeleteNSLicenseProxyServer(serverip string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nsLicenseProxyServerURL, serverip), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UpdateNSLicenseProxyServer(resource models.NSLicenseProxyServer) error {
	payload := map[string]any{"nslicenseproxyserver": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsLicenseProxyServerURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSLicenseProxyServer() ([]models.NSLicenseProxyServer, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsLicenseProxyServerURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSLicenseProxyServer `json:"nslicenseproxyserver"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) CountNSLicenseProxyServer() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsLicenseProxyServerURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSLicenseProxyServer `json:"nslicenseproxyserver"`
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

// nslicenseserver
func (s *NSService) AddNSLicenseServer(resource models.NSLicenseServer) error {
	payload := map[string]any{"nslicenseserver": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nsLicenseServerURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DeleteNSLicenseServer(servername string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nsLicenseServerURL, servername), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UpdateNSLicenseServer(resource models.NSLicenseServer) error {
	payload := map[string]any{"nslicenseserver": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsLicenseServerURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSLicenseServer() ([]models.NSLicenseServer, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsLicenseServerURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSLicenseServer `json:"nslicenseserver"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) CountNSLicenseServer() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsLicenseServerURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSLicenseServer `json:"nslicenseserver"`
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

// nslicenseserverpool
func (s *NSService) GetAllNSLicenseServerPool() ([]models.NSLicenseServerPool, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsLicenseServerPoolURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSLicenseServerPool `json:"nslicenseserverpool"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// nslimitidentifier
func (s *NSService) AddNSLimitIdentifier(resource models.NSLimitIdentifier) error {
	payload := map[string]any{"nslimitidentifier": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nsLimitIdentifierURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DeleteNSLimitIdentifier(limitidentifier string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nsLimitIdentifierURL, limitidentifier), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UpdateNSLimitIdentifier(resource models.NSLimitIdentifier) error {
	payload := map[string]any{"nslimitidentifier": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsLimitIdentifierURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UnsetNSLimitIdentifier(resource models.NSLimitIdentifier) error {
	payload := map[string]any{"nslimitidentifier": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", nsLimitIdentifierURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSLimitIdentifier() ([]models.NSLimitIdentifier, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsLimitIdentifierURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSLimitIdentifier `json:"nslimitidentifier"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSLimitIdentifier(limitidentifier string) (*models.NSLimitIdentifier, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsLimitIdentifierURL, limitidentifier), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSLimitIdentifier `json:"nslimitidentifier"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("nslimitidentifier %s not found", limitidentifier)
	}

	return &result.Data[0], nil
}

func (s *NSService) CountNSLimitIdentifier() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsLimitIdentifierURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSLimitIdentifier `json:"nslimitidentifier"`
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

// nslimitidentifier_binding
func (s *NSService) GetAllNSLimitIdentifierBinding() ([]models.NSLimitIdentifierBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsLimitIdentifierBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSLimitIdentifierBinding `json:"nslimitidentifier_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSLimitIdentifierBinding(limitidentifier string) (*models.NSLimitIdentifierBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsLimitIdentifierBindingURL, limitidentifier), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSLimitIdentifierBinding `json:"nslimitidentifier_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("nslimitidentifier_binding %s not found", limitidentifier)
	}

	return &result.Data[0], nil
}

// nslimitidentifier_nslimitsessions_binding
func (s *NSService) GetAllNSLimitIdentifierNSLimitSessionsBinding() ([]models.NSLimitIdentifierNSLimitSessionsBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsLimitIdentifierNSLimitSessionsBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSLimitIdentifierNSLimitSessionsBinding `json:"nslimitidentifier_nslimitsessions_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSLimitIdentifierNSLimitSessionsBinding(limitidentifier string) ([]models.NSLimitIdentifierNSLimitSessionsBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsLimitIdentifierNSLimitSessionsBindingURL, limitidentifier), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSLimitIdentifierNSLimitSessionsBinding `json:"nslimitidentifier_nslimitsessions_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) CountNSLimitIdentifierNSLimitSessionsBinding(limitidentifier string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", nsLimitIdentifierNSLimitSessionsBindingURL, limitidentifier), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSLimitIdentifierNSLimitSessionsBinding `json:"nslimitidentifier_nslimitsessions_binding"`
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

// nslimitselector
func (s *NSService) AddNSLimitSelector(resource models.NSLimitSelector) error {
	payload := map[string]any{"nslimitselector": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nsLimitSelectorURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DeleteNSLimitSelector(selectorname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nsLimitSelectorURL, selectorname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UpdateNSLimitSelector(resource models.NSLimitSelector) error {
	payload := map[string]any{"nslimitselector": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsLimitSelectorURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UnsetNSLimitSelector(resource models.NSLimitSelector) error {
	payload := map[string]any{"nslimitselector": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", nsLimitSelectorURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSLimitSelector() ([]models.NSLimitSelector, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsLimitSelectorURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSLimitSelector `json:"nslimitselector"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSLimitSelector(selectorname string) (*models.NSLimitSelector, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsLimitSelectorURL, selectorname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSLimitSelector `json:"nslimitselector"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("nslimitselector %s not found", selectorname)
	}

	return &result.Data[0], nil
}

func (s *NSService) CountNSLimitSelector() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsLimitSelectorURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSLimitSelector `json:"nslimitselector"`
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

// nslimitsessions
func (s *NSService) GetAllNSLimitSessions(args map[string]string) ([]models.NSLimitSessions, error) {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", nsLimitSessionsURL, argsStr), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSLimitSessions `json:"nslimitsessions"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) CountNSLimitSessions(args map[string]string) (int, error) {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",") + "&count=yes"
	} else {
		argsStr = "?count=yes"
	}
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", nsLimitSessionsURL, argsStr), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSLimitSessions `json:"nslimitsessions"`
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

func (s *NSService) ClearNSLimitSessions(resource models.NSLimitSessions) error {
	payload := map[string]any{"nslimitsessions": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=clear", nsLimitSessionsURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// nsmigration
func (s *NSService) GetAllNSMigration() ([]models.NSMigration, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsMigrationURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSMigration `json:"nsmigration"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) CountNSMigration() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsMigrationURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSMigration `json:"nsmigration"`
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

// nsmode
func (s *NSService) EnableNSMode(resource models.NSMode) error {
	payload := map[string]any{"nsmode": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=enable", nsModeURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DisableNSMode(resource models.NSMode) error {
	payload := map[string]any{"nsmode": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=disable", nsModeURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSMode() ([]models.NSMode, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsModeURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSMode `json:"nsmode"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// nsparam
func (s *NSService) UpdateNSParam(resource models.NSParam) error {
	payload := map[string]any{"nsparam": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsParamURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UnsetNSParam(resource models.NSParam) error {
	payload := map[string]any{"nsparam": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", nsParamURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSParam() ([]models.NSParam, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsParamURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSParam `json:"nsparam"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// nspartition
func (s *NSService) AddNSPartition(resource models.NSPartition) error {
	payload := map[string]any{"nspartition": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nsPartitionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DeleteNSPartition(partitionname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nsPartitionURL, partitionname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UpdateNSPartition(resource models.NSPartition) error {
	payload := map[string]any{"nspartition": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsPartitionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UnsetNSPartition(resource models.NSPartition) error {
	payload := map[string]any{"nspartition": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", nsPartitionURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) SwitchNSPartition(resource models.NSPartition) error {
	payload := map[string]any{"nspartition": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=switch", nsPartitionURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSPartition() ([]models.NSPartition, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsPartitionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSPartition `json:"nspartition"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSPartition(partitionname string) (*models.NSPartition, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsPartitionURL, partitionname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSPartition `json:"nspartition"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("nspartition %s not found", partitionname)
	}

	return &result.Data[0], nil
}

func (s *NSService) CountNSPartition() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsPartitionURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSPartition `json:"nspartition"`
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

// nspartitionmac
func (s *NSService) GetAllNSPartitionMAC() ([]models.NSPartitionMac, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsPartitionMACURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSPartitionMac `json:"nspartitionmac"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) CountNSPartitionMAC() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsPartitionMACURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSPartitionMac `json:"nspartitionmac"`
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

// nspartition_binding
func (s *NSService) GetAllNSPartitionBinding() ([]models.NSPartitionBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsPartitionBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSPartitionBinding `json:"nspartition_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSPartitionBinding(partitionname string) (*models.NSPartitionBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsPartitionBindingURL, partitionname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSPartitionBinding `json:"nspartition_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("nspartition_binding %s not found", partitionname)
	}

	return &result.Data[0], nil
}

// nspartition_bridgegroup_binding
func (s *NSService) AddNSPartitionBridgeGroupBinding(resource models.NSPartitionBridgeGroupBinding) error {
	payload := map[string]any{"nspartition_bridgegroup_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nsPartitionBridgeGroupBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DeleteNSPartitionBridgeGroupBinding(partitionname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nsPartitionBridgeGroupBindingURL, partitionname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSPartitionBridgeGroupBinding() ([]models.NSPartitionBridgeGroupBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsPartitionBridgeGroupBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSPartitionBridgeGroupBinding `json:"nspartition_bridgegroup_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSPartitionBridgeGroupBinding(partitionname string) ([]models.NSPartitionBridgeGroupBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsPartitionBridgeGroupBindingURL, partitionname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSPartitionBridgeGroupBinding `json:"nspartition_bridgegroup_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) CountNSPartitionBridgeGroupBinding(partitionname string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", nsPartitionBridgeGroupBindingURL, partitionname), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSPartitionBridgeGroupBinding `json:"nspartition_bridgegroup_binding"`
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

// nspartition_vlan_binding
func (s *NSService) AddNSPartitionVLANBinding(resource models.NSPartitionVlanBinding) error {
	payload := map[string]any{"nspartition_vlan_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nsPartitionVlanBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DeleteNSPartitionVLANBinding(partitionname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nsPartitionVlanBindingURL, partitionname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSPartitionVLANBinding() ([]models.NSPartitionVlanBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsPartitionVlanBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSPartitionVlanBinding `json:"nspartition_vlan_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSPartitionVLANBinding(partitionname string) ([]models.NSPartitionVlanBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsPartitionVlanBindingURL, partitionname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSPartitionVlanBinding `json:"nspartition_vlan_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) CountNSPartitionVLANBinding(partitionname string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", nsPartitionVlanBindingURL, partitionname), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSPartitionVlanBinding `json:"nspartition_vlan_binding"`
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

// nspartition_vxlan_binding
func (s *NSService) AddNSPartitionVXLANBinding(resource models.NSPartitionVXLANBinding) error {
	payload := map[string]any{"nspartition_vxlan_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nsPartitionVXLANBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DeleteNSPartitionVXLANBinding(partitionname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nsPartitionVXLANBindingURL, partitionname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSPartitionVXLANBinding() ([]models.NSPartitionVXLANBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsPartitionVXLANBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSPartitionVXLANBinding `json:"nspartition_vxlan_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSPartitionVXLANBinding(partitionname string) ([]models.NSPartitionVXLANBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsPartitionVXLANBindingURL, partitionname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSPartitionVXLANBinding `json:"nspartition_vxlan_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) CountNSPartitionVXLANBinding(partitionname string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", nsPartitionVXLANBindingURL, partitionname), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSPartitionVXLANBinding `json:"nspartition_vxlan_binding"`
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

// nspbr
func (s *NSService) AddNSPBR(resource models.NSPBR) error {
	payload := map[string]any{"nspbr": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nsPBRURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DeleteNSPBR(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nsPBRURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UpdateNSPBR(resource models.NSPBR) error {
	payload := map[string]any{"nspbr": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsPBRURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UnsetNSPBR(resource models.NSPBR) error {
	payload := map[string]any{"nspbr": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", nsPBRURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) EnableNSPBR(resource models.NSPBR) error {
	payload := map[string]any{"nspbr": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=enable", nsPBRURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DisableNSPBR(resource models.NSPBR) error {
	payload := map[string]any{"nspbr": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=disable", nsPBRURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSPBR() ([]models.NSPBR, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsPBRURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSPBR `json:"nspbr"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSPBR(name string) (*models.NSPBR, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsPBRURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSPBR `json:"nspbr"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("nspbr %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *NSService) CountNSPBR() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsPBRURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSPBR `json:"nspbr"`
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

// nspbr6
func (s *NSService) AddNSPBR6(resource models.NSPBR6) error {
	payload := map[string]any{"nspbr6": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nsPBR6URL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DeleteNSPBR6(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nsPBR6URL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UpdateNSPBR6(resource models.NSPBR6) error {
	payload := map[string]any{"nspbr6": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsPBR6URL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UnsetNSPBR6(resource models.NSPBR6) error {
	payload := map[string]any{"nspbr6": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", nsPBR6URL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) RenumberNSPBR6(resource models.NSPBR6) error {
	payload := map[string]any{"nspbr6": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=renumber", nsPBR6URL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) EnableNSPBR6(resource models.NSPBR6) error {
	payload := map[string]any{"nspbr6": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=enable", nsPBR6URL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DisableNSPBR6(resource models.NSPBR6) error {
	payload := map[string]any{"nspbr6": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=disable", nsPBR6URL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSPBR6() ([]models.NSPBR6, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsPBR6URL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSPBR6 `json:"nspbr6"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSPBR6(name string) (*models.NSPBR6, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsPBR6URL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSPBR6 `json:"nspbr6"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("nspbr6 %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *NSService) CountNSPBR6() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsPBR6URL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSPBR6 `json:"nspbr6"`
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

func (s *NSService) ClearNSPBR6(resource models.NSPBR6) error {
	payload := map[string]any{"nspbr6": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=clear", nsPBR6URL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) ApplyNSPBR6(resource models.NSPBR6) error {
	payload := map[string]any{"nspbr6": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=apply", nsPBR6URL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// nspbrs
func (s *NSService) RenumberNSPBRs() error {
	payload := map[string]any{"nspbrs": map[string]any{}}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=renumber", nsPBRsURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) ClearNSPBRs() error {
	payload := map[string]any{"nspbrs": map[string]any{}}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=clear", nsPBRsURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) ApplyNSPBRs() error {
	payload := map[string]any{"nspbrs": map[string]any{}}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=apply", nsPBRsURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// nsratecontrol
func (s *NSService) UpdateNSrateControl(resource models.NSRateControl) error {
	payload := map[string]any{"nsratecontrol": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsRateControlURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UnsetNSrateControl(resource models.NSRateControl) error {
	payload := map[string]any{"nsratecontrol": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", nsRateControlURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSrateControl() ([]models.NSRateControl, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsRateControlURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSRateControl `json:"nsratecontrol"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// nsrollbackcmd
func (s *NSService) GetAllNSRollBackCMD() ([]models.NSRollbackCmd, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsRollBackCMDURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSRollbackCmd `json:"nsrollbackcmd"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// nsrpcnode
func (s *NSService) UpdateNSRPCNode(resource models.NSRPCNode) error {
	payload := map[string]any{"nsrpcnode": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsRPCNodeURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UnsetNSRPCNode(resource models.NSRPCNode) error {
	payload := map[string]any{"nsrpcnode": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", nsRPCNodeURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSRPCNode() ([]models.NSRPCNode, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsRPCNodeURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSRPCNode `json:"nsrpcnode"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSRPCNode(ipaddress string) (*models.NSRPCNode, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsRPCNodeURL, ipaddress), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSRPCNode `json:"nsrpcnode"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("nsrpcnode %s not found", ipaddress)
	}

	return &result.Data[0], nil
}

func (s *NSService) CountNSRPCNode() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsRPCNodeURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSRPCNode `json:"nsrpcnode"`
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

// nsrunningconfig
func (s *NSService) GetAllNSRunningConfig() ([]models.NSRunningConfig, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsRunningConfigURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSRunningConfig `json:"nsrunningconfig"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// nssavedconfig
func (s *NSService) GetAllNSSavedConfig() ([]models.NSSavedConfig, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsSavedConfigURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSSavedConfig `json:"nssavedconfig"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// nsservicefunction
func (s *NSService) AddNSServiceFunction(resource models.NSServiceFunction) error {
	payload := map[string]any{"nsservicefunction": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nsServiceFunctionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DeleteNSServiceFunction(servicefunctionname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nsServiceFunctionURL, servicefunctionname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSServiceFunction() ([]models.NSServiceFunction, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsServiceFunctionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSServiceFunction `json:"nsservicefunction"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSServiceFunction(servicefunctionname string) (*models.NSServiceFunction, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsServiceFunctionURL, servicefunctionname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSServiceFunction `json:"nsservicefunction"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("nsservicefunction %s not found", servicefunctionname)
	}

	return &result.Data[0], nil
}

func (s *NSService) CountNSServiceFunction() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsServiceFunctionURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSServiceFunction `json:"nsservicefunction"`
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

// nsservicepath
func (s *NSService) AddNSServicePath(resource models.NSServicePath) error {
	payload := map[string]any{"nsservicepath": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nsServicePathURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DeleteNSServicePath(servicepathname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nsServicePathURL, servicepathname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSServicePath() ([]models.NSServicePath, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsServicePathURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSServicePath `json:"nsservicepath"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSServicePath(servicepathname string) (*models.NSServicePath, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsServicePathURL, servicepathname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSServicePath `json:"nsservicepath"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("nsservicepath %s not found", servicepathname)
	}

	return &result.Data[0], nil
}

func (s *NSService) CountNSServicePath() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsServicePathURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSServicePath `json:"nsservicepath"`
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

// nsservicepath_binding
func (s *NSService) GetAllNSServicePathBinding() ([]models.NSServicePathBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsServicePathBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSServicePathBinding `json:"nsservicepath_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSServicePathBinding(servicepathname string) (*models.NSServicePathBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsServicePathBindingURL, servicepathname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSServicePathBinding `json:"nsservicepath_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("nsservicepath_binding %s not found", servicepathname)
	}

	return &result.Data[0], nil
}

// nsservicepath_nsservicefunction_binding
func (s *NSService) AddNSServicePathNSServiceFunctionBinding(resource models.NSServicePathNSServiceFunctionBinding) error {
	payload := map[string]any{"nsservicepath_nsservicefunction_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nsServicePathNSServiceFunctionBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DeleteNSServicePathNSServiceFunctionBinding(servicepathname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nsServicePathNSServiceFunctionBindingURL, servicepathname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSServicePathNSServiceFunctionBinding() ([]models.NSServicePathNSServiceFunctionBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsServicePathNSServiceFunctionBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSServicePathNSServiceFunctionBinding `json:"nsservicepath_nsservicefunction_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSServicePathNSServiceFunctionBinding(servicepathname string) ([]models.NSServicePathNSServiceFunctionBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsServicePathNSServiceFunctionBindingURL, servicepathname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSServicePathNSServiceFunctionBinding `json:"nsservicepath_nsservicefunction_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) CountNSServicePathNSServiceFunctionBinding(servicepathname string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", nsServicePathNSServiceFunctionBindingURL, servicepathname), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSServicePathNSServiceFunctionBinding `json:"nsservicepath_nsservicefunction_binding"`
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

// nssimpleacl
func (s *NSService) AddNSSimpleACL(resource models.NSSimpleACL) error {
	payload := map[string]any{"nssimpleacl": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nsSimpleACLURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DeleteNSSimpleACL(aclname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nsSimpleACLURL, aclname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) FlushNSSimpleACL() error {
	payload := map[string]any{"nssimpleacl": map[string]any{}}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=flush", nsSimpleACLURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSSimpleACL() ([]models.NSSimpleACL, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsSimpleACLURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSSimpleACL `json:"nssimpleacl"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSSimpleACL(aclname string) (*models.NSSimpleACL, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsSimpleACLURL, aclname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSSimpleACL `json:"nssimpleacl"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("nssimpleacl %s not found", aclname)
	}

	return &result.Data[0], nil
}

func (s *NSService) CountNSSimpleACL() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsSimpleACLURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSSimpleACL `json:"nssimpleacl"`
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

func (s *NSService) ClearNSSimpleACL(resource models.NSSimpleACL) error {
	payload := map[string]any{"nssimpleacl": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=clear", nsSimpleACLURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// nssimpleacl6
func (s *NSService) AddNSSimpleACL6(resource models.NSSimpleACL6) error {
	payload := map[string]any{"nssimpleacl6": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nsSimpleACL6URL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DeleteNSSimpleACL6(aclname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nsSimpleACL6URL, aclname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) FlushNSSimpleACL6() error {
	payload := map[string]any{"nssimpleacl6": map[string]any{}}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=flush", nsSimpleACL6URL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSSimpleACL6() ([]models.NSSimpleACL6, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsSimpleACL6URL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSSimpleACL6 `json:"nssimpleacl6"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSSimpleACL6(aclname string) (*models.NSSimpleACL6, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsSimpleACL6URL, aclname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSSimpleACL6 `json:"nssimpleacl6"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("nssimpleacl6 %s not found", aclname)
	}

	return &result.Data[0], nil
}

func (s *NSService) CountNSSimpleACL6() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsSimpleACL6URL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSSimpleACL6 `json:"nssimpleacl6"`
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

func (s *NSService) ClearNSSimpleACL6(resource models.NSSimpleACL6) error {
	payload := map[string]any{"nssimpleacl6": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=clear", nsSimpleACL6URL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// nssourceroutecachetable
func (s *NSService) FlushNSSourceRouteCacheTable() error {
	payload := map[string]any{"nssourceroutecachetable": map[string]any{}}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=flush", nsSourceRouteCacheTableURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSSourceRouteCacheTable() ([]models.NSSourceRouteCacheTable, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsSourceRouteCacheTableURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSSourceRouteCacheTable `json:"nssourceroutecachetable"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) CountNSSourceRouteCacheTable() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsSourceRouteCacheTableURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSSourceRouteCacheTable `json:"nssourceroutecachetable"`
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

// nsspparams
func (s *NSService) UpdateNSSPParams(resource models.NSSPParams) error {
	payload := map[string]any{"nsspparams": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsSPParamsURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}
func (s *NSService) UnsetNSSPParams(resource models.NSSPParams) error {
	payload := map[string]any{"nsspparams": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", nsSPParamsURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSSPParams() ([]models.NSSPParams, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsSPParamsURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSSPParams `json:"nsspparams"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// nsstats
func (s *NSService) ClearNSStats(resource models.NSStats) error {
	payload := map[string]any{"nsstats": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=clear", nsStatsURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// nssurgeq
func (s *NSService) FlushNSSurgeQ() error {
	payload := map[string]any{"nssurgeq": map[string]any{}}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=flush", nsSurgeQURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// nstcpbufparam
func (s *NSService) UpdateNSTCPBufParam(resource models.NSTCPBufParam) error {
	payload := map[string]any{"nstcpbufparam": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsTCPBufParamURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UnsetNSTCPBufParam(resource models.NSTCPBufParam) error {
	payload := map[string]any{"nstcpbufparam": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", nsTCPBufParamURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSTCPBufParam() ([]models.NSTCPBufParam, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsTCPBufParamURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSTCPBufParam `json:"nstcpbufparam"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// nstcpparam
func (s *NSService) UpdateNSTCPParam(resource models.NSTCPParam) error {
	payload := map[string]any{"nstcpparam": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsTCPParamURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UnsetNSTCPParam(resource models.NSTCPParam) error {
	payload := map[string]any{"nstcpparam": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", nsTCPParamURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSTCPParam() ([]models.NSTCPParam, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsTCPParamURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSTCPParam `json:"nstcpparam"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// nstcpprofile
func (s *NSService) AddNSTCPProfile(resource models.NSTCPProfile) error {
	payload := map[string]any{"nstcpprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nsTCPProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DeleteNSTCPProfile(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nsTCPProfileURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UpdateNSTCPProfile(resource models.NSTCPProfile) error {
	payload := map[string]any{"nstcpprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsTCPProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UnsetNSTCPProfile(resource models.NSTCPProfile) error {
	payload := map[string]any{"nstcpprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", nsTCPProfileURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSTCPProfile() ([]models.NSTCPProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsTCPProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSTCPProfile `json:"nstcpprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSTCPProfile(name string) (*models.NSTCPProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsTCPProfileURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSTCPProfile `json:"nstcpprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("nstcpprofile %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *NSService) CountNSTCPProfile() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsTCPProfileURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSTCPProfile `json:"nstcpprofile"`
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

// nstimeout
func (s *NSService) UpdateNSTimeout(resource models.NSTimeout) error {
	payload := map[string]any{"nstimeout": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsTimeoutURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UnsetNSTimeout(resource models.NSTimeout) error {
	payload := map[string]any{"nstimeout": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", nsTimeoutURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSTimeout() ([]models.NSTimeout, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsTimeoutURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSTimeout `json:"nstimeout"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// nstimer
func (s *NSService) AddNSTimer(resource models.NSTimer) error {
	payload := map[string]any{"nstimer": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nsTimerURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DeleteNSTimer(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nsTimerURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UpdateNSTimer(resource models.NSTimer) error {
	payload := map[string]any{"nstimer": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsTimerURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UnsetNSTimer(resource models.NSTimer) error {
	payload := map[string]any{"nstimer": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", nsTimerURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSTimer() ([]models.NSTimer, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsTimerURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSTimer `json:"nstimer"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSTimer(name string) (*models.NSTimer, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsTimerURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSTimer `json:"nstimer"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("nstimer %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *NSService) CountNSTimer() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsTimerURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSTimer `json:"nstimer"`
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

func (s *NSService) RenameNSTimer(resource models.NSTimer) error {
	payload := map[string]any{"nstimer": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=rename", nsTimerURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// nstimer_autoscalepolicy_binding
func (s *NSService) AddNSTimerAutoScalePolicyBinding(resource models.NSTimerAutoScalePolicyBinding) error {
	payload := map[string]any{"nstimer_autoscalepolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nsTimerAutoScalePolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DeleteNSTimerAutoScalePolicyBinding(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nsTimerAutoScalePolicyBindingURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSTimerAutoScalePolicyBinding() ([]models.NSTimerAutoScalePolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsTimerAutoScalePolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSTimerAutoScalePolicyBinding `json:"nstimer_autoscalepolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSTimerAutoScalePolicyBinding(name string) ([]models.NSTimerAutoScalePolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsTimerAutoScalePolicyBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSTimerAutoScalePolicyBinding `json:"nstimer_autoscalepolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) CountNSTimerAutoScalePolicyBinding(name string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", nsTimerAutoScalePolicyBindingURL, name), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSTimerAutoScalePolicyBinding `json:"nstimer_autoscalepolicy_binding"`
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

// nstimer_binding
func (s *NSService) GetAllNSTimerBinding() ([]models.NSTimerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsTimerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSTimerBinding `json:"nstimer_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSTimerBinding(name string) (*models.NSTimerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsTimerBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSTimerBinding `json:"nstimer_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("nstimer_binding %s not found", name)
	}

	return &result.Data[0], nil
}

// nstimezone
func (s *NSService) GetAllNSTimeZone() ([]models.NSTimeZone, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsTimeZoneURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSTimeZone `json:"nstimezone"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSTimeZone() (*models.NSTimeZone, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsTimeZoneURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSTimeZone `json:"nstimezone"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("nstimezone not found")
	}

	return &result.Data[0], nil
}

func (s *NSService) CountNSTimeZone() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsTimeZoneURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSTimeZone `json:"nstimezone"`
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

// nstrafficdomain
func (s *NSService) AddNSTrafficDomain(resource models.NSTrafficDomain) error {
	payload := map[string]any{"nstrafficdomain": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nsTrafficDomainURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DeleteNSTrafficDomain(td string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nsTrafficDomainURL, td), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) EnableNSTrafficDomain(resource models.NSTrafficDomain) error {
	payload := map[string]any{"nstrafficdomain": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=enable", nsTrafficDomainURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DisableNSTrafficDomain(resource models.NSTrafficDomain) error {
	payload := map[string]any{"nstrafficdomain": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=disable", nsTrafficDomainURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSTrafficDomain() ([]models.NSTrafficDomain, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsTrafficDomainURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSTrafficDomain `json:"nstrafficdomain"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSTrafficDomain(td string) (*models.NSTrafficDomain, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsTrafficDomainURL, td), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSTrafficDomain `json:"nstrafficdomain"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("nstrafficdomain %s not found", td)
	}

	return &result.Data[0], nil
}

func (s *NSService) CountNSTrafficDomain() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsTrafficDomainURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSTrafficDomain `json:"nstrafficdomain"`
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

func (s *NSService) ClearNSTrafficDomain(resource models.NSTrafficDomain) error {
	payload := map[string]any{"nstrafficdomain": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=clear", nsTrafficDomainURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// nstrafficdomain_binding
func (s *NSService) GetAllNSTrafficDomainBinding() ([]models.NSTrafficDomainBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsTrafficDomainBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSTrafficDomainBinding `json:"nstrafficdomain_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSTrafficDomainBinding(td string) (*models.NSTrafficDomainBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsTrafficDomainBindingURL, td), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSTrafficDomainBinding `json:"nstrafficdomain_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("nstrafficdomain_binding %s not found", td)
	}

	return &result.Data[0], nil
}

// nstrafficdomain_bridgegroup_binding
func (s *NSService) AddNSTrafficDomainBridgeGroupBinding(resource models.NSTrafficDomainBridgeGroupBinding) error {
	payload := map[string]any{"nstrafficdomain_bridgegroup_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nsTrafficDomainBridgeGroupBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DeleteNSTrafficDomainBridgeGroupBinding(td string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nsTrafficDomainBridgeGroupBindingURL, td), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSTrafficDomainBridgeGroupBinding() ([]models.NSTrafficDomainBridgeGroupBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsTrafficDomainBridgeGroupBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSTrafficDomainBridgeGroupBinding `json:"nstrafficdomain_bridgegroup_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSTrafficDomainBridgeGroupBinding(td string) ([]models.NSTrafficDomainBridgeGroupBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsTrafficDomainBridgeGroupBindingURL, td), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSTrafficDomainBridgeGroupBinding `json:"nstrafficdomain_bridgegroup_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) CountNSTrafficDomainBridgeGroupBinding(td string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", nsTrafficDomainBridgeGroupBindingURL, td), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSTrafficDomainBridgeGroupBinding `json:"nstrafficdomain_bridgegroup_binding"`
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

// nstrafficdomain_vlan_binding
func (s *NSService) AddNSTrafficDomainVLANBinding(resource models.NSTrafficDomainVLANBinding) error {
	payload := map[string]any{"nstrafficdomain_vlan_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nsTrafficDomainVLANBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DeleteNSTrafficDomainVLANBinding(td string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nsTrafficDomainVLANBindingURL, td), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSTrafficDomainVLANBinding() ([]models.NSTrafficDomainVLANBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsTrafficDomainVLANBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSTrafficDomainVLANBinding `json:"nstrafficdomain_vlan_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSTrafficDomainVLANBinding(td string) ([]models.NSTrafficDomainVLANBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsTrafficDomainVLANBindingURL, td), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSTrafficDomainVLANBinding `json:"nstrafficdomain_vlan_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) CountNSTrafficDomainVLANBinding(td string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", nsTrafficDomainVLANBindingURL, td), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSTrafficDomainVLANBinding `json:"nstrafficdomain_vlan_binding"`
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

// nstrafficdomain_vxlan_binding
func (s *NSService) AddNSTrafficDomainVXLANBinding(resource models.NSTrafficDomainVXLANBinding) error {
	payload := map[string]any{"nstrafficdomain_vxlan_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nsTrafficDomainVXLANBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DeleteNSTrafficDomainVXLANBinding(td string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nsTrafficDomainVXLANBindingURL, td), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSTrafficDomainVXLANBinding() ([]models.NSTrafficDomainVXLANBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsTrafficDomainVXLANBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSTrafficDomainVXLANBinding `json:"nstrafficdomain_vxlan_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSTrafficDomainVXLANBinding(td string) ([]models.NSTrafficDomainVXLANBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsTrafficDomainVXLANBindingURL, td), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSTrafficDomainVXLANBinding `json:"nstrafficdomain_vxlan_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) CountNSTrafficDomainVXLANBinding(td string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", nsTrafficDomainVXLANBindingURL, td), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSTrafficDomainVXLANBinding `json:"nstrafficdomain_vxlan_binding"`
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

// nsvariable
func (s *NSService) AddNSVariable(resource models.NSVariable) error {
	payload := map[string]any{"nsvariable": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nsVariableURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DeleteNSVariable(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nsVariableURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UpdateNSVariable(resource models.NSVariable) error {
	payload := map[string]any{"nsvariable": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsVariableURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UnsetNSVariable(resource models.NSVariable) error {
	payload := map[string]any{"nsvariable": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", nsVariableURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSVariable() ([]models.NSVariable, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsVariableURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSVariable `json:"nsvariable"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSVariable(name string) (*models.NSVariable, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsVariableURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSVariable `json:"nsvariable"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("nsvariable %s not found", name)
	}

	return &result.Data[0], nil
}

func (s *NSService) CountNSVariable() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsVariableURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSVariable `json:"nsvariable"`
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

// nsversion
func (s *NSService) GetAllNSVersion() ([]models.NSVersion, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsVersionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSVersion `json:"nsversion"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// nsvpxparam
func (s *NSService) UpdateNSVPXParam(resource models.NSVPXParam) error {
	payload := map[string]any{"nsvpxparam": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsVPXParamURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UnsetNSVPXParam(resource models.NSVPXParam) error {
	payload := map[string]any{"nsvpxparam": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", nsVPXParamURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSVPXParam() ([]models.NSVPXParam, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsVPXParamURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSVPXParam `json:"nsvpxparam"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) CountNSVPXParam() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsVPXParamURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSVPXParam `json:"nsvpxparam"`
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

// nsweblogparam
func (s *NSService) UpdateNSWebLogParam(resource models.NSWebLogParam) error {
	payload := map[string]any{"nsweblogparam": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsWeblogParamURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UnsetNSWebLogParam(resource models.NSWebLogParam) error {
	payload := map[string]any{"nsweblogparam": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", nsWeblogParamURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSWebLogParam() ([]models.NSWebLogParam, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsWeblogParamURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSWebLogParam `json:"nsweblogparam"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// nsxmlnamespace
func (s *NSService) AddNSXMLNamespace(resource models.NSXMLNamespace) error {
	payload := map[string]any{"nsxmlnamespace": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nsXMLNamespaceURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) DeleteNSXMLNamespace(prefix string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nsXMLNamespaceURL, prefix), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UpdateNSXMLNamespace(resource models.NSXMLNamespace) error {
	payload := map[string]any{"nsxmlnamespace": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nsXMLNamespaceURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) UnsetNSXMLNamespace(resource models.NSXMLNamespace) error {
	payload := map[string]any{"nsxmlnamespace": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", nsXMLNamespaceURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NSService) GetAllNSXMLNamespace() ([]models.NSXMLNamespace, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsXMLNamespaceURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSXMLNamespace `json:"nsxmlnamespace"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NSService) GetNSXMLNamespace(prefix string) (*models.NSXMLNamespace, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nsXMLNamespaceURL, prefix), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NSXMLNamespace `json:"nsxmlnamespace"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("nsxmlnamespace %s not found", prefix)
	}

	return &result.Data[0], nil
}

func (s *NSService) CountNSXMLNamespace() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nsXMLNamespaceURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NSXMLNamespace `json:"nsxmlnamespace"`
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

// reboot
func (s *NSService) Reboot() error {
	payload := map[string]any{"reboot": map[string]any{}}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, rebootURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// shutdown
func (s *NSService) Shutdown() error {
	payload := map[string]any{"shutdown": map[string]any{}}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, shutdownURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}
