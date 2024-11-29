package nitrogo

type LLDPService struct {
	client *Client
}

// lldpneighbors
func (s *LLDPService) GetAllLLDPNeighbors() {}
func (s *LLDPService) GetLLDPNeighbors()    {}
func (s *LLDPService) CountLLDPNeighbors()  {}
func (s *LLDPService) ClearLLDPNeighbors()  {}

// lldpparam
func (s *LLDPService) UpdateLLDPParam() {}
func (s *LLDPService) UnsetLLDPParam()  {}
func (s *LLDPService) GetAllLLDPParam() {}
