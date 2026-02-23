package nitrogo

const (
	icaAccessProfileURL           = "/nitro/v1/config/icaaccessprofile"
	icaActionURL                  = "/nitro/v1/config/icaaction"
	icaGlobalBindingURL           = "/nitro/v1/config/icaglobal_binding"
	icaGlobalICAPolicyBindingURL  = "/nitro/v1/config/icaglobal_icapolicy_binding"
	icaLatencyProfileURL          = "/nitro/v1/config/icalatencyprofile"
	icaParameterURL               = "/nitro/v1/config/icaparameter"
	icaPolicyURL                  = "/nitro/v1/config/icapolicy"
	icaPolicyBindingURL           = "/nitro/v1/config/icapolicy_binding"
	icaPolicyCRVServerBindingURL  = "/nitro/v1/config/icapolicy_crvserver_binding"
	icaPolicyICAGlobalBindingURL  = "/nitro/v1/config/icapolicy_icaglobal_binding"
	icaPolicyVPNVServerBindingURL = "/nitro/v1/config/icapolicy_vpnvserver_binding"
)

// ICA configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/ica/ica
type ICAService struct {
	client *Client
}

// ucaaccessprofile
func (s *ICAService) AddICAAccessProfile()    {}
func (s *ICAService) DeleteICAAccessProfile() {}
func (s *ICAService) UpdateICAAccessProfile() {}
func (s *ICAService) UnsetICAAccessProfile()  {}
func (s *ICAService) GetAllICAAccessProfile() {}
func (s *ICAService) GetICAAccessProfile()    {}
func (s *ICAService) CountICAAccessProfile()  {}

// icaaction
func (s *ICAService) AddICAAction()    {}
func (s *ICAService) DeleteICAAction() {}
func (s *ICAService) UpdateICAAction() {}
func (s *ICAService) UnsetICAAction()  {}
func (s *ICAService) GetAllICAAction() {}
func (s *ICAService) GetICAAction()    {}
func (s *ICAService) CountICAAction()  {}
func (s *ICAService) RenameICAAction() {}

// icaglobal_binding
func (s *ICAService) GetICAGlobalBinding() {}

// icaglobal_icapolicy_binding
func (s *ICAService) AddICAGlobalICAPolicyBinding()    {}
func (s *ICAService) DeleteICAGlobalICAPolicyBinding() {}
func (s *ICAService) GetICAGlobalICAPolicyBinding()    {}
func (s *ICAService) CountICAGlobalICAPolicyBinding()  {}

// icalatencyprofile
func (s *ICAService) AddICALatencyProfile()    {}
func (s *ICAService) DeleteICALatencyProfile() {}
func (s *ICAService) UpdateICALatencyProfile() {}
func (s *ICAService) UnsetICALatencyProfile()  {}
func (s *ICAService) GetAllICALatencyProfile() {}
func (s *ICAService) GetICALatencyProfile()    {}
func (s *ICAService) CountICALatencyProfile()  {}

// icaparameter
func (s *ICAService) UpdateICAParameter() {}
func (s *ICAService) UnsetICAParameter()  {}
func (s *ICAService) GetAllICAParameter() {}

// icapolicy
func (s *ICAService) AddICAPolicy()    {}
func (s *ICAService) DeleteICAPolicy() {}
func (s *ICAService) UpdateICAPolicy() {}
func (s *ICAService) UnsetICAPolicy()  {}
func (s *ICAService) GetAllICAPolicy() {}
func (s *ICAService) GetICAPolicy()    {}
func (s *ICAService) CountICAPolicy()  {}
func (s *ICAService) RenameICAPolicy() {}

// icapolicy_binding
func (s *ICAService) GetAllICAPolicyBinding() {}
func (s *ICAService) GetICAPolicyBinding()    {}

// icapolicy_crvserver_binding
func (s *ICAService) GetAllICAPolicyCRVServerBinding() {}
func (s *ICAService) GetICAPolicyCRVServerBinding()    {}
func (s *ICAService) CountICAPolicyCRVServerBinding()  {}

// icapolicy_icaglobal_binding
func (s *ICAService) GetAllICAPolicyICAGlobalBinding() {}
func (s *ICAService) GetICAPolicyICAGlobalBinding()    {}
func (s *ICAService) CountICAPolicyICAGlobalBinding()  {}

// icapolicy_vpnvserver_binding
func (s *ICAService) GetAllICAPolicyVPNVServerBinding() {}
func (s *ICAService) GetICAPolicyVPNVServerBinding()    {}
func (s *ICAService) CountICAPolicyVPNVServerBinding()  {}
