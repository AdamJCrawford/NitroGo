package nitrogo

// Front end optimization configuration. The system’s feature to optimize Web content for performance.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/feo/feo
type FEOService struct {
	client *Client
}

// feoaction
// Configuration for Front end optimization action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/feo/feoaction
func (s *FEOService) AddFEOAction()    {}
func (s *FEOService) DeleteFEOAction() {}
func (s *FEOService) UpdateFEOAction() {}
func (s *FEOService) UnsetFEOAction()  {}
func (s *FEOService) GetAllFEOAction() {}
func (s *FEOService) GetFEOAction()    {}
func (s *FEOService) CountFEOAction()  {}

// feoglobal_binding
// Binding object which returns the resources bound to feoglobal.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/feo/feoglobal_binding
func (s *FEOService) GetFEOGlobalBinding() {}

// feoglobal_feopolicy_binding
// Binding object showing the feopolicy that can be bound to feoglobal.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/feo/feoglobal_feopolicy_binding
func (s *FEOService) AddFEOGlobalFEOPolicyBinding()    {}
func (s *FEOService) DeleteFEOGlobalFEOPolicyBinding() {}
func (s *FEOService) GetFEOGlobalFEOPolicyBinding()    {}
func (s *FEOService) CountFEOGlobalFEOPolicyBinding()  {}

// feoparameter
// Configuration for FEO parameter resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/feo/feoparameter
func (s *FEOService) UpdateFEOParameter() {}
func (s *FEOService) UnsetFEOParameter()  {}
func (s *FEOService) GetAllFEOParameter() {}

// feopolicy
// Configuration for Front end optimization policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/feo/feopolicy
func (s *FEOService) AddFEOPolicy()    {}
func (s *FEOService) DeleteFEOPolicy() {}
func (s *FEOService) UpdateFEOPolicy() {}
func (s *FEOService) UnsetFEOPolicy()  {}
func (s *FEOService) GetAllFEOPolicy() {}
func (s *FEOService) GetFEOPolicy()    {}
func (s *FEOService) CountFEOPolicy()  {}

// feopolicy_binding
// Binding object which returns the resources bound to feopolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/feo/feopolicy_binding
func (s *FEOService) GetAllFEOPolicyBinding() {}
func (s *FEOService) GetFEOPolicyBinding()    {}

// feopolicy_csvserver_binding
// Binding object showing the csvserver that can be bound to feopolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/feo/feopolicy_csvserver_binding
func (s *FEOService) GetAllFEOPolicyCSVServerBinding() {}
func (s *FEOService) GetFEOPolicyCSVServerBinding()    {}
func (s *FEOService) CountFEOPolicyCSVServerBinding()  {}

// feopolicy_feoglobal_binding
// Binding object showing the feoglobal that can be bound to feopolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/feo/feopolicy_feoglobal_binding
func (s *FEOService) GetAllFEOPolicyFEOGlobalBinding() {}
func (s *FEOService) GetFEOPolicyFEOGlobalBinding()    {}
func (s *FEOService) CountFEOPolicyFEOGlobalBinding()  {}

// feopolicy_lbvserver_binding
// Binding object showing the lbvserver that can be bound to feopolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/feo/feopolicy_lbvserver_binding
func (s *FEOService) GetAllFEOPolicyLBVServerBinding() {}
func (s *FEOService) GetFEOPolicyLBVServerBinding()    {}
func (s *FEOService) CountFEOPolicyLBVServerBinding()  {}
