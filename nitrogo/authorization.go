package nitrogo

// Authorization configuration. Authorization services check which resources users are authorized to access, and grant permissions accordingly.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authorization/authorization
type AuthorizationService struct {
	client *Client
}

// authorizationaction
// Configuration for authorization action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authorization/authorizationaction
func (s *AuthorizationService) GetAllAuthorizationAction() {}
func (s *AuthorizationService) GetAuthorizationAction()    {}
func (s *AuthorizationService) CountAuthorizationAction()  {}

// authorizationpolicy
// Configuration for authorization policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authorization/authorizationpolicy
func (s *AuthorizationService) AddAuthorizationPolicy()    {}
func (s *AuthorizationService) DeleteAuthorizationPolicy() {}
func (s *AuthorizationService) UpdateAuthorizationPolicy() {}
func (s *AuthorizationService) RenameAuthorizationPolicy() {}
func (s *AuthorizationService) GetAllAuthorizationPolicy() {}
func (s *AuthorizationService) GetAuthorizationPolicy()    {}
func (s *AuthorizationService) CountAuthorizationPolicy()  {}

// authorizationpolicylabel
// Configuration for authorization policy label resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authorization/authorizationpolicylabel
func (s *AuthorizationService) AddAuthorizationPolicyLabel()    {}
func (s *AuthorizationService) DeleteAuthorizationPolicyLabel() {}
func (s *AuthorizationService) RenameAuthorizationPolicyLabel() {}
func (s *AuthorizationService) GetAllAuthorizationPolicyLabel() {}
func (s *AuthorizationService) GetAuthorizationPolicyLabel()    {}
func (s *AuthorizationService) CountAuthorizationPolicyLabel()  {}

// authorizationpolicylabel_authorizationpolicy_binding
// Binding object showing the authorizationpolicy that can be bound to authorizationpolicylabel.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authorization/authorizationpolicylabel_authorizationpolicy_binding
func (s *AuthorizationService) AddAuthorizationPolicyLabelAuthorizationPolicyBinding()    {}
func (s *AuthorizationService) DeleteAuthorizationPolicyLabelAuthorizationPolicyBinding() {}
func (s *AuthorizationService) GetAllAuthorizationPolicyLabelAuthorizationPolicyBinding() {}
func (s *AuthorizationService) GetAuthorizationPolicyLabelAuthorizationPolicyBinding()    {}
func (s *AuthorizationService) CountAuthorizationPolicyLabelAuthorizationPolicyBinding()  {}

// authorizationpolicylabel_binding
// Binding object which returns the resources bound to authorizationpolicylabel.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authorization/authorizationpolicylabel_binding
func (s *AuthorizationService) GetAllAuthorizationPolicyLabelBinding() {}
func (s *AuthorizationService) GetAuthorizationPolicyLabelBinding()    {}

// authorizationpolicy_aaagroup_binding
// Binding object showing the aaagroup that can be bound to authorizationpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authorization/authorizationpolicy_aaagroup_binding
func (s *AuthorizationService) GetAllAuthorizationPolicyAAAGroupBinding() {}
func (s *AuthorizationService) GetAuthorizationPolicyAAAGroupBinding()    {}
func (s *AuthorizationService) CountAuthorizationPolicyAAAGroupBinding()  {}

// authorizationpolicy_aaauser_binding
// Binding object showing the aaauser that can be bound to authorizationpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authorization/authorizationpolicy_aaauser_binding
func (s *AuthorizationService) GetAllAuthorizationPolicyAAAUserBinding() {}
func (s *AuthorizationService) GetAuthorizationPolicyAAAUserBinding()    {}
func (s *AuthorizationService) CountAuthorizationPolicyAAAUserBinding()  {}

// authorizationpolicy_authorizationpolicylabel_binding
// Binding object showing the authorizationpolicylabel that can be bound to authorizationpolicy
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authorization/authorizationpolicy_authorizationpolicylabel_binding
func (s *AuthorizationService) GetAllAuthorizationPolicyAuthorizationPolicyLabelBinding() {}
func (s *AuthorizationService) GetAuthorizationPolicyAuthorizationPolicyLabelBinding()    {}
func (s *AuthorizationService) CountAuthorizationPolicyAuthorizationPolicyLabelBinding()  {}

// authorizationpolicy_binding
// Binding object which returns the resources bound to authorizationpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authorization/authorizationpolicy_binding
func (s *AuthorizationService) GetAllAuthorizationPolicyBinding() {}
func (s *AuthorizationService) GetAuthorizationPolicyBinding()    {}

// authorizationpolicy_csvserver_binding
// Binding object showing the csvserver that can be bound to authorizationpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authorization/authorizationpolicy_csvserver_binding
func (s *AuthorizationService) GetAllAuthorizationPolicyCSVServerBinding() {}
func (s *AuthorizationService) GetAuthorizationPolicyCSVServerBinding()    {}
func (s *AuthorizationService) CountAuthorizationPolicyCSVServerBinding()  {}

// authorizationpolicy_lbvserver_binding
// Binding object showing the lbvserver that can be bound to authorizationpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/authorization/authorizationpolicy_lbvserver_binding
func (s *AuthorizationService) GetAllAuthorizationPolicyLBVServerBinding() {}
func (s *AuthorizationService) GetAuthorizationPolicyLBVServerBinding()    {}
func (s *AuthorizationService) CountAuthorizationPolicyLBVServerBinding()  {}
