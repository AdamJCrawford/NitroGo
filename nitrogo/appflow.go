package nitrogo

// AppFlow configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appflow/appflow
type AppFlowService struct {
	client *Client
}

// appflowaction
// Configuration for AppFlow action resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appflow/appflowaction
func (s *AppFlowService) AddAppFlowAction()    {}
func (s *AppFlowService) DeleteAppFlowAction() {}
func (s *AppFlowService) UpdateAppFlowAction() {}
func (s *AppFlowService) UnsetAppFlowAction()  {}
func (s *AppFlowService) RenameAppFlowAction() {}
func (s *AppFlowService) GetAllAppFlowAction() {}
func (s *AppFlowService) GetAppFlowAction()    {}
func (s *AppFlowService) CountAppFlowAction()  {}

// appflowaction_analyticsprofile_binding
// Binding object showing the analyticsprofile that can be bound to appflowaction.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appflow/appflowaction_analyticsprofile_binding
func (s *AppFlowService) AddAppFlowActionAnalyticsProfileBinding()    {}
func (s *AppFlowService) DeleteAppFlowActionAnalyticsProfileBinding() {}
func (s *AppFlowService) GetAllAppFlowActionAnalyticsProfileBinding() {}
func (s *AppFlowService) GetAppFlowActionAnalyticsProfileBinding()    {}
func (s *AppFlowService) CountAppFlowActionAnalyticsProfileBinding()  {}

// appflowaction_binding
// Binding object which returns the resources bound to appflowaction.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appflow/appflowaction_binding
func (s *AppFlowService) GetAllAppFlowActionBinding() {}
func (s *AppFlowService) GetAppFlowActionBinding()    {}

// appflowcollector
// Configuration for AppFlow collector resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appflow/appflowcollector
func (s *AppFlowService) AddAppFlowCollector()    {}
func (s *AppFlowService) UpdateAppFlowCollector() {}
func (s *AppFlowService) UnsetAppFlowCollector()  {}
func (s *AppFlowService) DeleteAppFlowCollector() {}
func (s *AppFlowService) RenameAppFlowCollector() {}
func (s *AppFlowService) GetAllAppFlowCollector() {}
func (s *AppFlowService) GetAppFlowCollector()    {}
func (s *AppFlowService) CountAppFlowCollector()  {}

// appflowglobal_appflowpolicy_binding
// Binding object showing the appflowpolicy that can be bound to appflowglobal.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appflow/appflowglobal_appflowpolicy_binding
func (s *AppFlowService) AddAppFlowGlobalAppFlowPolicyBinding()    {}
func (s *AppFlowService) DeleteAppFlowGlobalAppFlowPolicyBinding() {}
func (s *AppFlowService) GetAppFlowGlobalAppFlowPolicyBinding()    {}
func (s *AppFlowService) CountAppFlowGlobalAppFlowPolicyBinding()  {}

// appflowglobal_binding
// Binding object which returns the resources bound to appflowglobal.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appflow/appflowglobal_binding
func (s *AppFlowService) GetAppFlowGlobalBinding() {}

// appflowparam
// Configuration for AppFlow parameter resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appflow/appflowparam
func (s *AppFlowService) UpdateAppFlowParam() {}
func (s *AppFlowService) UnsetAppFlowParam()  {}
func (s *AppFlowService) GetAllAppFlowParam() {}

// appflowpolicy
// Configuration for AppFlow policy resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appflow/appflowpolicy
func (s *AppFlowService) AddAppFlowPolicy()    {}
func (s *AppFlowService) DeleteAppFlowPolicy() {}
func (s *AppFlowService) UpdateAppFlowPolicy() {}
func (s *AppFlowService) UnsetAppFlowPolicy()  {}
func (s *AppFlowService) RenameAppFlowPolicy() {}
func (s *AppFlowService) GetAllAppFlowPolicy() {}
func (s *AppFlowService) GetAppFlowPolicy()    {}
func (s *AppFlowService) CountAppFlowPolicy()  {}

// appflowpolicylabel
// Configuration for AppFlow policy label resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appflow/appflowpolicylabel
func (s *AppFlowService) AddAppFlowPolicyLabel()    {}
func (s *AppFlowService) DeleteAppFlowPolicyLabel() {}
func (s *AppFlowService) RenameAppFlowPolicyLabel() {}
func (s *AppFlowService) GetAllAppFlowPolicyLabel() {}
func (s *AppFlowService) GetAppFlowPolicyLabel()    {}
func (s *AppFlowService) CountAppFlowPolicyLabel()  {}

// appflowpolicylabel_appflowpolicy_binding
// Binding object showing the appflowpolicy that can be bound to appflowpolicylabel.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appflow/appflowpolicylabel_appflowpolicy_binding
func (s *AppFlowService) AddAppFlowPolicyLabelAppFlowPolicyBinding()    {}
func (s *AppFlowService) DeleteAppFlowPolicyLabelAppFlowPolicyBinding() {}
func (s *AppFlowService) GetAllAppFlowPolicyLabelAppFlowPolicyBinding() {}
func (s *AppFlowService) GetAppFlowPolicyLabelAppFlowPolicyBinding()    {}
func (s *AppFlowService) CountAppFlowPolicyLabelAppFlowPolicyBinding()  {}

// appflowpolicylabel_binding
// Binding object which returns the resources bound to appflowpolicylabel.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appflow/appflowpolicylabel_binding
func (s *AppFlowService) GetAllAppFlowPolicyLabelBinding() {}
func (s *AppFlowService) GetAppFlowPolicyLabelBinding()    {}

// appflowpolicy_appflowglobal_binding
// Binding object showing the appflowglobal that can be bound to appflowpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appflow/appflowpolicy_appflowglobal_binding
func (s *AppFlowService) GetAllAppFlowPolicyAppFlowGlobalBinding() {}
func (s *AppFlowService) GetAppFlowPolicyAppFlowGlobalBinding()    {}
func (s *AppFlowService) CountAppFlowPolicyAppFlowGlobalBinding()  {}

// appflowpolicy_appflowpolicylabel_binding
// Binding object showing the appflowpolicylabel that can be bound to appflowpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appflow/appflowpolicy_appflowpolicylabel_binding
func (s *AppFlowService) GetAllAppFlowPolicyAppFlowPolicyLabelBinding() {}
func (s *AppFlowService) GetAppFlowPolicyAppFlowPolicyLabelBinding()    {}
func (s *AppFlowService) CountAppFlowPolicyAppFlowPolicyLabelBinding()  {}

// appflowpolicy_binding
// Binding object which returns the resources bound to appflowpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appflow/appflowpolicy_binding
func (s *AppFlowService) GetAllAppFlowPolicyBinding() {}
func (s *AppFlowService) GetAppFlowPolicyBinding()    {}

// appflowpolicy_csvserver_binding
// Binding object showing the csvserver that can be bound to appflowpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appflow/appflowpolicy_csvserver_binding
func (s *AppFlowService) GetAllAppFlowPolicyCSVServerBinding() {}
func (s *AppFlowService) GetAppFlowPolicyCSVServerBinding()    {}
func (s *AppFlowService) CountAppFlowPolicyCSVServerBinding()  {}

// appflowpolicy_lbvserver_binding
// Binding object showing the lbvserver that can be bound to appflowpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appflow/appflowpolicy_lbvserver_binding
func (s *AppFlowService) GetAllAppFlowPolicy_LBVServerBinding() {}
func (s *AppFlowService) GetAppFlowPolicy_LBVServerBinding()    {}
func (s *AppFlowService) CountAppFlowPolicy_LBVServerBinding()  {}

// appflowpolicy_vpnvserver_binding
// Binding object showing the vpnvserver that can be bound to appflowpolicy.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/appflow/appflowpolicy_vpnvserver_binding
func (s *AppFlowService) GetAllAppFlowPolicyVPNVServerBinding() {}
func (s *AppFlowService) GetAppFlowPolicyVPNVServerBinding()    {}
func (s *AppFlowService) CountAppFlowPolicyVPNVServerBinding()  {}
