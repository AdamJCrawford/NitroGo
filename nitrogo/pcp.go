package nitrogo

type PCPService struct {
	client *Client
}

// pcpmap
func (s *PCPService) GetAllPCPMap() {}
func (s *PCPService) CountPCPMap()  {}

// pcpprofile
func (s *PCPService) AddPCPProfile()    {}
func (s *PCPService) DeletePCPProfile() {}
func (s *PCPService) UpdatePCPProfile() {}
func (s *PCPService) UnsetPCPProfile()  {}
func (s *PCPService) GetAllPCPProfile() {}
func (s *PCPService) GetPCPProfile()    {}
func (s *PCPService) CountPCPProfile()  {}

// pcpserver
func (s *PCPService) AddPCPServer()    {}
func (s *PCPService) DeletePCPServer() {}
func (s *PCPService) UpdatePCPServer() {}
func (s *PCPService) UnsetPCPServer()  {}
func (s *PCPService) GetAllPCPServer() {}
func (s *PCPService) GetPCPServer()    {}
func (s *PCPService) CountPCPServer()  {}
