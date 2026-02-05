package nitrogo

const (
	botGlobalBindingURL                   = "/nitro/v1/config/botglobal_binding"
	botGlobalBotPolicyBindingURL          = "/nitro/v1/config/botglobal_botpolicy_binding"
	botPolicyURL                          = "/nitro/v1/config/botpolicy"
	botPolicyBindingURL                   = "/nitro/v1/config/botpolicy_binding"
	botPolicyBotGlobalBindingURL          = "/nitro/v1/config/botpolicy_botglobal_binding"
	botPolicyBotPolicyLabelBindingURL     = "/nitro/v1/config/botpolicy_botpolicylabel_binding"
	botPolicyCSVServerBindingURL          = "/nitro/v1/config/botpolicy_csvserver_binding"
	botPolicyLBVServerBindingURL          = "/nitro/v1/config/botpolicy_lbvserver_binding"
	botPolicyLabelURL                     = "/nitro/v1/config/botpolicylabel"
	botPolicyLabelBindingURL              = "/nitro/v1/config/botpolicylabel_binding"
	botPolicyLabelBotPolicyBindingURL     = "/nitro/v1/config/botpolicylabel_botpolicy_binding"
	botPolicyLabelPolicyBindingBindingURL = "/nitro/v1/config/botpolicylabel_policybinding_binding"
	botProfileURL                         = "/nitro/v1/config/botprofile"
	botProfileBindingURL                  = "/nitro/v1/config/botprofile_binding"
	botProfileBlacklistBindingURL         = "/nitro/v1/config/botprofile_blacklist_binding"
	botProfileCaptchaBindingURL           = "/nitro/v1/config/botprofile_captcha_binding"
	botProfileIPReputationBindingURL      = "/nitro/v1/config/botprofile_ipreputation_binding"
	botProfileRatelimitBindingURL         = "/nitro/v1/config/botprofile_ratelimit_binding"
	botProfileTPSBindingURL               = "/nitro/v1/config/botprofile_tps_binding"
	botProfileWhitelistBindingURL         = "/nitro/v1/config/botprofile_whitelist_binding"
	botSettingsURL                        = "/nitro/v1/config/botsettings"
	botSignatureURL                       = "/nitro/v1/config/botsignature"
)

// Bot Management.
// https://developer-docs.netscaler.com/en-us/adc-nitro-api/current-release/configuration/bot/bot
type BotService struct {
	client *Client
}

// botglobal_binding
func (s *BotService) GetBotGlobalBinding() {}

// botglobal_botpolicy_binding
func (s *BotService) AddBotGlobalBotPolicyBinding()    {}
func (s *BotService) DeleteBotGlobalBotPolicyBinding() {}
func (s *BotService) GetBotGlobalBotPolicyBinding()    {}
func (s *BotService) CountBotGlobalBotPolicyBinding()  {}

// botpolicy
func (s *BotService) AddBotPolicy()    {}
func (s *BotService) DeleteBotPolicy() {}
func (s *BotService) UpdateBotPolicy() {}
func (s *BotService) UnsetBotPolicy()  {}
func (s *BotService) GetAllBotPolicy() {}
func (s *BotService) GetBotPolicy()    {}
func (s *BotService) CountBotPolicy()  {}
func (s *BotService) RenameBotPolicy() {}

// botpolicylabel
func (s *BotService) AddBotPolicyLabel()    {}
func (s *BotService) DeleteBotPolicyLabel() {}
func (s *BotService) GetAllBotPolicyLabel() {}
func (s *BotService) GetBotPolicyLabel()    {}
func (s *BotService) CountBotPolicyLabel()  {}
func (s *BotService) RenameBotPolicyLabel() {}

// botpolicylabel_binding
func (s *BotService) GetAllBotPolicyLabelBinding() {}
func (s *BotService) GetBotPolicyLabelBinding()    {}

// botpolicylabel_botpolicy_binding
func (s *BotService) AddBotPolicyLabelBotPolicyBinding()    {}
func (s *BotService) DeleteBotPolicyLabelBotPolicyBinding() {}
func (s *BotService) GetAllBotPolicyLabelBotPolicyBinding() {}
func (s *BotService) GetBotPolicyLabelBotPolicyBinding()    {}
func (s *BotService) CountBotPolicyLabelBotPolicyBinding()  {}

// botpolicylabel_policybinding_binding
func (s *BotService) GetAllBotPolicyLabelPolicyBindingBinding() {}
func (s *BotService) GetBotPolicyLabelPolicyBindingBinding()    {}
func (s *BotService) CountBotPolicyLabelPolicyBindingBinding()  {}

// botpolcy_binding
func (s *BotService) GetAllBotPolicyBinding() {}
func (s *BotService) GetBotPolicyBinding()    {}

// botpolicy_botglobal_binding
func (s *BotService) GetAllBotPolicyBotGlobalBinding() {}
func (s *BotService) GetBotPolicyBotGlobalBinding()    {}
func (s *BotService) CountBotPolicyBotGlobalBinding()  {}

// gotpolicy_botpolicylabel_binding
func (s *BotService) GetAllBotPolicyBotPolicyLabelBinding() {}
func (s *BotService) GetBotPolicyBotPolicyLabelBinding()    {}
func (s *BotService) CountBotPolicyBotPolicyLabelBinding()  {}

// botpolicy_csvserver_binding
func (s *BotService) GetAllBotPolicyCSVServerBinding() {}
func (s *BotService) GetBotPolicyCSVServerBinding()    {}
func (s *BotService) CountBotPolicyCSVServerBinding()  {}

// botpolicy_lbvserver_binding
func (s *BotService) GetAllBotPolicyLBVServerBinding() {}
func (s *BotService) GetBotPolicyLBVServerBinding()    {}
func (s *BotService) CountBotPolicyLBVServerBinding()  {}

// botprofile
func (s *BotService) AddBotProfile()    {}
func (s *BotService) DeleteBotProfile() {}
func (s *BotService) UpdateBotProfile() {}
func (s *BotService) UnsetBotProfile()  {}
func (s *BotService) GetAllBotProfile() {}
func (s *BotService) GetBotProfile()    {}
func (s *BotService) CountBotProfile()  {}

// botprofile_binding
func (s *BotService) GetAllBotProfileBinding() {}
func (s *BotService) GetBotProfileBinding()    {}

// botprofile_blacklist_binding
func (s *BotService) AddBotProfileBlacklistBinding()    {}
func (s *BotService) DeleteBotProfileBlacklistBinding() {}
func (s *BotService) GetAllBotProfileBlacklistBinding() {}
func (s *BotService) GetBotProfileBlacklistBinding()    {}
func (s *BotService) CountBotProfileBlacklistBinding()  {}

// botprofile_captcha_binding
func (s *BotService) AddBotProfileCaptchaBinding()    {}
func (s *BotService) DeleteBotProfileCaptchaBinding() {}
func (s *BotService) GetAllBotProfileCaptchaBinding() {}
func (s *BotService) GetBotProfileCaptchaBinding()    {}
func (s *BotService) CountBotProfileCaptchaBinding()  {}

// botprofile_ipreputation_binding
func (s *BotService) AddBotProfileIPReputationBinding()    {}
func (s *BotService) DeleteBotProfileIPReputationBinding() {}
func (s *BotService) GetAllBotProfileIPReputationBinding() {}
func (s *BotService) GetBotProfileIPReputationBinding()    {}
func (s *BotService) CountBotProfileIPReputationBinding()  {}

// botprofile_ratelimit_binding
func (s *BotService) AddBotProfileRatelimitBinding()    {}
func (s *BotService) DeleteBotProfileRatelimitBinding() {}
func (s *BotService) GetAllBotProfileRatelimitBinding() {}
func (s *BotService) GetBotProfileRatelimitBinding()    {}
func (s *BotService) CountBotProfileRatelimitBinding()  {}

// botprofile_tps_binding
func (s *BotService) AddBotProfileTPSBinding()    {}
func (s *BotService) DeleteBotProfileTPSBinding() {}
func (s *BotService) GetAllBotProfileTPSBinding() {}
func (s *BotService) GetBotProfileTPSBinding()    {}
func (s *BotService) CountBotProfileTPSBinding()  {}

// botprofile_whitelist_binding
func (s *BotService) AddBotProfileWhitelistBinding()    {}
func (s *BotService) DeleteBotProfileWhitelistBinding() {}
func (s *BotService) GetAllBotProfileWhitelistBinding() {}
func (s *BotService) GetBotProfileWhitelistBinding()    {}
func (s *BotService) CountBotProfileWhitelistBinding()  {}

// botsettings
func (s *BotService) UpdateBotSettings() {}
func (s *BotService) UnsetBotSettings()  {}
func (s *BotService) GetAllBotSettings() {}

// botsignature
func (s *BotService) ImportBotSignature() {}
func (s *BotService) DeleteBotSignature() {}
func (s *BotService) ChangeBotSignature() {}
func (s *BotService) GetAllBotSignature() {}
func (s *BotService) GetBotSignature()    {}
