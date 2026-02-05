package nitrogo

const (
	dbDBProfileURL = "/nitro/v1/config/dbdbprofile"
	dbUserURL      = "/nitro/v1/config/dbuser"
)

// All the commands associated with database user
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/db/db
type DBService struct {
	client *Client
}

// dbdbprofile
// Configuration for DB profile resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/db/dbdbprofile
func (s *DBService) AddDBDBProfile()    {}
func (s *DBService) DeleteDBDBProfile() {}
func (s *DBService) UpdateDBDBProfile() {}
func (s *DBService) UnsetDBDBProfile()  {}
func (s *DBService) GetAllDBDBProfile() {}
func (s *DBService) GetDBDBProfile()    {}
func (s *DBService) CountDBDBProfile()  {}

// dbuser
// Configuration for DB user resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/db/dbuser
func (s *DBService) AddDBUser()    {}
func (s *DBService) DeleteDBUser() {}
func (s *DBService) UpdateDBUser() {}
func (s *DBService) GetAllDBUser() {}
func (s *DBService) GetDBUser()    {}
func (s *DBService) CountDBUser()  {}
