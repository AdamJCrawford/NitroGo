package nitrogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/AdamJCrawford/NitroGo/nitrogo/models"
)

const (
	dbsMonitorsURL                                      = "/nitro/v1/config/dbsmonitors"
	extendedMemoryParamURL                              = "/nitro/v1/config/extendedmemoryparam"
	locationURL                                         = "/nitro/v1/config/location"
	locationDataURL                                     = "/nitro/v1/config/locationdata"
	locationFileURL                                     = "/nitro/v1/config/locationfile"
	locationFile6URL                                    = "/nitro/v1/config/locationfile6"
	locationParameterURL                                = "/nitro/v1/config/locationparameter"
	nsTraceURL                                          = "/nitro/v1/config/nstrace"
	radiusNodeURL                                       = "/nitro/v1/config/radiusnode"
	reportingURL                                        = "/nitro/v1/config/reporting"
	serverURL                                           = "/nitro/v1/config/server"
	serverBindingURL                                    = "/nitro/v1/config/server_binding"
	serverGSLBServiceBindingURL                         = "/nitro/v1/config/server_gslbservice_binding"
	serverGSLBServiceGroupBindingURL                    = "/nitro/v1/config/server_gslbservicegroup_binding"
	serverServiceBindingURL                             = "/nitro/v1/config/server_service_binding"
	serverServiceGroupBindingURL                        = "/nitro/v1/config/server_servicegroup_binding"
	serviceURL                                          = "/nitro/v1/config/service"
	serviceBindingURL                                   = "/nitro/v1/config/service_binding"
	serviceGroupLBMonitorBindingURL                     = "/nitro/v1/config/servicegroup_lbmonitor_binding"
	serviceGroupURL                                     = "/nitro/v1/config/servicegroup"
	serviceGroupBindingURL                              = "/nitro/v1/config/servicegroup_binding"
	serviceLBMonitorBindingURL                          = "/nitro/v1/config/service_lbmonitor_binding"
	serviceGroupServiceGroupEntityMonBindingsBindingURL = "/nitro/v1/config/servicegroup_servicegroupentitymonbindings_binding"
	serviceGroupServiceGroupMemberBindingURL            = "/nitro/v1/config/servicegroup_servicegroupmember_binding"
	serviceGroupServiceGroupMemberListBindingURL        = "/nitro/v1/config/servicegroup_servicegoupmemberlist_binding"
	serviceGroupBindingsURL                             = "/nitro/v1/config/servicegroupbindings"
	svcBindingsURL                                      = "/nitro/v1/config/svcbindings"
	vServerURL                                          = "/nitro/v1/config/vserver"
)

// Basic system configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/basic
type BasicService struct {
	client *Client
}

// dbsmonitors
// Configuration for DB monitors resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/dbsmonitors
func (s *BasicService) RestartDBSMonitors() error {
	req, err := s.client.NewRequest(http.MethodPost, dbsMonitorsURL+"?action=restart", nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// extendedmemoryparam
// Configuration for Parameter for extended memory used by LSN and Subscriber Store resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/extendedmemoryparam
func (s *BasicService) UpdateExtendedMemoryParam(resource models.ExtendedMemoryParam) error {
	payload := map[string]any{"extendedmemoryparam": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, extendedMemoryParamURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) UnsetExtendedMemoryParam(resource models.ExtendedMemoryParam) error {
	payload := map[string]any{"extendedmemoryparam": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", extendedMemoryParamURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) GetAllExtendedMemoryParam() (models.ExtendedMemoryParam, error) {
	req, err := s.client.NewRequest(http.MethodGet, extendedMemoryParamURL, nil)
	if err != nil {
		return models.ExtendedMemoryParam{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.ExtendedMemoryParam{}, err
	}

	var result struct {
		ExtendedMemoryParam models.ExtendedMemoryParam `json:"extendedmemoryparam"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.ExtendedMemoryParam{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.ExtendedMemoryParam, nil
}

// location
// Configuration for location resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/location
func (s *BasicService) AddLocation(resource models.Location) error {
	payload := map[string]any{"location": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, locationURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) DeleteLocation(ip string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", locationURL, ip), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) GetAllLocation() ([]models.Location, error) {
	req, err := s.client.NewRequest(http.MethodGet, locationURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Location []models.Location `json:"location"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Location, nil
}

func (s *BasicService) GetLocation(ip string) (models.Location, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", locationURL, ip), nil)
	if err != nil {
		return models.Location{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.Location{}, err
	}

	var result struct {
		Location []models.Location `json:"location"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.Location{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Location) == 0 {
		return models.Location{}, fmt.Errorf("location %s not found", ip)
	}

	return result.Location[0], nil
}

// Struct lacks Count field, leaving for later
func (s *BasicService) CountLocation() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, locationURL+"?count=yes", nil)
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
		} `json:"location"`
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

// locationdata
// Configuration for location data resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/locationdata
func (s *BasicService) ClearLocationData() error {
	payload := map[string]any{"locationdata": models.LocationData{}}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=clear", locationDataURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// locationfile
// Configuration for location file resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/locationfile
func (s *BasicService) AddLocationFile(resource models.LocationFile) error {
	payload := map[string]any{"locationfile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, locationFileURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) DeleteLocationFile(LocationFile string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", locationFileURL, LocationFile), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) GetAllLocationFile() ([]models.LocationFile, error) {
	req, err := s.client.NewRequest(http.MethodGet, locationFileURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		LocationFile []models.LocationFile `json:"locationfile"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.LocationFile, nil
}

func (s *BasicService) ImportLocationFile(resource models.LocationFile) error {
	payload := map[string]any{"locationfile": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=import", locationFileURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// locationfile6
// Configuration for location file6 resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/locationfile6
func (s *BasicService) AddLocationFile6(resource models.LocationFile6) error {
	payload := map[string]any{"locationfile6": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, locationFile6URL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) DeleteLocationFile6() error {
	req, err := s.client.NewRequest(http.MethodDelete, locationFile6URL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) GetAllLocationFile6() ([]models.LocationFile6, error) {
	req, err := s.client.NewRequest(http.MethodGet, locationFile6URL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		LocationFile6 []models.LocationFile6 `json:"locationfile6"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.LocationFile6, nil
}

// Struct lacks Count field, leaving for later
func (s *BasicService) CountLocationFile6() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, locationFile6URL+"?count=yes", nil)
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
		} `json:"locationfile6"`
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

func (s *BasicService) ImportLocationFile6(resource models.LocationFile6) error {
	payload := map[string]any{"locationfile6": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=import", locationFile6URL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// locationparameter
// Configuration for location parameter resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/locationparameter
func (s *BasicService) UpdateLocationParameter(resource models.LocationParameter) error {
	payload := map[string]any{"locationparameter": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, locationParameterURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) UnsetLocationParameter(resource models.LocationParameter) error {
	payload := map[string]any{"locationparameter": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", locationParameterURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) GetAllLocationParameter() (models.LocationParameter, error) {
	req, err := s.client.NewRequest(http.MethodGet, locationParameterURL, nil)
	if err != nil {
		return models.LocationParameter{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.LocationParameter{}, err
	}

	var result struct {
		LocationParameter models.LocationParameter `json:"locationparameter"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.LocationParameter{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.LocationParameter, nil
}

// nstrace
// Configuration for nstrace operations resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/nstrace
func (s *BasicService) GetAllNSTrace() ([]models.NSTrace, error) {
	req, err := s.client.NewRequest(http.MethodGet, nsTraceURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		NSTrace []models.NSTrace `json:"nstrace"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.NSTrace, nil
}

// radiusnode
// Configuration for RADIUS Node resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/radiusnode
func (s *BasicService) AddRADIUSNode(resource models.RADIUSNode) error {
	payload := map[string]any{"radiusnode": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, radiusNodeURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) DeleteRADIUSNode(nodeprefix string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", radiusNodeURL, nodeprefix), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) UpdateRADIUSNode(resource models.RADIUSNode) error {
	payload := map[string]any{"radiusnode": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, radiusNodeURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) GetAllRADIUSNode() ([]models.RADIUSNode, error) {
	req, err := s.client.NewRequest(http.MethodGet, radiusNodeURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		RADIUSNode []models.RADIUSNode `json:"radiusnode"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.RADIUSNode, nil
}

func (s *BasicService) GetRADIUSNode(nodeprefix string) (models.RADIUSNode, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", radiusNodeURL, nodeprefix), nil)
	if err != nil {
		return models.RADIUSNode{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.RADIUSNode{}, err
	}

	var result struct {
		RADIUSNode []models.RADIUSNode `json:"radiusnode"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.RADIUSNode{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.RADIUSNode) == 0 {
		return models.RADIUSNode{}, fmt.Errorf("radiusnode %s not found", nodeprefix)
	}

	return result.RADIUSNode[0], nil
}

// Struct lacks Count field, leaving for later
func (s *BasicService) CountRADIUSNode() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, radiusNodeURL+"?count=yes", nil)
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
		} `json:"radiusnode"`
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

// reporting
// Configuration for reporting resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/reporting
func (s *BasicService) EnableReporting() error {
	payload := map[string]any{"reporting": models.Reporting{}}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=enable", reportingURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) DisableReporting() error {
	payload := map[string]any{"reporting": models.Reporting{}}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=disable", reportingURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) GetAllReporting() (models.Reporting, error) {
	req, err := s.client.NewRequest(http.MethodGet, reportingURL, nil)
	if err != nil {
		return models.Reporting{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.Reporting{}, err
	}

	var result struct {
		Data models.Reporting `json:"reporting"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.Reporting{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Data, nil
}

// server
// Configuration for server resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/server
func (s *BasicService) AddServer(resource models.Server) error {
	payload := map[string]any{"server": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, serverURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) DeleteServer(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", serverURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) UpdateServer(resource models.Server) error {
	payload := map[string]any{"server": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, serverURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) UnsetServer(resource models.Server) error {
	payload := map[string]any{"server": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", serverURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) EnableServer(resource models.Server) error {
	payload := map[string]any{"server": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=enable", serverURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) DisableServer(resource models.Server) error {
	payload := map[string]any{"server": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=disable", serverURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) GetAllServer() ([]models.Server, error) {
	req, err := s.client.NewRequest(http.MethodGet, serverURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Server []models.Server `json:"server"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Server, nil
}

func (s *BasicService) GetServer(name string) (models.Server, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", serverURL, name), nil)
	if err != nil {
		return models.Server{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.Server{}, err
	}

	var result struct {
		Server []models.Server `json:"server"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.Server{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Server) == 0 {
		return models.Server{}, fmt.Errorf("server %s not found", name)
	}

	return result.Server[0], nil
}

func (s *BasicService) CountServer() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", serverURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		Server []models.Server `json:"server"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Server) > 0 {
		return int(result.Server[0].Count), nil
	}
	return 0, fmt.Errorf("failed to parse count")
}

func (s *BasicService) RenameServer(resource models.Server) error {
	payload := map[string]any{"server": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=rename", serverURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// server_binding
// Binding object which returns the resources bound to server.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/server_binding
func (s *BasicService) GetAllServerBinding() ([]models.ServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, serverBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		ServerBinding []models.ServerBinding `json:"server_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.ServerBinding, nil
}

func (s *BasicService) GetServerBinding(name string) (models.ServerBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", serverBindingURL, name), nil)
	if err != nil {
		return models.ServerBinding{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.ServerBinding{}, err
	}

	var result struct {
		ServerBinding []models.ServerBinding `json:"server_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.ServerBinding{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.ServerBinding) == 0 {
		return models.ServerBinding{}, fmt.Errorf("server_binding %s not found", name)
	}

	return result.ServerBinding[0], nil
}

// server_gslbservicegroup_binding
// Binding object showing the gslbservicegroup that can be bound to server.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/server_gslbservicegroup_binding
func (s *BasicService) GetAllServerGSLBServiceGroupBinding() ([]models.ServerGSLBServiceGroupBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, serverGSLBServiceGroupBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		ServerGSLBServiceGroupBinding []models.ServerGSLBServiceGroupBinding `json:"server_gslbservicegroup_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.ServerGSLBServiceGroupBinding, nil
}

func (s *BasicService) GetServerGSLBServiceGroupBinding(name string) ([]models.ServerGSLBServiceGroupBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", serverGSLBServiceGroupBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		ServerGSLBServiceGroupBinding []models.ServerGSLBServiceGroupBinding `json:"server_gslbservicegroup_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.ServerGSLBServiceGroupBinding, nil
}

// Struct lacks Count field, leaving for later
func (s *BasicService) CountServerGSLBServiceGroupBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, serverGSLBServiceGroupBindingURL+"?count=yes", nil)
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
		} `json:"server_gslbservicegroup_binding"`
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

// server_gslbservice_binding
// Binding object showing the gslbservice that can be bound to server.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/server_gslbservice_binding
func (s *BasicService) GetAllServerGSLBServiceBinding() ([]models.ServerGSLBServiceBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, serverGSLBServiceBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		ServerGSLBServiceBinding []models.ServerGSLBServiceBinding `json:"server_gslbservice_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.ServerGSLBServiceBinding, nil
}

func (s *BasicService) GetServerGSLBServiceBinding(name string) ([]models.ServerGSLBServiceBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", serverGSLBServiceBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		ServerGSLBServiceBinding []models.ServerGSLBServiceBinding `json:"server_gslbservice_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.ServerGSLBServiceBinding, nil
}

func (s *BasicService) CountServerGSLBServiceBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, serverGSLBServiceBindingURL+"?count=yes", nil)
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
		} `json:"server_gslbservice_binding"`
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

// server_servicegroup_binding
// Binding object showing the servicegroup that can be bound to server.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/server_servicegroup_binding
func (s *BasicService) GetAllServerServiceGroupBinding() ([]models.ServerServiceGroupBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, serverServiceGroupBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		ServerServiceGroupBinding []models.ServerServiceGroupBinding `json:"server_servicegroup_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.ServerServiceGroupBinding, nil
}

func (s *BasicService) GetServerServiceGroupBinding(name string) ([]models.ServerServiceGroupBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", serverServiceGroupBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		ServerServiceGroupBinding []models.ServerServiceGroupBinding `json:"server_servicegroup_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.ServerServiceGroupBinding, nil
}

func (s *BasicService) CountServerServiceGroupBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, serverServiceGroupBindingURL+"?count=yes", nil)
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
		} `json:"server_servicegroup_binding"`
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

// server_service_binding
// Binding object showing the service that can be bound to server.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/server_service_binding
func (s *BasicService) GetAllServerServiceBinding() ([]models.ServerServiceBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, serverServiceBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		ServerServiceBinding []models.ServerServiceBinding `json:"server_service_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.ServerServiceBinding, nil
}

func (s *BasicService) GetServerServiceBinding(name string) ([]models.ServerServiceBinding, error) {
	reqURL := fmt.Sprintf("%s/%s", serverServiceBindingURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		ServerServiceBinding []models.ServerServiceBinding `json:"server_service_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.ServerServiceBinding, nil
}

// Struct lacks Count field, leaving for later
func (s *BasicService) CountServerServiceBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, serverServiceBindingURL+"?count=yes", nil)
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
		} `json:"server_service_binding"`
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

// service
// Configuration for service resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/service
func (s *BasicService) AddService(resource models.Service) error {
	payload := map[string]interface{}{
		"service": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, serviceURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) DeleteService(name string) error {
	reqURL := fmt.Sprintf("%s/%s", serviceURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodDelete, reqURL, nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) UpdateService(resource models.Service) error {
	payload := map[string]interface{}{
		"service": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPut, serviceURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) UnsetService(resource models.Service) error {
	payload := map[string]interface{}{
		"service": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, serviceURL+"?action=unset", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) EnableService(resource models.Service) error {
	payload := map[string]interface{}{
		"service": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, serviceURL+"?action=enable", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) DisableService(resource models.Service) error {
	payload := map[string]interface{}{
		"service": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, serviceURL+"?action=disable", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) GetAllService() ([]models.Service, error) {
	req, err := s.client.NewRequest(http.MethodGet, serviceURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Services []models.Service `json:"service"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.Services, nil
}

func (s *BasicService) GetService(name string) (*models.Service, error) {
	reqURL := fmt.Sprintf("%s/%s", serviceURL, url.PathEscape(name))
	req, err := s.client.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		Services []models.Service `json:"service"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.Services) == 0 {
		return nil, fmt.Errorf("service not found")
	}

	return &result.Services[0], nil
}

func (s *BasicService) CountService() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, serviceURL+"?count=yes", nil)
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
		} `json:"service"`
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

func (s *BasicService) RenameService(resource models.Service) error {
	payload := map[string]interface{}{
		"service": resource,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := s.client.NewRequest(http.MethodPost, serviceURL+"?action=rename", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// servicegroup
// Configuration for service group resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/servicegroup
func (s *BasicService) AddServiceGroup(resource models.ServiceGroup) error {
	payload := map[string]any{"servicegroup": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, serviceGroupURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) DeleteServiceGroup(servicegroupname string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", serviceGroupURL, servicegroupname), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) UpdateServiceGroup(resource models.ServiceGroup) error {
	payload := map[string]any{"servicegroup": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, serviceGroupURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) UnsetServiceGroup(resource models.ServiceGroup) error {
	payload := map[string]any{"servicegroup": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=unset", serviceGroupURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) EnableServiceGroup(resource models.ServiceGroup) error {
	payload := map[string]any{"servicegroup": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=enable", serviceGroupURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) DisableServiceGroup(resource models.ServiceGroup) error {
	payload := map[string]any{"servicegroup": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=disable", serviceGroupURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) GetAllServiceGroup() ([]models.ServiceGroup, error) {
	req, err := s.client.NewRequest(http.MethodGet, serviceGroupURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		ServiceGroup []models.ServiceGroup `json:"servicegroup"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.ServiceGroup, nil
}

func (s *BasicService) GetServiceGroup(servicegroupname string) (models.ServiceGroup, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", serviceGroupURL, servicegroupname), nil)
	if err != nil {
		return models.ServiceGroup{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.ServiceGroup{}, err
	}

	var result struct {
		ServiceGroup []models.ServiceGroup `json:"servicegroup"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.ServiceGroup{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.ServiceGroup) == 0 {
		return models.ServiceGroup{}, fmt.Errorf("servicegroup %s not found", servicegroupname)
	}

	return result.ServiceGroup[0], nil
}

func (s *BasicService) CountServiceGroup() (int, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s?count=yes", serviceGroupURL), nil)
	if err != nil {
		return 0, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}

	var result struct {
		ServiceGroup []models.ServiceGroup `json:"servicegroup"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.ServiceGroup) > 0 {
		return int(result.ServiceGroup[0].Count), nil
	}
	return 0, fmt.Errorf("failed to parse count")
}

func (s *BasicService) RenameServiceGroup(resource models.ServiceGroup) error {
	payload := map[string]any{"servicegroup": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=rename", serviceGroupURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// servicegroupbindings
// Configuration for servicegroupbind resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/servicegroupbindings
func (s *BasicService) GetAllServiceGroupBindings() ([]models.ServiceGroupBindings, error) {
	req, err := s.client.NewRequest(http.MethodGet, serviceGroupBindingsURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		ServiceGroupBindings []models.ServiceGroupBindings `json:"servicegroupbindings"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.ServiceGroupBindings, nil
}

func (s *BasicService) GetServiceGroupBindings(servicegroupname string) (models.ServiceGroupBindings, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", serviceGroupBindingsURL, servicegroupname), nil)
	if err != nil {
		return models.ServiceGroupBindings{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.ServiceGroupBindings{}, err
	}

	var result struct {
		ServiceGroupBindings []models.ServiceGroupBindings `json:"servicegroupbindings"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.ServiceGroupBindings{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.ServiceGroupBindings) == 0 {
		return models.ServiceGroupBindings{}, fmt.Errorf("servicegroupbindings %s not found", servicegroupname)
	}

	return result.ServiceGroupBindings[0], nil
}

// Struct lacks Count field, leaving for later
func (s *BasicService) CountServiceGroupBindings() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, serviceGroupBindingsURL+"?count=yes", nil)
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
		} `json:"servicegroup_binding"`
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

// servicegroup_binding
// Binding object which returns the resources bound to servicegroup.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/servicegroup_binding
func (s *BasicService) GetAllServiceGroupBinding() ([]models.ServiceGroupBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, serviceGroupBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		ServiceGroupBinding []models.ServiceGroupBinding `json:"servicegroup_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.ServiceGroupBinding, nil
}

func (s *BasicService) GetServiceGroupBinding(serviceGroupName string) ([]models.ServiceGroupBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", serviceGroupBindingURL, serviceGroupName), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		ServiceGroupBinding []models.ServiceGroupBinding `json:"servicegroup_binding"`
	}

	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.ServiceGroupBinding, nil
}

// servicegroup_lbmonitor_binding
// Binding object showing the lbmonitor that can be bound to servicegroup.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/servicegroup_lbmonitor_binding
func (s *BasicService) AddServiceGroupLBMonitorBinding(resource models.ServiceGroupLBMonitorBinding) error {
	payload := map[string]any{"servicegroup_lbmonitor_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, serviceGroupLBMonitorBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) DeleteServiceGroupLBMonitorBinding(servicegroupname string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", serviceGroupLBMonitorBindingURL, servicegroupname, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) GetAllServiceGroupLBMonitorBinding() ([]models.ServiceGroupLBMonitorBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, serviceGroupLBMonitorBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		ServiceGroupLBMonitorBinding []models.ServiceGroupLBMonitorBinding `json:"servicegroup_lbmonitor_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.ServiceGroupLBMonitorBinding, nil
}

func (s *BasicService) GetServiceGroupLBMonitorBinding(servicegroupname string) ([]models.ServiceGroupLBMonitorBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", serviceGroupLBMonitorBindingURL, servicegroupname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		ServiceGroupLBMonitorBinding []models.ServiceGroupLBMonitorBinding `json:"servicegroup_lbmonitor_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.ServiceGroupLBMonitorBinding, nil
}

// Struct lacks Count field, leaving for later
func (s *BasicService) CountServiceGroupLBMonitorBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, serviceGroupLBMonitorBindingURL+"?count=yes", nil)
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
		} `json:"servicegroup_lbmonitor_binding"`
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

// servicegroup_servicegroupentitymonbindings_binding
// Binding object showing the servicegroupentitymonbindings that can be bound to servicegroup.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/servicegroup_servicegroupentitymonbindings_binding
func (s *BasicService) GetAllServiceGroupServiceGroupEntityMonBindingsBinding() ([]models.ServiceGroupServiceGroupEntityMonBindingsBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, serviceGroupServiceGroupEntityMonBindingsBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		ServiceGroupServiceGroupEntityMonBindingsBinding []models.ServiceGroupServiceGroupEntityMonBindingsBinding `json:"servicegroup_servicegroupentitymonbindings_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.ServiceGroupServiceGroupEntityMonBindingsBinding, nil
}

func (s *BasicService) GetServiceGroupServiceGroupEntityMonBindingsBinding(servicegroupname string) ([]models.ServiceGroupServiceGroupEntityMonBindingsBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", serviceGroupServiceGroupEntityMonBindingsBindingURL, servicegroupname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		ServiceGroupServiceGroupEntityMonBindingsBinding []models.ServiceGroupServiceGroupEntityMonBindingsBinding `json:"servicegroup_servicegroupentitymonbindings_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.ServiceGroupServiceGroupEntityMonBindingsBinding, nil
}

// Struct lacks Count field, leaving for later
func (s *BasicService) CountServiceGroupServiceGroupEntityMonBindingsBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, serviceGroupServiceGroupEntityMonBindingsBindingURL+"?count=yes", nil)
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
		} `json:"servicegroup_servicegroupentitymonbindings_binding"`
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

// servicegroup_servicegroupmemberlist_binding
// Binding object showing the servicegroupmemberlist that can be bound to servicegroup.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/servicegroup_servicegroupmemberlist_binding
func (s *BasicService) AddServiceGroupServiceGroupMemberListBinding(resource models.ServiceGroupServiceGroupMemberListBinding) error {
	payload := map[string]any{"servicegroup_servicegroupmemberlist_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, serviceGroupServiceGroupMemberListBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) DeleteServiceGroupServiceGroupMemberListBinding(servicegroupname string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", serviceGroupServiceGroupMemberListBindingURL, servicegroupname, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

// servicegroup_servicegroupmember_binding
// Binding object showing the servicegroupmember that can be bound to servicegroup.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/servicegroup_servicegroupmember_binding
func (s *BasicService) AddServiceGroupServiceGroupMemberBinding(resource models.ServiceGroupServiceGroupMemberBinding) error {
	payload := map[string]any{"servicegroup_servicegroupmember_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, serviceGroupServiceGroupMemberBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) DeleteServiceGroupServiceGroupMemberBinding(servicegroupname string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", serviceGroupServiceGroupMemberBindingURL, servicegroupname, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) GetAllServiceGroupServiceGroupMemberBinding() ([]models.ServiceGroupServiceGroupMemberBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, serviceGroupServiceGroupMemberBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		ServiceGroupServiceGroupMemberBinding []models.ServiceGroupServiceGroupMemberBinding `json:"servicegroup_servicegroupmember_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.ServiceGroupServiceGroupMemberBinding, nil
}

func (s *BasicService) GetServiceGroupServiceGroupMemberBinding(servicegroupname string) ([]models.ServiceGroupServiceGroupMemberBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", serviceGroupServiceGroupMemberBindingURL, servicegroupname), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		ServiceGroupServiceGroupMemberBinding []models.ServiceGroupServiceGroupMemberBinding `json:"servicegroup_servicegroupmember_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.ServiceGroupServiceGroupMemberBinding, nil
}

// Struct lacks Count field, leaving for later
func (s *BasicService) CountServiceGroupServiceGroupMemberBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, serviceGroupServiceGroupMemberBindingURL+"?count=yes", nil)
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
		} `json:"servicegroup_servicegroupmember_binding"`
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

// service_binding
// Binding object which returns the resources bound to service.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/service_binding
func (s *BasicService) GetAllServiceBinding() ([]models.ServiceBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, serviceBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		ServiceBinding []models.ServiceBinding `json:"service_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.ServiceBinding, nil
}

func (s *BasicService) GetServiceBinding(name string) (models.ServiceBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", serviceBindingURL, name), nil)
	if err != nil {
		return models.ServiceBinding{}, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return models.ServiceBinding{}, err
	}

	var result struct {
		ServiceBinding []models.ServiceBinding `json:"service_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return models.ServiceBinding{}, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(result.ServiceBinding) == 0 {
		return models.ServiceBinding{}, fmt.Errorf("service_binding %s not found", name)
	}

	return result.ServiceBinding[0], nil
}

// service_lbmonitor_binding
// Binding object showing the lbmonitor that can be bound to service.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/service_lbmonitor_binding
func (s *BasicService) AddServiceLBMonitorBinding(resource models.ServiceLBMonitorBinding) error {
	payload := map[string]any{"service_lbmonitor_binding": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, serviceLBMonitorBindingURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) DeleteServiceLBMonitorBinding(name string, args map[string]string) error {
	var argsStr string
	if len(args) > 0 {
		var parts []string
		for k, v := range args {
			parts = append(parts, fmt.Sprintf("%s:%s", k, v))
		}
		argsStr = "?args=" + strings.Join(parts, ",")
	}

	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s%s", serviceLBMonitorBindingURL, name, argsStr), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) GetAllServiceLBMonitorBinding() ([]models.ServiceLBMonitorBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, serviceLBMonitorBindingURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		ServiceLBMonitorBinding []models.ServiceLBMonitorBinding `json:"service_lbmonitor_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.ServiceLBMonitorBinding, nil
}

func (s *BasicService) GetServiceLBMonitorBinding(name string) ([]models.ServiceLBMonitorBinding, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", serviceLBMonitorBindingURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		ServiceLBMonitorBinding []models.ServiceLBMonitorBinding `json:"service_lbmonitor_binding"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.ServiceLBMonitorBinding, nil
}

// Struct lacks Count field, leaving for later
func (s *BasicService) CountServiceLBMonitorBinding() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, serviceLBMonitorBindingURL+"?count=yes", nil)
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
		} `json:"service_lbmonitor_binding"`
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

// svcbindings
// Configuration for service bindings resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/svcbindings
func (s *BasicService) GetAllSVCBindings() ([]models.SVCBindings, error) {
	req, err := s.client.NewRequest(http.MethodGet, svcBindingsURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		SVCBindings []models.SVCBindings `json:"svcbindings"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.SVCBindings, nil
}

func (s *BasicService) GetSVCBindings(name string) ([]models.SVCBindings, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", svcBindingsURL, name), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	var result struct {
		SVCBindings []models.SVCBindings `json:"svcbindings"`
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result.SVCBindings, nil
}

// Struct lacks Count field, leaving for later
func (s *BasicService) CountSVCBindings() (float64, error) {
	req, err := s.client.NewRequest(http.MethodGet, svcBindingsURL+"?count=yes", nil)
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
		} `json:"svcbindings"`
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

// vserver
// Configuration for virtual server resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/vserver
func (s *BasicService) DeleteVServer(name string) error {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", vServerURL, name), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) UpdateVServer(resource models.VServer) error {
	payload := map[string]any{"vserver": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPut, vServerURL, bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) EnableVServer(resource models.VServer) error {
	payload := map[string]any{"vserver": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=enable", vServerURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}

func (s *BasicService) DisableVServer(resource models.VServer) error {
	payload := map[string]any{"vserver": resource}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest(http.MethodPost, fmt.Sprintf("%s?action=disable", vServerURL), bytes.NewReader(data))
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	return err
}
