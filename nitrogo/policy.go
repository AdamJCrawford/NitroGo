package nitrogo

const (
	policyDatasetURL                 = "/nitro/v1/config/policydataset"
	policyDatasetBindingURL          = "/nitro/v1/config/policydataset_binding"
	policyDatasetValueBindingURL     = "/nitro/v1/config/policydataset_value_binding"
	policyEvaluationURL              = "/nitro/v1/config/policyevaluation"
	policyExpressionURL              = "/nitro/v1/config/policyexpression"
	policyHTTPCalloutURL             = "/nitro/v1/config/policyhttpcallout"
	policyMapURL                     = "/nitro/v1/config/policymap"
	policyParamURL                   = "/nitro/v1/config/policyparam"
	policyPATSetURL                  = "/nitro/v1/config/policypatset"
	policyPATSetBindingURL           = "/nitro/v1/config/policypatset_binding"
	policyPATSetPatternBindingURL    = "/nitro/v1/config/policypatset_pattern_binding"
	policyStringMapURL               = "/nitro/v1/config/policystringmap"
	policyStringMapBindingURL        = "/nitro/v1/config/policystringmap_binding"
	policyStringMapPatternBindingURL = "/nitro/v1/config/policystringmap_pattern_binding"
	policyURLSetURL                  = "/nitro/v1/config/policyurlset"
)

// Policy configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/policy/policy
type PolicyService struct {
	client *Client
}

// policydataset
func (s *PolicyService) AddPolicyDataset()    {}
func (s *PolicyService) DeletePolicyDataset() {}
func (s *PolicyService) UpdatePolicyDataset() {}
func (s *PolicyService) UnsetPolicyDataset()  {}
func (s *PolicyService) GetAllPolicyDataset() {}
func (s *PolicyService) GetPolicyDataset()    {}
func (s *PolicyService) CountPolicyDataset()  {}

// policydataset_binding
func (s *PolicyService) GetAllPolicyDatasetBinding() {}
func (s *PolicyService) GetPolicyDatasetBinding()    {}

// policydataset_value_binding
func (s *PolicyService) AddPolicyDatasetValueBinding()    {}
func (s *PolicyService) DeletePolicyDatasetValueBinding() {}
func (s *PolicyService) GetAllPolicyDatasetValueBinding() {}
func (s *PolicyService) GetPolicyDatasetValueBinding()    {}
func (s *PolicyService) CountPolicyDatasetValueBinding()  {}

// policyevaluation
func (s *PolicyService) GetAllPolicyEvaluation() {}
func (s *PolicyService) CountPolicyEvaluation()  {}

// policyexpression
func (s *PolicyService) AddPolicyExpression()    {}
func (s *PolicyService) DeletePolicyExpression() {}
func (s *PolicyService) UpdatePolicyExpression() {}
func (s *PolicyService) UnsetPolicyExpression()  {}
func (s *PolicyService) GetAllPolicyExpression() {}
func (s *PolicyService) GetPolicyExpression()    {}
func (s *PolicyService) CountPolicyExpression()  {}

// policyhttpcallout
func (s *PolicyService) AddPolicyHTTPCallout()    {}
func (s *PolicyService) DeletePolicyHTTPCallout() {}
func (s *PolicyService) UpdatePolicyHTTPCallout() {}
func (s *PolicyService) UnsetPolicyHTTPCallout()  {}
func (s *PolicyService) GetAllPolicyHTTPCallout() {}
func (s *PolicyService) GetPolicyHTTPCallout()    {}
func (s *PolicyService) CountPolicyHTTPCallout()  {}

// policymap
func (s *PolicyService) AddPolicyMap()    {}
func (s *PolicyService) DeletePolicyMap() {}
func (s *PolicyService) GetAllPolicyMap() {}
func (s *PolicyService) GetPolicyMap()    {}
func (s *PolicyService) CountPolicyMap()  {}

// policyparam
func (s *PolicyService) UpdatePolicyParam() {}
func (s *PolicyService) UnsetPolicyParam()  {}
func (s *PolicyService) GetAllPolicyParam() {}

// policypatset
func (s *PolicyService) AddPolicyPATSet()    {}
func (s *PolicyService) DeletePolicyPATSet() {}
func (s *PolicyService) UpdatePolicyPATSet() {}
func (s *PolicyService) UnsetPolicyPATSet()  {}
func (s *PolicyService) GetAllPolicyPATSet() {}
func (s *PolicyService) GetPolicyPATSet()    {}
func (s *PolicyService) CountPolicyPATSet()  {}

// policypatset_binding
func (s *PolicyService) GetAllPolicyPATSetBinding() {}
func (s *PolicyService) GetPolicyPATSetBinding()    {}

// policypatset_pattern_binding
func (s *PolicyService) AddPolicyPATSetPatternBinding()    {}
func (s *PolicyService) DeletePolicyPATSetPatternBinding() {}
func (s *PolicyService) GetAllPolicyPATSetPatternBinding() {}
func (s *PolicyService) GetPolicyPATSetPatternBinding()    {}
func (s *PolicyService) CountPolicyPATSetPatternBinding()  {}

// policystringmap
func (s *PolicyService) AddPolicyStringMap()    {}
func (s *PolicyService) DeletePolicyStringMap() {}
func (s *PolicyService) UpdatePolicyStringMap() {}
func (s *PolicyService) UnsetPolicyStringMap()  {}
func (s *PolicyService) GetAllPolicyStringMap() {}
func (s *PolicyService) GetPolicyStringMap()    {}
func (s *PolicyService) CountPolicyStringMap()  {}

// policystringmap_binding
func (s *PolicyService) GetAllPolicyStringMapBinding() {}
func (s *PolicyService) GetPolicyStringMapBinding()    {}

// policystringmap_pattern_binding
func (s *PolicyService) AddPolicyStringMapPatternBinding()    {}
func (s *PolicyService) DeletePolicyStringMapPatternBinding() {}
func (s *PolicyService) GetAllPolicyStringMapPatternBinding() {}
func (s *PolicyService) GetPolicyStringMapPatternBinding()    {}
func (s *PolicyService) CountPolicyStringMapPatternBinding()  {}

// policyurlset
func (s *PolicyService) AddPolicyURLSet()    {}
func (s *PolicyService) DeletePolicyURLSet() {}
func (s *PolicyService) GetAllPolicyURLSet() {}
func (s *PolicyService) GetPolicyURLSet()    {}
func (s *PolicyService) CountPolicyURLSet()  {}
func (s *PolicyService) ImportPolicyURLSet() {}
func (s *PolicyService) ExportPolicyURLSet() {}
func (s *PolicyService) ChangePolicyURLSet() {}
