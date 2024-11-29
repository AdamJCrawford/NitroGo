package nitrogo

type TunnelService struct {
	client *Client
}

// tunnelglobal_binding
func (s *TunnelService) GetTunnelGlobalBinding() {}

// tunnelglobal_tunneltrafficpolicy_binding
func (s *TunnelService) AddTunnelGlobalTunnelTrafficPolicyBinding()    {}
func (s *TunnelService) DeleteTunnelGlobalTunnelTrafficPolicyBinding() {}
func (s *TunnelService) GetTunnelGlobalTunnelTrafficPolicyBinding()    {}
func (s *TunnelService) CountTunnelGlobalTunnelTrafficPolicyBinding()  {}

// tunneltrafficpolicy
func (s *TunnelService) AddTunnelTrafficPolicy()    {}
func (s *TunnelService) DeleteTunnelTrafficPolicy() {}
func (s *TunnelService) UpdateTunnelTrafficPolicy() {}
func (s *TunnelService) UnsetTunnelTrafficPolicy()  {}
func (s *TunnelService) GetAllTunnelTrafficPolicy() {}
func (s *TunnelService) GetTunnelTrafficPolicy()    {}
func (s *TunnelService) CountTunnelTrafficPolicy()  {}
func (s *TunnelService) RenameTunnelTrafficPolicy() {}

// tunneltrafficpolicy_binding
func (s *TunnelService) GetAllTunnelTrafficPolicyBinding() {}
func (s *TunnelService) GetTunnelTrafficPolicyBinding()    {}

// tunneltrafficpolicy_tunnelglobal_binding
func (s *TunnelService) GetAllTunnelTrafficPolicyTunnelGlobalBinding() {}
func (s *TunnelService) GetTunnelTrafficPolicyTunnelGlobalBinding()    {}
func (s *TunnelService) CountTunnelTrafficPolicyTunnelGlobalBinding()  {}
