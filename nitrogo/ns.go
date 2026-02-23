package nitrogo

const (
	nsACLURL                               = "/nitro/v1/config/nsacl"
	nsACL6URL                              = "/nitro/v1/config/nsacl6"
	nsAppFlowCollectorURL                  = "/nitro/v1/config/nsappflowcollector"
	nsAppFlowParamURL                      = "/nitro/v1/config/nsappflowparam"
	nsCapacityURL                          = "/nitro/v1/config/nscapacity"
	nsConfigURL                            = "/nitro/v1/config/nsconfig"
	nsConnectionTableURL                   = "/nitro/v1/config/nsconnectiontable"
	nsConsoleLoginPromptURL                = "/nitro/v1/config/nsconsoleloginprompt"
	nsDHCPParamsURL                        = "/nitro/v1/config/nsdhcpparams"
	nsDialogURL                            = "/nitro/v1/config/nsdialog"
	nsEncryptionParamsURL                  = "/nitro/v1/config/nsencryptionparams"
	nsEventsURL                            = "/nitro/v1/config/nsevents"
	nsExtensionURL                         = "/nitro/v1/config/nsextension"
	nsExtensionBindingURL                  = "/nitro/v1/config/nsextension_binding"
	nsExtensionExtensionFunctionBindingURL = "/nitro/v1/config/nsextension_extensionfunction_binding"
	nsFeatureURL                           = "/nitro/v1/config/nsfeature"
	nsHardwareURL                          = "/nitro/v1/config/nshardware"
	nsHostNameURL                          = "/nitro/v1/config/nshostname"
	nsHTTPParamURL                         = "/nitro/v1/config/nshttpparam"
	nsHTTPProfileURL                       = "/nitro/v1/config/nshttpprofile"
	nsICAPProfileURL                       = "/nitro/v1/config/nsicapprofile"
	nsIPURL                                = "/nitro/v1/config/nsip"
	nsIP6URL                               = "/nitro/v1/config/nsip6"
	nsLicenseURL                           = "/nitro/v1/config/nslicense"
	nsLicenseParametersURL                 = "/nitro/v1/config/nslicenseparameters"
	nsLicenseServerURL                     = "/nitro/v1/config/nslicenseserver"
	nsLimitIdentifierURL                   = "/nitro/v1/config/nslimitidentifier"
	nsLimitSelectorURL                     = "/nitro/v1/config/nslimitselector"
	nsLimitSessionsURL                     = "/nitro/v1/config/nslimitsessions"
	nsLogActionURL                         = "/nitro/v1/config/nslogaction"
	nsLogGlobalBindingURL                  = "/nitro/v1/config/nslogglobal_binding"
	nsLogGlobalAuditNSLogPolicyBindingURL  = "/nitro/v1/config/nslogglobal_auditnslogpolicy_binding"
	nsLogParamsURL                         = "/nitro/v1/config/nslogparams"
	nsLogPolicyURL                         = "/nitro/v1/config/nslogpolicy"
	nsLogPolicyBindingURL                  = "/nitro/v1/config/nslogpolicy_binding"
	nsModeURL                              = "/nitro/v1/config/nsmode"
	nsParamURL                             = "/nitro/v1/config/nsparam"
	nsPartitionURL                         = "/nitro/v1/config/nspartition"
	nsPartitionBindingURL                  = "/nitro/v1/config/nspartition_binding"
	nsPartitionBridgeGroupBindingURL       = "/nitro/v1/config/nspartition_bridgegroup_binding"
	nsPartitionVLANBindingURL              = "/nitro/v1/config/nspartition_vlan_binding"
	nsPartitionVXLANBindingURL             = "/nitro/v1/config/nspartition_vxlan_binding"
	nsRateControlURL                       = "/nitro/v1/config/nsratecontrol"
	nsRPCNodeURL                           = "/nitro/v1/config/nsrpcnode"
	nsRunningConfigURL                     = "/nitro/v1/config/nsrunningconfig"
	nsSavedConfigURL                       = "/nitro/v1/config/nssavedconfig"
	nsServicePathURL                       = "/nitro/v1/config/nsservicepath"
	nsSimpleACLURL                         = "/nitro/v1/config/nssimpleacl"
	nsSimpleACL6URL                        = "/nitro/v1/config/nssimpleacl6"
	nsSPParamsURL                          = "/nitro/v1/config/nsspparams"
	nsSurgeProtectionURL                   = "/nitro/v1/config/nssurgeprotection"
	nsTCPBufParamURL                       = "/nitro/v1/config/nstcpbufparam"
	nsTCPParamURL                          = "/nitro/v1/config/nstcpparam"
	nsTCPProfileURL                        = "/nitro/v1/config/nstcpprofile"
	nsTimeoutURL                           = "/nitro/v1/config/nstimeout"
	nsTimerURL                             = "/nitro/v1/config/nstimer"
	nsTrafficDomainURL                     = "/nitro/v1/config/nstrafficdomain"
	nsTrafficDomainBindingURL              = "/nitro/v1/config/nstrafficdomain_binding"
	nsTrafficDomainBridgeGroupBindingURL   = "/nitro/v1/config/nstrafficdomain_bridgegroup_binding"
	nsTrafficDomainVLANBindingURL          = "/nitro/v1/config/nstrafficdomain_vlan_binding"
	nsTrafficDomainVXLANBindingURL         = "/nitro/v1/config/nstrafficdomain_vxlan_binding"
	nsVariableURL                          = "/nitro/v1/config/nsvariable"
	nsVersionURL                           = "/nitro/v1/config/nsversion"
	nsWebLogParamURL                       = "/nitro/v1/config/nsweblogparam"
	nsXMLNamespaceURL                      = "/nitro/v1/config/nsxmlnamespace"
	nsXMLSchemaURL                         = "/nitro/v1/config/nsxmlschema"
)

// ns
// System/Global level configuration.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/ns/ns
type NSService struct {
	client *Client
}

// nsacl
func (s *NSService) AddNSACL()     {}
func (s *NSService) DeleteNSACL()  {}
func (s *NSService) UpdateNSACL()  {}
func (s *NSService) UnsetNSACL()   {}
func (s *NSService) EnableNSACL()  {}
func (s *NSService) DisableNSACL() {}
func (s *NSService) RenameNSACL()  {}
func (s *NSService) GetAllNSACL()  {}
func (s *NSService) GetNSACL()     {}
func (s *NSService) CountNSACL()   {}

// nsacl6
func (s *NSService) AddNSACL6()     {}
func (s *NSService) DeleteNSACL6()  {}
func (s *NSService) UpdateNSACL6()  {}
func (s *NSService) UnsetNSACL6()   {}
func (s *NSService) EnableNSACL6()  {}
func (s *NSService) DisableNSACL6() {}
func (s *NSService) RenameNSACL6()  {}
func (s *NSService) GetAllNSACL6()  {}
func (s *NSService) GetNSACL6()     {}
func (s *NSService) CountNSACL6()   {}

// nsacls
func (s *NSService) RenumberNSACLS() {}
func (s *NSService) ClearNSACLS()    {}
func (s *NSService) ApplyNSACLS()    {}

// nsacls6
func (s *NSService) RenumberNSACLS6() {}
func (s *NSService) ClearNSACLS6()    {}
func (s *NSService) ApplyNSACLS6()    {}

// nsappflowcollector
func (s *NSService) AddNSAppFlowCollector()    {}
func (s *NSService) DeleteNSAppFlowCollector() {}
func (s *NSService) GetAllNSAppFlowCollector() {}
func (s *NSService) GetNSAppFlowCollector()    {}
func (s *NSService) CountNSAppFlowCollector()  {}

// nsappflowparam
func (s *NSService) UpdateNSAppFlowParam() {}
func (s *NSService) UnsetNSAppFlowParam()  {}
func (s *NSService) GetAllNSAppFlowParam() {}

// nsaptlicense
func (s *NSService) GetAllNSAptLicense() {}
func (s *NSService) CountNSAptLicense()  {}
func (s *NSService) ChangeNSAptLicense() {}

// nsassignment
func (s *NSService) AddNSAssignment()    {}
func (s *NSService) DeleteNSAssignment() {}
func (s *NSService) UpdateNSAssignment() {}
func (s *NSService) UnsetNSAssignment()  {}
func (s *NSService) GetAllNSAssignment() {}
func (s *NSService) GetNSAssignment()    {}
func (s *NSService) CountNSAssignment()  {}
func (s *NSService) RenameNSAssignment() {}

// nscapacity
func (s *NSService) UpdateNSCapacity() {}
func (s *NSService) UnsetNSCapacity()  {}
func (s *NSService) GetAllNSCapacity() {}

// nscentralmanagementserver
func (s *NSService) AddNSCentralManagementServer()    {}
func (s *NSService) DeleteNSCentralManagementServer() {}
func (s *NSService) GetAllNSCentralManagementServer() {}
func (s *NSService) GetNSCentralManagementServer()    {}
func (s *NSService) CountNSCentralManagementServer()  {}

// nsconfig
func (s *NSService) ClearNSConfig()  {}
func (s *NSService) SaveNSConfig()   {}
func (s *NSService) DiffNSConfig()   {}
func (s *NSService) UpdateNSConfig() {}
func (s *NSService) UnsetNSConfig()  {}
func (s *NSService) GetAllNSConfig() {}

// nsconnectiontable
func (s *NSService) GetAllNSconnectionTable() {}
func (s *NSService) CountNSconnectionTable()  {}

// nsconsoleloginprompt
func (s *NSService) UpdateNSConsoleLoginPrompt() {}
func (s *NSService) UnsetNSConsoleLoginPrompt()  {}
func (s *NSService) GetAllNSConsoleLoginPrompt() {}

// nscqaparam
func (s *NSService) UpdateNSCQAParam() {}
func (s *NSService) UnsetNSCQAParam()  {}
func (s *NSService) GetAllNSCQAParam() {}

// nsdhcpip
func (s *NSService) ReleaseNSDHCPIP() {}

// nsdhcpparams
func (s *NSService) UpdateNSDHCPParams() {}
func (s *NSService) UnsetNSDHCPParams()  {}
func (s *NSService) GetAllNSDHCPParams() {}

// nsdiamter
func (s *NSService) UpdateNSDiameter() {}
func (s *NSService) UnsetNSDiameter()  {}
func (s *NSService) GetAllNSDiameter() {}
func (s *NSService) GetNSDiameter()    {}
func (s *NSService) CountNSDiameter()  {}

// nsencyrptionkey
func (s *NSService) AddNSEncryptionKey()    {}
func (s *NSService) DeleteNSEncryptionKey() {}
func (s *NSService) UpdateNSEncryptionKey() {}
func (s *NSService) UnsetNSEncryptionKey()  {}
func (s *NSService) GetAllNSEncryptionKey() {}
func (s *NSService) GetNSEncryptionKey()    {}
func (s *NSService) CountNSEncryptionKey()  {}

// nsencryptionparams
func (s *NSService) UpdateNSEncryptionParams() {}
func (s *NSService) GetAllNSEncryptionParams() {}

// nsevents
func (s *NSService) GetAllNSEvents() {}
func (s *NSService) CountNSEvents()  {}

// nsextension
func (s *NSService) ImportNSExtension() {}
func (s *NSService) AddNSExtension()    {}
func (s *NSService) DeleteNSExtension() {}
func (s *NSService) UpdateNSExtension() {}
func (s *NSService) UnsetNSExtension()  {}
func (s *NSService) GetAllNSExtension() {}
func (s *NSService) GetNSExtension()    {}
func (s *NSService) CountNSExtension()  {}
func (s *NSService) ChangeNSExtension() {}

// nsextension_binding
func (s *NSService) GetAllNSExtenionBinding() {}
func (s *NSService) GetNSExtenionBinding()    {}

// nsextension_extensionfunction_binding
func (s *NSService) GetAllNSExtensionExtensionFunctionBinding() {}
func (s *NSService) GetNSExtensionExtensionFunctionBinding()    {}
func (s *NSService) CountNSExtensionExtensionFunctionBinding()  {}

// nsfeature
func (s *NSService) EnableNSFeature()  {}
func (s *NSService) DisableNSFeature() {}
func (s *NSService) GetAllNSFeature()  {}

// nshardware
func (s *NSService) GetAllNSHardware() {}

// nshmackey
func (s *NSService) AddNSHMACKey()    {}
func (s *NSService) DeleteNSHMACKey() {}
func (s *NSService) UpdateNSHMACKey() {}
func (s *NSService) UnsetNSHMACKey()  {}
func (s *NSService) GetAllNSHMACKey() {}
func (s *NSService) GetNSHMACKey()    {}
func (s *NSService) CountNSHMACKey()  {}

// nshostname
func (s *NSService) UpdateNSHostname() {}
func (s *NSService) GetAllNSHostname() {}
func (s *NSService) CountNSHostname()  {}

// nshttpparam
func (s *NSService) UpdateNSHTTPParam() {}
func (s *NSService) GetAllNSHTTPParam() {}
func (s *NSService) CountNSHTTPParam()  {}

// nshttpprofile
func (s *NSService) AddNSHTTPProfile()    {}
func (s *NSService) DeleteNSHTTPProfile() {}
func (s *NSService) UpdateNSHTTPProfile() {}
func (s *NSService) UnsetNSHTTPProfile()  {}
func (s *NSService) GetAllNSHTTPProfile() {}
func (s *NSService) GetNSHTTPProfile()    {}
func (s *NSService) CountNSHTTPProfile()  {}

// nsicapprofile
func (s *NSService) AddNSICAPProfile()    {}
func (s *NSService) DeleteNSICAPProfile() {}
func (s *NSService) UpdateNSICAPProfile() {}
func (s *NSService) UnsetNSICAPProfile()  {}
func (s *NSService) GetAllNSICAPProfile() {}
func (s *NSService) GetNSICAPProfile()    {}
func (s *NSService) CountNSICAPProfile()  {}

// nsip
func (s *NSService) AddNSIP()     {}
func (s *NSService) DeleteNSIP()  {}
func (s *NSService) UpdateNSIP()  {}
func (s *NSService) UnsetNSIP()   {}
func (s *NSService) EnableNSIP()  {}
func (s *NSService) DisableNSIP() {}
func (s *NSService) GetAllNSIP()  {}
func (s *NSService) CountNSIP()   {}

// nsip6
func (s *NSService) AddNSIP6()    {}
func (s *NSService) DeleteNSIP6() {}
func (s *NSService) UpdateNSIP6() {}
func (s *NSService) UnsetNSIP6()  {}
func (s *NSService) GetAllNSIP6() {}
func (s *NSService) CountNSIP6()  {}

// nslicense
func (s *NSService) GetAllNSLicense() {}

// nslicenseproxyserver
func (s *NSService) AddNSLicenseProxyServer()    {}
func (s *NSService) DeleteNSLicenseProxyServer() {}
func (s *NSService) UpdateNSLicenseProxyServer() {}
func (s *NSService) GetAllNSLicenseProxyServer() {}
func (s *NSService) CountNSLicenseProxyServer()  {}

// nslicenseserver
func (s *NSService) AddNSLicenseServer()    {}
func (s *NSService) DeleteNSLicenseServer() {}
func (s *NSService) UpdateNSLicenseServer() {}
func (s *NSService) GetAllNSLicenseServer() {}
func (s *NSService) CountNSLicenseServer()  {}

// nslicenseserverpool
func (s *NSService) GetAllNSLicenseServerPool() {}

// nslimitidentifier
func (s *NSService) AddNSLimitIdentifier()    {}
func (s *NSService) DeleteNSLimitIdentifier() {}
func (s *NSService) UpdateNSLimitIdentifier() {}
func (s *NSService) UnsetNSLimitIdentifier()  {}
func (s *NSService) GetAllNSLimitIdentifier() {}
func (s *NSService) GetNSLimitIdentifier()    {}
func (s *NSService) CountNSLimitIdentifier()  {}

// nslimitidentifier_binding
func (s *NSService) GetAllNSLimitIdentifierBinding() {}
func (s *NSService) GetNSLimitIdentifierBinding()    {}

// nslimitidentifier_nslimitsessions_binding
func (s *NSService) GetAllNSLimitIdentifierNSLimitSessionsBinding() {}
func (s *NSService) GetNSLimitIdentifierNSLimitSessionsBinding()    {}
func (s *NSService) CountNSLimitIdentifierNSLimitSessionsBinding()  {}

// nslimitselector
func (s *NSService) AddNSLimitSelector()    {}
func (s *NSService) DeleteNSLimitSelector() {}
func (s *NSService) UpdateNSLimitSelector() {}
func (s *NSService) UnsetNSLimitSelector()  {}
func (s *NSService) GetAllNSLimitSelector() {}
func (s *NSService) GetNSLimitSelector()    {}
func (s *NSService) CountNSLimitSelector()  {}

// nslimitsessions
func (s *NSService) GetAllNSLimitSessions() {}
func (s *NSService) CountNSLimitSessions()  {}
func (s *NSService) ClearNSLimitSessions()  {}

// nsmigration
func (s *NSService) GetAllNSMigration() {}
func (s *NSService) CountNSMigration()  {}

// nsmode
func (s *NSService) EnableNSMode()  {}
func (s *NSService) DisableNSMode() {}
func (s *NSService) GetAllNSMode()  {}

// nsparam
func (s *NSService) UpdateNSParam() {}
func (s *NSService) UnsetNSParam()  {}
func (s *NSService) GetAllNSParam() {}

// nspartition
func (s *NSService) AddNSPartition()    {}
func (s *NSService) DeleteNSPartition() {}
func (s *NSService) UpdateNSPartition() {}
func (s *NSService) UnsetNSPartition()  {}
func (s *NSService) SwitchNSPartition() {}
func (s *NSService) GetAllNSPartition() {}
func (s *NSService) GetNSPartition()    {}
func (s *NSService) CountNSPartition()  {}

// nspartitionmac
func (s *NSService) GetAllNSPartitionMAC() {}
func (s *NSService) CountNSPartitionMAC()  {}

// nspartition_binding
func (s *NSService) GetAllNSPartitionBinding() {}
func (s *NSService) GetNSPartitionBinding()    {}

// nspartition_bridgegroup_binding
func (s *NSService) AddNSPartitionBridgeGroupBinding()    {}
func (s *NSService) DeleteNSPartitionBridgeGroupBinding() {}
func (s *NSService) GetAllNSPartitionBridgeGroupBinding() {}
func (s *NSService) GetNSPartitionBridgeGroupBinding()    {}
func (s *NSService) CountNSPartitionBridgeGroupBinding()  {}

// nspartition_vlan_binding
func (s *NSService) AddNSPartitionVLANBinding()    {}
func (s *NSService) DeleteNSPartitionVLANBinding() {}
func (s *NSService) GetAllNSPartitionVLANBinding() {}
func (s *NSService) GetNSPartitionVLANBinding()    {}
func (s *NSService) CountNSPartitionVLANBinding()  {}

// nspartition_vxlan_binding
func (s *NSService) AddNSPartitionVXLANBinding()    {}
func (s *NSService) DeleteNSPartitionVXLANBinding() {}
func (s *NSService) GetAllNSPartitionVXLANBinding() {}
func (s *NSService) GetNSPartitionVXLANBinding()    {}
func (s *NSService) CountNSPartitionVXLANBinding()  {}

// nspbr
func (s *NSService) AddNSPBR()     {}
func (s *NSService) DeleteNSPBR()  {}
func (s *NSService) UpdateNSPBR()  {}
func (s *NSService) UnsetNSPBR()   {}
func (s *NSService) EnableNSPBR()  {}
func (s *NSService) DisableNSPBR() {}
func (s *NSService) GetAllNSPBR()  {}
func (s *NSService) GetNSPBR()     {}
func (s *NSService) CountNSPBR()   {}

// nspbr6
func (s *NSService) AddNSPBR6()      {}
func (s *NSService) DeleteNSPBR6()   {}
func (s *NSService) UpdateNSPBR6()   {}
func (s *NSService) UnsetNSPBR6()    {}
func (s *NSService) RenumberNSPBR6() {}
func (s *NSService) EnableNSPBR6()   {}
func (s *NSService) DisableNSPBR6()  {}
func (s *NSService) GetAllNSPBR6()   {}
func (s *NSService) GetNSPBR6()      {}
func (s *NSService) CountNSPBR6()    {}
func (s *NSService) ClearNSPBR6()    {}
func (s *NSService) ApplyNSPBR6()    {}

// nspbrs
func (s *NSService) RenumberNSPBRs() {}
func (s *NSService) ClearNSPBRs()    {}
func (s *NSService) ApplyNSPBRs()    {}

// nsratecontrol
func (s *NSService) UpdateNSrateControl() {}
func (s *NSService) UnsetNSrateControl()  {}
func (s *NSService) GetAllNSrateControl() {}

// nsrollbackcmd
func (s *NSService) GetAllNSRollBackCMD() {}

// nsrpcnode
func (s *NSService) UpdateNSRPCNode() {}
func (s *NSService) UnsetNSRPCNode()  {}
func (s *NSService) GetAllNSRPCNode() {}
func (s *NSService) GetNSRPCNode()    {}
func (s *NSService) CountNSRPCNode()  {}

// nsrunningconfig
func (s *NSService) GetAllNSRunningConfig() {}

// nssavedconfig
func (s *NSService) GetAllNSSavedConfig() {}

// nsservicefunction
func (s *NSService) AddNSServiceFunction()    {}
func (s *NSService) DeleteNSServiceFunction() {}
func (s *NSService) GetAllNSServiceFunction() {}
func (s *NSService) GetNSServiceFunction()    {}
func (s *NSService) CountNSServiceFunction()  {}

// nsservicepath
func (s *NSService) AddNSServicePath()    {}
func (s *NSService) DeleteNSServicePath() {}
func (s *NSService) GetAllNSServicePath() {}
func (s *NSService) GetNSServicePath()    {}
func (s *NSService) CountNSServicePath()  {}

// nsservicepath_binding
func (s *NSService) GetAllNSServicePathBinding() {}
func (s *NSService) GetNSServicePathBinding()    {}

// nsservicepath_nsservicefunction_binding
func (s *NSService) AddNSServicePathNSServiceFunctionBinding()    {}
func (s *NSService) DeleteNSServicePathNSServiceFunctionBinding() {}
func (s *NSService) GetAllNSServicePathNSServiceFunctionBinding() {}
func (s *NSService) GetNSServicePathNSServiceFunctionBinding()    {}
func (s *NSService) CountNSServicePathNSServiceFunctionBinding()  {}

// nssimpleacl
func (s *NSService) AddNSSimpleACL()    {}
func (s *NSService) DeleteNSSimpleACL() {}
func (s *NSService) FlushNSSimpleACL()  {}
func (s *NSService) GetAllNSSimpleACL() {}
func (s *NSService) GetNSSimpleACL()    {}
func (s *NSService) CountNSSimpleACL()  {}
func (s *NSService) ClearNSSimpleACL()  {}

// nssimpleacl6
func (s *NSService) AddNSSimpleACL6()    {}
func (s *NSService) DeleteNSSimpleACL6() {}
func (s *NSService) FlushNSSimpleACL6()  {}
func (s *NSService) GetAllNSSimpleACL6() {}
func (s *NSService) GetNSSimpleACL6()    {}
func (s *NSService) CountNSSimpleACL6()  {}
func (s *NSService) ClearNSSimpleACL6()  {}

// nssourceroutecachetable
func (s *NSService) FlushNSSourceRouteCacheTable()  {}
func (s *NSService) GetAllNSSourceRouteCacheTable() {}
func (s *NSService) CountNSSourceRouteCacheTable()  {}

// nsspparams
func (s *NSService) UpdateNSSPParams() {}
func (s *NSService) UnsetNSSPParams()  {}
func (s *NSService) GetAllNSSPParams() {}

// nsstats
func (s *NSService) ClearNSStats() {}

// nssurgeq
func (s *NSService) FlushNSSurgeQ() {}

// nstcpbufparam
func (s *NSService) UpdateNSTCPBufParam() {}
func (s *NSService) UnsetNSTCPBufParam()  {}
func (s *NSService) GetAllNSTCPBufParam() {}

// nstcpparam
func (s *NSService) UpdateNSTCPParam() {}
func (s *NSService) UnsetNSTCPParam()  {}
func (s *NSService) GetAllNSTCPParam() {}

// nstcpprofile
func (s *NSService) AddNSTCPProfile()    {}
func (s *NSService) DeleteNSTCPProfile() {}
func (s *NSService) UpdateNSTCPProfile() {}
func (s *NSService) UnsetNSTCPProfile()  {}
func (s *NSService) GetAllNSTCPProfile() {}
func (s *NSService) GetNSTCPProfile()    {}
func (s *NSService) CountNSTCPProfile()  {}

// nstimeout
func (s *NSService) UpdateNSTimeout() {}
func (s *NSService) UnsetNSTimeout()  {}
func (s *NSService) GetAllNSTimeout() {}

// nstimer
func (s *NSService) AddNSTimer()    {}
func (s *NSService) DeleteNSTimer() {}
func (s *NSService) UpdateNSTimer() {}
func (s *NSService) UnsetNSTimer()  {}
func (s *NSService) GetAllNSTimer() {}
func (s *NSService) GetNSTimer()    {}
func (s *NSService) CountNSTimer()  {}
func (s *NSService) RenameNSTimer() {}

// nstimer_autoscalepolicy_binding
func (s *NSService) AddNSTimerAutoScalePolicyBinding()    {}
func (s *NSService) DeleteNSTimerAutoScalePolicyBinding() {}
func (s *NSService) GetAllNSTimerAutoScalePolicyBinding() {}
func (s *NSService) GetNSTimerAutoScalePolicyBinding()    {}
func (s *NSService) CountNSTimerAutoScalePolicyBinding()  {}

// nstimer_binding
func (s *NSService) GetAllNSTimerBinding() {}
func (s *NSService) GetNSTimerBinding()    {}

// nstimezone
func (s *NSService) GetAllNSTimeZone() {}
func (s *NSService) GetNSTimeZone()    {}
func (s *NSService) CountNSTimeZone()  {}

// nstrafficdomain
func (s *NSService) AddNSTrafficDomain()     {}
func (s *NSService) DeleteNSTrafficDomain()  {}
func (s *NSService) EnableNSTrafficDomain()  {}
func (s *NSService) DisableNSTrafficDomain() {}
func (s *NSService) GetAllNSTrafficDomain()  {}
func (s *NSService) GetNSTrafficDomain()     {}
func (s *NSService) CountNSTrafficDomain()   {}
func (s *NSService) ClearNSTrafficDomain()   {}

// nstrafficdomain_binding
func (s *NSService) GetAllNSTrafficDomainBinding() {}
func (s *NSService) GetNSTrafficDomainBinding()    {}

// nstrafficdomain_bridgegroup_binding
func (s *NSService) AddBSTrafficDomainBridgeGroupBinding()    {}
func (s *NSService) DeleteBSTrafficDomainBridgeGroupBinding() {}
func (s *NSService) GetAllBSTrafficDomainBridgeGroupBinding() {}
func (s *NSService) GetBSTrafficDomainBridgeGroupBinding()    {}
func (s *NSService) CountBSTrafficDomainBridgeGroupBinding()  {}

// nstrafficdomain_vlan_binding
func (s *NSService) AddNSTrafficDomainVLANBinding()    {}
func (s *NSService) DeleteNSTrafficDomainVLANBinding() {}
func (s *NSService) GetAllNSTrafficDomainVLANBinding() {}
func (s *NSService) GetNSTrafficDomainVLANBinding()    {}
func (s *NSService) CountNSTrafficDomainVLANBinding()  {}

// nstrafficdomain_vxlan_binding
func (s *NSService) AddNSTrafficDomainVXLANBinding()    {}
func (s *NSService) DeleteNSTrafficDomainVXLANBinding() {}
func (s *NSService) GetAllNSTrafficDomainVXLANBinding() {}
func (s *NSService) GetNSTrafficDomainVXLANBinding()    {}
func (s *NSService) CountNSTrafficDomainVXLANBinding()  {}

// nsvariable
func (s *NSService) AddNSVariable()    {}
func (s *NSService) DeleteNSVariable() {}
func (s *NSService) UpdateNSVariable() {}
func (s *NSService) UnsetNSVariable()  {}
func (s *NSService) GetAllNSVariable() {}
func (s *NSService) GetNSVariable()    {}
func (s *NSService) CountNSVariable()  {}

// nsversion
func (s *NSService) GetAllNSVersion() {}

// nsvpxparam
func (s *NSService) UpdateNSVPXParam() {}
func (s *NSService) UnsetNSVPXParam()  {}
func (s *NSService) GetAllNSVPXParam() {}
func (s *NSService) CountNSVPXParam()  {}

// nsweblogparam
func (s *NSService) UpdateNSWebLogParam() {}
func (s *NSService) UnsetNSWebLogParam()  {}
func (s *NSService) GetAllNSWebLogParam() {}

// nsxmlnamespace
func (s *NSService) AddNSXMLNamespace()    {}
func (s *NSService) DeleteNSXMLNamespace() {}
func (s *NSService) UpdateNSXMLNamespace() {}
func (s *NSService) UnsetNSXMLNamespace()  {}
func (s *NSService) GetAllNSXMLNamespace() {}
func (s *NSService) GetNSXMLNamespace()    {}
func (s *NSService) CountNSXMLNamespace()  {}

// reboot
func (s *NSService) Reboot() {}

// shutdown
func (s *NSService) Shutdown() {}
