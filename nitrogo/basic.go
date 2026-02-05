package nitrogo

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	serviceGroupServiceGroupMemberListBinding           = "/nitro/v1/config/servicegroup_servicegoupmemberlist_binding"
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
func (s *BasicService) RestartDBSMonitors() {}

// extendedmemoryparam
// Configuration for Parameter for extended memory used by LSN and Subscriber Store resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/extendedmemoryparam
func (s *BasicService) UpdateExtendedMemoryParam() {}
func (s *BasicService) UnsetExtendedMemoryParam()  {}
func (s *BasicService) GetAllExtendedMemoryParam() {}

// location
// Configuration for location resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/location
func (s *BasicService) AddLocation()    {}
func (s *BasicService) DeleteLocation() {}
func (s *BasicService) GetAllLocation() {}
func (s *BasicService) GetLocation()    {}
func (s *BasicService) CountLocation()  {}

// locationdata
// Configuration for location data resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/locationdata
func (s *BasicService) ClearLocationData() {}

// locationfile
// Configuration for location file resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/locationfile
func (s *BasicService) AddLocationFile()    {}
func (s *BasicService) DeleteLocationFile() {}
func (s *BasicService) GetAllLocationFile() {}
func (s *BasicService) ImportLocationFile() {}

// locationfile6
// Configuration for location file6 resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/locationfile6
func (s *BasicService) AddLocationFile6()    {}
func (s *BasicService) DeleteLocationFile6() {}
func (s *BasicService) GetAllLocationFile6() {}
func (s *BasicService) CountLocationFile6()  {}
func (s *BasicService) ImportLocationFile6() {}

// locationparameter
// Configuration for location parameter resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/locationparameter
func (s *BasicService) UpdateLocationParameter() {}
func (s *BasicService) UnsetLocationParameter()  {}
func (s *BasicService) GetAllLocationParameter() {}

// nstrace
// Configuration for nstrace operations resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/nstrace
func (s *BasicService) GetAllNSTrace() {}

// radiusnode
// Configuration for RADIUS Node resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/radiusnode
func (s *BasicService) AddRADIUSNode()    {}
func (s *BasicService) DeleteRADIUSNode() {}
func (s *BasicService) UpdateRADIUSNode() {}
func (s *BasicService) GetAllRADIUSNode() {}
func (s *BasicService) GetRADIUSNode()    {}
func (s *BasicService) CountRADIUSNode()  {}

// reporting
// Configuration for reporting resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/reporting
func (s *BasicService) EnableReporting()  {}
func (s *BasicService) DisableReporting() {}
func (s *BasicService) GetAllReporting()  {}

// server
// Configuration for server resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/server
func (s *BasicService) AddServer()     {}
func (s *BasicService) DeleteServer()  {}
func (s *BasicService) UpdateServer()  {}
func (s *BasicService) UnsetServer()   {}
func (s *BasicService) EnableServer()  {}
func (s *BasicService) DisableServer() {}
func (s *BasicService) GetAllServer()  {}
func (s *BasicService) GetServer()     {}
func (s *BasicService) CountServer()   {}
func (s *BasicService) RenameServer()  {}

// server_binding
// Binding object which returns the resources bound to server.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/server_binding
func (s *BasicService) GetAllServerBinding() {}
func (s *BasicService) GetServerBinding()    {}

// server_gslbservicegroup_binding
// Binding object showing the gslbservicegroup that can be bound to server.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/server_gslbservicegroup_binding
func (s *BasicService) GetAllServerGSLBServiceGroupBinding() {}
func (s *BasicService) GetServerGSLBServiceGroupBinding()    {}
func (s *BasicService) CountServerGSLBServiceGroupBinding()  {}

// server_gslbservice_binding
// Binding object showing the gslbservice that can be bound to server.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/server_gslbservice_binding
func (s *BasicService) GetAllServerGSLBServiceBinding() {}
func (s *BasicService) GetServerGSLBServiceBinding()    {}
func (s *BasicService) CountServerGSLBServiceBinding()  {}

// server_servicegroup_binding
// Binding object showing the servicegroup that can be bound to server.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/server_servicegroup_binding
func (s *BasicService) GetAllServerServiceGroupBinding() {}
func (s *BasicService) GetServerServiceGroupBinding()    {}
func (s *BasicService) CountServerServiceGroupBinding()  {}

// server_service_binding
// Binding object showing the service that can be bound to server.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/server_service_binding
func (s *BasicService) GetAllServerServiceBinding() {}
func (s *BasicService) GetServerServiceBinding()    {}
func (s *BasicService) CountServerServiceBinding()  {}

// service
// Configuration for service resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/service
func (s *BasicService) AddService()     {}
func (s *BasicService) DeleteService()  {}
func (s *BasicService) UpdateService()  {}
func (s *BasicService) UnsetService()   {}
func (s *BasicService) EnableService()  {}
func (s *BasicService) DisableService() {}
func (s *BasicService) GetAllService()  {}
func (s *BasicService) GetService()     {}
func (s *BasicService) CountService()   {}
func (s *BasicService) RenameService()  {}

// servicegroup
// Configuration for service group resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/servicegroup
func (s *BasicService) AddServiceGroup()     {}
func (s *BasicService) DeleteServiceGroup()  {}
func (s *BasicService) UpdateServiceGroup()  {}
func (s *BasicService) UnsetServiceGroup()   {}
func (s *BasicService) EnableServiceGroup()  {}
func (s *BasicService) DisableServiceGroup() {}
func (s *BasicService) GetAllServiceGroup()  {}
func (s *BasicService) GetServiceGroup()     {}
func (s *BasicService) CountServiceGroup()   {}
func (s *BasicService) RenameServiceGroup()  {}

// servicegroupbindings
// Configuration for servicegroupbind resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/servicegroupbindings
func (s *BasicService) GetAllServiceGroupBindings() {}
func (s *BasicService) GetServiceGroupBindings()    {}
func (s *BasicService) CountServiceGroupBindings()  {}

// servicegroup_binding
// Binding object which returns the resources bound to servicegroup.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/servicegroup_binding
func (s *BasicService) GetAllServiceGroupBinding() {}
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
func (s *BasicService) AddServiceGroupLBMonitorBinding()    {}
func (s *BasicService) DeleteServiceGroupLBMonitorBinding() {}
func (s *BasicService) GetAllServiceGroupLBMonitorBinding() {}
func (s *BasicService) GetServiceGroupLBMonitorBinding()    {}
func (s *BasicService) CountServiceGroupLBMonitorBinding()  {}

// servicegroup_servicegroupentitymonbindings_binding
// Binding object showing the servicegroupentitymonbindings that can be bound to servicegroup.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/servicegroup_servicegroupentitymonbindings_binding
func (s *BasicService) GetAllServiceGroupServiceGroupEntityMonBindingsBinding() {}
func (s *BasicService) GetServiceGroupServiceGroupEntityMonBindingsBinding()    {}
func (s *BasicService) CountServiceGroupServiceGroupEntityMonBindingsBinding()  {}

// servicegroup_servicegoupmemberlist_binding
// Binding object showing the servicegroupmemberlist that can be bound to servicegroup.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/servicegroup_servicegroupmemberlist_binding
func (s *BasicService) AddServiceGroupServiceGroupMemberListBinding()    {}
func (s *BasicService) DeleteServiceGroupServiceGroupMemberListBinding() {}

// servicegroup_servicegroupmember_binding
// Binding object showing the servicegroupmember that can be bound to servicegroup.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/servicegroup_servicegroupmember_binding
func (s *BasicService) AddServiceGroupServiceGroupMemberBinding()    {}
func (s *BasicService) DeleteServiceGroupServiceGroupMemberBinding() {}
func (s *BasicService) GetAllServiceGroupServiceGroupMemberBinding() {}
func (s *BasicService) GetServiceGroupServiceGroupMemberBinding()    {}
func (s *BasicService) CountServiceGroupServiceGroupMemberBinding()  {}

// service_binding
// Binding object which returns the resources bound to service.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/service_binding
func (s *BasicService) GetAllServiceBinding() {}
func (s *BasicService) GetServiceBinding()    {}

// service_lbmonitor_binding
// Binding object showing the lbmonitor that can be bound to service.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/service_lbmonitor_binding
func (s *BasicService) AddServiceLBMonitorBinding()    {}
func (s *BasicService) DeleteServiceLBMonitorBinding() {}
func (s *BasicService) GetAllServiceLBMonitorBinding() {}
func (s *BasicService) GetServiceLBMonitorBinding()    {}
func (s *BasicService) CountServiceLBMonitorBinding()  {}

// svcbindings
// Configuration for service bindings resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/svcbindings
func (s *BasicService) GetAllSVCBindings() {}
func (s *BasicService) GetSVCBindings()    {}
func (s *BasicService) CountSVCBindings()  {}

// vserver
// Configuration for virtual server resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/basic/vserver
func (s *BasicService) DeleteVServer()  {}
func (s *BasicService) UpdateVServer()  {}
func (s *BasicService) EnableVServer()  {}
func (s *BasicService) DisableVServer() {}
