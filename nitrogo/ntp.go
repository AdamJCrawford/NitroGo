package nitrogo

type NTPService struct {
	client *Client
}

// ntpparam
func (s *NTPService) UpdateNTPParam() {}
func (s *NTPService) UnsetNTPParam()  {}
func (s *NTPService) GetAllNTPParam() {}

// ntpserver
func (s *NTPService) AddNTPServer()    {}
func (s *NTPService) DeleteNTPServer() {}
func (s *NTPService) UpdateNTPServer() {}
func (s *NTPService) UnsetNTPServer()  {}
func (s *NTPService) GetAllNTPServer() {}
func (s *NTPService) CountNTPServer()  {}

// ntpstatus
func (s *NTPService) GetAllNTPStatus() {}

// ntpsync
func (s *NTPService) EnableNTPSync()  {}
func (s *NTPService) DisableNTPSync() {}
func (s *NTPService) GetAllNTPSync()  {}
