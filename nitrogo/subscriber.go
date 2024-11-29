package nitrogo

type SubscriberService struct {
	client *Client
}

// subscribergxinterface
func (s *SubscriberService) UpdateSubscriberGxInterface() {}
func (s *SubscriberService) UnsetSubscriberGxInterface()  {}
func (s *SubscriberService) GetAllSubscriberGxInterface() {}

// subscriberparam
func (s *SubscriberService) UpdateSubscriberParam() {}
func (s *SubscriberService) UnsetSubscriberParam()  {}
func (s *SubscriberService) GetAllSubscriberParam() {}

// subscriberprofile
func (s *SubscriberService) AddSubscriberProfile()    {}
func (s *SubscriberService) DeleteSubscriberProfile() {}
func (s *SubscriberService) UpdateSubscriberProfile() {}
func (s *SubscriberService) UnsetSubscriberProfile()  {}
func (s *SubscriberService) GetAllSubscriberProfile() {}
func (s *SubscriberService) CountSubscriberProfile()  {}

// subscriverradiusinterface
func (s *SubscriberService) UpdateSubscriberRADIUSInterface() {}
func (s *SubscriberService) UnsetSubscriberRADIUSInterface()  {}
func (s *SubscriberService) GetAllSubscriberRADIUSInterface() {}

// subscribersessions
func (s *SubscriberService) ClearSubscriberSessions()  {}
func (s *SubscriberService) GetAllSubscriberSessions() {}
func (s *SubscriberService) CountSubscriberSessions()  {}
