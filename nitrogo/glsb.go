package nitrogo

const (
	gslbConfigURL                                           = "/nitro/v1/config/gslbconfig"
	gslbDomainURL                                           = "/nitro/v1/config/gslbdomain"
	gslbDomainBindingURL                                    = "/nitro/v1/config/gslbdomain_binding"
	gslbDomainGSLBServiceGroupMemberBindingURL              = "/nitro/v1/config/gslbdomain_gslbservicegroupmember_binding"
	gslbDomainGSLBServiceGroupBindingURL                    = "/nitro/v1/config/gslbdomain_gslbservicegroup_binding"
	gslbDomainGSLBServiceBindingURL                         = "/nitro/v1/config/gslbdomain_gslbservice_binding"
	gslbDomainGSLBVServerBindingURL                         = "/nitro/v1/config/gslbdomain_gslbvserver_binding"
	gslbDomainLBMonitorBindingURL                           = "/nitro/v1/config/gslbdomain_lbmonitor_binding"
	gslbDNSEntriesURL                                       = "/nitro/v1/config/gslbdnsentries"
	gslbDNSEntryURL                                         = "/nitro/v1/config/gslbdnsentry"
	gslbParameterURL                                        = "/nitro/v1/config/gslbparameter"
	gslbRunningConfigURL                                    = "/nitro/v1/config/gslbrunningconfig"
	gslbServiceURL                                          = "/nitro/v1/config/gslbservice"
	gslbServiceGroupURL                                     = "/nitro/v1/config/gslbservicegroup"
	gslbServiceGroupBindingURL                              = "/nitro/v1/config/gslbservicegroup_binding"
	gslbServiceGroupGSLBServiceGroupMemberBindingURL        = "/nitro/v1/config/gslbservicegroup_gslbservicegroupmember_binding"
	gslbServiceGroupLBMonitorBindingURL                     = "/nitro/v1/config/gslbservicegroup_lbmonitor_binding"
	gslbServiceGroupServiceGroupEntityMonBindingsBindingURL = "/nitro/v1/config/gslbservicegroup_servicegroupentitymonbindings_binding"
	gslbServiceBindingURL                                   = "/nitro/v1/config/gslbservice_binding"
	gslbServiceDNSViewBindingURL                            = "/nitro/v1/config/gslbservice_dnsview_binding"
	gslbServiceLBMonitorBindingURL                          = "/nitro/v1/config/gslbservice_lbmonitor_binding"
	gslbSiteURL                                             = "/nitro/v1/config/gslbsite"
	gslbSiteBindingURL                                      = "/nitro/v1/config/gslbsite_binding"
	gslbSiteGSLBServiceGroupMemberBindingURL                = "/nitro/v1/config/gslbsite_gslbservicegroupmember_binding"
	gslbSiteGSLBServiceGroupBindingURL                      = "/nitro/v1/config/gslbsite_gslbservicegroup_binding"
	gslbSiteGSLBServiceBindingURL                           = "/nitro/v1/config/gslbsite_gslbservice_binding"
	gslbSyncStatusURL                                       = "/nitro/v1/config/gslbsyncstatus"
	gslbVServerURL                                          = "/nitro/v1/config/gslbvserver"
	gslbVServerBindingURL                                   = "/nitro/v1/config/gslbvserver_binding"
	gslbVServerDomainBindingURL                             = "/nitro/v1/config/gslbvserver_domain_binding"
	gslbVServerGSLBServiceGroupMemberBindingURL             = "/nitro/v1/config/gslbvserver_gslbservicegroupmember_binding"
	gslbVServerGSLBServiceGroupBindingURL                   = "/nitro/v1/config/gslbvserver_gslbservicegroup_binding"
	gslbVServerGSLBServiceBindingURL                        = "/nitro/v1/config/gslbvserver_gslbservice_binding"
	gslbVServerSpilloverPolicyBindingURL                    = "/nitro/v1/config/gslbvserver_spilloverpolicy_binding"
)

// Global Server Load Balancing (GSLB) configuration. GSLB feature ensures that client requests are
// directed to a best performing site available in a global enterprise and distributed Internet environment.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/gslb/gslb
type GSLBService struct {
	client *Client
}

// gslbconfig
func (s *GSLBService) SyncGSLBConfig() {}

// gslbdomain
func (s *GSLBService) GetAllGSLBDomain() {}
func (s *GSLBService) GetGSLBDomain()    {}
func (s *GSLBService) CountGSLBDomain()  {}

// gslbdomain_binding
func (s *GSLBService) GetAllGSLBDomainBinding() {}
func (s *GSLBService) GetGSLBDomainBinding()    {}

// dslbdomain_gslbservicegroupmember_binding
func (s *GSLBService) GetAllGSLBDomainGSLBServiceGroupMemberBinding() {}
func (s *GSLBService) GetGSLBDomainGSLBServiceGroupMemberBinding()    {}
func (s *GSLBService) CountGSLBDomainGSLBServiceGroupMemberBinding()  {}

// gslbdomain_gslbservicegroup_binding
func (s *GSLBService) GetAllGSLBDomainGSLBServiceGroupBinding() {}
func (s *GSLBService) GetGSLBDomainGSLBServiceGroupBinding()    {}
func (s *GSLBService) CountGSLBDomainGSLBServiceGroupBinding()  {}

// dslbdomain_gslbservice_binding
func (s *GSLBService) GetAllGSLBDomainGSLBServiceBinding() {}
func (s *GSLBService) GetGSLBDomainGSLBServiceBinding()    {}
func (s *GSLBService) CountGSLBDomainGSLBServiceBinding()  {}

// gslbdomain_gslbvserver_binding
func (s *GSLBService) GetAllGSLBDomainGSLBVServerBinding() {}
func (s *GSLBService) GetGSLBDomainGSLBVServerBinding()    {}
func (s *GSLBService) CountGSLBDomainGSLBVServerBinding()  {}

// gslbdomain_lbmonitor_binding
func (s *GSLBService) GetAllGSLBDomainLBMonitorBinding() {}
func (s *GSLBService) GetGSLBDomainLBMonitorBinding()    {}
func (s *GSLBService) CountGSLBDomainLBMonitorBinding()  {}

// gslbdnsentries
func (s *GSLBService) ClearGSLBDNSEntries()  {}
func (s *GSLBService) GetAllGSLBDNSEntries() {}
func (s *GSLBService) CountGSLBDNSEntries()  {}

// gslbdnsentry
func (s *GSLBService) DeleteGSLBDNSEntry() {}

// gslbparameter
func (s *GSLBService) UpdateGSLBParameter() {}
func (s *GSLBService) UnsetGSLBParameter()  {}
func (s *GSLBService) GetAllGSLBParameter() {}

// gslbrunningconfig
func (s *GSLBService) GetAllGSLBRunningConfig() {}

// gslbservice
func (s *GSLBService) AddGSLBService()    {}
func (s *GSLBService) DeleteGSLBService() {}
func (s *GSLBService) UpdateGSLBService() {}
func (s *GSLBService) UnsetGSLBService()  {}
func (s *GSLBService) GetAllGSLBService() {}
func (s *GSLBService) GetGSLBService()    {}
func (s *GSLBService) CountGSLBService()  {}
func (s *GSLBService) RenameGSLBService() {}

// gslbservicegroup
func (s *GSLBService) AddGSLBServiceGroup()     {}
func (s *GSLBService) DeleteGSLBServiceGroup()  {}
func (s *GSLBService) UpdateGSLBServiceGroup()  {}
func (s *GSLBService) UnsetGSLBServiceGroup()   {}
func (s *GSLBService) EnableGSLBServiceGroup()  {}
func (s *GSLBService) DisableGSLBServiceGroup() {}
func (s *GSLBService) GetAllGSLBServiceGroup()  {}
func (s *GSLBService) GetGSLBServiceGroup()     {}
func (s *GSLBService) CountGSLBServiceGroup()   {}
func (s *GSLBService) RenameGSLBServiceGroup()  {}

// gslbservicegroup_binding
func (s *GSLBService) GetAllGSLBServiceGroupBinding() {}
func (s *GSLBService) GetGSLBServiceGroupBinding()    {}

// gslbservicegroup_gslbservicegroupmember_binding
func (s *GSLBService) AddGSLBServiceGroupGSLBServiceGroupMemberBinding()    {}
func (s *GSLBService) DeleteGSLBServiceGroupGSLBServiceGroupMemberBinding() {}
func (s *GSLBService) GetAllGSLBServiceGroupGSLBServiceGroupMemberBinding() {}
func (s *GSLBService) GetGSLBServiceGroupGSLBServiceGroupMemberBinding()    {}
func (s *GSLBService) CountGSLBServiceGroupGSLBServiceGroupMemberBinding()  {}

// gslbservicegroup_lbmonitor_binding
func (s *GSLBService) AddGSLBServiceGroupLBMonitorBinding()    {}
func (s *GSLBService) DeleteGSLBServiceGroupLBMonitorBinding() {}
func (s *GSLBService) GetAllGSLBServiceGroupLBMonitorBinding() {}
func (s *GSLBService) GetGSLBServiceGroupLBMonitorBinding()    {}
func (s *GSLBService) CountGSLBServiceGroupLBMonitorBinding()  {}

// gslbservicegroup_servicegroupentitymonbindings_binding
func (s *GSLBService) GetAllGSLBServiceGroupServiceGroupEntityMonBindingsBinding() {}
func (s *GSLBService) GetGSLBServiceGroupServiceGroupEntityMonBindingsBinding()    {}
func (s *GSLBService) CountGSLBServiceGroupServiceGroupEntityMonBindingsBinding()  {}

// gslbservice_binding
func (s *GSLBService) GetAllGSLBServiceBinding() {}
func (s *GSLBService) GetGSLBServiceBinding()    {}

// gslbservice_dnsview_binding
func (s *GSLBService) AddGSLBServiceDNSViewBinding()    {}
func (s *GSLBService) DeleteGSLBServiceDNSViewBinding() {}
func (s *GSLBService) GetAllGSLBServiceDNSViewBinding() {}
func (s *GSLBService) GetGSLBServiceDNSViewBinding()    {}
func (s *GSLBService) CountGSLBServiceDNSViewBinding()  {}

// dslbservice_lbmonitor_binding
func (s *GSLBService) AddGSLBServiceLBMonitorBinding()    {}
func (s *GSLBService) DeleteGSLBServiceLBMonitorBinding() {}
func (s *GSLBService) GetAllGSLBServiceLBMonitorBinding() {}
func (s *GSLBService) GetGSLBServiceLBMonitorBinding()    {}
func (s *GSLBService) CountGSLBServiceLBMonitorBinding()  {}

// gslbsite
func (s *GSLBService) AddGSLBSite()    {}
func (s *GSLBService) DeleteGSLBSite() {}
func (s *GSLBService) UpdateGSLBSite() {}
func (s *GSLBService) UnsetGSLBSite()  {}
func (s *GSLBService) GetAllGSLBSite() {}
func (s *GSLBService) GetGSLBSite()    {}
func (s *GSLBService) CountGSLBSite()  {}
func (s *GSLBService) RenameGSLBSite() {}

// gslbsite_binding
func (s *GSLBService) GetAllGSLBSiteBinding() {}
func (s *GSLBService) GetGSLBSiteBinding()    {}

// gslbsite_gslbservicegroupmember_binding
func (s *GSLBService) GetAllGSLBSiteGSLBServiceGroupMemberBinding() {}
func (s *GSLBService) GetGSLBSiteGSLBServiceGroupMemberBinding()    {}
func (s *GSLBService) CountGSLBSiteGSLBServiceGroupMemberBinding()  {}

// gslbsite_gslbservicegroup_binding
func (s *GSLBService) GetAllGSLBSiteGSLBServiceGroupBinding() {}
func (s *GSLBService) GetGSLBSiteGSLBServiceGroupBinding()    {}
func (s *GSLBService) CountGSLBSiteGSLBServiceGroupBinding()  {}

// gslbsite_gslbservice_binding
func (s *GSLBService) GetAllGSLBSiteGSLBServiceBinding() {}
func (s *GSLBService) GetGSLBSiteGSLBServiceBinding()    {}
func (s *GSLBService) CountGSLBSiteGSLBServiceBinding()  {}

// gslbsyncstatus
func (s *GSLBService) GetAllGSLBSyncStatus() {}

// gslbvserver
func (s *GSLBService) AddGSLBVServer()     {}
func (s *GSLBService) DeleteGSLBVServer()  {}
func (s *GSLBService) UpdateGSLBVServer()  {}
func (s *GSLBService) UnsetGSLBVServer()   {}
func (s *GSLBService) EnableGSLBVServer()  {}
func (s *GSLBService) DisableGSLBVServer() {}
func (s *GSLBService) GetAllGSLBVServer()  {}
func (s *GSLBService) GetGSLBVServer()     {}
func (s *GSLBService) CountGSLBVServer()   {}
func (s *GSLBService) RenameGSLBVServer()  {}

// gslbvserver_binding
func (s *GSLBService) GetAllGSLBVServerBinding() {}
func (s *GSLBService) GetGSLBVServerBinding()    {}

// gslbvserver_domain_binding
func (s *GSLBService) AddGSLBVServerDomainBinding()    {}
func (s *GSLBService) DeleteGSLBVServerDomainBinding() {}
func (s *GSLBService) GetAllGSLBVServerDomainBinding() {}
func (s *GSLBService) GetGSLBVServerDomainBinding()    {}
func (s *GSLBService) CountGSLBVServerDomainBinding()  {}

// gslbvserver_gslbservicegroupmember_binding
func (s *GSLBService) GetAllGSLBVServerGSLBServiceGroupMemberBinding() {}
func (s *GSLBService) GetGSLBVServerGSLBServiceGroupMemberBinding()    {}
func (s *GSLBService) CountGSLBVServerGSLBServiceGroupMemberBinding()  {}

// gslbvserver_gslbservicegroup_binding
func (s *GSLBService) AddGSLBVServerGSLBServiceGroupBinding()    {}
func (s *GSLBService) DeleteGSLBVServerGSLBServiceGroupBinding() {}
func (s *GSLBService) GetAllGSLBVServerGSLBServiceGroupBinding() {}
func (s *GSLBService) GetGSLBVServerGSLBServiceGroupBinding()    {}
func (s *GSLBService) CountGSLBVServerGSLBServiceGroupBinding()  {}

// gslbvserver_gslbservice_binding
func (s *GSLBService) AddGSLBVServerGSLBServiceBinding()    {}
func (s *GSLBService) DeleteGSLBVServerGSLBServiceBinding() {}
func (s *GSLBService) GetAllGSLBVServerGSLBServiceBinding() {}
func (s *GSLBService) GetGSLBVServerGSLBServiceBinding()    {}
func (s *GSLBService) CountGSLBVServerGSLBServiceBinding()  {}

// gslbvserver_spilloverpolicy_binding
func (s *GSLBService) AddGSLBVServerSpilloverPolicyBinding()    {}
func (s *GSLBService) DeleteGSLBVServerSpilloverPolicyBinding() {}
func (s *GSLBService) GetAllGSLBVServerSpilloverPolicyBinding() {}
func (s *GSLBService) GetGSLBVServerSpilloverPolicyBinding()    {}
func (s *GSLBService) CountGSLBVServerSpilloverPolicyBinding()  {}
