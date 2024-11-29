package nitrogo

type UserService struct {
	client *Client
}

// userprotocol
func (s *UserService) AddUserProtocol()    {}
func (s *UserService) DeleteUserProtocol() {}
func (s *UserService) UpdateUserProtocol() {}
func (s *UserService) UnsetUserProtocol()  {}
func (s *UserService) GetAllUserProtocol() {}
func (s *UserService) GetUserProtocol()    {}
func (s *UserService) CountUserProtocol()  {}

// uservserver
func (s *UserService) AddUserVServer()    {}
func (s *UserService) DeleteUserVServer() {}
func (s *UserService) UpdateUserVServer() {}
func (s *UserService) UnsetUserVServer()  {}
func (s *UserService) EnableUserVServer() {}
func (s *UserService) GetAllUserVServer() {}
func (s *UserService) GetUserVServer()    {}
func (s *UserService) CountUserVServer()  {}
