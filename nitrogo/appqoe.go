package nitrogo

// Application Level Quality of Experience configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appqoe/appqoe
type AppQOEService struct {
	client *Client
}

// appqoeaction
// Configuration for AppQoS action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appqoe/appqoeaction
func (s *AppQOEService) AddAppQOEAction()    {}
func (s *AppQOEService) DeleteAppQOEAction() {}
func (s *AppQOEService) UpdateAppQOEAction() {}
func (s *AppQOEService) UnsetAppQOEAction()  {}
func (s *AppQOEService) GetAllAppQOEAction() {}
func (s *AppQOEService) GetAppQOEAction()    {}
func (s *AppQOEService) CountAppQOEAction()  {}

// appqoecustomerresp
// Configuration for AppQoE custom response page resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appqoe/appqoecustomresp
func (s *AppQOEService) ImportAppQOECustomerResp() {}
func (s *AppQOEService) DeleteAppQOECustomerResp() {}
func (s *AppQOEService) GetAllAppQOECustomerResp() {}
func (s *AppQOEService) CountAppQOECustomerResp()  {}
func (s *AppQOEService) ChangeAppQOECustomerResp() {}

// appqoeparameter
// Configuration for QOS parameter resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appqoe/appqoeparameter
func (s *AppQOEService) UpdateAppQOEParameter() {}
func (s *AppQOEService) UnsetAppQOEParameter()  {}
func (s *AppQOEService) GetAllAppQOEParameter() {}

// appqoepolicy
// Configuration for AppQoS policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appqoe/appqoepolicy
func (s *AppQOEService) AddAppQOEPolicy()    {}
func (s *AppQOEService) DeleteAppQOEPolicy() {}
func (s *AppQOEService) UpdateAppQOEPolicy() {}
func (s *AppQOEService) GetAllAppQOEPolicy() {}
func (s *AppQOEService) GetAppQOEPolicy()    {}
func (s *AppQOEService) CountAppQOEPolicy()  {}

// appqoepolicy_binding
// Binding object which returns the resources bound to appqoepolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appqoe/appqoepolicy_binding
func (s *AppQOEService) GetAllAppQOEPolicyBinding() {}
func (s *AppQOEService) GetAppQOEPolicyBinding()    {}

// appqoepolicy_lbvserver_binding
// Binding object showing the lbvserver that can be bound to appqoepolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appqoe/appqoepolicy_lbvserver_binding
func (s *AppQOEService) GetAllAppQOEPolicyLBVServerBinding() {}
func (s *AppQOEService) GetAppQOEPolicyLBVServerBinding()    {}
func (s *AppQOEService) CountAppQOEPolicyLBVServerBinding()  {}
