package nitrogo

type CloudService struct {
	client *Client
}

// cloudautoscalegroup
func (s *CloudService) GetAllCloudAutoscaleGroup() {}
func (s *CloudService) GetCloudAutoscaleGroup()    {}
func (s *CloudService) CountCloudAutoscaleGroup()  {}

// cloudcredential
func (s *CloudService) GetAllCloudCredential() {}
func (s *CloudService) UpdateCloudCredential() {}

// cloudparameter
func (s *CloudService) UpdateCloudParameter() {}
func (s *CloudService) UnsetCloudParameter()  {}
func (s *CloudService) GetAllCloudParameter() {}

// cloudparaminternal
func (s *CloudService) GetAllCloudParamInternal() {}
func (s *CloudService) CountCloudParamInternal()  {}
func (s *CloudService) UpdateCloudParamInternal() {}

// cloudprofile
func (s *CloudService) AddCloudProfile()    {}
func (s *CloudService) DeleteCloudProfile() {}
func (s *CloudService) GetAllCloudProfile() {}
func (s *CloudService) GetCloudProfile()    {}
func (s *CloudService) CountCloudProfile()  {}

// cloudservice
func (s *CloudService) CheckCloudService() {}

// cloudvserverip
func (s *CloudService) GetAllCloudVServerIP() {}
func (s *CloudService) CountCloudVServerIP()  {}
