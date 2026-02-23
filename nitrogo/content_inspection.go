package nitrogo

const (
	contentInspectionActionURL                                    = "/nitro/v1/config/contentinspectionaction"
	contentInspectionCalloutURL                                   = "/nitro/v1/config/contentinspectioncallout"
	contentInspectionGlobalBindingURL                             = "/nitro/v1/config/contentinspectionglobal_binding"
	contentInspectionGlobalContentInspectionPolicyBindingURL      = "/nitro/v1/config/contentinspectionglobal_contentinspectionpolicy_binding"
	contentInspectionParameterURL                                 = "/nitro/v1/config/contentinspectionparameter"
	contentInspectionPolicyURL                                    = "/nitro/v1/config/contentinspectionpolicy"
	contentInspectionPolicyLabelURL                               = "/nitro/v1/config/contentinspectionpolicylabel"
	contentInspectionPolicyLabelBindingURL                        = "/nitro/v1/config/contentinspectionpolicylabel_binding"
	contentInspectionPolicyLabelContentInspectionPolicyBindingURL = "/nitro/v1/config/contentinspectionpolicylabel_contentinspectionpolicy_binding"
	contentInspectionPolicyLabelPolicyBindingBindingURL           = "/nitro/v1/config/contentinspectionpolicylabel_policybinding_binding"
	contentInspectionPolicyBindingURL                             = "/nitro/v1/config/contentinspectionpolicy_binding"
	contentInspectionPolicyContentInspectionGlobalBindingURL      = "/nitro/v1/config/contentinspectionpolicy_contentinspectionglobal_binding"
	contentInspectionPolicyContentInspectionPolicyLabelBindingURL = "/nitro/v1/config/contentinspectionpolicy_contentinspectionpolicylabel_binding"
	contentInspectionPolicyCSVServerBindingURL                    = "/nitro/v1/config/contentinspectionpolicy_csvserver_binding"
	contentInspectionPolicyLBVServerBindingURL                    = "/nitro/v1/config/contentinspectionpolicy_lbvserver_binding"
	contentInspectionProfileURL                                   = "/nitro/v1/config/contentinspectionprofile"
)

// Content Inspection
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/contentinspection/contentinspection
type ContentInspectionService struct {
	client *Client
}

// contentinspectionaction
func (s *ContentInspectionService) AddContentInspectionAction()    {}
func (s *ContentInspectionService) DeleteContentInspectionAction() {}
func (s *ContentInspectionService) UpdateContentInspectionAction() {}
func (s *ContentInspectionService) UnsetContentInspectionAction()  {}
func (s *ContentInspectionService) GetAllContentInspectionAction() {}
func (s *ContentInspectionService) GetContentInspectionAction()    {}
func (s *ContentInspectionService) CountContentInspectionAction()  {}

// contentinspectioncallout
func (s *ContentInspectionService) AddContentInspectionCallout()    {}
func (s *ContentInspectionService) DeleteContentInspectionCallout() {}
func (s *ContentInspectionService) UpdateContentInspectionCallout() {}
func (s *ContentInspectionService) UnsetContentInspectionCallout()  {}
func (s *ContentInspectionService) GetAllContentInspectionCallout() {}
func (s *ContentInspectionService) GetContentInspectionCallout()    {}
func (s *ContentInspectionService) CountContentInspectionCallout()  {}

// contentinspectionglobal_binding
func (s *ContentInspectionService) GetContentInspectionGlobalBinding() {}

// contentinspectionglobal_contentinspectionpolicy_binding
func (s *ContentInspectionService) AddContentInspectionGlobalContentInspectionPolicyBinding()    {}
func (s *ContentInspectionService) DeleteContentInspectionGlobalContentInspectionPolicyBinding() {}
func (s *ContentInspectionService) GetContentInspectionGlobalContentInspectionPolicyBinding()    {}
func (s *ContentInspectionService) CountContentInspectionGlobalContentInspectionPolicyBinding()  {}

// contentinspectionparameter
func (s *ContentInspectionService) UpdateContentInspectionParameter() {}
func (s *ContentInspectionService) UnsetContentInspectionParameter()  {}
func (s *ContentInspectionService) GetAllContentInspectionParameter() {}

// contentinspectionpolicy
func (s *ContentInspectionService) AddContentInspectionPolicy()    {}
func (s *ContentInspectionService) DeleteContentInspectionPolicy() {}
func (s *ContentInspectionService) UpdateContentInspectionPolicy() {}
func (s *ContentInspectionService) UnsetContentInspectionPolicy()  {}
func (s *ContentInspectionService) GetAllContentInspectionPolicy() {}
func (s *ContentInspectionService) GetContentInspectionPolicy()    {}
func (s *ContentInspectionService) CountContentInspectionPolicy()  {}
func (s *ContentInspectionService) RenameContentInspectionPolicy() {}

// contentinspectionpolicylabel
func (s *ContentInspectionService) AddContentInspectionPolicyLabel()    {}
func (s *ContentInspectionService) DeleteContentInspectionPolicyLabel() {}
func (s *ContentInspectionService) GetAllContentInspectionPolicyLabel() {}
func (s *ContentInspectionService) GetContentInspectionPolicyLabel()    {}
func (s *ContentInspectionService) CountContentInspectionPolicyLabel()  {}
func (s *ContentInspectionService) RenameContentInspectionPolicyLabel() {}

// contentinspectionpolicylabel_binding
func (s *ContentInspectionService) GetAllContentInspectionPolicyLabelBinding() {}
func (s *ContentInspectionService) GetContentInspectionPolicyLabelBinding()    {}

// contentinspectionpolicylabel_contentinspectionpolicy_binding
func (s *ContentInspectionService) AddContentInspectionPolicyLabelContentInspectionPolicyBinding() {}
func (s *ContentInspectionService) DeleteContentInspectionPolicyLabelContentInspectionPolicyBinding() {
}
func (s *ContentInspectionService) GetAllContentInspectionPolicyLabelContentInspectionPolicyBinding() {
}
func (s *ContentInspectionService) GetContentInspectionPolicyLabelContentInspectionPolicyBinding() {}
func (s *ContentInspectionService) CountContentInspectionPolicyLabelContentInspectionPolicyBinding() {
}

// contentinspectionpolicylabel_policybinding_binding
func (s *ContentInspectionService) GetAllContentInspectionPolicyLabelPolicyBindingBinding() {}
func (s *ContentInspectionService) GetContentInspectionPolicyLabelPolicyBindingBinding()    {}
func (s *ContentInspectionService) CountContentInspectionPolicyLabelPolicyBindingBinding()  {}

// contentinspectionpolicy_binding
func (s *ContentInspectionService) GetAllContentInspectionPolicyBinding() {}
func (s *ContentInspectionService) GetContentInspectionPolicyBinding()    {}

// contentinspectionpolicy_contentinspectionglobal_binding
func (s *ContentInspectionService) GetAllContentInspectionPolcyContentInspectionGlobalBinding() {}
func (s *ContentInspectionService) GetContentInspectionPolcyContentInspectionGlobalBinding()    {}
func (s *ContentInspectionService) CountContentInspectionPolcyContentInspectionGlobalBinding()  {}

// contentinspectionpolicy_contentinspectionpolicylabel_binding
func (s *ContentInspectionService) GetAllContentInspectionPolicyContentInspectionPolicyLabelBinding() {
}
func (s *ContentInspectionService) GetContentInspectionPolicyContentInspectionPolicyLabelBinding() {}
func (s *ContentInspectionService) CountContentInspectionPolicyContentInspectionPolicyLabelBinding() {
}

// contentinspectionpolicy_csvserver_binding
func (s *ContentInspectionService) GetAllContentInspectionPolicyCSVServerBinding() {}
func (s *ContentInspectionService) GetContentInspectionPolicyCSVServerBinding()    {}
func (s *ContentInspectionService) CountContentInspectionPolicyCSVServerBinding()  {}

// contentinspectionpolicy_lbvserver_binding
func (s *ContentInspectionService) GetAllContentInspectionPolcyLBVServerBinding() {}
func (s *ContentInspectionService) GetContentInspectionPolcyLBVServerBinding()    {}
func (s *ContentInspectionService) CountContentInspectionPolcyLBVServerBinding()  {}

// contentinspectionprofile
func (s *ContentInspectionService) AddContentInspectionProfile()    {}
func (s *ContentInspectionService) DeleteContentInspectionProfile() {}
func (s *ContentInspectionService) UpdateContentInspectionProfile() {}
func (s *ContentInspectionService) UnsetContentInspectionProfile()  {}
func (s *ContentInspectionService) GetAllContentInspectionProfile() {}
func (s *ContentInspectionService) GetContentInspectionProfile()    {}
func (s *ContentInspectionService) CountContentInspectionProfile()  {}
