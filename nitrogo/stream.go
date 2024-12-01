package nitrogo

// Stream configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/stream/stream
type StreamService struct {
	client *Client
}

// streamidentifier
// Configuration for identifier resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/stream/streamidentifier
func (s *StreamService) AddStreamIdentifier()    {}
func (s *StreamService) DeleteStreamIdentifier() {}
func (s *StreamService) UpdateStreamIdentifier() {}
func (s *StreamService) UnsetStreamIdentifier()  {}
func (s *StreamService) GetAllStreamIdentifier() {}
func (s *StreamService) GetStreamIdentifier()    {}
func (s *StreamService) CountStreamIdentifier()  {}

// streamidentifier_binding
// Binding object which returns the resources bound to streamidentifier.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/stream/streamidentifier_binding
func (s *StreamService) GetAllStreamIdentifierBinding() {}
func (s *StreamService) GetStreamIdentifierBinding()    {}

// streamidentifier_streamsession_binding
// Binding object showing the streamsession that can be bound to streamidentifier.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/stream/streamidentifier_streamsession_binding
func (s *StreamService) GetAllStreamIdentifierStreamSessionBinding() {}
func (s *StreamService) GetStreamIdentifierStreamSessionBinding()    {}
func (s *StreamService) CountStreamIdentifierStreamSessionBinding()  {}

// streamselector
// Configuration for selector resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/stream/streamselector
func (s *StreamService) AddStreamSelector()    {}
func (s *StreamService) DeleteStreamSelector() {}
func (s *StreamService) UpdateStreamSelector() {}
func (s *StreamService) GetAllStreamSelector() {}
func (s *StreamService) GetStreamSelector()    {}
func (s *StreamService) CountStreamSelector()  {}

// streamsession
// Configuration for active connection resource.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/stream/streamsession
func (s *StreamService) ClearStreamSession() {}
