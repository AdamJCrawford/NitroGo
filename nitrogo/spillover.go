package nitrogo

type SpilloverService struct {
	client *Client
}

// spilloveraction
func (s *SpilloverService) AddSpilloverAction()    {}
func (s *SpilloverService) DeleteSpilloverAction() {}
func (s *SpilloverService) GetAllSpilloverAction() {}
func (s *SpilloverService) GetSpilloverAction()    {}
func (s *SpilloverService) CountSpilloverAction()  {}
func (s *SpilloverService) RenameSpilloverAction() {}

// spilloverpolicy
func (s *SpilloverService) AddSpilloverPolicy()    {}
func (s *SpilloverService) DeleteSpilloverPolicy() {}
func (s *SpilloverService) UpdateSpilloverPolicy() {}
func (s *SpilloverService) UnsetSpilloverPolicy()  {}
func (s *SpilloverService) GetAllSpilloverPolicy() {}
func (s *SpilloverService) GetSpilloverPolicy()    {}
func (s *SpilloverService) CountSpilloverPolicy()  {}
func (s *SpilloverService) RenameSpilloverPolicy() {}

// spilloverpolicy_binding
func (s *SpilloverService) GetAllSpilloverPolicyBinding() {}
func (s *SpilloverService) GetSpilloverPolicyBinding()    {}

// spilloverpolicy_csvserver_binding
func (s *SpilloverService) GetAllSpilloverPolicyCSVServerBinding() {}
func (s *SpilloverService) GetSpilloverPolicyCSVServerBinding()    {}
func (s *SpilloverService) CountSpilloverPolicyCSVServerBinding()  {}

// spilloverpolicy_gslbvserver_binding
func (s *SpilloverService) GetAllSpilloverPolicyGSLBVServerBinding() {}
func (s *SpilloverService) GetSpilloverPolicyGSLBVServerBinding()    {}
func (s *SpilloverService) CountSpilloverPolicyGSLBVServerBinding()  {}

// spilloverpolicy_lbvserver_binding
func (s *SpilloverService) GetAllSpilloverPolicyLBVServerBinding() {}
func (s *SpilloverService) GetSpilloverPolicyLBVServerBinding()    {}
func (s *SpilloverService) CountSpilloverPolicyLBVServerBinding()  {}
