package nitrogo

const (
	cmpActionURL                          = "/nitro/v1/config/cmpaction"
	cmpGlobalBindingURL                   = "/nitro/v1/config/cmpglobal_binding"
	cmpGlobalCMPPolicyBindingURL          = "/nitro/v1/config/cmpglobal_cmppolicy_binding"
	cmpParameterURL                       = "/nitro/v1/config/cmpparameter"
	cmpPolicyURL                          = "/nitro/v1/config/cmppolicy"
	cmpPolicyLabelURL                     = "/nitro/v1/config/cmppolicylabel"
	cmpPolicyLabelBindingURL              = "/nitro/v1/config/cmppolicylabel_binding"
	cmpPolicyLabelCMPPolicyBindingURL     = "/nitro/v1/config/cmppolicylabel_cmppolicy_binding"
	cmpPolicyLabelPolicyBindingBindingURL = "/nitro/v1/config/cmppolicylabel_policybinding_binding"
	cmpPolicyBindingURL                   = "/nitro/v1/config/cmppolicy_binding"
	cmpPolicyCMPGlobalBindingURL          = "/nitro/v1/config/cmppolicy_cmpglobal_binding"
	cmpPolicyCMPPolicyLabelBindingURL     = "/nitro/v1/config/cmppolicy_cmppolicylabel_binding"
	cmpPolicyCRVServerBindingURL          = "/nitro/v1/config/cmppolicy_crvserver_binding"
	cmpPolicyCSVServerBindingURL          = "/nitro/v1/config/cmppolicy_csvserver_binding"
	cmpPolicyLBVServerBindingURL          = "/nitro/v1/config/cmppolicy_lbvserver_binding"
)

// Compression configuration. The system’s feature for compressing HTTP responses to compression-aware browsers.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/cmp/cmp
type CMPService struct {
	client *Client
}

// cmpaction
func (s *CMPService) AddCMPAction()    {}
func (s *CMPService) DeleteCMPAction() {}
func (s *CMPService) UpdateCMPAction() {}
func (s *CMPService) UnsetCMPAction()  {}
func (s *CMPService) GetAllCMPAction() {}
func (s *CMPService) GetCMPAction()    {}
func (s *CMPService) CountCMPAction()  {}
func (s *CMPService) RenameCMPAction() {}

// cmpglobal_binding
func (s *CMPService) GetCMPGlobalBinding() {}

// cmpglobal_cpmpolicy_binding
func (s *CMPService) AddCMPGlobalCMPPolicyBinding()    {}
func (s *CMPService) DeleteCMPGlobalCMPPolicyBinding() {}
func (s *CMPService) GetCMPGlobalCMPPolicyBinding()    {}
func (s *CMPService) CountCMPGlobalCMPPolicyBinding()  {}

// cmpparameter
func (s *CMPService) UpdateCMPParameter() {}
func (s *CMPService) UnsetCMPParameter()  {}
func (s *CMPService) GetAllCMPParameter() {}

// cmppolicy
func (s *CMPService) AddCOMPolicy()    {}
func (s *CMPService) DeleteCOMPolicy() {}
func (s *CMPService) UpdateCOMPolicy() {}
func (s *CMPService) GetAllCOMPolicy() {}
func (s *CMPService) GetCOMPolicy()    {}
func (s *CMPService) CountCOMPolicy()  {}
func (s *CMPService) RenameCOMPolicy() {}

// cmppolicylabel
func (s *CMPService) AddCMPPolicyLabel()    {}
func (s *CMPService) DeleteCMPPolicyLabel() {}
func (s *CMPService) GetAllCMPPolicyLabel() {}
func (s *CMPService) GetCMPPolicyLabel()    {}
func (s *CMPService) CountCMPPolicyLabel()  {}
func (s *CMPService) RenameCMPPolicyLabel() {}

// cmppolicylabel_binding
func (s *CMPService) GetAllCMPPolicyLabelBinding() {}
func (s *CMPService) GetCMPPolicyLabelBinding()    {}

// cmppolicylabel_cmppolicy_binding
func (s *CMPService) AddCMPPolicyLabelCMPPolicyBinding()    {}
func (s *CMPService) DeleteCMPPolicyLabelCMPPolicyBinding() {}
func (s *CMPService) GetAllCMPPolicyLabelCMPPolicyBinding() {}
func (s *CMPService) GetCMPPolicyLabelCMPPolicyBinding()    {}
func (s *CMPService) CountCMPPolicyLabelCMPPolicyBinding()  {}

// cmppolicylabel_policybinding_binding
func (s *CMPService) GetAllCMPPolicyLabelPolicyBindingBinding() {}
func (s *CMPService) GetCMPPolicyLabelPolicyBindingBinding()    {}
func (s *CMPService) CountCMPPolicyLabelPolicyBindingBinding()  {}

// cmppolicy_binding
func (s *CMPService) GetAllCMPPolicyBinding() {}
func (s *CMPService) GetCMPPolicyBinding()    {}

// cmppolicy_cmpglobal_binding
func (s *CMPService) GetAllCMPPolicyCMPGlobalBinding() {}
func (s *CMPService) GetCMPPolicyCMPGlobalBinding()    {}
func (s *CMPService) CountCMPPolicyCMPGlobalBinding()  {}

// cmppolicy_cmppolicylabel_binding
func (s *CMPService) GetAllCMPPolicyCMPPolicyLabelBinding() {}
func (s *CMPService) GetCMPPolicyCMPPolicyLabelBinding()    {}
func (s *CMPService) CountCMPPolicyCMPPolicyLabelBinding()  {}

// cmppolicy_crvserver_binding
func (s *CMPService) GetAllCMPPolicyCRVserverBinding() {}
func (s *CMPService) GetCMPPolicyCRVserverBinding()    {}
func (s *CMPService) CountCMPPolicyCRVserverBinding()  {}

// cmppolicy_csvserver_binding
func (s *CMPService) GetAllCMPPolicyCSVServerBinding() {}
func (s *CMPService) GetCMPPolicyCSVServerBinding()    {}
func (s *CMPService) CountCMPPolicyCSVServerBinding()  {}

// cmppolicy_lbvserver_binding
func (s *CMPService) GetAllCMPPolicyLBVServerBinding() {}
func (s *CMPService) GetCMPPolicyLBVServerBinding()    {}
func (s *CMPService) CountCMPPolicyLBVServerBinding()  {}
