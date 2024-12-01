package nitrogo

// NTP configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/ntp/ntp
type NTPService struct {
	client *Client
}

// ntpparam
// Configuration for NTP parameter resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/ntp/ntpparam
func (s *NTPService) UpdateNTPParam() {}
func (s *NTPService) UnsetNTPParam()  {}
func (s *NTPService) GetAllNTPParam() {}

// ntpserver
// Configuration for NTP server resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/ntp/ntpserver
func (s *NTPService) AddNTPServer()    {}
func (s *NTPService) DeleteNTPServer() {}
func (s *NTPService) UpdateNTPServer() {}
func (s *NTPService) UnsetNTPServer()  {}
func (s *NTPService) GetAllNTPServer() {}
func (s *NTPService) CountNTPServer()  {}

// ntpstatus
// Configuration for ntp status resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/ntp/ntpstatus
func (s *NTPService) GetAllNTPStatus() {}

// ntpsync
// Configuration for NTP sync resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/ntp/ntpsync
func (s *NTPService) EnableNTPSync()  {}
func (s *NTPService) DisableNTPSync() {}
func (s *NTPService) GetAllNTPSync()  {}
