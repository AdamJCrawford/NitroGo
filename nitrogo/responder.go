package nitrogo

// Responder configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/responder/responder
type ResponderService struct {
	client *Client
}

// responderaction
func (s *ResponderService) AddResponderAction()    {}
func (s *ResponderService) DeleteResponderAction() {}
func (s *ResponderService) UpdateResponderAction() {}
func (s *ResponderService) UnsetResponderAction()  {}
func (s *ResponderService) GetAllResponderAction() {}
func (s *ResponderService) GetResponderAction()    {}
func (s *ResponderService) CountResponderAction()  {}
func (s *ResponderService) RenameResponderAction() {}

// responderglobal_binding
func (s *ResponderService) GetResponderGlobalBinding() {}

// responderglobal_responderpolicy_binding
func (s *ResponderService) AddResponderGlobalResponderPolicyBinding()    {}
func (s *ResponderService) DeleteResponderGlobalResponderPolicyBinding() {}
func (s *ResponderService) GetResponderGlobalResponderPolicyBinding()    {}
func (s *ResponderService) CountResponderGlobalResponderPolicyBinding()  {}

// responderhtmlpage
func (s *ResponderService) ImportResponderHTMLPage() {}
func (s *ResponderService) DeleteResponderHTMLPage() {}
func (s *ResponderService) GetAllResponderHTMLPage() {}
func (s *ResponderService) GetResponderHTMLPage()    {}
func (s *ResponderService) ChangeResponderHTMLPage() {}

// responderparam
func (s *ResponderService) UpdateResponderParam() {}
func (s *ResponderService) UnsetResponderParam()  {}
func (s *ResponderService) GetAllResponderParam() {}

// responderpolicy
func (s *ResponderService) AddResponderPolicy()    {}
func (s *ResponderService) DeleteResponderPolicy() {}
func (s *ResponderService) UpdateResponderPolicy() {}
func (s *ResponderService) UnsetResponderPolicy()  {}
func (s *ResponderService) GetAllResponderPolicy() {}
func (s *ResponderService) GetResponderPolicy()    {}
func (s *ResponderService) CountResponderPolicy()  {}
func (s *ResponderService) RenameResponderPolicy() {}

// responderpolicylabel
func (s *ResponderService) AddResponderPolicyLabel()    {}
func (s *ResponderService) DeleteResponderPolicyLabel() {}
func (s *ResponderService) GetAllResponderPolicyLabel() {}
func (s *ResponderService) GetResponderPolicyLabel()    {}
func (s *ResponderService) CountResponderPolicyLabel()  {}
func (s *ResponderService) RenameResponderPolicyLabel() {}

// responderpolicylabel_binding
func (s *ResponderService) GetAllResponderPolicyLabelBinding() {}
func (s *ResponderService) GetResponderPolicyLabelBinding()    {}

// respondrepolicylabel_policybinding_binding
func (s *ResponderService) GetAllResponderPolicyLabelPolicyBindingBinding() {}
func (s *ResponderService) GetResponderPolicyLabelPolicyBindingBinding()    {}
func (s *ResponderService) CountResponderPolicyLabelPolicyBindingBinding()  {}

// responderpolicylabel_responderpolicy_binding
func (s *ResponderService) AddResponderPolicyLabelResponderPolicyBinding()    {}
func (s *ResponderService) DeleteResponderPolicyLabelResponderPolicyBinding() {}
func (s *ResponderService) GetAllResponderPolicyLabelResponderPolicyBinding() {}
func (s *ResponderService) GetResponderPolicyLabelResponderPolicyBinding()    {}
func (s *ResponderService) CountResponderPolicyLabelResponderPolicyBinding()  {}

// responderpolicy_binding
func (s *ResponderService) GetAllResponderPolicyBinding() {}
func (s *ResponderService) GetResponderPolicyBinding()    {}

// responderpolicy_crvserver_binding
func (s *ResponderService) GetAllResponderPolicyCRVServerBinding() {}
func (s *ResponderService) GetResponderPolicyCRVServerBinding()    {}
func (s *ResponderService) CountResponderPolicyCRVServerBinding()  {}

// responderpolicy_csvserver_binding
func (s *ResponderService) GetAllResponderPolicyCSVServerBinding() {}
func (s *ResponderService) GetResponderPolicyCSVServerBinding()    {}
func (s *ResponderService) CountResponderPolicyCSVServerBinding()  {}

// responderpolicy_lbvserver_binding
func (s *ResponderService) GetAllResponderPolicyLBVServerBinding() {}
func (s *ResponderService) GetResponderPolicyLBVServerBinding()    {}
func (s *ResponderService) CountResponderPolicyLBVServerBinding()  {}

// responderpolicy_responderglobal_binding
func (s *ResponderService) GetAllResponderPolicyResponderGlobalBinding() {}
func (s *ResponderService) GetResponderPolicyResponderGlobalBinding()    {}
func (s *ResponderService) CountResponderPolicyResponderGlobalBinding()  {}

// responderpolicy_responderpolicylabel_binding
func (s *ResponderService) GetAllResponderPolicyResponderPolicyLabelBinding() {}
func (s *ResponderService) GetResponderPolicyResponderPolicyLabelBinding()    {}
func (s *ResponderService) CountResponderPolicyResponderPolicyLabelBinding()  {}

// responderpolicy_vpnvserver_binding
func (s *ResponderService) GetAllResponderPolicyVPNVServerBinding() {}
func (s *ResponderService) GetResponderPolicyVPNVServerBinding()    {}
func (s *ResponderService) CountResponderPolicyVPNVServerBinding()  {}
