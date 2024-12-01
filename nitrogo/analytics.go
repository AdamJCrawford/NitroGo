package nitrogo

// Analytics configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/analytics/analytics
type AnalyticsService struct {
	client *Client
}

// analyticsglobal_analyticsprofile_binding
// Binding object showing the analyticsprofile that can be bound to analyticsglobal.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/analytics/analyticsglobal_analyticsprofile_binding
func (s *AnalyticsService) AddAnalyticsGlobalAnalyticsProfile()    {}
func (s *AnalyticsService) DeleteAnalyticsGlobalAnalyticsProfile() {}
func (s *AnalyticsService) GetAnalyticsGlobalAnalyticsProfile()    {}
func (s *AnalyticsService) CountAnalyticsGlobalAnalyticsProfile()  {}

// analyticsglobal_binding
// Binding object which returns the resources bound to analyticsglobal.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/analytics/analyticsglobal_binding
func (s *AnalyticsService) AnalyticsGlobalBinding() {}

// analyticsprofile
// Configuration for Analytics profile resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/analytics/analyticsprofile
func (s *AnalyticsService) AddAnalyticsProfile()    {}
func (s *AnalyticsService) UpdateAnalyticsProfile() {}
func (s *AnalyticsService) UnsetAnalyticsProfile()  {}
func (s *AnalyticsService) DeleteAnalyticsProfile() {}
func (s *AnalyticsService) GetAllAnalyticsProfile() {}
func (s *AnalyticsService) GetAnalyticsProfile()    {}
func (s *AnalyticsService) CheckAnalyticsProfile()  {}
func (s *AnalyticsService) ChangeAnalyticsProfile() {}
