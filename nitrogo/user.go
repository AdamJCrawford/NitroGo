package nitrogo

// User protocol configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/user/user
type UserService struct {
	client *Client
}

// userprotocol
// Configuration for user protocol resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/user/userprotocol
func (s *UserService) AddUserProtocol()    {}
func (s *UserService) DeleteUserProtocol() {}
func (s *UserService) UpdateUserProtocol() {}
func (s *UserService) UnsetUserProtocol()  {}
func (s *UserService) GetAllUserProtocol() {}
func (s *UserService) GetUserProtocol()    {}
func (s *UserService) CountUserProtocol()  {}

// uservserver
// Configuration for virtual server resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/user/uservserver
func (s *UserService) AddUserVServer()    {}
func (s *UserService) DeleteUserVServer() {}
func (s *UserService) UpdateUserVServer() {}
func (s *UserService) UnsetUserVServer()  {}
func (s *UserService) EnableUserVServer() {}
func (s *UserService) GetAllUserVServer() {}
func (s *UserService) GetUserVServer()    {}
func (s *UserService) CountUserVServer()  {}
