package nitrogo

const (
	spilloverActionURL                   = "/nitro/v1/config/spilloveraction"
	spilloverPolicyURL                   = "/nitro/v1/config/spilloverpolicy"
	spilloverPolicyBindingURL            = "/nitro/v1/config/spilloverpolicy_binding"
	spilloverPolicyCSVServerBindingURL   = "/nitro/v1/config/spilloverpolicy_csvserver_binding"
	spilloverPolicyGSLBVServerBindingURL = "/nitro/v1/config/spilloverpolicy_gslbvserver_binding"
	spilloverPolicyLBVServerBindingURL   = "/nitro/v1/config/spilloverpolicy_lbvserver_binding"
)

// Spillover policies and actions.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/spillover/spillover
type SpilloverService struct {
	client *Client
}

// spilloveraction
// Configuration for Spillover action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/spillover/spilloveraction
func (s *SpilloverService) AddSpilloverAction()    {}
func (s *SpilloverService) DeleteSpilloverAction() {}
func (s *SpilloverService) GetAllSpilloverAction() {}
func (s *SpilloverService) GetSpilloverAction()    {}
func (s *SpilloverService) CountSpilloverAction()  {}
func (s *SpilloverService) RenameSpilloverAction() {}

// spilloverpolicy
// Configuration for Spillover policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/spillover/spilloverpolicy
func (s *SpilloverService) AddSpilloverPolicy()    {}
func (s *SpilloverService) DeleteSpilloverPolicy() {}
func (s *SpilloverService) UpdateSpilloverPolicy() {}
func (s *SpilloverService) UnsetSpilloverPolicy()  {}
func (s *SpilloverService) GetAllSpilloverPolicy() {}
func (s *SpilloverService) GetSpilloverPolicy()    {}
func (s *SpilloverService) CountSpilloverPolicy()  {}
func (s *SpilloverService) RenameSpilloverPolicy() {}

// spilloverpolicy_binding
// Binding object which returns the resources bound to spilloverpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/spillover/spilloverpolicy_binding
func (s *SpilloverService) GetAllSpilloverPolicyBinding() {}
func (s *SpilloverService) GetSpilloverPolicyBinding()    {}

// spilloverpolicy_csvserver_binding
// Binding object showing the csvserver that can be bound to spilloverpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/spillover/spilloverpolicy_csvserver_binding
func (s *SpilloverService) GetAllSpilloverPolicyCSVServerBinding() {}
func (s *SpilloverService) GetSpilloverPolicyCSVServerBinding()    {}
func (s *SpilloverService) CountSpilloverPolicyCSVServerBinding()  {}

// spilloverpolicy_gslbvserver_binding
// Binding object showing the gslbvserver that can be bound to spilloverpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/spillover/spilloverpolicy_gslbvserver_binding
func (s *SpilloverService) GetAllSpilloverPolicyGSLBVServerBinding() {}
func (s *SpilloverService) GetSpilloverPolicyGSLBVServerBinding()    {}
func (s *SpilloverService) CountSpilloverPolicyGSLBVServerBinding()  {}

// spilloverpolicy_lbvserver_binding
// Binding object showing the lbvserver that can be bound to spilloverpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/spillover/spilloverpolicy_lbvserver_binding
func (s *SpilloverService) GetAllSpilloverPolicyLBVServerBinding() {}
func (s *SpilloverService) GetSpilloverPolicyLBVServerBinding()    {}
func (s *SpilloverService) CountSpilloverPolicyLBVServerBinding()  {}
