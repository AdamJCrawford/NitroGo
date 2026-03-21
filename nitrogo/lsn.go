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
	lsnAppsAttributesURL                      = "/nitro/v1/config/lsnappsattributes"
	lsnAppsProfileURL                         = "/nitro/v1/config/lsnappsprofile"
	lsnAppsProfileBindingURL                  = "/nitro/v1/config/lsnappsprofile_binding"
	lsnAppsProfileLSNAppsAttributesBindingURL = "/nitro/v1/config/lsnappsprofile_lsnappsattributes_binding"
	lsnAppsProfilePortBindingURL              = "/nitro/v1/config/lsnappsprofile_port_binding"
	lsnClientURL                              = "/nitro/v1/config/lsnclient"
	lsnClientBindingURL                       = "/nitro/v1/config/lsnclient_binding"
	lsnClientNetwork6BindingURL               = "/nitro/v1/config/lsnclient_network6_binding"
	lsnClientNetworkBindingURL                = "/nitro/v1/config/lsnclient_network_binding"
	lsnClientNSACL6BindingURL                 = "/nitro/v1/config/lsnclient_nsacl6_binding"
	lsnClientNSACLBindingURL                  = "/nitro/v1/config/lsnclient_nsacl_binding"
	lsnDeterministicNATURL                    = "/nitro/v1/config/lsndeterministicnat"
	lsnGroupURL                               = "/nitro/v1/config/lsngroup"
	lsnGroupBindingURL                        = "/nitro/v1/config/lsngroup_binding"
	lsnGroupIPSecALGProfileBindingURL         = "/nitro/v1/config/lsngroup_ipsecalgprofile_binding"
	lsnGroupLSNAppsProfileBindingURL          = "/nitro/v1/config/lsngroup_lsnappsprofile_binding"
	lsnGroupLSNHTTPHDRLogProfileBindingURL    = "/nitro/v1/config/lsngroup_lsnhttphdrlogprofile_binding"
	lsnGroupLSNLogProfileBindingURL           = "/nitro/v1/config/lsngroup_lsnlogprofile_binding"
	lsnGroupLSNPoolBindingURL                 = "/nitro/v1/config/lsngroup_lsnpool_binding"
	lsnGroupLSNRTSPALGProfileBindingURL       = "/nitro/v1/config/lsngroup_lsnrtspalgprofile_binding"
	lsnGroupLSNSIPALGProfileBindingURL        = "/nitro/v1/config/lsngroup_lsnsipalgprofile_binding"
	lsnGroupLSNTransportProfileBindingURL     = "/nitro/v1/config/lsngroup_lsntransportprofile_binding"
	lsnGroupPCPServerBindingURL               = "/nitro/v1/config/lsngroup_pcpserver_binding"
	lsnHTTPHDRLogProfileURL                   = "/nitro/v1/config/lsnhttphdrlogprofile"
	lsnIP6ProfileURL                          = "/nitro/v1/config/lsnip6profile"
	lsnLogProfileURL                          = "/nitro/v1/config/lsnlogprofile"
	lsnParameterURL                           = "/nitro/v1/config/lsnparameter"
	lsnPoolURL                                = "/nitro/v1/config/lsnpool"
	lsnPoolBindingURL                         = "/nitro/v1/config/lsnpool_binding"
	lsnPoolLSNIPBindingURL                    = "/nitro/v1/config/lsnpool_lsnip_binding"
	lsnRTSPALGProfileURL                      = "/nitro/v1/config/lsnrtspalgprofile"
	lsnRTSPALGSessionURL                      = "/nitro/v1/config/lsnrtspalgsession"
	lsnRTSPALGSessionBindingURL               = "/nitro/v1/config/lsnrtspalgsession_binding"
	lsnRTSPALGSessionDataChannelBindingURL    = "/nitro/v1/config/lsnrtspalgsession_datachannel_binding"
	lsnSessionURL                             = "/nitro/v1/config/lsnsession"
	lsnSIPALGCallURL                          = "/nitro/v1/config/lsnsipalgcall"
	lsnSIPALGCallBindingURL                   = "/nitro/v1/config/lsnsipalgcall_binding"
	lsnSIPALGCallControlChannelBindingURL     = "/nitro/v1/config/lsnsipalgcall_controlchannel_binding"
	lsnSIPALGCallDataChannelBindingURL        = "/nitro/v1/config/lsnsipalgcall_datachannel_binding"
	lsnSIPALGProfileURL                       = "/nitro/v1/config/lsnsipalgprofile"
	lsnStaticURL                              = "/nitro/v1/config/lsnstatic"
	lsnTransportProfileURL                    = "/nitro/v1/config/lsntransportprofile"
)

// Large Scale NAT commands
type LSNService struct {
	client *Client
}

// lsnappsattributes
func (s *LSNService) AddLSNAppsAttributes(resource models.LSNAppsAttributes) error {
	payload := map[string]any{"lsnappsattributes": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lsnAppsAttributesURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) DeleteLSNAppsAttributes(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", lsnAppsAttributesURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) UpdateLSNAppsAttributes(resource models.LSNAppsAttributes) error {
	payload := map[string]any{"lsnappsattributes": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, lsnAppsAttributesURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) UnsetLSNAppsAttributes(resource models.LSNAppsAttributes) error {
	payload := map[string]any{"lsnappsattributes": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", lsnAppsAttributesURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) GetAllLSNAppsAttributes() ([]models.LSNAppsAttributes, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnAppsAttributesURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNAppsAttributes `json:"lsnappsattributes"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNAppsAttributes(name string) (models.LSNAppsAttributes, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnAppsAttributesURL, name), nil)
	if err != nil {
		return models.LSNAppsAttributes{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.LSNAppsAttributes{}, err
	}

	var result struct {
		Data []models.LSNAppsAttributes `json:"lsnappsattributes"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.LSNAppsAttributes{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.LSNAppsAttributes{}, fmt.Errorf("lsnappsattributes %s not found", name)
	}

	return result.Data[0], nil
}

func (s *LSNService) CountLSNAppsAttributes() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", lsnAppsAttributesURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LSNAppsAttributes `json:"lsnappsattributes"`
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

// lsnappsprofile
func (s *LSNService) AddLSNAppsProfile(resource models.LSNAppsProfile) error {
	payload := map[string]any{"lsnappsprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lsnAppsProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) DeleteLSNAppsProfile(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", lsnAppsProfileURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) UpdateLSNAppsProfile(resource models.LSNAppsProfile) error {
	payload := map[string]any{"lsnappsprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, lsnAppsProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) UnsetLSNAppsProfile(resource models.LSNAppsProfile) error {
	payload := map[string]any{"lsnappsprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", lsnAppsProfileURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) GetAllLSNAppsProfile() ([]models.LSNAppsProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnAppsProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNAppsProfile `json:"lsnappsprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNAppsProfile(name string) (models.LSNAppsProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnAppsProfileURL, name), nil)
	if err != nil {
		return models.LSNAppsProfile{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.LSNAppsProfile{}, err
	}

	var result struct {
		Data []models.LSNAppsProfile `json:"lsnappsprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.LSNAppsProfile{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.LSNAppsProfile{}, fmt.Errorf("lsnappsprofile %s not found", name)
	}

	return result.Data[0], nil
}

func (s *LSNService) CountLSNAppsProfile() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", lsnAppsProfileURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LSNAppsProfile `json:"lsnappsprofile"`
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

// lsnappsprofile_binding
func (s *LSNService) GetAllLSNAppsProfileBinding() ([]models.LSNAppsProfileBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnAppsProfileBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNAppsProfileBinding `json:"lsnappsprofile_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNAppsProfileBinding(name string) (models.LSNAppsProfileBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnAppsProfileBindingURL, name), nil)
	if err != nil {
		return models.LSNAppsProfileBinding{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.LSNAppsProfileBinding{}, err
	}

	var result struct {
		Data []models.LSNAppsProfileBinding `json:"lsnappsprofile_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.LSNAppsProfileBinding{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.LSNAppsProfileBinding{}, fmt.Errorf("lsnappsprofile_binding %s not found", name)
	}

	return result.Data[0], nil
}

// lsnappsprofile_lsnappsattributes_binding
func (s *LSNService) AddLSNAppsProfileLSNAppsAttributesBinding(resource models.LSNAppsProfileLSNAppsAttributesBinding) error {
	payload := map[string]any{"lsnappsprofile_lsnappsattributes_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lsnAppsProfileLSNAppsAttributesBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) DeleteLSNAppsProfileLSNAppsAttributesBinding(name string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lsnAppsProfileLSNAppsAttributesBindingURL, name, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) GetAllLSNAppsProfileLSNAppsAttributesBinding() ([]models.LSNAppsProfileLSNAppsAttributesBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnAppsProfileLSNAppsAttributesBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNAppsProfileLSNAppsAttributesBinding `json:"lsnappsprofile_lsnappsattributes_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNAppsProfileLSNAppsAttributesBinding(name string) ([]models.LSNAppsProfileLSNAppsAttributesBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnAppsProfileLSNAppsAttributesBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNAppsProfileLSNAppsAttributesBinding `json:"lsnappsprofile_lsnappsattributes_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) CountLSNAppsProfileLSNAppsAttributesBinding(appsprofilename string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lsnAppsProfileLSNAppsAttributesBindingURL, appsprofilename), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LSNAppsProfileLSNAppsAttributesBinding `json:"lsnappsprofile_lsnappsattributes_binding"`
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

// lsnappsprofile_port_binding
func (s *LSNService) AddLSNAppsProfilePortBinding(resource models.LSNAppsProfilePortBinding) error {
	payload := map[string]any{"lsnappsprofile_port_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lsnAppsProfilePortBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) DeleteLSNAppsProfilePortBinding(name string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lsnAppsProfilePortBindingURL, name, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) GetAllLSNAppsProfilePortBinding() ([]models.LSNAppsProfilePortBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnAppsProfilePortBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNAppsProfilePortBinding `json:"lsnappsprofile_port_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNAppsProfilePortBinding(name string) ([]models.LSNAppsProfilePortBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnAppsProfilePortBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNAppsProfilePortBinding `json:"lsnappsprofile_port_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) CountLSNAppsProfilePortBinding(appsprofilename string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lsnAppsProfilePortBindingURL, appsprofilename), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LSNAppsProfilePortBinding `json:"lsnappsprofile_port_binding"`
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

// lsnclient
func (s *LSNService) AddLSNClient(resource models.LSNClient) error {
	payload := map[string]any{"lsnclient": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lsnClientURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) DeleteLSNClient(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", lsnClientURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) GetAllLSNClient() ([]models.LSNClient, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnClientURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNClient `json:"lsnclient"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNClient(name string) (models.LSNClient, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnClientURL, name), nil)
	if err != nil {
		return models.LSNClient{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.LSNClient{}, err
	}

	var result struct {
		Data []models.LSNClient `json:"lsnclient"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.LSNClient{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.LSNClient{}, fmt.Errorf("lsnclient %s not found", name)
	}

	return result.Data[0], nil
}

func (s *LSNService) CountLSNClient() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", lsnClientURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LSNClient `json:"lsnclient"`
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

// lsnclient_binding
func (s *LSNService) GetAllLSNClientBinding() ([]models.LSNClientBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnClientBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNClientBinding `json:"lsnclient_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNClientBinding(name string) (models.LSNClientBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnClientBindingURL, name), nil)
	if err != nil {
		return models.LSNClientBinding{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.LSNClientBinding{}, err
	}

	var result struct {
		Data []models.LSNClientBinding `json:"lsnclient_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.LSNClientBinding{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.LSNClientBinding{}, fmt.Errorf("lsnclient_binding %s not found", name)
	}

	return result.Data[0], nil
}

// lsnclient_network6_binding
func (s *LSNService) AddLSNClientNetwork6Binding(resource models.LSNClientNetwork6Binding) error {
	payload := map[string]any{"lsnclient_network6_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lsnClientNetwork6BindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) DeleteLSNClientNetwork6Binding(name string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lsnClientNetwork6BindingURL, name, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) GetAllLSNClientNetwork6Binding() ([]models.LSNClientNetwork6Binding, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnClientNetwork6BindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNClientNetwork6Binding `json:"lsnclient_network6_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNClientNetwork6Binding(name string) ([]models.LSNClientNetwork6Binding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnClientNetwork6BindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNClientNetwork6Binding `json:"lsnclient_network6_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) CountLSNClientNetwork6Binding(clientname string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lsnClientNetwork6BindingURL, clientname), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LSNClientNetwork6Binding `json:"lsnclient_network6_binding"`
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

// lsnclient_network_binding
func (s *LSNService) AddLSNClientNetworkBinding(resource models.LSNClientNetworkBinding) error {
	payload := map[string]any{"lsnclient_network_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lsnClientNetworkBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) DeleteLSNClientNetworkBinding(name string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lsnClientNetworkBindingURL, name, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) GetAllLSNClientNetworkBinding() ([]models.LSNClientNetworkBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnClientNetworkBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNClientNetworkBinding `json:"lsnclient_network_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNClientNetworkBinding(name string) ([]models.LSNClientNetworkBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnClientNetworkBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNClientNetworkBinding `json:"lsnclient_network_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) CountLSNClientNetworkBinding(clientname string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lsnClientNetworkBindingURL, clientname), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LSNClientNetworkBinding `json:"lsnclient_network_binding"`
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

// lsnclient_nsacl6_binding
func (s *LSNService) AddLSNClientNSACL6Binding(resource models.LSNClientNSACL6Binding) error {
	payload := map[string]any{"lsnclient_nsacl6_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lsnClientNSACL6BindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) DeleteLSNClientNSACL6Binding(name string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lsnClientNSACL6BindingURL, name, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) GetAllLSNClientNSACL6Binding() ([]models.LSNClientNSACL6Binding, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnClientNSACL6BindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNClientNSACL6Binding `json:"lsnclient_nsacl6_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNClientNSACL6Binding(name string) ([]models.LSNClientNSACL6Binding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnClientNSACL6BindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNClientNSACL6Binding `json:"lsnclient_nsacl6_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) CountLSNClientNSACL6Binding(clientname string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lsnClientNSACL6BindingURL, clientname), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LSNClientNSACL6Binding `json:"lsnclient_nsacl6_binding"`
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

// lsnclient_nsacl_binding
func (s *LSNService) AddLSNClientNSACLBinding(resource models.LSNClientNSACLBinding) error {
	payload := map[string]any{"lsnclient_nsacl_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lsnClientNSACLBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) DeleteLSNClientNSACLBinding(name string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lsnClientNSACLBindingURL, name, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) GetAllLSNClientNSACLBinding() ([]models.LSNClientNSACLBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnClientNSACLBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNClientNSACLBinding `json:"lsnclient_nsacl_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNClientNSACLBinding(name string) ([]models.LSNClientNSACLBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnClientNSACLBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNClientNSACLBinding `json:"lsnclient_nsacl_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) CountLSNClientNSACLBinding(clientname string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lsnClientNSACLBindingURL, clientname), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LSNClientNSACLBinding `json:"lsnclient_nsacl_binding"`
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

// lsndeterministicnat
func (s *LSNService) GetAllLSNDeterministicNAT() ([]models.LSNDeterministicNAT, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnDeterministicNATURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNDeterministicNAT `json:"lsndeterministicnat"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) CountLSNDeterministicNAT() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", lsnDeterministicNATURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LSNDeterministicNAT `json:"lsndeterministicnat"`
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

// lsngroup
func (s *LSNService) AddLSNGroup(resource models.LSNGroup) error {
	payload := map[string]any{"lsngroup": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lsnGroupURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) DeleteLSNGroup(groupname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", lsnGroupURL, groupname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) UpdateLSNGroup(resource models.LSNGroup) error {
	payload := map[string]any{"lsngroup": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, lsnGroupURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) UnsetLSNGroup(resource models.LSNGroup) error {
	payload := map[string]any{"lsngroup": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", lsnGroupURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) GetAllLSNGroup() ([]models.LSNGroup, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnGroupURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNGroup `json:"lsngroup"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNGroup(groupname string) (models.LSNGroup, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnGroupURL, groupname), nil)
	if err != nil {
		return models.LSNGroup{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.LSNGroup{}, err
	}

	var result struct {
		Data []models.LSNGroup `json:"lsngroup"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.LSNGroup{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.LSNGroup{}, fmt.Errorf("lsngroup %s not found", groupname)
	}

	return result.Data[0], nil
}

func (s *LSNService) CountLSNGroup() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", lsnGroupURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LSNGroup `json:"lsngroup"`
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

// lsngroup_binding
func (s *LSNService) GetAllLSNGroupBinding() ([]models.LSNGroupBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnGroupBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNGroupBinding `json:"lsngroup_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNGroupBinding(groupname string) (models.LSNGroupBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnGroupBindingURL, groupname), nil)
	if err != nil {
		return models.LSNGroupBinding{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.LSNGroupBinding{}, err
	}

	var result struct {
		Data []models.LSNGroupBinding `json:"lsngroup_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.LSNGroupBinding{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.LSNGroupBinding{}, fmt.Errorf("lsngroup_binding %s not found", groupname)
	}

	return result.Data[0], nil
}

// lsngroup_ipsecalgprofile_binding
func (s *LSNService) AddLSNGroupIPSecALGProfileBinding(resource models.LSNGroupIPSECALGProfileBinding) error {
	payload := map[string]any{"lsngroup_ipsecalgprofile_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lsnGroupIPSecALGProfileBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) DeleteLSNGroupIPSecALGProfileBinding(groupname string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lsnGroupIPSecALGProfileBindingURL, groupname, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) GetAllLSNGroupIPSecALGProfileBinding() ([]models.LSNGroupIPSECALGProfileBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnGroupIPSecALGProfileBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNGroupIPSECALGProfileBinding `json:"lsngroup_ipsecalgprofile_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNGroupIPSecALGProfileBinding(groupname string) ([]models.LSNGroupIPSECALGProfileBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnGroupIPSecALGProfileBindingURL, groupname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNGroupIPSECALGProfileBinding `json:"lsngroup_ipsecalgprofile_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) CountLSNGroupIPSecALGProfileBinding(groupname string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lsnGroupIPSecALGProfileBindingURL, groupname), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LSNGroupIPSECALGProfileBinding `json:"lsngroup_ipsecalgprofile_binding"`
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

// lsngroup_lsnappsprofile_binding
func (s *LSNService) AddLSNGroupLSNAppsProfileBinding(resource models.LSNGroupLSNAppsProfileBinding) error {
	payload := map[string]any{"lsngroup_lsnappsprofile_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lsnGroupLSNAppsProfileBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) DeleteLSNGroupLSNAppsProfileBinding(groupname string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lsnGroupLSNAppsProfileBindingURL, groupname, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) GetAllLSNGroupLSNAppsProfileBinding() ([]models.LSNGroupLSNAppsProfileBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnGroupLSNAppsProfileBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNGroupLSNAppsProfileBinding `json:"lsngroup_lsnappsprofile_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNGroupLSNAppsProfileBinding(groupname string) ([]models.LSNGroupLSNAppsProfileBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnGroupLSNAppsProfileBindingURL, groupname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNGroupLSNAppsProfileBinding `json:"lsngroup_lsnappsprofile_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) CountLSNGroupLSNAppsProfileBinding(groupname string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lsnGroupLSNAppsProfileBindingURL, groupname), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LSNGroupLSNAppsProfileBinding `json:"lsngroup_lsnappsprofile_binding"`
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

// lsngroup_lsnhttphdrlogprofile_binding
func (s *LSNService) AddLSNGroupLSNHTTPHDRLogProfileBinding(resource models.LSNGroupLSNHTTPHdrLogProfileBinding) error {
	payload := map[string]any{"lsngroup_lsnhttphdrlogprofile_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lsnGroupLSNHTTPHDRLogProfileBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) DeleteLSNGroupLSNHTTPHDRLogProfileBinding(groupname string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lsnGroupLSNHTTPHDRLogProfileBindingURL, groupname, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) GetAllLSNGroupLSNHTTPHDRLogProfileBinding() ([]models.LSNGroupLSNHTTPHdrLogProfileBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnGroupLSNHTTPHDRLogProfileBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNGroupLSNHTTPHdrLogProfileBinding `json:"lsngroup_lsnhttphdrlogprofile_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNGroupLSNHTTPHDRLogProfileBinding(groupname string) ([]models.LSNGroupLSNHTTPHdrLogProfileBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnGroupLSNHTTPHDRLogProfileBindingURL, groupname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNGroupLSNHTTPHdrLogProfileBinding `json:"lsngroup_lsnhttphdrlogprofile_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) CountLSNGroupLSNHTTPHDRLogProfileBinding(groupname string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lsnGroupLSNHTTPHDRLogProfileBindingURL, groupname), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LSNGroupLSNHTTPHdrLogProfileBinding `json:"lsngroup_lsnhttphdrlogprofile_binding"`
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

// lsngroup_lsnlogprofile_binding
func (s *LSNService) AddLSNGroupLSNLogProfileBinding(resource models.LSNGroupLSNLogProfileBinding) error {
	payload := map[string]any{"lsngroup_lsnlogprofile_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lsnGroupLSNLogProfileBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) DeleteLSNGroupLSNLogProfileBinding(groupname string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lsnGroupLSNLogProfileBindingURL, groupname, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) GetAllLSNGroupLSNLogProfileBinding() ([]models.LSNGroupLSNLogProfileBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnGroupLSNLogProfileBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNGroupLSNLogProfileBinding `json:"lsngroup_lsnlogprofile_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNGroupLSNLogProfileBinding(groupname string) ([]models.LSNGroupLSNLogProfileBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnGroupLSNLogProfileBindingURL, groupname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNGroupLSNLogProfileBinding `json:"lsngroup_lsnlogprofile_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) CountLSNGroupLSNLogProfileBinding(groupname string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lsnGroupLSNLogProfileBindingURL, groupname), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LSNGroupLSNLogProfileBinding `json:"lsngroup_lsnlogprofile_binding"`
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

// lsngroup_lsnpool_binding
func (s *LSNService) AddLSNGroupLSNPoolBinding(resource models.LSNGroupLSNPoolBinding) error {
	payload := map[string]any{"lsngroup_lsnpool_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lsnGroupLSNPoolBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) DeleteLSNGroupLSNPoolBinding(groupname string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lsnGroupLSNPoolBindingURL, groupname, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) GetAllLSNGroupLSNPoolBinding() ([]models.LSNGroupLSNPoolBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnGroupLSNPoolBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNGroupLSNPoolBinding `json:"lsngroup_lsnpool_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNGroupLSNPoolBinding(groupname string) ([]models.LSNGroupLSNPoolBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnGroupLSNPoolBindingURL, groupname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNGroupLSNPoolBinding `json:"lsngroup_lsnpool_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) CountLSNGroupLSNPoolBinding(groupname string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lsnGroupLSNPoolBindingURL, groupname), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LSNGroupLSNPoolBinding `json:"lsngroup_lsnpool_binding"`
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

// lsngroup_lsnrtspalgprofile_binding
func (s *LSNService) AddLSNGroupLSNRTSPALGProfileBinding(resource models.LSNGroupLSNRTSPALGProfileBinding) error {
	payload := map[string]any{"lsngroup_lsnrtspalgprofile_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lsnGroupLSNRTSPALGProfileBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) DeleteLSNGroupLSNRTSPALGProfileBinding(groupname string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lsnGroupLSNRTSPALGProfileBindingURL, groupname, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) GetAllLSNGroupLSNRTSPALGProfileBinding() ([]models.LSNGroupLSNRTSPALGProfileBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnGroupLSNRTSPALGProfileBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNGroupLSNRTSPALGProfileBinding `json:"lsngroup_lsnrtspalgprofile_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNGroupLSNRTSPALGProfileBinding(groupname string) ([]models.LSNGroupLSNRTSPALGProfileBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnGroupLSNRTSPALGProfileBindingURL, groupname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNGroupLSNRTSPALGProfileBinding `json:"lsngroup_lsnrtspalgprofile_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) CountLSNGroupLSNRTSPALGProfileBinding(groupname string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lsnGroupLSNRTSPALGProfileBindingURL, groupname), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LSNGroupLSNRTSPALGProfileBinding `json:"lsngroup_lsnrtspalgprofile_binding"`
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

// lsngroup_lsnsipalgprofile_binding
func (s *LSNService) AddLSNGroupLSNSIPALGProfileBinding(resource models.LSNGroupLSNSIPALGProfileBinding) error {
	payload := map[string]any{"lsngroup_lsnsipalgprofile_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lsnGroupLSNSIPALGProfileBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) DeleteLSNGroupLSNSIPALGProfileBinding(groupname string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lsnGroupLSNSIPALGProfileBindingURL, groupname, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) GetAllLSNGroupLSNSIPALGProfileBinding() ([]models.LSNGroupLSNSIPALGProfileBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnGroupLSNSIPALGProfileBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNGroupLSNSIPALGProfileBinding `json:"lsngroup_lsnsipalgprofile_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNGroupLSNSIPALGProfileBinding(groupname string) ([]models.LSNGroupLSNSIPALGProfileBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnGroupLSNSIPALGProfileBindingURL, groupname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNGroupLSNSIPALGProfileBinding `json:"lsngroup_lsnsipalgprofile_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) CountLSNGroupLSNSIPALGProfileBinding(groupname string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lsnGroupLSNSIPALGProfileBindingURL, groupname), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LSNGroupLSNSIPALGProfileBinding `json:"lsngroup_lsnsipalgprofile_binding"`
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

// lsngroup_lsntransportprofile_binding
func (s *LSNService) AddLSNGroupLSNTransportProfileBinding(resource models.LSNGroupLSNTransportProfileBinding) error {
	payload := map[string]any{"lsngroup_lsntransportprofile_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lsnGroupLSNTransportProfileBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) DeleteLSNGroupLSNTransportProfileBinding(groupname string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lsnGroupLSNTransportProfileBindingURL, groupname, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) GetAllLSNGroupLSNTransportProfileBinding() ([]models.LSNGroupLSNTransportProfileBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnGroupLSNTransportProfileBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNGroupLSNTransportProfileBinding `json:"lsngroup_lsntransportprofile_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNGroupLSNTransportProfileBinding(groupname string) ([]models.LSNGroupLSNTransportProfileBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnGroupLSNTransportProfileBindingURL, groupname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNGroupLSNTransportProfileBinding `json:"lsngroup_lsntransportprofile_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) CountLSNGroupLSNTransportProfileBinding(groupname string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lsnGroupLSNTransportProfileBindingURL, groupname), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LSNGroupLSNTransportProfileBinding `json:"lsngroup_lsntransportprofile_binding"`
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

// lsngroup_pcpserver_binding
func (s *LSNService) AddLSNGroupPCPServerBinding(resource models.LSNGroupPCPServerBinding) error {
	payload := map[string]any{"lsngroup_pcpserver_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lsnGroupPCPServerBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) DeleteLSNGroupPCPServerBinding(groupname string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lsnGroupPCPServerBindingURL, groupname, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) GetAllLSNGroupPCPServerBinding() ([]models.LSNGroupPCPServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnGroupPCPServerBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNGroupPCPServerBinding `json:"lsngroup_pcpserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNGroupPCPServerBinding(groupname string) ([]models.LSNGroupPCPServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnGroupPCPServerBindingURL, groupname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNGroupPCPServerBinding `json:"lsngroup_pcpserver_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) CountLSNGroupPCPServerBinding(groupname string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lsnGroupPCPServerBindingURL, groupname), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LSNGroupPCPServerBinding `json:"lsngroup_pcpserver_binding"`
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

// lsnhttphdrlogprofile
func (s *LSNService) AddLSNHTTPHDRLogProfile(resource models.LSNHTTPHdrLogProfile) error {
	payload := map[string]any{"lsnhttphdrlogprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lsnHTTPHDRLogProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) DeleteLSNHTTPHDRLogProfile(httphdrlogprofilename string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", lsnHTTPHDRLogProfileURL, httphdrlogprofilename), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) UpdateLSNHTTPHDRLogProfile(resource models.LSNHTTPHdrLogProfile) error {
	payload := map[string]any{"lsnhttphdrlogprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, lsnHTTPHDRLogProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) UnsetLSNHTTPHDRLogProfile(resource models.LSNHTTPHdrLogProfile) error {
	payload := map[string]any{"lsnhttphdrlogprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", lsnHTTPHDRLogProfileURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) GetAllLSNHTTPHDRLogProfile() ([]models.LSNHTTPHdrLogProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnHTTPHDRLogProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNHTTPHdrLogProfile `json:"lsnhttphdrlogprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNHTTPHDRLogProfile(httphdrlogprofilename string) (models.LSNHTTPHdrLogProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnHTTPHDRLogProfileURL, httphdrlogprofilename), nil)
	if err != nil {
		return models.LSNHTTPHdrLogProfile{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.LSNHTTPHdrLogProfile{}, err
	}

	var result struct {
		Data []models.LSNHTTPHdrLogProfile `json:"lsnhttphdrlogprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.LSNHTTPHdrLogProfile{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.LSNHTTPHdrLogProfile{}, fmt.Errorf("lsnhttphdrlogprofile %s not found", httphdrlogprofilename)
	}

	return result.Data[0], nil
}

func (s *LSNService) CountLSNHTTPHDRLogProfile() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", lsnHTTPHDRLogProfileURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LSNHTTPHdrLogProfile `json:"lsnhttphdrlogprofile"`
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

// lsnip6profile
func (s *LSNService) AddLSNIP6Profile(resource models.LSNIP6Profile) error {
	payload := map[string]any{"lsnip6profile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lsnIP6ProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) DeleteLSNIP6Profile(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", lsnIP6ProfileURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) GetAllLSNIP6Profile() ([]models.LSNIP6Profile, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnIP6ProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNIP6Profile `json:"lsnip6profile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNIP6Profile(name string) (models.LSNIP6Profile, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnIP6ProfileURL, name), nil)
	if err != nil {
		return models.LSNIP6Profile{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.LSNIP6Profile{}, err
	}

	var result struct {
		Data []models.LSNIP6Profile `json:"lsnip6profile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.LSNIP6Profile{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.LSNIP6Profile{}, fmt.Errorf("lsnip6profile %s not found", name)
	}

	return result.Data[0], nil
}

func (s *LSNService) CountLSNIP6Profile() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", lsnIP6ProfileURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LSNIP6Profile `json:"lsnip6profile"`
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

// lsnlogprofile
func (s *LSNService) AddLSNLogProfile(resource models.LSNLogProfile) error {
	payload := map[string]any{"lsnlogprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lsnLogProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) DeleteLSNLogProfile(logprofilename string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", lsnLogProfileURL, logprofilename), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) UpdateLSNLogProfile(resource models.LSNLogProfile) error {
	payload := map[string]any{"lsnlogprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, lsnLogProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) UnsetLSNLogProfile(resource models.LSNLogProfile) error {
	payload := map[string]any{"lsnlogprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", lsnLogProfileURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) GetAllLSNLogProfile() ([]models.LSNLogProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnLogProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNLogProfile `json:"lsnlogprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNLogProfile(logprofilename string) (models.LSNLogProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnLogProfileURL, logprofilename), nil)
	if err != nil {
		return models.LSNLogProfile{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.LSNLogProfile{}, err
	}

	var result struct {
		Data []models.LSNLogProfile `json:"lsnlogprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.LSNLogProfile{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.LSNLogProfile{}, fmt.Errorf("lsnlogprofile %s not found", logprofilename)
	}

	return result.Data[0], nil
}

func (s *LSNService) CountLSNLogProfile() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", lsnLogProfileURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LSNLogProfile `json:"lsnlogprofile"`
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

// lsnparameter
func (s *LSNService) UpdateLSNParameter(resource models.LSNParameter) error {
	payload := map[string]any{"lsnparameter": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, lsnParameterURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) UnsetLSNParameter(resource models.LSNParameter) error {
	payload := map[string]any{"lsnparameter": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", lsnParameterURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) GetAllLSNParameter() (models.LSNParameter, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnParameterURL, nil)
	if err != nil {
		return models.LSNParameter{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.LSNParameter{}, err
	}

	var result struct {
		Data models.LSNParameter `json:"lsnparameter"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.LSNParameter{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// lsnpool
func (s *LSNService) AddLSNPool(resource models.LSNPool) error {
	payload := map[string]any{"lsnpool": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lsnPoolURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) DeleteLSNPool(poolname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", lsnPoolURL, poolname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) UpdateLSNPool(resource models.LSNPool) error {
	payload := map[string]any{"lsnpool": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, lsnPoolURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) UnsetLSNPool(resource models.LSNPool) error {
	payload := map[string]any{"lsnpool": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", lsnPoolURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) GetAllLSNPool() ([]models.LSNPool, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnPoolURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNPool `json:"lsnpool"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNPool(poolname string) (models.LSNPool, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnPoolURL, poolname), nil)
	if err != nil {
		return models.LSNPool{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.LSNPool{}, err
	}

	var result struct {
		Data []models.LSNPool `json:"lsnpool"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.LSNPool{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.LSNPool{}, fmt.Errorf("lsnpool %s not found", poolname)
	}

	return result.Data[0], nil
}

func (s *LSNService) CountLSNPool() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", lsnPoolURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LSNPool `json:"lsnpool"`
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

// lsnpool_binding
func (s *LSNService) GetAllLSNPoolBinding() ([]models.LSNPoolBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnPoolBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNPoolBinding `json:"lsnpool_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNPoolBinding(poolname string) (models.LSNPoolBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnPoolBindingURL, poolname), nil)
	if err != nil {
		return models.LSNPoolBinding{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.LSNPoolBinding{}, err
	}

	var result struct {
		Data []models.LSNPoolBinding `json:"lsnpool_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.LSNPoolBinding{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.LSNPoolBinding{}, fmt.Errorf("lsnpool_binding %s not found", poolname)
	}

	return result.Data[0], nil
}

// lsnpool_lsnip_binding
func (s *LSNService) AddLSNPoolLSNIPBinding(resource models.LSNPoolLSNIPBinding) error {
	payload := map[string]any{"lsnpool_lsnip_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lsnPoolLSNIPBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) DeleteLSNPoolLSNIPBinding(poolname string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", lsnPoolLSNIPBindingURL, poolname, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) GetAllLSNPoolLSNIPBinding() ([]models.LSNPoolLSNIPBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnPoolLSNIPBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNPoolLSNIPBinding `json:"lsnpool_lsnip_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNPoolLSNIPBinding(poolname string) ([]models.LSNPoolLSNIPBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnPoolLSNIPBindingURL, poolname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNPoolLSNIPBinding `json:"lsnpool_lsnip_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) CountLSNPoolLSNIPBinding(poolname string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lsnPoolLSNIPBindingURL, poolname), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LSNPoolLSNIPBinding `json:"lsnpool_lsnip_binding"`
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

// lsnrtspalgprofile
func (s *LSNService) AddLSNRTSPALGProfile(resource models.LSNRTSPALGProfile) error {
	payload := map[string]any{"lsnrtspalgprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lsnRTSPALGProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) DeleteLSNRTSPALGProfile(profilename string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", lsnRTSPALGProfileURL, profilename), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) UpdateLSNRTSPALGProfile(resource models.LSNRTSPALGProfile) error {
	payload := map[string]any{"lsnrtspalgprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, lsnRTSPALGProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) UnsetLSNRTSPALGProfile(resource models.LSNRTSPALGProfile) error {
	payload := map[string]any{"lsnrtspalgprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", lsnRTSPALGProfileURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) GetAllLSNRTSPALGProfile() ([]models.LSNRTSPALGProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnRTSPALGProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNRTSPALGProfile `json:"lsnrtspalgprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNRTSPALGProfile(profilename string) (models.LSNRTSPALGProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnRTSPALGProfileURL, profilename), nil)
	if err != nil {
		return models.LSNRTSPALGProfile{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.LSNRTSPALGProfile{}, err
	}

	var result struct {
		Data []models.LSNRTSPALGProfile `json:"lsnrtspalgprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.LSNRTSPALGProfile{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.LSNRTSPALGProfile{}, fmt.Errorf("lsnrtspalgprofile %s not found", profilename)
	}

	return result.Data[0], nil
}

func (s *LSNService) CountLSNRTSPALGProfile() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", lsnRTSPALGProfileURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LSNRTSPALGProfile `json:"lsnrtspalgprofile"`
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

// lsnrtspalgsession
func (s *LSNService) GetAllLSNRTSPALGSession() ([]models.LSNRTSPALGSession, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnRTSPALGSessionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNRTSPALGSession `json:"lsnrtspalgsession"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNRTSPALGSession(sessionid string) (models.LSNRTSPALGSession, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnRTSPALGSessionURL, sessionid), nil)
	if err != nil {
		return models.LSNRTSPALGSession{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.LSNRTSPALGSession{}, err
	}

	var result struct {
		Data []models.LSNRTSPALGSession `json:"lsnrtspalgsession"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.LSNRTSPALGSession{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.LSNRTSPALGSession{}, fmt.Errorf("lsnrtspalgsession %s not found", sessionid)
	}

	return result.Data[0], nil
}

func (s *LSNService) CountLSNRTSPALGSession() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", lsnRTSPALGSessionURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LSNRTSPALGSession `json:"lsnrtspalgsession"`
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

func (s *LSNService) FlushLSNRTSPALGSession(resource models.LSNRTSPALGSession) error {
	payload := map[string]any{"lsnrtspalgsession": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=flush", lsnRTSPALGSessionURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// lsnrtspalgsession_binding
func (s *LSNService) GetAllLSNRTSPALGSessionBinding() ([]models.LSNRTSPALGSessionBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnRTSPALGSessionBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNRTSPALGSessionBinding `json:"lsnrtspalgsession_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNRTSPALGSessionBinding(sessionid string) (models.LSNRTSPALGSessionBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnRTSPALGSessionBindingURL, sessionid), nil)
	if err != nil {
		return models.LSNRTSPALGSessionBinding{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.LSNRTSPALGSessionBinding{}, err
	}

	var result struct {
		Data []models.LSNRTSPALGSessionBinding `json:"lsnrtspalgsession_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.LSNRTSPALGSessionBinding{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.LSNRTSPALGSessionBinding{}, fmt.Errorf("lsnrtspalgsession_binding %s not found", sessionid)
	}

	return result.Data[0], nil
}

// lsnrtspalgsession_datachannel_binding
func (s *LSNService) GetAllLSNRTSPALGSessionDataChannelBinding() ([]models.LSNRTSPALGSessionDataChannelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnRTSPALGSessionDataChannelBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNRTSPALGSessionDataChannelBinding `json:"lsnrtspalgsession_datachannel_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNRTSPALGSessionDataChannelBinding(sessionid string) ([]models.LSNRTSPALGSessionDataChannelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnRTSPALGSessionDataChannelBindingURL, sessionid), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNRTSPALGSessionDataChannelBinding `json:"lsnrtspalgsession_datachannel_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) CountLSNRTSPALGSessionDataChannelBinding(sessionid string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lsnRTSPALGSessionDataChannelBindingURL, sessionid), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LSNRTSPALGSessionDataChannelBinding `json:"lsnrtspalgsession_datachannel_binding"`
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

// lsnsession
func (s *LSNService) GetAllLSNSession() ([]models.LSNSession, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnSessionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNSession `json:"lsnsession"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) CountLSNSession() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", lsnSessionURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LSNSession `json:"lsnsession"`
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

func (s *LSNService) FlushLSNSession(resource models.LSNSession) error {
	payload := map[string]any{"lsnsession": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=flush", lsnSessionURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// lsnsipalgcall
func (s *LSNService) GetAllLSNSIPALGCall() ([]models.LSNSIPALGCall, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnSIPALGCallURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNSIPALGCall `json:"lsnsipalgcall"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNSIPALGCall(callid string) (models.LSNSIPALGCall, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnSIPALGCallURL, callid), nil)
	if err != nil {
		return models.LSNSIPALGCall{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.LSNSIPALGCall{}, err
	}

	var result struct {
		Data []models.LSNSIPALGCall `json:"lsnsipalgcall"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.LSNSIPALGCall{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.LSNSIPALGCall{}, fmt.Errorf("lsnsipalgcall %s not found", callid)
	}

	return result.Data[0], nil
}

func (s *LSNService) CountLSNSIPALGCall() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", lsnSIPALGCallURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LSNSIPALGCall `json:"lsnsipalgcall"`
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

func (s *LSNService) FlushLSNSIPALGCall(resource models.LSNSIPALGCall) error {
	payload := map[string]any{"lsnsipalgcall": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=flush", lsnSIPALGCallURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// lsnsipalgcall_binding
func (s *LSNService) GetAllLSNSIPALGCallBinding() ([]models.LSNSIPALGCallBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnSIPALGCallBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNSIPALGCallBinding `json:"lsnsipalgcall_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNSIPALGCallBinding(callid string) (models.LSNSIPALGCallBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnSIPALGCallBindingURL, callid), nil)
	if err != nil {
		return models.LSNSIPALGCallBinding{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.LSNSIPALGCallBinding{}, err
	}

	var result struct {
		Data []models.LSNSIPALGCallBinding `json:"lsnsipalgcall_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.LSNSIPALGCallBinding{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.LSNSIPALGCallBinding{}, fmt.Errorf("lsnsipalgcall_binding %s not found", callid)
	}

	return result.Data[0], nil
}

// lsnsipalgcall_controlchannel_binding
func (s *LSNService) GetAllLSNSIPALGCallControlChannelBinding() ([]models.LSNSIPALGCallControlChannelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnSIPALGCallControlChannelBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNSIPALGCallControlChannelBinding `json:"lsnsipalgcall_controlchannel_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNSIPALGCallControlChannelBinding(callid string) ([]models.LSNSIPALGCallControlChannelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnSIPALGCallControlChannelBindingURL, callid), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNSIPALGCallControlChannelBinding `json:"lsnsipalgcall_controlchannel_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) CountLSNSIPALGCallControlChannelBinding(callid string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lsnSIPALGCallControlChannelBindingURL, callid), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LSNSIPALGCallControlChannelBinding `json:"lsnsipalgcall_controlchannel_binding"`
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

// lsnsipalgcall_datachannel_binding
func (s *LSNService) GetAllLSNSIPALGCallDataChannelBinding() ([]models.LSNSIPALGCallDataChannelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnSIPALGCallDataChannelBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNSIPALGCallDataChannelBinding `json:"lsnsipalgcall_datachannel_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNSIPALGCallDataChannelBinding(callid string) ([]models.LSNSIPALGCallDataChannelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnSIPALGCallDataChannelBindingURL, callid), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNSIPALGCallDataChannelBinding `json:"lsnsipalgcall_datachannel_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) CountLSNSIPALGCallDataChannelBinding(callid string) (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?count=yes", lsnSIPALGCallDataChannelBindingURL, callid), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LSNSIPALGCallDataChannelBinding `json:"lsnsipalgcall_datachannel_binding"`
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

// lsnsipalgprofile
func (s *LSNService) AddLSNSIPALGProfile(resource models.LSNSIPALGProfile) error {
	payload := map[string]any{"lsnsipalgprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lsnSIPALGProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) DeleteLSNSIPALGProfile(sipalgprofilename string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", lsnSIPALGProfileURL, sipalgprofilename), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) UpdateLSNSIPALGProfile(resource models.LSNSIPALGProfile) error {
	payload := map[string]any{"lsnsipalgprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, lsnSIPALGProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) UnsetLSNSIPALGProfile(resource models.LSNSIPALGProfile) error {
	payload := map[string]any{"lsnsipalgprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", lsnSIPALGProfileURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) GetAllLSNSIPALGProfile() ([]models.LSNSIPALGProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnSIPALGProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNSIPALGProfile `json:"lsnsipalgprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNSIPALGProfile(sipalgprofilename string) (models.LSNSIPALGProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnSIPALGProfileURL, sipalgprofilename), nil)
	if err != nil {
		return models.LSNSIPALGProfile{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.LSNSIPALGProfile{}, err
	}

	var result struct {
		Data []models.LSNSIPALGProfile `json:"lsnsipalgprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.LSNSIPALGProfile{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.LSNSIPALGProfile{}, fmt.Errorf("lsnsipalgprofile %s not found", sipalgprofilename)
	}

	return result.Data[0], nil
}

func (s *LSNService) CountLSNSIPALGProfile() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", lsnSIPALGProfileURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LSNSIPALGProfile `json:"lsnsipalgprofile"`
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

// lsnstatic
func (s *LSNService) AddLSNStatic(resource models.LSNStatic) error {
	payload := map[string]any{"lsnstatic": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lsnStaticURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) DeleteLSNStatic(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", lsnStaticURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) GetAllLSNStatic() ([]models.LSNStatic, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnStaticURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNStatic `json:"lsnstatic"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNStatic(name string) (models.LSNStatic, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnStaticURL, name), nil)
	if err != nil {
		return models.LSNStatic{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.LSNStatic{}, err
	}

	var result struct {
		Data []models.LSNStatic `json:"lsnstatic"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.LSNStatic{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.LSNStatic{}, fmt.Errorf("lsnstatic %s not found", name)
	}

	return result.Data[0], nil
}

func (s *LSNService) CountLSNStatic() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", lsnStaticURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LSNStatic `json:"lsnstatic"`
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

// lsntransportprofile
func (s *LSNService) AddLSNTransportProfile(resource models.LSNTransportProfile) error {
	payload := map[string]any{"lsntransportprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, lsnTransportProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) DeleteLSNTransportProfile(transportprofilename string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", lsnTransportProfileURL, transportprofilename), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) UpdateLSNTransportProfile(resource models.LSNTransportProfile) error {
	payload := map[string]any{"lsntransportprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, lsnTransportProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) UnsetLSNTransportProfile(resource models.LSNTransportProfile) error {
	payload := map[string]any{"lsntransportprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", lsnTransportProfileURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *LSNService) GetAllLSNTransportProfile() ([]models.LSNTransportProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, lsnTransportProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LSNTransportProfile `json:"lsntransportprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *LSNService) GetLSNTransportProfile(transportprofilename string) (models.LSNTransportProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lsnTransportProfileURL, transportprofilename), nil)
	if err != nil {
		return models.LSNTransportProfile{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.LSNTransportProfile{}, err
	}

	var result struct {
		Data []models.LSNTransportProfile `json:"lsntransportprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.LSNTransportProfile{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Data) == 0 {
		return models.LSNTransportProfile{}, fmt.Errorf("lsntransportprofile %s not found", transportprofilename)
	}

	return result.Data[0], nil
}

func (s *LSNService) CountLSNTransportProfile() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", lsnTransportProfileURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LSNTransportProfile `json:"lsntransportprofile"`
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
