package nitrogo

// Global Server Load Balancing (GSLB) configuration. GSLB feature ensures that client requests are
// directed to a best performing site available in a global enterprise and distributed Internet environment.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/gslb/gslb
type GLSBService struct {
	client *Client
}

// gslbconfig
func (s *GLSBService) SyncGSLBConfig() {}

// gslbdomain
func (s *GLSBService) GetAllGSLBDomain() {}
func (s *GLSBService) GetGSLBDomain()    {}
func (s *GLSBService) CountGSLBDomain()  {}

// gslbdomain_binding
func (s *GLSBService) GetAllGSLBDomainBinding() {}
func (s *GLSBService) GetGSLBDomainBinding()    {}

// dslbdomain_gslbservicegroupmember_binding
func (s *GLSBService) GetAllGSLBDomainGSLBServiceGroupMemberBinding() {}
func (s *GLSBService) GetGSLBDomainGSLBServiceGroupMemberBinding()    {}
func (s *GLSBService) CountGSLBDomainGSLBServiceGroupMemberBinding()  {}

// gslbdomain_gslbservicegroup_binding
func (s *GLSBService) GetAllGSLBDomainGSLBServiceGroupBinding() {}
func (s *GLSBService) GetGSLBDomainGSLBServiceGroupBinding()    {}
func (s *GLSBService) CountGSLBDomainGSLBServiceGroupBinding()  {}

// dslbdomain_gslbservice_binding
func (s *GLSBService) GetAllGSLBDomainGSLBServiceBinding() {}
func (s *GLSBService) GetGSLBDomainGSLBServiceBinding()    {}
func (s *GLSBService) CountGSLBDomainGSLBServiceBinding()  {}

// gslbdomain_gslbvserver_binding
func (s *GLSBService) GetAllGSLBDomainGSLBVServerBinding() {}
func (s *GLSBService) GetGSLBDomainGSLBVServerBinding()    {}
func (s *GLSBService) CountGSLBDomainGSLBVServerBinding()  {}

// gslbdomain_lbmonitor_binding
func (s *GLSBService) GetAllGSLBDomainLBMonitorBinding() {}
func (s *GLSBService) GetGSLBDomainLBMonitorBinding()    {}
func (s *GLSBService) CountGSLBDomainLBMonitorBinding()  {}

// gslbdnsentries
func (s *GLSBService) ClearGSLBDNSEntries()  {}
func (s *GLSBService) GetAllGSLBDNSEntries() {}
func (s *GLSBService) CountGSLBDNSEntries()  {}

// gslbdnsentry
func (s *GLSBService) DeleteGSLBDNSEntry() {}

// gslbparameter
func (s *GLSBService) UpdateGSLBParameter() {}
func (s *GLSBService) UnsetGSLBParameter()  {}
func (s *GLSBService) GetAllGSLBParameter() {}

// gslbrunningconfig
func (s *GLSBService) GetAllGSLBRunningConfig() {}

// gslbservice
func (s *GLSBService) AddGSLBService()    {}
func (s *GLSBService) DeleteGSLBService() {}
func (s *GLSBService) UpdateGSLBService() {}
func (s *GLSBService) UnsetGSLBService()  {}
func (s *GLSBService) GetAllGSLBService() {}
func (s *GLSBService) GetGSLBService()    {}
func (s *GLSBService) CountGSLBService()  {}
func (s *GLSBService) RenameGSLBService() {}

// gslbservicegroup
func (s *GLSBService) AddGSLBServiceGroup()     {}
func (s *GLSBService) DeleteGSLBServiceGroup()  {}
func (s *GLSBService) UpdateGSLBServiceGroup()  {}
func (s *GLSBService) UnsetGSLBServiceGroup()   {}
func (s *GLSBService) EnableGSLBServiceGroup()  {}
func (s *GLSBService) DisableGSLBServiceGroup() {}
func (s *GLSBService) GetAllGSLBServiceGroup()  {}
func (s *GLSBService) GetGSLBServiceGroup()     {}
func (s *GLSBService) CountGSLBServiceGroup()   {}
func (s *GLSBService) RenameGSLBServiceGroup()  {}

// gslbservicegroup_binding
func (s *GLSBService) GetAllGSLBServiceGroupBinding() {}
func (s *GLSBService) GetGSLBServiceGroupBinding()    {}

// gslbservicegroup_gslbservicegroupmember_binding
func (s *GLSBService) AddGSLBServiceGroupGSLBServiceGroupMemberBinding()    {}
func (s *GLSBService) DeleteGSLBServiceGroupGSLBServiceGroupMemberBinding() {}
func (s *GLSBService) GetAllGSLBServiceGroupGSLBServiceGroupMemberBinding() {}
func (s *GLSBService) GetGSLBServiceGroupGSLBServiceGroupMemberBinding()    {}
func (s *GLSBService) CountGSLBServiceGroupGSLBServiceGroupMemberBinding()  {}

// gslbservicegroup_lbmonitor_binding
func (s *GLSBService) AddGSLBServiceGroupLBMonitorBinding()    {}
func (s *GLSBService) DeleteGSLBServiceGroupLBMonitorBinding() {}
func (s *GLSBService) GetAllGSLBServiceGroupLBMonitorBinding() {}
func (s *GLSBService) GetGSLBServiceGroupLBMonitorBinding()    {}
func (s *GLSBService) CountGSLBServiceGroupLBMonitorBinding()  {}

// gslbservicegroup_servicegroupentitymonbindings_binding
func (s *GLSBService) GetAllGSLBServiceGroupServiceGroupEntityMonBindingsBinding() {}
func (s *GLSBService) GetGSLBServiceGroupServiceGroupEntityMonBindingsBinding()    {}
func (s *GLSBService) CountGSLBServiceGroupServiceGroupEntityMonBindingsBinding()  {}

// gslbservice_binding
func (s *GLSBService) GetAllGSLBServiceBinding() {}
func (s *GLSBService) GetGSLBServiceBinding()    {}

// gslbservice_dnsview_binding
func (s *GLSBService) AddGSLBServiceDNSViewBinding()    {}
func (s *GLSBService) DeleteGSLBServiceDNSViewBinding() {}
func (s *GLSBService) GetAllGSLBServiceDNSViewBinding() {}
func (s *GLSBService) GetGSLBServiceDNSViewBinding()    {}
func (s *GLSBService) CountGSLBServiceDNSViewBinding()  {}

// dslbservice_lbmonitor_binding
func (s *GLSBService) AddGSLBServiceLBMonitorBinding()    {}
func (s *GLSBService) DeleteGSLBServiceLBMonitorBinding() {}
func (s *GLSBService) GetAllGSLBServiceLBMonitorBinding() {}
func (s *GLSBService) GetGSLBServiceLBMonitorBinding()    {}
func (s *GLSBService) CountGSLBServiceLBMonitorBinding()  {}

// gslbsite
func (s *GLSBService) AddGSLBSite()    {}
func (s *GLSBService) DeleteGSLBSite() {}
func (s *GLSBService) UpdateGSLBSite() {}
func (s *GLSBService) UnsetGSLBSite()  {}
func (s *GLSBService) GetAllGSLBSite() {}
func (s *GLSBService) GetGSLBSite()    {}
func (s *GLSBService) CountGSLBSite()  {}
func (s *GLSBService) RenameGSLBSite() {}

// gslbsite_binding
func (s *GLSBService) GetAllGSLBSiteBinding() {}
func (s *GLSBService) GetGSLBSiteBinding()    {}

// gslbsite_gslbservicegroupmember_binding
func (s *GLSBService) GetAllGSLBSiteGSLBServiceGroupMemberBinding() {}
func (s *GLSBService) GetGSLBSiteGSLBServiceGroupMemberBinding()    {}
func (s *GLSBService) CountGSLBSiteGSLBServiceGroupMemberBinding()  {}

// gslbsite_gslbservicegroup_binding
func (s *GLSBService) GetAllGSLBSiteGSLBServiceGroupBinding() {}
func (s *GLSBService) GetGSLBSiteGSLBServiceGroupBinding()    {}
func (s *GLSBService) CountGSLBSiteGSLBServiceGroupBinding()  {}

// gslbsite_gslbservice_binding
func (s *GLSBService) GetAllGSLBSiteGSLBServiceBinding() {}
func (s *GLSBService) GetGSLBSiteGSLBServiceBinding()    {}
func (s *GLSBService) CountGSLBSiteGSLBServiceBinding()  {}

// gslbsyncstatus
func (s *GLSBService) GetAllGSLBSyncStatus() {}

// gslbvserver
func (s *GLSBService) AddGSLBVServer()     {}
func (s *GLSBService) DeleteGSLBVServer()  {}
func (s *GLSBService) UpdateGSLBVServer()  {}
func (s *GLSBService) UnsetGSLBVServer()   {}
func (s *GLSBService) EnableGSLBVServer()  {}
func (s *GLSBService) DisableGSLBVServer() {}
func (s *GLSBService) GetAllGSLBVServer()  {}
func (s *GLSBService) GetGSLBVServer()     {}
func (s *GLSBService) CountGSLBVServer()   {}
func (s *GLSBService) RenameGSLBVServer()  {}

// gslbvserver_binding
func (s *GLSBService) GetAllGSLBVServerBinding() {}
func (s *GLSBService) GetGSLBVServerBinding()    {}

// gslbvserver_domain_binding
func (s *GLSBService) AddGSLBVServerDomainBinding()    {}
func (s *GLSBService) DeleteGSLBVServerDomainBinding() {}
func (s *GLSBService) GetAllGSLBVServerDomainBinding() {}
func (s *GLSBService) GetGSLBVServerDomainBinding()    {}
func (s *GLSBService) CountGSLBVServerDomainBinding()  {}

// gslbvserver_gslbservicegroupmember_binding
func (s *GLSBService) GetAllGSLBVServerGSLBServiceGroupMemberBinding() {}
func (s *GLSBService) GetGSLBVServerGSLBServiceGroupMemberBinding()    {}
func (s *GLSBService) CountGSLBVServerGSLBServiceGroupMemberBinding()  {}

// gslbvserver_gslbservicegroup_binding
func (s *GLSBService) AddGSLBVServerGSLBServiceGroupBinding()    {}
func (s *GLSBService) DeleteGSLBVServerGSLBServiceGroupBinding() {}
func (s *GLSBService) GetAllGSLBVServerGSLBServiceGroupBinding() {}
func (s *GLSBService) GetGSLBVServerGSLBServiceGroupBinding()    {}
func (s *GLSBService) CountGSLBVServerGSLBServiceGroupBinding()  {}

// gslbvserver_gslbservice_binding
func (s *GLSBService) AddGSLBVServerGSLBServiceBinding()    {}
func (s *GLSBService) DeleteGSLBVServerGSLBServiceBinding() {}
func (s *GLSBService) GetAllGSLBVServerGSLBServiceBinding() {}
func (s *GLSBService) GetGSLBVServerGSLBServiceBinding()    {}
func (s *GLSBService) CountGSLBVServerGSLBServiceBinding()  {}

// gslbvserver_spilloverpolicy_binding
func (s *GLSBService) AddGSLBVServerSpilloverPolicyBinding()    {}
func (s *GLSBService) DeleteGSLBVServerSpilloverPolicyBinding() {}
func (s *GLSBService) GetAllGSLBVServerSpilloverPolicyBinding() {}
func (s *GLSBService) GetGSLBVServerSpilloverPolicyBinding()    {}
func (s *GLSBService) CountGSLBVServerSpilloverPolicyBinding()  {}
