package nitrogo

// Cache redirection configuration. The cache redirection feature can transparently redirect cacheable HTTP
// requests to a cache and send non-cacheable or dynamic HTTP requests to the origin server. A cache stores
// or caches frequently requested web content and serves such web content to a client on behalf of the origin
// servers, alleviating the load on the origin server farm.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/cr/cr
type CRService struct {
	client *Client
}

// craction
func (s *CRService) GetAllCRAction() {}
func (s *CRService) GetCRAction()    {}
func (s *CRService) CountCRAction()  {}

// crpolicy
func (s *CRService) AddCRPolicy()    {}
func (s *CRService) DeleteCRPolicy() {}
func (s *CRService) UpdateCRPolicy() {}
func (s *CRService) UnsetCRPolicy()  {}
func (s *CRService) GetAllCRPolicy() {}
func (s *CRService) GetCRPolicy()    {}
func (s *CRService) CountCRPolicy()  {}
func (s *CRService) RenameCRPolicy() {}

// crpolicy_binding
func (s *CRService) GetAllCRPolicyBinding() {}
func (s *CRService) GetCRPolicyBinding()    {}

// crpolicy_crvserver_binding
func (s *CRService) GetAllCRPolicyCRVServerBinding() {}
func (s *CRService) GetCRPolicyCRVServerBinding()    {}
func (s *CRService) CountCRPolicyCRVServerBinding()  {}

// crvserver
func (s *CRService) AddCRVServer()     {}
func (s *CRService) DeleteCRVServer()  {}
func (s *CRService) UpdateCRVServer()  {}
func (s *CRService) UnsetCRVServer()   {}
func (s *CRService) EnableCRVServer()  {}
func (s *CRService) DisableCRVServer() {}
func (s *CRService) GetAllCRVServer()  {}
func (s *CRService) GetCRVServer()     {}
func (s *CRService) CountCRVServer()   {}
func (s *CRService) RenameCRVServer()  {}

// crvserver_analyticsprofile_binding
func (s *CRService) AddCRVServerAnalyticsProfileBinding()    {}
func (s *CRService) DeleteCRVServerAnalyticsProfileBinding() {}
func (s *CRService) GetAllCRVServerAnalyticsProfileBinding() {}
func (s *CRService) GetCRVServerAnalyticsProfileBinding()    {}
func (s *CRService) CountCRVServerAnalyticsProfileBinding()  {}

// crvserver_appflowpolicy_binding
func (s *CRService) AddCRVServerAppFlowPolicyBinding()    {}
func (s *CRService) DeleteCRVServerAppFlowPolicyBinding() {}
func (s *CRService) GetAllCRVServerAppFlowPolicyBinding() {}
func (s *CRService) GetCRVServerAppFlowPolicyBinding()    {}
func (s *CRService) CountCRVServerAppFlowPolicyBinding()  {}

// crvserver_appfwpolicy_binding
func (s *CRService) AddCRVServerAppFWPolicyBinding()    {}
func (s *CRService) DeleteCRVServerAppFWPolicyBinding() {}
func (s *CRService) GetAllCRVServerAppFWPolicyBinding() {}
func (s *CRService) GetCRVServerAppFWPolicyBinding()    {}
func (s *CRService) CountCRVServerAppFWPolicyBinding()  {}

// crvserver_appqoepolicy_binding
func (s *CRService) AddCRVServerAppQOEPolicyBinding()    {}
func (s *CRService) DeleteCRVServerAppQOEPolicyBinding() {}
func (s *CRService) GetAllCRVServerAppQOEPolicyBinding() {}
func (s *CRService) GetCRVServerAppQOEPolicyBinding()    {}
func (s *CRService) CountCRVServerAppQOEPolicyBinding()  {}

// crvserver_binding
func (s *CRService) GetAllCRVServerBinding() {}
func (s *CRService) GetCRVServerBinding()    {}

// crvserver_cachepolicy_binding
func (s *CRService) AddCRVServerCachePolicyBinding()    {}
func (s *CRService) DeleteCRVServerCachePolicyBinding() {}
func (s *CRService) GetAllCRVServerCachePolicyBinding() {}
func (s *CRService) GetCRVServerCachePolicyBinding()    {}
func (s *CRService) CountCRVServerCachePolicyBinding()  {}

// crvserver_cmppolicy_binding
func (s *CRService) AddCRVServerCMPPolicyBinding()    {}
func (s *CRService) DeleteCRVServerCMPPolicyBinding() {}
func (s *CRService) GetAllCRVServerCMPPolicyBinding() {}
func (s *CRService) GetCRVServerCMPPolicyBinding()    {}
func (s *CRService) CountCRVServerCMPPolicyBinding()  {}

// crvserver_crpolicy_binding
func (s *CRService) AddCRVServerCRPolicyBinding()    {}
func (s *CRService) DeleteCRVServerCRPolicyBinding() {}
func (s *CRService) GetAllCRVServerCRPolicyBinding() {}
func (s *CRService) GetCRVServerCRPolicyBinding()    {}
func (s *CRService) CountCRVServerCRPolicyBinding()  {}

// crvserver_cspolicy_binding
func (s *CRService) AddCRVServerCSPolicyBinding()    {}
func (s *CRService) DeleteCRVServerCSPolicyBinding() {}
func (s *CRService) GetAllCRVServerCSPolicyBinding() {}
func (s *CRService) GetCRVServerCSPolicyBinding()    {}
func (s *CRService) CountCRVServerCSPolicyBinding()  {}

// crvserver_feopolicy_binding
func (s *CRService) AddCRVServerFEOPolicyBinding()    {}
func (s *CRService) DeleteCRVServerFEOPolicyBinding() {}
func (s *CRService) GetAllCRVServerFEOPolicyBinding() {}
func (s *CRService) GetCRVServerFEOPolicyBinding()    {}
func (s *CRService) CountCRVServerFEOPolicyBinding()  {}

// crvserver_icapolicy_binding
func (s *CRService) AddCRVServerICAPolicyBinding()    {}
func (s *CRService) DeleteCRVServerICAPolicyBinding() {}
func (s *CRService) GetAllCRVServerICAPolicyBinding() {}
func (s *CRService) GetCRVServerICAPolicyBinding()    {}
func (s *CRService) CountCRVServerICAPolicyBinding()  {}

// crvserver_lbvserver_binding
func (s *CRService) AddCRVServerLBVServerBinding()    {}
func (s *CRService) DeleteCRVServerLBVServerBinding() {}
func (s *CRService) GetAllCRVServerLBVServerBinding() {}
func (s *CRService) GetCRVServerLBVServerBinding()    {}
func (s *CRService) CountCRVServerLBVServerBinding()  {}

// crvserver_policymap_binding
func (s *CRService) AddCRVServerPolicyMapBinding()    {}
func (s *CRService) DeleteCRVServerPolicyMapBinding() {}
func (s *CRService) GetAllCRVServerPolicyMapBinding() {}
func (s *CRService) GetCRVServerPolicyMapBinding()    {}
func (s *CRService) CountCRVServerPolicyMapBinding()  {}

// crvserver_responderpolicy_binding
func (s *CRService) AddCRVServerResponderPolicyBinding()    {}
func (s *CRService) DeleteCRVServerResponderPolicyBinding() {}
func (s *CRService) GetAllCRVServerResponderPolicyBinding() {}
func (s *CRService) GetCRVServerResponderPolicyBinding()    {}
func (s *CRService) CountCRVServerResponderPolicyBinding()  {}

// crvserver_rewritepolicy_binding
func (s *CRService) AddCRVServerRewritePolcyBinding()    {}
func (s *CRService) DeleteCRVServerRewritePolcyBinding() {}
func (s *CRService) GetAllCRVServerRewritePolcyBinding() {}
func (s *CRService) GetCRVServerRewritePolcyBinding()    {}
func (s *CRService) CountCRVServerRewritePolcyBinding()  {}

// crvserver_spilloverpolicy_binding
func (s *CRService) AddCRVServerSpilloverPolicyBinding()    {}
func (s *CRService) DeleteCRVServerSpilloverPolicyBinding() {}
func (s *CRService) GetAllCRVServerSpilloverPolicyBinding() {}
func (s *CRService) GetCRVServerSpilloverPolicyBinding()    {}
func (s *CRService) CountCRVServerSpilloverPolicyBinding()  {}
