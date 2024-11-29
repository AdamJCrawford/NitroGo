package nitrogo

type AppFlowService struct {
	client *Client
}

// func (s *AppFlowService)

// appflowaction
func (s *AppFlowService) AddAppFlowAction()    {}
func (s *AppFlowService) DeleteAppFlowAction() {}
func (s *AppFlowService) UpdateAppFlowAction() {}
func (s *AppFlowService) UnsetAppFlowAction()  {}
func (s *AppFlowService) RenameAppFlowAction() {}
func (s *AppFlowService) GetAllAppFlowAction() {}
func (s *AppFlowService) GetAppFlowAction()    {}
func (s *AppFlowService) CountAppFlowAction()  {}

// appflowaction_analyticsprofile_binding
func (s *AppFlowService) AddAppFlowActionAnalyticsProfileBinding()    {}
func (s *AppFlowService) DeleteAppFlowActionAnalyticsProfileBinding() {}
func (s *AppFlowService) GetAllAppFlowActionAnalyticsProfileBinding() {}
func (s *AppFlowService) GetAppFlowActionAnalyticsProfileBinding()    {}
func (s *AppFlowService) CountAppFlowActionAnalyticsProfileBinding()  {}

// appflowaction_binding
func (s *AppFlowService) GetAllAppFlowActionBinding() {}
func (s *AppFlowService) GetAppFlowActionBinding()    {}

// appflowcollector
func (s *AppFlowService) AddAppFlowCollector()    {}
func (s *AppFlowService) UpdateAppFlowCollector() {}
func (s *AppFlowService) UnsetAppFlowCollector()  {}
func (s *AppFlowService) DeleteAppFlowCollector() {}
func (s *AppFlowService) RenameAppFlowCollector() {}
func (s *AppFlowService) GetAllAppFlowCollector() {}
func (s *AppFlowService) GetAppFlowCollector()    {}
func (s *AppFlowService) CountAppFlowCollector()  {}

// appflowglobal_appflowpolicy_binding
func (s *AppFlowService) AddAppFlowGlobalAppFlowPolicyBinding()    {}
func (s *AppFlowService) DeleteAppFlowGlobalAppFlowPolicyBinding() {}
func (s *AppFlowService) GetAppFlowGlobalAppFlowPolicyBinding()    {}
func (s *AppFlowService) CountAppFlowGlobalAppFlowPolicyBinding()  {}

// appflowglobal_binding
func (s *AppFlowService) GetAppFlowGlobalBinding() {}

// appflowparam
func (s *AppFlowService) UpdateAppFlowParam() {}
func (s *AppFlowService) UnsetAppFlowParam()  {}
func (s *AppFlowService) GetAllAppFlowParam() {}

// appflowpolicy
func (s *AppFlowService) AddAppFlowPolicy()    {}
func (s *AppFlowService) DeleteAppFlowPolicy() {}
func (s *AppFlowService) UpdateAppFlowPolicy() {}
func (s *AppFlowService) UnsetAppFlowPolicy()  {}
func (s *AppFlowService) RenameAppFlowPolicy() {}
func (s *AppFlowService) GetAllAppFlowPolicy() {}
func (s *AppFlowService) GetAppFlowPolicy()    {}
func (s *AppFlowService) CountAppFlowPolicy()  {}

// appflowpolicylabel
func (s *AppFlowService) AddAppFlowPolicyLabel()    {}
func (s *AppFlowService) DeleteAppFlowPolicyLabel() {}
func (s *AppFlowService) RenameAppFlowPolicyLabel() {}
func (s *AppFlowService) GetAllAppFlowPolicyLabel() {}
func (s *AppFlowService) GetAppFlowPolicyLabel()    {}
func (s *AppFlowService) CountAppFlowPolicyLabel()  {}

// appflowpolicylabel_appflowpolicy_binding
func (s *AppFlowService) AddAppFlowPolicyLabelAppFlowPolicyBinding()    {}
func (s *AppFlowService) DeleteAppFlowPolicyLabelAppFlowPolicyBinding() {}
func (s *AppFlowService) GetAllAppFlowPolicyLabelAppFlowPolicyBinding() {}
func (s *AppFlowService) GetAppFlowPolicyLabelAppFlowPolicyBinding()    {}
func (s *AppFlowService) CountAppFlowPolicyLabelAppFlowPolicyBinding()  {}

// appflowpolicylabel_binding
func (s *AppFlowService) GetAllAppFlowPolicyLabelBinding() {}
func (s *AppFlowService) GetAppFlowPolicyLabelBinding()    {}

// appflowpolicy_appflowglobal_binding
func (s *AppFlowService) GetAllAppFlowPolicyAppFlowGlobalBinding() {}
func (s *AppFlowService) GetAppFlowPolicyAppFlowGlobalBinding()    {}
func (s *AppFlowService) CountAppFlowPolicyAppFlowGlobalBinding()  {}

// appflowpolicy_appflowpolicylabel_binding
func (s *AppFlowService) GetAllAppFlowPolicyAppFlowPolicyLabelBinding() {}
func (s *AppFlowService) GetAppFlowPolicyAppFlowPolicyLabelBinding()    {}
func (s *AppFlowService) CountAppFlowPolicyAppFlowPolicyLabelBinding()  {}

// appflowpolicy_binding
func (s *AppFlowService) GetAllAppFlowPolicyBinding() {}
func (s *AppFlowService) GetAppFlowPolicyBinding()    {}

// appflowpolicy_csvserver_binding
func (s *AppFlowService) GetAllAppFlowPolicyCSVServerBinding() {}
func (s *AppFlowService) GetAppFlowPolicyCSVServerBinding()    {}
func (s *AppFlowService) CountAppFlowPolicyCSVServerBinding()  {}

// appflowpolicy_lbvserver_binding
func (s *AppFlowService) GetAllAppFlowPolicy_LBVServerBinding() {}
func (s *AppFlowService) GetAppFlowPolicy_LBVServerBinding()    {}
func (s *AppFlowService) CountAppFlowPolicy_LBVServerBinding()  {}

// appflowpolicy_vpnvserver_binding
func (s *AppFlowService) GetAllAppFlowPolicyVPNVServerBinding() {}
func (s *AppFlowService) GetAppFlowPolicyVPNVServerBinding()    {}
func (s *AppFlowService) CountAppFlowPolicyVPNVServerBinding()  {}
