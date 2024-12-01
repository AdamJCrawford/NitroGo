package nitrogo

// Citrix ADC as SD proxy Configuration and cloud discovery commands.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/cloud/cloud
type CloudService struct {
	client *Client
}

// cloudautoscalegroup
// Configuration for Allowed ticket profile for NGS resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/cloud/cloudallowedngsticketprofile
func (s *CloudService) GetAllCloudAutoscaleGroup() {}
func (s *CloudService) GetCloudAutoscaleGroup()    {}
func (s *CloudService) CountCloudAutoscaleGroup()  {}

// cloudcredential
// Configuration for cloud credentials resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/cloud/cloudcredential
func (s *CloudService) GetAllCloudCredential() {}
func (s *CloudService) UpdateCloudCredential() {}

// cloudparameter
// Configuration for cloud parameter resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/cloud/cloudparameter
func (s *CloudService) UpdateCloudParameter() {}
func (s *CloudService) UnsetCloudParameter()  {}
func (s *CloudService) GetAllCloudParameter() {}

// cloudparaminternal
// Configuration for cloud paramInternal resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/cloud/cloudparaminternal
func (s *CloudService) GetAllCloudParamInternal() {}
func (s *CloudService) CountCloudParamInternal()  {}
func (s *CloudService) UpdateCloudParamInternal() {}

// cloudprofile
// Configuration for cloud profile resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/cloud/cloudprofile
func (s *CloudService) AddCloudProfile()    {}
func (s *CloudService) DeleteCloudProfile() {}
func (s *CloudService) GetAllCloudProfile() {}
func (s *CloudService) GetCloudProfile()    {}
func (s *CloudService) CountCloudProfile()  {}

// cloudservice
// Configuration for cloud service resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/cloud/cloudservice
func (s *CloudService) CheckCloudService() {}

// cloudvserverip
// Configuration for Cloud virtual server IPs resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/cloud/cloudvserverip
func (s *CloudService) GetAllCloudVServerIP() {}
func (s *CloudService) CountCloudVServerIP()  {}
