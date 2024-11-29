package nitrogo

type StreamService struct {
	client *Client
}

// streamidentifier
func (s *StreamService) AddStreamIdentifier()    {}
func (s *StreamService) DeleteStreamIdentifier() {}
func (s *StreamService) UpdateStreamIdentifier() {}
func (s *StreamService) UnsetStreamIdentifier()  {}
func (s *StreamService) GetAllStreamIdentifier() {}
func (s *StreamService) GetStreamIdentifier()    {}
func (s *StreamService) CountStreamIdentifier()  {}

// streamidentifier_binding
func (s *StreamService) GetAllStreamIdentifierBinding() {}
func (s *StreamService) GetStreamIdentifierBinding()    {}

// streamidentifier_streamsession_binding
func (s *StreamService) GetAllStreamIdentifierStreamSessionBinding() {}
func (s *StreamService) GetStreamIdentifierStreamSessionBinding()    {}
func (s *StreamService) CountStreamIdentifierStreamSessionBinding()  {}

// streamselector
func (s *StreamService) AddStreamSelector()    {}
func (s *StreamService) DeleteStreamSelector() {}
func (s *StreamService) UpdateStreamSelector() {}
func (s *StreamService) GetAllStreamSelector() {}
func (s *StreamService) GetStreamSelector()    {}
func (s *StreamService) CountStreamSelector()  {}

// streamsession
func (s *StreamService) ClearStreamSession() {}
