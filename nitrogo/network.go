package nitrogo

type NetworkService struct {
	client *Client
}

// appalgparam
func (s *NetworkService) UpdateAppAlgParam() {}
func (s *NetworkService) UnsetAppAlgParam()  {}
func (s *NetworkService) GetAllAppAlgParam() {}

// arp
func (s *NetworkService) AddARP()    {}
func (s *NetworkService) DeleteARP() {}
func (s *NetworkService) SendARP()   {}
func (s *NetworkService) GetAllARP() {}
func (s *NetworkService) CountARP()  {}

// arpparam
func (s *NetworkService) UpdateARPParam() {}
func (s *NetworkService) UnsetARPParam()  {}
func (s *NetworkService) GetAllARPParam() {}

// bridgegroup
func (s *NetworkService) AddBridgeGroup()    {}
func (s *NetworkService) DeleteBridgeGroup() {}
func (s *NetworkService) UpdateBridgeGroup() {}
func (s *NetworkService) UnsetBridgeGroup()  {}
func (s *NetworkService) GetAllBridgeGroup() {}
func (s *NetworkService) GetBridgeGroup()    {}
func (s *NetworkService) CountBridgeGroup()  {}

// bridgegroup_binding
func (s *NetworkService) GetAllBridgeGroupBinding() {}
func (s *NetworkService) GetBridgeGroupBinding()    {}

// bridgegroup_nsip6_binding
func (s *NetworkService) AddBridgeGroupNSIP6Binding()    {}
func (s *NetworkService) DeleteBridgeGroupNSIP6Binding() {}
func (s *NetworkService) GetAllBridgeGroupNSIP6Binding() {}
func (s *NetworkService) GetBridgeGroupNSIP6Binding()    {}
func (s *NetworkService) CountBridgeGroupNSIP6Binding()  {}

// bridgegroup_nsip_binding
func (s *NetworkService) AddBrigeGroupNSIPBinding()    {}
func (s *NetworkService) DeleteBrigeGroupNSIPBinding() {}
func (s *NetworkService) GetAllBrigeGroupNSIPBinding() {}
func (s *NetworkService) GetBrigeGroupNSIPBinding()    {}
func (s *NetworkService) CountBrigeGroupNSIPBinding()  {}

// bridgegroup_vlan_binding
func (s *NetworkService) AddBridgeGroupVLANBinding()    {}
func (s *NetworkService) DeleteBridgeGroupVLANBinding() {}
func (s *NetworkService) GetAllBridgeGroupVLANBinding() {}
func (s *NetworkService) GetBridgeGroupVLANBinding()    {}
func (s *NetworkService) CountBridgeGroupVLANBinding()  {}

// bridgetable
func (s *NetworkService) AddBridgeTable()    {}
func (s *NetworkService) DeleteBridgeTable() {}
func (s *NetworkService) UpdateBridgeTable() {}
func (s *NetworkService) UnsetBridgeTable()  {}
func (s *NetworkService) GetAllBridgeTable() {}
func (s *NetworkService) CountBridgeTable()  {}

// channel
func (s *NetworkService) AddChannel()    {}
func (s *NetworkService) DeleteChannel() {}
func (s *NetworkService) UpdateChannel() {}
func (s *NetworkService) UnsetChannel()  {}
func (s *NetworkService) GetAllChannel() {}
func (s *NetworkService) GetChannel()    {}
func (s *NetworkService) CountChannel()  {}

// channel_binding
func (s *NetworkService) GetAllChannelBinding() {}
func (s *NetworkService) GetChannelBinding()    {}

// channel_interface_binding
func (s *NetworkService) AddChannelInterfaceBinding()    {}
func (s *NetworkService) DeleteChannelInterfaceBinding() {}
func (s *NetworkService) GetAllChannelInterfaceBinding() {}
func (s *NetworkService) GetChannelInterfaceBinding()    {}
func (s *NetworkService) CountChannelInterfaceBinding()  {}

// ci
func (s *NetworkService) GetAllCI() {}
func (s *NetworkService) CountCI()  {}

// fis
func (s *NetworkService) AddFIS()    {}
func (s *NetworkService) DeleteFIS() {}
func (s *NetworkService) GetAllFIS() {}
func (s *NetworkService) GetFIS()    {}
func (s *NetworkService) CountFIS()  {}

// fis_binding
func (s *NetworkService) GetAllFISBinding() {}
func (s *NetworkService) GetFISBinding()    {}

// fis_channel_binding
func (s *NetworkService) AddFISChannelBinding()    {}
func (s *NetworkService) DeleteFISChannelBinding() {}
func (s *NetworkService) GetAllFISChannelBinding() {}
func (s *NetworkService) GetFISChannelBinding()    {}
func (s *NetworkService) CountFISChannelBinding()  {}

// fis_interface_binding
func (s *NetworkService) AddFISInterfaceBinding()    {}
func (s *NetworkService) DeleteFISInterfaceBinding() {}

// forwardingsession
func (s *NetworkService) AddFowardingSession()    {}
func (s *NetworkService) DeleteFowardingSession() {}
func (s *NetworkService) UpdateFowardingSession() {}
func (s *NetworkService) GetAllFowardingSession() {}
func (s *NetworkService) GetFowardingSession()    {}
func (s *NetworkService) CountFowardingSession()  {}

// inat
func (s *NetworkService) AddINAT()    {}
func (s *NetworkService) DeleteINAT() {}
func (s *NetworkService) UpdateINAT() {}
func (s *NetworkService) UnsetINAT()  {}
func (s *NetworkService) GetAllINAT() {}
func (s *NetworkService) GetINAT()    {}
func (s *NetworkService) CountINAT()  {}

// inatparam
func (s *NetworkService) UpdateINATParam() {}
func (s *NetworkService) UnsetINATParam()  {}
func (s *NetworkService) GetAllINATParam() {}
func (s *NetworkService) GetINATParam()    {}
func (s *NetworkService) CountINATParam()  {}

// interface
func (s *NetworkService) ClearInterface()   {}
func (s *NetworkService) UpdateInterface()  {}
func (s *NetworkService) UnsetInterface()   {}
func (s *NetworkService) EnableInterface()  {}
func (s *NetworkService) DisableInterface() {}
func (s *NetworkService) ResetInterface()   {}
func (s *NetworkService) GetAllInterface()  {}
func (s *NetworkService) GetInterface()     {}
func (s *NetworkService) CountInterface()   {}

// interfacepair
func (s *NetworkService) AddInterfacePair()    {}
func (s *NetworkService) DeleteInterfacePair() {}
func (s *NetworkService) GetAllInterfacePair() {}
func (s *NetworkService) GetInterfacePair()    {}
func (s *NetworkService) CountInterfacePair()  {}

// ip6tunnel
func (s *NetworkService) AddIP6Tunnel()    {}
func (s *NetworkService) DeleteIP6Tunnel() {}
func (s *NetworkService) GetAllIP6Tunnel() {}
func (s *NetworkService) GetIP6Tunnel()    {}
func (s *NetworkService) CountIP6Tunnel()  {}

// ip6tunnelparam
func (s *NetworkService) UpdateIP6TunnelParam() {}
func (s *NetworkService) UnsetIP6TunnelParam()  {}
func (s *NetworkService) GetAllIP6TunnelParam() {}

// ipset
func (s *NetworkService) AddIPSet()    {}
func (s *NetworkService) DeleteIPSet() {}
func (s *NetworkService) GetAllIPSet() {}
func (s *NetworkService) GetIPSet()    {}
func (s *NetworkService) CountIPSet()  {}

// ipset_binding
func (s *NetworkService) GetAllIPSetBinding() {}
func (s *NetworkService) GetIPSetBinding()    {}

// ipset_nsip6_binding
func (s *NetworkService) AddIPSetNSIP6Binding()    {}
func (s *NetworkService) DeleteIPSetNSIP6Binding() {}
func (s *NetworkService) GetAllIPSetNSIP6Binding() {}
func (s *NetworkService) GetIPSetNSIP6Binding()    {}
func (s *NetworkService) CountIPSetNSIP6Binding()  {}

// ipset_nsip_binding
func (s *NetworkService) AddIPSetNSIPBinding()    {}
func (s *NetworkService) DeleteIPSetNSIPBinding() {}
func (s *NetworkService) GetAllIPSetNSIPBinding() {}
func (s *NetworkService) GetIPSetNSIPBinding()    {}
func (s *NetworkService) CountIPSetNSIPBinding()  {}

// iptunnel
func (s *NetworkService) AddIPTunnel()    {}
func (s *NetworkService) DeleteIPTunnel() {}
func (s *NetworkService) UpdateIPTunnel() {}
func (s *NetworkService) UnsetIPTunnel()  {}
func (s *NetworkService) GetAllIPTunnel() {}
func (s *NetworkService) GetIPTunnel()    {}
func (s *NetworkService) CountIPTunnel()  {}

// iptunnelparam
func (s *NetworkService) UpdateIPTunnelParam() {}
func (s *NetworkService) UnsetIPTunnelParam()  {}
func (s *NetworkService) GetAllIPTunnelParam() {}

// ipv6
func (s *NetworkService) UpdateIPV6() {}
func (s *NetworkService) UnsetIPV6()  {}
func (s *NetworkService) GetAllIPV6() {}
func (s *NetworkService) GetIPV6()    {}
func (s *NetworkService) CountIPV6()  {}

// l2param
func (s *NetworkService) UpdateL2Param() {}
func (s *NetworkService) UnsetL2Param()  {}
func (s *NetworkService) GetAllL2Param() {}

// l3param
func (s *NetworkService) UpdateL3Param() {}
func (s *NetworkService) UnsetL3Param()  {}
func (s *NetworkService) GetAllL3Param() {}

// l4param
func (s *NetworkService) UpdateL4Param() {}
func (s *NetworkService) UnsetL4Param()  {}
func (s *NetworkService) GetAllL4Param() {}

// lacp
func (s *NetworkService) UpdateLACP() {}
func (s *NetworkService) GetAllLACP() {}
func (s *NetworkService) GetLACP()    {}
func (s *NetworkService) CountLACP()  {}

// linkset
func (s *NetworkService) AddLinkSet()    {}
func (s *NetworkService) DeleteLinkSet() {}
func (s *NetworkService) GetAllLinkSet() {}
func (s *NetworkService) GetLinkSet()    {}
func (s *NetworkService) CountLinkSet()  {}

// linkset_binding
func (s *NetworkService) GetAllLinkSetBinding() {}
func (s *NetworkService) GetLinkSetBinding()    {}

// linkset_channel_binding
func (s *NetworkService) AddLinkSetChannelBinding()    {}
func (s *NetworkService) DeleteLinkSetChannelBinding() {}
func (s *NetworkService) GetAllLinkSetChannelBinding() {}
func (s *NetworkService) GetLinkSetChannelBinding()    {}
func (s *NetworkService) CountLinkSetChannelBinding()  {}

// linkset_interface_binding
func (s *NetworkService) AddLinkSetInterfaceBinding()    {}
func (s *NetworkService) DeleteLinkSetInterfaceBinding() {}
func (s *NetworkService) GetAllLinkSetInterfaceBinding() {}
func (s *NetworkService) GetLinkSetInterfaceBinding()    {}
func (s *NetworkService) CountLinkSetInterfaceBinding()  {}

// mapbmr
func (s *NetworkService) AddMAPBMR()    {}
func (s *NetworkService) DeleteMAPBMR() {}
func (s *NetworkService) GetAllMAPBMR() {}
func (s *NetworkService) GetMAPBMR()    {}
func (s *NetworkService) CountMAPBMR()  {}

// mapbr_binding
func (s *NetworkService) GetAllMAPBMRBinding() {}
func (s *NetworkService) GetMAPBMRBinding()    {}

// mapbmr_bmrv4network_binding
func (s *NetworkService) AddMAPBMRBMRV4NetworkBinding()    {}
func (s *NetworkService) DeleteMAPBMRBMRV4NetworkBinding() {}
func (s *NetworkService) GetAllMAPBMRBMRV4NetworkBinding() {}
func (s *NetworkService) GetMAPBMRBMRV4NetworkBinding()    {}
func (s *NetworkService) CountMAPBMRBMRV4NetworkBinding()  {}

// mapdmr
func (s *NetworkService) AddMAPDMR()    {}
func (s *NetworkService) DeleteMAPDMR() {}
func (s *NetworkService) GetAllMAPDMR() {}
func (s *NetworkService) GetMAPDMR()    {}
func (s *NetworkService) CountMAPDMR()  {}

// mapdomain
func (s *NetworkService) AddMAPDomain()    {}
func (s *NetworkService) DeleteMAPDomain() {}
func (s *NetworkService) GetAllMAPDomain() {}
func (s *NetworkService) GetMAPDomain()    {}
func (s *NetworkService) CountMAPDomain()  {}

// mapdomain_binding
func (s *NetworkService) GetAllMAPDomainBinding() {}
func (s *NetworkService) GetMAPDomainBinding()    {}

// mapdomain_mapbmr_binding
func (s *NetworkService) AddMAPDomainMAPBMRBinding()    {}
func (s *NetworkService) DeleteMAPDomainMAPBMRBinding() {}
func (s *NetworkService) GetAllMAPDomainMAPBMRBinding() {}
func (s *NetworkService) GetMAPDomainMAPBMRBinding()    {}
func (s *NetworkService) CountMAPDomainMAPBMRBinding()  {}

// nat64
func (s *NetworkService) AddNAT64()    {}
func (s *NetworkService) DeleteNAT64() {}
func (s *NetworkService) UpdateNAT64() {}
func (s *NetworkService) UnsetNAT64()  {}
func (s *NetworkService) GetAllNAT64() {}
func (s *NetworkService) GetNAT64()    {}
func (s *NetworkService) CountNAT64()  {}

// nat64param
func (s *NetworkService) UpdateNAT64Param() {}
func (s *NetworkService) UnsetNAT64Param()  {}
func (s *NetworkService) GetAllNAT64Param() {}
func (s *NetworkService) GetNAT64Param()    {}
func (s *NetworkService) CountNAT64Param()  {}

// nd6
func (s *NetworkService) AddND6()    {}
func (s *NetworkService) DeleteND6() {}
func (s *NetworkService) ClearND6()  {}
func (s *NetworkService) GetAllND6() {}
func (s *NetworkService) CountND6()  {}

// nd6ravariables
func (s *NetworkService) UpdateND6RAVariables() {}
func (s *NetworkService) UnsetND6RAVariables()  {}
func (s *NetworkService) GetAllND6RAVariables() {}
func (s *NetworkService) GetND6RAVariables()    {}
func (s *NetworkService) CountND6RAVariables()  {}

// nd6ravariables_binding
func (s *NetworkService) GetAllND6RAVariablesBinding() {}
func (s *NetworkService) GetND6RAVariablesBinding()    {}

// nd6ravariables_onlinkipv6prefix_binding
func (s *NetworkService) AddND6RAVariablesOnlinkIPV6PrefixBinding()    {}
func (s *NetworkService) DeleteND6RAVariablesOnlinkIPV6PrefixBinding() {}
func (s *NetworkService) GetAllND6RAVariablesOnlinkIPV6PrefixBinding() {}
func (s *NetworkService) GetND6RAVariablesOnlinkIPV6PrefixBinding()    {}
func (s *NetworkService) CountND6RAVariablesOnlinkIPV6PrefixBinding()  {}

// netbridge
func (s *NetworkService) AddNetBridge()    {}
func (s *NetworkService) DeleteNetBridge() {}
func (s *NetworkService) UpdateNetBridge() {}
func (s *NetworkService) UnsetNetBridge()  {}
func (s *NetworkService) GetAllNetBridge() {}
func (s *NetworkService) GetNetBridge()    {}
func (s *NetworkService) CountNetBridge()  {}

// netbridge_binding
func (s *NetworkService) GetAllNetBridgeBinding() {}
func (s *NetworkService) GetNetBridgeBinding()    {}

// netbridge_iptunnel_binding
func (s *NetworkService) AddNetBridgeIPTunnelBinding()    {}
func (s *NetworkService) DeleteNetBridgeIPTunnelBinding() {}
func (s *NetworkService) GetAllNetBridgeIPTunnelBinding() {}
func (s *NetworkService) GetNetBridgeIPTunnelBinding()    {}
func (s *NetworkService) CountNetBridgeIPTunnelBinding()  {}

// netbridge_nsip6_binding
func (s *NetworkService) AddNetBridgeNSIP6Binding()    {}
func (s *NetworkService) DeleteNetBridgeNSIP6Binding() {}
func (s *NetworkService) GetAllNetBridgeNSIP6Binding() {}
func (s *NetworkService) GetNetBridgeNSIP6Binding()    {}
func (s *NetworkService) CountNetBridgeNSIP6Binding()  {}

// netbridge_nsip_binding
func (s *NetworkService) AddNetBridgeNSIPBinding()    {}
func (s *NetworkService) DeleteNetBridgeNSIPBinding() {}
func (s *NetworkService) GetAllNetBridgeNSIPBinding() {}
func (s *NetworkService) GetNetBridgeNSIPBinding()    {}
func (s *NetworkService) CountNetBridgeNSIPBinding()  {}

// netbridge_vlan_binding
func (s *NetworkService) AddNetBridgeVLANBinding()    {}
func (s *NetworkService) DeleteNetBridgeVLANBinding() {}
func (s *NetworkService) GetAllNetBridgeVLANBinding() {}
func (s *NetworkService) GetNetBridgeVLANBinding()    {}
func (s *NetworkService) CountNetBridgeVLANBinding()  {}

// netprofile
func (s *NetworkService) AddNetProfile()    {}
func (s *NetworkService) DeleteNetProfile() {}
func (s *NetworkService) UpdateNetProfile() {}
func (s *NetworkService) UnsetNetProfile()  {}
func (s *NetworkService) GetAllNetProfile() {}
func (s *NetworkService) GetNetProfile()    {}
func (s *NetworkService) CountNetProfile()  {}

// netprofile_binding
func (s *NetworkService) GetAllNetProfileBinding() {}
func (s *NetworkService) GetNetProfileBinding()    {}

// netprofile_natrule_binding
func (s *NetworkService) AddNetProfileNATRuleBinding()    {}
func (s *NetworkService) DeleteNetProfileNATRuleBinding() {}
func (s *NetworkService) GetAllNetProfileNATRuleBinding() {}
func (s *NetworkService) GetNetProfileNATRuleBinding()    {}
func (s *NetworkService) CountNetProfileNATRuleBinding()  {}

// netprofile_srcportset_binding
func (s *NetworkService) AddNetProfileSrcPortSetBinding()    {}
func (s *NetworkService) DeleteNetProfileSrcPortSetBinding() {}
func (s *NetworkService) GetAllNetProfileSrcPortSetBinding() {}
func (s *NetworkService) GetNetProfileSrcPortSetBinding()    {}
func (s *NetworkService) CountNetProfileSrcPortSetBinding()  {}

// onlinkipv6prefix
func (s *NetworkService) AddOnlinkIPV6Prefix()    {}
func (s *NetworkService) DeleteOnlinkIPV6Prefix() {}
func (s *NetworkService) UpdateOnlinkIPV6Prefix() {}
func (s *NetworkService) UnsetOnlinkIPV6Prefix()  {}
func (s *NetworkService) GetAllOnlinkIPV6Prefix() {}
func (s *NetworkService) GetOnlinkIPV6Prefix()    {}
func (s *NetworkService) CountOnlinkIPV6Prefix()  {}

// ptp
func (s *NetworkService) UpdatePTP() {}
func (s *NetworkService) GetAllPTP() {}

// rnat
func (s *NetworkService) AddRNAT()    {}
func (s *NetworkService) DeleteRNAT() {}
func (s *NetworkService) UpdateRNAT() {}
func (s *NetworkService) UnsetRNAT()  {}
func (s *NetworkService) ClearRNAT()  {}
func (s *NetworkService) RenameRNAT() {}
func (s *NetworkService) GetAllRNAT() {}
func (s *NetworkService) GetRNAT()    {}
func (s *NetworkService) CountRNAT()  {}

// rnat6
func (s *NetworkService) AddRNAT6()    {}
func (s *NetworkService) DeleteRNAT6() {}
func (s *NetworkService) UpdateRNAT6() {}
func (s *NetworkService) UnsetRNAT6()  {}
func (s *NetworkService) GetAllRNAT6() {}
func (s *NetworkService) GetRNAT6()    {}
func (s *NetworkService) CountRNAT6()  {}

// rnat6_binding
func (s *NetworkService) GetAllRNAT6Binding() {}
func (s *NetworkService) GetRNAT6Binding()    {}

// rnat6_nsip6_binding
func (s *NetworkService) AddRNAT6NSIP6Binding()    {}
func (s *NetworkService) DeleteRNAT6NSIP6Binding() {}
func (s *NetworkService) GetAllRNAT6NSIP6Binding() {}
func (s *NetworkService) GetRNAT6NSIP6Binding()    {}
func (s *NetworkService) CountRNAT6NSIP6Binding()  {}

// rnatglobal_auditsyslogpolicy_binding
func (s *NetworkService) AddRNATGlobalAuditSyslogPolicyBinding()    {}
func (s *NetworkService) DeleteRNATGlobalAuditSyslogPolicyBinding() {}
func (s *NetworkService) GetRNATGlobalAuditSyslogPolicyBinding()    {}
func (s *NetworkService) CountRNATGlobalAuditSyslogPolicyBinding()  {}

// rnatglobal_binding
func (s *NetworkService) GetRBATGlobalBinding() {}

// rnatparam
func (s *NetworkService) UpdateRNATParam() {}
func (s *NetworkService) UnsetRNATParam()  {}
func (s *NetworkService) GetAllRNATParam() {}

// rnatsession
func (s *NetworkService) FlushRNATSession() {}

// rnat_binding
func (s *NetworkService) GetAllRNATBinding() {}
func (s *NetworkService) GetRNATBinding()    {}

// rnat_nsip_binding
func (s *NetworkService) AddRNATNSIPBinding()    {}
func (s *NetworkService) DeleteRNATNSIPBinding() {}
func (s *NetworkService) GetAllRNATNSIPBinding() {}
func (s *NetworkService) GetRNATNSIPBinding()    {}
func (s *NetworkService) CountRNATNSIPBinding()  {}

// rnat_retainsourceportset_binding
func (s *NetworkService) AddRNATRetainSourcePortSetBinding()    {}
func (s *NetworkService) DeleteRNATRetainSourcePortSetBinding() {}
func (s *NetworkService) GetAllRNATRetainSourcePortSetBinding() {}
func (s *NetworkService) GetRNATRetainSourcePortSetBinding()    {}
func (s *NetworkService) CountRNATRetainSourcePortSetBinding()  {}

// route
func (s *NetworkService) AddRoute()    {}
func (s *NetworkService) DeleteRoute() {}
func (s *NetworkService) UpdateRoute() {}
func (s *NetworkService) UnsetRoute()  {}
func (s *NetworkService) ClearRoute()  {}
func (s *NetworkService) GetAllRoute() {}
func (s *NetworkService) CountRoute()  {}

// route6
func (s *NetworkService) AddRoute6()    {}
func (s *NetworkService) DeleteRoute6() {}
func (s *NetworkService) UpdateRoute6() {}
func (s *NetworkService) UnsetRoute6()  {}
func (s *NetworkService) ClearRoute6()  {}
func (s *NetworkService) GetAllRoute6() {}
func (s *NetworkService) CountRoute6()  {}

// rsskeytype
func (s *NetworkService) UpdateRSSKeyType() {}
func (s *NetworkService) GetAllRSSKeyType() {}

// vlan
func (s *NetworkService) AddVLAN()    {}
func (s *NetworkService) DeleteVLAN() {}
func (s *NetworkService) UpdateVLAN() {}
func (s *NetworkService) UnsetVLAN()  {}
func (s *NetworkService) GetAllVLAN() {}
func (s *NetworkService) GetVLAN()    {}
func (s *NetworkService) CountVLAN()  {}

// vlan_binding
func (s *NetworkService) GetAllVLANBinding() {}
func (s *NetworkService) GetVLANBinding()    {}

// vlan_channel_binding
func (s *NetworkService) AddVLANChannelBinding()    {}
func (s *NetworkService) DeleteVLANChannelBinding() {}
func (s *NetworkService) GetAllVLANChannelBinding() {}
func (s *NetworkService) GetVLANChannelBinding()    {}
func (s *NetworkService) CountVLANChannelBinding()  {}

// vlan_interface_binding
func (s *NetworkService) AddVLANInterfaceBinding()    {}
func (s *NetworkService) DeleteVLANInterfaceBinding() {}
func (s *NetworkService) GetAllVLANInterfaceBinding() {}
func (s *NetworkService) GetVLANInterfaceBinding()    {}
func (s *NetworkService) CountVLANInterfaceBinding()  {}

// vlan_linkset_binding
func (s *NetworkService) AddVLANLinkSetBinding()    {}
func (s *NetworkService) DeleteVLANLinkSetBinding() {}
func (s *NetworkService) GetAllVLANLinkSetBinding() {}
func (s *NetworkService) GetVLANLinkSetBinding()    {}
func (s *NetworkService) CountVLANLinkSetBinding()  {}

// vlan_nsip6_binding
func (s *NetworkService) AddVLANNSIP6Binding()    {}
func (s *NetworkService) DeleteVLANNSIP6Binding() {}
func (s *NetworkService) GetAllVLANNSIP6Binding() {}
func (s *NetworkService) GetVLANNSIP6Binding()    {}
func (s *NetworkService) CountVLANNSIP6Binding()  {}

// vlan_nsip_binding
func (s *NetworkService) AddVLANNSIPBinding()    {}
func (s *NetworkService) DeleteVLANNSIPBinding() {}
func (s *NetworkService) GetAllVLANNSIPBinding() {}
func (s *NetworkService) GetVLANNSIPBinding()    {}
func (s *NetworkService) CountVLANNSIPBinding()  {}

// vrid
func (s *NetworkService) AddVRID()    {}
func (s *NetworkService) DeleteVRID() {}
func (s *NetworkService) UpdateVRID() {}
func (s *NetworkService) UnsetVRID()  {}
func (s *NetworkService) GetAllVRID() {}
func (s *NetworkService) GetVRID()    {}
func (s *NetworkService) CountVRID()  {}

// vrid6
func (s *NetworkService) AddVRID6()    {}
func (s *NetworkService) DeleteVRID6() {}
func (s *NetworkService) UpdateVRID6() {}
func (s *NetworkService) UnsetVRID6()  {}
func (s *NetworkService) GetAllVRID6() {}
func (s *NetworkService) GetVRID6()    {}
func (s *NetworkService) CountVRID6()  {}

// vrid6_binding
func (s *NetworkService) GetAllVRID6Biding() {}
func (s *NetworkService) GetVRID6Biding()    {}

// vrid6_channel_binding
func (s *NetworkService) AddVRID6ChannelBinding()    {}
func (s *NetworkService) DeleteVRID6ChannelBinding() {}
func (s *NetworkService) GetAllVRID6ChannelBinding() {}
func (s *NetworkService) GetVRID6ChannelBinding()    {}
func (s *NetworkService) CountVRID6ChannelBinding()  {}

// vrid6_interface_binding
func (s *NetworkService) AddVRID6InterfaceBinding()    {}
func (s *NetworkService) DeleteVRID6InterfaceBinding() {}
func (s *NetworkService) GetAllVRID6InterfaceBinding() {}
func (s *NetworkService) GetVRID6InterfaceBinding()    {}
func (s *NetworkService) CountVRID6InterfaceBinding()  {}

// vrid6_nsip6_binding
func (s *NetworkService) GetAllVRID6NSIP6Binding() {}
func (s *NetworkService) GetVRID6NSIP6Binding()    {}
func (s *NetworkService) CountVRID6NSIP6Binding()  {}

// vrid6_nsip_binding
func (s *NetworkService) GetAllVRID6NSIPBinding() {}
func (s *NetworkService) GetVRID6NSIPBinding()    {}
func (s *NetworkService) CountVRID6NSIPBinding()  {}

// vrid6_trackinterface_binding
func (s *NetworkService) AddVRID6TrackInterfaceBinding()    {}
func (s *NetworkService) DeleteVRID6TrackInterfaceBinding() {}
func (s *NetworkService) GetAllVRID6TrackInterfaceBinding() {}
func (s *NetworkService) GetVRID6TrackInterfaceBinding()    {}
func (s *NetworkService) CountVRID6TrackInterfaceBinding()  {}

// vridparam
func (s *NetworkService) UpdateVRIDParam() {}
func (s *NetworkService) UnsetVRIDParam()  {}
func (s *NetworkService) GetAllVRIDParam() {}

// vrid_binding
func (s *NetworkService) GetAllVRIDBinding() {}
func (s *NetworkService) GetVRIDBinding()    {}

// vrid_channel_binding
func (s *NetworkService) AddVRIDChannelBinding()    {}
func (s *NetworkService) DeleteVRIDChannelBinding() {}
func (s *NetworkService) GetAllVRIDChannelBinding() {}
func (s *NetworkService) GetVRIDChannelBinding()    {}
func (s *NetworkService) CountVRIDChannelBinding()  {}

// vrid_interface_binding
func (s *NetworkService) AddVRIDInterfaceBinding()    {}
func (s *NetworkService) DeleteVRIDInterfaceBinding() {}
func (s *NetworkService) GetAllVRIDInterfaceBinding() {}
func (s *NetworkService) GetVRIDInterfaceBinding()    {}
func (s *NetworkService) CountVRIDInterfaceBinding()  {}

// vrid_nsip6_binding
func (s *NetworkService) GetAllVRIDNSIP6Binding() {}
func (s *NetworkService) GetVRIDNSIP6Binding()    {}
func (s *NetworkService) CountVRIDNSIP6Binding()  {}

// vrid_nsip_binding
func (s *NetworkService) GetAllVRIDNSIPBinding() {}
func (s *NetworkService) GetVRIDNSIPBinding()    {}
func (s *NetworkService) CountVRIDNSIPBinding()  {}

// vrid_trackinterface_binding
func (s *NetworkService) AddVRIDTrackInterfaceBinding()    {}
func (s *NetworkService) DeleteVRIDTrackInterfaceBinding() {}
func (s *NetworkService) GetAllVRIDTrackInterfaceBinding() {}
func (s *NetworkService) GetVRIDTrackInterfaceBinding()    {}
func (s *NetworkService) CountVRIDTrackInterfaceBinding()  {}

// vxlan
func (s *NetworkService) AddVXLAN()    {}
func (s *NetworkService) DeleteVXLAN() {}
func (s *NetworkService) UpdateVXLAN() {}
func (s *NetworkService) UnsetVXLAN()  {}
func (s *NetworkService) GetAllVXLAN() {}
func (s *NetworkService) GetVXLAN()    {}
func (s *NetworkService) CountVXLAN()  {}

// vxlanvlanmap
func (s *NetworkService) AddVXLANVLANMap()    {}
func (s *NetworkService) DeleteVXLANVLANMap() {}
func (s *NetworkService) GetAllVXLANVLANMap() {}
func (s *NetworkService) GetVXLANVLANMap()    {}
func (s *NetworkService) CountVXLANVLANMap()  {}

// vxlanvlanmap_binding
func (s *NetworkService) GetAllVXLANVLANMapBinding() {}
func (s *NetworkService) GetVXLANVLANMapBinding()    {}

// vxlanvlanmap_vxlan_binding
func (s *NetworkService) AddVXLANVLANMapVXLANBinding()    {}
func (s *NetworkService) DeleteVXLANVLANMapVXLANBinding() {}
func (s *NetworkService) GetAllVXLANVLANMapVXLANBinding() {}
func (s *NetworkService) GetVXLANVLANMapVXLANBinding()    {}
func (s *NetworkService) CountVXLANVLANMapVXLANBinding()  {}

// vxlan_binding
func (s *NetworkService) GetAllVXLANBinding() {}
func (s *NetworkService) GetVXLANBinding()    {}

// vxlan_iptunnel_binding
func (s *NetworkService) GetAllVXLANIPTunnelBinding() {}
func (s *NetworkService) GetVXLANIPTunnelBinding()    {}
func (s *NetworkService) CountVXLANIPTunnelBinding()  {}

// vxlan_nsip6_binding
func (s *NetworkService) AddVXLANNSIP6Binding()    {}
func (s *NetworkService) DeleteVXLANNSIP6Binding() {}
func (s *NetworkService) GetAllVXLANNSIP6Binding() {}
func (s *NetworkService) GetVXLANNSIP6Binding()    {}
func (s *NetworkService) CountVXLANNSIP6Binding()  {}

// vxlan_nsip_binding
func (s *NetworkService) AddVXLANNSIPBinding()    {}
func (s *NetworkService) DeleteVXLANNSIPBinding() {}
func (s *NetworkService) GetAllVXLANNSIPBinding() {}
func (s *NetworkService) GetVXLANNSIPBinding()    {}
func (s *NetworkService) CountVXLANNSIPBinding()  {}

// vxlan_srcip_binding
func (s *NetworkService) AddVXLANSrcIPBinding()    {}
func (s *NetworkService) DeleteVXLANSrcIPBinding() {}
func (s *NetworkService) GetAllVXLANSrcIPBinding() {}
func (s *NetworkService) GetVXLANSrcIPBinding()    {}
func (s *NetworkService) CountVXLANSrcIPBinding()  {}
