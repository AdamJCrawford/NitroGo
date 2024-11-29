package nitrogo

type AnalyticsService struct {
	client *Client
}

// analyticsglobal_analyticsprofile_binding
func (s *AnalyticsService) AddAnalyticsGlobalAnalyticsProfile()    {}
func (s *AnalyticsService) DeleteAnalyticsGlobalAnalyticsProfile() {}
func (s *AnalyticsService) GetAnalyticsGlobalAnalyticsProfile()    {}
func (s *AnalyticsService) CountAnalyticsGlobalAnalyticsProfile()  {}

// analyticsglobal_binding
func (s *AnalyticsService) AnalyticsGlobalBinding() {}

// analyticsprofile
func (s *AnalyticsService) AddAnalyticsProfile()    {}
func (s *AnalyticsService) UpdateAnalyticsProfile() {}
func (s *AnalyticsService) UnsetAnalyticsProfile()  {}
func (s *AnalyticsService) DeleteAnalyticsProfile() {}
func (s *AnalyticsService) GetAllAnalyticsProfile() {}
func (s *AnalyticsService) GetAnalyticsProfile()    {}
func (s *AnalyticsService) CheckAnalyticsProfile()  {}
func (s *AnalyticsService) ChangeAnalyticsProfile() {}
