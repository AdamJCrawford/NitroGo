package nitrogo

type IPSECALGService struct {
	client *Client
}

// ipsecalgprofile
func (s *IPSECALGService) AddIPSECALGProfile()    {}
func (s *IPSECALGService) DeleteIPSECALGProfile() {}
func (s *IPSECALGService) UpdateIPSECALGProfile() {}
func (s *IPSECALGService) UnsetIPSECALGProfile()  {}
func (s *IPSECALGService) GetAllIPSECALGProfile() {}
func (s *IPSECALGService) GetIPSECALGProfile()    {}
func (s *IPSECALGService) CountIPSECALGProfile()  {}

// ipsecalgsession
func (s *IPSECALGService) GetAllIPSECALGSession() {}
func (s *IPSECALGService) CountIPSECALGSession()  {}
func (s *IPSECALGService) FlushIPSECALGSession()  {}
