package nitrogo

type DBService struct {
	client *Client
}

// dbdbprofile
func (s *DBService) AddDBDBProfile()    {}
func (s *DBService) DeleteDBDBProfile() {}
func (s *DBService) UpdateDBDBProfile() {}
func (s *DBService) UnsetDBDBProfile()  {}
func (s *DBService) GetAllDBDBProfile() {}
func (s *DBService) GetDBDBProfile()    {}
func (s *DBService) CountDBDBProfile()  {}

// dbuser
func (s *DBService) AddDBUser()    {}
func (s *DBService) DeleteDBUser() {}
func (s *DBService) UpdateDBUser() {}
func (s *DBService) GetAllDBUser() {}
func (s *DBService) GetDBUser()    {}
func (s *DBService) CountDBUser()  {}
