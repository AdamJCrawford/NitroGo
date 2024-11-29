package nitrogo

type FEOService struct {
	client *Client
}

// feoaction
func (s *FEOService) AddFEOAction()    {}
func (s *FEOService) DeleteFEOAction() {}
func (s *FEOService) UpdateFEOAction() {}
func (s *FEOService) UnsetFEOAction()  {}
func (s *FEOService) GetAllFEOAction() {}
func (s *FEOService) GetFEOAction()    {}
func (s *FEOService) CountFEOAction()  {}

// feoglobal_binding
func (s *FEOService) GetFEOGlobalBinding() {}

// feoglobal_feopolicy_binding
func (s *FEOService) AddFEOGlobalFEOPolicyBinding()    {}
func (s *FEOService) DeleteFEOGlobalFEOPolicyBinding() {}
func (s *FEOService) GetFEOGlobalFEOPolicyBinding()    {}
func (s *FEOService) CountFEOGlobalFEOPolicyBinding()  {}

// feoparameter
func (s *FEOService) UpdateFEOParameter() {}
func (s *FEOService) UnsetFEOParameter()  {}
func (s *FEOService) GetAllFEOParameter() {}

// feopolicy
func (s *FEOService) AddFEOPolicy()    {}
func (s *FEOService) DeleteFEOPolicy() {}
func (s *FEOService) UpdateFEOPolicy() {}
func (s *FEOService) UnsetFEOPolicy()  {}
func (s *FEOService) GetAllFEOPolicy() {}
func (s *FEOService) GetFEOPolicy()    {}
func (s *FEOService) CountFEOPolicy()  {}

// feopolicy_binding
func (s *FEOService) GetAllFEOPolicyBinding() {}
func (s *FEOService) GetFEOPolicyBinding()    {}

// feopolicy_csvserver_binding
func (s *FEOService) GetAllFEOPolicyCSVServerBinding() {}
func (s *FEOService) GetFEOPolicyCSVServerBinding()    {}
func (s *FEOService) CountFEOPolicyCSVServerBinding()  {}

// feopolicy_feoglobal_binding
func (s *FEOService) GetAllFEOPolicyFEOGlobalBinding() {}
func (s *FEOService) GetFEOPolicyFEOGlobalBinding()    {}
func (s *FEOService) CountFEOPolicyFEOGlobalBinding()  {}

// feopolicy_lbvserver_binding
func (s *FEOService) GetAllFEOPolicyLBVServerBinding() {}
func (s *FEOService) GetFEOPolicyLBVServerBinding()    {}
func (s *FEOService) CountFEOPolicyLBVServerBinding()  {}
