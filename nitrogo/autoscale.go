package nitrogo

type AutoscaleService struct {
	client *Client
}

// autoscaleaction
func (s *AutoscaleService) AddAutoscaleAction()    {}
func (s *AutoscaleService) DeleteAutoscaleAction() {}
func (s *AutoscaleService) UpdateAutoscaleAction() {}
func (s *AutoscaleService) UnsetAutoscaleAction()  {}
func (s *AutoscaleService) GetAllAutoscaleAction() {}
func (s *AutoscaleService) GetAutoscaleAction()    {}
func (s *AutoscaleService) CountAutoscaleAction()  {}

// autoscalepolicy
func (s *AutoscaleService) AddAutoscalePolicy()    {}
func (s *AutoscaleService) DeleteAutoscalePolicy() {}
func (s *AutoscaleService) UpdateAutoscalePolicy() {}
func (s *AutoscaleService) UnsetAutoscalePolicy()  {}
func (s *AutoscaleService) GetAllAutoscalePolicy() {}
func (s *AutoscaleService) GetAutoscalePolicy()    {}
func (s *AutoscaleService) CountAutoscalePolicy()  {}
func (s *AutoscaleService) RenameAutoscalePolicy() {}

// autoscalepolicy_binding
func (s *AutoscaleService) GetAllAutoscalePolicyBinding() {}
func (s *AutoscaleService) GetAutoscalePolicyBinding()    {}

// autoscalepolicy_nstimer_binding
func (s *AutoscaleService) GetAllAutoscalePolicyNSTimerBinding() {}
func (s *AutoscaleService) GetAutoscalePolicyNSTimerBinding()    {}
func (s *AutoscaleService) CountAutoscalePolicyNSTimerBinding()  {}

// autoscaleprofile
func (s *AutoscaleService) AddAutoscaleProfile()    {}
func (s *AutoscaleService) DeleteAutoscaleProfile() {}
func (s *AutoscaleService) UpdateAutoscaleProfile() {}
func (s *AutoscaleService) GetAllAutoscaleProfile() {}
func (s *AutoscaleService) GetAutoscaleProfile()    {}
func (s *AutoscaleService) CountAutoscaleProfile()  {}
