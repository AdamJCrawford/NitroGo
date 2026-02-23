package nitrogo

const (
	clusterFilesURL                                 = "/nitro/v1/config/clusterfiles"
	clusterInstanceURL                              = "/nitro/v1/config/clusterinstance"
	clusterInstanceBindingURL                       = "/nitro/v1/config/clusterinstance_binding"
	clusterInstanceClusterNodeBindingURL            = "/nitro/v1/config/clusterinstance_clusternode_binding"
	clusterNodeURL                                  = "/nitro/v1/config/clusternode"
	clusterNodeBindingURL                           = "/nitro/v1/config/clusternode_binding"
	clusterNodeRouteMonitorBindingURL               = "/nitro/v1/config/clusternode_routemonitor_binding"
	clusterNodeGroupURL                             = "/nitro/v1/config/clusternodegroup"
	clusterNodeGroupAuthenticationVServerBindingURL = "/nitro/v1/config/clusternodegroup_authenticationvserver_binding"
	clusterNodeGroupBindingURL                      = "/nitro/v1/config/clusternodegroup_binding"
	clusterNodeGroupClusterNodeBindingURL           = "/nitro/v1/config/clusternodegroup_clusternode_binding"
	clusterNodeGroupCRVServerBindingURL             = "/nitro/v1/config/clusternodegroup_crvserver_binding"
	clusterNodeGroupCSVServerBindingURL             = "/nitro/v1/config/clusternodegroup_csvserver_binding"
	clusterNodeGroupGSLBSiteBindingURL              = "/nitro/v1/config/clusternodegroup_gslbsite_binding"
	clusterNodeGroupGSLBVServerBindingURL           = "/nitro/v1/config/clusternodegroup_gslbvserver_binding"
	clusterNodeGroupLBVServerBindingURL             = "/nitro/v1/config/clusternodegroup_lbvserver_binding"
	clusterNodeGroupNSLimitIdentifierBindingURL     = "/nitro/v1/config/clusternodegroup_nslimitidentifier_binding"
	clusterNodeGroupServiceBindingURL               = "/nitro/v1/config/clusternodegroup_service_binding"
	clusterNodeGroupStreamIdentifierBindingURL      = "/nitro/v1/config/clusternodegroup_streamidentifier_binding"
	clusterNodeGroupVPNVServerBindingURL            = "/nitro/v1/config/clusternodegroup_vpnvserver_binding"
	clusterPropStatusURL                            = "/nitro/v1/config/clusterpropstatus"
	clusterSyncURL                                  = "/nitro/v1/config/clustersync"
)

// Configuration for 0 resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/cluster/cluster
type ClusterService struct {
	client *Client
}

// clusterfiles
func (s *ClusterService) SyncClusterFiles() {}

// clusterinstance
func (s *ClusterService) AddClusterInstance()     {}
func (s *ClusterService) DeleteClusterInstance()  {}
func (s *ClusterService) UpdateClusterInstance()  {}
func (s *ClusterService) UnsetClusterInstance()   {}
func (s *ClusterService) EnableClusterInstance()  {}
func (s *ClusterService) DisableClusterInstance() {}
func (s *ClusterService) GetAllClusterInstance()  {}
func (s *ClusterService) GetClusterInstance()     {}
func (s *ClusterService) CountClusterInstance()   {}

// clusterinstance_binding
func (s *ClusterService) GetAllClusterInstanceBinding() {}
func (s *ClusterService) GetClusterInstanceBinding()    {}

// cluisterinstance_clusternode_binding
func (s *ClusterService) GetAllClusterInstanceClusterNodeBinding() {}
func (s *ClusterService) GetClusterInstanceClusterNodeBinding()    {}
func (s *ClusterService) CountClusterInstanceClusterNodeBinding()  {}

// clusternode
func (s *ClusterService) AddClusterNode()    {}
func (s *ClusterService) DeleteClusterNode() {}
func (s *ClusterService) UpdateClusterNode() {}
func (s *ClusterService) UnsetClusterNode()  {}
func (s *ClusterService) GetAllClusterNode() {}
func (s *ClusterService) GetClusterNode()    {}
func (s *ClusterService) CountClusterNode()  {}

// clusternodegroup
func (s *ClusterService) AddClusterNodeGroup()    {}
func (s *ClusterService) DeleteClusterNodeGroup() {}
func (s *ClusterService) UpdateClusterNodeGroup() {}
func (s *ClusterService) UnsetClusterNodeGroup()  {}
func (s *ClusterService) GetAllClusterNodeGroup() {}
func (s *ClusterService) GetClusterNodeGroup()    {}
func (s *ClusterService) CountClusterNodeGroup()  {}

// clusternodegroup_authenticationvserver_binding
func (s *ClusterService) AddClusternodeGroupAuthenticationVServerBinding()    {}
func (s *ClusterService) DeleteClusternodeGroupAuthenticationVServerBinding() {}
func (s *ClusterService) GetAllClusternodeGroupAuthenticationVServerBinding() {}
func (s *ClusterService) GetClusternodeGroupAuthenticationVServerBinding()    {}
func (s *ClusterService) CountClusternodeGroupAuthenticationVServerBinding()  {}

// clusternodegroup_binding
func (s *ClusterService) GetAllClusterNodeGroupBinding() {}
func (s *ClusterService) GetClusterNodeGroupBinding()    {}

// clusternodegroup_clusternode_binding
func (s *ClusterService) AddClusterNodeGroupClusterNodeBinding()    {}
func (s *ClusterService) DeleteClusterNodeGroupClusterNodeBinding() {}
func (s *ClusterService) GetAllClusterNodeGroupClusterNodeBinding() {}
func (s *ClusterService) GetClusterNodeGroupClusterNodeBinding()    {}
func (s *ClusterService) CountClusterNodeGroupClusterNodeBinding()  {}

// clusternodegroup_crvserver_binding
func (s *ClusterService) AddClusterNodeGroupCRVServerBinding()    {}
func (s *ClusterService) DeleteClusterNodeGroupCRVServerBinding() {}
func (s *ClusterService) GetAllClusterNodeGroupCRVServerBinding() {}
func (s *ClusterService) GetClusterNodeGroupCRVServerBinding()    {}
func (s *ClusterService) CountClusterNodeGroupCRVServerBinding()  {}

// clusternodegroup_csvserver_binding
func (s *ClusterService) AddClusterNodeGroupCSVServerBinding()    {}
func (s *ClusterService) DeleteClusterNodeGroupCSVServerBinding() {}
func (s *ClusterService) GetAllClusterNodeGroupCSVServerBinding() {}
func (s *ClusterService) GetClusterNodeGroupCSVServerBinding()    {}
func (s *ClusterService) CountClusterNodeGroupCSVServerBinding()  {}

// clusternodegroup_gslbsite_binding
func (s *ClusterService) AddClusterNodeGroupGSLBSiteBinding()    {}
func (s *ClusterService) DeleteClusterNodeGroupGSLBSiteBinding() {}
func (s *ClusterService) GetAllClusterNodeGroupGSLBSiteBinding() {}
func (s *ClusterService) GetClusterNodeGroupGSLBSiteBinding()    {}
func (s *ClusterService) CountClusterNodeGroupGSLBSiteBinding()  {}

// clusternodegroup_gslbvserver_binding
func (s *ClusterService) AddClusterNodeGroupGSLBVServerBinding()    {}
func (s *ClusterService) DeleteClusterNodeGroupGSLBVServerBinding() {}
func (s *ClusterService) GetAllClusterNodeGroupGSLBVServerBinding() {}
func (s *ClusterService) GetClusterNodeGroupGSLBVServerBinding()    {}
func (s *ClusterService) CountClusterNodeGroupGSLBVServerBinding()  {}

// clusternodegroup_lbvserver_binding
func (s *ClusterService) AddClusterNodeGroupLBVServerBinding()    {}
func (s *ClusterService) DeleteClusterNodeGroupLBVServerBinding() {}
func (s *ClusterService) GetAllClusterNodeGroupLBVServerBinding() {}
func (s *ClusterService) GetClusterNodeGroupLBVServerBinding()    {}
func (s *ClusterService) CountClusterNodeGroupLBVServerBinding()  {}

// clusternodegroup_nslimitidentifier_binding
func (s *ClusterService) AddClusterNodeGroupNSLimitIdentifierBinding()    {}
func (s *ClusterService) DeleteClusterNodeGroupNSLimitIdentifierBinding() {}
func (s *ClusterService) GetAllClusterNodeGroupNSLimitIdentifierBinding() {}
func (s *ClusterService) GetClusterNodeGroupNSLimitIdentifierBinding()    {}
func (s *ClusterService) CountClusterNodeGroupNSLimitIdentifierBinding()  {}

// clusternodegroup_service_binding
func (s *ClusterService) AddClusterNodeGroupServiceBinding()    {}
func (s *ClusterService) DeleteClusterNodeGroupServiceBinding() {}
func (s *ClusterService) GetAllClusterNodeGroupServiceBinding() {}
func (s *ClusterService) GetClusterNodeGroupServiceBinding()    {}
func (s *ClusterService) CountClusterNodeGroupServiceBinding()  {}

// clusternodegroup_streamidentifier_binding
func (s *ClusterService) AddClusterNodeGroupStreamIdentifierBinding()    {}
func (s *ClusterService) DeleteClusterNodeGroupStreamIdentifierBinding() {}
func (s *ClusterService) GetAllClusterNodeGroupStreamIdentifierBinding() {}
func (s *ClusterService) GetClusterNodeGroupStreamIdentifierBinding()    {}
func (s *ClusterService) CountClusterNodeGroupStreamIdentifierBinding()  {}

// clusternodegroup_vpnvserver_binding
func (s *ClusterService) AddClusterNodeGroupVPNVServerBinding()    {}
func (s *ClusterService) DeleteClusterNodeGroupVPNVServerBinding() {}
func (s *ClusterService) GetAllClusterNodeGroupVPNVServerBinding() {}
func (s *ClusterService) GetClusterNodeGroupVPNVServerBinding()    {}
func (s *ClusterService) CountClusterNodeGroupVPNVServerBinding()  {}

// clusternode_binding
func (s *ClusterService) GetAllClusterNodeBinding() {}
func (s *ClusterService) GetClusterNodeBinding()    {}

// clusternode_routemonitor_binding
func (s *ClusterService) AddClusterNodeRouteMonitorBinding()    {}
func (s *ClusterService) DeleteClusterNodeRouteMonitorBinding() {}
func (s *ClusterService) GetAllClusterNodeRouteMonitorBinding() {}
func (s *ClusterService) GetClusterNodeRouteMonitorBinding()    {}
func (s *ClusterService) CountClusterNodeRouteMonitorBinding()  {}

// clusterpropstatus
func (s *ClusterService) GetAllClusterPropStatus() {}
func (s *ClusterService) CountClusterPropStatus()  {}
func (s *ClusterService) ClearClusterPropStatus()  {}

// clustersync
func (s *ClusterService) ForceClusterSync() {}
