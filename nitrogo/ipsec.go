package nitrogo

type IPSECService struct {
	client *Client
}

// ipsecparameter
func (s *IPSECService) UpdateIPSECParamter() {}
func (s *IPSECService) UnsetIPSECParamter()  {}
func (s *IPSECService) GetAllIPSECParamter() {}

// ipsecprofile
func (s *IPSECService) AddIPSECProfile()    {}
func (s *IPSECService) DeleteIPSECProfile() {}
func (s *IPSECService) GetAllIPSECProfile() {}
func (s *IPSECService) GetIPSECProfile()    {}
func (s *IPSECService) CountIPSECProfile()  {}
