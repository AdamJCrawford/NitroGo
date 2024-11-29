package nitrogo

type SMPPService struct {
	client *Client
}

// smppparam
func (s *SMPPService) UpdateSMPPParam() {}
func (s *SMPPService) UnsetSMPPParam()  {}
func (s *SMPPService) GetAllSMPPParam() {}

// smppuser
func (s *SMPPService) AddSMPPUser()    {}
func (s *SMPPService) DeleteSMPPUser() {}
func (s *SMPPService) UpdateSMPPUser() {}
func (s *SMPPService) GetAllSMPPUser() {}
func (s *SMPPService) GetSMPPUser()    {}
func (s *SMPPService) CountSMPPUser()  {}
