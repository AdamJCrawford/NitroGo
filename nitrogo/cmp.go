package nitrogo

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
