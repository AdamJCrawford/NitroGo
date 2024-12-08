package nitrogo

// Integrated Caching Configuration. The Integrated Caching feature is used to cache static and dynamic web application data.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/cache/cache
type CacheService struct {
	client *Client
}

// cachecontentgroup
func (s *CacheService) AddCacheContentGroup()    {}
func (s *CacheService) DeleteCacheContentGroup() {}
func (s *CacheService) UpdateCacheContentGroup() {}
func (s *CacheService) UnsetCacheContentGroup()  {}
func (s *CacheService) GetAllCacheContentGroup() {}
func (s *CacheService) GetCacheContentGroup()    {}
func (s *CacheService) CountCacheContentGroup()  {}
func (s *CacheService) ExpireCacheContentGroup() {}
func (s *CacheService) FlushCacheContentGroup()  {}
func (s *CacheService) SaveCacheContentGroup()   {}

// cacheforwardproxy
func (s *CacheService) AddCacheForwardProxy()    {}
func (s *CacheService) DeleteCacheForwardProxy() {}
func (s *CacheService) GetAllCacheForwardProxy() {}
func (s *CacheService) CountCacheForwardProxy()  {}

// cacheglobal_binding
func (s *CacheService) GetCacheGlobalBinding() {}

// cacheglobal_cachepolicy_binding
func (s *CacheService) AddCacheGlobalCachePolicyBinding()    {}
func (s *CacheService) DeleteCacheGlobalCachePolicyBinding() {}
func (s *CacheService) GetCacheGlobalCachePolicyBinding()    {}
func (s *CacheService) CountCacheGlobalCachePolicyBinding()  {}

// cacheobject
func (s *CacheService) GetAllCacheObject() {}
func (s *CacheService) CountCacheObject()  {}
func (s *CacheService) FlushCacheObject()  {}
func (s *CacheService) ExpireCacheObject() {}
func (s *CacheService) SaveCacheObject()   {}

// cacheparameter
func (s *CacheService) UpdateCacheParameter() {}
func (s *CacheService) UnsetCacheParameter()  {}
func (s *CacheService) GetAllCacheParameter() {}

// cachepolicy
func (s *CacheService) AddCachePolicy()    {}
func (s *CacheService) DeleteCachePolicy() {}
func (s *CacheService) UpdateCachePolicy() {}
func (s *CacheService) UnsetCachePolicy()  {}
func (s *CacheService) GetAllCachePolicy() {}
func (s *CacheService) GetCachePolicy()    {}
func (s *CacheService) CountCachePolicy()  {}
func (s *CacheService) RenameCachePolicy() {}

// cachepolicylabel
func (s *CacheService) AddCachePolicyLabel()    {}
func (s *CacheService) DeleteCachePolicyLabel() {}
func (s *CacheService) GetAllCachePolicyLabel() {}
func (s *CacheService) GetCachePolicyLabel()    {}
func (s *CacheService) CountCachePolicyLabel()  {}
func (s *CacheService) RenameCachePolicyLabel() {}

// cachepolicylabel_binding
func (s *CacheService) GetAllCachePolicyLabelBinding() {}
func (s *CacheService) GetCachePolicyLabelBinding()    {}

// cachepolicylabel_cachepolicy_binding
func (s *CacheService) AddCachePolicyLabelCachePolicyBinding()    {}
func (s *CacheService) DeleteCachePolicyLabelCachePolicyBinding() {}
func (s *CacheService) GetAllCachePolicyLabelCachePolicyBinding() {}
func (s *CacheService) GetCachePolicyLabelCachePolicyBinding()    {}
func (s *CacheService) CountCachePolicyLabelCachePolicyBinding()  {}

// cachepolicylabel_policybinding_binding
func (s *CacheService) GetAllCachePolicyLabelPolicyBindingBinding() {}
func (s *CacheService) GetCachePolicyLabelPolicyBindingBinding()    {}
func (s *CacheService) CountCachePolicyLabelPolicyBindingBinding()  {}

// cachepolicy_binding
func (s *CacheService) GetAllCachePolicyBinding() {}
func (s *CacheService) GetCachePolicyBinding()    {}

// cachepolicy_cacheglobal_binding
func (s *CacheService) GetAllCachePolicyCacheGlobalBinding() {}
func (s *CacheService) GetCachePolicyCacheGlobalBinding()    {}
func (s *CacheService) CountCachePolicyCacheGlobalBinding()  {}

// cachepolicy_cachepolicylabel_binding
func (s *CacheService) GetAllCachePolicyCachePolicyLabelBinding() {}
func (s *CacheService) GetCachePolicyCachePolicyLabelBinding()    {}
func (s *CacheService) CountCachePolicyCachePolicyLabelBinding()  {}

// cachepolicy_csvserver_binding
func (s *CacheService) GetAllCachePolicyCSVServerBinding() {}
func (s *CacheService) GetCachePolicyCSVServerBinding()    {}
func (s *CacheService) CountCachePolicyCSVServerBinding()  {}

// cachepolicy_lbvserver_binding
func (s *CacheService) GetAllCachePolicyLBVServerBinding() {}
func (s *CacheService) GetCachePolicyLBVServerBinding()    {}
func (s *CacheService) CountCachePolicyLBVServerBinding()  {}

// cacheselector
func (s *CacheService) AddCacheSelector()    {}
func (s *CacheService) DeleteCacheSelector() {}
func (s *CacheService) UpdateCacheSelector() {}
func (s *CacheService) GetAllCacheSelector() {}
func (s *CacheService) GetCacheSelector()    {}
func (s *CacheService) CountCacheSelector()  {}
