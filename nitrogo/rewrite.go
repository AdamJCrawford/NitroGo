package nitrogo

// Rewrite configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/rewrite/rewrite
type RewriteService struct {
	client *Client
}

// rewriteaction
func (s *RewriteService) AddRewriteAction()    {}
func (s *RewriteService) DeleteRewriteAction() {}
func (s *RewriteService) UpdateRewriteAction() {}
func (s *RewriteService) UnsetRewriteAction()  {}
func (s *RewriteService) GetAllRewriteAction() {}
func (s *RewriteService) GetRewriteAction()    {}
func (s *RewriteService) CountRewriteAction()  {}
func (s *RewriteService) RenameRewriteAction() {}

// rewriteglobal_binding
func (s *RewriteService) GetRewriteGlobalBinding() {}

// rewriteglobal_rewritepolicy_binding
func (s *RewriteService) AddRewriteGlobalRewritePolicyBinding()    {}
func (s *RewriteService) DeleteRewriteGlobalRewritePolicyBinding() {}
func (s *RewriteService) GetRewriteGlobalRewritePolicyBinding()    {}
func (s *RewriteService) CountRewriteGlobalRewritePolicyBinding()  {}

// rewriteparam
func (s *RewriteService) UpdateRewriteParam() {}
func (s *RewriteService) UnsetRewriteParam()  {}
func (s *RewriteService) GetAllRewriteParam() {}

// rewritepolicy
func (s *RewriteService) AddRewritePolicy()    {}
func (s *RewriteService) DeleteRewritePolicy() {}
func (s *RewriteService) UpdateRewritePolicy() {}
func (s *RewriteService) UnsetRewritePolicy()  {}
func (s *RewriteService) GetAllRewritePolicy() {}
func (s *RewriteService) GetRewritePolicy()    {}
func (s *RewriteService) CountRewritePolicy()  {}
func (s *RewriteService) RenameRewritePolicy() {}

// rewritepolicylabel
func (s *RewriteService) AddRewritePolicyLabel()    {}
func (s *RewriteService) DeleteRewritePolicyLabel() {}
func (s *RewriteService) GetAllRewritePolicyLabel() {}
func (s *RewriteService) GetRewritePolicyLabel()    {}
func (s *RewriteService) CountRewritePolicyLabel()  {}
func (s *RewriteService) RenameRewritePolicyLabel() {}

// rewritepolicylabel_binding
func (s *RewriteService) GetAllRewritePolicyLabelBinding() {}
func (s *RewriteService) GetRewritePolicyLabelBinding()    {}

// rewritepolicylabel_policybinding_binding
func (s *RewriteService) GetAllRewritePolicyLabelPolicyBindingBinding() {}
func (s *RewriteService) GetRewritePolicyLabelPolicyBindingBinding()    {}
func (s *RewriteService) CountRewritePolicyLabelPolicyBindingBinding()  {}

// rewritepolicylabel-rewritepolicy_binding
func (s *RewriteService) AddRewritePolicyLabelRewritePolicyBinding()    {}
func (s *RewriteService) DeleteRewritePolicyLabelRewritePolicyBinding() {}
func (s *RewriteService) GetAllRewritePolicyLabelRewritePolicyBinding() {}
func (s *RewriteService) GetRewritePolicyLabelRewritePolicyBinding()    {}
func (s *RewriteService) CountRewritePolicyLabelRewritePolicyBinding()  {}

// rewritepolicy_binding
func (s *RewriteService) GetAllRewritePolicyBinding() {}
func (s *RewriteService) GetRewritePolicyBinding()    {}

// rewritepolicy_csvserver_binding
func (s *RewriteService) GetAllRewritePolicyCSVServerBinding() {}
func (s *RewriteService) GetRewritePolicyCSVServerBinding()    {}
func (s *RewriteService) CountRewritePolicyCSVServerBinding()  {}

// rewritepolicy_lbvserver_binding
func (s *RewriteService) GetAllRewritePolicyLBVServerBinding() {}
func (s *RewriteService) GetRewritePolicyLBVServerBinding()    {}
func (s *RewriteService) CountRewritePolicyLBVServerBinding()  {}

// rewritepolicy_rewriteglobal_binding
func (s *RewriteService) GetAllRewritePolicyRewriteGlobalBinding() {}
func (s *RewriteService) GetRewritePolicyRewriteGlobalBinding()    {}
func (s *RewriteService) CountRewritePolicyRewriteGlobalBinding()  {}

// rewritepolicy_rewritepolicylabel_binding
func (s *RewriteService) GetAllRewritePolicyRewritePolicyLabelBinding() {}
func (s *RewriteService) GetRewritePolicyRewritePolicyLabelBinding()    {}
func (s *RewriteService) CountRewritePolicyRewritePolicyLabelBinding()  {}

// rewritepolicy_vpnvserver_binding
func (s *RewriteService) GetAllRewritePolicyVPNVServerBinding() {}
func (s *RewriteService) GetRewritePolicyVPNVServerBinding()    {}
func (s *RewriteService) CountRewritePolicyVPNVServerBinding()  {}
