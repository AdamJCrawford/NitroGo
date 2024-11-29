package nitrogo

type RDPService struct {
	client *Client
}

// rdpclientprofile
func (s *RDPService) AddRDPClientProfile()    {}
func (s *RDPService) DeleteRDPClientProfile() {}
func (s *RDPService) UpdateRDPClientProfile() {}
func (s *RDPService) UnsetRDPClientProfile()  {}
func (s *RDPService) GetAllRDPClientProfile() {}
func (s *RDPService) GetRDPClientProfile()    {}
func (s *RDPService) CountRDPClientProfile()  {}

// rdpconnections
func (s *RDPService) GetAllRDPConnections() {}
func (s *RDPService) CountRDPConnections()  {}
func (s *RDPService) KillRDPConnections()   {}

// rspserverprofile
func (s *RDPService) AddRDPServerProfile()    {}
func (s *RDPService) DeleteRDPServerProfile() {}
func (s *RDPService) UpdateRDPServerProfile() {}
func (s *RDPService) UnsetRDPServerProfile()  {}
func (s *RDPService) GetAllRDPServerProfile() {}
func (s *RDPService) GetRDPServerProfile()    {}
func (s *RDPService) CountRDPServerProfile()  {}
