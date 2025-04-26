package nitrogo

// Transform
// URL Transformation configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/transform/transform
type TransformService struct {
	client *Client
}

// transformaction
func (s *TransformService) AddTransformAction()    {}
func (s *TransformService) DeleteTransformAction() {}
func (s *TransformService) UpdateTransformAction() {}
func (s *TransformService) UnsetTransformAction()  {}
func (s *TransformService) GetAllTransformAction() {}
func (s *TransformService) GetTransformAction()    {}
func (s *TransformService) CountTransformAction()  {}

// transformglobal_binding
func (s *TransformService) GetTransformGlobalBinding() {}

// transformglobal_transformpolicy_binding
func (s *TransformService) AddTransformGlobalTransformPolicyBinding()    {}
func (s *TransformService) DeleteTransformGlobalTransformPolicyBinding() {}
func (s *TransformService) GetTransformGlobalTransformPolicyBinding()    {}
func (s *TransformService) CountTransformGlobalTransformPolicyBinding()  {}

// transformpolicy
func (s *TransformService) AddTransformPolicy()    {}
func (s *TransformService) DeleteTransformPolicy() {}
func (s *TransformService) UpdateTransformPolicy() {}
func (s *TransformService) UnsetTransformPolicy()  {}
func (s *TransformService) RenameTransformPolicy() {}
func (s *TransformService) GetAllTransformPolicy() {}
func (s *TransformService) GetTransformPolicy()    {}
func (s *TransformService) CountTransformPolicy()  {}

// transformpolicylabel
func (s *TransformService) AddTransformPolicyLabel()    {}
func (s *TransformService) DeleteTransformPolicyLabel() {}
func (s *TransformService) RenameTransformPolicyLabel() {}
func (s *TransformService) GetAllTransformPolicyLabel() {}
func (s *TransformService) GetTransformPolicyLabel()    {}
func (s *TransformService) CountTransformPolicyLabel()  {}

// transformpolicylabel_binding
func (s *TransformService) GetAllTransformPolicyLabelBinding() {}
func (s *TransformService) GetTransformPolicyLabelBinding()    {}

// transformpolicylabel_policybinding_binding
func (s *TransformService) GetAllTransformPolicyLabelPolicyBindingBinding() {}
func (s *TransformService) GetTransformPolicyLabelPolicyBindingBinding()    {}
func (s *TransformService) CountTransformPolicyLabelPolicyBindingBinding()  {}

// transformpolicytlabel_transformpolicy_binding
func (s *TransformService) AddTransformPolicyLabelTransformPolicyBinding()    {}
func (s *TransformService) DeleteTransformPolicyLabelTransformPolicyBinding() {}
func (s *TransformService) GetAllTransformPolicyLabelTransformPolicyBinding() {}
func (s *TransformService) GetTransformPolicyLabelTransformPolicyBinding()    {}
func (s *TransformService) CountTransformPolicyLabelTransformPolicyBinding()  {}

// transformpolicy_binding
func (s *TransformService) GetAllTransformPolicyBinding() {}
func (s *TransformService) GetTransformPolicyBinding()    {}

// transformpolicy_csvserver_binding
func (s *TransformService) GetAllTransformPolicyCSVServerBinding() {}
func (s *TransformService) GetTransformPolicyCSVServerBinding()    {}
func (s *TransformService) CountTransformPolicyCSVServerBinding()  {}

// transformpolicy_lbvserver_binding
func (s *TransformService) GetAllTransformPolicyLBVServerBinding() {}
func (s *TransformService) GetTransformPolicyLBVServerBinding()    {}
func (s *TransformService) CountTransformPolicyLBVServerBinding()  {}

// transformpolicy_transformglobal_binding
func (s *TransformService) GetAllTransformPolicyTransformGlobalBinding() {}
func (s *TransformService) GetTransformPolicyTransformGlobalBinding()    {}
func (s *TransformService) CountTransformPolicyTransformGlobalBinding()  {}

// transformpolicy_transformpolicylabel_binding
func (s *TransformService) GetAllTransformPolicyTransformPolicyLabelBinding() {}
func (s *TransformService) GetTransformPolicyTransformPolicyLabelBinding()    {}
func (s *TransformService) CountTransformPolicyTransformPolicyLabelBinding()  {}

// transformprofile
func (s *TransformService) AddTransformProfile()    {}
func (s *TransformService) DeleteTransformProfile() {}
func (s *TransformService) UpdateTransformProfile() {}
func (s *TransformService) UnsetTransformProfile()  {}
func (s *TransformService) GetAllTransformProfile() {}
func (s *TransformService) GetTransformProfile()    {}
func (s *TransformService) CountTransformProfile()  {}

// transformprofile_binding
func (s *TransformService) GetAllTransformProfileBinding() {}
func (s *TransformService) GetTransformProfileBinding()    {}

// transformprofile_transformaction_binding
func (s *TransformService) GetAllTransformProfileTransformActionBinding() {}
func (s *TransformService) GetTransformProfileTransformActionBinding()    {}
func (s *TransformService) CountTransformProfileTransformActionBinding()  {}
