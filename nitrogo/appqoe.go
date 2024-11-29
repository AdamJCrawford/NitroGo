package nitrogo

type AppQOEService struct {
	client *Client
}

// appqoeaction
func (s *AppQOEService) AddAppQOEAction()    {}
func (s *AppQOEService) DeleteAppQOEAction() {}
func (s *AppQOEService) UpdateAppQOEAction() {}
func (s *AppQOEService) UnsetAppQOEAction()  {}
func (s *AppQOEService) GetAllAppQOEAction() {}
func (s *AppQOEService) GetAppQOEAction()    {}
func (s *AppQOEService) CountAppQOEAction()  {}

// appqoecustomerresp
func (s *AppQOEService) ImportAppQOECustomerResp() {}
func (s *AppQOEService) DeleteAppQOECustomerResp() {}
func (s *AppQOEService) GetAllAppQOECustomerResp() {}
func (s *AppQOEService) CountAppQOECustomerResp()  {}
func (s *AppQOEService) ChangeAppQOECustomerResp() {}

// appqoeparameter
func (s *AppQOEService) UpdateAppQOEParameter() {}
func (s *AppQOEService) UnsetAppQOEParameter()  {}
func (s *AppQOEService) GetAllAppQOEParameter() {}

// appqoepolicy
func (s *AppQOEService) AddAppQOEPolicy()    {}
func (s *AppQOEService) DeleteAppQOEPolicy() {}
func (s *AppQOEService) UpdateAppQOEPolicy() {}
func (s *AppQOEService) GetAllAppQOEPolicy() {}
func (s *AppQOEService) GetAppQOEPolicy()    {}
func (s *AppQOEService) CountAppQOEPolicy()  {}

// appqoepolicy_binding
func (s *AppQOEService) GetAllAppQOEPolicyBinding() {}
func (s *AppQOEService) GetAppQOEPolicyBinding()    {}

// appqoepolicy_lbvserver_binding
func (s *AppQOEService) GetAllAppQOEPolicyLBVServerBinding() {}
func (s *AppQOEService) GetAppQOEPolicyLBVServerBinding()    {}
func (s *AppQOEService) CountAppQOEPolicyLBVServerBinding()  {}
