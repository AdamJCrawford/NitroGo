package nitrogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AdamJCrawford/NitroGo/nitrogo/models"
)

const (
	appAlgParamURL                           = "/nitro/v1/config/appalgparam"
	arpURL                                   = "/nitro/v1/config/arp"
	arpParamURL                              = "/nitro/v1/config/arpparam"
	bridgeGroupURL                           = "/nitro/v1/config/bridgegroup"
	bridgeGroupBindingURL                    = "/nitro/v1/config/bridgegroup_binding"
	bridgeGroupNSIP6BindingURL               = "/nitro/v1/config/bridgegroup_nsip6_binding"
	bridgeGroupNSIPBindingURL                = "/nitro/v1/config/bridgegroup_nsip_binding"
	bridgeGroupVLANBindingURL                = "/nitro/v1/config/bridgegroup_vlan_binding"
	bridgeGroupVXLANBindingURL               = "/nitro/v1/config/bridgegroup_vxlan_binding"
	bridgeTableURL                           = "/nitro/v1/config/bridgetable"
	channelURL                               = "/nitro/v1/config/channel"
	channelBindingURL                        = "/nitro/v1/config/channel_binding"
	channelInterfaceBindingURL               = "/nitro/v1/config/channel_interface_binding"
	ciURL                                    = "/nitro/v1/config/ci"
	fisURL                                   = "/nitro/v1/config/fis"
	fisBindingURL                            = "/nitro/v1/config/fis_binding"
	fisChannelBindingURL                     = "/nitro/v1/config/fis_channel_binding"
	fisInterfaceBindingURL                   = "/nitro/v1/config/fis_interface_binding"
	forwardingSessionURL                     = "/nitro/v1/config/forwardingsession"
	inatURL                                  = "/nitro/v1/config/inat"
	inatParamURL                             = "/nitro/v1/config/inatparam"
	interfaceURL                             = "/nitro/v1/config/interface"
	interfaceBindingURL                      = "/nitro/v1/config/interface_binding"
	interfacePairURL                         = "/nitro/v1/config/interfacepair"
	ip6TunnelURL                             = "/nitro/v1/config/ip6tunnel"
	ip6TunnelParamURL                        = "/nitro/v1/config/ip6tunnelparam"
	ipSetURL                                 = "/nitro/v1/config/ipset"
	ipSetBindingURL                          = "/nitro/v1/config/ipset_binding"
	ipSetNSIP6BindingURL                     = "/nitro/v1/config/ipset_nsip6_binding"
	ipSetNSIPBindingURL                      = "/nitro/v1/config/ipset_nsip_binding"
	ipTunnelURL                              = "/nitro/v1/config/iptunnel"
	ipTunnelParamURL                         = "/nitro/v1/config/iptunnelparam"
	ipv6URL                                  = "/nitro/v1/config/ipv6"
	l2ParamURL                               = "/nitro/v1/config/l2param"
	l3ParamURL                               = "/nitro/v1/config/l3param"
	l4ParamURL                               = "/nitro/v1/config/l4param"
	lacpURL                                  = "/nitro/v1/config/lacp"
	linkSetURL                               = "/nitro/v1/config/linkset"
	linkSetBindingURL                        = "/nitro/v1/config/linkset_binding"
	linkSetChannelBindingURL                 = "/nitro/v1/config/linkset_channel_binding"
	linkSetInterfaceBindingURL               = "/nitro/v1/config/linkset_interface_binding"
	mapBMRURL                                = "/nitro/v1/config/mapbmr"
	mapBMRBindingURL                         = "/nitro/v1/config/mapbmr_binding"
	mapBMRBMRV4NetworkBindingURL             = "/nitro/v1/config/mapbmr_bmrv4network_binding"
	mapDMRURL                                = "/nitro/v1/config/mapdmr"
	mapDomainURL                             = "/nitro/v1/config/mapdomain"
	mapDomainBindingURL                      = "/nitro/v1/config/mapdomain_binding"
	mapDomainMapBMRBindingURL                = "/nitro/v1/config/mapdomain_mapbmr_binding"
	nat64URL                                 = "/nitro/v1/config/nat64"
	nat64ParamURL                            = "/nitro/v1/config/nat64param"
	nd6URL                                   = "/nitro/v1/config/nd6"
	nd6RAVariablesURL                        = "/nitro/v1/config/nd6ravariables"
	nd6RAVariablesBindingURL                 = "/nitro/v1/config/nd6ravariables_binding"
	nd6RAVariablesOnLinkIPV6PrefixBindingURL = "/nitro/v1/config/nd6ravariables_onlinkipv6prefix_binding"
	netBridgeURL                             = "/nitro/v1/config/netbridge"
	netBridgeBindingURL                      = "/nitro/v1/config/netbridge_binding"
	netBridgeIPTunnelBindingURL              = "/nitro/v1/config/netbridge_iptunnel_binding"
	netBridgeNSIP6BindingURL                 = "/nitro/v1/config/netbridge_nsip6_binding"
	netBridgeNSIPBindingURL                  = "/nitro/v1/config/netbridge_nsip_binding"
	netBridgeVLANBindingURL                  = "/nitro/v1/config/netbridge_vlan_binding"
	netProfileURL                            = "/nitro/v1/config/netprofile"
	netProfileBindingURL                     = "/nitro/v1/config/netprofile_binding"
	netProfileNATRuleBindingURL              = "/nitro/v1/config/netprofile_natrule_binding"
	netProfileSrcPortSetBindingURL           = "/nitro/v1/config/netprofile_srcportset_binding"
	onLinkIPV6PrefixURL                      = "/nitro/v1/config/onlinkipv6prefix"
	ptpURL                                   = "/nitro/v1/config/ptp"
	rnatURL                                  = "/nitro/v1/config/rnat"
	rnatBindingURL                           = "/nitro/v1/config/rnat_binding"
	rnatNSIPBindingURL                       = "/nitro/v1/config/rnat_nsip_binding"
	rnatRetainSourcePortSetBindingURL        = "/nitro/v1/config/rnat_retainsourceportset_binding"
	rnat6URL                                 = "/nitro/v1/config/rnat6"
	rnat6BindingURL                          = "/nitro/v1/config/rnat6_binding"
	rnat6NSIP6BindingURL                     = "/nitro/v1/config/rnat6_nsip6_binding"
	rnatGlobalBindingURL                     = "/nitro/v1/config/rnatglobal_binding"
	rnatGlobalAuditSyslogPolicyBindingURL    = "/nitro/v1/config/rnatglobal_auditsyslogpolicy_binding"
	rnatParamURL                             = "/nitro/v1/config/rnatparam"
	rnatSessionURL                           = "/nitro/v1/config/rnatsession"
	routeURL                                 = "/nitro/v1/config/route"
	route6URL                                = "/nitro/v1/config/route6"
	rssKeyTypeURL                            = "/nitro/v1/config/rsskeytype"
	vlanURL                                  = "/nitro/v1/config/vlan"
	vlanBindingURL                           = "/nitro/v1/config/vlan_binding"
	vlanChannelBindingURL                    = "/nitro/v1/config/vlan_channel_binding"
	vlanInterfaceBindingURL                  = "/nitro/v1/config/vlan_interface_binding"
	vlanLinkSetBindingURL                    = "/nitro/v1/config/vlan_linkset_binding"
	vlanNSIP6BindingURL                      = "/nitro/v1/config/vlan_nsip6_binding"
	vlanNSIPBindingURL                       = "/nitro/v1/config/vlan_nsip_binding"
	vridURL                                  = "/nitro/v1/config/vrid"
	vridBindingURL                           = "/nitro/v1/config/vrid_binding"
	vridChannelBindingURL                    = "/nitro/v1/config/vrid_channel_binding"
	vridInterfaceBindingURL                  = "/nitro/v1/config/vrid_interface_binding"
	vridNSIP6BindingURL                      = "/nitro/v1/config/vrid_nsip6_binding"
	vridNSIPBindingURL                       = "/nitro/v1/config/vrid_nsip_binding"
	vridTrackInterfaceBindingURL             = "/nitro/v1/config/vrid_trackinterface_binding"
	vrid6URL                                 = "/nitro/v1/config/vrid6"
	vrid6BindingURL                          = "/nitro/v1/config/vrid6_binding"
	vrid6ChannelBindingURL                   = "/nitro/v1/config/vrid6_channel_binding"
	vrid6InterfaceBindingURL                 = "/nitro/v1/config/vrid6_interface_binding"
	vrid6NSIP6BindingURL                     = "/nitro/v1/config/vrid6_nsip6_binding"
	vrid6NSIPBindingURL                      = "/nitro/v1/config/vrid6_nsip_binding"
	vrid6TrackInterfaceBindingURL            = "/nitro/v1/config/vrid6_trackinterface_binding"
	vridParamURL                             = "/nitro/v1/config/vridparam"
	vxlanURL                                 = "/nitro/v1/config/vxlan"
	vxlanBindingURL                          = "/nitro/v1/config/vxlan_binding"
	vxlanIPTunnelBindingURL                  = "/nitro/v1/config/vxlan_iptunnel_binding"
	vxlanNSIP6BindingURL                     = "/nitro/v1/config/vxlan_nsip6_binding"
	vxlanNSIPBindingURL                      = "/nitro/v1/config/vxlan_nsip_binding"
	vxlanSrcIPBindingURL                     = "/nitro/v1/config/vxlan_srcip_binding"
	vxlanVLANMapURL                          = "/nitro/v1/config/vxlanvlanmap"
	vxlanVLANMapBindingURL                   = "/nitro/v1/config/vxlanvlanmap_binding"
	vxlanVLANMapVXLANBindingURL              = "/nitro/v1/config/vxlanvlanmap_vxlan_binding"
)

type NetworkService struct {
	client *Client
}

// appalgparam
func (s *NetworkService) UpdateAppAlgParam(resource models.AppALGParam) error {
	payload := map[string]any{"appalgparam": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, appAlgParamURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UnsetAppAlgParam(resource models.AppALGParam) error {
	payload := map[string]any{"appalgparam": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", appAlgParamURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllAppAlgParam() ([]models.AppALGParam, error) {
	req, err := s.client.NewRequest(http.MethodGet, appAlgParamURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.AppALGParam `json:"appalgparam"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// arp
func (s *NetworkService) AddARP(resource models.ARP) error {
	payload := map[string]any{"arp": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, arpURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteARP(ipaddress string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", arpURL, ipaddress), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) SendARP(resource models.ARP) error {
	payload := map[string]any{"arp": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=send", arpURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllARP() ([]models.ARP, error) {
	req, err := s.client.NewRequest(http.MethodGet, arpURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.ARP `json:"arp"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) CountARP() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", arpURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.ARP `json:"arp"`
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

// arpparam
func (s *NetworkService) UpdateARPParam(resource models.ARPParam) error {
	payload := map[string]any{"arpparam": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, arpParamURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UnsetARPParam(resource models.ARPParam) error {
	payload := map[string]any{"arpparam": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", arpParamURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllARPParam() ([]models.ARPParam, error) {
	req, err := s.client.NewRequest(http.MethodGet, arpParamURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.ARPParam `json:"arpparam"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// bridgegroup
func (s *NetworkService) AddBridgeGroup(resource models.BridgeGroup) error {
	payload := map[string]any{"bridgegroup": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, bridgeGroupURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteBridgeGroup(id string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", bridgeGroupURL, id), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UpdateBridgeGroup(resource models.BridgeGroup) error {
	payload := map[string]any{"bridgegroup": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, bridgeGroupURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UnsetBridgeGroup(resource models.BridgeGroup) error {
	payload := map[string]any{"bridgegroup": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", bridgeGroupURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllBridgeGroup() ([]models.BridgeGroup, error) {
	req, err := s.client.NewRequest(http.MethodGet, bridgeGroupURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.BridgeGroup `json:"bridgegroup"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetBridgeGroup(id string) (*models.BridgeGroup, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", bridgeGroupURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.BridgeGroup `json:"bridgegroup"`
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

func (s *NetworkService) CountBridgeGroup() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", bridgeGroupURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.BridgeGroup `json:"bridgegroup"`
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

// bridgegroup_binding
func (s *NetworkService) GetAllBridgeGroupBinding() ([]models.BridgeGroupBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, bridgeGroupBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.BridgeGroupBinding `json:"bridgegroup_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetBridgeGroupBinding(id string) (*models.BridgeGroupBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", bridgeGroupBindingURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.BridgeGroupBinding `json:"bridgegroup_binding"`
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

// bridgegroup_nsip6_binding
func (s *NetworkService) AddBridgeGroupNSIP6Binding(resource models.BridgeGroupNSIP6Binding) error {
	payload := map[string]any{"bridgegroup_nsip6_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, bridgeGroupNSIP6BindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteBridgeGroupNSIP6Binding(id string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", bridgeGroupNSIP6BindingURL, id), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllBridgeGroupNSIP6Binding() ([]models.BridgeGroupNSIP6Binding, error) {
	req, err := s.client.NewRequest(http.MethodGet, bridgeGroupNSIP6BindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.BridgeGroupNSIP6Binding `json:"bridgegroup_nsip6_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetBridgeGroupNSIP6Binding(id string) (*models.BridgeGroupNSIP6Binding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", bridgeGroupNSIP6BindingURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.BridgeGroupNSIP6Binding `json:"bridgegroup_nsip6_binding"`
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

func (s *NetworkService) CountBridgeGroupNSIP6Binding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", bridgeGroupNSIP6BindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.BridgeGroupNSIP6Binding `json:"bridgegroup_nsip6_binding"`
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

// bridgegroup_nsip_binding
func (s *NetworkService) AddBrigeGroupNSIPBinding(resource models.BridgeGroupNSIPBinding) error {
	payload := map[string]any{"bridgegroup_nsip_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, bridgeGroupNSIPBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteBrigeGroupNSIPBinding(id string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", bridgeGroupNSIPBindingURL, id), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllBrigeGroupNSIPBinding() ([]models.BridgeGroupNSIPBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, bridgeGroupNSIPBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.BridgeGroupNSIPBinding `json:"bridgegroup_nsip_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetBrigeGroupNSIPBinding(id string) (*models.BridgeGroupNSIPBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", bridgeGroupNSIPBindingURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.BridgeGroupNSIPBinding `json:"bridgegroup_nsip_binding"`
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

func (s *NetworkService) CountBrigeGroupNSIPBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", bridgeGroupNSIPBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.BridgeGroupNSIPBinding `json:"bridgegroup_nsip_binding"`
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

// bridgegroup_vlan_binding
func (s *NetworkService) AddBridgeGroupVLANBinding(resource models.BridgeGroupVLANBinding) error {
	payload := map[string]any{"bridgegroup_vlan_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, bridgeGroupVLANBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteBridgeGroupVLANBinding(id string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", bridgeGroupVLANBindingURL, id), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllBridgeGroupVLANBinding() ([]models.BridgeGroupVLANBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, bridgeGroupVLANBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.BridgeGroupVLANBinding `json:"bridgegroup_vlan_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetBridgeGroupVLANBinding(id string) (*models.BridgeGroupVLANBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", bridgeGroupVLANBindingURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.BridgeGroupVLANBinding `json:"bridgegroup_vlan_binding"`
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

func (s *NetworkService) CountBridgeGroupVLANBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", bridgeGroupVLANBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.BridgeGroupVLANBinding `json:"bridgegroup_vlan_binding"`
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

// bridgetable
func (s *NetworkService) AddBridgeTable(resource models.BridgeTable) error {
	payload := map[string]any{"bridgetable": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, bridgeTableURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteBridgeTable(mac string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", bridgeTableURL, mac), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UpdateBridgeTable(resource models.BridgeTable) error {
	payload := map[string]any{"bridgetable": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, bridgeTableURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UnsetBridgeTable(resource models.BridgeTable) error {
	payload := map[string]any{"bridgetable": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", bridgeTableURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllBridgeTable() ([]models.BridgeTable, error) {
	req, err := s.client.NewRequest(http.MethodGet, bridgeTableURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.BridgeTable `json:"bridgetable"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) CountBridgeTable() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", bridgeTableURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.BridgeTable `json:"bridgetable"`
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

// channel
func (s *NetworkService) AddChannel(resource models.Channel) error {
	payload := map[string]any{"channel": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, channelURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteChannel(id string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", channelURL, id), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UpdateChannel(resource models.Channel) error {
	payload := map[string]any{"channel": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, channelURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UnsetChannel(resource models.Channel) error {
	payload := map[string]any{"channel": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", channelURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllChannel() ([]models.Channel, error) {
	req, err := s.client.NewRequest(http.MethodGet, channelURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.Channel `json:"channel"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetChannel(id string) (*models.Channel, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", channelURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.Channel `json:"channel"`
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

func (s *NetworkService) CountChannel() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", channelURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.Channel `json:"channel"`
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

// channel_binding
func (s *NetworkService) GetAllChannelBinding() ([]models.ChannelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, channelBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.ChannelBinding `json:"channel_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetChannelBinding(id string) (*models.ChannelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", channelBindingURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.ChannelBinding `json:"channel_binding"`
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

// channel_interface_binding
func (s *NetworkService) AddChannelInterfaceBinding(resource models.ChannelInterfaceBinding) error {
	payload := map[string]any{"channel_interface_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, channelInterfaceBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteChannelInterfaceBinding(id string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", channelInterfaceBindingURL, id), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllChannelInterfaceBinding() ([]models.ChannelInterfaceBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, channelInterfaceBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.ChannelInterfaceBinding `json:"channel_interface_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetChannelInterfaceBinding(id string) (*models.ChannelInterfaceBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", channelInterfaceBindingURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.ChannelInterfaceBinding `json:"channel_interface_binding"`
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

func (s *NetworkService) CountChannelInterfaceBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", channelInterfaceBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.ChannelInterfaceBinding `json:"channel_interface_binding"`
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

// ci
func (s *NetworkService) GetAllCI() ([]models.CI, error) {
	req, err := s.client.NewRequest(http.MethodGet, ciURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.CI `json:"ci"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) CountCI() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", ciURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.CI `json:"ci"`
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

// fis
func (s *NetworkService) AddFIS(resource models.FIS) error {
	payload := map[string]any{"fis": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fisURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteFIS(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", fisURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllFIS() ([]models.FIS, error) {
	req, err := s.client.NewRequest(http.MethodGet, fisURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.FIS `json:"fis"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetFIS(name string) (*models.FIS, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", fisURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.FIS `json:"fis"`
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

func (s *NetworkService) CountFIS() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", fisURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.FIS `json:"fis"`
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

// fis_binding
func (s *NetworkService) GetAllFISBinding() ([]models.FISBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fisBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.FISBinding `json:"fis_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetFISBinding(name string) (*models.FISBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", fisBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.FISBinding `json:"fis_binding"`
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

// fis_channel_binding
func (s *NetworkService) AddFISChannelBinding(resource models.FISChannelBinding) error {
	payload := map[string]any{"fis_channel_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fisChannelBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteFISChannelBinding(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", fisChannelBindingURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllFISChannelBinding() ([]models.FISChannelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fisChannelBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.FISChannelBinding `json:"fis_channel_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetFISChannelBinding(name string) (*models.FISChannelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", fisChannelBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.FISChannelBinding `json:"fis_channel_binding"`
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

func (s *NetworkService) CountFISChannelBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", fisChannelBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.FISChannelBinding `json:"fis_channel_binding"`
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

// fis_interface_binding
func (s *NetworkService) AddFISInterfaceBinding(resource models.FISInterfaceBinding) error {
	payload := map[string]any{"fis_interface_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fisInterfaceBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteFISInterfaceBinding(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", fisInterfaceBindingURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// forwardingsession
func (s *NetworkService) AddFowardingSession(resource models.ForwardingSession) error {
	payload := map[string]any{"forwardingsession": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, forwardingSessionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteFowardingSession(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", forwardingSessionURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UpdateFowardingSession(resource models.ForwardingSession) error {
	payload := map[string]any{"forwardingsession": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, forwardingSessionURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllFowardingSession() ([]models.ForwardingSession, error) {
	req, err := s.client.NewRequest(http.MethodGet, forwardingSessionURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.ForwardingSession `json:"forwardingsession"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetFowardingSession(name string) (*models.ForwardingSession, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", forwardingSessionURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.ForwardingSession `json:"forwardingsession"`
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

func (s *NetworkService) CountFowardingSession() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", forwardingSessionURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.ForwardingSession `json:"forwardingsession"`
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

// inat
func (s *NetworkService) AddINAT(resource models.INAT) error {
	payload := map[string]any{"inat": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, inatURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteINAT(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", inatURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UpdateINAT(resource models.INAT) error {
	payload := map[string]any{"inat": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, inatURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UnsetINAT(resource models.INAT) error {
	payload := map[string]any{"inat": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", inatURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllINAT() ([]models.INAT, error) {
	req, err := s.client.NewRequest(http.MethodGet, inatURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.INAT `json:"inat"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetINAT(name string) (*models.INAT, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", inatURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.INAT `json:"inat"`
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

func (s *NetworkService) CountINAT() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", inatURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.INAT `json:"inat"`
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

// inatparam
func (s *NetworkService) UpdateINATParam(resource models.INATParam) error {
	payload := map[string]any{"inatparam": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, inatParamURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UnsetINATParam(resource models.INATParam) error {
	payload := map[string]any{"inatparam": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", inatParamURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllINATParam() ([]models.INATParam, error) {
	req, err := s.client.NewRequest(http.MethodGet, inatParamURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.INATParam `json:"inatparam"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetINATParam(name string) (*models.INATParam, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", inatParamURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.INATParam `json:"inatparam"`
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

func (s *NetworkService) CountINATParam() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", inatParamURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.INATParam `json:"inatparam"`
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

// interface
func (s *NetworkService) ClearInterface(id string) error {
	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=clear", interfaceURL), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UpdateInterface(resource models.Interface) error {
	payload := map[string]any{"interface": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, interfaceURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UnsetInterface(resource models.Interface) error {
	payload := map[string]any{"interface": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", interfaceURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) EnableInterface(resource models.Interface) error {
	payload := map[string]any{"interface": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=enable", interfaceURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DisableInterface(resource models.Interface) error {
	payload := map[string]any{"interface": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=disable", interfaceURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) ResetInterface(resource models.Interface) error {
	payload := map[string]any{"interface": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=reset", interfaceURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllInterface() ([]models.Interface, error) {
	req, err := s.client.NewRequest(http.MethodGet, interfaceURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.Interface `json:"interface"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetInterface(id string) (*models.Interface, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", interfaceURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.Interface `json:"interface"`
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

func (s *NetworkService) CountInterface() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", interfaceURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.Interface `json:"interface"`
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

// interfacepair
func (s *NetworkService) AddInterfacePair(resource models.InterfacePair) error {
	payload := map[string]any{"interfacepair": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, interfacePairURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteInterfacePair(id string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", interfacePairURL, id), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllInterfacePair() ([]models.InterfacePair, error) {
	req, err := s.client.NewRequest(http.MethodGet, interfacePairURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.InterfacePair `json:"interfacepair"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetInterfacePair(id string) (*models.InterfacePair, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", interfacePairURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.InterfacePair `json:"interfacepair"`
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

func (s *NetworkService) CountInterfacePair() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", interfacePairURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.InterfacePair `json:"interfacepair"`
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

// ip6tunnel
func (s *NetworkService) AddIP6Tunnel(resource models.IP6Tunnel) error {
	payload := map[string]any{"ip6tunnel": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, ip6TunnelURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteIP6Tunnel(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", ip6TunnelURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllIP6Tunnel() ([]models.IP6Tunnel, error) {
	req, err := s.client.NewRequest(http.MethodGet, ip6TunnelURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.IP6Tunnel `json:"ip6tunnel"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetIP6Tunnel(name string) (*models.IP6Tunnel, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", ip6TunnelURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.IP6Tunnel `json:"ip6tunnel"`
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

func (s *NetworkService) CountIP6Tunnel() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", ip6TunnelURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.IP6Tunnel `json:"ip6tunnel"`
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

// ip6tunnelparam
func (s *NetworkService) UpdateIP6TunnelParam(resource models.IP6TunnelParam) error {
	payload := map[string]any{"ip6tunnelparam": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, ip6TunnelParamURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UnsetIP6TunnelParam(resource models.IP6TunnelParam) error {
	payload := map[string]any{"ip6tunnelparam": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", ip6TunnelParamURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllIP6TunnelParam() ([]models.IP6TunnelParam, error) {
	req, err := s.client.NewRequest(http.MethodGet, ip6TunnelParamURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.IP6TunnelParam `json:"ip6tunnelparam"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// ipset
func (s *NetworkService) AddIPSet(resource models.IPSet) error {
	payload := map[string]any{"ipset": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, ipSetURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteIPSet(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", ipSetURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllIPSet() ([]models.IPSet, error) {
	req, err := s.client.NewRequest(http.MethodGet, ipSetURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.IPSet `json:"ipset"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetIPSet(name string) (*models.IPSet, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", ipSetURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.IPSet `json:"ipset"`
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

func (s *NetworkService) CountIPSet() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", ipSetURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.IPSet `json:"ipset"`
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

// ipset_binding
func (s *NetworkService) GetAllIPSetBinding() ([]models.IPSetBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, ipSetBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.IPSetBinding `json:"ipset_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetIPSetBinding(name string) (*models.IPSetBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", ipSetBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.IPSetBinding `json:"ipset_binding"`
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

// ipset_nsip6_binding
func (s *NetworkService) AddIPSetNSIP6Binding(resource models.IPSetNSIP6Binding) error {
	payload := map[string]any{"ipset_nsip6_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, ipSetNSIP6BindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteIPSetNSIP6Binding(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", ipSetNSIP6BindingURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllIPSetNSIP6Binding() ([]models.IPSetNSIP6Binding, error) {
	req, err := s.client.NewRequest(http.MethodGet, ipSetNSIP6BindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.IPSetNSIP6Binding `json:"ipset_nsip6_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetIPSetNSIP6Binding(name string) (*models.IPSetNSIP6Binding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", ipSetNSIP6BindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.IPSetNSIP6Binding `json:"ipset_nsip6_binding"`
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

func (s *NetworkService) CountIPSetNSIP6Binding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", ipSetNSIP6BindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.IPSetNSIP6Binding `json:"ipset_nsip6_binding"`
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

// ipset_nsip_binding
func (s *NetworkService) AddIPSetNSIPBinding(resource models.IPSetNSIPBinding) error {
	payload := map[string]any{"ipset_nsip_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, ipSetNSIPBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteIPSetNSIPBinding(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", ipSetNSIPBindingURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllIPSetNSIPBinding() ([]models.IPSetNSIPBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, ipSetNSIPBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.IPSetNSIPBinding `json:"ipset_nsip_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetIPSetNSIPBinding(name string) (*models.IPSetNSIPBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", ipSetNSIPBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.IPSetNSIPBinding `json:"ipset_nsip_binding"`
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

func (s *NetworkService) CountIPSetNSIPBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", ipSetNSIPBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.IPSetNSIPBinding `json:"ipset_nsip_binding"`
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

// iptunnel
func (s *NetworkService) AddIPTunnel(resource models.IPTunnel) error {
	payload := map[string]any{"iptunnel": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, ipTunnelURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteIPTunnel(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", ipTunnelURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UpdateIPTunnel(resource models.IPTunnel) error {
	payload := map[string]any{"iptunnel": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, ipTunnelURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UnsetIPTunnel(resource models.IPTunnel) error {
	payload := map[string]any{"iptunnel": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", ipTunnelURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllIPTunnel() ([]models.IPTunnel, error) {
	req, err := s.client.NewRequest(http.MethodGet, ipTunnelURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.IPTunnel `json:"iptunnel"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetIPTunnel(name string) (*models.IPTunnel, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", ipTunnelURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.IPTunnel `json:"iptunnel"`
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

func (s *NetworkService) CountIPTunnel() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", ipTunnelURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.IPTunnel `json:"iptunnel"`
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

// iptunnelparam
func (s *NetworkService) UpdateIPTunnelParam(resource models.IPTunnelParam) error {
	payload := map[string]any{"iptunnelparam": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, ipTunnelParamURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UnsetIPTunnelParam(resource models.IPTunnelParam) error {
	payload := map[string]any{"iptunnelparam": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", ipTunnelParamURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllIPTunnelParam() ([]models.IPTunnelParam, error) {
	req, err := s.client.NewRequest(http.MethodGet, ipTunnelParamURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.IPTunnelParam `json:"iptunnelparam"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// ipv6
func (s *NetworkService) UpdateIPV6(resource models.IPv6) error {
	payload := map[string]any{"ipv6": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, ipv6URL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UnsetIPV6(resource models.IPv6) error {
	payload := map[string]any{"ipv6": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", ipv6URL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllIPV6() ([]models.IPv6, error) {
	req, err := s.client.NewRequest(http.MethodGet, ipv6URL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.IPv6 `json:"ipv6"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetIPV6(dodad string) (*models.IPv6, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", ipv6URL, dodad), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.IPv6 `json:"ipv6"`
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

func (s *NetworkService) CountIPV6() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", ipv6URL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.IPv6 `json:"ipv6"`
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

// l2param
func (s *NetworkService) UpdateL2Param(resource models.L2Param) error {
	payload := map[string]any{"l2param": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, l2ParamURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UnsetL2Param(resource models.L2Param) error {
	payload := map[string]any{"l2param": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", l2ParamURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllL2Param() ([]models.L2Param, error) {
	req, err := s.client.NewRequest(http.MethodGet, l2ParamURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.L2Param `json:"l2param"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// l3param
func (s *NetworkService) UpdateL3Param(resource models.L3Param) error {
	payload := map[string]any{"l3param": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, l3ParamURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UnsetL3Param(resource models.L3Param) error {
	payload := map[string]any{"l3param": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", l3ParamURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllL3Param() ([]models.L3Param, error) {
	req, err := s.client.NewRequest(http.MethodGet, l3ParamURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.L3Param `json:"l3param"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// l4param
func (s *NetworkService) UpdateL4Param(resource models.L4Param) error {
	payload := map[string]any{"l4param": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, l4ParamURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UnsetL4Param(resource models.L4Param) error {
	payload := map[string]any{"l4param": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", l4ParamURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllL4Param() ([]models.L4Param, error) {
	req, err := s.client.NewRequest(http.MethodGet, l4ParamURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.L4Param `json:"l4param"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// lacp
func (s *NetworkService) UpdateLACP(resource models.LACP) error {
	payload := map[string]any{"lacp": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, lacpURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllLACP() ([]models.LACP, error) {
	req, err := s.client.NewRequest(http.MethodGet, lacpURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LACP `json:"lacp"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetLACP(syspriority string) (*models.LACP, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", lacpURL, syspriority), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LACP `json:"lacp"`
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

func (s *NetworkService) CountLACP() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", lacpURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LACP `json:"lacp"`
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

// linkset
func (s *NetworkService) AddLinkSet(resource models.LinkSet) error {
	payload := map[string]any{"linkset": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, linkSetURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteLinkSet(id string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", linkSetURL, id), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllLinkSet() ([]models.LinkSet, error) {
	req, err := s.client.NewRequest(http.MethodGet, linkSetURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LinkSet `json:"linkset"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetLinkSet(id string) (*models.LinkSet, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", linkSetURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LinkSet `json:"linkset"`
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

func (s *NetworkService) CountLinkSet() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", linkSetURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LinkSet `json:"linkset"`
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

// linkset_binding
func (s *NetworkService) GetAllLinkSetBinding() ([]models.LinkSetBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, linkSetBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LinkSetBinding `json:"linkset_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetLinkSetBinding(id string) (*models.LinkSetBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", linkSetBindingURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LinkSetBinding `json:"linkset_binding"`
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

// linkset_channel_binding
func (s *NetworkService) AddLinkSetChannelBinding(resource models.LinkSetChannelBinding) error {
	payload := map[string]any{"linkset_channel_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, linkSetChannelBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteLinkSetChannelBinding(id string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", linkSetChannelBindingURL, id), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllLinkSetChannelBinding() ([]models.LinkSetChannelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, linkSetChannelBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LinkSetChannelBinding `json:"linkset_channel_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetLinkSetChannelBinding(id string) (*models.LinkSetChannelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", linkSetChannelBindingURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LinkSetChannelBinding `json:"linkset_channel_binding"`
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

func (s *NetworkService) CountLinkSetChannelBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", linkSetChannelBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LinkSetChannelBinding `json:"linkset_channel_binding"`
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

// linkset_interface_binding
func (s *NetworkService) AddLinkSetInterfaceBinding(resource models.LinkSetInterfaceBinding) error {
	payload := map[string]any{"linkset_interface_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, linkSetInterfaceBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteLinkSetInterfaceBinding(id string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", linkSetInterfaceBindingURL, id), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllLinkSetInterfaceBinding() ([]models.LinkSetInterfaceBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, linkSetInterfaceBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LinkSetInterfaceBinding `json:"linkset_interface_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetLinkSetInterfaceBinding(id string) (*models.LinkSetInterfaceBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", linkSetInterfaceBindingURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.LinkSetInterfaceBinding `json:"linkset_interface_binding"`
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

func (s *NetworkService) CountLinkSetInterfaceBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", linkSetInterfaceBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.LinkSetInterfaceBinding `json:"linkset_interface_binding"`
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

// mapbmr
func (s *NetworkService) AddMAPBMR(resource models.MapBMR) error {
	payload := map[string]any{"mapbmr": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, mapBMRURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteMAPBMR(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", mapBMRURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllMAPBMR() ([]models.MapBMR, error) {
	req, err := s.client.NewRequest(http.MethodGet, mapBMRURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.MapBMR `json:"mapbmr"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetMAPBMR(name string) (*models.MapBMR, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", mapBMRURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.MapBMR `json:"mapbmr"`
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

func (s *NetworkService) CountMAPBMR() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", mapBMRURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.MapBMR `json:"mapbmr"`
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

// mapbmr_binding
func (s *NetworkService) GetAllMAPBMRBinding() ([]models.MapBMRBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, mapBMRBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.MapBMRBinding `json:"mapbmr_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetMAPBMRBinding(name string) (*models.MapBMRBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", mapBMRBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.MapBMRBinding `json:"mapbmr_binding"`
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

// mapbmr_bmrv4network_binding
func (s *NetworkService) AddMAPBMRBMRV4NetworkBinding(resource models.MapBMRBMRV4NetworkBinding) error {
	payload := map[string]any{"mapbmr_bmrv4network_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, mapBMRBMRV4NetworkBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteMAPBMRBMRV4NetworkBinding(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", mapBMRBMRV4NetworkBindingURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllMAPBMRBMRV4NetworkBinding() ([]models.MapBMRBMRV4NetworkBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, mapBMRBMRV4NetworkBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.MapBMRBMRV4NetworkBinding `json:"mapbmr_bmrv4network_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetMAPBMRBMRV4NetworkBinding(name string) (*models.MapBMRBMRV4NetworkBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", mapBMRBMRV4NetworkBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.MapBMRBMRV4NetworkBinding `json:"mapbmr_bmrv4network_binding"`
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

func (s *NetworkService) CountMAPBMRBMRV4NetworkBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", mapBMRBMRV4NetworkBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.MapBMRBMRV4NetworkBinding `json:"mapbmr_bmrv4network_binding"`
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

// mapdmr
func (s *NetworkService) AddMAPDMR(resource models.MapDMR) error {
	payload := map[string]any{"mapdmr": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, mapDMRURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteMAPDMR(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", mapDMRURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllMAPDMR() ([]models.MapDMR, error) {
	req, err := s.client.NewRequest(http.MethodGet, mapDMRURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.MapDMR `json:"mapdmr"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetMAPDMR(name string) (*models.MapDMR, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", mapDMRURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.MapDMR `json:"mapdmr"`
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
func (s *NetworkService) CountMAPDMR() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", mapDMRURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.MapDMR `json:"mapdmr"`
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

// mapdomain
func (s *NetworkService) AddMAPDomain(resource models.MapDomain) error {
	payload := map[string]any{"mapdomain": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, mapDomainURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteMAPDomain(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", mapDomainURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllMAPDomain() ([]models.MapDomain, error) {
	req, err := s.client.NewRequest(http.MethodGet, mapDomainURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.MapDomain `json:"mapdomain"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetMAPDomain(name string) (*models.MapDomain, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", mapDomainURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.MapDomain `json:"mapdomain"`
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

func (s *NetworkService) CountMAPDomain() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", mapDomainURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.MapDomain `json:"mapdomain"`
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

// mapdomain_binding
func (s *NetworkService) GetAllMAPDomainBinding() ([]models.MapDomainBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, mapDomainBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.MapDomainBinding `json:"mapdomain_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetMAPDomainBinding(name string) (*models.MapDomainBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", mapDomainBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.MapDomainBinding `json:"mapdomain_binding"`
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

// mapdomain_mapbmr_binding
func (s *NetworkService) AddMAPDomainMAPBMRBinding(resource models.MapDomainMapBMRBinding) error {
	payload := map[string]any{"mapdomain_mapbmr_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, mapDomainMapBMRBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteMAPDomainMAPBMRBinding(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", mapDomainMapBMRBindingURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllMAPDomainMAPBMRBinding() ([]models.MapDomainMapBMRBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, mapDomainMapBMRBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.MapDomainMapBMRBinding `json:"mapdomain_mapbmr_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetMAPDomainMAPBMRBinding(name string) (*models.MapDomainMapBMRBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", mapDomainMapBMRBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.MapDomainMapBMRBinding `json:"mapdomain_mapbmr_binding"`
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

func (s *NetworkService) CountMAPDomainMAPBMRBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", mapDomainMapBMRBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.MapDomainMapBMRBinding `json:"mapdomain_mapbmr_binding"`
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

// nat64
func (s *NetworkService) AddNAT64(resource models.NAT64) error {
	payload := map[string]any{"nat64": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nat64URL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteNAT64(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nat64URL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UpdateNAT64(resource models.NAT64) error {
	payload := map[string]any{"nat64": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nat64URL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UnsetNAT64(resource models.NAT64) error {
	payload := map[string]any{"nat64": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", nat64URL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllNAT64() ([]models.NAT64, error) {
	req, err := s.client.NewRequest(http.MethodGet, nat64URL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NAT64 `json:"nat64"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetNAT64(name string) (*models.NAT64, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nat64URL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NAT64 `json:"nat64"`
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

func (s *NetworkService) CountNAT64() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nat64URL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NAT64 `json:"nat64"`
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

// nat64param
func (s *NetworkService) UpdateNAT64Param(resource models.NAT64Param) error {
	payload := map[string]any{"nat64param": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nat64ParamURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UnsetNAT64Param(resource models.NAT64Param) error {
	payload := map[string]any{"nat64param": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", nat64ParamURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllNAT64Param() ([]models.NAT64Param, error) {
	req, err := s.client.NewRequest(http.MethodGet, nat64ParamURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NAT64Param `json:"nat64param"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetNAT64Param() (*models.NAT64Param, error) {
	req, err := s.client.NewRequest(http.MethodGet, nat64ParamURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NAT64Param `json:"nat64param"`
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

func (s *NetworkService) CountNAT64Param() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nat64ParamURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NAT64Param `json:"nat64param"`
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

// nd6
func (s *NetworkService) AddND6(resource models.ND6) error {
	payload := map[string]any{"nd6": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nd6URL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteND6(neighbor string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nd6URL, neighbor), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) ClearND6(neighbor string) error {
	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=clear", nd6URL), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllND6() ([]models.ND6, error) {
	req, err := s.client.NewRequest(http.MethodGet, nd6URL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.ND6 `json:"nd6"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}
func (s *NetworkService) CountND6() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nd6URL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.ND6 `json:"nd6"`
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

// nd6ravariables
func (s *NetworkService) UpdateND6RAVariables(resource models.ND6RAVariables) error {
	payload := map[string]any{"nd6ravariables": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, nd6RAVariablesURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UnsetND6RAVariables(resource models.ND6RAVariables) error {
	payload := map[string]any{"nd6ravariables": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", nd6RAVariablesURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllND6RAVariables() ([]models.ND6RAVariables, error) {
	req, err := s.client.NewRequest(http.MethodGet, nd6RAVariablesURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.ND6RAVariables `json:"nd6ravariables"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetND6RAVariables(vlan string) (*models.ND6RAVariables, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nd6RAVariablesURL, vlan), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.ND6RAVariables `json:"nd6ravariables"`
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

func (s *NetworkService) CountND6RAVariables() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nd6RAVariablesURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.ND6RAVariables `json:"nd6ravariables"`
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

// nd6ravariables_binding
func (s *NetworkService) GetAllND6RAVariablesBinding() ([]models.ND6RAVariablesBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, nd6RAVariablesBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.ND6RAVariablesBinding `json:"nd6ravariables_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetND6RAVariablesBinding(vlan string) (*models.ND6RAVariablesBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nd6RAVariablesBindingURL, vlan), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.ND6RAVariablesBinding `json:"nd6ravariables_binding"`
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

// nd6ravariables_onlinkipv6prefix_binding
func (s *NetworkService) AddND6RAVariablesOnlinkIPV6PrefixBinding(resource models.ND6RAVariablesOnLinkIPv6PrefixBinding) error {
	payload := map[string]any{"nd6ravariables_onlinkipv6prefix_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, nd6RAVariablesOnLinkIPV6PrefixBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteND6RAVariablesOnlinkIPV6PrefixBinding(vlan string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", nd6RAVariablesOnLinkIPV6PrefixBindingURL, vlan), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllND6RAVariablesOnlinkIPV6PrefixBinding() ([]models.ND6RAVariablesOnLinkIPv6PrefixBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, nd6RAVariablesOnLinkIPV6PrefixBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.ND6RAVariablesOnLinkIPv6PrefixBinding `json:"nd6ravariables_onlinkipv6prefix_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetND6RAVariablesOnlinkIPV6PrefixBinding(vlan string) (*models.ND6RAVariablesOnLinkIPv6PrefixBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", nd6RAVariablesOnLinkIPV6PrefixBindingURL, vlan), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.ND6RAVariablesOnLinkIPv6PrefixBinding `json:"nd6ravariables_onlinkipv6prefix_binding"`
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

func (s *NetworkService) CountND6RAVariablesOnlinkIPV6PrefixBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", nd6RAVariablesOnLinkIPV6PrefixBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.ND6RAVariablesOnLinkIPv6PrefixBinding `json:"nd6ravariables_onlinkipv6prefix_binding"`
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

// netbridge
func (s *NetworkService) AddNetBridge(resource models.NetBridge) error {
	payload := map[string]any{"netbridge": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, netBridgeURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteNetBridge(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", netBridgeURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UpdateNetBridge(resource models.NetBridge) error {
	payload := map[string]any{"netbridge": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, netBridgeURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UnsetNetBridge(resource models.NetBridge) error {
	payload := map[string]any{"netbridge": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", netBridgeURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllNetBridge() ([]models.NetBridge, error) {
	req, err := s.client.NewRequest(http.MethodGet, netBridgeURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NetBridge `json:"netbridge"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetNetBridge(name string) (*models.NetBridge, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", netBridgeURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NetBridge `json:"netbridge"`
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

func (s *NetworkService) CountNetBridge() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", netBridgeURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NetBridge `json:"netbridge"`
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

// netbridge_binding
func (s *NetworkService) GetAllNetBridgeBinding() ([]models.NetBridgeBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, netBridgeBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NetBridgeBinding `json:"netbridge_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetNetBridgeBinding(name string) (*models.NetBridgeBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", netBridgeBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NetBridgeBinding `json:"netbridge_binding"`
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

// netbridge_iptunnel_binding
func (s *NetworkService) AddNetBridgeIPTunnelBinding(resource models.NetBridgeIPTunnelBinding) error {
	payload := map[string]any{"netbridge_iptunnel_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, netBridgeIPTunnelBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteNetBridgeIPTunnelBinding(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", netBridgeIPTunnelBindingURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllNetBridgeIPTunnelBinding() ([]models.NetBridgeIPTunnelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, netBridgeIPTunnelBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NetBridgeIPTunnelBinding `json:"netbridge_iptunnel_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetNetBridgeIPTunnelBinding(name string) (*models.NetBridgeIPTunnelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", netBridgeIPTunnelBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NetBridgeIPTunnelBinding `json:"netbridge_iptunnel_binding"`
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

func (s *NetworkService) CountNetBridgeIPTunnelBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", netBridgeIPTunnelBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NetBridgeIPTunnelBinding `json:"netbridge_iptunnel_binding"`
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

// netbridge_nsip6_binding
func (s *NetworkService) AddNetBridgeNSIP6Binding(resource models.NetBridgeNSIP6Binding) error {
	payload := map[string]any{"netbridge_nsip6_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, netBridgeNSIP6BindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteNetBridgeNSIP6Binding(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", netBridgeNSIP6BindingURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllNetBridgeNSIP6Binding() ([]models.NetBridgeNSIP6Binding, error) {
	req, err := s.client.NewRequest(http.MethodGet, netBridgeNSIP6BindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NetBridgeNSIP6Binding `json:"netbridge_nsip6_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetNetBridgeNSIP6Binding(name string) (*models.NetBridgeNSIP6Binding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", netBridgeNSIP6BindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NetBridgeNSIP6Binding `json:"netbridge_nsip6_binding"`
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

func (s *NetworkService) CountNetBridgeNSIP6Binding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", netBridgeNSIP6BindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NetBridgeNSIP6Binding `json:"netbridge_nsip6_binding"`
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

// netbridge_nsip_binding
func (s *NetworkService) AddNetBridgeNSIPBinding(resource models.NetBridgeNSIPBinding) error {
	payload := map[string]any{"netbridge_nsip_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, netBridgeNSIPBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteNetBridgeNSIPBinding(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", netBridgeNSIPBindingURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllNetBridgeNSIPBinding() ([]models.NetBridgeNSIPBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, netBridgeNSIPBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NetBridgeNSIPBinding `json:"netbridge_nsip_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetNetBridgeNSIPBinding(name string) (*models.NetBridgeNSIPBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", netBridgeNSIPBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NetBridgeNSIPBinding `json:"netbridge_nsip_binding"`
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

func (s *NetworkService) CountNetBridgeNSIPBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", netBridgeNSIPBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NetBridgeNSIPBinding `json:"netbridge_nsip_binding"`
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

// netbridge_vlan_binding
func (s *NetworkService) AddNetBridgeVLANBinding(resource models.NetBridgeVLANBinding) error {
	payload := map[string]any{"netbridge_vlan_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, netBridgeVLANBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteNetBridgeVLANBinding(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", netBridgeVLANBindingURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllNetBridgeVLANBinding() ([]models.NetBridgeVLANBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, netBridgeVLANBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NetBridgeVLANBinding `json:"netbridge_vlan_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetNetBridgeVLANBinding(name string) (*models.NetBridgeVLANBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", netBridgeVLANBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NetBridgeVLANBinding `json:"netbridge_vlan_binding"`
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

func (s *NetworkService) CountNetBridgeVLANBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", netBridgeVLANBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NetBridgeVLANBinding `json:"netbridge_vlan_binding"`
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

// netprofile
func (s *NetworkService) AddNetProfile(resource models.NetProfile) error {
	payload := map[string]any{"netprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, netProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteNetProfile(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", netProfileURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UpdateNetProfile(resource models.NetProfile) error {
	payload := map[string]any{"netprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, netProfileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UnsetNetProfile(resource models.NetProfile) error {
	payload := map[string]any{"netprofile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", netProfileURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllNetProfile() ([]models.NetProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, netProfileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NetProfile `json:"netprofile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetNetProfile(name string) (*models.NetProfile, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", netProfileURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NetProfile `json:"netprofile"`
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

func (s *NetworkService) CountNetProfile() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", netProfileURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NetProfile `json:"netprofile"`
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

// netprofile_binding
func (s *NetworkService) GetAllNetProfileBinding() ([]models.NetProfileBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, netProfileBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NetProfileBinding `json:"netprofile_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetNetProfileBinding(name string) (*models.NetProfileBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", netProfileBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NetProfileBinding `json:"netprofile_binding"`
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

// netprofile_natrule_binding
func (s *NetworkService) AddNetProfileNATRuleBinding(resource models.NetProfileNATRuleBinding) error {
	payload := map[string]any{"netprofile_natrule_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, netProfileNATRuleBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteNetProfileNATRuleBinding(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", netProfileNATRuleBindingURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllNetProfileNATRuleBinding() ([]models.NetProfileNATRuleBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, netProfileNATRuleBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NetProfileNATRuleBinding `json:"netprofile_natrule_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetNetProfileNATRuleBinding(name string) (*models.NetProfileNATRuleBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", netProfileNATRuleBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NetProfileNATRuleBinding `json:"netprofile_natrule_binding"`
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

func (s *NetworkService) CountNetProfileNATRuleBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", netProfileNATRuleBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NetProfileNATRuleBinding `json:"netprofile_natrule_binding"`
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

// netprofile_srcportset_binding
func (s *NetworkService) AddNetProfileSrcPortSetBinding(resource models.NetProfileSrcPortSetBinding) error {
	payload := map[string]any{"netprofile_srcportset_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, netProfileSrcPortSetBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteNetProfileSrcPortSetBinding(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", netProfileSrcPortSetBindingURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllNetProfileSrcPortSetBinding() ([]models.NetProfileSrcPortSetBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, netProfileSrcPortSetBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NetProfileSrcPortSetBinding `json:"netprofile_srcportset_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetNetProfileSrcPortSetBinding(name string) (*models.NetProfileSrcPortSetBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", netProfileSrcPortSetBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.NetProfileSrcPortSetBinding `json:"netprofile_srcportset_binding"`
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

func (s *NetworkService) CountNetProfileSrcPortSetBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", netProfileSrcPortSetBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.NetProfileSrcPortSetBinding `json:"netprofile_srcportset_binding"`
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

// onlinkipv6prefix
func (s *NetworkService) AddOnlinkIPV6Prefix(resource models.OnLinkIPv6Prefix) error {
	payload := map[string]any{"onlinkipv6prefix": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, onLinkIPV6PrefixURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteOnlinkIPV6Prefix(ipv6prefix string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", onLinkIPV6PrefixURL, ipv6prefix), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UpdateOnlinkIPV6Prefix(resource models.OnLinkIPv6Prefix) error {
	payload := map[string]any{"onlinkipv6prefix": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, onLinkIPV6PrefixURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UnsetOnlinkIPV6Prefix(resource models.OnLinkIPv6Prefix) error {
	payload := map[string]any{"onlinkipv6prefix": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", onLinkIPV6PrefixURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllOnlinkIPV6Prefix() ([]models.OnLinkIPv6Prefix, error) {
	req, err := s.client.NewRequest(http.MethodGet, onLinkIPV6PrefixURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.OnLinkIPv6Prefix `json:"onlinkipv6prefix"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetOnlinkIPV6Prefix(ipv6prefix string) (*models.OnLinkIPv6Prefix, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", onLinkIPV6PrefixURL, ipv6prefix), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.OnLinkIPv6Prefix `json:"onlinkipv6prefix"`
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

func (s *NetworkService) CountOnlinkIPV6Prefix() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", onLinkIPV6PrefixURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.OnLinkIPv6Prefix `json:"onlinkipv6prefix"`
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

// ptp
func (s *NetworkService) UpdatePTP(resource models.PTP) error {
	payload := map[string]any{"ptp": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, ptpURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllPTP() ([]models.PTP, error) {
	req, err := s.client.NewRequest(http.MethodGet, ptpURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.PTP `json:"ptp"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// rnat
func (s *NetworkService) AddRNAT(resource models.RNAT) error {
	payload := map[string]any{"rnat": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, rnatURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteRNAT(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", rnatURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UpdateRNAT(resource models.RNAT) error {
	payload := map[string]any{"rnat": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, rnatURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UnsetRNAT(resource models.RNAT) error {
	payload := map[string]any{"rnat": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", rnatURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) ClearRNAT() error {
	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=clear", rnatURL), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) RenameRNAT(resource models.RNAT) error {
	payload := map[string]any{"rnat": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=rename", rnatURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllRNAT() ([]models.RNAT, error) {
	req, err := s.client.NewRequest(http.MethodGet, rnatURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.RNAT `json:"rnat"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetRNAT(name string) (*models.RNAT, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", rnatURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.RNAT `json:"rnat"`
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

func (s *NetworkService) CountRNAT() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", rnatURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.RNAT `json:"rnat"`
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

// rnat6
func (s *NetworkService) AddRNAT6(resource models.RNAT6) error {
	payload := map[string]any{"rnat6": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, rnat6URL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteRNAT6(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", rnat6URL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UpdateRNAT6(resource models.RNAT6) error {
	payload := map[string]any{"rnat6": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, rnat6URL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UnsetRNAT6(resource models.RNAT6) error {
	payload := map[string]any{"rnat6": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", rnat6URL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllRNAT6() ([]models.RNAT6, error) {
	req, err := s.client.NewRequest(http.MethodGet, rnat6URL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.RNAT6 `json:"rnat6"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetRNAT6(name string) (*models.RNAT6, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", rnat6URL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.RNAT6 `json:"rnat6"`
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

func (s *NetworkService) CountRNAT6() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", rnat6URL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.RNAT6 `json:"rnat6"`
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

// rnat6_binding
func (s *NetworkService) GetAllRNAT6Binding() ([]models.RNAT6Binding, error) {
	req, err := s.client.NewRequest(http.MethodGet, rnat6BindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.RNAT6Binding `json:"rnat6_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetRNAT6Binding(name string) (*models.RNAT6Binding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", rnat6BindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.RNAT6Binding `json:"rnat6_binding"`
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

// rnat6_nsip6_binding
func (s *NetworkService) AddRNAT6NSIP6Binding(resource models.RNAT6NSIP6Binding) error {
	payload := map[string]any{"rnat6_nsip6_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, rnat6NSIP6BindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteRNAT6NSIP6Binding(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", rnat6NSIP6BindingURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllRNAT6NSIP6Binding() ([]models.RNAT6NSIP6Binding, error) {
	req, err := s.client.NewRequest(http.MethodGet, rnat6NSIP6BindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.RNAT6NSIP6Binding `json:"rnat6_nsip6_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetRNAT6NSIP6Binding(name string) (*models.RNAT6NSIP6Binding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", rnat6NSIP6BindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.RNAT6NSIP6Binding `json:"rnat6_nsip6_binding"`
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

func (s *NetworkService) CountRNAT6NSIP6Binding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", rnat6NSIP6BindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.RNAT6NSIP6Binding `json:"rnat6_nsip6_binding"`
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

// rnatglobal_auditsyslogpolicy_binding
func (s *NetworkService) AddRNATGlobalAuditSyslogPolicyBinding(resource models.RNATGlobalAuditSyslogPolicyBinding) error {
	payload := map[string]any{"rnatglobal_auditsyslogpolicy_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, rnatGlobalAuditSyslogPolicyBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteRNATGlobalAuditSyslogPolicyBinding(policy string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", rnatGlobalAuditSyslogPolicyBindingURL, policy), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetRNATGlobalAuditSyslogPolicyBinding() ([]models.RNATGlobalAuditSyslogPolicyBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, rnatGlobalAuditSyslogPolicyBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.RNATGlobalAuditSyslogPolicyBinding `json:"rnatglobal_auditsyslogpolicy_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) CountRNATGlobalAuditSyslogPolicyBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", rnatGlobalAuditSyslogPolicyBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.RNATGlobalAuditSyslogPolicyBinding `json:"rnatglobal_auditsyslogpolicy_binding"`
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

// rnatglobal_binding
func (s *NetworkService) GetRBATGlobalBinding() ([]models.RNATGlobalBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, rnatGlobalBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.RNATGlobalBinding `json:"rnatglobal_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// rnatparam
func (s *NetworkService) UpdateRNATParam(resource models.RNATParam) error {
	payload := map[string]any{"rnatparam": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, rnatParamURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UnsetRNATParam(resource models.RNATParam) error {
	payload := map[string]any{"rnatparam": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", rnatParamURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllRNATParam() ([]models.RNATParam, error) {
	req, err := s.client.NewRequest(http.MethodGet, rnatParamURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.RNATParam `json:"rnatparam"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// rnatsession
func (s *NetworkService) FlushRNATSession(resource models.RNATSession) error {
	payload := map[string]any{"rnatsession": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=flush", rnatSessionURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// rnat_binding
func (s *NetworkService) GetAllRNATBinding() ([]models.RNATBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, rnatBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.RNATBinding `json:"rnat_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetRNATBinding(name string) (*models.RNATBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", rnatBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.RNATBinding `json:"rnat_binding"`
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

// rnat_nsip_binding
func (s *NetworkService) AddRNATNSIPBinding(resource models.RNATNSIPBinding) error {
	payload := map[string]any{"rnat_nsip_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, rnatNSIPBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteRNATNSIPBinding(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", rnatNSIPBindingURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllRNATNSIPBinding() ([]models.RNATNSIPBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, rnatNSIPBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.RNATNSIPBinding `json:"rnat_nsip_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetRNATNSIPBinding(name string) (*models.RNATNSIPBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", rnatNSIPBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.RNATNSIPBinding `json:"rnat_nsip_binding"`
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

func (s *NetworkService) CountRNATNSIPBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", rnatNSIPBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.RNATNSIPBinding `json:"rnat_nsip_binding"`
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

// rnat_retainsourceportset_binding
func (s *NetworkService) AddRNATRetainSourcePortSetBinding(resource models.RNATRetainSourcePortSetBinding) error {
	payload := map[string]any{"rnat_retainsourceportset_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, rnatRetainSourcePortSetBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteRNATRetainSourcePortSetBinding(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", rnatRetainSourcePortSetBindingURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllRNATRetainSourcePortSetBinding() ([]models.RNATRetainSourcePortSetBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, rnatRetainSourcePortSetBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.RNATRetainSourcePortSetBinding `json:"rnat_retainsourceportset_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetRNATRetainSourcePortSetBinding(name string) (*models.RNATRetainSourcePortSetBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", rnatRetainSourcePortSetBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.RNATRetainSourcePortSetBinding `json:"rnat_retainsourceportset_binding"`
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

func (s *NetworkService) CountRNATRetainSourcePortSetBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", rnatRetainSourcePortSetBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.RNATRetainSourcePortSetBinding `json:"rnat_retainsourceportset_binding"`
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

// route
func (s *NetworkService) AddRoute(resource models.Route) error {
	payload := map[string]any{"route": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, routeURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteRoute(network string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", routeURL, network), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UpdateRoute(resource models.Route) error {
	payload := map[string]any{"route": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, routeURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UnsetRoute(resource models.Route) error {
	payload := map[string]any{"route": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", routeURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) ClearRoute() error {
	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=clear", routeURL), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllRoute() ([]models.Route, error) {
	req, err := s.client.NewRequest(http.MethodGet, routeURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.Route `json:"route"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) CountRoute() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", routeURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.Route `json:"route"`
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

// route6
func (s *NetworkService) AddRoute6(resource models.Route6) error {
	payload := map[string]any{"route6": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, route6URL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteRoute6(network string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", route6URL, network), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UpdateRoute6(resource models.Route6) error {
	payload := map[string]any{"route6": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, route6URL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UnsetRoute6(resource models.Route6) error {
	payload := map[string]any{"route6": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", route6URL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) ClearRoute6() error {
	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=clear", route6URL), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllRoute6() ([]models.Route6, error) {
	req, err := s.client.NewRequest(http.MethodGet, route6URL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.Route6 `json:"route6"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) CountRoute6() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", route6URL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.Route6 `json:"route6"`
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

// rsskeytype
func (s *NetworkService) UpdateRSSKeyType(resource models.RSSKeyType) error {
	payload := map[string]any{"rsskeytype": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, rssKeyTypeURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllRSSKeyType() ([]models.RSSKeyType, error) {
	req, err := s.client.NewRequest(http.MethodGet, rssKeyTypeURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.RSSKeyType `json:"rsskeytype"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// vlan
func (s *NetworkService) AddVLAN(resource models.VLAN) error {
	payload := map[string]any{"vlan": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vlanURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteVLAN(id string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vlanURL, id), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UpdateVLAN(resource models.VLAN) error {
	payload := map[string]any{"vlan": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, vlanURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UnsetVLAN(resource models.VLAN) error {
	payload := map[string]any{"vlan": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", vlanURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllVLAN() ([]models.VLAN, error) {
	req, err := s.client.NewRequest(http.MethodGet, vlanURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VLAN `json:"vlan"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetVLAN(id string) (*models.VLAN, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vlanURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VLAN `json:"vlan"`
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

func (s *NetworkService) CountVLAN() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vlanURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VLAN `json:"vlan"`
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

// vlan_binding
func (s *NetworkService) GetAllVLANBinding() ([]models.VLANBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vlanBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VLANBinding `json:"vlan_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetVLANBinding(id string) (*models.VLANBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vlanBindingURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VLANBinding `json:"vlan_binding"`
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

// vlan_channel_binding
func (s *NetworkService) AddVLANChannelBinding(resource models.VLANChannelBinding) error {
	payload := map[string]any{"vlan_channel_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vlanChannelBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteVLANChannelBinding(id string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vlanChannelBindingURL, id), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllVLANChannelBinding() ([]models.VLANChannelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vlanChannelBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VLANChannelBinding `json:"vlan_channel_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetVLANChannelBinding(id string) (*models.VLANChannelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vlanChannelBindingURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VLANChannelBinding `json:"vlan_channel_binding"`
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

func (s *NetworkService) CountVLANChannelBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vlanChannelBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VLANChannelBinding `json:"vlan_channel_binding"`
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

// vlan_interface_binding
func (s *NetworkService) AddVLANInterfaceBinding(resource models.VLANInterfaceBinding) error {
	payload := map[string]any{"vlan_interface_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vlanInterfaceBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteVLANInterfaceBinding(id string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vlanInterfaceBindingURL, id), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllVLANInterfaceBinding() ([]models.VLANInterfaceBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vlanInterfaceBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VLANInterfaceBinding `json:"vlan_interface_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetVLANInterfaceBinding(id string) (*models.VLANInterfaceBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vlanInterfaceBindingURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VLANInterfaceBinding `json:"vlan_interface_binding"`
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

func (s *NetworkService) CountVLANInterfaceBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vlanInterfaceBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VLANInterfaceBinding `json:"vlan_interface_binding"`
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

// vlan_linkset_binding
func (s *NetworkService) AddVLANLinkSetBinding(resource models.VLANLinkSetBinding) error {
	payload := map[string]any{"vlan_linkset_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vlanLinkSetBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteVLANLinkSetBinding(id string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vlanLinkSetBindingURL, id), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllVLANLinkSetBinding() ([]models.VLANLinkSetBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vlanLinkSetBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VLANLinkSetBinding `json:"vlan_linkset_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetVLANLinkSetBinding(id string) (*models.VLANLinkSetBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vlanLinkSetBindingURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VLANLinkSetBinding `json:"vlan_linkset_binding"`
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

func (s *NetworkService) CountVLANLinkSetBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vlanLinkSetBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VLANLinkSetBinding `json:"vlan_linkset_binding"`
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

// vlan_nsip6_binding
func (s *NetworkService) AddVLANNSIP6Binding(resource models.VLANNSIP6Binding) error {
	payload := map[string]any{"vlan_nsip6_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vlanNSIP6BindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteVLANNSIP6Binding(id string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vlanNSIP6BindingURL, id), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllVLANNSIP6Binding() ([]models.VLANNSIP6Binding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vlanNSIP6BindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VLANNSIP6Binding `json:"vlan_nsip6_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetVLANNSIP6Binding(id string) (*models.VLANNSIP6Binding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vlanNSIP6BindingURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VLANNSIP6Binding `json:"vlan_nsip6_binding"`
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

func (s *NetworkService) CountVLANNSIP6Binding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vlanNSIP6BindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VLANNSIP6Binding `json:"vlan_nsip6_binding"`
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

// vlan_nsip_binding
func (s *NetworkService) AddVLANNSIPBinding(resource models.VLANNSIPBinding) error {
	payload := map[string]any{"vlan_nsip_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vlanNSIPBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteVLANNSIPBinding(id string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vlanNSIPBindingURL, id), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllVLANNSIPBinding() ([]models.VLANNSIPBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vlanNSIPBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VLANNSIPBinding `json:"vlan_nsip_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetVLANNSIPBinding(id string) (*models.VLANNSIPBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vlanNSIPBindingURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VLANNSIPBinding `json:"vlan_nsip_binding"`
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

func (s *NetworkService) CountVLANNSIPBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vlanNSIPBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VLANNSIPBinding `json:"vlan_nsip_binding"`
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

// vrid
func (s *NetworkService) AddVRID(resource models.VRID) error {
	payload := map[string]any{"vrid": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vridURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteVRID(id string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vridURL, id), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UpdateVRID(resource models.VRID) error {
	payload := map[string]any{"vrid": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, vridURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UnsetVRID(resource models.VRID) error {
	payload := map[string]any{"vrid": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", vridURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllVRID() ([]models.VRID, error) {
	req, err := s.client.NewRequest(http.MethodGet, vridURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VRID `json:"vrid"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetVRID(id string) (*models.VRID, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vridURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VRID `json:"vrid"`
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

func (s *NetworkService) CountVRID() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vridURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VRID `json:"vrid"`
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

// vrid6
func (s *NetworkService) AddVRID6(resource models.VRID6) error {
	payload := map[string]any{"vrid6": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vrid6URL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteVRID6(id string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vrid6URL, id), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UpdateVRID6(resource models.VRID6) error {
	payload := map[string]any{"vrid6": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, vrid6URL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UnsetVRID6(resource models.VRID6) error {
	payload := map[string]any{"vrid6": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", vrid6URL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllVRID6() ([]models.VRID6, error) {
	req, err := s.client.NewRequest(http.MethodGet, vrid6URL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VRID6 `json:"vrid6"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetVRID6(id string) (*models.VRID6, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vrid6URL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VRID6 `json:"vrid6"`
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

func (s *NetworkService) CountVRID6() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vrid6URL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VRID6 `json:"vrid6"`
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

// vrid6_binding
func (s *NetworkService) GetAllVRID6Biding() ([]models.VRID6Binding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vrid6BindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VRID6Binding `json:"vrid6_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetVRID6Biding(id string) (*models.VRID6Binding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vrid6BindingURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VRID6Binding `json:"vrid6_binding"`
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

// vrid6_channel_binding
func (s *NetworkService) AddVRID6ChannelBinding(resource models.VRID6ChannelBinding) error {
	payload := map[string]any{"vrid6_channel_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vrid6ChannelBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteVRID6ChannelBinding(id string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vrid6ChannelBindingURL, id), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllVRID6ChannelBinding() ([]models.VRID6ChannelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vrid6ChannelBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VRID6ChannelBinding `json:"vrid6_channel_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetVRID6ChannelBinding(id string) (*models.VRID6ChannelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vrid6ChannelBindingURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VRID6ChannelBinding `json:"vrid6_channel_binding"`
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

func (s *NetworkService) CountVRID6ChannelBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vrid6ChannelBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VRID6ChannelBinding `json:"vrid6_channel_binding"`
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

// vrid6_interface_binding
func (s *NetworkService) AddVRID6InterfaceBinding(resource models.VRID6InterfaceBinding) error {
	payload := map[string]any{"vrid6_interface_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vrid6InterfaceBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteVRID6InterfaceBinding(id string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vrid6InterfaceBindingURL, id), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllVRID6InterfaceBinding() ([]models.VRID6InterfaceBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vrid6InterfaceBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VRID6InterfaceBinding `json:"vrid6_interface_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetVRID6InterfaceBinding(id string) (*models.VRID6InterfaceBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vrid6InterfaceBindingURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VRID6InterfaceBinding `json:"vrid6_interface_binding"`
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

func (s *NetworkService) CountVRID6InterfaceBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vrid6InterfaceBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VRID6InterfaceBinding `json:"vrid6_interface_binding"`
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

// vrid6_nsip6_binding
func (s *NetworkService) GetAllVRID6NSIP6Binding() ([]models.VRID6NSIP6Binding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vrid6NSIP6BindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VRID6NSIP6Binding `json:"vrid6_nsip6_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetVRID6NSIP6Binding(id string) (*models.VRID6NSIP6Binding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vrid6NSIP6BindingURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VRID6NSIP6Binding `json:"vrid6_nsip6_binding"`
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

func (s *NetworkService) CountVRID6NSIP6Binding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vrid6NSIP6BindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VRID6NSIP6Binding `json:"vrid6_nsip6_binding"`
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

// vrid6_nsip_binding
func (s *NetworkService) GetAllVRID6NSIPBinding() ([]models.VRID6NSIPBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vrid6NSIPBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VRID6NSIPBinding `json:"vrid6_nsip_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetVRID6NSIPBinding(id string) (*models.VRID6NSIPBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vrid6NSIPBindingURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VRID6NSIPBinding `json:"vrid6_nsip_binding"`
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

func (s *NetworkService) CountVRID6NSIPBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vrid6NSIPBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VRID6NSIPBinding `json:"vrid6_nsip_binding"`
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

// vrid6_trackinterface_binding
func (s *NetworkService) AddVRID6TrackInterfaceBinding(resource models.VRID6TrackInterfaceBinding) error {
	payload := map[string]any{"vrid6_trackinterface_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vrid6TrackInterfaceBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteVRID6TrackInterfaceBinding(id string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vrid6TrackInterfaceBindingURL, id), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllVRID6TrackInterfaceBinding() ([]models.VRID6TrackInterfaceBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vrid6TrackInterfaceBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VRID6TrackInterfaceBinding `json:"vrid6_trackinterface_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetVRID6TrackInterfaceBinding(id string) (*models.VRID6TrackInterfaceBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vrid6TrackInterfaceBindingURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VRID6TrackInterfaceBinding `json:"vrid6_trackinterface_binding"`
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

func (s *NetworkService) CountVRID6TrackInterfaceBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vrid6TrackInterfaceBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VRID6TrackInterfaceBinding `json:"vrid6_trackinterface_binding"`
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

// vridparam
func (s *NetworkService) UpdateVRIDParam(resource models.VRIDParam) error {
	payload := map[string]any{"vridparam": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, vridParamURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UnsetVRIDParam(resource models.VRIDParam) error {
	payload := map[string]any{"vridparam": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", vridParamURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllVRIDParam() ([]models.VRIDParam, error) {
	req, err := s.client.NewRequest(http.MethodGet, vridParamURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VRIDParam `json:"vridparam"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// vrid_binding
func (s *NetworkService) GetAllVRIDBinding() ([]models.VRIDBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vridBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VRIDBinding `json:"vrid_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetVRIDBinding(id string) (*models.VRIDBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vridBindingURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VRIDBinding `json:"vrid_binding"`
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

// vrid_channel_binding
func (s *NetworkService) AddVRIDChannelBinding(resource models.VRIDChannelBinding) error {
	payload := map[string]any{"vrid_channel_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vridChannelBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteVRIDChannelBinding(id string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vridChannelBindingURL, id), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllVRIDChannelBinding() ([]models.VRIDChannelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vridChannelBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VRIDChannelBinding `json:"vrid_channel_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetVRIDChannelBinding(id string) (*models.VRIDChannelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vridChannelBindingURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VRIDChannelBinding `json:"vrid_channel_binding"`
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

func (s *NetworkService) CountVRIDChannelBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vridChannelBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VRIDChannelBinding `json:"vrid_channel_binding"`
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

// vrid_interface_binding
func (s *NetworkService) AddVRIDInterfaceBinding(resource models.VRIDInterfaceBinding) error {
	payload := map[string]any{"vrid_interface_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vridInterfaceBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteVRIDInterfaceBinding(id string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vridInterfaceBindingURL, id), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllVRIDInterfaceBinding() ([]models.VRIDInterfaceBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vridInterfaceBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VRIDInterfaceBinding `json:"vrid_interface_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetVRIDInterfaceBinding(id string) (*models.VRIDInterfaceBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vridInterfaceBindingURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VRIDInterfaceBinding `json:"vrid_interface_binding"`
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

func (s *NetworkService) CountVRIDInterfaceBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vridInterfaceBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VRIDInterfaceBinding `json:"vrid_interface_binding"`
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

// vrid_nsip6_binding
func (s *NetworkService) GetAllVRIDNSIP6Binding() ([]models.VRIDNSIP6Binding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vridNSIP6BindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VRIDNSIP6Binding `json:"vrid_nsip6_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetVRIDNSIP6Binding(id string) (*models.VRIDNSIP6Binding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vridNSIP6BindingURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VRIDNSIP6Binding `json:"vrid_nsip6_binding"`
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

func (s *NetworkService) CountVRIDNSIP6Binding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vridNSIP6BindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VRIDNSIP6Binding `json:"vrid_nsip6_binding"`
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

// vrid_nsip_binding
func (s *NetworkService) GetAllVRIDNSIPBinding() ([]models.VRIDNSIPBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vridNSIPBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VRIDNSIPBinding `json:"vrid_nsip_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetVRIDNSIPBinding(id string) (*models.VRIDNSIPBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vridNSIPBindingURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VRIDNSIPBinding `json:"vrid_nsip_binding"`
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

func (s *NetworkService) CountVRIDNSIPBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vridNSIPBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VRIDNSIPBinding `json:"vrid_nsip_binding"`
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

// vrid_trackinterface_binding
func (s *NetworkService) AddVRIDTrackInterfaceBinding(resource models.VRIDTrackInterfaceBinding) error {
	payload := map[string]any{"vrid_trackinterface_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vridTrackInterfaceBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteVRIDTrackInterfaceBinding(id string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vridTrackInterfaceBindingURL, id), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllVRIDTrackInterfaceBinding() ([]models.VRIDTrackInterfaceBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vridTrackInterfaceBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VRIDTrackInterfaceBinding `json:"vrid_trackinterface_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetVRIDTrackInterfaceBinding(id string) (*models.VRIDTrackInterfaceBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vridTrackInterfaceBindingURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VRIDTrackInterfaceBinding `json:"vrid_trackinterface_binding"`
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

func (s *NetworkService) CountVRIDTrackInterfaceBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vridTrackInterfaceBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VRIDTrackInterfaceBinding `json:"vrid_trackinterface_binding"`
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

// vxlan
func (s *NetworkService) AddVXLAN(resource models.VXLAN) error {
	payload := map[string]any{"vxlan": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vxlanURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteVXLAN(id string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vxlanURL, id), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UpdateVXLAN(resource models.VXLAN) error {
	payload := map[string]any{"vxlan": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, vxlanURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) UnsetVXLAN(resource models.VXLAN) error {
	payload := map[string]any{"vxlan": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", vxlanURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllVXLAN() ([]models.VXLAN, error) {
	req, err := s.client.NewRequest(http.MethodGet, vxlanURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VXLAN `json:"vxlan"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetVXLAN(id string) (*models.VXLAN, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vxlanURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VXLAN `json:"vxlan"`
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

func (s *NetworkService) CountVXLAN() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vxlanURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VXLAN `json:"vxlan"`
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

// vxlanvlanmap
func (s *NetworkService) AddVXLANVLANMap(resource models.VXLANVLANMap) error {
	payload := map[string]any{"vxlanvlanmap": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vxlanVLANMapURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteVXLANVLANMap(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vxlanVLANMapURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllVXLANVLANMap() ([]models.VXLANVLANMap, error) {
	req, err := s.client.NewRequest(http.MethodGet, vxlanVLANMapURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VXLANVLANMap `json:"vxlanvlanmap"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetVXLANVLANMap(name string) (*models.VXLANVLANMap, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vxlanVLANMapURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VXLANVLANMap `json:"vxlanvlanmap"`
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

func (s *NetworkService) CountVXLANVLANMap() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vxlanVLANMapURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VXLANVLANMap `json:"vxlanvlanmap"`
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

// vxlanvlanmap_binding
func (s *NetworkService) GetAllVXLANVLANMapBinding() ([]models.VXLANVLANMapBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vxlanVLANMapBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VXLANVLANMapBinding `json:"vxlanvlanmap_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetVXLANVLANMapBinding(name string) (*models.VXLANVLANMapBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vxlanVLANMapBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VXLANVLANMapBinding `json:"vxlanvlanmap_binding"`
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

// vxlanvlanmap_vxlan_binding
func (s *NetworkService) AddVXLANVLANMapVXLANBinding(resource models.VXLANVLANMapVXLANBinding) error {
	payload := map[string]any{"vxlanvlanmap_vxlan_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vxlanVLANMapVXLANBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteVXLANVLANMapVXLANBinding(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vxlanVLANMapVXLANBindingURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllVXLANVLANMapVXLANBinding() ([]models.VXLANVLANMapVXLANBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vxlanVLANMapVXLANBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VXLANVLANMapVXLANBinding `json:"vxlanvlanmap_vxlan_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetVXLANVLANMapVXLANBinding(name string) (*models.VXLANVLANMapVXLANBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vxlanVLANMapVXLANBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VXLANVLANMapVXLANBinding `json:"vxlanvlanmap_vxlan_binding"`
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

func (s *NetworkService) CountVXLANVLANMapVXLANBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vxlanVLANMapVXLANBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VXLANVLANMapVXLANBinding `json:"vxlanvlanmap_vxlan_binding"`
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

// vxlan_binding
func (s *NetworkService) GetAllVXLANBinding() ([]models.VxlanBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vxlanBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VxlanBinding `json:"vxlan_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetVXLANBinding(id string) (*models.VxlanBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vxlanBindingURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VxlanBinding `json:"vxlan_binding"`
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

// vxlan_iptunnel_binding
func (s *NetworkService) AddVXLANIPTunnelBinding(resource models.VXLANIPTunnelBinding) error {
	payload := map[string]any{"vxlan_iptunnel_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vxlanIPTunnelBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteVXLANIPTunnelBinding(id string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vxlanIPTunnelBindingURL, id), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllVXLANIPTunnelBinding() ([]models.VXLANIPTunnelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vxlanIPTunnelBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VXLANIPTunnelBinding `json:"vxlan_iptunnel_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetVXLANIPTunnelBinding(id string) (*models.VXLANIPTunnelBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vxlanIPTunnelBindingURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VXLANIPTunnelBinding `json:"vxlan_iptunnel_binding"`
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

func (s *NetworkService) CountVXLANIPTunnelBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vxlanIPTunnelBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VXLANIPTunnelBinding `json:"vxlan_iptunnel_binding"`
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

// vxlan_nsip6_binding
func (s *NetworkService) AddVXLANNSIP6Binding(resource models.VXLANNSIP6Binding) error {
	payload := map[string]any{"vxlan_nsip6_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vxlanNSIP6BindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteVXLANNSIP6Binding(id string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vxlanNSIP6BindingURL, id), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllVXLANNSIP6Binding() ([]models.VXLANNSIP6Binding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vxlanNSIP6BindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VXLANNSIP6Binding `json:"vxlan_nsip6_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetVXLANNSIP6Binding(id string) (*models.VXLANNSIP6Binding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vxlanNSIP6BindingURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VXLANNSIP6Binding `json:"vxlan_nsip6_binding"`
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

func (s *NetworkService) CountVXLANNSIP6Binding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vxlanNSIP6BindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VXLANNSIP6Binding `json:"vxlan_nsip6_binding"`
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

// vxlan_nsip_binding
func (s *NetworkService) AddVXLANNSIPBinding(resource models.VXLANNSIPBinding) error {
	payload := map[string]any{"vxlan_nsip_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vxlanNSIPBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteVXLANNSIPBinding(id string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vxlanNSIPBindingURL, id), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllVXLANNSIPBinding() ([]models.VXLANNSIPBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vxlanNSIPBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VXLANNSIPBinding `json:"vxlan_nsip_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetVXLANNSIPBinding(id string) (*models.VXLANNSIPBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vxlanNSIPBindingURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VXLANNSIPBinding `json:"vxlan_nsip_binding"`
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

func (s *NetworkService) CountVXLANNSIPBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vxlanNSIPBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VXLANNSIPBinding `json:"vxlan_nsip_binding"`
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

// vxlan_srcip_binding
func (s *NetworkService) AddVXLANSrcIPBinding(resource models.VXLANSrcIPBinding) error {
	payload := map[string]any{"vxlan_srcip_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, vxlanSrcIPBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) DeleteVXLANSrcIPBinding(id string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vxlanSrcIPBindingURL, id), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *NetworkService) GetAllVXLANSrcIPBinding() ([]models.VXLANSrcIPBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, vxlanSrcIPBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VXLANSrcIPBinding `json:"vxlan_srcip_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

func (s *NetworkService) GetVXLANSrcIPBinding(id string) (*models.VXLANSrcIPBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", vxlanSrcIPBindingURL, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []models.VXLANSrcIPBinding `json:"vxlan_srcip_binding"`
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

func (s *NetworkService) CountVXLANSrcIPBinding() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", vxlanSrcIPBindingURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Data []models.VXLANSrcIPBinding `json:"vxlan_srcip_binding"`
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
