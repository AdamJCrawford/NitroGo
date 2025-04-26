package nitrogo

// VideoOptimization
// Video optimization feature is used to show (i) the stats of different media types that are being served by the Citrix ADC and (ii) the details of optimization applied on ABR videos
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/videooptimization/videooptimization
type VideoOptimizationService struct {
	client *Client
}

// videooptimizationdetectionaction
func (s *VideoOptimizationService) AddVideoOptimizationDetectionAction()    {}
func (s *VideoOptimizationService) DeleteVideoOptimizationDetectionAction() {}
func (s *VideoOptimizationService) UpdateVideoOptimizationDetectionAction() {}
func (s *VideoOptimizationService) UnsetVideoOptimizationDetectionAction()  {}
func (s *VideoOptimizationService) RenameVideoOptimizationDetectionAction() {}
func (s *VideoOptimizationService) GetAllVideoOptimizationDetectionAction() {}
func (s *VideoOptimizationService) GetVideoOptimizationDetectionAction()    {}
func (s *VideoOptimizationService) CountVideoOptimizationDetectionAction()  {}

// videooptimizationdetectionpolicy
func (s *VideoOptimizationService) AddVideoOptimizationDetectionPolicy()    {}
func (s *VideoOptimizationService) DeleteVideoOptimizationDetectionPolicy() {}
func (s *VideoOptimizationService) UpdateVideoOptimizationDetectionPolicy() {}
func (s *VideoOptimizationService) UnsetVideoOptimizationDetectionPolicy()  {}
func (s *VideoOptimizationService) RenameVideoOptimizationDetectionPolicy() {}
func (s *VideoOptimizationService) GetAllVideoOptimizationDetectionPolicy() {}
func (s *VideoOptimizationService) GetVideoOptimizationDetectionPolicy()    {}
func (s *VideoOptimizationService) CountVideoOptimizationDetectionPolicy()  {}

// videooptimizationdetectionpolicylabel
func (s *VideoOptimizationService) AddVideoOptimizationDetectionPolicyLabel()    {}
func (s *VideoOptimizationService) DeleteVideoOptimizationDetectionPolicyLabel() {}
func (s *VideoOptimizationService) RenameVideoOptimizationDetectionPolicyLabel() {}
func (s *VideoOptimizationService) GetAllVideoOptimizationDetectionPolicyLabel() {}
func (s *VideoOptimizationService) GetVideoOptimizationDetectionPolicyLabel()    {}
func (s *VideoOptimizationService) CountVideoOptimizationDetectionPolicyLabel()  {}

// videooptimizationdetectionpolicylabel_binding
func (s *VideoOptimizationService) GetAllVideoOptimizationDetectionPolicyLabelBinding() {}
func (s *VideoOptimizationService) GetVideoOptimizationDetectionPolicyLabelBinding()    {}

// videooptimizationdetectionpolicylabel_policybinding_binding
func (s *VideoOptimizationService) GetAllVideoOptimizationDetectionPolicyLabelPolicyBindingBinding() {
}
func (s *VideoOptimizationService) GetVideoOptimizationDetectionPolicyLabelPolicyBindingBinding()   {}
func (s *VideoOptimizationService) CountVideoOptimizationDetectionPolicyLabelPolicyBindingBinding() {}

// videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding
func (s *VideoOptimizationService) AddVideoOptimizationDetectionPolicyLabelVideoOptimizationDetectionPolicyBinding() {
}
func (s *VideoOptimizationService) DeleteVideoOptimizationDetectionPolicyLabelVideoOptimizationDetectionPolicyBinding() {
}
func (s *VideoOptimizationService) GetAllVideoOptimizationDetectionPolicyLabelVideoOptimizationDetectionPolicyBinding() {
}
func (s *VideoOptimizationService) GetVideoOptimizationDetectionPolicyLabelVideoOptimizationDetectionPolicyBinding() {
}
func (s *VideoOptimizationService) CountVideoOptimizationDetectionPolicyLabelVideoOptimizationDetectionPolicyBinding() {
}

// videooptimizationdetectionpolicy_binding
func (s *VideoOptimizationService) GetAllVideoOptimizationDetectionPolicyBinding() {}
func (s *VideoOptimizationService) GetVideoOptimizationDetectionPolicyBinding()    {}

// videooptimizationdetectionpolicy_lbvserver_binding
func (s *VideoOptimizationService) GetAllVideoOptimizationDetectionPolicyLBVServerBinding() {}
func (s *VideoOptimizationService) GetVideoOptimizationDetectionPolicyLBVServerBinding()    {}
func (s *VideoOptimizationService) CountVideoOptimizationDetectionPolicyLBVServerBinding()  {}

// videooptimizationdetectionpolicy_videooptimizationglobaldetection_binding
func (s *VideoOptimizationService) GetAllVideoOptimizationDetectionPolicyVideoOptimizationGlobalDetectionBinding() {
}
func (s *VideoOptimizationService) GetVideoOptimizationDetectionPolicyVideoOptimizationGlobalDetectionBinding() {
}
func (s *VideoOptimizationService) CountVideoOptimizationDetectionPolicyVideoOptimizationGlobalDetectionBinding() {
}

// videooptimizationglobaldetection_binding
func (s *VideoOptimizationService) GetVideoOptimizationGlobalDetectionBinding() {}

// videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding
func (s *VideoOptimizationService) AddVideoOptimizationGlobalDetectionVideoOptimizationDetectionPolicyBinding() {
}
func (s *VideoOptimizationService) DeleteVideoOptimizationGlobalDetectionVideoOptimizationDetectionPolicyBinding() {
}
func (s *VideoOptimizationService) GetVideoOptimizationGlobalDetectionVideoOptimizationDetectionPolicyBinding() {
}
func (s *VideoOptimizationService) CountVideoOptimizationGlobalDetectionVideoOptimizationDetectionPolicyBinding() {
}

// videooptimizationglobalpacing_binding
func (s *VideoOptimizationService) GetVideoOptimizationGlobalPacingBinding() {}

// videooptimizationglobalpacing_videooptimizationpacingpolicy_binding
func (s *VideoOptimizationService) AddVideoOptimizationGlobalPacingVideoOptimizationPacingPolicyBinding() {
}
func (s *VideoOptimizationService) DeleteVideoOptimizationGlobalPacingVideoOptimizationPacingPolicyBinding() {
}
func (s *VideoOptimizationService) GetVideoOptimizationGlobalPacingVideoOptimizationPacingPolicyBinding() {
}
func (s *VideoOptimizationService) CountVideoOptimizationGlobalPacingVideoOptimizationPacingPolicyBinding() {
}

// videooptimizationpacingaction
func (s *VideoOptimizationService) AddVideoOptimizationPacingAction()    {}
func (s *VideoOptimizationService) DeleteVideoOptimizationPacingAction() {}
func (s *VideoOptimizationService) UpdateVideoOptimizationPacingAction() {}
func (s *VideoOptimizationService) UnsetVideoOptimizationPacingAction()  {}
func (s *VideoOptimizationService) RenameVideoOptimizationPacingAction() {}
func (s *VideoOptimizationService) GetAllVideoOptimizationPacingAction() {}
func (s *VideoOptimizationService) GetVideoOptimizationPacingAction()    {}
func (s *VideoOptimizationService) CountVideoOptimizationPacingAction()  {}

// videooptimizationpacingpolicy
func (s *VideoOptimizationService) AddVideoOptimizationPacingPolicy()    {}
func (s *VideoOptimizationService) DeleteVideoOptimizationPacingPolicy() {}
func (s *VideoOptimizationService) UpdateVideoOptimizationPacingPolicy() {}
func (s *VideoOptimizationService) UnsetVideoOptimizationPacingPolicy()  {}
func (s *VideoOptimizationService) RenameVideoOptimizationPacingPolicy() {}
func (s *VideoOptimizationService) GetAllVideoOptimizationPacingPolicy() {}
func (s *VideoOptimizationService) GetVideoOptimizationPacingPolicy()    {}
func (s *VideoOptimizationService) CountVideoOptimizationPacingPolicy()  {}

// videooptimizationpacingpolicylabel
func (s *VideoOptimizationService) AddVideoOptimizationPacingPolicyLabel()    {}
func (s *VideoOptimizationService) DeleteVideoOptimizationPacingPolicyLabel() {}
func (s *VideoOptimizationService) RenameVideoOptimizationPacingPolicyLabel() {}
func (s *VideoOptimizationService) GetAllVideoOptimizationPacingPolicyLabel() {}
func (s *VideoOptimizationService) GetVideoOptimizationPacingPolicyLabel()    {}
func (s *VideoOptimizationService) CountVideoOptimizationPacingPolicyLabel()  {}

// videooptimizationpacingpolicylabel_binding
func (s *VideoOptimizationService) GetAllVideoOptimizationPacingPolicyLabelBinding() {}
func (s *VideoOptimizationService) GetVideoOptimizationPacingPolicyLabelBinding()    {}

// videooptimizationpacingpolicylabel_policybinding_binding
func (s *VideoOptimizationService) GetAllVideoOptimizationPacingPolicyLabelPolicyBindingBinding() {}
func (s *VideoOptimizationService) GetVideoOptimizationPacingPolicyLabelPolicyBindingBinding()    {}
func (s *VideoOptimizationService) CountVideoOptimizationPacingPolicyLabelPolicyBindingBinding()  {}

// videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding
func (s *VideoOptimizationService) AddVideoOptimizationPacingPolicyLabelVideoOptimizationPacingPolicyBinding() {
}
func (s *VideoOptimizationService) DeleteVideoOptimizationPacingPolicyLabelVideoOptimizationPacingPolicyBinding() {
}
func (s *VideoOptimizationService) GetAllVideoOptimizationPacingPolicyLabelVideoOptimizationPacingPolicyBinding() {
}
func (s *VideoOptimizationService) GetVideoOptimizationPacingPolicyLabelVideoOptimizationPacingPolicyBinding() {
}
func (s *VideoOptimizationService) CountVideoOptimizationPacingPolicyLabelVideoOptimizationPacingPolicyBinding() {
}

// videooptimizationpacingpolicy_binding
func (s *VideoOptimizationService) GetAllVideoOptimizationPacingPolicyBinding() {}
func (s *VideoOptimizationService) GetVideoOptimizationPacingPolicyBinding()    {}

// videooptimizationpacingpolicy_lbvserver_binding
func (s *VideoOptimizationService) GetAllVideoOptimizationPacingPolicyLBVServerBinding() {}
func (s *VideoOptimizationService) GetVideoOptimizationPacingPolicyLBVServerBinding()    {}
func (s *VideoOptimizationService) CountVideoOptimizationPacingPolicyLBVServerBinding()  {}

// videooptimizationpacingpolicy_videooptimizationglobalpacing_binding
func (s *VideoOptimizationService) GetAllVideoOptimizationPacingPolicyVideoOptimizationGlobalPacingBinding() {
}
func (s *VideoOptimizationService) GetVideoOptimizationPacingPolicyVideoOptimizationGlobalPacingBinding() {
}
func (s *VideoOptimizationService) CountVideoOptimizationPacingPolicyVideoOptimizationGlobalPacingBinding() {
}

// videooptimizationparameter
func (s *VideoOptimizationService) UpdateVideoOptimizationParameter() {}
func (s *VideoOptimizationService) UnsetVideoOptimizationParameter()  {}
func (s *VideoOptimizationService) GetAllVideoOptimizationParameter() {}
