package nitrogo

const (
	haFilesURL                               = "/nitro/v1/config/hafiles"
	haNodeURL                                = "/nitro/v1/config/hanode"
	haNodeBindingURL                         = "/nitro/v1/config/hanode_binding"
	haNodeCIBindingURL                       = "/nitro/v1/config/hanode_ci_binding"
	haNodeFISBindingURL                      = "/nitro/v1/config/hanode_fis_binding"
	haNodePartialFailureInterfacesBindingURL = "/nitro/v1/config/hanode_partialfailureinterfaces_binding"
	haNodeRouteMonitor6BindingURL            = "/nitro/v1/config/hanode_routemonitor6_binding"
	haNodeRouteMonitorBindingURL             = "/nitro/v1/config/hanode_routemonitor_binding"
	haSyncURL                                = "/nitro/v1/config/hasync"
	haSyncFailuresURL                        = "/nitro/v1/config/hasyncfailures"
)

// Configuration for failover resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/ha/hafailover
type HAService struct {
	client *Client
}

// hafiles
func (s *HAService) SyncHAFiles() {}

// hanode
func (s *HAService) AddHANode()    {}
func (s *HAService) DeleteHANode() {}
func (s *HAService) UpdateHANode() {}
func (s *HAService) UnsetHANode()  {}
func (s *HAService) GetAllHANode() {}
func (s *HAService) GetHANode()    {}
func (s *HAService) CountHANode()  {}

// hanode_binding
func (s *HAService) GetAllHANodeBinding() {}
func (s *HAService) GetHANodeBinding()    {}

// hanode_ci_binding
func (s *HAService) GetAllHANodeCIBinding() {}
func (s *HAService) GetHANodeCIBinding()    {}
func (s *HAService) CountHANodeCIBinding()  {}

// hanode_fis_binding
func (s *HAService) GetAllHANodeFISBinding() {}
func (s *HAService) GetHANodeFISBinding()    {}
func (s *HAService) CountHANodeFISBinding()  {}

// hanode_partialfailureinterfaces_binding
func (s *HAService) GetAllHANodePartialFailureInterfacesBinding() {}
func (s *HAService) GetHANodePartialFailureInterfacesBinding()    {}
func (s *HAService) CountHANodePartialFailureInterfacesBinding()  {}

// hanode_routemonitor6_binding
func (s *HAService) AddHANodeRouteMonitor6Binding()    {}
func (s *HAService) DeleteHANodeRouteMonitor6Binding() {}
func (s *HAService) GetAllHANodeRouteMonitor6Binding() {}
func (s *HAService) GetHANodeRouteMonitor6Binding()    {}
func (s *HAService) CountHANodeRouteMonitor6Binding()  {}

// hanode_routemonitor_binding
func (s *HAService) AddHANodeRouteMonitorBinding()    {}
func (s *HAService) DeleteHANodeRouteMonitorBinding() {}
func (s *HAService) GetAllHANodeRouteMonitorBinding() {}
func (s *HAService) GetHANodeRouteMonitorBinding()    {}
func (s *HAService) CountHANodeRouteMonitorBinding()  {}

// hasync
func (s *HAService) ForceHASync() {}

// hasyncfailures
func (s *HAService) GetAllHASyncFailures() {}
